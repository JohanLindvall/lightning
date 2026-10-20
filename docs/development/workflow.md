# Development workflow

[Documentation](../README.md) · [Contributing](../../CONTRIBUTING.md)

Run commands from the repository root unless a command says otherwise.

## Commands

- `make check` — `lint` + `test`. `make test` regenerates the conformance decoder,
  then runs `go test -cover ./...`.
- `make lint` — golangci-lint under `GOARCH=amd64` and `arm64` (`unused` sees only
  build-tag-selected files, so one arch misreports the other's code as dead).
- `make vet` — `go vet ./...`. asmdecl checks only the host GOARCH's assembly and
  isn't in `go test`'s vet subset, so also run `GOARCH=arm64 go vet ./...` from an
  amd64 host (and vice versa).
- `make fmt-check` (gate) / `make fix` (`gofmt -w .` + `go mod tidy`). Generated
  decoders must be gofmt-clean too.
- `make sveasm` / `make sveasm-check` — regenerate / verify the `WORD`-encoded
  arm64 instructions in `SVEASM_FILES` from their mnemonics. Needs an aarch64 GNU
  `as`: `$SVEAS`, `aarch64-linux-gnu-as` (binutils-aarch64-linux-gnu, so it runs on
  amd64), or native `as`. CI runs the check.
- `make bench-test` — `bench/get`'s tests, which `go test ./...` can't reach.
- `make bench-md` — what CI's manual Benchmark workflow (`.github/workflows/bench.yml`)
  runs to produce the committed tables: `pkg_bench.sh` → `bench/pkg_results_<arch>.md`,
  `bench/run_bench.sh` → `bench/results_<arch>.md`. `pkg_bench.sh` takes a
  benchmark-name filter as `$1` and honors `BENCHTIME`/`BENCHCOUNT`.
- Inline costs: `go build -gcflags=-m=2` — they differ per arch, so check both
  GOARCHes. Surviving bounds checks: `go build -gcflags=-d=ssa/check_bce ./pkg/...`.
- Validation over the corpus: from `bench/`,
  `go test ./valid -run='^$' -bench=BenchmarkValidCorpus`.
- **Before pushing assembly**: `GOTOOLCHAIN=go1.25.0 go build -a ./... &&
  GOTOOLCHAIN=go1.25.0 go vet ./...` (the `go.mod` floor; `-a` and `./...` both
  matter, or the build cache hides a failure), `go test -race ./...`, `make lint`,
  `make sveasm-check`, and every dispatch arm the change touches (see [Testing rules](testing.md#testing-rules)).
- **End of session**: `go test ./...` and gofmt clean. Leave the committed benchmark
  tables to CI — a locally generated table silently changes the host the file
  describes.

## Conventions

- **The edit/transform API is two-tier.** `Set`/`SetMany`/`SetPaths`/`StripDefaults`
  return only a `[]byte` and are best effort — bracket balancers, not parsers, passing
  uninterpretable input through rather than failing. That keeps them zero-alloc; don't
  add error returns or validity checks to the hot walkers. Untrusted input goes through
  the `…Checked` wrappers. New edit operations take the same shape.
- **The `go` directive (1.25) is the assembler floor**; CI assembles with that
  toolchain under `GOTOOLCHAIN=local`. Use a `WORD` encoding for any mnemonic the
  minimum supported assembler does not accept, including SVE instructions.
- **`WORD`-encoded instructions are derived from their comments.** In `SVEASM_FILES`
  the trailing-comment mnemonic is the source of truth: write
  `WORD $0x00000000 // <mnemonic>` and run `make sveasm` — never hand-encode. A second
  `//` starts prose, which is preserved. Inside a `#define`, write `/* <mnemonic> */`
  before the continuation backslash (that's how `SHORTCONV`/`LONGCONV` share code).
  sveasm assembles under `.arch armv8.6-a+sve2`; an instruction outside it
  (`+sve2-bitperm`'s `BEXT`) needs that widened plus a runtime gate `x/sys/cpu`
  doesn't provide.
- **No legacy-SSE instruction while the upper halves of Y0–Y15/Z0–Z15 are dirty**
  (after writing one and before the next `VZEROUPPER`), and no `RET` or tail `JMP` in
  that state. Intel can take a microcode assist per occurrence; AMD measurements
  alone cannot detect this penalty. The usual culprit is a GP→XMM `MOVQ`/`MOVL`,
  which Go encodes as legacy SSE: use `VMOVQ` or broadcast from memory.
  `TestNoSSEAfterAVX` (`avxmix_test.go`) checks every `*_amd64.s` in program order with
  macros expanded (registers 16–31 exempt); on Intel,
  `perf record -e cpu_core/assists.sse_avx_mix/u` over the pkg/unstable test binary
  must record nothing.
- **The committed benchmark tables come from CI** and are never regenerated locally
  (see [Measuring performance](benchmarking.md)).
- Bench `data.go` files declare a single top-level `Benchmark` with **anonymous**
  nested structs, so only `Benchmark` gets a generated method and
  `type benchmarkStd Benchmark` is a reflection-only baseline.
- Keep one authoritative copy of anything enumerated (a divergence list, a dispatch
  table) and point to it; copies across files rot.
