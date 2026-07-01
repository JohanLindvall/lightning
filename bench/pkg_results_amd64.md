# lightning main-module benchmarks

- generated 2026-07-01T21:37:18Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.3 | 511.43 MB/s | 16 | 1 |
| sentence_clean | 41.7 | 1055.59 MB/s | 48 | 1 |
| url_clean | 40.8 | 1275.44 MB/s | 64 | 1 |
| log_line_clean | 95.4 | 3522.88 MB/s | 352 | 1 |
| path_with_backslash | 178.1 | 207.72 MB/s | 184 | 4 |
| json_in_json | 223.4 | 187.98 MB/s | 216 | 4 |
| prose_with_quotes | 131.2 | 289.72 MB/s | 128 | 3 |
| control_bytes | 146.6 | 163.71 MB/s | 104 | 3 |
| mostly_clean_one_quote | 120.7 | 2527.29 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.3 | 1296.68 MB/s | 0 | 0 |
| sentence_clean | 24.6 | 1791.68 MB/s | 0 | 0 |
| url_clean | 21.0 | 2474.05 MB/s | 0 | 0 |
| log_line_clean | 28.4 | 11810.39 MB/s | 0 | 0 |
| path_with_backslash | 67.5 | 548.02 MB/s | 0 | 0 |
| json_in_json | 109.7 | 382.90 MB/s | 0 | 0 |
| prose_with_quotes | 44.4 | 855.06 MB/s | 0 | 0 |
| control_bytes | 61.5 | 390.22 MB/s | 0 | 0 |
| mostly_clean_one_quote | 34.3 | 8898.07 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 817.4 | 2215.54 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1436.0 | 1260.76 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3245.00 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7343.99 MB/s | 0 | 0 |
| url_clean | 6.0 | 8650.48 MB/s | 0 | 0 |
| log_line_clean | 9.9 | 34055.27 MB/s | 0 | 0 |
| path_escaped | 80.2 | 536.21 MB/s | 48 | 1 |
| json_in_json | 104.9 | 514.68 MB/s | 64 | 1 |
| prose_with_quotes | 63.8 | 642.25 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5667.82 MB/s | 0 | 0 |
| mostly_clean_one_escape | 87.3 | 3503.05 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3017.51 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 6931.53 MB/s | 0 | 0 |
| url_clean | 6.4 | 8181.55 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32873.79 MB/s | 0 | 0 |
| path_escaped | 58.1 | 739.45 MB/s | 0 | 0 |
| json_in_json | 82.2 | 657.26 MB/s | 0 | 0 |
| prose_with_quotes | 45.1 | 908.44 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5135.32 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.4 | 12563.89 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 78.2 | — | 0 | 0 |
| append_empty | 21.8 | — | 0 | 0 |
| replace | 50.5 | — | 0 | 0 |
| create_nested | 53.4 | — | 0 | 0 |
| overwrite_nonobject | 61.0 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 127.2 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 390.8 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2230.0 | 1240.16 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2230.0 | 1240.57 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10418.0 | 2303.62 MB/s | 0 | 0 |
| stringObj/dispatch | 4137.0 | 5801.35 MB/s | 0 | 0 |
| numberObj/current | 5645.0 | 1806.56 MB/s | 0 | 0 |
| numberObj/dispatch | 1727.0 | 5904.60 MB/s | 0 | 0 |
| numberArr/current | 223.5 | 29533.33 MB/s | 0 | 0 |
| numberArr/dispatch | 226.5 | 29142.64 MB/s | 0 | 0 |
| nestedMixed/current | 16067.0 | 672.23 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2526.0 | 4275.65 MB/s | 0 | 0 |
