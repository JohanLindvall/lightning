package bench

import (
	stdjson "encoding/json"
	"reflect"
	"testing"
)

// TestCompactDecode guards the //lightning:compact decoder: on whitespace-free
// (between-token) input it must agree with encoding/json, it must still tolerate
// surrounding whitespace (the wrapper keeps its leading/trailing SkipWS), and it
// is expected to reject inter-token whitespace — the contract the tag asserts.
func TestCompactDecode(t *testing.T) {
	t.Run("matches stdlib", func(t *testing.T) {
		var got, want Benchmark
		if err := got.UnmarshalJSON(benchInput); err != nil {
			t.Fatalf("compact UnmarshalJSON: %v", err)
		}
		if err := stdjson.Unmarshal(benchInput, &want); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("compact decode != stdlib\n got %+v\nwant %+v", got, want)
		}
	})

	t.Run("tolerates surrounding whitespace", func(t *testing.T) {
		framed := append([]byte("\n  "), benchInput...)
		framed = append(framed, " \n"...)
		var v Benchmark
		if err := v.UnmarshalJSON(framed); err != nil {
			t.Fatalf("compact decoder rejected leading/trailing whitespace: %v", err)
		}
	})

	t.Run("rejects inter-token whitespace", func(t *testing.T) {
		var v Benchmark
		// A space after the first key's colon is between-token whitespace, which
		// the compact decoder is permitted to reject.
		if err := v.UnmarshalJSON([]byte(`{"EventTimestampMs": 1}`)); err == nil {
			t.Fatal("expected the compact decoder to reject inter-token whitespace")
		}
	})
}
