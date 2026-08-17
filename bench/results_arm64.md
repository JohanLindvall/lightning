# JSON Deserialization Benchmarks

- generated 2026-08-17T08:49:18Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 83775 | 1519.24 MB/s | 49760 | 3 | 13.2× |
| LightningArena | 83906 | 1516.88 MB/s | 49760 | 3 | 13.2× |
| LightningDestructive | 84369 | 1508.54 MB/s | 49280 | 2 | 13.1× |
| SonicFastest | 181868 | 699.82 MB/s | 194447 | 10 | 6.1× |
| Sonic | 184037 | 691.57 MB/s | 198485 | 10 | 6.0× |
| Goccy | 195174 | 652.11 MB/s | 224789 | 884 | 5.7× |
| Easyjson | 212796 | 598.11 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 418823 | 303.89 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 422910 | 223.81 MB/s | 463409 | 9708 | 2.6× |
| Stdlib | 1105062 | 115.17 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3576228 | 629.45 MB/s | 2532848 | 1143 | 7.3× |
| Lightning | 3587487 | 627.47 MB/s | 2532851 | 1143 | 7.3× |
| LightningArena | 3596428 | 625.91 MB/s | 2532848 | 1143 | 7.3× |
| SonicFastest | 4787891 | 470.16 MB/s | 15233824 | 970 | 5.4× |
| Sonic | 4869220 | 462.30 MB/s | 15233813 | 970 | 5.4× |
| Goccy | 10524582 | 213.89 MB/s | 4125462 | 56532 | 2.5× |
| Easyjson | 10937452 | 205.81 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12427855 | 181.13 MB/s | 19380210 | 223896 | 2.1× |
| JSONV2 | 16304822 | 138.06 MB/s | 3123207 | 3083 | 1.6× |
| Stdlib | 26075110 | 86.33 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 458226 | 590.11 MB/s | 397296 | 567 | 7.3× |
| LightningArena | 460370 | 587.36 MB/s | 397296 | 567 | 7.3× |
| Lightning | 461268 | 586.22 MB/s | 397296 | 567 | 7.3× |
| Sonic | 637330 | 424.27 MB/s | 476646 | 968 | 5.3× |
| SonicFastest | 641521 | 421.50 MB/s | 482033 | 968 | 5.2× |
| Easyjson | 1395886 | 193.71 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1424434 | 189.83 MB/s | 544164 | 8123 | 2.4× |
| LightningDecodeAny | 1636880 | 165.19 MB/s | 2543877 | 29687 | 2.0× |
| JSONV2 | 2134102 | 126.71 MB/s | 348152 | 1628 | 1.6× |
| Stdlib | 3352360 | 80.66 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1003026 | 1721.99 MB/s | 765560 | 2798 | 13.2× |
| Lightning | 1015371 | 1701.06 MB/s | 765603 | 2799 | 13.1× |
| LightningArena | 1019430 | 1694.28 MB/s | 768416 | 2440 | 13.0× |
| Sonic | 2072078 | 833.56 MB/s | 2664197 | 4020 | 6.4× |
| SonicFastest | 2077539 | 831.37 MB/s | 2720617 | 4020 | 6.4× |
| Goccy | 2410669 | 716.48 MB/s | 2582526 | 14604 | 5.5× |
| Easyjson | 4210479 | 410.22 MB/s | 972032 | 5389 | 3.2× |
| LightningDecodeAny | 4220035 | 118.55 MB/s | 4953694 | 76576 | 3.1× |
| JSONV2 | 4221579 | 409.14 MB/s | 1011631 | 7594 | 3.1× |
| Stdlib | 13264366 | 130.21 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 897 | 2019.71 MB/s | 0 | 0 | 15.6× |
| Lightning | 898 | 2018.57 MB/s | 0 | 0 | 15.6× |
| LightningDestructive | 911 | 1989.31 MB/s | 0 | 0 | 15.4× |
| Easyjson | 2541 | 713.08 MB/s | 24 | 1 | 5.5× |
| Goccy | 2845 | 636.87 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6100 | 297.07 MB/s | 3919 | 40 | 2.3× |
| Sonic | 6109 | 296.60 MB/s | 3922 | 40 | 2.3× |
| JSONV2 | 7838 | 231.18 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8181 | 221.38 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 14020 | 129.24 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 927 | 1954.68 MB/s | 0 | 0 | 15.1× |
| LightningArena | 929 | 1950.53 MB/s | 0 | 0 | 15.1× |
| LightningDestructive | 948 | 1910.47 MB/s | 0 | 0 | 14.8× |
| Easyjson | 2546 | 711.64 MB/s | 24 | 1 | 5.5× |
| Goccy | 2840 | 637.98 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6103 | 296.92 MB/s | 3780 | 40 | 2.3× |
| SonicFastest | 6109 | 296.61 MB/s | 3795 | 40 | 2.3× |
| JSONV2 | 7697 | 235.42 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7848 | 230.77 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 14027 | 129.18 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1117 | 1621.49 MB/s | 144 | 10 | 12.6× |
| Lightning | 1125 | 1611.34 MB/s | 144 | 10 | 12.5× |
| LightningDestructive | 1172 | 1546.48 MB/s | 144 | 10 | 12.0× |
| Easyjson | 2739 | 661.60 MB/s | 144 | 10 | 5.1× |
| Goccy | 2871 | 631.06 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6138 | 295.21 MB/s | 3835 | 42 | 2.3× |
| Sonic | 6143 | 294.97 MB/s | 3830 | 42 | 2.3× |
| LightningDecodeAny | 7862 | 230.35 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 8075 | 224.39 MB/s | 632 | 7 | 1.7× |
| Stdlib | 14024 | 129.21 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 728 | 678.19 MB/s | 160 | 1 | 7.6× |
| Lightning | 728 | 678.14 MB/s | 160 | 1 | 7.6× |
| Sonic | 1228 | 402.22 MB/s | 1027 | 6 | 4.5× |
| SonicFastest | 1237 | 399.44 MB/s | 1032 | 6 | 4.5× |
| LightningArena | 1389 | 355.62 MB/s | 4096 | 1 | 4.0× |
| LightningDecodeAny | 1434 | 343.75 MB/s | 1296 | 26 | 3.9× |
| Easyjson | 2227 | 221.80 MB/s | 448 | 3 | 2.5× |
| Goccy | 2407 | 205.22 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3242 | 152.36 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5528 | 89.36 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 432 | 532.79 MB/s | 160 | 1 | 9.5× |
| Lightning | 436 | 527.04 MB/s | 160 | 1 | 9.4× |
| Sonic | 876 | 262.56 MB/s | 656 | 6 | 4.7× |
| SonicFastest | 877 | 262.36 MB/s | 648 | 6 | 4.7× |
| LightningArena | 1075 | 213.91 MB/s | 4096 | 1 | 3.8× |
| LightningDecodeAny | 1145 | 199.91 MB/s | 1296 | 26 | 3.6× |
| Easyjson | 1396 | 164.80 MB/s | 448 | 3 | 2.9× |
| Goccy | 1557 | 147.73 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2442 | 94.20 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4107 | 56.00 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 51411 | 1266.89 MB/s | 97221 | 98 | 10.6× |
| Lightning | 52990 | 1229.13 MB/s | 103441 | 103 | 10.3× |
| LightningArena | 53837 | 1209.80 MB/s | 103441 | 103 | 10.2× |
| Sonic | 101452 | 642.00 MB/s | 160152 | 75 | 5.4× |
| SonicFastest | 101834 | 639.59 MB/s | 160241 | 75 | 5.4× |
| Goccy | 147431 | 441.78 MB/s | 229273 | 134 | 3.7× |
| LightningDecodeAny | 180995 | 294.64 MB/s | 180048 | 3245 | 3.0× |
| JSONV2 | 228296 | 285.30 MB/s | 206656 | 607 | 2.4× |
| Stdlib | 547175 | 119.03 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2304218 | 842.14 MB/s | 2864593 | 1380 | 10.2× |
| Lightning | 2392942 | 810.91 MB/s | 2864594 | 1380 | 9.8× |
| LightningArena | 2419501 | 802.01 MB/s | 2864593 | 1380 | 9.7× |
| Sonic | 4739555 | 409.42 MB/s | 14608618 | 1407 | 4.9× |
| SonicFastest | 4744341 | 409.01 MB/s | 14608621 | 1407 | 4.9× |
| Goccy | 4817996 | 402.75 MB/s | 4065508 | 13510 | 4.9× |
| Easyjson | 7418842 | 261.56 MB/s | 3871267 | 15043 | 3.2× |
| LightningDecodeAny | 9070683 | 213.93 MB/s | 7063040 | 218633 | 2.6× |
| JSONV2 | 11288997 | 171.89 MB/s | 3237215 | 13947 | 2.1× |
| Stdlib | 23454106 | 82.73 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 896474 | 3712.13 MB/s | 351704 | 1286 | 23.4× |
| Lightning | 1625200 | 2047.64 MB/s | 2488904 | 2995 | 12.9× |
| LightningArena | 1640225 | 2028.89 MB/s | 2488904 | 2995 | 12.8× |
| Sonic | 2786922 | 1194.09 MB/s | 6470726 | 4248 | 7.5× |
| SonicFastest | 2818719 | 1180.62 MB/s | 6464305 | 4248 | 7.4× |
| LightningDecodeAny | 3477255 | 883.96 MB/s | 4876911 | 56892 | 6.0× |
| Goccy | 4707039 | 706.99 MB/s | 3948909 | 3816 | 4.5× |
| JSONV2 | 7606935 | 437.47 MB/s | 5364512 | 13243 | 2.8× |
| Stdlib | 20949749 | 158.85 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 188434 | 1169.36 MB/s | 135872 | 226 | 10.8× |
| Lightning | 189204 | 1164.59 MB/s | 135872 | 226 | 10.8× |
| LightningDestructive | 190678 | 1155.59 MB/s | 135872 | 226 | 10.7× |
| SonicFastest | 384879 | 572.51 MB/s | 318559 | 398 | 5.3× |
| Sonic | 395480 | 557.16 MB/s | 343390 | 398 | 5.2× |
| Goccy | 436528 | 504.77 MB/s | 364468 | 1067 | 4.7× |
| Easyjson | 549214 | 401.20 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 724618 | 304.09 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 845510 | 128.10 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2042649 | 107.87 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 10545047 | 768.14 MB/s | 11845073 | 20816 | 8.5× |
| LightningArena | 10846081 | 746.82 MB/s | 11845072 | 20816 | 8.3× |
| Lightning | 10851003 | 746.48 MB/s | 11845083 | 20816 | 8.3× |
| SonicFastest | 17061715 | 474.75 MB/s | 70887620 | 40014 | 5.3× |
| Sonic | 17318589 | 467.71 MB/s | 70917119 | 40014 | 5.2× |
| Goccy | 24271037 | 333.73 MB/s | 17313127 | 107149 | 3.7× |
| Easyjson | 30747855 | 263.43 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 37076175 | 140.33 MB/s | 46279353 | 747112 | 2.4× |
| JSONV2 | 44012365 | 184.04 MB/s | 15233738 | 78972 | 2.0× |
| Stdlib | 89654204 | 90.35 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5237695 | 569.61 MB/s | 3764712 | 1504 | 9.0× |
| LightningDestructive | 5595549 | 533.19 MB/s | 3758856 | 29356 | 8.4× |
| Lightning | 5725854 | 521.05 MB/s | 3758856 | 29356 | 8.2× |
| SonicFastest | 8727406 | 341.85 MB/s | 26485908 | 56760 | 5.4× |
| Sonic | 8731317 | 341.70 MB/s | 26483425 | 56760 | 5.4× |
| Easyjson | 16457774 | 181.28 MB/s | 9479440 | 30115 | 2.9× |
| LightningDecodeAny | 16685975 | 109.92 MB/s | 23982584 | 351152 | 2.8× |
| Goccy | 16729517 | 178.34 MB/s | 10586264 | 273648 | 2.8× |
| JSONV2 | 24849471 | 120.06 MB/s | 9257187 | 86278 | 1.9× |
| Stdlib | 47185223 | 63.23 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1201811 | 602.09 MB/s | 911392 | 30 | 9.6× |
| LightningDestructive | 1234916 | 585.95 MB/s | 907601 | 3618 | 9.4× |
| Lightning | 1252211 | 577.86 MB/s | 907595 | 3618 | 9.2× |
| SonicFastest | 1792610 | 403.66 MB/s | 3183686 | 7226 | 6.5× |
| Sonic | 1799296 | 402.16 MB/s | 3181771 | 7226 | 6.4× |
| Easyjson | 4182339 | 173.01 MB/s | 2847907 | 3698 | 2.8× |
| LightningDecodeAny | 4331525 | 150.19 MB/s | 6500459 | 76546 | 2.7× |
| Goccy | 4860535 | 148.87 MB/s | 2835915 | 80275 | 2.4× |
| JSONV2 | 5582916 | 129.61 MB/s | 2704614 | 7318 | 2.1× |
| Stdlib | 11579545 | 62.49 MB/s | 2704553 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1718886 | 917.66 MB/s | 907601 | 3618 | 9.2× |
| LightningArena | 1721757 | 916.13 MB/s | 911393 | 30 | 9.1× |
| Lightning | 1771619 | 890.35 MB/s | 907595 | 3618 | 8.9× |
| Sonic | 2299387 | 685.99 MB/s | 5799447 | 7226 | 6.8× |
| SonicFastest | 2319675 | 679.99 MB/s | 5799279 | 7226 | 6.8× |
| LightningDecodeAny | 3968696 | 189.84 MB/s | 6500459 | 76546 | 4.0× |
| Easyjson | 5558385 | 283.78 MB/s | 2847906 | 3698 | 2.8× |
| Goccy | 5799120 | 272.00 MB/s | 3632925 | 80269 | 2.7× |
| JSONV2 | 6356010 | 248.17 MB/s | 2704596 | 7318 | 2.5× |
| Stdlib | 15740300 | 100.21 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 201747 | 744.12 MB/s | 81920 | 1 | 9.1× |
| LightningArena | 201785 | 743.98 MB/s | 81920 | 1 | 9.1× |
| Lightning | 201832 | 743.81 MB/s | 81920 | 1 | 9.1× |
| SonicFastest | 270210 | 555.58 MB/s | 249949 | 6 | 6.8× |
| Sonic | 270985 | 553.99 MB/s | 252963 | 6 | 6.8× |
| LightningDecodeAny | 482314 | 311.25 MB/s | 745764 | 10016 | 3.8× |
| Goccy | 870569 | 172.44 MB/s | 324522 | 10004 | 2.1× |
| JSONV2 | 1067775 | 140.60 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1843704 | 81.43 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 28721 | 978.97 MB/s | 29216 | 103 | 10.5× |
| LightningArena | 28893 | 973.14 MB/s | 29216 | 103 | 10.5× |
| LightningDestructive | 29192 | 963.19 MB/s | 29088 | 101 | 10.4× |
| Sonic | 64563 | 435.49 MB/s | 48795 | 103 | 4.7× |
| SonicFastest | 64946 | 432.93 MB/s | 48828 | 103 | 4.7× |
| Easyjson | 67956 | 413.75 MB/s | 32304 | 138 | 4.5× |
| Goccy | 72935 | 385.51 MB/s | 59228 | 188 | 4.1× |
| JSONV2 | 134511 | 209.03 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 149165 | 188.50 MB/s | 140576 | 2643 | 2.0× |
| Stdlib | 302460 | 92.96 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1551 | 1501.04 MB/s | 32 | 1 | 14.7× |
| LightningArena | 1559 | 1492.90 MB/s | 32 | 1 | 14.6× |
| LightningDestructive | 1627 | 1431.20 MB/s | 32 | 1 | 14.0× |
| Easyjson | 4210 | 552.98 MB/s | 192 | 2 | 5.4× |
| Goccy | 4232 | 550.10 MB/s | 3649 | 4 | 5.4× |
| Sonic | 5133 | 453.56 MB/s | 4285 | 6 | 4.4× |
| SonicFastest | 5163 | 450.90 MB/s | 4322 | 6 | 4.4× |
| JSONV2 | 8461 | 275.13 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9737 | 173.04 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22748 | 102.34 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 184 | 1026.70 MB/s | 0 | 0 | 13.1× |
| Lightning | 185 | 1022.84 MB/s | 0 | 0 | 13.1× |
| LightningDestructive | 187 | 1012.98 MB/s | 0 | 0 | 12.9× |
| Goccy | 380 | 497.70 MB/s | 304 | 2 | 6.4× |
| Easyjson | 496 | 380.99 MB/s | 0 | 0 | 4.9× |
| Sonic | 779 | 242.63 MB/s | 508 | 4 | 3.1× |
| SonicFastest | 783 | 241.35 MB/s | 508 | 4 | 3.1× |
| JSONV2 | 1026 | 184.14 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1174 | 114.10 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2413 | 78.33 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1195 | 1833.29 MB/s | 0 | 0 | 13.4× |
| LightningArena | 1201 | 1823.74 MB/s | 0 | 0 | 13.3× |
| LightningDestructive | 1222 | 1793.34 MB/s | 0 | 0 | 13.1× |
| Goccy | 3138 | 698.23 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3194 | 686.00 MB/s | 24 | 1 | 5.0× |
| Sonic | 6556 | 334.19 MB/s | 3928 | 40 | 2.4× |
| SonicFastest | 6557 | 334.14 MB/s | 3952 | 40 | 2.4× |
| LightningDecodeAny | 7836 | 231.11 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 7997 | 273.99 MB/s | 640 | 6 | 2.0× |
| Stdlib | 16026 | 136.71 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 557978 | 914.87 MB/s | 457537 | 1009 | 10.8× |
| Lightning | 567410 | 899.66 MB/s | 457537 | 1009 | 10.6× |
| LightningArena | 567933 | 898.83 MB/s | 457536 | 1009 | 10.6× |
| Sonic | 1153632 | 442.49 MB/s | 899792 | 2006 | 5.2× |
| SonicFastest | 1158107 | 440.78 MB/s | 901890 | 2006 | 5.2× |
| Goccy | 1158173 | 440.76 MB/s | 1139002 | 5006 | 5.2× |
| Easyjson | 1542479 | 330.95 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3174739 | 160.79 MB/s | 1076011 | 12646 | 1.9× |
| LightningDecodeAny | 3275364 | 140.89 MB/s | 2950648 | 64018 | 1.8× |
| Stdlib | 6006541 | 84.99 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 485 | 40795.27 MB/s | 0 | 0 | 230.1× |
| LightningArena | 485 | 40773.64 MB/s | 0 | 0 | 230.0× |
| LightningDestructive | 525 | 37723.41 MB/s | 0 | 0 | 212.7× |
| Goccy | 20494 | 965.58 MB/s | 20491 | 2 | 5.4× |
| SonicFastest | 27549 | 718.33 MB/s | 22861 | 4 | 4.1× |
| Sonic | 27552 | 718.25 MB/s | 22764 | 4 | 4.1× |
| JSONV2 | 29828 | 663.44 MB/s | 8 | 1 | 3.7× |
| LightningDecodeAny | 76140 | 259.89 MB/s | 116864 | 2015 | 1.5× |
| Easyjson | 95068 | 208.16 MB/s | 0 | 0 | 1.2× |
| Stdlib | 111604 | 177.31 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1730 | 10476.95 MB/s | 0 | 0 | 59.3× |
| LightningArena | 1853 | 9779.89 MB/s | 432 | 2 | 55.4× |
| Lightning | 1857 | 9758.98 MB/s | 432 | 2 | 55.3× |
| Easyjson | 3921 | 4622.71 MB/s | 432 | 2 | 26.2× |
| Sonic | 10036 | 1805.83 MB/s | 22917 | 6 | 10.2× |
| SonicFastest | 10260 | 1766.52 MB/s | 23586 | 6 | 10.0× |
| LightningDecodeAny | 15945 | 1121.46 MB/s | 29088 | 191 | 6.4× |
| Goccy | 16077 | 1127.33 MB/s | 19459 | 2 | 6.4× |
| JSONV2 | 45146 | 401.46 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102649 | 176.56 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2192544 | 916.06 MB/s | 3089566 | 6821 | 8.5× |
| LightningArena | 2269890 | 884.84 MB/s | 3094371 | 6703 | 8.2× |
| Lightning | 2301292 | 872.77 MB/s | 3091279 | 6827 | 8.1× |
| Goccy | 4348045 | 461.93 MB/s | 5413055 | 15837 | 4.3× |
| Sonic | 4579736 | 438.56 MB/s | 10922587 | 13683 | 4.1× |
| SonicFastest | 4594415 | 437.16 MB/s | 10934041 | 13683 | 4.1× |
| Easyjson | 4977378 | 403.52 MB/s | 2981514 | 7439 | 3.8× |
| LightningDecodeAny | 6944343 | 164.49 MB/s | 8503513 | 134008 | 2.7× |
| JSONV2 | 7032977 | 285.58 MB/s | 3173697 | 14563 | 2.7× |
| Stdlib | 18719905 | 107.29 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 868 | 632.51 MB/s | 480 | 1 | 6.6× |
| LightningArena | 873 | 628.71 MB/s | 480 | 1 | 6.5× |
| LightningDestructive | 878 | 624.93 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1605 | 341.37 MB/s | 2021 | 46 | 3.5× |
| Easyjson | 2145 | 255.95 MB/s | 1616 | 5 | 2.7× |
| SonicFastest | 2681 | 204.78 MB/s | 1951 | 26 | 2.1× |
| Sonic | 2688 | 204.28 MB/s | 1959 | 26 | 2.1× |
| Goccy | 3021 | 181.75 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3298 | 166.48 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5696 | 96.38 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 421964 | 1496.61 MB/s | 402729 | 545 | 12.8× |
| LightningArena | 492907 | 1281.20 MB/s | 453017 | 712 | 10.9× |
| Lightning | 493510 | 1279.64 MB/s | 451257 | 857 | 10.9× |
| SonicFastest | 1023734 | 616.87 MB/s | 1009053 | 1102 | 5.3× |
| Sonic | 1031970 | 611.95 MB/s | 1017657 | 1102 | 5.2× |
| Easyjson | 1143435 | 552.30 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1157475 | 545.60 MB/s | 986026 | 1201 | 4.7× |
| JSONV2 | 2140097 | 295.09 MB/s | 571613 | 3144 | 2.5× |
| LightningDecodeAny | 2336450 | 199.84 MB/s | 2076506 | 30126 | 2.3× |
| Stdlib | 5391525 | 117.13 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 566711 | 992.41 MB/s | 546572 | 429 | 9.3× |
| Lightning | 731260 | 769.09 MB/s | 769938 | 1235 | 7.2× |
| LightningArena | 736583 | 763.54 MB/s | 771666 | 1088 | 7.1× |
| SonicFastest | 1023564 | 549.46 MB/s | 938599 | 1476 | 5.1× |
| Sonic | 1023858 | 549.30 MB/s | 936292 | 1476 | 5.1× |
| Goccy | 1335204 | 421.21 MB/s | 1033779 | 1030 | 3.9× |
| Easyjson | 1750021 | 321.37 MB/s | 775153 | 1254 | 3.0× |
| LightningDecodeAny | 2556473 | 219.99 MB/s | 2180440 | 30126 | 2.1× |
| JSONV2 | 2740038 | 205.26 MB/s | 927445 | 3482 | 1.9× |
| Stdlib | 5264295 | 106.83 MB/s | 1011671 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 579303 | 920.38 MB/s | 333416 | 2084 | 9.5× |
| Lightning | 621826 | 857.44 MB/s | 368224 | 2293 | 8.8× |
| LightningArena | 625248 | 852.75 MB/s | 368224 | 2293 | 8.8× |
| Easyjson | 1124734 | 474.05 MB/s | 428362 | 3273 | 4.9× |
| SonicFastest | 1143719 | 466.18 MB/s | 1021871 | 4351 | 4.8× |
| Sonic | 1148792 | 464.12 MB/s | 1025019 | 4351 | 4.8× |
| Goccy | 1322050 | 403.30 MB/s | 1167228 | 5409 | 4.1× |
| JSONV2 | 2570861 | 207.39 MB/s | 745459 | 13288 | 2.1× |
| LightningDecodeAny | 3346756 | 159.31 MB/s | 2992878 | 50076 | 1.6× |
| Stdlib | 5477078 | 97.35 MB/s | 798693 | 17133 | 1.0× |
