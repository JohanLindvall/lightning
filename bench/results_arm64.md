# JSON Deserialization Benchmarks

- generated 2026-07-02T12:28:59Z
- go version go1.26.4 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 105451 | 1206.96 MB/s | 49280 | 2 | 10.5× |
| Lightning | 105646 | 1204.73 MB/s | 49760 | 3 | 10.4× |
| Sonic | 181845 | 699.91 MB/s | 196571 | 10 | 6.1× |
| SonicFastest | 182429 | 697.67 MB/s | 200271 | 10 | 6.1× |
| Goccy | 192519 | 661.10 MB/s | 225499 | 884 | 5.7× |
| Easyjson | 210836 | 603.67 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 423217 | 300.73 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 440355 | 214.95 MB/s | 465729 | 9708 | 2.5× |
| Stdlib | 1103962 | 115.29 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3585528 | 627.82 MB/s | 3008240 | 1158 | 7.3× |
| Lightning | 3621828 | 621.52 MB/s | 3008240 | 1158 | 7.2× |
| SonicFastest | 4572443 | 492.31 MB/s | 15233809 | 970 | 5.7× |
| Sonic | 4663490 | 482.70 MB/s | 15247018 | 970 | 5.6× |
| Goccy | 10463792 | 215.13 MB/s | 4126719 | 56532 | 2.5× |
| Easyjson | 10975721 | 205.09 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 12185461 | 184.73 MB/s | 19380215 | 223896 | 2.1× |
| JSONV2 | 16331187 | 137.84 MB/s | 3123213 | 3083 | 1.6× |
| Stdlib | 26095249 | 86.26 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 456918 | 591.80 MB/s | 392944 | 568 | 7.3× |
| LightningDestructive | 462862 | 584.20 MB/s | 392944 | 568 | 7.2× |
| SonicFastest | 631113 | 428.45 MB/s | 484804 | 968 | 5.3× |
| Sonic | 633960 | 426.53 MB/s | 497090 | 968 | 5.3× |
| Goccy | 1410945 | 191.65 MB/s | 542908 | 8122 | 2.4× |
| Easyjson | 1424222 | 189.86 MB/s | 330272 | 749 | 2.4× |
| LightningDecodeAny | 1596521 | 169.37 MB/s | 2543876 | 29687 | 2.1× |
| JSONV2 | 2107531 | 128.30 MB/s | 348156 | 1628 | 1.6× |
| Stdlib | 3350729 | 80.70 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1183183 | 1459.79 MB/s | 767864 | 2798 | 11.2× |
| Lightning | 1190723 | 1450.55 MB/s | 767906 | 2799 | 11.1× |
| Sonic | 2035272 | 848.64 MB/s | 2763661 | 4020 | 6.5× |
| SonicFastest | 2041269 | 846.14 MB/s | 2751273 | 4020 | 6.5× |
| Goccy | 2341108 | 737.77 MB/s | 2582453 | 14604 | 5.7× |
| JSONV2 | 4209054 | 410.35 MB/s | 1011636 | 7594 | 3.1× |
| Easyjson | 4222614 | 409.04 MB/s | 972032 | 5389 | 3.1× |
| LightningDecodeAny | 4272908 | 117.09 MB/s | 4954733 | 76576 | 3.1× |
| Stdlib | 13240451 | 130.45 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1190 | 1522.44 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1206 | 1502.36 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2555 | 709.18 MB/s | 24 | 1 | 5.5× |
| Goccy | 2796 | 648.01 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5892 | 307.53 MB/s | 3762 | 40 | 2.4× |
| SonicFastest | 5919 | 306.15 MB/s | 3771 | 40 | 2.4× |
| JSONV2 | 7709 | 235.05 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8109 | 223.34 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14009 | 129.34 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1230 | 1473.44 MB/s | 0 | 0 | 11.4× |
| LightningDestructive | 1235 | 1467.10 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2558 | 708.43 MB/s | 24 | 1 | 5.5× |
| Goccy | 2815 | 643.76 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5923 | 305.90 MB/s | 3734 | 40 | 2.4× |
| SonicFastest | 5944 | 304.85 MB/s | 3761 | 40 | 2.4× |
| JSONV2 | 7703 | 235.22 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8079 | 224.17 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14058 | 128.90 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1400 | 1293.94 MB/s | 144 | 10 | 10.0× |
| LightningDestructive | 1459 | 1241.84 MB/s | 144 | 10 | 9.6× |
| Easyjson | 2745 | 660.12 MB/s | 144 | 10 | 5.1× |
| Goccy | 2900 | 624.80 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6132 | 295.52 MB/s | 3863 | 42 | 2.3× |
| SonicFastest | 6177 | 293.35 MB/s | 3879 | 42 | 2.3× |
| JSONV2 | 7965 | 227.51 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8401 | 215.56 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14044 | 129.02 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 694 | 711.28 MB/s | 160 | 1 | 8.0× |
| Lightning | 696 | 709.85 MB/s | 160 | 1 | 7.9× |
| Sonic | 1227 | 402.64 MB/s | 973 | 6 | 4.5× |
| SonicFastest | 1229 | 402.10 MB/s | 983 | 6 | 4.5× |
| LightningDecodeAny | 1371 | 359.59 MB/s | 1296 | 26 | 4.0× |
| Easyjson | 2221 | 222.40 MB/s | 448 | 3 | 2.5× |
| Goccy | 2424 | 203.78 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3206 | 154.09 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5529 | 89.35 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 424 | 542.38 MB/s | 160 | 1 | 9.7× |
| Lightning | 426 | 539.95 MB/s | 160 | 1 | 9.7× |
| Sonic | 879 | 261.57 MB/s | 658 | 6 | 4.7× |
| SonicFastest | 880 | 261.21 MB/s | 660 | 6 | 4.7× |
| LightningDecodeAny | 1115 | 205.37 MB/s | 1296 | 26 | 3.7× |
| Easyjson | 1398 | 164.49 MB/s | 448 | 3 | 2.9× |
| Goccy | 1579 | 145.65 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2374 | 96.89 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4120 | 55.82 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 69191 | 941.34 MB/s | 158660 | 100 | 7.9× |
| Lightning | 70209 | 927.69 MB/s | 164880 | 105 | 7.8× |
| Sonic | 99554 | 654.24 MB/s | 159800 | 75 | 5.5× |
| SonicFastest | 100163 | 650.26 MB/s | 160855 | 75 | 5.4× |
| Goccy | 146811 | 443.65 MB/s | 228106 | 134 | 3.7× |
| LightningDecodeAny | 185503 | 287.48 MB/s | 180224 | 3245 | 2.9× |
| JSONV2 | 227914 | 285.77 MB/s | 206652 | 607 | 2.4× |
| Stdlib | 544291 | 119.66 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2511431 | 772.66 MB/s | 2813904 | 1358 | 9.3× |
| Lightning | 2603866 | 745.23 MB/s | 2813906 | 1358 | 9.0× |
| Sonic | 4542728 | 427.16 MB/s | 14608572 | 1407 | 5.1× |
| SonicFastest | 4619048 | 420.10 MB/s | 14608572 | 1407 | 5.1× |
| Goccy | 4802702 | 404.04 MB/s | 4065514 | 13510 | 4.9× |
| Easyjson | 7394287 | 262.43 MB/s | 3871265 | 15043 | 3.2× |
| LightningDecodeAny | 9352129 | 207.49 MB/s | 7064788 | 218633 | 2.5× |
| JSONV2 | 11156447 | 173.93 MB/s | 3237230 | 13947 | 2.1× |
| Stdlib | 23390832 | 82.96 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1096595 | 3034.69 MB/s | 351704 | 1286 | 19.0× |
| Lightning | 1779289 | 1870.31 MB/s | 2488904 | 2995 | 11.7× |
| Sonic | 2630901 | 1264.90 MB/s | 6452370 | 4248 | 7.9× |
| SonicFastest | 2641661 | 1259.75 MB/s | 6421501 | 4248 | 7.9× |
| LightningDecodeAny | 3704900 | 829.65 MB/s | 4886622 | 56892 | 5.6× |
| Goccy | 4454220 | 747.12 MB/s | 3948909 | 3816 | 4.7× |
| JSONV2 | 7395808 | 449.96 MB/s | 5364522 | 13243 | 2.8× |
| Stdlib | 20812201 | 159.90 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220479 | 999.39 MB/s | 135872 | 226 | 9.2× |
| LightningDestructive | 222240 | 991.48 MB/s | 135872 | 226 | 9.2× |
| Sonic | 373691 | 589.65 MB/s | 296904 | 398 | 5.5× |
| SonicFastest | 385332 | 571.83 MB/s | 331400 | 398 | 5.3× |
| Goccy | 434231 | 507.44 MB/s | 364136 | 1067 | 4.7× |
| Easyjson | 546317 | 403.33 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 725676 | 303.64 MB/s | 129742 | 470 | 2.8× |
| LightningDecodeAny | 842526 | 128.56 MB/s | 897521 | 11703 | 2.4× |
| Stdlib | 2039302 | 108.05 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12120777 | 668.28 MB/s | 15730640 | 20821 | 7.4× |
| Lightning | 12602832 | 642.72 MB/s | 15730647 | 20821 | 7.1× |
| Sonic | 16781278 | 482.68 MB/s | 70901559 | 40014 | 5.3× |
| SonicFastest | 16835917 | 481.12 MB/s | 70887412 | 40014 | 5.3× |
| Goccy | 23358344 | 346.77 MB/s | 16999092 | 107148 | 3.8× |
| Easyjson | 30891339 | 262.21 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 36849135 | 141.20 MB/s | 46191125 | 747112 | 2.4× |
| JSONV2 | 43734367 | 185.21 MB/s | 15233737 | 78972 | 2.0× |
| Stdlib | 89407573 | 90.60 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5685821 | 524.72 MB/s | 3908872 | 29356 | 8.3× |
| Lightning | 5793170 | 515.00 MB/s | 3908875 | 29356 | 8.1× |
| Sonic | 8552508 | 348.84 MB/s | 26591334 | 56760 | 5.5× |
| SonicFastest | 8594615 | 347.13 MB/s | 26512420 | 56760 | 5.5× |
| Easyjson | 16317285 | 182.84 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16328354 | 182.72 MB/s | 10602963 | 273647 | 2.9× |
| LightningDecodeAny | 16557951 | 110.77 MB/s | 23982397 | 351152 | 2.8× |
| JSONV2 | 24473333 | 121.91 MB/s | 9257173 | 86278 | 1.9× |
| Stdlib | 47110861 | 63.33 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1243095 | 582.09 MB/s | 907600 | 3618 | 9.3× |
| Lightning | 1257246 | 575.54 MB/s | 907596 | 3618 | 9.2× |
| Sonic | 1754898 | 412.33 MB/s | 3177721 | 7226 | 6.6× |
| SonicFastest | 1773017 | 408.12 MB/s | 3184960 | 7226 | 6.5× |
| LightningDecodeAny | 4181179 | 155.60 MB/s | 6500457 | 76546 | 2.8× |
| Easyjson | 4234910 | 170.86 MB/s | 2847907 | 3698 | 2.7× |
| Goccy | 4790078 | 151.06 MB/s | 2770968 | 80271 | 2.4× |
| JSONV2 | 5786960 | 125.04 MB/s | 2704695 | 7318 | 2.0× |
| Stdlib | 11556227 | 62.62 MB/s | 2704553 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1835355 | 859.43 MB/s | 907601 | 3618 | 8.5× |
| Lightning | 1885365 | 836.63 MB/s | 907594 | 3618 | 8.3× |
| Sonic | 2215822 | 711.86 MB/s | 5783791 | 7226 | 7.1× |
| SonicFastest | 2225386 | 708.80 MB/s | 5783083 | 7226 | 7.1× |
| LightningDecodeAny | 3837756 | 196.31 MB/s | 6500460 | 76546 | 4.1× |
| Easyjson | 5587572 | 282.30 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5600825 | 281.63 MB/s | 3589700 | 80267 | 2.8× |
| JSONV2 | 6586752 | 239.47 MB/s | 2704594 | 7318 | 2.4× |
| Stdlib | 15691385 | 100.52 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 210631 | 712.73 MB/s | 81920 | 1 | 8.6× |
| Lightning | 210657 | 712.65 MB/s | 81920 | 1 | 8.6× |
| Sonic | 271238 | 553.48 MB/s | 254941 | 6 | 6.7× |
| SonicFastest | 271654 | 552.63 MB/s | 256631 | 6 | 6.7× |
| LightningDecodeAny | 467395 | 321.19 MB/s | 745764 | 10016 | 3.9× |
| Goccy | 869139 | 172.73 MB/s | 324376 | 10004 | 2.1× |
| JSONV2 | 1074170 | 139.76 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1820388 | 82.47 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32177 | 873.82 MB/s | 29216 | 103 | 9.4× |
| LightningDestructive | 32190 | 873.47 MB/s | 29088 | 101 | 9.4× |
| Sonic | 62646 | 448.82 MB/s | 46491 | 103 | 4.8× |
| SonicFastest | 63081 | 445.73 MB/s | 46690 | 103 | 4.8× |
| Easyjson | 67599 | 415.94 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71673 | 392.30 MB/s | 59193 | 188 | 4.2× |
| JSONV2 | 134207 | 209.50 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 147774 | 190.27 MB/s | 140592 | 2643 | 2.1× |
| Stdlib | 303712 | 92.58 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1966 | 1184.27 MB/s | 32 | 1 | 11.5× |
| LightningDestructive | 2021 | 1151.82 MB/s | 32 | 1 | 11.2× |
| Goccy | 4160 | 559.65 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4198 | 554.53 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5034 | 462.43 MB/s | 4259 | 6 | 4.5× |
| Sonic | 5037 | 462.19 MB/s | 4257 | 6 | 4.5× |
| JSONV2 | 8457 | 275.27 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10123 | 166.45 MB/s | 10200 | 195 | 2.2× |
| Stdlib | 22674 | 102.67 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222 | 850.09 MB/s | 0 | 0 | 11.1× |
| LightningDestructive | 223 | 846.43 MB/s | 0 | 0 | 11.1× |
| Goccy | 374 | 504.76 MB/s | 304 | 2 | 6.6× |
| Easyjson | 488 | 387.14 MB/s | 0 | 0 | 5.1× |
| Sonic | 793 | 238.33 MB/s | 499 | 4 | 3.1× |
| SonicFastest | 794 | 238.13 MB/s | 502 | 4 | 3.1× |
| JSONV2 | 1036 | 182.39 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1209 | 110.85 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2471 | 76.50 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1555 | 1409.45 MB/s | 0 | 0 | 10.3× |
| LightningDestructive | 1575 | 1391.38 MB/s | 0 | 0 | 10.2× |
| Goccy | 3151 | 695.24 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3184 | 688.17 MB/s | 24 | 1 | 5.0× |
| Sonic | 6314 | 347.02 MB/s | 3952 | 40 | 2.5× |
| SonicFastest | 6336 | 345.79 MB/s | 3992 | 40 | 2.5× |
| JSONV2 | 7977 | 274.67 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8062 | 224.62 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 15995 | 136.98 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 685096 | 745.12 MB/s | 703297 | 1010 | 8.8× |
| Lightning | 725067 | 704.04 MB/s | 703300 | 1010 | 8.3× |
| Goccy | 1132844 | 450.61 MB/s | 1138310 | 5006 | 5.3× |
| Sonic | 1157032 | 441.19 MB/s | 912803 | 2006 | 5.2× |
| SonicFastest | 1160748 | 439.78 MB/s | 919309 | 2006 | 5.2× |
| Easyjson | 1535925 | 332.36 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3215224 | 158.77 MB/s | 1076006 | 12646 | 1.9× |
| LightningDecodeAny | 3353161 | 137.62 MB/s | 2929688 | 64018 | 1.8× |
| Stdlib | 6047099 | 84.42 MB/s | 1162120 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1351 | 14645.01 MB/s | 0 | 0 | 82.5× |
| LightningDestructive | 1388 | 14252.18 MB/s | 0 | 0 | 80.3× |
| Goccy | 19816 | 998.63 MB/s | 20491 | 2 | 5.6× |
| Sonic | 26922 | 735.04 MB/s | 22435 | 4 | 4.1× |
| SonicFastest | 26960 | 734.00 MB/s | 22409 | 4 | 4.1× |
| JSONV2 | 29853 | 662.89 MB/s | 8 | 1 | 3.7× |
| LightningDecodeAny | 73872 | 267.87 MB/s | 116864 | 2015 | 1.5× |
| Easyjson | 86037 | 230.01 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111510 | 177.46 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2666 | 6797.26 MB/s | 0 | 0 | 38.5× |
| Lightning | 2825 | 6414.73 MB/s | 432 | 2 | 36.3× |
| Easyjson | 3948 | 4590.73 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 9817 | 1846.25 MB/s | 22817 | 6 | 10.5× |
| Sonic | 9834 | 1842.92 MB/s | 23041 | 6 | 10.4× |
| Goccy | 15969 | 1134.93 MB/s | 19459 | 2 | 6.4× |
| LightningDecodeAny | 16429 | 1088.44 MB/s | 29088 | 191 | 6.2× |
| JSONV2 | 46204 | 392.26 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102677 | 176.52 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2339733 | 858.43 MB/s | 3089821 | 6821 | 8.0× |
| Lightning | 2396032 | 838.26 MB/s | 3091534 | 6827 | 7.8× |
| Goccy | 4241477 | 473.54 MB/s | 5412518 | 15832 | 4.4× |
| Sonic | 4386296 | 457.90 MB/s | 10929423 | 13683 | 4.2× |
| SonicFastest | 4392624 | 457.24 MB/s | 10944035 | 13683 | 4.2× |
| Easyjson | 4899710 | 409.92 MB/s | 2981482 | 7439 | 3.8× |
| JSONV2 | 6893665 | 291.35 MB/s | 3173688 | 14563 | 2.7× |
| LightningDecodeAny | 7369942 | 154.99 MB/s | 8498330 | 134008 | 2.5× |
| Stdlib | 18627114 | 107.83 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 868 | 632.59 MB/s | 480 | 1 | 6.5× |
| LightningDestructive | 871 | 630.18 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1665 | 329.22 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 2147 | 255.66 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2664 | 206.09 MB/s | 1908 | 26 | 2.1× |
| SonicFastest | 2668 | 205.79 MB/s | 1943 | 26 | 2.1× |
| Goccy | 3073 | 178.63 MB/s | 2128 | 43 | 1.8× |
| JSONV2 | 3320 | 165.38 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5644 | 97.27 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 501559 | 1259.10 MB/s | 402729 | 545 | 10.8× |
| Lightning | 563244 | 1121.21 MB/s | 451257 | 857 | 9.6× |
| Sonic | 1004963 | 628.40 MB/s | 996463 | 1102 | 5.4× |
| SonicFastest | 1009007 | 625.88 MB/s | 1010942 | 1102 | 5.4× |
| Easyjson | 1149525 | 549.37 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1154590 | 546.96 MB/s | 987756 | 1202 | 4.7× |
| JSONV2 | 2145683 | 294.32 MB/s | 571624 | 3144 | 2.5× |
| LightningDecodeAny | 2397962 | 194.71 MB/s | 2077368 | 30126 | 2.3× |
| Stdlib | 5400460 | 116.94 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 666578 | 843.72 MB/s | 579338 | 429 | 7.9× |
| Lightning | 838021 | 671.11 MB/s | 802710 | 1235 | 6.3× |
| Sonic | 1034019 | 543.91 MB/s | 977213 | 1476 | 5.1× |
| SonicFastest | 1047542 | 536.88 MB/s | 987077 | 1476 | 5.0× |
| Goccy | 1348764 | 416.98 MB/s | 1039322 | 1029 | 3.9× |
| Easyjson | 1741330 | 322.98 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2689830 | 209.09 MB/s | 2181319 | 30126 | 2.0× |
| JSONV2 | 2807269 | 200.34 MB/s | 927442 | 3482 | 1.9× |
| Stdlib | 5270513 | 106.71 MB/s | 1011674 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 650566 | 819.56 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 669011 | 796.96 MB/s | 368224 | 2293 | 8.1× |
| Easyjson | 1100906 | 484.31 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1135628 | 469.50 MB/s | 1026065 | 4351 | 4.8× |
| SonicFastest | 1147942 | 464.46 MB/s | 1035930 | 4351 | 4.7× |
| Goccy | 1294105 | 412.01 MB/s | 1167231 | 5409 | 4.2× |
| JSONV2 | 2534601 | 210.36 MB/s | 745449 | 13288 | 2.1× |
| LightningDecodeAny | 3368291 | 158.29 MB/s | 2991148 | 50076 | 1.6× |
| Stdlib | 5446520 | 97.89 MB/s | 798693 | 17133 | 1.0× |
