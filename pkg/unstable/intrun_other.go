//go:build !amd64 && !arm64

package unstable

// useIntRun is false where there is no SIMD integer-run kernel; the batch
// loops never call parseIntRun then.
var useIntRun = false

func parseIntRun(data []byte, i int, out []int64) (n, p, closed int) {
	return 0, i, 0
}

// intRunMinSlots is unused where there is no kernel.
const intRunMinSlots = 1
