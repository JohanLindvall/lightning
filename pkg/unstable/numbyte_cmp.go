//go:build amd64

package unstable

// isNumberByte reports whether c continues a number token. On amd64 it is the
// comparisons rather than the table of numbyte_table.go: a table-driven byte
// loop measured +9% on a ten-digit token and +13% on a three-digit one on
// Zen 4, where the load's latency lands on the loop's exit branch and a
// compare's does not. See numbyte_table.go for the other side.
func isNumberByte(c byte) bool {
	return (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-'
}
