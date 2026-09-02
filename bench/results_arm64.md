# JSON Deserialization Benchmarks

- generated 2026-09-02T12:33:18Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 83715 | 1520.34 MB/s | 49760 | 3 | 13.1× |
| Lightning | 83810 | 1518.62 MB/s | 49760 | 3 | 13.1× |
| LightningDestructive | 83976 | 1515.61 MB/s | 49280 | 2 | 13.0× |
| Sonic | 181233 | 702.27 MB/s | 193107 | 10 | 6.0× |
| SonicFastest | 182855 | 696.04 MB/s | 195662 | 10 | 6.0× |
| Goccy | 197655 | 643.92 MB/s | 225372 | 884 | 5.5× |
| Easyjson | 211724 | 601.14 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 423671 | 300.41 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 440014 | 215.11 MB/s | 463409 | 9708 | 2.5× |
| Stdlib | 1095708 | 116.16 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2752091 | 817.94 MB/s | 2532848 | 1143 | 9.6× |
| LightningArena | 2775671 | 810.99 MB/s | 2532849 | 1143 | 9.6× |
| Lightning | 2832127 | 794.83 MB/s | 2532850 | 1143 | 9.4× |
| Sonic | 4648519 | 484.25 MB/s | 15237190 | 970 | 5.7× |
| SonicFastest | 4690494 | 479.92 MB/s | 15235570 | 970 | 5.7× |
| Goccy | 10566611 | 213.03 MB/s | 4126873 | 56533 | 2.5× |
| Easyjson | 10912386 | 206.28 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11649229 | 193.23 MB/s | 19380211 | 223896 | 2.3× |
| JSONV2 | 16021692 | 140.50 MB/s | 3123222 | 3083 | 1.7× |
| Stdlib | 26545966 | 84.80 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 372702 | 725.52 MB/s | 397296 | 567 | 9.2× |
| LightningArena | 374744 | 721.57 MB/s | 397296 | 567 | 9.1× |
| Lightning | 377025 | 717.20 MB/s | 397297 | 567 | 9.1× |
| Sonic | 633959 | 426.53 MB/s | 472861 | 968 | 5.4× |
| SonicFastest | 640254 | 422.34 MB/s | 491807 | 968 | 5.4× |
| Goccy | 1407452 | 192.12 MB/s | 543001 | 8122 | 2.4× |
| Easyjson | 1409864 | 191.79 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1568749 | 172.37 MB/s | 2543877 | 29687 | 2.2× |
| JSONV2 | 2083210 | 129.80 MB/s | 348153 | 1628 | 1.6× |
| Stdlib | 3425810 | 78.93 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 970215 | 1780.23 MB/s | 765601 | 2799 | 13.6× |
| LightningDestructive | 972032 | 1776.90 MB/s | 765560 | 2798 | 13.6× |
| LightningArena | 976234 | 1769.25 MB/s | 768416 | 2440 | 13.5× |
| Sonic | 2065068 | 836.39 MB/s | 2711564 | 4020 | 6.4× |
| SonicFastest | 2068968 | 834.81 MB/s | 2707633 | 4020 | 6.4× |
| Goccy | 2419968 | 713.73 MB/s | 2584166 | 14605 | 5.5× |
| Easyjson | 4238677 | 407.49 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4267067 | 404.78 MB/s | 1011639 | 7594 | 3.1× |
| LightningDecodeAny | 4420163 | 113.19 MB/s | 4953693 | 76576 | 3.0× |
| Stdlib | 13202051 | 130.83 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 871 | 2079.57 MB/s | 0 | 0 | 16.0× |
| LightningArena | 873 | 2075.84 MB/s | 0 | 0 | 16.0× |
| LightningDestructive | 885 | 2047.14 MB/s | 0 | 0 | 15.8× |
| Easyjson | 2545 | 712.00 MB/s | 24 | 1 | 5.5× |
| Goccy | 2875 | 630.32 MB/s | 2608 | 4 | 4.9× |
| Sonic | 5916 | 306.26 MB/s | 3715 | 40 | 2.4× |
| SonicFastest | 5958 | 304.12 MB/s | 3769 | 40 | 2.3× |
| JSONV2 | 7828 | 231.49 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7920 | 228.65 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13964 | 129.76 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 903 | 2007.35 MB/s | 0 | 0 | 15.5× |
| Lightning | 904 | 2004.78 MB/s | 0 | 0 | 15.4× |
| LightningDestructive | 937 | 1934.76 MB/s | 0 | 0 | 14.9× |
| Easyjson | 2544 | 712.37 MB/s | 24 | 1 | 5.5× |
| Goccy | 2799 | 647.42 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6005 | 301.72 MB/s | 3762 | 40 | 2.3× |
| Sonic | 6026 | 300.70 MB/s | 3803 | 40 | 2.3× |
| JSONV2 | 7691 | 235.59 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7893 | 229.43 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13952 | 129.87 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1084 | 1671.96 MB/s | 144 | 10 | 12.8× |
| LightningArena | 1089 | 1664.54 MB/s | 144 | 10 | 12.8× |
| LightningDestructive | 1131 | 1601.75 MB/s | 144 | 10 | 12.3× |
| Easyjson | 2768 | 654.68 MB/s | 144 | 10 | 5.0× |
| Goccy | 2874 | 630.43 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6244 | 290.21 MB/s | 3742 | 42 | 2.2× |
| SonicFastest | 6253 | 289.77 MB/s | 3768 | 42 | 2.2× |
| LightningDecodeAny | 7950 | 227.79 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7967 | 227.44 MB/s | 632 | 7 | 1.7× |
| Stdlib | 13924 | 130.13 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 650 | 759.57 MB/s | 160 | 1 | 8.4× |
| LightningDestructive | 654 | 754.79 MB/s | 160 | 1 | 8.4× |
| Sonic | 1257 | 393.14 MB/s | 975 | 6 | 4.4× |
| SonicFastest | 1258 | 392.57 MB/s | 993 | 6 | 4.4× |
| LightningArena | 1300 | 380.05 MB/s | 4096 | 1 | 4.2× |
| LightningDecodeAny | 1360 | 362.42 MB/s | 1296 | 26 | 4.0× |
| Easyjson | 2215 | 222.99 MB/s | 448 | 3 | 2.5× |
| Goccy | 2457 | 201.06 MB/s | 856 | 23 | 2.2× |
| JSONV2 | 3266 | 151.28 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5480 | 90.15 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 379 | 607.08 MB/s | 160 | 1 | 10.8× |
| LightningDestructive | 383 | 599.93 MB/s | 160 | 1 | 10.7× |
| Sonic | 898 | 256.15 MB/s | 656 | 6 | 4.6× |
| SonicFastest | 899 | 255.85 MB/s | 659 | 6 | 4.6× |
| LightningArena | 1052 | 218.63 MB/s | 4096 | 1 | 3.9× |
| LightningDecodeAny | 1129 | 202.77 MB/s | 1296 | 26 | 3.6× |
| Easyjson | 1405 | 163.71 MB/s | 448 | 3 | 2.9× |
| Goccy | 1583 | 145.33 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2441 | 94.24 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4096 | 56.16 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 51033 | 1276.26 MB/s | 97220 | 98 | 10.8× |
| LightningArena | 52054 | 1251.23 MB/s | 103440 | 103 | 10.6× |
| Lightning | 52208 | 1247.54 MB/s | 103440 | 103 | 10.5× |
| SonicFastest | 97017 | 671.34 MB/s | 155055 | 75 | 5.7× |
| Sonic | 102313 | 636.60 MB/s | 161599 | 75 | 5.4× |
| Goccy | 145900 | 446.41 MB/s | 229336 | 134 | 3.8× |
| LightningDecodeAny | 179507 | 297.09 MB/s | 180048 | 3245 | 3.1× |
| JSONV2 | 224207 | 290.50 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 549197 | 118.59 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2041327 | 950.59 MB/s | 2864592 | 1380 | 11.3× |
| Lightning | 2089523 | 928.67 MB/s | 2864594 | 1380 | 11.1× |
| LightningArena | 2115977 | 917.06 MB/s | 2864593 | 1380 | 10.9× |
| SonicFastest | 4697662 | 413.07 MB/s | 14608632 | 1407 | 4.9× |
| Sonic | 4700835 | 412.79 MB/s | 14608667 | 1407 | 4.9× |
| Goccy | 4800135 | 404.25 MB/s | 4064191 | 13510 | 4.8× |
| Easyjson | 7540108 | 257.35 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9541063 | 203.38 MB/s | 7063039 | 218633 | 2.4× |
| JSONV2 | 11288476 | 171.90 MB/s | 3237212 | 13947 | 2.0× |
| Stdlib | 23091955 | 84.03 MB/s | 3551321 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 892574 | 3728.35 MB/s | 351704 | 1286 | 23.4× |
| Lightning | 1502866 | 2214.32 MB/s | 2488905 | 2995 | 13.9× |
| LightningArena | 1509271 | 2204.93 MB/s | 2488904 | 2995 | 13.8× |
| Sonic | 2622075 | 1269.16 MB/s | 6528821 | 4248 | 8.0× |
| SonicFastest | 2625463 | 1267.52 MB/s | 6522813 | 4248 | 8.0× |
| LightningDecodeAny | 3517654 | 873.81 MB/s | 4876911 | 56892 | 5.9× |
| Goccy | 4470791 | 744.35 MB/s | 3948908 | 3816 | 4.7× |
| JSONV2 | 7374950 | 451.23 MB/s | 5364517 | 13243 | 2.8× |
| Stdlib | 20880950 | 159.37 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 183008 | 1204.02 MB/s | 135872 | 226 | 11.0× |
| LightningArena | 183536 | 1200.56 MB/s | 135872 | 226 | 11.0× |
| Lightning | 183734 | 1199.27 MB/s | 135872 | 226 | 11.0× |
| SonicFastest | 389403 | 565.86 MB/s | 322622 | 398 | 5.2× |
| Sonic | 396522 | 555.70 MB/s | 339167 | 398 | 5.1× |
| Goccy | 433788 | 507.96 MB/s | 364870 | 1067 | 4.7× |
| Easyjson | 549015 | 401.35 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 723963 | 304.36 MB/s | 129740 | 470 | 2.8× |
| LightningDecodeAny | 839965 | 128.95 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2020055 | 109.08 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9356220 | 865.74 MB/s | 11845073 | 20816 | 9.4× |
| Lightning | 9687018 | 836.17 MB/s | 11845077 | 20816 | 9.1× |
| LightningArena | 9701645 | 834.91 MB/s | 11845073 | 20816 | 9.1× |
| SonicFastest | 16518229 | 490.37 MB/s | 70901324 | 40014 | 5.3× |
| Sonic | 16547247 | 489.51 MB/s | 70901242 | 40014 | 5.3× |
| Goccy | 23834060 | 339.85 MB/s | 16981379 | 107148 | 3.7× |
| Easyjson | 30895246 | 262.18 MB/s | 15059619 | 41643 | 2.8× |
| LightningDecodeAny | 35254421 | 147.59 MB/s | 46279371 | 747112 | 2.5× |
| JSONV2 | 44123351 | 183.58 MB/s | 15233744 | 78972 | 2.0× |
| Stdlib | 87910659 | 92.14 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4593295 | 649.53 MB/s | 3764714 | 1504 | 10.1× |
| LightningDestructive | 4949908 | 602.73 MB/s | 3758856 | 29356 | 9.4× |
| Lightning | 5080245 | 587.27 MB/s | 3758856 | 29356 | 9.2× |
| Sonic | 8743391 | 341.23 MB/s | 26542347 | 56760 | 5.3× |
| SonicFastest | 8762841 | 340.47 MB/s | 26511436 | 56760 | 5.3× |
| LightningDecodeAny | 16047714 | 114.30 MB/s | 23982580 | 351152 | 2.9× |
| Easyjson | 16598438 | 179.74 MB/s | 9479441 | 30115 | 2.8× |
| Goccy | 16725176 | 178.38 MB/s | 10574298 | 273647 | 2.8× |
| JSONV2 | 24881968 | 119.90 MB/s | 9257171 | 86278 | 1.9× |
| Stdlib | 46602625 | 64.02 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1009405 | 716.85 MB/s | 911393 | 30 | 11.3× |
| LightningDestructive | 1057202 | 684.45 MB/s | 907601 | 3618 | 10.8× |
| Lightning | 1063344 | 680.49 MB/s | 907595 | 3618 | 10.7× |
| Sonic | 1772764 | 408.17 MB/s | 3182194 | 7226 | 6.4× |
| SonicFastest | 1776675 | 407.28 MB/s | 3191324 | 7226 | 6.4× |
| LightningDecodeAny | 4016250 | 161.99 MB/s | 6500457 | 76546 | 2.8× |
| Easyjson | 4189769 | 172.71 MB/s | 2847906 | 3698 | 2.7× |
| Goccy | 4835964 | 149.63 MB/s | 2827329 | 80274 | 2.4× |
| JSONV2 | 5891454 | 122.82 MB/s | 2704627 | 7318 | 1.9× |
| Stdlib | 11391469 | 63.52 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1579763 | 998.47 MB/s | 911393 | 30 | 9.8× |
| LightningDestructive | 1586091 | 994.49 MB/s | 907600 | 3618 | 9.8× |
| Lightning | 1635346 | 964.54 MB/s | 907594 | 3618 | 9.5× |
| Sonic | 2254839 | 699.54 MB/s | 5782439 | 7226 | 6.9× |
| SonicFastest | 2261098 | 697.60 MB/s | 5789542 | 7226 | 6.9× |
| LightningDecodeAny | 3669188 | 205.33 MB/s | 6500459 | 76546 | 4.2× |
| Easyjson | 5589252 | 282.21 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5612692 | 281.03 MB/s | 3563987 | 80266 | 2.8× |
| JSONV2 | 6560001 | 240.45 MB/s | 2704592 | 7318 | 2.4× |
| Stdlib | 15535818 | 101.53 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 157120 | 955.48 MB/s | 81920 | 1 | 11.8× |
| Lightning | 157567 | 952.76 MB/s | 81920 | 1 | 11.8× |
| LightningArena | 157617 | 952.46 MB/s | 81920 | 1 | 11.8× |
| SonicFastest | 271525 | 552.89 MB/s | 246692 | 6 | 6.8× |
| Sonic | 271969 | 551.99 MB/s | 250317 | 6 | 6.8× |
| LightningDecodeAny | 427153 | 351.45 MB/s | 745765 | 10016 | 4.3× |
| Goccy | 859055 | 174.75 MB/s | 324387 | 10004 | 2.2× |
| JSONV2 | 1103127 | 136.09 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1857301 | 80.83 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 27153 | 1035.50 MB/s | 29216 | 103 | 11.1× |
| LightningDestructive | 27403 | 1026.05 MB/s | 29088 | 101 | 11.0× |
| Lightning | 27477 | 1023.30 MB/s | 29216 | 103 | 11.0× |
| Sonic | 63734 | 441.16 MB/s | 47175 | 103 | 4.7× |
| SonicFastest | 63794 | 440.75 MB/s | 47053 | 103 | 4.7× |
| Easyjson | 68424 | 410.92 MB/s | 32304 | 138 | 4.4× |
| Goccy | 71292 | 394.39 MB/s | 59217 | 188 | 4.2× |
| JSONV2 | 134287 | 209.38 MB/s | 36895 | 242 | 2.2× |
| LightningDecodeAny | 147637 | 190.45 MB/s | 140576 | 2643 | 2.0× |
| Stdlib | 301160 | 93.36 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1510 | 1541.34 MB/s | 32 | 1 | 14.9× |
| Lightning | 1522 | 1529.77 MB/s | 32 | 1 | 14.8× |
| LightningDestructive | 1589 | 1465.37 MB/s | 32 | 1 | 14.2× |
| Goccy | 4121 | 564.86 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4208 | 553.25 MB/s | 192 | 2 | 5.3× |
| SonicFastest | 5066 | 459.51 MB/s | 4208 | 6 | 4.4× |
| Sonic | 5084 | 457.93 MB/s | 4238 | 6 | 4.4× |
| JSONV2 | 8476 | 274.67 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9893 | 170.32 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22503 | 103.45 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 182 | 1038.68 MB/s | 0 | 0 | 13.3× |
| Lightning | 182 | 1037.59 MB/s | 0 | 0 | 13.3× |
| LightningDestructive | 183 | 1030.33 MB/s | 0 | 0 | 13.2× |
| Goccy | 384 | 491.96 MB/s | 304 | 2 | 6.3× |
| Easyjson | 494 | 382.94 MB/s | 0 | 0 | 4.9× |
| SonicFastest | 805 | 234.68 MB/s | 499 | 4 | 3.0× |
| Sonic | 809 | 233.68 MB/s | 505 | 4 | 3.0× |
| JSONV2 | 1036 | 182.48 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1210 | 110.73 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2429 | 77.82 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1189 | 1843.22 MB/s | 0 | 0 | 13.3× |
| LightningArena | 1200 | 1825.58 MB/s | 0 | 0 | 13.1× |
| LightningDestructive | 1215 | 1803.09 MB/s | 0 | 0 | 13.0× |
| Goccy | 3161 | 693.14 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3183 | 688.29 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6412 | 341.71 MB/s | 3999 | 40 | 2.5× |
| Sonic | 6423 | 341.11 MB/s | 3995 | 40 | 2.5× |
| JSONV2 | 7922 | 276.56 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 7969 | 227.25 MB/s | 7552 | 158 | 2.0× |
| Stdlib | 15764 | 138.98 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 556939 | 916.57 MB/s | 457536 | 1009 | 10.8× |
| LightningDestructive | 557041 | 916.41 MB/s | 457536 | 1009 | 10.8× |
| Lightning | 559939 | 911.66 MB/s | 457537 | 1009 | 10.7× |
| Sonic | 1161843 | 439.37 MB/s | 901471 | 2006 | 5.2× |
| Goccy | 1165375 | 438.04 MB/s | 1139792 | 5006 | 5.1× |
| SonicFastest | 1173148 | 435.13 MB/s | 912594 | 2006 | 5.1× |
| Easyjson | 1548137 | 329.74 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3266303 | 156.29 MB/s | 1076010 | 12646 | 1.8× |
| LightningDecodeAny | 3351793 | 137.68 MB/s | 2950648 | 64018 | 1.8× |
| Stdlib | 5994946 | 85.15 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 482 | 41026.77 MB/s | 0 | 0 | 224.8× |
| Lightning | 483 | 40993.08 MB/s | 0 | 0 | 224.6× |
| LightningDestructive | 501 | 39505.55 MB/s | 0 | 0 | 216.4× |
| Goccy | 20053 | 986.83 MB/s | 20491 | 2 | 5.4× |
| SonicFastest | 27932 | 708.47 MB/s | 22278 | 4 | 3.9× |
| Sonic | 27956 | 707.87 MB/s | 22311 | 4 | 3.9× |
| JSONV2 | 29573 | 669.16 MB/s | 8 | 1 | 3.7× |
| LightningDecodeAny | 81328 | 243.31 MB/s | 116864 | 2015 | 1.3× |
| Easyjson | 81972 | 241.41 MB/s | 0 | 0 | 1.3× |
| Stdlib | 108399 | 182.56 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1742 | 10402.34 MB/s | 0 | 0 | 59.0× |
| Lightning | 1877 | 9654.85 MB/s | 432 | 2 | 54.7× |
| LightningArena | 1882 | 9629.38 MB/s | 432 | 2 | 54.6× |
| Easyjson | 3954 | 4583.30 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 9904 | 1829.91 MB/s | 23003 | 6 | 10.4× |
| Sonic | 10057 | 1802.08 MB/s | 23219 | 6 | 10.2× |
| Goccy | 16044 | 1129.65 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 16186 | 1104.78 MB/s | 29088 | 191 | 6.3× |
| JSONV2 | 47498 | 381.58 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102694 | 176.49 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2128344 | 943.69 MB/s | 3089565 | 6821 | 8.7× |
| Lightning | 2159678 | 930.00 MB/s | 3091277 | 6827 | 8.6× |
| LightningArena | 2178407 | 922.00 MB/s | 3094371 | 6703 | 8.5× |
| Goccy | 4283090 | 468.94 MB/s | 5411343 | 15830 | 4.3× |
| Sonic | 4420386 | 454.37 MB/s | 10949762 | 13683 | 4.2× |
| SonicFastest | 4448541 | 451.49 MB/s | 10987439 | 13683 | 4.2× |
| Easyjson | 4922425 | 408.03 MB/s | 2981480 | 7439 | 3.8× |
| JSONV2 | 6964252 | 288.40 MB/s | 3173684 | 14563 | 2.7× |
| LightningDecodeAny | 7510322 | 152.10 MB/s | 8503512 | 134008 | 2.5× |
| Stdlib | 18529726 | 108.39 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 853 | 643.47 MB/s | 480 | 1 | 6.6× |
| LightningArena | 855 | 641.99 MB/s | 480 | 1 | 6.6× |
| LightningDestructive | 862 | 636.96 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1650 | 332.10 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 2160 | 254.15 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2679 | 204.94 MB/s | 1932 | 26 | 2.1× |
| SonicFastest | 2694 | 203.77 MB/s | 1945 | 26 | 2.1× |
| Goccy | 3007 | 182.56 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3305 | 166.13 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5611 | 97.84 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 418183 | 1510.14 MB/s | 402729 | 545 | 12.8× |
| Lightning | 486064 | 1299.24 MB/s | 451257 | 857 | 11.0× |
| LightningArena | 487848 | 1294.49 MB/s | 453017 | 712 | 11.0× |
| SonicFastest | 1019146 | 619.65 MB/s | 985760 | 1102 | 5.3× |
| Sonic | 1019216 | 619.61 MB/s | 989328 | 1102 | 5.3× |
| Easyjson | 1143302 | 552.36 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1180177 | 535.10 MB/s | 985926 | 1201 | 4.5× |
| JSONV2 | 2150959 | 293.60 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2354864 | 198.27 MB/s | 2076504 | 30126 | 2.3× |
| Stdlib | 5358249 | 117.86 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 547475 | 1027.28 MB/s | 546571 | 429 | 9.6× |
| Lightning | 711685 | 790.25 MB/s | 769938 | 1235 | 7.4× |
| LightningArena | 715410 | 786.13 MB/s | 771665 | 1088 | 7.3× |
| Sonic | 1024942 | 548.72 MB/s | 953708 | 1476 | 5.1× |
| SonicFastest | 1034750 | 543.52 MB/s | 964622 | 1476 | 5.1× |
| Goccy | 1345403 | 418.02 MB/s | 1040418 | 1030 | 3.9× |
| Easyjson | 1761989 | 319.19 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2570322 | 218.81 MB/s | 2180440 | 30126 | 2.0× |
| JSONV2 | 2751355 | 204.41 MB/s | 927441 | 3482 | 1.9× |
| Stdlib | 5245142 | 107.22 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 573969 | 928.93 MB/s | 333416 | 2084 | 9.5× |
| LightningArena | 596227 | 894.25 MB/s | 368224 | 2293 | 9.1× |
| Lightning | 597391 | 892.51 MB/s | 368224 | 2293 | 9.1× |
| Easyjson | 1112457 | 479.28 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1157240 | 460.73 MB/s | 1039499 | 4351 | 4.7× |
| Sonic | 1165772 | 457.36 MB/s | 1055508 | 4351 | 4.7× |
| Goccy | 1302436 | 409.37 MB/s | 1167226 | 5409 | 4.2× |
| JSONV2 | 2550865 | 209.02 MB/s | 745447 | 13288 | 2.1× |
| LightningDecodeAny | 3360814 | 158.65 MB/s | 2992878 | 50076 | 1.6× |
| Stdlib | 5430510 | 98.18 MB/s | 798692 | 17133 | 1.0× |
