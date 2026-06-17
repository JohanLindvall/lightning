# JSON Deserialization Benchmarks

- generated 2026-06-17T08:24:53Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/canada_geometry — 270403 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 519956 | 520.05 MB/s | 151227 | 495 | 5.0× |
| Sonic | 549201 | 492.36 MB/s | 643321 | 1147 | 4.7× |
| Easyjson | 1293540 | 209.04 MB/s | 330321 | 753 | 2.0× |
| JSONV2 | 1555996 | 173.78 MB/s | 348541 | 1628 | 1.7× |
| Stdlib | 2581123 | 104.76 MB/s | 348545 | 1641 | 1.0× |

## bench/citm_catalog — 1727204 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1552695 | 1112.39 MB/s | 979497 | 6803 | 5.6× |
| Sonic | 1601555 | 1078.45 MB/s | 2740016 | 5548 | 5.4× |
| Easyjson | 3086291 | 559.64 MB/s | 986875 | 6015 | 2.8× |
| JSONV2 | 3094835 | 558.09 MB/s | 1012484 | 7594 | 2.8× |
| Stdlib | 8634343 | 200.04 MB/s | 1234449 | 17027 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 928 | 1951.97 MB/s | 144 | 10 | 10.7× |
| Easyjson | 1857 | 975.97 MB/s | 144 | 10 | 5.3× |
| Sonic | 4331 | 418.39 MB/s | 3386 | 40 | 2.3× |
| JSONV2 | 5629 | 321.91 MB/s | 633 | 7 | 1.8× |
| Stdlib | 9913 | 182.80 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 624 | 2905.20 MB/s | 0 | 0 | 15.5× |
| Easyjson | 1687 | 1073.92 MB/s | 24 | 1 | 5.7× |
| Sonic | 4688 | 386.50 MB/s | 3372 | 38 | 2.1× |
| JSONV2 | 5401 | 335.47 MB/s | 640 | 6 | 1.8× |
| Stdlib | 9648 | 187.81 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 700 | 2587.28 MB/s | 0 | 0 | 13.7× |
| Easyjson | 1722 | 1052.44 MB/s | 24 | 1 | 5.6× |
| Sonic | 4521 | 400.78 MB/s | 3373 | 38 | 2.1× |
| JSONV2 | 5443 | 332.92 MB/s | 641 | 6 | 1.8× |
| Stdlib | 9565 | 189.44 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 352 | 653.68 MB/s | 160 | 1 | 8.8× |
| Sonic | 675 | 340.66 MB/s | 811 | 8 | 4.6× |
| Easyjson | 1108 | 207.64 MB/s | 448 | 3 | 2.8× |
| JSONV2 | 1653 | 139.10 MB/s | 528 | 7 | 1.9× |
| Stdlib | 3100 | 74.20 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 575 | 858.89 MB/s | 160 | 1 | 7.1× |
| Sonic | 939 | 526.26 MB/s | 1092 | 8 | 4.4× |
| Easyjson | 1682 | 293.78 MB/s | 448 | 3 | 2.4× |
| JSONV2 | 2189 | 225.71 MB/s | 528 | 7 | 1.9× |
| Stdlib | 4094 | 120.67 MB/s | 760 | 12 | 1.0× |

## bench/golang_source — 1940472 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 3246066 | 597.79 MB/s | 4883379 | 1736 | 5.5× |
| Lightning | 3376424 | 574.71 MB/s | 4282952 | 27849 | 5.2× |
| Easyjson | 5934022 | 327.01 MB/s | 4282949 | 27849 | 3.0× |
| JSONV2 | 7839626 | 247.52 MB/s | 3239409 | 13947 | 2.3× |
| Stdlib | 17696487 | 109.65 MB/s | 3551324 | 27166 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 13661500 | 592.91 MB/s | 23093855 | 175893 | 5.0× |
| Lightning | 13680962 | 592.07 MB/s | 14139267 | 107156 | 5.0× |
| Easyjson | 27773801 | 291.64 MB/s | 19531269 | 128767 | 2.5× |
| JSONV2 | 39502868 | 205.05 MB/s | 19028406 | 280356 | 1.7× |
| Stdlib | 68448439 | 118.34 MB/s | 19437619 | 352028 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 888 | 2467.23 MB/s | 0 | 0 | 12.6× |
| Easyjson | 2162 | 1013.51 MB/s | 24 | 1 | 5.2× |
| Sonic | 4803 | 456.18 MB/s | 3627 | 38 | 2.3× |
| JSONV2 | 5499 | 398.43 MB/s | 640 | 6 | 2.0× |
| Stdlib | 11217 | 195.33 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 519 | 38151.44 MB/s | 0 | 0 | 148.6× |
| JSONV2 | 22227 | 890.30 MB/s | 8 | 1 | 3.5× |
| Sonic | 27154 | 728.77 MB/s | 20716 | 3 | 2.8× |
| Easyjson | 46854 | 422.35 MB/s | 0 | 0 | 1.6× |
| Stdlib | 77061 | 256.80 MB/s | 240 | 6 | 1.0× |

## bench/string_unicode — 18124 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 5557 | 3261.20 MB/s | 18080 | 62 | 13.7× |
| Sonic | 5827 | 3110.38 MB/s | 20743 | 5 | 13.1× |
| Easyjson | 6894 | 2629.07 MB/s | 17648 | 60 | 11.1× |
| JSONV2 | 34676 | 522.66 MB/s | 16518 | 50 | 2.2× |
| Stdlib | 76350 | 237.38 MB/s | 19320 | 67 | 1.0× |

## bench/synthea_fhir — 2008494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 2532698 | 793.03 MB/s | 3587274 | 29254 | 5.2× |
| Sonic | 3257823 | 616.51 MB/s | 5160633 | 7085 | 4.0× |
| Easyjson | 4813197 | 417.29 MB/s | 3604078 | 29229 | 2.7× |
| JSONV2 | 5702668 | 352.20 MB/s | 3175481 | 14563 | 2.3× |
| Stdlib | 13137966 | 152.88 MB/s | 3589320 | 29340 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 842 | 652.04 MB/s | 480 | 1 | 5.3× |
| Easyjson | 1727 | 317.87 MB/s | 1616 | 5 | 2.6× |
| Sonic | 1764 | 311.29 MB/s | 2283 | 8 | 2.6× |
| JSONV2 | 2549 | 215.42 MB/s | 1665 | 7 | 1.8× |
| Stdlib | 4503 | 121.93 MB/s | 1896 | 12 | 1.0× |

## bench/twitter_status — 631514 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 629997 | 1002.41 MB/s | 634494 | 5544 | 6.4× |
| Sonic | 764485 | 826.06 MB/s | 1088008 | 814 | 5.3× |
| Easyjson | 1099370 | 574.43 MB/s | 591597 | 5205 | 3.7× |
| JSONV2 | 1616476 | 390.67 MB/s | 572286 | 3144 | 2.5× |
| Stdlib | 4037552 | 156.41 MB/s | 654689 | 6472 | 1.0× |
