package unstable

// This file holds the exported, inlinable wrappers over the per-architecture
// SIMD scanners (defined in simd_amd64.go / simd_arm64.go, with simd_scalar.go
// dispatching every other architecture to the generic implementations in
// simd_other.go). They let generated decoders reach the scanner inline at the
// call site rather than through a Read* call.

// IndexCloseOrEscape returns the index of the first '"' or '\\' byte in b, or
// len(b) if neither is present. It is exported (and inlinable) so generated
// decoders can write the object-key / string read inline at the call site — the
// no-escape fast path — instead of paying a ReadKey call; ReadKey stays the
// escape/error fallback.
func IndexCloseOrEscape(b []byte) int { return indexCloseOrEscape(b) }

// IndexEscape returns the index of the first byte that JSON string encoding must
// escape — a control byte < 0x20, '"' or '\\' — or len(b) if none. It is the
// scan behind EscapeString/EscapeStringInto: a clean run is copied out in bulk and
// only the escape byte at the returned index is expanded. SIMD (SSE2/AVX2) on
// amd64, SWAR elsewhere.
func IndexEscape(b []byte) int { return indexEscape(b) }

// IndexEscapeNonASCII returns the index of the first byte that is either one
// JSON string encoding must escape (a control byte < 0x20, '"' or '\\') or a
// non-ASCII byte (>= 0x80), or len(b) if none. It is the scan behind
// EscapeStringInto's UTF-8 handling: the walk runs on it until the first
// non-ASCII byte, where one utf8.Valid call decides between the plain escape
// path and U+FFFD substitution — so the widened predicate costs the pure-ASCII
// common case only an extra OR per SIMD block over IndexEscape.
func IndexEscapeNonASCII(b []byte) int { return indexEscapeNonASCII(b) }
