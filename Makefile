.PHONY: all download check bench bench-md update-tools fix

GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin
PATH := $(GOBIN):$(PATH)

export PATH

all: fix check

download: go.mod go.sum
	go mod download

check: $(GOBIN)/golangci-lint conformance/data_unmarshal.go
	golangci-lint run ./...
	go test -cover ./...

# The conformance test decodes test.json with a generated decoder, but that
# decoder is gitignored (like every *_unmarshal.go), so a clean checkout has none
# and the package would not compile. Regenerate it from its source struct whenever
# that struct or the generator (main.go) changes; listing it as a prerequisite of
# check makes the test run against an up-to-date decoder.
conformance/data_unmarshal.go: conformance/data.go main.go
	go run . conformance/data.go

# Run every benchmark: first the package microbenchmarks in the main module, then
# the comparison suite. bench/ is a separate module (so its benchmark-only deps
# stay out of the main module), and run_bench.sh regenerates the decoders, fetches
# the inputs, and benchmarks lightning against the other JSON libraries.
bench:
	go test -run='^$$' -bench=. -benchmem ./...
	bash bench/run_bench.sh

# Generate the committed benchmark results (bench/results.txt and the rendered
# bench/results_<arch>.md), running each benchmark for 10s for steadier numbers.
bench-md:
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
