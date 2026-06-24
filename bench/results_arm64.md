# JSON Deserialization Benchmarks

- generated 2026-06-24T11:19:49Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104955 | 1212.67 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 104981 | 1212.36 MB/s | 49280 | 2 | 10.5× |
| Sonic | 183152 | 694.91 MB/s | 192112 | 10 | 6.0× |
| SonicFastest | 185071 | 687.71 MB/s | 198400 | 10 | 6.0× |
| Goccy | 192705 | 660.47 MB/s | 224579 | 884 | 5.7× |
| Easyjson | 212910 | 597.79 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 419967 | 303.06 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 459078 | 206.18 MB/s | 465585 | 9714 | 2.4× |
| Stdlib | 1107472 | 114.92 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3764078 | 598.04 MB/s | 3122872 | 3081 | 7.0× |
| Lightning | 3861012 | 583.02 MB/s | 3122876 | 3081 | 6.8× |
| Sonic | 4602965 | 489.04 MB/s | 15239038 | 970 | 5.7× |
| SonicFastest | 4648654 | 484.24 MB/s | 15233781 | 970 | 5.6× |
| Goccy | 10386299 | 216.73 MB/s | 4118255 | 56532 | 2.5× |
| Easyjson | 11192712 | 201.12 MB/s | 3099808 | 2120 | 2.3× |
| LightningDecodeAny | 11602356 | 194.01 MB/s | 7938299 | 281383 | 2.3× |
| JSONV2 | 16408832 | 137.19 MB/s | 3123216 | 3083 | 1.6× |
| Stdlib | 26197500 | 85.93 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 494526 | 546.79 MB/s | 348025 | 1627 | 6.8× |
| LightningDestructive | 499431 | 541.42 MB/s | 348024 | 1627 | 6.7× |
| Sonic | 642763 | 420.69 MB/s | 475155 | 968 | 5.2× |
| SonicFastest | 655028 | 412.81 MB/s | 500321 | 968 | 5.1× |
| Goccy | 1429708 | 189.13 MB/s | 543710 | 8122 | 2.4× |
| Easyjson | 1452024 | 186.22 MB/s | 330272 | 749 | 2.3× |
| LightningDecodeAny | 1652034 | 163.68 MB/s | 1011484 | 37901 | 2.0× |
| JSONV2 | 2125250 | 127.23 MB/s | 348152 | 1628 | 1.6× |
| Stdlib | 3361776 | 80.43 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1374355 | 1256.74 MB/s | 959848 | 5881 | 9.6× |
| Lightning | 1398404 | 1235.13 MB/s | 959890 | 5882 | 9.5× |
| Sonic | 2088608 | 826.96 MB/s | 2744021 | 4020 | 6.3× |
| SonicFastest | 2099898 | 822.52 MB/s | 2724998 | 4020 | 6.3× |
| Goccy | 2406803 | 717.63 MB/s | 2583325 | 14605 | 5.5× |
| Easyjson | 4231221 | 408.20 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4354269 | 396.67 MB/s | 1011635 | 7594 | 3.0× |
| LightningDecodeAny | 4767356 | 104.94 MB/s | 4976204 | 81466 | 2.8× |
| Stdlib | 13255690 | 130.30 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1189 | 1523.93 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1205 | 1503.90 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2555 | 709.19 MB/s | 24 | 1 | 5.5× |
| Goccy | 2828 | 640.83 MB/s | 2608 | 4 | 5.0× |
| Sonic | 6032 | 300.41 MB/s | 3765 | 40 | 2.3× |
| SonicFastest | 6042 | 299.91 MB/s | 3755 | 40 | 2.3× |
| JSONV2 | 7839 | 231.15 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8342 | 217.10 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14017 | 129.27 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1220 | 1485.80 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1235 | 1466.78 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2550 | 710.49 MB/s | 24 | 1 | 5.5× |
| Goccy | 2830 | 640.27 MB/s | 2608 | 4 | 5.0× |
| Sonic | 6003 | 301.86 MB/s | 3798 | 40 | 2.3× |
| SonicFastest | 6009 | 301.53 MB/s | 3782 | 40 | 2.3× |
| JSONV2 | 7808 | 232.07 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8354 | 216.79 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14090 | 128.61 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1391 | 1303.00 MB/s | 144 | 10 | 10.1× |
| LightningDestructive | 1441 | 1257.85 MB/s | 144 | 10 | 9.8× |
| Easyjson | 2762 | 656.04 MB/s | 144 | 10 | 5.1× |
| Goccy | 2905 | 623.65 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6169 | 293.74 MB/s | 3796 | 42 | 2.3× |
| SonicFastest | 6184 | 293.01 MB/s | 3789 | 42 | 2.3× |
| JSONV2 | 8015 | 226.08 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8319 | 217.69 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14085 | 128.65 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 746 | 662.66 MB/s | 160 | 1 | 7.4× |
| Lightning | 748 | 660.56 MB/s | 160 | 1 | 7.4× |
| Sonic | 1248 | 395.73 MB/s | 984 | 6 | 4.4× |
| SonicFastest | 1248 | 395.75 MB/s | 982 | 6 | 4.4× |
| LightningDecodeAny | 1681 | 293.29 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2200 | 224.56 MB/s | 448 | 3 | 2.5× |
| Goccy | 2478 | 199.38 MB/s | 856 | 23 | 2.2× |
| JSONV2 | 3180 | 155.35 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5528 | 89.36 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 471 | 488.12 MB/s | 160 | 1 | 8.8× |
| LightningDestructive | 474 | 485.38 MB/s | 160 | 1 | 8.7× |
| Sonic | 892 | 257.86 MB/s | 659 | 6 | 4.6× |
| SonicFastest | 893 | 257.44 MB/s | 663 | 6 | 4.6× |
| Easyjson | 1386 | 165.94 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1408 | 162.64 MB/s | 1536 | 30 | 2.9× |
| Goccy | 1625 | 141.58 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2341 | 98.23 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4126 | 55.74 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 75117 | 867.07 MB/s | 166212 | 102 | 7.3× |
| Lightning | 76266 | 854.02 MB/s | 172432 | 107 | 7.2× |
| SonicFastest | 99509 | 654.54 MB/s | 155672 | 75 | 5.5× |
| Sonic | 99589 | 654.01 MB/s | 156379 | 75 | 5.5× |
| Goccy | 153164 | 425.24 MB/s | 229074 | 134 | 3.6× |
| LightningDecodeAny | 191909 | 277.89 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 230796 | 282.21 MB/s | 206652 | 607 | 2.4× |
| Stdlib | 548604 | 118.72 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2661022 | 729.22 MB/s | 2846865 | 2238 | 8.8× |
| Lightning | 2727057 | 711.56 MB/s | 2846866 | 2238 | 8.6× |
| SonicFastest | 4683955 | 414.28 MB/s | 14608664 | 1407 | 5.0× |
| Goccy | 4795459 | 404.65 MB/s | 4065305 | 13510 | 4.9× |
| Sonic | 4801247 | 404.16 MB/s | 14608640 | 1407 | 4.9× |
| Easyjson | 7639429 | 254.01 MB/s | 3871264 | 15043 | 3.1× |
| LightningDecodeAny | 9588729 | 202.37 MB/s | 7013528 | 219937 | 2.5× |
| JSONV2 | 11394965 | 170.29 MB/s | 3237216 | 13947 | 2.1× |
| Stdlib | 23517028 | 82.51 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1113630 | 2988.27 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1855733 | 1793.27 MB/s | 2488905 | 2995 | 11.3× |
| Sonic | 2737694 | 1215.56 MB/s | 6419298 | 4248 | 7.6× |
| SonicFastest | 2744210 | 1212.67 MB/s | 6394779 | 4248 | 7.6× |
| LightningDecodeAny | 3848232 | 798.75 MB/s | 4886623 | 56892 | 5.4× |
| Goccy | 4506903 | 738.39 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7458608 | 446.17 MB/s | 5364516 | 13243 | 2.8× |
| Stdlib | 20918416 | 159.09 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225731 | 976.14 MB/s | 136896 | 228 | 9.1× |
| LightningDestructive | 227343 | 969.22 MB/s | 136896 | 228 | 9.0× |
| Sonic | 390938 | 563.63 MB/s | 311983 | 398 | 5.2× |
| SonicFastest | 391072 | 563.44 MB/s | 316312 | 398 | 5.2× |
| Goccy | 442752 | 497.67 MB/s | 365158 | 1067 | 4.6× |
| Easyjson | 547472 | 402.48 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 738419 | 298.40 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 927509 | 116.78 MB/s | 861394 | 11944 | 2.2× |
| Stdlib | 2043735 | 107.82 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12869227 | 629.41 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13316075 | 608.29 MB/s | 15059841 | 51649 | 6.7× |
| SonicFastest | 16677988 | 485.67 MB/s | 70873024 | 40014 | 5.4× |
| Sonic | 16772819 | 482.93 MB/s | 70916128 | 40014 | 5.3× |
| Goccy | 23505786 | 344.60 MB/s | 16969747 | 107148 | 3.8× |
| Easyjson | 31297100 | 258.81 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 40797197 | 127.53 MB/s | 34333369 | 912810 | 2.2× |
| JSONV2 | 43937663 | 184.35 MB/s | 15233727 | 78972 | 2.0× |
| Stdlib | 89651845 | 90.35 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6103170 | 488.84 MB/s | 3985336 | 30015 | 7.8× |
| Lightning | 6245093 | 477.73 MB/s | 3985339 | 30015 | 7.6× |
| SonicFastest | 8682737 | 343.61 MB/s | 26566210 | 56760 | 5.5× |
| Sonic | 8710050 | 342.53 MB/s | 26641765 | 56760 | 5.4× |
| Goccy | 16655134 | 179.13 MB/s | 10631320 | 273651 | 2.8× |
| Easyjson | 16726470 | 178.37 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 18802782 | 97.55 MB/s | 20023837 | 408557 | 2.5× |
| JSONV2 | 24956948 | 119.54 MB/s | 9257169 | 86278 | 1.9× |
| Stdlib | 47330910 | 63.03 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1396007 | 518.33 MB/s | 907387 | 3618 | 8.3× |
| LightningDestructive | 1400091 | 516.82 MB/s | 907393 | 3618 | 8.3× |
| SonicFastest | 1804206 | 401.06 MB/s | 3176902 | 7226 | 6.4× |
| Sonic | 1809075 | 399.98 MB/s | 3164924 | 7226 | 6.4× |
| Easyjson | 4286586 | 168.80 MB/s | 2847907 | 3698 | 2.7× |
| LightningDecodeAny | 4463457 | 145.76 MB/s | 5752202 | 80172 | 2.6× |
| Goccy | 4724914 | 153.15 MB/s | 2797663 | 80273 | 2.5× |
| JSONV2 | 6045334 | 119.70 MB/s | 2704628 | 7318 | 1.9× |
| Stdlib | 11601334 | 62.37 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1948265 | 809.62 MB/s | 907393 | 3618 | 8.1× |
| Lightning | 1984125 | 794.99 MB/s | 907386 | 3618 | 7.9× |
| Sonic | 2260264 | 697.86 MB/s | 5784628 | 7226 | 7.0× |
| SonicFastest | 2264448 | 696.57 MB/s | 5782419 | 7226 | 6.9× |
| LightningDecodeAny | 3937613 | 191.33 MB/s | 5752200 | 80172 | 4.0× |
| Easyjson | 5621956 | 280.57 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5639195 | 279.71 MB/s | 3578551 | 80267 | 2.8× |
| JSONV2 | 6812415 | 231.54 MB/s | 2704592 | 7318 | 2.3× |
| Stdlib | 15735880 | 100.24 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224020 | 670.14 MB/s | 81920 | 1 | 8.1× |
| LightningDestructive | 225168 | 666.72 MB/s | 81920 | 1 | 8.1× |
| Sonic | 276069 | 543.79 MB/s | 254097 | 6 | 6.6× |
| SonicFastest | 276376 | 543.19 MB/s | 254605 | 6 | 6.6× |
| LightningDecodeAny | 479181 | 313.29 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 875997 | 171.37 MB/s | 331480 | 10005 | 2.1× |
| JSONV2 | 1116718 | 134.43 MB/s | 357715 | 20 | 1.6× |
| Stdlib | 1824878 | 82.27 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33252 | 845.58 MB/s | 30272 | 105 | 9.1× |
| LightningDestructive | 33520 | 838.81 MB/s | 30144 | 103 | 9.1× |
| SonicFastest | 63882 | 440.14 MB/s | 46675 | 103 | 4.8× |
| Sonic | 63936 | 439.77 MB/s | 46561 | 103 | 4.8× |
| Easyjson | 68666 | 409.48 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72423 | 388.24 MB/s | 59199 | 188 | 4.2× |
| JSONV2 | 135529 | 207.46 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 155195 | 181.17 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303940 | 92.51 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1951 | 1193.11 MB/s | 32 | 1 | 11.7× |
| LightningDestructive | 2006 | 1160.37 MB/s | 32 | 1 | 11.3× |
| Goccy | 4223 | 551.23 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4233 | 550.03 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5181 | 449.32 MB/s | 4340 | 6 | 4.4× |
| Sonic | 5196 | 448.07 MB/s | 4348 | 6 | 4.4× |
| JSONV2 | 8486 | 274.32 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10390 | 162.17 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22752 | 102.32 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 863.11 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 221 | 853.87 MB/s | 0 | 0 | 11.0× |
| Goccy | 392 | 482.74 MB/s | 304 | 2 | 6.2× |
| Easyjson | 482 | 392.49 MB/s | 0 | 0 | 5.1× |
| Sonic | 808 | 233.87 MB/s | 516 | 4 | 3.0× |
| SonicFastest | 812 | 232.68 MB/s | 520 | 4 | 3.0× |
| JSONV2 | 1025 | 184.39 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1310 | 102.32 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2435 | 77.61 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1546 | 1417.10 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1560 | 1404.48 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3182 | 688.65 MB/s | 24 | 1 | 5.0× |
| Goccy | 3232 | 677.88 MB/s | 2864 | 4 | 5.0× |
| SonicFastest | 6397 | 342.50 MB/s | 3945 | 40 | 2.5× |
| Sonic | 6406 | 342.03 MB/s | 3978 | 40 | 2.5× |
| JSONV2 | 8073 | 271.41 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8387 | 215.93 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16018 | 136.78 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 715337 | 713.62 MB/s | 703778 | 1012 | 8.4× |
| Lightning | 743075 | 686.98 MB/s | 703779 | 1012 | 8.1× |
| Goccy | 1173856 | 434.87 MB/s | 1136056 | 5006 | 5.1× |
| SonicFastest | 1176141 | 434.03 MB/s | 903780 | 2006 | 5.1× |
| Sonic | 1179068 | 432.95 MB/s | 903570 | 2006 | 5.1× |
| Easyjson | 1564122 | 326.37 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3245538 | 157.29 MB/s | 1076020 | 12646 | 1.9× |
| LightningDecodeAny | 3636346 | 126.90 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6026270 | 84.71 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14643.34 MB/s | 0 | 0 | 82.6× |
| LightningDestructive | 1402 | 14114.33 MB/s | 0 | 0 | 79.6× |
| Goccy | 20197 | 979.80 MB/s | 20491 | 2 | 5.5× |
| Sonic | 27562 | 717.99 MB/s | 22205 | 4 | 4.0× |
| SonicFastest | 27625 | 716.35 MB/s | 22711 | 4 | 4.0× |
| JSONV2 | 29771 | 664.71 MB/s | 8 | 1 | 3.7× |
| LightningDecodeAny | 82657 | 239.40 MB/s | 117104 | 2019 | 1.3× |
| Easyjson | 86019 | 230.05 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111549 | 177.40 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2655 | 6827.19 MB/s | 0 | 0 | 38.7× |
| Lightning | 2821 | 6423.85 MB/s | 432 | 2 | 36.4× |
| Easyjson | 3953 | 4584.58 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 10184 | 1779.64 MB/s | 22835 | 6 | 10.1× |
| Sonic | 10315 | 1757.06 MB/s | 23248 | 6 | 10.0× |
| Goccy | 15827 | 1145.14 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 17373 | 1029.30 MB/s | 29088 | 191 | 5.9× |
| JSONV2 | 45828 | 395.48 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102782 | 176.33 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2302705 | 872.23 MB/s | 2960388 | 7411 | 8.2× |
| Lightning | 2360790 | 850.77 MB/s | 2962101 | 7417 | 8.0× |
| Goccy | 4445506 | 451.80 MB/s | 5411862 | 15830 | 4.2× |
| Sonic | 4639908 | 432.87 MB/s | 10900576 | 13683 | 4.0× |
| SonicFastest | 4651118 | 431.83 MB/s | 10877051 | 13683 | 4.0× |
| Easyjson | 4984878 | 402.92 MB/s | 2981478 | 7438 | 3.8× |
| JSONV2 | 7056151 | 284.64 MB/s | 3173680 | 14563 | 2.7× |
| LightningDecodeAny | 7229320 | 158.01 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18788822 | 106.90 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 900 | 610.01 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 915 | 599.92 MB/s | 480 | 1 | 6.2× |
| LightningDecodeAny | 1942 | 282.23 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2188 | 250.92 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2729 | 201.16 MB/s | 1947 | 26 | 2.1× |
| Sonic | 2734 | 200.84 MB/s | 1957 | 26 | 2.1× |
| Goccy | 3077 | 178.43 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3334 | 164.67 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5685 | 96.57 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 501422 | 1259.45 MB/s | 364472 | 566 | 10.8× |
| Lightning | 560144 | 1127.41 MB/s | 413001 | 878 | 9.7× |
| Sonic | 1027817 | 614.42 MB/s | 993314 | 1102 | 5.3× |
| SonicFastest | 1036426 | 609.32 MB/s | 1005486 | 1102 | 5.2× |
| Easyjson | 1148142 | 550.03 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1186655 | 532.18 MB/s | 985121 | 1201 | 4.6× |
| JSONV2 | 2163523 | 291.89 MB/s | 571617 | 3144 | 2.5× |
| LightningDecodeAny | 2453037 | 190.34 MB/s | 2010200 | 30295 | 2.2× |
| Stdlib | 5414633 | 116.63 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 721737 | 779.24 MB/s | 544252 | 448 | 7.3× |
| Lightning | 901733 | 623.70 MB/s | 767620 | 1254 | 5.9× |
| SonicFastest | 1053859 | 533.67 MB/s | 945315 | 1476 | 5.0× |
| Sonic | 1060996 | 530.08 MB/s | 952659 | 1476 | 5.0× |
| Goccy | 1364747 | 412.10 MB/s | 1037052 | 1029 | 3.9× |
| Easyjson | 1766919 | 318.30 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2796832 | 201.09 MB/s | 2114151 | 30295 | 1.9× |
| JSONV2 | 2802942 | 200.65 MB/s | 927441 | 3482 | 1.9× |
| Stdlib | 5296355 | 106.19 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 664927 | 801.86 MB/s | 333416 | 2084 | 8.2× |
| Lightning | 689701 | 773.06 MB/s | 368224 | 2293 | 8.0× |
| Easyjson | 1122731 | 474.89 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1167702 | 456.60 MB/s | 1037609 | 4351 | 4.7× |
| SonicFastest | 1167943 | 456.51 MB/s | 1034881 | 4351 | 4.7× |
| Goccy | 1318606 | 404.35 MB/s | 1167239 | 5409 | 4.2× |
| JSONV2 | 2549936 | 209.09 MB/s | 745451 | 13288 | 2.2× |
| LightningDecodeAny | 3511017 | 151.86 MB/s | 2674619 | 50596 | 1.6× |
| Stdlib | 5483324 | 97.24 MB/s | 798692 | 17133 | 1.0× |
