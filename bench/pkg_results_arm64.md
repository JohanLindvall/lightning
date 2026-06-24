# lightning main-module benchmarks

- generated 2026-06-24T11:16:19Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.2 | 513.54 MB/s | 16 | 1 |
| sentence_clean | 46.2 | 953.17 MB/s | 48 | 1 |
| url_clean | 43.6 | 1193.28 MB/s | 64 | 1 |
| log_line_clean | 120.9 | 2779.13 MB/s | 352 | 1 |
| path_with_backslash | 211.9 | 174.60 MB/s | 184 | 4 |
| json_in_json | 234.0 | 179.48 MB/s | 216 | 4 |
| prose_with_quotes | 131.7 | 288.46 MB/s | 128 | 3 |
| control_bytes | 145.3 | 165.14 MB/s | 104 | 3 |
| mostly_clean_one_quote | 165.8 | 1839.81 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1870.23 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2608.58 MB/s | 0 | 0 |
| url_clean | 22.4 | 2318.46 MB/s | 0 | 0 |
| log_line_clean | 43.8 | 7670.19 MB/s | 0 | 0 |
| path_with_backslash | 58.1 | 636.24 MB/s | 0 | 0 |
| json_in_json | 94.0 | 447.03 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1095.77 MB/s | 0 | 0 |
| control_bytes | 52.4 | 457.75 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.5 | 6702.44 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1018.0 | 1778.96 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1349.0 | 1342.76 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3767.80 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9005.43 MB/s | 0 | 0 |
| url_clean | 4.9 | 10648.92 MB/s | 0 | 0 |
| log_line_clean | 10.8 | 31147.15 MB/s | 0 | 0 |
| path_escaped | 89.0 | 483.00 MB/s | 48 | 1 |
| json_in_json | 121.4 | 444.83 MB/s | 64 | 1 |
| prose_with_quotes | 75.0 | 546.64 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5652.18 MB/s | 0 | 0 |
| mostly_clean_one_escape | 124.6 | 2456.57 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3287.40 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8064.82 MB/s | 0 | 0 |
| url_clean | 5.5 | 9530.66 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29594.29 MB/s | 0 | 0 |
| path_escaped | 45.9 | 937.92 MB/s | 0 | 0 |
| json_in_json | 70.3 | 768.50 MB/s | 0 | 0 |
| prose_with_quotes | 33.5 | 1223.00 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5086.86 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.9 | 12282.24 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2297.0 | 1204.28 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2228.0 | 1241.45 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14358.0 | 1671.45 MB/s | 0 | 0 |
| stringObj/dispatch | 3873.0 | 6196.39 MB/s | 0 | 0 |
| numberObj/current | 6354.0 | 1605.00 MB/s | 0 | 0 |
| numberObj/dispatch | 1653.0 | 6169.50 MB/s | 0 | 0 |
| numberArr/current | 464.9 | 14197.79 MB/s | 0 | 0 |
| numberArr/dispatch | 466.9 | 14137.93 MB/s | 0 | 0 |
| nestedMixed/current | 19438.0 | 555.67 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2134.0 | 5062.27 MB/s | 0 | 0 |
