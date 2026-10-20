# Tags and directives

[Home](../README.md) · [Documentation](README.md)

Use field tags for local choices and `//lightning:` directives for an entire root's generated decoders.

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

Those escaped strings do not get one allocation each, though. A decoded string
is handed out aliasing the buffer it was unescaped into, so that buffer *is* the
string, and lightning carves them from a shared chunk rather than making one per
string. Chunks are capped at 16 KiB and sized to the document, with a 1 KiB
minimum. A string that outlives its document can keep its chunk alive. Bodies
larger than 4 KiB receive their own allocation; escape-free strings do not use
the chunk.

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
escape-free input decodes byte-identically to plain `nocopy`.

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
mismatch a no-op instead: the offending value is skipped and the field is left
unchanged (zero for a fresh target), while the rest of the object decodes normally:

```go
type Log struct {
    Status int64 `json:"Status,lax"` // a non-number Status is ignored, leaving 0
}
```

Only type mismatches are tolerated; malformed JSON (a syntax error in
the value) still fails, since a well-formed value of the wrong type can be
skipped but a broken one cannot. That is enforced rather than assumed: after a
failed decode the value is skipped with a walk that *parses* the grammar (the
same one [`Valid`](json.md#checking-validity) uses) instead of a bracket balancer, so
`{"l":[1,]}`, `{"l":[1 2 3]}` and `{"l":[1,,2]}` into a `[]int` field tagged
`lax` fail with `ErrInvalidJSON` exactly as they do without the tag. The
tolerance covers a mismatch nested anywhere inside an otherwise well-formed
value, which is the point of the option: `["a"]` into `[]int`, or `{"name":5}`
into a struct, is still swallowed.

`lax` works for every field type, including nested structs, slices, and maps,
where a decode error anywhere in the value leaves the whole field unset. It
combines with the other options and with alternate names —
`json:"Name|Title,nocopy,lax"`.

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

"A fresh input" holds at both ends: the body must be one JSON value and nothing
more, so content after it fails with `ErrInvalidJSON` exactly as it would at the
root — `"{\"id\":7} trailing garbage"` is an error, not an `{"id":7}` with the
rest ignored. Whitespace around the value is fine, and a body that is empty or
all whitespace still leaves the field at its zero value without an error.

## The `number` tag option

An `any` field decodes a JSON number to a `float64`, as `encoding/json` does by
default. Add `number` to the field's json tag and every number inside it — at
any depth of the value's objects and arrays — is a `json.Number` holding its
literal instead, which is what `encoding/json`'s `UseNumber` gives: `0.95`
survives as written, and an integer past 2^53 is not rounded.

```go
type Param struct {
    Default any `json:"default,omitempty,number"` // 0.95 stays "0.95"
}
```

It applies to a field that IS `any` (or `interface{}`); on any other field it
warns and is ignored. The toolkit has the same mode as `DecodeAnyNumber` (and
`DecodeAnyNumberCompact`).

## The `reader` tag option

Use `reader=readDate` to decode a field through a function you supply. This is
useful for sharing repeated strings, handling an application-specific value,
or using a reader without defining a new field type:

```go
type DayCost struct {
    Date string  `json:"date,reader=readDate"`
    Cost float64 `json:"cost"`
}

func readDate(data []byte, i int) (string, int, error) {
    return unstable.ReadStringOrNull(data, i)
}
```

The example uses `github.com/JohanLindvall/lightning/pkg/unstable`'s normal
string reader; the function can instead look up an owned string in a bounded
cache and avoid copying repeated values. The generator calls it directly,
without first scanning or copying the value's span.

The signature is `func([]byte, int) (T, int, error)`, with `T` assignable to
the field. The input index is the first byte of the value, after whitespace.
Return the absolute index immediately after the complete value, or the error
position on failure. The reader is responsible for valid offsets, parsing the
whole value, null handling, and copying any input it retains. The field is
assigned only on success; an absent key never calls the reader. A reader for
a pointer, slice, or struct returns that whole field's value, and takes
precedence over the type's `UnmarshalJSON` method.

The name can be a function in the same package or a qualified function from
an imported package, such as `readers.ReadDate`. A tag alone does not count
as a Go use of an import; a local wrapper is often convenient. Tagged fields
in sibling types, nested records, and embedded structs work too. A reader
on an embedded field decodes it under its type name (or explicit JSON name),
instead of promoting its fields.

`reader` composes with alternate names, `lax`, and `unwrap`: `lax` leaves the
field unchanged on a reader error only if the value is valid JSON, and
`unwrap` calls the reader on the unwrapped document. `nocopy` and `number`
cannot be combined with `reader`; the function controls those choices itself,
including under a `destructive` root. Duplicate reader options, invalid names,
and names shadowed by generated locals are generation errors. The Go compiler
checks the function's existence and signature.

## Comment directives

Some behavior is selected with a `//lightning:<name>` comment on the struct type
(or its declaration), separate from the per-field json tags above.

### `//lightning:emptyslice`

Preserve the distinction between an empty JSON array and null:

```go
//lightning:emptyslice
type Export struct {
    Days []DayCost `json:"days"`
}
```

`{"days":[]}` produces a non-nil empty slice, `{"days":null}` sets it to
nil, and `{}` leaves it unchanged. This applies throughout the root's generated
decoders: nested structs, slice elements, map values, named slice roots, and
batched numeric, byte, and coordinate slices. A reused slice keeps its backing
array and capacity, including for `[]`. Populated arrays retain their usual
allocation strategy. Custom readers and delegated `UnmarshalJSON` methods
remain responsible for their own results.

The directive composes with the other root directives. A shared nested type
follows the directive of the root decoding it. Without it, a nil slice stays
nil when decoded from `[]`.

### `//lightning:strict`

A member no field answers to is skipped by default, as `encoding/json` skips
it. Mark a root `//lightning:strict` and such a member fails the decode with
`ErrUnknownKey` instead — the `DisallowUnknownFields` posture, for a schema
where a misspelled key must be an error rather than a silently zero field —
at every object the root reaches, nested types included (a nested type
shared with a non-strict root gets a decoder of each kind, as it does for
the other directives). Maps are unaffected: every member of a map is a key.
The error is an `*UnknownKeyError` naming the member (`Key`), and it
matches `ErrUnknownKey` under `errors.Is` for a caller that only asks
whether the refusal was that.

```go
//lightning:strict
type Config struct {
    Listen string `json:"listen"`
}
```

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
— while each value/element type keeps its own field tags. Only slice and map
roots take the directive; a struct root uses per-field `nocopy` tags.

### `//lightning:arena`

Documents shaped like a skeletal animation or a mesh hold *many thousands of
tiny numeric arrays* — 3-element `pos`/`rot`/`scl` triples, small index lists.
Each decodes into its own exactly-sized `[]float64`/`[]int` backing, and those
small `make` calls can account for much of the allocation work.
Mark the root `//lightning:arena` to batch them: each `UnmarshalJSON` call
carves small numeric-slice backings out of shared 4 KiB arena chunks, turning
tens of thousands of allocations into a few hundred, with no change to what is
decoded:

```go
//lightning:arena
type Animation struct {
    Keys []Key `json:"keys"`
}
type Key struct {
    Pos []float64 `json:"pos"` // backing carved from the decode's arena
    Rot []float64 `json:"rot"`
}
```

The trade-off is retention granularity: a carved slice keeps its whole chunk
reachable, so holding on to one 3-element slice pins ~4 KiB. Use the directive
when you decode, process, and discard the result together (the natural shape
for these documents); don't use it if you retain a few small slices from a
large decode. Everything else is unchanged: results are ordinary slices,
appending to one reallocates it onto the heap without disturbing its
neighbours, backings over 512 bytes are allocated individually, and
reusing a target value still reuses its existing backings. Only slices of bare
`float64`/int/uint kinds participate; other field types decode exactly as
without the directive.

The chunks come from [`github.com/JohanLindvall/arena`](https://github.com/JohanLindvall/arena),
whose `Arena[T]` is typed — its chunks are `[]T` — so a schema storing several
element kinds gets one arena per kind, held in a small struct the generator
emits beside the method and threaded through that root's decoders. Being typed
is what makes the backings ordinary `[]T` the collector scans correctly, with no
pointer reinterpretation anywhere in the path.
