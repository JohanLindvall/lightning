package unstable

import (
	"bytes"
	"fmt"
	"testing"
)

// Small containers are the shape every walker over an array of records skips,
// and the shape BenchmarkSkipContainer's 300-member documents do not measure:
// there the block loop's prologue is amortized over dozens of blocks, here it
// is the whole cost. Each document is padded past the value so the 64-byte
// block path is reachable, which is what a real walk sees.
func smallShapes() []struct {
	name string
	data []byte
} {
	pad := bytes.Repeat([]byte(" "), 128)
	mk := func(s string) []byte { return append([]byte(s), pad...) }
	return []struct {
		name string
		data []byte
	}{
		{"pair", mk(`[1788087600,"0.0"]`)},                                         // 18B, a Prometheus point
		{"record", mk(`{"name":"record number 007","value":259,"active":true}`)},   // 52B
		{"tiny", mk(`{"a":1}`)},                                                    // 7B
		{"twoBlock", mk(fmt.Sprintf(`{"k":"%s"}`, bytes.Repeat([]byte("x"), 80)))}, // 90B
	}
}

func BenchmarkSkipSmall(b *testing.B) {
	for _, s := range smallShapes() {
		s := s
		b.Run(s.name, func(b *testing.B) {
			b.SetBytes(int64(len(s.data) - 128))
			for i := 0; i < b.N; i++ {
				e, err := SkipValue(s.data, 0)
				if err != nil {
					b.Fatal(err)
				}
				skipSmallSink += e
			}
		})
	}
}

// BenchmarkSkipSmallScalar is the same shapes through the scalar bracket
// balance (skipObject/skipArray), the path SkipValue's probe declines to use
// for them, so the two can be compared at the same document.
func BenchmarkSkipSmallScalar(b *testing.B) {
	for _, s := range smallShapes() {
		s := s
		b.Run(s.name, func(b *testing.B) {
			b.SetBytes(int64(len(s.data) - 128))
			for i := 0; i < b.N; i++ {
				var e int
				var err error
				if s.data[0] == '{' {
					e, err = skipObject(s.data, 0)
				} else {
					e, err = skipArray(s.data, 0)
				}
				if err != nil {
					b.Fatal(err)
				}
				skipSmallSink += e
			}
		})
	}
}

var skipSmallSink int
