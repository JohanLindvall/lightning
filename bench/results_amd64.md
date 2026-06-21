# JSON Deserialization Benchmarks

- generated 2026-06-21T13:05:19Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 80618 | 1578.73 MB/s | 49760 | 3 | 12.0× |
| SonicFastest | 150407 | 846.20 MB/s | 214662 | 15 | 6.4× |
| Sonic | 150558 | 845.36 MB/s | 214808 | 15 | 6.4× |
| Easyjson | 172395 | 738.27 MB/s | 122864 | 14 | 5.6× |
| Goccy | 192086 | 662.59 MB/s | 225433 | 884 | 5.0× |
| JSONV2 | 319974 | 397.77 MB/s | 195128 | 1805 | 3.0× |
| LightningDecodeAny | 339996 | 278.39 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 963699 | 132.07 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3230038 | 696.91 MB/s | 3122874 | 3081 | 7.5× |
| SonicFastest | 3931092 | 572.63 MB/s | 4870010 | 2584 | 6.2× |
| Sonic | 4103347 | 548.59 MB/s | 4870752 | 2584 | 5.9× |
| LightningDecodeAny | 8782181 | 256.32 MB/s | 7938298 | 281383 | 2.8× |
| Goccy | 9982337 | 225.50 MB/s | 4152481 | 56533 | 2.4× |
| Easyjson | 10635019 | 211.66 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 13149827 | 171.18 MB/s | 3123194 | 3083 | 1.8× |
| Stdlib | 24234537 | 92.89 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 445530 | 606.92 MB/s | 348024 | 1627 | 7.1× |
| SonicFastest | 574543 | 470.64 MB/s | 641419 | 1147 | 5.5× |
| Sonic | 586576 | 460.99 MB/s | 642392 | 1147 | 5.4× |
| LightningDecodeAny | 1259410 | 214.71 MB/s | 1011485 | 37901 | 2.5× |
| Goccy | 1356662 | 199.31 MB/s | 542077 | 8122 | 2.3× |
| Easyjson | 1366562 | 197.87 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 1713753 | 157.78 MB/s | 348159 | 1628 | 1.8× |
| Stdlib | 3157153 | 85.65 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1044791 | 1653.16 MB/s | 959890 | 5882 | 12.4× |
| SonicFastest | 1687042 | 1023.81 MB/s | 2695737 | 5547 | 7.7× |
| Sonic | 1696393 | 1018.16 MB/s | 2695434 | 5547 | 7.7× |
| Goccy | 2230950 | 774.20 MB/s | 2580782 | 14603 | 5.8× |
| JSONV2 | 3070884 | 562.45 MB/s | 1011613 | 7594 | 4.2× |
| Easyjson | 3087429 | 559.43 MB/s | 972032 | 5389 | 4.2× |
| LightningDecodeAny | 3185667 | 157.05 MB/s | 4976204 | 81466 | 4.1× |
| Stdlib | 12980288 | 133.06 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 781 | 2320.75 MB/s | 0 | 0 | 14.7× |
| Easyjson | 2121 | 854.25 MB/s | 24 | 1 | 5.4× |
| Goccy | 2656 | 682.19 MB/s | 2608 | 4 | 4.3× |
| SonicFastest | 4824 | 375.58 MB/s | 3349 | 38 | 2.4× |
| Sonic | 4964 | 365.02 MB/s | 3347 | 38 | 2.3× |
| JSONV2 | 6068 | 298.60 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 6600 | 274.39 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 11473 | 157.93 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 836 | 2168.09 MB/s | 0 | 0 | 13.5× |
| Easyjson | 2144 | 845.18 MB/s | 24 | 1 | 5.3× |
| Goccy | 2656 | 682.25 MB/s | 2608 | 4 | 4.3× |
| SonicFastest | 4637 | 390.74 MB/s | 3350 | 38 | 2.4× |
| Sonic | 4831 | 375.08 MB/s | 3348 | 38 | 2.3× |
| JSONV2 | 6180 | 293.18 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 6660 | 271.91 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 11292 | 160.46 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 960 | 1887.61 MB/s | 144 | 10 | 11.8× |
| Easyjson | 2397 | 755.85 MB/s | 144 | 10 | 4.7× |
| Goccy | 2450 | 739.50 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 4952 | 365.93 MB/s | 3367 | 40 | 2.3× |
| Sonic | 5111 | 354.53 MB/s | 3367 | 40 | 2.2× |
| JSONV2 | 5927 | 305.71 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 6698 | 270.38 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 11364 | 159.45 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 572 | 864.00 MB/s | 160 | 1 | 8.2× |
| Sonic | 944 | 523.12 MB/s | 1075 | 8 | 5.0× |
| SonicFastest | 945 | 522.52 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1266 | 389.37 MB/s | 1536 | 30 | 3.7× |
| Easyjson | 1871 | 263.98 MB/s | 448 | 3 | 2.5× |
| Goccy | 2005 | 246.37 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 2334 | 211.68 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4680 | 105.56 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 360 | 638.45 MB/s | 160 | 1 | 9.1× |
| SonicFastest | 674 | 341.31 MB/s | 801 | 8 | 4.9× |
| Sonic | 678 | 339.13 MB/s | 801 | 8 | 4.8× |
| LightningDecodeAny | 1112 | 206.00 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1223 | 188.04 MB/s | 448 | 3 | 2.7× |
| Goccy | 1342 | 171.40 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 1814 | 126.81 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3289 | 69.93 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 68135 | 955.93 MB/s | 172433 | 107 | 7.5× |
| SonicFastest | 117755 | 553.12 MB/s | 235797 | 65 | 4.4× |
| Sonic | 118729 | 548.58 MB/s | 235781 | 65 | 4.3× |
| LightningDecodeAny | 145768 | 365.85 MB/s | 176961 | 3252 | 3.5× |
| Goccy | 149193 | 436.56 MB/s | 228737 | 134 | 3.4× |
| JSONV2 | 199073 | 327.18 MB/s | 206667 | 607 | 2.6× |
| Stdlib | 512334 | 127.13 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2060755 | 941.63 MB/s | 2846866 | 2238 | 9.5× |
| Goccy | 3771230 | 514.55 MB/s | 4062689 | 13509 | 5.2× |
| Sonic | 4865762 | 398.80 MB/s | 4879031 | 1736 | 4.0× |
| SonicFastest | 4956437 | 391.51 MB/s | 4880805 | 1736 | 4.0× |
| Easyjson | 6116749 | 317.24 MB/s | 3871266 | 15043 | 3.2× |
| LightningDecodeAny | 7472011 | 259.70 MB/s | 7013525 | 219937 | 2.6× |
| JSONV2 | 8586588 | 225.99 MB/s | 3237189 | 13947 | 2.3× |
| Stdlib | 19648232 | 98.76 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1260760 | 2639.54 MB/s | 2488906 | 2995 | 15.7× |
| SonicFastest | 1788133 | 1861.07 MB/s | 5895086 | 4263 | 11.1× |
| Sonic | 1791686 | 1857.37 MB/s | 5894623 | 4263 | 11.1× |
| LightningDecodeAny | 2699491 | 1138.65 MB/s | 4886622 | 56892 | 7.3× |
| Goccy | 4398010 | 756.67 MB/s | 3948914 | 3816 | 4.5× |
| JSONV2 | 6309453 | 527.44 MB/s | 5364504 | 13243 | 3.1× |
| Stdlib | 19828500 | 167.83 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 181067 | 1216.93 MB/s | 136896 | 228 | 9.6× |
| Goccy | 354344 | 621.84 MB/s | 363802 | 1066 | 4.9× |
| Sonic | 387268 | 568.98 MB/s | 351183 | 262 | 4.5× |
| SonicFastest | 387318 | 568.90 MB/s | 350937 | 262 | 4.5× |
| Easyjson | 441829 | 498.71 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 531367 | 414.68 MB/s | 129746 | 470 | 3.3× |
| LightningDecodeAny | 732619 | 147.84 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 1732924 | 127.15 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 10725376 | 755.22 MB/s | 15059841 | 51649 | 7.3× |
| SonicFastest | 15991750 | 506.51 MB/s | 19862624 | 41640 | 4.9× |
| Sonic | 16004691 | 506.10 MB/s | 19861910 | 41640 | 4.9× |
| Goccy | 20144198 | 402.10 MB/s | 19090875 | 107156 | 3.9× |
| Easyjson | 25365220 | 319.34 MB/s | 15059616 | 41643 | 3.1× |
| LightningDecodeAny | 32633800 | 159.44 MB/s | 34333347 | 912810 | 2.4× |
| JSONV2 | 35012715 | 231.35 MB/s | 15233732 | 78972 | 2.2× |
| Stdlib | 78016897 | 103.82 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4765167 | 626.10 MB/s | 3985339 | 30015 | 8.3× |
| Sonic | 7332790 | 406.87 MB/s | 9129238 | 57804 | 5.4× |
| SonicFastest | 7384409 | 404.02 MB/s | 9130525 | 57804 | 5.3× |
| Easyjson | 14081719 | 211.87 MB/s | 9479441 | 30115 | 2.8× |
| Goccy | 14115215 | 211.37 MB/s | 9904648 | 273621 | 2.8× |
| LightningDecodeAny | 14721670 | 124.59 MB/s | 20023834 | 408557 | 2.7× |
| JSONV2 | 19478049 | 153.17 MB/s | 9257068 | 86278 | 2.0× |
| Stdlib | 39364326 | 75.79 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1124087 | 643.72 MB/s | 907386 | 3618 | 8.6× |
| SonicFastest | 1653995 | 437.48 MB/s | 2374718 | 3683 | 5.8× |
| Sonic | 1663606 | 434.96 MB/s | 2374557 | 3683 | 5.8× |
| LightningDecodeAny | 3983247 | 163.33 MB/s | 5752205 | 80172 | 2.4× |
| Easyjson | 4120635 | 175.60 MB/s | 2847907 | 3698 | 2.3× |
| Goccy | 4152627 | 174.25 MB/s | 2700081 | 80267 | 2.3× |
| JSONV2 | 4736699 | 152.76 MB/s | 2704708 | 7318 | 2.0× |
| Stdlib | 9652611 | 74.96 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1626862 | 969.57 MB/s | 907386 | 3618 | 8.5× |
| Sonic | 1902801 | 828.96 MB/s | 3226457 | 3683 | 7.3× |
| SonicFastest | 1909524 | 826.05 MB/s | 3226550 | 3683 | 7.3× |
| LightningDecodeAny | 3323828 | 226.67 MB/s | 5752201 | 80172 | 4.2× |
| Easyjson | 5070799 | 311.07 MB/s | 2847905 | 3698 | 2.7× |
| Goccy | 5131544 | 307.38 MB/s | 3496584 | 80262 | 2.7× |
| JSONV2 | 5239628 | 301.04 MB/s | 2704552 | 7318 | 2.6× |
| Stdlib | 13869608 | 113.73 MB/s | 2704555 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 182896 | 820.82 MB/s | 81920 | 1 | 9.4× |
| Sonic | 290444 | 516.88 MB/s | 407807 | 16 | 5.9× |
| SonicFastest | 301899 | 497.27 MB/s | 408580 | 16 | 5.7× |
| LightningDecodeAny | 436262 | 344.11 MB/s | 746003 | 10020 | 3.9× |
| Goccy | 819457 | 183.20 MB/s | 328470 | 10005 | 2.1× |
| JSONV2 | 882416 | 170.13 MB/s | 357727 | 20 | 2.0× |
| Stdlib | 1722045 | 87.18 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 26347 | 1067.16 MB/s | 30272 | 105 | 9.1× |
| Sonic | 52044 | 540.25 MB/s | 59440 | 83 | 4.6× |
| SonicFastest | 52127 | 539.40 MB/s | 59472 | 83 | 4.6× |
| Easyjson | 57109 | 492.34 MB/s | 32304 | 138 | 4.2× |
| Goccy | 61003 | 460.91 MB/s | 59265 | 188 | 3.9× |
| JSONV2 | 100579 | 279.55 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 117767 | 238.75 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 240486 | 116.92 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1548 | 1503.67 MB/s | 32 | 1 | 11.9× |
| Goccy | 3557 | 654.43 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 4104 | 567.29 MB/s | 192 | 2 | 4.5× |
| SonicFastest | 4799 | 485.07 MB/s | 3712 | 4 | 3.8× |
| Sonic | 4825 | 482.50 MB/s | 3711 | 4 | 3.8× |
| JSONV2 | 5811 | 400.64 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 7747 | 217.49 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 18458 | 126.12 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 172 | 1100.43 MB/s | 0 | 0 | 11.1× |
| Goccy | 325 | 581.66 MB/s | 304 | 2 | 5.9× |
| Easyjson | 406 | 465.09 MB/s | 0 | 0 | 4.7× |
| SonicFastest | 588 | 321.72 MB/s | 341 | 3 | 3.2× |
| Sonic | 590 | 320.33 MB/s | 341 | 3 | 3.2× |
| JSONV2 | 724 | 261.14 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 920 | 145.67 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 1903 | 99.31 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1111 | 1972.54 MB/s | 0 | 0 | 12.0× |
| Easyjson | 2565 | 854.34 MB/s | 24 | 1 | 5.2× |
| Goccy | 2866 | 764.39 MB/s | 2864 | 4 | 4.7× |
| SonicFastest | 5085 | 430.89 MB/s | 3603 | 38 | 2.6× |
| Sonic | 5218 | 419.93 MB/s | 3600 | 38 | 2.6× |
| JSONV2 | 6313 | 347.06 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 6656 | 272.08 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 13377 | 163.79 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 529875 | 963.39 MB/s | 703777 | 1012 | 9.4× |
| Goccy | 956993 | 533.42 MB/s | 1135241 | 5006 | 5.2× |
| Sonic | 1209640 | 422.01 MB/s | 1308513 | 2014 | 4.1× |
| Easyjson | 1214423 | 420.34 MB/s | 863776 | 3012 | 4.1× |
| SonicFastest | 1223648 | 417.18 MB/s | 1307922 | 2014 | 4.1× |
| JSONV2 | 2410911 | 211.74 MB/s | 1075963 | 12645 | 2.1× |
| LightningDecodeAny | 2564433 | 179.95 MB/s | 2785930 | 66022 | 1.9× |
| Stdlib | 4973374 | 102.64 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 473 | 41830.44 MB/s | 0 | 0 | 233.9× |
| SonicFastest | 5308 | 3728.13 MB/s | 21082 | 3 | 20.8× |
| Goccy | 19495 | 1015.06 MB/s | 20492 | 2 | 5.7× |
| Sonic | 22629 | 874.48 MB/s | 20636 | 3 | 4.9× |
| JSONV2 | 28075 | 704.87 MB/s | 8 | 1 | 3.9× |
| LightningDecodeAny | 75091 | 263.52 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 90145 | 219.52 MB/s | 0 | 0 | 1.2× |
| Stdlib | 110662 | 178.82 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1726 | 10502.09 MB/s | 432 | 2 | 54.4× |
| Easyjson | 3495 | 5186.18 MB/s | 432 | 2 | 26.9× |
| SonicFastest | 7178 | 2525.09 MB/s | 20416 | 5 | 13.1× |
| Sonic | 7347 | 2466.99 MB/s | 20413 | 5 | 12.8× |
| LightningDecodeAny | 14130 | 1265.56 MB/s | 29088 | 191 | 6.6× |
| Goccy | 20205 | 897.00 MB/s | 19460 | 2 | 4.6× |
| JSONV2 | 37401 | 484.59 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 93889 | 193.04 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1822684 | 1101.94 MB/s | 2962101 | 7417 | 8.9× |
| Goccy | 3587523 | 559.86 MB/s | 5410858 | 15831 | 4.5× |
| SonicFastest | 3802305 | 528.23 MB/s | 5152402 | 7085 | 4.2× |
| Sonic | 3861376 | 520.15 MB/s | 5153064 | 7085 | 4.2× |
| Easyjson | 3932443 | 510.75 MB/s | 2981481 | 7439 | 4.1× |
| LightningDecodeAny | 5258677 | 217.22 MB/s | 7386074 | 134751 | 3.1× |
| JSONV2 | 5580332 | 359.92 MB/s | 3173677 | 14563 | 2.9× |
| Stdlib | 16151411 | 124.35 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 630 | 871.47 MB/s | 480 | 1 | 7.2× |
| LightningDecodeAny | 1485 | 368.92 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 1578 | 347.91 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 1690 | 324.88 MB/s | 2262 | 8 | 2.7× |
| Sonic | 1760 | 311.96 MB/s | 2263 | 8 | 2.6× |
| JSONV2 | 2353 | 233.28 MB/s | 1664 | 7 | 1.9× |
| Goccy | 2441 | 224.90 MB/s | 2129 | 43 | 1.8× |
| Stdlib | 4507 | 121.82 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 430294 | 1467.63 MB/s | 413001 | 878 | 11.0× |
| SonicFastest | 829925 | 760.93 MB/s | 1065332 | 814 | 5.7× |
| Sonic | 836440 | 755.00 MB/s | 1065904 | 814 | 5.7× |
| Goccy | 1003564 | 629.27 MB/s | 989406 | 1200 | 4.7× |
| Easyjson | 1010594 | 624.89 MB/s | 422504 | 936 | 4.7× |
| JSONV2 | 1678373 | 376.27 MB/s | 571590 | 3144 | 2.8× |
| LightningDecodeAny | 1898737 | 245.90 MB/s | 2010198 | 30295 | 2.5× |
| Stdlib | 4731090 | 133.48 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 738272 | 761.79 MB/s | 767620 | 1254 | 6.0× |
| Sonic | 1094344 | 513.92 MB/s | 1349667 | 1184 | 4.0× |
| SonicFastest | 1096553 | 512.89 MB/s | 1350137 | 1185 | 4.0× |
| Goccy | 1171418 | 480.11 MB/s | 1044420 | 1028 | 3.8× |
| Easyjson | 1563414 | 359.73 MB/s | 775153 | 1254 | 2.8× |
| LightningDecodeAny | 2317396 | 242.69 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 2347635 | 239.56 MB/s | 927405 | 3482 | 1.9× |
| Stdlib | 4415042 | 127.38 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 508180 | 1049.19 MB/s | 368224 | 2293 | 9.1× |
| Easyjson | 958460 | 556.29 MB/s | 428362 | 3273 | 4.8× |
| Sonic | 1027758 | 518.78 MB/s | 981764 | 3082 | 4.5× |
| SonicFastest | 1033799 | 515.75 MB/s | 981765 | 3082 | 4.5× |
| Goccy | 1183340 | 450.57 MB/s | 1167050 | 5408 | 3.9× |
| JSONV2 | 2030937 | 262.53 MB/s | 745420 | 13288 | 2.3× |
| LightningDecodeAny | 2566271 | 207.76 MB/s | 2674618 | 50596 | 1.8× |
| Stdlib | 4602587 | 115.84 MB/s | 798692 | 17133 | 1.0× |
