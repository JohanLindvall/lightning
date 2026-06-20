# JSON Deserialization Benchmarks

- generated 2026-06-20T16:46:45Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 82309 | 1546.30 MB/s | 50208 | 4 | 11.6× |
| Sonic | 150650 | 844.84 MB/s | 213994 | 15 | 6.4× |
| SonicFastest | 151154 | 842.02 MB/s | 214099 | 15 | 6.3× |
| Easyjson | 179354 | 709.63 MB/s | 122864 | 14 | 5.3× |
| Goccy | 188317 | 675.86 MB/s | 225564 | 884 | 5.1× |
| JSONV2 | 305935 | 416.02 MB/s | 195128 | 1805 | 3.1× |
| LightningDecodeAny | 348270 | 271.78 MB/s | 466034 | 9715 | 2.8× |
| Stdlib | 958831 | 132.74 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3222489 | 698.54 MB/s | 3122873 | 3081 | 7.5× |
| SonicFastest | 3920339 | 574.20 MB/s | 4873169 | 2584 | 6.1× |
| Sonic | 4101904 | 548.78 MB/s | 4869739 | 2584 | 5.9× |
| LightningDecodeAny | 8755800 | 257.09 MB/s | 7938298 | 281383 | 2.7× |
| Goccy | 10009494 | 224.89 MB/s | 4169239 | 56534 | 2.4× |
| Easyjson | 10689073 | 210.59 MB/s | 3099809 | 2120 | 2.2× |
| JSONV2 | 13184371 | 170.74 MB/s | 3123239 | 3083 | 1.8× |
| Stdlib | 24035569 | 93.65 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 448355 | 603.10 MB/s | 348024 | 1627 | 7.0× |
| SonicFastest | 568629 | 475.53 MB/s | 642146 | 1147 | 5.5× |
| Sonic | 569344 | 474.94 MB/s | 642273 | 1147 | 5.5× |
| LightningDecodeAny | 1288272 | 209.90 MB/s | 1011485 | 37901 | 2.4× |
| Goccy | 1367098 | 197.79 MB/s | 542389 | 8122 | 2.3× |
| Easyjson | 1374871 | 196.68 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 1721024 | 157.12 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 3145593 | 85.96 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1041442 | 1658.47 MB/s | 964642 | 6177 | 12.4× |
| Sonic | 1721193 | 1003.49 MB/s | 2695309 | 5547 | 7.5× |
| SonicFastest | 1722621 | 1002.66 MB/s | 2695405 | 5547 | 7.5× |
| Goccy | 2250880 | 767.35 MB/s | 2580770 | 14603 | 5.8× |
| JSONV2 | 3092705 | 558.48 MB/s | 1011613 | 7594 | 4.2× |
| LightningDecodeAny | 3163701 | 158.14 MB/s | 4976251 | 81467 | 4.1× |
| Easyjson | 3170174 | 544.83 MB/s | 972032 | 5389 | 4.1× |
| Stdlib | 12950970 | 133.36 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 804 | 2254.91 MB/s | 0 | 0 | 14.2× |
| Easyjson | 2194 | 826.06 MB/s | 24 | 1 | 5.2× |
| Goccy | 2449 | 739.77 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 5003 | 362.19 MB/s | 3350 | 38 | 2.3× |
| Sonic | 5101 | 355.20 MB/s | 3348 | 38 | 2.2× |
| JSONV2 | 5794 | 312.73 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 6726 | 269.27 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 11440 | 158.39 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 858 | 2112.79 MB/s | 0 | 0 | 13.1× |
| Easyjson | 2199 | 824.07 MB/s | 24 | 1 | 5.1× |
| Goccy | 2448 | 740.32 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 4791 | 378.19 MB/s | 3350 | 38 | 2.3× |
| Sonic | 4931 | 367.47 MB/s | 3347 | 38 | 2.3× |
| JSONV2 | 5883 | 308.01 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 6729 | 269.15 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 11251 | 161.05 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1011 | 1792.50 MB/s | 144 | 10 | 11.2× |
| Easyjson | 2418 | 749.32 MB/s | 144 | 10 | 4.7× |
| Goccy | 2450 | 739.54 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 4782 | 378.92 MB/s | 3365 | 40 | 2.4× |
| Sonic | 4966 | 364.87 MB/s | 3370 | 40 | 2.3× |
| JSONV2 | 5913 | 306.45 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 6908 | 262.14 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 11312 | 160.19 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 568 | 869.21 MB/s | 160 | 1 | 8.3× |
| Sonic | 944 | 523.18 MB/s | 1075 | 8 | 5.0× |
| SonicFastest | 945 | 522.94 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1262 | 390.66 MB/s | 1536 | 30 | 3.7× |
| Easyjson | 1804 | 273.88 MB/s | 448 | 3 | 2.6× |
| Goccy | 1880 | 262.78 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 2402 | 205.63 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4695 | 105.22 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 360 | 639.33 MB/s | 160 | 1 | 9.1× |
| Sonic | 660 | 348.60 MB/s | 801 | 8 | 5.0× |
| SonicFastest | 662 | 347.67 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1089 | 210.31 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1193 | 192.86 MB/s | 448 | 3 | 2.8× |
| Goccy | 1295 | 177.60 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 1821 | 126.32 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3286 | 69.99 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 72305 | 900.79 MB/s | 178610 | 112 | 6.9× |
| Sonic | 118756 | 548.45 MB/s | 235865 | 65 | 4.2× |
| SonicFastest | 120485 | 540.58 MB/s | 236001 | 65 | 4.1× |
| LightningDecodeAny | 146863 | 363.12 MB/s | 183137 | 3257 | 3.4× |
| Goccy | 150654 | 432.33 MB/s | 228323 | 134 | 3.3× |
| JSONV2 | 196432 | 331.57 MB/s | 206666 | 607 | 2.5× |
| Stdlib | 498162 | 130.74 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2037518 | 952.37 MB/s | 2846867 | 2238 | 9.6× |
| Goccy | 3713231 | 522.58 MB/s | 4062629 | 13509 | 5.3× |
| Sonic | 4929779 | 393.62 MB/s | 4881567 | 1736 | 4.0× |
| SonicFastest | 4988470 | 388.99 MB/s | 4882183 | 1736 | 3.9× |
| Easyjson | 5930069 | 327.23 MB/s | 3871266 | 15043 | 3.3× |
| LightningDecodeAny | 7187334 | 269.98 MB/s | 7013525 | 219937 | 2.7× |
| JSONV2 | 8420874 | 230.44 MB/s | 3237189 | 13947 | 2.3× |
| Stdlib | 19499319 | 99.51 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1558526 | 2135.24 MB/s | 4611514 | 5958 | 12.7× |
| SonicFastest | 1701378 | 1955.96 MB/s | 5895392 | 4263 | 11.7× |
| Sonic | 1710921 | 1945.05 MB/s | 5895215 | 4263 | 11.6× |
| LightningDecodeAny | 2891902 | 1062.89 MB/s | 7005136 | 58601 | 6.9× |
| Goccy | 4458045 | 746.48 MB/s | 3948914 | 3816 | 4.4× |
| JSONV2 | 6066680 | 548.54 MB/s | 5364504 | 13243 | 3.3× |
| Stdlib | 19824008 | 167.87 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 195241 | 1128.58 MB/s | 123584 | 225 | 8.8× |
| Goccy | 350032 | 629.50 MB/s | 363946 | 1066 | 4.9× |
| SonicFastest | 406586 | 541.94 MB/s | 350810 | 262 | 4.2× |
| Sonic | 407301 | 540.99 MB/s | 350690 | 262 | 4.2× |
| Easyjson | 434392 | 507.25 MB/s | 130512 | 245 | 4.0× |
| JSONV2 | 527355 | 417.83 MB/s | 129746 | 470 | 3.3× |
| LightningDecodeAny | 738181 | 146.73 MB/s | 861394 | 11944 | 2.3× |
| Stdlib | 1727416 | 127.56 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 10531203 | 769.15 MB/s | 15059843 | 51649 | 7.4× |
| SonicFastest | 16124971 | 502.33 MB/s | 19861254 | 41640 | 4.8× |
| Sonic | 16466089 | 491.92 MB/s | 19861032 | 41640 | 4.7× |
| Goccy | 20410678 | 396.85 MB/s | 19121861 | 107156 | 3.8× |
| Easyjson | 25279138 | 320.42 MB/s | 15059618 | 41643 | 3.1× |
| LightningDecodeAny | 32357129 | 160.80 MB/s | 34333346 | 912810 | 2.4× |
| JSONV2 | 34266652 | 236.38 MB/s | 15233715 | 78972 | 2.3× |
| Stdlib | 77646903 | 104.32 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4733308 | 630.31 MB/s | 3985336 | 30015 | 8.4× |
| Sonic | 7353691 | 405.71 MB/s | 9131002 | 57804 | 5.4× |
| SonicFastest | 7356992 | 405.53 MB/s | 9131736 | 57804 | 5.4× |
| Goccy | 13641396 | 218.71 MB/s | 9944728 | 273623 | 2.9× |
| Easyjson | 14032002 | 212.62 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 14489824 | 126.59 MB/s | 20023836 | 408557 | 2.7× |
| JSONV2 | 21538666 | 138.52 MB/s | 9257070 | 86278 | 1.8× |
| Stdlib | 39582953 | 75.37 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1109761 | 652.03 MB/s | 907387 | 3618 | 8.5× |
| Sonic | 1604178 | 451.07 MB/s | 2372384 | 3683 | 5.9× |
| SonicFastest | 1613394 | 448.49 MB/s | 2373765 | 3683 | 5.9× |
| LightningDecodeAny | 3938798 | 165.17 MB/s | 5752203 | 80172 | 2.4× |
| Easyjson | 4027519 | 179.66 MB/s | 2847907 | 3698 | 2.4× |
| Goccy | 4168093 | 173.60 MB/s | 2724256 | 80268 | 2.3× |
| JSONV2 | 4670786 | 154.92 MB/s | 2704708 | 7318 | 2.0× |
| Stdlib | 9484471 | 76.29 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1607898 | 981.00 MB/s | 907387 | 3618 | 8.6× |
| SonicFastest | 1917555 | 822.59 MB/s | 3226289 | 3683 | 7.2× |
| Sonic | 1981878 | 795.89 MB/s | 3224631 | 3683 | 7.0× |
| LightningDecodeAny | 3322996 | 226.72 MB/s | 5752200 | 80172 | 4.2× |
| Easyjson | 4940096 | 319.30 MB/s | 2847906 | 3698 | 2.8× |
| Goccy | 5075923 | 310.75 MB/s | 3466752 | 80260 | 2.7× |
| JSONV2 | 5231819 | 301.49 MB/s | 2704556 | 7318 | 2.6× |
| Stdlib | 13827338 | 114.07 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 181453 | 827.35 MB/s | 81920 | 1 | 9.7× |
| Sonic | 288370 | 520.60 MB/s | 407311 | 16 | 6.1× |
| SonicFastest | 293420 | 511.64 MB/s | 407618 | 16 | 6.0× |
| LightningDecodeAny | 439302 | 341.73 MB/s | 746003 | 10020 | 4.0× |
| Goccy | 817518 | 183.63 MB/s | 332753 | 10005 | 2.2× |
| JSONV2 | 892149 | 168.27 MB/s | 357726 | 20 | 2.0× |
| Stdlib | 1760265 | 85.28 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 28150 | 998.83 MB/s | 29104 | 107 | 8.6× |
| Sonic | 53895 | 521.70 MB/s | 59470 | 83 | 4.5× |
| SonicFastest | 53940 | 521.26 MB/s | 59466 | 83 | 4.5× |
| Easyjson | 59018 | 476.41 MB/s | 32304 | 138 | 4.1× |
| Goccy | 59624 | 471.57 MB/s | 59281 | 188 | 4.1× |
| JSONV2 | 98327 | 285.95 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 120599 | 233.14 MB/s | 135136 | 2680 | 2.0× |
| Stdlib | 242454 | 115.97 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1581 | 1472.14 MB/s | 32 | 1 | 11.6× |
| Goccy | 3514 | 662.54 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 3986 | 583.99 MB/s | 192 | 2 | 4.6× |
| Sonic | 4757 | 489.43 MB/s | 3710 | 4 | 3.9× |
| SonicFastest | 4759 | 489.19 MB/s | 3709 | 4 | 3.9× |
| JSONV2 | 5810 | 400.65 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 7674 | 219.58 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 18413 | 126.43 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 174 | 1087.68 MB/s | 0 | 0 | 11.1× |
| Goccy | 330 | 572.78 MB/s | 304 | 2 | 5.8× |
| Easyjson | 408 | 462.70 MB/s | 0 | 0 | 4.7× |
| Sonic | 561 | 336.97 MB/s | 341 | 3 | 3.4× |
| SonicFastest | 566 | 333.67 MB/s | 341 | 3 | 3.4× |
| JSONV2 | 742 | 254.56 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 966 | 138.69 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 1926 | 98.13 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1130 | 1938.48 MB/s | 0 | 0 | 11.9× |
| Goccy | 2675 | 818.92 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 2692 | 814.02 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 5199 | 421.41 MB/s | 3601 | 38 | 2.6× |
| Sonic | 5407 | 405.25 MB/s | 3602 | 38 | 2.5× |
| JSONV2 | 6018 | 364.10 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 6814 | 265.78 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 13399 | 163.51 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 678284 | 752.60 MB/s | 703777 | 1012 | 7.4× |
| Goccy | 961503 | 530.91 MB/s | 1135725 | 5006 | 5.2× |
| Sonic | 1142433 | 446.83 MB/s | 1309837 | 2014 | 4.4× |
| SonicFastest | 1145137 | 445.78 MB/s | 1310146 | 2014 | 4.4× |
| Easyjson | 1205310 | 423.52 MB/s | 863777 | 3012 | 4.2× |
| JSONV2 | 2412742 | 211.58 MB/s | 1075961 | 12645 | 2.1× |
| LightningDecodeAny | 2766334 | 166.81 MB/s | 2785928 | 66022 | 1.8× |
| Stdlib | 5010745 | 101.88 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 500 | 39539.75 MB/s | 0 | 0 | 221.1× |
| SonicFastest | 5110 | 3872.48 MB/s | 21121 | 3 | 21.7× |
| Goccy | 20107 | 984.18 MB/s | 20492 | 2 | 5.5× |
| Sonic | 22823 | 867.06 MB/s | 20621 | 3 | 4.8× |
| JSONV2 | 27906 | 709.14 MB/s | 8 | 1 | 4.0× |
| LightningDecodeAny | 73698 | 268.50 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 86236 | 229.47 MB/s | 0 | 0 | 1.3× |
| Stdlib | 110675 | 178.80 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1869 | 9698.74 MB/s | 864 | 4 | 50.3× |
| Easyjson | 3470 | 5223.56 MB/s | 432 | 2 | 27.1× |
| Sonic | 7068 | 2564.23 MB/s | 20448 | 5 | 13.3× |
| SonicFastest | 7880 | 2300.04 MB/s | 20380 | 5 | 11.9× |
| LightningDecodeAny | 14575 | 1226.91 MB/s | 29520 | 193 | 6.5× |
| Goccy | 20920 | 866.36 MB/s | 19460 | 2 | 4.5× |
| JSONV2 | 39299 | 461.18 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 94061 | 192.68 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1847810 | 1086.96 MB/s | 2963982 | 7453 | 8.7× |
| Goccy | 3531899 | 568.67 MB/s | 5410529 | 15830 | 4.6× |
| SonicFastest | 3812415 | 526.83 MB/s | 5155062 | 7085 | 4.2× |
| Sonic | 3842783 | 522.67 MB/s | 5154671 | 7085 | 4.2× |
| Easyjson | 3893223 | 515.89 MB/s | 2981487 | 7439 | 4.1× |
| LightningDecodeAny | 5369442 | 212.74 MB/s | 7387753 | 134757 | 3.0× |
| JSONV2 | 5624355 | 357.11 MB/s | 3173674 | 14563 | 2.9× |
| Stdlib | 16106405 | 124.70 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 652 | 842.41 MB/s | 480 | 1 | 6.9× |
| LightningDecodeAny | 1507 | 363.69 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 1564 | 351.04 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 1651 | 332.43 MB/s | 2261 | 8 | 2.7× |
| Sonic | 1688 | 325.17 MB/s | 2261 | 8 | 2.7× |
| JSONV2 | 2286 | 240.14 MB/s | 1664 | 7 | 2.0× |
| Goccy | 2311 | 237.60 MB/s | 2129 | 43 | 2.0× |
| Stdlib | 4515 | 121.61 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 450225 | 1402.66 MB/s | 461113 | 1230 | 10.3× |
| Sonic | 818345 | 771.70 MB/s | 1064693 | 814 | 5.7× |
| SonicFastest | 820526 | 769.65 MB/s | 1064855 | 814 | 5.6× |
| Goccy | 989896 | 637.96 MB/s | 989390 | 1200 | 4.7× |
| Easyjson | 1010724 | 624.81 MB/s | 422504 | 936 | 4.6× |
| JSONV2 | 1616277 | 390.72 MB/s | 571591 | 3144 | 2.9× |
| LightningDecodeAny | 1857906 | 251.31 MB/s | 2058070 | 30607 | 2.5× |
| Stdlib | 4631808 | 136.34 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 805490 | 698.22 MB/s | 886430 | 2062 | 5.4× |
| Sonic | 1067101 | 527.04 MB/s | 1348274 | 1184 | 4.1× |
| SonicFastest | 1067725 | 526.73 MB/s | 1347911 | 1184 | 4.1× |
| Goccy | 1170671 | 480.42 MB/s | 1040534 | 1028 | 3.7× |
| Easyjson | 1554178 | 361.87 MB/s | 775153 | 1254 | 2.8× |
| JSONV2 | 2346944 | 239.63 MB/s | 927402 | 3482 | 1.9× |
| LightningDecodeAny | 2405378 | 233.81 MB/s | 2232992 | 31103 | 1.8× |
| Stdlib | 4372473 | 128.62 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 601533 | 886.37 MB/s | 413026 | 3155 | 7.7× |
| Easyjson | 1032871 | 516.21 MB/s | 428362 | 3273 | 4.5× |
| SonicFastest | 1064323 | 500.96 MB/s | 982233 | 3082 | 4.4× |
| Sonic | 1064537 | 500.85 MB/s | 981755 | 3082 | 4.4× |
| Goccy | 1109165 | 480.70 MB/s | 1167052 | 5408 | 4.2× |
| JSONV2 | 2084370 | 255.80 MB/s | 745423 | 13288 | 2.2× |
| LightningDecodeAny | 2537174 | 210.15 MB/s | 2709154 | 50805 | 1.8× |
| Stdlib | 4660689 | 114.40 MB/s | 798693 | 17133 | 1.0× |
