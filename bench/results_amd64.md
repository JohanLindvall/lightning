# JSON Deserialization Benchmarks

- generated 2026-07-26T10:38:29Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 107112 | 1188.24 MB/s | 49760 | 3 | 12.6× |
| LightningDestructive | 109106 | 1166.53 MB/s | 49280 | 2 | 12.4× |
| Sonic | 196312 | 648.33 MB/s | 214009 | 15 | 6.9× |
| SonicFastest | 202482 | 628.57 MB/s | 214953 | 15 | 6.7× |
| Goccy | 254766 | 499.58 MB/s | 225004 | 884 | 5.3× |
| Easyjson | 268076 | 474.77 MB/s | 122864 | 14 | 5.1× |
| LightningDecodeAny | 468860 | 201.88 MB/s | 465730 | 9708 | 2.9× |
| JSONV2 | 476659 | 267.01 MB/s | 195129 | 1805 | 2.8× |
| Stdlib | 1354070 | 93.99 MB/s | 199697 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4188021 | 537.50 MB/s | 3008241 | 1158 | 8.1× |
| Lightning | 4324697 | 520.51 MB/s | 3008242 | 1158 | 7.9× |
| Sonic | 5381346 | 418.31 MB/s | 4874259 | 2584 | 6.3× |
| SonicFastest | 5616490 | 400.79 MB/s | 4871605 | 2584 | 6.0× |
| Goccy | 13560630 | 166.00 MB/s | 4204477 | 56536 | 2.5× |
| LightningDecodeAny | 13871558 | 162.28 MB/s | 19380209 | 223896 | 2.4× |
| Easyjson | 14123999 | 159.38 MB/s | 3099810 | 2120 | 2.4× |
| JSONV2 | 17229464 | 130.65 MB/s | 3123198 | 3083 | 2.0× |
| Stdlib | 33959444 | 66.29 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 567765 | 476.26 MB/s | 392945 | 568 | 7.7× |
| LightningDestructive | 568995 | 475.23 MB/s | 392945 | 568 | 7.7× |
| SonicFastest | 762419 | 354.66 MB/s | 641463 | 1147 | 5.7× |
| Sonic | 766116 | 352.95 MB/s | 641120 | 1147 | 5.7× |
| Goccy | 1839019 | 147.04 MB/s | 540875 | 8121 | 2.4× |
| Easyjson | 1843472 | 146.68 MB/s | 330273 | 749 | 2.4× |
| LightningDecodeAny | 2285020 | 118.34 MB/s | 2543877 | 29687 | 1.9× |
| JSONV2 | 2322684 | 116.42 MB/s | 348160 | 1628 | 1.9× |
| Stdlib | 4382926 | 61.69 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1184700 | 1457.93 MB/s | 767864 | 2798 | 14.5× |
| Lightning | 1212481 | 1424.52 MB/s | 767906 | 2799 | 14.2× |
| Sonic | 2221879 | 777.36 MB/s | 2694138 | 5547 | 7.8× |
| SonicFastest | 2229147 | 774.83 MB/s | 2694523 | 5547 | 7.7× |
| Goccy | 2981270 | 579.35 MB/s | 2580995 | 14603 | 5.8× |
| LightningDecodeAny | 4259869 | 117.44 MB/s | 4954731 | 76576 | 4.0× |
| Easyjson | 4355091 | 396.59 MB/s | 972032 | 5389 | 4.0× |
| JSONV2 | 4832204 | 357.44 MB/s | 1011611 | 7594 | 3.6× |
| Stdlib | 17220825 | 100.30 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1011 | 1791.97 MB/s | 0 | 0 | 16.4× |
| LightningDestructive | 1045 | 1733.51 MB/s | 0 | 0 | 15.8× |
| Easyjson | 3013 | 601.43 MB/s | 24 | 1 | 5.5× |
| Goccy | 3292 | 550.48 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6303 | 287.47 MB/s | 3347 | 38 | 2.6× |
| Sonic | 6549 | 276.67 MB/s | 3348 | 38 | 2.5× |
| JSONV2 | 8639 | 209.74 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9605 | 188.54 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16544 | 109.52 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1101 | 1645.98 MB/s | 0 | 0 | 15.1× |
| LightningDestructive | 1149 | 1577.41 MB/s | 0 | 0 | 14.5× |
| Easyjson | 2982 | 607.74 MB/s | 24 | 1 | 5.6× |
| Goccy | 3301 | 548.94 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6308 | 287.27 MB/s | 3348 | 38 | 2.6× |
| Sonic | 6430 | 281.83 MB/s | 3348 | 38 | 2.6× |
| JSONV2 | 8185 | 221.38 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9506 | 190.51 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 16654 | 108.80 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1458 | 1243.03 MB/s | 144 | 10 | 11.3× |
| Lightning | 1516 | 1195.29 MB/s | 144 | 10 | 10.9× |
| Easyjson | 3170 | 571.68 MB/s | 144 | 10 | 5.2× |
| Goccy | 3618 | 500.80 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 6404 | 282.94 MB/s | 3366 | 40 | 2.6× |
| Sonic | 6649 | 272.53 MB/s | 3366 | 40 | 2.5× |
| JSONV2 | 8612 | 210.41 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9546 | 189.72 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16483 | 109.93 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 720 | 686.30 MB/s | 160 | 1 | 9.2× |
| LightningDestructive | 729 | 677.24 MB/s | 160 | 1 | 9.1× |
| SonicFastest | 1312 | 376.60 MB/s | 1077 | 8 | 5.1× |
| Sonic | 1320 | 374.30 MB/s | 1077 | 8 | 5.0× |
| LightningDecodeAny | 1529 | 322.46 MB/s | 1296 | 26 | 4.3× |
| Easyjson | 2525 | 195.66 MB/s | 448 | 3 | 2.6× |
| Goccy | 2610 | 189.28 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3605 | 137.03 MB/s | 528 | 7 | 1.8× |
| Stdlib | 6640 | 74.40 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 458 | 502.38 MB/s | 160 | 1 | 10.4× |
| LightningDestructive | 462 | 497.84 MB/s | 160 | 1 | 10.3× |
| SonicFastest | 947 | 242.79 MB/s | 801 | 8 | 5.0× |
| Sonic | 959 | 239.77 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1284 | 178.35 MB/s | 1296 | 26 | 3.7× |
| Easyjson | 1682 | 136.72 MB/s | 448 | 3 | 2.8× |
| Goccy | 1816 | 126.62 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2775 | 82.89 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4760 | 48.32 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 88795 | 733.51 MB/s | 164882 | 105 | 7.9× |
| LightningDestructive | 94571 | 688.71 MB/s | 158661 | 100 | 7.5× |
| Sonic | 147906 | 440.36 MB/s | 235634 | 65 | 4.8× |
| SonicFastest | 147920 | 440.32 MB/s | 235637 | 65 | 4.8× |
| Goccy | 178862 | 364.15 MB/s | 228258 | 134 | 3.9× |
| LightningDecodeAny | 207002 | 257.63 MB/s | 180224 | 3245 | 3.4× |
| JSONV2 | 265975 | 244.88 MB/s | 206666 | 607 | 2.7× |
| Stdlib | 704969 | 92.39 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2597434 | 747.07 MB/s | 2813905 | 1358 | 10.7× |
| Lightning | 2672646 | 726.05 MB/s | 2813905 | 1358 | 10.4× |
| SonicFastest | 5090812 | 381.17 MB/s | 4883438 | 1736 | 5.4× |
| Goccy | 5132009 | 378.11 MB/s | 4063077 | 13509 | 5.4× |
| Sonic | 5176151 | 374.89 MB/s | 4882360 | 1736 | 5.4× |
| Easyjson | 8512893 | 227.95 MB/s | 3871264 | 15043 | 3.3× |
| LightningDecodeAny | 9997336 | 194.10 MB/s | 7064789 | 218633 | 2.8× |
| JSONV2 | 12783808 | 151.79 MB/s | 3237192 | 13947 | 2.2× |
| Stdlib | 27720864 | 70.00 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1092396 | 3046.36 MB/s | 351704 | 1286 | 22.5× |
| Lightning | 1717196 | 1937.95 MB/s | 2488907 | 2995 | 14.3× |
| SonicFastest | 2450402 | 1358.08 MB/s | 5887776 | 4263 | 10.0× |
| Sonic | 2519252 | 1320.96 MB/s | 5879231 | 4263 | 9.7× |
| LightningDecodeAny | 3842052 | 800.03 MB/s | 4886621 | 56892 | 6.4× |
| Goccy | 6252086 | 532.28 MB/s | 3948914 | 3817 | 3.9× |
| JSONV2 | 8500213 | 391.50 MB/s | 5364503 | 13243 | 2.9× |
| Stdlib | 24527474 | 135.68 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 230819 | 954.63 MB/s | 135872 | 226 | 10.5× |
| LightningDestructive | 236015 | 933.61 MB/s | 135872 | 226 | 10.3× |
| SonicFastest | 491398 | 448.41 MB/s | 349743 | 262 | 4.9× |
| Goccy | 495291 | 444.88 MB/s | 363749 | 1066 | 4.9× |
| Sonic | 500218 | 440.50 MB/s | 349579 | 262 | 4.8× |
| Easyjson | 682331 | 322.93 MB/s | 130512 | 245 | 3.5× |
| JSONV2 | 764858 | 288.09 MB/s | 129747 | 470 | 3.2× |
| LightningDecodeAny | 1037364 | 104.41 MB/s | 897522 | 11703 | 2.3× |
| Stdlib | 2421381 | 91.00 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 13707831 | 590.91 MB/s | 15730640 | 20821 | 8.0× |
| Lightning | 14245662 | 568.60 MB/s | 15730641 | 20821 | 7.7× |
| Sonic | 18546213 | 436.75 MB/s | 19859770 | 41640 | 5.9× |
| SonicFastest | 18829878 | 430.17 MB/s | 19859626 | 41640 | 5.8× |
| Goccy | 28411028 | 285.10 MB/s | 18981173 | 107155 | 3.9× |
| Easyjson | 35846273 | 225.97 MB/s | 15059617 | 41643 | 3.1× |
| LightningDecodeAny | 41552329 | 125.22 MB/s | 46191120 | 747112 | 2.6× |
| JSONV2 | 50515933 | 160.35 MB/s | 15233745 | 78972 | 2.2× |
| Stdlib | 109797025 | 73.77 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5768545 | 517.20 MB/s | 3908872 | 29356 | 9.9× |
| Lightning | 6051046 | 493.05 MB/s | 3908873 | 29356 | 9.5× |
| Sonic | 9436489 | 316.16 MB/s | 9132382 | 57804 | 6.1× |
| SonicFastest | 9466059 | 315.18 MB/s | 9132387 | 57804 | 6.1× |
| Goccy | 18716349 | 159.40 MB/s | 9815162 | 273617 | 3.1× |
| LightningDecodeAny | 18920944 | 96.94 MB/s | 23982393 | 351152 | 3.0× |
| Easyjson | 20274581 | 147.15 MB/s | 9479440 | 30115 | 2.8× |
| JSONV2 | 29636097 | 100.67 MB/s | 9257029 | 86278 | 1.9× |
| Stdlib | 57389559 | 51.99 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1369115 | 528.51 MB/s | 907601 | 3618 | 10.1× |
| Lightning | 1442130 | 501.76 MB/s | 907595 | 3618 | 9.6× |
| SonicFastest | 2191160 | 330.23 MB/s | 2371166 | 3683 | 6.3× |
| Sonic | 2194999 | 329.66 MB/s | 2370971 | 3683 | 6.3× |
| Easyjson | 5652193 | 128.02 MB/s | 2847906 | 3698 | 2.5× |
| LightningDecodeAny | 5771607 | 112.72 MB/s | 6500458 | 76546 | 2.4× |
| Goccy | 5776061 | 125.28 MB/s | 2761518 | 80270 | 2.4× |
| JSONV2 | 6965974 | 103.88 MB/s | 2704706 | 7318 | 2.0× |
| Stdlib | 13850873 | 52.24 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2010282 | 784.64 MB/s | 907594 | 3618 | 9.7× |
| LightningDestructive | 2025069 | 778.91 MB/s | 907601 | 3618 | 9.6× |
| Sonic | 2497071 | 631.68 MB/s | 3225772 | 3683 | 7.8× |
| SonicFastest | 2498012 | 631.44 MB/s | 3225335 | 3683 | 7.8× |
| LightningDecodeAny | 5199663 | 144.89 MB/s | 6500455 | 76546 | 3.7× |
| Easyjson | 6676741 | 236.25 MB/s | 2847906 | 3698 | 2.9× |
| Goccy | 7169611 | 220.01 MB/s | 3503723 | 80263 | 2.7× |
| JSONV2 | 7655367 | 206.05 MB/s | 2704555 | 7318 | 2.5× |
| Stdlib | 19410086 | 81.26 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 234836 | 639.27 MB/s | 81920 | 1 | 9.1× |
| LightningDestructive | 244552 | 613.87 MB/s | 81920 | 1 | 8.8× |
| Sonic | 391393 | 383.56 MB/s | 407387 | 16 | 5.5× |
| SonicFastest | 397269 | 377.89 MB/s | 407291 | 16 | 5.4× |
| LightningDecodeAny | 580844 | 258.45 MB/s | 745765 | 10016 | 3.7× |
| Goccy | 1035730 | 144.95 MB/s | 325124 | 10005 | 2.1× |
| JSONV2 | 1214474 | 123.61 MB/s | 357727 | 20 | 1.8× |
| Stdlib | 2140425 | 70.14 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 34177 | 822.69 MB/s | 29216 | 103 | 10.4× |
| LightningDestructive | 35552 | 790.87 MB/s | 29088 | 101 | 10.0× |
| Sonic | 60364 | 465.79 MB/s | 59495 | 83 | 5.9× |
| SonicFastest | 60834 | 462.19 MB/s | 59494 | 83 | 5.8× |
| Goccy | 82264 | 341.79 MB/s | 59272 | 188 | 4.3× |
| Easyjson | 84620 | 332.27 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 142792 | 196.91 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 170951 | 164.47 MB/s | 140592 | 2643 | 2.1× |
| Stdlib | 353875 | 79.45 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1993 | 1167.93 MB/s | 32 | 1 | 13.3× |
| LightningDestructive | 2064 | 1127.99 MB/s | 32 | 1 | 12.9× |
| SonicFastest | 4911 | 474.06 MB/s | 3709 | 4 | 5.4× |
| Sonic | 4919 | 473.23 MB/s | 3711 | 4 | 5.4× |
| Goccy | 4929 | 472.27 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 5719 | 407.04 MB/s | 192 | 2 | 4.6× |
| JSONV2 | 8842 | 263.27 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 11158 | 151.02 MB/s | 10200 | 195 | 2.4× |
| Stdlib | 26544 | 87.70 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 214 | 882.46 MB/s | 0 | 0 | 13.1× |
| LightningDestructive | 222 | 851.20 MB/s | 0 | 0 | 12.6× |
| Goccy | 440 | 429.83 MB/s | 304 | 2 | 6.4× |
| Easyjson | 585 | 323.22 MB/s | 0 | 0 | 4.8× |
| SonicFastest | 649 | 291.03 MB/s | 341 | 3 | 4.3× |
| Sonic | 651 | 290.29 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1029 | 183.76 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1371 | 97.72 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2808 | 67.32 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1469 | 1491.78 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 1521 | 1440.44 MB/s | 0 | 0 | 12.5× |
| Easyjson | 3780 | 579.62 MB/s | 24 | 1 | 5.0× |
| Goccy | 3806 | 575.66 MB/s | 2864 | 4 | 5.0× |
| SonicFastest | 6860 | 319.38 MB/s | 3604 | 38 | 2.8× |
| Sonic | 6994 | 313.29 MB/s | 3604 | 38 | 2.7× |
| JSONV2 | 8431 | 259.87 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 9423 | 192.20 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 19038 | 115.09 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 724522 | 704.57 MB/s | 703298 | 1010 | 9.9× |
| Lightning | 768027 | 664.66 MB/s | 703297 | 1010 | 9.4× |
| SonicFastest | 1292543 | 394.94 MB/s | 1307696 | 2014 | 5.6× |
| Sonic | 1297862 | 393.32 MB/s | 1308518 | 2014 | 5.6× |
| Goccy | 1346619 | 379.08 MB/s | 1138189 | 5006 | 5.3× |
| Easyjson | 1741162 | 293.18 MB/s | 863777 | 3012 | 4.1× |
| JSONV2 | 3611854 | 141.33 MB/s | 1075971 | 12645 | 2.0× |
| LightningDecodeAny | 3859081 | 119.58 MB/s | 2929687 | 64018 | 1.9× |
| Stdlib | 7203572 | 70.86 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 547 | 36150.14 MB/s | 0 | 0 | 302.4× |
| LightningDestructive | 827 | 23937.10 MB/s | 0 | 0 | 200.2× |
| SonicFastest | 6769 | 2923.49 MB/s | 21137 | 3 | 24.5× |
| Goccy | 30751 | 643.52 MB/s | 20492 | 2 | 5.4× |
| JSONV2 | 33312 | 594.05 MB/s | 8 | 1 | 5.0× |
| Sonic | 33669 | 587.75 MB/s | 20627 | 3 | 4.9× |
| LightningDecodeAny | 102672 | 192.73 MB/s | 116864 | 2015 | 1.6× |
| Easyjson | 109229 | 181.17 MB/s | 0 | 0 | 1.5× |
| Stdlib | 165514 | 119.56 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2308 | 7851.69 MB/s | 432 | 2 | 54.9× |
| LightningDestructive | 2374 | 7635.41 MB/s | 0 | 0 | 53.4× |
| Easyjson | 5055 | 3585.39 MB/s | 432 | 2 | 25.1× |
| SonicFastest | 8457 | 2143.15 MB/s | 20471 | 5 | 15.0× |
| Sonic | 8510 | 2129.81 MB/s | 20450 | 5 | 14.9× |
| LightningDecodeAny | 19668 | 909.18 MB/s | 29088 | 191 | 6.4× |
| Goccy | 22584 | 802.51 MB/s | 19460 | 2 | 5.6× |
| JSONV2 | 51312 | 353.21 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 126672 | 143.08 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2466409 | 814.34 MB/s | 3089821 | 6821 | 9.0× |
| Lightning | 2609221 | 769.77 MB/s | 3091534 | 6827 | 8.5× |
| SonicFastest | 4323589 | 464.54 MB/s | 5161033 | 7085 | 5.1× |
| Sonic | 4389370 | 457.58 MB/s | 5161500 | 7085 | 5.1× |
| Goccy | 4887906 | 410.91 MB/s | 5409987 | 15830 | 4.5× |
| Easyjson | 6025474 | 333.33 MB/s | 2981481 | 7439 | 3.7× |
| LightningDecodeAny | 7447452 | 153.38 MB/s | 8498329 | 134008 | 3.0× |
| JSONV2 | 7770507 | 258.48 MB/s | 3173678 | 14563 | 2.9× |
| Stdlib | 22210344 | 90.43 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 853 | 643.78 MB/s | 480 | 1 | 7.9× |
| LightningDestructive | 866 | 634.06 MB/s | 480 | 1 | 7.8× |
| LightningDecodeAny | 1878 | 291.84 MB/s | 2021 | 46 | 3.6× |
| Easyjson | 2218 | 247.49 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2472 | 222.12 MB/s | 2260 | 8 | 2.7× |
| Sonic | 2477 | 221.63 MB/s | 2262 | 8 | 2.7× |
| Goccy | 3344 | 164.18 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3410 | 160.98 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 6736 | 81.50 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 529365 | 1192.97 MB/s | 402728 | 545 | 12.1× |
| Lightning | 593293 | 1064.42 MB/s | 451257 | 857 | 10.8× |
| SonicFastest | 989926 | 637.94 MB/s | 1065692 | 814 | 6.5× |
| Sonic | 995343 | 634.47 MB/s | 1065999 | 814 | 6.5× |
| Goccy | 1308400 | 482.66 MB/s | 988373 | 1200 | 4.9× |
| Easyjson | 1469964 | 429.61 MB/s | 422504 | 936 | 4.4× |
| JSONV2 | 2389466 | 264.29 MB/s | 571592 | 3144 | 2.7× |
| LightningDecodeAny | 2633758 | 177.28 MB/s | 2077366 | 30126 | 2.4× |
| Stdlib | 6426423 | 98.27 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 735555 | 764.60 MB/s | 579337 | 429 | 8.4× |
| Lightning | 951179 | 591.27 MB/s | 802709 | 1235 | 6.5× |
| Sonic | 1289065 | 436.29 MB/s | 1345222 | 1184 | 4.8× |
| SonicFastest | 1290775 | 435.71 MB/s | 1345045 | 1184 | 4.8× |
| Goccy | 1527713 | 368.14 MB/s | 1035749 | 1028 | 4.0× |
| Easyjson | 2215636 | 253.84 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 3120055 | 180.26 MB/s | 2181319 | 30126 | 2.0× |
| JSONV2 | 3268896 | 172.05 MB/s | 927401 | 3482 | 1.9× |
| Stdlib | 6142823 | 91.56 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 629077 | 847.56 MB/s | 333416 | 2084 | 10.5× |
| Lightning | 715361 | 745.33 MB/s | 368224 | 2293 | 9.2× |
| SonicFastest | 1122263 | 475.09 MB/s | 982107 | 3082 | 5.9× |
| Sonic | 1126872 | 473.15 MB/s | 982120 | 3082 | 5.9× |
| Easyjson | 1344088 | 396.68 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1464036 | 364.18 MB/s | 1167093 | 5409 | 4.5× |
| JSONV2 | 3073422 | 173.48 MB/s | 745418 | 13288 | 2.1× |
| LightningDecodeAny | 3722502 | 143.23 MB/s | 2991145 | 50076 | 1.8× |
| Stdlib | 6594746 | 80.85 MB/s | 798692 | 17133 | 1.0× |
