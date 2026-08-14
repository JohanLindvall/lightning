# JSON Deserialization Benchmarks

- generated 2026-08-14T16:28:03Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 94350 | 1348.97 MB/s | 49760 | 3 | 11.7× |
| Lightning | 94743 | 1343.37 MB/s | 49760 | 3 | 11.7× |
| LightningDestructive | 95402 | 1334.09 MB/s | 49280 | 2 | 11.6× |
| SonicFastest | 184010 | 691.67 MB/s | 199744 | 10 | 6.0× |
| Sonic | 185587 | 685.80 MB/s | 203136 | 10 | 6.0× |
| Goccy | 195599 | 650.69 MB/s | 225378 | 884 | 5.7× |
| Easyjson | 211724 | 601.14 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 421842 | 301.71 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 430705 | 219.76 MB/s | 463409 | 9708 | 2.6× |
| Stdlib | 1105949 | 115.08 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3631083 | 619.94 MB/s | 2532849 | 1143 | 7.2× |
| LightningArena | 3662709 | 614.59 MB/s | 2532852 | 1143 | 7.1× |
| Lightning | 3666286 | 613.99 MB/s | 2532850 | 1143 | 7.1× |
| SonicFastest | 4510393 | 499.08 MB/s | 15240439 | 970 | 5.8× |
| Sonic | 4543040 | 495.49 MB/s | 15233729 | 970 | 5.7× |
| Goccy | 10427715 | 215.87 MB/s | 4120704 | 56532 | 2.5× |
| Easyjson | 10944189 | 205.68 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12537425 | 179.54 MB/s | 19380211 | 223896 | 2.1× |
| JSONV2 | 16192885 | 139.01 MB/s | 3123222 | 3083 | 1.6× |
| Stdlib | 26034561 | 86.46 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 470903 | 574.22 MB/s | 397296 | 567 | 7.1× |
| Lightning | 474130 | 570.31 MB/s | 397296 | 567 | 7.1× |
| LightningArena | 475779 | 568.34 MB/s | 397297 | 567 | 7.0× |
| Sonic | 633688 | 426.71 MB/s | 496505 | 968 | 5.3× |
| SonicFastest | 638176 | 423.71 MB/s | 516521 | 968 | 5.2× |
| Easyjson | 1401173 | 192.98 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1407633 | 192.10 MB/s | 543806 | 8123 | 2.4× |
| LightningDecodeAny | 1663931 | 162.51 MB/s | 2543876 | 29687 | 2.0× |
| JSONV2 | 2115367 | 127.83 MB/s | 348151 | 1628 | 1.6× |
| Stdlib | 3348036 | 80.76 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1132958 | 1524.51 MB/s | 765560 | 2798 | 11.8× |
| LightningArena | 1144440 | 1509.21 MB/s | 768417 | 2440 | 11.6× |
| Lightning | 1145482 | 1507.84 MB/s | 765602 | 2799 | 11.6× |
| Sonic | 2042809 | 845.50 MB/s | 2740200 | 4020 | 6.5× |
| SonicFastest | 2060937 | 838.07 MB/s | 2705692 | 4020 | 6.5× |
| Goccy | 2430883 | 710.53 MB/s | 2582554 | 14605 | 5.5× |
| Easyjson | 4235368 | 407.80 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4267530 | 404.73 MB/s | 1011634 | 7594 | 3.1× |
| LightningDecodeAny | 4352197 | 114.95 MB/s | 4953694 | 76576 | 3.1× |
| Stdlib | 13331535 | 129.56 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 998 | 1815.39 MB/s | 0 | 0 | 14.1× |
| LightningArena | 1001 | 1810.22 MB/s | 0 | 0 | 14.0× |
| LightningDestructive | 1013 | 1788.92 MB/s | 0 | 0 | 13.9× |
| Easyjson | 2528 | 716.69 MB/s | 24 | 1 | 5.6× |
| Goccy | 2860 | 633.47 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 5958 | 304.13 MB/s | 3765 | 40 | 2.4× |
| Sonic | 5976 | 303.22 MB/s | 3726 | 40 | 2.4× |
| JSONV2 | 7855 | 230.68 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7887 | 229.61 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 14047 | 129.00 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1026 | 1765.28 MB/s | 0 | 0 | 13.7× |
| LightningArena | 1026 | 1765.34 MB/s | 0 | 0 | 13.7× |
| LightningDestructive | 1047 | 1730.54 MB/s | 0 | 0 | 13.4× |
| Easyjson | 2513 | 721.11 MB/s | 24 | 1 | 5.6× |
| Goccy | 2800 | 647.03 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 5897 | 307.26 MB/s | 3718 | 40 | 2.4× |
| Sonic | 5924 | 305.85 MB/s | 3777 | 40 | 2.4× |
| JSONV2 | 7725 | 234.56 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7881 | 229.80 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 14069 | 128.79 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1212 | 1494.89 MB/s | 144 | 10 | 11.6× |
| LightningArena | 1216 | 1489.57 MB/s | 144 | 10 | 11.6× |
| LightningDestructive | 1259 | 1439.33 MB/s | 144 | 10 | 11.2× |
| Easyjson | 2766 | 655.04 MB/s | 144 | 10 | 5.1× |
| Goccy | 2887 | 627.75 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6121 | 296.03 MB/s | 3804 | 42 | 2.3× |
| Sonic | 6146 | 294.81 MB/s | 3789 | 42 | 2.3× |
| LightningDecodeAny | 7886 | 229.63 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7924 | 228.68 MB/s | 632 | 7 | 1.8× |
| Stdlib | 14049 | 128.97 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 698 | 707.98 MB/s | 160 | 1 | 8.0× |
| Lightning | 699 | 707.03 MB/s | 160 | 1 | 8.0× |
| SonicFastest | 1235 | 399.94 MB/s | 976 | 6 | 4.5× |
| Sonic | 1238 | 398.88 MB/s | 978 | 6 | 4.5× |
| LightningArena | 1310 | 377.10 MB/s | 4096 | 1 | 4.3× |
| LightningDecodeAny | 1384 | 356.23 MB/s | 1296 | 26 | 4.0× |
| Easyjson | 2236 | 220.98 MB/s | 448 | 3 | 2.5× |
| Goccy | 2427 | 203.51 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3267 | 151.19 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5576 | 88.59 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 429 | 535.95 MB/s | 160 | 1 | 9.7× |
| LightningDestructive | 431 | 534.14 MB/s | 160 | 1 | 9.7× |
| Sonic | 891 | 258.24 MB/s | 657 | 6 | 4.7× |
| SonicFastest | 891 | 258.21 MB/s | 657 | 6 | 4.7× |
| LightningArena | 1088 | 211.40 MB/s | 4096 | 1 | 3.8× |
| LightningDecodeAny | 1133 | 202.19 MB/s | 1296 | 26 | 3.7× |
| Easyjson | 1406 | 163.59 MB/s | 448 | 3 | 3.0× |
| Goccy | 1582 | 145.36 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2456 | 93.66 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4169 | 55.17 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 54269 | 1200.18 MB/s | 97220 | 98 | 10.1× |
| LightningArena | 55820 | 1166.82 MB/s | 103440 | 103 | 9.8× |
| Lightning | 55923 | 1164.67 MB/s | 103440 | 103 | 9.8× |
| Sonic | 96608 | 674.19 MB/s | 155233 | 75 | 5.7× |
| SonicFastest | 96786 | 672.95 MB/s | 155957 | 75 | 5.6× |
| Goccy | 146719 | 443.92 MB/s | 229274 | 134 | 3.7× |
| LightningDecodeAny | 178599 | 298.60 MB/s | 180048 | 3245 | 3.1× |
| JSONV2 | 222955 | 292.13 MB/s | 206652 | 607 | 2.5× |
| Stdlib | 546685 | 119.14 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2419055 | 802.16 MB/s | 2864593 | 1380 | 9.7× |
| Lightning | 2486855 | 780.29 MB/s | 2864594 | 1380 | 9.4× |
| LightningArena | 2535020 | 765.47 MB/s | 2864594 | 1380 | 9.3× |
| SonicFastest | 4555768 | 425.94 MB/s | 14608566 | 1407 | 5.2× |
| Sonic | 4606610 | 421.24 MB/s | 14608585 | 1407 | 5.1× |
| Goccy | 4751587 | 408.38 MB/s | 4065661 | 13510 | 4.9× |
| Easyjson | 7562061 | 256.61 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9053652 | 214.33 MB/s | 7063040 | 218633 | 2.6× |
| JSONV2 | 11186668 | 173.46 MB/s | 3237220 | 13947 | 2.1× |
| Stdlib | 23488089 | 82.62 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1002648 | 3319.04 MB/s | 351704 | 1286 | 20.8× |
| Lightning | 1617946 | 2056.82 MB/s | 2488906 | 2995 | 12.9× |
| LightningArena | 1632407 | 2038.60 MB/s | 2488904 | 2995 | 12.8× |
| Sonic | 2644263 | 1258.51 MB/s | 6446631 | 4248 | 7.9× |
| SonicFastest | 2711826 | 1227.15 MB/s | 6453998 | 4248 | 7.7× |
| LightningDecodeAny | 3592238 | 855.67 MB/s | 4876912 | 56892 | 5.8× |
| Goccy | 4567950 | 728.52 MB/s | 3948907 | 3816 | 4.6× |
| JSONV2 | 7330587 | 453.97 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20867446 | 159.47 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 208919 | 1054.70 MB/s | 135872 | 226 | 9.8× |
| Lightning | 210063 | 1048.95 MB/s | 135872 | 226 | 9.7× |
| LightningArena | 211253 | 1043.04 MB/s | 135872 | 226 | 9.6× |
| SonicFastest | 377896 | 583.09 MB/s | 300586 | 398 | 5.4× |
| Sonic | 379313 | 580.91 MB/s | 304547 | 398 | 5.4× |
| Goccy | 433092 | 508.77 MB/s | 364658 | 1067 | 4.7× |
| Easyjson | 551663 | 399.42 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 738286 | 298.46 MB/s | 129743 | 470 | 2.8× |
| LightningDecodeAny | 869027 | 124.64 MB/s | 897218 | 11703 | 2.3× |
| Stdlib | 2037678 | 108.14 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11263831 | 719.12 MB/s | 11845073 | 20816 | 7.9× |
| Lightning | 11441061 | 707.98 MB/s | 11845078 | 20816 | 7.8× |
| LightningArena | 11526311 | 702.74 MB/s | 11845072 | 20816 | 7.7× |
| SonicFastest | 16833722 | 481.18 MB/s | 70887556 | 40014 | 5.3× |
| Sonic | 16854815 | 480.58 MB/s | 70901921 | 40014 | 5.3× |
| Goccy | 23423702 | 345.81 MB/s | 17018431 | 107148 | 3.8× |
| Easyjson | 30689689 | 263.93 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 37076907 | 140.33 MB/s | 46279353 | 747112 | 2.4× |
| JSONV2 | 43727264 | 185.24 MB/s | 15233769 | 78972 | 2.0× |
| Stdlib | 89151057 | 90.86 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5242067 | 569.14 MB/s | 3764714 | 1504 | 9.0× |
| LightningDestructive | 5708999 | 522.59 MB/s | 3758856 | 29356 | 8.2× |
| Lightning | 5771215 | 516.96 MB/s | 3758856 | 29356 | 8.1× |
| Sonic | 8697223 | 343.04 MB/s | 26748774 | 56760 | 5.4× |
| SonicFastest | 8747691 | 341.06 MB/s | 26588266 | 56760 | 5.4× |
| Goccy | 16447047 | 181.40 MB/s | 10517851 | 273645 | 2.9× |
| Easyjson | 16718639 | 178.45 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 16849895 | 108.86 MB/s | 23982582 | 351152 | 2.8× |
| JSONV2 | 24125479 | 123.66 MB/s | 9257151 | 86278 | 1.9× |
| Stdlib | 46983305 | 63.50 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1204003 | 600.99 MB/s | 911392 | 30 | 9.6× |
| Lightning | 1254569 | 576.77 MB/s | 907596 | 3618 | 9.2× |
| LightningDestructive | 1254653 | 576.73 MB/s | 907601 | 3618 | 9.2× |
| Sonic | 1772816 | 408.16 MB/s | 3183913 | 7226 | 6.5× |
| SonicFastest | 1775747 | 407.49 MB/s | 3179229 | 7226 | 6.5× |
| LightningDecodeAny | 4165083 | 156.20 MB/s | 6500457 | 76546 | 2.8× |
| Easyjson | 4195298 | 172.48 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 4795079 | 150.90 MB/s | 2750874 | 80271 | 2.4× |
| JSONV2 | 5741910 | 126.02 MB/s | 2704636 | 7318 | 2.0× |
| Stdlib | 11567689 | 62.55 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1841835 | 856.40 MB/s | 911393 | 30 | 8.6× |
| LightningDestructive | 1856004 | 849.87 MB/s | 907600 | 3618 | 8.5× |
| Lightning | 1896099 | 831.89 MB/s | 907594 | 3618 | 8.3× |
| SonicFastest | 2249170 | 701.30 MB/s | 5784570 | 7226 | 7.0× |
| Sonic | 2274311 | 693.55 MB/s | 5783107 | 7226 | 6.9× |
| LightningDecodeAny | 3881518 | 194.10 MB/s | 6500459 | 76546 | 4.1× |
| Goccy | 5554519 | 283.98 MB/s | 3543934 | 80265 | 2.8× |
| Easyjson | 5559349 | 283.73 MB/s | 2847904 | 3698 | 2.8× |
| JSONV2 | 6434225 | 245.15 MB/s | 2704594 | 7318 | 2.4× |
| Stdlib | 15751738 | 100.14 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 212050 | 707.97 MB/s | 81920 | 1 | 8.7× |
| Lightning | 212056 | 707.95 MB/s | 81920 | 1 | 8.7× |
| LightningDestructive | 212435 | 706.68 MB/s | 81920 | 1 | 8.7× |
| Sonic | 266184 | 563.99 MB/s | 241976 | 6 | 6.9× |
| SonicFastest | 268255 | 559.63 MB/s | 248300 | 6 | 6.9× |
| LightningDecodeAny | 467350 | 321.22 MB/s | 745764 | 10016 | 3.9× |
| Goccy | 863607 | 173.83 MB/s | 323656 | 10004 | 2.1× |
| JSONV2 | 1065430 | 140.90 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1840108 | 81.58 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 29388 | 956.75 MB/s | 29216 | 103 | 10.4× |
| LightningDestructive | 29407 | 956.15 MB/s | 29088 | 101 | 10.3× |
| Lightning | 29817 | 942.98 MB/s | 29216 | 103 | 10.2× |
| Sonic | 62666 | 448.68 MB/s | 46631 | 103 | 4.9× |
| SonicFastest | 62693 | 448.49 MB/s | 46767 | 103 | 4.9× |
| Easyjson | 68171 | 412.45 MB/s | 32304 | 138 | 4.5× |
| Goccy | 70814 | 397.05 MB/s | 59200 | 188 | 4.3× |
| JSONV2 | 134486 | 209.07 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 144992 | 193.92 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 304257 | 92.41 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1784 | 1304.98 MB/s | 32 | 1 | 12.7× |
| LightningArena | 1791 | 1299.67 MB/s | 32 | 1 | 12.6× |
| LightningDestructive | 1846 | 1260.76 MB/s | 32 | 1 | 12.3× |
| Goccy | 4131 | 563.55 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4205 | 553.60 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5077 | 458.55 MB/s | 4251 | 6 | 4.5× |
| Sonic | 5099 | 456.56 MB/s | 4291 | 6 | 4.4× |
| JSONV2 | 8565 | 271.79 MB/s | 1000 | 6 | 2.6× |
| LightningDecodeAny | 9823 | 171.54 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22650 | 102.78 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 207 | 913.18 MB/s | 0 | 0 | 12.0× |
| Lightning | 207 | 912.31 MB/s | 0 | 0 | 12.0× |
| LightningDestructive | 209 | 903.07 MB/s | 0 | 0 | 11.9× |
| Goccy | 384 | 492.13 MB/s | 304 | 2 | 6.5× |
| Easyjson | 501 | 376.93 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 797 | 237.04 MB/s | 500 | 4 | 3.1× |
| Sonic | 801 | 235.88 MB/s | 503 | 4 | 3.1× |
| JSONV2 | 1046 | 180.64 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1186 | 112.97 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2483 | 76.12 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1346 | 1627.90 MB/s | 0 | 0 | 12.0× |
| LightningArena | 1356 | 1615.65 MB/s | 0 | 0 | 11.9× |
| LightningDestructive | 1369 | 1600.93 MB/s | 0 | 0 | 11.8× |
| Easyjson | 3192 | 686.48 MB/s | 24 | 1 | 5.0× |
| Goccy | 3195 | 685.66 MB/s | 2864 | 4 | 5.0× |
| SonicFastest | 6325 | 346.41 MB/s | 3962 | 40 | 2.5× |
| Sonic | 6327 | 346.27 MB/s | 3958 | 40 | 2.5× |
| LightningDecodeAny | 7897 | 229.34 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 7999 | 273.91 MB/s | 640 | 6 | 2.0× |
| Stdlib | 16096 | 136.12 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 599231 | 851.88 MB/s | 457537 | 1009 | 10.2× |
| Lightning | 608350 | 839.12 MB/s | 457537 | 1009 | 10.0× |
| LightningArena | 611847 | 834.32 MB/s | 457536 | 1009 | 10.0× |
| Goccy | 1139783 | 447.87 MB/s | 1133548 | 5006 | 5.3× |
| Sonic | 1144662 | 445.96 MB/s | 891817 | 2006 | 5.3× |
| SonicFastest | 1147604 | 444.82 MB/s | 890139 | 2006 | 5.3× |
| Easyjson | 1546243 | 330.14 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3202837 | 159.38 MB/s | 1076013 | 12646 | 1.9× |
| LightningDecodeAny | 3319224 | 139.03 MB/s | 2950648 | 64018 | 1.8× |
| Stdlib | 6088720 | 83.84 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 883 | 22399.94 MB/s | 0 | 0 | 127.5× |
| Lightning | 885 | 22356.69 MB/s | 0 | 0 | 127.2× |
| LightningDestructive | 916 | 21602.16 MB/s | 0 | 0 | 122.9× |
| Goccy | 19825 | 998.21 MB/s | 20491 | 2 | 5.7× |
| Sonic | 27913 | 708.94 MB/s | 21967 | 4 | 4.0× |
| SonicFastest | 27959 | 707.78 MB/s | 22057 | 4 | 4.0× |
| JSONV2 | 29637 | 667.71 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 76373 | 259.10 MB/s | 116864 | 2015 | 1.5× |
| Easyjson | 86823 | 227.92 MB/s | 0 | 0 | 1.3× |
| Stdlib | 112596 | 175.75 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2202 | 8231.71 MB/s | 0 | 0 | 46.9× |
| Lightning | 2348 | 7718.90 MB/s | 432 | 2 | 44.0× |
| LightningArena | 2348 | 7719.63 MB/s | 432 | 2 | 44.0× |
| Easyjson | 3993 | 4539.06 MB/s | 432 | 2 | 25.9× |
| Sonic | 10158 | 1784.28 MB/s | 23494 | 6 | 10.2× |
| SonicFastest | 10258 | 1766.76 MB/s | 23596 | 6 | 10.1× |
| Goccy | 15898 | 1140.05 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 16090 | 1111.34 MB/s | 29088 | 191 | 6.4× |
| JSONV2 | 45044 | 402.36 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103267 | 175.51 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2315451 | 867.43 MB/s | 3089564 | 6821 | 8.1× |
| LightningArena | 2387477 | 841.26 MB/s | 3094372 | 6703 | 7.8× |
| Lightning | 2407963 | 834.10 MB/s | 3091277 | 6827 | 7.8× |
| Goccy | 4210843 | 476.98 MB/s | 5411709 | 15830 | 4.4× |
| SonicFastest | 4436020 | 452.77 MB/s | 10913154 | 13683 | 4.2× |
| Sonic | 4479820 | 448.34 MB/s | 10903887 | 13683 | 4.2× |
| Easyjson | 4901566 | 409.77 MB/s | 2981483 | 7439 | 3.8× |
| JSONV2 | 6955673 | 288.76 MB/s | 3173686 | 14563 | 2.7× |
| LightningDecodeAny | 7431266 | 153.71 MB/s | 8503515 | 134008 | 2.5× |
| Stdlib | 18679375 | 107.52 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 862 | 637.05 MB/s | 480 | 1 | 6.6× |
| Lightning | 866 | 634.13 MB/s | 480 | 1 | 6.5× |
| LightningArena | 870 | 631.08 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1617 | 338.94 MB/s | 2021 | 46 | 3.5× |
| Easyjson | 2163 | 253.81 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2682 | 204.71 MB/s | 1944 | 26 | 2.1× |
| SonicFastest | 2684 | 204.56 MB/s | 1953 | 26 | 2.1× |
| Goccy | 3058 | 179.54 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3293 | 166.72 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5656 | 97.07 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 464033 | 1360.93 MB/s | 402728 | 545 | 11.7× |
| Lightning | 531723 | 1187.67 MB/s | 451257 | 857 | 10.2× |
| LightningArena | 531822 | 1187.45 MB/s | 453017 | 712 | 10.2× |
| Sonic | 1018213 | 620.22 MB/s | 1025211 | 1102 | 5.3× |
| SonicFastest | 1021048 | 618.50 MB/s | 1019126 | 1102 | 5.3× |
| Easyjson | 1143469 | 552.28 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1151896 | 548.24 MB/s | 988234 | 1201 | 4.7× |
| JSONV2 | 2156183 | 292.89 MB/s | 571616 | 3144 | 2.5× |
| LightningDecodeAny | 2322159 | 201.07 MB/s | 2076505 | 30126 | 2.3× |
| Stdlib | 5413177 | 116.66 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 575209 | 977.75 MB/s | 546571 | 429 | 9.1× |
| LightningArena | 735345 | 764.82 MB/s | 771666 | 1088 | 7.1× |
| Lightning | 737006 | 763.10 MB/s | 769938 | 1235 | 7.1× |
| Sonic | 1010600 | 556.51 MB/s | 939439 | 1476 | 5.2× |
| SonicFastest | 1012032 | 555.72 MB/s | 945105 | 1476 | 5.2× |
| Goccy | 1306624 | 430.43 MB/s | 1040603 | 1030 | 4.0× |
| Easyjson | 1729114 | 325.26 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2559453 | 219.74 MB/s | 2180440 | 30126 | 2.1× |
| JSONV2 | 2761232 | 203.68 MB/s | 927440 | 3482 | 1.9× |
| Stdlib | 5256660 | 106.99 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 612325 | 870.74 MB/s | 333416 | 2084 | 8.9× |
| Lightning | 633377 | 841.80 MB/s | 368224 | 2293 | 8.6× |
| LightningArena | 636993 | 837.02 MB/s | 368224 | 2293 | 8.6× |
| Easyjson | 1120135 | 475.99 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1141018 | 467.28 MB/s | 1033832 | 4351 | 4.8× |
| Sonic | 1149281 | 463.92 MB/s | 1045793 | 4351 | 4.8× |
| Goccy | 1297864 | 410.81 MB/s | 1167220 | 5409 | 4.2× |
| JSONV2 | 2544907 | 209.51 MB/s | 745453 | 13288 | 2.2× |
| LightningDecodeAny | 3330471 | 160.09 MB/s | 2992877 | 50076 | 1.6× |
| Stdlib | 5477796 | 97.33 MB/s | 798692 | 17133 | 1.0× |
