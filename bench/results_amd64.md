# JSON Deserialization Benchmarks

- generated 2026-06-14T17:53:38Z
- go version go1.26.3 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1086 | 1669.11 MB/s | 144 | 10 | 10.5× |
| Easyjson | 2408 | 752.55 MB/s | 144 | 10 | 4.7× |
| Stdlib | 11430 | 158.54 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-compact — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 849 | 2133.85 MB/s | 0 | 0 | 13.6× |
| Easyjson | 2184 | 829.65 MB/s | 24 | 1 | 5.3× |
| Stdlib | 11579 | 156.49 MB/s | 928 | 16 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 900 | 2013.55 MB/s | 0 | 0 | 13.2× |
| Easyjson | 2208 | 820.77 MB/s | 24 | 1 | 5.4× |
| Stdlib | 11849 | 152.93 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 374 | 614.14 MB/s | 160 | 1 | 9.0× |
| Easyjson | 1340 | 171.61 MB/s | 448 | 3 | 2.5× |
| Stdlib | 3364 | 68.37 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1075 | 459.60 MB/s | 160 | 1 | 4.4× |
| Easyjson | 1958 | 252.35 MB/s | 448 | 3 | 2.4× |
| Stdlib | 4711 | 104.85 MB/s | 760 | 12 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1125 | 1947.00 MB/s | 0 | 0 | 11.9× |
| Easyjson | 2657 | 824.57 MB/s | 24 | 1 | 5.0× |
| Stdlib | 13393 | 163.60 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 514 | 38496.65 MB/s | 0 | 0 | 221.1× |
| Easyjson | 79834 | 247.88 MB/s | 0 | 0 | 1.4× |
| Stdlib | 113658 | 174.11 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 906 | 606.22 MB/s | 480 | 1 | 5.2× |
| Easyjson | 1725 | 318.17 MB/s | 1616 | 5 | 2.7× |
| Stdlib | 4712 | 116.51 MB/s | 1896 | 12 | 1.0× |
