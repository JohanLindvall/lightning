# JSON Deserialization Benchmarks

- generated 2026-06-20T15:43:02Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 107270 | 1186.49 MB/s | 50208 | 4 | 10.3× |
| SonicFastest | 180452 | 705.31 MB/s | 186878 | 10 | 6.1× |
| Sonic | 181362 | 701.77 MB/s | 187649 | 10 | 6.1× |
| Goccy | 196806 | 646.70 MB/s | 225209 | 884 | 5.6× |
| Easyjson | 212706 | 598.36 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 420527 | 302.66 MB/s | 195115 | 1805 | 2.6× |
| Stdlib | 1105641 | 115.11 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 4000151 | 562.74 MB/s | 3122883 | 3081 | 6.6× |
| SonicFastest | 4526502 | 497.30 MB/s | 15232103 | 970 | 5.8× |
| Sonic | 4627843 | 486.41 MB/s | 15232103 | 970 | 5.7× |
| Goccy | 10311768 | 218.30 MB/s | 4241767 | 56539 | 2.5× |
| Easyjson | 11388968 | 197.65 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 16434132 | 136.97 MB/s | 3123376 | 3083 | 1.6× |
| Stdlib | 26276347 | 85.67 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 483520 | 559.24 MB/s | 348024 | 1627 | 6.9× |
| Sonic | 639485 | 422.85 MB/s | 491750 | 968 | 5.2× |
| SonicFastest | 640786 | 421.99 MB/s | 492421 | 968 | 5.2× |
| Goccy | 1410677 | 191.68 MB/s | 541794 | 8122 | 2.4× |
| Easyjson | 1433238 | 188.67 MB/s | 330272 | 749 | 2.3× |
| JSONV2 | 2107924 | 128.28 MB/s | 348151 | 1628 | 1.6× |
| Stdlib | 3351380 | 80.68 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1483097 | 1164.59 MB/s | 964644 | 6177 | 9.0× |
| SonicFastest | 2134800 | 809.07 MB/s | 2931459 | 4020 | 6.3× |
| Sonic | 2141621 | 806.49 MB/s | 2963299 | 4020 | 6.2× |
| Goccy | 2476140 | 697.54 MB/s | 2587598 | 14607 | 5.4× |
| Easyjson | 4295564 | 402.09 MB/s | 972034 | 5389 | 3.1× |
| JSONV2 | 4321756 | 399.65 MB/s | 1011703 | 7594 | 3.1× |
| Stdlib | 13366384 | 129.22 MB/s | 1234450 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1238 | 1464.01 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2556 | 709.05 MB/s | 24 | 1 | 5.5× |
| Goccy | 2785 | 650.67 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5954 | 304.32 MB/s | 3743 | 40 | 2.4× |
| SonicFastest | 5964 | 303.81 MB/s | 3756 | 40 | 2.3× |
| JSONV2 | 7789 | 232.64 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14011 | 129.33 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1253 | 1445.61 MB/s | 0 | 0 | 11.2× |
| Easyjson | 2542 | 712.72 MB/s | 24 | 1 | 5.5× |
| Goccy | 2817 | 643.29 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 5996 | 302.21 MB/s | 3769 | 40 | 2.3× |
| Sonic | 5998 | 302.11 MB/s | 3759 | 40 | 2.3× |
| JSONV2 | 7852 | 230.78 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14032 | 129.14 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1439 | 1259.31 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2775 | 652.88 MB/s | 144 | 10 | 5.0× |
| Goccy | 2897 | 625.39 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6155 | 294.41 MB/s | 3742 | 42 | 2.3× |
| SonicFastest | 6202 | 292.18 MB/s | 3779 | 42 | 2.3× |
| JSONV2 | 8079 | 224.27 MB/s | 632 | 7 | 1.7× |
| Stdlib | 13988 | 129.54 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 744 | 663.84 MB/s | 160 | 1 | 7.4× |
| SonicFastest | 1240 | 398.38 MB/s | 977 | 6 | 4.5× |
| Sonic | 1243 | 397.28 MB/s | 977 | 6 | 4.4× |
| Easyjson | 2241 | 220.44 MB/s | 448 | 3 | 2.5× |
| Goccy | 2430 | 203.25 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3249 | 152.03 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5529 | 89.35 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 468 | 491.04 MB/s | 160 | 1 | 8.8× |
| SonicFastest | 890 | 258.49 MB/s | 697 | 6 | 4.6× |
| Sonic | 893 | 257.57 MB/s | 692 | 6 | 4.6× |
| Easyjson | 1423 | 161.58 MB/s | 448 | 3 | 2.9× |
| Goccy | 1610 | 142.85 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2452 | 93.80 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4114 | 55.90 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 81193 | 802.19 MB/s | 178609 | 112 | 6.8× |
| Sonic | 99346 | 655.61 MB/s | 155970 | 75 | 5.5× |
| SonicFastest | 103470 | 629.48 MB/s | 159567 | 75 | 5.3× |
| Goccy | 150318 | 433.30 MB/s | 228653 | 134 | 3.7× |
| JSONV2 | 231414 | 281.45 MB/s | 206652 | 607 | 2.4× |
| Stdlib | 550886 | 118.23 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2913891 | 665.94 MB/s | 2846869 | 2238 | 8.1× |
| SonicFastest | 4509672 | 430.29 MB/s | 14608601 | 1407 | 5.2× |
| Sonic | 4684897 | 414.20 MB/s | 14618261 | 1407 | 5.0× |
| Goccy | 4781512 | 405.83 MB/s | 4065220 | 13510 | 4.9× |
| Easyjson | 7729616 | 251.04 MB/s | 3871268 | 15043 | 3.0× |
| JSONV2 | 11486833 | 168.93 MB/s | 3237332 | 13947 | 2.1× |
| Stdlib | 23553588 | 82.39 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2492171 | 1335.31 MB/s | 4611519 | 5958 | 8.5× |
| Sonic | 2821905 | 1179.29 MB/s | 6682006 | 4249 | 7.5× |
| SonicFastest | 2838196 | 1172.52 MB/s | 6659210 | 4249 | 7.4× |
| Goccy | 4677727 | 711.42 MB/s | 3948956 | 3817 | 4.5× |
| JSONV2 | 7685219 | 433.02 MB/s | 5364561 | 13243 | 2.7× |
| Stdlib | 21118629 | 157.58 MB/s | 5565611 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 253803 | 868.18 MB/s | 123584 | 225 | 8.0× |
| Sonic | 376122 | 585.84 MB/s | 296839 | 398 | 5.4× |
| SonicFastest | 378971 | 581.43 MB/s | 304632 | 398 | 5.4× |
| Goccy | 427394 | 515.56 MB/s | 362799 | 1066 | 4.8× |
| Easyjson | 547062 | 402.78 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 723736 | 304.46 MB/s | 129738 | 470 | 2.8× |
| Stdlib | 2037251 | 108.16 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 14010805 | 578.13 MB/s | 15059836 | 51649 | 6.4× |
| Sonic | 17183990 | 471.37 MB/s | 72288724 | 40015 | 5.3× |
| SonicFastest | 17243959 | 469.73 MB/s | 72382758 | 40015 | 5.2× |
| Goccy | 24146597 | 335.45 MB/s | 17283231 | 107151 | 3.7× |
| Easyjson | 31486071 | 257.26 MB/s | 15059619 | 41643 | 2.9× |
| JSONV2 | 45124304 | 179.51 MB/s | 15233884 | 78973 | 2.0× |
| Stdlib | 90222590 | 89.78 MB/s | 15665088 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6490213 | 459.69 MB/s | 3985338 | 30015 | 7.4× |
| SonicFastest | 8419652 | 354.35 MB/s | 25564632 | 56758 | 5.7× |
| Sonic | 8469557 | 352.26 MB/s | 25670589 | 56758 | 5.6× |
| Goccy | 16654954 | 179.13 MB/s | 10918237 | 273667 | 2.9× |
| Easyjson | 16913398 | 176.40 MB/s | 9479451 | 30115 | 2.8× |
| JSONV2 | 25243693 | 118.19 MB/s | 9257241 | 86279 | 1.9× |
| Stdlib | 47838370 | 62.37 MB/s | 9258098 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1407600 | 514.06 MB/s | 907387 | 3618 | 8.3× |
| SonicFastest | 1775207 | 407.61 MB/s | 3233826 | 7226 | 6.6× |
| Sonic | 1800599 | 401.86 MB/s | 3287665 | 7226 | 6.5× |
| Easyjson | 4326418 | 167.25 MB/s | 2847912 | 3698 | 2.7× |
| Goccy | 4952542 | 146.11 MB/s | 3282326 | 80298 | 2.4× |
| JSONV2 | 5722363 | 126.45 MB/s | 2704785 | 7319 | 2.1× |
| Stdlib | 11740388 | 61.63 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2028371 | 777.65 MB/s | 907387 | 3618 | 7.8× |
| Sonic | 2290836 | 688.55 MB/s | 5780994 | 7226 | 6.9× |
| SonicFastest | 2310348 | 682.73 MB/s | 5782463 | 7226 | 6.9× |
| Easyjson | 5651972 | 279.08 MB/s | 2847911 | 3698 | 2.8× |
| Goccy | 5722333 | 275.65 MB/s | 3656366 | 80272 | 2.8× |
| JSONV2 | 6464994 | 243.98 MB/s | 2704770 | 7319 | 2.5× |
| Stdlib | 15842927 | 99.56 MB/s | 2704555 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224405 | 668.99 MB/s | 81920 | 1 | 8.2× |
| SonicFastest | 272971 | 549.96 MB/s | 257397 | 6 | 6.8× |
| Sonic | 276305 | 543.33 MB/s | 265751 | 6 | 6.7× |
| Goccy | 853119 | 175.97 MB/s | 319280 | 10004 | 2.2× |
| JSONV2 | 1093283 | 137.31 MB/s | 357713 | 20 | 1.7× |
| Stdlib | 1843354 | 81.44 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 38183 | 736.37 MB/s | 29104 | 107 | 7.9× |
| Sonic | 63776 | 440.87 MB/s | 46972 | 103 | 4.7× |
| SonicFastest | 64298 | 437.29 MB/s | 47430 | 103 | 4.7× |
| Easyjson | 67876 | 414.24 MB/s | 32304 | 138 | 4.5× |
| Goccy | 71052 | 395.72 MB/s | 59203 | 188 | 4.3× |
| JSONV2 | 133104 | 211.24 MB/s | 36896 | 242 | 2.3× |
| Stdlib | 302602 | 92.92 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2039 | 1141.81 MB/s | 32 | 1 | 11.1× |
| Goccy | 4116 | 565.60 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4211 | 552.79 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5078 | 458.42 MB/s | 4155 | 6 | 4.5× |
| Sonic | 5089 | 457.46 MB/s | 4179 | 6 | 4.4× |
| JSONV2 | 8399 | 277.18 MB/s | 1000 | 6 | 2.7× |
| Stdlib | 22635 | 102.85 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 228 | 827.56 MB/s | 0 | 0 | 10.6× |
| Goccy | 388 | 486.68 MB/s | 304 | 2 | 6.2× |
| Easyjson | 495 | 381.95 MB/s | 0 | 0 | 4.9× |
| SonicFastest | 781 | 241.96 MB/s | 510 | 4 | 3.1× |
| Sonic | 783 | 241.30 MB/s | 505 | 4 | 3.1× |
| JSONV2 | 1030 | 183.45 MB/s | 112 | 1 | 2.3× |
| Stdlib | 2414 | 78.29 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1546 | 1417.41 MB/s | 0 | 0 | 10.4× |
| Easyjson | 3185 | 687.95 MB/s | 24 | 1 | 5.0× |
| Goccy | 3268 | 670.45 MB/s | 2864 | 4 | 4.9× |
| Sonic | 6395 | 342.61 MB/s | 3943 | 40 | 2.5× |
| SonicFastest | 6485 | 337.86 MB/s | 4028 | 40 | 2.5× |
| JSONV2 | 8160 | 268.50 MB/s | 640 | 6 | 2.0× |
| Stdlib | 16007 | 136.88 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 889390 | 573.96 MB/s | 703777 | 1012 | 6.8× |
| Sonic | 1152313 | 443.00 MB/s | 909457 | 2006 | 5.3× |
| SonicFastest | 1159817 | 440.14 MB/s | 919323 | 2006 | 5.2× |
| Goccy | 1172223 | 435.48 MB/s | 1140058 | 5007 | 5.2× |
| Easyjson | 1545506 | 330.30 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3164787 | 161.30 MB/s | 1076002 | 12646 | 1.9× |
| Stdlib | 6069731 | 84.10 MB/s | 1162115 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2062 | 9595.08 MB/s | 0 | 0 | 54.2× |
| Goccy | 20447 | 967.80 MB/s | 20491 | 2 | 5.5× |
| Sonic | 27187 | 727.87 MB/s | 22139 | 4 | 4.1× |
| SonicFastest | 27728 | 713.68 MB/s | 23381 | 4 | 4.0× |
| JSONV2 | 29732 | 665.58 MB/s | 8 | 1 | 3.8× |
| Easyjson | 86041 | 230.00 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111660 | 177.23 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2919 | 6209.18 MB/s | 864 | 4 | 35.2× |
| Easyjson | 3941 | 4598.97 MB/s | 432 | 2 | 26.1× |
| SonicFastest | 10028 | 1807.33 MB/s | 22770 | 6 | 10.2× |
| Sonic | 10172 | 1781.73 MB/s | 23194 | 6 | 10.1× |
| Goccy | 15868 | 1142.16 MB/s | 19459 | 2 | 6.5× |
| JSONV2 | 44757 | 404.94 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102713 | 176.45 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2638906 | 761.11 MB/s | 2963987 | 7453 | 7.2× |
| Sonic | 4218300 | 476.14 MB/s | 10031253 | 13682 | 4.5× |
| SonicFastest | 4253358 | 472.21 MB/s | 10035700 | 13682 | 4.4× |
| Goccy | 4483072 | 448.02 MB/s | 5416454 | 15851 | 4.2× |
| Easyjson | 5212020 | 385.36 MB/s | 2981777 | 7442 | 3.6× |
| JSONV2 | 7269606 | 276.29 MB/s | 3173818 | 14563 | 2.6× |
| Stdlib | 18914936 | 106.19 MB/s | 3589321 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 898 | 611.65 MB/s | 480 | 1 | 6.3× |
| Easyjson | 2195 | 250.13 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2697 | 203.54 MB/s | 1940 | 26 | 2.1× |
| Sonic | 2709 | 202.66 MB/s | 1947 | 26 | 2.1× |
| Goccy | 3064 | 179.15 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3337 | 164.50 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5697 | 96.36 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 564041 | 1119.63 MB/s | 461113 | 1230 | 9.5× |
| Sonic | 1026628 | 615.13 MB/s | 1020821 | 1102 | 5.2× |
| SonicFastest | 1036721 | 609.15 MB/s | 1032151 | 1102 | 5.2× |
| Easyjson | 1130402 | 558.66 MB/s | 422504 | 936 | 4.8× |
| Goccy | 1180821 | 534.81 MB/s | 989496 | 1203 | 4.6× |
| JSONV2 | 2188030 | 288.62 MB/s | 571652 | 3144 | 2.5× |
| Stdlib | 5372742 | 117.54 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 924940 | 608.05 MB/s | 886430 | 2062 | 5.7× |
| Sonic | 1096001 | 513.15 MB/s | 1010383 | 1476 | 4.8× |
| SonicFastest | 1099178 | 511.66 MB/s | 1025913 | 1476 | 4.8× |
| Goccy | 1371237 | 410.15 MB/s | 1047123 | 1033 | 3.8× |
| Easyjson | 1708110 | 329.26 MB/s | 775153 | 1254 | 3.1× |
| JSONV2 | 2851244 | 197.25 MB/s | 927492 | 3483 | 1.8× |
| Stdlib | 5232887 | 107.48 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 775023 | 687.95 MB/s | 413025 | 3155 | 7.0× |
| Easyjson | 1103207 | 483.30 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1124649 | 474.08 MB/s | 1002998 | 4351 | 4.8× |
| Sonic | 1140608 | 467.45 MB/s | 1024614 | 4351 | 4.8× |
| Goccy | 1323693 | 402.80 MB/s | 1167463 | 5410 | 4.1× |
| JSONV2 | 2566698 | 207.73 MB/s | 745497 | 13288 | 2.1× |
| Stdlib | 5419369 | 98.38 MB/s | 798692 | 17133 | 1.0× |
