# lightning main-module benchmarks

- generated 2026-08-13T07:24:08Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.0 | 532.85 MB/s | 16 | 1 |
| sentence_clean | 40.5 | 1085.07 MB/s | 48 | 1 |
| url_clean | 38.3 | 1359.16 MB/s | 64 | 1 |
| log_line_clean | 94.9 | 3539.76 MB/s | 352 | 1 |
| path_with_backslash | 124.7 | 296.83 MB/s | 56 | 2 |
| json_in_json | 158.8 | 264.48 MB/s | 72 | 2 |
| prose_with_quotes | 98.4 | 386.03 MB/s | 64 | 2 |
| control_bytes | 114.8 | 209.14 MB/s | 56 | 2 |
| mostly_clean_one_quote | 99.9 | 3052.45 MB/s | 320 | 1 |
| unicode_clean | 276.8 | 852.60 MB/s | 240 | 1 |
| unicode_with_quotes | 163.7 | 384.92 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 110.1 | 2771.38 MB/s | 320 | 1 |
| invalid_utf8_dense | 658.2 | 182.31 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.0 | 1334.84 MB/s | 0 | 0 |
| sentence_clean | 23.3 | 1889.83 MB/s | 0 | 0 |
| url_clean | 20.9 | 2488.77 MB/s | 0 | 0 |
| log_line_clean | 28.9 | 11607.70 MB/s | 0 | 0 |
| path_with_backslash | 81.7 | 452.82 MB/s | 0 | 0 |
| json_in_json | 104.3 | 402.70 MB/s | 0 | 0 |
| prose_with_quotes | 43.4 | 876.68 MB/s | 0 | 0 |
| control_bytes | 61.0 | 393.67 MB/s | 0 | 0 |
| mostly_clean_one_quote | 33.8 | 9023.09 MB/s | 0 | 0 |
| unicode_clean | 234.8 | 1005.13 MB/s | 0 | 0 |
| unicode_with_quotes | 106.0 | 594.34 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 51.8 | 5884.30 MB/s | 0 | 0 |
| invalid_utf8_dense | 542.0 | 221.40 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2444.0 | 4109.02 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2448.0 | 4101.40 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2352.0 | 4268.85 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12686.0 | 791.52 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 779.2 | 2324.32 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1365.0 | 1326.73 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3235.40 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7331.28 MB/s | 0 | 0 |
| url_clean | 6.0 | 8680.62 MB/s | 0 | 0 |
| log_line_clean | 9.9 | 33846.37 MB/s | 0 | 0 |
| path_escaped | 77.8 | 552.76 MB/s | 48 | 1 |
| json_in_json | 105.2 | 513.36 MB/s | 64 | 1 |
| prose_with_quotes | 61.4 | 668.19 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5615.31 MB/s | 0 | 0 |
| unicode_escaped_dense | 331.8 | 578.64 MB/s | 192 | 1 |
| mostly_clean_one_escape | 86.3 | 3547.24 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3029.13 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 6942.41 MB/s | 0 | 0 |
| url_clean | 7.4 | 7025.80 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32861.80 MB/s | 0 | 0 |
| path_escaped | 56.2 | 765.01 MB/s | 0 | 0 |
| json_in_json | 82.0 | 658.37 MB/s | 0 | 0 |
| prose_with_quotes | 43.4 | 945.73 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5156.16 MB/s | 0 | 0 |
| unicode_escaped_dense | 291.9 | 657.87 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.4 | 12533.97 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 76.7 | — | 0 | 0 |
| append_empty | 21.4 | — | 0 | 0 |
| replace | 50.8 | — | 0 | 0 |
| create_nested | 49.8 | — | 0 | 0 |
| overwrite_nonobject | 53.3 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 130.7 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 318.1 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 103.0 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 127.1 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2281.0 | 1212.55 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2246.0 | 1231.37 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 261.7 | 710.74 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 81.8 | — | 24 | 1 |
| arena | 71.4 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4084.0 | 5876.20 MB/s | 0 | 0 |
| numberObj/goloop | 1382.0 | 7377.06 MB/s | 0 | 0 |
| nestedMixed/goloop | 1944.0 | 5555.09 MB/s | 0 | 0 |
| stringObj/avx2 | 2284.0 | 10505.34 MB/s | 0 | 0 |
| numberObj/avx2 | 813.7 | 12532.79 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1289.0 | 8377.82 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10248.0 | 2341.62 MB/s | 0 | 0 |
| stringObj/dispatch | 2284.0 | 10505.05 MB/s | 0 | 0 |
| numberObj/current | 5529.0 | 1844.42 MB/s | 0 | 0 |
| numberObj/dispatch | 816.9 | 12484.05 MB/s | 0 | 0 |
| numberArr/current | 223.5 | 29535.32 MB/s | 0 | 0 |
| numberArr/dispatch | 226.2 | 29181.41 MB/s | 0 | 0 |
| nestedMixed/current | 16118.0 | 670.12 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1278.0 | 8449.30 MB/s | 0 | 0 |
