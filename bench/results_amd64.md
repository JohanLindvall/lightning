# JSON Deserialization Benchmarks

- generated 2026-06-25T21:22:42Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 80005 | 1590.83 MB/s | 49760 | 3 | 12.0× |
| LightningDestructive | 82845 | 1536.31 MB/s | 49280 | 2 | 11.6× |
| Easyjson | 183622 | 693.14 MB/s | 122864 | 14 | 5.2× |
| Goccy | 187720 | 678.01 MB/s | 225338 | 884 | 5.1× |
| SonicFastest | 203504 | 625.42 MB/s | 214468 | 15 | 4.7× |
| Sonic | 204568 | 622.17 MB/s | 214474 | 15 | 4.7× |
| JSONV2 | 302857 | 420.25 MB/s | 195127 | 1805 | 3.2× |
| LightningDecodeAny | 331038 | 285.93 MB/s | 465586 | 9714 | 2.9× |
| Stdlib | 963557 | 132.09 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3188040 | 706.09 MB/s | 3122872 | 3081 | 7.6× |
| Lightning | 3206658 | 701.99 MB/s | 3122874 | 3081 | 7.5× |
| SonicFastest | 3962109 | 568.14 MB/s | 4871186 | 2584 | 6.1× |
| Sonic | 4028323 | 558.81 MB/s | 4869911 | 2584 | 6.0× |
| LightningDecodeAny | 8687425 | 259.11 MB/s | 7938298 | 281383 | 2.8× |
| Goccy | 9517356 | 236.52 MB/s | 4192023 | 56535 | 2.5× |
| Easyjson | 10533699 | 213.70 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 13336120 | 168.79 MB/s | 3123188 | 3083 | 1.8× |
| Stdlib | 24110638 | 93.36 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 432851 | 624.70 MB/s | 348024 | 1627 | 7.2× |
| LightningDestructive | 436946 | 618.85 MB/s | 348024 | 1627 | 7.1× |
| Sonic | 567138 | 476.79 MB/s | 641738 | 1147 | 5.5× |
| SonicFastest | 568053 | 476.02 MB/s | 641463 | 1147 | 5.5× |
| LightningDecodeAny | 1229940 | 219.85 MB/s | 1011486 | 37901 | 2.5× |
| Goccy | 1302890 | 207.54 MB/s | 541659 | 8122 | 2.4× |
| Easyjson | 1342551 | 201.41 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 1745512 | 154.91 MB/s | 348159 | 1628 | 1.8× |
| Stdlib | 3119600 | 86.68 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1004371 | 1719.69 MB/s | 959848 | 5881 | 13.0× |
| Lightning | 1040757 | 1659.57 MB/s | 959890 | 5882 | 12.5× |
| Sonic | 1722135 | 1002.94 MB/s | 2695890 | 5547 | 7.6× |
| SonicFastest | 1727289 | 999.95 MB/s | 2695381 | 5547 | 7.5× |
| Goccy | 1906344 | 906.03 MB/s | 2580859 | 14603 | 6.8× |
| Easyjson | 2793633 | 618.26 MB/s | 972032 | 5389 | 4.7× |
| LightningDecodeAny | 3105989 | 161.08 MB/s | 4976204 | 81466 | 4.2× |
| JSONV2 | 3235786 | 533.78 MB/s | 1011615 | 7594 | 4.0× |
| Stdlib | 13013481 | 132.72 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 754 | 2404.61 MB/s | 0 | 0 | 15.2× |
| LightningDestructive | 773 | 2344.30 MB/s | 0 | 0 | 14.9× |
| Easyjson | 2127 | 851.89 MB/s | 24 | 1 | 5.4× |
| Goccy | 2581 | 702.09 MB/s | 2608 | 4 | 4.5× |
| SonicFastest | 4815 | 376.33 MB/s | 3347 | 38 | 2.4× |
| Sonic | 5000 | 362.41 MB/s | 3345 | 38 | 2.3× |
| JSONV2 | 6042 | 299.89 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 6640 | 272.75 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 11492 | 157.68 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 842 | 2153.01 MB/s | 0 | 0 | 13.5× |
| LightningDestructive | 861 | 2104.65 MB/s | 0 | 0 | 13.2× |
| Easyjson | 2181 | 830.81 MB/s | 24 | 1 | 5.2× |
| Goccy | 2376 | 762.56 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 4726 | 383.45 MB/s | 3347 | 38 | 2.4× |
| Sonic | 4874 | 371.78 MB/s | 3345 | 38 | 2.3× |
| JSONV2 | 5760 | 314.60 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 6592 | 274.73 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 11396 | 159.00 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 968 | 1871.06 MB/s | 144 | 10 | 11.6× |
| LightningDestructive | 997 | 1816.76 MB/s | 144 | 10 | 11.3× |
| Easyjson | 2300 | 787.74 MB/s | 144 | 10 | 4.9× |
| Goccy | 2561 | 707.55 MB/s | 2600 | 5 | 4.4× |
| SonicFastest | 4801 | 377.40 MB/s | 3367 | 40 | 2.3× |
| Sonic | 4947 | 366.28 MB/s | 3367 | 40 | 2.3× |
| JSONV2 | 6241 | 290.33 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 6742 | 268.60 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 11256 | 160.98 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 575 | 859.09 MB/s | 160 | 1 | 8.2× |
| LightningDestructive | 582 | 848.66 MB/s | 160 | 1 | 8.1× |
| Sonic | 948 | 521.01 MB/s | 1075 | 8 | 4.9× |
| SonicFastest | 952 | 519.10 MB/s | 1075 | 8 | 4.9× |
| LightningDecodeAny | 1271 | 387.75 MB/s | 1536 | 30 | 3.7× |
| Easyjson | 1827 | 270.45 MB/s | 448 | 3 | 2.6× |
| Goccy | 1891 | 261.24 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 2399 | 205.90 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4687 | 105.39 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 360 | 638.65 MB/s | 160 | 1 | 9.1× |
| LightningDestructive | 364 | 631.96 MB/s | 160 | 1 | 9.0× |
| Sonic | 671 | 342.85 MB/s | 801 | 8 | 4.9× |
| SonicFastest | 671 | 342.64 MB/s | 801 | 8 | 4.9× |
| LightningDecodeAny | 1096 | 208.91 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1187 | 193.72 MB/s | 448 | 3 | 2.7× |
| Goccy | 1302 | 176.65 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 1809 | 127.15 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3264 | 70.46 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 65473 | 994.80 MB/s | 172433 | 107 | 7.8× |
| LightningDestructive | 68064 | 956.93 MB/s | 166213 | 102 | 7.5× |
| Sonic | 121335 | 536.79 MB/s | 236023 | 65 | 4.2× |
| SonicFastest | 121538 | 535.90 MB/s | 236076 | 65 | 4.2× |
| Goccy | 142157 | 458.17 MB/s | 228583 | 134 | 3.6× |
| LightningDecodeAny | 147672 | 361.13 MB/s | 176960 | 3252 | 3.4× |
| JSONV2 | 199688 | 326.17 MB/s | 206665 | 607 | 2.5× |
| Stdlib | 507966 | 128.22 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1969957 | 985.03 MB/s | 2846865 | 2238 | 9.9× |
| Lightning | 2042389 | 950.10 MB/s | 2846866 | 2238 | 9.6× |
| SonicFastest | 3727507 | 520.58 MB/s | 4881182 | 1736 | 5.2× |
| Sonic | 3738333 | 519.07 MB/s | 4880519 | 1736 | 5.2× |
| Goccy | 3887498 | 499.16 MB/s | 4062444 | 13509 | 5.0× |
| Easyjson | 5966624 | 325.22 MB/s | 3871265 | 15043 | 3.3× |
| LightningDecodeAny | 7214494 | 268.97 MB/s | 7013525 | 219937 | 2.7× |
| JSONV2 | 8741784 | 221.98 MB/s | 3237194 | 13947 | 2.2× |
| Stdlib | 19507236 | 99.47 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 794183 | 4190.26 MB/s | 351704 | 1286 | 25.0× |
| Lightning | 1232883 | 2699.23 MB/s | 2488906 | 2995 | 16.1× |
| Sonic | 1748116 | 1903.67 MB/s | 5895825 | 4263 | 11.4× |
| SonicFastest | 1755323 | 1895.85 MB/s | 5894506 | 4263 | 11.3× |
| LightningDecodeAny | 2622700 | 1171.99 MB/s | 4886621 | 56892 | 7.6× |
| Goccy | 5128156 | 648.93 MB/s | 3948913 | 3816 | 3.9× |
| JSONV2 | 6778017 | 490.97 MB/s | 5364504 | 13243 | 2.9× |
| Stdlib | 19874981 | 167.44 MB/s | 5565606 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 177345 | 1242.47 MB/s | 136896 | 228 | 9.7× |
| LightningDestructive | 184217 | 1196.12 MB/s | 136896 | 228 | 9.4× |
| Goccy | 357788 | 615.86 MB/s | 363916 | 1066 | 4.8× |
| Sonic | 398349 | 553.15 MB/s | 349951 | 262 | 4.3× |
| SonicFastest | 399952 | 550.93 MB/s | 350024 | 262 | 4.3× |
| Easyjson | 460474 | 478.52 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 505768 | 435.67 MB/s | 129746 | 470 | 3.4× |
| LightningDecodeAny | 731740 | 148.02 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 1727080 | 127.58 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 10250864 | 790.18 MB/s | 15059833 | 51649 | 7.6× |
| Lightning | 10517852 | 770.12 MB/s | 15059838 | 51649 | 7.4× |
| Sonic | 15927769 | 508.55 MB/s | 19860980 | 41640 | 4.9× |
| SonicFastest | 16006276 | 506.05 MB/s | 19860031 | 41640 | 4.9× |
| Goccy | 19292892 | 419.85 MB/s | 19024478 | 107155 | 4.0× |
| Easyjson | 26290606 | 308.10 MB/s | 15059616 | 41643 | 3.0× |
| LightningDecodeAny | 31666461 | 164.31 MB/s | 34333346 | 912810 | 2.5× |
| JSONV2 | 34245911 | 236.53 MB/s | 15233730 | 78972 | 2.3× |
| Stdlib | 77713704 | 104.23 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4676572 | 637.96 MB/s | 3985336 | 30015 | 8.4× |
| Lightning | 4736892 | 629.84 MB/s | 3985337 | 30015 | 8.3× |
| SonicFastest | 7386536 | 403.91 MB/s | 9131949 | 57804 | 5.3× |
| Sonic | 7390064 | 403.71 MB/s | 9132883 | 57804 | 5.3× |
| Goccy | 13724173 | 217.39 MB/s | 10021113 | 273625 | 2.9× |
| Easyjson | 14070835 | 212.03 MB/s | 9479441 | 30115 | 2.8× |
| LightningDecodeAny | 14943951 | 122.74 MB/s | 20023834 | 408557 | 2.6× |
| JSONV2 | 19208461 | 155.32 MB/s | 9257050 | 86278 | 2.0× |
| Stdlib | 39374243 | 75.77 MB/s | 9258086 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1070315 | 676.06 MB/s | 907395 | 3618 | 8.9× |
| Lightning | 1128951 | 640.95 MB/s | 907386 | 3618 | 8.4× |
| Sonic | 1608028 | 449.99 MB/s | 2374378 | 3683 | 5.9× |
| SonicFastest | 1617188 | 447.44 MB/s | 2375676 | 3683 | 5.9× |
| LightningDecodeAny | 3901747 | 166.74 MB/s | 5752204 | 80172 | 2.4× |
| Easyjson | 4089208 | 176.95 MB/s | 2847905 | 3698 | 2.3× |
| Goccy | 4145086 | 174.57 MB/s | 2699549 | 80267 | 2.3× |
| JSONV2 | 4754757 | 152.18 MB/s | 2704707 | 7318 | 2.0× |
| Stdlib | 9495461 | 76.20 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1604726 | 982.94 MB/s | 907388 | 3618 | 8.6× |
| LightningDestructive | 1618488 | 974.58 MB/s | 907392 | 3618 | 8.5× |
| SonicFastest | 1893282 | 833.13 MB/s | 3224576 | 3683 | 7.3× |
| Sonic | 1913850 | 824.18 MB/s | 3226362 | 3683 | 7.2× |
| LightningDecodeAny | 3260187 | 231.09 MB/s | 5752199 | 80172 | 4.2× |
| Easyjson | 4846562 | 325.46 MB/s | 2847906 | 3698 | 2.8× |
| Goccy | 5072068 | 310.99 MB/s | 3470020 | 80260 | 2.7× |
| JSONV2 | 5477181 | 287.99 MB/s | 2704553 | 7318 | 2.5× |
| Stdlib | 13755364 | 114.67 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 182574 | 822.26 MB/s | 81920 | 1 | 9.4× |
| LightningDestructive | 187449 | 800.88 MB/s | 81920 | 1 | 9.1× |
| Sonic | 301769 | 497.48 MB/s | 408791 | 16 | 5.7× |
| SonicFastest | 334200 | 449.20 MB/s | 411323 | 16 | 5.1× |
| LightningDecodeAny | 432632 | 346.99 MB/s | 746003 | 10020 | 4.0× |
| Goccy | 848122 | 177.01 MB/s | 333588 | 10005 | 2.0× |
| JSONV2 | 1046793 | 143.41 MB/s | 357724 | 20 | 1.6× |
| Stdlib | 1713443 | 87.62 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 24546 | 1145.48 MB/s | 30272 | 105 | 9.7× |
| LightningDestructive | 25317 | 1110.59 MB/s | 30144 | 103 | 9.4× |
| SonicFastest | 52365 | 536.94 MB/s | 59450 | 83 | 4.6× |
| Sonic | 52691 | 533.62 MB/s | 59469 | 83 | 4.5× |
| Easyjson | 58855 | 477.74 MB/s | 32304 | 138 | 4.1× |
| Goccy | 59581 | 471.91 MB/s | 59267 | 188 | 4.0× |
| JSONV2 | 96324 | 291.90 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 119799 | 234.70 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 238923 | 117.68 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1516 | 1535.18 MB/s | 32 | 1 | 12.1× |
| LightningDestructive | 1554 | 1498.07 MB/s | 32 | 1 | 11.8× |
| Goccy | 3575 | 651.13 MB/s | 3649 | 4 | 5.1× |
| Easyjson | 3747 | 621.24 MB/s | 192 | 2 | 4.9× |
| Sonic | 4800 | 485.03 MB/s | 3708 | 4 | 3.8× |
| SonicFastest | 4801 | 484.89 MB/s | 3710 | 4 | 3.8× |
| JSONV2 | 6009 | 387.40 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 7480 | 225.26 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 18310 | 127.14 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 172 | 1095.61 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 175 | 1081.68 MB/s | 0 | 0 | 11.0× |
| Goccy | 311 | 607.60 MB/s | 304 | 2 | 6.2× |
| Easyjson | 435 | 434.31 MB/s | 0 | 0 | 4.4× |
| Sonic | 570 | 331.42 MB/s | 341 | 3 | 3.4× |
| SonicFastest | 571 | 331.06 MB/s | 341 | 3 | 3.4× |
| JSONV2 | 715 | 264.38 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 939 | 142.64 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 1914 | 98.74 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1144 | 1914.54 MB/s | 0 | 0 | 11.7× |
| LightningDestructive | 1187 | 1845.47 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2661 | 823.48 MB/s | 24 | 1 | 5.0× |
| Goccy | 2684 | 816.24 MB/s | 2864 | 4 | 5.0× |
| SonicFastest | 4950 | 442.65 MB/s | 3602 | 38 | 2.7× |
| Sonic | 5131 | 427.00 MB/s | 3602 | 38 | 2.6× |
| JSONV2 | 5900 | 371.37 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 6632 | 273.08 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 13412 | 163.37 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 514540 | 992.10 MB/s | 703776 | 1012 | 9.7× |
| Lightning | 536716 | 951.11 MB/s | 703777 | 1012 | 9.3× |
| Goccy | 974540 | 523.81 MB/s | 1137827 | 5006 | 5.1× |
| Sonic | 1129403 | 451.99 MB/s | 1306298 | 2014 | 4.4× |
| SonicFastest | 1143540 | 446.40 MB/s | 1307950 | 2014 | 4.4× |
| Easyjson | 1170311 | 436.19 MB/s | 863776 | 3012 | 4.3× |
| JSONV2 | 2392739 | 213.34 MB/s | 1075958 | 12645 | 2.1× |
| LightningDecodeAny | 2633430 | 175.23 MB/s | 2785928 | 66022 | 1.9× |
| Stdlib | 4984640 | 102.41 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 475 | 41627.57 MB/s | 0 | 0 | 234.3× |
| LightningDestructive | 681 | 29057.10 MB/s | 0 | 0 | 163.6× |
| SonicFastest | 4894 | 4043.41 MB/s | 21140 | 3 | 22.8× |
| Goccy | 19806 | 999.12 MB/s | 20492 | 2 | 5.6× |
| Sonic | 22698 | 871.84 MB/s | 20613 | 3 | 4.9× |
| JSONV2 | 27943 | 708.19 MB/s | 8 | 1 | 4.0× |
| LightningDecodeAny | 73274 | 270.06 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 84617 | 233.87 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111405 | 177.63 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1728 | 10489.51 MB/s | 432 | 2 | 53.4× |
| LightningDestructive | 1792 | 10115.07 MB/s | 0 | 0 | 51.5× |
| Easyjson | 3641 | 4978.40 MB/s | 432 | 2 | 25.3× |
| Sonic | 7084 | 2558.37 MB/s | 20443 | 5 | 13.0× |
| SonicFastest | 7404 | 2447.97 MB/s | 20406 | 5 | 12.5× |
| LightningDecodeAny | 14112 | 1267.16 MB/s | 29088 | 191 | 6.5× |
| Goccy | 16234 | 1116.41 MB/s | 19460 | 2 | 5.7× |
| JSONV2 | 37272 | 486.26 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 92260 | 196.44 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1712696 | 1172.71 MB/s | 2960389 | 7411 | 9.4× |
| Lightning | 1769818 | 1134.86 MB/s | 2962101 | 7417 | 9.1× |
| Goccy | 3457524 | 580.91 MB/s | 5410488 | 15829 | 4.6× |
| SonicFastest | 3835891 | 523.61 MB/s | 5151962 | 7085 | 4.2× |
| Sonic | 3864108 | 519.78 MB/s | 5152831 | 7085 | 4.1× |
| Easyjson | 4087964 | 491.32 MB/s | 2981483 | 7439 | 3.9× |
| LightningDecodeAny | 5277135 | 216.46 MB/s | 7386074 | 134751 | 3.0× |
| JSONV2 | 5509789 | 364.53 MB/s | 3173673 | 14562 | 2.9× |
| Stdlib | 16030914 | 125.29 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 632 | 869.22 MB/s | 480 | 1 | 7.1× |
| LightningDestructive | 633 | 867.61 MB/s | 480 | 1 | 7.1× |
| LightningDecodeAny | 1473 | 371.92 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 1544 | 355.53 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 1656 | 331.43 MB/s | 2263 | 8 | 2.7× |
| Sonic | 1694 | 324.00 MB/s | 2261 | 8 | 2.7× |
| Goccy | 2313 | 237.34 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 2849 | 192.68 MB/s | 1664 | 7 | 1.6× |
| Stdlib | 4512 | 121.67 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 373174 | 1692.28 MB/s | 364472 | 566 | 12.5× |
| Lightning | 419439 | 1505.62 MB/s | 413001 | 878 | 11.1× |
| SonicFastest | 819027 | 771.05 MB/s | 1065074 | 814 | 5.7× |
| Sonic | 823333 | 767.02 MB/s | 1065210 | 814 | 5.7× |
| Easyjson | 940594 | 671.40 MB/s | 422504 | 936 | 5.0× |
| Goccy | 1071292 | 589.49 MB/s | 990837 | 1200 | 4.4× |
| JSONV2 | 1680958 | 375.69 MB/s | 571593 | 3144 | 2.8× |
| LightningDecodeAny | 1824182 | 255.95 MB/s | 2010198 | 30295 | 2.6× |
| Stdlib | 4661984 | 135.46 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 589662 | 953.78 MB/s | 544249 | 448 | 7.4× |
| Lightning | 731818 | 768.51 MB/s | 767620 | 1254 | 6.0× |
| Sonic | 1086790 | 517.49 MB/s | 1349778 | 1185 | 4.0× |
| SonicFastest | 1089535 | 516.19 MB/s | 1349544 | 1185 | 4.0× |
| Goccy | 1216704 | 462.24 MB/s | 1037333 | 1028 | 3.6× |
| Easyjson | 1540779 | 365.02 MB/s | 775153 | 1254 | 2.8× |
| LightningDecodeAny | 2271721 | 247.57 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 2321786 | 242.23 MB/s | 927405 | 3482 | 1.9× |
| Stdlib | 4386273 | 128.22 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 450028 | 1184.77 MB/s | 333416 | 2084 | 10.3× |
| Lightning | 519913 | 1025.51 MB/s | 368224 | 2293 | 8.9× |
| Easyjson | 959458 | 555.71 MB/s | 428362 | 3273 | 4.8× |
| Sonic | 1012786 | 526.45 MB/s | 982242 | 3082 | 4.6× |
| SonicFastest | 1018733 | 523.37 MB/s | 981392 | 3082 | 4.5× |
| Goccy | 1106101 | 482.03 MB/s | 1167062 | 5408 | 4.2× |
| JSONV2 | 2076770 | 256.73 MB/s | 745419 | 13288 | 2.2× |
| LightningDecodeAny | 2567815 | 207.64 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 4622620 | 115.34 MB/s | 798692 | 17133 | 1.0× |
