# JSON Deserialization Benchmarks

- generated 2026-08-14T16:28:04Z
- go version go1.26.5 linux/amd64
- cpu: AMD EPYC 7763 64-Core Processor (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 103460 | 1230.19 MB/s | 49760 | 3 | 13.1× |
| LightningArena | 103510 | 1229.59 MB/s | 49760 | 3 | 13.1× |
| LightningDestructive | 105111 | 1210.87 MB/s | 49280 | 2 | 12.9× |
| Sonic | 197929 | 643.03 MB/s | 214435 | 15 | 6.8× |
| SonicFastest | 203087 | 626.70 MB/s | 215453 | 15 | 6.7× |
| Goccy | 248110 | 512.98 MB/s | 225611 | 884 | 5.4× |
| Easyjson | 249122 | 510.89 MB/s | 122864 | 14 | 5.4× |
| LightningDecodeAny | 442329 | 213.99 MB/s | 463411 | 9708 | 3.1× |
| JSONV2 | 468653 | 271.58 MB/s | 195128 | 1805 | 2.9× |
| Stdlib | 1351469 | 94.18 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 4052504 | 555.47 MB/s | 2532848 | 1143 | 8.6× |
| Lightning | 4087487 | 550.72 MB/s | 2532850 | 1143 | 8.5× |
| LightningArena | 4093624 | 549.89 MB/s | 2532849 | 1143 | 8.5× |
| SonicFastest | 5267925 | 427.31 MB/s | 4878313 | 2584 | 6.6× |
| Sonic | 5287887 | 425.70 MB/s | 4877361 | 2584 | 6.6× |
| Goccy | 13373245 | 168.32 MB/s | 4211632 | 56536 | 2.6× |
| LightningDecodeAny | 13722910 | 164.03 MB/s | 19380209 | 223896 | 2.5× |
| Easyjson | 13802995 | 163.08 MB/s | 3099810 | 2120 | 2.5× |
| JSONV2 | 16819775 | 133.83 MB/s | 3123197 | 3083 | 2.1× |
| Stdlib | 34652894 | 64.96 MB/s | 3123392 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 555311 | 486.94 MB/s | 397297 | 567 | 7.9× |
| LightningArena | 555876 | 486.44 MB/s | 397297 | 567 | 7.9× |
| LightningDestructive | 565953 | 477.78 MB/s | 397296 | 567 | 7.8× |
| SonicFastest | 761224 | 355.22 MB/s | 642029 | 1147 | 5.8× |
| Sonic | 762630 | 354.57 MB/s | 641954 | 1147 | 5.8× |
| Goccy | 1814663 | 149.01 MB/s | 541440 | 8122 | 2.4× |
| Easyjson | 1816275 | 148.88 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2244336 | 120.48 MB/s | 348160 | 1628 | 2.0× |
| LightningDecodeAny | 2259637 | 119.67 MB/s | 2543877 | 29687 | 1.9× |
| Stdlib | 4405431 | 61.38 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1131634 | 1526.29 MB/s | 765560 | 2798 | 15.2× |
| LightningArena | 1160586 | 1488.22 MB/s | 768416 | 2440 | 14.9× |
| Lightning | 1169328 | 1477.09 MB/s | 765601 | 2799 | 14.8× |
| Sonic | 2075224 | 832.30 MB/s | 2694464 | 5547 | 8.3× |
| SonicFastest | 2098043 | 823.25 MB/s | 2694990 | 5547 | 8.2× |
| Goccy | 2624624 | 658.08 MB/s | 2581009 | 14603 | 6.6× |
| Easyjson | 3843390 | 449.40 MB/s | 972032 | 5389 | 4.5× |
| LightningDecodeAny | 4055117 | 123.37 MB/s | 4953692 | 76576 | 4.3× |
| JSONV2 | 4748171 | 363.76 MB/s | 1011614 | 7594 | 3.6× |
| Stdlib | 17250922 | 100.12 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 997 | 1817.18 MB/s | 0 | 0 | 16.4× |
| Lightning | 1003 | 1806.59 MB/s | 0 | 0 | 16.3× |
| LightningDestructive | 1067 | 1697.90 MB/s | 0 | 0 | 15.3× |
| Easyjson | 2994 | 605.27 MB/s | 24 | 1 | 5.5× |
| Goccy | 3350 | 540.84 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 6347 | 285.50 MB/s | 3347 | 38 | 2.6× |
| Sonic | 6746 | 268.62 MB/s | 3348 | 38 | 2.4× |
| JSONV2 | 8255 | 219.51 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 9140 | 198.15 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16328 | 110.98 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1102 | 1644.48 MB/s | 0 | 0 | 15.2× |
| Lightning | 1120 | 1618.03 MB/s | 0 | 0 | 15.0× |
| LightningDestructive | 1163 | 1557.67 MB/s | 0 | 0 | 14.4× |
| Easyjson | 3025 | 598.99 MB/s | 24 | 1 | 5.5× |
| Goccy | 3231 | 560.76 MB/s | 2608 | 4 | 5.2× |
| SonicFastest | 6125 | 295.84 MB/s | 3345 | 38 | 2.7× |
| Sonic | 6328 | 286.36 MB/s | 3347 | 38 | 2.7× |
| JSONV2 | 8321 | 217.75 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8989 | 201.46 MB/s | 7552 | 158 | 1.9× |
| Stdlib | 16783 | 107.96 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1291 | 1403.73 MB/s | 144 | 10 | 12.8× |
| Lightning | 1295 | 1399.64 MB/s | 144 | 10 | 12.8× |
| LightningDestructive | 1350 | 1342.20 MB/s | 144 | 10 | 12.3× |
| Easyjson | 3147 | 575.76 MB/s | 144 | 10 | 5.3× |
| Goccy | 3488 | 519.52 MB/s | 2600 | 5 | 4.7× |
| SonicFastest | 6441 | 281.31 MB/s | 3367 | 40 | 2.6× |
| Sonic | 6667 | 271.79 MB/s | 3370 | 40 | 2.5× |
| JSONV2 | 8670 | 208.99 MB/s | 632 | 7 | 1.9× |
| LightningDecodeAny | 9056 | 199.97 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 16566 | 109.38 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 707 | 698.60 MB/s | 160 | 1 | 9.4× |
| LightningDestructive | 724 | 682.45 MB/s | 160 | 1 | 9.1× |
| Sonic | 1272 | 388.24 MB/s | 1076 | 8 | 5.2× |
| SonicFastest | 1272 | 388.49 MB/s | 1076 | 8 | 5.2× |
| LightningDecodeAny | 1538 | 320.49 MB/s | 1296 | 26 | 4.3× |
| LightningArena | 1629 | 303.26 MB/s | 4096 | 1 | 4.1× |
| Goccy | 2542 | 194.31 MB/s | 856 | 23 | 2.6× |
| Easyjson | 2650 | 186.41 MB/s | 448 | 3 | 2.5× |
| JSONV2 | 3393 | 145.60 MB/s | 528 | 7 | 1.9× |
| Stdlib | 6614 | 74.69 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 454 | 507.17 MB/s | 160 | 1 | 10.6× |
| Lightning | 455 | 505.07 MB/s | 160 | 1 | 10.5× |
| Sonic | 929 | 247.66 MB/s | 801 | 8 | 5.2× |
| SonicFastest | 946 | 243.11 MB/s | 800 | 8 | 5.1× |
| LightningDecodeAny | 1280 | 178.95 MB/s | 1296 | 26 | 3.7× |
| LightningArena | 1324 | 173.66 MB/s | 4096 | 1 | 3.6× |
| Easyjson | 1648 | 139.57 MB/s | 448 | 3 | 2.9× |
| Goccy | 1765 | 130.28 MB/s | 584 | 23 | 2.7× |
| JSONV2 | 2642 | 87.05 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4790 | 48.02 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 67130 | 970.24 MB/s | 103441 | 103 | 10.1× |
| LightningArena | 68227 | 954.64 MB/s | 103441 | 103 | 10.0× |
| LightningDestructive | 76081 | 856.08 MB/s | 97220 | 98 | 9.0× |
| Sonic | 144606 | 450.41 MB/s | 235893 | 65 | 4.7× |
| SonicFastest | 145632 | 447.24 MB/s | 235825 | 65 | 4.7× |
| Goccy | 182382 | 357.12 MB/s | 227986 | 134 | 3.7× |
| LightningDecodeAny | 203730 | 261.76 MB/s | 180049 | 3245 | 3.3× |
| JSONV2 | 262275 | 248.33 MB/s | 206664 | 607 | 2.6× |
| Stdlib | 680986 | 95.64 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2520161 | 769.98 MB/s | 2864593 | 1380 | 11.0× |
| LightningArena | 2585365 | 750.56 MB/s | 2864593 | 1380 | 10.7× |
| Lightning | 2587622 | 749.91 MB/s | 2864593 | 1380 | 10.7× |
| Sonic | 4993309 | 388.61 MB/s | 4879649 | 1736 | 5.5× |
| SonicFastest | 5129024 | 378.33 MB/s | 4881796 | 1736 | 5.4× |
| Goccy | 5188723 | 373.98 MB/s | 4063711 | 13509 | 5.3× |
| Easyjson | 8217333 | 236.14 MB/s | 3871266 | 15043 | 3.4× |
| LightningDecodeAny | 9592874 | 202.28 MB/s | 7063040 | 218633 | 2.9× |
| JSONV2 | 12986692 | 149.42 MB/s | 3237189 | 13947 | 2.1× |
| Stdlib | 27605114 | 70.29 MB/s | 3551316 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1057777 | 3146.06 MB/s | 351704 | 1286 | 22.7× |
| LightningArena | 1589381 | 2093.79 MB/s | 2488907 | 2995 | 15.1× |
| Lightning | 1611910 | 2064.53 MB/s | 2488908 | 2995 | 14.9× |
| Sonic | 2025071 | 1643.32 MB/s | 5896341 | 4263 | 11.8× |
| SonicFastest | 2039335 | 1631.82 MB/s | 5895987 | 4263 | 11.7× |
| LightningDecodeAny | 3648015 | 842.59 MB/s | 4876913 | 56892 | 6.6× |
| Goccy | 5027714 | 661.90 MB/s | 3948915 | 3817 | 4.8× |
| JSONV2 | 7872883 | 422.70 MB/s | 5364505 | 13243 | 3.0× |
| Stdlib | 23959612 | 138.89 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 215496 | 1022.50 MB/s | 135872 | 226 | 11.1× |
| Lightning | 215688 | 1021.60 MB/s | 135872 | 226 | 11.1× |
| LightningDestructive | 231270 | 952.77 MB/s | 135872 | 226 | 10.4× |
| SonicFastest | 491599 | 448.22 MB/s | 350621 | 262 | 4.9× |
| Goccy | 497943 | 442.51 MB/s | 363675 | 1066 | 4.8× |
| Sonic | 502977 | 438.08 MB/s | 350611 | 262 | 4.8× |
| Easyjson | 653786 | 337.03 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 777780 | 283.30 MB/s | 129747 | 470 | 3.1× |
| LightningDecodeAny | 1011981 | 107.03 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2400836 | 91.78 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12586308 | 643.56 MB/s | 11845073 | 20816 | 8.7× |
| Lightning | 12921132 | 626.88 MB/s | 11845079 | 20816 | 8.4× |
| LightningArena | 12959330 | 625.04 MB/s | 11845072 | 20816 | 8.4× |
| SonicFastest | 18346376 | 441.51 MB/s | 19863105 | 41640 | 6.0× |
| Sonic | 18374280 | 440.84 MB/s | 19862497 | 41640 | 5.9× |
| Goccy | 27278552 | 296.94 MB/s | 19043885 | 107156 | 4.0× |
| Easyjson | 35302316 | 229.45 MB/s | 15059617 | 41643 | 3.1× |
| LightningDecodeAny | 40457340 | 128.61 MB/s | 46279351 | 747112 | 2.7× |
| JSONV2 | 52116745 | 155.42 MB/s | 15233724 | 78972 | 2.1× |
| Stdlib | 109162353 | 74.20 MB/s | 15665071 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 5621992 | 530.68 MB/s | 3764713 | 1504 | 10.2× |
| LightningDestructive | 5991024 | 497.99 MB/s | 3758857 | 29356 | 9.5× |
| Lightning | 6181432 | 482.65 MB/s | 3758857 | 29356 | 9.2× |
| SonicFastest | 9693696 | 307.77 MB/s | 9132230 | 57804 | 5.9× |
| Sonic | 9719062 | 306.97 MB/s | 9131997 | 57804 | 5.9× |
| Goccy | 19243433 | 155.04 MB/s | 9901683 | 273621 | 3.0× |
| LightningDecodeAny | 19520452 | 93.96 MB/s | 23982577 | 351152 | 2.9× |
| Easyjson | 19811111 | 150.60 MB/s | 9479440 | 30115 | 2.9× |
| JSONV2 | 27617486 | 108.03 MB/s | 9257070 | 86278 | 2.1× |
| Stdlib | 57135334 | 52.22 MB/s | 9258084 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1339548 | 540.18 MB/s | 907600 | 3618 | 10.3× |
| LightningArena | 1396261 | 518.24 MB/s | 911393 | 30 | 9.8× |
| Lightning | 1437662 | 503.31 MB/s | 907595 | 3618 | 9.6× |
| Sonic | 2085183 | 347.02 MB/s | 2369988 | 3683 | 6.6× |
| SonicFastest | 2090101 | 346.20 MB/s | 2370473 | 3683 | 6.6× |
| Easyjson | 5322040 | 135.96 MB/s | 2847906 | 3698 | 2.6× |
| Goccy | 5437209 | 133.08 MB/s | 2685688 | 80266 | 2.5× |
| LightningDecodeAny | 5462123 | 119.11 MB/s | 6500460 | 76546 | 2.5× |
| JSONV2 | 6437302 | 112.41 MB/s | 2704705 | 7318 | 2.1× |
| Stdlib | 13733901 | 52.69 MB/s | 2704549 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1890308 | 834.44 MB/s | 911393 | 30 | 10.2× |
| Lightning | 1959840 | 804.84 MB/s | 907595 | 3618 | 9.8× |
| LightningDestructive | 1971685 | 800.00 MB/s | 907600 | 3618 | 9.8× |
| Sonic | 2465712 | 639.72 MB/s | 3229484 | 3683 | 7.8× |
| SonicFastest | 2496011 | 631.95 MB/s | 3228432 | 3683 | 7.7× |
| LightningDecodeAny | 4746094 | 158.74 MB/s | 6500454 | 76546 | 4.1× |
| Easyjson | 6530841 | 241.52 MB/s | 2847904 | 3698 | 3.0× |
| Goccy | 6600641 | 238.97 MB/s | 3494624 | 80263 | 2.9× |
| JSONV2 | 7065150 | 223.26 MB/s | 2704555 | 7318 | 2.7× |
| Stdlib | 19285850 | 81.79 MB/s | 2704547 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 230043 | 652.59 MB/s | 81920 | 1 | 9.3× |
| LightningArena | 230117 | 652.38 MB/s | 81920 | 1 | 9.3× |
| LightningDestructive | 238808 | 628.64 MB/s | 81920 | 1 | 9.0× |
| Sonic | 378753 | 396.36 MB/s | 407514 | 16 | 5.7× |
| SonicFastest | 429400 | 349.61 MB/s | 409293 | 16 | 5.0× |
| LightningDecodeAny | 578149 | 259.66 MB/s | 745766 | 10016 | 3.7× |
| Goccy | 1054629 | 142.35 MB/s | 332842 | 10005 | 2.0× |
| JSONV2 | 1187048 | 126.47 MB/s | 357724 | 20 | 1.8× |
| Stdlib | 2144236 | 70.01 MB/s | 357801 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 33371 | 842.55 MB/s | 29216 | 103 | 10.5× |
| LightningArena | 33483 | 839.74 MB/s | 29216 | 103 | 10.4× |
| LightningDestructive | 34096 | 824.64 MB/s | 29088 | 101 | 10.2× |
| Sonic | 56667 | 496.18 MB/s | 59450 | 83 | 6.2× |
| SonicFastest | 56755 | 495.41 MB/s | 59452 | 83 | 6.2× |
| Goccy | 80052 | 351.24 MB/s | 59266 | 188 | 4.4× |
| Easyjson | 80872 | 347.67 MB/s | 32304 | 138 | 4.3× |
| JSONV2 | 138247 | 203.38 MB/s | 36897 | 242 | 2.5× |
| LightningDecodeAny | 165644 | 169.74 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 349138 | 80.53 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1971 | 1180.88 MB/s | 32 | 1 | 13.6× |
| Lightning | 1993 | 1167.82 MB/s | 32 | 1 | 13.5× |
| LightningDestructive | 2126 | 1094.76 MB/s | 32 | 1 | 12.6× |
| Sonic | 4750 | 490.06 MB/s | 3712 | 4 | 5.6× |
| SonicFastest | 4773 | 487.72 MB/s | 3714 | 4 | 5.6× |
| Goccy | 4906 | 474.56 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 5061 | 459.95 MB/s | 192 | 2 | 5.3× |
| JSONV2 | 8623 | 269.98 MB/s | 1000 | 6 | 3.1× |
| LightningDecodeAny | 10545 | 159.79 MB/s | 10200 | 195 | 2.5× |
| Stdlib | 26837 | 86.75 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 203 | 932.34 MB/s | 0 | 0 | 13.9× |
| LightningArena | 203 | 931.27 MB/s | 0 | 0 | 13.8× |
| LightningDestructive | 210 | 901.81 MB/s | 0 | 0 | 13.4× |
| Goccy | 444 | 425.27 MB/s | 304 | 2 | 6.3× |
| Easyjson | 551 | 343.18 MB/s | 0 | 0 | 5.1× |
| SonicFastest | 634 | 297.97 MB/s | 341 | 3 | 4.4× |
| Sonic | 637 | 296.64 MB/s | 341 | 3 | 4.4× |
| JSONV2 | 1038 | 182.09 MB/s | 112 | 1 | 2.7× |
| LightningDecodeAny | 1371 | 97.76 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2809 | 67.28 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1448 | 1512.71 MB/s | 0 | 0 | 13.2× |
| LightningArena | 1473 | 1487.72 MB/s | 0 | 0 | 13.0× |
| LightningDestructive | 1524 | 1438.02 MB/s | 0 | 0 | 12.5× |
| Goccy | 3693 | 593.35 MB/s | 2864 | 4 | 5.2× |
| Easyjson | 3831 | 571.89 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6581 | 332.92 MB/s | 3604 | 38 | 2.9× |
| Sonic | 6744 | 324.90 MB/s | 3603 | 38 | 2.8× |
| JSONV2 | 8532 | 256.81 MB/s | 640 | 6 | 2.2× |
| LightningDecodeAny | 9114 | 198.70 MB/s | 7552 | 158 | 2.1× |
| Stdlib | 19099 | 114.72 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 594491 | 858.68 MB/s | 457537 | 1009 | 12.1× |
| LightningArena | 612881 | 832.91 MB/s | 457537 | 1009 | 11.7× |
| Lightning | 616164 | 828.47 MB/s | 457537 | 1009 | 11.6× |
| Goccy | 1229317 | 415.25 MB/s | 1137032 | 5006 | 5.8× |
| Sonic | 1258002 | 405.78 MB/s | 1306647 | 2014 | 5.7× |
| SonicFastest | 1263671 | 403.96 MB/s | 1306433 | 2014 | 5.7× |
| Easyjson | 1786632 | 285.72 MB/s | 863779 | 3012 | 4.0× |
| LightningDecodeAny | 3549927 | 129.99 MB/s | 2950652 | 64018 | 2.0× |
| JSONV2 | 3597869 | 141.88 MB/s | 1075968 | 12645 | 2.0× |
| Stdlib | 7166196 | 71.23 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 546 | 36279.24 MB/s | 0 | 0 | 303.4× |
| LightningArena | 546 | 36259.20 MB/s | 0 | 0 | 303.2× |
| LightningDestructive | 822 | 24075.59 MB/s | 0 | 0 | 201.3× |
| SonicFastest | 6494 | 3047.37 MB/s | 21143 | 3 | 25.5× |
| Goccy | 23679 | 835.73 MB/s | 20492 | 2 | 7.0× |
| Sonic | 32028 | 617.87 MB/s | 20617 | 3 | 5.2× |
| JSONV2 | 34771 | 569.12 MB/s | 8 | 1 | 4.8× |
| LightningDecodeAny | 94957 | 208.39 MB/s | 116864 | 2015 | 1.7× |
| Easyjson | 104311 | 189.71 MB/s | 0 | 0 | 1.6× |
| Stdlib | 165478 | 119.59 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2215 | 8182.64 MB/s | 432 | 2 | 56.8× |
| LightningArena | 2245 | 8072.80 MB/s | 432 | 2 | 56.1× |
| LightningDestructive | 2427 | 7467.12 MB/s | 0 | 0 | 51.9× |
| Easyjson | 4763 | 3805.21 MB/s | 432 | 2 | 26.4× |
| Sonic | 8525 | 2125.91 MB/s | 20448 | 5 | 14.8× |
| SonicFastest | 8951 | 2024.85 MB/s | 20402 | 5 | 14.1× |
| LightningDecodeAny | 19019 | 940.22 MB/s | 29088 | 191 | 6.6× |
| Goccy | 24811 | 730.48 MB/s | 19460 | 2 | 5.1× |
| JSONV2 | 50864 | 356.32 MB/s | 16500 | 50 | 2.5× |
| Stdlib | 125904 | 143.95 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2332027 | 861.27 MB/s | 3089565 | 6821 | 9.3× |
| Lightning | 2475713 | 811.28 MB/s | 3091278 | 6827 | 8.7× |
| LightningArena | 2492939 | 805.67 MB/s | 3094370 | 6703 | 8.7× |
| SonicFastest | 4182364 | 480.23 MB/s | 5153658 | 7085 | 5.2× |
| Sonic | 4233642 | 474.41 MB/s | 5154555 | 7085 | 5.1× |
| Goccy | 4676825 | 429.46 MB/s | 5410797 | 15831 | 4.6× |
| Easyjson | 5237927 | 383.45 MB/s | 2981485 | 7439 | 4.1× |
| LightningDecodeAny | 6961283 | 164.09 MB/s | 8503512 | 134008 | 3.1× |
| JSONV2 | 7784813 | 258.00 MB/s | 3173672 | 14562 | 2.8× |
| Stdlib | 21655556 | 92.75 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 849 | 646.79 MB/s | 480 | 1 | 7.9× |
| Lightning | 857 | 640.32 MB/s | 480 | 1 | 7.8× |
| LightningDestructive | 871 | 630.62 MB/s | 480 | 1 | 7.7× |
| LightningDecodeAny | 1822 | 300.84 MB/s | 2021 | 46 | 3.7× |
| Easyjson | 2275 | 241.28 MB/s | 1616 | 5 | 2.9× |
| SonicFastest | 2374 | 231.25 MB/s | 2262 | 8 | 2.8× |
| Sonic | 2409 | 227.90 MB/s | 2261 | 8 | 2.8× |
| Goccy | 3384 | 162.26 MB/s | 2128 | 43 | 2.0× |
| JSONV2 | 3602 | 152.40 MB/s | 1664 | 7 | 1.9× |
| Stdlib | 6708 | 81.84 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 526843 | 1198.68 MB/s | 402729 | 545 | 12.2× |
| LightningArena | 602764 | 1047.70 MB/s | 453017 | 712 | 10.6× |
| Lightning | 615637 | 1025.79 MB/s | 451257 | 857 | 10.4× |
| Sonic | 993534 | 635.62 MB/s | 1066239 | 814 | 6.4× |
| SonicFastest | 1017905 | 620.41 MB/s | 1066639 | 814 | 6.3× |
| Easyjson | 1289051 | 489.91 MB/s | 422504 | 936 | 5.0× |
| Goccy | 1451502 | 435.08 MB/s | 989268 | 1201 | 4.4× |
| JSONV2 | 2468387 | 255.84 MB/s | 571586 | 3144 | 2.6× |
| LightningDecodeAny | 2568759 | 181.76 MB/s | 2076503 | 30126 | 2.5× |
| Stdlib | 6402597 | 98.63 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 659027 | 853.39 MB/s | 546569 | 429 | 9.2× |
| Lightning | 853475 | 658.96 MB/s | 769937 | 1235 | 7.1× |
| LightningArena | 873281 | 644.02 MB/s | 771665 | 1088 | 7.0× |
| SonicFastest | 1326332 | 424.03 MB/s | 1348889 | 1185 | 4.6× |
| Sonic | 1421493 | 395.65 MB/s | 1349709 | 1185 | 4.3× |
| Goccy | 1644912 | 341.91 MB/s | 1040700 | 1028 | 3.7× |
| Easyjson | 2137112 | 263.16 MB/s | 775153 | 1254 | 2.9× |
| LightningDecodeAny | 2909805 | 193.28 MB/s | 2180440 | 30126 | 2.1× |
| JSONV2 | 3381328 | 166.33 MB/s | 927408 | 3482 | 1.8× |
| Stdlib | 6093588 | 92.30 MB/s | 1011668 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 625406 | 852.53 MB/s | 333416 | 2084 | 10.5× |
| LightningArena | 701225 | 760.35 MB/s | 368224 | 2293 | 9.4× |
| Lightning | 717884 | 742.71 MB/s | 368224 | 2293 | 9.2× |
| Sonic | 1128967 | 472.27 MB/s | 981311 | 3082 | 5.8× |
| SonicFastest | 1142263 | 466.77 MB/s | 980916 | 3082 | 5.8× |
| Easyjson | 1346415 | 396.00 MB/s | 428362 | 3273 | 4.9× |
| Goccy | 1451595 | 367.30 MB/s | 1167081 | 5409 | 4.5× |
| JSONV2 | 2943180 | 181.16 MB/s | 745421 | 13288 | 2.2× |
| LightningDecodeAny | 3583996 | 148.77 MB/s | 2992874 | 50076 | 1.8× |
| Stdlib | 6596248 | 80.83 MB/s | 798692 | 17133 | 1.0× |
