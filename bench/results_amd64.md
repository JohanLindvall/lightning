# JSON Deserialization Benchmarks

- generated 2026-09-03T06:45:15Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 77134 | 1650.05 MB/s | 49760 | 3 | 16.9× |
| Lightning | 77244 | 1647.70 MB/s | 49760 | 3 | 16.9× |
| LightningDestructive | 79262 | 1605.76 MB/s | 49280 | 2 | 16.5× |
| SonicFastest | 192361 | 661.65 MB/s | 213845 | 15 | 6.8× |
| Sonic | 193367 | 658.20 MB/s | 214048 | 15 | 6.8× |
| Easyjson | 231873 | 548.90 MB/s | 122864 | 14 | 5.6× |
| Goccy | 248660 | 511.84 MB/s | 225140 | 884 | 5.3× |
| JSONV2 | 415597 | 306.25 MB/s | 195128 | 1805 | 3.1× |
| LightningDecodeAny | 423576 | 223.46 MB/s | 463410 | 9708 | 3.1× |
| Stdlib | 1306574 | 97.41 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3331755 | 675.64 MB/s | 2532848 | 1143 | 9.2× |
| Lightning | 3358076 | 670.34 MB/s | 2532849 | 1143 | 9.1× |
| LightningArena | 3365204 | 668.92 MB/s | 2532849 | 1143 | 9.1× |
| SonicFastest | 5541188 | 406.24 MB/s | 4867223 | 2584 | 5.5× |
| Sonic | 5661301 | 397.62 MB/s | 4865662 | 2584 | 5.4× |
| LightningDecodeAny | 12800843 | 175.85 MB/s | 19380209 | 223896 | 2.4× |
| Goccy | 13082968 | 172.06 MB/s | 4168734 | 56534 | 2.3× |
| Easyjson | 13717576 | 164.10 MB/s | 3099809 | 2120 | 2.2× |
| JSONV2 | 16930032 | 132.96 MB/s | 3123214 | 3083 | 1.8× |
| Stdlib | 30681118 | 73.37 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 460948 | 586.62 MB/s | 397296 | 567 | 8.7× |
| Lightning | 463314 | 583.63 MB/s | 397296 | 567 | 8.7× |
| LightningDestructive | 480371 | 562.90 MB/s | 397297 | 567 | 8.4× |
| SonicFastest | 753696 | 358.77 MB/s | 641960 | 1147 | 5.3× |
| Sonic | 755733 | 357.80 MB/s | 642007 | 1147 | 5.3× |
| Easyjson | 1749075 | 154.60 MB/s | 330272 | 749 | 2.3× |
| Goccy | 1769313 | 152.83 MB/s | 541816 | 8122 | 2.3× |
| LightningDecodeAny | 2164950 | 124.90 MB/s | 2543876 | 29687 | 1.9× |
| JSONV2 | 2242799 | 120.56 MB/s | 348159 | 1628 | 1.8× |
| Stdlib | 4025655 | 67.17 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 959456 | 1800.19 MB/s | 765560 | 2798 | 17.5× |
| LightningArena | 987583 | 1748.92 MB/s | 768416 | 2440 | 17.0× |
| Lightning | 990845 | 1743.16 MB/s | 765601 | 2799 | 17.0× |
| SonicFastest | 2168596 | 796.46 MB/s | 2693265 | 5547 | 7.8× |
| Sonic | 2180202 | 792.22 MB/s | 2693282 | 5547 | 7.7× |
| Goccy | 2511562 | 687.70 MB/s | 2581068 | 14603 | 6.7× |
| LightningDecodeAny | 3883510 | 128.83 MB/s | 4953692 | 76576 | 4.3× |
| Easyjson | 4007081 | 431.04 MB/s | 972032 | 5389 | 4.2× |
| JSONV2 | 4350553 | 397.01 MB/s | 1011616 | 7594 | 3.9× |
| Stdlib | 16809695 | 102.75 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 754 | 2404.15 MB/s | 0 | 0 | 21.4× |
| LightningArena | 781 | 2320.89 MB/s | 0 | 0 | 20.6× |
| LightningDestructive | 788 | 2298.87 MB/s | 0 | 0 | 20.4× |
| Easyjson | 2861 | 633.37 MB/s | 24 | 1 | 5.6× |
| Goccy | 3428 | 528.66 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6305 | 287.38 MB/s | 3352 | 38 | 2.6× |
| Sonic | 6375 | 284.23 MB/s | 3349 | 38 | 2.5× |
| JSONV2 | 7957 | 227.71 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8454 | 214.21 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 16093 | 112.60 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 802 | 2260.22 MB/s | 0 | 0 | 19.8× |
| LightningArena | 804 | 2255.15 MB/s | 0 | 0 | 19.7× |
| LightningDestructive | 830 | 2182.79 MB/s | 0 | 0 | 19.1× |
| Easyjson | 2840 | 638.09 MB/s | 24 | 1 | 5.6× |
| Goccy | 3336 | 543.09 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6315 | 286.94 MB/s | 3345 | 38 | 2.5× |
| Sonic | 6526 | 277.65 MB/s | 3344 | 38 | 2.4× |
| JSONV2 | 7958 | 227.68 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8315 | 217.79 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15844 | 114.37 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1021 | 1775.24 MB/s | 144 | 10 | 15.4× |
| Lightning | 1028 | 1762.06 MB/s | 144 | 10 | 15.3× |
| LightningDestructive | 1074 | 1686.71 MB/s | 144 | 10 | 14.7× |
| Easyjson | 2977 | 608.68 MB/s | 144 | 10 | 5.3× |
| Goccy | 3287 | 551.18 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6443 | 281.22 MB/s | 3365 | 40 | 2.4× |
| Sonic | 6652 | 272.38 MB/s | 3363 | 40 | 2.4× |
| JSONV2 | 7668 | 236.32 MB/s | 632 | 7 | 2.1× |
| LightningDecodeAny | 8396 | 215.70 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15767 | 114.92 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 657 | 751.57 MB/s | 160 | 1 | 9.4× |
| LightningDestructive | 669 | 738.03 MB/s | 160 | 1 | 9.2× |
| SonicFastest | 1232 | 401.04 MB/s | 1075 | 8 | 5.0× |
| Sonic | 1233 | 400.64 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1381 | 356.93 MB/s | 1296 | 26 | 4.5× |
| LightningArena | 1495 | 330.52 MB/s | 4096 | 1 | 4.1× |
| Easyjson | 2400 | 205.83 MB/s | 448 | 3 | 2.6× |
| Goccy | 2448 | 201.78 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3268 | 151.17 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6162 | 80.17 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 402 | 571.57 MB/s | 160 | 1 | 11.1× |
| LightningDestructive | 406 | 565.80 MB/s | 160 | 1 | 11.0× |
| Sonic | 890 | 258.41 MB/s | 802 | 8 | 5.0× |
| SonicFastest | 892 | 257.98 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1127 | 203.27 MB/s | 1296 | 26 | 4.0× |
| LightningArena | 1216 | 189.16 MB/s | 4096 | 1 | 3.7× |
| Easyjson | 1561 | 147.30 MB/s | 448 | 3 | 2.9× |
| Goccy | 1677 | 137.15 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2478 | 92.82 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4463 | 51.54 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 56160 | 1159.76 MB/s | 103441 | 103 | 11.9× |
| Lightning | 56595 | 1150.84 MB/s | 103441 | 103 | 11.8× |
| LightningDestructive | 57447 | 1133.78 MB/s | 97220 | 98 | 11.7× |
| SonicFastest | 152350 | 427.52 MB/s | 235753 | 65 | 4.4× |
| Sonic | 152473 | 427.17 MB/s | 235728 | 65 | 4.4× |
| LightningDecodeAny | 178909 | 298.08 MB/s | 180049 | 3245 | 3.7× |
| Goccy | 187617 | 347.15 MB/s | 227778 | 134 | 3.6× |
| JSONV2 | 252225 | 258.23 MB/s | 206660 | 607 | 2.7× |
| Stdlib | 670284 | 97.17 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2163799 | 896.79 MB/s | 2864592 | 1380 | 12.2× |
| LightningArena | 2226962 | 871.35 MB/s | 2864593 | 1380 | 11.9× |
| Lightning | 2244423 | 864.58 MB/s | 2864593 | 1380 | 11.8× |
| Goccy | 4950548 | 391.97 MB/s | 4063429 | 13509 | 5.4× |
| SonicFastest | 6391096 | 303.62 MB/s | 4884981 | 1736 | 4.1× |
| Sonic | 6407951 | 302.82 MB/s | 4886465 | 1736 | 4.1× |
| Easyjson | 7725786 | 251.17 MB/s | 3871264 | 15043 | 3.4× |
| LightningDecodeAny | 8921057 | 217.52 MB/s | 7063040 | 218633 | 3.0× |
| JSONV2 | 11207423 | 173.14 MB/s | 3237191 | 13947 | 2.4× |
| Stdlib | 26493752 | 73.24 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 834081 | 3989.82 MB/s | 351704 | 1286 | 31.1× |
| Lightning | 1373941 | 2422.11 MB/s | 2488905 | 2995 | 18.9× |
| LightningArena | 1374898 | 2420.42 MB/s | 2488905 | 2995 | 18.9× |
| SonicFastest | 2327603 | 1429.72 MB/s | 5896274 | 4263 | 11.1× |
| Sonic | 2330064 | 1428.21 MB/s | 5896346 | 4263 | 11.1× |
| LightningDecodeAny | 3257455 | 943.61 MB/s | 4876913 | 56892 | 8.0× |
| Goccy | 5221381 | 637.35 MB/s | 3948913 | 3816 | 5.0× |
| JSONV2 | 7880950 | 422.26 MB/s | 5364504 | 13243 | 3.3× |
| Stdlib | 25924918 | 128.36 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 169724 | 1298.26 MB/s | 135872 | 226 | 13.5× |
| Lightning | 171202 | 1287.05 MB/s | 135872 | 226 | 13.4× |
| LightningDestructive | 180782 | 1218.85 MB/s | 135872 | 226 | 12.7× |
| Goccy | 466440 | 472.40 MB/s | 364009 | 1066 | 4.9× |
| SonicFastest | 531941 | 414.23 MB/s | 351197 | 262 | 4.3× |
| Sonic | 533595 | 412.95 MB/s | 351434 | 262 | 4.3× |
| Easyjson | 598227 | 368.33 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 682223 | 322.98 MB/s | 129746 | 470 | 3.4× |
| LightningDecodeAny | 940996 | 115.10 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2295226 | 96.00 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 10091129 | 802.69 MB/s | 11845073 | 20816 | 10.1× |
| LightningArena | 10461352 | 774.28 MB/s | 11845073 | 20816 | 9.7× |
| Lightning | 10468788 | 773.73 MB/s | 11845078 | 20816 | 9.7× |
| Sonic | 21527728 | 376.26 MB/s | 19855003 | 41640 | 4.7× |
| SonicFastest | 21574420 | 375.45 MB/s | 19854024 | 41640 | 4.7× |
| Goccy | 26286608 | 308.14 MB/s | 19112261 | 107156 | 3.9× |
| Easyjson | 34417532 | 235.35 MB/s | 15059619 | 41643 | 3.0× |
| LightningDecodeAny | 37104659 | 140.23 MB/s | 46279351 | 747112 | 2.7× |
| JSONV2 | 46132303 | 175.58 MB/s | 15233743 | 78972 | 2.2× |
| Stdlib | 101542672 | 79.77 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4695809 | 635.35 MB/s | 3764713 | 1504 | 11.3× |
| LightningDestructive | 4925876 | 605.67 MB/s | 3758856 | 29356 | 10.8× |
| Lightning | 5052833 | 590.45 MB/s | 3758857 | 29356 | 10.5× |
| Sonic | 9245210 | 322.70 MB/s | 9131082 | 57804 | 5.7× |
| SonicFastest | 9279509 | 321.51 MB/s | 9131366 | 57804 | 5.7× |
| LightningDecodeAny | 17522742 | 104.68 MB/s | 23982579 | 351152 | 3.0× |
| Easyjson | 17953285 | 166.18 MB/s | 9479441 | 30115 | 3.0× |
| Goccy | 18072042 | 165.09 MB/s | 9877205 | 273621 | 2.9× |
| JSONV2 | 25150070 | 118.63 MB/s | 9257061 | 86278 | 2.1× |
| Stdlib | 52975143 | 56.32 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1128780 | 641.04 MB/s | 907600 | 3618 | 11.5× |
| LightningArena | 1154378 | 626.83 MB/s | 911395 | 30 | 11.2× |
| Lightning | 1204127 | 600.93 MB/s | 907598 | 3618 | 10.8× |
| SonicFastest | 2143386 | 337.60 MB/s | 2368269 | 3683 | 6.0× |
| Sonic | 2164593 | 334.29 MB/s | 2368253 | 3683 | 6.0× |
| Easyjson | 5234172 | 138.24 MB/s | 2847905 | 3698 | 2.5× |
| LightningDecodeAny | 5289257 | 123.00 MB/s | 6500461 | 76546 | 2.5× |
| Goccy | 5376691 | 134.58 MB/s | 2715117 | 80268 | 2.4× |
| JSONV2 | 6533276 | 110.76 MB/s | 2704708 | 7318 | 2.0× |
| Stdlib | 12959840 | 55.83 MB/s | 2704546 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1553575 | 1015.31 MB/s | 911393 | 30 | 11.6× |
| LightningDestructive | 1561289 | 1010.29 MB/s | 907600 | 3618 | 11.6× |
| Lightning | 1616221 | 975.95 MB/s | 907597 | 3618 | 11.2× |
| Sonic | 2470030 | 638.60 MB/s | 3222782 | 3683 | 7.3× |
| SonicFastest | 2483285 | 635.19 MB/s | 3221351 | 3683 | 7.3× |
| LightningDecodeAny | 4535114 | 166.13 MB/s | 6500456 | 76546 | 4.0× |
| Easyjson | 6505881 | 242.45 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 6606531 | 238.76 MB/s | 3472475 | 80261 | 2.7× |
| JSONV2 | 6854834 | 230.11 MB/s | 2704553 | 7318 | 2.6× |
| Stdlib | 18047358 | 87.40 MB/s | 2704547 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 204865 | 732.80 MB/s | 81920 | 1 | 10.0× |
| Lightning | 205308 | 731.21 MB/s | 81920 | 1 | 10.0× |
| LightningDestructive | 209187 | 717.65 MB/s | 81920 | 1 | 9.8× |
| Sonic | 383998 | 390.95 MB/s | 407244 | 16 | 5.3× |
| SonicFastest | 388593 | 386.33 MB/s | 407295 | 16 | 5.3× |
| LightningDecodeAny | 548229 | 273.83 MB/s | 745765 | 10016 | 3.7× |
| Goccy | 982207 | 152.84 MB/s | 324144 | 10005 | 2.1× |
| JSONV2 | 1149126 | 130.64 MB/s | 357731 | 20 | 1.8× |
| Stdlib | 2044183 | 73.44 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 25713 | 1093.49 MB/s | 29216 | 103 | 13.2× |
| Lightning | 25955 | 1083.31 MB/s | 29216 | 103 | 13.1× |
| LightningDestructive | 27107 | 1037.27 MB/s | 29088 | 101 | 12.5× |
| Sonic | 67103 | 419.01 MB/s | 59444 | 83 | 5.1× |
| SonicFastest | 67479 | 416.68 MB/s | 59458 | 83 | 5.0× |
| Easyjson | 75119 | 374.30 MB/s | 32304 | 138 | 4.5× |
| Goccy | 80403 | 349.70 MB/s | 59302 | 188 | 4.2× |
| JSONV2 | 128119 | 219.46 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 148579 | 189.24 MB/s | 140576 | 2643 | 2.3× |
| Stdlib | 339461 | 82.83 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1462 | 1592.41 MB/s | 32 | 1 | 17.0× |
| Lightning | 1466 | 1588.10 MB/s | 32 | 1 | 17.0× |
| LightningDestructive | 1546 | 1506.04 MB/s | 32 | 1 | 16.1× |
| Goccy | 4697 | 495.68 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 5115 | 455.16 MB/s | 192 | 2 | 4.9× |
| Sonic | 6119 | 380.47 MB/s | 3708 | 4 | 4.1× |
| SonicFastest | 6133 | 379.60 MB/s | 3709 | 4 | 4.1× |
| JSONV2 | 7926 | 293.73 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 9466 | 178.00 MB/s | 10200 | 195 | 2.6× |
| Stdlib | 24875 | 93.59 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 153 | 1234.55 MB/s | 0 | 0 | 17.2× |
| LightningArena | 153 | 1232.62 MB/s | 0 | 0 | 17.1× |
| LightningDestructive | 157 | 1203.52 MB/s | 0 | 0 | 16.7× |
| Goccy | 432 | 437.34 MB/s | 304 | 2 | 6.1× |
| Easyjson | 556 | 340.08 MB/s | 0 | 0 | 4.7× |
| Sonic | 718 | 263.10 MB/s | 341 | 3 | 3.7× |
| SonicFastest | 720 | 262.36 MB/s | 342 | 3 | 3.6× |
| JSONV2 | 947 | 199.49 MB/s | 112 | 1 | 2.8× |
| LightningDecodeAny | 1184 | 113.19 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 2627 | 71.94 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1068 | 2051.83 MB/s | 0 | 0 | 17.4× |
| LightningArena | 1082 | 2025.49 MB/s | 0 | 0 | 17.2× |
| LightningDestructive | 1115 | 1964.97 MB/s | 0 | 0 | 16.7× |
| Easyjson | 3559 | 615.60 MB/s | 24 | 1 | 5.2× |
| Goccy | 3762 | 582.34 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6870 | 318.91 MB/s | 3603 | 38 | 2.7× |
| Sonic | 7022 | 312.03 MB/s | 3600 | 38 | 2.6× |
| JSONV2 | 8172 | 268.11 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8386 | 215.97 MB/s | 7552 | 158 | 2.2× |
| Stdlib | 18572 | 117.97 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 482053 | 1058.96 MB/s | 457538 | 1009 | 14.1× |
| LightningArena | 491661 | 1038.27 MB/s | 457538 | 1009 | 13.8× |
| Lightning | 502162 | 1016.56 MB/s | 457537 | 1009 | 13.6× |
| Goccy | 1257750 | 405.86 MB/s | 1135618 | 5006 | 5.4× |
| Sonic | 1453758 | 351.14 MB/s | 1309738 | 2014 | 4.7× |
| SonicFastest | 1464051 | 348.67 MB/s | 1310525 | 2014 | 4.7× |
| Easyjson | 1589526 | 321.15 MB/s | 863780 | 3012 | 4.3× |
| JSONV2 | 3173374 | 160.86 MB/s | 1075954 | 12645 | 2.1× |
| LightningDecodeAny | 3380567 | 136.51 MB/s | 2950651 | 64018 | 2.0× |
| Stdlib | 6808690 | 74.97 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 611 | 32406.48 MB/s | 0 | 0 | 255.1× |
| LightningArena | 611 | 32374.99 MB/s | 0 | 0 | 254.8× |
| LightningDestructive | 863 | 22929.17 MB/s | 0 | 0 | 180.5× |
| SonicFastest | 6927 | 2856.98 MB/s | 21095 | 3 | 22.5× |
| Goccy | 25686 | 770.42 MB/s | 20492 | 2 | 6.1× |
| Sonic | 29212 | 677.43 MB/s | 20623 | 3 | 5.3× |
| JSONV2 | 36202 | 546.63 MB/s | 8 | 1 | 4.3× |
| LightningDecodeAny | 96819 | 204.38 MB/s | 116864 | 2015 | 1.6× |
| Easyjson | 120204 | 164.63 MB/s | 0 | 0 | 1.3× |
| Stdlib | 155745 | 127.06 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1757 | 10317.58 MB/s | 432 | 2 | 77.4× |
| Lightning | 1767 | 10256.94 MB/s | 432 | 2 | 76.9× |
| LightningDestructive | 1861 | 9741.00 MB/s | 0 | 0 | 73.1× |
| Easyjson | 4711 | 3847.27 MB/s | 432 | 2 | 28.9× |
| Sonic | 9363 | 1935.63 MB/s | 20433 | 5 | 14.5× |
| SonicFastest | 9409 | 1926.15 MB/s | 20461 | 5 | 14.5× |
| LightningDecodeAny | 18117 | 987.00 MB/s | 29088 | 191 | 7.5× |
| Goccy | 26364 | 687.45 MB/s | 19460 | 2 | 5.2× |
| JSONV2 | 48285 | 375.36 MB/s | 16500 | 50 | 2.8× |
| Stdlib | 135970 | 133.29 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1982174 | 1013.28 MB/s | 3089565 | 6821 | 10.7× |
| Lightning | 2082207 | 964.60 MB/s | 3091278 | 6827 | 10.2× |
| LightningArena | 2094900 | 958.75 MB/s | 3094370 | 6703 | 10.2× |
| Goccy | 4601808 | 436.46 MB/s | 5410739 | 15832 | 4.6× |
| SonicFastest | 5009679 | 400.92 MB/s | 5153683 | 7085 | 4.2× |
| Sonic | 5065396 | 396.51 MB/s | 5155326 | 7085 | 4.2× |
| Easyjson | 5533041 | 363.00 MB/s | 2981492 | 7439 | 3.8× |
| LightningDecodeAny | 6723300 | 169.90 MB/s | 8503511 | 134008 | 3.2× |
| JSONV2 | 7263803 | 276.51 MB/s | 3173682 | 14563 | 2.9× |
| Stdlib | 21266289 | 94.44 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 799 | 686.86 MB/s | 480 | 1 | 7.7× |
| Lightning | 807 | 680.45 MB/s | 480 | 1 | 7.6× |
| LightningDestructive | 819 | 670.68 MB/s | 480 | 1 | 7.5× |
| LightningDecodeAny | 1694 | 323.51 MB/s | 2021 | 46 | 3.6× |
| Easyjson | 1969 | 278.75 MB/s | 1616 | 5 | 3.1× |
| SonicFastest | 2115 | 259.58 MB/s | 2264 | 8 | 2.9× |
| Sonic | 2165 | 253.54 MB/s | 2263 | 8 | 2.8× |
| JSONV2 | 2945 | 186.40 MB/s | 1664 | 7 | 2.1× |
| Goccy | 2999 | 183.04 MB/s | 2129 | 43 | 2.0× |
| Stdlib | 6144 | 89.36 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 403706 | 1564.29 MB/s | 402728 | 545 | 15.5× |
| Lightning | 476814 | 1324.44 MB/s | 451257 | 857 | 13.1× |
| LightningArena | 477184 | 1323.42 MB/s | 453017 | 712 | 13.1× |
| Sonic | 1103580 | 572.24 MB/s | 1068673 | 814 | 5.7× |
| SonicFastest | 1134552 | 556.62 MB/s | 1067322 | 814 | 5.5× |
| Easyjson | 1289056 | 489.90 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1394916 | 452.73 MB/s | 990282 | 1200 | 4.5× |
| JSONV2 | 2202912 | 286.67 MB/s | 571590 | 3144 | 2.8× |
| LightningDecodeAny | 2341001 | 199.45 MB/s | 2076503 | 30126 | 2.7× |
| Stdlib | 6263853 | 100.82 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 573171 | 981.22 MB/s | 546569 | 429 | 10.4× |
| Lightning | 756290 | 743.64 MB/s | 769937 | 1235 | 7.9× |
| LightningArena | 765986 | 734.23 MB/s | 771665 | 1088 | 7.8× |
| Sonic | 1415634 | 397.28 MB/s | 1345293 | 1184 | 4.2× |
| SonicFastest | 1420001 | 396.06 MB/s | 1345526 | 1184 | 4.2× |
| Goccy | 1598209 | 351.90 MB/s | 1039826 | 1028 | 3.7× |
| Easyjson | 2039175 | 275.80 MB/s | 775154 | 1254 | 2.9× |
| LightningDecodeAny | 2813054 | 199.93 MB/s | 2180441 | 30126 | 2.1× |
| JSONV2 | 3208734 | 175.27 MB/s | 927406 | 3482 | 1.9× |
| Stdlib | 5942989 | 94.63 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 534764 | 997.03 MB/s | 333416 | 2084 | 12.0× |
| Lightning | 607737 | 877.32 MB/s | 368224 | 2293 | 10.5× |
| LightningArena | 617027 | 864.11 MB/s | 368224 | 2293 | 10.4× |
| Easyjson | 1282496 | 415.73 MB/s | 428362 | 3273 | 5.0× |
| Sonic | 1391386 | 383.20 MB/s | 981909 | 3082 | 4.6× |
| SonicFastest | 1402319 | 380.21 MB/s | 982092 | 3082 | 4.6× |
| Goccy | 1554147 | 343.07 MB/s | 1167084 | 5409 | 4.1× |
| JSONV2 | 2722497 | 195.84 MB/s | 745420 | 13288 | 2.4× |
| LightningDecodeAny | 3469822 | 153.66 MB/s | 2992874 | 50076 | 1.8× |
| Stdlib | 6399177 | 83.32 MB/s | 798692 | 17133 | 1.0× |
