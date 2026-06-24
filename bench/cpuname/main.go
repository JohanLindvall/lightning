// Command cpuname prints the detected CPU brand string (e.g. "Apple M2",
// "Intel(R) Core(TM) Ultra 9 185H") on its own line, then exits.
//
// It exists because Go's testing package fills the benchmark `cpu:` header from
// its internal cpu detection, which has no brand string on arm64 — there it
// prints only "unknown", losing the processor model from the committed
// pkg_results_arm64.md / results_arm64.md. github.com/klauspost/cpuid reads the
// brand from sysctl (darwin) or /proc-derived data and reports a real model on
// the arches Go leaves blank, so the bench scripts call this tool and feed the
// result to the markdown renderers.
//
// Output is empty when even cpuid cannot determine a brand; callers treat that
// the same as Go's "unknown" and fall back accordingly.
package main

import (
	"fmt"

	"github.com/klauspost/cpuid/v2"
)

func main() {
	fmt.Println(cpuid.CPU.BrandName)
}
