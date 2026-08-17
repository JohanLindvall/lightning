# lightning main-module benchmarks

- generated 2026-08-17T08:43:29Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 29.6 | 540.62 MB/s | 16 | 1 |
| sentence_clean | 36.2 | 1214.43 MB/s | 48 | 1 |
| url_clean | 39.4 | 1319.99 MB/s | 64 | 1 |
| log_line_clean | 112.3 | 2991.17 MB/s | 352 | 1 |
| path_with_backslash | 116.2 | 318.36 MB/s | 56 | 2 |
| json_in_json | 151.0 | 278.22 MB/s | 72 | 2 |
| prose_with_quotes | 93.7 | 405.69 MB/s | 64 | 2 |
| control_bytes | 113.1 | 212.14 MB/s | 56 | 2 |
| mostly_clean_one_quote | 132.9 | 2294.20 MB/s | 320 | 1 |
| unicode_clean | 278.0 | 848.83 MB/s | 240 | 1 |
| unicode_with_quotes | 158.1 | 398.41 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 141.6 | 2153.82 MB/s | 320 | 1 |
| invalid_utf8_dense | 635.1 | 188.94 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.9 | 2013.79 MB/s | 0 | 0 |
| sentence_clean | 15.4 | 2857.08 MB/s | 0 | 0 |
| url_clean | 10.6 | 4888.26 MB/s | 0 | 0 |
| log_line_clean | 33.5 | 10044.47 MB/s | 0 | 0 |
| path_with_backslash | 52.4 | 705.66 MB/s | 0 | 0 |
| json_in_json | 82.7 | 508.00 MB/s | 0 | 0 |
| prose_with_quotes | 31.3 | 1213.28 MB/s | 0 | 0 |
| control_bytes | 47.1 | 509.55 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.5 | 8590.97 MB/s | 0 | 0 |
| unicode_clean | 228.2 | 1034.37 MB/s | 0 | 0 |
| unicode_with_quotes | 86.2 | 730.59 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 50.3 | 6058.89 MB/s | 0 | 0 |
| invalid_utf8_dense | 445.5 | 269.38 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2959.0 | 3392.97 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2895.0 | 3467.88 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2748.0 | 3653.42 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 11257.0 | 891.99 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 823.2 | 2199.97 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1272.0 | 1423.77 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.4 | 3604.55 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9015.95 MB/s | 0 | 0 |
| url_clean | 4.9 | 10634.15 MB/s | 0 | 0 |
| log_line_clean | 10.6 | 31566.69 MB/s | 0 | 0 |
| path_escaped | 81.2 | 529.82 MB/s | 48 | 1 |
| json_in_json | 111.7 | 483.46 MB/s | 64 | 1 |
| prose_with_quotes | 67.4 | 608.21 MB/s | 48 | 1 |
| unicode_heavy | 4.1 | 7247.40 MB/s | 0 | 0 |
| unicode_escaped_dense | 289.5 | 663.18 MB/s | 192 | 1 |
| mostly_clean_one_escape | 116.6 | 2625.08 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.0 | 3192.11 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8065.50 MB/s | 0 | 0 |
| url_clean | 5.5 | 9533.34 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29587.68 MB/s | 0 | 0 |
| path_escaped | 45.2 | 950.93 MB/s | 0 | 0 |
| json_in_json | 67.9 | 795.23 MB/s | 0 | 0 |
| prose_with_quotes | 32.8 | 1248.81 MB/s | 0 | 0 |
| unicode_heavy | 4.9 | 6164.01 MB/s | 0 | 0 |
| unicode_escaped_dense | 220.6 | 870.35 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.6 | 12437.78 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 66.9 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 49.2 | — | 0 | 0 |
| create_nested | 41.0 | — | 0 | 0 |
| overwrite_nonobject | 49.8 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 110.3 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 286.3 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 100.8 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 123.4 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2005.0 | 1379.57 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1908.0 | 1449.88 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 228.9 | 812.75 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 87.9 | — | 24 | 1 |
| arena | 71.9 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4392.0 | 5463.76 MB/s | 0 | 0 |
| numberObj/goloop | 1828.0 | 5578.52 MB/s | 0 | 0 |
| nestedMixed/goloop | 2570.0 | 4202.90 MB/s | 0 | 0 |
| stringObj/neon | 2905.0 | 8260.35 MB/s | 0 | 0 |
| numberObj/neon | 1216.0 | 8384.41 MB/s | 0 | 0 |
| nestedMixed/neon | 1650.0 | 6547.12 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10268.0 | 2337.18 MB/s | 0 | 0 |
| stringObj/dispatch | 2883.0 | 8323.54 MB/s | 0 | 0 |
| numberObj/current | 5611.0 | 1817.47 MB/s | 0 | 0 |
| numberObj/dispatch | 1217.0 | 8378.66 MB/s | 0 | 0 |
| numberArr/current | 158.6 | 41618.23 MB/s | 0 | 0 |
| numberArr/dispatch | 160.5 | 41140.52 MB/s | 0 | 0 |
| nestedMixed/current | 14305.0 | 755.05 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1652.0 | 6538.31 MB/s | 0 | 0 |
