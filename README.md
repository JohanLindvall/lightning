# Lightning ⚡

A small Go code generator that emits fast, allocation-light
`json.Unmarshaler` implementations from your struct definitions.

Instead of decoding JSON with reflection at run time (like `encoding/json`),
lightning reads a struct definition at build time and writes a hand-written
style `UnmarshalJSON` method plus the recursive decoders it needs. The decoders
share a single set of scanning primitives in [`pkg/support`](pkg/support), so the
generated files stay small.

## Installation

Run it straight from the module path (no clone needed):

```sh
go run github.com/JohanLindvall/lightning@latest path/to/data.go
```

`@latest` can be any version, branch, or commit (`@v1.0.0`, `@main`, `@<sha>`).
Or install the binary once:

```sh
go install github.com/JohanLindvall/lightning@latest
lightning path/to/data.go
```

The generated code imports `github.com/JohanLindvall/lightning/pkg/support`, so
the module you generate into must depend on lightning:

```sh
go get github.com/JohanLindvall/lightning
```

A `go:generate` directive in the file that holds your structs works well:

```go
//go:generate go run github.com/JohanLindvall/lightning@latest $GOFILE
```

## How it works

Point the generator at a Go file containing one or more struct types. From
inside this repo:

```sh
go run . path/to/data.go
```

(`go run .` only works in this repo; elsewhere use the module path shown above.)

For each input file `FOO.go` it writes `FOO_unmarshal.go` next to it, containing
an `UnmarshalJSON` method for every top-level struct type. The generated code
imports `github.com/JohanLindvall/lightning/pkg/support` for the shared scanner.

Given:

```go
package cloudflare

type Log struct {
    RayID              string `json:"RayID"`
    EdgeResponseStatus int64  `json:"EdgeResponseStatus"`
    // ...
}
```

you get a `func (v *Log) UnmarshalJSON(data []byte) error` that parses the JSON
with an index-based scanner, no reflection, and no allocation on the common
paths (unescaped strings, integers, object keys).

## Supported types

`string`, `bool`, every sized `int`/`uint` kind, `float32`/`float64`,
`json.RawMessage` (and `RawValue`), `time.Time` (RFC 3339, like `encoding/json`),
nested named and anonymous structs, slices, maps with string keys, pointers, and
`interface{}`/`any` (decoded into the usual Go representation of an arbitrary
JSON value). Unknown object keys are skipped.

## The `nocopy` tag option

By default, string and `json.RawMessage` fields copy their bytes out of the
input, matching `encoding/json` semantics. Add `nocopy` to the json tag to make
a field alias the input buffer instead — zero-copy, but the caller must keep the
input `[]byte` unchanged while the result is in use:

```go
type Log struct {
    RayID string          `json:"RayID,nocopy"` // aliases the input
    Body  json.RawMessage `json:"Body,nocopy"`  // aliases the input
}
```

`nocopy` propagates through slices, maps, and pointers, but stops at struct
boundaries (each struct's own field tags govern). Strings containing escape
sequences still allocate, since they can't be a slice of the raw input.

## SIMD string scanning

The scanner finds the next `"` or `\` inside a string in a single SIMD pass,
instead of two `bytes.IndexByte` scans:

- **amd64 (AVX2)** — 32 bytes/pass (`pkg/support/index_amd64.s`), selected via
  `cpu.X86.HasAVX2`.
- **arm64 (NEON/ASIMD)** — 16 bytes/pass (`pkg/support/index_arm64.s`), selected
  via `cpu.ARM64.HasASIMD`.

Detection is at run time (`golang.org/x/sys/cpu`); other platforms, CPUs without
the feature, and strings shorter than the vector width fall back to the
(already SIMD-optimized) `bytes.IndexByte` path. This speeds up string-heavy
payloads — roughly 10–20% on the Cloudflare log case on amd64 — with no
behavioral change. The arm64 path is correctness-tested under qemu.

## Benchmarks

The [`bench/`](bench) directory is a separate module (so its benchmark-only
dependency on [easyjson](https://github.com/mailru/easyjson) stays out of the
main module). It benchmarks the same payload decoded three ways — lightning,
`encoding/json`, and easyjson — across each `bench/<case>/` folder.

**See [`bench/results.md`](bench/results.md) for the full, current results.**

Run them yourself with:

```sh
./bench/run_bench.sh
```

which (re)generates each case's deserializers and writes `bench/results.txt`
and `bench/results.md`.

Representative numbers for a 1.8 KB Cloudflare log (Go 1.26, amd64):

| Decoder | ns/op | B/op | allocs/op | vs stdlib |
|---|--:|--:|--:|--:|
| lightning (`nocopy`) | ~950 | 0 | 0 | ~9.9× |
| lightning (default)  | ~1090 | 120 | 9 | ~9.5× |
| easyjson             | ~1570–1800 | 24–120 | 1–9 | ~6× |
| `encoding/json`      | ~9400 | 928 | 16 | 1.0× |

## Layout

| Path | Description |
|---|---|
| [`main.go`](main.go) | the generator (`package main`) |
| [`pkg/support`](pkg/support) | shared JSON scanning primitives used by generated code |
| [`bench/`](bench) | benchmark module: hand-written `data.go` + `input.json` per case, plus the generated decoders, harness, and results |

Generated files (`*_unmarshal.go`, `bench/*/bench_test.go`, `bench/*/ej/`, and
the `bench/results.*` outputs) are reproducible and excluded from version
control via [`.gitignore`](.gitignore).
