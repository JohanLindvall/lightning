package unstable

import (
	"math/rand"
	"testing"
)

// skipWSRunOracle is the byte-at-a-time definition SkipWSRun must match exactly.
func skipWSRunOracle(data []byte, i int) int {
	for i < len(data) && data[i] <= ' ' {
		i++
	}
	return i
}

// TestSkipWSRunMatchesOracle is the guard on the all-spaces equality fast path: the
// word-at-a-time SWAR (and the one-compare shortcut in front of it) must agree with
// the byte loop for every input, not merely for well-formed indentation. The
// shortcut is only a *sufficient* test for eight spaces, so a divergence here would
// mean the exact classify behind it is being skipped wrongly.
func TestSkipWSRunMatchesOracle(t *testing.T) {
	check := func(data []byte, i int) {
		t.Helper()
		if got, want := SkipWSRun(data, i), skipWSRunOracle(data, i); got != want {
			t.Fatalf("SkipWSRun(%q, %d) = %d, oracle = %d", data, i, got, want)
		}
	}

	// Exhaustive: every byte value, at every offset of a spaces-filled buffer, for
	// every start offset — this covers each lane of the word and the boundary
	// between the word loop and the scalar tail.
	for b := 0; b < 256; b++ {
		for pos := 0; pos < 24; pos++ {
			buf := make([]byte, 24)
			for j := range buf {
				buf[j] = ' '
			}
			buf[pos] = byte(b)
			for start := 0; start <= 9; start++ {
				check(buf, start)
			}
		}
	}

	// Every whitespace byte as filler, so the fast path is exercised both when it
	// hits (spaces) and when it cannot (tabs, newlines, NUL).
	for _, fill := range []byte{' ', '\t', '\n', '\r', 0x00, 0x1f} {
		for n := 0; n <= 40; n++ {
			buf := make([]byte, n)
			for j := range buf {
				buf[j] = fill
			}
			for start := 0; start <= n; start++ {
				check(buf, start)
			}
			// ...and terminated by a structural byte.
			check(append(buf, '{'), 0)
			check(append(buf, 0x80), 0) // >= 0x80: must NOT count as whitespace
		}
	}

	// Random mixtures of whitespace, high bytes and structural bytes.
	alphabet := []byte{' ', ' ', ' ', ' ', '\t', '\n', '\r', 0x00, 0x1f, 0x20, 0x21,
		'{', '}', '"', ',', 0x7f, 0x80, 0xff, 'a'}
	rng := rand.New(rand.NewSource(7))
	for n := 0; n < 4000; n++ {
		buf := make([]byte, rng.Intn(48))
		for j := range buf {
			buf[j] = alphabet[rng.Intn(len(alphabet))]
		}
		for start := 0; start <= len(buf); start++ {
			check(buf, start)
		}
	}
}

// FuzzSkipWSRunMatchesOracle keeps the same contract under the native fuzzer.
func FuzzSkipWSRunMatchesOracle(f *testing.F) {
	f.Add([]byte("        {"), 0)
	f.Add([]byte("\n\t\t\t\t\t\t\t\t\t\t\"a\""), 1)
	f.Add([]byte("                        "), 2)
	f.Fuzz(func(t *testing.T, data []byte, i int) {
		if i < 0 || i > len(data) {
			t.Skip()
		}
		if got, want := SkipWSRun(data, i), skipWSRunOracle(data, i); got != want {
			t.Fatalf("SkipWSRun(%q, %d) = %d, oracle = %d", data, i, got, want)
		}
	})
}
