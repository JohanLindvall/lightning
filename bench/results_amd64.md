# JSON Deserialization Benchmarks

- generated 2026-08-08T12:49:45Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 100334 | 1268.51 MB/s | 49760 | 3 | 13.4× |
| Lightning | 102057 | 1247.10 MB/s | 49760 | 3 | 13.2× |
| LightningDestructive | 107158 | 1187.74 MB/s | 49280 | 2 | 12.6× |
| SonicFastest | 200043 | 636.24 MB/s | 214134 | 15 | 6.7× |
| Sonic | 201497 | 631.65 MB/s | 214395 | 15 | 6.7× |
| Easyjson | 243303 | 523.11 MB/s | 122864 | 14 | 5.5× |
| Goccy | 247121 | 515.03 MB/s | 225004 | 884 | 5.5× |
| LightningDecodeAny | 443992 | 213.19 MB/s | 463411 | 9708 | 3.0× |
| JSONV2 | 472880 | 269.15 MB/s | 195127 | 1805 | 2.9× |
| Stdlib | 1349170 | 94.34 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4033121 | 558.14 MB/s | 2532849 | 1143 | 8.4× |
| LightningArena | 4096504 | 549.51 MB/s | 2532849 | 1143 | 8.2× |
| Lightning | 4113294 | 547.26 MB/s | 2532852 | 1143 | 8.2× |
| SonicFastest | 5130470 | 438.76 MB/s | 4874968 | 2584 | 6.6× |
| Sonic | 5387925 | 417.80 MB/s | 4871238 | 2584 | 6.3× |
| Goccy | 12595902 | 178.71 MB/s | 4133238 | 56532 | 2.7× |
| LightningDecodeAny | 13548298 | 166.15 MB/s | 19380210 | 223896 | 2.5× |
| Easyjson | 13973908 | 161.09 MB/s | 3099810 | 2120 | 2.4× |
| JSONV2 | 17069070 | 131.88 MB/s | 3123189 | 3083 | 2.0× |
| Stdlib | 33741142 | 66.72 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 553763 | 488.30 MB/s | 397297 | 567 | 7.9× |
| Lightning | 555190 | 487.05 MB/s | 397297 | 567 | 7.9× |
| LightningDestructive | 565095 | 478.51 MB/s | 397296 | 567 | 7.7× |
| SonicFastest | 764376 | 353.76 MB/s | 642123 | 1147 | 5.7× |
| Sonic | 774260 | 349.24 MB/s | 643038 | 1147 | 5.6× |
| Goccy | 1754610 | 154.11 MB/s | 541353 | 8122 | 2.5× |
| Easyjson | 1758375 | 153.78 MB/s | 330272 | 749 | 2.5× |
| LightningDecodeAny | 2225526 | 121.50 MB/s | 2543877 | 29687 | 2.0× |
| JSONV2 | 2291339 | 118.01 MB/s | 348159 | 1628 | 1.9× |
| Stdlib | 4364384 | 61.96 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1119431 | 1542.93 MB/s | 765560 | 2798 | 15.4× |
| LightningArena | 1152275 | 1498.95 MB/s | 768417 | 2440 | 15.0× |
| Lightning | 1161237 | 1487.38 MB/s | 765602 | 2799 | 14.9× |
| SonicFastest | 2053846 | 840.96 MB/s | 2694866 | 5547 | 8.4× |
| Sonic | 2149388 | 803.58 MB/s | 2694394 | 5547 | 8.0× |
| Goccy | 2446475 | 706.00 MB/s | 2581564 | 14603 | 7.1× |
| Easyjson | 3889973 | 444.01 MB/s | 972032 | 5389 | 4.4× |
| LightningDecodeAny | 4020385 | 124.44 MB/s | 4953692 | 76576 | 4.3× |
| JSONV2 | 4665163 | 370.23 MB/s | 1011618 | 7594 | 3.7× |
| Stdlib | 17256542 | 100.09 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1007 | 1799.47 MB/s | 0 | 0 | 16.3× |
| Lightning | 1013 | 1788.87 MB/s | 0 | 0 | 16.2× |
| LightningDestructive | 1088 | 1665.26 MB/s | 0 | 0 | 15.1× |
| Easyjson | 2867 | 632.07 MB/s | 24 | 1 | 5.7× |
| Goccy | 3492 | 518.97 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6062 | 298.92 MB/s | 3345 | 38 | 2.7× |
| Sonic | 6314 | 286.97 MB/s | 3345 | 38 | 2.6× |
| JSONV2 | 8433 | 214.87 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9398 | 192.69 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16455 | 110.12 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1092 | 1659.08 MB/s | 0 | 0 | 15.3× |
| Lightning | 1103 | 1643.37 MB/s | 0 | 0 | 15.2× |
| LightningDestructive | 1179 | 1536.78 MB/s | 0 | 0 | 14.2× |
| Easyjson | 2902 | 624.30 MB/s | 24 | 1 | 5.8× |
| Goccy | 3480 | 520.63 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6223 | 291.19 MB/s | 3343 | 38 | 2.7× |
| Sonic | 6435 | 281.57 MB/s | 3346 | 38 | 2.6× |
| JSONV2 | 8471 | 213.89 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9080 | 199.45 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16718 | 108.39 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1282 | 1413.02 MB/s | 144 | 10 | 13.0× |
| Lightning | 1285 | 1410.43 MB/s | 144 | 10 | 13.0× |
| LightningDestructive | 1339 | 1352.95 MB/s | 144 | 10 | 12.5× |
| Goccy | 3216 | 563.51 MB/s | 2600 | 5 | 5.2× |
| Easyjson | 3221 | 562.60 MB/s | 144 | 10 | 5.2× |
| SonicFastest | 6185 | 292.98 MB/s | 3362 | 40 | 2.7× |
| Sonic | 6423 | 282.10 MB/s | 3367 | 40 | 2.6× |
| JSONV2 | 8408 | 215.51 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 9112 | 198.74 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16701 | 108.50 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 710 | 696.18 MB/s | 160 | 1 | 9.3× |
| LightningDestructive | 724 | 682.47 MB/s | 160 | 1 | 9.1× |
| Sonic | 1277 | 386.92 MB/s | 1076 | 8 | 5.2× |
| SonicFastest | 1281 | 385.65 MB/s | 1076 | 8 | 5.2× |
| LightningDecodeAny | 1520 | 324.37 MB/s | 1296 | 26 | 4.3× |
| LightningArena | 1634 | 302.29 MB/s | 4096 | 1 | 4.0× |
| Goccy | 2554 | 193.41 MB/s | 856 | 23 | 2.6× |
| Easyjson | 2702 | 182.84 MB/s | 448 | 3 | 2.4× |
| JSONV2 | 3525 | 140.14 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6598 | 74.88 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 449 | 511.90 MB/s | 160 | 1 | 10.5× |
| LightningDestructive | 454 | 506.36 MB/s | 160 | 1 | 10.4× |
| Sonic | 931 | 247.16 MB/s | 801 | 8 | 5.1× |
| SonicFastest | 933 | 246.54 MB/s | 801 | 8 | 5.1× |
| LightningDecodeAny | 1299 | 176.28 MB/s | 1296 | 26 | 3.6× |
| LightningArena | 1343 | 171.26 MB/s | 4096 | 1 | 3.5× |
| Easyjson | 1631 | 141.02 MB/s | 448 | 3 | 2.9× |
| Goccy | 1746 | 131.73 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2692 | 85.44 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4719 | 48.74 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 67074 | 971.05 MB/s | 103441 | 103 | 10.1× |
| Lightning | 67837 | 960.12 MB/s | 103441 | 103 | 10.0× |
| LightningDestructive | 72785 | 894.85 MB/s | 97220 | 98 | 9.3× |
| SonicFastest | 147873 | 440.46 MB/s | 235931 | 65 | 4.6× |
| Sonic | 149188 | 436.58 MB/s | 235993 | 65 | 4.5× |
| Goccy | 186552 | 349.14 MB/s | 228169 | 134 | 3.6× |
| LightningDecodeAny | 198690 | 268.40 MB/s | 180049 | 3245 | 3.4× |
| JSONV2 | 271021 | 240.32 MB/s | 206666 | 607 | 2.5× |
| Stdlib | 678329 | 96.02 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2502286 | 775.48 MB/s | 2864592 | 1380 | 11.0× |
| LightningArena | 2565805 | 756.28 MB/s | 2864592 | 1380 | 10.7× |
| Lightning | 2573666 | 753.97 MB/s | 2864594 | 1380 | 10.7× |
| SonicFastest | 4943873 | 392.50 MB/s | 4878829 | 1736 | 5.6× |
| Sonic | 4959455 | 391.27 MB/s | 4879474 | 1736 | 5.5× |
| Goccy | 5209014 | 372.52 MB/s | 4063309 | 13509 | 5.3× |
| Easyjson | 8144110 | 238.27 MB/s | 3871264 | 15043 | 3.4× |
| LightningDecodeAny | 9582998 | 202.49 MB/s | 7063041 | 218633 | 2.9× |
| JSONV2 | 13143275 | 147.64 MB/s | 3237182 | 13947 | 2.1× |
| Stdlib | 27497550 | 70.57 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1051453 | 3164.98 MB/s | 351704 | 1286 | 23.2× |
| Lightning | 1565236 | 2126.09 MB/s | 2488907 | 2995 | 15.6× |
| LightningArena | 1571359 | 2117.80 MB/s | 2488908 | 2995 | 15.6× |
| SonicFastest | 2012637 | 1653.47 MB/s | 5896581 | 4263 | 12.1× |
| Sonic | 2016848 | 1650.02 MB/s | 5896620 | 4263 | 12.1× |
| LightningDecodeAny | 3452789 | 890.23 MB/s | 4876913 | 56892 | 7.1× |
| Goccy | 4971476 | 669.38 MB/s | 3948915 | 3817 | 4.9× |
| JSONV2 | 7730185 | 430.50 MB/s | 5364505 | 13243 | 3.2× |
| Stdlib | 24439560 | 136.17 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 218619 | 1007.90 MB/s | 135872 | 226 | 11.1× |
| Lightning | 220634 | 998.70 MB/s | 135872 | 226 | 11.0× |
| LightningDestructive | 226090 | 974.60 MB/s | 135872 | 226 | 10.7× |
| Sonic | 475457 | 463.44 MB/s | 351080 | 262 | 5.1× |
| SonicFastest | 475950 | 462.96 MB/s | 351215 | 262 | 5.1× |
| Goccy | 486138 | 453.26 MB/s | 363915 | 1066 | 5.0× |
| Easyjson | 628213 | 350.75 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 852362 | 258.51 MB/s | 129747 | 470 | 2.8× |
| LightningDecodeAny | 1003579 | 107.93 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2418336 | 91.11 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12450585 | 650.57 MB/s | 11845073 | 20816 | 8.8× |
| LightningArena | 12582557 | 643.75 MB/s | 11845072 | 20816 | 8.7× |
| Lightning | 12695707 | 638.01 MB/s | 11845072 | 20816 | 8.6× |
| Sonic | 18382525 | 440.64 MB/s | 19859689 | 41640 | 5.9× |
| SonicFastest | 18409389 | 439.99 MB/s | 19857576 | 41640 | 5.9× |
| Goccy | 27194650 | 297.85 MB/s | 19222682 | 107156 | 4.0× |
| Easyjson | 35301751 | 229.45 MB/s | 15059618 | 41643 | 3.1× |
| LightningDecodeAny | 40068700 | 129.85 MB/s | 46279352 | 747112 | 2.7× |
| JSONV2 | 51065431 | 158.62 MB/s | 15233748 | 78972 | 2.1× |
| Stdlib | 109146928 | 74.21 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5423556 | 550.09 MB/s | 3764713 | 1504 | 10.5× |
| LightningDestructive | 5615003 | 531.34 MB/s | 3758856 | 29356 | 10.2× |
| Lightning | 5782940 | 515.91 MB/s | 3758857 | 29356 | 9.9× |
| Sonic | 9591633 | 311.05 MB/s | 9130301 | 57804 | 6.0× |
| SonicFastest | 9682505 | 308.13 MB/s | 9131636 | 57804 | 5.9× |
| LightningDecodeAny | 18512657 | 99.08 MB/s | 23982581 | 351152 | 3.1× |
| Goccy | 18970868 | 157.27 MB/s | 9918993 | 273623 | 3.0× |
| Easyjson | 19826651 | 150.48 MB/s | 9479440 | 30115 | 2.9× |
| JSONV2 | 27821278 | 107.24 MB/s | 9257041 | 86278 | 2.1× |
| Stdlib | 57216754 | 52.14 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1334003 | 542.43 MB/s | 907600 | 3618 | 10.3× |
| LightningArena | 1335805 | 541.69 MB/s | 911392 | 30 | 10.3× |
| Lightning | 1401959 | 516.13 MB/s | 907595 | 3618 | 9.8× |
| SonicFastest | 2121806 | 341.03 MB/s | 2371676 | 3683 | 6.5× |
| Sonic | 2137635 | 338.50 MB/s | 2372020 | 3683 | 6.4× |
| Easyjson | 5387052 | 134.32 MB/s | 2847907 | 3698 | 2.5× |
| Goccy | 5462882 | 132.46 MB/s | 2712034 | 80268 | 2.5× |
| LightningDecodeAny | 5486619 | 118.57 MB/s | 6500460 | 76546 | 2.5× |
| JSONV2 | 6912361 | 104.68 MB/s | 2704702 | 7318 | 2.0× |
| Stdlib | 13708163 | 52.79 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1937819 | 813.98 MB/s | 911394 | 30 | 9.9× |
| Lightning | 1946550 | 810.33 MB/s | 907594 | 3618 | 9.9× |
| LightningDestructive | 1970351 | 800.54 MB/s | 907600 | 3618 | 9.8× |
| SonicFastest | 2513986 | 627.43 MB/s | 3229376 | 3683 | 7.7× |
| Sonic | 2519338 | 626.10 MB/s | 3229932 | 3683 | 7.6× |
| LightningDecodeAny | 4778875 | 157.65 MB/s | 6500454 | 76546 | 4.0× |
| Easyjson | 6403001 | 246.35 MB/s | 2847906 | 3698 | 3.0× |
| Goccy | 6532139 | 241.48 MB/s | 3497001 | 80262 | 2.9× |
| JSONV2 | 7669895 | 205.66 MB/s | 2704551 | 7318 | 2.5× |
| Stdlib | 19248141 | 81.95 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 228340 | 657.46 MB/s | 81920 | 1 | 9.2× |
| Lightning | 229771 | 653.36 MB/s | 81920 | 1 | 9.2× |
| LightningDestructive | 238609 | 629.16 MB/s | 81920 | 1 | 8.8× |
| SonicFastest | 381303 | 393.71 MB/s | 407179 | 16 | 5.5× |
| Sonic | 384107 | 390.84 MB/s | 407137 | 16 | 5.5× |
| LightningDecodeAny | 591628 | 253.74 MB/s | 745765 | 10016 | 3.6× |
| Goccy | 1009286 | 148.74 MB/s | 324701 | 10005 | 2.1× |
| JSONV2 | 1106212 | 135.71 MB/s | 357724 | 20 | 1.9× |
| Stdlib | 2110415 | 71.13 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 32669 | 860.65 MB/s | 29216 | 103 | 10.6× |
| Lightning | 32794 | 857.39 MB/s | 29216 | 103 | 10.6× |
| LightningDestructive | 33905 | 829.29 MB/s | 29088 | 101 | 10.2× |
| SonicFastest | 58029 | 484.53 MB/s | 59527 | 83 | 6.0× |
| Sonic | 58059 | 484.28 MB/s | 59523 | 83 | 6.0× |
| Easyjson | 78975 | 356.02 MB/s | 32304 | 138 | 4.4× |
| Goccy | 83412 | 337.09 MB/s | 59282 | 188 | 4.2× |
| JSONV2 | 147419 | 190.73 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 164376 | 171.05 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 347099 | 81.01 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1961 | 1186.89 MB/s | 32 | 1 | 13.5× |
| Lightning | 2007 | 1159.96 MB/s | 32 | 1 | 13.2× |
| LightningDestructive | 2136 | 1089.80 MB/s | 32 | 1 | 12.4× |
| Sonic | 4828 | 482.19 MB/s | 3711 | 4 | 5.5× |
| SonicFastest | 4834 | 481.60 MB/s | 3709 | 4 | 5.5× |
| Goccy | 4906 | 474.56 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 5291 | 440.01 MB/s | 192 | 2 | 5.0× |
| JSONV2 | 8767 | 265.53 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 10635 | 158.43 MB/s | 10200 | 195 | 2.5× |
| Stdlib | 26395 | 88.20 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 203 | 932.85 MB/s | 0 | 0 | 13.8× |
| LightningArena | 204 | 927.80 MB/s | 0 | 0 | 13.7× |
| LightningDestructive | 209 | 902.39 MB/s | 0 | 0 | 13.3× |
| Goccy | 451 | 418.92 MB/s | 304 | 2 | 6.2× |
| Easyjson | 565 | 334.37 MB/s | 0 | 0 | 4.9× |
| Sonic | 639 | 295.62 MB/s | 341 | 3 | 4.4× |
| SonicFastest | 640 | 295.17 MB/s | 341 | 3 | 4.4× |
| JSONV2 | 1082 | 174.66 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1315 | 101.91 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2790 | 67.74 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1404 | 1560.51 MB/s | 0 | 0 | 13.7× |
| Lightning | 1411 | 1553.13 MB/s | 0 | 0 | 13.6× |
| LightningDestructive | 1559 | 1404.94 MB/s | 0 | 0 | 12.3× |
| Easyjson | 3570 | 613.71 MB/s | 24 | 1 | 5.4× |
| Goccy | 4058 | 539.93 MB/s | 2864 | 4 | 4.7× |
| SonicFastest | 6688 | 327.61 MB/s | 3602 | 38 | 2.9× |
| Sonic | 6905 | 317.31 MB/s | 3603 | 38 | 2.8× |
| JSONV2 | 8635 | 253.74 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9157 | 197.77 MB/s | 7552 | 158 | 2.1× |
| Stdlib | 19229 | 113.94 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 589061 | 866.59 MB/s | 457537 | 1009 | 12.1× |
| LightningArena | 612676 | 833.19 MB/s | 457536 | 1009 | 11.6× |
| Lightning | 618452 | 825.41 MB/s | 457536 | 1009 | 11.5× |
| Goccy | 1237398 | 412.54 MB/s | 1134177 | 5006 | 5.7× |
| Sonic | 1279325 | 399.02 MB/s | 1307225 | 2014 | 5.6× |
| SonicFastest | 1284236 | 397.49 MB/s | 1306938 | 2014 | 5.5× |
| Easyjson | 1811902 | 281.73 MB/s | 863777 | 3012 | 3.9× |
| LightningDecodeAny | 3524661 | 130.92 MB/s | 2950651 | 64018 | 2.0× |
| JSONV2 | 3558715 | 143.44 MB/s | 1075949 | 12645 | 2.0× |
| Stdlib | 7111273 | 71.78 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 544 | 36350.00 MB/s | 0 | 0 | 305.6× |
| LightningArena | 546 | 36256.17 MB/s | 0 | 0 | 304.8× |
| LightningDestructive | 863 | 22939.59 MB/s | 0 | 0 | 192.8× |
| SonicFastest | 6773 | 2921.96 MB/s | 21136 | 3 | 24.6× |
| Goccy | 23880 | 828.68 MB/s | 20492 | 2 | 7.0× |
| Sonic | 31958 | 619.22 MB/s | 20608 | 3 | 5.2× |
| JSONV2 | 33793 | 585.59 MB/s | 8 | 1 | 4.9× |
| LightningDecodeAny | 94324 | 209.79 MB/s | 116864 | 2015 | 1.8× |
| Easyjson | 106828 | 185.24 MB/s | 0 | 0 | 1.6× |
| Stdlib | 166355 | 118.96 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 2220 | 8165.13 MB/s | 432 | 2 | 56.9× |
| Lightning | 2237 | 8102.42 MB/s | 432 | 2 | 56.4× |
| LightningDestructive | 2379 | 7619.26 MB/s | 0 | 0 | 53.1× |
| Easyjson | 4809 | 3768.41 MB/s | 432 | 2 | 26.2× |
| SonicFastest | 8637 | 2098.53 MB/s | 20399 | 5 | 14.6× |
| Sonic | 8976 | 2019.26 MB/s | 20399 | 5 | 14.1× |
| LightningDecodeAny | 19096 | 936.44 MB/s | 29088 | 191 | 6.6× |
| Goccy | 25149 | 720.65 MB/s | 19460 | 2 | 5.0× |
| JSONV2 | 50860 | 356.35 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 126225 | 143.59 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2379265 | 844.17 MB/s | 3089565 | 6821 | 9.2× |
| LightningArena | 2514713 | 798.70 MB/s | 3094370 | 6703 | 8.7× |
| Lightning | 2519481 | 797.19 MB/s | 3091277 | 6827 | 8.7× |
| SonicFastest | 4295029 | 467.63 MB/s | 5154307 | 7085 | 5.1× |
| Sonic | 4340037 | 462.78 MB/s | 5154102 | 7085 | 5.1× |
| Goccy | 4670029 | 430.08 MB/s | 5409654 | 15831 | 4.7× |
| Easyjson | 5414823 | 370.93 MB/s | 2981486 | 7439 | 4.1× |
| LightningDecodeAny | 7082271 | 161.29 MB/s | 8503512 | 134008 | 3.1× |
| JSONV2 | 8105462 | 247.80 MB/s | 3173675 | 14563 | 2.7× |
| Stdlib | 21934092 | 91.57 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 849 | 646.92 MB/s | 480 | 1 | 7.9× |
| LightningArena | 851 | 644.99 MB/s | 480 | 1 | 7.9× |
| LightningDestructive | 857 | 640.84 MB/s | 480 | 1 | 7.8× |
| LightningDecodeAny | 1808 | 303.17 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 2278 | 241.04 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2372 | 231.43 MB/s | 2262 | 8 | 2.8× |
| Sonic | 2395 | 229.25 MB/s | 2262 | 8 | 2.8× |
| Goccy | 3286 | 167.07 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3488 | 157.41 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6704 | 81.89 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 502165 | 1257.58 MB/s | 402728 | 545 | 12.7× |
| LightningArena | 574443 | 1099.35 MB/s | 453017 | 712 | 11.1× |
| Lightning | 583607 | 1082.09 MB/s | 451257 | 857 | 10.9× |
| SonicFastest | 1002553 | 629.91 MB/s | 1067759 | 814 | 6.3× |
| Sonic | 1013959 | 622.82 MB/s | 1067897 | 814 | 6.3× |
| Easyjson | 1353967 | 466.42 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1412791 | 447.00 MB/s | 987698 | 1200 | 4.5× |
| JSONV2 | 2448159 | 257.95 MB/s | 571589 | 3144 | 2.6× |
| LightningDecodeAny | 2515142 | 185.64 MB/s | 2076504 | 30126 | 2.5× |
| Stdlib | 6353420 | 99.40 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 662106 | 849.42 MB/s | 546569 | 429 | 9.1× |
| LightningArena | 853010 | 659.32 MB/s | 771664 | 1088 | 7.1× |
| Lightning | 857449 | 655.91 MB/s | 769937 | 1235 | 7.1× |
| Sonic | 1318253 | 426.63 MB/s | 1349500 | 1185 | 4.6× |
| SonicFastest | 1332483 | 422.08 MB/s | 1350136 | 1185 | 4.5× |
| Goccy | 1626855 | 345.70 MB/s | 1042840 | 1029 | 3.7× |
| Easyjson | 2134256 | 263.51 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 2971525 | 189.27 MB/s | 2180441 | 30126 | 2.0× |
| JSONV2 | 3315349 | 169.64 MB/s | 927407 | 3482 | 1.8× |
| Stdlib | 6051291 | 92.94 MB/s | 1011667 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 634628 | 840.14 MB/s | 333416 | 2084 | 10.2× |
| LightningArena | 694251 | 767.99 MB/s | 368224 | 2293 | 9.3× |
| Lightning | 700892 | 760.71 MB/s | 368224 | 2293 | 9.2× |
| Sonic | 1136362 | 469.20 MB/s | 981480 | 3082 | 5.7× |
| SonicFastest | 1141613 | 467.04 MB/s | 981747 | 3082 | 5.7× |
| Easyjson | 1276528 | 417.68 MB/s | 428362 | 3273 | 5.1× |
| Goccy | 1524384 | 349.77 MB/s | 1167090 | 5409 | 4.2× |
| JSONV2 | 2797122 | 190.62 MB/s | 745422 | 13288 | 2.3× |
| LightningDecodeAny | 3549006 | 150.23 MB/s | 2992875 | 50076 | 1.8× |
| Stdlib | 6453084 | 82.62 MB/s | 798693 | 17133 | 1.0× |
