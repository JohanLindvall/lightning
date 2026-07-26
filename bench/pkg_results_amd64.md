# lightning main-module benchmarks

- generated 2026-07-26T15:04:13Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 23.3 | 686.38 MB/s | 16 | 1 |
| sentence_clean | 32.0 | 1375.49 MB/s | 48 | 1 |
| url_clean | 30.2 | 1722.70 MB/s | 64 | 1 |
| log_line_clean | 75.4 | 4453.94 MB/s | 352 | 1 |
| path_with_backslash | 95.7 | 386.79 MB/s | 56 | 2 |
| json_in_json | 129.1 | 325.25 MB/s | 72 | 2 |
| prose_with_quotes | 78.0 | 487.09 MB/s | 64 | 2 |
| control_bytes | 92.3 | 259.96 MB/s | 56 | 2 |
| mostly_clean_one_quote | 83.2 | 3667.72 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 9.8 | 1625.70 MB/s | 0 | 0 |
| sentence_clean | 18.9 | 2332.72 MB/s | 0 | 0 |
| url_clean | 16.3 | 3186.40 MB/s | 0 | 0 |
| log_line_clean | 22.0 | 15266.14 MB/s | 0 | 0 |
| path_with_backslash | 51.9 | 713.51 MB/s | 0 | 0 |
| json_in_json | 82.9 | 506.59 MB/s | 0 | 0 |
| prose_with_quotes | 34.1 | 1114.56 MB/s | 0 | 0 |
| control_bytes | 49.2 | 487.58 MB/s | 0 | 0 |
| mostly_clean_one_quote | 26.9 | 11356.22 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 602.5 | 3005.59 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1040.0 | 1741.22 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.8 | 4184.64 MB/s | 0 | 0 |
| sentence_clean | 4.6 | 9472.39 MB/s | 0 | 0 |
| url_clean | 4.7 | 11032.23 MB/s | 0 | 0 |
| log_line_clean | 7.7 | 43815.74 MB/s | 0 | 0 |
| path_escaped | 63.3 | 679.16 MB/s | 48 | 1 |
| json_in_json | 82.0 | 658.20 MB/s | 64 | 1 |
| prose_with_quotes | 51.4 | 798.45 MB/s | 48 | 1 |
| unicode_heavy | 4.1 | 7317.57 MB/s | 0 | 0 |
| unicode_escaped_dense | 300.1 | 639.74 MB/s | 192 | 1 |
| mostly_clean_one_escape | 69.3 | 4413.46 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.1 | 3903.40 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 8951.35 MB/s | 0 | 0 |
| url_clean | 4.9 | 10581.54 MB/s | 0 | 0 |
| log_line_clean | 7.9 | 42429.51 MB/s | 0 | 0 |
| path_escaped | 47.3 | 909.50 MB/s | 0 | 0 |
| json_in_json | 67.1 | 804.92 MB/s | 0 | 0 |
| prose_with_quotes | 36.5 | 1123.26 MB/s | 0 | 0 |
| unicode_heavy | 4.5 | 6605.21 MB/s | 0 | 0 |
| unicode_escaped_dense | 270.0 | 711.13 MB/s | 0 | 0 |
| mostly_clean_one_escape | 18.7 | 16364.22 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 60.3 | — | 0 | 0 |
| append_empty | 17.4 | — | 0 | 0 |
| replace | 40.1 | — | 0 | 0 |
| create_nested | 38.5 | — | 0 | 0 |
| overwrite_nonobject | 44.5 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 95.5 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 238.3 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1743.0 | 1587.16 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1700.0 | 1626.75 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 203.7 | 913.06 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3154.0 | 7608.02 MB/s | 0 | 0 |
| numberObj/goloop | 1047.0 | 9736.05 MB/s | 0 | 0 |
| nestedMixed/goloop | 1472.0 | 7335.73 MB/s | 0 | 0 |
| stringObj/avx2 | 1756.0 | 13664.16 MB/s | 0 | 0 |
| numberObj/avx2 | 631.6 | 16145.42 MB/s | 0 | 0 |
| nestedMixed/avx2 | 979.0 | 11032.49 MB/s | 0 | 0 |
| stringObj/avx512 | 1175.0 | 20426.17 MB/s | 0 | 0 |
| numberObj/avx512 | 429.4 | 23746.83 MB/s | 0 | 0 |
| nestedMixed/avx512 | 811.7 | 13306.64 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 8144.0 | 2946.70 MB/s | 0 | 0 |
| stringObj/dispatch | 1179.0 | 20352.36 MB/s | 0 | 0 |
| numberObj/current | 4542.0 | 2245.25 MB/s | 0 | 0 |
| numberObj/dispatch | 430.9 | 23666.70 MB/s | 0 | 0 |
| numberArr/current | 174.1 | 37908.04 MB/s | 0 | 0 |
| numberArr/dispatch | 175.8 | 37546.86 MB/s | 0 | 0 |
| nestedMixed/current | 13477.0 | 801.45 MB/s | 0 | 0 |
| nestedMixed/dispatch | 813.6 | 13274.83 MB/s | 0 | 0 |
