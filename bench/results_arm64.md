# JSON Deserialization Benchmarks

- generated 2026-08-11T11:20:07Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105321 | 1208.45 MB/s | 49760 | 3 | 10.5× |
| LightningArena | 105630 | 1204.92 MB/s | 49760 | 3 | 10.5× |
| LightningDestructive | 105724 | 1203.84 MB/s | 49280 | 2 | 10.5× |
| Sonic | 180202 | 706.29 MB/s | 192605 | 10 | 6.1× |
| SonicFastest | 181092 | 702.82 MB/s | 194183 | 10 | 6.1× |
| Goccy | 188462 | 675.34 MB/s | 224540 | 884 | 5.9× |
| Easyjson | 211268 | 602.43 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 416645 | 305.48 MB/s | 195119 | 1805 | 2.7× |
| LightningDecodeAny | 423135 | 223.69 MB/s | 463409 | 9708 | 2.6× |
| Stdlib | 1105865 | 115.09 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3647675 | 617.12 MB/s | 2532849 | 1143 | 7.2× |
| LightningArena | 3685135 | 610.85 MB/s | 2532848 | 1143 | 7.1× |
| Lightning | 3692230 | 609.67 MB/s | 2532851 | 1143 | 7.1× |
| Sonic | 4477342 | 502.77 MB/s | 15232100 | 970 | 5.8× |
| SonicFastest | 4508894 | 499.25 MB/s | 15233730 | 970 | 5.8× |
| Goccy | 10095436 | 222.98 MB/s | 4100088 | 56531 | 2.6× |
| Easyjson | 10941796 | 205.73 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12356753 | 182.17 MB/s | 19380210 | 223896 | 2.1× |
| JSONV2 | 15978307 | 140.88 MB/s | 3123245 | 3083 | 1.6× |
| Stdlib | 26114837 | 86.20 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 467104 | 578.89 MB/s | 397296 | 567 | 7.2× |
| LightningArena | 467728 | 578.12 MB/s | 397296 | 567 | 7.2× |
| LightningDestructive | 467990 | 577.80 MB/s | 397296 | 567 | 7.2× |
| SonicFastest | 623782 | 433.49 MB/s | 472585 | 968 | 5.4× |
| Sonic | 629094 | 429.83 MB/s | 479344 | 968 | 5.3× |
| Goccy | 1380549 | 195.87 MB/s | 542446 | 8122 | 2.4× |
| Easyjson | 1386985 | 194.96 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1629954 | 165.90 MB/s | 2543877 | 29687 | 2.1× |
| JSONV2 | 2070052 | 130.63 MB/s | 348153 | 1628 | 1.6× |
| Stdlib | 3350497 | 80.71 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1148391 | 1504.02 MB/s | 765560 | 2798 | 11.6× |
| LightningArena | 1153291 | 1497.63 MB/s | 768416 | 2440 | 11.5× |
| Lightning | 1163022 | 1485.10 MB/s | 765602 | 2799 | 11.4× |
| Sonic | 2015825 | 856.82 MB/s | 2746807 | 4020 | 6.6× |
| SonicFastest | 2020619 | 854.79 MB/s | 2747569 | 4020 | 6.6× |
| Goccy | 2305169 | 749.27 MB/s | 2581776 | 14604 | 5.8× |
| LightningDecodeAny | 4197968 | 119.18 MB/s | 4953694 | 76576 | 3.2× |
| Easyjson | 4224368 | 408.87 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4258974 | 405.54 MB/s | 1011633 | 7594 | 3.1× |
| Stdlib | 13306073 | 129.81 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1159 | 1563.03 MB/s | 0 | 0 | 12.1× |
| Lightning | 1160 | 1561.53 MB/s | 0 | 0 | 12.1× |
| LightningDestructive | 1173 | 1545.11 MB/s | 0 | 0 | 12.0× |
| Easyjson | 2550 | 710.53 MB/s | 24 | 1 | 5.5× |
| Goccy | 2742 | 660.85 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5883 | 308.01 MB/s | 3703 | 40 | 2.4× |
| Sonic | 5923 | 305.90 MB/s | 3744 | 40 | 2.4× |
| LightningDecodeAny | 7824 | 231.46 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7906 | 229.18 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14059 | 128.89 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1201 | 1508.54 MB/s | 0 | 0 | 11.7× |
| LightningDestructive | 1201 | 1508.83 MB/s | 0 | 0 | 11.7× |
| LightningArena | 1208 | 1500.40 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2547 | 711.29 MB/s | 24 | 1 | 5.5× |
| Goccy | 2764 | 655.51 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5931 | 305.52 MB/s | 3850 | 40 | 2.4× |
| Sonic | 5945 | 304.77 MB/s | 3869 | 40 | 2.4× |
| JSONV2 | 7800 | 232.30 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8139 | 222.51 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 14078 | 128.71 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1353 | 1339.42 MB/s | 144 | 10 | 10.4× |
| Lightning | 1355 | 1337.20 MB/s | 144 | 10 | 10.4× |
| LightningDestructive | 1396 | 1298.39 MB/s | 144 | 10 | 10.1× |
| Easyjson | 2757 | 657.14 MB/s | 144 | 10 | 5.1× |
| Goccy | 2806 | 645.74 MB/s | 2600 | 5 | 5.0× |
| Sonic | 6053 | 299.36 MB/s | 3739 | 42 | 2.3× |
| SonicFastest | 6079 | 298.10 MB/s | 3779 | 42 | 2.3× |
| LightningDecodeAny | 7813 | 231.78 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7862 | 230.46 MB/s | 632 | 7 | 1.8× |
| Stdlib | 14049 | 128.97 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 693 | 712.81 MB/s | 160 | 1 | 8.0× |
| LightningDestructive | 695 | 710.74 MB/s | 160 | 1 | 8.0× |
| SonicFastest | 1223 | 404.09 MB/s | 979 | 6 | 4.5× |
| Sonic | 1225 | 403.39 MB/s | 989 | 6 | 4.5× |
| LightningArena | 1266 | 390.31 MB/s | 4096 | 1 | 4.4× |
| LightningDecodeAny | 1360 | 362.54 MB/s | 1296 | 26 | 4.1× |
| Easyjson | 2229 | 221.61 MB/s | 448 | 3 | 2.5× |
| Goccy | 2433 | 203.08 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3293 | 150.01 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5557 | 88.89 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 427 | 538.64 MB/s | 160 | 1 | 9.8× |
| LightningDestructive | 427 | 538.39 MB/s | 160 | 1 | 9.7× |
| Sonic | 871 | 264.13 MB/s | 651 | 6 | 4.8× |
| SonicFastest | 873 | 263.40 MB/s | 654 | 6 | 4.8× |
| LightningArena | 1049 | 219.26 MB/s | 4096 | 1 | 4.0× |
| LightningDecodeAny | 1117 | 205.00 MB/s | 1296 | 26 | 3.7× |
| Easyjson | 1396 | 164.78 MB/s | 448 | 3 | 3.0× |
| Goccy | 1573 | 146.25 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2446 | 94.05 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4164 | 55.23 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 57296 | 1136.77 MB/s | 97220 | 98 | 9.5× |
| Lightning | 58120 | 1120.65 MB/s | 103440 | 103 | 9.3× |
| LightningArena | 58157 | 1119.94 MB/s | 103440 | 103 | 9.3× |
| SonicFastest | 96797 | 672.87 MB/s | 154195 | 75 | 5.6× |
| Sonic | 97154 | 670.40 MB/s | 155550 | 75 | 5.6× |
| Goccy | 140629 | 463.15 MB/s | 229220 | 134 | 3.9× |
| LightningDecodeAny | 173935 | 306.60 MB/s | 180048 | 3245 | 3.1× |
| JSONV2 | 219249 | 297.07 MB/s | 206651 | 607 | 2.5× |
| Stdlib | 543358 | 119.87 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2525467 | 768.36 MB/s | 2864592 | 1380 | 9.3× |
| LightningArena | 2585739 | 750.45 MB/s | 2864593 | 1380 | 9.1× |
| Lightning | 2605713 | 744.70 MB/s | 2864595 | 1380 | 9.0× |
| Sonic | 4526415 | 428.70 MB/s | 14606972 | 1407 | 5.2× |
| SonicFastest | 4620657 | 419.96 MB/s | 14608554 | 1407 | 5.1× |
| Goccy | 4760181 | 407.65 MB/s | 4064131 | 13510 | 4.9× |
| Easyjson | 7499912 | 258.73 MB/s | 3871264 | 15043 | 3.1× |
| LightningDecodeAny | 9524061 | 203.74 MB/s | 7063041 | 218633 | 2.5× |
| JSONV2 | 11100887 | 174.80 MB/s | 3237223 | 13947 | 2.1× |
| Stdlib | 23448431 | 82.75 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1095822 | 3036.83 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1659707 | 2005.07 MB/s | 2488904 | 2995 | 12.4× |
| LightningArena | 1666731 | 1996.62 MB/s | 2488905 | 2995 | 12.4× |
| Sonic | 2583239 | 1288.24 MB/s | 6453988 | 4248 | 8.0× |
| SonicFastest | 2595132 | 1282.34 MB/s | 6432094 | 4248 | 8.0× |
| LightningDecodeAny | 3524359 | 872.15 MB/s | 4876911 | 56892 | 5.9× |
| Goccy | 4439754 | 749.55 MB/s | 3948909 | 3816 | 4.7× |
| JSONV2 | 7397996 | 449.83 MB/s | 5364516 | 13243 | 2.8× |
| Stdlib | 20655700 | 161.11 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 216506 | 1017.74 MB/s | 135872 | 226 | 9.4× |
| LightningArena | 216819 | 1016.27 MB/s | 135872 | 226 | 9.4× |
| LightningDestructive | 218118 | 1010.22 MB/s | 135872 | 226 | 9.3× |
| SonicFastest | 374198 | 588.85 MB/s | 299351 | 398 | 5.4× |
| Sonic | 377517 | 583.67 MB/s | 306464 | 398 | 5.4× |
| Goccy | 429389 | 513.16 MB/s | 364300 | 1067 | 4.7× |
| Easyjson | 544651 | 404.56 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 717120 | 307.27 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 827254 | 130.93 MB/s | 897218 | 11703 | 2.5× |
| Stdlib | 2037416 | 108.15 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11468552 | 706.28 MB/s | 11845078 | 20816 | 7.8× |
| Lightning | 11786970 | 687.20 MB/s | 11845073 | 20816 | 7.6× |
| LightningArena | 11925110 | 679.24 MB/s | 11845073 | 20816 | 7.5× |
| SonicFastest | 16671702 | 485.86 MB/s | 70929614 | 40014 | 5.4× |
| Sonic | 16684103 | 485.49 MB/s | 70887293 | 40014 | 5.3× |
| Goccy | 23344317 | 346.98 MB/s | 16989235 | 107148 | 3.8× |
| Easyjson | 30583139 | 264.85 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 36787821 | 141.43 MB/s | 46279353 | 747112 | 2.4× |
| JSONV2 | 43562178 | 185.94 MB/s | 15233769 | 78972 | 2.0× |
| Stdlib | 89222278 | 90.78 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5322464 | 560.54 MB/s | 3764714 | 1504 | 8.8× |
| LightningDestructive | 5672248 | 525.98 MB/s | 3758856 | 29356 | 8.3× |
| Lightning | 5796799 | 514.67 MB/s | 3758859 | 29356 | 8.1× |
| Sonic | 8481051 | 351.78 MB/s | 26638483 | 56760 | 5.5× |
| SonicFastest | 8481464 | 351.76 MB/s | 26632107 | 56760 | 5.5× |
| Easyjson | 16277151 | 183.29 MB/s | 9479440 | 30115 | 2.9× |
| LightningDecodeAny | 16429850 | 111.64 MB/s | 23982582 | 351152 | 2.9× |
| Goccy | 16568435 | 180.07 MB/s | 10553148 | 273646 | 2.8× |
| JSONV2 | 24124483 | 123.67 MB/s | 9257175 | 86278 | 2.0× |
| Stdlib | 47057446 | 63.40 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1210521 | 597.76 MB/s | 911393 | 30 | 9.6× |
| LightningDestructive | 1247829 | 579.88 MB/s | 907601 | 3618 | 9.3× |
| Lightning | 1257708 | 575.33 MB/s | 907596 | 3618 | 9.2× |
| Sonic | 1750936 | 413.26 MB/s | 3183810 | 7226 | 6.6× |
| SonicFastest | 1764221 | 410.15 MB/s | 3193298 | 7226 | 6.6× |
| Easyjson | 4193618 | 172.55 MB/s | 2847906 | 3698 | 2.8× |
| LightningDecodeAny | 4233020 | 153.69 MB/s | 6500461 | 76546 | 2.7× |
| Goccy | 4820806 | 150.10 MB/s | 2842322 | 80275 | 2.4× |
| JSONV2 | 5575389 | 129.78 MB/s | 2704623 | 7318 | 2.1× |
| Stdlib | 11573254 | 62.52 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1834598 | 859.78 MB/s | 911393 | 30 | 8.6× |
| LightningDestructive | 1838415 | 858.00 MB/s | 907600 | 3618 | 8.6× |
| Lightning | 1882229 | 838.02 MB/s | 907594 | 3618 | 8.4× |
| SonicFastest | 2235640 | 705.55 MB/s | 5791524 | 7226 | 7.0× |
| Sonic | 2249052 | 701.34 MB/s | 5788176 | 7226 | 7.0× |
| LightningDecodeAny | 3876620 | 194.34 MB/s | 6500457 | 76546 | 4.1× |
| Easyjson | 5579498 | 282.71 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5751338 | 274.26 MB/s | 3558121 | 80266 | 2.7× |
| JSONV2 | 6373333 | 247.49 MB/s | 2704591 | 7318 | 2.5× |
| Stdlib | 15730881 | 100.27 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 211943 | 708.32 MB/s | 81920 | 1 | 8.6× |
| Lightning | 211956 | 708.28 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 212669 | 705.91 MB/s | 81920 | 1 | 8.6× |
| Sonic | 277858 | 540.29 MB/s | 259336 | 6 | 6.6× |
| SonicFastest | 278446 | 539.15 MB/s | 262203 | 6 | 6.5× |
| LightningDecodeAny | 476150 | 315.28 MB/s | 745765 | 10016 | 3.8× |
| Goccy | 882212 | 170.17 MB/s | 324806 | 10004 | 2.1× |
| JSONV2 | 1089858 | 137.75 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1823403 | 82.33 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 31678 | 887.59 MB/s | 29216 | 103 | 9.6× |
| LightningArena | 31703 | 886.90 MB/s | 29216 | 103 | 9.6× |
| LightningDestructive | 31763 | 885.22 MB/s | 29088 | 101 | 9.5× |
| Sonic | 63542 | 442.49 MB/s | 47040 | 103 | 4.8× |
| SonicFastest | 63636 | 441.84 MB/s | 46960 | 103 | 4.8× |
| Easyjson | 67996 | 413.51 MB/s | 32304 | 138 | 4.5× |
| Goccy | 70634 | 398.06 MB/s | 59210 | 188 | 4.3× |
| JSONV2 | 132775 | 211.76 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 143695 | 195.67 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 302792 | 92.86 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1956 | 1190.12 MB/s | 32 | 1 | 11.6× |
| LightningArena | 1966 | 1183.97 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2030 | 1146.80 MB/s | 32 | 1 | 11.2× |
| Goccy | 4083 | 570.11 MB/s | 3649 | 4 | 5.6× |
| Easyjson | 4217 | 552.10 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5113 | 455.27 MB/s | 4348 | 6 | 4.5× |
| Sonic | 5121 | 454.59 MB/s | 4349 | 6 | 4.4× |
| JSONV2 | 8379 | 277.84 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9631 | 174.96 MB/s | 10200 | 195 | 2.4× |
| Stdlib | 22760 | 102.28 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221 | 854.21 MB/s | 0 | 0 | 11.0× |
| LightningArena | 222 | 852.17 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 224 | 844.95 MB/s | 0 | 0 | 10.9× |
| Goccy | 376 | 502.63 MB/s | 304 | 2 | 6.5× |
| Easyjson | 485 | 389.51 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 783 | 241.30 MB/s | 501 | 4 | 3.1× |
| Sonic | 787 | 240.23 MB/s | 500 | 4 | 3.1× |
| JSONV2 | 1024 | 184.54 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1161 | 115.46 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2439 | 77.50 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1513 | 1448.23 MB/s | 0 | 0 | 10.6× |
| LightningArena | 1520 | 1441.61 MB/s | 0 | 0 | 10.5× |
| LightningDestructive | 1526 | 1436.11 MB/s | 0 | 0 | 10.5× |
| Easyjson | 3190 | 686.93 MB/s | 24 | 1 | 5.0× |
| Goccy | 3224 | 679.52 MB/s | 2864 | 4 | 5.0× |
| SonicFastest | 6318 | 346.80 MB/s | 3926 | 40 | 2.5× |
| Sonic | 6363 | 344.35 MB/s | 3944 | 40 | 2.5× |
| LightningDecodeAny | 7856 | 230.52 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 8015 | 273.37 MB/s | 640 | 6 | 2.0× |
| Stdlib | 16004 | 136.90 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 648587 | 787.06 MB/s | 457536 | 1009 | 9.3× |
| Lightning | 662601 | 770.41 MB/s | 457537 | 1009 | 9.1× |
| LightningArena | 664831 | 767.83 MB/s | 457536 | 1009 | 9.0× |
| Goccy | 1141716 | 447.11 MB/s | 1138769 | 5006 | 5.3× |
| SonicFastest | 1165494 | 437.99 MB/s | 906298 | 2006 | 5.2× |
| Sonic | 1169270 | 436.58 MB/s | 917630 | 2006 | 5.1× |
| Easyjson | 1516615 | 336.59 MB/s | 863777 | 3012 | 4.0× |
| JSONV2 | 3185839 | 160.23 MB/s | 1076008 | 12646 | 1.9× |
| LightningDecodeAny | 3227065 | 143.00 MB/s | 2950650 | 64018 | 1.9× |
| Stdlib | 6008325 | 84.96 MB/s | 1162119 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14648.63 MB/s | 0 | 0 | 103.8× |
| LightningArena | 1351 | 14642.90 MB/s | 0 | 0 | 103.8× |
| LightningDestructive | 1390 | 14232.35 MB/s | 0 | 0 | 100.9× |
| Goccy | 19873 | 995.79 MB/s | 20491 | 2 | 7.1× |
| Sonic | 27164 | 728.50 MB/s | 22586 | 4 | 5.2× |
| SonicFastest | 27254 | 726.09 MB/s | 22879 | 4 | 5.1× |
| JSONV2 | 29752 | 665.12 MB/s | 8 | 1 | 4.7× |
| LightningDecodeAny | 73869 | 267.88 MB/s | 116864 | 2015 | 1.9× |
| Easyjson | 90703 | 218.17 MB/s | 0 | 0 | 1.5× |
| Stdlib | 140266 | 141.08 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2610 | 6942.74 MB/s | 0 | 0 | 39.3× |
| LightningArena | 2757 | 6574.38 MB/s | 432 | 2 | 37.2× |
| Lightning | 2760 | 6567.07 MB/s | 432 | 2 | 37.1× |
| Easyjson | 3944 | 4595.67 MB/s | 432 | 2 | 26.0× |
| Sonic | 9981 | 1815.81 MB/s | 22880 | 6 | 10.3× |
| SonicFastest | 10089 | 1796.50 MB/s | 23204 | 6 | 10.2× |
| Goccy | 15571 | 1163.97 MB/s | 19459 | 2 | 6.6× |
| LightningDecodeAny | 16446 | 1087.32 MB/s | 29088 | 191 | 6.2× |
| JSONV2 | 45603 | 397.43 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102515 | 176.79 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2419919 | 829.98 MB/s | 3089564 | 6821 | 7.7× |
| LightningArena | 2481493 | 809.39 MB/s | 3094370 | 6703 | 7.5× |
| Lightning | 2496238 | 804.61 MB/s | 3091278 | 6827 | 7.5× |
| Goccy | 4286184 | 468.60 MB/s | 5412037 | 15831 | 4.4× |
| SonicFastest | 4518106 | 444.54 MB/s | 10880957 | 13683 | 4.1× |
| Sonic | 4561542 | 440.31 MB/s | 10882200 | 13683 | 4.1× |
| Easyjson | 4932070 | 407.23 MB/s | 2981482 | 7439 | 3.8× |
| LightningDecodeAny | 6904565 | 165.44 MB/s | 8503514 | 134008 | 2.7× |
| JSONV2 | 7036383 | 285.44 MB/s | 3173686 | 14563 | 2.7× |
| Stdlib | 18683512 | 107.50 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 915 | 600.03 MB/s | 480 | 1 | 6.2× |
| LightningDestructive | 918 | 598.22 MB/s | 480 | 1 | 6.2× |
| LightningArena | 920 | 596.78 MB/s | 480 | 1 | 6.2× |
| LightningDecodeAny | 1716 | 319.37 MB/s | 2021 | 46 | 3.3× |
| Easyjson | 2226 | 246.68 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2713 | 202.36 MB/s | 2032 | 26 | 2.1× |
| Sonic | 2715 | 202.18 MB/s | 2051 | 26 | 2.1× |
| Goccy | 3081 | 178.18 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3362 | 163.27 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5696 | 96.39 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 498658 | 1266.43 MB/s | 402729 | 545 | 10.8× |
| LightningArena | 559195 | 1129.33 MB/s | 453017 | 712 | 9.6× |
| Lightning | 559478 | 1128.76 MB/s | 451257 | 857 | 9.6× |
| SonicFastest | 1017221 | 620.82 MB/s | 993105 | 1102 | 5.3× |
| Sonic | 1018446 | 620.08 MB/s | 1007793 | 1102 | 5.3× |
| Easyjson | 1133226 | 557.27 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1178899 | 535.68 MB/s | 987726 | 1201 | 4.6× |
| JSONV2 | 2151239 | 293.56 MB/s | 571612 | 3144 | 2.5× |
| LightningDecodeAny | 2344957 | 199.11 MB/s | 2076505 | 30126 | 2.3× |
| Stdlib | 5378140 | 117.42 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 612801 | 917.77 MB/s | 546572 | 429 | 8.6× |
| Lightning | 765264 | 734.92 MB/s | 769938 | 1235 | 6.9× |
| LightningArena | 774219 | 726.42 MB/s | 771666 | 1088 | 6.8× |
| SonicFastest | 1033462 | 544.20 MB/s | 935243 | 1476 | 5.1× |
| Sonic | 1042653 | 539.40 MB/s | 945105 | 1476 | 5.1× |
| Goccy | 1348429 | 417.08 MB/s | 1036062 | 1029 | 3.9× |
| Easyjson | 1746773 | 321.97 MB/s | 775153 | 1254 | 3.0× |
| LightningDecodeAny | 2565981 | 219.18 MB/s | 2180440 | 30126 | 2.1× |
| JSONV2 | 2759056 | 203.84 MB/s | 927443 | 3482 | 1.9× |
| Stdlib | 5268965 | 106.74 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 652378 | 817.28 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 677608 | 786.85 MB/s | 368224 | 2293 | 8.1× |
| LightningArena | 678010 | 786.39 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1095393 | 486.75 MB/s | 428361 | 3273 | 5.0× |
| SonicFastest | 1147313 | 464.72 MB/s | 1036980 | 4351 | 4.8× |
| Sonic | 1155421 | 461.46 MB/s | 1040548 | 4351 | 4.7× |
| Goccy | 1286675 | 414.38 MB/s | 1167228 | 5409 | 4.2× |
| JSONV2 | 2534135 | 210.40 MB/s | 745446 | 13288 | 2.2× |
| LightningDecodeAny | 3285866 | 162.26 MB/s | 2992877 | 50076 | 1.7× |
| Stdlib | 5462477 | 97.61 MB/s | 798691 | 17133 | 1.0× |
