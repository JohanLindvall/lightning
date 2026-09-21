# lightning main-module benchmarks

- generated 2026-09-21T08:06:01Z
- go version go1.26.8 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 198.0 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.9 | 535.59 MB/s | 16 | 1 |
| sentence_clean | 37.8 | 1164.23 MB/s | 48 | 1 |
| url_clean | 41.3 | 1257.95 MB/s | 64 | 1 |
| log_line_clean | 119.9 | 2801.78 MB/s | 352 | 1 |
| path_with_backslash | 114.0 | 324.69 MB/s | 56 | 2 |
| json_in_json | 151.1 | 277.98 MB/s | 72 | 2 |
| prose_with_quotes | 92.4 | 411.16 MB/s | 64 | 2 |
| control_bytes | 110.0 | 218.28 MB/s | 56 | 2 |
| mostly_clean_one_quote | 139.4 | 2187.95 MB/s | 320 | 1 |
| unicode_clean | 273.4 | 863.09 MB/s | 240 | 1 |
| unicode_with_quotes | 160.1 | 393.44 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 148.4 | 2054.57 MB/s | 320 | 1 |
| invalid_utf8_dense | 621.3 | 193.15 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.5 | 2123.72 MB/s | 0 | 0 |
| sentence_clean | 15.1 | 2924.05 MB/s | 0 | 0 |
| url_clean | 10.6 | 4900.97 MB/s | 0 | 0 |
| log_line_clean | 33.6 | 10010.88 MB/s | 0 | 0 |
| path_with_backslash | 53.1 | 696.71 MB/s | 0 | 0 |
| json_in_json | 84.2 | 498.72 MB/s | 0 | 0 |
| prose_with_quotes | 31.2 | 1217.09 MB/s | 0 | 0 |
| control_bytes | 45.4 | 528.12 MB/s | 0 | 0 |
| mostly_clean_one_quote | 34.9 | 8730.85 MB/s | 0 | 0 |
| unicode_clean | 224.9 | 1049.25 MB/s | 0 | 0 |
| unicode_with_quotes | 86.2 | 730.58 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 49.9 | 6111.89 MB/s | 0 | 0 |
| invalid_utf8_dense | 432.5 | 277.48 MB/s | 0 | 0 |

## ValidEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean/value | 47.2 | 19005.05 MB/s | 0 | 0 |
| clean/key | 53.2 | 17007.65 MB/s | 0 | 0 |
| single/value | 369.5 | 1390.90 MB/s | 0 | 0 |
| single/key | 374.3 | 1391.94 MB/s | 0 | 0 |
| unicode/value | 366.4 | 2101.65 MB/s | 0 | 0 |
| unicode/key | 371.0 | 2094.44 MB/s | 0 | 0 |
| surrogates/value | 366.5 | 2101.13 MB/s | 0 | 0 |
| surrogates/key | 370.9 | 2095.00 MB/s | 0 | 0 |
| mixed/value | 568.2 | 1298.79 MB/s | 0 | 0 |
| mixed/key | 588.5 | 1265.93 MB/s | 0 | 0 |
| sparse/value | 54.3 | 16627.63 MB/s | 0 | 0 |
| sparse/key | 59.3 | 15336.03 MB/s | 0 | 0 |

## DecodeAnyEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean | 280.9 | 3197.06 MB/s | 912 | 2 |
| single | 887.7 | 578.99 MB/s | 357 | 1 |
| unicode | 966.0 | 797.07 MB/s | 1040 | 2 |
| surrogates | 903.7 | 852.02 MB/s | 528 | 1 |
| mixed | 1138.0 | 648.71 MB/s | 1040 | 2 |
| sparse | 386.1 | 2339.00 MB/s | 1040 | 2 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2709.0 | 3706.96 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2698.0 | 3722.31 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2696.0 | 3724.35 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10341.0 | 970.98 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 742.4 | 2439.25 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1059.0 | 1710.30 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.6 | 6254.12 MB/s | 0 | 0 |
| sentence_clean | 4.7 | 9284.59 MB/s | 0 | 0 |
| url_clean | 4.7 | 10972.84 MB/s | 0 | 0 |
| log_line_clean | 10.5 | 31960.45 MB/s | 0 | 0 |
| path_escaped | 79.8 | 538.63 MB/s | 48 | 1 |
| json_in_json | 106.9 | 505.30 MB/s | 64 | 1 |
| prose_with_quotes | 66.8 | 614.04 MB/s | 48 | 1 |
| unicode_heavy | 3.3 | 9225.94 MB/s | 0 | 0 |
| unicode_escaped_dense | 294.6 | 651.63 MB/s | 192 | 1 |
| mostly_clean_one_escape | 125.3 | 2441.60 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.1 | 5156.05 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8290.78 MB/s | 0 | 0 |
| url_clean | 5.3 | 9798.49 MB/s | 0 | 0 |
| log_line_clean | 11.1 | 30367.80 MB/s | 0 | 0 |
| path_escaped | 43.9 | 978.78 MB/s | 0 | 0 |
| json_in_json | 68.4 | 789.39 MB/s | 0 | 0 |
| prose_with_quotes | 32.7 | 1254.11 MB/s | 0 | 0 |
| unicode_heavy | 3.8 | 7827.08 MB/s | 0 | 0 |
| unicode_escaped_dense | 220.9 | 869.31 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.4 | 12564.37 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 2.9 | — | 0 | 0 |
| 3digit | 3.6 | — | 0 | 0 |
| 5digit | 4.0 | — | 0 | 0 |
| 10digit | 5.0 | — | 0 | 0 |
| 13digit | 5.0 | — | 0 | 0 |
| 16digit | 5.0 | — | 0 | 0 |
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
| empty | 2.4 | — | 0 | 0 |
| short | 3.4 | — | 0 | 0 |
| medium | 4.1 | — | 0 | 0 |
| long | 8.6 | — | 0 | 0 |
| escaped | 37.8 | — | 8 | 1 |
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
| short | 16.8 | — | 8 | 1 |
| medium | 28.6 | — | 32 | 1 |
| long | 50.8 | — | 128 | 1 |
| escaped | 35.8 | — | 8 | 1 |
| long_late_escape | 102.0 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1665.0 | 1321.74 MB/s | 0 | 0 |
| index | 1786.0 | 1232.63 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1783.0 | 1234.58 MB/s | 0 | 0 |
| strings | 1018.0 | 2261.37 MB/s | 0 | 0 |
| records | 984.6 | 2812.25 MB/s | 0 | 0 |

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

## DecodeAnyArrays

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| numbers/0 | 6.5 | 308.29 MB/s | 0 | 0 |
| numbers/1 | 103.9 | 86.59 MB/s | 48 | 3 |
| numbers/4 | 240.6 | 137.13 MB/s | 120 | 6 |
| numbers/16 | 653.1 | 197.52 MB/s | 408 | 18 |
| numbers/17 | 730.4 | 187.56 MB/s | 672 | 19 |
| numbers/256 | 9292.0 | 220.52 MB/s | 10904 | 261 |
| numbers/4096 | 159470.0 | 205.49 MB/s | 272281 | 4107 |
| strings/0 | 6.5 | 308.29 MB/s | 0 | 0 |
| strings/1 | 132.9 | 135.42 MB/s | 72 | 4 |
| strings/4 | 347.6 | 198.52 MB/s | 216 | 10 |
| strings/16 | 1078.0 | 253.19 MB/s | 792 | 34 |
| strings/17 | 1203.0 | 240.99 MB/s | 1080 | 36 |
| strings/256 | 16719.0 | 260.36 MB/s | 17048 | 517 |
| strings/4096 | 276299.0 | 252.02 MB/s | 370584 | 8203 |
| records/0 | 6.5 | 308.28 MB/s | 0 | 0 |
| records/1 | 592.7 | 72.55 MB/s | 488 | 12 |
| records/4 | 2155.0 | 78.43 MB/s | 1880 | 42 |
| records/16 | 8334.0 | 80.75 MB/s | 7448 | 162 |
| records/17 | 8826.0 | 81.01 MB/s | 8152 | 172 |
| records/256 | 130517.0 | 82.39 MB/s | 123544 | 2565 |
| records/4096 | 2383313.0 | 72.18 MB/s | 2074534 | 40971 |

## ValidShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar | 35.6 | 196.65 MB/s | 0 | 0 |
| empty_object | 24.7 | 81.07 MB/s | 0 | 0 |
| record | 923.1 | 1961.78 MB/s | 0 | 0 |
| pretty_record | 1023.0 | 2011.37 MB/s | 0 | 0 |
| records | 2218.0 | 1248.45 MB/s | 0 | 0 |
| strings | 847.0 | 2716.67 MB/s | 0 | 0 |
| numbers | 3654.0 | 602.33 MB/s | 0 | 0 |
| escapes | 2351.0 | 1446.33 MB/s | 0 | 0 |
| deep | 533.8 | 481.49 MB/s | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 4.0 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 63.1 | — | 0 | 0 |
| append_empty | 18.0 | — | 0 | 0 |
| replace | 47.0 | — | 0 | 0 |
| create_nested | 41.4 | — | 0 | 0 |
| overwrite_nonobject | 49.8 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 108.3 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 281.8 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 104.9 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 138.2 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 454417.0 | 2434.47 MB/s | 2293536 | 25 |
| stream | 231773.0 | 4773.05 MB/s | 67136 | 5 |
| stream_reused | 202917.0 | 5451.79 MB/s | 35 | 1 |
| stream_points | 1406521.0 | 786.52 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| strings/inmemory | 10471.0 | 2196.66 MB/s | 0 | 0 |
| strings/stream | 20115.0 | 1143.48 MB/s | 67120 | 4 |
| strings/stream_reused | 12213.0 | 1883.36 MB/s | 32 | 1 |
| records/inmemory | 9906.0 | 2846.65 MB/s | 0 | 0 |
| records/stream | 19945.0 | 1413.78 MB/s | 67120 | 4 |
| records/stream_reused | 11013.0 | 2560.36 MB/s | 32 | 1 |
| scalars/inmemory | 17113.0 | 1285.65 MB/s | 0 | 0 |
| scalars/stream | 28160.0 | 781.28 MB/s | 67120 | 4 |
| scalars/stream_reused | 19642.0 | 1120.12 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 199218.0 | 4411.88 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 986.7 | 1323.57 MB/s | 32 | 1 |
| inmemory | 799.1 | 1634.34 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 2129.0 | 1694.60 MB/s | 32 | 1 |
| arrayeach | 2102.0 | 1716.01 MB/s | 32 | 1 |
| inmemory | 1594.0 | 2262.98 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 291296.0 | 4089.53 MB/s | 69 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1986.0 | 1393.04 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1912.0 | 1446.88 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 27.9 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 209.7 | 886.96 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1019.0 | 2258.81 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 818.5 | 2212.68 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 822.8 | 2201.07 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 821.0 | 2205.96 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 990.9 | 2794.35 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1661.0 | 1324.94 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2674.0 | 495.52 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 87.9 | — | 24 | 1 |
| arena | 78.8 | — | 24 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 23.8 | — | 64 | 0 |
| make | 46.2 | — | 64 | 1 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 23700.0 | 828.37 MB/s | 0 | 0 |
| kernel/sep"," | 8539.0 | 2299.23 MB/s | 0 | 0 |
| scalar/sep",_" | 25908.0 | 912.12 MB/s | 0 | 0 |
| kernel/sep",_" | 11020.0 | 2144.35 MB/s | 0 | 0 |
| kernel-only | 8571.0 | 2290.46 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 7582.0 | 1065.71 MB/s | 0 | 0 |
| "12," | 7222.0 | 1672.61 MB/s | 0 | 0 |
| "123," | 7537.0 | 2133.44 MB/s | 0 | 0 |
| "1234," | 8169.0 | 2458.06 MB/s | 0 | 0 |
| "123456," | 8698.0 | 3228.31 MB/s | 0 | 0 |
| "1234567," | 8528.0 | 3761.82 MB/s | 0 | 0 |
| "1234,_" | 10788.0 | 2232.06 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 2.7 | — | 0 | 0 |
| d03 | 3.5 | — | 0 | 0 |
| d05 | 3.9 | — | 0 | 0 |
| d08 | 3.9 | — | 0 | 0 |
| d10 | 5.1 | — | 0 | 0 |
| d13 | 5.1 | — | 0 | 0 |
| d16 | 5.1 | — | 0 | 0 |
| d19 | 6.7 | — | 0 | 0 |
| d20 | 7.1 | — | 0 | 0 |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/whole | 12183.0 | 755.87 MB/s | 0 | 0 |
| records/4k | 12219.0 | 753.64 MB/s | 0 | 0 |
| strings/whole | 5303.0 | 2112.84 MB/s | 0 | 0 |
| strings/4k | 5316.0 | 2107.74 MB/s | 0 | 0 |
| numbers/whole | 192.3 | 43696.39 MB/s | 0 | 0 |
| numbers/4k | 219.5 | 38274.42 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 460.1 | — | 0 | 0 |
| canada | 333.9 | — | 0 | 0 |
| mesh | 272.0 | — | 0 | 0 |
| array | 223.5 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 459.3 | — | 0 | 0 |
| canada | 409.7 | — | 0 | 0 |
| mesh | 292.5 | — | 0 | 0 |
| array | 261.4 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4267.0 | 5623.57 MB/s | 0 | 0 |
| numberObj/goloop | 1776.0 | 5740.96 MB/s | 0 | 0 |
| nestedMixed/goloop | 2405.0 | 4491.49 MB/s | 0 | 0 |
| stringObj/neon | 2821.0 | 8507.17 MB/s | 0 | 0 |
| numberObj/neon | 1195.0 | 8533.34 MB/s | 0 | 0 |
| nestedMixed/neon | 1588.0 | 6802.17 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 11961.0 | 2006.36 MB/s | 0 | 0 |
| stringObj/dispatch | 2821.0 | 8505.77 MB/s | 0 | 0 |
| numberObj/current | 4419.0 | 2307.74 MB/s | 0 | 0 |
| numberObj/dispatch | 1196.0 | 8528.00 MB/s | 0 | 0 |
| numberArr/current | 150.8 | 43776.55 MB/s | 0 | 0 |
| numberArr/dispatch | 156.7 | 42135.60 MB/s | 0 | 0 |
| nestedMixed/current | 15956.0 | 676.91 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1591.0 | 6790.76 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 18.2 | 987.88 MB/s | 0 | 0 |
| record | 13.6 | 3966.39 MB/s | 0 | 0 |
| tiny | 13.6 | 515.59 MB/s | 0 | 0 |
| twoBlock | 21.0 | 4181.79 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 16.2 | 1112.53 MB/s | 0 | 0 |
| record | 43.5 | 1240.17 MB/s | 0 | 0 |
| tiny | 13.3 | 527.75 MB/s | 0 | 0 |
| twoBlock | 25.3 | 3477.52 MB/s | 0 | 0 |
