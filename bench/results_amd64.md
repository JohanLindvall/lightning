# JSON Deserialization Benchmarks

- generated 2026-06-22T19:49:47Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 107871 | 1179.88 MB/s | 49280 | 2 | 12.4× |
| Lightning | 108628 | 1171.66 MB/s | 49760 | 3 | 12.3× |
| SonicFastest | 201868 | 630.49 MB/s | 214433 | 15 | 6.6× |
| Sonic | 202585 | 628.25 MB/s | 214421 | 15 | 6.6× |
| Goccy | 253492 | 502.09 MB/s | 225265 | 884 | 5.3× |
| Easyjson | 258779 | 491.83 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 458895 | 277.35 MB/s | 195128 | 1805 | 2.9× |
| LightningDecodeAny | 478725 | 197.72 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1338372 | 95.10 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4314567 | 521.73 MB/s | 3122873 | 3081 | 8.0× |
| Lightning | 4331279 | 519.72 MB/s | 3122872 | 3081 | 8.0× |
| Sonic | 4868196 | 462.40 MB/s | 4873514 | 2584 | 7.1× |
| SonicFastest | 5473710 | 411.25 MB/s | 4869800 | 2584 | 6.3× |
| Goccy | 11987535 | 187.78 MB/s | 4201355 | 56535 | 2.9× |
| LightningDecodeAny | 12766909 | 176.32 MB/s | 7938300 | 281383 | 2.7× |
| Easyjson | 14067957 | 160.01 MB/s | 3099810 | 2120 | 2.5× |
| JSONV2 | 17303973 | 130.09 MB/s | 3123215 | 3083 | 2.0× |
| Stdlib | 34471319 | 65.30 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 590860 | 457.64 MB/s | 348025 | 1627 | 7.5× |
| LightningDestructive | 594872 | 454.56 MB/s | 348024 | 1627 | 7.5× |
| SonicFastest | 767935 | 352.12 MB/s | 641028 | 1147 | 5.8× |
| Sonic | 770259 | 351.05 MB/s | 640481 | 1147 | 5.8× |
| Goccy | 1697457 | 159.30 MB/s | 541316 | 8122 | 2.6× |
| LightningDecodeAny | 1743874 | 155.06 MB/s | 1011487 | 37901 | 2.5× |
| Easyjson | 1807717 | 149.58 MB/s | 330272 | 749 | 2.5× |
| JSONV2 | 2356990 | 114.72 MB/s | 348161 | 1628 | 1.9× |
| Stdlib | 4439950 | 60.90 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1324676 | 1303.87 MB/s | 959849 | 5881 | 13.5× |
| Lightning | 1386620 | 1245.62 MB/s | 959890 | 5882 | 12.9× |
| SonicFastest | 2203721 | 783.77 MB/s | 2694482 | 5547 | 8.1× |
| Sonic | 2204743 | 783.40 MB/s | 2694041 | 5547 | 8.1× |
| Goccy | 2643324 | 653.42 MB/s | 2581689 | 14604 | 6.8× |
| Easyjson | 4126311 | 418.58 MB/s | 972032 | 5389 | 4.3× |
| LightningDecodeAny | 4608525 | 108.56 MB/s | 4976203 | 81466 | 3.9× |
| JSONV2 | 4889140 | 353.27 MB/s | 1011615 | 7594 | 3.7× |
| Stdlib | 17856001 | 96.73 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1046 | 1732.57 MB/s | 0 | 0 | 15.5× |
| LightningDestructive | 1094 | 1656.13 MB/s | 0 | 0 | 14.8× |
| Easyjson | 2943 | 615.71 MB/s | 24 | 1 | 5.5× |
| Goccy | 3301 | 548.96 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6425 | 282.02 MB/s | 3345 | 38 | 2.5× |
| Sonic | 6711 | 270.01 MB/s | 3348 | 38 | 2.4× |
| JSONV2 | 8278 | 218.90 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9618 | 188.29 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16234 | 111.62 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1120 | 1617.73 MB/s | 0 | 0 | 14.5× |
| LightningDestructive | 1209 | 1498.60 MB/s | 0 | 0 | 13.5× |
| Easyjson | 3032 | 597.68 MB/s | 24 | 1 | 5.4× |
| Goccy | 3221 | 562.56 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6152 | 294.56 MB/s | 3346 | 38 | 2.6× |
| Sonic | 6318 | 286.80 MB/s | 3347 | 38 | 2.6× |
| JSONV2 | 8155 | 222.19 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9334 | 194.02 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16266 | 111.40 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1327 | 1365.41 MB/s | 144 | 10 | 12.2× |
| LightningDestructive | 1394 | 1300.13 MB/s | 144 | 10 | 11.6× |
| Easyjson | 3208 | 564.86 MB/s | 144 | 10 | 5.0× |
| Goccy | 3522 | 514.53 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 6317 | 286.86 MB/s | 3364 | 40 | 2.6× |
| Sonic | 6572 | 275.73 MB/s | 3365 | 40 | 2.5× |
| JSONV2 | 8881 | 204.03 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 9322 | 194.26 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16169 | 112.07 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 732 | 675.19 MB/s | 160 | 1 | 9.6× |
| LightningDestructive | 742 | 665.34 MB/s | 160 | 1 | 9.5× |
| Sonic | 1286 | 384.27 MB/s | 1075 | 8 | 5.5× |
| SonicFastest | 1288 | 383.47 MB/s | 1076 | 8 | 5.5× |
| LightningDecodeAny | 1801 | 273.66 MB/s | 1536 | 30 | 3.9× |
| Goccy | 2624 | 188.27 MB/s | 856 | 23 | 2.7× |
| Easyjson | 2718 | 181.78 MB/s | 448 | 3 | 2.6× |
| JSONV2 | 3500 | 141.16 MB/s | 528 | 7 | 2.0× |
| Stdlib | 7026 | 70.31 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 499 | 460.80 MB/s | 160 | 1 | 9.9× |
| Lightning | 500 | 460.37 MB/s | 160 | 1 | 9.9× |
| SonicFastest | 932 | 246.82 MB/s | 801 | 8 | 5.3× |
| Sonic | 935 | 246.09 MB/s | 800 | 8 | 5.3× |
| LightningDecodeAny | 1634 | 140.12 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1811 | 127.00 MB/s | 448 | 3 | 2.7× |
| Goccy | 1818 | 126.51 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2633 | 87.35 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4961 | 46.36 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 91065 | 715.23 MB/s | 172434 | 107 | 7.5× |
| LightningDestructive | 97263 | 669.65 MB/s | 166213 | 102 | 7.0× |
| SonicFastest | 149773 | 434.87 MB/s | 235811 | 65 | 4.6× |
| Sonic | 154193 | 422.41 MB/s | 235821 | 65 | 4.4× |
| Goccy | 181098 | 359.65 MB/s | 228080 | 134 | 3.8× |
| LightningDecodeAny | 206719 | 257.98 MB/s | 176960 | 3252 | 3.3× |
| JSONV2 | 265743 | 245.09 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 684400 | 95.17 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2573515 | 754.02 MB/s | 2846864 | 2238 | 11.0× |
| Lightning | 2650910 | 732.00 MB/s | 2846866 | 2238 | 10.7× |
| Goccy | 5067963 | 382.89 MB/s | 4063192 | 13509 | 5.6× |
| Sonic | 5441406 | 356.61 MB/s | 4882554 | 1736 | 5.2× |
| SonicFastest | 5464082 | 355.13 MB/s | 4882846 | 1736 | 5.2× |
| Easyjson | 8357935 | 232.17 MB/s | 3871264 | 15043 | 3.4× |
| LightningDecodeAny | 10163309 | 190.93 MB/s | 7013525 | 219937 | 2.8× |
| JSONV2 | 12639312 | 153.53 MB/s | 3237192 | 13947 | 2.2× |
| Stdlib | 28292043 | 68.59 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1090782 | 3050.87 MB/s | 351704 | 1286 | 22.4× |
| Lightning | 1741713 | 1910.67 MB/s | 2488907 | 2995 | 14.0× |
| SonicFastest | 2118321 | 1570.98 MB/s | 5896372 | 4263 | 11.5× |
| Sonic | 2137061 | 1557.20 MB/s | 5895712 | 4263 | 11.4× |
| LightningDecodeAny | 3974826 | 773.31 MB/s | 4886622 | 56892 | 6.1× |
| Goccy | 6194768 | 537.20 MB/s | 3948914 | 3817 | 3.9× |
| JSONV2 | 8173486 | 407.15 MB/s | 5364507 | 13243 | 3.0× |
| Stdlib | 24422353 | 136.26 MB/s | 5565619 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 227146 | 970.06 MB/s | 136896 | 228 | 10.8× |
| LightningDestructive | 236536 | 931.55 MB/s | 136896 | 228 | 10.4× |
| SonicFastest | 486642 | 452.79 MB/s | 350317 | 262 | 5.0× |
| Sonic | 489506 | 450.14 MB/s | 350511 | 262 | 5.0× |
| Goccy | 494221 | 445.84 MB/s | 364084 | 1066 | 5.0× |
| Easyjson | 676899 | 325.52 MB/s | 130512 | 245 | 3.6× |
| JSONV2 | 722174 | 305.11 MB/s | 129747 | 470 | 3.4× |
| LightningDecodeAny | 1031682 | 104.99 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2453298 | 89.82 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14809918 | 546.93 MB/s | 15059835 | 51649 | 7.6× |
| Lightning | 15919020 | 508.83 MB/s | 15059844 | 51649 | 7.0× |
| SonicFastest | 18018453 | 449.54 MB/s | 19859102 | 41640 | 6.2× |
| Sonic | 18124399 | 446.91 MB/s | 19857252 | 41640 | 6.2× |
| Goccy | 25128217 | 322.35 MB/s | 19016434 | 107155 | 4.5× |
| Easyjson | 36769111 | 220.29 MB/s | 15059620 | 41643 | 3.0× |
| LightningDecodeAny | 43597537 | 119.34 MB/s | 34333350 | 912810 | 2.6× |
| JSONV2 | 50031355 | 161.90 MB/s | 15233721 | 78972 | 2.2× |
| Stdlib | 112109093 | 72.25 MB/s | 15665073 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6324603 | 471.72 MB/s | 3985339 | 30015 | 9.1× |
| LightningDestructive | 6340665 | 470.53 MB/s | 3985336 | 30015 | 9.0× |
| Sonic | 9258497 | 322.24 MB/s | 9131774 | 57804 | 6.2× |
| SonicFastest | 9333510 | 319.65 MB/s | 9131304 | 57804 | 6.1× |
| Goccy | 18551188 | 160.82 MB/s | 9845246 | 273620 | 3.1× |
| Easyjson | 18916891 | 157.71 MB/s | 9479441 | 30115 | 3.0× |
| LightningDecodeAny | 20952806 | 87.54 MB/s | 20023834 | 408557 | 2.7× |
| JSONV2 | 29184986 | 102.23 MB/s | 9257069 | 86278 | 2.0× |
| Stdlib | 57363284 | 52.01 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1472919 | 491.27 MB/s | 907393 | 3618 | 9.6× |
| Lightning | 1546452 | 467.91 MB/s | 907387 | 3618 | 9.2× |
| SonicFastest | 2092179 | 345.86 MB/s | 2372464 | 3683 | 6.8× |
| Sonic | 2092524 | 345.80 MB/s | 2373054 | 3683 | 6.8× |
| Easyjson | 5321133 | 135.99 MB/s | 2847907 | 3698 | 2.7× |
| Goccy | 5538248 | 130.65 MB/s | 2709251 | 80268 | 2.6× |
| LightningDecodeAny | 5681556 | 114.51 MB/s | 5752204 | 80172 | 2.5× |
| JSONV2 | 6303173 | 114.80 MB/s | 2704707 | 7318 | 2.2× |
| Stdlib | 14163533 | 51.09 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2034010 | 775.49 MB/s | 907386 | 3618 | 9.7× |
| LightningDestructive | 2036545 | 774.52 MB/s | 907392 | 3618 | 9.7× |
| SonicFastest | 2443970 | 645.41 MB/s | 3226202 | 3683 | 8.1× |
| Sonic | 2488104 | 633.96 MB/s | 3228945 | 3683 | 7.9× |
| LightningDecodeAny | 4799003 | 156.99 MB/s | 5752199 | 80172 | 4.1× |
| Easyjson | 6278622 | 251.23 MB/s | 2847905 | 3698 | 3.1× |
| Goccy | 6664789 | 236.67 MB/s | 3486209 | 80262 | 3.0× |
| JSONV2 | 7186620 | 219.48 MB/s | 2704551 | 7318 | 2.7× |
| Stdlib | 19725817 | 79.96 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 250503 | 599.29 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 261265 | 574.60 MB/s | 81920 | 1 | 8.3× |
| Sonic | 429429 | 349.59 MB/s | 408077 | 16 | 5.0× |
| SonicFastest | 430375 | 348.82 MB/s | 408449 | 16 | 5.0× |
| LightningDecodeAny | 599621 | 250.36 MB/s | 746006 | 10020 | 3.6× |
| Goccy | 1063755 | 141.13 MB/s | 330951 | 10005 | 2.0× |
| JSONV2 | 1133100 | 132.49 MB/s | 357724 | 20 | 1.9× |
| Stdlib | 2166735 | 69.29 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 34419 | 816.91 MB/s | 30272 | 105 | 10.0× |
| LightningDestructive | 35759 | 786.30 MB/s | 30144 | 103 | 9.6× |
| SonicFastest | 59933 | 469.14 MB/s | 59522 | 83 | 5.8× |
| Sonic | 59961 | 468.92 MB/s | 59518 | 83 | 5.8× |
| Goccy | 81976 | 342.99 MB/s | 59277 | 188 | 4.2× |
| Easyjson | 82683 | 340.06 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 134079 | 209.70 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 168929 | 166.44 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 344967 | 81.51 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1971 | 1181.15 MB/s | 32 | 1 | 13.3× |
| LightningDestructive | 2105 | 1106.06 MB/s | 32 | 1 | 12.4× |
| SonicFastest | 4722 | 493.03 MB/s | 3707 | 4 | 5.5× |
| Sonic | 4769 | 488.14 MB/s | 3709 | 4 | 5.5× |
| Goccy | 4923 | 472.90 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 5101 | 456.39 MB/s | 192 | 2 | 5.1× |
| JSONV2 | 8424 | 276.36 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10652 | 158.18 MB/s | 9960 | 195 | 2.5× |
| Stdlib | 26140 | 89.06 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 215 | 880.57 MB/s | 0 | 0 | 13.1× |
| LightningDestructive | 225 | 841.20 MB/s | 0 | 0 | 12.5× |
| Goccy | 433 | 436.83 MB/s | 304 | 2 | 6.5× |
| Easyjson | 590 | 320.19 MB/s | 0 | 0 | 4.8× |
| SonicFastest | 646 | 292.63 MB/s | 341 | 3 | 4.3× |
| Sonic | 651 | 290.40 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1007 | 187.66 MB/s | 112 | 1 | 2.8× |
| LightningDecodeAny | 1363 | 98.33 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2806 | 67.35 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1491 | 1469.46 MB/s | 0 | 0 | 12.8× |
| LightningDestructive | 1545 | 1418.10 MB/s | 0 | 0 | 12.4× |
| Goccy | 3599 | 608.83 MB/s | 2864 | 4 | 5.3× |
| Easyjson | 3845 | 569.81 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6433 | 340.57 MB/s | 3597 | 38 | 3.0× |
| Sonic | 6635 | 330.20 MB/s | 3597 | 38 | 2.9× |
| JSONV2 | 8318 | 263.42 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 9207 | 196.71 MB/s | 7536 | 158 | 2.1× |
| Stdlib | 19129 | 114.54 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 703306 | 725.82 MB/s | 703779 | 1012 | 10.2× |
| Lightning | 740264 | 689.59 MB/s | 703778 | 1012 | 9.7× |
| SonicFastest | 1246923 | 409.39 MB/s | 1308037 | 2014 | 5.8× |
| Sonic | 1249209 | 408.64 MB/s | 1308028 | 2014 | 5.8× |
| Goccy | 1334330 | 382.57 MB/s | 1137784 | 5006 | 5.4× |
| Easyjson | 1710689 | 298.40 MB/s | 863778 | 3012 | 4.2× |
| JSONV2 | 3521556 | 144.96 MB/s | 1075957 | 12645 | 2.0× |
| LightningDecodeAny | 3686644 | 125.17 MB/s | 2785927 | 66022 | 2.0× |
| Stdlib | 7199487 | 70.90 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 643 | 30767.08 MB/s | 0 | 0 | 253.1× |
| LightningDestructive | 942 | 20998.79 MB/s | 0 | 0 | 172.7× |
| SonicFastest | 6515 | 3037.54 MB/s | 21119 | 3 | 25.0× |
| Goccy | 29370 | 673.79 MB/s | 20492 | 2 | 5.5× |
| Sonic | 32507 | 608.76 MB/s | 20631 | 3 | 5.0× |
| JSONV2 | 33260 | 594.97 MB/s | 8 | 1 | 4.9× |
| Easyjson | 101945 | 194.11 MB/s | 0 | 0 | 1.6× |
| LightningDecodeAny | 103872 | 190.50 MB/s | 117104 | 2019 | 1.6× |
| Stdlib | 162781 | 121.57 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2210 | 8200.76 MB/s | 432 | 2 | 58.0× |
| LightningDestructive | 2549 | 7109.49 MB/s | 0 | 0 | 50.3× |
| Easyjson | 5012 | 3615.97 MB/s | 432 | 2 | 25.6× |
| Sonic | 9655 | 1877.22 MB/s | 20394 | 5 | 13.3× |
| SonicFastest | 10224 | 1772.72 MB/s | 20382 | 5 | 12.5× |
| LightningDecodeAny | 19247 | 929.08 MB/s | 29088 | 191 | 6.7× |
| Goccy | 23251 | 779.51 MB/s | 19460 | 2 | 5.5× |
| JSONV2 | 51705 | 350.53 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 128104 | 141.48 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2276399 | 882.31 MB/s | 2960389 | 7411 | 9.8× |
| Lightning | 2408764 | 833.83 MB/s | 2962103 | 7417 | 9.2× |
| SonicFastest | 4222816 | 475.63 MB/s | 5155451 | 7085 | 5.3× |
| Sonic | 4341127 | 462.67 MB/s | 5154867 | 7085 | 5.1× |
| Goccy | 4876908 | 411.84 MB/s | 5411622 | 15833 | 4.6× |
| Easyjson | 5786720 | 347.09 MB/s | 2981492 | 7439 | 3.8× |
| LightningDecodeAny | 7271817 | 157.08 MB/s | 7386076 | 134751 | 3.1× |
| JSONV2 | 7804955 | 257.34 MB/s | 3173678 | 14563 | 2.9× |
| Stdlib | 22259815 | 90.23 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 855 | 642.14 MB/s | 480 | 1 | 8.0× |
| LightningDestructive | 856 | 641.27 MB/s | 480 | 1 | 8.0× |
| LightningDecodeAny | 2161 | 253.53 MB/s | 2261 | 50 | 3.2× |
| SonicFastest | 2378 | 230.90 MB/s | 2260 | 8 | 2.9× |
| Easyjson | 2387 | 229.98 MB/s | 1616 | 5 | 2.9× |
| Sonic | 2411 | 227.73 MB/s | 2260 | 8 | 2.8× |
| Goccy | 3356 | 163.57 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3634 | 151.08 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6842 | 80.24 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 503228 | 1254.93 MB/s | 364472 | 566 | 12.8× |
| Lightning | 582020 | 1085.04 MB/s | 413001 | 878 | 11.0× |
| Sonic | 1036129 | 609.49 MB/s | 1071956 | 814 | 6.2× |
| SonicFastest | 1043407 | 605.24 MB/s | 1071998 | 814 | 6.2× |
| Easyjson | 1309806 | 482.14 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1414193 | 446.55 MB/s | 987463 | 1200 | 4.5× |
| JSONV2 | 2313320 | 272.99 MB/s | 571591 | 3144 | 2.8× |
| LightningDecodeAny | 2554810 | 182.76 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6419484 | 98.37 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 818185 | 687.38 MB/s | 544249 | 448 | 7.4× |
| Lightning | 1038821 | 541.39 MB/s | 767622 | 1254 | 5.8× |
| Sonic | 1290535 | 435.79 MB/s | 1348753 | 1185 | 4.7× |
| SonicFastest | 1308897 | 429.68 MB/s | 1349307 | 1185 | 4.6× |
| Goccy | 1644597 | 341.97 MB/s | 1039903 | 1028 | 3.7× |
| Easyjson | 2152681 | 261.26 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 3143147 | 178.93 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 3204643 | 175.50 MB/s | 927407 | 3482 | 1.9× |
| Stdlib | 6072610 | 92.61 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 624299 | 854.04 MB/s | 333416 | 2084 | 10.3× |
| Lightning | 705851 | 755.37 MB/s | 368224 | 2293 | 9.1× |
| Sonic | 1113205 | 478.96 MB/s | 980851 | 3082 | 5.8× |
| SonicFastest | 1113231 | 478.95 MB/s | 980970 | 3082 | 5.8× |
| Easyjson | 1327003 | 401.79 MB/s | 428362 | 3273 | 4.8× |
| Goccy | 1426198 | 373.85 MB/s | 1167074 | 5408 | 4.5× |
| JSONV2 | 2840602 | 187.70 MB/s | 745421 | 13288 | 2.3× |
| LightningDecodeAny | 3553237 | 150.05 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6408744 | 83.20 MB/s | 798692 | 17133 | 1.0× |
