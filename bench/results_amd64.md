# JSON Deserialization Benchmarks

- generated 2026-07-01T22:11:41Z
- go version go1.26.4 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 107002 | 1189.46 MB/s | 49760 | 3 | 12.4× |
| LightningDestructive | 109159 | 1165.96 MB/s | 49280 | 2 | 12.2× |
| Sonic | 202752 | 627.74 MB/s | 214835 | 15 | 6.6× |
| SonicFastest | 203666 | 624.92 MB/s | 214707 | 15 | 6.5× |
| Goccy | 255702 | 497.75 MB/s | 225313 | 884 | 5.2× |
| Easyjson | 275149 | 462.57 MB/s | 122864 | 14 | 4.8× |
| JSONV2 | 473692 | 268.69 MB/s | 195129 | 1805 | 2.8× |
| LightningDecodeAny | 475720 | 198.97 MB/s | 465586 | 9714 | 2.8× |
| Stdlib | 1332082 | 95.55 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4310495 | 522.23 MB/s | 3122874 | 3081 | 8.0× |
| Lightning | 4398011 | 511.83 MB/s | 3122875 | 3081 | 7.8× |
| SonicFastest | 5042730 | 446.40 MB/s | 4873096 | 2584 | 6.8× |
| Sonic | 5177557 | 434.77 MB/s | 4872981 | 2584 | 6.6× |
| LightningDecodeAny | 12598712 | 178.67 MB/s | 7938298 | 281383 | 2.7× |
| Goccy | 13692983 | 164.39 MB/s | 4209300 | 56536 | 2.5× |
| Easyjson | 14032974 | 160.41 MB/s | 3099811 | 2120 | 2.4× |
| JSONV2 | 17044736 | 132.07 MB/s | 3123215 | 3083 | 2.0× |
| Stdlib | 34319343 | 65.59 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 589324 | 458.84 MB/s | 348025 | 1627 | 7.5× |
| LightningDestructive | 590504 | 457.92 MB/s | 348025 | 1627 | 7.5× |
| Sonic | 768668 | 351.78 MB/s | 641841 | 1147 | 5.7× |
| SonicFastest | 774970 | 348.92 MB/s | 641985 | 1147 | 5.7× |
| LightningDecodeAny | 1710791 | 158.06 MB/s | 1011488 | 37901 | 2.6× |
| Easyjson | 1761982 | 153.47 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1884760 | 143.47 MB/s | 541930 | 8122 | 2.3× |
| JSONV2 | 2301368 | 117.50 MB/s | 348162 | 1628 | 1.9× |
| Stdlib | 4407527 | 61.35 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1319992 | 1308.50 MB/s | 959848 | 5881 | 13.4× |
| Lightning | 1382947 | 1248.93 MB/s | 959890 | 5882 | 12.8× |
| Sonic | 2204927 | 783.34 MB/s | 2693939 | 5547 | 8.0× |
| SonicFastest | 2207008 | 782.60 MB/s | 2693869 | 5547 | 8.0× |
| Goccy | 2963140 | 582.90 MB/s | 2581977 | 14604 | 6.0× |
| Easyjson | 4315142 | 400.27 MB/s | 972032 | 5389 | 4.1× |
| LightningDecodeAny | 4478088 | 111.72 MB/s | 4976203 | 81466 | 4.0× |
| JSONV2 | 4853952 | 355.83 MB/s | 1011616 | 7594 | 3.7× |
| Stdlib | 17748948 | 97.31 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1019 | 1777.57 MB/s | 0 | 0 | 15.9× |
| LightningDestructive | 1050 | 1725.10 MB/s | 0 | 0 | 15.4× |
| Easyjson | 2995 | 604.93 MB/s | 24 | 1 | 5.4× |
| Goccy | 3305 | 548.29 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6595 | 274.76 MB/s | 3347 | 38 | 2.5× |
| Sonic | 6774 | 267.48 MB/s | 3345 | 38 | 2.4× |
| JSONV2 | 8186 | 221.36 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9555 | 189.54 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16178 | 112.01 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1098 | 1649.64 MB/s | 0 | 0 | 14.8× |
| LightningDestructive | 1156 | 1567.38 MB/s | 0 | 0 | 14.1× |
| Easyjson | 3010 | 602.00 MB/s | 24 | 1 | 5.4× |
| Goccy | 3219 | 562.98 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6192 | 292.64 MB/s | 3344 | 38 | 2.6× |
| Sonic | 6440 | 281.39 MB/s | 3346 | 38 | 2.5× |
| JSONV2 | 8286 | 218.69 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 10229 | 177.04 MB/s | 7536 | 158 | 1.6× |
| Stdlib | 16242 | 111.56 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1349 | 1343.16 MB/s | 144 | 10 | 12.1× |
| LightningDestructive | 1431 | 1266.06 MB/s | 144 | 10 | 11.4× |
| Easyjson | 3154 | 574.49 MB/s | 144 | 10 | 5.2× |
| Goccy | 3555 | 509.75 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 6346 | 285.54 MB/s | 3365 | 40 | 2.6× |
| Sonic | 6573 | 275.67 MB/s | 3365 | 40 | 2.5× |
| JSONV2 | 8618 | 210.26 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9608 | 188.48 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 16263 | 111.42 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 710 | 695.50 MB/s | 160 | 1 | 9.9× |
| LightningDestructive | 727 | 679.44 MB/s | 160 | 1 | 9.7× |
| Sonic | 1292 | 382.43 MB/s | 1077 | 8 | 5.4× |
| SonicFastest | 1303 | 379.24 MB/s | 1076 | 8 | 5.4× |
| LightningDecodeAny | 1976 | 249.44 MB/s | 1536 | 30 | 3.6× |
| Easyjson | 2550 | 193.71 MB/s | 448 | 3 | 2.8× |
| Goccy | 2736 | 180.58 MB/s | 856 | 23 | 2.6× |
| JSONV2 | 3298 | 149.77 MB/s | 528 | 7 | 2.1× |
| Stdlib | 7027 | 70.30 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 452 | 508.63 MB/s | 160 | 1 | 11.1× |
| LightningDestructive | 454 | 506.19 MB/s | 160 | 1 | 11.1× |
| Sonic | 928 | 247.89 MB/s | 801 | 8 | 5.4× |
| SonicFastest | 936 | 245.62 MB/s | 801 | 8 | 5.4× |
| LightningDecodeAny | 1666 | 137.43 MB/s | 1536 | 30 | 3.0× |
| Easyjson | 1747 | 131.63 MB/s | 448 | 3 | 2.9× |
| Goccy | 1843 | 124.83 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2641 | 87.08 MB/s | 528 | 7 | 1.9× |
| Stdlib | 5030 | 45.72 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 90919 | 716.38 MB/s | 172434 | 107 | 7.6× |
| LightningDestructive | 97146 | 670.45 MB/s | 166213 | 102 | 7.1× |
| Sonic | 146341 | 445.07 MB/s | 235783 | 65 | 4.7× |
| SonicFastest | 149129 | 436.75 MB/s | 235859 | 65 | 4.6× |
| Goccy | 183299 | 355.33 MB/s | 228078 | 134 | 3.8× |
| LightningDecodeAny | 209452 | 254.61 MB/s | 176961 | 3252 | 3.3× |
| JSONV2 | 266382 | 244.51 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 688390 | 94.61 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2620428 | 740.52 MB/s | 2846865 | 2238 | 10.9× |
| Lightning | 2713989 | 714.99 MB/s | 2846866 | 2238 | 10.5× |
| SonicFastest | 5087349 | 381.43 MB/s | 4882002 | 1736 | 5.6× |
| Sonic | 5118498 | 379.11 MB/s | 4881797 | 1736 | 5.6× |
| Goccy | 5377583 | 360.84 MB/s | 4061290 | 13509 | 5.3× |
| Easyjson | 8233828 | 235.67 MB/s | 3871264 | 15043 | 3.5× |
| LightningDecodeAny | 10248901 | 189.33 MB/s | 7013526 | 219937 | 2.8× |
| JSONV2 | 12984812 | 149.44 MB/s | 3237190 | 13947 | 2.2× |
| Stdlib | 28463634 | 68.17 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1075969 | 3092.87 MB/s | 351704 | 1286 | 22.5× |
| Lightning | 1705779 | 1950.92 MB/s | 2488907 | 2995 | 14.2× |
| Sonic | 2157517 | 1542.44 MB/s | 5895881 | 4263 | 11.2× |
| SonicFastest | 2157757 | 1542.26 MB/s | 5895904 | 4263 | 11.2× |
| LightningDecodeAny | 3861768 | 795.95 MB/s | 4886621 | 56892 | 6.3× |
| Goccy | 6204505 | 536.36 MB/s | 3948917 | 3817 | 3.9× |
| JSONV2 | 8310229 | 400.45 MB/s | 5364508 | 13243 | 2.9× |
| Stdlib | 24243534 | 137.27 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 227790 | 967.32 MB/s | 136896 | 228 | 10.8× |
| LightningDestructive | 238801 | 922.72 MB/s | 136896 | 228 | 10.3× |
| SonicFastest | 491868 | 447.98 MB/s | 350694 | 262 | 5.0× |
| Sonic | 493502 | 446.49 MB/s | 350537 | 262 | 5.0× |
| Goccy | 495321 | 444.85 MB/s | 363700 | 1066 | 5.0× |
| Easyjson | 673909 | 326.97 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 788229 | 279.55 MB/s | 129747 | 470 | 3.1× |
| LightningDecodeAny | 1048106 | 103.34 MB/s | 861394 | 11944 | 2.4× |
| Stdlib | 2465465 | 89.37 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 14720889 | 550.24 MB/s | 15059834 | 51649 | 7.6× |
| Lightning | 15320186 | 528.72 MB/s | 15059837 | 51649 | 7.3× |
| SonicFastest | 18208605 | 444.85 MB/s | 19859313 | 41640 | 6.2× |
| Sonic | 18246275 | 443.93 MB/s | 19861791 | 41640 | 6.2× |
| Goccy | 28058200 | 288.69 MB/s | 19221592 | 107156 | 4.0× |
| Easyjson | 34797863 | 232.77 MB/s | 15059617 | 41643 | 3.2× |
| LightningDecodeAny | 44224697 | 117.65 MB/s | 34333348 | 912810 | 2.5× |
| JSONV2 | 50951728 | 158.97 MB/s | 15233770 | 78972 | 2.2× |
| Stdlib | 112396937 | 72.07 MB/s | 15665067 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5832566 | 511.52 MB/s | 3985336 | 30015 | 9.9× |
| Lightning | 5975771 | 499.26 MB/s | 3985337 | 30015 | 9.7× |
| SonicFastest | 9540428 | 312.72 MB/s | 9130170 | 57804 | 6.1× |
| Sonic | 9571698 | 311.70 MB/s | 9130621 | 57804 | 6.1× |
| Goccy | 19384974 | 153.91 MB/s | 9868665 | 273619 | 3.0× |
| Easyjson | 19659673 | 151.76 MB/s | 9479441 | 30115 | 2.9× |
| LightningDecodeAny | 22019333 | 83.30 MB/s | 20023835 | 408557 | 2.6× |
| JSONV2 | 29113647 | 102.48 MB/s | 9257051 | 86278 | 2.0× |
| Stdlib | 57947781 | 51.49 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1342151 | 539.13 MB/s | 907392 | 3618 | 10.7× |
| Lightning | 1423342 | 508.38 MB/s | 907387 | 3618 | 10.1× |
| SonicFastest | 2229067 | 324.62 MB/s | 2372422 | 3683 | 6.4× |
| Sonic | 2240288 | 322.99 MB/s | 2371870 | 3683 | 6.4× |
| Easyjson | 5579744 | 129.68 MB/s | 2847907 | 3698 | 2.6× |
| LightningDecodeAny | 5699074 | 114.15 MB/s | 5752202 | 80172 | 2.5× |
| Goccy | 5893727 | 122.77 MB/s | 2751615 | 80270 | 2.4× |
| JSONV2 | 6621944 | 109.27 MB/s | 2704705 | 7318 | 2.2× |
| Stdlib | 14362623 | 50.38 MB/s | 2704550 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2056571 | 766.98 MB/s | 907388 | 3618 | 9.8× |
| LightningDestructive | 2068725 | 762.48 MB/s | 907392 | 3618 | 9.7× |
| SonicFastest | 2603447 | 605.87 MB/s | 3226473 | 3683 | 7.7× |
| Sonic | 2640584 | 597.35 MB/s | 3224093 | 3683 | 7.6× |
| LightningDecodeAny | 4903546 | 153.64 MB/s | 5752200 | 80172 | 4.1× |
| Easyjson | 6956454 | 226.75 MB/s | 2847907 | 3698 | 2.9× |
| Goccy | 7129802 | 221.23 MB/s | 3502274 | 80262 | 2.8× |
| JSONV2 | 7261113 | 217.23 MB/s | 2704558 | 7318 | 2.8× |
| Stdlib | 20101313 | 78.47 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 234248 | 640.88 MB/s | 81920 | 1 | 9.4× |
| LightningDestructive | 242625 | 618.75 MB/s | 81920 | 1 | 9.0× |
| Sonic | 385690 | 389.23 MB/s | 407718 | 16 | 5.7× |
| SonicFastest | 408421 | 367.57 MB/s | 407672 | 16 | 5.4× |
| LightningDecodeAny | 613903 | 244.54 MB/s | 746006 | 10020 | 3.6× |
| Goccy | 1005850 | 149.25 MB/s | 325597 | 10005 | 2.2× |
| JSONV2 | 1155025 | 129.97 MB/s | 357723 | 20 | 1.9× |
| Stdlib | 2191075 | 68.52 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 34486 | 815.31 MB/s | 30272 | 105 | 10.1× |
| LightningDestructive | 36005 | 780.91 MB/s | 30144 | 103 | 9.7× |
| SonicFastest | 60422 | 465.35 MB/s | 59503 | 83 | 5.8× |
| Sonic | 61402 | 457.92 MB/s | 59545 | 83 | 5.7× |
| Goccy | 82302 | 341.63 MB/s | 59279 | 188 | 4.2× |
| Easyjson | 84397 | 333.15 MB/s | 32304 | 138 | 4.1× |
| JSONV2 | 148524 | 189.31 MB/s | 36897 | 242 | 2.3× |
| LightningDecodeAny | 175161 | 160.52 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 349017 | 80.56 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1970 | 1181.96 MB/s | 32 | 1 | 13.4× |
| LightningDestructive | 2098 | 1109.79 MB/s | 32 | 1 | 12.6× |
| Sonic | 4852 | 479.80 MB/s | 3711 | 4 | 5.4× |
| SonicFastest | 4914 | 473.74 MB/s | 3715 | 4 | 5.4× |
| Goccy | 5047 | 461.30 MB/s | 3649 | 4 | 5.2× |
| Easyjson | 5077 | 458.53 MB/s | 192 | 2 | 5.2× |
| JSONV2 | 8590 | 271.01 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10930 | 154.17 MB/s | 9960 | 195 | 2.4× |
| Stdlib | 26419 | 88.12 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 216 | 874.63 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 223 | 849.02 MB/s | 0 | 0 | 12.6× |
| Goccy | 430 | 439.11 MB/s | 304 | 2 | 6.5× |
| Easyjson | 585 | 323.25 MB/s | 0 | 0 | 4.8× |
| Sonic | 631 | 299.66 MB/s | 341 | 3 | 4.5× |
| SonicFastest | 634 | 298.35 MB/s | 341 | 3 | 4.4× |
| JSONV2 | 1022 | 184.87 MB/s | 112 | 1 | 2.8× |
| LightningDecodeAny | 1364 | 98.23 MB/s | 1160 | 25 | 2.1× |
| Stdlib | 2814 | 67.16 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1469 | 1491.02 MB/s | 0 | 0 | 12.9× |
| LightningDestructive | 1542 | 1420.44 MB/s | 0 | 0 | 12.3× |
| Goccy | 3746 | 584.87 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3870 | 566.14 MB/s | 24 | 1 | 4.9× |
| SonicFastest | 6755 | 324.33 MB/s | 3602 | 38 | 2.8× |
| Sonic | 6997 | 313.15 MB/s | 3602 | 38 | 2.7× |
| JSONV2 | 8575 | 255.51 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9631 | 188.04 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 19006 | 115.28 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 713328 | 715.63 MB/s | 703778 | 1012 | 10.1× |
| Lightning | 761164 | 670.65 MB/s | 703778 | 1012 | 9.5× |
| Goccy | 1360106 | 375.32 MB/s | 1140434 | 5006 | 5.3× |
| Sonic | 1369569 | 372.73 MB/s | 1309162 | 2014 | 5.3× |
| SonicFastest | 1378189 | 370.40 MB/s | 1308436 | 2014 | 5.2× |
| Easyjson | 1747893 | 292.05 MB/s | 863778 | 3012 | 4.1× |
| JSONV2 | 3529064 | 144.65 MB/s | 1075967 | 12645 | 2.0× |
| LightningDecodeAny | 3760236 | 122.72 MB/s | 2785926 | 66022 | 1.9× |
| Stdlib | 7232378 | 70.58 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 546 | 36243.10 MB/s | 0 | 0 | 299.4× |
| LightningDestructive | 794 | 24932.18 MB/s | 0 | 0 | 205.9× |
| SonicFastest | 7775 | 2545.09 MB/s | 21069 | 3 | 21.0× |
| Goccy | 31350 | 631.22 MB/s | 20492 | 2 | 5.2× |
| Sonic | 33111 | 597.66 MB/s | 20642 | 3 | 4.9× |
| JSONV2 | 33349 | 593.38 MB/s | 8 | 1 | 4.9× |
| Easyjson | 100536 | 196.84 MB/s | 0 | 0 | 1.6× |
| LightningDecodeAny | 103078 | 191.97 MB/s | 117104 | 2019 | 1.6× |
| Stdlib | 163446 | 121.07 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2265 | 8002.64 MB/s | 432 | 2 | 54.6× |
| LightningDestructive | 2362 | 7674.75 MB/s | 0 | 0 | 52.3× |
| Easyjson | 5116 | 3542.52 MB/s | 432 | 2 | 24.2× |
| SonicFastest | 8659 | 2093.11 MB/s | 20459 | 5 | 14.3× |
| Sonic | 8688 | 2086.01 MB/s | 20429 | 5 | 14.2× |
| LightningDecodeAny | 19947 | 896.46 MB/s | 29088 | 191 | 6.2× |
| Goccy | 22342 | 811.20 MB/s | 19460 | 2 | 5.5× |
| JSONV2 | 51161 | 354.25 MB/s | 16501 | 50 | 2.4× |
| Stdlib | 123567 | 146.67 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2604202 | 771.25 MB/s | 2951716 | 7277 | 8.6× |
| Lightning | 2615989 | 767.78 MB/s | 2953430 | 7283 | 8.6× |
| Sonic | 4528429 | 443.53 MB/s | 5156481 | 7085 | 5.0× |
| SonicFastest | 4546342 | 441.78 MB/s | 5156988 | 7085 | 5.0× |
| Goccy | 5056969 | 397.17 MB/s | 5410912 | 15834 | 4.5× |
| Easyjson | 5584411 | 359.66 MB/s | 2981480 | 7439 | 4.0× |
| LightningDecodeAny | 7494780 | 152.41 MB/s | 7386074 | 134751 | 3.0× |
| JSONV2 | 8109529 | 247.67 MB/s | 3173676 | 14563 | 2.8× |
| Stdlib | 22525992 | 89.16 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 860 | 638.07 MB/s | 480 | 1 | 8.1× |
| LightningDestructive | 868 | 632.86 MB/s | 480 | 1 | 8.0× |
| LightningDecodeAny | 2188 | 250.41 MB/s | 2261 | 50 | 3.2× |
| Easyjson | 2252 | 243.80 MB/s | 1616 | 5 | 3.1× |
| SonicFastest | 2440 | 225.00 MB/s | 2260 | 8 | 2.9× |
| Sonic | 2476 | 221.68 MB/s | 2260 | 8 | 2.8× |
| Goccy | 3424 | 160.36 MB/s | 2129 | 43 | 2.0× |
| JSONV2 | 3576 | 153.52 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6971 | 78.76 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 522625 | 1208.35 MB/s | 364473 | 566 | 12.4× |
| Lightning | 594452 | 1062.35 MB/s | 413001 | 878 | 10.9× |
| SonicFastest | 1088074 | 580.40 MB/s | 1071944 | 814 | 6.0× |
| Sonic | 1101159 | 573.50 MB/s | 1072050 | 814 | 5.9× |
| Easyjson | 1373537 | 459.77 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1477362 | 427.46 MB/s | 991381 | 1201 | 4.4× |
| JSONV2 | 2453547 | 257.39 MB/s | 571593 | 3144 | 2.7× |
| LightningDecodeAny | 2630292 | 177.51 MB/s | 2010198 | 30295 | 2.5× |
| Stdlib | 6506676 | 97.06 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 822742 | 683.58 MB/s | 544249 | 448 | 7.6× |
| Lightning | 1067257 | 526.97 MB/s | 767622 | 1254 | 5.9× |
| Sonic | 1360480 | 413.39 MB/s | 1349315 | 1185 | 4.6× |
| SonicFastest | 1398591 | 402.12 MB/s | 1350461 | 1185 | 4.5× |
| Goccy | 1677911 | 335.18 MB/s | 1042736 | 1028 | 3.7× |
| Easyjson | 2213251 | 254.11 MB/s | 775154 | 1254 | 2.8× |
| LightningDecodeAny | 3220128 | 174.65 MB/s | 2114150 | 30295 | 1.9× |
| JSONV2 | 3323773 | 169.21 MB/s | 927405 | 3482 | 1.9× |
| Stdlib | 6270303 | 89.69 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 641601 | 831.01 MB/s | 333416 | 2084 | 10.1× |
| Lightning | 716843 | 743.79 MB/s | 368224 | 2293 | 9.0× |
| SonicFastest | 1144556 | 465.84 MB/s | 981650 | 3082 | 5.6× |
| Sonic | 1146112 | 465.21 MB/s | 982310 | 3082 | 5.6× |
| Easyjson | 1351296 | 394.57 MB/s | 428362 | 3273 | 4.8× |
| Goccy | 1458900 | 365.47 MB/s | 1167079 | 5408 | 4.4× |
| JSONV2 | 3013498 | 176.93 MB/s | 745420 | 13288 | 2.1× |
| LightningDecodeAny | 3729352 | 142.97 MB/s | 2674617 | 50596 | 1.7× |
| Stdlib | 6458875 | 82.55 MB/s | 798693 | 17133 | 1.0× |
