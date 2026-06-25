# JSON Deserialization Benchmarks

- generated 2026-06-25T21:22:37Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104697 | 1215.66 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 105040 | 1211.68 MB/s | 49280 | 2 | 10.5× |
| Sonic | 182338 | 698.02 MB/s | 194907 | 10 | 6.1× |
| SonicFastest | 183599 | 693.22 MB/s | 196627 | 10 | 6.0× |
| Goccy | 190453 | 668.28 MB/s | 224789 | 884 | 5.8× |
| Easyjson | 212012 | 600.32 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 418757 | 303.94 MB/s | 195120 | 1805 | 2.6× |
| LightningDecodeAny | 453778 | 208.59 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1106352 | 115.04 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3753285 | 599.75 MB/s | 3122872 | 3081 | 7.0× |
| Lightning | 3842881 | 585.77 MB/s | 3122875 | 3081 | 6.8× |
| SonicFastest | 4499679 | 500.27 MB/s | 15235403 | 970 | 5.8× |
| Sonic | 4541987 | 495.61 MB/s | 15233757 | 970 | 5.8× |
| Goccy | 10440727 | 215.60 MB/s | 4123585 | 56532 | 2.5× |
| Easyjson | 11064728 | 203.44 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11552589 | 194.85 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16320990 | 137.92 MB/s | 3123215 | 3083 | 1.6× |
| Stdlib | 26136506 | 86.13 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 490525 | 551.25 MB/s | 348024 | 1627 | 6.8× |
| LightningDestructive | 496228 | 544.92 MB/s | 348024 | 1627 | 6.8× |
| Sonic | 635081 | 425.78 MB/s | 481915 | 968 | 5.3× |
| SonicFastest | 637680 | 424.04 MB/s | 488771 | 968 | 5.3× |
| Goccy | 1405793 | 192.35 MB/s | 542479 | 8122 | 2.4× |
| Easyjson | 1436074 | 188.29 MB/s | 330272 | 749 | 2.3× |
| LightningDecodeAny | 1618553 | 167.06 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2105292 | 128.44 MB/s | 348154 | 1628 | 1.6× |
| Stdlib | 3354995 | 80.60 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1366453 | 1264.01 MB/s | 959848 | 5881 | 9.7× |
| Lightning | 1377143 | 1254.19 MB/s | 959890 | 5882 | 9.6× |
| SonicFastest | 2071114 | 833.95 MB/s | 2687685 | 4020 | 6.4× |
| Sonic | 2071174 | 833.93 MB/s | 2732369 | 4020 | 6.4× |
| Goccy | 2361794 | 731.31 MB/s | 2582524 | 14604 | 5.6× |
| Easyjson | 4227900 | 408.53 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4352175 | 396.86 MB/s | 1011634 | 7594 | 3.0× |
| LightningDecodeAny | 4730146 | 105.77 MB/s | 4976205 | 81466 | 2.8× |
| Stdlib | 13238297 | 130.47 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1203 | 1506.82 MB/s | 0 | 0 | 11.7× |
| LightningDestructive | 1218 | 1487.42 MB/s | 0 | 0 | 11.5× |
| Easyjson | 2544 | 712.26 MB/s | 24 | 1 | 5.5× |
| Goccy | 2876 | 630.07 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6005 | 301.75 MB/s | 3741 | 40 | 2.3× |
| SonicFastest | 6091 | 297.50 MB/s | 3790 | 40 | 2.3× |
| JSONV2 | 7884 | 229.84 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8258 | 219.31 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14064 | 128.84 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1218 | 1487.96 MB/s | 0 | 0 | 11.6× |
| LightningDestructive | 1231 | 1471.39 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2549 | 710.97 MB/s | 24 | 1 | 5.5× |
| Goccy | 2809 | 645.18 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5947 | 304.67 MB/s | 3750 | 40 | 2.4× |
| SonicFastest | 5957 | 304.16 MB/s | 3780 | 40 | 2.4× |
| JSONV2 | 7829 | 231.44 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8233 | 219.97 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14072 | 128.77 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1402 | 1292.73 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1440 | 1258.20 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2739 | 661.57 MB/s | 144 | 10 | 5.1× |
| Goccy | 2867 | 631.98 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6167 | 293.81 MB/s | 3778 | 42 | 2.3× |
| Sonic | 6174 | 293.47 MB/s | 3794 | 42 | 2.3× |
| JSONV2 | 7965 | 227.50 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8222 | 220.25 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14031 | 129.15 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 739 | 668.69 MB/s | 160 | 1 | 7.5× |
| Lightning | 742 | 665.34 MB/s | 160 | 1 | 7.4× |
| Sonic | 1237 | 399.28 MB/s | 981 | 6 | 4.5× |
| SonicFastest | 1237 | 399.29 MB/s | 990 | 6 | 4.5× |
| LightningDecodeAny | 1671 | 295.07 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2206 | 223.91 MB/s | 448 | 3 | 2.5× |
| Goccy | 2444 | 202.09 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3169 | 155.87 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5529 | 89.34 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 476 | 483.21 MB/s | 160 | 1 | 8.8× |
| LightningDestructive | 481 | 478.03 MB/s | 160 | 1 | 8.7× |
| SonicFastest | 906 | 253.93 MB/s | 690 | 6 | 4.6× |
| Sonic | 907 | 253.59 MB/s | 689 | 6 | 4.6× |
| Easyjson | 1405 | 163.76 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1459 | 156.93 MB/s | 1536 | 30 | 2.9× |
| Goccy | 1636 | 140.55 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2371 | 97.01 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4166 | 55.21 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 74276 | 876.89 MB/s | 166212 | 102 | 7.4× |
| Lightning | 75511 | 862.55 MB/s | 172433 | 107 | 7.3× |
| SonicFastest | 98917 | 658.45 MB/s | 155458 | 75 | 5.5× |
| Sonic | 99061 | 657.49 MB/s | 156138 | 75 | 5.5× |
| Goccy | 145928 | 446.33 MB/s | 229282 | 134 | 3.8× |
| LightningDecodeAny | 189399 | 281.57 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 225576 | 288.74 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 548411 | 118.76 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2626190 | 738.89 MB/s | 2846865 | 2238 | 8.9× |
| Lightning | 2679677 | 724.14 MB/s | 2846866 | 2238 | 8.8× |
| Sonic | 4685346 | 414.16 MB/s | 14608573 | 1407 | 5.0× |
| Goccy | 4785980 | 405.45 MB/s | 4065914 | 13510 | 4.9× |
| SonicFastest | 4962255 | 391.05 MB/s | 14608609 | 1407 | 4.7× |
| Easyjson | 7535302 | 257.52 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 9693405 | 200.18 MB/s | 7013523 | 219937 | 2.4× |
| JSONV2 | 11376010 | 170.58 MB/s | 3237223 | 13947 | 2.1× |
| Stdlib | 23453768 | 82.74 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1108116 | 3003.14 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1827929 | 1820.55 MB/s | 2488904 | 2995 | 11.4× |
| Sonic | 2661381 | 1250.42 MB/s | 6483622 | 4248 | 7.8× |
| SonicFastest | 2675914 | 1243.62 MB/s | 6444391 | 4248 | 7.8× |
| LightningDecodeAny | 3784724 | 812.15 MB/s | 4886622 | 56892 | 5.5× |
| Goccy | 4520520 | 736.16 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7392248 | 450.18 MB/s | 5364513 | 13243 | 2.8× |
| Stdlib | 20838208 | 159.70 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221737 | 993.73 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 224626 | 980.94 MB/s | 136897 | 228 | 9.1× |
| SonicFastest | 392226 | 561.78 MB/s | 318983 | 398 | 5.2× |
| Sonic | 395135 | 557.65 MB/s | 327441 | 398 | 5.2× |
| Goccy | 449866 | 489.80 MB/s | 364323 | 1067 | 4.5× |
| Easyjson | 546159 | 403.45 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 734381 | 300.04 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 896717 | 120.79 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2043864 | 107.81 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12914563 | 627.20 MB/s | 15059832 | 51649 | 6.9× |
| Lightning | 13328519 | 607.72 MB/s | 15059836 | 51649 | 6.7× |
| Sonic | 16908981 | 479.04 MB/s | 70901962 | 40014 | 5.3× |
| SonicFastest | 16910290 | 479.00 MB/s | 70887598 | 40014 | 5.3× |
| Goccy | 23776615 | 340.67 MB/s | 17107438 | 107149 | 3.8× |
| Easyjson | 31069468 | 260.71 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 40979071 | 126.97 MB/s | 34333352 | 912810 | 2.2× |
| JSONV2 | 44034775 | 183.95 MB/s | 15233768 | 78972 | 2.0× |
| Stdlib | 89676029 | 90.33 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6103583 | 488.81 MB/s | 3985336 | 30015 | 7.8× |
| Lightning | 6221381 | 479.55 MB/s | 3985336 | 30015 | 7.6× |
| SonicFastest | 8681091 | 343.67 MB/s | 26547088 | 56760 | 5.5× |
| Sonic | 8708353 | 342.60 MB/s | 26551728 | 56760 | 5.4× |
| Easyjson | 16550163 | 180.27 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16609828 | 179.62 MB/s | 10554531 | 273645 | 2.9× |
| LightningDecodeAny | 18569702 | 98.77 MB/s | 20023839 | 408557 | 2.6× |
| JSONV2 | 24527369 | 121.64 MB/s | 9257156 | 86278 | 1.9× |
| Stdlib | 47429762 | 62.90 MB/s | 9258096 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1376698 | 525.60 MB/s | 907393 | 3618 | 8.5× |
| Lightning | 1399049 | 517.21 MB/s | 907389 | 3618 | 8.4× |
| SonicFastest | 1790518 | 404.13 MB/s | 3199150 | 7226 | 6.5× |
| Sonic | 1791660 | 403.87 MB/s | 3181304 | 7226 | 6.5× |
| Easyjson | 4309583 | 167.90 MB/s | 2847906 | 3698 | 2.7× |
| LightningDecodeAny | 4414870 | 147.36 MB/s | 5752201 | 80172 | 2.7× |
| Goccy | 4699098 | 153.99 MB/s | 2811367 | 80274 | 2.5× |
| JSONV2 | 5973263 | 121.14 MB/s | 2704637 | 7318 | 2.0× |
| Stdlib | 11717175 | 61.76 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1924789 | 819.49 MB/s | 907393 | 3618 | 8.2× |
| Lightning | 1973371 | 799.32 MB/s | 907387 | 3618 | 8.0× |
| SonicFastest | 2249626 | 701.16 MB/s | 5783886 | 7226 | 7.0× |
| Sonic | 2258840 | 698.30 MB/s | 5790146 | 7226 | 7.0× |
| LightningDecodeAny | 3889673 | 193.69 MB/s | 5752201 | 80172 | 4.0× |
| Goccy | 5596558 | 281.84 MB/s | 3540691 | 80265 | 2.8× |
| Easyjson | 5624763 | 280.43 MB/s | 2847907 | 3698 | 2.8× |
| JSONV2 | 6659323 | 236.86 MB/s | 2704596 | 7318 | 2.4× |
| Stdlib | 15749493 | 100.15 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 223609 | 671.37 MB/s | 81920 | 1 | 8.2× |
| LightningDestructive | 224821 | 667.75 MB/s | 81920 | 1 | 8.1× |
| SonicFastest | 270867 | 554.24 MB/s | 245762 | 6 | 6.7× |
| Sonic | 273287 | 549.33 MB/s | 253785 | 6 | 6.7× |
| LightningDecodeAny | 477706 | 314.25 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 856574 | 175.26 MB/s | 325084 | 10004 | 2.1× |
| JSONV2 | 1106404 | 135.69 MB/s | 357715 | 20 | 1.6× |
| Stdlib | 1823922 | 82.31 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32913 | 854.30 MB/s | 30272 | 105 | 9.2× |
| LightningDestructive | 33068 | 850.27 MB/s | 30144 | 103 | 9.2× |
| SonicFastest | 63833 | 440.47 MB/s | 46880 | 103 | 4.8× |
| Sonic | 63856 | 440.32 MB/s | 47210 | 103 | 4.8× |
| Easyjson | 68379 | 411.19 MB/s | 32304 | 138 | 4.4× |
| Goccy | 71358 | 394.03 MB/s | 59183 | 188 | 4.3× |
| JSONV2 | 134934 | 208.38 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 152377 | 184.52 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303909 | 92.52 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1952 | 1192.42 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2005 | 1161.17 MB/s | 32 | 1 | 11.3× |
| Goccy | 4149 | 561.15 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4229 | 550.42 MB/s | 192 | 2 | 5.4× |
| Sonic | 5128 | 454.01 MB/s | 4283 | 6 | 4.4× |
| SonicFastest | 5136 | 453.28 MB/s | 4321 | 6 | 4.4× |
| JSONV2 | 8424 | 276.37 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10244 | 164.48 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22673 | 102.68 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 863.57 MB/s | 0 | 0 | 11.2× |
| LightningDestructive | 221 | 855.26 MB/s | 0 | 0 | 11.0× |
| Goccy | 377 | 501.01 MB/s | 304 | 2 | 6.5× |
| Easyjson | 485 | 389.98 MB/s | 0 | 0 | 5.0× |
| Sonic | 800 | 236.12 MB/s | 494 | 4 | 3.0× |
| SonicFastest | 801 | 235.83 MB/s | 502 | 4 | 3.0× |
| JSONV2 | 1023 | 184.82 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1218 | 109.98 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2441 | 77.41 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1540 | 1422.98 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1558 | 1406.67 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3189 | 687.04 MB/s | 24 | 1 | 5.0× |
| Goccy | 3317 | 660.58 MB/s | 2864 | 4 | 4.8× |
| Sonic | 6465 | 338.91 MB/s | 4126 | 40 | 2.5× |
| SonicFastest | 6470 | 338.64 MB/s | 4170 | 40 | 2.5× |
| JSONV2 | 8089 | 270.87 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8586 | 210.92 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16033 | 136.66 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 706945 | 722.09 MB/s | 703778 | 1012 | 8.5× |
| Lightning | 735509 | 694.04 MB/s | 703778 | 1012 | 8.2× |
| Sonic | 1162594 | 439.08 MB/s | 902311 | 2006 | 5.2× |
| SonicFastest | 1169767 | 436.39 MB/s | 911964 | 2006 | 5.1× |
| Goccy | 1172395 | 435.41 MB/s | 1141510 | 5007 | 5.1× |
| Easyjson | 1546308 | 330.13 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3204901 | 159.28 MB/s | 1076015 | 12646 | 1.9× |
| LightningDecodeAny | 3533916 | 130.58 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6007061 | 84.98 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14643.01 MB/s | 0 | 0 | 98.5× |
| LightningDestructive | 1389 | 14246.06 MB/s | 0 | 0 | 95.8× |
| Goccy | 19970 | 990.92 MB/s | 20491 | 2 | 6.7× |
| Sonic | 27436 | 721.27 MB/s | 22840 | 4 | 4.9× |
| SonicFastest | 27696 | 714.50 MB/s | 23320 | 4 | 4.8× |
| JSONV2 | 29798 | 664.10 MB/s | 8 | 1 | 4.5× |
| LightningDecodeAny | 78830 | 251.02 MB/s | 117104 | 2019 | 1.7× |
| Easyjson | 90419 | 218.86 MB/s | 0 | 0 | 1.5× |
| Stdlib | 133087 | 148.69 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2658 | 6819.10 MB/s | 0 | 0 | 38.6× |
| Lightning | 2829 | 6406.18 MB/s | 432 | 2 | 36.3× |
| Easyjson | 3942 | 4597.75 MB/s | 432 | 2 | 26.1× |
| Sonic | 9994 | 1813.42 MB/s | 22802 | 6 | 10.3× |
| SonicFastest | 10137 | 1787.84 MB/s | 23202 | 6 | 10.1× |
| Goccy | 15820 | 1145.64 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 16769 | 1066.39 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 46057 | 393.52 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102703 | 176.47 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2256891 | 889.94 MB/s | 2960388 | 7411 | 8.3× |
| Lightning | 2339858 | 858.38 MB/s | 2962101 | 7417 | 8.0× |
| Goccy | 4293475 | 467.80 MB/s | 5411735 | 15831 | 4.4× |
| Sonic | 4457779 | 450.56 MB/s | 10945930 | 13683 | 4.2× |
| SonicFastest | 4477818 | 448.54 MB/s | 10931473 | 13683 | 4.2× |
| Easyjson | 4925950 | 407.74 MB/s | 2981482 | 7439 | 3.8× |
| JSONV2 | 6948364 | 289.06 MB/s | 3173685 | 14563 | 2.7× |
| LightningDecodeAny | 7220775 | 158.20 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18729125 | 107.24 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 889 | 617.59 MB/s | 480 | 1 | 6.4× |
| LightningDestructive | 892 | 615.14 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1910 | 286.94 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2147 | 255.72 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2679 | 204.94 MB/s | 1938 | 26 | 2.1× |
| Sonic | 2687 | 204.30 MB/s | 1938 | 26 | 2.1× |
| Goccy | 3039 | 180.63 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3277 | 167.52 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5658 | 97.03 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 494965 | 1275.88 MB/s | 364472 | 566 | 10.9× |
| Lightning | 550837 | 1146.46 MB/s | 413001 | 878 | 9.8× |
| SonicFastest | 1011472 | 624.35 MB/s | 989957 | 1102 | 5.3× |
| Sonic | 1022205 | 617.80 MB/s | 1009263 | 1102 | 5.3× |
| Easyjson | 1142706 | 552.65 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1169675 | 539.91 MB/s | 985734 | 1201 | 4.6× |
| JSONV2 | 2141990 | 294.83 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2403029 | 194.30 MB/s | 2010200 | 30295 | 2.2× |
| Stdlib | 5394609 | 117.06 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 721631 | 779.36 MB/s | 544253 | 448 | 7.3× |
| Lightning | 895188 | 628.26 MB/s | 767621 | 1254 | 5.9× |
| Sonic | 1040452 | 540.54 MB/s | 958957 | 1476 | 5.1× |
| SonicFastest | 1040538 | 540.50 MB/s | 960214 | 1476 | 5.1× |
| Goccy | 1371477 | 410.07 MB/s | 1040855 | 1030 | 3.9× |
| Easyjson | 1753201 | 320.79 MB/s | 775154 | 1254 | 3.0× |
| JSONV2 | 2794061 | 201.29 MB/s | 927439 | 3482 | 1.9× |
| LightningDecodeAny | 2809119 | 200.21 MB/s | 2114151 | 30295 | 1.9× |
| Stdlib | 5289398 | 106.33 MB/s | 1011671 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 658471 | 809.72 MB/s | 333416 | 2084 | 8.3× |
| Lightning | 680720 | 783.26 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1116997 | 477.33 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1156610 | 460.98 MB/s | 1037610 | 4351 | 4.7× |
| SonicFastest | 1160825 | 459.31 MB/s | 1038030 | 4351 | 4.7× |
| Goccy | 1301052 | 409.81 MB/s | 1167221 | 5409 | 4.2× |
| JSONV2 | 2541482 | 209.79 MB/s | 745451 | 13288 | 2.2× |
| LightningDecodeAny | 3431755 | 155.37 MB/s | 2674619 | 50596 | 1.6× |
| Stdlib | 5490978 | 97.10 MB/s | 798692 | 17133 | 1.0× |
