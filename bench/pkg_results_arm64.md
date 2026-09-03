# lightning main-module benchmarks

- generated 2026-09-03T00:44:49Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.8 | 536.22 MB/s | 16 | 1 |
| sentence_clean | 37.1 | 1185.08 MB/s | 48 | 1 |
| url_clean | 40.7 | 1278.85 MB/s | 64 | 1 |
| log_line_clean | 116.1 | 2894.48 MB/s | 352 | 1 |
| path_with_backslash | 118.3 | 312.85 MB/s | 56 | 2 |
| json_in_json | 154.8 | 271.38 MB/s | 72 | 2 |
| prose_with_quotes | 92.8 | 409.51 MB/s | 64 | 2 |
| control_bytes | 113.0 | 212.34 MB/s | 56 | 2 |
| mostly_clean_one_quote | 134.3 | 2271.54 MB/s | 320 | 1 |
| unicode_clean | 271.8 | 868.20 MB/s | 240 | 1 |
| unicode_with_quotes | 159.5 | 394.92 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 145.7 | 2093.35 MB/s | 320 | 1 |
| invalid_utf8_dense | 627.8 | 191.15 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.5 | 2123.60 MB/s | 0 | 0 |
| sentence_clean | 15.0 | 2938.93 MB/s | 0 | 0 |
| url_clean | 10.8 | 4811.23 MB/s | 0 | 0 |
| log_line_clean | 33.2 | 10111.41 MB/s | 0 | 0 |
| path_with_backslash | 53.6 | 689.92 MB/s | 0 | 0 |
| json_in_json | 84.9 | 494.52 MB/s | 0 | 0 |
| prose_with_quotes | 31.2 | 1219.12 MB/s | 0 | 0 |
| control_bytes | 45.6 | 525.91 MB/s | 0 | 0 |
| mostly_clean_one_quote | 34.9 | 8729.45 MB/s | 0 | 0 |
| unicode_clean | 223.5 | 1055.76 MB/s | 0 | 0 |
| unicode_with_quotes | 86.0 | 732.78 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 49.8 | 6126.52 MB/s | 0 | 0 |
| invalid_utf8_dense | 435.0 | 275.84 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2889.0 | 3475.99 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2715.0 | 3698.34 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2706.0 | 3709.98 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10391.0 | 966.33 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 779.4 | 2323.59 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1220.0 | 1484.72 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3855.66 MB/s | 0 | 0 |
| sentence_clean | 4.7 | 9456.28 MB/s | 0 | 0 |
| url_clean | 4.7 | 11177.92 MB/s | 0 | 0 |
| log_line_clean | 10.5 | 31899.15 MB/s | 0 | 0 |
| path_escaped | 81.2 | 529.51 MB/s | 48 | 1 |
| json_in_json | 111.5 | 484.29 MB/s | 64 | 1 |
| prose_with_quotes | 68.0 | 603.08 MB/s | 48 | 1 |
| unicode_heavy | 5.0 | 5952.93 MB/s | 0 | 0 |
| unicode_escaped_dense | 292.2 | 657.00 MB/s | 192 | 1 |
| mostly_clean_one_escape | 120.6 | 2538.08 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3276.39 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8289.10 MB/s | 0 | 0 |
| url_clean | 5.3 | 9797.04 MB/s | 0 | 0 |
| log_line_clean | 11.2 | 29971.87 MB/s | 0 | 0 |
| path_escaped | 45.0 | 954.82 MB/s | 0 | 0 |
| json_in_json | 69.0 | 783.05 MB/s | 0 | 0 |
| prose_with_quotes | 32.0 | 1281.02 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5195.93 MB/s | 0 | 0 |
| unicode_escaped_dense | 220.8 | 869.72 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.5 | 12498.49 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 65.8 | — | 0 | 0 |
| append_empty | 18.2 | — | 0 | 0 |
| replace | 48.7 | — | 0 | 0 |
| create_nested | 41.9 | — | 0 | 0 |
| overwrite_nonobject | 52.0 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 108.7 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 290.8 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 99.5 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 123.4 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1970.0 | 1403.84 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1889.0 | 1464.51 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 222.1 | 837.63 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 90.1 | — | 24 | 1 |
| arena | 71.8 | — | 24 | 0 |

## DecodeIntSliceRun

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| scalar/sep"," | 22928.0 | 856.23 MB/s | 0 | 0 |
| kernel/sep"," | 8556.0 | 2294.43 MB/s | 0 | 0 |
| scalar/sep",_" | 25151.0 | 939.57 MB/s | 0 | 0 |
| kernel/sep",_" | 11005.0 | 2147.24 MB/s | 0 | 0 |
| kernel-only | 8530.0 | 2301.47 MB/s | 0 | 0 |

## ParseIntRunShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| "1," | 6878.0 | 1174.83 MB/s | 0 | 0 |
| "12," | 7186.0 | 1681.10 MB/s | 0 | 0 |
| "123," | 7593.0 | 2117.82 MB/s | 0 | 0 |
| "1234," | 8579.0 | 2340.55 MB/s | 0 | 0 |
| "123456," | 8689.0 | 3231.67 MB/s | 0 | 0 |
| "1234567," | 8533.0 | 3759.64 MB/s | 0 | 0 |
| "1234,_" | 10895.0 | 2210.13 MB/s | 0 | 0 |

## ScanFloatShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 459.9 | — | 0 | 0 |
| canada | 332.8 | — | 0 | 0 |
| mesh | 271.9 | — | 0 | 0 |
| array | 223.2 | — | 0 | 0 |

## ScanFloatSlowShapes

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| slow | 458.8 | — | 0 | 0 |
| canada | 412.8 | — | 0 | 0 |
| mesh | 292.4 | — | 0 | 0 |
| array | 261.7 | — | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4391.0 | 5465.77 MB/s | 0 | 0 |
| numberObj/goloop | 1820.0 | 5603.33 MB/s | 0 | 0 |
| nestedMixed/goloop | 2411.0 | 4479.24 MB/s | 0 | 0 |
| stringObj/neon | 2901.0 | 8271.62 MB/s | 0 | 0 |
| numberObj/neon | 1225.0 | 8323.00 MB/s | 0 | 0 |
| nestedMixed/neon | 1651.0 | 6542.77 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10291.0 | 2331.92 MB/s | 0 | 0 |
| stringObj/dispatch | 2901.0 | 8271.92 MB/s | 0 | 0 |
| numberObj/current | 5622.0 | 1813.83 MB/s | 0 | 0 |
| numberObj/dispatch | 1223.0 | 8335.13 MB/s | 0 | 0 |
| numberArr/current | 159.0 | 41510.92 MB/s | 0 | 0 |
| numberArr/dispatch | 160.3 | 41174.58 MB/s | 0 | 0 |
| nestedMixed/current | 15628.0 | 691.11 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1656.0 | 6523.47 MB/s | 0 | 0 |
