# JSON Deserialization Benchmarks

- generated 2026-06-23T18:19:13Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104394 | 1219.18 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 105276 | 1208.97 MB/s | 49280 | 2 | 10.5× |
| Sonic | 178482 | 713.10 MB/s | 191419 | 10 | 6.2× |
| SonicFastest | 183047 | 695.31 MB/s | 198788 | 10 | 6.0× |
| Goccy | 191415 | 664.92 MB/s | 224650 | 884 | 5.8× |
| Easyjson | 211398 | 602.06 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 419436 | 303.44 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 448087 | 211.24 MB/s | 465586 | 9714 | 2.5× |
| Stdlib | 1104372 | 115.25 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3783558 | 594.96 MB/s | 3122875 | 3081 | 6.9× |
| LightningDestructive | 3786830 | 594.44 MB/s | 3122872 | 3081 | 6.9× |
| SonicFastest | 4573567 | 492.19 MB/s | 15248680 | 970 | 5.7× |
| Sonic | 4645861 | 484.53 MB/s | 15233812 | 970 | 5.6× |
| Goccy | 10296588 | 218.62 MB/s | 4118054 | 56532 | 2.5× |
| Easyjson | 11060576 | 203.52 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11374016 | 197.91 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16044540 | 140.30 MB/s | 3123222 | 3083 | 1.6× |
| Stdlib | 26097414 | 86.26 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 488212 | 553.86 MB/s | 348024 | 1627 | 6.9× |
| LightningDestructive | 494532 | 546.79 MB/s | 348024 | 1627 | 6.8× |
| Sonic | 617264 | 438.07 MB/s | 462173 | 968 | 5.4× |
| SonicFastest | 619064 | 436.79 MB/s | 460445 | 968 | 5.4× |
| Goccy | 1385087 | 195.22 MB/s | 542321 | 8122 | 2.4× |
| Easyjson | 1405500 | 192.39 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1598534 | 169.16 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2096849 | 128.96 MB/s | 348150 | 1628 | 1.6× |
| Stdlib | 3350043 | 80.72 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1349062 | 1280.30 MB/s | 959848 | 5881 | 9.8× |
| Lightning | 1365560 | 1264.83 MB/s | 959890 | 5882 | 9.7× |
| Sonic | 2028779 | 851.35 MB/s | 2761007 | 4020 | 6.5× |
| SonicFastest | 2042484 | 845.64 MB/s | 2748381 | 4020 | 6.5× |
| Goccy | 2431982 | 710.20 MB/s | 2584977 | 14605 | 5.4× |
| Easyjson | 4219270 | 409.36 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4241169 | 407.25 MB/s | 1011634 | 7594 | 3.1× |
| LightningDecodeAny | 4629299 | 108.07 MB/s | 4976205 | 81466 | 2.9× |
| Stdlib | 13243396 | 130.42 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1194 | 1517.93 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1199 | 1511.62 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2545 | 712.04 MB/s | 24 | 1 | 5.5× |
| Goccy | 2834 | 639.28 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5932 | 305.48 MB/s | 3720 | 40 | 2.4× |
| SonicFastest | 5939 | 305.12 MB/s | 3755 | 40 | 2.4× |
| JSONV2 | 7785 | 232.77 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8080 | 224.13 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14078 | 128.71 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1219 | 1486.17 MB/s | 0 | 0 | 11.6× |
| LightningDestructive | 1238 | 1463.52 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2537 | 714.16 MB/s | 24 | 1 | 5.6× |
| Goccy | 2818 | 642.95 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 5891 | 307.59 MB/s | 3717 | 40 | 2.4× |
| Sonic | 5959 | 304.09 MB/s | 3732 | 40 | 2.4× |
| JSONV2 | 7765 | 233.37 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8109 | 223.33 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14110 | 128.42 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1407 | 1288.01 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1453 | 1246.75 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2776 | 652.79 MB/s | 144 | 10 | 5.1× |
| Goccy | 2864 | 632.76 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6083 | 297.86 MB/s | 3759 | 42 | 2.3× |
| Sonic | 6092 | 297.45 MB/s | 3770 | 42 | 2.3× |
| JSONV2 | 7931 | 228.46 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8113 | 223.23 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14086 | 128.64 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 736 | 670.93 MB/s | 160 | 1 | 7.6× |
| Lightning | 742 | 666.24 MB/s | 160 | 1 | 7.5× |
| Sonic | 1216 | 406.41 MB/s | 972 | 6 | 4.6× |
| SonicFastest | 1217 | 405.78 MB/s | 977 | 6 | 4.6× |
| LightningDecodeAny | 1636 | 301.42 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2237 | 220.86 MB/s | 448 | 3 | 2.5× |
| Goccy | 2426 | 203.65 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3220 | 153.40 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5567 | 88.74 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 475 | 484.57 MB/s | 160 | 1 | 8.8× |
| LightningDestructive | 477 | 482.48 MB/s | 160 | 1 | 8.8× |
| SonicFastest | 894 | 257.22 MB/s | 692 | 6 | 4.7× |
| Sonic | 896 | 256.60 MB/s | 693 | 6 | 4.7× |
| LightningDecodeAny | 1426 | 160.54 MB/s | 1536 | 30 | 2.9× |
| Easyjson | 1428 | 161.08 MB/s | 448 | 3 | 2.9× |
| Goccy | 1620 | 142.02 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2384 | 96.47 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4173 | 55.12 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 71196 | 914.82 MB/s | 166212 | 102 | 7.6× |
| Lightning | 72267 | 901.27 MB/s | 172432 | 107 | 7.5× |
| Sonic | 93625 | 695.67 MB/s | 154773 | 75 | 5.8× |
| SonicFastest | 94283 | 690.81 MB/s | 155098 | 75 | 5.7× |
| Goccy | 139705 | 466.21 MB/s | 229291 | 134 | 3.9× |
| LightningDecodeAny | 182745 | 291.82 MB/s | 176960 | 3252 | 3.0× |
| JSONV2 | 218706 | 297.81 MB/s | 206651 | 607 | 2.5× |
| Stdlib | 540986 | 120.40 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2583222 | 751.18 MB/s | 2846864 | 2238 | 9.1× |
| Lightning | 2668985 | 727.04 MB/s | 2846867 | 2238 | 8.8× |
| Sonic | 4529694 | 428.39 MB/s | 14608582 | 1407 | 5.2× |
| SonicFastest | 4612301 | 420.72 MB/s | 14608587 | 1407 | 5.1× |
| Goccy | 4777093 | 406.20 MB/s | 4064625 | 13510 | 4.9× |
| Easyjson | 7486100 | 259.21 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9831276 | 197.38 MB/s | 7013526 | 219937 | 2.4× |
| JSONV2 | 11276789 | 172.08 MB/s | 3237225 | 13947 | 2.1× |
| Stdlib | 23440749 | 82.78 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1106364 | 3007.90 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1756219 | 1894.88 MB/s | 2488905 | 2995 | 11.8× |
| Sonic | 2633787 | 1263.52 MB/s | 6492512 | 4248 | 7.9× |
| SonicFastest | 2652869 | 1254.43 MB/s | 6497080 | 4248 | 7.8× |
| LightningDecodeAny | 3688473 | 833.34 MB/s | 4886620 | 56892 | 5.6× |
| Goccy | 4518285 | 736.53 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7270787 | 457.70 MB/s | 5364515 | 13243 | 2.9× |
| Stdlib | 20798137 | 160.01 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221110 | 996.55 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 221286 | 995.75 MB/s | 136898 | 228 | 9.2× |
| SonicFastest | 395749 | 556.78 MB/s | 351110 | 398 | 5.1× |
| Sonic | 399291 | 551.84 MB/s | 354165 | 398 | 5.1× |
| Goccy | 429580 | 512.93 MB/s | 364570 | 1067 | 4.7× |
| Easyjson | 544391 | 404.76 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 732528 | 300.80 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 862587 | 125.57 MB/s | 861393 | 11944 | 2.4× |
| Stdlib | 2035774 | 108.24 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12741746 | 635.71 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13086871 | 618.94 MB/s | 15059840 | 51649 | 6.8× |
| Sonic | 16617856 | 487.43 MB/s | 70887292 | 40014 | 5.4× |
| SonicFastest | 16704038 | 484.92 MB/s | 70887390 | 40014 | 5.3× |
| Goccy | 23294802 | 347.72 MB/s | 16785296 | 107147 | 3.8× |
| Easyjson | 30852915 | 262.54 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 41941180 | 124.06 MB/s | 34333352 | 912810 | 2.1× |
| JSONV2 | 44016586 | 184.02 MB/s | 15233760 | 78972 | 2.0× |
| Stdlib | 89186753 | 90.82 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6070378 | 491.48 MB/s | 3985336 | 30015 | 7.7× |
| Lightning | 6159278 | 484.39 MB/s | 3985339 | 30015 | 7.6× |
| SonicFastest | 8631808 | 345.64 MB/s | 26691482 | 56760 | 5.4× |
| Sonic | 8642013 | 345.23 MB/s | 26673959 | 56760 | 5.4× |
| Easyjson | 16555463 | 180.21 MB/s | 9479440 | 30115 | 2.8× |
| Goccy | 16612462 | 179.59 MB/s | 10678068 | 273652 | 2.8× |
| LightningDecodeAny | 19130089 | 95.88 MB/s | 20023838 | 408557 | 2.5× |
| JSONV2 | 24057200 | 124.02 MB/s | 9257165 | 86278 | 1.9× |
| Stdlib | 46907393 | 63.60 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1363600 | 530.65 MB/s | 907393 | 3618 | 8.5× |
| Lightning | 1368918 | 528.59 MB/s | 907388 | 3618 | 8.4× |
| Sonic | 1735259 | 417.00 MB/s | 3178551 | 7226 | 6.7× |
| SonicFastest | 1735894 | 416.84 MB/s | 3187545 | 7226 | 6.7× |
| Easyjson | 4240517 | 170.64 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4285040 | 151.82 MB/s | 5752202 | 80172 | 2.7× |
| Goccy | 4767434 | 151.78 MB/s | 2823815 | 80275 | 2.4× |
| JSONV2 | 5637146 | 128.36 MB/s | 2704619 | 7318 | 2.0× |
| Stdlib | 11545169 | 62.68 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1940111 | 813.02 MB/s | 907392 | 3618 | 8.1× |
| Lightning | 1954594 | 807.00 MB/s | 907386 | 3618 | 8.0× |
| Sonic | 2252909 | 700.14 MB/s | 5798588 | 7226 | 7.0× |
| SonicFastest | 2273294 | 693.86 MB/s | 5799511 | 7226 | 6.9× |
| LightningDecodeAny | 3780350 | 199.29 MB/s | 5752203 | 80172 | 4.2× |
| Easyjson | 5575026 | 282.93 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5674912 | 277.95 MB/s | 3624115 | 80269 | 2.8× |
| JSONV2 | 6551316 | 240.77 MB/s | 2704594 | 7318 | 2.4× |
| Stdlib | 15692185 | 100.52 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 224053 | 670.04 MB/s | 81920 | 1 | 8.1× |
| Lightning | 224127 | 669.82 MB/s | 81920 | 1 | 8.1× |
| SonicFastest | 269255 | 557.55 MB/s | 251379 | 6 | 6.8× |
| Sonic | 270242 | 555.52 MB/s | 256276 | 6 | 6.7× |
| LightningDecodeAny | 470358 | 319.16 MB/s | 746005 | 10020 | 3.9× |
| Goccy | 857902 | 174.99 MB/s | 325227 | 10004 | 2.1× |
| JSONV2 | 1091655 | 137.52 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1821859 | 82.40 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33101 | 849.43 MB/s | 30272 | 105 | 9.2× |
| LightningDestructive | 33149 | 848.20 MB/s | 30144 | 103 | 9.1× |
| Sonic | 62905 | 446.98 MB/s | 46652 | 103 | 4.8× |
| SonicFastest | 63010 | 446.23 MB/s | 46725 | 103 | 4.8× |
| Easyjson | 67878 | 414.23 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71739 | 391.94 MB/s | 59207 | 188 | 4.2× |
| JSONV2 | 134009 | 209.81 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 149769 | 187.74 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303237 | 92.72 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1955 | 1191.06 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2012 | 1157.24 MB/s | 32 | 1 | 11.3× |
| Goccy | 4046 | 575.35 MB/s | 3649 | 4 | 5.6× |
| Easyjson | 4200 | 554.30 MB/s | 192 | 2 | 5.4× |
| Sonic | 5026 | 463.20 MB/s | 4166 | 6 | 4.5× |
| SonicFastest | 5069 | 459.30 MB/s | 4269 | 6 | 4.5× |
| JSONV2 | 8448 | 275.57 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10097 | 166.88 MB/s | 9960 | 195 | 2.3× |
| Stdlib | 22731 | 102.42 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220 | 857.43 MB/s | 0 | 0 | 11.3× |
| LightningDestructive | 222 | 849.80 MB/s | 0 | 0 | 11.2× |
| Goccy | 391 | 483.90 MB/s | 304 | 2 | 6.4× |
| Easyjson | 496 | 381.38 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 798 | 236.95 MB/s | 515 | 4 | 3.1× |
| Sonic | 801 | 235.83 MB/s | 515 | 4 | 3.1× |
| JSONV2 | 1040 | 181.77 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1272 | 105.31 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2483 | 76.12 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1542 | 1421.23 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1570 | 1395.86 MB/s | 0 | 0 | 10.2× |
| Goccy | 3188 | 687.28 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3201 | 684.49 MB/s | 24 | 1 | 5.0× |
| Sonic | 6284 | 348.67 MB/s | 3908 | 40 | 2.6× |
| SonicFastest | 6302 | 347.69 MB/s | 3923 | 40 | 2.5× |
| JSONV2 | 8034 | 272.71 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8134 | 222.65 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 16041 | 136.59 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 699545 | 729.73 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 727954 | 701.25 MB/s | 703779 | 1012 | 8.3× |
| Goccy | 1145777 | 445.53 MB/s | 1140324 | 5006 | 5.3× |
| SonicFastest | 1157848 | 440.88 MB/s | 886781 | 2006 | 5.2× |
| Sonic | 1159606 | 440.22 MB/s | 894757 | 2006 | 5.2× |
| Easyjson | 1521427 | 335.52 MB/s | 863778 | 3012 | 4.0× |
| JSONV2 | 3254522 | 156.85 MB/s | 1076012 | 12646 | 1.9× |
| LightningDecodeAny | 3502184 | 131.77 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6029403 | 84.66 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14647.38 MB/s | 0 | 0 | 93.0× |
| LightningDestructive | 1390 | 14237.85 MB/s | 0 | 0 | 90.4× |
| Goccy | 19754 | 1001.76 MB/s | 20491 | 2 | 6.4× |
| SonicFastest | 27821 | 711.30 MB/s | 22156 | 4 | 4.5× |
| Sonic | 27833 | 710.98 MB/s | 22174 | 4 | 4.5× |
| JSONV2 | 29604 | 668.46 MB/s | 8 | 1 | 4.2× |
| LightningDecodeAny | 78195 | 253.06 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 91716 | 215.76 MB/s | 0 | 0 | 1.4× |
| Stdlib | 125672 | 157.47 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2659 | 6816.48 MB/s | 0 | 0 | 38.8× |
| Lightning | 2817 | 6432.90 MB/s | 432 | 2 | 36.6× |
| Easyjson | 3964 | 4572.63 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 10066 | 1800.58 MB/s | 24199 | 6 | 10.3× |
| Sonic | 10110 | 1792.70 MB/s | 24288 | 6 | 10.2× |
| Goccy | 16027 | 1130.87 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 17427 | 1026.12 MB/s | 29088 | 191 | 5.9× |
| JSONV2 | 45257 | 400.47 MB/s | 16500 | 50 | 2.3× |
| Stdlib | 103220 | 175.59 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2193951 | 915.47 MB/s | 2960390 | 7411 | 8.5× |
| Lightning | 2253073 | 891.45 MB/s | 2962103 | 7417 | 8.3× |
| Goccy | 4185102 | 479.92 MB/s | 5412856 | 15837 | 4.4× |
| SonicFastest | 4419769 | 454.43 MB/s | 11019193 | 13683 | 4.2× |
| Sonic | 4430082 | 453.38 MB/s | 11023029 | 13683 | 4.2× |
| Easyjson | 4916278 | 408.54 MB/s | 2981522 | 7439 | 3.8× |
| JSONV2 | 7002732 | 286.82 MB/s | 3173686 | 14563 | 2.7× |
| LightningDecodeAny | 7284131 | 156.82 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18604165 | 107.96 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 890 | 616.96 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 890 | 616.59 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1882 | 291.14 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2135 | 257.11 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2673 | 205.39 MB/s | 1944 | 26 | 2.1× |
| Sonic | 2689 | 204.14 MB/s | 1966 | 26 | 2.1× |
| Goccy | 3015 | 182.10 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3302 | 166.27 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5589 | 98.23 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 499289 | 1264.83 MB/s | 364472 | 566 | 10.8× |
| Lightning | 551381 | 1145.33 MB/s | 413001 | 878 | 9.8× |
| SonicFastest | 1007171 | 627.02 MB/s | 997512 | 1102 | 5.4× |
| Sonic | 1007599 | 626.75 MB/s | 1005906 | 1102 | 5.4× |
| Easyjson | 1156648 | 545.99 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1157157 | 545.75 MB/s | 987141 | 1201 | 4.7× |
| JSONV2 | 2168767 | 291.19 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2382286 | 195.99 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5398030 | 116.99 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 716719 | 784.70 MB/s | 544252 | 448 | 7.3× |
| Lightning | 880738 | 638.56 MB/s | 767621 | 1254 | 6.0× |
| SonicFastest | 1018579 | 552.15 MB/s | 946364 | 1476 | 5.2× |
| Sonic | 1038556 | 541.53 MB/s | 983510 | 1476 | 5.1× |
| Goccy | 1299148 | 432.91 MB/s | 1037429 | 1029 | 4.0× |
| Easyjson | 1719565 | 327.06 MB/s | 775154 | 1254 | 3.1× |
| LightningDecodeAny | 2728601 | 206.12 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 2736168 | 205.55 MB/s | 927442 | 3482 | 1.9× |
| Stdlib | 5249467 | 107.14 MB/s | 1011670 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 644314 | 827.51 MB/s | 333416 | 2084 | 8.5× |
| Lightning | 664117 | 802.84 MB/s | 368224 | 2293 | 8.2× |
| Easyjson | 1104001 | 482.95 MB/s | 428361 | 3273 | 5.0× |
| SonicFastest | 1137223 | 468.84 MB/s | 1061532 | 4351 | 4.8× |
| Sonic | 1141834 | 466.95 MB/s | 1057127 | 4351 | 4.8× |
| Goccy | 1295889 | 411.44 MB/s | 1167243 | 5409 | 4.2× |
| JSONV2 | 2542918 | 209.67 MB/s | 745452 | 13288 | 2.1× |
| LightningDecodeAny | 3322965 | 160.45 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5465373 | 97.56 MB/s | 798692 | 17133 | 1.0× |
