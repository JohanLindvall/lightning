# lightning main-module benchmarks

- generated 2026-07-01T22:07:07Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 33.2 | 481.77 MB/s | 16 | 1 |
| sentence_clean | 45.1 | 974.69 MB/s | 48 | 1 |
| url_clean | 44.2 | 1177.33 MB/s | 64 | 1 |
| log_line_clean | 98.1 | 3423.52 MB/s | 352 | 1 |
| path_with_backslash | 208.6 | 177.37 MB/s | 184 | 4 |
| json_in_json | 253.5 | 165.69 MB/s | 216 | 4 |
| prose_with_quotes | 143.0 | 265.77 MB/s | 128 | 3 |
| control_bytes | 172.2 | 139.37 MB/s | 104 | 3 |
| mostly_clean_one_quote | 126.4 | 2412.51 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.8 | 1249.73 MB/s | 0 | 0 |
| sentence_clean | 25.4 | 1735.97 MB/s | 0 | 0 |
| url_clean | 21.6 | 2407.58 MB/s | 0 | 0 |
| log_line_clean | 29.7 | 11320.57 MB/s | 0 | 0 |
| path_with_backslash | 73.6 | 502.64 MB/s | 0 | 0 |
| json_in_json | 110.1 | 381.40 MB/s | 0 | 0 |
| prose_with_quotes | 47.3 | 804.13 MB/s | 0 | 0 |
| control_bytes | 65.4 | 367.03 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.0 | 8703.23 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 903.1 | 2005.27 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1605.0 | 1128.43 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3012.23 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8247.93 MB/s | 0 | 0 |
| url_clean | 5.3 | 9786.31 MB/s | 0 | 0 |
| log_line_clean | 8.8 | 38350.56 MB/s | 0 | 0 |
| path_escaped | 85.7 | 501.59 MB/s | 48 | 1 |
| json_in_json | 116.3 | 464.20 MB/s | 64 | 1 |
| prose_with_quotes | 70.8 | 579.46 MB/s | 48 | 1 |
| unicode_heavy | 6.6 | 4575.46 MB/s | 0 | 0 |
| mostly_clean_one_escape | 92.1 | 3321.82 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.0 | 2687.61 MB/s | 0 | 0 |
| sentence_clean | 5.9 | 7405.87 MB/s | 0 | 0 |
| url_clean | 6.0 | 8716.94 MB/s | 0 | 0 |
| log_line_clean | 9.4 | 35803.79 MB/s | 0 | 0 |
| path_escaped | 62.8 | 684.67 MB/s | 0 | 0 |
| json_in_json | 90.7 | 595.20 MB/s | 0 | 0 |
| prose_with_quotes | 47.2 | 868.54 MB/s | 0 | 0 |
| unicode_heavy | 7.2 | 4176.95 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12392.78 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 83.3 | — | 0 | 0 |
| append_empty | 22.8 | — | 0 | 0 |
| replace | 56.0 | — | 0 | 0 |
| create_nested | 51.8 | — | 0 | 0 |
| overwrite_nonobject | 64.3 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 136.7 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 393.6 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2420.0 | 1143.06 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2422.0 | 1142.14 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3822.0 | 6279.08 MB/s | 0 | 0 |
| numberObj/goloop | 1373.0 | 7427.84 MB/s | 0 | 0 |
| nestedMixed/goloop | 2304.0 | 4688.69 MB/s | 0 | 0 |
| stringObj/avx2 | 2026.0 | 11846.71 MB/s | 0 | 0 |
| numberObj/avx2 | 750.6 | 13587.12 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1332.0 | 8110.14 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10170.0 | 2359.72 MB/s | 0 | 0 |
| stringObj/dispatch | 2028.0 | 11836.10 MB/s | 0 | 0 |
| numberObj/current | 6164.0 | 1654.58 MB/s | 0 | 0 |
| numberObj/dispatch | 752.0 | 13560.47 MB/s | 0 | 0 |
| numberArr/current | 204.5 | 32281.55 MB/s | 0 | 0 |
| numberArr/dispatch | 207.3 | 31847.55 MB/s | 0 | 0 |
| nestedMixed/current | 17296.0 | 624.48 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1334.0 | 8097.79 MB/s | 0 | 0 |
