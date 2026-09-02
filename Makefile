.PHONY: all download check lint bench bench-md update-tools fix test generate vet fmt-check bench-test sveasm sveasm-check

GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin
PATH := $(GOBIN):$(PATH)

export PATH

all: fix check

download: go.mod go.sum
	go mod download

check: lint test

# Lint BOTH architectures, for the same reason `vet` does. The premise this
# target used to rest on — that static analysis is architecture-independent and
# so needs one run — is false wherever a package has per-arch files: `unused`
# only sees the files selected by the build tags, so a Go declaration reachable
# only from the *other* architecture's code reads as dead. The concrete case is
# pkg/unstable/simd_arm64.go, whose NEON routines are tail-called from the SVE2
# entry points in assembly; linting only amd64 never compiles that file at all.
# golangci-lint honours GOARCH, so both passes run on either host.
lint: $(GOBIN)/golangci-lint
	GOARCH=amd64 golangci-lint run ./...
	GOARCH=arm64 golangci-lint run ./...

test: conformance/data_unmarshal.go
	go test -cover ./...

# Build every gitignored generated source. A phony aggregate over the file targets
# below, so CI (and a fresh checkout) can produce them in one call before linting.
generate: conformance/data_unmarshal.go

# The conformance test decodes test.json with a generated decoder, but that
# decoder is gitignored (like every *_unmarshal.go), so a clean checkout has none
# and the package would not compile. Regenerate it from its source struct whenever
# that struct or the generator (main.go) changes; listing it as a prerequisite of
# test makes the suite run against an up-to-date decoder (and check, which depends
# on test, inherits it).
conformance/data_unmarshal.go: conformance/data.go main.go
	go run . conformance/data.go

# `go vet` over the whole main module. Run this on EVERY architecture that ships,
# not just one: asmdecl is inherently per-GOARCH. On an amd64 host it checks
# simd_amd64.s and skipfast_amd64.s against their Go declarations and never looks
# at the arm64 pair, and vice versa — so an amd64-only vet leaves the NEON
# assembly checked by nothing. `go test`'s built-in vet subset (atomic, bool,
# buildtags, directive, errorsas, ifaceassert, nilfunc, printf, stringintconv,
# tests) does not include asmdecl either, so running the tests is not a
# substitute. The class of bug this catches is real and recorded in CLAUDE.md:
# maskBlock's result offsets move because Go 8-aligns the result block after the
# `isArray bool` argument, and a hand-written +28(FP) where +32(FP) is required
# reads the wrong words at run time while assembling and testing cleanly on the
# other arch. Confirmed by experiment: that exact edit is reported by
# `GOARCH=arm64 go vet ./pkg/unstable/` and is invisible to `GOARCH=amd64 go vet`.
vet: conformance/data_unmarshal.go
	go vet ./...

# Fail if anything is not gofmt-clean. `fix` below is the fixer; this is the gate,
# so formatting drift is reported rather than landed. It walks the whole tree —
# gofmt follows directories, not modules, so the separate bench module is covered
# too — including the generated decoders, which the generator emits through
# go/format and which must therefore stay clean.
fmt-check:
	@unformatted=$$(gofmt -l .); \
	if [ -n "$$unformatted" ]; then \
		echo "not gofmt-clean (run 'make fix'):"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

# Regenerate the hand-encoded SVE instructions in the arm64 assembly from the
# mnemonics in their comments. The Go assembler has no SVE mnemonics (Go 1.26),
# so pkg/unstable's SVE2 scanners spell those instructions as WORD constants —
# and a bare 32-bit constant is unreviewable: one wrong nibble is a different
# instruction that still assembles, links and runs. So the mnemonic is the source
# of truth and the constant is derived from it by internal/sveasm, which
# assembles the comments with the GNU assembler.
#
# Run this after editing any SVE instruction; sveasm-check below is the gate that
# fails if the two ever drift apart. Needs an aarch64 assembler: binutils on an
# arm64 host, or binutils-aarch64-linux-gnu (the cross-assembler) on any host —
# the tool prefers the cross one so the check also runs on amd64.
SVEASM_FILES = pkg/unstable/simd_arm64.s pkg/unstable/intrun_arm64.s

sveasm:
	go run ./internal/sveasm -w $(SVEASM_FILES)

# The CI gate for the above. Kept out of `check` for the same reason `vet` and
# `fmt-check` are: it needs a tool a contributor may not have installed, so it is
# a separate target CI wires up explicitly (see .github/workflows/ci.yml), rather
# than a hidden prerequisite that breaks an unrelated local build.
sveasm-check:
	go run ./internal/sveasm $(SVEASM_FILES)

# Compile and run the bench module's own tests. bench/ is a separate module, so
# `make test`'s `go test ./...` cannot reach it, and both benchmark runners enter
# it with a -run pattern that deliberately matches no test (they are measuring,
# not testing). That left bench/get/get_bench_test.go — the jsonparser/gjson
# parity tests for Get and ObjectEach — compiled by nothing and executed by
# nothing, free to rot into a build failure unnoticed. This target is that guard.
# ./get is the only package listed because it is the only one whose test file is
# self-contained: its input.json and the sibling cloudflare/input.json are both
# committed, whereas every other case's test needs the gitignored harness that
# bench/run_bench.sh writes (see the test pass in that script, which covers those).
bench-test:
	cd bench && go test -count=1 ./get/

# Run every benchmark: first the package microbenchmarks in the main module, then
# the comparison suite. bench/ is a separate module (so its benchmark-only deps
# stay out of the main module), and run_bench.sh regenerates the decoders, fetches
# the inputs, and benchmarks lightning against the other JSON libraries.
bench:
	go test -run='^$$' -bench=. -benchmem ./...
	bash bench/run_bench.sh

# Generate the committed benchmark results. Two suites: the main-module
# microbenchmarks (pkg_bench.sh → bench/pkg_results.txt and the rendered
# bench/pkg_results_<arch>.md) and the competitor-comparison suite (run_bench.sh →
# bench/results.txt and bench/results_<arch>.md).
#
# The microbenchmarks are short and numerous (dozens of sub-benchmarks in one
# `go test ./...` run), so they run at 3s each — already millions of iterations,
# steady enough for the committed table — to keep the whole suite well under the
# CI wall-clock; the comparison suite runs each case for 10s for steadier numbers.
bench-md:
	BENCHTIME=3s bash pkg_bench.sh
	BENCHTIME=10s bash bench/run_bench.sh

# Force-install the latest version of each developer tool. Unlike a file target,
# a phony recipe runs every time, so @latest is actually re-fetched.
update-tools:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

# Install golangci-lint on demand for `check` when it is not already present.
$(GOBIN)/golangci-lint:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

fix:
	gofmt -w .
	go mod tidy
