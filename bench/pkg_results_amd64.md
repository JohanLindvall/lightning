# lightning main-module benchmarks

- generated 2026-06-23T14:24:39Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.3 | 511.43 MB/s | 16 | 1 |
| sentence_clean | 41.0 | 1071.98 MB/s | 48 | 1 |
| url_clean | 41.6 | 1248.87 MB/s | 64 | 1 |
| log_line_clean | 92.8 | 3621.11 MB/s | 352 | 1 |
| path_with_backslash | 177.2 | 208.78 MB/s | 184 | 4 |
| json_in_json | 226.3 | 185.60 MB/s | 216 | 4 |
| prose_with_quotes | 131.6 | 288.67 MB/s | 128 | 3 |
| control_bytes | 146.5 | 163.83 MB/s | 104 | 3 |
| mostly_clean_one_quote | 121.0 | 2521.48 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 10.9 | 1465.37 MB/s | 0 | 0 |
| sentence_clean | 17.3 | 2538.92 MB/s | 0 | 0 |
| url_clean | 13.4 | 3876.84 MB/s | 0 | 0 |
| log_line_clean | 23.2 | 14503.88 MB/s | 0 | 0 |
| path_with_backslash | 71.9 | 514.61 MB/s | 0 | 0 |
| json_in_json | 108.7 | 386.24 MB/s | 0 | 0 |
| prose_with_quotes | 45.9 | 828.54 MB/s | 0 | 0 |
| control_bytes | 64.3 | 373.05 MB/s | 0 | 0 |
| mostly_clean_one_quote | 28.9 | 10563.83 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 792.5 | 2285.32 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1442.0 | 1255.71 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3240.40 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7347.84 MB/s | 0 | 0 |
| url_clean | 6.0 | 8685.30 MB/s | 0 | 0 |
| log_line_clean | 9.9 | 34069.34 MB/s | 0 | 0 |
| path_escaped | 81.2 | 529.41 MB/s | 48 | 1 |
| json_in_json | 106.4 | 507.47 MB/s | 64 | 1 |
| prose_with_quotes | 63.9 | 642.09 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5677.74 MB/s | 0 | 0 |
| mostly_clean_one_escape | 87.1 | 3514.86 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3025.80 MB/s | 0 | 0 |
| sentence_clean | 6.5 | 6815.39 MB/s | 0 | 0 |
| url_clean | 6.4 | 8175.03 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32830.48 MB/s | 0 | 0 |
| path_escaped | 58.2 | 738.15 MB/s | 0 | 0 |
| json_in_json | 84.5 | 639.27 MB/s | 0 | 0 |
| prose_with_quotes | 45.0 | 911.78 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5156.38 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.4 | 12559.74 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2229.0 | 1241.00 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2186.0 | 1265.13 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10434.0 | 2299.96 MB/s | 0 | 0 |
| stringObj/dispatch | 4073.0 | 5892.21 MB/s | 0 | 0 |
| numberObj/current | 5685.0 | 1793.90 MB/s | 0 | 0 |
| numberObj/dispatch | 1717.0 | 5940.33 MB/s | 0 | 0 |
| numberArr/current | 223.8 | 29499.59 MB/s | 0 | 0 |
| numberArr/dispatch | 226.1 | 29195.27 MB/s | 0 | 0 |
| nestedMixed/current | 16122.0 | 669.95 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2511.0 | 4301.39 MB/s | 0 | 0 |
