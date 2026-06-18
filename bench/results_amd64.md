# JSON Deserialization Benchmarks

- generated 2026-06-18T06:38:09Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 483001 | 559.84 MB/s | 348074 | 1631 | 5.5× |
| SonicFastest | 567951 | 476.10 MB/s | 644426 | 1147 | 4.6× |
| Sonic | 575988 | 469.46 MB/s | 644420 | 1147 | 4.6× |
| Easyjson | 1287761 | 209.98 MB/s | 330321 | 753 | 2.0× |
| Goccy | 1369174 | 197.49 MB/s | 544027 | 8121 | 1.9× |
| JSONV2 | 1815060 | 148.98 MB/s | 348578 | 1628 | 1.5× |
| Stdlib | 2638940 | 102.47 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1189122 | 1452.50 MB/s | 979492 | 6803 | 7.8× |
| Sonic | 1748674 | 987.72 MB/s | 2730929 | 5548 | 5.3× |
| SonicFastest | 1774617 | 973.28 MB/s | 2734033 | 5548 | 5.2× |
| Goccy | 1939611 | 890.49 MB/s | 2603617 | 14610 | 4.8× |
| JSONV2 | 3019365 | 572.04 MB/s | 1012351 | 7594 | 3.1× |
| Easyjson | 3144850 | 549.22 MB/s | 986873 | 6015 | 2.9× |
| Stdlib | 9273153 | 186.26 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 899 | 2015.75 MB/s | 144 | 10 | 10.3× |
| Easyjson | 1966 | 921.67 MB/s | 144 | 10 | 4.7× |
| Goccy | 2600 | 697.04 MB/s | 2603 | 5 | 3.6× |
| SonicFastest | 4428 | 409.24 MB/s | 3386 | 40 | 2.1× |
| Sonic | 4614 | 392.68 MB/s | 3383 | 40 | 2.0× |
| JSONV2 | 5048 | 358.94 MB/s | 632 | 7 | 1.8× |
| Stdlib | 9293 | 195.00 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 603 | 3003.84 MB/s | 0 | 0 | 15.8× |
| Easyjson | 1666 | 1087.96 MB/s | 24 | 1 | 5.7× |
| Goccy | 2661 | 681.06 MB/s | 2611 | 4 | 3.6× |
| SonicFastest | 4481 | 404.36 MB/s | 3369 | 38 | 2.1× |
| Sonic | 4654 | 389.31 MB/s | 3366 | 38 | 2.0× |
| JSONV2 | 4929 | 367.60 MB/s | 641 | 6 | 1.9× |
| Stdlib | 9501 | 190.71 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 678 | 2673.88 MB/s | 0 | 0 | 13.8× |
| Easyjson | 1667 | 1087.00 MB/s | 24 | 1 | 5.6× |
| Goccy | 2536 | 714.51 MB/s | 2611 | 4 | 3.7× |
| SonicFastest | 4494 | 403.19 MB/s | 3367 | 38 | 2.1× |
| Sonic | 4613 | 392.78 MB/s | 3365 | 38 | 2.0× |
| JSONV2 | 4960 | 365.31 MB/s | 641 | 6 | 1.9× |
| Stdlib | 9320 | 194.42 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 334 | 687.89 MB/s | 160 | 1 | 8.9× |
| Sonic | 673 | 341.76 MB/s | 812 | 8 | 4.4× |
| SonicFastest | 679 | 338.72 MB/s | 813 | 8 | 4.4× |
| Easyjson | 1113 | 206.73 MB/s | 448 | 3 | 2.7× |
| Goccy | 1269 | 181.29 MB/s | 585 | 23 | 2.4× |
| JSONV2 | 1619 | 142.04 MB/s | 528 | 7 | 1.8× |
| Stdlib | 2985 | 77.05 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 521 | 948.76 MB/s | 160 | 1 | 7.9× |
| SonicFastest | 928 | 532.45 MB/s | 1092 | 8 | 4.4× |
| Sonic | 930 | 531.46 MB/s | 1091 | 8 | 4.4× |
| Easyjson | 1690 | 292.34 MB/s | 448 | 3 | 2.4× |
| Goccy | 1947 | 253.69 MB/s | 858 | 23 | 2.1× |
| JSONV2 | 2156 | 229.12 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4105 | 120.34 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2952334 | 657.27 MB/s | 3258550 | 15044 | 5.6× |
| Goccy | 3205460 | 605.36 MB/s | 4068059 | 13509 | 5.2× |
| Sonic | 3333141 | 582.18 MB/s | 4884851 | 1736 | 5.0× |
| SonicFastest | 3435382 | 564.85 MB/s | 4884097 | 1736 | 4.8× |
| Easyjson | 6319294 | 307.07 MB/s | 4282950 | 27849 | 2.6× |
| JSONV2 | 7864069 | 246.75 MB/s | 3239100 | 13947 | 2.1× |
| Stdlib | 16550482 | 117.25 MB/s | 3551319 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| SonicFastest | 10387909 | 779.76 MB/s | 19861278 | 41641 | 5.9× |
| Sonic | 10391913 | 779.46 MB/s | 19861415 | 41641 | 5.9× |
| Lightning | 10625188 | 762.34 MB/s | 15219864 | 71650 | 5.8× |
| Goccy | 18330745 | 441.88 MB/s | 17486117 | 107150 | 3.3× |
| Easyjson | 24950048 | 324.65 MB/s | 15219671 | 61644 | 2.5× |
| JSONV2 | 32199839 | 251.56 MB/s | 15236448 | 78973 | 1.9× |
| Stdlib | 61150326 | 132.46 MB/s | 15665070 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 918 | 2385.78 MB/s | 0 | 0 | 11.7× |
| Easyjson | 2051 | 1068.24 MB/s | 24 | 1 | 5.2× |
| Goccy | 2949 | 743.00 MB/s | 2867 | 4 | 3.6× |
| SonicFastest | 4771 | 459.28 MB/s | 3622 | 38 | 2.3× |
| Sonic | 5046 | 434.17 MB/s | 3622 | 38 | 2.1× |
| JSONV2 | 5059 | 433.05 MB/s | 640 | 6 | 2.1× |
| Stdlib | 10736 | 204.08 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 496 | 39875.89 MB/s | 0 | 0 | 151.3× |
| SonicFastest | 3898 | 5076.68 MB/s | 21449 | 3 | 19.3× |
| Goccy | 20751 | 953.64 MB/s | 20515 | 2 | 3.6× |
| JSONV2 | 21449 | 922.61 MB/s | 8 | 1 | 3.5× |
| Sonic | 26814 | 738.02 MB/s | 20709 | 3 | 2.8× |
| Easyjson | 46674 | 423.98 MB/s | 0 | 0 | 1.6× |
| Stdlib | 75111 | 263.46 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 5365 | 3377.88 MB/s | 20736 | 5 | 13.3× |
| SonicFastest | 5368 | 3376.16 MB/s | 20730 | 5 | 13.3× |
| Lightning | 5675 | 3193.90 MB/s | 18080 | 62 | 12.6× |
| Easyjson | 6964 | 2602.56 MB/s | 17648 | 60 | 10.3× |
| Goccy | 15796 | 1147.36 MB/s | 19483 | 2 | 4.5× |
| JSONV2 | 33951 | 533.84 MB/s | 16520 | 50 | 2.1× |
| Stdlib | 71556 | 253.28 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2507200 | 801.09 MB/s | 3584579 | 29240 | 5.2× |
| Sonic | 3076442 | 652.86 MB/s | 5154749 | 7085 | 4.2× |
| SonicFastest | 3148888 | 637.84 MB/s | 5154654 | 7085 | 4.1× |
| Goccy | 3477117 | 577.63 MB/s | 5484314 | 15838 | 3.7× |
| Easyjson | 4775537 | 420.58 MB/s | 3604049 | 29228 | 2.7× |
| JSONV2 | 5029822 | 399.32 MB/s | 3175179 | 14563 | 2.6× |
| Stdlib | 12997629 | 154.53 MB/s | 3589320 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 766 | 716.37 MB/s | 480 | 1 | 5.7× |
| SonicFastest | 1723 | 318.61 MB/s | 2283 | 8 | 2.5× |
| Easyjson | 1725 | 318.23 MB/s | 1616 | 5 | 2.5× |
| Sonic | 1804 | 304.25 MB/s | 2281 | 8 | 2.4× |
| Goccy | 2431 | 225.82 MB/s | 2133 | 43 | 1.8× |
| JSONV2 | 2448 | 224.28 MB/s | 1666 | 7 | 1.8× |
| Stdlib | 4339 | 126.53 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 571109 | 1105.77 MB/s | 630175 | 5499 | 7.1× |
| SonicFastest | 760959 | 829.89 MB/s | 1087202 | 814 | 5.3× |
| Sonic | 792262 | 797.10 MB/s | 1088082 | 814 | 5.1× |
| Goccy | 1002792 | 629.76 MB/s | 997147 | 1202 | 4.0× |
| Easyjson | 1075341 | 587.27 MB/s | 591597 | 5205 | 3.8× |
| JSONV2 | 1579884 | 399.72 MB/s | 572123 | 3144 | 2.6× |
| Stdlib | 4034478 | 156.53 MB/s | 654666 | 6472 | 1.0× |
