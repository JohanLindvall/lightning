# JSON Deserialization Benchmarks

- generated 2026-06-20T20:17:02Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 106776 | 1191.98 MB/s | 50208 | 4 | 12.4× |
| SonicFastest | 194797 | 653.37 MB/s | 214220 | 15 | 6.8× |
| Sonic | 199159 | 639.06 MB/s | 214815 | 15 | 6.6× |
| Goccy | 241668 | 526.65 MB/s | 225252 | 884 | 5.5× |
| Easyjson | 246165 | 517.03 MB/s | 122864 | 14 | 5.4× |
| JSONV2 | 466737 | 272.69 MB/s | 195127 | 1805 | 2.8× |
| LightningDecodeAny | 471778 | 200.63 MB/s | 466035 | 9715 | 2.8× |
| Stdlib | 1321944 | 96.28 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4172330 | 539.52 MB/s | 3122872 | 3081 | 8.2× |
| SonicFastest | 5491533 | 409.91 MB/s | 4869457 | 2584 | 6.2× |
| Sonic | 5497890 | 409.44 MB/s | 4869205 | 2584 | 6.2× |
| Goccy | 12842775 | 175.28 MB/s | 4194968 | 56535 | 2.7× |
| LightningDecodeAny | 12844492 | 175.25 MB/s | 7938298 | 281383 | 2.7× |
| Easyjson | 13464146 | 167.19 MB/s | 3099809 | 2120 | 2.5× |
| JSONV2 | 17542792 | 128.32 MB/s | 3123189 | 3083 | 1.9× |
| Stdlib | 34058141 | 66.09 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 577767 | 468.01 MB/s | 348025 | 1627 | 7.6× |
| Sonic | 749477 | 360.79 MB/s | 640239 | 1147 | 5.8× |
| SonicFastest | 778874 | 347.17 MB/s | 643293 | 1147 | 5.6× |
| Easyjson | 1759369 | 153.69 MB/s | 330272 | 749 | 2.5× |
| LightningDecodeAny | 1785942 | 151.41 MB/s | 1011489 | 37901 | 2.4× |
| Goccy | 1792656 | 150.84 MB/s | 543020 | 8122 | 2.4× |
| JSONV2 | 2371111 | 114.04 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4366846 | 61.92 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1342391 | 1286.66 MB/s | 959938 | 5883 | 13.3× |
| SonicFastest | 2093786 | 824.92 MB/s | 2694524 | 5547 | 8.5× |
| Sonic | 2101201 | 822.01 MB/s | 2694628 | 5547 | 8.5× |
| Goccy | 2493326 | 692.73 MB/s | 2580809 | 14603 | 7.2× |
| LightningDecodeAny | 4402163 | 113.65 MB/s | 4976252 | 81467 | 4.1× |
| JSONV2 | 4745304 | 363.98 MB/s | 1011616 | 7594 | 3.8× |
| Easyjson | 4947775 | 349.09 MB/s | 972032 | 5389 | 3.6× |
| Stdlib | 17846225 | 96.78 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1090 | 1661.73 MB/s | 0 | 0 | 14.8× |
| Easyjson | 2976 | 608.82 MB/s | 24 | 1 | 5.4× |
| Goccy | 3242 | 558.94 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6256 | 289.66 MB/s | 3347 | 38 | 2.6× |
| Sonic | 6463 | 280.38 MB/s | 3347 | 38 | 2.5× |
| JSONV2 | 8094 | 223.88 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9384 | 192.98 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16110 | 112.48 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1163 | 1557.51 MB/s | 0 | 0 | 13.9× |
| Easyjson | 2999 | 604.11 MB/s | 24 | 1 | 5.4× |
| Goccy | 3208 | 564.80 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6168 | 293.77 MB/s | 3346 | 38 | 2.6× |
| Sonic | 6359 | 284.96 MB/s | 3345 | 38 | 2.5× |
| JSONV2 | 8046 | 225.20 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9473 | 191.18 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16125 | 112.38 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1375 | 1318.28 MB/s | 144 | 10 | 11.8× |
| Goccy | 3207 | 564.95 MB/s | 2600 | 5 | 5.1× |
| Easyjson | 3261 | 555.69 MB/s | 144 | 10 | 5.0× |
| SonicFastest | 6156 | 294.33 MB/s | 3364 | 40 | 2.6× |
| Sonic | 6434 | 281.63 MB/s | 3367 | 40 | 2.5× |
| JSONV2 | 8319 | 217.82 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 9454 | 191.56 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16235 | 111.61 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 740 | 667.91 MB/s | 160 | 1 | 9.4× |
| Sonic | 1260 | 392.09 MB/s | 1076 | 8 | 5.5× |
| SonicFastest | 1270 | 389.09 MB/s | 1076 | 8 | 5.5× |
| LightningDecodeAny | 1742 | 282.99 MB/s | 1536 | 30 | 4.0× |
| Easyjson | 2490 | 198.39 MB/s | 448 | 3 | 2.8× |
| Goccy | 2682 | 184.18 MB/s | 856 | 23 | 2.6× |
| JSONV2 | 3373 | 146.46 MB/s | 528 | 7 | 2.1× |
| Stdlib | 6951 | 71.07 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 485 | 474.66 MB/s | 160 | 1 | 10.1× |
| SonicFastest | 919 | 250.18 MB/s | 801 | 8 | 5.3× |
| Sonic | 920 | 250.01 MB/s | 801 | 8 | 5.3× |
| LightningDecodeAny | 1535 | 149.21 MB/s | 1536 | 30 | 3.2× |
| Easyjson | 1663 | 138.34 MB/s | 448 | 3 | 3.0× |
| Goccy | 1834 | 125.40 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2602 | 88.40 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4916 | 46.79 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 97556 | 667.64 MB/s | 178610 | 112 | 6.9× |
| SonicFastest | 143531 | 453.78 MB/s | 235710 | 65 | 4.7× |
| Sonic | 143770 | 453.03 MB/s | 235718 | 65 | 4.7× |
| Goccy | 177535 | 366.87 MB/s | 227961 | 134 | 3.8× |
| LightningDecodeAny | 205764 | 259.18 MB/s | 183137 | 3257 | 3.3× |
| JSONV2 | 256406 | 254.02 MB/s | 206663 | 607 | 2.6× |
| Stdlib | 675945 | 96.36 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2685564 | 722.56 MB/s | 2846865 | 2238 | 10.5× |
| SonicFastest | 4739511 | 409.42 MB/s | 4881431 | 1736 | 5.9× |
| Sonic | 4810002 | 403.42 MB/s | 4880990 | 1736 | 5.9× |
| Goccy | 5128903 | 378.34 MB/s | 4063892 | 13509 | 5.5× |
| Easyjson | 8262010 | 234.87 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 10125367 | 191.64 MB/s | 7013525 | 219937 | 2.8× |
| JSONV2 | 12662690 | 153.24 MB/s | 3237201 | 13947 | 2.2× |
| Stdlib | 28189265 | 68.84 MB/s | 3551318 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 2123818 | 1566.91 MB/s | 5895414 | 4263 | 11.2× |
| Lightning | 2132945 | 1560.21 MB/s | 4607420 | 4704 | 11.2× |
| SonicFastest | 2135932 | 1558.02 MB/s | 5895330 | 4263 | 11.2× |
| LightningDecodeAny | 4371835 | 703.08 MB/s | 7005137 | 58601 | 5.5× |
| Goccy | 6213241 | 535.60 MB/s | 3948913 | 3816 | 3.8× |
| JSONV2 | 8220171 | 404.84 MB/s | 5364508 | 13243 | 2.9× |
| Stdlib | 23851095 | 139.53 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 236387 | 932.14 MB/s | 136896 | 228 | 10.4× |
| Sonic | 482022 | 457.13 MB/s | 350820 | 262 | 5.1× |
| SonicFastest | 482456 | 456.72 MB/s | 350780 | 262 | 5.1× |
| Goccy | 494118 | 445.94 MB/s | 364010 | 1066 | 5.0× |
| Easyjson | 650868 | 338.54 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 774314 | 284.57 MB/s | 129747 | 470 | 3.2× |
| LightningDecodeAny | 1039267 | 104.22 MB/s | 861395 | 11944 | 2.4× |
| Stdlib | 2453152 | 89.82 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 15616922 | 518.67 MB/s | 15059842 | 51649 | 7.2× |
| Sonic | 18468343 | 438.59 MB/s | 19862331 | 41640 | 6.1× |
| SonicFastest | 18579756 | 435.96 MB/s | 19862413 | 41640 | 6.0× |
| Goccy | 27553233 | 293.98 MB/s | 18999449 | 107156 | 4.1× |
| Easyjson | 36773842 | 220.27 MB/s | 15059619 | 41643 | 3.1× |
| LightningDecodeAny | 46416627 | 112.09 MB/s | 34333349 | 912810 | 2.4× |
| JSONV2 | 48059026 | 168.54 MB/s | 15233767 | 78972 | 2.3× |
| Stdlib | 112382478 | 72.08 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6424303 | 464.40 MB/s | 3985340 | 30015 | 9.0× |
| Sonic | 9642505 | 309.41 MB/s | 9133067 | 57804 | 6.0× |
| SonicFastest | 9791435 | 304.70 MB/s | 9133295 | 57804 | 5.9× |
| Goccy | 18561674 | 160.73 MB/s | 9896043 | 273620 | 3.1× |
| Easyjson | 18846193 | 158.31 MB/s | 9479441 | 30115 | 3.1× |
| LightningDecodeAny | 20764963 | 88.33 MB/s | 20023834 | 408557 | 2.8× |
| JSONV2 | 27780751 | 107.39 MB/s | 9257056 | 86278 | 2.1× |
| Stdlib | 57824446 | 51.60 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1535353 | 471.29 MB/s | 907387 | 3618 | 9.3× |
| Sonic | 2175108 | 332.67 MB/s | 2371737 | 3683 | 6.5× |
| SonicFastest | 2199496 | 328.98 MB/s | 2371011 | 3683 | 6.5× |
| LightningDecodeAny | 5642913 | 115.29 MB/s | 5752207 | 80172 | 2.5× |
| Easyjson | 5710232 | 126.72 MB/s | 2847908 | 3698 | 2.5× |
| Goccy | 5725560 | 126.38 MB/s | 2762844 | 80270 | 2.5× |
| JSONV2 | 6727257 | 107.56 MB/s | 2704701 | 7318 | 2.1× |
| Stdlib | 14211757 | 50.92 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2100591 | 750.91 MB/s | 907389 | 3618 | 9.6× |
| SonicFastest | 2503579 | 630.04 MB/s | 3224725 | 3683 | 8.0× |
| Sonic | 2551314 | 618.25 MB/s | 3229847 | 3683 | 7.9× |
| LightningDecodeAny | 4817829 | 156.38 MB/s | 5752200 | 80172 | 4.2× |
| Easyjson | 6932223 | 227.54 MB/s | 2847907 | 3698 | 2.9× |
| JSONV2 | 6978833 | 226.02 MB/s | 2704553 | 7318 | 2.9× |
| Goccy | 7097649 | 222.24 MB/s | 3498609 | 80263 | 2.8× |
| Stdlib | 20111926 | 78.43 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 256215 | 585.93 MB/s | 81920 | 1 | 8.5× |
| Sonic | 421333 | 356.31 MB/s | 408397 | 16 | 5.1× |
| SonicFastest | 434479 | 345.53 MB/s | 408506 | 16 | 5.0× |
| LightningDecodeAny | 612850 | 244.96 MB/s | 746006 | 10020 | 3.5× |
| Goccy | 1032649 | 145.38 MB/s | 324727 | 10005 | 2.1× |
| JSONV2 | 1105550 | 135.79 MB/s | 357723 | 20 | 2.0× |
| Stdlib | 2165181 | 69.34 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 37757 | 744.68 MB/s | 30384 | 107 | 9.2× |
| SonicFastest | 57457 | 489.35 MB/s | 59525 | 83 | 6.0× |
| Sonic | 57929 | 485.37 MB/s | 59515 | 83 | 6.0× |
| Goccy | 80431 | 349.58 MB/s | 59275 | 188 | 4.3× |
| Easyjson | 82139 | 342.31 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 142632 | 197.13 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 172078 | 163.40 MB/s | 135136 | 2680 | 2.0× |
| Stdlib | 345820 | 81.31 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1977 | 1177.68 MB/s | 32 | 1 | 13.3× |
| SonicFastest | 4763 | 488.75 MB/s | 3712 | 4 | 5.5× |
| Sonic | 4794 | 485.59 MB/s | 3714 | 4 | 5.5× |
| Goccy | 4962 | 469.12 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 5029 | 462.88 MB/s | 192 | 2 | 5.2× |
| JSONV2 | 8506 | 273.71 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10840 | 155.44 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 26344 | 88.37 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222 | 849.86 MB/s | 0 | 0 | 12.5× |
| Goccy | 415 | 454.99 MB/s | 304 | 2 | 6.7× |
| Easyjson | 580 | 325.97 MB/s | 0 | 0 | 4.8× |
| SonicFastest | 647 | 291.96 MB/s | 341 | 3 | 4.3× |
| Sonic | 650 | 290.89 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1035 | 182.53 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1403 | 95.53 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2788 | 67.78 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1456 | 1504.61 MB/s | 0 | 0 | 13.0× |
| Goccy | 3754 | 583.60 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3777 | 580.14 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6733 | 325.42 MB/s | 3603 | 38 | 2.8× |
| Sonic | 6928 | 316.25 MB/s | 3600 | 38 | 2.7× |
| JSONV2 | 8340 | 262.72 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 9561 | 189.41 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18946 | 115.65 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 744632 | 685.54 MB/s | 703778 | 1012 | 9.6× |
| SonicFastest | 1296849 | 393.63 MB/s | 1308650 | 2014 | 5.5× |
| Goccy | 1309689 | 389.77 MB/s | 1139371 | 5006 | 5.5× |
| Sonic | 1311905 | 389.11 MB/s | 1309721 | 2014 | 5.5× |
| Easyjson | 1704746 | 299.44 MB/s | 863779 | 3012 | 4.2× |
| JSONV2 | 3465810 | 147.29 MB/s | 1075968 | 12645 | 2.1× |
| LightningDecodeAny | 3755116 | 122.89 MB/s | 2785927 | 66022 | 1.9× |
| Stdlib | 7150793 | 71.39 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 547 | 36176.60 MB/s | 0 | 0 | 296.0× |
| SonicFastest | 7150 | 2767.76 MB/s | 21131 | 3 | 22.6× |
| Goccy | 23830 | 830.41 MB/s | 20492 | 2 | 6.8× |
| Sonic | 32173 | 615.08 MB/s | 20623 | 3 | 5.0× |
| JSONV2 | 34749 | 569.48 MB/s | 8 | 1 | 4.7× |
| LightningDecodeAny | 100906 | 196.10 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 103519 | 191.16 MB/s | 0 | 0 | 1.6× |
| Stdlib | 161937 | 122.20 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2445 | 7411.43 MB/s | 864 | 4 | 49.8× |
| Easyjson | 5007 | 3619.45 MB/s | 432 | 2 | 24.3× |
| Sonic | 8508 | 2130.33 MB/s | 20482 | 5 | 14.3× |
| SonicFastest | 9048 | 2003.03 MB/s | 20440 | 5 | 13.5× |
| LightningDecodeAny | 19818 | 902.33 MB/s | 29520 | 193 | 6.1× |
| Goccy | 20614 | 879.21 MB/s | 19460 | 2 | 5.9× |
| JSONV2 | 50168 | 361.26 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 121837 | 148.76 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2596117 | 773.65 MB/s | 2963781 | 7423 | 8.6× |
| SonicFastest | 4324516 | 464.44 MB/s | 5158970 | 7085 | 5.1× |
| Sonic | 4376580 | 458.92 MB/s | 5158929 | 7085 | 5.1× |
| Goccy | 4990475 | 402.47 MB/s | 5410887 | 15831 | 4.5× |
| Easyjson | 5872848 | 342.00 MB/s | 2981483 | 7439 | 3.8× |
| LightningDecodeAny | 7365427 | 155.09 MB/s | 7387754 | 134757 | 3.0× |
| JSONV2 | 8110498 | 247.64 MB/s | 3173677 | 14563 | 2.7× |
| Stdlib | 22210952 | 90.43 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 882 | 622.18 MB/s | 480 | 1 | 7.8× |
| LightningDecodeAny | 2148 | 255.07 MB/s | 2261 | 50 | 3.2× |
| Easyjson | 2265 | 242.37 MB/s | 1616 | 5 | 3.1× |
| SonicFastest | 2423 | 226.57 MB/s | 2262 | 8 | 2.9× |
| Sonic | 2444 | 224.67 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3470 | 158.21 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3623 | 151.51 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6919 | 79.35 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 626091 | 1008.66 MB/s | 460873 | 1190 | 10.4× |
| Sonic | 979059 | 645.02 MB/s | 1065646 | 814 | 6.7× |
| SonicFastest | 990180 | 637.78 MB/s | 1066690 | 814 | 6.6× |
| Easyjson | 1331018 | 474.46 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1445693 | 436.82 MB/s | 988965 | 1200 | 4.5× |
| JSONV2 | 2397644 | 263.39 MB/s | 571590 | 3144 | 2.7× |
| LightningDecodeAny | 2622880 | 178.01 MB/s | 2058069 | 30607 | 2.5× |
| Stdlib | 6521283 | 96.84 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1135629 | 495.24 MB/s | 886432 | 2062 | 5.4× |
| Sonic | 1290131 | 435.93 MB/s | 1348154 | 1185 | 4.8× |
| SonicFastest | 1302518 | 431.79 MB/s | 1348159 | 1185 | 4.7× |
| Goccy | 1648090 | 341.25 MB/s | 1041344 | 1028 | 3.7× |
| Easyjson | 2171262 | 259.02 MB/s | 775154 | 1254 | 2.8× |
| JSONV2 | 3247749 | 173.17 MB/s | 927405 | 3482 | 1.9× |
| LightningDecodeAny | 3372071 | 166.78 MB/s | 2232991 | 31103 | 1.8× |
| Stdlib | 6173291 | 91.10 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 737662 | 722.79 MB/s | 402760 | 2502 | 8.7× |
| SonicFastest | 1113108 | 479.00 MB/s | 982045 | 3082 | 5.7× |
| Sonic | 1123089 | 474.74 MB/s | 982323 | 3082 | 5.7× |
| Easyjson | 1312141 | 406.34 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1428429 | 373.26 MB/s | 1167077 | 5408 | 4.5× |
| JSONV2 | 2749318 | 193.93 MB/s | 745421 | 13288 | 2.3× |
| LightningDecodeAny | 3683876 | 144.73 MB/s | 2709154 | 50805 | 1.7× |
| Stdlib | 6389945 | 83.44 MB/s | 798692 | 17133 | 1.0× |
