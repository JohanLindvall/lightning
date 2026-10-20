# Generator internals

[Documentation](../README.md) · [Contributing](../../CONTRIBUTING.md)

Read alongside the templates and comments in `main.go`.

## Codegen patterns and bounds checks

- **Cursor tests are unsigned** — `uint(i) < uint(len(data))` — in generator
  templates and the runtime. The compiler can't prove a cursor returned by a reader
  non-negative, so a signed test leaves a check on every `data[i]`; the unsigned form
  proves both bounds at once. The `panicBounds` stub is cold, but its branch costs a
  dispatch slot on an issue-bound core, and in a small leaf one surviving check is a
  CALL that costs the whole frame (look for `morestack`). Prove an idiom with
  `check_bce` before relying on it:
  - one byte: `uint(i) < uint(len(x)) && x[i] == c` (survives a loop-carried
    cursor); two bytes need two probes (`uint(i)+1 < uint(n)` keeps both checks);
  - a word load in a loop: `if uint(i) > uint(len(data)) { return i }` plus
    `for i <= len(data)-8` (`i+8 <= len(data)` keeps a check per word, and
    `d := data[i:]; for len(d) >= 8` is check-free but pushes `SkipWSRun` over
    budget);
  - `uint(i+4) <= uint(len(raw))` and `for uint(i) < uint(len(b))` over `b[i]`, not
    `len(raw)-i >= 4` or `range b[i:]`.

  Keep a check that sits off a latency-bound chain (reslicing it away puts an op on
  the chain): `readUnicodeEscape`'s four, `EscapeStringInto`'s short-run
  `for i+8 <= n`, `ExpectNull`.
- **Pass offsets, not reslices.** `f(b[i:])` costs ~7 instructions and clobbers
  len/cap, so scanners take a start and return an absolute index
  (`indexCloseOrEscapeAt`/`IndexCloseOrEscapeAt`, `indexStructuralAt`). Load
  `data[i:i+8]`, never `data[i:]` (6 more instructions a word). The escape scanners
  take a resliced buffer on purpose: the per-run gate guarantees ≥ `minVectorRun`
  bytes per call.
- **One bound, then unchecked loads**: `load64`/`load32` are unchecked on
  amd64/arm64; the caller proves the window once (`scanFloat`:
  `uint(i)+48 <= uint(len(data))`).
- **`unsafeStr` reinterprets the slice header** (cost 3) because `unsafe.String`/
  `unsafe.Slice` add a negate, a compare and a panic branch per call. It has no
  empty-slice guard, so an empty nocopy string may point into the input.
- **AArch64 immediates.** `AND`/`ORR`/`EOR` fold only bitmask immediates (a rotated
  run of ones, replicated: `0x80…`, `0x7f…`, `0x30…`, `0xf0…`); any other wide
  constant costs `MOVZ` plus up to three `MOVK`s at every use, back edge included.
  `ADD`/`SUB`/`CMP` take 12-bit immediates. The compiler never treats a 32-bit splat
  as a bitmask immediate, so run 32-bit SWAR tests in 64-bit registers — and check
  the disassembly.
- **Return where you decide.** A flag or `(byte, bool)` result tested later costs a
  `CSET` and a branch even when inlined. Copy a short shared body rather than calling
  it (`ParseInt` ← `ParseUint`, `UnescapeString` ← `UnescapeStringScan`).
- **Keep data-dependent counts off the cursor's chain.** Object decoding serializes
  on the cursor (~60 cycles a member on amd64). A constant advance or a predicted
  branch lets the next loads issue early; a SWAR count feeding the next address adds
  ~12–14 cycles.
- **Inline watch list** (budget 80; re-check under both GOARCHes after editing):
  `DecodeValue` 78, `SkipWSRun` 74 amd64 / 78 arm64, `skipValueOrEnd` 77,
  `CountArrayScalars` and `SkipObject` 72.

## Inline generated reads

Generated decoders write the common case of their once-per-member reads inline and
call pkg/unstable only for the hard case (`g.skipWS`, `g.readKey` in `main.go`):

- **Whitespace**: `if uint(i) < uint(len(data)) && data[i] <= ' ' { i++; if
  uint(i) < uint(len(data)) && data[i] <= ' ' { i = unstable.SkipWSRun(data, i+1) } }`
  — the common 0–1 bytes cost one or two compares; only a run of ≥ 2 enters the SWAR
  loop.
- **Keys**: `ReadKey` (cost 198) never inlines, so the generator emits the no-escape
  path — `unstable.IndexCloseOrEscapeAt(data, ks)` (offset in, absolute index out)
  and `unstable.UnsafeStr(data[ks:ke])` — and calls `ReadKey` only for an escaped key
  or an error.
- **Unknown fields** (`skipUnknown`): `>= '['` → `SkipValue`, `'"'` → `SkipString`,
  anything else → `SkipNumber` — exactly `SkipValue`'s own partition, so the two agree
  on malformed input too (`TestUnknownFieldSkipMatchesSkipValue`). Testing `'['` first
  makes the container arm the fall-through. There is no `'{'` → `SkipObject` arm:
  `SkipObject` inlines (cost 72) into every struct decoder and the code growth cancels
  the saving.

Rules:

- **Only for once-per-loop reads.** Inlining a per-field read repeats the block per
  struct field; a wide decoder (string_unicode's 60 string fields) then outgrows the
  inliner's budget for large functions, `IndexCloseOrEscape` stops inlining, and
  instruction-cache pressure grows with each field.
- **It pays where the inline path skips work** (the no-escape alias, the whitespace
  fast path) **or bypasses a dispatch around short work** (`SkipValue`'s frame and
  comparison tree in front of a number, a short string or a one-block container —
  `skipUnknown`, the walkers' arms). It doesn't pay where it only removes a wrapper's
  frame around an inlined asm scan (`SkipString`): that flat time is the call's ABI0
  marshaling, which stays.
- **The block must inline directly into the big caller**: a helper wrapping it costs
  97 (whitespace) or 102 (close-quote scan) and stays a call.
- Hand-applied in `StripDefaults`' `handle` (its three `SkipString` sites and its
  whitespace skips), in `any.go`, in the key reads of the `get.go`/`set.go` walkers,
  and in the stream `Reader` (`objectEach`, `enter`). The whitespace block is
  not used in `get.go`/`set.go`; those walkers use the shared whitespace helpers.

## Generator rules

- **Directives.** A known directive where it can't act (a fixed-size array, a defined
  scalar, any other non-struct/slice/map type; bare `nocopy` on a struct root) warns,
  as does a directive not attached to a type declaration (a blank line detaches it).
  An unknown `//lightning:*` name is an error when attached to a type and a warning
  when detached. A directive on a type another root reaches gives it its own method
  too (marked in `emitted` after `entryTypes`); its copy inside the reaching root
  follows that root's directives.
- **Every memo key carries the variant** (`g.prefix + g.cmark()`, names through
  `g.decFn`/`g.csuf`) and every other body input: the nocopy/lax suffixes, the map key
  type, and the `root:` marker that keeps a named slice root's `GrowSliceEst` decoder
  apart from a field decoder of the same element type. Otherwise roots with different
  directives share a decoder — a plain root inheriting a destructive sibling's
  in-place unescape (`TestLaxDecoderIsolation`).
- **Names.** Reserve decoder names with `g.uniq`, never a bare `g.used[fn] = true`. A
  schema type named like a generated identifier — a decoder parameter or local, an
  import of the generated file, or a predeclared identifier the generated code uses on
  its own (`max`) — is a hard error (`reservedIdents`/`checkReservedNames`). **Adding
  a local to an emitted template means adding it to `reservedIdents`.** Predeclared
  names echoed only from the schema's own type text (`bool`, `uint16`, `any`) are
  excluded; the known gap is `type uint16 struct{…}` used as a field, which `isScalar`
  still treats as predeclared.
- **Imports** come from flags set where a qualified name is *emitted*
  (`noteQualifiers` in `typeStr`, `numberRead`, `sliceDecoder` →
  `g.needJSON`/`g.needTime`/`g.needUnsafe`), spelled with the schema's own qualifier —
  never from scanning the output, which contains JSON keys as string literals. A
  template that emits a new qualified name must set its flag. `isRaw`/`isNumber`/
  `isTime` match only through the schema's own qualifier.
- **Which types get `UnmarshalJSON`** (`entryTypes`): start from the types nothing
  references, mark what they reach, then promote a cycle nothing emitted enters — its
  whole strongly connected component, chosen among components that are sources of the
  still-uncovered subgraph (first in source order) — and repeat to a fixpoint. Both
  choices keep the result independent of declaration order
  (`TestEntryTypesOrderIndependent`). Generic and alias declarations warn and get no
  method; an alias to a struct literal still registers in `g.structTypes` (not
  `g.order`) so it decodes as a field type.
- **Defined types as roots.** A type defined over a struct, slice or map
  (`type ruleRaw Rule`; `underlying` resolves chains, in-file or in a sibling) becomes
  a root with that shape only when it carries a directive or `//lightning:root` —
  opt-in, because the same spelling is the methodless twin. That is the idiom for a
  hand-written `UnmarshalJSON` that decodes its own fields without recursing; declare
  the defined type at top level (one inside a function body is invisible to the
  generator).
- **Hand-written unmarshalers are delegated to.** `collectUnmarshalers` finds them in
  the input and its siblings (never in a `_unmarshal.go`); `field` checks for them
  before the struct lookup; `delegate` finds the value's span with `SkipValue` and
  passes it — null included, aliasing the input, as the stdlib does — reporting the
  value's start on failure. Such a type is never generated for (as a root it is
  dropped with a warning). Every foreign `pkg.Type` other than `time.Time`,
  `json.RawMessage` and `json.Number` is delegated the same way; a missing method is a
  compile error, which is the intended check. `nullAssigns` lets a delegated type
  decide what null means.
- **Siblings** (`registerSiblings`): same-package `.go` files, excluding `_test.go`,
  `_unmarshal.go` and `_`/`.`-prefixed files. Their struct, slice, map and scalar
  types register by name in `g.sibling`, never in `g.order`; the reaching root emits
  their decoders under its own directives (directives on sibling types aren't
  reported). A sibling whose import qualifiers disagree with the input's
  (`qualifiersAgree`) is skipped whole, with a warning.
- **Defined scalar types** (`type Severity string`) register in `g.scalarTypes`
  (`declaresScalar` resolves forward references; `scalarKind` follows the chain at
  use) and are read with the kind's own reader and stored through a conversion
  (`scalarAs`), so null and nocopy behave exactly as for the plain kind. No method;
  not counted by `isFlatScalarStringStruct`, which is only a presize hint.
- **A named slice or map used as a field** decodes through its element type's decoder
  with the destination converted (`callDecoderOn`, e.g. `(*[]Item)(&v.Items)`).
- **Map keys** (`mapKeyAssign`): a string, an integer kind, or a type defined over
  one. Integer keys parse with `unstable.ParseInt`/`ParseUint` from a non-escaping
  `[]byte(key)`; a non-numeric name is `ErrBadNumber`.
- **`//lightning:strict`** is part of `cmark`/`csuf` (`Strict` suffix), so a type
  reached from a strict and a loose root gets two decoders. `unknownKey()` emits
  either `skipUnknown` or `return i, &unstable.UnknownKeyError{Key: string([]byte(key))}`
  — the key copied because the error outlives the input; the position reported is the
  value's. Maps are unaffected.
- **Tag options**: `,number` only on an `any` field (`anyValueNumber`; elsewhere it
  warns — threading it into slices or maps of `any` would touch every element
  decoder); `,omitempty`/`,omitzero` pass silently; `,string` warns as unimplemented.
  `,reader=func` calls a custom `(data []byte, i int) (T, int, error)` reader for
  the whole field, null included, assigning only on success; it composes with
  `lax`/`unwrap` but conflicts with `nocopy`/`number`. Reader names are parsed,
  validated against generated locals, and imports claimed at emission. An
  embedded field with a reader is decoded as a named field, not promoted.
- **Types and embedding.** Only the empty interface and its spellings
  (`interface{ any }`) decode as `any` (`isAnyInterface`; `ast.InterfaceType.Methods`
  also holds embedded elements and type-set terms). An embedded `time.Time`,
  `json.RawMessage` or `json.Number` decodes as a field keyed by its type name, while
  Go promotes an embedded unmarshaling method to the outer struct, allowing
  encoding/json to hand that type the whole document. The available methods depend
  on the Go toolchain; the embedding tests compare against its actual behavior.
  Tag names identified by `invalidTagRune` are honored here with a warning. Preserve
  that contract: rejecting the tag or adopting a stdlib fallback changes which key
  an existing decoder answers to. A lax `[N]scalar` field reaches the batched
  array readers through a thin wrapper (`TestLaxFixedArrays`).
- **Field dispatch** (`keyDispatch`/`chunkedKeyEq`): the compiler inlines a
  string-vs-constant compare only up to 16 bytes and calls `runtime.memequal`
  (spilling registers) beyond. A struct with a key over 16 bytes dispatches on
  `switch len(key)` and compares long names in ≤ 16-byte chunks
  (`key[0:16] == "EdgeTimeToFirstB" && key[16:] == "yteMs"`); buckets whose names all
  fit keep a nested `switch key`, and structs without long keys use a plain
  `switch key`. Unmatched keys `goto lightningSkipKey` — one shared skip in its own
  block, so the `goto lightningKeyDone` over it crosses no declaration. A field whose
  pipe-separated alternates differ in length is emitted once per bucket
  (`TestLongKeyDispatch`).
- **Trailing commas are rejected.** The generated object/slice/array/map loops
  (`genStructBody`, `sliceDecoder`, `arrayDecoder`, `mapDecoder`), the batch loops and
  `decodeAnyObject` use a first-iteration flag — `for first := true; ; first = false`,
  checked in the loop-top closer case, where a closer on a non-first iteration means a
  trailing comma; `decodeAnyArray` returns `[]` before its loop and checks for `]`
  right after each comma. Preserve this loop shape unless measurements establish
  that a change works across alignments; wide decoders are layout-sensitive.
  Tests: `TestTrailingCommaRejected`,
  `TestBatchTrailingComma`, the `objErrs` arms in `any_test.go`.
- **`,lax`** (`laxField`) skips a value that failed to decode with
  `unstable.SkipValueStrict`, not `SkipValue`, so a balanced but invalid value
  (`[1,]`, `[1 2]`) is still an error and the outcome doesn't depend on the host's skip
  path; a `]`/`}` where a value must start is `ErrInvalidJSON`.
- **Nulls.** Every `*OrNull` reader returns the zero value and a nil error on null, and
  generated code assigns unconditionally, so an explicit null zeroes a leaf field
  (string, bool, number, `json.Number`, `time.Time`) where encoding/json leaves it
  alone — a documented divergence pinned by `TestNullFieldsDivergeFromStdlib` (via
  `zeroNullLeaves`). Unconditional assignment keeps the common decode path free
  of a per-leaf null probe. `laxField` tracks `nullGuard` through `nullAssigns`:
  `json:"n"` and `json:"n,lax"` must agree on null handling. Any change to that
  behavior must update both paths and the compatibility guide together. Composites: slice/map/pointer/`any` become nil,
  `json.RawMessage` takes the literal `null`, nested structs and `[N]T` are untouched.
  A null *document* nils a slice or map root (`nullReset`) and leaves a struct root
  alone.
- **`,unwrap`** closures guard an all-whitespace body and reject trailing content with
  `genUnmarshal`'s own `unstable.SkipWS(data, i) != len(data)` check
  (`ErrInvalidJSON`): the body is a whole document. Under `,lax` a wrong-type body is
  still tolerated, trailing content is not. Tests: `TestUnwrapWhitespaceBody`,
  `TestUnwrapRejectsTrailingContent`.
- **Generated code has no comments** except the
  `// Code generated by the lightning generator from …; DO NOT EDIT.` header;
  explanations belong in `main.go` next to the template. Templates are `fmt.Sprintf`
  strings, so a `%` in emitted text becomes `%!(MISSING)` (`go vet` catches it). Check:
  `rg -c '//' <generated file>` is 1.
