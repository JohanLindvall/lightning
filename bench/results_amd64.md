# JSON Deserialization Benchmarks

- generated 2026-09-23T13:35:01Z
- go version go1.26.8 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 80654 | 1578.04 MB/s | 49862 | 2 | 16.6× |
| Lightning | 81230 | 1566.84 MB/s | 49865 | 2 | 16.4× |
| LightningDestructive | 85735 | 1484.52 MB/s | 49280 | 2 | 15.6× |
| Sonic | 199986 | 636.42 MB/s | 214263 | 15 | 6.7× |
| SonicFastest | 200858 | 633.66 MB/s | 214336 | 15 | 6.6× |
| Easyjson | 255434 | 498.27 MB/s | 122864 | 14 | 5.2× |
| Goccy | 257093 | 495.05 MB/s | 225355 | 884 | 5.2× |
| LightningDecodeAny | 466128 | 203.06 MB/s | 465909 | 9706 | 2.9× |
| JSONV2 | 475127 | 267.88 MB/s | 195130 | 1805 | 2.8× |
| Stdlib | 1335011 | 95.34 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3437398 | 654.87 MB/s | 2532848 | 1143 | 9.4× |
| LightningArena | 3479008 | 647.04 MB/s | 2532849 | 1143 | 9.2× |
| Lightning | 3489920 | 645.01 MB/s | 2532850 | 1143 | 9.2× |
| Sonic | 5146486 | 437.40 MB/s | 4866976 | 2584 | 6.3× |
| SonicFastest | 5168550 | 435.53 MB/s | 4868506 | 2584 | 6.2× |
| LightningDecodeAny | 9465011 | 237.83 MB/s | 6828999 | 223498 | 3.4× |
| Goccy | 12706523 | 177.16 MB/s | 4209941 | 56536 | 2.5× |
| Easyjson | 13408844 | 167.88 MB/s | 3099808 | 2120 | 2.4× |
| JSONV2 | 17044212 | 132.07 MB/s | 3123206 | 3083 | 1.9× |
| Stdlib | 32175744 | 69.96 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 492849 | 548.65 MB/s | 397296 | 567 | 8.5× |
| Lightning | 499462 | 541.39 MB/s | 397297 | 567 | 8.4× |
| LightningDestructive | 508311 | 531.96 MB/s | 397296 | 567 | 8.2× |
| Sonic | 766877 | 352.60 MB/s | 641698 | 1147 | 5.4× |
| SonicFastest | 770746 | 350.83 MB/s | 641728 | 1147 | 5.4× |
| LightningDecodeAny | 1350609 | 200.21 MB/s | 845692 | 29656 | 3.1× |
| Goccy | 1772561 | 152.55 MB/s | 541141 | 8122 | 2.4× |
| Easyjson | 1780163 | 151.90 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2293312 | 117.91 MB/s | 348161 | 1628 | 1.8× |
| Stdlib | 4178122 | 64.72 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1041010 | 1659.16 MB/s | 765560 | 2798 | 16.5× |
| LightningArena | 1076205 | 1604.90 MB/s | 774939 | 2444 | 15.9× |
| Lightning | 1083129 | 1594.64 MB/s | 767828 | 2798 | 15.8× |
| Sonic | 2187012 | 789.76 MB/s | 2693183 | 5547 | 7.8× |
| SonicFastest | 2209987 | 781.54 MB/s | 2693619 | 5547 | 7.8× |
| Goccy | 2636889 | 655.02 MB/s | 2580897 | 14603 | 6.5× |
| LightningDecodeAny | 3902190 | 128.21 MB/s | 4491054 | 67881 | 4.4× |
| Easyjson | 4298535 | 401.81 MB/s | 972032 | 5389 | 4.0× |
| JSONV2 | 4916063 | 351.34 MB/s | 1011615 | 7594 | 3.5× |
| Stdlib | 17146317 | 100.73 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 691 | 2623.48 MB/s | 0 | 0 | 23.8× |
| LightningArena | 693 | 2613.66 MB/s | 0 | 0 | 23.7× |
| LightningDestructive | 771 | 2350.21 MB/s | 0 | 0 | 21.3× |
| Easyjson | 2878 | 629.55 MB/s | 24 | 1 | 5.7× |
| Goccy | 3255 | 556.61 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6368 | 284.54 MB/s | 3346 | 38 | 2.6× |
| Sonic | 6647 | 272.58 MB/s | 3347 | 38 | 2.5× |
| JSONV2 | 8236 | 220.01 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9239 | 196.01 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16437 | 110.24 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 708 | 2560.85 MB/s | 0 | 0 | 23.2× |
| LightningArena | 723 | 2507.44 MB/s | 0 | 0 | 22.7× |
| LightningDestructive | 790 | 2293.13 MB/s | 0 | 0 | 20.8× |
| Easyjson | 2848 | 636.28 MB/s | 24 | 1 | 5.8× |
| Goccy | 3295 | 549.94 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6233 | 290.69 MB/s | 3348 | 38 | 2.6× |
| Sonic | 6462 | 280.42 MB/s | 3348 | 38 | 2.5× |
| JSONV2 | 8073 | 224.45 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9456 | 191.52 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 16436 | 110.24 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 957 | 1893.25 MB/s | 144 | 10 | 17.2× |
| Lightning | 978 | 1853.25 MB/s | 144 | 10 | 16.8× |
| LightningDestructive | 1088 | 1665.36 MB/s | 144 | 10 | 15.1× |
| Easyjson | 3312 | 547.12 MB/s | 144 | 10 | 5.0× |
| Goccy | 3568 | 507.81 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 6356 | 285.08 MB/s | 3368 | 40 | 2.6× |
| Sonic | 6606 | 274.31 MB/s | 3367 | 40 | 2.5× |
| JSONV2 | 8533 | 212.36 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9363 | 193.42 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16422 | 110.34 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 792 | 623.69 MB/s | 160 | 1 | 8.4× |
| LightningDestructive | 807 | 612.19 MB/s | 160 | 1 | 8.2× |
| SonicFastest | 1299 | 380.22 MB/s | 1076 | 8 | 5.1× |
| Sonic | 1308 | 377.55 MB/s | 1076 | 8 | 5.1× |
| LightningDecodeAny | 1388 | 355.19 MB/s | 1040 | 25 | 4.8× |
| LightningArena | 1780 | 277.55 MB/s | 4120 | 2 | 3.7× |
| Easyjson | 2668 | 185.16 MB/s | 448 | 3 | 2.5× |
| Goccy | 2755 | 179.33 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 3505 | 140.92 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6635 | 74.45 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 402 | 572.83 MB/s | 160 | 1 | 11.8× |
| LightningDestructive | 404 | 568.98 MB/s | 160 | 1 | 11.7× |
| Sonic | 939 | 244.86 MB/s | 801 | 8 | 5.0× |
| SonicFastest | 940 | 244.66 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1155 | 198.19 MB/s | 1040 | 25 | 4.1× |
| LightningArena | 1355 | 169.79 MB/s | 4120 | 2 | 3.5× |
| Easyjson | 1806 | 127.37 MB/s | 448 | 3 | 2.6× |
| Goccy | 1873 | 122.80 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2588 | 88.86 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4721 | 48.72 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 60060 | 1084.45 MB/s | 103711 | 99 | 11.4× |
| Lightning | 60721 | 1072.64 MB/s | 103724 | 99 | 11.3× |
| LightningDestructive | 64996 | 1002.09 MB/s | 97220 | 98 | 10.6× |
| Sonic | 149562 | 435.48 MB/s | 235913 | 65 | 4.6× |
| SonicFastest | 158293 | 411.46 MB/s | 236382 | 65 | 4.3× |
| LightningDecodeAny | 195058 | 273.40 MB/s | 176734 | 3237 | 3.5× |
| Goccy | 196890 | 330.80 MB/s | 228697 | 134 | 3.5× |
| JSONV2 | 288416 | 225.83 MB/s | 206663 | 607 | 2.4× |
| Stdlib | 686500 | 94.88 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2135682 | 908.60 MB/s | 2185296 | 1350 | 12.9× |
| Lightning | 2225699 | 871.85 MB/s | 2185298 | 1350 | 12.3× |
| LightningArena | 2260392 | 858.47 MB/s | 2185297 | 1350 | 12.2× |
| Sonic | 4945624 | 392.36 MB/s | 4881427 | 1736 | 5.6× |
| SonicFastest | 4954837 | 391.63 MB/s | 4882576 | 1736 | 5.5× |
| Goccy | 5014296 | 386.99 MB/s | 4063208 | 13509 | 5.5× |
| Easyjson | 8091223 | 239.82 MB/s | 3871265 | 15043 | 3.4× |
| LightningDecodeAny | 9615560 | 201.81 MB/s | 6627984 | 206416 | 2.9× |
| JSONV2 | 13757762 | 141.05 MB/s | 3237179 | 13947 | 2.0× |
| Stdlib | 27468802 | 70.64 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 907280 | 3667.92 MB/s | 351704 | 1286 | 26.3× |
| LightningArena | 1353045 | 2459.51 MB/s | 2434826 | 1413 | 17.7× |
| Lightning | 1362272 | 2442.85 MB/s | 2434802 | 1413 | 17.5× |
| SonicFastest | 2092747 | 1590.17 MB/s | 5896612 | 4263 | 11.4× |
| Sonic | 2104332 | 1581.42 MB/s | 5896575 | 4263 | 11.4× |
| LightningDecodeAny | 3322638 | 925.10 MB/s | 4825765 | 55311 | 7.2× |
| Goccy | 4814428 | 691.22 MB/s | 3948914 | 3816 | 5.0× |
| JSONV2 | 7739382 | 429.99 MB/s | 5364507 | 13243 | 3.1× |
| Stdlib | 23886555 | 139.32 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 199330 | 1105.43 MB/s | 135392 | 226 | 12.1× |
| Lightning | 200611 | 1098.37 MB/s | 135392 | 226 | 12.1× |
| LightningDestructive | 207776 | 1060.50 MB/s | 135392 | 226 | 11.6× |
| Goccy | 488717 | 450.87 MB/s | 364286 | 1066 | 4.9× |
| SonicFastest | 490703 | 449.04 MB/s | 351371 | 262 | 4.9× |
| Sonic | 491461 | 448.35 MB/s | 351203 | 262 | 4.9× |
| Easyjson | 647937 | 340.07 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 747751 | 294.68 MB/s | 129747 | 470 | 3.2× |
| LightningDecodeAny | 1009793 | 107.26 MB/s | 854738 | 11700 | 2.4× |
| Stdlib | 2419147 | 91.08 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 10343786 | 783.08 MB/s | 8109648 | 20809 | 10.4× |
| LightningArena | 10356827 | 782.10 MB/s | 8109648 | 20809 | 10.4× |
| Lightning | 10482565 | 772.72 MB/s | 8109649 | 20809 | 10.3× |
| SonicFastest | 18354826 | 441.30 MB/s | 19863156 | 41640 | 5.9× |
| Sonic | 18439406 | 439.28 MB/s | 19860881 | 41640 | 5.9× |
| Goccy | 26381875 | 307.03 MB/s | 19193869 | 107156 | 4.1× |
| LightningDecodeAny | 34342140 | 151.51 MB/s | 28359828 | 746961 | 3.1× |
| Easyjson | 37148479 | 218.04 MB/s | 15059620 | 41643 | 2.9× |
| JSONV2 | 48439942 | 167.22 MB/s | 15233724 | 78972 | 2.2× |
| Stdlib | 108056355 | 74.96 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 3507990 | 850.48 MB/s | 3780457 | 1514 | 16.2× |
| LightningDestructive | 3699507 | 806.45 MB/s | 3758857 | 29356 | 15.3× |
| Lightning | 3764914 | 792.44 MB/s | 3758859 | 29356 | 15.1× |
| SonicFastest | 9370312 | 318.40 MB/s | 9131498 | 57804 | 6.1× |
| Sonic | 9402292 | 317.31 MB/s | 9131012 | 57804 | 6.0× |
| LightningDecodeAny | 17429536 | 105.23 MB/s | 18225371 | 350883 | 3.3× |
| Goccy | 18980877 | 157.18 MB/s | 9842938 | 273619 | 3.0× |
| Easyjson | 19398580 | 153.80 MB/s | 9479440 | 30115 | 2.9× |
| JSONV2 | 28017020 | 106.49 MB/s | 9257027 | 86278 | 2.0× |
| Stdlib | 56729694 | 52.59 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 690721 | 1047.60 MB/s | 907600 | 3618 | 19.8× |
| LightningArena | 690863 | 1047.38 MB/s | 916256 | 37 | 19.8× |
| Lightning | 797919 | 906.86 MB/s | 907596 | 3618 | 17.2× |
| SonicFastest | 2112278 | 342.57 MB/s | 2372494 | 3683 | 6.5× |
| Sonic | 2128549 | 339.95 MB/s | 2370954 | 3683 | 6.4× |
| LightningDecodeAny | 5130584 | 126.80 MB/s | 5691594 | 76540 | 2.7× |
| Easyjson | 5461839 | 132.48 MB/s | 2847906 | 3698 | 2.5× |
| Goccy | 5569824 | 129.91 MB/s | 2719094 | 80268 | 2.5× |
| JSONV2 | 6500502 | 111.31 MB/s | 2704700 | 7318 | 2.1× |
| Stdlib | 13703744 | 52.80 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1411042 | 1117.86 MB/s | 916256 | 37 | 13.5× |
| LightningDestructive | 1448669 | 1088.83 MB/s | 907600 | 3618 | 13.2× |
| Lightning | 1455121 | 1084.00 MB/s | 907594 | 3618 | 13.1× |
| SonicFastest | 2411667 | 654.05 MB/s | 3223273 | 3683 | 7.9× |
| Sonic | 2424476 | 650.60 MB/s | 3223433 | 3683 | 7.9× |
| LightningDecodeAny | 4386712 | 171.75 MB/s | 5691590 | 76540 | 4.3× |
| Goccy | 6633741 | 237.78 MB/s | 3491619 | 80262 | 2.9× |
| Easyjson | 6694002 | 235.64 MB/s | 2847905 | 3698 | 2.9× |
| JSONV2 | 6962441 | 226.55 MB/s | 2704554 | 7318 | 2.7× |
| Stdlib | 19080491 | 82.67 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 82692 | 1815.46 MB/s | 81920 | 1 | 26.0× |
| Lightning | 84383 | 1779.08 MB/s | 81920 | 1 | 25.5× |
| LightningDestructive | 87602 | 1713.71 MB/s | 81921 | 1 | 24.6× |
| Sonic | 408127 | 367.84 MB/s | 408020 | 16 | 5.3× |
| SonicFastest | 423250 | 354.69 MB/s | 408079 | 16 | 5.1× |
| LightningDecodeAny | 588286 | 255.18 MB/s | 745508 | 10015 | 3.7× |
| Goccy | 994193 | 151.00 MB/s | 327611 | 10005 | 2.2× |
| JSONV2 | 1151541 | 130.37 MB/s | 357723 | 20 | 1.9× |
| Stdlib | 2153611 | 69.71 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 28714 | 979.20 MB/s | 29279 | 101 | 12.0× |
| Lightning | 29266 | 960.74 MB/s | 29279 | 101 | 11.8× |
| LightningDestructive | 29984 | 937.73 MB/s | 29088 | 101 | 11.5× |
| Sonic | 62477 | 450.03 MB/s | 59457 | 83 | 5.5× |
| SonicFastest | 62903 | 446.99 MB/s | 59474 | 83 | 5.5× |
| Easyjson | 79716 | 352.72 MB/s | 32304 | 138 | 4.3× |
| Goccy | 81008 | 347.09 MB/s | 59257 | 188 | 4.3× |
| JSONV2 | 139005 | 202.27 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 163441 | 172.03 MB/s | 134045 | 2639 | 2.1× |
| Stdlib | 345447 | 81.39 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1604 | 1451.73 MB/s | 32 | 1 | 16.2× |
| LightningArena | 1610 | 1446.03 MB/s | 32 | 1 | 16.2× |
| LightningDestructive | 1725 | 1349.23 MB/s | 32 | 1 | 15.1× |
| SonicFastest | 4796 | 485.36 MB/s | 3711 | 4 | 5.4× |
| Sonic | 4826 | 482.43 MB/s | 3711 | 4 | 5.4× |
| Goccy | 4840 | 480.95 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 5083 | 458.04 MB/s | 192 | 2 | 5.1× |
| JSONV2 | 8692 | 267.83 MB/s | 1000 | 6 | 3.0× |
| LightningDecodeAny | 10559 | 159.58 MB/s | 9936 | 194 | 2.5× |
| Stdlib | 26041 | 89.40 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 154 | 1225.16 MB/s | 0 | 0 | 18.0× |
| LightningArena | 157 | 1203.14 MB/s | 0 | 0 | 17.6× |
| LightningDestructive | 164 | 1151.84 MB/s | 0 | 0 | 16.9× |
| Goccy | 462 | 409.02 MB/s | 304 | 2 | 6.0× |
| Easyjson | 590 | 320.33 MB/s | 0 | 0 | 4.7× |
| SonicFastest | 644 | 293.38 MB/s | 341 | 3 | 4.3× |
| Sonic | 648 | 291.68 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1090 | 173.40 MB/s | 112 | 1 | 2.5× |
| LightningDecodeAny | 1350 | 99.27 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2771 | 68.21 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1020 | 2148.40 MB/s | 0 | 0 | 18.6× |
| Lightning | 1024 | 2139.37 MB/s | 0 | 0 | 18.6× |
| LightningDestructive | 1069 | 2049.83 MB/s | 0 | 0 | 17.8× |
| Easyjson | 3520 | 622.45 MB/s | 24 | 1 | 5.4× |
| Goccy | 3654 | 599.61 MB/s | 2864 | 4 | 5.2× |
| SonicFastest | 6856 | 319.57 MB/s | 3602 | 38 | 2.8× |
| Sonic | 7038 | 311.33 MB/s | 3602 | 38 | 2.7× |
| JSONV2 | 8312 | 263.59 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 9254 | 195.69 MB/s | 7552 | 158 | 2.1× |
| Stdlib | 19009 | 115.26 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 588556 | 867.34 MB/s | 318400 | 1005 | 12.0× |
| Lightning | 609495 | 837.54 MB/s | 318400 | 1005 | 11.6× |
| LightningArena | 616794 | 827.63 MB/s | 318400 | 1005 | 11.4× |
| Sonic | 1246044 | 409.68 MB/s | 1307721 | 2014 | 5.7× |
| SonicFastest | 1254913 | 406.78 MB/s | 1308180 | 2014 | 5.6× |
| Goccy | 1380165 | 369.87 MB/s | 1140069 | 5006 | 5.1× |
| Easyjson | 1768527 | 288.64 MB/s | 863779 | 3012 | 4.0× |
| JSONV2 | 3477335 | 146.80 MB/s | 1075955 | 12645 | 2.0× |
| LightningDecodeAny | 3550435 | 129.97 MB/s | 2742393 | 64017 | 2.0× |
| Stdlib | 7056866 | 72.34 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 386 | 51246.39 MB/s | 0 | 0 | 426.5× |
| LightningArena | 387 | 51181.69 MB/s | 0 | 0 | 426.0× |
| LightningDestructive | 626 | 31600.87 MB/s | 0 | 0 | 263.0× |
| SonicFastest | 6949 | 2847.60 MB/s | 21150 | 3 | 23.7× |
| Goccy | 23930 | 826.96 MB/s | 20493 | 2 | 6.9× |
| Sonic | 32050 | 617.43 MB/s | 20608 | 3 | 5.1× |
| JSONV2 | 34430 | 574.76 MB/s | 8 | 1 | 4.8× |
| Easyjson | 97437 | 203.10 MB/s | 0 | 0 | 1.7× |
| LightningDecodeAny | 99148 | 199.58 MB/s | 116608 | 2014 | 1.7× |
| Stdlib | 164701 | 120.15 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1766 | 10262.13 MB/s | 405 | 0 | 70.7× |
| LightningArena | 1781 | 10178.55 MB/s | 405 | 0 | 70.1× |
| LightningDestructive | 2168 | 8360.04 MB/s | 0 | 0 | 57.6× |
| Easyjson | 5056 | 3584.34 MB/s | 432 | 2 | 24.7× |
| Sonic | 8666 | 2091.48 MB/s | 20448 | 5 | 14.4× |
| SonicFastest | 9099 | 1991.91 MB/s | 20430 | 5 | 13.7× |
| LightningDecodeAny | 20657 | 865.66 MB/s | 29135 | 189 | 6.0× |
| Goccy | 26500 | 683.91 MB/s | 19460 | 2 | 4.7× |
| JSONV2 | 52680 | 344.04 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 124895 | 145.11 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2395996 | 838.27 MB/s | 3089565 | 6821 | 9.1× |
| Lightning | 2580921 | 778.21 MB/s | 3097025 | 6822 | 8.4× |
| LightningArena | 2593652 | 774.39 MB/s | 3100034 | 6699 | 8.4× |
| SonicFastest | 4266197 | 470.79 MB/s | 5151419 | 7085 | 5.1× |
| Sonic | 4350904 | 461.63 MB/s | 5152109 | 7085 | 5.0× |
| Goccy | 4891685 | 410.59 MB/s | 5411525 | 15832 | 4.4× |
| Easyjson | 5765756 | 348.35 MB/s | 2981486 | 7439 | 3.8× |
| LightningDecodeAny | 7185502 | 158.97 MB/s | 7377329 | 134004 | 3.0× |
| JSONV2 | 7965362 | 252.15 MB/s | 3173671 | 14562 | 2.7× |
| Stdlib | 21714490 | 92.50 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 932 | 588.83 MB/s | 480 | 1 | 7.2× |
| LightningArena | 941 | 583.45 MB/s | 480 | 1 | 7.2× |
| LightningDestructive | 949 | 578.43 MB/s | 480 | 1 | 7.1× |
| LightningDecodeAny | 1840 | 297.85 MB/s | 1765 | 45 | 3.7× |
| Easyjson | 2256 | 243.30 MB/s | 1616 | 5 | 3.0× |
| SonicFastest | 2403 | 228.43 MB/s | 2261 | 8 | 2.8× |
| Sonic | 2431 | 225.87 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3366 | 163.09 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3575 | 153.58 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6729 | 81.59 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 445433 | 1417.75 MB/s | 400488 | 545 | 14.3× |
| Lightning | 503293 | 1254.77 MB/s | 446877 | 548 | 12.6× |
| LightningArena | 504620 | 1251.46 MB/s | 448674 | 404 | 12.6× |
| Sonic | 1025551 | 615.78 MB/s | 1067812 | 814 | 6.2× |
| SonicFastest | 1027298 | 614.73 MB/s | 1070080 | 814 | 6.2× |
| Easyjson | 1407076 | 448.81 MB/s | 422504 | 936 | 4.5× |
| Goccy | 1413358 | 446.82 MB/s | 990634 | 1200 | 4.5× |
| JSONV2 | 2326268 | 271.47 MB/s | 571592 | 3144 | 2.7× |
| LightningDecodeAny | 2444473 | 191.00 MB/s | 1990691 | 29072 | 2.6× |
| Stdlib | 6353228 | 99.40 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 614500 | 915.23 MB/s | 390531 | 426 | 9.8× |
| LightningArena | 727341 | 773.24 MB/s | 508020 | 287 | 8.3× |
| Lightning | 731767 | 768.56 MB/s | 506285 | 433 | 8.2× |
| SonicFastest | 1303484 | 431.47 MB/s | 1349450 | 1185 | 4.6× |
| Sonic | 1310554 | 429.14 MB/s | 1349375 | 1185 | 4.6× |
| Goccy | 1615769 | 348.07 MB/s | 1040911 | 1029 | 3.7× |
| Easyjson | 2190151 | 256.79 MB/s | 775153 | 1254 | 2.7× |
| LightningDecodeAny | 2662976 | 211.20 MB/s | 1990297 | 28580 | 2.3× |
| JSONV2 | 3278012 | 171.57 MB/s | 927407 | 3482 | 1.8× |
| Stdlib | 6009666 | 93.58 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 683657 | 779.89 MB/s | 333416 | 2084 | 9.4× |
| LightningArena | 737811 | 722.65 MB/s | 367712 | 2086 | 8.7× |
| Lightning | 746402 | 714.33 MB/s | 367708 | 2086 | 8.6× |
| SonicFastest | 1167757 | 456.58 MB/s | 981906 | 3082 | 5.5× |
| Sonic | 1183230 | 450.61 MB/s | 982013 | 3082 | 5.4× |
| Easyjson | 1331109 | 400.55 MB/s | 428362 | 3273 | 4.8× |
| Goccy | 1555699 | 342.73 MB/s | 1167103 | 5409 | 4.1× |
| JSONV2 | 2832524 | 188.23 MB/s | 745421 | 13288 | 2.3× |
| LightningDecodeAny | 3467667 | 153.76 MB/s | 2658098 | 49348 | 1.9× |
| Stdlib | 6416008 | 83.10 MB/s | 798692 | 17133 | 1.0× |
