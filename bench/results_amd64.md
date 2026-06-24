# JSON Deserialization Benchmarks

- generated 2026-06-24T16:33:40Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 104753 | 1215.00 MB/s | 49760 | 3 | 12.7× |
| LightningDestructive | 106519 | 1194.86 MB/s | 49280 | 2 | 12.5× |
| SonicFastest | 197219 | 645.35 MB/s | 214263 | 15 | 6.7× |
| Sonic | 198198 | 642.16 MB/s | 214372 | 15 | 6.7× |
| Goccy | 241966 | 526.00 MB/s | 224910 | 884 | 5.5× |
| Easyjson | 254058 | 500.97 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 433027 | 293.92 MB/s | 195128 | 1805 | 3.1× |
| LightningDecodeAny | 475778 | 198.94 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1328304 | 95.82 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4161948 | 540.86 MB/s | 3122872 | 3081 | 8.2× |
| Lightning | 4190519 | 537.18 MB/s | 3122875 | 3081 | 8.1× |
| Sonic | 4994675 | 450.69 MB/s | 4869605 | 2584 | 6.8× |
| SonicFastest | 5243273 | 429.32 MB/s | 4873538 | 2584 | 6.5× |
| Goccy | 11841944 | 190.09 MB/s | 4136729 | 56532 | 2.9× |
| LightningDecodeAny | 12361489 | 182.10 MB/s | 7938299 | 281383 | 2.8× |
| Easyjson | 13986806 | 160.94 MB/s | 3099809 | 2120 | 2.4× |
| JSONV2 | 17490379 | 128.70 MB/s | 3123206 | 3083 | 2.0× |
| Stdlib | 34117565 | 65.98 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 569892 | 474.48 MB/s | 348024 | 1627 | 7.7× |
| LightningDestructive | 578658 | 467.29 MB/s | 348025 | 1627 | 7.6× |
| Sonic | 753142 | 359.03 MB/s | 642408 | 1147 | 5.8× |
| SonicFastest | 766303 | 352.87 MB/s | 643456 | 1147 | 5.7× |
| Goccy | 1674090 | 161.52 MB/s | 541791 | 8122 | 2.6× |
| LightningDecodeAny | 1713007 | 157.85 MB/s | 1011488 | 37901 | 2.6× |
| Easyjson | 1819046 | 148.65 MB/s | 330273 | 749 | 2.4× |
| JSONV2 | 2359559 | 114.60 MB/s | 348160 | 1628 | 1.9× |
| Stdlib | 4399512 | 61.46 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1536281 | 1124.28 MB/s | 519272 | 2692 | 11.5× |
| Lightning | 1572966 | 1098.06 MB/s | 519313 | 2693 | 11.3× |
| SonicFastest | 2148563 | 803.89 MB/s | 2693589 | 5547 | 8.3× |
| Sonic | 2157733 | 800.47 MB/s | 2694059 | 5547 | 8.2× |
| Goccy | 2905029 | 594.56 MB/s | 2581192 | 14603 | 6.1× |
| Easyjson | 4254987 | 405.92 MB/s | 972032 | 5389 | 4.2× |
| LightningDecodeAny | 4464690 | 112.06 MB/s | 4976203 | 81466 | 4.0× |
| JSONV2 | 4972260 | 347.37 MB/s | 1011614 | 7594 | 3.6× |
| Stdlib | 17739901 | 97.36 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1023 | 1771.33 MB/s | 0 | 0 | 15.6× |
| LightningDestructive | 1089 | 1664.10 MB/s | 0 | 0 | 14.7× |
| Easyjson | 3025 | 599.02 MB/s | 24 | 1 | 5.3× |
| Goccy | 3226 | 561.60 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6308 | 287.27 MB/s | 3348 | 38 | 2.5× |
| Sonic | 6483 | 279.50 MB/s | 3346 | 38 | 2.5× |
| JSONV2 | 8036 | 225.49 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9369 | 193.31 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 15984 | 113.37 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1115 | 1624.89 MB/s | 0 | 0 | 14.4× |
| LightningDestructive | 1198 | 1512.25 MB/s | 0 | 0 | 13.4× |
| Easyjson | 2961 | 611.93 MB/s | 24 | 1 | 5.4× |
| Goccy | 3183 | 569.23 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5932 | 305.48 MB/s | 3342 | 38 | 2.7× |
| Sonic | 6137 | 295.28 MB/s | 3343 | 38 | 2.6× |
| JSONV2 | 7988 | 226.84 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9172 | 197.44 MB/s | 7536 | 158 | 1.8× |
| Stdlib | 16090 | 112.61 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1316 | 1377.04 MB/s | 144 | 10 | 12.2× |
| LightningDestructive | 1408 | 1287.19 MB/s | 144 | 10 | 11.4× |
| Easyjson | 3101 | 584.31 MB/s | 144 | 10 | 5.2× |
| Goccy | 3499 | 517.90 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 6317 | 286.85 MB/s | 3364 | 40 | 2.5× |
| Sonic | 6559 | 276.28 MB/s | 3366 | 40 | 2.5× |
| JSONV2 | 8670 | 208.99 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9345 | 193.80 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16080 | 112.69 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 726 | 680.92 MB/s | 160 | 1 | 9.6× |
| LightningDestructive | 739 | 668.80 MB/s | 160 | 1 | 9.4× |
| Sonic | 1273 | 388.07 MB/s | 1076 | 8 | 5.5× |
| SonicFastest | 1280 | 385.87 MB/s | 1075 | 8 | 5.4× |
| LightningDecodeAny | 1801 | 273.80 MB/s | 1536 | 30 | 3.9× |
| Goccy | 2602 | 189.85 MB/s | 856 | 23 | 2.7× |
| Easyjson | 2678 | 184.48 MB/s | 448 | 3 | 2.6× |
| JSONV2 | 3551 | 139.12 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6950 | 71.08 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 496 | 463.70 MB/s | 160 | 1 | 10.0× |
| LightningDestructive | 497 | 462.56 MB/s | 160 | 1 | 10.0× |
| Sonic | 922 | 249.37 MB/s | 801 | 8 | 5.4× |
| SonicFastest | 923 | 249.22 MB/s | 800 | 8 | 5.4× |
| LightningDecodeAny | 1599 | 143.20 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1724 | 133.40 MB/s | 448 | 3 | 2.9× |
| Goccy | 1842 | 124.87 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2544 | 90.42 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4949 | 46.47 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 89728 | 725.88 MB/s | 172434 | 107 | 7.6× |
| LightningDestructive | 95922 | 679.01 MB/s | 166213 | 102 | 7.1× |
| SonicFastest | 146296 | 445.21 MB/s | 235888 | 65 | 4.6× |
| Sonic | 146382 | 444.94 MB/s | 235857 | 65 | 4.6× |
| Goccy | 178826 | 364.22 MB/s | 228303 | 134 | 3.8× |
| LightningDecodeAny | 202400 | 263.48 MB/s | 176961 | 3252 | 3.4× |
| JSONV2 | 264295 | 246.44 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 678741 | 95.96 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2528983 | 767.29 MB/s | 2846866 | 2238 | 11.1× |
| Lightning | 2604726 | 744.98 MB/s | 2846865 | 2238 | 10.8× |
| Sonic | 4743884 | 409.05 MB/s | 4879534 | 1736 | 5.9× |
| Goccy | 4932037 | 393.44 MB/s | 4062883 | 13509 | 5.7× |
| SonicFastest | 5043498 | 384.75 MB/s | 4878987 | 1736 | 5.6× |
| Easyjson | 8535926 | 227.33 MB/s | 3871265 | 15043 | 3.3× |
| LightningDecodeAny | 10087979 | 192.35 MB/s | 7013524 | 219937 | 2.8× |
| JSONV2 | 12616146 | 153.81 MB/s | 3237180 | 13947 | 2.2× |
| Stdlib | 28124182 | 69.00 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1064190 | 3127.10 MB/s | 351704 | 1286 | 22.6× |
| Lightning | 1643074 | 2025.37 MB/s | 2488907 | 2995 | 14.6× |
| SonicFastest | 2006497 | 1658.53 MB/s | 5896315 | 4263 | 12.0× |
| Sonic | 2015437 | 1651.17 MB/s | 5896570 | 4263 | 11.9× |
| LightningDecodeAny | 3618472 | 849.47 MB/s | 4886622 | 56892 | 6.6× |
| Goccy | 6105968 | 545.01 MB/s | 3948915 | 3817 | 3.9× |
| JSONV2 | 7708272 | 431.72 MB/s | 5364499 | 13243 | 3.1× |
| Stdlib | 24040223 | 138.43 MB/s | 5565609 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 227230 | 969.71 MB/s | 136896 | 228 | 10.8× |
| LightningDestructive | 236145 | 933.10 MB/s | 136896 | 228 | 10.4× |
| Sonic | 490939 | 448.83 MB/s | 351385 | 262 | 5.0× |
| Goccy | 500879 | 439.92 MB/s | 363410 | 1066 | 4.9× |
| SonicFastest | 509645 | 432.35 MB/s | 352394 | 262 | 4.8× |
| Easyjson | 654631 | 336.60 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 715162 | 308.11 MB/s | 129746 | 470 | 3.4× |
| LightningDecodeAny | 1026287 | 105.54 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2456284 | 89.71 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14768706 | 548.46 MB/s | 15059834 | 51649 | 7.6× |
| Lightning | 15085758 | 536.93 MB/s | 15059836 | 51649 | 7.4× |
| Sonic | 17787512 | 455.38 MB/s | 19860870 | 41640 | 6.3× |
| SonicFastest | 17804680 | 454.94 MB/s | 19859504 | 41640 | 6.3× |
| Goccy | 25251561 | 320.77 MB/s | 19057286 | 107155 | 4.4× |
| Easyjson | 37287190 | 217.23 MB/s | 15059619 | 41643 | 3.0× |
| LightningDecodeAny | 44156124 | 117.83 MB/s | 34333346 | 912810 | 2.5× |
| JSONV2 | 48777838 | 166.06 MB/s | 15233770 | 78972 | 2.3× |
| Stdlib | 111917412 | 72.38 MB/s | 15665069 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6265325 | 476.19 MB/s | 3985336 | 30015 | 9.2× |
| Lightning | 6419929 | 464.72 MB/s | 3985340 | 30015 | 8.9× |
| Sonic | 9311747 | 320.40 MB/s | 9131433 | 57804 | 6.2× |
| SonicFastest | 9313062 | 320.35 MB/s | 9130868 | 57804 | 6.2× |
| Easyjson | 18620748 | 160.22 MB/s | 9479441 | 30115 | 3.1× |
| Goccy | 18628108 | 160.16 MB/s | 9889888 | 273621 | 3.1× |
| LightningDecodeAny | 20904018 | 87.74 MB/s | 20023837 | 408557 | 2.7× |
| JSONV2 | 26957406 | 110.67 MB/s | 9257059 | 86278 | 2.1× |
| Stdlib | 57364223 | 52.01 MB/s | 9258086 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1474615 | 490.70 MB/s | 907392 | 3618 | 9.5× |
| Lightning | 1552309 | 466.14 MB/s | 907387 | 3618 | 9.1× |
| Sonic | 2109783 | 342.97 MB/s | 2371791 | 3683 | 6.7× |
| SonicFastest | 2117814 | 341.67 MB/s | 2371986 | 3683 | 6.6× |
| Easyjson | 5403544 | 133.91 MB/s | 2847906 | 3698 | 2.6× |
| Goccy | 5523227 | 131.01 MB/s | 2702492 | 80267 | 2.5× |
| LightningDecodeAny | 5634785 | 115.46 MB/s | 5752202 | 80172 | 2.5× |
| JSONV2 | 6311093 | 114.65 MB/s | 2704701 | 7318 | 2.2× |
| Stdlib | 14056510 | 51.48 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2044853 | 771.38 MB/s | 907387 | 3618 | 9.8× |
| LightningDestructive | 2068012 | 762.74 MB/s | 907392 | 3618 | 9.7× |
| SonicFastest | 2427366 | 649.82 MB/s | 3224570 | 3683 | 8.2× |
| Sonic | 2431862 | 648.62 MB/s | 3224917 | 3683 | 8.2× |
| LightningDecodeAny | 4817626 | 156.38 MB/s | 5752199 | 80172 | 4.2× |
| Easyjson | 6360983 | 247.97 MB/s | 2847905 | 3698 | 3.1× |
| Goccy | 6650124 | 237.19 MB/s | 3477130 | 80261 | 3.0× |
| JSONV2 | 7139053 | 220.95 MB/s | 2704551 | 7318 | 2.8× |
| Stdlib | 20007799 | 78.84 MB/s | 2704548 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 252453 | 594.66 MB/s | 81920 | 1 | 8.6× |
| LightningDestructive | 263037 | 570.73 MB/s | 81920 | 1 | 8.2× |
| SonicFastest | 389408 | 385.52 MB/s | 407550 | 16 | 5.5× |
| Sonic | 389945 | 384.99 MB/s | 407655 | 16 | 5.5× |
| LightningDecodeAny | 600784 | 249.88 MB/s | 746007 | 10020 | 3.6× |
| Goccy | 1033465 | 145.26 MB/s | 326041 | 10005 | 2.1× |
| JSONV2 | 1102349 | 136.19 MB/s | 357723 | 20 | 2.0× |
| Stdlib | 2161094 | 69.47 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 34147 | 823.40 MB/s | 30272 | 105 | 10.1× |
| LightningDestructive | 35699 | 787.62 MB/s | 30144 | 103 | 9.7× |
| SonicFastest | 57659 | 487.64 MB/s | 59484 | 83 | 6.0× |
| Sonic | 58148 | 483.54 MB/s | 59468 | 83 | 5.9× |
| Goccy | 80328 | 350.03 MB/s | 59246 | 188 | 4.3× |
| Easyjson | 82185 | 342.12 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 132451 | 212.28 MB/s | 36897 | 242 | 2.6× |
| LightningDecodeAny | 169239 | 166.14 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 344744 | 81.56 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1954 | 1191.42 MB/s | 32 | 1 | 13.4× |
| LightningDestructive | 2099 | 1108.89 MB/s | 32 | 1 | 12.5× |
| Sonic | 4731 | 492.11 MB/s | 3713 | 4 | 5.5× |
| SonicFastest | 4733 | 491.87 MB/s | 3712 | 4 | 5.5× |
| Goccy | 4938 | 471.49 MB/s | 3649 | 4 | 5.3× |
| Easyjson | 5067 | 459.43 MB/s | 192 | 2 | 5.2× |
| JSONV2 | 8321 | 279.77 MB/s | 1000 | 6 | 3.2× |
| LightningDecodeAny | 10717 | 157.23 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 26234 | 88.74 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 215 | 879.08 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 226 | 837.97 MB/s | 0 | 0 | 12.4× |
| Goccy | 429 | 440.89 MB/s | 304 | 2 | 6.5× |
| Easyjson | 594 | 318.16 MB/s | 0 | 0 | 4.7× |
| SonicFastest | 665 | 284.15 MB/s | 341 | 3 | 4.2× |
| Sonic | 669 | 282.66 MB/s | 341 | 3 | 4.2× |
| JSONV2 | 1056 | 178.90 MB/s | 112 | 1 | 2.6× |
| LightningDecodeAny | 1345 | 99.62 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2786 | 67.84 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1491 | 1469.95 MB/s | 0 | 0 | 12.6× |
| LightningDestructive | 1559 | 1405.29 MB/s | 0 | 0 | 12.0× |
| Goccy | 3610 | 606.90 MB/s | 2864 | 4 | 5.2× |
| Easyjson | 3756 | 583.37 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6623 | 330.82 MB/s | 3601 | 38 | 2.8× |
| Sonic | 6846 | 320.05 MB/s | 3601 | 38 | 2.7× |
| JSONV2 | 8173 | 268.08 MB/s | 640 | 6 | 2.3× |
| LightningDecodeAny | 9427 | 192.12 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 18782 | 116.65 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 698858 | 730.44 MB/s | 703778 | 1012 | 10.1× |
| Lightning | 742055 | 687.92 MB/s | 703778 | 1012 | 9.5× |
| SonicFastest | 1246459 | 409.54 MB/s | 1310728 | 2014 | 5.7× |
| Sonic | 1247982 | 409.04 MB/s | 1311135 | 2014 | 5.7× |
| Goccy | 1338908 | 381.26 MB/s | 1138018 | 5006 | 5.3× |
| Easyjson | 1688789 | 302.27 MB/s | 863780 | 3012 | 4.2× |
| JSONV2 | 3488736 | 146.32 MB/s | 1075954 | 12645 | 2.0× |
| LightningDecodeAny | 3677343 | 125.49 MB/s | 2785927 | 66022 | 1.9× |
| Stdlib | 7052420 | 72.38 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 643 | 30772.03 MB/s | 0 | 0 | 249.5× |
| LightningDestructive | 975 | 20290.81 MB/s | 0 | 0 | 164.5× |
| SonicFastest | 6550 | 3021.09 MB/s | 21107 | 3 | 24.5× |
| Goccy | 29317 | 675.01 MB/s | 20492 | 2 | 5.5× |
| Sonic | 32330 | 612.09 MB/s | 20616 | 3 | 5.0× |
| JSONV2 | 33275 | 594.70 MB/s | 8 | 1 | 4.8× |
| Easyjson | 97355 | 203.27 MB/s | 0 | 0 | 1.6× |
| LightningDecodeAny | 103151 | 191.84 MB/s | 117104 | 2019 | 1.6× |
| Stdlib | 160463 | 123.32 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2225 | 8145.50 MB/s | 432 | 2 | 54.4× |
| LightningDestructive | 2570 | 7051.44 MB/s | 0 | 0 | 47.1× |
| Easyjson | 5038 | 3597.68 MB/s | 432 | 2 | 24.0× |
| SonicFastest | 8380 | 2162.73 MB/s | 20449 | 5 | 14.4× |
| Sonic | 8396 | 2158.72 MB/s | 20441 | 5 | 14.4× |
| LightningDecodeAny | 19308 | 926.16 MB/s | 29088 | 191 | 6.3× |
| Goccy | 19660 | 921.89 MB/s | 19460 | 2 | 6.2× |
| JSONV2 | 51353 | 352.93 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 121015 | 149.77 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2271846 | 884.08 MB/s | 2960389 | 7411 | 9.8× |
| Lightning | 2405257 | 835.04 MB/s | 2962102 | 7417 | 9.3× |
| SonicFastest | 4098143 | 490.10 MB/s | 5154554 | 7085 | 5.4× |
| Sonic | 4162812 | 482.48 MB/s | 5153828 | 7085 | 5.4× |
| Goccy | 4752462 | 422.62 MB/s | 5409447 | 15830 | 4.7× |
| Easyjson | 5724498 | 350.86 MB/s | 2981484 | 7439 | 3.9× |
| LightningDecodeAny | 7198347 | 158.69 MB/s | 7386072 | 134751 | 3.1× |
| JSONV2 | 7757761 | 258.90 MB/s | 3173679 | 14563 | 2.9× |
| Stdlib | 22328934 | 89.95 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 850 | 645.91 MB/s | 480 | 1 | 8.0× |
| Lightning | 853 | 643.94 MB/s | 480 | 1 | 8.0× |
| LightningDecodeAny | 2181 | 251.22 MB/s | 2261 | 50 | 3.1× |
| Easyjson | 2367 | 231.91 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2387 | 230.02 MB/s | 2260 | 8 | 2.9× |
| Sonic | 2401 | 228.66 MB/s | 2260 | 8 | 2.8× |
| Goccy | 3470 | 158.22 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3640 | 150.82 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6814 | 80.57 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 497896 | 1268.37 MB/s | 364472 | 566 | 12.7× |
| Lightning | 568557 | 1110.73 MB/s | 413001 | 878 | 11.2× |
| Sonic | 1003948 | 629.03 MB/s | 1066802 | 814 | 6.3× |
| SonicFastest | 1019782 | 619.26 MB/s | 1067941 | 814 | 6.2× |
| Easyjson | 1308260 | 482.71 MB/s | 422504 | 936 | 4.9× |
| Goccy | 1421391 | 444.29 MB/s | 990762 | 1201 | 4.5× |
| JSONV2 | 2293794 | 275.31 MB/s | 571592 | 3144 | 2.8× |
| LightningDecodeAny | 2541699 | 183.70 MB/s | 2010197 | 30295 | 2.5× |
| Stdlib | 6347271 | 99.49 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 815030 | 690.05 MB/s | 544249 | 448 | 7.8× |
| Lightning | 1023493 | 549.50 MB/s | 767621 | 1254 | 6.2× |
| SonicFastest | 1287244 | 436.91 MB/s | 1348872 | 1185 | 4.9× |
| Sonic | 1296319 | 433.85 MB/s | 1349941 | 1185 | 4.9× |
| Goccy | 1615598 | 348.11 MB/s | 1041560 | 1028 | 3.9× |
| Easyjson | 2163276 | 259.98 MB/s | 775154 | 1254 | 2.9× |
| LightningDecodeAny | 3096134 | 181.65 MB/s | 2114150 | 30295 | 2.1× |
| JSONV2 | 3209149 | 175.25 MB/s | 927405 | 3482 | 2.0× |
| Stdlib | 6371062 | 88.28 MB/s | 1011669 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 619272 | 860.98 MB/s | 333416 | 2084 | 10.1× |
| Lightning | 691600 | 770.93 MB/s | 368224 | 2293 | 9.1× |
| Sonic | 1110976 | 479.92 MB/s | 981917 | 3082 | 5.7× |
| SonicFastest | 1117041 | 477.31 MB/s | 981501 | 3082 | 5.6× |
| Easyjson | 1290999 | 413.00 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1466288 | 363.62 MB/s | 1167079 | 5408 | 4.3× |
| JSONV2 | 2826998 | 188.60 MB/s | 745421 | 13288 | 2.2× |
| LightningDecodeAny | 3548861 | 150.24 MB/s | 2674616 | 50596 | 1.8× |
| Stdlib | 6280426 | 84.90 MB/s | 798692 | 17133 | 1.0× |
