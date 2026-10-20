# JSON toolkit internals

[Documentation](../README.md) · [Contributing](../../CONTRIBUTING.md)

Implementation guidance for the public `pkg/json` API.

## pkg/json toolkit

- **The walkers dispatch `SkipValue`'s arms at the call site** (`objectEach`,
  `arrayEach`, `arrayEachIndex`, `getMany`, `objectField`, `walkPaths`, the `Reader`):
  `uint(c)-'-' <= 12` → `SkipNumber` (exactly `SkipValue`'s number bytes; widen before
  subtracting), `"` → `SkipString`, `{` → `SkipObject`, anything else → `SkipValue`.
  Array walkers test numbers first, object walkers strings first. On a short value the
  frame and switch cost as much as the scan. A missed byte is harmless; a wrongly
  claimed one is a bug (`TestArrayEachDispatchMatchesSkipValue`). The walker's own
  `uint(i) >= uint(len(data))` test proves `data[i]` in range. Callbacks are checked
  with `err == ErrStop` before `errors.Is`. `arrayEach`/`arrayEachIndex` are
  layout-sensitive — judge edits by instruction counts. `set.go`'s walkers
  deliberately keep `skipValueOrEnd`, which already inlines at every call site.
- **Scalar readers** (`KindOf`, `ParseInt`/`ParseUint`, `String`, `Bool`,
  `UnescapeString*`) are small leaves where one surviving bounds check costs a frame
  (see [bounds checks](codegen.md#codegen-patterns-and-bounds-checks)). `KindOf` dispatches through a static `[256]` table
  (`kindOfByte`): one load costs every kind the same, where a `switch`'s compare tree
  gets re-laid by unrelated edits. A literal must be the whole remainder, with trailing
  whitespace measured by `SkipWS`.
- **`GetPaths`** keeps one `[]int` scratch for every level's active-path set (sized
  `len(paths)*(maxDepth+1)`, backed by a stack `[32]int` for small sets — the output
  aliases only `data`, so it never escapes). There is no early exit once all paths are
  captured: first occurrence wins, so later duplicates must still be scanned. Each
  path behaves as if requested alone: `walkPaths` returns `(end, err, fatal)`; a
  failed descent is re-skipped with the lenient `SkipValue`, the walk continues, and
  the first error is still returned. Duplicate parent keys stop at the first, like
  `Get`. Tests: `TestGetPathsIndependenceRandomized`,
  `TestGetPathsIndependentOfCoRequestedPaths`.
- **`Set`/`SetMany`/`SetPaths` are zero-alloc with a reused `out`.** Creation streams
  into `out` (`appendMember` writes multi-key nesting directly; `setSpan` reports a
  non-object intermediate through `member`/`nested`). `SetMany`'s found flags are a
  stack `[64]bool`; `SetPaths`' per-key `sub` set and `setObject`'s `recurse`/`create`
  use per-frame `[8]int` backings (`subbuf`, `idxbuf`), because under `-race` the
  nil-start appends reach the heap. The key test runs first, once: a non-leaf match
  descends without pre-skipping and takes its end from the recursion, so each on-path
  container is scanned once. Early exits: `SetMany` copies the rest verbatim once every
  key is found; `setObject` does so only at the root frame
  (`depth == 0 && nmatched == len(active)` — a nested frame owes its parent the `}`
  offset). Shared semantics: the first occurrence of a duplicate key is edited, a key
  requested twice keeps the first request, and a non-object root has its value
  replaced with the surrounding bytes kept. Benches: `BenchmarkSet`,
  `BenchmarkSetMany`, `BenchmarkSetPaths`, `BenchmarkSetManyEarlyExit`,
  `BenchmarkSetPathsEarlyExit`; tests `TestSetManyMatchesSet`, `TestSetTruncatedNoPanic`.
- **`StripDefaults`** (`stripper.handle`, recursive, never inlined;
  `BenchmarkStripDefaults`, `BenchmarkStripDefaultsCompact`):
  - In-place mode (`output == input[:0]`) holds only while no write lands on bytes the
    walk will read again. A container member's key, colon and recursion output are
    written speculatively and rewound when the value strips to nothing, so the keep
    decision is made first (from bytes already read), and a kept member's original
    span is snapshotted into `stripper.scratch` — only when the output really aliases
    the input (`unstable.SameBuffer`, conservative), which keeps a separate output
    buffer at 0 allocs. One scratch serves every level (the argument is next to the
    buffer). The snapshot's end comes from `SkipValue`, which can disagree with the
    walk on malformed input, hence the guarded fallback.
  - A kept container member is re-emitted by `emitKeptCompact(src, base)` +
    `compactValue`, so `RemoveWhitespace` holds, reading from whichever buffer holds
    the original.
  - Keep `emitField` unchanged and inlinable (parameterizing it by `src`/`base` costs
    88 > 80); new arms go out of line, and `keepKey`'s length pre-filter leads the
    container branch.
  - A keep-key member whose container empties rewinds to just past its separating
    comma (`postComma`); a whitespace-only object in array position must not eject.
  - Tests: `TestStripDefaultsInPlaceFuzz` (in-place ≡ fresh),
    `TestStripDefaultsInPlaceMatchesFresh`, `TestStripDefaultsSnapshotShortOfMember`,
    `TestStripDefaultsInPlaceFuzzMalformed`, `TestStripDefaultsKeepKeyContainer`, and
    the `compact(preserve) == remove` oracle `TestStripDefaultsWhitespaceModes`.
- **Checked wrappers** (`checked.go`): `Valid` on the arguments and the result,
  `ErrValueCount` for a short `rawVal`, and `ErrUnsafeKey` — keys are written raw
  between quotes, so a key like `x":1,"role` would inject a well-formed member that
  `Valid` accepts. `StripDefaultsChecked` treats a token-free result as empty
  (`unstable.SkipWS(res, 0) == len(res)`; `PreserveWhitespace` can leave a
  whitespace-only remainder) and normalizes it to an empty slice.
- `pkg/json` re-exports all twelve `Err*` sentinels, `ErrUnknownKey` included
  (`TestSentinelsMatchable`); descending into a non-object returns `ErrExpectObject`.

## Streaming reader (`stream.go`, `unstable.ValueScanner`)

- **Each value is scanned once.** A value already buffered settles with `SkipValue`
  (parity with the in-memory walkers); otherwise `valueMore` refills once and feeds
  `ValueScanner`. Re-running `SkipValue` after each refill is O(n²)
  (`TestStreamHugeElementIsLinear`). After a successful scan, only a number ending
  exactly at the end of the buffered bytes is still undecided (`errMoreInput`).
- **`skip`** probes once, capped at `skipProbe` (4 KiB) so its cost tracks the value
  rather than the buffer; it accepts only `end < lim` and never refills to make a value
  fit, so a sibling bigger than the buffer streams past.
- **`ValueScanner` modes.** `Reset` behaves byte for byte like the scalar skip (typed
  brackets, `ErrMaxDepth`) and is the documented default (`TestValueScannerErrors`,
  `TestValueScannerMatchesScalarSkip`). The Reader uses `ResetFast` — the block scan,
  ~14× fewer instructions a byte — which inherits skipfast.go's malformed-input
  divergences; that is acceptable only because it steps past or delimits values
  exactly where the in-memory walkers use `SkipValue`. `skipBlocks` always starts
  outside strings, so `containerFast` walks out of a string the chunk begins in
  (`stringBody`), carries `inString`/`escaped` itself, and cuts the chunk at its last
  full block.
- **Fast paths are written out at the call site** (the skip arms; the buffered case
  of `space`/`colon`/`afterElement`), and nothing is consumed before the fallback.
  Keys use `IndexCloseOrEscapeAt` + `UnsafeStr` (`objectEach`, `enter`); `objectEach`
  decodes the key after the value, from `hold+klen`, because compaction shifts every
  index equally.
- **Several paths per document, forward only** (`enter`): each call resolves its path
  from where the previous call left the cursor, tracked by `open` (keys of nested
  objects entered and not yet closed), `entered` (the root's `{` passed), `after`
  (cursor just past a member's value), `pending` (the closer of a container a callback
  left with `ErrStop`; `finish` skips the rest first) and `done` (a keyless call
  consumed the root). A new path shares the open prefix, closes deeper objects with
  `closeObject`, and scans forward; a miss that reaches an object's `}` goes through
  `closed`, so the next path resolves from the parent; a key behind the cursor is
  `ErrKeyNotFound`, because the buffer is a bounded window. Tests:
  `TestStreamContinuesToTheNextPath`, `TestStreamContinuationIsForwardOnly`,
  `TestStreamContinuesPastAStoppedWalk`.
- A callback's window is valid only until the callback returns. `WithMaxElement`
  bounds one element (`ErrElementTooLarge`); `WithBufferSize` is the caller's buffer
  lever (`defaultReaderBuffer`, 64 KiB).

## Escaping (`escape.go`)

- `IndexEscape` finds the next byte that needs escaping: `indexEscapeSSE2` on amd64
  (built like the string scanner — SSE2 first 32 bytes, then AVX2, then a 16-byte loop
  and scalar tail; control bytes via `PMINUB(v, 0x1f) == v`; inputs under 16 bytes
  skip the splat loads), `indexEscapeArm64` on arm64 (SVE2 `CMPLO #32`, otherwise
  NEON `VUMIN`), SWAR `indexEscapeScalar` elsewhere. `SwarNeedsEscape` is the one
  spelling of the predicate.
- `EscapeStringInto` picks SWAR or vector once per run: with fewer than `minVectorRun`
  (48) bytes left it walks SWAR words, otherwise it probes one word and hands the clean
  bulk to `IndexEscape`. Deciding per run keeps `indexEscape` inlinable; a per-word
  budget, a SWAR prescan inside `indexEscape` and an asm scalar peek are all slower.
- `EscapeString` (Builder) escapes the tail into a stack `[128]byte` (non-escaping);
  longer tails grow on the heap.
- **Ill-formed UTF-8 becomes U+FFFD when escaping**, as encoding/json does when
  marshaling. The walk uses a predicate widened by non-ASCII bytes
  (`SwarNeedsEscapeOrNonASCII`, `IndexEscapeNonASCII`) until the first non-ASCII
  byte, where one `utf8.Valid` decides the rest: valid → `escapeValidInto` (plain
  predicate), invalid → `escapeInvalidInto` (a DecodeRune walk writing raw U+FFFD per
  ill-formed byte).
  - The widening costs no extra ops: amd64 ORs the raw chunk into the match vector
    before `PMOVMSKB` (sign bits are the non-ASCII lanes), NEON adds `VUSHR $7` +
    `VORR`, SVE2 tests `SUB 0x20`/`CMPHS #0x60`, scalar tails test
    `int8(c) < 0x20`, and the SWAR form `((v-0x20·lo)|v)&hi` guarantees only its
    lowest set bit (callers use only `TrailingZeros64`).
  - In `EscapeString`, a remainder first reached at an *escape* byte hasn't been
    UTF-8-checked and must go through `EscapeStringInto`, not `escapeValidInto`.
  - Tests: `escapeReference` (all bytes, UTF-8 corners straddling the 8/16/32/48-byte
    boundaries, fuzz, both entry points), the `indexEscapeNonASCII` arms of
    `TestIndexFunctionsMatchScalar`/`TestIndexVariantsFlip`,
    `TestIndexEscapeNonASCIIScalarOracle`, `TestEscapeStringMatchesStdlibCoercion`.
- Escaping has **no in-place form**: output grows, so `out` must not overlap the input
  (unchecked, per the two-tier convention). `UnescapeStringInto` allows
  `out == in[:0]` because unescaping shrinks — check the size direction before
  borrowing a sibling's buffer convention.
- Decoding passes invalid UTF-8 through ([compatibility guide](../compatibility.md); `TestStringsPassInvalidUTF8Through`
  covers all seven string paths).
