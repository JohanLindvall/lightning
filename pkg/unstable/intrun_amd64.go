//go:build amd64

package unstable

import "golang.org/x/sys/cpu"

// useIntRun gates parseIntRunSSE, the SIMD kernel for arrays of short
// integers: AVX for its VEX.128 encodings and BMI2 for SHRX (see the
// assembly for why the shift is SHRX). Haswell/Excavator and later have both.
var useIntRun = cpu.X86.HasAVX && cpu.X86.HasBMI2

// parseIntRunSSE parses a run of comma-separated integers of one to eight
// digits from data[i:] into out, one int64 per element, and returns how many
// it wrote, where the scalar loop should resume, and whether the run ended at
// the array's ']' (p is then that byte's index, the element before it having
// been written). See intrun_amd64.s for the contract.
//
//go:noescape
func parseIntRunSSE(data []byte, i int, out []int64) (n, p, closed int)

func parseIntRun(data []byte, i int, out []int64) (n, p, closed int) {
	return parseIntRunSSE(data, i, out)
}

// intRunShuffle holds the PSHUFB controls parseIntRunSSE folds a number with:
// entry s*8+L (unique for L in 1..8, the only lengths the kernel folds) for an
// L-digit number starting at block offset s. The eight bytes right-align the
// digits in the low lane with zeros in front (0x80 in a control byte zeroes
// the lane), so the fixed-weight folds read them as an eight-digit number with
// leading zeros; the kernel loads the entry with VMOVQ, so the upper lane's
// control is zero and its garbage never reaches the value (see the assembly).
var intRunShuffle [128][8]byte

func init() {
	for s := 0; s < 16; s++ {
		for l := 1; l <= 8 && s+l <= 16; l++ {
			m := &intRunShuffle[s*8+l]
			for t := range m {
				m[t] = 0x80
			}
			for t := 0; t < l; t++ {
				m[8-l+t] = byte(s + t)
			}
		}
	}
}
