# lightning main-module benchmarks

- generated 2026-06-25T21:19:03Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 24.5 | 654.06 MB/s | 16 | 1 |
| sentence_clean | 31.8 | 1384.67 MB/s | 48 | 1 |
| url_clean | 31.1 | 1671.32 MB/s | 64 | 1 |
| log_line_clean | 74.8 | 4493.27 MB/s | 352 | 1 |
| path_with_backslash | 175.4 | 210.90 MB/s | 184 | 4 |
| json_in_json | 176.4 | 238.07 MB/s | 216 | 4 |
| prose_with_quotes | 103.1 | 368.65 MB/s | 128 | 3 |
| control_bytes | 113.7 | 211.00 MB/s | 104 | 3 |
| mostly_clean_one_quote | 92.9 | 3282.30 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 9.6 | 1661.44 MB/s | 0 | 0 |
| sentence_clean | 18.9 | 2324.81 MB/s | 0 | 0 |
| url_clean | 16.2 | 3201.23 MB/s | 0 | 0 |
| log_line_clean | 21.9 | 15314.30 MB/s | 0 | 0 |
| path_with_backslash | 55.8 | 662.79 MB/s | 0 | 0 |
| json_in_json | 84.6 | 496.23 MB/s | 0 | 0 |
| prose_with_quotes | 34.6 | 1099.48 MB/s | 0 | 0 |
| control_bytes | 47.5 | 505.51 MB/s | 0 | 0 |
| mostly_clean_one_quote | 26.8 | 11397.84 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 641.1 | 2824.96 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1116.0 | 1622.09 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.9 | 4152.30 MB/s | 0 | 0 |
| sentence_clean | 4.6 | 9466.62 MB/s | 0 | 0 |
| url_clean | 4.7 | 11095.25 MB/s | 0 | 0 |
| log_line_clean | 7.7 | 43811.88 MB/s | 0 | 0 |
| path_escaped | 61.3 | 701.88 MB/s | 48 | 1 |
| json_in_json | 81.0 | 666.79 MB/s | 64 | 1 |
| prose_with_quotes | 49.0 | 837.60 MB/s | 48 | 1 |
| unicode_heavy | 4.1 | 7280.13 MB/s | 0 | 0 |
| mostly_clean_one_escape | 68.3 | 4480.64 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.1 | 3906.73 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 8906.97 MB/s | 0 | 0 |
| url_clean | 4.9 | 10581.60 MB/s | 0 | 0 |
| log_line_clean | 8.0 | 42163.78 MB/s | 0 | 0 |
| path_escaped | 45.1 | 954.26 MB/s | 0 | 0 |
| json_in_json | 63.2 | 854.54 MB/s | 0 | 0 |
| prose_with_quotes | 35.0 | 1171.46 MB/s | 0 | 0 |
| unicode_heavy | 4.5 | 6622.65 MB/s | 0 | 0 |
| mostly_clean_one_escape | 18.9 | 16164.47 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1707.0 | 1620.37 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1730.0 | 1598.57 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 8135.0 | 2950.04 MB/s | 0 | 0 |
| stringObj/dispatch | 3139.0 | 7645.93 MB/s | 0 | 0 |
| numberObj/current | 4358.0 | 2339.81 MB/s | 0 | 0 |
| numberObj/dispatch | 1325.0 | 7699.36 MB/s | 0 | 0 |
| numberArr/current | 173.0 | 38145.59 MB/s | 0 | 0 |
| numberArr/dispatch | 175.3 | 37647.90 MB/s | 0 | 0 |
| nestedMixed/current | 12644.0 | 854.22 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1947.0 | 5548.66 MB/s | 0 | 0 |
