# JSON Deserialization Benchmarks

- generated 2026-07-26T10:38:21Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105125 | 1210.70 MB/s | 49760 | 3 | 10.5× |
| LightningDestructive | 105370 | 1207.89 MB/s | 49280 | 2 | 10.5× |
| SonicFastest | 186682 | 681.77 MB/s | 210724 | 10 | 5.9× |
| Sonic | 189289 | 672.39 MB/s | 215314 | 10 | 5.8× |
| Goccy | 201227 | 632.50 MB/s | 225516 | 884 | 5.5× |
| Easyjson | 212147 | 599.94 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 423113 | 300.81 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 442255 | 214.02 MB/s | 465730 | 9708 | 2.5× |
| Stdlib | 1103920 | 115.29 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3630229 | 620.09 MB/s | 3008240 | 1158 | 7.2× |
| Lightning | 3676089 | 612.35 MB/s | 3008243 | 1158 | 7.1× |
| Sonic | 4450445 | 505.80 MB/s | 15232102 | 970 | 5.9× |
| SonicFastest | 4523199 | 497.67 MB/s | 15232101 | 970 | 5.8× |
| Goccy | 10243777 | 219.75 MB/s | 4116366 | 56532 | 2.5× |
| Easyjson | 10929472 | 205.96 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12257434 | 183.65 MB/s | 19380212 | 223896 | 2.1× |
| JSONV2 | 16154282 | 139.35 MB/s | 3123224 | 3083 | 1.6× |
| Stdlib | 26109483 | 86.22 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 460151 | 587.64 MB/s | 392944 | 568 | 7.3× |
| LightningDestructive | 467445 | 578.47 MB/s | 392944 | 568 | 7.2× |
| Sonic | 629906 | 429.28 MB/s | 463336 | 968 | 5.3× |
| SonicFastest | 632519 | 427.50 MB/s | 464373 | 968 | 5.3× |
| Goccy | 1395178 | 193.81 MB/s | 543334 | 8122 | 2.4× |
| Easyjson | 1402195 | 192.84 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1624530 | 166.45 MB/s | 2543877 | 29687 | 2.1× |
| JSONV2 | 2099248 | 128.81 MB/s | 348152 | 1628 | 1.6× |
| Stdlib | 3348217 | 80.76 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1178141 | 1466.04 MB/s | 767864 | 2798 | 11.2× |
| Lightning | 1188081 | 1453.78 MB/s | 767906 | 2799 | 11.1× |
| Sonic | 2053651 | 841.04 MB/s | 2736679 | 4020 | 6.4× |
| SonicFastest | 2057817 | 839.34 MB/s | 2755906 | 4020 | 6.4× |
| Goccy | 2342051 | 737.47 MB/s | 2581983 | 14604 | 5.6× |
| Easyjson | 4210753 | 410.19 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4277630 | 403.78 MB/s | 1011637 | 7594 | 3.1× |
| LightningDecodeAny | 4295135 | 116.48 MB/s | 4954733 | 76576 | 3.1× |
| Stdlib | 13214943 | 130.70 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1195 | 1516.79 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1210 | 1496.99 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2514 | 720.87 MB/s | 24 | 1 | 5.6× |
| Goccy | 2768 | 654.56 MB/s | 2608 | 4 | 5.1× |
| Sonic | 5911 | 306.57 MB/s | 3695 | 40 | 2.4× |
| SonicFastest | 5925 | 305.80 MB/s | 3714 | 40 | 2.4× |
| JSONV2 | 7797 | 232.41 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8082 | 224.08 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14099 | 128.52 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1225 | 1478.63 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1246 | 1454.32 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2537 | 714.35 MB/s | 24 | 1 | 5.6× |
| Goccy | 2861 | 633.45 MB/s | 2608 | 4 | 4.9× |
| Sonic | 5952 | 304.44 MB/s | 3805 | 40 | 2.4× |
| SonicFastest | 5966 | 303.71 MB/s | 3861 | 40 | 2.4× |
| JSONV2 | 7704 | 235.19 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8416 | 215.19 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14085 | 128.65 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1394 | 1300.30 MB/s | 144 | 10 | 10.1× |
| LightningDestructive | 1442 | 1256.52 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2765 | 655.26 MB/s | 144 | 10 | 5.1× |
| Goccy | 2883 | 628.62 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6073 | 298.39 MB/s | 3794 | 42 | 2.3× |
| SonicFastest | 6102 | 296.96 MB/s | 3770 | 42 | 2.3× |
| JSONV2 | 7995 | 226.65 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8073 | 224.32 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14012 | 129.32 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 728 | 678.74 MB/s | 160 | 1 | 7.6× |
| Lightning | 732 | 674.47 MB/s | 160 | 1 | 7.5× |
| Sonic | 1233 | 400.63 MB/s | 1031 | 6 | 4.5× |
| SonicFastest | 1238 | 399.09 MB/s | 1026 | 6 | 4.5× |
| LightningDecodeAny | 1437 | 343.18 MB/s | 1296 | 26 | 3.8× |
| Easyjson | 2200 | 224.50 MB/s | 448 | 3 | 2.5× |
| Goccy | 2443 | 202.21 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3243 | 152.32 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5524 | 89.43 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 422 | 545.34 MB/s | 160 | 1 | 9.8× |
| Lightning | 422 | 544.80 MB/s | 160 | 1 | 9.8× |
| Sonic | 870 | 264.31 MB/s | 652 | 6 | 4.7× |
| SonicFastest | 874 | 263.30 MB/s | 658 | 6 | 4.7× |
| LightningDecodeAny | 1121 | 204.25 MB/s | 1296 | 26 | 3.7× |
| Easyjson | 1383 | 166.32 MB/s | 448 | 3 | 3.0× |
| Goccy | 1570 | 146.52 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2443 | 94.13 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4128 | 55.71 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 68931 | 944.88 MB/s | 158660 | 100 | 7.9× |
| Lightning | 69739 | 933.94 MB/s | 164880 | 105 | 7.8× |
| SonicFastest | 96125 | 677.58 MB/s | 155213 | 75 | 5.6× |
| Sonic | 96346 | 676.02 MB/s | 155366 | 75 | 5.6× |
| Goccy | 140034 | 465.12 MB/s | 228907 | 134 | 3.9× |
| LightningDecodeAny | 182130 | 292.81 MB/s | 180224 | 3245 | 3.0× |
| JSONV2 | 221057 | 294.64 MB/s | 206648 | 607 | 2.5× |
| Stdlib | 542906 | 119.97 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2519251 | 770.26 MB/s | 2813905 | 1358 | 9.3× |
| Lightning | 2576218 | 753.23 MB/s | 2813906 | 1358 | 9.1× |
| SonicFastest | 4731849 | 410.09 MB/s | 14608604 | 1407 | 5.0× |
| Goccy | 4743504 | 409.08 MB/s | 4065523 | 13510 | 4.9× |
| Sonic | 4786922 | 405.37 MB/s | 14608603 | 1407 | 4.9× |
| Easyjson | 7502220 | 258.65 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9602142 | 202.09 MB/s | 7064790 | 218633 | 2.4× |
| JSONV2 | 11106843 | 174.71 MB/s | 3237227 | 13947 | 2.1× |
| Stdlib | 23439325 | 82.79 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1102249 | 3019.13 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1748731 | 1903.00 MB/s | 2488906 | 2995 | 11.9× |
| Sonic | 2641484 | 1259.83 MB/s | 6444455 | 4248 | 7.8× |
| SonicFastest | 2650873 | 1255.37 MB/s | 6442086 | 4248 | 7.8× |
| LightningDecodeAny | 3672129 | 837.05 MB/s | 4886620 | 56892 | 5.6× |
| Goccy | 4629810 | 718.78 MB/s | 3948909 | 3816 | 4.5× |
| JSONV2 | 7390792 | 450.27 MB/s | 5364510 | 13243 | 2.8× |
| Stdlib | 20732369 | 160.51 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220666 | 998.55 MB/s | 135872 | 226 | 9.2× |
| LightningDestructive | 221875 | 993.11 MB/s | 135872 | 226 | 9.2× |
| Sonic | 375961 | 586.09 MB/s | 306065 | 398 | 5.4× |
| SonicFastest | 385018 | 572.30 MB/s | 330848 | 398 | 5.3× |
| Goccy | 429482 | 513.05 MB/s | 364931 | 1067 | 4.7× |
| Easyjson | 548169 | 401.97 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 737477 | 298.78 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 864350 | 125.31 MB/s | 897522 | 11703 | 2.4× |
| Stdlib | 2038989 | 108.07 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12150098 | 666.66 MB/s | 15730640 | 20821 | 7.3× |
| Lightning | 12652706 | 640.18 MB/s | 15730646 | 20821 | 7.0× |
| Sonic | 16352386 | 495.34 MB/s | 70873026 | 40014 | 5.5× |
| SonicFastest | 16516887 | 490.41 MB/s | 70886998 | 40014 | 5.4× |
| Goccy | 23125718 | 350.26 MB/s | 17057115 | 107148 | 3.9× |
| Easyjson | 30452920 | 265.99 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 37237150 | 139.73 MB/s | 46191127 | 747112 | 2.4× |
| JSONV2 | 43692184 | 185.39 MB/s | 15233739 | 78972 | 2.0× |
| Stdlib | 89139127 | 90.87 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5619239 | 530.94 MB/s | 3908872 | 29356 | 8.4× |
| Lightning | 5722657 | 521.34 MB/s | 3908875 | 29356 | 8.2× |
| Sonic | 8584933 | 347.52 MB/s | 26719995 | 56760 | 5.5× |
| SonicFastest | 8586084 | 347.48 MB/s | 26760069 | 56760 | 5.5× |
| Easyjson | 16482536 | 181.01 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16630912 | 179.39 MB/s | 10716982 | 273652 | 2.8× |
| LightningDecodeAny | 16753293 | 109.48 MB/s | 23982394 | 351152 | 2.8× |
| JSONV2 | 23913492 | 124.76 MB/s | 9257157 | 86278 | 2.0× |
| Stdlib | 47085784 | 63.36 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1231537 | 587.56 MB/s | 907601 | 3618 | 9.5× |
| Lightning | 1242363 | 582.44 MB/s | 907596 | 3618 | 9.4× |
| Sonic | 1766002 | 409.74 MB/s | 3181556 | 7226 | 6.6× |
| SonicFastest | 1769326 | 408.97 MB/s | 3198643 | 7226 | 6.6× |
| LightningDecodeAny | 4121479 | 157.85 MB/s | 6500456 | 76546 | 2.8× |
| Easyjson | 4222288 | 171.38 MB/s | 2847906 | 3698 | 2.8× |
| Goccy | 4712124 | 153.56 MB/s | 2842910 | 80275 | 2.5× |
| JSONV2 | 5964238 | 121.32 MB/s | 2704643 | 7318 | 2.0× |
| Stdlib | 11643870 | 62.14 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1841704 | 856.46 MB/s | 907602 | 3618 | 8.5× |
| Lightning | 1881890 | 838.17 MB/s | 907595 | 3618 | 8.3× |
| Sonic | 2248806 | 701.42 MB/s | 5785257 | 7226 | 7.0× |
| SonicFastest | 2260942 | 697.65 MB/s | 5802163 | 7226 | 7.0× |
| LightningDecodeAny | 3826577 | 196.89 MB/s | 6500460 | 76546 | 4.1× |
| Easyjson | 5553029 | 284.05 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5687799 | 277.32 MB/s | 3631588 | 80269 | 2.8× |
| JSONV2 | 6694729 | 235.61 MB/s | 2704591 | 7318 | 2.3× |
| Stdlib | 15713681 | 100.38 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 207368 | 723.95 MB/s | 81920 | 1 | 8.8× |
| Lightning | 207590 | 723.18 MB/s | 81920 | 1 | 8.8× |
| SonicFastest | 269744 | 556.54 MB/s | 256399 | 6 | 6.7× |
| Sonic | 271001 | 553.96 MB/s | 258173 | 6 | 6.7× |
| LightningDecodeAny | 468097 | 320.71 MB/s | 745765 | 10016 | 3.9× |
| Goccy | 859936 | 174.58 MB/s | 324118 | 10004 | 2.1× |
| JSONV2 | 1097875 | 136.74 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1818996 | 82.53 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32090 | 876.20 MB/s | 29216 | 103 | 9.5× |
| LightningDestructive | 32178 | 873.79 MB/s | 29088 | 101 | 9.4× |
| Sonic | 62608 | 449.09 MB/s | 46602 | 103 | 4.8× |
| SonicFastest | 62876 | 447.18 MB/s | 46875 | 103 | 4.8× |
| Easyjson | 68047 | 413.20 MB/s | 32304 | 138 | 4.5× |
| Goccy | 70720 | 397.58 MB/s | 59209 | 188 | 4.3× |
| JSONV2 | 134201 | 209.51 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 147666 | 190.41 MB/s | 140592 | 2643 | 2.1× |
| Stdlib | 303532 | 92.63 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1966 | 1184.25 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2025 | 1149.62 MB/s | 32 | 1 | 11.2× |
| Goccy | 4104 | 567.25 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4224 | 551.15 MB/s | 192 | 2 | 5.4× |
| Sonic | 5002 | 465.43 MB/s | 4213 | 6 | 4.5× |
| SonicFastest | 5007 | 464.99 MB/s | 4263 | 6 | 4.5× |
| JSONV2 | 8456 | 275.30 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10099 | 166.84 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22731 | 102.42 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221 | 854.56 MB/s | 0 | 0 | 11.2× |
| LightningDestructive | 223 | 846.68 MB/s | 0 | 0 | 11.1× |
| Goccy | 384 | 491.63 MB/s | 304 | 2 | 6.4× |
| Easyjson | 498 | 379.91 MB/s | 0 | 0 | 5.0× |
| Sonic | 789 | 239.43 MB/s | 500 | 4 | 3.1× |
| SonicFastest | 793 | 238.42 MB/s | 498 | 4 | 3.1× |
| JSONV2 | 1040 | 181.72 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1212 | 110.61 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2475 | 76.36 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1559 | 1405.60 MB/s | 0 | 0 | 10.3× |
| LightningDestructive | 1575 | 1390.85 MB/s | 0 | 0 | 10.2× |
| Goccy | 3186 | 687.79 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3195 | 685.79 MB/s | 24 | 1 | 5.0× |
| Sonic | 6301 | 347.71 MB/s | 3938 | 40 | 2.5× |
| SonicFastest | 6325 | 346.42 MB/s | 4014 | 40 | 2.5× |
| JSONV2 | 7968 | 274.98 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8086 | 223.97 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 16031 | 136.67 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 678334 | 752.54 MB/s | 703297 | 1010 | 8.9× |
| Lightning | 724201 | 704.88 MB/s | 703302 | 1010 | 8.3× |
| Goccy | 1131768 | 451.04 MB/s | 1138709 | 5006 | 5.3× |
| SonicFastest | 1147400 | 444.90 MB/s | 885941 | 2006 | 5.3× |
| Sonic | 1155113 | 441.93 MB/s | 902520 | 2006 | 5.2× |
| Easyjson | 1512535 | 337.50 MB/s | 863779 | 3012 | 4.0× |
| JSONV2 | 3171709 | 160.95 MB/s | 1076011 | 12646 | 1.9× |
| LightningDecodeAny | 3328381 | 138.65 MB/s | 2929688 | 64018 | 1.8× |
| Stdlib | 6026140 | 84.71 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1337 | 14802.14 MB/s | 0 | 0 | 83.5× |
| LightningDestructive | 1361 | 14535.87 MB/s | 0 | 0 | 82.0× |
| Goccy | 19774 | 1000.73 MB/s | 20491 | 2 | 5.6× |
| SonicFastest | 27660 | 715.43 MB/s | 21792 | 4 | 4.0× |
| Sonic | 27684 | 714.81 MB/s | 21897 | 4 | 4.0× |
| JSONV2 | 29541 | 669.89 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 73427 | 269.49 MB/s | 116864 | 2015 | 1.5× |
| Easyjson | 86030 | 230.03 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111670 | 177.21 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2659 | 6815.31 MB/s | 0 | 0 | 38.7× |
| Lightning | 2835 | 6392.08 MB/s | 432 | 2 | 36.3× |
| Easyjson | 3984 | 4549.69 MB/s | 432 | 2 | 25.9× |
| Sonic | 9712 | 1866.19 MB/s | 22766 | 6 | 10.6× |
| SonicFastest | 9791 | 1851.12 MB/s | 22972 | 6 | 10.5× |
| Goccy | 15988 | 1133.58 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 16401 | 1090.31 MB/s | 29088 | 191 | 6.3× |
| JSONV2 | 44737 | 405.13 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103003 | 175.96 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2309933 | 869.50 MB/s | 3089820 | 6821 | 8.1× |
| Lightning | 2369132 | 847.78 MB/s | 3091533 | 6827 | 7.9× |
| Goccy | 4111357 | 488.52 MB/s | 5411403 | 15830 | 4.5× |
| Sonic | 4296676 | 467.45 MB/s | 10975105 | 13683 | 4.3× |
| SonicFastest | 4327453 | 464.13 MB/s | 10944842 | 13683 | 4.3× |
| Easyjson | 4847936 | 414.30 MB/s | 2981482 | 7439 | 3.8× |
| JSONV2 | 6849659 | 293.23 MB/s | 3173684 | 14563 | 2.7× |
| LightningDecodeAny | 7421067 | 153.93 MB/s | 8498330 | 134008 | 2.5× |
| Stdlib | 18619482 | 107.87 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 875 | 627.32 MB/s | 480 | 1 | 6.4× |
| LightningDestructive | 886 | 619.57 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1665 | 329.15 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 2116 | 259.40 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2633 | 208.50 MB/s | 1939 | 26 | 2.1× |
| SonicFastest | 2633 | 208.51 MB/s | 1934 | 26 | 2.1× |
| Goccy | 2958 | 185.58 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3240 | 169.47 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5584 | 98.32 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 499694 | 1263.80 MB/s | 402729 | 545 | 10.8× |
| Lightning | 560140 | 1127.42 MB/s | 451257 | 857 | 9.6× |
| Sonic | 991878 | 636.69 MB/s | 1004428 | 1102 | 5.4× |
| SonicFastest | 995908 | 634.11 MB/s | 1021822 | 1102 | 5.4× |
| Easyjson | 1140274 | 553.83 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1149755 | 549.26 MB/s | 987500 | 1202 | 4.7× |
| JSONV2 | 2137798 | 295.40 MB/s | 571619 | 3144 | 2.5× |
| LightningDecodeAny | 2339820 | 199.55 MB/s | 2077369 | 30126 | 2.3× |
| Stdlib | 5378162 | 117.42 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 663248 | 847.96 MB/s | 579337 | 429 | 7.9× |
| Lightning | 829807 | 677.76 MB/s | 802710 | 1235 | 6.3× |
| Sonic | 995057 | 565.20 MB/s | 946942 | 1476 | 5.3× |
| SonicFastest | 1013799 | 554.75 MB/s | 977842 | 1476 | 5.2× |
| Goccy | 1324875 | 424.50 MB/s | 1040543 | 1030 | 4.0× |
| Easyjson | 1732283 | 324.66 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2663739 | 211.13 MB/s | 2181320 | 30126 | 2.0× |
| JSONV2 | 2737879 | 205.42 MB/s | 927439 | 3482 | 1.9× |
| Stdlib | 5249094 | 107.14 MB/s | 1011674 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 644969 | 826.67 MB/s | 333416 | 2084 | 8.5× |
| Lightning | 666803 | 799.60 MB/s | 368224 | 2293 | 8.2× |
| Easyjson | 1099278 | 485.03 MB/s | 428361 | 3273 | 5.0× |
| Sonic | 1127095 | 473.06 MB/s | 1055027 | 4351 | 4.8× |
| SonicFastest | 1131581 | 471.18 MB/s | 1056708 | 4351 | 4.8× |
| Goccy | 1278720 | 416.96 MB/s | 1167246 | 5409 | 4.3× |
| JSONV2 | 2520858 | 211.51 MB/s | 745450 | 13288 | 2.2× |
| LightningDecodeAny | 3347045 | 159.30 MB/s | 2991147 | 50076 | 1.6× |
| Stdlib | 5457425 | 97.70 MB/s | 798693 | 17133 | 1.0× |
