# lightning main-module benchmarks

- generated 2026-07-27T11:14:59Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.5 | 523.89 MB/s | 16 | 1 |
| sentence_clean | 44.5 | 989.53 MB/s | 48 | 1 |
| url_clean | 41.7 | 1247.11 MB/s | 64 | 1 |
| log_line_clean | 93.5 | 3594.02 MB/s | 352 | 1 |
| path_with_backslash | 138.9 | 266.45 MB/s | 56 | 2 |
| json_in_json | 177.5 | 236.59 MB/s | 72 | 2 |
| prose_with_quotes | 108.1 | 351.45 MB/s | 64 | 2 |
| control_bytes | 131.3 | 182.77 MB/s | 56 | 2 |
| mostly_clean_one_quote | 102.3 | 2981.47 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 13.1 | 1219.45 MB/s | 0 | 0 |
| sentence_clean | 25.6 | 1717.94 MB/s | 0 | 0 |
| url_clean | 21.9 | 2376.13 MB/s | 0 | 0 |
| log_line_clean | 29.7 | 11319.52 MB/s | 0 | 0 |
| path_with_backslash | 73.3 | 504.65 MB/s | 0 | 0 |
| json_in_json | 110.2 | 381.06 MB/s | 0 | 0 |
| prose_with_quotes | 47.8 | 795.53 MB/s | 0 | 0 |
| control_bytes | 65.1 | 368.80 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.0 | 8717.86 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 833.9 | 2171.81 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1473.0 | 1229.68 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2845.54 MB/s | 0 | 0 |
| sentence_clean | 5.9 | 7417.79 MB/s | 0 | 0 |
| url_clean | 5.9 | 8763.40 MB/s | 0 | 0 |
| log_line_clean | 9.1 | 37126.21 MB/s | 0 | 0 |
| path_escaped | 86.1 | 499.54 MB/s | 48 | 1 |
| json_in_json | 110.3 | 489.53 MB/s | 64 | 1 |
| prose_with_quotes | 63.2 | 648.81 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4363.43 MB/s | 0 | 0 |
| unicode_escaped_dense | 337.4 | 569.14 MB/s | 192 | 1 |
| mostly_clean_one_escape | 87.7 | 3487.44 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.2 | 2563.17 MB/s | 0 | 0 |
| sentence_clean | 6.2 | 7057.14 MB/s | 0 | 0 |
| url_clean | 6.2 | 8321.70 MB/s | 0 | 0 |
| log_line_clean | 9.7 | 34754.80 MB/s | 0 | 0 |
| path_escaped | 64.5 | 667.11 MB/s | 0 | 0 |
| json_in_json | 87.9 | 614.43 MB/s | 0 | 0 |
| prose_with_quotes | 42.1 | 974.36 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 4004.01 MB/s | 0 | 0 |
| unicode_escaped_dense | 295.5 | 649.70 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.2 | 12640.38 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 77.9 | — | 0 | 0 |
| append_empty | 22.8 | — | 0 | 0 |
| replace | 52.3 | — | 0 | 0 |
| create_nested | 49.4 | — | 0 | 0 |
| overwrite_nonobject | 56.3 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 130.8 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 327.8 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 98.8 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 133.7 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2384.0 | 1160.45 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2308.0 | 1198.40 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 257.3 | 723.00 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 98.0 | — | 24 | 1 |
| arena | 80.9 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3823.0 | 6276.52 MB/s | 0 | 0 |
| numberObj/goloop | 1339.0 | 7617.57 MB/s | 0 | 0 |
| nestedMixed/goloop | 2300.0 | 4697.00 MB/s | 0 | 0 |
| stringObj/avx2 | 2019.0 | 11885.76 MB/s | 0 | 0 |
| numberObj/avx2 | 752.5 | 13552.96 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1330.0 | 8122.80 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10162.0 | 2361.45 MB/s | 0 | 0 |
| stringObj/dispatch | 2021.0 | 11874.90 MB/s | 0 | 0 |
| numberObj/current | 6106.0 | 1670.05 MB/s | 0 | 0 |
| numberObj/dispatch | 752.5 | 13552.48 MB/s | 0 | 0 |
| numberArr/current | 205.6 | 32112.72 MB/s | 0 | 0 |
| numberArr/dispatch | 207.4 | 31830.77 MB/s | 0 | 0 |
| nestedMixed/current | 17408.0 | 620.47 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1328.0 | 8136.07 MB/s | 0 | 0 |
