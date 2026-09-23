# lightning main-module benchmarks

- generated 2026-09-23T17:59:14Z
- go version go1.26.8 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 133.3 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 23.8 | 672.38 MB/s | 16 | 1 |
| sentence_clean | 31.5 | 1395.06 MB/s | 48 | 1 |
| url_clean | 30.7 | 1693.27 MB/s | 64 | 1 |
| log_line_clean | 73.0 | 4602.35 MB/s | 352 | 1 |
| path_with_backslash | 94.2 | 392.75 MB/s | 56 | 2 |
| json_in_json | 124.5 | 337.41 MB/s | 72 | 2 |
| prose_with_quotes | 76.1 | 499.13 MB/s | 64 | 2 |
| control_bytes | 90.1 | 266.27 MB/s | 56 | 2 |
| mostly_clean_one_quote | 79.3 | 3845.05 MB/s | 320 | 1 |
| unicode_clean | 240.6 | 980.68 MB/s | 240 | 1 |
| unicode_with_quotes | 126.3 | 498.84 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 88.2 | 3457.39 MB/s | 320 | 1 |
| invalid_utf8_dense | 541.1 | 221.76 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.5 | 1889.41 MB/s | 0 | 0 |
| sentence_clean | 16.4 | 2682.52 MB/s | 0 | 0 |
| url_clean | 15.5 | 3364.44 MB/s | 0 | 0 |
| log_line_clean | 22.1 | 15179.53 MB/s | 0 | 0 |
| path_with_backslash | 52.5 | 704.72 MB/s | 0 | 0 |
| json_in_json | 79.4 | 528.73 MB/s | 0 | 0 |
| prose_with_quotes | 30.9 | 1228.94 MB/s | 0 | 0 |
| control_bytes | 48.0 | 500.48 MB/s | 0 | 0 |
| mostly_clean_one_quote | 26.4 | 11564.17 MB/s | 0 | 0 |
| unicode_clean | 203.4 | 1160.03 MB/s | 0 | 0 |
| unicode_with_quotes | 80.8 | 779.19 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 41.2 | 7395.15 MB/s | 0 | 0 |
| invalid_utf8_dense | 442.5 | 271.19 MB/s | 0 | 0 |

## ValidEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean/value | 33.0 | 27203.91 MB/s | 0 | 0 |
| clean/key | 37.7 | 23991.08 MB/s | 0 | 0 |
| single/value | 441.8 | 1163.50 MB/s | 0 | 0 |
| single/key | 419.1 | 1243.27 MB/s | 0 | 0 |
| unicode/value | 271.5 | 2836.13 MB/s | 0 | 0 |
| unicode/key | 274.5 | 2830.96 MB/s | 0 | 0 |
| surrogates/value | 274.4 | 2805.68 MB/s | 0 | 0 |
| surrogates/key | 279.2 | 2782.60 MB/s | 0 | 0 |
| mixed/value | 390.4 | 1890.39 MB/s | 0 | 0 |
| mixed/key | 389.3 | 1913.60 MB/s | 0 | 0 |
| sparse/value | 39.5 | 22877.04 MB/s | 0 | 0 |
| sparse/key | 44.6 | 20421.87 MB/s | 0 | 0 |

## DecodeAnyEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean | 166.8 | 5382.14 MB/s | 912 | 2 |
| single | 670.6 | 766.48 MB/s | 357 | 1 |
| unicode | 726.0 | 1060.57 MB/s | 1040 | 2 |
| surrogates | 679.3 | 1133.48 MB/s | 528 | 1 |
| mixed | 734.1 | 1005.34 MB/s | 1040 | 2 |
| sparse | 215.9 | 4182.47 MB/s | 1040 | 2 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1431.0 | 7015.52 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1415.0 | 7094.00 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1380.0 | 7273.72 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 8562.0 | 1172.77 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 468.6 | 3864.89 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 919.6 | 1969.29 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.5 | 6498.62 MB/s | 0 | 0 |
| sentence_clean | 4.4 | 10060.97 MB/s | 0 | 0 |
| url_clean | 4.4 | 11818.09 MB/s | 0 | 0 |
| log_line_clean | 7.4 | 45559.53 MB/s | 0 | 0 |
| path_escaped | 65.6 | 655.72 MB/s | 48 | 1 |
| json_in_json | 86.9 | 621.29 MB/s | 64 | 1 |
| prose_with_quotes | 54.4 | 754.25 MB/s | 48 | 1 |
| unicode_heavy | 3.3 | 9071.84 MB/s | 0 | 0 |
| unicode_escaped_dense | 244.7 | 784.66 MB/s | 192 | 1 |
| mostly_clean_one_escape | 69.1 | 4428.43 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.0 | 5322.77 MB/s | 0 | 0 |
| sentence_clean | 4.6 | 9479.06 MB/s | 0 | 0 |
| url_clean | 5.0 | 10330.21 MB/s | 0 | 0 |
| log_line_clean | 7.7 | 43908.60 MB/s | 0 | 0 |
| path_escaped | 48.6 | 884.01 MB/s | 0 | 0 |
| json_in_json | 67.8 | 796.85 MB/s | 0 | 0 |
| prose_with_quotes | 38.0 | 1080.04 MB/s | 0 | 0 |
| unicode_heavy | 3.6 | 8438.27 MB/s | 0 | 0 |
| unicode_escaped_dense | 211.1 | 909.42 MB/s | 0 | 0 |
| mostly_clean_one_escape | 17.9 | 17126.84 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 3.6 | — | 0 | 0 |
| 3digit | 4.7 | — | 0 | 0 |
| 5digit | 4.0 | — | 0 | 0 |
| 10digit | 5.4 | — | 0 | 0 |
| 13digit | 5.4 | — | 0 | 0 |
| 16digit | 5.4 | — | 0 | 0 |
| 19digit | 7.0 | — | 0 | 0 |
| neg10digit | 5.2 | — | 0 | 0 |
| 20digit_overflow | 6.6 | — | 0 | 0 |
| notanint | 3.3 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 4.1 | — | 0 | 0 |
| 10digit | 4.8 | — | 0 | 0 |
| 13digit | 4.8 | — | 0 | 0 |
| 20digit | 6.7 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 2.2 | — | 0 | 0 |
| short | 3.0 | — | 0 | 0 |
| medium | 3.7 | — | 0 | 0 |
| long | 6.6 | — | 0 | 0 |
| escaped | 27.4 | — | 8 | 1 |
| notastring | 2.2 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.6 | — | 0 | 0 |
| false | 0.8 | — | 0 | 0 |
| null | 1.1 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 1.9 | — | 0 | 0 |
| number | 1.9 | — | 0 | 0 |
| object | 1.9 | — | 0 | 0 |
| array | 1.6 | — | 0 | 0 |
| null | 2.2 | — | 0 | 0 |
| true | 2.5 | — | 0 | 0 |
| invalid | 1.9 | — | 0 | 0 |
| ws_number | 2.7 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 13.1 | — | 8 | 1 |
| medium | 18.9 | — | 32 | 1 |
| long | 31.8 | — | 128 | 1 |
| escaped | 26.2 | — | 8 | 1 |
| long_late_escape | 42.0 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1415.0 | 1555.83 MB/s | 0 | 0 |
| index | 1487.0 | 1480.65 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1486.0 | 1480.91 MB/s | 0 | 0 |
| strings | 695.6 | 3308.00 MB/s | 0 | 0 |
| records | 569.5 | 4862.38 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12.0 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 7.3 | — | 0 | 0 |

## DecodeAnyArrays

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| numbers/0 | 7.1 | 280.87 MB/s | 0 | 0 |
| numbers/1 | 70.2 | 128.15 MB/s | 48 | 3 |
| numbers/4 | 152.9 | 215.83 MB/s | 120 | 6 |
| numbers/16 | 454.3 | 283.98 MB/s | 408 | 18 |
| numbers/17 | 521.1 | 262.91 MB/s | 672 | 19 |
| numbers/256 | 7034.0 | 291.29 MB/s | 10904 | 261 |
| numbers/4096 | 122550.0 | 267.39 MB/s | 272281 | 4107 |
| strings/0 | 7.1 | 281.54 MB/s | 0 | 0 |
| strings/1 | 85.0 | 211.82 MB/s | 72 | 4 |
| strings/4 | 205.7 | 335.46 MB/s | 216 | 10 |
| strings/16 | 698.8 | 390.67 MB/s | 792 | 34 |
| strings/17 | 782.2 | 370.75 MB/s | 1080 | 36 |
| strings/256 | 11090.0 | 392.52 MB/s | 17048 | 517 |
| strings/4096 | 193001.0 | 360.79 MB/s | 370584 | 8203 |
| records/0 | 7.4 | 271.03 MB/s | 0 | 0 |
| records/1 | 340.7 | 126.23 MB/s | 488 | 12 |
| records/4 | 1202.0 | 140.57 MB/s | 1880 | 42 |
| records/16 | 4684.0 | 143.68 MB/s | 7448 | 162 |
| records/17 | 5044.0 | 141.76 MB/s | 8152 | 172 |
| records/256 | 75189.0 | 143.01 MB/s | 123544 | 2565 |
| records/4096 | 1434451.0 | 119.93 MB/s | 2074533 | 40971 |

## ValidShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar | 28.2 | 248.62 MB/s | 0 | 0 |
| empty_object | 21.6 | 92.74 MB/s | 0 | 0 |
| record | 611.0 | 2963.99 MB/s | 0 | 0 |
| pretty_record | 695.9 | 2957.45 MB/s | 0 | 0 |
| records | 1599.0 | 1731.56 MB/s | 0 | 0 |
| strings | 541.6 | 4248.19 MB/s | 0 | 0 |
| numbers | 434.0 | 5070.91 MB/s | 0 | 0 |
| escapes | 1766.0 | 1925.72 MB/s | 0 | 0 |
| deep | 682.4 | 376.61 MB/s | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3.7 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 52.0 | — | 0 | 0 |
| append_empty | 17.8 | — | 0 | 0 |
| replace | 33.8 | — | 0 | 0 |
| create_nested | 38.2 | — | 0 | 0 |
| overwrite_nonobject | 39.9 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 95.2 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 229.7 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 75.0 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 100.5 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 262455.0 | 4215.06 MB/s | 2293538 | 25 |
| stream | 85381.0 | 12956.77 MB/s | 67136 | 5 |
| stream_reused | 77653.0 | 14246.22 MB/s | 33 | 1 |
| stream_points | 867526.0 | 1275.19 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 13500.0 | 1629.66 MB/s | 0 | 0 |
| scalars/stream | 25181.0 | 873.73 MB/s | 67120 | 4 |
| scalars/stream_reused | 19568.0 | 1124.34 MB/s | 32 | 1 |
| strings/inmemory | 6129.0 | 3752.58 MB/s | 0 | 0 |
| strings/stream | 13006.0 | 1768.51 MB/s | 67120 | 4 |
| strings/stream_reused | 7472.0 | 3078.10 MB/s | 32 | 1 |
| records/inmemory | 5454.0 | 5169.90 MB/s | 0 | 0 |
| records/stream | 12011.0 | 2347.74 MB/s | 67120 | 4 |
| records/stream_reused | 5817.0 | 4847.71 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 58603.0 | 14997.91 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 604.4 | 2160.90 MB/s | 32 | 1 |
| inmemory | 473.7 | 2757.16 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 1118.0 | 3225.15 MB/s | 32 | 1 |
| arrayeach | 1128.0 | 3198.38 MB/s | 32 | 1 |
| inmemory | 965.0 | 3737.78 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 125929.0 | 9459.82 MB/s | 47 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1689.0 | 1638.02 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1610.0 | 1718.47 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 19.1 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 131.6 | 1413.44 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 638.4 | 3604.43 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 529.3 | 3421.28 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 509.0 | 3557.65 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 509.0 | 3557.92 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 530.3 | 5221.24 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1416.0 | 1554.43 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2118.0 | 625.51 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 81.4 | — | 24 | 1 |
| arena | 74.4 | — | 24 | 0 |

## ReadIntShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 773.2 | — | 0 | 0 |
| d03 | 993.3 | — | 0 | 0 |
| d06 | 1320.0 | — | 0 | 0 |
| d09 | 992.2 | — | 0 | 0 |
| d10 | 1119.0 | — | 0 | 0 |
| d13 | 1559.0 | — | 0 | 0 |
| d18 | 2042.0 | — | 0 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 20.1 | — | 64 | 0 |
| make | 18.9 | — | 64 | 1 |

## FloatRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 15758.0 | 4067.57 MB/s | 0 | 0 |
| short_space | 17232.0 | 4183.96 MB/s | 0 | 0 |
| digits6 | 13297.0 | 2714.73 MB/s | 0 | 0 |
| long17 | 30840.0 | 2597.20 MB/s | 0 | 0 |
| long16 | 22012.0 | 3275.41 MB/s | 0 | 0 |

## DecodeFloat64Array

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short3/scalar | 39.7 | — | 0 | 0 |
| short3/kernel | 19.9 | — | 0 | 0 |
| coord2/scalar | 38.6 | — | 0 | 0 |
| coord2/kernel | 20.2 | — | 0 | 0 |
| pretty3/scalar | 48.4 | — | 0 | 0 |
| pretty3/kernel | 19.6 | — | 0 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 25532.0 | 768.91 MB/s | 0 | 0 |
| kernel/sep"," | 6900.0 | 2845.24 MB/s | 0 | 0 |
| scalar/sep",_" | 29067.0 | 812.99 MB/s | 0 | 0 |
| kernel/sep",_" | 7004.0 | 3374.03 MB/s | 0 | 0 |
| kernel-only | 6914.0 | 2839.26 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 5499.0 | 1469.45 MB/s | 0 | 0 |
| "12," | 5817.0 | 2076.68 MB/s | 0 | 0 |
| "123," | 6213.0 | 2588.20 MB/s | 0 | 0 |
| "1234," | 6629.0 | 3029.02 MB/s | 0 | 0 |
| "123456," | 7289.0 | 3852.48 MB/s | 0 | 0 |
| "1234567," | 8125.0 | 3948.26 MB/s | 0 | 0 |
| "1234,_" | 6860.0 | 3509.81 MB/s | 0 | 0 |

## DecodeIntSliceShort

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 2/scalar | 31.9 | — | 16 | 1 |
| 2/kernel | 30.3 | — | 16 | 1 |
| 3/scalar | 39.3 | — | 24 | 1 |
| 3/kernel | 32.9 | — | 24 | 1 |
| 4/scalar | 46.1 | — | 32 | 1 |
| 4/kernel | 35.8 | — | 32 | 1 |
| 6/scalar | 60.7 | — | 48 | 1 |
| 6/kernel | 40.4 | — | 48 | 1 |
| 8/scalar | 74.9 | — | 64 | 1 |
| 8/kernel | 44.3 | — | 64 | 1 |
| 12/scalar | 103.9 | — | 96 | 1 |
| 12/kernel | 53.9 | — | 96 | 1 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 3.6 | — | 0 | 0 |
| d03 | 4.6 | — | 0 | 0 |
| d05 | 3.8 | — | 0 | 0 |
| d08 | 3.8 | — | 0 | 0 |
| d10 | 5.3 | — | 0 | 0 |
| d13 | 5.3 | — | 0 | 0 |
| d16 | 5.3 | — | 0 | 0 |
| d19 | 7.2 | — | 0 | 0 |
| d20 | 7.0 | — | 0 | 0 |

## Float64Points

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| canada/ring6 | 119.2 | — | — | — |
| canada/ring1000 | 19152.0 | — | — | — |
| geometry/ring6 | 100.0 | — | — | — |
| geometry/ring1000 | 14581.0 | — | — | — |
| citylots/ring6 | 119.4 | — | — | — |
| citylots/ring1000 | 17814.0 | — | — | — |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| strings/whole | 3893.0 | 2878.37 MB/s | 0 | 0 |
| strings/4k | 3893.0 | 2877.99 MB/s | 0 | 0 |
| numbers/whole | 80.6 | 104271.87 MB/s | 0 | 0 |
| numbers/4k | 96.4 | 87144.50 MB/s | 0 | 0 |
| records/whole | 8846.0 | 1041.02 MB/s | 0 | 0 |
| records/4k | 8843.0 | 1041.39 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 411.9 | — | 0 | 0 |
| canada | 313.2 | — | 0 | 0 |
| mesh | 276.1 | — | 0 | 0 |
| array | 217.1 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 414.0 | — | 0 | 0 |
| canada | 389.2 | — | 0 | 0 |
| mesh | 278.8 | — | 0 | 0 |
| array | 236.5 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3109.0 | 7718.64 MB/s | 0 | 0 |
| numberObj/goloop | 1091.0 | 9343.27 MB/s | 0 | 0 |
| nestedMixed/goloop | 1434.0 | 7534.36 MB/s | 0 | 0 |
| stringObj/avx2 | 1696.0 | 14152.14 MB/s | 0 | 0 |
| numberObj/avx2 | 608.7 | 16752.63 MB/s | 0 | 0 |
| nestedMixed/avx2 | 893.6 | 12087.73 MB/s | 0 | 0 |
| stringObj/avx512 | 1095.0 | 21907.28 MB/s | 0 | 0 |
| numberObj/avx512 | 408.2 | 24983.60 MB/s | 0 | 0 |
| nestedMixed/avx512 | 787.5 | 13716.20 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 8436.0 | 2844.63 MB/s | 0 | 0 |
| stringObj/dispatch | 1101.0 | 21793.81 MB/s | 0 | 0 |
| numberObj/current | 3150.0 | 3237.37 MB/s | 0 | 0 |
| numberObj/dispatch | 408.0 | 24994.32 MB/s | 0 | 0 |
| numberArr/current | 62.6 | 105454.01 MB/s | 0 | 0 |
| numberArr/dispatch | 68.0 | 97025.41 MB/s | 0 | 0 |
| nestedMixed/current | 12492.0 | 864.61 MB/s | 0 | 0 |
| nestedMixed/dispatch | 788.6 | 13696.82 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 12.3 | 1461.70 MB/s | 0 | 0 |
| record | 7.4 | 7271.05 MB/s | 0 | 0 |
| tiny | 7.4 | 948.30 MB/s | 0 | 0 |
| twoBlock | 9.6 | 9195.79 MB/s | 0 | 0 |

## SkipSmallAtEnd

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 12.8 | 1404.33 MB/s | 0 | 0 |
| record | 7.7 | 7057.95 MB/s | 0 | 0 |
| tiny | 7.6 | 915.47 MB/s | 0 | 0 |
| twoBlock | 10.4 | 8475.98 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 13.7 | 1311.86 MB/s | 0 | 0 |
| record | 34.2 | 1577.92 MB/s | 0 | 0 |
| tiny | 10.8 | 647.14 MB/s | 0 | 0 |
| twoBlock | 19.6 | 4496.24 MB/s | 0 | 0 |

## IndexStructural

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 8 | 4.0 | 1979.87 MB/s | 0 | 0 |
| 24 | 6.6 | 3659.90 MB/s | 0 | 0 |
| 40 | 6.6 | 6063.71 MB/s | 0 | 0 |
| 72 | 6.6 | 10983.91 MB/s | 0 | 0 |
| 136 | 7.2 | 18771.12 MB/s | 0 | 0 |
| 520 | 9.0 | 57582.58 MB/s | 0 | 0 |
| 4104 | 38.3 | 107081.78 MB/s | 0 | 0 |
| 20000 | 172.9 | 115701.01 MB/s | 0 | 0 |

## ValidNumbers

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| flat | 10435.0 | 6133.56 MB/s | 0 | 0 |
| canada | 9529.0 | 4313.07 MB/s | 0 | 0 |
| geometry | 9523.0 | 4109.52 MB/s | 0 | 0 |
| objects | 21182.0 | 967.42 MB/s | 0 | 0 |
