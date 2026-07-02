# JSON Deserialization Benchmarks

- generated 2026-07-02T14:32:47Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 103049 | 1235.10 MB/s | 49760 | 3 | 12.1× |
| LightningDestructive | 104976 | 1212.41 MB/s | 49280 | 2 | 11.9× |
| Sonic | 203819 | 624.45 MB/s | 214442 | 15 | 6.1× |
| SonicFastest | 204682 | 621.82 MB/s | 214193 | 15 | 6.1× |
| Easyjson | 231027 | 550.91 MB/s | 122864 | 14 | 5.4× |
| Goccy | 243254 | 523.22 MB/s | 225185 | 884 | 5.1× |
| JSONV2 | 394433 | 322.68 MB/s | 195128 | 1805 | 3.2× |
| LightningDecodeAny | 434537 | 217.83 MB/s | 465730 | 9708 | 2.9× |
| Stdlib | 1245728 | 102.17 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4081216 | 551.56 MB/s | 3008241 | 1158 | 7.6× |
| Lightning | 4222072 | 533.16 MB/s | 3008242 | 1158 | 7.4× |
| Sonic | 5521670 | 407.68 MB/s | 4866660 | 2584 | 5.6× |
| SonicFastest | 5681067 | 396.24 MB/s | 4868410 | 2584 | 5.5× |
| Goccy | 12811760 | 175.70 MB/s | 4135077 | 56532 | 2.4× |
| LightningDecodeAny | 13444621 | 167.43 MB/s | 19380211 | 223896 | 2.3× |
| Easyjson | 13753693 | 163.67 MB/s | 3099810 | 2120 | 2.3× |
| JSONV2 | 16877913 | 133.37 MB/s | 3123190 | 3083 | 1.8× |
| Stdlib | 31129887 | 72.31 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 530375 | 509.83 MB/s | 392944 | 568 | 7.6× |
| LightningDestructive | 551330 | 490.46 MB/s | 392944 | 568 | 7.3× |
| Sonic | 742052 | 364.40 MB/s | 642908 | 1147 | 5.5× |
| SonicFastest | 750038 | 360.52 MB/s | 642524 | 1147 | 5.4× |
| Easyjson | 1764689 | 153.23 MB/s | 330272 | 749 | 2.3× |
| Goccy | 1773985 | 152.43 MB/s | 543481 | 8122 | 2.3× |
| LightningDecodeAny | 2171852 | 124.50 MB/s | 2543878 | 29687 | 1.9× |
| JSONV2 | 2238062 | 120.82 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4050861 | 66.75 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1201231 | 1437.86 MB/s | 767864 | 2798 | 14.1× |
| Lightning | 1213685 | 1423.11 MB/s | 767906 | 2799 | 13.9× |
| Sonic | 2245387 | 769.22 MB/s | 2693612 | 5547 | 7.5× |
| SonicFastest | 2273708 | 759.64 MB/s | 2693062 | 5547 | 7.4× |
| Goccy | 2924472 | 590.60 MB/s | 2581484 | 14603 | 5.8× |
| LightningDecodeAny | 3903644 | 128.16 MB/s | 4954731 | 76576 | 4.3× |
| Easyjson | 4030518 | 428.53 MB/s | 972032 | 5389 | 4.2× |
| JSONV2 | 4195620 | 411.67 MB/s | 1011616 | 7594 | 4.0× |
| Stdlib | 16902839 | 102.18 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1032 | 1755.96 MB/s | 0 | 0 | 14.2× |
| LightningDestructive | 1057 | 1714.29 MB/s | 0 | 0 | 13.8× |
| Easyjson | 2845 | 636.82 MB/s | 24 | 1 | 5.1× |
| Goccy | 3162 | 573.04 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6390 | 283.58 MB/s | 3348 | 38 | 2.3× |
| Sonic | 6641 | 272.85 MB/s | 3347 | 38 | 2.2× |
| JSONV2 | 7477 | 242.34 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8691 | 208.38 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14619 | 123.95 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1085 | 1670.45 MB/s | 0 | 0 | 13.4× |
| LightningDestructive | 1137 | 1593.60 MB/s | 0 | 0 | 12.8× |
| Easyjson | 2805 | 646.02 MB/s | 24 | 1 | 5.2× |
| Goccy | 3168 | 571.91 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6280 | 288.54 MB/s | 3349 | 38 | 2.3× |
| Sonic | 6433 | 281.67 MB/s | 3346 | 38 | 2.3× |
| JSONV2 | 7502 | 241.52 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 8599 | 210.60 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14507 | 124.90 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1272 | 1424.20 MB/s | 144 | 10 | 11.5× |
| LightningDestructive | 1323 | 1369.32 MB/s | 144 | 10 | 11.1× |
| Easyjson | 2964 | 611.42 MB/s | 144 | 10 | 4.9× |
| Goccy | 3351 | 540.74 MB/s | 2600 | 5 | 4.4× |
| SonicFastest | 6430 | 281.78 MB/s | 3366 | 40 | 2.3× |
| Sonic | 6743 | 268.73 MB/s | 3368 | 40 | 2.2× |
| JSONV2 | 7988 | 226.83 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8650 | 209.37 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14647 | 123.71 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 701 | 704.86 MB/s | 160 | 1 | 8.6× |
| LightningDestructive | 716 | 690.02 MB/s | 160 | 1 | 8.4× |
| SonicFastest | 1212 | 407.47 MB/s | 1075 | 8 | 5.0× |
| Sonic | 1213 | 407.34 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1392 | 354.21 MB/s | 1296 | 26 | 4.3× |
| Easyjson | 2330 | 212.00 MB/s | 448 | 3 | 2.6× |
| Goccy | 2440 | 202.48 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3154 | 156.61 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6029 | 81.93 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 429 | 535.77 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 435 | 529.12 MB/s | 160 | 1 | 9.8× |
| SonicFastest | 851 | 270.31 MB/s | 801 | 8 | 5.0× |
| Sonic | 854 | 269.19 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1156 | 198.14 MB/s | 1296 | 26 | 3.7× |
| Easyjson | 1546 | 148.76 MB/s | 448 | 3 | 2.8× |
| Goccy | 1670 | 137.73 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2391 | 96.20 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4252 | 54.09 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 83276 | 782.12 MB/s | 164881 | 105 | 7.8× |
| LightningDestructive | 88454 | 736.34 MB/s | 158661 | 100 | 7.3× |
| Sonic | 151120 | 430.99 MB/s | 235788 | 65 | 4.3× |
| SonicFastest | 152747 | 426.40 MB/s | 235884 | 65 | 4.3× |
| Goccy | 177274 | 367.41 MB/s | 228155 | 134 | 3.7× |
| LightningDecodeAny | 186364 | 286.16 MB/s | 180224 | 3245 | 3.5× |
| JSONV2 | 245560 | 265.24 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 649243 | 100.32 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2492949 | 778.38 MB/s | 2813904 | 1358 | 10.1× |
| Lightning | 2608632 | 743.87 MB/s | 2813906 | 1358 | 9.7× |
| Goccy | 4807630 | 403.62 MB/s | 4064228 | 13509 | 5.3× |
| Sonic | 6470974 | 299.87 MB/s | 4883825 | 1736 | 3.9× |
| SonicFastest | 6473506 | 299.76 MB/s | 4879378 | 1736 | 3.9× |
| Easyjson | 7660448 | 253.31 MB/s | 3871264 | 15043 | 3.3× |
| LightningDecodeAny | 9169201 | 211.63 MB/s | 7064790 | 218633 | 2.8× |
| JSONV2 | 11448595 | 169.49 MB/s | 3237192 | 13947 | 2.2× |
| Stdlib | 25259456 | 76.82 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 993795 | 3348.61 MB/s | 351704 | 1286 | 25.7× |
| Lightning | 1593309 | 2088.63 MB/s | 2488906 | 2995 | 16.0× |
| Sonic | 2472429 | 1345.98 MB/s | 5891368 | 4263 | 10.3× |
| SonicFastest | 2522551 | 1319.23 MB/s | 5892793 | 4263 | 10.1× |
| LightningDecodeAny | 3514398 | 874.62 MB/s | 4886621 | 56892 | 7.3× |
| Goccy | 6603203 | 503.97 MB/s | 3948910 | 3816 | 3.9× |
| JSONV2 | 8378886 | 397.17 MB/s | 5364512 | 13243 | 3.0× |
| Stdlib | 25549315 | 130.25 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225048 | 979.11 MB/s | 135872 | 226 | 9.9× |
| LightningDestructive | 235784 | 934.52 MB/s | 135872 | 226 | 9.5× |
| Goccy | 464181 | 474.70 MB/s | 363810 | 1066 | 4.8× |
| Sonic | 512634 | 429.83 MB/s | 351195 | 262 | 4.4× |
| SonicFastest | 514887 | 427.95 MB/s | 350786 | 262 | 4.3× |
| Easyjson | 593589 | 371.21 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 680972 | 323.58 MB/s | 129747 | 470 | 3.3× |
| LightningDecodeAny | 956473 | 113.24 MB/s | 897522 | 11703 | 2.3× |
| Stdlib | 2234418 | 98.61 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 13140336 | 616.43 MB/s | 15730641 | 20821 | 7.8× |
| Lightning | 13771109 | 588.19 MB/s | 15730641 | 20821 | 7.4× |
| SonicFastest | 19830524 | 408.46 MB/s | 19856701 | 41640 | 5.2× |
| Sonic | 20007048 | 404.86 MB/s | 19856328 | 41640 | 5.1× |
| Goccy | 26461310 | 306.11 MB/s | 19077720 | 107156 | 3.9× |
| Easyjson | 34348365 | 235.82 MB/s | 15059621 | 41643 | 3.0× |
| LightningDecodeAny | 38857967 | 133.90 MB/s | 46191121 | 747112 | 2.6× |
| JSONV2 | 47226400 | 171.52 MB/s | 15233722 | 78972 | 2.2× |
| Stdlib | 102209908 | 79.25 MB/s | 15665071 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5621928 | 530.68 MB/s | 3908872 | 29356 | 9.0× |
| Lightning | 5826698 | 512.03 MB/s | 3908876 | 29356 | 8.7× |
| SonicFastest | 9522316 | 313.31 MB/s | 9130119 | 57804 | 5.3× |
| Sonic | 9575164 | 311.58 MB/s | 9131327 | 57804 | 5.3× |
| LightningDecodeAny | 17732943 | 103.43 MB/s | 23982396 | 351152 | 2.9× |
| Goccy | 18124026 | 164.61 MB/s | 9815735 | 273617 | 2.8× |
| Easyjson | 18158284 | 164.30 MB/s | 9479441 | 30115 | 2.8× |
| JSONV2 | 24774339 | 120.43 MB/s | 9257069 | 86278 | 2.0× |
| Stdlib | 50614410 | 58.94 MB/s | 9258086 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1299096 | 557.00 MB/s | 907600 | 3618 | 9.3× |
| Lightning | 1369412 | 528.40 MB/s | 907595 | 3618 | 8.8× |
| SonicFastest | 2168770 | 333.64 MB/s | 2367503 | 3683 | 5.6× |
| Sonic | 2194226 | 329.77 MB/s | 2367415 | 3683 | 5.5× |
| Easyjson | 5141183 | 140.75 MB/s | 2847906 | 3698 | 2.4× |
| LightningDecodeAny | 5313517 | 122.44 MB/s | 6500457 | 76546 | 2.3× |
| Goccy | 5418188 | 133.55 MB/s | 2702405 | 80266 | 2.2× |
| JSONV2 | 6521686 | 110.95 MB/s | 2704701 | 7318 | 1.9× |
| Stdlib | 12102847 | 59.79 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2054977 | 767.58 MB/s | 907594 | 3618 | 8.7× |
| LightningDestructive | 2070511 | 761.82 MB/s | 907600 | 3618 | 8.7× |
| Sonic | 2614646 | 603.28 MB/s | 3233533 | 3683 | 6.9× |
| SonicFastest | 2746438 | 574.33 MB/s | 3230668 | 3683 | 6.5× |
| LightningDecodeAny | 4637083 | 162.47 MB/s | 6500459 | 76546 | 3.9× |
| Easyjson | 6281161 | 251.12 MB/s | 2847905 | 3698 | 2.9× |
| Goccy | 6585577 | 239.52 MB/s | 3487393 | 80261 | 2.7× |
| JSONV2 | 6875803 | 229.41 MB/s | 2704552 | 7318 | 2.6× |
| Stdlib | 17927119 | 87.99 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 218709 | 686.41 MB/s | 81920 | 1 | 10.2× |
| LightningDestructive | 227202 | 660.75 MB/s | 81920 | 1 | 9.8× |
| Sonic | 387728 | 387.19 MB/s | 407420 | 16 | 5.7× |
| SonicFastest | 388273 | 386.65 MB/s | 407451 | 16 | 5.7× |
| LightningDecodeAny | 564932 | 265.73 MB/s | 745765 | 10016 | 3.9× |
| Goccy | 1033829 | 145.21 MB/s | 324807 | 10005 | 2.2× |
| JSONV2 | 1132932 | 132.51 MB/s | 357728 | 20 | 2.0× |
| Stdlib | 2227294 | 67.40 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 31334 | 897.32 MB/s | 29216 | 103 | 9.9× |
| LightningDestructive | 32569 | 863.31 MB/s | 29088 | 101 | 9.5× |
| Sonic | 67449 | 416.87 MB/s | 59455 | 83 | 4.6× |
| SonicFastest | 67615 | 415.84 MB/s | 59439 | 83 | 4.6× |
| Easyjson | 75070 | 374.54 MB/s | 32304 | 138 | 4.1× |
| Goccy | 75862 | 370.63 MB/s | 59274 | 188 | 4.1× |
| JSONV2 | 126593 | 222.11 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 154464 | 182.03 MB/s | 140592 | 2643 | 2.0× |
| Stdlib | 309659 | 90.80 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1968 | 1182.97 MB/s | 32 | 1 | 12.0× |
| LightningDestructive | 1991 | 1169.14 MB/s | 32 | 1 | 11.9× |
| Goccy | 4496 | 517.78 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 5221 | 445.88 MB/s | 192 | 2 | 4.5× |
| SonicFastest | 6086 | 382.52 MB/s | 3715 | 4 | 3.9× |
| Sonic | 6102 | 381.50 MB/s | 3713 | 4 | 3.9× |
| JSONV2 | 7757 | 300.10 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 9697 | 173.77 MB/s | 10200 | 195 | 2.4× |
| Stdlib | 23633 | 98.51 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220 | 860.29 MB/s | 0 | 0 | 11.2× |
| LightningDestructive | 224 | 844.81 MB/s | 0 | 0 | 11.0× |
| Goccy | 393 | 480.76 MB/s | 304 | 2 | 6.2× |
| Easyjson | 567 | 333.18 MB/s | 0 | 0 | 4.3× |
| SonicFastest | 742 | 254.70 MB/s | 341 | 3 | 3.3× |
| Sonic | 743 | 254.48 MB/s | 341 | 3 | 3.3× |
| JSONV2 | 929 | 203.44 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1193 | 112.37 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2455 | 76.98 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1457 | 1504.18 MB/s | 0 | 0 | 11.9× |
| LightningDestructive | 1496 | 1464.71 MB/s | 0 | 0 | 11.5× |
| Goccy | 3429 | 639.00 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3472 | 631.01 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6458 | 339.26 MB/s | 3600 | 38 | 2.7× |
| Sonic | 6773 | 323.49 MB/s | 3603 | 38 | 2.6× |
| JSONV2 | 7645 | 286.59 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8445 | 214.45 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 17272 | 126.85 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 668505 | 763.61 MB/s | 703298 | 1010 | 9.6× |
| Lightning | 712689 | 716.27 MB/s | 703299 | 1010 | 9.0× |
| Goccy | 1262494 | 404.34 MB/s | 1134327 | 5006 | 5.1× |
| SonicFastest | 1378911 | 370.20 MB/s | 1310010 | 2014 | 4.7× |
| Sonic | 1397477 | 365.28 MB/s | 1310914 | 2014 | 4.6× |
| Easyjson | 1562099 | 326.79 MB/s | 863783 | 3012 | 4.1× |
| JSONV2 | 3119364 | 163.65 MB/s | 1075952 | 12645 | 2.1× |
| LightningDecodeAny | 3409959 | 135.33 MB/s | 2929689 | 64018 | 1.9× |
| Stdlib | 6439787 | 79.27 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 614 | 32244.04 MB/s | 0 | 0 | 232.3× |
| LightningDestructive | 882 | 22437.50 MB/s | 0 | 0 | 161.7× |
| SonicFastest | 7876 | 2512.50 MB/s | 21003 | 3 | 18.1× |
| Goccy | 26836 | 737.40 MB/s | 20492 | 2 | 5.3× |
| Sonic | 29242 | 676.73 MB/s | 20616 | 3 | 4.9× |
| JSONV2 | 36069 | 548.64 MB/s | 8 | 1 | 4.0× |
| LightningDecodeAny | 90911 | 217.66 MB/s | 116864 | 2015 | 1.6× |
| Easyjson | 112009 | 176.67 MB/s | 0 | 0 | 1.3× |
| Stdlib | 142585 | 138.79 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2219 | 8167.94 MB/s | 432 | 2 | 54.6× |
| LightningDestructive | 2373 | 7636.09 MB/s | 0 | 0 | 51.0× |
| Easyjson | 4687 | 3867.19 MB/s | 432 | 2 | 25.8× |
| Sonic | 11067 | 1637.72 MB/s | 20407 | 5 | 10.9× |
| SonicFastest | 11135 | 1627.70 MB/s | 20414 | 5 | 10.9× |
| LightningDecodeAny | 18557 | 963.64 MB/s | 29088 | 191 | 6.5× |
| Goccy | 25641 | 706.85 MB/s | 19460 | 2 | 4.7× |
| JSONV2 | 50516 | 358.78 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 121048 | 149.73 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2323676 | 864.36 MB/s | 3089821 | 6821 | 9.0× |
| Lightning | 2440063 | 823.13 MB/s | 3091534 | 6827 | 8.5× |
| Goccy | 4823137 | 416.43 MB/s | 5410772 | 15833 | 4.3× |
| Sonic | 5142727 | 390.55 MB/s | 5154893 | 7085 | 4.0× |
| SonicFastest | 5227394 | 384.22 MB/s | 5153242 | 7085 | 4.0× |
| Easyjson | 5382809 | 373.13 MB/s | 2981489 | 7439 | 3.9× |
| LightningDecodeAny | 6918243 | 165.11 MB/s | 8498329 | 134008 | 3.0× |
| JSONV2 | 6948395 | 289.06 MB/s | 3173679 | 14563 | 3.0× |
| Stdlib | 20825143 | 96.45 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 817 | 671.96 MB/s | 480 | 1 | 7.2× |
| LightningDestructive | 824 | 666.29 MB/s | 480 | 1 | 7.2× |
| LightningDecodeAny | 1716 | 319.36 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 1978 | 277.62 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2220 | 247.31 MB/s | 2263 | 8 | 2.7× |
| Sonic | 2233 | 245.91 MB/s | 2262 | 8 | 2.6× |
| Goccy | 2987 | 183.78 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3036 | 180.86 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 5897 | 93.11 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 492448 | 1282.40 MB/s | 402728 | 545 | 12.1× |
| Lightning | 556918 | 1133.95 MB/s | 451257 | 857 | 10.7× |
| Sonic | 1098566 | 574.85 MB/s | 1068587 | 814 | 5.4× |
| SonicFastest | 1110424 | 568.71 MB/s | 1070792 | 814 | 5.4× |
| Easyjson | 1274569 | 495.47 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1293649 | 488.16 MB/s | 992420 | 1201 | 4.6× |
| JSONV2 | 2117774 | 298.20 MB/s | 571596 | 3144 | 2.8× |
| LightningDecodeAny | 2399828 | 194.56 MB/s | 2077366 | 30126 | 2.5× |
| Stdlib | 5971379 | 105.76 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 703826 | 799.07 MB/s | 579337 | 429 | 8.1× |
| Lightning | 911086 | 617.29 MB/s | 802710 | 1235 | 6.2× |
| SonicFastest | 1405670 | 400.10 MB/s | 1350447 | 1185 | 4.0× |
| Sonic | 1407427 | 399.60 MB/s | 1350802 | 1185 | 4.0× |
| Goccy | 1562376 | 359.97 MB/s | 1046763 | 1029 | 3.6× |
| Easyjson | 2023289 | 277.97 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 2878586 | 195.38 MB/s | 2181319 | 30126 | 2.0× |
| JSONV2 | 3034149 | 185.36 MB/s | 927406 | 3482 | 1.9× |
| Stdlib | 5669581 | 99.20 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 582240 | 915.74 MB/s | 333416 | 2084 | 10.2× |
| Lightning | 659596 | 808.34 MB/s | 368224 | 2293 | 9.0× |
| Easyjson | 1259373 | 423.37 MB/s | 428362 | 3273 | 4.7× |
| Sonic | 1374199 | 387.99 MB/s | 981252 | 3082 | 4.3× |
| SonicFastest | 1375490 | 387.63 MB/s | 980966 | 3082 | 4.3× |
| Goccy | 1408201 | 378.62 MB/s | 1167060 | 5408 | 4.2× |
| JSONV2 | 2613201 | 204.03 MB/s | 745424 | 13288 | 2.3× |
| LightningDecodeAny | 3406567 | 156.51 MB/s | 2991145 | 50076 | 1.7× |
| Stdlib | 5938637 | 89.78 MB/s | 798693 | 17133 | 1.0× |
