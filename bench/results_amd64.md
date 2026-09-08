# JSON Deserialization Benchmarks

- generated 2026-09-08T12:32:46Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 59998 | 2121.34 MB/s | 49873 | 2 | 16.8× |
| Lightning | 60160 | 2115.60 MB/s | 49884 | 2 | 16.8× |
| LightningDestructive | 62881 | 2024.06 MB/s | 49280 | 2 | 16.1× |
| Easyjson | 181458 | 701.40 MB/s | 122864 | 14 | 5.6× |
| Goccy | 195636 | 650.57 MB/s | 225191 | 884 | 5.2× |
| Sonic | 198005 | 642.79 MB/s | 213932 | 15 | 5.1× |
| SonicFastest | 205487 | 619.38 MB/s | 213933 | 15 | 4.9× |
| JSONV2 | 323164 | 393.84 MB/s | 195127 | 1805 | 3.1× |
| LightningDecodeAny | 336671 | 281.14 MB/s | 466375 | 9707 | 3.0× |
| Stdlib | 1009327 | 126.10 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2574906 | 874.23 MB/s | 2532848 | 1143 | 9.2× |
| LightningArena | 2588949 | 869.48 MB/s | 2532849 | 1143 | 9.2× |
| Lightning | 2591593 | 868.60 MB/s | 2532848 | 1143 | 9.2× |
| Sonic | 3821276 | 589.08 MB/s | 4867783 | 2584 | 6.2× |
| SonicFastest | 3979090 | 565.72 MB/s | 4870109 | 2584 | 6.0× |
| LightningDecodeAny | 9683906 | 232.45 MB/s | 19380212 | 223896 | 2.5× |
| Goccy | 10033051 | 224.36 MB/s | 4164754 | 56533 | 2.4× |
| Easyjson | 10549972 | 213.37 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 13018563 | 172.91 MB/s | 3123194 | 3083 | 1.8× |
| Stdlib | 23784048 | 94.65 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 359723 | 751.70 MB/s | 397296 | 567 | 8.7× |
| Lightning | 362423 | 746.10 MB/s | 397296 | 567 | 8.6× |
| LightningDestructive | 371019 | 728.81 MB/s | 397297 | 567 | 8.4× |
| Sonic | 578834 | 467.15 MB/s | 641722 | 1147 | 5.4× |
| SonicFastest | 585962 | 461.47 MB/s | 641675 | 1147 | 5.3× |
| Easyjson | 1349813 | 200.33 MB/s | 330272 | 749 | 2.3× |
| Goccy | 1357298 | 199.22 MB/s | 541886 | 8122 | 2.3× |
| LightningDecodeAny | 1521688 | 177.70 MB/s | 2543878 | 29687 | 2.1× |
| JSONV2 | 1731296 | 156.19 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 3125462 | 86.52 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 744168 | 2320.99 MB/s | 765560 | 2798 | 17.4× |
| LightningArena | 757173 | 2281.12 MB/s | 774839 | 2444 | 17.1× |
| Lightning | 768169 | 2248.47 MB/s | 767731 | 2798 | 16.9× |
| SonicFastest | 1683506 | 1025.96 MB/s | 2696656 | 5547 | 7.7× |
| Sonic | 1684204 | 1025.53 MB/s | 2696277 | 5547 | 7.7× |
| Goccy | 1917161 | 900.92 MB/s | 2581604 | 14604 | 6.8× |
| LightningDecodeAny | 2956777 | 169.20 MB/s | 4963428 | 76577 | 4.4× |
| Easyjson | 3105139 | 556.24 MB/s | 972032 | 5389 | 4.2× |
| JSONV2 | 3359774 | 514.08 MB/s | 1011613 | 7594 | 3.9× |
| Stdlib | 12948944 | 133.39 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 540 | 3357.36 MB/s | 0 | 0 | 22.9× |
| LightningArena | 542 | 3343.50 MB/s | 0 | 0 | 22.8× |
| LightningDestructive | 581 | 3118.54 MB/s | 0 | 0 | 21.3× |
| Easyjson | 2224 | 814.88 MB/s | 24 | 1 | 5.6× |
| Goccy | 2628 | 689.40 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 4853 | 373.37 MB/s | 3348 | 38 | 2.5× |
| Sonic | 5097 | 355.52 MB/s | 3347 | 38 | 2.4× |
| JSONV2 | 6146 | 294.82 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 6598 | 274.48 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 12371 | 146.48 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 573 | 3164.30 MB/s | 0 | 0 | 21.3× |
| Lightning | 593 | 3054.81 MB/s | 0 | 0 | 20.6× |
| LightningDestructive | 604 | 2997.59 MB/s | 0 | 0 | 20.2× |
| Easyjson | 2225 | 814.31 MB/s | 24 | 1 | 5.5× |
| Goccy | 2658 | 681.69 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 5022 | 360.81 MB/s | 3348 | 38 | 2.4× |
| Sonic | 5173 | 350.30 MB/s | 3351 | 38 | 2.4× |
| JSONV2 | 6157 | 294.30 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 6562 | 275.98 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 12217 | 148.31 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 717 | 2526.93 MB/s | 144 | 10 | 17.1× |
| Lightning | 722 | 2510.83 MB/s | 144 | 10 | 17.0× |
| LightningDestructive | 766 | 2365.62 MB/s | 144 | 10 | 16.0× |
| Easyjson | 2299 | 788.07 MB/s | 144 | 10 | 5.3× |
| Goccy | 2480 | 730.68 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 5013 | 361.48 MB/s | 3365 | 40 | 2.4× |
| Sonic | 5253 | 344.92 MB/s | 3369 | 40 | 2.3× |
| JSONV2 | 6036 | 300.19 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 6695 | 270.49 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 12261 | 147.79 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 510 | 967.76 MB/s | 160 | 1 | 9.5× |
| LightningDestructive | 522 | 945.92 MB/s | 160 | 1 | 9.3× |
| SonicFastest | 962 | 513.22 MB/s | 1075 | 8 | 5.0× |
| Sonic | 969 | 509.73 MB/s | 1076 | 8 | 5.0× |
| LightningDecodeAny | 1094 | 450.76 MB/s | 1296 | 26 | 4.4× |
| LightningArena | 1228 | 402.24 MB/s | 4120 | 2 | 3.9× |
| Easyjson | 1839 | 268.64 MB/s | 448 | 3 | 2.6× |
| Goccy | 2000 | 246.96 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 2413 | 204.72 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4834 | 102.20 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 314 | 733.69 MB/s | 160 | 1 | 11.1× |
| LightningDestructive | 320 | 719.64 MB/s | 160 | 1 | 10.9× |
| Sonic | 682 | 337.17 MB/s | 801 | 8 | 5.1× |
| SonicFastest | 698 | 329.43 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 910 | 251.78 MB/s | 1296 | 26 | 3.8× |
| LightningArena | 1035 | 222.33 MB/s | 4120 | 2 | 3.4× |
| Easyjson | 1234 | 186.41 MB/s | 448 | 3 | 2.8× |
| Goccy | 1335 | 172.31 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 1879 | 122.38 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3470 | 66.28 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 43708 | 1490.16 MB/s | 103752 | 99 | 11.9× |
| Lightning | 43878 | 1484.38 MB/s | 103747 | 99 | 11.8× |
| LightningDestructive | 44450 | 1465.28 MB/s | 97220 | 98 | 11.7× |
| Sonic | 118431 | 549.96 MB/s | 235858 | 65 | 4.4× |
| SonicFastest | 119952 | 542.98 MB/s | 235956 | 65 | 4.3× |
| LightningDecodeAny | 142473 | 374.31 MB/s | 180559 | 3241 | 3.6× |
| Goccy | 145787 | 446.76 MB/s | 227804 | 134 | 3.6× |
| JSONV2 | 197598 | 329.62 MB/s | 206661 | 607 | 2.6× |
| Stdlib | 518238 | 125.68 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1705365 | 1137.86 MB/s | 2864593 | 1380 | 12.0× |
| LightningArena | 1738319 | 1116.29 MB/s | 2864593 | 1380 | 11.8× |
| Lightning | 1744727 | 1112.19 MB/s | 2864594 | 1380 | 11.7× |
| Goccy | 3975214 | 488.14 MB/s | 4062543 | 13509 | 5.2× |
| Sonic | 4987812 | 389.04 MB/s | 4881621 | 1736 | 4.1× |
| SonicFastest | 5068415 | 382.86 MB/s | 4880609 | 1736 | 4.0× |
| Easyjson | 5953658 | 325.93 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 7062036 | 274.78 MB/s | 7063040 | 218633 | 2.9× |
| JSONV2 | 9112704 | 212.94 MB/s | 3237188 | 13947 | 2.2× |
| Stdlib | 20495583 | 94.68 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 655784 | 5074.58 MB/s | 351704 | 1286 | 31.3× |
| Lightning | 993366 | 3350.06 MB/s | 2434746 | 1413 | 20.6× |
| LightningArena | 994924 | 3344.81 MB/s | 2434716 | 1413 | 20.6× |
| SonicFastest | 1845763 | 1802.96 MB/s | 5896162 | 4263 | 11.1× |
| Sonic | 1873435 | 1776.33 MB/s | 5892415 | 4263 | 10.9× |
| LightningDecodeAny | 2450822 | 1254.18 MB/s | 4825417 | 55311 | 8.4× |
| Goccy | 4487403 | 741.59 MB/s | 3948913 | 3816 | 4.6× |
| JSONV2 | 6480236 | 513.54 MB/s | 5364505 | 13243 | 3.2× |
| Stdlib | 20505307 | 162.29 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 132735 | 1660.04 MB/s | 135872 | 226 | 13.4× |
| Lightning | 133363 | 1652.23 MB/s | 135872 | 226 | 13.3× |
| LightningDestructive | 138555 | 1590.31 MB/s | 135872 | 226 | 12.8× |
| Goccy | 355882 | 619.15 MB/s | 364130 | 1066 | 5.0× |
| Sonic | 403725 | 545.78 MB/s | 350861 | 262 | 4.4× |
| SonicFastest | 404241 | 545.09 MB/s | 351016 | 262 | 4.4× |
| Easyjson | 458873 | 480.19 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 554559 | 397.34 MB/s | 129746 | 470 | 3.2× |
| LightningDecodeAny | 726193 | 149.15 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 1778370 | 123.90 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 7520510 | 1077.06 MB/s | 11845072 | 20816 | 10.4× |
| Lightning | 7746239 | 1045.67 MB/s | 11845077 | 20816 | 10.1× |
| LightningArena | 7768309 | 1042.70 MB/s | 11845073 | 20816 | 10.0× |
| Sonic | 15317080 | 528.82 MB/s | 19861092 | 41640 | 5.1× |
| SonicFastest | 15436647 | 524.73 MB/s | 19860332 | 41640 | 5.1× |
| Goccy | 20104827 | 402.89 MB/s | 19118073 | 107155 | 3.9× |
| Easyjson | 26847802 | 301.70 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 29444821 | 176.70 MB/s | 46279351 | 747112 | 2.6× |
| JSONV2 | 35343171 | 229.18 MB/s | 15233715 | 78972 | 2.2× |
| Stdlib | 77973571 | 103.88 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 3628093 | 822.32 MB/s | 3780456 | 1514 | 11.3× |
| LightningDestructive | 3749158 | 795.77 MB/s | 3758856 | 29356 | 10.9× |
| Lightning | 3867093 | 771.50 MB/s | 3758856 | 29356 | 10.6× |
| Sonic | 7381029 | 404.21 MB/s | 9130114 | 57804 | 5.6× |
| SonicFastest | 7382524 | 404.13 MB/s | 9130291 | 57804 | 5.5× |
| LightningDecodeAny | 12970352 | 141.41 MB/s | 23982579 | 351152 | 3.2× |
| Goccy | 13835333 | 215.64 MB/s | 9930730 | 273623 | 3.0× |
| Easyjson | 14133633 | 211.09 MB/s | 9479441 | 30115 | 2.9× |
| JSONV2 | 19270907 | 154.82 MB/s | 9257046 | 86278 | 2.1× |
| Stdlib | 40969048 | 72.82 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 868697 | 832.97 MB/s | 907600 | 3618 | 11.7× |
| LightningArena | 874895 | 827.07 MB/s | 916257 | 37 | 11.6× |
| Lightning | 912248 | 793.20 MB/s | 907596 | 3618 | 11.1× |
| SonicFastest | 1621636 | 446.21 MB/s | 2375373 | 3683 | 6.3× |
| Sonic | 1642580 | 440.52 MB/s | 2374399 | 3683 | 6.2× |
| LightningDecodeAny | 3808351 | 170.83 MB/s | 6500461 | 76546 | 2.7× |
| Goccy | 4129703 | 175.22 MB/s | 2734679 | 80268 | 2.5× |
| Easyjson | 4252613 | 170.15 MB/s | 2847907 | 3698 | 2.4× |
| JSONV2 | 4806018 | 150.56 MB/s | 2704705 | 7318 | 2.1× |
| Stdlib | 10159931 | 71.22 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1209485 | 1304.15 MB/s | 916256 | 37 | 11.7× |
| Lightning | 1219452 | 1293.49 MB/s | 907594 | 3618 | 11.6× |
| LightningDestructive | 1233424 | 1278.84 MB/s | 907600 | 3618 | 11.5× |
| Sonic | 2030163 | 776.96 MB/s | 3235044 | 3683 | 7.0× |
| SonicFastest | 2094889 | 752.95 MB/s | 3234295 | 3683 | 6.8× |
| LightningDecodeAny | 3453813 | 218.14 MB/s | 6500455 | 76546 | 4.1× |
| Easyjson | 4957493 | 318.18 MB/s | 2847906 | 3698 | 2.9× |
| Goccy | 5101696 | 309.18 MB/s | 3506217 | 80263 | 2.8× |
| JSONV2 | 5272192 | 299.18 MB/s | 2704552 | 7318 | 2.7× |
| Stdlib | 14201777 | 111.07 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 160428 | 935.77 MB/s | 81920 | 1 | 9.9× |
| LightningArena | 161362 | 930.36 MB/s | 81920 | 1 | 9.9× |
| LightningDestructive | 169758 | 884.34 MB/s | 81920 | 1 | 9.4× |
| Sonic | 298058 | 503.67 MB/s | 407713 | 16 | 5.3× |
| SonicFastest | 299251 | 501.67 MB/s | 407986 | 16 | 5.3× |
| LightningDecodeAny | 418861 | 358.40 MB/s | 745763 | 10016 | 3.8× |
| Goccy | 776783 | 193.26 MB/s | 327730 | 10005 | 2.1× |
| JSONV2 | 911159 | 164.76 MB/s | 357727 | 20 | 1.7× |
| Stdlib | 1593295 | 94.22 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 20566 | 1367.16 MB/s | 29281 | 101 | 12.6× |
| LightningArena | 21299 | 1320.09 MB/s | 29279 | 101 | 12.1× |
| LightningDestructive | 22500 | 1249.64 MB/s | 29088 | 101 | 11.5× |
| SonicFastest | 53975 | 520.93 MB/s | 59459 | 83 | 4.8× |
| Sonic | 54167 | 519.08 MB/s | 59463 | 83 | 4.8× |
| Easyjson | 58825 | 477.98 MB/s | 32304 | 138 | 4.4× |
| Goccy | 62737 | 448.17 MB/s | 59280 | 188 | 4.1× |
| JSONV2 | 101534 | 276.92 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 124290 | 226.22 MB/s | 141380 | 2641 | 2.1× |
| Stdlib | 258196 | 108.90 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1144 | 2035.28 MB/s | 32 | 1 | 17.0× |
| Lightning | 1160 | 2006.93 MB/s | 32 | 1 | 16.7× |
| LightningDestructive | 1196 | 1946.41 MB/s | 32 | 1 | 16.2× |
| Goccy | 3694 | 630.16 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 3984 | 584.33 MB/s | 192 | 2 | 4.9× |
| SonicFastest | 4744 | 490.75 MB/s | 3709 | 4 | 4.1× |
| Sonic | 4761 | 488.97 MB/s | 3710 | 4 | 4.1× |
| JSONV2 | 6084 | 382.66 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 7430 | 226.79 MB/s | 10200 | 195 | 2.6× |
| Stdlib | 19397 | 120.02 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 118 | 1603.53 MB/s | 0 | 0 | 17.2× |
| LightningArena | 119 | 1589.79 MB/s | 0 | 0 | 17.0× |
| LightningDestructive | 122 | 1548.08 MB/s | 0 | 0 | 16.6× |
| Goccy | 326 | 579.73 MB/s | 304 | 2 | 6.2× |
| Easyjson | 434 | 435.77 MB/s | 0 | 0 | 4.7× |
| Sonic | 600 | 315.14 MB/s | 341 | 3 | 3.4× |
| SonicFastest | 600 | 315.15 MB/s | 341 | 3 | 3.4× |
| JSONV2 | 743 | 254.29 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 919 | 145.76 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 2027 | 93.22 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 767 | 2856.81 MB/s | 0 | 0 | 18.7× |
| Lightning | 778 | 2816.90 MB/s | 0 | 0 | 18.4× |
| LightningDestructive | 802 | 2731.22 MB/s | 0 | 0 | 17.9× |
| Easyjson | 2780 | 788.21 MB/s | 24 | 1 | 5.2× |
| Goccy | 3029 | 723.26 MB/s | 2864 | 4 | 4.7× |
| SonicFastest | 5311 | 412.56 MB/s | 3601 | 38 | 2.7× |
| Sonic | 5488 | 399.23 MB/s | 3601 | 38 | 2.6× |
| JSONV2 | 6303 | 347.62 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 6584 | 275.05 MB/s | 7552 | 158 | 2.2× |
| Stdlib | 14350 | 152.68 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 373672 | 1366.11 MB/s | 457537 | 1009 | 14.1× |
| LightningArena | 389734 | 1309.80 MB/s | 457537 | 1009 | 13.5× |
| Lightning | 390401 | 1307.57 MB/s | 457537 | 1009 | 13.5× |
| Goccy | 979234 | 521.30 MB/s | 1138898 | 5006 | 5.4× |
| Easyjson | 1176325 | 433.96 MB/s | 863777 | 3012 | 4.5× |
| SonicFastest | 1203949 | 424.00 MB/s | 1310324 | 2014 | 4.4× |
| Sonic | 1221524 | 417.90 MB/s | 1310589 | 2014 | 4.3× |
| JSONV2 | 2397242 | 212.94 MB/s | 1075960 | 12645 | 2.2× |
| LightningDecodeAny | 2581341 | 178.77 MB/s | 2950649 | 64018 | 2.0× |
| Stdlib | 5263357 | 96.99 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 465 | 42553.38 MB/s | 0 | 0 | 261.0× |
| Lightning | 467 | 42405.31 MB/s | 0 | 0 | 260.0× |
| LightningDestructive | 659 | 30034.36 MB/s | 0 | 0 | 184.2× |
| SonicFastest | 6074 | 3258.08 MB/s | 21062 | 3 | 20.0× |
| Goccy | 20897 | 947.00 MB/s | 20492 | 2 | 5.8× |
| Sonic | 22746 | 870.00 MB/s | 20607 | 3 | 5.3× |
| JSONV2 | 28084 | 704.63 MB/s | 8 | 1 | 4.3× |
| LightningDecodeAny | 73265 | 270.09 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 93170 | 212.40 MB/s | 0 | 0 | 1.3× |
| Stdlib | 121343 | 163.08 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1306 | 13882.49 MB/s | 405 | 0 | 80.5× |
| LightningArena | 1310 | 13838.24 MB/s | 405 | 0 | 80.3× |
| LightningDestructive | 1436 | 12619.48 MB/s | 0 | 0 | 73.2× |
| Easyjson | 3661 | 4950.11 MB/s | 432 | 2 | 28.7× |
| Sonic | 7425 | 2441.03 MB/s | 20435 | 5 | 14.2× |
| SonicFastest | 7477 | 2423.86 MB/s | 20443 | 5 | 14.1× |
| LightningDecodeAny | 14487 | 1234.39 MB/s | 29137 | 189 | 7.3× |
| Goccy | 20473 | 885.27 MB/s | 19460 | 2 | 5.1× |
| JSONV2 | 37636 | 481.55 MB/s | 16500 | 50 | 2.8× |
| Stdlib | 105151 | 172.36 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1571021 | 1278.46 MB/s | 3089565 | 6821 | 10.5× |
| Lightning | 1638144 | 1226.08 MB/s | 3096847 | 6822 | 10.0× |
| LightningArena | 1640544 | 1224.29 MB/s | 3099763 | 6699 | 10.0× |
| Goccy | 3670577 | 547.19 MB/s | 5410269 | 15831 | 4.5× |
| SonicFastest | 3994744 | 502.78 MB/s | 5158494 | 7085 | 4.1× |
| Sonic | 4055964 | 495.20 MB/s | 5159913 | 7085 | 4.1× |
| Easyjson | 4253942 | 472.15 MB/s | 2981485 | 7439 | 3.9× |
| LightningDecodeAny | 5372731 | 212.61 MB/s | 8514780 | 134005 | 3.1× |
| JSONV2 | 5760005 | 348.70 MB/s | 3173676 | 14562 | 2.9× |
| Stdlib | 16461031 | 122.02 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 629 | 872.48 MB/s | 480 | 1 | 7.6× |
| Lightning | 632 | 869.31 MB/s | 480 | 1 | 7.6× |
| LightningDestructive | 640 | 857.40 MB/s | 480 | 1 | 7.5× |
| LightningDecodeAny | 1333 | 411.02 MB/s | 2021 | 46 | 3.6× |
| Easyjson | 1517 | 362.01 MB/s | 1616 | 5 | 3.2× |
| SonicFastest | 1662 | 330.25 MB/s | 2262 | 8 | 2.9× |
| Sonic | 1712 | 320.75 MB/s | 2262 | 8 | 2.8× |
| JSONV2 | 2306 | 238.09 MB/s | 1664 | 7 | 2.1× |
| Goccy | 2357 | 232.93 MB/s | 2129 | 43 | 2.0× |
| Stdlib | 4805 | 114.26 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 309058 | 2043.35 MB/s | 402728 | 545 | 15.5× |
| LightningArena | 350977 | 1799.30 MB/s | 450903 | 404 | 13.7× |
| Lightning | 354980 | 1779.01 MB/s | 449116 | 548 | 13.5× |
| Sonic | 898065 | 703.19 MB/s | 1069847 | 814 | 5.3× |
| SonicFastest | 907274 | 696.06 MB/s | 1069664 | 814 | 5.3× |
| Easyjson | 996779 | 633.55 MB/s | 422504 | 936 | 4.8× |
| Goccy | 1094354 | 577.07 MB/s | 989073 | 1200 | 4.4× |
| JSONV2 | 1703030 | 370.82 MB/s | 571593 | 3144 | 2.8× |
| LightningDecodeAny | 1814971 | 257.25 MB/s | 2079523 | 29819 | 2.6× |
| Stdlib | 4801184 | 131.53 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 465988 | 1206.91 MB/s | 549057 | 427 | 9.8× |
| LightningArena | 551335 | 1020.08 MB/s | 665769 | 288 | 8.3× |
| Lightning | 554756 | 1013.79 MB/s | 664099 | 434 | 8.2× |
| SonicFastest | 1106867 | 508.11 MB/s | 1346766 | 1184 | 4.1× |
| Sonic | 1110211 | 506.58 MB/s | 1346980 | 1184 | 4.1× |
| Goccy | 1254031 | 448.48 MB/s | 1043015 | 1028 | 3.6× |
| Easyjson | 1560917 | 360.31 MB/s | 775153 | 1254 | 2.9× |
| LightningDecodeAny | 2016566 | 278.89 MB/s | 2078706 | 29327 | 2.3× |
| JSONV2 | 2443096 | 230.20 MB/s | 927405 | 3482 | 1.9× |
| Stdlib | 4566783 | 123.15 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 410978 | 1297.34 MB/s | 333416 | 2084 | 12.0× |
| Lightning | 447789 | 1190.69 MB/s | 367820 | 2086 | 11.0× |
| LightningArena | 448907 | 1187.73 MB/s | 367824 | 2086 | 11.0× |
| Easyjson | 1003069 | 531.55 MB/s | 428362 | 3273 | 4.9× |
| SonicFastest | 1064801 | 500.73 MB/s | 983649 | 3082 | 4.6× |
| Sonic | 1074442 | 496.24 MB/s | 983182 | 3082 | 4.6× |
| Goccy | 1202443 | 443.41 MB/s | 1167089 | 5409 | 4.1× |
| JSONV2 | 2069030 | 257.69 MB/s | 745421 | 13288 | 2.4× |
| LightningDecodeAny | 2717821 | 196.18 MB/s | 3000254 | 49872 | 1.8× |
| Stdlib | 4917938 | 108.41 MB/s | 798692 | 17133 | 1.0× |
