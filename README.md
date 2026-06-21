# Lightning ⚡

A small Go code generator that emits fast, allocation-light
`json.Unmarshaler` implementations from your struct definitions.

Instead of decoding JSON with reflection at run time (like `encoding/json`),
lightning reads a struct definition at build time and writes a hand-written
style `UnmarshalJSON` method plus the recursive decoders it needs. The decoders
share a single set of scanning primitives in [`pkg/unstable`](pkg/unstable), so the
generated files stay small.

Those same primitives are also exposed directly as a small toolkit in
[`pkg/json`](pkg/json), for working with JSON without generating or decoding into
a struct at all: pull a few fields out of a document
([`Get`/`GetMany`](#key-lookups)), edit values in place
([`Set`/`SetMany`](#setting-a-value)), prune default members
([`StripDefaults`](#stripping-default-fields)), decode into a generic value
([`DecodeAny`](#decoding-into-any)), and escape, unescape, or parse a JSON number
on its own — each an allocation-light, single-pass operation over the raw bytes.
See [Layout](#layout) for the full list.

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

The generated code imports `github.com/JohanLindvall/lightning/pkg/unstable`, so
the module you generate into must depend on lightning:

```sh
go get github.com/JohanLindvall/lightning
```

`pkg/unstable` is the generator's runtime: it is exported only because the
generated `*_unmarshal.go` files (which live in your module) have to call into it.
As its name says, it is **not a stable API** — don't import it directly; use
[`pkg/json`](pkg/json) for the public toolkit.

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
an `UnmarshalJSON` method for every top-level type (a struct, or a named slice or
map root — see [Root types](#root-types)). The generated code imports
`github.com/JohanLindvall/lightning/pkg/unstable` for the shared scanner.

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
timestamps), nested named and anonymous structs, slices, fixed-size arrays
(`[N]T`), maps with string keys, pointers, and `interface{}`/`any` (decoded into
the usual Go representation of an arbitrary JSON value). Unknown object keys are
skipped.

A fixed-size array follows `encoding/json`: the leading elements are filled, a
shorter JSON array leaves the remaining elements zero, and a longer one's extras
are discarded.

Embedded struct fields are promoted like `encoding/json`: an embedded struct's
exported fields decode as if they were the outer struct's own (an embedded
pointer is allocated on demand), a name present on both the outer struct and an
embed is resolved by Go's shallower-wins rule, an equal-depth clash is dropped
unless a single field is tagged, and an embedded field with its own JSON tag name
is a plain named field rather than promoted. Embedding a type from another
package, whose fields aren't visible to the generator, is the one gap — it is
decoded as a single named field instead of being flattened.

## Root types

A top-level type need not be a struct: a named slice or a string-keyed map at the
root gets its own `UnmarshalJSON` too, so a document that is a bare array or a
data map needs no wrapper struct.

```go
type Records []Record          // a JSON array at the root
type ByID    map[string]Record // a JSON object used as a data map
```

`type Records []Record` decodes a top-level `[…]` with the slice element rules;
`type ByID map[string]Record` decodes a top-level `{…}` as a map, its keys the
object's member names. Either element/value type, and any nested types and field
options, behave exactly as the same type used for a struct field would. Several
root types (struct, slice, map, in any mix) can live in one input file. For a
root that is a *bare* `any`/`interface{}` — whose schema you don't know at all —
there is no method to generate (Go forbids methods on interface types); decode it
dynamically with [`json.DecodeAny`](#decoding-into-any) instead.

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

## The `//lightning:destructive` directive

The one allocation `nocopy` can't avoid is unescaping: a string like `"a\/b"` has
to be decoded somewhere, and it can't alias the raw input. The
`//lightning:destructive` directive removes even that allocation by unescaping
**into the input buffer** — overwriting the escaped bytes with the decoded ones (the
decoded form is always shorter) and aliasing the result. Put it on the type, above
its declaration:

```go
//lightning:destructive
type Log struct {
    RayID   string `json:"RayID,nocopy"`
    Message string `json:"Message,nocopy"`
}
```

The name says it plainly: this **mutates — and effectively destroys — the input
document**. The bytes of every escaped string are overwritten in place, so the
`[]byte` you pass in is no longer valid JSON afterward, and you must not read it
again or hold any other alias into it. In return, escaped `nocopy` strings cost zero
allocations. Use it when you own the input buffer and discard it after decoding (a
request body, a buffer you'll reuse). It upgrades the type's `nocopy` string fields;
escape-free input decodes byte-identically to plain `nocopy`. On escape-heavy
documents the savings are large — e.g. **−41% time and −86% bytes** on a
string-heavy corpus.

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

## The `unwrap` tag option

Some payloads carry a nested document as a *string* — JSON embedded in JSON,
sometimes base64-encoded. Add `unwrap` to a field's json tag to decode through
that wrapper: the field's value is read as a JSON string, its body unescaped,
and the result decoded as JSON into the field.

```go
type Envelope struct {
    Name    string  `json:"name"`
    Payload Message `json:"payload,unwrap"` // value is a string holding JSON
}
```

Both forms are accepted automatically. If the unescaped string is itself JSON
(its first non-whitespace byte starts a JSON value) it is decoded directly;
otherwise it is base64-decoded first (standard alphabet, with or without
padding) and the decoded bytes are the JSON. So a `"payload"` of
`"{\"id\":7}"` and of `"eyJpZCI6N30="` both fill `Payload`. A `null` or empty
string leaves the field at its zero value.

The field decodes with its normal rules, so `unwrap` composes with the field's
type (struct, slice, map, scalar…) and with `nocopy` — a `nocopy` string inside
the embedded document aliases the decoded buffer, which is retained for as long
as the result is in use. The embedded document is parsed as a fresh input, so
its own whitespace, escaping, and structure are independent of the outer JSON.

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

This runs a few percent faster on object-heavy payloads (the `cloudflare-compact`
benchmark beats `cloudflare-nocopy`, its non-compact equivalent, by ~4%).
Whitespace surrounding the whole document is still tolerated — a trailing newline
is fine — so only *inter-token* whitespace is assumed absent.

The directive is an assertion you make about the input: a compact decoder fed
input that does contain inter-token whitespace (for example the same document
pretty-printed) returns an error instead of parsing it. Use it only for sources
that are guaranteed compact. The directive applies to the whole type graph it
roots, including nested structs, slices, and maps.

### `//lightning:nocopy`

The [`nocopy`](#the-nocopy-tag-option) tag option lives on a *field*, but a named
slice or map **root** type has no field to tag — its element strings and map keys
would always be copied. Mark such a root `//lightning:nocopy` to alias them out of
the input instead, the same zero-copy trade as the tag (the caller must keep the
input `[]byte` unchanged while the result is in use):

```go
//lightning:nocopy
type Index map[string]Record // the map keys alias the input, not copied
```

It applies to what the root itself owns — a map's keys, a slice's string elements
— while each value/element type keeps its own field tags. On `map[string]struct`
documents with many keys this is a real saving: tagging the GeoSciences `gsoc_2018`
corpus's root map cut its allocations ~21%. (Only slice and map roots take the
directive; a struct root uses per-field `nocopy` tags.)

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
  it starts at), without reporting a value type. With no keys it returns the whole
  root value; a missing key returns `ErrKeyNotFound`.
- `GetMany(data []byte, keys []string, out [][]byte) ([][]byte, error)` — looks up
  several *top-level* keys in a **single pass** over the object, where N separate
  `Get` calls would rescan it N times. Results are written into `out[:0]` (pass a
  slice to reuse across calls, allocation-free; a `nil` reuses nothing) and
  returned with `len == len(keys)`: `out[n]` is the value for `keys[n]`, or `nil`
  if that key is absent. A missing key is reported by the `nil` slot, not an
  error (a present key whose value is JSON `null` yields the bytes `"null"`,
  distinct from absent); a non-object root or malformed JSON returns an error.
- `GetPaths(data []byte, paths [][]string, out [][]byte) ([][]byte, error)` — the
  multi-path form of `Get` (as `GetMany` is its multi-key form): each `paths[n]` is
  a nested key path and `out[n]` receives the value there. The document is walked
  **once** and paths that share a prefix share that descent, so pulling several
  nested fields — especially under a common parent — costs one traversal rather than
  one `Get` per field (≈40–50% faster on a record with a handful of nested paths).
  Same result conventions as `GetMany`; a `nil`/empty path selects the root.
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

## Decoding into `any`

When a document's shape isn't known ahead of time — too variable to model, or
genuinely arbitrary — `DecodeAny` reads the whole thing into the generic
representation `encoding/json` produces for `interface{}`: `nil`, `bool`,
`float64`, `string`, `[]any`, and `map[string]any`. It is the dynamic
counterpart to a generated unmarshaler, using the same scanner but with no target
type.

- `DecodeAny(data []byte, compact bool) (any, error)` — decodes the single JSON
  value in `data`. `compact` assumes no inter-token whitespace and skips the
  `SkipWS` scans (as [`GetCompact`](#key-lookups) does), faster on minified
  input. Whitespace around the whole document is tolerated; trailing
  non-whitespace content after the value is an error. Unlike the key-lookup
  helpers it builds Go values (so strings are unescaped and copied, numbers
  parsed), allocating the maps and slices the result needs.

## String escaping and unescaping

The [`pkg/json`](pkg/json) package exposes the scanner's string codec on its
own, for when you have a JSON string body (the bytes between the quotes) and
just want to decode or encode it.

**Unescaping** (escaped body → decoded value):

- `UnescapeString(in []byte) (string, error)` — decodes the escapes in `in`. If
  `in` contains no escapes the result aliases `in` with no copy; otherwise a new
  string is allocated. `in` is left unchanged.
- `UnescapeStringInto(in, out []byte) (string, error)` — same, but writes the
  decoded bytes into `out` instead of allocating.
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
  `float64`. It takes the scanner's Clinger fast path first — when the mantissa is
  exact and the decimal exponent is small, the result is a single multiply or
  divide by a power of ten. Numbers that miss it (a mantissa ≥ 2^53 or a larger
  exponent, e.g. high-precision coordinates) take an Eisel-Lemire pass that
  converts the extracted mantissa and exponent with a 128-bit multiply, the same
  fast path `strconv.ParseFloat` uses internally but without re-scanning the
  digits; only the rare ambiguous or >19-digit cases fall back to
  `strconv.ParseFloat`. The Eisel-Lemire result is bit-for-bit identical to
  `strconv` (verified by a differential fuzz test). `b` must be exactly one number
  with no surrounding whitespace; trailing bytes or an empty input return an
  error. Nothing is retained or copied.

## Stripping default fields

The [`pkg/json`](pkg/json) package can also prune a JSON document in a single
pass, dropping object members whose value is a "default" — useful for shrinking
verbose, mostly-default records (Cloudflare HTTP logs, say) before storing or
forwarding them:

- `StripDefaults(input, output []byte, defaults, keep [][]byte, ws WhitespaceMode) []byte`
  — copies `input` to `output`, dropping every object member whose value is a
  default and then dropping any object or array that this leaves empty. A value
  is a default when it is byte-equal to one of `defaults`, compared against the
  bare token — the unquoted contents for a string, the literal token for a number
  or keyword (e.g. `[]byte("none")`, `[]byte("false")`). Empty values are *not*
  special-cased: to drop empty strings (and other empty tokens) include an empty
  entry `[]byte("")` in `defaults`. A member is kept despite a default value when
  its unquoted key is byte-equal to one of `keep` (e.g. `[]byte("WallTimeMs")`).
  String values keep their surrounding quotes and escapes; scalars keep their
  literal token. (Empty `{}`/`[]` are always dropped, independent of `defaults`.)

`output` is filled from the front and the populated prefix is returned; `input`
is never modified. StripDefaults never lengthens the document, so `output` is grown
(allocated) only when `cap(output) < len(input)` — pass `input[:0]` to strip in
place, or a reused buffer to run allocation-free. It is best effort and copies
malformed input through unchanged.

`ws` chooses how inter-token whitespace is handled:
- `RemoveWhitespace` (the zero value) — tolerate any whitespace; output is compact.
- `AssumeCompact` — assert the input has no inter-token whitespace and skip the
  `SkipWS` scans (faster, as [`GetCompact`](#key-lookups) does); misreads spaced input.
- `PreserveWhitespace` — keep the input's whitespace around surviving content, so a
  pretty-printed document stays pretty-printed; only dropped members are removed.

```go
// "" opts empty strings in; "0"/"none"/"false"/"unknown" are the non-empty defaults.
defaults := [][]byte{[]byte(""), []byte("0"), []byte("none"), []byte("false"), []byte("unknown")}
keep := [][]byte{[]byte("WallTimeMs")} // retained even when their value is a default
var scratch []byte
scratch = json.StripDefaults(record, scratch[:0], defaults, keep, json.AssumeCompact)
// {"a":0,"b":"x","e":"","WallTimeMs":0}  ->  {"b":"x","WallTimeMs":0}
```

Matching is length-pre-filtered so a value or key longer than any candidate
skips the scan, and a kept member is moved with a single `copy` when its
`"key":value` span is contiguous in the input.

## Setting a value

The [`pkg/json`](pkg/json) package can also splice values into a document in a
single pass, without a full unmarshal/edit/marshal round-trip:

- `Set(in, out, rawVal []byte, keys []string) []byte` — returns `in` with the
  value at the object-key path `keys` replaced by the raw JSON `rawVal`, written
  into `out`. A path that doesn't exist is created: a missing member is inserted
  into its parent, and a missing intermediate object (or a non-object found where
  the path must still descend) is built up as nested objects. With no keys the
  whole document becomes `rawVal`.
- `SetMany(in, out []byte, rawVal [][]byte, keys []string) []byte` — sets several
  of the root object's own keys at once: `keys[i]`'s value becomes `rawVal[i]`,
  replacing it where the key exists and appending a member where it doesn't.
  `SetMany` is to `Set` what [`GetMany`](#key-lookups) is to `Get` — one walk over
  the object instead of rescanning and rewriting it once per key. A non-object
  root is replaced by a fresh object of all the members; a `rawVal` shorter than
  `keys` ignores the surplus keys.
- `SetPaths(in, out []byte, rawVal [][]byte, paths [][]string) []byte` — the
  multi-*path* form of `Set` (the write counterpart of `GetPaths`): each `paths[i]`
  is a nested key path set to `rawVal[i]`, replaced if present or created (with any
  missing intermediate objects) if absent. Paths that share a prefix are edited and
  created together, so the document is rewritten **once** rather than once per path
  — a large win over sequential `Set`, which re-reads and re-writes the whole
  document each call (≈2–3× faster, far fewer allocations). A `nil`/empty path
  replaces the whole document; when one path is a prefix of another the shorter
  wins.

```go
// {"a":{"b":1}}  ->  {"a":{"b":1},"c":[true]}
out = json.Set(doc, out[:0], []byte("[true]"), []string{"c"})

// {"a":1,"b":2}  ->  {"a":1,"b":9,"c":3}   (replace b, add c, one pass)
out = json.SetMany(doc, out[:0], [][]byte{[]byte("9"), []byte("3")}, []string{"b", "c"})

// {"a":{"b":1}}  ->  {"a":{"b":9,"c":8}}   (replace a.b, create a.c, one pass)
out = json.SetPaths(doc, out[:0], [][]byte{[]byte("9"), []byte("8")},
	[][]string{{"a", "b"}, {"a", "c"}})
```

Each `rawVal` is inserted verbatim and must be one well-formed JSON value; any
keys created along the way are written as plain JSON strings (no escaping, so
avoid keys needing it). `out` is filled from `out[:0]` and returned — pass a
reusable buffer to avoid allocation; `out` must not alias `in`, which is never
modified. Inter-token whitespace in `in` is preserved outside the edited spans.

## SIMD scanning

Two hot scan loops use a single vectorized pass instead of byte-at-a-time work,
with kernels in `pkg/unstable/index_amd64.s` and `pkg/unstable/index_arm64.s`
(arm64 uses NEON/ASIMD, 16 bytes/pass, for both):

- **next `"` or `\` in a string** — replaces two `bytes.IndexByte` scans; speeds
  up string-heavy payloads. On amd64 it uses SSE2 (16-byte vectors, two compares
  per 32-byte step), which avoids the `VZEROUPPER` an AVX2 routine must run on
  every call — pure overhead for the short keys and values typical of JSON.
- **next structural byte (`{`, `}`, `[`, `]`, `"`)** — AVX2 on amd64, 32
  bytes/pass, lets `skipObject` / `skipArray` jump over inert content (numbers,
  keys, whitespace) when skipping unknown values. Skipping a large ignored
  array/object is dramatically faster (the `skip-heavy` benchmark decodes at
  >50 GB/s, ~230× `encoding/json`).

### CPU requirements

The string scanner is used unconditionally on amd64 and arm64, relying only on
each architecture's **baseline, mandatory** vector ISA — no runtime gate:

- **amd64 requires SSE2.** SSE2 is part of the AMD64 ISA, so every 64-bit x86
  CPU has it (Go itself requires it); the string scanner uses it directly.
- **arm64 requires NEON / Advanced SIMD (ASIMD).** ASIMD is mandatory in the
  ARMv8-A baseline that Go targets, so every arm64 CPU has it; the string scanner
  uses it directly. (This unconditional call is also what lets the dispatch
  inline into its callers.)

The only **optional** feature is **AVX2** (amd64), used for the structural-byte
scanner and runtime-detected via `golang.org/x/sys/cpu`; without it the
structural scan falls back to scalar code. Inputs shorter than the vector width,
and platforms other than amd64/arm64, also take the scalar (`bytes.IndexByte`-
based) path. Behavior is identical across paths — verified by fuzzing each
primitive against a reference and by `SkipValue`/decode round-trips vs
`encoding/json`, on amd64 and on arm64 under qemu.

## Benchmarks

The [`bench/`](bench) directory is a separate module (so its benchmark-only
dependencies on [easyjson](https://github.com/mailru/easyjson) and
[sonic](https://github.com/bytedance/sonic) stay out of the main module). It
benchmarks the same payload decoded four ways — lightning, `encoding/json`,
easyjson, and bytedance/sonic — across each `bench/<case>/` folder.

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
| lightning (`nocopy`) | ~660 | 0 | 0 | ~13× |
| lightning (default)  | ~800 | 144 | 10 | ~10× |
| easyjson             | ~1600–1770 | 24–144 | 1–10 | ~5× |
| sonic                | ~4600 | 3380 | 40 | ~1.9× |
| `encoding/json`      | ~8250 | 920 | 17 | 1.0× |

## Layout

| Path | Description |
|---|---|
| [`main.go`](main.go) | the generator (`package main`) |
| [`pkg/unstable`](pkg/unstable) | the (unstable, do-not-import) runtime the generated decoders call into |
| [`pkg/json`](pkg/json) | small public API over the scanner (`Get`/`GetMany`/`ObjectEach`, `DecodeAny`, `UnescapeString`, `ParseFloat`, `StripDefaults`, `Set`/`SetMany`/`SetPaths`) |
| [`bench/`](bench) | benchmark module: hand-written `data.go` + `input.json` per case, plus the generated decoders, harness, and results |

Generated files (`*_unmarshal.go`, `bench/*/bench_test.go`, `bench/*/ej/`, and
the `bench/results.*` outputs) are reproducible and excluded from version
control via [`.gitignore`](.gitignore).

## Credits

Several of the hot-path techniques are borrowed from prior art:

- **[simdjson](https://github.com/simdjson/simdjson)** (Geoff Langdale and Daniel
  Lemire) and its Go port **[minio/simdjson-go](https://github.com/minio/simdjson-go)** —
  the SWAR "parse four digits at once" bit trick used in the float and integer
  scanners, and the two-`VPSHUFB` nibble-table classification that
  `indexStructuralAVX2` uses to find structural bytes.
- The float parser in `scanFloat` layers two published algorithms: William
  Clinger's fast path for exactly-representable values, then the
  **[Eisel–Lemire](https://github.com/fastfloat/fast_float)** algorithm (Michael
  Eisel and Daniel Lemire) for the rest — bit-for-bit identical to
  `strconv.ParseFloat` on the values it accepts.
- The benchmark corpus draws on
  [minio/simdjson-go](https://github.com/minio/simdjson-go/tree/master/testdata)
  and [go-json-experiment/jsonbench](https://github.com/go-json-experiment/jsonbench)
  test data, and the comparison suite measures against
  [encoding/json](https://pkg.go.dev/encoding/json),
  [mailru/easyjson](https://github.com/mailru/easyjson),
  [bytedance/sonic](https://github.com/bytedance/sonic),
  [goccy/go-json](https://github.com/goccy/go-json), and
  [go-json-experiment/json](https://github.com/go-json-experiment/json) (json/v2).
