//go:build amd64

package unstable

import "golang.org/x/sys/cpu"

// useIntRun gates parseIntRunAVX2, the SIMD kernel for arrays of short
// integers: AVX2 for its 32-byte classification and BMI2 for SHRX/SHLX (see
// the assembly). Haswell/Excavator and later have both.
var useIntRun = cpu.X86.HasAVX2 && cpu.X86.HasBMI2

// parseIntRunAVX2 parses a run of comma-separated integers of one to eight
// digits from data[i:] into out, one int64 per element, and returns how many
// it wrote, where the scalar loop should resume, and whether the run ended at
// the array's ']' (p is then that byte's index, the element before it having
// been written). See intrun_amd64.s for the contract.
//
//go:noescape
func parseIntRunAVX2(data []byte, i int, out []int64) (n, p, closed int)

func parseIntRun(data []byte, i int, out []int64) (n, p, closed int) {
	return parseIntRunAVX2(data, i, out)
}

// intRunMinSlots is the spare capacity below which the batch loops do not
// call parseIntRun.
const intRunMinSlots = 1
