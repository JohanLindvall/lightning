package json

import (
	stdjson "encoding/json"
	"strings"
	"testing"
)

// TestStringFindsEveryEscape is the coverage proof for the word tests String
// uses in place of a bytes.IndexByte call. They read a fixed pair (or two
// pairs) of eight- and four-byte windows at the token's ends, and the claim is
// that those windows cover every byte of a token up to 32 long. A window that
// fell one byte short would not error — it would return the token's bytes
// VERBATIM, escapes and all, which is silent corruption — so this walks a
// backslash escape across every position of every length either side of every
// boundary (4, 8, 16, 18, 19, 32), and holds the result to encoding/json.
func TestStringFindsEveryEscape(t *testing.T) {
	for n := 0; n <= 40; n++ {
		filler := strings.Repeat("a", n)
		// No escape anywhere: the fast path must accept and alias.
		checkString(t, `"`+filler+`"`)
		// One two-byte escape at every position it fits.
		for p := 0; p+2 <= n; p++ {
			for _, esc := range []string{`\n`, `\\`, `\"`, `\/`, `\t`} {
				checkString(t, `"`+filler[:p]+esc+filler[p+2:]+`"`)
			}
		}
		// A six-byte \uXXXX at every position it fits.
		for p := 0; p+6 <= n; p++ {
			checkString(t, `"`+filler[:p]+`é`+filler[p+6:]+`"`)
		}
	}
}

// checkString holds String's answer for one token to encoding/json's.
func checkString(t *testing.T, tok string) {
	t.Helper()
	var want string
	if err := stdjson.Unmarshal([]byte(tok), &want); err != nil {
		t.Fatalf("%s: the token is not one encoding/json accepts: %v", tok, err)
	}
	// A slice whose capacity stops at the token, so a window reading past it
	// would run past the allocation.
	raw := make([]byte, len(tok))
	copy(raw, tok)
	got, err := String(raw)
	if err != nil {
		t.Fatalf("String(%s) = %v", tok, err)
	}
	if got != want {
		t.Fatalf("String(%s) = %q, want %q", tok, got, want)
	}
}

// TestStringAliasesWhereItPromises pins the aliasing contract across the fast
// path: an escape-free token must still yield a string that shares the input's
// bytes, which is what makes the word tests worth having.
func TestStringAliasesWhereItPromises(t *testing.T) {
	for _, body := range []string{"", "a", "abc", "abcdefg", "abcdefgh", "0123456789abcdefg",
		strings.Repeat("x", 30), strings.Repeat("x", 31), strings.Repeat("x", 100)} {
		raw := []byte(`"` + body + `"`)
		got, err := String(raw)
		if err != nil {
			t.Fatal(err)
		}
		if got != body {
			t.Fatalf("String(%q) = %q", raw, got)
		}
		if len(body) > 0 {
			raw[1] = 'Z'
			if got[0] != 'Z' {
				t.Fatalf("String(%q) did not alias its input", raw)
			}
			raw[1] = body[0]
		}
	}
}

// TestUnescapeStringCopyFindsEveryEscape is TestStringFindsEveryEscape for the
// copying unescaper, which reads the same fixed windows over a body with no
// quotes to pad it. Held to UnescapeString, whose answer is the contract.
func TestUnescapeStringCopyFindsEveryEscape(t *testing.T) {
	for n := 0; n <= 40; n++ {
		filler := strings.Repeat("a", n)
		bodies := []string{filler}
		for p := 0; p+2 <= n; p++ {
			for _, esc := range []string{`\n`, `\\`, `\"`, `\/`} {
				bodies = append(bodies, filler[:p]+esc+filler[p+2:])
			}
		}
		for p := 0; p+6 <= n; p++ {
			bodies = append(bodies, filler[:p]+`é`+filler[p+6:])
		}
		for _, body := range bodies {
			in := make([]byte, len(body)) // capacity stops at the body
			copy(in, body)
			want, wantErr := UnescapeString(in)
			got, gotErr := UnescapeStringCopy(in)
			if (gotErr != nil) != (wantErr != nil) || got != want {
				t.Fatalf("UnescapeStringCopy(%q) = %q,%v; UnescapeString = %q,%v", body, got, gotErr, want, wantErr)
			}
			// And it must never be a window onto in.
			if len(got) > 0 && &[]byte(got)[0] == &in[0] {
				t.Fatalf("UnescapeStringCopy(%q) aliased its input", body)
			}
		}
	}
}
