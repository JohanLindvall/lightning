# JSON Deserialization Benchmarks

- generated 2026-09-03T00:51:53Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 84609 | 1504.26 MB/s | 49760 | 3 | 12.9× |
| Lightning | 84667 | 1503.25 MB/s | 49760 | 3 | 12.9× |
| LightningDestructive | 84717 | 1502.35 MB/s | 49280 | 2 | 12.9× |
| SonicFastest | 186921 | 680.90 MB/s | 197162 | 10 | 5.9× |
| Sonic | 187644 | 678.28 MB/s | 199085 | 10 | 5.8× |
| Goccy | 201737 | 630.90 MB/s | 225136 | 884 | 5.4× |
| Easyjson | 213502 | 596.13 MB/s | 122864 | 14 | 5.1× |
| JSONV2 | 423571 | 300.48 MB/s | 195120 | 1805 | 2.6× |
| LightningDecodeAny | 451786 | 209.51 MB/s | 463410 | 9708 | 2.4× |
| Stdlib | 1094204 | 116.32 MB/s | 199697 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2681666 | 839.42 MB/s | 2532848 | 1143 | 9.9× |
| LightningArena | 2699858 | 833.77 MB/s | 2532848 | 1143 | 9.8× |
| Lightning | 2715388 | 829.00 MB/s | 2532850 | 1143 | 9.8× |
| SonicFastest | 4768501 | 472.07 MB/s | 15233776 | 970 | 5.6× |
| Sonic | 4847333 | 464.39 MB/s | 15233851 | 970 | 5.5× |
| Goccy | 10442398 | 215.57 MB/s | 4125488 | 56532 | 2.5× |
| Easyjson | 11052854 | 203.66 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11722707 | 192.02 MB/s | 19380210 | 223896 | 2.3× |
| JSONV2 | 16402067 | 137.24 MB/s | 3123214 | 3083 | 1.6× |
| Stdlib | 26586872 | 84.67 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 371663 | 727.55 MB/s | 397296 | 567 | 9.2× |
| LightningArena | 371976 | 726.94 MB/s | 397297 | 567 | 9.2× |
| Lightning | 372029 | 726.83 MB/s | 397296 | 567 | 9.2× |
| Sonic | 634748 | 426.00 MB/s | 471879 | 968 | 5.4× |
| SonicFastest | 651042 | 415.34 MB/s | 499093 | 968 | 5.3× |
| Easyjson | 1381498 | 195.73 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1419381 | 190.51 MB/s | 543466 | 8122 | 2.4× |
| LightningDecodeAny | 1590209 | 170.04 MB/s | 2543876 | 29687 | 2.2× |
| JSONV2 | 2116426 | 127.76 MB/s | 348156 | 1628 | 1.6× |
| Stdlib | 3432454 | 78.78 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 959710 | 1799.72 MB/s | 765560 | 2798 | 13.8× |
| LightningArena | 965147 | 1789.58 MB/s | 768416 | 2440 | 13.7× |
| Lightning | 966555 | 1786.97 MB/s | 765601 | 2799 | 13.7× |
| Sonic | 2090727 | 826.13 MB/s | 2669328 | 4020 | 6.3× |
| SonicFastest | 2092815 | 825.30 MB/s | 2696780 | 4020 | 6.3× |
| Goccy | 2422136 | 713.09 MB/s | 2582327 | 14604 | 5.4× |
| JSONV2 | 4205767 | 410.68 MB/s | 1011633 | 7594 | 3.1× |
| Easyjson | 4222344 | 409.06 MB/s | 972032 | 5389 | 3.1× |
| LightningDecodeAny | 4415020 | 113.32 MB/s | 4953694 | 76576 | 3.0× |
| Stdlib | 13199914 | 130.85 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 866 | 2091.34 MB/s | 0 | 0 | 16.0× |
| LightningArena | 872 | 2077.30 MB/s | 0 | 0 | 15.9× |
| LightningDestructive | 885 | 2048.41 MB/s | 0 | 0 | 15.7× |
| Easyjson | 2532 | 715.58 MB/s | 24 | 1 | 5.5× |
| Goccy | 2851 | 635.65 MB/s | 2608 | 4 | 4.9× |
| Sonic | 6000 | 301.99 MB/s | 3708 | 40 | 2.3× |
| SonicFastest | 6011 | 301.47 MB/s | 3733 | 40 | 2.3× |
| JSONV2 | 7750 | 233.79 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7943 | 228.01 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13889 | 130.46 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 896 | 2022.91 MB/s | 0 | 0 | 15.6× |
| Lightning | 896 | 2021.09 MB/s | 0 | 0 | 15.5× |
| LightningDestructive | 927 | 1954.76 MB/s | 0 | 0 | 15.0× |
| Easyjson | 2533 | 715.43 MB/s | 24 | 1 | 5.5× |
| Goccy | 2833 | 639.56 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 5996 | 302.21 MB/s | 3767 | 40 | 2.3× |
| Sonic | 5998 | 302.09 MB/s | 3787 | 40 | 2.3× |
| JSONV2 | 7841 | 231.08 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7910 | 228.94 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13938 | 130.00 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1073 | 1688.07 MB/s | 144 | 10 | 13.0× |
| LightningArena | 1086 | 1668.03 MB/s | 144 | 10 | 12.8× |
| LightningDestructive | 1132 | 1600.43 MB/s | 144 | 10 | 12.3× |
| Easyjson | 2743 | 660.53 MB/s | 144 | 10 | 5.1× |
| Goccy | 2964 | 611.31 MB/s | 2600 | 5 | 4.7× |
| SonicFastest | 6230 | 290.87 MB/s | 3856 | 42 | 2.2× |
| Sonic | 6237 | 290.53 MB/s | 3923 | 42 | 2.2× |
| JSONV2 | 8032 | 225.60 MB/s | 632 | 7 | 1.7× |
| LightningDecodeAny | 8268 | 219.04 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13928 | 130.10 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 621 | 794.92 MB/s | 160 | 1 | 8.8× |
| LightningDestructive | 623 | 792.68 MB/s | 160 | 1 | 8.8× |
| Sonic | 1246 | 396.32 MB/s | 981 | 6 | 4.4× |
| SonicFastest | 1250 | 395.08 MB/s | 986 | 6 | 4.4× |
| LightningArena | 1289 | 383.25 MB/s | 4096 | 1 | 4.2× |
| LightningDecodeAny | 1299 | 379.53 MB/s | 1296 | 26 | 4.2× |
| Easyjson | 2244 | 220.10 MB/s | 448 | 3 | 2.4× |
| Goccy | 2428 | 203.50 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3245 | 152.25 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5468 | 90.35 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 372 | 618.18 MB/s | 160 | 1 | 10.9× |
| Lightning | 376 | 611.16 MB/s | 160 | 1 | 10.8× |
| Sonic | 892 | 257.93 MB/s | 660 | 6 | 4.6× |
| SonicFastest | 898 | 256.16 MB/s | 658 | 6 | 4.5× |
| LightningArena | 1038 | 221.55 MB/s | 4096 | 1 | 3.9× |
| LightningDecodeAny | 1135 | 201.77 MB/s | 1296 | 26 | 3.6× |
| Easyjson | 1384 | 166.16 MB/s | 448 | 3 | 2.9× |
| Goccy | 1587 | 144.96 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2434 | 94.49 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4065 | 56.59 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 52138 | 1249.23 MB/s | 97220 | 98 | 10.5× |
| LightningArena | 52475 | 1241.21 MB/s | 103440 | 103 | 10.5× |
| Lightning | 52735 | 1235.08 MB/s | 103440 | 103 | 10.4× |
| Sonic | 100560 | 647.69 MB/s | 156783 | 75 | 5.5× |
| SonicFastest | 102461 | 635.67 MB/s | 159134 | 75 | 5.4× |
| Goccy | 152188 | 427.97 MB/s | 228809 | 134 | 3.6× |
| LightningDecodeAny | 180762 | 295.02 MB/s | 180048 | 3245 | 3.0× |
| JSONV2 | 234838 | 277.35 MB/s | 206655 | 607 | 2.3× |
| Stdlib | 548627 | 118.72 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2063148 | 940.54 MB/s | 2864592 | 1380 | 11.2× |
| Lightning | 2128782 | 911.54 MB/s | 2864595 | 1380 | 10.8× |
| LightningArena | 2139212 | 907.10 MB/s | 2864594 | 1380 | 10.8× |
| Sonic | 4622059 | 419.83 MB/s | 14606973 | 1407 | 5.0× |
| SonicFastest | 4727311 | 410.48 MB/s | 14608590 | 1407 | 4.9× |
| Goccy | 4802719 | 404.04 MB/s | 4065310 | 13510 | 4.8× |
| Easyjson | 7499767 | 258.74 MB/s | 3871267 | 15043 | 3.1× |
| LightningDecodeAny | 9068485 | 213.98 MB/s | 7063039 | 218633 | 2.5× |
| JSONV2 | 11320773 | 171.41 MB/s | 3237224 | 13947 | 2.0× |
| Stdlib | 23086779 | 84.05 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 891403 | 3733.25 MB/s | 351704 | 1286 | 23.5× |
| Lightning | 1562321 | 2130.06 MB/s | 2488905 | 2995 | 13.4× |
| LightningArena | 1573058 | 2115.52 MB/s | 2488905 | 2995 | 13.3× |
| SonicFastest | 2763866 | 1204.05 MB/s | 6540146 | 4248 | 7.6× |
| Sonic | 2776823 | 1198.43 MB/s | 6519412 | 4248 | 7.6× |
| LightningDecodeAny | 3518221 | 873.67 MB/s | 4876915 | 56892 | 6.0× |
| Goccy | 4683016 | 710.62 MB/s | 3948908 | 3816 | 4.5× |
| JSONV2 | 7650412 | 434.99 MB/s | 5364532 | 13243 | 2.7× |
| Stdlib | 20973120 | 158.67 MB/s | 5565610 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 181089 | 1216.78 MB/s | 135872 | 226 | 11.2× |
| Lightning | 181542 | 1213.74 MB/s | 135872 | 226 | 11.1× |
| LightningDestructive | 182655 | 1206.35 MB/s | 135872 | 226 | 11.1× |
| Sonic | 386309 | 570.39 MB/s | 313402 | 398 | 5.2× |
| SonicFastest | 386825 | 569.63 MB/s | 316393 | 398 | 5.2× |
| Goccy | 445971 | 494.08 MB/s | 365025 | 1067 | 4.5× |
| Easyjson | 548252 | 401.91 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 743700 | 296.28 MB/s | 129741 | 470 | 2.7× |
| LightningDecodeAny | 851421 | 127.21 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2021269 | 109.01 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9330654 | 868.11 MB/s | 11845073 | 20816 | 9.4× |
| Lightning | 9447485 | 857.38 MB/s | 11845073 | 20816 | 9.3× |
| LightningArena | 9466344 | 855.67 MB/s | 11845072 | 20816 | 9.3× |
| Sonic | 17106208 | 473.51 MB/s | 70887480 | 40014 | 5.1× |
| SonicFastest | 17201081 | 470.90 MB/s | 70873031 | 40014 | 5.1× |
| Goccy | 23490859 | 344.82 MB/s | 17025194 | 107148 | 3.7× |
| Easyjson | 30848306 | 262.58 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 35272661 | 147.51 MB/s | 46279353 | 747112 | 2.5× |
| JSONV2 | 44042788 | 183.91 MB/s | 15233761 | 78972 | 2.0× |
| Stdlib | 88032705 | 92.01 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4277503 | 697.48 MB/s | 3764712 | 1504 | 10.9× |
| LightningDestructive | 4563839 | 653.72 MB/s | 3758856 | 29356 | 10.2× |
| Lightning | 4708601 | 633.62 MB/s | 3758859 | 29356 | 9.9× |
| SonicFastest | 8678401 | 343.78 MB/s | 26465280 | 56760 | 5.4× |
| Sonic | 8686039 | 343.48 MB/s | 26451229 | 56760 | 5.4× |
| LightningDecodeAny | 15816834 | 115.96 MB/s | 23982581 | 351152 | 2.9× |
| Easyjson | 16526595 | 180.53 MB/s | 9479440 | 30115 | 2.8× |
| Goccy | 16675993 | 178.91 MB/s | 10596903 | 273649 | 2.8× |
| JSONV2 | 25113710 | 118.80 MB/s | 9257171 | 86278 | 1.9× |
| Stdlib | 46597718 | 64.03 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 915446 | 790.43 MB/s | 907601 | 3618 | 12.5× |
| LightningArena | 928617 | 779.22 MB/s | 911396 | 30 | 12.3× |
| Lightning | 981125 | 737.52 MB/s | 907598 | 3618 | 11.6× |
| SonicFastest | 1822932 | 396.94 MB/s | 3183859 | 7226 | 6.3× |
| Sonic | 1823086 | 396.91 MB/s | 3187335 | 7226 | 6.3× |
| LightningDecodeAny | 4091497 | 159.01 MB/s | 6500459 | 76546 | 2.8× |
| Easyjson | 4214581 | 171.69 MB/s | 2847904 | 3698 | 2.7× |
| Goccy | 4898097 | 147.73 MB/s | 2960372 | 80282 | 2.3× |
| JSONV2 | 5843907 | 123.82 MB/s | 2704617 | 7318 | 2.0× |
| Stdlib | 11403793 | 63.45 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1365929 | 1154.78 MB/s | 907600 | 3618 | 11.4× |
| LightningArena | 1373285 | 1148.60 MB/s | 911392 | 30 | 11.3× |
| Lightning | 1420174 | 1110.68 MB/s | 907595 | 3618 | 10.9× |
| SonicFastest | 2275339 | 693.24 MB/s | 5787517 | 7226 | 6.8× |
| Sonic | 2291860 | 688.24 MB/s | 5790969 | 7226 | 6.8× |
| LightningDecodeAny | 3671714 | 205.19 MB/s | 6500455 | 76546 | 4.2× |
| Easyjson | 5614529 | 280.94 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5678395 | 277.78 MB/s | 3568441 | 80266 | 2.7× |
| JSONV2 | 6428896 | 245.35 MB/s | 2704592 | 7318 | 2.4× |
| Stdlib | 15535075 | 101.53 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 157018 | 956.09 MB/s | 81920 | 1 | 11.8× |
| LightningArena | 157057 | 955.86 MB/s | 81920 | 1 | 11.8× |
| Lightning | 157096 | 955.62 MB/s | 81920 | 1 | 11.8× |
| Sonic | 273842 | 548.21 MB/s | 253585 | 6 | 6.8× |
| SonicFastest | 275381 | 545.15 MB/s | 256869 | 6 | 6.8× |
| LightningDecodeAny | 431048 | 348.27 MB/s | 745765 | 10016 | 4.3× |
| Goccy | 875663 | 171.44 MB/s | 324945 | 10004 | 2.1× |
| JSONV2 | 1068423 | 140.51 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1859522 | 80.73 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 27894 | 1007.99 MB/s | 29216 | 103 | 10.8× |
| Lightning | 27981 | 1004.85 MB/s | 29216 | 103 | 10.8× |
| LightningDestructive | 28322 | 992.75 MB/s | 29088 | 101 | 10.6× |
| Sonic | 64473 | 436.11 MB/s | 48756 | 103 | 4.7× |
| SonicFastest | 64533 | 435.70 MB/s | 48697 | 103 | 4.7× |
| Easyjson | 68833 | 408.48 MB/s | 32304 | 138 | 4.4× |
| Goccy | 73134 | 384.46 MB/s | 59249 | 188 | 4.1× |
| JSONV2 | 135376 | 207.70 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 152435 | 184.45 MB/s | 140576 | 2643 | 2.0× |
| Stdlib | 301060 | 93.39 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1510 | 1541.92 MB/s | 32 | 1 | 15.0× |
| LightningArena | 1512 | 1540.05 MB/s | 32 | 1 | 14.9× |
| LightningDestructive | 1587 | 1466.73 MB/s | 32 | 1 | 14.2× |
| Easyjson | 4230 | 550.34 MB/s | 192 | 2 | 5.3× |
| Goccy | 4303 | 541.07 MB/s | 3649 | 4 | 5.2× |
| SonicFastest | 5161 | 451.10 MB/s | 4473 | 6 | 4.4× |
| Sonic | 5226 | 445.50 MB/s | 4525 | 6 | 4.3× |
| JSONV2 | 8599 | 270.74 MB/s | 1000 | 6 | 2.6× |
| LightningDecodeAny | 10397 | 162.07 MB/s | 10200 | 195 | 2.2× |
| Stdlib | 22586 | 103.07 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 178 | 1060.32 MB/s | 0 | 0 | 13.4× |
| Lightning | 178 | 1059.25 MB/s | 0 | 0 | 13.3× |
| LightningDestructive | 181 | 1044.93 MB/s | 0 | 0 | 13.2× |
| Goccy | 399 | 474.01 MB/s | 304 | 2 | 6.0× |
| Easyjson | 490 | 385.85 MB/s | 0 | 0 | 4.9× |
| SonicFastest | 790 | 239.15 MB/s | 521 | 4 | 3.0× |
| Sonic | 792 | 238.79 MB/s | 516 | 4 | 3.0× |
| JSONV2 | 1028 | 183.81 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1255 | 106.74 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2380 | 79.42 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1174 | 1866.67 MB/s | 0 | 0 | 13.5× |
| Lightning | 1175 | 1865.30 MB/s | 0 | 0 | 13.4× |
| LightningDestructive | 1198 | 1828.85 MB/s | 0 | 0 | 13.2× |
| Goccy | 3152 | 695.17 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3192 | 686.44 MB/s | 24 | 1 | 5.0× |
| Sonic | 6383 | 343.26 MB/s | 3992 | 40 | 2.5× |
| SonicFastest | 6410 | 341.79 MB/s | 4025 | 40 | 2.5× |
| LightningDecodeAny | 7908 | 229.02 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 8091 | 270.80 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15802 | 138.65 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 553753 | 921.85 MB/s | 457536 | 1009 | 10.8× |
| Lightning | 559921 | 911.69 MB/s | 457537 | 1009 | 10.6× |
| LightningArena | 564543 | 904.23 MB/s | 457536 | 1009 | 10.6× |
| Sonic | 1167182 | 437.36 MB/s | 900420 | 2006 | 5.1× |
| SonicFastest | 1168247 | 436.96 MB/s | 904409 | 2006 | 5.1× |
| Goccy | 1180984 | 432.25 MB/s | 1143540 | 5007 | 5.0× |
| Easyjson | 1549010 | 329.55 MB/s | 863777 | 3012 | 3.8× |
| JSONV2 | 3259283 | 156.62 MB/s | 1076015 | 12646 | 1.8× |
| LightningDecodeAny | 3362272 | 137.25 MB/s | 2950647 | 64018 | 1.8× |
| Stdlib | 5959896 | 85.65 MB/s | 1162117 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 481 | 41107.30 MB/s | 0 | 0 | 224.6× |
| Lightning | 482 | 41081.45 MB/s | 0 | 0 | 224.5× |
| LightningDestructive | 496 | 39878.07 MB/s | 0 | 0 | 217.9× |
| Goccy | 20552 | 962.89 MB/s | 20491 | 2 | 5.3× |
| SonicFastest | 27339 | 723.84 MB/s | 22561 | 4 | 4.0× |
| Sonic | 27439 | 721.20 MB/s | 22514 | 4 | 3.9× |
| JSONV2 | 29748 | 665.23 MB/s | 8 | 1 | 3.6× |
| Easyjson | 82111 | 241.00 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 83161 | 237.95 MB/s | 116864 | 2015 | 1.3× |
| Stdlib | 108137 | 183.00 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1765 | 10270.95 MB/s | 0 | 0 | 58.0× |
| LightningArena | 1872 | 9682.18 MB/s | 432 | 2 | 54.7× |
| Lightning | 1877 | 9654.65 MB/s | 432 | 2 | 54.6× |
| Easyjson | 3954 | 4583.79 MB/s | 432 | 2 | 25.9× |
| SonicFastest | 10139 | 1787.54 MB/s | 23200 | 6 | 10.1× |
| Sonic | 10188 | 1779.02 MB/s | 23221 | 6 | 10.1× |
| LightningDecodeAny | 16143 | 1107.69 MB/s | 29088 | 191 | 6.3× |
| Goccy | 16145 | 1122.56 MB/s | 19459 | 2 | 6.3× |
| JSONV2 | 45406 | 399.16 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102392 | 177.01 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2136459 | 940.10 MB/s | 3089564 | 6821 | 8.7× |
| Lightning | 2204579 | 911.06 MB/s | 3091277 | 6827 | 8.4× |
| LightningArena | 2207542 | 909.83 MB/s | 3094371 | 6703 | 8.4× |
| Goccy | 4330337 | 463.82 MB/s | 5411489 | 15830 | 4.3× |
| Sonic | 4566110 | 439.87 MB/s | 10985809 | 13683 | 4.1× |
| SonicFastest | 4567594 | 439.73 MB/s | 10959498 | 13683 | 4.1× |
| Easyjson | 4962095 | 404.77 MB/s | 2981487 | 7439 | 3.7× |
| JSONV2 | 7055225 | 284.68 MB/s | 3173687 | 14563 | 2.6× |
| LightningDecodeAny | 7309140 | 156.28 MB/s | 8503513 | 134008 | 2.5× |
| Stdlib | 18549002 | 108.28 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 854 | 642.46 MB/s | 480 | 1 | 6.6× |
| LightningArena | 859 | 639.21 MB/s | 480 | 1 | 6.6× |
| LightningDestructive | 862 | 637.07 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1637 | 334.73 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 2168 | 253.22 MB/s | 1616 | 5 | 2.6× |
| Sonic | 2627 | 208.95 MB/s | 1942 | 26 | 2.1× |
| SonicFastest | 2636 | 208.27 MB/s | 1937 | 26 | 2.1× |
| Goccy | 2969 | 184.91 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3322 | 165.27 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5633 | 97.47 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 411908 | 1533.14 MB/s | 402729 | 545 | 13.1× |
| Lightning | 485023 | 1302.03 MB/s | 451257 | 857 | 11.1× |
| LightningArena | 486647 | 1297.68 MB/s | 453017 | 712 | 11.0× |
| SonicFastest | 1035315 | 609.97 MB/s | 1022273 | 1102 | 5.2× |
| Sonic | 1035336 | 609.96 MB/s | 1018706 | 1102 | 5.2× |
| Easyjson | 1141746 | 553.11 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1183133 | 533.76 MB/s | 987028 | 1201 | 4.5× |
| JSONV2 | 2153292 | 293.28 MB/s | 571613 | 3144 | 2.5× |
| LightningDecodeAny | 2387243 | 195.58 MB/s | 2076505 | 30126 | 2.3× |
| Stdlib | 5375702 | 117.48 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 561740 | 1001.19 MB/s | 546569 | 429 | 9.4× |
| Lightning | 726552 | 774.08 MB/s | 769938 | 1235 | 7.2× |
| LightningArena | 729009 | 771.47 MB/s | 771666 | 1088 | 7.2× |
| Sonic | 1038326 | 541.65 MB/s | 954548 | 1476 | 5.1× |
| SonicFastest | 1047018 | 537.15 MB/s | 968820 | 1476 | 5.0× |
| Goccy | 1359928 | 413.56 MB/s | 1038782 | 1029 | 3.9× |
| Easyjson | 1765777 | 318.50 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2578375 | 218.12 MB/s | 2180441 | 30126 | 2.0× |
| JSONV2 | 2773363 | 202.79 MB/s | 927439 | 3482 | 1.9× |
| Stdlib | 5263232 | 106.86 MB/s | 1011674 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 582532 | 915.28 MB/s | 333416 | 2084 | 9.3× |
| Lightning | 602832 | 884.46 MB/s | 368224 | 2293 | 9.0× |
| LightningArena | 610248 | 873.71 MB/s | 368224 | 2293 | 8.9× |
| Easyjson | 1116741 | 477.44 MB/s | 428361 | 3273 | 4.8× |
| Sonic | 1157433 | 460.66 MB/s | 1025858 | 4351 | 4.7× |
| SonicFastest | 1158580 | 460.20 MB/s | 1035722 | 4351 | 4.7× |
| Goccy | 1315323 | 405.36 MB/s | 1167213 | 5409 | 4.1× |
| JSONV2 | 2549439 | 209.14 MB/s | 745451 | 13288 | 2.1× |
| LightningDecodeAny | 3406902 | 156.50 MB/s | 2992878 | 50076 | 1.6× |
| Stdlib | 5406490 | 98.62 MB/s | 798692 | 17133 | 1.0× |
