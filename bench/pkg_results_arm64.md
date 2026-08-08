# lightning main-module benchmarks

- generated 2026-08-08T12:44:21Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.1 | 532.53 MB/s | 16 | 1 |
| sentence_clean | 45.2 | 973.40 MB/s | 48 | 1 |
| url_clean | 41.4 | 1256.68 MB/s | 64 | 1 |
| log_line_clean | 112.5 | 2987.38 MB/s | 352 | 1 |
| path_with_backslash | 122.2 | 302.67 MB/s | 56 | 2 |
| json_in_json | 158.9 | 264.35 MB/s | 72 | 2 |
| prose_with_quotes | 94.5 | 402.18 MB/s | 64 | 2 |
| control_bytes | 117.7 | 203.99 MB/s | 56 | 2 |
| mostly_clean_one_quote | 134.1 | 2274.29 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1868.75 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2601.74 MB/s | 0 | 0 |
| url_clean | 22.4 | 2318.86 MB/s | 0 | 0 |
| log_line_clean | 43.9 | 7647.62 MB/s | 0 | 0 |
| path_with_backslash | 59.3 | 623.49 MB/s | 0 | 0 |
| json_in_json | 94.7 | 443.62 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1094.04 MB/s | 0 | 0 |
| control_bytes | 52.6 | 455.91 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.6 | 6683.00 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3075.0 | 3265.69 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2931.0 | 3425.49 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2914.0 | 3446.00 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12202.0 | 822.93 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 986.2 | 1836.33 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1320.0 | 1372.19 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.4 | 3615.11 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9028.08 MB/s | 0 | 0 |
| url_clean | 4.9 | 10669.18 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31557.75 MB/s | 0 | 0 |
| path_escaped | 83.2 | 517.16 MB/s | 48 | 1 |
| json_in_json | 114.5 | 471.60 MB/s | 64 | 1 |
| prose_with_quotes | 70.2 | 583.77 MB/s | 48 | 1 |
| unicode_heavy | 4.2 | 7211.78 MB/s | 0 | 0 |
| unicode_escaped_dense | 290.7 | 660.43 MB/s | 192 | 1 |
| mostly_clean_one_escape | 116.7 | 2622.73 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.0 | 3191.50 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8067.70 MB/s | 0 | 0 |
| url_clean | 5.5 | 9533.73 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29592.41 MB/s | 0 | 0 |
| path_escaped | 44.6 | 963.43 MB/s | 0 | 0 |
| json_in_json | 68.4 | 789.76 MB/s | 0 | 0 |
| prose_with_quotes | 32.1 | 1277.49 MB/s | 0 | 0 |
| unicode_heavy | 4.9 | 6165.23 MB/s | 0 | 0 |
| unicode_escaped_dense | 221.7 | 865.86 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12332.28 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 84.6 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 57.1 | — | 0 | 0 |
| create_nested | 46.3 | — | 0 | 0 |
| overwrite_nonobject | 54.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 116.7 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 294.0 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 109.4 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 139.2 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2328.0 | 1188.03 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2219.0 | 1246.65 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 271.3 | 685.58 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 84.5 | — | 24 | 1 |
| arena | 70.0 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4553.0 | 5270.99 MB/s | 0 | 0 |
| numberObj/goloop | 1840.0 | 5541.49 MB/s | 0 | 0 |
| nestedMixed/goloop | 2500.0 | 4319.71 MB/s | 0 | 0 |
| stringObj/neon | 2995.0 | 8012.27 MB/s | 0 | 0 |
| numberObj/neon | 1262.0 | 8079.71 MB/s | 0 | 0 |
| nestedMixed/neon | 1671.0 | 6463.77 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14419.0 | 1664.29 MB/s | 0 | 0 |
| stringObj/dispatch | 2997.0 | 8007.02 MB/s | 0 | 0 |
| numberObj/current | 6341.0 | 1608.27 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8069.72 MB/s | 0 | 0 |
| numberArr/current | 449.4 | 14688.77 MB/s | 0 | 0 |
| numberArr/dispatch | 450.9 | 14638.34 MB/s | 0 | 0 |
| nestedMixed/current | 17150.0 | 629.79 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1674.0 | 6451.20 MB/s | 0 | 0 |
