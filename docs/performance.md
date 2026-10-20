# Performance and CPU support

[Home](../README.md) · [Documentation](README.md)

Use the benchmark corpus to explore workloads, then measure your own data.
Performance depends on the schema, input, allocation strategy, CPU, and toolchain.

## Benchmarks

The [`bench/`](../bench) module compares Lightning's generated decoders,
`DecodeAny`, destructive and arena modes, `encoding/json`, easyjson, sonic,
goccy/go-json, and json/v2. Its dependencies are separate from the main module.

| Architecture | Decoder comparisons | Toolkit and runtime benchmarks |
|---|---|---|
| amd64 | [Results](../bench/results_amd64.md) | [Results](../bench/pkg_results_amd64.md) |
| arm64 | [Results](../bench/results_arm64.md) | [Results](../bench/pkg_results_arm64.md) |

These tables are single-run CI snapshots. Compare same-run ratios and inspect the
CPU and Go version headers. A difference between snapshots can reflect a different
runner; it does not establish a speedup or regression.

### Run locally

For a focused runtime benchmark from the repository root:

```sh
go test -run='^$' -bench=BenchmarkReadIntShapes -benchmem -count=8 ./pkg/unstable
```

For the comparison suite:

```sh
BENCHTIME=3s BENCHCOUNT=8 bash bench/run_bench.sh
```

The suite generates decoders and its harness, downloads missing input documents,
and installs the easyjson generator if necessary. It needs network access and may
need a newer Go toolchain than the main module. It writes `bench/results.txt` and
`bench/results_<arch>.md`; keep local outputs out of pull requests. The committed
tables are produced by CI.

For a code change, follow the [measurement guide](development/benchmarking.md):
use instruction counts, interleave baseline and candidate runs, and check layout
sensitivity. Include both fresh and reused targets when allocations matter.

## SIMD scanning

Lightning uses native scanning and number kernels where the CPU supports them:

- **Strings:** find the next quote or escape without examining each byte in Go.
  The amd64 scanner starts in SSE2 and uses AVX2 for long runs when available;
  arm64 uses NEON or SVE2.
- **Structural bytes:** locate braces, brackets, and quotes when skipping values.
- **Whole containers:** classify blocks, mask string contents, and balance brackets
  to skip an object or array without materializing it.
- **Number arrays:** classify and fold runs of integers, decimals, or coordinate
  points. Unsupported number shapes resume in the scalar readers.

The [SIMD developer guide](development/simd.md) documents the kernels, dispatch
gates, return contracts, and tests. The [runtime guide](development/runtime.md)
explains scalar paths, number conversion, and allocation strategies.

### CPU requirements

The baseline string scanners use SSE2 on amd64 and NEON/Advanced SIMD on arm64.
Optional features are detected with `golang.org/x/sys/cpu`; no build tags are
needed to choose a faster body.

| Architecture | Optional features | Used for |
|---|---|---|
| amd64 | AVX2 | Long string runs and structural scanning |
| amd64 | AVX2 with PCLMULQDQ, BMI1, BMI2, and POPCNT | Whole-container skip |
| amd64 | AVX-512BW with the skip kernel's other requirements | Wider container skip |
| amd64 | AVX2 and BMI2 | Number-array kernels |
| amd64 | AVX-512 VBMI with each kernel's base features | Decimal, coordinate, and structural variants |
| arm64 | DotProd | Number-array kernels |
| arm64 | SVE2 | String, structural, and escape scanners |

A CPU without an optional feature uses the corresponding baseline or scalar
path. Other architectures use portable implementations. No optional ISA is
required to use the library.

The variants agree on well-formed input and are tested against independent
oracles. The fast skip is a bracket balancer, so its behavior on malformed input
can depend on the dispatch path. Use the [documented validation and checked
APIs](compatibility.md#limits-and-untrusted-input) where that distinction matters.
