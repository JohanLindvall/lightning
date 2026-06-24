# JSON Deserialization Benchmarks

- generated 2026-06-24T06:17:05Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 102725 | 1238.99 MB/s | 49760 | 3 | 12.1× |
| LightningDestructive | 103997 | 1223.83 MB/s | 49280 | 2 | 12.0× |
| SonicFastest | 193248 | 658.61 MB/s | 214032 | 15 | 6.4× |
| Sonic | 193493 | 657.78 MB/s | 214098 | 15 | 6.4× |
| Goccy | 236387 | 538.42 MB/s | 225053 | 884 | 5.3× |
| Easyjson | 239194 | 532.10 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 390630 | 325.82 MB/s | 195127 | 1805 | 3.2× |
| LightningDecodeAny | 446545 | 211.97 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1245222 | 102.21 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4119550 | 546.43 MB/s | 3122873 | 3081 | 7.6× |
| Lightning | 4250050 | 529.65 MB/s | 3122872 | 3081 | 7.3× |
| SonicFastest | 5259866 | 427.97 MB/s | 4861205 | 2584 | 5.9× |
| Sonic | 5307368 | 424.14 MB/s | 4861229 | 2584 | 5.9× |
| LightningDecodeAny | 11485462 | 195.99 MB/s | 7938299 | 281383 | 2.7× |
| Goccy | 12214542 | 184.29 MB/s | 4148509 | 56533 | 2.5× |
| Easyjson | 13591134 | 165.63 MB/s | 3099808 | 2120 | 2.3× |
| JSONV2 | 17194085 | 130.92 MB/s | 3123198 | 3083 | 1.8× |
| Stdlib | 31108707 | 72.36 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 557830 | 484.74 MB/s | 348024 | 1627 | 7.2× |
| LightningDestructive | 565194 | 478.43 MB/s | 348024 | 1627 | 7.1× |
| SonicFastest | 745453 | 362.74 MB/s | 643485 | 1147 | 5.4× |
| Sonic | 746327 | 362.31 MB/s | 643777 | 1147 | 5.4× |
| LightningDecodeAny | 1588703 | 170.20 MB/s | 1011487 | 37901 | 2.5× |
| Goccy | 1690497 | 159.95 MB/s | 544178 | 8122 | 2.4× |
| Easyjson | 1733650 | 155.97 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2262850 | 119.50 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4027408 | 67.14 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1302179 | 1326.39 MB/s | 959848 | 5881 | 13.0× |
| Lightning | 1340724 | 1288.26 MB/s | 959890 | 5882 | 12.6× |
| SonicFastest | 2148832 | 803.79 MB/s | 2693081 | 5547 | 7.9× |
| Sonic | 2155329 | 801.36 MB/s | 2693115 | 5547 | 7.9× |
| Goccy | 2415474 | 715.06 MB/s | 2581161 | 14603 | 7.0× |
| Easyjson | 3583175 | 482.03 MB/s | 972032 | 5389 | 4.7× |
| LightningDecodeAny | 4045469 | 123.67 MB/s | 4976203 | 81466 | 4.2× |
| JSONV2 | 4166463 | 414.55 MB/s | 1011616 | 7594 | 4.1× |
| Stdlib | 16953347 | 101.88 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1019 | 1777.96 MB/s | 0 | 0 | 14.5× |
| LightningDestructive | 1053 | 1721.45 MB/s | 0 | 0 | 14.0× |
| Easyjson | 2808 | 645.27 MB/s | 24 | 1 | 5.3× |
| Goccy | 3100 | 584.44 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6356 | 285.07 MB/s | 3347 | 38 | 2.3× |
| Sonic | 6523 | 277.80 MB/s | 3346 | 38 | 2.3× |
| JSONV2 | 7477 | 242.35 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8444 | 214.47 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14752 | 122.83 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1085 | 1669.87 MB/s | 0 | 0 | 13.5× |
| LightningDestructive | 1140 | 1589.92 MB/s | 0 | 0 | 12.9× |
| Easyjson | 2803 | 646.56 MB/s | 24 | 1 | 5.2× |
| Goccy | 3053 | 593.48 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6039 | 300.05 MB/s | 3346 | 38 | 2.4× |
| Sonic | 6335 | 286.02 MB/s | 3346 | 38 | 2.3× |
| JSONV2 | 7435 | 243.70 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8484 | 213.46 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14650 | 123.69 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1250 | 1449.95 MB/s | 144 | 10 | 11.7× |
| LightningDestructive | 1295 | 1398.72 MB/s | 144 | 10 | 11.3× |
| Easyjson | 2947 | 614.90 MB/s | 144 | 10 | 5.0× |
| Goccy | 3299 | 549.26 MB/s | 2600 | 5 | 4.4× |
| SonicFastest | 6230 | 290.83 MB/s | 3366 | 40 | 2.3× |
| Sonic | 6442 | 281.30 MB/s | 3365 | 40 | 2.3× |
| JSONV2 | 8048 | 225.16 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8522 | 212.51 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14627 | 123.88 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 736 | 671.55 MB/s | 160 | 1 | 8.2× |
| LightningDestructive | 748 | 660.61 MB/s | 160 | 1 | 8.0× |
| Sonic | 1208 | 409.07 MB/s | 1074 | 8 | 5.0× |
| SonicFastest | 1215 | 406.52 MB/s | 1075 | 8 | 4.9× |
| LightningDecodeAny | 1602 | 307.79 MB/s | 1536 | 30 | 3.8× |
| Easyjson | 2344 | 210.74 MB/s | 448 | 3 | 2.6× |
| Goccy | 2442 | 202.28 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3092 | 159.78 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6008 | 82.23 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 467 | 492.66 MB/s | 160 | 1 | 9.0× |
| LightningDestructive | 469 | 490.75 MB/s | 160 | 1 | 8.9× |
| SonicFastest | 864 | 266.15 MB/s | 801 | 8 | 4.9× |
| Sonic | 881 | 260.94 MB/s | 800 | 8 | 4.8× |
| LightningDecodeAny | 1402 | 163.32 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1518 | 151.55 MB/s | 448 | 3 | 2.8× |
| Goccy | 1668 | 137.85 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2339 | 98.33 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4193 | 54.85 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 84944 | 766.76 MB/s | 172433 | 107 | 7.5× |
| LightningDestructive | 94728 | 687.57 MB/s | 166213 | 102 | 6.8× |
| Sonic | 151415 | 430.16 MB/s | 235786 | 65 | 4.2× |
| SonicFastest | 152001 | 428.50 MB/s | 235765 | 65 | 4.2× |
| Goccy | 173284 | 375.87 MB/s | 228197 | 134 | 3.7× |
| LightningDecodeAny | 181387 | 294.01 MB/s | 176960 | 3252 | 3.5× |
| JSONV2 | 248969 | 261.61 MB/s | 206663 | 607 | 2.6× |
| Stdlib | 639469 | 101.85 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2532550 | 766.21 MB/s | 2846865 | 2238 | 10.0× |
| Lightning | 2621352 | 740.26 MB/s | 2846867 | 2238 | 9.7× |
| Goccy | 4949974 | 392.02 MB/s | 4063063 | 13509 | 5.1× |
| SonicFastest | 6281055 | 308.94 MB/s | 4879748 | 1736 | 4.0× |
| Sonic | 6341596 | 305.99 MB/s | 4880477 | 1736 | 4.0× |
| Easyjson | 7669067 | 253.03 MB/s | 3871265 | 15043 | 3.3× |
| LightningDecodeAny | 9381027 | 206.85 MB/s | 7013525 | 219937 | 2.7× |
| JSONV2 | 11170324 | 173.72 MB/s | 3237193 | 13947 | 2.3× |
| Stdlib | 25307700 | 76.68 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1019281 | 3264.88 MB/s | 351704 | 1286 | 24.9× |
| Lightning | 1591872 | 2090.51 MB/s | 2488906 | 2995 | 15.9× |
| SonicFastest | 2309189 | 1441.13 MB/s | 5895520 | 4263 | 11.0× |
| Sonic | 2331109 | 1427.57 MB/s | 5893558 | 4263 | 10.9× |
| LightningDecodeAny | 3325987 | 924.17 MB/s | 4886621 | 56892 | 7.6× |
| Goccy | 6590796 | 504.92 MB/s | 3948911 | 3816 | 3.9× |
| JSONV2 | 8411548 | 395.63 MB/s | 5364506 | 13243 | 3.0× |
| Stdlib | 25381880 | 131.11 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225599 | 976.72 MB/s | 136896 | 228 | 9.8× |
| LightningDestructive | 235634 | 935.12 MB/s | 136896 | 228 | 9.4× |
| Goccy | 468007 | 470.82 MB/s | 363889 | 1066 | 4.7× |
| SonicFastest | 510282 | 431.81 MB/s | 351161 | 262 | 4.3× |
| Sonic | 511889 | 430.46 MB/s | 351103 | 262 | 4.3× |
| Easyjson | 590932 | 372.88 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 663509 | 332.09 MB/s | 129746 | 470 | 3.3× |
| LightningDecodeAny | 951450 | 113.84 MB/s | 861394 | 11944 | 2.3× |
| Stdlib | 2216875 | 99.39 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14744081 | 549.38 MB/s | 15059833 | 51649 | 6.9× |
| Lightning | 15134052 | 535.22 MB/s | 15059846 | 51649 | 6.7× |
| SonicFastest | 21668874 | 373.81 MB/s | 19850787 | 41640 | 4.7× |
| Sonic | 21728067 | 372.79 MB/s | 19851514 | 41640 | 4.7× |
| Goccy | 24740960 | 327.39 MB/s | 19157384 | 107156 | 4.1× |
| Easyjson | 34380924 | 235.60 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 40508305 | 128.44 MB/s | 34333346 | 912810 | 2.5× |
| JSONV2 | 45537260 | 177.88 MB/s | 15233698 | 78972 | 2.2× |
| Stdlib | 101269362 | 79.99 MB/s | 15665073 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5912162 | 504.63 MB/s | 3985336 | 30015 | 8.6× |
| Lightning | 6133309 | 486.44 MB/s | 3985337 | 30015 | 8.2× |
| Sonic | 9159490 | 325.72 MB/s | 9127516 | 57804 | 5.5× |
| SonicFastest | 9189863 | 324.65 MB/s | 9127785 | 57804 | 5.5× |
| Goccy | 17488465 | 170.60 MB/s | 9842591 | 273618 | 2.9× |
| Easyjson | 17815164 | 167.47 MB/s | 9479441 | 30115 | 2.8× |
| LightningDecodeAny | 19463893 | 94.24 MB/s | 20023836 | 408557 | 2.6× |
| JSONV2 | 24668669 | 120.94 MB/s | 9257084 | 86278 | 2.1× |
| Stdlib | 50596839 | 58.97 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1371156 | 527.73 MB/s | 907394 | 3618 | 9.0× |
| Lightning | 1436434 | 503.75 MB/s | 907387 | 3618 | 8.6× |
| Sonic | 2152302 | 336.20 MB/s | 2367746 | 3683 | 5.7× |
| SonicFastest | 2153601 | 335.99 MB/s | 2367295 | 3683 | 5.7× |
| LightningDecodeAny | 5245728 | 124.02 MB/s | 5752204 | 80172 | 2.3× |
| Easyjson | 5250325 | 137.82 MB/s | 2847906 | 3698 | 2.3× |
| Goccy | 5357001 | 135.08 MB/s | 2675904 | 80266 | 2.3× |
| JSONV2 | 5989835 | 120.80 MB/s | 2704706 | 7318 | 2.1× |
| Stdlib | 12305515 | 58.80 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2078504 | 758.89 MB/s | 907387 | 3618 | 8.5× |
| LightningDestructive | 2083452 | 757.09 MB/s | 907392 | 3618 | 8.5× |
| Sonic | 2434763 | 647.85 MB/s | 3221509 | 3683 | 7.3× |
| SonicFastest | 2435662 | 647.61 MB/s | 3221774 | 3683 | 7.3× |
| LightningDecodeAny | 4425722 | 170.23 MB/s | 5752203 | 80172 | 4.0× |
| Easyjson | 6304925 | 250.18 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 6906413 | 228.39 MB/s | 3450453 | 80259 | 2.6× |
| JSONV2 | 6956122 | 226.76 MB/s | 2704554 | 7318 | 2.5× |
| Stdlib | 17733709 | 88.95 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 237235 | 632.81 MB/s | 81920 | 1 | 9.3× |
| LightningDestructive | 245306 | 611.99 MB/s | 81920 | 1 | 9.0× |
| Sonic | 428780 | 350.12 MB/s | 407137 | 16 | 5.2× |
| SonicFastest | 431436 | 347.96 MB/s | 407549 | 16 | 5.1× |
| LightningDecodeAny | 591766 | 253.68 MB/s | 746005 | 10020 | 3.7× |
| Goccy | 1059850 | 141.65 MB/s | 338455 | 10005 | 2.1× |
| JSONV2 | 1165349 | 128.82 MB/s | 357725 | 20 | 1.9× |
| Stdlib | 2213280 | 67.83 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 31750 | 885.57 MB/s | 30272 | 105 | 9.8× |
| LightningDestructive | 32607 | 862.31 MB/s | 30144 | 103 | 9.5× |
| SonicFastest | 68542 | 410.22 MB/s | 59488 | 83 | 4.5× |
| Sonic | 68613 | 409.79 MB/s | 59494 | 83 | 4.5× |
| Easyjson | 75061 | 374.59 MB/s | 32304 | 138 | 4.1× |
| Goccy | 76702 | 366.58 MB/s | 59269 | 188 | 4.1× |
| JSONV2 | 124838 | 225.23 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 152569 | 184.29 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 310933 | 90.43 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1957 | 1189.74 MB/s | 32 | 1 | 12.1× |
| LightningDestructive | 2001 | 1163.20 MB/s | 32 | 1 | 11.9× |
| Goccy | 4644 | 501.27 MB/s | 3649 | 4 | 5.1× |
| Easyjson | 4788 | 486.24 MB/s | 192 | 2 | 5.0× |
| Sonic | 6039 | 385.48 MB/s | 3708 | 4 | 3.9× |
| SonicFastest | 6072 | 383.40 MB/s | 3711 | 4 | 3.9× |
| JSONV2 | 7758 | 300.08 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 9540 | 176.62 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 23776 | 97.91 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 223 | 848.92 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 226 | 836.99 MB/s | 0 | 0 | 10.8× |
| Goccy | 401 | 471.55 MB/s | 304 | 2 | 6.1× |
| Easyjson | 573 | 329.65 MB/s | 0 | 0 | 4.3× |
| Sonic | 720 | 262.31 MB/s | 341 | 3 | 3.4× |
| SonicFastest | 721 | 262.18 MB/s | 341 | 3 | 3.4× |
| JSONV2 | 925 | 204.25 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1189 | 112.69 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2440 | 77.47 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1473 | 1487.61 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1531 | 1431.21 MB/s | 0 | 0 | 11.4× |
| Goccy | 3447 | 635.58 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3457 | 633.74 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6561 | 333.94 MB/s | 3601 | 38 | 2.7× |
| Sonic | 6747 | 324.75 MB/s | 3599 | 38 | 2.6× |
| JSONV2 | 7542 | 290.49 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8503 | 212.98 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 17404 | 125.89 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 666407 | 766.01 MB/s | 703778 | 1012 | 9.7× |
| Lightning | 717064 | 711.90 MB/s | 703779 | 1012 | 9.0× |
| Goccy | 1274007 | 400.69 MB/s | 1133307 | 5006 | 5.1× |
| Sonic | 1471134 | 346.99 MB/s | 1307045 | 2014 | 4.4× |
| SonicFastest | 1492764 | 341.97 MB/s | 1308153 | 2014 | 4.3× |
| Easyjson | 1563335 | 326.53 MB/s | 863782 | 3012 | 4.1× |
| JSONV2 | 3116511 | 163.80 MB/s | 1075947 | 12645 | 2.1× |
| LightningDecodeAny | 3337824 | 138.25 MB/s | 2785927 | 66022 | 1.9× |
| Stdlib | 6435115 | 79.33 MB/s | 1162119 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 612 | 32307.41 MB/s | 0 | 0 | 232.9× |
| LightningDestructive | 905 | 21876.11 MB/s | 0 | 0 | 157.7× |
| SonicFastest | 6383 | 3100.04 MB/s | 21124 | 3 | 22.3× |
| Goccy | 26086 | 758.61 MB/s | 20492 | 2 | 5.5× |
| Sonic | 29102 | 679.98 MB/s | 20597 | 3 | 4.9× |
| JSONV2 | 35983 | 549.95 MB/s | 8 | 1 | 4.0× |
| LightningDecodeAny | 93673 | 211.25 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 111028 | 178.24 MB/s | 0 | 0 | 1.3× |
| Stdlib | 142650 | 138.72 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2228 | 8133.67 MB/s | 432 | 2 | 53.6× |
| LightningDestructive | 2365 | 7665.04 MB/s | 0 | 0 | 50.5× |
| Easyjson | 4670 | 3880.92 MB/s | 432 | 2 | 25.6× |
| Sonic | 9324 | 1943.76 MB/s | 20480 | 5 | 12.8× |
| SonicFastest | 9821 | 1845.38 MB/s | 20431 | 5 | 12.2× |
| LightningDecodeAny | 18132 | 986.24 MB/s | 29088 | 191 | 6.6× |
| Goccy | 21068 | 860.28 MB/s | 19460 | 2 | 5.7× |
| JSONV2 | 48699 | 372.16 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 119414 | 151.78 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2154881 | 932.07 MB/s | 2960389 | 7411 | 9.7× |
| Lightning | 2270269 | 884.69 MB/s | 2962101 | 7417 | 9.2× |
| Goccy | 4545754 | 441.84 MB/s | 5411968 | 15833 | 4.6× |
| SonicFastest | 5043005 | 398.27 MB/s | 5151559 | 7085 | 4.1× |
| Sonic | 5100628 | 393.77 MB/s | 5151290 | 7085 | 4.1× |
| Easyjson | 5260723 | 381.79 MB/s | 2981484 | 7439 | 4.0× |
| LightningDecodeAny | 6828772 | 167.28 MB/s | 7386073 | 134751 | 3.1× |
| JSONV2 | 7073174 | 283.96 MB/s | 3173678 | 14563 | 2.9× |
| Stdlib | 20847313 | 96.34 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 812 | 676.25 MB/s | 480 | 1 | 7.1× |
| LightningDestructive | 821 | 669.01 MB/s | 480 | 1 | 7.1× |
| LightningDecodeAny | 1869 | 293.23 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 1976 | 277.78 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2100 | 261.48 MB/s | 2261 | 8 | 2.8× |
| Sonic | 2135 | 257.18 MB/s | 2262 | 8 | 2.7× |
| Goccy | 2977 | 184.40 MB/s | 2129 | 43 | 1.9× |
| JSONV2 | 3631 | 151.20 MB/s | 1664 | 7 | 1.6× |
| Stdlib | 5790 | 94.82 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 481992 | 1310.22 MB/s | 364472 | 566 | 12.4× |
| Lightning | 552013 | 1144.02 MB/s | 413001 | 878 | 10.8× |
| SonicFastest | 1080939 | 584.23 MB/s | 1065960 | 814 | 5.5× |
| Sonic | 1095128 | 576.66 MB/s | 1066608 | 814 | 5.5× |
| Easyjson | 1219500 | 517.85 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1403844 | 449.85 MB/s | 992042 | 1201 | 4.3× |
| JSONV2 | 2194818 | 287.73 MB/s | 571587 | 3144 | 2.7× |
| LightningDecodeAny | 2372734 | 196.78 MB/s | 2010198 | 30295 | 2.5× |
| Stdlib | 5970752 | 105.77 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 776114 | 724.65 MB/s | 544249 | 448 | 7.3× |
| Lightning | 955880 | 588.37 MB/s | 767623 | 1254 | 5.9× |
| SonicFastest | 1362415 | 412.80 MB/s | 1348462 | 1184 | 4.2× |
| Sonic | 1363284 | 412.54 MB/s | 1348650 | 1185 | 4.2× |
| Goccy | 1600501 | 351.40 MB/s | 1044797 | 1028 | 3.6× |
| Easyjson | 2010201 | 279.78 MB/s | 775155 | 1254 | 2.8× |
| LightningDecodeAny | 2933023 | 191.75 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 3100670 | 181.38 MB/s | 927405 | 3482 | 1.8× |
| Stdlib | 5684958 | 98.93 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 586773 | 908.66 MB/s | 333416 | 2084 | 10.2× |
| Lightning | 663450 | 803.64 MB/s | 368224 | 2293 | 9.0× |
| Easyjson | 1236875 | 431.07 MB/s | 428362 | 3273 | 4.8× |
| SonicFastest | 1341423 | 397.47 MB/s | 979013 | 3082 | 4.5× |
| Sonic | 1347466 | 395.69 MB/s | 979250 | 3082 | 4.4× |
| Goccy | 1412055 | 377.59 MB/s | 1167078 | 5409 | 4.2× |
| JSONV2 | 2657669 | 200.62 MB/s | 745424 | 13288 | 2.3× |
| LightningDecodeAny | 3278416 | 162.63 MB/s | 2674618 | 50596 | 1.8× |
| Stdlib | 5982935 | 89.12 MB/s | 798692 | 17133 | 1.0× |
