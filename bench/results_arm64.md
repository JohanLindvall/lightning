# JSON Deserialization Benchmarks

- generated 2026-06-14T14:35:59Z
- go version go1.26.4 darwin/arm64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1071 | 1691.27 MB/s | 0 | 0 | 9.5× |
| Easyjson | 1840 | 984.65 MB/s | 24 | 1 | 5.5× |
| Stdlib | 10210 | 177.47 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1147 | 1579.34 MB/s | 144 | 10 | 9.0× |
| Easyjson | 1930 | 938.87 MB/s | 144 | 10 | 5.3× |
| Stdlib | 10309 | 175.77 MB/s | 920 | 17 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1633 | 302.46 MB/s | 160 | 1 | 2.4× |
| Easyjson | 1713 | 288.33 MB/s | 448 | 3 | 2.2× |
| Stdlib | 3843 | 128.54 MB/s | 760 | 12 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 373 | 616.81 MB/s | 160 | 1 | 7.1× |
| Easyjson | 1030 | 223.37 MB/s | 448 | 3 | 2.6× |
| Stdlib | 2648 | 86.87 MB/s | 760 | 12 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1268 | 1727.83 MB/s | 0 | 0 | 9.6× |
| Easyjson | 2306 | 950.04 MB/s | 24 | 1 | 5.3× |
| Stdlib | 12157 | 180.22 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1052 | 18817.17 MB/s | 0 | 0 | 79.7× |
| Easyjson | 70192 | 281.93 MB/s | 0 | 0 | 1.2× |
| Stdlib | 83815 | 236.10 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1035 | 530.57 MB/s | 480 | 1 | 3.9× |
| Easyjson | 1576 | 348.46 MB/s | 1616 | 5 | 2.5× |
| Stdlib | 3986 | 137.75 MB/s | 1896 | 12 | 1.0× |
