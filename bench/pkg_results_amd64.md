# lightning main-module benchmarks

- generated 2026-09-03T00:44:48Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.0 | 516.31 MB/s | 16 | 1 |
| sentence_clean | 40.3 | 1090.87 MB/s | 48 | 1 |
| url_clean | 40.6 | 1281.99 MB/s | 64 | 1 |
| log_line_clean | 93.7 | 3585.78 MB/s | 352 | 1 |
| path_with_backslash | 116.4 | 317.85 MB/s | 56 | 2 |
| json_in_json | 154.2 | 272.37 MB/s | 72 | 2 |
| prose_with_quotes | 96.6 | 393.29 MB/s | 64 | 2 |
| control_bytes | 119.6 | 200.62 MB/s | 56 | 2 |
| mostly_clean_one_quote | 101.8 | 2995.57 MB/s | 320 | 1 |
| unicode_clean | 276.9 | 852.19 MB/s | 240 | 1 |
| unicode_with_quotes | 157.8 | 399.28 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 113.9 | 2677.21 MB/s | 320 | 1 |
| invalid_utf8_dense | 672.9 | 178.33 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 11.3 | 1418.54 MB/s | 0 | 0 |
| sentence_clean | 21.2 | 2076.30 MB/s | 0 | 0 |
| url_clean | 20.3 | 2558.69 MB/s | 0 | 0 |
| log_line_clean | 28.9 | 11621.30 MB/s | 0 | 0 |
| path_with_backslash | 62.6 | 591.35 MB/s | 0 | 0 |
| json_in_json | 97.5 | 430.90 MB/s | 0 | 0 |
| prose_with_quotes | 40.2 | 944.86 MB/s | 0 | 0 |
| control_bytes | 63.2 | 379.64 MB/s | 0 | 0 |
| mostly_clean_one_quote | 33.0 | 9241.44 MB/s | 0 | 0 |
| unicode_clean | 232.3 | 1015.75 MB/s | 0 | 0 |
| unicode_with_quotes | 97.9 | 643.32 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 52.3 | 5834.67 MB/s | 0 | 0 |
| invalid_utf8_dense | 545.9 | 219.81 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2319.0 | 4329.48 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2239.0 | 4485.16 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2225.0 | 4512.36 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10788.0 | 930.72 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 709.3 | 2553.11 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1277.0 | 1418.66 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.6 | 3478.63 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7801.53 MB/s | 0 | 0 |
| url_clean | 12.4 | 4184.53 MB/s | 0 | 0 |
| log_line_clean | 9.5 | 35328.87 MB/s | 0 | 0 |
| path_escaped | 73.0 | 588.66 MB/s | 48 | 1 |
| json_in_json | 98.0 | 551.11 MB/s | 64 | 1 |
| prose_with_quotes | 59.8 | 686.22 MB/s | 48 | 1 |
| unicode_heavy | 5.1 | 5868.37 MB/s | 0 | 0 |
| unicode_escaped_dense | 287.3 | 668.38 MB/s | 192 | 1 |
| mostly_clean_one_escape | 85.1 | 3595.16 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.4 | 2980.69 MB/s | 0 | 0 |
| sentence_clean | 6.4 | 6915.77 MB/s | 0 | 0 |
| url_clean | 6.3 | 8196.56 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32854.39 MB/s | 0 | 0 |
| path_escaped | 52.7 | 815.66 MB/s | 0 | 0 |
| json_in_json | 76.3 | 707.81 MB/s | 0 | 0 |
| prose_with_quotes | 40.2 | 1018.65 MB/s | 0 | 0 |
| unicode_heavy | 5.7 | 5218.91 MB/s | 0 | 0 |
| unicode_escaped_dense | 248.8 | 771.62 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.1 | 12704.00 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 69.8 | — | 0 | 0 |
| append_empty | 22.0 | — | 0 | 0 |
| replace | 45.6 | — | 0 | 0 |
| create_nested | 48.6 | — | 0 | 0 |
| overwrite_nonobject | 53.0 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 129.9 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 313.0 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 99.5 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 128.3 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2173.0 | 1273.10 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2083.0 | 1327.75 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 219.5 | 847.39 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 86.0 | — | 24 | 1 |
| arena | 73.0 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 31667.0 | 619.96 MB/s | 0 | 0 |
| kernel/sep"," | 13080.0 | 1500.90 MB/s | 0 | 0 |
| scalar/sep",_" | 32923.0 | 717.78 MB/s | 0 | 0 |
| kernel/sep",_" | 17663.0 | 1337.87 MB/s | 0 | 0 |
| kernel-only | 13092.0 | 1499.58 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 8551.0 | 944.93 MB/s | 0 | 0 |
| "12," | 10271.0 | 1176.11 MB/s | 0 | 0 |
| "123," | 9657.0 | 1665.08 MB/s | 0 | 0 |
| "1234," | 13080.0 | 1535.13 MB/s | 0 | 0 |
| "123456," | 16094.0 | 1744.77 MB/s | 0 | 0 |
| "1234567," | 11051.0 | 2902.78 MB/s | 0 | 0 |
| "1234,_" | 20457.0 | 1177.03 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 534.2 | — | 0 | 0 |
| canada | 399.4 | — | 0 | 0 |
| mesh | 359.1 | — | 0 | 0 |
| array | 287.2 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 539.1 | — | 0 | 0 |
| canada | 501.1 | — | 0 | 0 |
| mesh | 360.4 | — | 0 | 0 |
| array | 316.5 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4083.0 | 5877.09 MB/s | 0 | 0 |
| numberObj/goloop | 1388.0 | 7346.02 MB/s | 0 | 0 |
| nestedMixed/goloop | 1950.0 | 5539.01 MB/s | 0 | 0 |
| stringObj/avx2 | 2289.0 | 10486.26 MB/s | 0 | 0 |
| numberObj/avx2 | 824.3 | 12371.82 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1297.0 | 8327.70 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 8117.0 | 2956.52 MB/s | 0 | 0 |
| stringObj/dispatch | 2297.0 | 10448.76 MB/s | 0 | 0 |
| numberObj/current | 5328.0 | 1914.15 MB/s | 0 | 0 |
| numberObj/dispatch | 825.9 | 12347.00 MB/s | 0 | 0 |
| numberArr/current | 223.6 | 29522.81 MB/s | 0 | 0 |
| numberArr/dispatch | 226.1 | 29188.61 MB/s | 0 | 0 |
| nestedMixed/current | 16876.0 | 640.03 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1303.0 | 8290.41 MB/s | 0 | 0 |
