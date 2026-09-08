package unstable

// structuralMask returns a mask with the high bit set in every byte lane of w
// that holds a JSON structural byte — '"', '[', ']', '{' or '}'. Only the
// LOWEST set lane is guaranteed to be a real match: a borrow out of a matching
// lane can flag the lane above it, exactly as in the SWAR digit tests, and the
// borrow can only ever travel upward. Every caller takes TrailingZeros64, so
// that is the property they need.
//
// The four brackets cost two tests rather than four. '[' and '{' differ only in
// bit 0x20, and so do ']' and '}', so ORing that bit in folds each pair onto its
// upper member: w|0x20 == '{' holds for exactly '{' and '[', and w|0x20 == '}'
// for exactly '}' and ']'. That is EXACT, unlike the tempting single test that
// masks off every bit the four differ in — the smallest cube containing all four
// also contains 'Y', '_', 'y' and DEL, and a scan that stopped on those would be
// a heuristic rather than the scanner it replaces.
//
// Three has-byte tests at two instructions a byte beat the byte loop's five
// compares, and that loop was 36% of a walk over Prometheus-shaped
// [timestamp,"value"] pairs, whose ten-digit timestamps put eleven bytes between
// one structural byte and the next.
func structuralMask(w uint64) uint64 {
	const (
		lo    = 0x0101010101010101
		hi    = 0x8080808080808080
		bit5  = lo * 0x20
		quote = lo * '"'
		open  = lo * '{'
		close = lo * '}'
	)
	t := w | bit5
	a := t ^ open
	b := t ^ close
	c := w ^ quote
	return ((a-lo)&^a | (b-lo)&^b | (c-lo)&^c) & hi
}
