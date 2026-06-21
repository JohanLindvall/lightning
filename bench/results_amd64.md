# JSON Deserialization Benchmarks

- generated 2026-06-21T16:42:03Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105478 | 1206.64 MB/s | 49760 | 3 | 12.4× |
| LightningDestructive | 105629 | 1204.92 MB/s | 49280 | 2 | 12.4× |
| Sonic | 202734 | 627.79 MB/s | 215064 | 15 | 6.4× |
| SonicFastest | 204346 | 622.84 MB/s | 215246 | 15 | 6.4× |
| Easyjson | 241449 | 527.13 MB/s | 122864 | 14 | 5.4× |
| Goccy | 247803 | 513.61 MB/s | 225093 | 884 | 5.3× |
| LightningDecodeAny | 472379 | 200.37 MB/s | 465586 | 9714 | 2.8× |
| JSONV2 | 476437 | 267.14 MB/s | 195127 | 1805 | 2.7× |
| Stdlib | 1307461 | 97.35 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4086776 | 550.81 MB/s | 3122872 | 3081 | 8.3× |
| Lightning | 4138641 | 543.91 MB/s | 3122875 | 3081 | 8.2× |
| SonicFastest | 4643079 | 484.82 MB/s | 4866459 | 2584 | 7.3× |
| Sonic | 4687767 | 480.20 MB/s | 4866728 | 2584 | 7.3× |
| LightningDecodeAny | 12532703 | 179.61 MB/s | 7938299 | 281383 | 2.7× |
| Goccy | 12971837 | 173.53 MB/s | 4195310 | 56535 | 2.6× |
| Easyjson | 13253332 | 169.85 MB/s | 3099809 | 2120 | 2.6× |
| JSONV2 | 17335756 | 129.85 MB/s | 3123189 | 3083 | 2.0× |
| Stdlib | 34063765 | 66.08 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 570065 | 474.34 MB/s | 348025 | 1627 | 7.7× |
| LightningDestructive | 574296 | 470.84 MB/s | 348025 | 1627 | 7.6× |
| Sonic | 751157 | 359.98 MB/s | 641741 | 1147 | 5.8× |
| SonicFastest | 751591 | 359.77 MB/s | 641802 | 1147 | 5.8× |
| LightningDecodeAny | 1749589 | 154.55 MB/s | 1011488 | 37901 | 2.5× |
| Easyjson | 1752914 | 154.26 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1774528 | 152.38 MB/s | 542060 | 8122 | 2.5× |
| JSONV2 | 2314239 | 116.84 MB/s | 348160 | 1628 | 1.9× |
| Stdlib | 4369327 | 61.89 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1289995 | 1338.92 MB/s | 959848 | 5881 | 13.7× |
| Lightning | 1338825 | 1290.09 MB/s | 959889 | 5882 | 13.2× |
| Sonic | 2111019 | 818.19 MB/s | 2695233 | 5547 | 8.3× |
| SonicFastest | 2112948 | 817.44 MB/s | 2695210 | 5547 | 8.3× |
| Goccy | 2921680 | 591.17 MB/s | 2581290 | 14603 | 6.0× |
| Easyjson | 4206705 | 410.58 MB/s | 972032 | 5389 | 4.2× |
| LightningDecodeAny | 4425346 | 113.05 MB/s | 4976204 | 81466 | 4.0× |
| JSONV2 | 4564946 | 378.36 MB/s | 1011614 | 7594 | 3.9× |
| Stdlib | 17612957 | 98.06 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1023 | 1771.93 MB/s | 0 | 0 | 15.6× |
| LightningDestructive | 1067 | 1697.54 MB/s | 0 | 0 | 15.0× |
| Easyjson | 2850 | 635.89 MB/s | 24 | 1 | 5.6× |
| Goccy | 3466 | 522.72 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6245 | 290.16 MB/s | 3344 | 38 | 2.6× |
| Sonic | 6461 | 280.45 MB/s | 3346 | 38 | 2.5× |
| JSONV2 | 8392 | 215.93 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9911 | 182.73 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 15994 | 113.29 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1102 | 1644.69 MB/s | 0 | 0 | 14.7× |
| LightningDestructive | 1163 | 1557.48 MB/s | 0 | 0 | 13.9× |
| Easyjson | 2855 | 634.57 MB/s | 24 | 1 | 5.7× |
| Goccy | 3436 | 527.39 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6097 | 297.22 MB/s | 3345 | 38 | 2.7× |
| Sonic | 6350 | 285.36 MB/s | 3345 | 38 | 2.5× |
| JSONV2 | 8352 | 216.97 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9466 | 191.32 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16170 | 112.06 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1320 | 1373.03 MB/s | 144 | 10 | 12.3× |
| LightningDestructive | 1442 | 1256.53 MB/s | 144 | 10 | 11.2× |
| Easyjson | 3229 | 561.25 MB/s | 144 | 10 | 5.0× |
| Goccy | 3230 | 560.93 MB/s | 2600 | 5 | 5.0× |
| SonicFastest | 6363 | 284.79 MB/s | 3368 | 40 | 2.5× |
| Sonic | 6528 | 277.56 MB/s | 3366 | 40 | 2.5× |
| JSONV2 | 8347 | 217.08 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9366 | 193.35 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16199 | 111.86 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 734 | 673.01 MB/s | 160 | 1 | 9.4× |
| LightningDestructive | 743 | 664.54 MB/s | 160 | 1 | 9.3× |
| Sonic | 1270 | 389.07 MB/s | 1076 | 8 | 5.4× |
| SonicFastest | 1271 | 388.80 MB/s | 1076 | 8 | 5.4× |
| LightningDecodeAny | 1718 | 286.98 MB/s | 1536 | 30 | 4.0× |
| Easyjson | 2509 | 196.85 MB/s | 448 | 3 | 2.8× |
| Goccy | 2813 | 175.58 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3302 | 149.62 MB/s | 528 | 7 | 2.1× |
| Stdlib | 6921 | 71.37 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 492 | 467.09 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 496 | 463.34 MB/s | 160 | 1 | 9.9× |
| Sonic | 927 | 248.13 MB/s | 800 | 8 | 5.3× |
| SonicFastest | 928 | 247.90 MB/s | 800 | 8 | 5.3× |
| LightningDecodeAny | 1555 | 147.24 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1691 | 136.01 MB/s | 448 | 3 | 2.9× |
| Goccy | 1832 | 125.57 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2566 | 89.63 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4896 | 46.98 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 91434 | 712.34 MB/s | 172434 | 107 | 7.3× |
| LightningDestructive | 98488 | 661.32 MB/s | 166213 | 102 | 6.8× |
| Sonic | 141740 | 459.52 MB/s | 235735 | 65 | 4.7× |
| SonicFastest | 144269 | 451.46 MB/s | 235860 | 65 | 4.6× |
| Goccy | 190348 | 342.17 MB/s | 228300 | 134 | 3.5× |
| LightningDecodeAny | 201410 | 264.78 MB/s | 176960 | 3252 | 3.3× |
| JSONV2 | 266448 | 244.45 MB/s | 206664 | 607 | 2.5× |
| Stdlib | 669360 | 97.30 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2508835 | 773.46 MB/s | 2846864 | 2238 | 11.3× |
| Lightning | 2581623 | 751.65 MB/s | 2846866 | 2238 | 11.0× |
| Sonic | 4667609 | 415.73 MB/s | 4875787 | 1736 | 6.1× |
| SonicFastest | 4728791 | 410.35 MB/s | 4880335 | 1736 | 6.0× |
| Goccy | 4915504 | 394.77 MB/s | 4063009 | 13509 | 5.8× |
| Easyjson | 8219186 | 236.09 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 9997534 | 194.10 MB/s | 7013526 | 219937 | 2.8× |
| JSONV2 | 12752178 | 152.17 MB/s | 3237192 | 13947 | 2.2× |
| Stdlib | 28275466 | 68.63 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1065933 | 3121.99 MB/s | 351704 | 1286 | 22.0× |
| Lightning | 1615431 | 2060.03 MB/s | 2488907 | 2995 | 14.5× |
| Sonic | 2015511 | 1651.11 MB/s | 5896506 | 4263 | 11.6× |
| SonicFastest | 2028114 | 1640.85 MB/s | 5896178 | 4263 | 11.6× |
| LightningDecodeAny | 3543240 | 867.50 MB/s | 4886621 | 56892 | 6.6× |
| Goccy | 4661436 | 713.91 MB/s | 3948916 | 3817 | 5.0× |
| JSONV2 | 7514466 | 442.86 MB/s | 5364511 | 13243 | 3.1× |
| Stdlib | 23433502 | 142.01 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 231484 | 951.88 MB/s | 136896 | 228 | 10.6× |
| LightningDestructive | 238560 | 923.65 MB/s | 136896 | 228 | 10.3× |
| Goccy | 480388 | 458.68 MB/s | 363955 | 1066 | 5.1× |
| SonicFastest | 490566 | 449.17 MB/s | 351083 | 262 | 5.0× |
| Sonic | 492089 | 447.78 MB/s | 351213 | 262 | 5.0× |
| Easyjson | 625178 | 352.45 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 760793 | 289.63 MB/s | 129747 | 470 | 3.2× |
| LightningDecodeAny | 1007017 | 107.56 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2446788 | 90.06 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14520844 | 557.82 MB/s | 15059835 | 51649 | 7.7× |
| Lightning | 15092561 | 536.69 MB/s | 15059843 | 51649 | 7.4× |
| SonicFastest | 18242466 | 444.02 MB/s | 19857810 | 41640 | 6.1× |
| Sonic | 18321916 | 442.10 MB/s | 19859205 | 41640 | 6.1× |
| Goccy | 26933551 | 300.74 MB/s | 19222838 | 107156 | 4.1× |
| Easyjson | 35200799 | 230.11 MB/s | 15059619 | 41643 | 3.2× |
| LightningDecodeAny | 43131737 | 120.63 MB/s | 34333348 | 912810 | 2.6× |
| JSONV2 | 48562951 | 166.79 MB/s | 15233722 | 78972 | 2.3× |
| Stdlib | 111642092 | 72.55 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5956471 | 500.88 MB/s | 3985336 | 30015 | 9.5× |
| Lightning | 6107143 | 488.52 MB/s | 3985337 | 30015 | 9.3× |
| Sonic | 9180992 | 324.96 MB/s | 9132088 | 57804 | 6.2× |
| SonicFastest | 9233834 | 323.10 MB/s | 9131614 | 57804 | 6.2× |
| Goccy | 18639970 | 160.06 MB/s | 9768721 | 273616 | 3.1× |
| Easyjson | 19009842 | 156.94 MB/s | 9479441 | 30115 | 3.0× |
| LightningDecodeAny | 20246265 | 90.59 MB/s | 20023836 | 408557 | 2.8× |
| JSONV2 | 26723658 | 111.64 MB/s | 9257034 | 86278 | 2.1× |
| Stdlib | 56860131 | 52.47 MB/s | 9258083 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1448251 | 499.64 MB/s | 907393 | 3618 | 9.6× |
| Lightning | 1522369 | 475.31 MB/s | 907387 | 3618 | 9.2× |
| SonicFastest | 2122187 | 340.97 MB/s | 2370898 | 3683 | 6.6× |
| Sonic | 2132752 | 339.28 MB/s | 2369796 | 3683 | 6.5× |
| Easyjson | 5351409 | 135.22 MB/s | 2847906 | 3698 | 2.6× |
| LightningDecodeAny | 5474262 | 118.84 MB/s | 5752202 | 80172 | 2.5× |
| Goccy | 5619289 | 128.77 MB/s | 2755793 | 80270 | 2.5× |
| JSONV2 | 6204321 | 116.63 MB/s | 2704703 | 7318 | 2.2× |
| Stdlib | 13939997 | 51.91 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2072985 | 760.91 MB/s | 907388 | 3618 | 9.5× |
| LightningDestructive | 2092317 | 753.88 MB/s | 907392 | 3618 | 9.4× |
| SonicFastest | 2402676 | 656.50 MB/s | 3225367 | 3683 | 8.2× |
| Sonic | 2409415 | 654.66 MB/s | 3225064 | 3683 | 8.2× |
| LightningDecodeAny | 4684620 | 160.82 MB/s | 5752197 | 80172 | 4.2× |
| Easyjson | 6584617 | 239.55 MB/s | 2847905 | 3698 | 3.0× |
| Goccy | 6757601 | 233.42 MB/s | 3490980 | 80262 | 2.9× |
| JSONV2 | 7064893 | 223.27 MB/s | 2704552 | 7318 | 2.8× |
| Stdlib | 19692641 | 80.10 MB/s | 2704547 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 248806 | 603.38 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 255344 | 587.93 MB/s | 81920 | 1 | 8.4× |
| SonicFastest | 374816 | 400.53 MB/s | 407394 | 16 | 5.7× |
| Sonic | 380083 | 394.98 MB/s | 407502 | 16 | 5.6× |
| LightningDecodeAny | 591533 | 253.78 MB/s | 746006 | 10020 | 3.6× |
| Goccy | 1023669 | 146.65 MB/s | 325045 | 10005 | 2.1× |
| JSONV2 | 1091854 | 137.49 MB/s | 357724 | 20 | 2.0× |
| Stdlib | 2145542 | 69.97 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 35919 | 782.80 MB/s | 30272 | 105 | 9.5× |
| LightningDestructive | 36469 | 770.99 MB/s | 30144 | 103 | 9.4× |
| SonicFastest | 57672 | 487.53 MB/s | 59441 | 83 | 5.9× |
| Sonic | 58059 | 484.28 MB/s | 59463 | 83 | 5.9× |
| Easyjson | 79007 | 355.88 MB/s | 32304 | 138 | 4.3× |
| Goccy | 82410 | 341.19 MB/s | 59254 | 188 | 4.1× |
| JSONV2 | 143073 | 196.52 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 166969 | 168.40 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 341891 | 82.24 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1997 | 1165.63 MB/s | 32 | 1 | 13.1× |
| LightningDestructive | 2057 | 1131.90 MB/s | 32 | 1 | 12.7× |
| Sonic | 4767 | 488.32 MB/s | 3713 | 4 | 5.5× |
| SonicFastest | 4780 | 487.01 MB/s | 3714 | 4 | 5.5× |
| Goccy | 4810 | 484.03 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 5427 | 429.00 MB/s | 192 | 2 | 4.8× |
| JSONV2 | 8765 | 265.62 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 10405 | 161.95 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 26065 | 89.31 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 214 | 882.21 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 223 | 849.05 MB/s | 0 | 0 | 12.5× |
| Goccy | 445 | 424.41 MB/s | 304 | 2 | 6.2× |
| Easyjson | 552 | 342.68 MB/s | 0 | 0 | 5.0× |
| Sonic | 643 | 293.96 MB/s | 341 | 3 | 4.3× |
| SonicFastest | 651 | 290.52 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1072 | 176.34 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1347 | 99.46 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2782 | 67.95 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1458 | 1503.07 MB/s | 0 | 0 | 12.9× |
| LightningDestructive | 1519 | 1442.09 MB/s | 0 | 0 | 12.3× |
| Easyjson | 3512 | 623.80 MB/s | 24 | 1 | 5.3× |
| Goccy | 3868 | 566.50 MB/s | 2864 | 4 | 4.8× |
| SonicFastest | 6581 | 332.93 MB/s | 3599 | 38 | 2.8× |
| Sonic | 6716 | 326.23 MB/s | 3601 | 38 | 2.8× |
| JSONV2 | 8511 | 257.43 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9427 | 192.11 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18752 | 116.84 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 719595 | 709.39 MB/s | 703778 | 1012 | 9.9× |
| Lightning | 764734 | 667.52 MB/s | 703778 | 1012 | 9.3× |
| SonicFastest | 1239939 | 411.69 MB/s | 1306878 | 2014 | 5.7× |
| Goccy | 1244211 | 410.28 MB/s | 1139860 | 5006 | 5.7× |
| Sonic | 1250637 | 408.17 MB/s | 1309367 | 2014 | 5.7× |
| Easyjson | 1796487 | 284.15 MB/s | 863778 | 3012 | 4.0× |
| JSONV2 | 3577173 | 142.70 MB/s | 1075963 | 12645 | 2.0× |
| LightningDecodeAny | 3639191 | 126.80 MB/s | 2785927 | 66022 | 2.0× |
| Stdlib | 7129349 | 71.60 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 544 | 36388.87 MB/s | 0 | 0 | 297.7× |
| LightningDestructive | 800 | 24738.96 MB/s | 0 | 0 | 202.4× |
| SonicFastest | 6959 | 2843.69 MB/s | 21077 | 3 | 23.3× |
| Goccy | 23295 | 849.49 MB/s | 20492 | 2 | 7.0× |
| Sonic | 31722 | 623.82 MB/s | 20595 | 3 | 5.1× |
| JSONV2 | 33773 | 585.94 MB/s | 8 | 1 | 4.8× |
| LightningDecodeAny | 100805 | 196.30 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 105091 | 188.30 MB/s | 0 | 0 | 1.5× |
| Stdlib | 161916 | 122.22 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2276 | 7963.34 MB/s | 432 | 2 | 52.5× |
| LightningDestructive | 2590 | 6996.85 MB/s | 0 | 0 | 46.1× |
| Easyjson | 4744 | 3820.11 MB/s | 432 | 2 | 25.2× |
| SonicFastest | 8321 | 2178.06 MB/s | 20462 | 5 | 14.4× |
| Sonic | 8774 | 2065.58 MB/s | 20411 | 5 | 13.6× |
| LightningDecodeAny | 19283 | 927.33 MB/s | 29088 | 191 | 6.2× |
| Goccy | 24673 | 734.57 MB/s | 19460 | 2 | 4.8× |
| JSONV2 | 49217 | 368.25 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 119481 | 151.69 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2275899 | 882.51 MB/s | 2960389 | 7411 | 9.6× |
| Lightning | 2460043 | 816.45 MB/s | 2962102 | 7417 | 8.9× |
| Sonic | 4104438 | 489.35 MB/s | 5153024 | 7085 | 5.3× |
| SonicFastest | 4111976 | 488.45 MB/s | 5154044 | 7085 | 5.3× |
| Goccy | 4760423 | 421.92 MB/s | 5411402 | 15833 | 4.6× |
| Easyjson | 5459962 | 367.86 MB/s | 2981487 | 7439 | 4.0× |
| LightningDecodeAny | 6976819 | 163.73 MB/s | 7386073 | 134751 | 3.1× |
| JSONV2 | 7782013 | 258.09 MB/s | 3173671 | 14562 | 2.8× |
| Stdlib | 21870507 | 91.84 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 843 | 650.91 MB/s | 480 | 1 | 8.0× |
| LightningDestructive | 847 | 648.21 MB/s | 480 | 1 | 7.9× |
| LightningDecodeAny | 2139 | 256.19 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 2262 | 242.71 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2362 | 232.45 MB/s | 2261 | 8 | 2.8× |
| Sonic | 2383 | 230.38 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3391 | 161.89 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3534 | 155.37 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6729 | 81.59 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 499951 | 1263.15 MB/s | 364473 | 566 | 12.6× |
| Lightning | 570696 | 1106.57 MB/s | 413001 | 878 | 11.0× |
| SonicFastest | 1037332 | 608.79 MB/s | 1072551 | 814 | 6.1× |
| Sonic | 1042920 | 605.52 MB/s | 1072701 | 814 | 6.0× |
| Goccy | 1261198 | 500.73 MB/s | 986806 | 1200 | 5.0× |
| Easyjson | 1391617 | 453.80 MB/s | 422504 | 936 | 4.5× |
| JSONV2 | 2332437 | 270.75 MB/s | 571591 | 3144 | 2.7× |
| LightningDecodeAny | 2512763 | 185.81 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6300347 | 100.23 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 817260 | 688.16 MB/s | 544249 | 448 | 7.4× |
| Lightning | 1015195 | 553.99 MB/s | 767621 | 1254 | 5.9× |
| Sonic | 1243895 | 452.13 MB/s | 1346682 | 1185 | 4.9× |
| SonicFastest | 1249663 | 450.05 MB/s | 1346897 | 1185 | 4.8× |
| Goccy | 1478168 | 380.48 MB/s | 1036956 | 1028 | 4.1× |
| Easyjson | 2146514 | 262.01 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 3092944 | 181.84 MB/s | 2114150 | 30295 | 2.0× |
| JSONV2 | 3198481 | 175.84 MB/s | 927409 | 3482 | 1.9× |
| Stdlib | 6035539 | 93.18 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 602512 | 884.93 MB/s | 333416 | 2084 | 10.4× |
| Lightning | 696599 | 765.40 MB/s | 368224 | 2293 | 9.0× |
| SonicFastest | 1121674 | 475.34 MB/s | 982812 | 3082 | 5.6× |
| Sonic | 1124981 | 473.94 MB/s | 983142 | 3082 | 5.6× |
| Easyjson | 1265786 | 421.22 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1508706 | 353.40 MB/s | 1167085 | 5409 | 4.1× |
| JSONV2 | 2898697 | 183.94 MB/s | 745420 | 13288 | 2.2× |
| LightningDecodeAny | 3450800 | 154.51 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6247625 | 85.34 MB/s | 798693 | 17133 | 1.0× |
