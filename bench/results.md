# JSON Deserialization Benchmarks

- generated 2026-06-01T18:06:10Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 964 | 1879.92 MB/s | 144 | 10 | 10.0× |
| Easyjson | 1857 | 975.79 MB/s | 144 | 10 | 5.2× |
| Stdlib | 9599 | 188.78 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 778 | 2328.35 MB/s | 0 | 0 | 12.2× |
| Easyjson | 1665 | 1088.28 MB/s | 24 | 1 | 5.7× |
| Stdlib | 9495 | 190.84 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 605 | 380.28 MB/s | 504 | 6 | 4.9× |
| Easyjson | 1114 | 206.42 MB/s | 448 | 3 | 2.6× |
| Stdlib | 2950 | 77.97 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Easyjson | 1708 | 289.31 MB/s | 448 | 3 | 2.4× |
| Lightning | 1760 | 280.64 MB/s | 504 | 6 | 2.3× |
| Stdlib | 4022 | 122.83 MB/s | 760 | 12 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1462 | 375.45 MB/s | 1640 | 6 | 2.9× |
| Easyjson | 1569 | 350.01 MB/s | 1616 | 5 | 2.7× |
| Stdlib | 4300 | 127.68 MB/s | 1896 | 12 | 1.0× |
