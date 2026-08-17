package unstable

import (
	"math/rand"
	"testing"
)

// TestTryParse4DigitsMatchesIs4Digits pins tryParse4Digits' accept set to
// is4Digits', which the SWAR folds used before it, and its value to
// parse4Digits'. The two forms are algebraically identical (see
// tryParse4Digits), and the point of this test is that the identity is not left
// to the argument: the byte set below is every value adjacent to the digit range
// where a wrong constant would land ('/' and ':' straddle it, 0x00/0xff and
// 0x80 exercise the high-bit and carry lanes), taken in every combination of the
// four lanes, plus random words.
func TestTryParse4DigitsMatchesIs4Digits(t *testing.T) {
	bytes := []byte{0x00, 0x01, 0x0a, 0x2f, '0', '1', '5', '8', '9', ':', ';', 0x40, 0x7f, 0x80, 0x8a, 0xd0, 0xfe, 0xff}
	for _, b0 := range bytes {
		for _, b1 := range bytes {
			for _, b2 := range bytes {
				for _, b3 := range bytes {
					w := uint32(b0) | uint32(b1)<<8 | uint32(b2)<<16 | uint32(b3)<<24
					v, ok := tryParse4Digits(w)
					if want := is4Digits(w); ok != want {
						t.Fatalf("tryParse4Digits(%08x) ok = %v, is4Digits = %v", w, ok, want)
					}
					if ok {
						if want := parse4Digits(w); v != want {
							t.Fatalf("tryParse4Digits(%08x) = %d, parse4Digits = %d", w, v, want)
						}
					}
				}
			}
		}
	}
	rng := rand.New(rand.NewSource(3))
	for iter := 0; iter < 2_000_000; iter++ {
		var w uint32
		switch iter % 3 {
		case 0:
			w = rng.Uint32()
		case 1: // mostly-digit words, so the accepting side is exercised too
			w = 0
			for k := 0; k < 4; k++ {
				c := uint32('0' + rng.Intn(10))
				if rng.Intn(16) == 0 {
					c = uint32(rng.Intn(256))
				}
				w |= c << (8 * k)
			}
		default:
			w = 0x30303030 + rng.Uint32()&0x0f0f0f0f
		}
		v, ok := tryParse4Digits(w)
		if want := is4Digits(w); ok != want {
			t.Fatalf("tryParse4Digits(%08x) ok = %v, is4Digits = %v", w, ok, want)
		}
		if ok && v != parse4Digits(w) {
			t.Fatalf("tryParse4Digits(%08x) = %d, parse4Digits = %d", w, v, parse4Digits(w))
		}
	}
}
