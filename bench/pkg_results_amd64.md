# lightning main-module benchmarks

- generated 2026-07-02T12:24:14Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 32.8 | 488.11 MB/s | 16 | 1 |
| sentence_clean | 44.9 | 980.42 MB/s | 48 | 1 |
| url_clean | 42.9 | 1213.44 MB/s | 64 | 1 |
| log_line_clean | 95.9 | 3504.94 MB/s | 352 | 1 |
| path_with_backslash | 138.3 | 267.54 MB/s | 56 | 2 |
| json_in_json | 178.8 | 234.95 MB/s | 72 | 2 |
| prose_with_quotes | 105.4 | 360.40 MB/s | 64 | 2 |
| control_bytes | 131.4 | 182.72 MB/s | 56 | 2 |
| mostly_clean_one_quote | 104.0 | 2931.37 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 13.4 | 1193.05 MB/s | 0 | 0 |
| sentence_clean | 25.0 | 1761.65 MB/s | 0 | 0 |
| url_clean | 21.6 | 2409.84 MB/s | 0 | 0 |
| log_line_clean | 29.7 | 11316.42 MB/s | 0 | 0 |
| path_with_backslash | 75.1 | 492.88 MB/s | 0 | 0 |
| json_in_json | 111.0 | 378.30 MB/s | 0 | 0 |
| prose_with_quotes | 47.7 | 796.80 MB/s | 0 | 0 |
| control_bytes | 65.9 | 364.08 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.4 | 8626.21 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 877.5 | 2063.81 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1589.0 | 1139.57 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.4 | 2987.87 MB/s | 0 | 0 |
| sentence_clean | 5.7 | 7780.33 MB/s | 0 | 0 |
| url_clean | 5.7 | 9204.23 MB/s | 0 | 0 |
| log_line_clean | 9.1 | 37065.30 MB/s | 0 | 0 |
| path_escaped | 78.6 | 547.37 MB/s | 48 | 1 |
| json_in_json | 107.9 | 500.49 MB/s | 64 | 1 |
| prose_with_quotes | 64.4 | 636.72 MB/s | 48 | 1 |
| unicode_heavy | 6.6 | 4565.49 MB/s | 0 | 0 |
| unicode_escaped_dense | 372.3 | 515.64 MB/s | 192 | 1 |
| mostly_clean_one_escape | 89.6 | 3415.29 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.9 | 2698.74 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7367.63 MB/s | 0 | 0 |
| url_clean | 5.9 | 8751.93 MB/s | 0 | 0 |
| log_line_clean | 9.4 | 35814.05 MB/s | 0 | 0 |
| path_escaped | 57.8 | 744.23 MB/s | 0 | 0 |
| json_in_json | 84.6 | 638.40 MB/s | 0 | 0 |
| prose_with_quotes | 42.8 | 956.95 MB/s | 0 | 0 |
| unicode_heavy | 7.2 | 4178.55 MB/s | 0 | 0 |
| unicode_escaped_dense | 327.5 | 586.24 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.1 | 12711.97 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 83.1 | — | 0 | 0 |
| append_empty | 23.3 | — | 0 | 0 |
| replace | 55.4 | — | 0 | 0 |
| create_nested | 55.4 | — | 0 | 0 |
| overwrite_nonobject | 59.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 132.5 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 316.4 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2382.0 | 1161.34 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2307.0 | 1198.92 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3832.0 | 6262.00 MB/s | 0 | 0 |
| numberObj/goloop | 1397.0 | 7299.46 MB/s | 0 | 0 |
| nestedMixed/goloop | 2356.0 | 4585.01 MB/s | 0 | 0 |
| stringObj/avx2 | 2032.0 | 11812.84 MB/s | 0 | 0 |
| numberObj/avx2 | 751.1 | 13577.74 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1343.0 | 8042.16 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10166.0 | 2360.69 MB/s | 0 | 0 |
| stringObj/dispatch | 2015.0 | 11910.86 MB/s | 0 | 0 |
| numberObj/current | 5398.0 | 1889.34 MB/s | 0 | 0 |
| numberObj/dispatch | 752.2 | 13558.38 MB/s | 0 | 0 |
| numberArr/current | 204.3 | 32311.38 MB/s | 0 | 0 |
| numberArr/dispatch | 207.1 | 31875.66 MB/s | 0 | 0 |
| nestedMixed/current | 17228.0 | 626.93 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1319.0 | 8188.09 MB/s | 0 | 0 |
