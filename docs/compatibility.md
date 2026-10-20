# Compatibility and input contracts

[Home](../README.md) · [Documentation](README.md)

Read this before replacing `encoding/json`, reusing decoded values, or processing untrusted input.

## Differences from `encoding/json`

Generated methods implement `json.Unmarshaler`, so existing call sites can use
them. Some decoding semantics differ from `encoding/json`; review the contracts
below before migrating, especially when reusing targets.

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
  [reuse](generator.md#reusing-a-decode-target) one across documents *and* the later document
  sends an explicit null. If you rely on that, keep the value yourself or compare
  against a sentinel; a field tagged `lax` follows the same rule, so the two
  spellings never disagree.
- **Key matching is exact and case-sensitive.** `encoding/json` tries an exact
  match and then falls back to a case-insensitive one, so `{"ITEMS":[…]}` fills a
  field tagged `json:"items"`. lightning has no fallback: a key that isn't
  byte-equal to the tag (or to one of its
  [alternate names](options.md#alternate-field-names)) is simply an unknown key, and unknown
  keys are skipped. That is faster — the field switch is a `switch len(key)` over
  constant strings, with no case folding — and less surprising to read, but it
  fails quietly: a mis-cased document decodes to zero values — or, into a
  [reused target](generator.md#reusing-a-decode-target), to the *previous* document's values.
  If a source has unstable capitalization, list the spellings explicitly:
  `json:"items|Items|ITEMS"`.
- **An empty JSON array can leave a nil slice nil.** By default, `[]` resets a
  slice's length and keeps its backing array, so a fresh nil slice stays nil.
  `encoding/json` produces a non-nil empty slice. Use
  [`//lightning:emptyslice`](options.md#lightningemptyslice) when the distinction matters;
  it produces non-nil empty slices while preserving existing backing arrays.
- **The `,string` tag option is not implemented.** `encoding/json` uses it to read
  a number or bool that the producer wrapped in a JSON string — `{"n":"7"}` into an
  `int` field tagged `json:"n,string"`. lightning ignores the option and decodes
  the value with the field's declared Go type, so a quoted number fails with
  `ErrBadNumber`; the generator prints a warning when it sees the option, rather
  than emitting a decoder that silently does the wrong thing. Declare the field as
  a `string` and convert, or — when the string holds a whole JSON *document* rather
  than one scalar — use [`unwrap`](options.md#the-unwrap-tag-option). `omitempty` and `omitzero` are accepted
  and ignored, since lightning only decodes.
- **Slice, array and map elements are reset before being decoded**, where
  `encoding/json` decodes a slice or array element into whatever it already holds.
  Only observable when you reuse a decode target; see
  [Reusing a decode target](generator.md#reusing-a-decode-target) for the full rule and an
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
  under [Checking validity](json.md#checking-validity) spells this out.
- **Tag names are honored as written.** For example, `json:"a\"b"` matches
  the key `a"b` in Lightning. The standard library may fall back to the Go
  field name or ignore a field with an invalid tag name, depending on the
  toolchain. The generator warns about these names; the conformance tests
  compare against the running standard library. See
  [How it works](generator.md#how-it-works).
- **An embedded type with its own `UnmarshalJSON` does not take over the struct.**
  Go promotes the embedded *method* along with the fields, so `struct { time.Time;
  … }` itself satisfies `json.Unmarshaler` and `encoding/json` hands the whole
  document to `time.Time.UnmarshalJSON`, never reaching the sibling fields.
  lightning promotes fields only: the embed decodes as a named field keyed by its
  type name and the siblings decode normally. The stdlib fails loudly for
  `time.Time` and *silently* for an embedded `json.RawMessage`, which captures the
  entire document and leaves the siblings zero. See
  [Supported types](generator.md#supported-types).
- **Errors match package sentinels.** Use `errors.Is` with the
  package [sentinels](json.md#errors); there is no
  `*json.UnmarshalTypeError` or `*json.SyntaxError` carrying the offending field,
  Go type and byte offset. Strict decoding returns an `*UnknownKeyError` with the
  key name, which matches `ErrUnknownKey`. Errors from custom readers and delegated
  unmarshalers are returned as supplied. Tolerating a type mismatch on a field is
  what the [`lax`](options.md#the-lax-tag-option) option is for.
- **Invalid UTF-8 is passed through when decoding** rather than replaced with
  U+FFFD — see [Limits and untrusted input](#limits-and-untrusted-input). Pinned
  across every string path (copy, nocopy, destructive, object keys, the dynamic
  `any` decoder, `UnescapeString`) by
  `unstable.TestStringsPassInvalidUTF8Through`. (The *escape* direction —
  `EscapeString`/`EscapeStringInto`, which produce JSON — substitutes U+FFFD
  exactly as `encoding/json` does when marshaling.)
- **Escaped timestamps are unescaped before parsing.** lightning parses the JSON
  string's value as RFC 3339. Standard-library acceptance can depend on the Go
  toolchain; `TestReadTimeAcceptsEscapedTimestamps` compares both implementations
  at run time and checks the resulting instant.
- **Duplicate keys: last wins when decoding, first wins in the toolkit.** A
  generated `UnmarshalJSON` decodes every occurrence in document order, so the last
  one survives — matching `encoding/json`. The [`pkg/json`](../pkg/json) lookup and
  edit helpers (`Get`, `Lookup`, `GetMany`, `GetPaths`, `Set`, `SetMany`,
  `SetPaths`) take the **first** occurrence and leave any later one untouched,
  because they return or rewrite a single span rather than reading the object to
  its end. (`ObjectEach`/`ArrayEach` are the exception: iteration visits every
  member, duplicates included.) Duplicate keys are rare and neither answer is
  wrong, but the two halves of the library disagree, so don't build on it.

Two things that are *not* differences, though they are commonly assumed to be:

- **Trailing commas are rejected**, exactly as `encoding/json` rejects them:
  `{"a":1,}` and `[1,]` both fail with `ErrInvalidJSON`, in generated decoders,
  `DecodeAny` and [`Valid`](json.md#checking-validity) alike — and in the value of a
  [`lax`](options.md#the-lax-tag-option) field, whose skip is that same grammar walk. (The
  lenient bracket balancer that *skips* unwanted values accepts them,
  which is precisely why neither `Valid` nor `lax` uses it. It remains what an
  *unknown* field's value is skipped with, where nothing downstream depends on
  those bytes.)
- **Unknown object keys are skipped**, as with `encoding/json`'s default. The
  `DisallowUnknownFields` posture is a directive on the root —
  [`//lightning:strict`](options.md#lightningstrict) — rather than a decoder option,
  since a generated `UnmarshalJSON` takes no options.

## Limits and untrusted input

Worth knowing before pointing this at input you don't control:

- **Nesting is bounded at 10 000 levels** (`unstable.MaxDepth`, matching
  `encoding/json`); past that the decode returns `ErrMaxDepth` instead of
  descending. `DecodeAny`, `Valid`, `StripDefaults` and generated decoders for
  [recursive schemas](generator.md#recursive-types) all enforce it. `Get`/`Lookup`/`GetMany`/`GetPaths`/
  `Set` walk the document without recursing per level, so the bound does not limit
  what they can read; the value skipping they rely on is bounded as well, so no
  entry point here can be driven into a stack overflow.

  The bound exists because a Go stack overflow is a **fatal** error that `recover`
  cannot catch — so without it, deeply nested input would take the process down
  instead of returning an error.
- **A slice presize is a hint, and it is bounded by the document's own bytes.**
  Before filling a slice the decoder cheaply estimates how many elements the JSON
  array holds — by counting commas or braces, or by sampling the first few
  elements — and sizes the backing array once instead of growing it repeatedly.
  Those scans deliberately don't parse, so bytes inside a *string value* can
  inflate the estimate: `[{"s":"{{{{…"}]` is one element that counts as thousands.
  Each counter is therefore clamped to the most elements its byte span could
  legally hold, which caps the effect at what an equally large honest document
  would already cost. What remains is the ordinary JSON-to-Go expansion ratio (a `[{},{},…]`
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
  field's value and the toolkit's walkers step over
  — takes a number token as a run of `[0-9.eE+-]`, so `+`, `-` and `1.2.3` are
  number *spans* to it. That is the same bracket-balancer trade this section
  already describes for containers, and the reason the `…Checked` wrappers exist:
  the typed readers and `Valid` do reject those.
- **Numbers are validated by arithmetic, not by grammar** — see the table under
  [Checking validity](json.md#checking-validity) for what that accepts. A number that
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
