# JSON Deserialization Benchmarks

- generated 2026-06-17T20:22:51Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 475943 | 568.14 MB/s | 151224 | 495 | 5.3× |
| SonicFastest | 551858 | 489.99 MB/s | 645077 | 1147 | 4.6× |
| Sonic | 557677 | 484.87 MB/s | 645138 | 1147 | 4.5× |
| Easyjson | 1238993 | 218.24 MB/s | 330320 | 753 | 2.0× |
| Goccy | 1260982 | 214.44 MB/s | 544557 | 8122 | 2.0× |
| JSONV2 | 1636054 | 165.28 MB/s | 348577 | 1628 | 1.5× |
| Stdlib | 2527600 | 106.98 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1434659 | 1203.91 MB/s | 979492 | 6803 | 6.8× |
| Sonic | 1728096 | 999.48 MB/s | 2729457 | 5548 | 5.7× |
| SonicFastest | 1731267 | 997.65 MB/s | 2730256 | 5548 | 5.7× |
| Goccy | 1977248 | 873.54 MB/s | 2603466 | 14609 | 4.9× |
| JSONV2 | 3061055 | 564.25 MB/s | 1012354 | 7594 | 3.2× |
| Easyjson | 3255420 | 530.56 MB/s | 986873 | 6015 | 3.0× |
| Stdlib | 9783270 | 176.55 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 885 | 2048.21 MB/s | 144 | 10 | 10.6× |
| Easyjson | 1870 | 968.90 MB/s | 144 | 10 | 5.0× |
| Goccy | 2474 | 732.44 MB/s | 2603 | 5 | 3.8× |
| SonicFastest | 4455 | 406.77 MB/s | 3381 | 40 | 2.1× |
| Sonic | 4664 | 388.53 MB/s | 3380 | 40 | 2.0× |
| JSONV2 | 4940 | 366.78 MB/s | 633 | 7 | 1.9× |
| Stdlib | 9415 | 192.47 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 610 | 2968.19 MB/s | 0 | 0 | 15.7× |
| Easyjson | 1679 | 1079.31 MB/s | 24 | 1 | 5.7× |
| Goccy | 2748 | 659.28 MB/s | 2611 | 4 | 3.5× |
| SonicFastest | 4569 | 396.62 MB/s | 3360 | 38 | 2.1× |
| Sonic | 4645 | 390.14 MB/s | 3360 | 38 | 2.1× |
| JSONV2 | 4884 | 371.04 MB/s | 640 | 6 | 2.0× |
| Stdlib | 9560 | 189.54 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 673 | 2690.64 MB/s | 0 | 0 | 13.8× |
| Easyjson | 1664 | 1088.72 MB/s | 24 | 1 | 5.6× |
| Goccy | 2663 | 680.50 MB/s | 2611 | 4 | 3.5× |
| SonicFastest | 4403 | 411.58 MB/s | 3365 | 38 | 2.1× |
| Sonic | 4536 | 399.45 MB/s | 3361 | 38 | 2.0× |
| JSONV2 | 5056 | 358.42 MB/s | 641 | 6 | 1.8× |
| Stdlib | 9295 | 194.94 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 348 | 661.39 MB/s | 160 | 1 | 8.9× |
| SonicFastest | 667 | 345.00 MB/s | 813 | 8 | 4.6× |
| Sonic | 689 | 333.71 MB/s | 812 | 8 | 4.5× |
| Easyjson | 1124 | 204.70 MB/s | 448 | 3 | 2.8× |
| Goccy | 1290 | 178.35 MB/s | 585 | 23 | 2.4× |
| JSONV2 | 1659 | 138.61 MB/s | 528 | 7 | 1.9× |
| Stdlib | 3094 | 74.34 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 526 | 939.58 MB/s | 160 | 1 | 8.0× |
| Sonic | 931 | 530.58 MB/s | 1091 | 8 | 4.5× |
| SonicFastest | 949 | 520.72 MB/s | 1092 | 8 | 4.5× |
| Easyjson | 1701 | 290.50 MB/s | 448 | 3 | 2.5× |
| Goccy | 1969 | 250.93 MB/s | 858 | 23 | 2.1× |
| JSONV2 | 2170 | 227.61 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4228 | 116.83 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2992827 | 648.37 MB/s | 3258552 | 15044 | 5.6× |
| Goccy | 3218926 | 602.83 MB/s | 4066639 | 13509 | 5.2× |
| SonicFastest | 3512940 | 552.38 MB/s | 4886358 | 1736 | 4.8× |
| Sonic | 3597253 | 539.43 MB/s | 4886729 | 1736 | 4.7× |
| Easyjson | 6414594 | 302.51 MB/s | 4282954 | 27849 | 2.6× |
| JSONV2 | 7897198 | 245.72 MB/s | 3239106 | 13947 | 2.1× |
| Stdlib | 16745538 | 115.88 MB/s | 3551320 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| SonicFastest | 10555929 | 767.34 MB/s | 19863403 | 41641 | 5.9× |
| Sonic | 10602672 | 763.96 MB/s | 19864615 | 41641 | 5.9× |
| Lightning | 11915201 | 679.81 MB/s | 12528038 | 40027 | 5.2× |
| Goccy | 18589564 | 435.73 MB/s | 17325247 | 107149 | 3.3× |
| Easyjson | 24560275 | 329.80 MB/s | 15219649 | 61644 | 2.5× |
| JSONV2 | 32035238 | 252.85 MB/s | 15236449 | 78973 | 1.9× |
| Stdlib | 62093414 | 130.45 MB/s | 15665066 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 859 | 2549.44 MB/s | 0 | 0 | 12.6× |
| Easyjson | 2093 | 1046.98 MB/s | 24 | 1 | 5.2× |
| Goccy | 2998 | 730.76 MB/s | 2867 | 4 | 3.6× |
| SonicFastest | 4795 | 456.97 MB/s | 3631 | 38 | 2.3× |
| Sonic | 4969 | 440.98 MB/s | 3623 | 38 | 2.2× |
| JSONV2 | 5176 | 423.27 MB/s | 641 | 6 | 2.1× |
| Stdlib | 10818 | 202.52 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 504 | 39306.68 MB/s | 0 | 0 | 156.3× |
| SonicFastest | 3840 | 5153.07 MB/s | 21444 | 3 | 20.5× |
| Goccy | 20477 | 966.40 MB/s | 20515 | 2 | 3.8× |
| JSONV2 | 22140 | 893.80 MB/s | 8 | 1 | 3.6× |
| Sonic | 27626 | 716.31 MB/s | 20811 | 3 | 2.8× |
| Easyjson | 47549 | 416.18 MB/s | 0 | 0 | 1.7× |
| Stdlib | 78709 | 251.42 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 5502 | 3294.12 MB/s | 18080 | 62 | 13.1× |
| Sonic | 5820 | 3114.04 MB/s | 20723 | 5 | 12.4× |
| SonicFastest | 5980 | 3030.54 MB/s | 20750 | 5 | 12.0× |
| Easyjson | 6977 | 2597.60 MB/s | 17648 | 60 | 10.3× |
| Goccy | 16364 | 1107.56 MB/s | 19483 | 2 | 4.4× |
| JSONV2 | 34781 | 521.08 MB/s | 16520 | 50 | 2.1× |
| Stdlib | 72017 | 251.66 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2667708 | 752.89 MB/s | 3584575 | 29240 | 4.9× |
| Sonic | 3092788 | 649.41 MB/s | 5156957 | 7085 | 4.2× |
| SonicFastest | 3103634 | 647.14 MB/s | 5155159 | 7085 | 4.2× |
| Goccy | 3480773 | 577.03 MB/s | 5484209 | 15839 | 3.8× |
| Easyjson | 4649445 | 431.99 MB/s | 3604022 | 29228 | 2.8× |
| JSONV2 | 5025733 | 399.64 MB/s | 3175176 | 14563 | 2.6× |
| Stdlib | 13071570 | 153.65 MB/s | 3589321 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 809 | 678.48 MB/s | 480 | 1 | 5.3× |
| Easyjson | 1733 | 316.81 MB/s | 1616 | 5 | 2.5× |
| SonicFastest | 1751 | 313.55 MB/s | 2283 | 8 | 2.4× |
| Sonic | 1791 | 306.54 MB/s | 2282 | 8 | 2.4× |
| Goccy | 2420 | 226.84 MB/s | 2133 | 43 | 1.8× |
| JSONV2 | 2496 | 219.99 MB/s | 1666 | 7 | 1.7× |
| Stdlib | 4272 | 128.50 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 597118 | 1057.60 MB/s | 630174 | 5499 | 6.4× |
| Sonic | 747261 | 845.11 MB/s | 1085439 | 814 | 5.1× |
| SonicFastest | 787457 | 801.97 MB/s | 1085198 | 814 | 4.9× |
| Goccy | 988674 | 638.75 MB/s | 996954 | 1202 | 3.9× |
| Easyjson | 1054795 | 598.71 MB/s | 591597 | 5205 | 3.6× |
| JSONV2 | 1528538 | 413.15 MB/s | 572116 | 3144 | 2.5× |
| Stdlib | 3844542 | 164.26 MB/s | 654666 | 6472 | 1.0× |
