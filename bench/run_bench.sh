#!/usr/bin/env bash
#
# Run from inside the bench module (this directory). For every <n>/ folder that
# contains a data.go struct definition:
#   1. generate the lightning JSON deserializer (data_unmarshal.go) by running
#      the generator (package main) from the parent (lightning) module,
#   2. generate a mailru/easyjson deserializer in a sibling ej/ sub-package
#      (separate package so its UnmarshalJSON does not clash with lightning's),
#   3. write a benchmark that decodes input.json with lightning, encoding/json
#      and easyjson,
#   4. run the benchmark and append the results to results.txt, then render an
#      architecture-specific summary (results_<goarch>.md).
#
# The bench module keeps the easyjson dependency out of the main module.
#
# Usage: bench/run_bench.sh
set -u

cd "$(dirname "$0")"       # the bench module root
ROOT=".."                  # the parent (lightning) module root

RESULTS="results.txt"
# The rendered markdown is per-architecture (results_<goarch>.md) so runs on
# different CPUs — whose SIMD paths and absolute timings differ — do not clobber
# each other's committed results.
RESULTSMD="results_$(go env GOARCH).md"
MODULE=$(go list -m)       # github.com/JohanLindvall/lightning/bench

# Each benchmark runs for BENCHTIME of wall clock (Go's -benchtime) and is
# repeated BENCHCOUNT times (Go's -count). Both default to Go's own defaults
# (1s, one run); override for steadier numbers, e.g. BENCHTIME=30s BENCHCOUNT=6
# run_bench.sh, which gives benchstat enough samples to estimate variance.
BENCHTIME="${BENCHTIME:-1s}"
BENCHCOUNT="${BENCHCOUNT:-1}"

# Locate (or install) the easyjson code generator.
EJBIN="$(command -v easyjson || true)"
if [ -z "$EJBIN" ] && [ -x "$(go env GOPATH)/bin/easyjson" ]; then
	EJBIN="$(go env GOPATH)/bin/easyjson"
fi
if [ -z "$EJBIN" ]; then
	echo "installing easyjson code generator..." >&2
	if go install github.com/mailru/easyjson/easyjson@latest; then
		EJBIN="$(go env GOPATH)/bin/easyjson"
	else
		echo "could not install easyjson; aborting" >&2
		exit 1
	fi
fi

# Build the generator once (in the parent module) so failures surface early.
if ! (cd "$ROOT" && go build -o /dev/null .); then
	echo "failed to build the lightning generator" >&2
	exit 1
fi

{
	echo "# JSON deserialization benchmarks"
	echo "# generated $(date -u +%Y-%m-%dT%H:%M:%SZ)"
	echo "# $(go version)"
	echo
} > "$RESULTS"

shopt -s nullglob
status=0
for dir in */; do
	[ -f "${dir}data.go" ]    || continue
	# A case whose input.json is too large to commit (e.g. bench/large-json) ships
	# a tiny input.url pointing at the data instead; fetch it on demand so the case
	# is runnable from a fresh checkout. A failed download leaves no input.json, so
	# the case is simply skipped below.
	if [ ! -f "${dir}input.json" ] && [ -f "${dir}input.url" ]; then
		url=$(tr -d '[:space:]' < "${dir}input.url")
		echo "==> bench/${dir}: downloading input.json from ${url}"
		# A .gz URL (e.g. the go-json-experiment testdata) is fetched to a temp
		# file and gunzipped into input.json; anything else is fetched directly.
		dl="${dir}input.json"
		case "$url" in
		*.gz) dl="${dir}input.json.gz" ;;
		*.zst) dl="${dir}input.json.zst" ;;
		esac
		if command -v curl >/dev/null 2>&1; then
			curl -fsSL -o "$dl" "$url" || rm -f "$dl"
		elif command -v wget >/dev/null 2>&1; then
			wget -qO "$dl" "$url" || rm -f "$dl"
		else
			echo "  no curl or wget available to download ${dir}input.json" >&2
		fi
		if [ "$dl" != "${dir}input.json" ] && [ -f "$dl" ]; then
			case "$dl" in
			*.gz)  gunzip -f "$dl"      || rm -f "$dl" "${dir}input.json" ;;
			*.zst) zstd -dfq "$dl" && rm -f "$dl" || rm -f "$dl" "${dir}input.json" ;;
			esac
		fi
	fi
	[ -f "${dir}input.json" ] || { echo "skip ${dir}: no input.json" >&2; continue; }
	echo "==> bench/${dir}"

	# 1. Generate the lightning UnmarshalJSON, running the generator from the
	#    parent module where the generator (package main) lives. The //lightning:*
	#    directives a case needs (e.g. cloudflare-compact's :compact, gsoc_2018's
	#    :nocopy) live in its data.go, so the generator call takes no extra arguments.
	if ! (cd "$ROOT" && go run . "bench/${dir}data.go"); then
		echo "  lightning generator failed" >&2
		status=1
		continue
	fi

	# Also build a destructive variant for BenchmarkLightningDestructive: duplicate
	# data.go, rename every top-level type (…Destructive) so it coexists in the same
	# package without clashing, and prepend //lightning:destructive to the root so its
	# UnmarshalJSON unescapes nocopy strings directly into the input buffer instead of
	# allocating. The helper types are renamed too because the generator parses the
	# variant file on its own and must find every type the root references; any
	# :compact / :nocopy directive in the source rides along in the copy.
	dsrc="${dir}data_destructive.go"
	cp "${dir}data.go" "$dsrc"
	names=$(awk '/^type \(/{b=1;next} b&&/^\)/{b=0;next} b&&/^\t[A-Za-z_]/{print $1;next} /^type [A-Za-z_]/{print $2}' "${dir}data.go")
	for n in $names; do gofmt -r "${n} -> ${n}Destructive" -w "$dsrc"; done
	sed -i '0,/^type /s//\/\/lightning:destructive\ntype /' "$dsrc"
	if ! (cd "$ROOT" && go run . "bench/${dsrc}"); then
		echo "  lightning destructive generator failed" >&2
		status=1
		continue
	fi

	# 2. Generate the easyjson deserializer in an ej/ sub-package. The struct is
	#    copied verbatim (only the package clause changes), so the json tags --
	#    including ",nocopy", which easyjson also understands -- carry over.
	ejdir="${dir}ej"
	mkdir -p "$ejdir"
	sed 's/^package .*/package ej/' "${dir}data.go" > "${ejdir}/data.go"
	if ! "$EJBIN" -all "${ejdir}/data.go"; then
		echo "  easyjson generator failed" >&2
		status=1
		continue
	fi
	ejimport="${MODULE}/${dir%/}/ej"

	# 3. Write the benchmark file (the deserialize functionality under test).
	cat > "${dir}bench_test.go" <<'EOF'
// Code generated by run_bench.sh; DO NOT EDIT.
package bench

import (
	"bytes"
	stdjson "encoding/json"
	"os"
	"testing"

	ljson "github.com/JohanLindvall/lightning/pkg/json"
	"github.com/bytedance/sonic"
	jsonv2 "github.com/go-json-experiment/json"
	gojson "github.com/goccy/go-json"
	ej "__EJ_IMPORT__"
)

var benchInput = func() []byte {
	b, err := os.ReadFile("input.json")
	if err != nil {
		panic(err)
	}
	return b
}()

// benchInputCompact is benchInput with insignificant whitespace removed, so the
// DecodeAny benchmark can drive the compact decode path (which assumes no
// inter-token whitespace) on every case regardless of how its input.json is
// formatted — many of the corpus inputs are pretty-printed.
var benchInputCompact = func() []byte {
	var buf bytes.Buffer
	if err := stdjson.Compact(&buf, benchInput); err != nil {
		panic(err)
	}
	return buf.Bytes()
}()

// BenchmarkLightning measures the generated (*Benchmark).UnmarshalJSON.
func BenchmarkLightning(b *testing.B) {
	b.SetBytes(int64(len(benchInput)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v Benchmark
		if err := v.UnmarshalJSON(benchInput); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkLightningDestructive measures the //lightning:destructive variant
// (*BenchmarkDestructive).UnmarshalJSON, which unescapes nocopy string fields
// directly into the input buffer instead of allocating a scratch buffer per escaped
// value. It destroys the input, so each iteration first restores a pristine copy
// into a reused buffer — that copy is a benchmark artifact (a real caller owns and
// discards the buffer), so the gap versus BenchmarkLightning understates the win.
func BenchmarkLightningDestructive(b *testing.B) {
	buf := make([]byte, len(benchInput))
	b.SetBytes(int64(len(benchInput)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		copy(buf, benchInput)
		var v BenchmarkDestructive
		if err := v.UnmarshalJSON(buf); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkLightningDecodeAny measures json.DecodeAny in compact mode: the same
// document decoded dynamically into the generic any value (map[string]any /
// []any / scalars) instead of the typed Benchmark struct, showing the cost of
// schema-less decoding versus the generated unmarshaler.
func BenchmarkLightningDecodeAny(b *testing.B) {
	b.SetBytes(int64(len(benchInputCompact)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := ljson.DecodeAny(benchInputCompact, true); err != nil {
			b.Fatal(err)
		}
	}
}

// benchmarkStd has Benchmark's fields but none of its methods, so encoding/json uses its
// reflection-based decoder, giving an apples-to-apples stdlib baseline.
type benchmarkStd Benchmark

// BenchmarkStdlib measures encoding/json's reflection decoder on the same data.
func BenchmarkStdlib(b *testing.B) {
	b.SetBytes(int64(len(benchInput)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v benchmarkStd
		if err := stdjson.Unmarshal(benchInput, &v); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkEasyjson measures the mailru/easyjson generated (*ej.Benchmark).UnmarshalJSON.
func BenchmarkEasyjson(b *testing.B) {
	b.SetBytes(int64(len(benchInput)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v ej.Benchmark
		if err := v.UnmarshalJSON(benchInput); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSonic measures bytedance/sonic's JIT decoder. It decodes into benchmarkStd
// (the methodless type) so sonic uses its own reflection/JIT decoder rather than
// the generated UnmarshalJSON, an apples-to-apples third-party comparison.
func BenchmarkSonic(b *testing.B) {
	b.SetBytes(int64(len(benchInput)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v benchmarkStd
		if err := sonic.Unmarshal(benchInput, &v); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSonicFastest is BenchmarkSonic with sonic's ConfigFastest, which trades
// strict validation (UTF-8 and duplicate-key checks) for speed.
func BenchmarkSonicFastest(b *testing.B) {
	b.SetBytes(int64(len(benchInput)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v benchmarkStd
		if err := sonic.ConfigFastest.Unmarshal(benchInput, &v); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGoccy measures goccy/go-json, a fast pure-Go drop-in for encoding/json
// (no JIT or codegen). It decodes into benchmarkStd, like the other third-party cases.
func BenchmarkGoccy(b *testing.B) {
	b.SetBytes(int64(len(benchInput)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v benchmarkStd
		if err := gojson.Unmarshal(benchInput, &v); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkJSONV2 measures the json/v2 reference implementation
// (github.com/go-json-experiment/json, the basis of encoding/json/v2). Like the
// stdlib and sonic cases it decodes into benchmarkStd, so it uses its own reflection
// decoder rather than the generated UnmarshalJSON.
func BenchmarkJSONV2(b *testing.B) {
	b.SetBytes(int64(len(benchInput)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v benchmarkStd
		if err := jsonv2.Unmarshal(benchInput, &v); err != nil {
			b.Fatal(err)
		}
	}
}
EOF
	perl -i -pe "s#__EJ_IMPORT__#${ejimport}#" "${dir}bench_test.go"

	# easyjson only emits an UnmarshalJSON for struct types; an array-root document
	# (type Benchmark []T) leaves ej.Benchmark without the method, so drop the
	# easyjson import and benchmark for those cases (lightning still decodes them —
	# it handles named slice roots — and the other libraries are reflection-based).
	if ! grep -rqs 'func (v \*Benchmark) UnmarshalJSON' "${ejdir}"; then
		perl -0777 -i -pe 's{\n\tej "[^"]*"}{}; s{\n// BenchmarkEasyjson.*?\n\}\n}{}s' "${dir}bench_test.go"
	fi

	# 4. Run the benchmarks and record the output.
	{
		echo "================================================================"
		echo "bench/${dir}  (input.json: $(wc -c < "${dir}input.json") bytes)"
		echo "================================================================"
	} >> "$RESULTS"
	if go test -mod=mod -run='^$' -bench=. -benchmem -benchtime="$BENCHTIME" -count="$BENCHCOUNT" "./${dir}" >> "$RESULTS" 2>&1; then
		echo "  ok"
	else
		echo "  benchmark failed (see ${RESULTS})" >&2
		status=1
	fi
	echo >> "$RESULTS"
done

# Render a markdown summary alongside the raw results.
if command -v python3 >/dev/null 2>&1; then
	python3 results_md.py "$RESULTS" "$RESULTSMD" || true
fi

echo
echo "results written to bench/${RESULTS} and bench/${RESULTSMD}"
exit $status
