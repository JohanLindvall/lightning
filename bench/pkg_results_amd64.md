# lightning main-module benchmarks

- generated 2026-09-23T13:15:47Z
- go version go1.26.8 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 194.7 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 35.1 | 455.97 MB/s | 16 | 1 |
| sentence_clean | 44.2 | 996.18 MB/s | 48 | 1 |
| url_clean | 41.9 | 1242.19 MB/s | 64 | 1 |
| log_line_clean | 95.3 | 3524.54 MB/s | 352 | 1 |
| path_with_backslash | 142.4 | 259.92 MB/s | 56 | 2 |
| json_in_json | 180.2 | 233.11 MB/s | 72 | 2 |
| prose_with_quotes | 115.1 | 330.12 MB/s | 64 | 2 |
| control_bytes | 139.4 | 172.13 MB/s | 56 | 2 |
| mostly_clean_one_quote | 103.1 | 2958.77 MB/s | 320 | 1 |
| unicode_clean | 303.9 | 776.45 MB/s | 240 | 1 |
| unicode_with_quotes | 185.0 | 340.60 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 115.0 | 2651.35 MB/s | 320 | 1 |
| invalid_utf8_dense | 734.8 | 163.32 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 11.6 | 1383.97 MB/s | 0 | 0 |
| sentence_clean | 21.3 | 2066.53 MB/s | 0 | 0 |
| url_clean | 20.6 | 2517.61 MB/s | 0 | 0 |
| log_line_clean | 28.4 | 11827.58 MB/s | 0 | 0 |
| path_with_backslash | 71.5 | 517.45 MB/s | 0 | 0 |
| json_in_json | 109.6 | 383.25 MB/s | 0 | 0 |
| prose_with_quotes | 44.8 | 847.93 MB/s | 0 | 0 |
| control_bytes | 66.4 | 361.66 MB/s | 0 | 0 |
| mostly_clean_one_quote | 34.2 | 8909.91 MB/s | 0 | 0 |
| unicode_clean | 251.5 | 938.20 MB/s | 0 | 0 |
| unicode_with_quotes | 107.0 | 588.74 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 52.8 | 5777.43 MB/s | 0 | 0 |
| invalid_utf8_dense | 577.9 | 207.66 MB/s | 0 | 0 |

## ValidEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean/value | 42.8 | 20977.66 MB/s | 0 | 0 |
| clean/key | 47.8 | 18950.75 MB/s | 0 | 0 |
| single/value | 429.7 | 1196.12 MB/s | 0 | 0 |
| single/key | 438.8 | 1187.25 MB/s | 0 | 0 |
| unicode/value | 396.3 | 1942.87 MB/s | 0 | 0 |
| unicode/key | 400.6 | 1939.47 MB/s | 0 | 0 |
| surrogates/value | 396.3 | 1942.99 MB/s | 0 | 0 |
| surrogates/key | 400.9 | 1938.30 MB/s | 0 | 0 |
| mixed/value | 489.5 | 1507.60 MB/s | 0 | 0 |
| mixed/key | 486.6 | 1530.90 MB/s | 0 | 0 |
| sparse/value | 50.6 | 17844.32 MB/s | 0 | 0 |
| sparse/key | 56.6 | 16067.07 MB/s | 0 | 0 |

## DecodeAnyEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean | 214.0 | 4196.57 MB/s | 912 | 2 |
| single | 902.9 | 569.27 MB/s | 357 | 1 |
| unicode | 982.1 | 784.07 MB/s | 1040 | 2 |
| surrogates | 903.1 | 852.59 MB/s | 528 | 1 |
| mixed | 1012.0 | 729.44 MB/s | 1040 | 2 |
| sparse | 292.0 | 3092.25 MB/s | 1040 | 2 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2030.0 | 4947.00 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2003.0 | 5012.42 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1873.0 | 5360.46 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12071.0 | 831.83 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 646.1 | 2802.77 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1277.0 | 1418.32 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.4 | 4664.29 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7819.89 MB/s | 0 | 0 |
| url_clean | 5.6 | 9240.47 MB/s | 0 | 0 |
| log_line_clean | 8.7 | 38444.06 MB/s | 0 | 0 |
| path_escaped | 89.6 | 479.97 MB/s | 48 | 1 |
| json_in_json | 117.5 | 459.63 MB/s | 64 | 1 |
| prose_with_quotes | 70.2 | 583.63 MB/s | 48 | 1 |
| unicode_heavy | 4.4 | 6864.80 MB/s | 0 | 0 |
| unicode_escaped_dense | 324.0 | 592.50 MB/s | 192 | 1 |
| mostly_clean_one_escape | 89.0 | 3439.47 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.8 | 4262.89 MB/s | 0 | 0 |
| sentence_clean | 5.9 | 7417.02 MB/s | 0 | 0 |
| url_clean | 6.0 | 8719.96 MB/s | 0 | 0 |
| log_line_clean | 9.4 | 35846.25 MB/s | 0 | 0 |
| path_escaped | 69.9 | 614.96 MB/s | 0 | 0 |
| json_in_json | 95.0 | 568.54 MB/s | 0 | 0 |
| prose_with_quotes | 49.1 | 834.39 MB/s | 0 | 0 |
| unicode_heavy | 4.7 | 6359.64 MB/s | 0 | 0 |
| unicode_escaped_dense | 280.2 | 685.19 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.5 | 12501.07 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 4.7 | — | 0 | 0 |
| 3digit | 6.0 | — | 0 | 0 |
| 5digit | 4.7 | — | 0 | 0 |
| 10digit | 6.6 | — | 0 | 0 |
| 13digit | 6.6 | — | 0 | 0 |
| 16digit | 6.6 | — | 0 | 0 |
| 19digit | 8.5 | — | 0 | 0 |
| neg10digit | 6.6 | — | 0 | 0 |
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
| empty | 2.8 | — | 0 | 0 |
| short | 4.1 | — | 0 | 0 |
| medium | 5.0 | — | 0 | 0 |
| long | 8.1 | — | 0 | 0 |
| escaped | 39.5 | — | 8 | 1 |
| notastring | 2.5 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.8 | — | 0 | 0 |
| false | 1.1 | — | 0 | 0 |
| null | 1.2 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 2.5 | — | 0 | 0 |
| number | 2.5 | — | 0 | 0 |
| object | 2.5 | — | 0 | 0 |
| array | 2.5 | — | 0 | 0 |
| null | 3.1 | — | 0 | 0 |
| true | 3.1 | — | 0 | 0 |
| invalid | 2.5 | — | 0 | 0 |
| ws_number | 3.8 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 17.7 | — | 8 | 1 |
| medium | 25.7 | — | 32 | 1 |
| long | 41.5 | — | 128 | 1 |
| escaped | 36.3 | — | 8 | 1 |
| long_late_escape | 57.0 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1696.0 | 1297.72 MB/s | 0 | 0 |
| index | 1845.0 | 1192.66 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1825.0 | 1205.78 MB/s | 0 | 0 |
| strings | 941.8 | 2443.27 MB/s | 0 | 0 |
| records | 678.8 | 4079.06 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 15.6 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 8.1 | — | 0 | 0 |

## DecodeAnyArrays

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| numbers/0 | 10.0 | 200.08 MB/s | 0 | 0 |
| numbers/1 | 97.7 | 92.13 MB/s | 48 | 3 |
| numbers/4 | 209.3 | 157.71 MB/s | 120 | 6 |
| numbers/16 | 622.9 | 207.10 MB/s | 408 | 18 |
| numbers/17 | 723.8 | 189.29 MB/s | 672 | 19 |
| numbers/256 | 9372.0 | 218.63 MB/s | 10904 | 261 |
| numbers/4096 | 164181.0 | 199.59 MB/s | 272281 | 4107 |
| strings/0 | 10.0 | 200.42 MB/s | 0 | 0 |
| strings/1 | 118.1 | 152.39 MB/s | 72 | 4 |
| strings/4 | 282.4 | 244.36 MB/s | 216 | 10 |
| strings/16 | 946.1 | 288.55 MB/s | 792 | 34 |
| strings/17 | 1056.0 | 274.59 MB/s | 1080 | 36 |
| strings/256 | 14813.0 | 293.86 MB/s | 17048 | 517 |
| strings/4096 | 264756.0 | 263.01 MB/s | 370585 | 8203 |
| records/0 | 10.0 | 200.40 MB/s | 0 | 0 |
| records/1 | 508.8 | 84.51 MB/s | 488 | 12 |
| records/4 | 1756.0 | 96.27 MB/s | 1880 | 42 |
| records/16 | 6775.0 | 99.34 MB/s | 7448 | 162 |
| records/17 | 7292.0 | 98.05 MB/s | 8152 | 172 |
| records/256 | 109450.0 | 98.25 MB/s | 123544 | 2565 |
| records/4096 | 2059911.0 | 83.51 MB/s | 2074532 | 40971 |

## ValidShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar | 33.1 | 211.36 MB/s | 0 | 0 |
| empty_object | 25.0 | 80.06 MB/s | 0 | 0 |
| record | 797.4 | 2271.10 MB/s | 0 | 0 |
| pretty_record | 900.9 | 2284.28 MB/s | 0 | 0 |
| records | 2144.0 | 1291.74 MB/s | 0 | 0 |
| strings | 721.4 | 3189.61 MB/s | 0 | 0 |
| numbers | 670.9 | 3280.47 MB/s | 0 | 0 |
| escapes | 2219.0 | 1532.43 MB/s | 0 | 0 |
| deep | 805.9 | 318.90 MB/s | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 4.7 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 71.5 | — | 0 | 0 |
| append_empty | 23.1 | — | 0 | 0 |
| replace | 48.1 | — | 0 | 0 |
| create_nested | 52.7 | — | 0 | 0 |
| overwrite_nonobject | 55.9 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 126.6 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 311.2 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 96.6 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 135.8 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 368355.0 | 3003.25 MB/s | 2293537 | 25 |
| stream | 135063.0 | 8190.70 MB/s | 67136 | 5 |
| stream_reused | 123422.0 | 8963.23 MB/s | 34 | 1 |
| stream_points | 1139136.0 | 971.14 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 16898.0 | 1302.02 MB/s | 0 | 0 |
| scalars/stream | 30811.0 | 714.06 MB/s | 67120 | 4 |
| scalars/stream_reused | 23020.0 | 955.72 MB/s | 32 | 1 |
| strings/inmemory | 9118.0 | 2522.63 MB/s | 0 | 0 |
| strings/stream | 17607.0 | 1306.35 MB/s | 67120 | 4 |
| strings/stream_reused | 9601.0 | 2395.64 MB/s | 32 | 1 |
| records/inmemory | 6746.0 | 4180.15 MB/s | 0 | 0 |
| records/stream | 15578.0 | 1810.15 MB/s | 67120 | 4 |
| records/stream_reused | 7658.0 | 3681.93 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 91340.0 | 9622.57 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 789.7 | 1653.75 MB/s | 32 | 1 |
| inmemory | 649.6 | 2010.40 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 1438.0 | 2508.79 MB/s | 32 | 1 |
| arrayeach | 1470.0 | 2453.41 MB/s | 32 | 1 |
| inmemory | 1383.0 | 2607.60 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 194439.0 | 6126.66 MB/s | 57 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2260.0 | 1223.95 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2167.0 | 1276.42 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 25.5 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 169.2 | 1099.32 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 891.2 | 2582.03 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 704.0 | 2572.58 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 688.6 | 2630.07 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 684.3 | 2646.39 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 656.0 | 4221.21 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1704.0 | 1291.45 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2806.0 | 472.25 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 102.1 | — | 24 | 1 |
| arena | 90.5 | — | 24 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 26.0 | — | 64 | 0 |
| make | 25.8 | — | 64 | 1 |

## FloatRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 24957.0 | 2568.27 MB/s | 0 | 0 |
| short_space | 26717.0 | 2698.51 MB/s | 0 | 0 |
| digits6 | 22136.0 | 1630.69 MB/s | 0 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 31012.0 | 633.05 MB/s | 0 | 0 |
| kernel/sep"," | 9637.0 | 2037.19 MB/s | 0 | 0 |
| scalar/sep",_" | 35092.0 | 673.40 MB/s | 0 | 0 |
| kernel/sep",_" | 9935.0 | 2378.59 MB/s | 0 | 0 |
| kernel-only | 9654.0 | 2033.60 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 8022.0 | 1007.24 MB/s | 0 | 0 |
| "12," | 8309.0 | 1453.83 MB/s | 0 | 0 |
| "123," | 8662.0 | 1856.47 MB/s | 0 | 0 |
| "1234," | 9010.0 | 2228.67 MB/s | 0 | 0 |
| "123456," | 9676.0 | 2901.97 MB/s | 0 | 0 |
| "1234567," | 10741.0 | 2986.80 MB/s | 0 | 0 |
| "1234,_" | 9337.0 | 2578.83 MB/s | 0 | 0 |

## DecodeIntSliceShort

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 2/scalar | 150.6 | — | 16 | 1 |
| 2/kernel | 152.2 | — | 16 | 1 |
| 3/scalar | 160.1 | — | 24 | 1 |
| 3/kernel | 158.6 | — | 24 | 1 |
| 4/scalar | 167.6 | — | 32 | 1 |
| 4/kernel | 160.7 | — | 32 | 1 |
| 6/scalar | 184.1 | — | 48 | 1 |
| 6/kernel | 167.9 | — | 48 | 1 |
| 8/scalar | 200.3 | — | 64 | 1 |
| 8/kernel | 177.1 | — | 64 | 1 |
| 12/scalar | 234.8 | — | 96 | 1 |
| 12/kernel | 186.2 | — | 96 | 1 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 4.4 | — | 0 | 0 |
| d03 | 5.0 | — | 0 | 0 |
| d05 | 5.0 | — | 0 | 0 |
| d08 | 5.0 | — | 0 | 0 |
| d10 | 6.6 | — | 0 | 0 |
| d13 | 6.6 | — | 0 | 0 |
| d16 | 6.6 | — | 0 | 0 |
| d19 | 8.9 | — | 0 | 0 |
| d20 | 8.4 | — | 0 | 0 |

## Float64Points

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| canada/ring6 | 323.3 | — | — | — |
| canada/ring1000 | 52329.0 | — | — | — |
| geometry/ring6 | 281.8 | — | — | — |
| geometry/ring1000 | 49289.0 | — | — | — |
| citylots/ring6 | 399.5 | — | — | — |
| citylots/ring1000 | 69493.0 | — | — | — |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| numbers/whole | 166.7 | 50407.38 MB/s | 0 | 0 |
| numbers/4k | 186.1 | 45143.24 MB/s | 0 | 0 |
| records/whole | 11382.0 | 809.07 MB/s | 0 | 0 |
| records/4k | 11443.0 | 804.74 MB/s | 0 | 0 |
| strings/whole | 4806.0 | 2331.46 MB/s | 0 | 0 |
| strings/4k | 4868.0 | 2301.62 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 529.5 | — | 0 | 0 |
| canada | 408.9 | — | 0 | 0 |
| mesh | 366.4 | — | 0 | 0 |
| array | 294.3 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 560.3 | — | 0 | 0 |
| canada | 526.6 | — | 0 | 0 |
| mesh | 385.4 | — | 0 | 0 |
| array | 314.2 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3688.0 | 6506.36 MB/s | 0 | 0 |
| numberObj/goloop | 1400.0 | 7282.85 MB/s | 0 | 0 |
| nestedMixed/goloop | 2126.0 | 5081.39 MB/s | 0 | 0 |
| stringObj/avx2 | 1968.0 | 12196.01 MB/s | 0 | 0 |
| numberObj/avx2 | 717.2 | 14219.97 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1193.0 | 9052.71 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10530.0 | 2279.01 MB/s | 0 | 0 |
| stringObj/dispatch | 1963.0 | 12222.21 MB/s | 0 | 0 |
| numberObj/current | 4059.0 | 2512.56 MB/s | 0 | 0 |
| numberObj/dispatch | 720.2 | 14159.39 MB/s | 0 | 0 |
| numberArr/current | 129.5 | 50968.72 MB/s | 0 | 0 |
| numberArr/dispatch | 131.3 | 50256.49 MB/s | 0 | 0 |
| nestedMixed/current | 16054.0 | 672.78 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1196.0 | 9031.52 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 16.6 | 1082.03 MB/s | 0 | 0 |
| record | 10.6 | 5075.00 MB/s | 0 | 0 |
| tiny | 10.6 | 658.16 MB/s | 0 | 0 |
| twoBlock | 14.2 | 6186.86 MB/s | 0 | 0 |

## SkipSmallAtEnd

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 17.6 | 1022.58 MB/s | 0 | 0 |
| record | 11.6 | 4650.85 MB/s | 0 | 0 |
| tiny | 11.6 | 604.15 MB/s | 0 | 0 |
| twoBlock | 14.9 | 5901.85 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 19.8 | 911.30 MB/s | 0 | 0 |
| record | 47.1 | 1146.61 MB/s | 0 | 0 |
| tiny | 17.3 | 405.52 MB/s | 0 | 0 |
| twoBlock | 26.8 | 3288.70 MB/s | 0 | 0 |

## IndexStructural

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 8 | 8.4 | 950.07 MB/s | 0 | 0 |
| 24 | 15.3 | 1568.77 MB/s | 0 | 0 |
| 40 | 8.8 | 4557.53 MB/s | 0 | 0 |
| 72 | 9.1 | 7917.71 MB/s | 0 | 0 |
| 136 | 10.0 | 13562.38 MB/s | 0 | 0 |
| 520 | 15.9 | 32609.46 MB/s | 0 | 0 |
| 4104 | 77.0 | 53266.01 MB/s | 0 | 0 |
| 20000 | 352.3 | 56766.42 MB/s | 0 | 0 |

## ValidNumbers

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| flat | 15483.0 | 4133.54 MB/s | 0 | 0 |
| canada | 18370.0 | 2237.36 MB/s | 0 | 0 |
| geometry | 18758.0 | 2086.35 MB/s | 0 | 0 |
| objects | 28206.0 | 726.51 MB/s | 0 | 0 |
