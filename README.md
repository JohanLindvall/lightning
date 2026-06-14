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
`json.RawMessage` (and `RawValue`), `time.Time` (RFC 3339, like `encoding/json`;
the [`lax`](#the-lax-tag-option) option also accepts a space separator and Unix
timestamps), nested named and anonymous structs, slices, maps with string keys,
pointers, and `interface{}`/`any` (decoded into the usual Go representation of an
arbitrary JSON value). Unknown object keys are skipped.

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

## The `lax` tag option

By default a value of the wrong type fails the whole decode: a string where a
number is expected returns an error. Add `lax` to the json tag to make such a
mismatch a no-op instead — the offending value is skipped and the field left at
its zero value, while the rest of the object decodes normally:

```go
type Log struct {
    Status int64 `json:"Status,lax"` // a non-number Status is ignored, leaving 0
}
```

Only type mismatches are tolerated; genuinely malformed JSON (a syntax error in
the value) still fails, since a well-formed value of the wrong type can be
skipped but a broken one cannot. `lax` works for every field type, including
nested structs, slices, and maps, where a decode error anywhere in the value
leaves the whole field unset. It combines with the other options and with
alternate names — `json:"Name|Title,nocopy,lax"`.

On a `time.Time` field, `lax` additionally widens what counts as a valid
timestamp. Besides strict RFC 3339, it accepts a space in place of the `T`
date/time separator and a Unix timestamp given as a JSON number or numeric
string, inferring seconds, milliseconds, or microseconds from the magnitude; the
result is normalized to UTC. An unrecognized timestamp is skipped and the field
left unset, like any other lax mismatch. As with `nocopy`, the lenient parser
propagates through slices, maps, and pointers (e.g. `[]time.Time`) but stops at
struct boundaries.

## Comment directives

Some behavior is selected with a `//lightning:<name>` comment on the struct type
(or its declaration), separate from the per-field json tags above.

### `//lightning:compact`

By default a decoder calls `SkipWS` around every token so it accepts JSON with
any whitespace. Mark a type `//lightning:compact` to assert the input has no
whitespace *between* tokens — the form `encoding/json`'s `Marshal` and most wire
protocols emit — and the generator drops those inter-token `SkipWS` calls,
decoding tokens back-to-back:

```go
//lightning:compact
type Log struct {
    RayID  string `json:"RayID"`
    Status int64  `json:"Status"`
}
```

This runs about 10% faster on object-heavy payloads (see the `cloudflare-compact`
benchmark). Whitespace surrounding the whole document is still tolerated — a
trailing newline is fine — so only *inter-token* whitespace is assumed absent.

The directive is an assertion you make about the input: a compact decoder fed
input that does contain inter-token whitespace (for example the same document
pretty-printed) returns an error instead of parsing it. Use it only for sources
that are guaranteed compact. The directive applies to the whole type graph it
roots, including nested structs, slices, and maps.

## Generated function names

The `UnmarshalJSON` methods keep their exact name (the `json.Unmarshaler`
interface requires it). The unexported decoder helpers they call are named
`lightning<ImportPath><Type>decode…` — a prefix derived from the package's import
path and the top-level type — so generating decoders for several types into one
package never produces colliding helper names. No annotation is needed; the
prefix is automatic.

## Key lookups

When you only need a few values out of a document and don't want to generate (or
decode into) a struct, the [`pkg/json`](pkg/json) package exposes the scanner's
key-lookup primitives. They walk the input with the same allocation-free
`Skip`/`ReadKey` machinery the generated decoders use, and every value they
return aliases the input `[]byte` (keep it unchanged while the result is in
use). A returned value follows the same conventions throughout: a string keeps
its surrounding quotes with escapes intact, an object or array spans the whole
`{`…`}` or `[`…`]`, and a scalar is the literal token.

- `Get(data []byte, keys ...string) ([]byte, int, error)` — walks the object-key
  path `keys` one level per key and returns the value's raw bytes (and the offset
  it starts at), modeled on [buger/jsonparser](https://github.com/buger/jsonparser)'s
  `Get`. With no keys it returns the whole root value; a missing key returns
  `ErrKeyNotFound`.
- `GetMany(data []byte, keys []string, out [][]byte) ([][]byte, error)` — looks up
  several *top-level* keys in a **single pass** over the object, where N separate
  `Get` calls would rescan it N times. Results are written into `out[:0]` (pass a
  slice to reuse across calls, allocation-free; a `nil` reuses nothing) and
  returned with `len == len(keys)`: `out[n]` is the value for `keys[n]`, or `nil`
  if that key is absent. A missing key is reported by the `nil` slot, not an
  error (a present key whose value is JSON `null` yields the bytes `"null"`,
  distinct from absent); a non-object root or malformed JSON returns an error.
- `ObjectEach(data []byte, fn func(key string, value []byte) error, keys ...string) error`
  — calls `fn` for every member of the object reached by the path `keys` (the
  root object with no keys). If `fn` returns an error, iteration stops and
  returns it.

```go
// Pull a few fields out of a log record in one pass, reusing a scratch slice.
keys := []string{"ClientIP", "EdgeResponseStatus", "RayID"}
vals, err := json.GetMany(data, keys, scratch[:0])
// vals[0] == []byte(`"203.0.113.23"`), vals[1] == []byte("599"), …
```

Each function has a **compact counterpart** — `GetCompact`, `GetManyCompact`,
`ObjectEachCompact` — with the identical signature and result. Like the
[`//lightning:compact`](#lightningcompact) directive, they assume the input has
no whitespace *between* tokens (the form `encoding/json`'s `Marshal` and most
wire protocols emit) and skip the inter-token `SkipWS` scans, running about 10%
faster; whitespace surrounding the whole document is still tolerated. Feed one
input that does contain inter-token whitespace and it may return an error, so use
them only for sources guaranteed compact.

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

## Number parsing

The [`pkg/json`](pkg/json) package also exposes the scanner's float parser:

- `ParseFloat(b []byte) (float64, error)` — parses the JSON number in `b` as a
  `float64`. It takes the scanner's Clinger fast path — when the mantissa is
  exact and the decimal exponent is small, the result is a single multiply or
  divide by a power of ten — and falls back to `strconv.ParseFloat` for the rest.
  `b` must be exactly one number with no surrounding whitespace; trailing bytes
  or an empty input return an error. Nothing is retained or copied.

## SIMD scanning

Two hot scan loops use a single vectorized pass instead of byte-at-a-time work,
with kernels in `pkg/support/index_amd64.s` and `pkg/support/index_arm64.s`
(arm64 uses NEON/ASIMD, 16 bytes/pass, for both):

- **next `"` or `\` in a string** — replaces two `bytes.IndexByte` scans; speeds
  up string-heavy payloads. On amd64 it uses SSE2 (16-byte vectors, two compares
  per 32-byte step), which avoids the `VZEROUPPER` an AVX2 routine must run on
  every call — pure overhead for the short keys and values typical of JSON.
- **next structural byte (`{`, `}`, `[`, `]`, `"`)** — AVX2 on amd64, 32
  bytes/pass, lets `skipObject` / `skipArray` jump over inert content (numbers,
  keys, whitespace) when skipping unknown values. Skipping a large ignored
  array/object is dramatically faster (the `skip-heavy` benchmark decodes at
  >20 GB/s, ~90× `encoding/json`).

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

**See the per-architecture results for the full, current numbers:
[`bench/results_amd64.md`](bench/results_amd64.md) and
[`bench/results_arm64.md`](bench/results_arm64.md).**

Run them yourself with:

```sh
./bench/run_bench.sh
```

which (re)generates each case's deserializers and writes `bench/results.txt`
and an architecture-specific `bench/results_<goarch>.md` (so runs on different
CPUs do not overwrite each other's committed results).

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
| [`pkg/json`](pkg/json) | small public API over the scanner (`Get`/`GetMany`/`ObjectEach`, `UnescapeString`, `ParseFloat`) |
| [`bench/`](bench) | benchmark module: hand-written `data.go` + `input.json` per case, plus the generated decoders, harness, and results |

Generated files (`*_unmarshal.go`, `bench/*/bench_test.go`, `bench/*/ej/`, and
the `bench/results.*` outputs) are reproducible and excluded from version
control via [`.gitignore`](.gitignore).
