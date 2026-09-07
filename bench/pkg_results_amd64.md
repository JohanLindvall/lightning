# lightning main-module benchmarks

- generated 2026-09-07T16:39:46Z
- go version go1.26.5 linux/amd64
- cpu: AMD Ryzen 7 8840HS w/ Radeon 780M Graphics (16 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 133.8 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 20.9 | 765.86 MB/s | 16 | 1 |
| sentence_clean | 30.6 | 1437.42 MB/s | 48 | 1 |
| url_clean | 29.8 | 1744.00 MB/s | 64 | 1 |
| log_line_clean | 66.6 | 5048.10 MB/s | 352 | 1 |
| path_with_backslash | 81.0 | 457.03 MB/s | 56 | 2 |
| json_in_json | 108.4 | 387.59 MB/s | 72 | 2 |
| prose_with_quotes | 67.9 | 559.96 MB/s | 64 | 2 |
| control_bytes | 80.1 | 299.79 MB/s | 56 | 2 |
| mostly_clean_one_quote | 77.5 | 3933.96 MB/s | 320 | 1 |
| unicode_clean | 198.0 | 1191.74 MB/s | 240 | 1 |
| unicode_with_quotes | 109.8 | 573.80 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 86.5 | 3527.45 MB/s | 320 | 1 |
| invalid_utf8_dense | 491.1 | 244.37 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.6 | 2427.89 MB/s | 0 | 0 |
| sentence_clean | 12.5 | 3510.49 MB/s | 0 | 0 |
| url_clean | 12.1 | 4285.03 MB/s | 0 | 0 |
| log_line_clean | 17.4 | 19334.96 MB/s | 0 | 0 |
| path_with_backslash | 39.6 | 934.69 MB/s | 0 | 0 |
| json_in_json | 60.7 | 692.21 MB/s | 0 | 0 |
| prose_with_quotes | 24.0 | 1584.58 MB/s | 0 | 0 |
| control_bytes | 36.9 | 651.16 MB/s | 0 | 0 |
| mostly_clean_one_quote | 20.2 | 15069.09 MB/s | 0 | 0 |
| unicode_clean | 140.8 | 1676.29 MB/s | 0 | 0 |
| unicode_with_quotes | 58.3 | 1080.30 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 31.8 | 9591.36 MB/s | 0 | 0 |
| invalid_utf8_dense | 347.1 | 345.69 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1232.0 | 8148.22 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1271.0 | 7901.97 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1207.0 | 8319.71 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 6580.0 | 1526.00 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 420.8 | 4303.23 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 785.8 | 2304.80 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.8 | 5663.88 MB/s | 0 | 0 |
| sentence_clean | 3.5 | 12616.38 MB/s | 0 | 0 |
| url_clean | 3.5 | 15061.28 MB/s | 0 | 0 |
| log_line_clean | 5.8 | 58181.29 MB/s | 0 | 0 |
| path_escaped | 51.3 | 838.03 MB/s | 48 | 1 |
| json_in_json | 67.5 | 800.54 MB/s | 64 | 1 |
| prose_with_quotes | 43.1 | 951.89 MB/s | 48 | 1 |
| unicode_heavy | 3.2 | 9458.25 MB/s | 0 | 0 |
| unicode_escaped_dense | 203.8 | 942.06 MB/s | 192 | 1 |
| mostly_clean_one_escape | 62.3 | 4911.52 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.2 | 4969.71 MB/s | 0 | 0 |
| sentence_clean | 4.0 | 10985.99 MB/s | 0 | 0 |
| url_clean | 8.4 | 6184.67 MB/s | 0 | 0 |
| log_line_clean | 7.0 | 48125.31 MB/s | 0 | 0 |
| path_escaped | 33.8 | 1273.77 MB/s | 0 | 0 |
| json_in_json | 46.8 | 1153.87 MB/s | 0 | 0 |
| prose_with_quotes | 23.9 | 1714.97 MB/s | 0 | 0 |
| unicode_heavy | 3.8 | 7880.33 MB/s | 0 | 0 |
| unicode_escaped_dense | 150.0 | 1280.04 MB/s | 0 | 0 |
| mostly_clean_one_escape | 15.0 | 20381.38 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 2.5 | — | 0 | 0 |
| 3digit | 3.2 | — | 0 | 0 |
| 5digit | 3.2 | — | 0 | 0 |
| 10digit | 4.2 | — | 0 | 0 |
| 13digit | 4.2 | — | 0 | 0 |
| 16digit | 4.2 | — | 0 | 0 |
| 19digit | 5.5 | — | 0 | 0 |
| neg10digit | 4.1 | — | 0 | 0 |
| 20digit_overflow | 5.3 | — | 0 | 0 |
| notanint | 2.5 | — | 0 | 0 |

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
| empty | 1.8 | — | 0 | 0 |
| short | 2.7 | — | 0 | 0 |
| medium | 3.1 | — | 0 | 0 |
| long | 5.2 | — | 0 | 0 |
| escaped | 22.2 | — | 8 | 1 |
| notastring | 1.7 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.5 | — | 0 | 0 |
| false | 0.6 | — | 0 | 0 |
| null | 0.8 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 1.5 | — | 0 | 0 |
| number | 1.6 | — | 0 | 0 |
| object | 1.5 | — | 0 | 0 |
| array | 1.7 | — | 0 | 0 |
| null | 1.9 | — | 0 | 0 |
| true | 2.1 | — | 0 | 0 |
| invalid | 1.7 | — | 0 | 0 |
| ws_number | 2.9 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 11.5 | — | 8 | 1 |
| medium | 17.9 | — | 32 | 1 |
| long | 29.0 | — | 128 | 1 |
| escaped | 21.0 | — | 8 | 1 |
| long_late_escape | 41.0 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1148.0 | 1916.82 MB/s | 0 | 0 |
| index | 1542.0 | 1427.11 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 11.6 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 5.6 | — | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3.1 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 41.5 | — | 0 | 0 |
| append_empty | 13.3 | — | 0 | 0 |
| replace | 27.7 | — | 0 | 0 |
| create_nested | 28.2 | — | 0 | 0 |
| overwrite_nonobject | 31.3 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 75.0 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 189.8 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 60.0 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 78.5 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1302.0 | 2124.64 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1254.0 | 2205.95 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 17.5 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 134.6 | 1382.29 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 530.2 | 4339.93 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 467.8 | 3870.93 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 460.9 | 3929.37 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 447.2 | 4049.26 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 548.2 | 5050.93 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1140.0 | 1931.18 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2038.0 | 650.06 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 56.8 | — | 24 | 1 |
| arena | 56.3 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 19201.0 | 1022.44 MB/s | 0 | 0 |
| kernel/sep"," | 7860.0 | 2497.74 MB/s | 0 | 0 |
| scalar/sep",_" | 21777.0 | 1085.15 MB/s | 0 | 0 |
| kernel/sep",_" | 10492.0 | 2252.39 MB/s | 0 | 0 |
| kernel-only | 7785.0 | 2521.73 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 5151.0 | 1568.50 MB/s | 0 | 0 |
| "12," | 6180.0 | 1954.75 MB/s | 0 | 0 |
| "123," | 5784.0 | 2779.96 MB/s | 0 | 0 |
| "1234," | 7733.0 | 2596.79 MB/s | 0 | 0 |
| "123456," | 9528.0 | 2947.16 MB/s | 0 | 0 |
| "1234567," | 6688.0 | 4796.35 MB/s | 0 | 0 |
| "1234,_" | 12054.0 | 1997.55 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 2.5 | — | 0 | 0 |
| d03 | 3.1 | — | 0 | 0 |
| d05 | 3.1 | — | 0 | 0 |
| d08 | 3.1 | — | 0 | 0 |
| d10 | 4.2 | — | 0 | 0 |
| d13 | 4.3 | — | 0 | 0 |
| d16 | 4.2 | — | 0 | 0 |
| d19 | 5.5 | — | 0 | 0 |
| d20 | 5.5 | — | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 317.6 | — | 0 | 0 |
| canada | 236.5 | — | 0 | 0 |
| mesh | 212.9 | — | 0 | 0 |
| array | 166.6 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 324.8 | — | 0 | 0 |
| canada | 303.5 | — | 0 | 0 |
| mesh | 216.4 | — | 0 | 0 |
| array | 190.4 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 2446.0 | 9810.13 MB/s | 0 | 0 |
| numberObj/goloop | 827.4 | 12324.96 MB/s | 0 | 0 |
| nestedMixed/goloop | 1179.0 | 9163.46 MB/s | 0 | 0 |
| stringObj/avx2 | 1347.0 | 17814.65 MB/s | 0 | 0 |
| numberObj/avx2 | 488.4 | 20878.70 MB/s | 0 | 0 |
| nestedMixed/avx2 | 768.2 | 14060.44 MB/s | 0 | 0 |
| stringObj/avx512 | 928.0 | 25860.69 MB/s | 0 | 0 |
| numberObj/avx512 | 342.0 | 29822.30 MB/s | 0 | 0 |
| nestedMixed/avx512 | 608.2 | 17759.76 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 4944.0 | 4854.10 MB/s | 0 | 0 |
| stringObj/dispatch | 927.2 | 25881.18 MB/s | 0 | 0 |
| numberObj/current | 3491.0 | 2920.93 MB/s | 0 | 0 |
| numberObj/dispatch | 342.6 | 29762.96 MB/s | 0 | 0 |
| numberArr/current | 135.8 | 48615.41 MB/s | 0 | 0 |
| numberArr/dispatch | 135.3 | 48774.03 MB/s | 0 | 0 |
| nestedMixed/current | 10393.0 | 1039.27 MB/s | 0 | 0 |
| nestedMixed/dispatch | 610.7 | 17687.16 MB/s | 0 | 0 |
