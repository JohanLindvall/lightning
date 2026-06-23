# JSON Deserialization Benchmarks

- generated 2026-06-23T19:53:17Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 102165 | 1245.77 MB/s | 49760 | 3 | 12.1× |
| LightningDestructive | 104072 | 1222.96 MB/s | 49280 | 2 | 11.8× |
| Sonic | 193896 | 656.41 MB/s | 214000 | 15 | 6.4× |
| SonicFastest | 196453 | 647.87 MB/s | 214840 | 15 | 6.3× |
| Easyjson | 230956 | 551.08 MB/s | 122864 | 14 | 5.3× |
| Goccy | 240751 | 528.66 MB/s | 225181 | 884 | 5.1× |
| JSONV2 | 389194 | 327.02 MB/s | 195128 | 1805 | 3.2× |
| LightningDecodeAny | 437715 | 216.24 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1232769 | 103.24 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4117627 | 546.69 MB/s | 3122873 | 3081 | 7.6× |
| Lightning | 4151580 | 542.22 MB/s | 3122874 | 3081 | 7.5× |
| Sonic | 5052298 | 445.55 MB/s | 4867689 | 2584 | 6.2× |
| SonicFastest | 5328988 | 422.42 MB/s | 4868120 | 2584 | 5.8× |
| LightningDecodeAny | 11502913 | 195.69 MB/s | 7938300 | 281383 | 2.7× |
| Goccy | 12623192 | 178.33 MB/s | 4237478 | 56537 | 2.5× |
| Easyjson | 13593667 | 165.60 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 17129685 | 131.41 MB/s | 3123189 | 3083 | 1.8× |
| Stdlib | 31103680 | 72.37 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 552618 | 489.31 MB/s | 348024 | 1627 | 7.3× |
| LightningDestructive | 560789 | 482.18 MB/s | 348024 | 1627 | 7.2× |
| SonicFastest | 732398 | 369.20 MB/s | 641951 | 1147 | 5.5× |
| Sonic | 733267 | 368.76 MB/s | 642071 | 1147 | 5.5× |
| LightningDecodeAny | 1584448 | 170.66 MB/s | 1011487 | 37901 | 2.5× |
| Goccy | 1672102 | 161.71 MB/s | 542521 | 8122 | 2.4× |
| Easyjson | 1731666 | 156.15 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2261183 | 119.58 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4033616 | 67.04 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1297304 | 1331.38 MB/s | 959848 | 5881 | 13.0× |
| Lightning | 1336572 | 1292.26 MB/s | 959890 | 5882 | 12.6× |
| Sonic | 2117145 | 815.82 MB/s | 2693290 | 5547 | 7.9× |
| SonicFastest | 2149000 | 803.72 MB/s | 2693243 | 5547 | 7.8× |
| Goccy | 2446132 | 706.10 MB/s | 2580454 | 14603 | 6.9× |
| Easyjson | 3691540 | 467.88 MB/s | 972032 | 5389 | 4.6× |
| LightningDecodeAny | 4018828 | 124.49 MB/s | 4976203 | 81466 | 4.2× |
| JSONV2 | 4291073 | 402.51 MB/s | 1011614 | 7594 | 3.9× |
| Stdlib | 16823297 | 102.67 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1018 | 1780.49 MB/s | 0 | 0 | 14.5× |
| LightningDestructive | 1086 | 1668.16 MB/s | 0 | 0 | 13.6× |
| Easyjson | 2794 | 648.46 MB/s | 24 | 1 | 5.3× |
| Goccy | 3155 | 574.24 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6223 | 291.16 MB/s | 3347 | 38 | 2.4× |
| Sonic | 6390 | 283.58 MB/s | 3345 | 38 | 2.3× |
| JSONV2 | 7472 | 242.51 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8484 | 213.47 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14719 | 123.10 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1087 | 1666.78 MB/s | 0 | 0 | 13.4× |
| LightningDestructive | 1129 | 1605.39 MB/s | 0 | 0 | 12.9× |
| Easyjson | 2813 | 644.14 MB/s | 24 | 1 | 5.2× |
| Goccy | 3100 | 584.49 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6052 | 299.42 MB/s | 3348 | 38 | 2.4× |
| Sonic | 6273 | 288.87 MB/s | 3347 | 38 | 2.3× |
| JSONV2 | 7474 | 242.42 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 8458 | 214.12 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14552 | 124.52 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1239 | 1463.05 MB/s | 144 | 10 | 11.8× |
| LightningDestructive | 1293 | 1401.87 MB/s | 144 | 10 | 11.3× |
| Easyjson | 2927 | 619.16 MB/s | 144 | 10 | 5.0× |
| Goccy | 3308 | 547.75 MB/s | 2600 | 5 | 4.4× |
| SonicFastest | 6312 | 287.08 MB/s | 3366 | 40 | 2.3× |
| Sonic | 6525 | 277.69 MB/s | 3367 | 40 | 2.2× |
| JSONV2 | 8026 | 225.76 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8534 | 212.22 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14559 | 124.46 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 738 | 669.83 MB/s | 160 | 1 | 8.2× |
| LightningDestructive | 746 | 661.92 MB/s | 160 | 1 | 8.1× |
| Sonic | 1223 | 403.97 MB/s | 1075 | 8 | 4.9× |
| SonicFastest | 1231 | 401.17 MB/s | 1075 | 8 | 4.9× |
| LightningDecodeAny | 1609 | 306.36 MB/s | 1536 | 30 | 3.7× |
| Easyjson | 2279 | 216.78 MB/s | 448 | 3 | 2.6× |
| Goccy | 2417 | 204.40 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3145 | 157.07 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6022 | 82.03 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 465 | 494.77 MB/s | 160 | 1 | 9.0× |
| LightningDestructive | 469 | 490.64 MB/s | 160 | 1 | 9.0× |
| SonicFastest | 857 | 268.48 MB/s | 801 | 8 | 4.9× |
| Sonic | 859 | 267.79 MB/s | 801 | 8 | 4.9× |
| LightningDecodeAny | 1390 | 164.71 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1587 | 144.89 MB/s | 448 | 3 | 2.7× |
| Goccy | 1658 | 138.71 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2373 | 96.92 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4207 | 54.68 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 83076 | 784.01 MB/s | 172433 | 107 | 7.7× |
| LightningDestructive | 86865 | 749.80 MB/s | 166213 | 102 | 7.4× |
| Sonic | 150389 | 433.09 MB/s | 235845 | 65 | 4.2× |
| SonicFastest | 150640 | 432.37 MB/s | 235853 | 65 | 4.2× |
| Goccy | 175292 | 371.56 MB/s | 228218 | 134 | 3.6× |
| LightningDecodeAny | 180975 | 294.68 MB/s | 176960 | 3252 | 3.5× |
| JSONV2 | 244868 | 265.99 MB/s | 206663 | 607 | 2.6× |
| Stdlib | 639055 | 101.92 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2536851 | 764.91 MB/s | 2846865 | 2238 | 9.9× |
| Lightning | 2590834 | 748.98 MB/s | 2846865 | 2238 | 9.7× |
| SonicFastest | 4768266 | 406.96 MB/s | 4880113 | 1736 | 5.3× |
| Sonic | 4789642 | 405.14 MB/s | 4879242 | 1736 | 5.2× |
| Goccy | 5012695 | 387.11 MB/s | 4063127 | 13509 | 5.0× |
| Easyjson | 7512380 | 258.30 MB/s | 3871264 | 15043 | 3.3× |
| LightningDecodeAny | 9275883 | 209.20 MB/s | 7013524 | 219937 | 2.7× |
| JSONV2 | 11526870 | 168.34 MB/s | 3237180 | 13947 | 2.2× |
| Stdlib | 25143427 | 77.18 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1009307 | 3297.15 MB/s | 351704 | 1286 | 25.1× |
| Lightning | 1535043 | 2167.91 MB/s | 2488906 | 2995 | 16.5× |
| Sonic | 2503574 | 1329.23 MB/s | 5870699 | 4263 | 10.1× |
| SonicFastest | 2519469 | 1320.85 MB/s | 5870694 | 4263 | 10.1× |
| LightningDecodeAny | 3270932 | 939.72 MB/s | 4886622 | 56892 | 7.7× |
| Goccy | 6568857 | 506.61 MB/s | 3948912 | 3816 | 3.9× |
| JSONV2 | 7903369 | 421.06 MB/s | 5364502 | 13243 | 3.2× |
| Stdlib | 25339576 | 131.33 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225326 | 977.90 MB/s | 136896 | 228 | 9.8× |
| LightningDestructive | 233887 | 942.10 MB/s | 136896 | 228 | 9.5× |
| Goccy | 455761 | 483.47 MB/s | 363651 | 1066 | 4.9× |
| Sonic | 499019 | 441.56 MB/s | 350806 | 262 | 4.4× |
| SonicFastest | 500963 | 439.84 MB/s | 350592 | 262 | 4.4× |
| Easyjson | 595677 | 369.91 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 648562 | 339.75 MB/s | 129746 | 470 | 3.4× |
| LightningDecodeAny | 945593 | 114.55 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2211781 | 99.62 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14420558 | 561.70 MB/s | 15059834 | 51649 | 7.0× |
| Lightning | 14595985 | 554.95 MB/s | 15059837 | 51649 | 7.0× |
| SonicFastest | 21137591 | 383.21 MB/s | 19849861 | 41640 | 4.8× |
| Sonic | 21215796 | 381.79 MB/s | 19849325 | 41640 | 4.8× |
| Goccy | 24391257 | 332.09 MB/s | 19225473 | 107156 | 4.2× |
| Easyjson | 34730663 | 233.22 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 40074351 | 129.83 MB/s | 34333345 | 912810 | 2.5× |
| JSONV2 | 45446099 | 178.23 MB/s | 15233741 | 78972 | 2.2× |
| Stdlib | 101518354 | 79.79 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5919567 | 504.00 MB/s | 3985336 | 30015 | 8.5× |
| Lightning | 6126732 | 486.96 MB/s | 3985337 | 30015 | 8.3× |
| Sonic | 9149739 | 326.07 MB/s | 9128589 | 57804 | 5.5× |
| SonicFastest | 9174963 | 325.17 MB/s | 9128822 | 57804 | 5.5× |
| Goccy | 17545017 | 170.05 MB/s | 9917701 | 273622 | 2.9× |
| Easyjson | 17898102 | 166.69 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 19518259 | 93.97 MB/s | 20023837 | 408557 | 2.6× |
| JSONV2 | 25104111 | 118.84 MB/s | 9257045 | 86278 | 2.0× |
| Stdlib | 50590700 | 58.97 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1363593 | 530.65 MB/s | 907394 | 3618 | 9.0× |
| Lightning | 1431702 | 505.41 MB/s | 907386 | 3618 | 8.5× |
| Sonic | 2190633 | 330.31 MB/s | 2367691 | 3683 | 5.6× |
| SonicFastest | 2197119 | 329.34 MB/s | 2367582 | 3683 | 5.6× |
| LightningDecodeAny | 5232710 | 124.33 MB/s | 5752201 | 80172 | 2.3× |
| Easyjson | 5237630 | 138.15 MB/s | 2847905 | 3698 | 2.3× |
| Goccy | 5352655 | 135.18 MB/s | 2640847 | 80264 | 2.3× |
| JSONV2 | 5992482 | 120.75 MB/s | 2704705 | 7318 | 2.0× |
| Stdlib | 12211013 | 59.26 MB/s | 2704547 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2071529 | 761.44 MB/s | 907387 | 3618 | 8.5× |
| LightningDestructive | 2082097 | 757.58 MB/s | 907392 | 3618 | 8.5× |
| SonicFastest | 2429566 | 649.23 MB/s | 3221617 | 3683 | 7.3× |
| Sonic | 2433205 | 648.26 MB/s | 3222244 | 3683 | 7.3× |
| LightningDecodeAny | 4374839 | 172.21 MB/s | 5752203 | 80172 | 4.0× |
| Easyjson | 6251844 | 252.30 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 6571850 | 240.02 MB/s | 3470094 | 80260 | 2.7× |
| JSONV2 | 6671830 | 236.42 MB/s | 2704554 | 7318 | 2.7× |
| Stdlib | 17692621 | 89.15 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 236519 | 634.72 MB/s | 81920 | 1 | 9.4× |
| LightningDestructive | 245093 | 612.52 MB/s | 81920 | 1 | 9.0× |
| Sonic | 388459 | 386.46 MB/s | 407553 | 16 | 5.7× |
| SonicFastest | 411937 | 364.43 MB/s | 406990 | 16 | 5.4× |
| LightningDecodeAny | 589735 | 254.56 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 1026456 | 146.25 MB/s | 324572 | 10005 | 2.2× |
| JSONV2 | 1119614 | 134.09 MB/s | 357728 | 20 | 2.0× |
| Stdlib | 2213912 | 67.81 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 31500 | 892.61 MB/s | 30272 | 105 | 9.9× |
| LightningDestructive | 32604 | 862.38 MB/s | 30144 | 103 | 9.5× |
| Sonic | 69229 | 406.14 MB/s | 59452 | 83 | 4.5× |
| SonicFastest | 69249 | 406.03 MB/s | 59461 | 83 | 4.5× |
| Goccy | 74982 | 374.98 MB/s | 59260 | 188 | 4.1× |
| Easyjson | 76465 | 367.71 MB/s | 32304 | 138 | 4.1× |
| JSONV2 | 122467 | 229.59 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 152757 | 184.06 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 310301 | 90.61 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1980 | 1175.87 MB/s | 32 | 1 | 12.1× |
| LightningDestructive | 2012 | 1157.21 MB/s | 32 | 1 | 11.9× |
| Goccy | 4650 | 500.60 MB/s | 3649 | 4 | 5.1× |
| Easyjson | 4770 | 488.08 MB/s | 192 | 2 | 5.0× |
| Sonic | 6036 | 385.67 MB/s | 3708 | 4 | 4.0× |
| SonicFastest | 6061 | 384.11 MB/s | 3712 | 4 | 3.9× |
| JSONV2 | 7819 | 297.74 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 9476 | 177.83 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 23881 | 97.48 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222 | 851.55 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 225 | 838.94 MB/s | 0 | 0 | 10.9× |
| Goccy | 394 | 480.04 MB/s | 304 | 2 | 6.2× |
| Easyjson | 556 | 339.84 MB/s | 0 | 0 | 4.4× |
| SonicFastest | 731 | 258.71 MB/s | 341 | 3 | 3.4× |
| Sonic | 735 | 256.99 MB/s | 341 | 3 | 3.3× |
| JSONV2 | 925 | 204.33 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1187 | 112.90 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2457 | 76.94 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1471 | 1489.79 MB/s | 0 | 0 | 11.7× |
| LightningDestructive | 1514 | 1447.37 MB/s | 0 | 0 | 11.4× |
| Goccy | 3444 | 636.18 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3492 | 627.48 MB/s | 24 | 1 | 4.9× |
| Sonic | 6676 | 328.18 MB/s | 3600 | 38 | 2.6× |
| SonicFastest | 6719 | 326.09 MB/s | 3602 | 38 | 2.6× |
| JSONV2 | 7607 | 288.02 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8502 | 213.00 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 17204 | 127.36 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 664484 | 768.23 MB/s | 703778 | 1012 | 9.6× |
| Lightning | 712674 | 716.28 MB/s | 703779 | 1012 | 8.9× |
| Goccy | 1270856 | 401.68 MB/s | 1132878 | 5006 | 5.0× |
| Sonic | 1451900 | 351.59 MB/s | 1308029 | 2014 | 4.4× |
| SonicFastest | 1469638 | 347.35 MB/s | 1307982 | 2014 | 4.3× |
| Easyjson | 1559777 | 327.28 MB/s | 863780 | 3012 | 4.1× |
| JSONV2 | 3160693 | 161.51 MB/s | 1075947 | 12645 | 2.0× |
| LightningDecodeAny | 3420227 | 134.92 MB/s | 2785930 | 66022 | 1.9× |
| Stdlib | 6377507 | 80.04 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 611 | 32377.52 MB/s | 0 | 0 | 234.0× |
| LightningDestructive | 870 | 22735.97 MB/s | 0 | 0 | 164.3× |
| SonicFastest | 6669 | 2967.19 MB/s | 21116 | 3 | 21.4× |
| Goccy | 25261 | 783.38 MB/s | 20492 | 2 | 5.7× |
| Sonic | 29175 | 678.29 MB/s | 20637 | 3 | 4.9× |
| JSONV2 | 36015 | 549.46 MB/s | 8 | 1 | 4.0× |
| LightningDecodeAny | 93626 | 211.35 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 111067 | 178.17 MB/s | 0 | 0 | 1.3× |
| Stdlib | 143045 | 138.34 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2223 | 8151.84 MB/s | 432 | 2 | 54.0× |
| LightningDestructive | 2329 | 7781.72 MB/s | 0 | 0 | 51.6× |
| Easyjson | 4802 | 3773.98 MB/s | 432 | 2 | 25.0× |
| SonicFastest | 9001 | 2013.46 MB/s | 20447 | 5 | 13.3× |
| Sonic | 9057 | 2001.12 MB/s | 20442 | 5 | 13.3× |
| LightningDecodeAny | 18317 | 976.26 MB/s | 29088 | 191 | 6.6× |
| Goccy | 21157 | 856.63 MB/s | 19460 | 2 | 5.7× |
| JSONV2 | 47759 | 379.49 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 120068 | 150.95 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2205121 | 910.83 MB/s | 2960389 | 7411 | 9.4× |
| Lightning | 2252211 | 891.79 MB/s | 2962100 | 7417 | 9.2× |
| Goccy | 4535040 | 442.88 MB/s | 5411405 | 15830 | 4.6× |
| SonicFastest | 4897822 | 410.08 MB/s | 5153101 | 7085 | 4.2× |
| Sonic | 4963439 | 404.66 MB/s | 5152929 | 7085 | 4.2× |
| Easyjson | 5285270 | 380.02 MB/s | 2981490 | 7439 | 3.9× |
| LightningDecodeAny | 6733861 | 169.63 MB/s | 7386073 | 134751 | 3.1× |
| JSONV2 | 7064905 | 284.29 MB/s | 3173673 | 14562 | 2.9× |
| Stdlib | 20667633 | 97.18 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 807 | 680.55 MB/s | 480 | 1 | 7.2× |
| LightningDestructive | 817 | 671.97 MB/s | 480 | 1 | 7.1× |
| LightningDecodeAny | 1865 | 293.83 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 1965 | 279.36 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2127 | 258.09 MB/s | 2263 | 8 | 2.7× |
| Sonic | 2162 | 253.98 MB/s | 2262 | 8 | 2.7× |
| Goccy | 2958 | 185.63 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3027 | 181.38 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 5785 | 94.91 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 479412 | 1317.27 MB/s | 364472 | 566 | 12.5× |
| Lightning | 544371 | 1160.08 MB/s | 413001 | 878 | 11.0× |
| SonicFastest | 1139885 | 554.02 MB/s | 1065512 | 814 | 5.2× |
| Sonic | 1140211 | 553.86 MB/s | 1065352 | 814 | 5.2× |
| Easyjson | 1222314 | 516.65 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1356905 | 465.41 MB/s | 991058 | 1200 | 4.4× |
| JSONV2 | 2148666 | 293.91 MB/s | 571591 | 3144 | 2.8× |
| LightningDecodeAny | 2362543 | 197.63 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 5970262 | 105.78 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 767259 | 733.01 MB/s | 544248 | 448 | 7.4× |
| Lightning | 957000 | 587.68 MB/s | 767623 | 1254 | 5.9× |
| SonicFastest | 1293537 | 434.78 MB/s | 1345524 | 1185 | 4.4× |
| Sonic | 1297166 | 433.57 MB/s | 1345917 | 1185 | 4.4× |
| Goccy | 1572959 | 357.55 MB/s | 1040083 | 1029 | 3.6× |
| Easyjson | 2008460 | 280.02 MB/s | 775155 | 1254 | 2.8× |
| LightningDecodeAny | 2964900 | 189.69 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 2980058 | 188.72 MB/s | 927406 | 3482 | 1.9× |
| Stdlib | 5652661 | 99.49 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 582984 | 914.57 MB/s | 333416 | 2084 | 10.2× |
| Lightning | 661434 | 806.09 MB/s | 368224 | 2293 | 9.0× |
| Easyjson | 1254322 | 425.07 MB/s | 428362 | 3273 | 4.7× |
| Sonic | 1396238 | 381.87 MB/s | 978847 | 3082 | 4.3× |
| SonicFastest | 1403377 | 379.93 MB/s | 979287 | 3082 | 4.2× |
| Goccy | 1407824 | 378.72 MB/s | 1167046 | 5408 | 4.2× |
| JSONV2 | 2607038 | 204.51 MB/s | 745421 | 13288 | 2.3× |
| LightningDecodeAny | 3251732 | 163.97 MB/s | 2674618 | 50596 | 1.8× |
| Stdlib | 5954761 | 89.54 MB/s | 798692 | 17133 | 1.0× |
