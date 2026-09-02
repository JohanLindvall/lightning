# JSON Deserialization Benchmarks

- generated 2026-09-02T12:33:18Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 59421 | 2141.91 MB/s | 49760 | 3 | 17.0× |
| LightningArena | 59532 | 2137.92 MB/s | 49760 | 3 | 17.0× |
| LightningDestructive | 62049 | 2051.21 MB/s | 49280 | 2 | 16.3× |
| SonicFastest | 148535 | 856.87 MB/s | 214012 | 15 | 6.8× |
| Sonic | 148720 | 855.80 MB/s | 213873 | 15 | 6.8× |
| Easyjson | 180697 | 704.36 MB/s | 122864 | 14 | 5.6× |
| Goccy | 193338 | 658.30 MB/s | 225495 | 884 | 5.2× |
| JSONV2 | 319351 | 398.54 MB/s | 195128 | 1805 | 3.2× |
| LightningDecodeAny | 328706 | 287.96 MB/s | 463410 | 9708 | 3.1× |
| Stdlib | 1009531 | 126.07 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2737161 | 822.40 MB/s | 2532849 | 1143 | 8.7× |
| LightningArena | 2780703 | 809.53 MB/s | 2532850 | 1143 | 8.6× |
| Lightning | 2786263 | 807.91 MB/s | 2532851 | 1143 | 8.6× |
| Sonic | 3963082 | 568.01 MB/s | 4871562 | 2584 | 6.0× |
| SonicFastest | 4025490 | 559.20 MB/s | 4869955 | 2584 | 5.9× |
| Goccy | 9456983 | 238.03 MB/s | 4179109 | 56534 | 2.5× |
| LightningDecodeAny | 9590279 | 234.72 MB/s | 19380213 | 223896 | 2.5× |
| Easyjson | 10542744 | 213.52 MB/s | 3099810 | 2120 | 2.3× |
| JSONV2 | 13566771 | 165.92 MB/s | 3123195 | 3083 | 1.8× |
| Stdlib | 23939409 | 94.03 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 367975 | 734.84 MB/s | 397297 | 567 | 8.4× |
| Lightning | 372277 | 726.35 MB/s | 397297 | 567 | 8.3× |
| LightningDestructive | 383190 | 705.66 MB/s | 397297 | 567 | 8.1× |
| Sonic | 577163 | 468.50 MB/s | 641755 | 1147 | 5.4× |
| SonicFastest | 578784 | 467.19 MB/s | 641611 | 1147 | 5.4× |
| Goccy | 1305830 | 207.07 MB/s | 541992 | 8122 | 2.4× |
| Easyjson | 1351722 | 200.04 MB/s | 330272 | 749 | 2.3× |
| LightningDecodeAny | 1658263 | 163.06 MB/s | 2543876 | 29687 | 1.9× |
| JSONV2 | 1783160 | 151.64 MB/s | 348160 | 1628 | 1.7× |
| Stdlib | 3100215 | 87.22 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 729737 | 2366.88 MB/s | 765560 | 2798 | 17.6× |
| LightningArena | 743327 | 2323.61 MB/s | 768416 | 2440 | 17.3× |
| Lightning | 747952 | 2309.25 MB/s | 765601 | 2799 | 17.2× |
| Sonic | 1732010 | 997.23 MB/s | 2695492 | 5547 | 7.4× |
| SonicFastest | 1732610 | 996.88 MB/s | 2696007 | 5547 | 7.4× |
| Goccy | 1890649 | 913.55 MB/s | 2581217 | 14603 | 6.8× |
| LightningDecodeAny | 2911551 | 171.83 MB/s | 4953692 | 76576 | 4.4× |
| Easyjson | 3082229 | 560.37 MB/s | 972032 | 5389 | 4.2× |
| JSONV2 | 3283196 | 526.07 MB/s | 1011612 | 7594 | 3.9× |
| Stdlib | 12832094 | 134.60 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 595 | 3044.47 MB/s | 0 | 0 | 20.8× |
| Lightning | 604 | 3000.48 MB/s | 0 | 0 | 20.5× |
| LightningDestructive | 616 | 2942.31 MB/s | 0 | 0 | 20.1× |
| Easyjson | 2203 | 822.68 MB/s | 24 | 1 | 5.6× |
| Goccy | 2588 | 700.18 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 4767 | 380.09 MB/s | 3346 | 38 | 2.6× |
| Sonic | 4929 | 367.65 MB/s | 3345 | 38 | 2.5× |
| JSONV2 | 6127 | 295.74 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 6587 | 274.94 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 12359 | 146.62 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 635 | 2852.31 MB/s | 0 | 0 | 19.5× |
| LightningArena | 651 | 2783.96 MB/s | 0 | 0 | 19.0× |
| LightningDestructive | 655 | 2766.34 MB/s | 0 | 0 | 18.9× |
| Easyjson | 2212 | 819.25 MB/s | 24 | 1 | 5.6× |
| Goccy | 2591 | 699.35 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 4853 | 373.41 MB/s | 3345 | 38 | 2.6× |
| Sonic | 4996 | 362.71 MB/s | 3345 | 38 | 2.5× |
| JSONV2 | 6044 | 299.81 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 6493 | 278.90 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 12378 | 146.39 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 806 | 2248.08 MB/s | 144 | 10 | 15.2× |
| LightningArena | 808 | 2241.65 MB/s | 144 | 10 | 15.1× |
| LightningDestructive | 815 | 2223.36 MB/s | 144 | 10 | 15.0× |
| Easyjson | 2313 | 783.27 MB/s | 144 | 10 | 5.3× |
| Goccy | 2389 | 758.56 MB/s | 2600 | 5 | 5.1× |
| SonicFastest | 4896 | 370.09 MB/s | 3362 | 40 | 2.5× |
| Sonic | 5073 | 357.18 MB/s | 3367 | 40 | 2.4× |
| JSONV2 | 5957 | 304.19 MB/s | 632 | 7 | 2.1× |
| LightningDecodeAny | 6433 | 281.51 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 12227 | 148.19 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 590 | 837.87 MB/s | 160 | 1 | 8.1× |
| LightningDestructive | 598 | 826.11 MB/s | 160 | 1 | 8.0× |
| Sonic | 955 | 517.50 MB/s | 1075 | 8 | 5.0× |
| SonicFastest | 956 | 516.58 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1135 | 434.54 MB/s | 1296 | 26 | 4.2× |
| LightningArena | 1248 | 395.69 MB/s | 4096 | 1 | 3.8× |
| Easyjson | 1781 | 277.31 MB/s | 448 | 3 | 2.7× |
| Goccy | 2010 | 245.76 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 2343 | 210.87 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4761 | 103.76 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 327 | 703.35 MB/s | 160 | 1 | 10.5× |
| LightningDestructive | 328 | 700.42 MB/s | 160 | 1 | 10.4× |
| SonicFastest | 685 | 335.97 MB/s | 801 | 8 | 5.0× |
| Sonic | 686 | 335.27 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 911 | 251.49 MB/s | 1296 | 26 | 3.8× |
| LightningArena | 968 | 237.57 MB/s | 4096 | 1 | 3.5× |
| Easyjson | 1171 | 196.33 MB/s | 448 | 3 | 2.9× |
| Goccy | 1341 | 171.55 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 1822 | 126.24 MB/s | 528 | 7 | 1.9× |
| Stdlib | 3421 | 67.23 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 44148 | 1475.31 MB/s | 103440 | 103 | 11.8× |
| LightningArena | 44438 | 1465.68 MB/s | 103440 | 103 | 11.7× |
| LightningDestructive | 44861 | 1451.88 MB/s | 97220 | 98 | 11.6× |
| Sonic | 118738 | 548.54 MB/s | 235826 | 65 | 4.4× |
| SonicFastest | 126401 | 515.28 MB/s | 236387 | 65 | 4.1× |
| LightningDecodeAny | 138576 | 384.84 MB/s | 180048 | 3245 | 3.7× |
| Goccy | 168469 | 386.61 MB/s | 229977 | 134 | 3.1× |
| JSONV2 | 207518 | 313.86 MB/s | 206666 | 607 | 2.5× |
| Stdlib | 519333 | 125.41 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1714822 | 1131.59 MB/s | 2864592 | 1380 | 11.9× |
| LightningArena | 1779828 | 1090.26 MB/s | 2864593 | 1380 | 11.5× |
| Lightning | 1793916 | 1081.70 MB/s | 2864593 | 1380 | 11.4× |
| Goccy | 3926745 | 494.17 MB/s | 4063630 | 13509 | 5.2× |
| Sonic | 4919553 | 394.44 MB/s | 4879294 | 1736 | 4.2× |
| SonicFastest | 4923772 | 394.10 MB/s | 4880609 | 1736 | 4.1× |
| Easyjson | 5985667 | 324.19 MB/s | 3871264 | 15043 | 3.4× |
| LightningDecodeAny | 6878792 | 282.09 MB/s | 7063040 | 218633 | 3.0× |
| JSONV2 | 8651976 | 224.28 MB/s | 3237183 | 13947 | 2.4× |
| Stdlib | 20422329 | 95.02 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 656819 | 5066.59 MB/s | 351704 | 1286 | 30.3× |
| Lightning | 1044719 | 3185.38 MB/s | 2488905 | 2995 | 19.0× |
| LightningArena | 1045901 | 3181.78 MB/s | 2488906 | 2995 | 19.0× |
| SonicFastest | 1750801 | 1900.75 MB/s | 5896467 | 4263 | 11.4× |
| Sonic | 1752200 | 1899.23 MB/s | 5896596 | 4263 | 11.3× |
| LightningDecodeAny | 2430557 | 1264.63 MB/s | 4876912 | 56892 | 8.2× |
| Goccy | 4288260 | 776.03 MB/s | 3948914 | 3816 | 4.6× |
| JSONV2 | 5948176 | 559.47 MB/s | 5364505 | 13243 | 3.3× |
| Stdlib | 19885084 | 167.35 MB/s | 5565610 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 147153 | 1497.39 MB/s | 135872 | 226 | 12.1× |
| LightningArena | 149401 | 1474.87 MB/s | 135872 | 226 | 11.9× |
| LightningDestructive | 156917 | 1404.22 MB/s | 135872 | 226 | 11.3× |
| Goccy | 356442 | 618.18 MB/s | 364214 | 1066 | 5.0× |
| SonicFastest | 412025 | 534.79 MB/s | 351193 | 262 | 4.3× |
| Sonic | 420030 | 524.60 MB/s | 351830 | 262 | 4.2× |
| Easyjson | 455812 | 483.41 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 523366 | 421.02 MB/s | 129747 | 470 | 3.4× |
| LightningDecodeAny | 728838 | 148.61 MB/s | 897218 | 11703 | 2.4× |
| Stdlib | 1773319 | 124.26 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 7863609 | 1030.07 MB/s | 11845072 | 20816 | 10.0× |
| Lightning | 8128285 | 996.52 MB/s | 11845074 | 20816 | 9.7× |
| LightningArena | 8149641 | 993.91 MB/s | 11845073 | 20816 | 9.7× |
| SonicFastest | 15798260 | 512.72 MB/s | 19859742 | 41640 | 5.0× |
| Sonic | 15965892 | 507.33 MB/s | 19857899 | 41640 | 4.9× |
| Goccy | 19341227 | 418.80 MB/s | 19122428 | 107155 | 4.1× |
| Easyjson | 26198746 | 309.18 MB/s | 15059618 | 41643 | 3.0× |
| LightningDecodeAny | 29716025 | 175.09 MB/s | 46279352 | 747112 | 2.7× |
| JSONV2 | 34901316 | 232.08 MB/s | 15233748 | 78972 | 2.3× |
| Stdlib | 78846693 | 102.73 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 3942517 | 756.74 MB/s | 3764712 | 1504 | 10.3× |
| LightningDestructive | 4097596 | 728.10 MB/s | 3758856 | 29356 | 10.0× |
| Lightning | 4190001 | 712.04 MB/s | 3758857 | 29356 | 9.7× |
| Sonic | 7242872 | 411.92 MB/s | 9129801 | 57804 | 5.6× |
| SonicFastest | 7253296 | 411.33 MB/s | 9129650 | 57804 | 5.6× |
| LightningDecodeAny | 13195576 | 139.00 MB/s | 23982580 | 351152 | 3.1× |
| Goccy | 13589421 | 219.54 MB/s | 9808245 | 273617 | 3.0× |
| Easyjson | 14064023 | 212.13 MB/s | 9479441 | 30115 | 2.9× |
| JSONV2 | 19736507 | 151.16 MB/s | 9257055 | 86278 | 2.1× |
| Stdlib | 40783630 | 73.15 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 955919 | 756.96 MB/s | 907600 | 3618 | 10.6× |
| LightningArena | 967876 | 747.61 MB/s | 911392 | 30 | 10.4× |
| Lightning | 997109 | 725.69 MB/s | 907594 | 3618 | 10.1× |
| Sonic | 1607493 | 450.14 MB/s | 2374066 | 3683 | 6.3× |
| SonicFastest | 1621700 | 446.20 MB/s | 2374197 | 3683 | 6.2× |
| LightningDecodeAny | 3892408 | 167.14 MB/s | 6500462 | 76546 | 2.6× |
| Easyjson | 4121588 | 175.56 MB/s | 2847908 | 3698 | 2.5× |
| Goccy | 4226306 | 171.21 MB/s | 2684082 | 80266 | 2.4× |
| JSONV2 | 4853897 | 149.08 MB/s | 2704707 | 7318 | 2.1× |
| Stdlib | 10108071 | 71.59 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1393373 | 1132.04 MB/s | 907600 | 3618 | 10.1× |
| Lightning | 1433573 | 1100.30 MB/s | 907595 | 3618 | 9.8× |
| LightningArena | 1463093 | 1078.09 MB/s | 911392 | 30 | 9.6× |
| SonicFastest | 1925770 | 819.08 MB/s | 3225752 | 3683 | 7.3× |
| Sonic | 1996686 | 789.99 MB/s | 3232093 | 3683 | 7.0× |
| LightningDecodeAny | 3420825 | 220.24 MB/s | 6500456 | 76546 | 4.1× |
| Easyjson | 4907035 | 321.45 MB/s | 2847906 | 3698 | 2.9× |
| Goccy | 5211596 | 302.66 MB/s | 3500825 | 80262 | 2.7× |
| JSONV2 | 5259768 | 299.89 MB/s | 2704553 | 7318 | 2.7× |
| Stdlib | 14037946 | 112.36 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 162127 | 925.97 MB/s | 81920 | 1 | 9.8× |
| Lightning | 162728 | 922.55 MB/s | 81920 | 1 | 9.7× |
| LightningDestructive | 168904 | 888.81 MB/s | 81920 | 1 | 9.4× |
| Sonic | 299256 | 501.66 MB/s | 408148 | 16 | 5.3× |
| SonicFastest | 315672 | 475.57 MB/s | 408969 | 16 | 5.0× |
| LightningDecodeAny | 407972 | 367.97 MB/s | 745763 | 10016 | 3.9× |
| Goccy | 767589 | 195.58 MB/s | 327165 | 10005 | 2.1× |
| JSONV2 | 896284 | 167.50 MB/s | 357728 | 20 | 1.8× |
| Stdlib | 1582010 | 94.89 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 20762 | 1354.27 MB/s | 29216 | 103 | 12.4× |
| Lightning | 20838 | 1349.34 MB/s | 29216 | 103 | 12.4× |
| LightningDestructive | 22001 | 1277.96 MB/s | 29088 | 101 | 11.7× |
| Sonic | 53682 | 523.77 MB/s | 59473 | 83 | 4.8× |
| SonicFastest | 53836 | 522.28 MB/s | 59466 | 83 | 4.8× |
| Easyjson | 58778 | 478.36 MB/s | 32304 | 138 | 4.4× |
| Goccy | 60599 | 463.98 MB/s | 59283 | 188 | 4.3× |
| JSONV2 | 99062 | 283.83 MB/s | 36896 | 242 | 2.6× |
| LightningDecodeAny | 116883 | 240.56 MB/s | 140576 | 2643 | 2.2× |
| Stdlib | 257551 | 109.17 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1154 | 2016.99 MB/s | 32 | 1 | 16.8× |
| LightningArena | 1158 | 2009.53 MB/s | 32 | 1 | 16.7× |
| LightningDestructive | 1222 | 1904.44 MB/s | 32 | 1 | 15.8× |
| Goccy | 3653 | 637.28 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 3938 | 591.21 MB/s | 192 | 2 | 4.9× |
| SonicFastest | 4770 | 488.10 MB/s | 3708 | 4 | 4.1× |
| Sonic | 4792 | 485.80 MB/s | 3708 | 4 | 4.0× |
| JSONV2 | 6051 | 384.74 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 7492 | 224.89 MB/s | 10200 | 195 | 2.6× |
| Stdlib | 19344 | 120.35 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 130 | 1459.81 MB/s | 0 | 0 | 15.6× |
| Lightning | 130 | 1454.87 MB/s | 0 | 0 | 15.6× |
| LightningDestructive | 132 | 1426.66 MB/s | 0 | 0 | 15.3× |
| Goccy | 340 | 556.01 MB/s | 304 | 2 | 6.0× |
| Easyjson | 443 | 426.49 MB/s | 0 | 0 | 4.6× |
| Sonic | 569 | 332.42 MB/s | 341 | 3 | 3.6× |
| SonicFastest | 572 | 330.61 MB/s | 341 | 3 | 3.5× |
| JSONV2 | 724 | 260.86 MB/s | 112 | 1 | 2.8× |
| LightningDecodeAny | 933 | 143.69 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 2024 | 93.38 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 823 | 2660.96 MB/s | 0 | 0 | 17.4× |
| LightningArena | 835 | 2623.44 MB/s | 0 | 0 | 17.1× |
| LightningDestructive | 851 | 2573.93 MB/s | 0 | 0 | 16.8× |
| Easyjson | 2732 | 801.87 MB/s | 24 | 1 | 5.2× |
| Goccy | 2911 | 752.79 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 5206 | 420.86 MB/s | 3597 | 38 | 2.7× |
| Sonic | 5588 | 392.08 MB/s | 3596 | 38 | 2.6× |
| JSONV2 | 6185 | 354.26 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 6422 | 282.02 MB/s | 7552 | 158 | 2.2× |
| Stdlib | 14312 | 153.09 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 379403 | 1345.47 MB/s | 457537 | 1009 | 13.7× |
| Lightning | 394590 | 1293.69 MB/s | 457537 | 1009 | 13.2× |
| LightningArena | 397023 | 1285.76 MB/s | 457537 | 1009 | 13.1× |
| Goccy | 979298 | 521.27 MB/s | 1136987 | 5006 | 5.3× |
| Sonic | 1094815 | 466.27 MB/s | 1309089 | 2014 | 4.7× |
| SonicFastest | 1095510 | 465.97 MB/s | 1308342 | 2014 | 4.7× |
| Easyjson | 1172020 | 435.55 MB/s | 863777 | 3012 | 4.4× |
| JSONV2 | 2387062 | 213.85 MB/s | 1075956 | 12645 | 2.2× |
| LightningDecodeAny | 2576483 | 179.11 MB/s | 2950650 | 64018 | 2.0× |
| Stdlib | 5190886 | 98.34 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 471 | 42033.55 MB/s | 0 | 0 | 256.9× |
| LightningArena | 473 | 41862.14 MB/s | 0 | 0 | 255.8× |
| LightningDestructive | 674 | 29379.35 MB/s | 0 | 0 | 179.5× |
| SonicFastest | 5460 | 3624.11 MB/s | 21091 | 3 | 22.2× |
| Goccy | 20792 | 951.76 MB/s | 20492 | 2 | 5.8× |
| Sonic | 22744 | 870.07 MB/s | 20610 | 3 | 5.3× |
| JSONV2 | 28098 | 704.29 MB/s | 8 | 1 | 4.3× |
| LightningDecodeAny | 74593 | 265.28 MB/s | 116864 | 2015 | 1.6× |
| Easyjson | 93077 | 212.61 MB/s | 0 | 0 | 1.3× |
| Stdlib | 120940 | 163.63 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1385 | 13081.47 MB/s | 432 | 2 | 75.7× |
| LightningArena | 1392 | 13023.99 MB/s | 432 | 2 | 75.3× |
| LightningDestructive | 1443 | 12562.19 MB/s | 0 | 0 | 72.6× |
| Easyjson | 3650 | 4965.91 MB/s | 432 | 2 | 28.7× |
| SonicFastest | 7284 | 2488.19 MB/s | 20461 | 5 | 14.4× |
| Sonic | 7841 | 2311.58 MB/s | 20408 | 5 | 13.4× |
| LightningDecodeAny | 14083 | 1269.80 MB/s | 29088 | 191 | 7.4× |
| Goccy | 20772 | 872.52 MB/s | 19460 | 2 | 5.0× |
| JSONV2 | 37366 | 485.04 MB/s | 16500 | 50 | 2.8× |
| Stdlib | 104826 | 172.90 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1563159 | 1284.89 MB/s | 3089565 | 6821 | 10.5× |
| LightningArena | 1616024 | 1242.86 MB/s | 3094370 | 6703 | 10.1× |
| Lightning | 1694027 | 1185.63 MB/s | 3091278 | 6827 | 9.7× |
| Goccy | 3636484 | 552.32 MB/s | 5410755 | 15832 | 4.5× |
| SonicFastest | 3809829 | 527.19 MB/s | 5154514 | 7085 | 4.3× |
| Sonic | 3874915 | 518.33 MB/s | 5153561 | 7085 | 4.2× |
| Easyjson | 3951666 | 508.27 MB/s | 2981488 | 7439 | 4.1× |
| LightningDecodeAny | 5172905 | 220.82 MB/s | 8503513 | 134008 | 3.2× |
| JSONV2 | 5273989 | 380.83 MB/s | 3173679 | 14563 | 3.1× |
| Stdlib | 16370122 | 122.69 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 628 | 874.42 MB/s | 480 | 1 | 7.7× |
| Lightning | 629 | 872.91 MB/s | 480 | 1 | 7.6× |
| LightningDestructive | 638 | 860.94 MB/s | 480 | 1 | 7.5× |
| LightningDecodeAny | 1331 | 411.83 MB/s | 2021 | 46 | 3.6× |
| Easyjson | 1569 | 350.00 MB/s | 1616 | 5 | 3.1× |
| SonicFastest | 1639 | 334.86 MB/s | 2262 | 8 | 2.9× |
| Sonic | 1694 | 323.99 MB/s | 2261 | 8 | 2.8× |
| Goccy | 2331 | 235.56 MB/s | 2129 | 43 | 2.1× |
| JSONV2 | 2362 | 232.46 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 4810 | 114.13 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 314139 | 2010.30 MB/s | 402728 | 545 | 15.5× |
| LightningArena | 367878 | 1716.64 MB/s | 453016 | 712 | 13.2× |
| Lightning | 377142 | 1674.47 MB/s | 451256 | 857 | 12.9× |
| Sonic | 827571 | 763.09 MB/s | 1065032 | 814 | 5.9× |
| SonicFastest | 831127 | 759.83 MB/s | 1065141 | 814 | 5.8× |
| Easyjson | 995729 | 634.22 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1084892 | 582.10 MB/s | 990237 | 1200 | 4.5× |
| JSONV2 | 1743393 | 362.23 MB/s | 571590 | 3144 | 2.8× |
| LightningDecodeAny | 1804294 | 258.77 MB/s | 2076504 | 30126 | 2.7× |
| Stdlib | 4854720 | 130.08 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 465578 | 1207.98 MB/s | 546569 | 429 | 9.9× |
| LightningArena | 615098 | 914.34 MB/s | 771664 | 1088 | 7.5× |
| Lightning | 623841 | 901.52 MB/s | 769936 | 1235 | 7.4× |
| Sonic | 1082862 | 519.37 MB/s | 1348366 | 1184 | 4.3× |
| SonicFastest | 1086600 | 517.59 MB/s | 1348165 | 1184 | 4.2× |
| Goccy | 1198570 | 469.23 MB/s | 1036334 | 1028 | 3.8× |
| Easyjson | 1558406 | 360.89 MB/s | 775153 | 1254 | 3.0× |
| LightningDecodeAny | 2157671 | 260.66 MB/s | 2180441 | 30126 | 2.1× |
| JSONV2 | 2314926 | 242.95 MB/s | 927403 | 3482 | 2.0× |
| Stdlib | 4607634 | 122.06 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 399104 | 1335.94 MB/s | 333416 | 2084 | 12.2× |
| Lightning | 460861 | 1156.92 MB/s | 368224 | 2293 | 10.6× |
| LightningArena | 467155 | 1141.33 MB/s | 368224 | 2293 | 10.4× |
| Easyjson | 971215 | 548.98 MB/s | 428362 | 3273 | 5.0× |
| SonicFastest | 1018665 | 523.41 MB/s | 980736 | 3082 | 4.8× |
| Sonic | 1035969 | 514.67 MB/s | 980619 | 3082 | 4.7× |
| Goccy | 1195854 | 445.86 MB/s | 1167073 | 5408 | 4.1× |
| JSONV2 | 2124408 | 250.98 MB/s | 745422 | 13288 | 2.3× |
| LightningDecodeAny | 2531051 | 210.65 MB/s | 2992875 | 50076 | 1.9× |
| Stdlib | 4878245 | 109.30 MB/s | 798692 | 17133 | 1.0× |
