# JSON Deserialization Benchmarks

- generated 2026-06-01T17:45:12Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1090 | 1662.17 MB/s | 144 | 10 | 8.3× |
| Easyjson | 1776 | 1020.39 MB/s | 144 | 10 | 5.1× |
| Stdlib | 9062 | 199.96 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 945 | 1916.59 MB/s | 0 | 0 | 9.9× |
| Easyjson | 1557 | 1164.10 MB/s | 24 | 1 | 6.0× |
| Stdlib | 9320 | 194.41 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 598 | 384.90 MB/s | 504 | 6 | 4.9× |
| Easyjson | 1103 | 208.59 MB/s | 448 | 3 | 2.7× |
| Stdlib | 2947 | 78.04 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Easyjson | 1584 | 311.96 MB/s | 448 | 3 | 2.6× |
| Lightning | 1718 | 287.54 MB/s | 504 | 6 | 2.4× |
| Stdlib | 4088 | 120.84 MB/s | 760 | 12 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1559 | 352.11 MB/s | 1640 | 6 | 2.6× |
| Easyjson | 1679 | 327.05 MB/s | 1616 | 5 | 2.4× |
| Stdlib | 4066 | 135.01 MB/s | 1896 | 12 | 1.0× |
