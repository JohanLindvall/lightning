//go:build !amd64

package unstable

// useFloatRun is false where there is no SIMD decimal-run kernel; the batch
// loop never calls parseFloatRun then.
var useFloatRun = false

// useFloatRunVBMI is false: there is no VBMI body here either.
var useFloatRunVBMI = false

func parseFloatRun(data []byte, i int, out []float64) (n, p, closed int) {
	return 0, i, 0
}

func parseFloatRunV(data []byte, i int, out []float64) (n, p, closed int) {
	return 0, i, 0
}

func parseFloatPoints(data []byte, i int, out []float64, n int) (np, p, closed int) {
	return 0, i, 0
}

// useValidRun and useValidRun512 are false: SkipValueStrict checks every
// number itself.
var useValidRun, useValidRun512 = false, false

func validNumberRun(data []byte, i int) (p, closed int) {
	return i, 0
}

func validPointsRun512(data []byte, i int) (p, closed int) {
	return i, 0
}
