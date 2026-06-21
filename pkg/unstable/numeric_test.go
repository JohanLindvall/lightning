package unstable

import "testing"

func TestFastFloat(t *testing.T) {
	fastFloat := func(b []byte) (float64, bool) {
		f, end, fast, ok := scanFloat(b, 0)
		return f, ok && fast && end == len(b)
	}
	okCases := map[string]float64{
		"0":      0,
		"1":      1,
		"-1":     -1,
		"+2":     2,
		"3.5":    3.5,
		"0.001":  0.001,
		"12e2":   1200,
		"12E2":   1200,
		"1.5e-1": 0.15,
		// Outside the Clinger exponent range [-22, 22] but an exact (<= 19 digit)
		// mantissa, so scanFloat resolves it on the Eisel-Lemire fast path.
		"1e-23": 1e-23,
	}
	for in, want := range okCases {
		got, ok := fastFloat([]byte(in))
		if !ok {
			t.Errorf("fastFloat(%q) ok=false, want true", in)
			continue
		}
		if got != want {
			t.Errorf("fastFloat(%q) = %v, want %v", in, got, want)
		}
	}

	declined := []string{
		"",                     // empty
		"1e23",                 // exactly halfway: Eisel-Lemire declines, defers to strconv
		"1e",                   // missing exp digits
		"1e+",                  // missing exp digits after sign
		"12345678901234567890", // 20 digits, > 19
		"1e99999",              // huge exponent
		".",                    // no digits
		"1x",                   // trailing junk -> not fully consumed
	}
	for _, in := range declined {
		if _, ok := fastFloat([]byte(in)); ok {
			t.Errorf("fastFloat(%q) ok=true, want false (should defer to strconv)", in)
		}
	}
}
