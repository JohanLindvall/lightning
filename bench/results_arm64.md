# JSON Deserialization Benchmarks

- generated 2026-07-27T11:19:54Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 105449 | 1206.99 MB/s | 49760 | 3 | 10.5× |
| Lightning | 105790 | 1203.09 MB/s | 49760 | 3 | 10.5× |
| LightningDestructive | 106644 | 1193.46 MB/s | 49280 | 2 | 10.4× |
| Sonic | 190141 | 669.37 MB/s | 204307 | 10 | 5.8× |
| SonicFastest | 190911 | 666.67 MB/s | 206944 | 10 | 5.8× |
| Goccy | 198303 | 641.82 MB/s | 225201 | 884 | 5.6× |
| Easyjson | 213462 | 596.24 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 421785 | 301.75 MB/s | 195122 | 1805 | 2.6× |
| LightningDecodeAny | 456219 | 207.47 MB/s | 463410 | 9708 | 2.4× |
| Stdlib | 1105999 | 115.08 MB/s | 199697 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3659136 | 615.19 MB/s | 2532848 | 1143 | 7.2× |
| Lightning | 3683054 | 611.19 MB/s | 2532851 | 1143 | 7.1× |
| LightningArena | 3690613 | 609.94 MB/s | 2532849 | 1143 | 7.1× |
| Sonic | 4675444 | 481.46 MB/s | 15244731 | 970 | 5.6× |
| SonicFastest | 4865312 | 462.67 MB/s | 15233811 | 970 | 5.4× |
| Goccy | 10244995 | 219.72 MB/s | 4111358 | 56532 | 2.6× |
| Easyjson | 11107536 | 202.66 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12951470 | 173.80 MB/s | 19380210 | 223896 | 2.0× |
| JSONV2 | 16139890 | 139.47 MB/s | 3123205 | 3083 | 1.6× |
| Stdlib | 26168503 | 86.02 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 478453 | 565.16 MB/s | 397296 | 567 | 7.0× |
| LightningDestructive | 480070 | 563.26 MB/s | 397296 | 567 | 7.0× |
| Lightning | 480987 | 562.18 MB/s | 397296 | 567 | 7.0× |
| Sonic | 633329 | 426.95 MB/s | 467907 | 968 | 5.3× |
| SonicFastest | 641120 | 421.77 MB/s | 487639 | 968 | 5.2× |
| Easyjson | 1403424 | 192.67 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1414881 | 191.11 MB/s | 542869 | 8122 | 2.4× |
| LightningDecodeAny | 1758736 | 153.75 MB/s | 2543881 | 29687 | 1.9× |
| JSONV2 | 2111945 | 128.04 MB/s | 348154 | 1628 | 1.6× |
| Stdlib | 3357424 | 80.54 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1169032 | 1477.47 MB/s | 765560 | 2798 | 11.3× |
| LightningArena | 1177656 | 1466.65 MB/s | 768416 | 2440 | 11.3× |
| Lightning | 1179188 | 1464.74 MB/s | 765602 | 2799 | 11.2× |
| SonicFastest | 2081209 | 829.90 MB/s | 2718801 | 4020 | 6.4× |
| Sonic | 2090644 | 826.16 MB/s | 2734967 | 4020 | 6.3× |
| Goccy | 2357129 | 732.76 MB/s | 2581789 | 14604 | 5.6× |
| Easyjson | 4223636 | 408.94 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4279729 | 403.58 MB/s | 1011631 | 7594 | 3.1× |
| LightningDecodeAny | 4294509 | 116.50 MB/s | 4953694 | 76576 | 3.1× |
| Stdlib | 13250641 | 130.35 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1163 | 1558.38 MB/s | 0 | 0 | 12.1× |
| Lightning | 1165 | 1555.26 MB/s | 0 | 0 | 12.1× |
| LightningDestructive | 1184 | 1529.93 MB/s | 0 | 0 | 11.9× |
| Easyjson | 2531 | 715.98 MB/s | 24 | 1 | 5.6× |
| Goccy | 2844 | 637.17 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6088 | 297.62 MB/s | 3841 | 40 | 2.3× |
| Sonic | 6101 | 296.98 MB/s | 3846 | 40 | 2.3× |
| JSONV2 | 7728 | 234.48 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8327 | 217.50 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 14063 | 128.85 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1193 | 1519.18 MB/s | 0 | 0 | 11.8× |
| Lightning | 1194 | 1517.60 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1210 | 1497.59 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2519 | 719.26 MB/s | 24 | 1 | 5.6× |
| Goccy | 2792 | 648.91 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6081 | 297.95 MB/s | 3779 | 40 | 2.3× |
| Sonic | 6087 | 297.67 MB/s | 3815 | 40 | 2.3× |
| JSONV2 | 7727 | 234.50 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8003 | 226.28 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 14085 | 128.65 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1346 | 1345.83 MB/s | 144 | 10 | 10.4× |
| LightningArena | 1348 | 1344.13 MB/s | 144 | 10 | 10.4× |
| LightningDestructive | 1399 | 1295.14 MB/s | 144 | 10 | 10.1× |
| Easyjson | 2779 | 651.95 MB/s | 144 | 10 | 5.1× |
| Goccy | 2875 | 630.27 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6135 | 295.37 MB/s | 3762 | 42 | 2.3× |
| Sonic | 6159 | 294.20 MB/s | 3767 | 42 | 2.3× |
| JSONV2 | 7965 | 227.48 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 7988 | 226.71 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 14063 | 128.85 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 701 | 704.40 MB/s | 160 | 1 | 7.9× |
| Lightning | 707 | 699.02 MB/s | 160 | 1 | 7.8× |
| Sonic | 1244 | 397.17 MB/s | 979 | 6 | 4.4× |
| SonicFastest | 1251 | 394.91 MB/s | 989 | 6 | 4.4× |
| LightningArena | 1346 | 367.13 MB/s | 4096 | 1 | 4.1× |
| LightningDecodeAny | 1395 | 353.51 MB/s | 1296 | 26 | 4.0× |
| Easyjson | 2229 | 221.66 MB/s | 448 | 3 | 2.5× |
| Goccy | 2420 | 204.14 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3240 | 152.48 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5526 | 89.39 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 434 | 530.08 MB/s | 160 | 1 | 9.5× |
| Lightning | 440 | 523.29 MB/s | 160 | 1 | 9.4× |
| Sonic | 882 | 260.72 MB/s | 651 | 6 | 4.7× |
| SonicFastest | 887 | 259.19 MB/s | 657 | 6 | 4.6× |
| LightningArena | 1080 | 213.00 MB/s | 4096 | 1 | 3.8× |
| LightningDecodeAny | 1149 | 199.33 MB/s | 1296 | 26 | 3.6× |
| Easyjson | 1387 | 165.80 MB/s | 448 | 3 | 3.0× |
| Goccy | 1581 | 145.47 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2447 | 93.98 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4120 | 55.82 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 58088 | 1121.26 MB/s | 97220 | 98 | 9.4× |
| LightningArena | 59888 | 1087.56 MB/s | 103440 | 103 | 9.1× |
| Lightning | 60117 | 1083.43 MB/s | 103440 | 103 | 9.1× |
| Sonic | 100446 | 648.43 MB/s | 156625 | 75 | 5.4× |
| SonicFastest | 102423 | 635.91 MB/s | 157805 | 75 | 5.3× |
| Goccy | 147476 | 441.64 MB/s | 229452 | 134 | 3.7× |
| LightningDecodeAny | 181005 | 294.63 MB/s | 180048 | 3245 | 3.0× |
| JSONV2 | 226351 | 287.75 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 546944 | 119.08 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2619430 | 740.80 MB/s | 2864593 | 1380 | 9.0× |
| LightningArena | 2664918 | 728.15 MB/s | 2864594 | 1380 | 8.8× |
| Lightning | 2672463 | 726.10 MB/s | 2864595 | 1380 | 8.8× |
| Sonic | 4778807 | 406.06 MB/s | 14608641 | 1407 | 4.9× |
| SonicFastest | 4820262 | 402.57 MB/s | 14608714 | 1407 | 4.9× |
| Goccy | 4911695 | 395.07 MB/s | 4065604 | 13510 | 4.8× |
| Easyjson | 7467307 | 259.86 MB/s | 3871266 | 15043 | 3.2× |
| LightningDecodeAny | 9158249 | 211.88 MB/s | 7063039 | 218633 | 2.6× |
| JSONV2 | 11254403 | 172.42 MB/s | 3237221 | 13947 | 2.1× |
| Stdlib | 23554489 | 82.38 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1099815 | 3025.81 MB/s | 351704 | 1286 | 19.1× |
| Lightning | 1828543 | 1819.94 MB/s | 2488905 | 2995 | 11.5× |
| LightningArena | 1846250 | 1802.48 MB/s | 2488904 | 2995 | 11.4× |
| Sonic | 2829977 | 1175.92 MB/s | 6419060 | 4248 | 7.4× |
| SonicFastest | 2851669 | 1166.98 MB/s | 6431392 | 4248 | 7.4× |
| LightningDecodeAny | 3550593 | 865.70 MB/s | 4876912 | 56892 | 5.9× |
| Goccy | 5005597 | 664.82 MB/s | 3948909 | 3816 | 4.2× |
| JSONV2 | 7580422 | 439.00 MB/s | 5364515 | 13243 | 2.8× |
| Stdlib | 20992877 | 158.52 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220404 | 999.74 MB/s | 135872 | 226 | 9.3× |
| LightningArena | 221192 | 996.18 MB/s | 135872 | 226 | 9.2× |
| LightningDestructive | 223408 | 986.29 MB/s | 135872 | 226 | 9.1× |
| SonicFastest | 381456 | 577.65 MB/s | 305883 | 398 | 5.4× |
| Sonic | 393531 | 559.92 MB/s | 336932 | 398 | 5.2× |
| Goccy | 435314 | 506.18 MB/s | 364261 | 1067 | 4.7× |
| Easyjson | 551396 | 399.61 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 736457 | 299.20 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 857532 | 126.31 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2042416 | 107.88 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11463870 | 706.57 MB/s | 11845073 | 20816 | 7.8× |
| Lightning | 11781685 | 687.51 MB/s | 11845078 | 20816 | 7.6× |
| LightningArena | 11792097 | 686.90 MB/s | 11845072 | 20816 | 7.6× |
| Sonic | 16959206 | 477.62 MB/s | 70887395 | 40014 | 5.3× |
| SonicFastest | 17091967 | 473.91 MB/s | 70887537 | 40014 | 5.2× |
| Goccy | 23329041 | 347.21 MB/s | 17338000 | 107149 | 3.8× |
| Easyjson | 30970210 | 261.54 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 37304167 | 139.48 MB/s | 46279352 | 747112 | 2.4× |
| JSONV2 | 43768365 | 185.07 MB/s | 15233783 | 78972 | 2.0× |
| Stdlib | 89455926 | 90.55 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5410499 | 551.42 MB/s | 3764712 | 1504 | 8.7× |
| LightningDestructive | 5758844 | 518.07 MB/s | 3758856 | 29356 | 8.2× |
| Lightning | 5843168 | 510.59 MB/s | 3758859 | 29356 | 8.1× |
| Sonic | 8671648 | 344.05 MB/s | 26553454 | 56760 | 5.4× |
| SonicFastest | 8752219 | 340.88 MB/s | 26548247 | 56760 | 5.4× |
| Easyjson | 16343882 | 182.54 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16503450 | 180.78 MB/s | 10563242 | 273647 | 2.9× |
| LightningDecodeAny | 16719375 | 109.70 MB/s | 23982580 | 351152 | 2.8× |
| JSONV2 | 24367246 | 122.44 MB/s | 9257182 | 86278 | 1.9× |
| Stdlib | 47201690 | 63.21 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1222531 | 591.88 MB/s | 911392 | 30 | 9.5× |
| LightningDestructive | 1254320 | 576.88 MB/s | 907600 | 3618 | 9.2× |
| Lightning | 1280803 | 564.96 MB/s | 907596 | 3618 | 9.0× |
| Sonic | 1783099 | 405.81 MB/s | 3188635 | 7226 | 6.5× |
| SonicFastest | 1784259 | 405.54 MB/s | 3190869 | 7226 | 6.5× |
| Easyjson | 4208394 | 171.94 MB/s | 2847905 | 3698 | 2.8× |
| LightningDecodeAny | 4339392 | 149.92 MB/s | 6500458 | 76546 | 2.7× |
| Goccy | 4716944 | 153.40 MB/s | 2800645 | 80273 | 2.5× |
| JSONV2 | 5689462 | 127.18 MB/s | 2704614 | 7318 | 2.0× |
| Stdlib | 11587292 | 62.45 MB/s | 2704553 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1845784 | 854.57 MB/s | 907601 | 3618 | 8.5× |
| LightningArena | 1849200 | 852.99 MB/s | 911393 | 30 | 8.5× |
| Lightning | 1894974 | 832.39 MB/s | 907594 | 3618 | 8.3× |
| SonicFastest | 2259596 | 698.07 MB/s | 5783116 | 7226 | 7.0× |
| Sonic | 2277412 | 692.61 MB/s | 5783836 | 7226 | 6.9× |
| LightningDecodeAny | 3978024 | 189.39 MB/s | 6500458 | 76546 | 4.0× |
| Easyjson | 5607598 | 281.29 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5710762 | 276.21 MB/s | 3572621 | 80266 | 2.8× |
| JSONV2 | 6459782 | 244.18 MB/s | 2704591 | 7318 | 2.4× |
| Stdlib | 15731540 | 100.27 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 212368 | 706.90 MB/s | 81920 | 1 | 8.6× |
| LightningArena | 212427 | 706.71 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 212814 | 705.42 MB/s | 81920 | 1 | 8.6× |
| SonicFastest | 275137 | 545.63 MB/s | 259620 | 6 | 6.7× |
| Sonic | 275610 | 544.70 MB/s | 260486 | 6 | 6.6× |
| LightningDecodeAny | 487231 | 308.11 MB/s | 745764 | 10016 | 3.8× |
| Goccy | 862214 | 174.11 MB/s | 324903 | 10004 | 2.1× |
| JSONV2 | 1103318 | 136.07 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1832102 | 81.94 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 31993 | 878.85 MB/s | 29216 | 103 | 9.5× |
| LightningDestructive | 32068 | 876.79 MB/s | 29088 | 101 | 9.5× |
| Lightning | 32094 | 876.10 MB/s | 29216 | 103 | 9.5× |
| Sonic | 64010 | 439.26 MB/s | 46986 | 103 | 4.7× |
| SonicFastest | 64489 | 436.00 MB/s | 47662 | 103 | 4.7× |
| Easyjson | 69272 | 405.90 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72216 | 389.35 MB/s | 59213 | 188 | 4.2× |
| JSONV2 | 134560 | 208.95 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 148046 | 189.92 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 303866 | 92.53 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1961 | 1187.29 MB/s | 32 | 1 | 11.6× |
| LightningArena | 1968 | 1182.66 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2027 | 1148.38 MB/s | 32 | 1 | 11.3× |
| Goccy | 4144 | 561.82 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4229 | 550.44 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5137 | 453.14 MB/s | 4286 | 6 | 4.4× |
| Sonic | 5146 | 452.35 MB/s | 4321 | 6 | 4.4× |
| JSONV2 | 8456 | 275.29 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10002 | 168.47 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22819 | 102.02 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221 | 854.11 MB/s | 0 | 0 | 11.1× |
| LightningArena | 222 | 853.35 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 223 | 846.27 MB/s | 0 | 0 | 11.0× |
| Goccy | 394 | 480.09 MB/s | 304 | 2 | 6.2× |
| Easyjson | 494 | 382.98 MB/s | 0 | 0 | 5.0× |
| Sonic | 797 | 237.07 MB/s | 499 | 4 | 3.1× |
| SonicFastest | 798 | 236.79 MB/s | 506 | 4 | 3.1× |
| JSONV2 | 1039 | 181.87 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1183 | 113.28 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2446 | 77.27 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1517 | 1444.59 MB/s | 0 | 0 | 10.5× |
| LightningArena | 1520 | 1441.22 MB/s | 0 | 0 | 10.5× |
| LightningDestructive | 1539 | 1423.28 MB/s | 0 | 0 | 10.4× |
| Goccy | 3148 | 695.90 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3191 | 686.53 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6454 | 339.46 MB/s | 3973 | 40 | 2.5× |
| Sonic | 6466 | 338.86 MB/s | 3965 | 40 | 2.5× |
| LightningDecodeAny | 7952 | 227.74 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 8033 | 272.74 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15984 | 137.08 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 656858 | 777.15 MB/s | 457536 | 1009 | 9.2× |
| LightningDestructive | 657154 | 776.80 MB/s | 457536 | 1009 | 9.2× |
| LightningArena | 661148 | 772.10 MB/s | 457536 | 1009 | 9.2× |
| Goccy | 1167776 | 437.14 MB/s | 1137942 | 5006 | 5.2× |
| Sonic | 1171862 | 435.61 MB/s | 911963 | 2006 | 5.2× |
| SonicFastest | 1172340 | 435.43 MB/s | 903360 | 2006 | 5.2× |
| Easyjson | 1555068 | 328.27 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3219386 | 158.56 MB/s | 1076009 | 12646 | 1.9× |
| LightningDecodeAny | 3377380 | 136.63 MB/s | 2950650 | 64018 | 1.8× |
| Stdlib | 6073285 | 84.05 MB/s | 1162119 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1338 | 14792.57 MB/s | 0 | 0 | 86.8× |
| LightningArena | 1339 | 14773.89 MB/s | 0 | 0 | 86.7× |
| LightningDestructive | 1367 | 14478.51 MB/s | 0 | 0 | 85.0× |
| Goccy | 19960 | 991.41 MB/s | 20491 | 2 | 5.8× |
| Sonic | 28225 | 701.13 MB/s | 22284 | 4 | 4.1× |
| SonicFastest | 28278 | 699.81 MB/s | 22430 | 4 | 4.1× |
| JSONV2 | 29545 | 669.80 MB/s | 8 | 1 | 3.9× |
| LightningDecodeAny | 77486 | 255.37 MB/s | 116864 | 2015 | 1.5× |
| Easyjson | 91075 | 217.28 MB/s | 0 | 0 | 1.3× |
| Stdlib | 116146 | 170.38 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2607 | 6952.34 MB/s | 0 | 0 | 39.5× |
| LightningArena | 2760 | 6567.06 MB/s | 432 | 2 | 37.3× |
| Lightning | 2765 | 6554.47 MB/s | 432 | 2 | 37.2× |
| Easyjson | 3947 | 4592.04 MB/s | 432 | 2 | 26.1× |
| Sonic | 10232 | 1771.24 MB/s | 24350 | 6 | 10.1× |
| SonicFastest | 10240 | 1769.86 MB/s | 24228 | 6 | 10.0× |
| Goccy | 16118 | 1124.48 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 17546 | 1019.12 MB/s | 29088 | 191 | 5.9× |
| JSONV2 | 45518 | 398.17 MB/s | 16500 | 50 | 2.3× |
| Stdlib | 102885 | 176.16 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2403633 | 835.61 MB/s | 3089564 | 6821 | 7.8× |
| Lightning | 2450788 | 819.53 MB/s | 3091277 | 6827 | 7.7× |
| LightningArena | 2480234 | 809.80 MB/s | 3094371 | 6703 | 7.6× |
| Goccy | 4305547 | 466.49 MB/s | 5412255 | 15833 | 4.4× |
| SonicFastest | 4464653 | 449.87 MB/s | 10943751 | 13683 | 4.2× |
| Sonic | 4475244 | 448.80 MB/s | 10924793 | 13683 | 4.2× |
| Easyjson | 4964343 | 404.58 MB/s | 2981486 | 7439 | 3.8× |
| JSONV2 | 7004446 | 286.75 MB/s | 3173686 | 14563 | 2.7× |
| LightningDecodeAny | 7426066 | 153.82 MB/s | 8503513 | 134008 | 2.5× |
| Stdlib | 18766026 | 107.03 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 908 | 604.95 MB/s | 480 | 1 | 6.3× |
| LightningArena | 913 | 601.32 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 918 | 598.04 MB/s | 480 | 1 | 6.2× |
| LightningDecodeAny | 1786 | 306.82 MB/s | 2021 | 46 | 3.2× |
| Easyjson | 2247 | 244.33 MB/s | 1616 | 5 | 2.5× |
| Sonic | 2751 | 199.58 MB/s | 2037 | 26 | 2.1× |
| SonicFastest | 2757 | 199.13 MB/s | 2024 | 26 | 2.1× |
| Goccy | 3087 | 177.83 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3359 | 163.42 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5708 | 96.18 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 502000 | 1258.00 MB/s | 402728 | 545 | 10.8× |
| Lightning | 567020 | 1113.74 MB/s | 451257 | 857 | 9.5× |
| LightningArena | 567020 | 1113.74 MB/s | 453017 | 712 | 9.5× |
| SonicFastest | 1025592 | 615.76 MB/s | 1010313 | 1102 | 5.3× |
| Sonic | 1025595 | 615.75 MB/s | 1008423 | 1102 | 5.3× |
| Easyjson | 1158983 | 544.89 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1178291 | 535.96 MB/s | 988210 | 1202 | 4.6× |
| JSONV2 | 2169023 | 291.15 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2360205 | 197.82 MB/s | 2076505 | 30126 | 2.3× |
| Stdlib | 5409705 | 116.74 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 608548 | 924.18 MB/s | 546572 | 429 | 8.7× |
| LightningArena | 768663 | 731.67 MB/s | 771666 | 1088 | 6.9× |
| Lightning | 771814 | 728.68 MB/s | 769938 | 1235 | 6.9× |
| SonicFastest | 1049193 | 536.04 MB/s | 955807 | 1476 | 5.0× |
| Sonic | 1053096 | 534.05 MB/s | 969028 | 1476 | 5.0× |
| Goccy | 1339135 | 419.98 MB/s | 1042007 | 1030 | 4.0× |
| Easyjson | 1754997 | 320.46 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2635861 | 213.37 MB/s | 2180439 | 30126 | 2.0× |
| JSONV2 | 2773062 | 202.81 MB/s | 927440 | 3482 | 1.9× |
| Stdlib | 5292368 | 106.27 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 661125 | 806.47 MB/s | 333416 | 2084 | 8.3× |
| Lightning | 679595 | 784.55 MB/s | 368224 | 2293 | 8.1× |
| LightningArena | 685746 | 777.52 MB/s | 368224 | 2293 | 8.0× |
| Easyjson | 1117825 | 476.98 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1146844 | 464.91 MB/s | 1038660 | 4351 | 4.8× |
| Sonic | 1149782 | 463.72 MB/s | 1044955 | 4351 | 4.8× |
| Goccy | 1305742 | 408.33 MB/s | 1167216 | 5409 | 4.2× |
| JSONV2 | 2540109 | 209.90 MB/s | 745447 | 13288 | 2.2× |
| LightningDecodeAny | 3362282 | 158.58 MB/s | 2992879 | 50076 | 1.6× |
| Stdlib | 5496546 | 97.00 MB/s | 798692 | 17133 | 1.0× |
