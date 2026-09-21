# Escape validation and dynamic decoding — 2026-09-21

Baseline: `5325310`. Host: Intel Core Ultra 9 185H (Meteor Lake), Linux amd64,
Go 1.27.1. This pass follows the separate Neoverse N2 work recorded in
`performance_2026-09-21.md`.

## What the profiles showed

Fresh CPU profiles used regenerated decoders, pinned to P-core CPU 1 with
`GOMAXPROCS=1`. Each benchmark ran for three seconds, excluding profiling
results from the A/B timing samples.

| Workload | Main observed costs | Decision |
| --- | --- | --- |
| Generated citm_catalog | Integer reader 17.9%; whitespace 10.8%; string scanner 9.6% | Keep the existing architecture-specific integer path. The recorded word-fold experiments do not justify repeating it unchanged. |
| Generated synthea_fhir | String scanner 13.2%; allocation, zeroing and GC spread across the profile | Preserve the current ownership and presizing behavior. |
| Generated twitterescaped | Unescaping 30.1% cumulative; string scanner 13.7% | Keep the existing decoding tables and chunked string backing. |
| Generated marine_ik | Float scanner 39.9% cumulative; scalar-array counting 12.6% | Preserve the measured float tiers and array strategy. |
| ValidShapes/escapes | Escape checker 83.6% cumulative, including repeated scanner calls and hexadecimal classification | Remove work already avoided by the string decoder. |

The generated paths have extensive prior tuning. The clearest remaining
redundancy was in validation, which still used a switch, four scalar hex checks,
and a scanner call after every escape. Separately, dynamic string decoding
scanned the clean prefix twice when it encountered an escape.

## Changes

`strictStringEscaped` now uses the existing `unescByte` and `hexNibble` tables.
Validation only needs to know whether four bytes are hexadecimal: OR-ing their
table entries preserves the invalid marker without constructing a code point.
The checker inspects the next byte before scanning a literal run, so consecutive
escapes and a closing quote avoid a scanner call.

`decodeValue` now resumes directly in `decodeStringEscaped` at the first
backslash it already found. Going back through `ReadStringOrNull` repeated both
the opening-quote check and the entire clean-prefix scan. Truncated strings
retain their original boxed empty-string result, end offset, and error.

Both changes reuse existing decoding rules. They add no unsafe operations,
allocations, dependencies, directives, or generated-code changes. The public
validation acceptance set, including its intentional differences from strict
JSON, remains unchanged.

## Measurement method

- Built independent baseline and candidate test executables from identical
  benchmark and test sources. A Go build overlay supplied the two original
  runtime files to baseline builds; the working tree did not need stashing.
- Ran eight interleaved rounds, alternating A/B and B/A order, pinned to CPU 1
  with `GOMAXPROCS=1`, using 150 ms per benchmark.
- Compared medians with `benchstat`. Values below roughly 2% are noise; “flat”
  means no significant difference at `p < 0.05`.
- Measured both `-ldflags=-funcalign=64` and default function alignment for the
  validation corpus and focused shapes. Timed benchmarks ran sequentially,
  without tests or compilation running alongside them.
- Kept CI-generated architecture summary tables untouched. These measurements
  are specific to this host; they make no arm64 performance claim.

## Results

The following table is the eight-round, 64-byte-alignment run. Every validation
case remains at zero allocated bytes and zero allocations per operation.

| Validation workload | Baseline | Candidate | Change |
| --- | ---: | ---: | ---: |
| twitterescaped corpus | 451.4 µs | 199.6 µs | -55.8% |
| gsoc_2018 corpus | 442.8 µs | 428.5 µs | -3.2% |
| Existing ValidShapes/escapes | 2.808 µs | 1.385 µs | -50.7% |
| Dense single-byte escapes, value | 1.064 µs | 234.2 ns | -78.0% |
| Dense Unicode escapes, value | 681.5 ns | 244.1 ns | -64.2% |
| Surrogate escapes, value | 707.3 ns | 239.6 ns | -66.1% |
| Mixed text and escapes, value | 445.8 ns | 315.4 ns | -29.3% |

The remaining eleven validation corpus cases were statistically flat. Escaped
object keys show the same dense-escape improvements as string values. Clean
strings, number arrays and deep nesting were flat in the shape run.

Dynamic decoding of a string with a long clean prefix and one escape fell from
199.5 ns to 176.5 ns (-11.5%). This shape directly exercises the removed rescan;
it is not a claim that all dynamic decoding improves by that percentage.

All 26 generated and dynamic decoder comparisons across the same 13 corpus
documents were statistically flat. In particular, the validation improvement on
twitterescaped does not imply a corresponding generated-decoder improvement:
the generated string readers already avoid these redundant scans.

The default-alignment repeat confirms the main results: twitterescaped validation
-56.4%, gsoc_2018 validation -4.5%, the existing escape-validation shape -48.8%,
and dynamic long-prefix escaped strings -11.6%. The other eleven validation
corpus cases remain flat. Do not count the default-only 10% clean-string
validation improvement as a win: that path executes neither changed function,
and it was flat at 64-byte alignment.

The default-alignment record-array micro initially read +3.1% (`p=0.028`),
despite containing no escaped strings. A targeted repeat using eight 500 ms
samples was statistically flat (1.372 → 1.404 µs, `p=0.244`), and the 64-byte
build moved in the other direction (1.389 → 1.360 µs). Its instruction counts
were unchanged within 0.05% at both alignments. This does not support an added
runtime cost, but the initial observation is retained here rather than omitted.

### Hardware counters

`perf stat` measured user-mode instructions, cycles, branches and branch misses
on the pinned P-core. For each variant, subtracting an N-iteration run from a
3N-iteration run removes startup cost; divide by 2N for per-operation counts.
The numbers below average two such differences, with N=200,000 for the existing
escape shape, 400,000 for dense escape strings, and 500,000 for dynamic strings.

| Shape | Instructions/op, base → candidate | Cycles/op, base → candidate |
| --- | ---: | ---: |
| Existing ValidShapes/escapes | 59,042 → 38,362 (-35.0%) | 12,218 → 5,899 (-51.7%) |
| Dense Unicode validation | 14,875 → 7,052 (-52.6%) | 2,866 → 1,087 (-62.1%) |
| Dense single-byte validation | 14,714 → 7,445 (-49.4%) | 4,930 → 1,054 (-78.6%) |
| Dynamic long-prefix escaped string | 3,071 → 2,621 (-14.6%) | 894 → 782 (-12.5%) |

These reductions support the mechanism independently of wall-clock timing:
validation removes per-escape dispatch and scans, and dynamic decoding removes
an entire redundant prefix scan. Sparse validation remains effectively flat;
its extra byte tests add about 1.3% instructions in the counter run.

## Correctness and compatibility

The new scalar oracle uses explicit character classes and a byte-at-a-time
literal scan, independent of the optimized tables and SIMD scanner. Tests
compare the exact error sentinel and end offset, not just acceptance. Coverage
includes all 256 escape-selector bytes, every possible byte in each of the four
Unicode-escape positions, consecutive escapes, surrogate halves and pairs,
raw control/invalid UTF-8 bytes, scanner boundaries, every truncated prefix of
the boundary corpus, and 20,000 arbitrary byte bodies.

Dynamic string tests compare all four reader modes to `ReadStringOrNull`,
including malformed/truncated input, long prefixes, nonzero starting offsets,
the type and value returned on errors, and independence from the input buffer.
Existing generator, conformance, numeric and ownership tests remain in place.

Passed on the final runtime implementation:

- `go test -cover ./...`: generator, conformance and both runtime packages.
  Coverage was 94.4% for `pkg/json` and 93.6% for `pkg/unstable`.
- `go test -race ./pkg/...`.
- `go test -p 4 ./...` in the separate benchmark module.
- `go vet ./...` for amd64 and arm64, plus arm64 test-binary cross-compilation
  for both runtime packages. The arm64 binaries were not executed on this host.
- `golangci-lint run ./...` for amd64 and arm64, and a final full test run
  after addressing a lint suggestion in the test oracle.
- `FuzzStrictStringEscaped`: 30 seconds, 2,097,742 executions, no mismatch.
- `FuzzValidMatchesDecodeAny`: 30-second fuzz budget, 1,586,456 executions,
  no acceptance mismatch.
- Formatting of changed Go files and `git diff --check`.

The sandbox exposes a synthetic `/tmp/.git`, which made the unchanged generator
test probes fail Go's VCS-stamping step. Setting `TMPDIR` to an empty temporary
directory inside the workspace allowed the full tests to run normally; no
generator or test-harness workaround was committed. The checks here use Go
1.27.1; the minimum Go 1.25 toolchain was not installed on this host.

## Reproduction

The focused benchmarks require no downloaded corpus:

```sh
go test ./pkg/json -run='^$' \
  -bench='^Benchmark(ValidEscapedStrings|DecodeAnyEscapedStrings|ValidShapes)$' \
  -benchmem -count=8
```

For corpus validation, from `bench/`:

```sh
go test ./valid -run='^$' -bench=BenchmarkValidCorpus -benchmem -count=8
```

Missing downloaded fixtures are explicitly skipped. For an A/B comparison,
build both versions with identical benchmark sources, then alternate:

```sh
GOMAXPROCS=1 taskset -c 1 /tmp/base.test -test.run='^$' \
  -test.bench='^BenchmarkValidCorpus$' -test.benchtime=150ms -test.benchmem
GOMAXPROCS=1 taskset -c 1 /tmp/opt.test -test.run='^$' \
  -test.bench='^BenchmarkValidCorpus$' -test.benchtime=150ms -test.benchmem
```

Run the corpus validation binaries from `bench/valid`, and generated-decoder
binaries from their corresponding `bench/<case>` directory. Repeat with reversed
order on alternate rounds and compare the captured outputs with `benchstat`.
