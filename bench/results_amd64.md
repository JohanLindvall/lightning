# JSON Deserialization Benchmarks

- generated 2026-08-11T11:20:10Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 104528 | 1217.62 MB/s | 49760 | 3 | 12.9× |
| Lightning | 104883 | 1213.50 MB/s | 49760 | 3 | 12.8× |
| LightningDestructive | 105555 | 1205.77 MB/s | 49280 | 2 | 12.7× |
| Sonic | 206183 | 617.29 MB/s | 214335 | 15 | 6.5× |
| SonicFastest | 206207 | 617.22 MB/s | 214302 | 15 | 6.5× |
| Goccy | 240390 | 529.45 MB/s | 224755 | 884 | 5.6× |
| Easyjson | 251141 | 506.79 MB/s | 122864 | 14 | 5.3× |
| LightningDecodeAny | 450536 | 210.09 MB/s | 463411 | 9708 | 3.0× |
| JSONV2 | 467493 | 272.25 MB/s | 195127 | 1805 | 2.9× |
| Stdlib | 1343521 | 94.73 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4012816 | 560.97 MB/s | 2532848 | 1143 | 8.4× |
| LightningArena | 4062721 | 554.07 MB/s | 2532849 | 1143 | 8.3× |
| Lightning | 4125690 | 545.62 MB/s | 2532851 | 1143 | 8.2× |
| SonicFastest | 5364018 | 419.66 MB/s | 4867128 | 2584 | 6.3× |
| Sonic | 5428279 | 414.69 MB/s | 4866869 | 2584 | 6.2× |
| Goccy | 13163489 | 171.01 MB/s | 4162939 | 56533 | 2.6× |
| LightningDecodeAny | 13553920 | 166.08 MB/s | 19380210 | 223896 | 2.5× |
| Easyjson | 13816438 | 162.93 MB/s | 3099811 | 2120 | 2.5× |
| JSONV2 | 16715394 | 134.67 MB/s | 3123200 | 3083 | 2.0× |
| Stdlib | 33877798 | 66.45 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 555688 | 486.61 MB/s | 397297 | 567 | 7.9× |
| Lightning | 558605 | 484.07 MB/s | 397297 | 567 | 7.8× |
| LightningDestructive | 565422 | 478.23 MB/s | 397296 | 567 | 7.7× |
| SonicFastest | 774104 | 349.31 MB/s | 641083 | 1147 | 5.6× |
| Sonic | 774779 | 349.01 MB/s | 641406 | 1147 | 5.6× |
| Goccy | 1831818 | 147.61 MB/s | 541265 | 8122 | 2.4× |
| Easyjson | 1832256 | 147.58 MB/s | 330273 | 749 | 2.4× |
| JSONV2 | 2253223 | 120.01 MB/s | 348160 | 1628 | 1.9× |
| LightningDecodeAny | 2285359 | 118.32 MB/s | 2543877 | 29687 | 1.9× |
| Stdlib | 4367442 | 61.91 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1134753 | 1522.10 MB/s | 765560 | 2798 | 15.0× |
| LightningArena | 1167783 | 1479.05 MB/s | 768416 | 2440 | 14.6× |
| Lightning | 1175350 | 1469.52 MB/s | 765602 | 2799 | 14.5× |
| Sonic | 2205238 | 783.23 MB/s | 2693671 | 5547 | 7.7× |
| SonicFastest | 2217545 | 778.88 MB/s | 2693209 | 5547 | 7.7× |
| Goccy | 2579788 | 669.51 MB/s | 2581192 | 14603 | 6.6× |
| Easyjson | 3978892 | 434.09 MB/s | 972032 | 5389 | 4.3× |
| LightningDecodeAny | 4014345 | 124.63 MB/s | 4953692 | 76576 | 4.3× |
| JSONV2 | 4787410 | 360.78 MB/s | 1011615 | 7594 | 3.6× |
| Stdlib | 17064202 | 101.22 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 996 | 1819.08 MB/s | 0 | 0 | 16.5× |
| Lightning | 1008 | 1796.97 MB/s | 0 | 0 | 16.3× |
| LightningDestructive | 1073 | 1689.33 MB/s | 0 | 0 | 15.3× |
| Easyjson | 2951 | 614.01 MB/s | 24 | 1 | 5.6× |
| Goccy | 3268 | 554.45 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6145 | 294.89 MB/s | 3344 | 38 | 2.7× |
| Sonic | 6410 | 282.68 MB/s | 3346 | 38 | 2.6× |
| JSONV2 | 8238 | 219.96 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9124 | 198.49 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16452 | 110.14 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1103 | 1642.68 MB/s | 0 | 0 | 15.1× |
| LightningArena | 1106 | 1638.14 MB/s | 0 | 0 | 15.0× |
| LightningDestructive | 1169 | 1549.81 MB/s | 0 | 0 | 14.2× |
| Easyjson | 3042 | 595.62 MB/s | 24 | 1 | 5.5× |
| Goccy | 3273 | 553.56 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 6171 | 293.63 MB/s | 3346 | 38 | 2.7× |
| Sonic | 6405 | 282.90 MB/s | 3345 | 38 | 2.6× |
| JSONV2 | 8258 | 219.42 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9263 | 195.50 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16606 | 109.12 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1281 | 1414.00 MB/s | 144 | 10 | 13.0× |
| LightningArena | 1291 | 1403.65 MB/s | 144 | 10 | 12.9× |
| LightningDestructive | 1358 | 1334.51 MB/s | 144 | 10 | 12.3× |
| Easyjson | 3170 | 571.67 MB/s | 144 | 10 | 5.3× |
| Goccy | 3574 | 507.06 MB/s | 2600 | 5 | 4.7× |
| SonicFastest | 6614 | 273.98 MB/s | 3366 | 40 | 2.5× |
| Sonic | 6825 | 265.49 MB/s | 3364 | 40 | 2.4× |
| JSONV2 | 8694 | 208.42 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9095 | 199.12 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16699 | 108.51 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 712 | 693.99 MB/s | 160 | 1 | 9.3× |
| LightningDestructive | 725 | 681.75 MB/s | 160 | 1 | 9.2× |
| Sonic | 1282 | 385.32 MB/s | 1077 | 8 | 5.2× |
| SonicFastest | 1284 | 384.72 MB/s | 1076 | 8 | 5.2× |
| LightningDecodeAny | 1552 | 317.57 MB/s | 1296 | 26 | 4.3× |
| LightningArena | 1734 | 284.82 MB/s | 4096 | 1 | 3.8× |
| Goccy | 2575 | 191.84 MB/s | 856 | 23 | 2.6× |
| Easyjson | 2670 | 185.00 MB/s | 448 | 3 | 2.5× |
| JSONV2 | 3422 | 144.35 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6638 | 74.42 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 449 | 512.11 MB/s | 160 | 1 | 10.6× |
| LightningDestructive | 456 | 503.84 MB/s | 160 | 1 | 10.4× |
| SonicFastest | 948 | 242.55 MB/s | 801 | 8 | 5.0× |
| Sonic | 953 | 241.40 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1304 | 175.55 MB/s | 1296 | 26 | 3.6× |
| LightningArena | 1363 | 168.72 MB/s | 4096 | 1 | 3.5× |
| Easyjson | 1665 | 138.14 MB/s | 448 | 3 | 2.8× |
| Goccy | 1784 | 128.90 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2651 | 86.77 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4742 | 48.50 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 71032 | 916.94 MB/s | 103441 | 103 | 9.9× |
| LightningArena | 74721 | 871.67 MB/s | 103440 | 103 | 9.4× |
| LightningDestructive | 84570 | 770.15 MB/s | 97220 | 98 | 8.3× |
| Sonic | 151140 | 430.94 MB/s | 236049 | 65 | 4.6× |
| SonicFastest | 151975 | 428.57 MB/s | 236169 | 65 | 4.6× |
| Goccy | 178669 | 364.54 MB/s | 228002 | 134 | 3.9× |
| LightningDecodeAny | 208128 | 256.23 MB/s | 180049 | 3245 | 3.4× |
| JSONV2 | 264332 | 246.40 MB/s | 206665 | 607 | 2.7× |
| Stdlib | 702621 | 92.70 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2524708 | 768.59 MB/s | 2864593 | 1380 | 10.9× |
| LightningArena | 2593144 | 748.31 MB/s | 2864593 | 1380 | 10.6× |
| Lightning | 2616672 | 741.58 MB/s | 2864595 | 1380 | 10.5× |
| SonicFastest | 4915570 | 394.76 MB/s | 4878945 | 1736 | 5.6× |
| Sonic | 5006877 | 387.56 MB/s | 4879336 | 1736 | 5.5× |
| Goccy | 5175470 | 374.94 MB/s | 4062622 | 13509 | 5.3× |
| Easyjson | 8116099 | 239.09 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 9494704 | 204.37 MB/s | 7063040 | 218633 | 2.9× |
| JSONV2 | 12780801 | 151.83 MB/s | 3237191 | 13947 | 2.2× |
| Stdlib | 27569575 | 70.38 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1058345 | 3144.37 MB/s | 351704 | 1286 | 22.8× |
| LightningArena | 1551662 | 2144.69 MB/s | 2488907 | 2995 | 15.6× |
| Lightning | 1564161 | 2127.55 MB/s | 2488907 | 2995 | 15.5× |
| Sonic | 2014504 | 1651.94 MB/s | 5896559 | 4263 | 12.0× |
| SonicFastest | 2017599 | 1649.40 MB/s | 5896447 | 4263 | 12.0× |
| LightningDecodeAny | 3581778 | 858.17 MB/s | 4876913 | 56892 | 6.7× |
| Goccy | 5017923 | 663.19 MB/s | 3948912 | 3816 | 4.8× |
| JSONV2 | 7823140 | 425.38 MB/s | 5364499 | 13243 | 3.1× |
| Stdlib | 24168578 | 137.69 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 217492 | 1013.12 MB/s | 135872 | 226 | 11.0× |
| LightningArena | 217551 | 1012.85 MB/s | 135872 | 226 | 11.0× |
| LightningDestructive | 230240 | 957.03 MB/s | 135872 | 226 | 10.4× |
| SonicFastest | 483741 | 455.50 MB/s | 351032 | 262 | 4.9× |
| Sonic | 484110 | 455.16 MB/s | 351069 | 262 | 4.9× |
| Goccy | 493717 | 446.30 MB/s | 363508 | 1066 | 4.8× |
| Easyjson | 650692 | 338.63 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 777081 | 283.56 MB/s | 129746 | 470 | 3.1× |
| LightningDecodeAny | 1025745 | 105.59 MB/s | 897217 | 11703 | 2.3× |
| Stdlib | 2392623 | 92.09 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12527992 | 646.56 MB/s | 11845073 | 20816 | 8.7× |
| Lightning | 12920041 | 626.94 MB/s | 11845073 | 20816 | 8.5× |
| LightningArena | 13158507 | 615.57 MB/s | 11845073 | 20816 | 8.3× |
| Sonic | 18315077 | 442.26 MB/s | 19856861 | 41640 | 6.0× |
| SonicFastest | 18427552 | 439.56 MB/s | 19856570 | 41640 | 5.9× |
| Goccy | 26877837 | 301.36 MB/s | 18963876 | 107154 | 4.1× |
| Easyjson | 35231364 | 229.91 MB/s | 15059617 | 41643 | 3.1× |
| LightningDecodeAny | 39998999 | 130.08 MB/s | 46279352 | 747112 | 2.7× |
| JSONV2 | 51099457 | 158.52 MB/s | 15233748 | 78972 | 2.1× |
| Stdlib | 109605467 | 73.90 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5492581 | 543.18 MB/s | 3764713 | 1504 | 10.3× |
| LightningDestructive | 5846503 | 510.30 MB/s | 3758856 | 29356 | 9.7× |
| Lightning | 6095883 | 489.42 MB/s | 3758857 | 29356 | 9.3× |
| Sonic | 9136637 | 326.54 MB/s | 9131932 | 57804 | 6.2× |
| SonicFastest | 9188496 | 324.70 MB/s | 9131081 | 57804 | 6.2× |
| LightningDecodeAny | 18617483 | 98.52 MB/s | 23982578 | 351152 | 3.0× |
| Goccy | 19009683 | 156.94 MB/s | 9749181 | 273615 | 3.0× |
| Easyjson | 19430623 | 153.54 MB/s | 9479440 | 30115 | 2.9× |
| JSONV2 | 27027492 | 110.39 MB/s | 9257048 | 86278 | 2.1× |
| Stdlib | 56741014 | 52.58 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1323442 | 546.75 MB/s | 907600 | 3618 | 10.5× |
| LightningArena | 1331565 | 543.42 MB/s | 911393 | 30 | 10.4× |
| Lightning | 1414798 | 511.45 MB/s | 907595 | 3618 | 9.8× |
| Sonic | 2112828 | 342.48 MB/s | 2371644 | 3683 | 6.6× |
| SonicFastest | 2121356 | 341.10 MB/s | 2372001 | 3683 | 6.5× |
| Easyjson | 5314853 | 136.15 MB/s | 2847908 | 3698 | 2.6× |
| Goccy | 5451636 | 132.73 MB/s | 2727288 | 80268 | 2.5× |
| LightningDecodeAny | 5470628 | 118.92 MB/s | 6500461 | 76546 | 2.5× |
| JSONV2 | 6712409 | 107.80 MB/s | 2704709 | 7318 | 2.1× |
| Stdlib | 13874128 | 52.15 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1908572 | 826.46 MB/s | 911392 | 30 | 10.1× |
| Lightning | 1979669 | 796.78 MB/s | 907595 | 3618 | 9.7× |
| LightningDestructive | 1981502 | 796.04 MB/s | 907600 | 3618 | 9.7× |
| Sonic | 2469373 | 638.77 MB/s | 3224952 | 3683 | 7.8× |
| SonicFastest | 2546154 | 619.50 MB/s | 3232663 | 3683 | 7.6× |
| LightningDecodeAny | 4753185 | 158.50 MB/s | 6500455 | 76546 | 4.1× |
| Easyjson | 6560245 | 240.44 MB/s | 2847906 | 3698 | 2.9× |
| Goccy | 6717911 | 234.80 MB/s | 3487205 | 80262 | 2.9× |
| JSONV2 | 7431056 | 212.26 MB/s | 2704552 | 7318 | 2.6× |
| Stdlib | 19274018 | 81.84 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 232188 | 646.56 MB/s | 81920 | 1 | 9.2× |
| Lightning | 233958 | 641.67 MB/s | 81920 | 1 | 9.1× |
| LightningDestructive | 243488 | 616.56 MB/s | 81920 | 1 | 8.8× |
| Sonic | 419562 | 357.81 MB/s | 408963 | 16 | 5.1× |
| SonicFastest | 443728 | 338.32 MB/s | 409320 | 16 | 4.8× |
| LightningDecodeAny | 586006 | 256.18 MB/s | 745766 | 10016 | 3.6× |
| Goccy | 1057812 | 141.92 MB/s | 328326 | 10005 | 2.0× |
| JSONV2 | 1205809 | 124.50 MB/s | 357723 | 20 | 1.8× |
| Stdlib | 2138450 | 70.20 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 32994 | 852.19 MB/s | 29216 | 103 | 10.6× |
| Lightning | 33044 | 850.89 MB/s | 29216 | 103 | 10.6× |
| LightningDestructive | 34866 | 806.44 MB/s | 29088 | 101 | 10.1× |
| Sonic | 57304 | 490.67 MB/s | 59459 | 83 | 6.1× |
| SonicFastest | 57593 | 488.20 MB/s | 59475 | 83 | 6.1× |
| Goccy | 81070 | 346.82 MB/s | 59288 | 188 | 4.3× |
| Easyjson | 81376 | 345.52 MB/s | 32304 | 138 | 4.3× |
| JSONV2 | 140797 | 199.70 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 164356 | 171.07 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 350553 | 80.21 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1992 | 1168.92 MB/s | 32 | 1 | 13.3× |
| Lightning | 2002 | 1163.09 MB/s | 32 | 1 | 13.3× |
| LightningDestructive | 2152 | 1081.94 MB/s | 32 | 1 | 12.3× |
| SonicFastest | 4865 | 478.56 MB/s | 3714 | 4 | 5.5× |
| Sonic | 4917 | 473.45 MB/s | 3714 | 4 | 5.4× |
| Easyjson | 5051 | 460.89 MB/s | 192 | 2 | 5.3× |
| Goccy | 5116 | 455.01 MB/s | 3649 | 4 | 5.2× |
| JSONV2 | 8602 | 270.63 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10757 | 156.64 MB/s | 10200 | 195 | 2.5× |
| Stdlib | 26545 | 87.70 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 203 | 932.77 MB/s | 0 | 0 | 13.9× |
| LightningArena | 203 | 930.20 MB/s | 0 | 0 | 13.9× |
| LightningDestructive | 210 | 900.97 MB/s | 0 | 0 | 13.4× |
| Goccy | 457 | 413.55 MB/s | 304 | 2 | 6.2× |
| Easyjson | 551 | 343.07 MB/s | 0 | 0 | 5.1× |
| SonicFastest | 650 | 290.71 MB/s | 341 | 3 | 4.3× |
| Sonic | 651 | 290.25 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1056 | 179.03 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1353 | 99.01 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2816 | 67.11 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1451 | 1509.85 MB/s | 0 | 0 | 13.3× |
| LightningArena | 1471 | 1489.13 MB/s | 0 | 0 | 13.1× |
| LightningDestructive | 1509 | 1451.70 MB/s | 0 | 0 | 12.7× |
| Goccy | 3739 | 585.91 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3855 | 568.32 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6853 | 319.70 MB/s | 3606 | 38 | 2.8× |
| Sonic | 7148 | 306.54 MB/s | 3603 | 38 | 2.7× |
| JSONV2 | 8568 | 255.72 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9298 | 194.76 MB/s | 7552 | 158 | 2.1× |
| Stdlib | 19237 | 113.90 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 593089 | 860.71 MB/s | 457537 | 1009 | 12.0× |
| LightningArena | 619221 | 824.38 MB/s | 457537 | 1009 | 11.5× |
| Lightning | 629096 | 811.44 MB/s | 457537 | 1009 | 11.3× |
| Sonic | 1261321 | 404.72 MB/s | 1308555 | 2014 | 5.6× |
| Goccy | 1263780 | 403.93 MB/s | 1136189 | 5006 | 5.6× |
| SonicFastest | 1287698 | 396.43 MB/s | 1309672 | 2014 | 5.5× |
| Easyjson | 1816866 | 280.97 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3579821 | 142.60 MB/s | 1075968 | 12645 | 2.0× |
| LightningDecodeAny | 3635254 | 126.94 MB/s | 2950650 | 64018 | 2.0× |
| Stdlib | 7113791 | 71.76 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 546 | 36227.44 MB/s | 0 | 0 | 303.5× |
| LightningArena | 549 | 36043.48 MB/s | 0 | 0 | 301.9× |
| LightningDestructive | 858 | 23076.69 MB/s | 0 | 0 | 193.3× |
| SonicFastest | 6808 | 2906.57 MB/s | 21112 | 3 | 24.3× |
| Goccy | 23642 | 837.02 MB/s | 20492 | 2 | 7.0× |
| Sonic | 32295 | 612.75 MB/s | 20619 | 3 | 5.1× |
| JSONV2 | 33914 | 583.50 MB/s | 8 | 1 | 4.9× |
| LightningDecodeAny | 95752 | 206.66 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 106826 | 185.24 MB/s | 0 | 0 | 1.6× |
| Stdlib | 165769 | 119.38 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2210 | 8199.12 MB/s | 432 | 2 | 57.5× |
| LightningArena | 2224 | 8149.82 MB/s | 432 | 2 | 57.2× |
| LightningDestructive | 2395 | 7567.79 MB/s | 0 | 0 | 53.1× |
| Easyjson | 4791 | 3782.72 MB/s | 432 | 2 | 26.5× |
| Sonic | 8779 | 2064.37 MB/s | 20405 | 5 | 14.5× |
| SonicFastest | 9195 | 1971.02 MB/s | 20373 | 5 | 13.8× |
| LightningDecodeAny | 19389 | 922.29 MB/s | 29088 | 191 | 6.6× |
| Goccy | 25857 | 700.94 MB/s | 19460 | 2 | 4.9× |
| JSONV2 | 51906 | 349.17 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 127158 | 142.53 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2434297 | 825.08 MB/s | 3089565 | 6821 | 9.0× |
| Lightning | 2536137 | 791.95 MB/s | 3091278 | 6827 | 8.7× |
| LightningArena | 2541220 | 790.37 MB/s | 3094370 | 6703 | 8.7× |
| SonicFastest | 3877544 | 517.98 MB/s | 5151790 | 7085 | 5.7× |
| Sonic | 3949274 | 508.57 MB/s | 5152230 | 7085 | 5.6× |
| Goccy | 4664771 | 430.57 MB/s | 5410485 | 15832 | 4.7× |
| Easyjson | 5412417 | 371.09 MB/s | 2981479 | 7439 | 4.1× |
| LightningDecodeAny | 7047505 | 162.08 MB/s | 8503512 | 134008 | 3.1× |
| JSONV2 | 7815849 | 256.98 MB/s | 3173669 | 14562 | 2.8× |
| Stdlib | 22007384 | 91.26 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 857 | 640.54 MB/s | 480 | 1 | 7.9× |
| LightningArena | 858 | 640.10 MB/s | 480 | 1 | 7.9× |
| LightningDestructive | 882 | 622.37 MB/s | 480 | 1 | 7.7× |
| LightningDecodeAny | 1823 | 300.66 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 2288 | 239.99 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2427 | 226.22 MB/s | 2262 | 8 | 2.8× |
| Sonic | 2437 | 225.28 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3397 | 161.61 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3662 | 149.91 MB/s | 1664 | 7 | 1.8× |
| Stdlib | 6763 | 81.18 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 506642 | 1246.47 MB/s | 402728 | 545 | 12.8× |
| LightningArena | 577056 | 1094.37 MB/s | 453017 | 712 | 11.2× |
| Lightning | 588336 | 1073.39 MB/s | 451257 | 857 | 11.0× |
| SonicFastest | 999777 | 631.65 MB/s | 1065109 | 814 | 6.5× |
| Sonic | 1000363 | 631.28 MB/s | 1065297 | 814 | 6.5× |
| Easyjson | 1323158 | 477.28 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1426371 | 442.74 MB/s | 988125 | 1200 | 4.5× |
| JSONV2 | 2487806 | 253.84 MB/s | 571589 | 3144 | 2.6× |
| LightningDecodeAny | 2539426 | 183.86 MB/s | 2076504 | 30126 | 2.5× |
| Stdlib | 6470465 | 97.60 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 681679 | 825.03 MB/s | 546569 | 429 | 9.1× |
| Lightning | 874600 | 643.05 MB/s | 769937 | 1235 | 7.1× |
| LightningArena | 879085 | 639.77 MB/s | 771665 | 1088 | 7.1× |
| SonicFastest | 1350192 | 416.54 MB/s | 1349096 | 1185 | 4.6× |
| Sonic | 1352242 | 415.91 MB/s | 1348555 | 1185 | 4.6× |
| Goccy | 1658674 | 339.07 MB/s | 1037922 | 1028 | 3.8× |
| Easyjson | 2193315 | 256.42 MB/s | 775153 | 1254 | 2.8× |
| LightningDecodeAny | 3050781 | 184.35 MB/s | 2180441 | 30126 | 2.0× |
| JSONV2 | 3354524 | 167.66 MB/s | 927402 | 3482 | 1.9× |
| Stdlib | 6236317 | 90.18 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 629134 | 847.48 MB/s | 333416 | 2084 | 10.4× |
| LightningArena | 709711 | 751.26 MB/s | 368224 | 2293 | 9.2× |
| Lightning | 722720 | 737.74 MB/s | 368224 | 2293 | 9.1× |
| Sonic | 1139404 | 467.94 MB/s | 982246 | 3082 | 5.7× |
| SonicFastest | 1157142 | 460.77 MB/s | 982630 | 3082 | 5.7× |
| Easyjson | 1371456 | 388.77 MB/s | 428362 | 3273 | 4.8× |
| Goccy | 1466535 | 363.56 MB/s | 1167106 | 5409 | 4.5× |
| JSONV2 | 2991440 | 178.23 MB/s | 745419 | 13288 | 2.2× |
| LightningDecodeAny | 3573672 | 149.20 MB/s | 2992876 | 50076 | 1.8× |
| Stdlib | 6549388 | 81.41 MB/s | 798693 | 17133 | 1.0× |
