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

## Alternate field names

A json tag may list several pipe-separated names. Any of them appearing in the
input fills the field, which is handy when an upstream source renamed a key and
you want to accept both spellings:

```go
type Log struct {
    EdgeResponseStatus int64 `json:"EdgeResponseStatus|AnotherField"`
}
```

Comma-separated options still follow the name as usual, so names and `nocopy`
combine freely — `json:"Name|Title,nocopy"` accepts both `Name` and `Title`,
zero-copy.

## String escaping and unescaping

The [`pkg/json`](pkg/json) package exposes the scanner's string codec on its
own, for when you have a JSON string body (the bytes between the quotes) and
just want to decode or encode it.

**Unescaping** (escaped body → decoded value):

- `UnescapeString(in []byte) (string, error)` — decodes the escapes in `in`. If
  `in` contains no escapes the result aliases `in` with no copy; otherwise a new
  string is allocated. `in` is left unchanged.
- `UnescapeStringInto(in, out []byte) (string, error)` — same, but writes the
  decoded bytes into `out` instead of allocating (mirroring the `Unescape(in,
  out)` convention of [buger/jsonparser](https://github.com/buger/jsonparser)).
  With no escapes the result aliases `in`; otherwise it aliases `out` and
  allocates nothing when `cap(out) >= len(in)`, since unescaping never lengthens
  a string. Pass `out == in` (e.g. `in[:0]`) to decode truly in place,
  overwriting `in`.

Both return a string that aliases a buffer, so keep that buffer unchanged while
the result is in use.

**Escaping** (raw value → escaped body, escaping `"`, `\`, and control bytes;
`/` is left as-is and `\b`/`\f` are written in `\u00XX` form):

- `EscapeString(s []byte, out *strings.Builder)` — writes the escaped form of
  `s` to `out`. The common no-escape case is detected with a vectorized scan and
  written straight to the builder, with no scratch buffer or copy.
- `EscapeStringInto(s, out []byte) []byte` — appends the escaped form of `s` to
  `out` and returns the extended slice, allocating nothing when `out` has room
  (escaping can grow a string up to 6× for control bytes). Pass `out[:0]` to
  reuse a buffer across calls.

Clean runs are skipped eight bytes at a time (SWAR), so strings that need little
or no escaping cost roughly one `memcpy`.

## SIMD scanning

Two hot scan loops use a single vectorized pass instead of byte-at-a-time work
— **amd64 (AVX2)**, 32 bytes/pass (`pkg/support/index_amd64.s`), and
**arm64 (NEON/ASIMD)**, 16 bytes/pass (`pkg/support/index_arm64.s`):

- **next `"` or `\` in a string** — replaces two `bytes.IndexByte` scans; speeds
  up string-heavy payloads (~10–20% on the Cloudflare case).
- **next structural byte (`{`, `}`, `[`, `]`, `"`)** — lets `skipObject` /
  `skipArray` jump over inert content (numbers, keys, whitespace) when skipping
  unknown values. Skipping a large ignored array/object is dramatically faster
  (the `skip-heavy` benchmark decodes at >20 GB/s, ~90× `encoding/json`).

Feature detection is at run time (`golang.org/x/sys/cpu`); other platforms, CPUs
without the feature, and inputs shorter than the vector width fall back to scalar
(`bytes.IndexByte`-based) code. Behavior is identical across paths — verified by
fuzzing each primitive against a reference and by `SkipValue`/decode round-trips
vs `encoding/json`, on amd64 and on arm64 under qemu.

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
| [`pkg/json`](pkg/json) | small public API over the scanner (e.g. `UnescapeString`) |
| [`bench/`](bench) | benchmark module: hand-written `data.go` + `input.json` per case, plus the generated decoders, harness, and results |

Generated files (`*_unmarshal.go`, `bench/*/bench_test.go`, `bench/*/ej/`, and
the `bench/results.*` outputs) are reproducible and excluded from version
control via [`.gitignore`](.gitignore).
