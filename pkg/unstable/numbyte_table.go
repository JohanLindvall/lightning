//go:build !amd64

package unstable

// isNumberByte reports whether c continues a number token, as one load from a
// table. The alternative — the six comparisons the accept set spells out — is
// what amd64 uses and what this was until it was measured here; see
// numbyte_cmp.go for that side of the trade.
//
// The table wins on a core whose limit is instruction issue rather than
// dispatch width, and the measurement says exactly where: the cost per DIGIT
// is the same either way (eight instructions, by the slope through token
// lengths 1, 3 and 10), and the saving is a flat 17 instructions per token on
// the byte that ENDS it — every one of the six comparisons has to fail before
// a non-number byte can be rejected, where the table rejects it with the load
// it was going to do anyway. So the shorter the token the larger the win,
// which is the opposite of what a SWAR fold offers and the reason the two
// belong to different cores. The dependent load looks like it should hurt and
// does not: the loop's branch is predicted, so each iteration's load of
// data[i] issues without waiting for the previous iteration's table load, and
// only the branch waits.
// Measured on a Neoverse N2 (2026-09-07) over a 200-element array walk, per
// token shape and against the comparisons, cycles / instructions: one digit
// -18% / -26%, three digits -9% / -21%, ten digits -5% / -12%, 1.5 -12% /
// -24%, 1.5e-7 -24% / -30%, 0.000698752666567719 -8% / -9%; five- and
// nineteen-digit tokens were timed but not counted, at -8% and -5%. No shape
// is slower. On Zen 4 the same table
// measured +9% and +13% on ten- and three-digit tokens (CLAUDE.md's rejected
// list), which is why this is a build tag and not a rewrite.
func isNumberByte(c byte) bool { return numberByte[c] }

// numberByte is [0-9.eE+-], the accept set SkipNumber's doc comment describes.
// A literal rather than an init function, so the table is data.
var numberByte = [256]bool{
	'0': true, '1': true, '2': true, '3': true, '4': true,
	'5': true, '6': true, '7': true, '8': true, '9': true,
	'.': true, 'e': true, 'E': true, '+': true, '-': true,
}
