# JSON Deserialization Benchmarks

- generated 2026-06-01T18:32:05Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1032 | 1756.60 MB/s | 144 | 10 | 9.3× |
| Easyjson | 1836 | 987.18 MB/s | 144 | 10 | 5.2× |
| Stdlib | 9571 | 189.33 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 791 | 2290.47 MB/s | 0 | 0 | 11.4× |
| Easyjson | 1633 | 1109.95 MB/s | 24 | 1 | 5.5× |
| Stdlib | 9034 | 200.57 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 601 | 382.51 MB/s | 504 | 6 | 5.0× |
| Easyjson | 1099 | 209.21 MB/s | 448 | 3 | 2.7× |
| Stdlib | 2992 | 76.87 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Easyjson | 1688 | 292.59 MB/s | 448 | 3 | 2.4× |
| Lightning | 1787 | 276.50 MB/s | 504 | 6 | 2.3× |
| Stdlib | 4028 | 122.64 MB/s | 760 | 12 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1299 | 1686.12 MB/s | 0 | 0 | 8.3× |
| Easyjson | 2020 | 1084.81 MB/s | 24 | 1 | 5.3× |
| Stdlib | 10807 | 202.73 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 823 | 24057.49 MB/s | 0 | 0 | 91.2× |
| Easyjson | 47671 | 415.11 MB/s | 0 | 0 | 1.6× |
| Stdlib | 75041 | 263.71 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1444 | 380.17 MB/s | 1640 | 6 | 2.8× |
| Easyjson | 1588 | 345.80 MB/s | 1616 | 5 | 2.6× |
| Stdlib | 4103 | 133.81 MB/s | 1896 | 12 | 1.0× |
