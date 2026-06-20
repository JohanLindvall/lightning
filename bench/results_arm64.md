# JSON Deserialization Benchmarks

- generated 2026-06-20T18:49:55Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 107262 | 1186.58 MB/s | 50208 | 4 | 10.3× |
| SonicFastest | 182959 | 695.65 MB/s | 199992 | 10 | 6.0× |
| Sonic | 186390 | 682.84 MB/s | 205148 | 10 | 5.9× |
| Goccy | 198467 | 641.29 MB/s | 225346 | 884 | 5.6× |
| Easyjson | 210825 | 603.70 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 420431 | 302.73 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 454530 | 208.24 MB/s | 466033 | 9715 | 2.4× |
| Stdlib | 1104353 | 115.25 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3794694 | 593.21 MB/s | 3122875 | 3081 | 6.9× |
| SonicFastest | 4506167 | 499.55 MB/s | 15233785 | 970 | 5.8× |
| Sonic | 4601097 | 489.24 MB/s | 15233763 | 970 | 5.7× |
| Goccy | 10343856 | 217.62 MB/s | 4119584 | 56532 | 2.5× |
| Easyjson | 11154728 | 201.80 MB/s | 3099808 | 2120 | 2.3× |
| LightningDecodeAny | 11289355 | 199.39 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16186471 | 139.07 MB/s | 3123216 | 3083 | 1.6× |
| Stdlib | 26050773 | 86.41 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 484214 | 558.44 MB/s | 348024 | 1627 | 6.9× |
| Sonic | 630853 | 428.63 MB/s | 473197 | 968 | 5.3× |
| SonicFastest | 641595 | 421.45 MB/s | 491333 | 968 | 5.2× |
| Goccy | 1413046 | 191.36 MB/s | 543738 | 8122 | 2.4× |
| Easyjson | 1415186 | 191.07 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1570873 | 172.14 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2092003 | 129.26 MB/s | 348159 | 1628 | 1.6× |
| Stdlib | 3350514 | 80.70 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1420208 | 1216.16 MB/s | 959939 | 5883 | 9.3× |
| Sonic | 2040163 | 846.60 MB/s | 2771704 | 4020 | 6.5× |
| SonicFastest | 2048153 | 843.30 MB/s | 2778978 | 4020 | 6.5× |
| Goccy | 2355300 | 733.33 MB/s | 2583744 | 14605 | 5.6× |
| JSONV2 | 4238763 | 407.48 MB/s | 1011638 | 7594 | 3.1× |
| Easyjson | 4259777 | 405.47 MB/s | 972032 | 5389 | 3.1× |
| LightningDecodeAny | 4720968 | 105.97 MB/s | 4976253 | 81467 | 2.8× |
| Stdlib | 13269594 | 130.16 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1235 | 1466.67 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2539 | 713.65 MB/s | 24 | 1 | 5.5× |
| Goccy | 2776 | 652.77 MB/s | 2608 | 4 | 5.1× |
| Sonic | 5933 | 305.43 MB/s | 3713 | 40 | 2.4× |
| SonicFastest | 5934 | 305.37 MB/s | 3743 | 40 | 2.4× |
| JSONV2 | 7841 | 231.08 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8246 | 219.62 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14033 | 129.12 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1249 | 1450.99 MB/s | 0 | 0 | 11.2× |
| Easyjson | 2553 | 709.81 MB/s | 24 | 1 | 5.5× |
| Goccy | 2801 | 646.86 MB/s | 2608 | 4 | 5.0× |
| Sonic | 6076 | 298.25 MB/s | 3870 | 40 | 2.3× |
| SonicFastest | 6076 | 298.22 MB/s | 3852 | 40 | 2.3× |
| JSONV2 | 7794 | 232.49 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8635 | 209.73 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14045 | 129.01 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1439 | 1259.28 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2790 | 649.37 MB/s | 144 | 10 | 5.0× |
| Goccy | 2884 | 628.26 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6134 | 295.38 MB/s | 3896 | 42 | 2.3× |
| Sonic | 6165 | 293.93 MB/s | 3925 | 42 | 2.3× |
| JSONV2 | 7880 | 229.95 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8611 | 210.32 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14018 | 129.26 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 747 | 661.07 MB/s | 160 | 1 | 7.4× |
| SonicFastest | 1228 | 402.29 MB/s | 969 | 6 | 4.5× |
| Sonic | 1229 | 402.09 MB/s | 983 | 6 | 4.5× |
| LightningDecodeAny | 1629 | 302.66 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2202 | 224.35 MB/s | 448 | 3 | 2.5× |
| Goccy | 2463 | 200.54 MB/s | 856 | 23 | 2.2× |
| JSONV2 | 3241 | 152.43 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5530 | 89.33 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 466 | 493.30 MB/s | 160 | 1 | 8.9× |
| Sonic | 874 | 263.22 MB/s | 648 | 6 | 4.7× |
| SonicFastest | 876 | 262.62 MB/s | 652 | 6 | 4.7× |
| LightningDecodeAny | 1366 | 167.69 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1382 | 166.41 MB/s | 448 | 3 | 3.0× |
| Goccy | 1615 | 142.43 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2388 | 96.32 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4149 | 55.44 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 79194 | 822.43 MB/s | 178609 | 112 | 7.0× |
| SonicFastest | 97452 | 668.35 MB/s | 154763 | 75 | 5.7× |
| Sonic | 98858 | 658.85 MB/s | 155793 | 75 | 5.6× |
| Goccy | 142606 | 456.73 MB/s | 229292 | 134 | 3.9× |
| LightningDecodeAny | 190888 | 279.37 MB/s | 183136 | 3257 | 2.9× |
| JSONV2 | 223032 | 292.03 MB/s | 206650 | 607 | 2.5× |
| Stdlib | 553021 | 117.77 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2710855 | 715.82 MB/s | 2846867 | 2238 | 8.7× |
| SonicFastest | 4619894 | 420.03 MB/s | 14606972 | 1407 | 5.1× |
| Sonic | 4629834 | 419.12 MB/s | 14608572 | 1407 | 5.1× |
| Goccy | 4784877 | 405.54 MB/s | 4065913 | 13510 | 4.9× |
| Easyjson | 7488611 | 259.12 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9652453 | 201.03 MB/s | 7013524 | 219937 | 2.4× |
| JSONV2 | 11173556 | 173.67 MB/s | 3237222 | 13947 | 2.1× |
| Stdlib | 23472510 | 82.67 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2256827 | 1474.56 MB/s | 4607420 | 4704 | 9.2× |
| Sonic | 2664296 | 1249.05 MB/s | 6514655 | 4248 | 7.8× |
| SonicFastest | 2669116 | 1246.79 MB/s | 6477563 | 4248 | 7.8× |
| LightningDecodeAny | 4339531 | 708.32 MB/s | 7005137 | 58601 | 4.8× |
| Goccy | 4573582 | 727.62 MB/s | 3948907 | 3816 | 4.6× |
| JSONV2 | 7484591 | 444.62 MB/s | 5364524 | 13243 | 2.8× |
| Stdlib | 20830723 | 159.76 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 253519 | 869.15 MB/s | 123584 | 225 | 8.0× |
| Sonic | 374315 | 588.66 MB/s | 292824 | 398 | 5.4× |
| SonicFastest | 375305 | 587.11 MB/s | 294942 | 398 | 5.4× |
| Goccy | 428544 | 514.17 MB/s | 364007 | 1067 | 4.8× |
| Easyjson | 548813 | 401.50 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 735447 | 299.61 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 877973 | 123.37 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2038419 | 108.10 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13204554 | 613.43 MB/s | 15059842 | 51649 | 6.7× |
| Sonic | 16876734 | 479.95 MB/s | 70873028 | 40014 | 5.3× |
| SonicFastest | 17009955 | 476.19 MB/s | 70887433 | 40014 | 5.2× |
| Goccy | 23752110 | 341.02 MB/s | 17132182 | 107149 | 3.7× |
| Easyjson | 30570103 | 264.97 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 40368132 | 128.89 MB/s | 34333353 | 912810 | 2.2× |
| JSONV2 | 44606482 | 181.59 MB/s | 15233745 | 78972 | 2.0× |
| Stdlib | 88965794 | 91.05 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6282601 | 474.88 MB/s | 3985339 | 30015 | 7.5× |
| Sonic | 8576951 | 347.85 MB/s | 26735703 | 56760 | 5.5× |
| SonicFastest | 8629531 | 345.73 MB/s | 26577565 | 56760 | 5.4× |
| Easyjson | 16314826 | 182.87 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16580870 | 179.93 MB/s | 10606233 | 273649 | 2.8× |
| LightningDecodeAny | 18474936 | 99.28 MB/s | 20023837 | 408557 | 2.5× |
| JSONV2 | 23974957 | 124.44 MB/s | 9257126 | 86278 | 2.0× |
| Stdlib | 46974363 | 63.51 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1370860 | 527.84 MB/s | 907388 | 3618 | 8.4× |
| Sonic | 1757692 | 411.67 MB/s | 3175815 | 7226 | 6.6× |
| SonicFastest | 1761740 | 410.73 MB/s | 3178776 | 7226 | 6.6× |
| Easyjson | 4178806 | 173.16 MB/s | 2847906 | 3698 | 2.8× |
| LightningDecodeAny | 4289081 | 151.68 MB/s | 5752202 | 80172 | 2.7× |
| Goccy | 4720135 | 153.30 MB/s | 2837887 | 80275 | 2.5× |
| JSONV2 | 5816423 | 124.41 MB/s | 2704637 | 7318 | 2.0× |
| Stdlib | 11567749 | 62.55 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1969638 | 800.83 MB/s | 907386 | 3618 | 8.0× |
| SonicFastest | 2222313 | 709.78 MB/s | 5793518 | 7226 | 7.1× |
| Sonic | 2222409 | 709.75 MB/s | 5787269 | 7226 | 7.1× |
| LightningDecodeAny | 3809836 | 197.75 MB/s | 5752200 | 80172 | 4.1× |
| Easyjson | 5584426 | 282.46 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5640670 | 279.64 MB/s | 3565438 | 80266 | 2.8× |
| JSONV2 | 6411577 | 246.02 MB/s | 2704594 | 7318 | 2.5× |
| Stdlib | 15713213 | 100.38 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224418 | 668.95 MB/s | 81920 | 1 | 8.2× |
| SonicFastest | 270518 | 554.95 MB/s | 248714 | 6 | 6.8× |
| Sonic | 271130 | 553.70 MB/s | 250916 | 6 | 6.8× |
| LightningDecodeAny | 476891 | 314.79 MB/s | 746004 | 10020 | 3.9× |
| Goccy | 857046 | 175.16 MB/s | 323567 | 10004 | 2.2× |
| JSONV2 | 1092487 | 137.41 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1843508 | 81.43 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 37967 | 740.56 MB/s | 29104 | 107 | 8.0× |
| Sonic | 63146 | 445.27 MB/s | 46467 | 103 | 4.8× |
| SonicFastest | 63370 | 443.69 MB/s | 46863 | 103 | 4.8× |
| Easyjson | 68211 | 412.21 MB/s | 32304 | 138 | 4.4× |
| Goccy | 70572 | 398.42 MB/s | 59193 | 188 | 4.3× |
| JSONV2 | 132784 | 211.75 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 152926 | 183.86 MB/s | 135136 | 2680 | 2.0× |
| Stdlib | 303195 | 92.74 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2039 | 1141.95 MB/s | 32 | 1 | 11.1× |
| Goccy | 4087 | 569.56 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4201 | 554.20 MB/s | 192 | 2 | 5.4× |
| Sonic | 5015 | 464.24 MB/s | 4223 | 6 | 4.5× |
| SonicFastest | 5032 | 462.61 MB/s | 4229 | 6 | 4.5× |
| JSONV2 | 8443 | 275.74 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10246 | 164.46 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22655 | 102.76 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 229 | 823.90 MB/s | 0 | 0 | 10.7× |
| Goccy | 380 | 497.65 MB/s | 304 | 2 | 6.5× |
| Easyjson | 491 | 385.21 MB/s | 0 | 0 | 5.0× |
| Sonic | 785 | 240.63 MB/s | 495 | 4 | 3.1× |
| SonicFastest | 788 | 239.86 MB/s | 503 | 4 | 3.1× |
| JSONV2 | 1030 | 183.48 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1247 | 107.49 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2462 | 76.77 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1549 | 1414.02 MB/s | 0 | 0 | 10.3× |
| Goccy | 3184 | 688.15 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3197 | 685.25 MB/s | 24 | 1 | 5.0× |
| Sonic | 6406 | 342.03 MB/s | 3997 | 40 | 2.5× |
| SonicFastest | 6410 | 341.82 MB/s | 3992 | 40 | 2.5× |
| JSONV2 | 8031 | 272.81 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8258 | 219.29 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 15960 | 137.28 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 923004 | 553.06 MB/s | 703779 | 1012 | 6.6× |
| Sonic | 1148966 | 444.29 MB/s | 881533 | 2006 | 5.3× |
| SonicFastest | 1150911 | 443.54 MB/s | 877548 | 2006 | 5.3× |
| Goccy | 1152347 | 442.99 MB/s | 1139604 | 5006 | 5.3× |
| Easyjson | 1534109 | 332.75 MB/s | 863778 | 3012 | 4.0× |
| JSONV2 | 3183650 | 160.34 MB/s | 1076012 | 12646 | 1.9× |
| LightningDecodeAny | 3542196 | 130.28 MB/s | 2785926 | 66022 | 1.7× |
| Stdlib | 6068947 | 84.11 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1337 | 14806.36 MB/s | 0 | 0 | 83.5× |
| Goccy | 20181 | 980.56 MB/s | 20491 | 2 | 5.5× |
| SonicFastest | 27277 | 725.50 MB/s | 22303 | 4 | 4.1× |
| Sonic | 27493 | 719.78 MB/s | 22865 | 4 | 4.1× |
| JSONV2 | 29603 | 668.48 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 79423 | 249.15 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86055 | 229.96 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111633 | 177.27 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2932 | 6182.46 MB/s | 864 | 4 | 35.1× |
| Easyjson | 3967 | 4568.54 MB/s | 432 | 2 | 25.9× |
| SonicFastest | 9849 | 1840.21 MB/s | 22828 | 6 | 10.4× |
| Sonic | 10008 | 1810.98 MB/s | 23248 | 6 | 10.3× |
| Goccy | 15670 | 1156.63 MB/s | 19459 | 2 | 6.6× |
| LightningDecodeAny | 16770 | 1066.31 MB/s | 29520 | 193 | 6.1× |
| JSONV2 | 44572 | 406.63 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102868 | 176.19 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2368009 | 848.18 MB/s | 2963783 | 7423 | 7.9× |
| Goccy | 4232699 | 474.52 MB/s | 5412528 | 15838 | 4.4× |
| Sonic | 4433049 | 453.07 MB/s | 10947421 | 13683 | 4.2× |
| SonicFastest | 4447031 | 451.65 MB/s | 10976599 | 13683 | 4.2× |
| Easyjson | 4924491 | 407.86 MB/s | 2981514 | 7439 | 3.8× |
| JSONV2 | 6899201 | 291.12 MB/s | 3173694 | 14563 | 2.7× |
| LightningDecodeAny | 7374059 | 154.91 MB/s | 7387753 | 134757 | 2.5× |
| Stdlib | 18625534 | 107.84 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 901 | 609.31 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1958 | 279.82 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2133 | 257.34 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2692 | 203.95 MB/s | 1966 | 26 | 2.1× |
| SonicFastest | 2698 | 203.45 MB/s | 1950 | 26 | 2.1× |
| Goccy | 3045 | 180.28 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3254 | 168.70 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5643 | 97.29 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 584433 | 1080.56 MB/s | 460873 | 1190 | 9.2× |
| Sonic | 1014986 | 622.19 MB/s | 1005694 | 1102 | 5.3× |
| SonicFastest | 1015430 | 621.92 MB/s | 1010732 | 1102 | 5.3× |
| Easyjson | 1147911 | 550.14 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1151928 | 548.22 MB/s | 986802 | 1201 | 4.7× |
| JSONV2 | 2141839 | 294.85 MB/s | 571613 | 3144 | 2.5× |
| LightningDecodeAny | 2442586 | 191.15 MB/s | 2058072 | 30607 | 2.2× |
| Stdlib | 5404409 | 116.85 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 962458 | 584.35 MB/s | 886433 | 2062 | 5.5× |
| Sonic | 1020623 | 551.04 MB/s | 946994 | 1476 | 5.2× |
| SonicFastest | 1023063 | 549.73 MB/s | 953499 | 1476 | 5.1× |
| Goccy | 1323868 | 424.82 MB/s | 1039383 | 1030 | 4.0× |
| Easyjson | 1738295 | 323.54 MB/s | 775154 | 1254 | 3.0× |
| JSONV2 | 2739600 | 205.29 MB/s | 927439 | 3482 | 1.9× |
| LightningDecodeAny | 2828785 | 198.82 MB/s | 2232993 | 31103 | 1.9× |
| Stdlib | 5256852 | 106.99 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 766608 | 695.50 MB/s | 413025 | 3155 | 7.1× |
| Easyjson | 1103101 | 483.34 MB/s | 428361 | 3273 | 5.0× |
| SonicFastest | 1141611 | 467.04 MB/s | 1036770 | 4351 | 4.8× |
| Sonic | 1147142 | 464.79 MB/s | 1054188 | 4351 | 4.8× |
| Goccy | 1277176 | 417.47 MB/s | 1167237 | 5409 | 4.3× |
| JSONV2 | 2522274 | 211.39 MB/s | 745447 | 13288 | 2.2× |
| LightningDecodeAny | 3404954 | 156.59 MB/s | 2709156 | 50805 | 1.6× |
| Stdlib | 5469773 | 97.48 MB/s | 798692 | 17133 | 1.0× |
