# JSON Deserialization Benchmarks

- generated 2026-06-23T19:53:27Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104624 | 1216.49 MB/s | 49760 | 3 | 12.7× |
| LightningDestructive | 106608 | 1193.86 MB/s | 49280 | 2 | 12.5× |
| SonicFastest | 203715 | 624.77 MB/s | 214346 | 15 | 6.5× |
| Sonic | 205676 | 618.81 MB/s | 214450 | 15 | 6.5× |
| Goccy | 247511 | 514.22 MB/s | 225085 | 884 | 5.4× |
| Easyjson | 255056 | 499.01 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 436316 | 291.70 MB/s | 195127 | 1805 | 3.0× |
| LightningDecodeAny | 493456 | 191.82 MB/s | 465586 | 9714 | 2.7× |
| Stdlib | 1330186 | 95.68 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4140456 | 543.67 MB/s | 3122872 | 3081 | 8.2× |
| Lightning | 4184248 | 537.98 MB/s | 3122874 | 3081 | 8.1× |
| Sonic | 5338281 | 421.68 MB/s | 4868507 | 2584 | 6.4× |
| SonicFastest | 5395974 | 417.17 MB/s | 4869276 | 2584 | 6.3× |
| Goccy | 12125623 | 185.64 MB/s | 4221271 | 56536 | 2.8× |
| LightningDecodeAny | 12291848 | 183.13 MB/s | 7938298 | 281383 | 2.8× |
| Easyjson | 13920593 | 161.71 MB/s | 3099809 | 2120 | 2.4× |
| JSONV2 | 17697181 | 127.20 MB/s | 3123198 | 3083 | 1.9× |
| Stdlib | 34049441 | 66.11 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 572110 | 472.64 MB/s | 348024 | 1627 | 7.7× |
| LightningDestructive | 584942 | 462.27 MB/s | 348025 | 1627 | 7.5× |
| SonicFastest | 759608 | 355.98 MB/s | 642290 | 1147 | 5.8× |
| Sonic | 759656 | 355.95 MB/s | 642028 | 1147 | 5.8× |
| Goccy | 1686455 | 160.34 MB/s | 546227 | 8123 | 2.6× |
| LightningDecodeAny | 1724781 | 156.78 MB/s | 1011488 | 37901 | 2.5× |
| Easyjson | 1813631 | 149.09 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2378329 | 113.69 MB/s | 348159 | 1628 | 1.8× |
| Stdlib | 4390526 | 61.59 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1297467 | 1331.21 MB/s | 959848 | 5881 | 13.7× |
| Lightning | 1360297 | 1269.73 MB/s | 959890 | 5882 | 13.1× |
| Sonic | 2132443 | 809.96 MB/s | 2694569 | 5547 | 8.4× |
| SonicFastest | 2150308 | 803.24 MB/s | 2694180 | 5547 | 8.3× |
| Goccy | 2492096 | 693.07 MB/s | 2581258 | 14603 | 7.1× |
| Easyjson | 4046721 | 426.82 MB/s | 972032 | 5389 | 4.4× |
| LightningDecodeAny | 4443114 | 112.60 MB/s | 4976203 | 81466 | 4.0× |
| JSONV2 | 4733850 | 364.86 MB/s | 1011612 | 7594 | 3.8× |
| Stdlib | 17811450 | 96.97 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1024 | 1769.82 MB/s | 0 | 0 | 15.7× |
| LightningDestructive | 1090 | 1661.88 MB/s | 0 | 0 | 14.8× |
| Easyjson | 3020 | 600.01 MB/s | 24 | 1 | 5.3× |
| Goccy | 3247 | 558.08 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6355 | 285.14 MB/s | 3344 | 38 | 2.5× |
| Sonic | 6551 | 276.59 MB/s | 3345 | 38 | 2.5× |
| JSONV2 | 8113 | 223.35 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9434 | 191.97 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16083 | 112.67 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1116 | 1623.01 MB/s | 0 | 0 | 14.7× |
| LightningDestructive | 1201 | 1508.37 MB/s | 0 | 0 | 13.6× |
| Easyjson | 2968 | 610.51 MB/s | 24 | 1 | 5.5× |
| Goccy | 3170 | 571.56 MB/s | 2608 | 4 | 5.2× |
| SonicFastest | 5989 | 302.58 MB/s | 3342 | 38 | 2.7× |
| Sonic | 6235 | 290.60 MB/s | 3342 | 38 | 2.6× |
| JSONV2 | 8046 | 225.22 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9676 | 187.16 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16367 | 110.71 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1316 | 1377.41 MB/s | 144 | 10 | 12.4× |
| LightningDestructive | 1402 | 1292.63 MB/s | 144 | 10 | 11.6× |
| Easyjson | 3106 | 583.43 MB/s | 144 | 10 | 5.2× |
| Goccy | 3492 | 518.91 MB/s | 2600 | 5 | 4.7× |
| SonicFastest | 6413 | 282.55 MB/s | 3365 | 40 | 2.5× |
| Sonic | 6659 | 272.10 MB/s | 3368 | 40 | 2.4× |
| JSONV2 | 8706 | 208.13 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9292 | 194.89 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 16278 | 111.32 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 724 | 682.73 MB/s | 160 | 1 | 9.6× |
| LightningDestructive | 741 | 666.33 MB/s | 160 | 1 | 9.4× |
| Sonic | 1295 | 381.33 MB/s | 1075 | 8 | 5.4× |
| SonicFastest | 1319 | 374.39 MB/s | 1075 | 8 | 5.3× |
| LightningDecodeAny | 1798 | 274.15 MB/s | 1536 | 30 | 3.9× |
| Goccy | 2611 | 189.20 MB/s | 856 | 23 | 2.7× |
| Easyjson | 2655 | 186.10 MB/s | 448 | 3 | 2.6× |
| JSONV2 | 3519 | 140.36 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6958 | 71.00 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 496 | 463.33 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 499 | 460.78 MB/s | 160 | 1 | 9.9× |
| Sonic | 928 | 247.88 MB/s | 801 | 8 | 5.3× |
| SonicFastest | 933 | 246.52 MB/s | 801 | 8 | 5.3× |
| LightningDecodeAny | 1613 | 141.95 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1732 | 132.82 MB/s | 448 | 3 | 2.8× |
| Goccy | 1850 | 124.31 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2553 | 90.10 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4929 | 46.67 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 90753 | 717.68 MB/s | 172434 | 107 | 7.5× |
| LightningDestructive | 96939 | 671.89 MB/s | 166213 | 102 | 7.1× |
| SonicFastest | 149963 | 434.32 MB/s | 235983 | 65 | 4.6× |
| Sonic | 150559 | 432.60 MB/s | 236028 | 65 | 4.5× |
| Goccy | 176968 | 368.04 MB/s | 228124 | 134 | 3.9× |
| LightningDecodeAny | 204471 | 260.81 MB/s | 176961 | 3252 | 3.3× |
| JSONV2 | 266871 | 244.06 MB/s | 206663 | 607 | 2.6× |
| Stdlib | 684522 | 95.15 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2541748 | 763.44 MB/s | 2846864 | 2238 | 11.2× |
| Lightning | 2620249 | 740.57 MB/s | 2846867 | 2238 | 10.8× |
| Sonic | 4833506 | 401.46 MB/s | 4879270 | 1736 | 5.9× |
| SonicFastest | 4858072 | 399.43 MB/s | 4880057 | 1736 | 5.8× |
| Goccy | 4944464 | 392.45 MB/s | 4061398 | 13509 | 5.7× |
| Easyjson | 8614246 | 225.26 MB/s | 3871264 | 15043 | 3.3× |
| LightningDecodeAny | 10092818 | 192.26 MB/s | 7013524 | 219937 | 2.8× |
| JSONV2 | 12817150 | 151.40 MB/s | 3237179 | 13947 | 2.2× |
| Stdlib | 28369481 | 68.40 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1064558 | 3126.02 MB/s | 351704 | 1286 | 22.2× |
| Lightning | 1636908 | 2033.00 MB/s | 2488907 | 2995 | 14.4× |
| SonicFastest | 2004380 | 1660.28 MB/s | 5896364 | 4263 | 11.8× |
| Sonic | 2008811 | 1656.62 MB/s | 5896393 | 4263 | 11.7× |
| LightningDecodeAny | 3534563 | 869.63 MB/s | 4886622 | 56892 | 6.7× |
| Goccy | 6069308 | 548.30 MB/s | 3948914 | 3817 | 3.9× |
| JSONV2 | 7594689 | 438.18 MB/s | 5364501 | 13243 | 3.1× |
| Stdlib | 23583587 | 141.11 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225776 | 975.95 MB/s | 136896 | 228 | 10.8× |
| LightningDestructive | 235568 | 935.38 MB/s | 136896 | 228 | 10.4× |
| Sonic | 484011 | 455.25 MB/s | 351094 | 262 | 5.0× |
| SonicFastest | 486387 | 453.03 MB/s | 351119 | 262 | 5.0× |
| Goccy | 489715 | 449.95 MB/s | 364158 | 1066 | 5.0× |
| Easyjson | 651117 | 338.41 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 693361 | 317.79 MB/s | 129747 | 470 | 3.5× |
| LightningDecodeAny | 1027114 | 105.45 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2442036 | 90.23 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 15191291 | 533.20 MB/s | 15059834 | 51649 | 7.4× |
| Lightning | 15500300 | 522.57 MB/s | 15059845 | 51649 | 7.2× |
| SonicFastest | 17783519 | 455.48 MB/s | 19857313 | 41640 | 6.3× |
| Sonic | 17816603 | 454.63 MB/s | 19861295 | 41640 | 6.3× |
| Goccy | 25395781 | 318.95 MB/s | 19259438 | 107156 | 4.4× |
| Easyjson | 36083827 | 224.48 MB/s | 15059618 | 41643 | 3.1× |
| LightningDecodeAny | 43648540 | 119.20 MB/s | 34333347 | 912810 | 2.6× |
| JSONV2 | 50147254 | 161.53 MB/s | 15233723 | 78972 | 2.2× |
| Stdlib | 112277384 | 72.14 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6397975 | 466.31 MB/s | 3985336 | 30015 | 9.1× |
| Lightning | 6599157 | 452.10 MB/s | 3985337 | 30015 | 8.8× |
| Sonic | 9655427 | 308.99 MB/s | 9132357 | 57804 | 6.0× |
| SonicFastest | 9664232 | 308.71 MB/s | 9131929 | 57804 | 6.0× |
| Goccy | 18724777 | 159.33 MB/s | 9830111 | 273617 | 3.1× |
| Easyjson | 19244873 | 155.03 MB/s | 9479441 | 30115 | 3.0× |
| LightningDecodeAny | 21461470 | 85.46 MB/s | 20023836 | 408557 | 2.7× |
| JSONV2 | 29426085 | 101.39 MB/s | 9257038 | 86278 | 2.0× |
| Stdlib | 57941957 | 51.49 MB/s | 9258086 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1479050 | 489.23 MB/s | 907393 | 3618 | 9.6× |
| Lightning | 1560200 | 463.78 MB/s | 907387 | 3618 | 9.1× |
| SonicFastest | 2119210 | 341.45 MB/s | 2370677 | 3683 | 6.7× |
| Sonic | 2145191 | 337.31 MB/s | 2371725 | 3683 | 6.7× |
| Easyjson | 5422842 | 133.44 MB/s | 2847907 | 3698 | 2.6× |
| Goccy | 5546111 | 130.47 MB/s | 2723144 | 80268 | 2.6× |
| LightningDecodeAny | 5665459 | 114.83 MB/s | 5752203 | 80172 | 2.5× |
| JSONV2 | 6360021 | 113.77 MB/s | 2704707 | 7318 | 2.2× |
| Stdlib | 14270807 | 50.70 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2047557 | 770.36 MB/s | 907387 | 3618 | 9.7× |
| LightningDestructive | 2050239 | 769.35 MB/s | 907392 | 3618 | 9.7× |
| Sonic | 2429287 | 649.31 MB/s | 3224741 | 3683 | 8.2× |
| SonicFastest | 2432604 | 648.42 MB/s | 3223961 | 3683 | 8.1× |
| LightningDecodeAny | 4792955 | 157.19 MB/s | 5752199 | 80172 | 4.1× |
| Easyjson | 6397246 | 246.57 MB/s | 2847905 | 3698 | 3.1× |
| Goccy | 6685357 | 235.94 MB/s | 3499523 | 80262 | 3.0× |
| JSONV2 | 7043308 | 223.95 MB/s | 2704553 | 7318 | 2.8× |
| Stdlib | 19819962 | 79.58 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 252930 | 593.54 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 262767 | 571.32 MB/s | 81920 | 1 | 8.3× |
| SonicFastest | 385422 | 389.51 MB/s | 407616 | 16 | 5.6× |
| Sonic | 388930 | 385.99 MB/s | 407547 | 16 | 5.6× |
| LightningDecodeAny | 605968 | 247.74 MB/s | 746006 | 10020 | 3.6× |
| Goccy | 1028687 | 145.94 MB/s | 324125 | 10005 | 2.1× |
| JSONV2 | 1098870 | 136.62 MB/s | 357724 | 20 | 2.0× |
| Stdlib | 2170256 | 69.17 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 34529 | 814.31 MB/s | 30272 | 105 | 10.1× |
| LightningDestructive | 35961 | 781.87 MB/s | 30144 | 103 | 9.7× |
| Sonic | 60383 | 465.64 MB/s | 59475 | 83 | 5.8× |
| SonicFastest | 60486 | 464.85 MB/s | 59496 | 83 | 5.8× |
| Goccy | 80486 | 349.34 MB/s | 59260 | 188 | 4.3× |
| Easyjson | 82669 | 340.11 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 135641 | 207.29 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 169178 | 166.20 MB/s | 135024 | 2678 | 2.1× |
| Stdlib | 348638 | 80.65 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1964 | 1185.40 MB/s | 32 | 1 | 13.5× |
| LightningDestructive | 2128 | 1093.93 MB/s | 32 | 1 | 12.5× |
| SonicFastest | 4857 | 479.33 MB/s | 3712 | 4 | 5.5× |
| Sonic | 4889 | 476.13 MB/s | 3713 | 4 | 5.4× |
| Goccy | 5064 | 459.69 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 5189 | 448.62 MB/s | 192 | 2 | 5.1× |
| JSONV2 | 8415 | 276.66 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 10807 | 155.92 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 26559 | 87.65 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 216 | 875.50 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 224 | 844.77 MB/s | 0 | 0 | 12.5× |
| Goccy | 431 | 438.35 MB/s | 304 | 2 | 6.5× |
| Easyjson | 598 | 316.01 MB/s | 0 | 0 | 4.7× |
| SonicFastest | 638 | 296.11 MB/s | 341 | 3 | 4.4× |
| Sonic | 640 | 295.47 MB/s | 341 | 3 | 4.4× |
| JSONV2 | 1053 | 179.48 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1356 | 98.82 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2804 | 67.39 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1491 | 1469.97 MB/s | 0 | 0 | 12.7× |
| LightningDestructive | 1556 | 1408.04 MB/s | 0 | 0 | 12.2× |
| Goccy | 3615 | 606.07 MB/s | 2864 | 4 | 5.2× |
| Easyjson | 3793 | 577.69 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6567 | 333.65 MB/s | 3601 | 38 | 2.9× |
| Sonic | 6727 | 325.71 MB/s | 3600 | 38 | 2.8× |
| JSONV2 | 8203 | 267.11 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 9352 | 193.66 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18922 | 115.79 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 704543 | 724.55 MB/s | 703779 | 1012 | 10.1× |
| Lightning | 750394 | 680.28 MB/s | 703779 | 1012 | 9.5× |
| SonicFastest | 1265268 | 403.45 MB/s | 1309445 | 2014 | 5.6× |
| Sonic | 1274048 | 400.67 MB/s | 1310377 | 2014 | 5.6× |
| Goccy | 1342342 | 380.29 MB/s | 1138894 | 5006 | 5.3× |
| Easyjson | 1687823 | 302.45 MB/s | 863778 | 3012 | 4.2× |
| JSONV2 | 3531908 | 144.53 MB/s | 1075964 | 12645 | 2.0× |
| LightningDecodeAny | 3601556 | 128.13 MB/s | 2785927 | 66022 | 2.0× |
| Stdlib | 7099634 | 71.90 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 645 | 30670.15 MB/s | 0 | 0 | 252.7× |
| LightningDestructive | 949 | 20847.62 MB/s | 0 | 0 | 171.8× |
| SonicFastest | 6823 | 2900.17 MB/s | 21133 | 3 | 23.9× |
| Goccy | 29642 | 667.60 MB/s | 20492 | 2 | 5.5× |
| Sonic | 32068 | 617.10 MB/s | 20612 | 3 | 5.1× |
| JSONV2 | 33295 | 594.36 MB/s | 8 | 1 | 4.9× |
| Easyjson | 102464 | 193.13 MB/s | 0 | 0 | 1.6× |
| LightningDecodeAny | 104445 | 189.46 MB/s | 117104 | 2019 | 1.6× |
| Stdlib | 163051 | 121.37 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2263 | 8008.66 MB/s | 432 | 2 | 52.9× |
| LightningDestructive | 2569 | 7055.98 MB/s | 0 | 0 | 46.6× |
| Easyjson | 5139 | 3526.93 MB/s | 432 | 2 | 23.3× |
| Sonic | 9984 | 1815.27 MB/s | 20358 | 5 | 12.0× |
| SonicFastest | 10036 | 1805.86 MB/s | 20366 | 5 | 11.9× |
| LightningDecodeAny | 19803 | 903.00 MB/s | 29088 | 191 | 6.0× |
| Goccy | 21268 | 852.17 MB/s | 19460 | 2 | 5.6× |
| JSONV2 | 53193 | 340.72 MB/s | 16501 | 50 | 2.3× |
| Stdlib | 119786 | 151.30 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2316658 | 866.98 MB/s | 2960389 | 7411 | 9.6× |
| Lightning | 2403595 | 835.62 MB/s | 2962101 | 7417 | 9.2× |
| SonicFastest | 4366078 | 460.02 MB/s | 5156278 | 7085 | 5.1× |
| Sonic | 4395785 | 456.91 MB/s | 5156900 | 7085 | 5.0× |
| Goccy | 4771598 | 420.93 MB/s | 5410210 | 15832 | 4.6× |
| Easyjson | 5670867 | 354.18 MB/s | 2981487 | 7439 | 3.9× |
| LightningDecodeAny | 7336711 | 155.70 MB/s | 7386073 | 134751 | 3.0× |
| JSONV2 | 7994020 | 251.25 MB/s | 3173677 | 14563 | 2.8× |
| Stdlib | 22135186 | 90.74 MB/s | 3589328 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 846 | 648.64 MB/s | 480 | 1 | 8.1× |
| LightningDestructive | 862 | 636.68 MB/s | 480 | 1 | 8.0× |
| LightningDecodeAny | 2185 | 250.81 MB/s | 2261 | 50 | 3.2× |
| SonicFastest | 2387 | 230.03 MB/s | 2260 | 8 | 2.9× |
| Easyjson | 2405 | 228.28 MB/s | 1616 | 5 | 2.9× |
| Sonic | 2443 | 224.74 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3481 | 157.71 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3693 | 148.67 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6883 | 79.76 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 506464 | 1246.91 MB/s | 364472 | 566 | 12.6× |
| Lightning | 572521 | 1103.04 MB/s | 413001 | 878 | 11.1× |
| Sonic | 1003298 | 629.44 MB/s | 1065015 | 814 | 6.3× |
| SonicFastest | 1004697 | 628.56 MB/s | 1065203 | 814 | 6.3× |
| Easyjson | 1316914 | 479.54 MB/s | 422504 | 936 | 4.8× |
| Goccy | 1419057 | 445.02 MB/s | 989361 | 1200 | 4.5× |
| JSONV2 | 2350305 | 268.69 MB/s | 571592 | 3144 | 2.7× |
| LightningDecodeAny | 2572742 | 181.48 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6366970 | 99.19 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 831725 | 676.19 MB/s | 544249 | 448 | 7.5× |
| Lightning | 1036112 | 542.81 MB/s | 767622 | 1254 | 6.1× |
| SonicFastest | 1353887 | 415.40 MB/s | 1350665 | 1185 | 4.6× |
| Sonic | 1366542 | 411.56 MB/s | 1350536 | 1185 | 4.6× |
| Goccy | 1659952 | 338.81 MB/s | 1043469 | 1029 | 3.8× |
| Easyjson | 2179144 | 258.09 MB/s | 775153 | 1254 | 2.9× |
| LightningDecodeAny | 3155401 | 178.24 MB/s | 2114150 | 30295 | 2.0× |
| JSONV2 | 3351750 | 167.80 MB/s | 927407 | 3482 | 1.9× |
| Stdlib | 6273690 | 89.65 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 625606 | 852.26 MB/s | 333416 | 2084 | 10.3× |
| Lightning | 704640 | 756.67 MB/s | 368224 | 2293 | 9.1× |
| SonicFastest | 1143926 | 466.10 MB/s | 981852 | 3082 | 5.6× |
| Sonic | 1149948 | 463.65 MB/s | 981959 | 3082 | 5.6× |
| Easyjson | 1319414 | 404.10 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1494266 | 356.82 MB/s | 1167079 | 5408 | 4.3× |
| JSONV2 | 2854433 | 186.79 MB/s | 745420 | 13288 | 2.3× |
| LightningDecodeAny | 3649081 | 146.11 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6430275 | 82.92 MB/s | 798692 | 17133 | 1.0× |
