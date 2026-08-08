# JSON Deserialization Benchmarks

- generated 2026-08-08T12:49:34Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 105019 | 1211.93 MB/s | 49760 | 3 | 10.5× |
| LightningDestructive | 105094 | 1211.06 MB/s | 49280 | 2 | 10.5× |
| Lightning | 105114 | 1210.82 MB/s | 49760 | 3 | 10.5× |
| Sonic | 179684 | 708.33 MB/s | 193035 | 10 | 6.2× |
| SonicFastest | 180847 | 703.77 MB/s | 197175 | 10 | 6.1× |
| Goccy | 197887 | 643.17 MB/s | 225333 | 884 | 5.6× |
| Easyjson | 211533 | 601.68 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 415860 | 306.05 MB/s | 195120 | 1805 | 2.7× |
| LightningDecodeAny | 426677 | 221.84 MB/s | 463409 | 9708 | 2.6× |
| Stdlib | 1105953 | 115.08 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3607430 | 624.00 MB/s | 2532848 | 1143 | 7.2× |
| Lightning | 3637612 | 618.83 MB/s | 2532851 | 1143 | 7.2× |
| LightningArena | 3640361 | 618.36 MB/s | 2532849 | 1143 | 7.2× |
| SonicFastest | 4526334 | 497.32 MB/s | 15233851 | 970 | 5.8× |
| Sonic | 4618292 | 487.42 MB/s | 15233759 | 970 | 5.7× |
| Goccy | 10218032 | 220.30 MB/s | 4110102 | 56532 | 2.6× |
| Easyjson | 10964620 | 205.30 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12255276 | 183.68 MB/s | 19380210 | 223896 | 2.1× |
| JSONV2 | 16119462 | 139.65 MB/s | 3123229 | 3083 | 1.6× |
| Stdlib | 26124605 | 86.17 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 463544 | 583.34 MB/s | 397296 | 567 | 7.2× |
| Lightning | 465072 | 581.42 MB/s | 397296 | 567 | 7.2× |
| LightningArena | 465294 | 581.14 MB/s | 397296 | 567 | 7.2× |
| SonicFastest | 621399 | 435.15 MB/s | 466776 | 968 | 5.4× |
| Sonic | 621988 | 434.74 MB/s | 466101 | 968 | 5.4× |
| Easyjson | 1400154 | 193.12 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1405930 | 192.33 MB/s | 542980 | 8122 | 2.4× |
| LightningDecodeAny | 1640964 | 164.78 MB/s | 2543876 | 29687 | 2.0× |
| JSONV2 | 2102578 | 128.61 MB/s | 348151 | 1628 | 1.6× |
| Stdlib | 3344498 | 80.85 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1150082 | 1501.81 MB/s | 765602 | 2799 | 11.5× |
| LightningDestructive | 1153040 | 1497.96 MB/s | 765560 | 2798 | 11.5× |
| LightningArena | 1155440 | 1494.85 MB/s | 768416 | 2440 | 11.5× |
| SonicFastest | 1999641 | 863.76 MB/s | 2709159 | 4020 | 6.6× |
| Sonic | 2012401 | 858.28 MB/s | 2730300 | 4020 | 6.6× |
| Goccy | 2356124 | 733.07 MB/s | 2583077 | 14605 | 5.6× |
| Easyjson | 4211056 | 410.16 MB/s | 972032 | 5389 | 3.1× |
| LightningDecodeAny | 4213761 | 118.73 MB/s | 4953693 | 76576 | 3.1× |
| JSONV2 | 4263913 | 405.07 MB/s | 1011634 | 7594 | 3.1× |
| Stdlib | 13244846 | 130.41 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1163 | 1557.39 MB/s | 0 | 0 | 12.1× |
| Lightning | 1169 | 1549.84 MB/s | 0 | 0 | 12.0× |
| LightningDestructive | 1182 | 1533.55 MB/s | 0 | 0 | 11.9× |
| Easyjson | 2541 | 713.23 MB/s | 24 | 1 | 5.5× |
| Goccy | 2813 | 644.16 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5934 | 305.37 MB/s | 3756 | 40 | 2.4× |
| SonicFastest | 5936 | 305.27 MB/s | 3748 | 40 | 2.4× |
| JSONV2 | 7800 | 232.30 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7903 | 229.15 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 14086 | 128.64 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1193 | 1518.59 MB/s | 0 | 0 | 11.8× |
| Lightning | 1195 | 1516.02 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1207 | 1500.74 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2539 | 713.76 MB/s | 24 | 1 | 5.5× |
| Goccy | 2858 | 634.09 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6072 | 298.41 MB/s | 3735 | 40 | 2.3× |
| SonicFastest | 6074 | 298.33 MB/s | 3765 | 40 | 2.3× |
| JSONV2 | 7671 | 236.23 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7930 | 228.36 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 14072 | 128.76 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1350 | 1342.71 MB/s | 144 | 10 | 10.4× |
| Lightning | 1351 | 1341.56 MB/s | 144 | 10 | 10.4× |
| LightningDestructive | 1411 | 1284.11 MB/s | 144 | 10 | 9.9× |
| Easyjson | 2784 | 650.90 MB/s | 144 | 10 | 5.0× |
| Goccy | 2846 | 636.65 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6108 | 296.68 MB/s | 3771 | 42 | 2.3× |
| SonicFastest | 6110 | 296.55 MB/s | 3763 | 42 | 2.3× |
| LightningDecodeAny | 7905 | 229.09 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7995 | 226.64 MB/s | 632 | 7 | 1.8× |
| Stdlib | 13997 | 129.45 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 701 | 705.09 MB/s | 160 | 1 | 7.9× |
| Lightning | 704 | 701.23 MB/s | 160 | 1 | 7.8× |
| SonicFastest | 1234 | 400.26 MB/s | 1023 | 6 | 4.5× |
| Sonic | 1240 | 398.40 MB/s | 1028 | 6 | 4.5× |
| LightningArena | 1363 | 362.31 MB/s | 4096 | 1 | 4.1× |
| LightningDecodeAny | 1421 | 346.85 MB/s | 1296 | 26 | 3.9× |
| Easyjson | 2235 | 220.99 MB/s | 448 | 3 | 2.5× |
| Goccy | 2438 | 202.61 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3238 | 152.54 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5523 | 89.44 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 432 | 533.07 MB/s | 160 | 1 | 9.6× |
| Lightning | 436 | 528.14 MB/s | 160 | 1 | 9.5× |
| Sonic | 890 | 258.38 MB/s | 692 | 6 | 4.7× |
| SonicFastest | 894 | 257.31 MB/s | 685 | 6 | 4.6× |
| LightningArena | 1119 | 205.51 MB/s | 4096 | 1 | 3.7× |
| LightningDecodeAny | 1171 | 195.62 MB/s | 1296 | 26 | 3.5× |
| Easyjson | 1409 | 163.25 MB/s | 448 | 3 | 2.9× |
| Goccy | 1596 | 144.13 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2472 | 93.05 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4156 | 55.34 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 59819 | 1088.81 MB/s | 103440 | 103 | 9.2× |
| Lightning | 60114 | 1083.48 MB/s | 103440 | 103 | 9.1× |
| LightningDestructive | 61042 | 1067.01 MB/s | 97220 | 98 | 9.0× |
| Sonic | 99751 | 652.94 MB/s | 156623 | 75 | 5.5× |
| SonicFastest | 100108 | 650.62 MB/s | 156963 | 75 | 5.5× |
| Goccy | 148961 | 437.24 MB/s | 229296 | 134 | 3.7× |
| LightningDecodeAny | 177644 | 300.20 MB/s | 180048 | 3245 | 3.1× |
| JSONV2 | 225624 | 288.67 MB/s | 206650 | 607 | 2.4× |
| Stdlib | 549737 | 118.48 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2599347 | 746.52 MB/s | 2864592 | 1380 | 9.0× |
| Lightning | 2671766 | 726.29 MB/s | 2864595 | 1380 | 8.8× |
| LightningArena | 2679160 | 724.28 MB/s | 2864594 | 1380 | 8.8× |
| Goccy | 4759476 | 407.71 MB/s | 4065792 | 13510 | 4.9× |
| SonicFastest | 4798024 | 404.43 MB/s | 14608615 | 1407 | 4.9× |
| Sonic | 4851434 | 399.98 MB/s | 14608631 | 1407 | 4.8× |
| Easyjson | 7475449 | 259.58 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 9288460 | 208.91 MB/s | 7063041 | 218633 | 2.5× |
| JSONV2 | 11220682 | 172.94 MB/s | 3237226 | 13947 | 2.1× |
| Stdlib | 23463893 | 82.70 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1095780 | 3036.95 MB/s | 351704 | 1286 | 19.1× |
| LightningArena | 1765079 | 1885.37 MB/s | 2488905 | 2995 | 11.8× |
| Lightning | 1774595 | 1875.26 MB/s | 2488906 | 2995 | 11.8× |
| Sonic | 2722401 | 1222.39 MB/s | 6530582 | 4248 | 7.7× |
| SonicFastest | 2735057 | 1216.73 MB/s | 6486116 | 4248 | 7.6× |
| LightningDecodeAny | 3571040 | 860.75 MB/s | 4876913 | 56892 | 5.8× |
| Goccy | 4659291 | 714.24 MB/s | 3948909 | 3816 | 4.5× |
| JSONV2 | 7526656 | 442.14 MB/s | 5364520 | 13243 | 2.8× |
| Stdlib | 20881078 | 159.37 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 219063 | 1005.86 MB/s | 135872 | 226 | 9.3× |
| Lightning | 219144 | 1005.49 MB/s | 135872 | 226 | 9.3× |
| LightningDestructive | 220801 | 997.94 MB/s | 135872 | 226 | 9.2× |
| Sonic | 377386 | 583.87 MB/s | 299261 | 398 | 5.4× |
| SonicFastest | 378439 | 582.25 MB/s | 304536 | 398 | 5.4× |
| Goccy | 427964 | 514.87 MB/s | 364327 | 1067 | 4.8× |
| Easyjson | 546825 | 402.96 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 722116 | 305.14 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 843091 | 128.47 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2038174 | 108.11 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11379871 | 711.79 MB/s | 11845073 | 20816 | 7.8× |
| LightningArena | 11733695 | 690.32 MB/s | 11845073 | 20816 | 7.6× |
| Lightning | 11769164 | 688.24 MB/s | 11845073 | 20816 | 7.6× |
| SonicFastest | 16948998 | 477.91 MB/s | 70873024 | 40014 | 5.3× |
| Sonic | 16957635 | 477.66 MB/s | 70902086 | 40014 | 5.3× |
| Goccy | 23666639 | 342.26 MB/s | 17344038 | 107149 | 3.8× |
| Easyjson | 30754086 | 263.38 MB/s | 15059617 | 41643 | 2.9× |
| LightningDecodeAny | 36622108 | 142.07 MB/s | 46279352 | 747112 | 2.4× |
| JSONV2 | 43540817 | 186.03 MB/s | 15233733 | 78972 | 2.0× |
| Stdlib | 89082607 | 90.93 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5415547 | 550.91 MB/s | 3764715 | 1504 | 8.7× |
| LightningDestructive | 5727982 | 520.86 MB/s | 3758856 | 29356 | 8.2× |
| Lightning | 5839550 | 510.91 MB/s | 3758857 | 29356 | 8.1× |
| SonicFastest | 8641298 | 345.26 MB/s | 26622339 | 56760 | 5.5× |
| Sonic | 8664076 | 344.35 MB/s | 26578547 | 56760 | 5.4× |
| Easyjson | 16415803 | 181.74 MB/s | 9479440 | 30115 | 2.9× |
| LightningDecodeAny | 16471419 | 111.36 MB/s | 23982584 | 351152 | 2.9× |
| Goccy | 16488023 | 180.95 MB/s | 10556639 | 273645 | 2.9× |
| JSONV2 | 24505811 | 121.75 MB/s | 9257161 | 86278 | 1.9× |
| Stdlib | 47137980 | 63.29 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1196249 | 604.89 MB/s | 911392 | 30 | 9.7× |
| LightningDestructive | 1234918 | 585.95 MB/s | 907601 | 3618 | 9.4× |
| Lightning | 1243433 | 581.93 MB/s | 907595 | 3618 | 9.3× |
| SonicFastest | 1742985 | 415.15 MB/s | 3187957 | 7226 | 6.7× |
| Sonic | 1783052 | 405.82 MB/s | 3187488 | 7226 | 6.5× |
| Easyjson | 4235737 | 170.83 MB/s | 2847906 | 3698 | 2.7× |
| LightningDecodeAny | 4252083 | 153.00 MB/s | 6500461 | 76546 | 2.7× |
| Goccy | 4747612 | 152.41 MB/s | 2853604 | 80276 | 2.4× |
| JSONV2 | 5579135 | 129.70 MB/s | 2704618 | 7318 | 2.1× |
| Stdlib | 11604295 | 62.36 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1827057 | 863.33 MB/s | 911393 | 30 | 8.6× |
| LightningDestructive | 1830401 | 861.75 MB/s | 907601 | 3618 | 8.6× |
| Lightning | 1879496 | 839.24 MB/s | 907595 | 3618 | 8.4× |
| SonicFastest | 2250644 | 700.85 MB/s | 5803016 | 7226 | 7.0× |
| Sonic | 2251099 | 700.70 MB/s | 5804744 | 7226 | 7.0× |
| LightningDecodeAny | 3839155 | 196.24 MB/s | 6500460 | 76546 | 4.1× |
| Easyjson | 5546909 | 284.37 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5714399 | 276.03 MB/s | 3593135 | 80268 | 2.7× |
| JSONV2 | 6333149 | 249.06 MB/s | 2704593 | 7318 | 2.5× |
| Stdlib | 15699039 | 100.47 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 211720 | 709.07 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 211775 | 708.89 MB/s | 81920 | 1 | 8.6× |
| LightningArena | 211898 | 708.47 MB/s | 81920 | 1 | 8.6× |
| Sonic | 272277 | 551.37 MB/s | 254124 | 6 | 6.7× |
| SonicFastest | 278231 | 539.57 MB/s | 271651 | 6 | 6.5× |
| LightningDecodeAny | 479344 | 313.18 MB/s | 745765 | 10016 | 3.8× |
| Goccy | 864494 | 173.66 MB/s | 324120 | 10004 | 2.1× |
| JSONV2 | 1087123 | 138.09 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1822008 | 82.39 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 31666 | 887.93 MB/s | 29216 | 103 | 9.6× |
| Lightning | 31678 | 887.59 MB/s | 29216 | 103 | 9.6× |
| LightningDestructive | 31916 | 880.97 MB/s | 29088 | 101 | 9.5× |
| SonicFastest | 62988 | 446.39 MB/s | 46974 | 103 | 4.8× |
| Sonic | 63054 | 445.92 MB/s | 47017 | 103 | 4.8× |
| Easyjson | 68786 | 408.76 MB/s | 32304 | 138 | 4.4× |
| Goccy | 71194 | 394.93 MB/s | 59227 | 188 | 4.3× |
| JSONV2 | 133437 | 210.71 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 145416 | 193.35 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 302746 | 92.87 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1957 | 1189.76 MB/s | 32 | 1 | 11.6× |
| LightningArena | 1965 | 1184.65 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2033 | 1145.21 MB/s | 32 | 1 | 11.2× |
| Goccy | 4149 | 561.09 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4263 | 546.09 MB/s | 192 | 2 | 5.3× |
| SonicFastest | 5058 | 460.25 MB/s | 4309 | 6 | 4.5× |
| Sonic | 5080 | 458.31 MB/s | 4313 | 6 | 4.5× |
| JSONV2 | 8348 | 278.87 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9849 | 171.08 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22723 | 102.45 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 221 | 854.35 MB/s | 0 | 0 | 11.0× |
| Lightning | 221 | 853.58 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 224 | 844.56 MB/s | 0 | 0 | 10.9× |
| Goccy | 379 | 498.14 MB/s | 304 | 2 | 6.4× |
| Easyjson | 489 | 386.35 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 782 | 241.81 MB/s | 507 | 4 | 3.1× |
| Sonic | 788 | 239.97 MB/s | 505 | 4 | 3.1× |
| JSONV2 | 1030 | 183.44 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1173 | 114.24 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2437 | 77.56 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1516 | 1444.87 MB/s | 0 | 0 | 10.6× |
| LightningArena | 1517 | 1444.37 MB/s | 0 | 0 | 10.6× |
| LightningDestructive | 1543 | 1420.25 MB/s | 0 | 0 | 10.4× |
| Goccy | 3160 | 693.26 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3202 | 684.19 MB/s | 24 | 1 | 5.0× |
| Sonic | 6469 | 338.69 MB/s | 3987 | 40 | 2.5× |
| SonicFastest | 6474 | 338.44 MB/s | 3973 | 40 | 2.5× |
| JSONV2 | 7907 | 277.11 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 7943 | 227.99 MB/s | 7552 | 158 | 2.0× |
| Stdlib | 16040 | 136.59 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 646184 | 789.99 MB/s | 457536 | 1009 | 9.3× |
| Lightning | 654881 | 779.49 MB/s | 457537 | 1009 | 9.1× |
| LightningArena | 656173 | 777.96 MB/s | 457536 | 1009 | 9.1× |
| Goccy | 1144189 | 446.15 MB/s | 1137379 | 5006 | 5.2× |
| SonicFastest | 1162507 | 439.12 MB/s | 907347 | 2006 | 5.2× |
| Sonic | 1162976 | 438.94 MB/s | 902310 | 2006 | 5.2× |
| Easyjson | 1525429 | 334.64 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3183208 | 160.37 MB/s | 1076008 | 12646 | 1.9× |
| LightningDecodeAny | 3243623 | 142.27 MB/s | 2950648 | 64018 | 1.8× |
| Stdlib | 5991778 | 85.20 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1336 | 14809.00 MB/s | 0 | 0 | 83.6× |
| LightningArena | 1337 | 14804.56 MB/s | 0 | 0 | 83.5× |
| LightningDestructive | 1359 | 14562.12 MB/s | 0 | 0 | 82.1× |
| Goccy | 19948 | 992.04 MB/s | 20491 | 2 | 5.6× |
| Sonic | 27851 | 710.53 MB/s | 22012 | 4 | 4.0× |
| SonicFastest | 28017 | 706.31 MB/s | 22392 | 4 | 4.0× |
| JSONV2 | 29543 | 669.83 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 74281 | 266.39 MB/s | 116864 | 2015 | 1.5× |
| Easyjson | 86095 | 229.85 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111630 | 177.27 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2739 | 6616.23 MB/s | 0 | 0 | 37.4× |
| Lightning | 2749 | 6593.29 MB/s | 432 | 2 | 37.3× |
| LightningArena | 2749 | 6592.41 MB/s | 432 | 2 | 37.3× |
| Easyjson | 3935 | 4605.77 MB/s | 432 | 2 | 26.1× |
| Sonic | 10134 | 1788.43 MB/s | 23005 | 6 | 10.1× |
| SonicFastest | 10540 | 1719.47 MB/s | 23823 | 6 | 9.7× |
| Goccy | 16334 | 1109.58 MB/s | 19459 | 2 | 6.3× |
| LightningDecodeAny | 16472 | 1085.57 MB/s | 29088 | 191 | 6.2× |
| JSONV2 | 44954 | 403.16 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102538 | 176.75 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2455263 | 818.04 MB/s | 3089564 | 6821 | 7.6× |
| LightningArena | 2511802 | 799.62 MB/s | 3094370 | 6703 | 7.5× |
| Lightning | 2535243 | 792.23 MB/s | 3091278 | 6827 | 7.4× |
| Goccy | 4418480 | 454.57 MB/s | 5412000 | 15831 | 4.2× |
| Sonic | 4492875 | 447.04 MB/s | 10948678 | 13683 | 4.2× |
| SonicFastest | 4538927 | 442.50 MB/s | 10948854 | 13683 | 4.1× |
| Easyjson | 4948375 | 405.89 MB/s | 2981484 | 7439 | 3.8× |
| LightningDecodeAny | 6795378 | 168.10 MB/s | 8503513 | 134008 | 2.8× |
| JSONV2 | 7023675 | 285.96 MB/s | 3173702 | 14563 | 2.7× |
| Stdlib | 18744288 | 107.15 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 903 | 608.16 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 910 | 603.08 MB/s | 480 | 1 | 6.3× |
| LightningArena | 912 | 601.65 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1732 | 316.39 MB/s | 2021 | 46 | 3.3× |
| Easyjson | 2266 | 242.28 MB/s | 1616 | 5 | 2.5× |
| Sonic | 2701 | 203.28 MB/s | 2053 | 26 | 2.1× |
| SonicFastest | 2711 | 202.48 MB/s | 2069 | 26 | 2.1× |
| Goccy | 3046 | 180.24 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3393 | 161.81 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5705 | 96.23 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 499005 | 1265.55 MB/s | 402728 | 545 | 10.8× |
| LightningArena | 561635 | 1124.42 MB/s | 453017 | 712 | 9.6× |
| Lightning | 563973 | 1119.76 MB/s | 451257 | 857 | 9.6× |
| Sonic | 1019851 | 619.22 MB/s | 1003178 | 1102 | 5.3× |
| SonicFastest | 1025258 | 615.96 MB/s | 1002758 | 1102 | 5.3× |
| Easyjson | 1142654 | 552.67 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1165445 | 541.86 MB/s | 986364 | 1201 | 4.6× |
| JSONV2 | 2135377 | 295.74 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2355985 | 198.18 MB/s | 2076505 | 30126 | 2.3× |
| Stdlib | 5399566 | 116.96 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 611327 | 919.98 MB/s | 546573 | 429 | 8.6× |
| LightningArena | 767527 | 732.75 MB/s | 771666 | 1088 | 6.9× |
| Lightning | 774735 | 725.94 MB/s | 769937 | 1235 | 6.8× |
| Sonic | 1036576 | 542.56 MB/s | 943008 | 1476 | 5.1× |
| SonicFastest | 1049876 | 535.69 MB/s | 954340 | 1476 | 5.0× |
| Goccy | 1334869 | 421.32 MB/s | 1039369 | 1030 | 4.0× |
| Easyjson | 1745726 | 322.16 MB/s | 775153 | 1254 | 3.0× |
| LightningDecodeAny | 2567214 | 219.07 MB/s | 2180440 | 30126 | 2.1× |
| JSONV2 | 2764933 | 203.41 MB/s | 927445 | 3482 | 1.9× |
| Stdlib | 5282367 | 106.47 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 653571 | 815.79 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 678246 | 786.11 MB/s | 368224 | 2293 | 8.1× |
| LightningArena | 679024 | 785.21 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1100981 | 484.28 MB/s | 428361 | 3273 | 5.0× |
| SonicFastest | 1138775 | 468.20 MB/s | 1044535 | 4351 | 4.8× |
| Sonic | 1141501 | 467.08 MB/s | 1038450 | 4351 | 4.8× |
| Goccy | 1318466 | 404.39 MB/s | 1167242 | 5409 | 4.2× |
| JSONV2 | 2515394 | 211.97 MB/s | 745449 | 13288 | 2.2× |
| LightningDecodeAny | 3322196 | 160.49 MB/s | 2992878 | 50076 | 1.6× |
| Stdlib | 5472259 | 97.43 MB/s | 798692 | 17133 | 1.0× |
