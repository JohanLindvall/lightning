# JSON Deserialization Benchmarks

- generated 2026-06-21T07:31:14Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 101030 | 1259.77 MB/s | 49760 | 3 | 12.2× |
| Easyjson | 220401 | 577.47 MB/s | 122864 | 14 | 5.6× |
| SonicFastest | 229427 | 554.75 MB/s | 214106 | 15 | 5.4× |
| Sonic | 230121 | 553.08 MB/s | 214184 | 15 | 5.4× |
| Goccy | 246754 | 515.80 MB/s | 225328 | 884 | 5.0× |
| JSONV2 | 408289 | 311.73 MB/s | 195128 | 1805 | 3.0× |
| LightningDecodeAny | 441166 | 214.55 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1234826 | 103.07 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4173880 | 539.32 MB/s | 3122874 | 3081 | 7.4× |
| SonicFastest | 5060388 | 444.84 MB/s | 4866231 | 2584 | 6.1× |
| Sonic | 5203437 | 432.61 MB/s | 4865921 | 2584 | 6.0× |
| LightningDecodeAny | 11550354 | 194.89 MB/s | 7938300 | 281383 | 2.7× |
| Goccy | 12884301 | 174.71 MB/s | 4180575 | 56534 | 2.4× |
| Easyjson | 13739317 | 163.84 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 16714150 | 134.68 MB/s | 3123189 | 3083 | 1.9× |
| Stdlib | 31076801 | 72.44 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 561692 | 481.41 MB/s | 348024 | 1627 | 7.2× |
| SonicFastest | 754119 | 358.57 MB/s | 643957 | 1147 | 5.3× |
| Sonic | 760990 | 355.33 MB/s | 644330 | 1147 | 5.3× |
| LightningDecodeAny | 1604661 | 168.51 MB/s | 1011487 | 37901 | 2.5× |
| Easyjson | 1752408 | 154.30 MB/s | 330272 | 749 | 2.3× |
| Goccy | 1780816 | 151.84 MB/s | 545213 | 8123 | 2.3× |
| JSONV2 | 2227798 | 121.38 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4027710 | 67.14 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1324816 | 1303.73 MB/s | 959890 | 5882 | 12.7× |
| SonicFastest | 2163105 | 798.48 MB/s | 2693209 | 5547 | 7.8× |
| Sonic | 2170222 | 795.87 MB/s | 2692861 | 5547 | 7.7× |
| Goccy | 2851001 | 605.82 MB/s | 2580989 | 14603 | 5.9× |
| JSONV2 | 3884624 | 444.63 MB/s | 1011615 | 7594 | 4.3× |
| Easyjson | 4038722 | 427.66 MB/s | 972032 | 5389 | 4.2× |
| LightningDecodeAny | 4051110 | 123.50 MB/s | 4976204 | 81466 | 4.1× |
| Stdlib | 16808419 | 102.76 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1055 | 1718.16 MB/s | 0 | 0 | 13.9× |
| Easyjson | 2785 | 650.67 MB/s | 24 | 1 | 5.3× |
| Goccy | 3288 | 551.03 MB/s | 2608 | 4 | 4.5× |
| SonicFastest | 6099 | 297.11 MB/s | 3347 | 38 | 2.4× |
| Sonic | 6272 | 288.89 MB/s | 3345 | 38 | 2.3× |
| JSONV2 | 7732 | 234.35 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 8596 | 210.67 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14687 | 123.37 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1115 | 1625.49 MB/s | 0 | 0 | 13.0× |
| Easyjson | 2759 | 656.69 MB/s | 24 | 1 | 5.3× |
| Goccy | 3277 | 552.95 MB/s | 2608 | 4 | 4.4× |
| SonicFastest | 5880 | 308.14 MB/s | 3349 | 38 | 2.5× |
| Sonic | 6132 | 295.51 MB/s | 3346 | 38 | 2.4× |
| JSONV2 | 7820 | 231.73 MB/s | 640 | 6 | 1.9× |
| LightningDecodeAny | 8630 | 209.86 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14536 | 124.66 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1287 | 1407.98 MB/s | 144 | 10 | 11.3× |
| Easyjson | 2975 | 609.12 MB/s | 144 | 10 | 4.9× |
| Goccy | 3306 | 548.17 MB/s | 2600 | 5 | 4.4× |
| SonicFastest | 6092 | 297.46 MB/s | 3367 | 40 | 2.4× |
| Sonic | 6279 | 288.56 MB/s | 3366 | 40 | 2.3× |
| JSONV2 | 8087 | 224.05 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8671 | 208.85 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14505 | 124.92 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 736 | 671.53 MB/s | 160 | 1 | 8.2× |
| Sonic | 1208 | 408.82 MB/s | 1075 | 8 | 5.0× |
| SonicFastest | 1210 | 408.27 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1613 | 305.59 MB/s | 1536 | 30 | 3.7× |
| Easyjson | 2319 | 213.02 MB/s | 448 | 3 | 2.6× |
| Goccy | 2412 | 204.78 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3093 | 159.72 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6019 | 82.07 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 463 | 496.39 MB/s | 160 | 1 | 9.0× |
| Sonic | 853 | 269.65 MB/s | 801 | 8 | 4.9× |
| SonicFastest | 855 | 268.89 MB/s | 801 | 8 | 4.9× |
| LightningDecodeAny | 1400 | 163.54 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1528 | 150.49 MB/s | 448 | 3 | 2.7× |
| Goccy | 1660 | 138.52 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2328 | 98.81 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4178 | 55.05 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 87095 | 747.83 MB/s | 172434 | 107 | 7.3× |
| SonicFastest | 152467 | 427.19 MB/s | 235794 | 65 | 4.2× |
| Sonic | 152985 | 425.74 MB/s | 235808 | 65 | 4.2× |
| LightningDecodeAny | 183315 | 290.91 MB/s | 176961 | 3252 | 3.5× |
| Goccy | 184451 | 353.11 MB/s | 228408 | 134 | 3.5× |
| JSONV2 | 249532 | 261.02 MB/s | 206666 | 607 | 2.6× |
| Stdlib | 638898 | 101.94 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2648421 | 732.69 MB/s | 2846866 | 2238 | 9.6× |
| SonicFastest | 4649904 | 417.31 MB/s | 4879806 | 1736 | 5.4× |
| Sonic | 4697645 | 413.07 MB/s | 4879912 | 1736 | 5.4× |
| Goccy | 4785163 | 405.52 MB/s | 4063673 | 13509 | 5.3× |
| Easyjson | 7786097 | 249.22 MB/s | 3871265 | 15043 | 3.3× |
| LightningDecodeAny | 9446054 | 205.43 MB/s | 7013524 | 219937 | 2.7× |
| JSONV2 | 10956507 | 177.11 MB/s | 3237191 | 13947 | 2.3× |
| Stdlib | 25321957 | 76.63 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1641435 | 2027.39 MB/s | 2488907 | 2995 | 15.6× |
| SonicFastest | 2292789 | 1451.43 MB/s | 5895854 | 4263 | 11.1× |
| Sonic | 2311326 | 1439.79 MB/s | 5895483 | 4263 | 11.0× |
| LightningDecodeAny | 3508583 | 876.07 MB/s | 4886621 | 56892 | 7.3× |
| Goccy | 5709660 | 582.84 MB/s | 3948913 | 3816 | 4.5× |
| JSONV2 | 7854733 | 423.67 MB/s | 5364501 | 13243 | 3.2× |
| Stdlib | 25525444 | 130.37 MB/s | 5565610 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 230942 | 954.12 MB/s | 136896 | 228 | 9.9× |
| Goccy | 454969 | 484.31 MB/s | 363789 | 1066 | 5.0× |
| Sonic | 516086 | 426.96 MB/s | 350527 | 262 | 4.4× |
| SonicFastest | 517459 | 425.82 MB/s | 350896 | 262 | 4.4× |
| Easyjson | 558901 | 394.25 MB/s | 130512 | 245 | 4.1× |
| JSONV2 | 684703 | 321.81 MB/s | 129746 | 470 | 3.3× |
| LightningDecodeAny | 947825 | 114.28 MB/s | 861393 | 11944 | 2.4× |
| Stdlib | 2293222 | 96.09 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 15475422 | 523.41 MB/s | 15059845 | 51649 | 6.6× |
| Sonic | 22054527 | 367.27 MB/s | 19853809 | 41640 | 4.7× |
| SonicFastest | 22305686 | 363.14 MB/s | 19854093 | 41640 | 4.6× |
| Goccy | 26177481 | 309.43 MB/s | 19114290 | 107156 | 3.9× |
| Easyjson | 33453165 | 242.13 MB/s | 15059620 | 41643 | 3.1× |
| LightningDecodeAny | 41769486 | 124.57 MB/s | 34333349 | 912810 | 2.5× |
| JSONV2 | 45399383 | 178.42 MB/s | 15233720 | 78972 | 2.3× |
| Stdlib | 102584261 | 78.96 MB/s | 15665072 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6265543 | 476.17 MB/s | 3985336 | 30015 | 8.1× |
| Sonic | 9519640 | 313.40 MB/s | 9129188 | 57804 | 5.3× |
| SonicFastest | 9578799 | 311.47 MB/s | 9129249 | 57804 | 5.3× |
| Goccy | 17744396 | 168.14 MB/s | 9750601 | 273615 | 2.9× |
| Easyjson | 18067987 | 165.12 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 19785252 | 92.71 MB/s | 20023835 | 408557 | 2.6× |
| JSONV2 | 24924523 | 119.70 MB/s | 9257063 | 86278 | 2.0× |
| Stdlib | 50925032 | 58.59 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1437081 | 503.52 MB/s | 907386 | 3618 | 8.6× |
| SonicFastest | 2189987 | 330.41 MB/s | 2367598 | 3683 | 5.7× |
| Sonic | 2211587 | 327.18 MB/s | 2367677 | 3683 | 5.6× |
| Easyjson | 5327275 | 135.83 MB/s | 2847906 | 3698 | 2.3× |
| Goccy | 5394369 | 134.14 MB/s | 2681738 | 80266 | 2.3× |
| LightningDecodeAny | 5489533 | 118.51 MB/s | 5752205 | 80172 | 2.3× |
| JSONV2 | 6006185 | 120.48 MB/s | 2704706 | 7318 | 2.1× |
| Stdlib | 12385663 | 58.42 MB/s | 2704547 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2086980 | 755.81 MB/s | 907386 | 3618 | 8.5× |
| Sonic | 2548788 | 618.86 MB/s | 3224694 | 3683 | 6.9× |
| SonicFastest | 2601003 | 606.44 MB/s | 3229169 | 3683 | 6.8× |
| LightningDecodeAny | 4506321 | 167.19 MB/s | 5752202 | 80172 | 3.9× |
| Easyjson | 6366877 | 247.74 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 6653902 | 237.06 MB/s | 3413114 | 80257 | 2.7× |
| JSONV2 | 7078439 | 222.84 MB/s | 2704552 | 7318 | 2.5× |
| Stdlib | 17678754 | 89.22 MB/s | 2704547 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 234978 | 638.88 MB/s | 81920 | 1 | 9.4× |
| Sonic | 389893 | 385.04 MB/s | 407077 | 16 | 5.7× |
| SonicFastest | 409878 | 366.27 MB/s | 407056 | 16 | 5.4× |
| LightningDecodeAny | 586559 | 255.93 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 994489 | 150.96 MB/s | 327146 | 10005 | 2.2× |
| JSONV2 | 1139755 | 131.72 MB/s | 357725 | 20 | 1.9× |
| Stdlib | 2217746 | 67.69 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 34996 | 803.44 MB/s | 30272 | 105 | 9.0× |
| Sonic | 71397 | 393.81 MB/s | 59559 | 83 | 4.4× |
| SonicFastest | 71433 | 393.61 MB/s | 59565 | 83 | 4.4× |
| Easyjson | 74779 | 376.00 MB/s | 32304 | 138 | 4.2× |
| Goccy | 83805 | 335.51 MB/s | 59258 | 188 | 3.8× |
| JSONV2 | 132257 | 212.59 MB/s | 36896 | 242 | 2.4× |
| LightningDecodeAny | 157278 | 178.77 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 314275 | 89.47 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1944 | 1197.68 MB/s | 32 | 1 | 12.2× |
| Goccy | 4560 | 510.52 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 5321 | 437.50 MB/s | 192 | 2 | 4.5× |
| Sonic | 6046 | 385.02 MB/s | 3708 | 4 | 3.9× |
| SonicFastest | 6055 | 384.49 MB/s | 3713 | 4 | 3.9× |
| JSONV2 | 7397 | 314.74 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 9881 | 170.53 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 23725 | 98.12 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 226 | 837.19 MB/s | 0 | 0 | 10.8× |
| Goccy | 414 | 456.11 MB/s | 304 | 2 | 5.9× |
| Easyjson | 525 | 360.18 MB/s | 0 | 0 | 4.7× |
| Sonic | 741 | 255.05 MB/s | 341 | 3 | 3.3× |
| SonicFastest | 741 | 254.98 MB/s | 341 | 3 | 3.3× |
| JSONV2 | 945 | 199.92 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1238 | 108.21 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2444 | 77.34 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1449 | 1512.58 MB/s | 0 | 0 | 11.8× |
| Easyjson | 3321 | 659.75 MB/s | 24 | 1 | 5.2× |
| Goccy | 3691 | 593.65 MB/s | 2864 | 4 | 4.6× |
| SonicFastest | 6590 | 332.46 MB/s | 3600 | 38 | 2.6× |
| Sonic | 6809 | 321.76 MB/s | 3599 | 38 | 2.5× |
| JSONV2 | 8004 | 273.74 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 9021 | 200.76 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 17153 | 127.73 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 714268 | 714.68 MB/s | 703779 | 1012 | 9.0× |
| Goccy | 1242200 | 410.95 MB/s | 1134535 | 5006 | 5.2× |
| Sonic | 1452035 | 351.56 MB/s | 1307322 | 2014 | 4.5× |
| SonicFastest | 1504419 | 339.32 MB/s | 1307342 | 2014 | 4.3× |
| Easyjson | 1636639 | 311.91 MB/s | 863781 | 3012 | 3.9× |
| JSONV2 | 3057326 | 166.97 MB/s | 1075953 | 12645 | 2.1× |
| LightningDecodeAny | 3448550 | 133.81 MB/s | 2785928 | 66022 | 1.9× |
| Stdlib | 6462082 | 79.00 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 611 | 32376.22 MB/s | 0 | 0 | 234.5× |
| SonicFastest | 6883 | 2874.91 MB/s | 21067 | 3 | 20.8× |
| Goccy | 25406 | 778.90 MB/s | 20492 | 2 | 5.6× |
| Sonic | 29430 | 672.41 MB/s | 20623 | 3 | 4.9× |
| JSONV2 | 35490 | 557.59 MB/s | 8 | 1 | 4.0× |
| LightningDecodeAny | 95979 | 206.17 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 109124 | 181.34 MB/s | 0 | 0 | 1.3× |
| Stdlib | 143346 | 138.05 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2261 | 8017.30 MB/s | 432 | 2 | 53.6× |
| Easyjson | 4486 | 4040.28 MB/s | 432 | 2 | 27.0× |
| Sonic | 10095 | 1795.26 MB/s | 20432 | 5 | 12.0× |
| SonicFastest | 10432 | 1737.30 MB/s | 20376 | 5 | 11.6× |
| LightningDecodeAny | 18720 | 955.24 MB/s | 29088 | 191 | 6.5× |
| Goccy | 26929 | 673.03 MB/s | 19460 | 2 | 4.5× |
| JSONV2 | 48720 | 372.00 MB/s | 16501 | 50 | 2.5× |
| Stdlib | 121237 | 149.49 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2405707 | 834.89 MB/s | 2962101 | 7417 | 8.6× |
| Goccy | 4561704 | 440.29 MB/s | 5410219 | 15830 | 4.6× |
| SonicFastest | 4796303 | 418.76 MB/s | 5152218 | 7085 | 4.3× |
| Sonic | 4844718 | 414.57 MB/s | 5152832 | 7085 | 4.3× |
| Easyjson | 5090202 | 394.58 MB/s | 2981485 | 7439 | 4.1× |
| LightningDecodeAny | 6918884 | 165.10 MB/s | 7386073 | 134751 | 3.0× |
| JSONV2 | 7121329 | 282.04 MB/s | 3173670 | 14562 | 2.9× |
| Stdlib | 20774010 | 96.68 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 834 | 658.32 MB/s | 480 | 1 | 6.9× |
| LightningDecodeAny | 1901 | 288.24 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2004 | 274.02 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2110 | 260.16 MB/s | 2262 | 8 | 2.7× |
| Sonic | 2159 | 254.33 MB/s | 2262 | 8 | 2.7× |
| JSONV2 | 2930 | 187.39 MB/s | 1664 | 7 | 2.0× |
| Goccy | 2963 | 185.28 MB/s | 2129 | 43 | 2.0× |
| Stdlib | 5782 | 94.94 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 548959 | 1150.38 MB/s | 413001 | 878 | 10.9× |
| Sonic | 1103328 | 572.37 MB/s | 1066287 | 814 | 5.4× |
| SonicFastest | 1133458 | 557.16 MB/s | 1066320 | 814 | 5.3× |
| Easyjson | 1311200 | 481.63 MB/s | 422504 | 936 | 4.6× |
| Goccy | 1312482 | 481.16 MB/s | 991816 | 1201 | 4.6× |
| JSONV2 | 2115370 | 298.54 MB/s | 571595 | 3144 | 2.8× |
| LightningDecodeAny | 2470426 | 189.00 MB/s | 2010198 | 30295 | 2.4× |
| Stdlib | 5972208 | 105.74 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 962420 | 584.37 MB/s | 767623 | 1254 | 5.8× |
| Sonic | 1393440 | 403.61 MB/s | 1349647 | 1185 | 4.0× |
| SonicFastest | 1395624 | 402.98 MB/s | 1349149 | 1184 | 4.0× |
| Goccy | 1508706 | 372.78 MB/s | 1042902 | 1029 | 3.7× |
| Easyjson | 2016041 | 278.97 MB/s | 775155 | 1254 | 2.8× |
| LightningDecodeAny | 2947025 | 190.84 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 2979095 | 188.78 MB/s | 927403 | 3482 | 1.9× |
| Stdlib | 5618382 | 100.10 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 670092 | 795.68 MB/s | 368224 | 2293 | 8.9× |
| Easyjson | 1230470 | 433.31 MB/s | 428362 | 3273 | 4.9× |
| Sonic | 1365899 | 390.35 MB/s | 980416 | 3082 | 4.4× |
| SonicFastest | 1368035 | 389.74 MB/s | 979545 | 3082 | 4.4× |
| Goccy | 1499623 | 355.54 MB/s | 1167053 | 5408 | 4.0× |
| JSONV2 | 2622673 | 203.30 MB/s | 745419 | 13288 | 2.3× |
| LightningDecodeAny | 3304961 | 161.33 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 5991968 | 88.98 MB/s | 798692 | 17133 | 1.0× |
