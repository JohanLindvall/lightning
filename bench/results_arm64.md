# JSON Deserialization Benchmarks

- generated 2026-10-13T08:26:47Z
- go version go1.26.8 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 81129 | 1568.80 MB/s | 49280 | 2 | 13.5× |
| Lightning | 81479 | 1562.06 MB/s | 49824 | 2 | 13.4× |
| LightningArena | 81581 | 1560.10 MB/s | 49823 | 2 | 13.4× |
| Sonic | 180613 | 704.68 MB/s | 194707 | 10 | 6.1× |
| SonicFastest | 181378 | 701.71 MB/s | 195413 | 10 | 6.0× |
| Goccy | 192169 | 662.31 MB/s | 224662 | 884 | 5.7× |
| Easyjson | 212075 | 600.14 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 419455 | 303.43 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 428000 | 221.15 MB/s | 464071 | 9706 | 2.6× |
| Stdlib | 1094899 | 116.24 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1575663 | 1428.64 MB/s | 2532849 | 1143 | 16.8× |
| Lightning | 1615246 | 1393.63 MB/s | 2532852 | 1143 | 16.4× |
| LightningArena | 2222375 | 1012.90 MB/s | 2532850 | 1143 | 11.9× |
| SonicFastest | 4576984 | 491.82 MB/s | 15247341 | 970 | 5.8× |
| Sonic | 4616202 | 487.64 MB/s | 15233780 | 970 | 5.8× |
| LightningDecodeAny | 8766636 | 256.77 MB/s | 6829000 | 223498 | 3.0× |
| Goccy | 10446345 | 215.49 MB/s | 4125277 | 56532 | 2.5× |
| Easyjson | 11165077 | 201.62 MB/s | 3099808 | 2120 | 2.4× |
| JSONV2 | 16167509 | 139.23 MB/s | 3123215 | 3083 | 1.6× |
| Stdlib | 26543327 | 84.81 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 254728 | 1061.54 MB/s | 397296 | 567 | 13.5× |
| LightningDestructive | 254918 | 1060.75 MB/s | 397296 | 567 | 13.4× |
| LightningArena | 338276 | 799.36 MB/s | 397296 | 567 | 10.1× |
| SonicFastest | 629465 | 429.58 MB/s | 466525 | 968 | 5.4× |
| Sonic | 632774 | 427.33 MB/s | 476173 | 968 | 5.4× |
| LightningDecodeAny | 1193275 | 226.61 MB/s | 845690 | 29656 | 2.9× |
| Easyjson | 1432416 | 188.77 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1437883 | 188.06 MB/s | 543646 | 8122 | 2.4× |
| JSONV2 | 2133262 | 126.76 MB/s | 348149 | 1628 | 1.6× |
| Stdlib | 3428007 | 78.88 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 925175 | 1866.89 MB/s | 765560 | 2798 | 14.3× |
| Lightning | 926541 | 1864.14 MB/s | 767597 | 2798 | 14.2× |
| LightningArena | 934166 | 1848.93 MB/s | 774939 | 2444 | 14.1× |
| Sonic | 2047568 | 843.54 MB/s | 2762258 | 4020 | 6.4× |
| SonicFastest | 2053823 | 840.97 MB/s | 2745622 | 4020 | 6.4× |
| Goccy | 2391806 | 722.13 MB/s | 2582080 | 14604 | 5.5× |
| LightningDecodeAny | 4060008 | 123.23 MB/s | 4492614 | 67881 | 3.2× |
| Easyjson | 4227088 | 408.60 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4324641 | 399.39 MB/s | 1011631 | 7594 | 3.0× |
| Stdlib | 13187187 | 130.98 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 829 | 2185.15 MB/s | 0 | 0 | 16.8× |
| Lightning | 836 | 2168.19 MB/s | 0 | 0 | 16.7× |
| LightningDestructive | 846 | 2141.04 MB/s | 0 | 0 | 16.5× |
| Easyjson | 2538 | 713.92 MB/s | 24 | 1 | 5.5× |
| Goccy | 2848 | 636.28 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6036 | 300.20 MB/s | 3794 | 40 | 2.3× |
| Sonic | 6044 | 299.82 MB/s | 3763 | 40 | 2.3× |
| JSONV2 | 7785 | 232.77 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7819 | 231.62 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13926 | 130.12 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 846 | 2140.92 MB/s | 0 | 0 | 16.4× |
| LightningArena | 853 | 2124.24 MB/s | 0 | 0 | 16.3× |
| LightningDestructive | 873 | 2076.02 MB/s | 0 | 0 | 15.9× |
| Easyjson | 2552 | 709.98 MB/s | 24 | 1 | 5.4× |
| Goccy | 2857 | 634.30 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 5957 | 304.19 MB/s | 3778 | 40 | 2.3× |
| Sonic | 5966 | 303.71 MB/s | 3780 | 40 | 2.3× |
| JSONV2 | 7832 | 231.35 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7869 | 230.14 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13901 | 130.35 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1020 | 1777.19 MB/s | 144 | 10 | 13.7× |
| LightningArena | 1025 | 1768.23 MB/s | 144 | 10 | 13.6× |
| LightningDestructive | 1063 | 1705.36 MB/s | 144 | 10 | 13.1× |
| Easyjson | 2773 | 653.39 MB/s | 144 | 10 | 5.0× |
| Goccy | 2907 | 623.41 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6155 | 294.41 MB/s | 3841 | 42 | 2.3× |
| Sonic | 6178 | 293.28 MB/s | 3883 | 42 | 2.3× |
| JSONV2 | 7969 | 227.37 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8105 | 223.44 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13953 | 129.87 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 637 | 775.17 MB/s | 160 | 1 | 8.6× |
| LightningDestructive | 638 | 773.95 MB/s | 160 | 1 | 8.6× |
| LightningDecodeAny | 1208 | 408.16 MB/s | 1040 | 25 | 4.5× |
| Sonic | 1239 | 398.87 MB/s | 977 | 6 | 4.4× |
| SonicFastest | 1241 | 398.06 MB/s | 986 | 6 | 4.4× |
| LightningArena | 1367 | 361.26 MB/s | 4120 | 2 | 4.0× |
| Easyjson | 2229 | 221.59 MB/s | 448 | 3 | 2.5× |
| Goccy | 2459 | 200.89 MB/s | 856 | 23 | 2.2× |
| JSONV2 | 3257 | 151.65 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5475 | 90.23 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 300 | 765.86 MB/s | 160 | 1 | 13.6× |
| Lightning | 303 | 758.49 MB/s | 160 | 1 | 13.5× |
| Sonic | 881 | 261.05 MB/s | 656 | 6 | 4.6× |
| SonicFastest | 885 | 259.77 MB/s | 658 | 6 | 4.6× |
| LightningDecodeAny | 1034 | 221.46 MB/s | 1040 | 25 | 4.0× |
| LightningArena | 1063 | 216.32 MB/s | 4120 | 2 | 3.8× |
| Easyjson | 1404 | 163.78 MB/s | 448 | 3 | 2.9× |
| Goccy | 1587 | 144.95 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2424 | 94.89 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4087 | 56.27 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 51112 | 1274.31 MB/s | 97220 | 98 | 10.8× |
| LightningArena | 52359 | 1243.94 MB/s | 103736 | 99 | 10.5× |
| Lightning | 52391 | 1243.19 MB/s | 103729 | 99 | 10.5× |
| SonicFastest | 101927 | 639.00 MB/s | 159419 | 75 | 5.4× |
| Sonic | 102788 | 633.65 MB/s | 160745 | 75 | 5.3× |
| Goccy | 145824 | 446.65 MB/s | 229413 | 134 | 3.8× |
| LightningDecodeAny | 177616 | 300.25 MB/s | 176707 | 3237 | 3.1× |
| JSONV2 | 226049 | 288.13 MB/s | 206655 | 607 | 2.4× |
| Stdlib | 549769 | 118.47 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2130511 | 910.80 MB/s | 2185296 | 1350 | 10.8× |
| LightningArena | 2196107 | 883.60 MB/s | 2185298 | 1350 | 10.5× |
| Lightning | 2197463 | 883.05 MB/s | 2185298 | 1350 | 10.5× |
| Goccy | 4815890 | 402.93 MB/s | 4066183 | 13510 | 4.8× |
| Sonic | 4845369 | 400.48 MB/s | 14608630 | 1407 | 4.8× |
| SonicFastest | 4874966 | 398.05 MB/s | 14608616 | 1407 | 4.7× |
| Easyjson | 7451636 | 260.41 MB/s | 3871264 | 15043 | 3.1× |
| LightningDecodeAny | 8694140 | 223.19 MB/s | 6627985 | 206416 | 2.7× |
| JSONV2 | 11379586 | 170.52 MB/s | 3237234 | 13947 | 2.0× |
| Stdlib | 23091234 | 84.04 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 874326 | 3806.17 MB/s | 351704 | 1286 | 24.0× |
| Lightning | 1325555 | 2510.52 MB/s | 2434300 | 1413 | 15.9× |
| LightningArena | 1328447 | 2505.05 MB/s | 2434383 | 1413 | 15.8× |
| SonicFastest | 2756181 | 1207.41 MB/s | 6451035 | 4248 | 7.6× |
| Sonic | 2777633 | 1198.08 MB/s | 6445580 | 4248 | 7.6× |
| LightningDecodeAny | 3262901 | 942.03 MB/s | 4825516 | 55311 | 6.4× |
| Goccy | 4579518 | 726.68 MB/s | 3948907 | 3816 | 4.6× |
| JSONV2 | 7454389 | 446.43 MB/s | 5364518 | 13243 | 2.8× |
| Stdlib | 21016393 | 158.34 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 174695 | 1261.32 MB/s | 135392 | 226 | 11.6× |
| Lightning | 174944 | 1259.53 MB/s | 135392 | 226 | 11.6× |
| LightningDestructive | 176536 | 1248.17 MB/s | 135392 | 226 | 11.5× |
| Sonic | 379566 | 580.52 MB/s | 302587 | 398 | 5.3× |
| SonicFastest | 379665 | 580.37 MB/s | 301297 | 398 | 5.3× |
| Goccy | 441664 | 498.90 MB/s | 365342 | 1067 | 4.6× |
| Easyjson | 548518 | 401.71 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 728765 | 302.36 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 852227 | 127.09 MB/s | 854738 | 11700 | 2.4× |
| Stdlib | 2023311 | 108.90 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 8521930 | 950.49 MB/s | 8109653 | 20809 | 10.3× |
| LightningDestructive | 8555719 | 946.74 MB/s | 8109652 | 20809 | 10.3× |
| LightningArena | 9682748 | 836.54 MB/s | 8109649 | 20809 | 9.1× |
| Sonic | 16899804 | 479.30 MB/s | 70902044 | 40014 | 5.2× |
| SonicFastest | 17018166 | 475.96 MB/s | 70901963 | 40014 | 5.2× |
| Goccy | 23770176 | 340.76 MB/s | 17094352 | 107148 | 3.7× |
| LightningDecodeAny | 30470387 | 170.76 MB/s | 28359832 | 746961 | 2.9× |
| Easyjson | 31011646 | 261.19 MB/s | 15059619 | 41643 | 2.8× |
| JSONV2 | 43835293 | 184.78 MB/s | 15233741 | 78972 | 2.0× |
| Stdlib | 88050836 | 91.99 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 3355239 | 889.20 MB/s | 3780458 | 1514 | 13.9× |
| LightningDestructive | 3520021 | 847.57 MB/s | 3758857 | 29356 | 13.2× |
| Lightning | 3650092 | 817.37 MB/s | 3758857 | 29356 | 12.8× |
| Sonic | 8755527 | 340.75 MB/s | 26529965 | 56760 | 5.3× |
| SonicFastest | 8784008 | 339.65 MB/s | 26545096 | 56760 | 5.3× |
| LightningDecodeAny | 14704534 | 124.74 MB/s | 18225372 | 350883 | 3.2× |
| Goccy | 16582815 | 179.91 MB/s | 10664952 | 273651 | 2.8× |
| Easyjson | 16698430 | 178.67 MB/s | 9479440 | 30115 | 2.8× |
| JSONV2 | 24716293 | 120.71 MB/s | 9257144 | 86278 | 1.9× |
| Stdlib | 46607120 | 64.01 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 697158 | 1037.92 MB/s | 907600 | 3618 | 16.3× |
| LightningArena | 709591 | 1019.74 MB/s | 916259 | 37 | 16.1× |
| Lightning | 740305 | 977.43 MB/s | 907599 | 3618 | 15.4× |
| Sonic | 1786958 | 404.93 MB/s | 3187310 | 7226 | 6.4× |
| SonicFastest | 1794878 | 403.15 MB/s | 3200847 | 7226 | 6.3× |
| LightningDecodeAny | 3962858 | 164.17 MB/s | 5691593 | 76540 | 2.9× |
| Easyjson | 4175010 | 173.32 MB/s | 2847905 | 3698 | 2.7× |
| Goccy | 4848724 | 149.23 MB/s | 2758654 | 80271 | 2.3× |
| JSONV2 | 5640122 | 128.29 MB/s | 2704636 | 7318 | 2.0× |
| Stdlib | 11393182 | 63.51 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1013517 | 1556.32 MB/s | 907600 | 3618 | 15.3× |
| LightningArena | 1020514 | 1545.65 MB/s | 916257 | 37 | 15.2× |
| Lightning | 1060776 | 1486.98 MB/s | 907594 | 3618 | 14.6× |
| Sonic | 2279735 | 691.90 MB/s | 5792411 | 7226 | 6.8× |
| SonicFastest | 2285844 | 690.05 MB/s | 5795266 | 7226 | 6.8× |
| LightningDecodeAny | 3469956 | 217.12 MB/s | 5691593 | 76540 | 4.5× |
| Easyjson | 5595408 | 281.90 MB/s | 2847908 | 3698 | 2.8× |
| Goccy | 5609363 | 281.20 MB/s | 3582557 | 80267 | 2.8× |
| JSONV2 | 6310160 | 249.97 MB/s | 2704595 | 7318 | 2.5× |
| Stdlib | 15540099 | 101.50 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 71113 | 2111.07 MB/s | 81920 | 1 | 26.2× |
| LightningArena | 72424 | 2072.84 MB/s | 81920 | 1 | 25.8× |
| Lightning | 72548 | 2069.31 MB/s | 81920 | 1 | 25.7× |
| Sonic | 271979 | 551.97 MB/s | 249139 | 6 | 6.9× |
| SonicFastest | 273386 | 549.13 MB/s | 251416 | 6 | 6.8× |
| LightningDecodeAny | 438996 | 341.96 MB/s | 745507 | 10015 | 4.3× |
| Goccy | 868680 | 172.82 MB/s | 326040 | 10005 | 2.1× |
| JSONV2 | 1074970 | 139.65 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1865934 | 80.46 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 26812 | 1048.66 MB/s | 29088 | 101 | 11.2× |
| LightningArena | 27194 | 1033.93 MB/s | 29237 | 101 | 11.1× |
| Lightning | 27215 | 1033.16 MB/s | 29236 | 101 | 11.0× |
| Sonic | 63335 | 443.94 MB/s | 46699 | 103 | 4.7× |
| SonicFastest | 63602 | 442.08 MB/s | 46869 | 103 | 4.7× |
| Easyjson | 68160 | 412.51 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72073 | 390.12 MB/s | 59205 | 188 | 4.2× |
| JSONV2 | 134378 | 209.24 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 145414 | 193.36 MB/s | 133690 | 2639 | 2.1× |
| Stdlib | 300657 | 93.52 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1501 | 1551.03 MB/s | 32 | 1 | 14.9× |
| Lightning | 1503 | 1548.69 MB/s | 32 | 1 | 14.9× |
| LightningDestructive | 1571 | 1481.61 MB/s | 32 | 1 | 14.3× |
| Goccy | 4142 | 562.10 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4203 | 553.84 MB/s | 192 | 2 | 5.3× |
| Sonic | 5071 | 459.04 MB/s | 4218 | 6 | 4.4× |
| SonicFastest | 5108 | 455.72 MB/s | 4192 | 6 | 4.4× |
| JSONV2 | 8452 | 275.43 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9647 | 174.67 MB/s | 9936 | 194 | 2.3× |
| Stdlib | 22413 | 103.87 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 180 | 1052.90 MB/s | 0 | 0 | 13.2× |
| LightningArena | 180 | 1049.84 MB/s | 0 | 0 | 13.2× |
| LightningDestructive | 182 | 1039.17 MB/s | 0 | 0 | 13.1× |
| Goccy | 374 | 505.90 MB/s | 304 | 2 | 6.4× |
| Easyjson | 483 | 391.53 MB/s | 0 | 0 | 4.9× |
| Sonic | 790 | 239.26 MB/s | 500 | 4 | 3.0× |
| SonicFastest | 793 | 238.39 MB/s | 502 | 4 | 3.0× |
| JSONV2 | 1024 | 184.60 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1193 | 112.32 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2375 | 79.57 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1093 | 2004.39 MB/s | 0 | 0 | 14.4× |
| LightningArena | 1094 | 2002.70 MB/s | 0 | 0 | 14.4× |
| LightningDestructive | 1118 | 1959.71 MB/s | 0 | 0 | 14.1× |
| Goccy | 3190 | 686.83 MB/s | 2864 | 4 | 4.9× |
| Easyjson | 3191 | 686.59 MB/s | 24 | 1 | 4.9× |
| SonicFastest | 6394 | 342.69 MB/s | 3992 | 40 | 2.5× |
| Sonic | 6446 | 339.91 MB/s | 4008 | 40 | 2.4× |
| LightningDecodeAny | 7875 | 229.96 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 8065 | 271.66 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15778 | 138.86 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 544069 | 938.26 MB/s | 318400 | 1005 | 10.9× |
| LightningArena | 544665 | 937.23 MB/s | 318400 | 1005 | 10.9× |
| LightningDestructive | 545446 | 935.89 MB/s | 318400 | 1005 | 10.9× |
| Goccy | 1150117 | 443.85 MB/s | 1137808 | 5006 | 5.2× |
| SonicFastest | 1154605 | 442.12 MB/s | 884472 | 2006 | 5.2× |
| Sonic | 1155286 | 441.86 MB/s | 884052 | 2006 | 5.2× |
| Easyjson | 1540732 | 331.32 MB/s | 863778 | 3012 | 3.9× |
| LightningDecodeAny | 3248289 | 142.06 MB/s | 2742392 | 64017 | 1.8× |
| JSONV2 | 3286785 | 155.31 MB/s | 1076006 | 12646 | 1.8× |
| Stdlib | 5956714 | 85.70 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 461 | 42956.18 MB/s | 0 | 0 | 235.3× |
| LightningArena | 461 | 42891.23 MB/s | 0 | 0 | 235.0× |
| LightningDestructive | 474 | 41718.46 MB/s | 0 | 0 | 228.6× |
| Goccy | 20167 | 981.24 MB/s | 20491 | 2 | 5.4× |
| Sonic | 27041 | 731.82 MB/s | 22261 | 4 | 4.0× |
| SonicFastest | 27140 | 729.16 MB/s | 22497 | 4 | 4.0× |
| JSONV2 | 29748 | 665.22 MB/s | 8 | 1 | 3.6× |
| Easyjson | 81826 | 241.84 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 86282 | 229.34 MB/s | 116608 | 2014 | 1.3× |
| Stdlib | 108406 | 182.55 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1741 | 10407.68 MB/s | 0 | 0 | 58.9× |
| LightningArena | 1819 | 9962.75 MB/s | 405 | 0 | 56.4× |
| Lightning | 1821 | 9952.59 MB/s | 405 | 0 | 56.3× |
| Easyjson | 3968 | 4567.60 MB/s | 432 | 2 | 25.9× |
| SonicFastest | 10158 | 1784.23 MB/s | 23630 | 6 | 10.1× |
| Sonic | 10206 | 1775.77 MB/s | 23508 | 6 | 10.1× |
| Goccy | 15780 | 1148.57 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 15963 | 1120.25 MB/s | 29103 | 189 | 6.4× |
| JSONV2 | 46810 | 387.18 MB/s | 16498 | 50 | 2.2× |
| Stdlib | 102606 | 176.64 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1999384 | 1004.56 MB/s | 3089564 | 6821 | 9.3× |
| Lightning | 2052349 | 978.63 MB/s | 3096412 | 6822 | 9.0× |
| LightningArena | 2074154 | 968.34 MB/s | 3100166 | 6700 | 8.9× |
| Goccy | 4295453 | 467.59 MB/s | 5411620 | 15830 | 4.3× |
| Sonic | 4378843 | 458.68 MB/s | 10908463 | 13683 | 4.2× |
| SonicFastest | 4397846 | 456.70 MB/s | 10933814 | 13683 | 4.2× |
| Easyjson | 4923227 | 407.96 MB/s | 2981483 | 7439 | 3.8× |
| JSONV2 | 7002876 | 286.81 MB/s | 3173684 | 14563 | 2.6× |
| LightningDecodeAny | 7054213 | 161.93 MB/s | 7376871 | 134004 | 2.6× |
| Stdlib | 18510772 | 108.50 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 834 | 658.41 MB/s | 480 | 1 | 6.8× |
| LightningArena | 836 | 656.34 MB/s | 480 | 1 | 6.7× |
| LightningDestructive | 850 | 645.98 MB/s | 480 | 1 | 6.6× |
| LightningDecodeAny | 1613 | 339.70 MB/s | 1765 | 45 | 3.5× |
| Easyjson | 2205 | 249.01 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2726 | 201.39 MB/s | 2005 | 26 | 2.1× |
| SonicFastest | 2742 | 200.23 MB/s | 2020 | 26 | 2.1× |
| Goccy | 3079 | 178.31 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3386 | 162.16 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5636 | 97.41 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 402886 | 1567.48 MB/s | 400488 | 545 | 13.3× |
| Lightning | 440463 | 1433.75 MB/s | 447156 | 548 | 12.1× |
| LightningArena | 444678 | 1420.16 MB/s | 449135 | 404 | 12.0× |
| Sonic | 1024964 | 616.13 MB/s | 1006745 | 1102 | 5.2× |
| SonicFastest | 1036737 | 609.14 MB/s | 1026890 | 1102 | 5.2× |
| Easyjson | 1149067 | 549.59 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1171007 | 539.29 MB/s | 986960 | 1201 | 4.6× |
| JSONV2 | 2178014 | 289.95 MB/s | 571612 | 3144 | 2.5× |
| LightningDecodeAny | 2270063 | 205.68 MB/s | 1992682 | 29072 | 2.4× |
| Stdlib | 5346249 | 118.12 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 568911 | 988.57 MB/s | 391211 | 426 | 9.2× |
| Lightning | 640201 | 878.49 MB/s | 506423 | 433 | 8.2× |
| LightningArena | 641259 | 877.04 MB/s | 508217 | 287 | 8.2× |
| Sonic | 1034461 | 543.67 MB/s | 958745 | 1476 | 5.1× |
| SonicFastest | 1044404 | 538.50 MB/s | 970918 | 1476 | 5.0× |
| Goccy | 1339533 | 419.85 MB/s | 1041072 | 1030 | 3.9× |
| Easyjson | 1759177 | 319.70 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2309221 | 243.55 MB/s | 1991711 | 28581 | 2.3× |
| JSONV2 | 2792712 | 201.38 MB/s | 927439 | 3482 | 1.9× |
| Stdlib | 5243221 | 107.26 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 547963 | 973.02 MB/s | 333416 | 2084 | 9.9× |
| Lightning | 559777 | 952.48 MB/s | 367506 | 2086 | 9.7× |
| LightningArena | 561263 | 949.96 MB/s | 367530 | 2086 | 9.7× |
| Easyjson | 1107181 | 481.56 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1135207 | 469.67 MB/s | 1029845 | 4351 | 4.8× |
| Sonic | 1136295 | 469.22 MB/s | 1021031 | 4351 | 4.8× |
| Goccy | 1299189 | 410.39 MB/s | 1167212 | 5409 | 4.2× |
| JSONV2 | 2515985 | 211.92 MB/s | 745449 | 13288 | 2.2× |
| LightningDecodeAny | 3197022 | 166.77 MB/s | 2659429 | 49349 | 1.7× |
| Stdlib | 5416759 | 98.43 MB/s | 798692 | 17133 | 1.0× |
