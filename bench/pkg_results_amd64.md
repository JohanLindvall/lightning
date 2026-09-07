# lightning main-module benchmarks

- generated 2026-09-07T17:38:26Z
- go version go1.26.5 linux/amd64
- cpu: AMD Ryzen 7 8840HS w/ Radeon 780M Graphics (16 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 181.8 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 26.9 | 595.54 MB/s | 16 | 1 |
| sentence_clean | 39.4 | 1117.44 MB/s | 48 | 1 |
| url_clean | 39.0 | 1334.82 MB/s | 64 | 1 |
| log_line_clean | 82.3 | 4084.57 MB/s | 352 | 1 |
| path_with_backslash | 114.6 | 322.97 MB/s | 56 | 2 |
| json_in_json | 141.3 | 297.26 MB/s | 72 | 2 |
| prose_with_quotes | 88.8 | 427.84 MB/s | 64 | 2 |
| control_bytes | 102.3 | 234.66 MB/s | 56 | 2 |
| mostly_clean_one_quote | 93.5 | 3262.57 MB/s | 320 | 1 |
| unicode_clean | 271.7 | 868.50 MB/s | 240 | 1 |
| unicode_with_quotes | 144.6 | 435.55 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 105.9 | 2879.66 MB/s | 320 | 1 |
| invalid_utf8_dense | 626.8 | 191.44 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.7 | 1838.27 MB/s | 0 | 0 |
| sentence_clean | 16.6 | 2645.88 MB/s | 0 | 0 |
| url_clean | 17.8 | 2927.35 MB/s | 0 | 0 |
| log_line_clean | 23.1 | 14579.65 MB/s | 0 | 0 |
| path_with_backslash | 51.7 | 715.06 MB/s | 0 | 0 |
| json_in_json | 81.2 | 517.50 MB/s | 0 | 0 |
| prose_with_quotes | 32.1 | 1184.74 MB/s | 0 | 0 |
| control_bytes | 49.5 | 484.50 MB/s | 0 | 0 |
| mostly_clean_one_quote | 26.7 | 11415.63 MB/s | 0 | 0 |
| unicode_clean | 210.0 | 1123.92 MB/s | 0 | 0 |
| unicode_with_quotes | 82.8 | 761.28 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 44.1 | 6917.45 MB/s | 0 | 0 |
| invalid_utf8_dense | 453.4 | 264.67 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1648.0 | 6091.34 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1683.0 | 5966.98 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1582.0 | 6346.05 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 8551.0 | 1174.21 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 531.8 | 3405.12 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1027.0 | 1763.28 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.7 | 4344.86 MB/s | 0 | 0 |
| sentence_clean | 4.5 | 9806.89 MB/s | 0 | 0 |
| url_clean | 4.6 | 11403.63 MB/s | 0 | 0 |
| log_line_clean | 7.6 | 44075.06 MB/s | 0 | 0 |
| path_escaped | 72.9 | 589.86 MB/s | 48 | 1 |
| json_in_json | 94.7 | 570.44 MB/s | 64 | 1 |
| prose_with_quotes | 53.2 | 770.20 MB/s | 48 | 1 |
| unicode_heavy | 4.1 | 7272.99 MB/s | 0 | 0 |
| unicode_escaped_dense | 283.9 | 676.38 MB/s | 192 | 1 |
| mostly_clean_one_escape | 76.5 | 3998.94 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3782.26 MB/s | 0 | 0 |
| sentence_clean | 5.0 | 8750.71 MB/s | 0 | 0 |
| url_clean | 5.1 | 10292.51 MB/s | 0 | 0 |
| log_line_clean | 8.2 | 41180.56 MB/s | 0 | 0 |
| path_escaped | 42.4 | 1013.33 MB/s | 0 | 0 |
| json_in_json | 62.3 | 867.35 MB/s | 0 | 0 |
| prose_with_quotes | 31.5 | 1303.36 MB/s | 0 | 0 |
| unicode_heavy | 4.5 | 6634.03 MB/s | 0 | 0 |
| unicode_escaped_dense | 227.5 | 844.03 MB/s | 0 | 0 |
| mostly_clean_one_escape | 19.1 | 15987.90 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 3.7 | — | 0 | 0 |
| 3digit | 4.8 | — | 0 | 0 |
| 5digit | 4.3 | — | 0 | 0 |
| 10digit | 5.5 | — | 0 | 0 |
| 13digit | 5.5 | — | 0 | 0 |
| 16digit | 5.6 | — | 0 | 0 |
| 19digit | 7.4 | — | 0 | 0 |
| neg10digit | 5.5 | — | 0 | 0 |
| 20digit_overflow | 7.0 | — | 0 | 0 |
| notanint | 3.7 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 4.6 | — | 0 | 0 |
| 10digit | 5.0 | — | 0 | 0 |
| 13digit | 5.1 | — | 0 | 0 |
| 20digit | 6.6 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 2.3 | — | 0 | 0 |
| short | 3.6 | — | 0 | 0 |
| medium | 4.1 | — | 0 | 0 |
| long | 6.8 | — | 0 | 0 |
| escaped | 32.0 | — | 8 | 1 |
| notastring | 1.7 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.6 | — | 0 | 0 |
| false | 0.8 | — | 0 | 0 |
| null | 0.8 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 2.0 | — | 0 | 0 |
| number | 2.0 | — | 0 | 0 |
| object | 2.0 | — | 0 | 0 |
| array | 2.0 | — | 0 | 0 |
| null | 2.5 | — | 0 | 0 |
| true | 2.8 | — | 0 | 0 |
| invalid | 2.0 | — | 0 | 0 |
| ws_number | 2.8 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 15.8 | — | 8 | 1 |
| medium | 22.0 | — | 32 | 1 |
| long | 36.6 | — | 128 | 1 |
| escaped | 28.7 | — | 8 | 1 |
| long_late_escape | 50.8 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1561.0 | 1409.64 MB/s | 0 | 0 |
| index | 1463.0 | 1504.59 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1452.0 | 1515.50 MB/s | 0 | 0 |
| strings | 652.9 | 3524.38 MB/s | 0 | 0 |
| records | 739.9 | 3742.44 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 14.9 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 7.7 | — | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 4.2 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 54.4 | — | 0 | 0 |
| append_empty | 17.5 | — | 0 | 0 |
| replace | 36.6 | — | 0 | 0 |
| create_nested | 39.2 | — | 0 | 0 |
| overwrite_nonobject | 41.8 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 99.0 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 248.0 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 75.8 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 102.3 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 379255.0 | 2916.94 MB/s | 2293540 | 25 |
| stream | 126380.0 | 8753.48 MB/s | 66992 | 4 |
| stream_reused | 97516.0 | 11344.39 MB/s | 37 | 1 |
| stream_points | 1265235.0 | 874.35 MB/s | 66992 | 4 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 14297.0 | 1538.91 MB/s | 0 | 0 |
| scalars/stream | 31735.0 | 693.27 MB/s | 66992 | 4 |
| strings/inmemory | 6269.0 | 3669.29 MB/s | 0 | 0 |
| strings/stream | 18613.0 | 1235.75 MB/s | 66992 | 4 |
| records/inmemory | 6806.0 | 4143.33 MB/s | 0 | 0 |
| records/stream | 16352.0 | 1724.41 MB/s | 66992 | 4 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 214885.0 | 4090.21 MB/s | 66992 | 4 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1759.0 | 1572.48 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1709.0 | 1618.12 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 22.2 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 182.7 | 1018.04 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 659.5 | 3488.84 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 618.9 | 2926.11 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 593.1 | 3053.63 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 561.4 | 3225.75 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 746.9 | 3707.21 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1503.0 | 1464.40 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2850.0 | 464.96 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 76.5 | — | 24 | 1 |
| arena | 74.5 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 25021.0 | 784.62 MB/s | 0 | 0 |
| kernel/sep"," | 10487.0 | 1872.10 MB/s | 0 | 0 |
| scalar/sep",_" | 27499.0 | 859.34 MB/s | 0 | 0 |
| kernel/sep",_" | 13981.0 | 1690.22 MB/s | 0 | 0 |
| kernel-only | 10439.0 | 1880.64 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 7109.0 | 1136.57 MB/s | 0 | 0 |
| "12," | 8578.0 | 1408.28 MB/s | 0 | 0 |
| "123," | 8129.0 | 1978.12 MB/s | 0 | 0 |
| "1234," | 10380.0 | 1934.43 MB/s | 0 | 0 |
| "123456," | 12617.0 | 2225.52 MB/s | 0 | 0 |
| "1234567," | 9375.0 | 3421.95 MB/s | 0 | 0 |
| "1234,_" | 16307.0 | 1476.61 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 3.4 | — | 0 | 0 |
| d03 | 4.2 | — | 0 | 0 |
| d05 | 4.1 | — | 0 | 0 |
| d08 | 4.3 | — | 0 | 0 |
| d10 | 5.7 | — | 0 | 0 |
| d13 | 5.7 | — | 0 | 0 |
| d16 | 5.8 | — | 0 | 0 |
| d19 | 7.8 | — | 0 | 0 |
| d20 | 7.7 | — | 0 | 0 |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| strings/whole | 3439.0 | 3257.80 MB/s | 0 | 0 |
| strings/4k | 3451.0 | 3246.79 MB/s | 0 | 0 |
| numbers/whole | 218.8 | 38412.90 MB/s | 0 | 0 |
| numbers/4k | 271.1 | 30992.56 MB/s | 0 | 0 |
| records/whole | 10362.0 | 888.70 MB/s | 0 | 0 |
| records/4k | 10460.0 | 880.44 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 426.2 | — | 0 | 0 |
| canada | 318.5 | — | 0 | 0 |
| mesh | 288.0 | — | 0 | 0 |
| array | 218.9 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 435.6 | — | 0 | 0 |
| canada | 417.0 | — | 0 | 0 |
| mesh | 301.7 | — | 0 | 0 |
| array | 262.0 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3347.0 | 7170.45 MB/s | 0 | 0 |
| numberObj/goloop | 1112.0 | 9170.57 MB/s | 0 | 0 |
| nestedMixed/goloop | 1606.0 | 6725.74 MB/s | 0 | 0 |
| stringObj/avx2 | 1865.0 | 12867.93 MB/s | 0 | 0 |
| numberObj/avx2 | 667.6 | 15276.35 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1059.0 | 10197.27 MB/s | 0 | 0 |
| stringObj/avx512 | 1277.0 | 18791.88 MB/s | 0 | 0 |
| numberObj/avx512 | 468.8 | 21752.80 MB/s | 0 | 0 |
| nestedMixed/avx512 | 861.7 | 12534.88 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 6487.0 | 3699.28 MB/s | 0 | 0 |
| stringObj/dispatch | 1255.0 | 19127.34 MB/s | 0 | 0 |
| numberObj/current | 4800.0 | 2124.70 MB/s | 0 | 0 |
| numberObj/dispatch | 467.2 | 21826.71 MB/s | 0 | 0 |
| numberArr/current | 175.9 | 37536.52 MB/s | 0 | 0 |
| numberArr/dispatch | 183.9 | 35896.24 MB/s | 0 | 0 |
| nestedMixed/current | 13377.0 | 807.40 MB/s | 0 | 0 |
| nestedMixed/dispatch | 870.9 | 12402.40 MB/s | 0 | 0 |
