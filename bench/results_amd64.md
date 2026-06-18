# JSON Deserialization Benchmarks

- generated 2026-06-18T21:37:22Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 565080 | 478.52 MB/s | 348024 | 1627 | 7.9× |
| SonicFastest | 794849 | 340.19 MB/s | 642410 | 1147 | 5.6× |
| Sonic | 798873 | 338.48 MB/s | 642635 | 1147 | 5.6× |
| Easyjson | 1768280 | 152.92 MB/s | 330273 | 749 | 2.5× |
| Goccy | 1769048 | 152.85 MB/s | 542520 | 8122 | 2.5× |
| JSONV2 | 2323934 | 116.36 MB/s | 348154 | 1628 | 1.9× |
| Stdlib | 4453897 | 60.71 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1525332 | 1132.35 MB/s | 964642 | 6177 | 11.7× |
| SonicFastest | 2089229 | 826.72 MB/s | 2704695 | 5548 | 8.6× |
| Sonic | 2113843 | 817.09 MB/s | 2705124 | 5548 | 8.5× |
| Goccy | 2487350 | 694.40 MB/s | 2590216 | 14606 | 7.2× |
| Easyjson | 3935814 | 438.84 MB/s | 972032 | 5389 | 4.5× |
| JSONV2 | 4987475 | 346.31 MB/s | 1011679 | 7594 | 3.6× |
| Stdlib | 17873080 | 96.64 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1031 | 1756.75 MB/s | 0 | 0 | 15.6× |
| Easyjson | 2991 | 605.73 MB/s | 24 | 1 | 5.4× |
| Goccy | 3182 | 569.52 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6236 | 290.56 MB/s | 3345 | 38 | 2.6× |
| Sonic | 6425 | 282.02 MB/s | 3345 | 38 | 2.5× |
| JSONV2 | 8371 | 216.45 MB/s | 640 | 6 | 1.9× |
| Stdlib | 16035 | 113.00 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1133 | 1598.86 MB/s | 0 | 0 | 14.2× |
| Easyjson | 2962 | 611.85 MB/s | 24 | 1 | 5.4× |
| Goccy | 3181 | 569.60 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 6211 | 291.75 MB/s | 3346 | 38 | 2.6× |
| Sonic | 6496 | 278.95 MB/s | 3346 | 38 | 2.5× |
| JSONV2 | 8220 | 220.44 MB/s | 640 | 6 | 2.0× |
| Stdlib | 16137 | 112.29 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 1341.03 MB/s | 144 | 10 | 12.0× |
| Goccy | 3192 | 567.70 MB/s | 2600 | 5 | 5.1× |
| Easyjson | 3371 | 537.60 MB/s | 144 | 10 | 4.8× |
| SonicFastest | 6302 | 287.54 MB/s | 3368 | 40 | 2.6× |
| Sonic | 6699 | 270.49 MB/s | 3370 | 40 | 2.4× |
| JSONV2 | 8307 | 218.13 MB/s | 632 | 7 | 1.9× |
| Stdlib | 16196 | 111.88 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 733 | 673.72 MB/s | 160 | 1 | 9.5× |
| SonicFastest | 1277 | 386.83 MB/s | 1075 | 8 | 5.4× |
| Sonic | 1278 | 386.60 MB/s | 1076 | 8 | 5.4× |
| Easyjson | 2430 | 203.31 MB/s | 448 | 3 | 2.9× |
| Goccy | 2582 | 191.34 MB/s | 856 | 23 | 2.7× |
| JSONV2 | 3451 | 143.13 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6929 | 71.29 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 482 | 477.30 MB/s | 160 | 1 | 10.3× |
| Sonic | 939 | 245.06 MB/s | 800 | 8 | 5.3× |
| SonicFastest | 943 | 243.94 MB/s | 801 | 8 | 5.2× |
| Easyjson | 1759 | 130.72 MB/s | 448 | 3 | 2.8× |
| Goccy | 1775 | 129.57 MB/s | 584 | 23 | 2.8× |
| JSONV2 | 2584 | 89.00 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4942 | 46.54 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3546317 | 547.18 MB/s | 2846867 | 2238 | 8.2× |
| Goccy | 5092556 | 381.04 MB/s | 4062390 | 13509 | 5.7× |
| Sonic | 5346810 | 362.92 MB/s | 4888353 | 1736 | 5.4× |
| SonicFastest | 5355606 | 362.33 MB/s | 4888016 | 1736 | 5.4× |
| Easyjson | 8808972 | 220.28 MB/s | 3871266 | 15043 | 3.3× |
| JSONV2 | 13495416 | 143.79 MB/s | 3237342 | 13947 | 2.2× |
| Stdlib | 29016150 | 66.88 MB/s | 3551324 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| SonicFastest | 17013313 | 476.10 MB/s | 19924043 | 41641 | 6.7× |
| Sonic | 17192206 | 471.15 MB/s | 19927427 | 41642 | 6.6× |
| Lightning | 18643668 | 434.47 MB/s | 15059838 | 51649 | 6.1× |
| Goccy | 28228629 | 286.94 MB/s | 17615926 | 107152 | 4.0× |
| Easyjson | 38406461 | 210.90 MB/s | 15059619 | 41643 | 3.0× |
| JSONV2 | 49221114 | 164.56 MB/s | 15233891 | 78973 | 2.3× |
| Stdlib | 113792875 | 71.18 MB/s | 15665073 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1444 | 1517.17 MB/s | 0 | 0 | 13.2× |
| Goccy | 3572 | 613.44 MB/s | 2864 | 4 | 5.3× |
| Easyjson | 3777 | 580.12 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6515 | 336.29 MB/s | 3600 | 38 | 2.9× |
| Sonic | 6775 | 323.38 MB/s | 3600 | 38 | 2.8× |
| JSONV2 | 8397 | 260.93 MB/s | 640 | 6 | 2.3× |
| Stdlib | 19022 | 115.18 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 586 | 33787.04 MB/s | 0 | 0 | 271.7× |
| SonicFastest | 7021 | 2818.56 MB/s | 21116 | 3 | 22.7× |
| Goccy | 23851 | 829.70 MB/s | 20492 | 2 | 6.7× |
| Sonic | 31940 | 619.56 MB/s | 20636 | 3 | 5.0× |
| JSONV2 | 34507 | 573.47 MB/s | 8 | 1 | 4.6× |
| Easyjson | 106057 | 186.59 MB/s | 0 | 0 | 1.5× |
| Stdlib | 159162 | 124.33 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2686 | 6748.21 MB/s | 864 | 4 | 44.3× |
| Easyjson | 4970 | 3646.32 MB/s | 432 | 2 | 23.9× |
| Sonic | 9197 | 1970.55 MB/s | 20385 | 5 | 12.9× |
| SonicFastest | 9277 | 1953.69 MB/s | 20361 | 5 | 12.8× |
| Goccy | 23243 | 779.75 MB/s | 19460 | 2 | 5.1× |
| JSONV2 | 52107 | 347.82 MB/s | 16500 | 50 | 2.3× |
| Stdlib | 118962 | 152.35 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2822555 | 711.59 MB/s | 2963984 | 7453 | 8.0× |
| Sonic | 3777207 | 531.74 MB/s | 5198232 | 7086 | 6.0× |
| SonicFastest | 3836901 | 523.47 MB/s | 5187218 | 7086 | 5.9× |
| Goccy | 4887963 | 410.91 MB/s | 5415289 | 15845 | 4.6× |
| Easyjson | 5623653 | 357.15 MB/s | 2981698 | 7441 | 4.0× |
| JSONV2 | 7817034 | 256.94 MB/s | 3173781 | 14563 | 2.9× |
| Stdlib | 22483585 | 89.33 MB/s | 3589320 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1060 | 517.91 MB/s | 480 | 1 | 6.4× |
| SonicFastest | 2349 | 233.74 MB/s | 2260 | 8 | 2.9× |
| Sonic | 2379 | 230.72 MB/s | 2260 | 8 | 2.9× |
| Easyjson | 2432 | 225.73 MB/s | 1616 | 5 | 2.8× |
| JSONV2 | 3376 | 162.60 MB/s | 1664 | 7 | 2.0× |
| Goccy | 3458 | 158.74 MB/s | 2128 | 43 | 2.0× |
| Stdlib | 6823 | 80.46 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 628478 | 1004.83 MB/s | 461113 | 1230 | 10.1× |
| Sonic | 1127289 | 560.21 MB/s | 1071029 | 814 | 5.6× |
| SonicFastest | 1134773 | 556.51 MB/s | 1070946 | 814 | 5.6× |
| Goccy | 1368234 | 461.55 MB/s | 992677 | 1202 | 4.6× |
| Easyjson | 1408460 | 448.37 MB/s | 422504 | 936 | 4.5× |
| JSONV2 | 2449305 | 257.83 MB/s | 571618 | 3144 | 2.6× |
| Stdlib | 6362280 | 99.26 MB/s | 654665 | 6472 | 1.0× |
