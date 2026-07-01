# lightning main-module benchmarks

- generated 2026-06-25T21:19:07Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.0 | 515.69 MB/s | 16 | 1 |
| sentence_clean | 45.6 | 964.33 MB/s | 48 | 1 |
| url_clean | 42.2 | 1231.76 MB/s | 64 | 1 |
| log_line_clean | 118.4 | 2836.91 MB/s | 352 | 1 |
| path_with_backslash | 204.4 | 181.05 MB/s | 184 | 4 |
| json_in_json | 232.2 | 180.89 MB/s | 216 | 4 |
| prose_with_quotes | 129.7 | 293.05 MB/s | 128 | 3 |
| control_bytes | 143.6 | 167.13 MB/s | 104 | 3 |
| mostly_clean_one_quote | 159.8 | 1909.19 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1870.48 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2607.42 MB/s | 0 | 0 |
| url_clean | 22.4 | 2315.98 MB/s | 0 | 0 |
| log_line_clean | 43.8 | 7671.52 MB/s | 0 | 0 |
| path_with_backslash | 58.4 | 634.13 MB/s | 0 | 0 |
| json_in_json | 94.1 | 446.45 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1096.24 MB/s | 0 | 0 |
| control_bytes | 52.7 | 455.27 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.5 | 6710.81 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1015.0 | 1783.84 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1347.0 | 1344.77 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3766.92 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9000.46 MB/s | 0 | 0 |
| url_clean | 4.9 | 10648.15 MB/s | 0 | 0 |
| log_line_clean | 10.8 | 31145.12 MB/s | 0 | 0 |
| path_escaped | 86.7 | 496.10 MB/s | 48 | 1 |
| json_in_json | 119.7 | 451.22 MB/s | 64 | 1 |
| prose_with_quotes | 72.5 | 565.32 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5652.45 MB/s | 0 | 0 |
| mostly_clean_one_escape | 121.8 | 2511.37 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3287.75 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8064.95 MB/s | 0 | 0 |
| url_clean | 5.5 | 9531.56 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29584.40 MB/s | 0 | 0 |
| path_escaped | 46.3 | 929.44 MB/s | 0 | 0 |
| json_in_json | 70.0 | 771.39 MB/s | 0 | 0 |
| prose_with_quotes | 34.1 | 1203.36 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5087.87 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12331.94 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2301.0 | 1202.28 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2233.0 | 1238.45 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14416.0 | 1664.72 MB/s | 0 | 0 |
| stringObj/dispatch | 3803.0 | 6310.42 MB/s | 0 | 0 |
| numberObj/current | 6341.0 | 1608.34 MB/s | 0 | 0 |
| numberObj/dispatch | 1626.0 | 6272.29 MB/s | 0 | 0 |
| numberArr/current | 464.4 | 14213.58 MB/s | 0 | 0 |
| numberArr/dispatch | 466.5 | 14151.47 MB/s | 0 | 0 |
| nestedMixed/current | 20101.0 | 537.33 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2146.0 | 5033.80 MB/s | 0 | 0 |
