//go:build !amd64 && !arm64

package unstable

import "encoding/binary"

// load64 is the checked, byte-order-independent form of load_le.go's word
// read, for architectures where an unaligned unsafe load is not guaranteed.
func load64(data []byte, i int) uint64 {
	return binary.LittleEndian.Uint64(data[i : i+8])
}
