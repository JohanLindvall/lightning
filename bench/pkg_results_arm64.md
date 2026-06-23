# lightning main-module benchmarks

- generated 2026-06-23T10:10:13Z
- go version go1.26.4 linux/arm64

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.9 | 500.77 MB/s | 16 | 1 |
| sentence_clean | 45.0 | 978.13 MB/s | 48 | 1 |
| url_clean | 49.0 | 1060.25 MB/s | 64 | 1 |
| log_line_clean | 164.7 | 2040.35 MB/s | 352 | 1 |
| path_with_backslash | 189.8 | 194.94 MB/s | 184 | 4 |
| json_in_json | 224.0 | 187.52 MB/s | 216 | 4 |
| prose_with_quotes | 125.4 | 303.02 MB/s | 128 | 3 |
| control_bytes | 144.5 | 166.13 MB/s | 104 | 3 |
| mostly_clean_one_quote | 203.5 | 1498.92 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.0 | 1996.43 MB/s | 0 | 0 |
| sentence_clean | 16.1 | 2732.53 MB/s | 0 | 0 |
| url_clean | 18.0 | 2894.59 MB/s | 0 | 0 |
| log_line_clean | 90.7 | 3703.10 MB/s | 0 | 0 |
| path_with_backslash | 54.1 | 683.48 MB/s | 0 | 0 |
| json_in_json | 90.0 | 466.85 MB/s | 0 | 0 |
| prose_with_quotes | 32.7 | 1161.26 MB/s | 0 | 0 |
| control_bytes | 51.0 | 470.31 MB/s | 0 | 0 |
| mostly_clean_one_quote | 86.6 | 3522.99 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1015.0 | 1783.70 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1361.0 | 1330.70 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3832.79 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 8985.95 MB/s | 0 | 0 |
| url_clean | 4.9 | 10679.73 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31403.03 MB/s | 0 | 0 |
| path_escaped | 85.7 | 501.80 MB/s | 48 | 1 |
| json_in_json | 116.5 | 463.55 MB/s | 64 | 1 |
| prose_with_quotes | 71.9 | 570.17 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5642.40 MB/s | 0 | 0 |
| mostly_clean_one_escape | 119.1 | 2570.08 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3288.17 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8066.53 MB/s | 0 | 0 |
| url_clean | 5.5 | 9533.91 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29592.03 MB/s | 0 | 0 |
| path_escaped | 45.7 | 941.19 MB/s | 0 | 0 |
| json_in_json | 70.0 | 771.54 MB/s | 0 | 0 |
| prose_with_quotes | 33.4 | 1226.48 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5088.38 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.9 | 12262.71 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2301.0 | 1201.93 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2226.0 | 1242.51 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14545.0 | 1649.91 MB/s | 0 | 0 |
| stringObj/dispatch | 6754.0 | 3552.94 MB/s | 0 | 0 |
| numberObj/current | 6354.0 | 1605.05 MB/s | 0 | 0 |
| numberObj/dispatch | 2875.0 | 3546.95 MB/s | 0 | 0 |
| numberArr/current | 449.3 | 14693.18 MB/s | 0 | 0 |
| numberArr/dispatch | 451.5 | 14621.45 MB/s | 0 | 0 |
| nestedMixed/current | 20177.0 | 535.31 MB/s | 0 | 0 |
| nestedMixed/dispatch | 3442.0 | 3137.93 MB/s | 0 | 0 |
