# JSON Deserialization Benchmarks

- generated 2026-07-01T12:54:37Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V45 96-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 58483 | 2176.26 MB/s | 49760 | 3 | 12.3× |
| LightningDestructive | 59804 | 2128.21 MB/s | 49280 | 2 | 12.0× |
| SonicFastest | 115734 | 1099.72 MB/s | 215837 | 15 | 6.2× |
| Sonic | 116292 | 1094.44 MB/s | 215713 | 15 | 6.2× |
| Easyjson | 124196 | 1024.79 MB/s | 122864 | 14 | 5.8× |
| Goccy | 140798 | 903.96 MB/s | 225706 | 884 | 5.1× |
| JSONV2 | 226043 | 563.06 MB/s | 195128 | 1805 | 3.2× |
| LightningDecodeAny | 264064 | 358.45 MB/s | 465586 | 9714 | 2.7× |
| Stdlib | 716601 | 177.61 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2228938 | 1009.92 MB/s | 3122873 | 3081 | 7.0× |
| Lightning | 2255820 | 997.89 MB/s | 3122873 | 3081 | 7.0× |
| Sonic | 2827473 | 796.14 MB/s | 4891933 | 2584 | 5.6× |
| SonicFastest | 2890054 | 778.90 MB/s | 4879674 | 2584 | 5.4× |
| LightningDecodeAny | 6191261 | 363.58 MB/s | 7938298 | 281383 | 2.5× |
| Goccy | 6484591 | 347.14 MB/s | 4258956 | 56538 | 2.4× |
| Easyjson | 6939894 | 324.36 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 8825664 | 255.06 MB/s | 3123189 | 3083 | 1.8× |
| Stdlib | 15712289 | 143.27 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 303363 | 891.35 MB/s | 348025 | 1627 | 6.8× |
| LightningDestructive | 308795 | 875.67 MB/s | 348025 | 1627 | 6.7× |
| Sonic | 420959 | 642.35 MB/s | 641709 | 1147 | 4.9× |
| SonicFastest | 422520 | 639.98 MB/s | 642034 | 1147 | 4.9× |
| LightningDecodeAny | 871420 | 310.30 MB/s | 1011486 | 37901 | 2.4× |
| Easyjson | 913001 | 296.17 MB/s | 330272 | 749 | 2.3× |
| Goccy | 914259 | 295.76 MB/s | 542145 | 8122 | 2.3× |
| JSONV2 | 1187936 | 227.62 MB/s | 348159 | 1628 | 1.7× |
| Stdlib | 2061557 | 131.16 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 737595 | 2341.67 MB/s | 959848 | 5881 | 12.1× |
| Lightning | 747261 | 2311.38 MB/s | 959889 | 5882 | 12.0× |
| Sonic | 1232356 | 1401.55 MB/s | 2704702 | 5547 | 7.3× |
| SonicFastest | 1237001 | 1396.28 MB/s | 2703864 | 5547 | 7.2× |
| Goccy | 1237811 | 1395.37 MB/s | 2580919 | 14603 | 7.2× |
| Easyjson | 1879738 | 918.85 MB/s | 972032 | 5389 | 4.8× |
| JSONV2 | 2153716 | 801.96 MB/s | 1011615 | 7594 | 4.2× |
| LightningDecodeAny | 2350056 | 212.89 MB/s | 4976204 | 81466 | 3.8× |
| Stdlib | 8952108 | 192.94 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 565 | 3205.28 MB/s | 0 | 0 | 15.6× |
| LightningDestructive | 608 | 2981.94 MB/s | 0 | 0 | 14.5× |
| Easyjson | 1548 | 1170.36 MB/s | 24 | 1 | 5.7× |
| Goccy | 1706 | 1062.24 MB/s | 2608 | 4 | 5.2× |
| SonicFastest | 3379 | 536.30 MB/s | 3343 | 38 | 2.6× |
| Sonic | 3487 | 519.65 MB/s | 3344 | 38 | 2.5× |
| JSONV2 | 4200 | 431.45 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 4767 | 379.90 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 8821 | 205.43 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 606 | 2989.41 MB/s | 0 | 0 | 14.4× |
| LightningDestructive | 636 | 2850.82 MB/s | 0 | 0 | 13.8× |
| Easyjson | 1462 | 1239.29 MB/s | 24 | 1 | 6.0× |
| Goccy | 1663 | 1089.33 MB/s | 2608 | 4 | 5.3× |
| SonicFastest | 3440 | 526.74 MB/s | 3344 | 38 | 2.5× |
| Sonic | 3568 | 507.82 MB/s | 3342 | 38 | 2.5× |
| JSONV2 | 4157 | 435.85 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 4779 | 378.97 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 8747 | 207.15 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 693 | 2614.30 MB/s | 144 | 10 | 12.5× |
| LightningDestructive | 722 | 2510.14 MB/s | 144 | 10 | 12.0× |
| Easyjson | 1581 | 1145.79 MB/s | 144 | 10 | 5.5× |
| Goccy | 1733 | 1045.67 MB/s | 2600 | 5 | 5.0× |
| SonicFastest | 3474 | 521.56 MB/s | 3360 | 40 | 2.5× |
| Sonic | 3590 | 504.70 MB/s | 3360 | 40 | 2.4× |
| JSONV2 | 4232 | 428.12 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 4778 | 379.06 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 8659 | 209.26 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 402 | 1227.25 MB/s | 160 | 1 | 8.1× |
| LightningDestructive | 409 | 1209.12 MB/s | 160 | 1 | 8.0× |
| Sonic | 716 | 690.08 MB/s | 1075 | 8 | 4.6× |
| SonicFastest | 720 | 685.70 MB/s | 1076 | 8 | 4.5× |
| LightningDecodeAny | 928 | 531.49 MB/s | 1536 | 30 | 3.5× |
| Easyjson | 1204 | 410.36 MB/s | 448 | 3 | 2.7× |
| Goccy | 1335 | 370.07 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 1693 | 291.79 MB/s | 528 | 7 | 1.9× |
| Stdlib | 3266 | 151.25 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 238 | 968.28 MB/s | 160 | 1 | 9.6× |
| LightningDestructive | 242 | 949.00 MB/s | 160 | 1 | 9.4× |
| SonicFastest | 488 | 471.72 MB/s | 801 | 8 | 4.7× |
| Sonic | 489 | 470.74 MB/s | 801 | 8 | 4.7× |
| Easyjson | 783 | 293.71 MB/s | 448 | 3 | 2.9× |
| LightningDecodeAny | 795 | 287.91 MB/s | 1536 | 30 | 2.9× |
| Goccy | 898 | 256.08 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 1245 | 184.80 MB/s | 528 | 7 | 1.8× |
| Stdlib | 2283 | 100.77 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 47358 | 1375.31 MB/s | 172434 | 107 | 7.7× |
| LightningDestructive | 48226 | 1350.56 MB/s | 166213 | 102 | 7.6× |
| Sonic | 87093 | 747.84 MB/s | 235902 | 65 | 4.2× |
| SonicFastest | 88332 | 737.35 MB/s | 235949 | 65 | 4.1× |
| Goccy | 95622 | 681.14 MB/s | 227620 | 134 | 3.8× |
| LightningDecodeAny | 103272 | 516.39 MB/s | 176961 | 3252 | 3.5× |
| JSONV2 | 131543 | 495.14 MB/s | 206660 | 607 | 2.8× |
| Stdlib | 364443 | 178.72 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1426741 | 1360.07 MB/s | 2846864 | 2238 | 9.6× |
| Lightning | 1443189 | 1344.57 MB/s | 2846866 | 2238 | 9.5× |
| Goccy | 2582866 | 751.29 MB/s | 4062857 | 13509 | 5.3× |
| Sonic | 2604356 | 745.09 MB/s | 4878528 | 1736 | 5.3× |
| SonicFastest | 2630092 | 737.80 MB/s | 4879716 | 1736 | 5.2× |
| Easyjson | 4132488 | 469.57 MB/s | 3871265 | 15043 | 3.3× |
| LightningDecodeAny | 5335113 | 363.72 MB/s | 7013525 | 219937 | 2.6× |
| JSONV2 | 5884096 | 329.78 MB/s | 3237186 | 13947 | 2.3× |
| Stdlib | 13686376 | 141.78 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 563073 | 5910.12 MB/s | 351704 | 1286 | 24.4× |
| Lightning | 873371 | 3810.33 MB/s | 2488905 | 2995 | 15.7× |
| SonicFastest | 1286314 | 2587.11 MB/s | 5896572 | 4263 | 10.7× |
| Sonic | 1308426 | 2543.39 MB/s | 5896642 | 4263 | 10.5× |
| LightningDecodeAny | 1870821 | 1643.00 MB/s | 4886622 | 56892 | 7.3× |
| Goccy | 3337409 | 997.13 MB/s | 3948914 | 3817 | 4.1× |
| JSONV2 | 4538941 | 733.17 MB/s | 5364499 | 13243 | 3.0× |
| Stdlib | 13727826 | 242.42 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 122789 | 1794.51 MB/s | 136896 | 228 | 10.2× |
| LightningDestructive | 126080 | 1747.67 MB/s | 136896 | 228 | 9.9× |
| Goccy | 247860 | 888.99 MB/s | 363966 | 1066 | 5.0× |
| Easyjson | 280042 | 786.83 MB/s | 130512 | 245 | 4.5× |
| SonicFastest | 318466 | 691.90 MB/s | 350294 | 262 | 3.9× |
| Sonic | 326621 | 674.62 MB/s | 350421 | 262 | 3.8× |
| JSONV2 | 345750 | 637.30 MB/s | 129747 | 470 | 3.6× |
| LightningDecodeAny | 530280 | 204.26 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 1246560 | 176.76 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 7367215 | 1099.47 MB/s | 15059833 | 51649 | 7.4× |
| Lightning | 7623998 | 1062.44 MB/s | 15059839 | 51649 | 7.1× |
| Sonic | 11399333 | 710.57 MB/s | 19868456 | 41640 | 4.8× |
| SonicFastest | 11426296 | 708.89 MB/s | 19869400 | 41640 | 4.8× |
| Goccy | 13458372 | 601.86 MB/s | 19086290 | 107156 | 4.0× |
| Easyjson | 17841553 | 454.00 MB/s | 15059616 | 41643 | 3.0× |
| JSONV2 | 23702415 | 341.74 MB/s | 15233697 | 78972 | 2.3× |
| LightningDecodeAny | 24206796 | 214.94 MB/s | 34333346 | 912810 | 2.2× |
| Stdlib | 54361532 | 149.00 MB/s | 15665071 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3263750 | 914.12 MB/s | 3985336 | 30015 | 8.5× |
| Lightning | 3315545 | 899.84 MB/s | 3985338 | 30015 | 8.4× |
| SonicFastest | 5048884 | 590.92 MB/s | 9133049 | 57804 | 5.5× |
| Sonic | 5086193 | 586.58 MB/s | 9133210 | 57804 | 5.4× |
| Easyjson | 9501055 | 314.01 MB/s | 9479441 | 30115 | 2.9× |
| Goccy | 9703851 | 307.45 MB/s | 9762467 | 273614 | 2.9× |
| LightningDecodeAny | 10674663 | 171.83 MB/s | 20023834 | 408557 | 2.6× |
| JSONV2 | 13441011 | 221.97 MB/s | 9257050 | 86278 | 2.1× |
| Stdlib | 27717065 | 107.64 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 746116 | 969.82 MB/s | 907393 | 3618 | 9.2× |
| Lightning | 774854 | 933.85 MB/s | 907387 | 3618 | 8.9× |
| Sonic | 1128119 | 641.42 MB/s | 2378187 | 3683 | 6.1× |
| SonicFastest | 1142675 | 633.25 MB/s | 2380432 | 3683 | 6.0× |
| LightningDecodeAny | 2743030 | 237.17 MB/s | 5752201 | 80172 | 2.5× |
| Easyjson | 2942778 | 245.89 MB/s | 2847907 | 3698 | 2.3× |
| Goccy | 3017305 | 239.82 MB/s | 2680931 | 80267 | 2.3× |
| JSONV2 | 3291587 | 219.83 MB/s | 2704711 | 7318 | 2.1× |
| Stdlib | 6866627 | 105.38 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1095463 | 1439.90 MB/s | 907388 | 3618 | 9.0× |
| LightningDestructive | 1124421 | 1402.81 MB/s | 907393 | 3618 | 8.8× |
| SonicFastest | 1510705 | 1044.12 MB/s | 3249989 | 3683 | 6.6× |
| Sonic | 1518461 | 1038.78 MB/s | 3249345 | 3683 | 6.5× |
| LightningDecodeAny | 2309821 | 326.17 MB/s | 5752198 | 80172 | 4.3× |
| Easyjson | 3376947 | 467.09 MB/s | 2847906 | 3698 | 2.9× |
| Goccy | 3631093 | 434.40 MB/s | 3490981 | 80262 | 2.7× |
| JSONV2 | 3678108 | 428.85 MB/s | 2704554 | 7318 | 2.7× |
| Stdlib | 9899992 | 159.33 MB/s | 2704547 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 127165 | 1180.54 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 130816 | 1147.60 MB/s | 81920 | 1 | 8.4× |
| Sonic | 222888 | 673.54 MB/s | 408050 | 16 | 4.9× |
| SonicFastest | 236786 | 634.01 MB/s | 409672 | 16 | 4.6× |
| LightningDecodeAny | 292747 | 512.80 MB/s | 746003 | 10020 | 3.8× |
| JSONV2 | 593013 | 253.15 MB/s | 357729 | 20 | 1.9× |
| Goccy | 600944 | 249.81 MB/s | 327652 | 10005 | 1.8× |
| Stdlib | 1099240 | 136.57 MB/s | 357802 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 18087 | 1554.55 MB/s | 30272 | 105 | 10.2× |
| LightningDestructive | 18436 | 1525.14 MB/s | 30144 | 103 | 10.0× |
| SonicFastest | 38154 | 736.93 MB/s | 59504 | 83 | 4.8× |
| Sonic | 38222 | 735.63 MB/s | 59504 | 83 | 4.8× |
| Easyjson | 40526 | 693.80 MB/s | 32304 | 138 | 4.5× |
| Goccy | 41478 | 677.87 MB/s | 59248 | 188 | 4.4× |
| JSONV2 | 74470 | 377.56 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 88151 | 318.96 MB/s | 135024 | 2678 | 2.1× |
| Stdlib | 184304 | 152.56 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1113 | 2091.27 MB/s | 32 | 1 | 12.3× |
| LightningDestructive | 1164 | 2000.49 MB/s | 32 | 1 | 11.7× |
| Goccy | 2457 | 947.62 MB/s | 3649 | 4 | 5.6× |
| Easyjson | 2526 | 921.59 MB/s | 192 | 2 | 5.4× |
| Sonic | 3846 | 605.28 MB/s | 3703 | 4 | 3.5× |
| SonicFastest | 3874 | 600.91 MB/s | 3705 | 4 | 3.5× |
| JSONV2 | 4188 | 555.82 MB/s | 1000 | 6 | 3.3× |
| LightningDecodeAny | 5545 | 303.86 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 13646 | 170.59 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 124 | 1527.11 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 126 | 1498.29 MB/s | 0 | 0 | 10.8× |
| Goccy | 205 | 923.74 MB/s | 304 | 2 | 6.7× |
| Easyjson | 288 | 656.45 MB/s | 0 | 0 | 4.7× |
| SonicFastest | 352 | 536.83 MB/s | 341 | 3 | 3.9× |
| Sonic | 358 | 527.98 MB/s | 341 | 3 | 3.8× |
| JSONV2 | 483 | 391.58 MB/s | 112 | 1 | 2.8× |
| LightningDecodeAny | 682 | 196.41 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 1361 | 138.82 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 846 | 2589.36 MB/s | 0 | 0 | 12.3× |
| LightningDestructive | 872 | 2514.08 MB/s | 0 | 0 | 11.9× |
| Easyjson | 1729 | 1267.46 MB/s | 24 | 1 | 6.0× |
| Goccy | 1833 | 1195.06 MB/s | 2864 | 4 | 5.7× |
| SonicFastest | 3750 | 584.26 MB/s | 3599 | 38 | 2.8× |
| Sonic | 3885 | 563.99 MB/s | 3598 | 38 | 2.7× |
| JSONV2 | 4155 | 527.27 MB/s | 640 | 6 | 2.5× |
| LightningDecodeAny | 4815 | 376.14 MB/s | 7536 | 158 | 2.2× |
| Stdlib | 10371 | 211.25 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 355156 | 1437.33 MB/s | 703777 | 1012 | 10.1× |
| Lightning | 376372 | 1356.31 MB/s | 703777 | 1012 | 9.5× |
| Goccy | 662149 | 770.94 MB/s | 1142104 | 5006 | 5.4× |
| Easyjson | 801225 | 637.12 MB/s | 863776 | 3012 | 4.5× |
| SonicFastest | 828321 | 616.28 MB/s | 1305206 | 2013 | 4.3× |
| Sonic | 834537 | 611.69 MB/s | 1304693 | 2013 | 4.3× |
| JSONV2 | 1757176 | 290.51 MB/s | 1075954 | 12645 | 2.0× |
| LightningDecodeAny | 1979550 | 233.12 MB/s | 2785926 | 66022 | 1.8× |
| Stdlib | 3570067 | 142.99 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 608 | 32545.98 MB/s | 0 | 0 | 145.6× |
| LightningDestructive | 685 | 28894.59 MB/s | 0 | 0 | 129.3× |
| SonicFastest | 3405 | 5811.41 MB/s | 21134 | 3 | 26.0× |
| Goccy | 11956 | 1655.14 MB/s | 20492 | 2 | 7.4× |
| JSONV2 | 16691 | 1185.63 MB/s | 8 | 1 | 5.3× |
| Sonic | 17704 | 1117.74 MB/s | 20692 | 3 | 5.0× |
| LightningDecodeAny | 53947 | 366.81 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 55185 | 358.59 MB/s | 0 | 0 | 1.6× |
| Stdlib | 88526 | 223.54 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1218 | 14885.93 MB/s | 432 | 2 | 57.7× |
| LightningDestructive | 1395 | 12990.35 MB/s | 0 | 0 | 50.4× |
| Easyjson | 2320 | 7813.56 MB/s | 432 | 2 | 30.3× |
| SonicFastest | 5406 | 3352.43 MB/s | 20475 | 5 | 13.0× |
| Sonic | 5594 | 3239.97 MB/s | 20450 | 5 | 12.6× |
| LightningDecodeAny | 10827 | 1651.63 MB/s | 29088 | 191 | 6.5× |
| Goccy | 11883 | 1525.26 MB/s | 19460 | 2 | 5.9× |
| JSONV2 | 26733 | 677.97 MB/s | 16500 | 50 | 2.6× |
| Stdlib | 70251 | 257.99 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1223525 | 1641.56 MB/s | 2960389 | 7411 | 9.8× |
| Lightning | 1323994 | 1517.00 MB/s | 2962101 | 7417 | 9.0× |
| Goccy | 2455897 | 817.83 MB/s | 5411060 | 15832 | 4.9× |
| Sonic | 2833321 | 708.88 MB/s | 5167001 | 7085 | 4.2× |
| Easyjson | 2834047 | 708.70 MB/s | 2981488 | 7439 | 4.2× |
| SonicFastest | 2930846 | 685.29 MB/s | 5165637 | 7085 | 4.1× |
| JSONV2 | 3734563 | 537.81 MB/s | 3173678 | 14563 | 3.2× |
| LightningDecodeAny | 3975808 | 287.31 MB/s | 7386074 | 134751 | 3.0× |
| Stdlib | 11965816 | 167.85 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 473 | 1161.40 MB/s | 480 | 1 | 7.3× |
| LightningDestructive | 478 | 1148.04 MB/s | 480 | 1 | 7.2× |
| LightningDecodeAny | 1108 | 494.56 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 1145 | 479.39 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 1216 | 451.61 MB/s | 2261 | 8 | 2.8× |
| Sonic | 1242 | 442.10 MB/s | 2260 | 8 | 2.8× |
| Goccy | 1592 | 344.92 MB/s | 2129 | 43 | 2.2× |
| JSONV2 | 1680 | 326.80 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 3439 | 159.62 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 266356 | 2370.94 MB/s | 364473 | 566 | 13.6× |
| Lightning | 304662 | 2072.84 MB/s | 413001 | 878 | 11.9× |
| Easyjson | 632793 | 997.98 MB/s | 422504 | 936 | 5.7× |
| Sonic | 655889 | 962.84 MB/s | 1066914 | 814 | 5.5× |
| SonicFastest | 674416 | 936.39 MB/s | 1067215 | 814 | 5.4× |
| Goccy | 733018 | 861.53 MB/s | 983711 | 1200 | 5.0× |
| JSONV2 | 1224319 | 515.81 MB/s | 571593 | 3144 | 3.0× |
| LightningDecodeAny | 1310485 | 356.28 MB/s | 2010198 | 30295 | 2.8× |
| Stdlib | 3634174 | 173.77 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 410204 | 1371.04 MB/s | 544248 | 448 | 8.2× |
| Lightning | 512142 | 1098.15 MB/s | 767621 | 1254 | 6.5× |
| Sonic | 818795 | 686.87 MB/s | 1347362 | 1184 | 4.1× |
| SonicFastest | 836342 | 672.46 MB/s | 1347137 | 1185 | 4.0× |
| Goccy | 841773 | 668.12 MB/s | 1035683 | 1028 | 4.0× |
| Easyjson | 1088772 | 516.55 MB/s | 775153 | 1254 | 3.1× |
| JSONV2 | 1673953 | 335.98 MB/s | 927407 | 3482 | 2.0× |
| LightningDecodeAny | 1691179 | 332.55 MB/s | 2114151 | 30295 | 2.0× |
| Stdlib | 3348148 | 167.98 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 336724 | 1583.43 MB/s | 333416 | 2084 | 10.7× |
| Lightning | 380108 | 1402.70 MB/s | 368224 | 2293 | 9.4× |
| Easyjson | 693574 | 768.74 MB/s | 428362 | 3273 | 5.2× |
| SonicFastest | 876662 | 608.19 MB/s | 982251 | 3082 | 4.1× |
| Goccy | 879375 | 606.32 MB/s | 1167023 | 5408 | 4.1× |
| Sonic | 882914 | 603.88 MB/s | 982505 | 3082 | 4.1× |
| JSONV2 | 1581532 | 337.13 MB/s | 745427 | 13288 | 2.3× |
| LightningDecodeAny | 1895943 | 281.22 MB/s | 2674617 | 50596 | 1.9× |
| Stdlib | 3589289 | 148.55 MB/s | 798693 | 17133 | 1.0× |
