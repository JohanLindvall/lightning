# lightning main-module benchmarks

- generated 2026-08-17T08:43:32Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 31.3 | 511.63 MB/s | 16 | 1 |
| sentence_clean | 46.0 | 956.65 MB/s | 48 | 1 |
| url_clean | 42.9 | 1211.11 MB/s | 64 | 1 |
| log_line_clean | 98.5 | 3412.52 MB/s | 352 | 1 |
| path_with_backslash | 142.4 | 259.82 MB/s | 56 | 2 |
| json_in_json | 187.8 | 223.63 MB/s | 72 | 2 |
| prose_with_quotes | 110.6 | 343.68 MB/s | 64 | 2 |
| control_bytes | 141.3 | 169.85 MB/s | 56 | 2 |
| mostly_clean_one_quote | 105.3 | 2896.67 MB/s | 320 | 1 |
| unicode_clean | 297.4 | 793.45 MB/s | 240 | 1 |
| unicode_with_quotes | 180.3 | 349.40 MB/s | 88 | 2 |
| invalid_utf8_one_byte | 115.9 | 2632.58 MB/s | 320 | 1 |
| invalid_utf8_dense | 708.2 | 169.44 MB/s | 456 | 3 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 11.6 | 1385.82 MB/s | 0 | 0 |
| sentence_clean | 21.9 | 2006.14 MB/s | 0 | 0 |
| url_clean | 21.0 | 2480.52 MB/s | 0 | 0 |
| log_line_clean | 28.8 | 11680.43 MB/s | 0 | 0 |
| path_with_backslash | 71.5 | 517.36 MB/s | 0 | 0 |
| json_in_json | 112.9 | 372.00 MB/s | 0 | 0 |
| prose_with_quotes | 46.3 | 821.07 MB/s | 0 | 0 |
| control_bytes | 68.3 | 351.20 MB/s | 0 | 0 |
| mostly_clean_one_quote | 34.4 | 8857.42 MB/s | 0 | 0 |
| unicode_clean | 234.6 | 1005.82 MB/s | 0 | 0 |
| unicode_with_quotes | 101.6 | 620.37 MB/s | 0 | 0 |
| invalid_utf8_one_byte | 52.9 | 5764.85 MB/s | 0 | 0 |
| invalid_utf8_dense | 537.3 | 223.33 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2417.0 | 4154.72 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2332.0 | 4305.34 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2260.0 | 4443.11 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 13156.0 | 763.21 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 773.1 | 2342.40 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1495.0 | 1211.55 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.6 | 2842.01 MB/s | 0 | 0 |
| sentence_clean | 5.7 | 7704.11 MB/s | 0 | 0 |
| url_clean | 5.6 | 9240.93 MB/s | 0 | 0 |
| log_line_clean | 9.2 | 36432.92 MB/s | 0 | 0 |
| path_escaped | 89.5 | 480.50 MB/s | 48 | 1 |
| json_in_json | 119.0 | 453.63 MB/s | 64 | 1 |
| prose_with_quotes | 69.4 | 590.51 MB/s | 48 | 1 |
| unicode_heavy | 6.9 | 4350.68 MB/s | 0 | 0 |
| unicode_escaped_dense | 323.6 | 593.27 MB/s | 192 | 1 |
| mostly_clean_one_escape | 91.3 | 3352.47 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 6.3 | 2557.18 MB/s | 0 | 0 |
| sentence_clean | 6.2 | 7045.29 MB/s | 0 | 0 |
| url_clean | 6.2 | 8323.52 MB/s | 0 | 0 |
| log_line_clean | 9.7 | 34674.61 MB/s | 0 | 0 |
| path_escaped | 65.5 | 656.50 MB/s | 0 | 0 |
| json_in_json | 92.3 | 585.01 MB/s | 0 | 0 |
| prose_with_quotes | 47.8 | 857.38 MB/s | 0 | 0 |
| unicode_heavy | 7.5 | 4000.26 MB/s | 0 | 0 |
| unicode_escaped_dense | 276.1 | 695.30 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.1 | 12718.35 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 74.5 | — | 0 | 0 |
| append_empty | 23.1 | — | 0 | 0 |
| replace | 50.9 | — | 0 | 0 |
| create_nested | 51.5 | — | 0 | 0 |
| overwrite_nonobject | 55.6 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 132.3 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 337.4 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 102.0 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 132.4 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2329.0 | 1187.39 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2215.0 | 1248.99 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 243.4 | 764.07 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 87.5 | — | 24 | 1 |
| arena | 74.9 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 3828.0 | 6268.99 MB/s | 0 | 0 |
| numberObj/goloop | 1340.0 | 7613.27 MB/s | 0 | 0 |
| nestedMixed/goloop | 2309.0 | 4678.54 MB/s | 0 | 0 |
| stringObj/avx2 | 2103.0 | 11411.81 MB/s | 0 | 0 |
| numberObj/avx2 | 756.3 | 13484.62 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1326.0 | 8144.74 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 7818.0 | 3069.75 MB/s | 0 | 0 |
| stringObj/dispatch | 2104.0 | 11407.46 MB/s | 0 | 0 |
| numberObj/current | 4955.0 | 2058.15 MB/s | 0 | 0 |
| numberObj/dispatch | 757.4 | 13464.30 MB/s | 0 | 0 |
| numberArr/current | 236.4 | 27918.71 MB/s | 0 | 0 |
| numberArr/dispatch | 239.5 | 27565.12 MB/s | 0 | 0 |
| nestedMixed/current | 15017.0 | 719.25 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1337.0 | 8076.57 MB/s | 0 | 0 |
