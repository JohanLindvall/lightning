# lightning main-module benchmarks

- generated 2026-06-23T19:50:17Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 35.3 | 453.01 MB/s | 16 | 1 |
| sentence_clean | 49.7 | 885.24 MB/s | 48 | 1 |
| url_clean | 52.9 | 982.52 MB/s | 64 | 1 |
| log_line_clean | 194.9 | 1723.85 MB/s | 352 | 1 |
| path_with_backslash | 199.1 | 185.80 MB/s | 184 | 4 |
| json_in_json | 239.3 | 175.49 MB/s | 216 | 4 |
| prose_with_quotes | 142.8 | 266.13 MB/s | 128 | 3 |
| control_bytes | 163.0 | 147.21 MB/s | 104 | 3 |
| mostly_clean_one_quote | 209.4 | 1456.78 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 11.9 | 1345.09 MB/s | 0 | 0 |
| sentence_clean | 22.2 | 1984.06 MB/s | 0 | 0 |
| url_clean | 24.7 | 2105.05 MB/s | 0 | 0 |
| log_line_clean | 123.4 | 2723.16 MB/s | 0 | 0 |
| path_with_backslash | 71.2 | 519.60 MB/s | 0 | 0 |
| json_in_json | 104.6 | 401.47 MB/s | 0 | 0 |
| prose_with_quotes | 44.1 | 861.86 MB/s | 0 | 0 |
| control_bytes | 65.9 | 364.31 MB/s | 0 | 0 |
| mostly_clean_one_quote | 111.6 | 2732.26 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 879.8 | 2058.35 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1602.0 | 1130.26 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2844.37 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7819.42 MB/s | 0 | 0 |
| url_clean | 5.6 | 9243.46 MB/s | 0 | 0 |
| log_line_clean | 9.1 | 37067.63 MB/s | 0 | 0 |
| path_escaped | 81.3 | 529.01 MB/s | 48 | 1 |
| json_in_json | 108.3 | 498.53 MB/s | 64 | 1 |
| prose_with_quotes | 64.8 | 632.20 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4369.39 MB/s | 0 | 0 |
| mostly_clean_one_escape | 88.8 | 3444.57 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.3 | 2559.62 MB/s | 0 | 0 |
| sentence_clean | 6.2 | 7050.42 MB/s | 0 | 0 |
| url_clean | 6.2 | 8331.01 MB/s | 0 | 0 |
| log_line_clean | 9.7 | 34726.31 MB/s | 0 | 0 |
| path_escaped | 58.6 | 734.09 MB/s | 0 | 0 |
| json_in_json | 85.2 | 633.41 MB/s | 0 | 0 |
| prose_with_quotes | 44.0 | 932.74 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 3998.56 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12402.88 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2364.0 | 1170.13 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2302.0 | 1201.64 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10179.0 | 2357.59 MB/s | 0 | 0 |
| stringObj/dispatch | 4741.0 | 5061.44 MB/s | 0 | 0 |
| numberObj/current | 5391.0 | 1891.66 MB/s | 0 | 0 |
| numberObj/dispatch | 2019.0 | 5050.11 MB/s | 0 | 0 |
| numberArr/current | 205.2 | 32164.01 MB/s | 0 | 0 |
| numberArr/dispatch | 207.9 | 31756.62 MB/s | 0 | 0 |
| nestedMixed/current | 15731.0 | 686.59 MB/s | 0 | 0 |
| nestedMixed/dispatch | 3043.0 | 3548.91 MB/s | 0 | 0 |
