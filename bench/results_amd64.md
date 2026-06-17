# JSON Deserialization Benchmarks

- generated 2026-06-17T07:08:11Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 696539 | 388.21 MB/s | 333834 | 7649 | 4.4× |
| Sonic | 726225 | 372.34 MB/s | 918593 | 8294 | 4.2× |
| Easyjson | 1708571 | 158.26 MB/s | 962174 | 8297 | 1.8× |
| JSONV2 | 2133224 | 126.76 MB/s | 688471 | 15929 | 1.4× |
| Stdlib | 3051185 | 88.62 MB/s | 688019 | 15942 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1388958 | 1243.53 MB/s | 979504 | 6803 | 6.4× |
| Sonic | 1829200 | 944.24 MB/s | 2730854 | 5548 | 4.8× |
| Easyjson | 3015652 | 572.75 MB/s | 986874 | 6015 | 2.9× |
| JSONV2 | 3043029 | 567.59 MB/s | 1012479 | 7594 | 2.9× |
| Stdlib | 8822553 | 195.77 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 928 | 1952.54 MB/s | 144 | 10 | 10.0× |
| Easyjson | 1791 | 1011.53 MB/s | 144 | 10 | 5.2× |
| Sonic | 4741 | 382.16 MB/s | 3377 | 40 | 2.0× |
| JSONV2 | 5271 | 343.77 MB/s | 633 | 7 | 1.8× |
| Stdlib | 9307 | 194.70 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 643 | 2817.77 MB/s | 0 | 0 | 15.3× |
| Easyjson | 1672 | 1083.65 MB/s | 24 | 1 | 5.9× |
| Sonic | 4822 | 375.74 MB/s | 3356 | 38 | 2.0× |
| JSONV2 | 5339 | 339.39 MB/s | 640 | 6 | 1.8× |
| Stdlib | 9871 | 183.57 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 674 | 2689.52 MB/s | 0 | 0 | 14.3× |
| Easyjson | 1665 | 1088.32 MB/s | 24 | 1 | 5.8× |
| Sonic | 4469 | 405.48 MB/s | 3372 | 38 | 2.1× |
| JSONV2 | 5256 | 344.77 MB/s | 640 | 6 | 1.8× |
| Stdlib | 9601 | 188.73 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 346 | 665.34 MB/s | 160 | 1 | 8.6× |
| Sonic | 673 | 341.69 MB/s | 812 | 8 | 4.4× |
| Easyjson | 1063 | 216.43 MB/s | 448 | 3 | 2.8× |
| JSONV2 | 1583 | 145.27 MB/s | 528 | 7 | 1.9× |
| Stdlib | 2972 | 77.40 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 532 | 928.39 MB/s | 160 | 1 | 7.5× |
| Sonic | 971 | 508.58 MB/s | 1090 | 8 | 4.1× |
| Easyjson | 1637 | 301.85 MB/s | 448 | 3 | 2.4× |
| JSONV2 | 2074 | 238.15 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4009 | 123.23 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3376985 | 574.62 MB/s | 4282957 | 27849 | 5.0× |
| Sonic | 3642351 | 532.75 MB/s | 4887237 | 1736 | 4.6× |
| Easyjson | 6562041 | 295.71 MB/s | 4282954 | 27849 | 2.6× |
| JSONV2 | 7951632 | 244.03 MB/s | 3239299 | 13947 | 2.1× |
| Stdlib | 16904473 | 114.79 MB/s | 3551319 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 12971673 | 624.44 MB/s | 23096153 | 175893 | 5.2× |
| Lightning | 13253149 | 611.18 MB/s | 14139126 | 107156 | 5.1× |
| Easyjson | 26803576 | 302.20 MB/s | 19531265 | 128767 | 2.5× |
| JSONV2 | 37489995 | 216.06 MB/s | 19028153 | 280355 | 1.8× |
| Stdlib | 67408366 | 120.16 MB/s | 19437634 | 352028 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 864 | 2535.61 MB/s | 0 | 0 | 12.3× |
| Easyjson | 2059 | 1064.07 MB/s | 24 | 1 | 5.2× |
| Sonic | 4885 | 448.49 MB/s | 3624 | 38 | 2.2× |
| JSONV2 | 5296 | 413.69 MB/s | 640 | 6 | 2.0× |
| Stdlib | 10638 | 205.96 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 510 | 38837.72 MB/s | 0 | 0 | 141.3× |
| JSONV2 | 21736 | 910.42 MB/s | 8 | 1 | 3.3× |
| Sonic | 25465 | 777.09 MB/s | 20656 | 3 | 2.8× |
| Easyjson | 46886 | 422.07 MB/s | 0 | 0 | 1.5× |
| Stdlib | 71974 | 274.95 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 5494 | 3299.07 MB/s | 18080 | 62 | 13.3× |
| Sonic | 5551 | 3265.25 MB/s | 20771 | 5 | 13.2× |
| Easyjson | 6991 | 2592.60 MB/s | 17648 | 60 | 10.5× |
| JSONV2 | 32153 | 563.68 MB/s | 16517 | 50 | 2.3× |
| Stdlib | 73085 | 247.99 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2553953 | 786.43 MB/s | 3587264 | 29254 | 5.0× |
| Sonic | 2941427 | 682.83 MB/s | 5159304 | 7085 | 4.3× |
| Easyjson | 4542118 | 442.19 MB/s | 3604001 | 29228 | 2.8× |
| JSONV2 | 5140521 | 390.72 MB/s | 3175378 | 14563 | 2.5× |
| Stdlib | 12742144 | 157.63 MB/s | 3589320 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 759 | 723.58 MB/s | 480 | 1 | 5.7× |
| Easyjson | 1768 | 310.60 MB/s | 1616 | 5 | 2.4× |
| Sonic | 1799 | 305.09 MB/s | 2283 | 8 | 2.4× |
| JSONV2 | 2410 | 227.79 MB/s | 1665 | 7 | 1.8× |
| Stdlib | 4315 | 127.24 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 621405 | 1016.27 MB/s | 634494 | 5544 | 6.4× |
| Sonic | 844393 | 747.89 MB/s | 1086112 | 814 | 4.7× |
| Easyjson | 1080168 | 584.64 MB/s | 591597 | 5205 | 3.7× |
| JSONV2 | 1618618 | 390.16 MB/s | 572271 | 3144 | 2.5× |
| Stdlib | 3984411 | 158.50 MB/s | 654665 | 6472 | 1.0× |
