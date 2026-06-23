# JSON Deserialization Benchmarks

- generated 2026-06-23T10:13:43Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 104313 | 1220.13 MB/s | 49280 | 2 | 10.6× |
| Lightning | 104349 | 1219.71 MB/s | 49760 | 3 | 10.6× |
| SonicFastest | 182178 | 698.63 MB/s | 195508 | 10 | 6.1× |
| Sonic | 183804 | 692.45 MB/s | 199719 | 10 | 6.0× |
| Goccy | 193569 | 657.52 MB/s | 224799 | 884 | 5.7× |
| Easyjson | 212538 | 598.83 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 421666 | 301.84 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 447109 | 211.70 MB/s | 465585 | 9714 | 2.5× |
| Stdlib | 1104903 | 115.19 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3778519 | 595.75 MB/s | 3122876 | 3081 | 6.9× |
| LightningDestructive | 3782598 | 595.11 MB/s | 3122872 | 3081 | 6.9× |
| SonicFastest | 4470322 | 503.55 MB/s | 15243790 | 970 | 5.8× |
| Sonic | 4560363 | 493.61 MB/s | 15232102 | 970 | 5.7× |
| Goccy | 10420892 | 216.01 MB/s | 4123800 | 56532 | 2.5× |
| Easyjson | 11073355 | 203.29 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11484012 | 196.01 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16183259 | 139.10 MB/s | 3123214 | 3083 | 1.6× |
| Stdlib | 26119046 | 86.18 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 489108 | 552.85 MB/s | 348024 | 1627 | 6.8× |
| LightningDestructive | 494682 | 546.62 MB/s | 348024 | 1627 | 6.8× |
| Sonic | 623810 | 433.47 MB/s | 473955 | 968 | 5.4× |
| SonicFastest | 625063 | 432.60 MB/s | 474886 | 968 | 5.4× |
| Goccy | 1405494 | 192.39 MB/s | 544008 | 8123 | 2.4× |
| Easyjson | 1406257 | 192.29 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1600060 | 169.00 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2126355 | 127.17 MB/s | 348153 | 1628 | 1.6× |
| Stdlib | 3349010 | 80.74 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1363212 | 1267.01 MB/s | 959848 | 5881 | 9.7× |
| Lightning | 1391044 | 1241.66 MB/s | 959890 | 5882 | 9.5× |
| SonicFastest | 2044747 | 844.70 MB/s | 2740246 | 4020 | 6.5× |
| Sonic | 2045088 | 844.56 MB/s | 2717019 | 4020 | 6.5× |
| Goccy | 2383967 | 724.51 MB/s | 2583257 | 14605 | 5.6× |
| JSONV2 | 4196154 | 411.62 MB/s | 1011636 | 7594 | 3.2× |
| Easyjson | 4222243 | 409.07 MB/s | 972032 | 5389 | 3.1× |
| LightningDecodeAny | 4652699 | 107.53 MB/s | 4976204 | 81466 | 2.9× |
| Stdlib | 13272001 | 130.14 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1192 | 1520.63 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1199 | 1511.37 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2547 | 711.52 MB/s | 24 | 1 | 5.5× |
| Goccy | 2881 | 628.87 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6012 | 301.39 MB/s | 3821 | 40 | 2.3× |
| Sonic | 6057 | 299.18 MB/s | 3835 | 40 | 2.3× |
| JSONV2 | 7795 | 232.47 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8504 | 212.96 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14044 | 129.02 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1219 | 1486.95 MB/s | 0 | 0 | 11.6× |
| LightningDestructive | 1241 | 1459.87 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2541 | 713.11 MB/s | 24 | 1 | 5.5× |
| Goccy | 2915 | 621.60 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 5983 | 302.84 MB/s | 3852 | 40 | 2.4× |
| Sonic | 5999 | 302.03 MB/s | 3833 | 40 | 2.3× |
| JSONV2 | 7791 | 232.57 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8547 | 211.88 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14093 | 128.57 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1397 | 1297.33 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1446 | 1253.21 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2767 | 654.80 MB/s | 144 | 10 | 5.1× |
| Goccy | 2852 | 635.36 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6162 | 294.08 MB/s | 3793 | 42 | 2.3× |
| Sonic | 6170 | 293.67 MB/s | 3776 | 42 | 2.3× |
| JSONV2 | 7995 | 226.65 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8178 | 221.43 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14022 | 129.23 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 736 | 670.78 MB/s | 160 | 1 | 7.5× |
| Lightning | 741 | 666.70 MB/s | 160 | 1 | 7.5× |
| Sonic | 1232 | 400.92 MB/s | 972 | 6 | 4.5× |
| SonicFastest | 1232 | 401.13 MB/s | 983 | 6 | 4.5× |
| LightningDecodeAny | 1647 | 299.35 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2231 | 221.42 MB/s | 448 | 3 | 2.5× |
| Goccy | 2445 | 202.02 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3230 | 152.96 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5538 | 89.21 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 475 | 484.12 MB/s | 160 | 1 | 8.7× |
| LightningDestructive | 477 | 481.74 MB/s | 160 | 1 | 8.7× |
| Sonic | 903 | 254.73 MB/s | 693 | 6 | 4.6× |
| SonicFastest | 904 | 254.41 MB/s | 694 | 6 | 4.6× |
| Easyjson | 1427 | 161.12 MB/s | 448 | 3 | 2.9× |
| LightningDecodeAny | 1433 | 159.80 MB/s | 1536 | 30 | 2.9× |
| Goccy | 1624 | 141.64 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2378 | 96.74 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4139 | 55.57 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 72585 | 897.31 MB/s | 166212 | 102 | 7.5× |
| Lightning | 74355 | 875.96 MB/s | 172433 | 107 | 7.3× |
| SonicFastest | 96538 | 674.68 MB/s | 155811 | 75 | 5.6× |
| Sonic | 98027 | 664.43 MB/s | 157215 | 75 | 5.6× |
| Goccy | 140486 | 463.62 MB/s | 229122 | 134 | 3.9× |
| LightningDecodeAny | 184279 | 289.39 MB/s | 176960 | 3252 | 3.0× |
| JSONV2 | 221196 | 294.45 MB/s | 206650 | 607 | 2.5× |
| Stdlib | 544426 | 119.63 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2587853 | 749.84 MB/s | 2846864 | 2238 | 9.1× |
| Lightning | 2681974 | 723.52 MB/s | 2846867 | 2238 | 8.8× |
| Goccy | 4796233 | 404.58 MB/s | 4065829 | 13510 | 4.9× |
| SonicFastest | 4872370 | 398.26 MB/s | 14608713 | 1407 | 4.8× |
| Sonic | 5011248 | 387.22 MB/s | 14608661 | 1407 | 4.7× |
| Easyjson | 7485078 | 259.25 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9643965 | 201.21 MB/s | 7013523 | 219937 | 2.4× |
| JSONV2 | 11270675 | 172.17 MB/s | 3237231 | 13947 | 2.1× |
| Stdlib | 23483582 | 82.63 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1107833 | 3003.91 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1779743 | 1869.84 MB/s | 2488905 | 2995 | 11.7× |
| SonicFastest | 2650145 | 1255.72 MB/s | 6418378 | 4248 | 7.9× |
| Sonic | 2651474 | 1255.09 MB/s | 6317592 | 4248 | 7.9× |
| LightningDecodeAny | 3732912 | 823.42 MB/s | 4886620 | 56892 | 5.6× |
| Goccy | 4610490 | 721.80 MB/s | 3948908 | 3816 | 4.5× |
| JSONV2 | 7329600 | 454.03 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20847326 | 159.63 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220958 | 997.23 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 222990 | 988.14 MB/s | 136897 | 228 | 9.2× |
| Sonic | 383979 | 573.85 MB/s | 315921 | 398 | 5.3× |
| SonicFastest | 390331 | 564.51 MB/s | 328925 | 398 | 5.2× |
| Goccy | 438715 | 502.25 MB/s | 365039 | 1067 | 4.7× |
| Easyjson | 547227 | 402.66 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 739729 | 297.87 MB/s | 129743 | 470 | 2.8× |
| LightningDecodeAny | 880528 | 123.01 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2041771 | 107.92 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12809151 | 632.36 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13157455 | 615.62 MB/s | 15059836 | 51649 | 6.8× |
| Sonic | 16771172 | 482.97 MB/s | 70873027 | 40014 | 5.3× |
| SonicFastest | 16820739 | 481.55 MB/s | 70901523 | 40014 | 5.3× |
| Goccy | 23353810 | 346.84 MB/s | 17033663 | 107148 | 3.8× |
| Easyjson | 30896997 | 262.16 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 40714530 | 127.79 MB/s | 34333353 | 912810 | 2.2× |
| JSONV2 | 44186241 | 183.32 MB/s | 15233739 | 78972 | 2.0× |
| Stdlib | 89551778 | 90.45 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6103004 | 488.85 MB/s | 3985336 | 30015 | 7.7× |
| Lightning | 6209671 | 480.45 MB/s | 3985336 | 30015 | 7.6× |
| Sonic | 8642744 | 345.20 MB/s | 26568256 | 56760 | 5.4× |
| SonicFastest | 8670976 | 344.07 MB/s | 26638834 | 56760 | 5.4× |
| Easyjson | 16492638 | 180.90 MB/s | 9479441 | 30115 | 2.9× |
| Goccy | 16585816 | 179.88 MB/s | 10580059 | 273647 | 2.8× |
| LightningDecodeAny | 18912310 | 96.98 MB/s | 20023837 | 408557 | 2.5× |
| JSONV2 | 24229910 | 123.13 MB/s | 9257149 | 86278 | 1.9× |
| Stdlib | 47076512 | 63.37 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1371945 | 527.42 MB/s | 907393 | 3618 | 8.4× |
| Lightning | 1387745 | 521.42 MB/s | 907388 | 3618 | 8.3× |
| Sonic | 1774383 | 407.80 MB/s | 3181416 | 7226 | 6.5× |
| SonicFastest | 1784740 | 405.44 MB/s | 3181063 | 7226 | 6.5× |
| Easyjson | 4253516 | 170.12 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4374041 | 148.73 MB/s | 5752203 | 80172 | 2.6× |
| Goccy | 4800556 | 150.73 MB/s | 2859389 | 80276 | 2.4× |
| JSONV2 | 5724852 | 126.40 MB/s | 2704627 | 7318 | 2.0× |
| Stdlib | 11584908 | 62.46 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1947466 | 809.95 MB/s | 907393 | 3618 | 8.1× |
| Lightning | 1979365 | 796.90 MB/s | 907387 | 3618 | 7.9× |
| SonicFastest | 2270442 | 694.73 MB/s | 5791743 | 7226 | 6.9× |
| Sonic | 2271620 | 694.37 MB/s | 5807219 | 7226 | 6.9× |
| LightningDecodeAny | 3859736 | 195.19 MB/s | 5752201 | 80172 | 4.1× |
| Easyjson | 5601158 | 281.61 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5703124 | 276.58 MB/s | 3645640 | 80270 | 2.8× |
| JSONV2 | 6530466 | 241.54 MB/s | 2704594 | 7318 | 2.4× |
| Stdlib | 15716438 | 100.36 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 223906 | 670.48 MB/s | 81920 | 1 | 8.2× |
| Lightning | 224355 | 669.14 MB/s | 81920 | 1 | 8.2× |
| SonicFastest | 270725 | 554.53 MB/s | 253100 | 6 | 6.8× |
| Sonic | 275648 | 544.62 MB/s | 268094 | 6 | 6.6× |
| LightningDecodeAny | 476987 | 314.73 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 860058 | 174.55 MB/s | 325429 | 10005 | 2.1× |
| JSONV2 | 1094567 | 137.15 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1830647 | 82.01 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33875 | 830.03 MB/s | 30272 | 105 | 9.0× |
| LightningDestructive | 34024 | 826.39 MB/s | 30144 | 103 | 8.9× |
| SonicFastest | 64148 | 438.31 MB/s | 48439 | 103 | 4.7× |
| Sonic | 64199 | 437.97 MB/s | 48136 | 103 | 4.7× |
| Easyjson | 68265 | 411.88 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72389 | 388.41 MB/s | 59259 | 188 | 4.2× |
| JSONV2 | 134736 | 208.68 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 156067 | 180.16 MB/s | 135024 | 2678 | 1.9× |
| Stdlib | 303499 | 92.64 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1955 | 1190.97 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2013 | 1156.57 MB/s | 32 | 1 | 11.3× |
| Goccy | 4058 | 573.66 MB/s | 3649 | 4 | 5.6× |
| Easyjson | 4200 | 554.27 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5010 | 464.66 MB/s | 4243 | 6 | 4.5× |
| Sonic | 5033 | 462.53 MB/s | 4253 | 6 | 4.5× |
| JSONV2 | 8433 | 276.06 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10223 | 164.82 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22750 | 102.33 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221 | 856.74 MB/s | 0 | 0 | 11.2× |
| LightningDestructive | 222 | 851.68 MB/s | 0 | 0 | 11.1× |
| Goccy | 387 | 488.32 MB/s | 304 | 2 | 6.4× |
| Easyjson | 496 | 381.37 MB/s | 0 | 0 | 5.0× |
| Sonic | 793 | 238.27 MB/s | 499 | 4 | 3.1× |
| SonicFastest | 794 | 237.93 MB/s | 505 | 4 | 3.1× |
| JSONV2 | 1035 | 182.56 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1230 | 108.96 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2464 | 76.69 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1539 | 1423.25 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1571 | 1394.69 MB/s | 0 | 0 | 10.2× |
| Easyjson | 3203 | 684.03 MB/s | 24 | 1 | 5.0× |
| Goccy | 3220 | 680.38 MB/s | 2864 | 4 | 5.0× |
| Sonic | 6389 | 342.92 MB/s | 3974 | 40 | 2.5× |
| SonicFastest | 6398 | 342.47 MB/s | 4011 | 40 | 2.5× |
| JSONV2 | 7996 | 274.02 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8235 | 219.91 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16045 | 136.56 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 699598 | 729.67 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 731815 | 697.55 MB/s | 703779 | 1012 | 8.2× |
| Goccy | 1140776 | 447.48 MB/s | 1135890 | 5006 | 5.3× |
| SonicFastest | 1159536 | 440.24 MB/s | 865165 | 2006 | 5.2× |
| Sonic | 1162329 | 439.18 MB/s | 871042 | 2006 | 5.2× |
| Easyjson | 1524032 | 334.95 MB/s | 863777 | 3012 | 4.0× |
| JSONV2 | 3230339 | 158.03 MB/s | 1076012 | 12646 | 1.9× |
| LightningDecodeAny | 3523151 | 130.98 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6036104 | 84.57 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14643.95 MB/s | 0 | 0 | 82.6× |
| LightningDestructive | 1392 | 14217.20 MB/s | 0 | 0 | 80.2× |
| Goccy | 19912 | 993.81 MB/s | 20491 | 2 | 5.6× |
| Sonic | 28092 | 704.42 MB/s | 22483 | 4 | 4.0× |
| SonicFastest | 28184 | 702.15 MB/s | 22699 | 4 | 4.0× |
| JSONV2 | 29554 | 669.59 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 79183 | 249.90 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86060 | 229.94 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111609 | 177.31 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2660 | 6814.22 MB/s | 0 | 0 | 38.6× |
| Lightning | 2814 | 6441.73 MB/s | 432 | 2 | 36.5× |
| Easyjson | 3949 | 4589.64 MB/s | 432 | 2 | 26.0× |
| Sonic | 9803 | 1848.80 MB/s | 22626 | 6 | 10.5× |
| SonicFastest | 9836 | 1842.67 MB/s | 22916 | 6 | 10.4× |
| Goccy | 16244 | 1115.74 MB/s | 19459 | 2 | 6.3× |
| LightningDecodeAny | 16712 | 1070.02 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 44976 | 402.97 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102629 | 176.60 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2243552 | 895.23 MB/s | 2960388 | 7411 | 8.3× |
| Lightning | 2291372 | 876.55 MB/s | 2962102 | 7417 | 8.1× |
| Goccy | 4275113 | 469.81 MB/s | 5411771 | 15830 | 4.4× |
| Sonic | 4458177 | 450.52 MB/s | 10892330 | 13683 | 4.2× |
| SonicFastest | 4466705 | 449.66 MB/s | 10909017 | 13683 | 4.2× |
| Easyjson | 4918985 | 408.31 MB/s | 2981487 | 7439 | 3.8× |
| JSONV2 | 7006098 | 286.68 MB/s | 3173683 | 14563 | 2.7× |
| LightningDecodeAny | 7227232 | 158.05 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18673357 | 107.56 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 885 | 620.35 MB/s | 480 | 1 | 6.4× |
| LightningDestructive | 889 | 617.56 MB/s | 480 | 1 | 6.4× |
| LightningDecodeAny | 1890 | 290.00 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2160 | 254.19 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2690 | 204.06 MB/s | 1962 | 26 | 2.1× |
| Sonic | 2698 | 203.49 MB/s | 1954 | 26 | 2.1× |
| Goccy | 3028 | 181.30 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3278 | 167.46 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5655 | 97.08 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 497822 | 1268.55 MB/s | 364472 | 566 | 10.8× |
| Lightning | 552136 | 1143.76 MB/s | 413001 | 878 | 9.8× |
| Sonic | 1018134 | 620.27 MB/s | 997092 | 1102 | 5.3× |
| SonicFastest | 1022040 | 617.90 MB/s | 1002757 | 1102 | 5.3× |
| Easyjson | 1148869 | 549.68 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1165940 | 541.64 MB/s | 988010 | 1202 | 4.6× |
| JSONV2 | 2169404 | 291.10 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2379032 | 196.26 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5397850 | 116.99 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 714298 | 787.36 MB/s | 544252 | 448 | 7.4× |
| Lightning | 878785 | 639.98 MB/s | 767621 | 1254 | 6.0× |
| SonicFastest | 1021892 | 550.36 MB/s | 950352 | 1476 | 5.1× |
| Sonic | 1023269 | 549.62 MB/s | 949093 | 1476 | 5.1× |
| Goccy | 1325538 | 424.29 MB/s | 1038877 | 1030 | 4.0× |
| Easyjson | 1729082 | 325.26 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2719810 | 206.78 MB/s | 2114152 | 30295 | 1.9× |
| JSONV2 | 2748748 | 204.61 MB/s | 927443 | 3482 | 1.9× |
| Stdlib | 5261772 | 106.89 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 651977 | 817.79 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 674573 | 790.39 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1108573 | 480.96 MB/s | 428361 | 3273 | 5.0× |
| Sonic | 1139244 | 468.01 MB/s | 1031523 | 4351 | 4.8× |
| SonicFastest | 1149192 | 463.96 MB/s | 1047053 | 4351 | 4.8× |
| Goccy | 1293077 | 412.33 MB/s | 1167220 | 5409 | 4.2× |
| JSONV2 | 2545629 | 209.45 MB/s | 745448 | 13288 | 2.2× |
| LightningDecodeAny | 3379774 | 157.76 MB/s | 2674619 | 50596 | 1.6× |
| Stdlib | 5490112 | 97.12 MB/s | 798692 | 17133 | 1.0× |
