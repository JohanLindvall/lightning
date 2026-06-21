# CLAUDE.md

Guidance for working on lightning — a code generator that emits fast,
allocation-light `json.Unmarshaler` implementations.

## Layout

- `main.go` — the generator (`package main`). Reads struct defs, emits
  `*_unmarshal.go`. Key bits: `field`/`sliceDecoder`/`mapDecoder` build decoders;
  `slicePresize` decides presizing; `g.compact`/`g.skipWS` handle the
  `//lightning:compact` directive (compile-time elision of inter-token `SkipWS`).
- `pkg/support` — the runtime scanning primitives every generated decoder calls.
  This is where almost all performance work happens.
- `pkg/json` — small public API over the scanner: `Get`/`GetMany`/`ObjectEach`
  (+ `*Compact`) read fields, `Set`/`SetMany` splice values, `StripDefaults`
  prunes default members, `DecodeAny` decodes a whole document into `interface{}`
  (schema-less), plus `UnescapeString`/`EscapeString` and `ParseFloat`.
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
- **Escaped-string decoding** (`decodeEscaped` / `readUnicodeEscape` /
  `decodeStringEscaped`) — three things make densely-escaped text fast
  (`\uXXXX`-per-character CJK, the twitterescaped workload): **(1)** the
  literal-run scan is skipped when already sitting on a `\` or `"` — consecutive
  escapes land on a `\` every other byte, so calling the SSE2
  `indexCloseOrEscape` each time just to find `\` at offset 0 was pure call
  overhead; **(2)** `readUnicodeEscape` parses the four hex digits through a
  256-byte `hexNibble` table (nibble value, or `0xFF` for a non-hex byte) — four
  loads + one `a|b|c|d >= 0x80` test replace four 3-way comparison chains, and
  validation falls out of the lookup because the `0xFF`'s high bit survives the
  OR; **(3)** the buffer cap hint in `decodeStringEscaped` finds the body end
  with a plain `bytes.IndexByte('"')` (one vectorized pass) instead of the
  escape-aware `skipString`, which stops at *every* backslash and re-scans the
  whole string. A `"` not preceded by `\` is definitively the unescaped close, so
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
- **`slicePresize`** skips presize when a struct element transitively holds a
  *multi-dimensional* slice (`hasMultiDimSlice`, e.g. GeoJSON `[][][]float64`):
  counting it would deep-scan every element's bulk for only ~log2(n) reallocs
  saved. Flat / string / 1-D-slice records (Cloudflare-style) keep presize and
  their low B/op.
- **SIMD scanners** in `index_amd64.s`/`index_arm64.s`: `indexCloseOrEscape`
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
  arm64 mirrors the change (correctness verified under qemu — full support suite +
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

## The inline trick — let the generator write hot bodies inline

Go's inliner refuses `SkipWS`, `ReadKey`, and `indexCloseOrEscape` (each exceeds
the cost budget once it holds a SIMD/asm call), so calling them from generated
code pays full call overhead per token. Instead the **generator writes the hot
fast path inline at the call site** and falls back to a support call only for the
rare/hard case. See `g.skipWS` and `g.readKey` in `main.go`:

- **skipWS** emits the whitespace check inline:
  `if i < len(data) && data[i] <= ' ' { i++; if i < len(data) && data[i] <= ' ' { i = support.SkipWSRun(data, i+1) } }`.
  The common 0–1 whitespace bytes cost one or two compares and no call; only a run
  of ≥2 (pretty-print indentation) reaches the SWAR `SkipWSRun`.
- **readKey** emits the no-escape key read inline — a `support.IndexCloseOrEscape`
  scan plus a `support.UnsafeStr` alias — falling back to `support.ReadKey` only
  for an escaped key or an error. It relies on tiny inlinable exported wrappers
  `IndexCloseOrEscape`/`UnsafeStr` that themselves inline, so the generated code
  reaches the SIMD scanner with no wrapper call. Won ~6–7% on the cloudflare
  family, no regressions.

Why it beats making the support funcs inlinable: the cheap common case skips the
call entirely, and the shared scanner keeps its tuned dispatch.

**Scope it to once-per-loop reads.** `skipWS`/`readKey` fire once per object member,
so the inline block stays small. Inlining a *per-field* read (e.g. string values)
emits the block once per struct field, which bloats a wide struct's decoder
(`string_unicode` has 60 string fields): the function grows past the inliner's
per-function budget, `IndexCloseOrEscape` stops inlining, and i-cache pressure
turned a −8% cloudflare win into a +9% regression there. Don't inline per-field
reads.

## Tried and rejected (don't re-attempt without a new idea)

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
  newly-presized slice.
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
  appears. The plain 8-byte SWAR is already near-optimal for these run lengths.
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

- Bench `data.go` files use a single top-level `Benchmark` with **anonymous**
  nested structs, so only `Benchmark` gets a generated method and
  `type benchmarkStd Benchmark` gives a clean reflection-only baseline for the
  stdlib/sonic benchmarks.
- Every case also gets a `BenchmarkLightningDecodeAny`: the same document decoded
  through `json.DecodeAny` (compact mode) into the generic `any` value instead of
  the typed `Benchmark`, so the rendered table contrasts schema-less decoding with
  the generated unmarshaler. The harness minifies each input with `json.Compact`
  first (many corpus inputs are pretty-printed) so the compact path is valid.
- End-of-session: run the full suite (`go test ./...`), keep gofmt clean
  (the `daysFromCivil` comment-alignment flag is pre-existing — ignore it),
  and regenerate `bench/results_<goarch>.md` via `run_bench.sh`.
