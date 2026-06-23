# JSON Deserialization Benchmarks

- generated 2026-06-23T19:53:44Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104406 | 1219.04 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 104529 | 1217.60 MB/s | 49280 | 2 | 10.6× |
| Sonic | 181293 | 702.04 MB/s | 192783 | 10 | 6.1× |
| SonicFastest | 182312 | 698.12 MB/s | 194511 | 10 | 6.1× |
| Goccy | 192066 | 662.66 MB/s | 224787 | 884 | 5.8× |
| Easyjson | 212523 | 598.88 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 418046 | 304.45 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 446879 | 211.81 MB/s | 465586 | 9714 | 2.5× |
| Stdlib | 1106366 | 115.04 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3788378 | 594.20 MB/s | 3122872 | 3081 | 6.9× |
| Lightning | 3788861 | 594.12 MB/s | 3122876 | 3081 | 6.9× |
| Sonic | 4554360 | 494.26 MB/s | 15233757 | 970 | 5.7× |
| SonicFastest | 4579704 | 491.53 MB/s | 15238660 | 970 | 5.7× |
| Goccy | 10441473 | 215.59 MB/s | 4124521 | 56532 | 2.5× |
| Easyjson | 11002176 | 204.60 MB/s | 3099809 | 2120 | 2.4× |
| LightningDecodeAny | 11679195 | 192.74 MB/s | 7938299 | 281383 | 2.2× |
| JSONV2 | 16156602 | 139.33 MB/s | 3123205 | 3083 | 1.6× |
| Stdlib | 26098785 | 86.25 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 490536 | 551.24 MB/s | 348024 | 1627 | 6.8× |
| LightningDestructive | 495605 | 545.60 MB/s | 348024 | 1627 | 6.8× |
| Sonic | 630003 | 429.21 MB/s | 473532 | 968 | 5.3× |
| SonicFastest | 635095 | 425.77 MB/s | 482535 | 968 | 5.3× |
| Goccy | 1398286 | 193.38 MB/s | 544386 | 8123 | 2.4× |
| Easyjson | 1402644 | 192.78 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1599547 | 169.05 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2111301 | 128.07 MB/s | 348149 | 1628 | 1.6× |
| Stdlib | 3354755 | 80.60 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1353273 | 1276.32 MB/s | 959848 | 5881 | 9.8× |
| Lightning | 1361768 | 1268.35 MB/s | 959889 | 5882 | 9.7× |
| Sonic | 2056182 | 840.01 MB/s | 2715263 | 4020 | 6.4× |
| SonicFastest | 2063546 | 837.01 MB/s | 2748314 | 4020 | 6.4× |
| Goccy | 2415654 | 715.00 MB/s | 2582822 | 14604 | 5.5× |
| Easyjson | 4213242 | 409.95 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4238107 | 407.54 MB/s | 1011632 | 7594 | 3.1× |
| LightningDecodeAny | 4620540 | 108.28 MB/s | 4976205 | 81466 | 2.9× |
| Stdlib | 13219350 | 130.66 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1192 | 1519.64 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1200 | 1510.30 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2521 | 718.63 MB/s | 24 | 1 | 5.6× |
| Goccy | 2796 | 648.05 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5931 | 305.53 MB/s | 3741 | 40 | 2.4× |
| SonicFastest | 5950 | 304.54 MB/s | 3750 | 40 | 2.4× |
| JSONV2 | 7838 | 231.17 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8189 | 221.15 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14054 | 128.93 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1217 | 1488.31 MB/s | 0 | 0 | 11.6× |
| LightningDestructive | 1232 | 1470.19 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2543 | 712.60 MB/s | 24 | 1 | 5.5× |
| Goccy | 2831 | 640.12 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 5965 | 303.78 MB/s | 3729 | 40 | 2.4× |
| Sonic | 5976 | 303.20 MB/s | 3713 | 40 | 2.4× |
| JSONV2 | 7880 | 229.95 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8228 | 220.11 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14058 | 128.89 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1395 | 1299.21 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1449 | 1250.10 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2774 | 653.22 MB/s | 144 | 10 | 5.0× |
| Goccy | 2834 | 639.48 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6166 | 293.87 MB/s | 3758 | 42 | 2.3× |
| Sonic | 6170 | 293.67 MB/s | 3750 | 42 | 2.3× |
| JSONV2 | 7908 | 229.14 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8204 | 220.73 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 13996 | 129.46 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 742 | 665.96 MB/s | 160 | 1 | 7.5× |
| LightningDestructive | 742 | 665.41 MB/s | 160 | 1 | 7.5× |
| Sonic | 1248 | 395.69 MB/s | 1016 | 6 | 4.4× |
| SonicFastest | 1251 | 394.76 MB/s | 1015 | 6 | 4.4× |
| LightningDecodeAny | 1695 | 290.92 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2204 | 224.18 MB/s | 448 | 3 | 2.5× |
| Goccy | 2464 | 200.51 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3181 | 155.31 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5553 | 88.96 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 473 | 485.85 MB/s | 160 | 1 | 8.7× |
| LightningDestructive | 479 | 479.76 MB/s | 160 | 1 | 8.6× |
| SonicFastest | 899 | 255.74 MB/s | 686 | 6 | 4.6× |
| Sonic | 902 | 254.96 MB/s | 685 | 6 | 4.6× |
| Easyjson | 1396 | 164.71 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1420 | 161.25 MB/s | 1536 | 30 | 2.9× |
| Goccy | 1618 | 142.18 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2353 | 97.76 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4137 | 55.59 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 72680 | 896.15 MB/s | 166212 | 102 | 7.5× |
| Lightning | 73962 | 880.61 MB/s | 172432 | 107 | 7.4× |
| SonicFastest | 96629 | 674.04 MB/s | 154599 | 75 | 5.6× |
| Sonic | 97129 | 670.57 MB/s | 155992 | 75 | 5.6× |
| Goccy | 141112 | 461.56 MB/s | 229308 | 134 | 3.9× |
| LightningDecodeAny | 186443 | 286.03 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 222478 | 292.76 MB/s | 206650 | 607 | 2.5× |
| Stdlib | 545537 | 119.39 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2599420 | 746.50 MB/s | 2846864 | 2238 | 9.0× |
| Lightning | 2656109 | 730.57 MB/s | 2846866 | 2238 | 8.8× |
| Goccy | 4740771 | 409.32 MB/s | 4064781 | 13510 | 4.9× |
| Sonic | 4761106 | 407.57 MB/s | 14606972 | 1407 | 4.9× |
| SonicFastest | 5018066 | 386.70 MB/s | 14608682 | 1407 | 4.7× |
| Easyjson | 7532083 | 257.63 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9604251 | 202.04 MB/s | 7013524 | 219937 | 2.4× |
| JSONV2 | 11226418 | 172.85 MB/s | 3237220 | 13947 | 2.1× |
| Stdlib | 23453598 | 82.74 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1104351 | 3013.38 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1776859 | 1872.87 MB/s | 2488905 | 2995 | 11.7× |
| Sonic | 2634224 | 1263.31 MB/s | 6502559 | 4248 | 7.9× |
| SonicFastest | 2634767 | 1263.05 MB/s | 6451708 | 4248 | 7.9× |
| LightningDecodeAny | 3698428 | 831.10 MB/s | 4886620 | 56892 | 5.6× |
| Goccy | 4612297 | 721.51 MB/s | 3948909 | 3816 | 4.5× |
| JSONV2 | 7324198 | 454.36 MB/s | 5364514 | 13243 | 2.8× |
| Stdlib | 20769014 | 160.23 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219899 | 1002.03 MB/s | 136896 | 228 | 9.3× |
| LightningDestructive | 221950 | 992.77 MB/s | 136896 | 228 | 9.2× |
| Sonic | 375192 | 587.29 MB/s | 300178 | 398 | 5.4× |
| SonicFastest | 379076 | 581.27 MB/s | 307059 | 398 | 5.4× |
| Goccy | 434423 | 507.22 MB/s | 364760 | 1067 | 4.7× |
| Easyjson | 547340 | 402.58 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 736988 | 298.98 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 875968 | 123.65 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2039221 | 108.05 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12801919 | 632.72 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13253890 | 611.14 MB/s | 15059838 | 51649 | 6.7× |
| Sonic | 16318620 | 496.37 MB/s | 70887350 | 40014 | 5.5× |
| SonicFastest | 16635226 | 486.92 MB/s | 70887175 | 40014 | 5.4× |
| Goccy | 23173799 | 349.53 MB/s | 17051483 | 107148 | 3.8× |
| Easyjson | 30698626 | 263.86 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 40270961 | 129.20 MB/s | 34333370 | 912810 | 2.2× |
| JSONV2 | 43947600 | 184.31 MB/s | 15233760 | 78972 | 2.0× |
| Stdlib | 89164414 | 90.84 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6077298 | 490.92 MB/s | 3985336 | 30015 | 7.8× |
| Lightning | 6185062 | 482.37 MB/s | 3985337 | 30015 | 7.7× |
| Sonic | 8718605 | 342.20 MB/s | 26564170 | 56760 | 5.4× |
| SonicFastest | 8741705 | 341.29 MB/s | 26663141 | 56760 | 5.4× |
| Goccy | 16565605 | 180.10 MB/s | 10722658 | 273654 | 2.9× |
| Easyjson | 16769435 | 177.91 MB/s | 9479448 | 30115 | 2.8× |
| LightningDecodeAny | 18713520 | 98.01 MB/s | 20023838 | 408557 | 2.5× |
| JSONV2 | 24277944 | 122.89 MB/s | 9257145 | 86278 | 2.0× |
| Stdlib | 47452533 | 62.87 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1378193 | 525.03 MB/s | 907393 | 3618 | 8.4× |
| Lightning | 1389462 | 520.77 MB/s | 907388 | 3618 | 8.3× |
| SonicFastest | 1764664 | 410.05 MB/s | 3186478 | 7226 | 6.6× |
| Sonic | 1769728 | 408.87 MB/s | 3191458 | 7226 | 6.5× |
| Easyjson | 4218991 | 171.51 MB/s | 2847905 | 3698 | 2.7× |
| LightningDecodeAny | 4322944 | 150.49 MB/s | 5752201 | 80172 | 2.7× |
| Goccy | 4740695 | 152.64 MB/s | 2807973 | 80274 | 2.4× |
| JSONV2 | 5695542 | 127.05 MB/s | 2704598 | 7318 | 2.0× |
| Stdlib | 11576063 | 62.51 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1921819 | 820.76 MB/s | 907392 | 3618 | 8.2× |
| Lightning | 1971900 | 799.92 MB/s | 907387 | 3618 | 8.0× |
| SonicFastest | 2251617 | 700.54 MB/s | 5786649 | 7226 | 7.0× |
| Sonic | 2270439 | 694.73 MB/s | 5793813 | 7226 | 6.9× |
| LightningDecodeAny | 3826291 | 196.90 MB/s | 5752203 | 80172 | 4.1× |
| Easyjson | 5540331 | 284.70 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5591524 | 282.10 MB/s | 3605268 | 80268 | 2.8× |
| JSONV2 | 6535607 | 241.35 MB/s | 2704593 | 7318 | 2.4× |
| Stdlib | 15707649 | 100.42 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 224276 | 669.37 MB/s | 81920 | 1 | 8.1× |
| Lightning | 224503 | 668.70 MB/s | 81920 | 1 | 8.1× |
| Sonic | 267522 | 561.16 MB/s | 246487 | 6 | 6.8× |
| SonicFastest | 268233 | 559.68 MB/s | 248310 | 6 | 6.8× |
| LightningDecodeAny | 470370 | 319.16 MB/s | 746005 | 10020 | 3.9× |
| Goccy | 861575 | 174.24 MB/s | 324603 | 10004 | 2.1× |
| JSONV2 | 1093568 | 137.28 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1821763 | 82.41 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33292 | 844.56 MB/s | 30272 | 105 | 9.1× |
| LightningDestructive | 33548 | 838.12 MB/s | 30144 | 103 | 9.1× |
| SonicFastest | 64011 | 439.25 MB/s | 48070 | 103 | 4.7× |
| Sonic | 64147 | 438.32 MB/s | 48203 | 103 | 4.7× |
| Easyjson | 68340 | 411.43 MB/s | 32304 | 138 | 4.4× |
| Goccy | 71939 | 390.85 MB/s | 59262 | 188 | 4.2× |
| JSONV2 | 134729 | 208.69 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 154339 | 182.18 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303706 | 92.58 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1952 | 1192.71 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2014 | 1155.72 MB/s | 32 | 1 | 11.2× |
| Goccy | 4110 | 566.45 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4196 | 554.87 MB/s | 192 | 2 | 5.4× |
| Sonic | 5085 | 457.77 MB/s | 4270 | 6 | 4.5× |
| SonicFastest | 5117 | 454.92 MB/s | 4350 | 6 | 4.4× |
| JSONV2 | 8437 | 275.93 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10184 | 165.45 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22657 | 102.75 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221 | 856.09 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 222 | 850.56 MB/s | 0 | 0 | 11.0× |
| Goccy | 381 | 496.11 MB/s | 304 | 2 | 6.4× |
| Easyjson | 492 | 384.49 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 794 | 238.01 MB/s | 500 | 4 | 3.1× |
| Sonic | 796 | 237.48 MB/s | 495 | 4 | 3.1× |
| JSONV2 | 1037 | 182.24 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1230 | 108.96 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2453 | 77.04 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1549 | 1414.42 MB/s | 0 | 0 | 10.3× |
| LightningDestructive | 1560 | 1404.07 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3210 | 682.53 MB/s | 24 | 1 | 5.0× |
| Goccy | 3247 | 674.84 MB/s | 2864 | 4 | 4.9× |
| Sonic | 6449 | 339.73 MB/s | 4109 | 40 | 2.5× |
| SonicFastest | 6449 | 339.74 MB/s | 4132 | 40 | 2.5× |
| JSONV2 | 8173 | 268.06 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8478 | 213.62 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16008 | 136.87 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 701076 | 728.13 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 727864 | 701.33 MB/s | 703779 | 1012 | 8.3× |
| Goccy | 1143667 | 446.35 MB/s | 1138180 | 5006 | 5.3× |
| SonicFastest | 1158974 | 440.46 MB/s | 887410 | 2006 | 5.2× |
| Sonic | 1159942 | 440.09 MB/s | 888040 | 2006 | 5.2× |
| Easyjson | 1533253 | 332.94 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3233481 | 157.87 MB/s | 1076008 | 12646 | 1.9× |
| LightningDecodeAny | 3509968 | 131.47 MB/s | 2785928 | 66022 | 1.7× |
| Stdlib | 6019532 | 84.80 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14650.67 MB/s | 0 | 0 | 82.5× |
| LightningDestructive | 1394 | 14199.28 MB/s | 0 | 0 | 80.0× |
| Goccy | 19888 | 995.04 MB/s | 20491 | 2 | 5.6× |
| SonicFastest | 27759 | 712.88 MB/s | 21910 | 4 | 4.0× |
| Sonic | 27823 | 711.25 MB/s | 21978 | 4 | 4.0× |
| JSONV2 | 29626 | 667.95 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 78911 | 250.76 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86040 | 230.00 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111485 | 177.50 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2677 | 6770.34 MB/s | 0 | 0 | 38.5× |
| Lightning | 2822 | 6421.50 MB/s | 432 | 2 | 36.5× |
| Easyjson | 3961 | 4575.43 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 10164 | 1783.15 MB/s | 24373 | 6 | 10.1× |
| Sonic | 10233 | 1771.21 MB/s | 24593 | 6 | 10.1× |
| Goccy | 15914 | 1138.90 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 17455 | 1024.47 MB/s | 29088 | 191 | 5.9× |
| JSONV2 | 45463 | 398.66 MB/s | 16500 | 50 | 2.3× |
| Stdlib | 103141 | 175.72 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2171443 | 924.96 MB/s | 2960388 | 7411 | 8.6× |
| Lightning | 2230453 | 900.49 MB/s | 2962101 | 7417 | 8.4× |
| Goccy | 4205392 | 477.60 MB/s | 5412540 | 15831 | 4.4× |
| Sonic | 4388948 | 457.63 MB/s | 10955582 | 13683 | 4.2× |
| SonicFastest | 4390785 | 457.43 MB/s | 10956122 | 13683 | 4.2× |
| Easyjson | 4876730 | 411.85 MB/s | 2981484 | 7439 | 3.8× |
| JSONV2 | 6910495 | 290.64 MB/s | 3173684 | 14563 | 2.7× |
| LightningDecodeAny | 7281207 | 156.88 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18643834 | 107.73 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 886 | 619.93 MB/s | 480 | 1 | 6.4× |
| LightningDestructive | 887 | 619.01 MB/s | 480 | 1 | 6.4× |
| LightningDecodeAny | 1885 | 290.68 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2144 | 256.01 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2686 | 204.40 MB/s | 1952 | 26 | 2.1× |
| SonicFastest | 2695 | 203.74 MB/s | 1944 | 26 | 2.1× |
| Goccy | 2991 | 183.55 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3260 | 168.41 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5638 | 97.37 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 492007 | 1283.55 MB/s | 364472 | 566 | 11.0× |
| Lightning | 545577 | 1157.52 MB/s | 413001 | 878 | 9.9× |
| Sonic | 1008328 | 626.30 MB/s | 1015676 | 1102 | 5.3× |
| SonicFastest | 1010742 | 624.80 MB/s | 1025421 | 1102 | 5.3× |
| Easyjson | 1140846 | 553.55 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1153713 | 547.38 MB/s | 987666 | 1201 | 4.7× |
| JSONV2 | 2151383 | 293.54 MB/s | 571615 | 3144 | 2.5× |
| LightningDecodeAny | 2344225 | 199.17 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5388287 | 117.20 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 715395 | 786.15 MB/s | 544254 | 448 | 7.4× |
| Lightning | 877876 | 640.65 MB/s | 767621 | 1254 | 6.0× |
| SonicFastest | 1014572 | 554.33 MB/s | 944055 | 1476 | 5.2× |
| Sonic | 1019949 | 551.41 MB/s | 955597 | 1476 | 5.2× |
| Goccy | 1324800 | 424.52 MB/s | 1036509 | 1029 | 4.0× |
| Easyjson | 1731448 | 324.82 MB/s | 775153 | 1254 | 3.0× |
| LightningDecodeAny | 2694574 | 208.72 MB/s | 2114151 | 30295 | 2.0× |
| JSONV2 | 2749186 | 204.57 MB/s | 927444 | 3482 | 1.9× |
| Stdlib | 5275823 | 106.60 MB/s | 1011673 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 648044 | 822.75 MB/s | 333416 | 2084 | 8.5× |
| Lightning | 671611 | 793.88 MB/s | 368224 | 2293 | 8.2× |
| Easyjson | 1112254 | 479.37 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1147018 | 464.84 MB/s | 1034462 | 4351 | 4.8× |
| Sonic | 1149639 | 463.78 MB/s | 1049991 | 4351 | 4.8× |
| Goccy | 1291392 | 412.87 MB/s | 1167221 | 5409 | 4.2× |
| JSONV2 | 2529563 | 210.78 MB/s | 745447 | 13288 | 2.2× |
| LightningDecodeAny | 3349340 | 159.19 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5476554 | 97.36 MB/s | 798692 | 17133 | 1.0× |
