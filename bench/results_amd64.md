# JSON Deserialization Benchmarks

- generated 2026-06-23T18:19:15Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105010 | 1212.02 MB/s | 49760 | 3 | 12.7× |
| LightningDestructive | 106860 | 1191.04 MB/s | 49280 | 2 | 12.4× |
| SonicFastest | 198095 | 642.49 MB/s | 214351 | 15 | 6.7× |
| Sonic | 199111 | 639.22 MB/s | 214386 | 15 | 6.7× |
| Goccy | 247438 | 514.37 MB/s | 224953 | 884 | 5.4× |
| Easyjson | 251366 | 506.33 MB/s | 122864 | 14 | 5.3× |
| JSONV2 | 446473 | 285.07 MB/s | 195129 | 1805 | 3.0× |
| LightningDecodeAny | 470127 | 201.33 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1328886 | 95.78 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4144282 | 543.17 MB/s | 3122872 | 3081 | 8.3× |
| Lightning | 4194083 | 536.72 MB/s | 3122873 | 3081 | 8.2× |
| SonicFastest | 5065055 | 444.43 MB/s | 4872418 | 2584 | 6.8× |
| Sonic | 5282516 | 426.13 MB/s | 4870619 | 2584 | 6.5× |
| Goccy | 11891341 | 189.30 MB/s | 4174131 | 56534 | 2.9× |
| LightningDecodeAny | 12380431 | 181.82 MB/s | 7938299 | 281383 | 2.8× |
| Easyjson | 13859546 | 162.42 MB/s | 3099810 | 2120 | 2.5× |
| JSONV2 | 17318838 | 129.98 MB/s | 3123215 | 3083 | 2.0× |
| Stdlib | 34201949 | 65.82 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 575144 | 470.15 MB/s | 348025 | 1627 | 7.6× |
| LightningDestructive | 579452 | 466.65 MB/s | 348024 | 1627 | 7.6× |
| SonicFastest | 752197 | 359.48 MB/s | 641900 | 1147 | 5.8× |
| Sonic | 753970 | 358.64 MB/s | 642214 | 1147 | 5.8× |
| Goccy | 1658548 | 163.04 MB/s | 541262 | 8122 | 2.6× |
| LightningDecodeAny | 1712099 | 157.94 MB/s | 1011487 | 37901 | 2.6× |
| Easyjson | 1781439 | 151.79 MB/s | 330272 | 749 | 2.5× |
| JSONV2 | 2308838 | 117.12 MB/s | 348161 | 1628 | 1.9× |
| Stdlib | 4390941 | 61.58 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1289384 | 1339.56 MB/s | 959848 | 5881 | 13.7× |
| Lightning | 1341880 | 1287.15 MB/s | 959890 | 5882 | 13.2× |
| Sonic | 2022279 | 854.09 MB/s | 2693960 | 5547 | 8.7× |
| SonicFastest | 2028605 | 851.42 MB/s | 2693237 | 5547 | 8.7× |
| Goccy | 2520175 | 685.35 MB/s | 2581034 | 14603 | 7.0× |
| Easyjson | 4033316 | 428.23 MB/s | 972032 | 5389 | 4.4× |
| LightningDecodeAny | 4429399 | 112.95 MB/s | 4976203 | 81466 | 4.0× |
| JSONV2 | 4772365 | 361.92 MB/s | 1011618 | 7594 | 3.7× |
| Stdlib | 17687675 | 97.65 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1023 | 1771.03 MB/s | 0 | 0 | 15.7× |
| LightningDestructive | 1116 | 1623.72 MB/s | 0 | 0 | 14.4× |
| Easyjson | 2961 | 611.96 MB/s | 24 | 1 | 5.4× |
| Goccy | 3280 | 552.48 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6301 | 287.56 MB/s | 3346 | 38 | 2.5× |
| Sonic | 6488 | 279.28 MB/s | 3344 | 38 | 2.5× |
| JSONV2 | 8297 | 218.39 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9378 | 193.12 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16031 | 113.03 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1120 | 1617.38 MB/s | 0 | 0 | 14.3× |
| LightningDestructive | 1196 | 1514.66 MB/s | 0 | 0 | 13.4× |
| Easyjson | 3042 | 595.67 MB/s | 24 | 1 | 5.3× |
| Goccy | 3203 | 565.80 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6012 | 301.38 MB/s | 3343 | 38 | 2.7× |
| Sonic | 6280 | 288.56 MB/s | 3346 | 38 | 2.6× |
| JSONV2 | 8052 | 225.04 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9251 | 195.76 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16024 | 113.08 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1312 | 1380.93 MB/s | 144 | 10 | 12.3× |
| LightningDestructive | 1399 | 1294.92 MB/s | 144 | 10 | 11.6× |
| Easyjson | 3092 | 585.97 MB/s | 144 | 10 | 5.2× |
| Goccy | 3463 | 523.31 MB/s | 2600 | 5 | 4.7× |
| SonicFastest | 6126 | 295.80 MB/s | 3363 | 40 | 2.6× |
| Sonic | 6338 | 285.91 MB/s | 3367 | 40 | 2.6× |
| JSONV2 | 8552 | 211.89 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9408 | 192.49 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16195 | 111.89 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 725 | 681.44 MB/s | 160 | 1 | 9.6× |
| LightningDestructive | 738 | 669.04 MB/s | 160 | 1 | 9.4× |
| Sonic | 1271 | 388.74 MB/s | 1075 | 8 | 5.5× |
| SonicFastest | 1273 | 388.08 MB/s | 1076 | 8 | 5.5× |
| LightningDecodeAny | 1814 | 271.73 MB/s | 1536 | 30 | 3.8× |
| Goccy | 2586 | 191.07 MB/s | 856 | 23 | 2.7× |
| Easyjson | 2694 | 183.37 MB/s | 448 | 3 | 2.6× |
| JSONV2 | 3460 | 142.77 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6970 | 70.88 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 495 | 464.74 MB/s | 160 | 1 | 10.0× |
| LightningDestructive | 498 | 461.83 MB/s | 160 | 1 | 9.9× |
| Sonic | 922 | 249.44 MB/s | 800 | 8 | 5.3× |
| SonicFastest | 924 | 249.04 MB/s | 800 | 8 | 5.3× |
| LightningDecodeAny | 1612 | 142.02 MB/s | 1536 | 30 | 3.1× |
| Goccy | 1804 | 127.51 MB/s | 584 | 23 | 2.7× |
| Easyjson | 1819 | 126.45 MB/s | 448 | 3 | 2.7× |
| JSONV2 | 2582 | 89.08 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4928 | 46.67 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 93918 | 693.50 MB/s | 172434 | 107 | 7.3× |
| LightningDestructive | 105591 | 616.83 MB/s | 166213 | 102 | 6.5× |
| Sonic | 144173 | 451.76 MB/s | 235829 | 65 | 4.7× |
| SonicFastest | 147412 | 441.84 MB/s | 235899 | 65 | 4.6× |
| Goccy | 179131 | 363.60 MB/s | 228135 | 134 | 3.8× |
| LightningDecodeAny | 206815 | 257.86 MB/s | 176960 | 3252 | 3.3× |
| JSONV2 | 264157 | 246.57 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 684086 | 95.21 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2562437 | 757.28 MB/s | 2846864 | 2238 | 11.0× |
| Lightning | 2631291 | 737.46 MB/s | 2846866 | 2238 | 10.7× |
| Sonic | 4922502 | 394.20 MB/s | 4874226 | 1736 | 5.7× |
| SonicFastest | 4977594 | 389.84 MB/s | 4874239 | 1736 | 5.6× |
| Goccy | 5051898 | 384.11 MB/s | 4062348 | 13509 | 5.6× |
| Easyjson | 8311839 | 233.46 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 10184974 | 190.52 MB/s | 7013525 | 219937 | 2.8× |
| JSONV2 | 12327210 | 157.41 MB/s | 3237182 | 13947 | 2.3× |
| Stdlib | 28064129 | 69.14 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1061672 | 3134.52 MB/s | 351704 | 1286 | 22.5× |
| Lightning | 1618887 | 2055.63 MB/s | 2488907 | 2995 | 14.8× |
| Sonic | 2002710 | 1661.66 MB/s | 5896561 | 4263 | 11.9× |
| SonicFastest | 2032730 | 1637.12 MB/s | 5895272 | 4263 | 11.8× |
| LightningDecodeAny | 3638113 | 844.88 MB/s | 4886622 | 56892 | 6.6× |
| Goccy | 6076008 | 547.70 MB/s | 3948914 | 3816 | 3.9× |
| JSONV2 | 7701338 | 432.11 MB/s | 5364506 | 13243 | 3.1× |
| Stdlib | 23931578 | 139.06 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225739 | 976.11 MB/s | 136896 | 228 | 10.9× |
| LightningDestructive | 235378 | 936.14 MB/s | 136896 | 228 | 10.4× |
| SonicFastest | 488862 | 450.73 MB/s | 351117 | 262 | 5.0× |
| Sonic | 489211 | 450.41 MB/s | 350926 | 262 | 5.0× |
| Goccy | 492105 | 447.76 MB/s | 363378 | 1066 | 5.0× |
| Easyjson | 674188 | 326.83 MB/s | 130512 | 245 | 3.6× |
| JSONV2 | 725016 | 303.92 MB/s | 129747 | 470 | 3.4× |
| LightningDecodeAny | 1022653 | 105.91 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2457014 | 89.68 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14918744 | 542.94 MB/s | 15059834 | 51649 | 7.5× |
| Lightning | 15807395 | 512.42 MB/s | 15059844 | 51649 | 7.1× |
| Sonic | 18270960 | 443.33 MB/s | 19861383 | 41640 | 6.1× |
| SonicFastest | 18291322 | 442.84 MB/s | 19861024 | 41640 | 6.1× |
| Goccy | 25082879 | 322.93 MB/s | 19155353 | 107156 | 4.5× |
| Easyjson | 37353034 | 216.85 MB/s | 15059619 | 41643 | 3.0× |
| LightningDecodeAny | 44158606 | 117.83 MB/s | 34333349 | 912810 | 2.5× |
| JSONV2 | 49047687 | 165.15 MB/s | 15233746 | 78972 | 2.3× |
| Stdlib | 112334528 | 72.11 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6254243 | 477.03 MB/s | 3985336 | 30015 | 9.2× |
| Lightning | 6292791 | 474.11 MB/s | 3985337 | 30015 | 9.1× |
| Sonic | 9260202 | 322.18 MB/s | 9132128 | 57804 | 6.2× |
| SonicFastest | 9279574 | 321.51 MB/s | 9132258 | 57804 | 6.2× |
| Goccy | 18513876 | 161.15 MB/s | 9913382 | 273622 | 3.1× |
| Easyjson | 18821930 | 158.51 MB/s | 9479441 | 30115 | 3.0× |
| LightningDecodeAny | 20744465 | 88.42 MB/s | 20023838 | 408557 | 2.8× |
| JSONV2 | 29028238 | 102.78 MB/s | 9257050 | 86278 | 2.0× |
| Stdlib | 57396105 | 51.98 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1476742 | 490.00 MB/s | 907394 | 3618 | 9.5× |
| Lightning | 1545549 | 468.18 MB/s | 907387 | 3618 | 9.1× |
| SonicFastest | 2099285 | 344.69 MB/s | 2370982 | 3683 | 6.7× |
| Sonic | 2101167 | 344.38 MB/s | 2370942 | 3683 | 6.7× |
| Easyjson | 5369619 | 134.76 MB/s | 2847906 | 3698 | 2.6× |
| Goccy | 5484242 | 131.94 MB/s | 2724766 | 80269 | 2.6× |
| LightningDecodeAny | 5640821 | 115.33 MB/s | 5752201 | 80172 | 2.5× |
| JSONV2 | 6362989 | 113.72 MB/s | 2704707 | 7318 | 2.2× |
| Stdlib | 14087862 | 51.36 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2046602 | 770.72 MB/s | 907387 | 3618 | 9.7× |
| LightningDestructive | 2051106 | 769.03 MB/s | 907392 | 3618 | 9.6× |
| Sonic | 2411763 | 654.03 MB/s | 3223276 | 3683 | 8.2× |
| SonicFastest | 2429397 | 649.28 MB/s | 3223261 | 3683 | 8.1× |
| LightningDecodeAny | 4827742 | 156.06 MB/s | 5752198 | 80172 | 4.1× |
| Easyjson | 6354839 | 248.21 MB/s | 2847905 | 3698 | 3.1× |
| Goccy | 6674225 | 236.33 MB/s | 3464637 | 80260 | 3.0× |
| JSONV2 | 6917319 | 228.03 MB/s | 2704553 | 7318 | 2.9× |
| Stdlib | 19770058 | 79.78 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 251179 | 597.68 MB/s | 81920 | 1 | 8.7× |
| LightningDestructive | 261309 | 574.51 MB/s | 81920 | 1 | 8.3× |
| Sonic | 380584 | 394.46 MB/s | 407537 | 16 | 5.7× |
| SonicFastest | 381922 | 393.08 MB/s | 407891 | 16 | 5.7× |
| LightningDecodeAny | 601989 | 249.38 MB/s | 746005 | 10020 | 3.6× |
| Goccy | 1030984 | 145.61 MB/s | 324516 | 10005 | 2.1× |
| JSONV2 | 1103883 | 136.00 MB/s | 357724 | 20 | 2.0× |
| Stdlib | 2176200 | 68.98 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33920 | 828.92 MB/s | 30272 | 105 | 10.0× |
| LightningDestructive | 35117 | 800.66 MB/s | 30144 | 103 | 9.7× |
| SonicFastest | 58387 | 481.57 MB/s | 59514 | 83 | 5.8× |
| Sonic | 59084 | 475.88 MB/s | 59555 | 83 | 5.7× |
| Goccy | 80864 | 347.71 MB/s | 59288 | 188 | 4.2× |
| Easyjson | 81567 | 344.71 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 132734 | 211.83 MB/s | 36896 | 242 | 2.6× |
| LightningDecodeAny | 167799 | 167.56 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 338956 | 82.95 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1955 | 1190.55 MB/s | 32 | 1 | 13.5× |
| LightningDestructive | 2104 | 1106.36 MB/s | 32 | 1 | 12.5× |
| Sonic | 4701 | 495.24 MB/s | 3712 | 4 | 5.6× |
| SonicFastest | 4705 | 494.83 MB/s | 3711 | 4 | 5.6× |
| Goccy | 4939 | 471.38 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 5045 | 461.44 MB/s | 192 | 2 | 5.2× |
| JSONV2 | 8378 | 277.87 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10673 | 157.87 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 26295 | 88.54 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 215 | 880.86 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 225 | 839.53 MB/s | 0 | 0 | 12.4× |
| Goccy | 427 | 443.00 MB/s | 304 | 2 | 6.5× |
| Easyjson | 591 | 319.98 MB/s | 0 | 0 | 4.7× |
| SonicFastest | 627 | 301.61 MB/s | 341 | 3 | 4.4× |
| Sonic | 635 | 297.81 MB/s | 341 | 3 | 4.4× |
| JSONV2 | 1001 | 188.86 MB/s | 112 | 1 | 2.8× |
| LightningDecodeAny | 1360 | 98.51 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2782 | 67.92 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1490 | 1470.81 MB/s | 0 | 0 | 12.7× |
| LightningDestructive | 1559 | 1405.49 MB/s | 0 | 0 | 12.2× |
| Goccy | 3607 | 607.45 MB/s | 2864 | 4 | 5.3× |
| Easyjson | 3816 | 574.13 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6646 | 329.68 MB/s | 3600 | 38 | 2.9× |
| Sonic | 6891 | 317.96 MB/s | 3601 | 38 | 2.8× |
| JSONV2 | 8271 | 264.90 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 9272 | 195.31 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18995 | 115.35 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 702356 | 726.81 MB/s | 703778 | 1012 | 10.1× |
| Lightning | 735026 | 694.50 MB/s | 703778 | 1012 | 9.6× |
| SonicFastest | 1292681 | 394.90 MB/s | 1307306 | 2014 | 5.5× |
| Sonic | 1294246 | 394.42 MB/s | 1307480 | 2014 | 5.5× |
| Goccy | 1338851 | 381.28 MB/s | 1138167 | 5006 | 5.3× |
| Easyjson | 1685934 | 302.79 MB/s | 863779 | 3012 | 4.2× |
| JSONV2 | 3519653 | 145.04 MB/s | 1075970 | 12645 | 2.0× |
| LightningDecodeAny | 3576788 | 129.02 MB/s | 2785927 | 66022 | 2.0× |
| Stdlib | 7069753 | 72.21 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 643 | 30759.71 MB/s | 0 | 0 | 253.3× |
| LightningDestructive | 966 | 20474.33 MB/s | 0 | 0 | 168.6× |
| SonicFastest | 6703 | 2952.21 MB/s | 21161 | 3 | 24.3× |
| Goccy | 29933 | 661.11 MB/s | 20492 | 2 | 5.4× |
| Sonic | 31919 | 619.97 MB/s | 20608 | 3 | 5.1× |
| JSONV2 | 33411 | 592.30 MB/s | 8 | 1 | 4.9× |
| Easyjson | 96761 | 204.51 MB/s | 0 | 0 | 1.7× |
| LightningDecodeAny | 103242 | 191.67 MB/s | 117104 | 2019 | 1.6× |
| Stdlib | 162939 | 121.45 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2211 | 8195.60 MB/s | 432 | 2 | 55.5× |
| LightningDestructive | 2458 | 7372.14 MB/s | 0 | 0 | 49.9× |
| Easyjson | 5030 | 3603.22 MB/s | 432 | 2 | 24.4× |
| Sonic | 8652 | 2094.69 MB/s | 20424 | 5 | 14.2× |
| SonicFastest | 8809 | 2057.49 MB/s | 20413 | 5 | 13.9× |
| LightningDecodeAny | 19076 | 937.42 MB/s | 29088 | 191 | 6.4× |
| Goccy | 23587 | 768.37 MB/s | 19460 | 2 | 5.2× |
| JSONV2 | 52655 | 344.21 MB/s | 16500 | 50 | 2.3× |
| Stdlib | 122701 | 147.71 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2335764 | 859.89 MB/s | 2960389 | 7411 | 9.5× |
| Lightning | 2398698 | 837.33 MB/s | 2962101 | 7417 | 9.3× |
| SonicFastest | 4122946 | 487.15 MB/s | 5158063 | 7085 | 5.4× |
| Sonic | 4161120 | 482.68 MB/s | 5155597 | 7085 | 5.3× |
| Goccy | 4857370 | 413.49 MB/s | 5410684 | 15833 | 4.6× |
| Easyjson | 5755242 | 348.99 MB/s | 2981490 | 7439 | 3.9× |
| LightningDecodeAny | 7378592 | 154.81 MB/s | 7386074 | 134751 | 3.0× |
| JSONV2 | 7879484 | 254.90 MB/s | 3173679 | 14563 | 2.8× |
| Stdlib | 22197441 | 90.48 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 848 | 647.09 MB/s | 480 | 1 | 8.1× |
| LightningDestructive | 850 | 645.53 MB/s | 480 | 1 | 8.0× |
| LightningDecodeAny | 2152 | 254.69 MB/s | 2261 | 50 | 3.2× |
| Easyjson | 2364 | 232.21 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2382 | 230.49 MB/s | 2261 | 8 | 2.9× |
| Sonic | 2410 | 227.75 MB/s | 2260 | 8 | 2.8× |
| Goccy | 3373 | 162.78 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3600 | 152.52 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6834 | 80.33 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 501191 | 1260.03 MB/s | 364473 | 566 | 12.8× |
| Lightning | 584104 | 1081.17 MB/s | 413001 | 878 | 11.0× |
| SonicFastest | 1002990 | 629.63 MB/s | 1066566 | 814 | 6.4× |
| Sonic | 1005170 | 628.27 MB/s | 1066364 | 814 | 6.4× |
| Easyjson | 1321846 | 477.75 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1414867 | 446.34 MB/s | 990055 | 1200 | 4.5× |
| JSONV2 | 2291603 | 275.58 MB/s | 571590 | 3144 | 2.8× |
| LightningDecodeAny | 2606204 | 179.15 MB/s | 2010198 | 30295 | 2.5× |
| Stdlib | 6423449 | 98.31 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 825850 | 681.01 MB/s | 544249 | 448 | 7.5× |
| Lightning | 1036611 | 542.54 MB/s | 767622 | 1254 | 6.0× |
| Sonic | 1281053 | 439.02 MB/s | 1347952 | 1185 | 4.8× |
| SonicFastest | 1289135 | 436.27 MB/s | 1348265 | 1185 | 4.8× |
| Goccy | 1639109 | 343.12 MB/s | 1039469 | 1028 | 3.8× |
| Easyjson | 2166247 | 259.62 MB/s | 775154 | 1254 | 2.9× |
| LightningDecodeAny | 3159326 | 178.02 MB/s | 2114150 | 30295 | 2.0× |
| JSONV2 | 3245662 | 173.28 MB/s | 927408 | 3482 | 1.9× |
| Stdlib | 6204367 | 90.65 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 626348 | 851.25 MB/s | 333416 | 2084 | 10.2× |
| Lightning | 701111 | 760.48 MB/s | 368224 | 2293 | 9.1× |
| Sonic | 1135307 | 469.63 MB/s | 981651 | 3082 | 5.6× |
| SonicFastest | 1140354 | 467.55 MB/s | 982172 | 3082 | 5.6× |
| Easyjson | 1331076 | 400.56 MB/s | 428362 | 3273 | 4.8× |
| Goccy | 1452722 | 367.02 MB/s | 1167090 | 5409 | 4.4× |
| JSONV2 | 2848807 | 187.16 MB/s | 745422 | 13288 | 2.2× |
| LightningDecodeAny | 3609790 | 147.70 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6408001 | 83.21 MB/s | 798693 | 17133 | 1.0× |
