# JSON Deserialization Benchmarks

- generated 2026-06-23T19:53:21Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104128 | 1222.29 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 104552 | 1217.34 MB/s | 49280 | 2 | 10.6× |
| Sonic | 180777 | 704.05 MB/s | 193579 | 10 | 6.1× |
| SonicFastest | 182654 | 696.81 MB/s | 199030 | 10 | 6.0× |
| Goccy | 188305 | 675.90 MB/s | 224768 | 884 | 5.9× |
| Easyjson | 211312 | 602.31 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 416511 | 305.57 MB/s | 195118 | 1805 | 2.7× |
| LightningDecodeAny | 443389 | 213.48 MB/s | 465586 | 9714 | 2.5× |
| Stdlib | 1103956 | 115.29 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3738922 | 602.06 MB/s | 3122872 | 3081 | 7.0× |
| Lightning | 3815261 | 590.01 MB/s | 3122876 | 3081 | 6.8× |
| Sonic | 4445667 | 506.35 MB/s | 15236969 | 970 | 5.9× |
| SonicFastest | 4446262 | 506.28 MB/s | 15232102 | 970 | 5.9× |
| Goccy | 10237457 | 219.88 MB/s | 4111394 | 56532 | 2.5× |
| Easyjson | 11045491 | 203.80 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11566857 | 194.61 MB/s | 7938299 | 281383 | 2.3× |
| JSONV2 | 16110189 | 139.73 MB/s | 3123206 | 3083 | 1.6× |
| Stdlib | 26069192 | 86.35 MB/s | 3123399 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 491600 | 550.05 MB/s | 348024 | 1627 | 6.8× |
| LightningDestructive | 494121 | 547.24 MB/s | 348024 | 1627 | 6.8× |
| SonicFastest | 632597 | 427.45 MB/s | 476862 | 968 | 5.3× |
| Sonic | 634807 | 425.96 MB/s | 486151 | 968 | 5.3× |
| Goccy | 1400781 | 193.04 MB/s | 542658 | 8122 | 2.4× |
| Easyjson | 1402300 | 192.83 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1641321 | 164.75 MB/s | 1011486 | 37901 | 2.0× |
| JSONV2 | 2108727 | 128.23 MB/s | 348149 | 1628 | 1.6× |
| Stdlib | 3354014 | 80.62 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1341680 | 1287.34 MB/s | 959848 | 5881 | 9.8× |
| Lightning | 1358508 | 1271.40 MB/s | 959890 | 5882 | 9.7× |
| Sonic | 2046833 | 843.84 MB/s | 2735068 | 4020 | 6.4× |
| SonicFastest | 2057936 | 839.29 MB/s | 2726097 | 4020 | 6.4× |
| Goccy | 2373842 | 727.60 MB/s | 2584031 | 14605 | 5.6× |
| Easyjson | 4196063 | 411.62 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4300414 | 401.64 MB/s | 1011632 | 7594 | 3.1× |
| LightningDecodeAny | 4543841 | 110.10 MB/s | 4976205 | 81466 | 2.9× |
| Stdlib | 13199135 | 130.86 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1190 | 1522.99 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1208 | 1500.11 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2534 | 714.99 MB/s | 24 | 1 | 5.5× |
| Goccy | 2803 | 646.47 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5902 | 307.03 MB/s | 3746 | 40 | 2.4× |
| SonicFastest | 5922 | 305.98 MB/s | 3792 | 40 | 2.4× |
| JSONV2 | 7844 | 231.00 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8143 | 222.40 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14012 | 129.32 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1226 | 1477.96 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1236 | 1465.58 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2546 | 711.59 MB/s | 24 | 1 | 5.5× |
| Goccy | 2806 | 645.68 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 5982 | 302.90 MB/s | 3852 | 40 | 2.3× |
| Sonic | 5988 | 302.59 MB/s | 3872 | 40 | 2.3× |
| JSONV2 | 7780 | 232.90 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8411 | 215.30 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14056 | 128.91 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1385 | 1307.85 MB/s | 144 | 10 | 10.1× |
| LightningDestructive | 1442 | 1256.80 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2747 | 659.63 MB/s | 144 | 10 | 5.1× |
| Goccy | 2849 | 635.92 MB/s | 2600 | 5 | 4.9× |
| SonicFastest | 6215 | 291.53 MB/s | 3814 | 42 | 2.3× |
| Sonic | 6216 | 291.52 MB/s | 3783 | 42 | 2.3× |
| JSONV2 | 8002 | 226.43 MB/s | 632 | 7 | 1.7× |
| LightningDecodeAny | 8172 | 221.60 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 13995 | 129.48 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 738 | 669.54 MB/s | 160 | 1 | 7.5× |
| Lightning | 748 | 660.71 MB/s | 160 | 1 | 7.4× |
| Sonic | 1233 | 400.59 MB/s | 1012 | 6 | 4.5× |
| SonicFastest | 1234 | 400.37 MB/s | 1022 | 6 | 4.5× |
| LightningDecodeAny | 1676 | 294.17 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2194 | 225.20 MB/s | 448 | 3 | 2.5× |
| Goccy | 2455 | 201.23 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3198 | 154.46 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5528 | 89.37 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 468 | 490.93 MB/s | 160 | 1 | 8.8× |
| Lightning | 470 | 489.49 MB/s | 160 | 1 | 8.7× |
| Sonic | 868 | 264.81 MB/s | 660 | 6 | 4.7× |
| SonicFastest | 871 | 264.02 MB/s | 660 | 6 | 4.7× |
| LightningDecodeAny | 1361 | 168.31 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1370 | 167.84 MB/s | 448 | 3 | 3.0× |
| Goccy | 1592 | 144.43 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2327 | 98.82 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4111 | 55.95 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 71958 | 905.14 MB/s | 166212 | 102 | 7.6× |
| Lightning | 73204 | 889.74 MB/s | 172432 | 107 | 7.4× |
| SonicFastest | 96571 | 674.45 MB/s | 155409 | 75 | 5.6× |
| Sonic | 97939 | 665.03 MB/s | 156094 | 75 | 5.6× |
| Goccy | 143028 | 455.38 MB/s | 228562 | 134 | 3.8× |
| LightningDecodeAny | 182706 | 291.88 MB/s | 176960 | 3252 | 3.0× |
| JSONV2 | 224077 | 290.67 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 544782 | 119.56 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2575525 | 753.43 MB/s | 2846865 | 2238 | 9.1× |
| Lightning | 2653140 | 731.39 MB/s | 2846865 | 2238 | 8.8× |
| SonicFastest | 4550424 | 426.44 MB/s | 14608617 | 1407 | 5.2× |
| Goccy | 4734217 | 409.88 MB/s | 4065011 | 13510 | 5.0× |
| Sonic | 4787573 | 405.31 MB/s | 14608567 | 1407 | 4.9× |
| Easyjson | 7482347 | 259.34 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9364474 | 207.22 MB/s | 7013524 | 219937 | 2.5× |
| JSONV2 | 11237408 | 172.68 MB/s | 3237223 | 13947 | 2.1× |
| Stdlib | 23440670 | 82.78 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1096294 | 3035.53 MB/s | 351704 | 1286 | 19.0× |
| Lightning | 1782423 | 1867.03 MB/s | 2488905 | 2995 | 11.7× |
| Sonic | 2638059 | 1261.47 MB/s | 6404632 | 4248 | 7.9× |
| SonicFastest | 2638531 | 1261.24 MB/s | 6439141 | 4248 | 7.9× |
| LightningDecodeAny | 3674000 | 836.63 MB/s | 4886621 | 56892 | 5.7× |
| Goccy | 4480295 | 742.77 MB/s | 3948909 | 3816 | 4.6× |
| JSONV2 | 7331792 | 453.89 MB/s | 5364516 | 13243 | 2.8× |
| Stdlib | 20813476 | 159.89 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222992 | 988.13 MB/s | 136896 | 228 | 9.2× |
| LightningDestructive | 223112 | 987.60 MB/s | 136897 | 228 | 9.1× |
| Sonic | 384033 | 573.77 MB/s | 308281 | 398 | 5.3× |
| SonicFastest | 387985 | 567.92 MB/s | 318193 | 398 | 5.3× |
| Goccy | 427647 | 515.25 MB/s | 364103 | 1067 | 4.8× |
| Easyjson | 547469 | 402.48 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 724558 | 304.11 MB/s | 129740 | 470 | 2.8× |
| LightningDecodeAny | 866127 | 125.05 MB/s | 861393 | 11944 | 2.4× |
| Stdlib | 2040621 | 107.98 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12839624 | 630.86 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13241562 | 611.71 MB/s | 15059835 | 51649 | 6.7× |
| SonicFastest | 16760429 | 483.28 MB/s | 70887493 | 40014 | 5.3× |
| Sonic | 16786249 | 482.54 MB/s | 70887392 | 40014 | 5.3× |
| Goccy | 23621305 | 342.91 MB/s | 16998425 | 107148 | 3.8× |
| Easyjson | 30951974 | 261.70 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 40043137 | 129.94 MB/s | 34333352 | 912810 | 2.2× |
| JSONV2 | 43717934 | 185.28 MB/s | 15233743 | 78972 | 2.0× |
| Stdlib | 89290398 | 90.72 MB/s | 15665065 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6074716 | 491.13 MB/s | 3985336 | 30015 | 7.7× |
| Lightning | 6144556 | 485.55 MB/s | 3985339 | 30015 | 7.6× |
| Sonic | 8464973 | 352.45 MB/s | 26633168 | 56760 | 5.5× |
| SonicFastest | 8470038 | 352.24 MB/s | 26610392 | 56760 | 5.5× |
| Goccy | 16320277 | 182.81 MB/s | 10558859 | 273647 | 2.9× |
| Easyjson | 16590057 | 179.83 MB/s | 9479441 | 30115 | 2.8× |
| LightningDecodeAny | 18426691 | 99.54 MB/s | 20023837 | 408557 | 2.5× |
| JSONV2 | 24811559 | 120.25 MB/s | 9257173 | 86278 | 1.9× |
| Stdlib | 46963349 | 63.53 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1359694 | 532.18 MB/s | 907393 | 3618 | 8.7× |
| Lightning | 1374617 | 526.40 MB/s | 907388 | 3618 | 8.6× |
| SonicFastest | 1749722 | 413.55 MB/s | 3179968 | 7226 | 6.7× |
| Sonic | 1751109 | 413.22 MB/s | 3179213 | 7226 | 6.7× |
| Easyjson | 4260486 | 169.84 MB/s | 2847904 | 3698 | 2.8× |
| LightningDecodeAny | 4319967 | 150.60 MB/s | 5752203 | 80172 | 2.7× |
| Goccy | 4715428 | 153.45 MB/s | 2930940 | 80280 | 2.5× |
| JSONV2 | 5739212 | 126.08 MB/s | 2704657 | 7318 | 2.1× |
| Stdlib | 11805173 | 61.29 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1919083 | 821.93 MB/s | 907393 | 3618 | 8.2× |
| Lightning | 1970302 | 800.56 MB/s | 907386 | 3618 | 8.0× |
| Sonic | 2213734 | 712.53 MB/s | 5784496 | 7226 | 7.1× |
| SonicFastest | 2214889 | 712.16 MB/s | 5783779 | 7226 | 7.1× |
| LightningDecodeAny | 3801174 | 198.20 MB/s | 5752203 | 80172 | 4.1× |
| Easyjson | 5571674 | 283.10 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5599222 | 281.71 MB/s | 3565351 | 80266 | 2.8× |
| JSONV2 | 6441177 | 244.89 MB/s | 2704594 | 7318 | 2.4× |
| Stdlib | 15717505 | 100.36 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 223107 | 672.88 MB/s | 81920 | 1 | 8.2× |
| LightningDestructive | 224247 | 669.46 MB/s | 81920 | 1 | 8.1× |
| Sonic | 273673 | 548.55 MB/s | 253651 | 6 | 6.7× |
| SonicFastest | 274979 | 545.95 MB/s | 257504 | 6 | 6.6× |
| LightningDecodeAny | 467419 | 321.17 MB/s | 746005 | 10020 | 3.9× |
| Goccy | 854911 | 175.60 MB/s | 324806 | 10004 | 2.1× |
| JSONV2 | 1101463 | 136.30 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1820289 | 82.47 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33037 | 851.07 MB/s | 30272 | 105 | 9.2× |
| LightningDestructive | 33135 | 848.55 MB/s | 30144 | 103 | 9.2× |
| Sonic | 63153 | 445.22 MB/s | 46858 | 103 | 4.8× |
| SonicFastest | 63403 | 443.47 MB/s | 47169 | 103 | 4.8× |
| Easyjson | 67775 | 414.85 MB/s | 32304 | 138 | 4.5× |
| Goccy | 70798 | 397.14 MB/s | 59219 | 188 | 4.3× |
| JSONV2 | 133808 | 210.13 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 147856 | 190.16 MB/s | 135024 | 2678 | 2.1× |
| Stdlib | 303376 | 92.68 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1957 | 1189.38 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2007 | 1159.78 MB/s | 32 | 1 | 11.3× |
| Goccy | 4092 | 568.85 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4206 | 553.54 MB/s | 192 | 2 | 5.4× |
| Sonic | 5005 | 465.13 MB/s | 4242 | 6 | 4.5× |
| SonicFastest | 5038 | 462.13 MB/s | 4259 | 6 | 4.5× |
| JSONV2 | 8427 | 276.25 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10060 | 167.50 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22635 | 102.85 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219 | 862.27 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 221 | 855.10 MB/s | 0 | 0 | 11.0× |
| Goccy | 379 | 499.17 MB/s | 304 | 2 | 6.4× |
| Easyjson | 485 | 389.87 MB/s | 0 | 0 | 5.0× |
| Sonic | 789 | 239.67 MB/s | 496 | 4 | 3.1× |
| SonicFastest | 791 | 238.89 MB/s | 497 | 4 | 3.1× |
| JSONV2 | 1025 | 184.40 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1211 | 110.70 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2426 | 77.91 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1542 | 1421.15 MB/s | 0 | 0 | 10.4× |
| LightningDestructive | 1559 | 1405.17 MB/s | 0 | 0 | 10.3× |
| Easyjson | 3197 | 685.41 MB/s | 24 | 1 | 5.0× |
| Goccy | 3267 | 670.74 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6327 | 346.31 MB/s | 3957 | 40 | 2.5× |
| Sonic | 6349 | 345.12 MB/s | 3989 | 40 | 2.5× |
| JSONV2 | 8083 | 271.06 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8147 | 222.29 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 16022 | 136.75 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 696338 | 733.09 MB/s | 703778 | 1012 | 8.6× |
| Lightning | 722504 | 706.54 MB/s | 703780 | 1012 | 8.3× |
| Goccy | 1129660 | 451.88 MB/s | 1136518 | 5006 | 5.3× |
| Sonic | 1143630 | 446.36 MB/s | 889929 | 2006 | 5.3× |
| SonicFastest | 1143684 | 446.34 MB/s | 882375 | 2006 | 5.3× |
| Easyjson | 1518159 | 336.25 MB/s | 863778 | 3012 | 4.0× |
| JSONV2 | 3213073 | 158.87 MB/s | 1076010 | 12646 | 1.9× |
| LightningDecodeAny | 3440339 | 134.13 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6008985 | 84.95 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1352 | 14642.09 MB/s | 0 | 0 | 89.9× |
| LightningDestructive | 1388 | 14258.12 MB/s | 0 | 0 | 87.5× |
| Goccy | 19841 | 997.39 MB/s | 20491 | 2 | 6.1× |
| SonicFastest | 26848 | 737.08 MB/s | 22166 | 4 | 4.5× |
| Sonic | 26918 | 735.16 MB/s | 22296 | 4 | 4.5× |
| JSONV2 | 29762 | 664.90 MB/s | 8 | 1 | 4.1× |
| LightningDecodeAny | 77606 | 254.98 MB/s | 117104 | 2019 | 1.6× |
| Easyjson | 86041 | 230.00 MB/s | 0 | 0 | 1.4× |
| Stdlib | 121490 | 162.89 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2654 | 6827.92 MB/s | 0 | 0 | 38.7× |
| Lightning | 2813 | 6442.16 MB/s | 432 | 2 | 36.5× |
| Easyjson | 3937 | 4603.18 MB/s | 432 | 2 | 26.1× |
| SonicFastest | 9931 | 1824.92 MB/s | 23109 | 6 | 10.3× |
| Sonic | 10040 | 1805.20 MB/s | 23641 | 6 | 10.2× |
| Goccy | 16080 | 1127.11 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 16431 | 1088.31 MB/s | 29088 | 191 | 6.2× |
| JSONV2 | 46399 | 390.61 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102626 | 176.60 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2198137 | 913.73 MB/s | 2960388 | 7411 | 8.5× |
| Lightning | 2268004 | 885.58 MB/s | 2962101 | 7417 | 8.2× |
| Goccy | 4301660 | 466.91 MB/s | 5411587 | 15830 | 4.3× |
| Sonic | 4502067 | 446.13 MB/s | 10902535 | 13683 | 4.1× |
| SonicFastest | 4571652 | 439.34 MB/s | 10861389 | 13683 | 4.1× |
| Easyjson | 4884754 | 411.18 MB/s | 2981480 | 7438 | 3.8× |
| JSONV2 | 7017789 | 286.20 MB/s | 3173685 | 14563 | 2.7× |
| LightningDecodeAny | 7105233 | 160.77 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18637137 | 107.77 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 905 | 606.59 MB/s | 480 | 1 | 6.3× |
| Lightning | 905 | 606.39 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1854 | 295.58 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 2178 | 252.01 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2660 | 206.37 MB/s | 1938 | 26 | 2.1× |
| SonicFastest | 2672 | 205.49 MB/s | 1949 | 26 | 2.1× |
| Goccy | 3022 | 181.64 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3337 | 164.54 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5667 | 96.88 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 496342 | 1272.34 MB/s | 364472 | 566 | 10.8× |
| Lightning | 547747 | 1152.93 MB/s | 413001 | 878 | 9.8× |
| SonicFastest | 1021289 | 618.35 MB/s | 996672 | 1102 | 5.3× |
| Sonic | 1023316 | 617.12 MB/s | 1000031 | 1102 | 5.3× |
| Easyjson | 1138988 | 554.45 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1161289 | 543.80 MB/s | 984884 | 1201 | 4.6× |
| JSONV2 | 2139728 | 295.14 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2343114 | 199.27 MB/s | 2010201 | 30295 | 2.3× |
| Stdlib | 5379425 | 117.39 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 711720 | 790.21 MB/s | 544252 | 448 | 7.4× |
| Lightning | 877094 | 641.22 MB/s | 767621 | 1254 | 6.0× |
| Sonic | 1036443 | 542.63 MB/s | 982459 | 1476 | 5.1× |
| SonicFastest | 1043399 | 539.02 MB/s | 984348 | 1476 | 5.0× |
| Goccy | 1354598 | 415.18 MB/s | 1042569 | 1030 | 3.9× |
| Easyjson | 1738287 | 323.54 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2710668 | 207.48 MB/s | 2114151 | 30295 | 1.9× |
| JSONV2 | 2755081 | 204.13 MB/s | 927442 | 3482 | 1.9× |
| Stdlib | 5246988 | 107.19 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 645279 | 826.27 MB/s | 333416 | 2084 | 8.5× |
| Lightning | 665105 | 801.64 MB/s | 368224 | 2293 | 8.2× |
| Easyjson | 1100323 | 484.57 MB/s | 428361 | 3273 | 5.0× |
| Sonic | 1137004 | 468.93 MB/s | 1041808 | 4351 | 4.8× |
| SonicFastest | 1139066 | 468.08 MB/s | 1048523 | 4351 | 4.8× |
| Goccy | 1272089 | 419.14 MB/s | 1167246 | 5409 | 4.3× |
| JSONV2 | 2524551 | 211.20 MB/s | 745447 | 13288 | 2.2× |
| LightningDecodeAny | 3300363 | 161.55 MB/s | 2674619 | 50596 | 1.7× |
| Stdlib | 5461153 | 97.63 MB/s | 798692 | 17133 | 1.0× |
