# lightning main-module benchmarks

- generated 2026-06-23T18:15:54Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 35.7 | 448.29 MB/s | 16 | 1 |
| sentence_clean | 51.3 | 857.73 MB/s | 48 | 1 |
| url_clean | 53.7 | 967.72 MB/s | 64 | 1 |
| log_line_clean | 190.3 | 1765.21 MB/s | 352 | 1 |
| path_with_backslash | 201.4 | 183.73 MB/s | 184 | 4 |
| json_in_json | 246.0 | 170.77 MB/s | 216 | 4 |
| prose_with_quotes | 143.9 | 264.10 MB/s | 128 | 3 |
| control_bytes | 164.5 | 145.90 MB/s | 104 | 3 |
| mostly_clean_one_quote | 212.1 | 1438.29 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 11.9 | 1347.80 MB/s | 0 | 0 |
| sentence_clean | 22.5 | 1956.17 MB/s | 0 | 0 |
| url_clean | 24.7 | 2103.47 MB/s | 0 | 0 |
| log_line_clean | 124.6 | 2697.67 MB/s | 0 | 0 |
| path_with_backslash | 71.0 | 521.04 MB/s | 0 | 0 |
| json_in_json | 104.4 | 402.24 MB/s | 0 | 0 |
| prose_with_quotes | 43.9 | 866.01 MB/s | 0 | 0 |
| control_bytes | 65.8 | 364.98 MB/s | 0 | 0 |
| mostly_clean_one_quote | 111.4 | 2737.59 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 894.6 | 2024.31 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1590.0 | 1139.06 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2844.03 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7804.86 MB/s | 0 | 0 |
| url_clean | 5.6 | 9223.83 MB/s | 0 | 0 |
| log_line_clean | 9.0 | 37140.24 MB/s | 0 | 0 |
| path_escaped | 82.9 | 518.66 MB/s | 48 | 1 |
| json_in_json | 109.1 | 494.91 MB/s | 64 | 1 |
| prose_with_quotes | 66.0 | 621.42 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4337.29 MB/s | 0 | 0 |
| mostly_clean_one_escape | 88.7 | 3448.48 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.2 | 2561.08 MB/s | 0 | 0 |
| sentence_clean | 6.2 | 7040.68 MB/s | 0 | 0 |
| url_clean | 6.2 | 8323.59 MB/s | 0 | 0 |
| log_line_clean | 9.7 | 34548.64 MB/s | 0 | 0 |
| path_escaped | 58.6 | 734.22 MB/s | 0 | 0 |
| json_in_json | 85.1 | 634.59 MB/s | 0 | 0 |
| prose_with_quotes | 43.7 | 937.37 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 4004.20 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12370.13 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2365.0 | 1169.38 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2301.0 | 1202.27 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10180.0 | 2357.33 MB/s | 0 | 0 |
| stringObj/dispatch | 4735.0 | 5068.74 MB/s | 0 | 0 |
| numberObj/current | 5391.0 | 1891.53 MB/s | 0 | 0 |
| numberObj/dispatch | 2006.0 | 5082.89 MB/s | 0 | 0 |
| numberArr/current | 207.4 | 31824.02 MB/s | 0 | 0 |
| numberArr/dispatch | 207.8 | 31768.55 MB/s | 0 | 0 |
| nestedMixed/current | 15673.0 | 689.14 MB/s | 0 | 0 |
| nestedMixed/dispatch | 3018.0 | 3578.95 MB/s | 0 | 0 |
