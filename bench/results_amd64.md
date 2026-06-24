# JSON Deserialization Benchmarks

- generated 2026-06-24T11:19:56Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 103429 | 1230.56 MB/s | 49760 | 3 | 12.8× |
| LightningDestructive | 104851 | 1213.87 MB/s | 49280 | 2 | 12.6× |
| SonicFastest | 193230 | 658.67 MB/s | 214151 | 15 | 6.8× |
| Sonic | 193911 | 656.36 MB/s | 214238 | 15 | 6.8× |
| Goccy | 240443 | 529.34 MB/s | 225497 | 884 | 5.5× |
| Easyjson | 252761 | 503.54 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 428267 | 297.19 MB/s | 195128 | 1805 | 3.1× |
| LightningDecodeAny | 466182 | 203.04 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1323347 | 96.18 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4074283 | 552.50 MB/s | 3122873 | 3081 | 8.3× |
| Lightning | 4130363 | 545.00 MB/s | 3122872 | 3081 | 8.2× |
| SonicFastest | 4903152 | 459.10 MB/s | 4871790 | 2584 | 6.9× |
| Sonic | 5100335 | 441.35 MB/s | 4870946 | 2584 | 6.7× |
| Goccy | 11826867 | 190.33 MB/s | 4189008 | 56535 | 2.9× |
| LightningDecodeAny | 12237577 | 183.94 MB/s | 7938299 | 281383 | 2.8× |
| Easyjson | 13797497 | 163.15 MB/s | 3099808 | 2120 | 2.5× |
| JSONV2 | 17430587 | 129.14 MB/s | 3123189 | 3083 | 2.0× |
| Stdlib | 34009092 | 66.19 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 565657 | 478.03 MB/s | 348025 | 1627 | 7.7× |
| LightningDestructive | 568004 | 476.06 MB/s | 348025 | 1627 | 7.7× |
| Sonic | 741878 | 364.48 MB/s | 641006 | 1147 | 5.9× |
| SonicFastest | 742051 | 364.40 MB/s | 641143 | 1147 | 5.9× |
| Goccy | 1646367 | 164.24 MB/s | 541054 | 8121 | 2.7× |
| LightningDecodeAny | 1690050 | 160.00 MB/s | 1011489 | 37901 | 2.6× |
| Easyjson | 1788061 | 151.23 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2348067 | 115.16 MB/s | 348159 | 1628 | 1.9× |
| Stdlib | 4363620 | 61.97 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1279030 | 1350.40 MB/s | 959848 | 5881 | 13.8× |
| Lightning | 1336964 | 1291.88 MB/s | 959889 | 5882 | 13.2× |
| Sonic | 2112183 | 817.73 MB/s | 2694940 | 5547 | 8.3× |
| SonicFastest | 2113350 | 817.28 MB/s | 2695080 | 5547 | 8.3× |
| Goccy | 2478513 | 696.87 MB/s | 2582002 | 14603 | 7.1× |
| Easyjson | 3847887 | 448.87 MB/s | 972032 | 5389 | 4.6× |
| LightningDecodeAny | 4380736 | 114.20 MB/s | 4976205 | 81466 | 4.0× |
| JSONV2 | 4551420 | 379.49 MB/s | 1011615 | 7594 | 3.9× |
| Stdlib | 17626735 | 97.99 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1024 | 1768.72 MB/s | 0 | 0 | 15.6× |
| LightningDestructive | 1086 | 1668.86 MB/s | 0 | 0 | 14.7× |
| Easyjson | 3033 | 597.53 MB/s | 24 | 1 | 5.3× |
| Goccy | 3186 | 568.74 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6310 | 287.18 MB/s | 3347 | 38 | 2.5× |
| Sonic | 6453 | 280.81 MB/s | 3345 | 38 | 2.5× |
| JSONV2 | 8002 | 226.43 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9395 | 192.77 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16015 | 113.14 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1118 | 1620.34 MB/s | 0 | 0 | 14.7× |
| LightningDestructive | 1188 | 1525.79 MB/s | 0 | 0 | 13.8× |
| Easyjson | 3003 | 603.48 MB/s | 24 | 1 | 5.5× |
| Goccy | 3200 | 566.23 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 6023 | 300.85 MB/s | 3344 | 38 | 2.7× |
| Sonic | 6265 | 289.21 MB/s | 3346 | 38 | 2.6× |
| JSONV2 | 7933 | 228.41 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 9161 | 197.68 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 16402 | 110.48 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1315 | 1378.00 MB/s | 144 | 10 | 12.2× |
| LightningDestructive | 1378 | 1314.86 MB/s | 144 | 10 | 11.6× |
| Easyjson | 3108 | 583.03 MB/s | 144 | 10 | 5.2× |
| Goccy | 3423 | 529.32 MB/s | 2600 | 5 | 4.7× |
| SonicFastest | 6197 | 292.41 MB/s | 3363 | 40 | 2.6× |
| Sonic | 6409 | 282.75 MB/s | 3364 | 40 | 2.5× |
| JSONV2 | 8635 | 209.83 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9419 | 192.26 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16043 | 112.95 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 722 | 684.04 MB/s | 160 | 1 | 9.8× |
| LightningDestructive | 734 | 672.94 MB/s | 160 | 1 | 9.6× |
| Sonic | 1261 | 391.70 MB/s | 1076 | 8 | 5.6× |
| SonicFastest | 1261 | 391.84 MB/s | 1076 | 8 | 5.6× |
| LightningDecodeAny | 1799 | 274.05 MB/s | 1536 | 30 | 3.9× |
| Goccy | 2567 | 192.43 MB/s | 856 | 23 | 2.7× |
| Easyjson | 2652 | 186.29 MB/s | 448 | 3 | 2.7× |
| JSONV2 | 3477 | 142.08 MB/s | 528 | 7 | 2.0× |
| Stdlib | 7050 | 70.07 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 490 | 469.10 MB/s | 160 | 1 | 10.0× |
| LightningDestructive | 491 | 468.76 MB/s | 160 | 1 | 10.0× |
| SonicFastest | 912 | 252.12 MB/s | 801 | 8 | 5.4× |
| Sonic | 916 | 251.02 MB/s | 800 | 8 | 5.4× |
| LightningDecodeAny | 1585 | 144.52 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1767 | 130.16 MB/s | 448 | 3 | 2.8× |
| Goccy | 1840 | 125.00 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2518 | 91.35 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4912 | 46.83 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 87480 | 744.54 MB/s | 172433 | 107 | 7.6× |
| LightningDestructive | 92314 | 705.55 MB/s | 166213 | 102 | 7.2× |
| SonicFastest | 137981 | 472.04 MB/s | 235657 | 65 | 4.8× |
| Sonic | 138476 | 470.35 MB/s | 235721 | 65 | 4.8× |
| Goccy | 172339 | 377.93 MB/s | 227809 | 134 | 3.9× |
| LightningDecodeAny | 196653 | 271.18 MB/s | 176960 | 3252 | 3.4× |
| JSONV2 | 256748 | 253.68 MB/s | 206662 | 607 | 2.6× |
| Stdlib | 667505 | 97.58 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2487501 | 780.09 MB/s | 2846864 | 2238 | 11.2× |
| Lightning | 2563832 | 756.86 MB/s | 2846866 | 2238 | 10.9× |
| Sonic | 4772199 | 406.62 MB/s | 4880763 | 1736 | 5.9× |
| SonicFastest | 4794870 | 404.70 MB/s | 4880429 | 1736 | 5.8× |
| Goccy | 4879144 | 397.71 MB/s | 4062312 | 13509 | 5.7× |
| Easyjson | 8472947 | 229.02 MB/s | 3871264 | 15043 | 3.3× |
| LightningDecodeAny | 9992755 | 194.19 MB/s | 7013525 | 219937 | 2.8× |
| JSONV2 | 12465857 | 155.66 MB/s | 3237183 | 13947 | 2.2× |
| Stdlib | 27919454 | 69.50 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1068094 | 3115.67 MB/s | 351704 | 1286 | 22.0× |
| Lightning | 1614395 | 2061.35 MB/s | 2488907 | 2995 | 14.5× |
| SonicFastest | 2018010 | 1649.07 MB/s | 5896313 | 4263 | 11.6× |
| Sonic | 2021727 | 1646.03 MB/s | 5896041 | 4263 | 11.6× |
| LightningDecodeAny | 3532602 | 870.11 MB/s | 4886622 | 56892 | 6.6× |
| Goccy | 6061212 | 549.04 MB/s | 3948914 | 3816 | 3.9× |
| JSONV2 | 7587258 | 438.61 MB/s | 5364508 | 13243 | 3.1× |
| Stdlib | 23453900 | 141.89 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 223102 | 987.65 MB/s | 136896 | 228 | 11.0× |
| LightningDestructive | 229915 | 958.38 MB/s | 136896 | 228 | 10.6× |
| SonicFastest | 480232 | 458.83 MB/s | 350936 | 262 | 5.1× |
| Goccy | 483482 | 455.75 MB/s | 363552 | 1066 | 5.1× |
| Sonic | 485993 | 453.39 MB/s | 351602 | 262 | 5.0× |
| Easyjson | 653676 | 337.09 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 695049 | 317.02 MB/s | 129747 | 470 | 3.5× |
| LightningDecodeAny | 1000361 | 108.27 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2446688 | 90.06 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 15029254 | 538.95 MB/s | 15059832 | 51649 | 7.4× |
| Lightning | 15384756 | 526.50 MB/s | 15059837 | 51649 | 7.3× |
| Sonic | 17800507 | 455.05 MB/s | 19864681 | 41640 | 6.3× |
| SonicFastest | 17839376 | 454.05 MB/s | 19864568 | 41640 | 6.3× |
| Goccy | 25009631 | 323.88 MB/s | 19196761 | 107156 | 4.5× |
| Easyjson | 35879206 | 225.76 MB/s | 15059618 | 41643 | 3.1× |
| LightningDecodeAny | 42947337 | 121.15 MB/s | 34333349 | 912810 | 2.6× |
| JSONV2 | 49718632 | 162.92 MB/s | 15233745 | 78972 | 2.2× |
| Stdlib | 111669654 | 72.54 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6120157 | 487.48 MB/s | 3985336 | 30015 | 9.3× |
| Lightning | 6193396 | 481.72 MB/s | 3985336 | 30015 | 9.2× |
| Sonic | 8933136 | 333.98 MB/s | 9131069 | 57804 | 6.4× |
| SonicFastest | 8997692 | 331.58 MB/s | 9132572 | 57804 | 6.3× |
| Goccy | 18401110 | 162.14 MB/s | 9851873 | 273619 | 3.1× |
| Easyjson | 18462509 | 161.60 MB/s | 9479440 | 30115 | 3.1× |
| LightningDecodeAny | 20730225 | 88.48 MB/s | 20023835 | 408557 | 2.8× |
| JSONV2 | 26596902 | 112.17 MB/s | 9257033 | 86278 | 2.1× |
| Stdlib | 57065162 | 52.28 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1449319 | 499.27 MB/s | 907392 | 3618 | 9.6× |
| Lightning | 1534648 | 471.51 MB/s | 907387 | 3618 | 9.1× |
| SonicFastest | 2140869 | 337.99 MB/s | 2370627 | 3683 | 6.5× |
| Sonic | 2145816 | 337.21 MB/s | 2371910 | 3683 | 6.5× |
| Easyjson | 5326778 | 135.84 MB/s | 2847905 | 3698 | 2.6× |
| Goccy | 5430227 | 133.25 MB/s | 2732032 | 80269 | 2.6× |
| LightningDecodeAny | 5545062 | 117.32 MB/s | 5752204 | 80172 | 2.5× |
| JSONV2 | 6210141 | 116.52 MB/s | 2704706 | 7318 | 2.2× |
| Stdlib | 13958047 | 51.84 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2019363 | 781.11 MB/s | 907388 | 3618 | 9.9× |
| LightningDestructive | 2031365 | 776.50 MB/s | 907392 | 3618 | 9.8× |
| SonicFastest | 2507635 | 629.02 MB/s | 3237159 | 3683 | 8.0× |
| Sonic | 2522794 | 625.24 MB/s | 3237189 | 3683 | 7.9× |
| LightningDecodeAny | 4772149 | 157.87 MB/s | 5752200 | 80172 | 4.2× |
| Easyjson | 6241110 | 252.74 MB/s | 2847904 | 3698 | 3.2× |
| Goccy | 6686860 | 235.89 MB/s | 3455891 | 80260 | 3.0× |
| JSONV2 | 6907473 | 228.35 MB/s | 2704550 | 7318 | 2.9× |
| Stdlib | 19963578 | 79.01 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 249456 | 601.81 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 255177 | 588.31 MB/s | 81920 | 1 | 8.4× |
| Sonic | 418783 | 358.48 MB/s | 408728 | 16 | 5.1× |
| SonicFastest | 421080 | 356.52 MB/s | 408623 | 16 | 5.1× |
| LightningDecodeAny | 595391 | 252.14 MB/s | 746007 | 10020 | 3.6× |
| Goccy | 1036621 | 144.82 MB/s | 329402 | 10005 | 2.1× |
| JSONV2 | 1110661 | 135.17 MB/s | 357724 | 20 | 1.9× |
| Stdlib | 2148562 | 69.87 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33750 | 833.08 MB/s | 30272 | 105 | 10.1× |
| LightningDestructive | 34831 | 807.23 MB/s | 30144 | 103 | 9.8× |
| Sonic | 55906 | 502.94 MB/s | 59509 | 83 | 6.1× |
| SonicFastest | 56197 | 500.33 MB/s | 59503 | 83 | 6.1× |
| Goccy | 79002 | 355.90 MB/s | 59285 | 188 | 4.3× |
| Easyjson | 81121 | 346.60 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 132781 | 211.75 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 167134 | 168.23 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 342239 | 82.16 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1951 | 1193.18 MB/s | 32 | 1 | 13.4× |
| LightningDestructive | 2093 | 1112.12 MB/s | 32 | 1 | 12.5× |
| Sonic | 4731 | 492.04 MB/s | 3706 | 4 | 5.5× |
| SonicFastest | 4755 | 489.59 MB/s | 3711 | 4 | 5.5× |
| Goccy | 4816 | 483.35 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 5026 | 463.20 MB/s | 192 | 2 | 5.2× |
| JSONV2 | 8343 | 279.04 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10440 | 161.40 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 26059 | 89.34 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 215 | 879.48 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 224 | 844.45 MB/s | 0 | 0 | 12.5× |
| Goccy | 420 | 449.60 MB/s | 304 | 2 | 6.6× |
| Easyjson | 592 | 319.22 MB/s | 0 | 0 | 4.7× |
| SonicFastest | 628 | 300.80 MB/s | 341 | 3 | 4.4× |
| Sonic | 634 | 297.91 MB/s | 341 | 3 | 4.4× |
| JSONV2 | 1030 | 183.53 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1343 | 99.75 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2790 | 67.73 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1489 | 1471.28 MB/s | 0 | 0 | 12.7× |
| LightningDestructive | 1551 | 1412.89 MB/s | 0 | 0 | 12.2× |
| Goccy | 3588 | 610.71 MB/s | 2864 | 4 | 5.3× |
| Easyjson | 3781 | 579.53 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6544 | 334.81 MB/s | 3600 | 38 | 2.9× |
| Sonic | 6652 | 329.38 MB/s | 3601 | 38 | 2.8× |
| JSONV2 | 8137 | 269.25 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 9225 | 196.31 MB/s | 7536 | 158 | 2.1× |
| Stdlib | 18930 | 115.74 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 691125 | 738.62 MB/s | 703778 | 1012 | 10.2× |
| Lightning | 725911 | 703.22 MB/s | 703778 | 1012 | 9.7× |
| SonicFastest | 1275324 | 400.27 MB/s | 1308395 | 2014 | 5.5× |
| Sonic | 1275348 | 400.26 MB/s | 1309203 | 2014 | 5.5× |
| Goccy | 1322398 | 386.02 MB/s | 1137104 | 5006 | 5.3× |
| Easyjson | 1665929 | 306.42 MB/s | 863780 | 3012 | 4.2× |
| JSONV2 | 3475990 | 146.86 MB/s | 1075961 | 12645 | 2.0× |
| LightningDecodeAny | 3576731 | 129.02 MB/s | 2785928 | 66022 | 2.0× |
| Stdlib | 7028908 | 72.63 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 642 | 30819.67 MB/s | 0 | 0 | 254.3× |
| LightningDestructive | 837 | 23646.30 MB/s | 0 | 0 | 195.1× |
| SonicFastest | 6687 | 2959.32 MB/s | 21175 | 3 | 24.4× |
| Goccy | 29626 | 667.96 MB/s | 20492 | 2 | 5.5× |
| Sonic | 31788 | 622.54 MB/s | 20609 | 3 | 5.1× |
| JSONV2 | 33278 | 594.66 MB/s | 8 | 1 | 4.9× |
| LightningDecodeAny | 101733 | 194.51 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 104286 | 189.76 MB/s | 0 | 0 | 1.6× |
| Stdlib | 163300 | 121.18 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2197 | 8248.43 MB/s | 432 | 2 | 54.6× |
| LightningDestructive | 2422 | 7481.55 MB/s | 0 | 0 | 49.5× |
| Easyjson | 4989 | 3632.52 MB/s | 432 | 2 | 24.0× |
| Sonic | 8553 | 2119.10 MB/s | 20473 | 5 | 14.0× |
| SonicFastest | 8786 | 2062.94 MB/s | 20465 | 5 | 13.6× |
| LightningDecodeAny | 18387 | 972.52 MB/s | 29088 | 191 | 6.5× |
| Goccy | 21213 | 854.36 MB/s | 19460 | 2 | 5.7× |
| JSONV2 | 54226 | 334.23 MB/s | 16500 | 50 | 2.2× |
| Stdlib | 119891 | 151.17 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2259253 | 889.01 MB/s | 2960389 | 7411 | 9.7× |
| Lightning | 2344242 | 856.78 MB/s | 2962101 | 7417 | 9.3× |
| SonicFastest | 4079393 | 492.35 MB/s | 5156590 | 7085 | 5.4× |
| Sonic | 4112489 | 488.39 MB/s | 5160105 | 7085 | 5.3× |
| Goccy | 4610949 | 435.59 MB/s | 5410851 | 15833 | 4.7× |
| Easyjson | 5542010 | 362.41 MB/s | 2981489 | 7439 | 3.9× |
| LightningDecodeAny | 7197140 | 158.71 MB/s | 7386072 | 134751 | 3.0× |
| JSONV2 | 7645710 | 262.70 MB/s | 3173680 | 14563 | 2.9× |
| Stdlib | 21866495 | 91.85 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 843 | 651.32 MB/s | 480 | 1 | 8.0× |
| Lightning | 847 | 648.23 MB/s | 480 | 1 | 7.9× |
| LightningDecodeAny | 2118 | 258.68 MB/s | 2261 | 50 | 3.2× |
| Easyjson | 2356 | 233.03 MB/s | 1616 | 5 | 2.9× |
| Sonic | 2379 | 230.80 MB/s | 2261 | 8 | 2.8× |
| SonicFastest | 2394 | 229.35 MB/s | 2262 | 8 | 2.8× |
| Goccy | 3421 | 160.47 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3587 | 153.04 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6727 | 81.61 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 493958 | 1278.48 MB/s | 364472 | 566 | 12.7× |
| Lightning | 564370 | 1118.97 MB/s | 413001 | 878 | 11.1× |
| Sonic | 963744 | 655.27 MB/s | 1065759 | 814 | 6.5× |
| SonicFastest | 965116 | 654.34 MB/s | 1065838 | 814 | 6.5× |
| Easyjson | 1299465 | 485.98 MB/s | 422504 | 936 | 4.8× |
| Goccy | 1393753 | 453.10 MB/s | 989618 | 1201 | 4.5× |
| JSONV2 | 2328007 | 271.27 MB/s | 571590 | 3144 | 2.7× |
| LightningDecodeAny | 2560137 | 182.38 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6291839 | 100.37 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 808536 | 695.59 MB/s | 544249 | 448 | 7.4× |
| Lightning | 1022332 | 550.12 MB/s | 767622 | 1254 | 5.8× |
| Sonic | 1294643 | 434.41 MB/s | 1349899 | 1185 | 4.6× |
| SonicFastest | 1309952 | 429.33 MB/s | 1350246 | 1185 | 4.5× |
| Goccy | 1579330 | 356.11 MB/s | 1036696 | 1028 | 3.8× |
| Easyjson | 2113777 | 266.07 MB/s | 775153 | 1254 | 2.8× |
| LightningDecodeAny | 3057694 | 183.93 MB/s | 2114149 | 30295 | 1.9× |
| JSONV2 | 3143952 | 178.89 MB/s | 927409 | 3482 | 1.9× |
| Stdlib | 5947822 | 94.56 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 608140 | 876.74 MB/s | 333416 | 2084 | 10.2× |
| Lightning | 685816 | 777.44 MB/s | 368224 | 2293 | 9.1× |
| SonicFastest | 1078530 | 494.36 MB/s | 981886 | 3082 | 5.8× |
| Sonic | 1079063 | 494.11 MB/s | 981513 | 3082 | 5.8× |
| Easyjson | 1280568 | 416.36 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1450783 | 367.51 MB/s | 1167076 | 5408 | 4.3× |
| JSONV2 | 2809148 | 189.80 MB/s | 745420 | 13288 | 2.2× |
| LightningDecodeAny | 3474228 | 153.47 MB/s | 2674618 | 50596 | 1.8× |
| Stdlib | 6226173 | 85.63 MB/s | 798692 | 17133 | 1.0× |
