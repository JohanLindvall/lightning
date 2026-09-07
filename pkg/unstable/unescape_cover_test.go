package unstable

import (
	stdjson "encoding/json"
	"strings"
	"testing"
)

// TestUnescapeFindsEveryEscape is the coverage proof for the word tests
// UnescapeString and UnescapeStringInto take in place of a bytes.IndexByte
// call, the same claim TestStringFindsEveryEscape makes for pkg/json.String:
// the fixed windows at the body's two ends cover every byte of a body up to 32
// long. A window one byte short would not error — it would return the body
// VERBATIM, escapes and all — so this walks an escape across every position of
// every length either side of every boundary (4, 8, 16, 18, 19, 32) and holds
// the answer to encoding/json's for the same body between quotes.
func TestUnescapeFindsEveryEscape(t *testing.T) {
	check := func(body string) {
		t.Helper()
		var want string
		if err := stdjson.Unmarshal([]byte(`"`+body+`"`), &want); err != nil {
			t.Fatalf("%q: not a body encoding/json accepts: %v", body, err)
		}
		// Capacity stops at the body, so a window reading past it reads past
		// the allocation rather than into the rest of a larger buffer.
		in := make([]byte, len(body))
		copy(in, body)
		got, err := UnescapeString(in)
		if err != nil {
			t.Fatalf("UnescapeString(%q) = %v", body, err)
		}
		if got != want {
			t.Fatalf("UnescapeString(%q) = %q, want %q", body, got, want)
		}
		in2 := make([]byte, len(body))
		copy(in2, body)
		out := make([]byte, 0, len(body))
		got, err = UnescapeStringInto(in2, out)
		if err != nil {
			t.Fatalf("UnescapeStringInto(%q) = %v", body, err)
		}
		if got != want {
			t.Fatalf("UnescapeStringInto(%q) = %q, want %q", body, got, want)
		}
	}
	for n := 0; n <= 40; n++ {
		filler := strings.Repeat("a", n)
		check(filler)
		for p := 0; p+2 <= n; p++ {
			for _, esc := range []string{`\n`, `\\`, `\"`, `\/`, `\t`} {
				check(filler[:p] + esc + filler[p+2:])
			}
		}
		for p := 0; p+6 <= n; p++ {
			check(filler[:p] + `é` + filler[p+6:])
		}
	}
}
