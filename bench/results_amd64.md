# JSON Deserialization Benchmarks

- generated 2026-09-07T19:11:49Z
- go version go1.26.7 linux/amd64
- cpu: Intel(R) Xeon(R) 6973P-C (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 67742 | 1878.82 MB/s | 49280 | 2 | 12.3× |
| Lightning | 67830 | 1876.38 MB/s | 49760 | 3 | 12.3× |
| LightningArena | 67884 | 1874.90 MB/s | 49760 | 3 | 12.2× |
| SonicFastest | 137445 | 926.01 MB/s | 213666 | 15 | 6.0× |
| Sonic | 138724 | 917.47 MB/s | 213844 | 15 | 6.0× |
| Easyjson | 168023 | 757.49 MB/s | 122864 | 14 | 4.9× |
| Goccy | 176709 | 720.25 MB/s | 225324 | 884 | 4.7× |
| JSONV2 | 309157 | 411.68 MB/s | 195132 | 1805 | 2.7× |
| LightningDecodeAny | 319107 | 296.62 MB/s | 463411 | 9708 | 2.6× |
| Stdlib | 831212 | 153.12 MB/s | 199697 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 2080883 | 1081.78 MB/s | 2532850 | 1143 | 9.7× |
| LightningDestructive | 2090587 | 1076.76 MB/s | 2532850 | 1143 | 9.6× |
| Lightning | 2214478 | 1016.52 MB/s | 2532851 | 1143 | 9.1× |
| Sonic | 3145901 | 715.55 MB/s | 4898495 | 2584 | 6.4× |
| SonicFastest | 3233808 | 696.10 MB/s | 4898685 | 2584 | 6.2× |
| LightningDecodeAny | 8263309 | 272.41 MB/s | 19380211 | 223896 | 2.4× |
| Goccy | 9429726 | 238.72 MB/s | 4262329 | 56539 | 2.1× |
| Easyjson | 10147206 | 221.84 MB/s | 3099810 | 2120 | 2.0× |
| JSONV2 | 11836054 | 190.19 MB/s | 3123188 | 3083 | 1.7× |
| Stdlib | 20119710 | 111.88 MB/s | 3123394 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 315421 | 857.28 MB/s | 397297 | 567 | 8.8× |
| LightningArena | 319738 | 845.70 MB/s | 397297 | 567 | 8.7× |
| LightningDestructive | 340300 | 794.60 MB/s | 397297 | 567 | 8.2× |
| Sonic | 537127 | 503.42 MB/s | 641488 | 1147 | 5.2× |
| SonicFastest | 539246 | 501.45 MB/s | 641383 | 1147 | 5.2× |
| Goccy | 1300343 | 207.95 MB/s | 541602 | 8122 | 2.1× |
| Easyjson | 1308446 | 206.66 MB/s | 330273 | 749 | 2.1× |
| LightningDecodeAny | 1402031 | 192.87 MB/s | 2543879 | 29687 | 2.0× |
| JSONV2 | 1640271 | 164.85 MB/s | 348161 | 1628 | 1.7× |
| Stdlib | 2778619 | 97.32 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 695190 | 2484.51 MB/s | 765601 | 2799 | 13.8× |
| LightningDestructive | 695498 | 2483.41 MB/s | 765560 | 2798 | 13.8× |
| LightningArena | 698198 | 2473.80 MB/s | 772704 | 2445 | 13.8× |
| Sonic | 1461381 | 1181.90 MB/s | 2711714 | 5548 | 6.6× |
| SonicFastest | 1461927 | 1181.46 MB/s | 2711680 | 5548 | 6.6× |
| Goccy | 1992098 | 867.03 MB/s | 2581658 | 14604 | 4.8× |
| LightningDecodeAny | 2893812 | 172.89 MB/s | 4953693 | 76576 | 3.3× |
| Easyjson | 3140937 | 549.90 MB/s | 972033 | 5389 | 3.1× |
| JSONV2 | 3231136 | 534.55 MB/s | 1011615 | 7594 | 3.0× |
| Stdlib | 9614049 | 179.65 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 636 | 2848.48 MB/s | 0 | 0 | 17.0× |
| LightningArena | 651 | 2784.56 MB/s | 0 | 0 | 16.6× |
| LightningDestructive | 652 | 2779.67 MB/s | 0 | 0 | 16.6× |
| Easyjson | 1880 | 963.86 MB/s | 24 | 1 | 5.7× |
| Goccy | 2613 | 693.54 MB/s | 2608 | 4 | 4.1× |
| SonicFastest | 4012 | 451.60 MB/s | 3351 | 38 | 2.7× |
| Sonic | 4166 | 434.96 MB/s | 3352 | 38 | 2.6× |
| JSONV2 | 5671 | 319.52 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 6187 | 292.70 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 10795 | 167.86 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 650 | 2787.80 MB/s | 0 | 0 | 16.1× |
| LightningArena | 650 | 2785.39 MB/s | 0 | 0 | 16.1× |
| LightningDestructive | 675 | 2686.15 MB/s | 0 | 0 | 15.5× |
| Easyjson | 1855 | 976.81 MB/s | 24 | 1 | 5.7× |
| Goccy | 2510 | 721.97 MB/s | 2608 | 4 | 4.2× |
| SonicFastest | 3872 | 467.96 MB/s | 3350 | 38 | 2.7× |
| Sonic | 4061 | 446.16 MB/s | 3351 | 38 | 2.6× |
| JSONV2 | 5701 | 317.83 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 5943 | 304.75 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 10490 | 172.74 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 790 | 2295.22 MB/s | 144 | 10 | 13.3× |
| Lightning | 796 | 2277.49 MB/s | 144 | 10 | 13.2× |
| LightningDestructive | 838 | 2161.18 MB/s | 144 | 10 | 12.5× |
| Easyjson | 1956 | 926.26 MB/s | 144 | 10 | 5.4× |
| Goccy | 2377 | 762.20 MB/s | 2600 | 5 | 4.4× |
| SonicFastest | 4110 | 440.90 MB/s | 3370 | 40 | 2.6× |
| Sonic | 4300 | 421.39 MB/s | 3371 | 40 | 2.4× |
| JSONV2 | 5356 | 338.29 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 6043 | 299.67 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 10509 | 172.43 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 401 | 1232.31 MB/s | 160 | 1 | 11.1× |
| LightningDestructive | 426 | 1160.15 MB/s | 160 | 1 | 10.4× |
| LightningDecodeAny | 930 | 530.28 MB/s | 1296 | 26 | 4.8× |
| SonicFastest | 999 | 494.54 MB/s | 1078 | 8 | 4.5× |
| Sonic | 1002 | 493.01 MB/s | 1079 | 8 | 4.4× |
| LightningArena | 1285 | 384.55 MB/s | 4120 | 2 | 3.5× |
| Easyjson | 1734 | 284.93 MB/s | 448 | 3 | 2.6× |
| Goccy | 1931 | 255.76 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 2335 | 211.55 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4446 | 111.11 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 270 | 851.82 MB/s | 160 | 1 | 11.7× |
| LightningDestructive | 274 | 838.33 MB/s | 160 | 1 | 11.5× |
| SonicFastest | 627 | 366.60 MB/s | 803 | 8 | 5.0× |
| Sonic | 644 | 357.35 MB/s | 803 | 8 | 4.9× |
| LightningDecodeAny | 847 | 270.32 MB/s | 1296 | 26 | 3.7× |
| Easyjson | 1114 | 206.49 MB/s | 448 | 3 | 2.8× |
| LightningArena | 1144 | 200.98 MB/s | 4120 | 2 | 2.8× |
| Goccy | 1274 | 180.60 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 1704 | 134.99 MB/s | 528 | 7 | 1.9× |
| Stdlib | 3154 | 72.93 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 45736 | 1424.08 MB/s | 97220 | 98 | 9.5× |
| Lightning | 46727 | 1393.89 MB/s | 103441 | 103 | 9.3× |
| LightningArena | 47004 | 1385.66 MB/s | 103441 | 103 | 9.2× |
| SonicFastest | 96332 | 676.12 MB/s | 235910 | 65 | 4.5× |
| Sonic | 98886 | 658.65 MB/s | 235940 | 65 | 4.4× |
| LightningDecodeAny | 128499 | 415.02 MB/s | 180049 | 3245 | 3.4× |
| Goccy | 128722 | 505.99 MB/s | 228156 | 134 | 3.4× |
| JSONV2 | 189166 | 344.31 MB/s | 206666 | 607 | 2.3× |
| Stdlib | 433207 | 150.35 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1599621 | 1213.08 MB/s | 2864594 | 1380 | 10.9× |
| Lightning | 1623890 | 1194.95 MB/s | 2864595 | 1380 | 10.8× |
| LightningArena | 1636909 | 1185.45 MB/s | 2864595 | 1380 | 10.7× |
| SonicFastest | 3036787 | 638.99 MB/s | 4885024 | 1736 | 5.8× |
| Sonic | 3059879 | 634.17 MB/s | 4883832 | 1736 | 5.7× |
| Goccy | 3612505 | 537.15 MB/s | 4064242 | 13509 | 4.8× |
| Easyjson | 5791900 | 335.03 MB/s | 3871265 | 15043 | 3.0× |
| LightningDecodeAny | 6627063 | 292.81 MB/s | 7063041 | 218633 | 2.6× |
| JSONV2 | 8088296 | 239.91 MB/s | 3237182 | 13947 | 2.2× |
| Stdlib | 17470543 | 111.07 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 697354 | 4772.08 MB/s | 351704 | 1286 | 22.1× |
| Lightning | 1098739 | 3028.77 MB/s | 2488905 | 2995 | 14.0× |
| LightningArena | 1099455 | 3026.80 MB/s | 2488906 | 2995 | 14.0× |
| SonicFastest | 1548200 | 2149.48 MB/s | 5892057 | 4263 | 10.0× |
| Sonic | 1562619 | 2129.65 MB/s | 5892240 | 4263 | 9.9× |
| LightningDecodeAny | 2391656 | 1285.20 MB/s | 4876913 | 56892 | 6.4× |
| Goccy | 3409990 | 975.91 MB/s | 3948918 | 3817 | 4.5× |
| JSONV2 | 5486331 | 606.57 MB/s | 5364501 | 13243 | 2.8× |
| Stdlib | 15405776 | 216.01 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 126298 | 1744.65 MB/s | 135872 | 226 | 11.6× |
| Lightning | 126707 | 1739.02 MB/s | 135872 | 226 | 11.6× |
| LightningDestructive | 128929 | 1709.04 MB/s | 135872 | 226 | 11.4× |
| Sonic | 259741 | 848.33 MB/s | 349775 | 262 | 5.7× |
| SonicFastest | 262378 | 839.80 MB/s | 349845 | 262 | 5.6× |
| Goccy | 342808 | 642.77 MB/s | 364234 | 1066 | 4.3× |
| Easyjson | 404005 | 545.40 MB/s | 130512 | 245 | 3.6× |
| JSONV2 | 502983 | 438.08 MB/s | 129746 | 470 | 2.9× |
| LightningDecodeAny | 695367 | 155.76 MB/s | 897217 | 11703 | 2.1× |
| Stdlib | 1471144 | 149.78 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 7449215 | 1087.37 MB/s | 11845072 | 20816 | 9.4× |
| Lightning | 7469650 | 1084.39 MB/s | 11845077 | 20816 | 9.4× |
| LightningArena | 7631613 | 1061.38 MB/s | 11845074 | 20816 | 9.2× |
| Sonic | 12039652 | 672.78 MB/s | 19868653 | 41640 | 5.8× |
| SonicFastest | 12246168 | 661.43 MB/s | 19869281 | 41640 | 5.7× |
| Goccy | 21414827 | 378.24 MB/s | 17888102 | 107150 | 3.3× |
| Easyjson | 25566784 | 316.82 MB/s | 15059617 | 41643 | 2.7× |
| LightningDecodeAny | 27927870 | 186.30 MB/s | 46279353 | 747112 | 2.5× |
| JSONV2 | 34632962 | 233.88 MB/s | 15233714 | 78972 | 2.0× |
| Stdlib | 69865295 | 115.94 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 3154635 | 945.74 MB/s | 3780458 | 1514 | 11.6× |
| Lightning | 3397947 | 878.02 MB/s | 3758857 | 29356 | 10.7× |
| LightningDestructive | 3401952 | 876.99 MB/s | 3758856 | 29356 | 10.7× |
| SonicFastest | 6402816 | 465.96 MB/s | 9133160 | 57804 | 5.7× |
| Sonic | 6567905 | 454.25 MB/s | 9132460 | 57804 | 5.6× |
| LightningDecodeAny | 11907387 | 154.04 MB/s | 23982579 | 351152 | 3.1× |
| Goccy | 13262742 | 224.95 MB/s | 9811212 | 273618 | 2.8× |
| Easyjson | 14206740 | 210.00 MB/s | 9479442 | 30115 | 2.6× |
| JSONV2 | 18827703 | 158.46 MB/s | 9257029 | 86278 | 1.9× |
| Stdlib | 36491774 | 81.76 MB/s | 9258083 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 742766 | 974.19 MB/s | 907600 | 3618 | 11.6× |
| LightningArena | 753046 | 960.89 MB/s | 916259 | 37 | 11.4× |
| Lightning | 779741 | 928.00 MB/s | 907598 | 3618 | 11.0× |
| Sonic | 1428162 | 506.66 MB/s | 2374929 | 3683 | 6.0× |
| SonicFastest | 1429719 | 506.11 MB/s | 2376914 | 3683 | 6.0× |
| LightningDecodeAny | 3323971 | 195.72 MB/s | 6500461 | 76546 | 2.6× |
| Easyjson | 3803174 | 190.26 MB/s | 2847912 | 3698 | 2.3× |
| Goccy | 3905444 | 185.28 MB/s | 2748527 | 80270 | 2.2× |
| JSONV2 | 4452946 | 162.50 MB/s | 2704714 | 7318 | 1.9× |
| Stdlib | 8614422 | 84.00 MB/s | 2704558 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1093468 | 1442.52 MB/s | 916257 | 37 | 10.6× |
| Lightning | 1116013 | 1413.38 MB/s | 907596 | 3618 | 10.4× |
| LightningDestructive | 1122463 | 1405.26 MB/s | 907600 | 3618 | 10.3× |
| SonicFastest | 1777529 | 887.39 MB/s | 3267245 | 3683 | 6.5× |
| Sonic | 1788255 | 882.06 MB/s | 3270096 | 3683 | 6.5× |
| LightningDecodeAny | 2928472 | 257.27 MB/s | 6500456 | 76546 | 3.9× |
| Easyjson | 4677155 | 337.25 MB/s | 2847907 | 3698 | 2.5× |
| Goccy | 4784954 | 329.65 MB/s | 3539643 | 80265 | 2.4× |
| JSONV2 | 5049242 | 312.39 MB/s | 2704552 | 7318 | 2.3× |
| Stdlib | 11553002 | 136.53 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 125597 | 1195.28 MB/s | 81920 | 1 | 10.6× |
| LightningArena | 128931 | 1164.37 MB/s | 81920 | 1 | 10.4× |
| LightningDestructive | 131400 | 1142.49 MB/s | 81920 | 1 | 10.2× |
| Sonic | 284896 | 526.94 MB/s | 408981 | 16 | 4.7× |
| SonicFastest | 303337 | 494.91 MB/s | 411249 | 16 | 4.4× |
| LightningDecodeAny | 359367 | 417.74 MB/s | 745764 | 10016 | 3.7× |
| Goccy | 744428 | 201.66 MB/s | 329900 | 10005 | 1.8× |
| JSONV2 | 832191 | 180.40 MB/s | 357726 | 20 | 1.6× |
| Stdlib | 1337082 | 112.28 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 21236 | 1324.01 MB/s | 29216 | 103 | 10.2× |
| LightningArena | 21384 | 1314.89 MB/s | 29216 | 103 | 10.1× |
| LightningDestructive | 22629 | 1242.50 MB/s | 29088 | 101 | 9.6× |
| SonicFastest | 36951 | 760.93 MB/s | 59506 | 83 | 5.9× |
| Sonic | 37316 | 753.48 MB/s | 59504 | 83 | 5.8× |
| Easyjson | 53502 | 525.53 MB/s | 32304 | 138 | 4.1× |
| Goccy | 57187 | 491.67 MB/s | 59297 | 188 | 3.8× |
| JSONV2 | 92077 | 305.36 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 109096 | 257.73 MB/s | 140576 | 2643 | 2.0× |
| Stdlib | 216878 | 129.64 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1166 | 1996.04 MB/s | 32 | 1 | 14.1× |
| LightningArena | 1171 | 1988.33 MB/s | 32 | 1 | 14.0× |
| LightningDestructive | 1220 | 1908.66 MB/s | 32 | 1 | 13.5× |
| SonicFastest | 2952 | 788.74 MB/s | 3718 | 4 | 5.6× |
| Sonic | 2993 | 777.76 MB/s | 3719 | 4 | 5.5× |
| Goccy | 3453 | 674.11 MB/s | 3649 | 4 | 4.8× |
| Easyjson | 3605 | 645.81 MB/s | 192 | 2 | 4.6× |
| JSONV2 | 5580 | 417.17 MB/s | 1000 | 6 | 2.9× |
| LightningDecodeAny | 7040 | 239.34 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 16443 | 141.58 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 130 | 1455.73 MB/s | 0 | 0 | 13.7× |
| Lightning | 131 | 1443.73 MB/s | 0 | 0 | 13.6× |
| LightningDestructive | 133 | 1424.12 MB/s | 0 | 0 | 13.4× |
| Goccy | 335 | 564.66 MB/s | 304 | 2 | 5.3× |
| SonicFastest | 350 | 540.35 MB/s | 343 | 3 | 5.1× |
| Sonic | 351 | 539.02 MB/s | 343 | 3 | 5.1× |
| Easyjson | 382 | 494.94 MB/s | 0 | 0 | 4.6× |
| JSONV2 | 679 | 278.46 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 865 | 154.95 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 1775 | 106.46 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 932 | 2350.42 MB/s | 0 | 0 | 12.9× |
| LightningDestructive | 943 | 2322.88 MB/s | 0 | 0 | 12.8× |
| Lightning | 948 | 2310.08 MB/s | 0 | 0 | 12.7× |
| Easyjson | 2300 | 952.57 MB/s | 24 | 1 | 5.2× |
| Goccy | 2859 | 766.28 MB/s | 2864 | 4 | 4.2× |
| SonicFastest | 4291 | 510.55 MB/s | 3607 | 38 | 2.8× |
| Sonic | 4398 | 498.18 MB/s | 3606 | 38 | 2.7× |
| JSONV2 | 5940 | 368.85 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 6171 | 293.47 MB/s | 7552 | 158 | 2.0× |
| Stdlib | 12069 | 181.54 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 364992 | 1398.59 MB/s | 457537 | 1009 | 12.6× |
| Lightning | 375171 | 1360.65 MB/s | 457537 | 1009 | 12.3× |
| LightningArena | 375214 | 1360.49 MB/s | 457537 | 1009 | 12.3× |
| Sonic | 792569 | 644.08 MB/s | 1308326 | 2014 | 5.8× |
| SonicFastest | 795627 | 641.60 MB/s | 1308290 | 2014 | 5.8× |
| Goccy | 840450 | 607.38 MB/s | 1142682 | 5006 | 5.5× |
| Easyjson | 1142966 | 446.62 MB/s | 863776 | 3012 | 4.0× |
| JSONV2 | 2365021 | 215.84 MB/s | 1075955 | 12645 | 1.9× |
| LightningDecodeAny | 2494590 | 184.99 MB/s | 2950651 | 64018 | 1.8× |
| Stdlib | 4611610 | 110.69 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 442 | 44743.34 MB/s | 0 | 0 | 173.0× |
| LightningArena | 447 | 44306.28 MB/s | 0 | 0 | 171.4× |
| LightningDestructive | 642 | 30820.47 MB/s | 0 | 0 | 119.2× |
| SonicFastest | 4617 | 4285.85 MB/s | 21102 | 3 | 16.6× |
| Goccy | 15750 | 1256.47 MB/s | 20492 | 2 | 4.9× |
| Sonic | 23788 | 831.88 MB/s | 20684 | 3 | 3.2× |
| JSONV2 | 25343 | 780.85 MB/s | 8 | 1 | 3.0× |
| Easyjson | 50954 | 388.37 MB/s | 0 | 0 | 1.5× |
| LightningDecodeAny | 69096 | 286.38 MB/s | 116864 | 2015 | 1.1× |
| Stdlib | 76536 | 258.56 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1490 | 12161.09 MB/s | 432 | 2 | 50.5× |
| Lightning | 1501 | 12077.60 MB/s | 432 | 2 | 50.1× |
| LightningDestructive | 1698 | 10672.89 MB/s | 0 | 0 | 44.3× |
| Easyjson | 3072 | 5900.56 MB/s | 432 | 2 | 24.5× |
| Sonic | 5369 | 3375.97 MB/s | 20416 | 5 | 14.0× |
| SonicFastest | 5419 | 3344.29 MB/s | 20427 | 5 | 13.9× |
| LightningDecodeAny | 13563 | 1318.46 MB/s | 29088 | 191 | 5.5× |
| Goccy | 14035 | 1291.36 MB/s | 19460 | 2 | 5.4× |
| JSONV2 | 33118 | 547.26 MB/s | 16501 | 50 | 2.3× |
| Stdlib | 75244 | 240.87 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1466155 | 1369.91 MB/s | 3089565 | 6821 | 9.2× |
| Lightning | 1577733 | 1273.03 MB/s | 3091278 | 6827 | 8.6× |
| LightningArena | 1577774 | 1272.99 MB/s | 3094394 | 6704 | 8.6× |
| Sonic | 2639031 | 761.07 MB/s | 5178584 | 7085 | 5.1× |
| SonicFastest | 2639105 | 761.05 MB/s | 5178543 | 7085 | 5.1× |
| Goccy | 3322161 | 604.57 MB/s | 5413783 | 15833 | 4.1× |
| Easyjson | 3808872 | 527.32 MB/s | 2981489 | 7439 | 3.6× |
| LightningDecodeAny | 5014349 | 227.80 MB/s | 8503512 | 134008 | 2.7× |
| JSONV2 | 5246856 | 382.80 MB/s | 3173678 | 14563 | 2.6× |
| Stdlib | 13554133 | 148.18 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 528 | 1039.37 MB/s | 480 | 1 | 7.9× |
| Lightning | 536 | 1023.46 MB/s | 480 | 1 | 7.7× |
| LightningDestructive | 557 | 985.19 MB/s | 480 | 1 | 7.4× |
| LightningDecodeAny | 1238 | 442.51 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 1485 | 369.60 MB/s | 1616 | 5 | 2.8× |
| Sonic | 1602 | 342.68 MB/s | 2265 | 8 | 2.6× |
| SonicFastest | 1622 | 338.42 MB/s | 2266 | 8 | 2.6× |
| JSONV2 | 2227 | 246.52 MB/s | 1664 | 7 | 1.9× |
| Goccy | 2231 | 246.05 MB/s | 2129 | 43 | 1.9× |
| Stdlib | 4148 | 132.34 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 315147 | 2003.87 MB/s | 402728 | 545 | 12.8× |
| LightningArena | 365502 | 1727.80 MB/s | 453041 | 713 | 11.0× |
| Lightning | 367837 | 1716.83 MB/s | 451257 | 857 | 11.0× |
| Sonic | 648376 | 973.99 MB/s | 1067351 | 814 | 6.2× |
| SonicFastest | 652113 | 968.41 MB/s | 1067647 | 814 | 6.2× |
| Easyjson | 866545 | 728.77 MB/s | 422505 | 936 | 4.7× |
| Goccy | 937825 | 673.38 MB/s | 988550 | 1201 | 4.3× |
| JSONV2 | 1557055 | 405.58 MB/s | 571592 | 3144 | 2.6× |
| LightningDecodeAny | 1716591 | 272.00 MB/s | 2076504 | 30126 | 2.4× |
| Stdlib | 4036990 | 156.43 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 419383 | 1341.04 MB/s | 546569 | 429 | 9.6× |
| LightningArena | 562500 | 999.84 MB/s | 771688 | 1089 | 7.1× |
| Lightning | 565764 | 994.07 MB/s | 769937 | 1235 | 7.1× |
| Sonic | 845216 | 665.40 MB/s | 1348256 | 1185 | 4.7× |
| SonicFastest | 851910 | 660.17 MB/s | 1348677 | 1185 | 4.7× |
| Goccy | 1173214 | 479.37 MB/s | 1035988 | 1028 | 3.4× |
| Easyjson | 1465429 | 383.78 MB/s | 775153 | 1254 | 2.7× |
| LightningDecodeAny | 2036408 | 276.18 MB/s | 2180441 | 30126 | 2.0× |
| JSONV2 | 2233920 | 251.76 MB/s | 927409 | 3482 | 1.8× |
| Stdlib | 4008834 | 140.29 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 401750 | 1327.14 MB/s | 333416 | 2084 | 10.2× |
| Lightning | 449868 | 1185.19 MB/s | 368224 | 2293 | 9.1× |
| LightningArena | 452923 | 1177.19 MB/s | 368225 | 2293 | 9.0× |
| SonicFastest | 692110 | 770.37 MB/s | 981928 | 3082 | 5.9× |
| Sonic | 694679 | 767.52 MB/s | 982542 | 3082 | 5.9× |
| Easyjson | 849988 | 627.28 MB/s | 428362 | 3273 | 4.8× |
| Goccy | 1031037 | 517.13 MB/s | 1167072 | 5409 | 4.0× |
| JSONV2 | 1852367 | 287.84 MB/s | 745427 | 13288 | 2.2× |
| LightningDecodeAny | 2482416 | 214.78 MB/s | 2992876 | 50076 | 1.6× |
| Stdlib | 4084872 | 130.53 MB/s | 798693 | 17133 | 1.0× |
