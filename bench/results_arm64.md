# JSON Deserialization Benchmarks

- generated 2026-06-20T13:10:16Z
- go version go1.26.4 linux/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/apache_builds — 127275 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 179637 | 708.51 MB/s | 186020 | 10 | 6.2× |
| SonicFastest | 180140 | 706.53 MB/s | 187272 | 10 | 6.1× |
| Lightning | 192629 | 660.73 MB/s | 50208 | 4 | 5.7× |
| Goccy | 193310 | 658.40 MB/s | 224785 | 884 | 5.7× |
| Easyjson | 211942 | 600.52 MB/s | 122864 | 14 | 5.2× |
| JSONV2 | 420363 | 302.77 MB/s | 195117 | 1805 | 2.6× |
| Stdlib | 1106591 | 115.02 MB/s | 199697 | 2661 | 1.0× |

## bench/canada — 2251051 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 3977654 | 565.92 MB/s | 3122879 | 3081 | 6.6× |
| Sonic | 4598027 | 489.57 MB/s | 15233761 | 970 | 5.7× |
| SonicFastest | 4612585 | 488.02 MB/s | 15232103 | 970 | 5.7× |
| Goccy | 10338344 | 217.74 MB/s | 4277294 | 56541 | 2.5× |
| Easyjson | 11217807 | 200.67 MB/s | 3099809 | 2120 | 2.3× |
| JSONV2 | 16129556 | 139.56 MB/s | 3123370 | 3083 | 1.6× |
| Stdlib | 26227608 | 85.83 MB/s | 3123395 | 3095 | 1.0× |

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 484518 | 558.09 MB/s | 348024 | 1627 | 6.9× |
| Sonic | 642602 | 420.79 MB/s | 501996 | 968 | 5.2× |
| SonicFastest | 644320 | 419.67 MB/s | 507641 | 968 | 5.2× |
| Goccy | 1394554 | 193.90 MB/s | 541291 | 8122 | 2.4× |
| Easyjson | 1414251 | 191.20 MB/s | 330272 | 749 | 2.4× |
| JSONV2 | 2085556 | 129.66 MB/s | 348150 | 1628 | 1.6× |
| Stdlib | 3350470 | 80.71 MB/s | 348544 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1477285 | 1169.17 MB/s | 964644 | 6177 | 9.1× |
| SonicFastest | 2079149 | 830.73 MB/s | 2928008 | 4020 | 6.5× |
| Sonic | 2082481 | 829.40 MB/s | 2983297 | 4020 | 6.5× |
| Goccy | 2450208 | 704.92 MB/s | 2587754 | 14607 | 5.5× |
| Easyjson | 4287482 | 402.85 MB/s | 972033 | 5389 | 3.1× |
| JSONV2 | 4299695 | 401.70 MB/s | 1011703 | 7594 | 3.1× |
| Stdlib | 13432747 | 128.58 MB/s | 1234450 | 17027 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1238 | 1463.49 MB/s | 0 | 0 | 11.3× |
| Easyjson | 2551 | 710.41 MB/s | 24 | 1 | 5.5× |
| Goccy | 2784 | 650.84 MB/s | 2608 | 4 | 5.0× |
| SonicFastest | 5915 | 306.33 MB/s | 3754 | 40 | 2.4× |
| Sonic | 5927 | 305.69 MB/s | 3788 | 40 | 2.4× |
| JSONV2 | 7782 | 232.83 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14031 | 129.15 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1253 | 1446.28 MB/s | 0 | 0 | 11.2× |
| Easyjson | 2560 | 707.90 MB/s | 24 | 1 | 5.5× |
| Goccy | 2759 | 656.78 MB/s | 2608 | 4 | 5.1× |
| SonicFastest | 5976 | 303.22 MB/s | 3764 | 40 | 2.4× |
| Sonic | 5980 | 302.99 MB/s | 3793 | 40 | 2.4× |
| JSONV2 | 7808 | 232.08 MB/s | 640 | 6 | 1.8× |
| Stdlib | 14076 | 128.73 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1443 | 1255.75 MB/s | 144 | 10 | 9.7× |
| Easyjson | 2782 | 651.29 MB/s | 144 | 10 | 5.0× |
| Goccy | 2847 | 636.41 MB/s | 2600 | 5 | 4.9× |
| Sonic | 6048 | 299.59 MB/s | 3749 | 42 | 2.3× |
| SonicFastest | 6088 | 297.65 MB/s | 3757 | 42 | 2.3× |
| JSONV2 | 7975 | 227.20 MB/s | 632 | 7 | 1.8× |
| Stdlib | 14012 | 129.32 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 774 | 638.68 MB/s | 160 | 1 | 7.2× |
| Sonic | 1244 | 397.24 MB/s | 988 | 6 | 4.5× |
| SonicFastest | 1245 | 396.71 MB/s | 986 | 6 | 4.5× |
| Easyjson | 2219 | 222.64 MB/s | 448 | 3 | 2.5× |
| Goccy | 2444 | 202.10 MB/s | 856 | 23 | 2.3× |
| JSONV2 | 3179 | 155.39 MB/s | 528 | 7 | 1.8× |
| Stdlib | 5564 | 88.78 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 456 | 503.93 MB/s | 160 | 1 | 9.1× |
| Sonic | 884 | 260.31 MB/s | 656 | 6 | 4.7× |
| SonicFastest | 887 | 259.37 MB/s | 663 | 6 | 4.7× |
| Easyjson | 1401 | 164.21 MB/s | 448 | 3 | 3.0× |
| Goccy | 1623 | 141.71 MB/s | 584 | 23 | 2.6× |
| JSONV2 | 2359 | 97.51 MB/s | 528 | 7 | 1.8× |
| Stdlib | 4173 | 55.12 MB/s | 760 | 12 | 1.0× |

## bench/github_events — 65132 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 80447 | 809.62 MB/s | 178609 | 112 | 6.8× |
| SonicFastest | 97911 | 665.21 MB/s | 157712 | 75 | 5.6× |
| Sonic | 98276 | 662.75 MB/s | 158825 | 75 | 5.5× |
| Goccy | 140915 | 462.21 MB/s | 229047 | 134 | 3.9× |
| JSONV2 | 221831 | 293.61 MB/s | 206651 | 607 | 2.5× |
| Stdlib | 544811 | 119.55 MB/s | 214617 | 842 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2931324 | 661.98 MB/s | 2846870 | 2238 | 8.1× |
| SonicFastest | 4488512 | 432.32 MB/s | 14608575 | 1407 | 5.3× |
| Sonic | 4495335 | 431.66 MB/s | 14608596 | 1407 | 5.2× |
| Goccy | 4825654 | 402.12 MB/s | 4064990 | 13510 | 4.9× |
| Easyjson | 7694274 | 252.20 MB/s | 3871269 | 15043 | 3.1× |
| JSONV2 | 11230793 | 172.78 MB/s | 3237332 | 13947 | 2.1× |
| Stdlib | 23597849 | 82.23 MB/s | 3551324 | 27166 | 1.0× |

## bench/gsoc_2018 — 3327831 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2713183 | 1226.54 MB/s | 4611524 | 5958 | 7.7× |
| Sonic | 2717013 | 1224.81 MB/s | 6769092 | 4249 | 7.7× |
| SonicFastest | 2730041 | 1218.97 MB/s | 6718017 | 4249 | 7.7× |
| Goccy | 4513365 | 737.33 MB/s | 3948956 | 3817 | 4.6× |
| JSONV2 | 7567792 | 439.74 MB/s | 5364587 | 13243 | 2.8× |
| Stdlib | 20959398 | 158.78 MB/s | 5565614 | 20690 | 1.0× |

## bench/instruments — 220346 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 277339 | 794.50 MB/s | 96320 | 224 | 7.3× |
| Sonic | 373681 | 589.66 MB/s | 294456 | 398 | 5.5× |
| SonicFastest | 375578 | 586.69 MB/s | 296790 | 398 | 5.4× |
| Goccy | 422012 | 522.13 MB/s | 362478 | 1066 | 4.8× |
| Easyjson | 550514 | 400.26 MB/s | 130512 | 245 | 3.7× |
| JSONV2 | 732580 | 300.78 MB/s | 129738 | 470 | 2.8× |
| Stdlib | 2037738 | 108.13 MB/s | 131240 | 619 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 13832527 | 585.58 MB/s | 15059836 | 51649 | 6.5× |
| SonicFastest | 17101056 | 473.66 MB/s | 71953074 | 40015 | 5.3× |
| Sonic | 17102229 | 473.62 MB/s | 71996728 | 40015 | 5.3× |
| Goccy | 24242348 | 334.13 MB/s | 17336842 | 107151 | 3.7× |
| Easyjson | 31140636 | 260.11 MB/s | 15059632 | 41643 | 2.9× |
| JSONV2 | 44222833 | 183.16 MB/s | 15233841 | 78972 | 2.0× |
| Stdlib | 89950696 | 90.05 MB/s | 15665082 | 150647 | 1.0× |

## bench/marine_ik — 2983466 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 6405025 | 465.80 MB/s | 3985341 | 30015 | 7.4× |
| SonicFastest | 8396705 | 355.31 MB/s | 25601214 | 56758 | 5.7× |
| Sonic | 8420068 | 354.33 MB/s | 25545038 | 56758 | 5.7× |
| Goccy | 16784132 | 177.76 MB/s | 10955854 | 273672 | 2.8× |
| Easyjson | 17116287 | 174.31 MB/s | 9479441 | 30115 | 2.8× |
| JSONV2 | 25531244 | 116.86 MB/s | 9257241 | 86279 | 1.9× |
| Stdlib | 47674707 | 62.58 MB/s | 9258100 | 86317 | 1.0× |

## bench/mesh — 723597 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1360595 | 531.82 MB/s | 696698 | 3610 | 8.5× |
| Sonic | 1714608 | 422.02 MB/s | 3199818 | 7226 | 6.8× |
| SonicFastest | 1716977 | 421.44 MB/s | 3214224 | 7226 | 6.8× |
| Easyjson | 4275029 | 169.26 MB/s | 2847909 | 3698 | 2.7× |
| Goccy | 4866502 | 148.69 MB/s | 3176172 | 80292 | 2.4× |
| JSONV2 | 5906829 | 122.50 MB/s | 2704783 | 7319 | 2.0× |
| Stdlib | 11607437 | 62.34 MB/s | 2704551 | 7324 | 1.0× |

## bench/mesh_pretty — 1577353 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2044332 | 771.57 MB/s | 696698 | 3610 | 7.7× |
| SonicFastest | 2287323 | 689.61 MB/s | 5794004 | 7226 | 6.9× |
| Sonic | 2294954 | 687.31 MB/s | 5802716 | 7226 | 6.9× |
| Goccy | 5670868 | 278.15 MB/s | 3678006 | 80273 | 2.8× |
| Easyjson | 5694370 | 277.00 MB/s | 2847911 | 3698 | 2.8× |
| JSONV2 | 6540649 | 241.16 MB/s | 2704769 | 7319 | 2.4× |
| Stdlib | 15822782 | 99.69 MB/s | 2704555 | 7324 | 1.0× |

## bench/numbers — 150124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 224377 | 669.07 MB/s | 81920 | 1 | 8.2× |
| Sonic | 272153 | 551.62 MB/s | 258100 | 6 | 6.8× |
| SonicFastest | 276053 | 543.82 MB/s | 276117 | 6 | 6.7× |
| Goccy | 856573 | 175.26 MB/s | 322809 | 10004 | 2.1× |
| JSONV2 | 1089580 | 137.78 MB/s | 357714 | 20 | 1.7× |
| Stdlib | 1841328 | 81.53 MB/s | 357800 | 22 | 1.0× |

## bench/payload_large — 28117 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 38156 | 736.89 MB/s | 29104 | 107 | 8.0× |
| SonicFastest | 63653 | 441.73 MB/s | 47085 | 103 | 4.8× |
| Sonic | 63719 | 441.27 MB/s | 47380 | 103 | 4.8× |
| Easyjson | 68483 | 410.57 MB/s | 32304 | 138 | 4.4× |
| Goccy | 69823 | 402.69 MB/s | 59201 | 188 | 4.4× |
| JSONV2 | 134134 | 209.62 MB/s | 36895 | 242 | 2.3× |
| Stdlib | 303902 | 92.52 MB/s | 43968 | 513 | 1.0× |

## bench/payload_medium — 2328 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2041 | 1140.75 MB/s | 32 | 1 | 11.2× |
| Goccy | 4090 | 569.22 MB/s | 3649 | 4 | 5.6× |
| Easyjson | 4269 | 545.36 MB/s | 192 | 2 | 5.3× |
| Sonic | 5033 | 462.57 MB/s | 4151 | 6 | 4.5× |
| SonicFastest | 5068 | 459.38 MB/s | 4208 | 6 | 4.5× |
| JSONV2 | 8474 | 274.73 MB/s | 1000 | 6 | 2.7× |
| Stdlib | 22800 | 102.11 MB/s | 2288 | 46 | 1.0× |

## bench/payload_small — 189 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 229 | 824.72 MB/s | 0 | 0 | 10.9× |
| Goccy | 388 | 486.63 MB/s | 304 | 2 | 6.4× |
| Easyjson | 490 | 385.65 MB/s | 0 | 0 | 5.1× |
| SonicFastest | 809 | 233.65 MB/s | 519 | 4 | 3.1× |
| Sonic | 810 | 233.28 MB/s | 515 | 4 | 3.1× |
| JSONV2 | 1045 | 180.80 MB/s | 112 | 1 | 2.4× |
| Stdlib | 2489 | 75.94 MB/s | 416 | 9 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1550 | 1413.71 MB/s | 0 | 0 | 10.3× |
| Goccy | 3131 | 699.67 MB/s | 2864 | 4 | 5.1× |
| Easyjson | 3225 | 679.32 MB/s | 24 | 1 | 5.0× |
| SonicFastest | 6317 | 346.83 MB/s | 3959 | 40 | 2.5× |
| Sonic | 6342 | 345.47 MB/s | 4001 | 40 | 2.5× |
| JSONV2 | 8023 | 273.07 MB/s | 640 | 6 | 2.0× |
| Stdlib | 15985 | 137.07 MB/s | 928 | 16 | 1.0× |

## bench/random — 510476 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 885062 | 576.77 MB/s | 703777 | 1012 | 7.0× |
| Sonic | 1162315 | 439.19 MB/s | 924987 | 2006 | 5.3× |
| SonicFastest | 1178821 | 433.04 MB/s | 967801 | 2006 | 5.3× |
| Goccy | 1184238 | 431.06 MB/s | 1141859 | 5007 | 5.2× |
| Easyjson | 1562953 | 326.61 MB/s | 863778 | 3012 | 4.0× |
| JSONV2 | 3229540 | 158.06 MB/s | 1076001 | 12646 | 1.9× |
| Stdlib | 6189826 | 82.47 MB/s | 1162116 | 16023 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2062 | 9595.90 MB/s | 0 | 0 | 54.2× |
| Goccy | 19899 | 994.47 MB/s | 20491 | 2 | 5.6× |
| Sonic | 27900 | 709.28 MB/s | 21860 | 4 | 4.0× |
| SonicFastest | 28239 | 700.77 MB/s | 22620 | 4 | 4.0× |
| JSONV2 | 29575 | 669.10 MB/s | 8 | 1 | 3.8× |
| Easyjson | 86040 | 230.00 MB/s | 0 | 0 | 1.3× |
| Stdlib | 111775 | 177.04 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2981 | 6080.63 MB/s | 864 | 4 | 34.7× |
| Easyjson | 3972 | 4562.89 MB/s | 432 | 2 | 26.0× |
| Sonic | 9889 | 1832.68 MB/s | 23197 | 6 | 10.5× |
| SonicFastest | 9900 | 1830.63 MB/s | 22996 | 6 | 10.4× |
| Goccy | 15627 | 1159.79 MB/s | 19459 | 2 | 6.6× |
| JSONV2 | 45607 | 397.39 MB/s | 16499 | 50 | 2.3× |
| Stdlib | 103409 | 175.26 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2583338 | 777.48 MB/s | 2963985 | 7453 | 7.3× |
| Sonic | 4087069 | 491.43 MB/s | 10028173 | 13682 | 4.6× |
| SonicFastest | 4119979 | 487.50 MB/s | 10036644 | 13682 | 4.6× |
| Goccy | 4281145 | 469.15 MB/s | 5415983 | 15844 | 4.4× |
| Easyjson | 5116747 | 392.53 MB/s | 2981724 | 7441 | 3.7× |
| JSONV2 | 7065673 | 284.26 MB/s | 3173807 | 14563 | 2.7× |
| Stdlib | 18803710 | 106.81 MB/s | 3589321 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 886 | 619.92 MB/s | 480 | 1 | 6.4× |
| Easyjson | 2134 | 257.28 MB/s | 1616 | 5 | 2.6× |
| SonicFastest | 2684 | 204.56 MB/s | 1942 | 26 | 2.1× |
| Sonic | 2687 | 204.29 MB/s | 1933 | 26 | 2.1× |
| Goccy | 3047 | 180.16 MB/s | 2128 | 43 | 1.9× |
| JSONV2 | 3262 | 168.28 MB/s | 1664 | 7 | 1.7× |
| Stdlib | 5641 | 97.33 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 569244 | 1109.39 MB/s | 461113 | 1230 | 9.4× |
| SonicFastest | 1011598 | 624.27 MB/s | 1010746 | 1102 | 5.3× |
| Sonic | 1016683 | 621.15 MB/s | 1021870 | 1102 | 5.3× |
| Easyjson | 1128944 | 559.38 MB/s | 422504 | 936 | 4.8× |
| Goccy | 1183913 | 533.41 MB/s | 986905 | 1202 | 4.5× |
| JSONV2 | 2159644 | 292.42 MB/s | 571656 | 3144 | 2.5× |
| Stdlib | 5371789 | 117.56 MB/s | 654665 | 6472 | 1.0× |

## bench/twitterescaped — 562408 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| SonicFastest | 1118719 | 502.73 MB/s | 1047880 | 2506 | 4.7× |
| Sonic | 1123092 | 500.77 MB/s | 1056485 | 2506 | 4.7× |
| Goccy | 1527141 | 368.28 MB/s | 1103875 | 3406 | 3.4× |
| Lightning | 1814935 | 309.88 MB/s | 945693 | 4207 | 2.9× |
| Easyjson | 1815750 | 309.74 MB/s | 833618 | 3391 | 2.9× |
| JSONV2 | 2886571 | 194.84 MB/s | 979919 | 5184 | 1.8× |
| Stdlib | 5251170 | 107.10 MB/s | 1067141 | 8759 | 1.0× |

## bench/update_center — 533178 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1327397 | 401.67 MB/s | 935534 | 16692 | 4.2× |
| SonicFastest | 1479591 | 360.35 MB/s | 1506534 | 10471 | 3.8× |
| Sonic | 1481397 | 359.92 MB/s | 1502056 | 10471 | 3.8× |
| Easyjson | 1777750 | 299.92 MB/s | 949523 | 15985 | 3.2× |
| Goccy | 2028737 | 262.81 MB/s | 1669058 | 18051 | 2.8× |
| JSONV2 | 3109649 | 171.46 MB/s | 1192599 | 22667 | 1.8× |
| Stdlib | 5611706 | 95.01 MB/s | 1245604 | 26510 | 1.0× |
