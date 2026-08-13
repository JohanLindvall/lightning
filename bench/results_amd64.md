# JSON Deserialization Benchmarks

- generated 2026-08-13T07:30:05Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 98374 | 1293.79 MB/s | 49760 | 3 | 13.4× |
| LightningArena | 98548 | 1291.50 MB/s | 49760 | 3 | 13.4× |
| LightningDestructive | 101283 | 1256.63 MB/s | 49280 | 2 | 13.0× |
| Sonic | 193144 | 658.96 MB/s | 214068 | 15 | 6.8× |
| SonicFastest | 193549 | 657.58 MB/s | 214082 | 15 | 6.8× |
| Goccy | 235767 | 539.83 MB/s | 225254 | 884 | 5.6× |
| Easyjson | 248348 | 512.49 MB/s | 122864 | 14 | 5.3× |
| JSONV2 | 399253 | 318.78 MB/s | 195127 | 1805 | 3.3× |
| LightningDecodeAny | 414210 | 228.51 MB/s | 463410 | 9708 | 3.2× |
| Stdlib | 1316084 | 96.71 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3971904 | 566.74 MB/s | 2532848 | 1143 | 7.8× |
| LightningArena | 4057486 | 554.79 MB/s | 2532849 | 1143 | 7.7× |
| Lightning | 4074679 | 552.45 MB/s | 2532851 | 1143 | 7.6× |
| SonicFastest | 5570391 | 404.11 MB/s | 4866199 | 2584 | 5.6× |
| Sonic | 5681451 | 396.21 MB/s | 4864895 | 2584 | 5.5× |
| Goccy | 12776741 | 176.18 MB/s | 4139505 | 56532 | 2.4× |
| LightningDecodeAny | 13070136 | 172.23 MB/s | 19380210 | 223896 | 2.4× |
| Easyjson | 13652733 | 164.88 MB/s | 3099810 | 2120 | 2.3× |
| JSONV2 | 16774026 | 134.20 MB/s | 3123189 | 3083 | 1.9× |
| Stdlib | 31145570 | 72.28 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 532004 | 508.27 MB/s | 397296 | 567 | 7.6× |
| LightningArena | 534494 | 505.90 MB/s | 397296 | 567 | 7.5× |
| LightningDestructive | 543901 | 497.16 MB/s | 397296 | 567 | 7.4× |
| Sonic | 729941 | 370.44 MB/s | 642373 | 1147 | 5.5× |
| SonicFastest | 742142 | 364.35 MB/s | 644064 | 1147 | 5.4× |
| Easyjson | 1748417 | 154.66 MB/s | 330272 | 749 | 2.3× |
| Goccy | 1748521 | 154.65 MB/s | 544901 | 8123 | 2.3× |
| LightningDecodeAny | 2115169 | 127.84 MB/s | 2543876 | 29687 | 1.9× |
| JSONV2 | 2199522 | 122.94 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4032514 | 67.06 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1090206 | 1584.29 MB/s | 765560 | 2798 | 15.5× |
| LightningArena | 1092268 | 1581.30 MB/s | 768416 | 2440 | 15.5× |
| Lightning | 1115831 | 1547.91 MB/s | 765601 | 2799 | 15.2× |
| Sonic | 2027508 | 851.89 MB/s | 2693228 | 5547 | 8.4× |
| SonicFastest | 2039468 | 846.89 MB/s | 2693963 | 5547 | 8.3× |
| Goccy | 2380940 | 725.43 MB/s | 2581342 | 14603 | 7.1× |
| Easyjson | 3634581 | 475.21 MB/s | 972032 | 5389 | 4.7× |
| LightningDecodeAny | 3709197 | 134.88 MB/s | 4953692 | 76576 | 4.6× |
| JSONV2 | 4169020 | 414.29 MB/s | 1011615 | 7594 | 4.1× |
| Stdlib | 16938982 | 101.97 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1008 | 1797.83 MB/s | 0 | 0 | 15.7× |
| Lightning | 1029 | 1761.53 MB/s | 0 | 0 | 15.4× |
| LightningDestructive | 1061 | 1707.41 MB/s | 0 | 0 | 14.9× |
| Easyjson | 2839 | 638.26 MB/s | 24 | 1 | 5.6× |
| Goccy | 3093 | 585.84 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 6080 | 298.04 MB/s | 3346 | 38 | 2.6× |
| Sonic | 6231 | 290.78 MB/s | 3344 | 38 | 2.5× |
| JSONV2 | 7579 | 239.07 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 8393 | 215.79 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15812 | 114.60 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1062 | 1706.13 MB/s | 0 | 0 | 14.9× |
| Lightning | 1084 | 1671.95 MB/s | 0 | 0 | 14.6× |
| LightningDestructive | 1137 | 1594.13 MB/s | 0 | 0 | 13.9× |
| Easyjson | 2859 | 633.72 MB/s | 24 | 1 | 5.5× |
| Goccy | 3301 | 548.86 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6242 | 290.31 MB/s | 3344 | 38 | 2.5× |
| Sonic | 6419 | 282.31 MB/s | 3344 | 38 | 2.5× |
| JSONV2 | 7459 | 242.94 MB/s | 640 | 6 | 2.1× |
| LightningDecodeAny | 8336 | 217.26 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15788 | 114.77 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1227 | 1477.17 MB/s | 144 | 10 | 12.8× |
| Lightning | 1231 | 1471.91 MB/s | 144 | 10 | 12.8× |
| LightningDestructive | 1272 | 1424.73 MB/s | 144 | 10 | 12.4× |
| Easyjson | 3021 | 599.86 MB/s | 144 | 10 | 5.2× |
| Goccy | 3274 | 553.44 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6309 | 287.21 MB/s | 3366 | 40 | 2.5× |
| Sonic | 6496 | 278.95 MB/s | 3367 | 40 | 2.4× |
| LightningDecodeAny | 8260 | 219.24 MB/s | 7552 | 158 | 1.9× |
| JSONV2 | 8924 | 203.04 MB/s | 632 | 7 | 1.8× |
| Stdlib | 15747 | 115.07 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 703 | 702.89 MB/s | 160 | 1 | 8.7× |
| LightningDestructive | 706 | 699.52 MB/s | 160 | 1 | 8.7× |
| SonicFastest | 1210 | 408.31 MB/s | 1075 | 8 | 5.1× |
| Sonic | 1218 | 405.67 MB/s | 1075 | 8 | 5.0× |
| LightningDecodeAny | 1391 | 354.30 MB/s | 1296 | 26 | 4.4× |
| LightningArena | 1525 | 324.03 MB/s | 4096 | 1 | 4.0× |
| Easyjson | 2382 | 207.38 MB/s | 448 | 3 | 2.6× |
| Goccy | 2407 | 205.23 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3141 | 157.26 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6123 | 80.68 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 433 | 531.06 MB/s | 160 | 1 | 10.1× |
| LightningDestructive | 437 | 525.90 MB/s | 160 | 1 | 10.0× |
| SonicFastest | 854 | 269.36 MB/s | 801 | 8 | 5.1× |
| Sonic | 856 | 268.86 MB/s | 801 | 8 | 5.1× |
| LightningDecodeAny | 1164 | 196.72 MB/s | 1296 | 26 | 3.7× |
| LightningArena | 1251 | 183.93 MB/s | 4096 | 1 | 3.5× |
| Easyjson | 1515 | 151.80 MB/s | 448 | 3 | 2.9× |
| Goccy | 1650 | 139.42 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2352 | 97.79 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4355 | 52.81 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 61519 | 1058.73 MB/s | 103441 | 103 | 10.7× |
| Lightning | 61966 | 1051.09 MB/s | 103441 | 103 | 10.7× |
| LightningDestructive | 62755 | 1037.88 MB/s | 97220 | 98 | 10.5× |
| SonicFastest | 150204 | 433.62 MB/s | 235994 | 65 | 4.4× |
| Sonic | 150750 | 432.05 MB/s | 235949 | 65 | 4.4× |
| Goccy | 172084 | 378.49 MB/s | 228426 | 134 | 3.8× |
| LightningDecodeAny | 177658 | 300.18 MB/s | 180048 | 3245 | 3.7× |
| JSONV2 | 238721 | 272.84 MB/s | 206663 | 607 | 2.8× |
| Stdlib | 660700 | 98.58 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2503612 | 775.07 MB/s | 2864592 | 1380 | 10.4× |
| LightningArena | 2531309 | 766.59 MB/s | 2864593 | 1380 | 10.3× |
| Lightning | 2555459 | 759.34 MB/s | 2864593 | 1380 | 10.2× |
| SonicFastest | 4791393 | 404.99 MB/s | 4879357 | 1736 | 5.5× |
| Sonic | 4808201 | 403.58 MB/s | 4879581 | 1736 | 5.4× |
| Goccy | 4998339 | 388.22 MB/s | 4063110 | 13509 | 5.2× |
| Easyjson | 7548352 | 257.07 MB/s | 3871265 | 15043 | 3.5× |
| LightningDecodeAny | 9027602 | 214.95 MB/s | 7063040 | 218633 | 2.9× |
| JSONV2 | 11424301 | 169.85 MB/s | 3237190 | 13947 | 2.3× |
| Stdlib | 26126174 | 74.27 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 968550 | 3435.89 MB/s | 351704 | 1286 | 26.2× |
| LightningArena | 1470911 | 2262.43 MB/s | 2488906 | 2995 | 17.2× |
| Lightning | 1475661 | 2255.15 MB/s | 2488906 | 2995 | 17.2× |
| SonicFastest | 2262451 | 1470.90 MB/s | 5896394 | 4263 | 11.2× |
| Sonic | 2276822 | 1461.61 MB/s | 5896139 | 4263 | 11.1× |
| LightningDecodeAny | 3135130 | 980.43 MB/s | 4876912 | 56892 | 8.1× |
| Goccy | 5588772 | 595.45 MB/s | 3948913 | 3816 | 4.5× |
| JSONV2 | 7844874 | 424.20 MB/s | 5364507 | 13243 | 3.2× |
| Stdlib | 25344471 | 131.30 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 208430 | 1057.17 MB/s | 135872 | 226 | 11.0× |
| Lightning | 209676 | 1050.89 MB/s | 135872 | 226 | 10.9× |
| LightningDestructive | 216428 | 1018.11 MB/s | 135872 | 226 | 10.6× |
| Goccy | 455670 | 483.57 MB/s | 363985 | 1066 | 5.0× |
| SonicFastest | 499726 | 440.93 MB/s | 351284 | 262 | 4.6× |
| Sonic | 500466 | 440.28 MB/s | 351113 | 262 | 4.6× |
| Easyjson | 602256 | 365.87 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 672535 | 327.64 MB/s | 129746 | 470 | 3.4× |
| LightningDecodeAny | 932591 | 116.14 MB/s | 897217 | 11703 | 2.5× |
| Stdlib | 2289754 | 96.23 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11703635 | 692.10 MB/s | 11845074 | 20816 | 8.7× |
| LightningArena | 12069075 | 671.14 MB/s | 11845074 | 20816 | 8.4× |
| Lightning | 12272805 | 660.00 MB/s | 11845079 | 20816 | 8.3× |
| Sonic | 19742072 | 410.29 MB/s | 19852809 | 41640 | 5.1× |
| SonicFastest | 19778833 | 409.53 MB/s | 19852290 | 41640 | 5.1× |
| Goccy | 25366168 | 319.32 MB/s | 19054409 | 107155 | 4.0× |
| Easyjson | 34411338 | 235.39 MB/s | 15059617 | 41643 | 3.0× |
| LightningDecodeAny | 37570734 | 138.49 MB/s | 46279351 | 747112 | 2.7× |
| JSONV2 | 45984058 | 176.15 MB/s | 15233723 | 78972 | 2.2× |
| Stdlib | 101617436 | 79.71 MB/s | 15665073 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5436957 | 548.74 MB/s | 3764713 | 1504 | 9.5× |
| LightningDestructive | 5603542 | 532.42 MB/s | 3758856 | 29356 | 9.2× |
| Lightning | 5829574 | 511.78 MB/s | 3758859 | 29356 | 8.9× |
| Sonic | 9026504 | 330.52 MB/s | 9130636 | 57804 | 5.7× |
| SonicFastest | 9086337 | 328.35 MB/s | 9130791 | 57804 | 5.7× |
| Goccy | 17544428 | 170.05 MB/s | 9882974 | 273620 | 2.9× |
| Easyjson | 17688203 | 168.67 MB/s | 9479440 | 30115 | 2.9× |
| LightningDecodeAny | 17723302 | 103.49 MB/s | 23982580 | 351152 | 2.9× |
| JSONV2 | 24311675 | 122.72 MB/s | 9257095 | 86278 | 2.1× |
| Stdlib | 51635472 | 57.78 MB/s | 9258086 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1269799 | 569.85 MB/s | 907601 | 3618 | 10.0× |
| LightningArena | 1278818 | 565.83 MB/s | 911393 | 30 | 9.9× |
| Lightning | 1335602 | 541.78 MB/s | 907595 | 3618 | 9.5× |
| SonicFastest | 2185295 | 331.12 MB/s | 2367814 | 3683 | 5.8× |
| Sonic | 2189939 | 330.42 MB/s | 2367645 | 3683 | 5.8× |
| Easyjson | 5081539 | 142.40 MB/s | 2847905 | 3698 | 2.5× |
| Goccy | 5301955 | 136.48 MB/s | 2694903 | 80266 | 2.4× |
| LightningDecodeAny | 5316676 | 122.36 MB/s | 6500461 | 76546 | 2.4× |
| JSONV2 | 6049692 | 119.61 MB/s | 2704706 | 7318 | 2.1× |
| Stdlib | 12716424 | 56.90 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1871576 | 842.79 MB/s | 911392 | 30 | 9.6× |
| LightningDestructive | 1917840 | 822.46 MB/s | 907600 | 3618 | 9.3× |
| Lightning | 1918298 | 822.27 MB/s | 907595 | 3618 | 9.3× |
| Sonic | 2424470 | 650.60 MB/s | 3221188 | 3683 | 7.4× |
| SonicFastest | 2527658 | 624.04 MB/s | 3225198 | 3683 | 7.1× |
| LightningDecodeAny | 4617398 | 163.17 MB/s | 6500460 | 76546 | 3.9× |
| Easyjson | 6453022 | 244.44 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 6513828 | 242.15 MB/s | 3490449 | 80261 | 2.8× |
| JSONV2 | 6721257 | 234.68 MB/s | 2704553 | 7318 | 2.7× |
| Stdlib | 17916728 | 88.04 MB/s | 2704549 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 216147 | 694.55 MB/s | 81920 | 1 | 9.4× |
| Lightning | 218152 | 688.16 MB/s | 81920 | 1 | 9.3× |
| LightningDestructive | 225785 | 664.90 MB/s | 81920 | 1 | 9.0× |
| Sonic | 390739 | 384.20 MB/s | 407080 | 16 | 5.2× |
| SonicFastest | 402448 | 373.03 MB/s | 406986 | 16 | 5.0× |
| LightningDecodeAny | 564468 | 265.95 MB/s | 745765 | 10016 | 3.6× |
| Goccy | 1031326 | 145.56 MB/s | 330952 | 10005 | 2.0× |
| JSONV2 | 1142184 | 131.44 MB/s | 357724 | 20 | 1.8× |
| Stdlib | 2023032 | 74.21 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 29884 | 940.86 MB/s | 29216 | 103 | 11.0× |
| Lightning | 30048 | 935.73 MB/s | 29216 | 103 | 10.9× |
| LightningDestructive | 30912 | 909.57 MB/s | 29088 | 101 | 10.6× |
| Sonic | 66168 | 424.93 MB/s | 59436 | 83 | 5.0× |
| SonicFastest | 66436 | 423.22 MB/s | 59422 | 83 | 4.9× |
| Goccy | 74079 | 379.55 MB/s | 59269 | 188 | 4.4× |
| Easyjson | 75907 | 370.41 MB/s | 32304 | 138 | 4.3× |
| JSONV2 | 125656 | 223.76 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 147368 | 190.79 MB/s | 140576 | 2643 | 2.2× |
| Stdlib | 328690 | 85.54 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1893 | 1229.77 MB/s | 32 | 1 | 13.2× |
| Lightning | 1902 | 1223.69 MB/s | 32 | 1 | 13.1× |
| LightningDestructive | 2031 | 1146.48 MB/s | 32 | 1 | 12.3× |
| Goccy | 4589 | 507.31 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4826 | 482.38 MB/s | 192 | 2 | 5.2× |
| Sonic | 6087 | 382.45 MB/s | 3708 | 4 | 4.1× |
| SonicFastest | 6107 | 381.17 MB/s | 3706 | 4 | 4.1× |
| JSONV2 | 7904 | 294.53 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 9458 | 178.15 MB/s | 10200 | 195 | 2.6× |
| Stdlib | 24953 | 93.29 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 207 | 911.78 MB/s | 0 | 0 | 12.7× |
| Lightning | 208 | 909.99 MB/s | 0 | 0 | 12.6× |
| LightningDestructive | 212 | 893.54 MB/s | 0 | 0 | 12.4× |
| Goccy | 414 | 456.77 MB/s | 304 | 2 | 6.3× |
| Easyjson | 524 | 360.96 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 759 | 249.08 MB/s | 341 | 3 | 3.5× |
| Sonic | 761 | 248.23 MB/s | 341 | 3 | 3.4× |
| JSONV2 | 952 | 198.44 MB/s | 112 | 1 | 2.8× |
| LightningDecodeAny | 1174 | 114.17 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 2623 | 72.05 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1399 | 1566.45 MB/s | 0 | 0 | 13.2× |
| LightningArena | 1406 | 1557.78 MB/s | 0 | 0 | 13.1× |
| LightningDestructive | 1441 | 1520.75 MB/s | 0 | 0 | 12.8× |
| Goccy | 3472 | 631.00 MB/s | 2864 | 4 | 5.3× |
| Easyjson | 3507 | 624.68 MB/s | 24 | 1 | 5.3× |
| SonicFastest | 6798 | 322.31 MB/s | 3601 | 38 | 2.7× |
| Sonic | 6988 | 313.52 MB/s | 3600 | 38 | 2.6× |
| JSONV2 | 7661 | 286.00 MB/s | 640 | 6 | 2.4× |
| LightningDecodeAny | 8344 | 217.04 MB/s | 7552 | 158 | 2.2× |
| Stdlib | 18420 | 118.95 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 576962 | 884.77 MB/s | 457537 | 1009 | 11.6× |
| Lightning | 619772 | 823.65 MB/s | 457537 | 1009 | 10.8× |
| LightningArena | 625400 | 816.24 MB/s | 457537 | 1009 | 10.7× |
| Goccy | 1238103 | 412.30 MB/s | 1136597 | 5006 | 5.4× |
| SonicFastest | 1604061 | 318.24 MB/s | 1308104 | 2014 | 4.2× |
| Sonic | 1612428 | 316.59 MB/s | 1308866 | 2014 | 4.2× |
| Easyjson | 1617732 | 315.55 MB/s | 863783 | 3012 | 4.1× |
| JSONV2 | 3069848 | 166.29 MB/s | 1075947 | 12645 | 2.2× |
| LightningDecodeAny | 3372759 | 136.82 MB/s | 2950653 | 64018 | 2.0× |
| Stdlib | 6694728 | 76.25 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 611 | 32409.42 MB/s | 0 | 0 | 255.4× |
| Lightning | 613 | 32268.70 MB/s | 0 | 0 | 254.2× |
| LightningDestructive | 871 | 22722.83 MB/s | 0 | 0 | 179.0× |
| SonicFastest | 6243 | 3169.83 MB/s | 21124 | 3 | 25.0× |
| Goccy | 25267 | 783.20 MB/s | 20492 | 2 | 6.2× |
| Sonic | 28797 | 687.19 MB/s | 20605 | 3 | 5.4× |
| JSONV2 | 36229 | 546.23 MB/s | 8 | 1 | 4.3× |
| LightningDecodeAny | 90673 | 218.24 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 126797 | 156.07 MB/s | 0 | 0 | 1.2× |
| Stdlib | 155931 | 126.91 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 2131 | 8506.86 MB/s | 432 | 2 | 62.7× |
| Lightning | 2157 | 8401.25 MB/s | 432 | 2 | 61.9× |
| LightningDestructive | 2293 | 7905.27 MB/s | 0 | 0 | 58.2× |
| Easyjson | 4483 | 4042.86 MB/s | 432 | 2 | 29.8× |
| Sonic | 9130 | 1985.07 MB/s | 20455 | 5 | 14.6× |
| SonicFastest | 10268 | 1765.10 MB/s | 20384 | 5 | 13.0× |
| LightningDecodeAny | 17774 | 1006.09 MB/s | 29088 | 191 | 7.5× |
| Goccy | 26704 | 678.69 MB/s | 19460 | 2 | 5.0× |
| JSONV2 | 48517 | 373.56 MB/s | 16500 | 50 | 2.8× |
| Stdlib | 133566 | 135.69 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2190675 | 916.84 MB/s | 3089564 | 6821 | 9.5× |
| LightningArena | 2295742 | 874.88 MB/s | 3094370 | 6703 | 9.1× |
| Lightning | 2310289 | 869.37 MB/s | 3091277 | 6827 | 9.0× |
| Goccy | 4544111 | 442.00 MB/s | 5411772 | 15833 | 4.6× |
| SonicFastest | 4864199 | 412.91 MB/s | 5153114 | 7085 | 4.3× |
| Sonic | 4914942 | 408.65 MB/s | 5153651 | 7085 | 4.2× |
| Easyjson | 5009369 | 400.95 MB/s | 2981490 | 7439 | 4.2× |
| LightningDecodeAny | 6519615 | 175.21 MB/s | 8503512 | 134008 | 3.2× |
| JSONV2 | 7144843 | 281.11 MB/s | 3173675 | 14563 | 2.9× |
| Stdlib | 20855103 | 96.31 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 808 | 679.78 MB/s | 480 | 1 | 7.6× |
| Lightning | 813 | 675.33 MB/s | 480 | 1 | 7.5× |
| LightningDestructive | 821 | 668.93 MB/s | 480 | 1 | 7.5× |
| LightningDecodeAny | 1661 | 329.86 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 1955 | 280.75 MB/s | 1616 | 5 | 3.1× |
| SonicFastest | 2046 | 268.29 MB/s | 2262 | 8 | 3.0× |
| Sonic | 2115 | 259.58 MB/s | 2262 | 8 | 2.9× |
| Goccy | 2926 | 187.61 MB/s | 2129 | 43 | 2.1× |
| JSONV2 | 2962 | 185.37 MB/s | 1664 | 7 | 2.1× |
| Stdlib | 6128 | 89.59 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 461210 | 1369.25 MB/s | 402728 | 545 | 13.3× |
| LightningArena | 535229 | 1179.90 MB/s | 453017 | 712 | 11.5× |
| Lightning | 543459 | 1162.03 MB/s | 451257 | 857 | 11.3× |
| Sonic | 1046922 | 603.21 MB/s | 1065563 | 814 | 5.9× |
| SonicFastest | 1068030 | 591.29 MB/s | 1066777 | 814 | 5.7× |
| Easyjson | 1226756 | 514.78 MB/s | 422504 | 936 | 5.0× |
| Goccy | 1366550 | 462.12 MB/s | 989370 | 1200 | 4.5× |
| JSONV2 | 2168037 | 291.28 MB/s | 571591 | 3144 | 2.8× |
| LightningDecodeAny | 2304710 | 202.59 MB/s | 2076503 | 30126 | 2.7× |
| Stdlib | 6132994 | 102.97 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 611755 | 919.33 MB/s | 546569 | 429 | 9.6× |
| LightningArena | 793790 | 708.51 MB/s | 771665 | 1088 | 7.4× |
| Lightning | 799286 | 703.64 MB/s | 769937 | 1235 | 7.3× |
| SonicFastest | 1379068 | 407.82 MB/s | 1347834 | 1184 | 4.2× |
| Sonic | 1379521 | 407.68 MB/s | 1347423 | 1184 | 4.2× |
| Goccy | 1567386 | 358.82 MB/s | 1040047 | 1028 | 3.7× |
| Easyjson | 2026430 | 277.54 MB/s | 775155 | 1254 | 2.9× |
| LightningDecodeAny | 2735974 | 205.56 MB/s | 2180441 | 30126 | 2.1× |
| JSONV2 | 3074441 | 182.93 MB/s | 927403 | 3482 | 1.9× |
| Stdlib | 5847218 | 96.18 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 580503 | 918.48 MB/s | 333416 | 2084 | 10.7× |
| LightningArena | 649756 | 820.58 MB/s | 368224 | 2293 | 9.5× |
| Lightning | 651499 | 818.39 MB/s | 368224 | 2293 | 9.5× |
| Easyjson | 1256371 | 424.38 MB/s | 428362 | 3273 | 4.9× |
| Sonic | 1355200 | 393.43 MB/s | 979416 | 3082 | 4.6× |
| SonicFastest | 1362305 | 391.38 MB/s | 979491 | 3082 | 4.6× |
| Goccy | 1417636 | 376.10 MB/s | 1167080 | 5408 | 4.4× |
| JSONV2 | 2648843 | 201.29 MB/s | 745423 | 13288 | 2.3× |
| LightningDecodeAny | 3297476 | 161.69 MB/s | 2992876 | 50076 | 1.9× |
| Stdlib | 6203372 | 85.95 MB/s | 798692 | 17133 | 1.0× |
