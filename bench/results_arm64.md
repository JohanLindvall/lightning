# JSON Deserialization Benchmarks

- generated 2026-06-18T21:37:22Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 485598 | 556.84 MB/s | 348024 | 1627 | 6.9× |
| SonicFastest | 641113 | 421.77 MB/s | 503251 | 968 | 5.2× |
| Sonic | 642912 | 420.59 MB/s | 508219 | 968 | 5.2× |
| Goccy | 1382153 | 195.64 MB/s | 541164 | 8122 | 2.4× |
| Easyjson | 1412135 | 191.49 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2081315 | 129.92 MB/s | 348150 | 1628 | 1.6× |
| Stdlib | 3353578 | 80.63 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1576188 | 1095.81 MB/s | 964643 | 6177 | 8.5× |
| Sonic | 2082014 | 829.58 MB/s | 2952870 | 4020 | 6.4× |
| SonicFastest | 2102024 | 821.69 MB/s | 2927447 | 4020 | 6.4× |
| Goccy | 2468441 | 699.71 MB/s | 2587872 | 14607 | 5.4× |
| Easyjson | 4293356 | 402.30 MB/s | 972033 | 5389 | 3.1× |
| JSONV2 | 4329059 | 398.98 MB/s | 1011702 | 7594 | 3.1× |
| Stdlib | 13377432 | 129.11 MB/s | 1234451 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1239 | 1462.03 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2549 | 710.91 MB/s | 24 | 1 | 5.5× |
| Goccy | 2770 | 654.07 MB/s | 2608 | 4 | 5.1× |
| Sonic | 5897 | 307.28 MB/s | 3743 | 40 | 2.4× |
| SonicFastest | 5905 | 306.84 MB/s | 3759 | 40 | 2.4× |
| JSONV2 | 7704 | 235.20 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14000 | 129.43 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1253 | 1446.27 MB/s | 0 | 0 | 11.2× |
| Easyjson | 2559 | 708.08 MB/s | 24 | 1 | 5.5× |
| Goccy | 2812 | 644.32 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5909 | 306.68 MB/s | 3723 | 40 | 2.4× |
| SonicFastest | 5927 | 305.74 MB/s | 3756 | 40 | 2.4× |
| JSONV2 | 7767 | 233.30 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14025 | 129.19 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1443 | 1256.08 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2779 | 652.00 MB/s | 144 | 10 | 5.0× |
| Goccy | 2869 | 631.52 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6056 | 299.20 MB/s | 3710 | 42 | 2.3× |
| SonicFastest | 6084 | 297.83 MB/s | 3759 | 42 | 2.3× |
| JSONV2 | 7980 | 227.07 MB/s | 632 | 7 | 1.8× |
| Stdlib | 13969 | 129.72 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 745 | 663.44 MB/s | 160 | 1 | 7.4× |
| SonicFastest | 1231 | 401.44 MB/s | 972 | 6 | 4.5× |
| Sonic | 1232 | 401.13 MB/s | 969 | 6 | 4.5× |
| Easyjson | 2220 | 222.54 MB/s | 448 | 3 | 2.5× |
| Goccy | 2447 | 201.90 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3244 | 152.27 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5525 | 89.42 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 455 | 505.02 MB/s | 160 | 1 | 9.0× |
| Sonic | 876 | 262.54 MB/s | 655 | 6 | 4.7× |
| SonicFastest | 877 | 262.33 MB/s | 655 | 6 | 4.7× |
| Easyjson | 1402 | 164.08 MB/s | 448 | 3 | 2.9× |
| Goccy | 1576 | 145.93 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2398 | 95.90 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4107 | 56.00 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3379033 | 574.27 MB/s | 2846868 | 2238 | 7.0× |
| SonicFastest | 4485634 | 432.60 MB/s | 14607505 | 1407 | 5.2× |
| Sonic | 4610779 | 420.86 MB/s | 14610790 | 1407 | 5.1× |
| Goccy | 4780901 | 405.88 MB/s | 4065802 | 13510 | 4.9× |
| Easyjson | 7693847 | 252.21 MB/s | 3871268 | 15043 | 3.1× |
| JSONV2 | 11414920 | 169.99 MB/s | 3237333 | 13947 | 2.1× |
| Stdlib | 23549144 | 82.40 MB/s | 3551322 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13885328 | 583.35 MB/s | 15059835 | 51649 | 6.5× |
| Sonic | 17129486 | 472.87 MB/s | 72455349 | 40015 | 5.3× |
| SonicFastest | 17129754 | 472.86 MB/s | 72424176 | 40015 | 5.3× |
| Goccy | 24242433 | 334.13 MB/s | 17479729 | 107152 | 3.7× |
| Easyjson | 31212545 | 259.51 MB/s | 15059618 | 41643 | 2.9× |
| JSONV2 | 43884843 | 184.57 MB/s | 15233876 | 78973 | 2.1× |
| Stdlib | 90298107 | 89.70 MB/s | 15665088 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1565 | 1399.72 MB/s | 0 | 0 | 10.2× |
| Goccy | 3203 | 683.98 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3216 | 681.31 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6371 | 343.91 MB/s | 4073 | 40 | 2.5× |
| Sonic | 6391 | 342.84 MB/s | 4072 | 40 | 2.5× |
| JSONV2 | 8009 | 273.57 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15947 | 137.39 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2062 | 9597.03 MB/s | 0 | 0 | 54.2× |
| Goccy | 19993 | 989.81 MB/s | 20491 | 2 | 5.6× |
| Sonic | 27947 | 708.10 MB/s | 22227 | 4 | 4.0× |
| SonicFastest | 27948 | 708.05 MB/s | 22252 | 4 | 4.0× |
| JSONV2 | 29625 | 667.98 MB/s | 8 | 1 | 3.8× |
| Easyjson | 86043 | 229.99 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111797 | 177.01 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2968 | 6105.73 MB/s | 864 | 4 | 34.5× |
| Easyjson | 3952 | 4586.37 MB/s | 432 | 2 | 25.9× |
| Sonic | 9931 | 1825.05 MB/s | 22967 | 6 | 10.3× |
| SonicFastest | 9934 | 1824.39 MB/s | 22880 | 6 | 10.3× |
| Goccy | 15487 | 1170.25 MB/s | 19459 | 2 | 6.6× |
| JSONV2 | 45357 | 399.58 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102501 | 176.82 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2577600 | 779.21 MB/s | 2963984 | 7453 | 7.3× |
| Sonic | 4139950 | 485.15 MB/s | 10031901 | 13682 | 4.6× |
| SonicFastest | 4150465 | 483.92 MB/s | 10031952 | 13682 | 4.5× |
| Goccy | 4379206 | 458.64 MB/s | 5417352 | 15845 | 4.3× |
| Easyjson | 5120343 | 392.26 MB/s | 2981713 | 7441 | 3.7× |
| JSONV2 | 6994487 | 287.15 MB/s | 3173801 | 14563 | 2.7× |
| Stdlib | 18874159 | 106.42 MB/s | 3589321 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1127 | 487.34 MB/s | 480 | 1 | 5.0× |
| Easyjson | 2129 | 257.83 MB/s | 1616 | 5 | 2.7× |
| SonicFastest | 2667 | 205.82 MB/s | 1937 | 26 | 2.1× |
| Sonic | 2673 | 205.37 MB/s | 1947 | 26 | 2.1× |
| Goccy | 3024 | 181.58 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3263 | 168.28 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5665 | 96.90 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 573147 | 1101.84 MB/s | 461113 | 1230 | 9.3× |
| SonicFastest | 1008987 | 625.89 MB/s | 1007628 | 1102 | 5.3× |
| Sonic | 1011692 | 624.22 MB/s | 1009382 | 1102 | 5.3× |
| Easyjson | 1133220 | 557.27 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1172359 | 538.67 MB/s | 987699 | 1202 | 4.6× |
| JSONV2 | 2159989 | 292.37 MB/s | 571655 | 3144 | 2.5× |
| Stdlib | 5356279 | 117.90 MB/s | 654666 | 6472 | 1.0× |
