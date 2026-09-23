# JSON Deserialization Benchmarks

- generated 2026-09-23T13:35:01Z
- go version go1.26.8 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 81642 | 1558.94 MB/s | 49821 | 2 | 13.4× |
| LightningArena | 81734 | 1557.18 MB/s | 49822 | 2 | 13.4× |
| LightningDestructive | 82150 | 1549.30 MB/s | 49280 | 2 | 13.3× |
| Sonic | 185032 | 687.85 MB/s | 197565 | 10 | 5.9× |
| SonicFastest | 188530 | 675.09 MB/s | 204498 | 10 | 5.8× |
| Goccy | 194630 | 653.93 MB/s | 224841 | 884 | 5.6× |
| Easyjson | 212689 | 598.41 MB/s | 122864 | 14 | 5.1× |
| JSONV2 | 419391 | 303.48 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 430867 | 219.68 MB/s | 464164 | 9706 | 2.5× |
| Stdlib | 1091338 | 116.62 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1587242 | 1418.22 MB/s | 2532848 | 1143 | 16.8× |
| Lightning | 1638995 | 1373.43 MB/s | 2532850 | 1143 | 16.2× |
| LightningArena | 2255720 | 997.93 MB/s | 2532848 | 1143 | 11.8× |
| SonicFastest | 4596446 | 489.74 MB/s | 15237106 | 970 | 5.8× |
| Sonic | 4647108 | 484.40 MB/s | 15248948 | 970 | 5.7× |
| LightningDecodeAny | 8423757 | 267.22 MB/s | 6828999 | 223498 | 3.2× |
| Goccy | 10637977 | 211.61 MB/s | 4131124 | 56533 | 2.5× |
| Easyjson | 11202505 | 200.94 MB/s | 3099808 | 2120 | 2.4× |
| JSONV2 | 16245800 | 138.56 MB/s | 3123221 | 3083 | 1.6× |
| Stdlib | 26600940 | 84.62 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 259059 | 1043.79 MB/s | 397296 | 567 | 13.3× |
| LightningDestructive | 264166 | 1023.61 MB/s | 397296 | 567 | 13.0× |
| LightningArena | 347077 | 779.09 MB/s | 397296 | 567 | 9.9× |
| SonicFastest | 642136 | 421.10 MB/s | 479333 | 968 | 5.4× |
| Sonic | 642824 | 420.65 MB/s | 474818 | 968 | 5.4× |
| LightningDecodeAny | 1182644 | 228.64 MB/s | 845690 | 29656 | 2.9× |
| Goccy | 1417224 | 190.80 MB/s | 542572 | 8122 | 2.4× |
| Easyjson | 1433323 | 188.65 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2094931 | 129.07 MB/s | 348160 | 1628 | 1.6× |
| Stdlib | 3441496 | 78.57 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 933162 | 1850.92 MB/s | 765560 | 2798 | 14.2× |
| Lightning | 948706 | 1820.59 MB/s | 768353 | 2798 | 14.0× |
| LightningArena | 954427 | 1809.68 MB/s | 775583 | 2444 | 13.9× |
| SonicFastest | 2067685 | 835.33 MB/s | 2673212 | 4020 | 6.4× |
| Sonic | 2070473 | 834.21 MB/s | 2695273 | 4020 | 6.4× |
| Goccy | 2422891 | 712.87 MB/s | 2582708 | 14605 | 5.5× |
| LightningDecodeAny | 3824882 | 130.80 MB/s | 4493709 | 67881 | 3.5× |
| Easyjson | 4216864 | 409.59 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4288470 | 402.76 MB/s | 1011637 | 7594 | 3.1× |
| Stdlib | 13234500 | 130.51 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 828 | 2187.15 MB/s | 0 | 0 | 16.8× |
| Lightning | 836 | 2166.57 MB/s | 0 | 0 | 16.7× |
| LightningDestructive | 847 | 2139.48 MB/s | 0 | 0 | 16.5× |
| Easyjson | 2541 | 712.98 MB/s | 24 | 1 | 5.5× |
| Goccy | 2850 | 635.89 MB/s | 2608 | 4 | 4.9× |
| Sonic | 5972 | 303.42 MB/s | 3766 | 40 | 2.3× |
| SonicFastest | 6040 | 299.99 MB/s | 3809 | 40 | 2.3× |
| JSONV2 | 7776 | 233.03 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7856 | 230.51 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13932 | 130.06 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 846 | 2140.89 MB/s | 0 | 0 | 16.5× |
| LightningArena | 851 | 2128.16 MB/s | 0 | 0 | 16.4× |
| LightningDestructive | 873 | 2075.07 MB/s | 0 | 0 | 16.0× |
| Easyjson | 2555 | 709.21 MB/s | 24 | 1 | 5.5× |
| Goccy | 2850 | 635.80 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6008 | 301.59 MB/s | 3818 | 40 | 2.3× |
| Sonic | 6018 | 301.08 MB/s | 3789 | 40 | 2.3× |
| JSONV2 | 7813 | 231.93 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7859 | 230.45 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13955 | 129.85 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1019 | 1777.91 MB/s | 144 | 10 | 13.6× |
| LightningArena | 1026 | 1766.30 MB/s | 144 | 10 | 13.6× |
| LightningDestructive | 1082 | 1674.64 MB/s | 144 | 10 | 12.9× |
| Easyjson | 2771 | 653.93 MB/s | 144 | 10 | 5.0× |
| Goccy | 2942 | 615.83 MB/s | 2600 | 5 | 4.7× |
| Sonic | 6158 | 294.24 MB/s | 3802 | 42 | 2.3× |
| SonicFastest | 6162 | 294.07 MB/s | 3810 | 42 | 2.3× |
| LightningDecodeAny | 7889 | 229.55 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7996 | 226.61 MB/s | 632 | 7 | 1.7× |
| Stdlib | 13908 | 130.28 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 641 | 770.19 MB/s | 160 | 1 | 8.5× |
| Lightning | 644 | 767.35 MB/s | 160 | 1 | 8.5× |
| LightningDecodeAny | 1195 | 412.64 MB/s | 1040 | 25 | 4.6× |
| Sonic | 1219 | 405.36 MB/s | 976 | 6 | 4.5× |
| SonicFastest | 1222 | 404.30 MB/s | 990 | 6 | 4.5× |
| LightningArena | 1374 | 359.60 MB/s | 4120 | 2 | 4.0× |
| Easyjson | 2227 | 221.80 MB/s | 448 | 3 | 2.5× |
| Goccy | 2428 | 203.48 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3224 | 153.23 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5475 | 90.23 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 300 | 765.27 MB/s | 160 | 1 | 13.5× |
| Lightning | 313 | 735.53 MB/s | 160 | 1 | 13.0× |
| Sonic | 861 | 267.20 MB/s | 665 | 6 | 4.7× |
| SonicFastest | 866 | 265.72 MB/s | 669 | 6 | 4.7× |
| LightningDecodeAny | 1029 | 222.53 MB/s | 1040 | 25 | 3.9× |
| LightningArena | 1101 | 208.94 MB/s | 4120 | 2 | 3.7× |
| Easyjson | 1399 | 164.45 MB/s | 448 | 3 | 2.9× |
| Goccy | 1574 | 146.14 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2423 | 94.91 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4060 | 56.65 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 50681 | 1285.13 MB/s | 97220 | 98 | 10.8× |
| LightningArena | 52045 | 1251.45 MB/s | 103665 | 99 | 10.5× |
| Lightning | 52140 | 1249.17 MB/s | 103667 | 99 | 10.5× |
| SonicFastest | 101054 | 644.53 MB/s | 158303 | 75 | 5.4× |
| Sonic | 101272 | 643.14 MB/s | 158040 | 75 | 5.4× |
| Goccy | 148758 | 437.84 MB/s | 228590 | 134 | 3.7× |
| LightningDecodeAny | 175675 | 303.57 MB/s | 176567 | 3237 | 3.1× |
| JSONV2 | 231803 | 280.98 MB/s | 206652 | 607 | 2.4× |
| Stdlib | 548537 | 118.74 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2121436 | 914.70 MB/s | 2185296 | 1350 | 10.9× |
| LightningArena | 2181893 | 889.35 MB/s | 2185297 | 1350 | 10.6× |
| Lightning | 2192649 | 884.99 MB/s | 2185298 | 1350 | 10.5× |
| SonicFastest | 4757197 | 407.90 MB/s | 14606973 | 1407 | 4.9× |
| Goccy | 4789298 | 405.17 MB/s | 4065917 | 13510 | 4.8× |
| Sonic | 4975853 | 389.98 MB/s | 14608596 | 1407 | 4.6× |
| Easyjson | 7472427 | 259.68 MB/s | 3871267 | 15043 | 3.1× |
| LightningDecodeAny | 8673485 | 223.72 MB/s | 6627984 | 206416 | 2.7× |
| JSONV2 | 11353755 | 170.91 MB/s | 3237225 | 13947 | 2.0× |
| Stdlib | 23077543 | 84.08 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 872257 | 3815.20 MB/s | 351704 | 1286 | 24.1× |
| LightningArena | 1318774 | 2523.43 MB/s | 2434273 | 1413 | 15.9× |
| Lightning | 1323691 | 2514.05 MB/s | 2434226 | 1413 | 15.9× |
| Sonic | 2706451 | 1229.59 MB/s | 6416862 | 4248 | 7.8× |
| SonicFastest | 2715103 | 1225.67 MB/s | 6472172 | 4248 | 7.7× |
| LightningDecodeAny | 3265305 | 941.34 MB/s | 4825406 | 55311 | 6.4× |
| Goccy | 4493635 | 740.57 MB/s | 3948908 | 3816 | 4.7× |
| JSONV2 | 7533729 | 441.72 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20990158 | 158.54 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 175441 | 1255.95 MB/s | 135392 | 226 | 11.5× |
| Lightning | 176105 | 1251.22 MB/s | 135392 | 226 | 11.5× |
| LightningDestructive | 177190 | 1243.56 MB/s | 135392 | 226 | 11.4× |
| Sonic | 375720 | 586.46 MB/s | 295229 | 398 | 5.4× |
| SonicFastest | 376437 | 585.35 MB/s | 295282 | 398 | 5.4× |
| Goccy | 437709 | 503.41 MB/s | 364463 | 1067 | 4.6× |
| Easyjson | 549468 | 401.02 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 728395 | 302.51 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 848759 | 127.61 MB/s | 854737 | 11700 | 2.4× |
| Stdlib | 2019177 | 109.13 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 8531106 | 949.47 MB/s | 8109650 | 20809 | 10.3× |
| LightningDestructive | 8557584 | 946.53 MB/s | 8109648 | 20809 | 10.3× |
| LightningArena | 9692874 | 835.67 MB/s | 8109650 | 20809 | 9.1× |
| Sonic | 17018322 | 475.96 MB/s | 70887372 | 40014 | 5.2× |
| SonicFastest | 17032990 | 475.55 MB/s | 70887497 | 40014 | 5.2× |
| Goccy | 23743408 | 341.15 MB/s | 16913956 | 107148 | 3.7× |
| LightningDecodeAny | 30456044 | 170.84 MB/s | 28359846 | 746961 | 2.9× |
| Easyjson | 30967537 | 261.57 MB/s | 15059621 | 41643 | 2.8× |
| JSONV2 | 43731966 | 185.22 MB/s | 15233716 | 78972 | 2.0× |
| Stdlib | 87972515 | 92.07 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 3370850 | 885.08 MB/s | 3780459 | 1514 | 13.8× |
| LightningDestructive | 3531983 | 844.70 MB/s | 3758856 | 29356 | 13.2× |
| Lightning | 3676615 | 811.47 MB/s | 3758857 | 29356 | 12.7× |
| Sonic | 8694464 | 343.15 MB/s | 26611592 | 56760 | 5.4× |
| SonicFastest | 8702676 | 342.82 MB/s | 26670481 | 56760 | 5.4× |
| LightningDecodeAny | 14767505 | 124.20 MB/s | 18225376 | 350883 | 3.2× |
| Goccy | 16612139 | 179.60 MB/s | 10613831 | 273649 | 2.8× |
| Easyjson | 16751589 | 178.10 MB/s | 9479440 | 30115 | 2.8× |
| JSONV2 | 25111424 | 118.81 MB/s | 9257151 | 86278 | 1.9× |
| Stdlib | 46667123 | 63.93 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 690953 | 1047.25 MB/s | 907601 | 3618 | 16.5× |
| LightningArena | 712748 | 1015.22 MB/s | 916260 | 37 | 16.0× |
| Lightning | 738130 | 980.31 MB/s | 907598 | 3618 | 15.4× |
| Sonic | 1778901 | 406.77 MB/s | 3187902 | 7226 | 6.4× |
| SonicFastest | 1781370 | 406.20 MB/s | 3199577 | 7226 | 6.4× |
| LightningDecodeAny | 3952471 | 164.60 MB/s | 5691595 | 76540 | 2.9× |
| Easyjson | 4170303 | 173.51 MB/s | 2847904 | 3698 | 2.7× |
| Goccy | 4782040 | 151.32 MB/s | 2790949 | 80272 | 2.4× |
| JSONV2 | 5575099 | 129.79 MB/s | 2704638 | 7318 | 2.0× |
| Stdlib | 11378660 | 63.59 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1010506 | 1560.95 MB/s | 907600 | 3618 | 15.4× |
| LightningArena | 1033201 | 1526.67 MB/s | 916257 | 37 | 15.0× |
| Lightning | 1061830 | 1485.50 MB/s | 907594 | 3618 | 14.6× |
| Sonic | 2241186 | 703.80 MB/s | 5787357 | 7226 | 6.9× |
| SonicFastest | 2255124 | 699.45 MB/s | 5792258 | 7226 | 6.9× |
| LightningDecodeAny | 3453421 | 218.16 MB/s | 5691593 | 76540 | 4.5× |
| Easyjson | 5563780 | 283.50 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5681245 | 277.64 MB/s | 3624676 | 80270 | 2.7× |
| JSONV2 | 6425353 | 245.49 MB/s | 2704594 | 7318 | 2.4× |
| Stdlib | 15532562 | 101.55 MB/s | 2704552 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 70691 | 2123.65 MB/s | 81920 | 1 | 26.3× |
| LightningArena | 72402 | 2073.48 MB/s | 81920 | 1 | 25.7× |
| Lightning | 72461 | 2071.80 MB/s | 81920 | 1 | 25.6× |
| Sonic | 273348 | 549.20 MB/s | 253127 | 6 | 6.8× |
| SonicFastest | 274004 | 547.89 MB/s | 253988 | 6 | 6.8× |
| LightningDecodeAny | 432497 | 347.10 MB/s | 745508 | 10015 | 4.3× |
| Goccy | 871303 | 172.30 MB/s | 324201 | 10004 | 2.1× |
| JSONV2 | 1074206 | 139.75 MB/s | 357717 | 20 | 1.7× |
| Stdlib | 1857582 | 80.82 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 26954 | 1043.13 MB/s | 29088 | 101 | 11.2× |
| LightningArena | 27217 | 1033.06 MB/s | 29235 | 101 | 11.1× |
| Lightning | 27236 | 1032.36 MB/s | 29235 | 101 | 11.0× |
| Sonic | 63479 | 442.94 MB/s | 46680 | 103 | 4.7× |
| SonicFastest | 64502 | 435.91 MB/s | 47697 | 103 | 4.7× |
| Easyjson | 68171 | 412.45 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72414 | 388.28 MB/s | 59210 | 188 | 4.2× |
| JSONV2 | 134814 | 208.56 MB/s | 36895 | 242 | 2.2× |
| LightningDecodeAny | 144896 | 194.05 MB/s | 133683 | 2639 | 2.1× |
| Stdlib | 300806 | 93.47 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1501 | 1550.85 MB/s | 32 | 1 | 15.1× |
| Lightning | 1503 | 1548.82 MB/s | 32 | 1 | 15.0× |
| LightningDestructive | 1581 | 1472.40 MB/s | 32 | 1 | 14.3× |
| Goccy | 4141 | 562.21 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4205 | 553.65 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5114 | 455.18 MB/s | 4291 | 6 | 4.4× |
| Sonic | 5115 | 455.15 MB/s | 4316 | 6 | 4.4× |
| JSONV2 | 8442 | 275.77 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9646 | 174.68 MB/s | 9936 | 194 | 2.3× |
| Stdlib | 22604 | 102.99 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 179 | 1055.07 MB/s | 0 | 0 | 13.3× |
| Lightning | 180 | 1051.68 MB/s | 0 | 0 | 13.2× |
| LightningDestructive | 182 | 1038.97 MB/s | 0 | 0 | 13.1× |
| Goccy | 377 | 501.52 MB/s | 304 | 2 | 6.3× |
| Easyjson | 482 | 391.74 MB/s | 0 | 0 | 4.9× |
| Sonic | 781 | 242.12 MB/s | 500 | 4 | 3.0× |
| SonicFastest | 791 | 238.96 MB/s | 503 | 4 | 3.0× |
| JSONV2 | 1029 | 183.75 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1195 | 112.18 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2376 | 79.53 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1092 | 2006.07 MB/s | 0 | 0 | 14.5× |
| LightningArena | 1092 | 2006.76 MB/s | 0 | 0 | 14.5× |
| LightningDestructive | 1120 | 1955.85 MB/s | 0 | 0 | 14.1× |
| Goccy | 3175 | 690.16 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3186 | 687.78 MB/s | 24 | 1 | 5.0× |
| Sonic | 6392 | 342.77 MB/s | 3986 | 40 | 2.5× |
| SonicFastest | 6395 | 342.61 MB/s | 4024 | 40 | 2.5× |
| LightningDecodeAny | 7808 | 231.94 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 8055 | 272.00 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15821 | 138.49 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 541059 | 943.48 MB/s | 318400 | 1005 | 11.0× |
| LightningArena | 542350 | 941.23 MB/s | 318400 | 1005 | 11.0× |
| Lightning | 543275 | 939.63 MB/s | 318400 | 1005 | 11.0× |
| SonicFastest | 1144434 | 446.05 MB/s | 853203 | 2006 | 5.2× |
| Goccy | 1150771 | 443.59 MB/s | 1139034 | 5006 | 5.2× |
| Sonic | 1151417 | 443.35 MB/s | 859079 | 2006 | 5.2× |
| Easyjson | 1534422 | 332.68 MB/s | 863778 | 3012 | 3.9× |
| LightningDecodeAny | 3230425 | 142.85 MB/s | 2742392 | 64017 | 1.8× |
| JSONV2 | 3275603 | 155.84 MB/s | 1076005 | 12646 | 1.8× |
| Stdlib | 5968807 | 85.52 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 461 | 42952.62 MB/s | 0 | 0 | 235.7× |
| LightningArena | 461 | 42890.98 MB/s | 0 | 0 | 235.4× |
| LightningDestructive | 492 | 40214.45 MB/s | 0 | 0 | 220.7× |
| Goccy | 19836 | 997.63 MB/s | 20491 | 2 | 5.5× |
| Sonic | 27206 | 727.37 MB/s | 22286 | 4 | 4.0× |
| SonicFastest | 27246 | 726.31 MB/s | 22442 | 4 | 4.0× |
| JSONV2 | 29745 | 665.30 MB/s | 8 | 1 | 3.7× |
| Easyjson | 81821 | 241.86 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 86032 | 230.01 MB/s | 116608 | 2014 | 1.3× |
| Stdlib | 108599 | 182.22 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1748 | 10371.15 MB/s | 0 | 0 | 58.6× |
| LightningArena | 1817 | 9972.53 MB/s | 405 | 0 | 56.4× |
| Lightning | 1821 | 9953.47 MB/s | 405 | 0 | 56.2× |
| Easyjson | 3958 | 4579.06 MB/s | 432 | 2 | 25.9× |
| Sonic | 10061 | 1801.44 MB/s | 23175 | 6 | 10.2× |
| SonicFastest | 10069 | 1800.06 MB/s | 23272 | 6 | 10.2× |
| Goccy | 15589 | 1162.64 MB/s | 19459 | 2 | 6.6× |
| LightningDecodeAny | 15964 | 1120.11 MB/s | 29103 | 189 | 6.4× |
| JSONV2 | 46777 | 387.45 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102415 | 176.97 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2017157 | 995.71 MB/s | 3089565 | 6821 | 9.2× |
| Lightning | 2078689 | 966.23 MB/s | 3096779 | 6823 | 8.9× |
| LightningArena | 2088595 | 961.65 MB/s | 3100277 | 6700 | 8.9× |
| Goccy | 4308313 | 466.19 MB/s | 5411930 | 15831 | 4.3× |
| Sonic | 4462304 | 450.10 MB/s | 10942019 | 13683 | 4.2× |
| SonicFastest | 4464105 | 449.92 MB/s | 10900729 | 13683 | 4.2× |
| Easyjson | 4936541 | 406.86 MB/s | 2981485 | 7439 | 3.8× |
| LightningDecodeAny | 6906537 | 165.39 MB/s | 7376728 | 134004 | 2.7× |
| JSONV2 | 7080340 | 283.67 MB/s | 3173682 | 14563 | 2.6× |
| Stdlib | 18541665 | 108.32 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 831 | 661.00 MB/s | 480 | 1 | 6.7× |
| LightningArena | 838 | 655.21 MB/s | 480 | 1 | 6.7× |
| LightningDestructive | 847 | 648.38 MB/s | 480 | 1 | 6.6× |
| LightningDecodeAny | 1512 | 362.46 MB/s | 1765 | 45 | 3.7× |
| Easyjson | 2144 | 256.10 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2686 | 204.41 MB/s | 1943 | 26 | 2.1× |
| SonicFastest | 2697 | 203.53 MB/s | 1949 | 26 | 2.1× |
| Goccy | 3047 | 180.19 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3324 | 165.18 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5589 | 98.23 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 407658 | 1549.13 MB/s | 400489 | 545 | 13.2× |
| Lightning | 439966 | 1435.37 MB/s | 447069 | 548 | 12.2× |
| LightningArena | 444861 | 1419.58 MB/s | 449121 | 404 | 12.1× |
| Sonic | 1019028 | 619.72 MB/s | 1004436 | 1102 | 5.3× |
| SonicFastest | 1020270 | 618.97 MB/s | 1010732 | 1102 | 5.3× |
| Easyjson | 1146124 | 551.00 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1161250 | 543.82 MB/s | 986020 | 1201 | 4.6× |
| JSONV2 | 2156178 | 292.89 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2262469 | 206.37 MB/s | 1992797 | 29072 | 2.4× |
| Stdlib | 5363993 | 117.73 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 574747 | 978.53 MB/s | 391185 | 426 | 9.1× |
| Lightning | 638224 | 881.21 MB/s | 506391 | 433 | 8.2× |
| LightningArena | 641161 | 877.17 MB/s | 508274 | 287 | 8.2× |
| SonicFastest | 1032714 | 544.59 MB/s | 965461 | 1476 | 5.1× |
| Sonic | 1033806 | 544.02 MB/s | 971757 | 1476 | 5.1× |
| Goccy | 1322830 | 425.16 MB/s | 1035855 | 1029 | 4.0× |
| Easyjson | 1746155 | 322.08 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2305696 | 243.92 MB/s | 1991359 | 28581 | 2.3× |
| JSONV2 | 2772923 | 202.82 MB/s | 927442 | 3482 | 1.9× |
| Stdlib | 5238948 | 107.35 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 550574 | 968.40 MB/s | 333416 | 2084 | 9.8× |
| Lightning | 560315 | 951.57 MB/s | 367499 | 2086 | 9.6× |
| LightningArena | 563080 | 946.90 MB/s | 367555 | 2086 | 9.6× |
| Easyjson | 1102187 | 483.75 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1146791 | 464.93 MB/s | 1034882 | 4351 | 4.7× |
| Sonic | 1149262 | 463.93 MB/s | 1030684 | 4351 | 4.7× |
| Goccy | 1323098 | 402.98 MB/s | 1167229 | 5409 | 4.1× |
| JSONV2 | 2518739 | 211.68 MB/s | 745449 | 13288 | 2.1× |
| LightningDecodeAny | 3218378 | 165.67 MB/s | 2659122 | 49349 | 1.7× |
| Stdlib | 5398233 | 98.77 MB/s | 798692 | 17133 | 1.0× |
