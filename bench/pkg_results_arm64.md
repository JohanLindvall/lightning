# lightning main-module benchmarks

- generated 2026-07-01T12:50:53Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.7 | 521.90 MB/s | 16 | 1 |
| sentence_clean | 45.3 | 970.66 MB/s | 48 | 1 |
| url_clean | 42.4 | 1227.21 MB/s | 64 | 1 |
| log_line_clean | 114.5 | 2933.74 MB/s | 352 | 1 |
| path_with_backslash | 204.0 | 181.36 MB/s | 184 | 4 |
| json_in_json | 230.2 | 182.48 MB/s | 216 | 4 |
| prose_with_quotes | 127.8 | 297.35 MB/s | 128 | 3 |
| control_bytes | 143.8 | 166.85 MB/s | 104 | 3 |
| mostly_clean_one_quote | 158.5 | 1924.72 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1870.08 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2609.75 MB/s | 0 | 0 |
| url_clean | 22.4 | 2318.54 MB/s | 0 | 0 |
| log_line_clean | 43.8 | 7674.21 MB/s | 0 | 0 |
| path_with_backslash | 58.2 | 635.26 MB/s | 0 | 0 |
| json_in_json | 94.4 | 444.90 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1094.80 MB/s | 0 | 0 |
| control_bytes | 52.4 | 458.10 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.6 | 6687.82 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1016.0 | 1782.80 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1341.0 | 1350.51 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3767.35 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9006.54 MB/s | 0 | 0 |
| url_clean | 4.9 | 10647.14 MB/s | 0 | 0 |
| log_line_clean | 10.8 | 31150.61 MB/s | 0 | 0 |
| path_escaped | 83.7 | 513.90 MB/s | 48 | 1 |
| json_in_json | 108.0 | 500.22 MB/s | 64 | 1 |
| prose_with_quotes | 72.1 | 568.97 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5652.39 MB/s | 0 | 0 |
| mostly_clean_one_escape | 116.9 | 2616.89 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3287.81 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8064.04 MB/s | 0 | 0 |
| url_clean | 5.5 | 9531.19 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29588.81 MB/s | 0 | 0 |
| path_escaped | 46.1 | 932.74 MB/s | 0 | 0 |
| json_in_json | 70.8 | 763.01 MB/s | 0 | 0 |
| prose_with_quotes | 33.5 | 1225.17 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5087.04 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12332.08 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 87.3 | — | 0 | 0 |
| append_empty | 19.7 | — | 0 | 0 |
| replace | 59.2 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2286.0 | 1209.92 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2227.0 | 1241.85 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14434.0 | 1662.57 MB/s | 0 | 0 |
| stringObj/dispatch | 3871.0 | 6199.87 MB/s | 0 | 0 |
| numberObj/current | 6343.0 | 1607.68 MB/s | 0 | 0 |
| numberObj/dispatch | 1652.0 | 6171.46 MB/s | 0 | 0 |
| numberArr/current | 464.6 | 14209.40 MB/s | 0 | 0 |
| numberArr/dispatch | 466.9 | 14136.47 MB/s | 0 | 0 |
| nestedMixed/current | 19867.0 | 543.66 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2131.0 | 5067.51 MB/s | 0 | 0 |
