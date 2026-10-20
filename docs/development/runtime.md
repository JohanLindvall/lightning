# Runtime internals

[Documentation](../README.md) · [Contributing](../../CONTRIBUTING.md)

Contracts and performance constraints for `pkg/unstable`.

## Numbers

- **`scanFloat` is a straight-line fast path** in front of `scanFloatSlow`, the loop
  form that is both fallback and oracle (`TestScanFloatFastMatchesSlow`). Fast shape:
  optional `-`, 1–7 integer digits, optional `.` + 1–24 fraction digits, a 1–5-digit
  exponent, ≤ 19 significant digits, and a 48-byte window in bounds; anything else
  goes whole to the slow path.
  - Each digit run gets one mask — `d := w ^ swarZero`, then
    `((d + swarSix) | d) & swarNib`, exact in its lowest flagged lane and as an
    all-digits verdict — and one `parse8Digits`.
  - The fraction words all load at `i+k+1`, an offset that depends only on the
    integer count, so loads and folds run in parallel. The exponent folds by shifts,
    so a token is never handed off after fast-path work. Leading fraction zeros are
    discounted only within the first word (that can send a number to the slow path
    early, never through wrongly).
  - Clinger (`mant>>53 == 0 && uint(exp+22) <= 44`) and Eisel-Lemire are inline,
    held bit-for-bit to `eiselLemire64`; the power-of-ten entry is taken by pointer
    (a copied `[2]uint64` can't store-forward).
  - `BenchmarkScanFloatShapes` must beat `BenchmarkScanFloatSlowShapes` on every
    shape.
- **Float tiers**: Clinger (exact mantissa < 2^53, |exp| ≤ 22, one multiply or
  divide) → Eisel-Lemire (`powers_table.go`; bit-identical to strconv whenever it
  returns ok) → `strconv`. Locked by the differential fuzz
  `TestParseFloatMatchesStrconv`. Don't remove EL. In `scanFloatSlow`, leading
  fraction zeros don't count toward the 19-digit budget (while `mant == 0` they only
  shift `exp`); a zero between nonzero digits is significant.
- **The four-digit SWAR step** (`tryParse4Digits`; `is4Digits`/`parse4Digits` remain
  only as its oracle, `TestTryParse4DigitsMatchesIs4Digits`) serves `scanFloatSlow`'s
  fraction loop and the batch integer loops. `n*10000+v` is bit-identical to
  `n*10+d`, wrap included.
- **Integer readers** (`ReadInt64OrNull`/`ReadUint64OrNull`) choose per GOARCH
  through `readerWordFold`:
  - off amd64, `digitRun` folds a word at a time — mod 2^64, so bit-identical to
    `n*10+d` with wrap and narrowing; cost 177, a leaf call; a run ending on a word
    boundary needs Go's unmasked shift, not `&63` (`TestDigitRunWordBoundary`);
  - amd64 keeps a byte loop, because a word fold would add a ~14-cycle chain to the
    cursor, behind an eight-digit step: test the byte at `i+7`, test one all-digits
    word, `parse8Digits`, advance by a constant 8. After a short number in an
    object, the byte at `i+7` is rarely a digit, so the gate avoids most failed folds. Tests: `TestIntReadersMatchByteLoop`,
    `BenchmarkReadIntShapes`.
  - Every byte loop accumulates as `n *= 5; n = int64(d) + n<<1` (two LEAs where
    `n*10 + d` is three). Don't "simplify" it: a `uint64` accumulator hoists the
    `- '0'` into four LEAs, and `n*5*2 + d` folds back.
- **`ParseInt`/`ParseUint`** (`pkg/unstable/parseint.go`) know the token length:
  ≤ 19 digits can't overflow the `uint64` fold, 20 need `bits.Mul64`, longer is valid
  only through leading zeros (a retry strips them). Words fold from the right; a
  4–8-digit token is two overlapping `load32`s ORed; nothing reads outside `b`;
  lengths dispatch on unsigned range tests. Tests:
  `TestParseIntMatchesStrconvUnstable`, `TestParseIntAtBufferEnd`,
  `FuzzParseIntMatchesStrconv`.
- **`isNumberByte`** (`SkipNumber`'s accept test) is a `[256]bool` table off amd64
  and six compares on amd64. The choice keeps the acceptance test appropriate to
  each architecture's instruction and dependency costs. `SkipNumber` must stay inlinable
  (`TestIsNumberByteMatchesComparisons`, `TestSkipNumberSpan`).
- **Batched scalar-array readers** (`batch.go`: `DecodeFloat64Slice`,
  `DecodeIntSlice`/`DecodeUintSlice`, `DecodeFloat64Array`/`DecodeIntArray`/
  `DecodeUintArray`, `DecodeByteSlice`). `batchSliceFn`/`batchArrayFn` route every
  slice or fixed-array field whose element is a bare float64/int/uint kind here
  (float32/bool keep the generated loop).
  - Runs go to the [SIMD kernels](simd.md#native-number-kernels) first, then to a scalar loop that
    calls the private `scanFloat` directly and inlines the integer fold: one
    unchecked four-digit attempt, a four-digit loop entered only past four digits (so
    short elements stay straight-line), a byte tail.
  - Semantics match the generated loop exactly (null root → nil slice / untouched
    array, null element → zero, overflow wraps, a truncated fraction is tolerated),
    locked by `batch_test.go`.
  - The slice readers work on a local header copy (`s := *out` — the compiler can't
    prove `*out` doesn't alias `data`) and store `*out = s` on every return, errors
    included.
  - A fresh `[]float64` starts in a stack `[32]float64`: if the array closes inside
    it, the slice is allocated at its exact size and never counted; otherwise the
    reader presizes and copies the prefix. The path is gated on ≥ 80 bytes left and a
    digit or `-` first, keeping the setup off small-slice paths. `-gcflags=-m` must
    show the buffer staying on the stack.

## Strings

- **Escaped strings** (`decodeEscaped`, `readUnicodeEscape`, `decodeStringEscaped`
  in `string.go`):
  - On `\` or `"` the literal-run scan is skipped (dense escapes land on `\` every
    other byte).
  - `readUnicodeEscape` reads hex through the `hexNibble` table (`uint32`, invalid
    marker `1<<16`) combined as `t[a]<<12|t[b]<<8|t[c]<<4|t[d]`, so one `>= 1<<16`
    compare validates all four digits and the function inlines (cost 65). The marker
    must sit at bit 16: `0xFF<<8` slips under the test.
  - Escape dispatch is one `unescByte[256]` load with `'u'` in the fallback; a
    `switch` reaches `'u'`, the hottest case, last.
  - BMP runes are UTF-8-encoded inline; `utf8.AppendRune` only handles supplementary
    runes from valid surrogate pairs; unpaired surrogates map to `RuneError`
    explicitly.
  - The buffer cap hint finds the closing quote with `bytes.IndexByte` from the
    first escape, resuming past a `"` preceded by an odd backslash run; only
    truncated input falls back to `SkipString`. The hint never changes acceptance.
  - The buffer (a chunk carve on the quoted path, a `make` on the unquoted one) is
    returned via `unsafeStr(buf)`; it is never retained or mutated, so `string(buf)`
    would only add an allocation and a copy.
  - `Unwrap` probes through an `unsafeBytes` alias and copies only in the arms that
    return the body itself.
- **Quoted escaped strings are carved from a pooled chunk** (`escapeScratch`/
  `escapeRelease`, `escbuf.go`), not one `make` each. Carves are exact (`cap == n`)
  and never reused, so only retention differs from `make`, and a rule bounds it: a
  chunk is ≤ `escapeChunkMax` (16 KiB) and ≤ the document length (floor 1 KiB), and a
  body over `escapeMaxCarve` (4 KiB) gets its own `make`. The chunk lives in a
  `sync.Pool`, so each `Get` gets an exclusive bump — and pools empty at every GC, so
  keep no adaptive state in the pooled object. The win is GC marking and span
  acquisition, which scale with chunk bytes rather than allocation count;
  `BenchmarkEscapeScratch` sees neither, so don't argue the design from it. Locked by
  `TestEscapedStringsSurviveChunkReuse`. Preserve the retention bound when changing
  chunk sizing.
- **"No escape" in ≤ 32 bytes takes word loads** (`NoBackslash4`/`8`/`16`, split by
  length so each inlines, cost 56): two words at each end of a 4–32-byte body prove it
  escape-free without the ~2 ns `bytes.IndexByte` call. They sit behind one
  `uint(n) < 33` gate in `json.String` (on the quoted token) and in
  `UnescapeString`/`…Into`/`…Copy`. The windows cover every byte, so a miss is
  conclusive (`String` then goes straight to `UnescapeStringScan`) — and a wrong
  bound silently returns escapes verbatim, so set bounds by what the windows
  *cover*, not what they can reach.
- **`//lightning:destructive`** (`ReadStringDestructiveOrNull`) unescapes into the
  input — `buf := data[i:i:len(data)]` — so the result aliases `data` with no
  allocation. Safe because unescaping only shrinks (each escape is ≥ 2 input bytes →
  ≤ 3 output; `\uXXXX` 6 → ≤ 3), so writes trail reads; the cap is the document tail,
  so `append` never moves away from `data`; and the closing quote, which the writes
  never reach, still bounds the value. It destroys the input, hence opt-in. Only
  nocopy string leaves change; escape-free input decodes exactly as nocopy. Tests:
  conformance `TestDestructiveDirective`, the destructive arm of
  `TestReadStringOrNull`.

## Slices: presize, growth, reuse, arena

- **Presize counters** (`count.go`, chosen by `slicePresize`) return hints: a
  miscount mis-sizes, never misdecodes.
  - `CountArrayScalars` — elements that can't contain `,` or `]` (numbers, bools,
    `json.Number`, `time.Time`): one `countKernel` pass, commas + 1, clamped to
    `(rb+1)/2`; a comma-free span counts 1 only if it holds a byte > 0x20.
  - `CountArrayObjects` — flat structs of number/bool/string fields
    (`isFlatScalarStringStruct`): `{` before the first `]`, clamped to `(rb+1)/3`;
    exact without strings, a hint with them.
  - `CountArrayElements` — strings, maps, `json.RawMessage`, structs that are cheap to
    skip, slices of leaf slices: skips elements with `SkipValue`, but after
    `countSampleCap` (64) extrapolates from the bytes the sample spans to the first
    `]`, clamped to `span/2`.
  - The clamps keep crafted input from buying more allocation than an honest document
    of the same size. Tighter would under-size honest documents; the
    `sizeof(element)` expansion is intrinsic (`bounds_test.go`).
- **`slicePresize` declines** struct elements that transitively reach any slice,
  array, map or interface field (`structSkipIsCheap`, cycle-safe through `seen` —
  counting would deep-scan every subtree the decoder then walks again), fixed-array
  elements (`[][3]float64` rings: counting descends every coordinate, and the memclr
  of a presized backing outweighs growth), and slices of slices of slices/arrays.
  Named struct elements resolve through `g.structTypes` like anonymous ones; `[]*Foo`
  is sized like `[]Foo`.
- **Growth without a count**:
  - The first append allocates `max(4, 256/sizeof(elem))` elements (a compile-time
    constant via `unsafe.Sizeof`); `[]` still yields nil.
  - `if len(*out) == cap(*out) { *out = unstable.GrowSlice(*out) }` grows a flat 2×.
    Plain `append` damps to ~1.25× above 256 elements, allocating ~5× and copying ~4×
    the final size; 4× buys no more time and wastes B/op (`TestGrowSlice`).
  - Named slice **roots** use `GrowSliceEst`, extrapolating from decode progress —
    `len * (end−start)/(i−start)`, padded by est/8+1, clamped to [2×, 8×] — from the
    root's `[` offset (`lightningArrStart`, no extra scan). The pad matters: a
    near-exact estimate plus the 2× floor overshoots (`TestGrowSliceEst`).
  - Nested slices of pointerful elements (`eltHasPointers`) use `GrowSliceSpan`:
    `arrayEndAt` finds the array's own `]` once, scanning forward from the cursor with
    the skip block loop at depth 1 (cached in `lightningArrEnd`), then applies
    `GrowSliceEst`'s estimate with a 64× ceiling. The scan costs ~0.9 instructions a
    byte against 1.5–7 per saved allocated byte, so `ScanWorth` (cost 24, emitted at
    the call site — `GrowSliceSpan` costs 231) requires ≥ 32 elements, ≥ 4 KiB of
    backing and JSON ≤ 5× that backing. Pointer-free slices keep flat 2×. Tests:
    `TestArrayEndAt`, `TestScanWorth`, `TestGrowSliceSpan`.
- **Elements decode in place** into the freshly grown zero slot (`&(*out)[last]`):
  decoding a composite element into a local and appending it heap-allocates per
  element. (For by-value leaves `append(v)` compiles to the same code.)
- **Reuse.** Decoding into a non-nil slice resets its length and appends into the
  existing backing, as encoding/json does. The generated reset is guarded —
  `if len(*out) != 0 { *out = (*out)[:0] }` — to avoid an unnecessary header store. Maps merge into existing entries (as in the stdlib), fixed
  arrays are zeroed then filled, slices reached as map values or through `lax`
  decode into a fresh scratch, and pointer fields reuse a non-nil pointee
  (`if dest == nil { dest = new(T) }`; null sets nil). Tests:
  `TestSliceReuseReplaces`, `TestSliceReuseKeepsBacking`,
  `TestDecodeFloat64SliceReplaces`, `TestPointerFieldReuse`.
- **`[]byte`** decodes from a base64 string or a numeric array (`DecodeByteSlice`),
  like encoding/json; `[N]byte` is numeric only. A failed base64 decode publishes
  `b[:n]`: the reused backing is already overwritten, so keeping the old length would
  splice two documents (`TestByteSliceStdlibParity`, `TestDecodeByteSlice`).
- **`//lightning:arena`**: `unstable.Arena[T]` aliases `github.com/JohanLindvall/arena`'s
  typed arena, built by `NewArena[T]` with `arenaChunkBytes` (4 KiB) chunks.
  `arenaCarve` reserves exclusive, exact regions (`len == cap`, so a caller's later
  `append` reallocates instead of clobbering a neighbor); backings over
  `arenaMaxCarve` (512 B) get their own `make`. The generator threads a per-root
  `<prefix>Arenas` struct as `a`, one field per element kind (`arenaField`; batched
  readers get `&a.<field>` via `arenaArgForExpr`). Never call `Reset` — freshly
  zeroed regions and no reuse are what the readers rely on. A surviving small slice
  pins its chunk, which is why this is a directive; schemas without it get no arena
  code. Chunking reduces allocation count and GC tracking; measure bytes separately. Tests: `TestArena*` in pkg/unstable, conformance `TestArenaDirective`
  (including arena × recursion, `ArenaTree`); micro `BenchmarkDecodeSmallSlices`.

## Whitespace

- **Every byte `<= 0x20` is whitespace, at both ends of every value, in every
  function** — `SkipWS`'s one-compare rule, which the decoder and `Valid` inherit.
  Anything that trims or bounds a token uses it too (`KindOf`, the walkers'
  `isNullToken`, `countKernel`'s blank-span test), or it answers for a document the
  rest of the library doesn't. Don't "fix" it. Tests:
  `TestKindOfWhitespaceIsThePackagesOwn`, `TestNullContainerWhitespaceIsThePackagesOwn`.
- `SkipWSRun` is an 8-byte SWAR loop: `data[i:i+8]` loads behind an unsigned entry
  test and `i <= len(data)-8` (no bounds check), and the mask
  `nws := ((w&^hi + 0x5f…) | w) & hi` uses constants that encode as AArch64
  immediates. The all-spaces shortcut (`w == 0x2020…20`) is per GOARCH
  (`skipWSSpaceShortcut`): on for arm64, off on amd64, where Meteor Lake pays for the
  extra taken branch. It must stay inlinable (74 amd64 / 78 arm64, budget 80): re-check
  `-gcflags=-m` under both GOARCHes after any edit, since the `g.skipWS` design
  depends on it. `TestSkipWSRunMatchesOracle` (exhaustive),
  `FuzzSkipWSRunMatchesOracle`.

## Dynamic `any`, `Valid`, depth bounds

- **`decodeAnyArray`** buffers up to 16 elements in a local `[16]any`, copies them
  into an exact-size backing at `]`, and switches to a capacity-32 backing at element
  17. The scratch stays separate from the returned slice (returning a slice of it
  would move it to the heap). `[]` returns the shared `emptyAnyArray` (non-nil, zero
  capacity) before the scratch is set up. Tests: `TestDecodeValueArrayBoundaries`,
  `TestDecodeValueArrayTruncation`, `TestDecodeValueEmptyArrayOwnership`.
- **`any.go` strings and numbers**: the number case calls `scanFloat` with the
  strconv fallback inline, mirroring `ReadFloat64OrNull`. String values scan inline
  and, on an escape, continue in `decodeStringEscaped` from the backslash already
  found; keys use the inline trick and fall back to `ReadKey`. Errors match the
  readers' (`TestDecodeValueStringMatchesReader`). Maps and boxing contribute
  additional allocation and CPU costs compared with typed decoding.
- **`DecodeValueNumber[Compact]`** (`json.DecodeAnyNumber[Compact]`) keeps numbers as
  `json.Number` — the literal copied and validated by the same `scanFloat` scan —
  threaded as a `number` flag through the object and array arms. pkg/unstable imports
  `encoding/json` only for this type (`TestDecodeAnyNumberMatchesUseNumber`).
- **Literals** compare as constant strings (`string(data[i:i+4]) == "null"` is one
  word compare) in `ExpectNull`, `ReadBoolOrNull` and `SkipValue`.
- **`Valid`** wraps `unstable.SkipValueStrict` (`pkg/unstable/valid.go`), the one
  strict single-value grammar walk (lax fields use it too).
  - It is a flat loop, not recursion: one bit per open container in a fixed
    `[MaxDepth/64+1]uint64`, goto-label states `scanValue`/`scanKey`/`scanAfter`
    (locals declared up front), and empty containers handled at the opening bracket
    so `scanKey` can reject `}` after a comma. Zero allocations.
  - Strings scan in its own frame; only an escape calls `strictStringEscaped`, which
    resumes at the backslash, reuses `unescByte`/`hexNibble` (ORing four entries —
    validity only; the bit-16 marker survives the OR) and skips the literal scan when
    already on `\` or `"`. Error offsets are contract: a bad or truncated `\u`
    reports at the `u`, while `readUnicodeEscape` reports at the first hex digit —
    don't merge them. Number arrays go to the `Valid` walks. Tests:
    `strictStringEscapedReference`, `TestStrictStringEscapedByteClasses`,
    `TestStrictStringEscapedBoundaries`, `FuzzStrictStringEscaped`,
    `TestSkipValueStrictStringErrorOffsets`, `TestSkipValueStrictRestoresContainerKind`.
  - Contract: `Valid(data)` ⇔ `DecodeAny` accepts `data` (`FuzzValidMatchesDecodeAny`).
    Numbers accept exactly what `ReadFloat64OrNull` accepts (`1e309` is invalid
    here); strings mirror the scanners' leniency (raw control bytes accepted, unpaired
    surrogates unchecked); whitespace is `SkipWS`'s. Divergences from
    `encoding/json.Valid` are pinned by `TestValidDivergesFromStdlib`.
  - It does **not** predict what a generated `UnmarshalJSON` accepts, in either
    direction: generated decoders are stricter where the schema is, looser where they
    skip (unknown fields), and looser where they read (`ReadInt64OrNull` stops after
    the digits, so `1.2.3` in an int field reads as 1).
  - It is slower than `SkipValue`'s balancer — fine for lax, which runs it only after
    a decode has failed; unknown-field skips stay on the lenient skips. Its end offset
    is tested against `SkipValue` on well-formed input: `Valid` reads only the error,
    so a wrong end would pass every `Valid` test while making a lax field resume
    mid-value.
- **Depth bounds.** `unstable.MaxDepth` is 10000, as in encoding/json. A stack
  overflow is fatal and `recover` can't catch it, so every recursive walker carries a
  depth: `decodeValue`↔`decodeAnyObject`/`decodeAnyArray` return `ErrMaxDepth`,
  `stripper.handle` ejects, and the scalar skips go through `skipObjectDepth`/
  `skipArrayDepth` (`TestSkipDepthBound`). The block scan is iterative.
- **Generated recursion** is bounded only where a cycle exists. `computeDepthThreading`
  builds the named-type reference graph over `allNamed()` (siblings included;
  `namedRefs`, which unlike `markReferenced` keeps self-edges) and threads
  `depth int` only through decoders of types that reach a cycle. A decoder that
  counts a level carries `depthGuard` at its top and passes `depth+1`: every struct
  decoder, plus each slice, array or map decoder whose element is a named slice or
  map type on a cycle (`elemCountsLevel`); other composite helpers pass `depth`
  through. That set bounds every cycle — a cycle runs through named types, and each
  step into one is decoded inside a struct decoder or inside a container decoder
  whose element is that type — including cycles with no struct on them at all
  (`type List []List`). The rule depends only on the element type, which also keys
  the decoder's memo entry, and leaves every schema without a named slice or map on
  a cycle byte-identical. Call `markDepthFn` before generating the body: like
  `g.memo[key]`, a recursive call emitted mid-body must spell the same signature.
  Tests: `TestRecursiveTypeDepthLimit`, `TestMutuallyRecursiveTypeDepthLimit`,
  `TestNamedContainerCycleDepthLimit` (each named slice/map shape exactly at the
  bound), `TestNonRecursiveTypesTakeNoDepthParam` (fails to compile if `Doc` gains a
  depth parameter).

## Runtime contracts

- `ReadNumberOrNull` accepts exactly what `ReadFloat64OrNull` accepts
  (`TestReadNumberAcceptSetMatchesFloat64`): `01` stays accepted because the float
  reader and `Valid` accept it — encoding/json rejects it, deliberately left as a
  divergence. `SkipNumber` *measures* a number token rather than validating it
  (`SkipValue([]byte("+"), 0)` is `(1, nil)`); `ParseFloat`/`DecodeValue` accept `+5`.
- `parseRFC3339` checks the day against the month's length; fractional seconds
  accumulate at most nine digits and scale with one `pow10nano[fd]` multiply;
  `daysFromCivilCached` uses a year-start table for 1970–2261. `ReadTimeOrNull`
  unescapes before parsing. `TestReadTimeAcceptsEscapedTimestamps` checks this
  contract and compares against the running toolchain's standard library. `time.Parse` copies its input into its
  errors, so passing aliased input is safe (`TestReadTimeErrorRetainsNoAlias`).
