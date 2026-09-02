# lightning main-module benchmarks

- generated 2026-09-02T12:27:22Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.7 | 538.26 MB/s | 16 | 1 |
| sentence_clean | 36.9 | 1192.94 MB/s | 48 | 1 |
| url_clean | 40.4 | 1286.52 MB/s | 64 | 1 |
| log_line_clean | 115.9 | 2900.15 MB/s | 352 | 1 |
| path_with_backslash | 117.3 | 315.39 MB/s | 56 | 2 |
| json_in_json | 153.5 | 273.64 MB/s | 72 | 2 |
| prose_with_quotes | 93.5 | 406.43 MB/s | 64 | 2 |
| control_bytes | 113.3 | 211.78 MB/s | 56 | 2 |
| mostly_clean_one_quote | 135.6 | 2249.26 MB/s | 320 | 1 |
| unicode_clean | 271.2 | 870.10 MB/s | 240 | 1 |
| unicode_with_quotes | 164.3 | 383.42 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 144.9 | 2105.32 MB/s | 320 | 1 |
| invalid_utf8_dense | 637.1 | 188.35 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.6 | 2106.08 MB/s | 0 | 0 |
| sentence_clean | 15.1 | 2920.52 MB/s | 0 | 0 |
| url_clean | 10.6 | 4924.66 MB/s | 0 | 0 |
| log_line_clean | 33.3 | 10083.57 MB/s | 0 | 0 |
| path_with_backslash | 52.9 | 699.13 MB/s | 0 | 0 |
| json_in_json | 83.6 | 502.43 MB/s | 0 | 0 |
| prose_with_quotes | 31.1 | 1223.29 MB/s | 0 | 0 |
| control_bytes | 46.2 | 519.08 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.5 | 8584.56 MB/s | 0 | 0 |
| unicode_clean | 222.2 | 1062.01 MB/s | 0 | 0 |
| unicode_with_quotes | 85.9 | 733.63 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 49.9 | 6116.68 MB/s | 0 | 0 |
| invalid_utf8_dense | 439.3 | 273.17 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2888.0 | 3476.32 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2743.0 | 3661.12 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2717.0 | 3695.57 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 10915.0 | 919.90 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 791.6 | 2287.78 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1232.0 | 1470.21 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.3 | 3721.34 MB/s | 0 | 0 |
| sentence_clean | 4.7 | 9328.30 MB/s | 0 | 0 |
| url_clean | 4.7 | 11025.01 MB/s | 0 | 0 |
| log_line_clean | 10.6 | 31841.56 MB/s | 0 | 0 |
| path_escaped | 84.2 | 510.71 MB/s | 48 | 1 |
| json_in_json | 113.3 | 476.48 MB/s | 64 | 1 |
| prose_with_quotes | 70.3 | 583.34 MB/s | 48 | 1 |
| unicode_heavy | 4.0 | 7411.13 MB/s | 0 | 0 |
| unicode_escaped_dense | 290.7 | 660.39 MB/s | 192 | 1 |
| mostly_clean_one_escape | 120.6 | 2537.23 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.7 | 3390.80 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8291.46 MB/s | 0 | 0 |
| url_clean | 5.3 | 9799.09 MB/s | 0 | 0 |
| log_line_clean | 11.1 | 30272.55 MB/s | 0 | 0 |
| path_escaped | 43.4 | 990.07 MB/s | 0 | 0 |
| json_in_json | 65.7 | 821.94 MB/s | 0 | 0 |
| prose_with_quotes | 31.7 | 1292.35 MB/s | 0 | 0 |
| unicode_heavy | 4.6 | 6539.30 MB/s | 0 | 0 |
| unicode_escaped_dense | 220.2 | 871.85 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.4 | 12515.76 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 66.0 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 48.3 | — | 0 | 0 |
| create_nested | 42.0 | — | 0 | 0 |
| overwrite_nonobject | 49.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 110.6 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 294.3 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 98.8 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 126.3 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1985.0 | 1393.13 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1893.0 | 1461.07 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 229.8 | 809.27 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 91.3 | — | 24 | 1 |
| arena | 72.9 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4388.0 | 5468.44 MB/s | 0 | 0 |
| numberObj/goloop | 1820.0 | 5604.44 MB/s | 0 | 0 |
| nestedMixed/goloop | 2417.0 | 4469.51 MB/s | 0 | 0 |
| stringObj/neon | 2884.0 | 8322.43 MB/s | 0 | 0 |
| numberObj/neon | 1217.0 | 8376.66 MB/s | 0 | 0 |
| nestedMixed/neon | 1651.0 | 6540.92 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10334.0 | 2322.16 MB/s | 0 | 0 |
| stringObj/dispatch | 2885.0 | 8319.20 MB/s | 0 | 0 |
| numberObj/current | 5677.0 | 1796.33 MB/s | 0 | 0 |
| numberObj/dispatch | 1218.0 | 8372.44 MB/s | 0 | 0 |
| numberArr/current | 158.7 | 41581.80 MB/s | 0 | 0 |
| numberArr/dispatch | 160.3 | 41178.63 MB/s | 0 | 0 |
| nestedMixed/current | 14237.0 | 758.66 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1653.0 | 6532.27 MB/s | 0 | 0 |
