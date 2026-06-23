# lightning main-module benchmarks

- generated 2026-06-23T14:24:39Z
- go version go1.26.4 linux/arm64

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.8 | 519.17 MB/s | 16 | 1 |
| sentence_clean | 45.4 | 969.87 MB/s | 48 | 1 |
| url_clean | 42.0 | 1238.68 MB/s | 64 | 1 |
| log_line_clean | 114.8 | 2927.90 MB/s | 352 | 1 |
| path_with_backslash | 227.7 | 162.46 MB/s | 184 | 4 |
| json_in_json | 275.2 | 152.62 MB/s | 216 | 4 |
| prose_with_quotes | 150.4 | 252.68 MB/s | 128 | 3 |
| control_bytes | 157.7 | 152.18 MB/s | 104 | 3 |
| mostly_clean_one_quote | 164.2 | 1857.25 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 7.8 | 2039.60 MB/s | 0 | 0 |
| sentence_clean | 20.6 | 2140.63 MB/s | 0 | 0 |
| url_clean | 12.5 | 4168.14 MB/s | 0 | 0 |
| log_line_clean | 39.0 | 8619.07 MB/s | 0 | 0 |
| path_with_backslash | 80.2 | 461.05 MB/s | 0 | 0 |
| json_in_json | 138.7 | 302.85 MB/s | 0 | 0 |
| prose_with_quotes | 53.6 | 709.28 MB/s | 0 | 0 |
| control_bytes | 64.5 | 372.03 MB/s | 0 | 0 |
| mostly_clean_one_quote | 41.1 | 7424.51 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1031.0 | 1756.86 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1353.0 | 1338.85 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3829.01 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9006.38 MB/s | 0 | 0 |
| url_clean | 4.9 | 10698.54 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31419.57 MB/s | 0 | 0 |
| path_escaped | 87.2 | 493.16 MB/s | 48 | 1 |
| json_in_json | 119.4 | 452.24 MB/s | 64 | 1 |
| prose_with_quotes | 73.0 | 561.62 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5642.17 MB/s | 0 | 0 |
| mostly_clean_one_escape | 117.5 | 2604.36 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3288.31 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8067.00 MB/s | 0 | 0 |
| url_clean | 5.5 | 9533.73 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29596.70 MB/s | 0 | 0 |
| path_escaped | 45.9 | 937.88 MB/s | 0 | 0 |
| json_in_json | 70.6 | 764.74 MB/s | 0 | 0 |
| prose_with_quotes | 33.5 | 1222.90 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5087.48 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12351.71 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2315.0 | 1194.58 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2271.0 | 1218.17 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14398.0 | 1666.81 MB/s | 0 | 0 |
| stringObj/dispatch | 6749.0 | 3555.56 MB/s | 0 | 0 |
| numberObj/current | 6354.0 | 1605.07 MB/s | 0 | 0 |
| numberObj/dispatch | 2882.0 | 3538.98 MB/s | 0 | 0 |
| numberArr/current | 464.6 | 14209.24 MB/s | 0 | 0 |
| numberArr/dispatch | 466.7 | 14144.00 MB/s | 0 | 0 |
| nestedMixed/current | 20376.0 | 530.09 MB/s | 0 | 0 |
| nestedMixed/dispatch | 3443.0 | 3137.32 MB/s | 0 | 0 |
