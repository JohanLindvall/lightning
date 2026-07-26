# JSON Deserialization Benchmarks

- generated 2026-07-26T15:08:51Z
- go version go1.26.5 linux/arm64
- cpu: unknown (4 cores)

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 105970 | 1201.05 MB/s | 49760 | 3 | 10.5× |
| LightningDestructive | 106640 | 1193.51 MB/s | 49280 | 2 | 10.4× |
| Sonic | 185966 | 684.40 MB/s | 196944 | 10 | 6.0× |
| SonicFastest | 189963 | 670.00 MB/s | 204014 | 10 | 5.8× |
| Goccy | 207507 | 613.35 MB/s | 225260 | 884 | 5.3× |
| Easyjson | 214333 | 593.82 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 432066 | 294.57 MB/s | 195118 | 1805 | 2.6× |
| LightningDecodeAny | 464494 | 203.78 MB/s | 465730 | 9708 | 2.4× |
| Stdlib | 1109236 | 114.74 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3663947 | 614.38 MB/s | 2532848 | 1143 | 7.1× |
| Lightning | 3693782 | 609.42 MB/s | 2532850 | 1143 | 7.1× |
| Sonic | 4770423 | 471.88 MB/s | 15232101 | 970 | 5.5× |
| SonicFastest | 4847034 | 464.42 MB/s | 15233865 | 970 | 5.4× |
| Goccy | 10279058 | 218.99 MB/s | 4113743 | 56532 | 2.5× |
| Easyjson | 11087903 | 203.02 MB/s | 3099809 | 2120 | 2.4× |
| LightningDecodeAny | 12619663 | 178.37 MB/s | 19380209 | 223896 | 2.1× |
| JSONV2 | 16112805 | 139.71 MB/s | 3123206 | 3083 | 1.6× |
| Stdlib | 26139687 | 86.12 MB/s | 3123396 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 479185 | 564.30 MB/s | 397296 | 567 | 7.0× |
| LightningDestructive | 479863 | 563.50 MB/s | 397296 | 567 | 7.0× |
| Sonic | 642501 | 420.86 MB/s | 482146 | 968 | 5.2× |
| SonicFastest | 652005 | 414.73 MB/s | 494694 | 968 | 5.1× |
| Easyjson | 1406020 | 192.32 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1424778 | 189.79 MB/s | 543312 | 8122 | 2.4× |
| LightningDecodeAny | 1764646 | 153.23 MB/s | 2543881 | 29687 | 1.9× |
| JSONV2 | 2112937 | 127.97 MB/s | 348158 | 1628 | 1.6× |
| Stdlib | 3356616 | 80.56 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1175046 | 1469.90 MB/s | 765560 | 2798 | 11.3× |
| Lightning | 1175835 | 1468.92 MB/s | 765602 | 2799 | 11.3× |
| Sonic | 2082504 | 829.39 MB/s | 2744778 | 4020 | 6.4× |
| SonicFastest | 2104881 | 820.57 MB/s | 2692650 | 4020 | 6.3× |
| Goccy | 2432003 | 710.20 MB/s | 2583286 | 14605 | 5.4× |
| Easyjson | 4244611 | 406.92 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4316120 | 400.18 MB/s | 1011631 | 7594 | 3.1× |
| LightningDecodeAny | 4512351 | 110.87 MB/s | 4954733 | 76576 | 2.9× |
| Stdlib | 13232937 | 130.52 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1172 | 1545.97 MB/s | 0 | 0 | 12.0× |
| LightningDestructive | 1188 | 1525.01 MB/s | 0 | 0 | 11.8× |
| Easyjson | 2520 | 718.97 MB/s | 24 | 1 | 5.6× |
| Goccy | 2799 | 647.26 MB/s | 2608 | 4 | 5.0× |
| Sonic | 5966 | 303.74 MB/s | 3726 | 40 | 2.4× |
| SonicFastest | 5985 | 302.76 MB/s | 3759 | 40 | 2.3× |
| JSONV2 | 7817 | 231.82 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8347 | 216.97 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14024 | 129.21 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1204 | 1504.96 MB/s | 0 | 0 | 11.7× |
| LightningDestructive | 1212 | 1494.52 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2534 | 715.07 MB/s | 24 | 1 | 5.6× |
| Goccy | 2858 | 634.07 MB/s | 2608 | 4 | 4.9× |
| SonicFastest | 5959 | 304.10 MB/s | 3714 | 40 | 2.4× |
| Sonic | 5964 | 303.84 MB/s | 3731 | 40 | 2.4× |
| JSONV2 | 7772 | 233.14 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8339 | 217.17 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14073 | 128.76 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1357 | 1334.85 MB/s | 144 | 10 | 10.4× |
| LightningDestructive | 1411 | 1283.97 MB/s | 144 | 10 | 10.0× |
| Easyjson | 2764 | 655.53 MB/s | 144 | 10 | 5.1× |
| Goccy | 2917 | 621.10 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6162 | 294.07 MB/s | 3796 | 42 | 2.3× |
| Sonic | 6190 | 292.74 MB/s | 3786 | 42 | 2.3× |
| JSONV2 | 8044 | 225.27 MB/s | 632 | 7 | 1.7× |
| LightningDecodeAny | 8290 | 218.44 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14049 | 128.98 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 730 | 677.21 MB/s | 160 | 1 | 7.6× |
| Lightning | 732 | 674.41 MB/s | 160 | 1 | 7.6× |
| SonicFastest | 1252 | 394.68 MB/s | 990 | 6 | 4.4× |
| Sonic | 1255 | 393.65 MB/s | 982 | 6 | 4.4× |
| LightningDecodeAny | 1425 | 345.88 MB/s | 1296 | 26 | 3.9× |
| Easyjson | 2209 | 223.66 MB/s | 448 | 3 | 2.5× |
| Goccy | 2456 | 201.11 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3253 | 151.84 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5553 | 88.96 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 435 | 528.95 MB/s | 160 | 1 | 9.6× |
| Lightning | 436 | 528.13 MB/s | 160 | 1 | 9.6× |
| Sonic | 918 | 250.61 MB/s | 687 | 6 | 4.6× |
| SonicFastest | 922 | 249.39 MB/s | 698 | 6 | 4.5× |
| LightningDecodeAny | 1208 | 189.54 MB/s | 1296 | 26 | 3.5× |
| Easyjson | 1413 | 162.74 MB/s | 448 | 3 | 3.0× |
| Goccy | 1619 | 142.02 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2485 | 92.55 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4178 | 55.05 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 71792 | 907.23 MB/s | 164880 | 105 | 7.6× |
| LightningDestructive | 72260 | 901.36 MB/s | 158660 | 100 | 7.6× |
| Sonic | 100256 | 649.66 MB/s | 156191 | 75 | 5.5× |
| SonicFastest | 100803 | 646.13 MB/s | 157810 | 75 | 5.4× |
| Goccy | 150097 | 433.93 MB/s | 229324 | 134 | 3.7× |
| LightningDecodeAny | 190449 | 280.02 MB/s | 180224 | 3245 | 2.9× |
| JSONV2 | 229153 | 284.23 MB/s | 206651 | 607 | 2.4× |
| Stdlib | 548214 | 118.81 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2563105 | 757.08 MB/s | 2864592 | 1380 | 9.2× |
| Lightning | 2640342 | 734.93 MB/s | 2864595 | 1380 | 8.9× |
| Sonic | 4705124 | 412.42 MB/s | 14608610 | 1407 | 5.0× |
| SonicFastest | 4734326 | 409.87 MB/s | 14608597 | 1407 | 5.0× |
| Goccy | 4794916 | 404.69 MB/s | 4065012 | 13510 | 4.9× |
| Easyjson | 7527936 | 257.77 MB/s | 3871266 | 15043 | 3.1× |
| LightningDecodeAny | 9534338 | 203.52 MB/s | 7064788 | 218633 | 2.5× |
| JSONV2 | 11184569 | 173.50 MB/s | 3237220 | 13947 | 2.1× |
| Stdlib | 23521587 | 82.50 MB/s | 3551322 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1110180 | 2997.56 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1849541 | 1799.27 MB/s | 2488905 | 2995 | 11.3× |
| SonicFastest | 2769742 | 1201.49 MB/s | 6372754 | 4248 | 7.5× |
| Sonic | 2795559 | 1190.40 MB/s | 6379851 | 4248 | 7.5× |
| LightningDecodeAny | 3818608 | 804.94 MB/s | 4886620 | 56892 | 5.5× |
| Goccy | 4658689 | 714.33 MB/s | 3948908 | 3816 | 4.5× |
| JSONV2 | 7514627 | 442.85 MB/s | 5364518 | 13243 | 2.8× |
| Stdlib | 20867538 | 159.47 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220338 | 1000.04 MB/s | 135872 | 226 | 9.3× |
| LightningDestructive | 222035 | 992.39 MB/s | 135872 | 226 | 9.2× |
| SonicFastest | 385072 | 572.22 MB/s | 310332 | 398 | 5.3× |
| Sonic | 396015 | 556.41 MB/s | 331571 | 398 | 5.2× |
| Goccy | 439785 | 501.03 MB/s | 364249 | 1067 | 4.6× |
| Easyjson | 551093 | 399.83 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 738623 | 298.32 MB/s | 129743 | 470 | 2.8× |
| LightningDecodeAny | 886309 | 122.21 MB/s | 897521 | 11703 | 2.3× |
| Stdlib | 2043347 | 107.84 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 11519562 | 703.16 MB/s | 11845072 | 20816 | 7.8× |
| Lightning | 11813577 | 685.66 MB/s | 11845073 | 20816 | 7.6× |
| Sonic | 16815006 | 481.71 MB/s | 70916870 | 40014 | 5.3× |
| SonicFastest | 16958792 | 477.63 MB/s | 70887315 | 40014 | 5.3× |
| Goccy | 23429599 | 345.72 MB/s | 17162658 | 107149 | 3.8× |
| Easyjson | 30796166 | 263.02 MB/s | 15059618 | 41643 | 2.9× |
| LightningDecodeAny | 37532937 | 138.63 MB/s | 46191128 | 747112 | 2.4× |
| JSONV2 | 43675508 | 185.46 MB/s | 15233740 | 78972 | 2.0× |
| Stdlib | 89433252 | 90.57 MB/s | 15665068 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 5694389 | 523.93 MB/s | 3758856 | 29356 | 8.3× |
| Lightning | 5813591 | 513.19 MB/s | 3758856 | 29356 | 8.1× |
| Sonic | 8683775 | 343.57 MB/s | 26570747 | 56760 | 5.4× |
| SonicFastest | 8745994 | 341.12 MB/s | 26505308 | 56760 | 5.4× |
| Easyjson | 16476814 | 181.07 MB/s | 9479440 | 30115 | 2.9× |
| Goccy | 16735651 | 178.27 MB/s | 10649202 | 273650 | 2.8× |
| LightningDecodeAny | 17036152 | 107.66 MB/s | 23982396 | 351152 | 2.8× |
| JSONV2 | 24409216 | 122.23 MB/s | 9257173 | 86278 | 1.9× |
| Stdlib | 47179547 | 63.24 MB/s | 9258092 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1240556 | 583.28 MB/s | 907600 | 3618 | 9.3× |
| Lightning | 1256751 | 575.77 MB/s | 907596 | 3618 | 9.2× |
| SonicFastest | 1780312 | 406.44 MB/s | 3187878 | 7226 | 6.5× |
| Sonic | 1781576 | 406.16 MB/s | 3174882 | 7226 | 6.5× |
| Easyjson | 4246189 | 170.41 MB/s | 2847906 | 3698 | 2.7× |
| LightningDecodeAny | 4257562 | 152.80 MB/s | 6500455 | 76546 | 2.7× |
| Goccy | 4767907 | 151.76 MB/s | 2836430 | 80275 | 2.4× |
| JSONV2 | 5988319 | 120.83 MB/s | 2704629 | 7318 | 1.9× |
| Stdlib | 11594474 | 62.41 MB/s | 2704553 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1844399 | 855.21 MB/s | 907601 | 3618 | 8.5× |
| Lightning | 1889506 | 834.80 MB/s | 907595 | 3618 | 8.3× |
| SonicFastest | 2295870 | 687.04 MB/s | 5807163 | 7226 | 6.9× |
| Sonic | 2303810 | 684.67 MB/s | 5794063 | 7226 | 6.8× |
| LightningDecodeAny | 3933799 | 191.52 MB/s | 6500457 | 76546 | 4.0× |
| Easyjson | 5590613 | 282.14 MB/s | 2847904 | 3698 | 2.8× |
| Goccy | 5697452 | 276.85 MB/s | 3607055 | 80268 | 2.8× |
| JSONV2 | 6708889 | 235.11 MB/s | 2704593 | 7318 | 2.3× |
| Stdlib | 15730113 | 100.28 MB/s | 2704551 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 212666 | 705.91 MB/s | 81920 | 1 | 8.6× |
| Lightning | 212822 | 705.40 MB/s | 81920 | 1 | 8.6× |
| Sonic | 270044 | 555.93 MB/s | 244642 | 6 | 6.8× |
| SonicFastest | 270759 | 554.46 MB/s | 247138 | 6 | 6.7× |
| LightningDecodeAny | 493902 | 303.95 MB/s | 745764 | 10016 | 3.7× |
| Goccy | 863680 | 173.82 MB/s | 324420 | 10004 | 2.1× |
| JSONV2 | 1102314 | 136.19 MB/s | 357715 | 20 | 1.7× |
| Stdlib | 1826186 | 82.21 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 32707 | 859.67 MB/s | 29216 | 103 | 9.3× |
| LightningDestructive | 33009 | 851.79 MB/s | 29088 | 101 | 9.2× |
| SonicFastest | 63904 | 439.99 MB/s | 47066 | 103 | 4.8× |
| Sonic | 63991 | 439.39 MB/s | 46954 | 103 | 4.7× |
| Easyjson | 68782 | 408.78 MB/s | 32304 | 138 | 4.4× |
| Goccy | 72180 | 389.54 MB/s | 59202 | 188 | 4.2× |
| JSONV2 | 134893 | 208.44 MB/s | 36895 | 242 | 2.3× |
| LightningDecodeAny | 154450 | 182.05 MB/s | 140592 | 2643 | 2.0× |
| Stdlib | 303612 | 92.61 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1967 | 1183.60 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2026 | 1149.23 MB/s | 32 | 1 | 11.2× |
| Goccy | 4168 | 558.60 MB/s | 3649 | 4 | 5.5× |
| Easyjson | 4228 | 550.57 MB/s | 192 | 2 | 5.4× |
| Sonic | 5117 | 454.96 MB/s | 4297 | 6 | 4.4× |
| SonicFastest | 5160 | 451.15 MB/s | 4342 | 6 | 4.4× |
| JSONV2 | 8515 | 273.40 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10539 | 159.88 MB/s | 10200 | 195 | 2.2× |
| Stdlib | 22744 | 102.36 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 222 | 853.38 MB/s | 0 | 0 | 11.0× |
| LightningDestructive | 224 | 844.20 MB/s | 0 | 0 | 10.8× |
| Goccy | 389 | 485.97 MB/s | 304 | 2 | 6.2× |
| Easyjson | 498 | 379.67 MB/s | 0 | 0 | 4.9× |
| Sonic | 803 | 235.48 MB/s | 502 | 4 | 3.0× |
| SonicFastest | 808 | 233.88 MB/s | 511 | 4 | 3.0× |
| JSONV2 | 1035 | 182.52 MB/s | 112 | 1 | 2.3× |
| LightningDecodeAny | 1249 | 107.31 MB/s | 1160 | 25 | 1.9× |
| Stdlib | 2429 | 77.80 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1527 | 1435.10 MB/s | 0 | 0 | 10.5× |
| LightningDestructive | 1537 | 1425.18 MB/s | 0 | 0 | 10.4× |
| Easyjson | 3195 | 685.66 MB/s | 24 | 1 | 5.0× |
| Goccy | 3250 | 674.12 MB/s | 2864 | 4 | 4.9× |
| Sonic | 6430 | 340.73 MB/s | 3993 | 40 | 2.5× |
| SonicFastest | 6462 | 339.05 MB/s | 4043 | 40 | 2.5× |
| JSONV2 | 7956 | 275.39 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8319 | 217.68 MB/s | 7536 | 158 | 1.9× |
| Stdlib | 16055 | 136.47 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 655911 | 778.27 MB/s | 457537 | 1009 | 9.2× |
| Lightning | 661264 | 771.97 MB/s | 457536 | 1009 | 9.2× |
| Goccy | 1167129 | 437.38 MB/s | 1138732 | 5006 | 5.2× |
| Sonic | 1178304 | 433.23 MB/s | 899791 | 2006 | 5.1× |
| SonicFastest | 1189907 | 429.00 MB/s | 920947 | 2006 | 5.1× |
| Easyjson | 1559282 | 327.38 MB/s | 863778 | 3012 | 3.9× |
| JSONV2 | 3218457 | 158.61 MB/s | 1076012 | 12646 | 1.9× |
| LightningDecodeAny | 3491820 | 132.16 MB/s | 2929689 | 64018 | 1.7× |
| Stdlib | 6054908 | 84.31 MB/s | 1162119 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1337 | 14797.30 MB/s | 0 | 0 | 85.2× |
| LightningDestructive | 1361 | 14540.54 MB/s | 0 | 0 | 83.7× |
| Goccy | 20150 | 982.11 MB/s | 20491 | 2 | 5.7× |
| Sonic | 27939 | 708.28 MB/s | 21663 | 4 | 4.1× |
| SonicFastest | 28126 | 703.57 MB/s | 22216 | 4 | 4.0× |
| JSONV2 | 29573 | 669.15 MB/s | 8 | 1 | 3.9× |
| LightningDecodeAny | 76957 | 257.13 MB/s | 116864 | 2015 | 1.5× |
| Easyjson | 86034 | 230.01 MB/s | 0 | 0 | 1.3× |
| Stdlib | 113867 | 173.79 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2625 | 6905.54 MB/s | 0 | 0 | 39.1× |
| Lightning | 2769 | 6545.45 MB/s | 432 | 2 | 37.1× |
| Easyjson | 3970 | 4565.54 MB/s | 432 | 2 | 25.9× |
| SonicFastest | 10154 | 1784.93 MB/s | 23208 | 6 | 10.1× |
| Sonic | 10417 | 1739.86 MB/s | 23466 | 6 | 9.9× |
| Goccy | 16189 | 1119.54 MB/s | 19459 | 2 | 6.3× |
| LightningDecodeAny | 17371 | 1029.40 MB/s | 29088 | 191 | 5.9× |
| JSONV2 | 45017 | 402.60 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102723 | 176.44 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2459605 | 816.59 MB/s | 3089564 | 6821 | 7.6× |
| Lightning | 2514234 | 798.85 MB/s | 3091278 | 6827 | 7.4× |
| Goccy | 4298097 | 467.30 MB/s | 5411605 | 15831 | 4.3× |
| Sonic | 4492330 | 447.09 MB/s | 10922278 | 13683 | 4.2× |
| SonicFastest | 4508382 | 445.50 MB/s | 10954597 | 13683 | 4.1× |
| Easyjson | 4953309 | 405.49 MB/s | 2981486 | 7439 | 3.8× |
| JSONV2 | 6994104 | 287.17 MB/s | 3173686 | 14563 | 2.7× |
| LightningDecodeAny | 7237177 | 157.84 MB/s | 8498331 | 134008 | 2.6× |
| Stdlib | 18693987 | 107.44 MB/s | 3589318 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 898 | 611.17 MB/s | 480 | 1 | 6.3× |
| LightningDestructive | 901 | 609.57 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 1762 | 311.06 MB/s | 2021 | 46 | 3.2× |
| Easyjson | 2183 | 251.53 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2701 | 203.30 MB/s | 1932 | 26 | 2.1× |
| Sonic | 2705 | 202.98 MB/s | 1952 | 26 | 2.1× |
| Goccy | 3032 | 181.04 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3325 | 165.10 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5675 | 96.74 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 515571 | 1224.88 MB/s | 402728 | 545 | 10.5× |
| Lightning | 571923 | 1104.19 MB/s | 451257 | 857 | 9.5× |
| SonicFastest | 1017758 | 620.50 MB/s | 995413 | 1102 | 5.3× |
| Sonic | 1018337 | 620.14 MB/s | 997511 | 1102 | 5.3× |
| Easyjson | 1152371 | 548.01 MB/s | 422505 | 936 | 4.7× |
| Goccy | 1173967 | 537.93 MB/s | 985000 | 1201 | 4.6× |
| JSONV2 | 2167871 | 291.31 MB/s | 571612 | 3144 | 2.5× |
| LightningDecodeAny | 2447574 | 190.76 MB/s | 2077367 | 30126 | 2.2× |
| Stdlib | 5415768 | 116.61 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 663060 | 848.20 MB/s | 546570 | 429 | 8.0× |
| Lightning | 840950 | 668.78 MB/s | 769941 | 1235 | 6.3× |
| SonicFastest | 1043210 | 539.11 MB/s | 978473 | 1476 | 5.1× |
| Sonic | 1052373 | 534.42 MB/s | 987077 | 1476 | 5.0× |
| Goccy | 1349399 | 416.78 MB/s | 1040903 | 1030 | 3.9× |
| Easyjson | 1758528 | 319.82 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2771532 | 202.92 MB/s | 2181319 | 30126 | 1.9× |
| JSONV2 | 2772823 | 202.83 MB/s | 927440 | 3482 | 1.9× |
| Stdlib | 5301578 | 106.08 MB/s | 1011672 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 666024 | 800.54 MB/s | 333416 | 2084 | 8.2× |
| Lightning | 690682 | 771.96 MB/s | 368224 | 2293 | 7.9× |
| Easyjson | 1124360 | 474.21 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1166229 | 457.18 MB/s | 1044324 | 4351 | 4.7× |
| Sonic | 1168187 | 456.42 MB/s | 1037610 | 4351 | 4.7× |
| Goccy | 1327495 | 401.64 MB/s | 1167246 | 5409 | 4.1× |
| JSONV2 | 2560754 | 208.21 MB/s | 745450 | 13288 | 2.1× |
| LightningDecodeAny | 3545067 | 150.40 MB/s | 2991148 | 50076 | 1.5× |
| Stdlib | 5489355 | 97.13 MB/s | 798692 | 17133 | 1.0× |
