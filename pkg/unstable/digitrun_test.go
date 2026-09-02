package unstable

import "testing"

// TestDigitRunWordBoundary pins the run lengths the batch parity tests never
// reach: a run that ends exactly on an eight-byte word boundary with more input
// behind it, so the word after it starts with the terminator (a zero-length run
// in that word), plus runs that end at the buffer and one past the 19-digit
// int64 range (the value wraps modulo 2^64, as the n*10+d chain does).
func TestDigitRunWordBoundary(t *testing.T) {
	inputs := []string{
		"12345678,xxxxxxxxxxxxxxxx", "1234567890123456,xxxxxxxxxx", "12345678", "1", "1,",
		"123456789012345678901234", "0000000000000000000000001x", "99999999999999999999,",
		"12345678]", "123456781234567812345678123456781",
	}
	for _, in := range inputs {
		want, i := uint64(0), 0
		for i < len(in) && in[i] >= '0' && in[i] <= '9' {
			want = want*10 + uint64(in[i]-'0')
			i++
		}
		if got, end := digitRun([]byte(in), 0); got != want || end != i {
			t.Errorf("digitRun(%q) = (%d, %d), want (%d, %d)", in, got, end, want, i)
		}
		if v, end, err := ReadUint64OrNull([]byte(in), 0); err != nil || v != want || end != i {
			t.Errorf("ReadUint64OrNull(%q) = (%d, %d, %v), want (%d, %d, nil)", in, v, end, err, want, i)
		}
		if v, end, err := ReadInt64OrNull([]byte(in), 0); err != nil || v != int64(want) || end != i {
			t.Errorf("ReadInt64OrNull(%q) = (%d, %d, %v), want (%d, %d, nil)", in, v, end, err, int64(want), i)
		}
		var s []int64
		if end, err := DecodeIntSlice(&s, []byte("["+in[:i]+"]"), 0); err != nil || len(s) != 1 || s[0] != int64(want) || end != i+2 {
			t.Errorf("DecodeIntSlice([%q]) = (%v, %d, %v), want ([%d], %d, nil)", in[:i], s, end, err, int64(want), i+2)
		}
	}
}
