//go:build arm64

package unstable

import "golang.org/x/sys/cpu"

// useIntRun gates parseIntRunNEON, the arm64 twin of the amd64 SIMD kernel
// for arrays of short integers, on the dot-product instructions its fold
// uses (UDOT, ARMv8.2 DotProd — mandatory from ARMv8.4, so every Neoverse,
// Graviton and Apple M-series core has it; a Cortex-A72-class core keeps
// the scalar loop). A var rather than a const so TestIntRunMatchesScalar
// can drive the batch loops with the kernel on and off.
var useIntRun = cpu.ARM64.HasASIMDDP

// parseIntRunNEON parses a run of comma-separated integers of one to eight
// digits from data[i:] into out, one int64 per element, and returns how many
// it wrote, where the scalar loop should resume, and whether the run ended at
// the array's ']' (p is then that byte's index, the element before it having
// been written). Same contract as the amd64 parseIntRunSSE; see
// intrun_arm64.s.
//
//go:noescape
func parseIntRunNEON(data []byte, i int, out []int64) (n, p, closed int)

func parseIntRun(data []byte, i int, out []int64) (n, p, closed int) {
	return parseIntRunNEON(data, i, out)
}

// intRunShuffleNEON holds the TBL controls parseIntRunNEON folds a number
// with: entry s*8+L (unique for L in 1..8, the only lengths the kernel
// folds) for an L-digit number starting at block offset s, over the 64-byte
// block. Each entry lays the eight digits of the zero-padded number out
// three, three and two to a word — bytes 0-2, 4-6 and 8-9 — with 0x80
// elsewhere, which the four-register TBL reads as zero (any index >= 64),
// so the kernel's UDOT sums each group with its byte weights and sees
// leading zeros where the number is shorter; see the assembly. 8 KiB,
// dense: the walk touches only the rows of the (offset, length) pairs it
// meets.
var intRunShuffleNEON [64 * 8][16]byte

func init() {
	pos := [8]int{0, 1, 2, 4, 5, 6, 8, 9}
	for s := 0; s < 64; s++ {
		for l := 1; l <= 8 && s+l <= 64; l++ {
			m := &intRunShuffleNEON[s*8+l]
			for t := range m {
				m[t] = 0x80
			}
			for t := 0; t < l; t++ {
				m[pos[8-l+t]] = byte(s + t)
			}
		}
	}
}

// intRunMinSlots is the spare capacity below which the batch loops do not
// call parseIntRun. The arm64 kernel classifies a 64-byte block before it
// stores anything, which breaks even against the scalar loop at four
// elements (N2, presized targets: three elements scalar 49 ns / kernel 53,
// four 56.5 / 57, six 76 / 70, twelve 128 / 95); a presized target's spare
// capacity is the array's remaining length, so below four the loop keeps
// the scalar path and pays no call — twitter_status's two-element indices
// arrays measured +1.0% before this gate.
const intRunMinSlots = 4
