# JSON Deserialization Benchmarks

- generated 2026-06-01T17:55:26Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 987 | 1835.43 MB/s | 144 | 10 | 9.5× |
| Easyjson | 1797 | 1008.56 MB/s | 144 | 10 | 5.2× |
| Stdlib | 9361 | 193.58 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 771 | 2349.69 MB/s | 0 | 0 | 12.1× |
| Easyjson | 1627 | 1113.64 MB/s | 24 | 1 | 5.7× |
| Stdlib | 9300 | 194.83 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 606 | 379.46 MB/s | 504 | 6 | 5.0× |
| Easyjson | 1125 | 204.48 MB/s | 448 | 3 | 2.7× |
| Stdlib | 3045 | 75.52 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Easyjson | 1689 | 292.54 MB/s | 448 | 3 | 2.4× |
| Lightning | 1827 | 270.44 MB/s | 504 | 6 | 2.2× |
| Stdlib | 4079 | 121.11 MB/s | 760 | 12 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1529 | 359.06 MB/s | 1640 | 6 | 2.7× |
| Easyjson | 1621 | 338.76 MB/s | 1616 | 5 | 2.6× |
| Stdlib | 4157 | 132.06 MB/s | 1896 | 12 | 1.0× |
