# JSON Deserialization Benchmarks

- generated 2026-06-21T16:42:00Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 104413 | 1218.95 MB/s | 49280 | 2 | 10.6× |
| Lightning | 104619 | 1216.55 MB/s | 49760 | 3 | 10.6× |
| Sonic | 185072 | 687.71 MB/s | 199882 | 10 | 6.0× |
| SonicFastest | 186056 | 684.07 MB/s | 200006 | 10 | 5.9× |
| Goccy | 196849 | 646.56 MB/s | 225485 | 884 | 5.6× |
| Easyjson | 212322 | 599.44 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 429988 | 296.00 MB/s | 195119 | 1805 | 2.6× |
| LightningDecodeAny | 460594 | 205.50 MB/s | 465586 | 9714 | 2.4× |
| Stdlib | 1106429 | 115.03 MB/s | 199696 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 3731623 | 603.24 MB/s | 3122872 | 3081 | 7.0× |
| Lightning | 3792211 | 593.60 MB/s | 3122876 | 3081 | 6.9× |
| Sonic | 4474292 | 503.11 MB/s | 15233737 | 970 | 5.8× |
| SonicFastest | 4516117 | 498.45 MB/s | 15243623 | 970 | 5.8× |
| Goccy | 10463809 | 215.13 MB/s | 4119875 | 56532 | 2.5× |
| Easyjson | 11180285 | 201.34 MB/s | 3099808 | 2120 | 2.3× |
| LightningDecodeAny | 11555203 | 194.81 MB/s | 7938298 | 281383 | 2.3× |
| JSONV2 | 15876050 | 141.79 MB/s | 3123214 | 3083 | 1.6× |
| Stdlib | 26095281 | 86.26 MB/s | 3123397 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 487875 | 554.25 MB/s | 348024 | 1627 | 6.9× |
| LightningDestructive | 495058 | 546.21 MB/s | 348024 | 1627 | 6.8× |
| Sonic | 632029 | 427.83 MB/s | 482938 | 968 | 5.3× |
| SonicFastest | 635958 | 425.19 MB/s | 482685 | 968 | 5.3× |
| Easyjson | 1402803 | 192.76 MB/s | 330272 | 749 | 2.4× |
| Goccy | 1437124 | 188.16 MB/s | 543093 | 8122 | 2.3× |
| LightningDecodeAny | 1665717 | 162.33 MB/s | 1011486 | 37901 | 2.0× |
| JSONV2 | 2097273 | 128.93 MB/s | 348156 | 1628 | 1.6× |
| Stdlib | 3354939 | 80.60 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1353534 | 1276.07 MB/s | 959848 | 5881 | 9.8× |
| Lightning | 1366060 | 1264.37 MB/s | 959890 | 5882 | 9.7× |
| Sonic | 2049457 | 842.76 MB/s | 2776422 | 4020 | 6.5× |
| SonicFastest | 2049786 | 842.63 MB/s | 2798992 | 4020 | 6.5× |
| Goccy | 2318573 | 744.94 MB/s | 2582489 | 14604 | 5.7× |
| Easyjson | 4239925 | 407.37 MB/s | 972032 | 5389 | 3.1× |
| JSONV2 | 4274182 | 404.10 MB/s | 1011632 | 7594 | 3.1× |
| LightningDecodeAny | 4587308 | 109.06 MB/s | 4976205 | 81466 | 2.9× |
| Stdlib | 13231010 | 130.54 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1186 | 1528.02 MB/s | 0 | 0 | 11.8× |
| LightningDestructive | 1206 | 1502.45 MB/s | 0 | 0 | 11.6× |
| Easyjson | 2551 | 710.24 MB/s | 24 | 1 | 5.5× |
| Goccy | 2773 | 653.44 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 6017 | 301.13 MB/s | 3749 | 40 | 2.3× |
| Sonic | 6020 | 300.97 MB/s | 3720 | 40 | 2.3× |
| JSONV2 | 7676 | 236.06 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8189 | 221.16 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14042 | 129.04 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1213 | 1493.92 MB/s | 0 | 0 | 11.6× |
| LightningDestructive | 1232 | 1470.38 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2550 | 710.66 MB/s | 24 | 1 | 5.5× |
| Goccy | 2798 | 647.62 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 6025 | 300.73 MB/s | 3753 | 40 | 2.3× |
| Sonic | 6041 | 299.96 MB/s | 3773 | 40 | 2.3× |
| JSONV2 | 7811 | 231.99 MB/s | 640 | 6 | 1.8× |
| LightningDecodeAny | 8152 | 222.16 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14073 | 128.75 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1389 | 1304.41 MB/s | 144 | 10 | 10.1× |
| LightningDestructive | 1412 | 1283.47 MB/s | 144 | 10 | 9.9× |
| Easyjson | 2744 | 660.30 MB/s | 144 | 10 | 5.1× |
| Goccy | 2899 | 625.09 MB/s | 2600 | 5 | 4.8× |
| SonicFastest | 6105 | 296.81 MB/s | 3764 | 42 | 2.3× |
| Sonic | 6126 | 295.78 MB/s | 3777 | 42 | 2.3× |
| JSONV2 | 7967 | 227.45 MB/s | 632 | 7 | 1.8× |
| LightningDecodeAny | 8152 | 222.16 MB/s | 7536 | 158 | 1.7× |
| Stdlib | 14021 | 129.23 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 768 | 643.04 MB/s | 160 | 1 | 7.2× |
| Lightning | 771 | 640.56 MB/s | 160 | 1 | 7.2× |
| Sonic | 1228 | 402.15 MB/s | 968 | 6 | 4.5× |
| SonicFastest | 1231 | 401.26 MB/s | 973 | 6 | 4.5× |
| LightningDecodeAny | 1638 | 300.95 MB/s | 1536 | 30 | 3.4× |
| Easyjson | 2209 | 223.64 MB/s | 448 | 3 | 2.5× |
| Goccy | 2415 | 204.59 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3256 | 151.71 MB/s | 528 | 7 | 1.7× |
| Stdlib | 5524 | 89.43 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 452 | 508.49 MB/s | 160 | 1 | 9.2× |
| LightningDestructive | 454 | 506.12 MB/s | 160 | 1 | 9.1× |
| Sonic | 877 | 262.15 MB/s | 656 | 6 | 4.7× |
| SonicFastest | 880 | 261.46 MB/s | 655 | 6 | 4.7× |
| LightningDecodeAny | 1352 | 169.34 MB/s | 1536 | 30 | 3.1× |
| Easyjson | 1388 | 165.66 MB/s | 448 | 3 | 3.0× |
| Goccy | 1568 | 146.67 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2450 | 93.87 MB/s | 528 | 7 | 1.7× |
| Stdlib | 4139 | 55.56 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 76252 | 854.17 MB/s | 172432 | 107 | 7.2× |
| LightningDestructive | 77968 | 835.37 MB/s | 166212 | 102 | 7.0× |
| Sonic | 99184 | 656.68 MB/s | 157842 | 75 | 5.5× |
| SonicFastest | 102898 | 632.98 MB/s | 161142 | 75 | 5.3× |
| Goccy | 146568 | 444.38 MB/s | 228229 | 134 | 3.7× |
| LightningDecodeAny | 185329 | 287.75 MB/s | 176960 | 3252 | 3.0× |
| JSONV2 | 227113 | 286.78 MB/s | 206652 | 607 | 2.4× |
| Stdlib | 549289 | 118.58 MB/s | 214616 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2588884 | 749.54 MB/s | 2846864 | 2238 | 9.1× |
| Lightning | 2650443 | 732.13 MB/s | 2846868 | 2238 | 8.9× |
| SonicFastest | 4587373 | 423.00 MB/s | 14608603 | 1407 | 5.1× |
| Sonic | 4672511 | 415.30 MB/s | 14608567 | 1407 | 5.0× |
| Goccy | 4758329 | 407.81 MB/s | 4065490 | 13510 | 4.9× |
| Easyjson | 7543011 | 257.25 MB/s | 3871265 | 15043 | 3.1× |
| LightningDecodeAny | 9598607 | 202.16 MB/s | 7013523 | 219937 | 2.4× |
| JSONV2 | 11285372 | 171.95 MB/s | 3237224 | 13947 | 2.1× |
| Stdlib | 23468061 | 82.69 MB/s | 3551323 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1107515 | 3004.77 MB/s | 351704 | 1286 | 18.8× |
| Lightning | 1771135 | 1878.93 MB/s | 2488906 | 2995 | 11.8× |
| SonicFastest | 2659138 | 1251.47 MB/s | 6439366 | 4248 | 7.8× |
| Sonic | 2673451 | 1244.77 MB/s | 6502009 | 4248 | 7.8× |
| LightningDecodeAny | 3707250 | 829.12 MB/s | 4886620 | 56892 | 5.6× |
| Goccy | 4564568 | 729.06 MB/s | 3948908 | 3816 | 4.6× |
| JSONV2 | 7437088 | 447.46 MB/s | 5364515 | 13243 | 2.8× |
| Stdlib | 20826280 | 159.79 MB/s | 5565607 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 230358 | 956.54 MB/s | 136896 | 228 | 8.9× |
| LightningDestructive | 230790 | 954.75 MB/s | 136896 | 228 | 8.8× |
| Sonic | 376479 | 585.28 MB/s | 299526 | 398 | 5.4× |
| SonicFastest | 382826 | 575.58 MB/s | 312428 | 398 | 5.3× |
| Goccy | 431393 | 510.78 MB/s | 364873 | 1067 | 4.7× |
| Easyjson | 550317 | 400.40 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 722336 | 305.05 MB/s | 129741 | 470 | 2.8× |
| LightningDecodeAny | 889359 | 121.79 MB/s | 861394 | 11944 | 2.3× |
| Stdlib | 2039682 | 108.03 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 12722877 | 636.65 MB/s | 15059832 | 51649 | 7.0× |
| Lightning | 13276516 | 610.10 MB/s | 15059842 | 51649 | 6.7× |
| Sonic | 16617117 | 487.45 MB/s | 70887515 | 40014 | 5.4× |
| SonicFastest | 16905232 | 479.14 MB/s | 70901722 | 40014 | 5.3× |
| Goccy | 23642981 | 342.60 MB/s | 16846998 | 107147 | 3.8× |
| Easyjson | 30561391 | 265.04 MB/s | 15059633 | 41643 | 2.9× |
| LightningDecodeAny | 40183976 | 129.48 MB/s | 34333351 | 912810 | 2.2× |
| JSONV2 | 43339744 | 186.90 MB/s | 15233725 | 78972 | 2.1× |
| Stdlib | 89331782 | 90.67 MB/s | 15665066 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 6069584 | 491.54 MB/s | 3985336 | 30015 | 7.8× |
| Lightning | 6149441 | 485.16 MB/s | 3985336 | 30015 | 7.7× |
| SonicFastest | 8642535 | 345.21 MB/s | 26582500 | 56760 | 5.4× |
| Sonic | 8652977 | 344.79 MB/s | 26492996 | 56760 | 5.4× |
| Goccy | 16463003 | 181.22 MB/s | 10678181 | 273650 | 2.9× |
| Easyjson | 16572971 | 180.02 MB/s | 9479440 | 30115 | 2.8× |
| LightningDecodeAny | 18663662 | 98.28 MB/s | 20023838 | 408557 | 2.5× |
| JSONV2 | 24604957 | 121.25 MB/s | 9257158 | 86278 | 1.9× |
| Stdlib | 47083515 | 63.37 MB/s | 9258093 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1361381 | 531.52 MB/s | 907393 | 3618 | 8.5× |
| Lightning | 1378496 | 524.92 MB/s | 907388 | 3618 | 8.4× |
| Sonic | 1759352 | 411.29 MB/s | 3196874 | 7226 | 6.6× |
| SonicFastest | 1768563 | 409.14 MB/s | 3183077 | 7226 | 6.5× |
| Easyjson | 4271196 | 169.41 MB/s | 2847907 | 3698 | 2.7× |
| LightningDecodeAny | 4322587 | 150.51 MB/s | 5752201 | 80172 | 2.7× |
| Goccy | 4736419 | 152.77 MB/s | 2772723 | 80272 | 2.4× |
| JSONV2 | 5679315 | 127.41 MB/s | 2704594 | 7318 | 2.0× |
| Stdlib | 11555933 | 62.62 MB/s | 2704553 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 1916464 | 823.05 MB/s | 907393 | 3618 | 8.2× |
| Lightning | 1966951 | 801.93 MB/s | 907387 | 3618 | 8.0× |
| SonicFastest | 2249385 | 701.24 MB/s | 5785236 | 7226 | 7.0× |
| Sonic | 2254462 | 699.66 MB/s | 5790894 | 7226 | 7.0× |
| LightningDecodeAny | 3811421 | 197.67 MB/s | 5752201 | 80172 | 4.1× |
| Easyjson | 5671902 | 278.10 MB/s | 2847906 | 3698 | 2.8× |
| Goccy | 5718526 | 275.83 MB/s | 3596708 | 80267 | 2.8× |
| JSONV2 | 6532144 | 241.48 MB/s | 2704593 | 7318 | 2.4× |
| Stdlib | 15750302 | 100.15 MB/s | 2704550 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 223221 | 672.53 MB/s | 81920 | 1 | 8.2× |
| Lightning | 223861 | 670.61 MB/s | 81920 | 1 | 8.1× |
| SonicFastest | 269951 | 556.12 MB/s | 253954 | 6 | 6.7× |
| Sonic | 271415 | 553.12 MB/s | 258987 | 6 | 6.7× |
| LightningDecodeAny | 474367 | 316.47 MB/s | 746005 | 10020 | 3.8× |
| Goccy | 857720 | 175.03 MB/s | 324570 | 10004 | 2.1× |
| JSONV2 | 1093066 | 137.34 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1821449 | 82.42 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 35376 | 794.80 MB/s | 30144 | 103 | 8.6× |
| Lightning | 35473 | 792.63 MB/s | 30272 | 105 | 8.6× |
| Sonic | 63533 | 442.56 MB/s | 46926 | 103 | 4.8× |
| SonicFastest | 63585 | 442.19 MB/s | 47039 | 103 | 4.8× |
| Easyjson | 68831 | 408.49 MB/s | 32304 | 138 | 4.4× |
| Goccy | 70648 | 397.99 MB/s | 59190 | 188 | 4.3× |
| JSONV2 | 134769 | 208.63 MB/s | 36896 | 242 | 2.3× |
| LightningDecodeAny | 150910 | 186.32 MB/s | 135024 | 2678 | 2.0× |
| Stdlib | 303704 | 92.58 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1958 | 1189.18 MB/s | 32 | 1 | 11.6× |
| LightningDestructive | 2010 | 1158.36 MB/s | 32 | 1 | 11.3× |
| Goccy | 4192 | 555.34 MB/s | 3649 | 4 | 5.4× |
| Easyjson | 4219 | 551.84 MB/s | 192 | 2 | 5.4× |
| SonicFastest | 5140 | 452.95 MB/s | 4491 | 6 | 4.4× |
| Sonic | 5143 | 452.62 MB/s | 4463 | 6 | 4.4× |
| JSONV2 | 8498 | 273.96 MB/s | 1000 | 6 | 2.7× |
| LightningDecodeAny | 10561 | 159.55 MB/s | 9960 | 195 | 2.1× |
| Stdlib | 22698 | 102.56 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 220 | 860.55 MB/s | 0 | 0 | 11.2× |
| LightningDestructive | 221 | 853.93 MB/s | 0 | 0 | 11.1× |
| Goccy | 380 | 497.52 MB/s | 304 | 2 | 6.5× |
| Easyjson | 492 | 384.19 MB/s | 0 | 0 | 5.0× |
| Sonic | 790 | 239.11 MB/s | 497 | 4 | 3.1× |
| SonicFastest | 792 | 238.76 MB/s | 496 | 4 | 3.1× |
| JSONV2 | 1032 | 183.19 MB/s | 112 | 1 | 2.4× |
| LightningDecodeAny | 1221 | 109.76 MB/s | 1160 | 25 | 2.0× |
| Stdlib | 2451 | 77.10 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1519 | 1442.15 MB/s | 0 | 0 | 10.6× |
| LightningDestructive | 1540 | 1423.06 MB/s | 0 | 0 | 10.4× |
| Goccy | 3189 | 687.13 MB/s | 2864 | 4 | 5.0× |
| Easyjson | 3190 | 686.75 MB/s | 24 | 1 | 5.0× |
| Sonic | 6426 | 340.96 MB/s | 3901 | 40 | 2.5× |
| SonicFastest | 6449 | 339.73 MB/s | 3931 | 40 | 2.5× |
| JSONV2 | 8124 | 269.71 MB/s | 640 | 6 | 2.0× |
| LightningDecodeAny | 8149 | 222.24 MB/s | 7536 | 158 | 2.0× |
| Stdlib | 16063 | 136.40 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 694486 | 735.04 MB/s | 703778 | 1012 | 8.7× |
| Lightning | 727778 | 701.42 MB/s | 703780 | 1012 | 8.3× |
| SonicFastest | 1148709 | 444.39 MB/s | 889718 | 2006 | 5.3× |
| Goccy | 1151081 | 443.48 MB/s | 1139859 | 5006 | 5.3× |
| Sonic | 1152210 | 443.04 MB/s | 898322 | 2006 | 5.3× |
| Easyjson | 1531128 | 333.40 MB/s | 863778 | 3012 | 4.0× |
| JSONV2 | 3229483 | 158.07 MB/s | 1076013 | 12646 | 1.9× |
| LightningDecodeAny | 3511527 | 131.41 MB/s | 2785927 | 66022 | 1.7× |
| Stdlib | 6065345 | 84.16 MB/s | 1162118 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1335 | 14826.28 MB/s | 0 | 0 | 83.7× |
| LightningDestructive | 1356 | 14589.22 MB/s | 0 | 0 | 82.4× |
| Goccy | 19866 | 996.11 MB/s | 20491 | 2 | 5.6× |
| SonicFastest | 27083 | 730.68 MB/s | 22200 | 4 | 4.1× |
| Sonic | 27089 | 730.51 MB/s | 22170 | 4 | 4.1× |
| JSONV2 | 29757 | 665.03 MB/s | 8 | 1 | 3.8× |
| LightningDecodeAny | 77590 | 255.03 MB/s | 117104 | 2019 | 1.4× |
| Easyjson | 86031 | 230.02 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111771 | 177.05 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2653 | 6832.75 MB/s | 0 | 0 | 38.7× |
| Lightning | 2828 | 6408.28 MB/s | 432 | 2 | 36.3× |
| Easyjson | 3954 | 4583.88 MB/s | 432 | 2 | 26.0× |
| SonicFastest | 9950 | 1821.45 MB/s | 23055 | 6 | 10.3× |
| Sonic | 10184 | 1779.61 MB/s | 23589 | 6 | 10.1× |
| Goccy | 15474 | 1171.29 MB/s | 19459 | 2 | 6.6× |
| LightningDecodeAny | 16523 | 1082.25 MB/s | 29088 | 191 | 6.2× |
| JSONV2 | 45414 | 399.08 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 102674 | 176.52 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 2280323 | 880.79 MB/s | 2960389 | 7411 | 8.2× |
| Lightning | 2329509 | 862.20 MB/s | 2962101 | 7417 | 8.0× |
| Goccy | 4245507 | 473.09 MB/s | 5411463 | 15830 | 4.4× |
| Sonic | 4396250 | 456.87 MB/s | 10953171 | 13683 | 4.2× |
| SonicFastest | 4457551 | 450.58 MB/s | 10925049 | 13683 | 4.2× |
| Easyjson | 4903183 | 409.63 MB/s | 2981487 | 7439 | 3.8× |
| JSONV2 | 6897810 | 291.18 MB/s | 3173685 | 14563 | 2.7× |
| LightningDecodeAny | 7212326 | 158.38 MB/s | 7386073 | 134751 | 2.6× |
| Stdlib | 18627616 | 107.82 MB/s | 3589319 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 886 | 619.42 MB/s | 480 | 1 | 6.4× |
| LightningDestructive | 896 | 612.40 MB/s | 480 | 1 | 6.3× |
| LightningDecodeAny | 2004 | 273.43 MB/s | 2261 | 50 | 2.8× |
| Easyjson | 2209 | 248.53 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2710 | 202.58 MB/s | 2048 | 26 | 2.1× |
| Sonic | 2723 | 201.62 MB/s | 2062 | 26 | 2.1× |
| Goccy | 3067 | 178.99 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3342 | 164.26 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5691 | 96.47 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 495905 | 1273.46 MB/s | 364472 | 566 | 10.9× |
| Lightning | 550630 | 1146.89 MB/s | 413001 | 878 | 9.8× |
| Sonic | 1004674 | 628.58 MB/s | 997510 | 1102 | 5.4× |
| SonicFastest | 1009205 | 625.75 MB/s | 1001287 | 1102 | 5.3× |
| Easyjson | 1137117 | 555.36 MB/s | 422504 | 936 | 4.7× |
| Goccy | 1169219 | 540.12 MB/s | 986304 | 1201 | 4.6× |
| JSONV2 | 2152101 | 293.44 MB/s | 571618 | 3144 | 2.5× |
| LightningDecodeAny | 2363688 | 197.53 MB/s | 2010200 | 30295 | 2.3× |
| Stdlib | 5392890 | 117.10 MB/s | 654667 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 714269 | 787.39 MB/s | 544252 | 448 | 7.4× |
| Lightning | 894089 | 629.03 MB/s | 767622 | 1254 | 5.9× |
| SonicFastest | 1022690 | 549.93 MB/s | 930835 | 1476 | 5.2× |
| Sonic | 1022744 | 549.90 MB/s | 927267 | 1476 | 5.2× |
| Goccy | 1346337 | 417.73 MB/s | 1039983 | 1030 | 3.9× |
| Easyjson | 1747701 | 321.80 MB/s | 775154 | 1254 | 3.0× |
| LightningDecodeAny | 2712047 | 207.37 MB/s | 2114152 | 30295 | 1.9× |
| JSONV2 | 2753820 | 204.23 MB/s | 927450 | 3482 | 1.9× |
| Stdlib | 5267910 | 106.76 MB/s | 1011671 | 6763 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| LightningDestructive | 655023 | 813.98 MB/s | 333416 | 2084 | 8.4× |
| Lightning | 671067 | 794.52 MB/s | 368224 | 2293 | 8.2× |
| Easyjson | 1106231 | 481.98 MB/s | 428361 | 3273 | 4.9× |
| SonicFastest | 1151272 | 463.12 MB/s | 1047053 | 4351 | 4.8× |
| Sonic | 1153823 | 462.10 MB/s | 1051879 | 4351 | 4.7× |
| Goccy | 1277712 | 417.29 MB/s | 1167226 | 5409 | 4.3× |
| JSONV2 | 2526841 | 211.01 MB/s | 745451 | 13288 | 2.2× |
| LightningDecodeAny | 3361175 | 158.63 MB/s | 2674620 | 50596 | 1.6× |
| Stdlib | 5472370 | 97.43 MB/s | 798692 | 17133 | 1.0× |
