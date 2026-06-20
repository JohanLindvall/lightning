# JSON Deserialization Benchmarks

- generated 2026-06-20T15:43:05Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 110692 | 1149.81 MB/s | 50208 | 4 | 11.7× |
| Sonic | 200892 | 633.55 MB/s | 213808 | 15 | 6.5× |
| SonicFastest | 201616 | 631.28 MB/s | 213873 | 15 | 6.4× |
| Easyjson | 239352 | 531.75 MB/s | 122864 | 14 | 5.4× |
| Goccy | 243950 | 521.73 MB/s | 224975 | 884 | 5.3× |
| JSONV2 | 453529 | 280.63 MB/s | 195125 | 1805 | 2.9× |
| Stdlib | 1298363 | 98.03 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4335733 | 519.19 MB/s | 3122876 | 3081 | 7.9× |
| SonicFastest | 5206769 | 432.33 MB/s | 4859071 | 2584 | 6.6× |
| Sonic | 5209601 | 432.10 MB/s | 4860592 | 2584 | 6.6× |
| Goccy | 13576920 | 165.80 MB/s | 4237705 | 56538 | 2.5× |
| Easyjson | 13814835 | 162.94 MB/s | 3099809 | 2120 | 2.5× |
| JSONV2 | 18040167 | 124.78 MB/s | 3123321 | 3083 | 1.9× |
| Stdlib | 34146526 | 65.92 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 566536 | 477.29 MB/s | 348024 | 1627 | 7.7× |
| Sonic | 743122 | 363.87 MB/s | 642177 | 1147 | 5.9× |
| SonicFastest | 752127 | 359.52 MB/s | 642528 | 1147 | 5.8× |
| Easyjson | 1775690 | 152.28 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1854008 | 145.85 MB/s | 542383 | 8122 | 2.4× |
| JSONV2 | 2356382 | 114.75 MB/s | 348156 | 1628 | 1.9× |
| Stdlib | 4373879 | 61.82 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1441670 | 1198.06 MB/s | 964643 | 6177 | 12.3× |
| SonicFastest | 2105610 | 820.29 MB/s | 2706879 | 5548 | 8.4× |
| Sonic | 2106488 | 819.95 MB/s | 2705503 | 5548 | 8.4× |
| Goccy | 2778744 | 621.58 MB/s | 2590191 | 14606 | 6.4× |
| Easyjson | 4302322 | 401.46 MB/s | 972032 | 5389 | 4.1× |
| JSONV2 | 4740625 | 364.34 MB/s | 1011683 | 7594 | 3.7× |
| Stdlib | 17756286 | 97.27 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1047 | 1731.21 MB/s | 0 | 0 | 15.4× |
| Easyjson | 2973 | 609.54 MB/s | 24 | 1 | 5.4× |
| Goccy | 3161 | 573.20 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 6186 | 292.94 MB/s | 3346 | 38 | 2.6× |
| Sonic | 6378 | 284.11 MB/s | 3346 | 38 | 2.5× |
| JSONV2 | 8187 | 221.31 MB/s | 640 | 6 | 2.0× |
| Stdlib | 16076 | 112.72 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1137 | 1593.59 MB/s | 0 | 0 | 14.2× |
| Easyjson | 2987 | 606.58 MB/s | 24 | 1 | 5.4× |
| Goccy | 3175 | 570.74 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5975 | 303.29 MB/s | 3344 | 38 | 2.7× |
| Sonic | 6227 | 290.98 MB/s | 3346 | 38 | 2.6× |
| JSONV2 | 8048 | 225.16 MB/s | 640 | 6 | 2.0× |
| Stdlib | 16151 | 112.19 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1365 | 1327.59 MB/s | 144 | 10 | 11.8× |
| Goccy | 3171 | 571.40 MB/s | 2600 | 5 | 5.1× |
| Easyjson | 3278 | 552.81 MB/s | 144 | 10 | 4.9× |
| SonicFastest | 6121 | 296.01 MB/s | 3365 | 40 | 2.6× |
| Sonic | 6442 | 281.28 MB/s | 3366 | 40 | 2.5× |
| JSONV2 | 8122 | 223.09 MB/s | 632 | 7 | 2.0× |
| Stdlib | 16066 | 112.79 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 739 | 668.51 MB/s | 160 | 1 | 9.4× |
| Sonic | 1287 | 383.79 MB/s | 1076 | 8 | 5.4× |
| SonicFastest | 1289 | 383.29 MB/s | 1075 | 8 | 5.4× |
| Easyjson | 2541 | 194.44 MB/s | 448 | 3 | 2.7× |
| Goccy | 2542 | 194.30 MB/s | 856 | 23 | 2.7× |
| JSONV2 | 3482 | 141.87 MB/s | 528 | 7 | 2.0× |
| Stdlib | 6918 | 71.41 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 491 | 468.84 MB/s | 160 | 1 | 10.0× |
| SonicFastest | 924 | 248.90 MB/s | 801 | 8 | 5.3× |
| Sonic | 926 | 248.24 MB/s | 801 | 8 | 5.3× |
| Easyjson | 1701 | 135.17 MB/s | 448 | 3 | 2.9× |
| Goccy | 1766 | 130.22 MB/s | 584 | 23 | 2.8× |
| JSONV2 | 2624 | 87.65 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4924 | 46.71 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 95146 | 684.55 MB/s | 178609 | 112 | 7.1× |
| SonicFastest | 138363 | 470.73 MB/s | 235439 | 65 | 4.9× |
| Sonic | 139998 | 465.24 MB/s | 235470 | 65 | 4.8× |
| Goccy | 183291 | 355.35 MB/s | 228043 | 134 | 3.7× |
| JSONV2 | 257915 | 252.53 MB/s | 206660 | 607 | 2.6× |
| Stdlib | 677681 | 96.11 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2828031 | 686.16 MB/s | 2846866 | 2238 | 10.1× |
| Goccy | 5159733 | 376.08 MB/s | 4062053 | 13509 | 5.6× |
| Sonic | 5836572 | 332.47 MB/s | 4886585 | 1736 | 4.9× |
| SonicFastest | 5865324 | 330.84 MB/s | 4886550 | 1736 | 4.9× |
| Easyjson | 8431616 | 230.14 MB/s | 3871267 | 15043 | 3.4× |
| JSONV2 | 12973456 | 149.57 MB/s | 3237342 | 13947 | 2.2× |
| Stdlib | 28661386 | 67.70 MB/s | 3551321 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| SonicFastest | 2375346 | 1400.99 MB/s | 5876507 | 4263 | 10.1× |
| Sonic | 2396898 | 1388.39 MB/s | 5874749 | 4263 | 10.0× |
| Lightning | 2485887 | 1338.69 MB/s | 4611522 | 5958 | 9.6× |
| Goccy | 4796320 | 693.83 MB/s | 3948963 | 3817 | 5.0× |
| JSONV2 | 7896055 | 421.45 MB/s | 5364545 | 13243 | 3.0× |
| Stdlib | 23889621 | 139.30 MB/s | 5565615 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 251997 | 874.40 MB/s | 123584 | 225 | 9.7× |
| SonicFastest | 461836 | 477.11 MB/s | 350837 | 262 | 5.3× |
| Sonic | 463451 | 475.45 MB/s | 351362 | 262 | 5.3× |
| Goccy | 489290 | 450.34 MB/s | 365689 | 1067 | 5.0× |
| Easyjson | 625311 | 352.38 MB/s | 130512 | 245 | 3.9× |
| JSONV2 | 726738 | 303.20 MB/s | 129745 | 470 | 3.4× |
| Stdlib | 2435828 | 90.46 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 16612307 | 487.59 MB/s | 19913976 | 41641 | 6.8× |
| SonicFastest | 16628247 | 487.13 MB/s | 19913975 | 41641 | 6.8× |
| Lightning | 17426323 | 464.82 MB/s | 15059840 | 51649 | 6.5× |
| Goccy | 29529162 | 274.31 MB/s | 17855593 | 107152 | 3.8× |
| Easyjson | 37950618 | 213.44 MB/s | 15059618 | 41643 | 3.0× |
| JSONV2 | 50158612 | 161.49 MB/s | 15233901 | 78973 | 2.3× |
| Stdlib | 112912455 | 71.74 MB/s | 15665079 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6585359 | 453.05 MB/s | 3985340 | 30015 | 9.1× |
| SonicFastest | 9323500 | 319.99 MB/s | 9180829 | 57805 | 6.4× |
| Sonic | 9538586 | 312.78 MB/s | 9174260 | 57805 | 6.3× |
| Goccy | 18702590 | 159.52 MB/s | 10206066 | 273637 | 3.2× |
| Easyjson | 20076107 | 148.61 MB/s | 9479442 | 30115 | 3.0× |
| JSONV2 | 27982461 | 106.62 MB/s | 9257222 | 86279 | 2.1× |
| Stdlib | 59709521 | 49.97 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1550342 | 466.73 MB/s | 907387 | 3618 | 9.0× |
| SonicFastest | 2215879 | 326.55 MB/s | 2376509 | 3683 | 6.3× |
| Sonic | 2230828 | 324.36 MB/s | 2376921 | 3683 | 6.2× |
| Easyjson | 5233602 | 138.26 MB/s | 2847904 | 3698 | 2.7× |
| Goccy | 5884376 | 122.97 MB/s | 2941239 | 80280 | 2.4× |
| JSONV2 | 6444073 | 112.29 MB/s | 2704691 | 7318 | 2.2× |
| Stdlib | 13875688 | 52.15 MB/s | 2704553 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2158515 | 730.76 MB/s | 907389 | 3618 | 9.5× |
| SonicFastest | 2955475 | 533.71 MB/s | 3227646 | 3683 | 7.0× |
| Sonic | 2960809 | 532.74 MB/s | 3228101 | 3683 | 6.9× |
| Easyjson | 6671200 | 236.44 MB/s | 2847906 | 3698 | 3.1× |
| JSONV2 | 6917225 | 228.03 MB/s | 2704691 | 7318 | 3.0× |
| Goccy | 6978166 | 226.04 MB/s | 3598539 | 80267 | 2.9× |
| Stdlib | 20565363 | 76.70 MB/s | 2704554 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 248842 | 603.29 MB/s | 81920 | 1 | 8.7× |
| SonicFastest | 383422 | 391.54 MB/s | 409204 | 16 | 5.6× |
| Sonic | 384044 | 390.90 MB/s | 409191 | 16 | 5.6× |
| Goccy | 1022289 | 146.85 MB/s | 327298 | 10005 | 2.1× |
| JSONV2 | 1144347 | 131.19 MB/s | 357724 | 20 | 1.9× |
| Stdlib | 2156361 | 69.62 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 37947 | 740.95 MB/s | 29104 | 107 | 9.0× |
| SonicFastest | 57918 | 485.46 MB/s | 59530 | 83 | 5.9× |
| Sonic | 57964 | 485.08 MB/s | 59545 | 83 | 5.9× |
| Goccy | 78196 | 359.57 MB/s | 59281 | 188 | 4.3× |
| Easyjson | 81106 | 346.67 MB/s | 32304 | 138 | 4.2× |
| JSONV2 | 134433 | 209.15 MB/s | 36897 | 242 | 2.5× |
| Stdlib | 339948 | 82.71 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2034 | 1144.49 MB/s | 32 | 1 | 12.8× |
| Sonic | 4654 | 500.26 MB/s | 3708 | 4 | 5.6× |
| SonicFastest | 4685 | 496.96 MB/s | 3711 | 4 | 5.5× |
| Goccy | 4771 | 487.99 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 5417 | 429.73 MB/s | 192 | 2 | 4.8× |
| JSONV2 | 8080 | 288.13 MB/s | 1000 | 6 | 3.2× |
| Stdlib | 25948 | 89.72 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225 | 838.73 MB/s | 0 | 0 | 12.3× |
| Goccy | 446 | 423.81 MB/s | 304 | 2 | 6.2× |
| Easyjson | 554 | 341.27 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 642 | 294.14 MB/s | 341 | 3 | 4.3× |
| Sonic | 646 | 292.44 MB/s | 341 | 3 | 4.3× |
| JSONV2 | 1044 | 180.99 MB/s | 112 | 1 | 2.7× |
| Stdlib | 2768 | 68.28 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1449 | 1512.03 MB/s | 0 | 0 | 13.1× |
| Goccy | 3578 | 612.44 MB/s | 2864 | 4 | 5.3× |
| Easyjson | 3871 | 566.03 MB/s | 24 | 1 | 4.9× |
| SonicFastest | 6489 | 337.65 MB/s | 3602 | 38 | 2.9× |
| Sonic | 6741 | 325.03 MB/s | 3601 | 38 | 2.8× |
| JSONV2 | 8252 | 265.50 MB/s | 640 | 6 | 2.3× |
| Stdlib | 19008 | 115.27 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1032629 | 494.35 MB/s | 703781 | 1012 | 6.7× |
| Goccy | 1351796 | 377.63 MB/s | 1144065 | 5007 | 5.2× |
| Sonic | 1411512 | 361.65 MB/s | 1316117 | 2014 | 4.9× |
| SonicFastest | 1415646 | 360.60 MB/s | 1315680 | 2014 | 4.9× |
| Easyjson | 1570908 | 324.96 MB/s | 863776 | 3012 | 4.4× |
| JSONV2 | 3494242 | 146.09 MB/s | 1076014 | 12646 | 2.0× |
| Stdlib | 6964759 | 73.29 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 645 | 30676.76 MB/s | 0 | 0 | 252.2× |
| SonicFastest | 6309 | 3136.88 MB/s | 21151 | 3 | 25.8× |
| Goccy | 29496 | 670.91 MB/s | 20492 | 2 | 5.5× |
| Sonic | 32301 | 612.64 MB/s | 20632 | 3 | 5.0× |
| JSONV2 | 33275 | 594.72 MB/s | 8 | 1 | 4.9× |
| Easyjson | 105862 | 186.93 MB/s | 0 | 0 | 1.5× |
| Stdlib | 162697 | 121.63 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2434 | 7444.69 MB/s | 864 | 4 | 48.8× |
| Easyjson | 4754 | 3812.50 MB/s | 432 | 2 | 25.0× |
| Sonic | 8120 | 2231.93 MB/s | 20478 | 5 | 14.6× |
| SonicFastest | 8499 | 2132.50 MB/s | 20440 | 5 | 14.0× |
| Goccy | 25437 | 712.49 MB/s | 19460 | 2 | 4.7× |
| JSONV2 | 50282 | 360.45 MB/s | 16501 | 50 | 2.4× |
| Stdlib | 118730 | 152.65 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2813187 | 713.96 MB/s | 2963985 | 7453 | 7.9× |
| SonicFastest | 3673164 | 546.80 MB/s | 5197549 | 7086 | 6.1× |
| Sonic | 3684454 | 545.13 MB/s | 5196861 | 7086 | 6.0× |
| Goccy | 4848669 | 414.24 MB/s | 5414666 | 15844 | 4.6× |
| Easyjson | 5868594 | 342.24 MB/s | 2981694 | 7441 | 3.8× |
| JSONV2 | 7830342 | 256.50 MB/s | 3173783 | 14563 | 2.8× |
| Stdlib | 22281284 | 90.14 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 868 | 632.28 MB/s | 480 | 1 | 7.8× |
| Easyjson | 2344 | 234.17 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2348 | 233.85 MB/s | 2259 | 8 | 2.9× |
| Sonic | 2374 | 231.23 MB/s | 2260 | 8 | 2.9× |
| JSONV2 | 3218 | 170.62 MB/s | 1664 | 7 | 2.1× |
| Goccy | 3386 | 162.12 MB/s | 2128 | 43 | 2.0× |
| Stdlib | 6797 | 80.77 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 614103 | 1028.35 MB/s | 461113 | 1230 | 10.3× |
| Sonic | 1117480 | 565.12 MB/s | 1069781 | 814 | 5.7× |
| SonicFastest | 1119234 | 564.24 MB/s | 1070655 | 814 | 5.7× |
| Goccy | 1325018 | 476.61 MB/s | 993777 | 1202 | 4.8× |
| Easyjson | 1443855 | 437.38 MB/s | 422505 | 936 | 4.4× |
| JSONV2 | 2330977 | 270.92 MB/s | 571620 | 3144 | 2.7× |
| Stdlib | 6352156 | 99.42 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1153573 | 487.54 MB/s | 886433 | 2062 | 5.2× |
| Sonic | 1385219 | 406.01 MB/s | 1351350 | 1185 | 4.3× |
| SonicFastest | 1400071 | 401.70 MB/s | 1351415 | 1185 | 4.3× |
| Goccy | 1520789 | 369.81 MB/s | 1041256 | 1029 | 3.9× |
| Easyjson | 2152833 | 261.24 MB/s | 775153 | 1254 | 2.8× |
| JSONV2 | 3190368 | 176.28 MB/s | 927455 | 3482 | 1.9× |
| Stdlib | 5960313 | 94.36 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 802772 | 664.17 MB/s | 413026 | 3155 | 7.9× |
| SonicFastest | 1048737 | 508.40 MB/s | 973241 | 3083 | 6.0× |
| Sonic | 1052915 | 506.38 MB/s | 973784 | 3083 | 6.0× |
| Easyjson | 1256370 | 424.38 MB/s | 428362 | 3273 | 5.0× |
| Goccy | 1540904 | 346.02 MB/s | 1167282 | 5409 | 4.1× |
| JSONV2 | 2863395 | 186.20 MB/s | 745485 | 13288 | 2.2× |
| Stdlib | 6316323 | 84.41 MB/s | 798693 | 17133 | 1.0× |
