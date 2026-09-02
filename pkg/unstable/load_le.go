//go:build amd64 || arm64

package unstable

import "unsafe"

// load64 returns the eight bytes at data[i:i+8] as a little-endian word with
// NO bounds check: the caller has already proven i+8 <= len(data) (scanFloat
// checks its whole 48-byte window once, up front). A checked data[i:i+8]
// costs five instructions of slice arithmetic and compares per load, and the
// float fast path issues four of them per number. Only on the little-endian
// architectures that permit unaligned word loads; elsewhere load_other.go
// spells the same read through encoding/binary.
func load64(data []byte, i int) uint64 {
	return *(*uint64)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(data)), i))
}
