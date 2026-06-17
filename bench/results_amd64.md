# JSON Deserialization Benchmarks

- generated 2026-06-17T10:58:29Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 501995 | 538.66 MB/s | 151230 | 495 | 5.2× |
| Sonic | 518352 | 521.66 MB/s | 644787 | 1147 | 5.1× |
| Easyjson | 1321610 | 204.60 MB/s | 330320 | 753 | 2.0× |
| JSONV2 | 1684446 | 160.53 MB/s | 348530 | 1628 | 1.6× |
| Stdlib | 2620719 | 103.18 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1448920 | 1192.06 MB/s | 979498 | 6803 | 6.4× |
| Sonic | 1791544 | 964.09 MB/s | 2733986 | 5548 | 5.2× |
| Easyjson | 3114842 | 554.51 MB/s | 986875 | 6015 | 3.0× |
| JSONV2 | 3146268 | 548.97 MB/s | 1012490 | 7594 | 2.9× |
| Stdlib | 9275107 | 186.22 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 960 | 1888.09 MB/s | 144 | 10 | 10.3× |
| Easyjson | 1865 | 971.47 MB/s | 144 | 10 | 5.3× |
| Sonic | 4897 | 370.03 MB/s | 3378 | 40 | 2.0× |
| JSONV2 | 5078 | 356.87 MB/s | 632 | 7 | 1.9× |
| Stdlib | 9859 | 183.79 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 630 | 2876.36 MB/s | 0 | 0 | 15.8× |
| Easyjson | 1664 | 1088.83 MB/s | 24 | 1 | 6.0× |
| Sonic | 4676 | 387.52 MB/s | 3367 | 38 | 2.1× |
| JSONV2 | 4946 | 366.37 MB/s | 640 | 6 | 2.0× |
| Stdlib | 9949 | 182.12 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 700 | 2587.22 MB/s | 0 | 0 | 14.1× |
| Easyjson | 1643 | 1102.64 MB/s | 24 | 1 | 6.0× |
| Sonic | 4716 | 384.24 MB/s | 3382 | 38 | 2.1× |
| JSONV2 | 5005 | 362.06 MB/s | 641 | 6 | 2.0× |
| Stdlib | 9842 | 184.11 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 343 | 669.95 MB/s | 160 | 1 | 8.7× |
| Sonic | 697 | 329.95 MB/s | 811 | 8 | 4.3× |
| Easyjson | 1092 | 210.66 MB/s | 448 | 3 | 2.7× |
| JSONV2 | 1666 | 138.03 MB/s | 528 | 7 | 1.8× |
| Stdlib | 2992 | 76.88 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 523 | 943.78 MB/s | 160 | 1 | 7.9× |
| Sonic | 892 | 553.82 MB/s | 1093 | 8 | 4.6× |
| Easyjson | 1615 | 305.97 MB/s | 448 | 3 | 2.6× |
| JSONV2 | 2120 | 232.99 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4127 | 119.70 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3458441 | 561.08 MB/s | 4282968 | 27849 | 4.9× |
| Sonic | 3530497 | 549.63 MB/s | 4882670 | 1736 | 4.8× |
| Easyjson | 6360882 | 305.06 MB/s | 4282948 | 27849 | 2.7× |
| JSONV2 | 7800061 | 248.78 MB/s | 3239349 | 13947 | 2.2× |
| Stdlib | 16883656 | 114.93 MB/s | 3551320 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 11107145 | 729.26 MB/s | 19849751 | 41641 | 5.7× |
| Lightning | 12110993 | 668.82 MB/s | 12528093 | 40027 | 5.2× |
| Easyjson | 24125252 | 335.75 MB/s | 15219646 | 61644 | 2.6× |
| JSONV2 | 32412019 | 249.91 MB/s | 15236614 | 78974 | 2.0× |
| Stdlib | 63314327 | 127.93 MB/s | 15665070 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 832 | 2634.41 MB/s | 0 | 0 | 13.2× |
| Easyjson | 2052 | 1067.82 MB/s | 24 | 1 | 5.4× |
| JSONV2 | 5073 | 431.92 MB/s | 640 | 6 | 2.2× |
| Sonic | 5147 | 425.69 MB/s | 3625 | 38 | 2.1× |
| Stdlib | 11001 | 199.17 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 516 | 38326.54 MB/s | 0 | 0 | 144.1× |
| JSONV2 | 21955 | 901.33 MB/s | 8 | 1 | 3.4× |
| Sonic | 26212 | 754.96 MB/s | 20683 | 3 | 2.8× |
| Easyjson | 46262 | 427.76 MB/s | 0 | 0 | 1.6× |
| Stdlib | 74406 | 265.96 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 5567 | 3255.45 MB/s | 18080 | 62 | 13.7× |
| Sonic | 5908 | 3067.56 MB/s | 20769 | 5 | 12.9× |
| Easyjson | 6912 | 2622.07 MB/s | 17648 | 60 | 11.0× |
| JSONV2 | 34091 | 531.64 MB/s | 16518 | 50 | 2.2× |
| Stdlib | 76005 | 238.46 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2544742 | 789.27 MB/s | 3587296 | 29254 | 5.3× |
| Sonic | 3362481 | 597.33 MB/s | 5156070 | 7085 | 4.0× |
| Easyjson | 4683073 | 428.88 MB/s | 3604108 | 29228 | 2.9× |
| JSONV2 | 5282102 | 380.25 MB/s | 3175522 | 14563 | 2.5× |
| Stdlib | 13407312 | 149.81 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 799 | 687.06 MB/s | 480 | 1 | 5.4× |
| Easyjson | 1759 | 312.05 MB/s | 1616 | 5 | 2.5× |
| Sonic | 1815 | 302.54 MB/s | 2283 | 8 | 2.4× |
| JSONV2 | 2389 | 229.81 MB/s | 1665 | 7 | 1.8× |
| Stdlib | 4351 | 126.17 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 629276 | 1003.56 MB/s | 634497 | 5544 | 6.4× |
| Sonic | 832968 | 758.15 MB/s | 1089298 | 814 | 4.8× |
| Easyjson | 1138014 | 554.93 MB/s | 591598 | 5205 | 3.5× |
| JSONV2 | 1650929 | 382.52 MB/s | 572258 | 3144 | 2.4× |
| Stdlib | 4016463 | 157.23 MB/s | 654667 | 6472 | 1.0× |
