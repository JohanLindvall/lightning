# JSON Deserialization Benchmarks

- generated 2026-06-16T20:26:41Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 838 | 2162.73 MB/s | 144 | 10 | 10.3× |
| Easyjson | 1726 | 1049.74 MB/s | 144 | 10 | 5.0× |
| Sonic | 4512 | 401.56 MB/s | 3366 | 40 | 1.9× |
| Stdlib | 8614 | 210.35 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 670 | 2704.00 MB/s | 0 | 0 | 13.0× |
| Easyjson | 1590 | 1139.38 MB/s | 24 | 1 | 5.5× |
| Sonic | 4325 | 418.97 MB/s | 3355 | 38 | 2.0× |
| Stdlib | 8742 | 207.28 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 691 | 2623.02 MB/s | 0 | 0 | 12.3× |
| Easyjson | 1598 | 1134.08 MB/s | 24 | 1 | 5.3× |
| Sonic | 4481 | 404.39 MB/s | 3358 | 38 | 1.9× |
| Stdlib | 8516 | 212.78 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 294 | 781.72 MB/s | 160 | 1 | 8.8× |
| Sonic | 636 | 361.57 MB/s | 806 | 8 | 4.1× |
| Easyjson | 972 | 236.62 MB/s | 448 | 3 | 2.7× |
| Stdlib | 2593 | 88.69 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 535 | 922.96 MB/s | 160 | 1 | 6.8× |
| Sonic | 930 | 531.33 MB/s | 1083 | 8 | 3.9× |
| Easyjson | 1487 | 332.27 MB/s | 448 | 3 | 2.4× |
| Stdlib | 3617 | 136.59 MB/s | 760 | 12 | 1.0× |

## bench/large-json — 8100039 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 11627043 | 696.66 MB/s | 14139127 | 107156 | 5.6× |
| Sonic | 15182625 | 533.51 MB/s | 23102291 | 175893 | 4.3× |
| Easyjson | 23295692 | 347.71 MB/s | 19531266 | 128767 | 2.8× |
| Stdlib | 65162897 | 124.30 MB/s | 19438214 | 352028 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 871 | 2516.34 MB/s | 0 | 0 | 12.0× |
| Easyjson | 1971 | 1111.87 MB/s | 24 | 1 | 5.3× |
| Sonic | 4779 | 458.43 MB/s | 3619 | 38 | 2.2× |
| Stdlib | 10409 | 210.50 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 387 | 51164.75 MB/s | 0 | 0 | 221.2× |
| Sonic | 20591 | 961.05 MB/s | 20753 | 3 | 4.2× |
| Easyjson | 66067 | 299.53 MB/s | 0 | 0 | 1.3× |
| Stdlib | 85552 | 231.31 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 709 | 774.08 MB/s | 480 | 1 | 5.2× |
| Easyjson | 1396 | 393.30 MB/s | 1616 | 5 | 2.7× |
| Sonic | 1601 | 342.84 MB/s | 2273 | 8 | 2.3× |
| Stdlib | 3716 | 147.73 MB/s | 1896 | 12 | 1.0× |
