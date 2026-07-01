# JSON Deserialization Benchmarks

- generated 2026-07-01T21:41:19Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 103352 | 1231.48 MB/s | 49760 | 3 | 11.9× |
| LightningDestructive | 104849 | 1213.88 MB/s | 49280 | 2 | 11.8× |
| Sonic | 232523 | 547.36 MB/s | 214465 | 15 | 5.3× |
| SonicFastest | 234436 | 542.90 MB/s | 214803 | 15 | 5.3× |
| Easyjson | 237684 | 535.48 MB/s | 122864 | 14 | 5.2× |
| Goccy | 241832 | 526.30 MB/s | 225518 | 884 | 5.1× |
| JSONV2 | 386518 | 329.29 MB/s | 195127 | 1805 | 3.2× |
| LightningDecodeAny | 446189 | 212.14 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1234132 | 103.13 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4166153 | 540.32 MB/s | 3122873 | 3081 | 7.5× |
| Lightning | 4182401 | 538.22 MB/s | 3122872 | 3081 | 7.5× |
| Sonic | 5422527 | 415.13 MB/s | 4863705 | 2584 | 5.8× |
| SonicFastest | 5486229 | 410.31 MB/s | 4863538 | 2584 | 5.7× |
| LightningDecodeAny | 11714157 | 192.16 MB/s | 7938298 | 281383 | 2.7× |
| Goccy | 12831019 | 175.44 MB/s | 4129726 | 56532 | 2.4× |
| Easyjson | 13649714 | 164.92 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 16801789 | 133.98 MB/s | 3123189 | 3083 | 1.9× |
| Stdlib | 31248208 | 72.04 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 560469 | 482.46 MB/s | 348024 | 1627 | 7.2× |
| LightningDestructive | 564383 | 479.11 MB/s | 348024 | 1627 | 7.2× |
| SonicFastest | 740205 | 365.31 MB/s | 642228 | 1147 | 5.5× |
| Sonic | 745567 | 362.68 MB/s | 642335 | 1147 | 5.4× |
| LightningDecodeAny | 1606401 | 168.33 MB/s | 1011487 | 37901 | 2.5× |
| Goccy | 1750450 | 154.48 MB/s | 542809 | 8122 | 2.3× |
| Easyjson | 1752125 | 154.33 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2189040 | 123.53 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4041643 | 66.90 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1300885 | 1327.71 MB/s | 959848 | 5881 | 12.9× |
| Lightning | 1341120 | 1287.88 MB/s | 959889 | 5882 | 12.5× |
| SonicFastest | 2202817 | 784.09 MB/s | 2692766 | 5547 | 7.6× |
| Sonic | 2205677 | 783.07 MB/s | 2692851 | 5547 | 7.6× |
| Goccy | 2886165 | 598.44 MB/s | 2581481 | 14603 | 5.8× |
| JSONV2 | 3932345 | 439.23 MB/s | 1011615 | 7594 | 4.3× |
| Easyjson | 3978539 | 434.13 MB/s | 972032 | 5389 | 4.2× |
| LightningDecodeAny | 4069661 | 122.93 MB/s | 4976203 | 81466 | 4.1× |
| Stdlib | 16782373 | 102.92 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1018 | 1780.17 MB/s | 0 | 0 | 14.5× |
| LightningDestructive | 1054 | 1718.86 MB/s | 0 | 0 | 14.0× |
| Easyjson | 2818 | 642.90 MB/s | 24 | 1 | 5.2× |
| Goccy | 3122 | 580.46 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 6296 | 287.80 MB/s | 3348 | 38 | 2.3× |
| Sonic | 6509 | 278.39 MB/s | 3348 | 38 | 2.3× |
| JSONV2 | 7494 | 241.80 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8500 | 213.07 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14780 | 122.60 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1086 | 1668.16 MB/s | 0 | 0 | 13.4× |
| LightningDestructive | 1136 | 1595.33 MB/s | 0 | 0 | 12.8× |
| Easyjson | 2810 | 644.93 MB/s | 24 | 1 | 5.2× |
| Goccy | 3069 | 590.47 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6060 | 298.99 MB/s | 3348 | 38 | 2.4× |
| Sonic | 6270 | 288.98 MB/s | 3345 | 38 | 2.3× |
| JSONV2 | 7445 | 243.38 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8549 | 211.84 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14591 | 124.18 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1248 | 1452.41 MB/s | 144 | 10 | 11.6× |
| LightningDestructive | 1301 | 1392.97 MB/s | 144 | 10 | 11.1× |
| Easyjson | 2971 | 609.82 MB/s | 144 | 10 | 4.9× |
| Goccy | 3300 | 549.06 MB/s | 2600 | 5 | 4.4× |
| SonicFastest | 6262 | 289.36 MB/s | 3365 | 40 | 2.3× |
| Sonic | 6471 | 280.03 MB/s | 3367 | 40 | 2.2× |
| JSONV2 | 7978 | 227.11 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8614 | 210.24 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14504 | 124.93 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 710 | 696.25 MB/s | 160 | 1 | 8.5× |
| LightningDestructive | 713 | 692.62 MB/s | 160 | 1 | 8.5× |
| Sonic | 1215 | 406.65 MB/s | 1075 | 8 | 5.0× |
| SonicFastest | 1222 | 404.39 MB/s | 1075 | 8 | 4.9× |
| LightningDecodeAny | 1607 | 306.86 MB/s | 1536 | 30 | 3.8× |
| Easyjson | 2392 | 206.51 MB/s | 448 | 3 | 2.5× |
| Goccy | 2576 | 191.81 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 2974 | 166.13 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6044 | 81.74 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 428 | 537.44 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 429 | 535.67 MB/s | 160 | 1 | 9.8× |
| Sonic | 852 | 269.94 MB/s | 801 | 8 | 5.0× |
| SonicFastest | 860 | 267.33 MB/s | 801 | 8 | 4.9× |
| LightningDecodeAny | 1396 | 164.02 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1542 | 149.19 MB/s | 448 | 3 | 2.7× |
| Goccy | 1726 | 133.27 MB/s | 584 | 23 | 2.4× |
| JSONV2 | 2309 | 99.60 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4225 | 54.44 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 85216 | 764.32 MB/s | 172433 | 107 | 7.6× |
| LightningDestructive | 89209 | 730.11 MB/s | 166213 | 102 | 7.2× |
| Sonic | 152723 | 426.47 MB/s | 235913 | 65 | 4.2× |
| SonicFastest | 157369 | 413.88 MB/s | 236064 | 65 | 4.1× |
| Goccy | 186786 | 348.70 MB/s | 228596 | 134 | 3.5× |
| LightningDecodeAny | 187242 | 284.81 MB/s | 176961 | 3252 | 3.4× |
| JSONV2 | 246578 | 264.14 MB/s | 206665 | 607 | 2.6× |
| Stdlib | 645269 | 100.94 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2652710 | 731.51 MB/s | 2846865 | 2238 | 9.7× |
| Lightning | 2685254 | 722.64 MB/s | 2846867 | 2238 | 9.6× |
| Sonic | 4805516 | 403.80 MB/s | 4875230 | 1736 | 5.4× |
| SonicFastest | 4837390 | 401.14 MB/s | 4875556 | 1736 | 5.3× |
| Goccy | 5295322 | 366.45 MB/s | 4062546 | 13509 | 4.9× |
| Easyjson | 7980582 | 243.15 MB/s | 3871265 | 15043 | 3.2× |
| LightningDecodeAny | 9743113 | 199.16 MB/s | 7013524 | 219937 | 2.6× |
| JSONV2 | 11584032 | 167.51 MB/s | 3237190 | 13947 | 2.2× |
| Stdlib | 25743240 | 75.38 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1013762 | 3282.65 MB/s | 351704 | 1286 | 25.1× |
| Lightning | 1655894 | 2009.69 MB/s | 2488906 | 2995 | 15.4× |
| Sonic | 2230884 | 1491.71 MB/s | 5896090 | 4263 | 11.4× |
| SonicFastest | 2244429 | 1482.71 MB/s | 5894304 | 4263 | 11.3× |
| LightningDecodeAny | 3357720 | 915.43 MB/s | 4886621 | 56892 | 7.6× |
| Goccy | 6638562 | 501.29 MB/s | 3948910 | 3816 | 3.8× |
| JSONV2 | 8588634 | 387.47 MB/s | 5364504 | 13243 | 3.0× |
| Stdlib | 25431134 | 130.86 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 227358 | 969.16 MB/s | 136896 | 228 | 9.8× |
| LightningDestructive | 236953 | 929.91 MB/s | 136896 | 228 | 9.4× |
| Goccy | 476139 | 462.78 MB/s | 363056 | 1066 | 4.7× |
| Sonic | 533436 | 413.07 MB/s | 352279 | 262 | 4.2× |
| SonicFastest | 535146 | 411.75 MB/s | 352142 | 262 | 4.2× |
| Easyjson | 596855 | 369.18 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 651425 | 338.25 MB/s | 129746 | 470 | 3.4× |
| LightningDecodeAny | 946221 | 114.47 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2226892 | 98.95 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 15198839 | 532.94 MB/s | 15059832 | 51649 | 6.7× |
| Lightning | 15270168 | 530.45 MB/s | 15059845 | 51649 | 6.6× |
| SonicFastest | 20522163 | 394.70 MB/s | 19854469 | 41640 | 4.9× |
| Sonic | 20564121 | 393.89 MB/s | 19856500 | 41640 | 4.9× |
| Goccy | 26381108 | 307.04 MB/s | 19048377 | 107156 | 3.8× |
| Easyjson | 34070949 | 237.74 MB/s | 15059618 | 41643 | 3.0× |
| LightningDecodeAny | 41886205 | 124.22 MB/s | 34333347 | 912810 | 2.4× |
| JSONV2 | 46161272 | 175.47 MB/s | 15233720 | 78972 | 2.2× |
| Stdlib | 101544085 | 79.77 MB/s | 15665071 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5657861 | 527.31 MB/s | 3985336 | 30015 | 9.0× |
| Lightning | 5860502 | 509.08 MB/s | 3985340 | 30015 | 8.6× |
| Sonic | 9373809 | 318.28 MB/s | 9130010 | 57804 | 5.4× |
| SonicFastest | 9483885 | 314.58 MB/s | 9129521 | 57804 | 5.3× |
| Goccy | 17799581 | 167.61 MB/s | 9925761 | 273621 | 2.8× |
| Easyjson | 18163425 | 164.26 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 19629614 | 93.44 MB/s | 20023838 | 408557 | 2.6× |
| JSONV2 | 24662454 | 120.97 MB/s | 9257029 | 86278 | 2.1× |
| Stdlib | 50680477 | 58.87 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1260497 | 574.06 MB/s | 907393 | 3618 | 9.7× |
| Lightning | 1321744 | 547.46 MB/s | 907388 | 3618 | 9.3× |
| Sonic | 2157147 | 335.44 MB/s | 2367456 | 3683 | 5.7× |
| SonicFastest | 2176935 | 332.39 MB/s | 2367692 | 3683 | 5.6× |
| Easyjson | 5248624 | 137.86 MB/s | 2847905 | 3698 | 2.3× |
| Goccy | 5288415 | 136.83 MB/s | 2661630 | 80265 | 2.3× |
| LightningDecodeAny | 5415314 | 120.14 MB/s | 5752203 | 80172 | 2.3× |
| JSONV2 | 6084932 | 118.92 MB/s | 2704704 | 7318 | 2.0× |
| Stdlib | 12288459 | 58.88 MB/s | 2704547 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1999699 | 788.80 MB/s | 907386 | 3618 | 8.9× |
| LightningDestructive | 2008392 | 785.38 MB/s | 907392 | 3618 | 8.8× |
| Sonic | 2696543 | 584.95 MB/s | 3237406 | 3683 | 6.6× |
| SonicFastest | 2832607 | 556.86 MB/s | 3240595 | 3683 | 6.2× |
| LightningDecodeAny | 4483627 | 168.03 MB/s | 5752202 | 80172 | 3.9× |
| Easyjson | 6436335 | 245.07 MB/s | 2847904 | 3698 | 2.7× |
| Goccy | 6600336 | 238.98 MB/s | 3466672 | 80260 | 2.7× |
| JSONV2 | 6782376 | 232.57 MB/s | 2704552 | 7318 | 2.6× |
| Stdlib | 17698166 | 89.13 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 214951 | 698.41 MB/s | 81920 | 1 | 10.4× |
| LightningDestructive | 223546 | 671.56 MB/s | 81920 | 1 | 10.0× |
| Sonic | 398220 | 376.99 MB/s | 407047 | 16 | 5.6× |
| SonicFastest | 417249 | 359.79 MB/s | 406832 | 16 | 5.4× |
| LightningDecodeAny | 588369 | 255.15 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 988999 | 151.79 MB/s | 325219 | 10005 | 2.3× |
| JSONV2 | 1136286 | 132.12 MB/s | 357728 | 20 | 2.0× |
| Stdlib | 2234056 | 67.20 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 31847 | 882.87 MB/s | 30272 | 105 | 9.8× |
| LightningDestructive | 32801 | 857.20 MB/s | 30144 | 103 | 9.5× |
| SonicFastest | 67603 | 415.92 MB/s | 59481 | 83 | 4.6× |
| Sonic | 67768 | 414.90 MB/s | 59461 | 83 | 4.6× |
| Easyjson | 75568 | 372.08 MB/s | 32304 | 138 | 4.1× |
| Goccy | 77153 | 364.43 MB/s | 59284 | 188 | 4.0× |
| JSONV2 | 126093 | 222.99 MB/s | 36896 | 242 | 2.5× |
| LightningDecodeAny | 153032 | 183.73 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 311779 | 90.18 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1963 | 1186.05 MB/s | 32 | 1 | 12.0× |
| LightningDestructive | 2001 | 1163.62 MB/s | 32 | 1 | 11.8× |
| Easyjson | 4790 | 485.98 MB/s | 192 | 2 | 4.9× |
| Goccy | 4856 | 479.42 MB/s | 3649 | 4 | 4.8× |
| SonicFastest | 6319 | 368.41 MB/s | 3717 | 4 | 3.7× |
| Sonic | 6341 | 367.13 MB/s | 3718 | 4 | 3.7× |
| JSONV2 | 7826 | 297.45 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 9638 | 174.82 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 23547 | 98.87 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222 | 852.26 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 228 | 830.11 MB/s | 0 | 0 | 10.8× |
| Goccy | 398 | 474.47 MB/s | 304 | 2 | 6.2× |
| Easyjson | 572 | 330.66 MB/s | 0 | 0 | 4.3× |
| SonicFastest | 767 | 246.45 MB/s | 341 | 3 | 3.2× |
| Sonic | 770 | 245.30 MB/s | 340 | 3 | 3.2× |
| JSONV2 | 920 | 205.55 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1196 | 112.08 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2453 | 77.04 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1473 | 1487.71 MB/s | 0 | 0 | 11.9× |
| LightningDestructive | 1532 | 1430.42 MB/s | 0 | 0 | 11.4× |
| Easyjson | 3478 | 629.88 MB/s | 24 | 1 | 5.0× |
| Goccy | 3547 | 617.65 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 6700 | 327.00 MB/s | 3604 | 38 | 2.6× |
| Sonic | 6848 | 319.96 MB/s | 3602 | 38 | 2.6× |
| JSONV2 | 7677 | 285.39 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8489 | 213.33 MB/s | 7536 | 158 | 2.1× |
| Stdlib | 17496 | 125.23 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 671046 | 760.72 MB/s | 703777 | 1012 | 9.7× |
| Lightning | 741032 | 688.87 MB/s | 703778 | 1012 | 8.8× |
| Goccy | 1270279 | 401.86 MB/s | 1133315 | 5006 | 5.1× |
| Sonic | 1371931 | 372.09 MB/s | 1307352 | 2014 | 4.8× |
| SonicFastest | 1408764 | 362.36 MB/s | 1309698 | 2014 | 4.6× |
| Easyjson | 1596143 | 319.82 MB/s | 863781 | 3012 | 4.1× |
| JSONV2 | 3121369 | 163.54 MB/s | 1075947 | 12645 | 2.1× |
| LightningDecodeAny | 3535460 | 130.53 MB/s | 2785927 | 66022 | 1.8× |
| Stdlib | 6533782 | 78.13 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 612 | 32334.47 MB/s | 0 | 0 | 233.5× |
| LightningDestructive | 879 | 22523.90 MB/s | 0 | 0 | 162.6× |
| SonicFastest | 6392 | 3095.72 MB/s | 21121 | 3 | 22.4× |
| Goccy | 25213 | 784.87 MB/s | 20492 | 2 | 5.7× |
| Sonic | 29203 | 677.64 MB/s | 20627 | 3 | 4.9× |
| JSONV2 | 35996 | 549.75 MB/s | 8 | 1 | 4.0× |
| LightningDecodeAny | 93819 | 210.92 MB/s | 117104 | 2019 | 1.5× |
| Easyjson | 107197 | 184.60 MB/s | 0 | 0 | 1.3× |
| Stdlib | 142879 | 138.50 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2258 | 8028.13 MB/s | 432 | 2 | 53.0× |
| LightningDestructive | 2382 | 7608.89 MB/s | 0 | 0 | 50.2× |
| Easyjson | 4706 | 3851.45 MB/s | 432 | 2 | 25.4× |
| SonicFastest | 9098 | 1992.15 MB/s | 20468 | 5 | 13.1× |
| Sonic | 9813 | 1846.90 MB/s | 20425 | 5 | 12.2× |
| LightningDecodeAny | 17927 | 997.46 MB/s | 29088 | 191 | 6.7× |
| Goccy | 21078 | 859.84 MB/s | 19460 | 2 | 5.7× |
| JSONV2 | 49114 | 369.02 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 119609 | 151.53 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2250069 | 892.64 MB/s | 2951717 | 7277 | 9.2× |
| Lightning | 2313357 | 868.22 MB/s | 2953429 | 7283 | 9.0× |
| Goccy | 4679241 | 429.24 MB/s | 5410637 | 15829 | 4.4× |
| SonicFastest | 4723409 | 425.22 MB/s | 5156392 | 7085 | 4.4× |
| Sonic | 4793977 | 418.96 MB/s | 5152522 | 7085 | 4.3× |
| Easyjson | 5105194 | 393.42 MB/s | 2981492 | 7439 | 4.1× |
| LightningDecodeAny | 6798124 | 168.03 MB/s | 7386072 | 134751 | 3.1× |
| JSONV2 | 7387629 | 271.87 MB/s | 3173675 | 14563 | 2.8× |
| Stdlib | 20786828 | 96.62 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 810 | 678.03 MB/s | 480 | 1 | 7.2× |
| LightningDestructive | 818 | 671.16 MB/s | 480 | 1 | 7.1× |
| LightningDecodeAny | 1885 | 290.65 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 2010 | 273.15 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2133 | 257.35 MB/s | 2262 | 8 | 2.7× |
| Sonic | 2182 | 251.55 MB/s | 2262 | 8 | 2.7× |
| Goccy | 2979 | 184.28 MB/s | 2129 | 43 | 1.9× |
| JSONV2 | 3666 | 149.75 MB/s | 1664 | 7 | 1.6× |
| Stdlib | 5807 | 94.54 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 479506 | 1317.01 MB/s | 364472 | 566 | 12.5× |
| Lightning | 549796 | 1148.63 MB/s | 413001 | 878 | 10.9× |
| Sonic | 1113686 | 567.05 MB/s | 1071304 | 814 | 5.4× |
| SonicFastest | 1123792 | 561.95 MB/s | 1070601 | 814 | 5.3× |
| Easyjson | 1221406 | 517.04 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1409458 | 448.05 MB/s | 990120 | 1200 | 4.3× |
| JSONV2 | 2214397 | 285.19 MB/s | 571593 | 3144 | 2.7× |
| LightningDecodeAny | 2364484 | 197.47 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6002047 | 105.22 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 780609 | 720.47 MB/s | 544249 | 448 | 7.4× |
| Lightning | 973022 | 578.00 MB/s | 767623 | 1254 | 5.9× |
| Sonic | 1456154 | 386.23 MB/s | 1349740 | 1185 | 3.9× |
| SonicFastest | 1481812 | 379.54 MB/s | 1350785 | 1185 | 3.9× |
| Goccy | 1618345 | 347.52 MB/s | 1040916 | 1028 | 3.6× |
| Easyjson | 2014196 | 279.22 MB/s | 775155 | 1254 | 2.9× |
| LightningDecodeAny | 2995957 | 187.72 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 3207820 | 175.32 MB/s | 927407 | 3482 | 1.8× |
| Stdlib | 5751063 | 97.79 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 595170 | 895.84 MB/s | 333416 | 2084 | 10.1× |
| Lightning | 672381 | 792.97 MB/s | 368224 | 2293 | 8.9× |
| Easyjson | 1250390 | 426.41 MB/s | 428362 | 3273 | 4.8× |
| SonicFastest | 1406097 | 379.19 MB/s | 981759 | 3082 | 4.3× |
| Sonic | 1409524 | 378.27 MB/s | 981833 | 3082 | 4.3× |
| Goccy | 1441293 | 369.93 MB/s | 1167056 | 5408 | 4.2× |
| JSONV2 | 2669035 | 199.76 MB/s | 745423 | 13288 | 2.3× |
| LightningDecodeAny | 3318891 | 160.65 MB/s | 2674617 | 50596 | 1.8× |
| Stdlib | 6014505 | 88.65 MB/s | 798693 | 17133 | 1.0× |
