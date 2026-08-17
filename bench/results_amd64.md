# JSON Deserialization Benchmarks

- generated 2026-08-17T08:49:25Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 84581 | 1504.76 MB/s | 49760 | 3 | 15.8× |
| Lightning | 85089 | 1495.79 MB/s | 49760 | 3 | 15.7× |
| LightningDestructive | 89926 | 1415.33 MB/s | 49280 | 2 | 14.9× |
| SonicFastest | 196257 | 648.51 MB/s | 214324 | 15 | 6.8× |
| Sonic | 196486 | 647.76 MB/s | 214236 | 15 | 6.8× |
| Goccy | 244130 | 521.34 MB/s | 225009 | 884 | 5.5× |
| Easyjson | 246609 | 516.10 MB/s | 122864 | 14 | 5.4× |
| LightningDecodeAny | 445545 | 212.44 MB/s | 463411 | 9708 | 3.0× |
| JSONV2 | 452368 | 281.35 MB/s | 195129 | 1805 | 3.0× |
| Stdlib | 1337513 | 95.16 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4018683 | 560.15 MB/s | 2532849 | 1143 | 8.6× |
| Lightning | 4052807 | 555.43 MB/s | 2532852 | 1143 | 8.5× |
| LightningArena | 4085628 | 550.97 MB/s | 2532849 | 1143 | 8.5× |
| SonicFastest | 5545873 | 405.90 MB/s | 4879980 | 2584 | 6.2× |
| Sonic | 5546776 | 405.83 MB/s | 4876480 | 2584 | 6.2× |
| Goccy | 12519293 | 179.81 MB/s | 4241109 | 56538 | 2.8× |
| LightningDecodeAny | 13517578 | 166.53 MB/s | 19380210 | 223896 | 2.6× |
| Easyjson | 14535009 | 154.87 MB/s | 3099809 | 2120 | 2.4× |
| JSONV2 | 18650297 | 120.70 MB/s | 3123190 | 3083 | 1.9× |
| Stdlib | 34621210 | 65.02 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 548588 | 492.91 MB/s | 397297 | 567 | 7.9× |
| Lightning | 550874 | 490.86 MB/s | 397297 | 567 | 7.9× |
| LightningDestructive | 560177 | 482.71 MB/s | 397296 | 567 | 7.8× |
| Sonic | 747226 | 361.88 MB/s | 641202 | 1147 | 5.8× |
| SonicFastest | 768214 | 351.99 MB/s | 643110 | 1147 | 5.7× |
| Goccy | 1698255 | 159.22 MB/s | 544952 | 8123 | 2.6× |
| Easyjson | 1789913 | 151.07 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 2178810 | 124.11 MB/s | 2543877 | 29687 | 2.0× |
| JSONV2 | 2516028 | 107.47 MB/s | 348160 | 1628 | 1.7× |
| Stdlib | 4357002 | 62.06 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1023670 | 1687.27 MB/s | 765560 | 2798 | 16.5× |
| LightningArena | 1070383 | 1613.63 MB/s | 768416 | 2440 | 15.8× |
| Lightning | 1091931 | 1581.79 MB/s | 765602 | 2799 | 15.5× |
| Sonic | 2158506 | 800.18 MB/s | 2693683 | 5547 | 7.8× |
| SonicFastest | 2165434 | 797.62 MB/s | 2693116 | 5547 | 7.8× |
| Goccy | 2932746 | 588.94 MB/s | 2581506 | 14603 | 5.8× |
| LightningDecodeAny | 3999189 | 125.10 MB/s | 4953692 | 76576 | 4.2× |
| Easyjson | 4147686 | 416.43 MB/s | 972032 | 5389 | 4.1× |
| JSONV2 | 4955335 | 348.55 MB/s | 1011615 | 7594 | 3.4× |
| Stdlib | 16888876 | 102.27 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 816 | 2220.56 MB/s | 0 | 0 | 20.1× |
| LightningArena | 819 | 2211.71 MB/s | 0 | 0 | 20.1× |
| LightningDestructive | 910 | 1992.27 MB/s | 0 | 0 | 18.1× |
| Easyjson | 2861 | 633.45 MB/s | 24 | 1 | 5.7× |
| Goccy | 3505 | 516.92 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6311 | 287.12 MB/s | 3346 | 38 | 2.6× |
| Sonic | 6465 | 280.27 MB/s | 3343 | 38 | 2.5× |
| JSONV2 | 8456 | 214.29 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9228 | 196.25 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16433 | 110.27 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 892 | 2032.03 MB/s | 0 | 0 | 18.7× |
| LightningArena | 902 | 2008.21 MB/s | 0 | 0 | 18.5× |
| LightningDestructive | 953 | 1900.93 MB/s | 0 | 0 | 17.5× |
| Easyjson | 2905 | 623.72 MB/s | 24 | 1 | 5.8× |
| Goccy | 3519 | 514.96 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6158 | 294.23 MB/s | 3346 | 38 | 2.7× |
| Sonic | 6385 | 283.79 MB/s | 3346 | 38 | 2.6× |
| JSONV2 | 8661 | 209.20 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 9138 | 198.18 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16718 | 108.39 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1128 | 1607.04 MB/s | 144 | 10 | 14.7× |
| LightningArena | 1137 | 1594.22 MB/s | 144 | 10 | 14.6× |
| LightningDestructive | 1226 | 1478.53 MB/s | 144 | 10 | 13.6× |
| Goccy | 3239 | 559.42 MB/s | 2600 | 5 | 5.1× |
| Easyjson | 3268 | 554.51 MB/s | 144 | 10 | 5.1× |
| SonicFastest | 6596 | 274.73 MB/s | 3368 | 40 | 2.5× |
| Sonic | 6786 | 267.03 MB/s | 3369 | 40 | 2.5× |
| JSONV2 | 8352 | 216.97 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 9117 | 198.64 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16637 | 108.91 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 705 | 701.10 MB/s | 160 | 1 | 9.4× |
| LightningDestructive | 718 | 688.48 MB/s | 160 | 1 | 9.2× |
| Sonic | 1274 | 387.76 MB/s | 1076 | 8 | 5.2× |
| SonicFastest | 1285 | 384.37 MB/s | 1076 | 8 | 5.1× |
| LightningDecodeAny | 1550 | 317.99 MB/s | 1296 | 26 | 4.3× |
| LightningArena | 1646 | 300.20 MB/s | 4096 | 1 | 4.0× |
| Goccy | 2695 | 183.31 MB/s | 856 | 23 | 2.4× |
| Easyjson | 2714 | 182.01 MB/s | 448 | 3 | 2.4× |
| JSONV2 | 3235 | 152.68 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6595 | 74.90 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 455 | 505.11 MB/s | 160 | 1 | 10.5× |
| LightningDestructive | 460 | 500.57 MB/s | 160 | 1 | 10.4× |
| Sonic | 962 | 238.97 MB/s | 800 | 8 | 4.9× |
| SonicFastest | 974 | 236.06 MB/s | 801 | 8 | 4.9× |
| LightningDecodeAny | 1348 | 169.83 MB/s | 1296 | 26 | 3.5× |
| LightningArena | 1352 | 170.12 MB/s | 4096 | 1 | 3.5× |
| Easyjson | 1694 | 135.75 MB/s | 448 | 3 | 2.8× |
| Goccy | 1850 | 124.31 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2579 | 89.20 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4764 | 48.28 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 63347 | 1028.18 MB/s | 103441 | 103 | 10.7× |
| Lightning | 64166 | 1015.05 MB/s | 103441 | 103 | 10.6× |
| LightningDestructive | 67386 | 966.55 MB/s | 97220 | 98 | 10.1× |
| Sonic | 144911 | 449.46 MB/s | 235812 | 65 | 4.7× |
| SonicFastest | 145082 | 448.93 MB/s | 235812 | 65 | 4.7× |
| Goccy | 186086 | 350.01 MB/s | 227939 | 134 | 3.6× |
| LightningDecodeAny | 201806 | 264.26 MB/s | 180048 | 3245 | 3.4× |
| JSONV2 | 270275 | 240.98 MB/s | 206665 | 607 | 2.5× |
| Stdlib | 677984 | 96.07 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2410270 | 805.09 MB/s | 2864593 | 1380 | 11.5× |
| Lightning | 2466374 | 786.77 MB/s | 2864593 | 1380 | 11.2× |
| LightningArena | 2503515 | 775.10 MB/s | 2864594 | 1380 | 11.1× |
| SonicFastest | 4922548 | 394.20 MB/s | 4879634 | 1736 | 5.6× |
| Sonic | 4935907 | 393.13 MB/s | 4879158 | 1736 | 5.6× |
| Goccy | 5070261 | 382.72 MB/s | 4063731 | 13509 | 5.5× |
| Easyjson | 8562301 | 226.63 MB/s | 3871264 | 15043 | 3.2× |
| LightningDecodeAny | 9730009 | 199.43 MB/s | 7063039 | 218633 | 2.8× |
| JSONV2 | 13101432 | 148.11 MB/s | 3237197 | 13947 | 2.1× |
| Stdlib | 27713765 | 70.02 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 954960 | 3484.79 MB/s | 351704 | 1286 | 24.8× |
| Lightning | 1483955 | 2242.54 MB/s | 2488907 | 2995 | 16.0× |
| LightningArena | 1485909 | 2239.59 MB/s | 2488907 | 2995 | 16.0× |
| Sonic | 2011891 | 1654.08 MB/s | 5896462 | 4263 | 11.8× |
| SonicFastest | 2024731 | 1643.59 MB/s | 5896272 | 4263 | 11.7× |
| LightningDecodeAny | 3536088 | 869.26 MB/s | 4876913 | 56892 | 6.7× |
| Goccy | 6123070 | 543.49 MB/s | 3948915 | 3817 | 3.9× |
| JSONV2 | 7874290 | 422.62 MB/s | 5364506 | 13243 | 3.0× |
| Stdlib | 23705094 | 140.38 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 198892 | 1107.87 MB/s | 135872 | 226 | 12.0× |
| LightningArena | 199591 | 1103.99 MB/s | 135872 | 226 | 12.0× |
| LightningDestructive | 208503 | 1056.80 MB/s | 135872 | 226 | 11.4× |
| Goccy | 483938 | 455.32 MB/s | 363981 | 1066 | 4.9× |
| Sonic | 490228 | 449.48 MB/s | 350639 | 262 | 4.9× |
| SonicFastest | 490244 | 449.46 MB/s | 350704 | 262 | 4.9× |
| Easyjson | 630382 | 349.54 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 819909 | 268.74 MB/s | 129746 | 470 | 2.9× |
| LightningDecodeAny | 1006297 | 107.64 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2385788 | 92.36 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11521164 | 703.06 MB/s | 11845073 | 20816 | 9.6× |
| LightningArena | 11753476 | 689.16 MB/s | 11845073 | 20816 | 9.4× |
| Lightning | 11821855 | 685.17 MB/s | 11845073 | 20816 | 9.3× |
| SonicFastest | 18513781 | 437.51 MB/s | 19860516 | 41640 | 5.9× |
| Sonic | 18696064 | 433.25 MB/s | 19857956 | 41640 | 5.9× |
| Goccy | 26092340 | 310.44 MB/s | 19123221 | 107156 | 4.2× |
| Easyjson | 36982545 | 219.02 MB/s | 15059619 | 41643 | 3.0× |
| LightningDecodeAny | 40631723 | 128.05 MB/s | 46279350 | 747112 | 2.7× |
| JSONV2 | 52548894 | 154.14 MB/s | 15233803 | 78973 | 2.1× |
| Stdlib | 110056917 | 73.60 MB/s | 15665075 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5607927 | 532.01 MB/s | 3764713 | 1504 | 10.4× |
| LightningDestructive | 5782773 | 515.92 MB/s | 3758856 | 29356 | 10.1× |
| Lightning | 6080206 | 490.69 MB/s | 3758860 | 29356 | 9.6× |
| SonicFastest | 9514953 | 313.56 MB/s | 9132023 | 57804 | 6.1× |
| Sonic | 9538992 | 312.77 MB/s | 9131875 | 57804 | 6.1× |
| LightningDecodeAny | 19016363 | 96.45 MB/s | 23982580 | 351152 | 3.1× |
| Goccy | 19325983 | 154.38 MB/s | 9899037 | 273622 | 3.0× |
| Easyjson | 20531855 | 145.31 MB/s | 9479440 | 30115 | 2.8× |
| JSONV2 | 27636427 | 107.95 MB/s | 9257072 | 86278 | 2.1× |
| Stdlib | 58186128 | 51.27 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1336766 | 541.30 MB/s | 907600 | 3618 | 10.5× |
| LightningArena | 1348804 | 536.47 MB/s | 911393 | 30 | 10.4× |
| Lightning | 1406346 | 514.52 MB/s | 907595 | 3618 | 9.9× |
| Sonic | 2131569 | 339.47 MB/s | 2371701 | 3683 | 6.6× |
| SonicFastest | 2143836 | 337.52 MB/s | 2371528 | 3683 | 6.5× |
| Easyjson | 5471718 | 132.24 MB/s | 2847906 | 3698 | 2.6× |
| Goccy | 5508400 | 131.36 MB/s | 2699290 | 80267 | 2.5× |
| LightningDecodeAny | 5512073 | 118.03 MB/s | 6500460 | 76546 | 2.5× |
| JSONV2 | 6835741 | 105.85 MB/s | 2704706 | 7318 | 2.0× |
| Stdlib | 13986197 | 51.74 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1806846 | 872.99 MB/s | 911393 | 30 | 10.8× |
| Lightning | 1845026 | 854.92 MB/s | 907595 | 3618 | 10.6× |
| LightningDestructive | 1865608 | 845.49 MB/s | 907600 | 3618 | 10.5× |
| Sonic | 2440387 | 646.35 MB/s | 3223016 | 3683 | 8.0× |
| SonicFastest | 2447905 | 644.37 MB/s | 3224268 | 3683 | 8.0× |
| LightningDecodeAny | 4741239 | 158.90 MB/s | 6500454 | 76546 | 4.1× |
| Easyjson | 6517191 | 242.03 MB/s | 2847906 | 3698 | 3.0× |
| Goccy | 6671743 | 236.42 MB/s | 3483657 | 80261 | 2.9× |
| JSONV2 | 7511886 | 209.98 MB/s | 2704553 | 7318 | 2.6× |
| Stdlib | 19507354 | 80.86 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 229585 | 653.89 MB/s | 81920 | 1 | 9.2× |
| Lightning | 230700 | 650.73 MB/s | 81920 | 1 | 9.2× |
| LightningDestructive | 238861 | 628.50 MB/s | 81920 | 1 | 8.9× |
| SonicFastest | 379481 | 395.60 MB/s | 407455 | 16 | 5.6× |
| Sonic | 384714 | 390.22 MB/s | 407447 | 16 | 5.5× |
| LightningDecodeAny | 600054 | 250.18 MB/s | 745765 | 10016 | 3.5× |
| Goccy | 974218 | 154.10 MB/s | 324682 | 10005 | 2.2× |
| JSONV2 | 1175460 | 127.72 MB/s | 357725 | 20 | 1.8× |
| Stdlib | 2114767 | 70.99 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 30592 | 919.10 MB/s | 29216 | 103 | 11.5× |
| LightningArena | 30660 | 917.07 MB/s | 29216 | 103 | 11.4× |
| LightningDestructive | 31674 | 887.69 MB/s | 29088 | 101 | 11.1× |
| Sonic | 60027 | 468.40 MB/s | 59497 | 83 | 5.8× |
| SonicFastest | 60270 | 466.52 MB/s | 59513 | 83 | 5.8× |
| Easyjson | 79464 | 353.83 MB/s | 32304 | 138 | 4.4× |
| Goccy | 82006 | 342.87 MB/s | 59257 | 188 | 4.3× |
| JSONV2 | 142675 | 197.07 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 165427 | 169.97 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 350636 | 80.19 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1791 | 1300.16 MB/s | 32 | 1 | 14.7× |
| LightningArena | 1800 | 1293.26 MB/s | 32 | 1 | 14.6× |
| LightningDestructive | 1894 | 1229.11 MB/s | 32 | 1 | 13.9× |
| Sonic | 4798 | 485.21 MB/s | 3705 | 4 | 5.5× |
| SonicFastest | 4821 | 482.85 MB/s | 3707 | 4 | 5.4× |
| Goccy | 4845 | 480.48 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 5496 | 423.58 MB/s | 192 | 2 | 4.8× |
| JSONV2 | 8777 | 265.23 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 10526 | 160.09 MB/s | 10200 | 195 | 2.5× |
| Stdlib | 26248 | 88.69 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 177 | 1069.96 MB/s | 0 | 0 | 15.8× |
| Lightning | 178 | 1060.06 MB/s | 0 | 0 | 15.7× |
| LightningDestructive | 188 | 1007.46 MB/s | 0 | 0 | 14.9× |
| Goccy | 416 | 453.83 MB/s | 304 | 2 | 6.7× |
| Easyjson | 592 | 319.34 MB/s | 0 | 0 | 4.7× |
| Sonic | 645 | 292.99 MB/s | 341 | 3 | 4.3× |
| SonicFastest | 647 | 292.03 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1028 | 183.85 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1344 | 99.73 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2791 | 67.73 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1146 | 1911.43 MB/s | 0 | 0 | 16.7× |
| LightningArena | 1168 | 1875.22 MB/s | 0 | 0 | 16.4× |
| LightningDestructive | 1249 | 1753.88 MB/s | 0 | 0 | 15.3× |
| Easyjson | 3573 | 613.22 MB/s | 24 | 1 | 5.3× |
| Goccy | 4010 | 546.38 MB/s | 2864 | 4 | 4.8× |
| SonicFastest | 6756 | 324.32 MB/s | 3600 | 38 | 2.8× |
| Sonic | 6986 | 313.63 MB/s | 3601 | 38 | 2.7× |
| JSONV2 | 8854 | 247.47 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9218 | 196.47 MB/s | 7552 | 158 | 2.1× |
| Stdlib | 19110 | 114.65 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 537759 | 949.26 MB/s | 457537 | 1009 | 13.2× |
| Lightning | 552827 | 923.39 MB/s | 457537 | 1009 | 12.9× |
| LightningArena | 555647 | 918.71 MB/s | 457537 | 1009 | 12.8× |
| Sonic | 1253726 | 407.17 MB/s | 1308019 | 2014 | 5.7× |
| SonicFastest | 1260469 | 404.99 MB/s | 1308638 | 2014 | 5.6× |
| Goccy | 1318039 | 387.30 MB/s | 1139287 | 5006 | 5.4× |
| Easyjson | 1699457 | 300.38 MB/s | 863777 | 3012 | 4.2× |
| JSONV2 | 3557019 | 143.51 MB/s | 1075952 | 12645 | 2.0× |
| LightningDecodeAny | 3607802 | 127.91 MB/s | 2950652 | 64018 | 2.0× |
| Stdlib | 7121332 | 71.68 MB/s | 1162119 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 640 | 30909.23 MB/s | 0 | 0 | 260.2× |
| LightningArena | 641 | 30871.36 MB/s | 0 | 0 | 259.9× |
| LightningDestructive | 979 | 20222.67 MB/s | 0 | 0 | 170.2× |
| SonicFastest | 7565 | 2615.92 MB/s | 21108 | 3 | 22.0× |
| Goccy | 30911 | 640.18 MB/s | 20492 | 2 | 5.4× |
| Sonic | 32624 | 606.57 MB/s | 20687 | 3 | 5.1× |
| JSONV2 | 33327 | 593.78 MB/s | 8 | 1 | 5.0× |
| LightningDecodeAny | 94000 | 210.51 MB/s | 116864 | 2015 | 1.8× |
| Easyjson | 102027 | 193.96 MB/s | 0 | 0 | 1.6× |
| Stdlib | 166581 | 118.79 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1956 | 9265.57 MB/s | 432 | 2 | 64.3× |
| LightningArena | 1971 | 9195.45 MB/s | 432 | 2 | 63.8× |
| LightningDestructive | 2346 | 7724.73 MB/s | 0 | 0 | 53.6× |
| Easyjson | 5064 | 3578.68 MB/s | 432 | 2 | 24.8× |
| Sonic | 8552 | 2119.15 MB/s | 20445 | 5 | 14.7× |
| SonicFastest | 8576 | 2113.27 MB/s | 20446 | 5 | 14.7× |
| LightningDecodeAny | 19076 | 937.42 MB/s | 29088 | 191 | 6.6× |
| Goccy | 20835 | 869.86 MB/s | 19460 | 2 | 6.0× |
| JSONV2 | 50741 | 357.19 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 125800 | 144.07 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2247600 | 893.62 MB/s | 3089565 | 6821 | 9.7× |
| LightningArena | 2408826 | 833.81 MB/s | 3094370 | 6703 | 9.0× |
| Lightning | 2410824 | 833.12 MB/s | 3091277 | 6827 | 9.0× |
| SonicFastest | 3860868 | 520.22 MB/s | 5151297 | 7085 | 5.6× |
| Sonic | 3899084 | 515.12 MB/s | 5150853 | 7085 | 5.6× |
| Goccy | 4635347 | 433.30 MB/s | 5409144 | 15830 | 4.7× |
| Easyjson | 5653309 | 355.28 MB/s | 2981486 | 7439 | 3.9× |
| LightningDecodeAny | 7084886 | 161.23 MB/s | 8503512 | 134008 | 3.1× |
| JSONV2 | 7361441 | 272.84 MB/s | 3173677 | 14563 | 3.0× |
| Stdlib | 21785537 | 92.19 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 866 | 633.66 MB/s | 480 | 1 | 7.7× |
| LightningArena | 868 | 632.73 MB/s | 480 | 1 | 7.7× |
| LightningDestructive | 870 | 630.76 MB/s | 480 | 1 | 7.7× |
| LightningDecodeAny | 1822 | 300.71 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 2269 | 241.99 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2352 | 233.43 MB/s | 2260 | 8 | 2.8× |
| Sonic | 2386 | 230.14 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3279 | 167.42 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3561 | 154.18 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6686 | 82.12 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 466789 | 1352.89 MB/s | 402728 | 545 | 13.5× |
| LightningArena | 534625 | 1181.23 MB/s | 453017 | 712 | 11.8× |
| Lightning | 541470 | 1166.30 MB/s | 451257 | 857 | 11.6× |
| SonicFastest | 1080391 | 584.52 MB/s | 1073382 | 814 | 5.8× |
| Sonic | 1081079 | 584.15 MB/s | 1073329 | 814 | 5.8× |
| Goccy | 1310204 | 482.00 MB/s | 989078 | 1200 | 4.8× |
| Easyjson | 1380247 | 457.54 MB/s | 422504 | 936 | 4.6× |
| JSONV2 | 2401646 | 262.95 MB/s | 571590 | 3144 | 2.6× |
| LightningDecodeAny | 2553417 | 182.86 MB/s | 2076504 | 30126 | 2.5× |
| Stdlib | 6294172 | 100.33 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 666314 | 844.06 MB/s | 546568 | 429 | 9.0× |
| LightningArena | 860767 | 653.38 MB/s | 771665 | 1088 | 7.0× |
| Lightning | 863088 | 651.62 MB/s | 769937 | 1235 | 7.0× |
| SonicFastest | 1271780 | 442.22 MB/s | 1347591 | 1185 | 4.7× |
| Sonic | 1310556 | 429.14 MB/s | 1349799 | 1185 | 4.6× |
| Goccy | 1498623 | 375.28 MB/s | 1038833 | 1028 | 4.0× |
| Easyjson | 2166446 | 259.60 MB/s | 775153 | 1254 | 2.8× |
| LightningDecodeAny | 2921647 | 192.50 MB/s | 2180440 | 30126 | 2.1× |
| JSONV2 | 3156186 | 178.19 MB/s | 927407 | 3482 | 1.9× |
| Stdlib | 6004463 | 93.67 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 602524 | 884.91 MB/s | 333416 | 2084 | 10.8× |
| Lightning | 671131 | 794.45 MB/s | 368224 | 2293 | 9.7× |
| LightningArena | 675775 | 788.99 MB/s | 368224 | 2293 | 9.6× |
| Sonic | 1112393 | 479.31 MB/s | 980584 | 3082 | 5.8× |
| SonicFastest | 1113222 | 478.95 MB/s | 980511 | 3082 | 5.8× |
| Easyjson | 1316776 | 404.91 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1506717 | 353.87 MB/s | 1167044 | 5408 | 4.3× |
| JSONV2 | 2836016 | 188.00 MB/s | 745423 | 13288 | 2.3× |
| LightningDecodeAny | 3568625 | 149.41 MB/s | 2992875 | 50076 | 1.8× |
| Stdlib | 6500518 | 82.02 MB/s | 798693 | 17133 | 1.0× |
