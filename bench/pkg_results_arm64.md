# lightning main-module benchmarks

- generated 2026-06-23T18:15:50Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.9 | 501.05 MB/s | 16 | 1 |
| sentence_clean | 44.5 | 988.93 MB/s | 48 | 1 |
| url_clean | 48.8 | 1066.74 MB/s | 64 | 1 |
| log_line_clean | 160.3 | 2095.79 MB/s | 352 | 1 |
| path_with_backslash | 188.4 | 196.36 MB/s | 184 | 4 |
| json_in_json | 224.5 | 187.08 MB/s | 216 | 4 |
| prose_with_quotes | 124.9 | 304.17 MB/s | 128 | 3 |
| control_bytes | 143.5 | 167.27 MB/s | 104 | 3 |
| mostly_clean_one_quote | 201.4 | 1514.05 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.0 | 1996.14 MB/s | 0 | 0 |
| sentence_clean | 16.1 | 2730.19 MB/s | 0 | 0 |
| url_clean | 18.0 | 2893.76 MB/s | 0 | 0 |
| log_line_clean | 90.8 | 3700.75 MB/s | 0 | 0 |
| path_with_backslash | 54.1 | 684.10 MB/s | 0 | 0 |
| json_in_json | 89.8 | 467.44 MB/s | 0 | 0 |
| prose_with_quotes | 32.9 | 1155.69 MB/s | 0 | 0 |
| control_bytes | 50.9 | 471.84 MB/s | 0 | 0 |
| mostly_clean_one_quote | 86.6 | 3522.75 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1017.0 | 1780.75 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1347.0 | 1344.55 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3833.15 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 8986.17 MB/s | 0 | 0 |
| url_clean | 4.9 | 10696.41 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31419.90 MB/s | 0 | 0 |
| path_escaped | 87.6 | 490.75 MB/s | 48 | 1 |
| json_in_json | 119.5 | 452.02 MB/s | 64 | 1 |
| prose_with_quotes | 74.1 | 553.32 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5640.15 MB/s | 0 | 0 |
| mostly_clean_one_escape | 113.9 | 2687.41 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3288.63 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8067.32 MB/s | 0 | 0 |
| url_clean | 5.5 | 9534.05 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29599.11 MB/s | 0 | 0 |
| path_escaped | 46.5 | 924.09 MB/s | 0 | 0 |
| json_in_json | 70.2 | 768.95 MB/s | 0 | 0 |
| prose_with_quotes | 33.5 | 1223.40 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5088.25 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12353.53 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2299.0 | 1203.24 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2223.0 | 1244.14 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14512.0 | 1653.63 MB/s | 0 | 0 |
| stringObj/dispatch | 6770.0 | 3545.02 MB/s | 0 | 0 |
| numberObj/current | 6357.0 | 1604.25 MB/s | 0 | 0 |
| numberObj/dispatch | 2865.0 | 3559.61 MB/s | 0 | 0 |
| numberArr/current | 449.4 | 14688.06 MB/s | 0 | 0 |
| numberArr/dispatch | 451.3 | 14625.33 MB/s | 0 | 0 |
| nestedMixed/current | 20027.0 | 539.31 MB/s | 0 | 0 |
| nestedMixed/dispatch | 3437.0 | 3142.47 MB/s | 0 | 0 |
