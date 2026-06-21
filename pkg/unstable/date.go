package unstable

import "time"

// pow10nano[fd] scales a fraction of fd significant digits (1..9) up to whole
// nanoseconds: pow10nano[fd] == 10^(9-fd). Index 0 is unused (fd >= 1).
var pow10nano = [...]int{1e9, 1e8, 1e7, 1e6, 1e5, 1e4, 1e3, 1e2, 1e1, 1e0}

// parseRFC3339 parses an RFC 3339 timestamp on the fast path shared by the time
// decoders, returning the same instant as time.Parse(time.RFC3339Nano, s). It
// returns ok=false — so the caller falls back to time.Parse — for any shape it
// does not handle exactly: a wrong-length or malformed token, an out-of-range
// field, or a timezone that is neither 'Z' nor ±HH:MM. allowSpace additionally
// permits a space in place of the 'T' date/time separator (the lax variant).
// It sidesteps time.Parse's reference-layout machinery, which dominates decoding
// of timestamp-heavy documents, by reading the fixed-offset fields directly.
func parseRFC3339(s string, allowSpace bool) (time.Time, bool) {
	// Shortest accepted form is "2006-01-02T15:04:05Z" (20 bytes).
	if len(s) < 20 {
		return time.Time{}, false
	}
	if sep := s[10]; sep != 'T' && (!allowSpace || sep != ' ') {
		return time.Time{}, false
	}
	if s[4] != '-' || s[7] != '-' || s[13] != ':' || s[16] != ':' {
		return time.Time{}, false
	}
	year, ok0 := atoi4(s, 0)
	month, ok1 := atoi2(s, 5)
	day, ok2 := atoi2(s, 8)
	hour, ok3 := atoi2(s, 11)
	min, ok4 := atoi2(s, 14)
	sec, ok5 := atoi2(s, 17)
	ok := ok0 && ok1 && ok2 && ok3 && ok4 && ok5
	if !ok {
		return time.Time{}, false
	}
	// Reject out-of-range fields so leap seconds and the like defer to time.Parse.
	if month < 1 || month > 12 || day < 1 || day > 31 || hour > 23 || min > 59 || sec > 59 {
		return time.Time{}, false
	}

	i := 19
	nsec := 0
	if s[i] == '.' {
		i++
		start := i
		// Accumulate up to 9 fractional digits (nanosecond precision). Bounding the
		// loop at nine keeps the per-digit `fd < 9` test out of it, and scaling the
		// result by one table multiply replaces the trailing `for fd<9 { nsec*=10 }`.
		end9 := i + 9
		if end9 > len(s) {
			end9 = len(s)
		}
		for i < end9 && s[i] >= '0' && s[i] <= '9' {
			nsec = nsec*10 + int(s[i]-'0')
			i++
		}
		fd := i - start
		if fd == 0 {
			return time.Time{}, false
		}
		nsec *= pow10nano[fd]
		// Any digits past the ninth are below nanosecond resolution; skip them.
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			i++
		}
	}

	if i >= len(s) {
		return time.Time{}, false
	}
	var loc *time.Location
	switch s[i] {
	case 'Z':
		if i+1 != len(s) {
			return time.Time{}, false
		}
		loc = time.UTC
	case '+', '-':
		if i+6 != len(s) || s[i+3] != ':' {
			return time.Time{}, false
		}
		oh, oka := atoi2(s, i+1)
		om, okb := atoi2(s, i+4)
		if !oka || !okb || oh > 23 || om > 59 {
			return time.Time{}, false
		}
		off := (oh*60 + om) * 60
		if s[i] == '-' {
			off = -off
		}
		if off == 0 {
			loc = time.UTC
		} else {
			loc = time.FixedZone("", off)
		}
	default:
		return time.Time{}, false
	}
	if loc == time.UTC {
		// The fields are already range-checked, so build the instant directly and
		// skip time.Date's field-normalization machinery (dateToAbsDays and the
		// surrounding range handling were ~18% of timestamp-array decoding).
		secs := daysFromCivilCached(year, month, day)*86400 + int64(hour)*3600 + int64(min)*60 + int64(sec)
		return time.Unix(secs, int64(nsec)).UTC(), true
	}
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc), true
}

// daysFromCivil returns the number of days from 1970-01-01 to the given
// proleptic-Gregorian date using Howard Hinnant's algorithm, which is branch-
// light and valid across time.Time's full year range. month is 1-12 and day
// 1-31; the caller has already range-checked them. It lets the RFC 3339 fast
// path construct UTC timestamps via time.Unix.
func daysFromCivil(y, m, d int) int64 {
	yy := int64(y)
	if m <= 2 {
		yy-- // treat Jan/Feb as months 13/14 of the previous year
	}
	var era int64
	if yy >= 0 {
		era = yy / 400
	} else {
		era = (yy - 399) / 400
	}
	yoe := yy - era*400 // year of era, [0, 399]
	mm := int64(m)
	mp := mm - 3 // March-based month index
	if mm <= 2 {
		mp = mm + 9
	}
	doy := (153*mp+2)/5 + int64(d) - 1     // day of year, [0, 365]
	doe := yoe*365 + yoe/4 - yoe/100 + doy // day of era, [0, 146096]
	return era*146097 + doe - 719468       // 719468 = days from 0000-03-01 to 1970-01-01
}

// monthYearDay[m-1] is the number of days before the first of month m in a
// non-leap year (January = index 0).
var monthYearDay = [12]int32{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

// yearStartDays[i] is the number of days from 1970-01-01 to (1970+i)-01-01, built
// once from daysFromCivil (so it cannot disagree with it). It spans 1970 through
// 2261, the years that dominate real timestamps.
var yearStartDays = func() [292]int32 {
	var t [292]int32
	for i := range t {
		t[i] = int32(daysFromCivil(1970+i, 1, 1))
	}
	return t
}()

// daysFromCivilCached is daysFromCivil specialized for the common case. A year in
// [1970, 2261] resolves to two table lookups, the day, and a single leap-day
// bump — skipping the general algorithm's several multiplies and divisions —
// while any other year defers to it. The caller has range-checked m and d.
func daysFromCivilCached(y, m, d int) int64 {
	if uint(y-1970) < uint(len(yearStartDays)) {
		days := int64(yearStartDays[y-1970]) + int64(monthYearDay[m-1]) + int64(d-1)
		if m > 2 && y%4 == 0 && (y%100 != 0 || y%400 == 0) {
			days++ // leap day of this year has already passed
		}
		return days
	}
	return daysFromCivil(y, m, d)
}

// atoi2 parses exactly two decimal digits of s at off into a non-negative int,
// returning false if either byte is not an ASCII digit. The branchless digit
// test and the absence of a loop (n is fixed) make it noticeably cheaper than a
// general atoiN on the RFC 3339 fast path, which parses six such fields per
// timestamp. The caller guarantees off+2 <= len(s).
func atoi2(s string, off int) (int, bool) {
	d0 := s[off] - '0'
	d1 := s[off+1] - '0'
	return int(d0)*10 + int(d1), d0 <= 9 && d1 <= 9
}

// atoi4 parses exactly four decimal digits of s at off (the RFC 3339 year).
// The caller guarantees off+4 <= len(s).
func atoi4(s string, off int) (int, bool) {
	d0 := s[off] - '0'
	d1 := s[off+1] - '0'
	d2 := s[off+2] - '0'
	d3 := s[off+3] - '0'
	return int(d0)*1000 + int(d1)*100 + int(d2)*10 + int(d3), d0 <= 9 && d1 <= 9 && d2 <= 9 && d3 <= 9
}
