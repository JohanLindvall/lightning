# JSON Deserialization Benchmarks

- generated 2026-06-21T09:44:39Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105969 | 1201.06 MB/s | 49760 | 3 | 12.5× |
| Sonic | 200182 | 635.80 MB/s | 213831 | 15 | 6.6× |
| SonicFastest | 201055 | 633.03 MB/s | 213944 | 15 | 6.6× |
| Easyjson | 241739 | 526.50 MB/s | 122864 | 14 | 5.5× |
| Goccy | 246671 | 515.97 MB/s | 224988 | 884 | 5.4× |
| JSONV2 | 472008 | 269.65 MB/s | 195129 | 1805 | 2.8× |
| LightningDecodeAny | 477688 | 198.15 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1320771 | 96.36 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4288176 | 524.94 MB/s | 3122872 | 3081 | 7.9× |
| SonicFastest | 5031267 | 447.41 MB/s | 4864106 | 2584 | 6.8× |
| Sonic | 5308683 | 424.03 MB/s | 4865758 | 2584 | 6.4× |
| LightningDecodeAny | 12369179 | 181.99 MB/s | 7938298 | 281383 | 2.8× |
| Goccy | 13185291 | 170.72 MB/s | 4171769 | 56534 | 2.6× |
| Easyjson | 13439210 | 167.50 MB/s | 3099808 | 2120 | 2.5× |
| JSONV2 | 17152998 | 131.23 MB/s | 3123189 | 3083 | 2.0× |
| Stdlib | 34023656 | 66.16 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 600288 | 450.46 MB/s | 348025 | 1627 | 7.3× |
| SonicFastest | 783994 | 344.90 MB/s | 644622 | 1147 | 5.6× |
| Sonic | 785698 | 344.16 MB/s | 644069 | 1147 | 5.6× |
| LightningDecodeAny | 1735318 | 155.82 MB/s | 1011488 | 37901 | 2.5× |
| Easyjson | 1776734 | 152.19 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1880111 | 143.82 MB/s | 548263 | 8124 | 2.3× |
| JSONV2 | 2336685 | 115.72 MB/s | 348158 | 1628 | 1.9× |
| Stdlib | 4395509 | 61.52 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1345688 | 1283.51 MB/s | 959890 | 5882 | 13.1× |
| SonicFastest | 2025383 | 852.78 MB/s | 2693806 | 5547 | 8.7× |
| Sonic | 2026347 | 852.37 MB/s | 2693639 | 5547 | 8.7× |
| Goccy | 2865961 | 602.66 MB/s | 2581074 | 14603 | 6.1× |
| Easyjson | 4194243 | 411.80 MB/s | 972032 | 5389 | 4.2× |
| LightningDecodeAny | 4330243 | 115.54 MB/s | 4976205 | 81466 | 4.1× |
| JSONV2 | 4521003 | 382.04 MB/s | 1011615 | 7594 | 3.9× |
| Stdlib | 17563152 | 98.34 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1050 | 1725.08 MB/s | 0 | 0 | 15.4× |
| Easyjson | 2861 | 633.40 MB/s | 24 | 1 | 5.6× |
| Goccy | 3493 | 518.80 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6316 | 286.87 MB/s | 3344 | 38 | 2.6× |
| Sonic | 6544 | 276.90 MB/s | 3348 | 38 | 2.5× |
| JSONV2 | 8527 | 212.49 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9532 | 190.00 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16149 | 112.20 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1149 | 1577.18 MB/s | 0 | 0 | 14.1× |
| Easyjson | 2843 | 637.29 MB/s | 24 | 1 | 5.7× |
| Goccy | 3472 | 521.88 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6203 | 292.13 MB/s | 3344 | 38 | 2.6× |
| Sonic | 6433 | 281.67 MB/s | 3343 | 38 | 2.5× |
| JSONV2 | 8388 | 216.01 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9484 | 190.96 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16184 | 111.96 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1374 | 1319.18 MB/s | 144 | 10 | 11.7× |
| Easyjson | 3146 | 575.94 MB/s | 144 | 10 | 5.1× |
| Goccy | 3510 | 516.24 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 6199 | 292.29 MB/s | 3368 | 40 | 2.6× |
| Sonic | 6454 | 280.77 MB/s | 3371 | 40 | 2.5× |
| JSONV2 | 8794 | 206.05 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 9423 | 192.18 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16097 | 112.57 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 745 | 662.76 MB/s | 160 | 1 | 9.4× |
| SonicFastest | 1263 | 391.03 MB/s | 1076 | 8 | 5.5× |
| Sonic | 1267 | 389.90 MB/s | 1076 | 8 | 5.5× |
| LightningDecodeAny | 1793 | 274.97 MB/s | 1536 | 30 | 3.9× |
| Easyjson | 2491 | 198.33 MB/s | 448 | 3 | 2.8× |
| Goccy | 2561 | 192.88 MB/s | 856 | 23 | 2.7× |
| JSONV2 | 3568 | 138.45 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6984 | 70.74 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 502 | 458.43 MB/s | 160 | 1 | 9.9× |
| Sonic | 910 | 252.69 MB/s | 801 | 8 | 5.5× |
| SonicFastest | 916 | 251.11 MB/s | 801 | 8 | 5.4× |
| LightningDecodeAny | 1561 | 146.69 MB/s | 1536 | 30 | 3.2× |
| Easyjson | 1696 | 135.62 MB/s | 448 | 3 | 2.9× |
| Goccy | 1784 | 128.91 MB/s | 584 | 23 | 2.8× |
| JSONV2 | 2766 | 83.15 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4986 | 46.13 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 95090 | 684.95 MB/s | 172434 | 107 | 7.1× |
| Sonic | 146514 | 444.54 MB/s | 235668 | 65 | 4.6× |
| SonicFastest | 147871 | 440.47 MB/s | 235757 | 65 | 4.6× |
| Goccy | 195294 | 333.51 MB/s | 228705 | 134 | 3.5× |
| LightningDecodeAny | 205435 | 259.59 MB/s | 176961 | 3252 | 3.3× |
| JSONV2 | 273837 | 237.85 MB/s | 206665 | 607 | 2.5× |
| Stdlib | 677814 | 96.09 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2626083 | 738.92 MB/s | 2846865 | 2238 | 10.7× |
| SonicFastest | 4869327 | 398.51 MB/s | 4874902 | 1736 | 5.8× |
| Sonic | 4885750 | 397.17 MB/s | 4875413 | 1736 | 5.8× |
| Goccy | 5079085 | 382.05 MB/s | 4062362 | 13509 | 5.5× |
| Easyjson | 8239488 | 235.51 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 9943984 | 195.14 MB/s | 7013525 | 219937 | 2.8× |
| JSONV2 | 12908990 | 150.32 MB/s | 3237190 | 13947 | 2.2× |
| Stdlib | 28116608 | 69.02 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1650609 | 2016.12 MB/s | 2488907 | 2995 | 14.3× |
| Sonic | 2005465 | 1659.38 MB/s | 5896425 | 4263 | 11.7× |
| SonicFastest | 2008494 | 1656.88 MB/s | 5896703 | 4263 | 11.7× |
| LightningDecodeAny | 3614267 | 850.45 MB/s | 4886622 | 56892 | 6.5× |
| Goccy | 5197327 | 640.30 MB/s | 3948915 | 3817 | 4.5× |
| JSONV2 | 7531995 | 441.83 MB/s | 5364504 | 13243 | 3.1× |
| Stdlib | 23557313 | 141.27 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 232808 | 946.47 MB/s | 136896 | 228 | 10.5× |
| Sonic | 479519 | 459.51 MB/s | 349834 | 262 | 5.1× |
| SonicFastest | 484200 | 455.07 MB/s | 350320 | 262 | 5.0× |
| Goccy | 484271 | 455.01 MB/s | 364007 | 1066 | 5.0× |
| Easyjson | 631391 | 348.99 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 796755 | 276.55 MB/s | 129746 | 470 | 3.1× |
| LightningDecodeAny | 1028022 | 105.36 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2443774 | 90.17 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 15415445 | 525.45 MB/s | 15059843 | 51649 | 7.2× |
| SonicFastest | 18022637 | 449.44 MB/s | 19859133 | 41640 | 6.2× |
| Sonic | 18142133 | 446.48 MB/s | 19857157 | 41640 | 6.2× |
| Goccy | 27301176 | 296.69 MB/s | 19107827 | 107155 | 4.1× |
| Easyjson | 34703346 | 233.41 MB/s | 15059617 | 41643 | 3.2× |
| LightningDecodeAny | 43769470 | 118.87 MB/s | 34333347 | 912810 | 2.5× |
| JSONV2 | 48628856 | 166.57 MB/s | 15233720 | 78972 | 2.3× |
| Stdlib | 111596431 | 72.58 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6476002 | 460.70 MB/s | 3985336 | 30015 | 8.9× |
| Sonic | 9363255 | 318.64 MB/s | 9130350 | 57804 | 6.1× |
| SonicFastest | 9373915 | 318.27 MB/s | 9130259 | 57804 | 6.1× |
| Goccy | 18551479 | 160.82 MB/s | 9912950 | 273622 | 3.1× |
| Easyjson | 19375523 | 153.98 MB/s | 9479450 | 30115 | 3.0× |
| LightningDecodeAny | 20613386 | 88.98 MB/s | 20023837 | 408557 | 2.8× |
| JSONV2 | 27486898 | 108.54 MB/s | 9257047 | 86278 | 2.1× |
| Stdlib | 57464979 | 51.92 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1562223 | 463.18 MB/s | 907387 | 3618 | 9.1× |
| SonicFastest | 2125812 | 340.39 MB/s | 2371544 | 3683 | 6.7× |
| Sonic | 2155766 | 335.66 MB/s | 2372027 | 3683 | 6.6× |
| Easyjson | 5439726 | 133.02 MB/s | 2847906 | 3698 | 2.6× |
| Goccy | 5496342 | 131.65 MB/s | 2708020 | 80267 | 2.6× |
| LightningDecodeAny | 5949998 | 109.34 MB/s | 5752201 | 80172 | 2.4× |
| JSONV2 | 6210166 | 116.52 MB/s | 2704707 | 7318 | 2.3× |
| Stdlib | 14156186 | 51.12 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2048766 | 769.90 MB/s | 907387 | 3618 | 9.7× |
| SonicFastest | 2404571 | 655.98 MB/s | 3222903 | 3683 | 8.3× |
| Sonic | 2412221 | 653.90 MB/s | 3222641 | 3683 | 8.3× |
| LightningDecodeAny | 5012853 | 150.29 MB/s | 5752200 | 80172 | 4.0× |
| Easyjson | 6463956 | 244.02 MB/s | 2847905 | 3698 | 3.1× |
| Goccy | 6601098 | 238.95 MB/s | 3506605 | 80263 | 3.0× |
| JSONV2 | 6820829 | 231.26 MB/s | 2704554 | 7318 | 2.9× |
| Stdlib | 19965615 | 79.00 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 247770 | 605.90 MB/s | 81920 | 1 | 8.7× |
| SonicFastest | 387239 | 387.68 MB/s | 407158 | 16 | 5.5× |
| Sonic | 391079 | 383.87 MB/s | 407178 | 16 | 5.5× |
| LightningDecodeAny | 607086 | 247.28 MB/s | 746006 | 10020 | 3.5× |
| Goccy | 991360 | 151.43 MB/s | 324652 | 10005 | 2.2× |
| JSONV2 | 1127270 | 133.17 MB/s | 357724 | 20 | 1.9× |
| Stdlib | 2148328 | 69.88 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 37486 | 750.07 MB/s | 30272 | 105 | 9.2× |
| Sonic | 59830 | 469.95 MB/s | 59496 | 83 | 5.8× |
| SonicFastest | 59989 | 468.70 MB/s | 59491 | 83 | 5.8× |
| Easyjson | 80815 | 347.92 MB/s | 32304 | 138 | 4.3× |
| Goccy | 83354 | 337.32 MB/s | 59259 | 188 | 4.2× |
| JSONV2 | 148269 | 189.64 MB/s | 36897 | 242 | 2.3× |
| LightningDecodeAny | 170972 | 164.45 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 346620 | 81.12 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2039 | 1141.95 MB/s | 32 | 1 | 12.9× |
| SonicFastest | 4884 | 476.65 MB/s | 3712 | 4 | 5.4× |
| Sonic | 4924 | 472.78 MB/s | 3715 | 4 | 5.3× |
| Goccy | 5085 | 457.84 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 5508 | 422.64 MB/s | 192 | 2 | 4.8× |
| JSONV2 | 8362 | 278.41 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10979 | 153.47 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 26317 | 88.46 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222 | 852.04 MB/s | 0 | 0 | 12.6× |
| Goccy | 447 | 422.62 MB/s | 304 | 2 | 6.2× |
| Easyjson | 550 | 343.66 MB/s | 0 | 0 | 5.1× |
| SonicFastest | 659 | 286.81 MB/s | 341 | 3 | 4.2× |
| Sonic | 663 | 284.92 MB/s | 341 | 3 | 4.2× |
| JSONV2 | 1062 | 177.94 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1398 | 95.86 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2791 | 67.71 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1450 | 1511.07 MB/s | 0 | 0 | 13.0× |
| Easyjson | 3508 | 624.56 MB/s | 24 | 1 | 5.4× |
| Goccy | 3964 | 552.74 MB/s | 2864 | 4 | 4.8× |
| SonicFastest | 6720 | 326.07 MB/s | 3602 | 38 | 2.8× |
| Sonic | 7051 | 310.73 MB/s | 3605 | 38 | 2.7× |
| JSONV2 | 8619 | 254.22 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9591 | 188.83 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18895 | 115.96 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 762204 | 669.74 MB/s | 703778 | 1012 | 9.5× |
| Goccy | 1268150 | 402.54 MB/s | 1137184 | 5006 | 5.7× |
| SonicFastest | 1280980 | 398.50 MB/s | 1308234 | 2014 | 5.6× |
| Sonic | 1290813 | 395.47 MB/s | 1308497 | 2014 | 5.6× |
| Easyjson | 1788463 | 285.43 MB/s | 863779 | 3012 | 4.0× |
| JSONV2 | 3552043 | 143.71 MB/s | 1075978 | 12646 | 2.0× |
| LightningDecodeAny | 3766662 | 122.51 MB/s | 2785928 | 66022 | 1.9× |
| Stdlib | 7227920 | 70.63 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 641 | 30889.66 MB/s | 0 | 0 | 254.2× |
| SonicFastest | 6809 | 2906.15 MB/s | 21104 | 3 | 23.9× |
| Goccy | 29966 | 660.37 MB/s | 20492 | 2 | 5.4× |
| Sonic | 32567 | 607.64 MB/s | 20631 | 3 | 5.0× |
| JSONV2 | 33287 | 594.49 MB/s | 8 | 1 | 4.9× |
| Easyjson | 101240 | 195.47 MB/s | 0 | 0 | 1.6× |
| LightningDecodeAny | 104069 | 190.14 MB/s | 117104 | 2019 | 1.6× |
| Stdlib | 162839 | 121.53 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2340 | 7745.13 MB/s | 432 | 2 | 52.1× |
| Easyjson | 4793 | 3781.10 MB/s | 432 | 2 | 25.5× |
| Sonic | 8570 | 2114.93 MB/s | 20456 | 5 | 14.2× |
| SonicFastest | 8590 | 2109.78 MB/s | 20457 | 5 | 14.2× |
| LightningDecodeAny | 19924 | 897.49 MB/s | 29088 | 191 | 6.1× |
| Goccy | 25269 | 717.24 MB/s | 19460 | 2 | 4.8× |
| JSONV2 | 49859 | 363.51 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 121988 | 148.57 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2663439 | 754.10 MB/s | 2962101 | 7417 | 8.4× |
| SonicFastest | 4272059 | 470.15 MB/s | 5159375 | 7085 | 5.2× |
| Sonic | 4325661 | 464.32 MB/s | 5159774 | 7085 | 5.2× |
| Goccy | 4939192 | 406.64 MB/s | 5411816 | 15833 | 4.5× |
| Easyjson | 5568914 | 360.66 MB/s | 2981483 | 7439 | 4.0× |
| LightningDecodeAny | 7336443 | 155.70 MB/s | 7386073 | 134751 | 3.0× |
| JSONV2 | 8019487 | 250.45 MB/s | 3173676 | 14563 | 2.8× |
| Stdlib | 22312611 | 90.02 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 872 | 629.38 MB/s | 480 | 1 | 7.9× |
| LightningDecodeAny | 2133 | 256.92 MB/s | 2261 | 50 | 3.2× |
| Easyjson | 2285 | 240.30 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2428 | 226.09 MB/s | 2262 | 8 | 2.9× |
| Sonic | 2443 | 224.73 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3431 | 159.99 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3438 | 159.69 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 6924 | 79.28 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 596713 | 1058.32 MB/s | 413001 | 878 | 10.8× |
| SonicFastest | 997000 | 633.41 MB/s | 1065349 | 814 | 6.5× |
| Sonic | 998405 | 632.52 MB/s | 1065465 | 814 | 6.5× |
| Goccy | 1335029 | 473.03 MB/s | 990060 | 1201 | 4.8× |
| Easyjson | 1417406 | 445.54 MB/s | 422504 | 936 | 4.6× |
| JSONV2 | 2431686 | 259.70 MB/s | 571589 | 3144 | 2.7× |
| LightningDecodeAny | 2670180 | 174.86 MB/s | 2010197 | 30295 | 2.4× |
| Stdlib | 6473897 | 97.55 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1049498 | 535.88 MB/s | 767621 | 1254 | 6.0× |
| SonicFastest | 1319970 | 426.08 MB/s | 1349235 | 1185 | 4.7× |
| Sonic | 1324581 | 424.59 MB/s | 1349684 | 1185 | 4.7× |
| Goccy | 1539489 | 365.32 MB/s | 1040373 | 1028 | 4.1× |
| Easyjson | 2203612 | 255.22 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 3184251 | 176.62 MB/s | 2114150 | 30295 | 2.0× |
| JSONV2 | 3286305 | 171.14 MB/s | 927411 | 3482 | 1.9× |
| Stdlib | 6249770 | 89.99 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 718050 | 742.54 MB/s | 368224 | 2293 | 9.0× |
| Sonic | 1147887 | 464.49 MB/s | 980203 | 3082 | 5.6× |
| SonicFastest | 1154456 | 461.84 MB/s | 980640 | 3082 | 5.6× |
| Easyjson | 1305243 | 408.49 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1547088 | 344.63 MB/s | 1167095 | 5409 | 4.2× |
| JSONV2 | 2897670 | 184.00 MB/s | 745421 | 13288 | 2.2× |
| LightningDecodeAny | 3658543 | 145.73 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6454724 | 82.60 MB/s | 798693 | 17133 | 1.0× |
