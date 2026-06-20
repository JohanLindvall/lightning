# JSON Deserialization Benchmarks

- generated 2026-06-20T13:10:14Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 186894 | 681.00 MB/s | 50208 | 4 | 6.6× |
| Easyjson | 234381 | 543.03 MB/s | 122864 | 14 | 5.3× |
| Goccy | 242025 | 525.88 MB/s | 225226 | 884 | 5.1× |
| Sonic | 262512 | 484.83 MB/s | 213922 | 15 | 4.7× |
| SonicFastest | 262566 | 484.74 MB/s | 214482 | 15 | 4.7× |
| JSONV2 | 395067 | 322.16 MB/s | 195125 | 1805 | 3.1× |
| Stdlib | 1238512 | 102.76 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4401163 | 511.47 MB/s | 3122874 | 3081 | 7.1× |
| SonicFastest | 5531700 | 406.94 MB/s | 4855239 | 2584 | 5.7× |
| Sonic | 5559937 | 404.87 MB/s | 4856450 | 2584 | 5.6× |
| Goccy | 13103207 | 171.79 MB/s | 4182620 | 56536 | 2.4× |
| Easyjson | 13916799 | 161.75 MB/s | 3099809 | 2120 | 2.2× |
| JSONV2 | 16837406 | 133.69 MB/s | 3123321 | 3083 | 1.9× |
| Stdlib | 31255523 | 72.02 MB/s | 3123395 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 553289 | 488.72 MB/s | 348024 | 1627 | 7.3× |
| SonicFastest | 733212 | 368.79 MB/s | 642207 | 1147 | 5.5× |
| Sonic | 743786 | 363.55 MB/s | 642635 | 1147 | 5.4× |
| Easyjson | 1745206 | 154.94 MB/s | 330272 | 749 | 2.3× |
| Goccy | 1779236 | 151.98 MB/s | 543680 | 8122 | 2.3× |
| JSONV2 | 2222675 | 121.66 MB/s | 348152 | 1628 | 1.8× |
| Stdlib | 4013754 | 67.37 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1416682 | 1219.19 MB/s | 964643 | 6177 | 11.9× |
| SonicFastest | 2354186 | 733.67 MB/s | 2701829 | 5548 | 7.2× |
| Sonic | 2420148 | 713.68 MB/s | 2701114 | 5548 | 7.0× |
| Goccy | 2783515 | 620.51 MB/s | 2589674 | 14606 | 6.1× |
| JSONV2 | 3986073 | 433.31 MB/s | 1011675 | 7594 | 4.2× |
| Easyjson | 4061300 | 425.28 MB/s | 972032 | 5389 | 4.2× |
| Stdlib | 16864362 | 102.42 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1052 | 1722.08 MB/s | 0 | 0 | 14.1× |
| Easyjson | 2710 | 668.74 MB/s | 24 | 1 | 5.5× |
| Goccy | 3358 | 539.53 MB/s | 2608 | 4 | 4.4× |
| SonicFastest | 6390 | 283.58 MB/s | 3346 | 38 | 2.3× |
| Sonic | 6457 | 280.61 MB/s | 3345 | 38 | 2.3× |
| JSONV2 | 7893 | 229.56 MB/s | 640 | 6 | 1.9× |
| Stdlib | 14819 | 122.27 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1109 | 1633.59 MB/s | 0 | 0 | 13.3× |
| Easyjson | 2723 | 665.51 MB/s | 24 | 1 | 5.4× |
| Goccy | 3284 | 551.80 MB/s | 2608 | 4 | 4.5× |
| SonicFastest | 5994 | 302.32 MB/s | 3349 | 38 | 2.5× |
| Sonic | 6157 | 294.28 MB/s | 3347 | 38 | 2.4× |
| JSONV2 | 7831 | 231.38 MB/s | 640 | 6 | 1.9× |
| Stdlib | 14783 | 122.57 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1443 | 1256.12 MB/s | 144 | 10 | 10.1× |
| Easyjson | 2964 | 611.38 MB/s | 144 | 10 | 4.9× |
| Goccy | 3313 | 546.94 MB/s | 2600 | 5 | 4.4× |
| SonicFastest | 6007 | 301.65 MB/s | 3368 | 40 | 2.4× |
| Sonic | 6218 | 291.40 MB/s | 3368 | 40 | 2.3× |
| JSONV2 | 8064 | 224.70 MB/s | 632 | 7 | 1.8× |
| Stdlib | 14567 | 124.39 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 732 | 674.43 MB/s | 160 | 1 | 8.3× |
| SonicFastest | 1211 | 407.91 MB/s | 1075 | 8 | 5.0× |
| Sonic | 1222 | 404.27 MB/s | 1075 | 8 | 4.9× |
| Easyjson | 2324 | 212.52 MB/s | 448 | 3 | 2.6× |
| Goccy | 2424 | 203.78 MB/s | 856 | 23 | 2.5× |
| JSONV2 | 3104 | 159.15 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6044 | 81.73 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 467 | 492.86 MB/s | 160 | 1 | 9.0× |
| SonicFastest | 856 | 268.66 MB/s | 801 | 8 | 4.9× |
| Sonic | 857 | 268.30 MB/s | 801 | 8 | 4.9× |
| Easyjson | 1538 | 149.51 MB/s | 448 | 3 | 2.7× |
| Goccy | 1665 | 138.14 MB/s | 584 | 23 | 2.5× |
| JSONV2 | 2329 | 98.76 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4209 | 54.64 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 97981 | 664.74 MB/s | 178609 | 112 | 6.6× |
| SonicFastest | 152081 | 428.27 MB/s | 235873 | 65 | 4.2× |
| Sonic | 157055 | 414.71 MB/s | 235863 | 65 | 4.1× |
| Goccy | 170160 | 382.77 MB/s | 228492 | 134 | 3.8× |
| JSONV2 | 244749 | 266.12 MB/s | 206661 | 607 | 2.6× |
| Stdlib | 642518 | 101.37 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2830936 | 685.45 MB/s | 2846867 | 2238 | 9.1× |
| Goccy | 5032678 | 385.57 MB/s | 4063358 | 13509 | 5.1× |
| Sonic | 5546950 | 349.83 MB/s | 4889609 | 1736 | 4.6× |
| SonicFastest | 5575327 | 348.05 MB/s | 4889799 | 1736 | 4.6× |
| Easyjson | 7883992 | 246.13 MB/s | 3871265 | 15043 | 3.3× |
| JSONV2 | 11431819 | 169.74 MB/s | 3237346 | 13947 | 2.2× |
| Stdlib | 25664181 | 75.61 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 2615551 | 1272.33 MB/s | 5874065 | 4263 | 9.8× |
| SonicFastest | 2643303 | 1258.97 MB/s | 5872607 | 4263 | 9.7× |
| Lightning | 2721004 | 1223.02 MB/s | 4611519 | 5958 | 9.4× |
| Goccy | 5379443 | 618.62 MB/s | 3948963 | 3817 | 4.7× |
| JSONV2 | 8625135 | 385.83 MB/s | 5364540 | 13243 | 3.0× |
| Stdlib | 25527202 | 130.36 MB/s | 5565611 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 268477 | 820.73 MB/s | 96320 | 224 | 8.2× |
| Goccy | 473583 | 465.27 MB/s | 365608 | 1067 | 4.7× |
| Sonic | 507883 | 433.85 MB/s | 351057 | 262 | 4.4× |
| SonicFastest | 509707 | 432.30 MB/s | 351317 | 262 | 4.3× |
| Easyjson | 585620 | 376.26 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 642628 | 342.88 MB/s | 129744 | 470 | 3.4× |
| Stdlib | 2212223 | 99.60 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 17692877 | 457.81 MB/s | 15059844 | 51649 | 5.8× |
| SonicFastest | 19327060 | 419.10 MB/s | 19877277 | 41641 | 5.3× |
| Sonic | 19376390 | 418.04 MB/s | 19882295 | 41641 | 5.3× |
| Goccy | 28320183 | 286.02 MB/s | 17448900 | 107151 | 3.6× |
| Easyjson | 36769498 | 220.29 MB/s | 15059620 | 41643 | 2.8× |
| JSONV2 | 48638001 | 166.54 MB/s | 15233927 | 78973 | 2.1× |
| Stdlib | 102511938 | 79.02 MB/s | 15665076 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6431189 | 463.91 MB/s | 3985337 | 30015 | 8.1× |
| Sonic | 9283823 | 321.36 MB/s | 9173654 | 57805 | 5.6× |
| SonicFastest | 9669372 | 308.55 MB/s | 9169751 | 57805 | 5.4× |
| Goccy | 18362884 | 162.47 MB/s | 10381202 | 273645 | 2.8× |
| Easyjson | 18705082 | 159.50 MB/s | 9479441 | 30115 | 2.8× |
| JSONV2 | 25706284 | 116.06 MB/s | 9257238 | 86279 | 2.0× |
| Stdlib | 51793732 | 57.60 MB/s | 9258090 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1452851 | 498.05 MB/s | 696699 | 3610 | 8.3× |
| Sonic | 2167295 | 333.87 MB/s | 2370114 | 3683 | 5.6× |
| SonicFastest | 2216333 | 326.48 MB/s | 2370383 | 3683 | 5.4× |
| Easyjson | 5094131 | 142.05 MB/s | 2847905 | 3698 | 2.4× |
| Goccy | 5381916 | 134.45 MB/s | 2913545 | 80278 | 2.2× |
| JSONV2 | 6224893 | 116.24 MB/s | 2704677 | 7318 | 1.9× |
| Stdlib | 12060283 | 60.00 MB/s | 2704552 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2167012 | 727.89 MB/s | 696697 | 3610 | 8.3× |
| SonicFastest | 2981692 | 529.01 MB/s | 3224877 | 3683 | 6.0× |
| Sonic | 3009758 | 524.08 MB/s | 3226316 | 3683 | 6.0× |
| Easyjson | 6761366 | 233.29 MB/s | 2847907 | 3698 | 2.7× |
| Goccy | 6984389 | 225.84 MB/s | 3547084 | 80263 | 2.6× |
| JSONV2 | 7417379 | 212.66 MB/s | 2704685 | 7318 | 2.4× |
| Stdlib | 17980632 | 87.73 MB/s | 2704553 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 235934 | 636.30 MB/s | 81920 | 1 | 9.4× |
| Sonic | 388840 | 386.08 MB/s | 409115 | 16 | 5.7× |
| SonicFastest | 390043 | 384.89 MB/s | 409294 | 16 | 5.7× |
| Goccy | 1006307 | 149.18 MB/s | 321438 | 10004 | 2.2× |
| JSONV2 | 1137593 | 131.97 MB/s | 357723 | 20 | 2.0× |
| Stdlib | 2219036 | 67.65 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 35736 | 786.81 MB/s | 29104 | 107 | 8.7× |
| Sonic | 66698 | 421.56 MB/s | 59497 | 83 | 4.7× |
| SonicFastest | 66951 | 419.96 MB/s | 59483 | 83 | 4.6× |
| Easyjson | 74031 | 379.80 MB/s | 32304 | 138 | 4.2× |
| Goccy | 79079 | 355.56 MB/s | 59273 | 188 | 3.9× |
| JSONV2 | 127656 | 220.26 MB/s | 36896 | 242 | 2.4× |
| Stdlib | 310652 | 90.51 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2045 | 1138.60 MB/s | 32 | 1 | 11.6× |
| Goccy | 4761 | 488.96 MB/s | 3649 | 4 | 5.0× |
| Easyjson | 4832 | 481.79 MB/s | 192 | 2 | 4.9× |
| Sonic | 6073 | 383.35 MB/s | 3711 | 4 | 3.9× |
| SonicFastest | 6087 | 382.48 MB/s | 3710 | 4 | 3.9× |
| JSONV2 | 7855 | 296.39 MB/s | 1000 | 6 | 3.0× |
| Stdlib | 23709 | 98.19 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224 | 841.69 MB/s | 0 | 0 | 10.9× |
| Goccy | 415 | 455.28 MB/s | 304 | 2 | 5.9× |
| Easyjson | 533 | 354.35 MB/s | 0 | 0 | 4.6× |
| SonicFastest | 806 | 234.55 MB/s | 341 | 3 | 3.0× |
| Sonic | 810 | 233.24 MB/s | 341 | 3 | 3.0× |
| JSONV2 | 906 | 208.51 MB/s | 112 | 1 | 2.7× |
| Stdlib | 2456 | 76.97 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1448 | 1513.04 MB/s | 0 | 0 | 11.9× |
| Easyjson | 3317 | 660.53 MB/s | 24 | 1 | 5.2× |
| Goccy | 3713 | 590.03 MB/s | 2864 | 4 | 4.6× |
| SonicFastest | 6584 | 332.76 MB/s | 3600 | 38 | 2.6× |
| Sonic | 6787 | 322.84 MB/s | 3599 | 38 | 2.5× |
| JSONV2 | 8014 | 273.39 MB/s | 640 | 6 | 2.2× |
| Stdlib | 17249 | 127.02 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 986903 | 517.25 MB/s | 703786 | 1012 | 6.4× |
| Goccy | 1252020 | 407.72 MB/s | 1136492 | 5006 | 5.1× |
| Easyjson | 1463351 | 348.84 MB/s | 863776 | 3012 | 4.3× |
| SonicFastest | 1466783 | 348.02 MB/s | 1315098 | 2014 | 4.3× |
| Sonic | 1470687 | 347.10 MB/s | 1313171 | 2014 | 4.3× |
| JSONV2 | 3130989 | 163.04 MB/s | 1076006 | 12646 | 2.0× |
| Stdlib | 6336907 | 80.56 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 645 | 30668.71 MB/s | 0 | 0 | 220.9× |
| SonicFastest | 7352 | 2691.71 MB/s | 21039 | 3 | 19.4× |
| Goccy | 26637 | 742.90 MB/s | 20492 | 2 | 5.4× |
| Sonic | 29186 | 678.04 MB/s | 20619 | 3 | 4.9× |
| JSONV2 | 35965 | 550.24 MB/s | 8 | 1 | 4.0× |
| Easyjson | 105664 | 187.28 MB/s | 0 | 0 | 1.3× |
| Stdlib | 142520 | 138.85 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2424 | 7478.05 MB/s | 864 | 4 | 49.6× |
| Easyjson | 4495 | 4032.18 MB/s | 432 | 2 | 26.8× |
| Sonic | 9405 | 1927.05 MB/s | 20411 | 5 | 12.8× |
| SonicFastest | 9418 | 1924.38 MB/s | 20425 | 5 | 12.8× |
| Goccy | 26238 | 690.75 MB/s | 19460 | 2 | 4.6× |
| JSONV2 | 48250 | 375.63 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 120340 | 150.61 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2701016 | 743.61 MB/s | 2963985 | 7453 | 7.8× |
| SonicFastest | 4469334 | 449.39 MB/s | 5179129 | 7086 | 4.7× |
| Sonic | 4522995 | 444.06 MB/s | 5175692 | 7086 | 4.7× |
| Goccy | 4741678 | 423.58 MB/s | 5414540 | 15843 | 4.4× |
| Easyjson | 5378719 | 373.41 MB/s | 2981698 | 7441 | 3.9× |
| JSONV2 | 7414462 | 270.89 MB/s | 3173778 | 14563 | 2.8× |
| Stdlib | 21071341 | 95.32 MB/s | 3589321 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 836 | 656.41 MB/s | 480 | 1 | 7.0× |
| Easyjson | 2021 | 271.58 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2125 | 258.36 MB/s | 2262 | 8 | 2.7× |
| Sonic | 2170 | 252.97 MB/s | 2262 | 8 | 2.7× |
| JSONV2 | 2942 | 186.61 MB/s | 1664 | 7 | 2.0× |
| Goccy | 2992 | 183.51 MB/s | 2129 | 43 | 1.9× |
| Stdlib | 5823 | 94.27 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 604706 | 1044.33 MB/s | 461113 | 1230 | 9.9× |
| Sonic | 1207002 | 523.21 MB/s | 1068990 | 814 | 5.0× |
| SonicFastest | 1222922 | 516.40 MB/s | 1069826 | 814 | 4.9× |
| Easyjson | 1291781 | 488.87 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1331026 | 474.46 MB/s | 995528 | 1202 | 4.5× |
| JSONV2 | 2174326 | 290.44 MB/s | 571617 | 3144 | 2.8× |
| Stdlib | 6013093 | 105.02 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 1553443 | 362.04 MB/s | 1416394 | 2239 | 3.7× |
| SonicFastest | 1558585 | 360.85 MB/s | 1416027 | 2239 | 3.6× |
| Goccy | 1702746 | 330.29 MB/s | 1115106 | 3404 | 3.3× |
| Lightning | 1736129 | 323.94 MB/s | 945690 | 4207 | 3.3× |
| Easyjson | 2168351 | 259.37 MB/s | 833619 | 3391 | 2.6× |
| JSONV2 | 3151330 | 178.47 MB/s | 979888 | 5184 | 1.8× |
| Stdlib | 5674491 | 99.11 MB/s | 1067140 | 8759 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1262930 | 422.18 MB/s | 935535 | 16692 | 5.0× |
| Sonic | 1738034 | 306.77 MB/s | 1389720 | 9201 | 3.6× |
| SonicFastest | 1741094 | 306.23 MB/s | 1389210 | 9201 | 3.6× |
| Easyjson | 1773197 | 300.69 MB/s | 949525 | 15985 | 3.5× |
| Goccy | 2157009 | 247.18 MB/s | 1668735 | 18049 | 2.9× |
| JSONV2 | 3283064 | 162.40 MB/s | 1192562 | 22667 | 1.9× |
| Stdlib | 6269868 | 85.04 MB/s | 1245606 | 26510 | 1.0× |
