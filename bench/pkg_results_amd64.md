# lightning main-module benchmarks

- generated 2026-09-08T12:19:06Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## ArrayEachIndex

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 147.9 | — | 0 | 0 |

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 24.6 | 650.18 MB/s | 16 | 1 |
| sentence_clean | 31.6 | 1390.46 MB/s | 48 | 1 |
| url_clean | 31.1 | 1669.22 MB/s | 64 | 1 |
| log_line_clean | 73.6 | 4562.63 MB/s | 352 | 1 |
| path_with_backslash | 91.1 | 406.21 MB/s | 56 | 2 |
| json_in_json | 121.2 | 346.60 MB/s | 72 | 2 |
| prose_with_quotes | 76.4 | 497.11 MB/s | 64 | 2 |
| control_bytes | 93.2 | 257.49 MB/s | 56 | 2 |
| mostly_clean_one_quote | 81.7 | 3735.28 MB/s | 320 | 1 |
| unicode_clean | 239.2 | 986.81 MB/s | 240 | 1 |
| unicode_with_quotes | 128.7 | 489.63 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 89.3 | 3414.16 MB/s | 320 | 1 |
| invalid_utf8_dense | 523.8 | 229.12 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.7 | 1829.15 MB/s | 0 | 0 |
| sentence_clean | 16.4 | 2683.34 MB/s | 0 | 0 |
| url_clean | 15.6 | 3342.22 MB/s | 0 | 0 |
| log_line_clean | 22.6 | 14877.04 MB/s | 0 | 0 |
| path_with_backslash | 49.3 | 751.08 MB/s | 0 | 0 |
| json_in_json | 75.7 | 554.93 MB/s | 0 | 0 |
| prose_with_quotes | 31.6 | 1201.60 MB/s | 0 | 0 |
| control_bytes | 49.2 | 487.53 MB/s | 0 | 0 |
| mostly_clean_one_quote | 25.7 | 11887.45 MB/s | 0 | 0 |
| unicode_clean | 198.7 | 1187.65 MB/s | 0 | 0 |
| unicode_with_quotes | 80.4 | 783.41 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 40.0 | 7629.54 MB/s | 0 | 0 |
| invalid_utf8_dense | 421.7 | 284.59 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1416.0 | 7091.69 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1462.0 | 6867.97 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1382.0 | 7264.72 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 8348.0 | 1202.74 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 469.8 | 3855.00 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 932.0 | 1943.08 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.7 | 5852.99 MB/s | 0 | 0 |
| sentence_clean | 4.6 | 9478.02 MB/s | 0 | 0 |
| url_clean | 4.6 | 11191.28 MB/s | 0 | 0 |
| log_line_clean | 7.7 | 43923.25 MB/s | 0 | 0 |
| path_escaped | 57.3 | 750.51 MB/s | 48 | 1 |
| json_in_json | 76.6 | 705.30 MB/s | 64 | 1 |
| prose_with_quotes | 46.6 | 879.27 MB/s | 48 | 1 |
| unicode_heavy | 3.3 | 9110.00 MB/s | 0 | 0 |
| unicode_escaped_dense | 223.4 | 859.27 MB/s | 192 | 1 |
| mostly_clean_one_escape | 67.7 | 4522.24 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.1 | 5092.77 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 8949.15 MB/s | 0 | 0 |
| url_clean | 4.9 | 10573.76 MB/s | 0 | 0 |
| log_line_clean | 7.9 | 42385.49 MB/s | 0 | 0 |
| path_escaped | 41.1 | 1046.95 MB/s | 0 | 0 |
| json_in_json | 60.2 | 897.00 MB/s | 0 | 0 |
| prose_with_quotes | 31.6 | 1296.66 MB/s | 0 | 0 |
| unicode_heavy | 3.7 | 8122.43 MB/s | 0 | 0 |
| unicode_escaped_dense | 192.5 | 997.45 MB/s | 0 | 0 |
| mostly_clean_one_escape | 17.9 | 17081.57 MB/s | 0 | 0 |

## ParseIntShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| 1digit | 3.6 | — | 0 | 0 |
| 3digit | 4.6 | — | 0 | 0 |
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
| 20digit | 6.8 | — | 0 | 0 |

## StringShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| empty | 2.2 | — | 0 | 0 |
| short | 3.3 | — | 0 | 0 |
| medium | 3.8 | — | 0 | 0 |
| long | 6.8 | — | 0 | 0 |
| escaped | 26.2 | — | 8 | 1 |
| notastring | 1.7 | — | 0 | 0 |

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
| array | 1.9 | — | 0 | 0 |
| null | 2.2 | — | 0 | 0 |
| true | 2.7 | — | 0 | 0 |
| invalid | 1.9 | — | 0 | 0 |
| ws_number | 3.0 | — | 0 | 0 |

## UnescapeCopyShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short | 12.9 | — | 8 | 1 |
| medium | 18.6 | — | 32 | 1 |
| long | 31.4 | — | 128 | 1 |
| escaped | 25.0 | — | 8 | 1 |
| long_late_escape | 41.4 | — | 144 | 1 |

## ArrayEachIndexVsEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| each | 1419.0 | 1551.13 MB/s | 0 | 0 |
| index | 1495.0 | 1472.58 MB/s | 0 | 0 |

## ArrayEachIndexShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars | 1493.0 | 1474.66 MB/s | 0 | 0 |
| strings | 641.9 | 3584.76 MB/s | 0 | 0 |
| records | 540.7 | 5121.15 MB/s | 0 | 0 |

## ErrStop

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12.0 | — | 0 | 0 |

## ParseInt

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 7.2 | — | 0 | 0 |

## String

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3.8 | — | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 52.2 | — | 0 | 0 |
| append_empty | 16.8 | — | 0 | 0 |
| replace | 34.4 | — | 0 | 0 |
| create_nested | 36.6 | — | 0 | 0 |
| overwrite_nonobject | 39.0 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 93.9 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 223.6 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 75.4 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 102.8 | — | 0 | 0 |

## StreamMatrix

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| readall+walk | 271686.0 | 4071.85 MB/s | 2293539 | 25 |
| stream | 85811.0 | 12891.85 MB/s | 67136 | 5 |
| stream_reused | 78688.0 | 14058.92 MB/s | 33 | 1 |
| stream_points | 891598.0 | 1240.76 MB/s | 67136 | 5 |

## StreamShapes

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalars/inmemory | 13527.0 | 1626.51 MB/s | 0 | 0 |
| scalars/stream | 24846.0 | 885.51 MB/s | 67120 | 4 |
| scalars/stream_reused | 19543.0 | 1125.80 MB/s | 32 | 1 |
| strings/inmemory | 6170.0 | 3728.11 MB/s | 0 | 0 |
| strings/stream | 12842.0 | 1791.11 MB/s | 67120 | 4 |
| strings/stream_reused | 7390.0 | 3112.64 MB/s | 32 | 1 |
| records/inmemory | 5701.0 | 4945.88 MB/s | 0 | 0 |
| records/stream | 12101.0 | 2330.31 MB/s | 67120 | 4 |
| records/stream_reused | 5937.0 | 4749.41 MB/s | 32 | 1 |

## StreamSkipToKey

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 59158.0 | 14857.25 MB/s | 67136 | 5 |

## StreamObjectEach

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stream | 613.6 | 2128.29 MB/s | 32 | 1 |
| inmemory | 500.6 | 2608.85 MB/s | 0 | 0 |

## StreamDescent

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| get | 1150.0 | 3135.20 MB/s | 32 | 1 |
| arrayeach | 1159.0 | 3111.27 MB/s | 32 | 1 |
| inmemory | 980.3 | 3679.32 MB/s | 0 | 0 |

## StreamLargeElements

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 124223.0 | 9589.72 MB/s | 47 | 1 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1679.0 | 1647.57 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1651.0 | 1675.02 MB/s | 0 | 0 |

## UnescapeStringCopy

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 19.6 | — | 32 | 1 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 170.2 | 1093.11 MB/s | 0 | 0 |

## ArrayEachStrings

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 604.0 | 3809.76 MB/s | 0 | 0 |

## ObjectEachRecord

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 547.8 | 3306.08 MB/s | 0 | 0 |

## ObjectEachRecordCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 522.0 | 3469.21 MB/s | 0 | 0 |

## ObjectEachNested

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 513.6 | 3526.02 MB/s | 0 | 0 |

## ArrayEachRecords

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 548.3 | 5049.81 MB/s | 0 | 0 |

## ArrayEachScalars

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1416.0 | 1554.82 MB/s | 0 | 0 |

## ArrayEachSeries

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2050.0 | 646.45 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 68.4 | — | 24 | 1 |
| arena | 61.1 | — | 24 | 0 |

## EscapeScratch

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| carve | 21.5 | — | 64 | 0 |
| make | 19.7 | — | 64 | 1 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 24476.0 | 802.08 MB/s | 0 | 0 |
| kernel/sep"," | 10184.0 | 1927.81 MB/s | 0 | 0 |
| scalar/sep",_" | 26711.0 | 884.70 MB/s | 0 | 0 |
| kernel/sep",_" | 13692.0 | 1725.86 MB/s | 0 | 0 |
| kernel-only | 10154.0 | 1933.52 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6623.0 | 1219.96 MB/s | 0 | 0 |
| "12," | 7945.0 | 1520.50 MB/s | 0 | 0 |
| "123," | 7479.0 | 2150.14 MB/s | 0 | 0 |
| "1234," | 10131.0 | 1982.11 MB/s | 0 | 0 |
| "123456," | 12569.0 | 2234.00 MB/s | 0 | 0 |
| "1234567," | 8608.0 | 3726.82 MB/s | 0 | 0 |
| "1234,_" | 15859.0 | 1518.35 MB/s | 0 | 0 |

## ParseIntUnstable

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| d01 | 3.6 | — | 0 | 0 |
| d03 | 4.7 | — | 0 | 0 |
| d05 | 3.8 | — | 0 | 0 |
| d08 | 3.8 | — | 0 | 0 |
| d10 | 5.3 | — | 0 | 0 |
| d13 | 5.3 | — | 0 | 0 |
| d16 | 5.4 | — | 0 | 0 |
| d19 | 7.2 | — | 0 | 0 |
| d20 | 7.0 | — | 0 | 0 |

## ValueScanner

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| records/whole | 9074.0 | 1014.86 MB/s | 0 | 0 |
| records/4k | 8998.0 | 1023.48 MB/s | 0 | 0 |
| strings/whole | 3694.0 | 3032.96 MB/s | 0 | 0 |
| strings/4k | 3738.0 | 2997.49 MB/s | 0 | 0 |
| numbers/whole | 194.5 | 43196.95 MB/s | 0 | 0 |
| numbers/4k | 242.3 | 34682.80 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 414.2 | — | 0 | 0 |
| canada | 308.0 | — | 0 | 0 |
| mesh | 280.7 | — | 0 | 0 |
| array | 211.8 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 415.4 | — | 0 | 0 |
| canada | 387.7 | — | 0 | 0 |
| mesh | 279.1 | — | 0 | 0 |
| array | 243.9 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3134.0 | 7657.63 MB/s | 0 | 0 |
| numberObj/goloop | 1095.0 | 9313.38 MB/s | 0 | 0 |
| nestedMixed/goloop | 1430.0 | 7553.79 MB/s | 0 | 0 |
| stringObj/avx2 | 1720.0 | 13948.62 MB/s | 0 | 0 |
| numberObj/avx2 | 617.0 | 16529.51 MB/s | 0 | 0 |
| nestedMixed/avx2 | 924.7 | 11680.62 MB/s | 0 | 0 |
| stringObj/avx512 | 1105.0 | 21709.44 MB/s | 0 | 0 |
| numberObj/avx512 | 412.9 | 24697.25 MB/s | 0 | 0 |
| nestedMixed/avx512 | 792.7 | 13624.77 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 8404.0 | 2855.55 MB/s | 0 | 0 |
| stringObj/dispatch | 1107.0 | 21680.97 MB/s | 0 | 0 |
| numberObj/current | 3155.0 | 3232.34 MB/s | 0 | 0 |
| numberObj/dispatch | 415.5 | 24546.74 MB/s | 0 | 0 |
| numberArr/current | 170.1 | 38807.03 MB/s | 0 | 0 |
| numberArr/dispatch | 173.6 | 38015.96 MB/s | 0 | 0 |
| nestedMixed/current | 12528.0 | 862.15 MB/s | 0 | 0 |
| nestedMixed/dispatch | 797.5 | 13543.46 MB/s | 0 | 0 |

## SkipSmall

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 12.7 | 1419.90 MB/s | 0 | 0 |
| record | 7.7 | 7052.03 MB/s | 0 | 0 |
| tiny | 7.6 | 915.41 MB/s | 0 | 0 |
| twoBlock | 9.7 | 9069.50 MB/s | 0 | 0 |

## SkipSmallScalar

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| pair | 14.1 | 1278.26 MB/s | 0 | 0 |
| record | 34.2 | 1577.77 MB/s | 0 | 0 |
| tiny | 11.6 | 604.68 MB/s | 0 | 0 |
| twoBlock | 19.5 | 4509.38 MB/s | 0 | 0 |
