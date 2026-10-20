# Generator guide

[Home](../README.md) · [Documentation](README.md)

Install the generator, choose root types, and understand how generated decoders handle your Go types.

## API at a glance

```sh
go run github.com/JohanLindvall/lightning@latest path/to/data.go
```

writes `path/to/data_unmarshal.go` with an `UnmarshalJSON` for each top-level
type in the file. Behaviour is steered by struct tags and per-type comment
directives:

| Struct tag option | Effect |
|---|---|
| `json:"name"` | names the key, as in `encoding/json` |
| `json:"a\|b\|c"` | [alternate names](options.md#alternate-field-names) — any listed name selects the field; the last occurrence in the document wins |
| `json:"-"` | skip the field |
| `,nocopy` | [alias the string into the input buffer](options.md#the-nocopy-tag-option) instead of copying it |
| `,lax` | [tolerate a value of the wrong shape](options.md#the-lax-tag-option) and leave the field unchanged |
| `,unwrap` | [decode a JSON document embedded in a string](options.md#the-unwrap-tag-option) (or base64) |
| `,reader=func` | [call a custom field reader](options.md#the-reader-tag-option) |
| `,number` | [decode an `any` field with `json.Number`](options.md#the-number-tag-option) rather than `float64` |
| `,omitempty` `,omitzero` | accepted and ignored (they are encode-only) |

| Comment directive | Effect |
|---|---|
| [`//lightning:emptyslice`](options.md#lightningemptyslice) | an empty JSON array produces a non-nil slice and retains reused backing arrays |
| [`//lightning:strict`](options.md#lightningstrict) | an unknown object key is an error (`ErrUnknownKey`), not a skip |
| [`//lightning:compact`](options.md#lightningcompact) | the input has no insignificant whitespace; elide every inter-token skip |
| [`//lightning:nocopy`](options.md#lightningnocopy) | a slice or map root aliases its keys and elements |
| [`//lightning:destructive`](options.md#the-lightningdestructive-directive) | unescape strings into the input buffer — zero allocation, destroys the input |
| [`//lightning:arena`](options.md#lightningarena) | carve small scalar-slice backings from chunks instead of one `make` each |
| [`//lightning:root`](#root-types) | make this a root — a method of its own for a type something else reaches, or one defined over a struct, slice or map |

## Installation

Run it straight from the module path (no clone needed):

```sh
go run github.com/JohanLindvall/lightning@latest path/to/data.go
```

Replace `@latest` with a published tag to pin a specific generator version.
Or install the binary once:

```sh
go install github.com/JohanLindvall/lightning@latest
lightning path/to/data.go
```

The generated code imports `github.com/JohanLindvall/lightning/pkg/unstable`, so
the module you generate into must depend on lightning:

```sh
go get github.com/JohanLindvall/lightning@latest
```

The import is `pkg/unstable`, the generator's runtime — **not a stable API**,
and not one to import yourself; [`pkg/json`](../pkg/json) is the public toolkit.
It pulls in two small dependencies of its own: `golang.org/x/sys` (CPU feature
detection for the SIMD scanners) and `github.com/JohanLindvall/arena` (the
chunk-backed store behind [`//lightning:arena`](options.md#lightningarena)).

For repeatable builds, run the generator version selected by your module's
`go.mod`:

```go
//go:generate go run github.com/JohanLindvall/lightning $GOFILE
```

Keep the generator and runtime at the same module version, and regenerate after
upgrading. Commit generated decoders if downstream builds will not run generation.
After generating for the first time, run `go mod tidy` to record the dependencies
imported by the generated code.

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

**Types declared in the package's other files resolve too.** The generator
reads every sibling `.go` file of the input (same package; not a test file,
not a generated `_unmarshal.go`) for its struct, slice, map and defined
scalar types, so a root in the input may name a record kept where it is
used. A sibling type gets no method from this run — the root that reaches
it emits its decoder, exactly as for an in-file nested type — under the
reaching root's directives; a `//lightning:` directive the sibling type
carries belongs to the run its own file is the input of, and is not
reported here. To give it a method of its own, generate its file. One
rule: a sibling that imports `encoding/json` or `time` under a different
alias than the input file — or any other package under a qualifier the
input file gives a different package — is skipped whole, with a warning,
because the generated file imports every package a decoder names under the
input's qualifiers and prints every type expression as written.

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
with an index-based scanner and no reflection. String fields copy by default;
use `nocopy` when their results can safely alias the input.

The `json:"..."` struct tag names the key, as with `encoding/json`, and the
[alternate-name](options.md#alternate-field-names), [`nocopy`](options.md#the-nocopy-tag-option),
[`lax`](options.md#the-lax-tag-option) and [`unwrap`](options.md#the-unwrap-tag-option) options
extend it. Lightning honors tag names as written and warns about names the
standard library may treat differently, such as a name containing a quote.
The standard library's response depends on the toolchain: it may fall back to
the Go field name or ignore the field. See the
[compatibility guide](compatibility.md#differences-from-encodingjson) before
sharing such a schema between implementations.

## Supported types

`string`, `bool`, every sized `int`/`uint` kind, `float32`/`float64`,
`json.RawMessage` (and `RawValue`), `time.Time` (RFC 3339, like `encoding/json`;
the [`lax`](options.md#the-lax-tag-option) option also accepts a space separator and Unix
timestamps), nested named and anonymous structs, slices, fixed-size arrays
(`[N]T`), maps keyed by a string or an integer kind (or a type defined over
one — an integer key is the member name parsed, as `encoding/json` writes it,
and a name that is not a number fails with `ErrBadNumber`), pointers, and the *empty* interface —
`any`, `interface{}`, or a spelling of the same type like `interface{ any }` —
decoded into the usual Go representation of an arbitrary JSON value. Unknown
object keys are skipped.

A **defined scalar type** — `type Severity string`, `type Level int32`, the
enum idiom, or a type defined over another such type — decodes as its
underlying kind and converts, wherever a scalar can stand: a field, a slice
or array element, a map value, a pointee. Its null and `nocopy` rules are
the underlying kind's exactly. The declaration order does not matter
(`type Chain Sev` may precede `type Sev string`). Such a type gets no
`UnmarshalJSON` of its own — there is no document it could be the root of —
see [Declarations that get no method](#declarations-that-get-no-method).

A type with an `UnmarshalJSON` of its own — declared in the input file or in
a sibling file of the package — is **delegated to** wherever it stands as a
field, an element, a map value or a pointee: the value's span (a JSON null
included, as `encoding/json` passes it) is handed to the method, and its
error is the decode's error. Such a type is never generated for; a root that
declares one is skipped with a warning, since a second method would not
compile and the hand-written one is the author's answer to a shape a
structural decode gets wrong. A type from **another package** is never
looked inside: `time.Time`, `json.RawMessage` and `json.Number` decode as
themselves, and any other is delegated to its own `UnmarshalJSON` the same
way — one that does not implement `json.Unmarshaler` fails to compile
("has no field or method UnmarshalJSON"), which is the generator's usual
assertion. An *embedded* type's method does not take over the struct it
is embedded in — see the [compatibility guide](compatibility.md).

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

One divergence comes from buffer reuse: a `[]byte` field decodes base64
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
  [the divergence list](compatibility.md#differences-from-encodingjson) for why. The rule is the same
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

- **A struct type from another package cannot be promoted.** Its fields aren't
  visible to the generator, so there is nothing to promote, and generation
  *fails* — `type Root struct { strings.Builder; … }` reports `unsupported type
  strings.Builder` (name it as a field instead, and it is delegated to its own
  `UnmarshalJSON`, or give the embedded field a [custom reader](options.md#the-reader-tag-option)).
  Without a custom reader, only the three foreign types the generator knows —
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
  same way under both. Whether an embedded `json.Number` exposes an unmarshaling
  method depends on the Go toolchain; the conformance tests compare against the
  standard library being used.

## Root types

A top-level type need not be a struct: a named slice or a string-keyed map at the
root gets its own `UnmarshalJSON` too, so a document that is a bare array or a
data map needs no wrapper struct.

```go
type Records []Record          // a JSON array at the root
type ByID    map[string]Record // a JSON object used as a data map
```

A type another root reaches gets no method of its own — the reaching root
emits its decoder — **unless it carries a `//lightning:` directive**: then it
is a root as well, with a method under its own directives, while the copy
inside the reaching root keeps the root's. That is how a record nested in a
schema is also decoded on its own: a dashboard's target arriving alone in an
API request, or a filter that another package's decoder delegates to and so
needs a method on. The bare `//lightning:root` asks for nothing else.

A type **defined over** a struct, slice or map type — `type raw Rule` — is a
root too, with the underlying type's shape, when it carries a `//lightning:`
directive: `//lightning:strict` or any other, or the bare `//lightning:root`
where no other applies. That is the idiom a hand-written `UnmarshalJSON`
uses to decode its own fields without recursing into itself: `Rule` keeps
the method it wrote, `raw` gets the generated one, and the two cannot
collide. The definition may precede or follow what it is defined over, and
may run through a chain. It is opt-in because the same spelling is the
methodless twin — `type rootStd Root`, the reflection-only baseline the
benchmarks and the `encoding/json` comparisons rest on — which must never
grow a method.

`type Records []Record` decodes a top-level `[…]` with the slice element rules;
`type ByID map[string]Record` decodes a top-level `{…}` as a map, its keys the
object's member names. Either element/value type, and any nested types and field
options, behave exactly as the same type used for a struct field would. Several
root types (struct, slice, map, in any mix) can live in one input file. A named
slice or map type may also be a **field** (`Items Items`, `*Items`, an element
of another slice): it decodes through the element type's own decoder with the
destination converted to the underlying type, and a null nils it, exactly as
the bare `[]Item` or `map[string]Item` would. For a
root that is a *bare* `any`/`interface{}` — whose schema you don't know at all —
there is no method to generate (Go forbids methods on interface types); decode it
dynamically with [`json.DecodeAny`](json.md#decoding-into-any) instead.

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

A defined scalar type (`type Severity string`) is the fourth kind: it is
registered so fields of that type resolve, and a `//lightning:` directive on
it warns, since there is nothing for the directive to govern.

## Generated function names

The `UnmarshalJSON` methods keep their exact name (the `json.Unmarshaler`
interface requires it). The unexported decoder helpers they call are named
`lightning<ImportPath><Type>decode…` — a prefix derived from the package's import
path and the top-level type — so generating decoders for several types into one
package never produces colliding helper names. No annotation is needed; the
prefix is automatic.

## Recursive types

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
their decoders threaded with a recursion depth, and the decoders among them that
count a level — every struct decoder, and each slice, array or map decoder whose
elements are a named slice or map type on the cycle, so a cycle with no struct on
it (`type List []List`) is bounded too — refuse to descend past
`unstable.MaxDepth`, returning `ErrMaxDepth`:

```go
// A document whose nesting exceeds unstable.MaxDepth.
err := node.UnmarshalJSON(deeplyNested)
// err is unstable.ErrMaxDepth
```

Schemas without a cycle need no depth parameter. Cycles are detected from the
types themselves; no annotation is required.
