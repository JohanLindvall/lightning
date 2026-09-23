//go:build !amd64 && !arm64

package unstable

// floatRunBodies and validRunBodies are empty: there are no number kernels
// here, and the batch readers and SkipValueStrict never call the stubs.
func floatRunBodies() []kernelBody { return nil }

func validRunBodies() []kernelBody { return nil }

// restoreKernels has nothing to restore beyond the flags the tests switch
// off, which are false here to begin with.
func restoreKernels() {
	useFloatRun, useFloatRunLong = floatRunHost, floatRunLongHost
	useValidRun, useValidPoints = validRunHost, validPointsHost
}
