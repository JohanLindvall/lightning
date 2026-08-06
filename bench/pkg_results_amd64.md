# lightning main-module benchmarks

- generated 2026-08-06T12:41:39Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.9 | 535.22 MB/s | 16 | 1 |
| sentence_clean | 41.0 | 1073.61 MB/s | 48 | 1 |
| url_clean | 38.9 | 1336.95 MB/s | 64 | 1 |
| log_line_clean | 94.4 | 3560.07 MB/s | 352 | 1 |
| path_with_backslash | 122.1 | 303.13 MB/s | 56 | 2 |
| json_in_json | 162.8 | 257.92 MB/s | 72 | 2 |
| prose_with_quotes | 99.5 | 381.92 MB/s | 64 | 2 |
| control_bytes | 117.8 | 203.73 MB/s | 56 | 2 |
| mostly_clean_one_quote | 102.5 | 2976.52 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.7 | 1261.63 MB/s | 0 | 0 |
| sentence_clean | 24.3 | 1807.97 MB/s | 0 | 0 |
| url_clean | 20.9 | 2487.52 MB/s | 0 | 0 |
| log_line_clean | 28.2 | 11902.34 MB/s | 0 | 0 |
| path_with_backslash | 67.4 | 548.84 MB/s | 0 | 0 |
| json_in_json | 107.4 | 391.15 MB/s | 0 | 0 |
| prose_with_quotes | 44.0 | 864.34 MB/s | 0 | 0 |
| control_bytes | 62.9 | 381.44 MB/s | 0 | 0 |
| mostly_clean_one_quote | 34.3 | 8889.91 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2514.0 | 3994.29 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2478.0 | 4052.63 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2383.0 | 4213.35 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12301.0 | 816.29 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 767.4 | 2360.00 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1344.0 | 1347.22 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3244.67 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7338.13 MB/s | 0 | 0 |
| url_clean | 6.0 | 8683.63 MB/s | 0 | 0 |
| log_line_clean | 9.9 | 34062.37 MB/s | 0 | 0 |
| path_escaped | 81.9 | 525.11 MB/s | 48 | 1 |
| json_in_json | 107.2 | 503.94 MB/s | 64 | 1 |
| prose_with_quotes | 66.4 | 617.46 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5672.30 MB/s | 0 | 0 |
| unicode_escaped_dense | 309.4 | 620.56 MB/s | 192 | 1 |
| mostly_clean_one_escape | 87.3 | 3506.30 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3027.38 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 6934.62 MB/s | 0 | 0 |
| url_clean | 6.3 | 8206.95 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32907.18 MB/s | 0 | 0 |
| path_escaped | 62.5 | 688.30 MB/s | 0 | 0 |
| json_in_json | 88.8 | 607.76 MB/s | 0 | 0 |
| prose_with_quotes | 47.8 | 858.66 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5141.16 MB/s | 0 | 0 |
| unicode_escaped_dense | 270.4 | 710.11 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.4 | 12546.22 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 75.7 | — | 0 | 0 |
| append_empty | 22.3 | — | 0 | 0 |
| replace | 50.1 | — | 0 | 0 |
| create_nested | 49.4 | — | 0 | 0 |
| overwrite_nonobject | 54.6 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 123.0 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 318.3 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 96.1 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 130.5 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2196.0 | 1259.59 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2151.0 | 1286.03 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 261.4 | 711.48 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 81.6 | — | 24 | 1 |
| arena | 70.1 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4093.0 | 5863.76 MB/s | 0 | 0 |
| numberObj/goloop | 1361.0 | 7492.52 MB/s | 0 | 0 |
| nestedMixed/goloop | 1967.0 | 5492.47 MB/s | 0 | 0 |
| stringObj/avx2 | 2284.0 | 10507.07 MB/s | 0 | 0 |
| numberObj/avx2 | 825.8 | 12349.41 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1286.0 | 8400.38 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10411.0 | 2305.00 MB/s | 0 | 0 |
| stringObj/dispatch | 2274.0 | 10552.73 MB/s | 0 | 0 |
| numberObj/current | 5649.0 | 1805.35 MB/s | 0 | 0 |
| numberObj/dispatch | 822.5 | 12399.06 MB/s | 0 | 0 |
| numberArr/current | 224.2 | 29439.07 MB/s | 0 | 0 |
| numberArr/dispatch | 226.4 | 29159.50 MB/s | 0 | 0 |
| nestedMixed/current | 15929.0 | 678.09 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1289.0 | 8376.35 MB/s | 0 | 0 |
