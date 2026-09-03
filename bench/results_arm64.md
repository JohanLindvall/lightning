# JSON Deserialization Benchmarks

- generated 2026-09-03T06:45:13Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 83946 | 1516.15 MB/s | 49760 | 3 | 13.0× |
| LightningArena | 83953 | 1516.02 MB/s | 49760 | 3 | 13.0× |
| LightningDestructive | 84453 | 1507.04 MB/s | 49280 | 2 | 12.9× |
| Sonic | 193866 | 656.51 MB/s | 213700 | 10 | 5.6× |
| SonicFastest | 194971 | 652.79 MB/s | 213972 | 10 | 5.6× |
| Goccy | 203298 | 626.05 MB/s | 225866 | 884 | 5.4× |
| Easyjson | 213489 | 596.17 MB/s | 122864 | 14 | 5.1× |
| LightningDecodeAny | 423938 | 223.27 MB/s | 463409 | 9708 | 2.6× |
| JSONV2 | 426039 | 298.74 MB/s | 195118 | 1805 | 2.6× |
| Stdlib | 1090581 | 116.70 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2669392 | 843.28 MB/s | 2532848 | 1143 | 9.9× |
| LightningArena | 2678532 | 840.40 MB/s | 2532849 | 1143 | 9.9× |
| Lightning | 2692330 | 836.10 MB/s | 2532850 | 1143 | 9.9× |
| SonicFastest | 4661440 | 482.91 MB/s | 15240548 | 970 | 5.7× |
| Sonic | 4861842 | 463.00 MB/s | 15249562 | 970 | 5.5× |
| Goccy | 10481979 | 214.75 MB/s | 4126855 | 56532 | 2.5× |
| Easyjson | 11056222 | 203.60 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11465028 | 196.34 MB/s | 19380210 | 223896 | 2.3× |
| JSONV2 | 16309135 | 138.02 MB/s | 3123215 | 3083 | 1.6× |
| Stdlib | 26523050 | 84.87 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 372400 | 726.11 MB/s | 397297 | 567 | 9.2× |
| Lightning | 372844 | 725.24 MB/s | 397296 | 567 | 9.2× |
| LightningDestructive | 373678 | 723.62 MB/s | 397296 | 567 | 9.2× |
| Sonic | 638830 | 423.28 MB/s | 483753 | 968 | 5.4× |
| SonicFastest | 641982 | 421.20 MB/s | 497093 | 968 | 5.3× |
| Easyjson | 1380213 | 195.91 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1412425 | 191.45 MB/s | 543390 | 8122 | 2.4× |
| LightningDecodeAny | 1588863 | 170.19 MB/s | 2543883 | 29687 | 2.2× |
| JSONV2 | 2117026 | 127.73 MB/s | 348151 | 1628 | 1.6× |
| Stdlib | 3431387 | 78.80 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 947976 | 1821.99 MB/s | 768416 | 2440 | 13.9× |
| LightningDestructive | 948467 | 1821.05 MB/s | 765560 | 2798 | 13.9× |
| Lightning | 951738 | 1814.79 MB/s | 765601 | 2799 | 13.9× |
| Sonic | 2088701 | 826.93 MB/s | 2655650 | 4020 | 6.3× |
| SonicFastest | 2093229 | 825.14 MB/s | 2731859 | 4020 | 6.3× |
| Goccy | 2434102 | 709.59 MB/s | 2582879 | 14604 | 5.4× |
| Easyjson | 4212107 | 410.06 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4238644 | 407.49 MB/s | 1011636 | 7594 | 3.1× |
| LightningDecodeAny | 4256562 | 117.54 MB/s | 4953694 | 76576 | 3.1× |
| Stdlib | 13207630 | 130.77 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 866 | 2091.67 MB/s | 0 | 0 | 16.1× |
| LightningArena | 873 | 2075.87 MB/s | 0 | 0 | 15.9× |
| LightningDestructive | 885 | 2046.49 MB/s | 0 | 0 | 15.7× |
| Easyjson | 2535 | 714.80 MB/s | 24 | 1 | 5.5× |
| Goccy | 2908 | 623.18 MB/s | 2608 | 4 | 4.8× |
| Sonic | 6056 | 299.20 MB/s | 3851 | 40 | 2.3× |
| SonicFastest | 6079 | 298.10 MB/s | 3905 | 40 | 2.3× |
| JSONV2 | 7781 | 232.89 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8139 | 222.51 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13921 | 130.16 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 895 | 2024.36 MB/s | 0 | 0 | 15.5× |
| Lightning | 896 | 2021.75 MB/s | 0 | 0 | 15.5× |
| LightningDestructive | 928 | 1951.85 MB/s | 0 | 0 | 15.0× |
| Easyjson | 2539 | 713.66 MB/s | 24 | 1 | 5.5× |
| Goccy | 2853 | 635.06 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 5971 | 303.45 MB/s | 3753 | 40 | 2.3× |
| Sonic | 5991 | 302.46 MB/s | 3776 | 40 | 2.3× |
| LightningDecodeAny | 7793 | 232.40 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7844 | 231.01 MB/s | 640 | 6 | 1.8× |
| Stdlib | 13916 | 130.21 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1073 | 1689.32 MB/s | 144 | 10 | 13.0× |
| LightningArena | 1086 | 1668.33 MB/s | 144 | 10 | 12.8× |
| LightningDestructive | 1131 | 1602.48 MB/s | 144 | 10 | 12.3× |
| Easyjson | 2747 | 659.58 MB/s | 144 | 10 | 5.1× |
| Goccy | 2919 | 620.84 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6173 | 293.54 MB/s | 3766 | 42 | 2.3× |
| SonicFastest | 6220 | 291.32 MB/s | 3837 | 42 | 2.2× |
| LightningDecodeAny | 7852 | 230.64 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 8039 | 225.41 MB/s | 632 | 7 | 1.7× |
| Stdlib | 13909 | 130.28 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 624 | 791.50 MB/s | 160 | 1 | 8.8× |
| LightningDestructive | 626 | 788.53 MB/s | 160 | 1 | 8.7× |
| SonicFastest | 1231 | 401.41 MB/s | 998 | 6 | 4.4× |
| Sonic | 1232 | 401.10 MB/s | 982 | 6 | 4.4× |
| LightningArena | 1266 | 390.28 MB/s | 4096 | 1 | 4.3× |
| LightningDecodeAny | 1287 | 382.97 MB/s | 1296 | 26 | 4.3× |
| Easyjson | 2240 | 220.58 MB/s | 448 | 3 | 2.4× |
| Goccy | 2392 | 206.49 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3248 | 152.11 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5477 | 90.20 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 372 | 617.82 MB/s | 160 | 1 | 10.9× |
| Lightning | 378 | 608.92 MB/s | 160 | 1 | 10.7× |
| Sonic | 882 | 260.87 MB/s | 652 | 6 | 4.6× |
| SonicFastest | 886 | 259.65 MB/s | 660 | 6 | 4.6× |
| LightningArena | 1033 | 222.62 MB/s | 4096 | 1 | 3.9× |
| LightningDecodeAny | 1118 | 204.85 MB/s | 1296 | 26 | 3.6× |
| Easyjson | 1379 | 166.75 MB/s | 448 | 3 | 2.9× |
| Goccy | 1573 | 146.26 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2450 | 93.89 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4060 | 56.66 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 51152 | 1273.31 MB/s | 97220 | 98 | 10.7× |
| LightningArena | 52658 | 1236.88 MB/s | 103440 | 103 | 10.4× |
| Lightning | 52785 | 1233.91 MB/s | 103440 | 103 | 10.4× |
| Sonic | 100685 | 646.89 MB/s | 155945 | 75 | 5.5× |
| SonicFastest | 100935 | 645.29 MB/s | 155818 | 75 | 5.4× |
| Goccy | 150009 | 434.19 MB/s | 229490 | 134 | 3.7× |
| LightningDecodeAny | 175554 | 303.78 MB/s | 180048 | 3245 | 3.1× |
| JSONV2 | 231054 | 281.89 MB/s | 206653 | 607 | 2.4× |
| Stdlib | 549219 | 118.59 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2096248 | 925.69 MB/s | 2864592 | 1380 | 11.0× |
| Lightning | 2204544 | 880.21 MB/s | 2864594 | 1380 | 10.5× |
| LightningArena | 2207361 | 879.09 MB/s | 2864593 | 1380 | 10.5× |
| Sonic | 4600814 | 421.77 MB/s | 14608574 | 1407 | 5.0× |
| Goccy | 4886077 | 397.14 MB/s | 4065090 | 13510 | 4.7× |
| SonicFastest | 5011494 | 387.20 MB/s | 14608622 | 1407 | 4.6× |
| Easyjson | 7410048 | 261.87 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 8933901 | 217.20 MB/s | 7063039 | 218633 | 2.6× |
| JSONV2 | 11387829 | 170.40 MB/s | 3237223 | 13947 | 2.0× |
| Stdlib | 23153306 | 83.81 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 891723 | 3731.91 MB/s | 351704 | 1286 | 23.6× |
| Lightning | 1651821 | 2014.64 MB/s | 2488904 | 2995 | 12.8× |
| LightningArena | 1670415 | 1992.22 MB/s | 2488905 | 2995 | 12.6× |
| Sonic | 2897261 | 1148.61 MB/s | 6461500 | 4248 | 7.3× |
| SonicFastest | 2908330 | 1144.24 MB/s | 6469121 | 4248 | 7.2× |
| LightningDecodeAny | 3470421 | 885.70 MB/s | 4876911 | 56892 | 6.1× |
| Goccy | 4762108 | 698.81 MB/s | 3948908 | 3816 | 4.4× |
| JSONV2 | 7640564 | 435.55 MB/s | 5364516 | 13243 | 2.8× |
| Stdlib | 21063181 | 157.99 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 180199 | 1222.79 MB/s | 135872 | 226 | 11.2× |
| LightningArena | 180532 | 1220.53 MB/s | 135872 | 226 | 11.2× |
| LightningDestructive | 181391 | 1214.76 MB/s | 135872 | 226 | 11.2× |
| SonicFastest | 376392 | 585.42 MB/s | 297867 | 398 | 5.4× |
| Sonic | 377430 | 583.81 MB/s | 297281 | 398 | 5.4× |
| Goccy | 433203 | 508.64 MB/s | 364171 | 1067 | 4.7× |
| Easyjson | 546474 | 403.21 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 738910 | 298.20 MB/s | 129743 | 470 | 2.7× |
| LightningDecodeAny | 837269 | 129.36 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2022664 | 108.94 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9329437 | 868.22 MB/s | 11845073 | 20816 | 9.5× |
| Lightning | 9534164 | 849.58 MB/s | 11845072 | 20816 | 9.3× |
| LightningArena | 9559579 | 847.32 MB/s | 11845073 | 20816 | 9.2× |
| Sonic | 17016085 | 476.02 MB/s | 70902295 | 40014 | 5.2× |
| SonicFastest | 17099570 | 473.70 MB/s | 70873026 | 40014 | 5.2× |
| Goccy | 23615267 | 343.00 MB/s | 17250289 | 107149 | 3.7× |
| Easyjson | 31177075 | 259.81 MB/s | 15059618 | 41643 | 2.8× |
| LightningDecodeAny | 35919957 | 144.85 MB/s | 46279373 | 747112 | 2.5× |
| JSONV2 | 43935638 | 184.36 MB/s | 15233741 | 78972 | 2.0× |
| Stdlib | 88315879 | 91.72 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4267991 | 699.03 MB/s | 3764712 | 1504 | 10.9× |
| LightningDestructive | 4548052 | 655.99 MB/s | 3758857 | 29356 | 10.2× |
| Lightning | 4725182 | 631.40 MB/s | 3758860 | 29356 | 9.9× |
| Sonic | 8645516 | 345.09 MB/s | 26501084 | 56760 | 5.4× |
| SonicFastest | 8719947 | 342.14 MB/s | 26503913 | 56760 | 5.3× |
| LightningDecodeAny | 15863828 | 115.62 MB/s | 23982581 | 351152 | 2.9× |
| Easyjson | 16505642 | 180.75 MB/s | 9479440 | 30115 | 2.8× |
| Goccy | 16723177 | 178.40 MB/s | 10488251 | 273644 | 2.8× |
| JSONV2 | 25237658 | 118.21 MB/s | 9257159 | 86278 | 1.8× |
| Stdlib | 46602082 | 64.02 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 919936 | 786.57 MB/s | 907601 | 3618 | 12.5× |
| LightningArena | 933185 | 775.41 MB/s | 911396 | 30 | 12.4× |
| Lightning | 990302 | 730.68 MB/s | 907599 | 3618 | 11.6× |
| Sonic | 1809418 | 399.91 MB/s | 3201811 | 7226 | 6.4× |
| SonicFastest | 1814027 | 398.89 MB/s | 3199453 | 7226 | 6.4× |
| LightningDecodeAny | 4053389 | 160.50 MB/s | 6500457 | 76546 | 2.8× |
| Easyjson | 4275352 | 169.25 MB/s | 2847906 | 3698 | 2.7× |
| Goccy | 4828142 | 149.87 MB/s | 2764333 | 80271 | 2.4× |
| JSONV2 | 5850094 | 123.69 MB/s | 2704627 | 7318 | 2.0× |
| Stdlib | 11525097 | 62.78 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1365436 | 1155.20 MB/s | 907601 | 3618 | 11.4× |
| LightningArena | 1376035 | 1146.30 MB/s | 911394 | 30 | 11.3× |
| Lightning | 1423551 | 1108.04 MB/s | 907595 | 3618 | 10.9× |
| Sonic | 2268885 | 695.21 MB/s | 5787433 | 7226 | 6.8× |
| SonicFastest | 2272096 | 694.23 MB/s | 5790975 | 7226 | 6.8× |
| LightningDecodeAny | 3645435 | 206.67 MB/s | 6500458 | 76546 | 4.3× |
| Easyjson | 5604624 | 281.44 MB/s | 2847907 | 3698 | 2.8× |
| Goccy | 5683776 | 277.52 MB/s | 3582165 | 80267 | 2.7× |
| JSONV2 | 6440339 | 244.92 MB/s | 2704592 | 7318 | 2.4× |
| Stdlib | 15540028 | 101.50 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 156908 | 956.77 MB/s | 81920 | 1 | 11.8× |
| LightningArena | 157003 | 956.19 MB/s | 81920 | 1 | 11.8× |
| Lightning | 157224 | 954.84 MB/s | 81920 | 1 | 11.8× |
| Sonic | 275238 | 545.43 MB/s | 257456 | 6 | 6.8× |
| SonicFastest | 280314 | 535.56 MB/s | 270053 | 6 | 6.6× |
| LightningDecodeAny | 426918 | 351.64 MB/s | 745764 | 10016 | 4.4× |
| Goccy | 865353 | 173.48 MB/s | 324527 | 10004 | 2.1× |
| JSONV2 | 1068001 | 140.57 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1858831 | 80.76 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 26992 | 1041.69 MB/s | 29216 | 103 | 11.1× |
| Lightning | 27079 | 1038.31 MB/s | 29216 | 103 | 11.1× |
| LightningDestructive | 27323 | 1029.05 MB/s | 29088 | 101 | 11.0× |
| Sonic | 62793 | 447.77 MB/s | 46521 | 103 | 4.8× |
| SonicFastest | 63056 | 445.91 MB/s | 46978 | 103 | 4.8× |
| Easyjson | 68170 | 412.45 MB/s | 32304 | 138 | 4.4× |
| Goccy | 70598 | 398.27 MB/s | 59184 | 188 | 4.3× |
| JSONV2 | 134852 | 208.50 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 144973 | 193.95 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 300316 | 93.62 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1511 | 1540.28 MB/s | 32 | 1 | 14.8× |
| LightningArena | 1514 | 1537.75 MB/s | 32 | 1 | 14.8× |
| LightningDestructive | 1595 | 1459.59 MB/s | 32 | 1 | 14.1× |
| Easyjson | 4238 | 549.37 MB/s | 192 | 2 | 5.3× |
| Goccy | 4343 | 536.00 MB/s | 3649 | 4 | 5.2× |
| Sonic | 5209 | 446.88 MB/s | 4332 | 6 | 4.3× |
| SonicFastest | 5238 | 444.46 MB/s | 4366 | 6 | 4.3× |
| JSONV2 | 8508 | 273.62 MB/s | 1000 | 6 | 2.6× |
| LightningDecodeAny | 9491 | 177.54 MB/s | 10200 | 195 | 2.4× |
| Stdlib | 22429 | 103.79 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 179 | 1057.35 MB/s | 0 | 0 | 13.3× |
| Lightning | 179 | 1055.39 MB/s | 0 | 0 | 13.3× |
| LightningDestructive | 182 | 1041.22 MB/s | 0 | 0 | 13.1× |
| Goccy | 383 | 493.60 MB/s | 304 | 2 | 6.2× |
| Easyjson | 490 | 385.78 MB/s | 0 | 0 | 4.9× |
| Sonic | 770 | 245.45 MB/s | 519 | 4 | 3.1× |
| SonicFastest | 772 | 244.78 MB/s | 518 | 4 | 3.1× |
| JSONV2 | 1031 | 183.38 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1190 | 112.61 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2379 | 79.44 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1174 | 1866.70 MB/s | 0 | 0 | 13.5× |
| LightningArena | 1175 | 1865.19 MB/s | 0 | 0 | 13.4× |
| LightningDestructive | 1197 | 1830.39 MB/s | 0 | 0 | 13.2× |
| Easyjson | 3196 | 685.55 MB/s | 24 | 1 | 4.9× |
| Goccy | 3282 | 667.65 MB/s | 2864 | 4 | 4.8× |
| Sonic | 6437 | 340.40 MB/s | 4035 | 40 | 2.5× |
| SonicFastest | 6450 | 339.71 MB/s | 4057 | 40 | 2.4× |
| LightningDecodeAny | 7846 | 230.81 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 8151 | 268.79 MB/s | 640 | 6 | 1.9× |
| Stdlib | 15796 | 138.71 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 556328 | 917.58 MB/s | 457536 | 1009 | 10.7× |
| Lightning | 558233 | 914.45 MB/s | 457536 | 1009 | 10.7× |
| LightningArena | 561647 | 908.89 MB/s | 457536 | 1009 | 10.6× |
| Goccy | 1159388 | 440.30 MB/s | 1138239 | 5006 | 5.1× |
| Sonic | 1181651 | 432.00 MB/s | 908395 | 2006 | 5.0× |
| SonicFastest | 1183637 | 431.28 MB/s | 911266 | 2006 | 5.0× |
| Easyjson | 1533402 | 332.90 MB/s | 863777 | 3012 | 3.9× |
| LightningDecodeAny | 3197720 | 144.31 MB/s | 2950650 | 64018 | 1.9× |
| JSONV2 | 3243732 | 157.37 MB/s | 1076011 | 12646 | 1.8× |
| Stdlib | 5956531 | 85.70 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 482 | 41072.83 MB/s | 0 | 0 | 225.1× |
| LightningArena | 482 | 41057.60 MB/s | 0 | 0 | 225.0× |
| LightningDestructive | 496 | 39921.54 MB/s | 0 | 0 | 218.8× |
| Goccy | 20260 | 976.75 MB/s | 20491 | 2 | 5.4× |
| Sonic | 27739 | 713.40 MB/s | 22644 | 4 | 3.9× |
| SonicFastest | 27777 | 712.42 MB/s | 22954 | 4 | 3.9× |
| JSONV2 | 29802 | 664.02 MB/s | 8 | 1 | 3.6× |
| Easyjson | 82034 | 241.23 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 82194 | 240.75 MB/s | 116864 | 2015 | 1.3× |
| Stdlib | 108440 | 182.49 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1753 | 10337.90 MB/s | 0 | 0 | 58.7× |
| LightningArena | 1876 | 9661.84 MB/s | 432 | 2 | 54.9× |
| Lightning | 1878 | 9652.97 MB/s | 432 | 2 | 54.8× |
| Easyjson | 3962 | 4574.36 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 10439 | 1736.16 MB/s | 23042 | 6 | 9.9× |
| Sonic | 10558 | 1716.66 MB/s | 22992 | 6 | 9.7× |
| LightningDecodeAny | 16271 | 1099.04 MB/s | 29088 | 191 | 6.3× |
| Goccy | 17022 | 1064.77 MB/s | 19459 | 2 | 6.0× |
| JSONV2 | 46771 | 387.50 MB/s | 16500 | 50 | 2.2× |
| Stdlib | 102908 | 176.12 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2248387 | 893.30 MB/s | 3089564 | 6821 | 8.4× |
| LightningArena | 2327789 | 862.83 MB/s | 3094370 | 6703 | 8.1× |
| Lightning | 2344560 | 856.66 MB/s | 3091277 | 6827 | 8.0× |
| Goccy | 4548274 | 441.59 MB/s | 5412499 | 15831 | 4.1× |
| Sonic | 4754541 | 422.44 MB/s | 10865889 | 13683 | 4.0× |
| SonicFastest | 4788540 | 419.44 MB/s | 10845947 | 13683 | 3.9× |
| Easyjson | 5049312 | 397.78 MB/s | 2981480 | 7438 | 3.7× |
| LightningDecodeAny | 6897247 | 165.62 MB/s | 8503513 | 134008 | 2.7× |
| JSONV2 | 7198245 | 279.03 MB/s | 3173680 | 14563 | 2.6× |
| Stdlib | 18780886 | 106.94 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 881 | 622.89 MB/s | 480 | 1 | 6.4× |
| Lightning | 884 | 621.08 MB/s | 480 | 1 | 6.4× |
| LightningArena | 887 | 619.07 MB/s | 480 | 1 | 6.4× |
| LightningDecodeAny | 1543 | 355.05 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 2306 | 238.05 MB/s | 1616 | 5 | 2.5× |
| Sonic | 2611 | 210.24 MB/s | 1944 | 26 | 2.2× |
| SonicFastest | 2637 | 208.23 MB/s | 1986 | 26 | 2.2× |
| Goccy | 2952 | 186.01 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3437 | 159.74 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5684 | 96.59 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 421347 | 1498.80 MB/s | 402729 | 545 | 12.7× |
| LightningArena | 480114 | 1315.34 MB/s | 453017 | 712 | 11.2× |
| Lightning | 487123 | 1296.42 MB/s | 451257 | 857 | 11.0× |
| Sonic | 1045986 | 603.75 MB/s | 993944 | 1102 | 5.1× |
| SonicFastest | 1048307 | 602.41 MB/s | 994994 | 1102 | 5.1× |
| Easyjson | 1140147 | 553.89 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1191782 | 529.89 MB/s | 986450 | 1201 | 4.5× |
| JSONV2 | 2158678 | 292.55 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2352563 | 198.47 MB/s | 2076505 | 30126 | 2.3× |
| Stdlib | 5368317 | 117.64 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 550571 | 1021.50 MB/s | 546571 | 429 | 9.6× |
| Lightning | 718742 | 782.49 MB/s | 769938 | 1235 | 7.3× |
| LightningArena | 722632 | 778.28 MB/s | 771665 | 1088 | 7.3× |
| Sonic | 1056203 | 532.48 MB/s | 976584 | 1476 | 5.0× |
| SonicFastest | 1058521 | 531.31 MB/s | 970918 | 1476 | 5.0× |
| Goccy | 1383476 | 406.52 MB/s | 1041181 | 1030 | 3.8× |
| Easyjson | 1763231 | 318.96 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2630776 | 213.78 MB/s | 2180439 | 30126 | 2.0× |
| JSONV2 | 2832099 | 198.58 MB/s | 927439 | 3482 | 1.9× |
| Stdlib | 5266086 | 106.80 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 571558 | 932.85 MB/s | 333416 | 2084 | 9.5× |
| Lightning | 595831 | 894.85 MB/s | 368224 | 2293 | 9.1× |
| LightningArena | 600229 | 888.29 MB/s | 368224 | 2293 | 9.0× |
| Easyjson | 1097674 | 485.73 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1160660 | 459.37 MB/s | 1037610 | 4351 | 4.7× |
| Sonic | 1172659 | 454.67 MB/s | 1050202 | 4351 | 4.6× |
| Goccy | 1309533 | 407.15 MB/s | 1167246 | 5409 | 4.1× |
| JSONV2 | 2534422 | 210.37 MB/s | 745450 | 13288 | 2.1× |
| LightningDecodeAny | 3344165 | 159.44 MB/s | 2992879 | 50076 | 1.6× |
| Stdlib | 5412746 | 98.50 MB/s | 798693 | 17133 | 1.0× |
