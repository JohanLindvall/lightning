# JSON Deserialization Benchmarks

- generated 2026-07-02T12:28:59Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105287 | 1208.84 MB/s | 49760 | 3 | 12.5× |
| LightningDestructive | 106681 | 1193.05 MB/s | 49280 | 2 | 12.3× |
| SonicFastest | 197505 | 644.41 MB/s | 214267 | 15 | 6.7× |
| Sonic | 198655 | 640.68 MB/s | 214454 | 15 | 6.6× |
| Goccy | 239298 | 531.87 MB/s | 225064 | 884 | 5.5× |
| Easyjson | 247774 | 513.67 MB/s | 122864 | 14 | 5.3× |
| JSONV2 | 455711 | 279.29 MB/s | 195127 | 1805 | 2.9× |
| LightningDecodeAny | 457311 | 206.98 MB/s | 465731 | 9708 | 2.9× |
| Stdlib | 1314795 | 96.80 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4032573 | 558.22 MB/s | 3008241 | 1158 | 8.4× |
| Lightning | 4183545 | 538.07 MB/s | 3008242 | 1158 | 8.1× |
| Sonic | 4851286 | 464.01 MB/s | 4867796 | 2584 | 7.0× |
| SonicFastest | 5012398 | 449.10 MB/s | 4869193 | 2584 | 6.8× |
| Goccy | 13278192 | 169.53 MB/s | 4202055 | 56535 | 2.6× |
| LightningDecodeAny | 13296971 | 169.29 MB/s | 19380210 | 223896 | 2.6× |
| Easyjson | 13527999 | 166.40 MB/s | 3099809 | 2120 | 2.5× |
| JSONV2 | 17251500 | 130.48 MB/s | 3123189 | 3083 | 2.0× |
| Stdlib | 34064163 | 66.08 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 545531 | 495.67 MB/s | 392945 | 568 | 8.0× |
| LightningDestructive | 560226 | 482.67 MB/s | 392945 | 568 | 7.8× |
| Sonic | 754332 | 358.47 MB/s | 641580 | 1147 | 5.8× |
| SonicFastest | 775579 | 348.65 MB/s | 642663 | 1147 | 5.6× |
| Easyjson | 1765104 | 153.19 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1820566 | 148.53 MB/s | 540879 | 8121 | 2.4× |
| LightningDecodeAny | 2170376 | 124.59 MB/s | 2543877 | 29687 | 2.0× |
| JSONV2 | 2310692 | 117.02 MB/s | 348160 | 1628 | 1.9× |
| Stdlib | 4369242 | 61.89 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1154448 | 1496.13 MB/s | 767864 | 2798 | 15.3× |
| Lightning | 1199735 | 1439.66 MB/s | 767905 | 2799 | 14.7× |
| SonicFastest | 2130605 | 810.66 MB/s | 2695306 | 5547 | 8.3× |
| Sonic | 2138291 | 807.75 MB/s | 2695960 | 5547 | 8.2× |
| Goccy | 2906176 | 594.32 MB/s | 2581132 | 14603 | 6.1× |
| LightningDecodeAny | 4101467 | 121.98 MB/s | 4954731 | 76576 | 4.3× |
| Easyjson | 4262085 | 405.25 MB/s | 972032 | 5389 | 4.1× |
| JSONV2 | 4997432 | 345.62 MB/s | 1011614 | 7594 | 3.5× |
| Stdlib | 17615572 | 98.05 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1012 | 1790.40 MB/s | 0 | 0 | 15.9× |
| LightningDestructive | 1059 | 1711.56 MB/s | 0 | 0 | 15.2× |
| Easyjson | 2993 | 605.36 MB/s | 24 | 1 | 5.4× |
| Goccy | 3181 | 569.70 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 6313 | 287.02 MB/s | 3347 | 38 | 2.5× |
| Sonic | 6494 | 279.02 MB/s | 3345 | 38 | 2.5× |
| JSONV2 | 7999 | 226.53 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9399 | 192.68 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16071 | 112.75 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1098 | 1650.21 MB/s | 0 | 0 | 14.8× |
| LightningDestructive | 1147 | 1580.28 MB/s | 0 | 0 | 14.1× |
| Easyjson | 2968 | 610.44 MB/s | 24 | 1 | 5.5× |
| Goccy | 3178 | 570.09 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5988 | 302.60 MB/s | 3345 | 38 | 2.7× |
| Sonic | 6219 | 291.39 MB/s | 3344 | 38 | 2.6× |
| JSONV2 | 8270 | 219.11 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9242 | 195.95 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 16226 | 111.67 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1358 | 1334.21 MB/s | 144 | 10 | 11.9× |
| LightningDestructive | 1435 | 1262.31 MB/s | 144 | 10 | 11.2× |
| Easyjson | 3158 | 573.76 MB/s | 144 | 10 | 5.1× |
| Goccy | 3544 | 511.33 MB/s | 2600 | 5 | 4.5× |
| SonicFastest | 6329 | 286.30 MB/s | 3364 | 40 | 2.5× |
| Sonic | 6504 | 278.61 MB/s | 3368 | 40 | 2.5× |
| JSONV2 | 8468 | 213.99 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9217 | 196.49 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16100 | 112.55 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 710 | 695.52 MB/s | 160 | 1 | 9.8× |
| LightningDestructive | 724 | 682.05 MB/s | 160 | 1 | 9.6× |
| SonicFastest | 1290 | 382.99 MB/s | 1076 | 8 | 5.4× |
| Sonic | 1297 | 380.83 MB/s | 1077 | 8 | 5.3× |
| LightningDecodeAny | 1504 | 327.72 MB/s | 1296 | 26 | 4.6× |
| Easyjson | 2485 | 198.79 MB/s | 448 | 3 | 2.8× |
| Goccy | 2673 | 184.80 MB/s | 856 | 23 | 2.6× |
| JSONV2 | 3662 | 134.91 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6938 | 71.21 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 455 | 505.02 MB/s | 160 | 1 | 10.8× |
| LightningDestructive | 455 | 505.00 MB/s | 160 | 1 | 10.8× |
| SonicFastest | 939 | 244.87 MB/s | 801 | 8 | 5.3× |
| Sonic | 940 | 244.61 MB/s | 801 | 8 | 5.3× |
| LightningDecodeAny | 1294 | 176.95 MB/s | 1296 | 26 | 3.8× |
| Easyjson | 1681 | 136.86 MB/s | 448 | 3 | 2.9× |
| Goccy | 1760 | 130.68 MB/s | 584 | 23 | 2.8× |
| JSONV2 | 2680 | 85.82 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4939 | 46.57 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 85728 | 759.75 MB/s | 164881 | 105 | 8.0× |
| LightningDestructive | 91207 | 714.11 MB/s | 158661 | 100 | 7.5× |
| Sonic | 147617 | 441.22 MB/s | 235824 | 65 | 4.6× |
| SonicFastest | 158755 | 410.27 MB/s | 236048 | 65 | 4.3× |
| Goccy | 195384 | 333.35 MB/s | 228896 | 134 | 3.5× |
| LightningDecodeAny | 203523 | 262.03 MB/s | 180225 | 3245 | 3.4× |
| JSONV2 | 272318 | 239.18 MB/s | 206665 | 607 | 2.5× |
| Stdlib | 682766 | 95.39 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2452211 | 791.32 MB/s | 2813904 | 1358 | 11.4× |
| Lightning | 2525581 | 768.33 MB/s | 2813905 | 1358 | 11.1× |
| Sonic | 4821668 | 402.45 MB/s | 4879382 | 1736 | 5.8× |
| SonicFastest | 4833131 | 401.49 MB/s | 4877930 | 1736 | 5.8× |
| Goccy | 4959759 | 391.24 MB/s | 4062348 | 13509 | 5.6× |
| Easyjson | 8074950 | 240.31 MB/s | 3871265 | 15043 | 3.5× |
| LightningDecodeAny | 9558339 | 203.01 MB/s | 7064789 | 218633 | 2.9× |
| JSONV2 | 12567174 | 154.41 MB/s | 3237188 | 13947 | 2.2× |
| Stdlib | 27955423 | 69.41 MB/s | 3551315 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1060105 | 3139.15 MB/s | 351704 | 1286 | 22.4× |
| Lightning | 1602844 | 2076.20 MB/s | 2488907 | 2995 | 14.8× |
| SonicFastest | 2135880 | 1558.06 MB/s | 5894654 | 4263 | 11.1× |
| Sonic | 2148131 | 1549.18 MB/s | 5895001 | 4263 | 11.1× |
| LightningDecodeAny | 3488813 | 881.03 MB/s | 4886621 | 56892 | 6.8× |
| Goccy | 6097615 | 545.76 MB/s | 3948914 | 3817 | 3.9× |
| JSONV2 | 7748973 | 429.45 MB/s | 5364504 | 13243 | 3.1× |
| Stdlib | 23758399 | 140.07 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 226190 | 974.16 MB/s | 135872 | 226 | 10.7× |
| LightningDestructive | 237305 | 928.53 MB/s | 135872 | 226 | 10.2× |
| Sonic | 485276 | 454.06 MB/s | 351110 | 262 | 5.0× |
| SonicFastest | 487053 | 452.41 MB/s | 351218 | 262 | 5.0× |
| Goccy | 490867 | 448.89 MB/s | 363710 | 1066 | 4.9× |
| Easyjson | 647904 | 340.09 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 753138 | 292.57 MB/s | 129747 | 470 | 3.2× |
| LightningDecodeAny | 1041817 | 103.97 MB/s | 897522 | 11703 | 2.3× |
| Stdlib | 2429437 | 90.70 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 13597242 | 595.71 MB/s | 15730640 | 20821 | 8.5× |
| Lightning | 14508894 | 558.28 MB/s | 15730640 | 20821 | 7.9× |
| SonicFastest | 18182249 | 445.49 MB/s | 19861771 | 41640 | 6.3× |
| Sonic | 18184125 | 445.45 MB/s | 19860563 | 41640 | 6.3× |
| Goccy | 27353658 | 296.12 MB/s | 18995810 | 107155 | 4.2× |
| Easyjson | 34994956 | 231.46 MB/s | 15059618 | 41643 | 3.3× |
| LightningDecodeAny | 39777948 | 130.80 MB/s | 46191121 | 747112 | 2.9× |
| JSONV2 | 50580658 | 160.14 MB/s | 15233698 | 78972 | 2.3× |
| Stdlib | 114991672 | 70.44 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5711335 | 522.38 MB/s | 3908872 | 29356 | 10.1× |
| Lightning | 5862880 | 508.87 MB/s | 3908873 | 29356 | 9.8× |
| SonicFastest | 9393339 | 317.62 MB/s | 9129841 | 57804 | 6.1× |
| Sonic | 9394338 | 317.58 MB/s | 9130976 | 57804 | 6.1× |
| LightningDecodeAny | 18609026 | 98.56 MB/s | 23982392 | 351152 | 3.1× |
| Goccy | 19198252 | 155.40 MB/s | 9899279 | 273621 | 3.0× |
| Easyjson | 19673383 | 151.65 MB/s | 9479440 | 30115 | 2.9× |
| JSONV2 | 28641385 | 104.17 MB/s | 9257069 | 86278 | 2.0× |
| Stdlib | 57556881 | 51.84 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1343018 | 538.78 MB/s | 907601 | 3618 | 10.5× |
| Lightning | 1419938 | 509.60 MB/s | 907595 | 3618 | 10.0× |
| SonicFastest | 2169037 | 333.60 MB/s | 2372476 | 3683 | 6.5× |
| Sonic | 2186683 | 330.91 MB/s | 2372500 | 3683 | 6.5× |
| Easyjson | 5405692 | 133.86 MB/s | 2847907 | 3698 | 2.6× |
| LightningDecodeAny | 5478211 | 118.76 MB/s | 6500457 | 76546 | 2.6× |
| Goccy | 5663583 | 127.76 MB/s | 2745146 | 80269 | 2.5× |
| JSONV2 | 6981392 | 103.65 MB/s | 2704709 | 7318 | 2.0× |
| Stdlib | 14141308 | 51.17 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1981794 | 795.92 MB/s | 907600 | 3618 | 10.0× |
| Lightning | 1995175 | 790.58 MB/s | 907595 | 3618 | 9.9× |
| SonicFastest | 2438144 | 646.95 MB/s | 3227081 | 3683 | 8.1× |
| Sonic | 2465223 | 639.84 MB/s | 3229019 | 3683 | 8.0× |
| LightningDecodeAny | 4809985 | 156.63 MB/s | 6500455 | 76546 | 4.1× |
| Easyjson | 6367795 | 247.71 MB/s | 2847905 | 3698 | 3.1× |
| Goccy | 6719578 | 234.74 MB/s | 3476138 | 80261 | 2.9× |
| JSONV2 | 7373631 | 213.92 MB/s | 2704556 | 7318 | 2.7× |
| Stdlib | 19740715 | 79.90 MB/s | 2704547 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 231485 | 648.52 MB/s | 81920 | 1 | 9.2× |
| LightningDestructive | 239316 | 627.30 MB/s | 81920 | 1 | 8.9× |
| SonicFastest | 379982 | 395.08 MB/s | 407503 | 16 | 5.6× |
| Sonic | 380401 | 394.65 MB/s | 407507 | 16 | 5.6× |
| LightningDecodeAny | 570893 | 262.96 MB/s | 745766 | 10016 | 3.7× |
| Goccy | 1024179 | 146.58 MB/s | 325831 | 10005 | 2.1× |
| JSONV2 | 1092524 | 137.41 MB/s | 357724 | 20 | 2.0× |
| Stdlib | 2139694 | 70.16 MB/s | 357802 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32933 | 853.77 MB/s | 29216 | 103 | 10.3× |
| LightningDestructive | 33775 | 832.48 MB/s | 29088 | 101 | 10.0× |
| SonicFastest | 58666 | 479.27 MB/s | 59496 | 83 | 5.8× |
| Sonic | 58905 | 477.33 MB/s | 59509 | 83 | 5.8× |
| Goccy | 78985 | 355.98 MB/s | 59279 | 188 | 4.3× |
| Easyjson | 80743 | 348.23 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 138707 | 202.71 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 166416 | 168.96 MB/s | 140592 | 2643 | 2.0× |
| Stdlib | 339274 | 82.87 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2031 | 1146.10 MB/s | 32 | 1 | 12.8× |
| LightningDestructive | 2074 | 1122.21 MB/s | 32 | 1 | 12.6× |
| Goccy | 4824 | 482.60 MB/s | 3649 | 4 | 5.4× |
| SonicFastest | 4904 | 474.73 MB/s | 3713 | 4 | 5.3× |
| Sonic | 4906 | 474.53 MB/s | 3710 | 4 | 5.3× |
| Easyjson | 5483 | 424.61 MB/s | 192 | 2 | 4.8× |
| JSONV2 | 8687 | 267.99 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 10708 | 157.36 MB/s | 10200 | 195 | 2.4× |
| Stdlib | 26050 | 89.37 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 213 | 886.62 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 223 | 847.36 MB/s | 0 | 0 | 12.5× |
| Goccy | 429 | 440.38 MB/s | 304 | 2 | 6.5× |
| Easyjson | 600 | 315.27 MB/s | 0 | 0 | 4.6× |
| SonicFastest | 632 | 298.84 MB/s | 341 | 3 | 4.4× |
| Sonic | 635 | 297.68 MB/s | 341 | 3 | 4.4× |
| JSONV2 | 1019 | 185.54 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1373 | 97.57 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2777 | 68.05 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1467 | 1493.65 MB/s | 0 | 0 | 12.8× |
| LightningDestructive | 1539 | 1423.65 MB/s | 0 | 0 | 12.2× |
| Goccy | 3564 | 614.75 MB/s | 2864 | 4 | 5.3× |
| Easyjson | 3895 | 562.46 MB/s | 24 | 1 | 4.8× |
| SonicFastest | 6572 | 333.39 MB/s | 3599 | 38 | 2.9× |
| Sonic | 6854 | 319.68 MB/s | 3597 | 38 | 2.7× |
| JSONV2 | 8428 | 259.97 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9242 | 195.95 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18840 | 116.30 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 713354 | 715.60 MB/s | 703300 | 1010 | 9.9× |
| Lightning | 746234 | 684.07 MB/s | 703298 | 1010 | 9.4× |
| Sonic | 1263577 | 403.99 MB/s | 1306696 | 2014 | 5.6× |
| SonicFastest | 1287384 | 396.52 MB/s | 1308813 | 2014 | 5.5× |
| Goccy | 1301149 | 392.33 MB/s | 1136972 | 5006 | 5.4× |
| Easyjson | 1684329 | 303.07 MB/s | 863778 | 3012 | 4.2× |
| JSONV2 | 3476524 | 146.84 MB/s | 1075954 | 12645 | 2.0× |
| LightningDecodeAny | 3556552 | 129.75 MB/s | 2929688 | 64018 | 2.0× |
| Stdlib | 7037966 | 72.53 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 550 | 35952.03 MB/s | 0 | 0 | 295.6× |
| LightningDestructive | 841 | 23519.31 MB/s | 0 | 0 | 193.4× |
| SonicFastest | 6683 | 2961.28 MB/s | 21142 | 3 | 24.3× |
| Goccy | 29433 | 672.34 MB/s | 20492 | 2 | 5.5× |
| Sonic | 31859 | 621.14 MB/s | 20608 | 3 | 5.1× |
| JSONV2 | 33288 | 594.48 MB/s | 8 | 1 | 4.9× |
| LightningDecodeAny | 96135 | 205.84 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 100967 | 196.00 MB/s | 0 | 0 | 1.6× |
| Stdlib | 162714 | 121.62 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2286 | 7927.28 MB/s | 432 | 2 | 52.7× |
| LightningDestructive | 2592 | 6992.74 MB/s | 0 | 0 | 46.4× |
| Easyjson | 4999 | 3625.36 MB/s | 432 | 2 | 24.1× |
| SonicFastest | 8390 | 2160.26 MB/s | 20461 | 5 | 14.3× |
| Sonic | 8498 | 2132.83 MB/s | 20434 | 5 | 14.2× |
| LightningDecodeAny | 19469 | 918.49 MB/s | 29088 | 191 | 6.2× |
| Goccy | 23292 | 778.13 MB/s | 19460 | 2 | 5.2× |
| JSONV2 | 51075 | 354.85 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 120380 | 150.56 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2366721 | 848.64 MB/s | 3089821 | 6821 | 9.4× |
| Lightning | 2562931 | 783.67 MB/s | 3091533 | 6827 | 8.7× |
| SonicFastest | 4220361 | 475.91 MB/s | 5163199 | 7085 | 5.3× |
| Sonic | 4348894 | 461.84 MB/s | 5160341 | 7085 | 5.1× |
| Goccy | 5013382 | 400.63 MB/s | 5410504 | 15831 | 4.4× |
| Easyjson | 5883080 | 341.40 MB/s | 2981486 | 7439 | 3.8× |
| LightningDecodeAny | 7246132 | 157.64 MB/s | 8498329 | 134008 | 3.1× |
| JSONV2 | 7584601 | 264.81 MB/s | 3173687 | 14563 | 2.9× |
| Stdlib | 22303884 | 90.05 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 852 | 643.99 MB/s | 480 | 1 | 8.0× |
| LightningDestructive | 858 | 639.50 MB/s | 480 | 1 | 8.0× |
| LightningDecodeAny | 1819 | 301.22 MB/s | 2021 | 46 | 3.8× |
| Easyjson | 2231 | 246.04 MB/s | 1616 | 5 | 3.1× |
| SonicFastest | 2428 | 226.13 MB/s | 2261 | 8 | 2.8× |
| Sonic | 2451 | 224.00 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3327 | 165.00 MB/s | 2129 | 43 | 2.1× |
| JSONV2 | 3434 | 159.86 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 6828 | 80.41 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 522686 | 1208.21 MB/s | 402728 | 545 | 12.4× |
| Lightning | 607549 | 1039.45 MB/s | 451257 | 857 | 10.7× |
| SonicFastest | 1117526 | 565.10 MB/s | 1068493 | 814 | 5.8× |
| Sonic | 1120634 | 563.53 MB/s | 1068226 | 814 | 5.8× |
| Goccy | 1306915 | 483.21 MB/s | 987502 | 1200 | 5.0× |
| Easyjson | 1405576 | 449.29 MB/s | 422504 | 936 | 4.6× |
| JSONV2 | 2407123 | 262.35 MB/s | 571593 | 3144 | 2.7× |
| LightningDecodeAny | 2689805 | 173.58 MB/s | 2077365 | 30126 | 2.4× |
| Stdlib | 6476688 | 97.51 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 757705 | 742.25 MB/s | 579336 | 429 | 8.0× |
| Lightning | 980320 | 573.70 MB/s | 802709 | 1235 | 6.2× |
| SonicFastest | 1321605 | 425.55 MB/s | 1349861 | 1185 | 4.6× |
| Sonic | 1334637 | 421.39 MB/s | 1349271 | 1185 | 4.6× |
| Goccy | 1545258 | 363.96 MB/s | 1040868 | 1028 | 3.9× |
| Easyjson | 2231145 | 252.07 MB/s | 775153 | 1254 | 2.7× |
| LightningDecodeAny | 3135963 | 179.34 MB/s | 2181319 | 30126 | 1.9× |
| JSONV2 | 3248739 | 173.12 MB/s | 927409 | 3482 | 1.9× |
| Stdlib | 6077211 | 92.54 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 631568 | 844.21 MB/s | 333416 | 2084 | 10.0× |
| Lightning | 717950 | 742.64 MB/s | 368224 | 2293 | 8.8× |
| SonicFastest | 1114554 | 478.38 MB/s | 979827 | 3082 | 5.7× |
| Sonic | 1120539 | 475.82 MB/s | 979445 | 3082 | 5.6× |
| Easyjson | 1314702 | 405.55 MB/s | 428362 | 3273 | 4.8× |
| Goccy | 1429112 | 373.08 MB/s | 1167071 | 5408 | 4.4× |
| JSONV2 | 2845918 | 187.35 MB/s | 745423 | 13288 | 2.2× |
| LightningDecodeAny | 3602859 | 147.99 MB/s | 2991145 | 50076 | 1.7× |
| Stdlib | 6298467 | 84.65 MB/s | 798692 | 17133 | 1.0× |
