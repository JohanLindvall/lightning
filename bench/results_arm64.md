# JSON Deserialization Benchmarks

- generated 2026-09-08T04:43:19Z
- go version go1.26.7 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 84603 | 1504.38 MB/s | 49760 | 3 | 12.9× |
| LightningArena | 84661 | 1503.34 MB/s | 49760 | 3 | 12.9× |
| LightningDestructive | 85112 | 1495.39 MB/s | 49280 | 2 | 12.9× |
| Sonic | 189037 | 673.28 MB/s | 209114 | 10 | 5.8× |
| SonicFastest | 189354 | 672.15 MB/s | 210246 | 10 | 5.8× |
| Goccy | 203314 | 626.00 MB/s | 225632 | 884 | 5.4× |
| Easyjson | 212982 | 597.59 MB/s | 122864 | 14 | 5.1× |
| JSONV2 | 427470 | 297.74 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 436688 | 216.75 MB/s | 463409 | 9708 | 2.5× |
| Stdlib | 1094815 | 116.25 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2671908 | 842.49 MB/s | 2532849 | 1143 | 9.9× |
| LightningArena | 2697555 | 834.48 MB/s | 2532850 | 1143 | 9.8× |
| Lightning | 2712669 | 829.83 MB/s | 2532852 | 1143 | 9.8× |
| Sonic | 4749402 | 473.97 MB/s | 15233873 | 970 | 5.6× |
| SonicFastest | 4792061 | 469.75 MB/s | 15232102 | 970 | 5.5× |
| Goccy | 10289770 | 218.77 MB/s | 4116510 | 56532 | 2.6× |
| Easyjson | 10989506 | 204.84 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11680067 | 192.72 MB/s | 19380212 | 223896 | 2.3× |
| JSONV2 | 16306846 | 138.04 MB/s | 3123216 | 3083 | 1.6× |
| Stdlib | 26569982 | 84.72 MB/s | 3123399 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 371290 | 728.28 MB/s | 397296 | 567 | 9.2× |
| LightningArena | 373480 | 724.01 MB/s | 397296 | 567 | 9.2× |
| Lightning | 374913 | 721.24 MB/s | 397297 | 567 | 9.2× |
| SonicFastest | 641252 | 421.68 MB/s | 484012 | 968 | 5.4× |
| Sonic | 647327 | 417.72 MB/s | 492401 | 968 | 5.3× |
| Easyjson | 1379891 | 195.96 MB/s | 330272 | 749 | 2.5× |
| Goccy | 1413710 | 191.27 MB/s | 543694 | 8122 | 2.4× |
| LightningDecodeAny | 1586897 | 170.40 MB/s | 2543877 | 29687 | 2.2× |
| JSONV2 | 2125822 | 127.20 MB/s | 348153 | 1628 | 1.6× |
| Stdlib | 3431778 | 78.79 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 950828 | 1816.53 MB/s | 765560 | 2798 | 13.9× |
| Lightning | 953787 | 1810.89 MB/s | 765602 | 2799 | 13.9× |
| LightningArena | 963478 | 1792.68 MB/s | 772704 | 2445 | 13.7× |
| SonicFastest | 2066289 | 835.90 MB/s | 2749315 | 4020 | 6.4× |
| Sonic | 2068963 | 834.82 MB/s | 2739600 | 4020 | 6.4× |
| Goccy | 2397264 | 720.49 MB/s | 2581722 | 14604 | 5.5× |
| Easyjson | 4231113 | 408.22 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4255966 | 405.83 MB/s | 1011633 | 7594 | 3.1× |
| LightningDecodeAny | 4380189 | 114.22 MB/s | 4953695 | 76576 | 3.0× |
| Stdlib | 13234895 | 130.50 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 851 | 2129.59 MB/s | 0 | 0 | 16.4× |
| LightningArena | 855 | 2118.25 MB/s | 0 | 0 | 16.3× |
| LightningDestructive | 868 | 2086.38 MB/s | 0 | 0 | 16.0× |
| Easyjson | 2520 | 719.00 MB/s | 24 | 1 | 5.5× |
| Goccy | 2797 | 647.89 MB/s | 2608 | 4 | 5.0× |
| Sonic | 6004 | 301.80 MB/s | 3718 | 40 | 2.3× |
| SonicFastest | 6027 | 300.66 MB/s | 3774 | 40 | 2.3× |
| JSONV2 | 7745 | 233.95 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7948 | 227.85 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13934 | 130.04 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 875 | 2071.71 MB/s | 0 | 0 | 15.9× |
| LightningArena | 876 | 2068.95 MB/s | 0 | 0 | 15.9× |
| LightningDestructive | 907 | 1998.07 MB/s | 0 | 0 | 15.3× |
| Easyjson | 2524 | 717.92 MB/s | 24 | 1 | 5.5× |
| Goccy | 2834 | 639.48 MB/s | 2608 | 4 | 4.9× |
| Sonic | 5948 | 304.64 MB/s | 3747 | 40 | 2.3× |
| SonicFastest | 5972 | 303.42 MB/s | 3799 | 40 | 2.3× |
| JSONV2 | 7798 | 232.35 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7901 | 229.21 MB/s | 7552 | 158 | 1.8× |
| Stdlib | 13919 | 130.18 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1055 | 1718.26 MB/s | 144 | 10 | 13.2× |
| LightningArena | 1066 | 1700.24 MB/s | 144 | 10 | 13.0× |
| LightningDestructive | 1109 | 1634.51 MB/s | 144 | 10 | 12.5× |
| Easyjson | 2744 | 660.44 MB/s | 144 | 10 | 5.1× |
| Goccy | 2913 | 622.14 MB/s | 2600 | 5 | 4.8× |
| Sonic | 6181 | 293.16 MB/s | 3813 | 42 | 2.2× |
| SonicFastest | 6230 | 290.87 MB/s | 3810 | 42 | 2.2× |
| LightningDecodeAny | 7870 | 230.12 MB/s | 7552 | 158 | 1.8× |
| JSONV2 | 7964 | 227.51 MB/s | 632 | 7 | 1.7× |
| Stdlib | 13901 | 130.35 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 621 | 795.19 MB/s | 160 | 1 | 8.8× |
| Lightning | 622 | 793.99 MB/s | 160 | 1 | 8.8× |
| Sonic | 1254 | 393.93 MB/s | 985 | 6 | 4.4× |
| SonicFastest | 1255 | 393.73 MB/s | 976 | 6 | 4.4× |
| LightningDecodeAny | 1299 | 379.57 MB/s | 1296 | 26 | 4.2× |
| LightningArena | 1378 | 358.56 MB/s | 4120 | 2 | 4.0× |
| Easyjson | 2195 | 225.09 MB/s | 448 | 3 | 2.5× |
| Goccy | 2445 | 202.08 MB/s | 856 | 23 | 2.2× |
| JSONV2 | 3238 | 152.54 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5483 | 90.09 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 367 | 627.45 MB/s | 160 | 1 | 11.1× |
| Lightning | 370 | 621.29 MB/s | 160 | 1 | 11.0× |
| SonicFastest | 887 | 259.28 MB/s | 659 | 6 | 4.6× |
| Sonic | 888 | 259.03 MB/s | 649 | 6 | 4.6× |
| LightningArena | 1117 | 205.82 MB/s | 4120 | 2 | 3.7× |
| LightningDecodeAny | 1124 | 203.77 MB/s | 1296 | 26 | 3.6× |
| Easyjson | 1381 | 166.57 MB/s | 448 | 3 | 3.0× |
| Goccy | 1593 | 144.43 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2424 | 94.87 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4082 | 56.34 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 52053 | 1251.27 MB/s | 103440 | 103 | 10.5× |
| LightningDestructive | 52072 | 1250.80 MB/s | 97220 | 98 | 10.5× |
| Lightning | 52123 | 1249.59 MB/s | 103440 | 103 | 10.5× |
| Sonic | 97406 | 668.67 MB/s | 154503 | 75 | 5.6× |
| SonicFastest | 97722 | 666.50 MB/s | 155790 | 75 | 5.6× |
| Goccy | 145316 | 448.21 MB/s | 229419 | 134 | 3.8× |
| LightningDecodeAny | 178960 | 297.99 MB/s | 180048 | 3245 | 3.1× |
| JSONV2 | 226825 | 287.15 MB/s | 206650 | 607 | 2.4× |
| Stdlib | 548836 | 118.67 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2039839 | 951.29 MB/s | 2864592 | 1380 | 11.3× |
| Lightning | 2102244 | 923.05 MB/s | 2864594 | 1380 | 11.0× |
| LightningArena | 2122764 | 914.13 MB/s | 2864593 | 1380 | 10.9× |
| Goccy | 4809497 | 403.47 MB/s | 4064724 | 13510 | 4.8× |
| Sonic | 4908894 | 395.30 MB/s | 14610308 | 1407 | 4.7× |
| SonicFastest | 4997811 | 388.26 MB/s | 14608712 | 1407 | 4.6× |
| Easyjson | 7524255 | 257.90 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9284347 | 209.00 MB/s | 7063039 | 218633 | 2.5× |
| JSONV2 | 11303577 | 171.67 MB/s | 3237231 | 13947 | 2.0× |
| Stdlib | 23058361 | 84.15 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 895165 | 3717.56 MB/s | 351704 | 1286 | 23.5× |
| Lightning | 1543466 | 2156.08 MB/s | 2488905 | 2995 | 13.6× |
| LightningArena | 1563193 | 2128.87 MB/s | 2488905 | 2995 | 13.4× |
| SonicFastest | 2679791 | 1241.82 MB/s | 6437672 | 4248 | 7.8× |
| Sonic | 2707817 | 1228.97 MB/s | 6480161 | 4248 | 7.8× |
| LightningDecodeAny | 3544829 | 867.11 MB/s | 4876911 | 56892 | 5.9× |
| Goccy | 4550079 | 731.38 MB/s | 3948907 | 3816 | 4.6× |
| JSONV2 | 7478682 | 444.98 MB/s | 5364515 | 13243 | 2.8× |
| Stdlib | 20994011 | 158.51 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 180169 | 1223.00 MB/s | 135872 | 226 | 11.2× |
| Lightning | 180930 | 1217.85 MB/s | 135872 | 226 | 11.1× |
| LightningDestructive | 182399 | 1208.05 MB/s | 135872 | 226 | 11.1× |
| SonicFastest | 378265 | 582.52 MB/s | 300898 | 398 | 5.3× |
| Sonic | 379532 | 580.57 MB/s | 303876 | 398 | 5.3× |
| Goccy | 435003 | 506.54 MB/s | 364403 | 1067 | 4.6× |
| Easyjson | 548364 | 401.82 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 740151 | 297.70 MB/s | 129743 | 470 | 2.7× |
| LightningDecodeAny | 841326 | 128.74 MB/s | 897217 | 11703 | 2.4× |
| Stdlib | 2017290 | 109.23 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 9340820 | 867.17 MB/s | 11845073 | 20816 | 9.4× |
| Lightning | 9536351 | 849.39 MB/s | 11845073 | 20816 | 9.2× |
| LightningArena | 9568449 | 846.54 MB/s | 11845077 | 20816 | 9.2× |
| Sonic | 16974853 | 477.18 MB/s | 70916803 | 40014 | 5.2× |
| SonicFastest | 17095438 | 473.81 MB/s | 70916804 | 40014 | 5.2× |
| Goccy | 23323025 | 347.30 MB/s | 17053230 | 107148 | 3.8× |
| Easyjson | 30564173 | 265.02 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 35145419 | 148.04 MB/s | 46279352 | 747112 | 2.5× |
| JSONV2 | 43737387 | 185.20 MB/s | 15233734 | 78972 | 2.0× |
| Stdlib | 88121253 | 91.92 MB/s | 15665066 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4354378 | 685.16 MB/s | 3780457 | 1514 | 10.7× |
| LightningDestructive | 4562897 | 653.85 MB/s | 3758856 | 29356 | 10.2× |
| Lightning | 4682441 | 637.16 MB/s | 3758859 | 29356 | 9.9× |
| SonicFastest | 8706231 | 342.68 MB/s | 26505707 | 56760 | 5.3× |
| Sonic | 8708965 | 342.57 MB/s | 26530440 | 56760 | 5.3× |
| LightningDecodeAny | 15776251 | 116.26 MB/s | 23982580 | 351152 | 3.0× |
| Goccy | 16627607 | 179.43 MB/s | 10634448 | 273649 | 2.8× |
| Easyjson | 16736341 | 178.26 MB/s | 9479440 | 30115 | 2.8× |
| JSONV2 | 24921128 | 119.72 MB/s | 9257149 | 86278 | 1.9× |
| Stdlib | 46556706 | 64.08 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 920669 | 785.95 MB/s | 907601 | 3618 | 12.4× |
| LightningArena | 938084 | 771.36 MB/s | 916259 | 37 | 12.1× |
| Lightning | 975055 | 742.11 MB/s | 907598 | 3618 | 11.7× |
| Sonic | 1777906 | 406.99 MB/s | 3175149 | 7226 | 6.4× |
| SonicFastest | 1778449 | 406.87 MB/s | 3187592 | 7226 | 6.4× |
| LightningDecodeAny | 4039856 | 161.04 MB/s | 6500458 | 76546 | 2.8× |
| Easyjson | 4198450 | 172.35 MB/s | 2847904 | 3698 | 2.7× |
| Goccy | 4827842 | 149.88 MB/s | 2858548 | 80276 | 2.4× |
| JSONV2 | 5428929 | 133.29 MB/s | 2704638 | 7318 | 2.1× |
| Stdlib | 11392532 | 63.52 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1366305 | 1154.47 MB/s | 907601 | 3618 | 11.4× |
| LightningArena | 1382083 | 1141.29 MB/s | 916256 | 37 | 11.3× |
| Lightning | 1415732 | 1114.16 MB/s | 907595 | 3618 | 11.0× |
| Sonic | 2269917 | 694.89 MB/s | 5784558 | 7226 | 6.9× |
| SonicFastest | 2292163 | 688.15 MB/s | 5783896 | 7226 | 6.8× |
| LightningDecodeAny | 3653566 | 206.21 MB/s | 6500460 | 76546 | 4.3× |
| Easyjson | 5590193 | 282.16 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5621668 | 280.58 MB/s | 3568690 | 80267 | 2.8× |
| JSONV2 | 6205053 | 254.20 MB/s | 2704594 | 7318 | 2.5× |
| Stdlib | 15558340 | 101.38 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 157161 | 955.22 MB/s | 81920 | 1 | 11.8× |
| LightningArena | 157245 | 954.71 MB/s | 81920 | 1 | 11.8× |
| Lightning | 157385 | 953.86 MB/s | 81920 | 1 | 11.8× |
| Sonic | 276533 | 542.88 MB/s | 256646 | 6 | 6.7× |
| SonicFastest | 277294 | 541.39 MB/s | 259402 | 6 | 6.7× |
| LightningDecodeAny | 428635 | 350.23 MB/s | 745765 | 10016 | 4.3× |
| Goccy | 864551 | 173.64 MB/s | 324367 | 10004 | 2.1× |
| JSONV2 | 1097548 | 136.78 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1858580 | 80.77 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 27254 | 1031.67 MB/s | 29216 | 103 | 11.0× |
| Lightning | 27335 | 1028.63 MB/s | 29216 | 103 | 11.0× |
| LightningDestructive | 27559 | 1020.24 MB/s | 29088 | 101 | 10.9× |
| Sonic | 63430 | 443.28 MB/s | 46987 | 103 | 4.7× |
| SonicFastest | 63815 | 440.60 MB/s | 47507 | 103 | 4.7× |
| Easyjson | 68531 | 410.28 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72113 | 389.90 MB/s | 59196 | 188 | 4.2× |
| JSONV2 | 135610 | 207.34 MB/s | 36895 | 242 | 2.2× |
| LightningDecodeAny | 146539 | 191.87 MB/s | 140576 | 2643 | 2.1× |
| Stdlib | 300680 | 93.51 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1511 | 1540.35 MB/s | 32 | 1 | 14.9× |
| LightningArena | 1512 | 1539.48 MB/s | 32 | 1 | 14.9× |
| LightningDestructive | 1585 | 1468.65 MB/s | 32 | 1 | 14.2× |
| Goccy | 4131 | 563.53 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4217 | 552.04 MB/s | 192 | 2 | 5.3× |
| Sonic | 5090 | 457.36 MB/s | 4264 | 6 | 4.4× |
| SonicFastest | 5112 | 455.44 MB/s | 4229 | 6 | 4.4× |
| JSONV2 | 8467 | 274.95 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 9868 | 170.75 MB/s | 10200 | 195 | 2.3× |
| Stdlib | 22462 | 103.64 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 179 | 1055.62 MB/s | 0 | 0 | 13.3× |
| Lightning | 179 | 1053.38 MB/s | 0 | 0 | 13.3× |
| LightningDestructive | 182 | 1037.32 MB/s | 0 | 0 | 13.1× |
| Goccy | 382 | 494.77 MB/s | 304 | 2 | 6.2× |
| Easyjson | 493 | 383.44 MB/s | 0 | 0 | 4.8× |
| SonicFastest | 801 | 236.06 MB/s | 505 | 4 | 3.0× |
| Sonic | 801 | 236.03 MB/s | 501 | 4 | 3.0× |
| JSONV2 | 1031 | 183.32 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1205 | 111.21 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2387 | 79.17 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1144 | 1914.99 MB/s | 0 | 0 | 13.8× |
| Lightning | 1149 | 1907.07 MB/s | 0 | 0 | 13.8× |
| LightningDestructive | 1173 | 1867.70 MB/s | 0 | 0 | 13.5× |
| Goccy | 3158 | 693.82 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3193 | 686.29 MB/s | 24 | 1 | 4.9× |
| Sonic | 6363 | 344.36 MB/s | 3952 | 40 | 2.5× |
| SonicFastest | 6395 | 342.64 MB/s | 3996 | 40 | 2.5× |
| LightningDecodeAny | 7903 | 229.16 MB/s | 7552 | 158 | 2.0× |
| JSONV2 | 8047 | 272.28 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15804 | 138.64 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 558256 | 914.41 MB/s | 457537 | 1009 | 10.8× |
| LightningArena | 560097 | 911.41 MB/s | 457537 | 1009 | 10.7× |
| Lightning | 568687 | 897.64 MB/s | 457536 | 1009 | 10.6× |
| Sonic | 1150193 | 443.82 MB/s | 860968 | 2006 | 5.2× |
| SonicFastest | 1158176 | 440.76 MB/s | 869152 | 2006 | 5.2× |
| Goccy | 1180661 | 432.36 MB/s | 1141744 | 5007 | 5.1× |
| Easyjson | 1543986 | 330.62 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3241851 | 157.46 MB/s | 1076012 | 12646 | 1.9× |
| LightningDecodeAny | 3346051 | 137.91 MB/s | 2950648 | 64018 | 1.8× |
| Stdlib | 6003217 | 85.03 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 481 | 41127.45 MB/s | 0 | 0 | 225.2× |
| Lightning | 482 | 41044.25 MB/s | 0 | 0 | 224.8× |
| LightningDestructive | 502 | 39391.80 MB/s | 0 | 0 | 215.7× |
| Goccy | 20785 | 952.10 MB/s | 20491 | 2 | 5.2× |
| Sonic | 27389 | 722.51 MB/s | 22759 | 4 | 4.0× |
| SonicFastest | 27660 | 715.45 MB/s | 23146 | 4 | 3.9× |
| JSONV2 | 29735 | 665.51 MB/s | 8 | 1 | 3.6× |
| Easyjson | 81965 | 241.43 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 82971 | 238.49 MB/s | 116864 | 2015 | 1.3× |
| Stdlib | 108357 | 182.63 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1752 | 10345.82 MB/s | 0 | 0 | 58.6× |
| LightningArena | 1879 | 9643.87 MB/s | 432 | 2 | 54.6× |
| Lightning | 1880 | 9642.61 MB/s | 432 | 2 | 54.6× |
| Easyjson | 3969 | 4566.48 MB/s | 432 | 2 | 25.8× |
| Sonic | 9920 | 1827.01 MB/s | 22804 | 6 | 10.3× |
| SonicFastest | 10122 | 1790.61 MB/s | 23196 | 6 | 10.1× |
| LightningDecodeAny | 16019 | 1116.33 MB/s | 29088 | 191 | 6.4× |
| Goccy | 16105 | 1125.34 MB/s | 19459 | 2 | 6.4× |
| JSONV2 | 45716 | 396.45 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102581 | 176.68 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2150746 | 933.86 MB/s | 3089565 | 6821 | 8.6× |
| Lightning | 2210263 | 908.71 MB/s | 3091279 | 6827 | 8.4× |
| LightningArena | 2214228 | 907.09 MB/s | 3094395 | 6704 | 8.4× |
| Goccy | 4342396 | 462.53 MB/s | 5413677 | 15838 | 4.3× |
| SonicFastest | 4484911 | 447.83 MB/s | 10948103 | 13683 | 4.1× |
| Sonic | 4504728 | 445.86 MB/s | 10936566 | 13683 | 4.1× |
| Easyjson | 4998287 | 401.84 MB/s | 2981518 | 7439 | 3.7× |
| JSONV2 | 7043554 | 285.15 MB/s | 3173696 | 14563 | 2.6× |
| LightningDecodeAny | 7313886 | 156.18 MB/s | 8503513 | 134008 | 2.5× |
| Stdlib | 18545161 | 108.30 MB/s | 3589317 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 871 | 630.18 MB/s | 480 | 1 | 6.5× |
| LightningArena | 875 | 627.69 MB/s | 480 | 1 | 6.5× |
| LightningDestructive | 876 | 626.53 MB/s | 480 | 1 | 6.5× |
| LightningDecodeAny | 1729 | 316.92 MB/s | 2021 | 46 | 3.3× |
| Easyjson | 2238 | 245.36 MB/s | 1616 | 5 | 2.5× |
| Sonic | 2710 | 202.59 MB/s | 2016 | 26 | 2.1× |
| SonicFastest | 2711 | 202.52 MB/s | 2014 | 26 | 2.1× |
| Goccy | 3043 | 180.42 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3376 | 162.60 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5670 | 96.82 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 419255 | 1506.28 MB/s | 402728 | 545 | 12.8× |
| LightningArena | 481796 | 1310.75 MB/s | 453041 | 713 | 11.2× |
| Lightning | 485333 | 1301.20 MB/s | 451257 | 857 | 11.1× |
| SonicFastest | 1015330 | 621.98 MB/s | 992474 | 1102 | 5.3× |
| Sonic | 1017084 | 620.91 MB/s | 999399 | 1102 | 5.3× |
| Easyjson | 1146867 | 550.64 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1180162 | 535.11 MB/s | 986065 | 1201 | 4.6× |
| JSONV2 | 2150104 | 293.71 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2374729 | 196.61 MB/s | 2076505 | 30126 | 2.3× |
| Stdlib | 5375273 | 117.49 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 549378 | 1023.72 MB/s | 546570 | 429 | 9.5× |
| Lightning | 707906 | 794.47 MB/s | 769938 | 1235 | 7.4× |
| LightningArena | 716601 | 784.83 MB/s | 771690 | 1089 | 7.3× |
| Sonic | 1022030 | 550.29 MB/s | 946154 | 1476 | 5.1× |
| SonicFastest | 1032487 | 544.71 MB/s | 956018 | 1476 | 5.1× |
| Goccy | 1356882 | 414.49 MB/s | 1043601 | 1030 | 3.9× |
| Easyjson | 1759391 | 319.66 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2612534 | 215.27 MB/s | 2180439 | 30126 | 2.0× |
| JSONV2 | 2808407 | 200.26 MB/s | 927445 | 3482 | 1.9× |
| Stdlib | 5245433 | 107.22 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 579772 | 919.63 MB/s | 333416 | 2084 | 9.4× |
| Lightning | 598095 | 891.46 MB/s | 368224 | 2293 | 9.1× |
| LightningArena | 604628 | 881.83 MB/s | 368224 | 2293 | 9.0× |
| Easyjson | 1114723 | 478.31 MB/s | 428361 | 3273 | 4.9× |
| Sonic | 1156656 | 460.97 MB/s | 1038028 | 4351 | 4.7× |
| SonicFastest | 1158628 | 460.18 MB/s | 1040128 | 4351 | 4.7× |
| Goccy | 1333107 | 399.95 MB/s | 1167242 | 5409 | 4.1× |
| JSONV2 | 2570602 | 207.41 MB/s | 745447 | 13288 | 2.1× |
| LightningDecodeAny | 3366183 | 158.39 MB/s | 2992877 | 50076 | 1.6× |
| Stdlib | 5428275 | 98.22 MB/s | 798692 | 17133 | 1.0× |
