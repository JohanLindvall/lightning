# JSON Deserialization Benchmarks

- generated 2026-06-01T17:06:15Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1087 | 1666.40 MB/s | 120 | 9 | 9.5× |
| Easyjson | 1802 | 1005.56 MB/s | 120 | 9 | 5.8× |
| Stdlib | 10374 | 174.67 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 948 | 1911.83 MB/s | 0 | 0 | 9.9× |
| Easyjson | 1572 | 1152.48 MB/s | 24 | 1 | 6.0× |
| Stdlib | 9365 | 193.48 MB/s | 928 | 16 | 1.0× |
