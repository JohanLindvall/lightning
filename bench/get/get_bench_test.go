// Package bench (get) compares lightning's json.Get key-path lookup against
// github.com/buger/jsonparser's Get over a realistic nested document. Both walk
// an object-key path and return the raw value bytes; lightning is built on the
// scanner's allocation-free Skip/ReadKey primitives.
package bench

import (
	"bytes"
	"os"
	"testing"

	"github.com/JohanLindvall/lightning/pkg/json"
	"github.com/buger/jsonparser"
)

var benchInput = func() []byte {
	b, err := os.ReadFile("input.json")
	if err != nil {
		panic(err)
	}
	return b
}()

// cloudflareInput is the flat cloudflare log record shared with the decode
// benchmarks (sibling bench/cloudflare). Unlike the deeply nested Kubernetes
// document above, it is a single wide object of ~50 scalar/string members —
// the shape where the common job is plucking a handful of top-level fields out
// of each record. It is read from the sibling directory rather than copied so
// the two benchmark suites stay on identical data.
var cloudflareInput = func() []byte {
	b, err := os.ReadFile("../cloudflare/input.json")
	if err != nil {
		panic(err)
	}
	return b
}()

// cloudflareTopKeys are six top-level fields spread from the front of the record
// to near its end and covering both string and numeric values, so the benchmark
// exercises the full range of skip distances a real field-extraction workload
// hits (an early key returns almost immediately; a late one must scan past every
// preceding member).
var cloudflareTopKeys = []string{
	"ClientIP",               // string, near the front
	"CacheResponseStatus",    // number, near the front
	"ClientRequestUserAgent", // long string, the middle
	"EdgeResponseStatus",     // number, the middle
	"OriginResponseStatus",   // number, the back
	"RayID",                  // string, near the end
}

// paths covers the access patterns that matter for a key-path lookup: a key at
// the very front of the document, deeply nested objects and scalars, and a key
// near the end of a long object (the worst case, since the scan must skip every
// preceding member).
var paths = []struct {
	name string
	keys []string
}{
	{"shallow_first_key", []string{"apiVersion"}},
	{"nested_scalar", []string{"metadata", "labels", "version"}},
	{"nested_object", []string{"spec", "strategy", "rollingUpdate"}},
	{"deep_object", []string{"spec", "template", "spec"}},
	{"late_key_in_long_object", []string{"status", "availableReplicas"}},
	{"deep_string_with_escape", []string{"status", "conditions"}},
}

// TestGetMatchesJSONParser pins down parity: for every path both libraries must
// agree the key exists and return the same value bytes. jsonparser strips the
// surrounding quotes from a string value, so a string result is compared with
// those quotes trimmed.
func TestGetMatchesJSONParser(t *testing.T) {
	for _, p := range paths {
		got, _, err := json.Get(benchInput, p.keys...)
		if err != nil {
			t.Fatalf("%s: lightning Get error: %v", p.name, err)
		}
		want, dt, _, err := jsonparser.Get(benchInput, p.keys...)
		if err != nil {
			t.Fatalf("%s: jsonparser Get error: %v", p.name, err)
		}
		g := got
		if dt == jsonparser.String { // jsonparser returns string bodies unquoted
			g = bytes.Trim(g, `"`)
		}
		if !bytes.Equal(g, want) {
			t.Fatalf("%s: got %q, jsonparser %q", p.name, g, want)
		}
	}
}

func BenchmarkLightningGet(b *testing.B) {
	for _, p := range paths {
		b.Run(p.name, func(b *testing.B) {
			b.SetBytes(int64(len(benchInput)))
			b.ReportAllocs()
			var sink []byte
			for i := 0; i < b.N; i++ {
				v, _, err := json.Get(benchInput, p.keys...)
				if err != nil {
					b.Fatal(err)
				}
				sink = v
			}
			_ = sink
		})
	}
}

func BenchmarkJSONParserGet(b *testing.B) {
	for _, p := range paths {
		b.Run(p.name, func(b *testing.B) {
			b.SetBytes(int64(len(benchInput)))
			b.ReportAllocs()
			var sink []byte
			for i := 0; i < b.N; i++ {
				v, _, _, err := jsonparser.Get(benchInput, p.keys...)
				if err != nil {
					b.Fatal(err)
				}
				sink = v
			}
			_ = sink
		})
	}
}

// objectEachPaths are the object-valued paths whose members ObjectEach walks.
var objectEachPaths = [][]string{
	nil,                          // the document root object
	{"metadata", "labels"},       // a small leaf object
	{"spec", "template", "spec"}, // a larger nested object
}

// TestObjectEachMatchesJSONParser confirms both libraries visit the same members
// in the same order with the same value bytes.
func TestObjectEachMatchesJSONParser(t *testing.T) {
	for _, keys := range objectEachPaths {
		var got, want []string
		if err := json.ObjectEach(benchInput, func(k string, v []byte) error {
			got = append(got, k+"="+string(v))
			return nil
		}, keys...); err != nil {
			t.Fatalf("%v: lightning ObjectEach: %v", keys, err)
		}
		err := jsonparser.ObjectEach(benchInput, func(k, v []byte, dt jsonparser.ValueType, _ int) error {
			s := string(v)
			if dt == jsonparser.String {
				s = `"` + s + `"` // jsonparser strips a string value's quotes
			}
			want = append(want, string(k)+"="+s)
			return nil
		}, keys...)
		if err != nil {
			t.Fatalf("%v: jsonparser ObjectEach: %v", keys, err)
		}
		if len(got) != len(want) {
			t.Fatalf("%v: got %d members, jsonparser %d", keys, len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Fatalf("%v member %d: got %q, jsonparser %q", keys, i, got[i], want[i])
			}
		}
	}
}

func BenchmarkLightningObjectEach(b *testing.B) {
	for _, keys := range objectEachPaths {
		b.Run(pathName(keys), func(b *testing.B) {
			b.SetBytes(int64(len(benchInput)))
			b.ReportAllocs()
			var sink int
			for i := 0; i < b.N; i++ {
				if err := json.ObjectEach(benchInput, func(k string, v []byte) error {
					sink += len(k) + len(v)
					return nil
				}, keys...); err != nil {
					b.Fatal(err)
				}
			}
			_ = sink
		})
	}
}

func BenchmarkJSONParserObjectEach(b *testing.B) {
	for _, keys := range objectEachPaths {
		b.Run(pathName(keys), func(b *testing.B) {
			b.SetBytes(int64(len(benchInput)))
			b.ReportAllocs()
			var sink int
			for i := 0; i < b.N; i++ {
				err := jsonparser.ObjectEach(benchInput, func(k, v []byte, _ jsonparser.ValueType, _ int) error {
					sink += len(k) + len(v)
					return nil
				}, keys...)
				if err != nil {
					b.Fatal(err)
				}
			}
			_ = sink
		})
	}
}

// TestGetCloudflareMatchesJSONParser pins down parity for the flat-record path:
// for every top-level key both libraries must agree on the value bytes.
func TestGetCloudflareMatchesJSONParser(t *testing.T) {
	for _, key := range cloudflareTopKeys {
		got, _, err := json.Get(cloudflareInput, key)
		if err != nil {
			t.Fatalf("%s: lightning Get error: %v", key, err)
		}
		want, dt, _, err := jsonparser.Get(cloudflareInput, key)
		if err != nil {
			t.Fatalf("%s: jsonparser Get error: %v", key, err)
		}
		g := got
		if dt == jsonparser.String { // jsonparser returns string bodies unquoted
			g = bytes.Trim(g, `"`)
		}
		if !bytes.Equal(g, want) {
			t.Fatalf("%s: got %q, jsonparser %q", key, g, want)
		}
	}
}

// BenchmarkLightningGetCloudflare models extracting a fixed set of fields from a
// flat log record: every iteration fetches all six cloudflareTopKeys. SetBytes
// reports the record size, so MB/s is records-per-second scaled by the document
// length (the six lookups together is one unit of work).
func BenchmarkLightningGetCloudflare(b *testing.B) {
	b.SetBytes(int64(len(cloudflareInput)))
	b.ReportAllocs()
	var sink int
	for i := 0; i < b.N; i++ {
		for _, key := range cloudflareTopKeys {
			v, _, err := json.Get(cloudflareInput, key)
			if err != nil {
				b.Fatal(err)
			}
			sink += len(v)
		}
	}
	_ = sink
}

func BenchmarkJSONParserGetCloudflare(b *testing.B) {
	b.SetBytes(int64(len(cloudflareInput)))
	b.ReportAllocs()
	var sink int
	for i := 0; i < b.N; i++ {
		for _, key := range cloudflareTopKeys {
			v, _, _, err := jsonparser.Get(cloudflareInput, key)
			if err != nil {
				b.Fatal(err)
			}
			sink += len(v)
		}
	}
	_ = sink
}

// TestGetManyMatchesGet confirms the single-pass GetMany returns, for each
// requested key, exactly what a separate Get for that key returns.
func TestGetManyMatchesGet(t *testing.T) {
	out, err := json.GetMany(cloudflareInput, cloudflareTopKeys, nil)
	if err != nil {
		t.Fatalf("GetMany: %v", err)
	}
	for n, key := range cloudflareTopKeys {
		want, _, err := json.Get(cloudflareInput, key)
		if err != nil {
			t.Fatalf("%s: Get error: %v", key, err)
		}
		if !bytes.Equal(out[n], want) {
			t.Fatalf("%s: GetMany %q, Get %q", key, out[n], want)
		}
	}
}

// BenchmarkLightningGetManyCloudflare is the batch counterpart of
// BenchmarkLightningGetCloudflare: it pulls the same six fields, but in one pass
// over the record (with a reused scratch slice) instead of one Get per field.
func BenchmarkLightningGetManyCloudflare(b *testing.B) {
	b.SetBytes(int64(len(cloudflareInput)))
	b.ReportAllocs()
	out := make([][]byte, 0, len(cloudflareTopKeys))
	var sink int
	for i := 0; i < b.N; i++ {
		var err error
		out, err = json.GetMany(cloudflareInput, cloudflareTopKeys, out)
		if err != nil {
			b.Fatal(err)
		}
		for _, v := range out {
			sink += len(v)
		}
	}
	_ = sink
}

// cloudflareInput is whitespace-free between tokens, so the compact variants
// decode it correctly; these benchmarks isolate the win from skipping the
// inter-token SkipWS scans.
func BenchmarkLightningGetManyCompactCloudflare(b *testing.B) {
	b.SetBytes(int64(len(cloudflareInput)))
	b.ReportAllocs()
	out := make([][]byte, 0, len(cloudflareTopKeys))
	var sink int
	for i := 0; i < b.N; i++ {
		var err error
		out, err = json.GetManyCompact(cloudflareInput, cloudflareTopKeys, out)
		if err != nil {
			b.Fatal(err)
		}
		for _, v := range out {
			sink += len(v)
		}
	}
	_ = sink
}

func pathName(keys []string) string {
	if len(keys) == 0 {
		return "root"
	}
	name := keys[0]
	for _, k := range keys[1:] {
		name += "." + k
	}
	return name
}
