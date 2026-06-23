# JSON Deserialization Benchmarks

- generated 2026-06-23T19:53:42Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 104363 | 1219.54 MB/s | 49280 | 2 | 10.6× |
| Lightning | 104498 | 1217.96 MB/s | 49760 | 3 | 10.6× |
| SonicFastest | 182657 | 696.80 MB/s | 194526 | 10 | 6.1× |
| Sonic | 184333 | 690.46 MB/s | 196524 | 10 | 6.0× |
| Goccy | 192973 | 659.55 MB/s | 224552 | 884 | 5.7× |
| Easyjson | 213998 | 594.75 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 421787 | 301.75 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 455045 | 208.01 MB/s | 465585 | 9714 | 2.4× |
| Stdlib | 1105431 | 115.14 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3791139 | 593.77 MB/s | 3122876 | 3081 | 6.9× |
| LightningDestructive | 3792831 | 593.50 MB/s | 3122872 | 3081 | 6.9× |
| SonicFastest | 4545888 | 495.18 MB/s | 15233741 | 970 | 5.7× |
| Sonic | 4547288 | 495.03 MB/s | 15233746 | 970 | 5.7× |
| Goccy | 10249394 | 219.63 MB/s | 4111949 | 56532 | 2.5× |
| Easyjson | 11086125 | 203.05 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11666528 | 192.95 MB/s | 7938298 | 281383 | 2.2× |
| JSONV2 | 16105079 | 139.77 MB/s | 3123206 | 3083 | 1.6× |
| Stdlib | 26121396 | 86.18 MB/s | 3123395 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 490290 | 551.52 MB/s | 348024 | 1627 | 6.8× |
| LightningDestructive | 495345 | 545.89 MB/s | 348024 | 1627 | 6.8× |
| Sonic | 645691 | 418.78 MB/s | 504025 | 968 | 5.2× |
| SonicFastest | 645743 | 418.75 MB/s | 499841 | 968 | 5.2× |
| Easyjson | 1409754 | 191.81 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1412273 | 191.47 MB/s | 543323 | 8122 | 2.4× |
| LightningDecodeAny | 1612463 | 167.70 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2111916 | 128.04 MB/s | 348155 | 1628 | 1.6× |
| Stdlib | 3352806 | 80.65 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1358184 | 1271.70 MB/s | 959848 | 5881 | 9.8× |
| Lightning | 1373049 | 1257.93 MB/s | 959890 | 5882 | 9.7× |
| Sonic | 2035244 | 848.65 MB/s | 2718227 | 4020 | 6.5× |
| SonicFastest | 2049841 | 842.60 MB/s | 2743787 | 4020 | 6.5× |
| Goccy | 2466452 | 700.28 MB/s | 2583709 | 14605 | 5.4× |
| Easyjson | 4204691 | 410.78 MB/s | 972032 | 5389 | 3.2× |
| JSONV2 | 4236982 | 407.65 MB/s | 1011633 | 7594 | 3.1× |
| LightningDecodeAny | 4624226 | 108.19 MB/s | 4976204 | 81466 | 2.9× |
| Stdlib | 13262613 | 130.23 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1193 | 1518.93 MB/s | 0 | 0 | 11.7× |
| LightningDestructive | 1200 | 1510.45 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2540 | 713.40 MB/s | 24 | 1 | 5.5× |
| Goccy | 2814 | 643.92 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6007 | 301.67 MB/s | 3749 | 40 | 2.3× |
| Sonic | 6024 | 300.79 MB/s | 3781 | 40 | 2.3× |
| JSONV2 | 7793 | 232.52 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8265 | 219.12 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14016 | 129.28 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1219 | 1486.82 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1239 | 1462.47 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2539 | 713.80 MB/s | 24 | 1 | 5.5× |
| Goccy | 2875 | 630.24 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 5946 | 304.76 MB/s | 3752 | 40 | 2.4× |
| Sonic | 5972 | 303.43 MB/s | 3759 | 40 | 2.4× |
| JSONV2 | 7791 | 232.58 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8298 | 218.24 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14071 | 128.77 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1400 | 1294.54 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1450 | 1249.77 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2767 | 654.81 MB/s | 144 | 10 | 5.1× |
| Goccy | 2875 | 630.29 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6121 | 296.04 MB/s | 3754 | 42 | 2.3× |
| Sonic | 6131 | 295.53 MB/s | 3754 | 42 | 2.3× |
| JSONV2 | 7950 | 227.93 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8264 | 219.15 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14053 | 128.94 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 738 | 668.94 MB/s | 160 | 1 | 7.5× |
| Lightning | 743 | 664.77 MB/s | 160 | 1 | 7.5× |
| SonicFastest | 1237 | 399.49 MB/s | 971 | 6 | 4.5× |
| Sonic | 1240 | 398.43 MB/s | 981 | 6 | 4.5× |
| LightningDecodeAny | 1657 | 297.55 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2235 | 221.07 MB/s | 448 | 3 | 2.5× |
| Goccy | 2455 | 201.25 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3252 | 151.92 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5541 | 89.16 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 469 | 490.62 MB/s | 160 | 1 | 8.9× |
| Lightning | 469 | 490.20 MB/s | 160 | 1 | 8.9× |
| Sonic | 888 | 258.86 MB/s | 653 | 6 | 4.7× |
| SonicFastest | 889 | 258.84 MB/s | 661 | 6 | 4.7× |
| LightningDecodeAny | 1382 | 165.70 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1414 | 162.63 MB/s | 448 | 3 | 2.9× |
| Goccy | 1591 | 144.52 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2364 | 97.28 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4161 | 55.28 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 74007 | 880.08 MB/s | 166212 | 102 | 7.4× |
| Lightning | 74741 | 871.44 MB/s | 172432 | 107 | 7.3× |
| Sonic | 96929 | 671.95 MB/s | 155059 | 75 | 5.6× |
| SonicFastest | 97316 | 669.29 MB/s | 155475 | 75 | 5.6× |
| Goccy | 145500 | 447.64 MB/s | 228600 | 134 | 3.7× |
| LightningDecodeAny | 187138 | 284.97 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 227897 | 285.80 MB/s | 206653 | 607 | 2.4× |
| Stdlib | 545426 | 119.41 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2611373 | 743.08 MB/s | 2846865 | 2238 | 9.0× |
| Lightning | 2703974 | 717.64 MB/s | 2846867 | 2238 | 8.7× |
| Sonic | 4586143 | 423.12 MB/s | 14606973 | 1407 | 5.1× |
| SonicFastest | 4594022 | 422.39 MB/s | 14606973 | 1407 | 5.1× |
| Goccy | 4793210 | 404.84 MB/s | 4065355 | 13510 | 4.9× |
| Easyjson | 7508935 | 258.42 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 9645250 | 201.18 MB/s | 7013525 | 219937 | 2.4× |
| JSONV2 | 11329822 | 171.27 MB/s | 3237231 | 13947 | 2.1× |
| Stdlib | 23480031 | 82.64 MB/s | 3551325 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1107894 | 3003.74 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1813219 | 1835.32 MB/s | 2488906 | 2995 | 11.5× |
| SonicFastest | 2677407 | 1242.93 MB/s | 6404783 | 4248 | 7.8× |
| Sonic | 2690870 | 1236.71 MB/s | 6480675 | 4248 | 7.8× |
| LightningDecodeAny | 3780107 | 813.14 MB/s | 4886620 | 56892 | 5.5× |
| Goccy | 4702275 | 707.71 MB/s | 3948909 | 3816 | 4.4× |
| JSONV2 | 7405963 | 449.34 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20855533 | 159.57 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221982 | 992.63 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 223954 | 983.89 MB/s | 136897 | 228 | 9.1× |
| Sonic | 379144 | 581.17 MB/s | 300421 | 398 | 5.4× |
| SonicFastest | 381602 | 577.42 MB/s | 304876 | 398 | 5.4× |
| Goccy | 432656 | 509.29 MB/s | 364586 | 1067 | 4.7× |
| Easyjson | 547385 | 402.54 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 734559 | 299.97 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 882256 | 122.77 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2041906 | 107.91 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12833105 | 631.18 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13225577 | 612.45 MB/s | 15059837 | 51649 | 6.8× |
| SonicFastest | 16595653 | 488.08 MB/s | 70887352 | 40014 | 5.4× |
| Sonic | 16799969 | 482.15 MB/s | 70901759 | 40014 | 5.3× |
| Goccy | 23576028 | 343.57 MB/s | 17070229 | 107148 | 3.8× |
| Easyjson | 30861545 | 262.46 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 40369936 | 128.88 MB/s | 34333350 | 912810 | 2.2× |
| JSONV2 | 44360216 | 182.60 MB/s | 15233742 | 78972 | 2.0× |
| Stdlib | 89352714 | 90.65 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6098319 | 489.23 MB/s | 3985336 | 30015 | 7.7× |
| Lightning | 6191943 | 481.83 MB/s | 3985337 | 30015 | 7.6× |
| Sonic | 8626111 | 345.86 MB/s | 26631008 | 56760 | 5.5× |
| SonicFastest | 8657807 | 344.60 MB/s | 26603303 | 56760 | 5.4× |
| Easyjson | 16522182 | 180.57 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16602598 | 179.70 MB/s | 10619884 | 273649 | 2.8× |
| LightningDecodeAny | 19004540 | 96.51 MB/s | 20023839 | 408557 | 2.5× |
| JSONV2 | 24446783 | 122.04 MB/s | 9257139 | 86278 | 1.9× |
| Stdlib | 47177891 | 63.24 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1374987 | 526.26 MB/s | 907393 | 3618 | 8.4× |
| Lightning | 1395490 | 518.53 MB/s | 907389 | 3618 | 8.3× |
| Sonic | 1762813 | 410.48 MB/s | 3189060 | 7226 | 6.6× |
| SonicFastest | 1767482 | 409.39 MB/s | 3186298 | 7226 | 6.6× |
| Easyjson | 4265659 | 169.63 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4361212 | 149.17 MB/s | 5752201 | 80172 | 2.7× |
| Goccy | 4792233 | 150.99 MB/s | 2912132 | 80279 | 2.4× |
| JSONV2 | 5642989 | 128.23 MB/s | 2704641 | 7318 | 2.1× |
| Stdlib | 11579179 | 62.49 MB/s | 2704553 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1951221 | 808.39 MB/s | 907393 | 3618 | 8.1× |
| Lightning | 1973559 | 799.24 MB/s | 907386 | 3618 | 8.0× |
| SonicFastest | 2230879 | 707.05 MB/s | 5785918 | 7226 | 7.1× |
| Sonic | 2232368 | 706.58 MB/s | 5788023 | 7226 | 7.0× |
| LightningDecodeAny | 3855893 | 195.39 MB/s | 5752201 | 80172 | 4.1× |
| Easyjson | 5582178 | 282.57 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5587563 | 282.30 MB/s | 3550578 | 80266 | 2.8× |
| JSONV2 | 6323518 | 249.44 MB/s | 2704596 | 7318 | 2.5× |
| Stdlib | 15728617 | 100.29 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 224058 | 670.02 MB/s | 81920 | 1 | 8.1× |
| Lightning | 224317 | 669.25 MB/s | 81920 | 1 | 8.1× |
| SonicFastest | 275278 | 545.35 MB/s | 258257 | 6 | 6.6× |
| Sonic | 276626 | 542.70 MB/s | 261141 | 6 | 6.6× |
| LightningDecodeAny | 475064 | 316.00 MB/s | 746004 | 10020 | 3.8× |
| Goccy | 863913 | 173.77 MB/s | 325132 | 10004 | 2.1× |
| JSONV2 | 1095825 | 137.00 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1822137 | 82.39 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 34022 | 826.43 MB/s | 30272 | 105 | 8.9× |
| LightningDestructive | 34173 | 822.77 MB/s | 30144 | 103 | 8.9× |
| Sonic | 64220 | 437.82 MB/s | 48369 | 103 | 4.7× |
| SonicFastest | 64516 | 435.82 MB/s | 48693 | 103 | 4.7× |
| Easyjson | 68652 | 409.56 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72720 | 386.65 MB/s | 59278 | 188 | 4.2× |
| JSONV2 | 134950 | 208.35 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 156884 | 179.22 MB/s | 135024 | 2678 | 1.9× |
| Stdlib | 303597 | 92.61 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1954 | 1191.15 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2014 | 1156.12 MB/s | 32 | 1 | 11.3× |
| Goccy | 4130 | 563.73 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4201 | 554.17 MB/s | 192 | 2 | 5.4× |
| Sonic | 5074 | 458.77 MB/s | 4298 | 6 | 4.5× |
| SonicFastest | 5077 | 458.56 MB/s | 4285 | 6 | 4.5× |
| JSONV2 | 8454 | 275.37 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10289 | 163.77 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22674 | 102.67 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220 | 857.17 MB/s | 0 | 0 | 11.2× |
| LightningDestructive | 222 | 850.53 MB/s | 0 | 0 | 11.1× |
| Goccy | 384 | 492.21 MB/s | 304 | 2 | 6.4× |
| Easyjson | 495 | 381.54 MB/s | 0 | 0 | 5.0× |
| Sonic | 796 | 237.55 MB/s | 494 | 4 | 3.1× |
| SonicFastest | 797 | 237.26 MB/s | 496 | 4 | 3.1× |
| JSONV2 | 1029 | 183.69 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1243 | 107.84 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2466 | 76.63 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1541 | 1421.52 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1569 | 1396.23 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3197 | 685.23 MB/s | 24 | 1 | 5.0× |
| Goccy | 3251 | 673.91 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6408 | 341.89 MB/s | 4029 | 40 | 2.5× |
| Sonic | 6440 | 340.19 MB/s | 4085 | 40 | 2.5× |
| JSONV2 | 7996 | 274.00 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8561 | 211.55 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16086 | 136.21 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 699914 | 729.34 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 732186 | 697.19 MB/s | 703780 | 1012 | 8.3× |
| Goccy | 1152671 | 442.86 MB/s | 1143233 | 5007 | 5.2× |
| SonicFastest | 1165972 | 437.81 MB/s | 901051 | 2006 | 5.2× |
| Sonic | 1171017 | 435.93 MB/s | 905038 | 2006 | 5.2× |
| Easyjson | 1531183 | 333.39 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3243620 | 157.38 MB/s | 1076021 | 12646 | 1.9× |
| LightningDecodeAny | 3540980 | 130.32 MB/s | 2785926 | 66022 | 1.7× |
| Stdlib | 6044934 | 84.45 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1352 | 14634.99 MB/s | 0 | 0 | 96.6× |
| LightningDestructive | 1392 | 14218.64 MB/s | 0 | 0 | 93.8× |
| Goccy | 20497 | 965.44 MB/s | 20491 | 2 | 6.4× |
| Sonic | 27889 | 709.56 MB/s | 22044 | 4 | 4.7× |
| SonicFastest | 28489 | 694.62 MB/s | 23473 | 4 | 4.6× |
| JSONV2 | 29892 | 662.02 MB/s | 8 | 1 | 4.4× |
| LightningDecodeAny | 79452 | 249.06 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 86049 | 229.97 MB/s | 0 | 0 | 1.5× |
| Stdlib | 130547 | 151.58 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2660 | 6812.79 MB/s | 0 | 0 | 38.8× |
| Lightning | 2823 | 6420.31 MB/s | 432 | 2 | 36.6× |
| Easyjson | 3958 | 4579.09 MB/s | 432 | 2 | 26.1× |
| Sonic | 9793 | 1850.63 MB/s | 22877 | 6 | 10.5× |
| SonicFastest | 9885 | 1833.57 MB/s | 22900 | 6 | 10.4× |
| Goccy | 15905 | 1139.53 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 16872 | 1059.86 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 44779 | 404.74 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103292 | 175.46 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2262937 | 887.56 MB/s | 2960388 | 7411 | 8.2× |
| Lightning | 2309425 | 869.69 MB/s | 2962102 | 7417 | 8.1× |
| Goccy | 4249147 | 472.68 MB/s | 5411819 | 15830 | 4.4× |
| Sonic | 4457918 | 450.55 MB/s | 10953101 | 13683 | 4.2× |
| SonicFastest | 4464070 | 449.92 MB/s | 10917387 | 13683 | 4.2× |
| Easyjson | 4941406 | 406.46 MB/s | 2981482 | 7439 | 3.8× |
| JSONV2 | 6970467 | 288.14 MB/s | 3173680 | 14563 | 2.7× |
| LightningDecodeAny | 7385198 | 154.67 MB/s | 7386073 | 134751 | 2.5× |
| Stdlib | 18656716 | 107.66 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 890 | 616.82 MB/s | 480 | 1 | 6.3× |
| Lightning | 894 | 613.96 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1914 | 286.37 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2161 | 254.09 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2701 | 203.25 MB/s | 1948 | 26 | 2.1× |
| SonicFastest | 2708 | 202.71 MB/s | 1950 | 26 | 2.1× |
| Goccy | 3087 | 177.86 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3323 | 165.20 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5639 | 97.35 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 500871 | 1260.83 MB/s | 364472 | 566 | 10.8× |
| Lightning | 552212 | 1143.61 MB/s | 413001 | 878 | 9.8× |
| SonicFastest | 1025910 | 615.56 MB/s | 1004436 | 1102 | 5.3× |
| Sonic | 1030591 | 612.77 MB/s | 1016608 | 1102 | 5.2× |
| Easyjson | 1150272 | 549.01 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1170067 | 539.72 MB/s | 986487 | 1201 | 4.6× |
| JSONV2 | 2155768 | 292.94 MB/s | 571617 | 3144 | 2.5× |
| LightningDecodeAny | 2406231 | 194.04 MB/s | 2010201 | 30295 | 2.2× |
| Stdlib | 5394014 | 117.08 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 710615 | 791.44 MB/s | 544254 | 448 | 7.4× |
| Lightning | 882419 | 637.35 MB/s | 767621 | 1254 | 6.0× |
| Sonic | 1023034 | 549.75 MB/s | 936291 | 1476 | 5.1× |
| SonicFastest | 1041247 | 540.13 MB/s | 956647 | 1476 | 5.1× |
| Goccy | 1351344 | 416.18 MB/s | 1042886 | 1030 | 3.9× |
| Easyjson | 1736158 | 323.94 MB/s | 775153 | 1254 | 3.0× |
| LightningDecodeAny | 2721287 | 206.67 MB/s | 2114151 | 30295 | 1.9× |
| JSONV2 | 2780724 | 202.25 MB/s | 927440 | 3482 | 1.9× |
| Stdlib | 5268229 | 106.75 MB/s | 1011670 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 662533 | 804.76 MB/s | 333416 | 2084 | 8.3× |
| Lightning | 700126 | 761.55 MB/s | 368224 | 2293 | 7.9× |
| Easyjson | 1132660 | 470.73 MB/s | 428362 | 3273 | 4.9× |
| SonicFastest | 1148729 | 464.15 MB/s | 1028377 | 4351 | 4.8× |
| Sonic | 1154052 | 462.01 MB/s | 1040336 | 4351 | 4.8× |
| Goccy | 1300661 | 409.93 MB/s | 1167229 | 5409 | 4.2× |
| JSONV2 | 2575506 | 207.02 MB/s | 745471 | 13288 | 2.1× |
| LightningDecodeAny | 3441130 | 154.94 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5500293 | 96.94 MB/s | 798692 | 17133 | 1.0× |
