# JSON Deserialization Benchmarks

- generated 2026-06-18T19:05:39Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 420505 | 643.04 MB/s | 348026 | 1627 | 5.9× |
| SonicFastest | 541827 | 499.06 MB/s | 644087 | 1147 | 4.6× |
| Sonic | 554057 | 488.04 MB/s | 644111 | 1147 | 4.4× |
| Easyjson | 1169341 | 231.24 MB/s | 330272 | 749 | 2.1× |
| Goccy | 1244338 | 217.31 MB/s | 544622 | 8122 | 2.0× |
| JSONV2 | 1534218 | 176.25 MB/s | 348485 | 1628 | 1.6× |
| Stdlib | 2465431 | 109.68 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1064900 | 1621.94 MB/s | 964648 | 6177 | 10.2× |
| SonicFastest | 1597292 | 1081.33 MB/s | 2735554 | 5548 | 6.8× |
| Sonic | 1614948 | 1069.51 MB/s | 2741884 | 5548 | 6.7× |
| Goccy | 1988944 | 868.40 MB/s | 2600149 | 14610 | 5.5× |
| Easyjson | 2459756 | 702.18 MB/s | 972033 | 5389 | 4.4× |
| JSONV2 | 2697886 | 640.21 MB/s | 1012189 | 7594 | 4.0× |
| Stdlib | 10882066 | 158.72 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 827 | 2190.40 MB/s | 144 | 10 | 11.0× |
| Easyjson | 1854 | 977.60 MB/s | 144 | 10 | 4.9× |
| Goccy | 2472 | 732.98 MB/s | 2602 | 5 | 3.7× |
| SonicFastest | 4618 | 392.35 MB/s | 3394 | 40 | 2.0× |
| Sonic | 4833 | 374.92 MB/s | 3393 | 40 | 1.9× |
| JSONV2 | 4888 | 370.69 MB/s | 632 | 7 | 1.9× |
| Stdlib | 9081 | 199.55 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 641 | 2827.85 MB/s | 0 | 0 | 15.3× |
| Easyjson | 1710 | 1059.40 MB/s | 24 | 1 | 5.7× |
| Goccy | 2420 | 748.70 MB/s | 2610 | 4 | 4.0× |
| SonicFastest | 4467 | 405.61 MB/s | 3367 | 38 | 2.2× |
| Sonic | 4653 | 389.44 MB/s | 3367 | 38 | 2.1× |
| JSONV2 | 4835 | 374.74 MB/s | 640 | 6 | 2.0× |
| Stdlib | 9790 | 185.08 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 679 | 2670.27 MB/s | 0 | 0 | 14.0× |
| Easyjson | 1709 | 1060.06 MB/s | 24 | 1 | 5.5× |
| Goccy | 2425 | 747.36 MB/s | 2610 | 4 | 3.9× |
| SonicFastest | 4389 | 412.86 MB/s | 3358 | 38 | 2.2× |
| Sonic | 4592 | 394.60 MB/s | 3359 | 38 | 2.1× |
| JSONV2 | 4686 | 386.68 MB/s | 640 | 6 | 2.0× |
| Stdlib | 9475 | 191.24 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 323 | 711.82 MB/s | 160 | 1 | 8.5× |
| SonicFastest | 665 | 345.92 MB/s | 805 | 8 | 4.1× |
| Sonic | 672 | 342.37 MB/s | 805 | 8 | 4.1× |
| Easyjson | 1036 | 222.06 MB/s | 448 | 3 | 2.7× |
| Goccy | 1127 | 204.07 MB/s | 585 | 23 | 2.4× |
| JSONV2 | 1543 | 149.06 MB/s | 528 | 7 | 1.8× |
| Stdlib | 2755 | 83.48 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 443 | 1114.83 MB/s | 160 | 1 | 8.7× |
| SonicFastest | 940 | 525.57 MB/s | 1082 | 8 | 4.1× |
| Sonic | 947 | 521.78 MB/s | 1082 | 8 | 4.1× |
| Easyjson | 1546 | 319.62 MB/s | 448 | 3 | 2.5× |
| Goccy | 1760 | 280.73 MB/s | 857 | 23 | 2.2× |
| JSONV2 | 2074 | 238.19 MB/s | 528 | 7 | 1.9× |
| Stdlib | 3841 | 128.61 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2429803 | 798.61 MB/s | 2846866 | 2238 | 7.0× |
| Goccy | 3308289 | 586.55 MB/s | 4067253 | 13509 | 5.1× |
| SonicFastest | 3727558 | 520.57 MB/s | 4878770 | 1736 | 4.6× |
| Sonic | 3811333 | 509.13 MB/s | 4879725 | 1736 | 4.5× |
| Easyjson | 5469385 | 354.79 MB/s | 3871269 | 15043 | 3.1× |
| JSONV2 | 7788869 | 249.13 MB/s | 3238677 | 13947 | 2.2× |
| Stdlib | 16978947 | 114.29 MB/s | 3551324 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 9511928 | 851.57 MB/s | 15059839 | 51649 | 6.5× |
| Sonic | 12161532 | 666.04 MB/s | 19887763 | 41641 | 5.1× |
| SonicFastest | 12230134 | 662.30 MB/s | 19894923 | 41641 | 5.1× |
| Goccy | 17966850 | 450.83 MB/s | 17635294 | 107149 | 3.5× |
| Easyjson | 20954759 | 386.55 MB/s | 15059616 | 41643 | 3.0× |
| JSONV2 | 29163819 | 277.74 MB/s | 15235800 | 78974 | 2.1× |
| Stdlib | 62293706 | 130.03 MB/s | 15665074 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 871 | 2515.30 MB/s | 0 | 0 | 12.6× |
| Easyjson | 1985 | 1103.95 MB/s | 24 | 1 | 5.5× |
| Goccy | 2676 | 818.81 MB/s | 2866 | 4 | 4.1× |
| JSONV2 | 4761 | 460.19 MB/s | 640 | 6 | 2.3× |
| SonicFastest | 4882 | 448.82 MB/s | 3624 | 38 | 2.2× |
| Sonic | 5062 | 432.82 MB/s | 3616 | 38 | 2.2× |
| Stdlib | 10956 | 199.98 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 387 | 51133.93 MB/s | 0 | 0 | 225.1× |
| SonicFastest | 3766 | 5254.00 MB/s | 21551 | 3 | 23.1× |
| Goccy | 19178 | 1031.89 MB/s | 20507 | 2 | 4.5× |
| Sonic | 21036 | 940.71 MB/s | 20708 | 3 | 4.1× |
| JSONV2 | 21866 | 905.02 MB/s | 8 | 1 | 4.0× |
| Easyjson | 67362 | 293.77 MB/s | 0 | 0 | 1.3× |
| Stdlib | 87103 | 227.19 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1758 | 10307.10 MB/s | 864 | 4 | 44.3× |
| Easyjson | 2866 | 6323.36 MB/s | 432 | 2 | 27.2× |
| SonicFastest | 6077 | 2982.32 MB/s | 20441 | 5 | 12.8× |
| Sonic | 6189 | 2928.52 MB/s | 20465 | 5 | 12.6× |
| Goccy | 16880 | 1073.69 MB/s | 19475 | 2 | 4.6× |
| JSONV2 | 33181 | 546.22 MB/s | 16514 | 50 | 2.3× |
| Stdlib | 77822 | 232.89 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1962147 | 1023.62 MB/s | 2963985 | 7453 | 7.2× |
| Goccy | 3465923 | 579.50 MB/s | 5466692 | 15852 | 4.1× |
| SonicFastest | 3569012 | 562.76 MB/s | 5154597 | 7085 | 4.0× |
| Sonic | 3584626 | 560.31 MB/s | 5155400 | 7085 | 4.0× |
| Easyjson | 3954745 | 507.87 MB/s | 2982878 | 7441 | 3.6× |
| JSONV2 | 5360851 | 374.66 MB/s | 3174798 | 14563 | 2.6× |
| Stdlib | 14172394 | 141.72 MB/s | 3589386 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 725 | 757.64 MB/s | 480 | 1 | 5.8× |
| Easyjson | 1460 | 376.00 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 1573 | 348.92 MB/s | 2274 | 8 | 2.7× |
| Sonic | 1655 | 331.77 MB/s | 2274 | 8 | 2.5× |
| Goccy | 2209 | 248.56 MB/s | 2132 | 43 | 1.9× |
| JSONV2 | 2215 | 247.83 MB/s | 1665 | 7 | 1.9× |
| Stdlib | 4186 | 131.14 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 445958 | 1416.08 MB/s | 461115 | 1230 | 8.8× |
| Sonic | 845955 | 746.51 MB/s | 1083049 | 814 | 4.6× |
| Easyjson | 853991 | 739.49 MB/s | 422504 | 936 | 4.6× |
| SonicFastest | 896039 | 704.78 MB/s | 1082982 | 814 | 4.4× |
| Goccy | 941160 | 671.00 MB/s | 1003781 | 1204 | 4.2× |
| JSONV2 | 1522003 | 414.92 MB/s | 572023 | 3144 | 2.6× |
| Stdlib | 3918137 | 161.18 MB/s | 654667 | 6472 | 1.0× |
