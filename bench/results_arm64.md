# JSON Deserialization Benchmarks

- generated 2026-09-07T19:12:00Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 84450 | 1507.11 MB/s | 49280 | 2 | 13.0× |
| Lightning | 84517 | 1505.92 MB/s | 49760 | 3 | 12.9× |
| LightningArena | 84598 | 1504.46 MB/s | 49760 | 3 | 12.9× |
| SonicFastest | 184252 | 690.77 MB/s | 197905 | 10 | 5.9× |
| Sonic | 184893 | 688.37 MB/s | 199273 | 10 | 5.9× |
| Goccy | 199191 | 638.96 MB/s | 224941 | 884 | 5.5× |
| Easyjson | 213632 | 595.77 MB/s | 122864 | 14 | 5.1× |
| JSONV2 | 422421 | 301.30 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 440109 | 215.07 MB/s | 463409 | 9708 | 2.5× |
| Stdlib | 1094325 | 116.30 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2664129 | 844.95 MB/s | 2532848 | 1143 | 10.0× |
| LightningArena | 2682275 | 839.23 MB/s | 2532849 | 1143 | 9.9× |
| Lightning | 2691749 | 836.28 MB/s | 2532850 | 1143 | 9.9× |
| SonicFastest | 4594586 | 489.94 MB/s | 15243685 | 970 | 5.8× |
| Sonic | 4628601 | 486.34 MB/s | 15240425 | 970 | 5.7× |
| Goccy | 10520348 | 213.97 MB/s | 4125373 | 56532 | 2.5× |
| Easyjson | 10970927 | 205.18 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11764408 | 191.34 MB/s | 19380211 | 223896 | 2.3× |
| JSONV2 | 16234687 | 138.66 MB/s | 3123205 | 3083 | 1.6× |
| Stdlib | 26515369 | 84.90 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 377309 | 716.66 MB/s | 397296 | 567 | 9.1× |
| LightningDestructive | 377568 | 716.17 MB/s | 397296 | 567 | 9.1× |
| Lightning | 377776 | 715.78 MB/s | 397297 | 567 | 9.1× |
| SonicFastest | 638382 | 423.58 MB/s | 483940 | 968 | 5.4× |
| Sonic | 638488 | 423.51 MB/s | 485861 | 968 | 5.4× |
| Easyjson | 1386019 | 195.09 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1420183 | 190.40 MB/s | 543111 | 8122 | 2.4× |
| LightningDecodeAny | 1648833 | 164.00 MB/s | 2543881 | 29687 | 2.1× |
| JSONV2 | 2118335 | 127.65 MB/s | 348153 | 1628 | 1.6× |
| Stdlib | 3433050 | 78.76 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 953556 | 1811.33 MB/s | 765601 | 2799 | 13.9× |
| LightningDestructive | 954196 | 1810.11 MB/s | 765560 | 2798 | 13.9× |
| LightningArena | 962104 | 1795.24 MB/s | 772704 | 2445 | 13.8× |
| Sonic | 2072083 | 833.56 MB/s | 2728145 | 4020 | 6.4× |
| SonicFastest | 2087005 | 827.60 MB/s | 2724651 | 4020 | 6.3× |
| Goccy | 2455969 | 703.27 MB/s | 2582516 | 14604 | 5.4× |
| Easyjson | 4238841 | 407.47 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4263520 | 405.11 MB/s | 1011637 | 7594 | 3.1× |
| LightningDecodeAny | 4489514 | 111.44 MB/s | 4953695 | 76576 | 2.9× |
| Stdlib | 13233814 | 130.51 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 850 | 2130.67 MB/s | 0 | 0 | 16.4× |
| LightningArena | 855 | 2120.12 MB/s | 0 | 0 | 16.3× |
| LightningDestructive | 868 | 2088.84 MB/s | 0 | 0 | 16.0× |
| Easyjson | 2525 | 717.59 MB/s | 24 | 1 | 5.5× |
| Goccy | 2816 | 643.37 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6027 | 300.64 MB/s | 3741 | 40 | 2.3× |
| Sonic | 6030 | 300.49 MB/s | 3763 | 40 | 2.3× |
| JSONV2 | 7726 | 234.54 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7954 | 227.68 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13916 | 130.21 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 875 | 2071.76 MB/s | 0 | 0 | 16.0× |
| LightningArena | 876 | 2069.60 MB/s | 0 | 0 | 16.0× |
| LightningDestructive | 905 | 2002.02 MB/s | 0 | 0 | 15.4× |
| Easyjson | 2528 | 716.75 MB/s | 24 | 1 | 5.5× |
| Goccy | 2850 | 635.80 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6003 | 301.87 MB/s | 3732 | 40 | 2.3× |
| Sonic | 6005 | 301.75 MB/s | 3751 | 40 | 2.3× |
| JSONV2 | 7850 | 230.82 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7929 | 228.40 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13977 | 129.65 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1061 | 1708.28 MB/s | 144 | 10 | 13.2× |
| LightningArena | 1078 | 1680.26 MB/s | 144 | 10 | 13.0× |
| LightningDestructive | 1119 | 1619.84 MB/s | 144 | 10 | 12.5× |
| Easyjson | 2757 | 657.32 MB/s | 144 | 10 | 5.1× |
| Goccy | 2873 | 630.66 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6173 | 293.55 MB/s | 3739 | 42 | 2.3× |
| Sonic | 6186 | 292.91 MB/s | 3759 | 42 | 2.3× |
| LightningDecodeAny | 7864 | 230.30 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7970 | 227.34 MB/s | 632 | 7 | 1.8× |
| Stdlib | 13990 | 129.52 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 618 | 799.03 MB/s | 160 | 1 | 8.9× |
| LightningDestructive | 619 | 798.47 MB/s | 160 | 1 | 8.9× |
| Sonic | 1260 | 391.94 MB/s | 983 | 6 | 4.4× |
| SonicFastest | 1260 | 391.95 MB/s | 990 | 6 | 4.4× |
| LightningDecodeAny | 1299 | 379.51 MB/s | 1296 | 26 | 4.3× |
| LightningArena | 1387 | 356.05 MB/s | 4120 | 2 | 4.0× |
| Easyjson | 2210 | 223.49 MB/s | 448 | 3 | 2.5× |
| Goccy | 2457 | 201.08 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3270 | 151.08 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5529 | 89.35 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 367 | 626.47 MB/s | 160 | 1 | 11.2× |
| Lightning | 369 | 622.56 MB/s | 160 | 1 | 11.2× |
| Sonic | 896 | 256.84 MB/s | 646 | 6 | 4.6× |
| SonicFastest | 900 | 255.45 MB/s | 655 | 6 | 4.6× |
| LightningArena | 1125 | 204.41 MB/s | 4120 | 2 | 3.7× |
| LightningDecodeAny | 1132 | 202.26 MB/s | 1296 | 26 | 3.6× |
| Easyjson | 1390 | 165.51 MB/s | 448 | 3 | 3.0× |
| Goccy | 1604 | 143.41 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2435 | 94.47 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4127 | 55.73 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 51188 | 1272.40 MB/s | 97220 | 98 | 10.7× |
| LightningArena | 51922 | 1254.42 MB/s | 103440 | 103 | 10.6× |
| Lightning | 52600 | 1238.25 MB/s | 103440 | 103 | 10.5× |
| SonicFastest | 97467 | 668.25 MB/s | 155581 | 75 | 5.6× |
| Sonic | 97492 | 668.07 MB/s | 155600 | 75 | 5.6× |
| Goccy | 145495 | 447.66 MB/s | 229265 | 134 | 3.8× |
| LightningDecodeAny | 180437 | 295.55 MB/s | 180048 | 3245 | 3.0× |
| JSONV2 | 226214 | 287.92 MB/s | 206652 | 607 | 2.4× |
| Stdlib | 549716 | 118.48 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2055641 | 943.97 MB/s | 2864593 | 1380 | 11.2× |
| Lightning | 2109660 | 919.80 MB/s | 2864595 | 1380 | 10.9× |
| LightningArena | 2126783 | 912.40 MB/s | 2864594 | 1380 | 10.9× |
| SonicFastest | 4556163 | 425.90 MB/s | 14606972 | 1407 | 5.1× |
| Goccy | 4800904 | 404.19 MB/s | 4065292 | 13510 | 4.8× |
| Sonic | 4981861 | 389.51 MB/s | 14608622 | 1407 | 4.6× |
| Easyjson | 7618129 | 254.72 MB/s | 3871265 | 15043 | 3.0× |
| LightningDecodeAny | 9591355 | 202.31 MB/s | 7063040 | 218633 | 2.4× |
| JSONV2 | 11296898 | 171.77 MB/s | 3237226 | 13947 | 2.0× |
| Stdlib | 23089049 | 84.04 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 905137 | 3676.61 MB/s | 351704 | 1286 | 23.1× |
| Lightning | 1521359 | 2187.41 MB/s | 2488905 | 2995 | 13.7× |
| LightningArena | 1525530 | 2181.43 MB/s | 2488904 | 2995 | 13.7× |
| Sonic | 2643511 | 1258.87 MB/s | 6479447 | 4248 | 7.9× |
| SonicFastest | 2671072 | 1245.88 MB/s | 6502720 | 4248 | 7.8× |
| LightningDecodeAny | 3570016 | 861.00 MB/s | 4876911 | 56892 | 5.8× |
| Goccy | 4541090 | 732.83 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7320638 | 454.58 MB/s | 5364516 | 13243 | 2.9× |
| Stdlib | 20869154 | 159.46 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 180534 | 1220.52 MB/s | 135872 | 226 | 11.2× |
| LightningDestructive | 180984 | 1217.49 MB/s | 135872 | 226 | 11.1× |
| Lightning | 181720 | 1212.56 MB/s | 135872 | 226 | 11.1× |
| SonicFastest | 380927 | 578.45 MB/s | 309105 | 398 | 5.3× |
| Sonic | 385162 | 572.09 MB/s | 316860 | 398 | 5.2× |
| Goccy | 433486 | 508.31 MB/s | 364974 | 1067 | 4.7× |
| Easyjson | 549816 | 400.76 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 738578 | 298.34 MB/s | 129742 | 470 | 2.7× |
| LightningDecodeAny | 874593 | 123.84 MB/s | 897218 | 11703 | 2.3× |
| Stdlib | 2017034 | 109.24 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9395129 | 862.15 MB/s | 11845072 | 20816 | 9.3× |
| Lightning | 9562437 | 847.07 MB/s | 11845077 | 20816 | 9.2× |
| LightningArena | 9581521 | 845.38 MB/s | 11845077 | 20816 | 9.2× |
| SonicFastest | 16901795 | 479.24 MB/s | 70887663 | 40014 | 5.2× |
| Sonic | 17112833 | 473.33 MB/s | 70873029 | 40014 | 5.1× |
| Goccy | 23615777 | 342.99 MB/s | 17046584 | 107148 | 3.7× |
| Easyjson | 30569915 | 264.97 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 35385554 | 147.04 MB/s | 46279353 | 747112 | 2.5× |
| JSONV2 | 44234918 | 183.11 MB/s | 15233724 | 78972 | 2.0× |
| Stdlib | 87765970 | 92.29 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4300051 | 693.82 MB/s | 3780456 | 1514 | 10.8× |
| LightningDestructive | 4560112 | 654.25 MB/s | 3758856 | 29356 | 10.2× |
| Lightning | 4658774 | 640.40 MB/s | 3758859 | 29356 | 10.0× |
| Sonic | 8699081 | 342.96 MB/s | 26542194 | 56760 | 5.4× |
| SonicFastest | 8714666 | 342.35 MB/s | 26580116 | 56760 | 5.3× |
| LightningDecodeAny | 16223028 | 113.06 MB/s | 23982579 | 351152 | 2.9× |
| Goccy | 16713090 | 178.51 MB/s | 10590461 | 273648 | 2.8× |
| Easyjson | 16887634 | 176.67 MB/s | 9479440 | 30115 | 2.8× |
| JSONV2 | 24785697 | 120.37 MB/s | 9257160 | 86278 | 1.9× |
| Stdlib | 46557104 | 64.08 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 920693 | 785.93 MB/s | 907601 | 3618 | 12.4× |
| LightningArena | 935775 | 773.26 MB/s | 916259 | 37 | 12.2× |
| Lightning | 975406 | 741.84 MB/s | 907599 | 3618 | 11.7× |
| Sonic | 1790580 | 404.11 MB/s | 3187798 | 7226 | 6.4× |
| SonicFastest | 1792577 | 403.66 MB/s | 3191735 | 7226 | 6.4× |
| LightningDecodeAny | 4072641 | 159.74 MB/s | 6500455 | 76546 | 2.8× |
| Easyjson | 4207246 | 171.99 MB/s | 2847904 | 3698 | 2.7× |
| Goccy | 4832847 | 149.72 MB/s | 2805446 | 80273 | 2.4× |
| JSONV2 | 5474644 | 132.17 MB/s | 2704655 | 7318 | 2.1× |
| Stdlib | 11395704 | 63.50 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1365977 | 1154.74 MB/s | 907601 | 3618 | 11.4× |
| LightningArena | 1368689 | 1152.46 MB/s | 916257 | 37 | 11.3× |
| Lightning | 1404008 | 1123.46 MB/s | 907596 | 3618 | 11.0× |
| Sonic | 2236980 | 705.13 MB/s | 5783814 | 7226 | 6.9× |
| SonicFastest | 2241867 | 703.59 MB/s | 5794967 | 7226 | 6.9× |
| LightningDecodeAny | 3626077 | 207.77 MB/s | 6500458 | 76546 | 4.3× |
| Easyjson | 5564055 | 283.49 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5659226 | 278.72 MB/s | 3558934 | 80266 | 2.7× |
| JSONV2 | 6162887 | 255.94 MB/s | 2704590 | 7318 | 2.5× |
| Stdlib | 15514037 | 101.67 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 157099 | 955.60 MB/s | 81920 | 1 | 11.8× |
| LightningArena | 157179 | 955.12 MB/s | 81920 | 1 | 11.8× |
| Lightning | 157252 | 954.67 MB/s | 81920 | 1 | 11.8× |
| Sonic | 272923 | 550.06 MB/s | 247425 | 6 | 6.8× |
| SonicFastest | 273629 | 548.64 MB/s | 251520 | 6 | 6.8× |
| LightningDecodeAny | 429376 | 349.63 MB/s | 745764 | 10016 | 4.3× |
| Goccy | 866847 | 173.18 MB/s | 324459 | 10004 | 2.1× |
| JSONV2 | 1097489 | 136.79 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1857537 | 80.82 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 27083 | 1038.17 MB/s | 29216 | 103 | 11.1× |
| LightningDestructive | 27299 | 1029.98 MB/s | 29088 | 101 | 11.1× |
| Lightning | 27372 | 1027.22 MB/s | 29216 | 103 | 11.0× |
| Sonic | 63440 | 443.21 MB/s | 47042 | 103 | 4.8× |
| SonicFastest | 63499 | 442.79 MB/s | 46945 | 103 | 4.8× |
| Easyjson | 68339 | 411.43 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72067 | 390.15 MB/s | 59221 | 188 | 4.2× |
| JSONV2 | 135108 | 208.11 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 147310 | 190.87 MB/s | 140576 | 2643 | 2.0× |
| Stdlib | 301743 | 93.18 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1509 | 1542.39 MB/s | 32 | 1 | 15.0× |
| LightningArena | 1512 | 1539.61 MB/s | 32 | 1 | 14.9× |
| LightningDestructive | 1589 | 1465.25 MB/s | 32 | 1 | 14.2× |
| Easyjson | 4223 | 551.24 MB/s | 192 | 2 | 5.3× |
| Goccy | 4227 | 550.72 MB/s | 3649 | 4 | 5.3× |
| Sonic | 5183 | 449.16 MB/s | 4537 | 6 | 4.4× |
| SonicFastest | 5225 | 445.55 MB/s | 4511 | 6 | 4.3× |
| JSONV2 | 8521 | 273.22 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10499 | 160.49 MB/s | 10200 | 195 | 2.2× |
| Stdlib | 22592 | 103.05 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 179 | 1058.21 MB/s | 0 | 0 | 13.7× |
| Lightning | 180 | 1052.97 MB/s | 0 | 0 | 13.6× |
| LightningDestructive | 182 | 1036.09 MB/s | 0 | 0 | 13.4× |
| Goccy | 393 | 480.94 MB/s | 304 | 2 | 6.2× |
| Easyjson | 493 | 383.32 MB/s | 0 | 0 | 4.9× |
| SonicFastest | 814 | 232.10 MB/s | 496 | 4 | 3.0× |
| Sonic | 816 | 231.77 MB/s | 498 | 4 | 3.0× |
| JSONV2 | 1044 | 181.11 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1204 | 111.26 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2439 | 77.49 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1143 | 1916.12 MB/s | 0 | 0 | 13.8× |
| Lightning | 1148 | 1907.87 MB/s | 0 | 0 | 13.8× |
| LightningDestructive | 1172 | 1869.15 MB/s | 0 | 0 | 13.5× |
| Easyjson | 3196 | 685.61 MB/s | 24 | 1 | 4.9× |
| Goccy | 3201 | 684.56 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6417 | 341.44 MB/s | 3961 | 40 | 2.5× |
| Sonic | 6435 | 340.51 MB/s | 3993 | 40 | 2.5× |
| LightningDecodeAny | 7910 | 228.96 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 8128 | 269.56 MB/s | 640 | 6 | 1.9× |
| Stdlib | 15818 | 138.51 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 556414 | 917.44 MB/s | 457536 | 1009 | 10.8× |
| LightningArena | 557198 | 916.15 MB/s | 457536 | 1009 | 10.8× |
| Lightning | 557267 | 916.03 MB/s | 457537 | 1009 | 10.8× |
| Sonic | 1173866 | 434.87 MB/s | 906298 | 2006 | 5.1× |
| SonicFastest | 1175291 | 434.34 MB/s | 906507 | 2006 | 5.1× |
| Goccy | 1182783 | 431.59 MB/s | 1139132 | 5006 | 5.1× |
| Easyjson | 1559742 | 327.28 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3239196 | 157.59 MB/s | 1076010 | 12646 | 1.9× |
| LightningDecodeAny | 3428325 | 134.60 MB/s | 2950650 | 64018 | 1.8× |
| Stdlib | 6024596 | 84.73 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 481 | 41103.52 MB/s | 0 | 0 | 225.5× |
| Lightning | 482 | 41090.79 MB/s | 0 | 0 | 225.4× |
| LightningDestructive | 508 | 38981.46 MB/s | 0 | 0 | 213.8× |
| Goccy | 20011 | 988.92 MB/s | 20491 | 2 | 5.4× |
| Sonic | 27539 | 718.58 MB/s | 22897 | 4 | 3.9× |
| SonicFastest | 27634 | 716.10 MB/s | 22964 | 4 | 3.9× |
| JSONV2 | 29764 | 664.86 MB/s | 8 | 1 | 3.6× |
| Easyjson | 81722 | 242.15 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 82986 | 238.45 MB/s | 116864 | 2015 | 1.3× |
| Stdlib | 108565 | 182.28 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1757 | 10316.17 MB/s | 0 | 0 | 58.9× |
| Lightning | 1896 | 9558.71 MB/s | 432 | 2 | 54.6× |
| LightningArena | 1898 | 9548.31 MB/s | 432 | 2 | 54.5× |
| Easyjson | 3992 | 4539.74 MB/s | 432 | 2 | 25.9× |
| SonicFastest | 10318 | 1756.48 MB/s | 24488 | 6 | 10.0× |
| Sonic | 10535 | 1720.31 MB/s | 24639 | 6 | 9.8× |
| Goccy | 16121 | 1124.24 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 17195 | 1039.93 MB/s | 29088 | 191 | 6.0× |
| JSONV2 | 46251 | 391.86 MB/s | 16501 | 50 | 2.2× |
| Stdlib | 103446 | 175.20 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2110179 | 951.81 MB/s | 3089564 | 6821 | 8.8× |
| Lightning | 2181490 | 920.70 MB/s | 3091277 | 6827 | 8.5× |
| LightningArena | 2200567 | 912.72 MB/s | 3094394 | 6704 | 8.4× |
| Goccy | 4311454 | 465.85 MB/s | 5411963 | 15830 | 4.3× |
| SonicFastest | 4442053 | 452.15 MB/s | 10980202 | 13683 | 4.2× |
| Sonic | 4444666 | 451.89 MB/s | 10971341 | 13683 | 4.2× |
| Easyjson | 4927602 | 407.60 MB/s | 2981483 | 7439 | 3.8× |
| JSONV2 | 7015589 | 286.29 MB/s | 3173685 | 14563 | 2.6× |
| LightningDecodeAny | 7625148 | 149.81 MB/s | 8503511 | 134008 | 2.4× |
| Stdlib | 18522150 | 108.44 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 852 | 644.18 MB/s | 480 | 1 | 6.6× |
| Lightning | 854 | 642.78 MB/s | 480 | 1 | 6.5× |
| LightningDestructive | 857 | 640.73 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1676 | 326.94 MB/s | 2021 | 46 | 3.3× |
| Easyjson | 2160 | 254.22 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2680 | 204.85 MB/s | 1934 | 26 | 2.1× |
| SonicFastest | 2691 | 203.99 MB/s | 1934 | 26 | 2.1× |
| Goccy | 2998 | 183.14 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3335 | 164.61 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5585 | 98.30 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 416150 | 1517.51 MB/s | 402729 | 545 | 13.0× |
| LightningArena | 484847 | 1302.50 MB/s | 453041 | 713 | 11.1× |
| Lightning | 489926 | 1289.00 MB/s | 451257 | 857 | 11.0× |
| Sonic | 1016420 | 621.31 MB/s | 994154 | 1102 | 5.3× |
| SonicFastest | 1017687 | 620.54 MB/s | 994154 | 1102 | 5.3× |
| Easyjson | 1148638 | 549.79 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1164729 | 542.20 MB/s | 985690 | 1201 | 4.6× |
| JSONV2 | 2159739 | 292.40 MB/s | 571618 | 3144 | 2.5× |
| LightningDecodeAny | 2411728 | 193.60 MB/s | 2076506 | 30126 | 2.2× |
| Stdlib | 5392581 | 117.11 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 541874 | 1037.89 MB/s | 546571 | 429 | 9.7× |
| Lightning | 708915 | 793.34 MB/s | 769938 | 1235 | 7.4× |
| LightningArena | 717142 | 784.24 MB/s | 771690 | 1089 | 7.3× |
| Sonic | 1034213 | 543.80 MB/s | 959585 | 1476 | 5.1× |
| SonicFastest | 1039256 | 541.16 MB/s | 965461 | 1476 | 5.1× |
| Goccy | 1352269 | 415.90 MB/s | 1042982 | 1030 | 3.9× |
| Easyjson | 1763993 | 318.83 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2626669 | 214.11 MB/s | 2180439 | 30126 | 2.0× |
| JSONV2 | 2793973 | 201.29 MB/s | 927439 | 3482 | 1.9× |
| Stdlib | 5266204 | 106.80 MB/s | 1011671 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 578723 | 921.30 MB/s | 333416 | 2084 | 9.4× |
| Lightning | 600720 | 887.56 MB/s | 368224 | 2293 | 9.1× |
| LightningArena | 607351 | 877.87 MB/s | 368224 | 2293 | 9.0× |
| Easyjson | 1116613 | 477.50 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1149867 | 463.69 MB/s | 1045584 | 4351 | 4.7× |
| SonicFastest | 1152567 | 462.60 MB/s | 1040128 | 4351 | 4.7× |
| Goccy | 1313595 | 405.89 MB/s | 1167230 | 5409 | 4.1× |
| JSONV2 | 2542207 | 209.73 MB/s | 745445 | 13288 | 2.1× |
| LightningDecodeAny | 3388539 | 157.35 MB/s | 2992878 | 50076 | 1.6× |
| Stdlib | 5448386 | 97.86 MB/s | 798692 | 17133 | 1.0× |
