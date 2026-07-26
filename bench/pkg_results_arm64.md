# lightning main-module benchmarks

- generated 2026-07-26T10:33:31Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.1 | 532.40 MB/s | 16 | 1 |
| sentence_clean | 44.9 | 979.59 MB/s | 48 | 1 |
| url_clean | 40.9 | 1271.59 MB/s | 64 | 1 |
| log_line_clean | 110.1 | 3050.76 MB/s | 352 | 1 |
| path_with_backslash | 121.7 | 304.11 MB/s | 56 | 2 |
| json_in_json | 159.8 | 262.89 MB/s | 72 | 2 |
| prose_with_quotes | 93.8 | 404.90 MB/s | 64 | 2 |
| control_bytes | 117.3 | 204.61 MB/s | 56 | 2 |
| mostly_clean_one_quote | 133.1 | 2291.58 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1869.73 MB/s | 0 | 0 |
| sentence_clean | 16.7 | 2630.16 MB/s | 0 | 0 |
| url_clean | 22.4 | 2322.92 MB/s | 0 | 0 |
| log_line_clean | 44.1 | 7621.62 MB/s | 0 | 0 |
| path_with_backslash | 58.4 | 634.04 MB/s | 0 | 0 |
| json_in_json | 94.9 | 442.60 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1094.30 MB/s | 0 | 0 |
| control_bytes | 52.6 | 455.95 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.5 | 6702.30 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1019.0 | 1777.29 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1379.0 | 1313.47 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.4 | 3604.51 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9034.88 MB/s | 0 | 0 |
| url_clean | 4.9 | 10676.12 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31562.44 MB/s | 0 | 0 |
| path_escaped | 84.7 | 507.96 MB/s | 48 | 1 |
| json_in_json | 117.8 | 458.26 MB/s | 64 | 1 |
| prose_with_quotes | 71.0 | 577.60 MB/s | 48 | 1 |
| unicode_heavy | 4.2 | 7213.89 MB/s | 0 | 0 |
| unicode_escaped_dense | 336.3 | 570.90 MB/s | 192 | 1 |
| mostly_clean_one_escape | 113.4 | 2697.55 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.0 | 3192.59 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8067.81 MB/s | 0 | 0 |
| url_clean | 5.5 | 9534.66 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29597.95 MB/s | 0 | 0 |
| path_escaped | 48.4 | 888.15 MB/s | 0 | 0 |
| json_in_json | 73.0 | 739.79 MB/s | 0 | 0 |
| prose_with_quotes | 33.6 | 1221.66 MB/s | 0 | 0 |
| unicode_heavy | 4.9 | 6165.49 MB/s | 0 | 0 |
| unicode_escaped_dense | 269.4 | 712.56 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12383.45 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 86.5 | — | 0 | 0 |
| append_empty | 18.2 | — | 0 | 0 |
| replace | 59.1 | — | 0 | 0 |
| create_nested | 48.0 | — | 0 | 0 |
| overwrite_nonobject | 58.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 119.9 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 299.5 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2298.0 | 1203.56 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2258.0 | 1225.15 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 270.7 | 687.06 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4573.0 | 5247.57 MB/s | 0 | 0 |
| numberObj/goloop | 1848.0 | 5518.32 MB/s | 0 | 0 |
| nestedMixed/goloop | 2500.0 | 4321.26 MB/s | 0 | 0 |
| stringObj/neon | 2978.0 | 8057.24 MB/s | 0 | 0 |
| numberObj/neon | 1263.0 | 8076.56 MB/s | 0 | 0 |
| nestedMixed/neon | 1673.0 | 6454.50 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14447.0 | 1661.10 MB/s | 0 | 0 |
| stringObj/dispatch | 2979.0 | 8056.62 MB/s | 0 | 0 |
| numberObj/current | 6371.0 | 1600.73 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8069.87 MB/s | 0 | 0 |
| numberArr/current | 464.6 | 14208.99 MB/s | 0 | 0 |
| numberArr/dispatch | 466.1 | 14161.92 MB/s | 0 | 0 |
| nestedMixed/current | 20349.0 | 530.78 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1674.0 | 6453.32 MB/s | 0 | 0 |
