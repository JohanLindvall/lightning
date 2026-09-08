# lightning main-module benchmarks

- generated 2026-09-08T04:30:36Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 269.2 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 32.8 | 488.39 MB/s | 16 | 1 |
| sentence_clean | 45.7 | 962.27 MB/s | 48 | 1 |
| url_clean | 42.8 | 1215.65 MB/s | 64 | 1 |
| log_line_clean | 98.0 | 3427.07 MB/s | 352 | 1 |
| path_with_backslash | 143.8 | 257.23 MB/s | 56 | 2 |
| json_in_json | 182.8 | 229.70 MB/s | 72 | 2 |
| prose_with_quotes | 115.9 | 327.97 MB/s | 64 | 2 |
| control_bytes | 139.7 | 171.84 MB/s | 56 | 2 |
| mostly_clean_one_quote | 106.2 | 2871.76 MB/s | 320 | 1 |
| unicode_clean | 300.9 | 784.44 MB/s | 240 | 1 |
| unicode_with_quotes | 180.0 | 349.94 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 117.0 | 2606.65 MB/s | 320 | 1 |
| invalid_utf8_dense | 723.1 | 165.96 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 11.6 | 1383.81 MB/s | 0 | 0 |
| sentence_clean | 21.3 | 2067.34 MB/s | 0 | 0 |
| url_clean | 20.9 | 2482.68 MB/s | 0 | 0 |
| log_line_clean | 30.6 | 10988.28 MB/s | 0 | 0 |
| path_with_backslash | 69.2 | 534.27 MB/s | 0 | 0 |
| json_in_json | 109.6 | 383.19 MB/s | 0 | 0 |
| prose_with_quotes | 45.0 | 844.47 MB/s | 0 | 0 |
| control_bytes | 66.5 | 360.71 MB/s | 0 | 0 |
| mostly_clean_one_quote | 36.0 | 8481.88 MB/s | 0 | 0 |
| unicode_clean | 250.1 | 943.70 MB/s | 0 | 0 |
| unicode_with_quotes | 106.4 | 592.11 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 54.3 | 5614.02 MB/s | 0 | 0 |
| invalid_utf8_dense | 580.5 | 206.73 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2388.0 | 4204.89 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2226.0 | 4511.40 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2142.0 | 4686.80 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12101.0 | 829.78 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 696.3 | 2600.97 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1441.0 | 1256.67 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.7 | 4273.30 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7824.74 MB/s | 0 | 0 |
| url_clean | 5.6 | 9247.15 MB/s | 0 | 0 |
| log_line_clean | 8.7 | 38428.11 MB/s | 0 | 0 |
| path_escaped | 77.1 | 557.75 MB/s | 48 | 1 |
| json_in_json | 102.3 | 527.88 MB/s | 64 | 1 |
| prose_with_quotes | 59.6 | 687.64 MB/s | 48 | 1 |
| unicode_heavy | 4.7 | 6414.92 MB/s | 0 | 0 |
| unicode_escaped_dense | 298.9 | 642.34 MB/s | 192 | 1 |
| mostly_clean_one_escape | 87.8 | 3484.54 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.1 | 3946.71 MB/s | 0 | 0 |
| sentence_clean | 5.9 | 7420.63 MB/s | 0 | 0 |
| url_clean | 5.9 | 8745.42 MB/s | 0 | 0 |
| log_line_clean | 9.4 | 35866.93 MB/s | 0 | 0 |
| path_escaped | 54.8 | 785.35 MB/s | 0 | 0 |
| json_in_json | 78.5 | 687.56 MB/s | 0 | 0 |
| prose_with_quotes | 39.3 | 1044.29 MB/s | 0 | 0 |
| unicode_heavy | 5.0 | 6015.16 MB/s | 0 | 0 |
| unicode_escaped_dense | 258.1 | 743.76 MB/s | 0 | 0 |
| mostly_clean_one_escape | 23.8 | 12867.73 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 4.7 | — | 0 | 0 |
| 3digit | 5.9 | — | 0 | 0 |
| 5digit | 4.7 | — | 0 | 0 |
| 10digit | 6.6 | — | 0 | 0 |
| 13digit | 6.6 | — | 0 | 0 |
| 16digit | 6.5 | — | 0 | 0 |
| 19digit | 8.5 | — | 0 | 0 |
| neg10digit | 6.5 | — | 0 | 0 |
| 20digit_overflow | 7.8 | — | 0 | 0 |
| notanint | 4.4 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 5.3 | — | 0 | 0 |
| 10digit | 5.9 | — | 0 | 0 |
| 13digit | 5.9 | — | 0 | 0 |
| 20digit | 7.5 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 2.9 | — | 0 | 0 |
| short | 4.4 | — | 0 | 0 |
| medium | 5.3 | — | 0 | 0 |
| long | 8.7 | — | 0 | 0 |
| escaped | 35.3 | — | 8 | 1 |
| notastring | 2.5 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.8 | — | 0 | 0 |
| false | 1.1 | — | 0 | 0 |
| null | 0.9 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 2.8 | — | 0 | 0 |
| number | 2.8 | — | 0 | 0 |
| object | 2.8 | — | 0 | 0 |
| array | 2.8 | — | 0 | 0 |
| null | 3.1 | — | 0 | 0 |
| true | 3.8 | — | 0 | 0 |
| invalid | 2.8 | — | 0 | 0 |
| ws_number | 3.7 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 17.3 | — | 8 | 1 |
| medium | 25.2 | — | 32 | 1 |
| long | 41.6 | — | 128 | 1 |
| escaped | 33.3 | — | 8 | 1 |
| long_late_escape | 55.2 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1647.0 | 1336.77 MB/s | 0 | 0 |
| index | 2319.0 | 948.97 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 2319.0 | 949.01 MB/s | 0 | 0 |
| strings | 899.2 | 2559.06 MB/s | 0 | 0 |
| records | 981.0 | 2822.50 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 14.4 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 7.9 | — | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 5.0 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 75.4 | — | 0 | 0 |
| append_empty | 22.5 | — | 0 | 0 |
| replace | 48.0 | — | 0 | 0 |
| create_nested | 48.7 | — | 0 | 0 |
| overwrite_nonobject | 56.6 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 121.6 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 327.1 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 97.6 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 129.8 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 387909.0 | 2851.86 MB/s | 2293537 | 25 |
| stream | 138302.0 | 7998.86 MB/s | 66992 | 4 |
| stream_reused | 127053.0 | 8707.10 MB/s | 34 | 1 |
| stream_points | 1698453.0 | 651.34 MB/s | 66992 | 4 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| strings/inmemory | 8940.0 | 2572.84 MB/s | 0 | 0 |
| strings/stream | 17902.0 | 1284.82 MB/s | 66992 | 4 |
| records/inmemory | 8712.0 | 3236.71 MB/s | 0 | 0 |
| records/stream | 17585.0 | 1603.50 MB/s | 66992 | 4 |
| scalars/inmemory | 15607.0 | 1409.73 MB/s | 0 | 0 |
| scalars/stream | 26121.0 | 842.27 MB/s | 66992 | 4 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 89918.0 | 9774.71 MB/s | 66992 | 4 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 817.3 | 1597.99 MB/s | 32 | 1 |
| inmemory | 717.3 | 1820.72 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 1796.0 | 2007.96 MB/s | 32 | 1 |
| arrayeach | 1795.0 | 2009.36 MB/s | 32 | 1 |
| inmemory | 1404.0 | 2568.61 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 195984.0 | 6078.39 MB/s | 56 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2254.0 | 1227.17 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2196.0 | 1259.79 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 25.5 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 227.7 | 816.92 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 895.6 | 2569.26 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 815.1 | 2221.89 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 801.2 | 2260.30 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 764.6 | 2368.58 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 964.0 | 2872.27 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1639.0 | 1343.26 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3829.0 | 346.04 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 85.3 | — | 24 | 1 |
| arena | 75.4 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 31176.0 | 629.71 MB/s | 0 | 0 |
| kernel/sep"," | 12897.0 | 1522.23 MB/s | 0 | 0 |
| scalar/sep",_" | 32872.0 | 718.87 MB/s | 0 | 0 |
| kernel/sep",_" | 19287.0 | 1225.21 MB/s | 0 | 0 |
| kernel-only | 12896.0 | 1522.37 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 9057.0 | 892.08 MB/s | 0 | 0 |
| "12," | 10259.0 | 1177.45 MB/s | 0 | 0 |
| "123," | 9707.0 | 1656.52 MB/s | 0 | 0 |
| "1234," | 12909.0 | 1555.47 MB/s | 0 | 0 |
| "123456," | 16463.0 | 1705.61 MB/s | 0 | 0 |
| "1234567," | 10086.0 | 3180.61 MB/s | 0 | 0 |
| "1234,_" | 19792.0 | 1216.57 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 4.4 | — | 0 | 0 |
| d03 | 5.0 | — | 0 | 0 |
| d05 | 5.0 | — | 0 | 0 |
| d08 | 5.0 | — | 0 | 0 |
| d10 | 6.5 | — | 0 | 0 |
| d13 | 6.5 | — | 0 | 0 |
| d16 | 6.6 | — | 0 | 0 |
| d19 | 8.8 | — | 0 | 0 |
| d20 | 8.4 | — | 0 | 0 |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/whole | 12269.0 | 750.60 MB/s | 0 | 0 |
| records/4k | 12217.0 | 753.80 MB/s | 0 | 0 |
| strings/whole | 4048.0 | 2767.75 MB/s | 0 | 0 |
| strings/4k | 4154.0 | 2697.33 MB/s | 0 | 0 |
| numbers/whole | 275.4 | 30507.14 MB/s | 0 | 0 |
| numbers/4k | 356.5 | 23570.21 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 542.7 | — | 0 | 0 |
| canada | 412.3 | — | 0 | 0 |
| mesh | 362.3 | — | 0 | 0 |
| array | 297.8 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 574.7 | — | 0 | 0 |
| canada | 535.7 | — | 0 | 0 |
| mesh | 397.7 | — | 0 | 0 |
| array | 316.9 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3776.0 | 6354.86 MB/s | 0 | 0 |
| numberObj/goloop | 1324.0 | 7702.19 MB/s | 0 | 0 |
| nestedMixed/goloop | 2247.0 | 4807.55 MB/s | 0 | 0 |
| stringObj/avx2 | 2117.0 | 11335.99 MB/s | 0 | 0 |
| numberObj/avx2 | 769.1 | 13259.31 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1347.0 | 8021.26 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 7678.0 | 3125.50 MB/s | 0 | 0 |
| stringObj/dispatch | 2119.0 | 11323.56 MB/s | 0 | 0 |
| numberObj/current | 6180.0 | 1650.12 MB/s | 0 | 0 |
| numberObj/dispatch | 767.6 | 13285.25 MB/s | 0 | 0 |
| numberArr/current | 236.1 | 27952.92 MB/s | 0 | 0 |
| numberArr/dispatch | 238.7 | 27653.34 MB/s | 0 | 0 |
| nestedMixed/current | 16957.0 | 636.97 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1354.0 | 7978.42 MB/s | 0 | 0 |
