# JSON Deserialization Benchmarks

- generated 2026-07-27T11:19:54Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 99863 | 1274.50 MB/s | 49760 | 3 | 13.4× |
| Lightning | 101508 | 1253.84 MB/s | 49760 | 3 | 13.2× |
| LightningDestructive | 107311 | 1186.04 MB/s | 49280 | 2 | 12.5× |
| Sonic | 199095 | 639.27 MB/s | 214379 | 15 | 6.7× |
| SonicFastest | 199132 | 639.15 MB/s | 214249 | 15 | 6.7× |
| Easyjson | 242607 | 524.61 MB/s | 122864 | 14 | 5.5× |
| Goccy | 244622 | 520.29 MB/s | 225436 | 884 | 5.5× |
| LightningDecodeAny | 449548 | 210.55 MB/s | 463410 | 9708 | 3.0× |
| JSONV2 | 471988 | 269.66 MB/s | 195127 | 1805 | 2.8× |
| Stdlib | 1338229 | 95.11 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4034842 | 557.90 MB/s | 2532849 | 1143 | 8.3× |
| Lightning | 4113334 | 547.26 MB/s | 2532852 | 1143 | 8.2× |
| LightningArena | 4121467 | 546.18 MB/s | 2532849 | 1143 | 8.2× |
| Sonic | 5390961 | 417.56 MB/s | 4868180 | 2584 | 6.2× |
| SonicFastest | 5395979 | 417.17 MB/s | 4868055 | 2584 | 6.2× |
| Goccy | 12849763 | 175.18 MB/s | 4231225 | 56537 | 2.6× |
| LightningDecodeAny | 13417167 | 167.77 MB/s | 19380210 | 223896 | 2.5× |
| Easyjson | 13662550 | 164.76 MB/s | 3099810 | 2120 | 2.5× |
| JSONV2 | 17058717 | 131.96 MB/s | 3123181 | 3083 | 2.0× |
| Stdlib | 33664816 | 66.87 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 551912 | 489.94 MB/s | 397296 | 567 | 7.9× |
| Lightning | 554040 | 488.06 MB/s | 397297 | 567 | 7.8× |
| LightningDestructive | 560190 | 482.70 MB/s | 397296 | 567 | 7.7× |
| Sonic | 743395 | 363.74 MB/s | 640267 | 1147 | 5.8× |
| SonicFastest | 747232 | 361.87 MB/s | 640428 | 1147 | 5.8× |
| Easyjson | 1733706 | 155.97 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1758354 | 153.78 MB/s | 540984 | 8121 | 2.5× |
| LightningDecodeAny | 2215756 | 122.04 MB/s | 2543877 | 29687 | 2.0× |
| JSONV2 | 2264067 | 119.43 MB/s | 348160 | 1628 | 1.9× |
| Stdlib | 4340831 | 62.29 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1118366 | 1544.40 MB/s | 765560 | 2798 | 15.2× |
| LightningArena | 1138250 | 1517.42 MB/s | 768416 | 2440 | 14.9× |
| Lightning | 1149014 | 1503.21 MB/s | 765602 | 2799 | 14.8× |
| Sonic | 2043071 | 845.40 MB/s | 2694282 | 5547 | 8.3× |
| SonicFastest | 2049838 | 842.61 MB/s | 2693929 | 5547 | 8.3× |
| Goccy | 2583442 | 668.57 MB/s | 2580925 | 14603 | 6.6× |
| Easyjson | 3858690 | 447.61 MB/s | 972032 | 5389 | 4.4× |
| LightningDecodeAny | 4077614 | 122.69 MB/s | 4953691 | 76576 | 4.2× |
| JSONV2 | 4911858 | 351.64 MB/s | 1011615 | 7594 | 3.5× |
| Stdlib | 17006172 | 101.56 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 988 | 1834.59 MB/s | 0 | 0 | 16.7× |
| Lightning | 1011 | 1793.16 MB/s | 0 | 0 | 16.3× |
| LightningDestructive | 1061 | 1708.47 MB/s | 0 | 0 | 15.6× |
| Easyjson | 2872 | 630.95 MB/s | 24 | 1 | 5.7× |
| Goccy | 3473 | 521.72 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6048 | 299.61 MB/s | 3346 | 38 | 2.7× |
| Sonic | 6367 | 284.60 MB/s | 3348 | 38 | 2.6× |
| JSONV2 | 8456 | 214.29 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9295 | 194.83 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16504 | 109.79 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1089 | 1664.00 MB/s | 0 | 0 | 15.2× |
| Lightning | 1102 | 1644.21 MB/s | 0 | 0 | 15.0× |
| LightningDestructive | 1168 | 1550.72 MB/s | 0 | 0 | 14.2× |
| Easyjson | 2891 | 626.73 MB/s | 24 | 1 | 5.7× |
| Goccy | 3491 | 519.05 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6180 | 293.20 MB/s | 3347 | 38 | 2.7× |
| Sonic | 6411 | 282.63 MB/s | 3346 | 38 | 2.6× |
| JSONV2 | 8423 | 215.13 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9027 | 200.63 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16528 | 109.63 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1272 | 1424.37 MB/s | 144 | 10 | 13.2× |
| LightningArena | 1275 | 1421.60 MB/s | 144 | 10 | 13.1× |
| LightningDestructive | 1340 | 1352.30 MB/s | 144 | 10 | 12.5× |
| Goccy | 3195 | 567.07 MB/s | 2600 | 5 | 5.2× |
| Easyjson | 3231 | 560.74 MB/s | 144 | 10 | 5.2× |
| SonicFastest | 6210 | 291.77 MB/s | 3367 | 40 | 2.7× |
| Sonic | 6449 | 280.98 MB/s | 3369 | 40 | 2.6× |
| JSONV2 | 8691 | 208.50 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9133 | 198.30 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16733 | 108.29 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 695 | 710.35 MB/s | 160 | 1 | 9.5× |
| LightningDestructive | 706 | 699.28 MB/s | 160 | 1 | 9.3× |
| Sonic | 1274 | 387.76 MB/s | 1076 | 8 | 5.2× |
| SonicFastest | 1276 | 387.18 MB/s | 1076 | 8 | 5.2× |
| LightningDecodeAny | 1510 | 326.50 MB/s | 1296 | 26 | 4.4× |
| LightningArena | 1600 | 308.85 MB/s | 4096 | 1 | 4.1× |
| Easyjson | 2528 | 195.42 MB/s | 448 | 3 | 2.6× |
| Goccy | 2746 | 179.87 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 3402 | 145.23 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6575 | 75.13 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 463 | 496.89 MB/s | 160 | 1 | 10.1× |
| LightningDestructive | 464 | 496.07 MB/s | 160 | 1 | 10.1× |
| Sonic | 919 | 250.27 MB/s | 801 | 8 | 5.1× |
| SonicFastest | 936 | 245.71 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1323 | 173.08 MB/s | 1296 | 26 | 3.5× |
| LightningArena | 1328 | 173.18 MB/s | 4096 | 1 | 3.5× |
| Easyjson | 1667 | 138.00 MB/s | 448 | 3 | 2.8× |
| Goccy | 1769 | 130.04 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2621 | 87.76 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4668 | 49.27 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 65764 | 990.39 MB/s | 103441 | 103 | 10.3× |
| Lightning | 66743 | 975.86 MB/s | 103441 | 103 | 10.1× |
| LightningDestructive | 70594 | 922.62 MB/s | 97220 | 98 | 9.6× |
| Sonic | 148665 | 438.11 MB/s | 235880 | 65 | 4.5× |
| SonicFastest | 157201 | 414.32 MB/s | 236137 | 65 | 4.3× |
| LightningDecodeAny | 198406 | 268.79 MB/s | 180048 | 3245 | 3.4× |
| Goccy | 215938 | 301.62 MB/s | 229217 | 134 | 3.1× |
| JSONV2 | 288809 | 225.52 MB/s | 206664 | 607 | 2.3× |
| Stdlib | 676416 | 96.29 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2510867 | 772.83 MB/s | 2864593 | 1380 | 10.9× |
| LightningArena | 2551225 | 760.60 MB/s | 2864593 | 1380 | 10.7× |
| Lightning | 2567689 | 755.73 MB/s | 2864595 | 1380 | 10.7× |
| SonicFastest | 4773164 | 406.54 MB/s | 4875033 | 1736 | 5.7× |
| Sonic | 4835458 | 401.30 MB/s | 4875919 | 1736 | 5.7× |
| Goccy | 5147366 | 376.98 MB/s | 4062939 | 13509 | 5.3× |
| Easyjson | 8027297 | 241.73 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 9637635 | 201.34 MB/s | 7063040 | 218633 | 2.8× |
| JSONV2 | 13125550 | 147.84 MB/s | 3237199 | 13947 | 2.1× |
| Stdlib | 27388303 | 70.85 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1049230 | 3171.69 MB/s | 351704 | 1286 | 22.4× |
| Lightning | 1528286 | 2177.49 MB/s | 2488907 | 2995 | 15.4× |
| LightningArena | 1537746 | 2164.10 MB/s | 2488907 | 2995 | 15.3× |
| Sonic | 2004245 | 1660.39 MB/s | 5896218 | 4263 | 11.7× |
| SonicFastest | 2005218 | 1659.59 MB/s | 5896565 | 4263 | 11.7× |
| LightningDecodeAny | 3361836 | 914.31 MB/s | 4876914 | 56892 | 7.0× |
| Goccy | 4979408 | 668.32 MB/s | 3948914 | 3817 | 4.7× |
| JSONV2 | 7447599 | 446.83 MB/s | 5364511 | 13243 | 3.2× |
| Stdlib | 23474731 | 141.76 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 217208 | 1014.45 MB/s | 135872 | 226 | 11.0× |
| Lightning | 219270 | 1004.91 MB/s | 135872 | 226 | 10.9× |
| LightningDestructive | 224012 | 983.63 MB/s | 135872 | 226 | 10.7× |
| SonicFastest | 475099 | 463.79 MB/s | 350747 | 262 | 5.0× |
| Sonic | 476657 | 462.27 MB/s | 350816 | 262 | 5.0× |
| Goccy | 479348 | 459.68 MB/s | 363929 | 1066 | 5.0× |
| Easyjson | 627479 | 351.16 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 825944 | 266.78 MB/s | 129745 | 470 | 2.9× |
| LightningDecodeAny | 1000005 | 108.31 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2396669 | 91.94 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12493483 | 648.34 MB/s | 11845072 | 20816 | 8.7× |
| LightningArena | 12617541 | 641.97 MB/s | 11845072 | 20816 | 8.6× |
| Lightning | 12736537 | 635.97 MB/s | 11845079 | 20816 | 8.5× |
| Sonic | 17475747 | 463.50 MB/s | 19859472 | 41640 | 6.2× |
| SonicFastest | 17624471 | 459.59 MB/s | 19859524 | 41640 | 6.2× |
| Goccy | 26290263 | 308.10 MB/s | 18794287 | 107154 | 4.1× |
| Easyjson | 34899495 | 232.10 MB/s | 15059618 | 41643 | 3.1× |
| LightningDecodeAny | 39508097 | 131.70 MB/s | 46279349 | 747112 | 2.8× |
| JSONV2 | 51292133 | 157.92 MB/s | 15233698 | 78972 | 2.1× |
| Stdlib | 108721910 | 74.50 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5405470 | 551.93 MB/s | 3764712 | 1504 | 10.6× |
| LightningDestructive | 5617297 | 531.12 MB/s | 3758856 | 29356 | 10.2× |
| Lightning | 5781879 | 516.00 MB/s | 3758857 | 29356 | 9.9× |
| Sonic | 9087615 | 328.30 MB/s | 9130940 | 57804 | 6.3× |
| SonicFastest | 9106324 | 327.63 MB/s | 9132412 | 57804 | 6.3× |
| LightningDecodeAny | 18523157 | 99.02 MB/s | 23982579 | 351152 | 3.1× |
| Goccy | 18662726 | 159.86 MB/s | 9795316 | 273616 | 3.1× |
| Easyjson | 19485383 | 153.11 MB/s | 9479441 | 30115 | 2.9× |
| JSONV2 | 27349070 | 109.09 MB/s | 9257077 | 86278 | 2.1× |
| Stdlib | 57151018 | 52.20 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1307423 | 553.45 MB/s | 907601 | 3618 | 10.5× |
| LightningArena | 1331889 | 543.29 MB/s | 911392 | 30 | 10.3× |
| Lightning | 1379005 | 524.72 MB/s | 907595 | 3618 | 9.9× |
| SonicFastest | 2101796 | 344.28 MB/s | 2371505 | 3683 | 6.5× |
| Sonic | 2117587 | 341.71 MB/s | 2371817 | 3683 | 6.5× |
| Easyjson | 5336406 | 135.60 MB/s | 2847907 | 3698 | 2.6× |
| LightningDecodeAny | 5351600 | 121.57 MB/s | 6500461 | 76546 | 2.6× |
| Goccy | 5423618 | 133.42 MB/s | 2735645 | 80269 | 2.5× |
| JSONV2 | 6897724 | 104.90 MB/s | 2704704 | 7318 | 2.0× |
| Stdlib | 13717508 | 52.75 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1845709 | 854.61 MB/s | 911393 | 30 | 10.3× |
| Lightning | 1895766 | 832.04 MB/s | 907596 | 3618 | 10.1× |
| LightningDestructive | 1933132 | 815.96 MB/s | 907600 | 3618 | 9.9× |
| Sonic | 2392162 | 659.38 MB/s | 3225577 | 3683 | 8.0× |
| SonicFastest | 2402117 | 656.65 MB/s | 3224945 | 3683 | 8.0× |
| LightningDecodeAny | 4717320 | 159.71 MB/s | 6500454 | 76546 | 4.0× |
| Easyjson | 6253608 | 252.23 MB/s | 2847905 | 3698 | 3.1× |
| Goccy | 6519422 | 241.95 MB/s | 3527587 | 80264 | 2.9× |
| JSONV2 | 7471102 | 211.13 MB/s | 2704553 | 7318 | 2.6× |
| Stdlib | 19099239 | 82.59 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 235020 | 638.77 MB/s | 81920 | 1 | 8.9× |
| Lightning | 236065 | 635.94 MB/s | 81920 | 1 | 8.9× |
| LightningDestructive | 244466 | 614.09 MB/s | 81920 | 1 | 8.6× |
| Sonic | 401618 | 373.80 MB/s | 408423 | 16 | 5.2× |
| SonicFastest | 428987 | 349.95 MB/s | 408488 | 16 | 4.9× |
| LightningDecodeAny | 580458 | 258.63 MB/s | 745765 | 10016 | 3.6× |
| Goccy | 1029403 | 145.84 MB/s | 331030 | 10005 | 2.0× |
| JSONV2 | 1116486 | 134.46 MB/s | 357725 | 20 | 1.9× |
| Stdlib | 2094837 | 71.66 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32498 | 865.18 MB/s | 29216 | 103 | 10.8× |
| LightningArena | 32631 | 861.66 MB/s | 29216 | 103 | 10.7× |
| LightningDestructive | 33968 | 827.75 MB/s | 29088 | 101 | 10.3× |
| Sonic | 57132 | 492.14 MB/s | 59468 | 83 | 6.1× |
| SonicFastest | 57602 | 488.12 MB/s | 59494 | 83 | 6.1× |
| Easyjson | 79020 | 355.82 MB/s | 32304 | 138 | 4.4× |
| Goccy | 85077 | 330.49 MB/s | 59259 | 188 | 4.1× |
| JSONV2 | 147048 | 191.21 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 163905 | 171.54 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 349372 | 80.48 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1990 | 1169.95 MB/s | 32 | 1 | 13.1× |
| LightningArena | 1990 | 1169.66 MB/s | 32 | 1 | 13.1× |
| LightningDestructive | 2138 | 1088.63 MB/s | 32 | 1 | 12.2× |
| Sonic | 4716 | 493.67 MB/s | 3711 | 4 | 5.5× |
| SonicFastest | 4735 | 491.61 MB/s | 3712 | 4 | 5.5× |
| Goccy | 4910 | 474.14 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 5296 | 439.55 MB/s | 192 | 2 | 4.9× |
| JSONV2 | 8553 | 272.17 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10393 | 162.13 MB/s | 10200 | 195 | 2.5× |
| Stdlib | 26134 | 89.08 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 203 | 932.98 MB/s | 0 | 0 | 13.7× |
| LightningArena | 203 | 929.66 MB/s | 0 | 0 | 13.7× |
| LightningDestructive | 211 | 897.38 MB/s | 0 | 0 | 13.2× |
| Goccy | 457 | 413.41 MB/s | 304 | 2 | 6.1× |
| Easyjson | 555 | 340.33 MB/s | 0 | 0 | 5.0× |
| Sonic | 648 | 291.67 MB/s | 341 | 3 | 4.3× |
| SonicFastest | 649 | 291.17 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1064 | 177.71 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1306 | 102.61 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2781 | 67.95 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1383 | 1584.73 MB/s | 0 | 0 | 13.8× |
| LightningArena | 1391 | 1574.72 MB/s | 0 | 0 | 13.7× |
| LightningDestructive | 1494 | 1466.05 MB/s | 0 | 0 | 12.7× |
| Easyjson | 3572 | 613.34 MB/s | 24 | 1 | 5.3× |
| Goccy | 3905 | 561.13 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6529 | 335.56 MB/s | 3601 | 38 | 2.9× |
| Sonic | 6728 | 325.65 MB/s | 3602 | 38 | 2.8× |
| JSONV2 | 8613 | 254.38 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9058 | 199.94 MB/s | 7552 | 158 | 2.1× |
| Stdlib | 19040 | 115.07 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 585239 | 872.25 MB/s | 457536 | 1009 | 12.3× |
| Lightning | 616128 | 828.52 MB/s | 457536 | 1009 | 11.7× |
| LightningArena | 618652 | 825.14 MB/s | 457536 | 1009 | 11.6× |
| Goccy | 1234849 | 413.39 MB/s | 1138235 | 5006 | 5.8× |
| SonicFastest | 1243706 | 410.45 MB/s | 1308385 | 2014 | 5.8× |
| Sonic | 1244621 | 410.15 MB/s | 1308624 | 2014 | 5.8× |
| Easyjson | 1744218 | 292.67 MB/s | 863779 | 3012 | 4.1× |
| JSONV2 | 3460842 | 147.50 MB/s | 1075964 | 12645 | 2.1× |
| LightningDecodeAny | 3533873 | 130.58 MB/s | 2950650 | 64018 | 2.0× |
| Stdlib | 7190573 | 70.99 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 643 | 30793.62 MB/s | 0 | 0 | 259.3× |
| LightningArena | 644 | 30707.75 MB/s | 0 | 0 | 258.6× |
| LightningDestructive | 919 | 21528.88 MB/s | 0 | 0 | 181.3× |
| SonicFastest | 6377 | 3103.18 MB/s | 21156 | 3 | 26.1× |
| Goccy | 23232 | 851.81 MB/s | 20492 | 2 | 7.2× |
| Sonic | 31693 | 624.40 MB/s | 20600 | 3 | 5.3× |
| JSONV2 | 33998 | 582.07 MB/s | 8 | 1 | 4.9× |
| LightningDecodeAny | 96368 | 205.34 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 105537 | 187.51 MB/s | 0 | 0 | 1.6× |
| Stdlib | 166620 | 118.77 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 2205 | 8218.64 MB/s | 432 | 2 | 56.8× |
| Lightning | 2209 | 8205.87 MB/s | 432 | 2 | 56.6× |
| LightningDestructive | 2465 | 7351.54 MB/s | 0 | 0 | 50.8× |
| Easyjson | 4746 | 3818.66 MB/s | 432 | 2 | 26.4× |
| SonicFastest | 8402 | 2157.04 MB/s | 20425 | 5 | 14.9× |
| Sonic | 8500 | 2132.28 MB/s | 20420 | 5 | 14.7× |
| LightningDecodeAny | 18963 | 943.00 MB/s | 29088 | 191 | 6.6× |
| Goccy | 24704 | 733.66 MB/s | 19460 | 2 | 5.1× |
| JSONV2 | 50131 | 361.54 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 125139 | 144.83 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2370845 | 847.16 MB/s | 3089565 | 6821 | 9.1× |
| LightningArena | 2440515 | 822.98 MB/s | 3094370 | 6703 | 8.9× |
| Lightning | 2464036 | 815.12 MB/s | 3091277 | 6827 | 8.8× |
| SonicFastest | 4163147 | 482.45 MB/s | 5152215 | 7085 | 5.2× |
| Sonic | 4239616 | 473.74 MB/s | 5153722 | 7085 | 5.1× |
| Goccy | 4650779 | 431.86 MB/s | 5410866 | 15832 | 4.7× |
| Easyjson | 5274946 | 380.76 MB/s | 2981491 | 7439 | 4.1× |
| LightningDecodeAny | 7000822 | 163.17 MB/s | 8503511 | 134008 | 3.1× |
| JSONV2 | 7572856 | 265.22 MB/s | 3173677 | 14563 | 2.9× |
| Stdlib | 21687844 | 92.61 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 838 | 655.29 MB/s | 480 | 1 | 7.9× |
| LightningArena | 843 | 651.47 MB/s | 480 | 1 | 7.9× |
| LightningDestructive | 849 | 646.50 MB/s | 480 | 1 | 7.8× |
| LightningDecodeAny | 1829 | 299.55 MB/s | 2021 | 46 | 3.6× |
| Easyjson | 2226 | 246.60 MB/s | 1616 | 5 | 3.0× |
| Sonic | 2395 | 229.23 MB/s | 2261 | 8 | 2.8× |
| SonicFastest | 2426 | 226.32 MB/s | 2261 | 8 | 2.7× |
| Goccy | 3288 | 166.98 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3424 | 160.33 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6655 | 82.49 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 502615 | 1256.46 MB/s | 402728 | 545 | 12.8× |
| Lightning | 571139 | 1105.71 MB/s | 451257 | 857 | 11.3× |
| LightningArena | 571844 | 1104.35 MB/s | 453017 | 712 | 11.2× |
| SonicFastest | 965927 | 653.79 MB/s | 1064910 | 814 | 6.7× |
| Sonic | 968114 | 652.31 MB/s | 1065040 | 814 | 6.6× |
| Easyjson | 1322700 | 477.44 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1383067 | 456.60 MB/s | 988210 | 1200 | 4.6× |
| JSONV2 | 2356709 | 267.96 MB/s | 571591 | 3144 | 2.7× |
| LightningDecodeAny | 2504456 | 186.43 MB/s | 2076503 | 30126 | 2.6× |
| Stdlib | 6428415 | 98.24 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 679460 | 827.73 MB/s | 546569 | 429 | 8.8× |
| Lightning | 863079 | 651.63 MB/s | 769936 | 1235 | 6.9× |
| LightningArena | 865662 | 649.69 MB/s | 771665 | 1088 | 6.9× |
| Sonic | 1293050 | 434.95 MB/s | 1349183 | 1185 | 4.6× |
| SonicFastest | 1295823 | 434.02 MB/s | 1348731 | 1185 | 4.6× |
| Goccy | 1588135 | 354.13 MB/s | 1036468 | 1028 | 3.8× |
| Easyjson | 2108443 | 266.74 MB/s | 775153 | 1254 | 2.8× |
| LightningDecodeAny | 2949190 | 190.70 MB/s | 2180441 | 30126 | 2.0× |
| JSONV2 | 3203273 | 175.57 MB/s | 927405 | 3482 | 1.9× |
| Stdlib | 5989153 | 93.90 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 640875 | 831.95 MB/s | 333416 | 2084 | 10.1× |
| LightningArena | 695379 | 766.74 MB/s | 368224 | 2293 | 9.3× |
| Lightning | 698569 | 763.24 MB/s | 368224 | 2293 | 9.3× |
| Sonic | 1118130 | 476.85 MB/s | 980928 | 3082 | 5.8× |
| SonicFastest | 1121000 | 475.63 MB/s | 981045 | 3082 | 5.8× |
| Easyjson | 1274652 | 418.29 MB/s | 428362 | 3273 | 5.1× |
| Goccy | 1511845 | 352.67 MB/s | 1167060 | 5408 | 4.3× |
| JSONV2 | 2809298 | 189.79 MB/s | 745425 | 13288 | 2.3× |
| LightningDecodeAny | 3507676 | 152.00 MB/s | 2992875 | 50076 | 1.8× |
| Stdlib | 6480910 | 82.27 MB/s | 798693 | 17133 | 1.0× |
