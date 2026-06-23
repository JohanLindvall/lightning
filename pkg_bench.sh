#!/usr/bin/env bash
#
# Run the main-module microbenchmarks — the Benchmark* functions that live in the
# lightning module itself (pkg/json, pkg/unstable, …), as opposed to the
# competitor-comparison suite in the separate bench/ module (see bench/run_bench.sh).
#
# It runs `go test -bench` over the whole main module, appends the raw output to
# bench/pkg_results.txt, then renders an architecture-specific markdown summary
# (bench/pkg_results_<goarch>.md) so runs on different CPUs do not clobber each
# other's committed results.
#
# Usage: pkg_bench.sh   (or `make bench-md`, which also runs the comparison suite)
set -u

cd "$(dirname "$0")"       # the lightning module root

RESULTS="bench/pkg_results.txt"
# Per-architecture so amd64/arm64 results (different SIMD paths, different absolute
# timings) do not overwrite each other.
RESULTSMD="bench/pkg_results_$(go env GOARCH).md"

# Each benchmark runs for BENCHTIME of wall clock (Go's -benchtime), repeated
# BENCHCOUNT times (Go's -count). Both default to Go's defaults; override for
# steadier numbers, e.g. BENCHTIME=10s pkg_bench.sh.
BENCHTIME="${BENCHTIME:-1s}"
BENCHCOUNT="${BENCHCOUNT:-1}"

{
	echo "# lightning main-module benchmarks"
	echo "# generated $(date -u +%Y-%m-%dT%H:%M:%SZ)"
	echo "# $(go version)"
	echo
} > "$RESULTS"

# -run='^$' disables the unit tests so only benchmarks run. The target is ./pkg/...,
# not ./..., on purpose: every main-module Benchmark* lives under pkg/ (pkg/json,
# pkg/unstable), and the other packages (conformance, the root generator) import
# *generated* decoders that are gitignored — absent on a fresh checkout (e.g. CI,
# which does not `make generate` before benchmarking), so `go test ./...` would fail
# to build them and report a benchmark failure even though no benchmark ran there.
# The bench/ comparison suite is a separate module and is not matched either way. A
# named benchmark filter can be passed as the first argument.
#
# -timeout=0 disables Go's default 10-minute test timeout: at a long BENCHTIME the
# whole suite (dozens of sub-benchmarks in one `go test` run) easily exceeds 10
# minutes and would otherwise be killed mid-run and reported as a failure.
status=0
if go test -run='^$' -bench="${1:-.}" -benchmem -timeout=0 -benchtime="$BENCHTIME" -count="$BENCHCOUNT" ./pkg/... >> "$RESULTS" 2>&1; then
	echo "  ok"
else
	echo "  benchmark failed (see ${RESULTS})" >&2
	status=1
fi

# Render a markdown summary alongside the raw results.
if command -v python3 >/dev/null 2>&1; then
	python3 bench/pkg_results_md.py "$RESULTS" "$RESULTSMD" || true
fi

echo
echo "results written to ${RESULTS} and ${RESULTSMD}"
exit $status
