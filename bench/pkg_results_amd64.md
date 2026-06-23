# lightning main-module benchmarks

- generated 2026-06-23T19:50:04Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 33.4 | 478.66 MB/s | 16 | 1 |
| sentence_clean | 43.7 | 1007.21 MB/s | 48 | 1 |
| url_clean | 42.5 | 1222.21 MB/s | 64 | 1 |
| log_line_clean | 100.2 | 3353.05 MB/s | 352 | 1 |
| path_with_backslash | 202.9 | 182.34 MB/s | 184 | 4 |
| json_in_json | 247.6 | 169.60 MB/s | 216 | 4 |
| prose_with_quotes | 145.2 | 261.69 MB/s | 128 | 3 |
| control_bytes | 171.1 | 140.25 MB/s | 104 | 3 |
| mostly_clean_one_quote | 132.5 | 2302.24 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.8 | 1249.97 MB/s | 0 | 0 |
| sentence_clean | 25.8 | 1708.79 MB/s | 0 | 0 |
| url_clean | 21.9 | 2374.84 MB/s | 0 | 0 |
| log_line_clean | 30.1 | 11150.55 MB/s | 0 | 0 |
| path_with_backslash | 73.7 | 501.96 MB/s | 0 | 0 |
| json_in_json | 110.7 | 379.30 MB/s | 0 | 0 |
| prose_with_quotes | 47.2 | 804.91 MB/s | 0 | 0 |
| control_bytes | 65.5 | 366.63 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.5 | 8589.86 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 915.2 | 1978.90 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1602.0 | 1130.23 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2840.75 MB/s | 0 | 0 |
| sentence_clean | 5.7 | 7772.71 MB/s | 0 | 0 |
| url_clean | 5.7 | 9153.80 MB/s | 0 | 0 |
| log_line_clean | 9.1 | 37045.03 MB/s | 0 | 0 |
| path_escaped | 82.0 | 524.35 MB/s | 48 | 1 |
| json_in_json | 109.7 | 492.19 MB/s | 64 | 1 |
| prose_with_quotes | 65.6 | 625.32 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4362.96 MB/s | 0 | 0 |
| mostly_clean_one_escape | 90.6 | 3377.01 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.3 | 2551.49 MB/s | 0 | 0 |
| sentence_clean | 6.2 | 7042.84 MB/s | 0 | 0 |
| url_clean | 6.2 | 8327.62 MB/s | 0 | 0 |
| log_line_clean | 9.7 | 34518.65 MB/s | 0 | 0 |
| path_escaped | 58.8 | 731.84 MB/s | 0 | 0 |
| json_in_json | 85.1 | 634.69 MB/s | 0 | 0 |
| prose_with_quotes | 43.8 | 935.89 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 4000.67 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12361.02 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2350.0 | 1177.23 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2295.0 | 1205.25 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10295.0 | 2331.12 MB/s | 0 | 0 |
| stringObj/dispatch | 4541.0 | 5284.93 MB/s | 0 | 0 |
| numberObj/current | 5393.0 | 1890.90 MB/s | 0 | 0 |
| numberObj/dispatch | 1916.0 | 5322.73 MB/s | 0 | 0 |
| numberArr/current | 236.6 | 27901.06 MB/s | 0 | 0 |
| numberArr/dispatch | 239.3 | 27581.28 MB/s | 0 | 0 |
| nestedMixed/current | 15679.0 | 688.87 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2869.0 | 3765.16 MB/s | 0 | 0 |
