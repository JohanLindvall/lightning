# Testing rules

[Documentation](../README.md) · [Contributing](../../CONTRIBUTING.md)

Read these rules before changing behavior or adding tests.

## Testing rules

- **Stdlib differentials need a methodless twin.** A generated `UnmarshalJSON` makes
  `encoding/json.Unmarshal(doc, &v)` call lightning. Compare against
  `type FooStd Foo` (as `bench/` does with `benchmarkStd`), and never put a
  `//lightning:` directive on the twin — that makes it a root with a method.
  `TestStdlibTwinsAreReflectionOnly` rejects any Unmarshaler in a twin's type graph
  except `json.RawMessage` and `time.Time`; a twin holding another stdlib type with
  its own `UnmarshalJSON` needs the same exception. A twin still inherits methods
  promoted from *embedded* fields, which is the behavior the embedding tests measure.
- **Pin deliberate divergences** with tests that fail whether the divergence widens
  or silently closes (`TestValidDivergesFromStdlib`,
  `TestNullFieldsDivergeFromStdlib`, `TestEmbeddedUnmarshalerDivergesFromStdlib`,
  `TestSkipPathsDivergeOnMalformed`, `TestReadTimeAcceptsEscapedTimestamps`,
  `TestStringsPassInvalidUTF8Through`), and keep the [compatibility guide](../compatibility.md) in sync. Standard-library behavior
  can depend on the Go toolchain. The stdlib-side pins
  (`TestReadTimeAcceptsEscapedTimestamps`, `TestGenerate/invalid_json_tag_names`,
  `TestEmbeddedUnmarshalerDivergesFromStdlib`) compute its answer at run time:
  when behavior differs, pin the difference; when it agrees, compare the values.
- **A differential proves nothing about shapes it never generates.** Check
  non-vacuity (assert the corpus reaches the shape) and ask what an oracle never
  sees: duplicate keys, reused targets, whitespace inside kept members. A kernel that
  stops early passes any differential, so assert it consumed the whole input
  (`TestIntRunWindows`, `walkFloatRun`).
- **Oracles share nothing with the code under test** — no tables, no scanners
  (`strictStringEscapedReference`) — and the error sentinel and the end offset are
  both part of the contract.
- **Run every dispatch arm a change touches.**
  - Live dispatch runs only the host's widest body. Flip the `use…` flags to test the
    narrower arms against the scalar oracle (`TestIndexVariantsFlip`,
    `TestIndexArm64Lengths`, `TestSkipBlocksVariants`), but only to bodies the CPU
    has: iterate `floatRunBodies()`/`validRunBodies()` with `defer restoreKernels()`.
    Forcing a flag regardless of the CPU dies with SIGILL. A reference run must clear
    every flag the code reads (`floatRunOff()`/`validRunOff()`), or the kernel is
    compared with itself.
  - All arms run from one x86 box: Intel SDE `sde64 -skx` (AVX-512BW without VBMI)
    and `-icx`/`-spr` (VBMI); `qemu-x86_64 -cpu Haswell-v4 | Nehalem | qemu64`
    (AVX2, SSE4.2, SSE2 — TCG has no AVX-512); `qemu-aarch64 -cpu cortex-a72 |
    neoverse-n1 | neoverse-n2 | max,sve-default-vector-length=64` (NEON without and
    with DotProd, SVE2, 512-bit SVE2); `qemu-riscv64`/`qemu-s390x` for the pure-Go
    fallbacks, little- and big-endian. Build each package's test binary once
    (`go test -c`, with `GOARCH` for foreign arches) and run it from its package
    directory; the matrix takes ~30 minutes.
  - Assembly counts as verified only once it has run on its arch, natively or under
    qemu: asmdecl checks frame offsets only. qemu and SDE verify correctness, never
    speed.
- **Guard pages** (`TestAssemblyStaysInBounds`, `TestNumberKernelsStayInBounds`,
  Linux): every assembly body runs over buffers flush against a `PROT_NONE` page at
  either end, at every length to 300 — the only proof of masked and overlapping-tail
  bounds. Patterns must put the farthest load in a window's last lanes (a number at
  lane 58+), or a shrunken bound faults nowhere. Add cases for every new body or
  window bound, and sabotage-check by cutting a bound.
- **Fixed-window fast paths fail silently** (a missed backslash returns escapes
  verbatim): test every position at every length around each boundary, with the
  buffer's capacity ending at the input (`TestStringFindsEveryEscape`,
  `TestUnescapeFindsEveryEscape`, `TestUnescapeStringCopyFindsEveryEscape`).
- **Resumable scanners need every chunk split**, sizes around 64 and boundaries
  inside strings — a dropped carry still passes at sizes 1, 7 and whole
  (`TestValueScannerChunkingIsInvisible`).
- **Allocator tests retain each buffer and re-read it** — never compare addresses,
  freed addresses recur — and run under `-race` (`TestEscapeScratch*`). Zero-alloc
  assertions must hold under `-race` too: race instrumentation changes inlining, and
  a nil-start `append` stays on the stack only while the compiler can see it bounded,
  so make the promise structural with a fixed per-frame `[N]T` backing.
- **Order-dependent logic needs permutation tests** (`TestEntryTypesOrderIndependent`).
- **Generator changes.** Generation must produce compilable code.
  Add new shapes to `TestGenerate`, which generates and then compiles them. `genCase`: `extra` adds sibling files (directories allowed, so a
  second package works); `wantMethods` pins the exact set of `UnmarshalJSON`
  receivers; a `wantNoWarn` substring must not appear in the case name (diagnostics
  carry the temp path, which carries the test name). Every generator change owes a
  dual-generator diff (parent vs tree) over conformance and every bench schema, run
  inside each case's directory (the sibling scan reads the package's other files):
  byte-identical except where intended, with identical diagnostics. A change that
  does alter output also regenerates `example/event_unmarshal.go`.
- **Prove a source-comment-only change** by comparing the normalized disassembly of a probe
  binary built before and after; raw binaries differ anyway (pclntab, DWARF lines).
- **Sabotage-check new tests**: break the guarded code and watch them fail. Copy the
  file aside first — `git checkout <file>` also discards your own uncommitted work.
- **An in-place mode is its own contract**: test in-place == fresh, byte for byte.
- **Before calling a divergence or regression yours**, rerun the input on the parent
  commit. Agent worktrees branch from `origin/main`, not local `main`. Build after
  merging parallel branches: two branches adding the same identifier merge cleanly
  and then fail to compile.
