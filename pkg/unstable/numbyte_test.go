package unstable

import "testing"

// numberByteCompare is isNumberByte's comparison spelling, the one numbyte_cmp.go
// compiles on amd64. Held against whichever form the build selected, so that an
// arm64 run — this repository's CI runs both — checks numbyte_table.go's table
// against the accept set SkipNumber's doc comment states. On amd64 the check is
// the comparisons against themselves and TestSkipNumberSpan carries the weight.
func numberByteCompare(c byte) bool {
	return (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-'
}

func TestIsNumberByteMatchesComparisons(t *testing.T) {
	for i := 0; i < 256; i++ {
		c := byte(i)
		if got, want := isNumberByte(c), numberByteCompare(c); got != want {
			t.Errorf("isNumberByte(%#02x) = %v, want %v", c, got, want)
		}
	}
}

// TestSkipNumberSpan pins the span and the error against the byte-at-a-time
// form SkipNumber had before the predicate moved behind isNumberByte — the
// maximal run of [0-9.eE+-], with an empty run rejected. A leading '-' used to
// be consumed by a step of its own in front of the loop, which the loop already
// accepted; the step is gone and nothing about the answer moved.
func TestSkipNumberSpan(t *testing.T) {
	oracle := func(data []byte, i int) (int, error) {
		start := i
		if uint(i) < uint(len(data)) && data[i] == '-' {
			i++
		}
		for uint(i) < uint(len(data)) {
			c := data[i]
			if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
				i++
				continue
			}
			break
		}
		if i == start {
			return i, ErrBadNumber
		}
		return i, nil
	}
	bodies := []string{"", "0", "7", "200", "-1", "+5", "1.5", "1e", "1e+308", "1.2.3", "--1", "-",
		"+", ".", "e", "E", "1788087600", "9223372036854775807", "0.000698752666567719",
		"12345678", "123456789", "1234567890123456789012345", "x", "null", "true"}
	tails := []string{"", ",", "]", "}", " ", "x", "\x00", ",1]", "e5", "00000000"}
	for _, b := range bodies {
		for _, tail := range tails {
			in := []byte(b + tail)
			for i := 0; i <= len(in); i++ {
				gotEnd, gotErr := SkipNumber(in, i)
				wantEnd, wantErr := oracle(in, i)
				if gotEnd != wantEnd || (gotErr == nil) != (wantErr == nil) {
					t.Fatalf("SkipNumber(%q, %d) = (%d, %v), want (%d, %v)",
						in, i, gotEnd, gotErr, wantEnd, wantErr)
				}
			}
		}
	}
}
