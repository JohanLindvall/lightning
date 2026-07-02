# lightning main-module benchmarks

- generated 2026-07-02T14:28:16Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.4 | 525.84 MB/s | 16 | 1 |
| sentence_clean | 45.2 | 973.11 MB/s | 48 | 1 |
| url_clean | 41.2 | 1261.52 MB/s | 64 | 1 |
| log_line_clean | 114.8 | 2925.63 MB/s | 352 | 1 |
| path_with_backslash | 122.8 | 301.18 MB/s | 56 | 2 |
| json_in_json | 159.9 | 262.70 MB/s | 72 | 2 |
| prose_with_quotes | 95.9 | 396.14 MB/s | 64 | 2 |
| control_bytes | 119.2 | 201.26 MB/s | 56 | 2 |
| mostly_clean_one_quote | 136.4 | 2235.36 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1870.01 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2605.60 MB/s | 0 | 0 |
| url_clean | 22.6 | 2306.12 MB/s | 0 | 0 |
| log_line_clean | 43.7 | 7691.86 MB/s | 0 | 0 |
| path_with_backslash | 58.9 | 627.98 MB/s | 0 | 0 |
| json_in_json | 94.0 | 446.95 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1094.38 MB/s | 0 | 0 |
| control_bytes | 52.1 | 460.96 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.9 | 6646.09 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1015.0 | 1784.03 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1351.0 | 1340.34 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.1 | 3865.45 MB/s | 0 | 0 |
| sentence_clean | 4.8 | 9099.78 MB/s | 0 | 0 |
| url_clean | 4.8 | 10774.05 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31441.18 MB/s | 0 | 0 |
| path_escaped | 85.8 | 500.98 MB/s | 48 | 1 |
| json_in_json | 117.2 | 460.73 MB/s | 64 | 1 |
| prose_with_quotes | 72.8 | 563.28 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5646.78 MB/s | 0 | 0 |
| unicode_escaped_dense | 342.3 | 560.84 MB/s | 192 | 1 |
| mostly_clean_one_escape | 119.6 | 2558.24 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3287.82 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7855.54 MB/s | 0 | 0 |
| url_clean | 5.6 | 9284.63 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29593.64 MB/s | 0 | 0 |
| path_escaped | 45.7 | 941.52 MB/s | 0 | 0 |
| json_in_json | 70.9 | 761.50 MB/s | 0 | 0 |
| prose_with_quotes | 33.1 | 1239.58 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5088.44 MB/s | 0 | 0 |
| unicode_escaped_dense | 270.0 | 711.10 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12332.81 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 86.5 | — | 0 | 0 |
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
| — | 298.8 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2329.0 | 1187.87 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2252.0 | 1228.10 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4530.0 | 5297.69 MB/s | 0 | 0 |
| numberObj/goloop | 1842.0 | 5535.16 MB/s | 0 | 0 |
| nestedMixed/goloop | 2511.0 | 4301.18 MB/s | 0 | 0 |
| stringObj/neon | 2996.0 | 8011.04 MB/s | 0 | 0 |
| numberObj/neon | 1263.0 | 8074.38 MB/s | 0 | 0 |
| nestedMixed/neon | 1672.0 | 6460.87 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14453.0 | 1660.42 MB/s | 0 | 0 |
| stringObj/dispatch | 2996.0 | 8009.86 MB/s | 0 | 0 |
| numberObj/current | 6330.0 | 1610.96 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8066.07 MB/s | 0 | 0 |
| numberArr/current | 452.6 | 14583.97 MB/s | 0 | 0 |
| numberArr/dispatch | 451.3 | 14625.48 MB/s | 0 | 0 |
| nestedMixed/current | 17326.0 | 623.41 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1675.0 | 6448.48 MB/s | 0 | 0 |
