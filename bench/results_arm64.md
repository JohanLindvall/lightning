# JSON Deserialization Benchmarks

- generated 2026-09-21T08:22:47Z
- go version go1.26.8 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 84696 | 1502.73 MB/s | 49819 | 2 | 12.9× |
| Lightning | 84720 | 1502.30 MB/s | 49818 | 2 | 12.9× |
| LightningDestructive | 84930 | 1498.58 MB/s | 49280 | 2 | 12.9× |
| Sonic | 185312 | 686.81 MB/s | 199347 | 10 | 5.9× |
| SonicFastest | 186861 | 681.12 MB/s | 201568 | 10 | 5.9× |
| Goccy | 199425 | 638.21 MB/s | 224850 | 884 | 5.5× |
| Easyjson | 213805 | 595.29 MB/s | 122864 | 14 | 5.1× |
| JSONV2 | 421829 | 301.72 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 438628 | 215.79 MB/s | 463881 | 9706 | 2.5× |
| Stdlib | 1094022 | 116.34 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2659745 | 846.34 MB/s | 2532848 | 1143 | 10.0× |
| LightningArena | 2695097 | 835.24 MB/s | 2532848 | 1143 | 9.9× |
| Lightning | 2699844 | 833.77 MB/s | 2532850 | 1143 | 9.8× |
| Sonic | 4538171 | 496.03 MB/s | 15233795 | 970 | 5.9× |
| SonicFastest | 4540418 | 495.78 MB/s | 15233749 | 970 | 5.8× |
| LightningDecodeAny | 8512765 | 264.43 MB/s | 6828999 | 223498 | 3.1× |
| Goccy | 10183108 | 221.06 MB/s | 4108020 | 56532 | 2.6× |
| Easyjson | 10989365 | 204.84 MB/s | 3099808 | 2120 | 2.4× |
| JSONV2 | 16220158 | 138.78 MB/s | 3123213 | 3083 | 1.6× |
| Stdlib | 26557306 | 84.76 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 371721 | 727.43 MB/s | 397296 | 567 | 9.2× |
| Lightning | 372579 | 725.76 MB/s | 397296 | 567 | 9.2× |
| LightningArena | 373624 | 723.73 MB/s | 397296 | 567 | 9.2× |
| SonicFastest | 643393 | 420.28 MB/s | 481709 | 968 | 5.3× |
| Sonic | 645649 | 418.81 MB/s | 479089 | 968 | 5.3× |
| LightningDecodeAny | 1218801 | 221.86 MB/s | 845691 | 29656 | 2.8× |
| Goccy | 1398132 | 193.40 MB/s | 542384 | 8122 | 2.5× |
| Easyjson | 1418103 | 190.68 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2096426 | 128.98 MB/s | 348155 | 1628 | 1.6× |
| Stdlib | 3431747 | 78.79 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 966739 | 1786.63 MB/s | 765560 | 2798 | 13.7× |
| Lightning | 978724 | 1764.75 MB/s | 768313 | 2798 | 13.5× |
| LightningArena | 983178 | 1756.76 MB/s | 775563 | 2444 | 13.4× |
| SonicFastest | 2074794 | 832.47 MB/s | 2746593 | 4020 | 6.4× |
| Sonic | 2077707 | 831.30 MB/s | 2754296 | 4020 | 6.4× |
| Goccy | 2441835 | 707.34 MB/s | 2584031 | 14605 | 5.4× |
| LightningDecodeAny | 4156112 | 120.38 MB/s | 4493904 | 67881 | 3.2× |
| Easyjson | 4247962 | 406.60 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4298478 | 401.82 MB/s | 1011634 | 7594 | 3.1× |
| Stdlib | 13204633 | 130.80 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 836 | 2167.95 MB/s | 0 | 0 | 16.7× |
| Lightning | 844 | 2146.15 MB/s | 0 | 0 | 16.5× |
| LightningDestructive | 856 | 2115.99 MB/s | 0 | 0 | 16.3× |
| Easyjson | 2533 | 715.34 MB/s | 24 | 1 | 5.5× |
| Goccy | 2838 | 638.38 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6030 | 300.50 MB/s | 3768 | 40 | 2.3× |
| Sonic | 6049 | 299.54 MB/s | 3762 | 40 | 2.3× |
| JSONV2 | 7832 | 231.37 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7959 | 227.54 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13924 | 130.13 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 849 | 2133.71 MB/s | 0 | 0 | 16.4× |
| LightningArena | 860 | 2107.87 MB/s | 0 | 0 | 16.2× |
| LightningDestructive | 875 | 2069.90 MB/s | 0 | 0 | 15.9× |
| Easyjson | 2544 | 712.36 MB/s | 24 | 1 | 5.5× |
| Goccy | 2854 | 634.91 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6112 | 296.46 MB/s | 3863 | 40 | 2.3× |
| SonicFastest | 6136 | 295.33 MB/s | 3954 | 40 | 2.3× |
| JSONV2 | 7709 | 235.06 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8341 | 217.13 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13943 | 129.96 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1024 | 1768.96 MB/s | 144 | 10 | 13.6× |
| LightningArena | 1039 | 1744.35 MB/s | 144 | 10 | 13.4× |
| LightningDestructive | 1087 | 1667.50 MB/s | 144 | 10 | 12.9× |
| Easyjson | 2769 | 654.31 MB/s | 144 | 10 | 5.0× |
| Goccy | 2965 | 611.22 MB/s | 2600 | 5 | 4.7× |
| Sonic | 6195 | 292.52 MB/s | 3875 | 42 | 2.3× |
| SonicFastest | 6226 | 291.06 MB/s | 3946 | 42 | 2.2× |
| JSONV2 | 7956 | 227.75 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8335 | 217.27 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13973 | 129.68 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 627 | 787.90 MB/s | 160 | 1 | 8.7× |
| LightningDestructive | 629 | 785.79 MB/s | 160 | 1 | 8.7× |
| LightningDecodeAny | 1222 | 403.37 MB/s | 1040 | 25 | 4.5× |
| Sonic | 1246 | 396.59 MB/s | 979 | 6 | 4.4× |
| SonicFastest | 1250 | 395.14 MB/s | 980 | 6 | 4.4× |
| LightningArena | 1415 | 349.12 MB/s | 4120 | 2 | 3.9× |
| Easyjson | 2250 | 219.55 MB/s | 448 | 3 | 2.4× |
| Goccy | 2453 | 201.39 MB/s | 856 | 23 | 2.2× |
| JSONV2 | 3245 | 152.23 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5472 | 90.27 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 375 | 614.02 MB/s | 160 | 1 | 10.9× |
| LightningDestructive | 377 | 610.41 MB/s | 160 | 1 | 10.8× |
| SonicFastest | 878 | 261.99 MB/s | 669 | 6 | 4.6× |
| Sonic | 880 | 261.40 MB/s | 658 | 6 | 4.6× |
| LightningDecodeAny | 1047 | 218.75 MB/s | 1040 | 25 | 3.9× |
| LightningArena | 1142 | 201.44 MB/s | 4120 | 2 | 3.6× |
| Easyjson | 1406 | 163.61 MB/s | 448 | 3 | 2.9× |
| Goccy | 1569 | 146.63 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2365 | 97.25 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4065 | 56.58 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 51739 | 1258.85 MB/s | 97220 | 98 | 10.7× |
| Lightning | 51948 | 1253.78 MB/s | 103674 | 99 | 10.6× |
| LightningArena | 52145 | 1249.06 MB/s | 103680 | 99 | 10.6× |
| Sonic | 98861 | 658.83 MB/s | 155645 | 75 | 5.6× |
| SonicFastest | 98989 | 657.97 MB/s | 155814 | 75 | 5.6× |
| Goccy | 150091 | 433.95 MB/s | 229286 | 134 | 3.7× |
| LightningDecodeAny | 181259 | 294.21 MB/s | 176541 | 3237 | 3.0× |
| JSONV2 | 228630 | 284.88 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 551863 | 118.02 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2119015 | 915.74 MB/s | 2185296 | 1350 | 10.9× |
| Lightning | 2183915 | 888.53 MB/s | 2185298 | 1350 | 10.6× |
| LightningArena | 2195575 | 883.81 MB/s | 2185297 | 1350 | 10.5× |
| Sonic | 4602077 | 421.65 MB/s | 14606973 | 1407 | 5.0× |
| SonicFastest | 4739479 | 409.43 MB/s | 14608573 | 1407 | 4.9× |
| Goccy | 4760602 | 407.61 MB/s | 4066352 | 13510 | 4.9× |
| Easyjson | 7570628 | 256.32 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 9095293 | 213.35 MB/s | 6627984 | 206416 | 2.5× |
| JSONV2 | 11251321 | 172.47 MB/s | 3237222 | 13947 | 2.1× |
| Stdlib | 23091836 | 84.03 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 890969 | 3735.07 MB/s | 351704 | 1286 | 23.6× |
| LightningArena | 1348548 | 2467.71 MB/s | 2434342 | 1413 | 15.6× |
| Lightning | 1353235 | 2459.17 MB/s | 2434332 | 1413 | 15.6× |
| Sonic | 2763741 | 1204.10 MB/s | 6455324 | 4248 | 7.6× |
| SonicFastest | 2782944 | 1195.80 MB/s | 6406953 | 4248 | 7.6× |
| LightningDecodeAny | 3358644 | 915.18 MB/s | 4825572 | 55311 | 6.3× |
| Goccy | 4669909 | 712.61 MB/s | 3948909 | 3816 | 4.5× |
| JSONV2 | 7555569 | 440.45 MB/s | 5364511 | 13243 | 2.8× |
| Stdlib | 21059031 | 158.02 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 180932 | 1217.84 MB/s | 135392 | 226 | 11.2× |
| LightningArena | 181279 | 1215.51 MB/s | 135392 | 226 | 11.1× |
| LightningDestructive | 183004 | 1204.05 MB/s | 135392 | 226 | 11.0× |
| Sonic | 383183 | 575.04 MB/s | 308870 | 398 | 5.3× |
| SonicFastest | 385849 | 571.07 MB/s | 315972 | 398 | 5.2× |
| Goccy | 444582 | 495.62 MB/s | 365034 | 1067 | 4.5× |
| Easyjson | 545066 | 404.26 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 724721 | 304.04 MB/s | 129743 | 470 | 2.8× |
| LightningDecodeAny | 865014 | 125.22 MB/s | 854737 | 11700 | 2.3× |
| Stdlib | 2020415 | 109.06 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9995842 | 810.34 MB/s | 8109648 | 20809 | 8.8× |
| Lightning | 10018536 | 808.51 MB/s | 8109654 | 20809 | 8.8× |
| LightningArena | 10046417 | 806.26 MB/s | 8109649 | 20809 | 8.8× |
| Sonic | 16712709 | 484.66 MB/s | 70902048 | 40014 | 5.3× |
| SonicFastest | 16933379 | 478.35 MB/s | 70915647 | 40014 | 5.2× |
| Goccy | 23738204 | 341.22 MB/s | 17083527 | 107148 | 3.7× |
| LightningDecodeAny | 30812429 | 168.86 MB/s | 28359831 | 746961 | 2.9× |
| Easyjson | 30817298 | 262.84 MB/s | 15059619 | 41643 | 2.9× |
| JSONV2 | 44309923 | 182.80 MB/s | 15233762 | 78972 | 2.0× |
| Stdlib | 88044131 | 92.00 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4414005 | 675.91 MB/s | 3780457 | 1514 | 10.6× |
| LightningDestructive | 4621501 | 645.56 MB/s | 3758856 | 29356 | 10.1× |
| Lightning | 4758766 | 626.94 MB/s | 3758859 | 29356 | 9.8× |
| Sonic | 8738626 | 341.41 MB/s | 26579934 | 56760 | 5.4× |
| SonicFastest | 8742024 | 341.28 MB/s | 26558788 | 56760 | 5.3× |
| LightningDecodeAny | 15214879 | 120.55 MB/s | 18225375 | 350883 | 3.1× |
| Goccy | 16774382 | 177.86 MB/s | 10631567 | 273649 | 2.8× |
| Easyjson | 16954712 | 175.97 MB/s | 9479440 | 30115 | 2.8× |
| JSONV2 | 24306060 | 122.75 MB/s | 9257185 | 86278 | 1.9× |
| Stdlib | 46759898 | 63.80 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 924189 | 782.95 MB/s | 907601 | 3618 | 12.3× |
| LightningArena | 941026 | 768.95 MB/s | 916258 | 37 | 12.1× |
| Lightning | 981920 | 736.92 MB/s | 907598 | 3618 | 11.6× |
| SonicFastest | 1776205 | 407.38 MB/s | 3180937 | 7226 | 6.4× |
| Sonic | 1782660 | 405.91 MB/s | 3186896 | 7226 | 6.4× |
| LightningDecodeAny | 4029721 | 161.44 MB/s | 5691594 | 76540 | 2.8× |
| Easyjson | 4193492 | 172.55 MB/s | 2847904 | 3698 | 2.7× |
| Goccy | 4799755 | 150.76 MB/s | 2787137 | 80272 | 2.4× |
| JSONV2 | 5450064 | 132.77 MB/s | 2704627 | 7318 | 2.1× |
| Stdlib | 11397402 | 63.49 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1378602 | 1144.17 MB/s | 907600 | 3618 | 11.3× |
| LightningArena | 1387621 | 1136.73 MB/s | 916256 | 37 | 11.2× |
| Lightning | 1424885 | 1107.00 MB/s | 907595 | 3618 | 10.9× |
| Sonic | 2269013 | 695.17 MB/s | 5785270 | 7226 | 6.8× |
| SonicFastest | 2292895 | 687.93 MB/s | 5785976 | 7226 | 6.8× |
| LightningDecodeAny | 3503082 | 215.07 MB/s | 5691592 | 76540 | 4.4× |
| Easyjson | 5540276 | 284.71 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5631856 | 280.08 MB/s | 3584374 | 80268 | 2.8× |
| JSONV2 | 6330759 | 249.16 MB/s | 2704594 | 7318 | 2.5× |
| Stdlib | 15532056 | 101.55 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 156349 | 960.19 MB/s | 81920 | 1 | 11.9× |
| LightningArena | 156771 | 957.60 MB/s | 81920 | 1 | 11.9× |
| Lightning | 156864 | 957.03 MB/s | 81920 | 1 | 11.8× |
| SonicFastest | 274341 | 547.22 MB/s | 252543 | 6 | 6.8× |
| Sonic | 279390 | 537.33 MB/s | 267104 | 6 | 6.7× |
| LightningDecodeAny | 434967 | 345.13 MB/s | 745508 | 10015 | 4.3× |
| Goccy | 860665 | 174.43 MB/s | 325428 | 10005 | 2.2× |
| JSONV2 | 1065333 | 140.92 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1858547 | 80.77 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 27223 | 1032.85 MB/s | 29088 | 101 | 11.1× |
| LightningArena | 27428 | 1025.13 MB/s | 29237 | 101 | 11.0× |
| Lightning | 27432 | 1024.96 MB/s | 29235 | 101 | 11.0× |
| SonicFastest | 63651 | 441.74 MB/s | 46275 | 103 | 4.7× |
| Sonic | 63703 | 441.37 MB/s | 46513 | 103 | 4.7× |
| Easyjson | 68372 | 411.24 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72006 | 390.48 MB/s | 59178 | 188 | 4.2× |
| JSONV2 | 133479 | 210.65 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 147569 | 190.53 MB/s | 133659 | 2639 | 2.0× |
| Stdlib | 300976 | 93.42 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1532 | 1520.04 MB/s | 32 | 1 | 14.7× |
| Lightning | 1536 | 1515.63 MB/s | 32 | 1 | 14.7× |
| LightningDestructive | 1589 | 1465.01 MB/s | 32 | 1 | 14.2× |
| Goccy | 4134 | 563.16 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4243 | 548.70 MB/s | 192 | 2 | 5.3× |
| Sonic | 5062 | 459.94 MB/s | 4231 | 6 | 4.5× |
| SonicFastest | 5070 | 459.21 MB/s | 4229 | 6 | 4.4× |
| JSONV2 | 8457 | 275.28 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9844 | 171.18 MB/s | 9936 | 194 | 2.3× |
| Stdlib | 22544 | 103.26 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 178 | 1062.74 MB/s | 0 | 0 | 13.8× |
| Lightning | 178 | 1062.63 MB/s | 0 | 0 | 13.8× |
| LightningDestructive | 180 | 1049.72 MB/s | 0 | 0 | 13.6× |
| Goccy | 399 | 473.39 MB/s | 304 | 2 | 6.1× |
| Easyjson | 501 | 377.51 MB/s | 0 | 0 | 4.9× |
| SonicFastest | 816 | 231.46 MB/s | 517 | 4 | 3.0× |
| Sonic | 817 | 231.21 MB/s | 515 | 4 | 3.0× |
| JSONV2 | 1037 | 182.23 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1266 | 105.88 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2448 | 77.20 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1097 | 1998.07 MB/s | 0 | 0 | 14.4× |
| LightningArena | 1097 | 1997.86 MB/s | 0 | 0 | 14.4× |
| LightningDestructive | 1126 | 1946.52 MB/s | 0 | 0 | 14.1× |
| Goccy | 3163 | 692.59 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3180 | 688.92 MB/s | 24 | 1 | 5.0× |
| Sonic | 6372 | 343.87 MB/s | 3967 | 40 | 2.5× |
| SonicFastest | 6387 | 343.02 MB/s | 3956 | 40 | 2.5× |
| LightningDecodeAny | 7908 | 229.01 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 7927 | 276.38 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15822 | 138.48 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 574804 | 888.09 MB/s | 318401 | 1005 | 10.5× |
| LightningArena | 578017 | 883.15 MB/s | 318401 | 1005 | 10.4× |
| LightningDestructive | 578976 | 881.69 MB/s | 318400 | 1005 | 10.4× |
| Sonic | 1158893 | 440.49 MB/s | 858240 | 2006 | 5.2× |
| SonicFastest | 1162510 | 439.12 MB/s | 864326 | 2006 | 5.2× |
| Goccy | 1184582 | 430.93 MB/s | 1133842 | 5006 | 5.1× |
| Easyjson | 1566740 | 325.82 MB/s | 863777 | 3012 | 3.8× |
| JSONV2 | 3235024 | 157.80 MB/s | 1076022 | 12646 | 1.9× |
| LightningDecodeAny | 3351578 | 137.69 MB/s | 2742394 | 64017 | 1.8× |
| Stdlib | 6018797 | 84.81 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 481 | 41156.31 MB/s | 0 | 0 | 226.2× |
| Lightning | 481 | 41137.41 MB/s | 0 | 0 | 226.1× |
| LightningDestructive | 481 | 41141.70 MB/s | 0 | 0 | 226.1× |
| Goccy | 20619 | 959.75 MB/s | 20491 | 2 | 5.3× |
| SonicFastest | 28141 | 703.21 MB/s | 22057 | 4 | 3.9× |
| Sonic | 28170 | 702.49 MB/s | 22205 | 4 | 3.9× |
| JSONV2 | 29709 | 666.09 MB/s | 8 | 1 | 3.7× |
| Easyjson | 81928 | 241.54 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 87043 | 227.34 MB/s | 116608 | 2014 | 1.2× |
| Stdlib | 108739 | 181.99 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1737 | 10435.12 MB/s | 0 | 0 | 59.1× |
| LightningArena | 1809 | 10016.89 MB/s | 405 | 0 | 56.7× |
| Lightning | 1811 | 10010.39 MB/s | 405 | 0 | 56.7× |
| Easyjson | 3961 | 4575.57 MB/s | 432 | 2 | 25.9× |
| Sonic | 10077 | 1798.61 MB/s | 22885 | 6 | 10.2× |
| SonicFastest | 10089 | 1796.49 MB/s | 23152 | 6 | 10.2× |
| Goccy | 15873 | 1141.84 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 16321 | 1095.61 MB/s | 29102 | 189 | 6.3× |
| JSONV2 | 45406 | 399.16 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102641 | 176.58 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2145765 | 936.03 MB/s | 3089564 | 6821 | 8.7× |
| LightningArena | 2220733 | 904.43 MB/s | 3100187 | 6700 | 8.4× |
| Lightning | 2225283 | 902.58 MB/s | 3096638 | 6822 | 8.4× |
| Goccy | 4227209 | 475.13 MB/s | 5412162 | 15831 | 4.4× |
| SonicFastest | 4530667 | 443.31 MB/s | 10942559 | 13683 | 4.1× |
| Sonic | 4532186 | 443.16 MB/s | 10972684 | 13683 | 4.1× |
| Easyjson | 4944674 | 406.19 MB/s | 2981483 | 7439 | 3.8× |
| JSONV2 | 6968817 | 288.21 MB/s | 3173681 | 14563 | 2.7× |
| LightningDecodeAny | 7044463 | 162.15 MB/s | 7376388 | 134004 | 2.6× |
| Stdlib | 18589531 | 108.04 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 849 | 646.42 MB/s | 480 | 1 | 6.6× |
| LightningArena | 855 | 642.32 MB/s | 480 | 1 | 6.5× |
| LightningDestructive | 855 | 641.90 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1553 | 352.86 MB/s | 1765 | 45 | 3.6× |
| Easyjson | 2157 | 254.50 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2628 | 208.94 MB/s | 1905 | 26 | 2.1× |
| Sonic | 2634 | 208.40 MB/s | 1930 | 26 | 2.1× |
| Goccy | 3038 | 180.74 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3360 | 163.38 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5598 | 98.07 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 410756 | 1537.44 MB/s | 400488 | 545 | 13.0× |
| LightningArena | 446361 | 1414.81 MB/s | 449120 | 404 | 12.0× |
| Lightning | 446585 | 1414.10 MB/s | 446947 | 548 | 12.0× |
| SonicFastest | 1031396 | 612.29 MB/s | 994363 | 1102 | 5.2× |
| Sonic | 1038152 | 608.31 MB/s | 1008214 | 1102 | 5.2× |
| Easyjson | 1154323 | 547.09 MB/s | 422505 | 936 | 4.6× |
| Goccy | 1177170 | 536.47 MB/s | 988346 | 1201 | 4.6× |
| JSONV2 | 2170966 | 290.89 MB/s | 571616 | 3144 | 2.5× |
| LightningDecodeAny | 2303430 | 202.70 MB/s | 1992814 | 29072 | 2.3× |
| Stdlib | 5360054 | 117.82 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 580065 | 969.56 MB/s | 391286 | 426 | 9.1× |
| Lightning | 643552 | 873.91 MB/s | 506296 | 433 | 8.2× |
| LightningArena | 648025 | 867.88 MB/s | 508217 | 287 | 8.1× |
| Sonic | 1048373 | 536.46 MB/s | 954757 | 1476 | 5.0× |
| SonicFastest | 1062450 | 529.35 MB/s | 965462 | 1476 | 5.0× |
| Goccy | 1343211 | 418.70 MB/s | 1039635 | 1030 | 3.9× |
| Easyjson | 1773924 | 317.04 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2380143 | 236.29 MB/s | 1991559 | 28581 | 2.2× |
| JSONV2 | 2784302 | 201.99 MB/s | 927444 | 3482 | 1.9× |
| Stdlib | 5262757 | 106.87 MB/s | 1011674 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 573493 | 929.70 MB/s | 333416 | 2084 | 9.5× |
| Lightning | 582968 | 914.59 MB/s | 367438 | 2086 | 9.3× |
| LightningArena | 588455 | 906.06 MB/s | 367520 | 2086 | 9.2× |
| Easyjson | 1123046 | 474.76 MB/s | 428361 | 3273 | 4.8× |
| Sonic | 1157318 | 460.70 MB/s | 1037192 | 4351 | 4.7× |
| SonicFastest | 1159417 | 459.87 MB/s | 1027747 | 4351 | 4.7× |
| Goccy | 1311724 | 406.47 MB/s | 1167222 | 5409 | 4.1× |
| JSONV2 | 2539231 | 209.98 MB/s | 745454 | 13288 | 2.1× |
| LightningDecodeAny | 3286706 | 162.22 MB/s | 2659130 | 49349 | 1.7× |
| Stdlib | 5435147 | 98.10 MB/s | 798693 | 17133 | 1.0× |
