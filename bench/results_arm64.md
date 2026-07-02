# JSON Deserialization Benchmarks

- generated 2026-07-01T22:11:41Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104395 | 1219.17 MB/s | 49760 | 3 | 10.6× |
| LightningDestructive | 104427 | 1218.79 MB/s | 49280 | 2 | 10.6× |
| SonicFastest | 183681 | 692.91 MB/s | 198092 | 10 | 6.0× |
| Sonic | 184839 | 688.57 MB/s | 200744 | 10 | 6.0× |
| Goccy | 192389 | 661.55 MB/s | 224927 | 884 | 5.7× |
| Easyjson | 212345 | 599.38 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 419226 | 303.59 MB/s | 195120 | 1805 | 2.6× |
| LightningDecodeAny | 457929 | 206.70 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1103148 | 115.37 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3691933 | 609.72 MB/s | 3122874 | 3081 | 7.1× |
| Lightning | 3743302 | 601.35 MB/s | 3122874 | 3081 | 7.0× |
| SonicFastest | 4512320 | 498.87 MB/s | 15233776 | 970 | 5.8× |
| Sonic | 4538463 | 495.99 MB/s | 15232101 | 970 | 5.8× |
| Goccy | 10161748 | 221.52 MB/s | 4109321 | 56532 | 2.6× |
| Easyjson | 10984153 | 204.94 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11201952 | 200.95 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 16023341 | 140.49 MB/s | 3123229 | 3083 | 1.6× |
| Stdlib | 26098195 | 86.25 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 483817 | 558.90 MB/s | 348024 | 1627 | 6.9× |
| LightningDestructive | 491115 | 550.59 MB/s | 348024 | 1627 | 6.8× |
| Sonic | 628055 | 430.54 MB/s | 469609 | 968 | 5.3× |
| SonicFastest | 632847 | 427.28 MB/s | 480210 | 968 | 5.3× |
| Easyjson | 1402504 | 192.80 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1409377 | 191.86 MB/s | 543632 | 8122 | 2.4× |
| LightningDecodeAny | 1582419 | 170.88 MB/s | 1011484 | 37901 | 2.1× |
| JSONV2 | 2095677 | 129.03 MB/s | 348153 | 1628 | 1.6× |
| Stdlib | 3350555 | 80.70 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1370321 | 1260.44 MB/s | 959848 | 5881 | 9.7× |
| Lightning | 1380813 | 1250.86 MB/s | 959890 | 5882 | 9.6× |
| Sonic | 2041783 | 845.93 MB/s | 2724423 | 4020 | 6.5× |
| SonicFastest | 2049752 | 842.64 MB/s | 2744297 | 4020 | 6.5× |
| Goccy | 2366610 | 729.82 MB/s | 2585113 | 14605 | 5.6× |
| Easyjson | 4219175 | 409.37 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4328895 | 398.99 MB/s | 1011641 | 7594 | 3.1× |
| LightningDecodeAny | 4596051 | 108.85 MB/s | 4976204 | 81466 | 2.9× |
| Stdlib | 13227995 | 130.57 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1193 | 1518.71 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1207 | 1501.15 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2525 | 717.59 MB/s | 24 | 1 | 5.6× |
| Goccy | 2851 | 635.67 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6056 | 299.20 MB/s | 3723 | 40 | 2.3× |
| SonicFastest | 6118 | 296.19 MB/s | 3795 | 40 | 2.3× |
| JSONV2 | 7757 | 233.61 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8149 | 222.24 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14025 | 129.20 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1226 | 1477.88 MB/s | 0 | 0 | 11.5× |
| LightningDestructive | 1235 | 1466.87 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2521 | 718.64 MB/s | 24 | 1 | 5.6× |
| Goccy | 2844 | 637.16 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6026 | 300.67 MB/s | 3764 | 40 | 2.3× |
| SonicFastest | 6053 | 299.38 MB/s | 3773 | 40 | 2.3× |
| JSONV2 | 7818 | 231.77 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8173 | 221.58 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14056 | 128.91 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1395 | 1299.09 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1449 | 1250.32 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2761 | 656.21 MB/s | 144 | 10 | 5.1× |
| Goccy | 2910 | 622.62 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6095 | 297.28 MB/s | 3794 | 42 | 2.3× |
| Sonic | 6105 | 296.81 MB/s | 3779 | 42 | 2.3× |
| JSONV2 | 7917 | 228.88 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8150 | 222.20 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14018 | 129.27 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 695 | 710.73 MB/s | 160 | 1 | 7.9× |
| LightningDestructive | 695 | 710.38 MB/s | 160 | 1 | 7.9× |
| Sonic | 1248 | 395.93 MB/s | 1026 | 6 | 4.4× |
| SonicFastest | 1248 | 395.91 MB/s | 1030 | 6 | 4.4× |
| LightningDecodeAny | 1671 | 295.02 MB/s | 1536 | 30 | 3.3× |
| Easyjson | 2254 | 219.18 MB/s | 448 | 3 | 2.4× |
| Goccy | 2443 | 202.18 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3228 | 153.04 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5518 | 89.53 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 424 | 542.63 MB/s | 160 | 1 | 9.7× |
| Lightning | 425 | 541.25 MB/s | 160 | 1 | 9.7× |
| SonicFastest | 891 | 258.05 MB/s | 692 | 6 | 4.6× |
| Sonic | 894 | 257.43 MB/s | 685 | 6 | 4.6× |
| LightningDecodeAny | 1418 | 161.52 MB/s | 1536 | 30 | 2.9× |
| Easyjson | 1431 | 160.68 MB/s | 448 | 3 | 2.9× |
| Goccy | 1599 | 143.81 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2431 | 94.61 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4117 | 55.87 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 71693 | 908.49 MB/s | 166212 | 102 | 7.6× |
| Lightning | 72593 | 897.22 MB/s | 172432 | 107 | 7.5× |
| Sonic | 96637 | 673.98 MB/s | 154626 | 75 | 5.6× |
| SonicFastest | 97042 | 671.17 MB/s | 154985 | 75 | 5.6× |
| Goccy | 142991 | 455.50 MB/s | 228894 | 134 | 3.8× |
| LightningDecodeAny | 183719 | 290.28 MB/s | 176960 | 3252 | 3.0× |
| JSONV2 | 221860 | 293.57 MB/s | 206652 | 607 | 2.5× |
| Stdlib | 544573 | 119.60 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2596308 | 747.40 MB/s | 2846864 | 2238 | 9.0× |
| Lightning | 2705088 | 717.34 MB/s | 2846867 | 2238 | 8.7× |
| Goccy | 4778942 | 406.05 MB/s | 4066523 | 13510 | 4.9× |
| Sonic | 4807951 | 403.60 MB/s | 14608610 | 1407 | 4.9× |
| SonicFastest | 4868844 | 398.55 MB/s | 14608675 | 1407 | 4.8× |
| Easyjson | 7452910 | 260.36 MB/s | 3871266 | 15043 | 3.2× |
| LightningDecodeAny | 9465113 | 205.01 MB/s | 7013524 | 219937 | 2.5× |
| JSONV2 | 11304943 | 171.65 MB/s | 3237214 | 13947 | 2.1× |
| Stdlib | 23476909 | 82.65 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1101259 | 3021.84 MB/s | 351704 | 1286 | 18.9× |
| Lightning | 1793895 | 1855.09 MB/s | 2488905 | 2995 | 11.6× |
| SonicFastest | 2657144 | 1252.41 MB/s | 6495271 | 4248 | 7.8× |
| Sonic | 2663090 | 1249.61 MB/s | 6498004 | 4248 | 7.8× |
| LightningDecodeAny | 3701609 | 830.39 MB/s | 4886621 | 56892 | 5.6× |
| Goccy | 4577678 | 726.97 MB/s | 3948909 | 3816 | 4.6× |
| JSONV2 | 7390598 | 450.28 MB/s | 5364516 | 13243 | 2.8× |
| Stdlib | 20832869 | 159.74 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 219942 | 1001.84 MB/s | 136896 | 228 | 9.3× |
| LightningDestructive | 221748 | 993.68 MB/s | 136897 | 228 | 9.2× |
| Sonic | 376230 | 585.67 MB/s | 296585 | 398 | 5.4× |
| SonicFastest | 380462 | 579.15 MB/s | 302551 | 398 | 5.4× |
| Goccy | 434702 | 506.89 MB/s | 364935 | 1067 | 4.7× |
| Easyjson | 548989 | 401.37 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 738203 | 298.49 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 873918 | 123.94 MB/s | 861393 | 11944 | 2.3× |
| Stdlib | 2038983 | 108.07 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12749797 | 635.31 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13247305 | 611.45 MB/s | 15059836 | 51649 | 6.7× |
| Sonic | 16625731 | 487.20 MB/s | 70887354 | 40014 | 5.4× |
| SonicFastest | 16627985 | 487.13 MB/s | 70887497 | 40014 | 5.4× |
| Goccy | 23467801 | 345.16 MB/s | 17040634 | 107148 | 3.8× |
| Easyjson | 30614257 | 264.58 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 40220985 | 129.36 MB/s | 34333352 | 912810 | 2.2× |
| JSONV2 | 44497086 | 182.04 MB/s | 15233712 | 78972 | 2.0× |
| Stdlib | 89383340 | 90.62 MB/s | 15665070 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5665117 | 526.64 MB/s | 3985336 | 30015 | 8.3× |
| Lightning | 5747047 | 519.13 MB/s | 3985336 | 30015 | 8.2× |
| Sonic | 8657175 | 344.62 MB/s | 26616345 | 56760 | 5.4× |
| SonicFastest | 8693957 | 343.17 MB/s | 26583795 | 56760 | 5.4× |
| Easyjson | 16348463 | 182.49 MB/s | 9479449 | 30115 | 2.9× |
| Goccy | 16675113 | 178.92 MB/s | 10635156 | 273649 | 2.8× |
| LightningDecodeAny | 18450723 | 99.41 MB/s | 20023838 | 408557 | 2.6× |
| JSONV2 | 24137971 | 123.60 MB/s | 9257162 | 86278 | 1.9× |
| Stdlib | 47051405 | 63.41 MB/s | 9258094 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1224017 | 591.17 MB/s | 907393 | 3618 | 9.5× |
| Lightning | 1236613 | 585.14 MB/s | 907388 | 3618 | 9.4× |
| Sonic | 1765472 | 409.86 MB/s | 3176064 | 7226 | 6.6× |
| SonicFastest | 1787909 | 404.72 MB/s | 3189997 | 7226 | 6.5× |
| Easyjson | 4199956 | 172.29 MB/s | 2847906 | 3698 | 2.8× |
| LightningDecodeAny | 4350052 | 149.56 MB/s | 5752201 | 80172 | 2.7× |
| Goccy | 4769300 | 151.72 MB/s | 2770001 | 80271 | 2.4× |
| JSONV2 | 5924050 | 122.15 MB/s | 2704651 | 7318 | 2.0× |
| Stdlib | 11574010 | 62.52 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1829946 | 861.97 MB/s | 907393 | 3618 | 8.6× |
| Lightning | 1879262 | 839.35 MB/s | 907386 | 3618 | 8.4× |
| SonicFastest | 2227870 | 708.01 MB/s | 5785279 | 7226 | 7.0× |
| Sonic | 2238612 | 704.61 MB/s | 5789429 | 7226 | 7.0× |
| LightningDecodeAny | 3868291 | 194.76 MB/s | 5752202 | 80172 | 4.1× |
| Goccy | 5530560 | 285.21 MB/s | 3551853 | 80266 | 2.8× |
| Easyjson | 5551082 | 284.15 MB/s | 2847908 | 3698 | 2.8× |
| JSONV2 | 6628076 | 237.98 MB/s | 2704594 | 7318 | 2.4× |
| Stdlib | 15700595 | 100.46 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 208926 | 718.55 MB/s | 81920 | 1 | 8.7× |
| LightningDestructive | 209270 | 717.37 MB/s | 81920 | 1 | 8.7× |
| SonicFastest | 276604 | 542.74 MB/s | 259308 | 6 | 6.6× |
| Sonic | 279850 | 536.45 MB/s | 270039 | 6 | 6.5× |
| LightningDecodeAny | 476655 | 314.95 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 855885 | 175.40 MB/s | 324639 | 10004 | 2.1× |
| JSONV2 | 1100313 | 136.44 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1820153 | 82.48 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32590 | 862.76 MB/s | 30272 | 105 | 9.3× |
| LightningDestructive | 32828 | 856.49 MB/s | 30144 | 103 | 9.2× |
| SonicFastest | 62902 | 447.00 MB/s | 47046 | 103 | 4.8× |
| Sonic | 63883 | 440.13 MB/s | 47934 | 103 | 4.7× |
| Easyjson | 68642 | 409.62 MB/s | 32304 | 138 | 4.4× |
| Goccy | 71630 | 392.53 MB/s | 59202 | 188 | 4.2× |
| JSONV2 | 133680 | 210.33 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 149830 | 187.66 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303214 | 92.73 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1952 | 1192.53 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2019 | 1153.26 MB/s | 32 | 1 | 11.2× |
| Goccy | 4162 | 559.34 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4211 | 552.85 MB/s | 192 | 2 | 5.4× |
| Sonic | 5027 | 463.14 MB/s | 4220 | 6 | 4.5× |
| SonicFastest | 5071 | 459.08 MB/s | 4314 | 6 | 4.5× |
| JSONV2 | 8430 | 276.16 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10164 | 165.78 MB/s | 9960 | 195 | 2.2× |
| Stdlib | 22636 | 102.85 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 221 | 856.31 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 222 | 850.50 MB/s | 0 | 0 | 10.9× |
| Goccy | 391 | 483.16 MB/s | 304 | 2 | 6.2× |
| Easyjson | 492 | 383.77 MB/s | 0 | 0 | 4.9× |
| SonicFastest | 792 | 238.55 MB/s | 517 | 4 | 3.1× |
| Sonic | 792 | 238.51 MB/s | 512 | 4 | 3.1× |
| JSONV2 | 1033 | 182.99 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1272 | 105.38 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2433 | 77.69 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1546 | 1417.07 MB/s | 0 | 0 | 10.3× |
| LightningDestructive | 1563 | 1402.03 MB/s | 0 | 0 | 10.2× |
| Goccy | 3130 | 700.06 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3189 | 686.99 MB/s | 24 | 1 | 5.0× |
| Sonic | 6420 | 341.29 MB/s | 3941 | 40 | 2.5× |
| SonicFastest | 6432 | 340.65 MB/s | 3975 | 40 | 2.5× |
| JSONV2 | 8064 | 271.69 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8169 | 221.68 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 15988 | 137.04 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 692693 | 736.94 MB/s | 703777 | 1012 | 8.7× |
| Lightning | 727810 | 701.39 MB/s | 703780 | 1012 | 8.3× |
| Goccy | 1151021 | 443.50 MB/s | 1139364 | 5006 | 5.2× |
| Sonic | 1163251 | 438.84 MB/s | 905038 | 2006 | 5.2× |
| SonicFastest | 1167209 | 437.35 MB/s | 911965 | 2006 | 5.2× |
| Easyjson | 1520638 | 335.70 MB/s | 863777 | 3012 | 4.0× |
| JSONV2 | 3204407 | 159.30 MB/s | 1076006 | 12646 | 1.9× |
| LightningDecodeAny | 3485125 | 132.41 MB/s | 2785926 | 66022 | 1.7× |
| Stdlib | 6029559 | 84.66 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1336 | 14812.14 MB/s | 0 | 0 | 86.3× |
| LightningDestructive | 1372 | 14419.93 MB/s | 0 | 0 | 84.1× |
| Goccy | 20183 | 980.46 MB/s | 20491 | 2 | 5.7× |
| SonicFastest | 28059 | 705.28 MB/s | 22396 | 4 | 4.1× |
| Sonic | 28100 | 704.24 MB/s | 22481 | 4 | 4.1× |
| JSONV2 | 29595 | 668.65 MB/s | 8 | 1 | 3.9× |
| LightningDecodeAny | 79829 | 247.88 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 89237 | 221.76 MB/s | 0 | 0 | 1.3× |
| Stdlib | 115330 | 171.59 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2660 | 6814.73 MB/s | 0 | 0 | 38.5× |
| Lightning | 2810 | 6450.59 MB/s | 432 | 2 | 36.5× |
| Easyjson | 3932 | 4609.26 MB/s | 432 | 2 | 26.1× |
| Sonic | 9793 | 1850.79 MB/s | 22946 | 6 | 10.5× |
| SonicFastest | 10029 | 1807.14 MB/s | 23336 | 6 | 10.2× |
| Goccy | 16439 | 1102.50 MB/s | 19459 | 2 | 6.2× |
| LightningDecodeAny | 16721 | 1069.43 MB/s | 29088 | 191 | 6.1× |
| JSONV2 | 45905 | 394.82 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102499 | 176.82 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2363125 | 849.93 MB/s | 2951717 | 7277 | 7.9× |
| Lightning | 2405523 | 834.95 MB/s | 2953428 | 7283 | 7.7× |
| Goccy | 4240247 | 473.67 MB/s | 5412112 | 15830 | 4.4× |
| SonicFastest | 4507526 | 445.59 MB/s | 10970569 | 13683 | 4.1× |
| Sonic | 4514262 | 444.92 MB/s | 10917927 | 13683 | 4.1× |
| Easyjson | 4895099 | 410.31 MB/s | 2981481 | 7439 | 3.8× |
| JSONV2 | 6940473 | 289.39 MB/s | 3173681 | 14563 | 2.7× |
| LightningDecodeAny | 7128218 | 160.25 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18617579 | 107.88 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 904 | 607.35 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 912 | 601.87 MB/s | 480 | 1 | 6.2× |
| LightningDecodeAny | 1987 | 275.73 MB/s | 2261 | 50 | 2.9× |
| Easyjson | 2210 | 248.47 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2722 | 201.71 MB/s | 2055 | 26 | 2.1× |
| SonicFastest | 2735 | 200.72 MB/s | 2070 | 26 | 2.1× |
| Goccy | 3068 | 178.92 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3350 | 163.86 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5687 | 96.53 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 498864 | 1265.90 MB/s | 364472 | 566 | 10.8× |
| Lightning | 549376 | 1149.51 MB/s | 413001 | 878 | 9.8× |
| Sonic | 1011726 | 624.19 MB/s | 1002129 | 1102 | 5.3× |
| SonicFastest | 1012689 | 623.60 MB/s | 999191 | 1102 | 5.3× |
| Easyjson | 1147276 | 550.45 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1157743 | 545.47 MB/s | 986332 | 1201 | 4.7× |
| JSONV2 | 2149491 | 293.80 MB/s | 571612 | 3144 | 2.5× |
| LightningDecodeAny | 2362727 | 197.61 MB/s | 2010201 | 30295 | 2.3× |
| Stdlib | 5397400 | 117.00 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 710931 | 791.09 MB/s | 544253 | 448 | 7.4× |
| Lightning | 883222 | 636.77 MB/s | 767621 | 1254 | 6.0× |
| Sonic | 1024486 | 548.97 MB/s | 947204 | 1476 | 5.1× |
| SonicFastest | 1058624 | 531.26 MB/s | 994842 | 1476 | 5.0× |
| Goccy | 1347835 | 417.27 MB/s | 1040157 | 1030 | 3.9× |
| Easyjson | 1737966 | 323.60 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2704955 | 207.92 MB/s | 2114151 | 30295 | 1.9× |
| JSONV2 | 2789130 | 201.64 MB/s | 927440 | 3482 | 1.9× |
| Stdlib | 5258432 | 106.95 MB/s | 1011671 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 650284 | 819.92 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 670746 | 794.90 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1103144 | 483.33 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1137500 | 468.73 MB/s | 1033203 | 4351 | 4.8× |
| SonicFastest | 1139620 | 467.86 MB/s | 1035932 | 4351 | 4.8× |
| Goccy | 1299716 | 410.23 MB/s | 1167248 | 5409 | 4.2× |
| JSONV2 | 2542964 | 209.67 MB/s | 745451 | 13288 | 2.1× |
| LightningDecodeAny | 3367721 | 158.32 MB/s | 2674621 | 50596 | 1.6× |
| Stdlib | 5454505 | 97.75 MB/s | 798693 | 17133 | 1.0× |
