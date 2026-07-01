# lightning main-module benchmarks

- generated 2026-07-01T12:50:51Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V45 96-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 16.8 | 950.90 MB/s | 16 | 1 |
| sentence_clean | 22.3 | 1970.80 MB/s | 48 | 1 |
| url_clean | 22.2 | 2338.03 MB/s | 64 | 1 |
| log_line_clean | 55.7 | 6030.88 MB/s | 352 | 1 |
| path_with_backslash | 104.3 | 354.78 MB/s | 184 | 4 |
| json_in_json | 129.0 | 325.57 MB/s | 216 | 4 |
| prose_with_quotes | 74.2 | 511.80 MB/s | 128 | 3 |
| control_bytes | 84.0 | 285.60 MB/s | 104 | 3 |
| mostly_clean_one_quote | 68.8 | 4436.39 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.7 | 2403.98 MB/s | 0 | 0 |
| sentence_clean | 12.9 | 3399.69 MB/s | 0 | 0 |
| url_clean | 9.6 | 5422.45 MB/s | 0 | 0 |
| log_line_clean | 14.8 | 22663.94 MB/s | 0 | 0 |
| path_with_backslash | 41.7 | 887.79 MB/s | 0 | 0 |
| json_in_json | 60.8 | 690.71 MB/s | 0 | 0 |
| prose_with_quotes | 25.8 | 1470.49 MB/s | 0 | 0 |
| control_bytes | 32.2 | 744.49 MB/s | 0 | 0 |
| mostly_clean_one_quote | 16.6 | 18379.27 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 455.5 | 3975.55 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 748.1 | 2420.86 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.5 | 6288.65 MB/s | 0 | 0 |
| sentence_clean | 2.9 | 15337.21 MB/s | 0 | 0 |
| url_clean | 2.9 | 18170.98 MB/s | 0 | 0 |
| log_line_clean | 4.7 | 71171.95 MB/s | 0 | 0 |
| path_escaped | 40.5 | 1060.38 MB/s | 48 | 1 |
| json_in_json | 54.2 | 996.99 MB/s | 64 | 1 |
| prose_with_quotes | 33.3 | 1230.77 MB/s | 48 | 1 |
| unicode_heavy | 2.9 | 10200.08 MB/s | 0 | 0 |
| mostly_clean_one_escape | 50.2 | 6100.07 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 2.7 | 5973.01 MB/s | 0 | 0 |
| sentence_clean | 3.0 | 14454.24 MB/s | 0 | 0 |
| url_clean | 3.1 | 16687.54 MB/s | 0 | 0 |
| log_line_clean | 5.1 | 65989.33 MB/s | 0 | 0 |
| path_escaped | 30.1 | 1426.81 MB/s | 0 | 0 |
| json_in_json | 42.9 | 1259.62 MB/s | 0 | 0 |
| prose_with_quotes | 21.6 | 1899.29 MB/s | 0 | 0 |
| unicode_heavy | 2.9 | 10292.41 MB/s | 0 | 0 |
| mostly_clean_one_escape | 12.2 | 25032.72 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 43.9 | — | 0 | 0 |
| append_empty | 11.7 | — | 0 | 0 |
| replace | 27.4 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1180.0 | 2344.05 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1151.0 | 2403.29 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 6174.0 | 3887.18 MB/s | 0 | 0 |
| stringObj/dispatch | 2026.0 | 11842.84 MB/s | 0 | 0 |
| numberObj/current | 2803.0 | 3637.85 MB/s | 0 | 0 |
| numberObj/dispatch | 860.4 | 11852.72 MB/s | 0 | 0 |
| numberArr/current | 209.6 | 31498.57 MB/s | 0 | 0 |
| numberArr/dispatch | 210.4 | 31377.57 MB/s | 0 | 0 |
| nestedMixed/current | 8597.0 | 1256.37 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1304.0 | 8283.13 MB/s | 0 | 0 |
