# lightning main-module benchmarks

- generated 2026-09-03T06:38:09Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.4 | 544.73 MB/s | 16 | 1 |
| sentence_clean | 35.9 | 1227.25 MB/s | 48 | 1 |
| url_clean | 39.0 | 1332.89 MB/s | 64 | 1 |
| log_line_clean | 112.0 | 3000.13 MB/s | 352 | 1 |
| path_with_backslash | 115.3 | 320.77 MB/s | 56 | 2 |
| json_in_json | 152.5 | 275.39 MB/s | 72 | 2 |
| prose_with_quotes | 91.2 | 416.78 MB/s | 64 | 2 |
| control_bytes | 110.5 | 217.14 MB/s | 56 | 2 |
| mostly_clean_one_quote | 131.8 | 2313.71 MB/s | 320 | 1 |
| unicode_clean | 270.7 | 871.69 MB/s | 240 | 1 |
| unicode_with_quotes | 158.1 | 398.56 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 142.3 | 2143.74 MB/s | 320 | 1 |
| invalid_utf8_dense | 627.3 | 191.28 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.5 | 2126.21 MB/s | 0 | 0 |
| sentence_clean | 15.0 | 2934.40 MB/s | 0 | 0 |
| url_clean | 10.8 | 4820.69 MB/s | 0 | 0 |
| log_line_clean | 33.3 | 10086.50 MB/s | 0 | 0 |
| path_with_backslash | 54.2 | 682.99 MB/s | 0 | 0 |
| json_in_json | 85.1 | 493.74 MB/s | 0 | 0 |
| prose_with_quotes | 31.2 | 1216.56 MB/s | 0 | 0 |
| control_bytes | 45.6 | 526.62 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.0 | 8725.82 MB/s | 0 | 0 |
| unicode_clean | 223.8 | 1054.67 MB/s | 0 | 0 |
| unicode_with_quotes | 85.8 | 733.99 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 49.7 | 6135.86 MB/s | 0 | 0 |
| invalid_utf8_dense | 434.6 | 276.11 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2881.0 | 3485.58 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2727.0 | 3682.69 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2706.0 | 3710.20 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10528.0 | 953.78 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 778.9 | 2325.00 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1222.0 | 1481.44 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.1 | 3859.31 MB/s | 0 | 0 |
| sentence_clean | 4.7 | 9450.42 MB/s | 0 | 0 |
| url_clean | 4.7 | 11169.30 MB/s | 0 | 0 |
| log_line_clean | 10.5 | 31891.43 MB/s | 0 | 0 |
| path_escaped | 82.7 | 520.00 MB/s | 48 | 1 |
| json_in_json | 112.6 | 479.43 MB/s | 64 | 1 |
| prose_with_quotes | 69.4 | 590.67 MB/s | 48 | 1 |
| unicode_heavy | 5.0 | 5951.84 MB/s | 0 | 0 |
| unicode_escaped_dense | 285.9 | 671.59 MB/s | 192 | 1 |
| mostly_clean_one_escape | 116.8 | 2618.84 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3275.36 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8287.45 MB/s | 0 | 0 |
| url_clean | 5.3 | 9795.97 MB/s | 0 | 0 |
| log_line_clean | 11.2 | 29969.08 MB/s | 0 | 0 |
| path_escaped | 45.2 | 952.27 MB/s | 0 | 0 |
| json_in_json | 69.0 | 782.30 MB/s | 0 | 0 |
| prose_with_quotes | 32.0 | 1279.48 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5196.72 MB/s | 0 | 0 |
| unicode_escaped_dense | 221.3 | 867.46 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.5 | 12501.98 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 65.7 | — | 0 | 0 |
| append_empty | 18.2 | — | 0 | 0 |
| replace | 48.6 | — | 0 | 0 |
| create_nested | 41.9 | — | 0 | 0 |
| overwrite_nonobject | 52.1 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 108.4 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 290.8 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 100.1 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 123.5 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1970.0 | 1403.76 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1893.0 | 1460.82 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 221.6 | 839.42 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 90.1 | — | 24 | 1 |
| arena | 71.7 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 22973.0 | 854.57 MB/s | 0 | 0 |
| kernel/sep"," | 8227.0 | 2386.24 MB/s | 0 | 0 |
| scalar/sep",_" | 25210.0 | 937.38 MB/s | 0 | 0 |
| kernel/sep",_" | 11008.0 | 2146.66 MB/s | 0 | 0 |
| kernel-only | 8551.0 | 2295.77 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6873.0 | 1175.63 MB/s | 0 | 0 |
| "12," | 7186.0 | 1681.01 MB/s | 0 | 0 |
| "123," | 7541.0 | 2132.38 MB/s | 0 | 0 |
| "1234," | 8186.0 | 2452.88 MB/s | 0 | 0 |
| "123456," | 8704.0 | 3226.01 MB/s | 0 | 0 |
| "1234567," | 8618.0 | 3722.65 MB/s | 0 | 0 |
| "1234,_" | 10785.0 | 2232.60 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 460.2 | — | 0 | 0 |
| canada | 333.9 | — | 0 | 0 |
| mesh | 272.2 | — | 0 | 0 |
| array | 223.2 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 458.6 | — | 0 | 0 |
| canada | 409.9 | — | 0 | 0 |
| mesh | 292.5 | — | 0 | 0 |
| array | 261.8 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4397.0 | 5457.48 MB/s | 0 | 0 |
| numberObj/goloop | 1820.0 | 5602.57 MB/s | 0 | 0 |
| nestedMixed/goloop | 2409.0 | 4482.93 MB/s | 0 | 0 |
| stringObj/neon | 2901.0 | 8270.91 MB/s | 0 | 0 |
| numberObj/neon | 1223.0 | 8337.65 MB/s | 0 | 0 |
| nestedMixed/neon | 1652.0 | 6537.34 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10279.0 | 2334.56 MB/s | 0 | 0 |
| stringObj/dispatch | 2902.0 | 8268.54 MB/s | 0 | 0 |
| numberObj/current | 5627.0 | 1812.45 MB/s | 0 | 0 |
| numberObj/dispatch | 1223.0 | 8336.36 MB/s | 0 | 0 |
| numberArr/current | 159.1 | 41487.27 MB/s | 0 | 0 |
| numberArr/dispatch | 160.4 | 41149.86 MB/s | 0 | 0 |
| nestedMixed/current | 16934.0 | 637.82 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1653.0 | 6533.72 MB/s | 0 | 0 |
