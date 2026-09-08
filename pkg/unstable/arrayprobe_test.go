package unstable

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// TestSkipValueArrayProbeMatchesScalar pins SkipValue's array arm to the scalar
// balancer it chooses between. The arm is a heuristic — it decides from the
// first sixteen bytes after the opening bracket whether the block scan or
// skipArray settles the value faster — and a heuristic is only allowed to cost
// speed, so on every WELL-FORMED array the two must return the same end.
//
// The shapes are chosen around the probe's boundaries rather than to look
// representative: a ']' or a '"' at every offset the two SWAR words cover and
// one past them, a first element long enough that neither appears, and the
// arrays that reach the probe at all (a scalar first element — an array opening
// with '{', '[' or '"' is decided before it).
func TestSkipValueArrayProbeMatchesScalar(t *testing.T) {
	var docs []string
	for pad := 0; pad < 24; pad++ {
		digits := strings.Repeat("7", pad+1)
		docs = append(docs,
			"["+digits+"]",
			"["+digits+`,"s"]`,
			"["+digits+`,{"k":1}]`,
			"["+digits+",[1,2]]",
			"["+digits+",true,null,1.5e3]",
			"["+digits+`,"a\"b",2]`,
			"[ "+digits+" , \"s\" ]",
		)
	}
	docs = append(docs,
		"[]", "[ ]", "[0]", "[0,1,2,3,4,5,6,7,8,9]",
		"["+strings.Repeat("1,", 600)+"1]",
		`[1,`+strings.Repeat(`"x",`, 200)+`2]`,
		`[1,{"a":[1,2,{"b":"c"}]},"d"]`,
	)
	r := rand.New(rand.NewSource(3))
	for n := 0; n < 3000; n++ {
		var b strings.Builder
		b.WriteByte('[')
		for k, m := 0, 1+r.Intn(6); k < m; k++ {
			if k > 0 {
				b.WriteByte(',')
			}
			switch r.Intn(5) {
			case 0:
				fmt.Fprintf(&b, "%d", r.Int63())
			case 1:
				fmt.Fprintf(&b, `"%s"`, strings.Repeat("v", r.Intn(20)))
			case 2:
				b.WriteString(`{"k":1,"s":"t"}`)
			case 3:
				b.WriteString("[1,2,3]")
			default:
				b.WriteString([]string{"true", "false", "null", "-1.5e-3"}[r.Intn(4)])
			}
		}
		b.WriteByte(']')
		// A leading scalar is what reaches the probe; the other openings are
		// answered by the cheap first switch and are covered above.
		s := b.String()
		if len(s) > 1 && (s[1] == '{' || s[1] == '[' || s[1] == '"') {
			s = "[0," + s[1:]
		}
		docs = append(docs, s)
	}
	// Trailing slack matters: the probe needs sixteen bytes past the first
	// element and the block scan needs sixty-four, so the same array answers
	// through different paths depending on what follows it in the document.
	for _, tail := range []string{"", " ", strings.Repeat(" ", 17), strings.Repeat(" ", 80)} {
		for _, d := range docs {
			data := []byte(d + tail)
			want, wantErr := skipArray(data, 0)
			got, gotErr := SkipValue(data, 0)
			if got != want || (gotErr == nil) != (wantErr == nil) {
				t.Fatalf("SkipValue(%q) = (%d, %v), skipArray = (%d, %v)", data, got, gotErr, want, wantErr)
			}
		}
	}
}
