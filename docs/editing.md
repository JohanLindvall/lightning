# Editing JSON

[Home](../README.md) · [Documentation](README.md)

Edit raw JSON with `github.com/JohanLindvall/lightning/pkg/json`. Use the checked forms when inputs or keys are untrusted.

## Stripping default fields

The [`pkg/json`](../pkg/json) package can also prune a JSON document in a single
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
  `SkipWS` scans (faster, as [`GetCompact`](json.md#key-lookups) does); misreads spaced input.
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

The [`pkg/json`](../pkg/json) package can also splice values into a document in a
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
  `SetMany` is to `Set` what [`GetMany`](json.md#key-lookups) is to `Get` — one walk over
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
  rather than re-reading and re-writing the whole document for each path. A `nil`/empty path
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
error. They validate their arguments with [`Valid`](json.md#checking-validity) before the
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
