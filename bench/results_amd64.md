# JSON Deserialization Benchmarks

- generated 2026-08-06T12:47:11Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 99720 | 1276.32 MB/s | 49760 | 3 | 13.2× |
| Lightning | 100133 | 1271.06 MB/s | 49760 | 3 | 13.1× |
| LightningDestructive | 101602 | 1252.68 MB/s | 49280 | 2 | 12.9× |
| SonicFastest | 195404 | 651.34 MB/s | 213803 | 15 | 6.7× |
| Sonic | 195951 | 649.53 MB/s | 213868 | 15 | 6.7× |
| Easyjson | 224963 | 565.76 MB/s | 122864 | 14 | 5.8× |
| Goccy | 246807 | 515.69 MB/s | 225067 | 884 | 5.3× |
| JSONV2 | 416634 | 305.48 MB/s | 195128 | 1805 | 3.2× |
| LightningDecodeAny | 423752 | 223.37 MB/s | 463410 | 9708 | 3.1× |
| Stdlib | 1313612 | 96.89 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4058878 | 554.60 MB/s | 2532849 | 1143 | 7.7× |
| LightningArena | 4156488 | 541.58 MB/s | 2532850 | 1143 | 7.5× |
| Lightning | 4159218 | 541.22 MB/s | 2532851 | 1143 | 7.5× |
| Sonic | 4693547 | 479.61 MB/s | 4869717 | 2584 | 6.7× |
| SonicFastest | 5109232 | 440.59 MB/s | 4872142 | 2584 | 6.1× |
| Goccy | 12780122 | 176.14 MB/s | 4123024 | 56531 | 2.4× |
| LightningDecodeAny | 13412128 | 167.84 MB/s | 19380210 | 223896 | 2.3× |
| Easyjson | 14007350 | 160.70 MB/s | 3099810 | 2120 | 2.2× |
| JSONV2 | 16947136 | 132.83 MB/s | 3123198 | 3083 | 1.8× |
| Stdlib | 31264891 | 72.00 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 536265 | 504.23 MB/s | 397296 | 567 | 7.6× |
| Lightning | 537046 | 503.50 MB/s | 397296 | 567 | 7.6× |
| LightningDestructive | 550681 | 491.03 MB/s | 397297 | 567 | 7.4× |
| SonicFastest | 741298 | 364.77 MB/s | 641300 | 1147 | 5.5× |
| Sonic | 741974 | 364.44 MB/s | 641417 | 1147 | 5.5× |
| Goccy | 1752623 | 154.28 MB/s | 541739 | 8122 | 2.3× |
| Easyjson | 1752705 | 154.28 MB/s | 330272 | 749 | 2.3× |
| LightningDecodeAny | 2119630 | 127.57 MB/s | 2543877 | 29687 | 1.9× |
| JSONV2 | 2212264 | 122.23 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4070865 | 66.42 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1105088 | 1562.96 MB/s | 768416 | 2440 | 15.2× |
| LightningDestructive | 1106968 | 1560.30 MB/s | 765560 | 2798 | 15.2× |
| Lightning | 1125204 | 1535.01 MB/s | 765602 | 2799 | 15.0× |
| SonicFastest | 2079346 | 830.65 MB/s | 2693625 | 5547 | 8.1× |
| Sonic | 2087417 | 827.44 MB/s | 2693705 | 5547 | 8.1× |
| Goccy | 2380746 | 725.49 MB/s | 2581132 | 14603 | 7.1× |
| Easyjson | 3646630 | 473.64 MB/s | 972032 | 5389 | 4.6× |
| LightningDecodeAny | 3776221 | 132.49 MB/s | 4953692 | 76576 | 4.5× |
| JSONV2 | 4159306 | 415.26 MB/s | 1011615 | 7594 | 4.0× |
| Stdlib | 16833390 | 102.61 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1022 | 1773.24 MB/s | 0 | 0 | 15.5× |
| Lightning | 1031 | 1758.29 MB/s | 0 | 0 | 15.3× |
| LightningDestructive | 1069 | 1695.24 MB/s | 0 | 0 | 14.8× |
| Easyjson | 2746 | 659.77 MB/s | 24 | 1 | 5.8× |
| Goccy | 3299 | 549.18 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6226 | 291.06 MB/s | 3346 | 38 | 2.5× |
| Sonic | 6429 | 281.85 MB/s | 3342 | 38 | 2.5× |
| JSONV2 | 7785 | 232.75 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8253 | 219.42 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15821 | 114.53 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1069 | 1695.39 MB/s | 0 | 0 | 14.8× |
| LightningArena | 1069 | 1694.88 MB/s | 0 | 0 | 14.8× |
| LightningDestructive | 1117 | 1621.66 MB/s | 0 | 0 | 14.1× |
| Easyjson | 2743 | 660.68 MB/s | 24 | 1 | 5.8× |
| Goccy | 3291 | 550.57 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6196 | 292.43 MB/s | 3346 | 38 | 2.5× |
| Sonic | 6337 | 285.95 MB/s | 3344 | 38 | 2.5× |
| JSONV2 | 7858 | 230.60 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8174 | 221.56 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15775 | 114.86 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1220 | 1485.72 MB/s | 144 | 10 | 12.9× |
| Lightning | 1224 | 1480.10 MB/s | 144 | 10 | 12.9× |
| LightningDestructive | 1271 | 1425.48 MB/s | 144 | 10 | 12.4× |
| Easyjson | 3037 | 596.70 MB/s | 144 | 10 | 5.2× |
| Goccy | 3104 | 583.75 MB/s | 2600 | 5 | 5.1× |
| SonicFastest | 6278 | 288.63 MB/s | 3368 | 40 | 2.5× |
| Sonic | 6442 | 281.28 MB/s | 3368 | 40 | 2.4× |
| JSONV2 | 7681 | 235.89 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 8258 | 219.30 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15746 | 115.08 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 706 | 699.32 MB/s | 160 | 1 | 8.6× |
| LightningDestructive | 711 | 694.45 MB/s | 160 | 1 | 8.6× |
| SonicFastest | 1214 | 406.84 MB/s | 1075 | 8 | 5.0× |
| Sonic | 1218 | 405.52 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1386 | 355.81 MB/s | 1296 | 26 | 4.4× |
| LightningArena | 1564 | 315.82 MB/s | 4096 | 1 | 3.9× |
| Easyjson | 2263 | 218.28 MB/s | 448 | 3 | 2.7× |
| Goccy | 2399 | 205.91 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3140 | 157.31 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6097 | 81.02 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 432 | 531.97 MB/s | 160 | 1 | 10.0× |
| LightningDestructive | 436 | 527.23 MB/s | 160 | 1 | 9.9× |
| Sonic | 857 | 268.48 MB/s | 801 | 8 | 5.0× |
| SonicFastest | 859 | 267.81 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1176 | 194.75 MB/s | 1296 | 26 | 3.7× |
| LightningArena | 1253 | 183.53 MB/s | 4096 | 1 | 3.5× |
| Easyjson | 1514 | 151.88 MB/s | 448 | 3 | 2.9× |
| Goccy | 1647 | 139.67 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2366 | 97.22 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4324 | 53.19 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 61946 | 1051.42 MB/s | 103441 | 103 | 10.7× |
| Lightning | 62331 | 1044.94 MB/s | 103441 | 103 | 10.7× |
| LightningDestructive | 64261 | 1013.55 MB/s | 97220 | 98 | 10.4× |
| Sonic | 154039 | 422.83 MB/s | 235972 | 65 | 4.3× |
| SonicFastest | 154147 | 422.53 MB/s | 236019 | 65 | 4.3× |
| LightningDecodeAny | 179275 | 297.47 MB/s | 180048 | 3245 | 3.7× |
| Goccy | 192334 | 338.64 MB/s | 228738 | 134 | 3.5× |
| JSONV2 | 256696 | 253.73 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 665422 | 97.88 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2499363 | 776.39 MB/s | 2864592 | 1380 | 10.5× |
| LightningArena | 2562346 | 757.30 MB/s | 2864593 | 1380 | 10.2× |
| Lightning | 2575898 | 753.32 MB/s | 2864594 | 1380 | 10.1× |
| Goccy | 5024548 | 386.20 MB/s | 4062802 | 13509 | 5.2× |
| SonicFastest | 5049705 | 384.27 MB/s | 4878899 | 1736 | 5.2× |
| Sonic | 5066584 | 382.99 MB/s | 4879674 | 1736 | 5.2× |
| Easyjson | 7576511 | 256.12 MB/s | 3871264 | 15043 | 3.4× |
| LightningDecodeAny | 8880443 | 218.51 MB/s | 7063040 | 218633 | 2.9× |
| JSONV2 | 11513194 | 168.54 MB/s | 3237203 | 13947 | 2.3× |
| Stdlib | 26126369 | 74.27 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 974706 | 3414.19 MB/s | 351704 | 1286 | 26.2× |
| LightningArena | 1491781 | 2230.78 MB/s | 2488906 | 2995 | 17.1× |
| Lightning | 1508387 | 2206.22 MB/s | 2488906 | 2995 | 17.0× |
| SonicFastest | 2326307 | 1430.52 MB/s | 5896222 | 4263 | 11.0× |
| Sonic | 2327095 | 1430.04 MB/s | 5896114 | 4263 | 11.0× |
| LightningDecodeAny | 3200121 | 960.52 MB/s | 4876912 | 56892 | 8.0× |
| Goccy | 5570984 | 597.35 MB/s | 3948914 | 3816 | 4.6× |
| JSONV2 | 7761646 | 428.75 MB/s | 5364510 | 13243 | 3.3× |
| Stdlib | 25577656 | 130.11 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 208965 | 1054.47 MB/s | 135872 | 226 | 11.0× |
| Lightning | 210215 | 1048.19 MB/s | 135872 | 226 | 11.0× |
| LightningDestructive | 220237 | 1000.50 MB/s | 135872 | 226 | 10.5× |
| Goccy | 453715 | 485.65 MB/s | 363916 | 1066 | 5.1× |
| SonicFastest | 513739 | 428.91 MB/s | 350897 | 262 | 4.5× |
| Sonic | 514996 | 427.86 MB/s | 350940 | 262 | 4.5× |
| Easyjson | 570869 | 385.98 MB/s | 130512 | 245 | 4.0× |
| JSONV2 | 715011 | 308.17 MB/s | 129746 | 470 | 3.2× |
| LightningDecodeAny | 922537 | 117.41 MB/s | 897217 | 11703 | 2.5× |
| Stdlib | 2308877 | 95.43 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11901832 | 680.57 MB/s | 11845074 | 20816 | 8.6× |
| LightningArena | 11986712 | 675.75 MB/s | 11845073 | 20816 | 8.5× |
| Lightning | 12012171 | 674.32 MB/s | 11845074 | 20816 | 8.5× |
| SonicFastest | 20337048 | 398.29 MB/s | 19854908 | 41640 | 5.0× |
| Sonic | 20701631 | 391.28 MB/s | 19852495 | 41640 | 4.9× |
| Goccy | 25603097 | 316.37 MB/s | 18830181 | 107154 | 4.0× |
| Easyjson | 33940530 | 238.65 MB/s | 15059617 | 41643 | 3.0× |
| LightningDecodeAny | 37668706 | 138.13 MB/s | 46279351 | 747112 | 2.7× |
| JSONV2 | 46052309 | 175.89 MB/s | 15233744 | 78972 | 2.2× |
| Stdlib | 101816519 | 79.56 MB/s | 15665073 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5355789 | 557.05 MB/s | 3764713 | 1504 | 9.7× |
| LightningDestructive | 5485024 | 543.93 MB/s | 3758856 | 29356 | 9.4× |
| Lightning | 5670302 | 526.16 MB/s | 3758856 | 29356 | 9.1× |
| Sonic | 9095103 | 328.03 MB/s | 9129439 | 57804 | 5.7× |
| SonicFastest | 9154114 | 325.92 MB/s | 9129863 | 57804 | 5.7× |
| LightningDecodeAny | 17615450 | 104.12 MB/s | 23982580 | 351152 | 2.9× |
| Easyjson | 18003610 | 165.71 MB/s | 9479441 | 30115 | 2.9× |
| Goccy | 18038290 | 165.40 MB/s | 9857695 | 273619 | 2.9× |
| JSONV2 | 24448795 | 122.03 MB/s | 9257074 | 86278 | 2.1× |
| Stdlib | 51808627 | 57.59 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1291925 | 560.09 MB/s | 911393 | 30 | 9.9× |
| LightningDestructive | 1294045 | 559.17 MB/s | 907601 | 3618 | 9.8× |
| Lightning | 1355355 | 533.88 MB/s | 907595 | 3618 | 9.4× |
| Sonic | 2165902 | 334.09 MB/s | 2367981 | 3683 | 5.9× |
| SonicFastest | 2174192 | 332.81 MB/s | 2367606 | 3683 | 5.9× |
| Easyjson | 5147609 | 140.57 MB/s | 2847906 | 3698 | 2.5× |
| Goccy | 5316830 | 136.10 MB/s | 2715715 | 80267 | 2.4× |
| LightningDecodeAny | 5353570 | 121.52 MB/s | 6500462 | 76546 | 2.4× |
| JSONV2 | 6481999 | 111.63 MB/s | 2704707 | 7318 | 2.0× |
| Stdlib | 12738682 | 56.80 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1887275 | 835.78 MB/s | 911392 | 30 | 9.4× |
| Lightning | 1935566 | 814.93 MB/s | 907595 | 3618 | 9.2× |
| LightningDestructive | 1954175 | 807.17 MB/s | 907600 | 3618 | 9.1× |
| Sonic | 2419425 | 651.95 MB/s | 3220481 | 3683 | 7.3× |
| SonicFastest | 2421066 | 651.51 MB/s | 3221354 | 3683 | 7.3× |
| LightningDecodeAny | 4673911 | 161.19 MB/s | 6500458 | 76546 | 3.8× |
| Easyjson | 6262055 | 251.89 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 6552229 | 240.74 MB/s | 3491680 | 80261 | 2.7× |
| JSONV2 | 7070687 | 223.08 MB/s | 2704556 | 7318 | 2.5× |
| Stdlib | 17724769 | 88.99 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 217367 | 690.65 MB/s | 81920 | 1 | 9.3× |
| Lightning | 221354 | 678.21 MB/s | 81920 | 1 | 9.1× |
| LightningDestructive | 224213 | 669.56 MB/s | 81920 | 1 | 9.0× |
| Sonic | 417074 | 359.95 MB/s | 408373 | 16 | 4.9× |
| SonicFastest | 427121 | 351.48 MB/s | 409402 | 16 | 4.7× |
| LightningDecodeAny | 569823 | 263.45 MB/s | 745764 | 10016 | 3.6× |
| Goccy | 1041674 | 144.12 MB/s | 331764 | 10005 | 1.9× |
| JSONV2 | 1151085 | 130.42 MB/s | 357724 | 20 | 1.8× |
| Stdlib | 2025320 | 74.12 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 30131 | 933.17 MB/s | 29216 | 103 | 11.0× |
| Lightning | 30656 | 917.18 MB/s | 29216 | 103 | 10.8× |
| LightningDestructive | 31588 | 890.11 MB/s | 29088 | 101 | 10.5× |
| SonicFastest | 69690 | 403.46 MB/s | 59411 | 83 | 4.8× |
| Sonic | 70525 | 398.68 MB/s | 59458 | 83 | 4.7× |
| Easyjson | 72306 | 388.86 MB/s | 32304 | 138 | 4.6× |
| Goccy | 78132 | 359.86 MB/s | 59239 | 188 | 4.2× |
| JSONV2 | 130491 | 215.47 MB/s | 36896 | 242 | 2.5× |
| LightningDecodeAny | 147859 | 190.16 MB/s | 140576 | 2643 | 2.2× |
| Stdlib | 331392 | 84.85 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1891 | 1231.40 MB/s | 32 | 1 | 13.1× |
| Lightning | 1911 | 1218.18 MB/s | 32 | 1 | 13.0× |
| LightningDestructive | 2124 | 1095.81 MB/s | 32 | 1 | 11.7× |
| Goccy | 4644 | 501.25 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 4887 | 476.38 MB/s | 192 | 2 | 5.1× |
| Sonic | 6074 | 383.29 MB/s | 3708 | 4 | 4.1× |
| SonicFastest | 6081 | 382.86 MB/s | 3708 | 4 | 4.1× |
| JSONV2 | 7874 | 295.67 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 9463 | 178.05 MB/s | 10200 | 195 | 2.6× |
| Stdlib | 24842 | 93.71 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 208 | 910.52 MB/s | 0 | 0 | 12.6× |
| Lightning | 208 | 909.36 MB/s | 0 | 0 | 12.6× |
| LightningDestructive | 211 | 894.76 MB/s | 0 | 0 | 12.4× |
| Goccy | 419 | 451.34 MB/s | 304 | 2 | 6.2× |
| Easyjson | 525 | 360.27 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 765 | 247.15 MB/s | 341 | 3 | 3.4× |
| Sonic | 767 | 246.54 MB/s | 341 | 3 | 3.4× |
| JSONV2 | 959 | 197.01 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1172 | 114.31 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 2617 | 72.23 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1407 | 1557.06 MB/s | 0 | 0 | 13.1× |
| LightningArena | 1411 | 1553.25 MB/s | 0 | 0 | 13.1× |
| LightningDestructive | 1456 | 1505.06 MB/s | 0 | 0 | 12.7× |
| Easyjson | 3312 | 661.49 MB/s | 24 | 1 | 5.6× |
| Goccy | 3702 | 591.84 MB/s | 2864 | 4 | 5.0× |
| SonicFastest | 6620 | 330.98 MB/s | 3600 | 38 | 2.8× |
| Sonic | 6789 | 322.73 MB/s | 3598 | 38 | 2.7× |
| JSONV2 | 7946 | 275.73 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8211 | 220.55 MB/s | 7552 | 158 | 2.3× |
| Stdlib | 18489 | 118.50 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 577548 | 883.87 MB/s | 457537 | 1009 | 11.6× |
| Lightning | 584764 | 872.96 MB/s | 457537 | 1009 | 11.5× |
| LightningArena | 591421 | 863.13 MB/s | 457538 | 1009 | 11.3× |
| Goccy | 1235962 | 413.02 MB/s | 1135310 | 5006 | 5.4× |
| Sonic | 1416836 | 360.29 MB/s | 1308261 | 2014 | 4.7× |
| SonicFastest | 1424584 | 358.33 MB/s | 1308034 | 2014 | 4.7× |
| Easyjson | 1594942 | 320.06 MB/s | 863780 | 3012 | 4.2× |
| JSONV2 | 3059742 | 166.84 MB/s | 1075950 | 12645 | 2.2× |
| LightningDecodeAny | 3333083 | 138.45 MB/s | 2950650 | 64018 | 2.0× |
| Stdlib | 6709674 | 76.08 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 611 | 32379.04 MB/s | 0 | 0 | 254.9× |
| LightningArena | 611 | 32375.88 MB/s | 0 | 0 | 254.9× |
| LightningDestructive | 866 | 22837.30 MB/s | 0 | 0 | 179.8× |
| SonicFastest | 6390 | 3096.97 MB/s | 21120 | 3 | 24.4× |
| Goccy | 25387 | 779.50 MB/s | 20492 | 2 | 6.1× |
| Sonic | 29109 | 679.83 MB/s | 20611 | 3 | 5.4× |
| JSONV2 | 36318 | 544.88 MB/s | 8 | 1 | 4.3× |
| LightningDecodeAny | 90398 | 218.90 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 126503 | 156.43 MB/s | 0 | 0 | 1.2× |
| Stdlib | 155805 | 127.01 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 2126 | 8524.63 MB/s | 432 | 2 | 62.8× |
| Lightning | 2167 | 8364.43 MB/s | 432 | 2 | 61.6× |
| LightningDestructive | 2304 | 7866.09 MB/s | 0 | 0 | 58.0× |
| Easyjson | 4479 | 4046.08 MB/s | 432 | 2 | 29.8× |
| SonicFastest | 8979 | 2018.50 MB/s | 20447 | 5 | 14.9× |
| Sonic | 8987 | 2016.72 MB/s | 20461 | 5 | 14.9× |
| LightningDecodeAny | 17784 | 1005.51 MB/s | 29088 | 191 | 7.5× |
| Goccy | 25863 | 700.78 MB/s | 19460 | 2 | 5.2× |
| JSONV2 | 48055 | 377.15 MB/s | 16500 | 50 | 2.8× |
| Stdlib | 133594 | 135.66 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2136531 | 940.07 MB/s | 3089565 | 6821 | 9.7× |
| LightningArena | 2295957 | 874.80 MB/s | 3094370 | 6703 | 9.0× |
| Lightning | 2306365 | 870.85 MB/s | 3091277 | 6827 | 9.0× |
| Goccy | 4486431 | 447.68 MB/s | 5410860 | 15831 | 4.6× |
| SonicFastest | 4823978 | 416.36 MB/s | 5152318 | 7085 | 4.3× |
| Sonic | 4953327 | 405.48 MB/s | 5152637 | 7085 | 4.2× |
| Easyjson | 4987940 | 402.67 MB/s | 2981484 | 7439 | 4.2× |
| LightningDecodeAny | 6507382 | 175.54 MB/s | 8503513 | 134008 | 3.2× |
| JSONV2 | 6983223 | 287.62 MB/s | 3173675 | 14563 | 3.0× |
| Stdlib | 20767372 | 96.71 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 806 | 681.10 MB/s | 480 | 1 | 7.6× |
| Lightning | 809 | 678.64 MB/s | 480 | 1 | 7.5× |
| LightningDestructive | 818 | 671.06 MB/s | 480 | 1 | 7.4× |
| LightningDecodeAny | 1707 | 321.05 MB/s | 2021 | 46 | 3.6× |
| Easyjson | 1915 | 286.69 MB/s | 1616 | 5 | 3.2× |
| SonicFastest | 2054 | 267.25 MB/s | 2263 | 8 | 3.0× |
| Sonic | 2117 | 259.38 MB/s | 2262 | 8 | 2.9× |
| Goccy | 2904 | 189.02 MB/s | 2129 | 43 | 2.1× |
| JSONV2 | 3001 | 182.96 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 6093 | 90.10 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 463651 | 1362.05 MB/s | 402728 | 545 | 13.4× |
| Lightning | 540391 | 1168.62 MB/s | 451257 | 857 | 11.5× |
| LightningArena | 540527 | 1168.33 MB/s | 453016 | 712 | 11.5× |
| SonicFastest | 1094171 | 577.16 MB/s | 1067885 | 814 | 5.7× |
| Sonic | 1094636 | 576.92 MB/s | 1067928 | 814 | 5.7× |
| Easyjson | 1217911 | 518.52 MB/s | 422504 | 936 | 5.1× |
| Goccy | 1382914 | 456.65 MB/s | 990554 | 1201 | 4.5× |
| JSONV2 | 2171401 | 290.83 MB/s | 571590 | 3144 | 2.9× |
| LightningDecodeAny | 2315263 | 201.66 MB/s | 2076504 | 30126 | 2.7× |
| Stdlib | 6222894 | 101.48 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 617022 | 911.49 MB/s | 546569 | 429 | 9.5× |
| Lightning | 805197 | 698.47 MB/s | 769937 | 1235 | 7.2× |
| LightningArena | 807910 | 696.13 MB/s | 771665 | 1088 | 7.2× |
| Sonic | 1344644 | 418.26 MB/s | 1347093 | 1185 | 4.3× |
| SonicFastest | 1351687 | 416.08 MB/s | 1347435 | 1185 | 4.3× |
| Goccy | 1551586 | 362.47 MB/s | 1040017 | 1029 | 3.8× |
| Easyjson | 1950052 | 288.41 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2780698 | 202.25 MB/s | 2180441 | 30126 | 2.1× |
| JSONV2 | 3017920 | 186.36 MB/s | 927406 | 3482 | 1.9× |
| Stdlib | 5831383 | 96.45 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 584417 | 912.32 MB/s | 333416 | 2084 | 10.7× |
| LightningArena | 657735 | 810.63 MB/s | 368224 | 2293 | 9.5× |
| Lightning | 663041 | 804.14 MB/s | 368224 | 2293 | 9.4× |
| Easyjson | 1215580 | 438.62 MB/s | 428362 | 3273 | 5.1× |
| Sonic | 1361149 | 391.71 MB/s | 980816 | 3082 | 4.6× |
| SonicFastest | 1389216 | 383.80 MB/s | 980815 | 3082 | 4.5× |
| Goccy | 1525378 | 349.54 MB/s | 1167061 | 5408 | 4.1× |
| JSONV2 | 2623245 | 203.25 MB/s | 745425 | 13288 | 2.4× |
| LightningDecodeAny | 3305464 | 161.30 MB/s | 2992875 | 50076 | 1.9× |
| Stdlib | 6244028 | 85.39 MB/s | 798692 | 17133 | 1.0× |
