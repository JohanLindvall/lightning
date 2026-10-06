# AGENTS.md

Guidance for working on lightning — a code generator that emits fast,
allocation-light `json.Unmarshaler` implementations, plus the SIMD scanning runtime
and JSON toolkit it is built on. README.md documents user-facing behavior
(directives, tag options, the deliberate differences from `encoding/json`); this
file covers how the code is built, how to measure it, and what must not break. Most
design rationale also sits in comments beside the code — read them before changing
a hot path.

## Layout

- `main.go` — the generator (`package main`): reads struct definitions (plus
  same-package sibling files), emits `*_unmarshal.go`. `field`/`sliceDecoder`/
  `mapDecoder` build decoders; `slicePresize` decides presizing. Per-type
  `//lightning:` directives are parsed by `hasDirective` (whitespace-insensitive:
  `//lightning:compact` ≡ `// lightning: compact`) into `compactTypes`/
  `nocopyTypes`/`destructiveTypes`/`arenaTypes`/`strictTypes`, which the per-root
  loop turns into `g.compact`/`g.nocopy`/`g.destructive`/`g.arena`/`g.strict`:
  - `compact` — elide inter-token whitespace skips (`g.skipWS`).
  - `nocopy` — a slice/map root aliases its keys/elements into the input.
  - `destructive` — the type's nocopy strings unescape in place, into the input
    buffer (implies nocopy).
  - `arena` — small scalar slices carve their backings from per-decode chunks.
  - `strict` — an unknown key fails with `*unstable.UnknownKeyError` instead of
    being skipped.
  - `root` — asks for a method and nothing else.

  `g.cmark`/`g.csuf` keep variants apart in memo keys and function names
  (`Compact`/`Destructive`/`Arena`/`Strict` suffixes); nocopy variants differ by
  the `nocopy` decoder param / `NoCopy` suffix.
- `generator_test.go` — generator tests; `TestGenerate` is the generate-then-compile
  table.
- `conformance/` — end-to-end tests against a generated decoder (`data_unmarshal.go`;
  every `*_unmarshal.go` is gitignored, and `make generate`/`make test` rebuild it).
- `internal/sveasm` — derives the `WORD`-encoded arm64 instructions from their
  comment mnemonics (see Conventions).
- `pkg/unstable` — the runtime the generated decoders call, plus primitives
  exported for `pkg/json`; nearly all performance work happens here.
  - Readers: `read.go` (`Read*`), `parseint.go`, `digitrun.go`,
    `digits_{amd64,other}.go` (`readerWordFold`), `numbyte_{table,cmp}.go`
    (`isNumberByte`), `numeric.go` (`scanFloat`, `scanFloatSlow`, Eisel-Lemire;
    `powers_table.go` is generated), `string.go` (unescaping, `Unwrap`), `escbuf.go`
    (pooled escaped-string chunk), `date.go`/`time.go`, `any.go` (dynamic
    `DecodeValue`), `valid.go` (`SkipValueStrict`).
  - Arrays: `batch.go` (batched scalar-array readers), `points.go`
    (`DecodeFloat64Points`), `count.go` + `count_{amd64,arm64}.s`/`count_other.go`
    (presize `countKernel`), `grow.go`, `arena.go`.
  - Number kernels: `intrun_*` (integer arrays), `floatrun_*` (decimal arrays, the
    ring points walk, `Valid`'s number walks).
  - Scanning: `skip.go`, `skipfast*` (SIMD container skip), `structural.go`
    (`structuralMask`), `scan.go` (`ValueScanner`, the resumable skip the stream
    reader uses), `simd_*` (string/structural/escape scanners), `ws_{amd64,other}.go`,
    `load_le.go`/`load_other.go` (`load64`/`load32`, unchecked on amd64/arm64);
    `unstable.go` holds the rest.
  - Per-arch `use…` dispatch flags live in `*_{amd64,arm64}.go`; `*_other.go`,
    `simd_scalar.go` and `skipfast_noasm.go` are the portable fallbacks.
- `pkg/json` — the public API, implemented on the exported primitives: `get.go`
  (`Get`/`GetMany`/`GetPaths`/`Lookup`/`ObjectEach`/`ArrayEach`/`ArrayEachIndex`,
  plus `*Compact`; `GetPaths` is `Get` for several nested paths in one
  prefix-sharing pass), `stream.go` (`Reader`: the walkers over an `io.Reader`
  through a bounded buffer), `set.go` (`Set`/`SetMany`/`SetPaths`),
  `strip_defaults.go`, `checked.go` (`…Checked` wrappers), `valid.go`, `escape.go`,
  `scalar.go` (`String`/`Bool`), `kind.go` (`KindOf`; the constants are
  `Kind`-prefixed because `String`/`Bool` are functions), `parseint.go`, and
  `json.go` (`DecodeAny`/`DecodeAnyNumber`/`UnescapeString`/`UnescapeStringCopy`/
  `ParseFloat`, wrappers over pkg/unstable internals).
- `bench/` — a separate module (keeps the competitor libraries out of the main
  module's deps), one directory per case. `bench/run_bench.sh` regenerates the
  decoders and benchmarks lightning against encoding/json, easyjson, sonic and
  others; `pkg_bench.sh` (repo root) runs the main-module microbenchmarks.
  `bench/large-json/input.json` (~8 MB GeoJSON) is gitignored and downloaded from
  `input.url`.

## Commands

- `make check` — `lint` + `test`. `make test` regenerates the conformance decoder,
  then runs `go test -cover ./...`.
- `make lint` — golangci-lint under `GOARCH=amd64` and `arm64` (`unused` sees only
  build-tag-selected files, so one arch misreports the other's code as dead).
- `make vet` — `go vet ./...`. asmdecl checks only the host GOARCH's assembly and
  isn't in `go test`'s vet subset, so also run `GOARCH=arm64 go vet ./...` from an
  amd64 host (and vice versa).
- `make fmt-check` (gate) / `make fix` (`gofmt -w .` + `go mod tidy`). Generated
  decoders must be gofmt-clean too.
- `make sveasm` / `make sveasm-check` — regenerate / verify the `WORD`-encoded
  arm64 instructions in `SVEASM_FILES` from their mnemonics. Needs an aarch64 GNU
  `as`: `$SVEAS`, `aarch64-linux-gnu-as` (binutils-aarch64-linux-gnu, so it runs on
  amd64), or native `as`. CI runs the check.
- `make bench-test` — `bench/get`'s tests, which `go test ./...` can't reach.
- `make bench-md` — what CI's manual Benchmark workflow (`.github/workflows/bench.yml`)
  runs to produce the committed tables: `pkg_bench.sh` → `bench/pkg_results_<arch>.md`,
  `bench/run_bench.sh` → `bench/results_<arch>.md`. `pkg_bench.sh` takes a
  benchmark-name filter as `$1` and honors `BENCHTIME`/`BENCHCOUNT`.
- Inline costs: `go build -gcflags=-m=2` — they differ per arch, so check both
  GOARCHes. Surviving bounds checks: `go build -gcflags=-d=ssa/check_bce ./pkg/...`.
- Validation over the corpus: from `bench/`,
  `go test ./valid -run='^$' -bench=BenchmarkValidCorpus`.
- **Before pushing assembly**: `GOTOOLCHAIN=go1.25.0 go build -a ./... &&
  GOTOOLCHAIN=go1.25.0 go vet ./...` (the `go.mod` floor; `-a` and `./...` both
  matter, or the build cache hides a failure), `go test -race ./...`, `make lint`,
  `make sveasm-check`, and every dispatch arm the change touches (see Testing rules).
- **End of session**: `go test ./...` and gofmt clean. Leave the committed benchmark
  tables to CI — a locally generated table silently changes the host the file
  describes.

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
  - re-measure, with counts, any rejection that rests on single-alignment wall time.
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
  trades by the corpus (~96% of corpus strings are ≤ 32 bytes), not by the benchmark
  shapes. A change whose sign depends on the document (key length, formatting, size)
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
  micro can't see GC pacing. Judge decoder work against the nocopy twin (a copying
  decode spends ~20% in malloc/GC), and size growth ideas against decoding into a
  reused target.
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
- **Where amd64 corpus time goes**: roughly a fifth each in the generated decoders,
  the string scanner (bounded by its ABI0 call), and allocation + GC marking + write
  barriers — the largest bucket still open to attack.

### Machines

Measurements come from four hosts. Allocation results transfer exactly between them;
time often doesn't. Where cores disagree, the choice is a per-GOARCH compile-time
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
  all-spaces shortcut. Instruction-identical scanner re-layouts can lose 5–8% in the
  back end, so test any cut in taken branches on cloudflare-compact first. A broad
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

## Testing rules

- **Stdlib differentials need a methodless twin.** A generated `UnmarshalJSON` makes
  `encoding/json.Unmarshal(doc, &v)` call lightning. Compare against
  `type FooStd Foo` (as `bench/` does with `benchmarkStd`), and never put a
  `//lightning:` directive on the twin — that makes it a root with a method.
  `TestStdlibTwinsAreReflectionOnly` rejects any Unmarshaler in a twin's type graph
  except `json.RawMessage` and `time.Time`; a twin holding another stdlib type with
  its own `UnmarshalJSON` needs the same exception. A twin still inherits methods
  promoted from *embedded* fields, which is the behavior the embedding tests measure.
- **Pin deliberate divergences** with tests that fail whether the divergence widens
  or silently closes (`TestValidDivergesFromStdlib`,
  `TestNullFieldsDivergeFromStdlib`, `TestEmbeddedUnmarshalerDivergesFromStdlib`,
  `TestSkipPathsDivergeOnMalformed`, `TestReadTimeAcceptsEscapedTimestamps`,
  `TestStringsPassInvalidUTF8Through`), and keep README's "Differences from
  `encoding/json`" in sync. Go 1.27's encoding/json is backed by json/v2, so the
  stdlib-side pins (`TestReadTimeAcceptsEscapedTimestamps`,
  `TestGenerate/invalid_json_tag_names`, `TestEmbeddedUnmarshalerDivergesFromStdlib`)
  compute the stdlib's answer at run time: a divergence that still exists must keep
  its shape, one that has closed must agree on the value. A new failure there means
  encoding/json changed.
- **A differential proves nothing about shapes it never generates.** Check
  non-vacuity (assert the corpus reaches the shape) and ask what an oracle never
  sees: duplicate keys, reused targets, whitespace inside kept members. A kernel that
  stops early passes any differential, so assert it consumed the whole input
  (`TestIntRunWindows`, `walkFloatRun`).
- **Oracles share nothing with the code under test** — no tables, no scanners
  (`strictStringEscapedReference`) — and the error sentinel and the end offset are
  both part of the contract.
- **Run every dispatch arm a change touches.**
  - Live dispatch runs only the host's widest body. Flip the `use…` flags to test the
    narrower arms against the scalar oracle (`TestIndexVariantsFlip`,
    `TestIndexArm64Lengths`, `TestSkipBlocksVariants`), but only to bodies the CPU
    has: iterate `floatRunBodies()`/`validRunBodies()` with `defer restoreKernels()`.
    Forcing a flag regardless of the CPU dies with SIGILL. A reference run must clear
    every flag the code reads (`floatRunOff()`/`validRunOff()`), or the kernel is
    compared with itself.
  - All arms run from one x86 box: Intel SDE `sde64 -skx` (AVX-512BW without VBMI)
    and `-icx`/`-spr` (VBMI); `qemu-x86_64 -cpu Haswell-v4 | Nehalem | qemu64`
    (AVX2, SSE4.2, SSE2 — TCG has no AVX-512); `qemu-aarch64 -cpu cortex-a72 |
    neoverse-n1 | neoverse-n2 | max,sve-default-vector-length=64` (NEON without and
    with DotProd, SVE2, 512-bit SVE2); `qemu-riscv64`/`qemu-s390x` for the pure-Go
    fallbacks, little- and big-endian. Build each package's test binary once
    (`go test -c`, with `GOARCH` for foreign arches) and run it from its package
    directory; the matrix takes ~30 minutes.
  - Assembly counts as verified only once it has run on its arch, natively or under
    qemu: asmdecl checks frame offsets only. qemu and SDE verify correctness, never
    speed.
- **Guard pages** (`TestAssemblyStaysInBounds`, `TestNumberKernelsStayInBounds`,
  Linux): every assembly body runs over buffers flush against a `PROT_NONE` page at
  either end, at every length to 300 — the only proof of masked and overlapping-tail
  bounds. Patterns must put the farthest load in a window's last lanes (a number at
  lane 58+), or a shrunken bound faults nowhere. Add cases for every new body or
  window bound, and sabotage-check by cutting a bound.
- **Fixed-window fast paths fail silently** (a missed backslash returns escapes
  verbatim): test every position at every length around each boundary, with the
  buffer's capacity ending at the input (`TestStringFindsEveryEscape`,
  `TestUnescapeFindsEveryEscape`, `TestUnescapeStringCopyFindsEveryEscape`).
- **Resumable scanners need every chunk split**, sizes around 64 and boundaries
  inside strings — a dropped carry still passes at sizes 1, 7 and whole
  (`TestValueScannerChunkingIsInvisible`).
- **Allocator tests retain each buffer and re-read it** — never compare addresses,
  freed addresses recur — and run under `-race` (`TestEscapeScratch*`). Zero-alloc
  assertions must hold under `-race` too: race instrumentation changes inlining, and
  a nil-start `append` stays on the stack only while the compiler can see it bounded,
  so make the promise structural with a fixed per-frame `[N]T` backing.
- **Order-dependent logic needs permutation tests** (`TestEntryTypesOrderIndependent`).
- **Generator changes.** The recurring bug class is "the generator exits 0 and emits
  code that doesn't compile"; `TestGenerate` (generate-then-compile) guards it — add
  new shapes there. `genCase`: `extra` adds sibling files (directories allowed, so a
  second package works); `wantMethods` pins the exact set of `UnmarshalJSON`
  receivers; a `wantNoWarn` substring must not appear in the case name (diagnostics
  carry the temp path, which carries the test name). Every generator change owes a
  dual-generator diff (parent vs tree) over conformance and every bench schema, run
  inside each case's directory (the sibling scan reads the package's other files):
  byte-identical except where intended, with identical diagnostics.
- **Prove a comment-only change** by comparing the normalized disassembly of a probe
  binary built before and after; raw binaries differ anyway (pclntab, DWARF lines).
- **Sabotage-check new tests**: break the guarded code and watch them fail. Copy the
  file aside first — `git checkout <file>` also discards your own uncommitted work.
- **An in-place mode is its own contract**: test in-place == fresh, byte for byte.
- **Before calling a divergence or regression yours**, rerun the input on the parent
  commit. Agent worktrees branch from `origin/main`, not local `main`. Build after
  merging parallel branches: two branches adding the same identifier merge cleanly
  and then fail to compile.

## Performance architecture (the load-bearing designs)

### Codegen patterns and bounds checks

- **Cursor tests are unsigned** — `uint(i) < uint(len(data))` — in generator
  templates and the runtime. The compiler can't prove a cursor returned by a reader
  non-negative, so a signed test leaves a check on every `data[i]`; the unsigned form
  proves both bounds at once. The `panicBounds` stub is cold, but its branch costs a
  dispatch slot on an issue-bound core, and in a small leaf one surviving check is a
  CALL that costs the whole frame (look for `morestack`). Prove an idiom with
  `check_bce` before relying on it:
  - one byte: `uint(i) < uint(len(x)) && x[i] == c` (survives a loop-carried
    cursor); two bytes need two probes (`uint(i)+1 < uint(n)` keeps both checks);
  - a word load in a loop: `if uint(i) > uint(len(data)) { return i }` plus
    `for i <= len(data)-8` (`i+8 <= len(data)` keeps a check per word, and
    `d := data[i:]; for len(d) >= 8` is check-free but pushes `SkipWSRun` over
    budget);
  - `uint(i+4) <= uint(len(raw))` and `for uint(i) < uint(len(b))` over `b[i]`, not
    `len(raw)-i >= 4` or `range b[i:]`.

  Keep a check that sits off a latency-bound chain (reslicing it away puts an op on
  the chain): `readUnicodeEscape`'s four, `EscapeStringInto`'s short-run
  `for i+8 <= n`, `ExpectNull`.
- **Pass offsets, not reslices.** `f(b[i:])` costs ~7 instructions and clobbers
  len/cap, so scanners take a start and return an absolute index
  (`indexCloseOrEscapeAt`/`IndexCloseOrEscapeAt`, `indexStructuralAt`). Load
  `data[i:i+8]`, never `data[i:]` (6 more instructions a word). The escape scanners
  take a resliced buffer on purpose: the per-run gate guarantees ≥ `minVectorRun`
  bytes per call.
- **One bound, then unchecked loads**: `load64`/`load32` are unchecked on
  amd64/arm64; the caller proves the window once (`scanFloat`:
  `uint(i)+48 <= uint(len(data))`).
- **`unsafeStr` reinterprets the slice header** (cost 3) because `unsafe.String`/
  `unsafe.Slice` add a negate, a compare and a panic branch per call. It has no
  empty-slice guard, so an empty nocopy string may point into the input.
- **AArch64 immediates.** `AND`/`ORR`/`EOR` fold only bitmask immediates (a rotated
  run of ones, replicated: `0x80…`, `0x7f…`, `0x30…`, `0xf0…`); any other wide
  constant costs `MOVZ` plus up to three `MOVK`s at every use, back edge included.
  `ADD`/`SUB`/`CMP` take 12-bit immediates. The compiler never treats a 32-bit splat
  as a bitmask immediate, so run 32-bit SWAR tests in 64-bit registers — and check
  the disassembly.
- **Return where you decide.** A flag or `(byte, bool)` result tested later costs a
  `CSET` and a branch even when inlined. Copy a short shared body rather than calling
  it (`ParseInt` ← `ParseUint`, `UnescapeString` ← `UnescapeStringScan`).
- **Keep data-dependent counts off the cursor's chain.** Object decoding serializes
  on the cursor (~60 cycles a member on amd64). A constant advance or a predicted
  branch lets the next loads issue early; a SWAR count feeding the next address adds
  ~12–14 cycles.
- **Inline watch list** (budget 80; re-check under both GOARCHes after editing):
  `DecodeValue` 78, `SkipWSRun` 74 amd64 / 78 arm64, `skipValueOrEnd` 77,
  `CountArrayScalars` and `SkipObject` 72.

### Numbers

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
    word, `parse8Digits`, advance by a constant 8. The `i+7` gate is what separates
    it from the rejected constant-advance `SkipNumber`: after a short number in an
    object that byte is rarely a digit. Tests: `TestIntReadersMatchByteLoop`,
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
  (~17 fewer instructions a token; wins on N2) and six compares on amd64 (the table
  is +9–13% on Zen 4). `SkipNumber` must stay inlinable
  (`TestIsNumberByteMatchesComparisons`, `TestSkipNumberSpan`).
- **Batched scalar-array readers** (`batch.go`: `DecodeFloat64Slice`,
  `DecodeIntSlice`/`DecodeUintSlice`, `DecodeFloat64Array`/`DecodeIntArray`/
  `DecodeUintArray`, `DecodeByteSlice`). `batchSliceFn`/`batchArrayFn` route every
  slice or fixed-array field whose element is a bare float64/int/uint kind here
  (float32/bool keep the generated loop).
  - Runs go to the SIMD kernels first (next section), then to a scalar loop that
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
    digit or `-` first (small slices lose ~6% without the gate), and `-gcflags=-m`
    must show the buffer staying on the stack.

### Native number kernels

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
  - Size a decline by its consequence, not its frequency: ~1 in 260 canada
    coordinates needs refinement, but a refusal at a ring's first point costs the
    whole ring (`TestFloatRunRefines`). The Eisel-Lemire shortcuts hold only for this
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
  unchanged (`TestValidNumberRunMatchesScalar`). The ring probe reads the byte after
  every nested `[`, a 10–13% cost on `ValidShapes/deep` that is accepted.
- **`countKernel`**: one pass finds `]` and counts separators before it, applying
  the clamp and the `<= 0x20` blank-span rule, so `CountArrayScalars` inlines (cost
  72). amd64 steps 32 bytes with `BZHI` and broadcasts `c` from the frame
  (`VPBROADCASTB c+32(FP)` — a GP→X `MOVQ` there is legacy SSE; see Conventions).
  arm64 accumulates per-lane byte counters flushed with `UADDLV` every 63 steps and
  builds exact masks only in the block holding the `]`.

### Strings

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
  `TestEscapedStringsSurviveChunkReuse`. What escape-heavy input still pays for is
  bytes (~2 MB per gsoc_2018 decode); the only levers are bigger chunks (traded away
  for the retention bound) or fewer bytes.
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

### Slices: presize, growth, reuse, arena

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
  `if len(*out) != 0 { *out = (*out)[:0] }` — because an unconditional header store
  costs ~1% on cloudflare. Maps merge into existing entries (as in the stdlib), fixed
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
  code. The win is allocation count and GC tracking (marine_ik −95% allocs/op), not
  bytes. Tests: `TestArena*` in pkg/unstable, conformance `TestArenaDirective`
  (including arena × recursion, `ArenaTree`); micro `BenchmarkDecodeSmallSlices`.

### SIMD scanning and skipping

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
  decoding, worth ~5% on cloudflare. Its wrapper is one unconditional asm call: SSE2
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
  Strings are absorbed into the scan. It is not the rejected two-stage design:
  skipping has no typed stage 2.
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

### Whitespace

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

### Dynamic `any`, `Valid`, depth bounds

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
  readers' (`TestDecodeValueStringMatchesReader`). The any path stays ~2.8× the typed
  path; maps and boxing are intrinsic.
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
  `skipArrayDepth` (`TestSkipDepthBound`). The block scan is iterative. Cost: one
  compare per `{`/`[`, measured flat.
- **Generated recursion** is bounded only where a cycle exists. `computeDepthThreading`
  builds the named-type reference graph over `allNamed()` (siblings included;
  `namedRefs`, which unlike `markReferenced` keeps self-edges) and threads
  `depth int` only through decoders of types that reach a cycle. Struct decoders hold
  the guard and pass `depth+1`; composite helpers pass `depth` through. Call
  `markDepthFn` before generating the body: like `g.memo[key]`, a recursive call
  emitted mid-body must spell the same signature. Tests:
  `TestRecursiveTypeDepthLimit`, `TestMutuallyRecursiveTypeDepthLimit`,
  `TestNonRecursiveTypesTakeNoDepthParam` (fails to compile if `Doc` gains a depth
  parameter). **Open bug:** named slice and map types can be fields, so a cycle can
  avoid every struct. `type Root struct{ L List }; type List []List` decodes
  20 000 levels with no error, and deeper input overflows the stack. The guard (and
  the comment beside `depthGuard` claiming every cycle passes through a struct) must
  extend to slice/map decoders on a cycle, with slice and map cycles added to the
  depth tests.

### pkg/json toolkit

- **The walkers dispatch `SkipValue`'s arms at the call site** (`objectEach`,
  `arrayEach`, `arrayEachIndex`, `getMany`, `objectField`, `walkPaths`, the `Reader`):
  `uint(c)-'-' <= 12` → `SkipNumber` (exactly `SkipValue`'s number bytes; widen before
  subtracting), `"` → `SkipString`, `{` → `SkipObject`, anything else → `SkipValue`.
  Array walkers test numbers first, object walkers strings first. On a short value the
  frame and switch cost as much as the scan. A missed byte is harmless; a wrongly
  claimed one is a bug (`TestArrayEachDispatchMatchesSkipValue`). The walker's own
  `uint(i) >= uint(len(data))` test proves `data[i]` in range. Callbacks are checked
  with `err == ErrStop` before `errors.Is`. `arrayEach`/`arrayEachIndex` are
  layout-sensitive — judge edits by instruction counts. `set.go`'s walkers
  deliberately keep `skipValueOrEnd`, which already inlines at every call site.
- **Scalar readers** (`KindOf`, `ParseInt`/`ParseUint`, `String`, `Bool`,
  `UnescapeString*`) are small leaves where one surviving bounds check costs a frame
  (see Codegen patterns). `KindOf` dispatches through a static `[256]` table
  (`kindOfByte`): one load costs every kind the same, where a `switch`'s compare tree
  gets re-laid by unrelated edits. A literal must be the whole remainder, with trailing
  whitespace measured by `SkipWS`.
- **`GetPaths`** keeps one `[]int` scratch for every level's active-path set (sized
  `len(paths)*(maxDepth+1)`, backed by a stack `[32]int` for small sets — the output
  aliases only `data`, so it never escapes). There is no early exit once all paths are
  captured: first occurrence wins, so later duplicates must still be scanned. Each
  path behaves as if requested alone: `walkPaths` returns `(end, err, fatal)`; a
  failed descent is re-skipped with the lenient `SkipValue`, the walk continues, and
  the first error is still returned. Duplicate parent keys stop at the first, like
  `Get`. Tests: `TestGetPathsIndependenceRandomized`,
  `TestGetPathsIndependentOfCoRequestedPaths`.
- **`Set`/`SetMany`/`SetPaths` are zero-alloc with a reused `out`.** Creation streams
  into `out` (`appendMember` writes multi-key nesting directly; `setSpan` reports a
  non-object intermediate through `member`/`nested`). `SetMany`'s found flags are a
  stack `[64]bool`; `SetPaths`' per-key `sub` set and `setObject`'s `recurse`/`create`
  use per-frame `[8]int` backings (`subbuf`, `idxbuf`), because under `-race` the
  nil-start appends reach the heap. The key test runs first, once: a non-leaf match
  descends without pre-skipping and takes its end from the recursion, so each on-path
  container is scanned once. Early exits: `SetMany` copies the rest verbatim once every
  key is found; `setObject` does so only at the root frame
  (`depth == 0 && nmatched == len(active)` — a nested frame owes its parent the `}`
  offset). Shared semantics: the first occurrence of a duplicate key is edited, a key
  requested twice keeps the first request, and a non-object root has its value
  replaced with the surrounding bytes kept. Benches: `BenchmarkSet`,
  `BenchmarkSetMany`, `BenchmarkSetPaths`, `BenchmarkSetManyEarlyExit`,
  `BenchmarkSetPathsEarlyExit`; tests `TestSetManyMatchesSet`, `TestSetTruncatedNoPanic`.
- **`StripDefaults`** (`stripper.handle`, recursive, never inlined;
  `BenchmarkStripDefaults`, `BenchmarkStripDefaultsCompact`):
  - In-place mode (`output == input[:0]`) holds only while no write lands on bytes the
    walk will read again. A container member's key, colon and recursion output are
    written speculatively and rewound when the value strips to nothing, so the keep
    decision is made first (from bytes already read), and a kept member's original
    span is snapshotted into `stripper.scratch` — only when the output really aliases
    the input (`unstable.SameBuffer`, conservative), which keeps a separate output
    buffer at 0 allocs. One scratch serves every level (the argument is next to the
    buffer). The snapshot's end comes from `SkipValue`, which can disagree with the
    walk on malformed input, hence the guarded fallback.
  - A kept container member is re-emitted by `emitKeptCompact(src, base)` +
    `compactValue`, so `RemoveWhitespace` holds, reading from whichever buffer holds
    the original.
  - Keep `emitField` unchanged and inlinable (parameterizing it by `src`/`base` costs
    88 > 80); new arms go out of line, and `keepKey`'s length pre-filter leads the
    container branch.
  - A keep-key member whose container empties rewinds to just past its separating
    comma (`postComma`); a whitespace-only object in array position must not eject.
  - Tests: `TestStripDefaultsInPlaceFuzz` (in-place ≡ fresh),
    `TestStripDefaultsInPlaceMatchesFresh`, `TestStripDefaultsSnapshotShortOfMember`,
    `TestStripDefaultsInPlaceFuzzMalformed`, `TestStripDefaultsKeepKeyContainer`, and
    the `compact(preserve) == remove` oracle `TestStripDefaultsWhitespaceModes`.
- **Checked wrappers** (`checked.go`): `Valid` on the arguments and the result,
  `ErrValueCount` for a short `rawVal`, and `ErrUnsafeKey` — keys are written raw
  between quotes, so a key like `x":1,"role` would inject a well-formed member that
  `Valid` accepts. `StripDefaultsChecked` treats a token-free result as empty
  (`unstable.SkipWS(res, 0) == len(res)`; `PreserveWhitespace` can leave a
  whitespace-only remainder) and normalizes it to an empty slice.
- `pkg/json` re-exports all twelve `Err*` sentinels, `ErrUnknownKey` included
  (`TestSentinelsMatchable`); descending into a non-object returns `ErrExpectObject`.

### Streaming reader (`stream.go`, `unstable.ValueScanner`)

- **Each value is scanned once.** A value already buffered settles with `SkipValue`
  (parity with the in-memory walkers); otherwise `valueMore` refills once and feeds
  `ValueScanner`. Re-running `SkipValue` after each refill is O(n²)
  (`TestStreamHugeElementIsLinear`). After a successful scan, only a number ending
  exactly at the end of the buffered bytes is still undecided (`errMoreInput`).
- **`skip`** probes once, capped at `skipProbe` (4 KiB) so its cost tracks the value
  rather than the buffer; it accepts only `end < lim` and never refills to make a value
  fit, so a sibling bigger than the buffer streams past.
- **`ValueScanner` modes.** `Reset` behaves byte for byte like the scalar skip (typed
  brackets, `ErrMaxDepth`) and is the documented default (`TestValueScannerErrors`,
  `TestValueScannerMatchesScalarSkip`). The Reader uses `ResetFast` — the block scan,
  ~14× fewer instructions a byte — which inherits skipfast.go's malformed-input
  divergences; that is acceptable only because it steps past or delimits values
  exactly where the in-memory walkers use `SkipValue`. `skipBlocks` always starts
  outside strings, so `containerFast` walks out of a string the chunk begins in
  (`stringBody`), carries `inString`/`escaped` itself, and cuts the chunk at its last
  full block.
- **Fast paths are written out at the call site** (the skip arms; the buffered case
  of `space`/`colon`/`afterElement`), and nothing is consumed before the fallback.
  Keys use `IndexCloseOrEscapeAt` + `UnsafeStr` (`objectEach`, `enter`); `objectEach`
  decodes the key after the value, from `hold+klen`, because compaction shifts every
  index equally.
- **Several paths per document, forward only** (`enter`): each call resolves its path
  from where the previous call left the cursor, tracked by `open` (keys of nested
  objects entered and not yet closed), `entered` (the root's `{` passed), `after`
  (cursor just past a member's value), `pending` (the closer of a container a callback
  left with `ErrStop`; `finish` skips the rest first) and `done` (a keyless call
  consumed the root). A new path shares the open prefix, closes deeper objects with
  `closeObject`, and scans forward; a miss that reaches an object's `}` goes through
  `closed`, so the next path resolves from the parent; a key behind the cursor is
  `ErrKeyNotFound`, because the buffer is a bounded window. Tests:
  `TestStreamContinuesToTheNextPath`, `TestStreamContinuationIsForwardOnly`,
  `TestStreamContinuesPastAStoppedWalk`.
- A callback's window is valid only until the callback returns. `WithMaxElement`
  bounds one element (`ErrElementTooLarge`); `WithBufferSize` is the caller's buffer
  lever (`defaultReaderBuffer`, 64 KiB).

### Escaping (`escape.go`)

- `IndexEscape` finds the next byte that needs escaping: `indexEscapeSSE2` on amd64
  (built like the string scanner — SSE2 first 32 bytes, then AVX2, then a 16-byte loop
  and scalar tail; control bytes via `PMINUB(v, 0x1f) == v`; inputs under 16 bytes
  skip the splat loads), `indexEscapeArm64` on arm64 (SVE2 `CMPLO #32`, otherwise
  NEON `VUMIN`), SWAR `indexEscapeScalar` elsewhere. `SwarNeedsEscape` is the one
  spelling of the predicate.
- `EscapeStringInto` picks SWAR or vector once per run: with fewer than `minVectorRun`
  (48) bytes left it walks SWAR words, otherwise it probes one word and hands the clean
  bulk to `IndexEscape`. Deciding per run keeps `indexEscape` inlinable; a per-word
  budget, a SWAR prescan inside `indexEscape` and an asm scalar peek are all slower.
- `EscapeString` (Builder) escapes the tail into a stack `[128]byte` (non-escaping);
  longer tails grow on the heap.
- **Ill-formed UTF-8 becomes U+FFFD when escaping**, as encoding/json does when
  marshaling. The walk uses a predicate widened by non-ASCII bytes
  (`SwarNeedsEscapeOrNonASCII`, `IndexEscapeNonASCII`) until the first non-ASCII
  byte, where one `utf8.Valid` decides the rest: valid → `escapeValidInto` (plain
  predicate), invalid → `escapeInvalidInto` (a DecodeRune walk writing raw U+FFFD per
  ill-formed byte).
  - The widening costs no extra ops: amd64 ORs the raw chunk into the match vector
    before `PMOVMSKB` (sign bits are the non-ASCII lanes), NEON adds `VUSHR $7` +
    `VORR`, SVE2 tests `SUB 0x20`/`CMPHS #0x60`, scalar tails test
    `int8(c) < 0x20`, and the SWAR form `((v-0x20·lo)|v)&hi` guarantees only its
    lowest set bit (callers use only `TrailingZeros64`).
  - In `EscapeString`, a remainder first reached at an *escape* byte hasn't been
    UTF-8-checked and must go through `EscapeStringInto`, not `escapeValidInto`.
  - Tests: `escapeReference` (all bytes, UTF-8 corners straddling the 8/16/32/48-byte
    boundaries, fuzz, both entry points), the `indexEscapeNonASCII` arms of
    `TestIndexFunctionsMatchScalar`/`TestIndexVariantsFlip`,
    `TestIndexEscapeNonASCIIScalarOracle`, `TestEscapeStringMatchesStdlibCoercion`.
- Escaping has **no in-place form**: output grows, so `out` must not overlap the input
  (unchecked, per the two-tier convention). `UnescapeStringInto` allows
  `out == in[:0]` because unescaping shrinks — check the size direction before
  borrowing a sibling's buffer convention.
- Decoding passes invalid UTF-8 through (README; `TestStringsPassInvalidUTF8Through`
  covers all seven string paths).

## The inline trick — let the generator write hot bodies inline

Generated decoders write the common case of their once-per-member reads inline and
call pkg/unstable only for the hard case (`g.skipWS`, `g.readKey` in `main.go`):

- **Whitespace**: `if uint(i) < uint(len(data)) && data[i] <= ' ' { i++; if
  uint(i) < uint(len(data)) && data[i] <= ' ' { i = unstable.SkipWSRun(data, i+1) } }`
  — the common 0–1 bytes cost one or two compares; only a run of ≥ 2 enters the SWAR
  loop.
- **Keys**: `ReadKey` (cost 198) never inlines, so the generator emits the no-escape
  path — `unstable.IndexCloseOrEscapeAt(data, ks)` (offset in, absolute index out)
  and `unstable.UnsafeStr(data[ks:ke])` — and calls `ReadKey` only for an escaped key
  or an error.
- **Unknown fields** (`skipUnknown`): `>= '['` → `SkipValue`, `'"'` → `SkipString`,
  anything else → `SkipNumber` — exactly `SkipValue`'s own partition, so the two agree
  on malformed input too (`TestUnknownFieldSkipMatchesSkipValue`). Testing `'['` first
  makes the container arm the fall-through. There is no `'{'` → `SkipObject` arm:
  `SkipObject` inlines (cost 72) into every struct decoder and the code growth cancels
  the saving.

Rules:
- **Only for once-per-loop reads.** Inlining a per-field read repeats the block per
  struct field; a wide decoder (string_unicode's 60 string fields) then outgrows the
  inliner's budget for large functions, `IndexCloseOrEscape` stops inlining, and
  i-cache pressure turns the win into a +9% loss.
- **It pays where the inline path skips work** (the no-escape alias, the whitespace
  fast path) **or bypasses a dispatch around short work** (`SkipValue`'s frame and
  comparison tree in front of a number, a short string or a one-block container —
  `skipUnknown`, the walkers' arms). It doesn't pay where it only removes a wrapper's
  frame around an inlined asm scan (`SkipString`): that flat time is the call's ABI0
  marshaling, which stays.
- **The block must inline directly into the big caller**: a helper wrapping it costs
  97 (whitespace) or 102 (close-quote scan) and stays a call.
- Hand-applied in `StripDefaults`' `handle` (its three `SkipString` sites and its
  whitespace skips), in `any.go`, in the key reads of the `get.go`/`set.go` walkers,
  and in the stream `Reader` (`objectEach`, `enter`). The key-read form wins on Zen 4
  and is neutral on M2; it stays. The whitespace block is *not* used in
  `get.go`/`set.go`: those walkers are skip-dominated and it cost 4–6% on pretty
  input.

## Generator rules

- **Directives.** A known directive where it can't act (a fixed-size array, a defined
  scalar, any other non-struct/slice/map type; bare `nocopy` on a struct root) warns,
  as does a directive not attached to a type declaration (a blank line detaches it).
  An unknown `//lightning:*` name is an error when attached to a type and a warning
  when detached. A directive on a type another root reaches gives it its own method
  too (marked in `emitted` after `entryTypes`); its copy inside the reaching root
  follows that root's directives.
- **Every memo key carries the variant** (`g.prefix + g.cmark()`, names through
  `g.decFn`/`g.csuf`) and every other body input: the nocopy/lax suffixes, the map key
  type, and the `root:` marker that keeps a named slice root's `GrowSliceEst` decoder
  apart from a field decoder of the same element type. Otherwise roots with different
  directives share a decoder — a plain root inheriting a destructive sibling's
  in-place unescape (`TestLaxDecoderIsolation`).
- **Names.** Reserve decoder names with `g.uniq`, never a bare `g.used[fn] = true`. A
  schema type named like a generated identifier — a decoder parameter or local, an
  import of the generated file, or a predeclared identifier the generated code uses on
  its own (`max`) — is a hard error (`reservedIdents`/`checkReservedNames`). **Adding
  a local to an emitted template means adding it to `reservedIdents`.** Predeclared
  names echoed only from the schema's own type text (`bool`, `uint16`, `any`) are
  excluded; the known gap is `type uint16 struct{…}` used as a field, which `isScalar`
  still treats as predeclared.
- **Imports** come from flags set where a qualified name is *emitted*
  (`noteQualifiers` in `typeStr`, `numberRead`, `sliceDecoder` →
  `g.needJSON`/`g.needTime`/`g.needUnsafe`), spelled with the schema's own qualifier —
  never from scanning the output, which contains JSON keys as string literals. A
  template that emits a new qualified name must set its flag. `isRaw`/`isNumber`/
  `isTime` match only through the schema's own qualifier.
- **Which types get `UnmarshalJSON`** (`entryTypes`): start from the types nothing
  references, mark what they reach, then promote a cycle nothing emitted enters — its
  whole strongly connected component, chosen among components that are sources of the
  still-uncovered subgraph (first in source order) — and repeat to a fixpoint. Both
  choices keep the result independent of declaration order
  (`TestEntryTypesOrderIndependent`). Generic and alias declarations warn and get no
  method; an alias to a struct literal still registers in `g.structTypes` (not
  `g.order`) so it decodes as a field type.
- **Defined types as roots.** A type defined over a struct, slice or map
  (`type ruleRaw Rule`; `underlying` resolves chains, in-file or in a sibling) becomes
  a root with that shape only when it carries a directive or `//lightning:root` —
  opt-in, because the same spelling is the methodless twin. That is the idiom for a
  hand-written `UnmarshalJSON` that decodes its own fields without recursing; declare
  the defined type at top level (one inside a function body is invisible to the
  generator).
- **Hand-written unmarshalers are delegated to.** `collectUnmarshalers` finds them in
  the input and its siblings (never in a `_unmarshal.go`); `field` checks for them
  before the struct lookup; `delegate` finds the value's span with `SkipValue` and
  passes it — null included, aliasing the input, as the stdlib does — reporting the
  value's start on failure. Such a type is never generated for (as a root it is
  dropped with a warning). Every foreign `pkg.Type` other than `time.Time`,
  `json.RawMessage` and `json.Number` is delegated the same way; a missing method is a
  compile error, which is the intended check. `nullAssigns` lets a delegated type
  decide what null means.
- **Siblings** (`registerSiblings`): same-package `.go` files, excluding `_test.go`,
  `_unmarshal.go` and `_`/`.`-prefixed files. Their struct, slice, map and scalar
  types register by name in `g.sibling`, never in `g.order`; the reaching root emits
  their decoders under its own directives (directives on sibling types aren't
  reported). A sibling whose import qualifiers disagree with the input's
  (`qualifiersAgree`) is skipped whole, with a warning.
- **Defined scalar types** (`type Severity string`) register in `g.scalarTypes`
  (`declaresScalar` resolves forward references; `scalarKind` follows the chain at
  use) and are read with the kind's own reader and stored through a conversion
  (`scalarAs`), so null and nocopy behave exactly as for the plain kind. No method;
  not counted by `isFlatScalarStringStruct`, which is only a presize hint.
- **A named slice or map used as a field** decodes through its element type's decoder
  with the destination converted (`callDecoderOn`, e.g. `(*[]Item)(&v.Items)`).
- **Map keys** (`mapKeyAssign`): a string, an integer kind, or a type defined over
  one. Integer keys parse with `unstable.ParseInt`/`ParseUint` from a non-escaping
  `[]byte(key)`; a non-numeric name is `ErrBadNumber`.
- **`//lightning:strict`** is part of `cmark`/`csuf` (`Strict` suffix), so a type
  reached from a strict and a loose root gets two decoders. `unknownKey()` emits
  either `skipUnknown` or `return i, &unstable.UnknownKeyError{Key: string([]byte(key))}`
  — the key copied because the error outlives the input; the position reported is the
  value's. Maps are unaffected.
- **Tag options**: `,number` only on an `any` field (`anyValueNumber`; elsewhere it
  warns — threading it into slices or maps of `any` would touch every element
  decoder); `,omitempty`/`,omitzero` pass silently; `,string` warns as unimplemented.
- **Types and embedding.** Only the empty interface and its spellings
  (`interface{ any }`) decode as `any` (`isAnyInterface`; `ast.InterfaceType.Methods`
  also holds embedded elements and type-set terms). An embedded `time.Time`,
  `json.RawMessage` or `json.Number` decodes as a field keyed by its type name, while
  Go promotes the embedded type's `UnmarshalJSON` (and, on Go 1.27, `json.Number`'s
  `UnmarshalJSONFrom`) to the outer struct, so encoding/json hands that type the whole
  document — a README divergence. Tag names that encoding/json ≤ 1.26 considers invalid
  (`invalidTagRune` is that `isValidTag` set, `|` included) are honored here and
  warned; Go 1.27 reserves only quotes, backslash and backtick and *ignores* such a
  field, and the warning names both. Don't "fix" this either way: rejecting breaks
  schemas that decode today, and adopting the stdlib's fallback silently moves which
  key an existing decoder answers to. A lax `[N]scalar` field reaches the batched
  array readers through a thin wrapper (`TestLaxFixedArrays`).
- **Field dispatch** (`keyDispatch`/`chunkedKeyEq`): the compiler inlines a
  string-vs-constant compare only up to 16 bytes and calls `runtime.memequal`
  (spilling registers) beyond. A struct with a key over 16 bytes dispatches on
  `switch len(key)` and compares long names in ≤ 16-byte chunks
  (`key[0:16] == "EdgeTimeToFirstB" && key[16:] == "yteMs"`); buckets whose names all
  fit keep a nested `switch key`, and structs without long keys use a plain
  `switch key`. Unmatched keys `goto lightningSkipKey` — one shared skip in its own
  block, so the `goto lightningKeyDone` over it crosses no declaration. A field whose
  pipe-separated alternates differ in length is emitted once per bucket
  (`TestLongKeyDispatch`).
- **Trailing commas are rejected.** The generated object/slice/array/map loops
  (`genStructBody`, `sliceDecoder`, `arrayDecoder`, `mapDecoder`), the batch loops and
  `decodeAnyObject` use a first-iteration flag — `for first := true; ; first = false`,
  checked in the loop-top closer case, where a closer on a non-first iteration means a
  trailing comma; `decodeAnyArray` returns `[]` before its loop and checks for `]`
  right after each comma. **Don't rotate the generated loops instead** (closer checked
  before the loop and after each value): cheaper on paper, +11% on cloudflare — the
  wide decoder is that layout-sensitive. Tests: `TestTrailingCommaRejected`,
  `TestBatchTrailingComma`, the `objErrs` arms in `any_test.go`.
- **`,lax`** (`laxField`) skips a value that failed to decode with
  `unstable.SkipValueStrict`, not `SkipValue`, so a balanced but invalid value
  (`[1,]`, `[1 2]`) is still an error and the outcome doesn't depend on the host's skip
  path; a `]`/`}` where a value must start is `ErrInvalidJSON`.
- **Nulls.** Every `*OrNull` reader returns the zero value and a nil error on null, and
  generated code assigns unconditionally, so an explicit null zeroes a leaf field
  (string, bool, number, `json.Number`, `time.Time`) where encoding/json leaves it
  alone — a documented divergence pinned by `TestNullFieldsDivergeFromStdlib` (via
  `zeroNullLeaves`). The parity fix is known and rejected on cost: guarding the
  assignment (`if data[i] != 'n' { dest = val }` in `nullGuard`, in bounds because the
  reader returned nil) adds ~20% instructions and a bounds check per leaf to a wide
  decoder, paid by every decode for behavior only seeded or reused targets can observe.
  `laxField` tracks `nullGuard` through `nullAssigns` — `json:"n"` and `json:"n,lax"`
  must never disagree — so restoring the guard means moving the lax leaf kinds to the
  guarded side in the same change. Composites: slice/map/pointer/`any` become nil,
  `json.RawMessage` takes the literal `null`, nested structs and `[N]T` are untouched.
  A null *document* nils a slice or map root (`nullReset`) and leaves a struct root
  alone.
- **`,unwrap`** closures guard an all-whitespace body and reject trailing content with
  `genUnmarshal`'s own `unstable.SkipWS(data, i) != len(data)` check
  (`ErrInvalidJSON`): the body is a whole document. Under `,lax` a wrong-type body is
  still tolerated, trailing content is not. Tests: `TestUnwrapWhitespaceBody`,
  `TestUnwrapRejectsTrailingContent`.
- **Generated code has no comments** except the
  `// Code generated by the lightning generator from …; DO NOT EDIT.` header;
  explanations belong in `main.go` next to the template. Templates are `fmt.Sprintf`
  strings, so a `%` in emitted text becomes `%!(MISSING)` (`go vet` catches it). Check:
  `grep -c '//' <generated file>` is 1.

## Runtime contracts

- `ReadNumberOrNull` accepts exactly what `ReadFloat64OrNull` accepts
  (`TestReadNumberAcceptSetMatchesFloat64`): `01` stays accepted because the float
  reader and `Valid` accept it — encoding/json rejects it, deliberately left as a
  divergence. `SkipNumber` *measures* a number token rather than validating it
  (`SkipValue([]byte("+"), 0)` is `(1, nil)`); `ParseFloat`/`DecodeValue` accept `+5`.
- `parseRFC3339` checks the day against the month's length; fractional seconds
  accumulate at most nine digits and scale with one `pow10nano[fd]` multiply;
  `daysFromCivilCached` uses a year-start table for 1970–2261. `ReadTimeOrNull`
  unescapes before parsing, where encoding/json through Go 1.26 parses the raw quoted
  bytes (Go 1.27 unescapes too) — pinned adaptively by
  `TestReadTimeAcceptsEscapedTimestamps`. `time.Parse` copies its input into its
  errors, so passing aliased input is safe (`TestReadTimeErrorRetainsNoAlias`).

## Conventions

- **The edit/transform API is two-tier.** `Set`/`SetMany`/`SetPaths`/`StripDefaults`
  return only a `[]byte` and are best effort — bracket balancers, not parsers, passing
  uninterpretable input through rather than failing. That keeps them zero-alloc; don't
  add error returns or validity checks to the hot walkers. Untrusted input goes through
  the `…Checked` wrappers. New edit operations take the same shape.
- **The `go` directive (1.25) is the assembler floor**; CI assembles with that
  toolchain under `GOTOOLCHAIN=local`. Go 1.25's arm64 assembler lacks mnemonics newer
  ones accept (`VCMHI`, `VCMHS`, vector `VMUL`, `VUMULL` and `VSHRN` arrive in 1.27;
  `VUDOT`, `VUADDLP` and all of SVE are absent everywhere). A missing mnemonic becomes
  a `WORD`.
- **`WORD`-encoded instructions are derived from their comments.** In `SVEASM_FILES`
  the trailing-comment mnemonic is the source of truth: write
  `WORD $0x00000000 // <mnemonic>` and run `make sveasm` — never hand-encode. A second
  `//` starts prose, which is preserved. Inside a `#define`, write `/* <mnemonic> */`
  before the continuation backslash (that's how `SHORTCONV`/`LONGCONV` share code).
  sveasm assembles under `.arch armv8.6-a+sve2`; an instruction outside it
  (`+sve2-bitperm`'s `BEXT`) needs that widened plus a runtime gate `x/sys/cpu`
  doesn't provide.
- **No legacy-SSE instruction while the upper halves of Y0–Y15/Z0–Z15 are dirty**
  (after writing one and before the next `VZEROUPPER`), and no `RET` or tail `JMP` in
  that state. Intel takes a microcode assist per occurrence — a single `MOVQ R11, X1` in
  `countKernel`'s prologue is worth +31% on update_center on Meteor Lake — and AMD
  takes none, so Zen 4 numbers can't catch it. The usual culprit is a GP→XMM `MOVQ`/`MOVL`,
  which Go encodes as legacy SSE: use `VMOVQ` or broadcast from memory.
  `TestNoSSEAfterAVX` (`avxmix_test.go`) checks every `*_amd64.s` in program order with
  macros expanded (registers 16–31 exempt); on Intel,
  `perf record -e cpu_core/assists.sse_avx_mix/u` over the pkg/unstable test binary
  must record nothing.
- **The committed benchmark tables come from CI** and are never regenerated locally
  (see Measuring performance).
- Bench `data.go` files declare a single top-level `Benchmark` with **anonymous**
  nested structs, so only `Benchmark` gets a generated method and
  `type benchmarkStd Benchmark` is a reflection-only baseline.
- Keep one authoritative copy of anything enumerated (a divergence list, a dispatch
  table) and point to it; copies across files rot.

## Open issues

- **Depth-bound gap** for cycles through named slice/map types (see Generated
  recursion).
- **A one-block container skip has a flat fixed cost** — `SkipValue`'s frame,
  `skipContainerFast`, `skipBlocks`' ABI0 call — the largest known unclaimed cost.
  Size attempts against `BenchmarkSkipSmall`/`BenchmarkSkipSmallAtEnd`.
- **ABI0 marshaling** (~11–13 instructions a scanner call, ~45 a run-kernel call) is the
  floor under every key and string read and limits the kernels on short arrays.
  `<ABIInternal>` is refused outside `package runtime` (Go 1.27.1), and
  `simd/archsimd` needs `GOEXPERIMENT=simd` (its arm64 half arrives in 1.27). Re-check
  when the Go version moves.
- Zen 4: port `LONGTAIL2`'s refinement and the second strike to the VBMI body, then A/B
  canada and large-json; the VBMI points walk still reloads the `]` from the frame (a
  store-forward on the point chain) where AVX2 uses `BTSQ`. `scanFloat`'s fast path has
  been A/B'd on Zen 4 only.
- A half-window `CLASSIFY` for arrays that close within 32 bytes (~2.5% of
  marine_ik). `DecodeIntArray` has no route into the integer kernel (no `[N]int` field
  in the corpus).
- SVE2 `NMATCH` over `SkipNumber`'s accept set (at most ~2% on skip-heavy schemas; it
  needs amd64 and scalar twins and turns an inlined loop into an asm call).
- A SIMD UTF-8 validator for escaping's single `utf8.Valid` pass; a `0x0909…` equality
  in `SkipWSRun` for tab-indented input.
- On compact input each inline whitespace probe jumps over its body — four taken
  branches a member. The compiler decides block placement, and no source shape changes
  it without a call in the body.

## Tried and rejected (don't re-attempt without a new idea)

**General**
- Churning fuzz-verified assembly to remove ops that hide under a port or issue
  bottleneck. Measure with counters first.

**Numbers**
- `digitRun` (the counted word fold) in the amd64 integer readers or in any batch loop:
  the cursor waits on its load → mask → count chain (citm +4% cycles at fewer
  instructions; mesh +1.1% in the batch loops, where guarded hybrids also hit an
  inliner size cliff). A four-digit SWAR step in the amd64 readers: its failing attempt
  on 1–3-digit ints costs more than it saves.
- SWAR `SkipNumber` (a counted run, a constant `i += 8`, or a flagged-lane walk):
  break-even is 8–9 digits, and in a comma-separated stream `data[i+7]` is the next
  token, so there is no free length signal. Removing the call around `SkipNumber` is
  what pays.
- An `isNumberByte` table on amd64 (+9–13% on Zen 4: the load lands on the loop's exit
  branch); a digits-first hybrid table (+3.4 instructions an element).
- Masking a table index to drop its bounds check: padding `pow10exact` to 32 entries
  for `&31` cost `numbers` +1.9% (72 bytes of rodata plus an alignment shift), and
  masking puts an `AND` on the value chain. Reading `scanFloat`'s sign from the wide
  load (canada +2.0%, mechanism unknown). Table multiplies in place of the float path's
  `CL` shifts (~8 cycles on the value chain).
- Routing Clinger's negative-exponent case through Eisel-Lemire: one FP divide beats
  EL's table load + 128-bit multiply.
- In `scanFloatSlow`'s loop: an 8-byte chunk in any arrangement (float-array up to
  +16%), a per-iteration count-and-fold, a one-byte `data[i+3]` guard. Count-and-shift
  wins only as `scanFloat`'s straight-line fast path. In that fast path: chained
  fraction-word loads (a 13-cycle chain) and funnel shifts (+20 instructions).
- SWAR for RFC 3339 fractional seconds: tied.
- Zeroing only the unfilled tail of a fixed array in the scalar loop: ≤ 1.7%
  attributed, below noise.
- A fused *Go* byte or SWAR presize scan in place of `bytes.IndexByte` + `bytes.Count`
  (float-array +6–14%): the fused pass pays only in assembly (`countKernel`).
- Kernel variants: the VBMI gather for short numbers (mesh +10%); choosing the body per
  call (a ring point is a call); a sign branch; exponents in the `Valid` walks (sound
  but rare); keying the AVX2 long gather to the delimiter; a `-` mask in the AVX2 flat
  walk (no free register); a whitespace look-ahead in the NEON points walk; a separate
  miss flag instead of the strike counter; 32- and 56-byte strides, a fall-through
  fold and a one-register `TBL` in the arm64 integer kernel. The NEON flat walk and
  integer kernel are at their issue floor.

**Slices and presize**
- Presizing slices whose elements nest a slice, array, map or `any`: counting costs as
  much as decoding, O(depth) times over (citm +155%); a bracket-only counter for
  coordinate rings still loses (large-json +14%).
- Ungated span scans (citm +23%); scanning from `[` instead of the cursor; a 16 KiB
  first-append hint (+385%); dropping the `GrowSliceSpan` scan on arm64 (random +3.5%).
- 4× slice growth.
- Presizing update_center's `map[string]struct` (a depth-aware count costs about the
  rehash it saves); a comma-count hint for flat-valued maps (≤ 0.3% of any case);
  `[]float32`/`[]bool` batch readers (no corpus fields — the `DecodeFloat64Slice`
  pattern ports mechanically if one appears).

**Strings and `any`**
- Carving escaped strings from 4 KiB chunks, through `unstable.Arena`, or by threading
  an arena through every `Read*`: flat at +13% B/op — the win tracks chunk bytes and GC
  work, not allocation count. Growing chunk sizes per document with the state in the
  pooled chunk (pools empty at every GC). Carving the one-shot `UnescapeString`/
  `…Scan`/`…Copy` from the chunk (+36% on a one-escape string: a pool round trip pays
  only at a decoder's allocation rate).
- Key interning or map presizing in `decodeAnyObject` (trades key allocations for a
  second hash; bimodal object sizes defeat a fixed hint); an 8-element `decodeAnyArray`
  scratch; caching the container kind in `SkipValueStrict` (deep nesting +25%).
- Removing `SkipString`'s frame from a clean-string skip, inline in the generated
  unknown-field skip or in `SkipValue`'s `"` arm: flat on M2, N2 and Meteor Lake.
- SWAR/uint64 key matching: the compiler already compiles ≤ 16-byte constant compares
  to word compares.

**SIMD**
- A Go SWAR probe in front of the string scanner, even armed per schema by key length:
  a miss adds ~20 instructions before the call it makes anyway (−15% on long-key
  documents), a bet on key length the schema can't make.
- Reshaping the amd64 string scanner's found path, or an AVX2 32-byte first block
  (cloudflare-compact +5…+8% at fewer instructions); SSSE3 `PSHUFB` or EVEX k-mask first
  blocks; AVX-512BW 64-byte tails (load-port-bound, and Zen 4 double-pumps zmm); AVX2
  64-byte `VPSHUFB` loops (lose at 40–120 B); AVX-512 `VPSHUFB` (12–18% slower).
- Pure-SSE2 `indexStructural` (~2× slower skips).
- arm64 NEON string-scanner block rewrites (mask folds, a `VUSHR`+`VUZP1` movemask, a
  deferred high-lane `VMOV`, an overlapping final block, 32-byte RODATA splats, a
  32-byte unroll): each regresses or trades cloudflare for long strings, on M2 and on
  issue-bound N2 alike. The width gap to amd64 is NEON's; SVE2 bodies close it on SVE2
  cores.
- Skip loop: breaking `BLOCKTAIL`'s escape-carry recurrence (+4% cycles — the mask is
  the bound); moving NEON bitmaps out through the stack with `LDP` (no
  store-forwarding, so the carried chain stalls); a 3-class bracket fold (30 ops vs
  26); SVE2 for the skip loop (no predicate→GP move); an 8-block SVE2 structural loop
  (unreliable lab, rare long scans). SVE2 `BEXT` is unmeasured (it needs
  `+sve2-bitperm` and a HWCAP2 gate).
- Array-probe variants: sending every array to the block scan (long number arrays +29%
  instructions), deciding through `indexStructuralAt`, a quote-only probe (it lifts the
  `MaxDepth` bound for `[scalar,{…}]`); a SWAR pre-walk in front of a one-block skip
  (wins only below ~24 bytes).
- Removing a bounds check whose only cost is its cold stub in a large function, or
  reslicing a load to drop a check that sits off a latency-bound chain.
- A two-stage structural index (simdjson-style): purely a whitespace play (pretty
  −11…−28%, compact +30%), because lightning's stage 2 is the typed parse, not a cheap
  tape copy.

**Whitespace**
- A standalone vectorized `SkipWS`, a `SkipWSRun` wider than 8 bytes per iteration, an
  SSE2 continuation for long runs, an inner `for w == sp` loop, SVE2 (it must stay
  inlinable), outlining it with `//go:noinline` (cloudflare +4.6% on compact input that
  never calls it). The 8-byte SWAR behind the inline two-compare guard is the design.
- Memoizing run lengths per call site: unsound — a two-byte check can match a space
  inside the next string token.

**Generator and harness**
- Rotating the container loops for trailing commas, and the null-assignment guard (both
  above; a hoisted null probe still can't prove its bounds).
- Fusing the member loop's whitespace probes with the structural test after them: a
  bet on formatting (wins compact, loses pretty), and guarding a probe with the
  expected byte loses wherever the gap isn't empty.
- An offset-taking `UnsafeStr` for the key read (~0.5%, for a new exported helper with
  an unchecked bound baked into generated code — bound `j` too if revisited).
- A generator flag emitting a second in-place method for the destructive benchmark: the
  per-case source copy is simpler.

**pkg/json toolkit**
- Writing `SkipValue`'s arms out in `set.go`'s walkers (SetPaths +2.25% instructions;
  `skipValueOrEnd` and `SkipObject` inline already).
- Porting `getPaths`' shared stack scratch into `set.go` (+2–4%).
- `if len(keys) != 0` around the walkers' key-descent loop (−0.3%); moving the
  `Reader`'s cursor fields into locals (~0.2%).
- A smaller or growing `Reader` buffer: large documents refill more (StreamSkipToKey
  +8.7%). `WithBufferSize` is the caller's lever.
