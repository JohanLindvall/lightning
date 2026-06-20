# JSON Deserialization Benchmarks

- generated 2026-06-20T14:15:52Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 192294 | 661.88 MB/s | 50208 | 4 | 6.9× |
| SonicFastest | 202588 | 628.24 MB/s | 213977 | 15 | 6.6× |
| Sonic | 207782 | 612.54 MB/s | 214670 | 15 | 6.4× |
| Goccy | 245077 | 519.33 MB/s | 225334 | 884 | 5.4× |
| Easyjson | 263369 | 483.26 MB/s | 122864 | 14 | 5.1× |
| JSONV2 | 452439 | 281.31 MB/s | 195125 | 1805 | 2.9× |
| Stdlib | 1330891 | 95.63 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4420731 | 509.20 MB/s | 3122874 | 3081 | 7.8× |
| SonicFastest | 5266666 | 427.41 MB/s | 4859839 | 2584 | 6.5× |
| Sonic | 5299459 | 424.77 MB/s | 4859560 | 2584 | 6.5× |
| Goccy | 13522787 | 166.46 MB/s | 4241681 | 56539 | 2.5× |
| Easyjson | 14022560 | 160.53 MB/s | 3099808 | 2120 | 2.5× |
| JSONV2 | 17890067 | 125.83 MB/s | 3123320 | 3083 | 1.9× |
| Stdlib | 34465620 | 65.31 MB/s | 3123393 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 567744 | 476.28 MB/s | 348024 | 1627 | 7.8× |
| Sonic | 759170 | 356.18 MB/s | 641776 | 1147 | 5.8× |
| SonicFastest | 763123 | 354.34 MB/s | 641930 | 1147 | 5.8× |
| Easyjson | 1830070 | 147.76 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1858110 | 145.53 MB/s | 543028 | 8122 | 2.4× |
| JSONV2 | 2351455 | 114.99 MB/s | 348155 | 1628 | 1.9× |
| Stdlib | 4400181 | 61.45 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1456938 | 1185.50 MB/s | 964644 | 6177 | 12.2× |
| SonicFastest | 2088408 | 827.04 MB/s | 2706334 | 5548 | 8.5× |
| Sonic | 2092442 | 825.45 MB/s | 2706783 | 5548 | 8.5× |
| Goccy | 2854895 | 605.00 MB/s | 2589922 | 14606 | 6.2× |
| Easyjson | 4368511 | 395.38 MB/s | 972032 | 5389 | 4.1× |
| JSONV2 | 4737400 | 364.59 MB/s | 1011678 | 7594 | 3.7× |
| Stdlib | 17756668 | 97.27 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1056 | 1715.29 MB/s | 0 | 0 | 15.3× |
| Easyjson | 2874 | 630.44 MB/s | 24 | 1 | 5.6× |
| Goccy | 3549 | 510.55 MB/s | 2608 | 4 | 4.5× |
| SonicFastest | 6443 | 281.25 MB/s | 3346 | 38 | 2.5× |
| Sonic | 6562 | 276.12 MB/s | 3348 | 38 | 2.5× |
| JSONV2 | 8386 | 216.08 MB/s | 640 | 6 | 1.9× |
| Stdlib | 16124 | 112.38 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1145 | 1582.64 MB/s | 0 | 0 | 14.2× |
| Easyjson | 2844 | 637.15 MB/s | 24 | 1 | 5.7× |
| Goccy | 3498 | 518.06 MB/s | 2608 | 4 | 4.6× |
| SonicFastest | 6079 | 298.06 MB/s | 3344 | 38 | 2.7× |
| Sonic | 6279 | 288.57 MB/s | 3345 | 38 | 2.6× |
| JSONV2 | 8374 | 216.39 MB/s | 640 | 6 | 1.9× |
| Stdlib | 16257 | 111.46 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1383 | 1310.65 MB/s | 144 | 10 | 11.6× |
| Easyjson | 3115 | 581.76 MB/s | 144 | 10 | 5.1× |
| Goccy | 3506 | 516.90 MB/s | 2600 | 5 | 4.6× |
| SonicFastest | 6334 | 286.09 MB/s | 3366 | 40 | 2.5× |
| Sonic | 6559 | 276.26 MB/s | 3367 | 40 | 2.4× |
| JSONV2 | 8597 | 210.76 MB/s | 632 | 7 | 1.9× |
| Stdlib | 16040 | 112.97 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 754 | 655.65 MB/s | 160 | 1 | 9.3× |
| SonicFastest | 1277 | 386.83 MB/s | 1075 | 8 | 5.5× |
| Sonic | 1287 | 383.79 MB/s | 1075 | 8 | 5.5× |
| Easyjson | 2562 | 192.81 MB/s | 448 | 3 | 2.7× |
| Goccy | 2583 | 191.27 MB/s | 856 | 23 | 2.7× |
| JSONV2 | 3502 | 141.06 MB/s | 528 | 7 | 2.0× |
| Stdlib | 7025 | 70.32 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 497 | 462.71 MB/s | 160 | 1 | 9.9× |
| SonicFastest | 923 | 249.26 MB/s | 801 | 8 | 5.4× |
| Sonic | 929 | 247.59 MB/s | 801 | 8 | 5.3× |
| Easyjson | 1696 | 135.62 MB/s | 448 | 3 | 2.9× |
| Goccy | 1795 | 128.13 MB/s | 584 | 23 | 2.8× |
| JSONV2 | 2668 | 86.22 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4942 | 46.54 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 97196 | 670.11 MB/s | 178609 | 112 | 7.1× |
| Sonic | 143603 | 453.56 MB/s | 235252 | 65 | 4.8× |
| SonicFastest | 154476 | 421.63 MB/s | 236072 | 65 | 4.5× |
| Goccy | 199598 | 326.32 MB/s | 228673 | 134 | 3.5× |
| JSONV2 | 279516 | 233.02 MB/s | 206660 | 607 | 2.5× |
| Stdlib | 694946 | 93.72 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2979742 | 651.22 MB/s | 2846866 | 2238 | 9.5× |
| Goccy | 5183421 | 374.36 MB/s | 4063043 | 13509 | 5.5× |
| SonicFastest | 5327184 | 364.26 MB/s | 4886860 | 1736 | 5.3× |
| Sonic | 5338181 | 363.51 MB/s | 4888287 | 1736 | 5.3× |
| Easyjson | 8453565 | 229.54 MB/s | 3871265 | 15043 | 3.4× |
| JSONV2 | 13396244 | 144.85 MB/s | 3237333 | 13947 | 2.1× |
| Stdlib | 28418773 | 68.28 MB/s | 3551321 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| SonicFastest | 2433328 | 1367.60 MB/s | 5875603 | 4263 | 10.2× |
| Sonic | 2511843 | 1324.86 MB/s | 5874121 | 4263 | 9.9× |
| Lightning | 2609858 | 1275.10 MB/s | 4611525 | 5958 | 9.5× |
| Goccy | 4795457 | 693.95 MB/s | 3948963 | 3817 | 5.2× |
| JSONV2 | 8336131 | 399.21 MB/s | 5364533 | 13243 | 3.0× |
| Stdlib | 24752881 | 134.44 MB/s | 5565613 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 270362 | 815.00 MB/s | 96320 | 224 | 9.1× |
| SonicFastest | 465679 | 473.17 MB/s | 350613 | 262 | 5.3× |
| Sonic | 467131 | 471.70 MB/s | 350668 | 262 | 5.2× |
| Goccy | 497629 | 442.79 MB/s | 365604 | 1067 | 4.9× |
| Easyjson | 642029 | 343.20 MB/s | 130512 | 245 | 3.8× |
| JSONV2 | 768694 | 286.65 MB/s | 129744 | 470 | 3.2× |
| Stdlib | 2449097 | 89.97 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| SonicFastest | 16334315 | 495.89 MB/s | 19907599 | 41641 | 7.0× |
| Sonic | 16475495 | 491.64 MB/s | 19909779 | 41641 | 6.9× |
| Lightning | 17456993 | 464.00 MB/s | 15059836 | 51649 | 6.5× |
| Goccy | 29721962 | 272.53 MB/s | 18021151 | 107152 | 3.8× |
| Easyjson | 37534472 | 215.80 MB/s | 15059619 | 41643 | 3.0× |
| JSONV2 | 49610067 | 163.27 MB/s | 15233901 | 78973 | 2.3× |
| Stdlib | 113799024 | 71.18 MB/s | 15665079 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6825161 | 437.13 MB/s | 3985336 | 30015 | 8.7× |
| Sonic | 9590296 | 311.09 MB/s | 9179580 | 57805 | 6.2× |
| SonicFastest | 9955293 | 299.69 MB/s | 9158395 | 57805 | 6.0× |
| Goccy | 19257837 | 154.92 MB/s | 10329927 | 273642 | 3.1× |
| Easyjson | 20736354 | 143.88 MB/s | 9479442 | 30115 | 2.9× |
| JSONV2 | 28786467 | 103.64 MB/s | 9257240 | 86279 | 2.1× |
| Stdlib | 59238509 | 50.36 MB/s | 9258090 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1528977 | 473.26 MB/s | 696698 | 3610 | 9.1× |
| SonicFastest | 2288045 | 316.25 MB/s | 2374591 | 3683 | 6.1× |
| Sonic | 2291631 | 315.76 MB/s | 2374993 | 3683 | 6.1× |
| Easyjson | 5517777 | 131.14 MB/s | 2847905 | 3698 | 2.5× |
| Goccy | 5719941 | 126.50 MB/s | 2919460 | 80279 | 2.4× |
| JSONV2 | 6368988 | 113.61 MB/s | 2704770 | 7319 | 2.2× |
| Stdlib | 13907048 | 52.03 MB/s | 2704554 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2131261 | 740.10 MB/s | 696699 | 3610 | 9.6× |
| Sonic | 3051475 | 516.91 MB/s | 3230644 | 3683 | 6.7× |
| SonicFastest | 3154593 | 500.02 MB/s | 3232066 | 3683 | 6.5× |
| Goccy | 7017287 | 224.78 MB/s | 3598054 | 80267 | 2.9× |
| Easyjson | 7029808 | 224.38 MB/s | 2847906 | 3698 | 2.9× |
| JSONV2 | 7461086 | 211.41 MB/s | 2704686 | 7318 | 2.8× |
| Stdlib | 20545892 | 76.77 MB/s | 2704552 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 249300 | 602.18 MB/s | 81920 | 1 | 8.8× |
| SonicFastest | 404535 | 371.10 MB/s | 410046 | 16 | 5.4× |
| Sonic | 408782 | 367.25 MB/s | 409617 | 16 | 5.4× |
| Goccy | 1000853 | 150.00 MB/s | 329609 | 10005 | 2.2× |
| JSONV2 | 1146215 | 130.97 MB/s | 357724 | 20 | 1.9× |
| Stdlib | 2198105 | 68.30 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 39069 | 719.68 MB/s | 29104 | 107 | 8.8× |
| SonicFastest | 57731 | 487.03 MB/s | 59542 | 83 | 6.0× |
| Sonic | 57799 | 486.46 MB/s | 59529 | 83 | 5.9× |
| Goccy | 82487 | 340.87 MB/s | 59283 | 188 | 4.2× |
| Easyjson | 83548 | 336.54 MB/s | 32304 | 138 | 4.1× |
| JSONV2 | 139875 | 201.02 MB/s | 36896 | 242 | 2.5× |
| Stdlib | 343556 | 81.84 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2079 | 1119.73 MB/s | 32 | 1 | 12.5× |
| SonicFastest | 4761 | 489.00 MB/s | 3711 | 4 | 5.5× |
| Sonic | 4778 | 487.19 MB/s | 3712 | 4 | 5.4× |
| Goccy | 5130 | 453.84 MB/s | 3649 | 4 | 5.1× |
| Easyjson | 5172 | 450.10 MB/s | 192 | 2 | 5.0× |
| JSONV2 | 8210 | 283.55 MB/s | 1000 | 6 | 3.2× |
| Stdlib | 25995 | 89.55 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 225 | 838.80 MB/s | 0 | 0 | 12.3× |
| Goccy | 452 | 418.15 MB/s | 304 | 2 | 6.2× |
| Easyjson | 552 | 342.06 MB/s | 0 | 0 | 5.0× |
| SonicFastest | 658 | 287.16 MB/s | 341 | 3 | 4.2× |
| Sonic | 662 | 285.60 MB/s | 341 | 3 | 4.2× |
| JSONV2 | 1046 | 180.73 MB/s | 112 | 1 | 2.7× |
| Stdlib | 2780 | 67.99 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1442 | 1519.01 MB/s | 0 | 0 | 13.1× |
| Easyjson | 3551 | 616.97 MB/s | 24 | 1 | 5.3× |
| Goccy | 3981 | 550.32 MB/s | 2864 | 4 | 4.8× |
| SonicFastest | 6689 | 327.55 MB/s | 3602 | 38 | 2.8× |
| Sonic | 6814 | 321.54 MB/s | 3600 | 38 | 2.8× |
| JSONV2 | 8574 | 255.55 MB/s | 640 | 6 | 2.2× |
| Stdlib | 18911 | 115.86 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1039396 | 491.13 MB/s | 703780 | 1012 | 6.7× |
| Goccy | 1303496 | 391.62 MB/s | 1144057 | 5007 | 5.4× |
| SonicFastest | 1398231 | 365.09 MB/s | 1316497 | 2014 | 5.0× |
| Sonic | 1425488 | 358.11 MB/s | 1317158 | 2014 | 4.9× |
| Easyjson | 1697411 | 300.74 MB/s | 863776 | 3012 | 4.1× |
| JSONV2 | 3548080 | 143.87 MB/s | 1076015 | 12646 | 2.0× |
| Stdlib | 7009400 | 72.83 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 644 | 30746.23 MB/s | 0 | 0 | 249.5× |
| SonicFastest | 6507 | 3041.02 MB/s | 21106 | 3 | 24.7× |
| Goccy | 29631 | 667.85 MB/s | 20492 | 2 | 5.4× |
| Sonic | 32729 | 604.62 MB/s | 20643 | 3 | 4.9× |
| JSONV2 | 33262 | 594.93 MB/s | 8 | 1 | 4.8× |
| Easyjson | 98216 | 201.48 MB/s | 0 | 0 | 1.6× |
| Stdlib | 160557 | 123.25 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2448 | 7403.14 MB/s | 864 | 4 | 49.2× |
| Easyjson | 4800 | 3775.96 MB/s | 432 | 2 | 25.1× |
| SonicFastest | 8591 | 2109.74 MB/s | 20418 | 5 | 14.0× |
| Sonic | 9117 | 1988.01 MB/s | 20402 | 5 | 13.2× |
| Goccy | 25372 | 714.32 MB/s | 19460 | 2 | 4.8× |
| JSONV2 | 50494 | 358.94 MB/s | 16500 | 50 | 2.4× |
| Stdlib | 120543 | 150.35 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2918099 | 688.29 MB/s | 2963985 | 7453 | 7.9× |
| SonicFastest | 3725961 | 539.05 MB/s | 5204784 | 7086 | 6.2× |
| Sonic | 3853450 | 521.22 MB/s | 5206745 | 7086 | 6.0× |
| Goccy | 4791131 | 419.21 MB/s | 5415455 | 15847 | 4.8× |
| Easyjson | 5720527 | 351.10 MB/s | 2981701 | 7441 | 4.0× |
| JSONV2 | 8153355 | 246.34 MB/s | 3173786 | 14563 | 2.8× |
| Stdlib | 23099154 | 86.95 MB/s | 3589320 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 873 | 628.64 MB/s | 480 | 1 | 7.8× |
| Easyjson | 2374 | 231.23 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2409 | 227.87 MB/s | 2261 | 8 | 2.8× |
| Sonic | 2425 | 226.40 MB/s | 2262 | 8 | 2.8× |
| JSONV2 | 3256 | 168.63 MB/s | 1664 | 7 | 2.1× |
| Goccy | 3437 | 159.74 MB/s | 2128 | 43 | 2.0× |
| Stdlib | 6770 | 81.10 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 626374 | 1008.21 MB/s | 461114 | 1230 | 10.1× |
| Sonic | 1108154 | 569.88 MB/s | 1071572 | 814 | 5.7× |
| SonicFastest | 1114460 | 566.65 MB/s | 1071722 | 814 | 5.7× |
| Goccy | 1411441 | 447.42 MB/s | 993477 | 1202 | 4.5× |
| Easyjson | 1449995 | 435.53 MB/s | 422504 | 936 | 4.4× |
| JSONV2 | 2386197 | 264.65 MB/s | 571625 | 3144 | 2.7× |
| Stdlib | 6345564 | 99.52 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1250138 | 449.88 MB/s | 945690 | 4207 | 4.8× |
| SonicFastest | 1497261 | 375.62 MB/s | 1418047 | 2239 | 4.0× |
| Sonic | 1528779 | 367.88 MB/s | 1418740 | 2239 | 3.9× |
| Goccy | 1718442 | 327.28 MB/s | 1110098 | 3403 | 3.5× |
| Easyjson | 2472800 | 227.44 MB/s | 833617 | 3391 | 2.4× |
| JSONV2 | 3390473 | 165.88 MB/s | 979881 | 5183 | 1.8× |
| Stdlib | 6032038 | 93.24 MB/s | 1067140 | 8759 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1370066 | 389.16 MB/s | 935535 | 16692 | 5.2× |
| SonicFastest | 1604772 | 332.25 MB/s | 1394379 | 9201 | 4.4× |
| Sonic | 1625416 | 328.03 MB/s | 1393764 | 9201 | 4.3× |
| Easyjson | 1918032 | 277.98 MB/s | 949526 | 15985 | 3.7× |
| Goccy | 2213902 | 240.83 MB/s | 1668735 | 18049 | 3.2× |
| JSONV2 | 3521845 | 151.39 MB/s | 1192568 | 22667 | 2.0× |
| Stdlib | 7069187 | 75.42 MB/s | 1245607 | 26510 | 1.0× |
