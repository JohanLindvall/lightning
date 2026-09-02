package unstable

import "math/bits"

// digitRun folds the run of decimal digits starting at data[i] and returns its
// value modulo 2^64 and the index just past the run. The caller has already
// established that data[i] is a digit. Eight bytes are taken per word: the SWAR
// mask flags every non-digit lane and its lowest flagged lane is exact, so a
// full word of digits is folded whole and the final partial word is folded
// after shifting its digits into the top lanes, the vacated low lanes reading
// as leading zeros. Every product and sum is modulo 2^64, so the result is bit
// for bit the n*10+d chain the scalar loop computes, wrap included — and the
// same holds after truncation to any narrower integer kind, since reduction
// modulo 2^32 (or 2^16, 2^8) commutes with + and ×.
func digitRun(data []byte, i int) (uint64, int) {
	var n uint64
	for uint(i)+8 <= uint(len(data)) {
		d := load64(data, i) ^ swarZero // in bounds: the loop condition
		m := ((d + swarSix) | d) & swarNib
		if m == 0 {
			n = n*100000000 + parse8Digits(d)
			i += 8
			continue
		}
		k := bits.TrailingZeros64(m) >> 3
		// k can be 0 here: the word after a full one may start with the byte
		// that ends the run. The shift is then 64, which Go defines as 0, so the
		// fold adds nothing; do NOT mask the shift count as scanFloat does (it
		// excludes k == 0 first), that folds eight lanes of whatever follows.
		return n*pow10u64[k] + parse8Digits(d<<(uint(8-k)*8)), i + k
	}
	for uint(i) < uint(len(data)) {
		d := data[i] - '0'
		if d > 9 {
			break
		}
		n = n*10 + uint64(d)
		i++
	}
	return n, i
}
