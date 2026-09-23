# lightning main-module benchmarks

- generated 2026-09-23T17:59:16Z
- go version go1.26.8 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 181.4 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.5 | 541.72 MB/s | 16 | 1 |
| sentence_clean | 36.0 | 1223.13 MB/s | 48 | 1 |
| url_clean | 39.4 | 1318.32 MB/s | 64 | 1 |
| log_line_clean | 113.7 | 2953.85 MB/s | 352 | 1 |
| path_with_backslash | 116.1 | 318.77 MB/s | 56 | 2 |
| json_in_json | 153.1 | 274.33 MB/s | 72 | 2 |
| prose_with_quotes | 91.5 | 415.35 MB/s | 64 | 2 |
| control_bytes | 111.8 | 214.68 MB/s | 56 | 2 |
| mostly_clean_one_quote | 133.0 | 2292.63 MB/s | 320 | 1 |
| unicode_clean | 272.8 | 865.08 MB/s | 240 | 1 |
| unicode_with_quotes | 165.8 | 380.06 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 141.8 | 2151.41 MB/s | 320 | 1 |
| invalid_utf8_dense | 621.6 | 193.05 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.5 | 2126.33 MB/s | 0 | 0 |
| sentence_clean | 15.0 | 2924.90 MB/s | 0 | 0 |
| url_clean | 10.9 | 4783.80 MB/s | 0 | 0 |
| log_line_clean | 33.3 | 10096.47 MB/s | 0 | 0 |
| path_with_backslash | 53.4 | 692.91 MB/s | 0 | 0 |
| json_in_json | 84.3 | 498.29 MB/s | 0 | 0 |
| prose_with_quotes | 31.1 | 1222.10 MB/s | 0 | 0 |
| control_bytes | 45.4 | 528.46 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.2 | 8660.61 MB/s | 0 | 0 |
| unicode_clean | 225.0 | 1048.96 MB/s | 0 | 0 |
| unicode_with_quotes | 86.3 | 729.82 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 50.5 | 6034.56 MB/s | 0 | 0 |
| invalid_utf8_dense | 433.3 | 276.96 MB/s | 0 | 0 |

## ValidEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean/value | 47.1 | 19081.81 MB/s | 0 | 0 |
| clean/key | 52.7 | 17165.51 MB/s | 0 | 0 |
| single/value | 368.6 | 1394.51 MB/s | 0 | 0 |
| single/key | 375.7 | 1386.63 MB/s | 0 | 0 |
| unicode/value | 366.6 | 2100.19 MB/s | 0 | 0 |
| unicode/key | 372.3 | 2086.81 MB/s | 0 | 0 |
| surrogates/value | 367.4 | 2095.65 MB/s | 0 | 0 |
| surrogates/key | 372.8 | 2084.39 MB/s | 0 | 0 |
| mixed/value | 569.3 | 1296.30 MB/s | 0 | 0 |
| mixed/key | 575.5 | 1294.54 MB/s | 0 | 0 |
| sparse/value | 53.8 | 16773.59 MB/s | 0 | 0 |
| sparse/key | 59.6 | 15269.59 MB/s | 0 | 0 |

## DecodeAnyEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean | 260.4 | 3448.46 MB/s | 912 | 2 |
| single | 871.7 | 589.66 MB/s | 357 | 1 |
| unicode | 949.8 | 810.67 MB/s | 1040 | 2 |
| surrogates | 887.4 | 867.71 MB/s | 528 | 1 |
| mixed | 1112.0 | 663.72 MB/s | 1040 | 2 |
| sparse | 362.8 | 2488.77 MB/s | 1040 | 2 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2726.0 | 3682.77 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2706.0 | 3710.66 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2690.0 | 3732.52 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10478.0 | 958.31 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 730.0 | 2480.83 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1036.0 | 1747.25 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.6 | 6251.67 MB/s | 0 | 0 |
| sentence_clean | 4.7 | 9296.41 MB/s | 0 | 0 |
| url_clean | 4.7 | 10990.31 MB/s | 0 | 0 |
| log_line_clean | 10.5 | 32025.31 MB/s | 0 | 0 |
| path_escaped | 77.6 | 554.19 MB/s | 48 | 1 |
| json_in_json | 105.7 | 510.90 MB/s | 64 | 1 |
| prose_with_quotes | 67.5 | 607.48 MB/s | 48 | 1 |
| unicode_heavy | 3.3 | 9201.01 MB/s | 0 | 0 |
| unicode_escaped_dense | 288.1 | 666.35 MB/s | 192 | 1 |
| mostly_clean_one_escape | 116.1 | 2635.76 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.1 | 5149.76 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8291.38 MB/s | 0 | 0 |
| url_clean | 5.3 | 9798.59 MB/s | 0 | 0 |
| log_line_clean | 11.1 | 30370.15 MB/s | 0 | 0 |
| path_escaped | 43.9 | 979.37 MB/s | 0 | 0 |
| json_in_json | 68.4 | 789.57 MB/s | 0 | 0 |
| prose_with_quotes | 32.6 | 1255.60 MB/s | 0 | 0 |
| unicode_heavy | 3.8 | 7827.28 MB/s | 0 | 0 |
| unicode_escaped_dense | 220.9 | 869.27 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.3 | 12582.27 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 2.7 | — | 0 | 0 |
| 3digit | 3.5 | — | 0 | 0 |
| 5digit | 4.1 | — | 0 | 0 |
| 10digit | 5.2 | — | 0 | 0 |
| 13digit | 5.2 | — | 0 | 0 |
| 16digit | 5.2 | — | 0 | 0 |
| 19digit | 6.8 | — | 0 | 0 |
| neg10digit | 5.1 | — | 0 | 0 |
| 20digit_overflow | 6.9 | — | 0 | 0 |
| notanint | 2.8 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 3.8 | — | 0 | 0 |
| 10digit | 4.5 | — | 0 | 0 |
| 13digit | 4.5 | — | 0 | 0 |
| 20digit | 6.5 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 2.4 | — | 0 | 0 |
| short | 3.4 | — | 0 | 0 |
| medium | 4.1 | — | 0 | 0 |
| long | 8.6 | — | 0 | 0 |
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
| short | 15.8 | — | 8 | 1 |
| medium | 26.5 | — | 32 | 1 |
| long | 47.6 | — | 128 | 1 |
| escaped | 34.9 | — | 8 | 1 |
| long_late_escape | 94.4 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1662.0 | 1324.18 MB/s | 0 | 0 |
| index | 1783.0 | 1234.22 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1782.0 | 1234.97 MB/s | 0 | 0 |
| strings | 1017.0 | 2263.30 MB/s | 0 | 0 |
| records | 971.9 | 2848.98 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 13.4 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 5.2 | — | 0 | 0 |

## DecodeAnyArrays

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| numbers/0 | 6.5 | 308.21 MB/s | 0 | 0 |
| numbers/1 | 104.6 | 86.07 MB/s | 48 | 3 |
| numbers/4 | 235.2 | 140.33 MB/s | 120 | 6 |
| numbers/16 | 636.3 | 202.74 MB/s | 408 | 18 |
| numbers/17 | 721.3 | 189.94 MB/s | 672 | 19 |
| numbers/256 | 9182.0 | 223.15 MB/s | 10904 | 261 |
| numbers/4096 | 157432.0 | 208.15 MB/s | 272281 | 4107 |
| strings/0 | 6.5 | 308.12 MB/s | 0 | 0 |
| strings/1 | 128.6 | 139.93 MB/s | 72 | 4 |
| strings/4 | 333.2 | 207.06 MB/s | 216 | 10 |
| strings/16 | 1046.0 | 260.94 MB/s | 792 | 34 |
| strings/17 | 1157.0 | 250.61 MB/s | 1080 | 36 |
| strings/256 | 16289.0 | 267.24 MB/s | 17048 | 517 |
| strings/4096 | 274780.0 | 253.41 MB/s | 370584 | 8203 |
| records/0 | 6.5 | 308.30 MB/s | 0 | 0 |
| records/1 | 575.2 | 74.76 MB/s | 488 | 12 |
| records/4 | 2079.0 | 81.30 MB/s | 1880 | 42 |
| records/16 | 8007.0 | 84.05 MB/s | 7448 | 162 |
| records/17 | 8537.0 | 83.75 MB/s | 8152 | 172 |
| records/256 | 127527.0 | 84.32 MB/s | 123544 | 2565 |
| records/4096 | 2290210.0 | 75.12 MB/s | 2074535 | 40971 |

## ValidShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar | 36.0 | 194.23 MB/s | 0 | 0 |
| empty_object | 24.7 | 81.09 MB/s | 0 | 0 |
| record | 909.9 | 1990.27 MB/s | 0 | 0 |
| pretty_record | 1017.0 | 2023.01 MB/s | 0 | 0 |
| records | 2228.0 | 1242.75 MB/s | 0 | 0 |
| strings | 851.3 | 2702.81 MB/s | 0 | 0 |
| numbers | 493.4 | 4461.31 MB/s | 0 | 0 |
| escapes | 2367.0 | 1436.60 MB/s | 0 | 0 |
| deep | 723.1 | 355.44 MB/s | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 4.0 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 63.0 | — | 0 | 0 |
| append_empty | 18.0 | — | 0 | 0 |
| replace | 46.9 | — | 0 | 0 |
| create_nested | 42.0 | — | 0 | 0 |
| overwrite_nonobject | 49.5 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 108.3 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 282.6 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 97.0 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 131.4 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 458872.0 | 2410.83 MB/s | 2293536 | 25 |
| stream | 231330.0 | 4782.18 MB/s | 67136 | 5 |
| stream_reused | 196356.0 | 5633.96 MB/s | 35 | 1 |
| stream_points | 1377719.0 | 802.97 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 17114.0 | 1285.57 MB/s | 0 | 0 |
| scalars/stream | 27280.0 | 806.47 MB/s | 67120 | 4 |
| scalars/stream_reused | 19648.0 | 1119.75 MB/s | 32 | 1 |
| strings/inmemory | 10486.0 | 2193.58 MB/s | 0 | 0 |
| strings/stream | 20013.0 | 1149.32 MB/s | 67120 | 4 |
| strings/stream_reused | 12161.0 | 1891.34 MB/s | 32 | 1 |
| records/inmemory | 9811.0 | 2873.98 MB/s | 0 | 0 |
| records/stream | 18669.0 | 1510.42 MB/s | 67120 | 4 |
| records/stream_reused | 10944.0 | 2576.54 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 196353.0 | 4476.24 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 988.0 | 1321.88 MB/s | 32 | 1 |
| inmemory | 798.5 | 1635.47 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 2086.0 | 1729.41 MB/s | 32 | 1 |
| arrayeach | 2099.0 | 1718.83 MB/s | 32 | 1 |
| inmemory | 1583.0 | 2278.25 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 292015.0 | 4079.46 MB/s | 68 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1969.0 | 1404.60 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1905.0 | 1452.28 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 26.2 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 188.2 | 988.19 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1017.0 | 2261.47 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 807.5 | 2242.85 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 804.6 | 2250.84 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 808.7 | 2239.37 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 968.4 | 2859.38 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1661.0 | 1324.89 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2687.0 | 493.08 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 93.0 | — | 24 | 1 |
| arena | 81.1 | — | 24 | 0 |

## ReadIntShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 1564.0 | — | 0 | 0 |
| d03 | 1575.0 | — | 0 | 0 |
| d06 | 1578.0 | — | 0 | 0 |
| d09 | 1785.0 | — | 0 | 0 |
| d10 | 1802.0 | — | 0 | 0 |
| d13 | 1793.0 | — | 0 | 0 |
| d18 | 2160.0 | — | 0 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 22.9 | — | 64 | 0 |
| make | 42.7 | — | 64 | 1 |

## FloatRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 22453.0 | 2854.73 MB/s | 0 | 0 |
| short_space | 34592.0 | 2084.17 MB/s | 0 | 0 |
| digits6 | 19784.0 | 1824.59 MB/s | 0 | 0 |
| long17 | 36187.0 | 2213.39 MB/s | 0 | 0 |
| long16 | 29195.0 | 2469.47 MB/s | 0 | 0 |

## DecodeFloat64Array

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short3/scalar | 41.3 | — | 0 | 0 |
| short3/kernel | 27.6 | — | 0 | 0 |
| coord2/scalar | 37.7 | — | 0 | 0 |
| coord2/kernel | 27.6 | — | 0 | 0 |
| pretty3/scalar | 52.7 | — | 0 | 0 |
| pretty3/kernel | 34.8 | — | 0 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 22906.0 | 857.06 MB/s | 0 | 0 |
| kernel/sep"," | 8034.0 | 2443.70 MB/s | 0 | 0 |
| scalar/sep",_" | 25515.0 | 926.16 MB/s | 0 | 0 |
| kernel/sep",_" | 11113.0 | 2126.37 MB/s | 0 | 0 |
| kernel-only | 8033.0 | 2444.04 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6181.0 | 1307.27 MB/s | 0 | 0 |
| "12," | 6542.0 | 1846.42 MB/s | 0 | 0 |
| "123," | 6911.0 | 2326.70 MB/s | 0 | 0 |
| "1234," | 7777.0 | 2581.91 MB/s | 0 | 0 |
| "123456," | 8467.0 | 3316.37 MB/s | 0 | 0 |
| "1234567," | 8260.0 | 3883.59 MB/s | 0 | 0 |
| "1234,_" | 10741.0 | 2241.78 MB/s | 0 | 0 |

## DecodeIntSliceShort

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 2/scalar | 45.8 | — | 16 | 1 |
| 2/kernel | 46.0 | — | 16 | 1 |
| 3/scalar | 57.4 | — | 24 | 1 |
| 3/kernel | 58.6 | — | 24 | 1 |
| 4/scalar | 60.0 | — | 32 | 1 |
| 4/kernel | 63.8 | — | 32 | 1 |
| 6/scalar | 73.8 | — | 48 | 1 |
| 6/kernel | 72.8 | — | 48 | 1 |
| 8/scalar | 93.0 | — | 64 | 1 |
| 8/kernel | 78.2 | — | 64 | 1 |
| 12/scalar | 118.3 | — | 96 | 1 |
| 12/kernel | 86.6 | — | 96 | 1 |

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

## Float64Points

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| canada/ring6 | 136.7 | — | — | — |
| canada/ring1000 | 19646.0 | — | — | — |
| geometry/ring6 | 121.3 | — | — | — |
| geometry/ring1000 | 20781.0 | — | — | — |
| citylots/ring6 | 205.4 | — | — | — |
| citylots/ring1000 | 35382.0 | — | — | — |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| strings/whole | 5404.0 | 2073.42 MB/s | 0 | 0 |
| strings/4k | 5370.0 | 2086.41 MB/s | 0 | 0 |
| numbers/whole | 191.9 | 43790.38 MB/s | 0 | 0 |
| numbers/4k | 214.9 | 39096.55 MB/s | 0 | 0 |
| records/whole | 12297.0 | 748.90 MB/s | 0 | 0 |
| records/4k | 12321.0 | 747.44 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 459.5 | — | 0 | 0 |
| canada | 334.3 | — | 0 | 0 |
| mesh | 271.8 | — | 0 | 0 |
| array | 223.3 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 458.8 | — | 0 | 0 |
| canada | 409.2 | — | 0 | 0 |
| mesh | 292.6 | — | 0 | 0 |
| array | 260.9 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4272.0 | 5617.04 MB/s | 0 | 0 |
| numberObj/goloop | 1776.0 | 5743.08 MB/s | 0 | 0 |
| nestedMixed/goloop | 2383.0 | 4531.98 MB/s | 0 | 0 |
| stringObj/neon | 2813.0 | 8532.19 MB/s | 0 | 0 |
| numberObj/neon | 1189.0 | 8580.09 MB/s | 0 | 0 |
| nestedMixed/neon | 1589.0 | 6795.27 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 11932.0 | 2011.31 MB/s | 0 | 0 |
| stringObj/dispatch | 2815.0 | 8526.50 MB/s | 0 | 0 |
| numberObj/current | 4424.0 | 2305.02 MB/s | 0 | 0 |
| numberObj/dispatch | 1190.0 | 8573.25 MB/s | 0 | 0 |
| numberArr/current | 148.8 | 44349.99 MB/s | 0 | 0 |
| numberArr/dispatch | 154.2 | 42795.34 MB/s | 0 | 0 |
| nestedMixed/current | 16052.0 | 672.89 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1592.0 | 6783.47 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 17.8 | 1014.07 MB/s | 0 | 0 |
| record | 13.7 | 3933.09 MB/s | 0 | 0 |
| tiny | 13.2 | 531.65 MB/s | 0 | 0 |
| twoBlock | 20.7 | 4246.18 MB/s | 0 | 0 |

## SkipSmallAtEnd

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 18.0 | 1001.34 MB/s | 0 | 0 |
| record | 14.0 | 3863.40 MB/s | 0 | 0 |
| tiny | 14.0 | 500.46 MB/s | 0 | 0 |
| twoBlock | 21.2 | 4154.94 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 16.2 | 1111.77 MB/s | 0 | 0 |
| record | 43.8 | 1233.00 MB/s | 0 | 0 |
| tiny | 13.3 | 527.20 MB/s | 0 | 0 |
| twoBlock | 25.1 | 3510.17 MB/s | 0 | 0 |

## IndexStructural

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 8 | 3.6 | 2221.46 MB/s | 0 | 0 |
| 24 | 6.4 | 3733.44 MB/s | 0 | 0 |
| 40 | 6.8 | 5924.53 MB/s | 0 | 0 |
| 72 | 7.7 | 9301.45 MB/s | 0 | 0 |
| 136 | 9.6 | 14217.28 MB/s | 0 | 0 |
| 520 | 17.9 | 29096.95 MB/s | 0 | 0 |
| 4104 | 96.8 | 42412.74 MB/s | 0 | 0 |
| 20000 | 449.2 | 44525.48 MB/s | 0 | 0 |

## ValidNumbers

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| flat | 10290.0 | 6219.46 MB/s | 0 | 0 |
| canada | 11926.0 | 3446.43 MB/s | 0 | 0 |
| geometry | 11882.0 | 3293.55 MB/s | 0 | 0 |
| objects | 29931.0 | 684.64 MB/s | 0 | 0 |
