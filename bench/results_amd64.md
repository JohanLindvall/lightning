# JSON Deserialization Benchmarks

- generated 2026-06-14T16:03:10Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 894 | 2026.06 MB/s | 144 | 10 | 10.3× |
| Easyjson | 1732 | 1045.93 MB/s | 144 | 10 | 5.3× |
| Stdlib | 9204 | 196.87 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 613 | 2957.29 MB/s | 0 | 0 | 15.4× |
| Easyjson | 1608 | 1126.92 MB/s | 24 | 1 | 5.9× |
| Stdlib | 9440 | 191.95 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 664 | 2729.72 MB/s | 0 | 0 | 13.9× |
| Easyjson | 1605 | 1128.64 MB/s | 24 | 1 | 5.8× |
| Stdlib | 9241 | 196.08 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 340 | 675.68 MB/s | 160 | 1 | 8.4× |
| Easyjson | 1038 | 221.54 MB/s | 448 | 3 | 2.8× |
| Stdlib | 2860 | 80.41 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1312 | 376.45 MB/s | 160 | 1 | 3.0× |
| Easyjson | 1568 | 315.06 MB/s | 448 | 3 | 2.5× |
| Stdlib | 3982 | 124.07 MB/s | 760 | 12 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 840 | 2608.39 MB/s | 0 | 0 | 12.6× |
| Easyjson | 1955 | 1120.95 MB/s | 24 | 1 | 5.4× |
| Stdlib | 10561 | 207.45 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 512 | 38647.41 MB/s | 0 | 0 | 147.0× |
| Easyjson | 47174 | 419.49 MB/s | 0 | 0 | 1.6× |
| Stdlib | 75286 | 262.85 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 732 | 749.76 MB/s | 480 | 1 | 5.3× |
| Easyjson | 1608 | 341.42 MB/s | 1616 | 5 | 2.4× |
| Stdlib | 3906 | 140.56 MB/s | 1896 | 12 | 1.0× |
