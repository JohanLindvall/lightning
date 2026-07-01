# lightning main-module benchmarks

- generated 2026-07-01T21:37:23Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.8 | 520.38 MB/s | 16 | 1 |
| sentence_clean | 45.9 | 959.04 MB/s | 48 | 1 |
| url_clean | 43.1 | 1206.15 MB/s | 64 | 1 |
| log_line_clean | 117.8 | 2852.12 MB/s | 352 | 1 |
| path_with_backslash | 206.4 | 179.26 MB/s | 184 | 4 |
| json_in_json | 232.2 | 180.84 MB/s | 216 | 4 |
| prose_with_quotes | 129.9 | 292.55 MB/s | 128 | 3 |
| control_bytes | 144.5 | 166.07 MB/s | 104 | 3 |
| mostly_clean_one_quote | 162.4 | 1877.93 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1869.06 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2605.03 MB/s | 0 | 0 |
| url_clean | 22.6 | 2303.48 MB/s | 0 | 0 |
| log_line_clean | 43.7 | 7686.91 MB/s | 0 | 0 |
| path_with_backslash | 58.8 | 629.58 MB/s | 0 | 0 |
| json_in_json | 94.5 | 444.61 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1095.72 MB/s | 0 | 0 |
| control_bytes | 52.9 | 453.76 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.6 | 6683.59 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1017.0 | 1781.31 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1341.0 | 1350.08 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.2 | 3823.25 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 8988.36 MB/s | 0 | 0 |
| url_clean | 4.9 | 10676.44 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31413.41 MB/s | 0 | 0 |
| path_escaped | 82.3 | 522.73 MB/s | 48 | 1 |
| json_in_json | 114.8 | 470.47 MB/s | 64 | 1 |
| prose_with_quotes | 69.8 | 587.19 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5640.08 MB/s | 0 | 0 |
| mostly_clean_one_escape | 121.3 | 2521.65 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3288.38 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8066.64 MB/s | 0 | 0 |
| url_clean | 5.5 | 9532.30 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29587.76 MB/s | 0 | 0 |
| path_escaped | 46.3 | 928.73 MB/s | 0 | 0 |
| json_in_json | 70.2 | 769.32 MB/s | 0 | 0 |
| prose_with_quotes | 33.8 | 1213.50 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5087.48 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12331.73 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 87.4 | — | 0 | 0 |
| append_empty | 18.3 | — | 0 | 0 |
| replace | 59.0 | — | 0 | 0 |
| create_nested | 49.0 | — | 0 | 0 |
| overwrite_nonobject | 63.3 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 120.6 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 352.8 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2323.0 | 1190.71 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2258.0 | 1224.74 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14430.0 | 1663.02 MB/s | 0 | 0 |
| stringObj/dispatch | 3887.0 | 6173.29 MB/s | 0 | 0 |
| numberObj/current | 6345.0 | 1607.36 MB/s | 0 | 0 |
| numberObj/dispatch | 1658.0 | 6150.46 MB/s | 0 | 0 |
| numberArr/current | 464.5 | 14212.15 MB/s | 0 | 0 |
| numberArr/dispatch | 466.3 | 14155.01 MB/s | 0 | 0 |
| nestedMixed/current | 21065.0 | 512.75 MB/s | 0 | 0 |
| nestedMixed/dispatch | 2136.0 | 5057.13 MB/s | 0 | 0 |
