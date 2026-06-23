# JSON Deserialization Benchmarks

- generated 2026-06-23T14:28:18Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 104513 | 1217.79 MB/s | 49280 | 2 | 10.6× |
| Lightning | 104518 | 1217.73 MB/s | 49760 | 3 | 10.6× |
| SonicFastest | 181386 | 701.68 MB/s | 196039 | 10 | 6.1× |
| Sonic | 181927 | 699.60 MB/s | 194431 | 10 | 6.1× |
| Goccy | 188857 | 673.92 MB/s | 224579 | 884 | 5.8× |
| Easyjson | 211852 | 600.77 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 417590 | 304.78 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 446899 | 211.80 MB/s | 465585 | 9714 | 2.5× |
| Stdlib | 1104717 | 115.21 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3741650 | 601.62 MB/s | 3122872 | 3081 | 7.0× |
| Lightning | 3825996 | 588.36 MB/s | 3122874 | 3081 | 6.8× |
| Sonic | 4517551 | 498.29 MB/s | 15233778 | 970 | 5.8× |
| SonicFastest | 4524771 | 497.49 MB/s | 15235426 | 970 | 5.8× |
| Goccy | 10384588 | 216.77 MB/s | 4121626 | 56532 | 2.5× |
| Easyjson | 11062880 | 203.48 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11628385 | 193.58 MB/s | 7938298 | 281383 | 2.2× |
| JSONV2 | 16298273 | 138.12 MB/s | 3123230 | 3083 | 1.6× |
| Stdlib | 26110714 | 86.21 MB/s | 3123395 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 493668 | 547.74 MB/s | 348024 | 1627 | 6.8× |
| LightningDestructive | 495844 | 545.34 MB/s | 348024 | 1627 | 6.8× |
| SonicFastest | 638874 | 423.25 MB/s | 493600 | 968 | 5.3× |
| Sonic | 639347 | 422.94 MB/s | 495592 | 968 | 5.2× |
| Easyjson | 1406390 | 192.27 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1417679 | 190.74 MB/s | 543431 | 8122 | 2.4× |
| LightningDecodeAny | 1674004 | 161.53 MB/s | 1011486 | 37901 | 2.0× |
| JSONV2 | 2125566 | 127.21 MB/s | 348153 | 1628 | 1.6× |
| Stdlib | 3355928 | 80.57 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1364049 | 1266.23 MB/s | 959848 | 5881 | 9.7× |
| Lightning | 1385976 | 1246.20 MB/s | 959890 | 5882 | 9.5× |
| Sonic | 2048880 | 843.00 MB/s | 2690232 | 4020 | 6.5× |
| SonicFastest | 2051469 | 841.94 MB/s | 2734573 | 4020 | 6.4× |
| Goccy | 2362523 | 731.08 MB/s | 2582926 | 14605 | 5.6× |
| Easyjson | 4224339 | 408.87 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4283355 | 403.24 MB/s | 1011643 | 7594 | 3.1× |
| LightningDecodeAny | 4666233 | 107.22 MB/s | 4976204 | 81466 | 2.8× |
| Stdlib | 13230319 | 130.55 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1190 | 1522.13 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1206 | 1502.94 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2537 | 714.11 MB/s | 24 | 1 | 5.6× |
| Goccy | 2863 | 633.01 MB/s | 2608 | 4 | 4.9× |
| Sonic | 5911 | 306.57 MB/s | 3724 | 40 | 2.4× |
| SonicFastest | 5966 | 303.74 MB/s | 3764 | 40 | 2.4× |
| JSONV2 | 7902 | 229.30 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8223 | 220.23 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14090 | 128.60 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1222 | 1483.40 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1238 | 1463.44 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2542 | 712.88 MB/s | 24 | 1 | 5.5× |
| Goccy | 2766 | 655.02 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5951 | 304.48 MB/s | 3751 | 40 | 2.4× |
| Sonic | 5954 | 304.34 MB/s | 3729 | 40 | 2.4× |
| JSONV2 | 7837 | 231.21 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8204 | 220.74 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14054 | 128.93 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1388 | 1305.92 MB/s | 144 | 10 | 10.2× |
| LightningDestructive | 1440 | 1258.73 MB/s | 144 | 10 | 9.8× |
| Easyjson | 2747 | 659.70 MB/s | 144 | 10 | 5.1× |
| Goccy | 2841 | 637.71 MB/s | 2600 | 5 | 5.0× |
| Sonic | 6242 | 290.31 MB/s | 3765 | 42 | 2.3× |
| SonicFastest | 6265 | 289.22 MB/s | 3805 | 42 | 2.3× |
| JSONV2 | 8041 | 225.36 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8178 | 221.45 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14147 | 128.08 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 740 | 667.16 MB/s | 160 | 1 | 7.5× |
| Lightning | 743 | 664.55 MB/s | 160 | 1 | 7.4× |
| Sonic | 1231 | 401.25 MB/s | 977 | 6 | 4.5× |
| SonicFastest | 1233 | 400.68 MB/s | 986 | 6 | 4.5× |
| LightningDecodeAny | 1640 | 300.55 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2194 | 225.19 MB/s | 448 | 3 | 2.5× |
| Goccy | 2438 | 202.66 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3183 | 155.18 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5521 | 89.47 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 471 | 488.22 MB/s | 160 | 1 | 8.8× |
| Lightning | 475 | 483.86 MB/s | 160 | 1 | 8.7× |
| Sonic | 879 | 261.71 MB/s | 652 | 6 | 4.7× |
| SonicFastest | 881 | 261.09 MB/s | 656 | 6 | 4.7× |
| LightningDecodeAny | 1377 | 166.31 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1380 | 166.67 MB/s | 448 | 3 | 3.0× |
| Goccy | 1592 | 144.50 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2342 | 98.20 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4134 | 55.63 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 73359 | 887.85 MB/s | 166212 | 102 | 7.4× |
| Lightning | 74887 | 869.73 MB/s | 172432 | 107 | 7.3× |
| Sonic | 96517 | 674.83 MB/s | 155277 | 75 | 5.6× |
| SonicFastest | 96705 | 673.51 MB/s | 155635 | 75 | 5.6× |
| Goccy | 143861 | 452.74 MB/s | 229210 | 134 | 3.8× |
| LightningDecodeAny | 186375 | 286.14 MB/s | 176960 | 3252 | 2.9× |
| JSONV2 | 223635 | 291.24 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 544619 | 119.59 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2608197 | 743.99 MB/s | 2846864 | 2238 | 9.0× |
| Lightning | 2669439 | 726.92 MB/s | 2846867 | 2238 | 8.8× |
| Goccy | 4724848 | 410.70 MB/s | 4066346 | 13510 | 5.0× |
| SonicFastest | 4988913 | 388.96 MB/s | 14608710 | 1407 | 4.7× |
| Sonic | 5015044 | 386.93 MB/s | 14608667 | 1407 | 4.7× |
| Easyjson | 7598446 | 255.38 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 9703107 | 199.98 MB/s | 7013526 | 219937 | 2.4× |
| JSONV2 | 11297499 | 171.76 MB/s | 3237229 | 13947 | 2.1× |
| Stdlib | 23470958 | 82.68 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1103212 | 3016.49 MB/s | 351704 | 1286 | 18.9× |
| Lightning | 1836561 | 1811.99 MB/s | 2488904 | 2995 | 11.3× |
| SonicFastest | 2685455 | 1239.21 MB/s | 6429509 | 4248 | 7.8× |
| Sonic | 2720555 | 1223.22 MB/s | 6451746 | 4248 | 7.7× |
| LightningDecodeAny | 3732278 | 823.56 MB/s | 4886621 | 56892 | 5.6× |
| Goccy | 4539122 | 733.14 MB/s | 3948909 | 3816 | 4.6× |
| JSONV2 | 7487312 | 444.46 MB/s | 5364520 | 13243 | 2.8× |
| Stdlib | 20815312 | 159.87 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 223031 | 987.96 MB/s | 136896 | 228 | 9.1× |
| LightningDestructive | 225291 | 978.05 MB/s | 136897 | 228 | 9.1× |
| Sonic | 381028 | 578.29 MB/s | 304415 | 398 | 5.4× |
| SonicFastest | 382001 | 576.82 MB/s | 302821 | 398 | 5.3× |
| Goccy | 438627 | 502.35 MB/s | 364793 | 1067 | 4.6× |
| Easyjson | 547599 | 402.39 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 726699 | 303.22 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 883754 | 122.56 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2039237 | 108.05 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12856376 | 630.04 MB/s | 15059832 | 51649 | 6.9× |
| Lightning | 13244321 | 611.59 MB/s | 15059842 | 51649 | 6.7× |
| Sonic | 16895637 | 479.42 MB/s | 70887273 | 40014 | 5.3× |
| SonicFastest | 16901106 | 479.26 MB/s | 70873027 | 40014 | 5.3× |
| Goccy | 23462212 | 345.24 MB/s | 17033540 | 107148 | 3.8× |
| Easyjson | 31009036 | 261.22 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 40568349 | 128.25 MB/s | 34333352 | 912810 | 2.2× |
| JSONV2 | 43705348 | 185.33 MB/s | 15233762 | 78972 | 2.0× |
| Stdlib | 89169259 | 90.84 MB/s | 15665066 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6102650 | 488.88 MB/s | 3985336 | 30015 | 7.7× |
| Lightning | 6191992 | 481.83 MB/s | 3985339 | 30015 | 7.6× |
| Sonic | 8585389 | 347.51 MB/s | 26663410 | 56760 | 5.5× |
| SonicFastest | 8611993 | 346.43 MB/s | 26651285 | 56760 | 5.5× |
| Goccy | 16426443 | 181.63 MB/s | 10625169 | 273649 | 2.9× |
| Easyjson | 16637576 | 179.32 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 18556647 | 98.84 MB/s | 20023836 | 408557 | 2.5× |
| JSONV2 | 24650059 | 121.03 MB/s | 9257163 | 86278 | 1.9× |
| Stdlib | 47073897 | 63.38 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1374020 | 526.63 MB/s | 907393 | 3618 | 8.4× |
| Lightning | 1390220 | 520.49 MB/s | 907388 | 3618 | 8.3× |
| Sonic | 1802136 | 401.52 MB/s | 3186997 | 7226 | 6.4× |
| SonicFastest | 1805660 | 400.74 MB/s | 3188040 | 7226 | 6.4× |
| Easyjson | 4247257 | 170.37 MB/s | 2847906 | 3698 | 2.7× |
| LightningDecodeAny | 4401877 | 147.79 MB/s | 5752201 | 80172 | 2.6× |
| Goccy | 4746634 | 152.44 MB/s | 2849787 | 80275 | 2.4× |
| JSONV2 | 5908515 | 122.47 MB/s | 2704610 | 7318 | 2.0× |
| Stdlib | 11573003 | 62.52 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1924855 | 819.47 MB/s | 907393 | 3618 | 8.2× |
| Lightning | 1976449 | 798.07 MB/s | 907387 | 3618 | 8.0× |
| Sonic | 2290206 | 688.74 MB/s | 5793617 | 7226 | 6.9× |
| SonicFastest | 2291002 | 688.50 MB/s | 5782438 | 7226 | 6.9× |
| LightningDecodeAny | 3887568 | 193.80 MB/s | 5752199 | 80172 | 4.0× |
| Easyjson | 5593529 | 282.00 MB/s | 2847907 | 3698 | 2.8× |
| Goccy | 5818567 | 271.09 MB/s | 3631760 | 80269 | 2.7× |
| JSONV2 | 6595729 | 239.15 MB/s | 2704591 | 7318 | 2.4× |
| Stdlib | 15742065 | 100.20 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 223432 | 671.90 MB/s | 81920 | 1 | 8.2× |
| LightningDestructive | 224931 | 667.42 MB/s | 81920 | 1 | 8.1× |
| SonicFastest | 272246 | 551.43 MB/s | 249836 | 6 | 6.7× |
| Sonic | 273822 | 548.25 MB/s | 254713 | 6 | 6.7× |
| LightningDecodeAny | 478794 | 313.54 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 855890 | 175.40 MB/s | 325073 | 10004 | 2.1× |
| JSONV2 | 1104880 | 135.87 MB/s | 357715 | 20 | 1.6× |
| Stdlib | 1822323 | 82.38 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33058 | 850.52 MB/s | 30272 | 105 | 9.2× |
| LightningDestructive | 33360 | 842.83 MB/s | 30144 | 103 | 9.1× |
| Sonic | 63466 | 443.03 MB/s | 47038 | 103 | 4.8× |
| SonicFastest | 63742 | 441.11 MB/s | 47181 | 103 | 4.8× |
| Easyjson | 67921 | 413.97 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71381 | 393.90 MB/s | 59215 | 188 | 4.2× |
| JSONV2 | 134382 | 209.23 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 149916 | 187.55 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303207 | 92.73 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1958 | 1188.80 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2007 | 1159.94 MB/s | 32 | 1 | 11.3× |
| Goccy | 4088 | 569.45 MB/s | 3649 | 4 | 5.6× |
| Easyjson | 4210 | 553.02 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5057 | 460.37 MB/s | 4225 | 6 | 4.5× |
| Sonic | 5068 | 459.36 MB/s | 4251 | 6 | 4.5× |
| JSONV2 | 8446 | 275.65 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10242 | 164.52 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22768 | 102.25 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 862.07 MB/s | 0 | 0 | 11.2× |
| LightningDestructive | 221 | 853.84 MB/s | 0 | 0 | 11.0× |
| Goccy | 377 | 501.73 MB/s | 304 | 2 | 6.5× |
| Easyjson | 487 | 388.44 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 795 | 237.78 MB/s | 499 | 4 | 3.1× |
| Sonic | 797 | 237.22 MB/s | 504 | 4 | 3.1× |
| JSONV2 | 1029 | 183.61 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1229 | 109.06 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2445 | 77.29 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1542 | 1420.77 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1560 | 1404.50 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3197 | 685.41 MB/s | 24 | 1 | 5.0× |
| Goccy | 3249 | 674.45 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6340 | 345.60 MB/s | 3960 | 40 | 2.5× |
| Sonic | 6394 | 342.66 MB/s | 4011 | 40 | 2.5× |
| JSONV2 | 8024 | 273.05 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8206 | 220.68 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 16019 | 136.77 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 699895 | 729.36 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 729082 | 700.16 MB/s | 703779 | 1012 | 8.3× |
| Goccy | 1144100 | 446.18 MB/s | 1134772 | 5006 | 5.3× |
| Sonic | 1155525 | 441.77 MB/s | 889720 | 2006 | 5.2× |
| SonicFastest | 1160836 | 439.75 MB/s | 892239 | 2006 | 5.2× |
| Easyjson | 1525641 | 334.60 MB/s | 863777 | 3012 | 4.0× |
| JSONV2 | 3217368 | 158.66 MB/s | 1076016 | 12646 | 1.9× |
| LightningDecodeAny | 3512270 | 131.39 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6039747 | 84.52 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14647.88 MB/s | 0 | 0 | 82.5× |
| LightningDestructive | 1392 | 14220.54 MB/s | 0 | 0 | 80.1× |
| Goccy | 20665 | 957.62 MB/s | 20491 | 2 | 5.4× |
| Sonic | 27194 | 727.71 MB/s | 22403 | 4 | 4.1× |
| SonicFastest | 27341 | 723.79 MB/s | 22778 | 4 | 4.1× |
| JSONV2 | 29648 | 667.46 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 78425 | 252.32 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86799 | 227.99 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111506 | 177.47 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2684 | 6753.63 MB/s | 0 | 0 | 38.2× |
| Lightning | 2821 | 6423.75 MB/s | 432 | 2 | 36.4× |
| Easyjson | 3945 | 4593.69 MB/s | 432 | 2 | 26.0× |
| Sonic | 9915 | 1827.93 MB/s | 22938 | 6 | 10.4× |
| SonicFastest | 10106 | 1793.42 MB/s | 23547 | 6 | 10.2× |
| Goccy | 15907 | 1139.39 MB/s | 19459 | 2 | 6.5× |
| LightningDecodeAny | 16864 | 1060.36 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 46045 | 393.62 MB/s | 16500 | 50 | 2.2× |
| Stdlib | 102653 | 176.56 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2225550 | 902.47 MB/s | 2960388 | 7411 | 8.4× |
| Lightning | 2303744 | 871.84 MB/s | 2962102 | 7417 | 8.1× |
| Goccy | 4237681 | 473.96 MB/s | 5412640 | 15832 | 4.4× |
| SonicFastest | 4413459 | 455.08 MB/s | 10943350 | 13683 | 4.2× |
| Sonic | 4428521 | 453.54 MB/s | 10941832 | 13683 | 4.2× |
| Easyjson | 4899352 | 409.95 MB/s | 2981486 | 7439 | 3.8× |
| JSONV2 | 6936978 | 289.53 MB/s | 3173685 | 14563 | 2.7× |
| LightningDecodeAny | 7179472 | 159.11 MB/s | 7386072 | 134751 | 2.6× |
| Stdlib | 18681516 | 107.51 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 887 | 619.24 MB/s | 480 | 1 | 6.4× |
| LightningDestructive | 888 | 618.06 MB/s | 480 | 1 | 6.4× |
| LightningDecodeAny | 1908 | 287.24 MB/s | 2261 | 50 | 3.0× |
| Easyjson | 2158 | 254.45 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2680 | 204.82 MB/s | 1937 | 26 | 2.1× |
| SonicFastest | 2685 | 204.46 MB/s | 1951 | 26 | 2.1× |
| Goccy | 3071 | 178.79 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3323 | 165.23 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5650 | 97.16 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 493134 | 1280.61 MB/s | 364472 | 566 | 10.9× |
| Lightning | 546789 | 1154.95 MB/s | 413001 | 878 | 9.9× |
| SonicFastest | 1001614 | 630.50 MB/s | 991636 | 1102 | 5.4× |
| Sonic | 1003007 | 629.62 MB/s | 992895 | 1102 | 5.4× |
| Easyjson | 1141876 | 553.05 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1165390 | 541.89 MB/s | 988153 | 1201 | 4.6× |
| JSONV2 | 2147107 | 294.12 MB/s | 571616 | 3144 | 2.5× |
| LightningDecodeAny | 2379227 | 196.24 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5391958 | 117.12 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 710292 | 791.80 MB/s | 544253 | 448 | 7.4× |
| Lightning | 894387 | 628.82 MB/s | 767622 | 1254 | 5.9× |
| SonicFastest | 1024916 | 548.74 MB/s | 945107 | 1476 | 5.1× |
| Sonic | 1052793 | 534.21 MB/s | 973437 | 1476 | 5.0× |
| Goccy | 1347158 | 417.48 MB/s | 1041195 | 1030 | 3.9× |
| Easyjson | 1748374 | 321.67 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2735764 | 205.58 MB/s | 2114151 | 30295 | 1.9× |
| JSONV2 | 2763101 | 203.54 MB/s | 927448 | 3482 | 1.9× |
| Stdlib | 5265397 | 106.81 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 652302 | 817.38 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 672748 | 792.54 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1107604 | 481.38 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1146858 | 464.90 MB/s | 1041806 | 4351 | 4.8× |
| SonicFastest | 1155154 | 461.56 MB/s | 1042858 | 4351 | 4.7× |
| Goccy | 1289440 | 413.50 MB/s | 1167228 | 5409 | 4.2× |
| JSONV2 | 2526563 | 211.03 MB/s | 745448 | 13288 | 2.2× |
| LightningDecodeAny | 3360940 | 158.64 MB/s | 2674619 | 50596 | 1.6× |
| Stdlib | 5477496 | 97.34 MB/s | 798692 | 17133 | 1.0× |
