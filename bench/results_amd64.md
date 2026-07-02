# JSON Deserialization Benchmarks

- generated 2026-07-02T07:16:07Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 102666 | 1239.70 MB/s | 49760 | 3 | 12.1× |
| LightningDestructive | 105023 | 1211.88 MB/s | 49280 | 2 | 11.8× |
| Sonic | 231710 | 549.28 MB/s | 214166 | 15 | 5.4× |
| SonicFastest | 232360 | 547.75 MB/s | 214491 | 15 | 5.3× |
| Easyjson | 232792 | 546.73 MB/s | 122864 | 14 | 5.3× |
| Goccy | 245461 | 518.51 MB/s | 225468 | 884 | 5.1× |
| JSONV2 | 402057 | 316.56 MB/s | 195127 | 1805 | 3.1× |
| LightningDecodeAny | 437375 | 216.41 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1239797 | 102.66 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4209547 | 534.75 MB/s | 3122872 | 3081 | 7.4× |
| LightningDestructive | 4220732 | 533.33 MB/s | 3122873 | 3081 | 7.4× |
| SonicFastest | 5597598 | 402.15 MB/s | 4861986 | 2584 | 5.6× |
| Sonic | 5623883 | 400.27 MB/s | 4863660 | 2584 | 5.5× |
| LightningDecodeAny | 11738209 | 191.77 MB/s | 7938299 | 281383 | 2.7× |
| Easyjson | 13742846 | 163.80 MB/s | 3099809 | 2120 | 2.3× |
| Goccy | 14306378 | 157.35 MB/s | 4152125 | 56533 | 2.2× |
| JSONV2 | 16747974 | 134.41 MB/s | 3123205 | 3083 | 1.9× |
| Stdlib | 31166135 | 72.23 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 557961 | 484.63 MB/s | 348024 | 1627 | 7.2× |
| LightningDestructive | 567667 | 476.34 MB/s | 348024 | 1627 | 7.1× |
| Sonic | 733573 | 368.61 MB/s | 642466 | 1147 | 5.5× |
| SonicFastest | 737268 | 366.76 MB/s | 642564 | 1147 | 5.5× |
| LightningDecodeAny | 1722148 | 157.01 MB/s | 1011486 | 37901 | 2.3× |
| Goccy | 1746103 | 154.86 MB/s | 541913 | 8122 | 2.3× |
| Easyjson | 1747344 | 154.75 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2198014 | 123.02 MB/s | 348161 | 1628 | 1.8× |
| Stdlib | 4033574 | 67.04 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1315398 | 1313.07 MB/s | 959848 | 5881 | 12.8× |
| Lightning | 1338048 | 1290.84 MB/s | 959890 | 5882 | 12.6× |
| Sonic | 2117168 | 815.81 MB/s | 2692758 | 5547 | 7.9× |
| SonicFastest | 2124894 | 812.84 MB/s | 2693516 | 5547 | 7.9× |
| Goccy | 2941775 | 587.13 MB/s | 2580947 | 14603 | 5.7× |
| Easyjson | 3987722 | 433.13 MB/s | 972032 | 5389 | 4.2× |
| JSONV2 | 4042101 | 427.30 MB/s | 1011615 | 7594 | 4.2× |
| LightningDecodeAny | 4092849 | 122.24 MB/s | 4976203 | 81466 | 4.1× |
| Stdlib | 16823030 | 102.67 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1022 | 1773.53 MB/s | 0 | 0 | 14.3× |
| LightningDestructive | 1046 | 1732.03 MB/s | 0 | 0 | 14.0× |
| Easyjson | 2853 | 635.14 MB/s | 24 | 1 | 5.1× |
| Goccy | 3137 | 577.67 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6434 | 281.62 MB/s | 3349 | 38 | 2.3× |
| Sonic | 6718 | 269.72 MB/s | 3349 | 38 | 2.2× |
| JSONV2 | 7535 | 240.46 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 8577 | 211.15 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14663 | 123.58 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1084 | 1671.65 MB/s | 0 | 0 | 13.5× |
| LightningDestructive | 1134 | 1598.06 MB/s | 0 | 0 | 12.9× |
| Easyjson | 2872 | 630.87 MB/s | 24 | 1 | 5.1× |
| Goccy | 3095 | 585.41 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6022 | 300.91 MB/s | 3347 | 38 | 2.4× |
| Sonic | 6350 | 285.34 MB/s | 3350 | 38 | 2.3× |
| JSONV2 | 7473 | 242.48 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8552 | 211.77 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14655 | 123.64 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1276 | 1420.40 MB/s | 144 | 10 | 11.4× |
| LightningDestructive | 1302 | 1391.44 MB/s | 144 | 10 | 11.1× |
| Easyjson | 2930 | 618.38 MB/s | 144 | 10 | 5.0× |
| Goccy | 3359 | 539.48 MB/s | 2600 | 5 | 4.3× |
| SonicFastest | 6208 | 291.87 MB/s | 3366 | 40 | 2.3× |
| Sonic | 6426 | 281.99 MB/s | 3364 | 40 | 2.3× |
| JSONV2 | 8053 | 225.01 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8894 | 203.63 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14516 | 124.83 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 700 | 705.88 MB/s | 160 | 1 | 8.7× |
| LightningDestructive | 712 | 694.11 MB/s | 160 | 1 | 8.5× |
| Sonic | 1216 | 406.16 MB/s | 1075 | 8 | 5.0× |
| SonicFastest | 1223 | 403.84 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1633 | 301.91 MB/s | 1536 | 30 | 3.7× |
| Easyjson | 2412 | 204.82 MB/s | 448 | 3 | 2.5× |
| Goccy | 2575 | 191.88 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 2989 | 165.28 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6064 | 81.47 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 426 | 540.39 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 431 | 534.04 MB/s | 160 | 1 | 9.8× |
| Sonic | 852 | 269.86 MB/s | 801 | 8 | 4.9× |
| SonicFastest | 861 | 267.10 MB/s | 801 | 8 | 4.9× |
| LightningDecodeAny | 1399 | 163.73 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1563 | 147.12 MB/s | 448 | 3 | 2.7× |
| Goccy | 1729 | 133.03 MB/s | 584 | 23 | 2.4× |
| JSONV2 | 2367 | 97.19 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4214 | 54.58 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 84115 | 774.32 MB/s | 172433 | 107 | 7.6× |
| LightningDestructive | 87376 | 745.42 MB/s | 166213 | 102 | 7.3× |
| Sonic | 151613 | 429.59 MB/s | 235703 | 65 | 4.2× |
| SonicFastest | 157761 | 412.85 MB/s | 236067 | 65 | 4.0× |
| LightningDecodeAny | 185371 | 287.69 MB/s | 176960 | 3252 | 3.4× |
| Goccy | 189381 | 343.92 MB/s | 228696 | 134 | 3.4× |
| JSONV2 | 252607 | 257.84 MB/s | 206664 | 607 | 2.5× |
| Stdlib | 637510 | 102.17 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2603212 | 745.41 MB/s | 2846867 | 2238 | 9.7× |
| LightningDestructive | 2653562 | 731.27 MB/s | 2846864 | 2238 | 9.5× |
| Goccy | 4989483 | 388.91 MB/s | 4064101 | 13509 | 5.1× |
| Sonic | 5099087 | 380.55 MB/s | 4878527 | 1736 | 5.0× |
| SonicFastest | 5188437 | 374.00 MB/s | 4880513 | 1736 | 4.9× |
| Easyjson | 7721666 | 251.30 MB/s | 3871265 | 15043 | 3.3× |
| LightningDecodeAny | 9516833 | 203.90 MB/s | 7013525 | 219937 | 2.7× |
| JSONV2 | 11781025 | 164.71 MB/s | 3237191 | 13947 | 2.1× |
| Stdlib | 25275298 | 76.77 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1015008 | 3278.63 MB/s | 351704 | 1286 | 25.3× |
| Lightning | 1640206 | 2028.91 MB/s | 2488906 | 2995 | 15.7× |
| Sonic | 2313002 | 1438.75 MB/s | 5894386 | 4263 | 11.1× |
| SonicFastest | 2373132 | 1402.30 MB/s | 5893207 | 4263 | 10.8× |
| LightningDecodeAny | 3533259 | 869.95 MB/s | 4886621 | 56892 | 7.3× |
| Goccy | 6595338 | 504.57 MB/s | 3948909 | 3816 | 3.9× |
| JSONV2 | 8637966 | 385.26 MB/s | 5364504 | 13243 | 3.0× |
| Stdlib | 25679705 | 129.59 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225719 | 976.20 MB/s | 136896 | 228 | 9.9× |
| LightningDestructive | 235938 | 933.92 MB/s | 136896 | 228 | 9.5× |
| Goccy | 461572 | 477.38 MB/s | 364145 | 1066 | 4.8× |
| Sonic | 511221 | 431.02 MB/s | 350761 | 262 | 4.4× |
| SonicFastest | 513620 | 429.01 MB/s | 351001 | 262 | 4.3× |
| Easyjson | 591516 | 372.51 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 682036 | 323.07 MB/s | 129747 | 470 | 3.3× |
| LightningDecodeAny | 950630 | 113.94 MB/s | 861394 | 11944 | 2.3× |
| Stdlib | 2230805 | 98.77 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14896063 | 543.77 MB/s | 15059832 | 51649 | 6.8× |
| Lightning | 14964861 | 541.27 MB/s | 15059837 | 51649 | 6.7× |
| Sonic | 21482523 | 377.05 MB/s | 19853711 | 41640 | 4.7× |
| SonicFastest | 21535777 | 376.12 MB/s | 19856446 | 41640 | 4.7× |
| Goccy | 25973177 | 311.86 MB/s | 19186747 | 107156 | 3.9× |
| Easyjson | 34378494 | 235.61 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 41458381 | 125.50 MB/s | 34333348 | 912810 | 2.4× |
| JSONV2 | 46580919 | 173.89 MB/s | 15233767 | 78972 | 2.2× |
| Stdlib | 101008885 | 80.19 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5469112 | 545.51 MB/s | 3985336 | 30015 | 9.2× |
| Lightning | 5645548 | 528.46 MB/s | 3985337 | 30015 | 8.9× |
| Sonic | 9168891 | 325.39 MB/s | 9128292 | 57804 | 5.5× |
| SonicFastest | 9172580 | 325.26 MB/s | 9126129 | 57804 | 5.5× |
| Goccy | 17750683 | 168.08 MB/s | 9845499 | 273618 | 2.8× |
| Easyjson | 18066697 | 165.14 MB/s | 9479441 | 30115 | 2.8× |
| LightningDecodeAny | 19193795 | 95.56 MB/s | 20023837 | 408557 | 2.6× |
| JSONV2 | 24439295 | 122.08 MB/s | 9257055 | 86278 | 2.1× |
| Stdlib | 50490187 | 59.09 MB/s | 9258083 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1256824 | 575.73 MB/s | 907393 | 3618 | 9.7× |
| Lightning | 1318848 | 548.66 MB/s | 907387 | 3618 | 9.3× |
| SonicFastest | 2175510 | 332.61 MB/s | 2367338 | 3683 | 5.6× |
| Sonic | 2181439 | 331.71 MB/s | 2367376 | 3683 | 5.6× |
| Easyjson | 5252548 | 137.76 MB/s | 2847906 | 3698 | 2.3× |
| LightningDecodeAny | 5319842 | 122.29 MB/s | 5752205 | 80172 | 2.3× |
| Goccy | 5387454 | 134.31 MB/s | 2631507 | 80263 | 2.3× |
| JSONV2 | 6295904 | 114.93 MB/s | 2704703 | 7318 | 1.9× |
| Stdlib | 12240049 | 59.12 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1990288 | 792.53 MB/s | 907392 | 3618 | 8.9× |
| Lightning | 2000032 | 788.66 MB/s | 907387 | 3618 | 8.9× |
| Sonic | 2434311 | 647.97 MB/s | 3221650 | 3683 | 7.3× |
| SonicFastest | 2508663 | 628.76 MB/s | 3227422 | 3683 | 7.1× |
| LightningDecodeAny | 4472618 | 168.45 MB/s | 5752203 | 80172 | 4.0× |
| Easyjson | 6462892 | 244.06 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 6557804 | 240.53 MB/s | 3428976 | 80258 | 2.7× |
| JSONV2 | 6699284 | 235.45 MB/s | 2704555 | 7318 | 2.7× |
| Stdlib | 17783897 | 88.70 MB/s | 2704547 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 213404 | 703.47 MB/s | 81920 | 1 | 10.4× |
| LightningDestructive | 223325 | 672.22 MB/s | 81920 | 1 | 9.9× |
| SonicFastest | 420211 | 357.26 MB/s | 408747 | 16 | 5.3× |
| Sonic | 422978 | 354.92 MB/s | 407976 | 16 | 5.3× |
| LightningDecodeAny | 585875 | 256.23 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 1000154 | 150.10 MB/s | 325752 | 10005 | 2.2× |
| JSONV2 | 1139651 | 131.73 MB/s | 357724 | 20 | 1.9× |
| Stdlib | 2221750 | 67.57 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 31506 | 892.43 MB/s | 30272 | 105 | 10.0× |
| LightningDestructive | 32513 | 864.78 MB/s | 30144 | 103 | 9.7× |
| SonicFastest | 65754 | 427.61 MB/s | 59452 | 83 | 4.8× |
| Sonic | 65894 | 426.70 MB/s | 59465 | 83 | 4.8× |
| Easyjson | 75292 | 373.44 MB/s | 32304 | 138 | 4.2× |
| Goccy | 76335 | 368.34 MB/s | 59267 | 188 | 4.1× |
| JSONV2 | 126267 | 222.68 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 153236 | 183.49 MB/s | 135024 | 2678 | 2.1× |
| Stdlib | 314478 | 89.41 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1987 | 1171.53 MB/s | 32 | 1 | 11.9× |
| LightningDestructive | 2010 | 1158.24 MB/s | 32 | 1 | 11.8× |
| Goccy | 4690 | 496.37 MB/s | 3649 | 4 | 5.0× |
| Easyjson | 4747 | 490.37 MB/s | 192 | 2 | 5.0× |
| SonicFastest | 6106 | 381.24 MB/s | 3709 | 4 | 3.9× |
| Sonic | 6122 | 380.25 MB/s | 3711 | 4 | 3.9× |
| JSONV2 | 7799 | 298.51 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 9550 | 176.44 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 23632 | 98.51 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220 | 858.07 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 225 | 841.55 MB/s | 0 | 0 | 10.9× |
| Goccy | 398 | 474.86 MB/s | 304 | 2 | 6.1× |
| Easyjson | 563 | 335.76 MB/s | 0 | 0 | 4.3× |
| Sonic | 715 | 264.39 MB/s | 341 | 3 | 3.4× |
| SonicFastest | 717 | 263.75 MB/s | 341 | 3 | 3.4× |
| JSONV2 | 926 | 204.21 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1194 | 112.21 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2439 | 77.48 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1469 | 1491.40 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1503 | 1458.16 MB/s | 0 | 0 | 11.5× |
| Goccy | 3431 | 638.66 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3496 | 626.71 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6588 | 332.56 MB/s | 3601 | 38 | 2.6× |
| Sonic | 6790 | 322.67 MB/s | 3602 | 38 | 2.6× |
| JSONV2 | 7644 | 286.63 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8375 | 216.25 MB/s | 7536 | 158 | 2.1× |
| Stdlib | 17315 | 126.54 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 666212 | 766.24 MB/s | 703778 | 1012 | 9.7× |
| Lightning | 714984 | 713.97 MB/s | 703779 | 1012 | 9.0× |
| Goccy | 1276758 | 399.82 MB/s | 1134348 | 5006 | 5.1× |
| SonicFastest | 1566155 | 325.94 MB/s | 1308219 | 2014 | 4.1× |
| Easyjson | 1576868 | 323.73 MB/s | 863782 | 3012 | 4.1× |
| Sonic | 1591888 | 320.67 MB/s | 1308077 | 2014 | 4.1× |
| JSONV2 | 3054329 | 167.13 MB/s | 1075952 | 12645 | 2.1× |
| LightningDecodeAny | 3452077 | 133.68 MB/s | 2785927 | 66022 | 1.9× |
| Stdlib | 6462207 | 78.99 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 612 | 32312.15 MB/s | 0 | 0 | 232.7× |
| LightningDestructive | 881 | 22461.48 MB/s | 0 | 0 | 161.8× |
| SonicFastest | 7075 | 2797.22 MB/s | 21116 | 3 | 20.1× |
| Goccy | 26486 | 747.14 MB/s | 20492 | 2 | 5.4× |
| Sonic | 29103 | 679.96 MB/s | 20606 | 3 | 4.9× |
| JSONV2 | 35913 | 551.03 MB/s | 8 | 1 | 4.0× |
| LightningDecodeAny | 93868 | 210.81 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 112535 | 175.85 MB/s | 0 | 0 | 1.3× |
| Stdlib | 142521 | 138.85 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2216 | 8177.15 MB/s | 432 | 2 | 53.9× |
| LightningDestructive | 2363 | 7669.66 MB/s | 0 | 0 | 50.6× |
| Easyjson | 4672 | 3878.89 MB/s | 432 | 2 | 25.6× |
| Sonic | 9162 | 1978.24 MB/s | 20444 | 5 | 13.0× |
| SonicFastest | 9951 | 1821.40 MB/s | 20415 | 5 | 12.0× |
| LightningDecodeAny | 18366 | 973.64 MB/s | 29088 | 191 | 6.5× |
| Goccy | 22394 | 809.34 MB/s | 19460 | 2 | 5.3× |
| JSONV2 | 48048 | 377.21 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 119498 | 151.67 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2219450 | 904.95 MB/s | 2951716 | 7277 | 9.2× |
| Lightning | 2304761 | 871.45 MB/s | 2953430 | 7283 | 8.9× |
| Goccy | 4487475 | 447.58 MB/s | 5410339 | 15832 | 4.6× |
| Easyjson | 4898714 | 410.00 MB/s | 2981490 | 7439 | 4.2× |
| SonicFastest | 4990083 | 402.50 MB/s | 5152207 | 7085 | 4.1× |
| Sonic | 5013823 | 400.59 MB/s | 5151606 | 7085 | 4.1× |
| LightningDecodeAny | 6690761 | 170.73 MB/s | 7386073 | 134751 | 3.1× |
| JSONV2 | 7207432 | 278.67 MB/s | 3173673 | 14563 | 2.8× |
| Stdlib | 20504688 | 97.95 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 808 | 679.33 MB/s | 480 | 1 | 7.2× |
| LightningDestructive | 817 | 671.84 MB/s | 480 | 1 | 7.1× |
| LightningDecodeAny | 1894 | 289.34 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 1960 | 280.17 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2133 | 257.34 MB/s | 2262 | 8 | 2.7× |
| Sonic | 2168 | 253.22 MB/s | 2262 | 8 | 2.7× |
| Goccy | 2985 | 183.93 MB/s | 2129 | 43 | 1.9× |
| JSONV2 | 3050 | 179.98 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 5784 | 94.92 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 478770 | 1319.03 MB/s | 364472 | 566 | 12.5× |
| Lightning | 545183 | 1158.35 MB/s | 413002 | 878 | 11.0× |
| Sonic | 1106870 | 570.54 MB/s | 1067694 | 814 | 5.4× |
| SonicFastest | 1137021 | 555.41 MB/s | 1067844 | 814 | 5.3× |
| Easyjson | 1234144 | 511.70 MB/s | 422504 | 936 | 4.8× |
| Goccy | 1478422 | 427.15 MB/s | 992486 | 1201 | 4.0× |
| JSONV2 | 2199297 | 287.14 MB/s | 571588 | 3144 | 2.7× |
| LightningDecodeAny | 2383181 | 195.92 MB/s | 2010198 | 30295 | 2.5× |
| Stdlib | 5981539 | 105.58 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 735557 | 764.60 MB/s | 544249 | 448 | 7.7× |
| Lightning | 930149 | 604.64 MB/s | 767623 | 1254 | 6.1× |
| Sonic | 1355957 | 414.77 MB/s | 1347973 | 1185 | 4.2× |
| SonicFastest | 1359269 | 413.76 MB/s | 1348226 | 1185 | 4.2× |
| Goccy | 1569682 | 358.29 MB/s | 1041413 | 1029 | 3.6× |
| Easyjson | 1979723 | 284.08 MB/s | 775156 | 1254 | 2.9× |
| LightningDecodeAny | 2920624 | 192.56 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 3000714 | 187.42 MB/s | 927408 | 3482 | 1.9× |
| Stdlib | 5650140 | 99.54 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 579378 | 920.26 MB/s | 333416 | 2084 | 10.2× |
| Lightning | 661114 | 806.48 MB/s | 368224 | 2293 | 8.9× |
| Easyjson | 1249452 | 426.73 MB/s | 428362 | 3273 | 4.7× |
| SonicFastest | 1359252 | 392.26 MB/s | 979545 | 3082 | 4.4× |
| Sonic | 1364973 | 390.61 MB/s | 979257 | 3082 | 4.3× |
| Goccy | 1398042 | 381.37 MB/s | 1167045 | 5408 | 4.2× |
| JSONV2 | 2643007 | 201.73 MB/s | 745421 | 13288 | 2.2× |
| LightningDecodeAny | 3352508 | 159.04 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 5913293 | 90.17 MB/s | 798692 | 17133 | 1.0× |
