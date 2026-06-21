# JSON Deserialization Benchmarks

- generated 2026-06-21T13:05:14Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104463 | 1218.37 MB/s | 49760 | 3 | 10.6× |
| SonicFastest | 180562 | 704.88 MB/s | 194562 | 10 | 6.1× |
| Sonic | 182536 | 697.26 MB/s | 196696 | 10 | 6.1× |
| Goccy | 197455 | 644.58 MB/s | 225244 | 884 | 5.6× |
| Easyjson | 211135 | 602.81 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 418017 | 304.47 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 447539 | 211.50 MB/s | 465585 | 9714 | 2.5× |
| Stdlib | 1106334 | 115.04 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3759154 | 598.82 MB/s | 3122875 | 3081 | 6.9× |
| Sonic | 4512396 | 498.86 MB/s | 15233773 | 970 | 5.8× |
| SonicFastest | 4523594 | 497.62 MB/s | 15236989 | 970 | 5.8× |
| Goccy | 10392797 | 216.60 MB/s | 4119841 | 56532 | 2.5× |
| Easyjson | 10957594 | 205.43 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11559440 | 194.73 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16095749 | 139.85 MB/s | 3123214 | 3083 | 1.6× |
| Stdlib | 26093377 | 86.27 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 485233 | 557.26 MB/s | 348024 | 1627 | 6.9× |
| Sonic | 632159 | 427.75 MB/s | 484401 | 968 | 5.3× |
| SonicFastest | 636846 | 424.60 MB/s | 493292 | 968 | 5.3× |
| Easyjson | 1401622 | 192.92 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1416498 | 190.90 MB/s | 543222 | 8122 | 2.4× |
| LightningDecodeAny | 1570040 | 172.23 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2104008 | 128.52 MB/s | 348157 | 1628 | 1.6× |
| Stdlib | 3350038 | 80.72 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1359721 | 1270.26 MB/s | 959890 | 5882 | 9.7× |
| SonicFastest | 2034755 | 848.85 MB/s | 2722518 | 4020 | 6.5× |
| Sonic | 2034798 | 848.83 MB/s | 2728951 | 4020 | 6.5× |
| Goccy | 2346314 | 736.14 MB/s | 2582768 | 14604 | 5.6× |
| Easyjson | 4220177 | 409.27 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4355702 | 396.54 MB/s | 1011637 | 7594 | 3.0× |
| LightningDecodeAny | 4601525 | 108.72 MB/s | 4976205 | 81466 | 2.9× |
| Stdlib | 13253856 | 130.32 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1188 | 1524.93 MB/s | 0 | 0 | 11.8× |
| Easyjson | 2548 | 711.26 MB/s | 24 | 1 | 5.5× |
| Goccy | 2762 | 656.04 MB/s | 2608 | 4 | 5.1× |
| Sonic | 5905 | 306.87 MB/s | 3706 | 40 | 2.4× |
| SonicFastest | 5923 | 305.90 MB/s | 3740 | 40 | 2.4× |
| JSONV2 | 7732 | 234.35 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8101 | 223.56 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14053 | 128.94 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1216 | 1489.66 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2561 | 707.63 MB/s | 24 | 1 | 5.5× |
| Goccy | 2770 | 654.19 MB/s | 2608 | 4 | 5.1× |
| Sonic | 5913 | 306.44 MB/s | 3708 | 40 | 2.4× |
| SonicFastest | 5928 | 305.69 MB/s | 3714 | 40 | 2.4× |
| JSONV2 | 7814 | 231.90 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8100 | 223.58 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14070 | 128.78 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1393 | 1300.36 MB/s | 144 | 10 | 10.1× |
| Easyjson | 2741 | 661.14 MB/s | 144 | 10 | 5.1× |
| Goccy | 2886 | 627.82 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6095 | 297.27 MB/s | 3727 | 42 | 2.3× |
| SonicFastest | 6139 | 295.14 MB/s | 3745 | 42 | 2.3× |
| JSONV2 | 7949 | 227.94 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8092 | 223.79 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14029 | 129.17 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 773 | 638.74 MB/s | 160 | 1 | 7.1× |
| Sonic | 1230 | 401.58 MB/s | 989 | 6 | 4.5× |
| SonicFastest | 1235 | 399.87 MB/s | 988 | 6 | 4.5× |
| LightningDecodeAny | 1653 | 298.28 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2202 | 224.37 MB/s | 448 | 3 | 2.5× |
| Goccy | 2416 | 204.47 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3265 | 151.29 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5515 | 89.57 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 454 | 506.08 MB/s | 160 | 1 | 9.2× |
| Sonic | 881 | 261.04 MB/s | 654 | 6 | 4.7× |
| SonicFastest | 884 | 260.31 MB/s | 658 | 6 | 4.7× |
| LightningDecodeAny | 1355 | 168.95 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1386 | 165.96 MB/s | 448 | 3 | 3.0× |
| Goccy | 1577 | 145.87 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2458 | 93.58 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4174 | 55.10 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 74912 | 869.44 MB/s | 172432 | 107 | 7.3× |
| SonicFastest | 102455 | 635.71 MB/s | 162223 | 75 | 5.3× |
| Sonic | 102469 | 635.63 MB/s | 162309 | 75 | 5.3× |
| Goccy | 153702 | 423.76 MB/s | 228671 | 134 | 3.5× |
| LightningDecodeAny | 186922 | 285.30 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 230854 | 282.13 MB/s | 206656 | 607 | 2.4× |
| Stdlib | 545550 | 119.39 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2673519 | 725.81 MB/s | 2846866 | 2238 | 8.8× |
| SonicFastest | 4668021 | 415.69 MB/s | 14606972 | 1407 | 5.0× |
| Sonic | 4721917 | 410.95 MB/s | 14611810 | 1407 | 5.0× |
| Goccy | 4751103 | 408.43 MB/s | 4064053 | 13510 | 4.9× |
| Easyjson | 7560453 | 256.66 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9617997 | 201.75 MB/s | 7013524 | 219937 | 2.4× |
| JSONV2 | 11201481 | 173.23 MB/s | 3237219 | 13947 | 2.1× |
| Stdlib | 23452454 | 82.74 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1731310 | 1922.15 MB/s | 2488905 | 2995 | 12.0× |
| SonicFastest | 2649119 | 1256.20 MB/s | 6408654 | 4248 | 7.8× |
| Sonic | 2652628 | 1254.54 MB/s | 6480115 | 4248 | 7.8× |
| LightningDecodeAny | 3680487 | 835.15 MB/s | 4886621 | 56892 | 5.6× |
| Goccy | 4508705 | 738.09 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7370489 | 451.51 MB/s | 5364515 | 13243 | 2.8× |
| Stdlib | 20759519 | 160.30 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 226336 | 973.53 MB/s | 136896 | 228 | 9.0× |
| Sonic | 368809 | 597.45 MB/s | 295744 | 398 | 5.5× |
| SonicFastest | 369137 | 596.92 MB/s | 296298 | 398 | 5.5× |
| Goccy | 424195 | 519.45 MB/s | 363785 | 1067 | 4.8× |
| Easyjson | 547628 | 402.36 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 723615 | 304.51 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 856322 | 126.49 MB/s | 861393 | 11944 | 2.4× |
| Stdlib | 2037276 | 108.16 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13181490 | 614.50 MB/s | 15059842 | 51649 | 6.8× |
| Sonic | 16823285 | 481.48 MB/s | 70887392 | 40014 | 5.3× |
| SonicFastest | 16829847 | 481.29 MB/s | 70887434 | 40014 | 5.3× |
| Goccy | 23136290 | 350.10 MB/s | 16906501 | 107148 | 3.9× |
| Easyjson | 30893180 | 262.20 MB/s | 15059620 | 41643 | 2.9× |
| LightningDecodeAny | 42399036 | 122.72 MB/s | 34333350 | 912810 | 2.1× |
| JSONV2 | 43457705 | 186.39 MB/s | 15233761 | 78972 | 2.1× |
| Stdlib | 89222893 | 90.78 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6184300 | 482.43 MB/s | 3985336 | 30015 | 7.6× |
| SonicFastest | 8596151 | 347.07 MB/s | 26544716 | 56760 | 5.5× |
| Sonic | 8596809 | 347.04 MB/s | 26684047 | 56760 | 5.5× |
| Goccy | 16625330 | 179.45 MB/s | 10640448 | 273650 | 2.8× |
| Easyjson | 16802256 | 177.56 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 19145045 | 95.81 MB/s | 20023838 | 408557 | 2.5× |
| JSONV2 | 25108894 | 118.82 MB/s | 9257186 | 86278 | 1.9× |
| Stdlib | 47157050 | 63.27 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1359502 | 532.25 MB/s | 907388 | 3618 | 8.5× |
| SonicFastest | 1732281 | 417.71 MB/s | 3167655 | 7226 | 6.7× |
| Sonic | 1740917 | 415.64 MB/s | 3172294 | 7226 | 6.6× |
| LightningDecodeAny | 4244704 | 153.27 MB/s | 5752201 | 80172 | 2.7× |
| Easyjson | 4245847 | 170.42 MB/s | 2847906 | 3698 | 2.7× |
| Goccy | 4681260 | 154.57 MB/s | 2798804 | 80273 | 2.5× |
| JSONV2 | 5523478 | 131.00 MB/s | 2704603 | 7318 | 2.1× |
| Stdlib | 11544160 | 62.68 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1951583 | 808.24 MB/s | 907386 | 3618 | 8.0× |
| Sonic | 2214482 | 712.29 MB/s | 5794936 | 7226 | 7.1× |
| SonicFastest | 2240361 | 704.06 MB/s | 5790730 | 7226 | 7.0× |
| LightningDecodeAny | 3749636 | 200.93 MB/s | 5752203 | 80172 | 4.2× |
| Easyjson | 5610782 | 281.13 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5625539 | 280.39 MB/s | 3579675 | 80267 | 2.8× |
| JSONV2 | 6367023 | 247.74 MB/s | 2704591 | 7318 | 2.5× |
| Stdlib | 15668617 | 100.67 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222473 | 674.80 MB/s | 81920 | 1 | 8.2× |
| SonicFastest | 264143 | 568.34 MB/s | 248301 | 6 | 6.9× |
| Sonic | 270288 | 555.42 MB/s | 266795 | 6 | 6.7× |
| LightningDecodeAny | 468163 | 320.66 MB/s | 746004 | 10020 | 3.9× |
| Goccy | 853848 | 175.82 MB/s | 325657 | 10005 | 2.1× |
| JSONV2 | 1087809 | 138.01 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1816316 | 82.65 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 35048 | 802.25 MB/s | 30272 | 105 | 8.7× |
| Sonic | 62404 | 450.56 MB/s | 46642 | 103 | 4.9× |
| SonicFastest | 62461 | 450.15 MB/s | 46700 | 103 | 4.9× |
| Easyjson | 67443 | 416.90 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71291 | 394.40 MB/s | 59224 | 188 | 4.3× |
| JSONV2 | 134520 | 209.02 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 148564 | 189.26 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 304291 | 92.40 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1949 | 1194.73 MB/s | 32 | 1 | 11.7× |
| Goccy | 4102 | 567.57 MB/s | 3649 | 4 | 5.6× |
| Easyjson | 4253 | 547.33 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 4940 | 471.29 MB/s | 4253 | 6 | 4.6× |
| Sonic | 4962 | 469.14 MB/s | 4219 | 6 | 4.6× |
| JSONV2 | 8523 | 273.14 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10089 | 167.01 MB/s | 9960 | 195 | 2.3× |
| Stdlib | 22788 | 102.16 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220 | 860.48 MB/s | 0 | 0 | 11.3× |
| Goccy | 380 | 497.91 MB/s | 304 | 2 | 6.6× |
| Easyjson | 489 | 386.18 MB/s | 0 | 0 | 5.1× |
| SonicFastest | 796 | 237.56 MB/s | 497 | 4 | 3.1× |
| Sonic | 798 | 236.92 MB/s | 503 | 4 | 3.1× |
| JSONV2 | 1047 | 180.44 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1227 | 109.25 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2490 | 75.89 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1522 | 1439.10 MB/s | 0 | 0 | 10.6× |
| Goccy | 3181 | 688.84 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3200 | 684.60 MB/s | 24 | 1 | 5.0× |
| Sonic | 6310 | 347.20 MB/s | 4101 | 40 | 2.6× |
| SonicFastest | 6336 | 345.80 MB/s | 4130 | 40 | 2.5× |
| JSONV2 | 8127 | 269.58 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8442 | 214.53 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16133 | 135.81 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 717142 | 711.82 MB/s | 703781 | 1012 | 8.4× |
| Goccy | 1120747 | 455.48 MB/s | 1138017 | 5006 | 5.4× |
| Sonic | 1133490 | 450.36 MB/s | 883004 | 2006 | 5.3× |
| SonicFastest | 1136923 | 449.00 MB/s | 893916 | 2006 | 5.3× |
| Easyjson | 1514827 | 336.99 MB/s | 863778 | 3012 | 4.0× |
| JSONV2 | 3199100 | 159.57 MB/s | 1076010 | 12646 | 1.9× |
| LightningDecodeAny | 3488794 | 132.27 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6049763 | 84.38 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1334 | 14829.73 MB/s | 0 | 0 | 83.7× |
| Goccy | 19955 | 991.68 MB/s | 20491 | 2 | 5.6× |
| SonicFastest | 26635 | 742.97 MB/s | 21806 | 4 | 4.2× |
| Sonic | 26844 | 737.19 MB/s | 22245 | 4 | 4.2× |
| JSONV2 | 29673 | 666.90 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 75437 | 262.31 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 86030 | 230.02 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111692 | 177.18 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2818 | 6432.39 MB/s | 432 | 2 | 36.7× |
| Easyjson | 4000 | 4531.53 MB/s | 432 | 2 | 25.8× |
| SonicFastest | 9916 | 1827.70 MB/s | 24183 | 6 | 10.4× |
| Sonic | 9997 | 1813.03 MB/s | 24366 | 6 | 10.3× |
| Goccy | 15438 | 1174.02 MB/s | 19459 | 2 | 6.7× |
| LightningDecodeAny | 17317 | 1032.61 MB/s | 29088 | 191 | 6.0× |
| JSONV2 | 46110 | 393.06 MB/s | 16500 | 50 | 2.2× |
| Stdlib | 103283 | 175.48 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2281876 | 880.19 MB/s | 2962101 | 7417 | 8.1× |
| Goccy | 4190504 | 479.30 MB/s | 5412509 | 15832 | 4.4× |
| Sonic | 4332172 | 463.62 MB/s | 10952234 | 13683 | 4.3× |
| SonicFastest | 4365742 | 460.06 MB/s | 10966888 | 13683 | 4.2× |
| Easyjson | 4848293 | 414.27 MB/s | 2981484 | 7439 | 3.8× |
| JSONV2 | 6853206 | 293.07 MB/s | 3173686 | 14563 | 2.7× |
| LightningDecodeAny | 7271667 | 157.09 MB/s | 7386074 | 134751 | 2.5× |
| Stdlib | 18500361 | 108.57 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 868 | 632.65 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1894 | 289.33 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2135 | 257.09 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2640 | 207.96 MB/s | 1952 | 26 | 2.1× |
| SonicFastest | 2660 | 206.36 MB/s | 1971 | 26 | 2.1× |
| Goccy | 3019 | 181.86 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3297 | 166.50 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5629 | 97.53 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 539605 | 1170.33 MB/s | 413001 | 878 | 10.0× |
| Sonic | 973370 | 648.79 MB/s | 984397 | 1102 | 5.5× |
| SonicFastest | 986007 | 640.48 MB/s | 1006242 | 1102 | 5.4× |
| Easyjson | 1124318 | 561.69 MB/s | 422505 | 936 | 4.8× |
| Goccy | 1148554 | 549.83 MB/s | 985967 | 1201 | 4.7× |
| JSONV2 | 2132891 | 296.08 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2324218 | 200.89 MB/s | 2010201 | 30295 | 2.3× |
| Stdlib | 5370819 | 117.58 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 865231 | 650.01 MB/s | 767621 | 1254 | 6.1× |
| SonicFastest | 1007932 | 557.98 MB/s | 948883 | 1476 | 5.2× |
| Sonic | 1014074 | 554.60 MB/s | 961894 | 1476 | 5.2× |
| Goccy | 1336154 | 420.92 MB/s | 1039356 | 1029 | 3.9× |
| Easyjson | 1732225 | 324.67 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2680831 | 209.79 MB/s | 2114152 | 30295 | 2.0× |
| JSONV2 | 2735793 | 205.57 MB/s | 927439 | 3482 | 1.9× |
| Stdlib | 5249323 | 107.14 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 666375 | 800.12 MB/s | 368224 | 2293 | 8.2× |
| Easyjson | 1098555 | 485.34 MB/s | 428361 | 3273 | 5.0× |
| SonicFastest | 1149510 | 463.83 MB/s | 1051460 | 4351 | 4.8× |
| Sonic | 1150000 | 463.63 MB/s | 1062372 | 4351 | 4.8× |
| Goccy | 1307378 | 407.82 MB/s | 1167261 | 5409 | 4.2× |
| JSONV2 | 2539662 | 209.94 MB/s | 745450 | 13288 | 2.2× |
| LightningDecodeAny | 3350710 | 159.12 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5489742 | 97.12 MB/s | 798692 | 17133 | 1.0× |
