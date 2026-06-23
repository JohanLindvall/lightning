# JSON Deserialization Benchmarks

- generated 2026-06-23T19:53:38Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104258 | 1220.76 MB/s | 49760 | 3 | 12.6× |
| LightningDestructive | 106065 | 1199.98 MB/s | 49280 | 2 | 12.4× |
| SonicFastest | 198228 | 642.06 MB/s | 214566 | 15 | 6.6× |
| Sonic | 200788 | 633.88 MB/s | 214607 | 15 | 6.6× |
| Goccy | 247695 | 513.84 MB/s | 225421 | 884 | 5.3× |
| Easyjson | 254085 | 500.92 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 446372 | 285.13 MB/s | 195128 | 1805 | 3.0× |
| LightningDecodeAny | 472654 | 200.26 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1317892 | 96.57 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4105657 | 548.28 MB/s | 3122872 | 3081 | 8.6× |
| Lightning | 4203280 | 535.55 MB/s | 3122873 | 3081 | 8.4× |
| SonicFastest | 5254802 | 428.38 MB/s | 4870174 | 2584 | 6.8× |
| Sonic | 5375328 | 418.77 MB/s | 4868193 | 2584 | 6.6× |
| Goccy | 11870844 | 189.63 MB/s | 4200966 | 56535 | 3.0× |
| LightningDecodeAny | 12310093 | 182.86 MB/s | 7938300 | 281383 | 2.9× |
| Easyjson | 13783593 | 163.31 MB/s | 3099809 | 2120 | 2.6× |
| JSONV2 | 17023176 | 132.23 MB/s | 3123189 | 3083 | 2.1× |
| Stdlib | 35490375 | 63.43 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 573904 | 471.16 MB/s | 348025 | 1627 | 7.6× |
| LightningDestructive | 577881 | 467.92 MB/s | 348025 | 1627 | 7.6× |
| Sonic | 800467 | 337.81 MB/s | 642285 | 1147 | 5.5× |
| SonicFastest | 803356 | 336.59 MB/s | 642456 | 1147 | 5.5× |
| Goccy | 1653790 | 163.51 MB/s | 542053 | 8122 | 2.7× |
| LightningDecodeAny | 1710724 | 158.06 MB/s | 1011488 | 37901 | 2.6× |
| Easyjson | 1773929 | 152.43 MB/s | 330272 | 749 | 2.5× |
| JSONV2 | 2318957 | 116.61 MB/s | 348160 | 1628 | 1.9× |
| Stdlib | 4390054 | 61.59 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1292599 | 1336.23 MB/s | 959848 | 5881 | 13.7× |
| Lightning | 1338673 | 1290.24 MB/s | 959890 | 5882 | 13.2× |
| Sonic | 2173065 | 794.82 MB/s | 2695457 | 5547 | 8.2× |
| SonicFastest | 2173927 | 794.51 MB/s | 2695254 | 5547 | 8.1× |
| Goccy | 2534090 | 681.59 MB/s | 2580567 | 14603 | 7.0× |
| Easyjson | 3846593 | 449.02 MB/s | 972032 | 5389 | 4.6× |
| LightningDecodeAny | 4439250 | 112.70 MB/s | 4976203 | 81466 | 4.0× |
| JSONV2 | 4735387 | 364.74 MB/s | 1011615 | 7594 | 3.7× |
| Stdlib | 17714534 | 97.50 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1021 | 1774.26 MB/s | 0 | 0 | 15.9× |
| LightningDestructive | 1109 | 1633.25 MB/s | 0 | 0 | 14.6× |
| Easyjson | 2932 | 618.05 MB/s | 24 | 1 | 5.5× |
| Goccy | 3267 | 554.63 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6243 | 290.27 MB/s | 3345 | 38 | 2.6× |
| Sonic | 6436 | 281.55 MB/s | 3343 | 38 | 2.5× |
| JSONV2 | 8235 | 220.03 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9236 | 196.09 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 16220 | 111.72 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1121 | 1616.31 MB/s | 0 | 0 | 14.5× |
| LightningDestructive | 1199 | 1511.72 MB/s | 0 | 0 | 13.6× |
| Easyjson | 3031 | 597.83 MB/s | 24 | 1 | 5.4× |
| Goccy | 3187 | 568.56 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5989 | 302.54 MB/s | 3345 | 38 | 2.7× |
| Sonic | 6200 | 292.27 MB/s | 3344 | 38 | 2.6× |
| JSONV2 | 8102 | 223.65 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9279 | 195.18 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 16299 | 111.17 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1318 | 1375.16 MB/s | 144 | 10 | 12.3× |
| LightningDestructive | 1408 | 1286.99 MB/s | 144 | 10 | 11.5× |
| Easyjson | 3092 | 586.03 MB/s | 144 | 10 | 5.2× |
| Goccy | 3500 | 517.70 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 6136 | 295.32 MB/s | 3362 | 40 | 2.6× |
| Sonic | 6346 | 285.53 MB/s | 3365 | 40 | 2.5× |
| JSONV2 | 8576 | 211.28 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9338 | 193.94 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16173 | 112.04 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 728 | 678.56 MB/s | 160 | 1 | 9.6× |
| LightningDestructive | 741 | 666.51 MB/s | 160 | 1 | 9.4× |
| Sonic | 1290 | 382.92 MB/s | 1076 | 8 | 5.4× |
| SonicFastest | 1293 | 382.13 MB/s | 1076 | 8 | 5.4× |
| LightningDecodeAny | 1782 | 276.64 MB/s | 1536 | 30 | 3.9× |
| Goccy | 2585 | 191.10 MB/s | 856 | 23 | 2.7× |
| Easyjson | 2698 | 183.13 MB/s | 448 | 3 | 2.6× |
| JSONV2 | 3468 | 142.45 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6984 | 70.73 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 495 | 464.25 MB/s | 160 | 1 | 10.0× |
| LightningDestructive | 498 | 461.91 MB/s | 160 | 1 | 9.9× |
| SonicFastest | 926 | 248.52 MB/s | 800 | 8 | 5.4× |
| Sonic | 928 | 247.87 MB/s | 801 | 8 | 5.3× |
| LightningDecodeAny | 1617 | 141.58 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1799 | 127.83 MB/s | 448 | 3 | 2.8× |
| Goccy | 1802 | 127.62 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2604 | 88.34 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4952 | 46.44 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 94815 | 686.94 MB/s | 172433 | 107 | 7.2× |
| LightningDestructive | 106596 | 611.02 MB/s | 166213 | 102 | 6.4× |
| SonicFastest | 147906 | 440.36 MB/s | 235997 | 65 | 4.6× |
| Sonic | 150111 | 433.89 MB/s | 236080 | 65 | 4.6× |
| Goccy | 177962 | 365.99 MB/s | 228185 | 134 | 3.9× |
| LightningDecodeAny | 207476 | 257.04 MB/s | 176960 | 3252 | 3.3× |
| JSONV2 | 262293 | 248.32 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 686285 | 94.91 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2512913 | 772.20 MB/s | 2846864 | 2238 | 11.3× |
| Lightning | 2590432 | 749.09 MB/s | 2846865 | 2238 | 11.0× |
| Sonic | 4957454 | 391.43 MB/s | 4878792 | 1736 | 5.7× |
| Goccy | 5003932 | 387.79 MB/s | 4062911 | 13509 | 5.7× |
| SonicFastest | 5022985 | 386.32 MB/s | 4880316 | 1736 | 5.6× |
| Easyjson | 8127213 | 238.76 MB/s | 3871265 | 15043 | 3.5× |
| LightningDecodeAny | 9950377 | 195.01 MB/s | 7013525 | 219937 | 2.9× |
| JSONV2 | 12399586 | 156.49 MB/s | 3237187 | 13947 | 2.3× |
| Stdlib | 28377280 | 68.38 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1072040 | 3104.20 MB/s | 351704 | 1286 | 21.9× |
| Lightning | 1619207 | 2055.22 MB/s | 2488908 | 2995 | 14.5× |
| SonicFastest | 2039271 | 1631.87 MB/s | 5896494 | 4263 | 11.5× |
| Sonic | 2043954 | 1628.13 MB/s | 5896627 | 4263 | 11.5× |
| LightningDecodeAny | 3553556 | 864.98 MB/s | 4886622 | 56892 | 6.6× |
| Goccy | 6078392 | 547.49 MB/s | 3948914 | 3817 | 3.9× |
| JSONV2 | 7673817 | 433.66 MB/s | 5364510 | 13243 | 3.1× |
| Stdlib | 23447641 | 141.93 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224608 | 981.02 MB/s | 136896 | 228 | 10.9× |
| LightningDestructive | 234780 | 938.52 MB/s | 136896 | 228 | 10.4× |
| SonicFastest | 477348 | 461.60 MB/s | 350012 | 262 | 5.1× |
| Sonic | 478988 | 460.02 MB/s | 350063 | 262 | 5.1× |
| Goccy | 485121 | 454.21 MB/s | 363935 | 1066 | 5.0× |
| Easyjson | 674686 | 326.59 MB/s | 130512 | 245 | 3.6× |
| JSONV2 | 720603 | 305.78 MB/s | 129746 | 470 | 3.4× |
| LightningDecodeAny | 1020106 | 106.18 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2444685 | 90.13 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14578342 | 555.62 MB/s | 15059835 | 51649 | 7.7× |
| Lightning | 15665904 | 517.05 MB/s | 15059844 | 51649 | 7.2× |
| Sonic | 17863822 | 453.43 MB/s | 19861285 | 41640 | 6.3× |
| SonicFastest | 17980255 | 450.50 MB/s | 19858611 | 41640 | 6.2× |
| Goccy | 24850632 | 325.95 MB/s | 19225163 | 107156 | 4.5× |
| Easyjson | 37486847 | 216.08 MB/s | 15059618 | 41643 | 3.0× |
| LightningDecodeAny | 43039679 | 120.89 MB/s | 34333346 | 912810 | 2.6× |
| JSONV2 | 50038920 | 161.87 MB/s | 15233698 | 78972 | 2.2× |
| Stdlib | 112113988 | 72.25 MB/s | 15665071 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6196996 | 481.44 MB/s | 3985336 | 30015 | 9.2× |
| Lightning | 6246467 | 477.62 MB/s | 3985339 | 30015 | 9.1× |
| Sonic | 9284185 | 321.35 MB/s | 9129376 | 57804 | 6.2× |
| SonicFastest | 9300008 | 320.80 MB/s | 9129864 | 57804 | 6.1× |
| Goccy | 18451131 | 161.70 MB/s | 9751093 | 273614 | 3.1× |
| Easyjson | 18739904 | 159.20 MB/s | 9479442 | 30115 | 3.0× |
| LightningDecodeAny | 20685373 | 88.67 MB/s | 20023835 | 408557 | 2.8× |
| JSONV2 | 28929467 | 103.13 MB/s | 9257040 | 86278 | 2.0× |
| Stdlib | 57144930 | 52.21 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1465737 | 493.67 MB/s | 907392 | 3618 | 9.6× |
| Lightning | 1542158 | 469.21 MB/s | 907387 | 3618 | 9.1× |
| SonicFastest | 2097878 | 344.92 MB/s | 2371621 | 3683 | 6.7× |
| Sonic | 2098837 | 344.76 MB/s | 2371587 | 3683 | 6.7× |
| Easyjson | 5401877 | 133.95 MB/s | 2847907 | 3698 | 2.6× |
| Goccy | 5466905 | 132.36 MB/s | 2705731 | 80268 | 2.6× |
| LightningDecodeAny | 5624018 | 115.68 MB/s | 5752202 | 80172 | 2.5× |
| JSONV2 | 6324494 | 114.41 MB/s | 2704706 | 7318 | 2.2× |
| Stdlib | 14032174 | 51.57 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2026717 | 778.28 MB/s | 907387 | 3618 | 9.9× |
| LightningDestructive | 2040075 | 773.18 MB/s | 907392 | 3618 | 9.8× |
| Sonic | 2419152 | 652.03 MB/s | 3223642 | 3683 | 8.3× |
| SonicFastest | 2430385 | 649.01 MB/s | 3224332 | 3683 | 8.2× |
| LightningDecodeAny | 4852460 | 155.26 MB/s | 5752200 | 80172 | 4.1× |
| Easyjson | 6354493 | 248.23 MB/s | 2847904 | 3698 | 3.2× |
| Goccy | 6634855 | 237.74 MB/s | 3493247 | 80262 | 3.0× |
| JSONV2 | 6894868 | 228.77 MB/s | 2704553 | 7318 | 2.9× |
| Stdlib | 20039337 | 78.71 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 249868 | 600.81 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 260535 | 576.21 MB/s | 81920 | 1 | 8.3× |
| Sonic | 382428 | 392.55 MB/s | 407277 | 16 | 5.6× |
| SonicFastest | 383617 | 391.34 MB/s | 407157 | 16 | 5.6× |
| LightningDecodeAny | 597932 | 251.07 MB/s | 746006 | 10020 | 3.6× |
| Goccy | 1027063 | 146.17 MB/s | 324441 | 10005 | 2.1× |
| JSONV2 | 1106728 | 135.65 MB/s | 357726 | 20 | 1.9× |
| Stdlib | 2154461 | 69.68 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 34193 | 822.31 MB/s | 30272 | 105 | 10.0× |
| LightningDestructive | 35284 | 796.89 MB/s | 30144 | 103 | 9.6× |
| Sonic | 56917 | 494.00 MB/s | 59548 | 83 | 6.0× |
| SonicFastest | 57073 | 492.65 MB/s | 59516 | 83 | 6.0× |
| Goccy | 80330 | 350.02 MB/s | 59282 | 188 | 4.2× |
| Easyjson | 81461 | 345.16 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 131964 | 213.07 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 167501 | 167.86 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 340451 | 82.59 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1966 | 1184.01 MB/s | 32 | 1 | 13.2× |
| LightningDestructive | 2101 | 1107.79 MB/s | 32 | 1 | 12.4× |
| SonicFastest | 4772 | 487.82 MB/s | 3711 | 4 | 5.5× |
| Sonic | 4777 | 487.35 MB/s | 3711 | 4 | 5.5× |
| Goccy | 4956 | 469.75 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 5100 | 456.44 MB/s | 192 | 2 | 5.1× |
| JSONV2 | 8384 | 277.66 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10758 | 156.62 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 26049 | 89.37 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 215 | 879.66 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 224 | 844.50 MB/s | 0 | 0 | 12.5× |
| Goccy | 427 | 442.30 MB/s | 304 | 2 | 6.5× |
| Easyjson | 590 | 320.05 MB/s | 0 | 0 | 4.7× |
| SonicFastest | 632 | 299.00 MB/s | 341 | 3 | 4.4× |
| Sonic | 641 | 294.71 MB/s | 341 | 3 | 4.4× |
| JSONV2 | 1011 | 187.02 MB/s | 112 | 1 | 2.8× |
| LightningDecodeAny | 1340 | 100.01 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2797 | 67.58 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1494 | 1466.48 MB/s | 0 | 0 | 12.7× |
| LightningDestructive | 1606 | 1364.55 MB/s | 0 | 0 | 11.8× |
| Goccy | 3570 | 613.66 MB/s | 2864 | 4 | 5.3× |
| Easyjson | 3814 | 574.50 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6424 | 341.09 MB/s | 3600 | 38 | 2.9× |
| Sonic | 6665 | 328.73 MB/s | 3599 | 38 | 2.8× |
| JSONV2 | 8220 | 266.53 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 9404 | 192.58 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18919 | 115.81 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 698542 | 730.77 MB/s | 703779 | 1012 | 10.1× |
| Lightning | 742748 | 687.28 MB/s | 703778 | 1012 | 9.5× |
| SonicFastest | 1306019 | 390.86 MB/s | 1308746 | 2014 | 5.4× |
| Sonic | 1310415 | 389.55 MB/s | 1310650 | 2014 | 5.4× |
| Goccy | 1325565 | 385.10 MB/s | 1138636 | 5006 | 5.3× |
| Easyjson | 1718098 | 297.12 MB/s | 863779 | 3012 | 4.1× |
| JSONV2 | 3493235 | 146.13 MB/s | 1075967 | 12645 | 2.0× |
| LightningDecodeAny | 3643247 | 126.66 MB/s | 2785927 | 66022 | 1.9× |
| Stdlib | 7083702 | 72.06 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 644 | 30733.14 MB/s | 0 | 0 | 251.4× |
| LightningDestructive | 960 | 20610.57 MB/s | 0 | 0 | 168.6× |
| SonicFastest | 6856 | 2886.51 MB/s | 21119 | 3 | 23.6× |
| Goccy | 29337 | 674.53 MB/s | 20492 | 2 | 5.5× |
| Sonic | 32771 | 603.86 MB/s | 20645 | 3 | 4.9× |
| JSONV2 | 33301 | 594.25 MB/s | 8 | 1 | 4.9× |
| Easyjson | 101155 | 195.63 MB/s | 0 | 0 | 1.6× |
| LightningDecodeAny | 102863 | 192.37 MB/s | 117104 | 2019 | 1.6× |
| Stdlib | 161852 | 122.27 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2200 | 8238.22 MB/s | 432 | 2 | 54.8× |
| LightningDestructive | 2483 | 7300.35 MB/s | 0 | 0 | 48.5× |
| Easyjson | 5010 | 3617.79 MB/s | 432 | 2 | 24.0× |
| Sonic | 8575 | 2113.48 MB/s | 20455 | 5 | 14.0× |
| SonicFastest | 8720 | 2078.33 MB/s | 20439 | 5 | 13.8× |
| LightningDecodeAny | 19316 | 925.76 MB/s | 29088 | 191 | 6.2× |
| Goccy | 22437 | 807.76 MB/s | 19460 | 2 | 5.4× |
| JSONV2 | 50624 | 358.01 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 120455 | 150.46 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2252650 | 891.61 MB/s | 2960389 | 7411 | 9.8× |
| Lightning | 2409160 | 833.69 MB/s | 2962102 | 7417 | 9.1× |
| SonicFastest | 4185015 | 479.93 MB/s | 5157522 | 7085 | 5.3× |
| Sonic | 4273827 | 469.95 MB/s | 5156225 | 7085 | 5.1× |
| Goccy | 4776023 | 420.54 MB/s | 5411630 | 15832 | 4.6× |
| Easyjson | 5701266 | 352.29 MB/s | 2981488 | 7439 | 3.9× |
| LightningDecodeAny | 7201387 | 158.62 MB/s | 7386073 | 134751 | 3.1× |
| JSONV2 | 7745045 | 259.33 MB/s | 3173675 | 14563 | 2.8× |
| Stdlib | 21982934 | 91.37 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 848 | 647.53 MB/s | 480 | 1 | 8.1× |
| Lightning | 851 | 644.85 MB/s | 480 | 1 | 8.1× |
| LightningDecodeAny | 2142 | 255.81 MB/s | 2261 | 50 | 3.2× |
| Easyjson | 2438 | 225.22 MB/s | 1616 | 5 | 2.8× |
| SonicFastest | 2457 | 223.46 MB/s | 2261 | 8 | 2.8× |
| Sonic | 2474 | 221.87 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3379 | 162.48 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3569 | 153.81 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6876 | 79.85 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 502498 | 1256.75 MB/s | 364473 | 566 | 12.7× |
| Lightning | 584763 | 1079.95 MB/s | 413001 | 878 | 10.9× |
| Sonic | 988584 | 638.81 MB/s | 1066627 | 814 | 6.5× |
| SonicFastest | 1014899 | 622.24 MB/s | 1067364 | 814 | 6.3× |
| Easyjson | 1305677 | 483.67 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1382943 | 456.65 MB/s | 988817 | 1200 | 4.6× |
| JSONV2 | 2279287 | 277.07 MB/s | 571591 | 3144 | 2.8× |
| LightningDecodeAny | 2571816 | 181.55 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6398801 | 98.69 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 828789 | 678.59 MB/s | 544248 | 448 | 7.3× |
| Lightning | 1036710 | 542.49 MB/s | 767621 | 1254 | 5.9× |
| Sonic | 1267160 | 443.83 MB/s | 1347791 | 1185 | 4.8× |
| SonicFastest | 1367256 | 411.34 MB/s | 1351163 | 1185 | 4.4× |
| Goccy | 1605145 | 350.38 MB/s | 1039267 | 1029 | 3.8× |
| Easyjson | 2117588 | 265.59 MB/s | 775154 | 1254 | 2.9× |
| LightningDecodeAny | 3115095 | 180.54 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 3125869 | 179.92 MB/s | 927407 | 3482 | 1.9× |
| Stdlib | 6073596 | 92.60 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 613082 | 869.67 MB/s | 333416 | 2084 | 10.2× |
| Lightning | 689976 | 772.75 MB/s | 368224 | 2293 | 9.1× |
| Sonic | 1108651 | 480.92 MB/s | 981235 | 3082 | 5.7× |
| SonicFastest | 1111564 | 479.66 MB/s | 981449 | 3082 | 5.6× |
| Easyjson | 1297186 | 411.03 MB/s | 428362 | 3273 | 4.8× |
| Goccy | 1393233 | 382.69 MB/s | 1167089 | 5409 | 4.5× |
| JSONV2 | 2832220 | 188.25 MB/s | 745420 | 13288 | 2.2× |
| LightningDecodeAny | 3464216 | 153.91 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6263908 | 85.12 MB/s | 798693 | 17133 | 1.0× |
