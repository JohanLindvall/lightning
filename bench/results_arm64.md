# JSON Deserialization Benchmarks

- generated 2026-06-20T16:46:40Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 107618 | 1182.65 MB/s | 50208 | 4 | 10.3× |
| Sonic | 185239 | 687.09 MB/s | 195839 | 10 | 6.0× |
| SonicFastest | 186338 | 683.03 MB/s | 196795 | 10 | 5.9× |
| Goccy | 197765 | 643.57 MB/s | 224807 | 884 | 5.6× |
| Easyjson | 212066 | 600.17 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 418604 | 304.05 MB/s | 195121 | 1805 | 2.6× |
| LightningDecodeAny | 477689 | 198.15 MB/s | 466034 | 9715 | 2.3× |
| Stdlib | 1107562 | 114.91 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3779473 | 595.60 MB/s | 3122876 | 3081 | 6.9× |
| Sonic | 4535655 | 496.30 MB/s | 15258526 | 970 | 5.7× |
| SonicFastest | 4591582 | 490.26 MB/s | 15238766 | 970 | 5.7× |
| Goccy | 10578055 | 212.80 MB/s | 4132265 | 56533 | 2.5× |
| Easyjson | 11236191 | 200.34 MB/s | 3099809 | 2120 | 2.3× |
| LightningDecodeAny | 11724835 | 191.99 MB/s | 7938299 | 281383 | 2.2× |
| JSONV2 | 16304800 | 138.06 MB/s | 3123207 | 3083 | 1.6× |
| Stdlib | 26071906 | 86.34 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 484588 | 558.01 MB/s | 348024 | 1627 | 6.9× |
| Sonic | 627016 | 431.25 MB/s | 485504 | 968 | 5.3× |
| SonicFastest | 635051 | 425.80 MB/s | 496540 | 968 | 5.3× |
| Easyjson | 1418580 | 190.62 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1426029 | 189.62 MB/s | 543930 | 8123 | 2.4× |
| LightningDecodeAny | 1598930 | 169.11 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2111105 | 128.09 MB/s | 348148 | 1628 | 1.6× |
| Stdlib | 3353449 | 80.63 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1406324 | 1228.17 MB/s | 964642 | 6177 | 9.5× |
| Sonic | 2030878 | 850.47 MB/s | 2692006 | 4020 | 6.6× |
| SonicFastest | 2042730 | 845.54 MB/s | 2720318 | 4020 | 6.5× |
| Goccy | 2343010 | 737.17 MB/s | 2582823 | 14604 | 5.7× |
| Easyjson | 4257080 | 405.73 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4321259 | 399.70 MB/s | 1011632 | 7594 | 3.1× |
| LightningDecodeAny | 4745070 | 105.44 MB/s | 4976253 | 81467 | 2.8× |
| Stdlib | 13357020 | 129.31 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1239 | 1462.63 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2528 | 716.67 MB/s | 24 | 1 | 5.6× |
| Goccy | 2827 | 640.88 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6011 | 301.44 MB/s | 3759 | 40 | 2.3× |
| Sonic | 6022 | 300.92 MB/s | 3756 | 40 | 2.3× |
| JSONV2 | 7865 | 230.38 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8359 | 216.66 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14061 | 128.87 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1248 | 1452.49 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2538 | 713.82 MB/s | 24 | 1 | 5.6× |
| Goccy | 2826 | 641.17 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5977 | 303.15 MB/s | 3677 | 40 | 2.4× |
| SonicFastest | 6016 | 301.22 MB/s | 3725 | 40 | 2.3× |
| JSONV2 | 7784 | 232.77 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8379 | 216.15 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14106 | 128.45 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1449 | 1250.33 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2757 | 657.26 MB/s | 144 | 10 | 5.1× |
| Goccy | 2865 | 632.39 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6223 | 291.19 MB/s | 3766 | 42 | 2.3× |
| SonicFastest | 6247 | 290.07 MB/s | 3847 | 42 | 2.2× |
| JSONV2 | 7869 | 230.26 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8349 | 216.92 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14025 | 129.20 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 758 | 651.43 MB/s | 160 | 1 | 7.4× |
| Sonic | 1247 | 396.22 MB/s | 966 | 6 | 4.5× |
| SonicFastest | 1255 | 393.64 MB/s | 979 | 6 | 4.4× |
| LightningDecodeAny | 1642 | 300.27 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2220 | 222.53 MB/s | 448 | 3 | 2.5× |
| Goccy | 2454 | 201.34 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3215 | 153.65 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5574 | 88.62 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 477 | 481.78 MB/s | 160 | 1 | 8.7× |
| SonicFastest | 906 | 253.98 MB/s | 685 | 6 | 4.6× |
| Sonic | 911 | 252.50 MB/s | 693 | 6 | 4.5× |
| Easyjson | 1403 | 163.88 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1464 | 156.40 MB/s | 1536 | 30 | 2.8× |
| Goccy | 1626 | 141.41 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2420 | 95.06 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4141 | 55.54 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 79706 | 817.15 MB/s | 178609 | 112 | 6.9× |
| SonicFastest | 101163 | 643.83 MB/s | 159986 | 75 | 5.4× |
| Sonic | 101346 | 642.67 MB/s | 160042 | 75 | 5.4× |
| Goccy | 147867 | 440.48 MB/s | 229376 | 134 | 3.7× |
| LightningDecodeAny | 190808 | 279.49 MB/s | 183136 | 3257 | 2.9× |
| JSONV2 | 226424 | 287.66 MB/s | 206654 | 607 | 2.4× |
| Stdlib | 546746 | 119.13 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2748241 | 706.08 MB/s | 2846865 | 2238 | 8.5× |
| SonicFastest | 4549412 | 426.53 MB/s | 14608574 | 1407 | 5.2× |
| Sonic | 4669044 | 415.60 MB/s | 14608593 | 1407 | 5.0× |
| Goccy | 4830693 | 401.70 MB/s | 4064522 | 13510 | 4.9× |
| Easyjson | 7522768 | 257.95 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9573370 | 202.69 MB/s | 7013524 | 219937 | 2.4× |
| JSONV2 | 11247274 | 172.53 MB/s | 3237214 | 13947 | 2.1× |
| Stdlib | 23445638 | 82.76 MB/s | 3551333 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2300683 | 1446.45 MB/s | 4611512 | 5958 | 9.1× |
| Sonic | 2730156 | 1218.92 MB/s | 6450073 | 4248 | 7.7× |
| SonicFastest | 2737690 | 1215.56 MB/s | 6400984 | 4248 | 7.6× |
| LightningDecodeAny | 4457515 | 689.57 MB/s | 7005135 | 58601 | 4.7× |
| Goccy | 4478683 | 743.04 MB/s | 3948907 | 3816 | 4.7× |
| JSONV2 | 7437368 | 447.45 MB/s | 5364515 | 13243 | 2.8× |
| Stdlib | 20903785 | 159.20 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 254348 | 866.32 MB/s | 123584 | 225 | 8.0× |
| Sonic | 377106 | 584.31 MB/s | 296254 | 398 | 5.4× |
| SonicFastest | 378993 | 581.40 MB/s | 297924 | 398 | 5.4× |
| Goccy | 430792 | 511.49 MB/s | 363999 | 1067 | 4.7× |
| Easyjson | 548014 | 402.08 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 727093 | 303.05 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 888900 | 121.85 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2039594 | 108.03 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13230421 | 612.23 MB/s | 15059843 | 51649 | 6.7× |
| Sonic | 16753799 | 483.47 MB/s | 70930819 | 40014 | 5.3× |
| SonicFastest | 16798933 | 482.18 MB/s | 70901598 | 40014 | 5.3× |
| Goccy | 23849334 | 339.63 MB/s | 16968274 | 107148 | 3.7× |
| Easyjson | 30911524 | 262.04 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 41243544 | 126.15 MB/s | 34333353 | 912810 | 2.2× |
| JSONV2 | 43818175 | 184.86 MB/s | 15233722 | 78972 | 2.0× |
| Stdlib | 88986299 | 91.03 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6282352 | 474.90 MB/s | 3985336 | 30015 | 7.5× |
| Sonic | 8733080 | 341.63 MB/s | 26486801 | 56760 | 5.4× |
| SonicFastest | 8795329 | 339.21 MB/s | 26564563 | 56760 | 5.4× |
| Easyjson | 16604425 | 179.68 MB/s | 9479440 | 30115 | 2.8× |
| Goccy | 16845724 | 177.11 MB/s | 10641370 | 273650 | 2.8× |
| LightningDecodeAny | 19066493 | 96.20 MB/s | 20023836 | 408557 | 2.5× |
| JSONV2 | 23726118 | 125.75 MB/s | 9257118 | 86278 | 2.0× |
| Stdlib | 47152403 | 63.27 MB/s | 9258116 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1370373 | 528.03 MB/s | 907388 | 3618 | 8.4× |
| SonicFastest | 1740319 | 415.78 MB/s | 3175415 | 7226 | 6.6× |
| Sonic | 1751502 | 413.13 MB/s | 3180717 | 7226 | 6.6× |
| Easyjson | 4170987 | 173.48 MB/s | 2847905 | 3698 | 2.8× |
| LightningDecodeAny | 4279150 | 152.03 MB/s | 5752200 | 80172 | 2.7× |
| Goccy | 4790465 | 151.05 MB/s | 2801351 | 80273 | 2.4× |
| JSONV2 | 5763686 | 125.54 MB/s | 2704627 | 7318 | 2.0× |
| Stdlib | 11566377 | 62.56 MB/s | 2704553 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1958818 | 805.26 MB/s | 907385 | 3618 | 8.0× |
| Sonic | 2233185 | 706.32 MB/s | 5792178 | 7226 | 7.0× |
| SonicFastest | 2238624 | 704.61 MB/s | 5785945 | 7226 | 7.0× |
| LightningDecodeAny | 3818611 | 197.30 MB/s | 5752203 | 80172 | 4.1× |
| Easyjson | 5533959 | 285.03 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5616892 | 280.82 MB/s | 3543071 | 80265 | 2.8× |
| JSONV2 | 6321229 | 249.53 MB/s | 2704595 | 7318 | 2.5× |
| Stdlib | 15710526 | 100.40 MB/s | 2704558 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224029 | 670.11 MB/s | 81920 | 1 | 8.2× |
| SonicFastest | 270921 | 554.13 MB/s | 260596 | 6 | 6.8× |
| Sonic | 271601 | 552.74 MB/s | 259811 | 6 | 6.8× |
| LightningDecodeAny | 477427 | 314.44 MB/s | 746005 | 10020 | 3.9× |
| Goccy | 859978 | 174.57 MB/s | 324051 | 10004 | 2.1× |
| JSONV2 | 1090382 | 137.68 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1842321 | 81.49 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 38138 | 737.25 MB/s | 29104 | 107 | 8.0× |
| Sonic | 63020 | 446.16 MB/s | 46742 | 103 | 4.8× |
| SonicFastest | 63240 | 444.61 MB/s | 47009 | 103 | 4.8× |
| Easyjson | 68146 | 412.60 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71311 | 394.29 MB/s | 59207 | 188 | 4.3× |
| JSONV2 | 134024 | 209.79 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 153289 | 183.43 MB/s | 135136 | 2680 | 2.0× |
| Stdlib | 303387 | 92.68 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2039 | 1141.75 MB/s | 32 | 1 | 11.1× |
| Goccy | 4117 | 565.47 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4214 | 552.46 MB/s | 192 | 2 | 5.4× |
| Sonic | 5026 | 463.21 MB/s | 4224 | 6 | 4.5× |
| SonicFastest | 5041 | 461.82 MB/s | 4230 | 6 | 4.5× |
| JSONV2 | 8460 | 275.19 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10467 | 160.98 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22724 | 102.45 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 229 | 824.99 MB/s | 0 | 0 | 10.9× |
| Goccy | 382 | 494.07 MB/s | 304 | 2 | 6.5× |
| Easyjson | 486 | 389.32 MB/s | 0 | 0 | 5.1× |
| Sonic | 800 | 236.10 MB/s | 498 | 4 | 3.1× |
| SonicFastest | 801 | 235.91 MB/s | 506 | 4 | 3.1× |
| JSONV2 | 1040 | 181.69 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1266 | 105.82 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2491 | 75.87 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1541 | 1422.13 MB/s | 0 | 0 | 10.4× |
| Easyjson | 3190 | 686.86 MB/s | 24 | 1 | 5.0× |
| Goccy | 3205 | 683.55 MB/s | 2864 | 4 | 5.0× |
| Sonic | 6383 | 343.27 MB/s | 3969 | 40 | 2.5× |
| SonicFastest | 6403 | 342.18 MB/s | 4001 | 40 | 2.5× |
| JSONV2 | 8062 | 271.77 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8281 | 218.68 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16022 | 136.75 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 919401 | 555.23 MB/s | 703780 | 1012 | 6.6× |
| Goccy | 1156212 | 441.51 MB/s | 1140232 | 5006 | 5.2× |
| SonicFastest | 1156322 | 441.47 MB/s | 911334 | 2006 | 5.2× |
| Sonic | 1170848 | 435.99 MB/s | 930852 | 2006 | 5.2× |
| Easyjson | 1531503 | 333.32 MB/s | 863777 | 3012 | 4.0× |
| JSONV2 | 3207747 | 159.14 MB/s | 1076005 | 12646 | 1.9× |
| LightningDecodeAny | 3583940 | 128.76 MB/s | 2785926 | 66022 | 1.7× |
| Stdlib | 6062162 | 84.21 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2063 | 9594.37 MB/s | 0 | 0 | 60.4× |
| Goccy | 19862 | 996.31 MB/s | 20491 | 2 | 6.3× |
| SonicFastest | 27838 | 710.86 MB/s | 21886 | 4 | 4.5× |
| Sonic | 27846 | 710.67 MB/s | 22015 | 4 | 4.5× |
| JSONV2 | 29760 | 664.94 MB/s | 8 | 1 | 4.2× |
| LightningDecodeAny | 77970 | 253.79 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 86079 | 229.89 MB/s | 0 | 0 | 1.4× |
| Stdlib | 124577 | 158.85 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2933 | 6179.40 MB/s | 864 | 4 | 35.2× |
| Easyjson | 3981 | 4552.11 MB/s | 432 | 2 | 25.9× |
| SonicFastest | 10071 | 1799.68 MB/s | 22921 | 6 | 10.3× |
| Sonic | 10199 | 1777.07 MB/s | 23197 | 6 | 10.1× |
| Goccy | 16253 | 1115.09 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 17196 | 1039.87 MB/s | 29520 | 193 | 6.0× |
| JSONV2 | 46809 | 387.19 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 103280 | 175.48 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2379403 | 844.12 MB/s | 2963981 | 7453 | 7.8× |
| Goccy | 4238057 | 473.92 MB/s | 5412280 | 15831 | 4.4× |
| Sonic | 4485003 | 447.82 MB/s | 10936662 | 13683 | 4.2× |
| SonicFastest | 4509356 | 445.41 MB/s | 10909472 | 13683 | 4.1× |
| Easyjson | 4891524 | 410.61 MB/s | 2981481 | 7439 | 3.8× |
| JSONV2 | 6854955 | 293.00 MB/s | 3173690 | 14563 | 2.7× |
| LightningDecodeAny | 7464816 | 153.02 MB/s | 7387754 | 134757 | 2.5× |
| Stdlib | 18646633 | 107.71 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 927 | 591.98 MB/s | 480 | 1 | 6.2× |
| LightningDecodeAny | 2116 | 258.95 MB/s | 2261 | 50 | 2.7× |
| Easyjson | 2222 | 247.06 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2771 | 198.11 MB/s | 2048 | 26 | 2.1× |
| SonicFastest | 2775 | 197.87 MB/s | 2048 | 26 | 2.1× |
| Goccy | 3169 | 173.26 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3358 | 163.50 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5740 | 95.65 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 589675 | 1070.95 MB/s | 461113 | 1230 | 9.2× |
| Sonic | 1008950 | 625.91 MB/s | 1004017 | 1102 | 5.4× |
| SonicFastest | 1017366 | 620.73 MB/s | 1006955 | 1102 | 5.3× |
| Easyjson | 1140824 | 553.56 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1164560 | 542.28 MB/s | 988841 | 1202 | 4.6× |
| JSONV2 | 2143964 | 294.55 MB/s | 571611 | 3144 | 2.5× |
| LightningDecodeAny | 2437886 | 191.52 MB/s | 2058071 | 30607 | 2.2× |
| Stdlib | 5401292 | 116.92 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 971052 | 579.17 MB/s | 886433 | 2062 | 5.4× |
| SonicFastest | 1025444 | 548.45 MB/s | 943846 | 1476 | 5.1× |
| Sonic | 1034153 | 543.83 MB/s | 962733 | 1476 | 5.1× |
| Goccy | 1323581 | 424.91 MB/s | 1039325 | 1030 | 4.0× |
| Easyjson | 1748644 | 321.63 MB/s | 775154 | 1254 | 3.0× |
| JSONV2 | 2771383 | 202.93 MB/s | 927439 | 3482 | 1.9× |
| LightningDecodeAny | 2863106 | 196.43 MB/s | 2232992 | 31103 | 1.8× |
| Stdlib | 5253883 | 107.05 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 805473 | 661.94 MB/s | 413026 | 3155 | 6.8× |
| Sonic | 1128145 | 472.61 MB/s | 1038030 | 4351 | 4.9× |
| SonicFastest | 1132090 | 470.97 MB/s | 1033202 | 4351 | 4.9× |
| Easyjson | 1132787 | 470.68 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1297432 | 410.95 MB/s | 1167255 | 5409 | 4.3× |
| JSONV2 | 2581781 | 206.52 MB/s | 745463 | 13288 | 2.1× |
| LightningDecodeAny | 3466829 | 153.79 MB/s | 2709157 | 50805 | 1.6× |
| Stdlib | 5514172 | 96.69 MB/s | 798692 | 17133 | 1.0× |
