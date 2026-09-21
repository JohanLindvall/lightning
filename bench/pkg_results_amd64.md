# lightning main-module benchmarks

- generated 2026-09-21T08:05:58Z
- go version go1.26.8 linux/amd64
- cpu: AMD EPYC 9V45 96-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 108.1 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 16.8 | 953.54 MB/s | 16 | 1 |
| sentence_clean | 22.4 | 1967.88 MB/s | 48 | 1 |
| url_clean | 22.7 | 2291.56 MB/s | 64 | 1 |
| log_line_clean | 55.9 | 6014.83 MB/s | 352 | 1 |
| path_with_backslash | 69.3 | 533.74 MB/s | 56 | 2 |
| json_in_json | 89.4 | 469.77 MB/s | 72 | 2 |
| prose_with_quotes | 54.7 | 694.49 MB/s | 64 | 2 |
| control_bytes | 68.8 | 348.91 MB/s | 56 | 2 |
| mostly_clean_one_quote | 61.2 | 4983.82 MB/s | 320 | 1 |
| unicode_clean | 192.0 | 1228.88 MB/s | 240 | 1 |
| unicode_with_quotes | 110.1 | 572.24 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 65.6 | 4648.16 MB/s | 320 | 1 |
| invalid_utf8_dense | 365.7 | 328.17 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.7 | 2787.23 MB/s | 0 | 0 |
| sentence_clean | 11.1 | 3970.41 MB/s | 0 | 0 |
| url_clean | 9.1 | 5736.74 MB/s | 0 | 0 |
| log_line_clean | 14.9 | 22584.71 MB/s | 0 | 0 |
| path_with_backslash | 32.5 | 1139.58 MB/s | 0 | 0 |
| json_in_json | 49.7 | 845.09 MB/s | 0 | 0 |
| prose_with_quotes | 20.9 | 1815.82 MB/s | 0 | 0 |
| control_bytes | 34.5 | 696.54 MB/s | 0 | 0 |
| mostly_clean_one_quote | 17.7 | 17205.41 MB/s | 0 | 0 |
| unicode_clean | 155.1 | 1521.30 MB/s | 0 | 0 |
| unicode_with_quotes | 55.8 | 1128.77 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 28.2 | 10809.15 MB/s | 0 | 0 |
| invalid_utf8_dense | 276.3 | 434.35 MB/s | 0 | 0 |

## ValidEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean/value | 21.3 | 42076.14 MB/s | 0 | 0 |
| clean/key | 24.4 | 37127.11 MB/s | 0 | 0 |
| single/value | 209.1 | 2458.35 MB/s | 0 | 0 |
| single/key | 210.7 | 2472.60 MB/s | 0 | 0 |
| unicode/value | 188.6 | 4083.42 MB/s | 0 | 0 |
| unicode/key | 194.2 | 4001.43 MB/s | 0 | 0 |
| surrogates/value | 193.9 | 3970.95 MB/s | 0 | 0 |
| surrogates/key | 196.8 | 3948.69 MB/s | 0 | 0 |
| mixed/value | 311.5 | 2369.02 MB/s | 0 | 0 |
| mixed/key | 301.3 | 2472.83 MB/s | 0 | 0 |
| sparse/value | 25.5 | 35352.45 MB/s | 0 | 0 |
| sparse/key | 28.0 | 32457.58 MB/s | 0 | 0 |

## DecodeAnyEscapedStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| clean | 129.1 | 6957.49 MB/s | 912 | 2 |
| single | 526.0 | 977.23 MB/s | 357 | 1 |
| unicode | 601.4 | 1280.40 MB/s | 1040 | 2 |
| surrogates | 498.8 | 1543.80 MB/s | 528 | 1 |
| mixed | 587.6 | 1255.90 MB/s | 1040 | 2 |
| sparse | 168.1 | 5371.29 MB/s | 1040 | 2 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 957.2 | 10490.47 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1015.0 | 9896.89 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1005.0 | 9986.12 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 6285.0 | 1597.71 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 345.0 | 5249.53 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 647.3 | 2797.63 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 1.8 | 9119.57 MB/s | 0 | 0 |
| sentence_clean | 3.4 | 12917.45 MB/s | 0 | 0 |
| url_clean | 3.0 | 17477.73 MB/s | 0 | 0 |
| log_line_clean | 5.0 | 67119.41 MB/s | 0 | 0 |
| path_escaped | 45.2 | 950.53 MB/s | 48 | 1 |
| json_in_json | 60.6 | 890.80 MB/s | 64 | 1 |
| prose_with_quotes | 37.3 | 1100.00 MB/s | 48 | 1 |
| unicode_heavy | 2.5 | 12029.51 MB/s | 0 | 0 |
| unicode_escaped_dense | 208.3 | 921.96 MB/s | 192 | 1 |
| mostly_clean_one_escape | 68.7 | 4456.60 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.0 | 7994.49 MB/s | 0 | 0 |
| sentence_clean | 3.2 | 13772.58 MB/s | 0 | 0 |
| url_clean | 3.2 | 16046.87 MB/s | 0 | 0 |
| log_line_clean | 5.6 | 60246.13 MB/s | 0 | 0 |
| path_escaped | 31.3 | 1375.55 MB/s | 0 | 0 |
| json_in_json | 44.8 | 1204.88 MB/s | 0 | 0 |
| prose_with_quotes | 22.5 | 1818.62 MB/s | 0 | 0 |
| unicode_heavy | 2.5 | 11862.02 MB/s | 0 | 0 |
| unicode_escaped_dense | 165.6 | 1159.77 MB/s | 0 | 0 |
| mostly_clean_one_escape | 13.8 | 22200.22 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 2.2 | — | 0 | 0 |
| 3digit | 2.6 | — | 0 | 0 |
| 5digit | 2.9 | — | 0 | 0 |
| 10digit | 4.0 | — | 0 | 0 |
| 13digit | 4.0 | — | 0 | 0 |
| 16digit | 3.9 | — | 0 | 0 |
| 19digit | 5.2 | — | 0 | 0 |
| neg10digit | 3.8 | — | 0 | 0 |
| 20digit_overflow | 4.9 | — | 0 | 0 |
| notanint | 2.2 | — | 0 | 0 |

## ParseUintShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 3digit | 2.5 | — | 0 | 0 |
| 10digit | 3.5 | — | 0 | 0 |
| 13digit | 3.5 | — | 0 | 0 |
| 20digit | 4.7 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 1.5 | — | 0 | 0 |
| short | 2.0 | — | 0 | 0 |
| medium | 2.4 | — | 0 | 0 |
| long | 4.2 | — | 0 | 0 |
| escaped | 20.3 | — | 8 | 1 |
| notastring | 1.3 | — | 0 | 0 |

## BoolShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| true | 0.4 | — | 0 | 0 |
| false | 0.5 | — | 0 | 0 |
| null | 0.4 | — | 0 | 0 |

## KindOfShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| string | 1.5 | — | 0 | 0 |
| number | 1.4 | — | 0 | 0 |
| object | 1.4 | — | 0 | 0 |
| array | 1.4 | — | 0 | 0 |
| null | 1.7 | — | 0 | 0 |
| true | 1.9 | — | 0 | 0 |
| invalid | 1.5 | — | 0 | 0 |
| ws_number | 2.2 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 10.0 | — | 8 | 1 |
| medium | 15.8 | — | 32 | 1 |
| long | 26.3 | — | 128 | 1 |
| escaped | 18.6 | — | 8 | 1 |
| long_late_escape | 33.9 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 856.0 | 2571.31 MB/s | 0 | 0 |
| index | 1202.0 | 1831.73 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1195.0 | 1841.57 MB/s | 0 | 0 |
| strings | 476.6 | 4828.21 MB/s | 0 | 0 |
| records | 455.9 | 6074.34 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 8.1 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 6.1 | — | 0 | 0 |

## DecodeAnyArrays

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| numbers/0 | 4.8 | 418.87 MB/s | 0 | 0 |
| numbers/1 | 58.7 | 153.26 MB/s | 48 | 3 |
| numbers/4 | 115.0 | 286.98 MB/s | 120 | 6 |
| numbers/16 | 365.2 | 353.25 MB/s | 408 | 18 |
| numbers/17 | 410.4 | 333.80 MB/s | 672 | 19 |
| numbers/256 | 5124.0 | 399.86 MB/s | 10904 | 261 |
| numbers/4096 | 88045.0 | 372.19 MB/s | 272281 | 4107 |
| strings/0 | 4.7 | 425.08 MB/s | 0 | 0 |
| strings/1 | 61.4 | 293.03 MB/s | 72 | 4 |
| strings/4 | 158.9 | 434.33 MB/s | 216 | 10 |
| strings/16 | 503.5 | 542.16 MB/s | 792 | 34 |
| strings/17 | 568.3 | 510.33 MB/s | 1080 | 36 |
| strings/256 | 8096.0 | 537.66 MB/s | 17048 | 517 |
| strings/4096 | 139424.0 | 499.43 MB/s | 370584 | 8203 |
| records/0 | 5.3 | 378.77 MB/s | 0 | 0 |
| records/1 | 275.9 | 155.83 MB/s | 488 | 12 |
| records/4 | 962.6 | 175.56 MB/s | 1880 | 42 |
| records/16 | 3884.0 | 173.29 MB/s | 7448 | 162 |
| records/17 | 4070.0 | 175.70 MB/s | 8152 | 172 |
| records/256 | 61128.0 | 175.91 MB/s | 123544 | 2565 |
| records/4096 | 1139395.0 | 150.99 MB/s | 2074533 | 40971 |

## ValidShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar | 17.1 | 408.95 MB/s | 0 | 0 |
| empty_object | 13.2 | 151.40 MB/s | 0 | 0 |
| record | 471.7 | 3839.10 MB/s | 0 | 0 |
| pretty_record | 531.8 | 3870.08 MB/s | 0 | 0 |
| records | 1213.0 | 2282.91 MB/s | 0 | 0 |
| strings | 450.5 | 5107.36 MB/s | 0 | 0 |
| numbers | 2438.0 | 902.67 MB/s | 0 | 0 |
| escapes | 1298.0 | 2619.89 MB/s | 0 | 0 |
| deep | 415.9 | 617.90 MB/s | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2.5 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 33.6 | — | 0 | 0 |
| append_empty | 10.6 | — | 0 | 0 |
| replace | 22.0 | — | 0 | 0 |
| create_nested | 22.6 | — | 0 | 0 |
| overwrite_nonobject | 24.3 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 59.3 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 142.0 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 48.5 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 63.6 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 180873.0 | 6116.22 MB/s | 2293537 | 25 |
| stream | 53940.0 | 20509.06 MB/s | 67136 | 5 |
| stream_reused | 48555.0 | 22783.75 MB/s | 32 | 1 |
| stream_points | 697153.0 | 1586.83 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 7954.0 | 2766.15 MB/s | 0 | 0 |
| scalars/stream | 12643.0 | 1740.18 MB/s | 67120 | 4 |
| scalars/stream_reused | 8675.0 | 2536.16 MB/s | 32 | 1 |
| strings/inmemory | 4289.0 | 5362.51 MB/s | 0 | 0 |
| strings/stream | 9099.0 | 2527.83 MB/s | 67120 | 4 |
| strings/stream_reused | 4804.0 | 4787.52 MB/s | 32 | 1 |
| records/inmemory | 4071.0 | 6927.04 MB/s | 0 | 0 |
| records/stream | 10079.0 | 2797.74 MB/s | 67120 | 4 |
| records/stream_reused | 4589.0 | 6144.99 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 35963.0 | 24439.87 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 420.8 | 3103.91 MB/s | 32 | 1 |
| inmemory | 359.5 | 3632.72 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 850.3 | 4242.07 MB/s | 32 | 1 |
| arrayeach | 865.3 | 4168.51 MB/s | 32 | 1 |
| inmemory | 741.8 | 4862.26 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 83618.0 | 14246.49 MB/s | 42 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1083.0 | 2554.51 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1067.0 | 2591.18 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 13.8 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 107.0 | 1737.92 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 437.7 | 5256.73 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 364.4 | 4970.09 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 381.9 | 4742.53 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 372.3 | 4864.16 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 419.6 | 6599.08 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 788.0 | 2793.03 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1372.0 | 965.99 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 45.9 | — | 24 | 1 |
| arena | 41.1 | — | 24 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 15.8 | — | 64 | 0 |
| make | 14.2 | — | 64 | 1 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 16058.0 | 1222.59 MB/s | 0 | 0 |
| kernel/sep"," | 8443.0 | 2325.17 MB/s | 0 | 0 |
| scalar/sep",_" | 18083.0 | 1306.84 MB/s | 0 | 0 |
| kernel/sep",_" | 10320.0 | 2289.93 MB/s | 0 | 0 |
| kernel-only | 8371.0 | 2345.24 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 5225.0 | 1546.52 MB/s | 0 | 0 |
| "12," | 6677.0 | 1809.12 MB/s | 0 | 0 |
| "123," | 5385.0 | 2985.94 MB/s | 0 | 0 |
| "1234," | 8296.0 | 2420.37 MB/s | 0 | 0 |
| "123456," | 9218.0 | 3046.15 MB/s | 0 | 0 |
| "1234567," | 5529.0 | 5802.58 MB/s | 0 | 0 |
| "1234,_" | 11993.0 | 2007.67 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 2.0 | — | 0 | 0 |
| d03 | 2.5 | — | 0 | 0 |
| d05 | 2.6 | — | 0 | 0 |
| d08 | 2.7 | — | 0 | 0 |
| d10 | 3.6 | — | 0 | 0 |
| d13 | 3.7 | — | 0 | 0 |
| d16 | 3.7 | — | 0 | 0 |
| d19 | 5.0 | — | 0 | 0 |
| d20 | 4.8 | — | 0 | 0 |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/whole | 6914.0 | 1331.88 MB/s | 0 | 0 |
| records/4k | 6994.0 | 1316.77 MB/s | 0 | 0 |
| strings/whole | 2889.0 | 3878.49 MB/s | 0 | 0 |
| strings/4k | 2997.0 | 3739.22 MB/s | 0 | 0 |
| numbers/whole | 136.9 | 61391.71 MB/s | 0 | 0 |
| numbers/4k | 162.3 | 51783.93 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 290.0 | — | 0 | 0 |
| canada | 216.1 | — | 0 | 0 |
| mesh | 191.3 | — | 0 | 0 |
| array | 149.1 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 308.5 | — | 0 | 0 |
| canada | 297.2 | — | 0 | 0 |
| mesh | 200.3 | — | 0 | 0 |
| array | 161.9 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 2256.0 | 10635.52 MB/s | 0 | 0 |
| numberObj/goloop | 727.6 | 14015.42 MB/s | 0 | 0 |
| nestedMixed/goloop | 1092.0 | 9889.81 MB/s | 0 | 0 |
| stringObj/avx2 | 922.0 | 26027.70 MB/s | 0 | 0 |
| numberObj/avx2 | 329.2 | 30978.62 MB/s | 0 | 0 |
| nestedMixed/avx2 | 641.9 | 16826.52 MB/s | 0 | 0 |
| stringObj/avx512 | 774.8 | 30971.38 MB/s | 0 | 0 |
| numberObj/avx512 | 237.4 | 42957.49 MB/s | 0 | 0 |
| nestedMixed/avx512 | 520.8 | 20740.08 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 6519.0 | 3681.11 MB/s | 0 | 0 |
| stringObj/dispatch | 777.5 | 30867.41 MB/s | 0 | 0 |
| numberObj/current | 2483.0 | 4106.44 MB/s | 0 | 0 |
| numberObj/dispatch | 248.5 | 41035.93 MB/s | 0 | 0 |
| numberArr/current | 115.1 | 57365.25 MB/s | 0 | 0 |
| numberArr/dispatch | 117.6 | 56118.46 MB/s | 0 | 0 |
| nestedMixed/current | 10255.0 | 1053.25 MB/s | 0 | 0 |
| nestedMixed/dispatch | 511.1 | 21133.67 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 10.0 | 1808.52 MB/s | 0 | 0 |
| record | 5.3 | 10139.74 MB/s | 0 | 0 |
| tiny | 5.1 | 1370.24 MB/s | 0 | 0 |
| twoBlock | 6.5 | 13434.45 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 9.8 | 1839.84 MB/s | 0 | 0 |
| record | 24.9 | 2164.22 MB/s | 0 | 0 |
| tiny | 9.3 | 756.31 MB/s | 0 | 0 |
| twoBlock | 13.4 | 6575.45 MB/s | 0 | 0 |
