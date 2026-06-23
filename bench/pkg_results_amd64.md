# lightning main-module benchmarks

- generated 2026-06-23T10:10:12Z
- go version go1.26.4 linux/amd64

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 34.0 | 470.67 MB/s | 16 | 1 |
| sentence_clean | 46.0 | 955.61 MB/s | 48 | 1 |
| url_clean | 50.9 | 1022.66 MB/s | 64 | 1 |
| log_line_clean | 187.5 | 1791.94 MB/s | 352 | 1 |
| path_with_backslash | 187.0 | 197.86 MB/s | 184 | 4 |
| json_in_json | 222.8 | 188.49 MB/s | 216 | 4 |
| prose_with_quotes | 131.2 | 289.53 MB/s | 128 | 3 |
| control_bytes | 149.2 | 160.88 MB/s | 104 | 3 |
| mostly_clean_one_quote | 202.7 | 1504.45 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 11.7 | 1370.30 MB/s | 0 | 0 |
| sentence_clean | 22.6 | 1947.60 MB/s | 0 | 0 |
| url_clean | 25.4 | 2047.25 MB/s | 0 | 0 |
| log_line_clean | 129.3 | 2599.40 MB/s | 0 | 0 |
| path_with_backslash | 70.8 | 522.68 MB/s | 0 | 0 |
| json_in_json | 101.4 | 414.14 MB/s | 0 | 0 |
| prose_with_quotes | 42.5 | 894.31 MB/s | 0 | 0 |
| control_bytes | 65.7 | 365.17 MB/s | 0 | 0 |
| mostly_clean_one_quote | 127.3 | 2395.81 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 792.6 | 2284.96 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1436.0 | 1260.80 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3237.11 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7342.53 MB/s | 0 | 0 |
| url_clean | 6.0 | 8669.91 MB/s | 0 | 0 |
| log_line_clean | 9.9 | 34068.41 MB/s | 0 | 0 |
| path_escaped | 79.2 | 543.21 MB/s | 48 | 1 |
| json_in_json | 105.0 | 514.46 MB/s | 64 | 1 |
| prose_with_quotes | 64.1 | 639.86 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5663.57 MB/s | 0 | 0 |
| mostly_clean_one_escape | 88.2 | 3469.22 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3022.49 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 6938.09 MB/s | 0 | 0 |
| url_clean | 6.3 | 8203.46 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32862.21 MB/s | 0 | 0 |
| path_escaped | 58.4 | 736.09 MB/s | 0 | 0 |
| json_in_json | 81.4 | 663.38 MB/s | 0 | 0 |
| prose_with_quotes | 45.0 | 911.13 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5161.73 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12330.41 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2280.0 | 1213.04 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2185.0 | 1265.63 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10251.0 | 2341.09 MB/s | 0 | 0 |
| stringObj/dispatch | 3886.0 | 6175.91 MB/s | 0 | 0 |
| numberObj/current | 5450.0 | 1871.06 MB/s | 0 | 0 |
| numberObj/dispatch | 1649.0 | 6185.43 MB/s | 0 | 0 |
| numberArr/current | 223.4 | 29553.49 MB/s | 0 | 0 |
| numberArr/dispatch | 226.3 | 29172.82 MB/s | 0 | 0 |
| nestedMixed/current | 15882.0 | 680.08 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2533.0 | 4263.88 MB/s | 0 | 0 |
