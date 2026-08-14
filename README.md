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
([`DecodeAny`](#decoding-into-any)), check a document without decoding it
([`Valid`](#checking-validity)), and escape, unescape, or parse a JSON number
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
map root — see [Root types](#root-types)). A type that another type in the file
*uses* as a field gets an internal decode function instead of a method, emitted
by the type that reaches it — so a shared record type stays a plain struct, and a
`type recordStd Record` reflection baseline still measures `encoding/json`. Where
a group of types reference each other in a cycle that nothing else enters, every
member of that cycle gets the method, so which of them you decode into is not
decided by declaration order. The generated code imports `github.com/JohanLindvall/lightning/pkg/unstable` for the
shared scanner.

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

The `json:"..."` struct tag names the key, as with `encoding/json`, and the
[alternate-name](#alternate-field-names), [`nocopy`](#the-nocopy-tag-option),
[`lax`](#the-lax-tag-option) and [`unwrap`](#the-unwrap-tag-option) options
extend it. One naming rule is *not* shared: `encoding/json` validates the tag
name and throws the whole tag away when it holds a character outside its allowed
set — letters, digits, `!#$%&()*+-./:;<=>?@[]^_{|}~` and space — keying the
field by its **Go field name** instead. A tag holding a quote, a backslash or a
control byte (`json:"a\"b"`) is therefore one `encoding/json` ignores, while
lightning has no such rule and matches the name as written: the two then answer
to different keys in both directions, with no error on either side. `|` is in
the allowed set, so alternate names are unaffected. The generator prints a
warning naming the field; it does not reject the tag (that would fail a schema
decoding correctly today) and does not adopt the field-name fallback (that would
silently change which key an existing decoder answers to).

## Supported types

`string`, `bool`, every sized `int`/`uint` kind, `float32`/`float64`,
`json.RawMessage` (and `RawValue`), `time.Time` (RFC 3339, like `encoding/json`;
the [`lax`](#the-lax-tag-option) option also accepts a space separator and Unix
timestamps), nested named and anonymous structs, slices, fixed-size arrays
(`[N]T`), maps with string keys, pointers, and the *empty* interface —
`any`, `interface{}`, or a spelling of the same type like `interface{ any }` —
decoded into the usual Go representation of an arbitrary JSON value. Unknown
object keys are skipped.

An interface with any content of its own (a method, an embedded named interface,
a type set) is **not** supported: a decoded value is an `any` and assigns to
nothing narrower, so the generator reports the field as an unsupported type
rather than emitting a decoder that does not compile.

A fixed-size array follows `encoding/json`: the leading elements are filled, a
shorter JSON array leaves the remaining elements zero, and a longer one's extras
are discarded.

A `[]byte` (`[]uint8`) field follows `encoding/json`: it accepts both the
base64 string form the stdlib marshals (`"AQID"`) and a JSON array of numbers
(`[1,2,3]`); `null` yields nil. A fixed-size `[N]byte` accepts only the numeric
form, also like the stdlib.

One divergence comes out of the reuse above: a `[]byte` field decodes base64
straight into the backing array it already has, so a *failed* decode cannot leave
the previous value intact — the bytes that did decode have overwritten it. The
field is therefore set to the prefix that decoded, which is what the other
readers do on error too, rather than left reporting its old length over rewritten
bytes. `encoding/json` decodes into a fresh buffer and so leaves its target
untouched; a caller that needs the old value across a failed decode has to keep
its own copy.

### Reusing a decode target

Decoding repeatedly into the same value is supported and is the cheapest way to
decode a stream of documents — the slices and maps already allocated get reused, so
a steady-state decode allocates nothing for them:

```go
var v Log
for _, doc := range docs {
    if err := v.UnmarshalJSON(doc); err != nil { // reuses v's backing arrays
        return err
    }
    use(&v)
}
```

Reuse here is *storage* reuse, not *value* reuse, and that is the one place where it
parts company with `encoding/json`. Three rules cover it, and two of them carry a
divergence:

- **Struct fields are decoded in place.** A field the document omits keeps the value
  it had, exactly as with `encoding/json` — that applies to the fields of a nested
  struct and to a non-nil **pointer** field, which is decoded *into the value it
  already points to* (a nil pointer allocates; JSON `null` sets it back to nil), so
  pointer-dense schemas reuse allocation-free. Zero the target yourself if you need
  absent-means-zero.
- **An explicit `null` splits by kind, and one of those kinds diverges.** A slice,
  map, pointer or `any` field becomes nil, a `json.RawMessage` field takes the four
  bytes `null`, and a nested struct or `[N]T` field is left **exactly as it was** —
  all as in `encoding/json`. A **leaf** field, though — `string`, `bool`, integer,
  float, `json.Number`, `time.Time` — is set to its **zero value**, where
  `encoding/json` would leave it alone ("unmarshaling a JSON null into any other Go
  type has no effect on the value"). So `{"n":null}` clears a seeded `Name`, and is
  *not* the same as omitting the key. See
  [the divergence list](#differences-from-encodingjson) for why. The rule is the same
  for a field tagged `lax`, deliberately: `json:"n"` and `json:"n,lax"` never
  disagree about a null. Like the pointer rule above, none of this is observable
  unless the target is not already zero — a struct seeded with defaults, or one
  reused across documents — which is exactly the pattern this section is about.
  At the **root**, `UnmarshalJSON("null")` on a named slice or map type sets it to
  nil and on a struct type does nothing, both as in `encoding/json`.
- **Container elements are reset before they are filled.** Every slice element,
  fixed-array element and map value starts from its zero value, so nothing carries
  over from the previous document into an element. The *storage* is still reused:
  a slice keeps its **backing array** (its length is reset and it is refilled, which
  is what makes the reuse free), a map keeps its existing **entries** and merges the
  document's members over them — clear or replace a map field yourself if you need it
  to hold only the latest document — and a fixed-size array is zeroed in full.

`encoding/json` differs on the third rule too: it decodes a slice or array element
*into* whatever that element already holds, so a field absent from the later document
survives from the earlier one. With `Items []Inner` where `Inner` is `{A int; B string}`:

```go
v.UnmarshalJSON([]byte(`{"items":[{"a":1,"b":"one"}]}`))
v.UnmarshalJSON([]byte(`{"items":[{"a":2}]}`))
// lightning:      v.Items == [{A:2, B:""}]     — the element was reset first
// encoding/json:  v.Items == [{A:2, B:"one"}]  — "one" bled through from document 1
```

The same holds for `[]*Inner` (lightning gives the element a fresh pointee; the
stdlib reuses the old one) and for `[N]Inner`; map values are reset by both. The
divergence is deliberate — stale-field bleed through a reused slice is a classic
`encoding/json` footgun, and resetting the element is the counterpart of the
length reset that makes `[1,2]` decoded twice yield `[1,2]` and not `[1,2,1,2]`.
Note that the difference is invisible unless you reuse a target *and* the documents
disagree about which fields are present.

Embedded struct fields are promoted like `encoding/json`: an embedded struct's
exported fields decode as if they were the outer struct's own (an embedded
pointer is allocated on demand), a name present on both the outer struct and an
embed is resolved by Go's shallower-wins rule, an equal-depth clash is dropped
unless a single field is tagged, and an embedded field with its own JSON tag name
is a plain named field rather than promoted.

Two limits on that:

- **A struct type from another package cannot be embedded.** Its fields aren't
  visible to the generator, so there is nothing to promote, and generation
  *fails* — `type Root struct { strings.Builder; … }` reports `unsupported type
  strings.Builder`. Only the three foreign types the generator knows —
  `time.Time`, `json.RawMessage` and `json.Number` — decode when embedded, and
  they decode as a single named field keyed by the type name (`"Time"`,
  `"RawMessage"`, `"Number"`), not flattened.
- **An embedded type that has its own `UnmarshalJSON` is decoded as that named
  field, where `encoding/json` lets the promoted method take over the whole
  struct.** Go promotes the *method* as well as the fields, so the outer struct
  itself satisfies `json.Unmarshaler` and the stdlib hands it the entire
  document; lightning promotes fields only. Given `struct { time.Time; A int }`
  and `{"Time":"2021-01-02T03:04:05Z","A":7}`, lightning fills both fields,
  while `encoding/json` fails the whole decode (`Time.UnmarshalJSON: input is
  not a JSON string`) — and with an embedded `json.RawMessage` it fails
  *silently*, capturing the entire document into the embed and leaving the
  sibling fields zero. lightning's answer is usually the wanted one, but it is a
  real divergence: don't embed such a type in a schema that has to decode the
  same way under both. An embedded `json.Number` is unaffected — it carries no
  `UnmarshalJSON`, so nothing is promoted and the two agree.

## Differences from `encoding/json`

Decoding is meant to be a drop-in swap: for every supported type a well-formed
document yields the values the stdlib yields. The list below is the set of
deliberate divergences — worth reading once before migrating, because most of them
are *silent*.

- **An explicit `null` clears a leaf field instead of leaving it alone.**
  `encoding/json` documents that "unmarshaling a JSON null into any other Go type
  has no effect on the value", so `{"name":null}` leaves a seeded `Name` at its
  previous value. lightning stores the **zero value** in a `string`, `bool`,
  integer, float, `json.Number` or `time.Time` field instead — `{"name":null}`
  clears it. The other kinds match the stdlib exactly: a slice, map, pointer or
  `any` becomes nil, a `json.RawMessage` takes the literal `null`, and a nested
  struct or `[N]T` is left untouched. The reason is cost: every `*OrNull` reader
  reports a null by returning the zero value, so honoring the stdlib rule means
  testing the value's first byte at **every leaf field of every decode** to change
  the outcome only for a null that arrives at an already-populated field. Since
  the two rules coincide whenever the target starts out zeroed — a freshly
  declared value, the overwhelmingly common case — the difference is only
  reachable if you seed a target with defaults or
  [reuse](#reusing-a-decode-target) one across documents *and* the later document
  sends an explicit null. If you rely on that, keep the value yourself or compare
  against a sentinel; a field tagged `lax` follows the same rule, so the two
  spellings never disagree.
- **Key matching is exact and case-sensitive.** `encoding/json` tries an exact
  match and then falls back to a case-insensitive one, so `{"ITEMS":[…]}` fills a
  field tagged `json:"items"`. lightning has no fallback: a key that isn't
  byte-equal to the tag (or to one of its
  [alternate names](#alternate-field-names)) is simply an unknown key, and unknown
  keys are skipped. That is faster — the field switch is a `switch len(key)` over
  constant strings, with no case folding — and less surprising to read, but it
  fails quietly: a mis-cased document decodes to zero values — or, into a
  [reused target](#reusing-a-decode-target), to the *previous* document's values.
  If a source has unstable capitalization, list the spellings explicitly:
  `json:"items|Items|ITEMS"`.
- **The `,string` tag option is not implemented.** `encoding/json` uses it to read
  a number or bool that the producer wrapped in a JSON string — `{"n":"7"}` into an
  `int` field tagged `json:"n,string"`. lightning ignores the option and decodes
  the value with the field's declared Go type, so a quoted number fails with
  `ErrBadNumber`; the generator prints a warning when it sees the option, rather
  than emitting a decoder that silently does the wrong thing. Declare the field as
  a `string` and convert, or — when the string holds a whole JSON *document* rather
  than one scalar — use [`unwrap`](#the-unwrap-tag-option). `omitempty` is accepted
  and ignored, since lightning only decodes.
- **Slice, array and map elements are reset before being decoded**, where
  `encoding/json` decodes a slice or array element into whatever it already holds.
  Only observable when you reuse a decode target; see
  [Reusing a decode target](#reusing-a-decode-target) for the full rule and an
  example.
- **Out-of-range numbers wrap or saturate instead of failing.** `encoding/json`
  reports an `UnmarshalTypeError` for a number that doesn't fit the destination
  and leaves the field alone. lightning parses into the widest type and converts,
  so `300` into an `int8` stores `44`, a 20-digit integer into an `int64` wraps,
  and `1e39` into a `float32` stores `+Inf` — the same for those types inside a
  slice or array. (A *negative* number in an unsigned field is the one range error
  that is reported, as `ErrBadNumber`.) The reason is the hot integer path: a SWAR
  digit fold with no per-width range check. If out-of-range input must be
  rejected, decode into `int64`/`float64`/`json.Number` and check the range
  yourself.
- **Numbers are validated by arithmetic, not by grammar**, so `01`, `+1`, `.5` and
  `5.` are accepted while `1e309` is rejected (it has no `float64`). The table
  under [Checking validity](#checking-validity) spells this out.
- **A json tag name `encoding/json` considers invalid is honored here.** The stdlib
  validates the tag name and, when it holds a character outside its allowed set
  (letters, digits, `!#$%&()*+-./:;<=>?@[]^_{|}~` and space), throws the whole tag
  away and keys the field by its **Go field name**. lightning matches the name as
  written, so `json:"a\"b"` is filled by the key `a"b` here and by the key `Q` (the
  Go field name) there — a disagreement in both directions, silent on both sides.
  The generator warns when it sees such a tag; see [How it works](#how-it-works).
- **An embedded type with its own `UnmarshalJSON` does not take over the struct.**
  Go promotes the embedded *method* along with the fields, so `struct { time.Time;
  … }` itself satisfies `json.Unmarshaler` and `encoding/json` hands the whole
  document to `time.Time.UnmarshalJSON`, never reaching the sibling fields.
  lightning promotes fields only: the embed decodes as a named field keyed by its
  type name and the siblings decode normally. The stdlib fails loudly for
  `time.Time` and *silently* for an embedded `json.RawMessage`, which captures the
  entire document and leaves the siblings zero. See
  [Supported types](#supported-types).
- **Errors are sentinel values, not typed errors.** Every failure is one of the
  package [sentinels](#errors), matched with `errors.Is`; there is no
  `*json.UnmarshalTypeError` or `*json.SyntaxError` carrying the offending field,
  Go type and byte offset. Tolerating a type mismatch on a particular field is
  what the [`lax`](#the-lax-tag-option) option is for.
- **Invalid UTF-8 is passed through when decoding** rather than replaced with
  U+FFFD — see [Limits and untrusted input](#limits-and-untrusted-input). Pinned
  across every string path (copy, nocopy, destructive, object keys, the dynamic
  `any` decoder, `UnescapeString`) by
  `unstable.TestStringsPassInvalidUTF8Through`. (The *escape* direction —
  `EscapeString`/`EscapeStringInto`, which produce JSON — substitutes U+FFFD
  exactly as `encoding/json` does when marshaling.)
- **A `time.Time` written with `\uXXXX` escapes decodes here and fails in
  `encoding/json`.** lightning reads the JSON string's *value* and parses that;
  `time.Time.UnmarshalJSON` parses the raw bytes between the quotes without
  unescaping them ([go.dev/issue/47353](https://go.dev/issue/47353)), so
  `"2021-01-01T00:00:00Z"` — legal JSON denoting a legal instant — errors
  there and decodes to `2021-01-01 00:00:00 +0000 UTC` here. The divergence only
  ever accepts *more*, and on an escape-free timestamp the two agree on both
  acceptance and instant. lightning's authority for the grammar is
  `time.Parse(time.RFC3339, …)`, which is also what the stdlib currently reduces
  to — its extra RFC 3339 strictness is compiled out pending
  [go.dev/issue/54580](https://go.dev/issue/54580), so that agreement is a
  premise the tests check rather than an identity.
- **Duplicate keys: last wins when decoding, first wins in the toolkit.** A
  generated `UnmarshalJSON` decodes every occurrence in document order, so the last
  one survives — matching `encoding/json`. The [`pkg/json`](pkg/json) lookup and
  edit helpers (`Get`, `Lookup`, `GetMany`, `GetPaths`, `Set`, `SetMany`,
  `SetPaths`) take the **first** occurrence and leave any later one untouched,
  because they return or rewrite a single span rather than reading the object to
  its end. (`ObjectEach`/`ArrayEach` are the exception: iteration visits every
  member, duplicates included.) Duplicate keys are rare and neither answer is
  wrong, but the two halves of the library disagree, so don't build on it.

Two things that are *not* differences, though they are commonly assumed to be:

- **Trailing commas are rejected**, exactly as `encoding/json` rejects them:
  `{"a":1,}` and `[1,]` both fail with `ErrInvalidJSON`, in generated decoders,
  `DecodeAny` and [`Valid`](#checking-validity) alike — and in the value of a
  [`lax`](#the-lax-tag-option) field, whose skip is that same grammar walk. (The
  lenient bracket balancer used to *skip* unwanted values would accept them,
  which is precisely why neither `Valid` nor `lax` uses it. It remains what an
  *unknown* field's value is skipped with, where nothing downstream depends on
  those bytes.)
- **Unknown object keys are skipped**, as with `encoding/json`'s default; there is
  no `DisallowUnknownFields` equivalent.

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

### Declarations that get no method

A schema file usually holds more than roots, so a declaration no method can be
attached to is **skipped** rather than failing the run — a generic helper or a
compatibility alias sitting next to the real root should not break the build:

```go
type Pair[T any] struct{ A T }  // generic: skipped, with a warning (a method needs an instantiated type)
type Legacy = Record            // alias:   skipped, with a warning (a method needs a defined type)
type Triple [3]int              // fixed-size array root: skipped
type ID string                  // not a struct, slice or map: skipped
```

The two forms that *look* like roots — the generic type and the alias — say so on
stderr, since silently dropping something spelled like a root is the surprising
case. Skipping costs an alias nothing else: `type Legacy = struct{…}` still
decodes wherever another type uses it as a field, it just gets no method of its
own. The remaining types generate as usual; if *nothing* is left, the run fails
with `no top-level struct, slice or map types found`.

One thing here is a hard error rather than a warning: a top-level type whose
**name is one the generated code already uses** for a parameter, a local, an
import, or a predeclared identifier — `data`, `i`, `err`, `key`, `zero`, `out`,
`val`, `unstable`, `max`, `nil`, … The generated bodies would capture it (`var
zero zero`, or `new(data)` where `data` is the `[]byte` parameter), so the
generator reports `type name "data" collides with an identifier used by the
generated code; rename it` instead of writing a package that does not compile.

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
skipped but a broken one cannot. That is enforced rather than assumed: after a
failed decode the value is skipped with a walk that *parses* the grammar (the
same one [`Valid`](#checking-validity) uses) instead of a bracket balancer, so
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

### `//lightning:arena`

Documents shaped like a skeletal animation or a mesh hold *many thousands of
tiny numeric arrays* — 3-element `pos`/`rot`/`scl` triples, small index lists.
Each decodes into its own exactly-sized `[]float64`/`[]int` backing, and those
tens of thousands of small `make` calls become the decode's dominant cost (on
the `marine_ik` benchmark they are 95% of allocated objects and ~20% of CPU).
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
neighbours, backings over 512 bytes are allocated individually as before, and
reusing a target value still reuses its existing backings. Only slices of bare
`float64`/int/uint kinds participate; other field types decode exactly as
without the directive.

A schema may refer back to itself — a tree node, a comment thread, a FHIR
extension — either directly or through a chain of types:

```go
type Node struct {
    Name string  `json:"name"`
    Kids []*Node `json:"kids"` // decoded by recursion
}
```

The decoders such a type needs call each other in a loop, so decoding recurses
once per level of the document. Since a Go stack overflow is **fatal** — `recover`
cannot catch it — unbounded recursion here would mean a deeply nested document
could kill the process rather than return an error.

So the generator looks for cycles in the type graph before emitting anything. Types
on a cycle (and those that can reach one, which must pass the counter down) get
their decoders threaded with a recursion depth, and the struct decoders among them
refuse to descend past `unstable.MaxDepth`, returning `ErrMaxDepth`:

```go
// {"kids":[{"kids":[ … 4 million levels … ]}]}
err := node.UnmarshalJSON(deeplyNested)
// err is unstable.ErrMaxDepth — previously: fatal error: stack overflow
```

**A schema without a cycle is completely unaffected** — not one signature changes,
so the generated code is byte-for-byte what it was and the hot paths pay nothing.
Of the 30 benchmark cases only three (`twitter_status`, whose `twitterURL` nests
`[]twitterURL`; `golang_source`; `synthea_fhir`) contain a cycle at all, and on
those the counter costs ~1–1.5%. Nothing is required of you: the cycle is detected
from the types themselves, with no annotation.

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
`Skip`/`ReadKey` machinery the generated decoders use — non-target objects and
arrays are jumped with the [SIMD whole-container skip](#simd-scanning), so
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
  one `Get` per field (≈40–50% faster on a record with a handful of nested paths).
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

```go
// Pull a few fields out of a log record in one pass, reusing a scratch slice.
keys := []string{"ClientIP", "EdgeResponseStatus", "RayID"}
vals, err := json.GetMany(data, keys, scratch[:0])
// vals[0] == []byte(`"203.0.113.23"`), vals[1] == []byte("599"), …
```

Each function has a **compact counterpart** — `GetCompact`, `LookupCompact`,
`GetManyCompact`, `GetPathsCompact`, `ObjectEachCompact`, `ArrayEachCompact` —
with the identical signature and result. Like the
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
- A [`lax`](#the-lax-tag-option) field is the one place a generated decoder runs
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
`ErrExpectArray`, `ErrExpectColon`, `ErrMaxDepth`, `ErrBadNumber`,
`ErrBadEscape`, `ErrBadUnicode`, `ErrBadTime` — so callers can match them with
`errors.Is` without importing the internal `pkg/unstable` package.

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
`indexEscapeNEON` or an SVE2 body on arm64 — see [SIMD scanning](#simd-scanning) —
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
  — copies `input` to `output`, dropping every default value and then dropping any
  object or array that this leaves empty (which cascades: the document itself can
  come out empty). A value is a default when it is byte-equal to one of
  `defaults`, compared against the bare token — the unquoted contents for a
  string, the literal token for a number or keyword (e.g. `[]byte("none")`,
  `[]byte("false")`). An object member with a default value is dropped whole, key
  included; an array *element* that is a default is dropped from the array, which
  reindexes it (`[0,1,0,2]` with `"0"` listed becomes `[1,2]`). Empty values are
  *not* special-cased: to drop empty strings (and other empty tokens) include an
  empty entry `[]byte("")` in `defaults`. A member is kept despite a default value
  when its unquoted key is byte-equal to one of `keep` (e.g.
  `[]byte("WallTimeMs")`) — it then keeps its original value, reformatted per
  `ws`; `keep` names object keys, so it has no effect on array elements. String
  values keep their surrounding quotes and escapes; scalars keep their literal
  token. (An already-empty `{}`/`[]` is dropped too, and that is the one rule that
  ignores `defaults` — though `keep` still rescues the member holding it.)

`output` is filled from the front and the populated prefix is returned; `input`
is never modified. StripDefaults never lengthens the document, so `output` is grown
(allocated) only when `cap(output) < len(input)` — pass `input[:0]` to strip in
place, or a reused buffer to run allocation-free. Stripping in place gives exactly
the bytes a separate output buffer would: the walk never writes past its own read
cursor, so it only ever overwrites input it has already consumed. It is best effort
and copies malformed input through unchanged.

`ws` chooses how inter-token whitespace is handled:
- `RemoveWhitespace` (the zero value) — tolerate any whitespace; output is compact.
- `AssumeCompact` — assert the input has no inter-token whitespace and skip the
  `SkipWS` scans (faster, as [`GetCompact`](#key-lookups) does); misreads spaced input.
- `PreserveWhitespace` — keep the input's whitespace around surviving content, so a
  pretty-printed document stays pretty-printed; only dropped members are removed.
  One exception: the run between a value and the `,` after it goes with the
  separator (`{"a":1 , "b":2}` → `{"a":1, "b":2}`); whitespace before a closing
  `}`/`]` is kept.

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
  root is replaced by a fresh object of all the members, in place of that root
  *value* (whitespace before it and anything after it survive, as with `Set`); a
  `rawVal` shorter than `keys` ignores the surplus keys. A key listed twice in
  `keys` is set once, from its first entry, so a degenerate request can't make
  `SetMany` write a duplicate-key document.
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

Each `rawVal` is inserted verbatim and must be one well-formed JSON value; any key
created along the way is written **raw between two quotes**, with no escaping pass
— part of what keeps these functions allocation-free. So a key containing `"`, `\`
or a control byte does not round-trip, and a key crafted to close its own string
(`x":1,"role`) splices whole extra members into the result. Never build a path out
of untrusted input: use the [checked forms](#checked-edits), which reject such a
key with `ErrUnsafeKey` before the edit runs. `out` is filled from `out[:0]` and
returned — pass a reusable buffer to avoid allocation; `out` must not alias `in`,
which is never modified. Inter-token whitespace in `in` is preserved outside the
edited spans.

## Checked edits

`Set`, `SetMany`, `SetPaths` and `StripDefaults` are **best effort by design**:
they return only a `[]byte`, walk the input with a bracket balancer rather than a
parser, and on input they can't make sense of they produce *something* — a
passthrough, a partial rewrite — instead of complaining. That is what keeps them
allocation-free on the hot path, where the input is typically a document you just
produced or already validated. It also means malformed input, or a `rawVal` that
isn't a single JSON value, propagates silently into the result.

For input you don't trust, each has a `…Checked` counterpart that returns an
error. They validate their arguments with [`Valid`](#checking-validity) before the
edit, check every key they are asked to create, and validate the result
afterwards, so a bad document, a bad inserted value, an unwritable key, or a
result that somehow came out malformed is reported instead of returned:

- `SetChecked(in, out, rawVal []byte, keys []string) ([]byte, error)`
- `SetManyChecked(in, out []byte, rawVal [][]byte, keys []string) ([]byte, error)`
- `SetPathsChecked(in, out []byte, rawVal [][]byte, paths [][]string) ([]byte, error)`
- `StripDefaultsChecked(input, output []byte, defaults, keep [][]byte, ws WhitespaceMode) ([]byte, error)`

```go
out, err := json.SetChecked(untrusted, out[:0], []byte(`{"ok":true}`), []string{"result"})
if err != nil { // ErrInvalidJSON: untrusted wasn't one well-formed value
    return err
}
```

They return `ErrInvalidJSON` for a malformed document, inserted value, or result;
`ErrValueCount` when `rawVal` has fewer entries than `keys`/`paths` — the case the
unchecked forms handle by silently ignoring the surplus; and `ErrUnsafeKey` when a
key, or any element of a path (the intermediate objects are built from them too),
holds a byte that would have to be escaped. That last one is a security check
rather than a tidiness one: created keys are written raw, so `x":1,"role` closes
its own string and injects a member — and since the forgery is itself well-formed
JSON, validating the *result* would not catch it, which is why keys are checked
before the edit. On error the returned slice is `nil`, so there's no half-edited
buffer to mistake for a result.

Two things they deliberately don't check. The compactness *assertion* is still a
promise you make: `AssumeCompact` is not verified, as nowhere else in this package
verifies a `Compact` claim. And stripping a document down to **nothing** is a
documented outcome rather than an error — when every member of the root is a
default, the empty containers cascade up and `StripDefaults` returns zero bytes,
so test `len()` before forwarding the result. `PreserveWhitespace` reaches that
outcome holding the document's outer whitespace, so its raw result is those bytes
rather than none; `StripDefaultsChecked` reports it the same way in every mode —
an empty slice and a `nil` error — so `len()` stays the one test.

The cost is the extra `Valid` passes, which allocate nothing; the unchecked
functions are untouched, so code that doesn't call these pays nothing for them.

## SIMD scanning

Several hot scan loops use a single vectorized pass instead of byte-at-a-time
work, with kernels in `pkg/unstable/simd_{amd64,arm64}.s` and
`pkg/unstable/skipfast_{amd64,arm64}.s` (amd64 uses SSE2/AVX2/AVX-512, arm64
NEON/ASIMD, plus **SVE2** where the core has it):

- **next `"` or `\` in a string** — replaces two `bytes.IndexByte` scans; speeds
  up string-heavy payloads. On amd64 it uses SSE2 (16-byte vectors, two compares
  per 32-byte step), which avoids the `VZEROUPPER` an AVX2 routine must run on
  every call — pure overhead for the short keys and values typical of JSON.
- **next structural byte (`{`, `}`, `[`, `]`, `"`)** — AVX2 on amd64, 32
  bytes/pass, lets `skipObject` / `skipArray` jump over inert content (numbers,
  keys, whitespace) when skipping unknown values. Skipping a large ignored
  array/object is dramatically faster (the `skip-heavy` benchmark decodes at
  >50 GB/s, ~230× `encoding/json`).
- **whole-container skip** — when an object or array must be skipped (a non-target
  value during a `Get`, an unknown struct field, a slice element counted for
  presizing), `SkipValue` reads **64 bytes at a time**, builds an *inside-string*
  bitmask so `{`/`}`/`[`/`]`/`"` within string values are ignored, and balances
  brackets to the matching close — absorbing string contents into the bulk scan
  instead of a separate scan per string. This makes `Get` on a document where the
  wanted key sits behind many nested-object siblings about **2× faster**. Objects
  and arrays of objects/strings take this path; a flat scalar array (`[1,2,…]`)
  keeps the structural-byte scan above, which already reaches the close in one pass.

  The whole block loop is itself assembly, so the per-block state (bracket depth,
  the escape and in-string carries) stays in registers rather than crossing the
  Go/asm boundary each block: `skipBlocksAVX2`, an **AVX-512** variant that classifies
  a 64-byte block with one load and a k-mask per class, and `skipBlocksNEON` on
  arm64. Measured against the per-block Go loop: −59…−67% on object-shaped
  containers (amd64), −29…−35% (Apple M2).

### CPU requirements

The string scanner is used unconditionally on amd64 and arm64, relying only on
each architecture's **baseline, mandatory** vector ISA — no runtime gate:

- **amd64 requires SSE2.** SSE2 is part of the AMD64 ISA, so every 64-bit x86
  CPU has it (Go itself requires it); the string scanner uses it directly.
- **arm64 requires NEON / Advanced SIMD (ASIMD).** ASIMD is mandatory in the
  ARMv8-A baseline that Go targets, so every arm64 CPU has it; the string scanner
  uses it directly. The Go-side call stays unconditional even where SVE2 is used
  (the choice is made in the assembly), which is what lets the dispatch inline
  into its callers.

The **optional** features are runtime-detected via `golang.org/x/sys/cpu`. On
arm64 there is one:

- **SVE2**, on cores that have it (Neoverse N2/V2, Graviton4 and later; not Apple
  M-series or Neoverse N1/V1), replaces the NEON body of all four scanners with a
  shorter one. `WHILELO` predication covers the ragged end of the buffer, so the
  scalar tail loop and the short-input entry disappear entirely; SVE2's `MATCH`
  tests a whole character class of up to 16 bytes in one instruction (it replaces
  two compares and an OR in the string scanner, and the entire five-op nibble
  shuffle in the structural one); and because `MATCH` sets the condition flags
  directly, the two cross-domain moves NEON needs to test a block are gone. The
  bodies are vector-length agnostic, so a 256-bit implementation scans 32 bytes
  per block with the same code. Measured on a Neoverse N2: **−8…−12% end-to-end
  on string-heavy documents** (cloudflare, string_unicode), −2…−4% on the rest,
  and −34% on skipping a scalar array. The feature check is made inside the
  assembly rather than in Go, so the dispatch stays a single call and keeps
  inlining into its callers.

And on amd64:

- **AVX2** drives the structural-byte scanner and the whole-container skip. Without
  it the structural scan falls back to scalar code and `SkipValue` stays on the
  structural-byte skip.
- **AVX-512** (specifically AVX512BW), when present, selects the faster
  `skipBlocksAVX512` container-skip kernel.
- **PCLMULQDQ, BMI1 and POPCNT** gate the assembly skip loop, which uses a
  carryless multiply for the in-string prefix XOR. All three are universally
  present alongside AVX2; the gate is a correctness belt, not a real branch.

On arm64 the NEON path needs no gate and is always available; SVE2 only ever
replaces a NEON body with a faster one, so the answers are identical either way
(a differential test drives both bodies against a scalar oracle on the same
machine). Inputs shorter than the vector width, and platforms other than
amd64/arm64, take the scalar (`bytes.IndexByte`-based) path.

Behavior is identical across every path, and that is tested rather than assumed:
each primitive is fuzzed against a scalar reference, `SkipValue`/decode round-trips
are checked against `encoding/json`, and `TestSkipBlocksVariants` flips the dispatch
flags to run the Go loop, AVX2 and AVX-512 kernels *each* against a scalar oracle
over the fuzz corpus plus boundary documents (backslash runs of every parity
crossing the 64-byte block boundary at every offset, quotes on the boundary,
close-dense blocks). Verified on amd64, on Apple M2, and on arm64 under qemu.

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

which (re)generates each case's deserializers, runs any hand-written tests sitting
beside its `data.go`, and writes `bench/results.txt` and an architecture-specific
`bench/results_<goarch>.md` (so runs on different CPUs do not overwrite each
other's committed results).

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
| [`pkg/json`](pkg/json) | small public API over the scanner (`Get`/`Lookup`/`GetMany`/`GetPaths`/`ObjectEach`/`ArrayEach`, `Valid`, `DecodeAny`, `Escape`/`UnescapeString`, `ParseFloat`, `StripDefaults`, `Set`/`SetMany`/`SetPaths` and their `…Checked` forms) |
| [`bench/`](bench) | benchmark module: hand-written `data.go` + `input.json` per case, plus the generated decoders, harness, and results |

Generated files (`*_unmarshal.go`, `bench/*/bench_test.go`, `bench/*/ej/`, and
the `bench/results.*` outputs) are reproducible and excluded from version
control via [`.gitignore`](.gitignore).

## Limits and untrusted input

Worth knowing before pointing this at input you don't control:

- **Nesting is bounded at 10 000 levels** (`unstable.MaxDepth`, matching
  `encoding/json`); past that the decode returns `ErrMaxDepth` instead of
  descending. `DecodeAny`, `Valid`, `StripDefaults` and generated decoders for
  recursive schemas (below) all enforce it. `Get`/`Lookup`/`GetMany`/`GetPaths`/
  `Set` walk the document without recursing per level, so the bound does not limit
  what they can read; the value skipping they rely on is bounded as well, so no
  entry point here can be driven into a stack overflow.

  The bound exists because a Go stack overflow is a **fatal** error that `recover`
  cannot catch — so without it, deeply nested input would take the process down
  instead of returning an error. That was measurable: a 4-million-level document
  aimed at a recursive schema used to die with `fatal error: stack overflow`, and
  now reports `ErrMaxDepth`.
- **A slice presize is a hint, and it is bounded by the document's own bytes.**
  Before filling a slice the decoder cheaply estimates how many elements the JSON
  array holds — by counting commas or braces, or by sampling the first few
  elements — and sizes the backing array once instead of growing it repeatedly.
  Those scans deliberately don't parse, so bytes inside a *string value* can
  inflate the estimate: `[{"s":"{{{{…"}]` is one element that counts as thousands.
  Each counter is therefore clamped to the most elements its byte span could
  legally hold, which caps the effect at what an equally large honest document
  would already cost — measured, a 2 MB crafted body allocates 98 MB where a 2 MB
  genuinely dense array of the same type allocates 96 MB, so crafting buys about
  1.02×. What remains is the ordinary JSON-to-Go expansion ratio (a `[{},{},…]`
  array of a 144-byte struct really does need ~48× its own JSON in memory), which
  is a capacity-planning fact rather than an attack, and applies equally to
  `encoding/json`. Size your request-body limit against your widest element type,
  not against the JSON.
- **Invalid UTF-8 is passed through when decoding**, not replaced.
  `encoding/json` substitutes U+FFFD for malformed bytes; lightning returns them
  verbatim, so a decoded `string` is not guaranteed to be valid UTF-8. (Unpaired
  `\uXXXX` surrogates *are* replaced with U+FFFD, matching `encoding/json`, and
  the escape direction — `EscapeString`/`EscapeStringInto` — substitutes U+FFFD
  for malformed bytes exactly as `encoding/json` does when marshaling, so JSON
  *produced* by this package is valid UTF-8.) This is lossless rather than
  lenient, but check with `utf8.Valid` if you need the guarantee. Pinned across
  every string path by `unstable.TestStringsPassInvalidUTF8Through`.
- **A skipped number is measured, not validated.** `SkipValue` — what an unknown
  field's value, a `lax` mismatch's leftovers and the toolkit's walkers step over
  — takes a number token as a run of `[0-9.eE+-]`, so `+`, `-` and `1.2.3` are
  number *spans* to it. That is the same bracket-balancer trade this section
  already describes for containers, and the reason the `…Checked` wrappers exist:
  the typed readers and `Valid` do reject those.
- **Numbers are validated by arithmetic, not by grammar** — see the table under
  [Checking validity](#checking-validity) for what that accepts. A number that
  overflows its destination **wraps** (sized ints) or **saturates to ±Inf**
  (`float32`) instead of erroring, so a field whose range matters wants an explicit
  check; see [Differences from `encoding/json`](#differences-from-encodingjson).
- **Key matching is exact**, with no case-insensitive fallback, so a document that
  varies the capitalization of a key decodes that field to its zero value (or, on a
  reused target, to the previous document's value) with no error — same section.
- `nocopy` results alias the input buffer and `//lightning:destructive` **overwrites
  it**; neither is safe for a buffer you don't own or intend to reuse.
- `//lightning:arena` batches small numeric-slice backings into shared chunks: a
  retained slice **pins its ~4 KiB chunk**, so opt in only for decode-and-discard
  use.

## License

[MIT](LICENSE) © Johan Lindvall

## Credits

Several of the hot-path techniques are borrowed from prior art:

- **[simdjson](https://github.com/simdjson/simdjson)** (Geoff Langdale and Daniel
  Lemire) and its Go port **[minio/simdjson-go](https://github.com/minio/simdjson-go)** —
  the SWAR "parse four digits at once" bit trick used in the float and integer
  scanners, the two-`VPSHUFB` nibble-table classification that
  `indexStructuralAVX2` uses to find structural bytes, and the branchless
  escaped-quote detection + quote-mask prefix-XOR that builds the *inside-string*
  bitmask for the whole-container skip.
- **[sonic-rs](https://github.com/cloudwego/sonic-rs)** (ByteDance/CloudWeGo) and
  the **[JSONSki](https://github.com/AutomataLab/JSONSki)** paper — the SIMD
  container-skip used by `SkipValue`/`Get`: balance a container's brackets over
  the inside-string bitmask, so string contents are absorbed into the bulk scan
  rather than re-scanned per string. (The arm64 `maskBlock` movemask uses the
  well-known byte-LSB gather, `(x & 0x01…01) * 0x0102…80 >> 56`.)
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
