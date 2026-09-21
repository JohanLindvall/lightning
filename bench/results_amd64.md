# JSON Deserialization Benchmarks

- generated 2026-09-21T08:22:10Z
- go version go1.26.8 linux/amd64
- cpu: AMD EPYC 9V45 96-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 46965 | 2710.02 MB/s | 49917 | 2 | 15.2× |
| LightningArena | 47976 | 2652.87 MB/s | 49941 | 2 | 14.8× |
| LightningDestructive | 57012 | 2232.44 MB/s | 49280 | 2 | 12.5× |
| SonicFastest | 105786 | 1203.14 MB/s | 215493 | 15 | 6.7× |
| Sonic | 106714 | 1192.67 MB/s | 215525 | 15 | 6.7× |
| Easyjson | 125653 | 1012.91 MB/s | 122864 | 14 | 5.7× |
| Goccy | 138173 | 921.13 MB/s | 225630 | 884 | 5.2× |
| JSONV2 | 231057 | 550.84 MB/s | 195128 | 1805 | 3.1× |
| LightningDecodeAny | 244677 | 386.85 MB/s | 465145 | 9706 | 2.9× |
| Stdlib | 712434 | 178.65 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1764260 | 1275.92 MB/s | 2532848 | 1143 | 8.9× |
| Lightning | 1772555 | 1269.95 MB/s | 2532850 | 1143 | 8.9× |
| LightningArena | 1775876 | 1267.57 MB/s | 2532849 | 1143 | 8.9× |
| Sonic | 2799104 | 804.20 MB/s | 4892422 | 2584 | 5.6× |
| SonicFastest | 2837078 | 793.44 MB/s | 4892246 | 2584 | 5.6× |
| LightningDecodeAny | 4854378 | 463.71 MB/s | 6828998 | 223498 | 3.2× |
| Goccy | 6738196 | 334.07 MB/s | 4263816 | 56539 | 2.3× |
| Easyjson | 7045358 | 319.51 MB/s | 3099808 | 2120 | 2.2× |
| JSONV2 | 8693658 | 258.93 MB/s | 3123190 | 3083 | 1.8× |
| Stdlib | 15749078 | 142.93 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 252666 | 1070.20 MB/s | 397297 | 567 | 8.2× |
| LightningArena | 252801 | 1069.63 MB/s | 397297 | 567 | 8.2× |
| LightningDestructive | 258095 | 1047.69 MB/s | 397297 | 567 | 8.0× |
| Sonic | 425675 | 635.23 MB/s | 641091 | 1147 | 4.9× |
| SonicFastest | 430917 | 627.51 MB/s | 641349 | 1147 | 4.8× |
| LightningDecodeAny | 692858 | 390.27 MB/s | 845691 | 29656 | 3.0× |
| Easyjson | 916201 | 295.13 MB/s | 330272 | 749 | 2.3× |
| Goccy | 920422 | 293.78 MB/s | 541481 | 8122 | 2.3× |
| JSONV2 | 1162528 | 232.60 MB/s | 348159 | 1628 | 1.8× |
| Stdlib | 2075593 | 130.28 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 519522 | 3324.60 MB/s | 765560 | 2798 | 17.6× |
| Lightning | 525913 | 3284.20 MB/s | 767712 | 2798 | 17.4× |
| LightningArena | 527023 | 3277.28 MB/s | 774871 | 2444 | 17.3× |
| Goccy | 1222644 | 1412.68 MB/s | 2580307 | 14603 | 7.5× |
| Sonic | 1247055 | 1385.03 MB/s | 2707353 | 5547 | 7.3× |
| SonicFastest | 1255888 | 1375.28 MB/s | 2707196 | 5547 | 7.3× |
| LightningDecodeAny | 1977846 | 252.95 MB/s | 4491649 | 67881 | 4.6× |
| Easyjson | 2014297 | 857.47 MB/s | 972032 | 5389 | 4.5× |
| JSONV2 | 2207521 | 782.42 MB/s | 1011611 | 7594 | 4.1× |
| Stdlib | 9141122 | 188.95 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 393 | 4611.56 MB/s | 0 | 0 | 22.7× |
| LightningArena | 393 | 4607.17 MB/s | 0 | 0 | 22.7× |
| LightningDestructive | 414 | 4374.95 MB/s | 0 | 0 | 21.5× |
| Easyjson | 1494 | 1212.94 MB/s | 24 | 1 | 6.0× |
| Goccy | 1795 | 1009.46 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 3506 | 516.79 MB/s | 3346 | 38 | 2.5× |
| Sonic | 3674 | 493.15 MB/s | 3346 | 38 | 2.4× |
| JSONV2 | 4139 | 437.78 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 4641 | 390.20 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 8916 | 203.24 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 399 | 4538.47 MB/s | 0 | 0 | 22.8× |
| LightningArena | 402 | 4510.86 MB/s | 0 | 0 | 22.6× |
| LightningDestructive | 426 | 4258.93 MB/s | 0 | 0 | 21.4× |
| Easyjson | 1480 | 1224.41 MB/s | 24 | 1 | 6.1× |
| Goccy | 1831 | 989.76 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 3519 | 514.90 MB/s | 3344 | 38 | 2.6× |
| Sonic | 3648 | 496.65 MB/s | 3344 | 38 | 2.5× |
| JSONV2 | 4379 | 413.76 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 4669 | 387.87 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 9086 | 199.44 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 504 | 3594.00 MB/s | 144 | 10 | 18.0× |
| LightningArena | 530 | 3420.80 MB/s | 144 | 10 | 17.1× |
| LightningDestructive | 548 | 3308.54 MB/s | 144 | 10 | 16.6× |
| Easyjson | 1579 | 1147.64 MB/s | 144 | 10 | 5.7× |
| Goccy | 1675 | 1081.68 MB/s | 2600 | 5 | 5.4× |
| SonicFastest | 3534 | 512.76 MB/s | 3361 | 40 | 2.6× |
| Sonic | 3557 | 509.40 MB/s | 3359 | 40 | 2.6× |
| JSONV2 | 4179 | 433.55 MB/s | 632 | 7 | 2.2× |
| LightningDecodeAny | 4655 | 389.05 MB/s | 7552 | 158 | 2.0× |
| Stdlib | 9079 | 199.59 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 358 | 1379.50 MB/s | 160 | 1 | 9.1× |
| LightningDestructive | 366 | 1349.26 MB/s | 160 | 1 | 8.9× |
| Sonic | 724 | 682.80 MB/s | 1076 | 8 | 4.5× |
| SonicFastest | 729 | 678.01 MB/s | 1076 | 8 | 4.5× |
| LightningDecodeAny | 739 | 667.36 MB/s | 1040 | 25 | 4.4× |
| LightningArena | 871 | 567.27 MB/s | 4120 | 2 | 3.8× |
| Easyjson | 1191 | 414.92 MB/s | 448 | 3 | 2.8× |
| Goccy | 1396 | 353.80 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 1778 | 277.79 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3276 | 150.81 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 217 | 1061.20 MB/s | 160 | 1 | 10.7× |
| LightningDestructive | 220 | 1047.72 MB/s | 160 | 1 | 10.5× |
| Sonic | 505 | 455.68 MB/s | 801 | 8 | 4.6× |
| SonicFastest | 512 | 449.59 MB/s | 800 | 8 | 4.5× |
| LightningDecodeAny | 579 | 395.35 MB/s | 1040 | 25 | 4.0× |
| LightningArena | 731 | 314.52 MB/s | 4120 | 2 | 3.2× |
| Easyjson | 793 | 290.06 MB/s | 448 | 3 | 2.9× |
| Goccy | 923 | 249.19 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 1245 | 184.80 MB/s | 528 | 7 | 1.9× |
| Stdlib | 2309 | 99.61 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 31294 | 2081.29 MB/s | 97220 | 98 | 11.8× |
| Lightning | 31310 | 2080.24 MB/s | 103782 | 99 | 11.8× |
| LightningArena | 32195 | 2023.03 MB/s | 103771 | 99 | 11.5× |
| Sonic | 92745 | 702.27 MB/s | 235805 | 65 | 4.0× |
| LightningDecodeAny | 97967 | 544.36 MB/s | 176776 | 3237 | 3.8× |
| SonicFastest | 103415 | 629.81 MB/s | 236030 | 65 | 3.6× |
| Goccy | 123758 | 526.28 MB/s | 228294 | 134 | 3.0× |
| JSONV2 | 181481 | 358.89 MB/s | 206666 | 607 | 2.0× |
| Stdlib | 369227 | 176.40 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1357631 | 1429.31 MB/s | 2185296 | 1350 | 11.1× |
| Lightning | 1371365 | 1414.99 MB/s | 2185297 | 1350 | 11.0× |
| LightningArena | 1418108 | 1368.35 MB/s | 2185297 | 1350 | 10.6× |
| Goccy | 2674256 | 725.61 MB/s | 4061975 | 13509 | 5.6× |
| SonicFastest | 2771924 | 700.05 MB/s | 4879269 | 1736 | 5.4× |
| Sonic | 2959672 | 655.64 MB/s | 4879131 | 1736 | 5.1× |
| Easyjson | 4589660 | 422.79 MB/s | 3871265 | 15043 | 3.3× |
| JSONV2 | 5623456 | 345.07 MB/s | 3237187 | 13947 | 2.7× |
| LightningDecodeAny | 6022665 | 322.19 MB/s | 6627986 | 206416 | 2.5× |
| Stdlib | 15057058 | 128.87 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 472351 | 7045.25 MB/s | 351704 | 1286 | 29.2× |
| LightningArena | 692852 | 4803.09 MB/s | 2434658 | 1413 | 19.9× |
| Lightning | 707656 | 4702.61 MB/s | 2434692 | 1413 | 19.5× |
| SonicFastest | 1359216 | 2448.35 MB/s | 5896635 | 4263 | 10.1× |
| Sonic | 1378832 | 2413.51 MB/s | 5896612 | 4263 | 10.0× |
| LightningDecodeAny | 1711323 | 1796.13 MB/s | 4825392 | 55311 | 8.1× |
| Goccy | 2786037 | 1194.47 MB/s | 3948914 | 3817 | 4.9× |
| JSONV2 | 4373197 | 760.96 MB/s | 5364503 | 13243 | 3.2× |
| Stdlib | 13788204 | 241.35 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 93969 | 2344.87 MB/s | 135392 | 226 | 13.9× |
| LightningDestructive | 94823 | 2323.77 MB/s | 135392 | 226 | 13.7× |
| Lightning | 94830 | 2323.60 MB/s | 135392 | 226 | 13.7× |
| Goccy | 250842 | 878.42 MB/s | 364559 | 1066 | 5.2× |
| Easyjson | 297012 | 741.88 MB/s | 130512 | 245 | 4.4× |
| Sonic | 336650 | 654.52 MB/s | 350866 | 262 | 3.9× |
| SonicFastest | 339659 | 648.73 MB/s | 350666 | 262 | 3.8× |
| JSONV2 | 355580 | 619.68 MB/s | 129746 | 470 | 3.7× |
| LightningDecodeAny | 544571 | 198.90 MB/s | 854738 | 11700 | 2.4× |
| Stdlib | 1303248 | 169.07 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 5335575 | 1518.12 MB/s | 8109649 | 20809 | 10.1× |
| LightningDestructive | 5509254 | 1470.26 MB/s | 8109648 | 20809 | 9.7× |
| LightningArena | 5534881 | 1463.45 MB/s | 8109648 | 20809 | 9.7× |
| Sonic | 11524117 | 702.88 MB/s | 19869813 | 41640 | 4.7× |
| SonicFastest | 11925214 | 679.24 MB/s | 19869915 | 41640 | 4.5× |
| Goccy | 14048637 | 576.57 MB/s | 19012549 | 107155 | 3.8× |
| Easyjson | 17232985 | 470.03 MB/s | 15059617 | 41643 | 3.1× |
| LightningDecodeAny | 18105508 | 287.37 MB/s | 28359829 | 746961 | 3.0× |
| JSONV2 | 25452458 | 318.24 MB/s | 15233710 | 78972 | 2.1× |
| Stdlib | 53707592 | 150.82 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2960072 | 1007.90 MB/s | 3758857 | 29356 | 10.5× |
| LightningArena | 3012423 | 990.39 MB/s | 3780456 | 1514 | 10.3× |
| LightningDestructive | 3086836 | 966.51 MB/s | 3758856 | 29356 | 10.0× |
| SonicFastest | 5481511 | 544.28 MB/s | 9133510 | 57804 | 5.6× |
| Sonic | 5647426 | 528.29 MB/s | 9133887 | 57804 | 5.5× |
| Goccy | 10463156 | 285.14 MB/s | 9909918 | 273621 | 3.0× |
| LightningDecodeAny | 10971002 | 167.19 MB/s | 18225371 | 350883 | 2.8× |
| Easyjson | 11167611 | 267.15 MB/s | 9479441 | 30115 | 2.8× |
| JSONV2 | 14827786 | 201.21 MB/s | 9257032 | 86278 | 2.1× |
| Stdlib | 30949585 | 96.40 MB/s | 9258083 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 593293 | 1219.63 MB/s | 907600 | 3618 | 11.5× |
| LightningArena | 596082 | 1213.92 MB/s | 916256 | 37 | 11.5× |
| Lightning | 686412 | 1054.17 MB/s | 907597 | 3618 | 10.0× |
| Sonic | 1155950 | 625.98 MB/s | 2382576 | 3683 | 5.9× |
| SonicFastest | 1217439 | 594.36 MB/s | 2382592 | 3683 | 5.6× |
| LightningDecodeAny | 2538487 | 256.28 MB/s | 5691591 | 76540 | 2.7× |
| Easyjson | 2909099 | 248.74 MB/s | 2847908 | 3698 | 2.4× |
| Goccy | 3034004 | 238.50 MB/s | 2678313 | 80266 | 2.3× |
| JSONV2 | 3360251 | 215.34 MB/s | 2704714 | 7318 | 2.0× |
| Stdlib | 6848056 | 105.66 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 831727 | 1896.48 MB/s | 916256 | 37 | 11.9× |
| Lightning | 845428 | 1865.74 MB/s | 907593 | 3618 | 11.7× |
| LightningDestructive | 854273 | 1846.43 MB/s | 907600 | 3618 | 11.6× |
| Sonic | 1543471 | 1021.95 MB/s | 3234976 | 3683 | 6.4× |
| SonicFastest | 1578009 | 999.58 MB/s | 3229830 | 3683 | 6.3× |
| LightningDecodeAny | 2131725 | 353.42 MB/s | 5691590 | 76540 | 4.6× |
| Easyjson | 3468390 | 454.78 MB/s | 2847905 | 3698 | 2.9× |
| Goccy | 3881428 | 406.38 MB/s | 3509333 | 80263 | 2.5× |
| JSONV2 | 4006546 | 393.69 MB/s | 2704552 | 7318 | 2.5× |
| Stdlib | 9892152 | 159.45 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 107461 | 1397.01 MB/s | 81920 | 1 | 10.5× |
| Lightning | 111962 | 1340.85 MB/s | 81920 | 1 | 10.0× |
| LightningDestructive | 112659 | 1332.55 MB/s | 81920 | 1 | 10.0× |
| Sonic | 238421 | 629.66 MB/s | 408047 | 16 | 4.7× |
| SonicFastest | 242843 | 618.19 MB/s | 408375 | 16 | 4.6× |
| LightningDecodeAny | 268400 | 559.32 MB/s | 745508 | 10015 | 4.2× |
| Goccy | 624933 | 240.22 MB/s | 326356 | 10005 | 1.8× |
| JSONV2 | 640872 | 234.25 MB/s | 357727 | 20 | 1.8× |
| Stdlib | 1124517 | 133.50 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 14496 | 1939.70 MB/s | 29276 | 101 | 12.7× |
| Lightning | 14606 | 1925.06 MB/s | 29276 | 101 | 12.6× |
| LightningDestructive | 15118 | 1859.85 MB/s | 29088 | 101 | 12.1× |
| SonicFastest | 40857 | 688.17 MB/s | 59496 | 83 | 4.5× |
| Easyjson | 41467 | 678.06 MB/s | 32304 | 138 | 4.4× |
| Sonic | 41702 | 674.24 MB/s | 59489 | 83 | 4.4× |
| Goccy | 44209 | 636.00 MB/s | 59287 | 188 | 4.1× |
| JSONV2 | 70665 | 397.89 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 84859 | 331.34 MB/s | 133983 | 2639 | 2.2× |
| Stdlib | 183435 | 153.28 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 809 | 2876.77 MB/s | 32 | 1 | 17.3× |
| Lightning | 822 | 2830.26 MB/s | 32 | 1 | 17.0× |
| LightningDestructive | 846 | 2750.32 MB/s | 32 | 1 | 16.6× |
| Goccy | 2543 | 915.31 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 2693 | 864.39 MB/s | 192 | 2 | 5.2× |
| SonicFastest | 3797 | 613.17 MB/s | 3704 | 4 | 3.7× |
| Sonic | 3852 | 604.35 MB/s | 3705 | 4 | 3.6× |
| JSONV2 | 4151 | 560.88 MB/s | 1000 | 6 | 3.4× |
| LightningDecodeAny | 5296 | 318.17 MB/s | 9936 | 194 | 2.6× |
| Stdlib | 14018 | 166.07 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 82 | 2314.76 MB/s | 0 | 0 | 16.9× |
| Lightning | 83 | 2268.73 MB/s | 0 | 0 | 16.6× |
| LightningDestructive | 86 | 2194.46 MB/s | 0 | 0 | 16.0× |
| Goccy | 229 | 825.11 MB/s | 304 | 2 | 6.0× |
| Easyjson | 297 | 635.81 MB/s | 0 | 0 | 4.6× |
| SonicFastest | 486 | 388.70 MB/s | 341 | 3 | 2.8× |
| Sonic | 500 | 378.14 MB/s | 341 | 3 | 2.8× |
| JSONV2 | 531 | 355.90 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 640 | 209.52 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 1381 | 136.84 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 572 | 3828.38 MB/s | 0 | 0 | 19.6× |
| LightningArena | 592 | 3699.34 MB/s | 0 | 0 | 19.0× |
| LightningDestructive | 640 | 3424.27 MB/s | 0 | 0 | 17.6× |
| Easyjson | 1876 | 1168.22 MB/s | 24 | 1 | 6.0× |
| Goccy | 2035 | 1076.76 MB/s | 2864 | 4 | 5.5× |
| SonicFastest | 4056 | 540.15 MB/s | 3600 | 38 | 2.8× |
| Sonic | 4211 | 520.26 MB/s | 3599 | 38 | 2.7× |
| JSONV2 | 4628 | 473.43 MB/s | 640 | 6 | 2.4× |
| LightningDecodeAny | 4925 | 367.69 MB/s | 7552 | 158 | 2.3× |
| Stdlib | 11239 | 194.95 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 276808 | 1844.15 MB/s | 318400 | 1005 | 15.1× |
| Lightning | 277436 | 1839.98 MB/s | 318400 | 1005 | 15.1× |
| LightningArena | 281760 | 1811.74 MB/s | 318400 | 1005 | 14.8× |
| Goccy | 684781 | 745.46 MB/s | 1140536 | 5006 | 6.1× |
| Easyjson | 879259 | 580.58 MB/s | 863776 | 3012 | 4.8× |
| SonicFastest | 931817 | 547.83 MB/s | 1305611 | 2013 | 4.5× |
| Sonic | 1014037 | 503.41 MB/s | 1306552 | 2013 | 4.1× |
| JSONV2 | 1814815 | 281.28 MB/s | 1075951 | 12645 | 2.3× |
| LightningDecodeAny | 2125841 | 217.07 MB/s | 2742394 | 64017 | 2.0× |
| Stdlib | 4183765 | 122.01 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 635 | 31184.24 MB/s | 0 | 0 | 144.0× |
| Lightning | 638 | 31005.51 MB/s | 0 | 0 | 143.2× |
| LightningDestructive | 760 | 26039.96 MB/s | 0 | 0 | 120.3× |
| SonicFastest | 4375 | 4523.11 MB/s | 21140 | 3 | 20.9× |
| Goccy | 14295 | 1384.32 MB/s | 20492 | 2 | 6.4× |
| JSONV2 | 18807 | 1052.23 MB/s | 8 | 1 | 4.9× |
| Sonic | 20390 | 970.54 MB/s | 20645 | 3 | 4.5× |
| Easyjson | 63558 | 311.35 MB/s | 0 | 0 | 1.4× |
| LightningDecodeAny | 66468 | 297.71 MB/s | 116608 | 2014 | 1.4× |
| Stdlib | 91384 | 216.55 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1024 | 17693.00 MB/s | 405 | 0 | 74.4× |
| LightningArena | 1068 | 16968.34 MB/s | 405 | 0 | 71.3× |
| LightningDestructive | 1282 | 14135.68 MB/s | 0 | 0 | 59.4× |
| Easyjson | 2434 | 7447.47 MB/s | 432 | 2 | 31.3× |
| SonicFastest | 6638 | 2730.38 MB/s | 20392 | 5 | 11.5× |
| Sonic | 6672 | 2716.33 MB/s | 20415 | 5 | 11.4× |
| LightningDecodeAny | 12743 | 1403.33 MB/s | 29134 | 189 | 6.0× |
| Goccy | 15589 | 1162.61 MB/s | 19460 | 2 | 4.9× |
| JSONV2 | 27429 | 660.76 MB/s | 16500 | 50 | 2.8× |
| Stdlib | 76174 | 237.93 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1251252 | 1605.19 MB/s | 3089565 | 6821 | 10.5× |
| LightningArena | 1254547 | 1600.97 MB/s | 3099653 | 6699 | 10.5× |
| Lightning | 1456481 | 1379.00 MB/s | 3096410 | 6822 | 9.1× |
| Goccy | 2498888 | 803.76 MB/s | 5411462 | 15832 | 5.3× |
| Easyjson | 2852063 | 704.23 MB/s | 2981484 | 7439 | 4.6× |
| Sonic | 3016995 | 665.73 MB/s | 5171271 | 7085 | 4.4× |
| SonicFastest | 3082181 | 651.65 MB/s | 5168616 | 7085 | 4.3× |
| LightningDecodeAny | 3944784 | 289.57 MB/s | 7376010 | 134004 | 3.3× |
| JSONV2 | 4049611 | 495.97 MB/s | 3173674 | 14562 | 3.3× |
| Stdlib | 13198409 | 152.18 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 479 | 1146.02 MB/s | 480 | 1 | 7.4× |
| Lightning | 481 | 1140.62 MB/s | 480 | 1 | 7.4× |
| LightningArena | 485 | 1132.42 MB/s | 480 | 1 | 7.4× |
| LightningDecodeAny | 901 | 608.07 MB/s | 1765 | 45 | 4.0× |
| Easyjson | 1181 | 465.00 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 1239 | 443.16 MB/s | 2261 | 8 | 2.9× |
| Sonic | 1293 | 424.59 MB/s | 2261 | 8 | 2.8× |
| Goccy | 1685 | 325.91 MB/s | 2129 | 43 | 2.1× |
| JSONV2 | 1807 | 303.83 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 3566 | 153.94 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 236750 | 2667.43 MB/s | 400488 | 545 | 16.3× |
| Lightning | 266195 | 2372.37 MB/s | 446803 | 548 | 14.5× |
| LightningArena | 267846 | 2357.75 MB/s | 448572 | 404 | 14.4× |
| Easyjson | 688714 | 916.95 MB/s | 422504 | 936 | 5.6× |
| Sonic | 696915 | 906.16 MB/s | 1066376 | 814 | 5.5× |
| SonicFastest | 712489 | 886.35 MB/s | 1067241 | 814 | 5.4× |
| Goccy | 764523 | 826.02 MB/s | 984672 | 1200 | 5.0× |
| JSONV2 | 1274707 | 495.42 MB/s | 571592 | 3144 | 3.0× |
| LightningDecodeAny | 1411577 | 330.77 MB/s | 1990861 | 29072 | 2.7× |
| Stdlib | 3852257 | 163.93 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 336400 | 1671.84 MB/s | 390676 | 426 | 10.7× |
| LightningArena | 390128 | 1441.60 MB/s | 507808 | 287 | 9.2× |
| Lightning | 394763 | 1424.67 MB/s | 506064 | 433 | 9.1× |
| Sonic | 893298 | 629.59 MB/s | 1350216 | 1185 | 4.0× |
| SonicFastest | 896451 | 627.37 MB/s | 1351011 | 1185 | 4.0× |
| Goccy | 904160 | 622.02 MB/s | 1031924 | 1027 | 4.0× |
| Easyjson | 1184822 | 474.68 MB/s | 775152 | 1254 | 3.0× |
| LightningDecodeAny | 1464652 | 383.99 MB/s | 1990034 | 28580 | 2.4× |
| JSONV2 | 1832880 | 306.84 MB/s | 927407 | 3482 | 2.0× |
| Stdlib | 3582785 | 156.98 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 306131 | 1741.67 MB/s | 333416 | 2084 | 12.1× |
| LightningArena | 336913 | 1582.54 MB/s | 367759 | 2086 | 11.0× |
| Lightning | 337405 | 1580.23 MB/s | 367698 | 2086 | 10.9× |
| Easyjson | 740426 | 720.10 MB/s | 428362 | 3273 | 5.0× |
| Goccy | 880417 | 605.60 MB/s | 1167043 | 5408 | 4.2× |
| Sonic | 912521 | 584.29 MB/s | 983596 | 3082 | 4.0× |
| SonicFastest | 918911 | 580.23 MB/s | 983139 | 3082 | 4.0× |
| JSONV2 | 1644561 | 324.21 MB/s | 745421 | 13288 | 2.2× |
| LightningDecodeAny | 1947416 | 273.79 MB/s | 2658083 | 49348 | 1.9× |
| Stdlib | 3693501 | 144.36 MB/s | 798692 | 17133 | 1.0× |
