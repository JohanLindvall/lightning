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
		// Trailing commas are rejected (the first-iteration flag: ']' at the
		// loop top is only reachable after a comma), as in encoding/json.
		{`[1,]`, ErrInvalidJSON},
		{`[1, ]`, ErrInvalidJSON},
		{`[1,`, ErrTruncated},
	}
	for _, c := range cases {
		var got []float64
		if _, err := DecodeFloat64Slice(&got, []byte(c.in), 0); !errors.Is(err, c.want) {
			t.Errorf("DecodeFloat64Slice(%q) err = %v, want %v", c.in, err, c.want)
		}
	}
}

// TestDecodeFloat64SliceReplaces checks the reuse contract shared with the
// generated loop: a non-nil *out has its length reset and is then filled, so
// decoding replaces its contents rather than appending to them. This is the rule
// encoding/json documents ("Unmarshal resets the slice length to zero and then
// appends each element to the slice").
//
// This test previously asserted the opposite — that a non-nil *out was appended to
// — which made decoding twice into one value accumulate ([1,2] read twice became
// [1,2,1,2]) and made a caller reusing a target to avoid allocation grow it without
// bound. The backing array is still reused, which is what keeps that pattern
// allocation-free; TestDecodeFloat64SliceReusesBacking covers that half.
func TestDecodeFloat64SliceReplaces(t *testing.T) {
	got := []float64{9}
	if _, err := DecodeFloat64Slice(&got, []byte(`[1,2]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := []float64{1, 2}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	// Decoding again must be idempotent, not cumulative.
	if _, err := DecodeFloat64Slice(&got, []byte(`[1,2]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := []float64{1, 2}; !reflect.DeepEqual(got, want) {
		t.Errorf("second decode got %v, want %v", got, want)
	}
	// A shorter array must shrink the result, not leave a stale tail.
	if _, err := DecodeFloat64Slice(&got, []byte(`[7]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := []float64{7}; !reflect.DeepEqual(got, want) {
		t.Errorf("shorter array got %v, want %v", got, want)
	}
}

// TestDecodeFloat64SliceReusesBacking is the performance half of the reset above:
// the length is reset but the backing array is kept, so repeated decoding into one
// slice never reallocates. Resetting by assigning a fresh slice instead would have
// turned the correctness fix into a per-decode allocation.
func TestDecodeFloat64SliceReusesBacking(t *testing.T) {
	var got []float64
	if _, err := DecodeFloat64Slice(&got, []byte(`[1,2,3,4,5,6,7,8]`), 0); err != nil {
		t.Fatal(err)
	}
	addr, capBefore := &got[0], cap(got)
	for i := 0; i < 3; i++ {
		if _, err := DecodeFloat64Slice(&got, []byte(`[1,2,3,4,5,6,7,8]`), 0); err != nil {
			t.Fatal(err)
		}
	}
	if &got[0] != addr || cap(got) != capBefore {
		t.Errorf("backing array replaced on reuse (cap %d -> %d)", capBefore, cap(got))
	}
}

// TestDecodeIntUintSliceReplace covers the same reset for the generic integer
// readers, which share the shape but not the code path.
func TestDecodeIntUintSliceReplace(t *testing.T) {
	ints := []int64{99, 98}
	if _, err := DecodeIntSlice(&ints, []byte(`[1,2]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := []int64{1, 2}; !reflect.DeepEqual(ints, want) {
		t.Errorf("DecodeIntSlice got %v, want %v", ints, want)
	}
	uints := []uint32{99}
	if _, err := DecodeUintSlice(&uints, []byte(`[5]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := []uint32{5}; !reflect.DeepEqual(uints, want) {
		t.Errorf("DecodeUintSlice got %v, want %v", uints, want)
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

func TestDecodeUintArray(t *testing.T) {
	arr := [3]uint32{9, 9, 9}
	if _, err := DecodeUintArray(arr[:], []byte(` [ 1 , 4000000000 ] `), 1); err != nil {
		t.Fatal(err)
	}
	if want := [3]uint32{1, 4000000000, 0}; arr != want {
		t.Errorf("got %v, want %v", arr, want)
	}

	arr = [3]uint32{1, 1, 1}
	if _, err := DecodeUintArray(arr[:], []byte(`[null,5,6,7,8]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := [3]uint32{0, 5, 6}; arr != want {
		t.Errorf("got %v, want %v", arr, want)
	}

	// Null root leaves the array untouched; overflow wraps like ReadUint64OrNull.
	arr = [3]uint32{7, 8, 9}
	if _, err := DecodeUintArray(arr[:], []byte(`null`), 0); err != nil {
		t.Fatal(err)
	}
	if want := [3]uint32{7, 8, 9}; arr != want {
		t.Errorf("null: got %v, want %v", arr, want)
	}
	arr = [3]uint32{}
	if _, err := DecodeUintArray(arr[:], []byte(`[4294967297]`), 0); err != nil {
		t.Fatal(err)
	}
	if want := [3]uint32{1, 0, 0}; arr != want {
		t.Errorf("wrap: got %v, want %v", arr, want)
	}
	if _, err := DecodeUintArray(arr[:], []byte(`[-1]`), 0); !errors.Is(err, ErrBadNumber) {
		t.Errorf("negative err = %v", err)
	}
}

// TestBatchTrailingComma locks the first-iteration flag across every batched
// reader: a trailing comma before ']' is an error, as in encoding/json.
func TestBatchTrailingComma(t *testing.T) {
	var f []float64
	if _, err := DecodeFloat64Slice(&f, []byte(`[1.5,]`), 0); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("float slice err = %v", err)
	}
	var n []int
	if _, err := DecodeIntSlice(&n, []byte(`[1,]`), 0); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("int slice err = %v", err)
	}
	var u []uint
	if _, err := DecodeUintSlice(&u, []byte(`[1,]`), 0); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("uint slice err = %v", err)
	}
	fa := [2]float64{}
	if _, err := DecodeFloat64Array(fa[:], []byte(`[1,]`), 0); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("float array err = %v", err)
	}
	ia := [2]int{}
	if _, err := DecodeIntArray(ia[:], []byte(`[1,]`), 0); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("int array err = %v", err)
	}
	ua := [2]uint{}
	if _, err := DecodeUintArray(ua[:], []byte(`[1,]`), 0); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("uint array err = %v", err)
	}
	// Also past the fixed array's length (the extras-skipped branch).
	sa := [1]int{}
	if _, err := DecodeIntArray(sa[:], []byte(`[1,2,]`), 0); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("skip branch err = %v", err)
	}
	// Empty arrays are unaffected.
	if _, err := DecodeIntSlice(&n, []byte(`[]`), 0); err != nil {
		t.Errorf("empty: %v", err)
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

// TestDecodeIntArrayMatchesReader locks the fixed-size array reader's inlined
// element parse to ReadInt64OrNull element by element, mirroring what
// TestDecodeIntSliceMatchesReader locks for the slice reader: overflow wrap,
// tolerated fraction/exponent truncation, null elements, whitespace, and the
// fixed-array-specific semantics (zero first, ignore surplus JSON elements,
// leave a short JSON array's tail zero).
func TestDecodeIntArrayMatchesReader(t *testing.T) {
	inputs := []string{
		`[0,1,-1,42]`,
		`[12345678901234567,-9223372036854775808,7,8]`,
		`[9223372036854775808]`, // overflow: wraps exactly as ReadInt64OrNull
		`[1.9,-2.5,3e2,4E-1]`,   // tolerated fraction/exponent, truncated toward zero
		`[null, 7 ,null]`,
		`[ 10 , 200 , 3000 ]`,
		`[1,2,3,4,5,6]`, // more elements than the array: surplus skipped
		`[]`,
	}
	for _, in := range inputs {
		data := []byte(in)
		got := [4]int64{-77, -77, -77, -77} // pre-dirtied: reader must zero it
		end, err := DecodeIntArray(got[:], data, 0)
		if err != nil {
			t.Fatalf("DecodeIntArray(%q): %v", in, err)
		}
		if end != len(data) {
			t.Errorf("DecodeIntArray(%q): end = %d, want %d", in, end, len(data))
		}
		var want [4]int64
		for i, n := range readInt64ArrayReference(t, data) {
			if i == len(want) {
				break
			}
			want[i] = n
		}
		if got != want {
			t.Errorf("DecodeIntArray(%q) = %v, want %v", in, got, want)
		}
	}
}

// TestDecodeUintArrayMatchesReader is the ReadUint64OrNull twin of
// TestDecodeIntArrayMatchesReader.
func TestDecodeUintArrayMatchesReader(t *testing.T) {
	inputs := []string{
		`[0,1,42,4294967297]`,
		`[18446744073709551615,7,8,9]`,
		`[18446744073709551616]`, // overflow: wraps exactly as ReadUint64OrNull
		`[1.9,2.5,3e2,4E-1]`,     // tolerated fraction/exponent
		`[null, 7 ,null]`,
		`[1,2,3,4,5,6]`,
		`[]`,
	}
	for _, in := range inputs {
		data := []byte(in)
		got := [4]uint64{77, 77, 77, 77}
		end, err := DecodeUintArray(got[:], data, 0)
		if err != nil {
			t.Fatalf("DecodeUintArray(%q): %v", in, err)
		}
		if end != len(data) {
			t.Errorf("DecodeUintArray(%q): end = %d, want %d", in, end, len(data))
		}
		var want [4]uint64
		i := SkipWS(data, 0)
		i++ // '['
		idx := 0
		for {
			i = SkipWS(data, i)
			if data[i] == ']' {
				break
			}
			n, end, err := ReadUint64OrNull(data, i)
			if err != nil {
				t.Fatalf("reference ReadUint64OrNull at %d: %v", i, err)
			}
			if idx < len(want) {
				want[idx] = n
			}
			idx++
			i = SkipWS(data, end)
			if data[i] == ']' {
				break
			}
			i++ // ','
		}
		if got != want {
			t.Errorf("DecodeUintArray(%q) = %v, want %v", in, got, want)
		}
	}
}

// TestDecodeByteSlice covers the []byte reader's base64 arm directly: reuse
// replaces, backing is reused when it fits, and the numeric/null arms delegate
// to the shared uint reader.
func TestDecodeByteSlice(t *testing.T) {
	var b []byte
	if _, err := DecodeByteSlice(&b, []byte(`"aGVsbG8="`), 0); err != nil || string(b) != "hello" {
		t.Fatalf("base64: %q, %v", b, err)
	}
	backing := &b[0]
	if _, err := DecodeByteSlice(&b, []byte(`"aGk="`), 0); err != nil || string(b) != "hi" {
		t.Fatalf("reuse: %q, %v", b, err)
	}
	if &b[0] != backing {
		t.Error("smaller decode did not reuse the backing")
	}
	if _, err := DecodeByteSlice(&b, []byte(`[104,105,33]`), 0); err != nil || string(b) != "hi!" {
		t.Fatalf("numeric: %q, %v", b, err)
	}
	if _, err := DecodeByteSlice(&b, []byte(`null`), 0); err != nil || b != nil {
		t.Fatalf("null: %v, %v", b, err)
	}
	if _, err := DecodeByteSlice(&b, []byte(`""`), 0); err != nil || b == nil || len(b) != 0 {
		t.Fatalf("empty base64: %#v, %v", b, err)
	}
	if _, err := DecodeByteSlice(&b, []byte(`"???"`), 0); err == nil {
		t.Fatal("invalid base64 accepted")
	}
}
