# JSON Deserialization Benchmarks

- generated 2026-06-20T20:17:02Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 106081 | 1199.79 MB/s | 50208 | 4 | 10.4× |
| SonicFastest | 184306 | 690.56 MB/s | 193502 | 10 | 6.0× |
| Sonic | 185567 | 685.87 MB/s | 193182 | 10 | 6.0× |
| Goccy | 203652 | 624.96 MB/s | 225125 | 884 | 5.4× |
| Easyjson | 213133 | 597.16 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 420985 | 302.33 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 462183 | 204.80 MB/s | 466033 | 9715 | 2.4× |
| Stdlib | 1108103 | 114.86 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3823609 | 588.72 MB/s | 3122875 | 3081 | 6.8× |
| SonicFastest | 4548805 | 494.87 MB/s | 15233838 | 970 | 5.7× |
| Sonic | 4737412 | 475.16 MB/s | 15239136 | 970 | 5.5× |
| Goccy | 10417297 | 216.09 MB/s | 4119921 | 56532 | 2.5× |
| Easyjson | 11193953 | 201.10 MB/s | 3099808 | 2120 | 2.3× |
| LightningDecodeAny | 11444426 | 196.69 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16130459 | 139.55 MB/s | 3123215 | 3083 | 1.6× |
| Stdlib | 26120203 | 86.18 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 486954 | 555.30 MB/s | 348024 | 1627 | 6.9× |
| Sonic | 634427 | 426.22 MB/s | 466300 | 968 | 5.3× |
| SonicFastest | 650611 | 415.61 MB/s | 494479 | 968 | 5.2× |
| Easyjson | 1419348 | 190.51 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1422844 | 190.04 MB/s | 543747 | 8122 | 2.4× |
| LightningDecodeAny | 1599739 | 169.03 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2093919 | 129.14 MB/s | 348156 | 1628 | 1.6× |
| Stdlib | 3356113 | 80.57 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1412826 | 1222.52 MB/s | 959938 | 5883 | 9.4× |
| SonicFastest | 2072922 | 833.22 MB/s | 2727389 | 4020 | 6.4× |
| Sonic | 2099938 | 822.50 MB/s | 2764131 | 4020 | 6.3× |
| Goccy | 2385118 | 724.16 MB/s | 2582396 | 14604 | 5.6× |
| JSONV2 | 4227682 | 408.55 MB/s | 1011634 | 7594 | 3.1× |
| Easyjson | 4273990 | 404.12 MB/s | 972033 | 5389 | 3.1× |
| LightningDecodeAny | 4800625 | 104.22 MB/s | 4976253 | 81467 | 2.8× |
| Stdlib | 13300902 | 129.86 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1235 | 1467.76 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2523 | 718.26 MB/s | 24 | 1 | 5.6× |
| Goccy | 2898 | 625.29 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6063 | 298.85 MB/s | 3874 | 40 | 2.3× |
| Sonic | 6080 | 298.04 MB/s | 3914 | 40 | 2.3× |
| JSONV2 | 7742 | 234.05 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8768 | 206.55 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14052 | 128.95 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1240 | 1461.28 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2559 | 707.99 MB/s | 24 | 1 | 5.5× |
| Goccy | 2894 | 626.03 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6078 | 298.11 MB/s | 3874 | 40 | 2.3× |
| Sonic | 6126 | 295.77 MB/s | 3885 | 40 | 2.3× |
| JSONV2 | 7676 | 236.07 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8763 | 206.67 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14061 | 128.86 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1444 | 1254.63 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2749 | 659.26 MB/s | 144 | 10 | 5.1× |
| Goccy | 2937 | 617.04 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6213 | 291.63 MB/s | 3787 | 42 | 2.3× |
| SonicFastest | 6240 | 290.39 MB/s | 3809 | 42 | 2.2× |
| JSONV2 | 7851 | 230.80 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8334 | 217.31 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14005 | 129.38 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 752 | 657.34 MB/s | 160 | 1 | 7.4× |
| Sonic | 1243 | 397.38 MB/s | 977 | 6 | 4.5× |
| SonicFastest | 1245 | 396.86 MB/s | 977 | 6 | 4.5× |
| LightningDecodeAny | 1658 | 297.30 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2213 | 223.25 MB/s | 448 | 3 | 2.5× |
| Goccy | 2450 | 201.64 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3212 | 153.81 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5561 | 88.83 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 477 | 482.12 MB/s | 160 | 1 | 8.8× |
| Sonic | 910 | 252.74 MB/s | 691 | 6 | 4.6× |
| SonicFastest | 912 | 252.30 MB/s | 689 | 6 | 4.6× |
| Easyjson | 1411 | 162.96 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1467 | 156.14 MB/s | 1536 | 30 | 2.9× |
| Goccy | 1626 | 141.46 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2408 | 95.53 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4187 | 54.93 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 79635 | 817.88 MB/s | 178609 | 112 | 6.9× |
| Sonic | 101047 | 644.57 MB/s | 159639 | 75 | 5.4× |
| SonicFastest | 102343 | 636.41 MB/s | 160659 | 75 | 5.4× |
| Goccy | 144764 | 449.92 MB/s | 229090 | 134 | 3.8× |
| LightningDecodeAny | 196454 | 271.46 MB/s | 183136 | 3257 | 2.8× |
| JSONV2 | 226959 | 286.98 MB/s | 206653 | 607 | 2.4× |
| Stdlib | 548144 | 118.82 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2742109 | 707.66 MB/s | 2846868 | 2238 | 8.6× |
| SonicFastest | 4536066 | 427.79 MB/s | 14606972 | 1407 | 5.2× |
| Sonic | 4542811 | 427.15 MB/s | 14608594 | 1407 | 5.2× |
| Goccy | 4800490 | 404.22 MB/s | 4065785 | 13510 | 4.9× |
| Easyjson | 7623177 | 254.55 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 10164686 | 190.90 MB/s | 7013526 | 219937 | 2.3× |
| JSONV2 | 11240158 | 172.64 MB/s | 3237226 | 13947 | 2.1× |
| Stdlib | 23454400 | 82.73 MB/s | 3551333 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2302697 | 1445.19 MB/s | 4607418 | 4704 | 9.1× |
| Sonic | 2688842 | 1237.64 MB/s | 6370320 | 4248 | 7.8× |
| SonicFastest | 2702636 | 1231.33 MB/s | 6422346 | 4248 | 7.7× |
| LightningDecodeAny | 4458856 | 689.36 MB/s | 7005135 | 58601 | 4.7× |
| Goccy | 4667611 | 712.96 MB/s | 3948909 | 3816 | 4.5× |
| JSONV2 | 7447283 | 446.85 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20905212 | 159.19 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 229792 | 958.89 MB/s | 136896 | 228 | 8.9× |
| Sonic | 379474 | 580.66 MB/s | 299487 | 398 | 5.4× |
| SonicFastest | 379797 | 580.17 MB/s | 296925 | 398 | 5.4× |
| Goccy | 436016 | 505.36 MB/s | 364501 | 1067 | 4.7× |
| Easyjson | 549631 | 400.90 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 738892 | 298.21 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 891412 | 121.51 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2040991 | 107.96 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13270117 | 610.40 MB/s | 15059837 | 51649 | 6.7× |
| Sonic | 17064507 | 474.67 MB/s | 70902171 | 40014 | 5.2× |
| SonicFastest | 17103284 | 473.60 MB/s | 70887853 | 40014 | 5.2× |
| Goccy | 23795926 | 340.40 MB/s | 17053295 | 107148 | 3.8× |
| Easyjson | 30641925 | 264.34 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 40790229 | 127.56 MB/s | 34333351 | 912810 | 2.2× |
| JSONV2 | 44049029 | 183.89 MB/s | 15233744 | 78972 | 2.0× |
| Stdlib | 89414444 | 90.59 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6306835 | 473.05 MB/s | 3985339 | 30015 | 7.5× |
| SonicFastest | 8726662 | 341.88 MB/s | 26664805 | 56760 | 5.4× |
| Sonic | 8732468 | 341.65 MB/s | 26726746 | 56760 | 5.4× |
| Easyjson | 16610438 | 179.61 MB/s | 9479441 | 30115 | 2.8× |
| Goccy | 16837153 | 177.20 MB/s | 10584557 | 273649 | 2.8× |
| LightningDecodeAny | 19099146 | 96.04 MB/s | 20023839 | 408557 | 2.5× |
| JSONV2 | 24273437 | 122.91 MB/s | 9257148 | 86278 | 1.9× |
| Stdlib | 47099279 | 63.34 MB/s | 9258096 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1384268 | 522.73 MB/s | 907388 | 3618 | 8.4× |
| SonicFastest | 1790857 | 404.05 MB/s | 3180841 | 7226 | 6.5× |
| Sonic | 1815379 | 398.59 MB/s | 3175768 | 7226 | 6.4× |
| Easyjson | 4193341 | 172.56 MB/s | 2847906 | 3698 | 2.8× |
| LightningDecodeAny | 4394196 | 148.05 MB/s | 5752199 | 80172 | 2.6× |
| Goccy | 4757315 | 152.10 MB/s | 2790980 | 80272 | 2.4× |
| JSONV2 | 5886390 | 122.93 MB/s | 2704670 | 7318 | 2.0× |
| Stdlib | 11593305 | 62.42 MB/s | 2704557 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1977489 | 797.65 MB/s | 907386 | 3618 | 8.0× |
| SonicFastest | 2256805 | 698.93 MB/s | 5788124 | 7226 | 7.0× |
| Sonic | 2289763 | 688.87 MB/s | 5805229 | 7226 | 6.9× |
| LightningDecodeAny | 3891195 | 193.62 MB/s | 5752201 | 80172 | 4.0× |
| Easyjson | 5614584 | 280.94 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5726392 | 275.45 MB/s | 3578890 | 80267 | 2.7× |
| JSONV2 | 6520464 | 241.91 MB/s | 2704594 | 7318 | 2.4× |
| Stdlib | 15735404 | 100.24 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224589 | 668.44 MB/s | 81920 | 1 | 8.2× |
| SonicFastest | 271684 | 552.57 MB/s | 245046 | 6 | 6.8× |
| Sonic | 273610 | 548.68 MB/s | 252848 | 6 | 6.7× |
| LightningDecodeAny | 481800 | 311.58 MB/s | 746004 | 10020 | 3.8× |
| Goccy | 874256 | 171.72 MB/s | 330643 | 10005 | 2.1× |
| JSONV2 | 1100899 | 136.36 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1845580 | 81.34 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 37394 | 751.91 MB/s | 30384 | 107 | 8.1× |
| SonicFastest | 65251 | 430.91 MB/s | 48254 | 103 | 4.7× |
| Sonic | 65484 | 429.37 MB/s | 48476 | 103 | 4.6× |
| Easyjson | 68884 | 408.18 MB/s | 32304 | 138 | 4.4× |
| Goccy | 73852 | 380.72 MB/s | 59279 | 188 | 4.1× |
| JSONV2 | 135976 | 206.78 MB/s | 36897 | 242 | 2.2× |
| LightningDecodeAny | 161606 | 173.99 MB/s | 135136 | 2680 | 1.9× |
| Stdlib | 303611 | 92.61 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1954 | 1191.15 MB/s | 32 | 1 | 11.6× |
| Goccy | 4159 | 559.72 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4203 | 553.93 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5147 | 452.27 MB/s | 4503 | 6 | 4.4× |
| Sonic | 5153 | 451.78 MB/s | 4535 | 6 | 4.4× |
| JSONV2 | 8426 | 276.28 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 11035 | 152.69 MB/s | 9960 | 195 | 2.1× |
| Stdlib | 22736 | 102.39 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 230 | 821.27 MB/s | 0 | 0 | 10.8× |
| Goccy | 393 | 480.58 MB/s | 304 | 2 | 6.3× |
| Easyjson | 490 | 385.35 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 800 | 236.20 MB/s | 496 | 4 | 3.1× |
| Sonic | 804 | 234.93 MB/s | 498 | 4 | 3.1× |
| JSONV2 | 1041 | 181.48 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1268 | 105.65 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2476 | 76.34 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1533 | 1429.00 MB/s | 0 | 0 | 10.4× |
| Goccy | 3197 | 685.30 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3247 | 674.71 MB/s | 24 | 1 | 4.9× |
| SonicFastest | 6414 | 341.61 MB/s | 3966 | 40 | 2.5× |
| Sonic | 6467 | 338.79 MB/s | 4022 | 40 | 2.5× |
| JSONV2 | 7884 | 277.90 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8320 | 217.67 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 15946 | 137.40 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 755134 | 676.01 MB/s | 703778 | 1012 | 8.0× |
| Sonic | 1171986 | 435.57 MB/s | 900002 | 2006 | 5.2× |
| Goccy | 1175985 | 434.08 MB/s | 1137680 | 5006 | 5.2× |
| SonicFastest | 1178754 | 433.06 MB/s | 908607 | 2006 | 5.1× |
| Easyjson | 1541007 | 331.26 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3207922 | 159.13 MB/s | 1076003 | 12646 | 1.9× |
| LightningDecodeAny | 3562452 | 129.54 MB/s | 2785926 | 66022 | 1.7× |
| Stdlib | 6065260 | 84.16 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1337 | 14804.62 MB/s | 0 | 0 | 99.8× |
| Goccy | 20261 | 976.68 MB/s | 20491 | 2 | 6.6× |
| Sonic | 27585 | 717.38 MB/s | 22726 | 4 | 4.8× |
| SonicFastest | 27665 | 715.31 MB/s | 22788 | 4 | 4.8× |
| JSONV2 | 29694 | 666.43 MB/s | 8 | 1 | 4.5× |
| LightningDecodeAny | 77595 | 255.02 MB/s | 117104 | 2019 | 1.7× |
| Easyjson | 86049 | 229.97 MB/s | 0 | 0 | 1.6× |
| Stdlib | 133441 | 148.30 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2933 | 6179.63 MB/s | 864 | 4 | 35.0× |
| Easyjson | 3963 | 4573.22 MB/s | 432 | 2 | 25.9× |
| SonicFastest | 10320 | 1756.14 MB/s | 24106 | 6 | 10.0× |
| Sonic | 10505 | 1725.32 MB/s | 24123 | 6 | 9.8× |
| Goccy | 16544 | 1095.54 MB/s | 19459 | 2 | 6.2× |
| LightningDecodeAny | 18490 | 967.11 MB/s | 29520 | 193 | 5.6× |
| JSONV2 | 45788 | 395.82 MB/s | 16500 | 50 | 2.2× |
| Stdlib | 102792 | 176.32 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2385048 | 842.12 MB/s | 2963781 | 7423 | 7.8× |
| Goccy | 4427184 | 453.67 MB/s | 5412153 | 15830 | 4.2× |
| Sonic | 4537014 | 442.69 MB/s | 10908126 | 13683 | 4.1× |
| SonicFastest | 4620068 | 434.73 MB/s | 10883887 | 13683 | 4.0× |
| Easyjson | 4966883 | 404.38 MB/s | 2981483 | 7439 | 3.8× |
| JSONV2 | 6978695 | 287.80 MB/s | 3173685 | 14563 | 2.7× |
| LightningDecodeAny | 7574872 | 150.80 MB/s | 7387754 | 134757 | 2.5× |
| Stdlib | 18651501 | 107.69 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 902 | 608.37 MB/s | 480 | 1 | 6.2× |
| LightningDecodeAny | 1975 | 277.53 MB/s | 2261 | 50 | 2.8× |
| Easyjson | 2155 | 254.71 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2714 | 202.26 MB/s | 1941 | 26 | 2.1× |
| Sonic | 2724 | 201.56 MB/s | 1947 | 26 | 2.1× |
| Goccy | 3051 | 179.94 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3254 | 168.71 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5625 | 97.59 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 591783 | 1067.14 MB/s | 460874 | 1190 | 9.2× |
| Sonic | 1037702 | 608.57 MB/s | 997302 | 1102 | 5.2× |
| SonicFastest | 1043149 | 605.39 MB/s | 1010103 | 1102 | 5.2× |
| Goccy | 1161820 | 543.56 MB/s | 987327 | 1202 | 4.7× |
| Easyjson | 1162712 | 543.14 MB/s | 422505 | 936 | 4.7× |
| JSONV2 | 2155231 | 293.01 MB/s | 571620 | 3144 | 2.5× |
| LightningDecodeAny | 2520340 | 185.26 MB/s | 2058072 | 30607 | 2.2× |
| Stdlib | 5421395 | 116.49 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 985709 | 570.56 MB/s | 886433 | 2062 | 5.3× |
| Sonic | 1047536 | 536.89 MB/s | 969449 | 1476 | 5.0× |
| SonicFastest | 1057894 | 531.63 MB/s | 978892 | 1476 | 5.0× |
| Goccy | 1330649 | 422.66 MB/s | 1039981 | 1030 | 4.0× |
| Easyjson | 1750326 | 321.32 MB/s | 775154 | 1254 | 3.0× |
| JSONV2 | 2744878 | 204.89 MB/s | 927443 | 3482 | 1.9× |
| LightningDecodeAny | 2894333 | 194.31 MB/s | 2232993 | 31103 | 1.8× |
| Stdlib | 5259639 | 106.93 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 728389 | 732.00 MB/s | 402760 | 2502 | 7.6× |
| Easyjson | 1144307 | 465.94 MB/s | 428362 | 3273 | 4.8× |
| Sonic | 1150543 | 463.41 MB/s | 1033833 | 4351 | 4.8× |
| SonicFastest | 1153032 | 462.41 MB/s | 1036142 | 4351 | 4.8× |
| Goccy | 1309333 | 407.21 MB/s | 1167250 | 5409 | 4.2× |
| JSONV2 | 2557104 | 208.51 MB/s | 745471 | 13288 | 2.2× |
| LightningDecodeAny | 3535873 | 150.79 MB/s | 2709157 | 50805 | 1.6× |
| Stdlib | 5525502 | 96.49 MB/s | 798693 | 17133 | 1.0× |
