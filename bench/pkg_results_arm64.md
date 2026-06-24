# lightning main-module benchmarks

- generated 2026-06-24T16:30:03Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.0 | 516.49 MB/s | 16 | 1 |
| sentence_clean | 45.7 | 962.31 MB/s | 48 | 1 |
| url_clean | 42.9 | 1212.45 MB/s | 64 | 1 |
| log_line_clean | 118.3 | 2840.80 MB/s | 352 | 1 |
| path_with_backslash | 210.5 | 175.75 MB/s | 184 | 4 |
| json_in_json | 232.9 | 180.35 MB/s | 216 | 4 |
| prose_with_quotes | 129.9 | 292.62 MB/s | 128 | 3 |
| control_bytes | 144.6 | 165.97 MB/s | 104 | 3 |
| mostly_clean_one_quote | 162.2 | 1880.41 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1870.87 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2609.96 MB/s | 0 | 0 |
| url_clean | 22.4 | 2316.72 MB/s | 0 | 0 |
| log_line_clean | 43.9 | 7659.19 MB/s | 0 | 0 |
| path_with_backslash | 58.9 | 628.16 MB/s | 0 | 0 |
| json_in_json | 93.9 | 447.21 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1096.14 MB/s | 0 | 0 |
| control_bytes | 52.7 | 455.57 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.5 | 6699.20 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1018.0 | 1779.42 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1342.0 | 1349.84 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3769.21 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9005.76 MB/s | 0 | 0 |
| url_clean | 4.9 | 10654.82 MB/s | 0 | 0 |
| log_line_clean | 10.8 | 31152.09 MB/s | 0 | 0 |
| path_escaped | 88.3 | 486.95 MB/s | 48 | 1 |
| json_in_json | 121.1 | 445.84 MB/s | 64 | 1 |
| prose_with_quotes | 74.7 | 549.04 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5652.87 MB/s | 0 | 0 |
| mostly_clean_one_escape | 122.8 | 2491.15 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3287.08 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8065.37 MB/s | 0 | 0 |
| url_clean | 5.5 | 9533.13 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29594.73 MB/s | 0 | 0 |
| path_escaped | 46.1 | 933.23 MB/s | 0 | 0 |
| json_in_json | 70.7 | 763.58 MB/s | 0 | 0 |
| prose_with_quotes | 33.4 | 1228.95 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5087.94 MB/s | 0 | 0 |
| mostly_clean_one_escape | 25.0 | 12256.59 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2307.0 | 1198.87 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2244.0 | 1232.80 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14418.0 | 1664.45 MB/s | 0 | 0 |
| stringObj/dispatch | 3871.0 | 6198.74 MB/s | 0 | 0 |
| numberObj/current | 6347.0 | 1606.80 MB/s | 0 | 0 |
| numberObj/dispatch | 1652.0 | 6172.46 MB/s | 0 | 0 |
| numberArr/current | 464.6 | 14207.69 MB/s | 0 | 0 |
| numberArr/dispatch | 466.9 | 14139.33 MB/s | 0 | 0 |
| nestedMixed/current | 20087.0 | 537.71 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2132.0 | 5067.21 MB/s | 0 | 0 |
