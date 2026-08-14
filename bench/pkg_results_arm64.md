# lightning main-module benchmarks

- generated 2026-08-14T16:22:12Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.2 | 530.57 MB/s | 16 | 1 |
| sentence_clean | 36.8 | 1196.40 MB/s | 48 | 1 |
| url_clean | 40.5 | 1282.88 MB/s | 64 | 1 |
| log_line_clean | 115.6 | 2906.57 MB/s | 352 | 1 |
| path_with_backslash | 123.3 | 300.13 MB/s | 56 | 2 |
| json_in_json | 163.2 | 257.43 MB/s | 72 | 2 |
| prose_with_quotes | 96.5 | 393.58 MB/s | 64 | 2 |
| control_bytes | 118.7 | 202.23 MB/s | 56 | 2 |
| mostly_clean_one_quote | 136.9 | 2227.49 MB/s | 320 | 1 |
| unicode_clean | 280.9 | 840.13 MB/s | 240 | 1 |
| unicode_with_quotes | 160.2 | 393.19 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 148.8 | 2050.33 MB/s | 320 | 1 |
| invalid_utf8_dense | 622.2 | 192.86 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.0 | 2001.98 MB/s | 0 | 0 |
| sentence_clean | 16.3 | 2694.10 MB/s | 0 | 0 |
| url_clean | 11.1 | 4699.62 MB/s | 0 | 0 |
| log_line_clean | 38.0 | 8842.09 MB/s | 0 | 0 |
| path_with_backslash | 58.1 | 637.06 MB/s | 0 | 0 |
| json_in_json | 94.4 | 444.72 MB/s | 0 | 0 |
| prose_with_quotes | 34.2 | 1112.04 MB/s | 0 | 0 |
| control_bytes | 51.4 | 466.62 MB/s | 0 | 0 |
| mostly_clean_one_quote | 39.4 | 7746.08 MB/s | 0 | 0 |
| unicode_clean | 231.6 | 1019.04 MB/s | 0 | 0 |
| unicode_with_quotes | 90.5 | 696.46 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 53.4 | 5708.54 MB/s | 0 | 0 |
| invalid_utf8_dense | 428.5 | 280.03 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3019.0 | 3325.77 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2883.0 | 3482.81 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2790.0 | 3598.41 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12353.0 | 812.87 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 903.2 | 2005.10 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1376.0 | 1316.57 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.4 | 3615.37 MB/s | 0 | 0 |
| sentence_clean | 4.8 | 9115.47 MB/s | 0 | 0 |
| url_clean | 4.8 | 10774.23 MB/s | 0 | 0 |
| log_line_clean | 10.6 | 31590.66 MB/s | 0 | 0 |
| path_escaped | 83.7 | 514.07 MB/s | 48 | 1 |
| json_in_json | 116.9 | 461.81 MB/s | 64 | 1 |
| prose_with_quotes | 70.6 | 581.10 MB/s | 48 | 1 |
| unicode_heavy | 4.1 | 7267.50 MB/s | 0 | 0 |
| unicode_escaped_dense | 293.7 | 653.79 MB/s | 192 | 1 |
| mostly_clean_one_escape | 119.4 | 2562.61 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.0 | 3192.31 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7855.50 MB/s | 0 | 0 |
| url_clean | 5.6 | 9283.15 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29587.70 MB/s | 0 | 0 |
| path_escaped | 46.3 | 929.45 MB/s | 0 | 0 |
| json_in_json | 69.4 | 778.29 MB/s | 0 | 0 |
| prose_with_quotes | 32.1 | 1275.20 MB/s | 0 | 0 |
| unicode_heavy | 4.9 | 6164.99 MB/s | 0 | 0 |
| unicode_escaped_dense | 225.2 | 852.39 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.6 | 12418.74 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 71.5 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 52.7 | — | 0 | 0 |
| create_nested | 42.2 | — | 0 | 0 |
| overwrite_nonobject | 53.9 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 118.4 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 295.1 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 106.2 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 130.7 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2156.0 | 1282.67 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2066.0 | 1339.04 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 251.5 | 739.48 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 87.4 | — | 24 | 1 |
| arena | 71.4 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4393.0 | 5462.85 MB/s | 0 | 0 |
| numberObj/goloop | 1821.0 | 5600.85 MB/s | 0 | 0 |
| nestedMixed/goloop | 2409.0 | 4483.08 MB/s | 0 | 0 |
| stringObj/neon | 2910.0 | 8247.99 MB/s | 0 | 0 |
| numberObj/neon | 1216.0 | 8383.84 MB/s | 0 | 0 |
| nestedMixed/neon | 1650.0 | 6544.69 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 12368.0 | 1940.28 MB/s | 0 | 0 |
| stringObj/dispatch | 2884.0 | 8320.97 MB/s | 0 | 0 |
| numberObj/current | 5954.0 | 1712.69 MB/s | 0 | 0 |
| numberObj/dispatch | 1218.0 | 8376.15 MB/s | 0 | 0 |
| numberArr/current | 296.9 | 22229.82 MB/s | 0 | 0 |
| numberArr/dispatch | 298.0 | 22153.17 MB/s | 0 | 0 |
| nestedMixed/current | 16055.0 | 672.76 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1652.0 | 6539.70 MB/s | 0 | 0 |
