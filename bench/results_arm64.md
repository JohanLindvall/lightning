# JSON Deserialization Benchmarks

- generated 2026-06-17T17:44:42Z
- go version go1.26.4 darwin/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 457111 | 591.55 MB/s | 466557 | 968 | 5.3× |
| Lightning | 521544 | 518.47 MB/s | 151224 | 495 | 4.6× |
| Easyjson | 1176828 | 229.77 MB/s | 330321 | 753 | 2.0× |
| JSONV2 | 1357814 | 199.15 MB/s | 348212 | 1628 | 1.8× |
| Stdlib | 2411478 | 112.13 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 1287955 | 1341.04 MB/s | 3571147 | 4021 | 7.6× |
| Lightning | 1308567 | 1319.92 MB/s | 979490 | 6803 | 7.5× |
| JSONV2 | 2492595 | 692.93 MB/s | 1012000 | 7594 | 3.9× |
| Easyjson | 2494818 | 692.32 MB/s | 986873 | 6015 | 3.9× |
| Stdlib | 9838474 | 175.56 MB/s | 1234450 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1036 | 1749.69 MB/s | 0 | 0 | 9.9× |
| Easyjson | 1795 | 1009.37 MB/s | 24 | 1 | 5.7× |
| Sonic | 3816 | 474.84 MB/s | 3822 | 40 | 2.7× |
| JSONV2 | 4762 | 380.54 MB/s | 640 | 6 | 2.1× |
| Stdlib | 10227 | 177.17 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1033 | 1754.61 MB/s | 0 | 0 | 9.9× |
| Easyjson | 1821 | 995.11 MB/s | 24 | 1 | 5.6× |
| Sonic | 3888 | 466.06 MB/s | 3807 | 40 | 2.6× |
| JSONV2 | 4794 | 377.99 MB/s | 640 | 6 | 2.1× |
| Stdlib | 10242 | 176.92 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1086 | 1668.67 MB/s | 144 | 10 | 9.4× |
| Easyjson | 1914 | 946.89 MB/s | 144 | 10 | 5.3× |
| Sonic | 3906 | 463.92 MB/s | 3810 | 42 | 2.6× |
| JSONV2 | 4940 | 366.79 MB/s | 632 | 7 | 2.1× |
| Stdlib | 10204 | 177.58 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 580 | 851.71 MB/s | 160 | 1 | 6.6× |
| Sonic | 735 | 671.95 MB/s | 977 | 6 | 5.2× |
| Easyjson | 1695 | 291.46 MB/s | 448 | 3 | 2.3× |
| JSONV2 | 2038 | 242.36 MB/s | 528 | 7 | 1.9× |
| Stdlib | 3845 | 128.47 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 344 | 669.44 MB/s | 160 | 1 | 7.8× |
| Sonic | 514 | 447.38 MB/s | 666 | 6 | 5.2× |
| Easyjson | 1011 | 227.59 MB/s | 448 | 3 | 2.6× |
| JSONV2 | 1482 | 155.17 MB/s | 528 | 7 | 1.8× |
| Stdlib | 2679 | 85.86 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3021721 | 642.17 MB/s | 4282957 | 27849 | 5.2× |
| Sonic | 3055739 | 635.03 MB/s | 14608361 | 1407 | 5.1× |
| Easyjson | 5198464 | 373.28 MB/s | 4282953 | 27849 | 3.0× |
| JSONV2 | 6505242 | 298.29 MB/s | 3237700 | 13947 | 2.4× |
| Stdlib | 15654506 | 123.96 MB/s | 3551324 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 11105424 | 729.38 MB/s | 12528028 | 40027 | 5.6× |
| Sonic | 11422821 | 709.11 MB/s | 72279828 | 40015 | 5.4× |
| Easyjson | 21641281 | 374.29 MB/s | 15219647 | 61644 | 2.9× |
| JSONV2 | 27034733 | 299.62 MB/s | 15234523 | 78973 | 2.3× |
| Stdlib | 61759270 | 131.16 MB/s | 15665091 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1125 | 1946.90 MB/s | 0 | 0 | 10.6× |
| Easyjson | 2296 | 954.31 MB/s | 24 | 1 | 5.2× |
| Sonic | 4207 | 520.83 MB/s | 4015 | 40 | 2.8× |
| JSONV2 | 4995 | 438.63 MB/s | 640 | 6 | 2.4× |
| Stdlib | 11946 | 183.42 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1021 | 19390.95 MB/s | 0 | 0 | 83.0× |
| Sonic | 16560 | 1194.98 MB/s | 22748 | 4 | 5.1× |
| JSONV2 | 22407 | 883.15 MB/s | 8 | 1 | 3.8× |
| Easyjson | 70250 | 281.69 MB/s | 0 | 0 | 1.2× |
| Stdlib | 84790 | 233.39 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3702 | 4895.86 MB/s | 18080 | 62 | 20.9× |
| Easyjson | 4707 | 3850.15 MB/s | 17648 | 60 | 16.4× |
| Sonic | 5396 | 3358.57 MB/s | 23120 | 6 | 14.3× |
| JSONV2 | 29811 | 607.97 MB/s | 16502 | 50 | 2.6× |
| Stdlib | 77224 | 234.69 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1807263 | 1111.35 MB/s | 3587264 | 29254 | 7.1× |
| Sonic | 2898392 | 692.97 MB/s | 10046079 | 13682 | 4.4× |
| Easyjson | 3520445 | 570.52 MB/s | 3602902 | 29230 | 3.7× |
| JSONV2 | 4628363 | 433.95 MB/s | 3174327 | 14563 | 2.8× |
| Stdlib | 12876488 | 155.98 MB/s | 3589326 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 717 | 765.76 MB/s | 480 | 1 | 5.1× |
| Easyjson | 1320 | 416.04 MB/s | 1616 | 5 | 2.8× |
| Sonic | 1624 | 338.01 MB/s | 1974 | 26 | 2.3× |
| JSONV2 | 1852 | 296.39 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 3664 | 149.82 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 431867 | 1462.29 MB/s | 634493 | 5544 | 8.9× |
| Sonic | 600481 | 1051.68 MB/s | 1097677 | 1102 | 6.4× |
| Easyjson | 845174 | 747.20 MB/s | 591598 | 5205 | 4.6× |
| JSONV2 | 1395918 | 452.40 MB/s | 571832 | 3144 | 2.8× |
| Stdlib | 3859340 | 163.63 MB/s | 654668 | 6472 | 1.0× |
