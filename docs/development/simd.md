# SIMD and assembly

[Documentation](../README.md) · [Contributing](../../CONTRIBUTING.md)

Read the [assembly conventions](workflow.md#conventions) and [dispatch testing rules](testing.md#testing-rules) before editing a kernel.

## Native number kernels

Every kernel returns `(n, p, closed)`: values are bit-identical to the scalar path,
and `p` is always a point the scalar path can resume from. Each 64-byte window
becomes lane masks (not-digit, comma, not-whitespace by `<= 0x20`) walked with
`TZCNT`/`BLSR` on amd64 or bit-reversed masks and `CLZ` on arm64. Anything outside a
kernel's grammar stops it exactly at the scalar loop's top state, so every value and
error comes from scalar code.

| | amd64 (AVX2 + BMI2) | arm64 (DotProd) |
|---|---|---|
| integer arrays (`useIntRun`) | `parseIntRunAVX2` | `parseIntRunNEON` |
| decimal arrays (`useFloatRun` = `useFloatRunLong`) | `parseFloatRunAVX2`; VBMI body under `useFloatRunVBMI` | `parseFloatRunNEON` |
| rings (`DecodeFloat64Points`) | `parseFloatPointsAVX2` / `parseFloatPointsVBMI` | `parseFloatPointsNEON` |
| `Valid` walks (`useValidRun` = `useValidPoints`) | `validNumberRun`, `validPointsRun` (512 bodies under `useValid512`) | `validNumberRunNEON`, `validPointsRunNEON` |
| presize | `countKernel` (AVX2 step under `useCountAVX2`, also needs POPCNT) | `countKernel` |

Other arches turn every kernel off (`count_other.go` keeps `bytes.IndexByte` +
`bytes.Count`). `useFloatRunVBMI` needs AVX512F/BW/VL/VBMI; `useValid512` needs
AVX512F/BW. The arm64 validation walks need only NEON but share the DotProd gate, so
one flag switches every walk.

- **Body selection lives in the assembly** wherever the Go wrapper must stay one
  inlinable call: the asm reads the flag and branches, or tail-jumps to another TEXT
  symbol (`indexStructuralAVX2`, `skipBlocks`, `parseFloatRunAVX2`,
  `validNumberRun`, `countKernel`; a two-call Go dispatch costs 164). Entries called
  once per ring (`parseFloatPoints`, `validPointsRun`) select in Go. A body reached
  only by tail jump still needs a Go declaration for asmdecl.
- **Handoff** (`decodeFloat64Slice`, `decodeIntSlice`, `decodeUintSlice`,
  `DecodeFloat64Points`): `hold` gives the element a call stopped at to the scalar
  code. The strike counter `run` starts at 1, resets to 2 after a productive call and
  drops by 1 per unproductive one, so a hopeless array costs one call and a helpful
  kernel survives one refusal (refusals happen mid-array — leading zeros count toward
  the 19 digits). `DecodeFloat64Array` keeps a plain bool and returns straight from
  `parseFloatRunV` when the kernel fills and closes the array, skipping `clear(out)`.
  `intRunMinSlots` (minimum spare capacity before a call) is 4 on arm64 and 1 on
  amd64.
- **Integer kernel** (1–8 unsigned digits an element): windows step a fixed 48 bytes,
  so the next address never waits on the walk; a length-indexed `PSHUFB` (`irCtrl`)
  right-aligns each number. Traps: `SHRX` sets no flags; a comma after an array's `]`
  belongs to the enclosing container (`closedBefore`); on arm64, classify block k+1
  while walking block k (otherwise each block exposes a ~25-cycle chain), zero
  `UDOT`'s accumulator first, check capacity per block, and compute free space from
  the cursor (an empty `out` may have a nil base). Tests: `TestIntRunMatchesScalar`,
  `TestParseIntRunDirect`, `TestIntRunWindows`; benches `BenchmarkDecodeIntSliceRun`,
  `BenchmarkParseIntRunShapes`.
- **Decimal kernel** (≤ 19 digits, no exponent): ≤ 15 digits take a per-length
  gather, the folds, and one divide by an exact ±10^L2 (Clinger; the signed divisor
  keeps `-0`); 16–19 digits take `LONGFOLD`, then Clinger below 2^53 or Eisel-Lemire
  in assembly.
  - The AVX2 body (`LONGGATHER2`: one `VPSHUFB` over the 16 bytes at each end of an
    80-byte window, no cross-lane permute) and the NEON body (`LONGCONV`;
    `TBL`/`MUL`/`UDOT`/`MUL`/`UADDLP`/`UCVTF`/`FDIV`, at most four V0 ops a number)
    include `eiselLemire64`'s low-word refinement and decline only where it
    declines. The VBMI body (`VPERMI2B`, 96-byte window) declines every case that
    needs refinement.
  - Size a decline by its consequence, not just its frequency: a refusal at a ring's
    first point sends the whole ring to the scalar path (`TestFloatRunRefines`). The Eisel-Lemire shortcuts hold only for this
    domain — no range tests, and rounding carries into the exponent
    (`TestFloatRunRoundsIntoExponent`; test the 1-in-512 low-bits pattern). Tests
    check each hand-back against `kernelDeclines`, not a decline rate.
  - The kernel is latency-bound: keep loads and the sign off the cursor chain. A stop
    after a `-` backs up over it (`TestFloatRunStopsAtSign`).
  - The arm64 walk: 64-byte blocks at a 48-byte stride, the next block classified
    only when the walk can reach it, the cursor following commas rather than each
    number's end so conversions overlap. Traps in any mask walk: `LSL` by 64 shifts
    by 0; restart a straddling element at its first non-whitespace byte (a run of 61+
    whitespace bytes otherwise loops); a capacity bound must not count commas past
    the `]`.
- **Rings** (`DecodeFloat64Points`): `sliceDecoder` routes `[][N]float64` with a
  literal `N` here (`isFloatPoint`), except under `//lightning:arena`. It is the
  generated ring loop, element for element, plus a walk that writes points into
  spare capacity. Each point is validated before conversion (`[`, the window's first
  `]`, n−1 commas); a point counts only once its separator is found (searching
  across whitespace windows); a hand-back restores the saved count, or the point
  decodes twice. AVX2 puts the point's `]` into the comma mask (`BTSQ`), so every
  delimiter is the mask's lowest bit; NEON classifies each successor's window early.
  Tests: `TestFloat64PointsMatchesPerPoint`, conformance `TestFloat64PointsMatchStdlib`.
- **`Valid` walks** are dispatched at `SkipValueStrict`'s `[` by the first element's
  first byte (rings only below `MaxDepth`). They take only `-?digits(.digits)?` within
  one window and hand everything else back (backing up over a `-`), so acceptance is
  unchanged (`TestValidNumberRunMatchesScalar`). Include `ValidShapes/deep` when
  measuring changes to the ring probe, which reads the byte after every nested `[`.
- **`countKernel`**: one pass finds `]` and counts separators before it, applying
  the clamp and the `<= 0x20` blank-span rule, so `CountArrayScalars` inlines (cost
  72). amd64 steps 32 bytes with `BZHI` and broadcasts `c` from the frame
  (`VPBROADCASTB c+32(FP)` — a GP→X `MOVQ` there is legacy SSE; see [Conventions](workflow.md#conventions)).
  arm64 accumulates per-lane byte counters flushed with `UADDLV` every 63 steps and
  builds exact masks only in the block holding the `]`.

## SIMD scanning and skipping

- **Scanners by arch.**
  - amd64: `indexCloseOrEscape` is two compares a block. `indexStructural` has a
    compare-classified AVX2 body (`c|0x20` against `{`/`}`, plus `"`; 64 bytes a
    step, finishing on the buffer's last 32 bytes, shifted) and a VBMI body
    (`useStructural512`: one `VPERMB` through `structPerm` and one `VPCMPEQB` per 64
    bytes, aligned loads with a masked tail, only Z16–Z31 so no `VZEROUPPER`). It is
    costed in ops per byte (Zen 4 is dispatch-bound here).
  - arm64 NEON: simdjson's shuffle trick for `indexStructural` (`TBL` on the low and
    high nibble, two bits so cross combinations stay non-structural); splats built
    with `VMOVI`; the first 16-byte block peeled, since most strings end there.
  - arm64 SVE2 (`useSVE2 = cpu.ARM64.HasSVE2`; M-series and N1 keep NEON):
    `WHILELO` removes the tail, `MATCH` tests a class in one op, `BRKB` + `INCP`
    recover the position, and the code is vector-length agnostic (each match set must
    fill every 128-bit segment). N2 has one predicate pipe, so count predicate ops:
    the quote and structural bodies run unpredicated `PTRUE` blocks, then a
    four-vector loop, with `WHILELO` only for the last < VL bytes; the escape bodies
    take the `PTRUE` loop only. Predicate registers are safe in leaf asm (async
    preemption never stops inside assembly).
  - Tests: `TestIndexStructuralBodies`, `TestIndexVariantsFlip`,
    `TestIndexArm64Lengths`.
- **The amd64 string scanner is length-adaptive** (`indexQuoteOrBackslashSSE2`): the
  first 32 bytes are SSE2, so short strings never touch AVX2 state or pay
  `VZEROUPPER`; only a string with 32 clean bytes switches to an AVX2 tail. Don't make
  it pure AVX2. Change its short path only with a measured win on cloudflare-compact,
  the corpus's most layout-sensitive case.
- **`indexCloseOrEscape` must inline** (cost 62) — it is the hottest call in object
  decoding. Its wrapper is one unconditional asm call: SSE2
  is the amd64 baseline (AVX2 is switched inside the asm), and arm64 reads `useSVE2`
  inside the asm and tail-branches to NEON (a Go `if` costs 124 and un-inlines it).
  The asm handles every length itself.
- **`indexStructuralAt(b, i)`** prescans two SWAR words of `structuralMask` —
  `w|0x20` folds `[`/`]` onto `{`/`}`, so three has-byte tests are exact (a bit-cube of
  the four brackets would also match `Y _ y DEL`), and only the lowest flagged lane
  is valid — then passes the start offset to the asm body, which takes any remainder
  itself (a Go byte loop costs ~11 instructions a byte). `skipObjectDepth`/
  `skipArrayDepth` write the first word inline so its six 64-bit immediates hoist out
  of the loop. Tests: `TestStructuralMaskMatchesByteScan`,
  `TestIndexStructuralAtMatchesScalar`.
- **Container skip** (`skipfast.go`, `skipfast_{amd64,arm64}.s`; the sonic-rs
  `skip_container`/JSONSki technique). `SkipValue` is the dispatch. Objects always
  take the block scan, and so does an array whose first element is `{`, `[` or `"`.
  After a scalar first element, `structuralMask` over the next 16 bytes routes it:
  `]` returns the end, `"` goes to the block scan (`[timestamp,"value"]` takes one
  block), and anything else goes to `skipArray` — all-number arrays, and `{`/`[`
  deliberately, because that path is the `MaxDepth` bound for the shape
  (`TestSkipValueArrayProbeMatchesScalar`). The block scan streams 64-byte blocks:
  four bitmaps — `"`, `\`, and only this container's open/close bracket (`isArray`
  picks which) — then `findEscaped64` (simdjson's branchless odd-run detection) and a
  prefix XOR build the in-string mask, and brackets outside strings are balanced.
  Strings are absorbed into the scan; skipping requires no typed second stage.
  - The whole loop is assembly. On amd64, `skipBlocks` is the AVX2 body and
    tail-jumps to `skipBlocksAVX512` under `useSkipBlocks512`; the gate is
    `useSkipBlocks = useAVX2 && HasPCLMULQDQ && HasBMI1 && HasBMI2 && HasPOPCNT`; the
    prefix XOR is one `VPCLMULQDQ` by all-ones; the bit math is shared in the
    `BLOCKTAIL` macro. On arm64, `skipBlocks` is NEON and unconditional: there is no
    predicate→GP move before SVE2.1 `PMOV`, so it stays NEON on SVE2 cores. Its
    movemask is a weight-and-fold `ADDP` cascade (`CLASS2` shares one final `ADDP`
    between two classes), and it has no PMULL prefix XOR — the GP→SIMD→GP round trip
    costs more than six shifted-register `EOR`s.
  - The assembly loops are mask-bound (on Zen 4 the four class masks share one
    once-a-cycle resource, ~8 cycles a block; on N2 they are vector-issue-bound). The
    Go loop (`prefixXor64`, other arches) is latency-bound on the escaped → in-string
    carry. Every variant skips `findEscaped64` when a block has no backslash and no
    carry, skips the prefix XOR when it has no unescaped quote, and updates depth by
    popcount when the block can't reach depth 0.
  - **The final < 64 bytes are a block** (`skipBlocksTakesTail`). On a buffer of at
    least 64 bytes, AVX2 and NEON reread the last 64 bytes and shift the bitmaps so
    bit 0 is the cursor — a shift, not a mask, because shifted-in zeros are inert and
    the carried bits stay at bit 0; AVX-512 masks the load; no-asm builds do it in
    `skipContainerBlocks`. Inputs under 64 bytes take the byte walk, so both amd64
    bodies see the same inputs. Carried results are written only when `end < 0`.
    Tests: `testSkipTailSweep`, `TestSkipContainerBoundaries`.
  - The Go glue is minimal: `skipContainerFast` holds the `fastSkipAvail` check (so
    `SkipObject` stays one inlinable call, cost 72) and moves its continuation out of
    line (`skipContainerBlocks`), so nothing stays live across the asm call.
    `fastSkipAvail` is false on other arches, where a scalar `maskBlock` is slower
    than `indexStructural`.
  - **Gotcha — `maskBlock`'s result offsets.** Go 8-aligns the result block after the
    `isArray bool` argument, so the first result sits at +32(FP). On amd64 the stores
    are inside the `CLASS(s, off)` macro, and asmdecl doesn't check FP references
    hidden in a macro, so a wrong offset passes `go vet` (a probe without the macro
    *is* flagged, which misleads). The variant differential (`testSkipVariantCorpus`)
    is what locks those offsets.
  - **The fast and scalar paths agree only on well-formed input.** On malformed input
    they diverge three ways — an unbalanced bracket of the other kind, nesting past
    `MaxDepth` (the iterative block scan accepts any depth), and a stray backslash
    outside a string, which diverges in both directions and splits the fast path
    against itself by length alone (`{\"a}` is truncated at 63 bytes and accepted at
    64). So `SkipValue`'s verdict on malformed input is host-dependent. `skipfast.go`'s
    header is the one authoritative list. Pinned by
    `TestSkipPathsDivergeOnMalformed`, `TestSkipBackslashLengthCliff`,
    `TestSkipDepthDivergence`.
  - `TestSkipBlocksVariants` flips the dispatch flags and checks each variant against
    the scalar oracle over random documents plus `boundaryDocs()`, with truncation
    safety. Benches: `BenchmarkSkipBlocksVariant`, `BenchmarkSkipContainer`,
    `BenchmarkSkipSmall`, `BenchmarkSkipSmallAtEnd`, `BenchmarkSkipSmallScalar` (the
    same shapes through `skipObject`/`skipArray`, which win only on the tiniest
    containers).
- `SkipString` peels its first scan out of its loop (escapes continue in
  `skipStringEscaped`), sparing clean strings a header spill; it still doesn't inline
  (cost 167).
