# JSON Deserialization Benchmarks

- generated 2026-06-17T19:44:31Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 485825 | 556.59 MB/s | 151224 | 495 | 5.3× |
| Sonic | 552766 | 489.18 MB/s | 644827 | 1147 | 4.6× |
| SonicFastest | 558825 | 483.88 MB/s | 644817 | 1147 | 4.6× |
| Easyjson | 1262226 | 214.23 MB/s | 330320 | 753 | 2.0× |
| Goccy | 1275672 | 211.97 MB/s | 544355 | 8122 | 2.0× |
| JSONV2 | 1635844 | 165.30 MB/s | 348578 | 1628 | 1.6× |
| Stdlib | 2562526 | 105.52 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1377731 | 1253.66 MB/s | 979492 | 6803 | 6.7× |
| SonicFastest | 1741455 | 991.82 MB/s | 2730263 | 5548 | 5.3× |
| Sonic | 1746856 | 988.75 MB/s | 2734694 | 5548 | 5.3× |
| Goccy | 1922716 | 898.31 MB/s | 2603137 | 14609 | 4.8× |
| JSONV2 | 3017835 | 572.33 MB/s | 1012353 | 7594 | 3.1× |
| Easyjson | 3148266 | 548.62 MB/s | 986873 | 6015 | 2.9× |
| Stdlib | 9240522 | 186.92 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 896 | 2021.50 MB/s | 144 | 10 | 10.4× |
| Easyjson | 1841 | 984.07 MB/s | 144 | 10 | 5.1× |
| Goccy | 2564 | 706.82 MB/s | 2603 | 5 | 3.6× |
| SonicFastest | 4513 | 401.50 MB/s | 3379 | 40 | 2.1× |
| Sonic | 4679 | 387.25 MB/s | 3378 | 40 | 2.0× |
| JSONV2 | 5023 | 360.77 MB/s | 633 | 7 | 1.9× |
| Stdlib | 9332 | 194.17 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 615 | 2946.19 MB/s | 0 | 0 | 15.6× |
| Easyjson | 1680 | 1078.57 MB/s | 24 | 1 | 5.7× |
| Goccy | 2647 | 684.66 MB/s | 2611 | 4 | 3.6× |
| SonicFastest | 4343 | 417.20 MB/s | 3365 | 38 | 2.2× |
| Sonic | 4553 | 397.97 MB/s | 3361 | 38 | 2.1× |
| JSONV2 | 4891 | 370.48 MB/s | 641 | 6 | 2.0× |
| Stdlib | 9596 | 188.83 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 689 | 2628.47 MB/s | 0 | 0 | 13.7× |
| Easyjson | 1676 | 1081.15 MB/s | 24 | 1 | 5.6× |
| Goccy | 2459 | 737.03 MB/s | 2611 | 4 | 3.8× |
| SonicFastest | 4328 | 418.62 MB/s | 3362 | 38 | 2.2× |
| Sonic | 4508 | 401.98 MB/s | 3360 | 38 | 2.1× |
| JSONV2 | 5133 | 353.02 MB/s | 641 | 6 | 1.8× |
| Stdlib | 9442 | 191.91 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 340 | 676.07 MB/s | 160 | 1 | 8.8× |
| Sonic | 665 | 345.89 MB/s | 812 | 8 | 4.5× |
| SonicFastest | 670 | 343.39 MB/s | 813 | 8 | 4.5× |
| Easyjson | 1133 | 202.99 MB/s | 448 | 3 | 2.7× |
| Goccy | 1287 | 178.64 MB/s | 585 | 23 | 2.3× |
| JSONV2 | 1703 | 135.05 MB/s | 528 | 7 | 1.8× |
| Stdlib | 3008 | 76.47 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 535 | 923.09 MB/s | 160 | 1 | 8.0× |
| Sonic | 918 | 538.05 MB/s | 1091 | 8 | 4.7× |
| SonicFastest | 947 | 521.48 MB/s | 1091 | 8 | 4.5× |
| Easyjson | 1742 | 283.59 MB/s | 448 | 3 | 2.5× |
| Goccy | 1934 | 255.47 MB/s | 858 | 23 | 2.2× |
| JSONV2 | 2187 | 225.91 MB/s | 528 | 7 | 2.0× |
| Stdlib | 4271 | 115.67 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3018772 | 642.80 MB/s | 3258552 | 15044 | 5.5× |
| Goccy | 3266924 | 593.98 MB/s | 4067021 | 13509 | 5.1× |
| SonicFastest | 3539132 | 548.29 MB/s | 4885701 | 1736 | 4.7× |
| Sonic | 3554300 | 545.95 MB/s | 4885842 | 1736 | 4.7× |
| Easyjson | 6405708 | 302.93 MB/s | 4282952 | 27849 | 2.6× |
| JSONV2 | 8134160 | 238.56 MB/s | 3239113 | 13947 | 2.0× |
| Stdlib | 16554226 | 117.22 MB/s | 3551319 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 10743557 | 753.94 MB/s | 19859940 | 41641 | 5.9× |
| SonicFastest | 10848219 | 746.67 MB/s | 19860034 | 41641 | 5.8× |
| Lightning | 12164686 | 665.87 MB/s | 12528037 | 40027 | 5.2× |
| Goccy | 18551379 | 436.63 MB/s | 17461568 | 107150 | 3.4× |
| Easyjson | 24704783 | 327.87 MB/s | 15219647 | 61644 | 2.6× |
| JSONV2 | 32210681 | 251.47 MB/s | 15236432 | 78973 | 2.0× |
| Stdlib | 63242045 | 128.08 MB/s | 15665065 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 857 | 2557.21 MB/s | 0 | 0 | 12.7× |
| Easyjson | 2084 | 1051.55 MB/s | 24 | 1 | 5.2× |
| Goccy | 2966 | 738.65 MB/s | 2867 | 4 | 3.7× |
| SonicFastest | 4914 | 445.88 MB/s | 3628 | 38 | 2.2× |
| Sonic | 4978 | 440.17 MB/s | 3627 | 38 | 2.2× |
| JSONV2 | 5244 | 417.81 MB/s | 641 | 6 | 2.1× |
| Stdlib | 10888 | 201.23 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 505 | 39155.02 MB/s | 0 | 0 | 153.9× |
| SonicFastest | 3743 | 5287.05 MB/s | 21436 | 3 | 20.8× |
| Goccy | 20337 | 973.03 MB/s | 20515 | 2 | 3.8× |
| JSONV2 | 21627 | 915.02 MB/s | 8 | 1 | 3.6× |
| Sonic | 26996 | 733.05 MB/s | 20709 | 3 | 2.9× |
| Easyjson | 47468 | 416.89 MB/s | 0 | 0 | 1.6× |
| Stdlib | 77761 | 254.48 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 5351 | 3387.34 MB/s | 18080 | 62 | 13.3× |
| SonicFastest | 5779 | 3136.29 MB/s | 20752 | 5 | 12.3× |
| Sonic | 5943 | 3049.88 MB/s | 20747 | 5 | 12.0× |
| Easyjson | 6969 | 2600.68 MB/s | 17648 | 60 | 10.2× |
| Goccy | 17207 | 1053.28 MB/s | 19483 | 2 | 4.1× |
| JSONV2 | 34799 | 520.82 MB/s | 16520 | 50 | 2.0× |
| Stdlib | 71174 | 254.64 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2605269 | 770.94 MB/s | 3584576 | 29240 | 5.1× |
| SonicFastest | 3098944 | 648.12 MB/s | 5158483 | 7085 | 4.3× |
| Sonic | 3124694 | 642.78 MB/s | 5158097 | 7085 | 4.3× |
| Goccy | 3519271 | 570.71 MB/s | 5485009 | 15840 | 3.8× |
| Easyjson | 4686797 | 428.54 MB/s | 3604045 | 29228 | 2.9× |
| JSONV2 | 5147827 | 390.16 MB/s | 3175169 | 14563 | 2.6× |
| Stdlib | 13414643 | 149.72 MB/s | 3589322 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 811 | 676.90 MB/s | 480 | 1 | 5.3× |
| SonicFastest | 1702 | 322.59 MB/s | 2283 | 8 | 2.5× |
| Easyjson | 1710 | 321.11 MB/s | 1616 | 5 | 2.5× |
| Sonic | 1786 | 307.37 MB/s | 2282 | 8 | 2.4× |
| Goccy | 2432 | 225.78 MB/s | 2133 | 43 | 1.8× |
| JSONV2 | 2435 | 225.49 MB/s | 1666 | 7 | 1.8× |
| Stdlib | 4326 | 126.92 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 603929 | 1045.68 MB/s | 630174 | 5499 | 6.4× |
| SonicFastest | 786410 | 803.03 MB/s | 1087506 | 814 | 4.9× |
| Sonic | 792583 | 796.78 MB/s | 1086732 | 814 | 4.9× |
| Easyjson | 1054758 | 598.73 MB/s | 591597 | 5205 | 3.7× |
| Goccy | 1055802 | 598.14 MB/s | 1003648 | 1203 | 3.7× |
| JSONV2 | 1587691 | 397.76 MB/s | 572133 | 3144 | 2.4× |
| Stdlib | 3859741 | 163.62 MB/s | 654666 | 6472 | 1.0× |
