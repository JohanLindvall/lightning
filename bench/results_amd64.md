# JSON Deserialization Benchmarks

- generated 2026-06-18T14:23:50Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 485841 | 556.57 MB/s | 348024 | 1627 | 5.3× |
| SonicFastest | 557498 | 485.03 MB/s | 644313 | 1147 | 4.6× |
| Sonic | 563838 | 479.58 MB/s | 644232 | 1147 | 4.5× |
| Easyjson | 1241042 | 217.88 MB/s | 330272 | 749 | 2.1× |
| Goccy | 1254610 | 215.53 MB/s | 544464 | 8122 | 2.0× |
| JSONV2 | 1603764 | 168.61 MB/s | 348580 | 1628 | 1.6× |
| Stdlib | 2557310 | 105.74 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1164596 | 1483.09 MB/s | 964644 | 6177 | 7.8× |
| Sonic | 1624714 | 1063.08 MB/s | 2740075 | 5548 | 5.6× |
| SonicFastest | 1661927 | 1039.28 MB/s | 2737831 | 5548 | 5.5× |
| Goccy | 1898899 | 909.58 MB/s | 2603913 | 14610 | 4.8× |
| JSONV2 | 3045731 | 567.09 MB/s | 1012361 | 7594 | 3.0× |
| Easyjson | 3125919 | 552.54 MB/s | 972033 | 5389 | 2.9× |
| Stdlib | 9064590 | 190.54 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 889 | 2038.32 MB/s | 144 | 10 | 10.6× |
| Easyjson | 1838 | 985.62 MB/s | 144 | 10 | 5.1× |
| Goccy | 2613 | 693.46 MB/s | 2603 | 5 | 3.6× |
| SonicFastest | 4548 | 398.42 MB/s | 3387 | 40 | 2.1× |
| Sonic | 4691 | 386.25 MB/s | 3386 | 40 | 2.0× |
| JSONV2 | 4949 | 366.11 MB/s | 632 | 7 | 1.9× |
| Stdlib | 9438 | 191.99 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 622 | 2911.54 MB/s | 0 | 0 | 15.4× |
| Easyjson | 1665 | 1088.40 MB/s | 24 | 1 | 5.8× |
| Goccy | 2610 | 694.13 MB/s | 2611 | 4 | 3.7× |
| SonicFastest | 4355 | 416.05 MB/s | 3372 | 38 | 2.2× |
| Sonic | 4489 | 403.65 MB/s | 3368 | 38 | 2.1× |
| JSONV2 | 4863 | 372.59 MB/s | 640 | 6 | 2.0× |
| Stdlib | 9598 | 188.79 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 678 | 2673.85 MB/s | 0 | 0 | 13.8× |
| Easyjson | 1663 | 1089.49 MB/s | 24 | 1 | 5.6× |
| Goccy | 2536 | 714.39 MB/s | 2611 | 4 | 3.7× |
| SonicFastest | 4405 | 411.35 MB/s | 3372 | 38 | 2.1× |
| Sonic | 4643 | 390.31 MB/s | 3371 | 38 | 2.0× |
| JSONV2 | 4874 | 371.76 MB/s | 641 | 6 | 1.9× |
| Stdlib | 9361 | 193.57 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 326 | 706.50 MB/s | 160 | 1 | 9.3× |
| SonicFastest | 671 | 342.89 MB/s | 813 | 8 | 4.5× |
| Sonic | 684 | 336.03 MB/s | 812 | 8 | 4.4× |
| Easyjson | 1122 | 204.91 MB/s | 448 | 3 | 2.7× |
| Goccy | 1287 | 178.73 MB/s | 585 | 23 | 2.3× |
| JSONV2 | 1667 | 137.94 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3013 | 76.33 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 519 | 951.34 MB/s | 160 | 1 | 8.1× |
| Sonic | 912 | 541.40 MB/s | 1091 | 8 | 4.6× |
| SonicFastest | 953 | 518.50 MB/s | 1091 | 8 | 4.4× |
| Easyjson | 1710 | 288.97 MB/s | 448 | 3 | 2.5× |
| Goccy | 1958 | 252.35 MB/s | 858 | 23 | 2.1× |
| JSONV2 | 2170 | 227.64 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4201 | 117.59 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2766040 | 701.53 MB/s | 2846871 | 2238 | 6.0× |
| Goccy | 3234997 | 599.84 MB/s | 4067114 | 13509 | 5.1× |
| SonicFastest | 3466762 | 559.74 MB/s | 4885648 | 1736 | 4.7× |
| Sonic | 3467885 | 559.55 MB/s | 4886365 | 1736 | 4.7× |
| Easyjson | 6128436 | 316.63 MB/s | 3871268 | 15043 | 2.7× |
| JSONV2 | 8000073 | 242.56 MB/s | 3239106 | 13947 | 2.1× |
| Stdlib | 16465703 | 117.85 MB/s | 3551319 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 10542352 | 768.33 MB/s | 15059843 | 51649 | 5.9× |
| Sonic | 10715845 | 755.89 MB/s | 19864953 | 41641 | 5.8× |
| SonicFastest | 10937988 | 740.54 MB/s | 19861757 | 41641 | 5.6× |
| Goccy | 18929346 | 427.91 MB/s | 17509980 | 107150 | 3.3× |
| Easyjson | 23868756 | 339.36 MB/s | 15059619 | 41643 | 2.6× |
| JSONV2 | 31599050 | 256.34 MB/s | 15236450 | 78973 | 2.0× |
| Stdlib | 61756256 | 131.16 MB/s | 15665063 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 917 | 2390.33 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2062 | 1062.74 MB/s | 24 | 1 | 5.2× |
| Goccy | 3039 | 720.99 MB/s | 2867 | 4 | 3.5× |
| JSONV2 | 4993 | 438.84 MB/s | 641 | 6 | 2.1× |
| SonicFastest | 5011 | 437.20 MB/s | 3628 | 38 | 2.1× |
| Sonic | 5160 | 424.58 MB/s | 3626 | 38 | 2.1× |
| Stdlib | 10719 | 204.40 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 496 | 39886.88 MB/s | 0 | 0 | 152.0× |
| SonicFastest | 3667 | 5396.80 MB/s | 21447 | 3 | 20.6× |
| Goccy | 19233 | 1028.93 MB/s | 20515 | 2 | 3.9× |
| JSONV2 | 21584 | 916.83 MB/s | 8 | 1 | 3.5× |
| Sonic | 26871 | 736.44 MB/s | 20715 | 3 | 2.8× |
| Easyjson | 46735 | 423.43 MB/s | 0 | 0 | 1.6× |
| Stdlib | 75383 | 262.51 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1976 | 9169.84 MB/s | 864 | 4 | 37.8× |
| Easyjson | 2978 | 6085.75 MB/s | 432 | 2 | 25.1× |
| SonicFastest | 5698 | 3180.61 MB/s | 20707 | 5 | 13.1× |
| Sonic | 5843 | 3101.96 MB/s | 20719 | 5 | 12.8× |
| Goccy | 17233 | 1051.70 MB/s | 19483 | 2 | 4.3× |
| JSONV2 | 35032 | 517.35 MB/s | 16520 | 50 | 2.1× |
| Stdlib | 74697 | 242.63 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1909009 | 1052.11 MB/s | 2963988 | 7453 | 7.1× |
| Sonic | 3063523 | 655.62 MB/s | 5158850 | 7085 | 4.4× |
| SonicFastest | 3097181 | 648.49 MB/s | 5158657 | 7085 | 4.3× |
| Goccy | 3494652 | 574.73 MB/s | 5485051 | 15841 | 3.9× |
| Easyjson | 4172782 | 481.33 MB/s | 2983386 | 7442 | 3.2× |
| JSONV2 | 5272481 | 380.94 MB/s | 3175176 | 14563 | 2.6× |
| Stdlib | 13465269 | 149.16 MB/s | 3589322 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 825 | 665.67 MB/s | 480 | 1 | 5.4× |
| SonicFastest | 1706 | 321.74 MB/s | 2284 | 8 | 2.6× |
| Easyjson | 1722 | 318.86 MB/s | 1616 | 5 | 2.6× |
| Sonic | 1739 | 315.78 MB/s | 2282 | 8 | 2.6× |
| Goccy | 2445 | 224.53 MB/s | 2133 | 43 | 1.8× |
| JSONV2 | 2503 | 219.31 MB/s | 1666 | 7 | 1.8× |
| Stdlib | 4439 | 123.68 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 476693 | 1324.78 MB/s | 461113 | 1230 | 8.4× |
| Sonic | 799221 | 790.16 MB/s | 1090146 | 814 | 5.0× |
| SonicFastest | 819066 | 771.02 MB/s | 1089193 | 814 | 4.9× |
| Easyjson | 942150 | 670.29 MB/s | 422505 | 936 | 4.2× |
| Goccy | 1070343 | 590.01 MB/s | 1006401 | 1204 | 3.7× |
| JSONV2 | 1671584 | 377.79 MB/s | 572140 | 3144 | 2.4× |
| Stdlib | 3997363 | 157.98 MB/s | 654666 | 6472 | 1.0× |
