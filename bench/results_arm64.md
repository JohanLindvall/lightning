# JSON Deserialization Benchmarks

- generated 2026-07-02T14:32:54Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 105325 | 1208.40 MB/s | 49280 | 2 | 10.5× |
| Lightning | 105763 | 1203.40 MB/s | 49760 | 3 | 10.5× |
| Sonic | 185284 | 686.92 MB/s | 199777 | 10 | 6.0× |
| SonicFastest | 185412 | 686.44 MB/s | 199138 | 10 | 6.0× |
| Goccy | 198420 | 641.44 MB/s | 225492 | 884 | 5.6× |
| Easyjson | 211569 | 601.58 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 428832 | 296.79 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 468132 | 202.19 MB/s | 465730 | 9708 | 2.4× |
| Stdlib | 1107728 | 114.90 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3602723 | 624.82 MB/s | 3008240 | 1158 | 7.2× |
| Lightning | 3641195 | 618.22 MB/s | 3008242 | 1158 | 7.2× |
| Sonic | 4488717 | 501.49 MB/s | 15232101 | 970 | 5.8× |
| SonicFastest | 4610284 | 488.27 MB/s | 15233758 | 970 | 5.7× |
| Goccy | 10501162 | 214.36 MB/s | 4124347 | 56532 | 2.5× |
| Easyjson | 10990384 | 204.82 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12364332 | 182.06 MB/s | 19380209 | 223896 | 2.1× |
| JSONV2 | 16304853 | 138.06 MB/s | 3123214 | 3083 | 1.6× |
| Stdlib | 26106513 | 86.23 MB/s | 3123399 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 460036 | 587.79 MB/s | 392944 | 568 | 7.3× |
| LightningDestructive | 464356 | 582.32 MB/s | 392944 | 568 | 7.2× |
| Sonic | 629624 | 429.47 MB/s | 478321 | 968 | 5.3× |
| SonicFastest | 633569 | 426.79 MB/s | 484836 | 968 | 5.3× |
| Easyjson | 1424219 | 189.86 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1426662 | 189.54 MB/s | 544694 | 8123 | 2.4× |
| LightningDecodeAny | 1644536 | 164.43 MB/s | 2543877 | 29687 | 2.0× |
| JSONV2 | 2142200 | 126.23 MB/s | 348150 | 1628 | 1.6× |
| Stdlib | 3353701 | 80.63 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1186010 | 1456.31 MB/s | 767864 | 2798 | 11.2× |
| Lightning | 1204689 | 1433.73 MB/s | 767906 | 2799 | 11.0× |
| SonicFastest | 2040858 | 846.31 MB/s | 2756588 | 4020 | 6.5× |
| Sonic | 2051769 | 841.81 MB/s | 2775991 | 4020 | 6.5× |
| Goccy | 2380234 | 725.64 MB/s | 2582658 | 14604 | 5.6× |
| Easyjson | 4241417 | 407.22 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4251698 | 406.24 MB/s | 1011634 | 7594 | 3.1× |
| LightningDecodeAny | 4358932 | 114.78 MB/s | 4954733 | 76576 | 3.0× |
| Stdlib | 13289794 | 129.96 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1190 | 1523.01 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1206 | 1502.57 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2555 | 709.30 MB/s | 24 | 1 | 5.5× |
| Goccy | 2821 | 642.23 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5866 | 308.92 MB/s | 3701 | 40 | 2.4× |
| SonicFastest | 5897 | 307.25 MB/s | 3733 | 40 | 2.4× |
| JSONV2 | 7776 | 233.03 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8158 | 222.00 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14021 | 129.24 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1230 | 1472.83 MB/s | 0 | 0 | 11.4× |
| LightningDestructive | 1234 | 1468.58 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2555 | 709.30 MB/s | 24 | 1 | 5.5× |
| Goccy | 2869 | 631.60 MB/s | 2608 | 4 | 4.9× |
| Sonic | 5938 | 305.14 MB/s | 3764 | 40 | 2.4× |
| SonicFastest | 5944 | 304.85 MB/s | 3735 | 40 | 2.4× |
| JSONV2 | 7698 | 235.38 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8224 | 220.20 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14071 | 128.78 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1402 | 1292.40 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1456 | 1244.93 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2748 | 659.37 MB/s | 144 | 10 | 5.1× |
| Goccy | 2867 | 632.08 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6156 | 294.37 MB/s | 3784 | 42 | 2.3× |
| Sonic | 6163 | 294.02 MB/s | 3771 | 42 | 2.3× |
| JSONV2 | 7950 | 227.93 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8234 | 219.94 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14053 | 128.94 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 696 | 710.17 MB/s | 160 | 1 | 7.9× |
| Lightning | 698 | 708.09 MB/s | 160 | 1 | 7.9× |
| SonicFastest | 1243 | 397.58 MB/s | 979 | 6 | 4.4× |
| Sonic | 1248 | 395.96 MB/s | 979 | 6 | 4.4× |
| LightningDecodeAny | 1393 | 353.94 MB/s | 1296 | 26 | 4.0× |
| Easyjson | 2229 | 221.62 MB/s | 448 | 3 | 2.5× |
| Goccy | 2421 | 204.08 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3197 | 154.53 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5524 | 89.43 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 427 | 538.86 MB/s | 160 | 1 | 9.6× |
| LightningDestructive | 428 | 536.97 MB/s | 160 | 1 | 9.6× |
| Sonic | 875 | 262.89 MB/s | 655 | 6 | 4.7× |
| SonicFastest | 881 | 261.06 MB/s | 661 | 6 | 4.7× |
| LightningDecodeAny | 1126 | 203.32 MB/s | 1296 | 26 | 3.7× |
| Easyjson | 1399 | 164.39 MB/s | 448 | 3 | 2.9× |
| Goccy | 1589 | 144.73 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2383 | 96.50 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4114 | 55.91 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 70274 | 926.83 MB/s | 158660 | 100 | 7.8× |
| Lightning | 71223 | 914.48 MB/s | 164880 | 105 | 7.6× |
| SonicFastest | 98550 | 660.90 MB/s | 156061 | 75 | 5.5× |
| Sonic | 99218 | 656.45 MB/s | 157084 | 75 | 5.5× |
| Goccy | 145328 | 448.17 MB/s | 229080 | 134 | 3.7× |
| LightningDecodeAny | 185259 | 287.86 MB/s | 180224 | 3245 | 2.9× |
| JSONV2 | 224958 | 289.53 MB/s | 206650 | 607 | 2.4× |
| Stdlib | 544707 | 119.57 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2541634 | 763.47 MB/s | 2813904 | 1358 | 9.2× |
| Lightning | 2626869 | 738.70 MB/s | 2813907 | 1358 | 8.9× |
| Sonic | 4781564 | 405.82 MB/s | 14608698 | 1407 | 4.9× |
| Goccy | 4837752 | 401.11 MB/s | 4065223 | 13510 | 4.9× |
| SonicFastest | 5016520 | 386.82 MB/s | 14606973 | 1407 | 4.7× |
| Easyjson | 7468974 | 259.80 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9439086 | 205.58 MB/s | 7064788 | 218633 | 2.5× |
| JSONV2 | 11226187 | 172.85 MB/s | 3237225 | 13947 | 2.1× |
| Stdlib | 23481756 | 82.64 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1101242 | 3021.89 MB/s | 351704 | 1286 | 18.9× |
| Lightning | 1783280 | 1866.13 MB/s | 2488904 | 2995 | 11.7× |
| Sonic | 2690681 | 1236.80 MB/s | 6467403 | 4248 | 7.7× |
| SonicFastest | 2711973 | 1227.09 MB/s | 6408792 | 4248 | 7.7× |
| LightningDecodeAny | 3757600 | 818.01 MB/s | 4886621 | 56892 | 5.5× |
| Goccy | 4613808 | 721.28 MB/s | 3948908 | 3816 | 4.5× |
| JSONV2 | 7454543 | 446.42 MB/s | 5364523 | 13243 | 2.8× |
| Stdlib | 20811422 | 159.90 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221150 | 996.37 MB/s | 135872 | 226 | 9.2× |
| LightningDestructive | 223384 | 986.40 MB/s | 135872 | 226 | 9.1× |
| SonicFastest | 379683 | 580.34 MB/s | 303881 | 398 | 5.4× |
| Sonic | 394398 | 558.69 MB/s | 342896 | 398 | 5.2× |
| Goccy | 433630 | 508.14 MB/s | 364208 | 1067 | 4.7× |
| Easyjson | 546538 | 403.17 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 726629 | 303.24 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 857929 | 126.25 MB/s | 897521 | 11703 | 2.4× |
| Stdlib | 2039639 | 108.03 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12185905 | 664.71 MB/s | 15730640 | 20821 | 7.3× |
| Lightning | 12619993 | 641.84 MB/s | 15730641 | 20821 | 7.1× |
| SonicFastest | 16583123 | 488.45 MB/s | 70887311 | 40014 | 5.4× |
| Sonic | 16584507 | 488.41 MB/s | 70915528 | 40014 | 5.4× |
| Goccy | 23417365 | 345.90 MB/s | 17366248 | 107149 | 3.8× |
| Easyjson | 30925919 | 261.92 MB/s | 15059616 | 41643 | 2.9× |
| LightningDecodeAny | 37130474 | 140.13 MB/s | 46191127 | 747112 | 2.4× |
| JSONV2 | 43781605 | 185.01 MB/s | 15233740 | 78972 | 2.0× |
| Stdlib | 89328662 | 90.68 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5720688 | 521.52 MB/s | 3908872 | 29356 | 8.3× |
| Lightning | 5798485 | 514.53 MB/s | 3908873 | 29356 | 8.1× |
| Sonic | 8673393 | 343.98 MB/s | 26667922 | 56760 | 5.4× |
| SonicFastest | 8684266 | 343.55 MB/s | 26667367 | 56760 | 5.4× |
| Goccy | 16390298 | 182.03 MB/s | 10621791 | 273648 | 2.9× |
| Easyjson | 16464612 | 181.20 MB/s | 9479447 | 30115 | 2.9× |
| LightningDecodeAny | 16775484 | 109.34 MB/s | 23982395 | 351152 | 2.8× |
| JSONV2 | 24546194 | 121.54 MB/s | 9257173 | 86278 | 1.9× |
| Stdlib | 47235987 | 63.16 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1249368 | 579.17 MB/s | 907601 | 3618 | 9.3× |
| Lightning | 1260725 | 573.95 MB/s | 907596 | 3618 | 9.2× |
| Sonic | 1784029 | 405.60 MB/s | 3191865 | 7226 | 6.5× |
| SonicFastest | 1791868 | 403.82 MB/s | 3198857 | 7226 | 6.5× |
| LightningDecodeAny | 4216647 | 154.29 MB/s | 6500456 | 76546 | 2.7× |
| Easyjson | 4250291 | 170.25 MB/s | 2847907 | 3698 | 2.7× |
| Goccy | 4864078 | 148.76 MB/s | 2998352 | 80284 | 2.4× |
| JSONV2 | 5762454 | 125.57 MB/s | 2704653 | 7318 | 2.0× |
| Stdlib | 11583512 | 62.47 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1839356 | 857.56 MB/s | 907601 | 3618 | 8.5× |
| Lightning | 1891418 | 833.95 MB/s | 907594 | 3618 | 8.3× |
| Sonic | 2241957 | 703.56 MB/s | 5788033 | 7226 | 7.0× |
| SonicFastest | 2249522 | 701.19 MB/s | 5791050 | 7226 | 7.0× |
| LightningDecodeAny | 3898636 | 193.25 MB/s | 6500458 | 76546 | 4.0× |
| Goccy | 5599436 | 281.70 MB/s | 3566488 | 80266 | 2.8× |
| Easyjson | 5615709 | 280.88 MB/s | 2847905 | 3698 | 2.8× |
| JSONV2 | 6462882 | 244.06 MB/s | 2704592 | 7318 | 2.4× |
| Stdlib | 15714783 | 100.37 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 210791 | 712.19 MB/s | 81920 | 1 | 8.7× |
| Lightning | 210920 | 711.76 MB/s | 81920 | 1 | 8.7× |
| SonicFastest | 269786 | 556.46 MB/s | 240358 | 6 | 6.8× |
| Sonic | 271446 | 553.05 MB/s | 245331 | 6 | 6.7× |
| LightningDecodeAny | 472417 | 317.77 MB/s | 745764 | 10016 | 3.9× |
| Goccy | 872210 | 172.12 MB/s | 325546 | 10005 | 2.1× |
| JSONV2 | 1079195 | 139.11 MB/s | 357717 | 20 | 1.7× |
| Stdlib | 1828600 | 82.10 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32464 | 866.10 MB/s | 29216 | 103 | 9.3× |
| LightningDestructive | 32545 | 863.96 MB/s | 29088 | 101 | 9.3× |
| SonicFastest | 63375 | 443.66 MB/s | 46416 | 103 | 4.8× |
| Sonic | 63454 | 443.11 MB/s | 46536 | 103 | 4.8× |
| Easyjson | 67769 | 414.89 MB/s | 32304 | 138 | 4.5× |
| Goccy | 72452 | 388.08 MB/s | 59199 | 188 | 4.2× |
| JSONV2 | 134690 | 208.75 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 151126 | 186.05 MB/s | 140592 | 2643 | 2.0× |
| Stdlib | 302883 | 92.83 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1968 | 1183.00 MB/s | 32 | 1 | 11.5× |
| LightningDestructive | 2025 | 1149.50 MB/s | 32 | 1 | 11.2× |
| Goccy | 4111 | 566.26 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4203 | 553.84 MB/s | 192 | 2 | 5.4× |
| Sonic | 5111 | 455.47 MB/s | 4288 | 6 | 4.4× |
| SonicFastest | 5112 | 455.38 MB/s | 4320 | 6 | 4.4× |
| JSONV2 | 8437 | 275.94 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10353 | 162.75 MB/s | 10200 | 195 | 2.2× |
| Stdlib | 22694 | 102.58 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222 | 850.61 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 223 | 845.85 MB/s | 0 | 0 | 10.9× |
| Goccy | 387 | 488.82 MB/s | 304 | 2 | 6.3× |
| Easyjson | 489 | 386.75 MB/s | 0 | 0 | 5.0× |
| Sonic | 790 | 239.28 MB/s | 523 | 4 | 3.1× |
| SonicFastest | 794 | 238.02 MB/s | 518 | 4 | 3.1× |
| JSONV2 | 1029 | 183.59 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1274 | 105.15 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2435 | 77.62 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1572 | 1393.98 MB/s | 0 | 0 | 10.2× |
| Lightning | 1574 | 1391.80 MB/s | 0 | 0 | 10.2× |
| Easyjson | 3184 | 688.02 MB/s | 24 | 1 | 5.0× |
| Goccy | 3184 | 688.18 MB/s | 2864 | 4 | 5.0× |
| Sonic | 6371 | 343.90 MB/s | 3925 | 40 | 2.5× |
| SonicFastest | 6384 | 343.20 MB/s | 3969 | 40 | 2.5× |
| JSONV2 | 7956 | 275.38 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8251 | 219.50 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16010 | 136.85 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 699079 | 730.21 MB/s | 703298 | 1010 | 8.6× |
| Lightning | 737042 | 692.60 MB/s | 703300 | 1010 | 8.2× |
| Sonic | 1160847 | 439.74 MB/s | 899792 | 2006 | 5.2× |
| Goccy | 1161119 | 439.64 MB/s | 1140324 | 5007 | 5.2× |
| SonicFastest | 1164245 | 438.46 MB/s | 905249 | 2006 | 5.2× |
| Easyjson | 1553898 | 328.51 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3234806 | 157.81 MB/s | 1076012 | 12646 | 1.9× |
| LightningDecodeAny | 3439880 | 134.15 MB/s | 2929689 | 64018 | 1.8× |
| Stdlib | 6044302 | 84.46 MB/s | 1162119 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14642.67 MB/s | 0 | 0 | 95.9× |
| LightningDestructive | 1390 | 14239.75 MB/s | 0 | 0 | 93.2× |
| Goccy | 20526 | 964.11 MB/s | 20491 | 2 | 6.3× |
| SonicFastest | 27375 | 722.88 MB/s | 22559 | 4 | 4.7× |
| Sonic | 27471 | 720.35 MB/s | 22685 | 4 | 4.7× |
| JSONV2 | 29798 | 664.09 MB/s | 8 | 1 | 4.3× |
| LightningDecodeAny | 75583 | 261.80 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 86070 | 229.92 MB/s | 0 | 0 | 1.5× |
| Stdlib | 129579 | 152.72 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2667 | 6796.40 MB/s | 0 | 0 | 38.5× |
| Lightning | 2822 | 6423.24 MB/s | 432 | 2 | 36.4× |
| Easyjson | 3942 | 4598.14 MB/s | 432 | 2 | 26.1× |
| SonicFastest | 10067 | 1800.33 MB/s | 22975 | 6 | 10.2× |
| Sonic | 10071 | 1799.66 MB/s | 22948 | 6 | 10.2× |
| Goccy | 16171 | 1120.75 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 16931 | 1056.16 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 46434 | 390.32 MB/s | 16500 | 50 | 2.2× |
| Stdlib | 102694 | 176.49 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2397845 | 837.62 MB/s | 3089821 | 6821 | 7.8× |
| Lightning | 2455059 | 818.10 MB/s | 3091534 | 6827 | 7.6× |
| Goccy | 4281564 | 469.10 MB/s | 5412364 | 15831 | 4.4× |
| SonicFastest | 4473247 | 449.00 MB/s | 10912541 | 13683 | 4.2× |
| Sonic | 4476499 | 448.68 MB/s | 10905217 | 13683 | 4.2× |
| Easyjson | 4931601 | 407.27 MB/s | 2981482 | 7439 | 3.8× |
| JSONV2 | 6929956 | 289.83 MB/s | 3173691 | 14563 | 2.7× |
| LightningDecodeAny | 7456606 | 153.19 MB/s | 8498331 | 134008 | 2.5× |
| Stdlib | 18665825 | 107.60 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 879 | 624.29 MB/s | 480 | 1 | 6.4× |
| LightningDestructive | 882 | 622.67 MB/s | 480 | 1 | 6.4× |
| LightningDecodeAny | 1692 | 323.88 MB/s | 2021 | 46 | 3.3× |
| Easyjson | 2161 | 254.09 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2700 | 203.36 MB/s | 1918 | 26 | 2.1× |
| SonicFastest | 2714 | 202.28 MB/s | 1954 | 26 | 2.1× |
| Goccy | 3121 | 175.91 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3380 | 162.44 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5645 | 97.25 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 504353 | 1252.13 MB/s | 402729 | 545 | 10.7× |
| Lightning | 560714 | 1126.27 MB/s | 451257 | 857 | 9.6× |
| SonicFastest | 1025450 | 615.84 MB/s | 1011571 | 1102 | 5.3× |
| Sonic | 1028688 | 613.90 MB/s | 1018706 | 1102 | 5.2× |
| Easyjson | 1149655 | 549.31 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1162847 | 543.08 MB/s | 985174 | 1201 | 4.6× |
| JSONV2 | 2134773 | 295.82 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2393810 | 195.05 MB/s | 2077367 | 30126 | 2.3× |
| Stdlib | 5388155 | 117.20 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 670953 | 838.22 MB/s | 579337 | 429 | 7.9× |
| Lightning | 845285 | 665.35 MB/s | 802710 | 1235 | 6.2× |
| Sonic | 1035347 | 543.21 MB/s | 954969 | 1476 | 5.1× |
| SonicFastest | 1040242 | 540.65 MB/s | 961475 | 1476 | 5.1× |
| Goccy | 1333443 | 421.77 MB/s | 1035855 | 1029 | 4.0× |
| Easyjson | 1754646 | 320.52 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2671374 | 210.53 MB/s | 2181321 | 30126 | 2.0× |
| JSONV2 | 2743366 | 205.01 MB/s | 927443 | 3482 | 1.9× |
| Stdlib | 5273774 | 106.64 MB/s | 1011674 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 654802 | 814.26 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 675981 | 788.75 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1107496 | 481.43 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1152913 | 462.46 MB/s | 1035511 | 4351 | 4.7× |
| Sonic | 1157469 | 460.64 MB/s | 1049992 | 4351 | 4.7× |
| Goccy | 1290707 | 413.09 MB/s | 1167233 | 5409 | 4.2× |
| JSONV2 | 2546180 | 209.40 MB/s | 745447 | 13288 | 2.1× |
| LightningDecodeAny | 3436783 | 155.14 MB/s | 2991149 | 50076 | 1.6× |
| Stdlib | 5472766 | 97.42 MB/s | 798693 | 17133 | 1.0× |
