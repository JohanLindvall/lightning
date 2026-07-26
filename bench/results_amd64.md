# JSON Deserialization Benchmarks

- generated 2026-07-26T15:08:55Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 78132 | 1628.98 MB/s | 49760 | 3 | 12.9× |
| LightningDestructive | 78308 | 1625.32 MB/s | 49280 | 2 | 12.9× |
| Sonic | 146856 | 866.67 MB/s | 213832 | 15 | 6.9× |
| SonicFastest | 148280 | 858.34 MB/s | 213781 | 15 | 6.8× |
| Easyjson | 183745 | 692.67 MB/s | 122864 | 14 | 5.5× |
| Goccy | 186175 | 683.63 MB/s | 225690 | 884 | 5.4× |
| JSONV2 | 302684 | 420.49 MB/s | 195127 | 1805 | 3.3× |
| LightningDecodeAny | 327586 | 288.94 MB/s | 465730 | 9708 | 3.1× |
| Stdlib | 1011168 | 125.87 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3121568 | 721.13 MB/s | 2532848 | 1143 | 7.8× |
| Lightning | 3178531 | 708.20 MB/s | 2532849 | 1143 | 7.6× |
| SonicFastest | 3926546 | 573.29 MB/s | 4871457 | 2584 | 6.2× |
| Sonic | 4102595 | 548.69 MB/s | 4869143 | 2584 | 5.9× |
| Goccy | 9979643 | 225.56 MB/s | 4151506 | 56533 | 2.4× |
| LightningDecodeAny | 9990903 | 225.31 MB/s | 19380211 | 223896 | 2.4× |
| Easyjson | 10583405 | 212.70 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 13142347 | 171.28 MB/s | 3123193 | 3083 | 1.8× |
| Stdlib | 24241869 | 92.86 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 416902 | 648.60 MB/s | 397296 | 567 | 7.6× |
| LightningDestructive | 422581 | 639.89 MB/s | 397297 | 567 | 7.5× |
| SonicFastest | 572227 | 472.55 MB/s | 641705 | 1147 | 5.5× |
| Sonic | 574870 | 470.37 MB/s | 641472 | 1147 | 5.5× |
| Easyjson | 1349377 | 200.39 MB/s | 330272 | 749 | 2.3× |
| Goccy | 1359801 | 198.85 MB/s | 541882 | 8122 | 2.3× |
| LightningDecodeAny | 1620236 | 166.89 MB/s | 2543877 | 29687 | 1.9× |
| JSONV2 | 1713428 | 157.81 MB/s | 348159 | 1628 | 1.8× |
| Stdlib | 3148332 | 85.89 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 848501 | 2035.59 MB/s | 765560 | 2798 | 15.4× |
| Lightning | 870086 | 1985.10 MB/s | 765601 | 2799 | 15.0× |
| SonicFastest | 1684656 | 1025.26 MB/s | 2696630 | 5547 | 7.8× |
| Sonic | 1688131 | 1023.15 MB/s | 2695633 | 5547 | 7.7× |
| Goccy | 2193031 | 787.59 MB/s | 2580896 | 14603 | 6.0× |
| Easyjson | 2866602 | 602.53 MB/s | 972032 | 5389 | 4.6× |
| LightningDecodeAny | 2955029 | 169.30 MB/s | 4954731 | 76576 | 4.4× |
| JSONV2 | 3102114 | 556.78 MB/s | 1011613 | 7594 | 4.2× |
| Stdlib | 13068436 | 132.17 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 783 | 2314.53 MB/s | 0 | 0 | 15.7× |
| LightningDestructive | 792 | 2287.40 MB/s | 0 | 0 | 15.6× |
| Easyjson | 2196 | 825.29 MB/s | 24 | 1 | 5.6× |
| Goccy | 2400 | 755.02 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 4760 | 380.70 MB/s | 3350 | 38 | 2.6× |
| Sonic | 4939 | 366.85 MB/s | 3347 | 38 | 2.5× |
| JSONV2 | 5821 | 311.29 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 6738 | 268.76 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 12328 | 146.98 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 824 | 2200.08 MB/s | 0 | 0 | 15.0× |
| LightningDestructive | 845 | 2143.60 MB/s | 0 | 0 | 14.6× |
| Easyjson | 2246 | 806.70 MB/s | 24 | 1 | 5.5× |
| Goccy | 2407 | 752.79 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 4914 | 368.71 MB/s | 3349 | 38 | 2.5× |
| Sonic | 5088 | 356.11 MB/s | 3348 | 38 | 2.4× |
| JSONV2 | 5820 | 311.34 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 6522 | 277.69 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 12338 | 146.86 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 957 | 1893.12 MB/s | 144 | 10 | 12.8× |
| LightningDestructive | 993 | 1824.72 MB/s | 144 | 10 | 12.4× |
| Easyjson | 2315 | 782.69 MB/s | 144 | 10 | 5.3× |
| Goccy | 2559 | 708.06 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 4975 | 364.24 MB/s | 3367 | 40 | 2.5× |
| Sonic | 5109 | 354.68 MB/s | 3369 | 40 | 2.4× |
| JSONV2 | 6291 | 288.04 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 6572 | 275.55 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 12268 | 147.70 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 538 | 919.13 MB/s | 160 | 1 | 8.8× |
| LightningDestructive | 547 | 903.17 MB/s | 160 | 1 | 8.7× |
| Sonic | 940 | 525.44 MB/s | 1075 | 8 | 5.0× |
| SonicFastest | 944 | 523.08 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1077 | 457.61 MB/s | 1296 | 26 | 4.4× |
| Easyjson | 1866 | 264.79 MB/s | 448 | 3 | 2.5× |
| Goccy | 1979 | 249.60 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 2324 | 212.57 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4733 | 104.37 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 340 | 677.33 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 342 | 673.28 MB/s | 160 | 1 | 9.9× |
| SonicFastest | 668 | 344.28 MB/s | 802 | 8 | 5.0× |
| Sonic | 671 | 342.62 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 929 | 246.56 MB/s | 1296 | 26 | 3.6× |
| Easyjson | 1199 | 191.81 MB/s | 448 | 3 | 2.8× |
| Goccy | 1319 | 174.34 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 1808 | 127.19 MB/s | 528 | 7 | 1.9× |
| Stdlib | 3369 | 68.27 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 63547 | 1024.95 MB/s | 164881 | 105 | 8.3× |
| LightningDestructive | 65768 | 990.33 MB/s | 158661 | 100 | 8.0× |
| Sonic | 115187 | 565.44 MB/s | 235838 | 65 | 4.6× |
| SonicFastest | 116056 | 561.21 MB/s | 235787 | 65 | 4.5× |
| Goccy | 141315 | 460.90 MB/s | 228470 | 134 | 3.7× |
| LightningDecodeAny | 146185 | 364.81 MB/s | 180224 | 3245 | 3.6× |
| JSONV2 | 195497 | 333.16 MB/s | 206664 | 607 | 2.7× |
| Stdlib | 525555 | 123.93 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1975548 | 982.24 MB/s | 2864593 | 1380 | 10.3× |
| Lightning | 1996851 | 971.77 MB/s | 2864594 | 1380 | 10.2× |
| Goccy | 3780541 | 513.28 MB/s | 4063181 | 13509 | 5.4× |
| SonicFastest | 4967766 | 390.61 MB/s | 4878693 | 1736 | 4.1× |
| Sonic | 4978068 | 389.80 MB/s | 4880611 | 1736 | 4.1× |
| Easyjson | 6019124 | 322.38 MB/s | 3871266 | 15043 | 3.4× |
| LightningDecodeAny | 7284247 | 266.39 MB/s | 7064789 | 218633 | 2.8× |
| JSONV2 | 9000942 | 215.59 MB/s | 3237188 | 13947 | 2.3× |
| Stdlib | 20408959 | 95.08 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 756219 | 4400.62 MB/s | 351704 | 1286 | 26.4× |
| Lightning | 1174367 | 2833.72 MB/s | 2488906 | 2995 | 17.0× |
| Sonic | 1687977 | 1971.49 MB/s | 5895642 | 4263 | 11.8× |
| SonicFastest | 1701031 | 1956.36 MB/s | 5894467 | 4263 | 11.8× |
| LightningDecodeAny | 2564594 | 1198.54 MB/s | 4886621 | 56892 | 7.8× |
| Goccy | 5120932 | 649.85 MB/s | 3948913 | 3816 | 3.9× |
| JSONV2 | 6708724 | 496.05 MB/s | 5364503 | 13243 | 3.0× |
| Stdlib | 19988165 | 166.49 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 165473 | 1331.61 MB/s | 135872 | 226 | 10.8× |
| LightningDestructive | 168961 | 1304.12 MB/s | 135872 | 226 | 10.5× |
| Goccy | 361650 | 609.28 MB/s | 363702 | 1066 | 4.9× |
| Sonic | 404450 | 544.80 MB/s | 351237 | 262 | 4.4× |
| SonicFastest | 407405 | 540.85 MB/s | 351545 | 262 | 4.4× |
| Easyjson | 455150 | 484.12 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 524707 | 419.94 MB/s | 129746 | 470 | 3.4× |
| LightningDecodeAny | 731986 | 147.97 MB/s | 897522 | 11703 | 2.4× |
| Stdlib | 1781976 | 123.65 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9071802 | 892.88 MB/s | 11845072 | 20816 | 8.6× |
| Lightning | 9282902 | 872.58 MB/s | 11845074 | 20816 | 8.4× |
| Sonic | 15848037 | 511.11 MB/s | 19864512 | 41640 | 4.9× |
| SonicFastest | 16010068 | 505.93 MB/s | 19862947 | 41640 | 4.9× |
| Goccy | 20439211 | 396.30 MB/s | 19150608 | 107156 | 3.8× |
| Easyjson | 25591371 | 316.51 MB/s | 15059620 | 41643 | 3.0× |
| LightningDecodeAny | 31228347 | 166.61 MB/s | 46191120 | 747112 | 2.5× |
| JSONV2 | 35841341 | 226.00 MB/s | 15233732 | 78972 | 2.2× |
| Stdlib | 77988041 | 103.86 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4502715 | 662.59 MB/s | 3758856 | 29356 | 9.0× |
| Lightning | 4540260 | 657.11 MB/s | 3758857 | 29356 | 9.0× |
| Sonic | 7134125 | 418.20 MB/s | 9130013 | 57804 | 5.7× |
| SonicFastest | 7137266 | 418.01 MB/s | 9131539 | 57804 | 5.7× |
| Goccy | 13826580 | 215.78 MB/s | 9890351 | 273620 | 2.9× |
| Easyjson | 13882113 | 214.91 MB/s | 9479441 | 30115 | 2.9× |
| LightningDecodeAny | 14087367 | 130.20 MB/s | 23982399 | 351152 | 2.9× |
| JSONV2 | 19341631 | 154.25 MB/s | 9257040 | 86278 | 2.1× |
| Stdlib | 40641625 | 73.41 MB/s | 9258083 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1036441 | 698.16 MB/s | 907600 | 3618 | 9.6× |
| Lightning | 1056978 | 684.59 MB/s | 907595 | 3618 | 9.4× |
| Sonic | 1605396 | 450.73 MB/s | 2373182 | 3683 | 6.2× |
| SonicFastest | 1620992 | 446.39 MB/s | 2371995 | 3683 | 6.2× |
| LightningDecodeAny | 3893015 | 167.11 MB/s | 6500458 | 76546 | 2.6× |
| Easyjson | 4046211 | 178.83 MB/s | 2847906 | 3698 | 2.5× |
| Goccy | 4136235 | 174.94 MB/s | 2733863 | 80268 | 2.4× |
| JSONV2 | 4726446 | 153.10 MB/s | 2704704 | 7318 | 2.1× |
| Stdlib | 9980098 | 72.50 MB/s | 2704547 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1485046 | 1062.16 MB/s | 907594 | 3618 | 9.3× |
| LightningDestructive | 1507641 | 1046.24 MB/s | 907600 | 3618 | 9.2× |
| SonicFastest | 1903843 | 828.51 MB/s | 3228131 | 3683 | 7.3× |
| Sonic | 1919223 | 821.87 MB/s | 3228425 | 3683 | 7.2× |
| LightningDecodeAny | 3616480 | 208.32 MB/s | 6500456 | 76546 | 3.8× |
| Easyjson | 5046464 | 312.57 MB/s | 2847906 | 3698 | 2.8× |
| Goccy | 5194436 | 303.66 MB/s | 3495394 | 80262 | 2.7× |
| JSONV2 | 5237219 | 301.18 MB/s | 2704552 | 7318 | 2.7× |
| Stdlib | 13882117 | 113.62 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 169811 | 884.06 MB/s | 81920 | 1 | 9.3× |
| LightningDestructive | 177562 | 845.47 MB/s | 81920 | 1 | 8.9× |
| Sonic | 305975 | 490.64 MB/s | 407917 | 16 | 5.2× |
| SonicFastest | 323285 | 464.37 MB/s | 409762 | 16 | 4.9× |
| LightningDecodeAny | 425143 | 353.11 MB/s | 745764 | 10016 | 3.7× |
| Goccy | 789298 | 190.20 MB/s | 329596 | 10005 | 2.0× |
| JSONV2 | 902824 | 166.28 MB/s | 357725 | 20 | 1.8× |
| Stdlib | 1583265 | 94.82 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 24040 | 1169.58 MB/s | 29216 | 103 | 10.7× |
| LightningDestructive | 24535 | 1145.98 MB/s | 29088 | 101 | 10.5× |
| SonicFastest | 52188 | 538.77 MB/s | 59441 | 83 | 4.9× |
| Sonic | 52297 | 537.64 MB/s | 59435 | 83 | 4.9× |
| Easyjson | 59002 | 476.54 MB/s | 32304 | 138 | 4.4× |
| Goccy | 59550 | 472.16 MB/s | 59257 | 188 | 4.3× |
| JSONV2 | 97027 | 289.79 MB/s | 36897 | 242 | 2.7× |
| LightningDecodeAny | 118133 | 238.01 MB/s | 140592 | 2643 | 2.2× |
| Stdlib | 257217 | 109.31 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1479 | 1573.60 MB/s | 32 | 1 | 13.6× |
| LightningDestructive | 1523 | 1528.12 MB/s | 32 | 1 | 13.2× |
| Goccy | 3515 | 662.31 MB/s | 3649 | 4 | 5.7× |
| Easyjson | 4018 | 579.40 MB/s | 192 | 2 | 5.0× |
| Sonic | 4762 | 488.90 MB/s | 3707 | 4 | 4.2× |
| SonicFastest | 4801 | 484.90 MB/s | 3712 | 4 | 4.2× |
| JSONV2 | 6052 | 384.65 MB/s | 1000 | 6 | 3.3× |
| LightningDecodeAny | 7620 | 221.13 MB/s | 10200 | 195 | 2.6× |
| Stdlib | 20085 | 115.91 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 162 | 1166.36 MB/s | 0 | 0 | 12.6× |
| LightningDestructive | 166 | 1140.01 MB/s | 0 | 0 | 12.4× |
| Goccy | 306 | 616.75 MB/s | 304 | 2 | 6.7× |
| Easyjson | 431 | 438.43 MB/s | 0 | 0 | 4.8× |
| Sonic | 558 | 338.64 MB/s | 341 | 3 | 3.7× |
| SonicFastest | 559 | 338.20 MB/s | 341 | 3 | 3.7× |
| JSONV2 | 719 | 262.88 MB/s | 112 | 1 | 2.8× |
| LightningDecodeAny | 921 | 145.50 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 2049 | 92.26 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1081 | 2027.15 MB/s | 0 | 0 | 13.2× |
| LightningDestructive | 1117 | 1961.74 MB/s | 0 | 0 | 12.8× |
| Goccy | 2662 | 823.03 MB/s | 2864 | 4 | 5.4× |
| Easyjson | 2763 | 792.96 MB/s | 24 | 1 | 5.2× |
| SonicFastest | 5152 | 425.28 MB/s | 3600 | 38 | 2.8× |
| Sonic | 5350 | 409.57 MB/s | 3603 | 38 | 2.7× |
| JSONV2 | 5932 | 369.37 MB/s | 640 | 6 | 2.4× |
| LightningDecodeAny | 6634 | 272.97 MB/s | 7536 | 158 | 2.2× |
| Stdlib | 14294 | 153.28 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 454422 | 1123.35 MB/s | 457537 | 1009 | 11.6× |
| Lightning | 460838 | 1107.71 MB/s | 457537 | 1009 | 11.4× |
| Goccy | 986253 | 517.59 MB/s | 1136733 | 5006 | 5.3× |
| Sonic | 1156009 | 441.58 MB/s | 1308162 | 2014 | 4.6× |
| SonicFastest | 1156648 | 441.34 MB/s | 1307960 | 2014 | 4.6× |
| Easyjson | 1167144 | 437.37 MB/s | 863778 | 3012 | 4.5× |
| JSONV2 | 2433384 | 209.78 MB/s | 1075960 | 12645 | 2.2× |
| LightningDecodeAny | 2600770 | 177.43 MB/s | 2929688 | 64018 | 2.0× |
| Stdlib | 5267137 | 96.92 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 474 | 41732.80 MB/s | 0 | 0 | 255.7× |
| LightningDestructive | 653 | 30317.51 MB/s | 0 | 0 | 185.7× |
| SonicFastest | 4782 | 4138.06 MB/s | 21126 | 3 | 25.4× |
| Goccy | 19891 | 994.89 MB/s | 20492 | 2 | 6.1× |
| Sonic | 22634 | 874.29 MB/s | 20617 | 3 | 5.4× |
| JSONV2 | 27867 | 710.13 MB/s | 8 | 1 | 4.4× |
| LightningDecodeAny | 71849 | 275.41 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 92735 | 213.39 MB/s | 0 | 0 | 1.3× |
| Stdlib | 121233 | 163.23 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1667 | 10873.88 MB/s | 432 | 2 | 61.9× |
| LightningDestructive | 1695 | 10690.14 MB/s | 0 | 0 | 60.9× |
| Easyjson | 3654 | 4960.53 MB/s | 432 | 2 | 28.2× |
| SonicFastest | 7101 | 2552.17 MB/s | 20469 | 5 | 14.5× |
| Sonic | 7529 | 2407.15 MB/s | 20430 | 5 | 13.7× |
| LightningDecodeAny | 14411 | 1240.89 MB/s | 29088 | 191 | 7.2× |
| Goccy | 17537 | 1033.48 MB/s | 19460 | 2 | 5.9× |
| JSONV2 | 37158 | 487.75 MB/s | 16500 | 50 | 2.8× |
| Stdlib | 103143 | 175.72 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1745087 | 1150.94 MB/s | 3089565 | 6821 | 9.3× |
| Lightning | 1856142 | 1082.08 MB/s | 3091278 | 6827 | 8.8× |
| Goccy | 3727815 | 538.79 MB/s | 5410325 | 15830 | 4.4× |
| SonicFastest | 3841824 | 522.80 MB/s | 5156471 | 7085 | 4.2× |
| Sonic | 3882478 | 517.32 MB/s | 5156457 | 7085 | 4.2× |
| Easyjson | 4235946 | 474.15 MB/s | 2981482 | 7439 | 3.9× |
| JSONV2 | 5385534 | 372.94 MB/s | 3173681 | 14563 | 3.0× |
| LightningDecodeAny | 5429517 | 210.39 MB/s | 8498329 | 134008 | 3.0× |
| Stdlib | 16312229 | 123.13 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 630 | 871.35 MB/s | 480 | 1 | 7.7× |
| LightningDestructive | 639 | 858.90 MB/s | 480 | 1 | 7.5× |
| LightningDecodeAny | 1304 | 420.19 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 1498 | 366.45 MB/s | 1616 | 5 | 3.2× |
| SonicFastest | 1619 | 339.10 MB/s | 2263 | 8 | 3.0× |
| Sonic | 1655 | 331.63 MB/s | 2263 | 8 | 2.9× |
| Goccy | 2255 | 243.43 MB/s | 2129 | 43 | 2.1× |
| JSONV2 | 2374 | 231.23 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 4823 | 113.83 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 366159 | 1724.70 MB/s | 402728 | 545 | 13.1× |
| Lightning | 426324 | 1481.30 MB/s | 451257 | 857 | 11.3× |
| Sonic | 845593 | 746.83 MB/s | 1066944 | 814 | 5.7× |
| SonicFastest | 889945 | 709.61 MB/s | 1069863 | 814 | 5.4× |
| Easyjson | 992483 | 636.30 MB/s | 422504 | 936 | 4.8× |
| Goccy | 1006377 | 627.51 MB/s | 991678 | 1201 | 4.8× |
| JSONV2 | 1700406 | 371.39 MB/s | 571590 | 3144 | 2.8× |
| LightningDecodeAny | 1902980 | 245.36 MB/s | 2077365 | 30126 | 2.5× |
| Stdlib | 4801731 | 131.52 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 527994 | 1065.18 MB/s | 546569 | 429 | 8.6× |
| Lightning | 683756 | 822.53 MB/s | 769941 | 1235 | 6.7× |
| Sonic | 1074297 | 523.51 MB/s | 1348283 | 1184 | 4.2× |
| SonicFastest | 1083370 | 519.13 MB/s | 1348452 | 1184 | 4.2× |
| Goccy | 1184437 | 474.83 MB/s | 1037944 | 1028 | 3.8× |
| Easyjson | 1547249 | 363.49 MB/s | 775154 | 1254 | 2.9× |
| LightningDecodeAny | 2218148 | 253.55 MB/s | 2181319 | 30126 | 2.1× |
| JSONV2 | 2343745 | 239.96 MB/s | 927404 | 3482 | 1.9× |
| Stdlib | 4552351 | 123.54 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 457793 | 1164.67 MB/s | 333416 | 2084 | 10.8× |
| Lightning | 515592 | 1034.11 MB/s | 368224 | 2293 | 9.6× |
| Easyjson | 989566 | 538.80 MB/s | 428362 | 3273 | 5.0× |
| SonicFastest | 1057400 | 504.23 MB/s | 980497 | 3082 | 4.7× |
| Sonic | 1058617 | 503.66 MB/s | 980817 | 3082 | 4.7× |
| Goccy | 1119303 | 476.35 MB/s | 1167088 | 5409 | 4.4× |
| JSONV2 | 2081755 | 256.12 MB/s | 745420 | 13288 | 2.4× |
| LightningDecodeAny | 2658135 | 200.58 MB/s | 2991146 | 50076 | 1.9× |
| Stdlib | 4957477 | 107.55 MB/s | 798693 | 17133 | 1.0× |
