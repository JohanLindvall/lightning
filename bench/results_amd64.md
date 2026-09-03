# JSON Deserialization Benchmarks

- generated 2026-09-03T00:51:59Z
- go version go1.26.7 linux/amd64
- cpu: AMD EPYC 9V74 80-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 76775 | 1657.77 MB/s | 49760 | 3 | 17.0× |
| Lightning | 77030 | 1652.28 MB/s | 49760 | 3 | 16.9× |
| LightningDestructive | 78590 | 1619.49 MB/s | 49280 | 2 | 16.6× |
| SonicFastest | 206699 | 615.75 MB/s | 213912 | 15 | 6.3× |
| Sonic | 206896 | 615.16 MB/s | 213926 | 15 | 6.3× |
| Easyjson | 231023 | 550.92 MB/s | 122864 | 14 | 5.6× |
| Goccy | 250300 | 508.49 MB/s | 225478 | 884 | 5.2× |
| JSONV2 | 413217 | 308.01 MB/s | 195129 | 1805 | 3.2× |
| LightningDecodeAny | 418006 | 226.44 MB/s | 463410 | 9708 | 3.1× |
| Stdlib | 1304791 | 97.54 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3266207 | 689.19 MB/s | 2532848 | 1143 | 9.4× |
| LightningArena | 3294882 | 683.20 MB/s | 2532849 | 1143 | 9.3× |
| Lightning | 3312217 | 679.62 MB/s | 2532850 | 1143 | 9.3× |
| Sonic | 5648669 | 398.51 MB/s | 4867458 | 2584 | 5.5× |
| SonicFastest | 5914899 | 380.57 MB/s | 4865497 | 2584 | 5.2× |
| LightningDecodeAny | 12412935 | 181.35 MB/s | 19380211 | 223896 | 2.5× |
| Goccy | 12938927 | 173.98 MB/s | 4139260 | 56532 | 2.4× |
| Easyjson | 13642392 | 165.00 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 16950758 | 132.80 MB/s | 3123182 | 3083 | 1.8× |
| Stdlib | 30805487 | 73.07 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 463481 | 583.42 MB/s | 397296 | 567 | 8.7× |
| LightningArena | 463681 | 583.17 MB/s | 397296 | 567 | 8.7× |
| LightningDestructive | 478195 | 565.47 MB/s | 397297 | 567 | 8.4× |
| SonicFastest | 748352 | 361.33 MB/s | 642354 | 1147 | 5.4× |
| Sonic | 748619 | 361.20 MB/s | 642461 | 1147 | 5.4× |
| Easyjson | 1725064 | 156.75 MB/s | 330272 | 749 | 2.3× |
| Goccy | 1758626 | 153.76 MB/s | 543047 | 8122 | 2.3× |
| LightningDecodeAny | 2005889 | 134.80 MB/s | 2543876 | 29687 | 2.0× |
| JSONV2 | 2215736 | 122.04 MB/s | 348160 | 1628 | 1.8× |
| Stdlib | 4014722 | 67.35 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 943673 | 1830.30 MB/s | 765560 | 2798 | 17.8× |
| LightningArena | 964572 | 1790.64 MB/s | 768416 | 2440 | 17.5× |
| Lightning | 968205 | 1783.92 MB/s | 765601 | 2799 | 17.4× |
| SonicFastest | 2172698 | 794.96 MB/s | 2692750 | 5547 | 7.8× |
| Sonic | 2173528 | 794.65 MB/s | 2692794 | 5547 | 7.7× |
| Goccy | 2421383 | 713.31 MB/s | 2580089 | 14603 | 7.0× |
| LightningDecodeAny | 3818009 | 131.04 MB/s | 4953692 | 76576 | 4.4× |
| Easyjson | 3984829 | 433.44 MB/s | 972032 | 5389 | 4.2× |
| JSONV2 | 4233271 | 408.01 MB/s | 1011612 | 7594 | 4.0× |
| Stdlib | 16842202 | 102.55 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 759 | 2387.57 MB/s | 0 | 0 | 20.8× |
| Lightning | 773 | 2343.76 MB/s | 0 | 0 | 20.4× |
| LightningDestructive | 790 | 2292.69 MB/s | 0 | 0 | 20.0× |
| Easyjson | 2859 | 633.78 MB/s | 24 | 1 | 5.5× |
| Goccy | 3438 | 526.99 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6155 | 294.40 MB/s | 3347 | 38 | 2.6× |
| Sonic | 6426 | 281.96 MB/s | 3346 | 38 | 2.5× |
| JSONV2 | 7901 | 229.35 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8307 | 218.01 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15786 | 114.79 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 799 | 2267.65 MB/s | 0 | 0 | 19.9× |
| LightningArena | 804 | 2254.94 MB/s | 0 | 0 | 19.7× |
| LightningDestructive | 837 | 2164.50 MB/s | 0 | 0 | 18.9× |
| Easyjson | 2855 | 634.58 MB/s | 24 | 1 | 5.6× |
| Goccy | 3331 | 543.92 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6449 | 280.98 MB/s | 3344 | 38 | 2.5× |
| Sonic | 6511 | 278.28 MB/s | 3342 | 38 | 2.4× |
| JSONV2 | 7920 | 228.80 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8479 | 213.58 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15863 | 114.23 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1017 | 1781.47 MB/s | 144 | 10 | 15.5× |
| Lightning | 1028 | 1762.88 MB/s | 144 | 10 | 15.3× |
| LightningDestructive | 1074 | 1687.83 MB/s | 144 | 10 | 14.7× |
| Easyjson | 2945 | 615.29 MB/s | 144 | 10 | 5.4× |
| Goccy | 3113 | 581.99 MB/s | 2600 | 5 | 5.1× |
| SonicFastest | 6343 | 285.65 MB/s | 3366 | 40 | 2.5× |
| Sonic | 6544 | 276.89 MB/s | 3368 | 40 | 2.4× |
| JSONV2 | 7663 | 236.47 MB/s | 632 | 7 | 2.1× |
| LightningDecodeAny | 8388 | 215.90 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 15759 | 114.98 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 657 | 751.54 MB/s | 160 | 1 | 9.4× |
| LightningDestructive | 669 | 738.28 MB/s | 160 | 1 | 9.2× |
| Sonic | 1241 | 398.17 MB/s | 1075 | 8 | 5.0× |
| SonicFastest | 1244 | 397.23 MB/s | 1075 | 8 | 4.9× |
| LightningDecodeAny | 1413 | 348.93 MB/s | 1296 | 26 | 4.4× |
| LightningArena | 1481 | 333.50 MB/s | 4096 | 1 | 4.2× |
| Easyjson | 2394 | 206.36 MB/s | 448 | 3 | 2.6× |
| Goccy | 2432 | 203.13 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3266 | 151.24 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6151 | 80.31 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 402 | 572.91 MB/s | 160 | 1 | 11.0× |
| LightningDestructive | 406 | 566.19 MB/s | 160 | 1 | 10.9× |
| SonicFastest | 881 | 261.14 MB/s | 801 | 8 | 5.0× |
| Sonic | 883 | 260.35 MB/s | 801 | 8 | 5.0× |
| LightningDecodeAny | 1145 | 200.08 MB/s | 1296 | 26 | 3.8× |
| LightningArena | 1242 | 185.19 MB/s | 4096 | 1 | 3.5× |
| Easyjson | 1553 | 148.07 MB/s | 448 | 3 | 2.8× |
| Goccy | 1672 | 137.59 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2483 | 92.62 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4408 | 52.17 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 56109 | 1160.82 MB/s | 103440 | 103 | 12.0× |
| Lightning | 56964 | 1143.38 MB/s | 103440 | 103 | 11.8× |
| LightningDestructive | 58371 | 1115.82 MB/s | 97220 | 98 | 11.5× |
| Sonic | 153578 | 424.10 MB/s | 235766 | 65 | 4.4× |
| SonicFastest | 154152 | 422.52 MB/s | 235810 | 65 | 4.4× |
| LightningDecodeAny | 178298 | 299.10 MB/s | 180048 | 3245 | 3.8× |
| Goccy | 188948 | 344.71 MB/s | 227753 | 134 | 3.6× |
| JSONV2 | 257323 | 253.11 MB/s | 206661 | 607 | 2.6× |
| Stdlib | 673323 | 96.73 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2241948 | 865.53 MB/s | 2864593 | 1380 | 11.8× |
| Lightning | 2296075 | 845.13 MB/s | 2864593 | 1380 | 11.5× |
| LightningArena | 2315219 | 838.14 MB/s | 2864593 | 1380 | 11.4× |
| Goccy | 4982475 | 389.46 MB/s | 4063852 | 13509 | 5.3× |
| SonicFastest | 4990542 | 388.83 MB/s | 4881239 | 1736 | 5.3× |
| Sonic | 5209172 | 372.51 MB/s | 4880655 | 1736 | 5.1× |
| Easyjson | 7848627 | 247.24 MB/s | 3871266 | 15043 | 3.4× |
| LightningDecodeAny | 9107275 | 213.07 MB/s | 7063039 | 218633 | 2.9× |
| JSONV2 | 11277451 | 172.07 MB/s | 3237189 | 13947 | 2.3× |
| Stdlib | 26484232 | 73.27 MB/s | 3551317 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 837446 | 3973.79 MB/s | 351704 | 1286 | 30.7× |
| Lightning | 1378202 | 2414.62 MB/s | 2488905 | 2995 | 18.7× |
| LightningArena | 1388926 | 2395.97 MB/s | 2488905 | 2995 | 18.5× |
| Sonic | 2336175 | 1424.48 MB/s | 5896571 | 4263 | 11.0× |
| SonicFastest | 2338096 | 1423.31 MB/s | 5896529 | 4263 | 11.0× |
| LightningDecodeAny | 3284282 | 935.90 MB/s | 4876913 | 56892 | 7.8× |
| Goccy | 5237514 | 635.38 MB/s | 3948912 | 3816 | 4.9× |
| JSONV2 | 7603979 | 437.64 MB/s | 5364504 | 13243 | 3.4× |
| Stdlib | 25718409 | 129.39 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 168387 | 1308.57 MB/s | 135872 | 226 | 13.6× |
| Lightning | 170267 | 1294.12 MB/s | 135872 | 226 | 13.4× |
| LightningDestructive | 178163 | 1236.77 MB/s | 135872 | 226 | 12.8× |
| Goccy | 457853 | 481.26 MB/s | 364285 | 1066 | 5.0× |
| Sonic | 515513 | 427.43 MB/s | 351282 | 262 | 4.4× |
| SonicFastest | 517339 | 425.92 MB/s | 351408 | 262 | 4.4× |
| Easyjson | 592584 | 371.84 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 683133 | 322.55 MB/s | 129746 | 470 | 3.3× |
| LightningDecodeAny | 934541 | 115.90 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2281805 | 96.57 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9846228 | 822.65 MB/s | 11845073 | 20816 | 10.3× |
| LightningArena | 10415400 | 777.70 MB/s | 11845073 | 20816 | 9.7× |
| Lightning | 10420020 | 777.35 MB/s | 11845078 | 20816 | 9.7× |
| Sonic | 21033380 | 385.10 MB/s | 19851805 | 41640 | 4.8× |
| SonicFastest | 21043525 | 384.92 MB/s | 19852979 | 41640 | 4.8× |
| Goccy | 26090125 | 310.46 MB/s | 19122606 | 107156 | 3.9× |
| Easyjson | 34773130 | 232.94 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 37350971 | 139.30 MB/s | 46279352 | 747112 | 2.7× |
| JSONV2 | 45063193 | 179.75 MB/s | 15233751 | 78972 | 2.3× |
| Stdlib | 101469249 | 79.83 MB/s | 15665074 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4627908 | 644.67 MB/s | 3764712 | 1504 | 11.4× |
| LightningDestructive | 4855361 | 614.47 MB/s | 3758856 | 29356 | 10.9× |
| Lightning | 4963027 | 601.14 MB/s | 3758857 | 29356 | 10.7× |
| Sonic | 9333009 | 319.67 MB/s | 9130662 | 57804 | 5.7× |
| SonicFastest | 9420787 | 316.69 MB/s | 9130209 | 57804 | 5.6× |
| LightningDecodeAny | 17294606 | 106.06 MB/s | 23982579 | 351152 | 3.1× |
| Goccy | 17849808 | 167.14 MB/s | 9878361 | 273619 | 3.0× |
| Easyjson | 17885820 | 166.81 MB/s | 9479440 | 30115 | 3.0× |
| JSONV2 | 25022594 | 119.23 MB/s | 9257052 | 86278 | 2.1× |
| Stdlib | 52867661 | 56.43 MB/s | 9258085 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1142433 | 633.38 MB/s | 907601 | 3618 | 11.4× |
| LightningArena | 1154491 | 626.77 MB/s | 911396 | 30 | 11.2× |
| Lightning | 1217189 | 594.48 MB/s | 907597 | 3618 | 10.7× |
| SonicFastest | 2143949 | 337.51 MB/s | 2368230 | 3683 | 6.1× |
| Sonic | 2148874 | 336.73 MB/s | 2368463 | 3683 | 6.0× |
| LightningDecodeAny | 5223568 | 124.55 MB/s | 6500461 | 76546 | 2.5× |
| Easyjson | 5259637 | 137.58 MB/s | 2847909 | 3698 | 2.5× |
| Goccy | 5404083 | 133.90 MB/s | 2725408 | 80268 | 2.4× |
| JSONV2 | 6594842 | 109.72 MB/s | 2704705 | 7318 | 2.0× |
| Stdlib | 12982980 | 55.73 MB/s | 2704548 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1565711 | 1007.44 MB/s | 907600 | 3618 | 11.5× |
| LightningArena | 1574294 | 1001.94 MB/s | 911393 | 30 | 11.4× |
| Lightning | 1616160 | 975.99 MB/s | 907597 | 3618 | 11.1× |
| Sonic | 2418296 | 652.26 MB/s | 3220631 | 3683 | 7.4× |
| SonicFastest | 2426428 | 650.07 MB/s | 3221375 | 3683 | 7.4× |
| LightningDecodeAny | 4545188 | 165.76 MB/s | 6500457 | 76546 | 4.0× |
| Easyjson | 6461538 | 244.11 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 6511513 | 242.24 MB/s | 3487035 | 80261 | 2.8× |
| JSONV2 | 6685622 | 235.93 MB/s | 2704553 | 7318 | 2.7× |
| Stdlib | 17988078 | 87.69 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 204195 | 735.20 MB/s | 81920 | 1 | 10.0× |
| Lightning | 204289 | 734.86 MB/s | 81920 | 1 | 10.0× |
| LightningDestructive | 209532 | 716.47 MB/s | 81920 | 1 | 9.8× |
| Sonic | 386302 | 388.62 MB/s | 407337 | 16 | 5.3× |
| SonicFastest | 425439 | 352.87 MB/s | 407527 | 16 | 4.8× |
| LightningDecodeAny | 544269 | 275.82 MB/s | 745765 | 10016 | 3.8× |
| Goccy | 995885 | 150.74 MB/s | 330201 | 10005 | 2.1× |
| JSONV2 | 1164801 | 128.88 MB/s | 357724 | 20 | 1.8× |
| Stdlib | 2049296 | 73.26 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 25959 | 1083.11 MB/s | 29216 | 103 | 12.8× |
| LightningArena | 25976 | 1082.43 MB/s | 29216 | 103 | 12.8× |
| LightningDestructive | 27452 | 1024.21 MB/s | 29088 | 101 | 12.1× |
| Sonic | 68484 | 410.56 MB/s | 59445 | 83 | 4.8× |
| SonicFastest | 68693 | 409.31 MB/s | 59516 | 83 | 4.8× |
| Easyjson | 75273 | 373.53 MB/s | 32304 | 138 | 4.4× |
| Goccy | 79760 | 352.52 MB/s | 59276 | 188 | 4.2× |
| JSONV2 | 127909 | 219.82 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 158763 | 177.10 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 332082 | 84.67 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1436 | 1621.59 MB/s | 32 | 1 | 17.3× |
| Lightning | 1459 | 1595.48 MB/s | 32 | 1 | 17.1× |
| LightningDestructive | 1604 | 1451.15 MB/s | 32 | 1 | 15.5× |
| Goccy | 4774 | 487.64 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 5149 | 452.11 MB/s | 192 | 2 | 4.8× |
| SonicFastest | 6179 | 376.77 MB/s | 3710 | 4 | 4.0× |
| Sonic | 6199 | 375.52 MB/s | 3711 | 4 | 4.0× |
| JSONV2 | 7860 | 296.17 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 9552 | 176.40 MB/s | 10200 | 195 | 2.6× |
| Stdlib | 24895 | 93.51 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 153 | 1236.17 MB/s | 0 | 0 | 17.1× |
| LightningArena | 154 | 1226.64 MB/s | 0 | 0 | 17.0× |
| LightningDestructive | 157 | 1204.85 MB/s | 0 | 0 | 16.6× |
| Goccy | 430 | 439.71 MB/s | 304 | 2 | 6.1× |
| Easyjson | 570 | 331.70 MB/s | 0 | 0 | 4.6× |
| Sonic | 756 | 249.98 MB/s | 341 | 3 | 3.5× |
| SonicFastest | 764 | 247.40 MB/s | 341 | 3 | 3.4× |
| JSONV2 | 951 | 198.81 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1185 | 113.09 MB/s | 1160 | 25 | 2.2× |
| Stdlib | 2612 | 72.35 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1068 | 2052.06 MB/s | 0 | 0 | 17.4× |
| LightningArena | 1083 | 2023.57 MB/s | 0 | 0 | 17.2× |
| LightningDestructive | 1123 | 1950.92 MB/s | 0 | 0 | 16.5× |
| Easyjson | 3607 | 607.49 MB/s | 24 | 1 | 5.1× |
| Goccy | 3738 | 586.21 MB/s | 2864 | 4 | 5.0× |
| SonicFastest | 6876 | 318.63 MB/s | 3602 | 38 | 2.7× |
| Sonic | 7149 | 306.48 MB/s | 3601 | 38 | 2.6× |
| JSONV2 | 8162 | 268.45 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 8414 | 215.23 MB/s | 7552 | 158 | 2.2× |
| Stdlib | 18576 | 117.95 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 482627 | 1057.70 MB/s | 457537 | 1009 | 14.0× |
| LightningArena | 494845 | 1031.59 MB/s | 457537 | 1009 | 13.6× |
| Lightning | 496272 | 1028.62 MB/s | 457537 | 1009 | 13.6× |
| Goccy | 1258475 | 405.63 MB/s | 1136350 | 5006 | 5.4× |
| Sonic | 1518235 | 336.23 MB/s | 1310832 | 2014 | 4.4× |
| Easyjson | 1549550 | 329.44 MB/s | 863782 | 3012 | 4.3× |
| SonicFastest | 1551813 | 328.95 MB/s | 1310504 | 2014 | 4.3× |
| JSONV2 | 3174249 | 160.82 MB/s | 1075950 | 12645 | 2.1× |
| LightningDecodeAny | 3358742 | 137.39 MB/s | 2950650 | 64018 | 2.0× |
| Stdlib | 6733949 | 75.81 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 607 | 32593.79 MB/s | 0 | 0 | 258.4× |
| Lightning | 609 | 32484.06 MB/s | 0 | 0 | 257.5× |
| LightningDestructive | 873 | 22672.94 MB/s | 0 | 0 | 179.7× |
| SonicFastest | 6470 | 3058.78 MB/s | 21095 | 3 | 24.2× |
| Goccy | 25735 | 768.95 MB/s | 20492 | 2 | 6.1× |
| Sonic | 29435 | 672.29 MB/s | 20624 | 3 | 5.3× |
| JSONV2 | 36236 | 546.11 MB/s | 8 | 1 | 4.3× |
| LightningDecodeAny | 96774 | 204.48 MB/s | 116864 | 2015 | 1.6× |
| Easyjson | 120310 | 164.48 MB/s | 0 | 0 | 1.3× |
| Stdlib | 156871 | 126.15 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1754 | 10332.56 MB/s | 432 | 2 | 77.5× |
| Lightning | 1763 | 10277.35 MB/s | 432 | 2 | 77.1× |
| LightningDestructive | 1919 | 9445.96 MB/s | 0 | 0 | 70.8× |
| Easyjson | 4715 | 3844.11 MB/s | 432 | 2 | 28.8× |
| SonicFastest | 9279 | 1953.21 MB/s | 20450 | 5 | 14.6× |
| Sonic | 9405 | 1926.99 MB/s | 20433 | 5 | 14.4× |
| LightningDecodeAny | 18093 | 988.34 MB/s | 29088 | 191 | 7.5× |
| Goccy | 26220 | 691.23 MB/s | 19460 | 2 | 5.2× |
| JSONV2 | 48787 | 371.49 MB/s | 16500 | 50 | 2.8× |
| Stdlib | 135862 | 133.40 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1968387 | 1020.38 MB/s | 3089565 | 6821 | 10.8× |
| Lightning | 2050996 | 979.28 MB/s | 3091277 | 6827 | 10.4× |
| LightningArena | 2052815 | 978.41 MB/s | 3094370 | 6703 | 10.4× |
| Goccy | 4746747 | 423.13 MB/s | 5409768 | 15830 | 4.5× |
| SonicFastest | 5180951 | 387.67 MB/s | 5153277 | 7085 | 4.1× |
| Sonic | 5183643 | 387.47 MB/s | 5152792 | 7085 | 4.1× |
| Easyjson | 5612499 | 357.86 MB/s | 2981482 | 7439 | 3.8× |
| LightningDecodeAny | 6593479 | 173.25 MB/s | 8503512 | 134008 | 3.2× |
| JSONV2 | 7405685 | 271.21 MB/s | 3173674 | 14563 | 2.9× |
| Stdlib | 21257339 | 94.48 MB/s | 3589316 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 803 | 683.60 MB/s | 480 | 1 | 7.8× |
| Lightning | 809 | 678.85 MB/s | 480 | 1 | 7.7× |
| LightningDestructive | 819 | 669.99 MB/s | 480 | 1 | 7.6× |
| LightningDecodeAny | 1707 | 321.00 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 1973 | 278.27 MB/s | 1616 | 5 | 3.2× |
| SonicFastest | 2128 | 257.97 MB/s | 2263 | 8 | 2.9× |
| Sonic | 2188 | 250.96 MB/s | 2262 | 8 | 2.9× |
| JSONV2 | 2982 | 184.12 MB/s | 1664 | 7 | 2.1× |
| Goccy | 3010 | 182.41 MB/s | 2129 | 43 | 2.1× |
| Stdlib | 6248 | 87.87 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 402977 | 1567.12 MB/s | 402728 | 545 | 15.5× |
| LightningArena | 481009 | 1312.89 MB/s | 453017 | 712 | 13.0× |
| Lightning | 483571 | 1305.94 MB/s | 451256 | 857 | 12.9× |
| SonicFastest | 1112330 | 567.74 MB/s | 1068375 | 814 | 5.6× |
| Sonic | 1131217 | 558.26 MB/s | 1067747 | 814 | 5.5× |
| Easyjson | 1286161 | 491.01 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1399988 | 451.09 MB/s | 991776 | 1201 | 4.5× |
| JSONV2 | 2194472 | 287.77 MB/s | 571588 | 3144 | 2.9× |
| LightningDecodeAny | 2345042 | 199.10 MB/s | 2076504 | 30126 | 2.7× |
| Stdlib | 6256482 | 100.94 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 588627 | 955.46 MB/s | 546569 | 429 | 10.1× |
| Lightning | 768991 | 731.36 MB/s | 769937 | 1235 | 7.7× |
| LightningArena | 774144 | 726.49 MB/s | 771665 | 1088 | 7.6× |
| SonicFastest | 1414025 | 397.74 MB/s | 1348067 | 1184 | 4.2× |
| Sonic | 1420576 | 395.90 MB/s | 1348460 | 1184 | 4.2× |
| Goccy | 1596373 | 352.30 MB/s | 1044482 | 1028 | 3.7× |
| Easyjson | 2020485 | 278.35 MB/s | 775154 | 1254 | 2.9× |
| LightningDecodeAny | 2771900 | 202.90 MB/s | 2180441 | 30126 | 2.1× |
| JSONV2 | 3070086 | 183.19 MB/s | 927404 | 3482 | 1.9× |
| Stdlib | 5918201 | 95.03 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 522075 | 1021.27 MB/s | 333416 | 2084 | 12.1× |
| Lightning | 598743 | 890.50 MB/s | 368224 | 2293 | 10.5× |
| LightningArena | 604822 | 881.54 MB/s | 368224 | 2293 | 10.4× |
| Easyjson | 1267280 | 420.73 MB/s | 428362 | 3273 | 5.0× |
| SonicFastest | 1431075 | 372.57 MB/s | 982945 | 3082 | 4.4× |
| Sonic | 1436112 | 371.26 MB/s | 982860 | 3082 | 4.4× |
| Goccy | 1539869 | 346.25 MB/s | 1167085 | 5409 | 4.1× |
| JSONV2 | 2701931 | 197.33 MB/s | 745420 | 13288 | 2.3× |
| LightningDecodeAny | 3307850 | 161.19 MB/s | 2992875 | 50076 | 1.9× |
| Stdlib | 6301242 | 84.61 MB/s | 798691 | 17133 | 1.0× |
