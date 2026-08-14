# lightning main-module benchmarks

- generated 2026-08-14T16:22:09Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 32.0 | 499.46 MB/s | 16 | 1 |
| sentence_clean | 45.5 | 966.42 MB/s | 48 | 1 |
| url_clean | 41.9 | 1241.90 MB/s | 64 | 1 |
| log_line_clean | 96.4 | 3486.06 MB/s | 352 | 1 |
| path_with_backslash | 140.2 | 263.95 MB/s | 56 | 2 |
| json_in_json | 178.9 | 234.74 MB/s | 72 | 2 |
| prose_with_quotes | 109.3 | 347.80 MB/s | 64 | 2 |
| control_bytes | 134.1 | 178.97 MB/s | 56 | 2 |
| mostly_clean_one_quote | 105.0 | 2903.60 MB/s | 320 | 1 |
| unicode_clean | 295.2 | 799.36 MB/s | 240 | 1 |
| unicode_with_quotes | 184.4 | 341.63 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 114.8 | 2656.87 MB/s | 320 | 1 |
| invalid_utf8_dense | 703.7 | 170.53 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.5 | 1282.53 MB/s | 0 | 0 |
| sentence_clean | 24.7 | 1783.72 MB/s | 0 | 0 |
| url_clean | 21.3 | 2442.07 MB/s | 0 | 0 |
| log_line_clean | 30.9 | 10881.20 MB/s | 0 | 0 |
| path_with_backslash | 72.5 | 510.36 MB/s | 0 | 0 |
| json_in_json | 110.0 | 381.95 MB/s | 0 | 0 |
| prose_with_quotes | 47.7 | 796.22 MB/s | 0 | 0 |
| control_bytes | 66.2 | 362.43 MB/s | 0 | 0 |
| mostly_clean_one_quote | 35.6 | 8562.49 MB/s | 0 | 0 |
| unicode_clean | 234.8 | 1004.96 MB/s | 0 | 0 |
| unicode_with_quotes | 106.2 | 593.43 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 53.4 | 5712.19 MB/s | 0 | 0 |
| invalid_utf8_dense | 542.4 | 221.22 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2517.0 | 3988.90 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2411.0 | 4165.35 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2338.0 | 4294.06 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 14305.0 | 701.94 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 819.5 | 2209.79 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1512.0 | 1197.40 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2844.97 MB/s | 0 | 0 |
| sentence_clean | 5.6 | 7822.55 MB/s | 0 | 0 |
| url_clean | 5.6 | 9247.03 MB/s | 0 | 0 |
| log_line_clean | 9.1 | 37114.48 MB/s | 0 | 0 |
| path_escaped | 85.2 | 504.64 MB/s | 48 | 1 |
| json_in_json | 114.8 | 470.54 MB/s | 64 | 1 |
| prose_with_quotes | 64.3 | 638.01 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4357.91 MB/s | 0 | 0 |
| unicode_escaped_dense | 345.6 | 555.51 MB/s | 192 | 1 |
| mostly_clean_one_escape | 90.1 | 3397.20 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.2 | 2562.24 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 7036.06 MB/s | 0 | 0 |
| url_clean | 6.3 | 8318.70 MB/s | 0 | 0 |
| log_line_clean | 9.7 | 34525.97 MB/s | 0 | 0 |
| path_escaped | 58.4 | 736.61 MB/s | 0 | 0 |
| json_in_json | 86.4 | 625.07 MB/s | 0 | 0 |
| prose_with_quotes | 41.9 | 979.03 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 4003.02 MB/s | 0 | 0 |
| unicode_escaped_dense | 296.0 | 648.64 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.1 | 12692.36 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 78.2 | — | 0 | 0 |
| append_empty | 22.8 | — | 0 | 0 |
| replace | 52.4 | — | 0 | 0 |
| create_nested | 52.0 | — | 0 | 0 |
| overwrite_nonobject | 56.2 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 128.7 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 337.2 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 104.7 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 136.1 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2395.0 | 1155.10 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2306.0 | 1199.74 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 261.8 | 710.42 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 86.6 | — | 24 | 1 |
| arena | 73.8 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3826.0 | 6272.34 MB/s | 0 | 0 |
| numberObj/goloop | 1346.0 | 7578.73 MB/s | 0 | 0 |
| nestedMixed/goloop | 2315.0 | 4666.60 MB/s | 0 | 0 |
| stringObj/avx2 | 2025.0 | 11852.46 MB/s | 0 | 0 |
| numberObj/avx2 | 758.2 | 13450.12 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1338.0 | 8074.36 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10132.0 | 2368.57 MB/s | 0 | 0 |
| stringObj/dispatch | 2018.0 | 11891.88 MB/s | 0 | 0 |
| numberObj/current | 5247.0 | 1943.41 MB/s | 0 | 0 |
| numberObj/dispatch | 755.4 | 13499.70 MB/s | 0 | 0 |
| numberArr/current | 205.1 | 32176.81 MB/s | 0 | 0 |
| numberArr/dispatch | 208.1 | 31716.56 MB/s | 0 | 0 |
| nestedMixed/current | 15691.0 | 688.34 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1338.0 | 8072.45 MB/s | 0 | 0 |
