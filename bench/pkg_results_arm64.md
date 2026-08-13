# lightning main-module benchmarks

- generated 2026-08-13T07:24:09Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.3 | 527.53 MB/s | 16 | 1 |
| sentence_clean | 46.2 | 952.95 MB/s | 48 | 1 |
| url_clean | 43.1 | 1205.02 MB/s | 64 | 1 |
| log_line_clean | 123.7 | 2715.32 MB/s | 352 | 1 |
| path_with_backslash | 119.6 | 309.25 MB/s | 56 | 2 |
| json_in_json | 160.0 | 262.42 MB/s | 72 | 2 |
| prose_with_quotes | 94.5 | 402.09 MB/s | 64 | 2 |
| control_bytes | 115.5 | 207.76 MB/s | 56 | 2 |
| mostly_clean_one_quote | 146.4 | 2083.81 MB/s | 320 | 1 |
| unicode_clean | 281.0 | 839.98 MB/s | 240 | 1 |
| unicode_with_quotes | 163.8 | 384.61 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 155.5 | 1960.94 MB/s | 320 | 1 |
| invalid_utf8_dense | 614.0 | 195.45 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.0 | 1996.29 MB/s | 0 | 0 |
| sentence_clean | 16.5 | 2673.70 MB/s | 0 | 0 |
| url_clean | 22.3 | 2330.88 MB/s | 0 | 0 |
| log_line_clean | 48.6 | 6911.84 MB/s | 0 | 0 |
| path_with_backslash | 57.8 | 640.29 MB/s | 0 | 0 |
| json_in_json | 94.4 | 444.81 MB/s | 0 | 0 |
| prose_with_quotes | 34.5 | 1100.93 MB/s | 0 | 0 |
| control_bytes | 51.2 | 468.91 MB/s | 0 | 0 |
| mostly_clean_one_quote | 49.8 | 6124.06 MB/s | 0 | 0 |
| unicode_clean | 226.5 | 1042.15 MB/s | 0 | 0 |
| unicode_with_quotes | 91.4 | 689.56 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 65.1 | 4683.68 MB/s | 0 | 0 |
| invalid_utf8_dense | 439.1 | 273.26 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3011.0 | 3335.32 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2860.0 | 3511.31 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2844.0 | 3530.88 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12531.0 | 801.27 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 961.2 | 1884.10 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1440.0 | 1257.79 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.4 | 3615.07 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9015.64 MB/s | 0 | 0 |
| url_clean | 4.9 | 10652.82 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31302.73 MB/s | 0 | 0 |
| path_escaped | 74.0 | 581.16 MB/s | 48 | 1 |
| json_in_json | 102.9 | 524.60 MB/s | 64 | 1 |
| prose_with_quotes | 63.0 | 650.48 MB/s | 48 | 1 |
| unicode_heavy | 4.2 | 7085.48 MB/s | 0 | 0 |
| unicode_escaped_dense | 293.2 | 654.83 MB/s | 192 | 1 |
| mostly_clean_one_escape | 126.3 | 2422.79 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.0 | 3192.61 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7855.53 MB/s | 0 | 0 |
| url_clean | 5.6 | 9283.59 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29582.85 MB/s | 0 | 0 |
| path_escaped | 44.0 | 977.63 MB/s | 0 | 0 |
| json_in_json | 68.3 | 790.37 MB/s | 0 | 0 |
| prose_with_quotes | 32.2 | 1274.26 MB/s | 0 | 0 |
| unicode_heavy | 4.9 | 6163.43 MB/s | 0 | 0 |
| unicode_escaped_dense | 221.6 | 866.42 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.6 | 12415.74 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 83.2 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 55.7 | — | 0 | 0 |
| create_nested | 45.6 | — | 0 | 0 |
| overwrite_nonobject | 55.5 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 122.0 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 294.8 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 110.3 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 134.5 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2296.0 | 1204.59 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2192.0 | 1261.82 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 264.0 | 704.49 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 88.0 | — | 24 | 1 |
| arena | 72.6 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4399.0 | 5455.63 MB/s | 0 | 0 |
| numberObj/goloop | 1818.0 | 5608.94 MB/s | 0 | 0 |
| nestedMixed/goloop | 2407.0 | 4486.47 MB/s | 0 | 0 |
| stringObj/neon | 2905.0 | 8261.69 MB/s | 0 | 0 |
| numberObj/neon | 1216.0 | 8384.56 MB/s | 0 | 0 |
| nestedMixed/neon | 1650.0 | 6546.44 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14090.0 | 1703.23 MB/s | 0 | 0 |
| stringObj/dispatch | 2886.0 | 8314.35 MB/s | 0 | 0 |
| numberObj/current | 6285.0 | 1622.57 MB/s | 0 | 0 |
| numberObj/dispatch | 1217.0 | 8377.31 MB/s | 0 | 0 |
| numberArr/current | 448.5 | 14719.17 MB/s | 0 | 0 |
| numberArr/dispatch | 450.8 | 14643.13 MB/s | 0 | 0 |
| nestedMixed/current | 17019.0 | 634.63 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1651.0 | 6541.11 MB/s | 0 | 0 |
