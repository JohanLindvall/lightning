# lightning main-module benchmarks

- generated 2026-08-11T11:14:22Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.5 | 524.83 MB/s | 16 | 1 |
| sentence_clean | 45.5 | 967.52 MB/s | 48 | 1 |
| url_clean | 40.0 | 1298.59 MB/s | 64 | 1 |
| log_line_clean | 115.5 | 2908.62 MB/s | 352 | 1 |
| path_with_backslash | 119.7 | 309.01 MB/s | 56 | 2 |
| json_in_json | 159.4 | 263.48 MB/s | 72 | 2 |
| prose_with_quotes | 94.6 | 401.53 MB/s | 64 | 2 |
| control_bytes | 116.1 | 206.80 MB/s | 56 | 2 |
| mostly_clean_one_quote | 138.7 | 2199.64 MB/s | 320 | 1 |
| unicode_clean | 282.6 | 835.11 MB/s | 240 | 1 |
| unicode_with_quotes | 164.3 | 383.53 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 149.7 | 2037.23 MB/s | 320 | 1 |
| invalid_utf8_dense | 620.6 | 193.35 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.0 | 1996.02 MB/s | 0 | 0 |
| sentence_clean | 16.5 | 2670.95 MB/s | 0 | 0 |
| url_clean | 22.6 | 2304.93 MB/s | 0 | 0 |
| log_line_clean | 49.1 | 6848.07 MB/s | 0 | 0 |
| path_with_backslash | 57.5 | 643.39 MB/s | 0 | 0 |
| json_in_json | 94.1 | 446.27 MB/s | 0 | 0 |
| prose_with_quotes | 34.4 | 1104.45 MB/s | 0 | 0 |
| control_bytes | 51.2 | 468.45 MB/s | 0 | 0 |
| mostly_clean_one_quote | 50.2 | 6074.98 MB/s | 0 | 0 |
| unicode_clean | 228.0 | 1034.93 MB/s | 0 | 0 |
| unicode_with_quotes | 91.9 | 685.70 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 65.5 | 4658.80 MB/s | 0 | 0 |
| invalid_utf8_dense | 440.7 | 272.27 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3075.0 | 3265.26 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2919.0 | 3439.76 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2916.0 | 3443.93 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12799.0 | 784.48 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 987.0 | 1834.84 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1467.0 | 1234.75 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.4 | 3613.88 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9014.07 MB/s | 0 | 0 |
| url_clean | 4.9 | 10648.46 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31301.53 MB/s | 0 | 0 |
| path_escaped | 79.8 | 538.70 MB/s | 48 | 1 |
| json_in_json | 110.1 | 490.27 MB/s | 64 | 1 |
| prose_with_quotes | 66.8 | 613.90 MB/s | 48 | 1 |
| unicode_heavy | 4.2 | 7085.08 MB/s | 0 | 0 |
| unicode_escaped_dense | 290.3 | 661.36 MB/s | 192 | 1 |
| mostly_clean_one_escape | 114.4 | 2675.04 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.0 | 3192.86 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7854.72 MB/s | 0 | 0 |
| url_clean | 5.6 | 9283.09 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29586.74 MB/s | 0 | 0 |
| path_escaped | 45.5 | 945.20 MB/s | 0 | 0 |
| json_in_json | 67.6 | 798.38 MB/s | 0 | 0 |
| prose_with_quotes | 32.3 | 1267.89 MB/s | 0 | 0 |
| unicode_heavy | 4.9 | 6162.50 MB/s | 0 | 0 |
| unicode_escaped_dense | 224.9 | 853.54 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.6 | 12421.13 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 84.6 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 57.1 | — | 0 | 0 |
| create_nested | 46.4 | — | 0 | 0 |
| overwrite_nonobject | 55.3 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 125.2 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 291.3 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 112.2 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 135.3 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2364.0 | 1170.00 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2264.0 | 1221.96 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 272.1 | 683.68 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 87.6 | — | 24 | 1 |
| arena | 72.0 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4541.0 | 5284.65 MB/s | 0 | 0 |
| numberObj/goloop | 1840.0 | 5543.47 MB/s | 0 | 0 |
| nestedMixed/goloop | 2492.0 | 4334.09 MB/s | 0 | 0 |
| stringObj/neon | 2974.0 | 8069.09 MB/s | 0 | 0 |
| numberObj/neon | 1263.0 | 8077.41 MB/s | 0 | 0 |
| nestedMixed/neon | 1671.0 | 6464.45 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14498.0 | 1655.30 MB/s | 0 | 0 |
| stringObj/dispatch | 2975.0 | 8066.30 MB/s | 0 | 0 |
| numberObj/current | 6298.0 | 1619.14 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8071.16 MB/s | 0 | 0 |
| numberArr/current | 449.0 | 14699.95 MB/s | 0 | 0 |
| numberArr/dispatch | 451.4 | 14624.56 MB/s | 0 | 0 |
| nestedMixed/current | 17394.0 | 620.96 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1676.0 | 6445.37 MB/s | 0 | 0 |
