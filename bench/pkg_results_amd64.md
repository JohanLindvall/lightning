# lightning main-module benchmarks

- generated 2026-09-08T16:42:00Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 222.6 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 37.1 | 430.98 MB/s | 16 | 1 |
| sentence_clean | 47.0 | 935.61 MB/s | 48 | 1 |
| url_clean | 46.0 | 1131.20 MB/s | 64 | 1 |
| log_line_clean | 108.2 | 3105.17 MB/s | 352 | 1 |
| path_with_backslash | 137.2 | 269.69 MB/s | 56 | 2 |
| json_in_json | 182.7 | 229.89 MB/s | 72 | 2 |
| prose_with_quotes | 119.3 | 318.58 MB/s | 64 | 2 |
| control_bytes | 139.3 | 172.35 MB/s | 56 | 2 |
| mostly_clean_one_quote | 119.9 | 2542.75 MB/s | 320 | 1 |
| unicode_clean | 357.6 | 659.94 MB/s | 240 | 1 |
| unicode_with_quotes | 190.4 | 330.80 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 126.2 | 2417.58 MB/s | 320 | 1 |
| invalid_utf8_dense | 781.5 | 153.56 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 13.2 | 1214.16 MB/s | 0 | 0 |
| sentence_clean | 24.7 | 1783.12 MB/s | 0 | 0 |
| url_clean | 25.9 | 2010.68 MB/s | 0 | 0 |
| log_line_clean | 33.8 | 9952.63 MB/s | 0 | 0 |
| path_with_backslash | 72.5 | 510.07 MB/s | 0 | 0 |
| json_in_json | 113.9 | 368.66 MB/s | 0 | 0 |
| prose_with_quotes | 46.9 | 811.03 MB/s | 0 | 0 |
| control_bytes | 74.0 | 324.08 MB/s | 0 | 0 |
| mostly_clean_one_quote | 38.6 | 7900.67 MB/s | 0 | 0 |
| unicode_clean | 306.9 | 769.06 MB/s | 0 | 0 |
| unicode_with_quotes | 121.4 | 519.13 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 61.6 | 4950.17 MB/s | 0 | 0 |
| invalid_utf8_dense | 640.1 | 187.48 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2442.0 | 4111.68 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2515.0 | 3991.71 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2477.0 | 4053.08 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12735.0 | 788.43 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 715.8 | 2529.96 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1390.0 | 1303.10 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.7 | 4321.34 MB/s | 0 | 0 |
| sentence_clean | 6.6 | 6679.66 MB/s | 0 | 0 |
| url_clean | 6.6 | 7908.01 MB/s | 0 | 0 |
| log_line_clean | 11.2 | 30070.06 MB/s | 0 | 0 |
| path_escaped | 99.9 | 430.30 MB/s | 48 | 1 |
| json_in_json | 130.1 | 415.22 MB/s | 64 | 1 |
| prose_with_quotes | 80.5 | 509.46 MB/s | 48 | 1 |
| unicode_heavy | 5.0 | 6044.39 MB/s | 0 | 0 |
| unicode_escaped_dense | 360.6 | 532.40 MB/s | 192 | 1 |
| mostly_clean_one_escape | 99.5 | 3074.96 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.5 | 3539.80 MB/s | 0 | 0 |
| sentence_clean | 7.0 | 6293.86 MB/s | 0 | 0 |
| url_clean | 7.0 | 7420.81 MB/s | 0 | 0 |
| log_line_clean | 11.5 | 29123.14 MB/s | 0 | 0 |
| path_escaped | 72.9 | 589.87 MB/s | 0 | 0 |
| json_in_json | 101.8 | 530.44 MB/s | 0 | 0 |
| prose_with_quotes | 57.3 | 715.05 MB/s | 0 | 0 |
| unicode_heavy | 5.3 | 5614.81 MB/s | 0 | 0 |
| unicode_escaped_dense | 314.9 | 609.78 MB/s | 0 | 0 |
| mostly_clean_one_escape | 26.8 | 11422.66 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 5.4 | — | 0 | 0 |
| 3digit | 7.0 | — | 0 | 0 |
| 5digit | 6.0 | — | 0 | 0 |
| 10digit | 8.1 | — | 0 | 0 |
| 13digit | 8.1 | — | 0 | 0 |
| 16digit | 8.1 | — | 0 | 0 |
| 19digit | 10.5 | — | 0 | 0 |
| neg10digit | 7.8 | — | 0 | 0 |
| 20digit_overflow | 9.9 | — | 0 | 0 |
| notanint | 4.9 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 6.2 | — | 0 | 0 |
| 10digit | 7.2 | — | 0 | 0 |
| 13digit | 7.2 | — | 0 | 0 |
| 20digit | 10.1 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 3.3 | — | 0 | 0 |
| short | 4.5 | — | 0 | 0 |
| medium | 5.6 | — | 0 | 0 |
| long | 10.3 | — | 0 | 0 |
| escaped | 41.3 | — | 8 | 1 |
| notastring | 2.5 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 1.0 | — | 0 | 0 |
| false | 1.2 | — | 0 | 0 |
| null | 1.6 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 2.9 | — | 0 | 0 |
| number | 2.9 | — | 0 | 0 |
| object | 2.9 | — | 0 | 0 |
| array | 2.9 | — | 0 | 0 |
| null | 3.3 | — | 0 | 0 |
| true | 4.1 | — | 0 | 0 |
| invalid | 2.9 | — | 0 | 0 |
| ws_number | 4.1 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 19.3 | — | 8 | 1 |
| medium | 27.9 | — | 32 | 1 |
| long | 45.7 | — | 128 | 1 |
| escaped | 39.9 | — | 8 | 1 |
| long_late_escape | 61.9 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 2080.0 | 1058.15 MB/s | 0 | 0 |
| index | 2896.0 | 759.96 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 2909.0 | 756.73 MB/s | 0 | 0 |
| strings | 931.7 | 2469.67 MB/s | 0 | 0 |
| records | 831.2 | 3331.34 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 18.1 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10.8 | — | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 5.6 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 75.5 | — | 0 | 0 |
| append_empty | 24.9 | — | 0 | 0 |
| replace | 52.1 | — | 0 | 0 |
| create_nested | 53.7 | — | 0 | 0 |
| overwrite_nonobject | 58.0 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 142.9 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 337.4 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 110.7 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 158.3 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 422525.0 | 2618.21 MB/s | 2293537 | 25 |
| stream | 167898.0 | 6588.90 MB/s | 67136 | 5 |
| stream_reused | 154059.0 | 7180.76 MB/s | 34 | 1 |
| stream_points | 1428906.0 | 774.20 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 20586.0 | 1068.73 MB/s | 0 | 0 |
| scalars/stream | 30544.0 | 720.30 MB/s | 67120 | 4 |
| scalars/stream_reused | 22928.0 | 959.55 MB/s | 32 | 1 |
| strings/inmemory | 9564.0 | 2405.01 MB/s | 0 | 0 |
| strings/stream | 18966.0 | 1212.73 MB/s | 67120 | 4 |
| strings/stream_reused | 11108.0 | 2070.69 MB/s | 32 | 1 |
| records/inmemory | 8807.0 | 3201.60 MB/s | 0 | 0 |
| records/stream | 17775.0 | 1586.41 MB/s | 67120 | 4 |
| records/stream_reused | 9656.0 | 2920.39 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 114296.0 | 7689.89 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 958.5 | 1362.55 MB/s | 32 | 1 |
| inmemory | 723.2 | 1805.93 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 1786.0 | 2019.74 MB/s | 32 | 1 |
| arrayeach | 1844.0 | 1955.84 MB/s | 32 | 1 |
| inmemory | 1610.0 | 2240.06 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 235796.0 | 5052.10 MB/s | 62 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2555.0 | 1082.71 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2483.0 | 1114.13 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 28.2 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 256.2 | 725.91 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 923.0 | 2492.91 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 818.8 | 2211.72 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 771.7 | 2346.64 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 769.3 | 2353.97 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 849.3 | 3260.51 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2074.0 | 1061.00 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3139.0 | 422.12 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 102.6 | — | 24 | 1 |
| arena | 116.4 | — | 24 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 31.2 | — | 64 | 0 |
| make | 28.6 | — | 64 | 1 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 36826.0 | 533.10 MB/s | 0 | 0 |
| kernel/sep"," | 15335.0 | 1280.19 MB/s | 0 | 0 |
| scalar/sep",_" | 40353.0 | 585.61 MB/s | 0 | 0 |
| kernel/sep",_" | 20619.0 | 1146.07 MB/s | 0 | 0 |
| kernel-only | 15285.0 | 1284.43 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 9948.0 | 812.19 MB/s | 0 | 0 |
| "12," | 11976.0 | 1008.69 MB/s | 0 | 0 |
| "123," | 11255.0 | 1428.76 MB/s | 0 | 0 |
| "1234," | 15286.0 | 1313.60 MB/s | 0 | 0 |
| "123456," | 18835.0 | 1490.82 MB/s | 0 | 0 |
| "1234567," | 12897.0 | 2487.41 MB/s | 0 | 0 |
| "1234,_" | 23875.0 | 1008.56 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 4.9 | — | 0 | 0 |
| d03 | 6.1 | — | 0 | 0 |
| d05 | 5.8 | — | 0 | 0 |
| d08 | 5.8 | — | 0 | 0 |
| d10 | 8.1 | — | 0 | 0 |
| d13 | 8.1 | — | 0 | 0 |
| d16 | 8.1 | — | 0 | 0 |
| d19 | 10.8 | — | 0 | 0 |
| d20 | 10.9 | — | 0 | 0 |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/whole | 13522.0 | 681.02 MB/s | 0 | 0 |
| records/4k | 13623.0 | 675.99 MB/s | 0 | 0 |
| strings/whole | 5579.0 | 2008.55 MB/s | 0 | 0 |
| strings/4k | 5678.0 | 1973.48 MB/s | 0 | 0 |
| numbers/whole | 292.8 | 28702.24 MB/s | 0 | 0 |
| numbers/4k | 364.1 | 23081.11 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 624.1 | — | 0 | 0 |
| canada | 399.2 | — | 0 | 0 |
| mesh | 357.9 | — | 0 | 0 |
| array | 286.3 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 539.2 | — | 0 | 0 |
| canada | 501.0 | — | 0 | 0 |
| mesh | 362.4 | — | 0 | 0 |
| array | 318.7 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4013.0 | 5979.48 MB/s | 0 | 0 |
| numberObj/goloop | 1412.0 | 7220.16 MB/s | 0 | 0 |
| nestedMixed/goloop | 1831.0 | 5898.15 MB/s | 0 | 0 |
| stringObj/avx2 | 2187.0 | 10972.86 MB/s | 0 | 0 |
| numberObj/avx2 | 793.4 | 12854.28 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1203.0 | 8978.10 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10819.0 | 2218.15 MB/s | 0 | 0 |
| stringObj/dispatch | 2195.0 | 10931.95 MB/s | 0 | 0 |
| numberObj/current | 4042.0 | 2522.91 MB/s | 0 | 0 |
| numberObj/dispatch | 795.0 | 12828.13 MB/s | 0 | 0 |
| numberArr/current | 212.2 | 31104.22 MB/s | 0 | 0 |
| numberArr/dispatch | 219.0 | 30148.37 MB/s | 0 | 0 |
| nestedMixed/current | 16088.0 | 671.36 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1211.0 | 8918.20 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 17.2 | 1044.25 MB/s | 0 | 0 |
| record | 11.0 | 4920.66 MB/s | 0 | 0 |
| tiny | 11.1 | 633.01 MB/s | 0 | 0 |
| twoBlock | 15.4 | 5710.91 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 18.0 | 998.14 MB/s | 0 | 0 |
| record | 43.8 | 1234.24 MB/s | 0 | 0 |
| tiny | 15.4 | 455.87 MB/s | 0 | 0 |
| twoBlock | 25.3 | 3483.91 MB/s | 0 | 0 |
