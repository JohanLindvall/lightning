# lightning main-module benchmarks

- generated 2026-09-07T16:50:27Z
- go version go1.26.5 linux/amd64
- cpu: AMD Ryzen 7 8840HS w/ Radeon 780M Graphics (16 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 126.4 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 20.3 | 788.33 MB/s | 16 | 1 |
| sentence_clean | 31.8 | 1382.33 MB/s | 48 | 1 |
| url_clean | 30.1 | 1724.61 MB/s | 64 | 1 |
| log_line_clean | 69.6 | 4825.98 MB/s | 352 | 1 |
| path_with_backslash | 82.0 | 451.41 MB/s | 56 | 2 |
| json_in_json | 108.6 | 386.59 MB/s | 72 | 2 |
| prose_with_quotes | 67.5 | 562.97 MB/s | 64 | 2 |
| control_bytes | 79.0 | 303.89 MB/s | 56 | 2 |
| mostly_clean_one_quote | 78.6 | 3880.38 MB/s | 320 | 1 |
| unicode_clean | 204.7 | 1153.03 MB/s | 240 | 1 |
| unicode_with_quotes | 111.9 | 563.05 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 86.3 | 3532.28 MB/s | 320 | 1 |
| invalid_utf8_dense | 491.8 | 244.00 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.5 | 2454.65 MB/s | 0 | 0 |
| sentence_clean | 12.4 | 3553.45 MB/s | 0 | 0 |
| url_clean | 12.1 | 4297.38 MB/s | 0 | 0 |
| log_line_clean | 17.7 | 18966.70 MB/s | 0 | 0 |
| path_with_backslash | 39.0 | 947.49 MB/s | 0 | 0 |
| json_in_json | 61.7 | 680.67 MB/s | 0 | 0 |
| prose_with_quotes | 24.0 | 1585.18 MB/s | 0 | 0 |
| control_bytes | 36.9 | 649.76 MB/s | 0 | 0 |
| mostly_clean_one_quote | 20.1 | 15178.09 MB/s | 0 | 0 |
| unicode_clean | 141.1 | 1672.48 MB/s | 0 | 0 |
| unicode_with_quotes | 59.0 | 1068.16 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 32.1 | 9487.79 MB/s | 0 | 0 |
| invalid_utf8_dense | 346.7 | 346.15 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1243.0 | 8075.93 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1262.0 | 7956.29 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1217.0 | 8253.74 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 6728.0 | 1492.33 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 426.3 | 4248.59 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 782.8 | 2313.58 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.9 | 5612.43 MB/s | 0 | 0 |
| sentence_clean | 3.5 | 12658.68 MB/s | 0 | 0 |
| url_clean | 3.5 | 14932.34 MB/s | 0 | 0 |
| log_line_clean | 5.8 | 58064.77 MB/s | 0 | 0 |
| path_escaped | 51.8 | 829.96 MB/s | 48 | 1 |
| json_in_json | 69.8 | 774.07 MB/s | 64 | 1 |
| prose_with_quotes | 42.6 | 961.24 MB/s | 48 | 1 |
| unicode_heavy | 3.1 | 9538.54 MB/s | 0 | 0 |
| unicode_escaped_dense | 202.5 | 948.36 MB/s | 192 | 1 |
| mostly_clean_one_escape | 63.6 | 4807.73 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.3 | 4910.73 MB/s | 0 | 0 |
| sentence_clean | 3.9 | 11271.04 MB/s | 0 | 0 |
| url_clean | 3.9 | 13311.76 MB/s | 0 | 0 |
| log_line_clean | 6.1 | 54672.23 MB/s | 0 | 0 |
| path_escaped | 31.1 | 1380.88 MB/s | 0 | 0 |
| json_in_json | 46.0 | 1175.09 MB/s | 0 | 0 |
| prose_with_quotes | 23.5 | 1744.15 MB/s | 0 | 0 |
| unicode_heavy | 3.5 | 8475.75 MB/s | 0 | 0 |
| unicode_escaped_dense | 149.2 | 1286.45 MB/s | 0 | 0 |
| mostly_clean_one_escape | 14.6 | 21002.64 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 2.5 | — | 0 | 0 |
| 3digit | 3.1 | — | 0 | 0 |
| 5digit | 3.2 | — | 0 | 0 |
| 10digit | 4.1 | — | 0 | 0 |
| 13digit | 4.2 | — | 0 | 0 |
| 16digit | 4.2 | — | 0 | 0 |
| 19digit | 5.5 | — | 0 | 0 |
| neg10digit | 4.1 | — | 0 | 0 |
| 20digit_overflow | 5.2 | — | 0 | 0 |
| notanint | 2.6 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 3.0 | — | 0 | 0 |
| 10digit | 3.8 | — | 0 | 0 |
| 13digit | 3.8 | — | 0 | 0 |
| 20digit | 4.9 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 1.7 | — | 0 | 0 |
| short | 2.8 | — | 0 | 0 |
| medium | 3.1 | — | 0 | 0 |
| long | 5.2 | — | 0 | 0 |
| escaped | 22.5 | — | 8 | 1 |
| notastring | 1.3 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.5 | — | 0 | 0 |
| false | 0.6 | — | 0 | 0 |
| null | 0.6 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 1.7 | — | 0 | 0 |
| number | 1.5 | — | 0 | 0 |
| object | 1.5 | — | 0 | 0 |
| array | 1.5 | — | 0 | 0 |
| null | 1.9 | — | 0 | 0 |
| true | 2.1 | — | 0 | 0 |
| invalid | 1.5 | — | 0 | 0 |
| ws_number | 2.1 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 11.6 | — | 8 | 1 |
| medium | 18.2 | — | 32 | 1 |
| long | 30.5 | — | 128 | 1 |
| escaped | 20.9 | — | 8 | 1 |
| long_late_escape | 42.6 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1176.0 | 1871.85 MB/s | 0 | 0 |
| index | 1092.0 | 2015.32 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1098.0 | 2004.54 MB/s | 0 | 0 |
| strings | 512.6 | 4489.22 MB/s | 0 | 0 |
| records | 553.4 | 5003.55 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 11.4 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 5.7 | — | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3.1 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 42.0 | — | 0 | 0 |
| append_empty | 13.3 | — | 0 | 0 |
| replace | 27.5 | — | 0 | 0 |
| create_nested | 29.9 | — | 0 | 0 |
| overwrite_nonobject | 31.6 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 75.9 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 190.9 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 59.0 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 78.0 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1324.0 | 2089.69 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1296.0 | 2133.88 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 17.7 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 134.5 | 1383.02 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 525.4 | 4379.72 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 466.8 | 3879.93 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 464.2 | 3901.40 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 445.2 | 4068.20 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 548.4 | 5049.29 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1135.0 | 1939.88 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2037.0 | 650.57 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 57.0 | — | 24 | 1 |
| arena | 55.8 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 20044.0 | 979.45 MB/s | 0 | 0 |
| kernel/sep"," | 7796.0 | 2518.33 MB/s | 0 | 0 |
| scalar/sep",_" | 21581.0 | 1094.98 MB/s | 0 | 0 |
| kernel/sep",_" | 10483.0 | 2254.26 MB/s | 0 | 0 |
| kernel-only | 7786.0 | 2521.32 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 5224.0 | 1546.64 MB/s | 0 | 0 |
| "12," | 6155.0 | 1962.67 MB/s | 0 | 0 |
| "123," | 5903.0 | 2724.20 MB/s | 0 | 0 |
| "1234," | 7728.0 | 2598.39 MB/s | 0 | 0 |
| "123456," | 9554.0 | 2939.02 MB/s | 0 | 0 |
| "1234567," | 6631.0 | 4838.04 MB/s | 0 | 0 |
| "1234,_" | 12016.0 | 2003.92 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 2.5 | — | 0 | 0 |
| d03 | 3.0 | — | 0 | 0 |
| d05 | 3.1 | — | 0 | 0 |
| d08 | 3.1 | — | 0 | 0 |
| d10 | 4.3 | — | 0 | 0 |
| d13 | 4.3 | — | 0 | 0 |
| d16 | 4.2 | — | 0 | 0 |
| d19 | 5.5 | — | 0 | 0 |
| d20 | 5.5 | — | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 317.1 | — | 0 | 0 |
| canada | 236.1 | — | 0 | 0 |
| mesh | 209.3 | — | 0 | 0 |
| array | 164.7 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 324.0 | — | 0 | 0 |
| canada | 303.8 | — | 0 | 0 |
| mesh | 220.2 | — | 0 | 0 |
| array | 192.5 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 2469.0 | 9720.35 MB/s | 0 | 0 |
| numberObj/goloop | 824.0 | 12376.08 MB/s | 0 | 0 |
| nestedMixed/goloop | 1177.0 | 9176.02 MB/s | 0 | 0 |
| stringObj/avx2 | 1345.0 | 17848.50 MB/s | 0 | 0 |
| numberObj/avx2 | 488.3 | 20885.07 MB/s | 0 | 0 |
| nestedMixed/avx2 | 774.6 | 13944.55 MB/s | 0 | 0 |
| stringObj/avx512 | 925.2 | 25937.65 MB/s | 0 | 0 |
| numberObj/avx512 | 337.3 | 30230.41 MB/s | 0 | 0 |
| nestedMixed/avx512 | 604.1 | 17879.43 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 4963.0 | 4835.05 MB/s | 0 | 0 |
| stringObj/dispatch | 928.4 | 25848.88 MB/s | 0 | 0 |
| numberObj/current | 3384.0 | 3013.23 MB/s | 0 | 0 |
| numberObj/dispatch | 341.9 | 29823.24 MB/s | 0 | 0 |
| numberArr/current | 134.7 | 49023.14 MB/s | 0 | 0 |
| numberArr/dispatch | 135.9 | 48561.75 MB/s | 0 | 0 |
| nestedMixed/current | 10134.0 | 1065.82 MB/s | 0 | 0 |
| nestedMixed/dispatch | 605.4 | 17841.31 MB/s | 0 | 0 |
