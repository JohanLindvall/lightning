# Performance analysis — 2026-09-21

Measured against `6e0bd5b` on a two-core Neoverse N2, Linux arm64, Go 1.27.1.
The changes improve dynamic array decoding and allocation-free validation.
The generated decoder templates and assembly are unchanged.

## Method

- Built separate baseline and candidate test executables, regenerating the
  benchmark decoders before building them.
- Pinned execution to CPU 1, with `GOMAXPROCS=2`. Alternated A/B and B/A over
  eight rounds, using 150 ms per benchmark and comparing with `benchstat`.
- Measured both `BenchmarkLightningDecodeAny` and `BenchmarkLightning` on 17
  corpus cases, plus the array and validation shape benchmarks in `pkg/json`.
- Repeated the toolkit comparison with both binaries linked using
  `-ldflags=-funcalign=64` to distinguish code-layout effects.
- Used CPU/allocation profiles and `perf stat` user-mode counters. The reported
  per-operation counters subtract runs at N and 3N iterations, dividing by 2N;
  validation used N=200,000 and two repetitions.

Times below are medians. "Flat" means the comparison was not significant at
`p < 0.05`; small significant changes below about 2% are not treated as wins.
The existing CI-generated architecture tables were not regenerated locally.
No amd64 runtime performance claim is made.

A second run on three large inputs used `GOMAXPROCS=1`, eight
interleaved rounds, and 500 ms samples. It confirmed canada **-27.99%**, large-json
**-15.41%**, and marine_ik **-7.76%** in time, with the same byte reductions.

## Dynamic arrays: allocate for the values actually present

The baseline's `decodeAnyArray` allocated a 16-element backing for every
nonempty array. A two-number coordinate therefore reserved 256 bytes for its
slice backing on this machine. On the large GeoJSON dynamic decode, the array
reader accounted for 54% of allocated bytes. Its initial `makeslice` calls
accounted for about 20% of sampled CPU time; later `growslice` calls accounted
for less than 1%.

The decoder now buffers up to 16 values in a local `[16]any`. When a short array
closes, it allocates exactly the required backing and copies the buffered values.
A longer array transfers them into a 32-element backing on element 17 and then
uses ordinary append growth. The local buffer never escapes. This also removes
one backing allocation from long arrays.

Empty arrays return a shared boxed, non-nil `[]any` with zero capacity. Its
header and backing cannot be modified through the returned value: appending
creates a separate backing. This removes the 24-byte interface-box allocation
from every empty array. The empty-array return precedes scratch initialization.

### Corpus: `DecodeAnyCompact`

| Input | Before | After | Time change | Allocated bytes change |
|---|---:|---:|---:|---:|
| canada | 14.117 ms | 9.940 ms | -29.59% | -64.76% |
| large-json | 43.794 ms | 37.478 ms | -14.42% | -38.72% |
| marine_ik | 20.755 ms | 18.393 ms | -11.38% | -24.01% |
| instruments | 1.129 ms | 0.942 ms | -16.57% | -4.73% |
| float-array | 1.403 µs | 1.177 µs | -16.11% | -19.75% |
| citm_catalog | 4.664 ms | 4.330 ms | -7.15% | -9.51% |
| synthea_fhir | 8.268 ms | 7.866 ms | -4.86% | -13.39% |
| mesh_pretty | 5.041 ms | 4.825 ms | -4.27% | -12.44% |
| random | 3.913 ms | 3.750 ms | -4.17% | -7.06% |
| golang_source | 10.442 ms | 10.086 ms | -3.41% | -6.16% |
| twitter_status | 2.722 ms | 2.668 ms | Flat | -4.26% |
| twitterescaped | 2.702 ms | 2.723 ms | Flat | -4.28% |
| time-array | 1.869 µs | 1.800 µs | Flat | -12.67% |
| cloudflare | 9.222 µs | 9.247 µs | Flat | Unchanged |
| cloudflare-nocopy | 8.985 µs | 9.258 µs | Flat | Unchanged |
| cloudflare-compact | 8.879 µs | 8.920 µs | Flat | Unchanged |
| gsoc_2018 | 3.566 ms | 3.567 ms | Flat | Unchanged |

The generated-decoder controls ranged from -1.75% to +2.82% in time, with
unchanged allocation counts. The +2.82% twitterescaped result was not
significant (`p=0.279`). These measurements establish no material improvement
or regression in generated decoding.

### Small inputs and the tradeoff

| Numeric array length | Before | After | Before B/op | After B/op |
|---:|---:|---:|---:|---:|
| 0 | 35.2 ns | 6.4 ns | 24 | 0 |
| 1 | 184.5 ns | 107.3 ns | 288 | 48 |
| 4 | 290.1 ns | 244.8 ns | 312 | 120 |
| 16 | 681.5 ns | 717.8 ns | 408 | 408 |
| 17 | 1,014.5 ns | 838.2 ns | 928 | 672 |

An exactly 16-element numeric array already fit the old hint perfectly. It
now pays the scratch initialization and copy without saving heap bytes:
**+5.33%, about 36 ns** in this microbenchmark. This is a retained tradeoff for
the measured corpus gains. Large numeric arrays were statistically flat.
Arrays of records improved 7–17% across all measured nonempty lengths because
their nested short arrays also benefit. A smaller eight-element scratch buffer
was evaluated; 16 elements saves an additional allocation on longer arrays and
improved the 17-element transition substantially.

Smaller returned capacities can cause a caller's later append to grow sooner.
Values remain independently owned, and the decoder adds 256 bytes of bounded
stack scratch per active array frame on 64-bit systems.

## Validation: keep clean string scans in the parser

`SkipValueStrict` now scans clean keys and string values directly, avoiding a
non-inlined string-reader call per token. Escaped strings resume validation at
the first backslash through `strictStringEscaped`; their clean prefix is scanned
once. Number parsing, whitespace acceptance, depth bounds, error identities,
and error offsets retain their existing behavior.

| Validation input | Before | After | Time change |
|---|---:|---:|---:|
| Flat record | 1,012.5 ns | 911.9 ns | -9.93% |
| Pretty record | 1,101 ns | 1,020 ns | -7.36% |
| Array of records | 2,381 ns | 2,227 ns | -6.45% |
| Existing mixed benchmark | 220.9 ns | 214.4 ns | -2.94% |

All validation cases still allocate zero bytes. Scalars, empty containers,
numbers, escape-dense strings, and deep arrays moved by about 1% or less.
The pure-string array showed **+2.11%** at default function alignment and
**-2.64%** at 64-byte alignment. Its timing is layout-sensitive, so it is not
counted as a throughput win. Record improvements held at both alignments.

Hardware counters on the final implementation at 64-byte alignment confirm
less work:

| Input | Instructions before → after | Cycles before → after |
|---|---:|---:|
| Flat record | 12,106 → 10,225 (-15.5%) | 3,434 → 3,119 (-9.2%) |
| Array of strings | 11,555 → 9,578 (-17.1%) | 3,000 → 2,886 (-3.8%) |

Caching the current container kind was also tried. It slowed deep nesting by
25% and reduced the record gains, so that experiment was removed.

## Correctness and compatibility

Passed:

- `go test -cover ./...` on Go 1.27.1, including generator and conformance tests.
- `GOTOOLCHAIN=go1.25.0 go test ./...` on the minimum supported toolchain.
- `go test -race ./pkg/unstable ./pkg/json`.
- `golangci-lint run ./...` and `go vet ./...` for both arm64 and amd64.
- Cross-compilation of both affected test packages for amd64; these binaries
  were not executed on this arm64 host.
- `go test -p 2 ./...` in the separate bench module.
- `FuzzValidMatchesDecodeAny`: 30 seconds, 1,158,994 executions, no mismatches.
- Formatting of all changed Go files and `git diff --check`.

New regression tests compare the array values against `encoding/json` in all
four dynamic modes, including `UseNumber`, whitespace, nested siblings, input
ownership, and lengths around both the scratch and heap growth boundaries.
They also check every truncated prefix around those boundaries, trailing
commas, independent empty-array appends, zero empty-array allocations, mixed
container nesting across bitset words, the depth limit, and exact string-error
offsets at scanner boundaries.

`make` was unavailable on the host, so its lint and coverage-test recipes were
run directly. No dependency or minimum-toolchain changes were needed.

## Reproduction

The new benchmarks can be run without the external corpus:

```sh
go test ./pkg/json -run='^$' \
  -bench='^Benchmark(DecodeAnyArrays|ValidShapes)$' -benchmem -count=8
```

For an A/B comparison, build test executables from the baseline and candidate
with identical benchmark sources, then alternate pinned runs. For example,
after building `/tmp/base.test` and `/tmp/opt.test`:

```sh
GOMAXPROCS=2 taskset -c 1 /tmp/base.test -test.run='^$' \
  -test.bench='^BenchmarkValidShapes$' -test.benchtime=150ms -test.benchmem
GOMAXPROCS=2 taskset -c 1 /tmp/opt.test -test.run='^$' \
  -test.bench='^BenchmarkValidShapes$' -test.benchtime=150ms -test.benchmem
```

Repeat in alternating order and compare the saved outputs with `benchstat`.
Corpus executables must run from their corresponding `bench/<case>` directory
because their harness reads `input.json` relative to the working directory.
