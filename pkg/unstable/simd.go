package unstable

// This file holds the exported, inlinable wrappers over the per-architecture
// SIMD scanners (defined in simd_amd64.go / simd_arm64.go / simd_noasm.go with
// the generic fallback in simd_other.go). They let generated decoders reach the
// scanner inline at the call site rather than through a Read* call.

// IndexCloseOrEscape returns the index of the first '"' or '\\' byte in b, or
// len(b) if neither is present. It is exported (and inlinable) so generated
// decoders can write the object-key / string read inline at the call site — the
// no-escape fast path — instead of paying a ReadKey call; ReadKey stays the
// escape/error fallback.
func IndexCloseOrEscape(b []byte) int { return indexCloseOrEscape(b) }
