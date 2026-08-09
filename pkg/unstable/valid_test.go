package unstable

import (
	"errors"
	"strings"
	"testing"
)

// TestSkipValueStrictEndMatchesSkipValue is the property the ",lax" caller
// depends on and pkg/json.Valid does not: the *end offset*. Valid only reads the
// error (it re-checks that nothing but whitespace follows), so a walk that
// returned the wrong end would still pass every Valid test — including the
// 14-million-execution differential fuzz — while making a generated decoder
// resume in the middle of the value it just skipped.
//
// SkipValue is the oracle because it is the walk the lax skip used before, and
// the two must agree on every well-formed value; they are allowed to differ only
// in *which* malformed input they reject (SkipValue balances brackets and
// accepts a great deal that this rejects, which is the whole point of the
// change).
func TestSkipValueStrictEndMatchesSkipValue(t *testing.T) {
	values := []string{
		`null`, `true`, `false`, `0`, `-1.5e-3`, `123456789012345678901234567890`,
		`""`, `"x"`, `"a\"b"`, `"é😀"`, `"\\"`, `"a b c"`,
		`{}`, `[]`, `{"a":1}`, `[1,2,3]`, `[[[]]]`, `{"a":{"b":[1,{"c":null}]}}`,
		`{ "a" : 1 , "b" : [ 2 , 3 ] }`, "{\n\t\"a\": [1,\r2]\n}",
		`["a","b",{"c":[true,false,null]}]`, `{"a":"}","b":"]"}`,
		`{"esc":"\\\"{[,"}`, `[{"x":[{"y":[]}]}]`,
	}
	for _, v := range values {
		// A trailing tail proves the walk stops at the value's end rather than
		// running to the end of the buffer.
		for _, tail := range []string{"", ",7]", " , \"next\" : 1 }", "xyz"} {
			in := []byte(v + tail)
			wantEnd, wantErr := SkipValue(in, 0)
			gotEnd, gotErr := SkipValueStrict(in, 0)
			if wantErr != nil {
				t.Fatalf("oracle rejected well-formed %q: %v", v, wantErr)
			}
			if gotErr != nil {
				t.Errorf("SkipValueStrict(%q) = error %v, want nil", v+tail, gotErr)
				continue
			}
			if gotEnd != wantEnd {
				t.Errorf("SkipValueStrict(%q) end = %d, SkipValue = %d (value %q)",
					v+tail, gotEnd, wantEnd, v)
			}
			if gotEnd != len(v) {
				t.Errorf("SkipValueStrict(%q) end = %d, want %d (just past the value)",
					v+tail, gotEnd, len(v))
			}
		}
	}
}

// TestSkipValueStrictRejects covers the inputs the lenient balancer accepts and
// this must not, each with the sentinel a caller would see. The error identities
// are checked, not just non-nil, because they are what a ",lax" field surfaces
// to the caller of UnmarshalJSON, and they are chosen to match what the
// generated container loops report for the same input.
func TestSkipValueStrictRejects(t *testing.T) {
	cases := []struct {
		in   string
		want error
	}{
		{`[1,]`, ErrInvalidJSON},
		{`{"a":1,}`, ErrInvalidJSON},
		{`[1 2 3]`, ErrInvalidJSON},
		{`[1,,2]`, ErrBadNumber},
		{`{"a" 1}`, ErrExpectColon},
		{`{a:1}`, ErrInvalidJSON},
		{`{1:2}`, ErrInvalidJSON},
		{`[1,2}`, ErrInvalidJSON},
		{`{"a":[}`, ErrInvalidJSON},
		{`[tru]`, ErrInvalidJSON},
		{`"a\q"`, ErrBadEscape},
		{`"\uZZZZ"`, ErrBadUnicode},
		{`"unterminated`, ErrTruncated},
		{`[1,2`, ErrTruncated},
		{`{"a":`, ErrTruncated},
		{``, ErrTruncated},
		{`1e309`, ErrBadNumber}, // no float64; Valid rejects it too
		{`x`, ErrBadNumber},
	}
	for _, c := range cases {
		_, err := SkipValueStrict([]byte(c.in), 0)
		if !errors.Is(err, c.want) {
			t.Errorf("SkipValueStrict(%q) = %v, want %v", c.in, err, c.want)
		}
	}
	// The leniencies are deliberate and shared with the decoder's own readers;
	// see pkg/json.Valid's documentation. Pinned here so a "tightening" of this
	// walk cannot silently make a lax field stricter than the plain one.
	for _, in := range []string{`01`, `+1`, `.5`, `5.`, "\"\t\"", "{\"a\":\x001}"} {
		if _, err := SkipValueStrict([]byte(in), 0); err != nil {
			t.Errorf("SkipValueStrict(%q) = %v, want nil (this library accepts it)", in, err)
		}
	}
}

// TestSkipValueStrictDepthBound: the walk is iterative, so deep nesting costs
// bits in a fixed local rather than goroutine stack — but it still refuses past
// MaxDepth, as encoding/json does, instead of accepting arbitrarily deep input.
func TestSkipValueStrictDepthBound(t *testing.T) {
	atLimit := []byte(strings.Repeat("[", MaxDepth) + strings.Repeat("]", MaxDepth))
	if end, err := SkipValueStrict(atLimit, 0); err != nil || end != len(atLimit) {
		t.Errorf("at MaxDepth: end=%d err=%v, want %d and nil", end, err, len(atLimit))
	}
	past := []byte(strings.Repeat("[", MaxDepth+1) + strings.Repeat("]", MaxDepth+1))
	if _, err := SkipValueStrict(past, 0); !errors.Is(err, ErrMaxDepth) {
		t.Errorf("past MaxDepth: err=%v, want ErrMaxDepth", err)
	}
}

// TestSkipValueStrictTruncationSafe: every prefix of a valid document must
// return an error or a bounded end, never panic — the walk indexes data[i] in
// several places and a caller may hand it a cut-off buffer.
func TestSkipValueStrictTruncationSafe(t *testing.T) {
	doc := []byte(`{"a":[1,{"b":"éx","c":[true,null]}],"d":-1.5e3,"e":"\\"}`)
	for n := 0; n <= len(doc); n++ {
		end, err := SkipValueStrict(doc[:n], 0)
		if end < 0 || end > n {
			t.Fatalf("prefix %d: end=%d out of [0,%d]", n, end, n)
		}
		if n < len(doc) && err == nil {
			t.Errorf("prefix %d (%q) accepted as a complete value", n, doc[:n])
		}
	}
}
