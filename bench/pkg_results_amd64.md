# lightning main-module benchmarks

- generated 2026-07-02T14:28:09Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.2 | 529.78 MB/s | 16 | 1 |
| sentence_clean | 43.9 | 1001.74 MB/s | 48 | 1 |
| url_clean | 40.9 | 1270.73 MB/s | 64 | 1 |
| log_line_clean | 94.1 | 3571.86 MB/s | 352 | 1 |
| path_with_backslash | 122.0 | 303.21 MB/s | 56 | 2 |
| json_in_json | 162.7 | 258.20 MB/s | 72 | 2 |
| prose_with_quotes | 98.9 | 384.11 MB/s | 64 | 2 |
| control_bytes | 118.5 | 202.60 MB/s | 56 | 2 |
| mostly_clean_one_quote | 104.0 | 2932.83 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.7 | 1259.91 MB/s | 0 | 0 |
| sentence_clean | 24.3 | 1810.72 MB/s | 0 | 0 |
| url_clean | 20.9 | 2490.62 MB/s | 0 | 0 |
| log_line_clean | 29.2 | 11511.05 MB/s | 0 | 0 |
| path_with_backslash | 71.2 | 519.96 MB/s | 0 | 0 |
| json_in_json | 110.4 | 380.58 MB/s | 0 | 0 |
| prose_with_quotes | 44.0 | 864.32 MB/s | 0 | 0 |
| control_bytes | 63.5 | 377.82 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.4 | 8616.76 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 794.2 | 2280.37 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1451.0 | 1247.78 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3239.56 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7347.01 MB/s | 0 | 0 |
| url_clean | 6.0 | 8666.40 MB/s | 0 | 0 |
| log_line_clean | 9.9 | 34069.67 MB/s | 0 | 0 |
| path_escaped | 76.8 | 560.01 MB/s | 48 | 1 |
| json_in_json | 104.7 | 515.59 MB/s | 64 | 1 |
| prose_with_quotes | 62.3 | 658.29 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5634.48 MB/s | 0 | 0 |
| unicode_escaped_dense | 355.6 | 539.90 MB/s | 192 | 1 |
| mostly_clean_one_escape | 87.0 | 3517.88 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3024.33 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 6929.41 MB/s | 0 | 0 |
| url_clean | 6.4 | 8151.32 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32892.48 MB/s | 0 | 0 |
| path_escaped | 57.0 | 754.10 MB/s | 0 | 0 |
| json_in_json | 81.8 | 660.27 MB/s | 0 | 0 |
| prose_with_quotes | 44.0 | 930.91 MB/s | 0 | 0 |
| unicode_heavy | 5.9 | 5124.74 MB/s | 0 | 0 |
| unicode_escaped_dense | 320.4 | 599.21 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.6 | 12433.16 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 76.8 | — | 0 | 0 |
| append_empty | 22.9 | — | 0 | 0 |
| replace | 50.9 | — | 0 | 0 |
| create_nested | 50.2 | — | 0 | 0 |
| overwrite_nonobject | 57.7 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 128.1 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 310.7 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2273.0 | 1217.02 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2209.0 | 1252.17 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4103.0 | 5848.28 MB/s | 0 | 0 |
| numberObj/goloop | 1387.0 | 7354.13 MB/s | 0 | 0 |
| nestedMixed/goloop | 1907.0 | 5664.06 MB/s | 0 | 0 |
| stringObj/avx2 | 2267.0 | 10585.84 MB/s | 0 | 0 |
| numberObj/avx2 | 814.3 | 12523.02 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1278.0 | 8451.90 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10227.0 | 2346.55 MB/s | 0 | 0 |
| stringObj/dispatch | 2290.0 | 10480.58 MB/s | 0 | 0 |
| numberObj/current | 5629.0 | 1811.70 MB/s | 0 | 0 |
| numberObj/dispatch | 816.7 | 12486.91 MB/s | 0 | 0 |
| numberArr/current | 223.1 | 29583.74 MB/s | 0 | 0 |
| numberArr/dispatch | 226.0 | 29202.85 MB/s | 0 | 0 |
| nestedMixed/current | 17574.0 | 614.58 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1278.0 | 8449.37 MB/s | 0 | 0 |
