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
- `pkg/json` — small public API over the scanner (`Get`/`GetMany`/`ObjectEach`
  + `*Compact`, `Unwrap`, `UnescapeString`, `ParseFloat`).
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
- **`CountArrayElements`** (slice presize) skips each element with `SkipValue`
  (vectorized via `indexStructural`), not a byte-by-byte depth walk. Element types
  whose JSON can hold no comma or bracket use the much cheaper `CountArrayScalars`
  (find `]`, count commas — two vectorized scans, no per-element work): bare
  numbers/bools, `json.Number`, **and `time.Time`** — an RFC 3339 / Unix-timestamp
  value never contains a `,` or `]`, so a `[]time.Time` sizes by comma count. That
  avoids a `skipString` over every element (which re-scanned each date string just
  to count it, doubling the per-date work): time-array −16%.
- **`CountArrayObjects`** extends that to a *struct* of only number/bool fields
  (`isBracketFreeStruct`): its JSON (`{"a":1,"b":2}`) holds no string, `[`, `]`, or
  nested `{`, so the array's `]` is the first `]` and the element count is the
  number of `{` before it — again two vectorized scans instead of a `SkipValue`
  per struct. citm_catalog's price entries hit this: −6% with *no* change in
  allocations (the slice is still presized exactly, just counted cheaply). Any
  string/array/pointer/nested field disqualifies the struct (its JSON could carry
  a bracket); those keep `CountArrayElements`. All these counters are presize
  hints — a wrong count only mis-sizes the slice, never misdecodes.
- **`slicePresize`** skips presize when a struct element transitively holds a
  *multi-dimensional* slice (`hasMultiDimSlice`, e.g. GeoJSON `[][][]float64`):
  counting it would deep-scan every element's bulk for only ~log2(n) reallocs
  saved. Flat / string / 1-D-slice records (Cloudflare-style) keep presize and
  their low B/op.
- **SIMD scanners** in `index_amd64.s`/`index_arm64.s`: `indexCloseOrEscape`
  (next `"`/`\`, SSE2 — no VZEROUPPER), `indexStructural` (next `{}[]"`, AVX2).
- **Date parsing**: `daysFromCivilCached` uses a year-start-days table (built from
  `daysFromCivil`) for 1970–2261; falls back to the general algorithm otherwise.

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
- **Pure-SSE2 `indexStructural`** (dropping AVX2): ~2× slower on the skip path
  (`skip-heavy`); the throughput loss dwarfs the VZEROUPPER saving.

## Conventions

- Bench `data.go` files use a single top-level `Benchmark` with **anonymous**
  nested structs, so only `Benchmark` gets a generated method and
  `type benchmarkStd Benchmark` gives a clean reflection-only baseline for the
  stdlib/sonic benchmarks.
- End-of-session: run the full suite (`go test ./...`), keep gofmt clean
  (the `daysFromCivil` comment-alignment flag is pre-existing — ignore it),
  and regenerate `bench/results_<goarch>.md` via `run_bench.sh`.
