# Contributing to Lightning

Thanks for helping improve Lightning. Documentation fixes, small reproductions,
tests, and implementation changes are all welcome. You do not need SIMD or
assembly experience to contribute.

## Find a starting point

- Try the [quick start](README.md#quick-start) and improve anything unclear.
- Turn a decoding issue into a small regression test.
- Improve an example or document an input shape you use.
- For a new API or a broad optimization, open an issue describing the use case
  before investing in a large change. Small fixes can go straight to a pull request.

Use the [repository map](docs/development/layout.md) to find the relevant code.
The [documentation index](docs/README.md) links both user guides and internals.

## Set up a checkout

Install Go 1.25 or newer, Git, and Make. Fork the repository on GitHub if you plan
to submit a change, then clone your fork. To explore the upstream repository:

```sh
git clone https://github.com/JohanLindvall/lightning.git
cd lightning
go mod download
make generate
go test ./...
```

`make generate` builds the ignored conformance decoder; a fresh checkout needs
this before running tests or static analysis. No benchmark downloads are needed
for the main-module tests.

Create a branch for your change:

```sh
git switch -c describe-your-change
```

## Make and check a change

Keep changes focused and explain the behavior they affect. Read
[AGENTS.md](AGENTS.md) for the task-specific guides, and keep user-facing behavior
in the [user documentation](docs/README.md).

For code changes, run:

```sh
make generate
make fmt-check
make check
make vet
GOARCH=amd64 go vet ./...
GOARCH=arm64 go vet ./...
```

`make check` runs lint for both architectures and the tests. It installs
`golangci-lint` if needed. `make fmt-check` only checks formatting; use `gofmt -w`
on the Go files you edited to fix it. `make fix` also runs `go mod tidy`.

For documentation-only changes, check relative links and examples, then run
`make generate`, `go test ./...`, and `make fmt-check`. Keep performance tables
under `bench/` unchanged; CI owns those snapshots.

Additional checks depend on the change:

| Change | What to check |
|---|---|
| Generator | Add generate-then-compile cases, compare both generators across schemas, and regenerate the committed example when output changes. See [testing rules](docs/development/testing.md). |
| Runtime or toolkit behavior | Test errors, offsets, reused targets, and documented differences. See [testing rules](docs/development/testing.md). |
| Assembly or CPU dispatch | Run all affected bodies on supported CPUs or emulators, plus the toolchain, race, lint, and encoding checks in [the workflow](docs/development/workflow.md). |
| Benchmark module | Run `make bench-test`; `bench/` is a separate module and may need a newer Go toolchain. |
| Performance | Follow the [measurement guide](docs/development/benchmarking.md) and include reproducible results. |

Compare with `encoding/json` using a methodless twin. Calling it on a type with a
generated `UnmarshalJSON` invokes Lightning and does not provide an independent
baseline.

## Report an issue

[Open an issue](https://github.com/JohanLindvall/lightning/issues/new/choose) and
include the smallest example that shows the problem:

- Go types, JSON input, and the generator command or toolkit call.
- Expected behavior and actual output or error.
- Lightning version, `go version`, operating system, and architecture.
- For performance reports, the CPU model, benchmark command, allocation results,
  and baseline. A small input that reproduces the shape is more useful than a
  timing number alone.

Check the [compatibility guide](docs/compatibility.md) when comparing against the
standard library. Remove credentials and private data from reproductions.

## Submit a pull request

Describe the problem, the resulting behavior, and how you checked it. Link the
related issue if there is one. Include test commands and outcomes; explain any
checks you could not run so the reviewer knows what remains.

For a performance change, include the workload, CPU, Go version, and paired
measurements. Do not use a single timing run or a comparison between CI snapshots
as evidence of an improvement.

Only `example/event_unmarshal.go` is committed generated decoder output. Update
it with `cd example && go generate` when appropriate; keep other generated files
and local benchmark results out of the pull request.

Pull requests run CI on amd64 and arm64. Keep review conversations respectful,
specific, and focused on the change.
