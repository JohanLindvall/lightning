# lightning main-module benchmarks

- generated 2026-09-07T18:59:37Z
- go version go1.26.7 linux/amd64
- cpu: Intel(R) Xeon(R) 6973P-C (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 156.0 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 19.1 | 836.91 MB/s | 16 | 1 |
| sentence_clean | 30.2 | 1457.73 MB/s | 48 | 1 |
| url_clean | 29.8 | 1744.53 MB/s | 64 | 1 |
| log_line_clean | 95.5 | 3518.68 MB/s | 352 | 1 |
| path_with_backslash | 85.7 | 431.72 MB/s | 56 | 2 |
| json_in_json | 127.2 | 330.10 MB/s | 72 | 2 |
| prose_with_quotes | 70.2 | 541.54 MB/s | 64 | 2 |
| control_bytes | 88.3 | 271.91 MB/s | 56 | 2 |
| mostly_clean_one_quote | 97.2 | 3137.62 MB/s | 320 | 1 |
| unicode_clean | 222.9 | 1058.64 MB/s | 240 | 1 |
| unicode_with_quotes | 127.2 | 495.32 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 101.9 | 2992.93 MB/s | 320 | 1 |
| invalid_utf8_dense | 512.1 | 234.32 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.4 | 2153.25 MB/s | 0 | 0 |
| sentence_clean | 13.7 | 3220.72 MB/s | 0 | 0 |
| url_clean | 13.9 | 3743.72 MB/s | 0 | 0 |
| log_line_clean | 20.1 | 16722.30 MB/s | 0 | 0 |
| path_with_backslash | 44.7 | 828.44 MB/s | 0 | 0 |
| json_in_json | 69.6 | 603.42 MB/s | 0 | 0 |
| prose_with_quotes | 28.9 | 1312.62 MB/s | 0 | 0 |
| control_bytes | 42.8 | 560.74 MB/s | 0 | 0 |
| mostly_clean_one_quote | 23.0 | 13280.88 MB/s | 0 | 0 |
| unicode_clean | 177.9 | 1326.36 MB/s | 0 | 0 |
| unicode_with_quotes | 66.5 | 947.30 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 32.9 | 9278.42 MB/s | 0 | 0 |
| invalid_utf8_dense | 374.2 | 320.70 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1300.0 | 7726.70 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1213.0 | 8280.42 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1175.0 | 8546.71 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 7633.0 | 1315.54 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 474.6 | 3815.89 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 873.5 | 2073.37 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.0 | 8082.97 MB/s | 0 | 0 |
| sentence_clean | 3.5 | 12748.55 MB/s | 0 | 0 |
| url_clean | 3.5 | 14876.37 MB/s | 0 | 0 |
| log_line_clean | 8.3 | 40703.89 MB/s | 0 | 0 |
| path_escaped | 57.3 | 749.89 MB/s | 48 | 1 |
| json_in_json | 82.7 | 653.36 MB/s | 64 | 1 |
| prose_with_quotes | 50.3 | 815.63 MB/s | 48 | 1 |
| unicode_heavy | 2.7 | 11038.70 MB/s | 0 | 0 |
| unicode_escaped_dense | 211.2 | 909.20 MB/s | 192 | 1 |
| mostly_clean_one_escape | 92.0 | 3324.52 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.6 | 6150.17 MB/s | 0 | 0 |
| sentence_clean | 3.9 | 11312.36 MB/s | 0 | 0 |
| url_clean | 3.8 | 13555.01 MB/s | 0 | 0 |
| log_line_clean | 6.8 | 49197.13 MB/s | 0 | 0 |
| path_escaped | 38.6 | 1113.32 MB/s | 0 | 0 |
| json_in_json | 55.6 | 970.53 MB/s | 0 | 0 |
| prose_with_quotes | 27.8 | 1475.29 MB/s | 0 | 0 |
| unicode_heavy | 2.7 | 11024.80 MB/s | 0 | 0 |
| unicode_escaped_dense | 166.9 | 1150.67 MB/s | 0 | 0 |
| mostly_clean_one_escape | 16.3 | 18815.30 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 2.6 | — | 0 | 0 |
| 3digit | 4.5 | — | 0 | 0 |
| 5digit | 3.2 | — | 0 | 0 |
| 10digit | 4.8 | — | 0 | 0 |
| 13digit | 4.8 | — | 0 | 0 |
| 16digit | 4.6 | — | 0 | 0 |
| 19digit | 5.4 | — | 0 | 0 |
| neg10digit | 4.5 | — | 0 | 0 |
| 20digit_overflow | 5.1 | — | 0 | 0 |
| notanint | 2.1 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 2.8 | — | 0 | 0 |
| 10digit | 4.1 | — | 0 | 0 |
| 13digit | 3.9 | — | 0 | 0 |
| 20digit | 5.1 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 1.6 | — | 0 | 0 |
| short | 2.5 | — | 0 | 0 |
| medium | 3.0 | — | 0 | 0 |
| long | 5.4 | — | 0 | 0 |
| escaped | 22.0 | — | 8 | 1 |
| notastring | 1.2 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.6 | — | 0 | 0 |
| false | 0.7 | — | 0 | 0 |
| null | 0.5 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 1.0 | — | 0 | 0 |
| number | 1.0 | — | 0 | 0 |
| object | 1.0 | — | 0 | 0 |
| array | 1.0 | — | 0 | 0 |
| null | 1.5 | — | 0 | 0 |
| true | 1.6 | — | 0 | 0 |
| invalid | 1.3 | — | 0 | 0 |
| ws_number | 1.6 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 10.8 | — | 8 | 1 |
| medium | 17.3 | — | 32 | 1 |
| long | 37.6 | — | 128 | 1 |
| escaped | 22.2 | — | 8 | 1 |
| long_late_escape | 47.6 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1875.0 | 1173.95 MB/s | 0 | 0 |
| index | 1412.0 | 1559.31 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1579.0 | 1393.59 MB/s | 0 | 0 |
| strings | 582.4 | 3950.70 MB/s | 0 | 0 |
| records | 591.1 | 4684.19 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 13.2 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 8.8 | — | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2.9 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 44.3 | — | 0 | 0 |
| append_empty | 14.9 | — | 0 | 0 |
| replace | 29.1 | — | 0 | 0 |
| create_nested | 31.8 | — | 0 | 0 |
| overwrite_nonobject | 35.8 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 83.0 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 208.3 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 61.8 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 79.5 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 414239.0 | 2670.59 MB/s | 2293536 | 25 |
| stream | 87923.0 | 12582.16 MB/s | 66992 | 4 |
| stream_reused | 75694.0 | 14614.93 MB/s | 33 | 1 |
| stream_points | 984519.0 | 1123.66 MB/s | 66992 | 4 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 16237.0 | 1355.02 MB/s | 0 | 0 |
| scalars/stream | 27005.0 | 814.71 MB/s | 66992 | 4 |
| strings/inmemory | 5241.0 | 4388.57 MB/s | 0 | 0 |
| strings/stream | 17629.0 | 1304.70 MB/s | 66992 | 4 |
| records/inmemory | 5089.0 | 5540.69 MB/s | 0 | 0 |
| records/stream | 15003.0 | 1879.44 MB/s | 66992 | 4 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 148084.0 | 5935.30 MB/s | 66992 | 4 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1523.0 | 1815.86 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1467.0 | 1885.38 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 16.8 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 148.9 | 1249.16 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 514.7 | 4470.74 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 589.4 | 3072.44 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 584.2 | 3099.87 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 571.1 | 3170.81 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 574.8 | 4817.12 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1843.0 | 1194.08 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2479.0 | 534.59 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 58.3 | — | 24 | 1 |
| arena | 53.9 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 22858.0 | 858.88 MB/s | 0 | 0 |
| kernel/sep"," | 8802.0 | 2230.41 MB/s | 0 | 0 |
| scalar/sep",_" | 24382.0 | 969.19 MB/s | 0 | 0 |
| kernel/sep",_" | 14136.0 | 1671.67 MB/s | 0 | 0 |
| kernel-only | 8709.0 | 2254.19 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6051.0 | 1335.31 MB/s | 0 | 0 |
| "12," | 7396.0 | 1633.40 MB/s | 0 | 0 |
| "123," | 6664.0 | 2412.94 MB/s | 0 | 0 |
| "1234," | 8590.0 | 2337.69 MB/s | 0 | 0 |
| "123456," | 10752.0 | 2611.49 MB/s | 0 | 0 |
| "1234567," | 7460.0 | 4300.54 MB/s | 0 | 0 |
| "1234,_" | 14553.0 | 1654.57 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 2.0 | — | 0 | 0 |
| d03 | 2.8 | — | 0 | 0 |
| d05 | 3.0 | — | 0 | 0 |
| d08 | 3.0 | — | 0 | 0 |
| d10 | 4.1 | — | 0 | 0 |
| d13 | 4.1 | — | 0 | 0 |
| d16 | 4.1 | — | 0 | 0 |
| d19 | 5.3 | — | 0 | 0 |
| d20 | 5.4 | — | 0 | 0 |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/whole | 7538.0 | 1221.62 MB/s | 0 | 0 |
| records/4k | 7686.0 | 1198.08 MB/s | 0 | 0 |
| strings/whole | 2466.0 | 4543.07 MB/s | 0 | 0 |
| strings/4k | 2551.0 | 4392.93 MB/s | 0 | 0 |
| numbers/whole | 197.5 | 42540.44 MB/s | 0 | 0 |
| numbers/4k | 251.6 | 33398.14 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 330.4 | — | 0 | 0 |
| canada | 239.8 | — | 0 | 0 |
| mesh | 191.7 | — | 0 | 0 |
| array | 152.3 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 346.3 | — | 0 | 0 |
| canada | 328.7 | — | 0 | 0 |
| mesh | 259.2 | — | 0 | 0 |
| array | 243.8 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 2586.0 | 9278.95 MB/s | 0 | 0 |
| numberObj/goloop | 877.5 | 11621.06 MB/s | 0 | 0 |
| nestedMixed/goloop | 1324.0 | 8158.58 MB/s | 0 | 0 |
| stringObj/avx2 | 1424.0 | 16856.99 MB/s | 0 | 0 |
| numberObj/avx2 | 523.2 | 19491.82 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1159.0 | 9315.26 MB/s | 0 | 0 |
| stringObj/avx512 | 1079.0 | 22241.34 MB/s | 0 | 0 |
| numberObj/avx512 | 308.7 | 33031.04 MB/s | 0 | 0 |
| nestedMixed/avx512 | 1139.0 | 9478.95 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 5753.0 | 4171.46 MB/s | 0 | 0 |
| stringObj/dispatch | 1062.0 | 22589.52 MB/s | 0 | 0 |
| numberObj/current | 3110.0 | 3278.71 MB/s | 0 | 0 |
| numberObj/dispatch | 299.4 | 34061.94 MB/s | 0 | 0 |
| numberArr/current | 170.5 | 38712.47 MB/s | 0 | 0 |
| numberArr/dispatch | 172.2 | 38329.08 MB/s | 0 | 0 |
| nestedMixed/current | 10572.0 | 1021.65 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1155.0 | 9354.92 MB/s | 0 | 0 |
