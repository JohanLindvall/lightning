# lightning main-module benchmarks

- generated 2026-09-02T12:27:20Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 24.0 | 667.29 MB/s | 16 | 1 |
| sentence_clean | 32.0 | 1374.98 MB/s | 48 | 1 |
| url_clean | 31.4 | 1658.12 MB/s | 64 | 1 |
| log_line_clean | 73.8 | 4550.98 MB/s | 352 | 1 |
| path_with_backslash | 91.9 | 402.81 MB/s | 56 | 2 |
| json_in_json | 120.4 | 348.79 MB/s | 72 | 2 |
| prose_with_quotes | 75.3 | 504.77 MB/s | 64 | 2 |
| control_bytes | 92.7 | 258.96 MB/s | 56 | 2 |
| mostly_clean_one_quote | 79.9 | 3815.49 MB/s | 320 | 1 |
| unicode_clean | 215.2 | 1096.89 MB/s | 240 | 1 |
| unicode_with_quotes | 122.0 | 516.22 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 87.7 | 3477.13 MB/s | 320 | 1 |
| invalid_utf8_dense | 516.7 | 232.26 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.8 | 1828.14 MB/s | 0 | 0 |
| sentence_clean | 16.7 | 2636.34 MB/s | 0 | 0 |
| url_clean | 15.5 | 3358.55 MB/s | 0 | 0 |
| log_line_clean | 22.4 | 14992.32 MB/s | 0 | 0 |
| path_with_backslash | 49.6 | 745.99 MB/s | 0 | 0 |
| json_in_json | 76.6 | 548.57 MB/s | 0 | 0 |
| prose_with_quotes | 31.4 | 1208.21 MB/s | 0 | 0 |
| control_bytes | 49.8 | 482.05 MB/s | 0 | 0 |
| mostly_clean_one_quote | 25.9 | 11785.58 MB/s | 0 | 0 |
| unicode_clean | 181.5 | 1300.50 MB/s | 0 | 0 |
| unicode_with_quotes | 75.6 | 832.91 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 40.7 | 7487.90 MB/s | 0 | 0 |
| invalid_utf8_dense | 429.4 | 279.43 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1636.0 | 6138.54 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1625.0 | 6180.49 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1543.0 | 6508.45 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 8808.0 | 1139.95 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 523.6 | 3458.49 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1009.0 | 1795.31 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 3.6 | 4499.46 MB/s | 0 | 0 |
| sentence_clean | 4.4 | 10063.49 MB/s | 0 | 0 |
| url_clean | 4.4 | 11892.21 MB/s | 0 | 0 |
| log_line_clean | 7.4 | 45332.62 MB/s | 0 | 0 |
| path_escaped | 57.9 | 742.92 MB/s | 48 | 1 |
| json_in_json | 78.2 | 690.41 MB/s | 64 | 1 |
| prose_with_quotes | 47.2 | 867.78 MB/s | 48 | 1 |
| unicode_heavy | 4.0 | 7560.46 MB/s | 0 | 0 |
| unicode_escaped_dense | 248.8 | 771.79 MB/s | 192 | 1 |
| mostly_clean_one_escape | 67.5 | 4529.69 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.1 | 3905.86 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 8944.12 MB/s | 0 | 0 |
| url_clean | 4.9 | 10576.55 MB/s | 0 | 0 |
| log_line_clean | 7.9 | 42429.91 MB/s | 0 | 0 |
| path_escaped | 41.5 | 1035.54 MB/s | 0 | 0 |
| json_in_json | 60.4 | 894.72 MB/s | 0 | 0 |
| prose_with_quotes | 31.9 | 1285.16 MB/s | 0 | 0 |
| unicode_heavy | 4.4 | 6858.67 MB/s | 0 | 0 |
| unicode_escaped_dense | 218.3 | 879.62 MB/s | 0 | 0 |
| mostly_clean_one_escape | 18.4 | 16596.20 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 55.1 | — | 0 | 0 |
| append_empty | 17.0 | — | 0 | 0 |
| replace | 35.5 | — | 0 | 0 |
| create_nested | 39.0 | — | 0 | 0 |
| overwrite_nonobject | 41.8 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 101.2 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 242.0 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 75.2 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 99.9 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1742.0 | 1587.49 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1691.0 | 1635.48 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 179.7 | 1035.29 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 67.6 | — | 24 | 1 |
| arena | 57.1 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3140.0 | 7643.46 MB/s | 0 | 0 |
| numberObj/goloop | 1079.0 | 9447.42 MB/s | 0 | 0 |
| nestedMixed/goloop | 1480.0 | 7299.76 MB/s | 0 | 0 |
| stringObj/avx2 | 1751.0 | 13703.97 MB/s | 0 | 0 |
| numberObj/avx2 | 633.4 | 16100.73 MB/s | 0 | 0 |
| nestedMixed/avx2 | 970.4 | 11130.52 MB/s | 0 | 0 |
| stringObj/avx512 | 1163.0 | 20629.67 MB/s | 0 | 0 |
| numberObj/avx512 | 433.0 | 23551.96 MB/s | 0 | 0 |
| nestedMixed/avx512 | 809.2 | 13347.28 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 6319.0 | 3797.48 MB/s | 0 | 0 |
| stringObj/dispatch | 1166.0 | 20585.56 MB/s | 0 | 0 |
| numberObj/current | 4326.0 | 2357.61 MB/s | 0 | 0 |
| numberObj/dispatch | 432.3 | 23588.18 MB/s | 0 | 0 |
| numberArr/current | 173.5 | 38055.30 MB/s | 0 | 0 |
| numberArr/dispatch | 175.7 | 37569.75 MB/s | 0 | 0 |
| nestedMixed/current | 13373.0 | 807.66 MB/s | 0 | 0 |
| nestedMixed/dispatch | 809.5 | 13342.52 MB/s | 0 | 0 |
