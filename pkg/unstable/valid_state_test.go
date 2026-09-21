package unstable

import (
	"strconv"
	"strings"
	"testing"
)

func TestSkipValueStrictRestoresContainerKind(t *testing.T) {
	for _, depth := range []int{1, 2, 3, 63, 64, 65, 127, 128, 129, MaxDepth} {
		t.Run(strconv.Itoa(depth), func(t *testing.T) {
			var doc strings.Builder
			for d := 0; d < depth; d++ {
				if d%2 == 0 {
					doc.WriteString(`[{"empty":[]},`)
				} else {
					doc.WriteString(`{"empty":{},"child":`)
				}
			}
			doc.WriteString(`"leaf\n"`)
			for d := depth - 1; d >= 0; d-- {
				if d%2 == 0 {
					doc.WriteString(`,true]`)
				} else {
					doc.WriteString(`,"tail":false}`)
				}
			}
			in := []byte(doc.String())
			// The leading empty siblings add two levels before descent.
			if depth == MaxDepth {
				in = []byte(strings.Repeat(`[{"x":`, depth/2) + `0` + strings.Repeat(`},true]`, depth/2))
			}
			end, err := SkipValueStrict(in, 0)
			if err != nil || end != len(in) {
				t.Fatalf("end=%d err=%v, want end=%d", end, err, len(in))
			}
			// After unwinding every child, the root must still expect an array
			// closer. Reading the last child's kind would accept this instead.
			in[len(in)-1] = '}'
			if _, err := SkipValueStrict(in, 0); err != ErrInvalidJSON {
				t.Fatalf("mismatched root closer: %v", err)
			}
		})
	}
}

func TestSkipValueStrictStringErrorOffsets(t *testing.T) {
	for _, n := range []int{0, 1, 15, 16, 31, 32, 63, 64, 129} {
		prefix := `"` + strings.Repeat("x", n)
		for _, c := range []struct {
			tail string
			end  int
			err  error
		}{
			{`\q"`, n + 2, ErrBadEscape},
			{`\uZZZZ"`, n + 2, ErrBadUnicode},
			{`\u12`, n + 2, ErrTruncated},
			{`\`, n + 2, ErrTruncated},
			{`\n`, n + 3, ErrTruncated},
			{"", n + 1, ErrTruncated},
		} {
			// Both a string value and an object key take the inline scan.
			for _, lead := range []string{"", "{"} {
				in := []byte(lead + prefix + c.tail)
				end, err := SkipValueStrict(in, 0)
				if end != len(lead)+c.end || err != c.err {
					t.Errorf("%q: end=%d err=%v, want end=%d err=%v", in, end, err, len(lead)+c.end, c.err)
				}
			}
		}
	}
}
