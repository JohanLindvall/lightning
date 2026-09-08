# lightning main-module benchmarks

- generated 2026-09-08T04:30:35Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 190.6 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.6 | 540.30 MB/s | 16 | 1 |
| sentence_clean | 36.9 | 1193.06 MB/s | 48 | 1 |
| url_clean | 40.3 | 1290.06 MB/s | 64 | 1 |
| log_line_clean | 116.9 | 2873.47 MB/s | 352 | 1 |
| path_with_backslash | 115.9 | 319.17 MB/s | 56 | 2 |
| json_in_json | 153.7 | 273.21 MB/s | 72 | 2 |
| prose_with_quotes | 93.2 | 407.80 MB/s | 64 | 2 |
| control_bytes | 114.6 | 209.39 MB/s | 56 | 2 |
| mostly_clean_one_quote | 135.2 | 2255.56 MB/s | 320 | 1 |
| unicode_clean | 272.1 | 867.39 MB/s | 240 | 1 |
| unicode_with_quotes | 158.8 | 396.71 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 145.7 | 2093.54 MB/s | 320 | 1 |
| invalid_utf8_dense | 635.7 | 188.77 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.9 | 2013.47 MB/s | 0 | 0 |
| sentence_clean | 14.9 | 2944.53 MB/s | 0 | 0 |
| url_clean | 10.6 | 4904.85 MB/s | 0 | 0 |
| log_line_clean | 33.2 | 10127.54 MB/s | 0 | 0 |
| path_with_backslash | 51.6 | 716.53 MB/s | 0 | 0 |
| json_in_json | 84.2 | 499.00 MB/s | 0 | 0 |
| prose_with_quotes | 30.8 | 1235.76 MB/s | 0 | 0 |
| control_bytes | 47.3 | 507.79 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.0 | 8710.09 MB/s | 0 | 0 |
| unicode_clean | 222.1 | 1062.79 MB/s | 0 | 0 |
| unicode_with_quotes | 85.8 | 734.60 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 49.7 | 6140.00 MB/s | 0 | 0 |
| invalid_utf8_dense | 439.9 | 272.77 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2949.0 | 3405.08 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2723.0 | 3687.93 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2705.0 | 3711.95 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10465.0 | 959.46 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 790.7 | 2290.43 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1212.0 | 1493.63 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.8 | 5781.93 MB/s | 0 | 0 |
| sentence_clean | 4.7 | 9265.30 MB/s | 0 | 0 |
| url_clean | 4.8 | 10942.91 MB/s | 0 | 0 |
| log_line_clean | 10.6 | 31565.40 MB/s | 0 | 0 |
| path_escaped | 82.9 | 518.78 MB/s | 48 | 1 |
| json_in_json | 114.7 | 470.85 MB/s | 64 | 1 |
| prose_with_quotes | 70.7 | 580.35 MB/s | 48 | 1 |
| unicode_heavy | 3.5 | 8478.33 MB/s | 0 | 0 |
| unicode_escaped_dense | 294.5 | 651.91 MB/s | 192 | 1 |
| mostly_clean_one_escape | 121.7 | 2514.58 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.4 | 4694.36 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8274.95 MB/s | 0 | 0 |
| url_clean | 5.3 | 9786.29 MB/s | 0 | 0 |
| log_line_clean | 11.2 | 29957.62 MB/s | 0 | 0 |
| path_escaped | 45.0 | 956.43 MB/s | 0 | 0 |
| json_in_json | 68.9 | 783.87 MB/s | 0 | 0 |
| prose_with_quotes | 32.9 | 1245.53 MB/s | 0 | 0 |
| unicode_heavy | 4.0 | 7513.62 MB/s | 0 | 0 |
| unicode_escaped_dense | 220.9 | 869.32 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.5 | 12472.14 MB/s | 0 | 0 |

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
| long | 8.7 | — | 0 | 0 |
| escaped | 36.9 | — | 8 | 1 |
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
| short | 16.3 | — | 8 | 1 |
| medium | 27.8 | — | 32 | 1 |
| long | 50.6 | — | 128 | 1 |
| escaped | 34.8 | — | 8 | 1 |
| long_late_escape | 98.9 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1661.0 | 1325.27 MB/s | 0 | 0 |
| index | 1661.0 | 1325.21 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1661.0 | 1325.34 MB/s | 0 | 0 |
| strings | 1025.0 | 2244.76 MB/s | 0 | 0 |
| records | 1155.0 | 2396.77 MB/s | 0 | 0 |

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
| append | 64.0 | — | 0 | 0 |
| append_empty | 18.2 | — | 0 | 0 |
| replace | 47.7 | — | 0 | 0 |
| create_nested | 42.2 | — | 0 | 0 |
| overwrite_nonobject | 49.7 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 107.4 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 281.6 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 98.8 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 123.2 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 475897.0 | 2324.59 MB/s | 2293536 | 25 |
| stream | 240088.0 | 4607.73 MB/s | 66992 | 4 |
| stream_reused | 200034.0 | 5530.38 MB/s | 35 | 1 |
| stream_points | 1416218.0 | 781.14 MB/s | 66992 | 4 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/inmemory | 11151.0 | 2528.73 MB/s | 0 | 0 |
| records/stream | 20158.0 | 1398.88 MB/s | 66992 | 4 |
| scalars/inmemory | 16848.0 | 1305.86 MB/s | 0 | 0 |
| scalars/stream | 26770.0 | 821.84 MB/s | 66992 | 4 |
| strings/inmemory | 10386.0 | 2214.54 MB/s | 0 | 0 |
| strings/stream | 20013.0 | 1149.28 MB/s | 66992 | 4 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 169018.0 | 5200.17 MB/s | 66992 | 4 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 975.7 | 1338.56 MB/s | 32 | 1 |
| inmemory | 796.1 | 1640.56 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 2192.0 | 1645.76 MB/s | 32 | 1 |
| arrayeach | 2199.0 | 1640.64 MB/s | 32 | 1 |
| inmemory | 1621.0 | 2224.83 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 294203.0 | 4049.13 MB/s | 69 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1982.0 | 1395.26 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1912.0 | 1446.49 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 27.9 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 222.2 | 837.15 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1034.0 | 2225.66 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 850.0 | 2130.47 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 875.3 | 2068.89 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 854.2 | 2120.21 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1160.0 | 2386.51 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1661.0 | 1325.42 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2972.0 | 445.83 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 87.8 | — | 24 | 1 |
| arena | 79.2 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 23391.0 | 839.31 MB/s | 0 | 0 |
| kernel/sep"," | 8532.0 | 2301.06 MB/s | 0 | 0 |
| scalar/sep",_" | 25552.0 | 924.81 MB/s | 0 | 0 |
| kernel/sep",_" | 11003.0 | 2147.74 MB/s | 0 | 0 |
| kernel-only | 8545.0 | 2297.40 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6873.0 | 1175.67 MB/s | 0 | 0 |
| "12," | 7186.0 | 1681.02 MB/s | 0 | 0 |
| "123," | 7541.0 | 2132.37 MB/s | 0 | 0 |
| "1234," | 8177.0 | 2455.61 MB/s | 0 | 0 |
| "123456," | 8697.0 | 3228.82 MB/s | 0 | 0 |
| "1234567," | 8636.0 | 3714.75 MB/s | 0 | 0 |
| "1234,_" | 10838.0 | 2221.73 MB/s | 0 | 0 |

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
| records/whole | 11414.0 | 806.82 MB/s | 0 | 0 |
| records/4k | 11399.0 | 807.86 MB/s | 0 | 0 |
| strings/whole | 4192.0 | 2673.16 MB/s | 0 | 0 |
| strings/4k | 4217.0 | 2657.24 MB/s | 0 | 0 |
| numbers/whole | 199.3 | 42161.96 MB/s | 0 | 0 |
| numbers/4k | 240.3 | 34966.63 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 462.1 | — | 0 | 0 |
| canada | 334.8 | — | 0 | 0 |
| mesh | 275.2 | — | 0 | 0 |
| array | 223.1 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 459.3 | — | 0 | 0 |
| canada | 407.8 | — | 0 | 0 |
| mesh | 291.1 | — | 0 | 0 |
| array | 260.9 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4408.0 | 5444.11 MB/s | 0 | 0 |
| numberObj/goloop | 1826.0 | 5584.49 MB/s | 0 | 0 |
| nestedMixed/goloop | 2413.0 | 4475.90 MB/s | 0 | 0 |
| stringObj/neon | 2899.0 | 8278.49 MB/s | 0 | 0 |
| numberObj/neon | 1220.0 | 8356.84 MB/s | 0 | 0 |
| nestedMixed/neon | 1648.0 | 6552.35 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10450.0 | 2296.50 MB/s | 0 | 0 |
| stringObj/dispatch | 2900.0 | 8275.42 MB/s | 0 | 0 |
| numberObj/current | 5635.0 | 1809.69 MB/s | 0 | 0 |
| numberObj/dispatch | 1221.0 | 8349.75 MB/s | 0 | 0 |
| numberArr/current | 158.8 | 41569.51 MB/s | 0 | 0 |
| numberArr/dispatch | 160.5 | 41132.27 MB/s | 0 | 0 |
| nestedMixed/current | 14294.0 | 755.65 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1649.0 | 6549.83 MB/s | 0 | 0 |
