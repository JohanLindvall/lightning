# JSON Deserialization Benchmarks

- generated 2026-06-14T14:42:48Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 997 | 1818.08 MB/s | 144 | 10 | 9.2× |
| Easyjson | 1759 | 1030.19 MB/s | 144 | 10 | 5.2× |
| Stdlib | 9125 | 198.57 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 806 | 2247.95 MB/s | 0 | 0 | 11.9× |
| Easyjson | 1663 | 1089.89 MB/s | 24 | 1 | 5.8× |
| Stdlib | 9620 | 188.36 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 364 | 631.14 MB/s | 160 | 1 | 8.1× |
| Easyjson | 1059 | 217.25 MB/s | 448 | 3 | 2.8× |
| Stdlib | 2953 | 77.89 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1427 | 346.29 MB/s | 160 | 1 | 2.9× |
| Easyjson | 1658 | 297.92 MB/s | 448 | 3 | 2.5× |
| Stdlib | 4139 | 119.37 MB/s | 760 | 12 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1011 | 2168.07 MB/s | 0 | 0 | 10.5× |
| Easyjson | 2008 | 1091.29 MB/s | 24 | 1 | 5.3× |
| Stdlib | 10652 | 205.68 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 532 | 37183.36 MB/s | 0 | 0 | 141.9× |
| Easyjson | 47909 | 413.06 MB/s | 0 | 0 | 1.6× |
| Stdlib | 75496 | 262.12 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 974 | 563.57 MB/s | 480 | 1 | 4.4× |
| Easyjson | 1571 | 349.42 MB/s | 1616 | 5 | 2.7× |
| Stdlib | 4267 | 128.67 MB/s | 1896 | 12 | 1.0× |
