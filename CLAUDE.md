# CLAUDE.md

Guidance for working on lightning — a code generator that emits fast,
allocation-light `json.Unmarshaler` implementations.

## Layout

- `main.go` — the generator (`package main`). Reads struct defs, emits
  `*_unmarshal.go`. Key bits: `field`/`sliceDecoder`/`mapDecoder` build decoders;
  `slicePresize` decides presizing. Behavior is selected by **per-type
  `//lightning:` comment directives** (parsed by `hasDirective`, which tolerates
  whitespace anywhere in the directive — `//lightning:compact`,
  `// lightning:compact ` and `// lightning: compact` are all equivalent, via
  `strings.Join(strings.Fields(...), "")`; collected into
  `compactTypes`/`nocopyTypes`/`destructiveTypes`): `//lightning:compact`
  (`g.compact`/`g.skipWS`, compile-time elision of inter-token `SkipWS`),
  `//lightning:destructive` (`g.destructive`/`g.scalar`, the type's nocopy strings
  unescape into the input buffer instead of allocating — see below; implies nocopy),
  and `//lightning:nocopy` (`g.nocopy`, a slice/map root aliases its keys/elements).
  The per-root loop sets the working `g.compact`/`g.destructive`/`g.nocopy` from
  those sets. `g.cmark`/`g.csuf` keep compact/destructive variants distinct in memo
  keys / function names (the `Destructive` suffix); nocopy variants are already
  distinguished by the `nocopy` decoder param / `NoCopy` suffix.
- `pkg/unstable` — the runtime scanning primitives the generated decoders call, plus
  the handful exported for the `pkg/json` toolkit (`SkipWS`/`SkipWSCompact`/`SkipValue`/
  `SkipString`/`ReadKey`/`DecodeValue`/`UnescapeString`/`ParseFloat` and the `Err*`
  sentinels). This is where almost all performance work happens. Split into topical
  files: `read.go` (the `Read*` readers), `batch.go` (the batched scalar-array
  readers), `skip.go`, `skipfast.go` (+ `skipfast_{amd64,arm64,noasm}`,
  the SIMD container skip), `count.go` (slice-presize counters), `numeric.go` (`scanFloat`
  + Eisel-Lemire), `string.go` (unescape/`Unwrap`), `date.go`/`time.go`, `any.go` (the
  dynamic `DecodeValue`), and the `simd_*` SIMD kernels; `unstable.go` holds the rest.
- `pkg/json` — small public API over the scanner, **implemented here** (not just
  wrappers) on the exported pkg/unstable primitives: `get.go` holds the read toolkit
  `Get`/`GetMany`/`GetPaths`/`ObjectEach` (+ `*Compact`) — `GetPaths` pulls several
  *nested* paths in one prefix-sharing pass (the multi-path form of `Get`); `set.go`
  holds `Set`/`SetMany`/`SetPaths` (`SetPaths` edits/creates several nested paths in
  one rewrite); `strip_defaults.go` holds `StripDefaults`; `json.go` keeps the
  decode-internal-bound wrappers `DecodeAny`/`UnescapeString`/`ParseFloat` (they need
  private `decodeEscaped`/`scanFloat`, so they stay in pkg/unstable). `EscapeString` lives
  in `escape.go`.
- `bench/` — separate module (keeps easyjson/sonic deps out of the main module).
  `run_bench.sh` regenerates decoders and benchmarks lightning vs
  encoding/json, easyjson, and bytedance/sonic. `bench/large-json/input.json` is
  the ~8 MB GeoJSON, gitignored and downloaded from `input.url`.

## Benchmarking — read this before claiming any speedup

**`run_bench.sh` runs each benchmark ONCE; those numbers are noisy** (first-run
GC/cache effects can make the faster decoder look slower — sonic "beat" lightning
on large-json in a single run while being ~17% slower over repeated runs). The
machine also drifts ~5–8% over minutes.

**Always measure with interleaved A/B**: `go test -c` a binary for each variant
(base via `git stash`, opt from the tree), alternate running them, feed both to
`~/go/bin/benchstat`. Build inside `bench/`; run each from its data dir (it reads
`input.json` relative to CWD). Treat <~2% as noise. For lightning-vs-competitor,
run `-count=8+` and compare medians, not the single run.

Hot functions are sensitive to **linker/code-alignment**: adding code (e.g. a
lookup table) can shift a tight micro-benchmark like `float-array` a few % with no
change to its executed instructions. Keep hot paths (`scanFloat`'s Clinger loop)
byte-identical when adding cold paths; push new logic out-of-line.

## Performance architecture (the load-bearing designs)

- **Float parsing** tiers in `scanFloat`: Clinger (exact mantissa <2^53, |exp|≤22,
  one mul/div) → **Eisel-Lemire** (`eiselLemire64` + generated `powers_table.go`,
  for mantissa ≥2^53 or |exp|>22) → `strconv`. EL is inline in `scanFloat` (no
  rescan) and bit-identical to strconv when it returns ok; guarded by the
  differential fuzz `TestParseFloatMatchesStrconv`. Don't remove EL.
- **Leading fraction zeros don't consume the 19-digit budget** in `scanFloat`:
  while `mant == 0`, leading `0`s after the decimal point (the `000` of
  `0.000698…`) only shift `exp`; they are skipped before the digit loops so they
  aren't counted toward the significant-digit total. Without this a small decimal
  like `0.0006988752666567719` (16 significant digits, easily exact) counts as 20
  digits, trips the `digits > mdigits` overflow guard, and falls all the way to
  `strconv` instead of Clinger/EL. Real-world win: golang_source −16.6% (its
  `cl_weight` weights are all sub-thousandth decimals); no regression elsewhere.
  A zero *between* nonzero digits stays significant (the guard is `mant == 0`).
- **SWAR fractional digits** in `scanFloat`: the fraction loop folds **four bytes
  at a time** (`is4Digits`/`parse4Digits`, the simdjson bit trick) instead of one
  `mant*10+d` per byte, with the 1–3 digit tail dropping to the scalar loop. Just
  a flat 4-byte loop — no 8-byte chunk, no trailing chunk. Net vs scalar:
  canada_geometry −5.6%, large-json −3%, float-array-slow −13%, golang_source −1%,
  and only +1.4% on the synthetic `float-array`. An earlier version added 8-byte
  runs (`is8Digits`/`parse8Digits`) for long fractions plus a trailing 4-byte; a
  direct A/B showed that machinery was **statistically tied** with the flat 4-byte
  loop on the real-world long-fraction cases (canada/large-json/golang_source —
  their 14–15-digit fractions still fold mostly in SWAR either way) while *costing*
  2.56% on `float-array`; it won only on `float-array-slow`'s 16-digit synthetic
  mantissas (+1.4%). So the 8-byte path was net-negative once its float-array cost
  is counted — dropped for the simpler, smaller, faster-on-balance 4-only loop.
- **SWAR integer digits** in `ReadInt64OrNull`/`ReadUint64OrNull`: same flat
  4-byte fold (`is4Digits`/`parse4Digits`) as the float fraction, scalar tail for
  the last 1-3 digits. `n*10000+v` is bit-identical to the scalar `n*10+d` chain,
  overflow wrap included. Win comes from long IDs/timestamps: golang_source −2.4%
  (10-digit ints), twitter_status −1.3% (18-digit IDs). No regression on short-int
  workloads (synthea/cloudflare flat) — unlike float-array, the per-int SWAR
  overhead is diluted by the surrounding string/object work. citm_catalog is flat
  despite the most 9-digit IDs: its bottleneck is key reading and map building, not
  int parsing.
- **Batched scalar-array readers** (`batch.go`: `DecodeFloat64Slice`, the generic
  `DecodeIntSlice`/`DecodeUintSlice`, and the fixed-size `DecodeFloat64Array`/
  `DecodeIntArray`/`DecodeUintArray`). The generated per-element loop paid a non-inlinable reader
  call per number — two frames for floats (`ReadFloat64OrNull` → `scanFloat`) —
  plus its own append/branch machinery; ~18% of the canada profile was that
  dispatch. The generator (`batchSliceFn`/`batchArrayFn` in `main.go`) now routes
  any slice or fixed-size-array field whose element is a bare float64/int/uint
  kind to these pkg/unstable loops, which call the private `scanFloat` directly
  (one frame per float) and inline the SWAR digit fold (no call at all per int).
  Presize (`CountArrayScalars`, only when `*out == nil`) and null handling live
  inside; semantics match the generated loop exactly — null root → nil slice /
  untouched fixed array, null element → zero, overflow wrap, tolerated truncated
  fraction — locked by parity tests against `ReadInt64OrNull` in `batch_test.go`.
  float32/bool elements keep the generated loop. Interleaved A/B (n=8, amd64):
  **numbers −9.2%, float-array −8.5%, marine_ik −6.4%, float-array-slow −5.4%,
  mesh −5.3%, mesh_pretty −3.7%**; canada/canada_geometry/large-json flat (their
  `[N]float64` ring points are scanFloat-bound, so the saved frame is a small
  fraction) and citm/cloudflare/twitter/synthea flat — no regressions. Toolchain
  note discovered here: the Go 1.25 inliner now inlines functions *with loops*
  (`SkipWS` cost 16, `SkipWSRun` 62), so `SkipWSRun` inlines straight into these
  readers — but a tiny `skipws` helper wrapping the two-compare fast path +
  `SkipWSRun` exceeds the budget (cost 97 > 80), which is why the whitespace fast
  path is expanded manually at each site, generator-style. The slice readers work
  on a **local copy of the slice header** (`s := *out`), not through the pointer:
  the compiler cannot prove `*out` doesn't alias `data`, so `*out = append(...)`
  reloaded ptr/len/cap and stored the new len every element; the local keeps the
  header in registers across the call-free int loop, at the price of writing
  `*out = s` back on *every* return (errors included — the parity tests lock
  partial progress on error).
- **Named-struct slice presize parity.** `slicePresize`'s `*ast.Ident` case now
  resolves a named struct element through `g.structTypes` and applies the same
  `isFlatScalarStringStruct`/`structSkipIsCheap` decision as an anonymous struct
  element; previously a schema that *named* a shared record type (a FHIR coding)
  silently lost the `CountArrayObjects`/`CountArrayElements` presize its inline
  twin received. synthea_fhir (the only bench with named struct slices):
  allocs/op −1.8%, B/op −0.3%, time statistically flat (p=0.065; its coding
  arrays are almost all single-element, so the win is bounded there).
- **Escaped-string decoding** (`decodeEscaped` / `readUnicodeEscape` /
  `decodeStringEscaped`) — three things make densely-escaped text fast
  (`\uXXXX`-per-character CJK, the twitterescaped workload): **(1)** the
  literal-run scan is skipped when already sitting on a `\` or `"` — consecutive
  escapes land on a `\` every other byte, so calling the SSE2
  `indexCloseOrEscape` each time just to find `\` at offset 0 was pure call
  overhead; **(2)** `readUnicodeEscape` parses the four hex digits through a
  `hexNibble` table of `uint32` entries (nibble value, or the invalid marker
  `1<<16`) and **inlines into `decodeEscaped`'s loop** (cost 65): the shifts fold
  into the combine — `t[a]<<12|t[b]<<8|t[c]<<4|t[d]`, one `>= 1<<16` compare
  validates all four digits — which is what pulled it under the budget (the
  earlier `uint8`/`0xFF` form cost 89 and was the loop's only non-inlined call,
  a frame per `\uXXXX`). The marker must sit at bit 16: an `0xFF`-style marker
  is wrong because `0xFF<<8 = 0xFF00 < 1<<16` slips through the combined test.
  Measured −8.8% on a dense-`\uXXXX` body; **(2b)** the escape *dispatch* is one
  `unescByte[256]` load for the eight single-byte escapes with `'u'` in the
  fallback — a Go `switch` lowers to a comparison tree that reached `'u'`, the
  hottest case on unicode text, *last*. Another −15.3% dense (combined with (2):
  **dense −22.9%, mixed −1.9%, sparse flat**, micro A/B n=8 M2); **(3)** the
  buffer cap hint in `decodeStringEscaped` finds the body end
  with a plain `bytes.IndexByte('"')` (one vectorized pass) instead of the
  escape-aware `skipString`, which stops at *every* backslash and re-scans the
  whole string — and the scan starts **at the first escape `i`, not at `start`**
  (the prefix `data[start:i]` is already proven clean by the caller's scan, so
  re-scanning it was pure waste). A `"` not preceded by `\` is definitively the
  unescaped close, so
  its offset is the exact body length (decoded ≤ escaped, always); only when the
  found `"` *is* preceded by `\` (a possible `\"`, rare) does it fall back to
  `skipString` for the true end — and the dense-escape strings that made
  `skipString` costly never hit that branch. Net: twitterescaped −33%, no
  regression on twitter_status, string_unicode, citm, golang_source, synthea, or
  cloudflare. **(4)** both return points hand out the scratch buffer with
  `unsafeStr(buf)`, not `string(buf)`: `buf` is freshly `make`-allocated by
  `decodeStringEscaped` (the only quoted caller) and never retained or mutated
  afterward, so `string(buf)` was pure waste — a *second* allocation plus a
  `slicebytetostring` memmove per escaped value, leaving the scratch buffer as
  immediate garbage. Aliasing the buffer into the result string instead is one
  alloc, not two, and no copy. The unquoted (`UnescapeString`) path already did
  this; the hot JSON-scanner (quoted) path was still copying. Net on
  escaped-string-heavy input: **gsoc_2018 −24% time, −46% B/op, −36% allocs**
  (4704→2995 — it had trailed sonic on amd64, now beats it ~28%), twitter_status
  −6%, string_unicode −7%, twitterescaped −4.5%; flat where escapes are rare
  (golang_source, citm).
- **`//lightning:destructive` — in-place unescape** (`ReadStringDestructiveOrNull`,
  `g.destructive`). A nocopy string still *aliases* the input when it has no escapes;
  the remaining per-string allocation is the scratch buffer `decodeStringEscaped`
  makes for an *escaped* value (it can't alias an escaped body). Under this directive,
  the type's nocopy strings are instead unescaped **into the input buffer**:
  `ReadStringDestructiveOrNull` hands `decodeEscaped` a `buf := data[i:i:len(data)]`
  aliasing the body, so the appends write through into `data` and the result aliases
  it — zero allocation. Safe because unescaping only *shrinks* (every escape is ≥2
  input bytes → ≤3 output bytes, `\uXXXX` 6→≤3), so the write cursor always trails
  the read cursor and never clobbers an unconsumed byte; cap is the document tail so
  `append` never reallocates away from `data`, and the closing quote (which the write
  never reaches) still bounds the value. It **destroys the input** — every escaped
  string's bytes are overwritten and any overlapping alias is invalidated — so it is
  opt-in (the directive name says so) for callers that own the buffer and discard it.
  It upgrades the type's `nocopy` string leaves (raw/number aliases are verbatim, no
  escapes); escape-free input is byte-identical to nocopy. Wins (with a per-iteration
  input-restore copy real usage omits): **gsoc_2018 −41% time / −86% B/op / −57%
  allocs**, twitterescaped −9.5% / −29% / −64%. Distinct decoder variants vs the
  plain/compact forms via `g.cmark`/`g.csuf` (`Destructive` suffix). Covered by
  conformance `TestDestructiveDirective` and the `destructive` arm of pkg/unstable's
  `TestReadStringOrNull`.
- **`//lightning:arena` — chunked backings for small scalar slices**
  (`unstable.Arena` + `arenaCarve` in `pkg/unstable/arena.go`, the
  `Decode*SliceArena` reader twins in `batch.go`, `g.arena`/`arenaParam`/
  `arenaArg` threading in `main.go`). Documents shaped like marine_ik hold tens
  of thousands of 3–4-element `[]float64` fields; each decoded into its own
  exact-fit `make` backing, `DecodeFloat64Slice` was **95% of allocated
  objects**. Under the directive, each `UnmarshalJSON` declares a local
  `unstable.Arena` and threads `a *unstable.Arena` through every decoder of the
  variant (uniformly — unlike depth's cycle gating, the directive itself is the
  gate; the batch slice readers are rerouted to `...Arena` twins that carve the
  presized backing from 4 KiB chunks instead of one `make` per slice). Safety
  is by construction: `arenaCarve` is constrained to **noscan element kinds**
  (chunks are `[]byte`, which the GC doesn't scan — pointerful types must never
  live there); the cursor bumps past the **full capacity**, so carves are
  exclusive and a caller's later `append` can never clobber a neighbour (exact
  scalar counts leave `len == cap`, so such an append reallocates to the heap
  anyway); carve offsets stay 8-aligned; chunks are make-zeroed and regions
  never reused, so `s[len:cap]` reads zeros exactly like a `make` backing; and
  backings over `arenaMaxCarve` (512 B) fall back to a direct `make`, keeping
  both chunk waste and the pinning trade-off bounded to small slices. The
  trade-off — a surviving small slice pins its ~4 KiB chunk — is why it is a
  directive and not the default. Measured (interleaved A/B, n=24, pinned Zen 4):
  **marine_ik allocs/op −94.9% (29 356 → 1 504) and time −2.80% (p=0.001);
  mesh allocs/op −99.2% (3 618 → 30) and time −3.30% (p=0.000)**; B/op +0.2–0.4%
  (chunk-tail waste); numbers exactly flat (its one big array bypasses the
  threshold by design). Non-arena schemas are **byte-identical** (dual-generator
  diff over all 60 bench inputs) and the shared-body restructure of the batch
  readers costs nothing (`Decode*Slice` became inlinable wrappers over a private
  body with a nil-arena parameter — regression A/B over marine_ik/mesh/numbers/
  float-array/canada/cloudflare/citm all flat, n=8). **Calibration lesson worth
  keeping:** the old "~18% target" here was sized from mallocgc's cumulative
  profile share; the isolated micro (`BenchmarkDecodeSmallSlices`, −20.7%/slice)
  over-predicted the same way. Most of that profile share is bytes-proportional
  work the arena deliberately keeps (chunk zeroing costs the same bytes as the
  makes it replaced, and GC assist paces on bytes) — the removable part is only
  the per-object malloc fast path, ~12 ns on Zen 4. What the wall-clock number
  understates is the whole-program effect: a service decoding at rate carries
  29k → 1.5k objects/decode into every GC mark, which the benchmark loop's
  short windows underweight. Locked by `TestArenaSliceParity`/`Exclusive`/
  `AlignmentAndKinds`/`ThresholdBypass`/`ReuseKeepsBacking`/`RandomizedParity`
  (pkg/unstable), conformance `TestArenaDirective` (stdlib parity, sibling
  safety after append, reuse, and the arena×recursive combo `ArenaTree`, whose
  decoders thread depth *and* arena), and the standing micro
  `BenchmarkDecodeSmallSlices`. The bench harness generates a `data_arena.go`
  twin per case (like destructive) and a `BenchmarkLightningArena` row.
- **`CountArrayElements`** (slice presize) skips each element with `SkipValue`
  (vectorized via `indexStructural`), not a byte-by-byte depth walk — but **gives
  up the per-element walk after `countSampleCap` (64) elements** and extrapolates
  the total from the bytes the sample spans (`n * (approxEnd − open) / sampled`,
  with `approxEnd` the first `]` via `bytes.IndexByte`). A huge uniform array
  (apache_builds' **875** `{name,url,color}` job objects, all strings, so the
  cheap counters below don't apply) is then sized from a 64-element sample instead
  of a full `skipObject`-per-element pass that re-scans every URL: apache_builds
  **−41%**, allocs unchanged (the estimate, 912 vs 875, still presizes in one
  alloc). The first `]` is at or before the true close (a `]` inside a string only
  moves it earlier) and the result is only ever a presize hint — a wrong count
  mis-sizes the slice, never misdecodes — so an over/under-estimate is harmless;
  arrays ≤ 64 elements still get an exact count. No regression on citm, large-json,
  golang_source, synthea, twitter, marine_ik, payload_large, update_center.
  Element types whose JSON can hold no comma or bracket use the much cheaper
  `CountArrayScalars`
  (find `]`, count commas — two vectorized scans, no per-element work): bare
  numbers/bools, `json.Number`, **and `time.Time`** — an RFC 3339 / Unix-timestamp
  value never contains a `,` or `]`, so a `[]time.Time` sizes by comma count. That
  avoids a `skipString` over every element (which re-scanned each date string just
  to count it, doubling the per-date work): time-array −16%.
- **`CountArrayObjects`** extends that to a *flat struct of number/bool/**string**
  fields* (`isFlatScalarStringStruct`): its JSON holds no `[`, `]` or nested `{` of
  its own, so the array's `]` is the first `]` and the element count is the number
  of `{` before it — two vectorized scans instead of a `SkipValue` per struct.
  Number/bool-only is *exact* (citm_catalog price entries: −6%, no alloc change).
  With string fields it is a presize *hint* — a `[`/`]`/`{` inside a string value
  could mis-size the slice, but a miscount only mis-sizes, never misdecodes — and
  it pays where the array holds many small `{name,version}`-style records:
  update_center dependencies/developers and **apache_builds jobs −2.7%** (cheaper
  than the cap-and-extrapolate `CountArrayElements` it replaces, since two
  `IndexByte`/`Count` passes beat skipping 64 elements). A *nested*
  struct/array/pointer/map field disqualifies it (its JSON carries real brackets);
  those keep `CountArrayElements`. All these counters are presize hints.
  (Tried-and-rejected for update_center: presizing the `map[string]struct` plugins
  map saves the ~7 rehashes of its large struct values, −5% — but counting its 654
  members needs a depth-aware scan or a matching-`}` extent for extrapolation, both
  ~as costly as the rehash they'd save. The `//lightning:nocopy`-equivalent
  `,nocopy` on the plugins map *field* does land though: aliasing the 654 plugin-name
  keys is −21% allocs / −2.5% time.)
- **`slicePresize`** skips presize in two cases. **(1)** When a struct element
  transitively holds a *multi-dimensional* slice (e.g. GeoJSON `[][][]float64`):
  counting it would deep-scan every element's bulk for only ~log2(n) reallocs
  saved. (The test is `structSkipIsCheap`, which walks the element's fields
  transitively with a `seen` map for cycle safety and answers the more general
  question "is skipping one of these cheap"; an earlier `hasMultiDimSlice` helper
  named here until 2026-08-09 no longer exists.) **(2)** When the slice's element is itself a *fixed-size
  array* `[N]T` (the `ArrayType` case is gated on `t.Len == nil`, so a `[3]float64`
  coordinate point in large-json's `[][3]float64` rings disqualifies the ring):
  presizing such a ring runs `CountArrayElements`, which descends through every
  coordinate number — work the element decoders then repeat — and zeroes the
  presized backing, costing more than the doubling growth it saves. Letting the
  ring append instead is **−8.8% on large-json** and makes **canada_geometry beat
  sonic** (its `[][2]float64` rings); the alloc *count* rises (large-json
  ~40k→~72k as the rings double) but wall-time drops and B/op stays *below* sonic.
  Flat / string / 1-D-slice records (Cloudflare-style) keep presize and their low
  B/op — a slice whose element is a bare scalar (`[]float64`), string, struct, or
  time is still counted (those land in the `Ident`/`StructType`/`SelectorExpr`
  cases, not the fixed-array `ArrayType` branch).
- **Static first-append capacity hint for un-presized slices** (`sliceDecoder`,
  the `presize == ""` case). Every slice `slicePresize` declines to count (struct
  elements nesting a slice/map/any — the citm areas tree — plus the `[][N]T`
  rings and multi-dim slices above) used to grow from nil by pure append-doubling
  (1→2→4→…); an M2 alloc profile put that one `append` line at **69% of
  citm_catalog's allocated objects** (~4k tiny growslice allocs/op — its areas
  arrays are mostly 6–11 elements of a 32-byte struct). The first append now
  allocates `max(4, 256/sizeof(elem))` capacity — a compile-time constant via
  `unsafe.Sizeof`, ~256 bytes of elements — which is *not* the rejected counting
  presize: no extra scan, no full-array memclr; a too-large hint wastes only
  spare cap, a too-small one regrows as before. `[]` still yields nil (the hint
  fires only when an element exists). (The hint's `*out == nil` test still works
  after the length reset described in **Slice reuse replaces** below — `nil[:0]`
  is still nil — so a fresh decode gets the hint and a reused slice appends into
  the capacity it already has.) Interleaved
  A/B (n=8, M2): **citm_catalog −7.7%, canada_geometry −5.3%, golang_source
  −2.7%, canada −1.7%, marine_ik −1.1%**, mesh/large-json/cloudflare flat — no
  time regressions; allocs/op geomean **−41%** (citm −52%, canada −62%,
  large-json −60%). B/op trade: mostly down (citm −20%), but the rings carry
  spare cap (canada_geometry +13%, large-json +4.4% B/op) — the same
  time-over-B/op trade the ring presize-skip already accepted.
- **SIMD scanners** in `simd_amd64.s`/`simd_arm64.s`: `indexCloseOrEscape`
  (next `"`/`\`) and `indexStructural` (next `{}[]"`). Both arches classify the 5
  target bytes with simdjson's **shuffle trick** — two table lookups
  (`structLo[lowNibble] & structHi[highNibble] != 0`) instead of five
  compare+or. Two bits suffice: one for `"` (lo 0x2/hi 0x2), one for the
  brackets/braces (lo 0xB|0xD / hi 0x5|0x7), so cross combos like `0x52` 'R'
  stay non-structural. amd64 (`VPSHUFB`) needs a trailing compare-to-zero +
  `VPMOVMSKB`+`NOT` (the movemask reads bit 7); arm64 (`TBL`) needs neither — its
  RBIT/CLZ recovery finds the first nonzero byte directly, so the loop is just
  `VAND`/`VTBL`/`VUSHR`/`VTBL`/`VAND` (5 ops vs the old 9). Fewer ops win on the
  throughput-bound skip loop where blocks run back-to-back: **amd64 skip-heavy −5%**,
  no regression on citm / large-json / canada (whose `SkipValue`/presize use
  early-exits within a block via the scalar prescan, so it never reaches the loop).
  arm64 mirrors the change (correctness verified under qemu — full pkg/unstable test suite +
  an exhaustive 0–255 × boundary-offset differential); its speedup is unmeasured
  here (qemu isn't cycle-accurate) and wants a real-arm64/CI run to confirm. The
  string scanner
  (`indexQuoteOrBackslashSSE2`) is a **length-adaptive hybrid**: the first 32-byte
  block is SSE2, so short keys/values return with no AVX2 state and no VZEROUPPER;
  only a string whose first 32 bytes hold no `"`/`\` (a long text field) switches
  to an AVX2 tail loop — one 32-byte compare per iteration vs SSE2's two 16-byte —
  amortizing the lone VZEROUPPER over the rest. Short-string benches (cloudflare)
  are unchanged; long strings win (string_unicode −9%, twitter/large-json ~−1%).
  Don't make it pure-AVX2: the per-call VZEROUPPER regresses the short-string
  common case (the reason SSE2 was chosen originally).
- **Date parsing**: `daysFromCivilCached` uses a year-start-days table (built from
  `daysFromCivil`) for 1970–2261; falls back to the general algorithm otherwise.
  `parseRFC3339`'s fractional-seconds loop accumulates at most nine digits (bounded
  so the per-digit `<9` test stays out of the loop) and scales to nanoseconds with
  one `pow10nano[fd]` multiply instead of a trailing `for fd<9 { nsec*=10 }` pad:
  time-array −1%.
- **SIMD in-string-mask container skip** (`skipfast.go` + `skipfast_amd64.s` /
  `skipfast_arm64.s`, the sonic-rs `skip_container` / JSONSki technique). `SkipValue`
  used to land on each structural byte with `indexStructural` and call `SkipString`
  *per string*, so skipping a string-heavy container paid N calls. `skipContainerFast`
  instead streams **64-byte** blocks: `maskBlock` (AVX2 / NEON) returns four uint64
  bitmaps — `"`, `\`, and only the container's *own* open/close brackets (it is told
  which via an `isArray` arg and branches to the `{`/`}` or `[`/`]` splats, so a
  stray bracket of the other type is never counted — which on *well-formed* input
  matches `skipObject`/`skipArray`, and on malformed input does **not**; see the
  divergence note below). 64 bytes/call (vs an
  earlier 32) halves the call/marshal/`VZEROUPPER` overhead and the
  `findEscaped`/`prefixXor` frequency; computing only 4 classes (not 6) cuts the
  per-block movemasks 12→8. A direct A/B (both builds 4-mask type-selected, only the
  block size differing) measured **64-byte ~5% faster than 32-byte** on amd64
  `GetSkipHeavy` (12.33 vs 12.95 µs, p=0.000): the fixed per-call costs outweigh
  64-byte's extra 32-byte load + the `SHLQ`/`ORQ` that folds two 32-bit movemasks
  into a uint64. (**asm gotcha**: `go vet` asmdecl does *not* validate `maskBlock`'s
  result offsets — a 32-byte `uint32`-return build silently miscounted because Go
  8-aligns the result block *after the `isArray bool` arg*, so the first return sits
  at +32, not +28. Verify `maskBlock`'s masks with a direct dump if you touch the
  signature; the live 64-byte form returns `uint64`s, which are 8-aligned anyway.
  **The reason, re-confirmed by experiment 2026-08-09 — do not "correct" this
  entry**: the four result stores go through the `CLASS(s, off)` macro, whose body
  ends `MOVQ AX, off(FP)`. asmdecl scans the assembly *source*, so in the macro
  definition it sees the token `off(FP)`, never the expanded `quote+32(FP)`, and
  has nothing to check. Sabotaging `CLASS(Y10, quote+32)` to `quote+28`, and
  replacing a named reference with a bare `40`, are both accepted in silence —
  while the *control*, shifting the non-macro `b_base+0(FP)` to `+8`, is reported
  at once (`invalid offset b_base+8(FP); expected b_base+0(FP)`). So asmdecl is
  running and is not the problem; macro-hidden FP references are. A probe on a
  hand-written maskBlock-shaped function *is* flagged and will mislead you into
  thinking the gotcha is stale — it is the macro, not the toolchain.)
  `findEscaped64` (simdjson's branchless odd-run detection)
  + `prefixXor64` (the carryless-multiply-by-ones done in six shift/XORs — **no
  PCLMULQDQ**, so the bit math is plain Go) build the *inside-string* mask; the
  bracket bitmaps are masked to bytes outside strings and balanced. Strings (keys and
  values) are absorbed into the bulk scan — no per-string call. **`SkipValue` is the
  dispatch**, so the win reaches every caller (`Get`/`GetMany`/`GetPaths`, `Set`,
  `CountArrayElements`, generated unknown-field skips) with no call-site change:
  objects always go fast (string keys dominate); an array goes fast only if its first
  element is `{`/`[`/`"` — a *scalar* array (`[1,2,…]`) stays on the current path,
  where one `indexStructural` scan already reaches the close and the mask path would
  only add per-block work. The probe is a heuristic; a wrong guess (or a miss on
  malformed input) only costs speed, never correctness — both paths are bracket
  balancers that return the identical end index on every **well-formed** value
  (50k-doc differential fuzz + truncation safety vs `skipObject`/`skipArray`).
  **They are not interchangeable off that set.** An earlier version of this entry
  claimed they "agree on every value the other accepts"; a 2026-08 audit falsified
  it, and `skipfast.go`'s header now enumerates three divergence classes — an
  unbalanced bracket of the *other* type (fast ignores, scalar descends), nesting
  past `MaxDepth` (fast is iterative and accepts, scalar now returns `ErrMaxDepth`),
  and a stray backslash *outside* a string (`findEscaped64` is pure bit math and
  cannot know it is not in a string, so the verdict flips on 64-byte-grid alignment
  alone: `{`+N spaces+`\}` is accepted at N=62, rejected at 63, accepted at 64).
  Because `SkipValue` picks the path by CPU feature, all three make its answer on
  **malformed** input host-dependent. Pinned by `TestSkipPathsDivergeOnMalformed` /
  `TestSkipBackslashAlignmentCliff` / `TestSkipDepthDivergence`, in the spirit of
  `TestValidDivergesFromStdlib`. **This is not the rejected two-stage
  feed** (below): the skip path has no typed stage-2, so the index-like scan *is* the
  work and the economics that sank two-stage do not apply. Wins: **`Get` end-to-end
  +105%** (skip-heavy doc, skipping 500 nested-object siblings: 27.9→13.6 µs),
  micro-skip −36 % (string object) / −49 % (number-valued object) / −79 % (array of
  records), flat on scalar arrays, zero allocs. Gated `fastSkipAvail = useAVX2`
  (amd64) / `true` (arm64, NEON baseline) / `false` (other, where the scalar
  `maskBlock` is slower than `indexStructural`). The arm64 NEON `maskBlock` builds the
  movemask NEON lacks (`PMOVMSKB`) by **weight-and-fold over an ADDP cascade**, computing
  each class's full 64-bit mask (all four 16-byte chunks of the block) with **one**
  vector→GP move. `VAND` each chunk's 0x00/0xFF compare with the `{1,2,4,…,128}`-repeated
  bit-weight vector (a matching lane becomes its bit value), then a four-step `VADDP`
  cascade — `P=ADDP(A0,A1)`, `Q=ADDP(A2,A3)`, `R=ADDP(P,Q)`, `S=ADDP(R,R)` — collapses
  each chunk's two 8-lane halves to one mask-byte and *packs all four chunks* into S's
  low eight bytes in order `[lo0,hi0,lo1,hi1,lo2,hi2,lo3,hi3]`, which is exactly the
  uint64 mask (chunk k at bit 16k), so a single `VMOV S.D[0]` lifts it (no per-half
  extract, no shift/OR stitching). `ADDP(Vn,Vm)` puts Vn's pairwise sums in the low half;
  a full 8-lane half sums to 255 so no byte overflows. This replaced (1) an `AND/MUL/LSR`
  byte-LSB gather with two `VMOV`s + two `MUL`s per 16-byte half, then (2) a per-half
  three-`VADDP` fold that still did one `VMOV` per 16-byte half (four per class for the
  64-byte block) plus `ORR`-shift stitching. The cascade drops a class from **27 → 13
  ops** (16→4 cross-domain `VMOV`s, 12→0 `ORR`s for the whole block). On **Apple M2**,
  successive: the per-half fold cut `BenchmarkGetSkipHeavy` −28% (30.8→22.2 µs); the
  cascade then cut **another −24…−33% on `BenchmarkSkipContainer`** (stringObj −32%,
  numberObj −33%, nestedMixed −28%; scalar arrays flat), dropping `maskBlock` from ~76%
  to ~45% of the skip profile (its bit-math sibling `skipContainerFast` is now the larger
  share). This directly attacks the skip path, which the bench-md comparison flags as
  arm64's worst lag vs amd64 (skip-heavy ~0.36 of amd64's speedup-over-stdlib) — the
  residual gap is intrinsic: NEON is 16-byte and has no `PMOVMSKB`, so even the cascade
  does more work than amd64's `VPMOVMSKB` over 32-byte AVX2.
- **amd64 whole-loop skip assembly (`skipBlocksAVX2`/`skipBlocksAVX512`) + Go-loop
  fast paths.** Profiling the fast skip showed `maskBlock` (the vector kernel) at
  only ~15%: the loop was **latency-bound on the loop-carried
  escaped→inStr→prevInString chain** (each block's in-string mask depends on the
  previous block's through `findEscaped64`'s add-carry plus `prefixXor64`'s six
  serial shift-XORs — those few ALU lines out-sample the SIMD), plus a full asm
  call, four results through memory, five splat loads and an `isArray` branch *per
  block*. Two layers of fix. **(1) Go-loop fast paths** (all arches, and the
  amd64 fallback): skip `findEscaped64` when `bslash|prevEscaped == 0` (escapes
  are rare), skip the prefix-XOR when the block has no unescaped quote (the mask
  is the carried `prevInString`), and replace the per-bit bracket walk with
  **popcount bulk updates** whenever the block cannot cross depth 0 (`cl == 0`,
  or `depth > popcount(cl)` — opens only raise depth, so the running minimum
  stays ≥ 1); only genuinely-might-close blocks walk bits. Alone: numberObj −22%,
  nestedMixed −23% (amd64 single-run; arm64 inherits, unmeasured). **(2) The
  whole block loop in assembly** (`useSkipBlocks`, amd64): splats loaded once,
  depth/prevEscaped/prevInString carried in registers, the same fast paths and
  popcount bulk in GP code, and the prefix XOR as **one `VPCLMULQDQ` carryless
  multiply by all-ones** (the reason the Go form is six shift-XORs is exactly
  that CLMUL isn't reachable from Go without a call). The shared per-block bit
  math is one `BLOCKTAIL` macro expanded into both variants (labels are
  function-scoped). The **AVX-512 variant** (`useSkipBlocks512`, needs
  AVX512BW) does one 64-byte load and a `VPCMPEQB`→k-mask→`KMOVQ` per class (2
  instructions vs AVX2's 7); measured on Zen 4: goloop→AVX2 −36…−45%,
  AVX2→AVX-512 **another −18…−30%** — it earns the gate. Net vs the maskBlock
  Go loop (interleaved n=8): **stringObj −62.8%, numberObj −66.9%, nestedMixed
  −58.8%**; scalar arrays flat (they keep the `indexStructural` path by design).
  The real-workload suite is **flat** — its unknown-field skips are small and
  skip-heavy's content is scalar-array-dominated (~49 GB/s on the
  `indexStructural` path already), so this win is for Get/GetPaths/Set-style
  callers skipping big object/record containers, which is what
  `BenchmarkSkipContainer` models; cloudflare's apparent +4% was rechecked at
  n=10 and is alignment noise (p=0.225). Gates
  `useAVX2 && HasPCLMULQDQ && HasBMI1 && HasPOPCNT` (all universal with AVX2 —
  correctness belts). **arm64 has the same whole-loop form** (`skipBlocksNEON`,
  unconditional — NEON is baseline): splats and the bit-weight vector hoisted
  out of the loop, state in registers, the CLASS/ADDP-cascade movemask per
  block, popcounts via `CNT`/`UADDLV` (no GP popcount on arm64), the bit walk
  via `RBIT`+`CLZ`. Deliberately **no PMULL prefix XOR**: the mask is in the GP
  domain by then (the escape math needs GP add-with-carry) and a GP→SIMD→GP
  round trip costs more on M-class cores than the shift-XOR chain — which
  arm64's shifted-register `EOR` does in six single instructions anyway. What
  the arm64 loop saves vs the Go loop is the per-block call, four results
  through memory and five per-call `VDUP` splats. Verified under qemu (full
  pkg/unstable + pkg/json suites, all variants) and natively on **Apple M2**
  (full suite + `TestSkipBlocksVariants`), where `BenchmarkSkipBlocksVariant`
  (n=8, benchstat) measures NEON-loop vs Go maskBlock loop **stringObj −34.0%,
  numberObj −34.7%, nestedMixed −28.7%** (geomean −32.5%, all p=0.000) —
  10–15 GB/s end-to-end on the object shapes, scalar arrays flat by design.
  Correctness: `TestSkipBlocksVariants` flips the dispatch flags and
  differentially tests **goloop, AVX2 and AVX-512 each** (goloop and NEON on
  arm64) against the scalar
  oracle over the random fuzz corpus plus `boundaryDocs()` — backslash runs of
  every parity crossing the 64-byte block boundary at every offset (the
  `prevEscaped` carry), quotes on the boundary, deep bracket runs, close-dense
  blocks, closes at exact block multiples — plus per-variant truncation safety;
  `BenchmarkSkipBlocksVariant` is the standing AVX2-vs-512-vs-Go comparison.
- **Un-presized slices grow at a flat 2×, not Go's damped 1.25×**
  (`unstable.GrowSlice` in `pkg/unstable/grow.go`, emitted by `sliceDecoder`'s
  `presize == ""` path). Bare `append` lets `runtime.nextslicecap` decide capacity,
  and it doubles only while cap is under **256 elements**, then grows
  `cap += (cap+768)>>2` ≈ 1.25×. The total bytes a growing slice allocates come to
  `final_cap × f/(f−1)`, so the damped regime allocates ~5× the final size and
  memmoves ~4× it, where a flat 2× allocates 2× and memmoves 1×. That is exactly
  where large-json's 10k-element `features` array and canada's long rings live —
  large-json's profile showed **memmove 8.8% + `memclrNoHeapPointersChunked` 5.6%**,
  i.e. the growth itself, not the parsing. The generated loop now does
  `if len(*out) == cap(*out) { *out = unstable.GrowSlice(*out) }` before the append.
  Interleaved A/B (n=10, pinned): **large-json −10.74% time / −24.70% B/op**
  (p=0.000), **canada −2.13% / −15.80%** (p=0.029), marine_ik −3.84% B/op,
  citm −0.30% B/op, all times flat elsewhere. **Arrays under 256 elements are
  untouched either way** — citm's areas (≤16), mesh, cloudflare are byte-identical —
  which is what makes this safe; the only cost is spare capacity on a mid-sized
  array (canada_geometry **+1.11% B/op**, time flat). **4× was measured and
  rejected**: it buys *no further time* on large-json or canada while pushing
  canada_geometry to **+7.10% B/op**, so 2× captures the whole win with the least
  waste. Note this is orthogonal to the rejected *counting* presizes and to the
  static first-append hint: it changes the growth **factor**, not whether a scan
  sizes the array. Locked by `TestGrowSlice` (length/contents preserved, capacity at
  least doubled either side of the 256-element threshold, no aliasing of the old
  backing).
- **All-spaces equality fast path in `SkipWSRun`** — the whitespace attempt that
  finally worked, after three that did not. The loop's per-word classify
  `ws := (g - w&^hi) &^ w & hi` answers "are all eight bytes `<= 0x20`", but inside
  an indentation run the overwhelmingly common word is **eight literal spaces**, so
  a single `w != sp` compare against the `0x2020…20` splat now guards it and the
  exact SWAR runs only for the other words. Sound by construction: equality with
  the splat is a *sufficient* test (every byte is exactly `0x20`), so the fast path
  can only skip work — the exact classify still decides every other word, including
  the run-terminating one, and no input changes acceptance. All-space word share of
  the corpus: citm 68%, marine_ik 66%, instruments 50%, mesh_pretty 49%, synthea 43%.
  **Why this is orthogonal to the three rejections** above — and the reason to keep
  it distinct from them: it changes neither the loop *width* (still 8 bytes, still
  the same number of loads, so "more loads than needed when the run ends mid-chunk"
  cannot apply), nor uses any vector instruction or call (so unamortised per-call
  setup cannot apply), and `SkipWSRun` still inlines (cost 62 → **66**, budget 80 —
  re-check `-gcflags=-m` after any edit here, the whole `g.skipWS` design depends on
  it). What it attacks is the *cost of classifying one word*, which none of the
  earlier attempts touched. Interleaved A/B (n=10, pinned to one core): **mesh_pretty
  −5.61%, instruments −4.42%, citm_catalog −4.11%** (all p≤0.002), synthea_fhir
  −3.2% (p=0.143); marine_ik flat despite 66% all-space words, because `SkipWSRun`
  is only ~2.5% of its profile. Compact cases (cloudflare, canada, large-json) are
  flat **by construction** — they never enter `SkipWSRun` — which is what makes this
  unusually safe to A/B: an apparent regression there can only be alignment noise.
  Equivalence to the byte loop is not left to the soundness argument:
  `TestSkipWSRunMatchesOracle` checks every byte value at every lane offset for
  every start offset, each whitespace byte as filler, `>= 0x80` bytes (which must
  *not* count as whitespace), and random mixtures, plus
  `FuzzSkipWSRunMatchesOracle` (2.4M execs). A tab-indented document takes the
  general path always and pays one extra compare per word; a second equality
  against `0x0909…` would cover it if that ever matters.

  **Companion micro-change, same loop: the mask is computed complemented.** The
  loop builds `nws := ^((g - w&^hi) &^ w) & hi` — a bit per lane that is *not*
  whitespace — rather than the whitespace mask, exploiting
  `(x & hi) ^ hi == (^x) & hi`. It pays off only at the run-terminating exit, which
  runs once per whitespace run: there is no `XOR` to invert the mask, and testing
  `nws != 0` *proves* the `TrailingZeros64` operand non-zero, so the compiler stops
  materialising 64 and `CMOVE`-ing it as the all-zero guard. Verified in the emitted
  code: the exit collapses to `TESTQ` / `BSFQ` / `SHRQ` with **zero** `CMOVQEQ` or
  `MOVL $64` remaining, and the inline cost drops 66 → **65**. Honest sizing: it is
  **below the 2% noise floor** — instruments −1.84% (p=0.001) is the only
  significant result; citm −0.8% (p=0.075), mesh_pretty −1.2% (p=0.105), synthea
  −1.6% (p=0.280), large-json and canada flat. Kept anyway because it is strictly
  less code, provably equivalent (the identity plus the exhaustive oracle test), and
  favourable in direction on every case measured — not because it is a measured win.
- **Slice reuse replaces, and the reset is guarded.** Every slice decoder — the
  generated `sliceDecoder` loop and `batch.go`'s three readers — used to start from
  `*out`, so decoding into a **non-nil** slice *appended* to it. `[1,2]` decoded
  twice into one value became `[1,2,1,2]`, silently and with no error, and a caller
  reusing a target to avoid allocation (the only reason to reuse one) grew it
  without bound. It contradicted `encoding/json`, which documents "Unmarshal resets
  the slice length to zero and then appends each element to the slice". **This was
  deliberate, not an oversight** — `TestDecodeFloat64SliceAppends` asserted "a
  non-nil `*out` is appended to, not reset" — which is why it survived; that test
  now locks the opposite. No benchmark caught it either, because every benchmark
  decodes into a freshly declared `var v Benchmark`. Found only by reading
  `DecodeFloat64Slice` while chasing marine_ik's allocation profile.
  The fix is `(*out)[:0]`, which keeps the backing array so reuse stays
  allocation-free. Two details worth keeping: **(1)** the generated reset is
  **guarded** — `if len(*out) != 0 { *out = (*out)[:0] }` — because an
  unconditional store of the 3-word slice header measured **cloudflare +1.26%
  (p=0.001, n=8)** for its three slice fields; the guard makes a fresh decode pay a
  load and a not-taken branch instead, and every slice-heavy case (cloudflare,
  marine_ik, citm_catalog, mesh, canada, large-json) is then **flat**. So this is a
  correctness fix that costs nothing. **(2)** Only slices were wrong: maps
  *merge* existing entries, which is exactly what `encoding/json` does (verified
  with different keys per round); fixed arrays are zeroed then index-filled; and a
  slice reached as a **map value** or through `lax` was already correct, since both
  decode through a fresh scratch variable. Locked by `TestSliceReuseReplaces` /
  `TestSliceReuseKeepsBacking` (conformance, incl. a named slice root and a nocopy
  slice) and the `batch.go` trio above; all fail against the pre-fix code.
  Codegen gotcha found here: the reset text sits inside a `fmt.Sprintf` template,
  so a `%` written into that comment is emitted into every generated file as
  `%!(MISSING)` — `go vet` catches it, which is why measurement notes live in
  `sliceDecoder`'s Go doc comment rather than in the emitted comment.
- **`Valid` is a grammar walk, not a decode** (`pkg/json/valid.go`). It used to be
  `_, err := decodeAny(data, false); return err == nil` — which built the entire
  `map[string]any`/`[]any` tree just to throw it away: **8065 ns / 13312 B / 201
  allocs** on a 1.1 KB record, i.e. *slower than `encoding/json.Valid`* (2760 ns,
  0 allocs), the one place this library lost to the stdlib. `validValue` now checks
  the document in place: **1220 ns, 0 allocs** (6.6× the old form, ~2.3× the
  stdlib). Three load-bearing choices. **(1)** It is a **flat loop with a bitset**,
  not recursion — one bit per open container (set = object, clear = array) in a
  fixed `[MaxDepth/64+1]uint64` local, so nesting costs bits instead of stack
  frames. Three `goto` labels are the states (`scanValue`/`scanKey`/`scanAfter`);
  all locals are declared up front because Go forbids a goto jumping over a
  declaration. The empty-container case is handled *at the open bracket*, which is
  precisely what lets `scanKey` reject a `}` as a trailing comma with no extra
  flag. **(2)** Its contract is **agreement with lightning's own decoder, not with
  `encoding/json`** — `Valid(data)` answers "will `DecodeAny`/a generated
  `UnmarshalJSON` accept this?", which is the useful question for a gate in front
  of them. So numbers go through the decoder's own `ReadFloat64OrNull` (same
  Clinger/EL/strconv tiers, same span consumed, same overflow rejection) rather
  than a reimplemented grammar check: agreement is by construction, and
  `1e309` is *invalid* here though `encoding/json.Valid` accepts it. Strings are
  the one hand-written part (the decoder's readers unescape, so they allocate) and
  deliberately mirror the scanners' leniency — `IndexCloseOrEscape` stops at `"`
  and `\` only, so a **raw control byte is accepted**, and unpaired surrogates are
  not checked. Locked by `FuzzValidMatchesDecodeAny` (20M execs, zero divergence)
  plus `TestValidDivergesFromStdlib`, which pins each deliberate disagreement so
  neither side drifts silently. **(3)** It needs its own `skipWSStrict`… **no** —
  it needs the opposite: it uses `unstable.SkipWS`, whose `<= 0x20` test is a
  deliberate one-compare-instead-of-four shortcut, *because* matching the decoder
  means inheriting that leniency (a NUL between tokens is whitespace here). An
  earlier draft aimed at `encoding/json` parity and did add a strict four-byte
  skip; the differential fuzz against the stdlib is what exposed the `<= 0x20`
  divergence in the first place. Don't "fix" `SkipWS` — the decode path is tuned.
- **Depth bound on the recursive walkers** (`unstable.MaxDepth` = 10000, matching
  `encoding/json`). `decodeValue`↔`decodeAnyObject`/`decodeAnyArray` and
  `stripper.handle` recurse once per nesting level and had **no bound**: measured,
  `StripDefaults` died at ~1M nesting and `DecodeAny` at ~4M with
  `fatal error: stack overflow` — which `recover` **cannot** catch, so one hostile
  document took the process down instead of returning an error. Both now carry a
  `depth` param: `decodeAny*` returns the new `ErrMaxDepth`, and `handle` ejects
  (its existing best-effort response to input it cannot interpret, keeping
  `StripDefaults`' no-error signature). Cost is one compare per `{`/`[` — citm
  `DecodeAny` measured flat (p=0.161, n=8). `Get`/`Set`/`SkipValue` are iterative
  or path-bounded and needed nothing.
- **Cycle-gated depth threading in the generator** (`computeDepthThreading`,
  `exprThreadsDepth`, `enterBody`/`depthParam`/`depthArgFor`/`depthGuard` in
  `main.go`) — the generated-code half of the bound above. A self-referential
  schema (`type Node struct { Kids []*Node }`) emits decoders that call each other
  in a loop, so decoding recurses per document level: measured, a 4M-level document
  died with `fatal error: stack overflow` (2M survived — Go grows the stack to 1 GB
  first, which is why the crash threshold is high and the naive test looks fine).
  The fix is **gated on a cycle actually existing**, which is what makes it free:
  `computeDepthThreading` builds the named-type reference graph (`namedRefs`, which
  unlike `markReferenced` keeps self-edges), marks types that reach themselves as
  cyclic, and threads `depth int` only through decoders for types that reach a
  cycle. Everything else is emitted byte-identically — verified by regenerating all
  30 bench cases + conformance under the old and new generator and diffing: **27 of
  30 identical**, the three that differ (`twitter_status`'s `twitterURL` nests
  `[]twitterURL`, `golang_source`'s `golangNode`, `synthea_fhir`'s
  `syntheaExtension`) being genuinely recursive — they were live crash vectors *in
  the benchmark corpus*. Two design points worth keeping: **(1)** the guard lives in
  the **struct** decoders only, because every cycle in a decodable schema must pass
  through a named struct — a named slice/map type is only decodable at the *root*,
  never as a field type (`field`'s `*ast.Ident` case rejects it), so an
  all-slice/map cycle cannot be generated. Struct frames pass `depth+1`, composite
  helpers (slice/array/map/lax-value) thread `depth` unchanged, so depth counts
  document levels rather than frames. **(2)** `markDepthFn` must be called
  **before** the body is generated, for the same reason `g.memo[key]` is: a
  recursive schema calls back into the function while its own body is still being
  built, and that call has to spell the same signature. Cost on the three recursive
  cases (interleaved A/B, n=8–16): twitter_status geomean −0.74%, synthea_fhir
  −1.24%, golang_source **+1.3%** (its `Lightning` p=0.270, `Destructive` +1.55%
  p=0.019) — golang_source is the deepest tree so it takes the most guard checks;
  under the 2% noise floor and the price of closing a fatal crash on exactly that
  shape. Locked by `TestRecursiveTypeDepthLimit` (self-reference) and
  `TestMutuallyRecursiveTypeDepthLimit` (Ring1↔Ring2 reached through a
  non-cyclic `RingRoot` — note a cycle with no member outside it gets *no*
  `UnmarshalJSON` at all, since `referenced` marks every member, so such a test
  needs a root above the cycle) plus
  `TestNonRecursiveTypesTakeNoDepthParam`, whose value is that it fails to
  *compile* if `Doc` ever grows a depth parameter.
- **`GetPaths` stack-backed active-index scratch.** `getPaths` keeps one shared
  `[]int` scratch holding the active path-index set for every recursion level
  (sized `len(paths)*(maxDepth+1)` so the depth-first walk's per-level sub-slices
  never reallocate — see the comment there). It was the function's *only*
  allocation. For the common small lookup the set fits in a `var stackbuf [32]int`
  used as the backing (`scratch = stackbuf[:0]`); only a larger set falls back to a
  single `make`. Safe because the backing never escapes — `out` only ever aliases
  `data`, never `scratch` (escape analysis confirms no heap move). Net on
  `BenchmarkGetPathsWithSkip` (3 paths, two sharing a nested parent): **−5.4% time,
  1→0 allocs, 80→0 B/op**. The rest of `GetPaths`' time is irreducible scanning
  (`ReadKey`/`SkipValue`/`SkipString`/`SkipWS`, ~75% of the profile). An early-exit
  when all paths are captured was considered and rejected: "first occurrence wins"
  means a not-yet-captured leaf must keep scanning later duplicate keys, so it
  wouldn't beat the current path. Covered by `BenchmarkGetPathsWithSkip` in
  `pkg/json/get_skipbench_test.go` (alongside `BenchmarkGetManyWithSkip`).
- **Zero-alloc `Set`/`SetMany`/`SetPaths`** (`pkg/json/set.go`). Four pieces.
  **(1)** All creation streams into `out`: `appendMember` writes a multi-key
  member's nesting directly (open-braces loop, rawVal, close-braces), and the
  non-object-intermediate case is signaled out of `setSpan` (`member` + `nested`
  flag) instead of returning a `nestValue` temporary — completing the
  member-append streaming that removed the single-key temp. Set create_nested
  **−46%**, overwrite_nonobject **−46%**, both 1–2 → 0 allocs; append_empty −15%
  from the same restructure. **(2)** `SetMany`'s found flags are a stack `[64]bool`
  (heap fallback beyond) and the key test leads with the string compare like
  `getMany`: **−4.7%**. **(3)** `SetPaths`' one heap set — `appendMembers`' per-key
  `sub`, whose append-from-nil growth defeats stack placement inside the
  `appendMembers`↔`appendMergedObject` mutual recursion — is built in a per-frame
  `[8]int` stack array instead (spilling to heap only past 8 deeper paths per
  key): **−1.6% time, 1→0 allocs**. Important negative result baked into that
  shape: porting `getPaths`' *shared* stack-backed scratch (threading
  active/recurse/create/matched through `setObject`) measured **+2–4%** — escape
  analysis already keeps `setObject`'s own `make([]bool, len(active))`, `recurse`,
  `create`, and `idx` locals on the stack (verified by alloc profile: `sub` was
  the *only* heap set), so the threading replaced free allocations with real
  bookkeeping and a 256-byte stackbuf memclr per call. Don't re-port it; check
  the alloc profile before assuming a `make` in this file is heap. **(4)** Both
  walkers used to scan every on-path container **twice**: `setSpan` pre-skipped
  each member's value (`skipValueOrEnd`) *before* the key test and then descended
  into that same value on a match, and `setObject`'s recurse branch *discarded*
  the recursion's returned end and used the pre-computed skip. Now the key test
  runs first (a non-leaf match descends without pre-skipping — the next level
  walks it member by member anyway) and the recurse branch takes its end from the
  recursion's return; only the leaf-replace/no-match/non-object branches skip.
  Measured **SetPaths −20.8%, overwrite_nonobject −9.8%**, and −80% on a
  deep-descend shape (descend keys first in a ~60 KB doc) no committed bench
  covered — the reason a 2× walk went unnoticed. A first draft that key-compared
  twice cost `replace` +4.4%; keep the single compare. Covered by
  `BenchmarkSet` (append/append_empty/replace/create_nested/overwrite_nonobject),
  `BenchmarkSetMany`, `BenchmarkSetPaths` — all zero allocs with a reused `out`.
- **SIMD escape scan (`IndexEscape` / `indexEscapeSSE2` / `indexEscapeNEON`).** `EscapeString`/
  `EscapeStringInto` used a SWAR clean-run scan (`swarHasLess|swarHasByte('"')|
  swarHasByte('\\')` per 8 bytes); the three SWAR tests were ~80% of a clean-string
  escape (the two `swarHasByte` alone ~70%). Replaced by an amd64 SSE2/AVX2 scanner
  `indexEscapeSSE2` (exported `IndexEscape`), structured exactly like
  `indexQuoteOrBackslashSSE2` (SSE2 first 32 bytes — no VZEROUPPER — then AVX2 once
  32 bytes are clean, then a 16-byte loop and scalar tail) with one extra per-block
  test for control bytes: `PMINUB(v, 0x1f) == v` (min(c,0x1f) equals c iff c ≤ 0x1f),
  `VPMINUB` on the AVX2 path. A `len < 16` short-circuit skips the three splat loads
  for sub-block buffers. **arm64 has a NEON twin** `indexEscapeNEON` (mirrors
  `indexQuoteOrBackslashNEON` — 16 bytes/iter, `VDUP`-built splats, `VMOV`+`RBIT`/
  `CLZ` position recovery, scalar tail) with the control test as `VUMIN(chunk, 0x1f)
  == chunk` (NEON's form of `PMINUB`); correctness verified (full pkg/unstable +
  pkg/json suites, incl. the `indexEscape` arm of `TestIndexFunctionsMatchScalar` and
  the all-bytes/fuzz `TestEscapeStringIntoReference`). Other arches keep the SWAR
  `indexEscapeScalar` (dispatch in `simd_amd64.go`/`simd_arm64.go`/`simd_scalar.go`,
  fallback + SWAR helpers in `simd_other.go`). Net (`BenchmarkEscapeStringInto`,
  amd64): **log_line_clean −80%, mostly_clean_one_quote −73%, url_clean −43%,
  sentence_clean −15%, short_clean −11%, control_bytes/prose/path −6…−9%**;
  `EscapeString` (Builder) mirrors it (log_line −52%, mostly_clean −40%).
- **`EscapeStringInto`'s per-run length gate** (the SWAR/vector chooser, shared by
  all arches). The pure always-vector form (`i += unstable.IndexEscape(s[i:])` per
  run) made escape-*dense* input regress: each short run between escapes pays the
  scanner's per-call setup (3 splats + a block + position recovery) to find an
  escape a few bytes in, where SWAR finds it in one word. The cost is intrinsic to
  the vector call and is **worse on arm64 than amd64** — NEON's setup + `VMOV`/`RBIT`
  recovery is heavier than SSE2's `PMOVMSKB`/`BSF`, so on **Apple M2** the dense
  cases ran **json_in_json +19%, path +34%, prose +43%** vs the SWAR baseline (amd64
  saw a milder +12% on json). The fix decides the scanner **once per run** by how
  much input is left: a run with `< minVectorRun` (48) bytes remaining — every short
  string and every short gap between escapes — is walked a word at a time with SWAR
  (exact offset via `TrailingZeros`, no vector call); only a longer run probes its
  first word with SWAR and hands the clean bulk to `IndexEscape`. Crucially the gate
  is **one length compare per run**, not per word, and leaves `indexEscape` inlinable
  — which is why it succeeds where earlier clawbacks failed (a per-word budget taxed
  pure-SWAR strings; a SWAR-prescan *inside* `indexEscape` broke its inlining, json
  +33%; an asm scalar peek added cost to every clean prefix, short_clean +56%). M2,
  vs the SWAR baseline (`main`): dense fixed — **path −4%, json −4.4%, prose +3.3%**
  (from +34/+19/+43%) — clean wins kept — **log_line_clean −58%, mostly_clean −51%**
  — small residual +3…+5% on short/medium clean (the lone gate compare on a 5–13 ns
  op); geomean −16%. `EscapeString` (Builder) is all wins (geomean −13.5%), its
  scratch alloc hiding the gate. The vector-vs-SWAR boundary is a heuristic — a
  *long* escape-dense run (rare; real escape-dense strings are short or
  frequently-escaped) can still take a vector call — but it only ever costs speed,
  never correctness, both paths gated by the all-bytes + fuzz
  `TestEscapeStringIntoReference`. asm vs scalar locked by the `indexEscape` arm of
  `TestIndexFunctionsMatchScalar` (control bytes in the inserted set).
- **`EscapeString`'s escaped-tail scratch is stack-backed** (`escape.go`): the
  Builder path used `make([]byte, 0, len(s)-pos)` per escaped string — a heap
  alloc whose cap *under*-estimates (escaping lengthens, `\u00XX` is 6
  bytes/byte), so dense tails regrew, 3–4 allocs/op. A `var buf [128]byte;
  EscapeStringInto(s[pos:], buf[:0])` never escapes (`EscapeStringInto` leaks its
  buffer only into the result, and `Builder.Write`'s param doesn't escape —
  verified with `-m`); only the escaped *tail* must fit since the clean prefix is
  written directly, and a longer tail regrows on the heap exactly as before.
  Measured (M2, n=8): **path_with_backslash −26.8%, json_in_json −20.1%,
  mostly_clean_one_quote −19.7%, prose −18.5%, control_bytes −15.1%**, B/op
  −46…−70%; clean cases exactly flat (they never reach the scratch).
- **Dynamic `any` path (`any.go`)**: the number case calls the private
  `scanFloat` directly (strconv fallback inlined at the site, mirroring
  `ReadFloat64OrNull` byte for byte) instead of going through the non-inlinable
  `ReadFloat64OrNull` — two frames per number plus re-checks of the truncation
  and `'n'` cases `decodeValue`'s switch already excludes; the same mechanism as
  the batched array readers. And `decodeAnyArray` gets the generated decoders'
  static first-append capacity hint (`cap(a)==0` → `make([]any, 0, 16)`, ~256
  bytes) instead of append growing 1→2→4→…; `[]` still returns the non-nil empty
  slice. The string case and `decodeAnyObject`'s key read carry the readKey
  inline trick by hand: both host functions are recursive and never inline, so
  the no-escape fast path (`indexCloseOrEscape` + one `string(rest[:k])` copy)
  is free body-size-wise and skips the non-inlined `ReadStringOrNull`/`ReadKey`
  call per clean string — the key path also drops the old alias-then-recopy
  `string([]byte(key))` at the map insert (one copy, not two). Escaped/truncated
  fall back to the old calls with identical error identities. Measured (with the
  word-compare literal rewrite below, interleaved n=10): **DecodeAny citm −4.15%,
  twitterescaped −3.45%**. The dynamic path remains ~2.8× the typed path
  (map+boxing, intrinsic — see the rejected key-interning entry); these only
  shave the removable overhead.
- **Literal matching is constant-string compares** (`ExpectNull`,
  `ReadBoolOrNull`, `SkipValue`'s true/false arms): `string(data[i:i+4]) ==
  "null"` compiles to one word load + compare against an immediate (no
  allocation, no memequal for constants ≤16 bytes) instead of 3–4 byte compares.
  Bounds handling and error positions are byte-identical (a partial "nul" at EOF
  fails the length test and returns i as before). `ExpectNull`'s inline cost
  drops 45 → **25** — headroom for the many decoders that inline it — and
  nothing that inlined before stopped inlining. Sub-noise alone; landed as a
  rider on the any-path work above.
- **Escaped-string decode, second pass** (on top of the four-part entry above;
  both changes in `string.go`, twitterescaped **−8.4%** and gsoc_2018 **−2.3%**
  interleaved). **(1)** `decodeEscaped`'s `\uXXXX` branch hand-encodes BMP runes
  (1/2/3-byte UTF-8 appends) instead of calling `utf8.AppendRune`, whose
  non-ASCII side is the non-inlinable `appendRuneNonASCII` — a call per rune
  that profiled 5.9% flat on `\uXXXX`-dense text; `AppendRune` remains only for
  supplementary-plane runes from valid surrogate pairs. Correctness subtlety: the
  pairing logic passed *unpaired* surrogates through raw and relied on
  AppendRune to rewrite them to U+FFFD, so the inline arm needs the explicit
  `utf16.IsSurrogate(r) → RuneError` normalization to stay byte-identical
  (locked by a differential test vs encoding/json over surrogate corners during
  development; the conformance suite covers the encode arms). **(2)** the
  close-quote cap hint's escaped-`\"` fallback no longer rescans via
  `SkipString` from the string start — profiling gsoc/twitterescaped showed that
  "rare" branch at ~2.5% flat (HTML/tweet text is full of `\"`), contradicting
  the original entry's rarity claim. It now *resumes* the `bytes.IndexByte`
  scan past each escaped quote, deciding by backslash-run parity (odd run =
  escaped; the run cannot extend past the first escape, which bounds the
  backward count), so each `\"` costs one vectorized sweep instead of a
  per-backslash rescan. Only truncated input still calls `SkipString`, keeping
  malformed-input sizing identical; the hint never changes acceptance. Also:
  `Unwrap` probes the string body through an `unsafeBytes` alias and copies
  only in the arms that return the body itself, so the base64 arm no longer
  pays a dead copy of the whole embedded document.
- **`Set` walkers, second pass** (`pkg/json/set.go`, on top of the four-part
  zero-alloc entry). **(1)** The three walkers read keys with the readKey
  inline trick (the get.go port CLAUDE.md left open pending amd64 evidence —
  measured here: `Set/overwrite_nonobject` **−8.8%**, `append_empty` **−6.5%**,
  rest flat, n=10 Zen 4). **(2)** `SetMany` counts found keys and, when all are
  found, splices the rest of the input verbatim (`append(out, in[prev:]...)`) —
  sound because all-found ⇔ no member remains to append at the close, and
  duplicates pass through either way (first occurrence wins). `setObject` has
  the same exit at the **root frame only** (`depth == 0 && nmatched ==
  len(active)`; a nested frame would still need its `}` offset for the parent,
  which costs the very scan the exit skips). The committed SetMany/SetPaths
  benches append members and so measure flat by design;
  `BenchmarkSetManyEarlyExit`/`BenchmarkSetPathsEarlyExit` (edit 2–3 early keys
  of a 45-member record) pin the shape the exit serves: the walk drops from
  O(doc) to O(prefix), 65/86 ns, zero allocs. One deliberate change rode along:
  on a *duplicate-key* document with every path matched, SetPaths used to
  re-edit later duplicates; the exit makes it first-occurrence-wins — which is
  what Set and SetMany already did, so this is a consistency fix, verified by a
  200k-random-document differential (unique-keyed docs byte-identical old vs
  new).
- **Root-slice growth extrapolates from decode progress**
  (`unstable.GrowSliceEst` in `grow.go`; emitted by `sliceDecoder` for **named
  slice roots only**, gated by a `root` flag whose memo marker keeps root and
  field decoders for the same element type distinct). A github_events-shaped
  document — a root array of large pointer-dense records — grew by pure flat-2×
  doubling: ~40% of its allocated bytes were dead doubled backings + memmove.
  At the grow point the decoder knows the array's `[` index (captured as
  `lightningArrStart`, the only extra emitted code — **no extra scanning**, the
  mechanism that distinguishes this from the rejected counting presizes), the
  cursor, and `len(data)`, so the new capacity is `len * (end−start)/(i−start)`,
  **padded by est/8+1** and clamped to [2×, 8×] of the current cap. The pad is
  load-bearing: the raw estimate landed at 29 of the true 30 on github_events
  and the mandatory 2× floor then doubled 29→58 — B/op measured **+42% worse**
  than flat-2× before the pad, −37% after; a near-exact estimate plus a growth
  floor overshoots, so make the estimate genuinely upper-ish. **Root-gating is
  the safety design**: the premise (array spans the rest of the document, so
  progress is a faithful density sample) is structurally true only at the root —
  for a nested slice the estimate always saturates its clamp, and ≥4× growth was
  measured and rejected for exactly those shapes in the flat-2× entry. Gated,
  the dual-generator diff shows **only github_events' decoders change** among
  all bench outputs; every nested slice keeps the tuned flat-2× byte-identically,
  so no guard measurements were even needed. Interleaved A/B (n=8):
  **github_events −27.6% time, −37.3% B/op, 1.31 GB/s** (was 0.95). Locked by
  `TestGrowSliceEst` (clamps, pad, degenerate inputs, no aliasing). A possible
  future extension — the estimate for *fields* that dominate the document tail
  (large-json's `features`) — needs a way to bound nested-slice waste first.
- **Trailing commas are rejected (first-iteration flag), matching
  encoding/json.** Every container loop — generated object/slice/fixed-array/map
  (`genStructBody`/`sliceDecoder`/`arrayDecoder`/`mapDecoder` in `main.go`), the
  batched readers, and `decodeAnyObject`/`decodeAnyArray` — used to check the
  closer at the loop *top*, so control flowed from `,` back to that check and
  `{"a":1,}` / `[1,]` silently decoded. The fix exploits an invariant of the
  existing shape: the loop top is reached only after the opener (first
  iteration) or after a comma, and a closer *after a member* returns from the
  post-value check — so a closer at the loop top on a non-first iteration ⇔
  trailing comma. `for first := true; ; first = false` + a branch inside the
  loop-top closer case (taken once per container) rejects it with
  `ErrInvalidJSON`, keeping the loop structure byte-identical otherwise (one
  register move per member). **Do not fix this by rotating the loop instead**
  (closer checked once pre-loop and post-value only, post-comma flowing straight
  into the key/element read): that shape was built first, is per-member-cheaper
  on paper, and measured **cloudflare +11% (p=0.000, n=8, isolated to the
  generated code — old runtime + new codegen reproduced it, new runtime + old
  codegen was +0.8%)** — the wide 45-field decoder is that sensitive to
  restructuring/layout; the flag form measures cloudflare +0.95% (within the
  runtime-only baseline) and citm flat. Locked by `TestTrailingCommaRejected`
  (conformance, incl. stdlib premise checks), `TestBatchTrailingComma`, and the
  `objErrs` trailing-comma arms in `any_test.go`. The same change closed two
  parity gaps: `slicePresize` resolves a `*ast.StarExpr` element (`[]*Foo`
  presizes by `[]Foo`'s rules — previously silently skipped), and
  `batchArrayFn` routes uint kinds to the new `DecodeUintArray` (a `[N]uint32`
  field used to keep the generated loop).

## The inline trick — let the generator write hot bodies inline

Go's inliner refuses `SkipWS`, `ReadKey`, and `indexCloseOrEscape` (each exceeds
the cost budget once it holds a SIMD/asm call), so calling them from generated
code pays full call overhead per token. Instead the **generator writes the hot
fast path inline at the call site** and falls back to a pkg/unstable call only for the
rare/hard case. See `g.skipWS` and `g.readKey` in `main.go`:

- **skipWS** emits the whitespace check inline:
  `if i < len(data) && data[i] <= ' ' { i++; if i < len(data) && data[i] <= ' ' { i = unstable.SkipWSRun(data, i+1) } }`.
  The common 0–1 whitespace bytes cost one or two compares and no call; only a run
  of ≥2 (pretty-print indentation) reaches the SWAR `SkipWSRun`.
- **readKey** emits the no-escape key read inline — a `unstable.IndexCloseOrEscape`
  scan plus a `unstable.UnsafeStr` alias — falling back to `unstable.ReadKey` only
  for an escaped key or an error. It relies on tiny inlinable exported wrappers
  `IndexCloseOrEscape`/`UnsafeStr` that themselves inline, so the generated code
  reaches the SIMD scanner with no wrapper call. Won ~6–7% on the cloudflare
  family, no regressions.

Why it beats making the pkg/unstable funcs inlinable: the cheap common case skips the
call entirely, and the shared scanner keeps its tuned dispatch.

The same trick applies by hand in **`StripDefaults`** (`pkg/json/strip_defaults.go`),
whose hot `handle` reads a key and a value string per object member — `SkipString`
was ~40% cumulative in the profile, all of it call frames since `SkipString`'s loop
keeps it un-inlinable. Each of the three `SkipString` call sites (key, value,
top-level string) now emits the no-escape close-quote scan inline —
`rest := in[i+1:]; k := unstable.IndexCloseOrEscape(rest); end := i+k+2` with a
`rest[k] == '"'` test — and falls back to `unstable.SkipString` only for an escaped
or truncated string. `IndexCloseOrEscape` inlines into the (never-inlined, recursive)
`handle` at all three sites; a helper wrapping the scan did **not** work — it costs
102 > the 80 budget once the SIMD scan inlines into it, so it stays a call frame just
like `SkipString` (the scan must inline *directly* into the big caller). Net:
**StripDefaults −10.5%, StripDefaultsCompact −12.0%** (geomean −11.3%), still zero
allocs. Covered by `BenchmarkStripDefaults`/`BenchmarkStripDefaultsCompact`.

**Scope it to once-per-loop reads.** `skipWS`/`readKey` fire once per object member,
so the inline block stays small. Inlining a *per-field* read (e.g. string values)
emits the block once per struct field, which bloats a wide struct's decoder
(`string_unicode` has 60 string fields): the function grows past the inliner's
per-function budget, `IndexCloseOrEscape` stops inlining, and i-cache pressure
turned a −8% cloudflare win into a +9% regression there. Don't inline per-field
reads.

**Make the dispatch wrapper itself inlinable (arm64).** `indexCloseOrEscape` is
the single hottest call in object decoding — on cloudflare it (and the NEON
scanner it guards) is ~40% of the work, and `pprof` showed the *wrapper* alone at
~11% flat because it wasn't inlined (cost 129 > 80: two calls — the NEON asm and
the `bytes.IndexByte` scalar fallback — plus a `useNEON && len(b)>=16` guard).
Collapsing it to a single unconditional `return indexQuoteOrBackslashNEON(b)`
drops the cost under budget so it inlines into every caller (`ReadStringOrNull`,
`skipString`, `decodeEscaped`, the generated decoders), removing that call frame.
Two facts make the collapse safe: Advanced SIMD (NEON) is **mandatory** in the
ARMv8-A baseline Go targets, so the `useNEON` (`cpu.ARM64.HasASIMD`) gate is dead
— always true — and can go; and the asm **already handles every length itself**
(its 16-byte loop falls through to a scalar byte tail for the final <16 bytes, no
out-of-bounds 16-byte load), so dropping the Go-side `len<16` → `bytes.IndexByte`
branch only changes who scans short buffers. That short-buffer path is rare in
decode anyway — the scanner is called on `string-body + rest-of-document`, so the
buffer is almost always ≥16 and the close quote is found in a NEON *block*, not
the tail (the tail fires once at end-of-document). Net, no regressions and broad
wins: **golang_source −5.6%, cloudflare −5.4%, citm_catalog −3.0%, twitter_status
−2.6%, string_unicode −2.0%, twitterescaped −2.0%**; gsoc_2018/synthea/large-json
flat. (amd64's `indexCloseOrEscape` was structured the same way and *also* not
inlined — cost 127 — and the analogous collapse to a single unconditional
`return indexQuoteOrBackslashSSE2(b)` lands the same win there: it drops to cost 61
and inlines into every caller. Safe for the same two reasons — SSE2 is the amd64
baseline so the SSE path needs no feature gate (the AVX2 switch is gated *inside*
the asm), and the asm handles every length itself (the 32- and 16-byte loops fall
through to a scalar tail), so dropping the Go-side `len(b) >= 16` →
`indexCloseOrEscapeScalar` branch only changes who scans the rare <16-byte buffer.
Measured on amd64: **cloudflare −4.8%, cloudflare-compact −4.1%, golang_source
−2.0%, string_unicode −1.7%, update_center −1.3%**; citm/twitter/synthea/gsoc flat,
no regressions.)

## Tried and rejected (don't re-attempt without a new idea)

- **skipWS inline trick in the pkg/json walkers — split verdict, measured 2026-08.**
  The two-compare + `SkipWSRun` block (the generated decoders' whitespace shape)
  was ported to every `SkipWSCompact` site in both `strip_defaults.go` and
  `get.go`, with new pretty-input benches (`get_pretty_bench_test.go` — no
  committed bench covered non-compact pkg/json input before). Interleaved A/B
  (final n=10, pinned Zen 4, idle machine): **strip_defaults.go kept** —
  StripDefaultsPretty −11.7% and StripDefaultsCompact −5.1% (both p=0.000;
  the compact win comes from the `for !compact` loop-entry test becoming a
  single predictable `if`), everything else flat — but **get.go reverted**:
  the first A/B (n=8) measured GetPretty +4.1%, ObjectEachPretty +5.7% (both
  p≤0.005) on exactly the pretty shapes it targeted, with the WithSkip micros
  flat; after the revert the Get family measures exactly flat again. The Get walkers are skip-dominated
  (`SkipValue` absorbs member values, whitespace and all, in the maskBlock bulk
  scan; only the short `": "`/newline-indent gaps between key tokens remain), so
  the 22 expanded SkipWSRun inline blocks bought nothing and cost i-cache/layout.
  `stripper.handle` visits every member of every object it keeps, so its 8 sites
  amortize. Don't re-port get.go without a new idea; set.go was never ported
  (its walkers are also skip-dominated — same economics as get.go).


- **Dropping the up-front `clear(out)` in the batched fixed-size-array readers**
  (`DecodeFloat64Array`/`DecodeIntArray`/`DecodeUintArray`) in favour of zeroing only
  the unfilled tail. The setup looks compelling: `out`'s length is dynamic to the
  callee, so `clear` cannot be sized at compile time and really does emit
  `CALL runtime.memclrNoHeapPointers` once per array — i.e. once per coordinate
  *point* for canada's `[][2]float64` and large-json's `[][3]float64` rings — and the
  clear is redundant whenever the JSON array fills every slot. **Rejected on
  sizing, before writing it.** `memclrNoHeapPointers` is 2.71% of canada and 9.04%
  of large-json, but `pprof -peek` attributes only **0.02s/2.58s = 0.78%** (canada)
  and **0.06s/3.54s = 1.69%** (large-json) to `DecodeFloat64Array`; the rest is
  `growslice`, `mallocgc` and `memclrNoHeapPointersChunked` — slice-backing zeroing,
  a different cost with a different fix (see the flat-2× growth entry, which is what
  that share was actually worth). So the achievable win is **below the 2% noise
  floor** while the change needs a labelled-break restructure of a hot loop to get a
  single exit for the tail clear, against the documented layout sensitivity. The
  lesson is the recurring one: an estimate built from "a CALL plus a 16-byte memclr
  is ~12–20 cycles × N points" predicted −4…−7%; the *attributed* profile says
  ~1%. Size a candidate with `pprof -peek`, not with cycle arithmetic.
- **Replacing `CountArrayScalars`' two vectorized calls with a fused inline scan.**
  A marine_ik CPU profile puts `CountArrayScalars` at **9.5% cum** (`bytes.Count`
  4.95% + `indexbytebody` 2.80%), and the reason looks damning: for the 3- and
  4-element `[]float64` coordinate arrays that dominate that document (`Pos`/`Rot`/
  `Scl`, ~20 bytes each) it makes **two SIMD library calls** — `bytes.IndexByte`
  for the `]`, then `bytes.Count` for the commas, two passes over the same 20
  bytes — where one inline pass could find both. Tried twice, both **net-negative**
  (interleaved A/B, n=8, amd64). **(1) A byte-at-a-time fused loop** (bounded at 64
  bytes, falling back to the vectorized path): float-array **+14.4%**,
  float-array-slow +9.8%, time-array +5.2%, mesh_pretty +3.3%, and marine_ik itself
  **+3.9% — worse**. Two flaws: the per-byte cost is ~6 compares (the 4-way
  whitespace test dominates), so a 20-byte array costs ~120 ops — *more* than the
  two calls it replaced; and `float-array` is a **~157-byte** array (293 ns/op), so
  it lands in the worst zone where the 64-byte prescan fails to find `]` and the
  fallback then **rescans from zero**. **(2) A SWAR fused loop** (8 bytes/iter,
  has-byte masks + `OnesCount64` for the commas, `TrailingZeros64` for the `]`,
  with the fallback *resuming* from where the prescan stopped rather than
  restarting — both flaws of (1) fixed): still net-negative — float-array **+6.1%**,
  time-array +2.5%, citm +2.3%, and marine_ik merely **flat (p=0.083)**. The lesson:
  Go's `bytes.IndexByte`/`bytes.Count` have cheap short-input fast paths, so the
  "call overhead" being blamed is a few ns, and *neither a scalar nor a SWAR loop
  can beat them even on a 20-byte buffer*. The 9.5% profile share is **not
  removable overhead** — it is the irreducible cost of reading those bytes, the
  same "a profile's cumulative % for a call chain is not the saveable overhead"
  trap as the unknown-field-skip inline above. Don't re-attempt by tuning the cap:
  the fast path measured *equal*, not slower, on exactly the short arrays it was
  built for, so there is no cap at which it starts winning.
  **The real lever on marine_ik is allocation, not counting** — see the next entry.
- **Formerly "un-landed": the slab for small `[]float64` backings is now live as
  `//lightning:arena`** (see that entry in the performance architecture above).
  The original sizing note stands: marine_ik's `DecodeFloat64Slice` was **95% of
  all allocated objects** (29 356 allocs/op — one tiny `make([]float64, 0, n)`
  per `Pos`/`Rot`/`Scl`), `mallocgc` ~20% of CPU. What made it a "design
  decision, not a local optimization" — the retention-semantics change (one
  surviving 3-float slice pins its whole chunk) and the arena threading through
  generated code — is resolved the same way `destructive` was: an opt-in
  directive, so the default semantics never change and non-arena schemas
  generate byte-identical code.

- **`switch len(key)` field dispatch, so no key comparison calls
  `runtime.memequal`** (`keyDispatch`/`chunkedKeyEq` in `main.go`). `cmd/compile`
  inlines a string-vs-constant comparison as word loads only while the constant is
  <= `maxRewriteLen` = 2*RegSize = **16 bytes**; past that it emits
  `CALL runtime.memequal`, which also spills and reloads the live registers around
  each failed compare and gives the decoder a stack frame. cloudflare's 45-field
  decoder made **10** such calls. The dispatch is now `switch len(key)`, and inside a
  bucket whose names exceed 16 bytes each name is compared in <=16-byte chunks
  (`key[0:16] == "EdgeTimeToFirstB" && key[16:] == "yteMs"`); buckets whose names all
  fit keep a nested `switch key`, so short keys dispatch exactly as before.
  **Gated on the struct actually having a name > 16 bytes**, which is what bounds the
  blast radius: regenerating all 30 bench cases + conformance under both generators
  leaves **16 of 31 byte-identical**, and the 15 that change are precisely those with
  a long key. Emitted-code check on cloudflare: `memequal` **10 → 0**, instruction
  count **1685 → 1645** (smaller), bounds-check panic sites **33 → 33** (the
  `switch len(key)` lets the slice bounds be proved, so the chunking adds none).
  Non-matching keys reach the skip via `goto lightningSkipKey` rather than a copy of
  the skip per bucket — a wide struct has ~20 buckets and duplicating it would add
  real code to the hot loop; the matched path costs nothing, falling out of the
  switch, which is why this is a goto and not a `handled` flag. (The skip sits in its
  own block so the `goto lightningKeyDone` over it does not cross a declaration,
  which Go forbids.) Interleaved A/B, pinned: **cloudflare −2.16% (p=0.000, n=16),
  cloudflare-compact −2.67% (p=0.002), cloudflare-nocopy −1.96% (p=0.000)**;
  string_unicode −1.6% (p=0.118); citm_catalog, twitter_status, synthea_fhir,
  update_center, marine_ik, instruments all flat. **Calibration note worth keeping:**
  an isolated microbenchmark of the two dispatch forms over cloudflare's real 43 keys
  measured −28% on the dispatch and predicted ~5% end-to-end; the delivered win is
  ~2–2.7%, because the isolated loop had perfect branch prediction and none of the
  surrounding memory traffic. Isolated dispatch benchmarks over-predict — halve them.
  A field with pipe-separated names of differing lengths (`shortAlt` /
  `aVeryMuchLongerAlternateName`) lands in two buckets and so has its decode code
  emitted twice; that is accepted as rarer and cheaper than a second dispatch.
  Locked by `TestLongKeyDispatch`, whose cases are chosen to break a careless
  implementation rather than to look representative: `sharedPrefix16xxA`/`B` are the
  same length and share their whole first chunk (comparing only chunk 1 swaps them —
  verified by sabotaging `chunkedKeyEq` and watching the test fail), a 33-byte name
  needs three chunks, and eight near-miss unknown keys must all be skipped.
- **A string arena / batched allocation for copied & escaped strings.** The alloc
  *count* looks addressable — on nocopy twitter_status `decodeStringEscaped` is
  **36% of allocations** (every escaped string — tweets are full of `\/`, `\"`,
  `\uXXXX` — must decode into a fresh buffer; nocopy can't alias an escaped body),
  and the copy-string variants (cloudflare default) allocate per value. The plan
  was a chunked arena (escaped strings appended into pooled chunks, old chunks kept
  valid so the aliases survive) to turn N small allocs into a few large ones. The
  *general, non-destructive* arena (threaded through every
  `Read*String`/`decodeEscaped` signature) is still unbuilt — but the
  highest-value slice of it shipped instead as the `//lightning:destructive`
  directive (above), with zero API churn. **Correction to an earlier note here:** a
  `GOGC=off` test once suggested the escaped-string allocs cost only ~1.4% — that
  was *misleading*. `GOGC=off` keeps the `mallocgc` + zeroing and removes only
  *collection*, so it measures the GC pause, not the allocate-plus-zero-plus-cache
  cost. Eliminating the allocations outright (`//lightning:destructive`) is **−41% time
  / −86% B/op on gsoc_2018** — allocation reduction *is* a large lever for
  escape-heavy input. Lesson: never size an allocation-reduction idea with
  `GOGC=off`; measure against actually not allocating. **Sized properly (2026-07),
  the general arena is dead**: the bench suite is nocopy-dominated and the pure
  copy workload (cloudflare) allocates only 144 B / 10 allocs/op — nothing to
  batch. The arena's real target is escaped-string scratch buffers on nocopy
  workloads (twitter_status 312 allocs/op, gsoc_2018 1 709), but a fresh-chunk
  arena allocates the same total *bytes* — GC rate and zeroing unchanged — saving
  only mallocgc *count*: ~25 ns × count ≈ **1.4% on twitter_status, 2.4% on
  gsoc** for signature churn through every `Read*` and generated decoder. The
  destructive directive already owns the eliminate-the-bytes win. Don't build it.

  **CORRECTION (2026-08-09) — that sizing model is wrong, and this verdict is
  withdrawn for the escaped-string case.** The `~25 ns × alloc count` estimate
  above predicts a win proportional to *allocation count*; measurement says the
  cost is span acquisition inside `makeslice`, which is proportional to *chunk
  count*. `pprof -peek` on gsoc_2018 attributes **10.4% of total** to
  `runtime.makeslice` under `decodeStringEscaped` (0.49s of 4.73s), not 2.4%. A
  measurement-only bump allocator replacing the `make` at `string.go:102`
  (interleaved A/B, n=10, pinned) measures **−8.78% time (p=0.001), allocs/op
  −56%, B/op −4.2%** — that is the *ceiling*, since a real arena also pays
  threading and bookkeeping, and it is below the −11…−13% an isolated estimate
  predicted (the usual over-prediction).
  The decisive experiment is the chunk-size sweep, and it is what proves the
  mechanism: at **4 KiB** chunks — the size `unstable.Arena` already uses — the
  same probe is **statistically flat (p=0.065)** and B/op is **+13.4% worse**
  (chunk-tail waste) while allocs/op still fall 37%. So a 37% reduction in
  allocation count buys *nothing*, which falsifies the count model outright, and
  **the existing `//lightning:arena` infrastructure cannot deliver this win** —
  it needs large chunks (≥64 KiB), which changes the pinning trade-off that made
  4 KiB the right choice there. Anyone picking this up should size it against the
  8.78% ceiling, not the 2.4% figure, and should not assume the Arena type is
  reusable as-is.
- **Inline clean-string fast path in the unknown-field `default:` skip** (the
  readKey inline trick applied to the generated `SkipValue` branch: emit
  `IndexCloseOrEscape` + close-quote test inline, fall back to `SkipValue` for a
  non-string, escaped, or truncated value). Motivated by an M2 cloudflare profile
  showing `SkipValue` 17% / `SkipString` 11% cumulative, with 26 of the doc's 42
  skipped values being clean strings. The fast path *worked* — all 16
  `IndexCloseOrEscape` sites still inlined, and the opt profile showed
  `SkipValue`/`SkipString` gone from the top-10 — yet interleaved A/B (n=8, M2)
  measured **cloudflare +0.48% (p=0.000)**, citm flat, string_unicode +0.19%. The
  two saved call frames were never on the critical path: each skip is
  latency-bound on the SIMD scan (identical either way), and the M-class OoO
  window hides call overhead behind it — the same "latency-bound, not call-bound"
  lesson as the NEON scanner micro-opts. Lesson: a profile's *cumulative* % for a
  call chain is not the saveable overhead; only the frames were removable and
  they cost ~nothing. The readKey/StripDefaults inline trick pays where the
  inline path *skips work* (no-escape key alias, WS fast path), not where it
  merely removes a call around identical work.
- **The readKey inline trick in pkg/json's `get.go`** (`getMany`/`walkPaths`/
  `objectEach`/`objectField` pay a non-inlined `unstable.ReadKey`, cost 212, per
  member; `SkipWS`/`SkipWSCompact`/`IndexCloseOrEscape` already inline there).
  Implemented the inline no-escape fast path at all four sites, confirmed the
  SIMD scan inlined at each — and measured `BenchmarkGetManyWithSkip`/
  `BenchmarkGetPathsWithSkip` **statistically flat on M2** (−0.14% geomean,
  p≥0.06): each member read is latency-bound on the SIMD scan and the M-class
  OoO window hides the frame, the same outcome as the unknown-field-skip inline
  above. The identical trick won 6–7% in generated code on amd64, so one
  interleaved A/B on an amd64 box could still justify it — don't land it on
  arm64 evidence. **Resolution: landed** (the get.go blocks are in-tree), and
  the amd64 evidence arrived via the set.go port of the same block, which
  measured `Set/overwrite_nonobject` −8.8% / `append_empty` −6.5% on Zen 4 —
  the trick pays on amd64 walkers where the M2 hid it. The Get-family micros
  themselves still measure flat on Zen 4 (their shapes are skip-dominated).
- **SWAR / uint64 key matching in the generated field switch.** The thought was to
  load a key's bytes as a `uint64` and compare against precomputed constants instead
  of `memcmp`. Already done — by the Go compiler: `key == "8bytechars"` against a
  constant compiles to a word load + compare (and the switch length-buckets first),
  so field dispatch is only ~3.5% even on cloudflare's 45-field struct (`memequal`
  2.0% + `memeqbody` 1.5%, the >8-byte names). No codegen change can beat what the
  compiler already emits.
- **AVX-512BW kmask tails for the string scanners** (`indexQuoteOrBackslashSSE2`
  / `indexEscapeSSE2`: a 64-byte `VPCMPEQB→k` loop replacing the AVX2 32-byte
  tail, GP-broadcast splats, entered where the AVX2 loop was). Implemented,
  byte-parity-verified under every dispatch variant at every offset — and
  **reverted on measurement**: string_unicode (the long-string case the tail
  serves) **+1.55% (p=0.040, n=16)**, gsoc flat. The skipBlocksAVX512 analogy
  does not transfer: that win came from collapsing a 7-instruction classify to
  2 per class, while this loop is already just load+compare+or+movemask and is
  **load-port-bound** — Zen 4 double-pumps zmm on a 256-bit datapath (64 B/iter
  costs the same load work as 2×32 B), so the wide form removes ~1 instruction
  per 64 bytes while its added code shifts alignment. Same lesson as the
  `VPCMPGTB` entry below: don't churn fuzz-verified SIMD asm for arithmetic
  that only removes ops hiding under a port bottleneck. The strengthened tests
  outlived the revert: `TestIndexFunctionsMatchScalar` now covers multi-block
  lengths to 320 and the 0x20 boundary byte, and `TestIndexVariantsFlip`
  differentially tests the pure-SSE2 and AVX2 arms under flag control (the
  live-dispatch test alone never exercises the narrower paths on a wide
  machine).
- **Eliminating the SWAR wide-load bounds checks** (SkipWSRun,
  ReadInt64/Uint64OrNull, scanFloat's fraction fold, the batch digit loops) —
  verified negative by direct codegen audit: the ~66 `panicBounds` sites
  visible in every objdump are *cold* branch targets of perfectly-predicted
  compares that issue in parallel with the loads they guard; restructuring to
  remove them (uint casts, explicit re-slicing) changes no hot-path schedule.
  The same "predicted branches are free; attributed profile beats cycle
  arithmetic" lesson as the memclr and unknown-field-skip entries — recorded so
  the panic sites stop looking like removable waste in disassembly sessions.
  Related watch-list from the same audit: `DecodeValue` (inline cost 77) and
  `skipNumber` (74) sit just under the 80 budget — re-check `-gcflags=-m` after
  any edit to either, as the SkipWSRun entry already mandates for itself.
- **Whitespace run-length memoization (`wsGuess`)** — the tempting fix for
  SkipWSRun's ~24% flat share on citm-shaped indentation: remember the last
  run's length g per call site and verify a repeat with two compares
  (`data[i+g-1] <= ' ' && data[i+g] > ' '`). **Unsound**: on a shorter actual
  run the byte at `i+g-1` can be a ≤0x20 byte *inside the next string token*
  (a space in `"hello world"`), so the check can skip over real token bytes
  and misdecode valid documents. The sound form must verify all g bytes —
  which is SkipWSRun's existing word loop minus only its final classify, so
  the ceiling is roughly half the naive estimate, against generator threading
  of per-site state and inline-budget risk. Not attempted; don't retry without
  an O(1) sound verification.
- **amd64 `indexStructuralAVX2` — `VPCMPGTB` to drop the trailing `NOTL`.** The
  shuffle-AND marker bytes are always `0x01`/`0x02` (positive), so
  `VPCMPGTB Yzero, Y6` yields `0xFF` where structural *directly* — no
  compare-to-zero + `NOTL` to invert a "non-structural" movemask. Correct (passes
  the differential fuzz) and one instruction shorter, but measured **flat**
  everywhere — micro (latency + throughput), skip-heavy, citm, large-json, canada
  all within noise. The loop is bound by port 5 (the two `VPSHUFB`); the `NOTL` is a
  GP op that already hid under that bottleneck, so removing it frees nothing.
  Reverted — not worth churning fuzz-verified SIMD asm for a non-win. (A real
  single-call latency win for the SSE2 string scanner is likewise elusive: load →
  `PCMPEQB` → `POR` → `PMOVMSKB` → `BSF` is the irreducible chain, and the
  destructive-compare copy `MOVOU X2,X3` is move-eliminated to ~0 latency on
  current uarchs, so there is nothing to shave.)

- **No-corpus determinations (quantified, don't build without a workload):**
  `[]float32`/`[]bool` batch readers — zero bench-corpus fields of those types
  (if a real workload appears, the `DecodeFloat64Slice` pattern ports
  mechanically and `arenaScalar` already includes `~float32`); a comma-count
  presize hint for flat-valued maps (`map[string]string`) — mechanism sound and
  unclaimed, but its entire measurable corpus is ~106 citm members, ≤0.3%
  there; and the update_center struct-valued-map presize re-check *with* the
  AVX-512 skip machinery — still a wash (~5% saved vs ~5–8% extent-scan cost).
- **Routing Clinger's negative-exponent case through Eisel-Lemire** (replace the
  `f /= pow10exact[-exp]` division with EL's 128-bit multiply, which is also
  correctly-rounded so it's a legal swap): measured *worse* everywhere —
  canada_geometry +12%, `float-array` +24%, large-json +3%. A single FP division
  is cheaper than EL's table load + `bits.Mul64` + normalization. The tier order
  (Clinger first, EL only for mantissa ≥2^53 or |exp|>22) is correct as-is.
- **Presizing slices whose element nests a slice/array/struct** (forcing
  `CountArrayElements` on the multi-dim coordinate slices and struct-with-nested
  slices that `slicePresize` deliberately skips): catastrophic — citm_catalog
  +155%, canada +61%, large-json +48%. Counting a nested element costs the same
  as decoding it (`SkipValue` recurses through the whole subtree), and presizing
  at every nesting level re-counts the sub-structure O(depth) times. `growslice`
  (12–20% on these benchmarks) is the lesser evil; the presize-skip rules stand.
  A follow-up tried presizing *just* the leaf coordinate rings (`[][N]scalar`)
  with a cheap bracket-only counter (`CountNestedScalarElements`, an
  `indexStructural` depth walk that needs no per-element `SkipValue`): still
  net-negative — canada +7%, large-json +14%. The count plus the `memclr` of the
  presized backing outweighs the rings' append growth. Only the *date* half of
  that idea paid off (the cheap comma count for `[]time.Time`, above), because
  there the array was already presized — the win was a cheaper *counter*, not a
  newly-presized slice. **Resolution (now live):** the lesson — count + `memclr`
  beats nothing for these rings — is acted on by *not presizing them at all*. A
  `[][N]scalar` ring (large-json `[][3]float64`, canada `[][2]float64`) was being
  presized by default via `CountArrayElements`; the `t.Len == nil` guard in
  `slicePresize` now skips it (see that entry above), −8.8% on large-json. Not
  presizing beats every counter because it pays neither the count scan nor the
  `memclr`.
- **SWAR-folding the RFC 3339 fractional-seconds (nanosecond) loop** (`is4Digits`/
  `parse4Digits` on the `.190533` digits): statistically tied on time-array — the
  digit accumulation isn't the bottleneck there (validation + `time.Unix`
  construction dominate), and assembling the uint32 from string bytes doesn't fold
  to one load. Not worth the extra code.
- **SWAR fraction fold with an 8-byte chunk** (any arrangement): the live decoder
  folds fractions four bytes at a time; adding an 8-byte path never paid off.
  8-digit-only regressed common `float-array` and was flat on `large-json` (4–7
  digit fractions fail the 8-byte gate). 4-then-8-runs-then-trailing-4 beat scalar
  but a direct A/B vs the flat 4-byte loop was *tied* on real data and cost 2.56%
  on `float-array` (it only won `float-array-slow`'s synthetic 16-digit mantissas).
  Scalar-scanning the first four then switching to 8-byte runs was worse still
  (canada +2%, `float-array` +16%). See the "SWAR fractional digits" entry above.
  Don't reintroduce an 8-byte chunk without a new idea.
- **One-byte `data[i+3]` digit guard** before the 4-byte fold (skip the load+
  `is4Digits` for <4-digit fractions): didn't move `float-array` (its regression
  wasn't the failed test — it's alignment-class) and cost ~1.5% on canada by adding
  a load to its always-passing 14-digit path. Rejected.
- **Variable-length SWAR fraction fold** (count leading digits, fold k at once): a
  single 8-byte loop where `leadingDigits8` (nibble test + carry-safe haszero +
  `TrailingZeros`) returns the digit count and a shift-padded `parse8Digits` folds
  the final partial chunk (no divide). Correct (passes the strconv fuzz) and
  elegant — folds the whole fraction in SWAR — but measured worse everywhere
  (canada/large-json flat, `float-array` +14%, slow −11%): the count's
  mask+`TrailingZeros` is dearer than `is4Digits`, and the variable shift + `pow10`
  table load beats the fixed-width `parse4Digits` whose multipliers are
  compile-time immediates. Fixed widths win *because* they're fixed.
- **Vectorizing `SkipWS` as a standalone function**: regresses — the SWAR/SIMD
  setup isn't amortized on the common 0–1 byte runs. It pays off *only* via the
  inline trick above (inline fast path for short runs; `SkipWSRun` SWAR for ≥2-byte
  runs). Done — not rejected — once moved to the call site.
- **Widening `SkipWSRun` past 8 bytes/iter** (a 32-byte unrolled SWAR after the
  first 8-byte chunk, for deeply-indented pretty JSON whose runs are 16–28 bytes):
  regressed everywhere it was tried (citm +5%, twitter +3%, synthea +4.5%). The
  wide loop reads 32 bytes (four loads) even when the run ends mid-chunk, so it
  does *more* loads than the 8-byte loop that stops as soon as a non-space chunk
  appears. The plain 8-byte SWAR is already near-optimal for these run lengths. (Its
  per-word *constant* was not near-optimal, though — see the landed all-spaces
  equality fast path above, which is orthogonal to width.)
- **SSE2 `skipNonWS` continuation for long whitespace runs** (`SkipWSRun` keeps
  its 8-byte SWAR word for ≤8-byte runs but, when the first 8 bytes are all
  whitespace, calls an SSE2 `PMINUB`+`PCMPEQB`+`PMOVMSKB` find-first-non-space for
  the rest — `min`/`eq` against a 0x20 splat, 16 bytes/iter): a tight-loop
  microbenchmark looked great (run 9 −25%, run 25 −52% vs SWAR), but in the real
  decode it was **flat on citm and regressed mesh_pretty +6.5%, twitter +3%,
  cloudflare +1.5%**. `SkipWSRun` is called once per inter-token gap; the SSE
  routine's per-call setup (splat load, the SSE→GP `PMOVMSKB`, the asm call that
  can't inline) is not amortized across a single run the way the microbench's tight
  loop hid. citm's 21–29-byte runs are long enough that SSE2 *should* win, yet the
  per-call overhead cancels it; the 9–13-byte runs of the other pretty cases tip
  net-negative. Confirms the two entries above: whitespace skipping resists SIMD —
  the 8-byte SWAR's cheapness per call beats any vector setup. Don't retry without
  eliminating the per-call cost (e.g. inlining the SSE into the generated `skipWS`,
  which would bloat every call site — see the "inline trick" scoping note).
- **Pure-SSE2 `indexStructural`** (dropping AVX2): ~2× slower on the skip path
  (`skip-heavy`); the throughput loss dwarfs the VZEROUPPER saving.
- **arm64 `indexQuoteOrBackslashNEON` hot-loop micro-opts** — five plausible
  rewrites of the per-block match-test, each killed by the same lesson: the
  decoder consumes each scanner result *immediately* (to slice the string), so a
  single call is **latency-bound**, not throughput-bound. A tight independent-call
  microbenchmark measures throughput and *lies* about these — its `huge −13%`
  became `+6%` real. The block already does two *parallel* `VMOV`s of the D lanes
  + two `CBNZ`s; that two-move form is **latency-optimal** (both lanes land at
  once, recovery reads them directly), and every attempt to "reduce" it adds an op
  to the chunk→branch critical path. **(1) Fold the 16-lane mask to one 64-bit
  detect word** (`VEXT $8`+`VORR`, one `VMOV`/`CBNZ`, exact lane only in a cold
  `found` path): throughput microbench loved it (huge −13%) but **regressed every
  real decode +2.4%…+6.3%** — longer per-block path and the cold path re-does both
  `VMOV`s. **(2) True NEON movemask** (`VUSHR $4, .H8` + `VUZP1` narrow the mask to
  4 bits/byte so one `VMOV` carries all 16 lanes and `RBIT`/`CLZ`/`>>2` recover the
  byte — the SHRN trick the Go assembler can't spell directly): genuinely halves
  cross-domain traffic with *no* cold re-move, the cleanest version of the idea —
  and still **regressed +2.7%…+9.2%**, because the two extra vector ops sit on the
  latency path and the loop was never throughput-bound. (Also: it only works where
  the mask byte is `0xFF`; `indexStructuralNEON`'s marker is the `0x01` table-AND,
  which `>>4` annihilates — the differential fuzz catches it instantly.) **(3)
  Defer the high-lane `VMOV` past the first `CBNZ`** (so a low-half match skips
  it): a *real, data-dependent* split — golang_source −5.9% (short tokens match in
  the low half) but cloudflare **+3.9%** (16–31-byte strings match in the high half
  of block 1); a cloudflare regression is a blocker and there's no static way to
  know which half a string ends in. **(4) Overlapping final NEON block** instead of
  the scalar tail: flat — the tail is only hit once at end-of-document. **(5) One
  `VLD1` of a 32-byte RODATA splat pair** instead of two `MOVD`-imm+`VDUP`: a
  tradeoff (long strings −1.5…−1.9%, cloudflare +3.7% load-use latency), the
  short-vs-long tension the asm comment already settles for `VDUP`. The hot loop is
  genuinely well-tuned; the real arm64 string-handling win was making the dispatch
  wrapper inline (above), not touching the block body. Don't re-attempt mask-
  reduction tricks without a way to *shorten single-call latency*, not block
  throughput.
- **arm64 register-ABI (`<ABIInternal>`) for the NEON scanners** — the asm is
  ABI0, so every (very hot) call spills base/len/cap to the stack and reloads the
  result; declaring the functions `<ABIInternal>` would pass the slice in R0/R1/R2
  and return in R0, removing that marshaling. The Go toolchain **forbids the
  `<ABIInternal>` selector outside `package runtime`** ("ABI selector only
  permitted when compiling runtime"), so this is simply unavailable to a
  third-party package. No workaround that doesn't fork the toolchain.
- **arm64 32-byte-unrolled `indexQuoteOrBackslashNEON` for long strings.** The
  bench-md flags string_unicode (long unicode text fields) as arm64's #2 lag vs
  amd64 (~0.68), and the scanner is ~51% of that decode; amd64 wins partly by
  switching to a 32-byte AVX2 tail. The arm64 analogue: a `loop32` (one 2-register
  `VLD1` of 32 bytes) reached only when ≥32 bytes remain — so short keys/values keep
  the latency-tuned 16-byte `loop16` — using the `VUSHR`+`VUZP1` movemask (one
  cross-domain `VMOV` per 16-byte half instead of two) to cut the throughput-bound
  long-string scan's `VMOV` traffic in half. **Regressed both**: string_unicode +13%
  *and* cloudflare +13% — and cloudflare never executes `loop32`, so its regression
  is pure **code-alignment** shift from inserting the block (the float-array lesson).
  string_unicode regressing too means the scan isn't `VMOV`-bound — `loop16`'s two D
  extracts already issue in parallel, so the movemask's extra `VUSHR`+`VUZP1` on the
  path only adds latency. Confirms the existing "arm64 string scanner is latency-
  bound and resists mask-reduction" finding; the string_unicode lag is **intrinsic**
  (NEON 16-byte vs AVX2 32-byte compare width), not closable in the block body.
- **Key interning / map presizing in the dynamic `any` decoder**
  (`decodeAnyObject`, the `DecodeValue`/`json.DecodeAny` path): twitterescaped's
  `DecodeAny` is bound by building `map[string]any` — `mapassign` plus bucket
  allocation is ~40% of the decode, and the per-key copy `m[string([]byte(key))]`
  is the single biggest allocation (54% of allocs; twitter keys like `id`,
  `created_at`, `text` repeat thousands of times). **Interning** the keys — a
  document-wide `map[string]string` threaded through `decodeValue`/`decodeAnyObject`
  (nil for the per-field path, a real map for `json.DecodeAny`), copying each
  distinct key out of the input once — cut allocs −42% and B/op −9% but left
  wall-time flat (−1.6%, noise): it trades ~14k key *allocations* for ~14k
  intern-map *lookups* (`mapaccess2`, a second hash of the same bytes), a wash.
  **Presizing** the result map didn't help either — twitter objects are bimodal
  (many 2–5-field nested objects, a few 30–40-field ones) so a fixed `make(map, N)`
  hint mis-sizes both and `make(map,8)` moved nothing; an *exact* member count
  needs a structural pre-scan that re-walks nested content depth-times, net-negative
  for the same reason `slicePresize` skips complex elements. The `any` path is
  ~2.8× the typed path (twitterescaped 1.9 ms vs 0.67 ms) purely from `map[string]any`
  + interface boxing, intrinsic to the dynamic representation. Reverted both; don't
  re-attempt without a way to populate the result map without the per-key hash.
- **Two-stage structural feed** (simdjson-style: a stage-1 SIMD pass builds a
  structural-position *index*, then stage-2 navigates it by position instead of
  scanning inline). Fully prototyped (the `VPSHUFB`/nibble-table classification
  extends to `{}[]":,.\` within 4 bits for free; stage-1 indexes at 7–15 GB/s; a
  hand-written typed stage-2 produced byte-identical output). Measured **net-negative
  on the common case and a win only on heavily pretty-printed input** — it is purely
  a *whitespace play*, winning in proportion to the whitespace fraction (stage-2
  jumps token-to-token, never skipping whitespace; stage-1 absorbs it in the bulk
  SIMD scan) and losing a flat ~30% on compact (stage-1 is pure overhead with
  nothing to reclaim). Fair A/B (both allocate output; two-stage reuses only its
  index scratch): synthetic flat-int records (59% ws) pretty **−28%** / compact
  **+29%**; cloudflare-like string records (shallow, 34% ws) pretty **−11%** /
  compact **+30%**. A **string-aware** index (the simdjson inside-string mask —
  `find_escaped` + a quote-mask prefix-XOR, verified bit-exact, ~13 GB/s — so stage-2
  can alias clean strings straight from the boundaries: gsoc 95% of strings clean)
  did *not* rescue the compact case: the mask *adds* stage-1 cost, and single-pass
  `nocopy` already aliases a short string after one `indexCloseOrEscape` that finds
  its close-quote in the first 16–32 bytes — little to amortize. The deeper reason
  the economics differ from simdjson: simdjson's stage-2 is a cheap *tape copy* that
  amortizes stage-1, whereas lightning's "stage-2" is the expensive *typed* parse, so
  the index is mostly redundant work plus a second pass over the bytes. Not worth a
  wholesale rewrite (it would regress every compact wire-format decode); only ever an
  opt-in mode for sources known to be both large and deeply indented.
- **`append(value)` instead of grow-zero-then-assign for scalar slice elements.**
  `sliceDecoder` decodes every element into the freshly-grown last slot
  (`var zero T; *out = append(*out, zero); …(*out)[len(*out)-1] = v`). For a
  *composite* element (struct/slice/map/pointer) this is load-bearing — decoding
  through `&(*out)[last]` keeps the element from living in an escaping local, which
  would heap-allocate per element. For a *by-value* leaf (number/string/time/raw,
  whose reader returns the value) it looks wasteful — a dead zero-store plus an
  indexed assign with a bounds check, where `*out = append(*out, v)` is one write.
  Split the codegen so by-value leaves append directly: **statistically flat
  everywhere** (marine_ik, numbers, canada, mesh, float-array, golang_source, citm —
  all within noise). The Go compiler already dead-stores the zero and elides the
  `(*out)[len(*out)-1]` bounds check, so the two forms compile to the same code.
  Reverted — it only added an `isByValueLeaf` branch to the generator for no gain.

## Conventions

- **The edit/transform API is deliberately two-tier.** `Set`/`SetMany`/`SetPaths`/
  `StripDefaults` return only a `[]byte` and are best effort — bracket balancers,
  not parsers, that pass uninterpretable input through rather than failing. That is
  what keeps them zero-alloc on the hot path, and it is not a defect to be fixed by
  adding error returns to them. Untrusted input is served by the `…Checked`
  counterparts in `pkg/json/checked.go`, which wrap the unchanged fast functions
  with `Valid` on the arguments and on the result (plus `ErrValueCount` for a short
  `rawVal`). Keep new edit operations to that shape: fast and silent, with a checked
  wrapper — never a validity check inside the hot walker.
- Bench `data.go` files use a single top-level `Benchmark` with **anonymous**
  nested structs, so only `Benchmark` gets a generated method and
  `type benchmarkStd Benchmark` gives a clean reflection-only baseline for the
  stdlib/sonic benchmarks.
- Every case also gets a `BenchmarkLightningDecodeAny`: the same document decoded
  through `json.DecodeAny` (compact mode) into the generic `any` value instead of
  the typed `Benchmark`, so the rendered table contrasts schema-less decoding with
  the generated unmarshaler. The harness minifies each input with `json.Compact`
  first (many corpus inputs are pretty-printed) so the compact path is valid.
- And a `BenchmarkLightningDestructive`: the `//lightning:destructive` variant.
  `run_bench.sh` duplicates `data.go` into a gitignored `data_destructive.go`,
  renaming **every** top-level type `…Destructive` (an `awk` extractor handles both
  the `type (...)` block and single-`type` forms; helpers are renamed too because the
  generator parses the variant file alone and must resolve every referenced type) and
  prepending `//lightning:destructive` to the root, then generates a second decoder.
  Any `:compact`/`:nocopy` directive in the source rides along in the copy. The
  benchmark restores a pristine copy of the input into a reused buffer each iteration
  (the destructive decode mutates it) and decodes through
  `(*BenchmarkDestructive).UnmarshalJSON` — so the gap vs `BenchmarkLightning`
  *understates* the real win (a true owner wouldn't pay the restore-copy). The
  restore also perturbs cache state, which cuts the other way on byte-bound
  cases: skip-heavy's apparent Destructive "+38%" in the committed amd64 table
  is this rotating-buffer cache effect on a >10 GB/s case, not a regression —
  read Destructive rows on such cases accordingly. Cases with
  no `nocopy` string fields generate an identical decoder (destructive is a no-op
  there) and read ~flat. Implemented as a per-case source duplicate, **not** a
  generator twin (a `-inplace-twin` flag emitting a second `UnmarshalJSONInPlace`
  method was prototyped and dropped in favour of the simpler duplicate).
- End-of-session: run the full suite (`go test ./...`), keep gofmt clean
  (the `daysFromCivil` comment-alignment flag is pre-existing — ignore it),
  and regenerate the committed benchmark markdown via `make bench-md`. That runs
  **two** suites: `pkg_bench.sh` (the main-module microbenchmarks — every
  `Benchmark*` in `pkg/json`/`pkg/unstable`, rendered to `bench/pkg_results_<arch>.md`
  by `bench/pkg_results_md.py`) and `bench/run_bench.sh` (the competitor-comparison
  suite, rendered to `bench/results_<arch>.md` by `bench/results_md.py`). Both write
  a raw `*results.txt` (gitignored) and commit the per-arch `.md`. `pkg_bench.sh`
  takes an optional benchmark-name filter as `$1` and honours `BENCHTIME`/`BENCHCOUNT`.

## Apple M2 re-validation of the Zen-4-measured stack (2026-08-08)

Every perf commit landed since the NEON whole-loop port was interleaved-A/B'd
(n=10, benchstat, parent-vs-commit builds from worktrees) on Apple M2, since all
of them had been measured only on the pinned Zen 4 box. Everything transfers in
direction; four entries diverge enough in magnitude to matter when reading the
per-entry numbers above (M2 vs Zen 4):

- **Transfers cleanly**: escape second pass (twitterescaped −7.1% vs −8.4%, gsoc
  −4.6% vs −2.3%; the UnescapeString micro isolates it — unicode_escaped_dense
  −9.3%/−10.8%, every other case exactly flat), any-path inlines (DecodeAny citm
  −3.5% vs −4.15%, twitterescaped −1.9% vs −3.45%), len(key) dispatch
  (cloudflare −2.05% vs −2.16%), SetMany/SetPaths all-found early exits (−92% on
  the early-exit shapes, 1005→82 ns / 1086→89 ns), and the arena directive,
  which is *better* on M2: marine_ik **−6.6%** time (Zen 4 −2.8%), mesh −3.0%,
  allocs/op identical to Zen 4 (−94.9%/−99.2%).
- **GrowSliceEst**: github_events B/op −37.3% — *exactly* the Zen 4 number, as
  allocation mechanics must be — but time −10.8% vs Zen 4's −27.6% (M2's
  faster allocator/memmove shrinks the doubling share it removes).
- **Smaller on M2**: strip_defaults skipWS port — StripDefaultsPretty −4.0% (vs
  −11.7%), StripDefaults −1.0%, but StripDefaultsCompact **+2.7%** (vs −5.1%);
  net ≈ wash on M2, clearly positive on Zen 4. All-spaces SkipWSRun fast path —
  citm −0.9% (vs −4.1%): the M-class OoO core hides the classify ALU ops the
  fast path skips, even though SkipWSRun is still ~22% of citm's decode tree on
  M2. Same lesson as the unknown-field-skip rejection: M2 hides short
  ALU/call-frame savings that Zen 4 exposes.
- **Set-walker readKey inline**: the early exits dominate the commit, but the
  inline-key-read half is mildly *negative* on M2 — overwrite_nonobject +5.4%
  (~2 ns), append +2.5%, replace/SetMany/SetPaths +1.6–1.9%, create_nested
  −3.7% — where Zen 4 measured −8.8%/−6.5%. Third confirmation (after the
  get.go flats and the unknown-field-skip revert) that this trick pays on Zen 4
  and is ~neutral-to-slightly-negative on M2. Keep: the Zen 4 win is real and
  the M2 cost is ~2 ns on ns-scale micros.

M2 profiling note: macOS CPU profiles are dominated by `runtime.kevent` +
`runtime.madvise` samples from background threads (81%+18% flat on a citm run —
they swamp the denominator). Read decode shares with `pprof
-ignore='kevent|madvise|pthread'` or as fractions of the decode tree, not of
total samples. The filtered M2 profiles match the documented cost structure:
citm SkipWSRun ~22% of decode, cloudflare scanner ~30% + unknown-field skips
~33%, gsoc scanner+decodeEscaped — all already-attacked or documented-intrinsic;
no new addressable hot spot surfaced.

## Session 2026-08 fixes worth knowing (correctness/API)

- **`valueDecoder` memo keys now carry `g.prefix + g.cmark()`** like every other
  emitter (and names go through `g.decFn`/`g.csuf`). Before, roots with
  different directives sharing a lax field type shared one decoder — a plain
  root could inherit a destructive sibling's in-place unescape (caller-buffer
  mutation) or a compact sibling's whitespace rejection, and two files with the
  same lax field type generated colliding names. Locked by
  `TestLaxDecoderIsolation` (conformance).
- **StripDefaults keep-key + emptied container value**: the rewind now backs up
  to just past the separator comma (`postComma`), not to `localStartWrite` —
  rewinding past the comma emitted `{"a":1"b":{...}}`, invalid JSON. Locked by
  `TestStripDefaultsKeepKeyContainer`.
- **Generated unwrap closures guard an all-whitespace body** (`i >= len(data)`
  after the SkipWS): a pointer field's null probe reads `data[i]` unguarded and
  panicked on `{"p":" "}`. Locked by `TestUnwrapWhitespaceBody`.
- **Pointer fields reuse a non-nil pointee** (`if dest == nil { dest = new(T) }`),
  matching encoding/json's documented pointer semantics and making reuse
  allocation-free on pointer-dense schemas; null still nils. Locked by
  `TestPointerFieldReuse` (with a stdlib premise check).
- **`[]byte` follows encoding/json**: base64 string or numeric array both decode
  (`unstable.DecodeByteSlice` + `...Arena`; `batchSliceFn` routes byte/uint8
  there); `[N]byte` stays numeric-only like the stdlib. Locked by
  `TestByteSliceStdlibParity`, `TestDecodeByteSlice`.
- **Directives are validated**: unknown `//lightning:*` names are errors; a known
  directive on a non-root/referenced type or bare `nocopy` on a struct root
  warns on stderr (it silently did nothing before).
- **pkg/json re-exports all eleven sentinels** (added ErrBadNumber/ErrBadEscape/
  ErrBadUnicode/ErrBadTime); `TestSentinelsMatchable` locks the contract. Get's
  non-object descent is documented as ErrExpectObject (was misdocumented as
  ErrKeyNotFound).
- **`time.Parse` does NOT retain its input in errors on modern Go** (the stdlib
  copies into ParseError) — a plausible unsafe-alias finding refuted by test;
  `TestReadTimeErrorRetainsNoAlias` guards the property against toolchain drift.
- **`SwarNeedsEscape`** (pkg/unstable) is now the one spelling of the JSON
  escape-byte predicate, shared by `indexEscapeScalar` and `EscapeStringInto`'s
  probe — the dedup was verified byte-identical in `EscapeStringInto`'s asm.
- The generator's dead `need*` import flags were deleted (`assemble` scans the
  generated text); lax `[N]scalar` fields route through the batched array
  readers via a thin generated wrapper (`TestLaxFixedArrays`).

## Session 2026-08-09 audit: 23 findings fixed, and the traps found along the way

A full correctness/API/performance/DRY audit. Every finding below was reproduced
before being fixed and re-reproduced by an independent verifier afterwards. The
entries worth carrying forward are the *reasoning*, not the diffs.

### Two memory-safety defects, both reachable from the public API

- **`skipObject`/`skipArray` recursed with no bound** (`skip.go`). This entry's own
  MaxDepth note used to say "Get/Set/SkipValue are iterative or path-bounded and
  needed nothing" — true of `skipContainerFast` and of `Get`'s loop, **false of the
  scalar fallback underneath them**. Measured: `{"k":1,"a":[0,[[[…10M…]]]],"zz":7}`
  (20 MB) through `json.Get` died with `fatal error: stack overflow`, which
  `recover` cannot catch; 4M nesting survived. Reaching it on an AVX2 host needs one
  detail — a *scalar first element* makes `SkipValue`'s fast-path probe decline, and
  every nested container from there recurses unconditionally. Fixed by threading
  `depth` through unexported `skipObjectDepth`/`skipArrayDepth` (the two-arg entries
  stay inlinable wrappers, so `SkipValue` is unchanged). Locked by
  `TestSkipDepthBound`.
- **Presize hints were bounded by document bytes, not by possible element count**
  (`count.go`). `CountArrayObjects` counts `{` before the first `]` *including braces
  inside string values* — documented as harmless because "a miscount only mis-sizes,
  never misdecodes", which is true of correctness and silent about memory, since the
  count becomes a `make` capacity. Measured 65× amplification, linear in input
  (2 MB → 130 MB, `len(items)=1 cap=2000001`).
  **The fix, and the sizing lesson.** All three counters now clamp to what the byte
  span can structurally hold (`(rb+1)/3` for objects, `(rb+1)/2` for scalars,
  `span/2` for the extrapolation). The right target was not "bound the allocation"
  but "bound the hint to what a *legal* document of this span could have produced" —
  verified directly: post-fix, a 2 MB brace-bomb allocates 98 MB and a 2 MB **honest**
  densest-legal array of the same element type allocates 96 MB, so crafted input buys
  **1.02×** over simply sending real data. Anything tighter would under-size honest
  documents. The residual `sizeof(element)` factor is the intrinsic JSON→Go expansion
  ratio, not an attack — an absolute cap was considered and rejected on that basis.
  Note the clamp is *deliberately* a no-op on unrepresentative leading elements (64
  leading `""`/`{}` then one huge element extrapolate identically with or without
  it); that shape is already inside the honest-document ceiling.

### The methodology trap that nearly lost two findings

A generated `UnmarshalJSON` makes its type a `json.Unmarshaler`, so
`encoding/json.Unmarshal(doc, &target)` **delegates to lightning's own decoder** and
compares it against itself. Measured against that broken baseline, the RawMessage-null
and slice-element-reuse findings both "matched the stdlib exactly" and were nearly
dismissed. Every stdlib differential must go through a defined twin with no methods —
`type FooStd Foo`, which is exactly why `bench/` uses `type benchmarkStd Benchmark`.
Against a real baseline both were genuine divergences. **Check any differential test
in this repo for this trap before trusting it.**

### Correctness fixes worth knowing

- **`json.RawMessage` ignored JSON null**: a fresh target got nil where the stdlib
  gets the four bytes `null`, and — worse under the reuse this library encourages —
  a reused target silently kept the *previous* document's value.
- **`parseRFC3339` validated the day only as 1..31**, so `2021-02-31` decoded to
  March 3 with a nil error (month was already checked correctly). The failure mode
  is the bad one: a plausible wrong date, not a rejection.
- **`ReadNumberOrNull` captured malformed literals verbatim** (`1.2.3`, `-`, `1e`).
  The fix's acceptance rule is **agreement with this library's own float reader**,
  not with encoding/json — so `01` stays accepted, because `Valid` accepts it and
  changing that would create a *new* Valid-vs-decoder divergence in the other
  direction. That invariant is now a differential test.
- **`SetPaths` edited every duplicate of a matched key**, where `Set`/`SetMany` stop
  at the first — and whether it did depended on whether an *unrelated* path was
  found, because the all-found early exit suppressed it. The 200k-document
  differential that blessed the early exit as "a consistency fix" used unique-keyed
  documents and could not see this. **A differential over inputs that lack the
  feature under test proves nothing about it.**
- **`GetPaths` descended duplicate parents where `Get` stops**, so the documented
  "multi-path form of Get" disagreed with `Get`.
- **`StripDefaults` ejected on a whitespace-only object in array position**, copying
  the entire remainder through unstripped — output still valid JSON, so nothing
  downstream noticed it had quietly stopped working.
- **The checked edit wrappers validated everything except the keys.** `appendMember`
  writes keys raw between quotes, so key `x":1,"role` injected a member and the
  trailing `Valid(res)` passed it *because the forgery is well-formed*. The
  precondition was documented — on the *unchecked* function, while the checked one
  is sold as the untrusted-input path. New `ErrUnsafeKey`; the hot walker is
  untouched, per the two-tier convention.

### Generator: the whole class came from one hole

`main.go` had **0.0% statement coverage** and no test file in the root package.
Four confirmed defects lived in it, all of the same shape — *the generator exits 0
printing "wrote …" and emits a package that does not compile*:
`namedStruct` reserving its decoder name with a bare `g.used[fn] = true` instead of
`g.uniq` (duplicate function names, order-dependent — swapping two fields fixes or
breaks it); `assemble` deciding imports by substring-scanning the generated text,
which a JSON *key literal* containing `time.Time` trivially fools; `isRaw` matching
`RawMessage` from any package qualifier; and directive validation only seeing
directives already attached to a type decl, so a blank line silently disables both
the directive and its typo check. The fix that matters is the **generate-then-compile
table test** (`generator_test.go`) — compilation is the assertion, and it closes the
class rather than the instances. `//lightning:compact` likewise had *zero executed
coverage*: its only test lives in the `bench` module, which `go test ./...` cannot
reach and which every runner enters with `-run='^$'`; it now has a conformance case.

### The SIMD/scalar skip paths are not interchangeable — now pinned

See the corrected note in the `skipContainerFast` entry above. Three divergence
classes, all confined to malformed input, all host-dependent because `SkipValue`
picks the path by CPU feature. The third was found only by an exhaustive
differential (708k malformed documents): a **stray backslash outside a string**,
where `findEscaped64`'s pure bit math cannot know it is not in a string, producing a
verdict that flips on 64-byte-grid alignment alone. Pinned by
`TestSkipPathsDivergeOnMalformed` / `TestSkipBackslashAlignmentCliff` /
`TestSkipDepthDivergence` — the `TestValidDivergesFromStdlib` pattern, which exists
precisely so a documented disagreement cannot rot into an undocumented one.

### `Valid`'s contract was overstated in four places

`Valid` ≡ `DecodeAny` is real and fuzz-locked. "And therefore a generated
`UnmarshalJSON` accepts" was not, in **either** direction, and had propagated to the
package doc, `checked.go` and `valid_test.go`'s rationale. A generated decoder is
stricter where the schema is, and looser in two ways — where it does not look
(unknown fields take `SkipValue`'s bracket balancing) and, less obviously, **where it
does**: `ReadInt64OrNull`/`ReadUint64OrNull` consume a digit-ish run and stop, so a
known int field reads `1.2.3` as 1 with no error while `Valid` rejects the document.

### CI gaps closed

`go vet` ran on amd64 only, so **`asmdecl` never inspected the arm64 assembly** —
and `go test`'s built-in vet subset does not include `asmdecl`, so the NEON side was
checked by nothing. This is exactly the class of the recorded `maskBlock`
result-offset bug (confirmed by experiment: that edit is reported by
`GOARCH=arm64 go vet` and invisible to amd64). New `vet` and `fmt-check` targets;
`bench-test` compiles the otherwise-unreachable `bench/get`.

## Session 2026-08-09 (second pass): nine fixes from a four-domain review

A second correctness review — generator, pure-Go runtime, SIMD/skip, `pkg/json`,
one reviewer each — reported 21 findings; the nine load-bearing ones (3 high, 6
medium) were fixed here. The SIMD and float/escape machinery came back **clean**
under a guard-page OOB harness, a 6.4M-exec fast-vs-scalar differential, and an
exact-arithmetic re-derivation of the Eisel-Lemire table, so nothing below touches
it. As with the first audit, the entries worth carrying are the *reasoning*.

### Generator front-end: four more "exits 0, does not compile" shapes, and one silent no-op

`generator_test.go`'s generate-then-compile table closed the *class*; these are the
instances it then found in the front end (type collection and root selection). All
are output-neutral — generators built from the parent commit and from the fix emit
**byte-identical decoders for conformance and all 30 bench schemas**, which is the
proof any change in this area owes.

- **Generic and alias declarations were collected as roots.** The collect loop
  switched on `ts.Type` and never read `ts.TypeParams` or `ts.Assign`, so
  `type Root[T any] struct{…}` got a method ("cannot use generic type without
  instantiation") and `type Root = struct{…}` got one too ("invalid receiver type
  *Root"). Both now **warn and skip**, following the fixed-size-array precedent in
  the same loop: a schema file may legitimately hold a generic helper or a compat
  alias, and failing the whole run over an incidental declaration is worse than
  naming it. If that empties the file, the existing "no top-level struct, slice or
  map types found" error fires. **One subtlety worth keeping:** an alias to a
  *struct literal* is still registered in `g.structTypes` (but not `g.order`) — it
  cannot take a method, but a decoder taking a `*Legacy` is legal Go and that is how
  such an alias already decoded as a **field** type. Skipping it outright was a real
  regression, caught only by generating the shape under both generators.
- **Every `*ast.InterfaceType` routed to the `any` decoder**, so a field typed
  `interface{ Foo() }` was assigned an `any`. `isAnyInterface` now accepts only the
  empty interface and its spellings (`interface{ any }`, `interface{ interface{} }`
  — the same type); a method, an embedded name, or a type set is reported as the
  unsupported type it is. Note `ast.InterfaceType`'s element list is `Methods`,
  which holds embedded elements and type-set terms too.
- **A schema type named like a generated identifier was captured by it.**
  `type data` met the `[]byte` parameter ("data (parameter) is not a type"),
  `type unstable` collided with the scanner import, `type max` shadowed the builtin
  the un-presized slice's capacity hint calls. Now a hard error naming the type
  (`reservedIdents`/`checkReservedNames`); renaming the generated locals instead was
  rejected as churn across every committed generated file. **The set is derived, not
  guessed** — decoder parameters/locals, the generated file's imports, and the
  predeclared identifiers the generated code names *of its own accord* — and the
  first two groups were cross-checked by parsing the decoders generated for
  conformance plus all 30 bench schemas, which declare exactly that list.
  Predeclared names that reach the output only as an echo of the schema's own type
  text (`bool`, `uint16`, `any` — `typeStr` prints what the field declares) are
  deliberately excluded, which leaves one adjacent class unaddressed:
  `type uint16 struct{…}` used as a field still gets predeclared meaning from
  `isScalar`. Scope check when editing: adding a local to an emitted template means
  adding its name to `reservedIdents`. The rule is deliberately a property of the
  *name* rather than of how far a given schema gets, so the generator now rejects a
  few schemas that happened to compile — a lone `type data` root is one (a
  parameter's scope is the function body, so `*data` in the signature still resolves
  to the type; it breaks only once a *body* says `var zero data`).
- **A mutually recursive pair beside any other type got NO method and no decoder**,
  silently. `referenced` marks both members, so both were skipped as "nested", and
  the `allReferenced` rescue only fired when the cycle was the *entire* file — one
  unrelated type alongside brought the hole back, and the user's code then failed on
  a missing `UnmarshalJSON`. The special case is now the general rule
  (`entryTypes`): start from the types nothing references, mark what they reach,
  promote a cycle nothing emitted enters, recompute until stable.

  **Two decisions inside that promotion are load-bearing, and both exist to make the
  result independent of declaration order** — "reorder two `type` lines and a
  different type becomes decodable" is not a property anyone can reason about.

  *How much to promote: the candidate's whole strongly-connected component.*
  Promoting one member leaves the rest merely reachable from it, so the decodable
  member of `MutA ↔ MutB` would be whichever came first — and the fully-cyclic file,
  which `allReferenced` gave a method per member, would silently lose one on
  upgrade. The SCC is exactly "the cycle", so the general rule is a strict superset
  of the special case it replaced.

  *Which candidate: one whose component is a SOURCE of the still-uncovered
  subgraph* (no other uncovered type outside it reaches it). Taking the first
  uncovered type in source order is **not** enough, and the gap opens only when the
  file has no entry type at all — every type referenced by something, so nothing is
  covered on the first round. Then a record that merely hangs off a cycle is
  uncovered like everything else, and *declared before* that cycle it gets picked
  and handed a method, where declaring it after leaves it covered and correctly
  nested. That is a measured witness, not a hypothetical: the same three
  declarations in two orders produced `{MutA, MutB}` and `{MutA, MutB, Other}`.
  Requiring a source makes the first round behave exactly like a file that does have
  an entry type, which is what keeps a hanger-on nested and its
  `type recordStd Record` reflection baseline reaching `encoding/json` rather than
  the generated decoder. A source always exists while anything is uncovered (only
  uncovered-internal edges can matter — anything referenced by an emitted or covered
  type would itself be covered — so the condensation of the induced subgraph is a
  finite non-empty DAG); among sources the first in source order is taken, and that
  choice is immaterial since distinct sources are mutually unreachable and all of
  them are promoted before the fixpoint ends.

  **The table test cannot see this class** — each case fixes one declaration order,
  so an order-sensitive rule passes all of them, and the first fix's per-shape checks
  did pass while still being order-dependent. `TestEntryTypesOrderIndependent` is the
  guard that can: it generates every permutation of a schema's top-level
  declarations (6 and 120 orders over two shapes) and asserts the set of types
  receiving `UnmarshalJSON` never moves. It runs the generator only, no
  per-permutation `go build`, which is what makes exhausting 5! orders cheap. Point
  it at any future change to `entryTypes`.

### `,lax` swallowed syntax errors, and the fix deduplicated a grammar

`laxField` skipped a failed field's value with `unstable.SkipValue` — a bracket
balancer — so every *balanced but invalid* value was dropped silently and the whole
decode returned nil: `{"l":[1,]}`, `{"l":[1 2 3]}`, `{"l":[1,,2]}` into a `[]int`
field tagged `lax` all gave `err=nil, L=[], X=9`, where the same field without the
tag fails. That contradicted the option's own doc comment and README, punched a hole
through the trailing-comma rejection the rest of the decoder enforces, and was
**host-dependent** on some shapes (SkipValue picks a SIMD or scalar balancer by CPU
feature, and the two differ off the well-formed set — the divergence classes
`TestSkipPathsDivergeOnMalformed` already pins).

The fix needed a *validating* skip and one already existed: `pkg/json`'s
`validValue` is exactly a strict single-value parse. It moved to
**`unstable.SkipValueStrict`** (`pkg/unstable/valid.go`) returning
`(end int, err error)`; `pkg/json.Valid` is now a six-line wrapper and `laxField`
calls it in place of `SkipValue`. **One grammar, two callers** is the load-bearing
part: `Valid`'s contract is "accepts exactly what `DecodeAny` accepts", and a lax
field now inherits that by construction instead of a parallel implementation
drifting alongside. Acceptance unchanged: `FuzzValidMatchesDecodeAny` **14.07M
execs, zero divergence**; `TestValidDivergesFromStdlib` untouched. One deliberate
change is an *error identity* — a `]`/`}` where a value must start is
`ErrInvalidJSON` rather than the number reader's `ErrBadNumber`, so lax reports the
same sentinel a non-lax container loop reports for `[1,]`. The strict walk is a
scalar parse and slower per byte than the SIMD balancer, but runs only after a
decode has already failed; `SkipValue` remains what an *unknown* field's value is
skipped with, where nothing downstream depends on those bytes.

New test worth knowing: `SkipValueStrict`'s **end offset** is pinned against
`SkipValue` on well-formed input (`pkg/unstable/valid_test.go`). `Valid` only reads
the error, so a walk that returned the wrong end would pass every Valid test —
including the 14M-exec fuzz — while making a lax field resume in the middle of the
value it just skipped.

### An explicit JSON null zeroes a leaf field — found, fixed, then reverted to a documented divergence

Every `*OrNull` reader signals a null by returning the **zero value with a nil
error**, and the emitted code assigns it unconditionally — so `{"s":null}` wipes a
string/bool/int/uint/float/`json.Number`/`time.Time` field where encoding/json
documents "unmarshaling a JSON null into any other Go type has no effect on the
value". Invisible on a fresh target (the zero was already there) and observable on
exactly the targets this library promotes: a struct seeded with defaults, and one
reused across documents. Composite kinds were already right — slice/map/pointer/any
nil, `json.RawMessage` takes the literal `null`, nested struct and `[N]T` untouched.

**The parity fix was built, verified, and then deliberately reverted.** Keep the
mechanism, because it is correct and one line away: guard the *assignment*, not
the read — `if data[i] != 'n' { dest = val }` in `nullGuard`. It is exact (only the
first byte separates "was null" from "was an empty string / 0 / false") and in
bounds at every emission site by construction: `i` still holds the value's start
(the templates assign `i = end` only after, and no `*OrNull` reader skips leading
whitespace of its own), and **`data[i]` is in range *because the reader returned
nil*** — every one of them returns `ErrTruncated` when `i >= len(data)`, so the
bound does not depend on the caller. It survives `//lightning:destructive`, whose
in-place unescape writes into `data` while decoding, since the writes start at the
string *body* and never touch the value's first byte.

What sank it was cost against benefit. The guard fires once per leaf field read:
cloudflare's 45-field decoder grew **1620 → 1937 instructions**, and the compiler
could not prove the index, so each guard also carries a bounds check (**+25
`panicBounds` stubs**; 52 of 53 land out of line in the cold tail, but the compare
stays inline). Static estimate ~1–2% on that decoder — under the repo's own noise
floor, and no measurement available could separate it from zero (see the harness
entry below). Against that: because the two rules coincide whenever the target
starts out zeroed, the guard makes *every* decode pay to change behavior only for
callers who seed or reuse a target **and** receive an explicit null. Documented
beats paid-for-by-everyone. The README's divergence list now carries it, and
`TestNullFieldsDivergeFromStdlib` pins it the way `TestValidDivergesFromStdlib`
pins Valid's: the expectation is a methodless twin's result with a **named**
transformation (`zeroNullLeaves`) applied, so the test fails both if the
divergence widens and if it silently closes.

**`laxField` has to track whatever `nullGuard` does** (`nullAssigns`), because a
lax value decodes into a fresh zero scratch and then commits it. With the guard
reverted, a lax leaf commits unconditionally (matching the plain leaf's zeroing)
while a lax nested struct or `[N]T` stays guarded (matching the plain form's
"leave it alone"). The invariant to preserve is not "follow encoding/json" but
**`json:"n"` and `json:"n,lax"` never disagree** — restoring the parity guard means
moving the three leaf kinds back to the guarded side in the same change.
### A null at the ROOT left a named slice or map alone

`genUnmarshal` intercepts a null document and returns before calling the root
decoder — whose own null arm is the thing that nils a slice or map. So
`List{1,2}.UnmarshalJSON("null")` left `[1 2]` where the stdlib gives nil, while the
same named type used as a *field* was nil'd by the very arm the intercept skipped.
Fixed with a per-kind `nullReset` (`*v = nil` for slice and map roots, nothing for
struct roots, which the stdlib also no-ops). Conformance had a null-root test that
passed either way because it decoded into an *already nil* root — the same "a
differential over inputs that lack the feature under test proves nothing" lesson as
the SetPaths duplicate-key finding.

### A failed base64 decode left `[]byte` reporting a stale length over rewritten bytes

`decodeByteSlice` reuses `*out`'s backing when the decoded bytes fit, and
`base64.Decode` fills that backing quantum by quantum before it reports a bad one.
Returning the error without touching `*out` therefore never preserved the caller's
previous value — the bytes are already gone — it only left the old **length**
describing rewritten bytes, so a caller that handles the error and keeps using the
value saw a silent splice of two documents: `"QUJD####"` into a retained
`"0123456789abcdef"` left `"ABC\x00456789abcdef"`. It now publishes `b[:n]` on every
return, the partial-progress convention every other reader in `batch.go` already
follows (`*out = s` on all paths, errors included), so `len(*out)` says exactly what
decoded and the semantics no longer depend on whether the target had spare capacity.
encoding/json leaves its target untouched because it always decodes into a fresh
buffer; that divergence is documented on `DecodeByteSlice`.

### Measuring the null guard: interleaving two *binaries* does not control for layout

The guard was the only change in this pass that adds hot-path work, and measuring
it is the reason it was reverted rather than kept on faith. Two lessons outlived
the change itself.

**Interleaving two binaries does not control for link layout.** Cross-binary
interleaved runs (n=14 and n=30, and an ABBA order-alternating variant) put
cloudflare at +0.8% to +4.0% with robust paired estimators reading +2.5…+3.5%,
while putting **both decoders in ONE binary** — rename the schema's top-level
types, as `run_bench.sh` already does for its destructive/arena twins — measured
−1.2% (p=0.756). That is the same "adding code shifts a micro-benchmark a few %
with no change to executed instructions" effect the float-array entry records,
showing up as a *between-binary* rather than a within-binary artifact. A second
methodology note from the same exercise: a fixed base-then-opt arm order *within*
each rep biases every rep the same way, so alternate it (ABBA). Interleaving alone
is not enough — position within a rep has to alternate too.

**Same-binary is still not enough when the two decoders are different functions.**
A follow-up role-swap experiment — build the pair twice, swapping which *type name*
carries the guard, so position and benchmark order stay fixed and only the guard
moves — produced medians saying the guard costs 3–10% and minimums saying it saves
3–5%, *in both orientations*. Contradictory estimators are the signature of a
noise-dominated measurement, and the machine was the reason: an unrelated `ffmpeg`
transcode was holding ~11 of 16 threads at load ~15. **Check what else is on the
box before trusting any of these numbers** — the harness cannot rescue a busy
machine, and "interleaved" is not a synonym for "valid".

What did decide it was static and load-independent: instruction counts and an
**identical call structure** (same 14 `ReadStringOrNull`, 11 `ReadInt64OrNull`, 3
`SkipValue`, 17 write barriers, 2 `memmove` — no lost inlining, no new frames, so
the guard did *not* push the wide decoder past the inliner's big-caller threshold,
the documented failure mode for `string_unicode`-sized decoders). A hoisted variant
(probe `data[i]` *before* the reader call) does not let the compiler prove the
index either — still 53 `panicBounds` — and saves only 57 instructions, so it is
not the way out if this is ever revisited.
### Testing note that generalized

`time.Time` implements `json.Unmarshaler`, so `TestStdlibTwinsAreReflectionOnly`'s
guard (it fails on any Unmarshaler in a twin's reflect graph) rejects any twin
containing a `time.Time` field. The exception list is now `{json.RawMessage,
time.Time}` — both standard-library-owned, i.e. the behaviour the differentials
measure lightning *against* rather than a way for the comparison to short-circuit
back into lightning's own decoder. Any future twin holding a stdlib type with its
own `UnmarshalJSON` needs the same treatment.

### The in-place mode of a rewriter is a separate contract from its output

`StripDefaults` documents `output == input[:0]`, and the walk earns it by never
writing past its own read cursor — writes only ever land on consumed bytes. One
member shape violated it: a member whose value is a container has its key, colon
and the recursion's whole output written speculatively and rewound when the value
strips to nothing, and the code after the rewind re-read the key (for `keepKey`)
and the member's original span (to re-emit a kept member) out of the input it had
just overwritten. Under any left shift — a dropped sibling ahead of it, or removed
whitespace — the member silently vanished from the output or came out garbled,
*while the output stayed valid JSON*. The fresh-buffer path was always correct, so
every committed test passed; `TestStripDefaultsInPlace` existed and happened to
use an input with no keep-container member.

Two lessons. **(1)** In-place is not "the same function with a different buffer":
it needs its own property test, and the right one is **`in-place ≡ fresh, byte for
byte, for every input`** — a differential over generated documents ×
defaults/keep pools × every whitespace mode, checked for **non-vacuity** (17.8% of
the generated documents actually reach the shape under test; 3M pairs run during
development, 120k committed). **(2)** The invariant to restore is "nothing the code
will still read has been overwritten", not "add a copy somewhere". Here that means:
hoist the decision that is a pure function of already-read bytes (the keep test),
and snapshot only the span a *later* read needs, on the one path that can need it —
a `stripper.scratch` grown with `append(s.scratch[:0], …)`, so an ordinary member
allocates nothing and 0 allocs/op survives. One shared scratch is enough for all
nesting levels because a deeper frame replaces it only by snapshotting a kept
member of its own, and such a member is always emitted, which leaves every
enclosing container non-empty — so no outer frame ever reaches back for the
snapshot it lost. That argument is the load-bearing part; it lives next to the
buffer.

The snapshot's end comes from `SkipValue` over the same value, which agrees with
the walk on every well-formed one. They part on malformed input — a bare token
holding a quote (`{"b":{"k":q"w,"a}b":0}}`) makes `SkipValue` open a string the
walk does not, leaving it outside a string where the walk is inside one, so it
stops at an earlier `}` — hence the guarded fallback to emitting from the input.
Reachable, and tested.

**Cost control on a fix like this took three measured rounds.** Parameterizing
`emitField` with a `src`/`base` pair costs **88** against the inliner's 80, so
`emitField` stops inlining and every *kept member's* emit pays a call frame — the
obvious "DRY the two forms" move is the expensive one here. Expanding the early
decision inline grew `handle` by **9.4%** of its instructions and cost the pretty
benchmark ~4%; calling a helper per container member cost the same again. Landed:
`emitField` byte-identical, both new arms out of line, and `keepKey`'s own length
pre-filter leading the container branch so an ordinary container member pays one
compare. `handle` +3.0% instructions; all three StripDefaults benchmarks flat
(n=16, pinned, geomean −0.4%, p ≥ 0.10) and still 0 B/op. **Instruction count of
the hot function is the noise-free proxy** to lean on when the box cannot resolve
2% — the same conclusion the null-guard measurement reached from the other
direction.

### `Valid` is not a test for "the transform produced nothing"

`StripDefaultsChecked` gated on `len(res) > 0 && !Valid(res)`, which reads as
"empty is the documented consumed-document case, anything else must be valid". It
is wrong for `PreserveWhitespace`, which keeps the document's outer whitespace: a
fully-stripped document leaves a whitespace-only, *non-empty* remainder, so the
wrapper returned `ErrInvalidJSON` for input it had just validated. The predicate
has to be "holds no token" (`unstable.SkipWS(res, 0) == len(res)` — the same
lenient `<= 0x20` notion the stripper copied those bytes through with), and the
wrapper normalizes to an empty slice so `len()` stays the caller's one test in
every mode.

### Escaping has no in-place form — and the doc said it did

`EscapeStringInto`'s doc said it mirrors "the in/out convention of
`UnescapeStringInto`", whose doc explicitly blesses `out == in[:0]`. That is safe
only because unescaping *shrinks* — the same property `//lightning:destructive` is
built on. Escaping lengthens, so an aliasing `out` overruns input the scan has not
read and then rescans its own output: `EscapeStringInto(s, s[:0])` on `he"llo`
gives `he\"\"\\\"`. Fixed in the doc on both sides (the non-overlap requirement
here, the reason and a pointer back on `UnescapeStringInto`), with no runtime
overlap check on the hot path, per the two-tier convention. Whenever a doc points
at a sibling's buffer convention, check the direction of the size change first.

### Two more stdlib divergences, pinned rather than "fixed"

- **`ReadTimeOrNull` unescapes; `time.Time.UnmarshalJSON` does not.** Its doc
  claimed to match "how encoding/json decodes time.Time". It reads the string
  *value* and parses that, where the stdlib parses the raw quoted bytes without
  unescaping (`TODO(https://go.dev/issue/47353)` in time.go), so any `\uXXXX` in a
  timestamp — legal JSON for a legal instant — is accepted here and rejected there.
  Tightening would be a regression, so the doc was scoped and
  `TestReadTimeAcceptsEscapedTimestamps` pins both halves. **A second correction
  came out of writing it:** the stdlib's authority is `parseStrictRFC3339`, not
  `time.Parse`, and its extra RFC 3339 checks are currently *compiled out*
  (`case true: return t, nil`, pending go.dev/issue/54580). So "lightning ≡
  encoding/json on escape-free input" is true on this toolchain and is a **premise,
  not an identity** — the test runs the whole `dateCorpus()` through both, so
  re-enabling 54580 surfaces here rather than as a silent new divergence.
- **Raw invalid UTF-8 passes through where encoding/json coerces to U+FFFD.**
  Stated in the README since forever, tested by nothing.
  `TestStringsPassInvalidUTF8Through` now covers all seven paths that could differ
  (copy / nocopy / destructive readers, `ReadKey`, `DecodeValue` value and map key,
  `decodeEscaped`'s literal runs, `UnescapeString`) and asserts the stdlib premise
  in the same test. Sabotage-verified in both places it could be closed — the
  reader's `string(rest[:k])` and `decodeEscaped`'s `unsafeStr(buf)`.

Prose rot found in the same pass, all verified before editing: `simd.go` /
`simd_other.go` pointed at a `simd_noasm.go` that does not exist (the scalar
dispatch file is `simd_scalar.go`); `skipfast_noasm.go` still called the whole-loop
block scan "amd64-only" after `skipBlocksNEON` landed; and `SkipValue`'s doc
re-enumerated `skipfast.go`'s divergence classes and listed two of three — the exact
cross-file-copy rot the audit section warns about, now replaced by a pointer to the
single authoritative list. `skipNumber` gained a seam note: it *measures* a number
token rather than validating one, so `SkipValue([]byte("+"), 0)` is `(1, nil)` and
`-`, `e`, `.`, `1.2.3` are spans too — consistent with the bracket-balancer
contract, but it reads as validation without the note. `ParseFloat`/`DecodeValue`
docs likewise now state their accept set (`+5` is 5) instead of saying "the JSON
number in b"; the behavior is `Valid`-consistent and deliberate.

**Method note worth keeping.** The claim "these are comment-only changes" was not
left to `git diff`: a probe binary exercising the touched entry points was built
against the tree before and after, and the normalized whole-binary disassembly
(175 569 instruction lines) hashed identically. The raw binaries *do* differ —
pclntab and DWARF line numbers shift when you insert comments — so a naive binary
comparison would have looked like a code change and a naive `diff` of the source
proves nothing about codegen. Disassemble and normalize.

### Generator front end: two silent stdlib disagreements, one real bug, one wrong doc

Same area as the entry above, same proof obligation — parent-vs-tree generators
over conformance plus all 30 bench schemas: **30 of 31 byte-identical**, the one
that differs adding only the unwrap check below (+25/−0, five identical sites), and
identical diagnostic streams (no committed schema has a tag name the new check
flags, so CI stderr does not change).

- **`,unwrap` tolerated trailing garbage in the wrapped body.** The generated
  closure decoded one value and stopped, so `"{\"n\":1} trailing garbage"` decoded
  happily while the identical bytes decoded as a *root* are `ErrInvalidJSON`. The
  body is a whole document, so `unwrapField` now ends with `genUnmarshal`'s own
  `unstable.SkipWS(data, i) != len(data)` check and the same sentinel. Two
  neighbours it must not disturb, both covered: the whitespace-only-body guard
  returns *before* `inner` (so `TestUnwrapWhitespaceBody` is untouched), and
  `unwrapField` wraps the **lax** code too — a wrapped value of the wrong *type* is
  still swallowed, while trailing content is malformed JSON, which lax has never
  tolerated. Locked by `TestUnwrapRejectsTrailingContent`, which decodes every body
  as a root as well, so the wrapped path is pinned to the root's answer rather than
  to a hardcoded expectation.
- **Tag names encoding/json rejects are honored here — warned, not "fixed".**
  `isValidTag` discards a whole tag holding a character outside its allowed set and
  keys the field by the **Go field name**; lightning matched the name as written, so
  the two answered to different keys *in both directions* with no error either side.
  Both candidate fixes are worse than the warning: rejecting fails a schema that
  decodes correctly today, and adopting the field-name fallback silently moves which
  key an existing decoder answers to. `invalidTagRune` copies the stdlib set
  verbatim — note it contains `|`, which is exactly what makes checking each
  pipe-separated name separately equivalent to checking the joined tag.
- **Field promotion is not method promotion.** An embedded `time.Time` or
  `json.RawMessage` decodes here as a named field keyed by the type name, while Go
  promotes the embedded `UnmarshalJSON` to the outer struct, so `encoding/json`
  hands it the entire document: loud for `time.Time`, **silent** for
  `json.RawMessage` (whole document into the embed, siblings zero). An embedded
  `json.Number` is the control — no method, no divergence — which is what makes the
  cause attributable to the promoted method rather than to foreign embedding. Note
  the methodless twin `type FooStd Foo` still inherits the *embedded* type's
  methods, which is precisely why it is the right baseline here.
- **The README's "embedding a foreign type is decoded as a single named field" was
  false**: generation *fails* (`unsupported type strings.Builder`). The failure mode
  is right; the doc was wrong, and only the three known selector types embed at all.

### Toolkit consistency pass: the same span, read twice

Five low-severity findings in `pkg/json`, four sharing a shape worth naming: **a
value the walker emits without having decided every byte of it**.

- **`StripDefaults`' keep path re-emitted a member's original input span**, so
  `RemoveWhitespace` handed back whitespace it promised to remove
  (`{"b":{ "x" : 0 }}` → `{"b":{ "x" : 0 }}`). The repo's own oracle
  (`compact(preserve) == remove`) could not see it because no shipped test
  generated a keep-listed container that empties *with interior whitespace*. **An
  oracle only tests the inputs you generate; when one exists, ask what shapes it
  never reaches.** Fixed with `emitKeptCompact` + `compactValue`, both off the hot
  path — and `handle` came out ~155 instructions *smaller*, because the keep path's
  inlined `emitField` copy became a call. Outlining a cold branch is a real win
  here, not a wash.
- **`SetMany` was the odd one of the family in two ways**: its non-object-root
  branch replaced the whole *input* instead of the root *value* (dropping leading
  whitespace and trailing bytes `Set`/`SetPaths` keep), and a key listed twice
  appended a second member, producing a **duplicate-key document** no later `Valid`
  would flag. `TestSetManyMatchesSet` locked byte equality on object roots only —
  the branch it did not cover was the divergent one. Both fixed toward SetPaths'
  answer (first request wins); `SetMany` is ~60 instructions smaller than before.
- **`GetPaths` — adding a request changed another request's outcome.** Descending
  for a deeper path is stricter than skipping a value (`Get(doc,"a")` reads
  `{"b" 1}`; `Get(doc,"a","x")` reports `ErrExpectColon`), so co-requesting `{a}`
  and `{a,x}` returned `[nil nil]` plus an error. Capture-before-descend fixes only
  the same-key case — a path at a *later member* still lost its value. `walkPaths`
  now returns `(end, err, fatal)`: a failed descent is re-skipped with the same
  lenient `SkipValue` a solo lookup uses, the walk continues, and the failure is
  held as the frame's first error and still returned. Nothing is swallowed; each
  path gets what it would have got alone. Locked by a randomized solo-vs-combined
  differential — **20 354 mismatches over 300k documents before, 0 after** — which
  is the kind of check worth writing precisely because the property ("each path
  behaves as if requested alone") is otherwise only assertable one hand-picked
  shape at a time.

**Integration note, and the reason it needed care.** The compaction fix and the
in-place snapshot fix were written on separate branches from the same parent, and
they touch the same three lines: one re-emits the kept member's original bytes, the
other changes *where those bytes live*. Composed naively, the compaction would read
`in` while the snapshot existed precisely because `in` had been overwritten.
`emitKeptCompact` therefore takes `(src, base)` and compacts out of whichever buffer
holds the original — snapshot in place, input otherwise. The proof that they compose
is that the three cases the compaction work had to exclude from its in-place
assertions now pass with the exclusion removed; that flag is the artifact to look
for when merging two fixes to one path.

Composing them also cost the zero-alloc contract, since a kept container member
always snapshotted. **The snapshot exists only because writing the output over the
input destroys bytes a kept member still needs**, so it is now taken only when the
output actually aliases the input (`unstable.SameBuffer`, which answers "same
backing array" and is conservative in the safe direction — anything it cannot prove
separate still snapshots). With an output buffer of its own the input is only ever
read, so the copy had no reader. All twelve StripDefaults/Set benchmark rows are
back to 0 allocs/op.

Method note that paid off twice in this pass: **before claiming a divergence is
yours, re-run the same input against the parent commit.** The compaction agent
correctly identified its in-place divergences as pre-existing rather than
self-inflicted, which is what made the merge a composition problem instead of a bug
hunt. Note also that agent worktrees branch from *origin/main*, not local `main`, so
parallel fixes see none of the merges that happened while they ran — check the
branch parent before assuming a report's baseline matches the tree.
