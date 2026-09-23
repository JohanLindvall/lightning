//go:build !amd64

package unstable

import "runtime"

// countBodies lists the one countKernel body a host other than amd64 has:
// the NEON pass on arm64, the runtime scans elsewhere.
func countBodies() []countBody {
	return []countBody{{runtime.GOARCH, func() func() { return func() {} }}}
}
