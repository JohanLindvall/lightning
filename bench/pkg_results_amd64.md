# lightning main-module benchmarks

- generated 2026-08-08T12:44:25Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 32.2 | 497.19 MB/s | 16 | 1 |
| sentence_clean | 43.8 | 1004.35 MB/s | 48 | 1 |
| url_clean | 50.1 | 1037.78 MB/s | 64 | 1 |
| log_line_clean | 94.9 | 3542.18 MB/s | 352 | 1 |
| path_with_backslash | 143.7 | 257.42 MB/s | 56 | 2 |
| json_in_json | 182.2 | 230.58 MB/s | 72 | 2 |
| prose_with_quotes | 106.8 | 355.90 MB/s | 64 | 2 |
| control_bytes | 135.3 | 177.38 MB/s | 56 | 2 |
| mostly_clean_one_quote | 105.5 | 2890.83 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 13.7 | 1166.20 MB/s | 0 | 0 |
| sentence_clean | 25.3 | 1739.87 MB/s | 0 | 0 |
| url_clean | 21.9 | 2378.11 MB/s | 0 | 0 |
| log_line_clean | 30.3 | 11100.42 MB/s | 0 | 0 |
| path_with_backslash | 77.3 | 478.67 MB/s | 0 | 0 |
| json_in_json | 110.5 | 380.06 MB/s | 0 | 0 |
| prose_with_quotes | 54.3 | 699.88 MB/s | 0 | 0 |
| control_bytes | 82.1 | 292.44 MB/s | 0 | 0 |
| mostly_clean_one_quote | 37.7 | 8096.94 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2497.0 | 4021.44 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2406.0 | 4173.81 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2330.0 | 4309.72 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12600.0 | 796.93 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 817.6 | 2215.03 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1476.0 | 1227.29 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 2992.91 MB/s | 0 | 0 |
| sentence_clean | 9.8 | 4480.00 MB/s | 0 | 0 |
| url_clean | 5.6 | 9248.66 MB/s | 0 | 0 |
| log_line_clean | 11.8 | 28376.41 MB/s | 0 | 0 |
| path_escaped | 89.5 | 480.58 MB/s | 48 | 1 |
| json_in_json | 114.8 | 470.56 MB/s | 64 | 1 |
| prose_with_quotes | 68.5 | 598.11 MB/s | 48 | 1 |
| unicode_heavy | 6.6 | 4560.76 MB/s | 0 | 0 |
| unicode_escaped_dense | 310.3 | 618.85 MB/s | 192 | 1 |
| mostly_clean_one_escape | 89.0 | 3438.38 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.9 | 2692.86 MB/s | 0 | 0 |
| sentence_clean | 5.9 | 7417.24 MB/s | 0 | 0 |
| url_clean | 5.9 | 8767.13 MB/s | 0 | 0 |
| log_line_clean | 9.4 | 35817.82 MB/s | 0 | 0 |
| path_escaped | 67.7 | 635.11 MB/s | 0 | 0 |
| json_in_json | 90.4 | 597.18 MB/s | 0 | 0 |
| prose_with_quotes | 48.6 | 843.97 MB/s | 0 | 0 |
| unicode_heavy | 7.2 | 4175.09 MB/s | 0 | 0 |
| unicode_escaped_dense | 267.7 | 717.23 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.4 | 12547.32 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 81.5 | — | 0 | 0 |
| append_empty | 23.6 | — | 0 | 0 |
| replace | 53.2 | — | 0 | 0 |
| create_nested | 52.7 | — | 0 | 0 |
| overwrite_nonobject | 57.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 125.1 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 338.9 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 98.4 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 139.9 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2303.0 | 1200.87 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2197.0 | 1259.18 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 252.8 | 735.67 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 85.9 | — | 24 | 1 |
| arena | 72.1 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3775.0 | 6357.71 MB/s | 0 | 0 |
| numberObj/goloop | 1317.0 | 7744.70 MB/s | 0 | 0 |
| nestedMixed/goloop | 2241.0 | 4820.79 MB/s | 0 | 0 |
| stringObj/avx2 | 2103.0 | 11411.71 MB/s | 0 | 0 |
| numberObj/avx2 | 760.9 | 13402.93 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1336.0 | 8082.14 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10264.0 | 2338.17 MB/s | 0 | 0 |
| stringObj/dispatch | 2110.0 | 11374.10 MB/s | 0 | 0 |
| numberObj/current | 5377.0 | 1896.72 MB/s | 0 | 0 |
| numberObj/dispatch | 765.6 | 13320.60 MB/s | 0 | 0 |
| numberArr/current | 236.6 | 27899.74 MB/s | 0 | 0 |
| numberArr/dispatch | 239.1 | 27606.83 MB/s | 0 | 0 |
| nestedMixed/current | 15681.0 | 688.78 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1341.0 | 8053.28 MB/s | 0 | 0 |
