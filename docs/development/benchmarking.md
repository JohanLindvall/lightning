# Measuring performance

[Documentation](../README.md) · [Contributing](../../CONTRIBUTING.md)

Use reproducible measurements to support performance changes.

## Measuring performance — read before claiming a speedup

`run_bench.sh` runs each benchmark once; those numbers are noise-dominated and never
back a claim. Neither do the committed tables (see below).

- **Executed instructions per op decide; wall time confirms.** Build
  `go test -c -ldflags=-funcalign=64` binaries, run one benchmark under `perf stat`
  (`:u` events) at `-test.benchtime=Nx` and `3Nx`, and take `(c₃ₙ − cₙ)/2N`, which
  cancels start-up and warm-up. Instruction counts repeat to ~4 digits; cycles drift
  5–15%. Judge instructions first, then cycles and mispredicts, then confirm with
  time — fewer instructions can still lose when a change lengthens a dependency
  chain. `:u` counting needs `perf_event_paranoid` ≤ 2 (Ubuntu's default 4 silently
  blocks it). Without PMU access, compare the hot function's instruction count and
  call structure (inlining, frames, bounds-check sites) in the disassembly.
- **Interleaved A/B for time.** Rebuild the base every time (stash or worktree;
  regenerate decoders with the base generator), alternate the binaries in ABBA
  order, compare with `~/go/bin/benchstat`. Build inside `bench/` and run each binary
  from its case directory (it reads `input.json` relative to CWD). Treat < ~2% as
  noise; for lightning vs competitors use `-count=8+` and medians. For a small change
  to generated code, put both variants in one binary (rename the schema's top-level
  types, as `run_bench.sh` does for its twins).
- **Layout is a lottery.** Go aligns functions to 32 bytes on amd64 (16 on arm64), so
  a size change shifts everything after it. `-funcalign=64` on both sides removes
  that, but freezes each function's own placement at one draw, which can itself
  alias in the branch predictor; instruction-identical builds differ ±5–15% on
  micros, and some functions (`arrayEachIndex`, the amd64 string scanner's short
  path) swing tens of percent. So:
  - a wall-clock move in code whose instruction count didn't change is layout, and a
    case that never executes the changed code is the control;
  - when the default and 64-byte alignment disagree in sign, add a third alignment
    and the mispredict counter, and compare each binary's cycles across alignments
    rather than taking a majority vote (the unstable one may be the baseline);
  - treat walker deltas under ~5% as layout unless they hold at both alignments;
  - alignment NOPs inside a hot loop count as instructions — diff both builds'
    `objdump` (addresses and line numbers stripped) before trusting a loop's delta;
  - keep hot paths byte-identical when adding cold ones, and push new logic out of
    line;
  - re-measure with counts before drawing a conclusion from single-alignment wall time.
- **Check the machine before and after a run** (`uptime`,
  `ps aux --sort=-%cpu | head -3`, the absolute time of a known case). A stray
  process slows both arms alike, so deltas still look plausible; the tells are
  unchanged shapes moving and absolute times off. Medians and minimums disagreeing in
  sign is noise.
- **Size before building.** `pprof -peek` gives the share attributed to a call site;
  a chain's cumulative % is not saveable overhead, cycle arithmetic over-predicts,
  and isolated micros over-predict end-to-end wins ~2×. A call costs ~2 ns, so
  removing one pays only when the work it wraps is under ~10 ns, and widening a
  short scan (SWAR) loses on the short tokens that dominate. Weigh short-vs-long
  trades by the corpus's token-length distribution, not by the benchmark shapes.
  A change whose sign depends on the document (key length, formatting, size)
  is a bet, not an improvement: it needs a miss that costs nothing.
- **Profiles around assembly mislead.** Samples on the instruction after a load,
  vector compare or branch are usually retirement skid — check IPC and mispredicts
  before calling it a stall. A large asm symbol (`skipBlocks`) is not removable work;
  the removable glue is billed to neighbors (a non-inlinable dispatch, values spilled
  across the call, constants re-materialized per call). A Go function wrapping an
  inlined asm call carries the call's ABI0 marshaling in its own flat time, so
  removing its frame removes none of it.
- **Allocation.** Allocated bytes cost time, not allocation count: batching wins
  scale with the bytes and objects the collector paces. Size an idea with a
  throwaway probe and a chunk-size sweep; run `GOGC=off` beside a normal run to split
  allocator from GC-marking work (never `GOGC=off` alone); an isolated allocator
  micro can't see GC pacing. Judge decoder work against the nocopy twin to separate copying and allocation
  costs, and size growth ideas against decoding into a reused target.
- **Read the helper's disassembly, not just the caller's** (`-gcflags=-S` on a small
  probe package). The compiler won't find strength reductions such as the two-LEA
  `n *= 5; n = d + n<<1`.
- **Benchmarks are deliverables.** Parameterize by what the cost depends on (digit
  count, string length and escapes, dispatch arm — `newapi_bench_test.go`,
  `walk_bench_test.go`). Make micros take the caller's shapes: a long buffer with an
  early match for scanners, elements that do and don't straddle blocks for kernels
  (`BenchmarkParseIntRunShapes`). Pair padded documents with end-of-buffer ones, and
  commit setup-confounded rows beside clean ones (`stream` vs `stream_reused` — a
  fresh `NewReader` is mostly its 64 KiB buffer). Glue changes show in
  one-small-op-per-call benchmarks (`BenchmarkSkipSmall`, `BenchmarkSkipSmallAtEnd`);
  skip-path changes don't show in the case suite at all — use
  `BenchmarkSkipContainer`/`BenchmarkSkipBlocksVariant`.
- **Committed tables are a snapshot, never evidence**: single runs on unpinned
  GitHub runners. Compare two only when their `cpu:` headers match (the amd64 pool
  rotates EPYC SKUs and per-host speed varies up to 2×; arm64 headers read
  `cpu: unknown`, so nothing says whether SVE2 ran). The third-party rows (`Stdlib`,
  `Sonic`, `Easyjson`, `Goccy`, `JSONV2`, …) are the control — if they moved, the
  machine did. The Speedup column (same-run stdlib ÷ decoder) survives a host
  change; raw ns/op doesn't.
- **Harness rows.** Besides `BenchmarkLightning`, each case gets
  `BenchmarkLightningDecodeAny` (input minified with `json.Compact`, decoded by
  `json.DecodeAny`), `BenchmarkLightningDestructive` and `BenchmarkLightningArena`,
  generated from gitignored per-case source copies (`data_destructive.go`,
  `data_arena.go`) with every top-level type renamed (the generator parses the copy
  alone, so helpers too; an `awk` extractor handles `type (...)` blocks) and the
  directive prepended to the root. The destructive row restores a pristine input
  each iteration, which understates the win and perturbs the cache — an apparent
  regression on a byte-bound case (skip-heavy) is that effect. Cases without nocopy
  strings generate an identical destructive decoder.
### Machines

Use representative cores when evaluating a change. Allocation behavior and timing
must be checked separately; a timing result on one core does not establish a win
on another. Where cores disagree, the choice is a per-GOARCH compile-time
constant (`isNumberByte`, `readerWordFold`, `skipWSSpaceShortcut`) whose verdict must
hold on both cores of that arch. Most x86 hosts run the AVX2 bodies, so a win only in
a VBMI body misses them.

- **Zen 4** (Ryzen 7 8840HS; AVX-512 + VBMI): ALU-port- and latency-bound at once,
  so fewer instructions and shorter chains both pay. Counters: `ex_ret_brn_misp`
  (mispredicts), `ls_bad_status2.stli_other` (failed store forwarding — take
  `[2]uint64` table entries by pointer), `de_src_op_disp.decoder` (microcoded ops —
  never branch on the flags of `SHR r, CL`; use `SHRX` + `TEST`),
  `ex_no_retire.not_complete` (waits on a chain). `VPERMB zmm` issues every 2 cycles;
  an unaligned zmm load spans two lines; `VPCMPEQB zmm→k`, `KMOVQ k→r` and
  `VPMOVMSKB ymm` share one once-a-cycle resource. AMD doesn't penalize SSE/AVX
  mixing, so Zen 4 can't clear an amd64 body for Intel.
- **Meteor Lake** (Core Ultra 9 185H; AVX2, no AVX-512; a shared box — pin to P-cores
  0–11 and decide on instruction counts): object decoding is front-end bound (taken
  branches, DSB→legacy-decode switches), which is why amd64's `SkipWSRun` has no
  all-spaces shortcut. Test cuts in taken branches on cloudflare-compact and
  inspect instruction counts and branch behavior. A broad
  regression at unchanged instruction counts is an SSE/AVX assist: read
  `assists.sse_avx_mix`.
- **Apple M2** (NEON + DotProd): its out-of-order window hides short ALU and
  call-frame savings that Zen 4 exposes. macOS profiles are swamped by background
  `runtime.kevent`/`runtime.madvise` samples — use
  `pprof -ignore='kevent|madvise|pthread'`.
- **Neoverse N2** (Hyper-V guest; SVE2 + DotProd): issue-bound — instructions are
  cycles, a never-taken branch is a dispatch slot, and tables beat compare chains.
  Only these events count (others are accepted and read zero): `cpu_cycles
  inst_retired op_retired br_retired br_mis_pred_retired stall_frontend
  stall_backend inst_spec op_spec br_pred br_mis_pred l1d_cache l1d_cache_refill`.
  Vector `MUL`, `UCVTF` and `FDIV` share pipe V0 (one per cycle in total), while
  `UDOT`, `UADDLP`, one- and two-register `TBL`, `FMUL`, the compares and `ADDP`
  issue on either pipe — stacking V0 ops runs ~2× slower than the count suggests.
  SVE2 `MATCH` uses a single pipe (the structural loop's floor is 4 cycles per 64 B).
  N2 runs the SVE2 scanner bodies; M2 runs the NEON ones.
- Harness traps: a `pkill -f` pattern that matches the harness's own shell kills it,
  and bash reads a running script incrementally.
