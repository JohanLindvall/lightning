# lightning main-module benchmarks

- generated 2026-06-24T06:13:30Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.7 | 521.36 MB/s | 16 | 1 |
| sentence_clean | 45.2 | 973.60 MB/s | 48 | 1 |
| url_clean | 41.7 | 1247.05 MB/s | 64 | 1 |
| log_line_clean | 115.4 | 2910.76 MB/s | 352 | 1 |
| path_with_backslash | 206.0 | 179.59 MB/s | 184 | 4 |
| json_in_json | 230.6 | 182.10 MB/s | 216 | 4 |
| prose_with_quotes | 127.6 | 297.69 MB/s | 128 | 3 |
| control_bytes | 143.4 | 167.33 MB/s | 104 | 3 |
| mostly_clean_one_quote | 156.6 | 1947.34 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1870.16 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2609.13 MB/s | 0 | 0 |
| url_clean | 22.4 | 2316.16 MB/s | 0 | 0 |
| log_line_clean | 43.7 | 7680.95 MB/s | 0 | 0 |
| path_with_backslash | 59.4 | 622.76 MB/s | 0 | 0 |
| json_in_json | 94.3 | 445.18 MB/s | 0 | 0 |
| prose_with_quotes | 34.6 | 1096.61 MB/s | 0 | 0 |
| control_bytes | 53.0 | 453.23 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.6 | 6692.32 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1016.0 | 1782.72 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1341.0 | 1350.96 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3767.58 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 8987.57 MB/s | 0 | 0 |
| url_clean | 4.9 | 10638.29 MB/s | 0 | 0 |
| log_line_clean | 10.8 | 31156.43 MB/s | 0 | 0 |
| path_escaped | 88.2 | 487.49 MB/s | 48 | 1 |
| json_in_json | 120.4 | 448.44 MB/s | 64 | 1 |
| prose_with_quotes | 74.5 | 550.58 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5652.48 MB/s | 0 | 0 |
| mostly_clean_one_escape | 117.7 | 2599.42 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3287.60 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8064.73 MB/s | 0 | 0 |
| url_clean | 5.5 | 9530.16 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29598.11 MB/s | 0 | 0 |
| path_escaped | 46.0 | 935.50 MB/s | 0 | 0 |
| json_in_json | 70.1 | 770.62 MB/s | 0 | 0 |
| prose_with_quotes | 33.2 | 1233.21 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5088.01 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12354.42 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2294.0 | 1205.67 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2232.0 | 1239.50 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14437.0 | 1662.23 MB/s | 0 | 0 |
| stringObj/dispatch | 3873.0 | 6196.18 MB/s | 0 | 0 |
| numberObj/current | 6364.0 | 1602.47 MB/s | 0 | 0 |
| numberObj/dispatch | 1652.0 | 6171.87 MB/s | 0 | 0 |
| numberArr/current | 464.5 | 14210.38 MB/s | 0 | 0 |
| numberArr/dispatch | 466.8 | 14141.93 MB/s | 0 | 0 |
| nestedMixed/current | 20574.0 | 524.98 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2135.0 | 5058.52 MB/s | 0 | 0 |
