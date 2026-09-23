# lightning main-module benchmarks

- generated 2026-09-23T13:15:49Z
- go version go1.26.8 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 181.5 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.6 | 540.11 MB/s | 16 | 1 |
| sentence_clean | 36.5 | 1206.28 MB/s | 48 | 1 |
| url_clean | 39.8 | 1305.79 MB/s | 64 | 1 |
| log_line_clean | 115.8 | 2901.36 MB/s | 352 | 1 |
| path_with_backslash | 113.8 | 325.15 MB/s | 56 | 2 |
| json_in_json | 151.5 | 277.21 MB/s | 72 | 2 |
| prose_with_quotes | 90.5 | 419.92 MB/s | 64 | 2 |
| control_bytes | 109.4 | 219.35 MB/s | 56 | 2 |
| mostly_clean_one_quote | 135.7 | 2247.89 MB/s | 320 | 1 |
| unicode_clean | 271.7 | 868.58 MB/s | 240 | 1 |
| unicode_with_quotes | 165.1 | 381.64 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 144.0 | 2118.27 MB/s | 320 | 1 |
| invalid_utf8_dense | 613.5 | 195.61 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.5 | 2123.16 MB/s | 0 | 0 |
| sentence_clean | 15.1 | 2922.42 MB/s | 0 | 0 |
| url_clean | 10.9 | 4776.59 MB/s | 0 | 0 |
| log_line_clean | 33.4 | 10067.32 MB/s | 0 | 0 |
| path_with_backslash | 52.4 | 706.75 MB/s | 0 | 0 |
| json_in_json | 84.3 | 498.34 MB/s | 0 | 0 |
| prose_with_quotes | 31.2 | 1217.89 MB/s | 0 | 0 |
| control_bytes | 45.6 | 525.79 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.1 | 8679.64 MB/s | 0 | 0 |
| unicode_clean | 224.7 | 1050.49 MB/s | 0 | 0 |
| unicode_with_quotes | 86.3 | 730.11 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 50.5 | 6033.28 MB/s | 0 | 0 |
| invalid_utf8_dense | 432.6 | 277.38 MB/s | 0 | 0 |

## ValidEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean/value | 47.1 | 19080.72 MB/s | 0 | 0 |
| clean/key | 52.7 | 17165.28 MB/s | 0 | 0 |
| single/value | 370.0 | 1389.37 MB/s | 0 | 0 |
| single/key | 374.4 | 1391.38 MB/s | 0 | 0 |
| unicode/value | 366.5 | 2100.74 MB/s | 0 | 0 |
| unicode/key | 372.4 | 2086.38 MB/s | 0 | 0 |
| surrogates/value | 367.8 | 2093.77 MB/s | 0 | 0 |
| surrogates/key | 372.2 | 2087.46 MB/s | 0 | 0 |
| mixed/value | 566.2 | 1303.32 MB/s | 0 | 0 |
| mixed/key | 575.0 | 1295.68 MB/s | 0 | 0 |
| sparse/value | 53.8 | 16788.44 MB/s | 0 | 0 |
| sparse/key | 59.7 | 15247.34 MB/s | 0 | 0 |

## DecodeAnyEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean | 270.5 | 3320.00 MB/s | 912 | 2 |
| single | 892.1 | 576.20 MB/s | 357 | 1 |
| unicode | 988.2 | 779.22 MB/s | 1040 | 2 |
| surrogates | 919.6 | 837.32 MB/s | 528 | 1 |
| mixed | 1206.0 | 611.88 MB/s | 1040 | 2 |
| sparse | 393.8 | 2292.83 MB/s | 1040 | 2 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2716.0 | 3697.06 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2701.0 | 3717.11 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2694.0 | 3727.80 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10484.0 | 957.71 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 726.9 | 2491.45 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1040.0 | 1741.71 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.6 | 6253.03 MB/s | 0 | 0 |
| sentence_clean | 4.7 | 9296.21 MB/s | 0 | 0 |
| url_clean | 4.7 | 10989.71 MB/s | 0 | 0 |
| log_line_clean | 10.5 | 32043.30 MB/s | 0 | 0 |
| path_escaped | 74.9 | 573.88 MB/s | 48 | 1 |
| json_in_json | 101.3 | 533.23 MB/s | 64 | 1 |
| prose_with_quotes | 61.0 | 672.29 MB/s | 48 | 1 |
| unicode_heavy | 3.3 | 9200.52 MB/s | 0 | 0 |
| unicode_escaped_dense | 292.1 | 657.25 MB/s | 192 | 1 |
| mostly_clean_one_escape | 121.4 | 2520.91 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.1 | 5150.87 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8291.23 MB/s | 0 | 0 |
| url_clean | 5.3 | 9798.79 MB/s | 0 | 0 |
| log_line_clean | 11.1 | 30357.03 MB/s | 0 | 0 |
| path_escaped | 44.0 | 977.71 MB/s | 0 | 0 |
| json_in_json | 68.8 | 785.08 MB/s | 0 | 0 |
| prose_with_quotes | 32.7 | 1253.59 MB/s | 0 | 0 |
| unicode_heavy | 3.8 | 7827.16 MB/s | 0 | 0 |
| unicode_escaped_dense | 220.9 | 869.23 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.3 | 12571.00 MB/s | 0 | 0 |

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
| escaped | 36.6 | — | 8 | 1 |
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
| short | 14.7 | — | 8 | 1 |
| medium | 25.4 | — | 32 | 1 |
| long | 51.7 | — | 128 | 1 |
| escaped | 35.1 | — | 8 | 1 |
| long_late_escape | 91.4 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1661.0 | 1325.34 MB/s | 0 | 0 |
| index | 1782.0 | 1235.03 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1782.0 | 1235.20 MB/s | 0 | 0 |
| strings | 1020.0 | 2256.70 MB/s | 0 | 0 |
| records | 971.6 | 2849.88 MB/s | 0 | 0 |

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
| numbers/0 | 6.5 | 308.28 MB/s | 0 | 0 |
| numbers/1 | 101.7 | 88.52 MB/s | 48 | 3 |
| numbers/4 | 228.5 | 144.43 MB/s | 120 | 6 |
| numbers/16 | 643.6 | 200.45 MB/s | 408 | 18 |
| numbers/17 | 729.1 | 187.90 MB/s | 672 | 19 |
| numbers/256 | 9252.0 | 221.47 MB/s | 10904 | 261 |
| numbers/4096 | 158161.0 | 207.19 MB/s | 272281 | 4107 |
| strings/0 | 6.5 | 308.32 MB/s | 0 | 0 |
| strings/1 | 118.6 | 151.80 MB/s | 72 | 4 |
| strings/4 | 301.9 | 228.56 MB/s | 216 | 10 |
| strings/16 | 987.2 | 276.53 MB/s | 792 | 34 |
| strings/17 | 1141.0 | 254.09 MB/s | 1080 | 36 |
| strings/256 | 16207.0 | 268.59 MB/s | 17048 | 517 |
| strings/4096 | 275754.0 | 252.52 MB/s | 370584 | 8203 |
| records/0 | 6.5 | 308.30 MB/s | 0 | 0 |
| records/1 | 561.2 | 76.63 MB/s | 488 | 12 |
| records/4 | 2067.0 | 81.75 MB/s | 1880 | 42 |
| records/16 | 8008.0 | 84.04 MB/s | 7448 | 162 |
| records/17 | 8452.0 | 84.59 MB/s | 8152 | 172 |
| records/256 | 126043.0 | 85.31 MB/s | 123544 | 2565 |
| records/4096 | 2305999.0 | 74.60 MB/s | 2074534 | 40971 |

## ValidShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar | 36.0 | 194.21 MB/s | 0 | 0 |
| empty_object | 24.7 | 81.01 MB/s | 0 | 0 |
| record | 909.6 | 1991.00 MB/s | 0 | 0 |
| pretty_record | 1011.0 | 2034.84 MB/s | 0 | 0 |
| records | 2228.0 | 1243.01 MB/s | 0 | 0 |
| strings | 851.4 | 2702.69 MB/s | 0 | 0 |
| numbers | 492.9 | 4465.18 MB/s | 0 | 0 |
| escapes | 2352.0 | 1445.81 MB/s | 0 | 0 |
| deep | 724.0 | 354.98 MB/s | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 4.0 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 63.2 | — | 0 | 0 |
| append_empty | 18.0 | — | 0 | 0 |
| replace | 47.0 | — | 0 | 0 |
| create_nested | 42.1 | — | 0 | 0 |
| overwrite_nonobject | 49.5 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 108.4 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 282.6 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 97.3 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 131.4 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 487846.0 | 2267.65 MB/s | 2293540 | 25 |
| stream | 237301.0 | 4661.84 MB/s | 67136 | 5 |
| stream_reused | 199861.0 | 5535.17 MB/s | 35 | 1 |
| stream_points | 1389120.0 | 796.38 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 17114.0 | 1285.56 MB/s | 0 | 0 |
| scalars/stream | 27739.0 | 793.15 MB/s | 67120 | 4 |
| scalars/stream_reused | 19563.0 | 1124.65 MB/s | 32 | 1 |
| strings/inmemory | 10478.0 | 2195.08 MB/s | 0 | 0 |
| strings/stream | 21024.0 | 1094.02 MB/s | 67120 | 4 |
| strings/stream_reused | 12136.0 | 1895.26 MB/s | 32 | 1 |
| records/inmemory | 9802.0 | 2876.69 MB/s | 0 | 0 |
| records/stream | 20341.0 | 1386.29 MB/s | 67120 | 4 |
| records/stream_reused | 10938.0 | 2577.89 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 197820.0 | 4443.04 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 988.1 | 1321.68 MB/s | 32 | 1 |
| inmemory | 799.1 | 1634.25 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 2089.0 | 1727.00 MB/s | 32 | 1 |
| arrayeach | 2099.0 | 1718.16 MB/s | 32 | 1 |
| inmemory | 1586.0 | 2273.96 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 290439.0 | 4101.61 MB/s | 69 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1967.0 | 1406.55 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1899.0 | 1456.93 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 26.2 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 188.5 | 986.64 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1019.0 | 2259.17 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 806.7 | 2245.06 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 800.9 | 2261.20 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 808.9 | 2238.73 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 969.1 | 2857.21 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1661.0 | 1325.30 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2685.0 | 493.48 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 92.8 | — | 24 | 1 |
| arena | 80.2 | — | 24 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 23.1 | — | 64 | 0 |
| make | 43.4 | — | 64 | 1 |

## FloatRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 22448.0 | 2855.40 MB/s | 0 | 0 |
| short_space | 34537.0 | 2087.52 MB/s | 0 | 0 |
| digits6 | 19739.0 | 1828.68 MB/s | 0 | 0 |
| long17 | 36159.0 | 2215.11 MB/s | 0 | 0 |
| long16 | 29207.0 | 2468.46 MB/s | 0 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 22966.0 | 854.82 MB/s | 0 | 0 |
| kernel/sep"," | 7713.0 | 2545.20 MB/s | 0 | 0 |
| scalar/sep",_" | 25644.0 | 921.49 MB/s | 0 | 0 |
| kernel/sep",_" | 11072.0 | 2134.36 MB/s | 0 | 0 |
| kernel-only | 7687.0 | 2553.77 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6212.0 | 1300.64 MB/s | 0 | 0 |
| "12," | 6544.0 | 1846.04 MB/s | 0 | 0 |
| "123," | 6921.0 | 2323.35 MB/s | 0 | 0 |
| "1234," | 7756.0 | 2588.89 MB/s | 0 | 0 |
| "123456," | 8469.0 | 3315.56 MB/s | 0 | 0 |
| "1234567," | 8285.0 | 3872.07 MB/s | 0 | 0 |
| "1234,_" | 10761.0 | 2237.61 MB/s | 0 | 0 |

## DecodeIntSliceShort

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 2/scalar | 45.6 | — | 16 | 1 |
| 2/kernel | 46.3 | — | 16 | 1 |
| 3/scalar | 57.4 | — | 24 | 1 |
| 3/kernel | 58.5 | — | 24 | 1 |
| 4/scalar | 60.0 | — | 32 | 1 |
| 4/kernel | 61.4 | — | 32 | 1 |
| 6/scalar | 74.0 | — | 48 | 1 |
| 6/kernel | 70.3 | — | 48 | 1 |
| 8/scalar | 92.8 | — | 64 | 1 |
| 8/kernel | 78.6 | — | 64 | 1 |
| 12/scalar | 118.9 | — | 96 | 1 |
| 12/kernel | 87.7 | — | 96 | 1 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 2.9 | — | 0 | 0 |
| d03 | 3.6 | — | 0 | 0 |
| d05 | 3.9 | — | 0 | 0 |
| d08 | 3.9 | — | 0 | 0 |
| d10 | 5.1 | — | 0 | 0 |
| d13 | 5.1 | — | 0 | 0 |
| d16 | 5.0 | — | 0 | 0 |
| d19 | 6.7 | — | 0 | 0 |
| d20 | 7.1 | — | 0 | 0 |

## Float64Points

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| canada/ring6 | 137.0 | — | — | — |
| canada/ring1000 | 19601.0 | — | — | — |
| geometry/ring6 | 121.2 | — | — | — |
| geometry/ring1000 | 20705.0 | — | — | — |
| citylots/ring6 | 206.0 | — | — | — |
| citylots/ring1000 | 35473.0 | — | — | — |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/whole | 12281.0 | 749.85 MB/s | 0 | 0 |
| records/4k | 12310.0 | 748.11 MB/s | 0 | 0 |
| strings/whole | 5406.0 | 2072.86 MB/s | 0 | 0 |
| strings/4k | 5422.0 | 2066.47 MB/s | 0 | 0 |
| numbers/whole | 191.9 | 43782.72 MB/s | 0 | 0 |
| numbers/4k | 215.3 | 39025.01 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 461.7 | — | 0 | 0 |
| canada | 343.3 | — | 0 | 0 |
| mesh | 275.2 | — | 0 | 0 |
| array | 223.0 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 457.5 | — | 0 | 0 |
| canada | 409.1 | — | 0 | 0 |
| mesh | 292.4 | — | 0 | 0 |
| array | 260.5 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4279.0 | 5607.94 MB/s | 0 | 0 |
| numberObj/goloop | 1775.0 | 5745.22 MB/s | 0 | 0 |
| nestedMixed/goloop | 2379.0 | 4540.84 MB/s | 0 | 0 |
| stringObj/neon | 2813.0 | 8532.16 MB/s | 0 | 0 |
| numberObj/neon | 1189.0 | 8580.00 MB/s | 0 | 0 |
| nestedMixed/neon | 1590.0 | 6793.27 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 11923.0 | 2012.71 MB/s | 0 | 0 |
| stringObj/dispatch | 2814.0 | 8527.42 MB/s | 0 | 0 |
| numberObj/current | 4424.0 | 2304.91 MB/s | 0 | 0 |
| numberObj/dispatch | 1189.0 | 8574.77 MB/s | 0 | 0 |
| numberArr/current | 148.9 | 44317.57 MB/s | 0 | 0 |
| numberArr/dispatch | 154.3 | 42776.91 MB/s | 0 | 0 |
| nestedMixed/current | 16052.0 | 672.90 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1593.0 | 6779.95 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 18.1 | 996.12 MB/s | 0 | 0 |
| record | 13.1 | 4116.15 MB/s | 0 | 0 |
| tiny | 13.1 | 533.58 MB/s | 0 | 0 |
| twoBlock | 20.6 | 4267.26 MB/s | 0 | 0 |

## SkipSmallAtEnd

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 18.0 | 1000.85 MB/s | 0 | 0 |
| record | 14.0 | 3845.55 MB/s | 0 | 0 |
| tiny | 14.1 | 495.56 MB/s | 0 | 0 |
| twoBlock | 21.2 | 4153.49 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 16.1 | 1117.52 MB/s | 0 | 0 |
| record | 43.9 | 1230.49 MB/s | 0 | 0 |
| tiny | 13.3 | 526.54 MB/s | 0 | 0 |
| twoBlock | 25.2 | 3485.82 MB/s | 0 | 0 |

## IndexStructural

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 8 | 3.4 | 2329.66 MB/s | 0 | 0 |
| 24 | 6.5 | 3690.87 MB/s | 0 | 0 |
| 40 | 6.8 | 5862.58 MB/s | 0 | 0 |
| 72 | 7.8 | 9289.92 MB/s | 0 | 0 |
| 136 | 9.6 | 14150.59 MB/s | 0 | 0 |
| 520 | 17.9 | 29051.74 MB/s | 0 | 0 |
| 4104 | 96.9 | 42363.34 MB/s | 0 | 0 |
| 20000 | 449.1 | 44529.02 MB/s | 0 | 0 |

## ValidNumbers

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| flat | 10655.0 | 6006.87 MB/s | 0 | 0 |
| canada | 11900.0 | 3453.87 MB/s | 0 | 0 |
| geometry | 11965.0 | 3270.79 MB/s | 0 | 0 |
| objects | 29901.0 | 685.32 MB/s | 0 | 0 |
