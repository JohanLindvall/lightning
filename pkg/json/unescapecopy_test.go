package json

import (
	"errors"
	"testing"
	"unsafe"
)

// UnescapeStringCopy decodes exactly as UnescapeString does — same value,
// same errors, invalid UTF-8 passed through — on every body shape.
func TestUnescapeStringCopyMatchesUnescapeString(t *testing.T) {
	bodies := []string{
		``, `plain`, `a b`, `line\nbreak`, `quote\"inside`, `back\\slash`, `tab\there`,
		`é`, `😀`, `\ud83d`, `☃ raw`, "\xff\xfe raw invalid utf-8",
		`\/`, `\b\f\r`, `mixed \n text A end`, `trailing\\`,
		// Bad escapes: both must refuse, with the same sentinel.
		`\x`, `\u12`, `\`, `\uZZZZ`,
	}
	for _, s := range bodies {
		in := []byte(s)
		want, wantErr := UnescapeString(in)
		got, err := UnescapeStringCopy(in)
		switch {
		case (err == nil) != (wantErr == nil):
			t.Errorf("UnescapeStringCopy(%q): err %v, UnescapeString %v", s, err, wantErr)
		case err != nil && !errors.Is(err, wantErr):
			t.Errorf("UnescapeStringCopy(%q): err %v, UnescapeString %v", s, err, wantErr)
		case err == nil && got != want:
			t.Errorf("UnescapeStringCopy(%q) = %q, UnescapeString %q", s, got, want)
		}
	}
}

// The property the function exists for: the result never points into in,
// on the escape-free path (where UnescapeString aliases) or the escaped one,
// so overwriting the input afterwards changes nothing.
func TestUnescapeStringCopyNeverAliasesItsInput(t *testing.T) {
	for _, s := range []string{`pod-name`, `pod\tname`, `x`} {
		in := []byte(s)
		got, err := UnescapeStringCopy(in)
		if err != nil {
			t.Fatal(err)
		}
		if p := unsafe.StringData(got); len(got) > 0 && uintptr(unsafe.Pointer(p)) >= uintptr(unsafe.Pointer(&in[0])) && uintptr(unsafe.Pointer(p)) < uintptr(unsafe.Pointer(&in[0]))+uintptr(len(in)) {
			t.Fatalf("UnescapeStringCopy(%q) aliases its input", s)
		}
		before := got
		for i := range in {
			in[i] = 'X'
		}
		if got != before {
			t.Fatalf("UnescapeStringCopy(%q) changed when the input was overwritten: %q", s, got)
		}
	}
	// And the contrast that makes this function necessary, pinned so it is
	// a deliberate change on the day it moves: UnescapeString DOES alias an
	// escape-free body.
	in := []byte(`plain`)
	s, err := UnescapeString(in)
	if err != nil {
		t.Fatal(err)
	}
	if unsafe.StringData(s) != &in[0] {
		t.Fatal("UnescapeString no longer aliases an escape-free body; the Copy form's reason for existing has moved")
	}
}

func BenchmarkUnescapeStringCopy(b *testing.B) {
	in := []byte(`entity-service-20260903-1`)
	b.ReportAllocs()
	for range b.N {
		if _, err := UnescapeStringCopy(in); err != nil {
			b.Fatal(err)
		}
	}
}
