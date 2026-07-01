# JSON Deserialization Benchmarks

- generated 2026-07-01T12:54:36Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104636 | 1216.36 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 104696 | 1215.67 MB/s | 49280 | 2 | 10.6× |
| Sonic | 181205 | 702.38 MB/s | 194424 | 10 | 6.1× |
| SonicFastest | 182445 | 697.61 MB/s | 195657 | 10 | 6.1× |
| Goccy | 189695 | 670.95 MB/s | 224736 | 884 | 5.8× |
| Easyjson | 212210 | 599.76 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 418940 | 303.80 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 447645 | 211.45 MB/s | 465585 | 9714 | 2.5× |
| Stdlib | 1105127 | 115.17 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3748508 | 600.52 MB/s | 3122874 | 3081 | 7.0× |
| Lightning | 3829838 | 587.77 MB/s | 3122875 | 3081 | 6.8× |
| Sonic | 4484037 | 502.01 MB/s | 15255276 | 970 | 5.8× |
| SonicFastest | 4528912 | 497.04 MB/s | 15233755 | 970 | 5.8× |
| Goccy | 10329058 | 217.93 MB/s | 4118117 | 56532 | 2.5× |
| Easyjson | 11108688 | 202.64 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11357916 | 198.19 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16395039 | 137.30 MB/s | 3123239 | 3083 | 1.6× |
| Stdlib | 26111645 | 86.21 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 490959 | 550.76 MB/s | 348024 | 1627 | 6.8× |
| LightningDestructive | 494383 | 546.95 MB/s | 348024 | 1627 | 6.8× |
| SonicFastest | 627536 | 430.90 MB/s | 472833 | 968 | 5.3× |
| Sonic | 634519 | 426.15 MB/s | 479371 | 968 | 5.3× |
| Goccy | 1408864 | 191.93 MB/s | 543303 | 8122 | 2.4× |
| Easyjson | 1446596 | 186.92 MB/s | 330272 | 749 | 2.3× |
| LightningDecodeAny | 1610922 | 167.86 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2112600 | 128.00 MB/s | 348149 | 1628 | 1.6× |
| Stdlib | 3355499 | 80.59 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1352580 | 1276.97 MB/s | 959848 | 5881 | 9.8× |
| Lightning | 1377822 | 1253.58 MB/s | 959890 | 5882 | 9.6× |
| Sonic | 2044279 | 844.90 MB/s | 2721498 | 4020 | 6.5× |
| SonicFastest | 2051319 | 842.00 MB/s | 2732474 | 4020 | 6.5× |
| Goccy | 2331026 | 740.96 MB/s | 2583803 | 14605 | 5.7× |
| Easyjson | 4221396 | 409.15 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4282034 | 403.36 MB/s | 1011636 | 7594 | 3.1× |
| LightningDecodeAny | 4668636 | 107.16 MB/s | 4976205 | 81466 | 2.8× |
| Stdlib | 13235875 | 130.49 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1190 | 1522.31 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1205 | 1503.55 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2559 | 708.09 MB/s | 24 | 1 | 5.5× |
| Goccy | 2792 | 648.96 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 5894 | 307.44 MB/s | 3717 | 40 | 2.4× |
| Sonic | 5928 | 305.64 MB/s | 3755 | 40 | 2.4× |
| JSONV2 | 7858 | 230.61 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8157 | 222.01 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14042 | 129.04 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1219 | 1486.03 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1234 | 1468.30 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2547 | 711.46 MB/s | 24 | 1 | 5.5× |
| Goccy | 2907 | 623.30 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 5979 | 303.08 MB/s | 3873 | 40 | 2.4× |
| Sonic | 5983 | 302.84 MB/s | 3862 | 40 | 2.4× |
| JSONV2 | 7864 | 230.40 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8501 | 213.03 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14068 | 128.80 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1413 | 1282.64 MB/s | 144 | 10 | 9.9× |
| LightningDestructive | 1456 | 1244.92 MB/s | 144 | 10 | 9.6× |
| Easyjson | 2760 | 656.54 MB/s | 144 | 10 | 5.1× |
| Goccy | 2902 | 624.37 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6120 | 296.07 MB/s | 3897 | 42 | 2.3× |
| SonicFastest | 6126 | 295.77 MB/s | 3927 | 42 | 2.3× |
| JSONV2 | 8003 | 226.42 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8524 | 212.46 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14028 | 129.17 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 744 | 664.47 MB/s | 160 | 1 | 7.4× |
| Lightning | 745 | 662.90 MB/s | 160 | 1 | 7.4× |
| SonicFastest | 1229 | 401.83 MB/s | 992 | 6 | 4.5× |
| Sonic | 1231 | 401.39 MB/s | 986 | 6 | 4.5× |
| LightningDecodeAny | 1652 | 298.40 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2200 | 224.51 MB/s | 448 | 3 | 2.5× |
| Goccy | 2461 | 200.69 MB/s | 856 | 23 | 2.2× |
| JSONV2 | 3183 | 155.20 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5523 | 89.44 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 471 | 488.70 MB/s | 160 | 1 | 8.8× |
| LightningDestructive | 472 | 487.38 MB/s | 160 | 1 | 8.8× |
| SonicFastest | 877 | 262.22 MB/s | 650 | 6 | 4.7× |
| Sonic | 879 | 261.66 MB/s | 651 | 6 | 4.7× |
| LightningDecodeAny | 1380 | 165.95 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1385 | 166.04 MB/s | 448 | 3 | 3.0× |
| Goccy | 1624 | 141.60 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2355 | 97.65 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4130 | 55.69 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 74073 | 879.29 MB/s | 172432 | 107 | 7.4× |
| LightningDestructive | 76038 | 856.57 MB/s | 166212 | 102 | 7.2× |
| Sonic | 96745 | 673.23 MB/s | 156567 | 75 | 5.7× |
| SonicFastest | 104671 | 622.26 MB/s | 163247 | 75 | 5.2× |
| Goccy | 153978 | 423.00 MB/s | 228746 | 134 | 3.6× |
| LightningDecodeAny | 187756 | 284.03 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 232221 | 280.47 MB/s | 206654 | 607 | 2.4× |
| Stdlib | 549372 | 118.56 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2615337 | 741.96 MB/s | 2846864 | 2238 | 9.0× |
| Lightning | 2662096 | 728.93 MB/s | 2846867 | 2238 | 8.8× |
| Goccy | 4754782 | 408.11 MB/s | 4064379 | 13510 | 4.9× |
| SonicFastest | 5017930 | 386.71 MB/s | 14608665 | 1407 | 4.7× |
| Sonic | 5021510 | 386.43 MB/s | 14606973 | 1407 | 4.7× |
| Easyjson | 7541567 | 257.30 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9662847 | 200.82 MB/s | 7013524 | 219937 | 2.4× |
| JSONV2 | 11351797 | 170.94 MB/s | 3237215 | 13947 | 2.1× |
| Stdlib | 23470045 | 82.68 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1105918 | 3009.11 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1823634 | 1824.83 MB/s | 2488905 | 2995 | 11.4× |
| SonicFastest | 2687851 | 1238.10 MB/s | 6420242 | 4248 | 7.7× |
| Sonic | 2704976 | 1230.26 MB/s | 6470947 | 4248 | 7.7× |
| LightningDecodeAny | 3746342 | 820.47 MB/s | 4886621 | 56892 | 5.6× |
| Goccy | 4505790 | 738.57 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7334143 | 453.75 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20820178 | 159.84 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222491 | 990.36 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 225046 | 979.11 MB/s | 136897 | 228 | 9.1× |
| SonicFastest | 378323 | 582.43 MB/s | 293833 | 398 | 5.4× |
| Sonic | 378990 | 581.40 MB/s | 296713 | 398 | 5.4× |
| Goccy | 432473 | 509.50 MB/s | 364141 | 1067 | 4.7× |
| Easyjson | 545405 | 404.00 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 732951 | 300.63 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 889070 | 121.83 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2040160 | 108.00 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12814963 | 632.08 MB/s | 15059833 | 51649 | 7.0× |
| Lightning | 13265376 | 610.62 MB/s | 15059835 | 51649 | 6.7× |
| Sonic | 16718969 | 484.48 MB/s | 70901923 | 40014 | 5.4× |
| SonicFastest | 16803230 | 482.05 MB/s | 70930577 | 40014 | 5.3× |
| Goccy | 23921148 | 338.61 MB/s | 17400477 | 107150 | 3.7× |
| Easyjson | 31157870 | 259.97 MB/s | 15059617 | 41643 | 2.9× |
| LightningDecodeAny | 40640422 | 128.03 MB/s | 34333350 | 912810 | 2.2× |
| JSONV2 | 44090850 | 183.71 MB/s | 15233746 | 78972 | 2.0× |
| Stdlib | 89477607 | 90.53 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6068284 | 491.65 MB/s | 3985336 | 30015 | 7.8× |
| Lightning | 6168246 | 483.68 MB/s | 3985336 | 30015 | 7.7× |
| Sonic | 8660224 | 344.50 MB/s | 26533424 | 56760 | 5.5× |
| SonicFastest | 8660572 | 344.49 MB/s | 26540260 | 56760 | 5.5× |
| Goccy | 16497113 | 180.85 MB/s | 10665605 | 273651 | 2.9× |
| Easyjson | 16497621 | 180.84 MB/s | 9479440 | 30115 | 2.9× |
| LightningDecodeAny | 18524280 | 99.02 MB/s | 20023837 | 408557 | 2.6× |
| JSONV2 | 24549668 | 121.53 MB/s | 9257144 | 86278 | 1.9× |
| Stdlib | 47269688 | 63.12 MB/s | 9258095 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1386871 | 521.75 MB/s | 907388 | 3618 | 8.3× |
| LightningDestructive | 1389373 | 520.81 MB/s | 907393 | 3618 | 8.3× |
| Sonic | 1777366 | 407.12 MB/s | 3189974 | 7226 | 6.5× |
| SonicFastest | 1778576 | 406.84 MB/s | 3201794 | 7226 | 6.5× |
| Easyjson | 4245211 | 170.45 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4395907 | 148.00 MB/s | 5752203 | 80172 | 2.6× |
| Goccy | 4711514 | 153.58 MB/s | 2757439 | 80271 | 2.5× |
| JSONV2 | 5965584 | 121.30 MB/s | 2704631 | 7318 | 1.9× |
| Stdlib | 11575860 | 62.51 MB/s | 2704553 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1941298 | 812.53 MB/s | 907393 | 3618 | 8.1× |
| Lightning | 1972289 | 799.76 MB/s | 907385 | 3618 | 8.0× |
| SonicFastest | 2243404 | 703.11 MB/s | 5788244 | 7226 | 7.0× |
| Sonic | 2278454 | 692.29 MB/s | 5798021 | 7226 | 6.9× |
| LightningDecodeAny | 3876438 | 194.35 MB/s | 5752200 | 80172 | 4.1× |
| Easyjson | 5594690 | 281.94 MB/s | 2847908 | 3698 | 2.8× |
| Goccy | 5611694 | 281.08 MB/s | 3583370 | 80268 | 2.8× |
| JSONV2 | 6670603 | 236.46 MB/s | 2704595 | 7318 | 2.4× |
| Stdlib | 15718440 | 100.35 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 223463 | 671.81 MB/s | 81920 | 1 | 8.2× |
| LightningDestructive | 224689 | 668.14 MB/s | 81920 | 1 | 8.1× |
| SonicFastest | 270171 | 555.66 MB/s | 241332 | 6 | 6.7× |
| Sonic | 272727 | 550.46 MB/s | 248168 | 6 | 6.7× |
| LightningDecodeAny | 474529 | 316.36 MB/s | 746004 | 10020 | 3.8× |
| Goccy | 852035 | 176.19 MB/s | 324109 | 10004 | 2.1× |
| JSONV2 | 1105378 | 135.81 MB/s | 357716 | 20 | 1.6× |
| Stdlib | 1823283 | 82.34 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32822 | 856.64 MB/s | 30272 | 105 | 9.3× |
| LightningDestructive | 32926 | 853.94 MB/s | 30144 | 103 | 9.2× |
| SonicFastest | 64000 | 439.33 MB/s | 47253 | 103 | 4.7× |
| Sonic | 64061 | 438.91 MB/s | 47459 | 103 | 4.7× |
| Easyjson | 68070 | 413.06 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71242 | 394.67 MB/s | 59222 | 188 | 4.3× |
| JSONV2 | 134528 | 209.01 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 150838 | 186.41 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303736 | 92.57 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1951 | 1193.29 MB/s | 32 | 1 | 11.7× |
| LightningDestructive | 2004 | 1161.73 MB/s | 32 | 1 | 11.4× |
| Goccy | 4125 | 564.38 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4232 | 550.11 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5030 | 462.81 MB/s | 4238 | 6 | 4.5× |
| Sonic | 5108 | 455.78 MB/s | 4331 | 6 | 4.5× |
| JSONV2 | 8403 | 277.06 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10237 | 164.60 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22763 | 102.27 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 862.79 MB/s | 0 | 0 | 11.3× |
| LightningDestructive | 221 | 855.70 MB/s | 0 | 0 | 11.2× |
| Goccy | 379 | 498.67 MB/s | 304 | 2 | 6.5× |
| Easyjson | 482 | 392.09 MB/s | 0 | 0 | 5.1× |
| SonicFastest | 798 | 236.81 MB/s | 498 | 4 | 3.1× |
| Sonic | 800 | 236.21 MB/s | 499 | 4 | 3.1× |
| JSONV2 | 1028 | 183.84 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1231 | 108.90 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2467 | 76.63 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1541 | 1421.95 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1558 | 1405.95 MB/s | 0 | 0 | 10.3× |
| Goccy | 3166 | 692.06 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3184 | 688.21 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6311 | 347.18 MB/s | 3937 | 40 | 2.5× |
| Sonic | 6371 | 343.88 MB/s | 4015 | 40 | 2.5× |
| JSONV2 | 8088 | 270.89 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8184 | 221.29 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 16008 | 136.87 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 702494 | 726.66 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 723868 | 705.21 MB/s | 703779 | 1012 | 8.4× |
| Goccy | 1137265 | 448.86 MB/s | 1137616 | 5006 | 5.3× |
| Sonic | 1150727 | 443.61 MB/s | 888459 | 2006 | 5.3× |
| SonicFastest | 1155714 | 441.70 MB/s | 905038 | 2006 | 5.2× |
| Easyjson | 1536954 | 332.13 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3190196 | 160.01 MB/s | 1076010 | 12646 | 1.9× |
| LightningDecodeAny | 3525349 | 130.90 MB/s | 2785926 | 66022 | 1.7× |
| Stdlib | 6050342 | 84.37 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14645.92 MB/s | 0 | 0 | 82.7× |
| LightningDestructive | 1391 | 14225.02 MB/s | 0 | 0 | 80.3× |
| Goccy | 19876 | 995.64 MB/s | 20491 | 2 | 5.6× |
| SonicFastest | 27194 | 727.69 MB/s | 22513 | 4 | 4.1× |
| Sonic | 27329 | 724.10 MB/s | 22810 | 4 | 4.1× |
| JSONV2 | 29728 | 665.68 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 79140 | 250.04 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86019 | 230.05 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111683 | 177.19 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2648 | 6843.31 MB/s | 0 | 0 | 38.9× |
| Lightning | 2830 | 6405.28 MB/s | 432 | 2 | 36.4× |
| Easyjson | 3962 | 4574.65 MB/s | 432 | 2 | 26.0× |
| Sonic | 9928 | 1825.51 MB/s | 22996 | 6 | 10.4× |
| SonicFastest | 10015 | 1809.63 MB/s | 23175 | 6 | 10.3× |
| Goccy | 16011 | 1131.97 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 16659 | 1073.42 MB/s | 29088 | 191 | 6.2× |
| JSONV2 | 46249 | 391.88 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102907 | 176.12 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2264053 | 887.12 MB/s | 2960388 | 7411 | 8.3× |
| Lightning | 2298665 | 873.77 MB/s | 2962101 | 7417 | 8.1× |
| Goccy | 4294475 | 467.69 MB/s | 5411948 | 15830 | 4.4× |
| Sonic | 4487715 | 447.55 MB/s | 10976719 | 13683 | 4.2× |
| SonicFastest | 4498541 | 446.48 MB/s | 10913376 | 13683 | 4.2× |
| Easyjson | 4953998 | 405.43 MB/s | 2981482 | 7438 | 3.8× |
| JSONV2 | 6887893 | 291.60 MB/s | 3173684 | 14563 | 2.7× |
| LightningDecodeAny | 7193628 | 158.79 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18704893 | 107.38 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 888 | 618.45 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 892 | 615.59 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1908 | 287.27 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2141 | 256.44 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2677 | 205.08 MB/s | 1961 | 26 | 2.1× |
| Sonic | 2680 | 204.83 MB/s | 1938 | 26 | 2.1× |
| Goccy | 3035 | 180.86 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3261 | 168.35 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5619 | 97.71 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 493213 | 1280.41 MB/s | 364473 | 566 | 11.0× |
| Lightning | 553394 | 1141.17 MB/s | 413002 | 878 | 9.8× |
| SonicFastest | 1026179 | 615.40 MB/s | 992686 | 1102 | 5.3× |
| Sonic | 1026201 | 615.39 MB/s | 989748 | 1102 | 5.3× |
| Easyjson | 1143230 | 552.39 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1161849 | 543.54 MB/s | 987904 | 1201 | 4.7× |
| JSONV2 | 2163867 | 291.85 MB/s | 571618 | 3144 | 2.5× |
| LightningDecodeAny | 2385436 | 195.73 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5403858 | 116.86 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 710363 | 791.72 MB/s | 544251 | 448 | 7.4× |
| Lightning | 886690 | 634.28 MB/s | 767622 | 1254 | 5.9× |
| SonicFastest | 1015471 | 553.84 MB/s | 945735 | 1476 | 5.2× |
| Sonic | 1020020 | 551.37 MB/s | 947624 | 1476 | 5.2× |
| Goccy | 1329665 | 422.97 MB/s | 1040733 | 1030 | 4.0× |
| Easyjson | 1736736 | 323.83 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2709334 | 207.58 MB/s | 2114151 | 30295 | 1.9× |
| JSONV2 | 2741869 | 205.12 MB/s | 927450 | 3482 | 1.9× |
| Stdlib | 5268999 | 106.74 MB/s | 1011670 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 650024 | 820.24 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 670132 | 795.63 MB/s | 368224 | 2293 | 8.2× |
| Easyjson | 1106439 | 481.89 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1143067 | 466.44 MB/s | 1036770 | 4351 | 4.8× |
| SonicFastest | 1146827 | 464.92 MB/s | 1041177 | 4351 | 4.8× |
| Goccy | 1285288 | 414.83 MB/s | 1167202 | 5409 | 4.3× |
| JSONV2 | 2523201 | 211.31 MB/s | 745449 | 13288 | 2.2× |
| LightningDecodeAny | 3353338 | 159.00 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5475249 | 97.38 MB/s | 798692 | 17133 | 1.0× |
