# JSON Deserialization Benchmarks

- generated 2026-07-02T07:16:08Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104779 | 1214.70 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 104939 | 1212.85 MB/s | 49280 | 2 | 10.5× |
| Sonic | 181126 | 702.69 MB/s | 192916 | 10 | 6.1× |
| SonicFastest | 181522 | 701.16 MB/s | 193390 | 10 | 6.1× |
| Goccy | 195603 | 650.68 MB/s | 225301 | 884 | 5.7× |
| Easyjson | 212928 | 597.74 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 420936 | 302.36 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 456721 | 207.24 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1106509 | 115.02 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3690544 | 609.95 MB/s | 3122872 | 3081 | 7.1× |
| Lightning | 3757829 | 599.03 MB/s | 3122874 | 3081 | 6.9× |
| Sonic | 4563173 | 493.31 MB/s | 15233790 | 970 | 5.7× |
| SonicFastest | 4654015 | 483.68 MB/s | 15233774 | 970 | 5.6× |
| Goccy | 10363959 | 217.20 MB/s | 4120903 | 56532 | 2.5× |
| Easyjson | 10991808 | 204.79 MB/s | 3099813 | 2120 | 2.4× |
| LightningDecodeAny | 11749743 | 191.58 MB/s | 7938298 | 281383 | 2.2× |
| JSONV2 | 16230881 | 138.69 MB/s | 3123221 | 3083 | 1.6× |
| Stdlib | 26108748 | 86.22 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 487656 | 554.50 MB/s | 348024 | 1627 | 6.9× |
| LightningDestructive | 496148 | 545.01 MB/s | 348024 | 1627 | 6.8× |
| Sonic | 632368 | 427.60 MB/s | 470777 | 968 | 5.3× |
| SonicFastest | 634479 | 426.18 MB/s | 473194 | 968 | 5.3× |
| Easyjson | 1405781 | 192.35 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1409713 | 191.81 MB/s | 544075 | 8123 | 2.4× |
| LightningDecodeAny | 1597423 | 169.27 MB/s | 1011485 | 37901 | 2.1× |
| JSONV2 | 2102614 | 128.60 MB/s | 348152 | 1628 | 1.6× |
| Stdlib | 3362097 | 80.43 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1369835 | 1260.88 MB/s | 959890 | 5882 | 9.7× |
| LightningDestructive | 1377697 | 1253.69 MB/s | 959848 | 5881 | 9.6× |
| SonicFastest | 2064109 | 836.78 MB/s | 2720272 | 4020 | 6.4× |
| Sonic | 2073032 | 833.18 MB/s | 2747797 | 4020 | 6.4× |
| Goccy | 2380328 | 725.62 MB/s | 2582983 | 14604 | 5.6× |
| Easyjson | 4221781 | 409.12 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4313504 | 400.42 MB/s | 1011634 | 7594 | 3.1× |
| LightningDecodeAny | 4724052 | 105.90 MB/s | 4976205 | 81466 | 2.8× |
| Stdlib | 13244986 | 130.40 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1195 | 1516.70 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1204 | 1504.59 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2527 | 717.07 MB/s | 24 | 1 | 5.6× |
| Goccy | 2889 | 627.30 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6093 | 297.41 MB/s | 3776 | 40 | 2.3× |
| SonicFastest | 6102 | 296.95 MB/s | 3789 | 40 | 2.3× |
| JSONV2 | 7793 | 232.51 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8274 | 218.87 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14061 | 128.87 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1228 | 1475.26 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1236 | 1466.16 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2528 | 716.64 MB/s | 24 | 1 | 5.6× |
| Goccy | 2812 | 644.38 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6081 | 297.99 MB/s | 3835 | 40 | 2.3× |
| Sonic | 6133 | 295.46 MB/s | 3901 | 40 | 2.3× |
| JSONV2 | 7740 | 234.11 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8630 | 209.84 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14064 | 128.84 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1406 | 1289.06 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1458 | 1243.22 MB/s | 144 | 10 | 9.6× |
| Easyjson | 2761 | 656.21 MB/s | 144 | 10 | 5.1× |
| Goccy | 2932 | 617.98 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6151 | 294.59 MB/s | 3827 | 42 | 2.3× |
| SonicFastest | 6170 | 293.70 MB/s | 3862 | 42 | 2.3× |
| JSONV2 | 7924 | 228.67 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8538 | 212.10 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14033 | 129.12 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 693 | 712.39 MB/s | 160 | 1 | 8.0× |
| Lightning | 694 | 711.83 MB/s | 160 | 1 | 8.0× |
| SonicFastest | 1242 | 397.88 MB/s | 972 | 6 | 4.5× |
| Sonic | 1244 | 397.14 MB/s | 974 | 6 | 4.5× |
| LightningDecodeAny | 1636 | 301.36 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2265 | 218.12 MB/s | 448 | 3 | 2.5× |
| Goccy | 2416 | 204.43 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3240 | 152.48 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5557 | 88.90 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 421 | 546.47 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 421 | 546.47 MB/s | 160 | 1 | 9.9× |
| SonicFastest | 887 | 259.18 MB/s | 662 | 6 | 4.7× |
| Sonic | 889 | 258.78 MB/s | 655 | 6 | 4.7× |
| LightningDecodeAny | 1388 | 164.98 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1431 | 160.72 MB/s | 448 | 3 | 2.9× |
| Goccy | 1582 | 145.42 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2438 | 94.34 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4161 | 55.28 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 71988 | 904.76 MB/s | 166212 | 102 | 7.5× |
| Lightning | 73043 | 891.70 MB/s | 172432 | 107 | 7.4× |
| Sonic | 96796 | 672.88 MB/s | 155066 | 75 | 5.6× |
| SonicFastest | 97006 | 671.42 MB/s | 155068 | 75 | 5.6× |
| Goccy | 140723 | 462.84 MB/s | 229239 | 134 | 3.9× |
| LightningDecodeAny | 186432 | 286.05 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 221250 | 294.38 MB/s | 206650 | 607 | 2.5× |
| Stdlib | 543275 | 119.89 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2609519 | 743.61 MB/s | 2846865 | 2238 | 9.0× |
| Lightning | 2712771 | 715.31 MB/s | 2846867 | 2238 | 8.7× |
| Sonic | 4550951 | 426.39 MB/s | 14606973 | 1407 | 5.2× |
| SonicFastest | 4675060 | 415.07 MB/s | 14610265 | 1407 | 5.0× |
| Goccy | 4796573 | 404.55 MB/s | 4066063 | 13510 | 4.9× |
| Easyjson | 7542074 | 257.29 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 9932771 | 195.36 MB/s | 7013526 | 219937 | 2.4× |
| JSONV2 | 11254526 | 172.42 MB/s | 3237232 | 13947 | 2.1× |
| Stdlib | 23543726 | 82.42 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1109768 | 2998.67 MB/s | 351704 | 1286 | 18.7× |
| Lightning | 1777399 | 1872.30 MB/s | 2488905 | 2995 | 11.7× |
| Sonic | 2647071 | 1257.17 MB/s | 6453084 | 4248 | 7.8× |
| SonicFastest | 2650751 | 1255.43 MB/s | 6473316 | 4248 | 7.8× |
| LightningDecodeAny | 3769327 | 815.47 MB/s | 4886622 | 56892 | 5.5× |
| Goccy | 4599967 | 723.45 MB/s | 3948908 | 3816 | 4.5× |
| JSONV2 | 7361623 | 452.05 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20725648 | 160.57 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219191 | 1005.27 MB/s | 136896 | 228 | 9.3× |
| LightningDestructive | 221711 | 993.84 MB/s | 136897 | 228 | 9.2× |
| SonicFastest | 378435 | 582.26 MB/s | 303963 | 398 | 5.4× |
| Sonic | 393236 | 560.34 MB/s | 334081 | 398 | 5.2× |
| Goccy | 429048 | 513.57 MB/s | 364411 | 1067 | 4.8× |
| Easyjson | 550346 | 400.38 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 733207 | 300.52 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 882765 | 122.70 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2038642 | 108.08 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12721500 | 636.72 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13251918 | 611.24 MB/s | 15059843 | 51649 | 6.7× |
| Sonic | 16678410 | 485.66 MB/s | 70887494 | 40014 | 5.3× |
| SonicFastest | 17175151 | 471.61 MB/s | 70887293 | 40014 | 5.2× |
| Goccy | 23059814 | 351.26 MB/s | 16816758 | 107147 | 3.9× |
| Easyjson | 30685740 | 263.97 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 41573358 | 125.15 MB/s | 34333352 | 912810 | 2.1× |
| JSONV2 | 43713728 | 185.30 MB/s | 15233713 | 78972 | 2.0× |
| Stdlib | 89147009 | 90.86 MB/s | 15665066 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5632576 | 529.68 MB/s | 3985336 | 30015 | 8.4× |
| Lightning | 5716342 | 521.92 MB/s | 3985339 | 30015 | 8.3× |
| SonicFastest | 8658481 | 344.57 MB/s | 26558457 | 56760 | 5.5× |
| Sonic | 8670522 | 344.09 MB/s | 26570718 | 56760 | 5.5× |
| Easyjson | 16542780 | 180.35 MB/s | 9479441 | 30115 | 2.9× |
| Goccy | 16752371 | 178.09 MB/s | 10682762 | 273652 | 2.8× |
| LightningDecodeAny | 19370578 | 94.69 MB/s | 20023837 | 408557 | 2.4× |
| JSONV2 | 24628815 | 121.14 MB/s | 9257144 | 86278 | 1.9× |
| Stdlib | 47298625 | 63.08 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1225408 | 590.49 MB/s | 907393 | 3618 | 9.4× |
| Lightning | 1241447 | 582.87 MB/s | 907387 | 3618 | 9.3× |
| Sonic | 1780146 | 406.48 MB/s | 3182859 | 7226 | 6.5× |
| SonicFastest | 1785648 | 405.23 MB/s | 3192053 | 7226 | 6.5× |
| Easyjson | 4205185 | 172.07 MB/s | 2847905 | 3698 | 2.8× |
| LightningDecodeAny | 4349304 | 149.58 MB/s | 5752200 | 80172 | 2.7× |
| Goccy | 4838703 | 149.54 MB/s | 2935809 | 80281 | 2.4× |
| JSONV2 | 5956365 | 121.48 MB/s | 2704626 | 7318 | 1.9× |
| Stdlib | 11566779 | 62.56 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1834194 | 859.97 MB/s | 907393 | 3618 | 8.6× |
| Lightning | 1883738 | 837.35 MB/s | 907386 | 3618 | 8.3× |
| SonicFastest | 2233224 | 706.31 MB/s | 5785227 | 7226 | 7.0× |
| Sonic | 2236377 | 705.32 MB/s | 5785263 | 7226 | 7.0× |
| LightningDecodeAny | 3904869 | 192.94 MB/s | 5752200 | 80172 | 4.0× |
| Goccy | 5562483 | 283.57 MB/s | 3545489 | 80265 | 2.8× |
| Easyjson | 5583167 | 282.52 MB/s | 2847906 | 3698 | 2.8× |
| JSONV2 | 6557017 | 240.56 MB/s | 2704592 | 7318 | 2.4× |
| Stdlib | 15717401 | 100.36 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 209244 | 717.46 MB/s | 81920 | 1 | 8.7× |
| LightningDestructive | 209471 | 716.68 MB/s | 81920 | 1 | 8.7× |
| SonicFastest | 271447 | 553.05 MB/s | 246912 | 6 | 6.7× |
| Sonic | 271560 | 552.82 MB/s | 247288 | 6 | 6.7× |
| LightningDecodeAny | 482652 | 311.03 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 859604 | 174.64 MB/s | 326277 | 10005 | 2.1× |
| JSONV2 | 1101336 | 136.31 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1823115 | 82.34 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32789 | 857.52 MB/s | 30272 | 105 | 9.3× |
| LightningDestructive | 32836 | 856.30 MB/s | 30144 | 103 | 9.3× |
| Sonic | 63517 | 442.67 MB/s | 46697 | 103 | 4.8× |
| SonicFastest | 63788 | 440.79 MB/s | 46928 | 103 | 4.8× |
| Easyjson | 69105 | 406.88 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72123 | 389.85 MB/s | 59198 | 188 | 4.2× |
| JSONV2 | 133720 | 210.27 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 151845 | 185.17 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 304072 | 92.47 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1953 | 1192.15 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2012 | 1156.89 MB/s | 32 | 1 | 11.3× |
| Easyjson | 4212 | 552.75 MB/s | 192 | 2 | 5.4× |
| Goccy | 4249 | 547.91 MB/s | 3649 | 4 | 5.3× |
| Sonic | 5081 | 458.13 MB/s | 4317 | 6 | 4.5× |
| SonicFastest | 5107 | 455.84 MB/s | 4309 | 6 | 4.4× |
| JSONV2 | 8450 | 275.51 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10348 | 162.83 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22661 | 102.73 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221 | 856.69 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 222 | 850.20 MB/s | 0 | 0 | 10.9× |
| Goccy | 384 | 492.53 MB/s | 304 | 2 | 6.3× |
| Easyjson | 493 | 383.70 MB/s | 0 | 0 | 4.9× |
| SonicFastest | 795 | 237.77 MB/s | 493 | 4 | 3.1× |
| Sonic | 799 | 236.56 MB/s | 499 | 4 | 3.0× |
| JSONV2 | 1033 | 182.93 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1237 | 108.36 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2426 | 77.90 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1549 | 1414.67 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1565 | 1399.95 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3192 | 686.35 MB/s | 24 | 1 | 5.0× |
| Goccy | 3193 | 686.23 MB/s | 2864 | 4 | 5.0× |
| Sonic | 6572 | 333.39 MB/s | 4060 | 40 | 2.4× |
| SonicFastest | 6577 | 333.12 MB/s | 4103 | 40 | 2.4× |
| JSONV2 | 8046 | 272.31 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8580 | 211.08 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16060 | 136.43 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 704912 | 724.17 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 745697 | 684.56 MB/s | 703778 | 1012 | 8.1× |
| Sonic | 1167307 | 437.31 MB/s | 893915 | 2006 | 5.2× |
| SonicFastest | 1168303 | 436.94 MB/s | 904618 | 2006 | 5.2× |
| Goccy | 1181623 | 432.01 MB/s | 1141781 | 5007 | 5.1× |
| Easyjson | 1538658 | 331.77 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3229637 | 158.06 MB/s | 1076011 | 12646 | 1.9× |
| LightningDecodeAny | 3573684 | 129.13 MB/s | 2785929 | 66022 | 1.7× |
| Stdlib | 6072611 | 84.06 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1336 | 14813.94 MB/s | 0 | 0 | 83.5× |
| LightningDestructive | 1364 | 14510.28 MB/s | 0 | 0 | 81.8× |
| Goccy | 20901 | 946.81 MB/s | 20491 | 2 | 5.3× |
| SonicFastest | 27921 | 708.75 MB/s | 21973 | 4 | 4.0× |
| Sonic | 27949 | 708.03 MB/s | 22026 | 4 | 4.0× |
| JSONV2 | 29645 | 667.54 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 79121 | 250.10 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86014 | 230.07 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111510 | 177.46 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2672 | 6783.53 MB/s | 0 | 0 | 38.5× |
| Lightning | 2817 | 6434.16 MB/s | 432 | 2 | 36.5× |
| Easyjson | 3950 | 4588.10 MB/s | 432 | 2 | 26.1× |
| Sonic | 9869 | 1836.55 MB/s | 22756 | 6 | 10.4× |
| SonicFastest | 10331 | 1754.29 MB/s | 23674 | 6 | 10.0× |
| Goccy | 16391 | 1105.72 MB/s | 19459 | 2 | 6.3× |
| LightningDecodeAny | 16929 | 1056.29 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 45576 | 397.67 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102951 | 176.05 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2360740 | 850.79 MB/s | 2951717 | 7277 | 7.9× |
| Lightning | 2423059 | 828.91 MB/s | 2953429 | 7283 | 7.7× |
| Goccy | 4275457 | 469.77 MB/s | 5411356 | 15830 | 4.4× |
| Sonic | 4410048 | 455.44 MB/s | 10902071 | 13683 | 4.2× |
| SonicFastest | 4431842 | 453.20 MB/s | 10925510 | 13683 | 4.2× |
| Easyjson | 4906749 | 409.33 MB/s | 2981483 | 7439 | 3.8× |
| JSONV2 | 6839165 | 293.68 MB/s | 3173684 | 14563 | 2.7× |
| LightningDecodeAny | 7399447 | 154.38 MB/s | 7386073 | 134751 | 2.5× |
| Stdlib | 18639910 | 107.75 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 891 | 616.38 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 896 | 612.44 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1921 | 285.23 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2170 | 253.05 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2680 | 204.83 MB/s | 1944 | 26 | 2.1× |
| SonicFastest | 2702 | 203.19 MB/s | 1966 | 26 | 2.1× |
| Goccy | 3039 | 180.62 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3302 | 166.25 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5640 | 97.35 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 493933 | 1278.54 MB/s | 364472 | 566 | 11.0× |
| Lightning | 552126 | 1143.79 MB/s | 413001 | 878 | 9.8× |
| SonicFastest | 1009786 | 625.39 MB/s | 1003598 | 1102 | 5.4× |
| Sonic | 1010620 | 624.88 MB/s | 1005066 | 1102 | 5.4× |
| Easyjson | 1146033 | 551.04 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1162658 | 543.16 MB/s | 986049 | 1201 | 4.7× |
| JSONV2 | 2134787 | 295.82 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2389036 | 195.44 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5416598 | 116.59 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 711129 | 790.87 MB/s | 544253 | 448 | 7.4× |
| Lightning | 886627 | 634.32 MB/s | 767621 | 1254 | 5.9× |
| SonicFastest | 1043217 | 539.11 MB/s | 978892 | 1476 | 5.0× |
| Sonic | 1051499 | 534.86 MB/s | 989175 | 1476 | 5.0× |
| Goccy | 1333242 | 421.83 MB/s | 1041536 | 1030 | 3.9× |
| Easyjson | 1744818 | 322.33 MB/s | 775154 | 1254 | 3.0× |
| JSONV2 | 2748855 | 204.60 MB/s | 927442 | 3482 | 1.9× |
| LightningDecodeAny | 2762300 | 203.60 MB/s | 2114151 | 30295 | 1.9× |
| Stdlib | 5265209 | 106.82 MB/s | 1011671 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 662583 | 804.70 MB/s | 333416 | 2084 | 8.3× |
| Lightning | 700949 | 760.65 MB/s | 368224 | 2293 | 7.9× |
| Easyjson | 1136066 | 469.32 MB/s | 428362 | 3273 | 4.9× |
| Sonic | 1148667 | 464.17 MB/s | 1036770 | 4351 | 4.8× |
| SonicFastest | 1149442 | 463.86 MB/s | 1035092 | 4351 | 4.8× |
| Goccy | 1309613 | 407.13 MB/s | 1167244 | 5409 | 4.2× |
| JSONV2 | 2566735 | 207.73 MB/s | 745471 | 13288 | 2.1× |
| LightningDecodeAny | 3459844 | 154.10 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5513848 | 96.70 MB/s | 798692 | 17133 | 1.0× |
