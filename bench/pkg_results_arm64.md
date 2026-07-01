# lightning main-module benchmarks

- generated 2026-07-01T22:07:08Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.1 | 531.01 MB/s | 16 | 1 |
| sentence_clean | 45.1 | 976.25 MB/s | 48 | 1 |
| url_clean | 40.7 | 1276.72 MB/s | 64 | 1 |
| log_line_clean | 112.4 | 2988.87 MB/s | 352 | 1 |
| path_with_backslash | 195.8 | 189.00 MB/s | 184 | 4 |
| json_in_json | 232.0 | 181.00 MB/s | 216 | 4 |
| prose_with_quotes | 126.5 | 300.51 MB/s | 128 | 3 |
| control_bytes | 143.4 | 167.32 MB/s | 104 | 3 |
| mostly_clean_one_quote | 155.8 | 1957.68 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.7 | 1829.64 MB/s | 0 | 0 |
| sentence_clean | 17.1 | 2572.04 MB/s | 0 | 0 |
| url_clean | 22.3 | 2327.64 MB/s | 0 | 0 |
| log_line_clean | 44.2 | 7609.47 MB/s | 0 | 0 |
| path_with_backslash | 58.6 | 631.21 MB/s | 0 | 0 |
| json_in_json | 96.6 | 434.66 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1095.10 MB/s | 0 | 0 |
| control_bytes | 52.1 | 460.92 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.7 | 6669.00 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1016.0 | 1782.40 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1347.0 | 1344.71 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3775.70 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9009.31 MB/s | 0 | 0 |
| url_clean | 4.9 | 10647.83 MB/s | 0 | 0 |
| log_line_clean | 10.8 | 31169.87 MB/s | 0 | 0 |
| path_escaped | 84.8 | 507.08 MB/s | 48 | 1 |
| json_in_json | 113.4 | 476.16 MB/s | 64 | 1 |
| prose_with_quotes | 69.9 | 586.72 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5653.31 MB/s | 0 | 0 |
| mostly_clean_one_escape | 115.2 | 2656.17 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3289.10 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7856.03 MB/s | 0 | 0 |
| url_clean | 5.6 | 9284.50 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29588.22 MB/s | 0 | 0 |
| path_escaped | 47.0 | 914.28 MB/s | 0 | 0 |
| json_in_json | 73.3 | 736.28 MB/s | 0 | 0 |
| prose_with_quotes | 34.2 | 1197.79 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5088.51 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12385.61 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 86.0 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 59.7 | — | 0 | 0 |
| create_nested | 47.4 | — | 0 | 0 |
| overwrite_nonobject | 62.3 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 118.2 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 360.6 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2299.0 | 1203.18 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2245.0 | 1232.29 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4568.0 | 5253.71 MB/s | 0 | 0 |
| numberObj/goloop | 1850.0 | 5511.09 MB/s | 0 | 0 |
| nestedMixed/goloop | 2518.0 | 4289.45 MB/s | 0 | 0 |
| stringObj/neon | 2995.0 | 8011.48 MB/s | 0 | 0 |
| numberObj/neon | 1263.0 | 8074.29 MB/s | 0 | 0 |
| nestedMixed/neon | 1674.0 | 6453.31 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14474.0 | 1657.98 MB/s | 0 | 0 |
| stringObj/dispatch | 2996.0 | 8010.12 MB/s | 0 | 0 |
| numberObj/current | 6298.0 | 1619.30 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8065.66 MB/s | 0 | 0 |
| numberArr/current | 449.5 | 14685.28 MB/s | 0 | 0 |
| numberArr/dispatch | 451.3 | 14627.07 MB/s | 0 | 0 |
| nestedMixed/current | 17071.0 | 632.71 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1674.0 | 6451.26 MB/s | 0 | 0 |
