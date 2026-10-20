# Repository layout

[Documentation](../README.md) · [Contributing](../../CONTRIBUTING.md)

Use this map to find the implementation and its tests.

## Layout

- `main.go` — the generator (`package main`): reads struct definitions (plus
  same-package sibling files), emits `*_unmarshal.go`. `field`/`sliceDecoder`/
  `mapDecoder` build decoders; `slicePresize` decides presizing. Per-type
  `//lightning:` directives are parsed by `hasDirective` (whitespace-insensitive:
  `//lightning:compact` ≡ `// lightning: compact`) into `compactTypes`/
  `nocopyTypes`/`destructiveTypes`/`arenaTypes`/`strictTypes`, which the per-root
  loop turns into `g.compact`/`g.nocopy`/`g.destructive`/`g.arena`/`g.strict`:
  - `compact` — elide inter-token whitespace skips (`g.skipWS`).
  - `nocopy` — a slice/map root aliases its keys/elements into the input.
  - `destructive` — the type's nocopy strings unescape in place, into the input
    buffer (implies nocopy).
  - `arena` — small scalar slices carve their backings from per-decode chunks.
  - `strict` — an unknown key fails with `*unstable.UnknownKeyError` instead of
    being skipped.
  - `emptyslice` — an empty JSON array produces a non-nil slice, retaining reused
    backing arrays. `batchSliceDecoder` wraps batched readers only in this mode.
  - `root` — asks for a method and nothing else.

  `g.cmark`/`g.csuf` keep variants apart in memo keys and function names
  (`Compact`/`Destructive`/`Arena`/`Strict`/`EmptySlice` suffixes); nocopy variants differ by
  the `nocopy` decoder param / `NoCopy` suffix.
- `generator_test.go` — generator tests; `TestGenerate` is the generate-then-compile
  table.
- `conformance/` — end-to-end tests against a generated decoder (`data_unmarshal.go`;
  every `*_unmarshal.go` is gitignored, and `make generate`/`make test` rebuild it).
- `example/` — the runnable example and its live demo: the schema, the
  generated decoder (the one committed `*_unmarshal.go`, so the package builds on
  pkg.go.dev, where its `Example` functions run in the browser), and those
  examples. `TestExampleDecoderIsCurrent` fails when the generator's output
  changes; regenerate with `cd example && go generate`.
- `internal/sveasm` — derives the `WORD`-encoded arm64 instructions from their
  comment mnemonics (see [Conventions](workflow.md#conventions)).
- `pkg/unstable` — the runtime the generated decoders call, plus primitives
  exported for `pkg/json`; nearly all performance work happens here.
  - Readers: `read.go` (`Read*`), `parseint.go`, `digitrun.go`,
    `digits_{amd64,other}.go` (`readerWordFold`), `numbyte_{table,cmp}.go`
    (`isNumberByte`), `numeric.go` (`scanFloat`, `scanFloatSlow`, Eisel-Lemire;
    `powers_table.go` is generated), `string.go` (unescaping, `Unwrap`), `escbuf.go`
    (pooled escaped-string chunk), `date.go`/`time.go`, `any.go` (dynamic
    `DecodeValue`), `valid.go` (`SkipValueStrict`).
  - Arrays: `batch.go` (batched scalar-array readers), `points.go`
    (`DecodeFloat64Points`), `count.go` + `count_{amd64,arm64}.s`/`count_other.go`
    (presize `countKernel`), `grow.go`, `arena.go`.
  - Number kernels: `intrun_*` (integer arrays), `floatrun_*` (decimal arrays, the
    ring points walk, `Valid`'s number walks).
  - Scanning: `skip.go`, `skipfast*` (SIMD container skip), `structural.go`
    (`structuralMask`), `scan.go` (`ValueScanner`, the resumable skip the stream
    reader uses), `simd_*` (string/structural/escape scanners), `ws_{amd64,other}.go`,
    `load_le.go`/`load_other.go` (`load64`/`load32`, unchecked on amd64/arm64);
    `unstable.go` holds the rest.
  - Per-arch `use…` dispatch flags live in `*_{amd64,arm64}.go`; `*_other.go`,
    `simd_scalar.go` and `skipfast_noasm.go` are the portable fallbacks.
- `pkg/json` — the public API, implemented on the exported primitives: `get.go`
  (`Get`/`GetMany`/`GetPaths`/`Lookup`/`ObjectEach`/`ArrayEach`/`ArrayEachIndex`,
  plus `*Compact`; `GetPaths` is `Get` for several nested paths in one
  prefix-sharing pass), `stream.go` (`Reader`: the walkers over an `io.Reader`
  through a bounded buffer), `set.go` (`Set`/`SetMany`/`SetPaths`),
  `strip_defaults.go`, `checked.go` (`…Checked` wrappers), `valid.go`, `escape.go`,
  `scalar.go` (`String`/`Bool`), `kind.go` (`KindOf`; the constants are
  `Kind`-prefixed because `String`/`Bool` are functions), `parseint.go`, and
  `json.go` (`DecodeAny`/`DecodeAnyNumber`/`UnescapeString`/`UnescapeStringCopy`/
  `ParseFloat`, wrappers over pkg/unstable internals).
- `bench/` — a separate module (keeps the competitor libraries out of the main
  module's deps), one directory per case. `bench/run_bench.sh` regenerates the
  decoders and benchmarks lightning against encoding/json, easyjson, sonic and
  others; `pkg_bench.sh` (repo root) runs the main-module microbenchmarks.
  `bench/large-json/input.json` (~8 MB GeoJSON) is gitignored and downloaded from
  `input.url`.
