# lightning main-module benchmarks

- generated 2026-06-23T19:49:53Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 33.6 | 476.44 MB/s | 16 | 1 |
| sentence_clean | 47.4 | 928.61 MB/s | 48 | 1 |
| url_clean | 50.5 | 1029.87 MB/s | 64 | 1 |
| log_line_clean | 186.5 | 1801.70 MB/s | 352 | 1 |
| path_with_backslash | 188.5 | 196.25 MB/s | 184 | 4 |
| json_in_json | 223.3 | 188.11 MB/s | 216 | 4 |
| prose_with_quotes | 130.9 | 290.22 MB/s | 128 | 3 |
| control_bytes | 149.2 | 160.83 MB/s | 104 | 3 |
| mostly_clean_one_quote | 198.8 | 1534.11 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 11.7 | 1370.22 MB/s | 0 | 0 |
| sentence_clean | 22.6 | 1948.28 MB/s | 0 | 0 |
| url_clean | 25.5 | 2041.75 MB/s | 0 | 0 |
| log_line_clean | 129.5 | 2594.20 MB/s | 0 | 0 |
| path_with_backslash | 70.8 | 522.65 MB/s | 0 | 0 |
| json_in_json | 101.5 | 413.80 MB/s | 0 | 0 |
| prose_with_quotes | 42.8 | 888.85 MB/s | 0 | 0 |
| control_bytes | 64.8 | 370.53 MB/s | 0 | 0 |
| mostly_clean_one_quote | 122.0 | 2499.61 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 791.5 | 2287.93 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1437.0 | 1260.26 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3239.46 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7330.70 MB/s | 0 | 0 |
| url_clean | 6.0 | 8678.46 MB/s | 0 | 0 |
| log_line_clean | 9.9 | 34018.23 MB/s | 0 | 0 |
| path_escaped | 79.0 | 543.93 MB/s | 48 | 1 |
| json_in_json | 104.5 | 516.59 MB/s | 64 | 1 |
| prose_with_quotes | 63.1 | 649.88 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5651.75 MB/s | 0 | 0 |
| mostly_clean_one_escape | 85.9 | 3561.21 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3024.43 MB/s | 0 | 0 |
| sentence_clean | 6.4 | 6926.96 MB/s | 0 | 0 |
| url_clean | 6.3 | 8200.29 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32858.87 MB/s | 0 | 0 |
| path_escaped | 60.5 | 710.62 MB/s | 0 | 0 |
| json_in_json | 81.4 | 663.51 MB/s | 0 | 0 |
| prose_with_quotes | 45.2 | 907.84 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5161.32 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.4 | 12553.51 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2260.0 | 1223.98 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2193.0 | 1261.38 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10240.0 | 2343.44 MB/s | 0 | 0 |
| stringObj/dispatch | 3890.0 | 6169.48 MB/s | 0 | 0 |
| numberObj/current | 5506.0 | 1852.23 MB/s | 0 | 0 |
| numberObj/dispatch | 1643.0 | 6205.85 MB/s | 0 | 0 |
| numberArr/current | 223.0 | 29598.94 MB/s | 0 | 0 |
| numberArr/dispatch | 226.4 | 29157.10 MB/s | 0 | 0 |
| nestedMixed/current | 15781.0 | 684.42 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2559.0 | 4220.84 MB/s | 0 | 0 |
