# lightning main-module benchmarks

- generated 2026-06-24T11:16:22Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 34.0 | 470.99 MB/s | 16 | 1 |
| sentence_clean | 42.7 | 1031.51 MB/s | 48 | 1 |
| url_clean | 41.9 | 1240.53 MB/s | 64 | 1 |
| log_line_clean | 99.3 | 3383.84 MB/s | 352 | 1 |
| path_with_backslash | 201.9 | 183.30 MB/s | 184 | 4 |
| json_in_json | 248.7 | 168.87 MB/s | 216 | 4 |
| prose_with_quotes | 137.2 | 276.92 MB/s | 128 | 3 |
| control_bytes | 166.8 | 143.93 MB/s | 104 | 3 |
| mostly_clean_one_quote | 127.7 | 2389.14 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.8 | 1248.07 MB/s | 0 | 0 |
| sentence_clean | 25.3 | 1740.66 MB/s | 0 | 0 |
| url_clean | 21.8 | 2381.69 MB/s | 0 | 0 |
| log_line_clean | 30.2 | 11120.32 MB/s | 0 | 0 |
| path_with_backslash | 77.8 | 475.52 MB/s | 0 | 0 |
| json_in_json | 110.7 | 379.35 MB/s | 0 | 0 |
| prose_with_quotes | 48.6 | 781.66 MB/s | 0 | 0 |
| control_bytes | 65.2 | 368.21 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.3 | 8649.16 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 912.7 | 1984.16 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1600.0 | 1132.04 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2840.21 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7823.14 MB/s | 0 | 0 |
| url_clean | 5.6 | 9214.12 MB/s | 0 | 0 |
| log_line_clean | 9.1 | 37111.25 MB/s | 0 | 0 |
| path_escaped | 80.3 | 535.63 MB/s | 48 | 1 |
| json_in_json | 107.2 | 503.67 MB/s | 64 | 1 |
| prose_with_quotes | 64.2 | 638.77 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4362.86 MB/s | 0 | 0 |
| mostly_clean_one_escape | 87.0 | 3518.07 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.2 | 2563.80 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 7009.99 MB/s | 0 | 0 |
| url_clean | 11.4 | 4544.63 MB/s | 0 | 0 |
| log_line_clean | 9.7 | 34710.82 MB/s | 0 | 0 |
| path_escaped | 58.6 | 734.20 MB/s | 0 | 0 |
| json_in_json | 85.4 | 632.03 MB/s | 0 | 0 |
| prose_with_quotes | 43.1 | 950.81 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 4008.67 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.7 | 12382.98 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2341.0 | 1181.57 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2295.0 | 1205.04 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10247.0 | 2342.02 MB/s | 0 | 0 |
| stringObj/dispatch | 4504.0 | 5328.46 MB/s | 0 | 0 |
| numberObj/current | 5382.0 | 1894.67 MB/s | 0 | 0 |
| numberObj/dispatch | 1914.0 | 5328.94 MB/s | 0 | 0 |
| numberArr/current | 236.2 | 27946.88 MB/s | 0 | 0 |
| numberArr/dispatch | 238.7 | 27650.27 MB/s | 0 | 0 |
| nestedMixed/current | 15672.0 | 689.20 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2835.0 | 3810.08 MB/s | 0 | 0 |
