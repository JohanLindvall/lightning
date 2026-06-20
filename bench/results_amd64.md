# JSON Deserialization Benchmarks

- generated 2026-06-20T18:49:56Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104967 | 1212.53 MB/s | 50208 | 4 | 11.8× |
| Easyjson | 220857 | 576.28 MB/s | 122864 | 14 | 5.6× |
| Goccy | 245925 | 517.54 MB/s | 225326 | 884 | 5.0× |
| Sonic | 258293 | 492.75 MB/s | 214118 | 15 | 4.8× |
| SonicFastest | 258375 | 492.60 MB/s | 214297 | 15 | 4.8× |
| JSONV2 | 418302 | 304.27 MB/s | 195128 | 1805 | 3.0× |
| LightningDecodeAny | 445538 | 212.45 MB/s | 466034 | 9715 | 2.8× |
| Stdlib | 1234541 | 103.10 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4145204 | 543.05 MB/s | 3122874 | 3081 | 7.5× |
| Sonic | 5477687 | 410.95 MB/s | 4864060 | 2584 | 5.7× |
| SonicFastest | 5768073 | 390.26 MB/s | 4862413 | 2584 | 5.4× |
| LightningDecodeAny | 11549806 | 194.90 MB/s | 7938299 | 281383 | 2.7× |
| Goccy | 12528083 | 179.68 MB/s | 4215479 | 56536 | 2.5× |
| Easyjson | 13551440 | 166.11 MB/s | 3099810 | 2120 | 2.3× |
| JSONV2 | 17218999 | 130.73 MB/s | 3123189 | 3083 | 1.8× |
| Stdlib | 31251866 | 72.03 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 556017 | 486.32 MB/s | 348024 | 1627 | 7.3× |
| Sonic | 737378 | 366.71 MB/s | 641447 | 1147 | 5.5× |
| SonicFastest | 739642 | 365.59 MB/s | 641485 | 1147 | 5.5× |
| LightningDecodeAny | 1591437 | 169.91 MB/s | 1011488 | 37901 | 2.5× |
| Goccy | 1692915 | 159.73 MB/s | 542388 | 8122 | 2.4× |
| Easyjson | 1723758 | 156.87 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2269641 | 119.14 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4036239 | 66.99 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1319657 | 1308.83 MB/s | 959938 | 5883 | 12.6× |
| SonicFastest | 2025792 | 852.61 MB/s | 2692955 | 5547 | 8.2× |
| Sonic | 2028901 | 851.30 MB/s | 2694102 | 5547 | 8.2× |
| Goccy | 2410365 | 716.57 MB/s | 2581027 | 14603 | 6.9× |
| Easyjson | 3598458 | 479.98 MB/s | 972032 | 5389 | 4.6× |
| LightningDecodeAny | 4081179 | 122.59 MB/s | 4976251 | 81467 | 4.1× |
| JSONV2 | 4247003 | 406.69 MB/s | 1011614 | 7594 | 3.9× |
| Stdlib | 16651860 | 103.72 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1047 | 1730.42 MB/s | 0 | 0 | 14.0× |
| Easyjson | 2734 | 662.67 MB/s | 24 | 1 | 5.4× |
| Goccy | 3287 | 551.25 MB/s | 2608 | 4 | 4.5× |
| SonicFastest | 6193 | 292.61 MB/s | 3346 | 38 | 2.4× |
| Sonic | 6343 | 285.67 MB/s | 3343 | 38 | 2.3× |
| JSONV2 | 7793 | 232.51 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 8489 | 213.33 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14673 | 123.49 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1104 | 1641.22 MB/s | 0 | 0 | 13.1× |
| Easyjson | 2726 | 664.75 MB/s | 24 | 1 | 5.3× |
| Goccy | 3332 | 543.86 MB/s | 2608 | 4 | 4.4× |
| SonicFastest | 6084 | 297.82 MB/s | 3347 | 38 | 2.4× |
| Sonic | 6186 | 292.90 MB/s | 3346 | 38 | 2.3× |
| JSONV2 | 7998 | 226.56 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8535 | 212.18 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14512 | 124.86 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1293 | 1401.09 MB/s | 144 | 10 | 11.4× |
| Easyjson | 2975 | 609.06 MB/s | 144 | 10 | 4.9× |
| Goccy | 3301 | 548.86 MB/s | 2600 | 5 | 4.5× |
| SonicFastest | 6153 | 294.48 MB/s | 3369 | 40 | 2.4× |
| Sonic | 6315 | 286.94 MB/s | 3368 | 40 | 2.3× |
| JSONV2 | 8008 | 226.28 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8675 | 208.77 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14691 | 123.34 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 734 | 673.43 MB/s | 160 | 1 | 8.2× |
| Sonic | 1217 | 405.90 MB/s | 1076 | 8 | 5.0× |
| SonicFastest | 1228 | 402.20 MB/s | 1076 | 8 | 4.9× |
| LightningDecodeAny | 1619 | 304.44 MB/s | 1536 | 30 | 3.7× |
| Easyjson | 2402 | 205.62 MB/s | 448 | 3 | 2.5× |
| Goccy | 2581 | 191.43 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 2970 | 166.34 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6032 | 81.89 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 466 | 493.10 MB/s | 160 | 1 | 9.0× |
| Sonic | 848 | 271.15 MB/s | 801 | 8 | 5.0× |
| SonicFastest | 851 | 270.15 MB/s | 801 | 8 | 4.9× |
| LightningDecodeAny | 1382 | 165.73 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1549 | 148.52 MB/s | 448 | 3 | 2.7× |
| Goccy | 1714 | 134.17 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2298 | 100.10 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4208 | 54.66 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 91368 | 712.86 MB/s | 178610 | 112 | 7.0× |
| Sonic | 153109 | 425.40 MB/s | 235966 | 65 | 4.2× |
| SonicFastest | 154359 | 421.95 MB/s | 235932 | 65 | 4.1× |
| Goccy | 173353 | 375.72 MB/s | 228547 | 134 | 3.7× |
| LightningDecodeAny | 186487 | 285.97 MB/s | 183137 | 3257 | 3.4× |
| JSONV2 | 238545 | 273.04 MB/s | 206664 | 607 | 2.7× |
| Stdlib | 638930 | 101.94 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2635660 | 736.24 MB/s | 2846867 | 2238 | 9.6× |
| SonicFastest | 4648717 | 417.42 MB/s | 4873949 | 1736 | 5.4× |
| Sonic | 4836829 | 401.19 MB/s | 4880143 | 1736 | 5.2× |
| Goccy | 5019440 | 386.59 MB/s | 4063278 | 13509 | 5.0× |
| Easyjson | 7638863 | 254.03 MB/s | 3871265 | 15043 | 3.3× |
| LightningDecodeAny | 9478155 | 204.73 MB/s | 7013525 | 219937 | 2.7× |
| JSONV2 | 10945605 | 177.28 MB/s | 3237182 | 13947 | 2.3× |
| Stdlib | 25236619 | 76.89 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1988507 | 1673.53 MB/s | 4607418 | 4704 | 12.8× |
| SonicFastest | 2335803 | 1424.71 MB/s | 5896313 | 4263 | 10.9× |
| Sonic | 2349617 | 1416.33 MB/s | 5895804 | 4263 | 10.8× |
| LightningDecodeAny | 3820801 | 804.48 MB/s | 7005136 | 58601 | 6.7× |
| Goccy | 6595845 | 504.53 MB/s | 3948913 | 3816 | 3.9× |
| JSONV2 | 8117003 | 409.98 MB/s | 5364505 | 13243 | 3.1× |
| Stdlib | 25414417 | 130.94 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 248082 | 888.20 MB/s | 123584 | 225 | 9.0× |
| Goccy | 461833 | 477.11 MB/s | 363277 | 1066 | 4.8× |
| SonicFastest | 514883 | 427.95 MB/s | 351388 | 262 | 4.3× |
| Sonic | 516460 | 426.65 MB/s | 351430 | 262 | 4.3× |
| Easyjson | 583593 | 377.57 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 665248 | 331.22 MB/s | 129747 | 470 | 3.3× |
| LightningDecodeAny | 953470 | 113.60 MB/s | 861394 | 11944 | 2.3× |
| Stdlib | 2221027 | 99.21 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 14893707 | 543.86 MB/s | 15059836 | 51649 | 6.9× |
| Sonic | 19939923 | 406.22 MB/s | 19853310 | 41640 | 5.1× |
| SonicFastest | 19975272 | 405.50 MB/s | 19855158 | 41640 | 5.1× |
| Goccy | 25044697 | 323.42 MB/s | 19052077 | 107155 | 4.1× |
| Easyjson | 34477728 | 234.94 MB/s | 15059617 | 41643 | 3.0× |
| LightningDecodeAny | 41926129 | 124.10 MB/s | 34333348 | 912810 | 2.4× |
| JSONV2 | 46085990 | 175.76 MB/s | 15233719 | 78972 | 2.2× |
| Stdlib | 102550218 | 78.99 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6026407 | 495.07 MB/s | 3985339 | 30015 | 8.4× |
| Sonic | 9344543 | 319.27 MB/s | 9128480 | 57804 | 5.4× |
| SonicFastest | 9428804 | 316.42 MB/s | 9127522 | 57804 | 5.4× |
| Goccy | 17623437 | 169.29 MB/s | 9774287 | 273616 | 2.9× |
| Easyjson | 18010995 | 165.65 MB/s | 9479441 | 30115 | 2.8× |
| LightningDecodeAny | 19162390 | 95.72 MB/s | 20023838 | 408557 | 2.6× |
| JSONV2 | 24753879 | 120.53 MB/s | 9257042 | 86278 | 2.0× |
| Stdlib | 50533646 | 59.04 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1441341 | 502.03 MB/s | 907387 | 3618 | 8.6× |
| SonicFastest | 2255132 | 320.87 MB/s | 2367732 | 3683 | 5.5× |
| Sonic | 2256780 | 320.63 MB/s | 2367843 | 3683 | 5.5× |
| Goccy | 5250134 | 137.82 MB/s | 2635671 | 80263 | 2.3× |
| Easyjson | 5281851 | 137.00 MB/s | 2847907 | 3698 | 2.3× |
| LightningDecodeAny | 5304310 | 122.65 MB/s | 5752206 | 80172 | 2.3× |
| JSONV2 | 6105063 | 118.52 MB/s | 2704703 | 7318 | 2.0× |
| Stdlib | 12336667 | 58.65 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2078463 | 758.90 MB/s | 907387 | 3618 | 8.5× |
| Sonic | 2434286 | 647.97 MB/s | 3222665 | 3683 | 7.3× |
| SonicFastest | 2672099 | 590.30 MB/s | 3227391 | 3683 | 6.6× |
| LightningDecodeAny | 4422959 | 170.34 MB/s | 5752204 | 80172 | 4.0× |
| Easyjson | 6484945 | 243.23 MB/s | 2847905 | 3698 | 2.7× |
| Goccy | 6621767 | 238.21 MB/s | 3451524 | 80259 | 2.7× |
| JSONV2 | 6734617 | 234.22 MB/s | 2704554 | 7318 | 2.6× |
| Stdlib | 17737711 | 88.93 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 233795 | 642.12 MB/s | 81920 | 1 | 9.5× |
| Sonic | 382711 | 392.26 MB/s | 407202 | 16 | 5.8× |
| SonicFastest | 391409 | 383.55 MB/s | 407287 | 16 | 5.7× |
| LightningDecodeAny | 594625 | 252.46 MB/s | 746005 | 10020 | 3.7× |
| Goccy | 1027560 | 146.10 MB/s | 325595 | 10005 | 2.2× |
| JSONV2 | 1124082 | 133.55 MB/s | 357729 | 20 | 2.0× |
| Stdlib | 2215551 | 67.76 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 35790 | 785.61 MB/s | 29104 | 107 | 8.7× |
| SonicFastest | 69175 | 406.46 MB/s | 59451 | 83 | 4.5× |
| Sonic | 69292 | 405.77 MB/s | 59466 | 83 | 4.5× |
| Easyjson | 73866 | 380.65 MB/s | 32304 | 138 | 4.2× |
| Goccy | 78958 | 356.10 MB/s | 59265 | 188 | 3.9× |
| JSONV2 | 126896 | 221.58 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 153681 | 182.96 MB/s | 135136 | 2680 | 2.0× |
| Stdlib | 309659 | 90.80 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2060 | 1129.89 MB/s | 32 | 1 | 11.4× |
| Goccy | 4710 | 494.26 MB/s | 3649 | 4 | 5.0× |
| Easyjson | 4820 | 482.95 MB/s | 192 | 2 | 4.9× |
| SonicFastest | 6139 | 379.19 MB/s | 3708 | 4 | 3.8× |
| Sonic | 6145 | 378.87 MB/s | 3711 | 4 | 3.8× |
| JSONV2 | 7877 | 295.55 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 9722 | 173.33 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 23521 | 98.97 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224 | 844.08 MB/s | 0 | 0 | 11.0× |
| Goccy | 395 | 478.60 MB/s | 304 | 2 | 6.2× |
| Easyjson | 557 | 339.27 MB/s | 0 | 0 | 4.4× |
| Sonic | 814 | 232.32 MB/s | 341 | 3 | 3.0× |
| SonicFastest | 814 | 232.11 MB/s | 341 | 3 | 3.0× |
| JSONV2 | 938 | 201.41 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1214 | 110.35 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2457 | 76.93 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1449 | 1511.90 MB/s | 0 | 0 | 11.9× |
| Easyjson | 3368 | 650.57 MB/s | 24 | 1 | 5.1× |
| Goccy | 3751 | 584.16 MB/s | 2864 | 4 | 4.6× |
| SonicFastest | 6458 | 339.24 MB/s | 3601 | 38 | 2.7× |
| Sonic | 6676 | 328.19 MB/s | 3599 | 38 | 2.6× |
| JSONV2 | 8085 | 271.00 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 8631 | 209.82 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 17239 | 127.10 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 880012 | 580.08 MB/s | 703779 | 1012 | 7.3× |
| Goccy | 1268642 | 402.38 MB/s | 1135716 | 5006 | 5.1× |
| Sonic | 1561657 | 326.88 MB/s | 1308031 | 2014 | 4.1× |
| SonicFastest | 1562315 | 326.74 MB/s | 1307814 | 2014 | 4.1× |
| Easyjson | 1566993 | 325.77 MB/s | 863782 | 3012 | 4.1× |
| JSONV2 | 3202742 | 159.39 MB/s | 1075946 | 12645 | 2.0× |
| LightningDecodeAny | 3422298 | 134.84 MB/s | 2785927 | 66022 | 1.9× |
| Stdlib | 6438521 | 79.28 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 609 | 32478.59 MB/s | 0 | 0 | 233.9× |
| SonicFastest | 7574 | 2612.75 MB/s | 21048 | 3 | 18.8× |
| Goccy | 26452 | 748.10 MB/s | 20492 | 2 | 5.4× |
| Sonic | 29199 | 677.74 MB/s | 20608 | 3 | 4.9× |
| JSONV2 | 36122 | 547.84 MB/s | 8 | 1 | 3.9× |
| LightningDecodeAny | 93525 | 211.58 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 118594 | 166.86 MB/s | 0 | 0 | 1.2× |
| Stdlib | 142489 | 138.88 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2401 | 7548.26 MB/s | 864 | 4 | 50.2× |
| Easyjson | 4682 | 3870.68 MB/s | 432 | 2 | 25.7× |
| SonicFastest | 10368 | 1748.14 MB/s | 20400 | 5 | 11.6× |
| Sonic | 10425 | 1738.53 MB/s | 20405 | 5 | 11.6× |
| LightningDecodeAny | 18445 | 969.47 MB/s | 29520 | 193 | 6.5× |
| Goccy | 21643 | 837.41 MB/s | 19460 | 2 | 5.6× |
| JSONV2 | 47786 | 379.28 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 120417 | 150.51 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2351040 | 854.30 MB/s | 2963782 | 7423 | 8.8× |
| Goccy | 4556423 | 440.80 MB/s | 5409720 | 15830 | 4.5× |
| SonicFastest | 4853287 | 413.84 MB/s | 5151472 | 7085 | 4.2× |
| Sonic | 4904733 | 409.50 MB/s | 5150416 | 7085 | 4.2× |
| Easyjson | 5347303 | 375.61 MB/s | 2981484 | 7439 | 3.9× |
| LightningDecodeAny | 6754764 | 169.11 MB/s | 7387753 | 134757 | 3.1× |
| JSONV2 | 7088926 | 283.33 MB/s | 3173678 | 14563 | 2.9× |
| Stdlib | 20621628 | 97.40 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 835 | 657.34 MB/s | 480 | 1 | 6.9× |
| LightningDecodeAny | 1910 | 286.94 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 1976 | 277.89 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2158 | 254.44 MB/s | 2262 | 8 | 2.7× |
| Sonic | 2178 | 252.04 MB/s | 2262 | 8 | 2.7× |
| Goccy | 2991 | 183.55 MB/s | 2129 | 43 | 1.9× |
| JSONV2 | 3089 | 177.73 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 5804 | 94.58 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 576749 | 1094.95 MB/s | 460873 | 1190 | 10.3× |
| SonicFastest | 1054904 | 598.65 MB/s | 1064882 | 814 | 5.7× |
| Sonic | 1058632 | 596.54 MB/s | 1064882 | 814 | 5.6× |
| Easyjson | 1252618 | 504.16 MB/s | 422504 | 936 | 4.8× |
| Goccy | 1367322 | 461.86 MB/s | 989774 | 1201 | 4.4× |
| JSONV2 | 2138199 | 295.35 MB/s | 571589 | 3144 | 2.8× |
| LightningDecodeAny | 2437130 | 191.58 MB/s | 2058070 | 30607 | 2.4× |
| Stdlib | 5966758 | 105.84 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1051034 | 535.10 MB/s | 886432 | 2062 | 5.4× |
| Sonic | 1405732 | 400.08 MB/s | 1346639 | 1184 | 4.0× |
| SonicFastest | 1407889 | 399.47 MB/s | 1346980 | 1184 | 4.0× |
| Goccy | 1582542 | 355.38 MB/s | 1044004 | 1028 | 3.6× |
| Easyjson | 1980135 | 284.03 MB/s | 775155 | 1254 | 2.9× |
| JSONV2 | 3013711 | 186.62 MB/s | 927407 | 3482 | 1.9× |
| LightningDecodeAny | 3022094 | 186.10 MB/s | 2232991 | 31103 | 1.9× |
| Stdlib | 5645197 | 99.63 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 769644 | 692.76 MB/s | 413026 | 3155 | 7.7× |
| Easyjson | 1229134 | 433.78 MB/s | 428362 | 3273 | 4.8× |
| SonicFastest | 1395371 | 382.10 MB/s | 979656 | 3082 | 4.2× |
| Sonic | 1397579 | 381.50 MB/s | 979213 | 3082 | 4.2× |
| Goccy | 1492047 | 357.35 MB/s | 1167054 | 5408 | 4.0× |
| JSONV2 | 2636941 | 202.20 MB/s | 745421 | 13288 | 2.2× |
| LightningDecodeAny | 3374129 | 158.02 MB/s | 2709153 | 50805 | 1.8× |
| Stdlib | 5928530 | 89.93 MB/s | 798692 | 17133 | 1.0× |
