# Agent guidance

Lightning generates allocation-light `json.Unmarshaler` implementations and
provides a SIMD scanning runtime and a public JSON toolkit.

[README.md](README.md) is the user entry point. [CONTRIBUTING.md](CONTRIBUTING.md)
covers setup and the contribution workflow. Detailed guidance is split by task
below; read the relevant guides and the comments beside code before changing it.
All paths mentioned in the guides are relative to the repository root unless
stated otherwise.

## Read for your task

| Task | Required guidance |
|---|---|
| Find an implementation or test | [Repository layout](docs/development/layout.md) |
| Build, format, lint, or prepare a change | [Workflow and conventions](docs/development/workflow.md) |
| Change behavior or add tests | [Testing rules](docs/development/testing.md) |
| Measure or claim a performance change | [Benchmarking](docs/development/benchmarking.md) |
| Change templates, directives, root selection, or generated reads | [Generator internals](docs/development/codegen.md) |
| Change readers, numbers, strings, slices, whitespace, or validation | [Runtime internals](docs/development/runtime.md) |
| Change native kernels, scanning, or dispatch | [SIMD and assembly](docs/development/simd.md), plus workflow and testing |
| Change lookup, edits, streaming, or escaping | [Toolkit internals](docs/development/toolkit.md) |
| Change documented behavior | [User guides](docs/README.md) and [compatibility contracts](docs/compatibility.md) |

## Rules for every change

- Preserve the documented contracts. Keep the compatibility guide and regression
  tests aligned when behavior changes.
- Keep one authoritative copy of an enumerated contract or dispatch table; link
  to it from other guides. Keep guidance about the current implementation rather
  than experiment logs or historical results.
- Generated decoders must compile and be gofmt-clean. Generate conformance code
  with `make generate`; update `example/event_unmarshal.go` if generator output
  changes. Other `*_unmarshal.go` files are ignored.
- A generator change requires a parent-versus-tree generator diff over conformance
  and every benchmark schema, from each case's directory. Preserve diagnostics
  and unrelated output.
- Test every dispatch arm a change touches, with CPU feature gates respected.
  Assembly is verified only when it runs on its architecture, natively or under
  emulation. See the assembly checks in the workflow guide.
- Performance claims need instruction counts and interleaved comparisons. The
  committed benchmark tables are CI snapshots; do not update them from local runs.
- At the end of a session, run `go test ./...` and `make fmt-check`. On a fresh
  checkout, run `make generate` first.
