# JSON Deserialization Benchmarks

- generated 2026-07-02T08:51:24Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105189 | 1209.97 MB/s | 49760 | 3 | 12.4× |
| LightningDestructive | 108187 | 1176.43 MB/s | 49280 | 2 | 12.1× |
| Sonic | 201720 | 630.95 MB/s | 214406 | 15 | 6.5× |
| SonicFastest | 201775 | 630.78 MB/s | 214477 | 15 | 6.5× |
| Goccy | 244746 | 520.03 MB/s | 225282 | 884 | 5.3× |
| Easyjson | 269995 | 471.40 MB/s | 122864 | 14 | 4.8× |
| JSONV2 | 454217 | 280.21 MB/s | 195127 | 1805 | 2.9× |
| LightningDecodeAny | 465510 | 203.33 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1305975 | 97.46 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4110877 | 547.58 MB/s | 3008241 | 1158 | 8.3× |
| Lightning | 4240131 | 530.89 MB/s | 3008242 | 1158 | 8.0× |
| Sonic | 4817690 | 467.25 MB/s | 4873342 | 2584 | 7.1× |
| SonicFastest | 4910049 | 458.46 MB/s | 4872422 | 2584 | 6.9× |
| LightningDecodeAny | 12232032 | 184.03 MB/s | 7938298 | 281383 | 2.8× |
| Goccy | 13431373 | 167.60 MB/s | 4159048 | 56533 | 2.5× |
| Easyjson | 13796646 | 163.16 MB/s | 3099809 | 2120 | 2.5× |
| JSONV2 | 16909509 | 133.12 MB/s | 3123189 | 3083 | 2.0× |
| Stdlib | 34080961 | 66.05 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 568150 | 475.94 MB/s | 392945 | 568 | 7.8× |
| LightningDestructive | 586209 | 461.27 MB/s | 392944 | 568 | 7.5× |
| Sonic | 761832 | 354.94 MB/s | 641622 | 1147 | 5.8× |
| SonicFastest | 762221 | 354.76 MB/s | 641643 | 1147 | 5.8× |
| LightningDecodeAny | 1748797 | 154.62 MB/s | 1011488 | 37901 | 2.5× |
| Easyjson | 1762392 | 153.43 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1860765 | 145.32 MB/s | 542228 | 8122 | 2.4× |
| JSONV2 | 2267305 | 119.26 MB/s | 348160 | 1628 | 1.9× |
| Stdlib | 4405294 | 61.38 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1151409 | 1500.08 MB/s | 767864 | 2798 | 15.3× |
| Lightning | 1192310 | 1448.62 MB/s | 767905 | 2799 | 14.8× |
| SonicFastest | 2064951 | 836.44 MB/s | 2694899 | 5547 | 8.5× |
| Sonic | 2068772 | 834.89 MB/s | 2695094 | 5547 | 8.5× |
| Goccy | 2810774 | 614.49 MB/s | 2580678 | 14603 | 6.3× |
| Easyjson | 4235287 | 407.81 MB/s | 972032 | 5389 | 4.2× |
| LightningDecodeAny | 4317471 | 115.88 MB/s | 4976203 | 81466 | 4.1× |
| JSONV2 | 4685714 | 368.61 MB/s | 1011615 | 7594 | 3.8× |
| Stdlib | 17594280 | 98.17 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1014 | 1787.22 MB/s | 0 | 0 | 15.8× |
| LightningDestructive | 1078 | 1681.38 MB/s | 0 | 0 | 14.9× |
| Easyjson | 2974 | 609.22 MB/s | 24 | 1 | 5.4× |
| Goccy | 3210 | 564.45 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6276 | 288.74 MB/s | 3345 | 38 | 2.6× |
| Sonic | 6491 | 279.14 MB/s | 3345 | 38 | 2.5× |
| JSONV2 | 8161 | 222.04 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9366 | 193.37 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16042 | 112.95 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1096 | 1653.49 MB/s | 0 | 0 | 14.7× |
| LightningDestructive | 1166 | 1554.66 MB/s | 0 | 0 | 13.8× |
| Easyjson | 3007 | 602.66 MB/s | 24 | 1 | 5.4× |
| Goccy | 3262 | 555.52 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6072 | 298.40 MB/s | 3347 | 38 | 2.7× |
| Sonic | 6289 | 288.11 MB/s | 3345 | 38 | 2.6× |
| JSONV2 | 8204 | 220.86 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9328 | 194.14 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16123 | 112.39 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1333 | 1358.99 MB/s | 144 | 10 | 12.2× |
| LightningDestructive | 1436 | 1262.17 MB/s | 144 | 10 | 11.3× |
| Easyjson | 3123 | 580.13 MB/s | 144 | 10 | 5.2× |
| Goccy | 3462 | 523.36 MB/s | 2600 | 5 | 4.7× |
| SonicFastest | 6106 | 296.77 MB/s | 3365 | 40 | 2.7× |
| Sonic | 6300 | 287.64 MB/s | 3366 | 40 | 2.6× |
| JSONV2 | 8538 | 212.22 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9253 | 195.73 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 16234 | 111.62 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 703 | 702.45 MB/s | 160 | 1 | 9.8× |
| LightningDestructive | 724 | 682.35 MB/s | 160 | 1 | 9.6× |
| Sonic | 1287 | 383.70 MB/s | 1075 | 8 | 5.4× |
| SonicFastest | 1294 | 381.68 MB/s | 1076 | 8 | 5.4× |
| LightningDecodeAny | 1826 | 270.05 MB/s | 1536 | 30 | 3.8× |
| Easyjson | 2521 | 195.99 MB/s | 448 | 3 | 2.7× |
| Goccy | 2687 | 183.84 MB/s | 856 | 23 | 2.6× |
| JSONV2 | 3253 | 151.87 MB/s | 528 | 7 | 2.1× |
| Stdlib | 6925 | 71.33 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 446 | 515.28 MB/s | 160 | 1 | 11.3× |
| LightningDestructive | 451 | 510.14 MB/s | 160 | 1 | 11.2× |
| Sonic | 927 | 248.06 MB/s | 801 | 8 | 5.4× |
| SonicFastest | 928 | 247.89 MB/s | 801 | 8 | 5.4× |
| LightningDecodeAny | 1666 | 137.48 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1735 | 132.54 MB/s | 448 | 3 | 2.9× |
| Goccy | 1821 | 126.27 MB/s | 584 | 23 | 2.8× |
| JSONV2 | 2617 | 87.90 MB/s | 528 | 7 | 1.9× |
| Stdlib | 5035 | 45.68 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 89260 | 729.69 MB/s | 164881 | 105 | 7.6× |
| LightningDestructive | 92319 | 705.51 MB/s | 158661 | 100 | 7.3× |
| Sonic | 148939 | 437.31 MB/s | 235850 | 65 | 4.5× |
| SonicFastest | 154891 | 420.50 MB/s | 236034 | 65 | 4.4× |
| Goccy | 181042 | 359.76 MB/s | 228345 | 134 | 3.7× |
| LightningDecodeAny | 207323 | 257.23 MB/s | 176960 | 3252 | 3.3× |
| JSONV2 | 267195 | 243.76 MB/s | 206666 | 607 | 2.5× |
| Stdlib | 673919 | 96.65 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2492221 | 778.61 MB/s | 2813904 | 1358 | 11.3× |
| Lightning | 2651938 | 731.72 MB/s | 2813906 | 1358 | 10.6× |
| Sonic | 5020791 | 386.49 MB/s | 4880184 | 1736 | 5.6× |
| SonicFastest | 5031435 | 385.67 MB/s | 4880053 | 1736 | 5.6× |
| Goccy | 5292189 | 366.67 MB/s | 4062320 | 13509 | 5.3× |
| Easyjson | 8132611 | 238.60 MB/s | 3871265 | 15043 | 3.5× |
| LightningDecodeAny | 10091002 | 192.30 MB/s | 7013525 | 219937 | 2.8× |
| JSONV2 | 12804738 | 151.54 MB/s | 3237191 | 13947 | 2.2× |
| Stdlib | 28181019 | 68.86 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1066839 | 3119.34 MB/s | 351704 | 1286 | 22.5× |
| Lightning | 1624256 | 2048.83 MB/s | 2488908 | 2995 | 14.8× |
| Sonic | 2074643 | 1604.05 MB/s | 5896113 | 4263 | 11.6× |
| SonicFastest | 2091445 | 1591.16 MB/s | 5894311 | 4263 | 11.5× |
| LightningDecodeAny | 3666029 | 838.45 MB/s | 4886622 | 56892 | 6.6× |
| Goccy | 6125662 | 543.26 MB/s | 3948917 | 3817 | 3.9× |
| JSONV2 | 8150087 | 408.32 MB/s | 5364500 | 13243 | 2.9× |
| Stdlib | 24018426 | 138.55 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 228810 | 963.01 MB/s | 135392 | 226 | 10.7× |
| LightningDestructive | 243462 | 905.05 MB/s | 135392 | 226 | 10.0× |
| Sonic | 483564 | 455.67 MB/s | 350128 | 262 | 5.0× |
| SonicFastest | 486843 | 452.60 MB/s | 350568 | 262 | 5.0× |
| Goccy | 491592 | 448.23 MB/s | 363833 | 1066 | 5.0× |
| Easyjson | 670944 | 328.41 MB/s | 130512 | 245 | 3.6× |
| JSONV2 | 774064 | 284.66 MB/s | 129747 | 470 | 3.2× |
| LightningDecodeAny | 1044783 | 103.67 MB/s | 861394 | 11944 | 2.3× |
| Stdlib | 2441535 | 90.25 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 13455841 | 601.97 MB/s | 15730641 | 20821 | 8.4× |
| Lightning | 13899728 | 582.75 MB/s | 15730641 | 20821 | 8.1× |
| Sonic | 17870835 | 453.25 MB/s | 19859979 | 41640 | 6.3× |
| SonicFastest | 18116405 | 447.11 MB/s | 19859580 | 41640 | 6.2× |
| Goccy | 28119101 | 288.06 MB/s | 19113004 | 107156 | 4.0× |
| Easyjson | 35141141 | 230.50 MB/s | 15059620 | 41643 | 3.2× |
| LightningDecodeAny | 42729726 | 121.77 MB/s | 34333348 | 912810 | 2.6× |
| JSONV2 | 50897267 | 159.14 MB/s | 15233772 | 78972 | 2.2× |
| Stdlib | 112675862 | 71.89 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5707760 | 522.70 MB/s | 3908872 | 29356 | 10.1× |
| Lightning | 5862649 | 508.89 MB/s | 3908876 | 29356 | 9.9× |
| SonicFastest | 9603525 | 310.66 MB/s | 9130716 | 57804 | 6.0× |
| Sonic | 9691594 | 307.84 MB/s | 9130799 | 57804 | 6.0× |
| Goccy | 19531701 | 152.75 MB/s | 9944262 | 273623 | 3.0× |
| Easyjson | 19640888 | 151.90 MB/s | 9479441 | 30115 | 2.9× |
| LightningDecodeAny | 21874932 | 83.85 MB/s | 20023835 | 408557 | 2.6× |
| JSONV2 | 29335423 | 101.70 MB/s | 9257042 | 86278 | 2.0× |
| Stdlib | 57835883 | 51.59 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1356562 | 533.41 MB/s | 907601 | 3618 | 10.5× |
| Lightning | 1412365 | 512.33 MB/s | 907596 | 3618 | 10.1× |
| Sonic | 2165540 | 334.14 MB/s | 2369843 | 3683 | 6.6× |
| SonicFastest | 2193820 | 329.83 MB/s | 2369595 | 3683 | 6.5× |
| Easyjson | 5569852 | 129.91 MB/s | 2847907 | 3698 | 2.6× |
| LightningDecodeAny | 5724970 | 113.64 MB/s | 5752202 | 80172 | 2.5× |
| Goccy | 5815237 | 124.43 MB/s | 2726114 | 80268 | 2.5× |
| JSONV2 | 6534480 | 110.74 MB/s | 2704707 | 7318 | 2.2× |
| Stdlib | 14305572 | 50.58 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2025399 | 778.79 MB/s | 907596 | 3618 | 9.7× |
| LightningDestructive | 2067734 | 762.84 MB/s | 907601 | 3618 | 9.5× |
| Sonic | 2437962 | 647.00 MB/s | 3224812 | 3683 | 8.1× |
| SonicFastest | 2440909 | 646.22 MB/s | 3224396 | 3683 | 8.1× |
| LightningDecodeAny | 4861830 | 154.96 MB/s | 5752198 | 80172 | 4.1× |
| Easyjson | 6719794 | 234.73 MB/s | 2847906 | 3698 | 2.9× |
| Goccy | 6957790 | 226.70 MB/s | 3470387 | 80261 | 2.8× |
| JSONV2 | 7308540 | 215.82 MB/s | 2704551 | 7318 | 2.7× |
| Stdlib | 19690662 | 80.11 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 229417 | 654.37 MB/s | 81920 | 1 | 9.4× |
| LightningDestructive | 242526 | 619.00 MB/s | 81920 | 1 | 8.9× |
| Sonic | 392686 | 382.30 MB/s | 407416 | 16 | 5.5× |
| SonicFastest | 413445 | 363.10 MB/s | 408158 | 16 | 5.2× |
| LightningDecodeAny | 605199 | 248.05 MB/s | 746006 | 10020 | 3.6× |
| Goccy | 1001580 | 149.89 MB/s | 328037 | 10005 | 2.1× |
| JSONV2 | 1180233 | 127.20 MB/s | 357725 | 20 | 1.8× |
| Stdlib | 2150715 | 69.80 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33134 | 848.57 MB/s | 29216 | 103 | 10.3× |
| LightningDestructive | 34191 | 822.35 MB/s | 29088 | 101 | 10.0× |
| Sonic | 60563 | 464.26 MB/s | 59505 | 83 | 5.7× |
| SonicFastest | 61034 | 460.68 MB/s | 59520 | 83 | 5.6× |
| Goccy | 80210 | 350.54 MB/s | 59290 | 188 | 4.3× |
| Easyjson | 84837 | 331.42 MB/s | 32304 | 138 | 4.0× |
| JSONV2 | 143973 | 195.29 MB/s | 36897 | 242 | 2.4× |
| LightningDecodeAny | 167753 | 167.61 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 342750 | 82.03 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1954 | 1191.13 MB/s | 32 | 1 | 13.4× |
| LightningDestructive | 2103 | 1107.08 MB/s | 32 | 1 | 12.5× |
| Sonic | 4698 | 495.57 MB/s | 3711 | 4 | 5.6× |
| SonicFastest | 4726 | 492.56 MB/s | 3713 | 4 | 5.6× |
| Goccy | 4931 | 472.15 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 5047 | 461.28 MB/s | 192 | 2 | 5.2× |
| JSONV2 | 8550 | 272.27 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10684 | 157.72 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 26236 | 88.73 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 216 | 873.79 MB/s | 0 | 0 | 12.9× |
| LightningDestructive | 226 | 835.13 MB/s | 0 | 0 | 12.4× |
| Goccy | 425 | 444.57 MB/s | 304 | 2 | 6.6× |
| Easyjson | 583 | 324.22 MB/s | 0 | 0 | 4.8× |
| SonicFastest | 646 | 292.51 MB/s | 341 | 3 | 4.3× |
| Sonic | 649 | 291.25 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1018 | 185.58 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1356 | 98.79 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2797 | 67.57 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1459 | 1502.01 MB/s | 0 | 0 | 13.2× |
| LightningDestructive | 1547 | 1416.05 MB/s | 0 | 0 | 12.4× |
| Goccy | 3635 | 602.77 MB/s | 2864 | 4 | 5.3× |
| Easyjson | 3824 | 572.93 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6499 | 337.15 MB/s | 3603 | 38 | 3.0× |
| Sonic | 6719 | 326.08 MB/s | 3602 | 38 | 2.9× |
| JSONV2 | 9030 | 242.63 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 9242 | 195.96 MB/s | 7536 | 158 | 2.1× |
| Stdlib | 19191 | 114.17 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 698910 | 730.39 MB/s | 703299 | 1010 | 10.1× |
| Lightning | 737875 | 691.82 MB/s | 703298 | 1010 | 9.6× |
| SonicFastest | 1284740 | 397.34 MB/s | 1308029 | 2014 | 5.5× |
| Sonic | 1303741 | 391.55 MB/s | 1310108 | 2014 | 5.4× |
| Goccy | 1319644 | 386.83 MB/s | 1137730 | 5006 | 5.3× |
| Easyjson | 1693752 | 301.39 MB/s | 863778 | 3012 | 4.2× |
| JSONV2 | 3447994 | 148.05 MB/s | 1075953 | 12645 | 2.0× |
| LightningDecodeAny | 3584458 | 128.74 MB/s | 2785928 | 66022 | 2.0× |
| Stdlib | 7054350 | 72.36 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 547 | 36189.68 MB/s | 0 | 0 | 295.7× |
| LightningDestructive | 856 | 23119.96 MB/s | 0 | 0 | 188.9× |
| SonicFastest | 6577 | 3008.89 MB/s | 21142 | 3 | 24.6× |
| Goccy | 29256 | 676.41 MB/s | 20492 | 2 | 5.5× |
| Sonic | 31949 | 619.40 MB/s | 20612 | 3 | 5.1× |
| JSONV2 | 33273 | 594.75 MB/s | 8 | 1 | 4.9× |
| LightningDecodeAny | 98673 | 200.54 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 100321 | 197.26 MB/s | 0 | 0 | 1.6× |
| Stdlib | 161685 | 122.39 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2220 | 8164.86 MB/s | 432 | 2 | 53.5× |
| LightningDestructive | 2488 | 7286.00 MB/s | 0 | 0 | 47.7× |
| Easyjson | 5057 | 3584.13 MB/s | 432 | 2 | 23.5× |
| SonicFastest | 9137 | 1983.48 MB/s | 20406 | 5 | 13.0× |
| Sonic | 9229 | 1963.90 MB/s | 20415 | 5 | 12.9× |
| LightningDecodeAny | 19196 | 931.53 MB/s | 29088 | 191 | 6.2× |
| Goccy | 22985 | 788.51 MB/s | 19460 | 2 | 5.2× |
| JSONV2 | 50469 | 359.11 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 118688 | 152.70 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2444894 | 821.51 MB/s | 3089821 | 6821 | 9.1× |
| Lightning | 2521954 | 796.40 MB/s | 3091533 | 6827 | 8.8× |
| SonicFastest | 4244814 | 473.16 MB/s | 5153031 | 7085 | 5.2× |
| Sonic | 4251451 | 472.43 MB/s | 5152632 | 7085 | 5.2× |
| Goccy | 4801208 | 418.33 MB/s | 5411484 | 15833 | 4.6× |
| Easyjson | 5482476 | 366.35 MB/s | 2981482 | 7439 | 4.1× |
| LightningDecodeAny | 7188522 | 158.90 MB/s | 7386073 | 134751 | 3.1× |
| JSONV2 | 7843566 | 256.07 MB/s | 3173677 | 14563 | 2.8× |
| Stdlib | 22274124 | 90.17 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 855 | 642.20 MB/s | 480 | 1 | 8.0× |
| Lightning | 856 | 641.69 MB/s | 480 | 1 | 7.9× |
| LightningDecodeAny | 2126 | 257.79 MB/s | 2261 | 50 | 3.2× |
| Easyjson | 2243 | 244.76 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2390 | 229.68 MB/s | 2261 | 8 | 2.8× |
| Sonic | 2410 | 227.81 MB/s | 2260 | 8 | 2.8× |
| Goccy | 3325 | 165.11 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3455 | 158.91 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 6801 | 80.72 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 518039 | 1219.05 MB/s | 400648 | 555 | 12.2× |
| Lightning | 592108 | 1066.55 MB/s | 449177 | 867 | 10.7× |
| Sonic | 978143 | 645.63 MB/s | 1065370 | 814 | 6.5× |
| SonicFastest | 1028250 | 614.16 MB/s | 1067631 | 814 | 6.1× |
| Easyjson | 1331017 | 474.46 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1404969 | 449.49 MB/s | 989730 | 1200 | 4.5× |
| JSONV2 | 2338050 | 270.10 MB/s | 571591 | 3144 | 2.7× |
| LightningDecodeAny | 2578662 | 181.07 MB/s | 2010197 | 30295 | 2.4× |
| Stdlib | 6311057 | 100.06 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 808098 | 695.96 MB/s | 577673 | 437 | 7.4× |
| Lightning | 1028284 | 546.94 MB/s | 801045 | 1243 | 5.8× |
| Sonic | 1271983 | 442.15 MB/s | 1348382 | 1185 | 4.7× |
| SonicFastest | 1278760 | 439.81 MB/s | 1348132 | 1185 | 4.7× |
| Goccy | 1622645 | 346.60 MB/s | 1044082 | 1028 | 3.7× |
| Easyjson | 2174755 | 258.61 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 3117968 | 180.38 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 3221471 | 174.58 MB/s | 927402 | 3482 | 1.9× |
| Stdlib | 6008610 | 93.60 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 624867 | 853.27 MB/s | 333416 | 2084 | 10.1× |
| Lightning | 694553 | 767.66 MB/s | 368224 | 2293 | 9.1× |
| SonicFastest | 1128206 | 472.59 MB/s | 981077 | 3082 | 5.6× |
| Sonic | 1128706 | 472.38 MB/s | 981214 | 3082 | 5.6× |
| Easyjson | 1301134 | 409.78 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1429838 | 372.89 MB/s | 1167080 | 5409 | 4.4× |
| JSONV2 | 2912532 | 183.06 MB/s | 745418 | 13288 | 2.2× |
| LightningDecodeAny | 3499867 | 152.34 MB/s | 2674616 | 50596 | 1.8× |
| Stdlib | 6333779 | 84.18 MB/s | 798693 | 17133 | 1.0× |
