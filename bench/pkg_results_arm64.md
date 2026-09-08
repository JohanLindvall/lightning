# lightning main-module benchmarks

- generated 2026-09-08T12:19:06Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 197.7 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.0 | 533.29 MB/s | 16 | 1 |
| sentence_clean | 37.1 | 1184.82 MB/s | 48 | 1 |
| url_clean | 40.5 | 1283.71 MB/s | 64 | 1 |
| log_line_clean | 116.9 | 2874.60 MB/s | 352 | 1 |
| path_with_backslash | 117.8 | 314.17 MB/s | 56 | 2 |
| json_in_json | 155.0 | 270.97 MB/s | 72 | 2 |
| prose_with_quotes | 95.6 | 397.39 MB/s | 64 | 2 |
| control_bytes | 114.3 | 209.92 MB/s | 56 | 2 |
| mostly_clean_one_quote | 136.1 | 2240.85 MB/s | 320 | 1 |
| unicode_clean | 273.2 | 863.87 MB/s | 240 | 1 |
| unicode_with_quotes | 161.0 | 391.31 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 144.8 | 2106.05 MB/s | 320 | 1 |
| invalid_utf8_dense | 630.8 | 190.24 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.5 | 2124.19 MB/s | 0 | 0 |
| sentence_clean | 15.1 | 2918.01 MB/s | 0 | 0 |
| url_clean | 10.6 | 4885.91 MB/s | 0 | 0 |
| log_line_clean | 33.4 | 10074.76 MB/s | 0 | 0 |
| path_with_backslash | 53.1 | 696.79 MB/s | 0 | 0 |
| json_in_json | 84.3 | 498.16 MB/s | 0 | 0 |
| prose_with_quotes | 31.1 | 1221.57 MB/s | 0 | 0 |
| control_bytes | 45.4 | 528.18 MB/s | 0 | 0 |
| mostly_clean_one_quote | 34.8 | 8756.95 MB/s | 0 | 0 |
| unicode_clean | 224.7 | 1050.28 MB/s | 0 | 0 |
| unicode_with_quotes | 86.2 | 730.67 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 49.6 | 6146.45 MB/s | 0 | 0 |
| invalid_utf8_dense | 433.1 | 277.08 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2734.0 | 3672.63 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2703.0 | 3714.72 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2677.0 | 3750.86 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10368.0 | 968.45 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 737.4 | 2456.02 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1048.0 | 1728.02 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.9 | 5596.54 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9027.90 MB/s | 0 | 0 |
| url_clean | 4.9 | 10667.95 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31455.61 MB/s | 0 | 0 |
| path_escaped | 76.9 | 559.36 MB/s | 48 | 1 |
| json_in_json | 105.5 | 512.02 MB/s | 64 | 1 |
| prose_with_quotes | 64.9 | 631.71 MB/s | 48 | 1 |
| unicode_heavy | 3.5 | 8476.87 MB/s | 0 | 0 |
| unicode_escaped_dense | 292.6 | 656.27 MB/s | 192 | 1 |
| mostly_clean_one_escape | 125.5 | 2437.65 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.2 | 4931.71 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8292.03 MB/s | 0 | 0 |
| url_clean | 5.3 | 9799.47 MB/s | 0 | 0 |
| log_line_clean | 11.2 | 29956.77 MB/s | 0 | 0 |
| path_escaped | 42.4 | 1013.97 MB/s | 0 | 0 |
| json_in_json | 66.7 | 809.78 MB/s | 0 | 0 |
| prose_with_quotes | 31.5 | 1301.54 MB/s | 0 | 0 |
| unicode_heavy | 4.0 | 7528.73 MB/s | 0 | 0 |
| unicode_escaped_dense | 220.5 | 870.67 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.5 | 12504.44 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 2.7 | — | 0 | 0 |
| 3digit | 3.8 | — | 0 | 0 |
| 5digit | 4.0 | — | 0 | 0 |
| 10digit | 5.0 | — | 0 | 0 |
| 13digit | 5.0 | — | 0 | 0 |
| 16digit | 5.0 | — | 0 | 0 |
| 19digit | 6.7 | — | 0 | 0 |
| neg10digit | 4.9 | — | 0 | 0 |
| 20digit_overflow | 6.7 | — | 0 | 0 |
| notanint | 2.9 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 3.4 | — | 0 | 0 |
| 10digit | 4.6 | — | 0 | 0 |
| 13digit | 4.6 | — | 0 | 0 |
| 20digit | 6.5 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 2.4 | — | 0 | 0 |
| short | 3.5 | — | 0 | 0 |
| medium | 4.3 | — | 0 | 0 |
| long | 8.9 | — | 0 | 0 |
| escaped | 37.1 | — | 8 | 1 |
| notastring | 2.1 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.9 | — | 0 | 0 |
| false | 1.1 | — | 0 | 0 |
| null | 0.9 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 1.8 | — | 0 | 0 |
| number | 1.8 | — | 0 | 0 |
| object | 1.8 | — | 0 | 0 |
| array | 1.8 | — | 0 | 0 |
| null | 2.6 | — | 0 | 0 |
| true | 2.6 | — | 0 | 0 |
| invalid | 2.1 | — | 0 | 0 |
| ws_number | 2.7 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 14.8 | — | 8 | 1 |
| medium | 26.5 | — | 32 | 1 |
| long | 51.5 | — | 128 | 1 |
| escaped | 35.0 | — | 8 | 1 |
| long_late_escape | 98.7 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1660.0 | 1325.65 MB/s | 0 | 0 |
| index | 1782.0 | 1235.10 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1782.0 | 1235.24 MB/s | 0 | 0 |
| strings | 1017.0 | 2263.56 MB/s | 0 | 0 |
| records | 984.2 | 2813.32 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 13.3 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 5.1 | — | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 4.1 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 64.1 | — | 0 | 0 |
| append_empty | 18.1 | — | 0 | 0 |
| replace | 47.5 | — | 0 | 0 |
| create_nested | 41.3 | — | 0 | 0 |
| overwrite_nonobject | 49.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 107.8 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 285.0 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 98.4 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 132.4 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 474605.0 | 2330.91 MB/s | 2293537 | 25 |
| stream | 236684.0 | 4674.01 MB/s | 67136 | 5 |
| stream_reused | 199762.0 | 5537.89 MB/s | 35 | 1 |
| stream_points | 1416248.0 | 781.12 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 17110.0 | 1285.84 MB/s | 0 | 0 |
| scalars/stream | 27722.0 | 793.64 MB/s | 67120 | 4 |
| scalars/stream_reused | 19570.0 | 1124.21 MB/s | 32 | 1 |
| strings/inmemory | 10500.0 | 2190.51 MB/s | 0 | 0 |
| strings/stream | 20139.0 | 1142.13 MB/s | 67120 | 4 |
| strings/stream_reused | 12106.0 | 1900.03 MB/s | 32 | 1 |
| records/inmemory | 9901.0 | 2848.09 MB/s | 0 | 0 |
| records/stream | 19316.0 | 1459.86 MB/s | 67120 | 4 |
| records/stream_reused | 11015.0 | 2560.07 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 167168.0 | 5257.73 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 981.5 | 1330.57 MB/s | 32 | 1 |
| inmemory | 799.9 | 1632.77 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 2072.0 | 1740.78 MB/s | 32 | 1 |
| arrayeach | 2086.0 | 1728.89 MB/s | 32 | 1 |
| inmemory | 1590.0 | 2268.87 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 293102.0 | 4064.34 MB/s | 69 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1970.0 | 1404.34 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1904.0 | 1452.81 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 27.9 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 222.0 | 837.92 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 990.8 | 2322.32 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 814.6 | 2223.11 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 806.2 | 2246.24 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 813.1 | 2227.29 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 991.4 | 2792.96 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1661.0 | 1325.23 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2735.0 | 484.50 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 87.4 | — | 24 | 1 |
| arena | 79.2 | — | 24 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 23.6 | — | 64 | 0 |
| make | 41.7 | — | 64 | 1 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 23321.0 | 841.83 MB/s | 0 | 0 |
| kernel/sep"," | 8030.0 | 2444.70 MB/s | 0 | 0 |
| scalar/sep",_" | 25450.0 | 928.52 MB/s | 0 | 0 |
| kernel/sep",_" | 10903.0 | 2167.30 MB/s | 0 | 0 |
| kernel-only | 8010.0 | 2451.09 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6985.0 | 1156.83 MB/s | 0 | 0 |
| "12," | 6546.0 | 1845.42 MB/s | 0 | 0 |
| "123," | 6956.0 | 2311.52 MB/s | 0 | 0 |
| "1234," | 7733.0 | 2596.65 MB/s | 0 | 0 |
| "123456," | 8454.0 | 3321.44 MB/s | 0 | 0 |
| "1234567," | 8312.0 | 3859.63 MB/s | 0 | 0 |
| "1234,_" | 10759.0 | 2237.94 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 2.7 | — | 0 | 0 |
| d03 | 3.6 | — | 0 | 0 |
| d05 | 3.9 | — | 0 | 0 |
| d08 | 3.9 | — | 0 | 0 |
| d10 | 5.0 | — | 0 | 0 |
| d13 | 5.0 | — | 0 | 0 |
| d16 | 5.0 | — | 0 | 0 |
| d19 | 6.5 | — | 0 | 0 |
| d20 | 7.0 | — | 0 | 0 |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/whole | 12182.0 | 755.98 MB/s | 0 | 0 |
| records/4k | 12254.0 | 751.50 MB/s | 0 | 0 |
| strings/whole | 5301.0 | 2113.68 MB/s | 0 | 0 |
| strings/4k | 5310.0 | 2110.23 MB/s | 0 | 0 |
| numbers/whole | 192.0 | 43766.24 MB/s | 0 | 0 |
| numbers/4k | 219.2 | 38336.68 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 459.2 | — | 0 | 0 |
| canada | 335.4 | — | 0 | 0 |
| mesh | 272.0 | — | 0 | 0 |
| array | 223.1 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 458.3 | — | 0 | 0 |
| canada | 410.6 | — | 0 | 0 |
| mesh | 292.2 | — | 0 | 0 |
| array | 260.9 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4266.0 | 5625.24 MB/s | 0 | 0 |
| numberObj/goloop | 1777.0 | 5739.98 MB/s | 0 | 0 |
| nestedMixed/goloop | 2383.0 | 4531.70 MB/s | 0 | 0 |
| stringObj/neon | 2819.0 | 8512.14 MB/s | 0 | 0 |
| numberObj/neon | 1195.0 | 8533.89 MB/s | 0 | 0 |
| nestedMixed/neon | 1596.0 | 6768.92 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 11926.0 | 2012.32 MB/s | 0 | 0 |
| stringObj/dispatch | 2820.0 | 8509.47 MB/s | 0 | 0 |
| numberObj/current | 4417.0 | 2308.72 MB/s | 0 | 0 |
| numberObj/dispatch | 1196.0 | 8528.50 MB/s | 0 | 0 |
| numberArr/current | 150.7 | 43803.77 MB/s | 0 | 0 |
| numberArr/dispatch | 156.2 | 42263.87 MB/s | 0 | 0 |
| nestedMixed/current | 15989.0 | 675.54 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1596.0 | 6767.59 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 18.2 | 988.72 MB/s | 0 | 0 |
| record | 13.6 | 3976.33 MB/s | 0 | 0 |
| tiny | 13.6 | 515.65 MB/s | 0 | 0 |
| twoBlock | 21.0 | 4187.00 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 16.2 | 1112.51 MB/s | 0 | 0 |
| record | 43.6 | 1239.16 MB/s | 0 | 0 |
| tiny | 13.3 | 526.84 MB/s | 0 | 0 |
| twoBlock | 25.3 | 3481.31 MB/s | 0 | 0 |
