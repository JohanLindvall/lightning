package unstable

import (
	"encoding/json"
	"errors"
	"math"
	"reflect"
	"testing"
)

func TestDecodeFloat64Slice(t *testing.T) {
	cases := []struct {
		in   string
		want []float64
	}{
		{`[]`, nil}, // empty array leaves a nil *out untouched, like the generated loop
		{`[1]`, []float64{1}},
		{`[1,2.5,-3e2]`, []float64{1, 2.5, -300}},
		{`[ 1 , 2 ,\n 3 ]`, []float64{1, 2, 3}},
		{`[null,1,null]`, []float64{0, 1, 0}},
		{`[0.0006988752666567719,-65.613616999999977]`, []float64{0.0006988752666567719, -65.613616999999977}},
		{`[1e308,2.2250738585072014e-308]`, []float64{1e308, 2.2250738585072014e-308}},
		// 20+ significant digits: scanFloat declines, strconv fallback path
		{`[3.141592653589793238462643]`, []float64{3.141592653589793238462643}},
		{`null`, nil},
	}
	for _, c := range cases {
		in := []byte(unescapeNL(c.in))
		var got []float64
		end, err := DecodeFloat64Slice(&got, in, 0)
		if err != nil {
			t.Errorf("DecodeFloat64Slice(%q): %v", c.in, err)
			continue
		}
		if end != len(in) {
			t.Errorf("DecodeFloat64Slice(%q): end = %d, want %d", c.in, end, len(in))
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("DecodeFloat64Slice(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestDecodeFloat64SliceMatchesStdlib(t *testing.T) {
	inputs := []string{
		`[0.1,0.2,0.3]`,
		`[1,22,333,4444,55555,666666]`,
		`[-1.5e-10, 6.02214076e23, 0, -0]`,
		`[ 100.25 ]`,
		"[\n  1.5,\n  2.5\n]",
	}
	for _, in := range inputs {
		var want []float64
		if err := json.Unmarshal([]byte(in), &want); err != nil {
			t.Fatalf("stdlib rejected %q: %v", in, err)
		}
		var got []float64
		if _, err := DecodeFloat64Slice(&got, []byte(in), 0); err != nil {
			t.Fatalf("DecodeFloat64Slice(%q): %v", in, err)
		}
		if len(got) != len(want) {
			t.Fatalf("DecodeFloat64Slice(%q) = %v, want %v", in, got, want)
		}
		for k := range got {
			if math.Float64bits(got[k]) != math.Float64bits(want[k]) {
				t.Errorf("DecodeFloat64Slice(%q)[%d] = %v, want %v (bit mismatch)", in, k, got[k], want[k])
			}
		}
	}
}

func TestDecodeFloat64SliceErrors(t *testing.T) {
	cases := []struct {
		in   string
		want error
	}{
		{``, ErrTruncated},
		{`{`, ErrExpectArray},
		{`[1,2`, ErrTruncated},
		{`[1;2]`, ErrInvalidJSON},
		{`[true]`, ErrBadNumber},
		{`["x"]`, ErrBadNumber},
		{`[1.2.3]`, ErrBadNumber},
		{`nul`, ErrInvalidJSON},
	}
	for _, c := range cases {
		var got []float64
		if _, err := DecodeFloat64Slice(&got, []byte(c.in), 0); !errors.Is(err, c.want) {
			t.Errorf("DecodeFloat64Slice(%q) err = %v, want %v", c.in, err, c.want)
		}
	}
}

// TestDecodeFloat64SliceAppends checks the reuse contract shared with the
// generated loop: a non-nil *out is appended to, not reset.
func TestDecodeFloat64SliceAppends(t *testing.T) {
	got := []float64{9}
	if _, err := DecodeFloat64Slice(&got, []byte(`[1,2]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := []float64{9, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// TestDecodeIntSliceMatchesReader locks the batched element parse to
// ReadInt64OrNull element by element, including the behaviors stdlib does not
// share: truncated fractions/exponents and overflow wrap.
func TestDecodeIntSliceMatchesReader(t *testing.T) {
	inputs := []string{
		`[0,1,-1,42,12345678901234567,-9223372036854775808]`,
		`[9223372036854775808]`, // overflow: wraps exactly as ReadInt64OrNull
		`[1.9,-2.5,3e2,4E-1]`,   // tolerated fraction/exponent, truncated toward zero
		`[null, 7 ,null]`,
		`[ 10 , 200 , 3000 ]`,
		`[]`,
		`null`,
	}
	for _, in := range inputs {
		data := []byte(in)
		var got []int64
		end, err := DecodeIntSlice(&got, data, 0)
		if err != nil {
			t.Fatalf("DecodeIntSlice(%q): %v", in, err)
		}
		if end != len(data) {
			t.Errorf("DecodeIntSlice(%q): end = %d, want %d", in, end, len(data))
		}
		want := readInt64ArrayReference(t, data)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DecodeIntSlice(%q) = %v, want %v", in, got, want)
		}
	}
}

// readInt64ArrayReference decodes a JSON integer array the way the generated
// per-element loop did: SkipWS + ReadInt64OrNull per element.
func readInt64ArrayReference(t *testing.T, data []byte) []int64 {
	t.Helper()
	i := SkipWS(data, 0)
	if data[i] == 'n' {
		return nil
	}
	i++ // '['
	var out []int64
	for {
		i = SkipWS(data, i)
		if data[i] == ']' {
			return out
		}
		n, end, err := ReadInt64OrNull(data, i)
		if err != nil {
			t.Fatalf("reference ReadInt64OrNull at %d: %v", i, err)
		}
		out = append(out, n)
		i = SkipWS(data, end)
		if data[i] == ']' {
			return out
		}
		i++ // ','
	}
}

func TestDecodeIntSliceKinds(t *testing.T) {
	var i16 []int16
	if _, err := DecodeIntSlice(&i16, []byte(`[32000,-32000,70000]`), 0); err != nil {
		t.Fatal(err)
	}
	// 70000 wraps into int16 exactly as the generated int16(n) conversion did.
	wrap := int64(70000)
	if want := []int16{32000, -32000, int16(wrap)}; !reflect.DeepEqual(i16, want) {
		t.Errorf("got %v, want %v", i16, want)
	}

	var u []uint64
	if _, err := DecodeUintSlice(&u, []byte(`[0,18446744073709551615,null]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := []uint64{0, math.MaxUint64, 0}; !reflect.DeepEqual(u, want) {
		t.Errorf("got %v, want %v", u, want)
	}

	var u8 []uint8
	if _, err := DecodeUintSlice(&u8, []byte(`[1,2,255]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := []uint8{1, 2, 255}; !reflect.DeepEqual(u8, want) {
		t.Errorf("got %v, want %v", u8, want)
	}

	var bad []int
	if _, err := DecodeUintSlice(&u, []byte(`[-1]`), 0); !errors.Is(err, ErrBadNumber) {
		t.Errorf("DecodeUintSlice([-1]) err = %v, want ErrBadNumber", err)
	}
	if _, err := DecodeIntSlice(&bad, []byte(`[-]`), 0); !errors.Is(err, ErrBadNumber) {
		t.Errorf("DecodeIntSlice([-]) err = %v, want ErrBadNumber", err)
	}
}

func TestDecodeFloat64Array(t *testing.T) {
	// Exact fill.
	arr := [2]float64{9, 9}
	end, err := DecodeFloat64Array(arr[:], []byte(`[-65.613617,43.420273]`), 0)
	if err != nil {
		t.Fatal(err)
	}
	if end != len(`[-65.613617,43.420273]`) {
		t.Errorf("end = %d", end)
	}
	if want := [2]float64{-65.613617, 43.420273}; arr != want {
		t.Errorf("got %v, want %v", arr, want)
	}

	// Short JSON array: array is zeroed first, tail stays zero.
	arr = [2]float64{9, 9}
	if _, err := DecodeFloat64Array(arr[:], []byte(`[1]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := [2]float64{1, 0}; arr != want {
		t.Errorf("short: got %v, want %v", arr, want)
	}

	// Long JSON array: extras are skipped whole, even non-numbers.
	arr = [2]float64{}
	if _, err := DecodeFloat64Array(arr[:], []byte(`[1,2,{"a":[3]},"x"]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := [2]float64{1, 2}; arr != want {
		t.Errorf("long: got %v, want %v", arr, want)
	}

	// Null root leaves the array untouched (matching the generated decoder).
	arr = [2]float64{7, 8}
	if _, err := DecodeFloat64Array(arr[:], []byte(`null`), 0); err != nil {
		t.Fatal(err)
	}
	if want := [2]float64{7, 8}; arr != want {
		t.Errorf("null: got %v, want %v", arr, want)
	}

	// Null element decodes as zero.
	arr = [2]float64{7, 8}
	if _, err := DecodeFloat64Array(arr[:], []byte(`[null,2]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := [2]float64{0, 2}; arr != want {
		t.Errorf("null elem: got %v, want %v", arr, want)
	}

	if _, err := DecodeFloat64Array(arr[:], []byte(`[1,`), 0); !errors.Is(err, ErrTruncated) {
		t.Errorf("truncated err = %v", err)
	}
	if _, err := DecodeFloat64Array(arr[:], []byte(`7`), 0); !errors.Is(err, ErrExpectArray) {
		t.Errorf("non-array err = %v", err)
	}
}

func TestDecodeIntArray(t *testing.T) {
	arr := [3]int{9, 9, 9}
	if _, err := DecodeIntArray(arr[:], []byte(` [ 1 , -2 ] `), 1); err != nil {
		t.Fatal(err)
	}
	if want := [3]int{1, -2, 0}; arr != want {
		t.Errorf("got %v, want %v", arr, want)
	}

	arr = [3]int{1, 1, 1}
	if _, err := DecodeIntArray(arr[:], []byte(`[null,5,6,7,8]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := [3]int{0, 5, 6}; arr != want {
		t.Errorf("got %v, want %v", arr, want)
	}
}

// unescapeNL turns literal \n escapes in test-case strings into newlines, so a
// case can spell whitespace-run inputs inline.
func unescapeNL(s string) string {
	out := make([]byte, 0, len(s))
	for k := 0; k < len(s); k++ {
		if s[k] == '\\' && k+1 < len(s) && s[k+1] == 'n' {
			out = append(out, '\n')
			k++
			continue
		}
		out = append(out, s[k])
	}
	return string(out)
}
