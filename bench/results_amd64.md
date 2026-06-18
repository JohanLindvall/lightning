# JSON Deserialization Benchmarks

- generated 2026-06-18T18:45:19Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 424761 | 636.60 MB/s | 348026 | 1627 | 5.8× |
| Sonic | 541221 | 499.62 MB/s | 642327 | 1147 | 4.6× |
| SonicFastest | 547789 | 493.63 MB/s | 643305 | 1147 | 4.5× |
| Goccy | 1126176 | 240.11 MB/s | 543537 | 8122 | 2.2× |
| Easyjson | 1136189 | 237.99 MB/s | 330272 | 749 | 2.2× |
| JSONV2 | 1449087 | 186.60 MB/s | 348450 | 1628 | 1.7× |
| Stdlib | 2462777 | 109.80 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1078042 | 1602.17 MB/s | 964649 | 6177 | 10.0× |
| Sonic | 1570001 | 1100.13 MB/s | 2741732 | 5548 | 6.9× |
| SonicFastest | 1592125 | 1084.84 MB/s | 2741322 | 5548 | 6.8× |
| Goccy | 1626473 | 1061.93 MB/s | 2599476 | 14610 | 6.7× |
| Easyjson | 2585967 | 667.91 MB/s | 972033 | 5389 | 4.2× |
| JSONV2 | 2875664 | 600.63 MB/s | 1012191 | 7594 | 3.8× |
| Stdlib | 10826540 | 159.53 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 818 | 2216.25 MB/s | 144 | 10 | 11.4× |
| Easyjson | 1866 | 971.31 MB/s | 144 | 10 | 5.0× |
| Goccy | 2342 | 773.71 MB/s | 2602 | 5 | 4.0× |
| SonicFastest | 4404 | 411.48 MB/s | 3382 | 40 | 2.1× |
| Sonic | 4600 | 393.88 MB/s | 3387 | 40 | 2.0× |
| JSONV2 | 4840 | 374.42 MB/s | 632 | 7 | 1.9× |
| Stdlib | 9314 | 194.54 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 611 | 2964.42 MB/s | 0 | 0 | 15.3× |
| Easyjson | 1671 | 1084.39 MB/s | 24 | 1 | 5.6× |
| Goccy | 2375 | 762.88 MB/s | 2610 | 4 | 3.9× |
| SonicFastest | 4442 | 407.89 MB/s | 3364 | 38 | 2.1× |
| Sonic | 4534 | 399.64 MB/s | 3365 | 38 | 2.1× |
| JSONV2 | 4756 | 381.01 MB/s | 640 | 6 | 2.0× |
| Stdlib | 9334 | 194.14 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 663 | 2731.53 MB/s | 0 | 0 | 13.7× |
| Easyjson | 1641 | 1104.31 MB/s | 24 | 1 | 5.5× |
| Goccy | 2429 | 746.00 MB/s | 2610 | 4 | 3.7× |
| SonicFastest | 4540 | 399.16 MB/s | 3364 | 38 | 2.0× |
| Sonic | 4565 | 396.95 MB/s | 3362 | 38 | 2.0× |
| JSONV2 | 4872 | 371.94 MB/s | 640 | 6 | 1.9× |
| Stdlib | 9104 | 199.03 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 312 | 736.77 MB/s | 160 | 1 | 8.7× |
| SonicFastest | 660 | 348.41 MB/s | 805 | 8 | 4.1× |
| Sonic | 660 | 348.21 MB/s | 806 | 8 | 4.1× |
| Easyjson | 1015 | 226.61 MB/s | 448 | 3 | 2.7× |
| Goccy | 1163 | 197.82 MB/s | 585 | 23 | 2.3× |
| JSONV2 | 1550 | 148.34 MB/s | 528 | 7 | 1.8× |
| Stdlib | 2726 | 84.36 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 543 | 910.11 MB/s | 160 | 1 | 7.2× |
| Sonic | 901 | 548.08 MB/s | 1081 | 8 | 4.3× |
| SonicFastest | 919 | 537.28 MB/s | 1084 | 8 | 4.2× |
| Easyjson | 1500 | 329.26 MB/s | 448 | 3 | 2.6× |
| Goccy | 1816 | 271.98 MB/s | 857 | 23 | 2.1× |
| JSONV2 | 1982 | 249.26 MB/s | 528 | 7 | 2.0× |
| Stdlib | 3893 | 126.89 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2498110 | 776.78 MB/s | 2846866 | 2238 | 6.6× |
| Goccy | 3199624 | 606.47 MB/s | 4064217 | 13509 | 5.2× |
| Sonic | 4742561 | 409.16 MB/s | 4879153 | 1736 | 3.5× |
| SonicFastest | 4777198 | 406.19 MB/s | 4879437 | 1736 | 3.5× |
| Easyjson | 5498250 | 352.93 MB/s | 3871267 | 15043 | 3.0× |
| JSONV2 | 7590388 | 255.65 MB/s | 3238624 | 13947 | 2.2× |
| Stdlib | 16551259 | 117.24 MB/s | 3551320 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 9409703 | 860.82 MB/s | 15059842 | 51649 | 6.9× |
| Sonic | 12617322 | 641.98 MB/s | 19896170 | 41641 | 5.1× |
| SonicFastest | 12719354 | 636.83 MB/s | 19903220 | 41641 | 5.1× |
| Goccy | 17561790 | 461.23 MB/s | 17660820 | 107150 | 3.7× |
| Easyjson | 20959489 | 386.46 MB/s | 15059620 | 41643 | 3.1× |
| JSONV2 | 29734091 | 272.42 MB/s | 15235674 | 78973 | 2.2× |
| Stdlib | 64528420 | 125.53 MB/s | 15665076 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 883 | 2480.15 MB/s | 0 | 0 | 12.7× |
| Easyjson | 2054 | 1066.63 MB/s | 24 | 1 | 5.5× |
| Goccy | 2659 | 823.84 MB/s | 2866 | 4 | 4.2× |
| SonicFastest | 4758 | 460.51 MB/s | 3623 | 38 | 2.4× |
| Sonic | 4972 | 440.70 MB/s | 3622 | 38 | 2.3× |
| JSONV2 | 5030 | 435.60 MB/s | 640 | 6 | 2.2× |
| Stdlib | 11249 | 194.77 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 387 | 51081.09 MB/s | 0 | 0 | 226.2× |
| SonicFastest | 3791 | 5219.78 MB/s | 21555 | 3 | 23.1× |
| Goccy | 18494 | 1070.02 MB/s | 20507 | 2 | 4.7× |
| Sonic | 20960 | 944.11 MB/s | 20748 | 3 | 4.2× |
| JSONV2 | 21985 | 900.11 MB/s | 8 | 1 | 4.0× |
| Easyjson | 67027 | 295.24 MB/s | 0 | 0 | 1.3× |
| Stdlib | 87627 | 225.83 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1748 | 10366.99 MB/s | 864 | 4 | 45.0× |
| Easyjson | 2885 | 6282.56 MB/s | 432 | 2 | 27.2× |
| SonicFastest | 6132 | 2955.65 MB/s | 20406 | 5 | 12.8× |
| Sonic | 6245 | 2901.98 MB/s | 20421 | 5 | 12.6× |
| Goccy | 15784 | 1148.23 MB/s | 19475 | 2 | 5.0× |
| JSONV2 | 33025 | 548.80 MB/s | 16514 | 50 | 2.4× |
| Stdlib | 78605 | 230.57 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1915863 | 1048.35 MB/s | 2963984 | 7453 | 7.3× |
| Goccy | 3514319 | 571.52 MB/s | 5465558 | 15848 | 4.0× |
| Sonic | 3630581 | 553.22 MB/s | 5158693 | 7085 | 3.8× |
| SonicFastest | 3653292 | 549.78 MB/s | 5154654 | 7085 | 3.8× |
| Easyjson | 3665745 | 547.91 MB/s | 2982889 | 7441 | 3.8× |
| JSONV2 | 5147166 | 390.21 MB/s | 3174775 | 14563 | 2.7× |
| Stdlib | 13972348 | 143.75 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 736 | 745.94 MB/s | 480 | 1 | 5.7× |
| Easyjson | 1433 | 383.14 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 1572 | 349.34 MB/s | 2275 | 8 | 2.6× |
| Sonic | 1627 | 337.53 MB/s | 2272 | 8 | 2.6× |
| JSONV2 | 2166 | 253.47 MB/s | 1665 | 7 | 1.9× |
| Goccy | 2242 | 244.87 MB/s | 2132 | 43 | 1.9× |
| Stdlib | 4162 | 131.90 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 454311 | 1390.05 MB/s | 461119 | 1230 | 8.7× |
| SonicFastest | 836760 | 754.71 MB/s | 1082787 | 814 | 4.7× |
| Sonic | 838010 | 753.59 MB/s | 1088163 | 814 | 4.7× |
| Easyjson | 917937 | 687.97 MB/s | 422504 | 936 | 4.3× |
| Goccy | 938298 | 673.04 MB/s | 998910 | 1203 | 4.2× |
| JSONV2 | 1510007 | 418.22 MB/s | 572021 | 3144 | 2.6× |
| Stdlib | 3955634 | 159.65 MB/s | 654667 | 6472 | 1.0× |
