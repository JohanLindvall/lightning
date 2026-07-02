# lightning main-module benchmarks

- generated 2026-07-02T08:46:52Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.5 | 524.14 MB/s | 16 | 1 |
| sentence_clean | 45.5 | 966.75 MB/s | 48 | 1 |
| url_clean | 42.2 | 1230.94 MB/s | 64 | 1 |
| log_line_clean | 113.7 | 2955.27 MB/s | 352 | 1 |
| path_with_backslash | 203.6 | 181.76 MB/s | 184 | 4 |
| json_in_json | 236.4 | 177.63 MB/s | 216 | 4 |
| prose_with_quotes | 129.2 | 294.18 MB/s | 128 | 3 |
| control_bytes | 144.9 | 165.66 MB/s | 104 | 3 |
| mostly_clean_one_quote | 157.3 | 1938.44 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.7 | 1830.13 MB/s | 0 | 0 |
| sentence_clean | 17.1 | 2568.64 MB/s | 0 | 0 |
| url_clean | 22.4 | 2323.53 MB/s | 0 | 0 |
| log_line_clean | 44.1 | 7617.77 MB/s | 0 | 0 |
| path_with_backslash | 58.6 | 631.48 MB/s | 0 | 0 |
| json_in_json | 96.6 | 434.65 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1096.00 MB/s | 0 | 0 |
| control_bytes | 52.0 | 461.15 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.7 | 6668.41 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1016.0 | 1781.64 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1353.0 | 1338.65 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3775.53 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9009.07 MB/s | 0 | 0 |
| url_clean | 4.9 | 10649.97 MB/s | 0 | 0 |
| log_line_clean | 10.8 | 31169.09 MB/s | 0 | 0 |
| path_escaped | 89.0 | 482.91 MB/s | 48 | 1 |
| json_in_json | 120.3 | 449.05 MB/s | 64 | 1 |
| prose_with_quotes | 74.0 | 554.40 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5653.61 MB/s | 0 | 0 |
| mostly_clean_one_escape | 118.1 | 2592.03 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3288.92 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7856.39 MB/s | 0 | 0 |
| url_clean | 5.6 | 9285.04 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29586.00 MB/s | 0 | 0 |
| path_escaped | 50.8 | 847.26 MB/s | 0 | 0 |
| json_in_json | 72.7 | 742.68 MB/s | 0 | 0 |
| prose_with_quotes | 34.3 | 1195.44 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5088.24 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12371.63 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 86.0 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 59.6 | — | 0 | 0 |
| create_nested | 47.6 | — | 0 | 0 |
| overwrite_nonobject | 62.2 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 118.5 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 360.6 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2303.0 | 1200.81 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2234.0 | 1238.40 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4570.0 | 5251.52 MB/s | 0 | 0 |
| numberObj/goloop | 1850.0 | 5511.90 MB/s | 0 | 0 |
| nestedMixed/goloop | 2520.0 | 4286.59 MB/s | 0 | 0 |
| stringObj/neon | 2996.0 | 8010.08 MB/s | 0 | 0 |
| numberObj/neon | 1263.0 | 8074.36 MB/s | 0 | 0 |
| nestedMixed/neon | 1675.0 | 6449.82 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14452.0 | 1660.58 MB/s | 0 | 0 |
| stringObj/dispatch | 2996.0 | 8010.06 MB/s | 0 | 0 |
| numberObj/current | 6294.0 | 1620.15 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8066.16 MB/s | 0 | 0 |
| numberArr/current | 449.6 | 14682.61 MB/s | 0 | 0 |
| numberArr/dispatch | 451.1 | 14632.19 MB/s | 0 | 0 |
| nestedMixed/current | 16760.0 | 644.44 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1674.0 | 6451.55 MB/s | 0 | 0 |
