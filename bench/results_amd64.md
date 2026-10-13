# JSON Deserialization Benchmarks

- generated 2026-10-13T08:26:39Z
- go version go1.26.8 linux/amd64
- cpu: INTEL(R) XEON(R) PLATINUM 8573C (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 78353 | 1624.39 MB/s | 49893 | 2 | 13.3× |
| Lightning | 78876 | 1613.62 MB/s | 49894 | 2 | 13.2× |
| LightningDestructive | 80338 | 1584.24 MB/s | 49280 | 2 | 12.9× |
| SonicFastest | 151625 | 839.40 MB/s | 213446 | 15 | 6.9× |
| Sonic | 153018 | 831.77 MB/s | 213453 | 15 | 6.8× |
| Easyjson | 198115 | 642.43 MB/s | 122864 | 14 | 5.2× |
| Goccy | 211350 | 602.20 MB/s | 225128 | 884 | 4.9× |
| LightningDecodeAny | 381182 | 248.31 MB/s | 465547 | 9706 | 2.7× |
| JSONV2 | 383714 | 331.69 MB/s | 195128 | 1805 | 2.7× |
| Stdlib | 1038650 | 122.54 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2107623 | 1068.05 MB/s | 2532848 | 1143 | 12.1× |
| LightningArena | 2111569 | 1066.06 MB/s | 2532850 | 1143 | 12.1× |
| Lightning | 2127591 | 1058.03 MB/s | 2532849 | 1143 | 12.0× |
| Sonic | 4029413 | 558.65 MB/s | 4892153 | 2584 | 6.3× |
| SonicFastest | 4043936 | 556.65 MB/s | 4896551 | 2584 | 6.3× |
| LightningDecodeAny | 7299784 | 308.37 MB/s | 6828998 | 223498 | 3.5× |
| Goccy | 10404434 | 216.35 MB/s | 4255865 | 56538 | 2.4× |
| Easyjson | 12216623 | 184.26 MB/s | 3099808 | 2120 | 2.1× |
| JSONV2 | 14958256 | 150.49 MB/s | 3123188 | 3083 | 1.7× |
| Stdlib | 25464319 | 88.40 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 315237 | 857.78 MB/s | 397296 | 567 | 10.6× |
| LightningDestructive | 328590 | 822.92 MB/s | 397296 | 567 | 10.1× |
| LightningArena | 353206 | 765.57 MB/s | 397297 | 567 | 9.4× |
| Sonic | 636573 | 424.78 MB/s | 641681 | 1147 | 5.2× |
| SonicFastest | 645199 | 419.10 MB/s | 641594 | 1147 | 5.2× |
| LightningDecodeAny | 1062701 | 254.45 MB/s | 845692 | 29656 | 3.1× |
| Goccy | 1414231 | 191.20 MB/s | 541768 | 8122 | 2.4× |
| Easyjson | 1595325 | 169.50 MB/s | 330273 | 749 | 2.1× |
| JSONV2 | 1957219 | 138.16 MB/s | 348161 | 1628 | 1.7× |
| Stdlib | 3333306 | 81.12 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 758108 | 2278.31 MB/s | 765560 | 2798 | 15.7× |
| Lightning | 770700 | 2241.08 MB/s | 767904 | 2798 | 15.4× |
| LightningArena | 783430 | 2204.67 MB/s | 775016 | 2444 | 15.2× |
| SonicFastest | 1701345 | 1015.20 MB/s | 2711741 | 5547 | 7.0× |
| Sonic | 1714748 | 1007.26 MB/s | 2711750 | 5547 | 6.9× |
| Goccy | 2204191 | 783.60 MB/s | 2581855 | 14604 | 5.4× |
| LightningDecodeAny | 3240136 | 154.41 MB/s | 4492036 | 67881 | 3.7× |
| Easyjson | 3693600 | 467.62 MB/s | 972032 | 5389 | 3.2× |
| JSONV2 | 3890581 | 443.94 MB/s | 1011614 | 7594 | 3.1× |
| Stdlib | 11885178 | 145.32 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 752 | 2410.35 MB/s | 0 | 0 | 17.1× |
| LightningArena | 757 | 2393.71 MB/s | 0 | 0 | 16.9× |
| LightningDestructive | 804 | 2255.00 MB/s | 0 | 0 | 16.0× |
| Easyjson | 2191 | 827.08 MB/s | 24 | 1 | 5.9× |
| Goccy | 2772 | 653.57 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 4869 | 372.15 MB/s | 3348 | 38 | 2.6× |
| Sonic | 5022 | 360.81 MB/s | 3346 | 38 | 2.6× |
| JSONV2 | 6432 | 281.70 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 7147 | 253.39 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 12827 | 141.27 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 767 | 2363.69 MB/s | 0 | 0 | 17.1× |
| LightningArena | 772 | 2347.66 MB/s | 0 | 0 | 17.0× |
| LightningDestructive | 812 | 2231.72 MB/s | 0 | 0 | 16.1× |
| Easyjson | 2221 | 815.85 MB/s | 24 | 1 | 5.9× |
| Goccy | 2718 | 666.74 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 4811 | 376.62 MB/s | 3348 | 38 | 2.7× |
| Sonic | 5045 | 359.17 MB/s | 3347 | 38 | 2.6× |
| JSONV2 | 6455 | 280.71 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 7117 | 254.45 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13103 | 138.28 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 900 | 2013.29 MB/s | 144 | 10 | 14.5× |
| Lightning | 913 | 1984.94 MB/s | 144 | 10 | 14.3× |
| LightningDestructive | 968 | 1871.56 MB/s | 144 | 10 | 13.5× |
| Easyjson | 2564 | 706.73 MB/s | 144 | 10 | 5.1× |
| Goccy | 2915 | 621.52 MB/s | 2600 | 5 | 4.5× |
| SonicFastest | 5128 | 353.34 MB/s | 3371 | 40 | 2.5× |
| Sonic | 5324 | 340.38 MB/s | 3371 | 40 | 2.5× |
| JSONV2 | 7063 | 256.55 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 7274 | 248.97 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13049 | 138.86 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 511 | 966.16 MB/s | 160 | 1 | 10.1× |
| LightningDestructive | 522 | 945.75 MB/s | 160 | 1 | 9.9× |
| SonicFastest | 1004 | 491.87 MB/s | 1076 | 8 | 5.2× |
| LightningDecodeAny | 1007 | 489.77 MB/s | 1040 | 25 | 5.2× |
| Sonic | 1011 | 488.81 MB/s | 1077 | 8 | 5.1× |
| LightningArena | 1448 | 341.06 MB/s | 4120 | 2 | 3.6× |
| Easyjson | 2040 | 242.14 MB/s | 448 | 3 | 2.5× |
| Goccy | 2145 | 230.35 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 2653 | 186.18 MB/s | 528 | 7 | 2.0× |
| Stdlib | 5189 | 95.20 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 253 | 908.41 MB/s | 160 | 1 | 14.8× |
| LightningDestructive | 260 | 885.03 MB/s | 160 | 1 | 14.4× |
| Sonic | 737 | 312.18 MB/s | 801 | 8 | 5.1× |
| SonicFastest | 742 | 309.91 MB/s | 802 | 8 | 5.0× |
| LightningDecodeAny | 867 | 264.19 MB/s | 1040 | 25 | 4.3× |
| LightningArena | 1218 | 188.91 MB/s | 4120 | 2 | 3.1× |
| Easyjson | 1350 | 170.37 MB/s | 448 | 3 | 2.8× |
| Goccy | 1483 | 155.08 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2151 | 106.93 MB/s | 528 | 7 | 1.7× |
| Stdlib | 3742 | 61.47 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 52257 | 1246.38 MB/s | 103749 | 99 | 9.9× |
| Lightning | 52711 | 1235.65 MB/s | 103758 | 99 | 9.8× |
| LightningDestructive | 54627 | 1192.31 MB/s | 97220 | 98 | 9.5× |
| Sonic | 118961 | 547.51 MB/s | 235511 | 65 | 4.4× |
| SonicFastest | 125758 | 517.91 MB/s | 236371 | 65 | 4.1× |
| LightningDecodeAny | 156183 | 341.45 MB/s | 176848 | 3237 | 3.3× |
| Goccy | 157450 | 413.67 MB/s | 227971 | 134 | 3.3× |
| JSONV2 | 228788 | 284.68 MB/s | 206671 | 607 | 2.3× |
| Stdlib | 518454 | 125.63 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1711298 | 1133.92 MB/s | 2185297 | 1350 | 12.8× |
| LightningArena | 1751013 | 1108.20 MB/s | 2185297 | 1350 | 12.6× |
| Lightning | 1760465 | 1102.25 MB/s | 2185298 | 1350 | 12.5× |
| Sonic | 3825555 | 507.24 MB/s | 4882731 | 1736 | 5.7× |
| SonicFastest | 3847366 | 504.36 MB/s | 4883220 | 1736 | 5.7× |
| Goccy | 4339242 | 447.19 MB/s | 4063949 | 13509 | 5.1× |
| Easyjson | 7032823 | 275.92 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 7662702 | 253.24 MB/s | 6627985 | 206416 | 2.9× |
| JSONV2 | 10448947 | 185.71 MB/s | 3237181 | 13947 | 2.1× |
| Stdlib | 21987009 | 88.26 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 830114 | 4008.89 MB/s | 351704 | 1286 | 23.5× |
| LightningArena | 1202144 | 2768.25 MB/s | 2434797 | 1413 | 16.2× |
| Lightning | 1228167 | 2709.59 MB/s | 2434738 | 1413 | 15.9× |
| SonicFastest | 1667583 | 1995.60 MB/s | 5896672 | 4263 | 11.7× |
| Sonic | 1685661 | 1974.20 MB/s | 5896641 | 4263 | 11.6× |
| LightningDecodeAny | 2831543 | 1085.54 MB/s | 4825539 | 55311 | 6.9× |
| Goccy | 4542582 | 732.59 MB/s | 3948914 | 3816 | 4.3× |
| JSONV2 | 6709673 | 495.98 MB/s | 5364498 | 13243 | 2.9× |
| Stdlib | 19475154 | 170.88 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 149015 | 1478.68 MB/s | 135392 | 226 | 11.9× |
| LightningArena | 150052 | 1468.46 MB/s | 135392 | 226 | 11.8× |
| LightningDestructive | 157100 | 1402.58 MB/s | 135392 | 226 | 11.3× |
| SonicFastest | 301979 | 729.67 MB/s | 350543 | 262 | 5.9× |
| Sonic | 306127 | 719.79 MB/s | 350662 | 262 | 5.8× |
| Goccy | 405346 | 543.60 MB/s | 363890 | 1066 | 4.4× |
| Easyjson | 465958 | 472.89 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 606492 | 363.31 MB/s | 129748 | 470 | 2.9× |
| LightningDecodeAny | 804091 | 134.70 MB/s | 854737 | 11700 | 2.2× |
| Stdlib | 1775000 | 124.14 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 7349678 | 1102.09 MB/s | 8109652 | 20809 | 11.3× |
| Lightning | 7488231 | 1081.70 MB/s | 8109652 | 20809 | 11.1× |
| LightningArena | 7700379 | 1051.90 MB/s | 8109648 | 20809 | 10.8× |
| Sonic | 14074387 | 575.52 MB/s | 19865578 | 41640 | 5.9× |
| SonicFastest | 14113677 | 573.91 MB/s | 19864247 | 41640 | 5.9× |
| Goccy | 22709973 | 356.67 MB/s | 18100654 | 107152 | 3.7× |
| LightningDecodeAny | 28735945 | 181.06 MB/s | 28359831 | 746961 | 2.9× |
| Easyjson | 30858232 | 262.49 MB/s | 15059616 | 41643 | 2.7× |
| JSONV2 | 42174151 | 192.06 MB/s | 15233715 | 78972 | 2.0× |
| Stdlib | 83350298 | 97.18 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 3061823 | 974.41 MB/s | 3780457 | 1514 | 14.3× |
| LightningDestructive | 3148807 | 947.49 MB/s | 3758856 | 29356 | 13.9× |
| Lightning | 3251447 | 917.58 MB/s | 3758860 | 29356 | 13.4× |
| SonicFastest | 7539268 | 395.72 MB/s | 9133129 | 57804 | 5.8× |
| Sonic | 7574086 | 393.90 MB/s | 9133705 | 57804 | 5.8× |
| LightningDecodeAny | 13222639 | 138.72 MB/s | 18225371 | 350883 | 3.3× |
| Goccy | 15444307 | 193.18 MB/s | 9808500 | 273617 | 2.8× |
| Easyjson | 16646300 | 179.23 MB/s | 9479441 | 30115 | 2.6× |
| JSONV2 | 22518395 | 132.49 MB/s | 9257088 | 86278 | 1.9× |
| Stdlib | 43666199 | 68.32 MB/s | 9258083 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 590504 | 1225.39 MB/s | 907600 | 3618 | 18.3× |
| LightningArena | 617977 | 1170.91 MB/s | 916256 | 37 | 17.5× |
| Lightning | 635767 | 1138.15 MB/s | 907595 | 3618 | 17.0× |
| Sonic | 1679611 | 430.81 MB/s | 2378660 | 3683 | 6.5× |
| SonicFastest | 1694159 | 427.11 MB/s | 2379981 | 3683 | 6.4× |
| LightningDecodeAny | 3614811 | 179.97 MB/s | 5691592 | 76540 | 3.0× |
| Easyjson | 4559990 | 158.68 MB/s | 2847908 | 3698 | 2.4× |
| Goccy | 4644316 | 155.80 MB/s | 2760094 | 80270 | 2.3× |
| JSONV2 | 5277882 | 137.10 MB/s | 2704706 | 7318 | 2.1× |
| Stdlib | 10833544 | 66.79 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 835476 | 1887.97 MB/s | 907600 | 3618 | 17.2× |
| LightningArena | 867924 | 1817.39 MB/s | 916256 | 37 | 16.6× |
| Lightning | 875274 | 1802.12 MB/s | 907594 | 3618 | 16.5× |
| SonicFastest | 2083070 | 757.23 MB/s | 3269288 | 3683 | 6.9× |
| Sonic | 2088348 | 755.31 MB/s | 3268905 | 3683 | 6.9× |
| LightningDecodeAny | 3155023 | 238.79 MB/s | 5691587 | 76540 | 4.6× |
| Easyjson | 5690259 | 277.20 MB/s | 2847905 | 3698 | 2.5× |
| Goccy | 5733741 | 275.10 MB/s | 3555190 | 80266 | 2.5× |
| JSONV2 | 6066660 | 260.00 MB/s | 2704552 | 7318 | 2.4× |
| Stdlib | 14409279 | 109.47 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 84753 | 1771.31 MB/s | 81920 | 1 | 19.5× |
| LightningArena | 84760 | 1771.16 MB/s | 81920 | 1 | 19.5× |
| LightningDestructive | 94108 | 1595.23 MB/s | 81920 | 1 | 17.6× |
| Sonic | 336319 | 446.37 MB/s | 408526 | 16 | 4.9× |
| SonicFastest | 347468 | 432.05 MB/s | 409711 | 16 | 4.8× |
| LightningDecodeAny | 415237 | 361.53 MB/s | 745507 | 10015 | 4.0× |
| Goccy | 824325 | 182.12 MB/s | 331554 | 10005 | 2.0× |
| JSONV2 | 965184 | 155.54 MB/s | 357725 | 20 | 1.7× |
| Stdlib | 1651866 | 90.88 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 24743 | 1136.36 MB/s | 29277 | 101 | 10.9× |
| Lightning | 24775 | 1134.90 MB/s | 29275 | 101 | 10.9× |
| LightningDestructive | 26158 | 1074.90 MB/s | 29088 | 101 | 10.3× |
| Sonic | 49411 | 569.04 MB/s | 59548 | 83 | 5.5× |
| SonicFastest | 50153 | 560.63 MB/s | 59563 | 83 | 5.4× |
| Easyjson | 62267 | 451.56 MB/s | 32304 | 138 | 4.3× |
| Goccy | 63715 | 441.29 MB/s | 59247 | 188 | 4.2× |
| JSONV2 | 111069 | 253.15 MB/s | 36896 | 242 | 2.4× |
| LightningDecodeAny | 127994 | 219.67 MB/s | 134136 | 2639 | 2.1× |
| Stdlib | 270391 | 103.99 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1418 | 1642.07 MB/s | 32 | 1 | 14.0× |
| LightningArena | 1424 | 1634.27 MB/s | 32 | 1 | 13.9× |
| LightningDestructive | 1499 | 1552.58 MB/s | 32 | 1 | 13.2× |
| SonicFastest | 3509 | 663.41 MB/s | 3710 | 4 | 5.7× |
| Sonic | 3520 | 661.37 MB/s | 3711 | 4 | 5.6× |
| Goccy | 3826 | 608.44 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 4247 | 548.20 MB/s | 192 | 2 | 4.7× |
| JSONV2 | 6682 | 348.40 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 8175 | 206.11 MB/s | 9936 | 194 | 2.4× |
| Stdlib | 19828 | 117.41 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 151 | 1249.16 MB/s | 0 | 0 | 14.6× |
| LightningArena | 152 | 1240.18 MB/s | 0 | 0 | 14.5× |
| LightningDestructive | 157 | 1203.19 MB/s | 0 | 0 | 14.0× |
| Goccy | 393 | 480.90 MB/s | 304 | 2 | 5.6× |
| SonicFastest | 437 | 432.16 MB/s | 342 | 3 | 5.0× |
| Sonic | 438 | 431.39 MB/s | 342 | 3 | 5.0× |
| Easyjson | 479 | 394.33 MB/s | 0 | 0 | 4.6× |
| JSONV2 | 844 | 223.88 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1040 | 128.83 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2207 | 85.65 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1035 | 2116.05 MB/s | 0 | 0 | 14.2× |
| Lightning | 1041 | 2104.03 MB/s | 0 | 0 | 14.1× |
| LightningDestructive | 1110 | 1974.29 MB/s | 0 | 0 | 13.2× |
| Easyjson | 2733 | 801.59 MB/s | 24 | 1 | 5.4× |
| Goccy | 3087 | 709.68 MB/s | 2864 | 4 | 4.8× |
| SonicFastest | 5243 | 417.89 MB/s | 3602 | 38 | 2.8× |
| Sonic | 5375 | 407.64 MB/s | 3603 | 38 | 2.7× |
| JSONV2 | 6698 | 327.11 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 7241 | 250.09 MB/s | 7552 | 158 | 2.0× |
| Stdlib | 14683 | 149.22 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 417334 | 1223.18 MB/s | 318401 | 1005 | 13.5× |
| LightningArena | 420393 | 1214.28 MB/s | 318400 | 1005 | 13.4× |
| Lightning | 423206 | 1206.21 MB/s | 318401 | 1005 | 13.4× |
| SonicFastest | 943513 | 541.04 MB/s | 1309828 | 2014 | 6.0× |
| Sonic | 944371 | 540.55 MB/s | 1310658 | 2014 | 6.0× |
| Goccy | 1056640 | 483.11 MB/s | 1139806 | 5006 | 5.3× |
| Easyjson | 1457415 | 350.26 MB/s | 863776 | 3012 | 3.9× |
| JSONV2 | 2871968 | 177.74 MB/s | 1075960 | 12645 | 2.0× |
| LightningDecodeAny | 2919397 | 158.07 MB/s | 2742393 | 64017 | 1.9× |
| Stdlib | 5652627 | 90.31 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 206 | 96147.39 MB/s | 0 | 0 | 468.2× |
| LightningArena | 207 | 95574.26 MB/s | 0 | 0 | 465.2× |
| LightningDestructive | 341 | 58004.15 MB/s | 0 | 0 | 282.4× |
| SonicFastest | 5434 | 3641.84 MB/s | 21074 | 3 | 17.7× |
| Goccy | 21666 | 913.38 MB/s | 20492 | 2 | 4.4× |
| Sonic | 28694 | 689.65 MB/s | 20608 | 3 | 3.4× |
| JSONV2 | 30182 | 655.66 MB/s | 8 | 1 | 3.2× |
| LightningDecodeAny | 78812 | 251.08 MB/s | 116608 | 2014 | 1.2× |
| Easyjson | 84722 | 233.58 MB/s | 0 | 0 | 1.1× |
| Stdlib | 96351 | 205.38 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1787 | 10141.00 MB/s | 405 | 0 | 53.7× |
| Lightning | 1803 | 10050.85 MB/s | 405 | 0 | 53.2× |
| LightningDestructive | 2213 | 8190.94 MB/s | 0 | 0 | 43.3× |
| Easyjson | 3968 | 4567.29 MB/s | 432 | 2 | 24.2× |
| Sonic | 6481 | 2796.42 MB/s | 20444 | 5 | 14.8× |
| SonicFastest | 6538 | 2772.30 MB/s | 20432 | 5 | 14.7× |
| LightningDecodeAny | 16101 | 1110.58 MB/s | 29140 | 189 | 6.0× |
| Goccy | 18388 | 985.63 MB/s | 19460 | 2 | 5.2× |
| JSONV2 | 40281 | 449.94 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 95923 | 188.94 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1795880 | 1118.39 MB/s | 3089565 | 6821 | 9.5× |
| Lightning | 1870843 | 1073.58 MB/s | 3096858 | 6822 | 9.1× |
| LightningArena | 1871399 | 1073.26 MB/s | 3099845 | 6699 | 9.1× |
| SonicFastest | 3296833 | 609.22 MB/s | 5175821 | 7085 | 5.1× |
| Sonic | 3323951 | 604.25 MB/s | 5176058 | 7085 | 5.1× |
| Goccy | 3891007 | 516.19 MB/s | 5413628 | 15835 | 4.4× |
| Easyjson | 4797686 | 418.64 MB/s | 2981490 | 7439 | 3.5× |
| LightningDecodeAny | 6187238 | 184.62 MB/s | 7376881 | 134004 | 2.7× |
| JSONV2 | 6308054 | 318.40 MB/s | 3173671 | 14562 | 2.7× |
| Stdlib | 16974738 | 118.32 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 626 | 876.27 MB/s | 480 | 1 | 8.4× |
| Lightning | 636 | 863.61 MB/s | 480 | 1 | 8.3× |
| LightningDestructive | 656 | 837.42 MB/s | 480 | 1 | 8.0× |
| LightningDecodeAny | 1415 | 387.30 MB/s | 1765 | 45 | 3.7× |
| Easyjson | 1801 | 304.78 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 1964 | 279.46 MB/s | 2262 | 8 | 2.7× |
| Sonic | 2012 | 272.83 MB/s | 2263 | 8 | 2.6× |
| Goccy | 2695 | 203.68 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 2738 | 200.53 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 5259 | 104.40 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 373281 | 1691.79 MB/s | 400488 | 545 | 13.0× |
| LightningArena | 408038 | 1547.68 MB/s | 448581 | 404 | 11.9× |
| Lightning | 418555 | 1508.79 MB/s | 446826 | 548 | 11.6× |
| SonicFastest | 799172 | 790.21 MB/s | 1071641 | 814 | 6.1× |
| Sonic | 806205 | 783.32 MB/s | 1071679 | 814 | 6.0× |
| Easyjson | 1064722 | 593.13 MB/s | 422504 | 936 | 4.6× |
| Goccy | 1096610 | 575.88 MB/s | 987566 | 1200 | 4.4× |
| JSONV2 | 1876270 | 336.58 MB/s | 571590 | 3144 | 2.6× |
| LightningDecodeAny | 1979401 | 235.88 MB/s | 1991443 | 29072 | 2.5× |
| Stdlib | 4866252 | 129.77 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 485172 | 1159.19 MB/s | 390799 | 426 | 10.0× |
| LightningArena | 572295 | 982.72 MB/s | 507980 | 287 | 8.5× |
| Lightning | 572515 | 982.35 MB/s | 506281 | 433 | 8.5× |
| Sonic | 1013588 | 554.87 MB/s | 1348612 | 1185 | 4.8× |
| SonicFastest | 1019123 | 551.85 MB/s | 1348666 | 1185 | 4.8× |
| Goccy | 1394926 | 403.18 MB/s | 1042825 | 1029 | 3.5× |
| Easyjson | 1759129 | 319.71 MB/s | 775153 | 1254 | 2.8× |
| LightningDecodeAny | 2189254 | 256.89 MB/s | 1990275 | 28580 | 2.2× |
| JSONV2 | 2767950 | 203.19 MB/s | 927404 | 3482 | 1.8× |
| Stdlib | 4853678 | 115.87 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 472358 | 1128.76 MB/s | 333416 | 2084 | 10.8× |
| LightningArena | 507684 | 1050.22 MB/s | 367837 | 2086 | 10.0× |
| Lightning | 508853 | 1047.80 MB/s | 367849 | 2086 | 10.0× |
| SonicFastest | 872184 | 611.31 MB/s | 984704 | 3082 | 5.8× |
| Sonic | 882702 | 604.03 MB/s | 985427 | 3082 | 5.8× |
| Easyjson | 1049740 | 507.91 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1225507 | 435.07 MB/s | 1167086 | 5409 | 4.2× |
| JSONV2 | 2304072 | 231.41 MB/s | 745422 | 13288 | 2.2× |
| LightningDecodeAny | 2800806 | 190.37 MB/s | 2659144 | 49348 | 1.8× |
| Stdlib | 5094167 | 104.66 MB/s | 798694 | 17133 | 1.0× |
