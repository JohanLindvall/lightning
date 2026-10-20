# Reading and parsing JSON

[Home](../README.md) · [Documentation](README.md)

Use `github.com/JohanLindvall/lightning/pkg/json` without generating a decoder. Examples use `json` as its import name.

## API at a glance

Lookup and walker results alias the input, so keep it alive and unchanged while
a result is in use. Scalar helpers describe their own ownership rules below;
`UnescapeStringCopy` provides an owned string when you need one. The
readers and walkers each have a `…Compact` twin that assumes the document holds
no insignificant whitespace, which lets them skip the inter-token scans.

| | |
|---|---|
| **Read** | [`Lookup`](#key-lookups), [`Get`](#key-lookups) (with the value's offset), [`GetMany`](#key-lookups), [`GetPaths`](#key-lookups) (several nested paths in one pass), [`ObjectEach`](#key-lookups), [`ArrayEach`](#key-lookups), [`ArrayEachIndex`](#key-lookups) |
| **Interpret** | [`KindOf`](#key-lookups), [`String`](#key-lookups), [`Bool`](#key-lookups), [`ParseInt`/`ParseUint`](#number-parsing), [`ParseFloat`](#number-parsing), [`DecodeAny`](#decoding-into-any), [`DecodeAnyNumber`](#decoding-into-any) |
| **Edit** | [`Set`](editing.md#setting-a-value), [`SetMany`](editing.md#setting-a-value), [`SetPaths`](editing.md#setting-a-value), [`StripDefaults`](editing.md#stripping-default-fields) — and a [`…Checked`](editing.md#checked-edits) form of each, for untrusted input |
| **Strings** | [`EscapeString`](#string-escaping-and-unescaping), [`EscapeStringInto`](#string-escaping-and-unescaping), [`UnescapeString`](#string-escaping-and-unescaping), [`UnescapeStringInto`](#string-escaping-and-unescaping), [`UnescapeStringCopy`](#string-escaping-and-unescaping) |
| **Validate** | [`Valid`](#checking-validity) |
| **Stream** | [`NewReader`](streaming.md#streaming) (`WithBufferSize`, `WithMaxElement`) and its `Get`, `ObjectEach`, `ArrayEach`, `Elements` (a range-over-func iterator), `Reset`, `Buffered`, `Consumed` |

Errors are matchable sentinels — `ErrKeyNotFound`, `ErrExpectObject`,
`ErrTruncated`, `ErrUnknownKey` and the rest; see [Errors](#errors).
`ErrStop` ends any of the walkers early.

`pkg/json` is the public API. [`pkg/unstable`](../pkg/unstable) is the generator's
runtime, exported only because the generated files in *your* module have to call
into it — as its name says, it is not a stable API; don't import it directly.

## Key lookups

When you only need a few values out of a document and don't want to generate (or
decode into) a struct, the [`pkg/json`](../pkg/json) package exposes the scanner's
key-lookup primitives. They walk the input with the same allocation-free
`Skip`/`ReadKey` machinery the generated decoders use — non-target objects and
arrays are jumped with the [SIMD whole-container skip](performance.md#simd-scanning), so
reaching a key behind many siblings in a large document stays fast — and every
value they return aliases the input `[]byte` (keep it unchanged while the result
is in use). A returned value follows the same conventions throughout: a string keeps
its surrounding quotes with escapes intact, an object or array spans the whole
`{`…`}` or `[`…`]`, and a scalar is the literal token.

- `Get(data []byte, keys ...string) ([]byte, int, error)` — walks the object-key
  path `keys` one level per key and returns the value's raw bytes (and the offset
  it starts at), without reporting a value type. With no keys it returns the whole
  root value; a missing key returns `ErrKeyNotFound`.
- `Lookup(data []byte, keys ...string) ([]byte, error)` — `Get` without the
  offset return, for the common read-only case where the value's position in
  `data` isn't needed (use `Get` when it is, e.g. to splice the value back in).
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
  one `Get` per field.
  Same result conventions as `GetMany`; a `nil`/empty path selects the root. What a
  path resolves to never depends on what else was requested: descending for a
  deeper path is stricter than skipping a value, so on malformed input a subtree
  only a co-requested path enters can fail to be walked — that is reported as the
  call's error, but the other paths still get the values `Get` would have given
  them, and `out` comes back filled as far as the walk got.
- `ObjectEach(data []byte, fn func(key string, value []byte) error, keys ...string) error`
  — calls `fn` for every member of the object reached by the path `keys` (the
  root object with no keys). If `fn` returns an error, iteration stops and
  returns it.
- `ArrayEach(data []byte, fn func(value []byte) error, keys ...string) error`
  — the array counterpart of `ObjectEach`: calls `fn` for every element of the
  array reached by the path `keys` (the root array with no keys), same
  error-stops-iteration contract.
- `ArrayEachIndex(data []byte, fn func(index int, value []byte) error, keys ...string) error`
  — `ArrayEach` with the element's position handed along, `0` for the first.
  A fixed-shape array — the `[timestamp, value]` pair of a time series, a
  `[key, value]` tuple — reads by position without the caller keeping a
  counter in the closure.
- `KindOf(raw []byte) Kind` — says what a returned value (or a whole
  document) is: `KindNull`, `KindBool`, `KindNumber`, `KindString`,
  `KindArray` or `KindObject`, or `KindInvalid` for an empty input or a byte no
  value starts with. The read functions decide this while scanning and do not
  report it, so without it every caller branched on `raw[0]` itself. A literal
  is matched whole (`nul` and `truex` are `KindInvalid`); a string, number,
  array or object is classified by its opening byte alone, as the scanner
  dispatches — `KindOf` says what a value *is*,
  [`Valid`](#checking-validity) says whether it is well-formed. Whitespace
  around the value is tolerated at both ends, by this package's rule (every
  byte `<= 0x20`, what `Get` tolerates at a document's root), so a value
  `Valid` accepts is one `KindOf` classifies.

A value those functions return is a raw token, and the scalar readers turn
one into its Go value without the caller re-deriving the grammar:

- `String(raw []byte) (string, error)` — decodes a string token (`"…"`, quotes
  on, escapes intact) with `UnescapeString`'s contract: an escape-free token
  yields a string that aliases `raw` (keep it unchanged while the result is in
  use, or clone it when it outlives the input), an escaped one a fresh string.
  Anything that is not a string token — a number, `null`, a container, an
  unterminated string — is `ErrExpectString`; a bad escape inside the quotes is
  `ErrBadEscape`/`ErrBadUnicode`.
- `Bool(raw []byte) (bool, error)` — reads `true` or `false`; anything else,
  `null` included, is `ErrExpectBool`.
- A number is read with [`ParseFloat`](#number-parsing).

Both walkers treat a JSON `null` in the container's place as a container with
nothing in it — `fn` is not called and the call returns `nil` — which is what
unmarshalling `null` into a map or a slice does (it leaves them `nil`) and what
`{"metric":null}` or `"values":null` mean on the wire. The literal must end at
a token boundary (whitespace by the rule above, a comma, a closing bracket, or
the end of the input), so any other non-container value there, a misspelt
literal included, is still `ErrExpectObject`/`ErrExpectArray`.

A callback that returns `ErrStop` ends any of these walks early **without**
making it a failure: the walker returns `nil`. Any other error stops the walk
too and is returned as it is. For example, stop at the first matching element:

```go
var first []byte
err := json.ArrayEach(data, func(v []byte) error {
    if matches(v) {
        first = v
        return json.ErrStop
    }
    return nil
})
```

```go
// Pull a few fields out of a log record in one pass, reusing a scratch slice.
keys := []string{"ClientIP", "EdgeResponseStatus", "RayID"}
vals, err := json.GetMany(data, keys, scratch[:0])
// vals[0] == []byte(`"203.0.113.23"`), vals[1] == []byte("599"), …
```

Each function has a **compact counterpart** — `GetCompact`, `LookupCompact`,
`GetManyCompact`, `GetPathsCompact`, `ObjectEachCompact`, `ArrayEachCompact`,
`ArrayEachIndexCompact` —
with the identical signature and result. Like the
[`//lightning:compact`](options.md#lightningcompact) directive, they assume the input has
no whitespace *between* tokens (the form `encoding/json`'s `Marshal` and most
wire protocols emit) and skip the inter-token `SkipWS` scans; whitespace surrounding the whole document is still tolerated. Feed one
input that does contain inter-token whitespace and it may return an error, so use
them only for sources guaranteed compact.

## Decoding into `any`

When a document's shape isn't known ahead of time — too variable to model, or
genuinely arbitrary — `DecodeAny` reads the whole thing into the generic
representation `encoding/json` produces for `interface{}`: `nil`, `bool`,
`float64`, `string`, `[]any`, and `map[string]any`. It is the dynamic
counterpart to a generated unmarshaler, using the same scanner but with no target
type.

- `DecodeAny(data []byte) (any, error)` — decodes the single JSON value in
  `data`. Whitespace around the whole document is tolerated; trailing
  non-whitespace content after the value is an error. Unlike the key-lookup
  helpers it builds Go values (so strings are unescaped and copied, numbers
  parsed), allocating the maps and slices the result needs.
- `DecodeAnyCompact(data []byte) (any, error)` — the same for compact input,
  skipping the inter-token `SkipWS` scans (as [`GetCompact`](#key-lookups) does),
  faster on minified input; may error if the input does contain inter-token
  whitespace.

A document nested deeper than 10 000 levels is rejected with `ErrMaxDepth` rather
than descended into — the same bound `encoding/json` applies, and for the same
reason: `DecodeAny` recurses once per nesting level, and a Go stack overflow is a
*fatal* error that `recover` cannot catch, so without the bound a hostile document
would take the process down instead of returning an error. `Get`, `Set` and value
skipping walk the document without recursing per level, so the bound does not
constrain them (the one recursive skip fallback carries the same limit, so nothing
here can be driven into a stack overflow either).

## Checking validity

`Valid(data []byte) bool` reports whether `data` is one well-formed JSON value
(optionally whitespace-surrounded, no trailing content), checking the document in
place — **no allocation, and no decoded value built**.

`Valid` answers *"will lightning's own scanner accept this?"*, which is a
deliberately different question from `encoding/json.Valid`'s *"does this match the
JSON grammar?"* — the decoder's scanners are tuned for input already known to be
JSON, so the two disagree in a few places. What "lightning's own scanner" means is
worth being precise about, since it is what makes `Valid` useful as a gate:

- Against [`DecodeAny`](#decoding-into-any) the two are **held equal by test**: a
  20-million-execution differential fuzz runs both over the same inputs and
  requires them to agree, so `Valid` is exactly `DecodeAny`'s precondition. It
  holds because `Valid` reuses the decoder's own readers instead of
  reimplementing them.
- Against a generated `UnmarshalJSON` it is a **guide, not an equivalence**, and
  nothing tests it as one. A generated decoder shares these scanners — so every
  divergence in the table below applies to it too — and then adds the schema,
  which cuts both ways. It is *stricter* where the schema looks: a value `Valid`
  accepts is rejected when its type doesn't fit the field it lands in. And it is
  *looser* in two ways. Looser where the schema doesn't look: an unknown field's
  value is jumped with the lenient bracket balancer rather than parsed, so
  `{"zzz":[1,,2],"i8":5}` decodes fine into a struct with no `zzz` field even
  though `Valid` rejects it. And — the surprising one — looser where it *does*
  look, on integer fields: the integer readers consume a run of digit-ish bytes
  and stop, so `{"n":1.2.3}` decodes `n` as `1` with no error while `Valid`
  rejects the document. Read `Valid` as "this is well-formed JSON as lightning
  defines it", not as "every generated decoder will accept this".
- A [`lax`](options.md#the-lax-tag-option) field is the one place a generated decoder runs
  *this exact walk*: when the typed decode of such a field fails, the value is
  skipped by the same grammar check rather than by the bracket balancer, which is
  what makes "lax tolerates type mismatches, not syntax errors" true. So the
  "unknown field" leniency above stops at a `lax` field's value.

The divergences from `encoding/json.Valid` are these:

| Input | `Valid` | `encoding/json.Valid` | Why |
|---|:--:|:--:|---|
| `01`, `+1`, `.5`, `5.` | ✅ | ❌ | `scanFloat` does arithmetic, not grammar checking |
| `"…<0x09>…"` (raw control byte) | ✅ | ❌ | the string scanners stop at `"` and `\` only |
| `{"a":<0x00>1}` | ✅ | ❌ | inter-token whitespace is every byte `<= 0x20` — one compare, not four |
| `1e309` | ❌ | ✅ | overflows `float64`, so the decoder cannot represent it |

Everything else is checked strictly, as the decoder checks it: trailing commas,
non-string object keys, missing colons, unknown escapes, a `\u` without four hex
digits, unterminated strings, mismatched brackets, bad keywords, and trailing
bytes are all rejected. Semantics are not checked — duplicate keys are accepted,
and neither UTF-8 well-formedness nor surrogate pairing is verified (an unpaired
`\uD800` decodes to U+FFFD rather than failing, as in `encoding/json`).

Because it checks the document in place rather than decoding it, it is also
considerably faster than the standard library's validator on the same bytes:

| | ns/op | B/op | allocs/op |
|---|--:|--:|--:|
| `json.Valid` | ~1220 | 0 | 0 |
| `encoding/json.Valid` | ~2760 | 0 | 0 |

(1.1 KB record, amd64 — ~2.3× faster, and 6.6× faster than checking a document by
running `DecodeAny` and discarding the result, which also costs ~200 allocations.
`Valid` does *not* use the lenient bracket-balancing `SkipValue` path, which would
accept a trailing comma.)

## Errors

Errors returned by these helpers are the package's exported sentinels —
`ErrKeyNotFound`, `ErrInvalidJSON`, `ErrTruncated`, `ErrExpectObject`,
`ErrExpectArray`, `ErrExpectString`, `ErrExpectBool`, `ErrExpectColon`,
`ErrMaxDepth`, `ErrBadNumber`, `ErrBadEscape`, `ErrBadUnicode`, `ErrBadTime` —
so callers can match them with `errors.Is` without importing the internal
`pkg/unstable` package. `ErrStop` is the one that is not a failure: returned
from a walker's callback, it ends the walk with a `nil` result.
[Streaming](streaming.md#streaming) adds `ErrElementTooLarge`, and returns whatever the
underlying `io.Reader` reported wrapped, so `errors.Is` reaches that too.

## String escaping and unescaping

The [`pkg/json`](../pkg/json) package exposes the scanner's string codec on its
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

- `UnescapeStringCopy(in []byte) (string, error)` — the same decode, for a
  result that **outlives** its input: the returned string is always the
  caller's own, never a window onto `in`. Aliasing is right for a value read
  and dropped inside one call and a hazard for one that is kept — a label
  value or a cached row that aliases a response body pins the whole body for
  as long as anything holds it — and until this form a caller in that position
  wrote `string(in)` on the escape-free path itself and lost the library's
  escape detection doing it. One copy with no escapes; the same fresh
  allocation `UnescapeString` makes with them.

**Escaping** (raw value → escaped body, escaping `"`, `\`, and control bytes;
`/` is left as-is and `\b`/`\f` are written in `\u00XX` form. Bytes that are not
part of a well-formed UTF-8 sequence are replaced with U+FFFD, as `encoding/json`
does when marshaling — a JSON text must be valid UTF-8 (RFC 8259), and decoders
substitute U+FFFD rather than error, so passing raw bytes through would become
silent corruption downstream; the replacement is written as the raw three-byte
UTF-8 encoding, not the `\ufffd` escape):

- `EscapeString(s []byte, out *strings.Builder)` — writes the escaped form of
  `s` to `out`. The common no-escape case is detected with a vectorized scan and
  written straight to the builder, with no scratch buffer or copy.
- `EscapeStringInto(s, out []byte) []byte` — appends the escaped form of `s` to
  `out` and returns the extended slice, allocating nothing when `out` has room
  (escaping can grow a string up to 6× for control bytes). Pass `out[:0]` to
  reuse a buffer across calls. `out` must not overlap `s`: escaping has no
  in-place form, because every escape writes more bytes than it consumes and the
  appends would run over input the scan hasn't read yet. (That is the one place
  the two directions differ — `UnescapeStringInto` takes `out == in[:0]` precisely
  because unescaping only shrinks.)

Clean runs are skipped with a vectorized scan (`indexEscapeSSE2` on amd64,
`indexEscapeNEON` or an SVE2 body on arm64 — see [SIMD scanning](performance.md#simd-scanning) —
which classify `"`, `\` and control bytes in one pass), so a string needing little
or no escaping costs roughly one `memcpy`. The scanner is chosen per run by how
much input is left: a run shorter than 48 bytes — every short string, and every
short gap between two escapes — is walked a word at a time with SWAR instead, since
the vector call's setup would cost more than it saves there.

The UTF-8 handling rides in that same scan rather than costing a pass of its own:
the walk runs on a variant widened by non-ASCII bytes until the first one decides
the rest of the input once with `utf8.Valid` — valid text (unicode prose)
continues under the plain scanner so multibyte characters don't break clean runs,
and only actually ill-formed input takes the U+FFFD substitution walk. A
pure-ASCII string never pays more than the widened predicate, which costs the
same per word as the plain one.

## Number parsing

The [`pkg/json`](../pkg/json) package also exposes the scanner's float parser:

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
- `ParseInt(b []byte) (int64, error)` and `ParseUint(b []byte) (uint64, error)`
  — the integer counterparts, for the value `Get` or `ObjectEach` hands back
  when it is an integer, or for a quoted 64-bit integer once its quotes are
  stripped (ClickHouse quotes every `Int64`/`UInt64` so it survives
  JavaScript's 53-bit numbers; without these the only way to read such a cell
  was to allocate a string for `strconv.ParseInt`). Same grammar as
  `ParseFloat` where it applies — an optional sign, `+` included, leading zeros
  accepted — and strictly an integer: a fraction or an exponent (`1.0`, `1e3`),
  an empty input, stray bytes and a magnitude outside the type are all
  `ErrBadNumber`, never a truncated value. `ParseUint` takes a `+` and refuses
  a `-`. Held to `strconv.ParseInt`/`ParseUint` over a generated corpus, the
  `+` on `ParseUint` being the one documented difference.

  Both fold the digits eight at a time rather than one multiply-add per digit,
  which they can because they are handed the whole token: the digit count is
  known before a byte is read, so the per-digit overflow test disappears (only
  a 20-digit token needs a checked multiply) and the folds of the separate
  words issue in parallel. A thirteen-digit timestamp costs 4.3 ns and a
  nineteen-digit id 5.6 ns, against 9.7 and 13.4 for the digit loop.
