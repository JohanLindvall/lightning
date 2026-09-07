package json

import "github.com/JohanLindvall/lightning/pkg/unstable"

// ParseInt parses the JSON integer in b as an int64. It is the integer
// counterpart of [ParseFloat] and the reader a caller wants for the value a
// [Get] or an [ObjectEach] hands back when that value is an integer, or for a
// quoted 64-bit integer once its quotes are stripped — ClickHouse, for one,
// quotes every Int64/UInt64 so it survives JavaScript's 53-bit numbers, and
// the alternative to this function is allocating a string for
// strconv.ParseInt on every such cell.
//
// b must be exactly one integer token with no surrounding whitespace: an
// optional sign followed by decimal digits. It accepts the sign forms the rest
// of this library accepts — a leading '+' as well as '-' — and leading zeros,
// like [ParseFloat]. A fraction or an exponent (1.0, 1e3), an empty input,
// stray bytes, and a magnitude outside int64 are all ErrBadNumber: this is a
// parser for an integer, not a float that happens to be whole. Nothing is
// retained or copied.
func ParseInt(b []byte) (int64, error) { return unstable.ParseInt(b) }

// ParseUint parses the JSON integer in b as a uint64, with [ParseInt]'s
// grammar and rules: a leading '+' is accepted (and only a '+' — a '-' is
// ErrBadNumber, as is any magnitude past uint64).
func ParseUint(b []byte) (uint64, error) { return unstable.ParseUint(b) }
