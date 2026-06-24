# lightning main-module benchmarks

- generated 2026-06-24T06:13:34Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 32.5 | 491.65 MB/s | 16 | 1 |
| sentence_clean | 40.9 | 1074.93 MB/s | 48 | 1 |
| url_clean | 41.2 | 1261.96 MB/s | 64 | 1 |
| log_line_clean | 97.9 | 3430.82 MB/s | 352 | 1 |
| path_with_backslash | 178.5 | 207.32 MB/s | 184 | 4 |
| json_in_json | 223.7 | 187.74 MB/s | 216 | 4 |
| prose_with_quotes | 131.4 | 289.26 MB/s | 128 | 3 |
| control_bytes | 145.6 | 164.82 MB/s | 104 | 3 |
| mostly_clean_one_quote | 120.3 | 2535.45 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.4 | 1294.57 MB/s | 0 | 0 |
| sentence_clean | 24.3 | 1807.36 MB/s | 0 | 0 |
| url_clean | 21.0 | 2481.21 MB/s | 0 | 0 |
| log_line_clean | 28.2 | 11902.16 MB/s | 0 | 0 |
| path_with_backslash | 72.9 | 507.49 MB/s | 0 | 0 |
| json_in_json | 109.4 | 383.98 MB/s | 0 | 0 |
| prose_with_quotes | 44.4 | 855.47 MB/s | 0 | 0 |
| control_bytes | 61.7 | 389.05 MB/s | 0 | 0 |
| mostly_clean_one_quote | 34.6 | 8822.18 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 831.6 | 2177.81 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1437.0 | 1260.36 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3239.62 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7307.34 MB/s | 0 | 0 |
| url_clean | 6.0 | 8675.32 MB/s | 0 | 0 |
| log_line_clean | 9.9 | 34010.25 MB/s | 0 | 0 |
| path_escaped | 80.3 | 535.50 MB/s | 48 | 1 |
| json_in_json | 104.9 | 514.96 MB/s | 64 | 1 |
| prose_with_quotes | 63.5 | 646.04 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5672.05 MB/s | 0 | 0 |
| mostly_clean_one_escape | 87.6 | 3494.72 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3026.84 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 6932.67 MB/s | 0 | 0 |
| url_clean | 6.3 | 8196.87 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32802.47 MB/s | 0 | 0 |
| path_escaped | 58.7 | 732.60 MB/s | 0 | 0 |
| json_in_json | 85.6 | 630.79 MB/s | 0 | 0 |
| prose_with_quotes | 45.0 | 910.25 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5158.62 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.5 | 12487.61 MB/s | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2194.0 | 1260.85 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2231.0 | 1239.62 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10438.0 | 2299.07 MB/s | 0 | 0 |
| stringObj/dispatch | 4072.0 | 5892.94 MB/s | 0 | 0 |
| numberObj/current | 5671.0 | 1798.42 MB/s | 0 | 0 |
| numberObj/dispatch | 1713.0 | 5954.24 MB/s | 0 | 0 |
| numberArr/current | 224.9 | 29351.81 MB/s | 0 | 0 |
| numberArr/dispatch | 225.3 | 29304.56 MB/s | 0 | 0 |
| nestedMixed/current | 16286.0 | 663.20 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2512.0 | 4300.47 MB/s | 0 | 0 |
