//go:build amd64

package unstable

import "golang.org/x/sys/cpu"

// useCountAVX2 gates countBeforeClose's vector body; BMI2 supplies the SHRX
// and BZHI it takes the tail and the found block's count with.
var useCountAVX2 = useAVX2 && cpu.X86.HasBMI2 && cpu.X86.HasPOPCNT

// countKernel is the presize scan in count_amd64.s: the first ']' in data[i:]
// and the c bytes before it, or with hint the element count CountArrayScalars
// wants (see countBeforeClose and CountArrayScalars in count.go).
//
//go:noescape
func countKernel(data []byte, i int, c byte, hint bool) (rb, n int)
