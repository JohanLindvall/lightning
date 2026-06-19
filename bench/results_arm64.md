# JSON Deserialization Benchmarks

- generated 2026-06-19T22:02:31Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 484737 | 557.83 MB/s | 348024 | 1627 | 6.9× |
| SonicFastest | 634989 | 425.84 MB/s | 498654 | 968 | 5.3× |
| Sonic | 635485 | 425.51 MB/s | 498217 | 968 | 5.3× |
| Goccy | 1384070 | 195.37 MB/s | 540802 | 8121 | 2.4× |
| Easyjson | 1417001 | 190.83 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2079991 | 130.00 MB/s | 348145 | 1628 | 1.6× |
| Stdlib | 3352144 | 80.67 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1466711 | 1177.60 MB/s | 964643 | 6177 | 9.2× |
| SonicFastest | 2054590 | 840.66 MB/s | 2939384 | 4020 | 6.5× |
| Sonic | 2060820 | 838.11 MB/s | 2944724 | 4020 | 6.5× |
| Goccy | 2424381 | 712.43 MB/s | 2587822 | 14607 | 5.5× |
| Easyjson | 4289475 | 402.66 MB/s | 972033 | 5389 | 3.1× |
| JSONV2 | 4323609 | 399.48 MB/s | 1011700 | 7594 | 3.1× |
| Stdlib | 13443916 | 128.47 MB/s | 1234450 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1243 | 1457.54 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2549 | 710.99 MB/s | 24 | 1 | 5.5× |
| Goccy | 2772 | 653.70 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5877 | 308.33 MB/s | 3742 | 40 | 2.4× |
| Sonic | 5881 | 308.13 MB/s | 3736 | 40 | 2.4× |
| JSONV2 | 7739 | 234.15 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14060 | 128.87 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1253 | 1446.19 MB/s | 0 | 0 | 11.2× |
| Easyjson | 2559 | 707.98 MB/s | 24 | 1 | 5.5× |
| Goccy | 2762 | 656.06 MB/s | 2608 | 4 | 5.1× |
| Sonic | 5883 | 307.99 MB/s | 3712 | 40 | 2.4× |
| SonicFastest | 5892 | 307.53 MB/s | 3719 | 40 | 2.4× |
| JSONV2 | 7787 | 232.70 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14089 | 128.61 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1442 | 1256.34 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2787 | 650.07 MB/s | 144 | 10 | 5.0× |
| Goccy | 2825 | 641.34 MB/s | 2600 | 5 | 5.0× |
| SonicFastest | 6023 | 300.87 MB/s | 3746 | 42 | 2.3× |
| Sonic | 6032 | 300.41 MB/s | 3750 | 42 | 2.3× |
| JSONV2 | 8012 | 226.17 MB/s | 632 | 7 | 1.7× |
| Stdlib | 14014 | 129.30 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 772 | 639.81 MB/s | 160 | 1 | 7.2× |
| SonicFastest | 1232 | 401.14 MB/s | 983 | 6 | 4.5× |
| Sonic | 1233 | 400.51 MB/s | 972 | 6 | 4.5× |
| Easyjson | 2216 | 222.92 MB/s | 448 | 3 | 2.5× |
| Goccy | 2451 | 201.57 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3208 | 154.01 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5541 | 89.15 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 463 | 496.29 MB/s | 160 | 1 | 9.0× |
| Sonic | 894 | 257.15 MB/s | 684 | 6 | 4.7× |
| SonicFastest | 901 | 255.31 MB/s | 687 | 6 | 4.6× |
| Easyjson | 1412 | 162.84 MB/s | 448 | 3 | 3.0× |
| Goccy | 1625 | 141.57 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2341 | 98.24 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4180 | 55.03 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2879426 | 673.91 MB/s | 2846868 | 2238 | 8.2× |
| Sonic | 4509291 | 430.33 MB/s | 14611215 | 1407 | 5.2× |
| SonicFastest | 4732183 | 410.06 MB/s | 14608054 | 1407 | 5.0× |
| Goccy | 4833158 | 401.49 MB/s | 4065269 | 13510 | 4.9× |
| Easyjson | 7648499 | 253.71 MB/s | 3871268 | 15043 | 3.1× |
| JSONV2 | 11223418 | 172.89 MB/s | 3237327 | 13947 | 2.1× |
| Stdlib | 23541849 | 82.43 MB/s | 3551323 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13798374 | 587.03 MB/s | 15059835 | 51649 | 6.5× |
| Sonic | 17114597 | 473.28 MB/s | 72373982 | 40015 | 5.3× |
| SonicFastest | 17190535 | 471.19 MB/s | 72301855 | 40015 | 5.2× |
| Goccy | 24158189 | 335.29 MB/s | 17414240 | 107152 | 3.7× |
| Easyjson | 31044278 | 260.92 MB/s | 15059618 | 41643 | 2.9× |
| JSONV2 | 44154404 | 183.45 MB/s | 15233878 | 78973 | 2.0× |
| Stdlib | 89985060 | 90.02 MB/s | 15665080 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1550 | 1413.42 MB/s | 0 | 0 | 10.3× |
| Goccy | 3148 | 695.90 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3202 | 684.27 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6321 | 346.60 MB/s | 3962 | 40 | 2.5× |
| Sonic | 6337 | 345.72 MB/s | 3966 | 40 | 2.5× |
| JSONV2 | 8004 | 273.75 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15963 | 137.25 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2062 | 9596.60 MB/s | 0 | 0 | 54.2× |
| Goccy | 20146 | 982.30 MB/s | 20491 | 2 | 5.5× |
| Sonic | 27734 | 713.52 MB/s | 21909 | 4 | 4.0× |
| SonicFastest | 27794 | 711.99 MB/s | 21951 | 4 | 4.0× |
| JSONV2 | 29604 | 668.46 MB/s | 8 | 1 | 3.8× |
| Easyjson | 86031 | 230.02 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111745 | 177.09 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2972 | 6097.38 MB/s | 864 | 4 | 34.7× |
| Easyjson | 3970 | 4565.23 MB/s | 432 | 2 | 26.0× |
| Sonic | 9790 | 1851.31 MB/s | 23084 | 6 | 10.5× |
| SonicFastest | 9810 | 1847.57 MB/s | 23331 | 6 | 10.5× |
| Goccy | 15884 | 1140.99 MB/s | 19459 | 2 | 6.5× |
| JSONV2 | 45770 | 395.98 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103182 | 175.65 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2505799 | 801.54 MB/s | 2963984 | 7453 | 7.5× |
| SonicFastest | 4019487 | 499.69 MB/s | 10035032 | 13682 | 4.7× |
| Sonic | 4042418 | 496.85 MB/s | 10028649 | 13682 | 4.6× |
| Goccy | 4206283 | 477.50 MB/s | 5416808 | 15845 | 4.5× |
| Easyjson | 5060045 | 396.93 MB/s | 2981730 | 7441 | 3.7× |
| JSONV2 | 6967300 | 288.27 MB/s | 3173803 | 14563 | 2.7× |
| Stdlib | 18786656 | 106.91 MB/s | 3589321 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 902 | 608.57 MB/s | 480 | 1 | 6.2× |
| Easyjson | 2118 | 259.24 MB/s | 1616 | 5 | 2.7× |
| Sonic | 2681 | 204.76 MB/s | 1936 | 26 | 2.1× |
| SonicFastest | 2689 | 204.14 MB/s | 1952 | 26 | 2.1× |
| Goccy | 3069 | 178.90 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3310 | 165.88 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5616 | 97.75 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 572462 | 1103.15 MB/s | 461113 | 1230 | 9.4× |
| Sonic | 1019636 | 619.35 MB/s | 1050705 | 1102 | 5.3× |
| SonicFastest | 1021670 | 618.12 MB/s | 1060195 | 1102 | 5.3× |
| Easyjson | 1125459 | 561.12 MB/s | 422505 | 936 | 4.8× |
| Goccy | 1185748 | 532.59 MB/s | 992768 | 1203 | 4.5× |
| JSONV2 | 2172997 | 290.62 MB/s | 571663 | 3144 | 2.5× |
| Stdlib | 5364766 | 117.72 MB/s | 654666 | 6472 | 1.0× |
