package unstable

import (
	"strconv"
	"testing"
)

// BenchmarkIndexStructural measures the structural scanner alone over a run of
// digits and commas with the only structural byte at the end — the shape a
// skipped scalar array presents (bench/skip-heavy is one such scan of 20 KB) —
// at lengths either side of the prescan, the 32- and 64-byte steps and the
// tails.
func BenchmarkIndexStructural(b *testing.B) {
	for _, n := range []int{8, 24, 40, 72, 136, 520, 4104, 20000} {
		buf := make([]byte, n+1)
		for k := range buf {
			buf[k] = "0123456789,"[k%11]
		}
		buf[n] = ']'
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			b.SetBytes(int64(n))
			for i := 0; i < b.N; i++ {
				if indexStructural(buf) != n {
					b.Fatal("wrong")
				}
			}
		})
	}
}
