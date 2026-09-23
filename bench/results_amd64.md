# JSON Deserialization Benchmarks

- generated 2026-09-23T18:18:58Z
- go version go1.26.8 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 59735 | 2130.66 MB/s | 49892 | 2 | 17.1× |
| LightningDestructive | 61997 | 2052.93 MB/s | 49280 | 2 | 16.4× |
| LightningArena | 68121 | 1868.38 MB/s | 49882 | 2 | 15.0× |
| Sonic | 158852 | 801.22 MB/s | 213721 | 15 | 6.4× |
| SonicFastest | 160730 | 791.86 MB/s | 214070 | 15 | 6.3× |
| Easyjson | 181784 | 700.14 MB/s | 122864 | 14 | 5.6× |
| Goccy | 201438 | 631.83 MB/s | 225639 | 884 | 5.1× |
| JSONV2 | 330153 | 385.50 MB/s | 195128 | 1805 | 3.1× |
| LightningDecodeAny | 335839 | 281.84 MB/s | 465155 | 9706 | 3.0× |
| Stdlib | 1019780 | 124.81 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1480613 | 1520.35 MB/s | 2532849 | 1143 | 16.2× |
| Lightning | 1502453 | 1498.25 MB/s | 2532849 | 1143 | 16.0× |
| LightningArena | 1795859 | 1253.47 MB/s | 2532849 | 1143 | 13.4× |
| SonicFastest | 4277468 | 526.26 MB/s | 4876335 | 2584 | 5.6× |
| Sonic | 4317205 | 521.41 MB/s | 4876663 | 2584 | 5.6× |
| LightningDecodeAny | 7100059 | 317.04 MB/s | 6828998 | 223498 | 3.4× |
| Goccy | 9552788 | 235.64 MB/s | 4135295 | 56532 | 2.5× |
| Easyjson | 10724308 | 209.90 MB/s | 3099809 | 2120 | 2.2× |
| JSONV2 | 13490999 | 166.86 MB/s | 3123209 | 3083 | 1.8× |
| Stdlib | 24037177 | 93.65 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 229334 | 1179.08 MB/s | 397296 | 567 | 13.6× |
| LightningDestructive | 241673 | 1118.88 MB/s | 397297 | 567 | 12.9× |
| LightningArena | 278130 | 972.22 MB/s | 397296 | 567 | 11.2× |
| Sonic | 585489 | 461.84 MB/s | 642452 | 1147 | 5.3× |
| SonicFastest | 597069 | 452.88 MB/s | 643068 | 1147 | 5.2× |
| LightningDecodeAny | 1016736 | 265.95 MB/s | 845691 | 29656 | 3.1× |
| Goccy | 1371856 | 197.11 MB/s | 542404 | 8122 | 2.3× |
| Easyjson | 1377621 | 196.28 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 1787620 | 151.26 MB/s | 348159 | 1628 | 1.7× |
| Stdlib | 3123052 | 86.58 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 703373 | 2455.60 MB/s | 765560 | 2798 | 18.6× |
| LightningArena | 721449 | 2394.07 MB/s | 774869 | 2444 | 18.2× |
| Lightning | 724019 | 2385.58 MB/s | 767772 | 2798 | 18.1× |
| Sonic | 1714939 | 1007.15 MB/s | 2695073 | 5547 | 7.6× |
| SonicFastest | 1737961 | 993.81 MB/s | 2695367 | 5547 | 7.5× |
| Goccy | 1945063 | 887.99 MB/s | 2581862 | 14604 | 6.7× |
| LightningDecodeAny | 2818791 | 177.49 MB/s | 4491489 | 67881 | 4.6× |
| Easyjson | 3114705 | 554.53 MB/s | 972032 | 5389 | 4.2× |
| JSONV2 | 3355076 | 514.80 MB/s | 1011615 | 7594 | 3.9× |
| Stdlib | 13095561 | 131.89 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 533 | 3397.55 MB/s | 0 | 0 | 23.0× |
| Lightning | 533 | 3396.83 MB/s | 0 | 0 | 22.9× |
| LightningDestructive | 567 | 3193.54 MB/s | 0 | 0 | 21.6× |
| Easyjson | 2133 | 849.49 MB/s | 24 | 1 | 5.7× |
| Goccy | 2473 | 732.73 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 4718 | 384.07 MB/s | 3350 | 38 | 2.6× |
| Sonic | 4856 | 373.12 MB/s | 3348 | 38 | 2.5× |
| JSONV2 | 5792 | 312.83 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 6552 | 276.42 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 12241 | 148.03 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 556 | 3259.69 MB/s | 0 | 0 | 22.2× |
| Lightning | 556 | 3256.66 MB/s | 0 | 0 | 22.2× |
| LightningDestructive | 587 | 3087.17 MB/s | 0 | 0 | 21.0× |
| Easyjson | 2143 | 845.71 MB/s | 24 | 1 | 5.8× |
| Goccy | 2492 | 727.01 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 4837 | 374.60 MB/s | 3346 | 38 | 2.5× |
| Sonic | 4996 | 362.68 MB/s | 3346 | 38 | 2.5× |
| JSONV2 | 5838 | 310.40 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 6511 | 278.16 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 12326 | 147.01 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 701 | 2583.74 MB/s | 144 | 10 | 17.4× |
| Lightning | 703 | 2578.33 MB/s | 144 | 10 | 17.4× |
| LightningDestructive | 760 | 2382.52 MB/s | 144 | 10 | 16.1× |
| Easyjson | 2376 | 762.77 MB/s | 144 | 10 | 5.1× |
| Goccy | 2637 | 687.27 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 4994 | 362.80 MB/s | 3366 | 40 | 2.4× |
| Sonic | 5135 | 352.85 MB/s | 3366 | 40 | 2.4× |
| JSONV2 | 6247 | 290.08 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 6508 | 278.25 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 12229 | 148.17 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 521 | 947.86 MB/s | 160 | 1 | 9.2× |
| LightningDestructive | 531 | 930.67 MB/s | 160 | 1 | 9.0× |
| SonicFastest | 958 | 515.65 MB/s | 1075 | 8 | 5.0× |
| Sonic | 984 | 502.28 MB/s | 1075 | 8 | 4.9× |
| LightningDecodeAny | 992 | 496.81 MB/s | 1040 | 25 | 4.8× |
| LightningArena | 1214 | 406.92 MB/s | 4120 | 2 | 3.9× |
| Easyjson | 1771 | 278.97 MB/s | 448 | 3 | 2.7× |
| Goccy | 1993 | 247.81 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 2374 | 208.11 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4779 | 103.38 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 200 | 1150.68 MB/s | 160 | 1 | 17.2× |
| LightningDestructive | 202 | 1136.48 MB/s | 160 | 1 | 16.9× |
| Sonic | 690 | 333.51 MB/s | 800 | 8 | 5.0× |
| SonicFastest | 703 | 327.16 MB/s | 800 | 8 | 4.9× |
| LightningDecodeAny | 822 | 278.71 MB/s | 1040 | 25 | 4.2× |
| LightningArena | 928 | 247.73 MB/s | 4120 | 2 | 3.7× |
| Easyjson | 1180 | 194.92 MB/s | 448 | 3 | 2.9× |
| Goccy | 1340 | 171.67 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 1859 | 123.74 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3429 | 67.08 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 43254 | 1505.81 MB/s | 103769 | 99 | 12.0× |
| Lightning | 43527 | 1496.34 MB/s | 103765 | 99 | 12.0× |
| LightningDestructive | 43945 | 1482.12 MB/s | 97220 | 98 | 11.9× |
| Sonic | 121301 | 536.94 MB/s | 236001 | 65 | 4.3× |
| SonicFastest | 127721 | 509.95 MB/s | 236443 | 65 | 4.1× |
| LightningDecodeAny | 138048 | 386.31 MB/s | 176746 | 3237 | 3.8× |
| Goccy | 150253 | 433.48 MB/s | 229394 | 134 | 3.5× |
| JSONV2 | 206369 | 315.61 MB/s | 206664 | 607 | 2.5× |
| Stdlib | 521100 | 124.99 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1525367 | 1272.13 MB/s | 2185296 | 1350 | 13.4× |
| Lightning | 1581261 | 1227.17 MB/s | 2185298 | 1350 | 12.9× |
| LightningArena | 1590133 | 1220.32 MB/s | 2185297 | 1350 | 12.9× |
| Sonic | 3524697 | 550.54 MB/s | 4880125 | 1736 | 5.8× |
| SonicFastest | 3538395 | 548.40 MB/s | 4878382 | 1736 | 5.8× |
| Goccy | 3764176 | 515.51 MB/s | 4062712 | 13509 | 5.4× |
| Easyjson | 5910282 | 328.32 MB/s | 3871265 | 15043 | 3.5× |
| LightningDecodeAny | 6741534 | 287.84 MB/s | 6627984 | 206416 | 3.0× |
| JSONV2 | 8806425 | 220.35 MB/s | 3237182 | 13947 | 2.3× |
| Stdlib | 20447902 | 94.90 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 646583 | 5146.80 MB/s | 351704 | 1286 | 31.2× |
| LightningArena | 975876 | 3410.10 MB/s | 2434707 | 1413 | 20.7× |
| Lightning | 979024 | 3399.13 MB/s | 2434668 | 1413 | 20.6× |
| SonicFastest | 1784626 | 1864.72 MB/s | 5895971 | 4263 | 11.3× |
| Sonic | 1796273 | 1852.63 MB/s | 5896169 | 4263 | 11.2× |
| LightningDecodeAny | 2417053 | 1271.70 MB/s | 4825094 | 55311 | 8.3× |
| Goccy | 4119427 | 807.84 MB/s | 3948913 | 3816 | 4.9× |
| JSONV2 | 6140363 | 541.96 MB/s | 5364497 | 13243 | 3.3× |
| Stdlib | 20166783 | 165.02 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 132429 | 1663.88 MB/s | 135392 | 226 | 13.5× |
| Lightning | 132821 | 1658.97 MB/s | 135392 | 226 | 13.4× |
| LightningDestructive | 137811 | 1598.90 MB/s | 135392 | 226 | 12.9× |
| Goccy | 363479 | 606.21 MB/s | 364324 | 1066 | 4.9× |
| SonicFastest | 415744 | 530.00 MB/s | 351089 | 262 | 4.3× |
| Sonic | 416760 | 528.71 MB/s | 351177 | 262 | 4.3× |
| Easyjson | 449553 | 490.14 MB/s | 130512 | 245 | 4.0× |
| JSONV2 | 524408 | 420.18 MB/s | 129747 | 470 | 3.4× |
| LightningDecodeAny | 719215 | 150.60 MB/s | 854738 | 11700 | 2.5× |
| Stdlib | 1784334 | 123.49 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5623052 | 1440.51 MB/s | 8109648 | 20809 | 14.0× |
| Lightning | 5746720 | 1409.51 MB/s | 8109652 | 20809 | 13.7× |
| LightningArena | 6248217 | 1296.38 MB/s | 8109649 | 20809 | 12.6× |
| Sonic | 16127373 | 502.25 MB/s | 19860092 | 41640 | 4.9× |
| SonicFastest | 16187209 | 500.40 MB/s | 19860243 | 41640 | 4.8× |
| Goccy | 19475040 | 415.92 MB/s | 19019352 | 107155 | 4.0× |
| LightningDecodeAny | 25524055 | 203.85 MB/s | 28359831 | 746961 | 3.1× |
| Easyjson | 26453736 | 306.20 MB/s | 15059617 | 41643 | 3.0× |
| JSONV2 | 34762512 | 233.01 MB/s | 15233747 | 78972 | 2.3× |
| Stdlib | 78499541 | 103.19 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 2420708 | 1232.48 MB/s | 3780457 | 1514 | 17.0× |
| LightningDestructive | 2490375 | 1198.00 MB/s | 3758856 | 29356 | 16.5× |
| Lightning | 2581413 | 1155.75 MB/s | 3758857 | 29356 | 15.9× |
| Sonic | 7228453 | 412.74 MB/s | 9130647 | 57804 | 5.7× |
| SonicFastest | 7290311 | 409.24 MB/s | 9132133 | 57804 | 5.6× |
| LightningDecodeAny | 12151812 | 150.94 MB/s | 18225370 | 350883 | 3.4× |
| Goccy | 13904675 | 214.57 MB/s | 9871028 | 273619 | 3.0× |
| Easyjson | 14204193 | 210.04 MB/s | 9479440 | 30115 | 2.9× |
| JSONV2 | 20030673 | 148.94 MB/s | 9257090 | 86278 | 2.1× |
| Stdlib | 41065729 | 72.65 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 468280 | 1545.22 MB/s | 907602 | 3618 | 21.7× |
| Lightning | 505030 | 1432.78 MB/s | 907596 | 3618 | 20.1× |
| LightningArena | 506429 | 1428.82 MB/s | 916257 | 37 | 20.0× |
| Sonic | 1643889 | 440.17 MB/s | 2376453 | 3683 | 6.2× |
| SonicFastest | 1649128 | 438.78 MB/s | 2376035 | 3683 | 6.1× |
| LightningDecodeAny | 3731281 | 174.36 MB/s | 5691594 | 76540 | 2.7× |
| Easyjson | 4154163 | 174.19 MB/s | 2847906 | 3698 | 2.4× |
| Goccy | 4232380 | 170.97 MB/s | 2742319 | 80269 | 2.4× |
| JSONV2 | 4688053 | 154.35 MB/s | 2704703 | 7318 | 2.2× |
| Stdlib | 10139006 | 71.37 MB/s | 2704547 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 667233 | 2364.02 MB/s | 916256 | 37 | 20.9× |
| LightningDestructive | 686509 | 2297.64 MB/s | 907600 | 3618 | 20.4× |
| Lightning | 690604 | 2284.02 MB/s | 907594 | 3618 | 20.2× |
| SonicFastest | 1953389 | 807.50 MB/s | 3222643 | 3683 | 7.2× |
| Sonic | 1995992 | 790.26 MB/s | 3222667 | 3683 | 7.0× |
| LightningDecodeAny | 3071117 | 245.32 MB/s | 5691590 | 76540 | 4.5× |
| Easyjson | 5096698 | 309.49 MB/s | 2847905 | 3698 | 2.7× |
| Goccy | 5170026 | 305.10 MB/s | 3502542 | 80262 | 2.7× |
| JSONV2 | 5260126 | 299.87 MB/s | 2704552 | 7318 | 2.7× |
| Stdlib | 13972937 | 112.89 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 53445 | 2808.92 MB/s | 81920 | 1 | 30.5× |
| LightningArena | 54266 | 2766.45 MB/s | 81920 | 1 | 30.0× |
| LightningDestructive | 62979 | 2383.73 MB/s | 81920 | 1 | 25.9× |
| SonicFastest | 327167 | 458.86 MB/s | 410477 | 16 | 5.0× |
| Sonic | 333643 | 449.95 MB/s | 411731 | 16 | 4.9× |
| LightningDecodeAny | 428002 | 350.75 MB/s | 745507 | 10015 | 3.8× |
| Goccy | 788177 | 190.47 MB/s | 332564 | 10005 | 2.1× |
| JSONV2 | 916858 | 163.74 MB/s | 357725 | 20 | 1.8× |
| Stdlib | 1629271 | 92.14 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 20184 | 1393.02 MB/s | 29280 | 101 | 12.9× |
| LightningArena | 20467 | 1373.78 MB/s | 29279 | 101 | 12.7× |
| LightningDestructive | 20564 | 1367.27 MB/s | 29088 | 101 | 12.7× |
| Sonic | 54844 | 512.68 MB/s | 59504 | 83 | 4.8× |
| SonicFastest | 54875 | 512.38 MB/s | 59489 | 83 | 4.8× |
| Easyjson | 57432 | 489.57 MB/s | 32304 | 138 | 4.5× |
| Goccy | 60229 | 466.84 MB/s | 59272 | 188 | 4.3× |
| JSONV2 | 100510 | 279.74 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 117418 | 239.46 MB/s | 134075 | 2639 | 2.2× |
| Stdlib | 260876 | 107.78 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1089 | 2137.93 MB/s | 32 | 1 | 17.9× |
| Lightning | 1099 | 2118.65 MB/s | 32 | 1 | 17.7× |
| LightningDestructive | 1171 | 1987.56 MB/s | 32 | 1 | 16.6× |
| Goccy | 3585 | 649.46 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 3796 | 613.31 MB/s | 192 | 2 | 5.1× |
| SonicFastest | 4897 | 475.41 MB/s | 3708 | 4 | 4.0× |
| Sonic | 4917 | 473.49 MB/s | 3713 | 4 | 4.0× |
| JSONV2 | 6128 | 379.90 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 7464 | 225.76 MB/s | 9936 | 194 | 2.6× |
| Stdlib | 19452 | 119.68 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 118 | 1601.95 MB/s | 0 | 0 | 17.2× |
| Lightning | 119 | 1591.60 MB/s | 0 | 0 | 17.1× |
| LightningDestructive | 130 | 1453.33 MB/s | 0 | 0 | 15.6× |
| Goccy | 344 | 548.76 MB/s | 304 | 2 | 5.9× |
| Easyjson | 434 | 435.16 MB/s | 0 | 0 | 4.7× |
| SonicFastest | 576 | 328.25 MB/s | 341 | 3 | 3.5× |
| Sonic | 578 | 327.18 MB/s | 341 | 3 | 3.5× |
| JSONV2 | 744 | 254.07 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 919 | 145.85 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 2033 | 92.97 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 758 | 2889.42 MB/s | 0 | 0 | 19.0× |
| LightningArena | 760 | 2884.76 MB/s | 0 | 0 | 19.0× |
| LightningDestructive | 800 | 2737.23 MB/s | 0 | 0 | 18.0× |
| Easyjson | 2568 | 853.20 MB/s | 24 | 1 | 5.6× |
| Goccy | 2747 | 797.49 MB/s | 2864 | 4 | 5.3× |
| SonicFastest | 5253 | 417.12 MB/s | 3602 | 38 | 2.7× |
| Sonic | 5463 | 401.05 MB/s | 3600 | 38 | 2.6× |
| JSONV2 | 5891 | 371.94 MB/s | 640 | 6 | 2.4× |
| LightningDecodeAny | 6578 | 275.29 MB/s | 7552 | 158 | 2.2× |
| Stdlib | 14424 | 151.90 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 347757 | 1467.91 MB/s | 318401 | 1005 | 15.1× |
| Lightning | 362782 | 1407.12 MB/s | 318400 | 1005 | 14.5× |
| LightningArena | 363050 | 1406.07 MB/s | 318400 | 1005 | 14.5× |
| Goccy | 1005063 | 507.90 MB/s | 1138550 | 5006 | 5.2× |
| SonicFastest | 1134550 | 449.94 MB/s | 1309102 | 2014 | 4.6× |
| Sonic | 1144523 | 446.02 MB/s | 1309325 | 2014 | 4.6× |
| Easyjson | 1226888 | 416.07 MB/s | 863778 | 3012 | 4.3× |
| JSONV2 | 2467671 | 206.87 MB/s | 1075957 | 12645 | 2.1× |
| LightningDecodeAny | 2497003 | 184.81 MB/s | 2742394 | 64017 | 2.1× |
| Stdlib | 5256472 | 97.11 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 196 | 100732.63 MB/s | 0 | 0 | 612.8× |
| LightningArena | 197 | 100583.95 MB/s | 0 | 0 | 612.2× |
| LightningDestructive | 291 | 67965.95 MB/s | 0 | 0 | 413.5× |
| SonicFastest | 5149 | 3843.61 MB/s | 21100 | 3 | 23.4× |
| Goccy | 20069 | 986.03 MB/s | 20492 | 2 | 6.0× |
| Sonic | 22895 | 864.33 MB/s | 20617 | 3 | 5.3× |
| JSONV2 | 28066 | 705.09 MB/s | 8 | 1 | 4.3× |
| LightningDecodeAny | 75132 | 263.38 MB/s | 116608 | 2014 | 1.6× |
| Easyjson | 92703 | 213.47 MB/s | 0 | 0 | 1.3× |
| Stdlib | 120416 | 164.34 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1343 | 13496.82 MB/s | 405 | 0 | 78.1× |
| LightningArena | 1353 | 13391.04 MB/s | 405 | 0 | 77.6× |
| LightningDestructive | 1461 | 12409.05 MB/s | 0 | 0 | 71.8× |
| Easyjson | 3636 | 4984.88 MB/s | 432 | 2 | 28.9× |
| Sonic | 7292 | 2485.45 MB/s | 20424 | 5 | 14.4× |
| SonicFastest | 7998 | 2266.16 MB/s | 20404 | 5 | 13.1× |
| LightningDecodeAny | 14115 | 1266.90 MB/s | 29136 | 189 | 7.4× |
| Goccy | 22639 | 800.58 MB/s | 19460 | 2 | 4.6× |
| JSONV2 | 39164 | 462.77 MB/s | 16500 | 50 | 2.7× |
| Stdlib | 104928 | 172.73 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1541142 | 1303.25 MB/s | 3089565 | 6821 | 10.7× |
| Lightning | 1602814 | 1253.11 MB/s | 3096831 | 6822 | 10.3× |
| LightningArena | 1609727 | 1247.72 MB/s | 3099946 | 6699 | 10.2× |
| Goccy | 3604241 | 557.26 MB/s | 5411654 | 15832 | 4.6× |
| SonicFastest | 3791776 | 529.70 MB/s | 5155965 | 7085 | 4.3× |
| Sonic | 3868482 | 519.19 MB/s | 5158654 | 7085 | 4.3× |
| Easyjson | 4146888 | 484.34 MB/s | 2981488 | 7439 | 4.0× |
| LightningDecodeAny | 5318661 | 214.77 MB/s | 7375378 | 134004 | 3.1× |
| JSONV2 | 5678842 | 353.68 MB/s | 3173673 | 14562 | 2.9× |
| Stdlib | 16442646 | 122.15 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 611 | 898.67 MB/s | 480 | 1 | 7.8× |
| LightningArena | 611 | 897.92 MB/s | 480 | 1 | 7.8× |
| LightningDestructive | 618 | 888.53 MB/s | 480 | 1 | 7.8× |
| LightningDecodeAny | 1251 | 438.22 MB/s | 1765 | 45 | 3.8× |
| Easyjson | 1512 | 363.19 MB/s | 1616 | 5 | 3.2× |
| SonicFastest | 1641 | 334.49 MB/s | 2261 | 8 | 2.9× |
| Sonic | 1700 | 323.02 MB/s | 2262 | 8 | 2.8× |
| JSONV2 | 2289 | 239.79 MB/s | 1664 | 7 | 2.1× |
| Goccy | 2329 | 235.72 MB/s | 2129 | 43 | 2.1× |
| Stdlib | 4790 | 114.61 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 307295 | 2055.07 MB/s | 400488 | 545 | 15.8× |
| LightningArena | 339857 | 1858.18 MB/s | 448640 | 404 | 14.3× |
| Lightning | 345869 | 1825.87 MB/s | 446859 | 548 | 14.1× |
| Sonic | 892453 | 707.62 MB/s | 1069979 | 814 | 5.4× |
| SonicFastest | 893039 | 707.15 MB/s | 1067423 | 814 | 5.4× |
| Easyjson | 990890 | 637.32 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1056103 | 597.97 MB/s | 989237 | 1200 | 4.6× |
| JSONV2 | 1670703 | 377.99 MB/s | 571591 | 3144 | 2.9× |
| LightningDecodeAny | 1807550 | 258.31 MB/s | 1991108 | 29072 | 2.7× |
| Stdlib | 4862667 | 129.87 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 421050 | 1335.73 MB/s | 390677 | 426 | 10.9× |
| LightningArena | 498594 | 1127.99 MB/s | 508048 | 287 | 9.2× |
| Lightning | 502069 | 1120.18 MB/s | 506284 | 433 | 9.1× |
| Sonic | 1092239 | 514.91 MB/s | 1350041 | 1185 | 4.2× |
| SonicFastest | 1108461 | 507.38 MB/s | 1351191 | 1185 | 4.1× |
| Goccy | 1227617 | 458.13 MB/s | 1041834 | 1028 | 3.7× |
| Easyjson | 1558502 | 360.86 MB/s | 775153 | 1254 | 2.9× |
| LightningDecodeAny | 1895020 | 296.78 MB/s | 1990396 | 28580 | 2.4× |
| JSONV2 | 2356880 | 238.62 MB/s | 927408 | 3482 | 1.9× |
| Stdlib | 4584180 | 122.68 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 381493 | 1397.61 MB/s | 333416 | 2084 | 12.8× |
| LightningArena | 428459 | 1244.41 MB/s | 367781 | 2086 | 11.4× |
| Lightning | 429599 | 1241.11 MB/s | 367801 | 2086 | 11.4× |
| Easyjson | 983445 | 542.15 MB/s | 428362 | 3273 | 5.0× |
| Sonic | 1075348 | 495.82 MB/s | 982246 | 3082 | 4.5× |
| SonicFastest | 1078450 | 494.39 MB/s | 982117 | 3082 | 4.5× |
| Goccy | 1185778 | 449.64 MB/s | 1167070 | 5408 | 4.1× |
| JSONV2 | 2074242 | 257.05 MB/s | 745420 | 13288 | 2.4× |
| LightningDecodeAny | 2472031 | 215.68 MB/s | 2658331 | 49348 | 2.0× |
| Stdlib | 4880168 | 109.25 MB/s | 798692 | 17133 | 1.0× |
