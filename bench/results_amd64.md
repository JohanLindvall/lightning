# JSON Deserialization Benchmarks

- generated 2026-06-21T14:51:27Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 106090 | 1199.69 MB/s | 49760 | 3 | 12.5× |
| LightningDestructive | 107044 | 1189.00 MB/s | 49281 | 2 | 12.4× |
| SonicFastest | 200213 | 635.70 MB/s | 214317 | 15 | 6.6× |
| Sonic | 202506 | 628.50 MB/s | 214645 | 15 | 6.5× |
| Easyjson | 242508 | 524.83 MB/s | 122864 | 14 | 5.5× |
| Goccy | 248810 | 511.53 MB/s | 225276 | 884 | 5.3× |
| LightningDecodeAny | 465412 | 203.37 MB/s | 465586 | 9714 | 2.8× |
| JSONV2 | 479497 | 265.43 MB/s | 195129 | 1805 | 2.8× |
| Stdlib | 1324549 | 96.09 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4189092 | 537.36 MB/s | 3122874 | 3081 | 8.1× |
| LightningDestructive | 4219630 | 533.47 MB/s | 3123685 | 3081 | 8.1× |
| SonicFastest | 5139421 | 438.00 MB/s | 4871595 | 2584 | 6.6× |
| Sonic | 5180422 | 434.53 MB/s | 4873545 | 2584 | 6.6× |
| LightningDecodeAny | 12568835 | 179.10 MB/s | 7938300 | 281383 | 2.7× |
| Goccy | 12945567 | 173.89 MB/s | 4177737 | 56534 | 2.6× |
| Easyjson | 13293415 | 169.34 MB/s | 3099808 | 2120 | 2.6× |
| JSONV2 | 17374396 | 129.56 MB/s | 3123198 | 3083 | 2.0× |
| Stdlib | 34034628 | 66.14 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 571471 | 473.17 MB/s | 348024 | 1627 | 7.7× |
| LightningDestructive | 595758 | 453.88 MB/s | 348039 | 1627 | 7.4× |
| SonicFastest | 745488 | 362.72 MB/s | 640211 | 1147 | 5.9× |
| Sonic | 746669 | 362.15 MB/s | 640550 | 1147 | 5.9× |
| LightningDecodeAny | 1757413 | 153.86 MB/s | 1011488 | 37901 | 2.5× |
| Easyjson | 1768152 | 152.93 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1806026 | 149.72 MB/s | 541809 | 8122 | 2.4× |
| JSONV2 | 2349855 | 115.07 MB/s | 348160 | 1628 | 1.9× |
| Stdlib | 4394704 | 61.53 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1355611 | 1274.11 MB/s | 959889 | 5882 | 13.0× |
| LightningDestructive | 1383303 | 1248.61 MB/s | 960042 | 5881 | 12.7× |
| Sonic | 2160347 | 799.50 MB/s | 2693927 | 5547 | 8.1× |
| SonicFastest | 2173952 | 794.50 MB/s | 2694875 | 5547 | 8.1× |
| Goccy | 2961682 | 583.18 MB/s | 2581180 | 14603 | 5.9× |
| Easyjson | 4271195 | 404.38 MB/s | 972032 | 5389 | 4.1× |
| LightningDecodeAny | 4368897 | 114.51 MB/s | 4976203 | 81466 | 4.0× |
| JSONV2 | 4583294 | 376.85 MB/s | 1011613 | 7594 | 3.8× |
| Stdlib | 17594898 | 98.17 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1023 | 1771.04 MB/s | 0 | 0 | 15.6× |
| LightningDestructive | 1042 | 1738.84 MB/s | 0 | 0 | 15.4× |
| Easyjson | 2845 | 636.87 MB/s | 24 | 1 | 5.6× |
| Goccy | 3503 | 517.28 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6195 | 292.47 MB/s | 3344 | 38 | 2.6× |
| Sonic | 6713 | 269.94 MB/s | 3347 | 38 | 2.4× |
| JSONV2 | 8476 | 213.78 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9928 | 182.41 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 15999 | 113.26 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1112 | 1630.08 MB/s | 0 | 0 | 14.6× |
| LightningDestructive | 1120 | 1618.41 MB/s | 0 | 0 | 14.5× |
| Easyjson | 2869 | 631.59 MB/s | 24 | 1 | 5.7× |
| Goccy | 3513 | 515.85 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6293 | 287.95 MB/s | 3343 | 38 | 2.6× |
| Sonic | 6450 | 280.91 MB/s | 3346 | 38 | 2.5× |
| JSONV2 | 8433 | 214.86 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9352 | 193.65 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16227 | 111.66 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1335 | 1357.01 MB/s | 144 | 10 | 12.0× |
| LightningDestructive | 1369 | 1323.53 MB/s | 144 | 10 | 11.7× |
| Easyjson | 3254 | 556.78 MB/s | 144 | 10 | 4.9× |
| Goccy | 3282 | 552.10 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6337 | 285.93 MB/s | 3367 | 40 | 2.5× |
| Sonic | 6495 | 278.99 MB/s | 3369 | 40 | 2.5× |
| JSONV2 | 8394 | 215.86 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9471 | 191.22 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16017 | 113.13 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 738 | 669.34 MB/s | 160 | 1 | 9.4× |
| LightningDestructive | 747 | 660.96 MB/s | 160 | 1 | 9.3× |
| Sonic | 1264 | 390.93 MB/s | 1075 | 8 | 5.5× |
| SonicFastest | 1273 | 387.91 MB/s | 1075 | 8 | 5.5× |
| LightningDecodeAny | 1748 | 282.08 MB/s | 1536 | 30 | 4.0× |
| Easyjson | 2549 | 193.82 MB/s | 448 | 3 | 2.7× |
| Goccy | 2736 | 180.59 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3325 | 148.55 MB/s | 528 | 7 | 2.1× |
| Stdlib | 6956 | 71.02 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 495 | 464.86 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 498 | 461.90 MB/s | 160 | 1 | 9.9× |
| SonicFastest | 924 | 248.96 MB/s | 800 | 8 | 5.3× |
| Sonic | 924 | 248.94 MB/s | 801 | 8 | 5.3× |
| LightningDecodeAny | 1566 | 146.24 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1698 | 135.43 MB/s | 448 | 3 | 2.9× |
| Goccy | 1838 | 125.15 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2585 | 88.98 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4907 | 46.87 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 91570 | 711.28 MB/s | 166210 | 102 | 7.5× |
| Lightning | 93213 | 698.75 MB/s | 172434 | 107 | 7.3× |
| Sonic | 145896 | 446.43 MB/s | 235779 | 65 | 4.7× |
| SonicFastest | 146096 | 445.82 MB/s | 235839 | 65 | 4.7× |
| Goccy | 191226 | 340.60 MB/s | 228197 | 134 | 3.6× |
| LightningDecodeAny | 201788 | 264.28 MB/s | 176961 | 3252 | 3.4× |
| JSONV2 | 267052 | 243.89 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 683007 | 95.36 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2600018 | 746.33 MB/s | 2846866 | 2238 | 10.8× |
| LightningDestructive | 2643812 | 733.97 MB/s | 2847303 | 2238 | 10.6× |
| Goccy | 4991606 | 388.75 MB/s | 4062692 | 13509 | 5.6× |
| Sonic | 5002108 | 387.93 MB/s | 4879175 | 1736 | 5.6× |
| SonicFastest | 5020611 | 386.50 MB/s | 4878932 | 1736 | 5.6× |
| Easyjson | 8261206 | 234.89 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 10053915 | 193.01 MB/s | 7013525 | 219937 | 2.8× |
| JSONV2 | 12950379 | 149.84 MB/s | 3237182 | 13947 | 2.2× |
| Stdlib | 28032551 | 69.22 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1237759 | 2688.59 MB/s | 352059 | 1286 | 19.4× |
| Lightning | 1676744 | 1984.70 MB/s | 2488907 | 2995 | 14.3× |
| SonicFastest | 2036208 | 1634.33 MB/s | 5896509 | 4263 | 11.8× |
| Sonic | 2050430 | 1622.99 MB/s | 5896208 | 4263 | 11.7× |
| LightningDecodeAny | 3695003 | 831.87 MB/s | 4886621 | 56892 | 6.5× |
| Goccy | 4766637 | 698.15 MB/s | 3948916 | 3817 | 5.0× |
| JSONV2 | 7660480 | 434.42 MB/s | 5364505 | 13243 | 3.1× |
| Stdlib | 23978621 | 138.78 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 232945 | 945.92 MB/s | 136896 | 228 | 10.5× |
| LightningDestructive | 241480 | 912.48 MB/s | 136900 | 228 | 10.1× |
| Goccy | 488355 | 451.20 MB/s | 363696 | 1066 | 5.0× |
| Sonic | 489094 | 450.52 MB/s | 352124 | 262 | 5.0× |
| SonicFastest | 490178 | 449.52 MB/s | 352136 | 262 | 5.0× |
| Easyjson | 628638 | 350.51 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 766799 | 287.36 MB/s | 129746 | 470 | 3.2× |
| LightningDecodeAny | 1029325 | 105.23 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2447331 | 90.04 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14999258 | 540.03 MB/s | 15070013 | 51649 | 7.5× |
| Lightning | 15113038 | 535.96 MB/s | 15059842 | 51649 | 7.4× |
| Sonic | 17617919 | 459.76 MB/s | 19857345 | 41640 | 6.4× |
| SonicFastest | 17670558 | 458.39 MB/s | 19859801 | 41640 | 6.3× |
| Goccy | 27223424 | 297.54 MB/s | 18906820 | 107155 | 4.1× |
| Easyjson | 35506364 | 228.13 MB/s | 15059616 | 41643 | 3.2× |
| LightningDecodeAny | 44064004 | 118.08 MB/s | 34333349 | 912810 | 2.5× |
| JSONV2 | 48946427 | 165.49 MB/s | 15233747 | 78972 | 2.3× |
| Stdlib | 112004798 | 72.32 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6159264 | 484.39 MB/s | 3985337 | 30015 | 9.3× |
| LightningDestructive | 6168459 | 483.66 MB/s | 3986876 | 30015 | 9.3× |
| Sonic | 9265448 | 322.00 MB/s | 9131631 | 57804 | 6.2× |
| SonicFastest | 9318345 | 320.17 MB/s | 9131247 | 57804 | 6.2× |
| Goccy | 18887982 | 157.96 MB/s | 9957101 | 273624 | 3.0× |
| Easyjson | 19334964 | 154.30 MB/s | 9479441 | 30115 | 3.0× |
| LightningDecodeAny | 20574893 | 89.15 MB/s | 20023836 | 408557 | 2.8× |
| JSONV2 | 27283328 | 109.35 MB/s | 9257055 | 86278 | 2.1× |
| Stdlib | 57456421 | 51.93 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1500636 | 482.19 MB/s | 907477 | 3618 | 9.4× |
| Lightning | 1540208 | 469.80 MB/s | 907387 | 3618 | 9.1× |
| Sonic | 2130869 | 339.58 MB/s | 2372285 | 3683 | 6.6× |
| SonicFastest | 2138896 | 338.30 MB/s | 2373058 | 3683 | 6.6× |
| Easyjson | 5469771 | 132.29 MB/s | 2847908 | 3698 | 2.6× |
| LightningDecodeAny | 5531622 | 117.61 MB/s | 5752203 | 80172 | 2.5× |
| Goccy | 5632310 | 128.47 MB/s | 2708354 | 80267 | 2.5× |
| JSONV2 | 6278392 | 115.25 MB/s | 2704702 | 7318 | 2.2× |
| Stdlib | 14079207 | 51.39 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2085261 | 756.43 MB/s | 907389 | 3618 | 9.5× |
| LightningDestructive | 2184893 | 721.94 MB/s | 907673 | 3618 | 9.1× |
| SonicFastest | 2453862 | 642.80 MB/s | 3224113 | 3683 | 8.1× |
| Sonic | 2503777 | 629.99 MB/s | 3223208 | 3683 | 7.9× |
| LightningDecodeAny | 4684367 | 160.83 MB/s | 5752199 | 80172 | 4.2× |
| Easyjson | 6673209 | 236.37 MB/s | 2847905 | 3698 | 3.0× |
| JSONV2 | 6911798 | 228.21 MB/s | 2704552 | 7318 | 2.9× |
| Goccy | 6937120 | 227.38 MB/s | 3499671 | 80262 | 2.9× |
| Stdlib | 19828450 | 79.55 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 249492 | 601.72 MB/s | 81920 | 1 | 8.7× |
| LightningDestructive | 251952 | 595.84 MB/s | 81923 | 1 | 8.6× |
| Sonic | 380677 | 394.36 MB/s | 407513 | 16 | 5.7× |
| SonicFastest | 405034 | 370.65 MB/s | 408226 | 16 | 5.3× |
| LightningDecodeAny | 596541 | 251.65 MB/s | 746006 | 10020 | 3.6× |
| Goccy | 1023098 | 146.73 MB/s | 324517 | 10005 | 2.1× |
| JSONV2 | 1106698 | 135.65 MB/s | 357723 | 20 | 2.0× |
| Stdlib | 2165460 | 69.33 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 36060 | 779.72 MB/s | 30144 | 103 | 9.5× |
| Lightning | 36128 | 778.26 MB/s | 30272 | 105 | 9.5× |
| Sonic | 60100 | 467.84 MB/s | 59503 | 83 | 5.7× |
| SonicFastest | 60996 | 460.97 MB/s | 59537 | 83 | 5.6× |
| Easyjson | 80138 | 350.86 MB/s | 32304 | 138 | 4.3× |
| Goccy | 83535 | 336.59 MB/s | 59258 | 188 | 4.1× |
| JSONV2 | 145401 | 193.38 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 169695 | 165.69 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 343376 | 81.88 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1996 | 1166.34 MB/s | 32 | 1 | 13.2× |
| LightningDestructive | 1998 | 1165.29 MB/s | 32 | 1 | 13.2× |
| SonicFastest | 4762 | 488.85 MB/s | 3713 | 4 | 5.5× |
| Sonic | 4764 | 488.68 MB/s | 3714 | 4 | 5.5× |
| Goccy | 4887 | 476.37 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 5444 | 427.63 MB/s | 192 | 2 | 4.8× |
| JSONV2 | 8831 | 263.62 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 10753 | 156.71 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 26340 | 88.38 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 214 | 881.98 MB/s | 0 | 0 | 13.1× |
| LightningDestructive | 221 | 854.83 MB/s | 0 | 0 | 12.7× |
| Goccy | 446 | 423.90 MB/s | 304 | 2 | 6.3× |
| Easyjson | 550 | 343.73 MB/s | 0 | 0 | 5.1× |
| SonicFastest | 638 | 296.42 MB/s | 341 | 3 | 4.4× |
| Sonic | 644 | 293.72 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1072 | 176.28 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1358 | 98.71 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2799 | 67.53 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1456 | 1505.11 MB/s | 0 | 0 | 12.9× |
| LightningDestructive | 1469 | 1491.93 MB/s | 0 | 0 | 12.8× |
| Easyjson | 3523 | 621.89 MB/s | 24 | 1 | 5.3× |
| Goccy | 3880 | 564.69 MB/s | 2864 | 4 | 4.8× |
| SonicFastest | 6610 | 331.45 MB/s | 3599 | 38 | 2.8× |
| Sonic | 6841 | 320.29 MB/s | 3602 | 38 | 2.7× |
| JSONV2 | 8571 | 255.62 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9373 | 193.21 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18749 | 116.86 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 730125 | 699.16 MB/s | 703808 | 1012 | 9.7× |
| Lightning | 750808 | 679.90 MB/s | 703778 | 1012 | 9.5× |
| Goccy | 1257157 | 406.06 MB/s | 1138776 | 5006 | 5.7× |
| Sonic | 1297106 | 393.55 MB/s | 1310525 | 2014 | 5.5× |
| SonicFastest | 1299634 | 392.78 MB/s | 1309817 | 2014 | 5.5× |
| Easyjson | 1760684 | 289.93 MB/s | 863779 | 3012 | 4.0× |
| LightningDecodeAny | 3646157 | 126.56 MB/s | 2785927 | 66022 | 2.0× |
| JSONV2 | 3655433 | 139.65 MB/s | 1075967 | 12645 | 1.9× |
| Stdlib | 7110353 | 71.79 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 544 | 36346.70 MB/s | 0 | 0 | 300.3× |
| LightningDestructive | 883 | 22401.13 MB/s | 0 | 0 | 185.1× |
| SonicFastest | 7268 | 2722.92 MB/s | 21111 | 3 | 22.5× |
| Goccy | 25846 | 765.65 MB/s | 20492 | 2 | 6.3× |
| Sonic | 32122 | 616.05 MB/s | 20622 | 3 | 5.1× |
| JSONV2 | 33800 | 585.48 MB/s | 8 | 1 | 4.8× |
| Easyjson | 103113 | 191.92 MB/s | 0 | 0 | 1.6× |
| LightningDecodeAny | 103463 | 191.26 MB/s | 117104 | 2019 | 1.6× |
| Stdlib | 163505 | 121.03 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2276 | 7964.16 MB/s | 432 | 2 | 53.0× |
| LightningDestructive | 2393 | 7572.86 MB/s | 0 | 0 | 50.4× |
| Easyjson | 4799 | 3776.76 MB/s | 432 | 2 | 25.1× |
| SonicFastest | 8904 | 2035.43 MB/s | 20419 | 5 | 13.6× |
| Sonic | 9282 | 1952.65 MB/s | 20392 | 5 | 13.0× |
| LightningDecodeAny | 19296 | 926.71 MB/s | 29088 | 191 | 6.3× |
| Goccy | 26016 | 696.64 MB/s | 19460 | 2 | 4.6× |
| JSONV2 | 50355 | 359.93 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 120651 | 150.22 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2413474 | 832.20 MB/s | 2960792 | 7411 | 9.1× |
| Lightning | 2492209 | 805.91 MB/s | 2962101 | 7417 | 8.8× |
| SonicFastest | 4098370 | 490.07 MB/s | 5155957 | 7085 | 5.4× |
| Sonic | 4132164 | 486.06 MB/s | 5155050 | 7085 | 5.3× |
| Goccy | 4788753 | 419.42 MB/s | 5411076 | 15832 | 4.6× |
| Easyjson | 5332893 | 376.62 MB/s | 2981486 | 7439 | 4.1× |
| LightningDecodeAny | 7067310 | 161.63 MB/s | 7386073 | 134751 | 3.1× |
| JSONV2 | 7803054 | 257.40 MB/s | 3173673 | 14562 | 2.8× |
| Stdlib | 21988344 | 91.34 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 849 | 646.58 MB/s | 480 | 1 | 8.0× |
| LightningDestructive | 850 | 646.21 MB/s | 480 | 1 | 8.0× |
| LightningDecodeAny | 2155 | 254.32 MB/s | 2261 | 50 | 3.2× |
| Easyjson | 2278 | 240.97 MB/s | 1616 | 5 | 3.0× |
| Sonic | 2395 | 229.21 MB/s | 2260 | 8 | 2.9× |
| SonicFastest | 2412 | 227.58 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3422 | 160.41 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3584 | 153.18 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6828 | 80.41 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 532938 | 1184.97 MB/s | 364485 | 566 | 12.0× |
| Lightning | 567772 | 1112.27 MB/s | 413001 | 878 | 11.3× |
| Sonic | 964427 | 654.81 MB/s | 1064828 | 814 | 6.6× |
| SonicFastest | 979797 | 644.54 MB/s | 1065369 | 814 | 6.5× |
| Goccy | 1310232 | 481.99 MB/s | 989344 | 1200 | 4.9× |
| Easyjson | 1400181 | 451.02 MB/s | 422504 | 936 | 4.6× |
| JSONV2 | 2460736 | 256.64 MB/s | 571595 | 3144 | 2.6× |
| LightningDecodeAny | 2566604 | 181.92 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6403823 | 98.62 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 846781 | 664.17 MB/s | 544288 | 448 | 7.1× |
| Lightning | 1032249 | 544.84 MB/s | 767621 | 1254 | 5.8× |
| Sonic | 1258553 | 446.87 MB/s | 1348324 | 1185 | 4.8× |
| SonicFastest | 1281578 | 438.84 MB/s | 1348689 | 1185 | 4.7× |
| Goccy | 1497054 | 375.68 MB/s | 1039111 | 1028 | 4.0× |
| Easyjson | 2167967 | 259.42 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 3114472 | 180.58 MB/s | 2114149 | 30295 | 1.9× |
| JSONV2 | 3198657 | 175.83 MB/s | 927404 | 3482 | 1.9× |
| Stdlib | 5999471 | 93.74 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 631953 | 843.70 MB/s | 333444 | 2084 | 10.0× |
| Lightning | 705418 | 755.83 MB/s | 368224 | 2293 | 9.0× |
| Sonic | 1095356 | 486.76 MB/s | 978814 | 3082 | 5.8× |
| SonicFastest | 1096774 | 486.13 MB/s | 978962 | 3082 | 5.8× |
| Easyjson | 1293555 | 412.18 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1518192 | 351.19 MB/s | 1167071 | 5408 | 4.2× |
| JSONV2 | 2912471 | 183.07 MB/s | 745427 | 13288 | 2.2× |
| LightningDecodeAny | 3466174 | 153.82 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6334965 | 84.16 MB/s | 798693 | 17133 | 1.0× |
