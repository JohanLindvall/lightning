# lightning main-module benchmarks

- generated 2026-07-26T09:39:04Z
- go version go1.26.5 linux/amd64
- cpu: AMD Ryzen 7 8840HS w/ Radeon 780M Graphics (16 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 20.1 | 794.45 MB/s | 16 | 1 |
| sentence_clean | 31.0 | 1418.19 MB/s | 48 | 1 |
| url_clean | 30.4 | 1713.00 MB/s | 64 | 1 |
| log_line_clean | 65.3 | 5142.61 MB/s | 352 | 1 |
| path_with_backslash | 84.0 | 440.74 MB/s | 56 | 2 |
| json_in_json | 113.4 | 370.40 MB/s | 72 | 2 |
| prose_with_quotes | 71.7 | 529.70 MB/s | 64 | 2 |
| control_bytes | 79.9 | 300.23 MB/s | 56 | 2 |
| mostly_clean_one_quote | 78.2 | 3899.13 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.7 | 2080.00 MB/s | 0 | 0 |
| sentence_clean | 15.1 | 2909.59 MB/s | 0 | 0 |
| url_clean | 13.0 | 4012.36 MB/s | 0 | 0 |
| log_line_clean | 17.8 | 18832.28 MB/s | 0 | 0 |
| path_with_backslash | 42.0 | 880.20 MB/s | 0 | 0 |
| json_in_json | 67.8 | 619.16 MB/s | 0 | 0 |
| prose_with_quotes | 27.5 | 1382.61 MB/s | 0 | 0 |
| control_bytes | 38.0 | 631.93 MB/s | 0 | 0 |
| mostly_clean_one_quote | 21.1 | 14427.38 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 508.6 | 3560.64 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 924.4 | 1959.21 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.2 | 5005.41 MB/s | 0 | 0 |
| sentence_clean | 3.7 | 11778.38 MB/s | 0 | 0 |
| url_clean | 3.7 | 14045.77 MB/s | 0 | 0 |
| log_line_clean | 6.1 | 55137.46 MB/s | 0 | 0 |
| path_escaped | 57.2 | 752.01 MB/s | 48 | 1 |
| json_in_json | 75.0 | 720.46 MB/s | 64 | 1 |
| prose_with_quotes | 47.4 | 865.85 MB/s | 48 | 1 |
| unicode_heavy | 3.4 | 8855.23 MB/s | 0 | 0 |
| unicode_escaped_dense | 268.1 | 716.21 MB/s | 192 | 1 |
| mostly_clean_one_escape | 63.1 | 4847.43 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.4 | 4724.61 MB/s | 0 | 0 |
| sentence_clean | 4.0 | 10919.70 MB/s | 0 | 0 |
| url_clean | 4.0 | 12945.46 MB/s | 0 | 0 |
| log_line_clean | 6.4 | 52323.53 MB/s | 0 | 0 |
| path_escaped | 37.5 | 1145.25 MB/s | 0 | 0 |
| json_in_json | 52.8 | 1022.59 MB/s | 0 | 0 |
| prose_with_quotes | 28.6 | 1435.01 MB/s | 0 | 0 |
| unicode_heavy | 3.7 | 8159.35 MB/s | 0 | 0 |
| unicode_escaped_dense | 215.4 | 891.37 MB/s | 0 | 0 |
| mostly_clean_one_escape | 14.9 | 20495.70 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 48.2 | — | 0 | 0 |
| append_empty | 14.1 | — | 0 | 0 |
| replace | 32.5 | — | 0 | 0 |
| create_nested | 31.3 | — | 0 | 0 |
| overwrite_nonobject | 35.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 77.7 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 197.1 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1400.0 | 1976.02 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1468.0 | 1884.10 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 166.2 | 1118.90 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 2544.0 | 9434.68 MB/s | 0 | 0 |
| numberObj/goloop | 865.4 | 11784.60 MB/s | 0 | 0 |
| nestedMixed/goloop | 1189.0 | 9081.56 MB/s | 0 | 0 |
| stringObj/avx2 | 1373.0 | 17476.37 MB/s | 0 | 0 |
| numberObj/avx2 | 492.3 | 20716.85 MB/s | 0 | 0 |
| nestedMixed/avx2 | 798.1 | 13533.90 MB/s | 0 | 0 |
| stringObj/avx512 | 940.9 | 25504.27 MB/s | 0 | 0 |
| numberObj/avx512 | 340.9 | 29916.64 MB/s | 0 | 0 |
| nestedMixed/avx512 | 614.3 | 17583.12 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 6267.0 | 3829.25 MB/s | 0 | 0 |
| stringObj/dispatch | 924.6 | 25954.35 MB/s | 0 | 0 |
| numberObj/current | 3683.0 | 2769.01 MB/s | 0 | 0 |
| numberObj/dispatch | 348.1 | 29299.32 MB/s | 0 | 0 |
| numberArr/current | 140.7 | 46906.09 MB/s | 0 | 0 |
| numberArr/dispatch | 143.2 | 46098.18 MB/s | 0 | 0 |
| nestedMixed/current | 10885.0 | 992.30 MB/s | 0 | 0 |
| nestedMixed/dispatch | 615.1 | 17559.45 MB/s | 0 | 0 |
