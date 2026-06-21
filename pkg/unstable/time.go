package unstable

import (
	"strconv"
	"time"
)

// rfc3339NanoSpaceLayout is RFC 3339 with the date and time separated by a space
// instead of a 'T', a common variant emitted by databases and log pipelines.
const rfc3339NanoSpaceLayout = "2006-01-02 15:04:05.999999999Z07:00"

// parseJSONTS parses an RFC 3339 timestamp, tolerating a space in place of the
// 'T' date/time separator, and normalizes the result to UTC.
func parseJSONTS(ts string) (time.Time, bool) {
	if t, ok := parseRFC3339(ts, true); ok {
		return t.UTC(), true
	}
	pattern := time.RFC3339Nano
	if len(ts) > 11 && ts[10] == ' ' {
		pattern = rfc3339NanoSpaceLayout
	}
	t, err := time.Parse(pattern, ts)
	if err == nil {
		t = t.UTC()
	}
	return t, err == nil
}

// parseNumTS parses a Unix timestamp expressed as a decimal integer, inferring
// the unit (seconds, milliseconds, or microseconds) from its magnitude. Values
// outside the recognized ranges are rejected.
func parseNumTS(ts string) (time.Time, bool) {
	val, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, false
	}
	if val > 1_000_000_000_000_000 && val < 10_000_000_000_000_000 {
		return time.UnixMicro(val).UTC(), true
	}
	if val > 1_000_000_000_000 && val < 10_000_000_000_000 {
		return time.UnixMilli(val).UTC(), true
	}
	if val > 1_000_000_000 && val < 10_000_000_000 {
		return time.Unix(val, 0).UTC(), true
	}
	return time.Time{}, false
}
