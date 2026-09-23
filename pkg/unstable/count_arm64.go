//go:build arm64

package unstable

// countKernel is the presize scan in count_arm64.s: the first ']' in data[i:]
// and the c bytes before it, or with hint the element count
// CountArrayScalars wants (see countBeforeClose and CountArrayScalars in
// count.go).
//
//go:noescape
func countKernel(data []byte, i int, c byte, hint bool) (rb, n int)
