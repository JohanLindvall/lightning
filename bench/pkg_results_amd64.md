# lightning main-module benchmarks

- generated 2026-07-26T10:33:34Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 33.2 | 481.51 MB/s | 16 | 1 |
| sentence_clean | 43.5 | 1010.74 MB/s | 48 | 1 |
| url_clean | 42.7 | 1216.65 MB/s | 64 | 1 |
| log_line_clean | 96.5 | 3481.91 MB/s | 352 | 1 |
| path_with_backslash | 139.2 | 265.86 MB/s | 56 | 2 |
| json_in_json | 178.8 | 234.84 MB/s | 72 | 2 |
| prose_with_quotes | 107.3 | 354.06 MB/s | 64 | 2 |
| control_bytes | 130.4 | 184.09 MB/s | 56 | 2 |
| mostly_clean_one_quote | 106.2 | 2871.83 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 13.7 | 1164.25 MB/s | 0 | 0 |
| sentence_clean | 25.4 | 1734.97 MB/s | 0 | 0 |
| url_clean | 21.9 | 2377.28 MB/s | 0 | 0 |
| log_line_clean | 29.7 | 11325.92 MB/s | 0 | 0 |
| path_with_backslash | 73.1 | 505.97 MB/s | 0 | 0 |
| json_in_json | 110.2 | 381.28 MB/s | 0 | 0 |
| prose_with_quotes | 47.4 | 801.92 MB/s | 0 | 0 |
| control_bytes | 65.5 | 366.24 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.3 | 8634.43 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 857.0 | 2113.20 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1584.0 | 1143.35 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2844.12 MB/s | 0 | 0 |
| sentence_clean | 5.9 | 7417.30 MB/s | 0 | 0 |
| url_clean | 5.9 | 8754.62 MB/s | 0 | 0 |
| log_line_clean | 9.1 | 37046.31 MB/s | 0 | 0 |
| path_escaped | 86.8 | 495.46 MB/s | 48 | 1 |
| json_in_json | 113.9 | 474.26 MB/s | 64 | 1 |
| prose_with_quotes | 69.1 | 593.24 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4367.11 MB/s | 0 | 0 |
| unicode_escaped_dense | 374.8 | 512.31 MB/s | 192 | 1 |
| mostly_clean_one_escape | 89.0 | 3439.37 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.2 | 2561.28 MB/s | 0 | 0 |
| sentence_clean | 6.2 | 7040.76 MB/s | 0 | 0 |
| url_clean | 6.2 | 8322.11 MB/s | 0 | 0 |
| log_line_clean | 9.7 | 34697.80 MB/s | 0 | 0 |
| path_escaped | 63.1 | 681.47 MB/s | 0 | 0 |
| json_in_json | 89.0 | 606.86 MB/s | 0 | 0 |
| prose_with_quotes | 46.9 | 874.24 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 3996.03 MB/s | 0 | 0 |
| unicode_escaped_dense | 328.0 | 585.31 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.1 | 12695.70 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 83.2 | — | 0 | 0 |
| append_empty | 22.8 | — | 0 | 0 |
| replace | 54.9 | — | 0 | 0 |
| create_nested | 53.1 | — | 0 | 0 |
| overwrite_nonobject | 62.2 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 126.1 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 329.0 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2394.0 | 1155.30 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2318.0 | 1193.28 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 252.6 | 736.32 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3846.0 | 6239.39 MB/s | 0 | 0 |
| numberObj/goloop | 1338.0 | 7622.79 MB/s | 0 | 0 |
| nestedMixed/goloop | 2301.0 | 4694.66 MB/s | 0 | 0 |
| stringObj/avx2 | 2019.0 | 11888.74 MB/s | 0 | 0 |
| numberObj/avx2 | 750.2 | 13594.55 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1333.0 | 8105.59 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10156.0 | 2362.98 MB/s | 0 | 0 |
| stringObj/dispatch | 2008.0 | 11953.17 MB/s | 0 | 0 |
| numberObj/current | 5423.0 | 1880.49 MB/s | 0 | 0 |
| numberObj/dispatch | 758.3 | 13449.32 MB/s | 0 | 0 |
| numberArr/current | 206.0 | 32050.30 MB/s | 0 | 0 |
| numberArr/dispatch | 207.3 | 31847.14 MB/s | 0 | 0 |
| nestedMixed/current | 17290.0 | 624.70 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1328.0 | 8135.63 MB/s | 0 | 0 |
