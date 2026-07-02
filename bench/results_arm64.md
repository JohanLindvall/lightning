# JSON Deserialization Benchmarks

- generated 2026-07-02T08:51:25Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104021 | 1223.55 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 104358 | 1219.60 MB/s | 49280 | 2 | 10.6× |
| Sonic | 176460 | 721.27 MB/s | 188209 | 10 | 6.3× |
| SonicFastest | 177720 | 716.16 MB/s | 190297 | 10 | 6.2× |
| Goccy | 194846 | 653.21 MB/s | 225049 | 884 | 5.7× |
| Easyjson | 212135 | 599.97 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 421721 | 301.80 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 458452 | 206.46 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1104624 | 115.22 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3611861 | 623.24 MB/s | 3008240 | 1158 | 7.2× |
| Lightning | 3645756 | 617.44 MB/s | 3008243 | 1158 | 7.2× |
| Sonic | 4506947 | 499.46 MB/s | 15233785 | 970 | 5.8× |
| SonicFastest | 4677904 | 481.21 MB/s | 15233790 | 970 | 5.6× |
| Goccy | 10387267 | 216.71 MB/s | 4118192 | 56532 | 2.5× |
| Easyjson | 10952002 | 205.54 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11740788 | 191.73 MB/s | 7938298 | 281383 | 2.2× |
| JSONV2 | 16165476 | 139.25 MB/s | 3123205 | 3083 | 1.6× |
| Stdlib | 26069787 | 86.35 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 457580 | 590.94 MB/s | 392944 | 568 | 7.3× |
| LightningDestructive | 464125 | 582.61 MB/s | 392944 | 568 | 7.2× |
| SonicFastest | 623774 | 433.50 MB/s | 469275 | 968 | 5.4× |
| Sonic | 626320 | 431.73 MB/s | 479635 | 968 | 5.3× |
| Goccy | 1393007 | 194.11 MB/s | 542561 | 8122 | 2.4× |
| Easyjson | 1398453 | 193.36 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1608092 | 168.15 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2087280 | 129.55 MB/s | 348149 | 1628 | 1.6× |
| Stdlib | 3348266 | 80.76 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1181672 | 1461.66 MB/s | 767864 | 2798 | 11.3× |
| Lightning | 1191158 | 1450.02 MB/s | 767907 | 2799 | 11.2× |
| Sonic | 2031257 | 850.31 MB/s | 2743496 | 4020 | 6.6× |
| SonicFastest | 2038129 | 847.45 MB/s | 2728735 | 4020 | 6.5× |
| Goccy | 2344808 | 736.61 MB/s | 2582894 | 14605 | 5.7× |
| Easyjson | 4234104 | 407.93 MB/s | 972032 | 5389 | 3.2× |
| JSONV2 | 4330773 | 398.82 MB/s | 1011632 | 7594 | 3.1× |
| LightningDecodeAny | 4696073 | 106.54 MB/s | 4976205 | 81466 | 2.8× |
| Stdlib | 13349241 | 129.39 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1195 | 1516.76 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1205 | 1503.25 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2525 | 717.69 MB/s | 24 | 1 | 5.6× |
| Goccy | 2866 | 632.22 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6103 | 296.90 MB/s | 3868 | 40 | 2.3× |
| SonicFastest | 6111 | 296.51 MB/s | 3888 | 40 | 2.3× |
| JSONV2 | 7784 | 232.78 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8579 | 211.09 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14088 | 128.62 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1226 | 1478.43 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1236 | 1465.85 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2520 | 718.91 MB/s | 24 | 1 | 5.6× |
| Goccy | 2830 | 640.32 MB/s | 2608 | 4 | 5.0× |
| Sonic | 6022 | 300.89 MB/s | 3722 | 40 | 2.3× |
| SonicFastest | 6045 | 299.74 MB/s | 3762 | 40 | 2.3× |
| JSONV2 | 7742 | 234.05 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8214 | 220.47 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14124 | 128.29 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1406 | 1288.36 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1454 | 1246.32 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2770 | 654.12 MB/s | 144 | 10 | 5.1× |
| Goccy | 2893 | 626.30 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6154 | 294.42 MB/s | 3834 | 42 | 2.3× |
| SonicFastest | 6163 | 294.00 MB/s | 3856 | 42 | 2.3× |
| JSONV2 | 7891 | 229.63 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8553 | 211.75 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14116 | 128.37 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 689 | 716.63 MB/s | 160 | 1 | 8.1× |
| LightningDestructive | 694 | 711.40 MB/s | 160 | 1 | 8.0× |
| SonicFastest | 1242 | 397.76 MB/s | 982 | 6 | 4.5× |
| Sonic | 1243 | 397.48 MB/s | 988 | 6 | 4.5× |
| LightningDecodeAny | 1642 | 300.18 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2270 | 217.63 MB/s | 448 | 3 | 2.5× |
| Goccy | 2428 | 203.46 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3257 | 151.68 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5571 | 88.67 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 421 | 546.87 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 421 | 546.45 MB/s | 160 | 1 | 9.9× |
| Sonic | 888 | 259.04 MB/s | 663 | 6 | 4.7× |
| SonicFastest | 889 | 258.78 MB/s | 665 | 6 | 4.7× |
| LightningDecodeAny | 1389 | 164.91 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1435 | 160.27 MB/s | 448 | 3 | 2.9× |
| Goccy | 1583 | 145.31 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2443 | 94.13 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4183 | 54.99 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 68815 | 946.48 MB/s | 158660 | 100 | 7.9× |
| Lightning | 69467 | 937.60 MB/s | 164880 | 105 | 7.8× |
| Sonic | 95674 | 680.77 MB/s | 155287 | 75 | 5.7× |
| SonicFastest | 96033 | 678.23 MB/s | 155681 | 75 | 5.7× |
| Goccy | 141253 | 461.10 MB/s | 229062 | 134 | 3.8× |
| LightningDecodeAny | 185822 | 286.99 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 221045 | 294.66 MB/s | 206653 | 607 | 2.5× |
| Stdlib | 542843 | 119.98 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2501264 | 775.80 MB/s | 2813904 | 1358 | 9.4× |
| Lightning | 2575485 | 753.44 MB/s | 2813907 | 1358 | 9.1× |
| SonicFastest | 4555060 | 426.00 MB/s | 14606972 | 1407 | 5.2× |
| Sonic | 4577707 | 423.90 MB/s | 14608544 | 1407 | 5.1× |
| Goccy | 4806622 | 403.71 MB/s | 4064139 | 13510 | 4.9× |
| Easyjson | 7510968 | 258.35 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9934173 | 195.33 MB/s | 7013523 | 219937 | 2.4× |
| JSONV2 | 11201961 | 173.23 MB/s | 3237224 | 13947 | 2.1× |
| Stdlib | 23459971 | 82.71 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1109646 | 2999.00 MB/s | 351704 | 1286 | 18.7× |
| Lightning | 1777780 | 1871.90 MB/s | 2488905 | 2995 | 11.6× |
| Sonic | 2627286 | 1266.64 MB/s | 6487741 | 4248 | 7.9× |
| SonicFastest | 2643695 | 1258.78 MB/s | 6408807 | 4248 | 7.8× |
| LightningDecodeAny | 3770850 | 815.14 MB/s | 4886622 | 56892 | 5.5× |
| Goccy | 4545822 | 732.06 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7298621 | 455.95 MB/s | 5364522 | 13243 | 2.8× |
| Stdlib | 20698504 | 160.78 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 218901 | 1006.60 MB/s | 135392 | 226 | 9.3× |
| LightningDestructive | 220240 | 1000.48 MB/s | 135392 | 226 | 9.3× |
| SonicFastest | 372212 | 591.99 MB/s | 294122 | 398 | 5.5× |
| Sonic | 374143 | 588.94 MB/s | 296702 | 398 | 5.4× |
| Goccy | 427382 | 515.57 MB/s | 364477 | 1067 | 4.8× |
| Easyjson | 550103 | 400.55 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 733534 | 300.39 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 876641 | 123.55 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2038365 | 108.10 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12107733 | 669.00 MB/s | 15730640 | 20821 | 7.4× |
| Lightning | 12576342 | 644.07 MB/s | 15730641 | 20821 | 7.1× |
| SonicFastest | 17165728 | 471.87 MB/s | 70902294 | 40014 | 5.2× |
| Sonic | 17336007 | 467.24 MB/s | 70887851 | 40014 | 5.1× |
| Goccy | 23111752 | 350.47 MB/s | 17069837 | 107148 | 3.9× |
| Easyjson | 30786548 | 263.10 MB/s | 15059620 | 41643 | 2.9× |
| LightningDecodeAny | 41085315 | 126.64 MB/s | 34333351 | 912810 | 2.2× |
| JSONV2 | 43624906 | 185.67 MB/s | 15233739 | 78972 | 2.0× |
| Stdlib | 89088049 | 90.92 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5576495 | 535.01 MB/s | 3908872 | 29356 | 8.4× |
| Lightning | 5712520 | 522.27 MB/s | 3908875 | 29356 | 8.2× |
| Sonic | 8714785 | 342.35 MB/s | 26457935 | 56760 | 5.4× |
| SonicFastest | 8736299 | 341.50 MB/s | 26502206 | 56760 | 5.4× |
| Easyjson | 16589079 | 179.85 MB/s | 9479440 | 30115 | 2.8× |
| Goccy | 16810876 | 177.47 MB/s | 10585807 | 273648 | 2.8× |
| LightningDecodeAny | 19392691 | 94.58 MB/s | 20023836 | 408557 | 2.4× |
| JSONV2 | 24424219 | 122.15 MB/s | 9257171 | 86278 | 1.9× |
| Stdlib | 47092172 | 63.35 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1222106 | 592.09 MB/s | 907601 | 3618 | 9.5× |
| Lightning | 1232244 | 587.22 MB/s | 907595 | 3618 | 9.4× |
| Sonic | 1771025 | 408.58 MB/s | 3184814 | 7226 | 6.5× |
| SonicFastest | 1774801 | 407.71 MB/s | 3182779 | 7226 | 6.5× |
| Easyjson | 4201074 | 172.24 MB/s | 2847905 | 3698 | 2.8× |
| LightningDecodeAny | 4336455 | 150.02 MB/s | 5752202 | 80172 | 2.7× |
| Goccy | 4857261 | 148.97 MB/s | 2897974 | 80278 | 2.4× |
| JSONV2 | 5961389 | 121.38 MB/s | 2704635 | 7318 | 1.9× |
| Stdlib | 11561454 | 62.59 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1834237 | 859.95 MB/s | 907601 | 3618 | 8.6× |
| Lightning | 1870459 | 843.30 MB/s | 907594 | 3618 | 8.4× |
| Sonic | 2219557 | 710.66 MB/s | 5784492 | 7226 | 7.1× |
| SonicFastest | 2225343 | 708.81 MB/s | 5792949 | 7226 | 7.1× |
| LightningDecodeAny | 3865450 | 194.91 MB/s | 5752202 | 80172 | 4.1× |
| Goccy | 5587327 | 282.31 MB/s | 3578343 | 80266 | 2.8× |
| Easyjson | 5595071 | 281.92 MB/s | 2847904 | 3698 | 2.8× |
| JSONV2 | 6513334 | 242.17 MB/s | 2704591 | 7318 | 2.4× |
| Stdlib | 15739091 | 100.22 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 208790 | 719.02 MB/s | 81920 | 1 | 8.7× |
| LightningDestructive | 208884 | 718.69 MB/s | 81920 | 1 | 8.7× |
| SonicFastest | 276088 | 543.75 MB/s | 261270 | 6 | 6.6× |
| Sonic | 277027 | 541.91 MB/s | 263671 | 6 | 6.6× |
| LightningDecodeAny | 478377 | 313.81 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 861745 | 174.21 MB/s | 324992 | 10004 | 2.1× |
| JSONV2 | 1101708 | 136.26 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1822791 | 82.36 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32209 | 872.94 MB/s | 29216 | 103 | 9.5× |
| LightningDestructive | 32389 | 868.09 MB/s | 29088 | 101 | 9.4× |
| SonicFastest | 62997 | 446.32 MB/s | 46609 | 103 | 4.8× |
| Sonic | 63067 | 445.83 MB/s | 46720 | 103 | 4.8× |
| Easyjson | 68723 | 409.14 MB/s | 32304 | 138 | 4.4× |
| Goccy | 71994 | 390.55 MB/s | 59218 | 188 | 4.2× |
| JSONV2 | 134705 | 208.73 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 150972 | 186.24 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 305017 | 92.18 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1951 | 1193.34 MB/s | 32 | 1 | 11.7× |
| LightningDestructive | 2010 | 1158.18 MB/s | 32 | 1 | 11.4× |
| Goccy | 4164 | 559.03 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4222 | 551.44 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5038 | 462.08 MB/s | 4236 | 6 | 4.5× |
| Sonic | 5074 | 458.85 MB/s | 4306 | 6 | 4.5× |
| JSONV2 | 8428 | 276.24 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10341 | 162.94 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22858 | 101.85 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221 | 856.20 MB/s | 0 | 0 | 11.4× |
| LightningDestructive | 222 | 851.66 MB/s | 0 | 0 | 11.4× |
| Goccy | 389 | 485.84 MB/s | 304 | 2 | 6.5× |
| Easyjson | 493 | 383.59 MB/s | 0 | 0 | 5.1× |
| SonicFastest | 798 | 236.82 MB/s | 493 | 4 | 3.2× |
| Sonic | 798 | 236.71 MB/s | 495 | 4 | 3.2× |
| JSONV2 | 1048 | 180.27 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1231 | 108.81 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2523 | 74.91 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1548 | 1415.25 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1560 | 1404.52 MB/s | 0 | 0 | 10.3× |
| Goccy | 3172 | 690.81 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3188 | 687.32 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6521 | 335.98 MB/s | 4085 | 40 | 2.5× |
| Sonic | 6559 | 334.06 MB/s | 4107 | 40 | 2.5× |
| JSONV2 | 8032 | 272.78 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8578 | 211.13 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16088 | 136.19 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 681794 | 748.72 MB/s | 703297 | 1010 | 8.9× |
| Lightning | 736020 | 693.56 MB/s | 703300 | 1010 | 8.3× |
| Goccy | 1154028 | 442.34 MB/s | 1136950 | 5006 | 5.3× |
| SonicFastest | 1157092 | 441.17 MB/s | 896435 | 2006 | 5.3× |
| Sonic | 1158111 | 440.78 MB/s | 896855 | 2006 | 5.2× |
| Easyjson | 1527044 | 334.29 MB/s | 863778 | 3012 | 4.0× |
| JSONV2 | 3190600 | 159.99 MB/s | 1076013 | 12646 | 1.9× |
| LightningDecodeAny | 3567084 | 129.37 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6076965 | 84.00 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1336 | 14814.20 MB/s | 0 | 0 | 83.5× |
| LightningDestructive | 1359 | 14556.14 MB/s | 0 | 0 | 82.1× |
| Goccy | 20530 | 963.89 MB/s | 20491 | 2 | 5.4× |
| SonicFastest | 27990 | 707.02 MB/s | 22231 | 4 | 4.0× |
| Sonic | 28017 | 706.32 MB/s | 22266 | 4 | 4.0× |
| JSONV2 | 29574 | 669.12 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 80377 | 246.19 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86063 | 229.94 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111534 | 177.43 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2666 | 6798.58 MB/s | 0 | 0 | 38.7× |
| Lightning | 2833 | 6396.71 MB/s | 432 | 2 | 36.5× |
| Easyjson | 3959 | 4577.53 MB/s | 432 | 2 | 26.1× |
| Sonic | 9657 | 1876.70 MB/s | 22880 | 6 | 10.7× |
| SonicFastest | 9730 | 1862.74 MB/s | 22988 | 6 | 10.6× |
| Goccy | 15737 | 1151.68 MB/s | 19459 | 2 | 6.6× |
| LightningDecodeAny | 16724 | 1069.22 MB/s | 29088 | 191 | 6.2× |
| JSONV2 | 44858 | 404.03 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103305 | 175.44 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2311471 | 868.92 MB/s | 3089821 | 6821 | 8.0× |
| Lightning | 2379058 | 844.24 MB/s | 3091533 | 6827 | 7.8× |
| Goccy | 4137181 | 485.47 MB/s | 5412462 | 15832 | 4.5× |
| SonicFastest | 4364958 | 460.14 MB/s | 10941994 | 13683 | 4.2× |
| Sonic | 4372850 | 459.31 MB/s | 10987437 | 13683 | 4.2× |
| Easyjson | 4862253 | 413.08 MB/s | 2981481 | 7439 | 3.8× |
| JSONV2 | 6795592 | 295.56 MB/s | 3173686 | 14563 | 2.7× |
| LightningDecodeAny | 7440949 | 153.51 MB/s | 7386073 | 134751 | 2.5× |
| Stdlib | 18536738 | 108.35 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 907 | 605.15 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 914 | 600.37 MB/s | 480 | 1 | 6.2× |
| LightningDecodeAny | 2050 | 267.29 MB/s | 2261 | 50 | 2.8× |
| Easyjson | 2227 | 246.47 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2717 | 202.04 MB/s | 2057 | 26 | 2.1× |
| SonicFastest | 2720 | 201.84 MB/s | 2066 | 26 | 2.1× |
| Goccy | 3069 | 178.86 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3328 | 164.94 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5710 | 96.15 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 494544 | 1276.96 MB/s | 400649 | 555 | 10.9× |
| Lightning | 555852 | 1136.12 MB/s | 449177 | 867 | 9.7× |
| Sonic | 991011 | 637.24 MB/s | 1000651 | 1102 | 5.4× |
| SonicFastest | 996556 | 633.70 MB/s | 1006439 | 1102 | 5.4× |
| Easyjson | 1142082 | 552.95 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1152469 | 547.97 MB/s | 986744 | 1201 | 4.7× |
| JSONV2 | 2134972 | 295.80 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2364674 | 197.45 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5386635 | 117.24 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 715266 | 786.29 MB/s | 577673 | 437 | 7.3× |
| Lightning | 900227 | 624.74 MB/s | 801047 | 1243 | 5.8× |
| SonicFastest | 1012850 | 555.27 MB/s | 915516 | 1476 | 5.2× |
| Sonic | 1013935 | 554.68 MB/s | 920762 | 1476 | 5.2× |
| Goccy | 1312317 | 428.56 MB/s | 1034656 | 1030 | 4.0× |
| Easyjson | 1733458 | 324.44 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2730257 | 205.99 MB/s | 2114152 | 30295 | 1.9× |
| JSONV2 | 2758111 | 203.91 MB/s | 927450 | 3482 | 1.9× |
| Stdlib | 5251620 | 107.09 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 651239 | 818.71 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 674300 | 790.71 MB/s | 368224 | 2293 | 8.2× |
| Easyjson | 1114789 | 478.28 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1131918 | 471.04 MB/s | 1040966 | 4351 | 4.9× |
| Sonic | 1138726 | 468.22 MB/s | 1044325 | 4351 | 4.8× |
| Goccy | 1297840 | 410.82 MB/s | 1167236 | 5409 | 4.2× |
| JSONV2 | 2531422 | 210.62 MB/s | 745451 | 13288 | 2.2× |
| LightningDecodeAny | 3390380 | 157.26 MB/s | 2674619 | 50596 | 1.6× |
| Stdlib | 5498248 | 96.97 MB/s | 798692 | 17133 | 1.0× |
