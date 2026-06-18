# JSON Deserialization Benchmarks

- generated 2026-06-18T20:48:59Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 423554 | 638.41 MB/s | 348024 | 1627 | 7.3× |
| Sonic | 566050 | 477.70 MB/s | 642398 | 1147 | 5.5× |
| SonicFastest | 568164 | 475.92 MB/s | 642668 | 1147 | 5.5× |
| Goccy | 1355656 | 199.46 MB/s | 542177 | 8122 | 2.3× |
| Easyjson | 1356197 | 199.38 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 1718649 | 157.33 MB/s | 348158 | 1628 | 1.8× |
| Stdlib | 3107845 | 87.01 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1177311 | 1467.07 MB/s | 964643 | 6177 | 11.2× |
| Sonic | 1631603 | 1058.59 MB/s | 2705400 | 5548 | 8.1× |
| SonicFastest | 1634539 | 1056.69 MB/s | 2704898 | 5548 | 8.0× |
| Goccy | 1878405 | 919.51 MB/s | 2592933 | 14607 | 7.0× |
| Easyjson | 2956025 | 584.30 MB/s | 972032 | 5389 | 4.4× |
| JSONV2 | 3361081 | 513.88 MB/s | 1011674 | 7594 | 3.9× |
| Stdlib | 13139863 | 131.45 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 803 | 2256.91 MB/s | 0 | 0 | 14.3× |
| Easyjson | 2193 | 826.11 MB/s | 24 | 1 | 5.2× |
| Goccy | 2408 | 752.60 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 4780 | 379.05 MB/s | 3348 | 38 | 2.4× |
| Sonic | 4918 | 368.43 MB/s | 3346 | 38 | 2.3× |
| JSONV2 | 5821 | 311.27 MB/s | 640 | 6 | 2.0× |
| Stdlib | 11464 | 158.06 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 854 | 2121.99 MB/s | 0 | 0 | 13.3× |
| Easyjson | 2163 | 837.61 MB/s | 24 | 1 | 5.3× |
| Goccy | 2417 | 749.57 MB/s | 2608 | 4 | 4.7× |
| SonicFastest | 4758 | 380.82 MB/s | 3349 | 38 | 2.4× |
| Sonic | 4861 | 372.79 MB/s | 3347 | 38 | 2.3× |
| JSONV2 | 5844 | 310.07 MB/s | 640 | 6 | 1.9× |
| Stdlib | 11381 | 159.22 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1090 | 1661.95 MB/s | 144 | 10 | 10.3× |
| Easyjson | 2371 | 764.23 MB/s | 144 | 10 | 4.7× |
| Goccy | 2588 | 700.12 MB/s | 2600 | 5 | 4.4× |
| SonicFastest | 4825 | 375.54 MB/s | 3367 | 40 | 2.3× |
| Sonic | 5021 | 360.90 MB/s | 3367 | 40 | 2.2× |
| JSONV2 | 6244 | 290.18 MB/s | 632 | 7 | 1.8× |
| Stdlib | 11262 | 160.90 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 555 | 890.31 MB/s | 160 | 1 | 8.4× |
| SonicFastest | 946 | 522.23 MB/s | 1075 | 8 | 4.9× |
| Sonic | 950 | 519.84 MB/s | 1075 | 8 | 4.9× |
| Easyjson | 1742 | 283.54 MB/s | 448 | 3 | 2.7× |
| Goccy | 1908 | 258.94 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 2399 | 205.88 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4682 | 105.52 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 358 | 642.46 MB/s | 160 | 1 | 9.1× |
| SonicFastest | 673 | 341.91 MB/s | 801 | 8 | 4.9× |
| Sonic | 691 | 332.81 MB/s | 800 | 8 | 4.7× |
| Easyjson | 1182 | 194.60 MB/s | 448 | 3 | 2.8× |
| Goccy | 1285 | 178.94 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 1808 | 127.19 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3272 | 70.30 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2686575 | 722.28 MB/s | 2846868 | 2238 | 7.4× |
| Goccy | 3949885 | 491.27 MB/s | 4064404 | 13509 | 5.1× |
| Sonic | 3984145 | 487.05 MB/s | 4890185 | 1736 | 5.0× |
| SonicFastest | 4073697 | 476.34 MB/s | 4890121 | 1736 | 4.9× |
| Easyjson | 6176276 | 314.18 MB/s | 3871265 | 15043 | 3.2× |
| JSONV2 | 9087818 | 213.52 MB/s | 3237339 | 13947 | 2.2× |
| Stdlib | 19969562 | 97.17 MB/s | 3551321 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 11547254 | 701.47 MB/s | 15059837 | 51649 | 6.8× |
| SonicFastest | 15406573 | 525.75 MB/s | 19919991 | 41642 | 5.1× |
| Sonic | 15456819 | 524.04 MB/s | 19922937 | 41642 | 5.1× |
| Goccy | 21525003 | 376.31 MB/s | 17776770 | 107152 | 3.6× |
| Easyjson | 26363207 | 307.25 MB/s | 15059618 | 41643 | 3.0× |
| JSONV2 | 35228468 | 229.93 MB/s | 15233896 | 78973 | 2.2× |
| Stdlib | 78377123 | 103.35 MB/s | 15665072 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1122 | 1952.70 MB/s | 0 | 0 | 12.0× |
| Easyjson | 2722 | 804.91 MB/s | 24 | 1 | 5.0× |
| Goccy | 2759 | 794.00 MB/s | 2864 | 4 | 4.9× |
| SonicFastest | 5203 | 421.13 MB/s | 3607 | 38 | 2.6× |
| Sonic | 5330 | 411.07 MB/s | 3603 | 38 | 2.5× |
| JSONV2 | 5962 | 367.47 MB/s | 640 | 6 | 2.3× |
| Stdlib | 13519 | 162.07 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 502 | 39456.86 MB/s | 0 | 0 | 219.7× |
| SonicFastest | 5613 | 3525.39 MB/s | 21059 | 3 | 19.6× |
| Goccy | 20369 | 971.52 MB/s | 20492 | 2 | 5.4× |
| Sonic | 22694 | 871.98 MB/s | 20644 | 3 | 4.9× |
| JSONV2 | 28225 | 701.12 MB/s | 8 | 1 | 3.9× |
| Easyjson | 92720 | 213.43 MB/s | 0 | 0 | 1.2× |
| Stdlib | 110189 | 179.59 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2036 | 8903.27 MB/s | 864 | 4 | 46.2× |
| Easyjson | 3622 | 5004.22 MB/s | 432 | 2 | 25.9× |
| Sonic | 7317 | 2476.86 MB/s | 20421 | 5 | 12.8× |
| SonicFastest | 7415 | 2444.18 MB/s | 20424 | 5 | 12.7× |
| Goccy | 17498 | 1035.79 MB/s | 19460 | 2 | 5.4× |
| JSONV2 | 37172 | 487.58 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 93972 | 192.87 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2088734 | 961.58 MB/s | 2963985 | 7453 | 7.8× |
| Goccy | 3593761 | 558.88 MB/s | 5415120 | 15844 | 4.5× |
| Sonic | 3633278 | 552.80 MB/s | 5180269 | 7086 | 4.5× |
| SonicFastest | 4059433 | 494.77 MB/s | 5175004 | 7086 | 4.0× |
| Easyjson | 4158355 | 483.00 MB/s | 2981689 | 7441 | 3.9× |
| JSONV2 | 5710711 | 351.71 MB/s | 3173777 | 14563 | 2.9× |
| Stdlib | 16320134 | 123.07 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 803 | 683.42 MB/s | 480 | 1 | 5.7× |
| Easyjson | 1567 | 350.29 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 1678 | 327.18 MB/s | 2262 | 8 | 2.7× |
| Sonic | 1714 | 320.39 MB/s | 2263 | 8 | 2.7× |
| JSONV2 | 2329 | 235.76 MB/s | 1664 | 7 | 2.0× |
| Goccy | 2352 | 233.46 MB/s | 2129 | 43 | 1.9× |
| Stdlib | 4565 | 120.25 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 464543 | 1359.43 MB/s | 461114 | 1230 | 10.1× |
| Sonic | 967823 | 652.51 MB/s | 1069219 | 814 | 4.8× |
| SonicFastest | 977824 | 645.84 MB/s | 1069250 | 814 | 4.8× |
| Easyjson | 1008111 | 626.43 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1037888 | 608.46 MB/s | 993092 | 1202 | 4.5× |
| JSONV2 | 1787271 | 353.34 MB/s | 571617 | 3144 | 2.6× |
| Stdlib | 4687996 | 134.71 MB/s | 654665 | 6472 | 1.0× |
