# JSON Deserialization Benchmarks

- generated 2026-06-21T21:08:48Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104489 | 1218.07 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 104976 | 1212.42 MB/s | 49280 | 2 | 10.5× |
| Sonic | 178652 | 712.42 MB/s | 190081 | 10 | 6.2× |
| SonicFastest | 179794 | 707.89 MB/s | 192744 | 10 | 6.2× |
| Goccy | 195556 | 650.84 MB/s | 225237 | 884 | 5.7× |
| Easyjson | 213656 | 595.70 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 429669 | 296.22 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 458740 | 206.33 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1106620 | 115.01 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3745224 | 601.05 MB/s | 3122872 | 3081 | 7.0× |
| Lightning | 3824449 | 588.59 MB/s | 3122876 | 3081 | 6.8× |
| Sonic | 4551733 | 494.55 MB/s | 15238736 | 970 | 5.7× |
| SonicFastest | 4607884 | 488.52 MB/s | 15232102 | 970 | 5.7× |
| Goccy | 10511948 | 214.14 MB/s | 4124532 | 56532 | 2.5× |
| Easyjson | 11176934 | 201.40 MB/s | 3099808 | 2120 | 2.3× |
| LightningDecodeAny | 11938017 | 188.56 MB/s | 7938298 | 281383 | 2.2× |
| JSONV2 | 15893585 | 141.63 MB/s | 3123221 | 3083 | 1.6× |
| Stdlib | 26116216 | 86.19 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 490186 | 551.63 MB/s | 348024 | 1627 | 6.9× |
| LightningDestructive | 499342 | 541.52 MB/s | 348024 | 1627 | 6.7× |
| Sonic | 635124 | 425.75 MB/s | 491988 | 968 | 5.3× |
| SonicFastest | 640863 | 421.94 MB/s | 488083 | 968 | 5.2× |
| Easyjson | 1402399 | 192.81 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1447082 | 186.86 MB/s | 544122 | 8123 | 2.3× |
| LightningDecodeAny | 1703996 | 158.69 MB/s | 1011486 | 37901 | 2.0× |
| JSONV2 | 2109822 | 128.16 MB/s | 348149 | 1628 | 1.6× |
| Stdlib | 3359083 | 80.50 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1358650 | 1271.26 MB/s | 959848 | 5881 | 9.8× |
| Lightning | 1379377 | 1252.16 MB/s | 959890 | 5882 | 9.6× |
| Sonic | 2050818 | 842.20 MB/s | 2717020 | 4020 | 6.5× |
| SonicFastest | 2054485 | 840.70 MB/s | 2760724 | 4020 | 6.5× |
| Goccy | 2389222 | 722.91 MB/s | 2582796 | 14604 | 5.6× |
| Easyjson | 4261420 | 405.31 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4345507 | 397.47 MB/s | 1011634 | 7594 | 3.1× |
| LightningDecodeAny | 4704903 | 106.34 MB/s | 4976206 | 81466 | 2.8× |
| Stdlib | 13282810 | 130.03 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1186 | 1527.47 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1203 | 1506.53 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2553 | 709.76 MB/s | 24 | 1 | 5.5× |
| Goccy | 2795 | 648.31 MB/s | 2608 | 4 | 5.0× |
| Sonic | 6003 | 301.87 MB/s | 3696 | 40 | 2.3× |
| SonicFastest | 6073 | 298.37 MB/s | 3759 | 40 | 2.3× |
| JSONV2 | 7699 | 235.35 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8262 | 219.18 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14044 | 129.02 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1217 | 1488.52 MB/s | 0 | 0 | 11.6× |
| LightningDestructive | 1232 | 1471.02 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2550 | 710.67 MB/s | 24 | 1 | 5.5× |
| Goccy | 2796 | 648.14 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6027 | 300.67 MB/s | 3694 | 40 | 2.3× |
| Sonic | 6069 | 298.56 MB/s | 3730 | 40 | 2.3× |
| JSONV2 | 7811 | 231.97 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8242 | 219.73 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14098 | 128.53 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1397 | 1297.37 MB/s | 144 | 10 | 10.1× |
| LightningDestructive | 1417 | 1278.46 MB/s | 144 | 10 | 9.9× |
| Easyjson | 2749 | 659.26 MB/s | 144 | 10 | 5.1× |
| Goccy | 2897 | 625.48 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6066 | 298.71 MB/s | 3745 | 42 | 2.3× |
| SonicFastest | 6106 | 296.76 MB/s | 3800 | 42 | 2.3× |
| JSONV2 | 8007 | 226.30 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8264 | 219.15 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14061 | 128.87 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 740 | 667.25 MB/s | 160 | 1 | 7.5× |
| Lightning | 746 | 662.15 MB/s | 160 | 1 | 7.5× |
| Sonic | 1241 | 398.01 MB/s | 979 | 6 | 4.5× |
| SonicFastest | 1242 | 397.80 MB/s | 989 | 6 | 4.5× |
| LightningDecodeAny | 1665 | 296.16 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2223 | 222.26 MB/s | 448 | 3 | 2.5× |
| Goccy | 2433 | 203.04 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3281 | 150.58 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5571 | 88.67 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 478 | 481.40 MB/s | 160 | 1 | 8.7× |
| Lightning | 479 | 480.47 MB/s | 160 | 1 | 8.7× |
| Sonic | 906 | 253.93 MB/s | 693 | 6 | 4.6× |
| SonicFastest | 906 | 253.84 MB/s | 693 | 6 | 4.6× |
| Easyjson | 1409 | 163.26 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1469 | 155.93 MB/s | 1536 | 30 | 2.8× |
| Goccy | 1610 | 142.83 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2486 | 92.50 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4168 | 55.18 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 77201 | 843.67 MB/s | 172432 | 107 | 7.1× |
| LightningDestructive | 80252 | 811.60 MB/s | 166212 | 102 | 6.8× |
| Sonic | 97495 | 668.05 MB/s | 157134 | 75 | 5.6× |
| SonicFastest | 99683 | 653.39 MB/s | 157894 | 75 | 5.5× |
| Goccy | 145030 | 449.09 MB/s | 229219 | 134 | 3.8× |
| LightningDecodeAny | 189569 | 281.32 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 222401 | 292.86 MB/s | 206651 | 607 | 2.5× |
| Stdlib | 548854 | 118.67 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2608493 | 743.91 MB/s | 2846864 | 2238 | 9.0× |
| Lightning | 2676012 | 725.14 MB/s | 2846868 | 2238 | 8.8× |
| Sonic | 4653138 | 417.02 MB/s | 14608593 | 1407 | 5.0× |
| SonicFastest | 4725996 | 410.60 MB/s | 14608656 | 1407 | 5.0× |
| Goccy | 4775091 | 406.37 MB/s | 4065330 | 13510 | 4.9× |
| Easyjson | 7600197 | 255.32 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9681668 | 200.43 MB/s | 7013524 | 219937 | 2.4× |
| JSONV2 | 11250413 | 172.48 MB/s | 3237215 | 13947 | 2.1× |
| Stdlib | 23496862 | 82.58 MB/s | 3551325 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1112877 | 2990.30 MB/s | 351704 | 1286 | 18.7× |
| Lightning | 1783207 | 1866.21 MB/s | 2488905 | 2995 | 11.7× |
| SonicFastest | 2671337 | 1245.75 MB/s | 6472819 | 4248 | 7.8× |
| Sonic | 2672718 | 1245.11 MB/s | 6488607 | 4248 | 7.8× |
| LightningDecodeAny | 3783997 | 812.31 MB/s | 4886620 | 56892 | 5.5× |
| Goccy | 4509553 | 737.95 MB/s | 3948909 | 3816 | 4.6× |
| JSONV2 | 7458327 | 446.19 MB/s | 5364515 | 13243 | 2.8× |
| Stdlib | 20791887 | 160.05 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 228789 | 963.10 MB/s | 136896 | 228 | 8.9× |
| LightningDestructive | 230535 | 955.80 MB/s | 136897 | 228 | 8.8× |
| SonicFastest | 378489 | 582.17 MB/s | 298555 | 398 | 5.4× |
| Sonic | 386072 | 570.74 MB/s | 329633 | 398 | 5.3× |
| Goccy | 432977 | 508.91 MB/s | 365021 | 1067 | 4.7× |
| Easyjson | 549496 | 401.00 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 720130 | 305.98 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 889875 | 121.72 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2039025 | 108.06 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12782421 | 633.69 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13285762 | 609.68 MB/s | 15059841 | 51649 | 6.7× |
| SonicFastest | 16718875 | 484.48 MB/s | 70916256 | 40014 | 5.4× |
| Sonic | 16747309 | 483.66 MB/s | 70887359 | 40014 | 5.4× |
| Goccy | 23687738 | 341.95 MB/s | 17027813 | 107148 | 3.8× |
| Easyjson | 30695224 | 263.89 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 42057627 | 123.71 MB/s | 34333351 | 912810 | 2.1× |
| JSONV2 | 43416602 | 186.57 MB/s | 15233738 | 78972 | 2.1× |
| Stdlib | 89629029 | 90.37 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6120351 | 487.47 MB/s | 3985336 | 30015 | 7.7× |
| Lightning | 6235347 | 478.48 MB/s | 3985336 | 30015 | 7.6× |
| Sonic | 8664690 | 344.32 MB/s | 26595816 | 56760 | 5.4× |
| SonicFastest | 8728335 | 341.81 MB/s | 26572324 | 56760 | 5.4× |
| Goccy | 16582334 | 179.92 MB/s | 10542746 | 273646 | 2.8× |
| Easyjson | 16620889 | 179.50 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 19416537 | 94.47 MB/s | 20023839 | 408557 | 2.4× |
| JSONV2 | 24334518 | 122.60 MB/s | 9257137 | 86278 | 1.9× |
| Stdlib | 47100066 | 63.34 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1376625 | 525.63 MB/s | 907393 | 3618 | 8.4× |
| Lightning | 1391271 | 520.10 MB/s | 907388 | 3618 | 8.3× |
| Sonic | 1786142 | 405.12 MB/s | 3175631 | 7226 | 6.5× |
| SonicFastest | 1789989 | 404.25 MB/s | 3181985 | 7226 | 6.5× |
| Easyjson | 4279567 | 169.08 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4370154 | 148.87 MB/s | 5752202 | 80172 | 2.7× |
| Goccy | 4761025 | 151.98 MB/s | 2763165 | 80271 | 2.4× |
| JSONV2 | 5723626 | 126.42 MB/s | 2704630 | 7318 | 2.0× |
| Stdlib | 11602633 | 62.36 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1927411 | 818.38 MB/s | 907393 | 3618 | 8.2× |
| Lightning | 1974979 | 798.67 MB/s | 907386 | 3618 | 8.0× |
| SonicFastest | 2261714 | 697.41 MB/s | 5796993 | 7226 | 7.0× |
| Sonic | 2281629 | 691.33 MB/s | 5788889 | 7226 | 6.9× |
| LightningDecodeAny | 3851818 | 195.60 MB/s | 5752201 | 80172 | 4.1× |
| Goccy | 5657111 | 278.83 MB/s | 3552830 | 80266 | 2.8× |
| Easyjson | 5669897 | 278.20 MB/s | 2847908 | 3698 | 2.8× |
| JSONV2 | 6324991 | 249.38 MB/s | 2704593 | 7318 | 2.5× |
| Stdlib | 15741905 | 100.20 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 224178 | 669.66 MB/s | 81920 | 1 | 8.1× |
| Lightning | 224301 | 669.30 MB/s | 81920 | 1 | 8.1× |
| Sonic | 274019 | 547.86 MB/s | 262020 | 6 | 6.7× |
| SonicFastest | 274400 | 547.10 MB/s | 260242 | 6 | 6.6× |
| LightningDecodeAny | 478397 | 313.80 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 861183 | 174.32 MB/s | 324964 | 10004 | 2.1× |
| JSONV2 | 1095809 | 137.00 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1823554 | 82.32 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 35487 | 792.32 MB/s | 30272 | 105 | 8.6× |
| LightningDestructive | 35634 | 789.05 MB/s | 30144 | 103 | 8.5× |
| SonicFastest | 63891 | 440.08 MB/s | 47063 | 103 | 4.8× |
| Sonic | 63913 | 439.93 MB/s | 46970 | 103 | 4.8× |
| Easyjson | 68998 | 407.50 MB/s | 32304 | 138 | 4.4× |
| Goccy | 71035 | 395.82 MB/s | 59216 | 188 | 4.3× |
| JSONV2 | 134939 | 208.37 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 151808 | 185.21 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 304205 | 92.43 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1957 | 1189.53 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2008 | 1159.15 MB/s | 32 | 1 | 11.3× |
| Goccy | 4107 | 566.90 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4219 | 551.82 MB/s | 192 | 2 | 5.4× |
| Sonic | 5113 | 455.35 MB/s | 4288 | 6 | 4.4× |
| SonicFastest | 5120 | 454.71 MB/s | 4297 | 6 | 4.4× |
| JSONV2 | 8486 | 274.34 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10313 | 163.39 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22689 | 102.61 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 862.39 MB/s | 0 | 0 | 11.2× |
| LightningDestructive | 221 | 855.39 MB/s | 0 | 0 | 11.2× |
| Goccy | 381 | 496.12 MB/s | 304 | 2 | 6.5× |
| Easyjson | 493 | 383.19 MB/s | 0 | 0 | 5.0× |
| Sonic | 801 | 235.91 MB/s | 499 | 4 | 3.1× |
| SonicFastest | 806 | 234.61 MB/s | 504 | 4 | 3.1× |
| JSONV2 | 1037 | 182.34 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1252 | 107.06 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2465 | 76.67 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1525 | 1436.45 MB/s | 0 | 0 | 10.5× |
| LightningDestructive | 1536 | 1426.42 MB/s | 0 | 0 | 10.4× |
| Easyjson | 3187 | 687.56 MB/s | 24 | 1 | 5.0× |
| Goccy | 3191 | 686.60 MB/s | 2864 | 4 | 5.0× |
| Sonic | 6505 | 336.79 MB/s | 4001 | 40 | 2.5× |
| SonicFastest | 6545 | 334.75 MB/s | 4014 | 40 | 2.4× |
| JSONV2 | 8067 | 271.62 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8269 | 219.00 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16026 | 136.72 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 713921 | 715.03 MB/s | 703778 | 1012 | 8.5× |
| Lightning | 744552 | 685.61 MB/s | 703778 | 1012 | 8.2× |
| Sonic | 1159865 | 440.12 MB/s | 913223 | 2006 | 5.2× |
| Goccy | 1164059 | 438.53 MB/s | 1138373 | 5006 | 5.2× |
| SonicFastest | 1165678 | 437.92 MB/s | 908187 | 2006 | 5.2× |
| Easyjson | 1540200 | 331.43 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3240550 | 157.53 MB/s | 1076005 | 12646 | 1.9× |
| LightningDecodeAny | 3566194 | 129.40 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6077495 | 83.99 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1334 | 14832.19 MB/s | 0 | 0 | 83.7× |
| LightningDestructive | 1360 | 14555.13 MB/s | 0 | 0 | 82.1× |
| Goccy | 19953 | 991.79 MB/s | 20491 | 2 | 5.6× |
| Sonic | 27372 | 722.98 MB/s | 22461 | 4 | 4.1× |
| SonicFastest | 27449 | 720.95 MB/s | 22703 | 4 | 4.1× |
| JSONV2 | 29701 | 666.28 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 78952 | 250.63 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86034 | 230.01 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111614 | 177.30 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2652 | 6833.79 MB/s | 0 | 0 | 38.9× |
| Lightning | 2852 | 6354.67 MB/s | 432 | 2 | 36.2× |
| Easyjson | 3968 | 4567.69 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 10002 | 1812.09 MB/s | 23143 | 6 | 10.3× |
| Sonic | 10161 | 1783.65 MB/s | 23292 | 6 | 10.2× |
| Goccy | 15823 | 1145.42 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 16942 | 1055.47 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 46037 | 393.68 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 103255 | 175.53 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2314858 | 867.65 MB/s | 2960388 | 7411 | 8.0× |
| Lightning | 2350810 | 854.38 MB/s | 2962102 | 7417 | 7.9× |
| Goccy | 4141274 | 484.99 MB/s | 5412085 | 15831 | 4.5× |
| Sonic | 4404087 | 456.05 MB/s | 10921852 | 13683 | 4.2× |
| SonicFastest | 4496176 | 446.71 MB/s | 10906444 | 13683 | 4.1× |
| Easyjson | 4905051 | 409.47 MB/s | 2981483 | 7439 | 3.8× |
| JSONV2 | 6881375 | 291.87 MB/s | 3173688 | 14563 | 2.7× |
| LightningDecodeAny | 7440978 | 153.51 MB/s | 7386073 | 134751 | 2.5× |
| Stdlib | 18590291 | 108.04 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 885 | 620.29 MB/s | 480 | 1 | 6.4× |
| LightningDestructive | 901 | 609.17 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1925 | 284.62 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2161 | 254.04 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2688 | 204.27 MB/s | 1947 | 26 | 2.1× |
| SonicFastest | 2689 | 204.18 MB/s | 1953 | 26 | 2.1× |
| Goccy | 3014 | 182.14 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3271 | 167.82 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5646 | 97.24 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 497847 | 1268.49 MB/s | 364472 | 566 | 10.9× |
| Lightning | 569346 | 1109.19 MB/s | 413001 | 878 | 9.5× |
| SonicFastest | 1010144 | 625.17 MB/s | 996462 | 1102 | 5.4× |
| Sonic | 1022449 | 617.65 MB/s | 998350 | 1102 | 5.3× |
| Easyjson | 1148656 | 549.78 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1164310 | 542.39 MB/s | 983991 | 1201 | 4.6× |
| JSONV2 | 2136287 | 295.61 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2394879 | 194.96 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5408063 | 116.77 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 712267 | 789.60 MB/s | 544252 | 448 | 7.4× |
| Lightning | 895060 | 628.35 MB/s | 767621 | 1254 | 5.9× |
| SonicFastest | 1018970 | 551.94 MB/s | 949932 | 1476 | 5.2× |
| Sonic | 1038072 | 541.78 MB/s | 971337 | 1476 | 5.1× |
| Goccy | 1331798 | 422.29 MB/s | 1039273 | 1030 | 4.0× |
| Easyjson | 1754041 | 320.64 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2731302 | 205.91 MB/s | 2114151 | 30295 | 1.9× |
| JSONV2 | 2749959 | 204.52 MB/s | 927441 | 3482 | 1.9× |
| Stdlib | 5291758 | 106.28 MB/s | 1011671 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 653197 | 816.26 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 683253 | 780.35 MB/s | 368224 | 2293 | 8.0× |
| Easyjson | 1110321 | 480.20 MB/s | 428361 | 3273 | 5.0× |
| SonicFastest | 1155498 | 461.43 MB/s | 1052510 | 4351 | 4.8× |
| Sonic | 1157223 | 460.74 MB/s | 1035933 | 4351 | 4.8× |
| Goccy | 1295004 | 411.72 MB/s | 1167230 | 5409 | 4.2× |
| JSONV2 | 2551591 | 208.96 MB/s | 745447 | 13288 | 2.2× |
| LightningDecodeAny | 3423106 | 155.76 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5497557 | 96.98 MB/s | 798692 | 17133 | 1.0× |
