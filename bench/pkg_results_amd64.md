# lightning main-module benchmarks

- generated 2026-06-24T16:30:09Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 33.1 | 483.86 MB/s | 16 | 1 |
| sentence_clean | 44.0 | 999.31 MB/s | 48 | 1 |
| url_clean | 42.4 | 1225.71 MB/s | 64 | 1 |
| log_line_clean | 97.8 | 3436.73 MB/s | 352 | 1 |
| path_with_backslash | 200.4 | 184.64 MB/s | 184 | 4 |
| json_in_json | 246.0 | 170.74 MB/s | 216 | 4 |
| prose_with_quotes | 141.2 | 269.10 MB/s | 128 | 3 |
| control_bytes | 166.8 | 143.86 MB/s | 104 | 3 |
| mostly_clean_one_quote | 131.1 | 2325.61 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.8 | 1248.52 MB/s | 0 | 0 |
| sentence_clean | 25.3 | 1738.42 MB/s | 0 | 0 |
| url_clean | 21.9 | 2377.79 MB/s | 0 | 0 |
| log_line_clean | 30.0 | 11198.58 MB/s | 0 | 0 |
| path_with_backslash | 73.2 | 505.82 MB/s | 0 | 0 |
| json_in_json | 110.4 | 380.47 MB/s | 0 | 0 |
| prose_with_quotes | 47.2 | 804.46 MB/s | 0 | 0 |
| control_bytes | 65.3 | 367.52 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.3 | 8641.57 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 913.4 | 1982.79 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1602.0 | 1130.76 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2848.51 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7824.39 MB/s | 0 | 0 |
| url_clean | 5.6 | 9245.34 MB/s | 0 | 0 |
| log_line_clean | 9.1 | 37067.21 MB/s | 0 | 0 |
| path_escaped | 81.1 | 530.48 MB/s | 48 | 1 |
| json_in_json | 107.9 | 500.35 MB/s | 64 | 1 |
| prose_with_quotes | 65.0 | 630.47 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4369.80 MB/s | 0 | 0 |
| mostly_clean_one_escape | 87.8 | 3484.72 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.3 | 2557.98 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 7021.23 MB/s | 0 | 0 |
| url_clean | 6.3 | 8314.68 MB/s | 0 | 0 |
| log_line_clean | 9.7 | 34734.40 MB/s | 0 | 0 |
| path_escaped | 58.8 | 731.42 MB/s | 0 | 0 |
| json_in_json | 85.1 | 634.56 MB/s | 0 | 0 |
| prose_with_quotes | 43.2 | 948.00 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 4004.59 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12411.00 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2344.0 | 1179.86 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2295.0 | 1205.02 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10249.0 | 2341.48 MB/s | 0 | 0 |
| stringObj/dispatch | 4494.0 | 5340.41 MB/s | 0 | 0 |
| numberObj/current | 5374.0 | 1897.71 MB/s | 0 | 0 |
| numberObj/dispatch | 1915.0 | 5325.33 MB/s | 0 | 0 |
| numberArr/current | 236.3 | 27931.75 MB/s | 0 | 0 |
| numberArr/dispatch | 238.8 | 27642.32 MB/s | 0 | 0 |
| nestedMixed/current | 15757.0 | 685.49 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2883.0 | 3746.52 MB/s | 0 | 0 |
