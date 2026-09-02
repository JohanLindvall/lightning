//go:build !amd64

package unstable

// skipWSSpaceShortcut keeps SkipWSRun's all-spaces equality test on every
// architecture but amd64. On an issue-bound core (Neoverse N2, IPC 3.6 on
// citm_catalog) instructions removed are cycles removed, and dropping the
// shortcut measured +16% on citm-shaped indent runs there; on amd64 the
// taken branch it costs dominates instead (see ws_amd64.go). A compile-time
// constant, so each architecture's loop is exactly what it measured best.
const skipWSSpaceShortcut = true
