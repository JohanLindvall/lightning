# lightning main-module benchmarks

- generated 2026-06-23T19:50:13Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.9 | 501.93 MB/s | 16 | 1 |
| sentence_clean | 43.1 | 1021.66 MB/s | 48 | 1 |
| url_clean | 47.1 | 1103.03 MB/s | 64 | 1 |
| log_line_clean | 159.9 | 2101.92 MB/s | 352 | 1 |
| path_with_backslash | 188.3 | 196.51 MB/s | 184 | 4 |
| json_in_json | 227.8 | 184.38 MB/s | 216 | 4 |
| prose_with_quotes | 125.2 | 303.63 MB/s | 128 | 3 |
| control_bytes | 144.8 | 165.70 MB/s | 104 | 3 |
| mostly_clean_one_quote | 200.5 | 1521.49 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.0 | 1996.64 MB/s | 0 | 0 |
| sentence_clean | 16.1 | 2733.15 MB/s | 0 | 0 |
| url_clean | 18.0 | 2893.79 MB/s | 0 | 0 |
| log_line_clean | 90.8 | 3699.86 MB/s | 0 | 0 |
| path_with_backslash | 54.3 | 681.51 MB/s | 0 | 0 |
| json_in_json | 89.7 | 468.19 MB/s | 0 | 0 |
| prose_with_quotes | 32.5 | 1169.25 MB/s | 0 | 0 |
| control_bytes | 51.1 | 469.95 MB/s | 0 | 0 |
| mostly_clean_one_quote | 86.3 | 3532.78 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1016.0 | 1783.10 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1344.0 | 1347.89 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3830.21 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 8994.96 MB/s | 0 | 0 |
| url_clean | 4.9 | 10685.46 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31423.86 MB/s | 0 | 0 |
| path_escaped | 85.7 | 502.05 MB/s | 48 | 1 |
| json_in_json | 117.2 | 460.95 MB/s | 64 | 1 |
| prose_with_quotes | 71.6 | 572.93 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5641.74 MB/s | 0 | 0 |
| mostly_clean_one_escape | 115.7 | 2644.71 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3288.99 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8067.99 MB/s | 0 | 0 |
| url_clean | 5.5 | 9534.98 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29593.22 MB/s | 0 | 0 |
| path_escaped | 45.5 | 945.33 MB/s | 0 | 0 |
| json_in_json | 70.1 | 770.06 MB/s | 0 | 0 |
| prose_with_quotes | 33.2 | 1235.07 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5088.57 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12350.35 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2297.0 | 1204.33 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2228.0 | 1241.54 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14546.0 | 1649.84 MB/s | 0 | 0 |
| stringObj/dispatch | 3872.0 | 6197.86 MB/s | 0 | 0 |
| numberObj/current | 6359.0 | 1603.59 MB/s | 0 | 0 |
| numberObj/dispatch | 1653.0 | 6170.06 MB/s | 0 | 0 |
| numberArr/current | 449.4 | 14688.29 MB/s | 0 | 0 |
| numberArr/dispatch | 451.3 | 14625.13 MB/s | 0 | 0 |
| nestedMixed/current | 19400.0 | 556.76 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2134.0 | 5061.32 MB/s | 0 | 0 |
