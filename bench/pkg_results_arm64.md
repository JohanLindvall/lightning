# lightning main-module benchmarks

- generated 2026-09-07T18:59:38Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 189.8 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.0 | 533.39 MB/s | 16 | 1 |
| sentence_clean | 37.6 | 1170.90 MB/s | 48 | 1 |
| url_clean | 41.3 | 1259.24 MB/s | 64 | 1 |
| log_line_clean | 118.7 | 2830.84 MB/s | 352 | 1 |
| path_with_backslash | 118.2 | 313.01 MB/s | 56 | 2 |
| json_in_json | 155.2 | 270.66 MB/s | 72 | 2 |
| prose_with_quotes | 94.9 | 400.48 MB/s | 64 | 2 |
| control_bytes | 117.0 | 205.15 MB/s | 56 | 2 |
| mostly_clean_one_quote | 137.4 | 2220.22 MB/s | 320 | 1 |
| unicode_clean | 272.1 | 867.46 MB/s | 240 | 1 |
| unicode_with_quotes | 162.3 | 388.17 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 147.7 | 2065.65 MB/s | 320 | 1 |
| invalid_utf8_dense | 640.1 | 187.48 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.9 | 2012.98 MB/s | 0 | 0 |
| sentence_clean | 14.9 | 2942.30 MB/s | 0 | 0 |
| url_clean | 10.6 | 4908.35 MB/s | 0 | 0 |
| log_line_clean | 33.3 | 10097.79 MB/s | 0 | 0 |
| path_with_backslash | 51.6 | 716.89 MB/s | 0 | 0 |
| json_in_json | 84.1 | 499.65 MB/s | 0 | 0 |
| prose_with_quotes | 30.6 | 1239.97 MB/s | 0 | 0 |
| control_bytes | 47.8 | 502.47 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.6 | 8564.34 MB/s | 0 | 0 |
| unicode_clean | 222.1 | 1062.53 MB/s | 0 | 0 |
| unicode_with_quotes | 85.8 | 734.46 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 50.0 | 6103.69 MB/s | 0 | 0 |
| invalid_utf8_dense | 440.9 | 272.14 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2931.0 | 3425.70 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2726.0 | 3683.02 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2709.0 | 3706.95 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10502.0 | 956.13 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 796.8 | 2272.85 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1227.0 | 1475.95 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.8 | 5759.55 MB/s | 0 | 0 |
| sentence_clean | 4.8 | 9256.35 MB/s | 0 | 0 |
| url_clean | 4.8 | 10932.58 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31561.02 MB/s | 0 | 0 |
| path_escaped | 86.7 | 495.89 MB/s | 48 | 1 |
| json_in_json | 118.6 | 455.39 MB/s | 64 | 1 |
| prose_with_quotes | 73.2 | 559.82 MB/s | 48 | 1 |
| unicode_heavy | 3.5 | 8477.00 MB/s | 0 | 0 |
| unicode_escaped_dense | 293.8 | 653.59 MB/s | 192 | 1 |
| mostly_clean_one_escape | 123.6 | 2475.89 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.4 | 4712.51 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8287.33 MB/s | 0 | 0 |
| url_clean | 5.3 | 9794.41 MB/s | 0 | 0 |
| log_line_clean | 11.2 | 29953.86 MB/s | 0 | 0 |
| path_escaped | 45.0 | 954.73 MB/s | 0 | 0 |
| json_in_json | 68.9 | 784.10 MB/s | 0 | 0 |
| prose_with_quotes | 32.6 | 1258.09 MB/s | 0 | 0 |
| unicode_heavy | 4.0 | 7521.76 MB/s | 0 | 0 |
| unicode_escaped_dense | 221.2 | 868.02 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.5 | 12507.15 MB/s | 0 | 0 |

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
| notanint | 3.0 | — | 0 | 0 |

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
| long | 8.7 | — | 0 | 0 |
| escaped | 37.0 | — | 8 | 1 |
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
| short | 16.5 | — | 8 | 1 |
| medium | 28.5 | — | 32 | 1 |
| long | 50.9 | — | 128 | 1 |
| escaped | 34.4 | — | 8 | 1 |
| long_late_escape | 101.5 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1661.0 | 1325.29 MB/s | 0 | 0 |
| index | 1661.0 | 1325.28 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1661.0 | 1325.17 MB/s | 0 | 0 |
| strings | 1025.0 | 2245.01 MB/s | 0 | 0 |
| records | 1152.0 | 2404.26 MB/s | 0 | 0 |

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
| append | 64.8 | — | 0 | 0 |
| append_empty | 18.2 | — | 0 | 0 |
| replace | 47.7 | — | 0 | 0 |
| create_nested | 41.8 | — | 0 | 0 |
| overwrite_nonobject | 49.9 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 106.9 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 286.1 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 99.0 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 123.6 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 450874.0 | 2453.59 MB/s | 2293536 | 25 |
| stream | 255325.0 | 4332.76 MB/s | 66992 | 4 |
| stream_reused | 223115.0 | 4958.25 MB/s | 36 | 1 |
| stream_points | 1427961.0 | 774.71 MB/s | 66992 | 4 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 16918.0 | 1300.42 MB/s | 0 | 0 |
| scalars/stream | 30147.0 | 729.79 MB/s | 66992 | 4 |
| strings/inmemory | 10375.0 | 2216.98 MB/s | 0 | 0 |
| strings/stream | 21714.0 | 1059.28 MB/s | 66992 | 4 |
| records/inmemory | 11142.0 | 2530.75 MB/s | 0 | 0 |
| records/stream | 20010.0 | 1409.23 MB/s | 66992 | 4 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 241311.0 | 3642.28 MB/s | 66992 | 4 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1991.0 | 1389.12 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1914.0 | 1445.03 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 28.2 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 222.2 | 837.27 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1029.0 | 2235.64 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 857.9 | 2110.89 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 877.0 | 2065.10 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 872.1 | 2076.64 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1155.0 | 2397.89 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1661.0 | 1324.80 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2951.0 | 449.06 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 89.7 | — | 24 | 1 |
| arena | 79.2 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 23116.0 | 849.30 MB/s | 0 | 0 |
| kernel/sep"," | 8034.0 | 2443.76 MB/s | 0 | 0 |
| scalar/sep",_" | 25410.0 | 929.98 MB/s | 0 | 0 |
| kernel/sep",_" | 11058.0 | 2137.09 MB/s | 0 | 0 |
| kernel-only | 8031.0 | 2444.58 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6176.0 | 1308.33 MB/s | 0 | 0 |
| "12," | 6539.0 | 1847.38 MB/s | 0 | 0 |
| "123," | 6938.0 | 2317.70 MB/s | 0 | 0 |
| "1234," | 8156.0 | 2461.95 MB/s | 0 | 0 |
| "123456," | 8466.0 | 3316.83 MB/s | 0 | 0 |
| "1234567," | 8281.0 | 3874.08 MB/s | 0 | 0 |
| "1234,_" | 10731.0 | 2243.85 MB/s | 0 | 0 |

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
| records/whole | 11510.0 | 800.10 MB/s | 0 | 0 |
| records/4k | 11486.0 | 801.78 MB/s | 0 | 0 |
| strings/whole | 4150.0 | 2700.30 MB/s | 0 | 0 |
| strings/4k | 4164.0 | 2691.24 MB/s | 0 | 0 |
| numbers/whole | 199.2 | 42182.40 MB/s | 0 | 0 |
| numbers/4k | 240.8 | 34891.10 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 459.6 | — | 0 | 0 |
| canada | 334.2 | — | 0 | 0 |
| mesh | 272.0 | — | 0 | 0 |
| array | 223.2 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 457.8 | — | 0 | 0 |
| canada | 409.3 | — | 0 | 0 |
| mesh | 292.7 | — | 0 | 0 |
| array | 261.1 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4391.0 | 5464.83 MB/s | 0 | 0 |
| numberObj/goloop | 1819.0 | 5607.76 MB/s | 0 | 0 |
| nestedMixed/goloop | 2413.0 | 4476.16 MB/s | 0 | 0 |
| stringObj/neon | 2901.0 | 8273.25 MB/s | 0 | 0 |
| numberObj/neon | 1215.0 | 8396.61 MB/s | 0 | 0 |
| nestedMixed/neon | 1647.0 | 6559.54 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10298.0 | 2330.30 MB/s | 0 | 0 |
| stringObj/dispatch | 2881.0 | 8330.98 MB/s | 0 | 0 |
| numberObj/current | 5648.0 | 1805.73 MB/s | 0 | 0 |
| numberObj/dispatch | 1216.0 | 8387.85 MB/s | 0 | 0 |
| numberArr/current | 158.9 | 41535.83 MB/s | 0 | 0 |
| numberArr/dispatch | 160.4 | 41158.71 MB/s | 0 | 0 |
| nestedMixed/current | 14301.0 | 755.26 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1647.0 | 6556.90 MB/s | 0 | 0 |
