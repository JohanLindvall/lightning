# JSON Deserialization Benchmarks

- generated 2026-06-18T20:49:08Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 483378 | 559.40 MB/s | 348024 | 1627 | 6.9× |
| SonicFastest | 636524 | 424.81 MB/s | 487626 | 968 | 5.3× |
| Sonic | 638911 | 423.23 MB/s | 502051 | 968 | 5.2× |
| Goccy | 1379527 | 196.01 MB/s | 541023 | 8122 | 2.4× |
| Easyjson | 1411312 | 191.60 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2081699 | 129.90 MB/s | 348154 | 1628 | 1.6× |
| Stdlib | 3350594 | 80.70 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1570380 | 1099.86 MB/s | 964644 | 6177 | 8.6× |
| Sonic | 2067801 | 835.29 MB/s | 2892135 | 4020 | 6.5× |
| SonicFastest | 2069926 | 834.43 MB/s | 2951100 | 4020 | 6.5× |
| Goccy | 2437309 | 708.65 MB/s | 2588000 | 14608 | 5.5× |
| Easyjson | 4288466 | 402.76 MB/s | 972033 | 5389 | 3.1× |
| JSONV2 | 4323518 | 399.49 MB/s | 1011699 | 7594 | 3.1× |
| Stdlib | 13457649 | 128.34 MB/s | 1234451 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1239 | 1463.00 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2550 | 710.63 MB/s | 24 | 1 | 5.5× |
| Goccy | 2743 | 660.48 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5889 | 307.68 MB/s | 3739 | 40 | 2.4× |
| Sonic | 5908 | 306.69 MB/s | 3773 | 40 | 2.4× |
| JSONV2 | 7745 | 233.96 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14056 | 128.91 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1253 | 1445.78 MB/s | 0 | 0 | 11.2× |
| Easyjson | 2568 | 705.60 MB/s | 24 | 1 | 5.5× |
| Goccy | 2837 | 638.66 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5909 | 306.66 MB/s | 3743 | 40 | 2.4× |
| SonicFastest | 5962 | 303.91 MB/s | 3792 | 40 | 2.4× |
| JSONV2 | 7838 | 231.19 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14051 | 128.95 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1439 | 1258.81 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2777 | 652.43 MB/s | 144 | 10 | 5.0× |
| Goccy | 2820 | 642.49 MB/s | 2600 | 5 | 5.0× |
| Sonic | 6052 | 299.40 MB/s | 3765 | 42 | 2.3× |
| SonicFastest | 6057 | 299.17 MB/s | 3746 | 42 | 2.3× |
| JSONV2 | 7986 | 226.89 MB/s | 632 | 7 | 1.8× |
| Stdlib | 13976 | 129.65 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 743 | 665.26 MB/s | 160 | 1 | 7.5× |
| Sonic | 1240 | 398.44 MB/s | 978 | 6 | 4.5× |
| SonicFastest | 1242 | 397.61 MB/s | 983 | 6 | 4.5× |
| Easyjson | 2222 | 222.36 MB/s | 448 | 3 | 2.5× |
| Goccy | 2450 | 201.67 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3265 | 151.31 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5536 | 89.23 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 461 | 498.56 MB/s | 160 | 1 | 9.0× |
| SonicFastest | 901 | 255.28 MB/s | 687 | 6 | 4.6× |
| Sonic | 905 | 254.16 MB/s | 691 | 6 | 4.6× |
| Easyjson | 1423 | 161.59 MB/s | 448 | 3 | 2.9× |
| Goccy | 1605 | 143.30 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2441 | 94.21 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4167 | 55.19 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3383953 | 573.43 MB/s | 2846871 | 2238 | 7.0× |
| Sonic | 4675735 | 415.01 MB/s | 14608568 | 1407 | 5.0× |
| Goccy | 4751923 | 408.36 MB/s | 4065962 | 13510 | 5.0× |
| SonicFastest | 4888875 | 396.92 MB/s | 14613580 | 1407 | 4.8× |
| Easyjson | 7746100 | 250.51 MB/s | 3871267 | 15043 | 3.0× |
| JSONV2 | 11428952 | 169.79 MB/s | 3237298 | 13947 | 2.1× |
| Stdlib | 23560757 | 82.36 MB/s | 3551324 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13886929 | 583.29 MB/s | 15059842 | 51649 | 6.5× |
| SonicFastest | 17079796 | 474.25 MB/s | 72111828 | 40015 | 5.3× |
| Sonic | 17125224 | 472.99 MB/s | 72199847 | 40015 | 5.3× |
| Goccy | 24061280 | 336.64 MB/s | 17474376 | 107151 | 3.7× |
| Easyjson | 31189910 | 259.70 MB/s | 15059620 | 41643 | 2.9× |
| JSONV2 | 44131779 | 183.54 MB/s | 15233890 | 78973 | 2.0× |
| Stdlib | 90101894 | 89.90 MB/s | 15665073 | 150647 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1553 | 1410.38 MB/s | 0 | 0 | 10.3× |
| Goccy | 3139 | 697.92 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3222 | 680.09 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6313 | 347.07 MB/s | 3946 | 40 | 2.5× |
| Sonic | 6332 | 346.02 MB/s | 4001 | 40 | 2.5× |
| JSONV2 | 8013 | 273.45 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15958 | 137.29 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2062 | 9596.98 MB/s | 0 | 0 | 54.2× |
| Goccy | 19975 | 990.71 MB/s | 20491 | 2 | 5.6× |
| Sonic | 27911 | 709.01 MB/s | 21936 | 4 | 4.0× |
| SonicFastest | 27968 | 707.57 MB/s | 22131 | 4 | 4.0× |
| JSONV2 | 29683 | 666.68 MB/s | 8 | 1 | 3.8× |
| Easyjson | 86037 | 230.01 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111831 | 176.96 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2972 | 6098.19 MB/s | 864 | 4 | 34.5× |
| Easyjson | 3945 | 4594.33 MB/s | 432 | 2 | 26.0× |
| Sonic | 9961 | 1819.51 MB/s | 22798 | 6 | 10.3× |
| SonicFastest | 10002 | 1811.96 MB/s | 22982 | 6 | 10.3× |
| Goccy | 15554 | 1165.23 MB/s | 19459 | 2 | 6.6× |
| JSONV2 | 45400 | 399.21 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102670 | 176.53 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2569317 | 781.72 MB/s | 2963985 | 7453 | 7.3× |
| SonicFastest | 4085007 | 491.67 MB/s | 10020888 | 13682 | 4.6× |
| Sonic | 4138365 | 485.34 MB/s | 10025123 | 13682 | 4.5× |
| Goccy | 4315191 | 465.45 MB/s | 5417147 | 15845 | 4.4× |
| Easyjson | 5117574 | 392.47 MB/s | 2981719 | 7441 | 3.7× |
| JSONV2 | 7106398 | 282.63 MB/s | 3173801 | 14563 | 2.6× |
| Stdlib | 18796931 | 106.85 MB/s | 3589320 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1125 | 487.93 MB/s | 480 | 1 | 5.0× |
| Easyjson | 2140 | 256.55 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2683 | 204.62 MB/s | 1942 | 26 | 2.1× |
| SonicFastest | 2696 | 203.62 MB/s | 1955 | 26 | 2.1× |
| Goccy | 3025 | 181.51 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3245 | 169.19 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5630 | 97.51 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 570409 | 1107.13 MB/s | 461113 | 1230 | 9.4× |
| SonicFastest | 1009527 | 625.55 MB/s | 993959 | 1102 | 5.3× |
| Sonic | 1017775 | 620.48 MB/s | 1012007 | 1102 | 5.3× |
| Easyjson | 1134085 | 556.85 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1170311 | 539.61 MB/s | 988669 | 1203 | 4.6× |
| JSONV2 | 2178801 | 289.84 MB/s | 571653 | 3144 | 2.5× |
| Stdlib | 5366402 | 117.68 MB/s | 654665 | 6472 | 1.0× |
