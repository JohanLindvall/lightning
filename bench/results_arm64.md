# JSON Deserialization Benchmarks

- generated 2026-08-13T07:29:53Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 103109 | 1234.37 MB/s | 49760 | 3 | 10.7× |
| LightningArena | 103168 | 1233.67 MB/s | 49760 | 3 | 10.7× |
| LightningDestructive | 103348 | 1231.52 MB/s | 49280 | 2 | 10.7× |
| Sonic | 187765 | 677.84 MB/s | 197640 | 10 | 5.9× |
| SonicFastest | 188071 | 676.74 MB/s | 199623 | 10 | 5.9× |
| Goccy | 199497 | 637.98 MB/s | 224933 | 884 | 5.6× |
| Easyjson | 213636 | 595.76 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 423302 | 300.67 MB/s | 195120 | 1805 | 2.6× |
| LightningDecodeAny | 451473 | 209.65 MB/s | 463410 | 9708 | 2.5× |
| Stdlib | 1107350 | 114.94 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3676495 | 612.28 MB/s | 2532848 | 1143 | 7.1× |
| Lightning | 3706680 | 607.30 MB/s | 2532851 | 1143 | 7.1× |
| LightningArena | 3715106 | 605.92 MB/s | 2532849 | 1143 | 7.0× |
| SonicFastest | 4517267 | 498.32 MB/s | 15232101 | 970 | 5.8× |
| Sonic | 4521258 | 497.88 MB/s | 15232100 | 970 | 5.8× |
| Goccy | 10248073 | 219.66 MB/s | 4114629 | 56532 | 2.6× |
| Easyjson | 10925235 | 206.04 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12756865 | 176.46 MB/s | 19380210 | 223896 | 2.1× |
| JSONV2 | 15918483 | 141.41 MB/s | 3123216 | 3083 | 1.6× |
| Stdlib | 26161638 | 86.04 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 479668 | 563.73 MB/s | 397296 | 567 | 7.0× |
| LightningDestructive | 480665 | 562.56 MB/s | 397297 | 567 | 7.0× |
| LightningArena | 481055 | 562.10 MB/s | 397296 | 567 | 7.0× |
| Sonic | 644758 | 419.39 MB/s | 481572 | 968 | 5.2× |
| SonicFastest | 646152 | 418.48 MB/s | 485922 | 968 | 5.2× |
| Easyjson | 1398876 | 193.30 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1408235 | 192.02 MB/s | 542685 | 8122 | 2.4× |
| LightningDecodeAny | 1771130 | 152.67 MB/s | 2543882 | 29687 | 1.9× |
| JSONV2 | 2078686 | 130.08 MB/s | 348147 | 1628 | 1.6× |
| Stdlib | 3362753 | 80.41 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1159305 | 1489.86 MB/s | 765560 | 2798 | 11.4× |
| LightningArena | 1169788 | 1476.51 MB/s | 768416 | 2440 | 11.3× |
| Lightning | 1170634 | 1475.44 MB/s | 765602 | 2799 | 11.3× |
| SonicFastest | 2073328 | 833.06 MB/s | 2673300 | 4020 | 6.4× |
| Sonic | 2074146 | 832.73 MB/s | 2674799 | 4020 | 6.4× |
| Goccy | 2363942 | 730.65 MB/s | 2583130 | 14605 | 5.6× |
| Easyjson | 4229390 | 408.38 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4289331 | 402.67 MB/s | 1011633 | 7594 | 3.1× |
| LightningDecodeAny | 4329952 | 115.54 MB/s | 4953693 | 76576 | 3.1× |
| Stdlib | 13238670 | 130.47 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1122 | 1615.42 MB/s | 0 | 0 | 12.5× |
| LightningArena | 1123 | 1613.84 MB/s | 0 | 0 | 12.5× |
| LightningDestructive | 1137 | 1593.50 MB/s | 0 | 0 | 12.3× |
| Easyjson | 2562 | 707.32 MB/s | 24 | 1 | 5.5× |
| Goccy | 2824 | 641.65 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6017 | 301.14 MB/s | 3791 | 40 | 2.3× |
| Sonic | 6051 | 299.48 MB/s | 3798 | 40 | 2.3× |
| JSONV2 | 7839 | 231.15 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8022 | 225.75 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 14026 | 129.19 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1142 | 1587.27 MB/s | 0 | 0 | 12.4× |
| LightningArena | 1144 | 1584.52 MB/s | 0 | 0 | 12.3× |
| LightningDestructive | 1156 | 1567.60 MB/s | 0 | 0 | 12.2× |
| Easyjson | 2539 | 713.66 MB/s | 24 | 1 | 5.6× |
| Goccy | 2892 | 626.48 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6077 | 298.16 MB/s | 3878 | 40 | 2.3× |
| Sonic | 6084 | 297.81 MB/s | 3870 | 40 | 2.3× |
| JSONV2 | 7803 | 232.20 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8374 | 216.26 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 14105 | 128.46 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1305 | 1388.18 MB/s | 144 | 10 | 10.8× |
| Lightning | 1311 | 1382.60 MB/s | 144 | 10 | 10.7× |
| LightningDestructive | 1372 | 1320.79 MB/s | 144 | 10 | 10.2× |
| Easyjson | 2761 | 656.32 MB/s | 144 | 10 | 5.1× |
| Goccy | 2980 | 608.03 MB/s | 2600 | 5 | 4.7× |
| Sonic | 6199 | 292.33 MB/s | 3851 | 42 | 2.3× |
| SonicFastest | 6222 | 291.22 MB/s | 3806 | 42 | 2.3× |
| JSONV2 | 7897 | 229.45 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8023 | 225.71 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 14029 | 129.16 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 704 | 701.56 MB/s | 160 | 1 | 7.8× |
| Lightning | 705 | 700.74 MB/s | 160 | 1 | 7.8× |
| Sonic | 1241 | 397.95 MB/s | 980 | 6 | 4.5× |
| SonicFastest | 1246 | 396.56 MB/s | 986 | 6 | 4.4× |
| LightningArena | 1345 | 367.24 MB/s | 4096 | 1 | 4.1× |
| LightningDecodeAny | 1397 | 352.86 MB/s | 1296 | 26 | 4.0× |
| Easyjson | 2236 | 220.94 MB/s | 448 | 3 | 2.5× |
| Goccy | 2415 | 204.59 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3241 | 152.41 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5523 | 89.45 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 436 | 527.70 MB/s | 160 | 1 | 9.5× |
| Lightning | 445 | 516.93 MB/s | 160 | 1 | 9.3× |
| SonicFastest | 883 | 260.56 MB/s | 657 | 6 | 4.7× |
| Sonic | 887 | 259.41 MB/s | 659 | 6 | 4.6× |
| LightningDecodeAny | 1160 | 197.41 MB/s | 1296 | 26 | 3.6× |
| LightningArena | 1161 | 198.03 MB/s | 4096 | 1 | 3.5× |
| Easyjson | 1410 | 163.13 MB/s | 448 | 3 | 2.9× |
| Goccy | 1606 | 143.24 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2447 | 93.98 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4121 | 55.82 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 58445 | 1114.42 MB/s | 97220 | 98 | 9.4× |
| Lightning | 59701 | 1090.97 MB/s | 103440 | 103 | 9.2× |
| LightningArena | 59705 | 1090.89 MB/s | 103440 | 103 | 9.2× |
| SonicFastest | 99940 | 651.71 MB/s | 156000 | 75 | 5.5× |
| Sonic | 100388 | 648.80 MB/s | 156377 | 75 | 5.4× |
| Goccy | 148556 | 438.43 MB/s | 229234 | 134 | 3.7× |
| LightningDecodeAny | 181204 | 294.30 MB/s | 180048 | 3245 | 3.0× |
| JSONV2 | 226915 | 287.03 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 546603 | 119.16 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2543735 | 762.84 MB/s | 2864592 | 1380 | 9.2× |
| LightningArena | 2591492 | 748.79 MB/s | 2864593 | 1380 | 9.1× |
| Lightning | 2610330 | 743.38 MB/s | 2864594 | 1380 | 9.0× |
| Goccy | 4835501 | 401.30 MB/s | 4064924 | 13510 | 4.9× |
| SonicFastest | 5011211 | 387.23 MB/s | 14606973 | 1407 | 4.7× |
| Sonic | 5017640 | 386.73 MB/s | 14608594 | 1407 | 4.7× |
| Easyjson | 7514034 | 258.25 MB/s | 3871267 | 15043 | 3.1× |
| LightningDecodeAny | 9304833 | 208.54 MB/s | 7063039 | 218633 | 2.5× |
| JSONV2 | 11338612 | 171.14 MB/s | 3237222 | 13947 | 2.1× |
| Stdlib | 23495079 | 82.59 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1075921 | 3093.01 MB/s | 351704 | 1286 | 19.5× |
| Lightning | 1767524 | 1882.76 MB/s | 2488904 | 2995 | 11.8× |
| LightningArena | 1785514 | 1863.79 MB/s | 2488904 | 2995 | 11.7× |
| Sonic | 2754987 | 1207.93 MB/s | 6441754 | 4248 | 7.6× |
| SonicFastest | 2760193 | 1205.65 MB/s | 6410104 | 4248 | 7.6× |
| LightningDecodeAny | 3570233 | 860.94 MB/s | 4876912 | 56892 | 5.9× |
| Goccy | 4598567 | 723.67 MB/s | 3948909 | 3816 | 4.6× |
| JSONV2 | 7526029 | 442.18 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20930559 | 158.99 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 216731 | 1016.68 MB/s | 135872 | 226 | 9.4× |
| Lightning | 217687 | 1012.22 MB/s | 135872 | 226 | 9.4× |
| LightningDestructive | 219549 | 1003.63 MB/s | 135872 | 226 | 9.3× |
| Sonic | 381847 | 577.05 MB/s | 303568 | 398 | 5.4× |
| SonicFastest | 382461 | 576.13 MB/s | 301949 | 398 | 5.3× |
| Goccy | 440788 | 499.89 MB/s | 364264 | 1067 | 4.6× |
| Easyjson | 545646 | 403.83 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 730694 | 301.56 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 862437 | 125.59 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2043027 | 107.85 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11426567 | 708.88 MB/s | 11845078 | 20816 | 7.8× |
| Lightning | 11697748 | 692.44 MB/s | 11845078 | 20816 | 7.7× |
| LightningArena | 11772165 | 688.07 MB/s | 11845073 | 20816 | 7.6× |
| Sonic | 16846785 | 480.81 MB/s | 70902343 | 40014 | 5.3× |
| SonicFastest | 17184965 | 471.34 MB/s | 70887315 | 40014 | 5.2× |
| Goccy | 23310517 | 347.48 MB/s | 16999961 | 107148 | 3.8× |
| Easyjson | 30647842 | 264.29 MB/s | 15059617 | 41643 | 2.9× |
| LightningDecodeAny | 37148072 | 140.06 MB/s | 46279355 | 747112 | 2.4× |
| JSONV2 | 43853816 | 184.71 MB/s | 15233734 | 78972 | 2.0× |
| Stdlib | 89536918 | 90.47 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5400622 | 552.43 MB/s | 3764712 | 1504 | 8.8× |
| LightningDestructive | 5699437 | 523.47 MB/s | 3758856 | 29356 | 8.3× |
| Lightning | 5835796 | 511.24 MB/s | 3758859 | 29356 | 8.1× |
| Sonic | 8635771 | 345.48 MB/s | 26640862 | 56760 | 5.5× |
| SonicFastest | 8666271 | 344.26 MB/s | 26599991 | 56760 | 5.5× |
| Easyjson | 16490708 | 180.92 MB/s | 9479441 | 30115 | 2.9× |
| LightningDecodeAny | 16597048 | 110.51 MB/s | 23982580 | 351152 | 2.9× |
| Goccy | 16779156 | 177.81 MB/s | 10492557 | 273644 | 2.8× |
| JSONV2 | 24768451 | 120.45 MB/s | 9257145 | 86278 | 1.9× |
| Stdlib | 47370573 | 62.98 MB/s | 9258095 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1211074 | 597.48 MB/s | 911393 | 30 | 9.6× |
| LightningDestructive | 1255330 | 576.42 MB/s | 907601 | 3618 | 9.2× |
| Lightning | 1268926 | 570.24 MB/s | 907595 | 3618 | 9.1× |
| Sonic | 1801901 | 401.57 MB/s | 3198798 | 7226 | 6.4× |
| SonicFastest | 1811090 | 399.54 MB/s | 3190928 | 7226 | 6.4× |
| Easyjson | 4246879 | 170.38 MB/s | 2847906 | 3698 | 2.7× |
| LightningDecodeAny | 4318773 | 150.64 MB/s | 6500459 | 76546 | 2.7× |
| Goccy | 4813912 | 150.31 MB/s | 2805211 | 80273 | 2.4× |
| JSONV2 | 5823231 | 124.26 MB/s | 2704690 | 7318 | 2.0× |
| Stdlib | 11598478 | 62.39 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1840066 | 857.23 MB/s | 907601 | 3618 | 8.6× |
| LightningArena | 1845406 | 854.75 MB/s | 911393 | 30 | 8.5× |
| Lightning | 1894368 | 832.65 MB/s | 907594 | 3618 | 8.3× |
| SonicFastest | 2265077 | 696.38 MB/s | 5785967 | 7226 | 7.0× |
| Sonic | 2265628 | 696.21 MB/s | 5783135 | 7226 | 7.0× |
| LightningDecodeAny | 3984363 | 189.09 MB/s | 6500456 | 76546 | 4.0× |
| Easyjson | 5620886 | 280.62 MB/s | 2847907 | 3698 | 2.8× |
| Goccy | 5702551 | 276.60 MB/s | 3590193 | 80268 | 2.8× |
| JSONV2 | 6434799 | 245.13 MB/s | 2704594 | 7318 | 2.5× |
| Stdlib | 15766566 | 100.04 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 212373 | 706.89 MB/s | 81920 | 1 | 8.6× |
| LightningArena | 212589 | 706.17 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 212597 | 706.14 MB/s | 81920 | 1 | 8.6× |
| SonicFastest | 287209 | 522.70 MB/s | 289668 | 6 | 6.4× |
| Sonic | 288265 | 520.78 MB/s | 288150 | 6 | 6.3× |
| LightningDecodeAny | 483904 | 310.23 MB/s | 745764 | 10016 | 3.8× |
| Goccy | 914013 | 164.25 MB/s | 328681 | 10005 | 2.0× |
| JSONV2 | 1109154 | 135.35 MB/s | 357715 | 20 | 1.6× |
| Stdlib | 1824511 | 82.28 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 31272 | 899.11 MB/s | 29216 | 103 | 9.7× |
| LightningArena | 31395 | 895.58 MB/s | 29216 | 103 | 9.7× |
| LightningDestructive | 31628 | 888.99 MB/s | 29088 | 101 | 9.6× |
| SonicFastest | 64192 | 438.02 MB/s | 46974 | 103 | 4.7× |
| Sonic | 64440 | 436.33 MB/s | 47432 | 103 | 4.7× |
| Easyjson | 68245 | 412.00 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72564 | 387.48 MB/s | 59207 | 188 | 4.2× |
| JSONV2 | 134784 | 208.61 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 148465 | 189.39 MB/s | 140576 | 2643 | 2.0× |
| Stdlib | 303568 | 92.62 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1899 | 1226.00 MB/s | 32 | 1 | 12.0× |
| LightningArena | 1908 | 1220.36 MB/s | 32 | 1 | 11.9× |
| LightningDestructive | 1976 | 1177.95 MB/s | 32 | 1 | 11.5× |
| Goccy | 4162 | 559.34 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4202 | 554.00 MB/s | 192 | 2 | 5.4× |
| Sonic | 5133 | 453.50 MB/s | 4289 | 6 | 4.4× |
| SonicFastest | 5174 | 449.91 MB/s | 4356 | 6 | 4.4× |
| JSONV2 | 8461 | 275.15 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9918 | 169.90 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22757 | 102.30 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 216 | 873.73 MB/s | 0 | 0 | 11.2× |
| Lightning | 217 | 869.49 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 218 | 868.99 MB/s | 0 | 0 | 11.1× |
| Goccy | 386 | 490.27 MB/s | 304 | 2 | 6.3× |
| Easyjson | 484 | 390.08 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 788 | 239.72 MB/s | 504 | 4 | 3.1× |
| Sonic | 790 | 239.25 MB/s | 510 | 4 | 3.1× |
| JSONV2 | 1037 | 182.24 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1195 | 112.17 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2424 | 77.97 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1470 | 1490.61 MB/s | 0 | 0 | 10.9× |
| LightningArena | 1479 | 1481.26 MB/s | 0 | 0 | 10.8× |
| LightningDestructive | 1486 | 1474.38 MB/s | 0 | 0 | 10.8× |
| Easyjson | 3181 | 688.71 MB/s | 24 | 1 | 5.0× |
| Goccy | 3273 | 669.45 MB/s | 2864 | 4 | 4.9× |
| Sonic | 6536 | 335.23 MB/s | 4117 | 40 | 2.5× |
| SonicFastest | 6542 | 334.89 MB/s | 4145 | 40 | 2.4× |
| JSONV2 | 8005 | 273.72 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8370 | 216.37 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 16018 | 136.79 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 642823 | 794.12 MB/s | 457536 | 1009 | 9.4× |
| Lightning | 649348 | 786.14 MB/s | 457536 | 1009 | 9.3× |
| LightningArena | 649805 | 785.58 MB/s | 457536 | 1009 | 9.3× |
| Goccy | 1163715 | 438.66 MB/s | 1137544 | 5006 | 5.2× |
| SonicFastest | 1172159 | 435.50 MB/s | 907767 | 2006 | 5.1× |
| Sonic | 1173358 | 435.06 MB/s | 910669 | 2006 | 5.1× |
| Easyjson | 1550010 | 329.34 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3205918 | 159.23 MB/s | 1076012 | 12646 | 1.9× |
| LightningDecodeAny | 3316719 | 139.13 MB/s | 2950649 | 64018 | 1.8× |
| Stdlib | 6017998 | 84.82 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1334 | 14833.70 MB/s | 0 | 0 | 93.6× |
| LightningArena | 1334 | 14835.83 MB/s | 0 | 0 | 93.6× |
| LightningDestructive | 1359 | 14561.44 MB/s | 0 | 0 | 91.9× |
| Goccy | 20049 | 987.01 MB/s | 20491 | 2 | 6.2× |
| SonicFastest | 27415 | 721.82 MB/s | 22788 | 4 | 4.6× |
| Sonic | 27494 | 719.75 MB/s | 22865 | 4 | 4.5× |
| JSONV2 | 29837 | 663.24 MB/s | 8 | 1 | 4.2× |
| LightningDecodeAny | 75780 | 261.13 MB/s | 116864 | 2015 | 1.6× |
| Easyjson | 91602 | 216.03 MB/s | 0 | 0 | 1.4× |
| Stdlib | 124849 | 158.50 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2534 | 7152.03 MB/s | 0 | 0 | 40.6× |
| LightningArena | 2681 | 6759.04 MB/s | 432 | 2 | 38.4× |
| Lightning | 2686 | 6747.59 MB/s | 432 | 2 | 38.3× |
| Easyjson | 3952 | 4586.58 MB/s | 432 | 2 | 26.0× |
| Sonic | 10111 | 1792.45 MB/s | 22816 | 6 | 10.2× |
| SonicFastest | 10330 | 1754.50 MB/s | 23389 | 6 | 10.0× |
| Goccy | 16009 | 1132.11 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 17003 | 1051.72 MB/s | 29088 | 191 | 6.0× |
| JSONV2 | 46112 | 393.04 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102842 | 176.23 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2464868 | 814.85 MB/s | 3089565 | 6821 | 7.6× |
| Lightning | 2541234 | 790.36 MB/s | 3091279 | 6827 | 7.4× |
| LightningArena | 2565255 | 782.96 MB/s | 3094371 | 6703 | 7.3× |
| Goccy | 4457165 | 450.62 MB/s | 5413525 | 15837 | 4.2× |
| SonicFastest | 4636827 | 433.16 MB/s | 10953998 | 13683 | 4.1× |
| Sonic | 4672841 | 429.82 MB/s | 10902387 | 13683 | 4.0× |
| Easyjson | 4994332 | 402.15 MB/s | 2981516 | 7439 | 3.8× |
| LightningDecodeAny | 6908563 | 165.34 MB/s | 8503513 | 134008 | 2.7× |
| JSONV2 | 7123169 | 281.97 MB/s | 3173684 | 14563 | 2.6× |
| Stdlib | 18819828 | 106.72 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 905 | 606.61 MB/s | 480 | 1 | 6.3× |
| LightningArena | 906 | 605.80 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 908 | 604.36 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1685 | 325.14 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 2226 | 246.65 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2697 | 203.56 MB/s | 1955 | 26 | 2.1× |
| SonicFastest | 2711 | 202.51 MB/s | 1965 | 26 | 2.1× |
| Goccy | 3059 | 179.47 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3360 | 163.38 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5691 | 96.46 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 499034 | 1265.47 MB/s | 402728 | 545 | 10.9× |
| Lightning | 570024 | 1107.87 MB/s | 451257 | 857 | 9.5× |
| LightningArena | 572031 | 1103.99 MB/s | 453017 | 712 | 9.5× |
| SonicFastest | 1025796 | 615.63 MB/s | 986390 | 1102 | 5.3× |
| Sonic | 1034292 | 610.58 MB/s | 996253 | 1102 | 5.2× |
| Easyjson | 1158486 | 545.12 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1172447 | 538.63 MB/s | 986147 | 1202 | 4.6× |
| JSONV2 | 2171773 | 290.78 MB/s | 571617 | 3144 | 2.5× |
| LightningDecodeAny | 2408161 | 193.88 MB/s | 2076505 | 30126 | 2.3× |
| Stdlib | 5423476 | 116.44 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 611082 | 920.35 MB/s | 546571 | 429 | 8.7× |
| LightningArena | 786025 | 715.51 MB/s | 771666 | 1088 | 6.7× |
| Lightning | 789516 | 712.35 MB/s | 769938 | 1235 | 6.7× |
| Sonic | 1029652 | 546.21 MB/s | 921811 | 1476 | 5.1× |
| SonicFastest | 1042138 | 539.67 MB/s | 941539 | 1476 | 5.1× |
| Goccy | 1349943 | 416.62 MB/s | 1038293 | 1030 | 3.9× |
| Easyjson | 1754460 | 320.56 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2652718 | 212.01 MB/s | 2180440 | 30126 | 2.0× |
| JSONV2 | 2776760 | 202.54 MB/s | 927451 | 3482 | 1.9× |
| Stdlib | 5294255 | 106.23 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 652645 | 816.95 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 673943 | 791.13 MB/s | 368224 | 2293 | 8.1× |
| LightningArena | 677572 | 786.89 MB/s | 368224 | 2293 | 8.0× |
| Easyjson | 1095552 | 486.68 MB/s | 428361 | 3273 | 5.0× |
| SonicFastest | 1160334 | 459.50 MB/s | 1048103 | 4351 | 4.7× |
| Sonic | 1160943 | 459.26 MB/s | 1053350 | 4351 | 4.7× |
| Goccy | 1310406 | 406.88 MB/s | 1167224 | 5409 | 4.2× |
| JSONV2 | 2527541 | 210.95 MB/s | 745449 | 13288 | 2.2× |
| LightningDecodeAny | 3412053 | 156.26 MB/s | 2992877 | 50076 | 1.6× |
| Stdlib | 5451415 | 97.81 MB/s | 798692 | 17133 | 1.0× |
