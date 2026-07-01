# JSON Deserialization Benchmarks

- generated 2026-07-01T21:41:22Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104806 | 1214.39 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 105511 | 1206.27 MB/s | 49280 | 2 | 10.5× |
| Sonic | 184491 | 689.87 MB/s | 195614 | 10 | 6.0× |
| SonicFastest | 188085 | 676.69 MB/s | 202188 | 10 | 5.9× |
| Goccy | 199255 | 638.76 MB/s | 225264 | 884 | 5.6× |
| Easyjson | 213299 | 596.70 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 426210 | 298.62 MB/s | 195120 | 1805 | 2.6× |
| LightningDecodeAny | 461976 | 204.89 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1107754 | 114.89 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3717169 | 605.58 MB/s | 3122873 | 3081 | 7.0× |
| Lightning | 3780220 | 595.48 MB/s | 3122876 | 3081 | 6.9× |
| Sonic | 4637371 | 485.42 MB/s | 15233804 | 970 | 5.6× |
| SonicFastest | 4638306 | 485.32 MB/s | 15245592 | 970 | 5.6× |
| Goccy | 10234202 | 219.95 MB/s | 4108704 | 56532 | 2.6× |
| LightningDecodeAny | 11153560 | 201.82 MB/s | 7938299 | 281383 | 2.3× |
| Easyjson | 11368615 | 198.01 MB/s | 3099808 | 2120 | 2.3× |
| JSONV2 | 16160533 | 139.29 MB/s | 3123215 | 3083 | 1.6× |
| Stdlib | 26140293 | 86.11 MB/s | 3123398 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 488250 | 553.82 MB/s | 348024 | 1627 | 6.9× |
| LightningDestructive | 495171 | 546.08 MB/s | 348024 | 1627 | 6.8× |
| SonicFastest | 638574 | 423.45 MB/s | 472170 | 968 | 5.3× |
| Sonic | 642019 | 421.18 MB/s | 480365 | 968 | 5.2× |
| Goccy | 1409942 | 191.78 MB/s | 543459 | 8122 | 2.4× |
| Easyjson | 1439737 | 187.81 MB/s | 330272 | 749 | 2.3× |
| LightningDecodeAny | 1615847 | 167.34 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2102853 | 128.59 MB/s | 348159 | 1628 | 1.6× |
| Stdlib | 3359288 | 80.49 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1369633 | 1261.07 MB/s | 959848 | 5881 | 9.7× |
| Lightning | 1390559 | 1242.09 MB/s | 959890 | 5882 | 9.5× |
| Sonic | 2084885 | 828.44 MB/s | 2665078 | 4020 | 6.4× |
| SonicFastest | 2090808 | 826.09 MB/s | 2671535 | 4020 | 6.3× |
| Goccy | 2446456 | 706.00 MB/s | 2582928 | 14605 | 5.4× |
| Easyjson | 4226500 | 408.66 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4298464 | 401.82 MB/s | 1011635 | 7594 | 3.1× |
| LightningDecodeAny | 4627110 | 108.12 MB/s | 4976206 | 81466 | 2.9× |
| Stdlib | 13272932 | 130.13 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1190 | 1523.13 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1205 | 1503.46 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2555 | 709.23 MB/s | 24 | 1 | 5.5× |
| Goccy | 2842 | 637.63 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6092 | 297.44 MB/s | 3827 | 40 | 2.3× |
| Sonic | 6100 | 297.03 MB/s | 3771 | 40 | 2.3× |
| JSONV2 | 7835 | 231.27 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8305 | 218.07 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14047 | 128.99 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1223 | 1481.81 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1234 | 1468.80 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2546 | 711.82 MB/s | 24 | 1 | 5.5× |
| Goccy | 2819 | 642.86 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 5986 | 302.70 MB/s | 3790 | 40 | 2.3× |
| Sonic | 6016 | 301.18 MB/s | 3792 | 40 | 2.3× |
| JSONV2 | 7815 | 231.88 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8297 | 218.26 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14067 | 128.81 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1390 | 1304.01 MB/s | 144 | 10 | 10.1× |
| LightningDestructive | 1444 | 1255.18 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2759 | 656.72 MB/s | 144 | 10 | 5.1× |
| Goccy | 2887 | 627.68 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6119 | 296.11 MB/s | 3767 | 42 | 2.3× |
| SonicFastest | 6140 | 295.13 MB/s | 3808 | 42 | 2.3× |
| JSONV2 | 7995 | 226.64 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8282 | 218.67 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14073 | 128.76 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 697 | 709.06 MB/s | 160 | 1 | 7.9× |
| LightningDestructive | 698 | 707.25 MB/s | 160 | 1 | 7.9× |
| SonicFastest | 1243 | 397.27 MB/s | 984 | 6 | 4.4× |
| Sonic | 1244 | 397.00 MB/s | 978 | 6 | 4.4× |
| LightningDecodeAny | 1655 | 297.96 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2212 | 223.34 MB/s | 448 | 3 | 2.5× |
| Goccy | 2393 | 206.48 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3237 | 152.61 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5516 | 89.56 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 426 | 539.57 MB/s | 160 | 1 | 9.6× |
| Lightning | 428 | 537.14 MB/s | 160 | 1 | 9.6× |
| Sonic | 884 | 260.09 MB/s | 658 | 6 | 4.7× |
| SonicFastest | 891 | 258.03 MB/s | 666 | 6 | 4.6× |
| Easyjson | 1394 | 164.94 MB/s | 448 | 3 | 3.0× |
| LightningDecodeAny | 1401 | 163.46 MB/s | 1536 | 30 | 2.9× |
| Goccy | 1580 | 145.56 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2421 | 94.98 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4113 | 55.91 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 75260 | 865.42 MB/s | 166212 | 102 | 7.3× |
| Lightning | 76447 | 851.98 MB/s | 172433 | 107 | 7.2× |
| Sonic | 99144 | 656.94 MB/s | 155793 | 75 | 5.5× |
| SonicFastest | 99396 | 655.28 MB/s | 156782 | 75 | 5.5× |
| Goccy | 151268 | 430.57 MB/s | 228686 | 134 | 3.6× |
| LightningDecodeAny | 189683 | 281.15 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 231906 | 280.85 MB/s | 206653 | 607 | 2.4× |
| Stdlib | 547818 | 118.89 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2635458 | 736.29 MB/s | 2846864 | 2238 | 8.9× |
| Lightning | 2703576 | 717.74 MB/s | 2846867 | 2238 | 8.7× |
| Sonic | 4741507 | 409.25 MB/s | 14606973 | 1407 | 5.0× |
| SonicFastest | 4763766 | 407.34 MB/s | 14606973 | 1407 | 4.9× |
| Goccy | 4772832 | 406.57 MB/s | 4064895 | 13510 | 4.9× |
| Easyjson | 7551430 | 256.97 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9444384 | 205.46 MB/s | 7013524 | 219937 | 2.5× |
| JSONV2 | 11368202 | 170.69 MB/s | 3237224 | 13947 | 2.1× |
| Stdlib | 23521342 | 82.50 MB/s | 3551321 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1110639 | 2996.32 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1830689 | 1817.80 MB/s | 2488905 | 2995 | 11.4× |
| Sonic | 2730144 | 1218.92 MB/s | 6448056 | 4248 | 7.6× |
| SonicFastest | 2744744 | 1212.44 MB/s | 6414993 | 4248 | 7.6× |
| LightningDecodeAny | 3799584 | 808.97 MB/s | 4886620 | 56892 | 5.5× |
| Goccy | 4503848 | 738.89 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7554214 | 440.53 MB/s | 5364509 | 13243 | 2.8× |
| Stdlib | 20865784 | 159.49 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222307 | 991.18 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 225332 | 977.87 MB/s | 136897 | 228 | 9.1× |
| SonicFastest | 377240 | 584.10 MB/s | 293046 | 398 | 5.4× |
| Sonic | 378307 | 582.45 MB/s | 295331 | 398 | 5.4× |
| Goccy | 435328 | 506.16 MB/s | 364185 | 1067 | 4.7× |
| Easyjson | 545617 | 403.85 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 734486 | 300.00 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 898236 | 120.58 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2044210 | 107.79 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12888300 | 628.48 MB/s | 15059832 | 51649 | 6.9× |
| Lightning | 13407505 | 604.14 MB/s | 15059836 | 51649 | 6.7× |
| Sonic | 16858375 | 480.48 MB/s | 70902007 | 40014 | 5.3× |
| SonicFastest | 16886759 | 479.67 MB/s | 70901680 | 40014 | 5.3× |
| Goccy | 23362520 | 346.71 MB/s | 17108998 | 107148 | 3.8× |
| Easyjson | 31026422 | 261.07 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 40504471 | 128.46 MB/s | 34333352 | 912810 | 2.2× |
| JSONV2 | 44270990 | 182.96 MB/s | 15233751 | 78972 | 2.0× |
| Stdlib | 89545481 | 90.46 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5688677 | 524.46 MB/s | 3985336 | 30015 | 8.3× |
| Lightning | 5820492 | 512.58 MB/s | 3985339 | 30015 | 8.1× |
| Sonic | 8683844 | 343.57 MB/s | 26575742 | 56760 | 5.4× |
| SonicFastest | 8704470 | 342.75 MB/s | 26460880 | 56760 | 5.4× |
| Easyjson | 16460323 | 181.25 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16594092 | 179.79 MB/s | 10568824 | 273648 | 2.8× |
| LightningDecodeAny | 18674772 | 98.22 MB/s | 20023838 | 408557 | 2.5× |
| JSONV2 | 24260159 | 122.98 MB/s | 9257166 | 86278 | 1.9× |
| Stdlib | 47150558 | 63.28 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1237719 | 584.62 MB/s | 907392 | 3618 | 9.5× |
| Lightning | 1253789 | 577.13 MB/s | 907388 | 3618 | 9.3× |
| Sonic | 1802084 | 401.53 MB/s | 3183769 | 7226 | 6.5× |
| SonicFastest | 1802456 | 401.45 MB/s | 3188538 | 7226 | 6.5× |
| Easyjson | 4285884 | 168.83 MB/s | 2847906 | 3698 | 2.7× |
| LightningDecodeAny | 4479564 | 145.23 MB/s | 5752202 | 80172 | 2.6× |
| Goccy | 4782978 | 151.29 MB/s | 2940190 | 80281 | 2.4× |
| JSONV2 | 5676937 | 127.46 MB/s | 2704624 | 7318 | 2.1× |
| Stdlib | 11702606 | 61.83 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1848653 | 853.24 MB/s | 907393 | 3618 | 8.5× |
| Lightning | 1897305 | 831.37 MB/s | 907386 | 3618 | 8.3× |
| SonicFastest | 2256282 | 699.09 MB/s | 5783870 | 7226 | 7.0× |
| Sonic | 2262071 | 697.30 MB/s | 5789604 | 7226 | 6.9× |
| LightningDecodeAny | 3942282 | 191.11 MB/s | 5752201 | 80172 | 4.0× |
| Easyjson | 5580278 | 282.67 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5726337 | 275.46 MB/s | 3574088 | 80267 | 2.7× |
| JSONV2 | 6422808 | 245.59 MB/s | 2704593 | 7318 | 2.4× |
| Stdlib | 15716320 | 100.36 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 209411 | 716.89 MB/s | 81920 | 1 | 8.7× |
| LightningDestructive | 209658 | 716.04 MB/s | 81920 | 1 | 8.7× |
| SonicFastest | 271285 | 553.38 MB/s | 244341 | 6 | 6.7× |
| Sonic | 271586 | 552.77 MB/s | 243266 | 6 | 6.7× |
| LightningDecodeAny | 485900 | 308.95 MB/s | 746004 | 10020 | 3.8× |
| Goccy | 856656 | 175.24 MB/s | 324701 | 10004 | 2.1× |
| JSONV2 | 1104489 | 135.92 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1824952 | 82.26 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33989 | 827.25 MB/s | 30272 | 105 | 9.0× |
| LightningDestructive | 34061 | 825.48 MB/s | 30144 | 103 | 8.9× |
| SonicFastest | 65390 | 429.99 MB/s | 48746 | 103 | 4.7× |
| Sonic | 65503 | 429.25 MB/s | 48918 | 103 | 4.6× |
| Easyjson | 68594 | 409.90 MB/s | 32304 | 138 | 4.4× |
| Goccy | 73189 | 384.17 MB/s | 59236 | 188 | 4.2× |
| JSONV2 | 135825 | 207.01 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 158612 | 177.27 MB/s | 135024 | 2678 | 1.9× |
| Stdlib | 304224 | 92.42 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1951 | 1193.31 MB/s | 32 | 1 | 11.7× |
| LightningDestructive | 2007 | 1160.07 MB/s | 32 | 1 | 11.4× |
| Goccy | 4214 | 552.41 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4233 | 549.98 MB/s | 192 | 2 | 5.4× |
| Sonic | 5119 | 454.75 MB/s | 4272 | 6 | 4.5× |
| SonicFastest | 5132 | 453.59 MB/s | 4282 | 6 | 4.4× |
| JSONV2 | 8503 | 273.77 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10345 | 162.87 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22817 | 102.03 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 862.69 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 221 | 854.85 MB/s | 0 | 0 | 11.0× |
| Goccy | 393 | 481.02 MB/s | 304 | 2 | 6.2× |
| Easyjson | 482 | 391.92 MB/s | 0 | 0 | 5.0× |
| Sonic | 803 | 235.38 MB/s | 519 | 4 | 3.0× |
| SonicFastest | 808 | 233.98 MB/s | 523 | 4 | 3.0× |
| JSONV2 | 1026 | 184.19 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1300 | 103.09 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2431 | 77.74 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1562 | 1402.48 MB/s | 0 | 0 | 10.3× |
| LightningDestructive | 1562 | 1402.66 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3188 | 687.37 MB/s | 24 | 1 | 5.0× |
| Goccy | 3196 | 685.54 MB/s | 2864 | 4 | 5.0× |
| Sonic | 6413 | 341.64 MB/s | 4007 | 40 | 2.5× |
| SonicFastest | 6413 | 341.64 MB/s | 3991 | 40 | 2.5× |
| JSONV2 | 8055 | 272.02 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8329 | 217.44 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16014 | 136.82 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 710278 | 718.70 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 746620 | 683.72 MB/s | 703778 | 1012 | 8.2× |
| Goccy | 1163022 | 438.92 MB/s | 1139231 | 5006 | 5.2× |
| Sonic | 1165905 | 437.84 MB/s | 898534 | 2006 | 5.2× |
| SonicFastest | 1166663 | 437.55 MB/s | 895325 | 2006 | 5.2× |
| Easyjson | 1568493 | 325.46 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3205973 | 159.23 MB/s | 1076008 | 12646 | 1.9× |
| LightningDecodeAny | 3581509 | 128.85 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6101893 | 83.66 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14642.57 MB/s | 0 | 0 | 99.5× |
| LightningDestructive | 1390 | 14236.13 MB/s | 0 | 0 | 96.7× |
| Goccy | 20117 | 983.71 MB/s | 20491 | 2 | 6.7× |
| Sonic | 27223 | 726.91 MB/s | 22213 | 4 | 4.9× |
| SonicFastest | 27572 | 717.72 MB/s | 22833 | 4 | 4.9× |
| JSONV2 | 29764 | 664.86 MB/s | 8 | 1 | 4.5× |
| LightningDecodeAny | 82224 | 240.66 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 93931 | 210.68 MB/s | 0 | 0 | 1.4× |
| Stdlib | 134389 | 147.25 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2650 | 6838.88 MB/s | 0 | 0 | 38.8× |
| Lightning | 2822 | 6421.46 MB/s | 432 | 2 | 36.4× |
| Easyjson | 3949 | 4589.22 MB/s | 432 | 2 | 26.0× |
| Sonic | 10113 | 1792.10 MB/s | 22883 | 6 | 10.2× |
| SonicFastest | 10434 | 1736.99 MB/s | 23739 | 6 | 9.8× |
| Goccy | 16427 | 1103.33 MB/s | 19459 | 2 | 6.3× |
| LightningDecodeAny | 17009 | 1051.35 MB/s | 29088 | 191 | 6.0× |
| JSONV2 | 46908 | 386.37 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102690 | 176.49 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2410828 | 833.11 MB/s | 2951717 | 7277 | 7.8× |
| Lightning | 2454078 | 818.43 MB/s | 2953429 | 7283 | 7.6× |
| Goccy | 4331455 | 463.70 MB/s | 5411494 | 15830 | 4.3× |
| SonicFastest | 4492479 | 447.08 MB/s | 10922889 | 13683 | 4.2× |
| Sonic | 4512928 | 445.05 MB/s | 10938189 | 13683 | 4.1× |
| Easyjson | 4954642 | 405.38 MB/s | 2981484 | 7439 | 3.8× |
| JSONV2 | 7057634 | 284.58 MB/s | 3173685 | 14563 | 2.6× |
| LightningDecodeAny | 7230151 | 157.99 MB/s | 7386072 | 134751 | 2.6× |
| Stdlib | 18696063 | 107.43 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 897 | 612.12 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 900 | 609.85 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1930 | 283.87 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2166 | 253.44 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2704 | 203.05 MB/s | 1968 | 26 | 2.1× |
| Sonic | 2705 | 202.96 MB/s | 1952 | 26 | 2.1× |
| Goccy | 3068 | 178.92 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3313 | 165.69 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5670 | 96.82 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 500267 | 1262.35 MB/s | 364472 | 566 | 10.8× |
| Lightning | 559784 | 1128.14 MB/s | 413001 | 878 | 9.7× |
| SonicFastest | 1024063 | 616.67 MB/s | 1008422 | 1102 | 5.3× |
| Sonic | 1025966 | 615.53 MB/s | 1007374 | 1102 | 5.3× |
| Easyjson | 1147323 | 550.42 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1183937 | 533.40 MB/s | 988215 | 1201 | 4.6× |
| JSONV2 | 2190369 | 288.31 MB/s | 571611 | 3144 | 2.5× |
| LightningDecodeAny | 2437569 | 191.55 MB/s | 2010201 | 30295 | 2.2× |
| Stdlib | 5405068 | 116.84 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 720630 | 780.44 MB/s | 544251 | 448 | 7.3× |
| Lightning | 901657 | 623.75 MB/s | 767621 | 1254 | 5.9× |
| SonicFastest | 1037643 | 542.01 MB/s | 955807 | 1476 | 5.1× |
| Sonic | 1037940 | 541.85 MB/s | 957487 | 1476 | 5.1× |
| Goccy | 1352267 | 415.90 MB/s | 1036817 | 1029 | 3.9× |
| Easyjson | 1760147 | 319.52 MB/s | 775154 | 1254 | 3.0× |
| JSONV2 | 2772310 | 202.87 MB/s | 927439 | 3482 | 1.9× |
| LightningDecodeAny | 2789546 | 201.61 MB/s | 2114151 | 30295 | 1.9× |
| Stdlib | 5284492 | 106.43 MB/s | 1011671 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 667798 | 798.41 MB/s | 333416 | 2084 | 8.2× |
| Lightning | 712230 | 748.60 MB/s | 368224 | 2293 | 7.7× |
| Easyjson | 1146480 | 465.06 MB/s | 428362 | 3273 | 4.8× |
| Sonic | 1171796 | 455.01 MB/s | 1033412 | 4351 | 4.7× |
| SonicFastest | 1172632 | 454.68 MB/s | 1043276 | 4351 | 4.7× |
| Goccy | 1306746 | 408.02 MB/s | 1167232 | 5409 | 4.2× |
| JSONV2 | 2567944 | 207.63 MB/s | 745469 | 13288 | 2.1× |
| LightningDecodeAny | 3506580 | 152.05 MB/s | 2674622 | 50596 | 1.6× |
| Stdlib | 5503982 | 96.87 MB/s | 798692 | 17133 | 1.0× |
