# lightning main-module benchmarks

- generated 2026-07-27T11:14:58Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.9 | 516.98 MB/s | 16 | 1 |
| sentence_clean | 45.8 | 961.69 MB/s | 48 | 1 |
| url_clean | 42.8 | 1213.78 MB/s | 64 | 1 |
| log_line_clean | 116.4 | 2886.95 MB/s | 352 | 1 |
| path_with_backslash | 124.3 | 297.66 MB/s | 56 | 2 |
| json_in_json | 166.7 | 251.90 MB/s | 72 | 2 |
| prose_with_quotes | 96.6 | 393.55 MB/s | 64 | 2 |
| control_bytes | 119.8 | 200.40 MB/s | 56 | 2 |
| mostly_clean_one_quote | 138.3 | 2204.90 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.8 | 1828.00 MB/s | 0 | 0 |
| sentence_clean | 17.1 | 2569.50 MB/s | 0 | 0 |
| url_clean | 22.4 | 2317.95 MB/s | 0 | 0 |
| log_line_clean | 44.1 | 7623.94 MB/s | 0 | 0 |
| path_with_backslash | 59.0 | 627.66 MB/s | 0 | 0 |
| json_in_json | 97.0 | 432.91 MB/s | 0 | 0 |
| prose_with_quotes | 34.8 | 1093.66 MB/s | 0 | 0 |
| control_bytes | 52.1 | 460.94 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.5 | 6709.21 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 999.7 | 1811.50 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1318.0 | 1373.66 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.4 | 3617.16 MB/s | 0 | 0 |
| sentence_clean | 4.8 | 9121.09 MB/s | 0 | 0 |
| url_clean | 4.8 | 10782.20 MB/s | 0 | 0 |
| log_line_clean | 10.6 | 31582.06 MB/s | 0 | 0 |
| path_escaped | 84.1 | 511.34 MB/s | 48 | 1 |
| json_in_json | 113.0 | 477.78 MB/s | 64 | 1 |
| prose_with_quotes | 70.8 | 579.51 MB/s | 48 | 1 |
| unicode_heavy | 4.1 | 7261.24 MB/s | 0 | 0 |
| unicode_escaped_dense | 295.5 | 649.66 MB/s | 192 | 1 |
| mostly_clean_one_escape | 122.0 | 2508.52 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.0 | 3192.67 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7856.16 MB/s | 0 | 0 |
| url_clean | 5.6 | 9284.92 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29595.25 MB/s | 0 | 0 |
| path_escaped | 46.9 | 917.00 MB/s | 0 | 0 |
| json_in_json | 68.9 | 783.33 MB/s | 0 | 0 |
| prose_with_quotes | 32.8 | 1251.76 MB/s | 0 | 0 |
| unicode_heavy | 4.9 | 6164.67 MB/s | 0 | 0 |
| unicode_escaped_dense | 229.2 | 837.64 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.6 | 12415.78 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 83.7 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 57.2 | — | 0 | 0 |
| create_nested | 46.2 | — | 0 | 0 |
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
| — | 295.7 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 108.6 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 138.2 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2342.0 | 1180.90 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2244.0 | 1232.71 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 269.2 | 690.84 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 84.9 | — | 24 | 1 |
| arena | 70.2 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4541.0 | 5285.27 MB/s | 0 | 0 |
| numberObj/goloop | 1841.0 | 5538.73 MB/s | 0 | 0 |
| nestedMixed/goloop | 2486.0 | 4345.30 MB/s | 0 | 0 |
| stringObj/neon | 2996.0 | 8010.45 MB/s | 0 | 0 |
| numberObj/neon | 1262.0 | 8078.48 MB/s | 0 | 0 |
| nestedMixed/neon | 1671.0 | 6463.49 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14424.0 | 1663.74 MB/s | 0 | 0 |
| stringObj/dispatch | 2997.0 | 8006.86 MB/s | 0 | 0 |
| numberObj/current | 6357.0 | 1604.19 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8070.94 MB/s | 0 | 0 |
| numberArr/current | 449.4 | 14687.38 MB/s | 0 | 0 |
| numberArr/dispatch | 450.9 | 14639.44 MB/s | 0 | 0 |
| nestedMixed/current | 17169.0 | 629.09 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1675.0 | 6449.72 MB/s | 0 | 0 |
