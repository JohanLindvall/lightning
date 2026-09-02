//go:build amd64

package unstable

// readerWordFold selects how ReadInt64OrNull and ReadUint64OrNull fold a digit
// run. On amd64 it is off: the word-at-a-time digitRun puts the cursor on a
// load, mask, count, add data chain of ~14 cycles, and in an object decoder
// the next member's key scan waits on that cursor, where a byte loop advances
// it under a predicted branch and the loads that follow issue speculatively.
// Measured on Zen 4 (2026-09-02) against digitRun, with a plain byte loop:
// instructions per decode instruments -6.8%, golang_source +0.9%,
// payload_large -2.8%, random -2.4%, citm_catalog +1.4%, twitter_status
// -0.7%; cycles golang_source -4.5%, instruments -5%, update_center -3% —
// and a four-digit SWAR step in front of the byte loop cost 0.5-4.8% more
// instructions than the byte loop alone on every one of them (its failing
// attempt on the short ints that dominate). See digits_other.go.
const readerWordFold = false
