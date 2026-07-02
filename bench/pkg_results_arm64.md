# lightning main-module benchmarks

- generated 2026-07-02T07:11:36Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.3 | 527.84 MB/s | 16 | 1 |
| sentence_clean | 45.6 | 964.00 MB/s | 48 | 1 |
| url_clean | 42.3 | 1229.09 MB/s | 64 | 1 |
| log_line_clean | 114.1 | 2943.89 MB/s | 352 | 1 |
| path_with_backslash | 201.2 | 183.91 MB/s | 184 | 4 |
| json_in_json | 236.4 | 177.63 MB/s | 216 | 4 |
| prose_with_quotes | 129.1 | 294.33 MB/s | 128 | 3 |
| control_bytes | 144.8 | 165.69 MB/s | 104 | 3 |
| mostly_clean_one_quote | 158.5 | 1924.60 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.7 | 1829.59 MB/s | 0 | 0 |
| sentence_clean | 17.1 | 2571.75 MB/s | 0 | 0 |
| url_clean | 22.4 | 2326.88 MB/s | 0 | 0 |
| log_line_clean | 44.1 | 7616.97 MB/s | 0 | 0 |
| path_with_backslash | 58.6 | 630.87 MB/s | 0 | 0 |
| json_in_json | 96.8 | 433.86 MB/s | 0 | 0 |
| prose_with_quotes | 34.6 | 1097.70 MB/s | 0 | 0 |
| control_bytes | 52.0 | 461.29 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.7 | 6676.45 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1025.0 | 1766.23 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1350.0 | 1341.41 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3777.07 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9011.07 MB/s | 0 | 0 |
| url_clean | 4.9 | 10652.65 MB/s | 0 | 0 |
| log_line_clean | 10.8 | 31180.37 MB/s | 0 | 0 |
| path_escaped | 87.9 | 489.43 MB/s | 48 | 1 |
| json_in_json | 120.8 | 446.89 MB/s | 64 | 1 |
| prose_with_quotes | 72.9 | 562.32 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5653.91 MB/s | 0 | 0 |
| mostly_clean_one_escape | 117.8 | 2598.10 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3289.03 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7856.13 MB/s | 0 | 0 |
| url_clean | 5.6 | 9284.73 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29585.78 MB/s | 0 | 0 |
| path_escaped | 51.2 | 839.32 MB/s | 0 | 0 |
| json_in_json | 71.5 | 754.87 MB/s | 0 | 0 |
| prose_with_quotes | 34.0 | 1204.05 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5088.72 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12392.44 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 86.0 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 59.7 | — | 0 | 0 |
| create_nested | 47.6 | — | 0 | 0 |
| overwrite_nonobject | 62.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 117.9 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 358.4 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2299.0 | 1202.94 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2232.0 | 1239.19 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4569.0 | 5252.69 MB/s | 0 | 0 |
| numberObj/goloop | 1851.0 | 5510.77 MB/s | 0 | 0 |
| nestedMixed/goloop | 2514.0 | 4295.54 MB/s | 0 | 0 |
| stringObj/neon | 2996.0 | 8011.34 MB/s | 0 | 0 |
| numberObj/neon | 1263.0 | 8073.86 MB/s | 0 | 0 |
| nestedMixed/neon | 1672.0 | 6459.10 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14464.0 | 1659.16 MB/s | 0 | 0 |
| stringObj/dispatch | 2996.0 | 8009.40 MB/s | 0 | 0 |
| numberObj/current | 6301.0 | 1618.43 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8066.15 MB/s | 0 | 0 |
| numberArr/current | 449.3 | 14691.50 MB/s | 0 | 0 |
| numberArr/dispatch | 451.2 | 14631.04 MB/s | 0 | 0 |
| nestedMixed/current | 17005.0 | 635.15 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1674.0 | 6450.63 MB/s | 0 | 0 |
