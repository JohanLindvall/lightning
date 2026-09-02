//go:build !amd64

package unstable

// readerWordFold keeps the word-at-a-time digitRun in the single-value integer
// readers on every architecture but amd64: on a Neoverse N2 it measured
// golang_source -5.2%, twitter_status -1.8%, citm_catalog -1.4% against the
// four-digit fold and byte tail (CLAUDE.md, 2026-09-02), the trade going the
// other way on Zen 4 (see digits_amd64.go). A compile-time constant, so each
// architecture's reader is exactly the one it measured best.
const readerWordFold = true
