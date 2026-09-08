# JSON Deserialization Benchmarks

- generated 2026-09-08T16:56:11Z
- go version go1.26.8 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 84935 | 1498.51 MB/s | 49828 | 2 | 12.9× |
| LightningArena | 84964 | 1497.98 MB/s | 49833 | 2 | 12.9× |
| LightningDestructive | 85584 | 1487.14 MB/s | 49280 | 2 | 12.8× |
| SonicFastest | 189696 | 670.94 MB/s | 198488 | 10 | 5.8× |
| Sonic | 189896 | 670.23 MB/s | 198988 | 10 | 5.8× |
| Goccy | 210289 | 605.24 MB/s | 225399 | 884 | 5.2× |
| Easyjson | 214369 | 593.72 MB/s | 122864 | 14 | 5.1× |
| JSONV2 | 433698 | 293.46 MB/s | 195120 | 1805 | 2.5× |
| LightningDecodeAny | 460753 | 205.43 MB/s | 465086 | 9707 | 2.4× |
| Stdlib | 1094814 | 116.25 MB/s | 199697 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2674106 | 841.80 MB/s | 2532848 | 1143 | 9.9× |
| LightningArena | 2708951 | 830.97 MB/s | 2532848 | 1143 | 9.8× |
| Lightning | 2718927 | 827.92 MB/s | 2532850 | 1143 | 9.8× |
| Sonic | 4588649 | 490.57 MB/s | 15233813 | 970 | 5.8× |
| SonicFastest | 4600577 | 489.30 MB/s | 15238794 | 970 | 5.8× |
| Goccy | 10543852 | 213.49 MB/s | 4131218 | 56533 | 2.5× |
| Easyjson | 11029763 | 204.09 MB/s | 3099808 | 2120 | 2.4× |
| LightningDecodeAny | 11759984 | 191.41 MB/s | 19380210 | 223896 | 2.3× |
| JSONV2 | 16270632 | 138.35 MB/s | 3123214 | 3083 | 1.6× |
| Stdlib | 26589852 | 84.66 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 376944 | 717.36 MB/s | 397296 | 567 | 9.1× |
| LightningArena | 378799 | 713.84 MB/s | 397297 | 567 | 9.1× |
| Lightning | 378914 | 713.63 MB/s | 397297 | 567 | 9.1× |
| Sonic | 644245 | 419.72 MB/s | 478511 | 968 | 5.3× |
| SonicFastest | 655485 | 412.52 MB/s | 493359 | 968 | 5.2× |
| Easyjson | 1428668 | 189.27 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1433044 | 188.69 MB/s | 544290 | 8123 | 2.4× |
| LightningDecodeAny | 1648741 | 164.01 MB/s | 2543878 | 29687 | 2.1× |
| JSONV2 | 2128035 | 127.07 MB/s | 348158 | 1628 | 1.6× |
| Stdlib | 3438023 | 78.65 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 961148 | 1797.02 MB/s | 767576 | 2798 | 13.7× |
| LightningDestructive | 962841 | 1793.86 MB/s | 765560 | 2798 | 13.7× |
| LightningArena | 975311 | 1770.93 MB/s | 774977 | 2444 | 13.5× |
| Sonic | 2081064 | 829.96 MB/s | 2686811 | 4020 | 6.3× |
| SonicFastest | 2089587 | 826.58 MB/s | 2698795 | 4020 | 6.3× |
| Goccy | 2365067 | 730.30 MB/s | 2583372 | 14605 | 5.6× |
| Easyjson | 4253273 | 406.09 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4296519 | 402.00 MB/s | 1011631 | 7594 | 3.1× |
| LightningDecodeAny | 4451734 | 112.38 MB/s | 4964635 | 76577 | 3.0× |
| Stdlib | 13208622 | 130.76 MB/s | 1234448 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 841 | 2154.74 MB/s | 0 | 0 | 16.6× |
| Lightning | 842 | 2152.98 MB/s | 0 | 0 | 16.6× |
| LightningDestructive | 854 | 2122.15 MB/s | 0 | 0 | 16.3× |
| Easyjson | 2554 | 709.37 MB/s | 24 | 1 | 5.5× |
| Goccy | 2810 | 644.82 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6111 | 296.54 MB/s | 3757 | 40 | 2.3× |
| Sonic | 6123 | 295.95 MB/s | 3747 | 40 | 2.3× |
| JSONV2 | 7810 | 232.02 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7998 | 226.44 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13930 | 130.08 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 850 | 2132.12 MB/s | 0 | 0 | 16.4× |
| LightningArena | 857 | 2113.46 MB/s | 0 | 0 | 16.3× |
| LightningDestructive | 877 | 2065.35 MB/s | 0 | 0 | 15.9× |
| Easyjson | 2547 | 711.40 MB/s | 24 | 1 | 5.5× |
| Goccy | 2805 | 646.01 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6088 | 297.63 MB/s | 3800 | 40 | 2.3× |
| Sonic | 6096 | 297.22 MB/s | 3777 | 40 | 2.3× |
| JSONV2 | 7660 | 236.55 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 7984 | 226.84 MB/s | 7552 | 158 | 1.7× |
| Stdlib | 13970 | 129.70 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1023 | 1771.80 MB/s | 144 | 10 | 13.6× |
| LightningArena | 1029 | 1760.79 MB/s | 144 | 10 | 13.5× |
| LightningDestructive | 1083 | 1672.77 MB/s | 144 | 10 | 12.8× |
| Easyjson | 2767 | 654.82 MB/s | 144 | 10 | 5.0× |
| Goccy | 2989 | 606.14 MB/s | 2600 | 5 | 4.7× |
| Sonic | 6193 | 292.60 MB/s | 3807 | 42 | 2.2× |
| SonicFastest | 6224 | 291.15 MB/s | 3851 | 42 | 2.2× |
| LightningDecodeAny | 8030 | 225.54 MB/s | 7552 | 158 | 1.7× |
| JSONV2 | 8096 | 223.80 MB/s | 632 | 7 | 1.7× |
| Stdlib | 13906 | 130.31 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 604 | 818.27 MB/s | 160 | 1 | 9.1× |
| LightningDestructive | 607 | 813.87 MB/s | 160 | 1 | 9.0× |
| SonicFastest | 1252 | 394.59 MB/s | 995 | 6 | 4.4× |
| Sonic | 1255 | 393.72 MB/s | 978 | 6 | 4.4× |
| LightningDecodeAny | 1307 | 377.13 MB/s | 1296 | 26 | 4.2× |
| LightningArena | 1395 | 354.03 MB/s | 4120 | 2 | 3.9× |
| Easyjson | 2234 | 221.17 MB/s | 448 | 3 | 2.5× |
| Goccy | 2423 | 203.90 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3249 | 152.04 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5477 | 90.19 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 376 | 612.41 MB/s | 160 | 1 | 10.8× |
| Lightning | 385 | 597.18 MB/s | 160 | 1 | 10.6× |
| Sonic | 897 | 256.53 MB/s | 659 | 6 | 4.5× |
| SonicFastest | 901 | 255.24 MB/s | 663 | 6 | 4.5× |
| LightningDecodeAny | 1148 | 199.47 MB/s | 1296 | 26 | 3.5× |
| LightningArena | 1187 | 193.78 MB/s | 4120 | 2 | 3.4× |
| Easyjson | 1410 | 163.10 MB/s | 448 | 3 | 2.9× |
| Goccy | 1585 | 145.12 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2448 | 93.95 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4067 | 56.55 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 51289 | 1269.91 MB/s | 97220 | 98 | 10.7× |
| Lightning | 53190 | 1224.51 MB/s | 103662 | 99 | 10.3× |
| LightningArena | 53956 | 1207.13 MB/s | 103669 | 99 | 10.2× |
| SonicFastest | 100415 | 648.63 MB/s | 155591 | 75 | 5.5× |
| Sonic | 100609 | 647.37 MB/s | 155677 | 75 | 5.5× |
| Goccy | 149237 | 436.43 MB/s | 228949 | 134 | 3.7× |
| LightningDecodeAny | 184126 | 289.63 MB/s | 180444 | 3241 | 3.0× |
| JSONV2 | 229079 | 284.32 MB/s | 206652 | 607 | 2.4× |
| Stdlib | 550060 | 118.41 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2128401 | 911.70 MB/s | 2185297 | 1350 | 10.8× |
| LightningArena | 2185947 | 887.70 MB/s | 2185297 | 1350 | 10.6× |
| Lightning | 2200537 | 881.82 MB/s | 2185298 | 1350 | 10.5× |
| SonicFastest | 4709566 | 412.03 MB/s | 14608703 | 1407 | 4.9× |
| Goccy | 4797351 | 404.49 MB/s | 4064291 | 13510 | 4.8× |
| Sonic | 4815786 | 402.94 MB/s | 14608649 | 1407 | 4.8× |
| Easyjson | 7616951 | 254.76 MB/s | 3871266 | 15043 | 3.0× |
| LightningDecodeAny | 9323362 | 208.13 MB/s | 7063039 | 218633 | 2.5× |
| JSONV2 | 11249628 | 172.49 MB/s | 3237224 | 13947 | 2.1× |
| Stdlib | 23091434 | 84.03 MB/s | 3551321 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 893109 | 3726.12 MB/s | 351704 | 1286 | 23.4× |
| LightningArena | 1344895 | 2474.42 MB/s | 2434368 | 1413 | 15.6× |
| Lightning | 1356035 | 2454.09 MB/s | 2434434 | 1413 | 15.4× |
| Sonic | 2697340 | 1233.75 MB/s | 6441540 | 4248 | 7.8× |
| SonicFastest | 2698174 | 1233.36 MB/s | 6493350 | 4248 | 7.8× |
| LightningDecodeAny | 3432088 | 895.60 MB/s | 4825432 | 55311 | 6.1× |
| Goccy | 4651706 | 715.40 MB/s | 3948909 | 3816 | 4.5× |
| JSONV2 | 7496505 | 443.92 MB/s | 5364515 | 13243 | 2.8× |
| Stdlib | 20920494 | 159.07 MB/s | 5565608 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 181469 | 1214.23 MB/s | 135872 | 226 | 11.1× |
| LightningArena | 182311 | 1208.63 MB/s | 135872 | 226 | 11.1× |
| LightningDestructive | 182529 | 1207.18 MB/s | 135872 | 226 | 11.1× |
| Sonic | 393418 | 560.08 MB/s | 323233 | 398 | 5.1× |
| SonicFastest | 413220 | 533.24 MB/s | 362734 | 398 | 4.9× |
| Goccy | 454562 | 484.74 MB/s | 365078 | 1067 | 4.5× |
| Easyjson | 548917 | 401.42 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 731713 | 301.14 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 862354 | 125.60 MB/s | 897217 | 11703 | 2.3× |
| Stdlib | 2023047 | 108.92 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 9930712 | 815.66 MB/s | 8109654 | 20809 | 8.9× |
| LightningDestructive | 9933879 | 815.40 MB/s | 8109648 | 20809 | 8.9× |
| Lightning | 9942335 | 814.70 MB/s | 8109655 | 20809 | 8.9× |
| Sonic | 17102480 | 473.62 MB/s | 70887643 | 40014 | 5.2× |
| SonicFastest | 17318203 | 467.72 MB/s | 70902510 | 40014 | 5.1× |
| Goccy | 23807525 | 340.23 MB/s | 17221092 | 107149 | 3.7× |
| Easyjson | 30780536 | 263.15 MB/s | 15059619 | 41643 | 2.9× |
| LightningDecodeAny | 35644330 | 145.97 MB/s | 46279354 | 747112 | 2.5× |
| JSONV2 | 44368463 | 182.56 MB/s | 15233760 | 78972 | 2.0× |
| Stdlib | 88326900 | 91.71 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 4443246 | 671.46 MB/s | 3780456 | 1514 | 10.6× |
| LightningDestructive | 4586722 | 650.46 MB/s | 3758856 | 29356 | 10.2× |
| Lightning | 4749856 | 628.12 MB/s | 3758859 | 29356 | 9.9× |
| SonicFastest | 8842956 | 337.38 MB/s | 26604049 | 56760 | 5.3× |
| Sonic | 8883420 | 335.85 MB/s | 26566220 | 56760 | 5.3× |
| LightningDecodeAny | 15899732 | 115.36 MB/s | 23982581 | 351152 | 2.9× |
| Easyjson | 16786647 | 177.73 MB/s | 9479440 | 30115 | 2.8× |
| Goccy | 16998709 | 175.51 MB/s | 10576520 | 273648 | 2.8× |
| JSONV2 | 24776233 | 120.42 MB/s | 9257190 | 86278 | 1.9× |
| Stdlib | 46886494 | 63.63 MB/s | 9258091 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 931574 | 776.75 MB/s | 907601 | 3618 | 12.2× |
| LightningArena | 955518 | 757.28 MB/s | 916259 | 37 | 11.9× |
| Lightning | 991649 | 729.69 MB/s | 907599 | 3618 | 11.5× |
| Sonic | 1813983 | 398.90 MB/s | 3182905 | 7226 | 6.3× |
| SonicFastest | 1814261 | 398.84 MB/s | 3190781 | 7226 | 6.3× |
| LightningDecodeAny | 4127534 | 157.62 MB/s | 6500457 | 76546 | 2.8× |
| Easyjson | 4200478 | 172.27 MB/s | 2847905 | 3698 | 2.7× |
| Goccy | 4889248 | 148.00 MB/s | 2809531 | 80274 | 2.3× |
| JSONV2 | 5625959 | 128.62 MB/s | 2704621 | 7318 | 2.0× |
| Stdlib | 11408896 | 63.42 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1383346 | 1140.24 MB/s | 907600 | 3618 | 11.3× |
| LightningArena | 1383879 | 1139.81 MB/s | 916256 | 37 | 11.3× |
| Lightning | 1427268 | 1105.16 MB/s | 907595 | 3618 | 10.9× |
| Sonic | 2324348 | 678.62 MB/s | 5803425 | 7226 | 6.7× |
| SonicFastest | 2339651 | 674.18 MB/s | 5803087 | 7226 | 6.7× |
| LightningDecodeAny | 3759228 | 200.41 MB/s | 6500457 | 76546 | 4.1× |
| Easyjson | 5601384 | 281.60 MB/s | 2847905 | 3698 | 2.8× |
| Goccy | 5676767 | 277.86 MB/s | 3594501 | 80268 | 2.7× |
| JSONV2 | 6386599 | 246.98 MB/s | 2704591 | 7318 | 2.4× |
| Stdlib | 15578148 | 101.25 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 157908 | 950.70 MB/s | 81920 | 1 | 11.8× |
| LightningArena | 158227 | 948.79 MB/s | 81920 | 1 | 11.8× |
| Lightning | 158246 | 948.67 MB/s | 81920 | 1 | 11.8× |
| SonicFastest | 274465 | 546.97 MB/s | 252440 | 6 | 6.8× |
| Sonic | 277568 | 540.85 MB/s | 260598 | 6 | 6.7× |
| LightningDecodeAny | 431621 | 347.81 MB/s | 745764 | 10016 | 4.3× |
| Goccy | 861078 | 174.34 MB/s | 323643 | 10004 | 2.2× |
| JSONV2 | 1069623 | 140.35 MB/s | 357716 | 20 | 1.7× |
| Stdlib | 1860548 | 80.69 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 27974 | 1005.12 MB/s | 29088 | 101 | 10.8× |
| Lightning | 28004 | 1004.02 MB/s | 29256 | 101 | 10.8× |
| LightningArena | 28026 | 1003.24 MB/s | 29260 | 101 | 10.8× |
| Sonic | 65259 | 430.85 MB/s | 48784 | 103 | 4.6× |
| SonicFastest | 65267 | 430.80 MB/s | 48765 | 103 | 4.6× |
| Easyjson | 69436 | 404.93 MB/s | 32304 | 138 | 4.3× |
| Goccy | 73271 | 383.74 MB/s | 59253 | 188 | 4.1× |
| JSONV2 | 134368 | 209.25 MB/s | 36896 | 242 | 2.2× |
| LightningDecodeAny | 155585 | 180.72 MB/s | 141163 | 2641 | 1.9× |
| Stdlib | 301423 | 93.28 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 1506 | 1545.97 MB/s | 32 | 1 | 15.0× |
| Lightning | 1509 | 1543.03 MB/s | 32 | 1 | 15.0× |
| LightningDestructive | 1583 | 1470.73 MB/s | 32 | 1 | 14.3× |
| Easyjson | 4235 | 549.77 MB/s | 192 | 2 | 5.3× |
| Goccy | 4240 | 549.02 MB/s | 3649 | 4 | 5.3× |
| Sonic | 5166 | 450.66 MB/s | 4369 | 6 | 4.4× |
| SonicFastest | 5172 | 450.08 MB/s | 4333 | 6 | 4.4× |
| JSONV2 | 8524 | 273.10 MB/s | 1000 | 6 | 2.6× |
| LightningDecodeAny | 10088 | 167.03 MB/s | 10200 | 195 | 2.2× |
| Stdlib | 22582 | 103.09 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningArena | 178 | 1061.91 MB/s | 0 | 0 | 13.4× |
| Lightning | 178 | 1060.94 MB/s | 0 | 0 | 13.4× |
| LightningDestructive | 181 | 1046.72 MB/s | 0 | 0 | 13.2× |
| Goccy | 387 | 488.31 MB/s | 304 | 2 | 6.2× |
| Easyjson | 492 | 384.23 MB/s | 0 | 0 | 4.9× |
| Sonic | 808 | 233.92 MB/s | 497 | 4 | 3.0× |
| SonicFastest | 809 | 233.50 MB/s | 502 | 4 | 2.9× |
| JSONV2 | 1031 | 183.28 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1235 | 108.55 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2387 | 79.18 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1098 | 1995.81 MB/s | 0 | 0 | 14.4× |
| LightningArena | 1098 | 1996.33 MB/s | 0 | 0 | 14.4× |
| LightningDestructive | 1122 | 1952.17 MB/s | 0 | 0 | 14.1× |
| Goccy | 3190 | 686.88 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3201 | 684.52 MB/s | 24 | 1 | 4.9× |
| Sonic | 6477 | 338.27 MB/s | 4000 | 40 | 2.4× |
| SonicFastest | 6501 | 337.05 MB/s | 4002 | 40 | 2.4× |
| JSONV2 | 7923 | 276.54 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8040 | 225.25 MB/s | 7552 | 158 | 2.0× |
| Stdlib | 15806 | 138.62 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 574572 | 888.45 MB/s | 318400 | 1005 | 10.4× |
| LightningDestructive | 578329 | 882.67 MB/s | 318400 | 1005 | 10.4× |
| LightningArena | 578885 | 881.83 MB/s | 318400 | 1005 | 10.3× |
| SonicFastest | 1182408 | 431.73 MB/s | 910441 | 2006 | 5.1× |
| Sonic | 1184233 | 431.06 MB/s | 919100 | 2006 | 5.1× |
| Goccy | 1196759 | 426.55 MB/s | 1140681 | 5006 | 5.0× |
| Easyjson | 1553947 | 328.50 MB/s | 863777 | 3012 | 3.9× |
| JSONV2 | 3250678 | 157.04 MB/s | 1076017 | 12646 | 1.8× |
| LightningDecodeAny | 3420131 | 134.93 MB/s | 2950648 | 64018 | 1.8× |
| Stdlib | 5986940 | 85.26 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 480 | 41193.65 MB/s | 0 | 0 | 225.9× |
| LightningArena | 481 | 41172.08 MB/s | 0 | 0 | 225.8× |
| LightningDestructive | 502 | 39452.85 MB/s | 0 | 0 | 216.4× |
| Goccy | 21006 | 942.06 MB/s | 20491 | 2 | 5.2× |
| Sonic | 28693 | 689.68 MB/s | 22721 | 4 | 3.8× |
| SonicFastest | 29015 | 682.03 MB/s | 23397 | 4 | 3.7× |
| JSONV2 | 29532 | 670.09 MB/s | 8 | 1 | 3.7× |
| Easyjson | 82031 | 241.24 MB/s | 0 | 0 | 1.3× |
| LightningDecodeAny | 87064 | 227.28 MB/s | 116864 | 2015 | 1.2× |
| Stdlib | 108532 | 182.33 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1747 | 10374.47 MB/s | 0 | 0 | 58.8× |
| LightningArena | 1812 | 10000.49 MB/s | 405 | 0 | 56.7× |
| Lightning | 1814 | 9990.29 MB/s | 405 | 0 | 56.6× |
| Easyjson | 3948 | 4591.13 MB/s | 432 | 2 | 26.0× |
| Sonic | 10297 | 1760.20 MB/s | 22961 | 6 | 10.0× |
| SonicFastest | 10304 | 1758.90 MB/s | 23080 | 6 | 10.0× |
| Goccy | 16575 | 1093.44 MB/s | 19459 | 2 | 6.2× |
| LightningDecodeAny | 16685 | 1071.72 MB/s | 29103 | 189 | 6.2× |
| JSONV2 | 45787 | 395.84 MB/s | 16499 | 50 | 2.2× |
| Stdlib | 102702 | 176.47 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2152259 | 933.20 MB/s | 3089564 | 6821 | 8.6× |
| LightningArena | 2205919 | 910.50 MB/s | 3100185 | 6700 | 8.4× |
| Lightning | 2211558 | 908.18 MB/s | 3096252 | 6822 | 8.4× |
| Goccy | 4343823 | 462.38 MB/s | 5412471 | 15831 | 4.3× |
| SonicFastest | 4530565 | 443.32 MB/s | 10925298 | 13683 | 4.1× |
| Sonic | 4555629 | 440.88 MB/s | 10972585 | 13683 | 4.1× |
| Easyjson | 4962853 | 404.71 MB/s | 2981488 | 7439 | 3.7× |
| JSONV2 | 6990607 | 287.31 MB/s | 3173684 | 14563 | 2.7× |
| LightningDecodeAny | 7459375 | 153.14 MB/s | 8516390 | 134005 | 2.5× |
| Stdlib | 18540901 | 108.33 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 850 | 645.66 MB/s | 480 | 1 | 6.6× |
| LightningDestructive | 858 | 639.69 MB/s | 480 | 1 | 6.6× |
| LightningArena | 859 | 639.22 MB/s | 480 | 1 | 6.6× |
| LightningDecodeAny | 1660 | 330.07 MB/s | 2021 | 46 | 3.4× |
| Easyjson | 2200 | 249.53 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2674 | 205.27 MB/s | 1962 | 26 | 2.1× |
| Sonic | 2679 | 204.90 MB/s | 1954 | 26 | 2.1× |
| Goccy | 2977 | 184.39 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3375 | 162.66 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5631 | 97.50 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 419248 | 1506.30 MB/s | 402728 | 545 | 12.8× |
| LightningArena | 455122 | 1387.57 MB/s | 451338 | 404 | 11.8× |
| Lightning | 456492 | 1383.41 MB/s | 449415 | 548 | 11.8× |
| Sonic | 1037777 | 608.53 MB/s | 1010103 | 1102 | 5.2× |
| SonicFastest | 1039422 | 607.56 MB/s | 1010102 | 1102 | 5.2× |
| Easyjson | 1159048 | 544.86 MB/s | 422505 | 936 | 4.6× |
| Goccy | 1180813 | 534.81 MB/s | 985324 | 1201 | 4.6× |
| JSONV2 | 2161335 | 292.19 MB/s | 571614 | 3144 | 2.5× |
| LightningDecodeAny | 2440428 | 191.32 MB/s | 2081274 | 29820 | 2.2× |
| Stdlib | 5373193 | 117.53 MB/s | 654666 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 576477 | 975.60 MB/s | 393091 | 426 | 9.2× |
| Lightning | 649174 | 866.34 MB/s | 508175 | 433 | 8.1× |
| LightningArena | 651573 | 863.15 MB/s | 510001 | 287 | 8.1× |
| Sonic | 1048445 | 536.42 MB/s | 965461 | 1476 | 5.0× |
| SonicFastest | 1060700 | 530.22 MB/s | 977634 | 1476 | 5.0× |
| Goccy | 1361504 | 413.08 MB/s | 1041814 | 1030 | 3.9× |
| Easyjson | 1771831 | 317.42 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2466614 | 228.01 MB/s | 2079981 | 29328 | 2.1× |
| JSONV2 | 2811073 | 200.07 MB/s | 927441 | 3482 | 1.9× |
| Stdlib | 5277155 | 106.57 MB/s | 1011674 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 582928 | 914.66 MB/s | 333416 | 2084 | 9.4× |
| Lightning | 608214 | 876.63 MB/s | 367876 | 2086 | 9.0× |
| LightningArena | 610202 | 873.77 MB/s | 367907 | 2086 | 8.9× |
| Easyjson | 1143714 | 466.18 MB/s | 428362 | 3273 | 4.8× |
| Sonic | 1164879 | 457.71 MB/s | 1042646 | 4351 | 4.7× |
| SonicFastest | 1171319 | 455.19 MB/s | 1037191 | 4351 | 4.7× |
| Goccy | 1338925 | 398.21 MB/s | 1167222 | 5409 | 4.1× |
| JSONV2 | 2583075 | 206.41 MB/s | 745466 | 13288 | 2.1× |
| LightningDecodeAny | 3451766 | 154.46 MB/s | 3002945 | 49873 | 1.6× |
| Stdlib | 5459408 | 97.66 MB/s | 798693 | 17133 | 1.0× |
