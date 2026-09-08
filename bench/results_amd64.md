# JSON Deserialization Benchmarks

- generated 2026-09-08T16:56:06Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 77798 | 1635.97 MB/s | 49938 | 2 | 16.8× |
| LightningArena | 79844 | 1594.05 MB/s | 49953 | 2 | 16.4× |
| LightningDestructive | 84290 | 1509.97 MB/s | 49280 | 2 | 15.5× |
| SonicFastest | 194996 | 652.71 MB/s | 214028 | 15 | 6.7× |
| Sonic | 200581 | 634.53 MB/s | 214664 | 15 | 6.5× |
| Easyjson | 231216 | 550.46 MB/s | 122864 | 14 | 5.7× |
| Goccy | 250060 | 508.98 MB/s | 225341 | 884 | 5.2× |
| JSONV2 | 415256 | 306.50 MB/s | 195127 | 1805 | 3.2× |
| LightningDecodeAny | 432713 | 218.74 MB/s | 466196 | 9707 | 3.0× |
| Stdlib | 1309526 | 97.19 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3271712 | 688.03 MB/s | 2532848 | 1143 | 9.4× |
| LightningArena | 3308080 | 680.47 MB/s | 2532849 | 1143 | 9.3× |
| Lightning | 3345718 | 672.82 MB/s | 2532849 | 1143 | 9.2× |
| Sonic | 5723183 | 393.32 MB/s | 4865283 | 2584 | 5.4× |
| SonicFastest | 5831080 | 386.04 MB/s | 4863833 | 2584 | 5.3× |
| LightningDecodeAny | 12329585 | 182.57 MB/s | 19380211 | 223896 | 2.5× |
| Goccy | 12677133 | 177.57 MB/s | 4207017 | 56536 | 2.4× |
| Easyjson | 13716896 | 164.11 MB/s | 3099808 | 2120 | 2.2× |
| JSONV2 | 17527717 | 128.43 MB/s | 3123180 | 3083 | 1.7× |
| Stdlib | 30634706 | 73.48 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 460761 | 586.86 MB/s | 397296 | 567 | 8.7× |
| LightningArena | 462311 | 584.89 MB/s | 397296 | 567 | 8.7× |
| LightningDestructive | 471386 | 573.63 MB/s | 397296 | 567 | 8.5× |
| Sonic | 726755 | 372.07 MB/s | 639483 | 1147 | 5.5× |
| SonicFastest | 741223 | 364.81 MB/s | 640724 | 1147 | 5.4× |
| Goccy | 1702163 | 158.86 MB/s | 545461 | 8123 | 2.4× |
| Easyjson | 1748272 | 154.67 MB/s | 330272 | 749 | 2.3× |
| LightningDecodeAny | 2117173 | 127.72 MB/s | 2543877 | 29687 | 1.9× |
| JSONV2 | 2313557 | 116.88 MB/s | 348160 | 1628 | 1.7× |
| Stdlib | 4002906 | 67.55 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 930837 | 1855.54 MB/s | 765560 | 2798 | 17.9× |
| Lightning | 962057 | 1795.32 MB/s | 767759 | 2798 | 17.3× |
| LightningArena | 963389 | 1792.84 MB/s | 774932 | 2444 | 17.3× |
| Sonic | 2080654 | 830.13 MB/s | 2694021 | 5547 | 8.0× |
| SonicFastest | 2084511 | 828.59 MB/s | 2693367 | 5547 | 8.0× |
| Goccy | 2435692 | 709.12 MB/s | 2581055 | 14603 | 6.8× |
| LightningDecodeAny | 3821335 | 130.92 MB/s | 4963272 | 76577 | 4.4× |
| Easyjson | 3973286 | 434.70 MB/s | 972032 | 5389 | 4.2× |
| JSONV2 | 4220954 | 409.20 MB/s | 1011616 | 7594 | 3.9× |
| Stdlib | 16651426 | 103.73 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 688 | 2631.99 MB/s | 0 | 0 | 23.1× |
| LightningArena | 690 | 2625.82 MB/s | 0 | 0 | 23.0× |
| LightningDestructive | 723 | 2504.72 MB/s | 0 | 0 | 22.0× |
| Easyjson | 2864 | 632.65 MB/s | 24 | 1 | 5.5× |
| Goccy | 3312 | 547.02 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6100 | 297.03 MB/s | 3346 | 38 | 2.6× |
| Sonic | 6312 | 287.09 MB/s | 3346 | 38 | 2.5× |
| JSONV2 | 7810 | 232.00 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8380 | 216.12 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15894 | 114.01 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 736 | 2460.92 MB/s | 0 | 0 | 21.5× |
| LightningArena | 738 | 2455.89 MB/s | 0 | 0 | 21.5× |
| LightningDestructive | 767 | 2363.65 MB/s | 0 | 0 | 20.7× |
| Easyjson | 2867 | 632.01 MB/s | 24 | 1 | 5.5× |
| Goccy | 3320 | 545.81 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6282 | 288.44 MB/s | 3346 | 38 | 2.5× |
| Sonic | 6475 | 279.85 MB/s | 3345 | 38 | 2.4× |
| JSONV2 | 8024 | 225.83 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8390 | 215.86 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15841 | 114.39 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 925 | 1959.56 MB/s | 144 | 10 | 17.1× |
| LightningArena | 939 | 1929.85 MB/s | 144 | 10 | 16.9× |
| LightningDestructive | 983 | 1843.66 MB/s | 144 | 10 | 16.1× |
| Easyjson | 2990 | 606.10 MB/s | 144 | 10 | 5.3× |
| Goccy | 3285 | 551.61 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6496 | 278.96 MB/s | 3368 | 40 | 2.4× |
| Sonic | 6696 | 270.61 MB/s | 3368 | 40 | 2.4× |
| JSONV2 | 7762 | 233.44 MB/s | 632 | 7 | 2.0× |
| LightningDecodeAny | 8426 | 214.93 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15836 | 114.42 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 654 | 755.63 MB/s | 160 | 1 | 9.4× |
| LightningDestructive | 664 | 744.38 MB/s | 160 | 1 | 9.3× |
| Sonic | 1223 | 403.78 MB/s | 1075 | 8 | 5.0× |
| SonicFastest | 1228 | 402.26 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1388 | 355.13 MB/s | 1296 | 26 | 4.4× |
| LightningArena | 1566 | 315.45 MB/s | 4120 | 2 | 3.9× |
| Easyjson | 2281 | 216.57 MB/s | 448 | 3 | 2.7× |
| Goccy | 2561 | 192.86 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 3114 | 158.64 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6139 | 80.46 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 406 | 566.72 MB/s | 160 | 1 | 10.9× |
| Lightning | 407 | 565.28 MB/s | 160 | 1 | 10.8× |
| Sonic | 874 | 263.03 MB/s | 801 | 8 | 5.0× |
| SonicFastest | 878 | 262.05 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1145 | 199.94 MB/s | 1296 | 26 | 3.9× |
| LightningArena | 1309 | 175.67 MB/s | 4120 | 2 | 3.4× |
| Easyjson | 1533 | 149.99 MB/s | 448 | 3 | 2.9× |
| Goccy | 1726 | 133.23 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2444 | 94.10 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4412 | 52.13 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 56553 | 1151.70 MB/s | 103754 | 99 | 11.8× |
| LightningArena | 56829 | 1146.11 MB/s | 103770 | 99 | 11.8× |
| LightningDestructive | 57350 | 1135.69 MB/s | 97220 | 98 | 11.7× |
| Sonic | 155328 | 419.32 MB/s | 235789 | 65 | 4.3× |
| SonicFastest | 157089 | 414.62 MB/s | 235750 | 65 | 4.3× |
| LightningDecodeAny | 179941 | 296.37 MB/s | 180581 | 3241 | 3.7× |
| Goccy | 187403 | 347.55 MB/s | 227708 | 134 | 3.6× |
| JSONV2 | 251312 | 259.17 MB/s | 206664 | 607 | 2.7× |
| Stdlib | 669869 | 97.23 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2179111 | 890.49 MB/s | 2185296 | 1350 | 12.1× |
| Lightning | 2229651 | 870.30 MB/s | 2185297 | 1350 | 11.9× |
| LightningArena | 2238087 | 867.02 MB/s | 2185297 | 1350 | 11.8× |
| Goccy | 5190946 | 373.82 MB/s | 4063576 | 13509 | 5.1× |
| SonicFastest | 6517985 | 297.71 MB/s | 4881765 | 1736 | 4.1× |
| Sonic | 6673075 | 290.79 MB/s | 4881604 | 1736 | 4.0× |
| Easyjson | 7830205 | 247.82 MB/s | 3871266 | 15043 | 3.4× |
| LightningDecodeAny | 9181538 | 211.34 MB/s | 7063041 | 218633 | 2.9× |
| JSONV2 | 11506422 | 168.64 MB/s | 3237179 | 13947 | 2.3× |
| Stdlib | 26458740 | 73.34 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 841415 | 3955.04 MB/s | 351704 | 1286 | 30.7× |
| Lightning | 1286394 | 2586.95 MB/s | 2434771 | 1413 | 20.1× |
| LightningArena | 1294886 | 2569.98 MB/s | 2434825 | 1413 | 20.0× |
| Sonic | 2360695 | 1409.68 MB/s | 5896511 | 4263 | 11.0× |
| SonicFastest | 2361227 | 1409.37 MB/s | 5896387 | 4263 | 11.0× |
| LightningDecodeAny | 3259628 | 942.98 MB/s | 4825517 | 55311 | 7.9× |
| Goccy | 5636541 | 590.40 MB/s | 3948911 | 3816 | 4.6× |
| JSONV2 | 8307010 | 400.61 MB/s | 5364513 | 13243 | 3.1× |
| Stdlib | 25871697 | 128.63 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 169036 | 1303.54 MB/s | 135872 | 226 | 13.7× |
| Lightning | 170714 | 1290.73 MB/s | 135872 | 226 | 13.6× |
| LightningDestructive | 178781 | 1232.49 MB/s | 135872 | 226 | 13.0× |
| Goccy | 462913 | 476.00 MB/s | 364111 | 1066 | 5.0× |
| SonicFastest | 514144 | 428.57 MB/s | 351006 | 262 | 4.5× |
| Sonic | 516025 | 427.01 MB/s | 350997 | 262 | 4.5× |
| Easyjson | 589154 | 374.00 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 714141 | 308.55 MB/s | 129747 | 470 | 3.3× |
| LightningDecodeAny | 942952 | 114.87 MB/s | 897217 | 11703 | 2.5× |
| Stdlib | 2321174 | 94.93 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9599746 | 843.78 MB/s | 8109648 | 20809 | 10.6× |
| LightningArena | 9857292 | 821.73 MB/s | 8109648 | 20809 | 10.3× |
| Lightning | 9859818 | 821.52 MB/s | 8109649 | 20809 | 10.3× |
| Sonic | 20363471 | 397.77 MB/s | 19854268 | 41640 | 5.0× |
| SonicFastest | 20406658 | 396.93 MB/s | 19853815 | 41640 | 5.0× |
| Goccy | 24685143 | 328.13 MB/s | 19156141 | 107156 | 4.1× |
| Easyjson | 34058248 | 237.83 MB/s | 15059620 | 41643 | 3.0× |
| LightningDecodeAny | 37103571 | 140.23 MB/s | 46279351 | 747112 | 2.7× |
| JSONV2 | 45931940 | 176.35 MB/s | 15233723 | 78972 | 2.2× |
| Stdlib | 101338610 | 79.93 MB/s | 15665071 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4686970 | 636.54 MB/s | 3780456 | 1514 | 11.3× |
| Lightning | 4991011 | 597.77 MB/s | 3758856 | 29356 | 10.6× |
| LightningDestructive | 4995878 | 597.19 MB/s | 3758856 | 29356 | 10.6× |
| SonicFastest | 9215278 | 323.75 MB/s | 9130767 | 57804 | 5.7× |
| Sonic | 9221467 | 323.53 MB/s | 9129957 | 57804 | 5.7× |
| LightningDecodeAny | 17500710 | 104.81 MB/s | 23982579 | 351152 | 3.0× |
| Goccy | 18029279 | 165.48 MB/s | 9875703 | 273620 | 2.9× |
| Easyjson | 18054754 | 165.25 MB/s | 9479440 | 30115 | 2.9× |
| JSONV2 | 24864458 | 119.99 MB/s | 9257019 | 86278 | 2.1× |
| Stdlib | 52886039 | 56.41 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1111756 | 650.86 MB/s | 907600 | 3618 | 11.7× |
| LightningArena | 1144608 | 632.18 MB/s | 916258 | 37 | 11.3× |
| Lightning | 1190068 | 608.03 MB/s | 907596 | 3618 | 10.9× |
| SonicFastest | 2107940 | 343.27 MB/s | 2368292 | 3683 | 6.2× |
| Sonic | 2147309 | 336.98 MB/s | 2367839 | 3683 | 6.0× |
| LightningDecodeAny | 5159111 | 126.10 MB/s | 6500460 | 76546 | 2.5× |
| Easyjson | 5316004 | 136.12 MB/s | 2847906 | 3698 | 2.4× |
| Goccy | 5340620 | 135.49 MB/s | 2719958 | 80268 | 2.4× |
| JSONV2 | 6160723 | 117.45 MB/s | 2704708 | 7318 | 2.1× |
| Stdlib | 12978647 | 55.75 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1562367 | 1009.59 MB/s | 907600 | 3618 | 11.5× |
| LightningArena | 1567141 | 1006.52 MB/s | 916257 | 37 | 11.4× |
| Lightning | 1611804 | 978.63 MB/s | 907597 | 3618 | 11.1× |
| SonicFastest | 2549281 | 618.74 MB/s | 3222847 | 3683 | 7.0× |
| Sonic | 2567693 | 614.31 MB/s | 3225357 | 3683 | 7.0× |
| LightningDecodeAny | 4481957 | 168.10 MB/s | 6500457 | 76546 | 4.0× |
| Easyjson | 6397142 | 246.57 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 6624758 | 238.10 MB/s | 3506521 | 80263 | 2.7× |
| JSONV2 | 6815582 | 231.43 MB/s | 2704554 | 7318 | 2.6× |
| Stdlib | 17940699 | 87.92 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 204579 | 733.82 MB/s | 81920 | 1 | 9.9× |
| Lightning | 205470 | 730.64 MB/s | 81920 | 1 | 9.9× |
| LightningDestructive | 209470 | 716.69 MB/s | 81920 | 1 | 9.7× |
| SonicFastest | 382843 | 392.13 MB/s | 407280 | 16 | 5.3× |
| Sonic | 397033 | 378.11 MB/s | 407325 | 16 | 5.1× |
| LightningDecodeAny | 545304 | 275.30 MB/s | 745765 | 10016 | 3.7× |
| Goccy | 985861 | 152.28 MB/s | 325513 | 10005 | 2.1× |
| JSONV2 | 1148619 | 130.70 MB/s | 357724 | 20 | 1.8× |
| Stdlib | 2029693 | 73.96 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 25976 | 1082.44 MB/s | 29278 | 101 | 12.8× |
| Lightning | 26075 | 1078.31 MB/s | 29278 | 101 | 12.8× |
| LightningDestructive | 26992 | 1041.70 MB/s | 29088 | 101 | 12.3× |
| Sonic | 72631 | 387.12 MB/s | 59443 | 83 | 4.6× |
| SonicFastest | 73247 | 383.86 MB/s | 59424 | 83 | 4.5× |
| Easyjson | 76192 | 369.03 MB/s | 32304 | 138 | 4.4× |
| Goccy | 80403 | 349.70 MB/s | 59282 | 188 | 4.1× |
| JSONV2 | 132663 | 211.94 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 151831 | 185.19 MB/s | 141390 | 2641 | 2.2× |
| Stdlib | 333197 | 84.39 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1435 | 1622.44 MB/s | 32 | 1 | 17.4× |
| LightningArena | 1472 | 1581.44 MB/s | 32 | 1 | 17.0× |
| LightningDestructive | 1508 | 1543.97 MB/s | 32 | 1 | 16.6× |
| Goccy | 4761 | 488.94 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 5356 | 434.69 MB/s | 192 | 2 | 4.7× |
| Sonic | 6021 | 386.67 MB/s | 3708 | 4 | 4.2× |
| SonicFastest | 6069 | 383.58 MB/s | 3710 | 4 | 4.1× |
| JSONV2 | 7852 | 296.48 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 9512 | 177.15 MB/s | 10200 | 195 | 2.6× |
| Stdlib | 24993 | 93.15 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 149 | 1264.91 MB/s | 0 | 0 | 17.6× |
| LightningArena | 150 | 1259.07 MB/s | 0 | 0 | 17.5× |
| LightningDestructive | 157 | 1206.13 MB/s | 0 | 0 | 16.7× |
| Goccy | 416 | 453.78 MB/s | 304 | 2 | 6.3× |
| Easyjson | 546 | 345.94 MB/s | 0 | 0 | 4.8× |
| SonicFastest | 771 | 245.14 MB/s | 341 | 3 | 3.4× |
| Sonic | 776 | 243.59 MB/s | 341 | 3 | 3.4× |
| JSONV2 | 966 | 195.60 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1201 | 111.56 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 2624 | 72.03 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 982 | 2229.92 MB/s | 0 | 0 | 19.0× |
| Lightning | 983 | 2229.51 MB/s | 0 | 0 | 19.0× |
| LightningDestructive | 1014 | 2161.02 MB/s | 0 | 0 | 18.4× |
| Easyjson | 3492 | 627.35 MB/s | 24 | 1 | 5.3× |
| Goccy | 3748 | 584.62 MB/s | 2864 | 4 | 5.0× |
| SonicFastest | 6744 | 324.86 MB/s | 3601 | 38 | 2.8× |
| Sonic | 6982 | 313.81 MB/s | 3599 | 38 | 2.7× |
| JSONV2 | 8151 | 268.79 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8395 | 215.71 MB/s | 7552 | 158 | 2.2× |
| Stdlib | 18668 | 117.37 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 470992 | 1083.83 MB/s | 318401 | 1005 | 14.3× |
| LightningArena | 491348 | 1038.93 MB/s | 318400 | 1005 | 13.7× |
| Lightning | 492529 | 1036.44 MB/s | 318400 | 1005 | 13.7× |
| Goccy | 1241076 | 411.32 MB/s | 1135606 | 5006 | 5.4× |
| Sonic | 1562891 | 326.62 MB/s | 1309857 | 2014 | 4.3× |
| Easyjson | 1563144 | 326.57 MB/s | 863782 | 3012 | 4.3× |
| SonicFastest | 1571368 | 324.86 MB/s | 1311157 | 2014 | 4.3× |
| JSONV2 | 3071976 | 166.17 MB/s | 1075950 | 12645 | 2.2× |
| LightningDecodeAny | 3266683 | 141.26 MB/s | 2950651 | 64018 | 2.1× |
| Stdlib | 6746044 | 75.67 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 599 | 33023.09 MB/s | 0 | 0 | 259.8× |
| LightningArena | 601 | 32942.22 MB/s | 0 | 0 | 259.1× |
| LightningDestructive | 846 | 23383.66 MB/s | 0 | 0 | 183.9× |
| SonicFastest | 6889 | 2872.56 MB/s | 21107 | 3 | 22.6× |
| Goccy | 25948 | 762.64 MB/s | 20492 | 2 | 6.0× |
| Sonic | 29190 | 677.93 MB/s | 20631 | 3 | 5.3× |
| JSONV2 | 36209 | 546.52 MB/s | 8 | 1 | 4.3× |
| LightningDecodeAny | 96482 | 205.09 MB/s | 116864 | 2015 | 1.6× |
| Easyjson | 120486 | 164.24 MB/s | 0 | 0 | 1.3× |
| Stdlib | 155662 | 127.13 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1683 | 10770.15 MB/s | 405 | 0 | 80.9× |
| LightningArena | 1716 | 10560.54 MB/s | 405 | 0 | 79.3× |
| LightningDestructive | 1975 | 9177.79 MB/s | 0 | 0 | 68.9× |
| Easyjson | 4687 | 3866.88 MB/s | 432 | 2 | 29.0× |
| SonicFastest | 9563 | 1895.13 MB/s | 20431 | 5 | 14.2× |
| Sonic | 9614 | 1885.09 MB/s | 20427 | 5 | 14.2× |
| LightningDecodeAny | 18219 | 981.49 MB/s | 29136 | 189 | 7.5× |
| Goccy | 26360 | 687.56 MB/s | 19460 | 2 | 5.2× |
| JSONV2 | 49177 | 368.55 MB/s | 16500 | 50 | 2.8× |
| Stdlib | 136126 | 133.14 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2009620 | 999.44 MB/s | 3089565 | 6821 | 10.6× |
| LightningArena | 2102376 | 955.34 MB/s | 3100035 | 6699 | 10.1× |
| Lightning | 2135984 | 940.31 MB/s | 3096638 | 6822 | 10.0× |
| Goccy | 4732467 | 424.41 MB/s | 5410219 | 15832 | 4.5× |
| Sonic | 5133887 | 391.22 MB/s | 5152757 | 7085 | 4.1× |
| SonicFastest | 5251060 | 382.49 MB/s | 5151286 | 7085 | 4.1× |
| Easyjson | 5447475 | 368.70 MB/s | 2981488 | 7439 | 3.9× |
| LightningDecodeAny | 6832730 | 167.18 MB/s | 8515157 | 134005 | 3.1× |
| JSONV2 | 7393110 | 271.67 MB/s | 3173681 | 14563 | 2.9× |
| Stdlib | 21302161 | 94.29 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 799 | 686.91 MB/s | 480 | 1 | 7.9× |
| LightningArena | 810 | 677.54 MB/s | 480 | 1 | 7.8× |
| LightningDestructive | 821 | 668.96 MB/s | 480 | 1 | 7.7× |
| LightningDecodeAny | 1691 | 324.06 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 1983 | 276.92 MB/s | 1616 | 5 | 3.2× |
| SonicFastest | 2153 | 255.05 MB/s | 2263 | 8 | 2.9× |
| Sonic | 2210 | 248.45 MB/s | 2264 | 8 | 2.8× |
| Goccy | 3061 | 179.36 MB/s | 2129 | 43 | 2.1× |
| JSONV2 | 3108 | 176.65 MB/s | 1664 | 7 | 2.0× |
| Stdlib | 6283 | 87.38 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 406205 | 1554.67 MB/s | 402728 | 545 | 15.6× |
| LightningArena | 469440 | 1345.25 MB/s | 450926 | 404 | 13.5× |
| Lightning | 469507 | 1345.06 MB/s | 449095 | 548 | 13.5× |
| Sonic | 1101922 | 573.10 MB/s | 1067385 | 814 | 5.7× |
| SonicFastest | 1122420 | 562.64 MB/s | 1068875 | 814 | 5.6× |
| Easyjson | 1297212 | 486.82 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1393995 | 453.02 MB/s | 989067 | 1200 | 4.5× |
| JSONV2 | 2277873 | 277.24 MB/s | 571589 | 3144 | 2.8× |
| LightningDecodeAny | 2436782 | 191.61 MB/s | 2079734 | 29819 | 2.6× |
| Stdlib | 6327696 | 99.80 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 593415 | 947.75 MB/s | 392417 | 426 | 9.9× |
| LightningArena | 682694 | 823.81 MB/s | 509733 | 287 | 8.6× |
| Lightning | 696659 | 807.29 MB/s | 507977 | 433 | 8.5× |
| Sonic | 1381739 | 407.03 MB/s | 1348777 | 1185 | 4.3× |
| SonicFastest | 1412031 | 398.30 MB/s | 1349308 | 1185 | 4.2× |
| Goccy | 1604503 | 350.52 MB/s | 1041156 | 1029 | 3.7× |
| Easyjson | 2037339 | 276.05 MB/s | 775154 | 1254 | 2.9× |
| LightningDecodeAny | 2598549 | 216.43 MB/s | 2078536 | 29327 | 2.3× |
| JSONV2 | 3069537 | 183.22 MB/s | 927406 | 3482 | 1.9× |
| Stdlib | 5901889 | 95.29 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 525528 | 1014.56 MB/s | 333416 | 2084 | 12.1× |
| Lightning | 570599 | 934.42 MB/s | 367884 | 2086 | 11.2× |
| LightningArena | 574641 | 927.85 MB/s | 367878 | 2086 | 11.1× |
| Easyjson | 1282514 | 415.73 MB/s | 428362 | 3273 | 5.0× |
| Sonic | 1403144 | 379.99 MB/s | 981752 | 3082 | 4.5× |
| SonicFastest | 1412148 | 377.57 MB/s | 982426 | 3082 | 4.5× |
| Goccy | 1561553 | 341.44 MB/s | 1167069 | 5408 | 4.1× |
| JSONV2 | 2721590 | 195.91 MB/s | 745421 | 13288 | 2.3× |
| LightningDecodeAny | 3439858 | 155.00 MB/s | 3000540 | 49872 | 1.9× |
| Stdlib | 6375312 | 83.63 MB/s | 798692 | 17133 | 1.0× |
