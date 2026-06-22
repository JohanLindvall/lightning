# JSON Deserialization Benchmarks

- generated 2026-06-22T19:49:52Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 104370 | 1219.46 MB/s | 49280 | 2 | 10.6× |
| Lightning | 104538 | 1217.49 MB/s | 49760 | 3 | 10.6× |
| SonicFastest | 181474 | 701.34 MB/s | 193430 | 10 | 6.1× |
| Sonic | 182099 | 698.93 MB/s | 194667 | 10 | 6.1× |
| Goccy | 193371 | 658.19 MB/s | 224652 | 884 | 5.7× |
| Easyjson | 212221 | 599.73 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 418684 | 303.99 MB/s | 195120 | 1805 | 2.6× |
| LightningDecodeAny | 444400 | 212.99 MB/s | 465586 | 9714 | 2.5× |
| Stdlib | 1103179 | 115.37 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3782567 | 595.11 MB/s | 3122876 | 3081 | 6.9× |
| LightningDestructive | 3791716 | 593.68 MB/s | 3122872 | 3081 | 6.9× |
| Sonic | 4525323 | 497.43 MB/s | 15233794 | 970 | 5.8× |
| SonicFastest | 4527070 | 497.24 MB/s | 15232101 | 970 | 5.8× |
| Goccy | 10223735 | 220.18 MB/s | 4112174 | 56532 | 2.6× |
| Easyjson | 11072036 | 203.31 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11329938 | 198.68 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16088526 | 139.92 MB/s | 3123230 | 3083 | 1.6× |
| Stdlib | 26095962 | 86.26 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 489675 | 552.21 MB/s | 348024 | 1627 | 6.9× |
| LightningDestructive | 494281 | 547.06 MB/s | 348024 | 1627 | 6.8× |
| Sonic | 629844 | 429.32 MB/s | 472611 | 968 | 5.3× |
| SonicFastest | 633740 | 426.68 MB/s | 481018 | 968 | 5.3× |
| Goccy | 1395222 | 193.81 MB/s | 543325 | 8122 | 2.4× |
| Easyjson | 1407581 | 192.10 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1586862 | 170.40 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2106201 | 128.38 MB/s | 348150 | 1628 | 1.6× |
| Stdlib | 3355240 | 80.59 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1349994 | 1279.42 MB/s | 959848 | 5881 | 9.8× |
| Lightning | 1363443 | 1266.80 MB/s | 959890 | 5882 | 9.7× |
| Sonic | 2036853 | 847.98 MB/s | 2768530 | 4020 | 6.5× |
| SonicFastest | 2041373 | 846.10 MB/s | 2731142 | 4020 | 6.5× |
| Goccy | 2391637 | 722.18 MB/s | 2582507 | 14604 | 5.5× |
| Easyjson | 4213834 | 409.89 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4246287 | 406.76 MB/s | 1011634 | 7594 | 3.1× |
| LightningDecodeAny | 4610014 | 108.52 MB/s | 4976205 | 81466 | 2.9× |
| Stdlib | 13230994 | 130.54 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1193 | 1518.39 MB/s | 0 | 0 | 11.7× |
| LightningDestructive | 1199 | 1511.58 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2547 | 711.47 MB/s | 24 | 1 | 5.5× |
| Goccy | 2820 | 642.52 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5949 | 304.56 MB/s | 3744 | 40 | 2.4× |
| SonicFastest | 5982 | 302.91 MB/s | 3777 | 40 | 2.3× |
| JSONV2 | 7772 | 233.15 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8139 | 222.50 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14005 | 129.39 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1219 | 1487.06 MB/s | 0 | 0 | 11.6× |
| LightningDestructive | 1240 | 1461.01 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2540 | 713.35 MB/s | 24 | 1 | 5.5× |
| Goccy | 2869 | 631.55 MB/s | 2608 | 4 | 4.9× |
| Sonic | 5907 | 306.77 MB/s | 3693 | 40 | 2.4× |
| SonicFastest | 5932 | 305.44 MB/s | 3719 | 40 | 2.4× |
| JSONV2 | 7772 | 233.14 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8178 | 221.46 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14085 | 128.65 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1397 | 1296.61 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1447 | 1252.33 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2766 | 655.21 MB/s | 144 | 10 | 5.1× |
| Goccy | 2869 | 631.66 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6051 | 299.44 MB/s | 3768 | 42 | 2.3× |
| Sonic | 6078 | 298.11 MB/s | 3766 | 42 | 2.3× |
| JSONV2 | 7978 | 227.12 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8144 | 222.36 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14024 | 129.20 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 735 | 672.04 MB/s | 160 | 1 | 7.5× |
| Lightning | 740 | 668.00 MB/s | 160 | 1 | 7.5× |
| Sonic | 1228 | 402.32 MB/s | 977 | 6 | 4.5× |
| SonicFastest | 1231 | 401.23 MB/s | 985 | 6 | 4.5× |
| LightningDecodeAny | 1637 | 301.09 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2232 | 221.33 MB/s | 448 | 3 | 2.5× |
| Goccy | 2428 | 203.44 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3210 | 153.91 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5541 | 89.15 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 465 | 494.90 MB/s | 160 | 1 | 8.8× |
| Lightning | 466 | 493.24 MB/s | 160 | 1 | 8.8× |
| SonicFastest | 874 | 263.17 MB/s | 666 | 6 | 4.7× |
| Sonic | 874 | 263.10 MB/s | 654 | 6 | 4.7× |
| LightningDecodeAny | 1367 | 167.50 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1404 | 163.82 MB/s | 448 | 3 | 2.9× |
| Goccy | 1576 | 145.97 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2362 | 97.39 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4111 | 55.95 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 73204 | 889.73 MB/s | 166212 | 102 | 7.4× |
| Lightning | 73713 | 883.59 MB/s | 172432 | 107 | 7.4× |
| SonicFastest | 96766 | 673.09 MB/s | 154738 | 75 | 5.6× |
| Sonic | 97600 | 667.34 MB/s | 155415 | 75 | 5.6× |
| Goccy | 142147 | 458.20 MB/s | 229169 | 134 | 3.8× |
| LightningDecodeAny | 184730 | 288.69 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 221722 | 293.76 MB/s | 206651 | 607 | 2.5× |
| Stdlib | 543870 | 119.76 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2597420 | 747.08 MB/s | 2846864 | 2238 | 9.0× |
| Lightning | 2698182 | 719.18 MB/s | 2846867 | 2238 | 8.7× |
| SonicFastest | 4751277 | 408.41 MB/s | 14606972 | 1407 | 4.9× |
| Sonic | 4753343 | 408.23 MB/s | 14608641 | 1407 | 4.9× |
| Goccy | 4758896 | 407.76 MB/s | 4065964 | 13510 | 4.9× |
| Easyjson | 7496576 | 258.85 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 9758652 | 198.85 MB/s | 7013527 | 219937 | 2.4× |
| JSONV2 | 11303613 | 171.67 MB/s | 3237224 | 13947 | 2.1× |
| Stdlib | 23477300 | 82.65 MB/s | 3551325 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1105773 | 3009.51 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1758214 | 1892.73 MB/s | 2488905 | 2995 | 11.8× |
| Sonic | 2627621 | 1266.48 MB/s | 6411795 | 4248 | 7.9× |
| SonicFastest | 2662948 | 1249.68 MB/s | 6419219 | 4248 | 7.8× |
| LightningDecodeAny | 3727963 | 824.52 MB/s | 4886620 | 56892 | 5.6× |
| Goccy | 4638901 | 717.37 MB/s | 3948909 | 3816 | 4.5× |
| JSONV2 | 7399511 | 449.74 MB/s | 5364517 | 13243 | 2.8× |
| Stdlib | 20783958 | 160.12 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221513 | 994.73 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 222690 | 989.48 MB/s | 136897 | 228 | 9.2× |
| Sonic | 375497 | 586.81 MB/s | 299266 | 398 | 5.4× |
| SonicFastest | 395870 | 556.61 MB/s | 340202 | 398 | 5.2× |
| Goccy | 435106 | 506.42 MB/s | 365339 | 1067 | 4.7× |
| Easyjson | 546401 | 403.27 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 738679 | 298.30 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 878996 | 123.22 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2040972 | 107.96 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12757941 | 634.90 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13168830 | 615.09 MB/s | 15059835 | 51649 | 6.8× |
| SonicFastest | 16398163 | 493.96 MB/s | 70887434 | 40014 | 5.4× |
| Sonic | 16572762 | 488.76 MB/s | 70928765 | 40014 | 5.4× |
| Goccy | 23373171 | 346.55 MB/s | 16874129 | 107148 | 3.8× |
| Easyjson | 30856416 | 262.51 MB/s | 15059620 | 41643 | 2.9× |
| LightningDecodeAny | 40261918 | 129.23 MB/s | 34333351 | 912810 | 2.2× |
| JSONV2 | 44321092 | 182.76 MB/s | 15233736 | 78972 | 2.0× |
| Stdlib | 89354576 | 90.65 MB/s | 15665071 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6080489 | 490.66 MB/s | 3985336 | 30015 | 7.8× |
| Lightning | 6177400 | 482.96 MB/s | 3985336 | 30015 | 7.6× |
| Sonic | 8529655 | 349.78 MB/s | 26496501 | 56760 | 5.5× |
| SonicFastest | 8574331 | 347.95 MB/s | 26559768 | 56760 | 5.5× |
| Easyjson | 16403486 | 181.88 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16481048 | 181.02 MB/s | 10595349 | 273648 | 2.9× |
| LightningDecodeAny | 18417484 | 99.59 MB/s | 20023838 | 408557 | 2.6× |
| JSONV2 | 24728711 | 120.65 MB/s | 9257155 | 86278 | 1.9× |
| Stdlib | 47171950 | 63.25 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1368738 | 528.66 MB/s | 907393 | 3618 | 8.5× |
| Lightning | 1385394 | 522.30 MB/s | 907387 | 3618 | 8.4× |
| Sonic | 1767724 | 409.34 MB/s | 3183395 | 7226 | 6.5× |
| SonicFastest | 1779265 | 406.68 MB/s | 3187797 | 7226 | 6.5× |
| Easyjson | 4236440 | 170.80 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4338865 | 149.94 MB/s | 5752203 | 80172 | 2.7× |
| Goccy | 4810247 | 150.43 MB/s | 2810198 | 80274 | 2.4× |
| JSONV2 | 5729804 | 126.29 MB/s | 2704643 | 7318 | 2.0× |
| Stdlib | 11572027 | 62.53 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1945308 | 810.85 MB/s | 907393 | 3618 | 8.1× |
| Lightning | 1976076 | 798.22 MB/s | 907386 | 3618 | 8.0× |
| SonicFastest | 2229520 | 707.49 MB/s | 5789442 | 7226 | 7.1× |
| Sonic | 2247343 | 701.87 MB/s | 5785223 | 7226 | 7.0× |
| LightningDecodeAny | 3858244 | 195.27 MB/s | 5752202 | 80172 | 4.1× |
| Goccy | 5568485 | 283.26 MB/s | 3553172 | 80266 | 2.8× |
| Easyjson | 5595154 | 281.91 MB/s | 2847906 | 3698 | 2.8× |
| JSONV2 | 6383876 | 247.08 MB/s | 2704589 | 7318 | 2.5× |
| Stdlib | 15723079 | 100.32 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 223913 | 670.46 MB/s | 81920 | 1 | 8.2× |
| Lightning | 224317 | 669.25 MB/s | 81920 | 1 | 8.1× |
| SonicFastest | 268620 | 558.87 MB/s | 248307 | 6 | 6.8× |
| Sonic | 270558 | 554.87 MB/s | 254905 | 6 | 6.7× |
| LightningDecodeAny | 472514 | 317.71 MB/s | 746005 | 10020 | 3.9× |
| Goccy | 867783 | 173.00 MB/s | 329984 | 10005 | 2.1× |
| JSONV2 | 1100853 | 136.37 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1825235 | 82.25 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 33236 | 845.97 MB/s | 30144 | 103 | 9.1× |
| Lightning | 33245 | 845.76 MB/s | 30272 | 105 | 9.1× |
| SonicFastest | 62918 | 446.89 MB/s | 46351 | 103 | 4.8× |
| Sonic | 63043 | 446.00 MB/s | 46580 | 103 | 4.8× |
| Easyjson | 67625 | 415.78 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71226 | 394.76 MB/s | 59196 | 188 | 4.3× |
| JSONV2 | 133442 | 210.71 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 148882 | 188.85 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303432 | 92.66 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1954 | 1191.28 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2012 | 1157.13 MB/s | 32 | 1 | 11.3× |
| Goccy | 4128 | 563.98 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4202 | 553.97 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5068 | 459.35 MB/s | 4214 | 6 | 4.5× |
| Sonic | 5073 | 458.88 MB/s | 4223 | 6 | 4.5× |
| JSONV2 | 8454 | 275.36 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10022 | 168.13 MB/s | 9960 | 195 | 2.3× |
| Stdlib | 22657 | 102.75 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220 | 857.49 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 222 | 850.63 MB/s | 0 | 0 | 10.9× |
| Goccy | 382 | 494.87 MB/s | 304 | 2 | 6.3× |
| Easyjson | 496 | 381.14 MB/s | 0 | 0 | 4.9× |
| SonicFastest | 775 | 243.97 MB/s | 495 | 4 | 3.1× |
| Sonic | 780 | 242.22 MB/s | 495 | 4 | 3.1× |
| JSONV2 | 1021 | 185.11 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1220 | 109.86 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2424 | 77.96 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1540 | 1422.79 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1569 | 1396.40 MB/s | 0 | 0 | 10.2× |
| Easyjson | 3206 | 683.37 MB/s | 24 | 1 | 5.0× |
| Goccy | 3257 | 672.75 MB/s | 2864 | 4 | 4.9× |
| Sonic | 6381 | 343.37 MB/s | 3988 | 40 | 2.5× |
| SonicFastest | 6409 | 341.85 MB/s | 4001 | 40 | 2.5× |
| JSONV2 | 7999 | 273.92 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8168 | 221.73 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 15999 | 136.95 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 706303 | 722.74 MB/s | 703777 | 1012 | 8.6× |
| Lightning | 731576 | 697.78 MB/s | 703778 | 1012 | 8.3× |
| Goccy | 1142214 | 446.92 MB/s | 1138603 | 5006 | 5.3× |
| Sonic | 1173194 | 435.12 MB/s | 914063 | 2006 | 5.2× |
| SonicFastest | 1173718 | 434.92 MB/s | 914692 | 2006 | 5.2× |
| Easyjson | 1528595 | 333.95 MB/s | 863777 | 3012 | 4.0× |
| JSONV2 | 3220314 | 158.52 MB/s | 1076017 | 12646 | 1.9× |
| LightningDecodeAny | 3508484 | 131.53 MB/s | 2785929 | 66022 | 1.7× |
| Stdlib | 6048206 | 84.40 MB/s | 1162119 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14647.03 MB/s | 0 | 0 | 82.6× |
| LightningDestructive | 1390 | 14234.42 MB/s | 0 | 0 | 80.3× |
| Goccy | 19792 | 999.83 MB/s | 20491 | 2 | 5.6× |
| SonicFastest | 27941 | 708.25 MB/s | 22200 | 4 | 4.0× |
| Sonic | 28029 | 706.03 MB/s | 22162 | 4 | 4.0× |
| JSONV2 | 29552 | 669.63 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 77655 | 254.82 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86017 | 230.06 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111644 | 177.25 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2658 | 6818.87 MB/s | 0 | 0 | 38.6× |
| Lightning | 2811 | 6447.53 MB/s | 432 | 2 | 36.5× |
| Easyjson | 3943 | 4596.66 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 9804 | 1848.62 MB/s | 22957 | 6 | 10.5× |
| Sonic | 9935 | 1824.20 MB/s | 22890 | 6 | 10.3× |
| Goccy | 16014 | 1131.79 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 16549 | 1080.53 MB/s | 29088 | 191 | 6.2× |
| JSONV2 | 44961 | 403.10 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102552 | 176.73 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2244707 | 894.77 MB/s | 2960389 | 7411 | 8.3× |
| Lightning | 2312326 | 868.60 MB/s | 2962103 | 7417 | 8.1× |
| Goccy | 4314305 | 465.54 MB/s | 5413052 | 15836 | 4.3× |
| SonicFastest | 4426091 | 453.79 MB/s | 10961192 | 13683 | 4.2× |
| Sonic | 4441152 | 452.25 MB/s | 10939593 | 13683 | 4.2× |
| Easyjson | 4936913 | 406.83 MB/s | 2981515 | 7439 | 3.8× |
| JSONV2 | 7004288 | 286.75 MB/s | 3173694 | 14563 | 2.7× |
| LightningDecodeAny | 7157686 | 159.59 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18693338 | 107.44 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 894 | 614.35 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 899 | 610.72 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1874 | 292.45 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2198 | 249.77 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2677 | 205.12 MB/s | 1962 | 26 | 2.1× |
| SonicFastest | 2686 | 204.42 MB/s | 1966 | 26 | 2.1× |
| Goccy | 3056 | 179.63 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3336 | 164.57 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5657 | 97.04 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 494990 | 1275.81 MB/s | 364472 | 566 | 10.9× |
| Lightning | 554105 | 1139.70 MB/s | 413002 | 878 | 9.8× |
| Sonic | 1008599 | 626.13 MB/s | 985341 | 1102 | 5.4× |
| SonicFastest | 1009276 | 625.71 MB/s | 991847 | 1102 | 5.4× |
| Easyjson | 1147849 | 550.17 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1161063 | 543.91 MB/s | 987020 | 1201 | 4.7× |
| JSONV2 | 2170078 | 291.01 MB/s | 571618 | 3144 | 2.5× |
| LightningDecodeAny | 2367821 | 197.19 MB/s | 2010201 | 30295 | 2.3× |
| Stdlib | 5403721 | 116.87 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 710935 | 791.08 MB/s | 544252 | 448 | 7.4× |
| Lightning | 876131 | 641.92 MB/s | 767621 | 1254 | 6.0× |
| SonicFastest | 1026484 | 547.90 MB/s | 955599 | 1476 | 5.1× |
| Sonic | 1034199 | 543.81 MB/s | 961894 | 1476 | 5.1× |
| Goccy | 1332564 | 422.05 MB/s | 1042389 | 1030 | 3.9× |
| Easyjson | 1728960 | 325.29 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2724812 | 206.40 MB/s | 2114151 | 30295 | 1.9× |
| JSONV2 | 2745371 | 204.86 MB/s | 927442 | 3482 | 1.9× |
| Stdlib | 5256996 | 106.98 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 642165 | 830.28 MB/s | 333416 | 2084 | 8.5× |
| Lightning | 662498 | 804.80 MB/s | 368224 | 2293 | 8.3× |
| Easyjson | 1098820 | 485.23 MB/s | 428361 | 3273 | 5.0× |
| SonicFastest | 1128714 | 472.38 MB/s | 1050830 | 4351 | 4.8× |
| Sonic | 1129099 | 472.22 MB/s | 1050830 | 4351 | 4.8× |
| Goccy | 1287280 | 414.19 MB/s | 1167239 | 5409 | 4.2× |
| JSONV2 | 2531172 | 210.64 MB/s | 745449 | 13288 | 2.2× |
| LightningDecodeAny | 3306214 | 161.27 MB/s | 2674619 | 50596 | 1.7× |
| Stdlib | 5466777 | 97.53 MB/s | 798693 | 17133 | 1.0× |
