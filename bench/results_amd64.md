# JSON Deserialization Benchmarks

- generated 2026-06-14T15:19:31Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 914 | 1981.79 MB/s | 144 | 10 | 10.0× |
| Easyjson | 1739 | 1042.25 MB/s | 144 | 10 | 5.3× |
| Stdlib | 9176 | 197.47 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 665 | 2723.67 MB/s | 0 | 0 | 13.9× |
| Easyjson | 1596 | 1135.06 MB/s | 24 | 1 | 5.8× |
| Stdlib | 9240 | 196.10 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 314 | 731.26 MB/s | 160 | 1 | 8.9× |
| Easyjson | 1009 | 227.87 MB/s | 448 | 3 | 2.8× |
| Stdlib | 2813 | 81.77 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1291 | 382.75 MB/s | 160 | 1 | 3.1× |
| Easyjson | 1547 | 319.28 MB/s | 448 | 3 | 2.6× |
| Stdlib | 3977 | 124.20 MB/s | 760 | 12 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 832 | 2632.46 MB/s | 0 | 0 | 12.6× |
| Easyjson | 1948 | 1125.00 MB/s | 24 | 1 | 5.4× |
| Stdlib | 10457 | 209.52 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 503 | 39321.74 MB/s | 0 | 0 | 146.7× |
| Easyjson | 47231 | 418.98 MB/s | 0 | 0 | 1.6× |
| Stdlib | 73819 | 268.07 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 729 | 752.84 MB/s | 480 | 1 | 5.7× |
| Easyjson | 1641 | 334.54 MB/s | 1616 | 5 | 2.5× |
| Stdlib | 4123 | 133.14 MB/s | 1896 | 12 | 1.0× |
