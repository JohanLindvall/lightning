# lightning main-module benchmarks

- generated 2026-09-08T16:42:04Z
- go version go1.26.8 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 200.1 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.1 | 531.86 MB/s | 16 | 1 |
| sentence_clean | 37.8 | 1163.70 MB/s | 48 | 1 |
| url_clean | 41.7 | 1246.14 MB/s | 64 | 1 |
| log_line_clean | 122.7 | 2738.65 MB/s | 352 | 1 |
| path_with_backslash | 118.6 | 311.92 MB/s | 56 | 2 |
| json_in_json | 157.5 | 266.70 MB/s | 72 | 2 |
| prose_with_quotes | 96.1 | 395.58 MB/s | 64 | 2 |
| control_bytes | 116.6 | 205.78 MB/s | 56 | 2 |
| mostly_clean_one_quote | 140.6 | 2169.43 MB/s | 320 | 1 |
| unicode_clean | 278.0 | 848.91 MB/s | 240 | 1 |
| unicode_with_quotes | 167.4 | 376.25 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 150.6 | 2025.60 MB/s | 320 | 1 |
| invalid_utf8_dense | 650.3 | 184.54 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.0 | 2009.78 MB/s | 0 | 0 |
| sentence_clean | 15.2 | 2905.03 MB/s | 0 | 0 |
| url_clean | 10.5 | 4947.24 MB/s | 0 | 0 |
| log_line_clean | 33.1 | 10162.77 MB/s | 0 | 0 |
| path_with_backslash | 52.0 | 712.04 MB/s | 0 | 0 |
| json_in_json | 83.6 | 502.15 MB/s | 0 | 0 |
| prose_with_quotes | 30.8 | 1234.98 MB/s | 0 | 0 |
| control_bytes | 47.3 | 507.88 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.1 | 8687.20 MB/s | 0 | 0 |
| unicode_clean | 224.3 | 1052.07 MB/s | 0 | 0 |
| unicode_with_quotes | 86.2 | 731.07 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 49.9 | 6111.21 MB/s | 0 | 0 |
| invalid_utf8_dense | 438.5 | 273.65 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2716.0 | 3697.49 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2703.0 | 3714.49 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2685.0 | 3739.03 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10524.0 | 954.08 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 732.5 | 2472.26 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1033.0 | 1753.54 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.7 | 5865.31 MB/s | 0 | 0 |
| sentence_clean | 4.7 | 9324.54 MB/s | 0 | 0 |
| url_clean | 4.7 | 11019.15 MB/s | 0 | 0 |
| log_line_clean | 10.6 | 31642.93 MB/s | 0 | 0 |
| path_escaped | 83.3 | 516.29 MB/s | 48 | 1 |
| json_in_json | 109.3 | 494.24 MB/s | 64 | 1 |
| prose_with_quotes | 70.1 | 584.65 MB/s | 48 | 1 |
| unicode_heavy | 3.4 | 8816.31 MB/s | 0 | 0 |
| unicode_escaped_dense | 295.8 | 649.10 MB/s | 192 | 1 |
| mostly_clean_one_escape | 127.2 | 2405.71 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.2 | 4933.83 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8290.23 MB/s | 0 | 0 |
| url_clean | 5.3 | 9798.04 MB/s | 0 | 0 |
| log_line_clean | 11.1 | 30373.56 MB/s | 0 | 0 |
| path_escaped | 42.7 | 1007.86 MB/s | 0 | 0 |
| json_in_json | 66.8 | 808.80 MB/s | 0 | 0 |
| prose_with_quotes | 31.7 | 1293.94 MB/s | 0 | 0 |
| unicode_heavy | 3.8 | 7823.08 MB/s | 0 | 0 |
| unicode_escaped_dense | 221.0 | 868.74 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.4 | 12531.00 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 2.9 | — | 0 | 0 |
| 3digit | 3.6 | — | 0 | 0 |
| 5digit | 4.0 | — | 0 | 0 |
| 10digit | 5.1 | — | 0 | 0 |
| 13digit | 5.1 | — | 0 | 0 |
| 16digit | 5.1 | — | 0 | 0 |
| 19digit | 6.7 | — | 0 | 0 |
| neg10digit | 5.0 | — | 0 | 0 |
| 20digit_overflow | 6.7 | — | 0 | 0 |
| notanint | 2.9 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 3.5 | — | 0 | 0 |
| 10digit | 4.6 | — | 0 | 0 |
| 13digit | 4.6 | — | 0 | 0 |
| 20digit | 6.5 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 2.3 | — | 0 | 0 |
| short | 3.4 | — | 0 | 0 |
| medium | 4.1 | — | 0 | 0 |
| long | 8.6 | — | 0 | 0 |
| escaped | 37.5 | — | 8 | 1 |
| notastring | 2.0 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.9 | — | 0 | 0 |
| false | 1.1 | — | 0 | 0 |
| null | 1.0 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 1.8 | — | 0 | 0 |
| number | 1.8 | — | 0 | 0 |
| object | 1.8 | — | 0 | 0 |
| array | 1.8 | — | 0 | 0 |
| null | 2.5 | — | 0 | 0 |
| true | 2.6 | — | 0 | 0 |
| invalid | 2.1 | — | 0 | 0 |
| ws_number | 2.7 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 15.8 | — | 8 | 1 |
| medium | 28.9 | — | 32 | 1 |
| long | 52.7 | — | 128 | 1 |
| escaped | 34.9 | — | 8 | 1 |
| long_late_escape | 102.5 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1661.0 | 1325.38 MB/s | 0 | 0 |
| index | 1720.0 | 1279.74 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1720.0 | 1279.37 MB/s | 0 | 0 |
| strings | 1047.0 | 2198.16 MB/s | 0 | 0 |
| records | 991.4 | 2792.91 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 13.1 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 5.1 | — | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3.9 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 63.2 | — | 0 | 0 |
| append_empty | 18.2 | — | 0 | 0 |
| replace | 47.4 | — | 0 | 0 |
| create_nested | 41.2 | — | 0 | 0 |
| overwrite_nonobject | 48.7 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 109.2 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 284.7 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 98.7 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 131.8 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 482256.0 | 2293.93 MB/s | 2293536 | 25 |
| stream | 237182.0 | 4664.19 MB/s | 67136 | 5 |
| stream_reused | 199986.0 | 5531.69 MB/s | 35 | 1 |
| stream_points | 1418113.0 | 780.09 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 17042.0 | 1291.02 MB/s | 0 | 0 |
| scalars/stream | 27129.0 | 810.99 MB/s | 67120 | 4 |
| scalars/stream_reused | 18546.0 | 1186.26 MB/s | 32 | 1 |
| strings/inmemory | 10394.0 | 2212.92 MB/s | 0 | 0 |
| strings/stream | 20645.0 | 1114.10 MB/s | 67120 | 4 |
| strings/stream_reused | 12216.0 | 1882.80 MB/s | 32 | 1 |
| records/inmemory | 9828.0 | 2869.29 MB/s | 0 | 0 |
| records/stream | 19825.0 | 1422.34 MB/s | 67120 | 4 |
| records/stream_reused | 10990.0 | 2565.74 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 199038.0 | 4415.86 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 989.0 | 1320.46 MB/s | 32 | 1 |
| inmemory | 794.2 | 1644.47 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 2086.0 | 1729.52 MB/s | 32 | 1 |
| arrayeach | 2102.0 | 1715.60 MB/s | 32 | 1 |
| inmemory | 1587.0 | 2272.66 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 291924.0 | 4080.73 MB/s | 69 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1955.0 | 1415.09 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1878.0 | 1472.54 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 28.4 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 222.8 | 834.73 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1033.0 | 2226.58 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 814.7 | 2222.96 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 808.4 | 2240.09 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 822.9 | 2200.64 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 987.3 | 2804.65 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1660.0 | 1325.63 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2752.0 | 481.49 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 88.2 | — | 24 | 1 |
| arena | 79.8 | — | 24 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 23.8 | — | 64 | 0 |
| make | 47.4 | — | 64 | 1 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 23014.0 | 853.05 MB/s | 0 | 0 |
| kernel/sep"," | 8042.0 | 2441.14 MB/s | 0 | 0 |
| scalar/sep",_" | 25375.0 | 931.28 MB/s | 0 | 0 |
| kernel/sep",_" | 11096.0 | 2129.67 MB/s | 0 | 0 |
| kernel-only | 8031.0 | 2444.57 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6995.0 | 1155.04 MB/s | 0 | 0 |
| "12," | 6551.0 | 1843.90 MB/s | 0 | 0 |
| "123," | 6912.0 | 2326.33 MB/s | 0 | 0 |
| "1234," | 8119.0 | 2473.20 MB/s | 0 | 0 |
| "123456," | 8463.0 | 3317.86 MB/s | 0 | 0 |
| "1234567," | 8270.0 | 3878.95 MB/s | 0 | 0 |
| "1234,_" | 10694.0 | 2251.63 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 2.7 | — | 0 | 0 |
| d03 | 3.5 | — | 0 | 0 |
| d05 | 3.9 | — | 0 | 0 |
| d08 | 3.9 | — | 0 | 0 |
| d10 | 5.0 | — | 0 | 0 |
| d13 | 5.0 | — | 0 | 0 |
| d16 | 5.0 | — | 0 | 0 |
| d19 | 6.7 | — | 0 | 0 |
| d20 | 7.1 | — | 0 | 0 |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| numbers/whole | 191.9 | 43796.01 MB/s | 0 | 0 |
| numbers/4k | 219.3 | 38317.42 MB/s | 0 | 0 |
| records/whole | 12209.0 | 754.26 MB/s | 0 | 0 |
| records/4k | 12239.0 | 752.45 MB/s | 0 | 0 |
| strings/whole | 5221.0 | 2146.05 MB/s | 0 | 0 |
| strings/4k | 5249.0 | 2134.54 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 428.8 | — | 0 | 0 |
| canada | 312.1 | — | 0 | 0 |
| mesh | 271.5 | — | 0 | 0 |
| array | 223.7 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 499.1 | — | 0 | 0 |
| canada | 410.2 | — | 0 | 0 |
| mesh | 291.5 | — | 0 | 0 |
| array | 262.6 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4304.0 | 5575.42 MB/s | 0 | 0 |
| numberObj/goloop | 1776.0 | 5740.69 MB/s | 0 | 0 |
| nestedMixed/goloop | 2379.0 | 4539.86 MB/s | 0 | 0 |
| stringObj/neon | 2819.0 | 8511.61 MB/s | 0 | 0 |
| numberObj/neon | 1195.0 | 8532.00 MB/s | 0 | 0 |
| nestedMixed/neon | 1596.0 | 6768.67 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 11940.0 | 2009.84 MB/s | 0 | 0 |
| stringObj/dispatch | 2820.0 | 8510.41 MB/s | 0 | 0 |
| numberObj/current | 4456.0 | 2288.62 MB/s | 0 | 0 |
| numberObj/dispatch | 1196.0 | 8526.73 MB/s | 0 | 0 |
| numberArr/current | 150.5 | 43847.37 MB/s | 0 | 0 |
| numberArr/dispatch | 156.2 | 42273.12 MB/s | 0 | 0 |
| nestedMixed/current | 16065.0 | 672.33 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1597.0 | 6762.20 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 18.5 | 973.58 MB/s | 0 | 0 |
| record | 13.6 | 3979.25 MB/s | 0 | 0 |
| tiny | 13.6 | 515.66 MB/s | 0 | 0 |
| twoBlock | 21.0 | 4182.89 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 16.3 | 1106.09 MB/s | 0 | 0 |
| record | 43.8 | 1233.33 MB/s | 0 | 0 |
| tiny | 13.3 | 527.93 MB/s | 0 | 0 |
| twoBlock | 25.1 | 3506.24 MB/s | 0 | 0 |
