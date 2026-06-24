# JSON Deserialization Benchmarks

- generated 2026-06-24T06:16:58Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104408 | 1219.01 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 104726 | 1215.32 MB/s | 49280 | 2 | 10.5× |
| SonicFastest | 176966 | 719.20 MB/s | 190366 | 10 | 6.2× |
| Sonic | 178584 | 712.69 MB/s | 193288 | 10 | 6.2× |
| Goccy | 182879 | 695.95 MB/s | 224126 | 883 | 6.0× |
| Easyjson | 210180 | 605.55 MB/s | 122864 | 14 | 5.3× |
| JSONV2 | 414127 | 307.33 MB/s | 195118 | 1805 | 2.7× |
| LightningDecodeAny | 451556 | 209.62 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1104255 | 115.26 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3716064 | 605.76 MB/s | 3122872 | 3081 | 7.0× |
| Lightning | 3805936 | 591.46 MB/s | 3122873 | 3081 | 6.8× |
| SonicFastest | 4482866 | 502.15 MB/s | 15238674 | 970 | 5.8× |
| Sonic | 4490541 | 501.29 MB/s | 15238679 | 970 | 5.8× |
| Goccy | 10563691 | 213.09 MB/s | 4127044 | 56532 | 2.5× |
| Easyjson | 11064613 | 203.45 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11883638 | 189.42 MB/s | 7938298 | 281383 | 2.2× |
| JSONV2 | 16211978 | 138.85 MB/s | 3123206 | 3083 | 1.6× |
| Stdlib | 26049562 | 86.41 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 485282 | 557.21 MB/s | 348024 | 1627 | 6.9× |
| LightningDestructive | 489103 | 552.85 MB/s | 348024 | 1627 | 6.8× |
| SonicFastest | 615630 | 439.23 MB/s | 459954 | 968 | 5.4× |
| Sonic | 615668 | 439.20 MB/s | 463521 | 968 | 5.4× |
| Goccy | 1402189 | 192.84 MB/s | 543163 | 8122 | 2.4× |
| Easyjson | 1439148 | 187.89 MB/s | 330272 | 749 | 2.3× |
| LightningDecodeAny | 1597506 | 169.27 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2104112 | 128.51 MB/s | 348148 | 1628 | 1.6× |
| Stdlib | 3346981 | 80.79 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1340022 | 1288.94 MB/s | 959848 | 5881 | 9.9× |
| Lightning | 1362574 | 1267.60 MB/s | 959890 | 5882 | 9.8× |
| SonicFastest | 2003488 | 862.10 MB/s | 2730281 | 4020 | 6.6× |
| Sonic | 2005808 | 861.10 MB/s | 2761054 | 4020 | 6.6× |
| Goccy | 2317623 | 745.25 MB/s | 2582607 | 14604 | 5.7× |
| Easyjson | 4216032 | 409.68 MB/s | 972032 | 5389 | 3.2× |
| JSONV2 | 4286939 | 402.90 MB/s | 1011637 | 7594 | 3.1× |
| LightningDecodeAny | 4636616 | 107.90 MB/s | 4976205 | 81466 | 2.9× |
| Stdlib | 13308134 | 129.79 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1189 | 1524.32 MB/s | 0 | 0 | 11.9× |
| LightningDestructive | 1204 | 1504.89 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2550 | 710.64 MB/s | 24 | 1 | 5.5× |
| Goccy | 2779 | 651.96 MB/s | 2608 | 4 | 5.1× |
| Sonic | 5894 | 307.41 MB/s | 3679 | 40 | 2.4× |
| SonicFastest | 5895 | 307.40 MB/s | 3696 | 40 | 2.4× |
| JSONV2 | 7826 | 231.54 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8114 | 223.20 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14093 | 128.57 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1220 | 1485.14 MB/s | 0 | 0 | 11.6× |
| LightningDestructive | 1233 | 1469.79 MB/s | 0 | 0 | 11.5× |
| Easyjson | 2544 | 712.19 MB/s | 24 | 1 | 5.6× |
| Goccy | 2791 | 649.14 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5829 | 310.87 MB/s | 3708 | 40 | 2.4× |
| Sonic | 5845 | 310.00 MB/s | 3738 | 40 | 2.4× |
| JSONV2 | 7824 | 231.61 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8114 | 223.21 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14141 | 128.13 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1399 | 1295.05 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1444 | 1255.12 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2761 | 656.32 MB/s | 144 | 10 | 5.1× |
| Goccy | 2885 | 628.17 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6044 | 299.80 MB/s | 3733 | 42 | 2.3× |
| SonicFastest | 6058 | 299.09 MB/s | 3761 | 42 | 2.3× |
| JSONV2 | 7993 | 226.69 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8183 | 221.32 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14030 | 129.15 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 739 | 668.70 MB/s | 160 | 1 | 7.5× |
| LightningDestructive | 744 | 664.36 MB/s | 160 | 1 | 7.5× |
| SonicFastest | 1225 | 403.39 MB/s | 976 | 6 | 4.5× |
| Sonic | 1228 | 402.19 MB/s | 978 | 6 | 4.5× |
| LightningDecodeAny | 1639 | 300.86 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2207 | 223.83 MB/s | 448 | 3 | 2.5× |
| Goccy | 2471 | 199.95 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3198 | 154.48 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5572 | 88.65 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 470 | 489.88 MB/s | 160 | 1 | 8.9× |
| LightningDestructive | 474 | 485.55 MB/s | 160 | 1 | 8.8× |
| Sonic | 880 | 261.30 MB/s | 652 | 6 | 4.7× |
| SonicFastest | 881 | 261.08 MB/s | 655 | 6 | 4.7× |
| LightningDecodeAny | 1380 | 165.96 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1386 | 165.94 MB/s | 448 | 3 | 3.0× |
| Goccy | 1622 | 141.81 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2342 | 98.22 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4169 | 55.17 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 70594 | 922.63 MB/s | 166212 | 102 | 7.7× |
| Lightning | 72176 | 902.40 MB/s | 172432 | 107 | 7.5× |
| Sonic | 93932 | 693.40 MB/s | 154443 | 75 | 5.8× |
| SonicFastest | 94656 | 688.09 MB/s | 155236 | 75 | 5.7× |
| Goccy | 142303 | 457.70 MB/s | 228751 | 134 | 3.8× |
| LightningDecodeAny | 184618 | 288.86 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 222430 | 292.82 MB/s | 206652 | 607 | 2.4× |
| Stdlib | 542420 | 120.08 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2583127 | 751.21 MB/s | 2846865 | 2238 | 9.1× |
| Lightning | 2632079 | 737.24 MB/s | 2846866 | 2238 | 8.9× |
| Sonic | 4626818 | 419.40 MB/s | 14608611 | 1407 | 5.1× |
| SonicFastest | 4650564 | 417.26 MB/s | 14608593 | 1407 | 5.0× |
| Goccy | 4781579 | 405.82 MB/s | 4064326 | 13510 | 4.9× |
| Easyjson | 7525800 | 257.84 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9860209 | 196.80 MB/s | 7013523 | 219937 | 2.4× |
| JSONV2 | 11257249 | 172.38 MB/s | 3237222 | 13947 | 2.1× |
| Stdlib | 23447851 | 82.76 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1109760 | 2998.69 MB/s | 351704 | 1286 | 18.7× |
| Lightning | 1747904 | 1903.90 MB/s | 2488904 | 2995 | 11.9× |
| SonicFastest | 2628349 | 1266.13 MB/s | 6426190 | 4248 | 7.9× |
| Sonic | 2644654 | 1258.32 MB/s | 6453736 | 4248 | 7.8× |
| LightningDecodeAny | 3719968 | 826.29 MB/s | 4886621 | 56892 | 5.6× |
| Goccy | 4452476 | 747.41 MB/s | 3948908 | 3816 | 4.7× |
| JSONV2 | 7303026 | 455.68 MB/s | 5364516 | 13243 | 2.8× |
| Stdlib | 20712979 | 160.66 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220265 | 1000.37 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 222346 | 991.01 MB/s | 136897 | 228 | 9.2× |
| SonicFastest | 380896 | 578.49 MB/s | 318992 | 398 | 5.3× |
| Sonic | 391612 | 562.66 MB/s | 347183 | 398 | 5.2× |
| Goccy | 428204 | 514.58 MB/s | 364286 | 1067 | 4.8× |
| Easyjson | 543018 | 405.78 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 733194 | 300.53 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 868975 | 124.64 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2034535 | 108.30 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12718825 | 636.85 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13165271 | 615.26 MB/s | 15059836 | 51649 | 6.8× |
| Sonic | 16744395 | 483.75 MB/s | 70887434 | 40014 | 5.3× |
| SonicFastest | 16754966 | 483.44 MB/s | 70887394 | 40014 | 5.3× |
| Goccy | 23552971 | 343.91 MB/s | 16752843 | 107147 | 3.8× |
| Easyjson | 31065740 | 260.74 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 42534480 | 122.32 MB/s | 34333354 | 912810 | 2.1× |
| JSONV2 | 43360717 | 186.81 MB/s | 15233737 | 78972 | 2.1× |
| Stdlib | 89236877 | 90.77 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6024422 | 495.23 MB/s | 3985336 | 30015 | 7.8× |
| Lightning | 6131084 | 486.61 MB/s | 3985339 | 30015 | 7.7× |
| Sonic | 8568090 | 348.21 MB/s | 26534556 | 56760 | 5.5× |
| SonicFastest | 8576859 | 347.85 MB/s | 26513726 | 56760 | 5.5× |
| Easyjson | 16718483 | 178.45 MB/s | 9479440 | 30115 | 2.8× |
| Goccy | 16728901 | 178.34 MB/s | 10769823 | 273655 | 2.8× |
| LightningDecodeAny | 19310722 | 94.98 MB/s | 20023840 | 408557 | 2.4× |
| JSONV2 | 24282979 | 122.86 MB/s | 9257160 | 86278 | 1.9× |
| Stdlib | 46971743 | 63.52 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1366903 | 529.37 MB/s | 907389 | 3618 | 8.4× |
| LightningDestructive | 1372588 | 527.18 MB/s | 907393 | 3618 | 8.4× |
| Sonic | 1744293 | 414.84 MB/s | 3170228 | 7226 | 6.6× |
| SonicFastest | 1759406 | 411.27 MB/s | 3187186 | 7226 | 6.6× |
| Easyjson | 4248621 | 170.31 MB/s | 2847904 | 3698 | 2.7× |
| LightningDecodeAny | 4276525 | 152.13 MB/s | 5752201 | 80172 | 2.7× |
| Goccy | 4697122 | 154.05 MB/s | 2730817 | 80269 | 2.5× |
| JSONV2 | 5986903 | 120.86 MB/s | 2704625 | 7318 | 1.9× |
| Stdlib | 11545990 | 62.67 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1931635 | 816.59 MB/s | 907393 | 3618 | 8.1× |
| Lightning | 1942160 | 812.16 MB/s | 907386 | 3618 | 8.1× |
| Sonic | 2187952 | 720.93 MB/s | 5783759 | 7226 | 7.2× |
| SonicFastest | 2211324 | 713.31 MB/s | 5787234 | 7226 | 7.1× |
| LightningDecodeAny | 3807996 | 197.85 MB/s | 5752202 | 80172 | 4.1× |
| Easyjson | 5555038 | 283.95 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5689218 | 277.25 MB/s | 3591936 | 80267 | 2.8× |
| JSONV2 | 6703857 | 235.29 MB/s | 2704590 | 7318 | 2.3× |
| Stdlib | 15701361 | 100.46 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222987 | 673.24 MB/s | 81920 | 1 | 8.1× |
| LightningDestructive | 224100 | 669.90 MB/s | 81920 | 1 | 8.1× |
| SonicFastest | 266582 | 563.14 MB/s | 246221 | 6 | 6.8× |
| Sonic | 269811 | 556.40 MB/s | 259517 | 6 | 6.7× |
| LightningDecodeAny | 466974 | 321.48 MB/s | 746005 | 10020 | 3.9× |
| Goccy | 852926 | 176.01 MB/s | 325588 | 10005 | 2.1× |
| JSONV2 | 1100561 | 136.41 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1816984 | 82.62 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33008 | 851.82 MB/s | 30272 | 105 | 9.2× |
| LightningDestructive | 33386 | 842.17 MB/s | 30144 | 103 | 9.1× |
| Sonic | 63847 | 440.38 MB/s | 47975 | 103 | 4.8× |
| SonicFastest | 64107 | 438.59 MB/s | 47984 | 103 | 4.8× |
| Easyjson | 68013 | 413.41 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71970 | 390.68 MB/s | 59260 | 188 | 4.2× |
| JSONV2 | 134926 | 208.39 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 155964 | 180.28 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 305154 | 92.14 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1950 | 1194.02 MB/s | 32 | 1 | 11.7× |
| LightningDestructive | 2008 | 1159.47 MB/s | 32 | 1 | 11.4× |
| Goccy | 4092 | 568.90 MB/s | 3649 | 4 | 5.6× |
| Easyjson | 4255 | 547.12 MB/s | 192 | 2 | 5.4× |
| Sonic | 4996 | 465.97 MB/s | 4175 | 6 | 4.6× |
| SonicFastest | 5056 | 460.45 MB/s | 4305 | 6 | 4.5× |
| JSONV2 | 8404 | 277.00 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10192 | 165.32 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22830 | 101.97 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 862.20 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 221 | 855.66 MB/s | 0 | 0 | 11.4× |
| Goccy | 376 | 503.30 MB/s | 304 | 2 | 6.7× |
| Easyjson | 482 | 392.05 MB/s | 0 | 0 | 5.2× |
| SonicFastest | 798 | 236.99 MB/s | 491 | 4 | 3.2× |
| Sonic | 799 | 236.62 MB/s | 494 | 4 | 3.2× |
| JSONV2 | 1039 | 181.96 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1222 | 109.70 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2526 | 74.81 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1543 | 1420.24 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1559 | 1405.35 MB/s | 0 | 0 | 10.3× |
| Goccy | 3155 | 694.41 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3181 | 688.79 MB/s | 24 | 1 | 5.0× |
| Sonic | 6295 | 348.05 MB/s | 3913 | 40 | 2.6× |
| SonicFastest | 6329 | 346.17 MB/s | 3970 | 40 | 2.5× |
| JSONV2 | 8087 | 270.93 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8124 | 222.92 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 16055 | 136.47 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 689128 | 740.76 MB/s | 703778 | 1012 | 8.7× |
| Lightning | 719929 | 709.06 MB/s | 703780 | 1012 | 8.4× |
| Goccy | 1137805 | 448.65 MB/s | 1139800 | 5006 | 5.3× |
| SonicFastest | 1163841 | 438.61 MB/s | 929802 | 2006 | 5.2× |
| Sonic | 1165434 | 438.01 MB/s | 937148 | 2006 | 5.2× |
| Easyjson | 1530087 | 333.63 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3195725 | 159.74 MB/s | 1076012 | 12646 | 1.9× |
| LightningDecodeAny | 3510422 | 131.46 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6025152 | 84.72 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14645.14 MB/s | 0 | 0 | 82.6× |
| LightningDestructive | 1388 | 14255.14 MB/s | 0 | 0 | 80.4× |
| Goccy | 19751 | 1001.93 MB/s | 20491 | 2 | 5.6× |
| SonicFastest | 27272 | 725.60 MB/s | 22814 | 4 | 4.1× |
| Sonic | 27274 | 725.55 MB/s | 22749 | 4 | 4.1× |
| JSONV2 | 29710 | 666.07 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 79934 | 247.56 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86022 | 230.05 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111564 | 177.38 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2660 | 6812.61 MB/s | 0 | 0 | 38.8× |
| Lightning | 2835 | 6393.44 MB/s | 432 | 2 | 36.4× |
| Easyjson | 3976 | 4558.18 MB/s | 432 | 2 | 26.0× |
| Sonic | 9769 | 1855.24 MB/s | 22832 | 6 | 10.6× |
| SonicFastest | 9938 | 1823.66 MB/s | 23131 | 6 | 10.4× |
| Goccy | 15450 | 1173.05 MB/s | 19459 | 2 | 6.7× |
| LightningDecodeAny | 16524 | 1082.19 MB/s | 29088 | 191 | 6.3× |
| JSONV2 | 45561 | 397.80 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103278 | 175.49 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2188320 | 917.82 MB/s | 2960389 | 7411 | 8.5× |
| Lightning | 2245872 | 894.30 MB/s | 2962102 | 7417 | 8.3× |
| Goccy | 4104872 | 489.30 MB/s | 5412450 | 15832 | 4.5× |
| Sonic | 4283152 | 468.93 MB/s | 11003971 | 13683 | 4.3× |
| SonicFastest | 4287110 | 468.50 MB/s | 10963333 | 13683 | 4.3× |
| Easyjson | 4868392 | 412.56 MB/s | 2981482 | 7438 | 3.8× |
| JSONV2 | 6781724 | 296.16 MB/s | 3173687 | 14563 | 2.7× |
| LightningDecodeAny | 7369569 | 155.00 MB/s | 7386072 | 134751 | 2.5× |
| Stdlib | 18578910 | 108.11 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 886 | 619.63 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 894 | 614.26 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1904 | 287.86 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2142 | 256.33 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2673 | 205.39 MB/s | 1957 | 26 | 2.1× |
| SonicFastest | 2684 | 204.51 MB/s | 1979 | 26 | 2.1× |
| Goccy | 3056 | 179.66 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3274 | 167.71 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5596 | 98.11 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 492635 | 1281.91 MB/s | 364472 | 566 | 10.9× |
| Lightning | 543601 | 1161.72 MB/s | 413002 | 878 | 9.9× |
| SonicFastest | 1004580 | 628.63 MB/s | 1009683 | 1102 | 5.4× |
| Sonic | 1008147 | 626.41 MB/s | 1015139 | 1102 | 5.3× |
| Easyjson | 1142492 | 552.75 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1167532 | 540.90 MB/s | 989485 | 1202 | 4.6× |
| JSONV2 | 2153583 | 293.24 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2386087 | 195.68 MB/s | 2010201 | 30295 | 2.3× |
| Stdlib | 5392593 | 117.11 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 710385 | 791.69 MB/s | 544249 | 448 | 7.4× |
| Lightning | 880384 | 638.82 MB/s | 767621 | 1254 | 6.0× |
| SonicFastest | 1004689 | 559.78 MB/s | 940279 | 1476 | 5.2× |
| Sonic | 1007775 | 558.07 MB/s | 952451 | 1476 | 5.2× |
| Goccy | 1333884 | 421.63 MB/s | 1044119 | 1030 | 3.9× |
| Easyjson | 1743139 | 322.64 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2696932 | 208.54 MB/s | 2114152 | 30295 | 2.0× |
| JSONV2 | 2744016 | 204.96 MB/s | 927438 | 3482 | 1.9× |
| Stdlib | 5264541 | 106.83 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 648714 | 821.90 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 668315 | 797.79 MB/s | 368224 | 2293 | 8.2× |
| Easyjson | 1102466 | 483.62 MB/s | 428361 | 3273 | 5.0× |
| SonicFastest | 1132348 | 470.86 MB/s | 1048521 | 4351 | 4.8× |
| Sonic | 1135313 | 469.63 MB/s | 1041178 | 4351 | 4.8× |
| Goccy | 1291002 | 413.00 MB/s | 1167246 | 5409 | 4.2× |
| JSONV2 | 2529930 | 210.75 MB/s | 745455 | 13288 | 2.2× |
| LightningDecodeAny | 3358891 | 158.74 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5472444 | 97.43 MB/s | 798691 | 17133 | 1.0× |
