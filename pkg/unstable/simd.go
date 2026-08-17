package unstable

// The arm64 SVE2 scanners spell their instructions as WORD constants, because
// the Go assembler has no SVE mnemonics; internal/sveasm regenerates those
// constants from the mnemonics written in their comments. The directive lives
// here rather than in simd_arm64.go so that `go generate ./...` reaches it from
// any host — the tool prefers the aarch64 cross-assembler and so is not itself
// arm64-only, whereas a directive inside a `//go:build arm64` file would be
// skipped everywhere else. `make sveasm-check` is the gate; see the Makefile.
//
//go:generate go run ../../internal/sveasm -w simd_arm64.s

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

// IndexCloseOrEscapeAt is IndexCloseOrEscape starting the scan at i and
// returning an ABSOLUTE index into b (len(b) if there is no '"' or '\\' at or
// after i). It is what generated decoders and the readers here call, because
// expressing the start as IndexCloseOrEscape(b[i:]) makes the caller pay for a
// reslice — seven instructions on arm64: the len and cap subtractions and the
// negative-length clamp on the base pointer — once per object key and once per
// string value, where handing the offset to the scanner costs one argument word
// and nothing in the scan itself. Measured -20.7% on a key-read-shaped micro,
// and the absolute result is what every caller wanted anyway.
//
// An i past the end of b is not an error: the answer is len(b), the same as a
// scan that found nothing. Every implementation agrees on that, and the amd64
// one needs an explicit guard to (its byte count would otherwise go negative and
// walk off the buffer), which TestIndexCloseOrEscapeAtPastEnd pins.
func IndexCloseOrEscapeAt(b []byte, i int) int { return indexCloseOrEscapeAt(b, i) }

// IndexEscape returns the index of the first byte that JSON string encoding must
// escape — a control byte < 0x20, '"' or '\\' — or len(b) if none. It is the
// scan behind EscapeString/EscapeStringInto: a clean run is copied out in bulk and
// only the escape byte at the returned index is expanded. SIMD on amd64
// (SSE2/AVX2) and arm64 (NEON, or SVE2 where the core has it); SWAR elsewhere.
func IndexEscape(b []byte) int { return indexEscape(b) }

// IndexEscapeNonASCII returns the index of the first byte that is either one
// JSON string encoding must escape (a control byte < 0x20, '"' or '\\') or a
// non-ASCII byte (>= 0x80), or len(b) if none. It is the scan behind
// EscapeStringInto's UTF-8 handling: the walk runs on it until the first
// non-ASCII byte, where one utf8.Valid call decides between the plain escape
// path and U+FFFD substitution — so the widened predicate costs the pure-ASCII
// common case only an extra OR per SIMD block over IndexEscape.
func IndexEscapeNonASCII(b []byte) int { return indexEscapeNonASCII(b) }
