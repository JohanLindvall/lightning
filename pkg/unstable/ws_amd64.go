//go:build amd64

package unstable

// skipWSSpaceShortcut selects SkipWSRun's per-word test for an all-spaces
// word. On amd64 it is off: the shortcut's equality test costs a taken branch
// per space word on top of the loop's back edge, and on a taken-branch-bound
// front end (Meteor Lake decodes citm_catalog at one taken branch every 2.5
// cycles) that second jump outweighs the three ALU operations the shortcut
// skips — measured citm_catalog −3.6%, instruments −2.1%, twitter_status
// −1.6% with the shortcut removed, nothing worse. See ws_other.go for the
// arm64 side of that trade.
const skipWSSpaceShortcut = false
