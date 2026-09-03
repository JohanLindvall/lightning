# lightning main-module benchmarks

- generated 2026-09-03T06:38:10Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.1 | 513.66 MB/s | 16 | 1 |
| sentence_clean | 41.3 | 1065.59 MB/s | 48 | 1 |
| url_clean | 39.8 | 1307.66 MB/s | 64 | 1 |
| log_line_clean | 95.4 | 3521.40 MB/s | 352 | 1 |
| path_with_backslash | 116.5 | 317.51 MB/s | 56 | 2 |
| json_in_json | 155.3 | 270.37 MB/s | 72 | 2 |
| prose_with_quotes | 97.6 | 389.18 MB/s | 64 | 2 |
| control_bytes | 119.7 | 200.57 MB/s | 56 | 2 |
| mostly_clean_one_quote | 103.8 | 2939.67 MB/s | 320 | 1 |
| unicode_clean | 277.6 | 850.04 MB/s | 240 | 1 |
| unicode_with_quotes | 157.0 | 401.37 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 113.0 | 2699.66 MB/s | 320 | 1 |
| invalid_utf8_dense | 678.7 | 176.81 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 11.3 | 1418.18 MB/s | 0 | 0 |
| sentence_clean | 21.1 | 2080.49 MB/s | 0 | 0 |
| url_clean | 20.3 | 2563.68 MB/s | 0 | 0 |
| log_line_clean | 28.9 | 11614.82 MB/s | 0 | 0 |
| path_with_backslash | 62.5 | 592.01 MB/s | 0 | 0 |
| json_in_json | 97.6 | 430.21 MB/s | 0 | 0 |
| prose_with_quotes | 40.2 | 944.65 MB/s | 0 | 0 |
| control_bytes | 63.7 | 376.60 MB/s | 0 | 0 |
| mostly_clean_one_quote | 33.0 | 9230.97 MB/s | 0 | 0 |
| unicode_clean | 231.5 | 1019.42 MB/s | 0 | 0 |
| unicode_with_quotes | 97.5 | 646.51 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 52.2 | 5839.74 MB/s | 0 | 0 |
| invalid_utf8_dense | 550.4 | 218.02 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2310.0 | 4346.94 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2240.0 | 4482.51 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2212.0 | 4539.66 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10881.0 | 922.79 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 711.5 | 2545.44 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1278.0 | 1417.00 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.6 | 3469.89 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7797.72 MB/s | 0 | 0 |
| url_clean | 5.7 | 9078.79 MB/s | 0 | 0 |
| log_line_clean | 9.6 | 34920.43 MB/s | 0 | 0 |
| path_escaped | 74.1 | 580.18 MB/s | 48 | 1 |
| json_in_json | 98.5 | 548.47 MB/s | 64 | 1 |
| prose_with_quotes | 59.4 | 690.43 MB/s | 48 | 1 |
| unicode_heavy | 5.1 | 5848.29 MB/s | 0 | 0 |
| unicode_escaped_dense | 288.0 | 666.64 MB/s | 192 | 1 |
| mostly_clean_one_escape | 87.1 | 3513.77 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3028.68 MB/s | 0 | 0 |
| sentence_clean | 6.4 | 6908.19 MB/s | 0 | 0 |
| url_clean | 6.3 | 8192.79 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32906.44 MB/s | 0 | 0 |
| path_escaped | 51.9 | 828.29 MB/s | 0 | 0 |
| json_in_json | 75.7 | 713.54 MB/s | 0 | 0 |
| prose_with_quotes | 38.5 | 1066.01 MB/s | 0 | 0 |
| unicode_heavy | 5.7 | 5299.53 MB/s | 0 | 0 |
| unicode_escaped_dense | 247.3 | 776.29 MB/s | 0 | 0 |
| mostly_clean_one_escape | 23.8 | 12840.45 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 70.3 | — | 0 | 0 |
| append_empty | 21.6 | — | 0 | 0 |
| replace | 45.0 | — | 0 | 0 |
| create_nested | 47.8 | — | 0 | 0 |
| overwrite_nonobject | 53.0 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 127.2 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 315.1 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 97.5 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 127.9 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2189.0 | 1263.64 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2084.0 | 1327.41 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 219.4 | 847.60 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 86.3 | — | 24 | 1 |
| arena | 72.8 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 31656.0 | 620.17 MB/s | 0 | 0 |
| kernel/sep"," | 13086.0 | 1500.26 MB/s | 0 | 0 |
| scalar/sep",_" | 32980.0 | 716.53 MB/s | 0 | 0 |
| kernel/sep",_" | 17675.0 | 1336.99 MB/s | 0 | 0 |
| kernel-only | 13073.0 | 1501.77 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 8546.0 | 945.42 MB/s | 0 | 0 |
| "12," | 10258.0 | 1177.60 MB/s | 0 | 0 |
| "123," | 9651.0 | 1666.08 MB/s | 0 | 0 |
| "1234," | 13101.0 | 1532.66 MB/s | 0 | 0 |
| "123456," | 15999.0 | 1755.14 MB/s | 0 | 0 |
| "1234567," | 11047.0 | 2903.99 MB/s | 0 | 0 |
| "1234,_" | 20457.0 | 1177.08 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 534.0 | — | 0 | 0 |
| canada | 398.5 | — | 0 | 0 |
| mesh | 360.9 | — | 0 | 0 |
| array | 273.1 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 539.3 | — | 0 | 0 |
| canada | 502.5 | — | 0 | 0 |
| mesh | 360.5 | — | 0 | 0 |
| array | 317.3 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4114.0 | 5833.59 MB/s | 0 | 0 |
| numberObj/goloop | 1360.0 | 7497.37 MB/s | 0 | 0 |
| nestedMixed/goloop | 1948.0 | 5543.66 MB/s | 0 | 0 |
| stringObj/avx2 | 2287.0 | 10492.55 MB/s | 0 | 0 |
| numberObj/avx2 | 824.3 | 12371.27 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1299.0 | 8312.94 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 8156.0 | 2942.44 MB/s | 0 | 0 |
| stringObj/dispatch | 2296.0 | 10453.56 MB/s | 0 | 0 |
| numberObj/current | 5848.0 | 1743.87 MB/s | 0 | 0 |
| numberObj/dispatch | 826.1 | 12344.97 MB/s | 0 | 0 |
| numberArr/current | 223.5 | 29537.85 MB/s | 0 | 0 |
| numberArr/dispatch | 225.9 | 29215.75 MB/s | 0 | 0 |
| nestedMixed/current | 16253.0 | 664.54 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1303.0 | 8287.92 MB/s | 0 | 0 |
