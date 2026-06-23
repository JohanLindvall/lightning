# JSON Deserialization Benchmarks

- generated 2026-06-23T10:13:40Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 102360 | 1243.41 MB/s | 49760 | 3 | 12.1× |
| LightningDestructive | 104209 | 1221.34 MB/s | 49280 | 2 | 11.9× |
| Easyjson | 231342 | 550.16 MB/s | 122864 | 14 | 5.3× |
| Goccy | 237691 | 535.46 MB/s | 224547 | 884 | 5.2× |
| SonicFastest | 255232 | 498.66 MB/s | 213885 | 15 | 4.8× |
| Sonic | 256560 | 496.08 MB/s | 214011 | 15 | 4.8× |
| JSONV2 | 391630 | 324.99 MB/s | 195128 | 1805 | 3.2× |
| LightningDecodeAny | 441000 | 214.63 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1237171 | 102.88 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4127547 | 545.37 MB/s | 3122873 | 3081 | 7.5× |
| Lightning | 4144270 | 543.17 MB/s | 3122874 | 3081 | 7.5× |
| SonicFastest | 5070435 | 443.96 MB/s | 4862742 | 2584 | 6.1× |
| Sonic | 5421866 | 415.18 MB/s | 4862617 | 2584 | 5.7× |
| LightningDecodeAny | 11517634 | 195.44 MB/s | 7938299 | 281383 | 2.7× |
| Goccy | 12100493 | 186.03 MB/s | 4150982 | 56533 | 2.6× |
| Easyjson | 13617280 | 165.31 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 17202143 | 130.86 MB/s | 3123188 | 3083 | 1.8× |
| Stdlib | 31055672 | 72.48 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 566760 | 477.10 MB/s | 348024 | 1627 | 7.2× |
| LightningDestructive | 570149 | 474.27 MB/s | 348024 | 1627 | 7.1× |
| Sonic | 746467 | 362.24 MB/s | 643141 | 1147 | 5.4× |
| SonicFastest | 765721 | 353.14 MB/s | 645709 | 1147 | 5.3× |
| LightningDecodeAny | 1629025 | 165.99 MB/s | 1011487 | 37901 | 2.5× |
| Goccy | 1714096 | 157.75 MB/s | 545415 | 8123 | 2.4× |
| Easyjson | 1744046 | 155.04 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2325355 | 116.28 MB/s | 348160 | 1628 | 1.7× |
| Stdlib | 4061302 | 66.58 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1319655 | 1308.83 MB/s | 959848 | 5881 | 12.7× |
| Lightning | 1344461 | 1284.68 MB/s | 959890 | 5882 | 12.5× |
| SonicFastest | 2127958 | 811.67 MB/s | 2692610 | 5547 | 7.9× |
| Sonic | 2137945 | 807.88 MB/s | 2692679 | 5547 | 7.9× |
| Goccy | 2457502 | 702.83 MB/s | 2581493 | 14603 | 6.8× |
| Easyjson | 3689797 | 468.10 MB/s | 972032 | 5389 | 4.6× |
| LightningDecodeAny | 4122005 | 121.37 MB/s | 4976203 | 81466 | 4.1× |
| JSONV2 | 4254461 | 405.97 MB/s | 1011614 | 7594 | 3.9× |
| Stdlib | 16789011 | 102.88 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1019 | 1777.64 MB/s | 0 | 0 | 14.5× |
| LightningDestructive | 1089 | 1663.88 MB/s | 0 | 0 | 13.6× |
| Easyjson | 2800 | 647.03 MB/s | 24 | 1 | 5.3× |
| Goccy | 3191 | 567.76 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6248 | 290.01 MB/s | 3346 | 38 | 2.4× |
| Sonic | 6443 | 281.24 MB/s | 3347 | 38 | 2.3× |
| JSONV2 | 7523 | 240.87 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8554 | 211.72 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14777 | 122.62 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1090 | 1662.84 MB/s | 0 | 0 | 13.2× |
| LightningDestructive | 1144 | 1583.55 MB/s | 0 | 0 | 12.6× |
| Easyjson | 2809 | 645.12 MB/s | 24 | 1 | 5.1× |
| Goccy | 3299 | 549.30 MB/s | 2608 | 4 | 4.4× |
| SonicFastest | 6112 | 296.46 MB/s | 3348 | 38 | 2.4× |
| Sonic | 6325 | 286.49 MB/s | 3347 | 38 | 2.3× |
| JSONV2 | 7500 | 241.61 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 8602 | 210.52 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14431 | 125.57 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1245 | 1455.39 MB/s | 144 | 10 | 11.7× |
| LightningDestructive | 1296 | 1398.65 MB/s | 144 | 10 | 11.3× |
| Easyjson | 2941 | 616.04 MB/s | 144 | 10 | 5.0× |
| Goccy | 3376 | 536.76 MB/s | 2600 | 5 | 4.3× |
| SonicFastest | 6174 | 293.50 MB/s | 3367 | 40 | 2.4× |
| Sonic | 6406 | 282.88 MB/s | 3368 | 40 | 2.3× |
| JSONV2 | 8076 | 224.37 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8577 | 211.14 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14589 | 124.20 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 742 | 665.51 MB/s | 160 | 1 | 8.2× |
| LightningDestructive | 749 | 659.72 MB/s | 160 | 1 | 8.1× |
| Sonic | 1227 | 402.64 MB/s | 1076 | 8 | 4.9× |
| SonicFastest | 1271 | 388.54 MB/s | 1076 | 8 | 4.8× |
| LightningDecodeAny | 1625 | 303.40 MB/s | 1536 | 30 | 3.7× |
| Easyjson | 2286 | 216.14 MB/s | 448 | 3 | 2.6× |
| Goccy | 2477 | 199.47 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 3186 | 155.06 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6057 | 81.56 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 466 | 493.99 MB/s | 160 | 1 | 9.1× |
| LightningDestructive | 470 | 489.67 MB/s | 160 | 1 | 9.0× |
| SonicFastest | 862 | 266.95 MB/s | 801 | 8 | 4.9× |
| Sonic | 869 | 264.69 MB/s | 801 | 8 | 4.9× |
| LightningDecodeAny | 1411 | 162.26 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1596 | 144.15 MB/s | 448 | 3 | 2.6× |
| Goccy | 1675 | 137.32 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2394 | 96.08 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4226 | 54.43 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 84403 | 771.68 MB/s | 172433 | 107 | 7.6× |
| LightningDestructive | 88005 | 740.09 MB/s | 166213 | 102 | 7.3× |
| Sonic | 151496 | 429.93 MB/s | 235895 | 65 | 4.3× |
| SonicFastest | 160301 | 406.31 MB/s | 236260 | 65 | 4.0× |
| LightningDecodeAny | 183827 | 290.10 MB/s | 176960 | 3252 | 3.5× |
| Goccy | 186739 | 348.79 MB/s | 228841 | 134 | 3.5× |
| JSONV2 | 250981 | 259.51 MB/s | 206667 | 607 | 2.6× |
| Stdlib | 645342 | 100.93 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2626944 | 738.68 MB/s | 2846865 | 2238 | 9.6× |
| Lightning | 2631705 | 737.34 MB/s | 2846867 | 2238 | 9.6× |
| SonicFastest | 5082315 | 381.81 MB/s | 4879763 | 1736 | 5.0× |
| Sonic | 5168368 | 375.45 MB/s | 4879772 | 1736 | 4.9× |
| Goccy | 5211320 | 372.36 MB/s | 4062992 | 13509 | 4.9× |
| Easyjson | 7617230 | 254.75 MB/s | 3871265 | 15043 | 3.3× |
| LightningDecodeAny | 9502792 | 204.20 MB/s | 7013525 | 219937 | 2.7× |
| JSONV2 | 11713012 | 165.67 MB/s | 3237191 | 13947 | 2.2× |
| Stdlib | 25307516 | 76.68 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1011389 | 3290.36 MB/s | 351704 | 1286 | 25.3× |
| Lightning | 1651718 | 2014.77 MB/s | 2488906 | 2995 | 15.5× |
| SonicFastest | 2271660 | 1464.93 MB/s | 5895725 | 4263 | 11.3× |
| Sonic | 2310415 | 1440.36 MB/s | 5895564 | 4263 | 11.1× |
| LightningDecodeAny | 3621873 | 848.67 MB/s | 4886621 | 56892 | 7.1× |
| Goccy | 6659747 | 499.69 MB/s | 3948912 | 3816 | 3.8× |
| JSONV2 | 8645451 | 384.92 MB/s | 5364500 | 13243 | 3.0× |
| Stdlib | 25594842 | 130.02 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 227736 | 967.55 MB/s | 136896 | 228 | 9.8× |
| LightningDestructive | 238825 | 922.63 MB/s | 136896 | 228 | 9.3× |
| Goccy | 465171 | 473.69 MB/s | 363518 | 1066 | 4.8× |
| Sonic | 524842 | 419.83 MB/s | 350058 | 262 | 4.2× |
| SonicFastest | 525367 | 419.41 MB/s | 350326 | 262 | 4.2× |
| Easyjson | 597615 | 368.71 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 655156 | 336.33 MB/s | 129746 | 470 | 3.4× |
| LightningDecodeAny | 942359 | 114.94 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2228666 | 98.87 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14795128 | 547.48 MB/s | 15059832 | 51649 | 6.8× |
| Lightning | 15142888 | 534.91 MB/s | 15059844 | 51649 | 6.7× |
| SonicFastest | 20980000 | 386.08 MB/s | 19851766 | 41640 | 4.8× |
| Sonic | 21044057 | 384.91 MB/s | 19852881 | 41640 | 4.8× |
| Goccy | 24762366 | 327.11 MB/s | 19095689 | 107156 | 4.1× |
| Easyjson | 35284417 | 229.56 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 42364095 | 122.82 MB/s | 34333350 | 912810 | 2.4× |
| JSONV2 | 45325321 | 178.71 MB/s | 15233696 | 78972 | 2.2× |
| Stdlib | 101162949 | 80.07 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6025713 | 495.12 MB/s | 3985336 | 30015 | 8.5× |
| Lightning | 6203577 | 480.93 MB/s | 3985337 | 30015 | 8.2× |
| SonicFastest | 9458091 | 315.44 MB/s | 9131878 | 57804 | 5.4× |
| Sonic | 9498392 | 314.10 MB/s | 9130875 | 57804 | 5.4× |
| Goccy | 17669191 | 168.85 MB/s | 10000408 | 273625 | 2.9× |
| Easyjson | 18071559 | 165.09 MB/s | 9479441 | 30115 | 2.8× |
| LightningDecodeAny | 19806442 | 92.61 MB/s | 20023836 | 408557 | 2.6× |
| JSONV2 | 25158082 | 118.59 MB/s | 9257072 | 86278 | 2.0× |
| Stdlib | 51066839 | 58.42 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1426695 | 507.18 MB/s | 907392 | 3618 | 8.6× |
| Lightning | 1462987 | 494.60 MB/s | 907387 | 3618 | 8.4× |
| SonicFastest | 2233644 | 323.95 MB/s | 2367465 | 3683 | 5.5× |
| Sonic | 2236607 | 323.52 MB/s | 2367367 | 3683 | 5.5× |
| Easyjson | 5260881 | 137.54 MB/s | 2847906 | 3698 | 2.3× |
| LightningDecodeAny | 5323038 | 122.22 MB/s | 5752203 | 80172 | 2.3× |
| Goccy | 5390303 | 134.24 MB/s | 2686056 | 80266 | 2.3× |
| JSONV2 | 6194999 | 116.80 MB/s | 2704707 | 7318 | 2.0× |
| Stdlib | 12332807 | 58.67 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2086142 | 756.11 MB/s | 907387 | 3618 | 8.5× |
| LightningDestructive | 2087629 | 755.57 MB/s | 907393 | 3618 | 8.5× |
| Sonic | 2452702 | 643.11 MB/s | 3222217 | 3683 | 7.2× |
| SonicFastest | 2456863 | 642.02 MB/s | 3222593 | 3683 | 7.2× |
| LightningDecodeAny | 4438412 | 169.75 MB/s | 5752203 | 80172 | 4.0× |
| Easyjson | 6317306 | 249.69 MB/s | 2847907 | 3698 | 2.8× |
| Goccy | 6609728 | 238.64 MB/s | 3474170 | 80261 | 2.7× |
| JSONV2 | 6728049 | 234.44 MB/s | 2704555 | 7318 | 2.6× |
| Stdlib | 17733295 | 88.95 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 237505 | 632.09 MB/s | 81920 | 1 | 9.4× |
| LightningDestructive | 245870 | 610.58 MB/s | 81920 | 1 | 9.0× |
| Sonic | 386245 | 388.68 MB/s | 407401 | 16 | 5.8× |
| SonicFastest | 395200 | 379.87 MB/s | 407436 | 16 | 5.6× |
| LightningDecodeAny | 592934 | 253.18 MB/s | 746005 | 10020 | 3.7× |
| Goccy | 1032138 | 145.45 MB/s | 325124 | 10005 | 2.2× |
| JSONV2 | 1125500 | 133.38 MB/s | 357727 | 20 | 2.0× |
| Stdlib | 2221256 | 67.59 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 31878 | 882.02 MB/s | 30272 | 105 | 9.8× |
| LightningDestructive | 32874 | 855.29 MB/s | 30144 | 103 | 9.5× |
| SonicFastest | 66910 | 420.22 MB/s | 59493 | 83 | 4.7× |
| Sonic | 67207 | 418.37 MB/s | 59486 | 83 | 4.6× |
| Goccy | 75418 | 372.81 MB/s | 59279 | 188 | 4.1× |
| Easyjson | 76122 | 369.37 MB/s | 32304 | 138 | 4.1× |
| JSONV2 | 125394 | 224.23 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 153643 | 183.00 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 311186 | 90.35 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1975 | 1178.53 MB/s | 32 | 1 | 12.0× |
| LightningDestructive | 2021 | 1152.06 MB/s | 32 | 1 | 11.7× |
| Goccy | 4672 | 498.24 MB/s | 3649 | 4 | 5.1× |
| Easyjson | 4810 | 484.03 MB/s | 192 | 2 | 4.9× |
| SonicFastest | 6135 | 379.49 MB/s | 3708 | 4 | 3.9× |
| Sonic | 6141 | 379.06 MB/s | 3711 | 4 | 3.9× |
| JSONV2 | 7818 | 297.78 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 9736 | 173.07 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 23705 | 98.21 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222 | 851.70 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 225 | 838.57 MB/s | 0 | 0 | 10.9× |
| Goccy | 397 | 476.58 MB/s | 304 | 2 | 6.2× |
| Easyjson | 562 | 336.15 MB/s | 0 | 0 | 4.4× |
| SonicFastest | 763 | 247.85 MB/s | 341 | 3 | 3.2× |
| Sonic | 769 | 245.86 MB/s | 341 | 3 | 3.2× |
| JSONV2 | 924 | 204.58 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1199 | 111.74 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2461 | 76.81 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1470 | 1490.64 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1518 | 1443.16 MB/s | 0 | 0 | 11.5× |
| Easyjson | 3520 | 622.52 MB/s | 24 | 1 | 4.9× |
| Goccy | 3565 | 614.65 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6701 | 326.95 MB/s | 3606 | 38 | 2.6× |
| Sonic | 6859 | 319.45 MB/s | 3601 | 38 | 2.5× |
| JSONV2 | 7703 | 284.43 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8699 | 208.20 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 17385 | 126.03 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 666307 | 766.13 MB/s | 703778 | 1012 | 9.7× |
| Lightning | 722971 | 706.08 MB/s | 703779 | 1012 | 8.9× |
| Goccy | 1306544 | 390.71 MB/s | 1136193 | 5006 | 4.9× |
| SonicFastest | 1386561 | 368.16 MB/s | 1308221 | 2014 | 4.7× |
| Sonic | 1400549 | 364.48 MB/s | 1307774 | 2014 | 4.6× |
| Easyjson | 1576874 | 323.73 MB/s | 863781 | 3012 | 4.1× |
| JSONV2 | 3210532 | 159.00 MB/s | 1075956 | 12645 | 2.0× |
| LightningDecodeAny | 3375824 | 136.70 MB/s | 2785928 | 66022 | 1.9× |
| Stdlib | 6454522 | 79.09 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 611 | 32381.01 MB/s | 0 | 0 | 231.4× |
| LightningDestructive | 879 | 22518.27 MB/s | 0 | 0 | 160.9× |
| SonicFastest | 6360 | 3111.57 MB/s | 21123 | 3 | 22.2× |
| Goccy | 25562 | 774.17 MB/s | 20492 | 2 | 5.5× |
| Sonic | 29642 | 667.61 MB/s | 20632 | 3 | 4.8× |
| JSONV2 | 35997 | 549.74 MB/s | 8 | 1 | 3.9× |
| LightningDecodeAny | 93599 | 211.41 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 106581 | 185.67 MB/s | 0 | 0 | 1.3× |
| Stdlib | 141387 | 139.96 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2238 | 8098.22 MB/s | 432 | 2 | 53.9× |
| LightningDestructive | 2339 | 7749.68 MB/s | 0 | 0 | 51.6× |
| Easyjson | 4701 | 3855.31 MB/s | 432 | 2 | 25.7× |
| Sonic | 9366 | 1935.12 MB/s | 20436 | 5 | 12.9× |
| SonicFastest | 10828 | 1673.76 MB/s | 20375 | 5 | 11.1× |
| LightningDecodeAny | 18279 | 978.28 MB/s | 29088 | 191 | 6.6× |
| Goccy | 23114 | 784.11 MB/s | 19460 | 2 | 5.2× |
| JSONV2 | 49106 | 369.08 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 120658 | 150.21 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2312844 | 868.41 MB/s | 2960388 | 7411 | 9.0× |
| Lightning | 2351775 | 854.03 MB/s | 2962101 | 7417 | 8.9× |
| Goccy | 4747473 | 423.07 MB/s | 5410392 | 15829 | 4.4× |
| Sonic | 5154227 | 389.68 MB/s | 5152390 | 7085 | 4.1× |
| SonicFastest | 5215462 | 385.10 MB/s | 5153798 | 7085 | 4.0× |
| Easyjson | 5563781 | 360.99 MB/s | 2981484 | 7439 | 3.8× |
| LightningDecodeAny | 6987415 | 163.48 MB/s | 7386073 | 134751 | 3.0× |
| JSONV2 | 7321246 | 274.34 MB/s | 3173676 | 14563 | 2.9× |
| Stdlib | 20906548 | 96.07 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 813 | 675.60 MB/s | 480 | 1 | 7.2× |
| LightningDestructive | 823 | 667.14 MB/s | 480 | 1 | 7.1× |
| LightningDecodeAny | 1894 | 289.26 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 2000 | 274.55 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2138 | 256.81 MB/s | 2263 | 8 | 2.7× |
| Sonic | 2184 | 251.42 MB/s | 2262 | 8 | 2.7× |
| Goccy | 2996 | 183.26 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3073 | 178.63 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 5852 | 93.81 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 480184 | 1315.15 MB/s | 364472 | 566 | 12.6× |
| Lightning | 545104 | 1158.52 MB/s | 413002 | 878 | 11.1× |
| SonicFastest | 1054122 | 599.09 MB/s | 1065701 | 814 | 5.7× |
| Sonic | 1054820 | 598.69 MB/s | 1065606 | 814 | 5.7× |
| Easyjson | 1207158 | 523.14 MB/s | 422504 | 936 | 5.0× |
| Goccy | 1383726 | 456.39 MB/s | 990916 | 1200 | 4.4× |
| JSONV2 | 2130519 | 296.41 MB/s | 571593 | 3144 | 2.8× |
| LightningDecodeAny | 2430401 | 192.11 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6032080 | 104.69 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 766827 | 733.42 MB/s | 544249 | 448 | 7.3× |
| Lightning | 960610 | 585.47 MB/s | 767622 | 1254 | 5.9× |
| SonicFastest | 1381085 | 407.22 MB/s | 1348447 | 1185 | 4.1× |
| Sonic | 1382295 | 406.87 MB/s | 1348711 | 1185 | 4.1× |
| Goccy | 1576614 | 356.72 MB/s | 1037616 | 1028 | 3.6× |
| Easyjson | 1994038 | 282.04 MB/s | 775155 | 1254 | 2.8× |
| LightningDecodeAny | 2963785 | 189.76 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 2998812 | 187.54 MB/s | 927403 | 3482 | 1.9× |
| Stdlib | 5629630 | 99.90 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 592900 | 899.27 MB/s | 333416 | 2084 | 10.0× |
| Lightning | 661076 | 806.53 MB/s | 368224 | 2293 | 9.0× |
| Easyjson | 1249756 | 426.63 MB/s | 428362 | 3273 | 4.7× |
| Sonic | 1397339 | 381.57 MB/s | 979818 | 3082 | 4.2× |
| Goccy | 1400191 | 380.79 MB/s | 1167064 | 5408 | 4.2× |
| SonicFastest | 1404717 | 379.56 MB/s | 979603 | 3082 | 4.2× |
| JSONV2 | 2606465 | 204.56 MB/s | 745426 | 13288 | 2.3× |
| LightningDecodeAny | 3333736 | 159.93 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 5926662 | 89.96 MB/s | 798693 | 17133 | 1.0× |
