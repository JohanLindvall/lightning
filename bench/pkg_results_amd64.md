# lightning main-module benchmarks

- generated 2026-08-11T11:14:21Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 32.3 | 495.59 MB/s | 16 | 1 |
| sentence_clean | 43.7 | 1006.57 MB/s | 48 | 1 |
| url_clean | 42.4 | 1226.08 MB/s | 64 | 1 |
| log_line_clean | 100.5 | 3342.38 MB/s | 352 | 1 |
| path_with_backslash | 140.4 | 263.55 MB/s | 56 | 2 |
| json_in_json | 180.0 | 233.34 MB/s | 72 | 2 |
| prose_with_quotes | 109.9 | 345.78 MB/s | 64 | 2 |
| control_bytes | 134.4 | 178.53 MB/s | 56 | 2 |
| mostly_clean_one_quote | 105.4 | 2892.59 MB/s | 320 | 1 |
| unicode_clean | 292.4 | 807.11 MB/s | 240 | 1 |
| unicode_with_quotes | 185.0 | 340.56 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 115.5 | 2641.65 MB/s | 320 | 1 |
| invalid_utf8_dense | 705.2 | 170.16 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.5 | 1281.63 MB/s | 0 | 0 |
| sentence_clean | 24.7 | 1782.15 MB/s | 0 | 0 |
| url_clean | 21.2 | 2448.12 MB/s | 0 | 0 |
| log_line_clean | 30.9 | 10874.50 MB/s | 0 | 0 |
| path_with_backslash | 72.7 | 509.15 MB/s | 0 | 0 |
| json_in_json | 110.0 | 381.84 MB/s | 0 | 0 |
| prose_with_quotes | 47.8 | 795.59 MB/s | 0 | 0 |
| control_bytes | 66.6 | 360.37 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.7 | 8552.45 MB/s | 0 | 0 |
| unicode_clean | 234.8 | 1005.12 MB/s | 0 | 0 |
| unicode_with_quotes | 106.3 | 592.43 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 53.4 | 5707.00 MB/s | 0 | 0 |
| invalid_utf8_dense | 547.9 | 219.03 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2530.0 | 3969.06 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2409.0 | 4168.52 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2338.0 | 4295.53 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 14323.0 | 701.06 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 819.6 | 2209.72 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1509.0 | 1199.80 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2842.39 MB/s | 0 | 0 |
| sentence_clean | 5.7 | 7748.58 MB/s | 0 | 0 |
| url_clean | 5.6 | 9227.46 MB/s | 0 | 0 |
| log_line_clean | 9.1 | 37033.50 MB/s | 0 | 0 |
| path_escaped | 86.5 | 497.37 MB/s | 48 | 1 |
| json_in_json | 115.6 | 467.09 MB/s | 64 | 1 |
| prose_with_quotes | 65.7 | 624.36 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4354.70 MB/s | 0 | 0 |
| unicode_escaped_dense | 347.1 | 553.11 MB/s | 192 | 1 |
| mostly_clean_one_escape | 91.3 | 3352.82 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.3 | 2559.02 MB/s | 0 | 0 |
| sentence_clean | 6.2 | 7043.93 MB/s | 0 | 0 |
| url_clean | 6.2 | 8325.95 MB/s | 0 | 0 |
| log_line_clean | 9.8 | 34201.61 MB/s | 0 | 0 |
| path_escaped | 59.7 | 720.80 MB/s | 0 | 0 |
| json_in_json | 86.9 | 621.22 MB/s | 0 | 0 |
| prose_with_quotes | 42.2 | 970.84 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 3993.97 MB/s | 0 | 0 |
| unicode_escaped_dense | 295.9 | 648.76 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.2 | 12625.45 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 78.5 | — | 0 | 0 |
| append_empty | 22.8 | — | 0 | 0 |
| replace | 52.8 | — | 0 | 0 |
| create_nested | 49.6 | — | 0 | 0 |
| overwrite_nonobject | 56.2 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 129.0 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 337.3 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 104.6 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 138.0 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2398.0 | 1153.49 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2327.0 | 1188.84 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 266.1 | 698.95 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 87.8 | — | 24 | 1 |
| arena | 73.8 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3875.0 | 6193.62 MB/s | 0 | 0 |
| numberObj/goloop | 1345.0 | 7580.19 MB/s | 0 | 0 |
| nestedMixed/goloop | 2312.0 | 4672.51 MB/s | 0 | 0 |
| stringObj/avx2 | 2026.0 | 11845.72 MB/s | 0 | 0 |
| numberObj/avx2 | 762.2 | 13380.33 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1345.0 | 8030.33 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10163.0 | 2361.34 MB/s | 0 | 0 |
| stringObj/dispatch | 2036.0 | 11785.55 MB/s | 0 | 0 |
| numberObj/current | 5227.0 | 1951.11 MB/s | 0 | 0 |
| numberObj/dispatch | 762.6 | 13372.18 MB/s | 0 | 0 |
| numberArr/current | 206.9 | 31900.14 MB/s | 0 | 0 |
| numberArr/dispatch | 208.1 | 31718.33 MB/s | 0 | 0 |
| nestedMixed/current | 15673.0 | 689.14 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1337.0 | 8079.99 MB/s | 0 | 0 |
