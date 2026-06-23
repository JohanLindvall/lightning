# lightning main-module benchmarks

- generated 2026-06-23T19:49:58Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.4 | 526.48 MB/s | 16 | 1 |
| sentence_clean | 44.9 | 980.85 MB/s | 48 | 1 |
| url_clean | 40.5 | 1283.62 MB/s | 64 | 1 |
| log_line_clean | 110.8 | 3031.39 MB/s | 352 | 1 |
| path_with_backslash | 197.4 | 187.41 MB/s | 184 | 4 |
| json_in_json | 226.0 | 185.81 MB/s | 216 | 4 |
| prose_with_quotes | 124.5 | 305.17 MB/s | 128 | 3 |
| control_bytes | 144.2 | 166.38 MB/s | 104 | 3 |
| mostly_clean_one_quote | 153.4 | 1988.12 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1870.77 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2609.36 MB/s | 0 | 0 |
| url_clean | 22.5 | 2315.62 MB/s | 0 | 0 |
| log_line_clean | 43.7 | 7695.32 MB/s | 0 | 0 |
| path_with_backslash | 58.2 | 635.43 MB/s | 0 | 0 |
| json_in_json | 94.0 | 446.55 MB/s | 0 | 0 |
| prose_with_quotes | 34.6 | 1096.94 MB/s | 0 | 0 |
| control_bytes | 53.1 | 452.30 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.8 | 6663.82 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1024.0 | 1768.07 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1359.0 | 1332.30 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3767.68 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9004.99 MB/s | 0 | 0 |
| url_clean | 4.9 | 10649.03 MB/s | 0 | 0 |
| log_line_clean | 10.8 | 31161.11 MB/s | 0 | 0 |
| path_escaped | 84.5 | 508.85 MB/s | 48 | 1 |
| json_in_json | 114.7 | 470.73 MB/s | 64 | 1 |
| prose_with_quotes | 71.9 | 569.92 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5652.21 MB/s | 0 | 0 |
| mostly_clean_one_escape | 112.7 | 2714.17 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3288.39 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8065.50 MB/s | 0 | 0 |
| url_clean | 5.5 | 9528.07 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29583.40 MB/s | 0 | 0 |
| path_escaped | 46.0 | 935.18 MB/s | 0 | 0 |
| json_in_json | 71.0 | 761.14 MB/s | 0 | 0 |
| prose_with_quotes | 33.3 | 1232.46 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5087.62 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12350.00 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2296.0 | 1204.62 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2227.0 | 1242.30 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14381.0 | 1668.68 MB/s | 0 | 0 |
| stringObj/dispatch | 6769.0 | 3545.44 MB/s | 0 | 0 |
| numberObj/current | 6353.0 | 1605.26 MB/s | 0 | 0 |
| numberObj/dispatch | 2871.0 | 3551.54 MB/s | 0 | 0 |
| numberArr/current | 464.5 | 14211.78 MB/s | 0 | 0 |
| numberArr/dispatch | 466.9 | 14137.14 MB/s | 0 | 0 |
| nestedMixed/current | 19874.0 | 543.49 MB/s | 0 | 0 |
| nestedMixed/dispatch | 3450.0 | 3130.36 MB/s | 0 | 0 |
