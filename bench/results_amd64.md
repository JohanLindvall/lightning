# JSON Deserialization Benchmarks

- generated 2026-06-14T14:58:24Z
- go version go1.26.4 linux/amd64

Lower ns/op is better; throughput (MB/s) and allocations are reported by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.

## bench/cloudflare — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 886 | 2045.43 MB/s | 144 | 10 | 10.4× |
| Easyjson | 1827 | 991.74 MB/s | 144 | 10 | 5.1× |
| Stdlib | 9234 | 196.24 MB/s | 920 | 17 | 1.0× |

## bench/cloudflare-nocopy — 1812 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 723 | 2507.19 MB/s | 0 | 0 | 13.2× |
| Easyjson | 1677 | 1080.80 MB/s | 24 | 1 | 5.7× |
| Stdlib | 9511 | 190.51 MB/s | 928 | 16 | 1.0× |

## bench/float-array — 230 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 343 | 670.45 MB/s | 160 | 1 | 8.7× |
| Easyjson | 1097 | 209.66 MB/s | 448 | 3 | 2.7× |
| Stdlib | 2998 | 76.71 MB/s | 760 | 12 | 1.0× |

## bench/float-array-slow — 494 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 1329 | 371.84 MB/s | 160 | 1 | 3.0× |
| Easyjson | 1672 | 295.54 MB/s | 448 | 3 | 2.4× |
| Stdlib | 4034 | 122.46 MB/s | 760 | 12 | 1.0× |

## bench/pretty — 2191 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 910 | 2408.54 MB/s | 0 | 0 | 11.4× |
| Easyjson | 2070 | 1058.46 MB/s | 24 | 1 | 5.0× |
| Stdlib | 10345 | 211.79 MB/s | 928 | 16 | 1.0× |

## bench/skip-heavy — 19789 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 505 | 39175.39 MB/s | 0 | 0 | 150.2× |
| Easyjson | 47577 | 415.93 MB/s | 0 | 0 | 1.6× |
| Stdlib | 75851 | 260.89 MB/s | 240 | 6 | 1.0× |

## bench/time-array — 549 byte input

| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |
|---|--:|--:|--:|--:|--:|
| Lightning | 767 | 715.38 MB/s | 480 | 1 | 5.1× |
| Easyjson | 1563 | 351.33 MB/s | 1616 | 5 | 2.5× |
| Stdlib | 3944 | 139.21 MB/s | 1896 | 12 | 1.0× |
