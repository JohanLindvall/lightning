# lightning main-module benchmarks

- generated 2026-07-02T08:46:48Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 32.3 | 494.68 MB/s | 16 | 1 |
| sentence_clean | 44.9 | 979.40 MB/s | 48 | 1 |
| url_clean | 43.6 | 1191.57 MB/s | 64 | 1 |
| log_line_clean | 97.1 | 3460.94 MB/s | 352 | 1 |
| path_with_backslash | 201.7 | 183.43 MB/s | 184 | 4 |
| json_in_json | 248.5 | 169.00 MB/s | 216 | 4 |
| prose_with_quotes | 140.2 | 271.06 MB/s | 128 | 3 |
| control_bytes | 168.6 | 142.35 MB/s | 104 | 3 |
| mostly_clean_one_quote | 123.6 | 2467.87 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.8 | 1251.50 MB/s | 0 | 0 |
| sentence_clean | 25.6 | 1722.06 MB/s | 0 | 0 |
| url_clean | 21.5 | 2413.61 MB/s | 0 | 0 |
| log_line_clean | 29.7 | 11328.69 MB/s | 0 | 0 |
| path_with_backslash | 73.1 | 505.93 MB/s | 0 | 0 |
| json_in_json | 110.4 | 380.50 MB/s | 0 | 0 |
| prose_with_quotes | 47.2 | 804.93 MB/s | 0 | 0 |
| control_bytes | 65.2 | 367.82 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.0 | 8723.14 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 901.2 | 2009.64 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1601.0 | 1131.25 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3011.97 MB/s | 0 | 0 |
| sentence_clean | 5.3 | 8285.88 MB/s | 0 | 0 |
| url_clean | 5.3 | 9776.34 MB/s | 0 | 0 |
| log_line_clean | 8.7 | 38418.69 MB/s | 0 | 0 |
| path_escaped | 85.0 | 505.83 MB/s | 48 | 1 |
| json_in_json | 115.2 | 468.90 MB/s | 64 | 1 |
| prose_with_quotes | 69.5 | 589.72 MB/s | 48 | 1 |
| unicode_heavy | 6.5 | 4580.58 MB/s | 0 | 0 |
| mostly_clean_one_escape | 89.2 | 3429.46 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.9 | 2698.85 MB/s | 0 | 0 |
| sentence_clean | 5.9 | 7412.28 MB/s | 0 | 0 |
| url_clean | 5.9 | 8752.77 MB/s | 0 | 0 |
| log_line_clean | 9.4 | 35818.85 MB/s | 0 | 0 |
| path_escaped | 62.2 | 690.90 MB/s | 0 | 0 |
| json_in_json | 90.5 | 596.86 MB/s | 0 | 0 |
| prose_with_quotes | 47.7 | 860.00 MB/s | 0 | 0 |
| unicode_heavy | 7.2 | 4143.36 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12398.30 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 83.3 | — | 0 | 0 |
| append_empty | 22.8 | — | 0 | 0 |
| replace | 56.0 | — | 0 | 0 |
| create_nested | 51.6 | — | 0 | 0 |
| overwrite_nonobject | 64.3 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 134.7 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 398.7 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2408.0 | 1148.60 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2425.0 | 1140.69 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3819.0 | 6283.40 MB/s | 0 | 0 |
| numberObj/goloop | 1346.0 | 7575.93 MB/s | 0 | 0 |
| nestedMixed/goloop | 2301.0 | 4693.62 MB/s | 0 | 0 |
| stringObj/avx2 | 2023.0 | 11864.38 MB/s | 0 | 0 |
| numberObj/avx2 | 753.3 | 13538.65 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1319.0 | 8190.24 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10151.0 | 2364.06 MB/s | 0 | 0 |
| stringObj/dispatch | 2022.0 | 11869.08 MB/s | 0 | 0 |
| numberObj/current | 5515.0 | 1849.03 MB/s | 0 | 0 |
| numberObj/dispatch | 751.9 | 13562.65 MB/s | 0 | 0 |
| numberArr/current | 204.7 | 32249.05 MB/s | 0 | 0 |
| numberArr/dispatch | 207.0 | 31886.30 MB/s | 0 | 0 |
| nestedMixed/current | 17302.0 | 624.26 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1328.0 | 8134.54 MB/s | 0 | 0 |
