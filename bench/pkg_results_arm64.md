# lightning main-module benchmarks

- generated 2026-07-26T15:04:14Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.0 | 516.38 MB/s | 16 | 1 |
| sentence_clean | 46.3 | 950.82 MB/s | 48 | 1 |
| url_clean | 43.5 | 1195.91 MB/s | 64 | 1 |
| log_line_clean | 121.0 | 2776.20 MB/s | 352 | 1 |
| path_with_backslash | 124.2 | 297.95 MB/s | 56 | 2 |
| json_in_json | 162.1 | 259.12 MB/s | 72 | 2 |
| prose_with_quotes | 98.0 | 387.83 MB/s | 64 | 2 |
| control_bytes | 119.6 | 200.64 MB/s | 56 | 2 |
| mostly_clean_one_quote | 142.5 | 2140.07 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1870.55 MB/s | 0 | 0 |
| sentence_clean | 16.8 | 2615.38 MB/s | 0 | 0 |
| url_clean | 22.4 | 2322.55 MB/s | 0 | 0 |
| log_line_clean | 44.1 | 7620.72 MB/s | 0 | 0 |
| path_with_backslash | 58.7 | 630.26 MB/s | 0 | 0 |
| json_in_json | 94.2 | 445.92 MB/s | 0 | 0 |
| prose_with_quotes | 34.8 | 1093.50 MB/s | 0 | 0 |
| control_bytes | 52.9 | 454.02 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.5 | 6701.00 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 986.1 | 1836.52 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1334.0 | 1357.53 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.4 | 3614.72 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9030.83 MB/s | 0 | 0 |
| url_clean | 4.9 | 10671.25 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31557.39 MB/s | 0 | 0 |
| path_escaped | 89.4 | 480.93 MB/s | 48 | 1 |
| json_in_json | 122.7 | 439.97 MB/s | 64 | 1 |
| prose_with_quotes | 75.8 | 540.69 MB/s | 48 | 1 |
| unicode_heavy | 4.2 | 7210.13 MB/s | 0 | 0 |
| unicode_escaped_dense | 344.8 | 556.78 MB/s | 192 | 1 |
| mostly_clean_one_escape | 124.8 | 2451.73 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.0 | 3192.12 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8067.05 MB/s | 0 | 0 |
| url_clean | 5.5 | 9533.47 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29593.20 MB/s | 0 | 0 |
| path_escaped | 46.8 | 918.33 MB/s | 0 | 0 |
| json_in_json | 74.6 | 723.95 MB/s | 0 | 0 |
| prose_with_quotes | 33.5 | 1223.26 MB/s | 0 | 0 |
| unicode_heavy | 4.9 | 6166.08 MB/s | 0 | 0 |
| unicode_escaped_dense | 269.4 | 712.76 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12373.13 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 86.8 | — | 0 | 0 |
| append_empty | 18.2 | — | 0 | 0 |
| replace | 59.4 | — | 0 | 0 |
| create_nested | 48.2 | — | 0 | 0 |
| overwrite_nonobject | 58.1 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 119.2 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 299.3 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2311.0 | 1197.13 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2261.0 | 1223.51 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 271.9 | 684.02 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4534.0 | 5292.97 MB/s | 0 | 0 |
| numberObj/goloop | 1843.0 | 5532.90 MB/s | 0 | 0 |
| nestedMixed/goloop | 2499.0 | 4322.98 MB/s | 0 | 0 |
| stringObj/neon | 2974.0 | 8068.15 MB/s | 0 | 0 |
| numberObj/neon | 1263.0 | 8073.00 MB/s | 0 | 0 |
| nestedMixed/neon | 1679.0 | 6432.70 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14439.0 | 1662.03 MB/s | 0 | 0 |
| stringObj/dispatch | 2975.0 | 8065.56 MB/s | 0 | 0 |
| numberObj/current | 6294.0 | 1620.17 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8068.03 MB/s | 0 | 0 |
| numberArr/current | 464.4 | 14212.89 MB/s | 0 | 0 |
| numberArr/dispatch | 466.6 | 14147.79 MB/s | 0 | 0 |
| nestedMixed/current | 17027.0 | 634.34 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1681.0 | 6425.50 MB/s | 0 | 0 |
