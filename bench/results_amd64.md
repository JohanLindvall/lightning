# JSON Deserialization Benchmarks

- generated 2026-06-16T19:25:59Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 846 | 2141.29 MB/s | 144 | 10 | 10.0× |
| Easyjson | 1726 | 1049.97 MB/s | 144 | 10 | 4.9× |
| Sonic | 4504 | 402.31 MB/s | 3375 | 40 | 1.9× |
| Stdlib | 8445 | 214.56 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 662 | 2738.60 MB/s | 0 | 0 | 12.9× |
| Easyjson | 1601 | 1132.09 MB/s | 24 | 1 | 5.3× |
| Sonic | 4271 | 424.28 MB/s | 3360 | 38 | 2.0× |
| Stdlib | 8555 | 211.79 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 675 | 2682.72 MB/s | 0 | 0 | 12.8× |
| Easyjson | 1590 | 1139.51 MB/s | 24 | 1 | 5.5× |
| Sonic | 4416 | 410.36 MB/s | 3359 | 38 | 2.0× |
| Stdlib | 8672 | 208.94 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 280 | 820.94 MB/s | 160 | 1 | 9.2× |
| Sonic | 653 | 352.18 MB/s | 805 | 8 | 4.0× |
| Easyjson | 974 | 236.05 MB/s | 448 | 3 | 2.7× |
| Stdlib | 2583 | 89.05 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 817 | 604.67 MB/s | 160 | 1 | 4.4× |
| Sonic | 912 | 541.82 MB/s | 1082 | 8 | 4.0× |
| Easyjson | 1472 | 335.62 MB/s | 448 | 3 | 2.5× |
| Stdlib | 3610 | 136.83 MB/s | 760 | 12 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Sonic | 14428492 | 561.39 MB/s | 23116096 | 175893 | 4.5× |
| Lightning | 14860773 | 545.06 MB/s | 5693903 | 107137 | 4.4× |
| Easyjson | 23268968 | 348.10 MB/s | 19531264 | 128767 | 2.8× |
| Stdlib | 64913137 | 124.78 MB/s | 19438504 | 352029 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 864 | 2537.03 MB/s | 0 | 0 | 11.9× |
| Easyjson | 1968 | 1113.18 MB/s | 24 | 1 | 5.2× |
| Sonic | 4787 | 457.73 MB/s | 3613 | 38 | 2.1× |
| Stdlib | 10275 | 213.24 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 384 | 51474.40 MB/s | 0 | 0 | 223.2× |
| Sonic | 20585 | 961.32 MB/s | 20763 | 3 | 4.2× |
| Easyjson | 68600 | 288.47 MB/s | 0 | 0 | 1.3× |
| Stdlib | 85787 | 230.67 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 706 | 777.21 MB/s | 480 | 1 | 5.4× |
| Easyjson | 1404 | 391.02 MB/s | 1616 | 5 | 2.7× |
| Sonic | 1654 | 331.89 MB/s | 2271 | 8 | 2.3× |
| Stdlib | 3794 | 144.71 MB/s | 1896 | 12 | 1.0× |
