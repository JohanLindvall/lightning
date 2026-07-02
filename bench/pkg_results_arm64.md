# lightning main-module benchmarks

- generated 2026-07-02T12:24:16Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.8 | 537.03 MB/s | 16 | 1 |
| sentence_clean | 44.7 | 983.85 MB/s | 48 | 1 |
| url_clean | 40.1 | 1297.75 MB/s | 64 | 1 |
| log_line_clean | 110.4 | 3043.00 MB/s | 352 | 1 |
| path_with_backslash | 121.8 | 303.83 MB/s | 56 | 2 |
| json_in_json | 159.2 | 263.76 MB/s | 72 | 2 |
| prose_with_quotes | 94.6 | 401.54 MB/s | 64 | 2 |
| control_bytes | 118.3 | 202.85 MB/s | 56 | 2 |
| mostly_clean_one_quote | 133.7 | 2280.83 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1869.72 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2605.37 MB/s | 0 | 0 |
| url_clean | 22.5 | 2306.50 MB/s | 0 | 0 |
| log_line_clean | 43.6 | 7697.27 MB/s | 0 | 0 |
| path_with_backslash | 58.6 | 631.88 MB/s | 0 | 0 |
| json_in_json | 93.8 | 447.85 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1095.17 MB/s | 0 | 0 |
| control_bytes | 51.8 | 463.38 MB/s | 0 | 0 |
| mostly_clean_one_quote | 46.0 | 6635.85 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1017.0 | 1781.05 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1353.0 | 1338.20 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.1 | 3867.16 MB/s | 0 | 0 |
| sentence_clean | 4.8 | 9091.20 MB/s | 0 | 0 |
| url_clean | 4.8 | 10773.59 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31443.03 MB/s | 0 | 0 |
| path_escaped | 84.3 | 510.32 MB/s | 48 | 1 |
| json_in_json | 115.8 | 466.37 MB/s | 64 | 1 |
| prose_with_quotes | 71.4 | 573.90 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5644.73 MB/s | 0 | 0 |
| unicode_escaped_dense | 340.3 | 564.18 MB/s | 192 | 1 |
| mostly_clean_one_escape | 114.0 | 2685.25 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3288.20 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7855.65 MB/s | 0 | 0 |
| url_clean | 5.6 | 9284.03 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29588.93 MB/s | 0 | 0 |
| path_escaped | 45.5 | 945.97 MB/s | 0 | 0 |
| json_in_json | 70.7 | 763.55 MB/s | 0 | 0 |
| prose_with_quotes | 33.0 | 1242.04 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5087.79 MB/s | 0 | 0 |
| unicode_escaped_dense | 269.9 | 711.41 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.9 | 12313.75 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 86.4 | — | 0 | 0 |
| append_empty | 18.2 | — | 0 | 0 |
| replace | 59.2 | — | 0 | 0 |
| create_nested | 47.8 | — | 0 | 0 |
| overwrite_nonobject | 57.8 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 120.5 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 299.5 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2316.0 | 1194.24 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2263.0 | 1222.17 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4530.0 | 5297.61 MB/s | 0 | 0 |
| numberObj/goloop | 1849.0 | 5515.89 MB/s | 0 | 0 |
| nestedMixed/goloop | 2511.0 | 4301.12 MB/s | 0 | 0 |
| stringObj/neon | 2996.0 | 8010.33 MB/s | 0 | 0 |
| numberObj/neon | 1263.0 | 8071.83 MB/s | 0 | 0 |
| nestedMixed/neon | 1672.0 | 6461.14 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14482.0 | 1657.05 MB/s | 0 | 0 |
| stringObj/dispatch | 2996.0 | 8009.21 MB/s | 0 | 0 |
| numberObj/current | 6333.0 | 1610.19 MB/s | 0 | 0 |
| numberObj/dispatch | 1265.0 | 8063.45 MB/s | 0 | 0 |
| numberArr/current | 449.6 | 14682.01 MB/s | 0 | 0 |
| numberArr/dispatch | 451.3 | 14625.93 MB/s | 0 | 0 |
| nestedMixed/current | 16866.0 | 640.42 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1675.0 | 6449.13 MB/s | 0 | 0 |
