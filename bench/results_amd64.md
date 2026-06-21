# JSON Deserialization Benchmarks

- generated 2026-06-21T21:08:54Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105373 | 1207.86 MB/s | 49760 | 3 | 12.5× |
| LightningDestructive | 106021 | 1200.47 MB/s | 49280 | 2 | 12.4× |
| Sonic | 199425 | 638.21 MB/s | 214980 | 15 | 6.6× |
| SonicFastest | 200395 | 635.12 MB/s | 215218 | 15 | 6.6× |
| Easyjson | 239659 | 531.07 MB/s | 122864 | 14 | 5.5× |
| Goccy | 247675 | 513.88 MB/s | 225361 | 884 | 5.3× |
| LightningDecodeAny | 472241 | 200.43 MB/s | 465586 | 9714 | 2.8× |
| JSONV2 | 477184 | 266.72 MB/s | 195127 | 1805 | 2.8× |
| Stdlib | 1314120 | 96.85 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4072821 | 552.70 MB/s | 3122872 | 3081 | 8.3× |
| Lightning | 4138798 | 543.89 MB/s | 3122872 | 3081 | 8.2× |
| Sonic | 5105797 | 440.88 MB/s | 4872080 | 2584 | 6.7× |
| SonicFastest | 5326534 | 422.61 MB/s | 4872632 | 2584 | 6.4× |
| LightningDecodeAny | 12213176 | 184.31 MB/s | 7938298 | 281383 | 2.8× |
| Goccy | 12771748 | 176.25 MB/s | 4170642 | 56534 | 2.7× |
| Easyjson | 13225878 | 170.20 MB/s | 3099809 | 2120 | 2.6× |
| JSONV2 | 17347730 | 129.76 MB/s | 3123215 | 3083 | 2.0× |
| Stdlib | 33953694 | 66.30 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 571063 | 473.51 MB/s | 348025 | 1627 | 7.6× |
| Lightning | 571067 | 473.51 MB/s | 348025 | 1627 | 7.6× |
| SonicFastest | 745373 | 362.78 MB/s | 641194 | 1147 | 5.9× |
| Sonic | 746439 | 362.26 MB/s | 641026 | 1147 | 5.8× |
| LightningDecodeAny | 1700386 | 159.02 MB/s | 1011488 | 37901 | 2.6× |
| Easyjson | 1744749 | 154.98 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1775674 | 152.28 MB/s | 541407 | 8122 | 2.5× |
| JSONV2 | 2315953 | 116.76 MB/s | 348159 | 1628 | 1.9× |
| Stdlib | 4365739 | 61.94 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1284012 | 1345.16 MB/s | 959848 | 5881 | 13.7× |
| Lightning | 1346348 | 1282.88 MB/s | 959889 | 5882 | 13.0× |
| SonicFastest | 2084533 | 828.58 MB/s | 2695099 | 5547 | 8.4× |
| Sonic | 2090361 | 826.27 MB/s | 2695063 | 5547 | 8.4× |
| Goccy | 2957963 | 583.92 MB/s | 2580961 | 14603 | 5.9× |
| Easyjson | 4207799 | 410.48 MB/s | 972032 | 5389 | 4.2× |
| LightningDecodeAny | 4429000 | 112.96 MB/s | 4976204 | 81466 | 4.0× |
| JSONV2 | 4538734 | 380.55 MB/s | 1011611 | 7594 | 3.9× |
| Stdlib | 17560009 | 98.36 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1004 | 1804.02 MB/s | 0 | 0 | 15.9× |
| LightningDestructive | 1044 | 1736.10 MB/s | 0 | 0 | 15.3× |
| Easyjson | 2857 | 634.25 MB/s | 24 | 1 | 5.6× |
| Goccy | 3462 | 523.41 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6211 | 291.75 MB/s | 3345 | 38 | 2.6× |
| Sonic | 6398 | 283.23 MB/s | 3344 | 38 | 2.5× |
| JSONV2 | 8441 | 214.66 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9376 | 193.14 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 15966 | 113.49 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1101 | 1646.45 MB/s | 0 | 0 | 14.6× |
| LightningDestructive | 1155 | 1569.37 MB/s | 0 | 0 | 13.9× |
| Easyjson | 2839 | 638.35 MB/s | 24 | 1 | 5.7× |
| Goccy | 3421 | 529.68 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6101 | 296.99 MB/s | 3342 | 38 | 2.6× |
| Sonic | 6274 | 288.83 MB/s | 3342 | 38 | 2.6× |
| JSONV2 | 8333 | 217.45 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9362 | 193.44 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16054 | 112.87 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1324 | 1368.20 MB/s | 144 | 10 | 12.2× |
| LightningDestructive | 1434 | 1263.95 MB/s | 144 | 10 | 11.3× |
| Goccy | 3207 | 565.04 MB/s | 2600 | 5 | 5.0× |
| Easyjson | 3265 | 555.03 MB/s | 144 | 10 | 5.0× |
| SonicFastest | 6199 | 292.30 MB/s | 3366 | 40 | 2.6× |
| Sonic | 6396 | 283.31 MB/s | 3366 | 40 | 2.5× |
| JSONV2 | 8316 | 217.89 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9248 | 195.83 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16162 | 112.11 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 722 | 684.05 MB/s | 160 | 1 | 9.6× |
| LightningDestructive | 734 | 672.52 MB/s | 160 | 1 | 9.4× |
| SonicFastest | 1269 | 389.25 MB/s | 1075 | 8 | 5.5× |
| Sonic | 1271 | 388.55 MB/s | 1075 | 8 | 5.4× |
| LightningDecodeAny | 1783 | 276.49 MB/s | 1536 | 30 | 3.9× |
| Easyjson | 2524 | 195.76 MB/s | 448 | 3 | 2.7× |
| Goccy | 2693 | 183.42 MB/s | 856 | 23 | 2.6× |
| JSONV2 | 3299 | 149.76 MB/s | 528 | 7 | 2.1× |
| Stdlib | 6924 | 71.34 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 489 | 470.13 MB/s | 160 | 1 | 10.1× |
| LightningDestructive | 494 | 465.85 MB/s | 160 | 1 | 10.0× |
| Sonic | 919 | 250.35 MB/s | 801 | 8 | 5.4× |
| SonicFastest | 922 | 249.49 MB/s | 801 | 8 | 5.4× |
| LightningDecodeAny | 1601 | 143.02 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1705 | 134.86 MB/s | 448 | 3 | 2.9× |
| Goccy | 1839 | 125.06 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2608 | 88.20 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4945 | 46.52 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 92754 | 702.20 MB/s | 172434 | 107 | 7.3× |
| LightningDestructive | 99113 | 657.15 MB/s | 166213 | 102 | 6.8× |
| Sonic | 148465 | 438.70 MB/s | 235835 | 65 | 4.5× |
| SonicFastest | 148886 | 437.46 MB/s | 235770 | 65 | 4.5× |
| Goccy | 186285 | 349.64 MB/s | 228020 | 134 | 3.6× |
| LightningDecodeAny | 199784 | 266.93 MB/s | 176961 | 3252 | 3.4× |
| JSONV2 | 263478 | 247.20 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 672909 | 96.79 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2498344 | 776.70 MB/s | 2846864 | 2238 | 11.2× |
| Lightning | 2587393 | 749.97 MB/s | 2846866 | 2238 | 10.8× |
| SonicFastest | 4691407 | 413.62 MB/s | 4878689 | 1736 | 6.0× |
| Sonic | 4745767 | 408.88 MB/s | 4881597 | 1736 | 5.9× |
| Goccy | 4916728 | 394.67 MB/s | 4061841 | 13509 | 5.7× |
| Easyjson | 8256173 | 235.03 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 9942190 | 195.18 MB/s | 7013525 | 219937 | 2.8× |
| JSONV2 | 12634027 | 153.59 MB/s | 3237200 | 13947 | 2.2× |
| Stdlib | 27966851 | 69.38 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1053922 | 3157.57 MB/s | 351704 | 1286 | 22.2× |
| Lightning | 1593525 | 2088.35 MB/s | 2488907 | 2995 | 14.7× |
| SonicFastest | 2004545 | 1660.14 MB/s | 5896590 | 4263 | 11.7× |
| Sonic | 2011909 | 1654.07 MB/s | 5896516 | 4263 | 11.6× |
| LightningDecodeAny | 3538613 | 868.64 MB/s | 4886622 | 56892 | 6.6× |
| Goccy | 4674507 | 711.91 MB/s | 3948915 | 3817 | 5.0× |
| JSONV2 | 7394185 | 450.06 MB/s | 5364504 | 13243 | 3.2× |
| Stdlib | 23423047 | 142.08 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 231540 | 951.65 MB/s | 136896 | 228 | 10.6× |
| LightningDestructive | 238075 | 925.53 MB/s | 136896 | 228 | 10.3× |
| Goccy | 480333 | 458.74 MB/s | 364024 | 1066 | 5.1× |
| SonicFastest | 488164 | 451.38 MB/s | 350088 | 262 | 5.0× |
| Sonic | 490363 | 449.35 MB/s | 350120 | 262 | 5.0× |
| Easyjson | 629797 | 349.87 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 755731 | 291.57 MB/s | 129746 | 470 | 3.2× |
| LightningDecodeAny | 1011712 | 107.06 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2453765 | 89.80 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14713176 | 550.53 MB/s | 15059834 | 51649 | 7.6× |
| Lightning | 15414299 | 525.49 MB/s | 15059843 | 51649 | 7.3× |
| SonicFastest | 17881557 | 452.98 MB/s | 19856834 | 41640 | 6.3× |
| Sonic | 17926354 | 451.85 MB/s | 19857696 | 41640 | 6.3× |
| Goccy | 26966880 | 300.37 MB/s | 19043690 | 107155 | 4.2× |
| Easyjson | 34462069 | 235.04 MB/s | 15059620 | 41643 | 3.3× |
| LightningDecodeAny | 42875531 | 121.35 MB/s | 34333347 | 912810 | 2.6× |
| JSONV2 | 48254618 | 167.86 MB/s | 15233721 | 78972 | 2.3× |
| Stdlib | 112313892 | 72.12 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6194862 | 481.60 MB/s | 3985336 | 30015 | 9.1× |
| Lightning | 6406169 | 465.72 MB/s | 3985340 | 30015 | 8.8× |
| SonicFastest | 8972643 | 332.51 MB/s | 9131945 | 57804 | 6.3× |
| Sonic | 9022422 | 330.67 MB/s | 9131795 | 57804 | 6.3× |
| Goccy | 18658014 | 159.90 MB/s | 9878220 | 273621 | 3.0× |
| Easyjson | 19013711 | 156.91 MB/s | 9479440 | 30115 | 3.0× |
| LightningDecodeAny | 20624938 | 88.93 MB/s | 20023835 | 408557 | 2.7× |
| JSONV2 | 26603116 | 112.15 MB/s | 9257029 | 86278 | 2.1× |
| Stdlib | 56416661 | 52.88 MB/s | 9258082 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1438109 | 503.16 MB/s | 907394 | 3618 | 9.7× |
| Lightning | 1516121 | 477.27 MB/s | 907386 | 3618 | 9.2× |
| Sonic | 2130274 | 339.67 MB/s | 2373919 | 3683 | 6.6× |
| SonicFastest | 2131384 | 339.50 MB/s | 2373041 | 3683 | 6.6× |
| Easyjson | 5428580 | 133.29 MB/s | 2847906 | 3698 | 2.6× |
| LightningDecodeAny | 5562953 | 116.95 MB/s | 5752202 | 80172 | 2.5× |
| Goccy | 5584928 | 129.56 MB/s | 2728447 | 80269 | 2.5× |
| JSONV2 | 6251233 | 115.75 MB/s | 2704701 | 7318 | 2.2× |
| Stdlib | 13991377 | 51.72 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2069943 | 762.03 MB/s | 907387 | 3618 | 9.6× |
| LightningDestructive | 2085458 | 756.36 MB/s | 907392 | 3618 | 9.5× |
| SonicFastest | 2472103 | 638.06 MB/s | 3230930 | 3683 | 8.0× |
| Sonic | 2502203 | 630.39 MB/s | 3230296 | 3683 | 7.9× |
| LightningDecodeAny | 4817704 | 156.38 MB/s | 5752199 | 80172 | 4.1× |
| Easyjson | 6667541 | 236.57 MB/s | 2847905 | 3698 | 3.0× |
| Goccy | 6768576 | 233.04 MB/s | 3487008 | 80262 | 2.9× |
| JSONV2 | 6992812 | 225.57 MB/s | 2704551 | 7318 | 2.8× |
| Stdlib | 19819673 | 79.59 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 249709 | 601.20 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 259032 | 579.56 MB/s | 81920 | 1 | 8.3× |
| Sonic | 378466 | 396.66 MB/s | 407279 | 16 | 5.7× |
| SonicFastest | 402466 | 373.01 MB/s | 407359 | 16 | 5.3× |
| LightningDecodeAny | 598895 | 250.66 MB/s | 746006 | 10020 | 3.6× |
| Goccy | 1038208 | 144.60 MB/s | 329033 | 10005 | 2.1× |
| JSONV2 | 1111125 | 135.11 MB/s | 357724 | 20 | 1.9× |
| Stdlib | 2144082 | 70.02 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 36054 | 779.87 MB/s | 30272 | 105 | 9.4× |
| LightningDestructive | 37025 | 759.41 MB/s | 30144 | 103 | 9.2× |
| SonicFastest | 58571 | 480.05 MB/s | 59467 | 83 | 5.8× |
| Sonic | 58759 | 478.52 MB/s | 59496 | 83 | 5.8× |
| Easyjson | 78333 | 358.94 MB/s | 32304 | 138 | 4.3× |
| Goccy | 82264 | 341.79 MB/s | 59263 | 188 | 4.1× |
| JSONV2 | 142026 | 197.97 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 166476 | 168.90 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 340339 | 82.61 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1986 | 1172.28 MB/s | 32 | 1 | 13.1× |
| LightningDestructive | 2056 | 1132.37 MB/s | 32 | 1 | 12.7× |
| Sonic | 4677 | 497.80 MB/s | 3710 | 4 | 5.6× |
| SonicFastest | 4703 | 495.03 MB/s | 3710 | 4 | 5.5× |
| Goccy | 4853 | 479.73 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 5453 | 426.89 MB/s | 192 | 2 | 4.8× |
| JSONV2 | 8780 | 265.13 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 10620 | 158.66 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 26045 | 89.38 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 214 | 882.77 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 221 | 855.72 MB/s | 0 | 0 | 12.6× |
| Goccy | 446 | 423.71 MB/s | 304 | 2 | 6.2× |
| Easyjson | 552 | 342.30 MB/s | 0 | 0 | 5.0× |
| Sonic | 649 | 291.15 MB/s | 341 | 3 | 4.3× |
| SonicFastest | 654 | 289.08 MB/s | 341 | 3 | 4.2× |
| JSONV2 | 1072 | 176.27 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1365 | 98.20 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2778 | 68.03 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1437 | 1524.83 MB/s | 0 | 0 | 13.1× |
| LightningDestructive | 1501 | 1459.50 MB/s | 0 | 0 | 12.5× |
| Easyjson | 3528 | 621.11 MB/s | 24 | 1 | 5.3× |
| Goccy | 3866 | 566.73 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6477 | 338.28 MB/s | 3600 | 38 | 2.9× |
| Sonic | 6675 | 328.24 MB/s | 3600 | 38 | 2.8× |
| JSONV2 | 8469 | 258.71 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9236 | 196.08 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18784 | 116.64 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 697247 | 732.13 MB/s | 703779 | 1012 | 10.1× |
| Lightning | 731880 | 697.49 MB/s | 703778 | 1012 | 9.6× |
| Goccy | 1227235 | 415.96 MB/s | 1139585 | 5006 | 5.7× |
| SonicFastest | 1286452 | 396.81 MB/s | 1309812 | 2014 | 5.5× |
| Sonic | 1294194 | 394.44 MB/s | 1310242 | 2014 | 5.4× |
| Easyjson | 1749439 | 291.79 MB/s | 863779 | 3012 | 4.0× |
| JSONV2 | 3586505 | 142.33 MB/s | 1075955 | 12645 | 2.0× |
| LightningDecodeAny | 3590160 | 128.54 MB/s | 2785927 | 66022 | 2.0× |
| Stdlib | 7028343 | 72.63 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 641 | 30860.68 MB/s | 0 | 0 | 254.2× |
| LightningDestructive | 964 | 20530.99 MB/s | 0 | 0 | 169.1× |
| SonicFastest | 6767 | 2924.25 MB/s | 21121 | 3 | 24.1× |
| Goccy | 24379 | 811.72 MB/s | 20492 | 2 | 6.7× |
| Sonic | 32121 | 616.08 MB/s | 20621 | 3 | 5.1× |
| JSONV2 | 33769 | 586.01 MB/s | 8 | 1 | 4.8× |
| LightningDecodeAny | 102142 | 193.73 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 104768 | 188.88 MB/s | 0 | 0 | 1.6× |
| Stdlib | 163001 | 121.40 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2198 | 8244.41 MB/s | 432 | 2 | 54.2× |
| LightningDestructive | 2550 | 7108.46 MB/s | 0 | 0 | 46.7× |
| Easyjson | 4769 | 3800.11 MB/s | 432 | 2 | 25.0× |
| SonicFastest | 8261 | 2193.86 MB/s | 20446 | 5 | 14.4× |
| Sonic | 8308 | 2181.52 MB/s | 20430 | 5 | 14.3× |
| LightningDecodeAny | 18952 | 943.56 MB/s | 29088 | 191 | 6.3× |
| Goccy | 24666 | 734.77 MB/s | 19460 | 2 | 4.8× |
| JSONV2 | 49131 | 368.89 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 119025 | 152.27 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2296508 | 874.59 MB/s | 2960389 | 7411 | 9.6× |
| Lightning | 2463833 | 815.19 MB/s | 2962101 | 7417 | 8.9× |
| SonicFastest | 4044309 | 496.62 MB/s | 5155590 | 7085 | 5.4× |
| Sonic | 4132948 | 485.97 MB/s | 5155408 | 7085 | 5.3× |
| Goccy | 4697061 | 427.61 MB/s | 5411451 | 15832 | 4.7× |
| Easyjson | 5212917 | 385.29 MB/s | 2981488 | 7439 | 4.2× |
| LightningDecodeAny | 7093762 | 161.03 MB/s | 7386073 | 134751 | 3.1× |
| JSONV2 | 7668711 | 261.91 MB/s | 3173683 | 14563 | 2.9× |
| Stdlib | 22003220 | 91.28 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 852 | 644.29 MB/s | 480 | 1 | 8.0× |
| LightningDestructive | 859 | 638.84 MB/s | 480 | 1 | 7.9× |
| LightningDecodeAny | 2212 | 247.71 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 2262 | 242.67 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2397 | 229.02 MB/s | 2261 | 8 | 2.8× |
| Sonic | 2410 | 227.76 MB/s | 2260 | 8 | 2.8× |
| Goccy | 3401 | 161.43 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3532 | 155.42 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6789 | 80.87 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 502810 | 1255.97 MB/s | 364473 | 566 | 12.6× |
| Lightning | 569453 | 1108.98 MB/s | 413001 | 878 | 11.1× |
| SonicFastest | 1003163 | 629.52 MB/s | 1071054 | 814 | 6.3× |
| Sonic | 1003944 | 629.03 MB/s | 1071084 | 814 | 6.3× |
| Goccy | 1274395 | 495.54 MB/s | 987012 | 1200 | 5.0× |
| Easyjson | 1390575 | 454.14 MB/s | 422504 | 936 | 4.5× |
| JSONV2 | 2384579 | 264.83 MB/s | 571590 | 3144 | 2.7× |
| LightningDecodeAny | 2502526 | 186.57 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6324828 | 99.85 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 821939 | 684.25 MB/s | 544249 | 448 | 7.3× |
| Lightning | 1009295 | 557.23 MB/s | 767621 | 1254 | 5.9× |
| Sonic | 1267108 | 443.85 MB/s | 1348849 | 1185 | 4.7× |
| SonicFastest | 1302499 | 431.79 MB/s | 1350278 | 1185 | 4.6× |
| Goccy | 1479138 | 380.23 MB/s | 1040847 | 1028 | 4.0× |
| Easyjson | 2154962 | 260.98 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 3110095 | 180.83 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 3156358 | 178.18 MB/s | 927407 | 3482 | 1.9× |
| Stdlib | 5985895 | 93.96 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 607193 | 878.10 MB/s | 333416 | 2084 | 10.3× |
| Lightning | 692532 | 769.90 MB/s | 368224 | 2293 | 9.0× |
| Sonic | 1110567 | 480.10 MB/s | 982078 | 3082 | 5.6× |
| SonicFastest | 1119155 | 476.41 MB/s | 982183 | 3082 | 5.6× |
| Easyjson | 1261352 | 422.70 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1494593 | 356.74 MB/s | 1167080 | 5409 | 4.2× |
| JSONV2 | 2865319 | 186.08 MB/s | 745422 | 13288 | 2.2× |
| LightningDecodeAny | 3487585 | 152.88 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6241825 | 85.42 MB/s | 798692 | 17133 | 1.0× |
