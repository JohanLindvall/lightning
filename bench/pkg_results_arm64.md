# lightning main-module benchmarks

- generated 2026-06-23T19:49:57Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 32.1 | 498.70 MB/s | 16 | 1 |
| sentence_clean | 45.0 | 978.43 MB/s | 48 | 1 |
| url_clean | 49.5 | 1051.50 MB/s | 64 | 1 |
| log_line_clean | 162.7 | 2064.85 MB/s | 352 | 1 |
| path_with_backslash | 191.5 | 193.21 MB/s | 184 | 4 |
| json_in_json | 226.6 | 185.32 MB/s | 216 | 4 |
| prose_with_quotes | 126.7 | 299.98 MB/s | 128 | 3 |
| control_bytes | 144.4 | 166.23 MB/s | 104 | 3 |
| mostly_clean_one_quote | 203.7 | 1497.22 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.0 | 1996.72 MB/s | 0 | 0 |
| sentence_clean | 16.1 | 2732.85 MB/s | 0 | 0 |
| url_clean | 18.0 | 2892.63 MB/s | 0 | 0 |
| log_line_clean | 90.8 | 3702.65 MB/s | 0 | 0 |
| path_with_backslash | 54.0 | 685.04 MB/s | 0 | 0 |
| json_in_json | 89.4 | 469.86 MB/s | 0 | 0 |
| prose_with_quotes | 33.0 | 1150.39 MB/s | 0 | 0 |
| control_bytes | 50.7 | 473.61 MB/s | 0 | 0 |
| mostly_clean_one_quote | 86.5 | 3523.90 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1017.0 | 1781.33 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1345.0 | 1346.37 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3823.48 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9001.18 MB/s | 0 | 0 |
| url_clean | 4.9 | 10674.92 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31413.18 MB/s | 0 | 0 |
| path_escaped | 87.7 | 490.60 MB/s | 48 | 1 |
| json_in_json | 120.6 | 447.88 MB/s | 64 | 1 |
| prose_with_quotes | 74.2 | 552.79 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5642.59 MB/s | 0 | 0 |
| mostly_clean_one_escape | 119.3 | 2563.91 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3288.47 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8063.36 MB/s | 0 | 0 |
| url_clean | 5.5 | 9535.01 MB/s | 0 | 0 |
| log_line_clean | 11.3 | 29595.85 MB/s | 0 | 0 |
| path_escaped | 45.9 | 936.14 MB/s | 0 | 0 |
| json_in_json | 70.2 | 769.82 MB/s | 0 | 0 |
| prose_with_quotes | 33.6 | 1219.21 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5088.62 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12336.54 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2292.0 | 1206.93 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2226.0 | 1242.68 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14527.0 | 1652.01 MB/s | 0 | 0 |
| stringObj/dispatch | 6754.0 | 3553.26 MB/s | 0 | 0 |
| numberObj/current | 6359.0 | 1603.82 MB/s | 0 | 0 |
| numberObj/dispatch | 2873.0 | 3549.04 MB/s | 0 | 0 |
| numberArr/current | 449.6 | 14681.80 MB/s | 0 | 0 |
| numberArr/dispatch | 451.3 | 14626.46 MB/s | 0 | 0 |
| nestedMixed/current | 20507.0 | 526.70 MB/s | 0 | 0 |
| nestedMixed/dispatch | 3441.0 | 3139.07 MB/s | 0 | 0 |
