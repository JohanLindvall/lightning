# JSON Deserialization Benchmarks

- generated 2026-06-24T16:33:36Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104793 | 1214.54 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 105141 | 1210.52 MB/s | 49280 | 2 | 10.5× |
| SonicFastest | 181512 | 701.19 MB/s | 192718 | 10 | 6.1× |
| Sonic | 181859 | 699.85 MB/s | 192845 | 10 | 6.1× |
| Goccy | 188720 | 674.41 MB/s | 224418 | 884 | 5.9× |
| Easyjson | 212203 | 599.78 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 417601 | 304.78 MB/s | 195119 | 1805 | 2.7× |
| LightningDecodeAny | 460233 | 205.66 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1107273 | 114.94 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3749309 | 600.39 MB/s | 3122872 | 3081 | 7.0× |
| Lightning | 3838240 | 586.48 MB/s | 3122874 | 3081 | 6.8× |
| Sonic | 4652182 | 483.87 MB/s | 15233751 | 970 | 5.6× |
| SonicFastest | 4668708 | 482.16 MB/s | 15233741 | 970 | 5.6× |
| Goccy | 10557019 | 213.23 MB/s | 4129380 | 56533 | 2.5× |
| Easyjson | 11116787 | 202.49 MB/s | 3099808 | 2120 | 2.3× |
| LightningDecodeAny | 11971066 | 188.04 MB/s | 7938299 | 281383 | 2.2× |
| JSONV2 | 16379326 | 137.43 MB/s | 3123223 | 3083 | 1.6× |
| Stdlib | 26114804 | 86.20 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 490751 | 551.00 MB/s | 348024 | 1627 | 6.8× |
| LightningDestructive | 495179 | 546.07 MB/s | 348024 | 1627 | 6.8× |
| SonicFastest | 651460 | 415.07 MB/s | 515254 | 968 | 5.1× |
| Sonic | 652341 | 414.51 MB/s | 511941 | 968 | 5.1× |
| Goccy | 1423661 | 189.93 MB/s | 543236 | 8122 | 2.4× |
| Easyjson | 1448792 | 186.64 MB/s | 330272 | 749 | 2.3× |
| LightningDecodeAny | 1624354 | 166.47 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2140080 | 126.35 MB/s | 348149 | 1628 | 1.6× |
| Stdlib | 3354711 | 80.60 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1533972 | 1125.97 MB/s | 519314 | 2693 | 8.7× |
| LightningDestructive | 1541210 | 1120.68 MB/s | 519272 | 2692 | 8.6× |
| Sonic | 2051900 | 841.76 MB/s | 2723811 | 4020 | 6.5× |
| SonicFastest | 2065132 | 836.36 MB/s | 2730079 | 4020 | 6.4× |
| Goccy | 2393844 | 721.52 MB/s | 2583260 | 14605 | 5.5× |
| JSONV2 | 4199954 | 411.24 MB/s | 1011634 | 7594 | 3.2× |
| Easyjson | 4250248 | 406.38 MB/s | 972032 | 5389 | 3.1× |
| LightningDecodeAny | 4748967 | 105.35 MB/s | 4976204 | 81466 | 2.8× |
| Stdlib | 13270741 | 130.15 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1190 | 1522.66 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1206 | 1502.09 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2554 | 709.37 MB/s | 24 | 1 | 5.5× |
| Goccy | 2874 | 630.45 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6104 | 296.87 MB/s | 3850 | 40 | 2.3× |
| Sonic | 6122 | 296.00 MB/s | 3888 | 40 | 2.3× |
| JSONV2 | 7869 | 230.28 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8591 | 210.79 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14056 | 128.91 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1221 | 1484.25 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1234 | 1468.74 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2548 | 711.18 MB/s | 24 | 1 | 5.5× |
| Goccy | 2873 | 630.69 MB/s | 2608 | 4 | 4.9× |
| Sonic | 5962 | 303.91 MB/s | 3772 | 40 | 2.4× |
| SonicFastest | 6006 | 301.72 MB/s | 3822 | 40 | 2.3× |
| JSONV2 | 7844 | 231.00 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8294 | 218.36 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14100 | 128.51 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1391 | 1302.26 MB/s | 144 | 10 | 10.1× |
| LightningDestructive | 1448 | 1251.71 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2761 | 656.31 MB/s | 144 | 10 | 5.1× |
| Goccy | 2965 | 611.09 MB/s | 2600 | 5 | 4.7× |
| Sonic | 6192 | 292.64 MB/s | 3871 | 42 | 2.3× |
| SonicFastest | 6213 | 291.65 MB/s | 3901 | 42 | 2.3× |
| JSONV2 | 7992 | 226.73 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8648 | 209.41 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14014 | 129.30 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 742 | 665.92 MB/s | 160 | 1 | 7.5× |
| LightningDestructive | 746 | 662.59 MB/s | 160 | 1 | 7.4× |
| Sonic | 1235 | 399.87 MB/s | 966 | 6 | 4.5× |
| SonicFastest | 1243 | 397.54 MB/s | 986 | 6 | 4.5× |
| LightningDecodeAny | 1676 | 294.23 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2203 | 224.24 MB/s | 448 | 3 | 2.5× |
| Goccy | 2495 | 198.00 MB/s | 856 | 23 | 2.2× |
| JSONV2 | 3198 | 154.47 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5533 | 89.29 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 472 | 487.06 MB/s | 160 | 1 | 8.8× |
| LightningDestructive | 476 | 483.23 MB/s | 160 | 1 | 8.7× |
| Sonic | 887 | 259.27 MB/s | 650 | 6 | 4.7× |
| SonicFastest | 892 | 257.70 MB/s | 659 | 6 | 4.7× |
| Easyjson | 1388 | 165.73 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1409 | 162.50 MB/s | 1536 | 30 | 2.9× |
| Goccy | 1627 | 141.41 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2332 | 98.63 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4156 | 55.35 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 73848 | 881.98 MB/s | 166212 | 102 | 7.4× |
| Lightning | 74837 | 870.32 MB/s | 172432 | 107 | 7.3× |
| Sonic | 97867 | 665.52 MB/s | 155760 | 75 | 5.6× |
| SonicFastest | 97898 | 665.30 MB/s | 155966 | 75 | 5.6× |
| Goccy | 148315 | 439.15 MB/s | 229380 | 134 | 3.7× |
| LightningDecodeAny | 189848 | 280.90 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 227074 | 286.83 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 547177 | 119.03 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2635457 | 736.29 MB/s | 2846865 | 2238 | 8.9× |
| Lightning | 2685737 | 722.51 MB/s | 2846866 | 2238 | 8.8× |
| SonicFastest | 4669978 | 415.52 MB/s | 14606973 | 1407 | 5.0× |
| Goccy | 4785755 | 405.47 MB/s | 4065172 | 13510 | 4.9× |
| Sonic | 4884544 | 397.27 MB/s | 14606972 | 1407 | 4.8× |
| Easyjson | 7586418 | 255.78 MB/s | 3871267 | 15043 | 3.1× |
| LightningDecodeAny | 9582297 | 202.51 MB/s | 7013524 | 219937 | 2.5× |
| JSONV2 | 11375200 | 170.59 MB/s | 3237222 | 13947 | 2.1× |
| Stdlib | 23529216 | 82.47 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1109179 | 3000.27 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1813204 | 1835.33 MB/s | 2488906 | 2995 | 11.5× |
| Sonic | 2679559 | 1241.93 MB/s | 6483975 | 4248 | 7.8× |
| SonicFastest | 2693679 | 1235.42 MB/s | 6458111 | 4248 | 7.7× |
| LightningDecodeAny | 3808110 | 807.16 MB/s | 4886620 | 56892 | 5.5× |
| Goccy | 4492818 | 740.70 MB/s | 3948909 | 3816 | 4.6× |
| JSONV2 | 7491799 | 444.20 MB/s | 5364515 | 13243 | 2.8× |
| Stdlib | 20823391 | 159.81 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222273 | 991.33 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 224054 | 983.45 MB/s | 136897 | 228 | 9.1× |
| Sonic | 381154 | 578.10 MB/s | 297436 | 398 | 5.4× |
| SonicFastest | 382838 | 575.56 MB/s | 301376 | 398 | 5.3× |
| Goccy | 435487 | 505.98 MB/s | 364354 | 1067 | 4.7× |
| Easyjson | 545893 | 403.64 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 735477 | 299.60 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 894273 | 121.12 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2043819 | 107.81 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12788038 | 633.41 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13267577 | 610.51 MB/s | 15059836 | 51649 | 6.7× |
| Sonic | 16751819 | 483.53 MB/s | 70873026 | 40014 | 5.3× |
| SonicFastest | 16831097 | 481.25 MB/s | 70901802 | 40014 | 5.3× |
| Goccy | 23380056 | 346.45 MB/s | 16930622 | 107148 | 3.8× |
| Easyjson | 31122043 | 260.27 MB/s | 15059617 | 41643 | 2.9× |
| LightningDecodeAny | 40588570 | 128.19 MB/s | 34333352 | 912810 | 2.2× |
| JSONV2 | 43747621 | 185.15 MB/s | 15233739 | 78972 | 2.0× |
| Stdlib | 89279825 | 90.73 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6076388 | 490.99 MB/s | 3985336 | 30015 | 7.8× |
| Lightning | 6280662 | 475.02 MB/s | 3985336 | 30015 | 7.5× |
| SonicFastest | 8670276 | 344.10 MB/s | 26497418 | 56760 | 5.4× |
| Sonic | 8670432 | 344.10 MB/s | 26608151 | 56760 | 5.4× |
| Goccy | 16534751 | 180.44 MB/s | 10625938 | 273649 | 2.9× |
| Easyjson | 16670466 | 178.97 MB/s | 9479441 | 30115 | 2.8× |
| LightningDecodeAny | 19094628 | 96.06 MB/s | 20023837 | 408557 | 2.5× |
| JSONV2 | 24502824 | 121.76 MB/s | 9257139 | 86278 | 1.9× |
| Stdlib | 47207223 | 63.20 MB/s | 9258115 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1387455 | 521.53 MB/s | 907388 | 3618 | 8.4× |
| LightningDestructive | 1394622 | 518.85 MB/s | 907394 | 3618 | 8.4× |
| Sonic | 1774882 | 407.69 MB/s | 3181638 | 7226 | 6.6× |
| SonicFastest | 1780621 | 406.37 MB/s | 3185796 | 7226 | 6.6× |
| Easyjson | 4307012 | 168.00 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4406949 | 147.62 MB/s | 5752201 | 80172 | 2.6× |
| Goccy | 4699728 | 153.97 MB/s | 2770680 | 80271 | 2.5× |
| JSONV2 | 5908674 | 122.46 MB/s | 2704609 | 7318 | 2.0× |
| Stdlib | 11671477 | 62.00 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1946254 | 810.46 MB/s | 907393 | 3618 | 8.1× |
| Lightning | 1973335 | 799.33 MB/s | 907386 | 3618 | 8.0× |
| SonicFastest | 2252253 | 700.34 MB/s | 5786083 | 7226 | 7.0× |
| Sonic | 2261112 | 697.60 MB/s | 5790379 | 7226 | 7.0× |
| LightningDecodeAny | 3902157 | 193.07 MB/s | 5752201 | 80172 | 4.0× |
| Easyjson | 5606533 | 281.34 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5613618 | 280.99 MB/s | 3538655 | 80265 | 2.8× |
| JSONV2 | 6715742 | 234.87 MB/s | 2704595 | 7318 | 2.3× |
| Stdlib | 15724854 | 100.31 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 223851 | 670.64 MB/s | 81920 | 1 | 8.1× |
| LightningDestructive | 225077 | 666.99 MB/s | 81920 | 1 | 8.1× |
| Sonic | 278131 | 539.76 MB/s | 261951 | 6 | 6.6× |
| SonicFastest | 278583 | 538.88 MB/s | 264062 | 6 | 6.5× |
| LightningDecodeAny | 478136 | 313.97 MB/s | 746004 | 10020 | 3.8× |
| Goccy | 862582 | 174.04 MB/s | 325188 | 10004 | 2.1× |
| JSONV2 | 1105301 | 135.82 MB/s | 357715 | 20 | 1.6× |
| Stdlib | 1823161 | 82.34 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33002 | 851.98 MB/s | 30272 | 105 | 9.2× |
| LightningDestructive | 33162 | 847.87 MB/s | 30144 | 103 | 9.2× |
| SonicFastest | 64086 | 438.74 MB/s | 47009 | 103 | 4.7× |
| Sonic | 64352 | 436.93 MB/s | 47683 | 103 | 4.7× |
| Easyjson | 68162 | 412.50 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71894 | 391.09 MB/s | 59215 | 188 | 4.2× |
| JSONV2 | 134776 | 208.62 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 153284 | 183.43 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 304135 | 92.45 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1950 | 1193.66 MB/s | 32 | 1 | 11.7× |
| LightningDestructive | 2023 | 1150.71 MB/s | 32 | 1 | 11.3× |
| Goccy | 4124 | 564.47 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4232 | 550.13 MB/s | 192 | 2 | 5.4× |
| Sonic | 5091 | 457.25 MB/s | 4275 | 6 | 4.5× |
| SonicFastest | 5147 | 452.30 MB/s | 4333 | 6 | 4.4× |
| JSONV2 | 8448 | 275.57 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10391 | 162.17 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22785 | 102.17 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 862.03 MB/s | 0 | 0 | 11.4× |
| LightningDestructive | 221 | 855.17 MB/s | 0 | 0 | 11.3× |
| Goccy | 386 | 489.05 MB/s | 304 | 2 | 6.5× |
| Easyjson | 482 | 392.25 MB/s | 0 | 0 | 5.2× |
| SonicFastest | 806 | 234.63 MB/s | 499 | 4 | 3.1× |
| Sonic | 809 | 233.75 MB/s | 499 | 4 | 3.1× |
| JSONV2 | 1039 | 181.99 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1246 | 107.56 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2495 | 75.76 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1542 | 1420.98 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1561 | 1403.80 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3188 | 687.21 MB/s | 24 | 1 | 5.0× |
| Goccy | 3258 | 672.42 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6418 | 341.37 MB/s | 4012 | 40 | 2.5× |
| Sonic | 6432 | 340.63 MB/s | 3986 | 40 | 2.5× |
| JSONV2 | 8093 | 270.73 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8301 | 218.15 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16031 | 136.67 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 712967 | 715.99 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 734283 | 695.20 MB/s | 703780 | 1012 | 8.4× |
| SonicFastest | 1159640 | 440.20 MB/s | 890978 | 2006 | 5.3× |
| Sonic | 1165136 | 438.13 MB/s | 900421 | 2006 | 5.3× |
| Goccy | 1171381 | 435.79 MB/s | 1142308 | 5007 | 5.2× |
| Easyjson | 1569471 | 325.25 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3229956 | 158.04 MB/s | 1076006 | 12646 | 1.9× |
| LightningDecodeAny | 3581662 | 128.84 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6135105 | 83.21 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14648.48 MB/s | 0 | 0 | 93.4× |
| LightningDestructive | 1409 | 14043.31 MB/s | 0 | 0 | 89.5× |
| Goccy | 20351 | 972.36 MB/s | 20491 | 2 | 6.2× |
| Sonic | 27070 | 731.04 MB/s | 22397 | 4 | 4.7× |
| SonicFastest | 27234 | 726.63 MB/s | 22793 | 4 | 4.6× |
| JSONV2 | 29709 | 666.09 MB/s | 8 | 1 | 4.2× |
| LightningDecodeAny | 79517 | 248.85 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 88334 | 224.03 MB/s | 0 | 0 | 1.4× |
| Stdlib | 126150 | 156.87 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2661 | 6810.17 MB/s | 0 | 0 | 38.9× |
| Lightning | 2842 | 6377.39 MB/s | 432 | 2 | 36.4× |
| Easyjson | 3978 | 4556.39 MB/s | 432 | 2 | 26.0× |
| Sonic | 9901 | 1830.53 MB/s | 23036 | 6 | 10.4× |
| SonicFastest | 10056 | 1802.23 MB/s | 23456 | 6 | 10.3× |
| Goccy | 15620 | 1160.31 MB/s | 19459 | 2 | 6.6× |
| LightningDecodeAny | 16928 | 1056.34 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 45887 | 394.97 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103405 | 175.27 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2250036 | 892.65 MB/s | 2960389 | 7411 | 8.3× |
| Lightning | 2283647 | 879.51 MB/s | 2962101 | 7417 | 8.2× |
| Goccy | 4280472 | 469.22 MB/s | 5413174 | 15833 | 4.4× |
| SonicFastest | 4411948 | 455.24 MB/s | 11007636 | 13683 | 4.2× |
| Sonic | 4433342 | 453.04 MB/s | 10980115 | 13683 | 4.2× |
| Easyjson | 4914783 | 408.66 MB/s | 2981483 | 7439 | 3.8× |
| JSONV2 | 6962872 | 288.46 MB/s | 3173684 | 14563 | 2.7× |
| LightningDecodeAny | 7507548 | 152.15 MB/s | 7386074 | 134751 | 2.5× |
| Stdlib | 18664976 | 107.61 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 893 | 614.78 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 895 | 613.13 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1954 | 280.38 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2179 | 251.91 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2694 | 203.77 MB/s | 1956 | 26 | 2.1× |
| Sonic | 2712 | 202.43 MB/s | 1976 | 26 | 2.1× |
| Goccy | 3062 | 179.30 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3291 | 166.80 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5662 | 96.97 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 495177 | 1275.33 MB/s | 364472 | 566 | 10.9× |
| Lightning | 548448 | 1151.46 MB/s | 413002 | 878 | 9.8× |
| SonicFastest | 1006521 | 627.42 MB/s | 1010522 | 1102 | 5.4× |
| Sonic | 1016723 | 621.13 MB/s | 1012200 | 1102 | 5.3× |
| Easyjson | 1142289 | 552.85 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1163772 | 542.64 MB/s | 988203 | 1201 | 4.6× |
| JSONV2 | 2149643 | 293.78 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2402396 | 194.35 MB/s | 2010201 | 30295 | 2.2× |
| Stdlib | 5401954 | 116.90 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 706372 | 796.19 MB/s | 544253 | 448 | 7.6× |
| Lightning | 882954 | 636.96 MB/s | 767621 | 1254 | 6.1× |
| Sonic | 1031912 | 545.02 MB/s | 961474 | 1476 | 5.2× |
| SonicFastest | 1036647 | 542.53 MB/s | 971966 | 1476 | 5.2× |
| Goccy | 1339897 | 419.74 MB/s | 1042660 | 1030 | 4.0× |
| Easyjson | 1740705 | 323.09 MB/s | 775154 | 1254 | 3.1× |
| JSONV2 | 2761350 | 203.67 MB/s | 927442 | 3482 | 1.9× |
| LightningDecodeAny | 2805098 | 200.49 MB/s | 2114151 | 30295 | 1.9× |
| Stdlib | 5351370 | 105.10 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 656491 | 812.16 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 679195 | 785.01 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1119851 | 476.12 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1158625 | 460.18 MB/s | 1030894 | 4351 | 4.8× |
| SonicFastest | 1162213 | 458.76 MB/s | 1032152 | 4351 | 4.7× |
| Goccy | 1306742 | 408.02 MB/s | 1167225 | 5409 | 4.2× |
| JSONV2 | 2542413 | 209.71 MB/s | 745450 | 13288 | 2.2× |
| LightningDecodeAny | 3420712 | 155.87 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5505755 | 96.84 MB/s | 798692 | 17133 | 1.0× |
