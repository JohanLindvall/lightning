# JSON Deserialization Benchmarks

- generated 2026-08-06T12:47:00Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 105699 | 1204.13 MB/s | 49760 | 3 | 10.5× |
| Lightning | 106048 | 1200.16 MB/s | 49760 | 3 | 10.4× |
| LightningDestructive | 106927 | 1190.29 MB/s | 49280 | 2 | 10.4× |
| SonicFastest | 182427 | 697.68 MB/s | 193161 | 10 | 6.1× |
| Sonic | 183221 | 694.65 MB/s | 196211 | 10 | 6.0× |
| Goccy | 196811 | 646.68 MB/s | 224830 | 884 | 5.6× |
| Easyjson | 213566 | 595.95 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 419886 | 303.12 MB/s | 195120 | 1805 | 2.6× |
| LightningDecodeAny | 433451 | 218.37 MB/s | 463409 | 9708 | 2.6× |
| Stdlib | 1107335 | 114.94 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3647225 | 617.20 MB/s | 2532848 | 1143 | 7.2× |
| Lightning | 3681383 | 611.47 MB/s | 2532849 | 1143 | 7.1× |
| LightningArena | 3687331 | 610.48 MB/s | 2532850 | 1143 | 7.1× |
| Sonic | 4747760 | 474.13 MB/s | 15237334 | 970 | 5.5× |
| SonicFastest | 4780133 | 470.92 MB/s | 15233893 | 970 | 5.5× |
| Goccy | 10408963 | 216.26 MB/s | 4123817 | 56532 | 2.5× |
| Easyjson | 11013994 | 204.38 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12625755 | 178.29 MB/s | 19380209 | 223896 | 2.1× |
| JSONV2 | 16365560 | 137.55 MB/s | 3123223 | 3083 | 1.6× |
| Stdlib | 26188229 | 85.96 MB/s | 3123395 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 474187 | 570.25 MB/s | 397296 | 567 | 7.1× |
| Lightning | 475139 | 569.10 MB/s | 397297 | 567 | 7.1× |
| LightningArena | 475723 | 568.40 MB/s | 397297 | 567 | 7.1× |
| SonicFastest | 641860 | 421.28 MB/s | 477063 | 968 | 5.2× |
| Sonic | 646840 | 418.04 MB/s | 487717 | 968 | 5.2× |
| Easyjson | 1404867 | 192.48 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1419773 | 190.46 MB/s | 544165 | 8123 | 2.4× |
| LightningDecodeAny | 1698468 | 159.20 MB/s | 2543877 | 29687 | 2.0× |
| JSONV2 | 2137724 | 126.49 MB/s | 348152 | 1628 | 1.6× |
| Stdlib | 3360826 | 80.46 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1176192 | 1468.47 MB/s | 765560 | 2798 | 11.3× |
| Lightning | 1187786 | 1454.14 MB/s | 765602 | 2799 | 11.2× |
| LightningArena | 1191449 | 1449.67 MB/s | 768416 | 2440 | 11.1× |
| Sonic | 2078487 | 830.99 MB/s | 2692515 | 4020 | 6.4× |
| SonicFastest | 2098659 | 823.00 MB/s | 2713359 | 4020 | 6.3× |
| Goccy | 2428729 | 711.16 MB/s | 2585035 | 14605 | 5.5× |
| LightningDecodeAny | 4183239 | 119.60 MB/s | 4953694 | 76576 | 3.2× |
| Easyjson | 4217387 | 409.54 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4310242 | 400.72 MB/s | 1011635 | 7594 | 3.1× |
| Stdlib | 13266232 | 130.20 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1163 | 1557.97 MB/s | 0 | 0 | 12.1× |
| Lightning | 1167 | 1552.95 MB/s | 0 | 0 | 12.0× |
| LightningDestructive | 1181 | 1534.30 MB/s | 0 | 0 | 11.9× |
| Easyjson | 2538 | 714.08 MB/s | 24 | 1 | 5.5× |
| Goccy | 2893 | 626.29 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6033 | 300.32 MB/s | 3758 | 40 | 2.3× |
| SonicFastest | 6119 | 296.11 MB/s | 3846 | 40 | 2.3× |
| JSONV2 | 7850 | 230.83 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8020 | 225.81 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 14042 | 129.04 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1195 | 1516.16 MB/s | 0 | 0 | 11.8× |
| Lightning | 1196 | 1514.64 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1209 | 1499.09 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2537 | 714.26 MB/s | 24 | 1 | 5.5× |
| Goccy | 2848 | 636.13 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6141 | 295.06 MB/s | 3784 | 40 | 2.3× |
| SonicFastest | 6144 | 294.90 MB/s | 3794 | 40 | 2.3× |
| JSONV2 | 7678 | 236.00 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8089 | 223.89 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 14070 | 128.79 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1352 | 1340.16 MB/s | 144 | 10 | 10.3× |
| LightningArena | 1353 | 1338.98 MB/s | 144 | 10 | 10.3× |
| LightningDestructive | 1398 | 1295.79 MB/s | 144 | 10 | 10.0× |
| Easyjson | 2790 | 649.48 MB/s | 144 | 10 | 5.0× |
| Goccy | 2912 | 622.30 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6181 | 293.15 MB/s | 3819 | 42 | 2.3× |
| SonicFastest | 6204 | 292.09 MB/s | 3806 | 42 | 2.3× |
| LightningDecodeAny | 8022 | 225.74 MB/s | 7552 | 158 | 1.7× |
| JSONV2 | 8089 | 224.00 MB/s | 632 | 7 | 1.7× |
| Stdlib | 13989 | 129.53 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 702 | 704.07 MB/s | 160 | 1 | 7.9× |
| Lightning | 706 | 699.45 MB/s | 160 | 1 | 7.8× |
| SonicFastest | 1240 | 398.41 MB/s | 989 | 6 | 4.4× |
| Sonic | 1241 | 398.17 MB/s | 981 | 6 | 4.4× |
| LightningArena | 1327 | 372.30 MB/s | 4096 | 1 | 4.2× |
| LightningDecodeAny | 1387 | 355.51 MB/s | 1296 | 26 | 4.0× |
| Easyjson | 2238 | 220.77 MB/s | 448 | 3 | 2.5× |
| Goccy | 2412 | 204.84 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3249 | 152.02 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5516 | 89.55 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 433 | 531.14 MB/s | 160 | 1 | 9.5× |
| Lightning | 435 | 528.47 MB/s | 160 | 1 | 9.4× |
| Sonic | 883 | 260.60 MB/s | 656 | 6 | 4.7× |
| SonicFastest | 887 | 259.22 MB/s | 663 | 6 | 4.6× |
| LightningArena | 1100 | 209.08 MB/s | 4096 | 1 | 3.7× |
| LightningDecodeAny | 1136 | 201.57 MB/s | 1296 | 26 | 3.6× |
| Easyjson | 1391 | 165.35 MB/s | 448 | 3 | 3.0× |
| Goccy | 1579 | 145.68 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2450 | 93.87 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4110 | 55.96 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 59279 | 1098.74 MB/s | 97220 | 98 | 9.2× |
| Lightning | 60042 | 1084.77 MB/s | 103440 | 103 | 9.1× |
| LightningArena | 60338 | 1079.46 MB/s | 103440 | 103 | 9.1× |
| Sonic | 99487 | 654.68 MB/s | 155055 | 75 | 5.5× |
| SonicFastest | 100247 | 649.72 MB/s | 156602 | 75 | 5.5× |
| Goccy | 150382 | 433.11 MB/s | 228667 | 134 | 3.6× |
| LightningDecodeAny | 180344 | 295.71 MB/s | 180048 | 3245 | 3.0× |
| JSONV2 | 230898 | 282.08 MB/s | 206654 | 607 | 2.4× |
| Stdlib | 547271 | 119.01 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2583543 | 751.09 MB/s | 2864592 | 1380 | 9.1× |
| LightningArena | 2649880 | 732.29 MB/s | 2864593 | 1380 | 8.9× |
| Lightning | 2670590 | 726.61 MB/s | 2864595 | 1380 | 8.8× |
| SonicFastest | 4762760 | 407.43 MB/s | 14608559 | 1407 | 4.9× |
| Goccy | 4787979 | 405.28 MB/s | 4064958 | 13510 | 4.9× |
| Sonic | 4885245 | 397.21 MB/s | 14608574 | 1407 | 4.8× |
| Easyjson | 7488385 | 259.13 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9291857 | 208.84 MB/s | 7063039 | 218633 | 2.5× |
| JSONV2 | 11245777 | 172.55 MB/s | 3237230 | 13947 | 2.1× |
| Stdlib | 23455625 | 82.73 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1099189 | 3027.53 MB/s | 351704 | 1286 | 19.1× |
| Lightning | 1787876 | 1861.33 MB/s | 2488905 | 2995 | 11.7× |
| LightningArena | 1800454 | 1848.33 MB/s | 2488904 | 2995 | 11.6× |
| Sonic | 2717502 | 1224.59 MB/s | 6483402 | 4248 | 7.7× |
| SonicFastest | 2733128 | 1217.59 MB/s | 6506142 | 4248 | 7.7× |
| LightningDecodeAny | 3634703 | 845.67 MB/s | 4876913 | 56892 | 5.8× |
| Goccy | 4652045 | 715.35 MB/s | 3948910 | 3816 | 4.5× |
| JSONV2 | 7485522 | 444.57 MB/s | 5364521 | 13243 | 2.8× |
| Stdlib | 20961735 | 158.76 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220180 | 1000.75 MB/s | 135872 | 226 | 9.3× |
| LightningArena | 221143 | 996.40 MB/s | 135872 | 226 | 9.2× |
| LightningDestructive | 221432 | 995.10 MB/s | 135872 | 226 | 9.2× |
| SonicFastest | 383292 | 574.88 MB/s | 310402 | 398 | 5.3× |
| Sonic | 384166 | 573.57 MB/s | 314312 | 398 | 5.3× |
| Goccy | 434496 | 507.13 MB/s | 365656 | 1067 | 4.7× |
| Easyjson | 548965 | 401.38 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 723794 | 304.43 MB/s | 129740 | 470 | 2.8× |
| LightningDecodeAny | 868994 | 124.64 MB/s | 897218 | 11703 | 2.3× |
| Stdlib | 2039550 | 108.04 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11477450 | 705.74 MB/s | 11845072 | 20816 | 7.8× |
| Lightning | 11795355 | 686.71 MB/s | 11845073 | 20816 | 7.6× |
| LightningArena | 11800553 | 686.41 MB/s | 11845073 | 20816 | 7.6× |
| SonicFastest | 16860153 | 480.43 MB/s | 70916930 | 40014 | 5.3× |
| Sonic | 17040405 | 475.34 MB/s | 70901838 | 40014 | 5.3× |
| Goccy | 23801180 | 340.32 MB/s | 17334569 | 107150 | 3.8× |
| Easyjson | 30787873 | 263.09 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 36726899 | 141.67 MB/s | 46279353 | 747112 | 2.4× |
| JSONV2 | 43634083 | 185.64 MB/s | 15233716 | 78972 | 2.1× |
| Stdlib | 89581224 | 90.42 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5402548 | 552.23 MB/s | 3764714 | 1504 | 8.7× |
| LightningDestructive | 5765102 | 517.50 MB/s | 3758856 | 29356 | 8.2× |
| Lightning | 5884843 | 506.97 MB/s | 3758856 | 29356 | 8.0× |
| SonicFastest | 8589778 | 347.33 MB/s | 26537411 | 56760 | 5.5× |
| Sonic | 8609968 | 346.51 MB/s | 26611835 | 56760 | 5.5× |
| Easyjson | 16425527 | 181.64 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16493987 | 180.88 MB/s | 10680618 | 273651 | 2.9× |
| LightningDecodeAny | 16534416 | 110.93 MB/s | 23982581 | 351152 | 2.9× |
| JSONV2 | 24396772 | 122.29 MB/s | 9257180 | 86278 | 1.9× |
| Stdlib | 47261333 | 63.13 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1202846 | 601.57 MB/s | 911392 | 30 | 9.6× |
| LightningDestructive | 1248267 | 579.68 MB/s | 907601 | 3618 | 9.3× |
| Lightning | 1258600 | 574.92 MB/s | 907596 | 3618 | 9.2× |
| SonicFastest | 1798138 | 402.41 MB/s | 3196236 | 7226 | 6.4× |
| Sonic | 1800356 | 401.92 MB/s | 3194129 | 7226 | 6.4× |
| Easyjson | 4237307 | 170.77 MB/s | 2847906 | 3698 | 2.7× |
| LightningDecodeAny | 4322622 | 150.50 MB/s | 6500459 | 76546 | 2.7× |
| Goccy | 4798897 | 150.78 MB/s | 2955620 | 80282 | 2.4× |
| JSONV2 | 5561255 | 130.11 MB/s | 2704632 | 7318 | 2.1× |
| Stdlib | 11593676 | 62.41 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1836318 | 858.98 MB/s | 907601 | 3618 | 8.6× |
| LightningArena | 1846584 | 854.20 MB/s | 911393 | 30 | 8.6× |
| Lightning | 1895730 | 832.06 MB/s | 907594 | 3618 | 8.3× |
| SonicFastest | 2263791 | 696.77 MB/s | 5786662 | 7226 | 7.0× |
| Sonic | 2280209 | 691.76 MB/s | 5786680 | 7226 | 6.9× |
| LightningDecodeAny | 3972675 | 189.65 MB/s | 6500459 | 76546 | 4.0× |
| Easyjson | 5622766 | 280.53 MB/s | 2847907 | 3698 | 2.8× |
| Goccy | 5832684 | 270.43 MB/s | 3604128 | 80268 | 2.7× |
| JSONV2 | 6435093 | 245.12 MB/s | 2704594 | 7318 | 2.5× |
| Stdlib | 15797603 | 99.85 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 212463 | 706.59 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 212618 | 706.07 MB/s | 81920 | 1 | 8.6× |
| LightningArena | 212683 | 705.86 MB/s | 81920 | 1 | 8.6× |
| SonicFastest | 271583 | 552.77 MB/s | 245278 | 6 | 6.7× |
| Sonic | 273037 | 549.83 MB/s | 250315 | 6 | 6.7× |
| LightningDecodeAny | 482093 | 311.39 MB/s | 745764 | 10016 | 3.8× |
| Goccy | 878302 | 170.93 MB/s | 328555 | 10005 | 2.1× |
| JSONV2 | 1105963 | 135.74 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1826946 | 82.17 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32466 | 866.04 MB/s | 29216 | 103 | 9.4× |
| LightningArena | 32505 | 865.01 MB/s | 29216 | 103 | 9.3× |
| LightningDestructive | 32803 | 857.15 MB/s | 29088 | 101 | 9.3× |
| Sonic | 64673 | 434.75 MB/s | 48960 | 103 | 4.7× |
| SonicFastest | 64767 | 434.12 MB/s | 48892 | 103 | 4.7× |
| Easyjson | 69496 | 404.59 MB/s | 32304 | 138 | 4.4× |
| Goccy | 73037 | 384.97 MB/s | 59240 | 188 | 4.2× |
| JSONV2 | 133924 | 209.95 MB/s | 36897 | 242 | 2.3× |
| LightningDecodeAny | 152700 | 184.13 MB/s | 140576 | 2643 | 2.0× |
| Stdlib | 303569 | 92.62 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1957 | 1189.36 MB/s | 32 | 1 | 11.6× |
| LightningArena | 1966 | 1184.14 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2038 | 1142.29 MB/s | 32 | 1 | 11.2× |
| Goccy | 4223 | 551.21 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4264 | 546.00 MB/s | 192 | 2 | 5.3× |
| Sonic | 5131 | 453.69 MB/s | 4352 | 6 | 4.4× |
| SonicFastest | 5140 | 452.88 MB/s | 4358 | 6 | 4.4× |
| JSONV2 | 8423 | 276.40 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9862 | 170.86 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22738 | 102.38 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 221 | 854.00 MB/s | 0 | 0 | 11.0× |
| Lightning | 222 | 853.44 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 224 | 845.00 MB/s | 0 | 0 | 10.9× |
| Goccy | 390 | 484.93 MB/s | 304 | 2 | 6.2× |
| Easyjson | 489 | 386.79 MB/s | 0 | 0 | 5.0× |
| Sonic | 798 | 236.78 MB/s | 500 | 4 | 3.1× |
| SonicFastest | 798 | 236.71 MB/s | 502 | 4 | 3.0× |
| JSONV2 | 1033 | 183.00 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1186 | 112.98 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2435 | 77.63 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1517 | 1444.18 MB/s | 0 | 0 | 10.6× |
| LightningArena | 1518 | 1443.27 MB/s | 0 | 0 | 10.5× |
| LightningDestructive | 1539 | 1423.69 MB/s | 0 | 0 | 10.4× |
| Easyjson | 3205 | 683.63 MB/s | 24 | 1 | 5.0× |
| Goccy | 3253 | 673.50 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6603 | 331.80 MB/s | 4100 | 40 | 2.4× |
| Sonic | 6610 | 331.48 MB/s | 4094 | 40 | 2.4× |
| JSONV2 | 7957 | 275.35 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8333 | 217.33 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 16014 | 136.82 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 657045 | 776.93 MB/s | 457537 | 1009 | 9.2× |
| LightningDestructive | 657859 | 775.97 MB/s | 457537 | 1009 | 9.2× |
| LightningArena | 658996 | 774.63 MB/s | 457536 | 1009 | 9.1× |
| Goccy | 1162593 | 439.08 MB/s | 1140387 | 5006 | 5.2× |
| Sonic | 1169913 | 436.34 MB/s | 899372 | 2006 | 5.1× |
| SonicFastest | 1173061 | 435.17 MB/s | 904829 | 2006 | 5.1× |
| Easyjson | 1534700 | 332.62 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3186734 | 160.19 MB/s | 1076021 | 12646 | 1.9× |
| LightningDecodeAny | 3347971 | 137.83 MB/s | 2950648 | 64018 | 1.8× |
| Stdlib | 6024885 | 84.73 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1336 | 14807.30 MB/s | 0 | 0 | 83.5× |
| Lightning | 1337 | 14805.74 MB/s | 0 | 0 | 83.5× |
| LightningDestructive | 1360 | 14551.33 MB/s | 0 | 0 | 82.0× |
| Goccy | 20028 | 988.04 MB/s | 20491 | 2 | 5.6× |
| SonicFastest | 27956 | 707.87 MB/s | 22074 | 4 | 4.0× |
| Sonic | 28053 | 705.41 MB/s | 22258 | 4 | 4.0× |
| JSONV2 | 29566 | 669.32 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 74828 | 264.45 MB/s | 116864 | 2015 | 1.5× |
| Easyjson | 86034 | 230.01 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111577 | 177.36 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2618 | 6923.99 MB/s | 0 | 0 | 39.2× |
| Lightning | 2752 | 6586.18 MB/s | 432 | 2 | 37.3× |
| LightningArena | 2753 | 6583.37 MB/s | 432 | 2 | 37.3× |
| Easyjson | 3938 | 4602.10 MB/s | 432 | 2 | 26.1× |
| SonicFastest | 10143 | 1786.85 MB/s | 22803 | 6 | 10.1× |
| Sonic | 10208 | 1775.46 MB/s | 22973 | 6 | 10.1× |
| Goccy | 16260 | 1114.65 MB/s | 19459 | 2 | 6.3× |
| LightningDecodeAny | 16850 | 1061.22 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 45152 | 401.40 MB/s | 16498 | 50 | 2.3× |
| Stdlib | 102626 | 176.60 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2447489 | 820.63 MB/s | 3089564 | 6821 | 7.6× |
| Lightning | 2511602 | 799.69 MB/s | 3091277 | 6827 | 7.5× |
| LightningArena | 2531900 | 793.28 MB/s | 3094371 | 6703 | 7.4× |
| Goccy | 4311291 | 465.87 MB/s | 5412332 | 15830 | 4.3× |
| Sonic | 4521465 | 444.21 MB/s | 10872990 | 13683 | 4.1× |
| SonicFastest | 4529696 | 443.41 MB/s | 10895428 | 13683 | 4.1× |
| Easyjson | 4960105 | 404.93 MB/s | 2981486 | 7439 | 3.8× |
| JSONV2 | 7003079 | 286.80 MB/s | 3173685 | 14563 | 2.7× |
| LightningDecodeAny | 7035954 | 162.35 MB/s | 8503513 | 134008 | 2.7× |
| Stdlib | 18721501 | 107.28 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 886 | 619.97 MB/s | 480 | 1 | 6.4× |
| LightningArena | 890 | 616.87 MB/s | 480 | 1 | 6.4× |
| LightningDestructive | 895 | 613.27 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1674 | 327.35 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 2186 | 251.19 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2684 | 204.54 MB/s | 1947 | 26 | 2.1× |
| SonicFastest | 2692 | 203.97 MB/s | 1950 | 26 | 2.1× |
| Goccy | 3002 | 182.86 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3293 | 166.74 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5675 | 96.73 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 499684 | 1263.83 MB/s | 402728 | 545 | 10.8× |
| Lightning | 566719 | 1114.33 MB/s | 451257 | 857 | 9.5× |
| LightningArena | 570439 | 1107.07 MB/s | 453017 | 712 | 9.4× |
| SonicFastest | 1025883 | 615.58 MB/s | 1006326 | 1102 | 5.3× |
| Sonic | 1028349 | 614.10 MB/s | 1003598 | 1102 | 5.2× |
| Easyjson | 1152933 | 547.75 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1158028 | 545.34 MB/s | 985343 | 1201 | 4.7× |
| JSONV2 | 2149505 | 293.80 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2393214 | 195.10 MB/s | 2076506 | 30126 | 2.3× |
| Stdlib | 5389579 | 117.17 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 616587 | 912.13 MB/s | 546571 | 429 | 8.6× |
| Lightning | 788770 | 713.02 MB/s | 769938 | 1235 | 6.7× |
| LightningArena | 789962 | 711.94 MB/s | 771666 | 1088 | 6.7× |
| SonicFastest | 1016104 | 553.49 MB/s | 911737 | 1476 | 5.2× |
| Sonic | 1026142 | 548.08 MB/s | 925589 | 1476 | 5.1× |
| Goccy | 1328181 | 423.44 MB/s | 1037348 | 1030 | 4.0× |
| Easyjson | 1753694 | 320.70 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2638927 | 213.12 MB/s | 2180440 | 30126 | 2.0× |
| JSONV2 | 2758229 | 203.90 MB/s | 927451 | 3482 | 1.9× |
| Stdlib | 5282026 | 106.48 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 661316 | 806.24 MB/s | 333416 | 2084 | 8.3× |
| Lightning | 683486 | 780.09 MB/s | 368224 | 2293 | 8.0× |
| LightningArena | 688754 | 774.12 MB/s | 368224 | 2293 | 8.0× |
| Easyjson | 1122267 | 475.09 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1150885 | 463.28 MB/s | 1036770 | 4351 | 4.8× |
| Sonic | 1155779 | 461.32 MB/s | 1038871 | 4351 | 4.7× |
| Goccy | 1318837 | 404.28 MB/s | 1167251 | 5409 | 4.2× |
| JSONV2 | 2528851 | 210.84 MB/s | 745453 | 13288 | 2.2× |
| LightningDecodeAny | 3384915 | 157.52 MB/s | 2992878 | 50076 | 1.6× |
| Stdlib | 5479303 | 97.31 MB/s | 798692 | 17133 | 1.0× |
