# JSON Deserialization Benchmarks

- generated 2026-06-19T22:02:35Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 550251 | 491.42 MB/s | 348024 | 1627 | 7.3× |
| Sonic | 727100 | 371.89 MB/s | 642128 | 1147 | 5.5× |
| SonicFastest | 734127 | 368.33 MB/s | 642252 | 1147 | 5.5× |
| Goccy | 1701136 | 158.95 MB/s | 544569 | 8123 | 2.4× |
| Easyjson | 1725899 | 156.67 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2292634 | 117.94 MB/s | 348153 | 1628 | 1.8× |
| Stdlib | 4014518 | 67.36 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1419254 | 1216.98 MB/s | 964643 | 6177 | 11.9× |
| Sonic | 2080508 | 830.18 MB/s | 2698670 | 5548 | 8.1× |
| SonicFastest | 2097271 | 823.55 MB/s | 2699883 | 5548 | 8.1× |
| Goccy | 2568027 | 672.58 MB/s | 2589175 | 14606 | 6.6× |
| Easyjson | 3767406 | 458.46 MB/s | 972032 | 5389 | 4.5× |
| JSONV2 | 4256403 | 405.79 MB/s | 1011674 | 7594 | 4.0× |
| Stdlib | 16883916 | 102.30 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1034 | 1752.77 MB/s | 0 | 0 | 14.2× |
| Easyjson | 2834 | 639.27 MB/s | 24 | 1 | 5.2× |
| Goccy | 3075 | 589.32 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 6250 | 289.94 MB/s | 3346 | 38 | 2.4× |
| Sonic | 6442 | 281.28 MB/s | 3343 | 38 | 2.3× |
| JSONV2 | 7488 | 241.98 MB/s | 640 | 6 | 2.0× |
| Stdlib | 14694 | 123.31 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1101 | 1645.56 MB/s | 0 | 0 | 13.3× |
| Easyjson | 2809 | 645.02 MB/s | 24 | 1 | 5.2× |
| Goccy | 3049 | 594.38 MB/s | 2608 | 4 | 4.8× |
| SonicFastest | 5880 | 308.15 MB/s | 3349 | 38 | 2.5× |
| Sonic | 6145 | 294.88 MB/s | 3345 | 38 | 2.4× |
| JSONV2 | 7519 | 240.99 MB/s | 640 | 6 | 1.9× |
| Stdlib | 14640 | 123.77 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1296 | 1397.97 MB/s | 144 | 10 | 11.3× |
| Easyjson | 3048 | 594.54 MB/s | 144 | 10 | 4.8× |
| Goccy | 3112 | 582.29 MB/s | 2600 | 5 | 4.7× |
| SonicFastest | 6076 | 298.20 MB/s | 3369 | 40 | 2.4× |
| Sonic | 6286 | 288.28 MB/s | 3368 | 40 | 2.3× |
| JSONV2 | 7712 | 234.95 MB/s | 632 | 7 | 1.9× |
| Stdlib | 14606 | 124.05 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 730 | 677.19 MB/s | 160 | 1 | 8.5× |
| Sonic | 1220 | 404.98 MB/s | 1075 | 8 | 5.1× |
| SonicFastest | 1224 | 403.66 MB/s | 1075 | 8 | 5.1× |
| Easyjson | 2392 | 206.54 MB/s | 448 | 3 | 2.6× |
| Goccy | 2593 | 190.54 MB/s | 856 | 23 | 2.4× |
| JSONV2 | 2995 | 164.93 MB/s | 528 | 7 | 2.1× |
| Stdlib | 6183 | 79.90 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 475 | 484.47 MB/s | 160 | 1 | 8.9× |
| Sonic | 857 | 268.38 MB/s | 801 | 8 | 4.9× |
| SonicFastest | 863 | 266.47 MB/s | 801 | 8 | 4.9× |
| Easyjson | 1564 | 147.04 MB/s | 448 | 3 | 2.7× |
| Goccy | 1709 | 134.57 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2283 | 100.73 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4225 | 54.44 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2811235 | 690.26 MB/s | 2846867 | 2238 | 9.2× |
| Goccy | 4993083 | 388.63 MB/s | 4063195 | 13509 | 5.2× |
| Sonic | 5409860 | 358.69 MB/s | 4888224 | 1736 | 4.8× |
| SonicFastest | 5431885 | 357.24 MB/s | 4887919 | 1736 | 4.8× |
| Easyjson | 7810630 | 248.44 MB/s | 3871265 | 15043 | 3.3× |
| JSONV2 | 11338310 | 171.14 MB/s | 3237339 | 13947 | 2.3× |
| Stdlib | 25860559 | 75.04 MB/s | 3551323 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 17822417 | 454.49 MB/s | 15059837 | 51649 | 5.7× |
| SonicFastest | 19780660 | 409.49 MB/s | 19893141 | 41641 | 5.2× |
| Sonic | 19828302 | 408.51 MB/s | 19898238 | 41641 | 5.2× |
| Goccy | 27503678 | 294.51 MB/s | 17485112 | 107151 | 3.7× |
| Easyjson | 38062605 | 212.81 MB/s | 15059620 | 41643 | 2.7× |
| JSONV2 | 48032017 | 168.64 MB/s | 15233883 | 78973 | 2.1× |
| Stdlib | 102256604 | 79.21 MB/s | 15665078 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1447 | 1514.39 MB/s | 0 | 0 | 12.0× |
| Goccy | 3426 | 639.59 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3510 | 624.25 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6688 | 327.62 MB/s | 3603 | 38 | 2.6× |
| Sonic | 6850 | 319.86 MB/s | 3601 | 38 | 2.5× |
| JSONV2 | 7629 | 287.19 MB/s | 640 | 6 | 2.3× |
| Stdlib | 17402 | 125.91 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 647 | 30593.93 MB/s | 0 | 0 | 220.9× |
| SonicFastest | 7118 | 2780.32 MB/s | 21068 | 3 | 20.1× |
| Goccy | 25190 | 785.59 MB/s | 20492 | 2 | 5.7× |
| Sonic | 29538 | 669.96 MB/s | 20642 | 3 | 4.8× |
| JSONV2 | 36208 | 546.53 MB/s | 8 | 1 | 3.9× |
| Easyjson | 118079 | 167.59 MB/s | 0 | 0 | 1.2× |
| Stdlib | 142892 | 138.49 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2605 | 6958.07 MB/s | 864 | 4 | 46.5× |
| Easyjson | 4665 | 3885.22 MB/s | 432 | 2 | 25.9× |
| Sonic | 9157 | 1979.30 MB/s | 20449 | 5 | 13.2× |
| SonicFastest | 9296 | 1949.60 MB/s | 20434 | 5 | 13.0× |
| Goccy | 22302 | 812.67 MB/s | 19460 | 2 | 5.4× |
| JSONV2 | 48304 | 375.21 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 121028 | 149.75 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2725112 | 737.03 MB/s | 2963983 | 7453 | 7.7× |
| Sonic | 4512745 | 445.07 MB/s | 5172640 | 7086 | 4.7× |
| SonicFastest | 4564738 | 440.00 MB/s | 5173515 | 7086 | 4.6× |
| Goccy | 4785012 | 419.75 MB/s | 5415046 | 15844 | 4.4× |
| Easyjson | 5594808 | 358.99 MB/s | 2981692 | 7441 | 3.8× |
| JSONV2 | 6926162 | 289.99 MB/s | 3173782 | 14563 | 3.0× |
| Stdlib | 21069717 | 95.33 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 848 | 647.47 MB/s | 480 | 1 | 6.8× |
| Easyjson | 1990 | 275.83 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2139 | 256.62 MB/s | 2262 | 8 | 2.7× |
| Sonic | 2186 | 251.18 MB/s | 2262 | 8 | 2.6× |
| JSONV2 | 3039 | 180.63 MB/s | 1664 | 7 | 1.9× |
| Goccy | 3046 | 180.26 MB/s | 2129 | 43 | 1.9× |
| Stdlib | 5784 | 94.92 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 598027 | 1055.99 MB/s | 461113 | 1230 | 10.0× |
| Sonic | 1218477 | 518.28 MB/s | 1069947 | 814 | 4.9× |
| SonicFastest | 1234805 | 511.43 MB/s | 1069835 | 814 | 4.8× |
| Easyjson | 1245500 | 507.04 MB/s | 422505 | 936 | 4.8× |
| Goccy | 1418146 | 445.31 MB/s | 994723 | 1202 | 4.2× |
| JSONV2 | 2251968 | 280.43 MB/s | 571619 | 3144 | 2.7× |
| Stdlib | 5974040 | 105.71 MB/s | 654666 | 6472 | 1.0× |
