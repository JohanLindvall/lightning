# lightning main-module benchmarks

- generated 2026-10-13T08:06:35Z
- go version go1.26.8 linux/amd64
- cpu: INTEL(R) XEON(R) PLATINUM 8573C (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 154.6 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 23.6 | 678.77 MB/s | 16 | 1 |
| sentence_clean | 35.8 | 1227.76 MB/s | 48 | 1 |
| url_clean | 34.0 | 1530.07 MB/s | 64 | 1 |
| log_line_clean | 93.4 | 3598.32 MB/s | 352 | 1 |
| path_with_backslash | 102.6 | 360.73 MB/s | 56 | 2 |
| json_in_json | 148.8 | 282.17 MB/s | 72 | 2 |
| prose_with_quotes | 85.1 | 446.44 MB/s | 64 | 2 |
| control_bytes | 100.5 | 238.73 MB/s | 56 | 2 |
| mostly_clean_one_quote | 102.3 | 2981.08 MB/s | 320 | 1 |
| unicode_clean | 257.0 | 918.18 MB/s | 240 | 1 |
| unicode_with_quotes | 156.1 | 403.52 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 114.5 | 2664.11 MB/s | 320 | 1 |
| invalid_utf8_dense | 641.1 | 187.19 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.4 | 1896.30 MB/s | 0 | 0 |
| sentence_clean | 16.2 | 2715.34 MB/s | 0 | 0 |
| url_clean | 16.2 | 3214.82 MB/s | 0 | 0 |
| log_line_clean | 24.4 | 13781.73 MB/s | 0 | 0 |
| path_with_backslash | 55.5 | 666.27 MB/s | 0 | 0 |
| json_in_json | 86.5 | 485.37 MB/s | 0 | 0 |
| prose_with_quotes | 35.3 | 1075.41 MB/s | 0 | 0 |
| control_bytes | 51.1 | 469.22 MB/s | 0 | 0 |
| mostly_clean_one_quote | 27.4 | 11147.74 MB/s | 0 | 0 |
| unicode_clean | 211.3 | 1116.89 MB/s | 0 | 0 |
| unicode_with_quotes | 90.9 | 693.08 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 45.4 | 6713.53 MB/s | 0 | 0 |
| invalid_utf8_dense | 493.5 | 243.18 MB/s | 0 | 0 |

## ValidEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean/value | 37.8 | 23771.22 MB/s | 0 | 0 |
| clean/key | 43.0 | 21059.73 MB/s | 0 | 0 |
| single/value | 331.0 | 1552.64 MB/s | 0 | 0 |
| single/key | 378.4 | 1376.80 MB/s | 0 | 0 |
| unicode/value | 315.9 | 2437.25 MB/s | 0 | 0 |
| unicode/key | 324.5 | 2394.76 MB/s | 0 | 0 |
| surrogates/value | 315.7 | 2438.81 MB/s | 0 | 0 |
| surrogates/key | 326.2 | 2382.28 MB/s | 0 | 0 |
| mixed/value | 378.8 | 1948.46 MB/s | 0 | 0 |
| mixed/key | 387.5 | 1922.74 MB/s | 0 | 0 |
| sparse/value | 45.6 | 19798.04 MB/s | 0 | 0 |
| sparse/key | 51.3 | 17730.95 MB/s | 0 | 0 |

## DecodeAnyEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean | 213.5 | 4205.81 MB/s | 912 | 2 |
| single | 731.0 | 703.13 MB/s | 357 | 1 |
| unicode | 769.8 | 1000.32 MB/s | 1040 | 2 |
| surrogates | 704.0 | 1093.72 MB/s | 528 | 1 |
| mixed | 804.4 | 917.51 MB/s | 1040 | 2 |
| sparse | 282.2 | 3199.36 MB/s | 1040 | 2 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1519.0 | 6610.84 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1381.0 | 7268.85 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1340.0 | 7493.32 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 9021.0 | 1113.01 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 537.1 | 3371.54 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 941.5 | 1923.43 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.2 | 7150.81 MB/s | 0 | 0 |
| sentence_clean | 4.2 | 10558.62 MB/s | 0 | 0 |
| url_clean | 4.2 | 12443.64 MB/s | 0 | 0 |
| log_line_clean | 8.1 | 41732.52 MB/s | 0 | 0 |
| path_escaped | 74.2 | 579.61 MB/s | 48 | 1 |
| json_in_json | 97.1 | 556.01 MB/s | 64 | 1 |
| prose_with_quotes | 61.0 | 671.60 MB/s | 48 | 1 |
| unicode_heavy | 2.8 | 10543.87 MB/s | 0 | 0 |
| unicode_escaped_dense | 245.3 | 782.67 MB/s | 192 | 1 |
| mostly_clean_one_escape | 85.4 | 3582.90 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.6 | 6239.08 MB/s | 0 | 0 |
| sentence_clean | 4.3 | 10216.49 MB/s | 0 | 0 |
| url_clean | 4.3 | 11962.26 MB/s | 0 | 0 |
| log_line_clean | 8.0 | 41758.22 MB/s | 0 | 0 |
| path_escaped | 57.5 | 747.56 MB/s | 0 | 0 |
| json_in_json | 77.8 | 694.14 MB/s | 0 | 0 |
| prose_with_quotes | 42.7 | 959.79 MB/s | 0 | 0 |
| unicode_heavy | 3.1 | 9797.31 MB/s | 0 | 0 |
| unicode_escaped_dense | 206.2 | 931.03 MB/s | 0 | 0 |
| mostly_clean_one_escape | 18.2 | 16823.11 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 4.5 | — | 0 | 0 |
| 3digit | 6.2 | — | 0 | 0 |
| 5digit | 3.7 | — | 0 | 0 |
| 10digit | 5.1 | — | 0 | 0 |
| 13digit | 5.1 | — | 0 | 0 |
| 16digit | 5.1 | — | 0 | 0 |
| 19digit | 6.2 | — | 0 | 0 |
| neg10digit | 5.2 | — | 0 | 0 |
| 20digit_overflow | 6.0 | — | 0 | 0 |
| notanint | 4.1 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 5.8 | — | 0 | 0 |
| 10digit | 4.4 | — | 0 | 0 |
| 13digit | 4.4 | — | 0 | 0 |
| 20digit | 5.8 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 2.1 | — | 0 | 0 |
| short | 2.8 | — | 0 | 0 |
| medium | 3.3 | — | 0 | 0 |
| long | 6.7 | — | 0 | 0 |
| escaped | 28.4 | — | 8 | 1 |
| notastring | 1.8 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.6 | — | 0 | 0 |
| false | 0.9 | — | 0 | 0 |
| null | 1.1 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 2.2 | — | 0 | 0 |
| number | 2.1 | — | 0 | 0 |
| object | 2.1 | — | 0 | 0 |
| array | 2.2 | — | 0 | 0 |
| null | 2.4 | — | 0 | 0 |
| true | 2.6 | — | 0 | 0 |
| invalid | 2.0 | — | 0 | 0 |
| ws_number | 4.0 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 12.4 | — | 8 | 1 |
| medium | 19.2 | — | 32 | 1 |
| long | 37.2 | — | 128 | 1 |
| escaped | 27.1 | — | 8 | 1 |
| long_late_escape | 49.3 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1644.0 | 1339.04 MB/s | 0 | 0 |
| index | 1781.0 | 1235.64 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1749.0 | 1258.31 MB/s | 0 | 0 |
| strings | 694.5 | 3313.31 MB/s | 0 | 0 |
| records | 540.1 | 5126.91 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 15.6 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 8.4 | — | 0 | 0 |

## DecodeAnyArrays

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| numbers/0 | 8.3 | 240.37 MB/s | 0 | 0 |
| numbers/1 | 80.9 | 111.28 MB/s | 48 | 3 |
| numbers/4 | 176.2 | 187.34 MB/s | 120 | 6 |
| numbers/16 | 485.7 | 265.59 MB/s | 408 | 18 |
| numbers/17 | 570.6 | 240.10 MB/s | 672 | 19 |
| numbers/256 | 7139.0 | 287.03 MB/s | 10904 | 261 |
| numbers/4096 | 126350.0 | 259.35 MB/s | 272281 | 4107 |
| strings/0 | 8.5 | 235.63 MB/s | 0 | 0 |
| strings/1 | 94.5 | 190.49 MB/s | 72 | 4 |
| strings/4 | 230.9 | 298.78 MB/s | 216 | 10 |
| strings/16 | 764.5 | 357.10 MB/s | 792 | 34 |
| strings/17 | 858.7 | 337.72 MB/s | 1080 | 36 |
| strings/256 | 12244.0 | 355.51 MB/s | 17048 | 517 |
| strings/4096 | 216538.0 | 321.57 MB/s | 370584 | 8203 |
| records/0 | 8.7 | 228.85 MB/s | 0 | 0 |
| records/1 | 415.2 | 103.58 MB/s | 488 | 12 |
| records/4 | 1462.0 | 115.62 MB/s | 1880 | 42 |
| records/16 | 5637.0 | 119.39 MB/s | 7448 | 162 |
| records/17 | 6117.0 | 116.88 MB/s | 8152 | 172 |
| records/256 | 91630.0 | 117.35 MB/s | 123544 | 2565 |
| records/4096 | 1716134.0 | 100.24 MB/s | 2074533 | 40971 |

## ValidShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar | 31.6 | 221.27 MB/s | 0 | 0 |
| empty_object | 17.0 | 117.81 MB/s | 0 | 0 |
| record | 795.8 | 2275.72 MB/s | 0 | 0 |
| pretty_record | 803.0 | 2562.91 MB/s | 0 | 0 |
| records | 1905.0 | 1453.64 MB/s | 0 | 0 |
| strings | 599.8 | 3836.20 MB/s | 0 | 0 |
| numbers | 450.7 | 4883.55 MB/s | 0 | 0 |
| escapes | 1739.0 | 1955.64 MB/s | 0 | 0 |
| deep | 889.7 | 288.86 MB/s | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3.2 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 54.7 | — | 0 | 0 |
| append_empty | 20.4 | — | 0 | 0 |
| replace | 34.4 | — | 0 | 0 |
| create_nested | 41.2 | — | 0 | 0 |
| overwrite_nonobject | 43.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 108.8 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 263.2 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 76.6 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 107.9 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 393912.0 | 2808.40 MB/s | 2293536 | 25 |
| stream | 91744.0 | 12058.12 MB/s | 67136 | 5 |
| stream_reused | 79365.0 | 13938.86 MB/s | 33 | 1 |
| stream_points | 763873.0 | 1448.23 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 15565.0 | 1413.51 MB/s | 0 | 0 |
| scalars/stream | 31408.0 | 700.49 MB/s | 67120 | 4 |
| scalars/stream_reused | 19272.0 | 1141.58 MB/s | 32 | 1 |
| strings/inmemory | 6208.0 | 3704.97 MB/s | 0 | 0 |
| strings/stream | 16530.0 | 1391.47 MB/s | 67120 | 4 |
| strings/stream_reused | 6339.0 | 3628.70 MB/s | 32 | 1 |
| records/inmemory | 5477.0 | 5148.35 MB/s | 0 | 0 |
| records/stream | 15005.0 | 1879.19 MB/s | 67120 | 4 |
| records/stream_reused | 5978.0 | 4717.11 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 61252.0 | 14349.32 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 725.9 | 1799.18 MB/s | 32 | 1 |
| inmemory | 670.0 | 1949.31 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 1134.0 | 3180.82 MB/s | 32 | 1 |
| arrayeach | 1137.0 | 3173.31 MB/s | 32 | 1 |
| inmemory | 1040.0 | 3467.77 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 150135.0 | 7934.62 MB/s | 51 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1875.0 | 1475.12 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1766.0 | 1565.83 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 19.6 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 151.6 | 1226.83 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 689.2 | 3338.64 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 683.8 | 2648.32 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 648.6 | 2792.03 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 690.5 | 2622.81 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 546.4 | 5067.31 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1660.0 | 1325.58 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2487.0 | 532.85 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 85.3 | — | 24 | 1 |
| arena | 85.0 | — | 24 | 0 |

## ReadIntShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 849.0 | — | 0 | 0 |
| d03 | 1092.0 | — | 0 | 0 |
| d06 | 1425.0 | — | 0 | 0 |
| d09 | 1081.0 | — | 0 | 0 |
| d10 | 1142.0 | — | 0 | 0 |
| d13 | 1495.0 | — | 0 | 0 |
| d18 | 2064.0 | — | 0 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 21.4 | — | 64 | 0 |
| make | 21.9 | — | 64 | 1 |

## FloatRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 27969.0 | 2291.73 MB/s | 0 | 0 |
| short_space | 29811.0 | 2418.42 MB/s | 0 | 0 |
| digits6 | 24377.0 | 1480.80 MB/s | 0 | 0 |
| long17 | 47955.0 | 1670.26 MB/s | 0 | 0 |
| long16 | 36227.0 | 1990.14 MB/s | 0 | 0 |

## DecodeFloat64Array

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short3/scalar | 37.5 | — | 0 | 0 |
| short3/kernel | 21.1 | — | 0 | 0 |
| coord2/scalar | 33.8 | — | 0 | 0 |
| coord2/kernel | 23.4 | — | 0 | 0 |
| pretty3/scalar | 45.9 | — | 0 | 0 |
| pretty3/kernel | 27.2 | — | 0 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 26479.0 | 741.42 MB/s | 0 | 0 |
| kernel/sep"," | 8384.0 | 2341.68 MB/s | 0 | 0 |
| scalar/sep",_" | 29683.0 | 796.12 MB/s | 0 | 0 |
| kernel/sep",_" | 8725.0 | 2708.50 MB/s | 0 | 0 |
| kernel-only | 8372.0 | 2345.04 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 7016.0 | 1151.68 MB/s | 0 | 0 |
| "12," | 7356.0 | 1642.23 MB/s | 0 | 0 |
| "123," | 7622.0 | 2109.78 MB/s | 0 | 0 |
| "1234," | 8041.0 | 2497.30 MB/s | 0 | 0 |
| "123456," | 8712.0 | 3222.99 MB/s | 0 | 0 |
| "1234567," | 8963.0 | 3578.99 MB/s | 0 | 0 |
| "1234,_" | 8621.0 | 2793.00 MB/s | 0 | 0 |

## DecodeIntSliceShort

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 2/scalar | 33.4 | — | 16 | 1 |
| 2/kernel | 28.8 | — | 16 | 1 |
| 3/scalar | 40.8 | — | 24 | 1 |
| 3/kernel | 31.9 | — | 24 | 1 |
| 4/scalar | 48.3 | — | 32 | 1 |
| 4/kernel | 35.0 | — | 32 | 1 |
| 6/scalar | 63.9 | — | 48 | 1 |
| 6/kernel | 42.3 | — | 48 | 1 |
| 8/scalar | 79.3 | — | 64 | 1 |
| 8/kernel | 49.1 | — | 64 | 1 |
| 12/scalar | 110.7 | — | 96 | 1 |
| 12/kernel | 60.9 | — | 96 | 1 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 4.5 | — | 0 | 0 |
| d03 | 6.6 | — | 0 | 0 |
| d05 | 3.7 | — | 0 | 0 |
| d08 | 3.7 | — | 0 | 0 |
| d10 | 5.2 | — | 0 | 0 |
| d13 | 5.2 | — | 0 | 0 |
| d16 | 5.2 | — | 0 | 0 |
| d19 | 6.5 | — | 0 | 0 |
| d20 | 6.4 | — | 0 | 0 |

## Float64Points

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| canada/ring6 | 169.4 | — | — | — |
| canada/ring1000 | 29024.0 | — | — | — |
| geometry/ring6 | 141.5 | — | — | — |
| geometry/ring1000 | 21995.0 | — | — | — |
| citylots/ring6 | 184.5 | — | — | — |
| citylots/ring1000 | 31393.0 | — | — | — |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/whole | 11261.0 | 817.75 MB/s | 0 | 0 |
| records/4k | 11274.0 | 816.81 MB/s | 0 | 0 |
| strings/whole | 4549.0 | 2462.98 MB/s | 0 | 0 |
| strings/4k | 4595.0 | 2438.61 MB/s | 0 | 0 |
| numbers/whole | 79.5 | 105759.88 MB/s | 0 | 0 |
| numbers/4k | 86.3 | 97349.51 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 374.2 | — | 0 | 0 |
| canada | 281.6 | — | 0 | 0 |
| mesh | 230.3 | — | 0 | 0 |
| array | 187.7 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 436.7 | — | 0 | 0 |
| canada | 403.6 | — | 0 | 0 |
| mesh | 318.3 | — | 0 | 0 |
| array | 301.8 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3107.0 | 7723.51 MB/s | 0 | 0 |
| numberObj/goloop | 1030.0 | 9900.87 MB/s | 0 | 0 |
| nestedMixed/goloop | 1581.0 | 6830.88 MB/s | 0 | 0 |
| stringObj/avx2 | 1556.0 | 15422.71 MB/s | 0 | 0 |
| numberObj/avx2 | 525.4 | 19409.97 MB/s | 0 | 0 |
| nestedMixed/avx2 | 980.5 | 11015.57 MB/s | 0 | 0 |
| stringObj/avx512 | 1204.0 | 19930.75 MB/s | 0 | 0 |
| numberObj/avx512 | 356.8 | 28584.33 MB/s | 0 | 0 |
| nestedMixed/avx512 | 1039.0 | 10398.49 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10724.0 | 2237.79 MB/s | 0 | 0 |
| stringObj/dispatch | 1197.0 | 20054.10 MB/s | 0 | 0 |
| numberObj/current | 3843.0 | 2653.71 MB/s | 0 | 0 |
| numberObj/dispatch | 359.9 | 28335.15 MB/s | 0 | 0 |
| numberArr/current | 62.0 | 106409.68 MB/s | 0 | 0 |
| numberArr/dispatch | 64.9 | 101730.35 MB/s | 0 | 0 |
| nestedMixed/current | 15302.0 | 705.87 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1037.0 | 10414.57 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 10.6 | 1703.18 MB/s | 0 | 0 |
| record | 6.9 | 7871.98 MB/s | 0 | 0 |
| tiny | 6.9 | 1015.93 MB/s | 0 | 0 |
| twoBlock | 9.8 | 8965.21 MB/s | 0 | 0 |

## SkipSmallAtEnd

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 10.8 | 1671.63 MB/s | 0 | 0 |
| record | 7.3 | 7412.22 MB/s | 0 | 0 |
| tiny | 7.5 | 934.96 MB/s | 0 | 0 |
| twoBlock | 10.6 | 8284.13 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 13.6 | 1326.31 MB/s | 0 | 0 |
| record | 31.7 | 1704.80 MB/s | 0 | 0 |
| tiny | 9.9 | 704.41 MB/s | 0 | 0 |
| twoBlock | 18.4 | 4768.76 MB/s | 0 | 0 |

## IndexStructural

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 8 | 3.6 | 2240.44 MB/s | 0 | 0 |
| 24 | 5.5 | 4332.31 MB/s | 0 | 0 |
| 40 | 5.5 | 7210.55 MB/s | 0 | 0 |
| 72 | 5.6 | 12949.54 MB/s | 0 | 0 |
| 136 | 6.2 | 21931.67 MB/s | 0 | 0 |
| 520 | 8.3 | 62392.51 MB/s | 0 | 0 |
| 4104 | 38.4 | 106785.39 MB/s | 0 | 0 |
| 20000 | 187.1 | 106866.77 MB/s | 0 | 0 |

## ValidNumbers

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| flat | 10599.0 | 6038.66 MB/s | 0 | 0 |
| canada | 10891.0 | 3773.77 MB/s | 0 | 0 |
| geometry | 10899.0 | 3590.61 MB/s | 0 | 0 |
| objects | 25949.0 | 789.72 MB/s | 0 | 0 |
