# JSON Deserialization Benchmarks

- generated 2026-09-23T18:19:20Z
- go version go1.26.8 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 81179 | 1567.84 MB/s | 49280 | 2 | 13.5× |
| LightningArena | 81444 | 1562.73 MB/s | 49819 | 2 | 13.4× |
| Lightning | 81504 | 1561.58 MB/s | 49821 | 2 | 13.4× |
| SonicFastest | 182516 | 697.34 MB/s | 195984 | 10 | 6.0× |
| Sonic | 183604 | 693.20 MB/s | 196766 | 10 | 5.9× |
| Goccy | 192187 | 662.25 MB/s | 224672 | 884 | 5.7× |
| Easyjson | 212690 | 598.41 MB/s | 122864 | 14 | 5.1× |
| JSONV2 | 420164 | 302.92 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 428576 | 220.85 MB/s | 464118 | 9706 | 2.5× |
| Stdlib | 1091870 | 116.57 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1581272 | 1423.57 MB/s | 2532848 | 1143 | 16.8× |
| Lightning | 1612190 | 1396.27 MB/s | 2532850 | 1143 | 16.4× |
| LightningArena | 2222372 | 1012.90 MB/s | 2532848 | 1143 | 11.9× |
| Sonic | 4475833 | 502.93 MB/s | 15233735 | 970 | 5.9× |
| SonicFastest | 4484260 | 501.99 MB/s | 15233735 | 970 | 5.9× |
| LightningDecodeAny | 8363093 | 269.16 MB/s | 6828999 | 223498 | 3.2× |
| Goccy | 10610172 | 212.16 MB/s | 4129952 | 56533 | 2.5× |
| Easyjson | 11171504 | 201.50 MB/s | 3099808 | 2120 | 2.4× |
| JSONV2 | 16185277 | 139.08 MB/s | 3123214 | 3083 | 1.6× |
| Stdlib | 26516727 | 84.89 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 257304 | 1050.91 MB/s | 397296 | 567 | 13.3× |
| LightningDestructive | 258217 | 1047.19 MB/s | 397296 | 567 | 13.3× |
| LightningArena | 339399 | 796.71 MB/s | 397296 | 567 | 10.1× |
| Sonic | 626784 | 431.41 MB/s | 469603 | 968 | 5.5× |
| SonicFastest | 637938 | 423.87 MB/s | 485040 | 968 | 5.4× |
| LightningDecodeAny | 1190981 | 227.04 MB/s | 845690 | 29656 | 2.9× |
| Goccy | 1429070 | 189.22 MB/s | 543488 | 8122 | 2.4× |
| Easyjson | 1432826 | 188.72 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2108043 | 128.27 MB/s | 348160 | 1628 | 1.6× |
| Stdlib | 3428144 | 78.88 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 925331 | 1866.58 MB/s | 765560 | 2798 | 14.3× |
| Lightning | 940432 | 1836.61 MB/s | 767766 | 2798 | 14.0× |
| LightningArena | 941789 | 1833.96 MB/s | 775122 | 2444 | 14.0× |
| Sonic | 2060514 | 838.24 MB/s | 2703533 | 4020 | 6.4× |
| SonicFastest | 2062699 | 837.35 MB/s | 2698442 | 4020 | 6.4× |
| Goccy | 2405371 | 718.06 MB/s | 2582861 | 14605 | 5.5× |
| LightningDecodeAny | 4029258 | 124.17 MB/s | 4492673 | 67881 | 3.3× |
| Easyjson | 4218693 | 409.42 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4278179 | 403.72 MB/s | 1011634 | 7594 | 3.1× |
| Stdlib | 13189830 | 130.95 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 829 | 2186.32 MB/s | 0 | 0 | 16.8× |
| Lightning | 836 | 2167.68 MB/s | 0 | 0 | 16.6× |
| LightningDestructive | 846 | 2141.05 MB/s | 0 | 0 | 16.4× |
| Easyjson | 2542 | 712.83 MB/s | 24 | 1 | 5.5× |
| Goccy | 2787 | 650.20 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6021 | 300.96 MB/s | 3872 | 40 | 2.3× |
| Sonic | 6029 | 300.55 MB/s | 3837 | 40 | 2.3× |
| JSONV2 | 7758 | 233.55 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8064 | 224.58 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13905 | 130.31 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 846 | 2143.14 MB/s | 0 | 0 | 16.4× |
| LightningArena | 851 | 2129.32 MB/s | 0 | 0 | 16.3× |
| LightningDestructive | 873 | 2074.77 MB/s | 0 | 0 | 15.9× |
| Easyjson | 2550 | 710.49 MB/s | 24 | 1 | 5.5× |
| Goccy | 2817 | 643.31 MB/s | 2608 | 4 | 4.9× |
| Sonic | 5952 | 304.45 MB/s | 3727 | 40 | 2.3× |
| SonicFastest | 6028 | 300.60 MB/s | 3812 | 40 | 2.3× |
| LightningDecodeAny | 7760 | 233.38 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7889 | 229.67 MB/s | 640 | 6 | 1.8× |
| Stdlib | 13905 | 130.31 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1019 | 1778.22 MB/s | 144 | 10 | 13.6× |
| LightningArena | 1025 | 1766.95 MB/s | 144 | 10 | 13.6× |
| LightningDestructive | 1083 | 1673.02 MB/s | 144 | 10 | 12.8× |
| Easyjson | 2773 | 653.53 MB/s | 144 | 10 | 5.0× |
| Goccy | 2861 | 633.26 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6121 | 296.01 MB/s | 3822 | 42 | 2.3× |
| SonicFastest | 6125 | 295.84 MB/s | 3815 | 42 | 2.3× |
| LightningDecodeAny | 7837 | 231.08 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7975 | 227.22 MB/s | 632 | 7 | 1.7× |
| Stdlib | 13909 | 130.28 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 636 | 776.24 MB/s | 160 | 1 | 8.6× |
| LightningDestructive | 638 | 773.89 MB/s | 160 | 1 | 8.6× |
| LightningDecodeAny | 1197 | 411.86 MB/s | 1040 | 25 | 4.6× |
| SonicFastest | 1235 | 400.09 MB/s | 986 | 6 | 4.4× |
| Sonic | 1238 | 399.18 MB/s | 998 | 6 | 4.4× |
| LightningArena | 1368 | 361.21 MB/s | 4120 | 2 | 4.0× |
| Easyjson | 2222 | 222.34 MB/s | 448 | 3 | 2.5× |
| Goccy | 2441 | 202.38 MB/s | 856 | 23 | 2.2× |
| JSONV2 | 3223 | 153.29 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5462 | 90.44 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 300 | 766.87 MB/s | 160 | 1 | 13.6× |
| Lightning | 303 | 758.59 MB/s | 160 | 1 | 13.5× |
| Sonic | 880 | 261.44 MB/s | 655 | 6 | 4.6× |
| SonicFastest | 882 | 260.79 MB/s | 649 | 6 | 4.6× |
| LightningDecodeAny | 1021 | 224.26 MB/s | 1040 | 25 | 4.0× |
| LightningArena | 1074 | 214.18 MB/s | 4120 | 2 | 3.8× |
| Easyjson | 1398 | 164.58 MB/s | 448 | 3 | 2.9× |
| Goccy | 1581 | 145.45 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2421 | 95.02 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4086 | 56.29 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 49560 | 1314.22 MB/s | 97220 | 98 | 11.0× |
| LightningArena | 50946 | 1278.44 MB/s | 103684 | 99 | 10.7× |
| Lightning | 51622 | 1261.72 MB/s | 103684 | 99 | 10.6× |
| SonicFastest | 97332 | 669.17 MB/s | 155152 | 75 | 5.6× |
| Sonic | 97571 | 667.54 MB/s | 156003 | 75 | 5.6× |
| Goccy | 141965 | 458.79 MB/s | 229291 | 134 | 3.8× |
| LightningDecodeAny | 173943 | 306.59 MB/s | 176544 | 3237 | 3.1× |
| JSONV2 | 224600 | 289.99 MB/s | 206652 | 607 | 2.4× |
| Stdlib | 546074 | 119.27 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2092506 | 927.34 MB/s | 2185296 | 1350 | 11.0× |
| LightningArena | 2162513 | 897.32 MB/s | 2185297 | 1350 | 10.7× |
| Lightning | 2168482 | 894.85 MB/s | 2185298 | 1350 | 10.6× |
| Sonic | 4725627 | 410.63 MB/s | 14608621 | 1407 | 4.9× |
| SonicFastest | 4743931 | 409.04 MB/s | 14606973 | 1407 | 4.9× |
| Goccy | 4761939 | 407.50 MB/s | 4064764 | 13510 | 4.8× |
| Easyjson | 7515537 | 258.19 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 8955194 | 216.69 MB/s | 6627983 | 206416 | 2.6× |
| JSONV2 | 11306115 | 171.63 MB/s | 3237231 | 13947 | 2.0× |
| Stdlib | 23092754 | 84.03 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 876084 | 3798.53 MB/s | 351704 | 1286 | 24.0× |
| Lightning | 1310686 | 2539.00 MB/s | 2434281 | 1413 | 16.1× |
| LightningArena | 1319544 | 2521.96 MB/s | 2434285 | 1413 | 15.9× |
| Sonic | 2655210 | 1253.32 MB/s | 6472621 | 4248 | 7.9× |
| SonicFastest | 2699235 | 1232.88 MB/s | 6458598 | 4248 | 7.8× |
| LightningDecodeAny | 3237950 | 949.29 MB/s | 4825392 | 55311 | 6.5× |
| Goccy | 4479493 | 742.90 MB/s | 3948908 | 3816 | 4.7× |
| JSONV2 | 7502377 | 443.57 MB/s | 5364513 | 13243 | 2.8× |
| Stdlib | 21041960 | 158.15 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 173577 | 1269.44 MB/s | 135392 | 226 | 11.6× |
| Lightning | 173891 | 1267.15 MB/s | 135392 | 226 | 11.6× |
| LightningDestructive | 175651 | 1254.45 MB/s | 135392 | 226 | 11.5× |
| SonicFastest | 379143 | 581.17 MB/s | 303345 | 398 | 5.3× |
| Sonic | 386735 | 569.76 MB/s | 324117 | 398 | 5.2× |
| Goccy | 447333 | 492.58 MB/s | 365080 | 1067 | 4.5× |
| Easyjson | 549855 | 400.73 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 732145 | 300.96 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 829282 | 130.61 MB/s | 854737 | 11700 | 2.4× |
| Stdlib | 2017183 | 109.23 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 8554482 | 946.88 MB/s | 8109648 | 20809 | 10.3× |
| Lightning | 8577462 | 944.34 MB/s | 8109649 | 20809 | 10.3× |
| LightningArena | 9780909 | 828.15 MB/s | 8109653 | 20809 | 9.1× |
| SonicFastest | 17378003 | 466.11 MB/s | 70888004 | 40014 | 5.1× |
| Sonic | 17415503 | 465.11 MB/s | 70887849 | 40014 | 5.1× |
| Goccy | 23426917 | 345.76 MB/s | 16775216 | 107147 | 3.8× |
| LightningDecodeAny | 30478956 | 170.71 MB/s | 28359832 | 746961 | 2.9× |
| Easyjson | 30971923 | 261.53 MB/s | 15059618 | 41643 | 2.9× |
| JSONV2 | 43679725 | 185.44 MB/s | 15233748 | 78972 | 2.0× |
| Stdlib | 88526880 | 91.50 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 3331233 | 895.60 MB/s | 3780457 | 1514 | 14.0× |
| LightningDestructive | 3507797 | 850.52 MB/s | 3758857 | 29356 | 13.3× |
| Lightning | 3619545 | 824.27 MB/s | 3758859 | 29356 | 12.8× |
| SonicFastest | 8678807 | 343.76 MB/s | 26607303 | 56760 | 5.4× |
| Sonic | 8698300 | 342.99 MB/s | 26517296 | 56760 | 5.3× |
| LightningDecodeAny | 14777347 | 124.12 MB/s | 18225373 | 350883 | 3.1× |
| Goccy | 16678642 | 178.88 MB/s | 10618950 | 273649 | 2.8× |
| Easyjson | 16877688 | 176.77 MB/s | 9479440 | 30115 | 2.8× |
| JSONV2 | 24822852 | 120.19 MB/s | 9257175 | 86278 | 1.9× |
| Stdlib | 46511068 | 64.15 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 684557 | 1057.03 MB/s | 907601 | 3618 | 16.6× |
| LightningArena | 707761 | 1022.38 MB/s | 916260 | 37 | 16.1× |
| Lightning | 733610 | 986.35 MB/s | 907599 | 3618 | 15.5× |
| SonicFastest | 1759080 | 411.35 MB/s | 3179625 | 7226 | 6.5× |
| Sonic | 1765812 | 409.78 MB/s | 3181412 | 7226 | 6.4× |
| LightningDecodeAny | 3917557 | 166.07 MB/s | 5691591 | 76540 | 2.9× |
| Easyjson | 4168679 | 173.58 MB/s | 2847906 | 3698 | 2.7× |
| Goccy | 4821153 | 150.09 MB/s | 2779574 | 80272 | 2.4× |
| JSONV2 | 5553320 | 130.30 MB/s | 2704658 | 7318 | 2.0× |
| Stdlib | 11382533 | 63.57 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1012134 | 1558.44 MB/s | 907600 | 3618 | 15.4× |
| LightningArena | 1012292 | 1558.20 MB/s | 916257 | 37 | 15.4× |
| Lightning | 1051243 | 1500.46 MB/s | 907594 | 3618 | 14.8× |
| Sonic | 2219322 | 710.74 MB/s | 5789360 | 7226 | 7.0× |
| SonicFastest | 2247924 | 701.69 MB/s | 5797792 | 7226 | 6.9× |
| LightningDecodeAny | 3463291 | 217.54 MB/s | 5691592 | 76540 | 4.5× |
| Easyjson | 5595731 | 281.89 MB/s | 2847909 | 3698 | 2.8× |
| Goccy | 5645602 | 279.40 MB/s | 3590883 | 80267 | 2.8× |
| JSONV2 | 6388670 | 246.90 MB/s | 2704588 | 7318 | 2.4× |
| Stdlib | 15551421 | 101.43 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 72468 | 2071.59 MB/s | 81920 | 1 | 25.6× |
| LightningArena | 73063 | 2054.72 MB/s | 81920 | 1 | 25.4× |
| LightningDestructive | 73094 | 2053.84 MB/s | 81920 | 1 | 25.4× |
| SonicFastest | 271736 | 552.46 MB/s | 250125 | 6 | 6.8× |
| Sonic | 272937 | 550.03 MB/s | 253253 | 6 | 6.8× |
| LightningDecodeAny | 430418 | 348.78 MB/s | 745508 | 10015 | 4.3× |
| Goccy | 868538 | 172.85 MB/s | 324954 | 10004 | 2.1× |
| JSONV2 | 1072574 | 139.97 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1855127 | 80.92 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 26881 | 1045.99 MB/s | 29088 | 101 | 11.2× |
| LightningArena | 27097 | 1037.64 MB/s | 29236 | 101 | 11.1× |
| Lightning | 27135 | 1036.21 MB/s | 29235 | 101 | 11.1× |
| Sonic | 63102 | 445.58 MB/s | 46619 | 103 | 4.8× |
| SonicFastest | 63434 | 443.25 MB/s | 47118 | 103 | 4.7× |
| Easyjson | 68445 | 410.79 MB/s | 32304 | 138 | 4.4× |
| Goccy | 71056 | 395.70 MB/s | 59206 | 188 | 4.2× |
| JSONV2 | 133870 | 210.03 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 144036 | 195.21 MB/s | 133685 | 2639 | 2.1× |
| Stdlib | 300953 | 93.43 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1501 | 1551.29 MB/s | 32 | 1 | 15.0× |
| Lightning | 1503 | 1549.18 MB/s | 32 | 1 | 14.9× |
| LightningDestructive | 1569 | 1483.89 MB/s | 32 | 1 | 14.3× |
| Goccy | 4159 | 559.79 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4201 | 554.20 MB/s | 192 | 2 | 5.3× |
| Sonic | 5160 | 451.17 MB/s | 4339 | 6 | 4.4× |
| SonicFastest | 5161 | 451.08 MB/s | 4346 | 6 | 4.4× |
| JSONV2 | 8477 | 274.62 MB/s | 1000 | 6 | 2.6× |
| LightningDecodeAny | 9626 | 175.04 MB/s | 9936 | 194 | 2.3× |
| Stdlib | 22463 | 103.64 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 179 | 1056.56 MB/s | 0 | 0 | 13.4× |
| LightningArena | 180 | 1052.48 MB/s | 0 | 0 | 13.4× |
| LightningDestructive | 182 | 1039.10 MB/s | 0 | 0 | 13.2× |
| Goccy | 386 | 489.62 MB/s | 304 | 2 | 6.2× |
| Easyjson | 483 | 391.59 MB/s | 0 | 0 | 5.0× |
| Sonic | 796 | 237.44 MB/s | 513 | 4 | 3.0× |
| SonicFastest | 803 | 235.32 MB/s | 513 | 4 | 3.0× |
| JSONV2 | 1035 | 182.54 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1231 | 108.89 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2398 | 78.80 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1093 | 2005.28 MB/s | 0 | 0 | 14.4× |
| LightningArena | 1093 | 2004.26 MB/s | 0 | 0 | 14.4× |
| LightningDestructive | 1117 | 1960.67 MB/s | 0 | 0 | 14.1× |
| Easyjson | 3191 | 686.58 MB/s | 24 | 1 | 4.9× |
| Goccy | 3203 | 684.11 MB/s | 2864 | 4 | 4.9× |
| Sonic | 6312 | 347.12 MB/s | 3950 | 40 | 2.5× |
| SonicFastest | 6348 | 345.16 MB/s | 3996 | 40 | 2.5× |
| LightningDecodeAny | 7865 | 230.25 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 8091 | 270.80 MB/s | 640 | 6 | 1.9× |
| Stdlib | 15776 | 138.88 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 537952 | 948.92 MB/s | 318400 | 1005 | 11.1× |
| Lightning | 540932 | 943.70 MB/s | 318400 | 1005 | 11.0× |
| LightningArena | 542532 | 940.91 MB/s | 318400 | 1005 | 11.0× |
| Goccy | 1144176 | 446.15 MB/s | 1137547 | 5006 | 5.2× |
| SonicFastest | 1159341 | 440.32 MB/s | 882373 | 2006 | 5.1× |
| Sonic | 1163101 | 438.89 MB/s | 886571 | 2006 | 5.1× |
| Easyjson | 1529159 | 333.83 MB/s | 863778 | 3012 | 3.9× |
| LightningDecodeAny | 3201537 | 144.14 MB/s | 2742391 | 64017 | 1.9× |
| JSONV2 | 3292154 | 155.06 MB/s | 1076010 | 12646 | 1.8× |
| Stdlib | 5952415 | 85.76 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 461 | 42932.53 MB/s | 0 | 0 | 235.0× |
| LightningArena | 462 | 42878.71 MB/s | 0 | 0 | 234.7× |
| LightningDestructive | 474 | 41725.88 MB/s | 0 | 0 | 228.3× |
| Goccy | 19946 | 992.13 MB/s | 20491 | 2 | 5.4× |
| SonicFastest | 27217 | 727.09 MB/s | 22488 | 4 | 4.0× |
| Sonic | 27253 | 726.11 MB/s | 22708 | 4 | 4.0× |
| JSONV2 | 29776 | 664.60 MB/s | 8 | 1 | 3.6× |
| Easyjson | 82211 | 240.71 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 86191 | 229.58 MB/s | 116608 | 2014 | 1.3× |
| Stdlib | 108292 | 182.74 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1761 | 10290.61 MB/s | 0 | 0 | 58.2× |
| LightningArena | 1818 | 9968.72 MB/s | 405 | 0 | 56.4× |
| Lightning | 1821 | 9950.35 MB/s | 405 | 0 | 56.3× |
| Easyjson | 3963 | 4572.84 MB/s | 432 | 2 | 25.9× |
| Sonic | 9889 | 1832.81 MB/s | 23043 | 6 | 10.4× |
| SonicFastest | 10007 | 1811.04 MB/s | 23149 | 6 | 10.2× |
| Goccy | 15577 | 1163.52 MB/s | 19459 | 2 | 6.6× |
| LightningDecodeAny | 15908 | 1124.06 MB/s | 29105 | 189 | 6.4× |
| JSONV2 | 46678 | 388.27 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102475 | 176.86 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1979295 | 1014.75 MB/s | 3089565 | 6821 | 9.3× |
| LightningArena | 2078540 | 966.30 MB/s | 3101345 | 6700 | 8.9× |
| Lightning | 2090699 | 960.68 MB/s | 3097858 | 6823 | 8.8× |
| Goccy | 4271248 | 470.24 MB/s | 5411967 | 15835 | 4.3× |
| Sonic | 4436525 | 452.72 MB/s | 10951066 | 13683 | 4.2× |
| SonicFastest | 4479258 | 448.40 MB/s | 10968795 | 13683 | 4.1× |
| Easyjson | 4934061 | 407.07 MB/s | 2981518 | 7439 | 3.7× |
| JSONV2 | 6994140 | 287.17 MB/s | 3173692 | 14563 | 2.6× |
| LightningDecodeAny | 7049546 | 162.04 MB/s | 7378476 | 134004 | 2.6× |
| Stdlib | 18475055 | 108.71 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 823 | 666.98 MB/s | 480 | 1 | 6.8× |
| LightningArena | 825 | 665.53 MB/s | 480 | 1 | 6.8× |
| LightningDestructive | 837 | 656.17 MB/s | 480 | 1 | 6.7× |
| LightningDecodeAny | 1509 | 363.05 MB/s | 1765 | 45 | 3.7× |
| Easyjson | 2138 | 256.80 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2680 | 204.86 MB/s | 1949 | 26 | 2.1× |
| SonicFastest | 2685 | 204.48 MB/s | 1956 | 26 | 2.1× |
| Goccy | 3010 | 182.37 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3306 | 166.06 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5574 | 98.49 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 398821 | 1583.45 MB/s | 400489 | 545 | 13.4× |
| Lightning | 437471 | 1443.56 MB/s | 447123 | 548 | 12.2× |
| LightningArena | 445752 | 1416.74 MB/s | 449185 | 404 | 12.0× |
| SonicFastest | 1011681 | 624.22 MB/s | 1004437 | 1102 | 5.3× |
| Sonic | 1012191 | 623.91 MB/s | 999191 | 1102 | 5.3× |
| Easyjson | 1143927 | 552.06 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1157399 | 545.63 MB/s | 986042 | 1201 | 4.6× |
| JSONV2 | 2147582 | 294.06 MB/s | 571616 | 3144 | 2.5× |
| LightningDecodeAny | 2234786 | 208.93 MB/s | 1992711 | 29072 | 2.4× |
| Stdlib | 5354656 | 117.94 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 574928 | 978.22 MB/s | 391284 | 426 | 9.1× |
| Lightning | 638599 | 880.69 MB/s | 506671 | 433 | 8.2× |
| LightningArena | 645681 | 871.03 MB/s | 508601 | 287 | 8.1× |
| SonicFastest | 1038251 | 541.69 MB/s | 952451 | 1476 | 5.1× |
| Sonic | 1049002 | 536.14 MB/s | 963364 | 1476 | 5.0× |
| Goccy | 1316436 | 427.22 MB/s | 1038677 | 1029 | 4.0× |
| Easyjson | 1750992 | 321.19 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2322756 | 242.13 MB/s | 1992891 | 28581 | 2.3× |
| JSONV2 | 2760867 | 203.71 MB/s | 927451 | 3482 | 1.9× |
| Stdlib | 5244564 | 107.24 MB/s | 1011674 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 546832 | 975.03 MB/s | 333416 | 2084 | 9.9× |
| Lightning | 559062 | 953.70 MB/s | 367496 | 2086 | 9.7× |
| LightningArena | 561573 | 949.44 MB/s | 367544 | 2086 | 9.6× |
| Easyjson | 1099390 | 484.98 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1134691 | 469.89 MB/s | 1033412 | 4351 | 4.8× |
| SonicFastest | 1139637 | 467.85 MB/s | 1033833 | 4351 | 4.7× |
| Goccy | 1307521 | 407.78 MB/s | 1167226 | 5409 | 4.1× |
| JSONV2 | 2515847 | 211.93 MB/s | 745447 | 13288 | 2.1× |
| LightningDecodeAny | 3183259 | 167.49 MB/s | 2659464 | 49349 | 1.7× |
| Stdlib | 5407175 | 98.61 MB/s | 798692 | 17133 | 1.0× |
