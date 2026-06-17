# JSON Deserialization Benchmarks

- generated 2026-06-17T06:55:19Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 708953 | 381.41 MB/s | 333834 | 7649 | 4.4× |
| Sonic | 745957 | 362.49 MB/s | 917555 | 8294 | 4.2× |
| Easyjson | 1729550 | 156.34 MB/s | 962165 | 8297 | 1.8× |
| JSONV2 | 2135365 | 126.63 MB/s | 688474 | 15929 | 1.5× |
| Stdlib | 3098857 | 87.26 MB/s | 688019 | 15942 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 1845398 | 935.95 MB/s | 2736963 | 5548 | 4.9× |
| Lightning | 2890066 | 597.63 MB/s | 459888 | 3126 | 3.1× |
| JSONV2 | 3002634 | 575.23 MB/s | 1012459 | 7594 | 3.0× |
| Easyjson | 3030931 | 569.86 MB/s | 986872 | 6015 | 3.0× |
| Stdlib | 9078775 | 190.25 MB/s | 1234489 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 901 | 2011.14 MB/s | 144 | 10 | 10.6× |
| Easyjson | 1803 | 1005.07 MB/s | 144 | 10 | 5.3× |
| Sonic | 4653 | 389.45 MB/s | 3384 | 40 | 2.1× |
| JSONV2 | 5254 | 344.91 MB/s | 632 | 7 | 1.8× |
| Stdlib | 9570 | 189.34 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 621 | 2916.50 MB/s | 0 | 0 | 15.5× |
| Easyjson | 1653 | 1096.26 MB/s | 24 | 1 | 5.8× |
| Sonic | 4657 | 389.09 MB/s | 3367 | 38 | 2.1× |
| JSONV2 | 5181 | 349.73 MB/s | 640 | 6 | 1.9× |
| Stdlib | 9606 | 188.63 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 694 | 2612.72 MB/s | 0 | 0 | 13.7× |
| Easyjson | 1662 | 1090.45 MB/s | 24 | 1 | 5.7× |
| Sonic | 4620 | 392.23 MB/s | 3359 | 38 | 2.1× |
| JSONV2 | 5158 | 351.32 MB/s | 640 | 6 | 1.8× |
| Stdlib | 9493 | 190.88 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 343 | 670.66 MB/s | 160 | 1 | 8.5× |
| Sonic | 669 | 343.58 MB/s | 812 | 8 | 4.3× |
| Easyjson | 1075 | 213.89 MB/s | 448 | 3 | 2.7× |
| JSONV2 | 1585 | 145.12 MB/s | 528 | 7 | 1.8× |
| Stdlib | 2901 | 79.29 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 514 | 961.98 MB/s | 160 | 1 | 7.9× |
| Sonic | 934 | 529.16 MB/s | 1092 | 8 | 4.3× |
| Easyjson | 1602 | 308.45 MB/s | 448 | 3 | 2.5× |
| JSONV2 | 2106 | 234.55 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4054 | 121.85 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3513531 | 552.29 MB/s | 4282973 | 27849 | 4.7× |
| Sonic | 3918996 | 495.15 MB/s | 4886951 | 1736 | 4.2× |
| Easyjson | 6225806 | 311.68 MB/s | 4282948 | 27849 | 2.7× |
| JSONV2 | 8121224 | 238.94 MB/s | 3239325 | 13947 | 2.0× |
| Stdlib | 16628661 | 116.69 MB/s | 3551322 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 13589175 | 596.07 MB/s | 23105689 | 175893 | 5.0× |
| Lightning | 13636239 | 594.01 MB/s | 14139125 | 107156 | 4.9× |
| Easyjson | 26920918 | 300.88 MB/s | 19531273 | 128767 | 2.5× |
| JSONV2 | 38003387 | 213.14 MB/s | 19028356 | 280355 | 1.8× |
| Stdlib | 67393938 | 120.19 MB/s | 19438240 | 352028 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 857 | 2557.86 MB/s | 0 | 0 | 12.5× |
| Easyjson | 2102 | 1042.23 MB/s | 24 | 1 | 5.1× |
| Sonic | 4740 | 462.23 MB/s | 3616 | 38 | 2.3× |
| JSONV2 | 5559 | 394.13 MB/s | 641 | 6 | 1.9× |
| Stdlib | 10736 | 204.07 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 500 | 39544.27 MB/s | 0 | 0 | 143.3× |
| JSONV2 | 21749 | 909.87 MB/s | 8 | 1 | 3.3× |
| Sonic | 25302 | 782.10 MB/s | 20692 | 3 | 2.8× |
| Easyjson | 45724 | 432.79 MB/s | 0 | 0 | 1.6× |
| Stdlib | 71730 | 275.88 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 5373 | 3373.21 MB/s | 20721 | 5 | 13.6× |
| Lightning | 5508 | 3290.57 MB/s | 18080 | 62 | 13.3× |
| Easyjson | 6904 | 2625.28 MB/s | 17648 | 60 | 10.6× |
| JSONV2 | 32987 | 549.44 MB/s | 16517 | 50 | 2.2× |
| Stdlib | 73245 | 247.44 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 3372370 | 595.57 MB/s | 5157660 | 7085 | 3.8× |
| Lightning | 4198403 | 478.39 MB/s | 3378232 | 28699 | 3.0× |
| Easyjson | 4675556 | 429.57 MB/s | 3604034 | 29228 | 2.7× |
| JSONV2 | 5246113 | 382.85 MB/s | 3175421 | 14563 | 2.4× |
| Stdlib | 12758803 | 157.42 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 779 | 704.65 MB/s | 480 | 1 | 5.8× |
| Easyjson | 1728 | 317.76 MB/s | 1616 | 5 | 2.6× |
| Sonic | 1842 | 298.02 MB/s | 2282 | 8 | 2.4× |
| JSONV2 | 2484 | 221.04 MB/s | 1665 | 7 | 1.8× |
| Stdlib | 4489 | 122.31 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 620635 | 1017.53 MB/s | 633934 | 5539 | 6.3× |
| Sonic | 852025 | 741.19 MB/s | 1086133 | 814 | 4.6× |
| Easyjson | 1108389 | 569.76 MB/s | 591598 | 5205 | 3.5× |
| JSONV2 | 1669081 | 378.36 MB/s | 572274 | 3144 | 2.3× |
| Stdlib | 3887856 | 162.43 MB/s | 654666 | 6472 | 1.0× |
