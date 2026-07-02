# lightning main-module benchmarks

- generated 2026-07-02T07:11:32Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

The Benchmark* functions in the lightning module itself (`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison suite in `bench/` (see `results_<arch>.md`). One table per benchmark; lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`.

## EscapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 30.8 | 519.70 MB/s | 16 | 1 |
| sentence_clean | 44.3 | 993.52 MB/s | 48 | 1 |
| url_clean | 40.9 | 1273.08 MB/s | 64 | 1 |
| log_line_clean | 95.5 | 3517.71 MB/s | 352 | 1 |
| path_with_backslash | 179.1 | 206.60 MB/s | 184 | 4 |
| json_in_json | 222.8 | 188.55 MB/s | 216 | 4 |
| prose_with_quotes | 130.6 | 291.00 MB/s | 128 | 3 |
| control_bytes | 148.0 | 162.13 MB/s | 104 | 3 |
| mostly_clean_one_quote | 120.4 | 2532.82 MB/s | 328 | 2 |

## EscapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 12.3 | 1297.09 MB/s | 0 | 0 |
| sentence_clean | 24.3 | 1809.36 MB/s | 0 | 0 |
| url_clean | 20.9 | 2484.26 MB/s | 0 | 0 |
| log_line_clean | 28.6 | 11758.14 MB/s | 0 | 0 |
| path_with_backslash | 71.8 | 515.56 MB/s | 0 | 0 |
| json_in_json | 108.9 | 385.60 MB/s | 0 | 0 |
| prose_with_quotes | 45.1 | 842.38 MB/s | 0 | 0 |
| control_bytes | 61.5 | 390.17 MB/s | 0 | 0 |
| mostly_clean_one_quote | 34.7 | 8786.70 MB/s | 0 | 0 |

## GetManyWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 843.7 | 2146.44 MB/s | 0 | 0 |

## GetPathsWithSkip

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 1447.0 | 1251.29 MB/s | 0 | 0 |

## UnescapeString

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 4.9 | 3243.21 MB/s | 0 | 0 |
| sentence_clean | 6.0 | 7348.04 MB/s | 0 | 0 |
| url_clean | 6.0 | 8675.06 MB/s | 0 | 0 |
| log_line_clean | 9.9 | 34060.10 MB/s | 0 | 0 |
| path_escaped | 84.3 | 510.14 MB/s | 48 | 1 |
| json_in_json | 110.0 | 490.96 MB/s | 64 | 1 |
| prose_with_quotes | 69.3 | 591.23 MB/s | 48 | 1 |
| unicode_heavy | 5.3 | 5677.07 MB/s | 0 | 0 |
| mostly_clean_one_escape | 89.4 | 3423.39 MB/s | 320 | 1 |

## UnescapeStringInto

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| short_clean | 5.3 | 3027.18 MB/s | 0 | 0 |
| sentence_clean | 6.3 | 6939.86 MB/s | 0 | 0 |
| url_clean | 6.3 | 8200.73 MB/s | 0 | 0 |
| log_line_clean | 10.2 | 32880.47 MB/s | 0 | 0 |
| path_escaped | 63.2 | 680.78 MB/s | 0 | 0 |
| json_in_json | 86.7 | 622.71 MB/s | 0 | 0 |
| prose_with_quotes | 49.5 | 827.88 MB/s | 0 | 0 |
| unicode_heavy | 5.8 | 5161.14 MB/s | 0 | 0 |
| mostly_clean_one_escape | 24.6 | 12418.85 MB/s | 0 | 0 |

## Set

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| append | 77.3 | — | 0 | 0 |
| append_empty | 21.9 | — | 0 | 0 |
| replace | 50.6 | — | 0 | 0 |
| create_nested | 52.9 | — | 0 | 0 |
| overwrite_nonobject | 60.4 | — | 0 | 0 |

## SetMany

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 125.7 | — | 0 | 0 |

## SetPaths

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 384.9 | — | 0 | 0 |

## StripDefaults

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2309.0 | 1197.85 MB/s | 0 | 0 |

## StripDefaultsCompact

`github.com/JohanLindvall/lightning/pkg/json`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| — | 2290.0 | 1207.90 MB/s | 0 | 0 |

## SkipBlocksVariant

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/goloop | 4048.0 | 5928.19 MB/s | 0 | 0 |
| numberObj/goloop | 1351.0 | 7549.06 MB/s | 0 | 0 |
| nestedMixed/goloop | 1897.0 | 5694.21 MB/s | 0 | 0 |
| stringObj/avx2 | 2263.0 | 10604.52 MB/s | 0 | 0 |
| numberObj/avx2 | 813.4 | 12537.71 MB/s | 0 | 0 |
| nestedMixed/avx2 | 1280.0 | 8439.09 MB/s | 0 | 0 |

## SkipContainer

`github.com/JohanLindvall/lightning/pkg/unstable`

| Case | ns/op | Throughput | B/op | allocs/op |
|---|--:|--:|--:|--:|
| stringObj/current | 10232.0 | 2345.37 MB/s | 0 | 0 |
| stringObj/dispatch | 2312.0 | 10381.66 MB/s | 0 | 0 |
| numberObj/current | 5490.0 | 1857.47 MB/s | 0 | 0 |
| numberObj/dispatch | 816.4 | 12491.18 MB/s | 0 | 0 |
| numberArr/current | 223.3 | 29566.41 MB/s | 0 | 0 |
| numberArr/dispatch | 225.6 | 29259.21 MB/s | 0 | 0 |
| nestedMixed/current | 17903.0 | 603.31 MB/s | 0 | 0 |
| nestedMixed/dispatch | 1282.0 | 8426.30 MB/s | 0 | 0 |
