//go:build !amd64 && !arm64

package unstable

import "encoding/binary"

// load64 is the checked, byte-order-independent form of load_le.go's word
// read, for architectures where an unaligned unsafe load is not guaranteed.
func load64(data []byte, i int) uint64 {
	return binary.LittleEndian.Uint64(data[i : i+8])
}

// load32 is the checked form of load_le.go's four-byte read.
func load32(data []byte, i int) uint32 {
	return binary.LittleEndian.Uint32(data[i : i+4])
}
