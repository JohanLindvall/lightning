# JSON Deserialization Benchmarks

- generated 2026-06-17T10:32:25Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 501401 | 539.29 MB/s | 151226 | 495 | 5.2× |
| Sonic | 552999 | 488.98 MB/s | 644917 | 1147 | 4.7× |
| Easyjson | 1293886 | 208.99 MB/s | 330320 | 753 | 2.0× |
| JSONV2 | 1595239 | 169.51 MB/s | 348531 | 1628 | 1.6× |
| Stdlib | 2588654 | 104.46 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1490013 | 1159.19 MB/s | 979505 | 6803 | 6.1× |
| Sonic | 1699569 | 1016.26 MB/s | 2735052 | 5548 | 5.3× |
| Easyjson | 3149969 | 548.32 MB/s | 986873 | 6015 | 2.9× |
| JSONV2 | 3173089 | 544.33 MB/s | 1012478 | 7594 | 2.9× |
| Stdlib | 9088838 | 190.04 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 958 | 1892.38 MB/s | 144 | 10 | 10.4× |
| Easyjson | 1880 | 963.97 MB/s | 144 | 10 | 5.3× |
| Sonic | 4690 | 386.38 MB/s | 3379 | 40 | 2.1× |
| JSONV2 | 5013 | 361.45 MB/s | 632 | 7 | 2.0× |
| Stdlib | 9987 | 181.43 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 648 | 2795.74 MB/s | 0 | 0 | 15.7× |
| Easyjson | 1713 | 1057.62 MB/s | 24 | 1 | 5.9× |
| Sonic | 4641 | 390.39 MB/s | 3363 | 38 | 2.2× |
| JSONV2 | 5096 | 355.56 MB/s | 640 | 6 | 2.0× |
| Stdlib | 10147 | 178.57 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 706 | 2566.53 MB/s | 0 | 0 | 14.4× |
| Easyjson | 1664 | 1089.21 MB/s | 24 | 1 | 6.1× |
| Sonic | 4616 | 392.53 MB/s | 3372 | 38 | 2.2× |
| JSONV2 | 5037 | 359.75 MB/s | 640 | 6 | 2.0× |
| Stdlib | 10165 | 178.26 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 347 | 662.75 MB/s | 160 | 1 | 8.5× |
| Sonic | 689 | 333.86 MB/s | 811 | 8 | 4.3× |
| Easyjson | 1085 | 212.00 MB/s | 448 | 3 | 2.7× |
| JSONV2 | 1644 | 139.92 MB/s | 528 | 7 | 1.8× |
| Stdlib | 2941 | 78.20 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 543 | 909.60 MB/s | 160 | 1 | 7.6× |
| Sonic | 900 | 549.06 MB/s | 1090 | 8 | 4.6× |
| Easyjson | 1704 | 289.87 MB/s | 448 | 3 | 2.4× |
| JSONV2 | 2223 | 222.18 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4149 | 119.07 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3392270 | 572.03 MB/s | 4282970 | 27849 | 5.0× |
| Sonic | 3937879 | 492.77 MB/s | 4882395 | 1736 | 4.3× |
| Easyjson | 6786159 | 285.95 MB/s | 4282959 | 27849 | 2.5× |
| JSONV2 | 8184718 | 237.08 MB/s | 3239390 | 13947 | 2.1× |
| Stdlib | 16862984 | 115.07 MB/s | 3551322 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 13746302 | 589.25 MB/s | 23094906 | 175893 | 5.1× |
| Lightning | 15085669 | 536.94 MB/s | 14139206 | 107156 | 4.6× |
| Easyjson | 28262543 | 286.60 MB/s | 19531267 | 128767 | 2.5× |
| JSONV2 | 38115121 | 212.52 MB/s | 19028154 | 280355 | 1.8× |
| Stdlib | 69588884 | 116.40 MB/s | 19437988 | 352028 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 837 | 2618.85 MB/s | 0 | 0 | 13.2× |
| Easyjson | 2013 | 1088.44 MB/s | 24 | 1 | 5.5× |
| JSONV2 | 5123 | 427.67 MB/s | 640 | 6 | 2.2× |
| Sonic | 5306 | 412.95 MB/s | 3626 | 38 | 2.1× |
| Stdlib | 11033 | 198.59 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 505 | 39151.83 MB/s | 0 | 0 | 147.7× |
| JSONV2 | 22379 | 884.28 MB/s | 8 | 1 | 3.3× |
| Sonic | 26910 | 735.38 MB/s | 20795 | 3 | 2.8× |
| Easyjson | 46873 | 422.19 MB/s | 0 | 0 | 1.6× |
| Stdlib | 74670 | 265.02 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 5392 | 3361.25 MB/s | 20757 | 5 | 13.9× |
| Lightning | 5610 | 3230.61 MB/s | 18080 | 62 | 13.4× |
| Easyjson | 6845 | 2647.92 MB/s | 17648 | 60 | 10.9× |
| JSONV2 | 32101 | 564.60 MB/s | 16517 | 50 | 2.3× |
| Stdlib | 74952 | 241.81 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2589949 | 775.50 MB/s | 3587261 | 29254 | 5.1× |
| Sonic | 3139319 | 639.79 MB/s | 5160612 | 7085 | 4.2× |
| Easyjson | 4409041 | 455.54 MB/s | 3604100 | 29229 | 3.0× |
| JSONV2 | 5087561 | 394.79 MB/s | 3175474 | 14563 | 2.6× |
| Stdlib | 13204385 | 152.11 MB/s | 3589322 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 736 | 745.80 MB/s | 480 | 1 | 5.6× |
| Easyjson | 1610 | 340.91 MB/s | 1616 | 5 | 2.5× |
| Sonic | 1730 | 317.27 MB/s | 2282 | 8 | 2.4× |
| JSONV2 | 2396 | 229.11 MB/s | 1665 | 7 | 1.7× |
| Stdlib | 4097 | 134.01 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 630582 | 1001.48 MB/s | 634497 | 5544 | 6.2× |
| Sonic | 830733 | 760.19 MB/s | 1088058 | 814 | 4.7× |
| Easyjson | 1135664 | 556.07 MB/s | 591597 | 5205 | 3.4× |
| JSONV2 | 1687831 | 374.16 MB/s | 572270 | 3144 | 2.3× |
| Stdlib | 3896679 | 162.06 MB/s | 654665 | 6472 | 1.0× |
