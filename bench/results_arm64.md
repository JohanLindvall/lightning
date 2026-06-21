# JSON Deserialization Benchmarks

- generated 2026-06-21T07:31:16Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105573 | 1205.56 MB/s | 49760 | 3 | 10.5× |
| SonicFastest | 182833 | 696.13 MB/s | 192248 | 10 | 6.1× |
| Sonic | 183123 | 695.03 MB/s | 193176 | 10 | 6.0× |
| Goccy | 196015 | 649.31 MB/s | 224523 | 884 | 5.6× |
| Easyjson | 212316 | 599.46 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 418373 | 304.21 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 470443 | 201.20 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1107073 | 114.97 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3821295 | 589.08 MB/s | 3122877 | 3081 | 6.8× |
| SonicFastest | 4629245 | 486.27 MB/s | 15233809 | 970 | 5.6× |
| Sonic | 4699221 | 479.03 MB/s | 15240506 | 970 | 5.5× |
| Goccy | 10764348 | 209.12 MB/s | 4132974 | 56533 | 2.4× |
| Easyjson | 11175268 | 201.43 MB/s | 3099808 | 2120 | 2.3× |
| LightningDecodeAny | 11540121 | 195.06 MB/s | 7938299 | 281383 | 2.3× |
| JSONV2 | 16164092 | 139.26 MB/s | 3123207 | 3083 | 1.6× |
| Stdlib | 26063875 | 86.37 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 488193 | 553.89 MB/s | 348024 | 1627 | 6.9× |
| Sonic | 650661 | 415.58 MB/s | 484834 | 968 | 5.1× |
| SonicFastest | 660911 | 409.14 MB/s | 502927 | 968 | 5.1× |
| Easyjson | 1423564 | 189.95 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1453741 | 186.00 MB/s | 543553 | 8122 | 2.3× |
| LightningDecodeAny | 1588656 | 170.21 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2116793 | 127.74 MB/s | 348155 | 1628 | 1.6× |
| Stdlib | 3348884 | 80.74 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1411687 | 1223.50 MB/s | 959890 | 5882 | 9.4× |
| Sonic | 2065033 | 836.41 MB/s | 2705901 | 4020 | 6.4× |
| SonicFastest | 2065167 | 836.35 MB/s | 2683826 | 4020 | 6.4× |
| Goccy | 2508955 | 688.42 MB/s | 2584323 | 14605 | 5.3× |
| Easyjson | 4254539 | 405.97 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4267143 | 404.77 MB/s | 1011633 | 7594 | 3.1× |
| LightningDecodeAny | 4790468 | 104.44 MB/s | 4976205 | 81466 | 2.8× |
| Stdlib | 13271587 | 130.14 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1234 | 1467.96 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2524 | 717.92 MB/s | 24 | 1 | 5.6× |
| Goccy | 2846 | 636.78 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 5980 | 303.00 MB/s | 3792 | 40 | 2.3× |
| Sonic | 5991 | 302.43 MB/s | 3804 | 40 | 2.3× |
| JSONV2 | 7715 | 234.88 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8398 | 215.64 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14038 | 129.08 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1241 | 1460.06 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2540 | 713.44 MB/s | 24 | 1 | 5.5× |
| Goccy | 2865 | 632.51 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6012 | 301.37 MB/s | 3762 | 40 | 2.3× |
| SonicFastest | 6035 | 300.24 MB/s | 3780 | 40 | 2.3× |
| JSONV2 | 7792 | 232.54 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8379 | 216.14 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14046 | 129.01 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1443 | 1255.62 MB/s | 144 | 10 | 9.8× |
| Easyjson | 2748 | 659.47 MB/s | 144 | 10 | 5.1× |
| Goccy | 2936 | 617.10 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6322 | 286.61 MB/s | 3873 | 42 | 2.2× |
| SonicFastest | 6342 | 285.73 MB/s | 3894 | 42 | 2.2× |
| JSONV2 | 7921 | 228.77 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8698 | 208.21 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 14072 | 128.77 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 749 | 659.38 MB/s | 160 | 1 | 7.4× |
| Sonic | 1245 | 396.88 MB/s | 976 | 6 | 4.5× |
| SonicFastest | 1250 | 395.34 MB/s | 975 | 6 | 4.4× |
| LightningDecodeAny | 1664 | 296.34 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2219 | 222.60 MB/s | 448 | 3 | 2.5× |
| Goccy | 2439 | 202.54 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3221 | 153.37 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5555 | 88.93 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 464 | 495.55 MB/s | 160 | 1 | 9.0× |
| Sonic | 887 | 259.34 MB/s | 649 | 6 | 4.7× |
| SonicFastest | 887 | 259.26 MB/s | 650 | 6 | 4.7× |
| Easyjson | 1391 | 165.31 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1397 | 163.87 MB/s | 1536 | 30 | 3.0× |
| Goccy | 1602 | 143.56 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2374 | 96.89 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4174 | 55.10 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 77258 | 843.05 MB/s | 172432 | 107 | 7.1× |
| Sonic | 99720 | 653.15 MB/s | 158385 | 75 | 5.5× |
| SonicFastest | 99759 | 652.90 MB/s | 158639 | 75 | 5.5× |
| Goccy | 143724 | 453.17 MB/s | 229167 | 134 | 3.8× |
| LightningDecodeAny | 190288 | 280.25 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 224053 | 290.70 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 546738 | 119.13 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2721876 | 712.92 MB/s | 2846866 | 2238 | 8.6× |
| SonicFastest | 4736488 | 409.69 MB/s | 14608661 | 1407 | 5.0× |
| Goccy | 4839444 | 400.97 MB/s | 4064673 | 13510 | 4.8× |
| Sonic | 5015152 | 386.92 MB/s | 14608697 | 1407 | 4.7× |
| Easyjson | 7554621 | 256.86 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 10091786 | 192.28 MB/s | 7013525 | 219937 | 2.3× |
| JSONV2 | 11246000 | 172.55 MB/s | 3237221 | 13947 | 2.1× |
| Stdlib | 23450290 | 82.75 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1816197 | 1832.31 MB/s | 2488905 | 2995 | 11.5× |
| Sonic | 2690967 | 1236.67 MB/s | 6477901 | 4248 | 7.8× |
| SonicFastest | 2735908 | 1216.35 MB/s | 6455768 | 4248 | 7.6× |
| LightningDecodeAny | 3864589 | 795.37 MB/s | 4886622 | 56892 | 5.4× |
| Goccy | 4597560 | 723.83 MB/s | 3948907 | 3816 | 4.5× |
| JSONV2 | 7435950 | 447.53 MB/s | 5364521 | 13243 | 2.8× |
| Stdlib | 20867986 | 159.47 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 230078 | 957.70 MB/s | 136896 | 228 | 8.9× |
| Sonic | 375814 | 586.32 MB/s | 299916 | 398 | 5.4× |
| SonicFastest | 377069 | 584.37 MB/s | 301354 | 398 | 5.4× |
| Goccy | 429585 | 512.93 MB/s | 364647 | 1067 | 4.7× |
| Easyjson | 548923 | 401.42 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 737915 | 298.61 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 897005 | 120.75 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2039442 | 108.04 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13233639 | 612.08 MB/s | 15059848 | 51649 | 6.7× |
| Sonic | 16683980 | 485.50 MB/s | 70887414 | 40014 | 5.3× |
| SonicFastest | 16747587 | 483.65 MB/s | 70887312 | 40014 | 5.3× |
| Goccy | 23670017 | 342.21 MB/s | 17005889 | 107148 | 3.8× |
| Easyjson | 30816203 | 262.85 MB/s | 15059620 | 41643 | 2.9× |
| LightningDecodeAny | 41461549 | 125.49 MB/s | 34333352 | 912810 | 2.2× |
| JSONV2 | 43857172 | 184.69 MB/s | 15233736 | 78972 | 2.0× |
| Stdlib | 89180929 | 90.83 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6260002 | 476.59 MB/s | 3985339 | 30015 | 7.5× |
| Sonic | 8791805 | 339.35 MB/s | 26457995 | 56760 | 5.4× |
| SonicFastest | 8801415 | 338.98 MB/s | 26597579 | 56760 | 5.4× |
| Easyjson | 16637688 | 179.32 MB/s | 9479440 | 30115 | 2.8× |
| Goccy | 16820271 | 177.37 MB/s | 10696798 | 273653 | 2.8× |
| LightningDecodeAny | 19220991 | 95.43 MB/s | 20023839 | 408557 | 2.5× |
| JSONV2 | 24386666 | 122.34 MB/s | 9257138 | 86278 | 1.9× |
| Stdlib | 47187361 | 63.23 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1382216 | 523.50 MB/s | 907388 | 3618 | 8.5× |
| Sonic | 1775776 | 407.48 MB/s | 3174137 | 7226 | 6.6× |
| SonicFastest | 1780517 | 406.40 MB/s | 3175326 | 7226 | 6.6× |
| Easyjson | 4268356 | 169.53 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4371728 | 148.81 MB/s | 5752202 | 80172 | 2.7× |
| Goccy | 4759815 | 152.02 MB/s | 2831159 | 80275 | 2.5× |
| JSONV2 | 5844388 | 123.81 MB/s | 2704672 | 7318 | 2.0× |
| Stdlib | 11705500 | 61.82 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1977882 | 797.50 MB/s | 907386 | 3618 | 8.0× |
| SonicFastest | 2244830 | 702.66 MB/s | 5783140 | 7226 | 7.0× |
| Sonic | 2245902 | 702.33 MB/s | 5788681 | 7226 | 7.0× |
| LightningDecodeAny | 3878912 | 194.23 MB/s | 5752203 | 80172 | 4.1× |
| Easyjson | 5654206 | 278.97 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5684037 | 277.51 MB/s | 3553737 | 80266 | 2.8× |
| JSONV2 | 6645631 | 237.35 MB/s | 2704592 | 7318 | 2.4× |
| Stdlib | 15762672 | 100.07 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225439 | 665.92 MB/s | 81920 | 1 | 8.2× |
| Sonic | 278472 | 539.10 MB/s | 265958 | 6 | 6.6× |
| SonicFastest | 279041 | 538.00 MB/s | 268069 | 6 | 6.6× |
| LightningDecodeAny | 481182 | 311.98 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 866466 | 173.26 MB/s | 325610 | 10005 | 2.1× |
| JSONV2 | 1096526 | 136.91 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1843857 | 81.42 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 36424 | 771.94 MB/s | 30272 | 105 | 8.4× |
| SonicFastest | 63429 | 443.28 MB/s | 46867 | 103 | 4.8× |
| Sonic | 63497 | 442.81 MB/s | 47125 | 103 | 4.8× |
| Easyjson | 68372 | 411.23 MB/s | 32304 | 138 | 4.5× |
| Goccy | 72156 | 389.67 MB/s | 59228 | 188 | 4.2× |
| JSONV2 | 134468 | 209.10 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 154912 | 181.50 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 304413 | 92.36 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1954 | 1191.34 MB/s | 32 | 1 | 11.6× |
| Goccy | 4146 | 561.53 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4208 | 553.22 MB/s | 192 | 2 | 5.4× |
| Sonic | 5155 | 451.58 MB/s | 4513 | 6 | 4.4× |
| SonicFastest | 5171 | 450.21 MB/s | 4534 | 6 | 4.4× |
| JSONV2 | 8420 | 276.47 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10955 | 153.81 MB/s | 9960 | 195 | 2.1× |
| Stdlib | 22725 | 102.44 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 230 | 822.61 MB/s | 0 | 0 | 10.8× |
| Goccy | 391 | 483.40 MB/s | 304 | 2 | 6.3× |
| Easyjson | 490 | 385.40 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 807 | 234.24 MB/s | 504 | 4 | 3.1× |
| Sonic | 808 | 233.83 MB/s | 498 | 4 | 3.1× |
| JSONV2 | 1035 | 182.62 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1274 | 105.15 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2476 | 76.32 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1538 | 1424.68 MB/s | 0 | 0 | 10.4× |
| Goccy | 3181 | 688.83 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3224 | 679.60 MB/s | 24 | 1 | 5.0× |
| Sonic | 6385 | 343.17 MB/s | 3989 | 40 | 2.5× |
| SonicFastest | 6402 | 342.22 MB/s | 3984 | 40 | 2.5× |
| JSONV2 | 8025 | 273.02 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8386 | 215.97 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 15967 | 137.22 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 745413 | 684.82 MB/s | 703780 | 1012 | 8.1× |
| Goccy | 1157358 | 441.07 MB/s | 1137052 | 5006 | 5.2× |
| SonicFastest | 1159463 | 440.27 MB/s | 891398 | 2006 | 5.2× |
| Sonic | 1166927 | 437.45 MB/s | 908396 | 2006 | 5.2× |
| Easyjson | 1532776 | 333.04 MB/s | 863778 | 3012 | 4.0× |
| JSONV2 | 3203256 | 159.36 MB/s | 1076010 | 12646 | 1.9× |
| LightningDecodeAny | 3598916 | 128.22 MB/s | 2785926 | 66022 | 1.7× |
| Stdlib | 6064908 | 84.17 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1337 | 14799.40 MB/s | 0 | 0 | 83.5× |
| Goccy | 20043 | 987.34 MB/s | 20491 | 2 | 5.6× |
| Sonic | 27237 | 726.55 MB/s | 22369 | 4 | 4.1× |
| SonicFastest | 27618 | 716.52 MB/s | 22968 | 4 | 4.0× |
| JSONV2 | 29744 | 665.31 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 78569 | 251.86 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86091 | 229.86 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111648 | 177.25 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2849 | 6361.25 MB/s | 432 | 2 | 36.2× |
| Easyjson | 3967 | 4568.91 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 10021 | 1808.52 MB/s | 23072 | 6 | 10.3× |
| Sonic | 10033 | 1806.49 MB/s | 23217 | 6 | 10.3× |
| Goccy | 15916 | 1138.72 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 17216 | 1038.66 MB/s | 29088 | 191 | 6.0× |
| JSONV2 | 44861 | 404.00 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103053 | 175.87 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2368616 | 847.96 MB/s | 2962101 | 7417 | 7.9× |
| Goccy | 4315484 | 465.42 MB/s | 5411351 | 15830 | 4.3× |
| Sonic | 4491205 | 447.21 MB/s | 10904729 | 13683 | 4.2× |
| SonicFastest | 4491448 | 447.18 MB/s | 10925769 | 13683 | 4.2× |
| Easyjson | 4934840 | 407.00 MB/s | 2981485 | 7439 | 3.8× |
| JSONV2 | 6936014 | 289.57 MB/s | 3173686 | 14563 | 2.7× |
| LightningDecodeAny | 7537689 | 151.54 MB/s | 7386073 | 134751 | 2.5× |
| Stdlib | 18651126 | 107.69 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 893 | 614.82 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 2024 | 270.69 MB/s | 2261 | 50 | 2.8× |
| Easyjson | 2167 | 253.31 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2712 | 202.40 MB/s | 1942 | 26 | 2.1× |
| SonicFastest | 2726 | 201.39 MB/s | 1948 | 26 | 2.1× |
| Goccy | 3080 | 178.24 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3305 | 166.11 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5665 | 96.90 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 558636 | 1130.46 MB/s | 413001 | 878 | 9.7× |
| Sonic | 1024479 | 616.42 MB/s | 1011362 | 1102 | 5.3× |
| SonicFastest | 1031709 | 612.10 MB/s | 1023533 | 1102 | 5.2× |
| Easyjson | 1155676 | 546.45 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1162498 | 543.24 MB/s | 987913 | 1201 | 4.7× |
| JSONV2 | 2171271 | 290.85 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2434167 | 191.81 MB/s | 2010201 | 30295 | 2.2× |
| Stdlib | 5415633 | 116.61 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 913723 | 615.51 MB/s | 767622 | 1254 | 5.8× |
| Sonic | 1013761 | 554.77 MB/s | 918033 | 1476 | 5.2× |
| SonicFastest | 1031485 | 545.24 MB/s | 951822 | 1476 | 5.1× |
| Goccy | 1347585 | 417.35 MB/s | 1040034 | 1030 | 3.9× |
| Easyjson | 1749329 | 321.50 MB/s | 775154 | 1254 | 3.0× |
| JSONV2 | 2775039 | 202.67 MB/s | 927446 | 3482 | 1.9× |
| LightningDecodeAny | 2826003 | 199.01 MB/s | 2114152 | 30295 | 1.9× |
| Stdlib | 5280413 | 106.51 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 682187 | 781.57 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1127919 | 472.71 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1149746 | 463.74 MB/s | 1032361 | 4351 | 4.8× |
| SonicFastest | 1154344 | 461.89 MB/s | 1041387 | 4351 | 4.8× |
| Goccy | 1316377 | 405.03 MB/s | 1167241 | 5409 | 4.2× |
| JSONV2 | 2522831 | 211.34 MB/s | 745446 | 13288 | 2.2× |
| LightningDecodeAny | 3457050 | 154.23 MB/s | 2674619 | 50596 | 1.6× |
| Stdlib | 5503586 | 96.88 MB/s | 798692 | 17133 | 1.0× |
