package unstable

import "testing"

// TestCountArrayElements covers the SkipValue-based element counter: nested
// arrays, strings with structural bytes inside, objects, whitespace, and the
// malformed/empty cases that must return 0 (so the caller falls back to append).
func TestCountArrayElements(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{`[]`, 0},
		{`[1]`, 1},
		{`[1,2,3]`, 3},
		{` [ 1 , 2 , 3 ] `, 3}, // callers must position i at '[' (after their own whitespace skip)
		{`["a","b"]`, 2},
		{`["a,b","c]d"]`, 2},         // commas/brackets inside strings are not counted
		{`[{"x":1},{"y":[2,3]}]`, 2}, // object elements
		{`[[1,2],[3,4],[5,6]]`, 3},   // nested arrays (the coordinate shape)
		{`[[[1,2],[3,4]]]`, 1},       // one outer element
		{`[null,true,false]`, 3},
		{`[`, 0},      // truncated
		{`[1,2`, 0},   // unterminated
		{`[1,2,]`, 0}, // trailing comma -> bail to append
		{`{}`, 0},     // not an array
		{``, 0},       // empty input
	}
	for _, tt := range tests {
		// The caller positions i at '[' (after its own SkipWS); mirror that.
		i := 0
		for i < len(tt.in) && (tt.in[i] == ' ' || tt.in[i] == '\t') {
			i++
		}
		if got := CountArrayElements([]byte(tt.in), i); got != tt.want {
			t.Errorf("CountArrayElements(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

// TestCountArrayObjects covers the cheap counter for arrays of bracket-free
// objects (number/bool fields only): it finds the array's ']' and counts the
// '{' before it. The generator only emits it for structs it has proved hold no
// string/array field, so brace-free element JSON is the contract; a miscount on
// other input is harmless (it only mis-sizes a presize hint).
func TestCountArrayObjects(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{`[]`, 0},
		{`[{"a":1}]`, 1},
		{`[{"a":1,"b":2},{"a":3,"b":4}]`, 2},
		{`[ {"a":1} , {"b":2} , {"c":3} ]`, 3},
		{`[{"amount":10,"id":205705999,"x":-3.5e2,"ok":true}]`, 1},
		{`[`, 0},  // truncated (no ']')
		{`{}`, 0}, // not an array
		{``, 0},   // empty input
	}
	for _, tt := range tests {
		i := 0
		for i < len(tt.in) && (tt.in[i] == ' ' || tt.in[i] == '\t') {
			i++
		}
		if got := CountArrayObjects([]byte(tt.in), i); got != tt.want {
			t.Errorf("CountArrayObjects(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

// TestDecodeValueCompact checks the compact dynamic-decode path: on whitespace-
// free input it must agree with DecodeValue (objects, arrays, and nesting), and
// it still tolerates leading whitespace at the value itself (only inter-token
// whitespace is assumed absent).
