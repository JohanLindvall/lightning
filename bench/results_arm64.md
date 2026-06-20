# JSON Deserialization Benchmarks

- generated 2026-06-20T14:15:48Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| SonicFastest | 179602 | 708.65 MB/s | 185304 | 10 | 6.2× |
| Sonic | 180197 | 706.31 MB/s | 188259 | 10 | 6.1× |
| Goccy | 190154 | 669.33 MB/s | 224645 | 884 | 5.8× |
| Lightning | 192594 | 660.85 MB/s | 50208 | 4 | 5.7× |
| Easyjson | 211976 | 600.42 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 417951 | 304.52 MB/s | 195115 | 1805 | 2.6× |
| Stdlib | 1106371 | 115.04 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3980847 | 565.47 MB/s | 3122881 | 3081 | 6.6× |
| Sonic | 4610699 | 488.22 MB/s | 15247536 | 970 | 5.7× |
| SonicFastest | 4657350 | 483.33 MB/s | 15233766 | 970 | 5.6× |
| Goccy | 10595258 | 212.46 MB/s | 4256237 | 56540 | 2.5× |
| Easyjson | 11378147 | 197.84 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 16519586 | 136.27 MB/s | 3123376 | 3083 | 1.6× |
| Stdlib | 26253463 | 85.74 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 486123 | 556.24 MB/s | 348024 | 1627 | 6.9× |
| Sonic | 648436 | 417.01 MB/s | 506505 | 968 | 5.2× |
| SonicFastest | 649472 | 416.34 MB/s | 502833 | 968 | 5.2× |
| Goccy | 1423705 | 189.93 MB/s | 542658 | 8122 | 2.4× |
| Easyjson | 1438022 | 188.04 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2121122 | 127.48 MB/s | 348152 | 1628 | 1.6× |
| Stdlib | 3358563 | 80.51 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1491618 | 1157.94 MB/s | 964644 | 6177 | 9.0× |
| SonicFastest | 2141517 | 806.53 MB/s | 2953650 | 4020 | 6.3× |
| Sonic | 2152777 | 802.31 MB/s | 2917807 | 4020 | 6.2× |
| Goccy | 2499153 | 691.12 MB/s | 2587632 | 14607 | 5.4× |
| JSONV2 | 4290504 | 402.56 MB/s | 1011703 | 7594 | 3.1× |
| Easyjson | 4307965 | 400.93 MB/s | 972033 | 5389 | 3.1× |
| Stdlib | 13403643 | 128.86 MB/s | 1234456 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1232 | 1470.69 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2532 | 715.70 MB/s | 24 | 1 | 5.5× |
| Goccy | 2816 | 643.55 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6108 | 296.68 MB/s | 3747 | 40 | 2.3× |
| Sonic | 6119 | 296.13 MB/s | 3766 | 40 | 2.3× |
| JSONV2 | 7686 | 235.74 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14026 | 129.19 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1247 | 1452.54 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2529 | 716.41 MB/s | 24 | 1 | 5.6× |
| Goccy | 2829 | 640.56 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5998 | 302.11 MB/s | 3721 | 40 | 2.4× |
| SonicFastest | 6009 | 301.56 MB/s | 3743 | 40 | 2.3× |
| JSONV2 | 7803 | 232.23 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14112 | 128.40 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1431 | 1266.48 MB/s | 144 | 10 | 9.8× |
| Easyjson | 2735 | 662.41 MB/s | 144 | 10 | 5.1× |
| Goccy | 2884 | 628.37 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6248 | 290.01 MB/s | 3702 | 42 | 2.2× |
| Sonic | 6273 | 288.87 MB/s | 3746 | 42 | 2.2× |
| JSONV2 | 7967 | 227.42 MB/s | 632 | 7 | 1.8× |
| Stdlib | 13975 | 129.66 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 752 | 656.84 MB/s | 160 | 1 | 7.4× |
| Sonic | 1251 | 394.82 MB/s | 984 | 6 | 4.4× |
| SonicFastest | 1253 | 394.14 MB/s | 979 | 6 | 4.4× |
| Easyjson | 2246 | 219.97 MB/s | 448 | 3 | 2.5× |
| Goccy | 2434 | 202.96 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3240 | 152.47 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5530 | 89.33 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 463 | 497.04 MB/s | 160 | 1 | 8.9× |
| Sonic | 886 | 259.71 MB/s | 653 | 6 | 4.6× |
| SonicFastest | 891 | 258.26 MB/s | 655 | 6 | 4.6× |
| Easyjson | 1416 | 162.47 MB/s | 448 | 3 | 2.9× |
| Goccy | 1576 | 145.96 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2423 | 94.94 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4118 | 55.86 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 81948 | 794.80 MB/s | 178609 | 112 | 6.7× |
| Sonic | 98540 | 660.97 MB/s | 155628 | 75 | 5.6× |
| SonicFastest | 100783 | 646.26 MB/s | 157534 | 75 | 5.5× |
| Goccy | 148641 | 438.18 MB/s | 229049 | 134 | 3.7× |
| JSONV2 | 227579 | 286.20 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 552994 | 117.78 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2935362 | 661.07 MB/s | 2846870 | 2238 | 8.0× |
| Sonic | 4578461 | 423.83 MB/s | 14608651 | 1407 | 5.2× |
| SonicFastest | 4603274 | 421.54 MB/s | 14606973 | 1407 | 5.1× |
| Goccy | 4815071 | 403.00 MB/s | 4066428 | 13510 | 4.9× |
| Easyjson | 7778363 | 249.47 MB/s | 3871268 | 15043 | 3.0× |
| JSONV2 | 11528902 | 168.31 MB/s | 3237332 | 13947 | 2.0× |
| Stdlib | 23597080 | 82.23 MB/s | 3551321 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2561344 | 1299.25 MB/s | 4611520 | 5958 | 8.3× |
| SonicFastest | 2857442 | 1164.62 MB/s | 6758952 | 4249 | 7.4× |
| Sonic | 2875575 | 1157.28 MB/s | 6732634 | 4249 | 7.4× |
| Goccy | 4607528 | 722.26 MB/s | 3948952 | 3816 | 4.6× |
| JSONV2 | 7615988 | 436.95 MB/s | 5364579 | 13243 | 2.8× |
| Stdlib | 21174063 | 157.17 MB/s | 5565614 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 277829 | 793.10 MB/s | 96320 | 224 | 7.4× |
| Sonic | 381199 | 578.03 MB/s | 300665 | 398 | 5.4× |
| SonicFastest | 381598 | 577.43 MB/s | 301139 | 398 | 5.4× |
| Goccy | 435597 | 505.85 MB/s | 362915 | 1066 | 4.7× |
| Easyjson | 543460 | 405.45 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 717702 | 307.02 MB/s | 129739 | 470 | 2.9× |
| Stdlib | 2046498 | 107.67 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13981912 | 579.32 MB/s | 15059839 | 51649 | 6.5× |
| SonicFastest | 17437667 | 464.51 MB/s | 72287867 | 40015 | 5.2× |
| Sonic | 17456228 | 464.02 MB/s | 72420553 | 40015 | 5.2× |
| Goccy | 24598887 | 329.28 MB/s | 17518378 | 107152 | 3.7× |
| Easyjson | 31587128 | 256.43 MB/s | 15059619 | 41643 | 2.9× |
| JSONV2 | 45168386 | 179.33 MB/s | 15233864 | 78973 | 2.0× |
| Stdlib | 90465486 | 89.54 MB/s | 15665081 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6564383 | 454.49 MB/s | 3985342 | 30015 | 7.3× |
| SonicFastest | 8528177 | 349.84 MB/s | 25593019 | 56758 | 5.6× |
| Sonic | 8561485 | 348.48 MB/s | 25644967 | 56758 | 5.6× |
| Goccy | 16889994 | 176.64 MB/s | 10840994 | 273667 | 2.8× |
| Easyjson | 16975924 | 175.75 MB/s | 9479442 | 30115 | 2.8× |
| JSONV2 | 25022181 | 119.23 MB/s | 9257228 | 86279 | 1.9× |
| Stdlib | 48110351 | 62.01 MB/s | 9258096 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1392939 | 519.48 MB/s | 696699 | 3610 | 8.4× |
| Sonic | 1846218 | 391.93 MB/s | 3352407 | 7226 | 6.3× |
| SonicFastest | 1902301 | 380.38 MB/s | 3486420 | 7227 | 6.1× |
| Easyjson | 4355105 | 166.15 MB/s | 2847910 | 3698 | 2.7× |
| Goccy | 4880516 | 148.26 MB/s | 3228956 | 80295 | 2.4× |
| JSONV2 | 5966358 | 121.28 MB/s | 2704783 | 7319 | 2.0× |
| Stdlib | 11651792 | 62.10 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2059627 | 765.84 MB/s | 696697 | 3610 | 7.7× |
| Sonic | 2321403 | 679.48 MB/s | 5782454 | 7226 | 6.8× |
| SonicFastest | 2467246 | 639.32 MB/s | 6066804 | 7226 | 6.4× |
| Easyjson | 5704956 | 276.49 MB/s | 2847910 | 3698 | 2.8× |
| Goccy | 5857712 | 269.28 MB/s | 3717837 | 80275 | 2.7× |
| JSONV2 | 6725508 | 234.53 MB/s | 2704766 | 7319 | 2.4× |
| Stdlib | 15880587 | 99.33 MB/s | 2704555 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224719 | 668.05 MB/s | 81920 | 1 | 8.2× |
| SonicFastest | 280258 | 535.66 MB/s | 264516 | 6 | 6.6× |
| Sonic | 283171 | 530.15 MB/s | 271484 | 6 | 6.5× |
| Goccy | 876776 | 171.22 MB/s | 333121 | 10005 | 2.1× |
| JSONV2 | 1117677 | 134.32 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1848317 | 81.22 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 38336 | 733.44 MB/s | 29104 | 107 | 7.9× |
| Sonic | 63710 | 441.33 MB/s | 46765 | 103 | 4.8× |
| SonicFastest | 63749 | 441.06 MB/s | 46729 | 103 | 4.8× |
| Easyjson | 68078 | 413.01 MB/s | 32304 | 138 | 4.5× |
| Goccy | 72469 | 387.99 MB/s | 59197 | 188 | 4.2× |
| JSONV2 | 134831 | 208.54 MB/s | 36895 | 242 | 2.3× |
| Stdlib | 303495 | 92.64 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2039 | 1141.73 MB/s | 32 | 1 | 11.1× |
| Goccy | 4184 | 556.46 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4209 | 553.14 MB/s | 192 | 2 | 5.4× |
| Sonic | 5104 | 456.15 MB/s | 4184 | 6 | 4.4× |
| SonicFastest | 5138 | 453.08 MB/s | 4226 | 6 | 4.4× |
| JSONV2 | 8540 | 272.62 MB/s | 1000 | 6 | 2.7× |
| Stdlib | 22660 | 102.74 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 229 | 826.65 MB/s | 0 | 0 | 10.6× |
| Goccy | 390 | 484.46 MB/s | 304 | 2 | 6.2× |
| Easyjson | 495 | 381.53 MB/s | 0 | 0 | 4.9× |
| Sonic | 774 | 244.32 MB/s | 506 | 4 | 3.1× |
| SonicFastest | 778 | 242.81 MB/s | 502 | 4 | 3.1× |
| JSONV2 | 1037 | 182.22 MB/s | 112 | 1 | 2.3× |
| Stdlib | 2420 | 78.11 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1545 | 1417.97 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3202 | 684.32 MB/s | 24 | 1 | 5.0× |
| Goccy | 3218 | 680.96 MB/s | 2864 | 4 | 5.0× |
| SonicFastest | 6432 | 340.65 MB/s | 3966 | 40 | 2.5× |
| Sonic | 6450 | 339.70 MB/s | 3997 | 40 | 2.5× |
| JSONV2 | 8086 | 270.98 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15970 | 137.20 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 890842 | 573.03 MB/s | 703777 | 1012 | 6.9× |
| Sonic | 1170747 | 436.03 MB/s | 890151 | 2006 | 5.2× |
| SonicFastest | 1185067 | 430.76 MB/s | 930237 | 2006 | 5.2× |
| Goccy | 1198186 | 426.04 MB/s | 1145844 | 5007 | 5.1× |
| Easyjson | 1560915 | 327.04 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3271994 | 156.01 MB/s | 1076000 | 12646 | 1.9× |
| Stdlib | 6112588 | 83.51 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2063 | 9594.03 MB/s | 0 | 0 | 54.2× |
| Goccy | 19979 | 990.47 MB/s | 20491 | 2 | 5.6× |
| SonicFastest | 27218 | 727.05 MB/s | 22170 | 4 | 4.1× |
| Sonic | 27404 | 722.12 MB/s | 22391 | 4 | 4.1× |
| JSONV2 | 29781 | 664.49 MB/s | 8 | 1 | 3.8× |
| Easyjson | 86033 | 230.02 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111716 | 177.14 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2934 | 6176.78 MB/s | 864 | 4 | 35.0× |
| Easyjson | 3944 | 4594.77 MB/s | 432 | 2 | 26.1× |
| SonicFastest | 10057 | 1802.22 MB/s | 22791 | 6 | 10.2× |
| Sonic | 10148 | 1785.93 MB/s | 22870 | 6 | 10.1× |
| Goccy | 15891 | 1140.53 MB/s | 19459 | 2 | 6.5× |
| JSONV2 | 44793 | 404.61 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102770 | 176.35 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2624814 | 765.19 MB/s | 2963985 | 7453 | 7.2× |
| SonicFastest | 4268525 | 470.54 MB/s | 10035604 | 13682 | 4.4× |
| Sonic | 4278271 | 469.46 MB/s | 10035415 | 13682 | 4.4× |
| Goccy | 4434281 | 452.95 MB/s | 5416895 | 15844 | 4.3× |
| Easyjson | 5194153 | 386.68 MB/s | 2981720 | 7441 | 3.7× |
| JSONV2 | 7125888 | 281.86 MB/s | 3173801 | 14563 | 2.7× |
| Stdlib | 18961270 | 105.93 MB/s | 3589320 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 900 | 609.75 MB/s | 480 | 1 | 6.3× |
| Easyjson | 2183 | 251.53 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2735 | 200.73 MB/s | 1942 | 26 | 2.1× |
| SonicFastest | 2740 | 200.36 MB/s | 1942 | 26 | 2.1× |
| Goccy | 3051 | 179.92 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3302 | 166.24 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5702 | 96.29 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 573141 | 1101.85 MB/s | 461113 | 1230 | 9.4× |
| Sonic | 1041152 | 606.55 MB/s | 1036350 | 1102 | 5.2× |
| SonicFastest | 1044129 | 604.82 MB/s | 1040128 | 1102 | 5.1× |
| Easyjson | 1138675 | 554.60 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1185482 | 532.71 MB/s | 986815 | 1202 | 4.5× |
| JSONV2 | 2183723 | 289.19 MB/s | 571652 | 3144 | 2.5× |
| Stdlib | 5377117 | 117.44 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1054759 | 533.21 MB/s | 945693 | 4207 | 5.0× |
| Sonic | 1143559 | 491.81 MB/s | 1048930 | 2506 | 4.6× |
| SonicFastest | 1156161 | 486.44 MB/s | 1054598 | 2506 | 4.5× |
| Goccy | 1534397 | 366.53 MB/s | 1109505 | 3406 | 3.4× |
| Easyjson | 1804426 | 311.68 MB/s | 833617 | 3391 | 2.9× |
| JSONV2 | 2958640 | 190.09 MB/s | 979922 | 5184 | 1.8× |
| Stdlib | 5246967 | 107.19 MB/s | 1067141 | 8759 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1343913 | 396.74 MB/s | 935533 | 16692 | 4.1× |
| SonicFastest | 1500547 | 355.32 MB/s | 1507307 | 10471 | 3.7× |
| Sonic | 1504458 | 354.40 MB/s | 1506067 | 10471 | 3.7× |
| Easyjson | 1793067 | 297.36 MB/s | 949523 | 15985 | 3.1× |
| Goccy | 2094758 | 254.53 MB/s | 1669066 | 18051 | 2.7× |
| JSONV2 | 3156198 | 168.93 MB/s | 1192603 | 22667 | 1.8× |
| Stdlib | 5561553 | 95.87 MB/s | 1245604 | 26510 | 1.0× |
