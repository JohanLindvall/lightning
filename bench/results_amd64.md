# JSON Deserialization Benchmarks

- generated 2026-06-23T14:28:08Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 102293 | 1244.22 MB/s | 49760 | 3 | 12.2× |
| LightningDestructive | 104774 | 1214.75 MB/s | 49280 | 2 | 11.9× |
| SonicFastest | 192821 | 660.07 MB/s | 214318 | 15 | 6.5× |
| Sonic | 195372 | 651.45 MB/s | 214583 | 15 | 6.4× |
| Goccy | 238797 | 532.99 MB/s | 225236 | 884 | 5.2× |
| Easyjson | 240451 | 529.32 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 392936 | 323.91 MB/s | 195130 | 1805 | 3.2× |
| LightningDecodeAny | 445742 | 212.35 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1251871 | 101.67 MB/s | 199697 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4153019 | 542.03 MB/s | 3122872 | 3081 | 7.5× |
| Lightning | 4174046 | 539.30 MB/s | 3122874 | 3081 | 7.5× |
| Sonic | 5610372 | 401.23 MB/s | 4864395 | 2584 | 5.6× |
| SonicFastest | 5740628 | 392.13 MB/s | 4863771 | 2584 | 5.4× |
| LightningDecodeAny | 11475725 | 196.16 MB/s | 7938299 | 281383 | 2.7× |
| Goccy | 12376217 | 181.89 MB/s | 4161973 | 56533 | 2.5× |
| Easyjson | 13607593 | 165.43 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 17306471 | 130.07 MB/s | 3123224 | 3083 | 1.8× |
| Stdlib | 31193653 | 72.16 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 562935 | 480.35 MB/s | 348024 | 1627 | 7.2× |
| LightningDestructive | 568412 | 475.72 MB/s | 348024 | 1627 | 7.1× |
| SonicFastest | 743431 | 363.72 MB/s | 642754 | 1147 | 5.4× |
| Sonic | 743887 | 363.50 MB/s | 643086 | 1147 | 5.4× |
| LightningDecodeAny | 1597720 | 169.24 MB/s | 1011487 | 37901 | 2.5× |
| Goccy | 1682826 | 160.68 MB/s | 542424 | 8122 | 2.4× |
| Easyjson | 1741184 | 155.30 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2270265 | 119.11 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4043032 | 66.88 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1310821 | 1317.65 MB/s | 959848 | 5881 | 12.8× |
| Lightning | 1348227 | 1281.09 MB/s | 959889 | 5882 | 12.4× |
| SonicFastest | 2033683 | 849.30 MB/s | 2693088 | 5547 | 8.2× |
| Sonic | 2039300 | 846.96 MB/s | 2693390 | 5547 | 8.2× |
| Goccy | 2423513 | 712.69 MB/s | 2581295 | 14603 | 6.9× |
| Easyjson | 3597375 | 480.13 MB/s | 972032 | 5389 | 4.7× |
| LightningDecodeAny | 4065555 | 123.06 MB/s | 4976203 | 81466 | 4.1× |
| JSONV2 | 4247048 | 406.68 MB/s | 1011613 | 7594 | 3.9× |
| Stdlib | 16747233 | 103.13 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1018 | 1780.57 MB/s | 0 | 0 | 14.3× |
| LightningDestructive | 1060 | 1709.15 MB/s | 0 | 0 | 13.7× |
| Easyjson | 2806 | 645.78 MB/s | 24 | 1 | 5.2× |
| Goccy | 3176 | 570.45 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6437 | 281.50 MB/s | 3349 | 38 | 2.3× |
| Sonic | 6653 | 272.35 MB/s | 3346 | 38 | 2.2× |
| JSONV2 | 7438 | 243.60 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8548 | 211.87 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14548 | 124.55 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1087 | 1666.48 MB/s | 0 | 0 | 13.4× |
| LightningDestructive | 1137 | 1594.36 MB/s | 0 | 0 | 12.8× |
| Easyjson | 2798 | 647.66 MB/s | 24 | 1 | 5.2× |
| Goccy | 3088 | 586.71 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 5914 | 306.40 MB/s | 3348 | 38 | 2.5× |
| Sonic | 6166 | 293.88 MB/s | 3349 | 38 | 2.4× |
| JSONV2 | 7457 | 243.01 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 8696 | 208.26 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14526 | 124.74 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1248 | 1452.25 MB/s | 144 | 10 | 11.7× |
| LightningDestructive | 1300 | 1394.20 MB/s | 144 | 10 | 11.2× |
| Easyjson | 2971 | 609.90 MB/s | 144 | 10 | 4.9× |
| Goccy | 3403 | 532.43 MB/s | 2600 | 5 | 4.3× |
| SonicFastest | 6232 | 290.75 MB/s | 3367 | 40 | 2.3× |
| Sonic | 6428 | 281.89 MB/s | 3367 | 40 | 2.3× |
| JSONV2 | 8081 | 224.24 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8658 | 209.16 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14562 | 124.43 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 738 | 669.65 MB/s | 160 | 1 | 8.2× |
| LightningDestructive | 754 | 655.49 MB/s | 160 | 1 | 8.1× |
| SonicFastest | 1218 | 405.46 MB/s | 1076 | 8 | 5.0× |
| Sonic | 1237 | 399.37 MB/s | 1076 | 8 | 4.9× |
| LightningDecodeAny | 1654 | 298.06 MB/s | 1536 | 30 | 3.7× |
| Easyjson | 2369 | 208.53 MB/s | 448 | 3 | 2.6× |
| Goccy | 2438 | 202.67 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3113 | 158.67 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6077 | 81.28 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 467 | 492.67 MB/s | 160 | 1 | 9.0× |
| LightningDestructive | 470 | 488.93 MB/s | 160 | 1 | 9.0× |
| Sonic | 864 | 266.32 MB/s | 800 | 8 | 4.9× |
| SonicFastest | 864 | 266.06 MB/s | 801 | 8 | 4.9× |
| LightningDecodeAny | 1391 | 164.68 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1516 | 151.74 MB/s | 448 | 3 | 2.8× |
| Goccy | 1685 | 136.48 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2366 | 97.20 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4222 | 54.48 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 84498 | 770.81 MB/s | 172433 | 107 | 7.6× |
| LightningDestructive | 88161 | 738.79 MB/s | 166213 | 102 | 7.3× |
| SonicFastest | 153381 | 424.64 MB/s | 235816 | 65 | 4.2× |
| Sonic | 153957 | 423.05 MB/s | 235894 | 65 | 4.2× |
| Goccy | 176119 | 369.82 MB/s | 228485 | 134 | 3.7× |
| LightningDecodeAny | 183411 | 290.76 MB/s | 176960 | 3252 | 3.5× |
| JSONV2 | 242839 | 268.21 MB/s | 206663 | 607 | 2.7× |
| Stdlib | 645769 | 100.86 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2603986 | 745.19 MB/s | 2846864 | 2238 | 9.7× |
| Lightning | 2608181 | 743.99 MB/s | 2846867 | 2238 | 9.7× |
| Goccy | 4977109 | 389.88 MB/s | 4063221 | 13509 | 5.1× |
| Sonic | 5201209 | 373.08 MB/s | 4880327 | 1736 | 4.9× |
| SonicFastest | 5463924 | 355.14 MB/s | 4881174 | 1736 | 4.6× |
| Easyjson | 7732383 | 250.95 MB/s | 3871265 | 15043 | 3.3× |
| LightningDecodeAny | 9410926 | 206.19 MB/s | 7013524 | 219937 | 2.7× |
| JSONV2 | 11323250 | 171.37 MB/s | 3237191 | 13947 | 2.2× |
| Stdlib | 25246734 | 76.86 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1017740 | 3269.83 MB/s | 351704 | 1286 | 25.0× |
| Lightning | 1678129 | 1983.06 MB/s | 2488907 | 2995 | 15.2× |
| Sonic | 2249905 | 1479.10 MB/s | 5895558 | 4263 | 11.3× |
| SonicFastest | 2325670 | 1430.91 MB/s | 5892001 | 4263 | 10.9× |
| LightningDecodeAny | 3677365 | 835.86 MB/s | 4886621 | 56892 | 6.9× |
| Goccy | 6710772 | 495.89 MB/s | 3948909 | 3816 | 3.8× |
| JSONV2 | 8664668 | 384.07 MB/s | 5364508 | 13243 | 2.9× |
| Stdlib | 25457891 | 130.72 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 226551 | 972.61 MB/s | 136896 | 228 | 10.0× |
| LightningDestructive | 239943 | 918.33 MB/s | 136896 | 228 | 9.5× |
| Goccy | 462720 | 476.20 MB/s | 363902 | 1066 | 4.9× |
| SonicFastest | 508129 | 433.64 MB/s | 351296 | 262 | 4.5× |
| Sonic | 508784 | 433.08 MB/s | 351070 | 262 | 4.5× |
| Easyjson | 591123 | 372.76 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 654248 | 336.79 MB/s | 129746 | 470 | 3.5× |
| LightningDecodeAny | 958345 | 113.02 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2270199 | 97.06 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 15048810 | 538.25 MB/s | 15059832 | 51649 | 6.8× |
| Lightning | 15104843 | 536.25 MB/s | 15059837 | 51649 | 6.7× |
| Sonic | 20763874 | 390.10 MB/s | 19852581 | 41640 | 4.9× |
| SonicFastest | 20770841 | 389.97 MB/s | 19855207 | 41640 | 4.9× |
| Goccy | 24770539 | 327.00 MB/s | 19088691 | 107156 | 4.1× |
| Easyjson | 34766822 | 232.98 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 41452677 | 125.52 MB/s | 34333349 | 912810 | 2.5× |
| JSONV2 | 44462269 | 182.18 MB/s | 15233765 | 78972 | 2.3× |
| Stdlib | 101922147 | 79.47 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6034475 | 494.40 MB/s | 3985336 | 30015 | 8.4× |
| Lightning | 6188522 | 482.10 MB/s | 3985337 | 30015 | 8.2× |
| SonicFastest | 9196178 | 324.42 MB/s | 9127613 | 57804 | 5.5× |
| Sonic | 9298677 | 320.85 MB/s | 9127775 | 57804 | 5.4× |
| Goccy | 17745547 | 168.12 MB/s | 9798024 | 273617 | 2.9× |
| Easyjson | 17976372 | 165.97 MB/s | 9479441 | 30115 | 2.8× |
| LightningDecodeAny | 19619862 | 93.49 MB/s | 20023837 | 408557 | 2.6× |
| JSONV2 | 24942971 | 119.61 MB/s | 9257027 | 86278 | 2.0× |
| Stdlib | 50591567 | 58.97 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1373328 | 526.89 MB/s | 907392 | 3618 | 9.0× |
| Lightning | 1444453 | 500.95 MB/s | 907387 | 3618 | 8.5× |
| SonicFastest | 2177556 | 332.30 MB/s | 2367379 | 3683 | 5.7× |
| Sonic | 2196609 | 329.42 MB/s | 2367245 | 3683 | 5.6× |
| LightningDecodeAny | 5284275 | 123.11 MB/s | 5752204 | 80172 | 2.3× |
| Easyjson | 5295564 | 136.64 MB/s | 2847906 | 3698 | 2.3× |
| Goccy | 5427745 | 133.31 MB/s | 2708973 | 80267 | 2.3× |
| JSONV2 | 6138336 | 117.88 MB/s | 2704706 | 7318 | 2.0× |
| Stdlib | 12306282 | 58.80 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2075305 | 760.06 MB/s | 907387 | 3618 | 8.6× |
| LightningDestructive | 2078902 | 758.74 MB/s | 907392 | 3618 | 8.6× |
| Sonic | 2581556 | 611.01 MB/s | 3222463 | 3683 | 6.9× |
| SonicFastest | 2614416 | 603.33 MB/s | 3222768 | 3683 | 6.8× |
| LightningDecodeAny | 4457628 | 169.01 MB/s | 5752203 | 80172 | 4.0× |
| Easyjson | 6310693 | 249.95 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 6572974 | 239.98 MB/s | 3461937 | 80260 | 2.7× |
| JSONV2 | 6717268 | 234.82 MB/s | 2704555 | 7318 | 2.6× |
| Stdlib | 17798594 | 88.62 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 236137 | 635.75 MB/s | 81920 | 1 | 9.4× |
| LightningDestructive | 244429 | 614.18 MB/s | 81920 | 1 | 9.1× |
| Sonic | 393617 | 381.40 MB/s | 407368 | 16 | 5.6× |
| SonicFastest | 425122 | 353.13 MB/s | 407049 | 16 | 5.2× |
| LightningDecodeAny | 592878 | 253.21 MB/s | 746005 | 10020 | 3.7× |
| Goccy | 1044540 | 143.72 MB/s | 327145 | 10005 | 2.1× |
| JSONV2 | 1129509 | 132.91 MB/s | 357724 | 20 | 2.0× |
| Stdlib | 2218917 | 67.66 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 31846 | 882.90 MB/s | 30272 | 105 | 9.8× |
| LightningDestructive | 32756 | 858.37 MB/s | 30144 | 103 | 9.5× |
| Sonic | 67162 | 418.64 MB/s | 59459 | 83 | 4.6× |
| SonicFastest | 67272 | 417.96 MB/s | 59482 | 83 | 4.6× |
| Easyjson | 75288 | 373.46 MB/s | 32304 | 138 | 4.1× |
| Goccy | 76825 | 365.99 MB/s | 59262 | 188 | 4.1× |
| JSONV2 | 125400 | 224.22 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 153476 | 183.20 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 311148 | 90.37 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1956 | 1190.46 MB/s | 32 | 1 | 12.1× |
| LightningDestructive | 1997 | 1165.57 MB/s | 32 | 1 | 11.8× |
| Goccy | 4735 | 491.66 MB/s | 3649 | 4 | 5.0× |
| Easyjson | 4772 | 487.88 MB/s | 192 | 2 | 5.0× |
| SonicFastest | 6153 | 378.35 MB/s | 3711 | 4 | 3.8× |
| Sonic | 6299 | 369.58 MB/s | 3710 | 4 | 3.8× |
| JSONV2 | 7814 | 297.92 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 9597 | 175.58 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 23649 | 98.44 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222 | 850.17 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 226 | 837.18 MB/s | 0 | 0 | 10.9× |
| Goccy | 397 | 476.57 MB/s | 304 | 2 | 6.2× |
| Easyjson | 572 | 330.67 MB/s | 0 | 0 | 4.3× |
| SonicFastest | 768 | 246.23 MB/s | 341 | 3 | 3.2× |
| Sonic | 772 | 244.84 MB/s | 341 | 3 | 3.2× |
| JSONV2 | 939 | 201.36 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1179 | 113.61 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2456 | 76.96 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1471 | 1489.10 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1534 | 1428.48 MB/s | 0 | 0 | 11.4× |
| Easyjson | 3440 | 636.88 MB/s | 24 | 1 | 5.1× |
| Goccy | 3441 | 636.65 MB/s | 2864 | 4 | 5.1× |
| SonicFastest | 6563 | 333.87 MB/s | 3601 | 38 | 2.7× |
| Sonic | 6747 | 324.75 MB/s | 3600 | 38 | 2.6× |
| JSONV2 | 7710 | 284.18 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8673 | 208.82 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 17429 | 125.71 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 671234 | 760.50 MB/s | 703778 | 1012 | 9.6× |
| Lightning | 731869 | 697.50 MB/s | 703778 | 1012 | 8.8× |
| Goccy | 1284275 | 397.48 MB/s | 1130969 | 5006 | 5.0× |
| Easyjson | 1566538 | 325.86 MB/s | 863781 | 3012 | 4.1× |
| Sonic | 1600175 | 319.01 MB/s | 1309105 | 2014 | 4.0× |
| SonicFastest | 1603761 | 318.30 MB/s | 1309614 | 2014 | 4.0× |
| JSONV2 | 3126276 | 163.29 MB/s | 1075957 | 12645 | 2.1× |
| LightningDecodeAny | 3422968 | 134.81 MB/s | 2785928 | 66022 | 1.9× |
| Stdlib | 6464654 | 78.96 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 647 | 30593.62 MB/s | 0 | 0 | 220.3× |
| LightningDestructive | 897 | 22062.87 MB/s | 0 | 0 | 158.8× |
| SonicFastest | 7017 | 2820.12 MB/s | 21106 | 3 | 20.3× |
| Goccy | 26735 | 740.18 MB/s | 20492 | 2 | 5.3× |
| Sonic | 29220 | 677.24 MB/s | 20613 | 3 | 4.9× |
| JSONV2 | 36049 | 548.95 MB/s | 8 | 1 | 4.0× |
| LightningDecodeAny | 93821 | 210.91 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 111962 | 176.75 MB/s | 0 | 0 | 1.3× |
| Stdlib | 142465 | 138.90 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2232 | 8119.00 MB/s | 432 | 2 | 53.9× |
| LightningDestructive | 2403 | 7540.94 MB/s | 0 | 0 | 50.1× |
| Easyjson | 4680 | 3872.65 MB/s | 432 | 2 | 25.7× |
| Sonic | 9182 | 1973.85 MB/s | 20447 | 5 | 13.1× |
| SonicFastest | 9200 | 1969.93 MB/s | 20440 | 5 | 13.1× |
| LightningDecodeAny | 18346 | 974.71 MB/s | 29088 | 191 | 6.6× |
| Goccy | 20916 | 866.52 MB/s | 19460 | 2 | 5.8× |
| JSONV2 | 48236 | 375.74 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 120369 | 150.57 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2194928 | 915.06 MB/s | 2960388 | 7411 | 9.5× |
| Lightning | 2314341 | 867.85 MB/s | 2962101 | 7417 | 9.0× |
| Goccy | 4545278 | 441.89 MB/s | 5410649 | 15831 | 4.6× |
| SonicFastest | 4676237 | 429.51 MB/s | 5152977 | 7085 | 4.4× |
| Sonic | 4905163 | 409.47 MB/s | 5153790 | 7085 | 4.2× |
| Easyjson | 5367858 | 374.17 MB/s | 2981484 | 7439 | 3.9× |
| LightningDecodeAny | 6750719 | 169.21 MB/s | 7386073 | 134751 | 3.1× |
| JSONV2 | 7304192 | 274.98 MB/s | 3173673 | 14563 | 2.8× |
| Stdlib | 20762222 | 96.74 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 811 | 676.97 MB/s | 480 | 1 | 7.2× |
| LightningDestructive | 814 | 674.20 MB/s | 480 | 1 | 7.1× |
| LightningDecodeAny | 1877 | 292.02 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 1984 | 276.77 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2115 | 259.54 MB/s | 2261 | 8 | 2.8× |
| Sonic | 2154 | 254.88 MB/s | 2261 | 8 | 2.7× |
| Goccy | 2965 | 185.17 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3604 | 152.32 MB/s | 1664 | 7 | 1.6× |
| Stdlib | 5817 | 94.38 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 481828 | 1310.66 MB/s | 364472 | 566 | 12.4× |
| Lightning | 547968 | 1152.47 MB/s | 413001 | 878 | 10.9× |
| SonicFastest | 1079001 | 585.28 MB/s | 1066139 | 814 | 5.5× |
| Sonic | 1079760 | 584.87 MB/s | 1066160 | 814 | 5.5× |
| Easyjson | 1210761 | 521.58 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1404826 | 449.53 MB/s | 991798 | 1201 | 4.3× |
| JSONV2 | 2150937 | 293.60 MB/s | 571592 | 3144 | 2.8× |
| LightningDecodeAny | 2407802 | 193.91 MB/s | 2010198 | 30295 | 2.5× |
| Stdlib | 5978456 | 105.63 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 778130 | 722.77 MB/s | 544249 | 448 | 7.3× |
| Lightning | 963300 | 583.83 MB/s | 767623 | 1254 | 5.9× |
| Sonic | 1417783 | 396.68 MB/s | 1349980 | 1185 | 4.0× |
| SonicFastest | 1438836 | 390.88 MB/s | 1349809 | 1185 | 4.0× |
| Goccy | 1573305 | 357.47 MB/s | 1041252 | 1028 | 3.6× |
| Easyjson | 1996327 | 281.72 MB/s | 775155 | 1254 | 2.9× |
| LightningDecodeAny | 2980343 | 188.71 MB/s | 2114151 | 30295 | 1.9× |
| JSONV2 | 3027693 | 185.75 MB/s | 927406 | 3482 | 1.9× |
| Stdlib | 5692124 | 98.80 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 585603 | 910.48 MB/s | 333416 | 2084 | 10.3× |
| Lightning | 668974 | 797.01 MB/s | 368224 | 2293 | 9.0× |
| Easyjson | 1250452 | 426.39 MB/s | 428362 | 3273 | 4.8× |
| Sonic | 1349556 | 395.08 MB/s | 979942 | 3082 | 4.5× |
| SonicFastest | 1350905 | 394.68 MB/s | 979419 | 3082 | 4.5× |
| Goccy | 1416680 | 376.36 MB/s | 1167035 | 5408 | 4.3× |
| JSONV2 | 2652419 | 201.02 MB/s | 745419 | 13288 | 2.3× |
| LightningDecodeAny | 3369074 | 158.26 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6045986 | 88.19 MB/s | 798692 | 17133 | 1.0× |
