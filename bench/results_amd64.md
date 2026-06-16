# JSON Deserialization Benchmarks

- generated 2026-06-16T18:22:14Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 800 | 2264.58 MB/s | 144 | 10 | 10.3× |
| Easyjson | 1768 | 1025.00 MB/s | 144 | 10 | 4.7× |
| Stdlib | 8246 | 219.73 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 630 | 2876.37 MB/s | 0 | 0 | 13.5× |
| Easyjson | 1602 | 1130.91 MB/s | 24 | 1 | 5.3× |
| Stdlib | 8481 | 213.66 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 658 | 2753.07 MB/s | 0 | 0 | 12.9× |
| Easyjson | 1602 | 1131.35 MB/s | 24 | 1 | 5.3× |
| Stdlib | 8478 | 213.73 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 280 | 821.25 MB/s | 160 | 1 | 8.8× |
| Easyjson | 951 | 241.81 MB/s | 448 | 3 | 2.6× |
| Stdlib | 2474 | 92.96 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1256 | 393.45 MB/s | 160 | 1 | 2.8× |
| Easyjson | 1453 | 339.88 MB/s | 448 | 3 | 2.4× |
| Stdlib | 3494 | 141.38 MB/s | 760 | 12 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 812 | 2699.32 MB/s | 0 | 0 | 12.5× |
| Easyjson | 1975 | 1109.15 MB/s | 24 | 1 | 5.1× |
| Stdlib | 10152 | 215.83 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 375 | 52750.00 MB/s | 0 | 0 | 229.4× |
| Easyjson | 55016 | 359.69 MB/s | 0 | 0 | 1.6× |
| Stdlib | 86038 | 230.00 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 668 | 821.85 MB/s | 480 | 1 | 5.4× |
| Easyjson | 1331 | 412.45 MB/s | 1616 | 5 | 2.7× |
| Stdlib | 3611 | 152.05 MB/s | 1896 | 12 | 1.0× |
