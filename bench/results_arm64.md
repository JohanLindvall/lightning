# JSON Deserialization Benchmarks

- generated 2026-06-21T09:44:35Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104166 | 1221.85 MB/s | 49760 | 3 | 10.6× |
| SonicFastest | 179965 | 707.22 MB/s | 192488 | 10 | 6.1× |
| Sonic | 180867 | 703.70 MB/s | 193269 | 10 | 6.1× |
| Goccy | 195024 | 652.61 MB/s | 224993 | 884 | 5.7× |
| Easyjson | 210734 | 603.96 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 419016 | 303.75 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 445004 | 212.70 MB/s | 465585 | 9714 | 2.5× |
| Stdlib | 1104100 | 115.27 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3757584 | 599.07 MB/s | 3122875 | 3081 | 6.9× |
| Sonic | 4420937 | 509.18 MB/s | 15233756 | 970 | 5.9× |
| SonicFastest | 4463769 | 504.29 MB/s | 15235373 | 970 | 5.8× |
| Goccy | 10340749 | 217.69 MB/s | 4115290 | 56532 | 2.5× |
| Easyjson | 10959735 | 205.39 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11402465 | 197.42 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16091376 | 139.89 MB/s | 3123214 | 3083 | 1.6× |
| Stdlib | 26066990 | 86.36 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 484567 | 558.03 MB/s | 348024 | 1627 | 6.9× |
| SonicFastest | 630444 | 428.91 MB/s | 481167 | 968 | 5.3× |
| Sonic | 631078 | 428.48 MB/s | 485103 | 968 | 5.3× |
| Goccy | 1399743 | 193.18 MB/s | 542668 | 8122 | 2.4× |
| Easyjson | 1401975 | 192.87 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1567546 | 172.50 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2097964 | 128.89 MB/s | 348150 | 1628 | 1.6× |
| Stdlib | 3350897 | 80.70 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1361599 | 1268.51 MB/s | 959890 | 5882 | 9.7× |
| SonicFastest | 2040798 | 846.34 MB/s | 2721043 | 4020 | 6.5× |
| Sonic | 2043285 | 845.31 MB/s | 2752652 | 4020 | 6.5× |
| Goccy | 2337322 | 738.97 MB/s | 2583282 | 14605 | 5.7× |
| Easyjson | 4208185 | 410.44 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4369460 | 395.29 MB/s | 1011634 | 7594 | 3.0× |
| LightningDecodeAny | 4545722 | 110.06 MB/s | 4976205 | 81466 | 2.9× |
| Stdlib | 13217821 | 130.67 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1188 | 1525.64 MB/s | 0 | 0 | 11.8× |
| Easyjson | 2542 | 712.76 MB/s | 24 | 1 | 5.5× |
| Goccy | 2763 | 655.90 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5864 | 308.99 MB/s | 3720 | 40 | 2.4× |
| Sonic | 5870 | 308.69 MB/s | 3738 | 40 | 2.4× |
| JSONV2 | 7747 | 233.90 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8140 | 222.49 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14022 | 129.22 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1215 | 1491.22 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2566 | 706.08 MB/s | 24 | 1 | 5.5× |
| Goccy | 2753 | 658.19 MB/s | 2608 | 4 | 5.1× |
| Sonic | 5874 | 308.50 MB/s | 3721 | 40 | 2.4× |
| SonicFastest | 5895 | 307.39 MB/s | 3761 | 40 | 2.4× |
| JSONV2 | 7818 | 231.76 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8128 | 222.81 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14060 | 128.88 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1388 | 1305.04 MB/s | 144 | 10 | 10.1× |
| Easyjson | 2743 | 660.56 MB/s | 144 | 10 | 5.1× |
| Goccy | 2845 | 636.95 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6043 | 299.83 MB/s | 3754 | 42 | 2.3× |
| Sonic | 6070 | 298.52 MB/s | 3789 | 42 | 2.3× |
| JSONV2 | 7942 | 228.16 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8112 | 223.26 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14011 | 129.33 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 776 | 636.83 MB/s | 160 | 1 | 7.1× |
| Sonic | 1217 | 405.98 MB/s | 975 | 6 | 4.5× |
| SonicFastest | 1218 | 405.58 MB/s | 979 | 6 | 4.5× |
| LightningDecodeAny | 1639 | 300.78 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2204 | 224.15 MB/s | 448 | 3 | 2.5× |
| Goccy | 2422 | 203.96 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3295 | 149.93 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5509 | 89.66 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 456 | 504.04 MB/s | 160 | 1 | 9.0× |
| SonicFastest | 862 | 266.74 MB/s | 655 | 6 | 4.8× |
| Sonic | 869 | 264.66 MB/s | 656 | 6 | 4.7× |
| LightningDecodeAny | 1344 | 170.34 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1380 | 166.70 MB/s | 448 | 3 | 3.0× |
| Goccy | 1551 | 148.27 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2441 | 94.24 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4103 | 56.06 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 75621 | 861.30 MB/s | 172432 | 107 | 7.2× |
| SonicFastest | 96598 | 674.26 MB/s | 155076 | 75 | 5.6× |
| Sonic | 97529 | 667.82 MB/s | 155252 | 75 | 5.6× |
| Goccy | 141285 | 461.00 MB/s | 228951 | 134 | 3.8× |
| LightningDecodeAny | 183045 | 291.34 MB/s | 176960 | 3252 | 3.0× |
| JSONV2 | 220706 | 295.11 MB/s | 206650 | 607 | 2.5× |
| Stdlib | 543694 | 119.80 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2670301 | 726.69 MB/s | 2846867 | 2238 | 8.8× |
| Goccy | 4739837 | 409.40 MB/s | 4064356 | 13510 | 4.9× |
| Sonic | 5015511 | 386.89 MB/s | 14606973 | 1407 | 4.7× |
| SonicFastest | 5016795 | 386.80 MB/s | 14608712 | 1407 | 4.7× |
| Easyjson | 7527630 | 257.78 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9505000 | 204.15 MB/s | 7013524 | 219937 | 2.5× |
| JSONV2 | 11228702 | 172.81 MB/s | 3237229 | 13947 | 2.1× |
| Stdlib | 23452011 | 82.74 MB/s | 3551321 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1774477 | 1875.39 MB/s | 2488905 | 2995 | 11.7× |
| Sonic | 2675360 | 1243.88 MB/s | 6469657 | 4248 | 7.8× |
| SonicFastest | 2679223 | 1242.09 MB/s | 6416738 | 4248 | 7.8× |
| LightningDecodeAny | 3700693 | 830.59 MB/s | 4886621 | 56892 | 5.6× |
| Goccy | 4547984 | 731.72 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7398729 | 449.78 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20771440 | 160.21 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 229675 | 959.38 MB/s | 136896 | 228 | 8.9× |
| SonicFastest | 377779 | 583.27 MB/s | 304360 | 398 | 5.4× |
| Sonic | 383498 | 574.57 MB/s | 318463 | 398 | 5.3× |
| Goccy | 437534 | 503.61 MB/s | 365189 | 1067 | 4.7× |
| Easyjson | 547634 | 402.36 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 730189 | 301.77 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 877255 | 123.47 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2039009 | 108.07 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13289678 | 609.50 MB/s | 15059836 | 51649 | 6.7× |
| Sonic | 16954127 | 477.76 MB/s | 70887373 | 40014 | 5.3× |
| SonicFastest | 16974669 | 477.18 MB/s | 70887372 | 40014 | 5.3× |
| Goccy | 23223003 | 348.79 MB/s | 17034485 | 107148 | 3.8× |
| Easyjson | 30937133 | 261.82 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 40238733 | 129.30 MB/s | 34333353 | 912810 | 2.2× |
| JSONV2 | 43401394 | 186.63 MB/s | 15233736 | 78972 | 2.1× |
| Stdlib | 89161460 | 90.85 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6204421 | 480.86 MB/s | 3985337 | 30015 | 7.6× |
| Sonic | 8498342 | 351.06 MB/s | 26589864 | 56760 | 5.5× |
| SonicFastest | 8517616 | 350.27 MB/s | 26604051 | 56760 | 5.5× |
| Goccy | 16311426 | 182.91 MB/s | 10641262 | 273649 | 2.9× |
| Easyjson | 16635458 | 179.34 MB/s | 9479442 | 30115 | 2.8× |
| LightningDecodeAny | 18489280 | 99.20 MB/s | 20023839 | 408557 | 2.5× |
| JSONV2 | 24948125 | 119.59 MB/s | 9257146 | 86278 | 1.9× |
| Stdlib | 47034003 | 63.43 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1373489 | 526.83 MB/s | 907388 | 3618 | 8.4× |
| Sonic | 1742227 | 415.33 MB/s | 3166906 | 7226 | 6.6× |
| SonicFastest | 1746482 | 414.32 MB/s | 3175371 | 7226 | 6.6× |
| Easyjson | 4240871 | 170.62 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4263215 | 152.60 MB/s | 5752203 | 80172 | 2.7× |
| Goccy | 4719487 | 153.32 MB/s | 2777735 | 80272 | 2.5× |
| JSONV2 | 5614160 | 128.89 MB/s | 2704684 | 7318 | 2.1× |
| Stdlib | 11576588 | 62.51 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1963746 | 803.24 MB/s | 907386 | 3618 | 8.0× |
| SonicFastest | 2238094 | 704.78 MB/s | 5785945 | 7226 | 7.0× |
| Sonic | 2240619 | 703.98 MB/s | 5788835 | 7226 | 7.0× |
| LightningDecodeAny | 3777165 | 199.46 MB/s | 5752202 | 80172 | 4.2× |
| Easyjson | 5611043 | 281.12 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5618416 | 280.75 MB/s | 3590449 | 80268 | 2.8× |
| JSONV2 | 6300073 | 250.37 MB/s | 2704590 | 7318 | 2.5× |
| Stdlib | 15760898 | 100.08 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222793 | 673.83 MB/s | 81920 | 1 | 8.2× |
| SonicFastest | 269042 | 557.99 MB/s | 250408 | 6 | 6.8× |
| Sonic | 269222 | 557.62 MB/s | 251460 | 6 | 6.8× |
| LightningDecodeAny | 473708 | 316.91 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 864691 | 173.62 MB/s | 328941 | 10005 | 2.1× |
| JSONV2 | 1094601 | 137.15 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1820651 | 82.46 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 35884 | 783.56 MB/s | 30272 | 105 | 8.5× |
| Sonic | 64952 | 432.89 MB/s | 48365 | 103 | 4.7× |
| SonicFastest | 65267 | 430.80 MB/s | 48128 | 103 | 4.6× |
| Easyjson | 68419 | 410.95 MB/s | 32304 | 138 | 4.4× |
| Goccy | 73084 | 384.72 MB/s | 59232 | 188 | 4.2× |
| JSONV2 | 135841 | 206.98 MB/s | 36895 | 242 | 2.2× |
| LightningDecodeAny | 155889 | 180.37 MB/s | 135024 | 2678 | 1.9× |
| Stdlib | 303435 | 92.66 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1949 | 1194.67 MB/s | 32 | 1 | 11.6× |
| Goccy | 4096 | 568.31 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4220 | 551.65 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5048 | 461.15 MB/s | 4258 | 6 | 4.5× |
| Sonic | 5063 | 459.80 MB/s | 4265 | 6 | 4.5× |
| JSONV2 | 8517 | 273.32 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10115 | 166.59 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22641 | 102.82 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 861.30 MB/s | 0 | 0 | 11.1× |
| Goccy | 372 | 508.78 MB/s | 304 | 2 | 6.5× |
| Easyjson | 489 | 386.26 MB/s | 0 | 0 | 5.0× |
| Sonic | 790 | 239.14 MB/s | 495 | 4 | 3.1× |
| SonicFastest | 791 | 239.00 MB/s | 497 | 4 | 3.1× |
| JSONV2 | 1038 | 182.17 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1230 | 108.91 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2426 | 77.89 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1521 | 1440.47 MB/s | 0 | 0 | 10.5× |
| Goccy | 3197 | 685.26 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3199 | 684.91 MB/s | 24 | 1 | 5.0× |
| Sonic | 6304 | 347.58 MB/s | 3949 | 40 | 2.5× |
| SonicFastest | 6370 | 343.98 MB/s | 4003 | 40 | 2.5× |
| JSONV2 | 8099 | 270.52 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8109 | 223.32 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 16030 | 136.68 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 739114 | 690.66 MB/s | 703778 | 1012 | 8.2× |
| Goccy | 1135834 | 449.43 MB/s | 1138577 | 5006 | 5.4× |
| Sonic | 1156794 | 441.29 MB/s | 898533 | 2006 | 5.3× |
| SonicFastest | 1159110 | 440.40 MB/s | 899163 | 2006 | 5.3× |
| Easyjson | 1546751 | 330.03 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3224186 | 158.33 MB/s | 1076004 | 12646 | 1.9× |
| LightningDecodeAny | 3474211 | 132.83 MB/s | 2785927 | 66022 | 1.8× |
| Stdlib | 6085592 | 83.88 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1334 | 14828.89 MB/s | 0 | 0 | 83.7× |
| Goccy | 19907 | 994.09 MB/s | 20491 | 2 | 5.6× |
| Sonic | 27060 | 731.31 MB/s | 22305 | 4 | 4.1× |
| SonicFastest | 27314 | 724.49 MB/s | 22930 | 4 | 4.1× |
| JSONV2 | 29771 | 664.71 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 76818 | 257.60 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 86029 | 230.03 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111691 | 177.18 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2809 | 6452.14 MB/s | 432 | 2 | 36.5× |
| Easyjson | 3984 | 4549.57 MB/s | 432 | 2 | 25.8× |
| Sonic | 9734 | 1862.02 MB/s | 22652 | 6 | 10.5× |
| SonicFastest | 9845 | 1840.96 MB/s | 22902 | 6 | 10.4× |
| Goccy | 16061 | 1128.45 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 16486 | 1084.70 MB/s | 29088 | 191 | 6.2× |
| JSONV2 | 46528 | 389.53 MB/s | 16500 | 50 | 2.2× |
| Stdlib | 102653 | 176.56 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2332268 | 861.18 MB/s | 2962101 | 7417 | 8.0× |
| Goccy | 4209675 | 477.11 MB/s | 5413114 | 15836 | 4.4× |
| SonicFastest | 4364221 | 460.22 MB/s | 10978595 | 13683 | 4.3× |
| Sonic | 4365511 | 460.08 MB/s | 10963472 | 13683 | 4.3× |
| Easyjson | 4886005 | 411.07 MB/s | 2981519 | 7439 | 3.8× |
| JSONV2 | 6849500 | 293.23 MB/s | 3173692 | 14563 | 2.7× |
| LightningDecodeAny | 7166166 | 159.40 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18556903 | 108.23 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 868 | 632.65 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1887 | 290.47 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2143 | 256.18 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2660 | 206.42 MB/s | 1941 | 26 | 2.1× |
| SonicFastest | 2660 | 206.35 MB/s | 1942 | 26 | 2.1× |
| Goccy | 3023 | 181.61 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3306 | 166.06 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5641 | 97.33 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 546395 | 1155.78 MB/s | 413001 | 878 | 9.9× |
| SonicFastest | 1004064 | 628.96 MB/s | 1003176 | 1102 | 5.4× |
| Sonic | 1004068 | 628.96 MB/s | 1007653 | 1102 | 5.4× |
| Easyjson | 1134586 | 556.60 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1162367 | 543.30 MB/s | 985460 | 1201 | 4.6× |
| JSONV2 | 2150632 | 293.64 MB/s | 571613 | 3144 | 2.5× |
| LightningDecodeAny | 2343914 | 199.20 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5386351 | 117.24 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 874534 | 643.09 MB/s | 767621 | 1254 | 6.0× |
| Sonic | 1027980 | 547.10 MB/s | 952240 | 1476 | 5.1× |
| SonicFastest | 1053079 | 534.06 MB/s | 981411 | 1476 | 5.0× |
| Goccy | 1345233 | 418.07 MB/s | 1041294 | 1030 | 3.9× |
| Easyjson | 1738062 | 323.58 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2680396 | 209.82 MB/s | 2114152 | 30295 | 2.0× |
| JSONV2 | 2767032 | 203.25 MB/s | 927438 | 3482 | 1.9× |
| Stdlib | 5269176 | 106.74 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 669915 | 795.89 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1103265 | 483.27 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1146636 | 464.99 MB/s | 1039288 | 4351 | 4.8× |
| SonicFastest | 1150542 | 463.41 MB/s | 1055027 | 4351 | 4.7× |
| Goccy | 1285702 | 414.70 MB/s | 1167224 | 5409 | 4.2× |
| JSONV2 | 2518974 | 211.66 MB/s | 745448 | 13288 | 2.2× |
| LightningDecodeAny | 3346661 | 159.32 MB/s | 2674621 | 50596 | 1.6× |
| Stdlib | 5458903 | 97.67 MB/s | 798692 | 17133 | 1.0× |
