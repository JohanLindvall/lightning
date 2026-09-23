package unstable

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

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

// refReadInt is the integer readers' contract written as the plain byte loop
// they were before the eight-digit step (amd64) and digitRun (elsewhere): a
// sign for the signed reader, a digit run folded as n*10+d modulo 2^64, and a
// fraction or exponent tail measured and dropped.
func refReadInt(data []byte, i int, signed bool) (uint64, int, error) {
	if i >= len(data) {
		return 0, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return 0, end, err
	}
	neg := false
	if signed && data[i] == '-' {
		neg = true
		i++
		if i >= len(data) {
			return 0, i, ErrBadNumber
		}
	}
	if data[i] < '0' || data[i] > '9' {
		return 0, i, ErrBadNumber
	}
	var n uint64
	for i < len(data) && data[i] >= '0' && data[i] <= '9' {
		n = n*10 + uint64(data[i]-'0')
		i++
	}
	if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
		for i < len(data) && strings.IndexByte("0123456789.eE+-", data[i]) >= 0 {
			i++
		}
	}
	if neg {
		n = -n
	}
	return n, i, nil
}

// TestIntReadersMatchByteLoop holds ReadInt64OrNull and ReadUint64OrNull to
// refReadInt over every digit-run length from 0 to 26 — either side of the
// eight-digit step (and of digitRun's word), past the 19 digits a uint64 holds
// — behind every terminator shape, with the byte seven past the run's start a
// digit and not (the step's pre-filter reads it, and a short number followed
// closely by more digits must still stop at its own end), and cut at every
// length, so a run that ends at the buffer and a word that would read past it
// are both reached.
func TestIntReadersMatchByteLoop(t *testing.T) {
	rng := rand.New(rand.NewSource(88))
	tails := []string{"", ",", "}", "]", " ", `"`, ".5,", "e7}", ".25e-3,", "x", ",12345678901234", ", 9", "123"}
	for l := 0; l <= 26; l++ {
		for _, tail := range tails {
			for _, sign := range []string{"", "-"} {
				for rep := 0; rep < 4; rep++ {
					var b strings.Builder
					b.WriteString(sign)
					for k := 0; k < l; k++ {
						c := byte('0' + rng.Intn(10))
						if rep == 0 {
							c = '9' // the largest value at each length: the wrap
						}
						b.WriteByte(c)
					}
					b.WriteString(tail)
					full := b.String()
					for cut := 0; cut <= len(full); cut++ {
						data := []byte(full[:cut])
						for _, signed := range []bool{true, false} {
							wv, we, werr := refReadInt(data, 0, signed)
							var gv uint64
							var ge int
							var gerr error
							if signed {
								var v int64
								v, ge, gerr = ReadInt64OrNull(data, 0)
								gv = uint64(v)
							} else {
								gv, ge, gerr = ReadUint64OrNull(data, 0)
							}
							if gv != wv || ge != we || !errors.Is(gerr, werr) || !errors.Is(werr, gerr) {
								t.Fatalf("signed=%v %q: (%d, %d, %v), want (%d, %d, %v)", signed, data, gv, ge, gerr, wv, we, werr)
							}
						}
					}
				}
			}
		}
	}
}

// BenchmarkReadIntShapes is ReadInt64OrNull over 200 integers of one digit
// count, each read starting where the one before it ended — the cursor chain a
// generated decoder has from one member to the next — so a reader whose
// cursor waits on its value (a counted word fold) pays here as it does in a
// decode. Short counts are the small integers of instruments and random, nine
// and ten citm's and golang_source's ids, thirteen a millisecond timestamp,
// eighteen twitter's ids.
func BenchmarkReadIntShapes(b *testing.B) {
	for _, nd := range []int{1, 3, 6, 9, 10, 13, 18} {
		rng := rand.New(rand.NewSource(int64(nd)))
		var sb strings.Builder
		for k := 0; k < 200; k++ {
			sb.WriteByte(byte('1' + rng.Intn(9)))
			for d := 1; d < nd; d++ {
				sb.WriteByte(byte('0' + rng.Intn(10)))
			}
			sb.WriteString(`,"key":`)
		}
		data := []byte(sb.String())
		b.Run(fmt.Sprintf("d%02d", nd), func(b *testing.B) {
			var sum int64
			for b.Loop() {
				for i, k := 0, 0; k < 200; k++ {
					v, end, err := ReadInt64OrNull(data, i)
					if err != nil {
						b.Fatal(err)
					}
					sum += v
					i = end + 7
				}
			}
			if sum == 42 {
				b.Log(sum)
			}
		})
	}
}
