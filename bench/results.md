# JSON Deserialization Benchmarks

- generated 2026-06-01T17:20:15Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1064 | 1703.26 MB/s | 120 | 9 | 8.6× |
| Easyjson | 1770 | 1023.61 MB/s | 120 | 9 | 5.2× |
| Stdlib | 9200 | 196.95 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 904 | 2004.20 MB/s | 0 | 0 | 10.3× |
| Easyjson | 1543 | 1174.57 MB/s | 24 | 1 | 6.0× |
| Stdlib | 9305 | 194.73 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 981 | 234.38 MB/s | 504 | 6 | 3.0× |
| Easyjson | 1076 | 213.75 MB/s | 448 | 3 | 2.8× |
| Stdlib | 2986 | 77.03 MB/s | 760 | 12 | 1.0× |
