# lightning main-module benchmarks

- generated 2026-08-06T12:41:37Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.2 | 529.95 MB/s | 16 | 1 |
| sentence_clean | 46.2 | 951.31 MB/s | 48 | 1 |
| url_clean | 43.2 | 1204.33 MB/s | 64 | 1 |
| log_line_clean | 120.0 | 2798.88 MB/s | 352 | 1 |
| path_with_backslash | 120.1 | 307.96 MB/s | 56 | 2 |
| json_in_json | 159.0 | 264.08 MB/s | 72 | 2 |
| prose_with_quotes | 95.5 | 397.79 MB/s | 64 | 2 |
| control_bytes | 116.3 | 206.31 MB/s | 56 | 2 |
| mostly_clean_one_quote | 140.5 | 2170.34 MB/s | 320 | 1 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 8.6 | 1868.83 MB/s | 0 | 0 |
| sentence_clean | 16.9 | 2602.34 MB/s | 0 | 0 |
| url_clean | 22.4 | 2325.07 MB/s | 0 | 0 |
| log_line_clean | 43.9 | 7657.94 MB/s | 0 | 0 |
| path_with_backslash | 58.4 | 633.32 MB/s | 0 | 0 |
| json_in_json | 94.6 | 443.86 MB/s | 0 | 0 |
| prose_with_quotes | 34.7 | 1094.52 MB/s | 0 | 0 |
| control_bytes | 52.2 | 459.98 MB/s | 0 | 0 |
| mostly_clean_one_quote | 45.6 | 6684.88 MB/s | 0 | 0 |

## GetManyPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 3070.0 | 3270.72 MB/s | 0 | 0 |

## GetPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2913.0 | 3446.99 MB/s | 0 | 0 |

## ObjectEachPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2901.0 | 3460.70 MB/s | 0 | 0 |

## StripDefaultsPretty

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 12218.0 | 821.81 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 989.0 | 1831.11 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1324.0 | 1367.59 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.4 | 3615.49 MB/s | 0 | 0 |
| sentence_clean | 4.9 | 9026.28 MB/s | 0 | 0 |
| url_clean | 4.9 | 10672.06 MB/s | 0 | 0 |
| log_line_clean | 10.7 | 31559.93 MB/s | 0 | 0 |
| path_escaped | 84.8 | 506.99 MB/s | 48 | 1 |
| json_in_json | 117.0 | 461.51 MB/s | 64 | 1 |
| prose_with_quotes | 71.2 | 576.26 MB/s | 48 | 1 |
| unicode_heavy | 4.2 | 7209.68 MB/s | 0 | 0 |
| unicode_escaped_dense | 296.5 | 647.48 MB/s | 192 | 1 |
| mostly_clean_one_escape | 121.8 | 2512.80 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.0 | 3190.80 MB/s | 0 | 0 |
| sentence_clean | 5.5 | 8066.67 MB/s | 0 | 0 |
| url_clean | 5.5 | 9533.91 MB/s | 0 | 0 |
| log_line_clean | 11.4 | 29587.52 MB/s | 0 | 0 |
| path_escaped | 44.3 | 971.63 MB/s | 0 | 0 |
| json_in_json | 67.7 | 797.91 MB/s | 0 | 0 |
| prose_with_quotes | 32.4 | 1266.99 MB/s | 0 | 0 |
| unicode_heavy | 4.9 | 6164.51 MB/s | 0 | 0 |
| unicode_escaped_dense | 221.1 | 868.27 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.8 | 12341.67 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 84.0 | — | 0 | 0 |
| append_empty | 18.4 | — | 0 | 0 |
| replace | 57.2 | — | 0 | 0 |
| create_nested | 46.2 | — | 0 | 0 |
| overwrite_nonobject | 54.6 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 116.7 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 293.9 | — | 0 | 0 |

## SetManyEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 109.8 | — | 0 | 0 |

## SetPathsEarlyExit

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 138.9 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2326.0 | 1189.03 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2222.0 | 1245.09 MB/s | 0 | 0 |

## Valid

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 271.6 | 684.74 MB/s | 0 | 0 |

## DecodeSmallSlices

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| plain | 84.8 | — | 24 | 1 |
| arena | 70.8 | — | 24 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4543.0 | 5282.53 MB/s | 0 | 0 |
| numberObj/goloop | 1840.0 | 5541.39 MB/s | 0 | 0 |
| nestedMixed/goloop | 2501.0 | 4318.12 MB/s | 0 | 0 |
| stringObj/neon | 2996.0 | 8008.77 MB/s | 0 | 0 |
| numberObj/neon | 1262.0 | 8078.37 MB/s | 0 | 0 |
| nestedMixed/neon | 1672.0 | 6461.21 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 14376.0 | 1669.32 MB/s | 0 | 0 |
| stringObj/dispatch | 2997.0 | 8007.11 MB/s | 0 | 0 |
| numberObj/current | 6327.0 | 1611.88 MB/s | 0 | 0 |
| numberObj/dispatch | 1264.0 | 8068.22 MB/s | 0 | 0 |
| numberArr/current | 449.7 | 14679.40 MB/s | 0 | 0 |
| numberArr/dispatch | 450.8 | 14641.78 MB/s | 0 | 0 |
| nestedMixed/current | 16905.0 | 638.91 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1674.0 | 6450.68 MB/s | 0 | 0 |
