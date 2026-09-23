//go:build amd64

package unstable

// countBodies lists the countKernel bodies an amd64 host has: the byte loop
// always, the AVX2 step with useCountAVX2.
func countBodies() []countBody {
	b := []countBody{{"bytes", func() func() {
		saved := useCountAVX2
		useCountAVX2 = false
		return func() { useCountAVX2 = saved }
	}}}
	if useCountAVX2 {
		b = append(b, countBody{"avx2", func() func() { return func() {} }})
	}
	return b
}
