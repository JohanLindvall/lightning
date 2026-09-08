# JSON Deserialization Benchmarks

- generated 2026-09-08T12:32:45Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 84920 | 1498.76 MB/s | 49823 | 2 | 12.9× |
| LightningArena | 85039 | 1496.67 MB/s | 49825 | 2 | 12.9× |
| LightningDestructive | 86085 | 1478.49 MB/s | 49280 | 2 | 12.7× |
| SonicFastest | 184270 | 690.70 MB/s | 195915 | 10 | 5.9× |
| Sonic | 184674 | 689.19 MB/s | 196799 | 10 | 5.9× |
| Goccy | 203470 | 625.52 MB/s | 225404 | 884 | 5.4× |
| Easyjson | 213091 | 597.28 MB/s | 122864 | 14 | 5.1× |
| JSONV2 | 432199 | 294.48 MB/s | 195120 | 1805 | 2.5× |
| LightningDecodeAny | 440301 | 214.97 MB/s | 464779 | 9707 | 2.5× |
| Stdlib | 1093926 | 116.35 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2666352 | 844.24 MB/s | 2532848 | 1143 | 10.0× |
| Lightning | 2698852 | 834.08 MB/s | 2532850 | 1143 | 9.8× |
| LightningArena | 2707615 | 831.38 MB/s | 2532848 | 1143 | 9.8× |
| Sonic | 4901494 | 459.26 MB/s | 15233878 | 970 | 5.4× |
| SonicFastest | 4916031 | 457.90 MB/s | 15233846 | 970 | 5.4× |
| Goccy | 10498846 | 214.41 MB/s | 4126981 | 56532 | 2.5× |
| Easyjson | 10986552 | 204.89 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11614028 | 193.82 MB/s | 19380209 | 223896 | 2.3× |
| JSONV2 | 16338348 | 137.78 MB/s | 3123221 | 3083 | 1.6× |
| Stdlib | 26575123 | 84.71 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 370401 | 730.03 MB/s | 397296 | 567 | 9.3× |
| Lightning | 370861 | 729.12 MB/s | 397296 | 567 | 9.3× |
| LightningDestructive | 371883 | 727.12 MB/s | 397296 | 567 | 9.2× |
| Sonic | 647730 | 417.46 MB/s | 490333 | 968 | 5.3× |
| SonicFastest | 648808 | 416.77 MB/s | 496628 | 968 | 5.3× |
| Easyjson | 1406465 | 192.26 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1418408 | 190.64 MB/s | 542834 | 8122 | 2.4× |
| LightningDecodeAny | 1583811 | 170.73 MB/s | 2543877 | 29687 | 2.2× |
| JSONV2 | 2111645 | 128.05 MB/s | 348155 | 1628 | 1.6× |
| Stdlib | 3431562 | 78.80 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 958724 | 1801.56 MB/s | 765560 | 2798 | 13.8× |
| Lightning | 979132 | 1764.02 MB/s | 768140 | 2798 | 13.5× |
| LightningArena | 979679 | 1763.03 MB/s | 775473 | 2444 | 13.5× |
| SonicFastest | 2093205 | 825.15 MB/s | 2746187 | 4020 | 6.3× |
| Sonic | 2108087 | 819.32 MB/s | 2725377 | 4020 | 6.3× |
| Goccy | 2482828 | 695.66 MB/s | 2584941 | 14605 | 5.3× |
| Easyjson | 4242330 | 407.14 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4267199 | 404.76 MB/s | 1011635 | 7594 | 3.1× |
| LightningDecodeAny | 4372655 | 114.42 MB/s | 4964938 | 76577 | 3.0× |
| Stdlib | 13233058 | 130.52 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 834 | 2173.19 MB/s | 0 | 0 | 16.7× |
| Lightning | 835 | 2168.96 MB/s | 0 | 0 | 16.6× |
| LightningDestructive | 849 | 2134.53 MB/s | 0 | 0 | 16.4× |
| Easyjson | 2527 | 717.17 MB/s | 24 | 1 | 5.5× |
| Goccy | 2795 | 648.30 MB/s | 2608 | 4 | 5.0× |
| Sonic | 6105 | 296.78 MB/s | 3785 | 40 | 2.3× |
| SonicFastest | 6105 | 296.79 MB/s | 3793 | 40 | 2.3× |
| JSONV2 | 7782 | 232.85 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7907 | 229.05 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13906 | 130.30 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 858 | 2112.67 MB/s | 0 | 0 | 16.3× |
| LightningArena | 859 | 2108.81 MB/s | 0 | 0 | 16.2× |
| LightningDestructive | 882 | 2053.23 MB/s | 0 | 0 | 15.8× |
| Easyjson | 2524 | 717.91 MB/s | 24 | 1 | 5.5× |
| Goccy | 2850 | 635.89 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6036 | 300.18 MB/s | 3776 | 40 | 2.3× |
| SonicFastest | 6041 | 299.94 MB/s | 3804 | 40 | 2.3× |
| JSONV2 | 7668 | 236.31 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7936 | 228.21 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13944 | 129.95 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1029 | 1761.28 MB/s | 144 | 10 | 13.6× |
| LightningArena | 1036 | 1749.75 MB/s | 144 | 10 | 13.5× |
| LightningDestructive | 1089 | 1663.42 MB/s | 144 | 10 | 12.9× |
| Easyjson | 2748 | 659.33 MB/s | 144 | 10 | 5.1× |
| Goccy | 2952 | 613.87 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6202 | 292.18 MB/s | 3850 | 42 | 2.3× |
| Sonic | 6234 | 290.64 MB/s | 3871 | 42 | 2.2× |
| JSONV2 | 8058 | 224.87 MB/s | 632 | 7 | 1.7× |
| LightningDecodeAny | 8195 | 220.98 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 14023 | 129.21 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 594 | 832.36 MB/s | 160 | 1 | 9.3× |
| LightningDestructive | 594 | 831.18 MB/s | 160 | 1 | 9.3× |
| Sonic | 1248 | 395.72 MB/s | 979 | 6 | 4.4× |
| SonicFastest | 1257 | 392.85 MB/s | 994 | 6 | 4.4× |
| LightningDecodeAny | 1296 | 380.50 MB/s | 1296 | 26 | 4.3× |
| LightningArena | 1355 | 364.49 MB/s | 4120 | 2 | 4.1× |
| Easyjson | 2226 | 221.96 MB/s | 448 | 3 | 2.5× |
| Goccy | 2428 | 203.42 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3299 | 149.76 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5515 | 89.58 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 372 | 617.95 MB/s | 160 | 1 | 11.1× |
| LightningDestructive | 372 | 617.74 MB/s | 160 | 1 | 11.1× |
| Sonic | 890 | 258.45 MB/s | 651 | 6 | 4.6× |
| SonicFastest | 891 | 258.21 MB/s | 658 | 6 | 4.6× |
| LightningArena | 1100 | 209.16 MB/s | 4120 | 2 | 3.8× |
| LightningDecodeAny | 1127 | 203.15 MB/s | 1296 | 26 | 3.7× |
| Easyjson | 1400 | 164.34 MB/s | 448 | 3 | 2.9× |
| Goccy | 1581 | 145.46 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2479 | 92.79 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4129 | 55.70 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 51437 | 1266.24 MB/s | 97220 | 98 | 10.7× |
| Lightning | 52106 | 1249.98 MB/s | 103647 | 99 | 10.6× |
| LightningArena | 52434 | 1242.16 MB/s | 103640 | 99 | 10.5× |
| SonicFastest | 100995 | 644.90 MB/s | 157753 | 75 | 5.5× |
| Sonic | 102438 | 635.82 MB/s | 158676 | 75 | 5.4× |
| Goccy | 148208 | 439.46 MB/s | 228822 | 134 | 3.7× |
| LightningDecodeAny | 179976 | 296.31 MB/s | 180432 | 3241 | 3.1× |
| JSONV2 | 232267 | 280.42 MB/s | 206653 | 607 | 2.4× |
| Stdlib | 551278 | 118.15 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2099679 | 924.18 MB/s | 2864592 | 1380 | 11.0× |
| Lightning | 2157731 | 899.31 MB/s | 2864595 | 1380 | 10.7× |
| LightningArena | 2178801 | 890.61 MB/s | 2864594 | 1380 | 10.6× |
| SonicFastest | 4670581 | 415.47 MB/s | 14608590 | 1407 | 4.9× |
| Sonic | 4748279 | 408.67 MB/s | 14606973 | 1407 | 4.9× |
| Goccy | 4796766 | 404.54 MB/s | 4065031 | 13510 | 4.8× |
| Easyjson | 7507327 | 258.48 MB/s | 3871264 | 15043 | 3.1× |
| LightningDecodeAny | 9100770 | 213.22 MB/s | 7063040 | 218633 | 2.5× |
| JSONV2 | 11179895 | 173.57 MB/s | 3237217 | 13947 | 2.1× |
| Stdlib | 23074509 | 84.10 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 885130 | 3759.71 MB/s | 351704 | 1286 | 23.7× |
| LightningArena | 1351293 | 2462.70 MB/s | 2434667 | 1413 | 15.5× |
| Lightning | 1351887 | 2461.62 MB/s | 2434671 | 1413 | 15.5× |
| Sonic | 2718451 | 1224.16 MB/s | 6479624 | 4248 | 7.7× |
| SonicFastest | 2768181 | 1202.17 MB/s | 6509224 | 4248 | 7.6× |
| LightningDecodeAny | 3329444 | 923.21 MB/s | 4825778 | 55311 | 6.3× |
| Goccy | 4632395 | 718.38 MB/s | 3948907 | 3816 | 4.5× |
| JSONV2 | 7454994 | 446.39 MB/s | 5364520 | 13243 | 2.8× |
| Stdlib | 20983975 | 158.59 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 180948 | 1217.73 MB/s | 135872 | 226 | 11.2× |
| LightningArena | 181918 | 1211.24 MB/s | 135872 | 226 | 11.1× |
| LightningDestructive | 182753 | 1205.71 MB/s | 135872 | 226 | 11.1× |
| SonicFastest | 378268 | 582.51 MB/s | 294246 | 398 | 5.3× |
| Sonic | 380708 | 578.78 MB/s | 299657 | 398 | 5.3× |
| Goccy | 438340 | 502.68 MB/s | 364333 | 1067 | 4.6× |
| Easyjson | 546983 | 402.84 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 733785 | 300.29 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 851585 | 127.19 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2022936 | 108.92 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9287309 | 872.16 MB/s | 11845072 | 20816 | 9.5× |
| Lightning | 9498455 | 852.77 MB/s | 11845077 | 20816 | 9.3× |
| LightningArena | 9526804 | 850.24 MB/s | 11845073 | 20816 | 9.3× |
| Sonic | 16708642 | 484.78 MB/s | 70915411 | 40014 | 5.3× |
| SonicFastest | 16750340 | 483.57 MB/s | 70901598 | 40014 | 5.3× |
| Goccy | 23330490 | 347.19 MB/s | 17027273 | 107148 | 3.8× |
| Easyjson | 30966202 | 261.58 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 35443775 | 146.80 MB/s | 46279354 | 747112 | 2.5× |
| JSONV2 | 43832005 | 184.80 MB/s | 15233716 | 78972 | 2.0× |
| Stdlib | 88371475 | 91.66 MB/s | 15665066 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4351120 | 685.68 MB/s | 3780457 | 1514 | 10.7× |
| LightningDestructive | 4603684 | 648.06 MB/s | 3758856 | 29356 | 10.1× |
| Lightning | 4679355 | 637.58 MB/s | 3758859 | 29356 | 9.9× |
| SonicFastest | 8714608 | 342.35 MB/s | 26580119 | 56760 | 5.3× |
| Sonic | 8728423 | 341.81 MB/s | 26498846 | 56760 | 5.3× |
| LightningDecodeAny | 16198165 | 113.23 MB/s | 23982581 | 351152 | 2.9× |
| Goccy | 16804112 | 177.54 MB/s | 10696047 | 273652 | 2.8× |
| Easyjson | 16970503 | 175.80 MB/s | 9479440 | 30115 | 2.7× |
| JSONV2 | 24662900 | 120.97 MB/s | 9257138 | 86278 | 1.9× |
| Stdlib | 46494255 | 64.17 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 924250 | 782.90 MB/s | 907601 | 3618 | 12.3× |
| LightningArena | 928712 | 779.14 MB/s | 916258 | 37 | 12.2× |
| Lightning | 969141 | 746.64 MB/s | 907598 | 3618 | 11.7× |
| Sonic | 1763400 | 410.34 MB/s | 3190122 | 7226 | 6.4× |
| SonicFastest | 1767582 | 409.37 MB/s | 3194806 | 7226 | 6.4× |
| LightningDecodeAny | 3972142 | 163.78 MB/s | 6500455 | 76546 | 2.9× |
| Easyjson | 4167942 | 173.61 MB/s | 2847904 | 3698 | 2.7× |
| Goccy | 4825021 | 149.97 MB/s | 2860689 | 80276 | 2.4× |
| JSONV2 | 5521150 | 131.06 MB/s | 2704611 | 7318 | 2.1× |
| Stdlib | 11372651 | 63.63 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1373464 | 1148.45 MB/s | 907600 | 3618 | 11.3× |
| LightningArena | 1373951 | 1148.04 MB/s | 916257 | 37 | 11.3× |
| Lightning | 1415000 | 1114.74 MB/s | 907596 | 3618 | 11.0× |
| Sonic | 2230094 | 707.30 MB/s | 5781690 | 7226 | 7.0× |
| SonicFastest | 2271772 | 694.33 MB/s | 5782421 | 7226 | 6.8× |
| LightningDecodeAny | 3600307 | 209.26 MB/s | 6500457 | 76546 | 4.3× |
| Easyjson | 5537027 | 284.87 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5662252 | 278.57 MB/s | 3595921 | 80268 | 2.7× |
| JSONV2 | 6393846 | 246.70 MB/s | 2704590 | 7318 | 2.4× |
| Stdlib | 15532347 | 101.55 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 154975 | 968.70 MB/s | 81920 | 1 | 12.0× |
| LightningArena | 155244 | 967.02 MB/s | 81920 | 1 | 11.9× |
| Lightning | 155327 | 966.50 MB/s | 81920 | 1 | 11.9× |
| SonicFastest | 272611 | 550.69 MB/s | 259436 | 6 | 6.8× |
| Sonic | 274268 | 547.36 MB/s | 264724 | 6 | 6.8× |
| LightningDecodeAny | 419012 | 358.27 MB/s | 745764 | 10016 | 4.4× |
| Goccy | 871494 | 172.26 MB/s | 326784 | 10005 | 2.1× |
| JSONV2 | 1069701 | 140.34 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1854267 | 80.96 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 27269 | 1031.11 MB/s | 29236 | 101 | 11.0× |
| Lightning | 27339 | 1028.46 MB/s | 29237 | 101 | 11.0× |
| LightningDestructive | 27470 | 1023.57 MB/s | 29088 | 101 | 10.9× |
| SonicFastest | 62822 | 447.56 MB/s | 47036 | 103 | 4.8× |
| Sonic | 62946 | 446.69 MB/s | 46971 | 103 | 4.8× |
| Easyjson | 67988 | 413.56 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72183 | 389.52 MB/s | 59232 | 188 | 4.2× |
| JSONV2 | 134167 | 209.57 MB/s | 36895 | 242 | 2.2× |
| LightningDecodeAny | 148578 | 189.24 MB/s | 140938 | 2641 | 2.0× |
| Stdlib | 300773 | 93.48 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1514 | 1538.01 MB/s | 32 | 1 | 14.9× |
| Lightning | 1519 | 1532.32 MB/s | 32 | 1 | 14.9× |
| LightningDestructive | 1582 | 1471.31 MB/s | 32 | 1 | 14.3× |
| Goccy | 4076 | 571.15 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4253 | 547.38 MB/s | 192 | 2 | 5.3× |
| SonicFastest | 5036 | 462.28 MB/s | 4185 | 6 | 4.5× |
| Sonic | 5050 | 461.00 MB/s | 4192 | 6 | 4.5× |
| JSONV2 | 8404 | 277.00 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9881 | 170.52 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22562 | 103.18 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 179 | 1056.64 MB/s | 0 | 0 | 13.6× |
| Lightning | 179 | 1056.10 MB/s | 0 | 0 | 13.6× |
| LightningDestructive | 182 | 1037.87 MB/s | 0 | 0 | 13.4× |
| Goccy | 396 | 476.99 MB/s | 304 | 2 | 6.2× |
| Easyjson | 490 | 385.75 MB/s | 0 | 0 | 5.0× |
| Sonic | 813 | 232.51 MB/s | 515 | 4 | 3.0× |
| SonicFastest | 814 | 232.21 MB/s | 520 | 4 | 3.0× |
| JSONV2 | 1059 | 178.53 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1262 | 106.15 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2440 | 77.47 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1104 | 1984.34 MB/s | 0 | 0 | 14.3× |
| LightningArena | 1105 | 1983.60 MB/s | 0 | 0 | 14.3× |
| LightningDestructive | 1127 | 1943.28 MB/s | 0 | 0 | 14.1× |
| Goccy | 3112 | 703.96 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3197 | 685.42 MB/s | 24 | 1 | 5.0× |
| Sonic | 6384 | 343.21 MB/s | 3986 | 40 | 2.5× |
| SonicFastest | 6387 | 343.03 MB/s | 3978 | 40 | 2.5× |
| LightningDecodeAny | 7793 | 232.37 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 7897 | 277.45 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15840 | 138.32 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 545151 | 936.39 MB/s | 457537 | 1009 | 11.0× |
| Lightning | 555886 | 918.31 MB/s | 457537 | 1009 | 10.8× |
| LightningArena | 557640 | 915.42 MB/s | 457536 | 1009 | 10.8× |
| Sonic | 1159812 | 440.14 MB/s | 915532 | 2006 | 5.2× |
| SonicFastest | 1160800 | 439.76 MB/s | 906298 | 2006 | 5.2× |
| Goccy | 1165980 | 437.81 MB/s | 1139959 | 5006 | 5.1× |
| Easyjson | 1523533 | 335.06 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3215283 | 158.77 MB/s | 1076014 | 12646 | 1.9× |
| LightningDecodeAny | 3341582 | 138.10 MB/s | 2950649 | 64018 | 1.8× |
| Stdlib | 5999601 | 85.08 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 482 | 41075.25 MB/s | 0 | 0 | 224.8× |
| LightningArena | 482 | 41068.42 MB/s | 0 | 0 | 224.7× |
| LightningDestructive | 489 | 40492.09 MB/s | 0 | 0 | 221.6× |
| Goccy | 19960 | 991.42 MB/s | 20491 | 2 | 5.4× |
| SonicFastest | 26989 | 733.23 MB/s | 22159 | 4 | 4.0× |
| Sonic | 27013 | 732.57 MB/s | 22434 | 4 | 4.0× |
| JSONV2 | 29807 | 663.91 MB/s | 8 | 1 | 3.6× |
| Easyjson | 82038 | 241.22 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 82131 | 240.93 MB/s | 116864 | 2015 | 1.3× |
| Stdlib | 108286 | 182.75 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1746 | 10379.39 MB/s | 0 | 0 | 59.2× |
| LightningArena | 1811 | 10008.40 MB/s | 405 | 0 | 57.0× |
| Lightning | 1815 | 9984.25 MB/s | 405 | 0 | 56.9× |
| Easyjson | 3990 | 4542.49 MB/s | 432 | 2 | 25.9× |
| SonicFastest | 9826 | 1844.55 MB/s | 23077 | 6 | 10.5× |
| Sonic | 9924 | 1826.23 MB/s | 23522 | 6 | 10.4× |
| Goccy | 15635 | 1159.23 MB/s | 19459 | 2 | 6.6× |
| LightningDecodeAny | 16067 | 1112.95 MB/s | 29107 | 189 | 6.4× |
| JSONV2 | 45383 | 399.36 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103277 | 175.49 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2063878 | 973.17 MB/s | 3089566 | 6821 | 8.9× |
| LightningArena | 2137285 | 939.74 MB/s | 3100961 | 6700 | 8.6× |
| Lightning | 2137830 | 939.50 MB/s | 3097402 | 6823 | 8.6× |
| Goccy | 4169430 | 481.72 MB/s | 5413018 | 15837 | 4.4× |
| SonicFastest | 4356408 | 461.04 MB/s | 10952984 | 13683 | 4.2× |
| Sonic | 4361638 | 460.49 MB/s | 10990466 | 13683 | 4.2× |
| Easyjson | 4902034 | 409.73 MB/s | 2981520 | 7439 | 3.8× |
| JSONV2 | 6860483 | 292.76 MB/s | 3173696 | 14563 | 2.7× |
| LightningDecodeAny | 7490163 | 152.51 MB/s | 8517441 | 134006 | 2.5× |
| Stdlib | 18415783 | 109.06 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 847 | 648.11 MB/s | 480 | 1 | 6.5× |
| LightningArena | 850 | 646.26 MB/s | 480 | 1 | 6.5× |
| LightningDestructive | 851 | 644.93 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1628 | 336.63 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 2140 | 256.53 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2618 | 209.70 MB/s | 1961 | 26 | 2.1× |
| SonicFastest | 2619 | 209.64 MB/s | 1970 | 26 | 2.1× |
| Goccy | 2952 | 186.00 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3277 | 167.51 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5537 | 99.16 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 402233 | 1570.02 MB/s | 402729 | 545 | 13.3× |
| Lightning | 439957 | 1435.40 MB/s | 449280 | 548 | 12.2× |
| LightningArena | 441032 | 1431.90 MB/s | 451301 | 404 | 12.2× |
| Sonic | 989895 | 637.96 MB/s | 998406 | 1102 | 5.4× |
| SonicFastest | 994566 | 634.96 MB/s | 998098 | 1102 | 5.4× |
| Easyjson | 1142007 | 552.99 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1169865 | 539.82 MB/s | 987199 | 1201 | 4.6× |
| JSONV2 | 2190559 | 288.29 MB/s | 571615 | 3144 | 2.4× |
| LightningDecodeAny | 2297763 | 203.20 MB/s | 2081319 | 29820 | 2.3× |
| Stdlib | 5360421 | 117.81 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 536080 | 1049.11 MB/s | 549455 | 427 | 9.8× |
| Lightning | 613103 | 917.31 MB/s | 664281 | 435 | 8.5× |
| LightningArena | 614545 | 915.16 MB/s | 666162 | 289 | 8.5× |
| Sonic | 1010592 | 556.51 MB/s | 950562 | 1476 | 5.2× |
| SonicFastest | 1013908 | 554.69 MB/s | 954129 | 1476 | 5.2× |
| Goccy | 1324987 | 424.46 MB/s | 1038206 | 1029 | 4.0× |
| Easyjson | 1741878 | 322.87 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2378517 | 236.45 MB/s | 2080130 | 29328 | 2.2× |
| JSONV2 | 2756597 | 204.02 MB/s | 927442 | 3482 | 1.9× |
| Stdlib | 5234763 | 107.44 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 564670 | 944.23 MB/s | 333416 | 2084 | 9.6× |
| Lightning | 577968 | 922.50 MB/s | 367537 | 2086 | 9.4× |
| LightningArena | 581138 | 917.47 MB/s | 367549 | 2086 | 9.4× |
| Easyjson | 1106603 | 481.82 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1131001 | 471.42 MB/s | 1042645 | 4351 | 4.8× |
| SonicFastest | 1132475 | 470.81 MB/s | 1037399 | 4351 | 4.8× |
| Goccy | 1294734 | 411.80 MB/s | 1167233 | 5409 | 4.2× |
| JSONV2 | 2535757 | 210.26 MB/s | 745448 | 13288 | 2.1× |
| LightningDecodeAny | 3291278 | 162.00 MB/s | 3001770 | 49873 | 1.7× |
| Stdlib | 5439465 | 98.02 MB/s | 798693 | 17133 | 1.0× |
