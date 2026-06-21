# JSON Deserialization Benchmarks

- generated 2026-06-21T14:51:24Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104671 | 1215.95 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 107393 | 1185.13 MB/s | 49281 | 2 | 10.3× |
| SonicFastest | 182275 | 698.26 MB/s | 193603 | 10 | 6.1× |
| Sonic | 183401 | 693.97 MB/s | 197550 | 10 | 6.0× |
| Goccy | 190434 | 668.34 MB/s | 224730 | 884 | 5.8× |
| Easyjson | 212079 | 600.13 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 423709 | 300.38 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 449318 | 210.66 MB/s | 465585 | 9714 | 2.5× |
| Stdlib | 1104994 | 115.18 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3807697 | 591.18 MB/s | 3122874 | 3081 | 6.9× |
| LightningDestructive | 3838584 | 586.43 MB/s | 3123593 | 3081 | 6.8× |
| SonicFastest | 4514847 | 498.59 MB/s | 15233807 | 970 | 5.8× |
| Sonic | 4619600 | 487.28 MB/s | 15238825 | 970 | 5.7× |
| Goccy | 10625737 | 211.85 MB/s | 4131951 | 56533 | 2.5× |
| Easyjson | 11183506 | 201.28 MB/s | 3099813 | 2120 | 2.3× |
| LightningDecodeAny | 11277274 | 199.61 MB/s | 7938299 | 281383 | 2.3× |
| JSONV2 | 15994814 | 140.74 MB/s | 3123206 | 3083 | 1.6× |
| Stdlib | 26112900 | 86.20 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 486781 | 555.49 MB/s | 348024 | 1627 | 6.9× |
| LightningDestructive | 504594 | 535.88 MB/s | 348036 | 1627 | 6.6× |
| Sonic | 630269 | 429.03 MB/s | 481597 | 968 | 5.3× |
| SonicFastest | 632630 | 427.43 MB/s | 486124 | 968 | 5.3× |
| Easyjson | 1401868 | 192.89 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1415772 | 190.99 MB/s | 542346 | 8122 | 2.4× |
| LightningDecodeAny | 1596510 | 169.37 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2078172 | 130.12 MB/s | 348151 | 1628 | 1.6× |
| Stdlib | 3353736 | 80.63 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1390897 | 1241.79 MB/s | 959890 | 5882 | 9.5× |
| LightningDestructive | 1425173 | 1211.93 MB/s | 960047 | 5881 | 9.3× |
| Sonic | 2043889 | 845.06 MB/s | 2738699 | 4020 | 6.5× |
| SonicFastest | 2051543 | 841.90 MB/s | 2717381 | 4020 | 6.5× |
| Goccy | 2378832 | 726.07 MB/s | 2583680 | 14605 | 5.6× |
| Easyjson | 4250488 | 406.35 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4350664 | 397.00 MB/s | 1011638 | 7594 | 3.0× |
| LightningDecodeAny | 4620640 | 108.27 MB/s | 4976204 | 81466 | 2.9× |
| Stdlib | 13235162 | 130.50 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1186 | 1528.21 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1230 | 1473.03 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2554 | 709.53 MB/s | 24 | 1 | 5.5× |
| Goccy | 2825 | 641.44 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6108 | 296.66 MB/s | 3872 | 40 | 2.3× |
| Sonic | 6121 | 296.01 MB/s | 3889 | 40 | 2.3× |
| JSONV2 | 7700 | 235.34 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8505 | 212.94 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14051 | 128.96 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1214 | 1492.88 MB/s | 0 | 0 | 11.6× |
| LightningDestructive | 1259 | 1439.60 MB/s | 0 | 0 | 11.2× |
| Easyjson | 2548 | 711.20 MB/s | 24 | 1 | 5.5× |
| Goccy | 2782 | 651.32 MB/s | 2608 | 4 | 5.1× |
| Sonic | 6003 | 301.87 MB/s | 3704 | 40 | 2.3× |
| SonicFastest | 6018 | 301.08 MB/s | 3727 | 40 | 2.3× |
| JSONV2 | 7843 | 231.04 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8152 | 222.17 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14090 | 128.60 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1388 | 1305.42 MB/s | 144 | 10 | 10.1× |
| LightningDestructive | 1433 | 1264.57 MB/s | 144 | 10 | 9.8× |
| Easyjson | 2746 | 659.76 MB/s | 144 | 10 | 5.1× |
| Goccy | 2923 | 619.98 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6131 | 295.52 MB/s | 3802 | 42 | 2.3× |
| Sonic | 6191 | 292.67 MB/s | 3825 | 42 | 2.3× |
| JSONV2 | 7963 | 227.55 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8225 | 220.18 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14055 | 128.92 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 772 | 639.81 MB/s | 160 | 1 | 7.2× |
| LightningDestructive | 789 | 625.85 MB/s | 160 | 1 | 7.0× |
| SonicFastest | 1223 | 403.79 MB/s | 978 | 6 | 4.5× |
| Sonic | 1226 | 402.84 MB/s | 977 | 6 | 4.5× |
| LightningDecodeAny | 1649 | 299.02 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2204 | 224.13 MB/s | 448 | 3 | 2.5× |
| Goccy | 2425 | 203.71 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3258 | 151.61 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5522 | 89.46 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 462 | 497.35 MB/s | 160 | 1 | 8.9× |
| LightningDestructive | 472 | 487.84 MB/s | 160 | 1 | 8.8× |
| Sonic | 892 | 257.88 MB/s | 685 | 6 | 4.6× |
| SonicFastest | 892 | 257.86 MB/s | 690 | 6 | 4.6× |
| Easyjson | 1399 | 164.42 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1426 | 160.54 MB/s | 1536 | 30 | 2.9× |
| Goccy | 1603 | 143.48 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2452 | 93.81 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4134 | 55.64 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 75758 | 859.73 MB/s | 172432 | 107 | 7.3× |
| LightningDestructive | 77670 | 838.57 MB/s | 166209 | 102 | 7.1× |
| SonicFastest | 98104 | 663.91 MB/s | 156230 | 75 | 5.6× |
| Sonic | 100285 | 649.47 MB/s | 156476 | 75 | 5.5× |
| Goccy | 145020 | 449.12 MB/s | 229145 | 134 | 3.8× |
| LightningDecodeAny | 184692 | 288.75 MB/s | 176960 | 3252 | 3.0× |
| JSONV2 | 223017 | 292.05 MB/s | 206650 | 607 | 2.5× |
| Stdlib | 551596 | 118.08 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2661999 | 728.95 MB/s | 2846865 | 2238 | 8.8× |
| LightningDestructive | 2674772 | 725.47 MB/s | 2847301 | 2238 | 8.8× |
| Sonic | 4547448 | 426.72 MB/s | 14608596 | 1407 | 5.2× |
| SonicFastest | 4575171 | 424.13 MB/s | 14606973 | 1407 | 5.1× |
| Goccy | 4766250 | 407.13 MB/s | 4065796 | 13510 | 4.9× |
| Easyjson | 7539365 | 257.38 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9646987 | 201.15 MB/s | 7013524 | 219937 | 2.4× |
| JSONV2 | 11288730 | 171.89 MB/s | 3237215 | 13947 | 2.1× |
| Stdlib | 23456875 | 82.73 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1231089 | 2703.16 MB/s | 352053 | 1286 | 16.9× |
| Lightning | 1806208 | 1842.44 MB/s | 2488905 | 2995 | 11.5× |
| Sonic | 2647187 | 1257.12 MB/s | 6441478 | 4248 | 7.9× |
| SonicFastest | 2669021 | 1246.84 MB/s | 6501227 | 4248 | 7.8× |
| LightningDecodeAny | 3745226 | 820.72 MB/s | 4886622 | 56892 | 5.6× |
| Goccy | 4506010 | 738.53 MB/s | 3948909 | 3816 | 4.6× |
| JSONV2 | 7445643 | 446.95 MB/s | 5364523 | 13243 | 2.8× |
| Stdlib | 20831155 | 159.75 MB/s | 5565610 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 229286 | 961.01 MB/s | 136896 | 228 | 8.9× |
| LightningDestructive | 239636 | 919.50 MB/s | 136900 | 228 | 8.5× |
| SonicFastest | 377642 | 583.48 MB/s | 302026 | 398 | 5.4× |
| Sonic | 378812 | 581.68 MB/s | 301873 | 398 | 5.4× |
| Goccy | 428516 | 514.21 MB/s | 364369 | 1067 | 4.8× |
| Easyjson | 549504 | 400.99 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 720344 | 305.89 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 880929 | 122.95 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2039850 | 108.02 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 13256304 | 611.03 MB/s | 15068746 | 51649 | 6.7× |
| Lightning | 13370701 | 605.81 MB/s | 15059838 | 51649 | 6.7× |
| SonicFastest | 17141668 | 472.54 MB/s | 70902336 | 40014 | 5.2× |
| Sonic | 17229035 | 470.14 MB/s | 70887743 | 40014 | 5.2× |
| Goccy | 23913474 | 338.72 MB/s | 16906681 | 107148 | 3.7× |
| Easyjson | 30668956 | 264.11 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 40278477 | 129.18 MB/s | 34333352 | 912810 | 2.2× |
| JSONV2 | 43605814 | 185.76 MB/s | 15233741 | 78972 | 2.0× |
| Stdlib | 89371671 | 90.63 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6158005 | 484.49 MB/s | 3985339 | 30015 | 7.7× |
| LightningDestructive | 6190475 | 481.94 MB/s | 3986877 | 30015 | 7.6× |
| Sonic | 8664055 | 344.35 MB/s | 26610137 | 56760 | 5.4× |
| SonicFastest | 8674331 | 343.94 MB/s | 26522277 | 56760 | 5.4× |
| Goccy | 16488294 | 180.94 MB/s | 10705922 | 273651 | 2.9× |
| Easyjson | 16709864 | 178.55 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 19006367 | 96.50 MB/s | 20023839 | 408557 | 2.5× |
| JSONV2 | 24074277 | 123.93 MB/s | 9257155 | 86278 | 2.0× |
| Stdlib | 47160191 | 63.26 MB/s | 9258095 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1367605 | 529.10 MB/s | 907388 | 3618 | 8.4× |
| LightningDestructive | 1384049 | 522.81 MB/s | 907472 | 3618 | 8.3× |
| Sonic | 1761220 | 410.85 MB/s | 3195498 | 7226 | 6.6× |
| SonicFastest | 1773133 | 408.09 MB/s | 3202886 | 7226 | 6.5× |
| LightningDecodeAny | 4251935 | 153.01 MB/s | 5752202 | 80172 | 2.7× |
| Easyjson | 4253887 | 170.10 MB/s | 2847905 | 3698 | 2.7× |
| Goccy | 4716197 | 153.43 MB/s | 2878875 | 80277 | 2.4× |
| JSONV2 | 5647829 | 128.12 MB/s | 2704651 | 7318 | 2.0× |
| Stdlib | 11546037 | 62.67 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1967095 | 801.87 MB/s | 907387 | 3618 | 8.0× |
| LightningDestructive | 1980416 | 796.48 MB/s | 907648 | 3618 | 7.9× |
| SonicFastest | 2279156 | 692.08 MB/s | 5793102 | 7226 | 6.9× |
| Sonic | 2281583 | 691.34 MB/s | 5792352 | 7226 | 6.9× |
| LightningDecodeAny | 3829377 | 196.74 MB/s | 5752202 | 80172 | 4.1× |
| Easyjson | 5634283 | 279.96 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5757558 | 273.96 MB/s | 3652096 | 80270 | 2.7× |
| JSONV2 | 6534006 | 241.41 MB/s | 2704595 | 7318 | 2.4× |
| Stdlib | 15711967 | 100.39 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 223948 | 670.35 MB/s | 81920 | 1 | 8.1× |
| LightningDestructive | 227424 | 660.11 MB/s | 81923 | 1 | 8.0× |
| Sonic | 267664 | 560.87 MB/s | 250708 | 6 | 6.8× |
| SonicFastest | 272278 | 551.36 MB/s | 266930 | 6 | 6.7× |
| LightningDecodeAny | 474070 | 316.66 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 867926 | 172.97 MB/s | 328065 | 10005 | 2.1× |
| JSONV2 | 1092613 | 137.40 MB/s | 357713 | 20 | 1.7× |
| Stdlib | 1823169 | 82.34 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 35412 | 793.99 MB/s | 30272 | 105 | 8.6× |
| LightningDestructive | 35775 | 785.95 MB/s | 30144 | 103 | 8.5× |
| Sonic | 63445 | 443.17 MB/s | 47220 | 103 | 4.8× |
| SonicFastest | 63559 | 442.38 MB/s | 47387 | 103 | 4.8× |
| Easyjson | 68780 | 408.80 MB/s | 32304 | 138 | 4.4× |
| Goccy | 70244 | 400.27 MB/s | 59205 | 188 | 4.3× |
| JSONV2 | 134604 | 208.89 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 150920 | 186.30 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303718 | 92.58 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1957 | 1189.46 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2002 | 1162.85 MB/s | 32 | 1 | 11.3× |
| Goccy | 4114 | 565.81 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4215 | 552.34 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5010 | 464.70 MB/s | 4195 | 6 | 4.5× |
| Sonic | 5032 | 462.61 MB/s | 4232 | 6 | 4.5× |
| JSONV2 | 8467 | 274.96 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10222 | 164.83 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22700 | 102.56 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220 | 860.61 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 224 | 842.03 MB/s | 0 | 0 | 10.8× |
| Goccy | 381 | 495.91 MB/s | 304 | 2 | 6.4× |
| Easyjson | 492 | 383.87 MB/s | 0 | 0 | 4.9× |
| SonicFastest | 796 | 237.36 MB/s | 497 | 4 | 3.1× |
| Sonic | 798 | 236.70 MB/s | 501 | 4 | 3.0× |
| JSONV2 | 1037 | 182.31 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1233 | 108.65 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2431 | 77.74 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1523 | 1438.68 MB/s | 0 | 0 | 10.5× |
| LightningDestructive | 1575 | 1391.35 MB/s | 0 | 0 | 10.2× |
| Easyjson | 3186 | 687.69 MB/s | 24 | 1 | 5.0× |
| Goccy | 3200 | 684.79 MB/s | 2864 | 4 | 5.0× |
| Sonic | 6538 | 335.14 MB/s | 4085 | 40 | 2.5× |
| SonicFastest | 6541 | 334.96 MB/s | 4074 | 40 | 2.5× |
| JSONV2 | 8053 | 272.07 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8424 | 214.98 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16033 | 136.66 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 726225 | 702.92 MB/s | 703808 | 1012 | 8.4× |
| Lightning | 728471 | 700.75 MB/s | 703779 | 1012 | 8.3× |
| Goccy | 1149466 | 444.10 MB/s | 1138213 | 5006 | 5.3× |
| Sonic | 1150986 | 443.51 MB/s | 895385 | 2006 | 5.3× |
| SonicFastest | 1154405 | 442.20 MB/s | 901682 | 2006 | 5.3× |
| Easyjson | 1536385 | 332.26 MB/s | 863777 | 3012 | 4.0× |
| JSONV2 | 3232490 | 157.92 MB/s | 1076012 | 12646 | 1.9× |
| LightningDecodeAny | 3501806 | 131.78 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6077094 | 84.00 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1335 | 14827.12 MB/s | 0 | 0 | 83.5× |
| LightningDestructive | 1732 | 11425.58 MB/s | 0 | 0 | 64.4× |
| Goccy | 19935 | 992.69 MB/s | 20491 | 2 | 5.6× |
| Sonic | 27071 | 731.00 MB/s | 22198 | 4 | 4.1× |
| SonicFastest | 27097 | 730.31 MB/s | 22286 | 4 | 4.1× |
| JSONV2 | 29746 | 665.27 MB/s | 8 | 1 | 3.7× |
| LightningDecodeAny | 77851 | 254.18 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86020 | 230.05 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111534 | 177.43 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2833 | 6396.46 MB/s | 432 | 2 | 36.4× |
| LightningDestructive | 2967 | 6108.19 MB/s | 0 | 0 | 34.7× |
| Easyjson | 3964 | 4572.08 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 9827 | 1844.39 MB/s | 23005 | 6 | 10.5× |
| Sonic | 9845 | 1840.93 MB/s | 23107 | 6 | 10.5× |
| Goccy | 15591 | 1162.44 MB/s | 19459 | 2 | 6.6× |
| LightningDecodeAny | 17105 | 1045.41 MB/s | 29088 | 191 | 6.0× |
| JSONV2 | 45575 | 397.67 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103093 | 175.80 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2364775 | 849.34 MB/s | 2962103 | 7417 | 7.9× |
| LightningDestructive | 2371312 | 847.00 MB/s | 2960787 | 7411 | 7.9× |
| Goccy | 4243729 | 473.29 MB/s | 5412896 | 15836 | 4.4× |
| Sonic | 4552852 | 441.15 MB/s | 10949315 | 13683 | 4.1× |
| SonicFastest | 4573089 | 439.20 MB/s | 10939733 | 13683 | 4.1× |
| Easyjson | 4969023 | 404.20 MB/s | 2981514 | 7439 | 3.8× |
| JSONV2 | 6920105 | 290.24 MB/s | 3173687 | 14563 | 2.7× |
| LightningDecodeAny | 7226903 | 158.06 MB/s | 7386074 | 134751 | 2.6× |
| Stdlib | 18721734 | 107.28 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 874 | 627.98 MB/s | 480 | 1 | 6.5× |
| LightningDestructive | 896 | 612.94 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1902 | 288.09 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2164 | 253.67 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2658 | 206.56 MB/s | 1947 | 26 | 2.1× |
| Sonic | 2662 | 206.26 MB/s | 1940 | 26 | 2.1× |
| Goccy | 3071 | 178.76 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3353 | 163.71 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5649 | 97.18 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 533013 | 1184.80 MB/s | 364485 | 566 | 10.1× |
| Lightning | 556974 | 1133.83 MB/s | 413002 | 878 | 9.7× |
| Sonic | 1008359 | 626.28 MB/s | 982824 | 1102 | 5.4× |
| SonicFastest | 1018335 | 620.14 MB/s | 987230 | 1102 | 5.3× |
| Easyjson | 1145372 | 551.36 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1172212 | 538.74 MB/s | 984858 | 1201 | 4.6× |
| JSONV2 | 2158602 | 292.56 MB/s | 571620 | 3144 | 2.5× |
| LightningDecodeAny | 2388435 | 195.49 MB/s | 2010201 | 30295 | 2.3× |
| Stdlib | 5404447 | 116.85 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 745172 | 754.74 MB/s | 544283 | 448 | 7.1× |
| Lightning | 884183 | 636.08 MB/s | 767621 | 1254 | 6.0× |
| Sonic | 1036003 | 542.86 MB/s | 982880 | 1476 | 5.1× |
| SonicFastest | 1041532 | 539.98 MB/s | 991064 | 1476 | 5.1× |
| Goccy | 1345093 | 418.12 MB/s | 1046288 | 1030 | 3.9× |
| Easyjson | 1738284 | 323.54 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2756082 | 204.06 MB/s | 2114152 | 30295 | 1.9× |
| JSONV2 | 2768388 | 203.15 MB/s | 927440 | 3482 | 1.9× |
| Stdlib | 5282538 | 106.47 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 669576 | 796.29 MB/s | 333446 | 2084 | 8.2× |
| Lightning | 673213 | 791.99 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1104675 | 482.66 MB/s | 428361 | 3273 | 5.0× |
| Sonic | 1144931 | 465.69 MB/s | 1047263 | 4351 | 4.8× |
| SonicFastest | 1146570 | 465.02 MB/s | 1035931 | 4351 | 4.8× |
| Goccy | 1300546 | 409.96 MB/s | 1167236 | 5409 | 4.2× |
| JSONV2 | 2534856 | 210.34 MB/s | 745446 | 13288 | 2.2× |
| LightningDecodeAny | 3343658 | 159.46 MB/s | 2674619 | 50596 | 1.6× |
| Stdlib | 5476756 | 97.35 MB/s | 798692 | 17133 | 1.0× |
