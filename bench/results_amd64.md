# JSON Deserialization Benchmarks

- generated 2026-09-08T04:43:21Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 82278 | 1546.90 MB/s | 49760 | 3 | 16.1× |
| Lightning | 82371 | 1545.15 MB/s | 49760 | 3 | 16.1× |
| LightningDestructive | 86788 | 1466.50 MB/s | 49280 | 2 | 15.3× |
| Sonic | 197465 | 644.55 MB/s | 213948 | 15 | 6.7× |
| SonicFastest | 197781 | 643.51 MB/s | 214100 | 15 | 6.7× |
| Easyjson | 248862 | 511.43 MB/s | 122864 | 14 | 5.3× |
| Goccy | 260149 | 489.24 MB/s | 224714 | 884 | 5.1× |
| JSONV2 | 453309 | 280.77 MB/s | 195128 | 1805 | 2.9× |
| LightningDecodeAny | 456560 | 207.32 MB/s | 463411 | 9708 | 2.9× |
| Stdlib | 1327925 | 95.85 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3226333 | 697.71 MB/s | 2532848 | 1143 | 9.9× |
| LightningArena | 3277553 | 686.81 MB/s | 2532849 | 1143 | 9.8× |
| Lightning | 3297029 | 682.75 MB/s | 2532850 | 1143 | 9.7× |
| SonicFastest | 5153906 | 436.77 MB/s | 4874394 | 2584 | 6.2× |
| Sonic | 5373545 | 418.91 MB/s | 4868170 | 2584 | 5.9× |
| LightningDecodeAny | 12635716 | 178.15 MB/s | 19380209 | 223896 | 2.5× |
| Goccy | 12836879 | 175.36 MB/s | 4206177 | 56536 | 2.5× |
| Easyjson | 13198004 | 170.56 MB/s | 3099809 | 2120 | 2.4× |
| JSONV2 | 17478453 | 128.79 MB/s | 3123189 | 3083 | 1.8× |
| Stdlib | 31968153 | 70.42 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 475301 | 568.91 MB/s | 397297 | 567 | 8.7× |
| Lightning | 475818 | 568.29 MB/s | 397297 | 567 | 8.7× |
| LightningDestructive | 485512 | 556.94 MB/s | 397296 | 567 | 8.5× |
| Sonic | 755529 | 357.90 MB/s | 640492 | 1147 | 5.5× |
| SonicFastest | 755734 | 357.80 MB/s | 640899 | 1147 | 5.5× |
| Easyjson | 1723324 | 156.91 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1763640 | 153.32 MB/s | 541259 | 8122 | 2.3× |
| LightningDecodeAny | 2085808 | 129.64 MB/s | 2543877 | 29687 | 2.0× |
| JSONV2 | 2304656 | 117.33 MB/s | 348161 | 1628 | 1.8× |
| Stdlib | 4140705 | 65.30 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1002483 | 1722.93 MB/s | 765560 | 2798 | 16.9× |
| Lightning | 1012452 | 1705.96 MB/s | 765601 | 2799 | 16.7× |
| LightningArena | 1013402 | 1704.36 MB/s | 772704 | 2445 | 16.7× |
| SonicFastest | 2157241 | 800.65 MB/s | 2695173 | 5547 | 7.9× |
| Sonic | 2166238 | 797.33 MB/s | 2695286 | 5547 | 7.8× |
| Goccy | 2503714 | 689.86 MB/s | 2581101 | 14603 | 6.8× |
| LightningDecodeAny | 4141678 | 120.80 MB/s | 4953691 | 76576 | 4.1× |
| Easyjson | 4208588 | 410.40 MB/s | 972032 | 5389 | 4.0× |
| JSONV2 | 4812653 | 358.89 MB/s | 1011615 | 7594 | 3.5× |
| Stdlib | 16952856 | 101.88 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 765 | 2367.68 MB/s | 0 | 0 | 21.5× |
| LightningArena | 772 | 2346.47 MB/s | 0 | 0 | 21.3× |
| LightningDestructive | 836 | 2167.59 MB/s | 0 | 0 | 19.7× |
| Easyjson | 3039 | 596.33 MB/s | 24 | 1 | 5.4× |
| Goccy | 3513 | 515.81 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6112 | 296.49 MB/s | 3346 | 38 | 2.7× |
| Sonic | 6348 | 285.46 MB/s | 3345 | 38 | 2.6× |
| JSONV2 | 8788 | 206.19 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9077 | 199.52 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16473 | 110.00 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 844 | 2146.08 MB/s | 0 | 0 | 19.5× |
| LightningArena | 855 | 2119.64 MB/s | 0 | 0 | 19.3× |
| LightningDestructive | 901 | 2010.14 MB/s | 0 | 0 | 18.3× |
| Easyjson | 3022 | 599.56 MB/s | 24 | 1 | 5.5× |
| Goccy | 3491 | 518.98 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6154 | 294.43 MB/s | 3345 | 38 | 2.7× |
| Sonic | 6379 | 284.06 MB/s | 3345 | 38 | 2.6× |
| JSONV2 | 8383 | 216.15 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9166 | 197.57 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16488 | 109.90 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1105 | 1639.59 MB/s | 144 | 10 | 15.0× |
| Lightning | 1113 | 1628.22 MB/s | 144 | 10 | 14.9× |
| LightningDestructive | 1173 | 1544.96 MB/s | 144 | 10 | 14.1× |
| Easyjson | 3130 | 578.92 MB/s | 144 | 10 | 5.3× |
| Goccy | 3281 | 552.20 MB/s | 2600 | 5 | 5.1× |
| SonicFastest | 6307 | 287.30 MB/s | 3365 | 40 | 2.6× |
| Sonic | 6539 | 277.12 MB/s | 3367 | 40 | 2.5× |
| JSONV2 | 8407 | 215.52 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 9177 | 197.34 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16585 | 109.25 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 668 | 739.75 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 679 | 727.53 MB/s | 160 | 1 | 9.7× |
| Sonic | 1273 | 388.16 MB/s | 1076 | 8 | 5.2× |
| SonicFastest | 1277 | 386.97 MB/s | 1076 | 8 | 5.2× |
| LightningDecodeAny | 1429 | 344.94 MB/s | 1296 | 26 | 4.6× |
| LightningArena | 1585 | 311.70 MB/s | 4120 | 2 | 4.2× |
| Easyjson | 2724 | 181.38 MB/s | 448 | 3 | 2.4× |
| Goccy | 2814 | 175.52 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3483 | 141.83 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6594 | 74.92 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 416 | 552.70 MB/s | 160 | 1 | 11.3× |
| Lightning | 417 | 551.29 MB/s | 160 | 1 | 11.3× |
| SonicFastest | 943 | 243.86 MB/s | 801 | 8 | 5.0× |
| Sonic | 946 | 243.02 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1253 | 182.82 MB/s | 1296 | 26 | 3.8× |
| LightningArena | 1360 | 169.07 MB/s | 4120 | 2 | 3.5× |
| Easyjson | 1665 | 138.16 MB/s | 448 | 3 | 2.8× |
| Goccy | 1911 | 120.35 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2568 | 89.55 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4713 | 48.80 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 61220 | 1063.91 MB/s | 103441 | 103 | 11.0× |
| Lightning | 63045 | 1033.10 MB/s | 103441 | 103 | 10.7× |
| LightningDestructive | 65491 | 994.52 MB/s | 97220 | 98 | 10.3× |
| Sonic | 154807 | 420.73 MB/s | 236127 | 65 | 4.4× |
| SonicFastest | 170860 | 381.20 MB/s | 236735 | 65 | 4.0× |
| LightningDecodeAny | 193984 | 274.92 MB/s | 180049 | 3245 | 3.5× |
| Goccy | 218345 | 298.30 MB/s | 230034 | 134 | 3.1× |
| JSONV2 | 290043 | 224.56 MB/s | 206664 | 607 | 2.3× |
| Stdlib | 676103 | 96.33 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2183070 | 888.87 MB/s | 2864592 | 1380 | 12.5× |
| LightningArena | 2287433 | 848.32 MB/s | 2864593 | 1380 | 12.0× |
| Lightning | 2293155 | 846.20 MB/s | 2864593 | 1380 | 11.9× |
| SonicFastest | 4797272 | 404.49 MB/s | 4879207 | 1736 | 5.7× |
| Sonic | 4802749 | 404.03 MB/s | 4881185 | 1736 | 5.7× |
| Goccy | 4903533 | 395.73 MB/s | 4063734 | 13509 | 5.6× |
| Easyjson | 8493940 | 228.45 MB/s | 3871265 | 15043 | 3.2× |
| LightningDecodeAny | 9732554 | 199.38 MB/s | 7063041 | 218633 | 2.8× |
| JSONV2 | 12081278 | 160.62 MB/s | 3237188 | 13947 | 2.3× |
| Stdlib | 27379259 | 70.87 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 903804 | 3682.03 MB/s | 351704 | 1286 | 25.8× |
| Lightning | 1415673 | 2350.71 MB/s | 2488906 | 2995 | 16.5× |
| LightningArena | 1420606 | 2342.54 MB/s | 2488907 | 2995 | 16.4× |
| SonicFastest | 2025882 | 1642.66 MB/s | 5896674 | 4263 | 11.5× |
| Sonic | 2045707 | 1626.74 MB/s | 5896572 | 4263 | 11.4× |
| LightningDecodeAny | 3317605 | 926.50 MB/s | 4876914 | 56892 | 7.0× |
| Goccy | 4660767 | 714.01 MB/s | 3948914 | 3817 | 5.0× |
| JSONV2 | 7291453 | 456.40 MB/s | 5364507 | 13243 | 3.2× |
| Stdlib | 23320675 | 142.70 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 180163 | 1223.04 MB/s | 135872 | 226 | 13.3× |
| Lightning | 182281 | 1208.83 MB/s | 135872 | 226 | 13.2× |
| LightningDestructive | 189433 | 1163.19 MB/s | 135872 | 226 | 12.7× |
| SonicFastest | 492179 | 447.69 MB/s | 351249 | 262 | 4.9× |
| Goccy | 493127 | 446.83 MB/s | 364621 | 1066 | 4.9× |
| Sonic | 494101 | 445.95 MB/s | 351191 | 262 | 4.9× |
| Easyjson | 648775 | 339.63 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 716211 | 307.66 MB/s | 129747 | 470 | 3.4× |
| LightningDecodeAny | 1002486 | 108.04 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2400385 | 91.80 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 10133277 | 799.35 MB/s | 11845073 | 20816 | 10.5× |
| LightningArena | 10430876 | 776.54 MB/s | 11845072 | 20816 | 10.2× |
| Lightning | 10486840 | 772.40 MB/s | 11845077 | 20816 | 10.2× |
| SonicFastest | 18061719 | 448.46 MB/s | 19866048 | 41640 | 5.9× |
| Sonic | 18145260 | 446.40 MB/s | 19866302 | 41640 | 5.9× |
| Goccy | 26531984 | 305.29 MB/s | 18901144 | 107155 | 4.0× |
| Easyjson | 35893248 | 225.67 MB/s | 15059624 | 41643 | 3.0× |
| LightningDecodeAny | 38412205 | 135.45 MB/s | 46279353 | 747112 | 2.8× |
| JSONV2 | 48468418 | 167.12 MB/s | 15233722 | 78972 | 2.2× |
| Stdlib | 106821059 | 75.83 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4798275 | 621.78 MB/s | 3780456 | 1514 | 11.8× |
| LightningDestructive | 4956771 | 601.90 MB/s | 3758856 | 29356 | 11.4× |
| Lightning | 5131741 | 581.37 MB/s | 3758857 | 29356 | 11.0× |
| Sonic | 8954685 | 333.17 MB/s | 9131722 | 57804 | 6.3× |
| SonicFastest | 8957025 | 333.09 MB/s | 9131831 | 57804 | 6.3× |
| LightningDecodeAny | 17991262 | 101.95 MB/s | 23982579 | 351152 | 3.1× |
| Goccy | 18635853 | 160.09 MB/s | 9955645 | 273624 | 3.0× |
| Easyjson | 19343148 | 154.24 MB/s | 9479441 | 30115 | 2.9× |
| JSONV2 | 27213791 | 109.63 MB/s | 9257061 | 86278 | 2.1× |
| Stdlib | 56466556 | 52.84 MB/s | 9258083 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1153435 | 627.34 MB/s | 916256 | 37 | 11.8× |
| LightningDestructive | 1158471 | 624.61 MB/s | 907601 | 3618 | 11.8× |
| Lightning | 1241684 | 582.75 MB/s | 907596 | 3618 | 11.0× |
| SonicFastest | 2084633 | 347.11 MB/s | 2372233 | 3683 | 6.6× |
| Sonic | 2094878 | 345.41 MB/s | 2371132 | 3683 | 6.5× |
| LightningDecodeAny | 5244260 | 124.05 MB/s | 6500460 | 76546 | 2.6× |
| Easyjson | 5383620 | 134.41 MB/s | 2847907 | 3698 | 2.5× |
| Goccy | 5494093 | 131.70 MB/s | 2747210 | 80269 | 2.5× |
| JSONV2 | 6367653 | 113.64 MB/s | 2704700 | 7318 | 2.1× |
| Stdlib | 13660013 | 52.97 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1570508 | 1004.36 MB/s | 916256 | 37 | 12.2× |
| LightningDestructive | 1600929 | 985.27 MB/s | 907600 | 3618 | 11.9× |
| Lightning | 1617027 | 975.47 MB/s | 907595 | 3618 | 11.8× |
| Sonic | 2414359 | 653.32 MB/s | 3222792 | 3683 | 7.9× |
| SonicFastest | 2422684 | 651.08 MB/s | 3222754 | 3683 | 7.9× |
| LightningDecodeAny | 4673518 | 161.21 MB/s | 6500454 | 76546 | 4.1× |
| Easyjson | 6298022 | 250.45 MB/s | 2847904 | 3698 | 3.0× |
| Goccy | 6685382 | 235.94 MB/s | 3494958 | 80263 | 2.9× |
| JSONV2 | 6900503 | 228.59 MB/s | 2704553 | 7318 | 2.8× |
| Stdlib | 19097145 | 82.60 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 226114 | 663.93 MB/s | 81920 | 1 | 9.4× |
| LightningArena | 227589 | 659.63 MB/s | 81920 | 1 | 9.3× |
| LightningDestructive | 234144 | 641.16 MB/s | 81920 | 1 | 9.1× |
| Sonic | 396952 | 378.19 MB/s | 407560 | 16 | 5.3× |
| SonicFastest | 418478 | 358.74 MB/s | 407817 | 16 | 5.1× |
| LightningDecodeAny | 568706 | 263.97 MB/s | 745766 | 10016 | 3.7× |
| Goccy | 981934 | 152.89 MB/s | 328643 | 10005 | 2.2× |
| JSONV2 | 1152639 | 130.24 MB/s | 357727 | 20 | 1.8× |
| Stdlib | 2121070 | 70.78 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 29112 | 965.81 MB/s | 29216 | 103 | 11.9× |
| Lightning | 29378 | 957.08 MB/s | 29216 | 103 | 11.8× |
| LightningDestructive | 30945 | 908.62 MB/s | 29088 | 101 | 11.2× |
| SonicFastest | 60551 | 464.35 MB/s | 59424 | 83 | 5.7× |
| Sonic | 60680 | 463.36 MB/s | 59437 | 83 | 5.7× |
| Easyjson | 81053 | 346.90 MB/s | 32304 | 138 | 4.3× |
| Goccy | 83772 | 335.64 MB/s | 59282 | 188 | 4.1× |
| JSONV2 | 139905 | 200.97 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 162329 | 173.21 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 345891 | 81.29 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1561 | 1491.24 MB/s | 32 | 1 | 16.7× |
| Lightning | 1563 | 1489.86 MB/s | 32 | 1 | 16.6× |
| LightningDestructive | 1710 | 1361.28 MB/s | 32 | 1 | 15.2× |
| SonicFastest | 4801 | 484.86 MB/s | 3705 | 4 | 5.4× |
| Sonic | 4802 | 484.79 MB/s | 3706 | 4 | 5.4× |
| Goccy | 5006 | 465.06 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 5580 | 417.17 MB/s | 192 | 2 | 4.7× |
| JSONV2 | 8376 | 277.92 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10515 | 160.25 MB/s | 10200 | 195 | 2.5× |
| Stdlib | 26014 | 89.49 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 162 | 1168.51 MB/s | 0 | 0 | 17.1× |
| LightningDestructive | 167 | 1134.22 MB/s | 0 | 0 | 16.6× |
| LightningArena | 168 | 1125.17 MB/s | 0 | 0 | 16.5× |
| Goccy | 471 | 401.22 MB/s | 304 | 2 | 5.9× |
| Easyjson | 590 | 320.33 MB/s | 0 | 0 | 4.7× |
| Sonic | 644 | 293.50 MB/s | 341 | 3 | 4.3× |
| SonicFastest | 646 | 292.56 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1090 | 173.38 MB/s | 112 | 1 | 2.5× |
| LightningDecodeAny | 1351 | 99.22 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2765 | 68.35 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1121 | 1954.22 MB/s | 0 | 0 | 17.0× |
| Lightning | 1123 | 1951.83 MB/s | 0 | 0 | 17.0× |
| LightningDestructive | 1201 | 1824.56 MB/s | 0 | 0 | 15.9× |
| Easyjson | 3925 | 558.15 MB/s | 24 | 1 | 4.9× |
| Goccy | 3965 | 552.64 MB/s | 2864 | 4 | 4.8× |
| SonicFastest | 6674 | 328.30 MB/s | 3601 | 38 | 2.9× |
| Sonic | 6847 | 320.00 MB/s | 3599 | 38 | 2.8× |
| JSONV2 | 8644 | 253.48 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9181 | 197.25 MB/s | 7552 | 158 | 2.1× |
| Stdlib | 19072 | 114.88 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 510595 | 999.77 MB/s | 457537 | 1009 | 13.8× |
| Lightning | 539875 | 945.55 MB/s | 457537 | 1009 | 13.1× |
| LightningArena | 548197 | 931.19 MB/s | 457537 | 1009 | 12.9× |
| Goccy | 1233217 | 413.94 MB/s | 1137743 | 5006 | 5.7× |
| Sonic | 1282936 | 397.90 MB/s | 1307402 | 2014 | 5.5× |
| SonicFastest | 1292861 | 394.84 MB/s | 1309632 | 2014 | 5.5× |
| Easyjson | 1675009 | 304.76 MB/s | 863781 | 3012 | 4.2× |
| JSONV2 | 3522898 | 144.90 MB/s | 1075955 | 12645 | 2.0× |
| LightningDecodeAny | 3552764 | 129.89 MB/s | 2950650 | 64018 | 2.0× |
| Stdlib | 7059502 | 72.31 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 540 | 36635.89 MB/s | 0 | 0 | 305.9× |
| Lightning | 540 | 36626.75 MB/s | 0 | 0 | 305.8× |
| LightningDestructive | 783 | 25263.07 MB/s | 0 | 0 | 211.0× |
| SonicFastest | 6927 | 2856.67 MB/s | 21067 | 3 | 23.9× |
| Goccy | 24803 | 797.84 MB/s | 20492 | 2 | 6.7× |
| Sonic | 31805 | 622.19 MB/s | 20582 | 3 | 5.2× |
| JSONV2 | 34378 | 575.63 MB/s | 8 | 1 | 4.8× |
| LightningDecodeAny | 98195 | 201.52 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 98586 | 200.73 MB/s | 0 | 0 | 1.7× |
| Stdlib | 165245 | 119.76 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1779 | 10186.16 MB/s | 432 | 2 | 69.9× |
| LightningArena | 1798 | 10078.99 MB/s | 432 | 2 | 69.1× |
| LightningDestructive | 2105 | 8608.25 MB/s | 0 | 0 | 59.0× |
| Easyjson | 5012 | 3616.24 MB/s | 432 | 2 | 24.8× |
| Sonic | 8690 | 2085.54 MB/s | 20460 | 5 | 14.3× |
| SonicFastest | 8742 | 2073.12 MB/s | 20461 | 5 | 14.2× |
| LightningDecodeAny | 19101 | 936.19 MB/s | 29088 | 191 | 6.5× |
| Goccy | 25258 | 717.56 MB/s | 19460 | 2 | 4.9× |
| JSONV2 | 51094 | 354.72 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 124283 | 145.83 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2101031 | 955.96 MB/s | 3089565 | 6821 | 10.3× |
| LightningArena | 2193669 | 915.59 MB/s | 3094394 | 6704 | 9.9× |
| Lightning | 2205340 | 910.74 MB/s | 3091277 | 6827 | 9.8× |
| SonicFastest | 4146127 | 484.43 MB/s | 5153905 | 7085 | 5.2× |
| Sonic | 4214723 | 476.54 MB/s | 5153629 | 7085 | 5.1× |
| Goccy | 4737790 | 423.93 MB/s | 5411613 | 15832 | 4.6× |
| Easyjson | 5723705 | 350.91 MB/s | 2981483 | 7439 | 3.8× |
| LightningDecodeAny | 6964013 | 164.03 MB/s | 8503512 | 134008 | 3.1× |
| JSONV2 | 7830768 | 256.49 MB/s | 3173678 | 14563 | 2.8× |
| Stdlib | 21631924 | 92.85 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 820 | 669.80 MB/s | 480 | 1 | 8.2× |
| LightningArena | 826 | 664.48 MB/s | 480 | 1 | 8.2× |
| LightningDestructive | 833 | 658.71 MB/s | 480 | 1 | 8.1× |
| LightningDecodeAny | 1826 | 300.16 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 2266 | 242.24 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2404 | 228.34 MB/s | 2262 | 8 | 2.8× |
| Sonic | 2413 | 227.48 MB/s | 2262 | 8 | 2.8× |
| Goccy | 3254 | 168.74 MB/s | 2129 | 43 | 2.1× |
| JSONV2 | 3539 | 155.12 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6743 | 81.42 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 434752 | 1452.59 MB/s | 402728 | 545 | 14.4× |
| LightningArena | 511388 | 1234.90 MB/s | 453041 | 713 | 12.3× |
| Lightning | 513748 | 1229.23 MB/s | 451257 | 857 | 12.2× |
| SonicFastest | 1047432 | 602.92 MB/s | 1068013 | 814 | 6.0× |
| Sonic | 1048197 | 602.48 MB/s | 1068205 | 814 | 6.0× |
| Goccy | 1378933 | 457.97 MB/s | 988117 | 1200 | 4.5× |
| Easyjson | 1394504 | 452.86 MB/s | 422504 | 936 | 4.5× |
| JSONV2 | 2294104 | 275.28 MB/s | 571593 | 3144 | 2.7× |
| LightningDecodeAny | 2571368 | 181.58 MB/s | 2076504 | 30126 | 2.4× |
| Stdlib | 6273752 | 100.66 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 631203 | 891.01 MB/s | 546569 | 429 | 9.4× |
| LightningArena | 815987 | 689.24 MB/s | 771688 | 1089 | 7.3× |
| Lightning | 818998 | 686.70 MB/s | 769936 | 1235 | 7.3× |
| SonicFastest | 1290003 | 435.97 MB/s | 1347730 | 1185 | 4.6× |
| Sonic | 1291369 | 435.51 MB/s | 1347786 | 1185 | 4.6× |
| Goccy | 1583193 | 355.24 MB/s | 1035782 | 1028 | 3.8× |
| Easyjson | 2158446 | 260.56 MB/s | 775153 | 1254 | 2.8× |
| LightningDecodeAny | 2956813 | 190.21 MB/s | 2180440 | 30126 | 2.0× |
| JSONV2 | 3155397 | 178.24 MB/s | 927407 | 3482 | 1.9× |
| Stdlib | 5956225 | 94.42 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 565171 | 943.39 MB/s | 333416 | 2084 | 11.4× |
| Lightning | 632049 | 843.57 MB/s | 368224 | 2293 | 10.2× |
| LightningArena | 644147 | 827.73 MB/s | 368224 | 2293 | 10.0× |
| SonicFastest | 1128298 | 472.55 MB/s | 980927 | 3082 | 5.7× |
| Sonic | 1131292 | 471.30 MB/s | 981928 | 3082 | 5.7× |
| Easyjson | 1282700 | 415.67 MB/s | 428362 | 3273 | 5.0× |
| Goccy | 1509045 | 353.32 MB/s | 1167081 | 5408 | 4.3× |
| JSONV2 | 2913975 | 182.97 MB/s | 745424 | 13288 | 2.2× |
| LightningDecodeAny | 3499889 | 152.34 MB/s | 2992875 | 50076 | 1.8× |
| Stdlib | 6421554 | 83.03 MB/s | 798692 | 17133 | 1.0× |
