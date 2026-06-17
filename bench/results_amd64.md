# JSON Deserialization Benchmarks

- generated 2026-06-17T17:19:58Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 459311 | 588.71 MB/s | 151224 | 495 | 5.3× |
| Sonic | 552727 | 489.22 MB/s | 645440 | 1147 | 4.4× |
| Easyjson | 1240746 | 217.94 MB/s | 330320 | 753 | 2.0× |
| JSONV2 | 1518912 | 178.02 MB/s | 348528 | 1628 | 1.6× |
| Stdlib | 2430931 | 111.23 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1388931 | 1243.55 MB/s | 979492 | 6803 | 6.3× |
| Sonic | 1719762 | 1004.33 MB/s | 2730897 | 5548 | 5.1× |
| JSONV2 | 2972392 | 581.08 MB/s | 1012458 | 7594 | 3.0× |
| Easyjson | 3053891 | 565.57 MB/s | 986873 | 6015 | 2.9× |
| Stdlib | 8811174 | 196.02 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 889 | 2037.54 MB/s | 144 | 10 | 10.7× |
| Easyjson | 1856 | 976.39 MB/s | 144 | 10 | 5.1× |
| Sonic | 4706 | 385.07 MB/s | 3377 | 40 | 2.0× |
| JSONV2 | 5306 | 341.52 MB/s | 632 | 7 | 1.8× |
| Stdlib | 9487 | 191.00 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 619 | 2928.76 MB/s | 0 | 0 | 15.4× |
| Easyjson | 1676 | 1081.40 MB/s | 24 | 1 | 5.7× |
| Sonic | 4607 | 393.35 MB/s | 3359 | 38 | 2.1× |
| JSONV2 | 5271 | 343.77 MB/s | 640 | 6 | 1.8× |
| Stdlib | 9517 | 190.40 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 692 | 2616.86 MB/s | 0 | 0 | 14.2× |
| Easyjson | 1675 | 1081.83 MB/s | 24 | 1 | 5.9× |
| Sonic | 4498 | 402.88 MB/s | 3362 | 38 | 2.2× |
| JSONV2 | 5181 | 349.75 MB/s | 640 | 6 | 1.9× |
| Stdlib | 9859 | 183.79 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 317 | 724.59 MB/s | 160 | 1 | 9.2× |
| Sonic | 675 | 340.57 MB/s | 812 | 8 | 4.3× |
| Easyjson | 1069 | 215.09 MB/s | 448 | 3 | 2.7× |
| JSONV2 | 1698 | 135.43 MB/s | 528 | 7 | 1.7× |
| Stdlib | 2923 | 78.69 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 537 | 920.09 MB/s | 160 | 1 | 7.4× |
| Sonic | 899 | 549.59 MB/s | 1090 | 8 | 4.4× |
| Easyjson | 1635 | 302.13 MB/s | 448 | 3 | 2.4× |
| JSONV2 | 2254 | 219.17 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3958 | 124.80 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 3414159 | 568.36 MB/s | 4888328 | 1736 | 4.8× |
| Lightning | 3579464 | 542.11 MB/s | 4282959 | 27849 | 4.5× |
| Easyjson | 6296029 | 308.21 MB/s | 4282951 | 27849 | 2.6× |
| JSONV2 | 7817433 | 248.22 MB/s | 3239264 | 13947 | 2.1× |
| Stdlib | 16276473 | 119.22 MB/s | 3551319 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 10226159 | 792.09 MB/s | 19858332 | 41641 | 6.1× |
| Lightning | 11843471 | 683.92 MB/s | 12528031 | 40027 | 5.3× |
| Easyjson | 24279453 | 333.62 MB/s | 15219650 | 61644 | 2.6× |
| JSONV2 | 31489488 | 257.23 MB/s | 15236434 | 78973 | 2.0× |
| Stdlib | 62344019 | 129.92 MB/s | 15665068 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 844 | 2596.66 MB/s | 0 | 0 | 12.9× |
| Easyjson | 2093 | 1046.87 MB/s | 24 | 1 | 5.2× |
| Sonic | 4977 | 440.24 MB/s | 3618 | 38 | 2.2× |
| JSONV2 | 5472 | 400.42 MB/s | 640 | 6 | 2.0× |
| Stdlib | 10843 | 202.07 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 502 | 39431.62 MB/s | 0 | 0 | 146.5× |
| JSONV2 | 21819 | 906.97 MB/s | 8 | 1 | 3.4× |
| Sonic | 25873 | 764.84 MB/s | 20693 | 3 | 2.8× |
| Easyjson | 45818 | 431.90 MB/s | 0 | 0 | 1.6× |
| Stdlib | 73509 | 269.20 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 5169 | 3506.33 MB/s | 20751 | 5 | 14.3× |
| Lightning | 5290 | 3425.98 MB/s | 18080 | 62 | 14.0× |
| Easyjson | 6865 | 2640.05 MB/s | 17648 | 60 | 10.8× |
| JSONV2 | 32858 | 551.58 MB/s | 16516 | 50 | 2.2× |
| Stdlib | 73808 | 245.56 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2627575 | 764.39 MB/s | 3587264 | 29254 | 5.1× |
| Sonic | 3160655 | 635.47 MB/s | 5156138 | 7085 | 4.3× |
| Easyjson | 4728130 | 424.80 MB/s | 3604107 | 29229 | 2.9× |
| JSONV2 | 5358319 | 374.84 MB/s | 3175426 | 14563 | 2.5× |
| Stdlib | 13506767 | 148.70 MB/s | 3589326 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 775 | 708.51 MB/s | 480 | 1 | 5.8× |
| Easyjson | 1711 | 320.82 MB/s | 1616 | 5 | 2.6× |
| Sonic | 1828 | 300.25 MB/s | 2281 | 8 | 2.5× |
| JSONV2 | 2491 | 220.38 MB/s | 1665 | 7 | 1.8× |
| Stdlib | 4486 | 122.39 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 613117 | 1030.01 MB/s | 634494 | 5544 | 6.7× |
| Sonic | 797957 | 791.41 MB/s | 1087592 | 814 | 5.2× |
| Easyjson | 1067396 | 591.64 MB/s | 591598 | 5205 | 3.9× |
| JSONV2 | 1681767 | 375.51 MB/s | 572266 | 3144 | 2.5× |
| Stdlib | 4135788 | 152.69 MB/s | 654667 | 6472 | 1.0× |
