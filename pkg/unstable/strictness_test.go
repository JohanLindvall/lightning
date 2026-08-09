package unstable

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// This file holds the strictness invariants for the two decoders that used to
// accept input the rest of the library rejects: the RFC 3339 fast path (which
// let an impossible day-of-month roll over into a plausible wrong date) and the
// json.Number readers (which captured any run of number bytes verbatim). Both
// bugs shared a shape — a scanner that measures a token without validating it —
// so both fixes are pinned here by differential tests against the authority each
// decoder must agree with: time.Parse for dates, ReadFloat64OrNull for numbers.

// --- C4: RFC 3339 day-of-month validation ----------------------------------

// quoteJSON wraps s as a JSON string. Every timestamp used here is escape-free,
// so plain concatenation is exact.
func quoteJSON(s string) []byte { return []byte(`"` + s + `"`) }

// TestReadTimeRejectsImpossibleDates covers the C4 defect: the fast path
// range-checked the day as 1..31 only, so daysFromCivil silently normalised an
// impossible date into a real one — "2021-02-31T00:00:00Z" decoded as 2021-03-03
// with a nil error, and Feb 29 of a non-leap year (the case a real producer is
// most likely to emit) as March 1.
func TestReadTimeRejectsImpossibleDates(t *testing.T) {
	for _, s := range []string{
		"2021-02-29T00:00:00Z", // not a leap year
		"1900-02-29T00:00:00Z", // century, not a leap year
		"2100-02-29T00:00:00Z", // century, not a leap year, outside the cached window's habits
		"2021-02-30T00:00:00Z",
		"2021-02-31T00:00:00Z",
		"2021-04-31T00:00:00Z",
		"2021-06-31T00:00:00Z",
		"2021-09-31T00:00:00Z",
		"2021-11-31T00:00:00Z",
		"2021-01-32T00:00:00Z",
		"2021-00-10T00:00:00Z", // month 0
		"2021-13-01T00:00:00Z", // month 13
		"2021-01-00T00:00:00Z", // day 0
		"2021-02-31T12:34:56.789+02:00",
	} {
		t.Run(s, func(t *testing.T) {
			// The fast path itself must decline, in both separator modes.
			if _, ok := parseRFC3339(s, false); ok {
				t.Errorf("parseRFC3339(%q, false) accepted an impossible date", s)
			}
			if _, ok := parseRFC3339(s, true); ok {
				t.Errorf("parseRFC3339(%q, true) accepted an impossible date", s)
			}
			// And the readers must surface the stdlib's error rather than a value.
			got, _, err := ReadTimeOrNull(quoteJSON(s), 0)
			if err == nil {
				t.Errorf("ReadTimeOrNull(%q) = %v, want an error", s, got)
			}
			if _, _, err := ReadTimeLaxOrNull(quoteJSON(s), 0); err == nil {
				t.Errorf("ReadTimeLaxOrNull(%q) accepted an impossible date", s)
			}
			// The premise: encoding/json's parser rejects these too.
			if _, perr := time.Parse(time.RFC3339, s); perr == nil {
				t.Fatalf("premise broken: time.Parse accepts %q", s)
			}
		})
	}
}

// TestReadTimeAcceptsEdgeDates guards the other side of the C4 fix: the added
// day-length check must not reject any real date. It covers 29 February in leap
// years (including the 400-year century rule), month ends, and years on both
// sides of the daysFromCivilCached [1970, 2261] window so the general
// daysFromCivil fallback is exercised too.
func TestReadTimeAcceptsEdgeDates(t *testing.T) {
	for _, s := range []string{
		"2020-02-29T00:00:00Z", // leap year
		"2000-02-29T00:00:00Z", // divisible by 400: leap
		"1600-02-29T00:00:00Z", // leap, below the cached window
		"2400-02-29T00:00:00Z", // leap, above the cached window
		"2021-01-01T00:00:00Z",
		"2021-12-31T23:59:59Z",
		"2021-01-31T00:00:00Z",
		"2021-04-30T00:00:00Z",
		"2021-02-28T00:00:00Z",
		"1969-12-31T23:59:59Z", // below the cached window
		"1970-01-01T00:00:00Z", // first cached year
		"2261-12-31T00:00:00Z", // last cached year
		"2262-01-01T00:00:00Z", // first year past the cache
		"2020-02-29T12:34:56.987654321+02:00",
	} {
		t.Run(s, func(t *testing.T) {
			want, perr := time.Parse(time.RFC3339, s)
			if perr != nil {
				t.Fatalf("premise broken: time.Parse rejects %q: %v", s, perr)
			}
			got, _, err := ReadTimeOrNull(quoteJSON(s), 0)
			if err != nil {
				t.Fatalf("ReadTimeOrNull(%q): %v", s, err)
			}
			if !got.Equal(want) {
				t.Errorf("ReadTimeOrNull(%q) = %v, want %v", s, got, want)
			}
		})
	}
}

// dateCorpus enumerates every month 0..13 crossed with the interesting days
// (including the impossible 0 and 32) for years chosen to cover the leap rules
// and both sides of the daysFromCivilCached window.
func dateCorpus() []string {
	years := []int{1600, 1900, 1969, 1970, 1999, 2000, 2019, 2020, 2021, 2100, 2261, 2262, 2400}
	days := []int{0, 1, 27, 28, 29, 30, 31, 32}
	out := make([]string, 0, len(years)*14*len(days))
	for _, y := range years {
		for m := 0; m <= 13; m++ {
			for _, d := range days {
				out = append(out, fmt.Sprintf("%04d-%02d-%02dT01:02:03Z", y, m, d))
			}
		}
	}
	return out
}

// TestReadTimeMatchesStdlibAcceptance is the C4 differential: over the whole
// date corpus, ReadTimeOrNull must accept exactly the timestamps
// time.Parse(time.RFC3339, ...) accepts, and produce the same instant. The fast
// path is only ever allowed to be *more* conservative than the stdlib (anything
// it declines falls through to time.Parse); this pins the other direction, which
// is where C4 lived.
func TestReadTimeMatchesStdlibAcceptance(t *testing.T) {
	for _, s := range dateCorpus() {
		want, perr := time.Parse(time.RFC3339, s)
		got, _, err := ReadTimeOrNull(quoteJSON(s), 0)
		if (err == nil) != (perr == nil) {
			t.Fatalf("%q: lightning err=%v (value %v), stdlib err=%v", s, err, got, perr)
		}
		if err == nil && !got.Equal(want) {
			t.Fatalf("%q: lightning = %v, stdlib = %v", s, got, want)
		}
	}
}

// --- C5: json.Number literal validation ------------------------------------

// TestReadNumberRejectsMalformed covers the C5 defect: the json.Number readers
// captured skipNumber's raw token span verbatim, so a field could end up holding
// a string that fails at every later .Float64()/.Int64() call, far from the
// decode — and that this library's own json.Valid rejects.
func TestReadNumberRejectsMalformed(t *testing.T) {
	for _, in := range []string{
		"1.2.3",
		"-",
		"--1",
		"1e",
		"1e+",
		"1e-",
		"1E",
		".",
		"-.",
		".e1",
		"1e2e3",
		"1..2",
		"+",
		"1e309", // magnitude overflows float64, so the float reader rejects it
		"-1e309",
		"1e400",
	} {
		t.Run(in, func(t *testing.T) {
			for _, nocopy := range []bool{false, true} {
				var got string
				var err error
				if nocopy {
					got, _, err = ReadNumberNoCopyOrNull([]byte(in), 0)
				} else {
					got, _, err = ReadNumberOrNull([]byte(in), 0)
				}
				if err == nil {
					t.Errorf("nocopy=%v: ReadNumber(%q) = %q, want an error", nocopy, in, got)
				}
			}
		})
	}
}

// TestReadNumberAcceptsOrdinary guards the other side: the validation must not
// narrow the accepted set beyond ReadFloat64OrNull's. In particular "01" stays
// accepted — the float reader accepts it and so does json.Valid, so rejecting it
// here would create a fresh Valid-vs-decoder disagreement in the opposite
// direction (encoding/json rejects it; that divergence is pre-existing and
// deliberate, see CLAUDE.md).
func TestReadNumberAcceptsOrdinary(t *testing.T) {
	for _, in := range []string{
		"0", "1", "-1", "01", "007", "-0",
		"3.14159", "-6.022e23", "1e3", "2E+2", "5e-2",
		// Leniencies the float reader already has, kept deliberately: a missing
		// integer or fraction part, and an exponent that underflows to zero
		// (strconv reports a range error only on overflow).
		".5", "-.5", "1.", "1e-400",
		"12345678901234567890",                    // beyond uint64, kept verbatim
		"123456789012345678901234567890.12345678", // far beyond float64 precision
		"1e308", "-1e308", "1e-300",
	} {
		t.Run(in, func(t *testing.T) {
			for _, nocopy := range []bool{false, true} {
				var got string
				var end int
				var err error
				if nocopy {
					got, end, err = ReadNumberNoCopyOrNull([]byte(in), 0)
				} else {
					got, end, err = ReadNumberOrNull([]byte(in), 0)
				}
				if err != nil {
					t.Fatalf("nocopy=%v: ReadNumber(%q): %v", nocopy, in, err)
				}
				if got != in || end != len(in) {
					t.Errorf("nocopy=%v: ReadNumber(%q) = %q, end %d; want the literal verbatim", nocopy, in, got, end)
				}
			}
		})
	}
}

// numberCorpus builds every combination of a sign, an integer part, a fraction
// and an exponent — well-formed and malformed alike — plus a set of literals
// followed by the delimiters a decoder would really see.
func numberCorpus() []string {
	signs := []string{"", "-", "+"}
	ints := []string{"", "0", "1", "01", "10", "9007199254740993", "1234567890123456789012345"}
	fracs := []string{"", ".", ".0", ".5", ".00", ".0000000000000000001", ".123456789012345678901"}
	exps := []string{"", "e", "E", "e+", "e-", "e0", "e5", "E12", "e+308", "e309", "e-400", "e-9999999999"}
	out := make([]string, 0, len(signs)*len(ints)*len(fracs)*len(exps)+32)
	for _, s := range signs {
		for _, i := range ints {
			for _, f := range fracs {
				for _, e := range exps {
					out = append(out, s+i+f+e)
				}
			}
		}
	}
	// Tokens as they appear inside a document, plus shapes that stress where the
	// scanner decides a token ends.
	out = append(out,
		"", " ", "x", `"7.5"`, "null", "nul", "n",
		"1]", "1}", "1,", "1 ", "12,34", "1-2", "1+2", "1.2}", "-0]",
		"1.2.3", "--1", "++1", "-+1", "1e2e3", "0x10", "0b1", "Infinity", "NaN",
	)
	return out
}

// TestReadNumberAcceptSetMatchesFloat64 is the invariant that keeps json.Valid
// honest. Valid checks a number by running it through ReadFloat64OrNull, so a
// json.Number field must capture a literal exactly when that reader would have
// parsed one — same accept/reject decision, same error identity, same token end.
// Without this the two disagree in whichever direction the readers happen to
// drift, which is precisely the C5 defect (skipNumber's span is a superset of
// scanFloat's grammar).
func TestReadNumberAcceptSetMatchesFloat64(t *testing.T) {
	for _, in := range numberCorpus() {
		data := []byte(in)
		_, fend, ferr := ReadFloat64OrNull(data, 0)
		for _, nocopy := range []bool{false, true} {
			var got string
			var end int
			var err error
			if nocopy {
				got, end, err = ReadNumberNoCopyOrNull(data, 0)
			} else {
				got, end, err = ReadNumberOrNull(data, 0)
			}
			if !errors.Is(err, ferr) || !errors.Is(ferr, err) {
				t.Fatalf("%q (nocopy=%v): ReadNumber err = %v, ReadFloat64 err = %v", in, nocopy, err, ferr)
			}
			if end != fend {
				t.Fatalf("%q (nocopy=%v): ReadNumber end = %d, ReadFloat64 end = %d", in, nocopy, end, fend)
			}
			if err != nil || in == "null" {
				continue
			}
			// An accepted literal is the verbatim token and is usable as a
			// json.Number: the whole point of rejecting the malformed ones is that
			// .Float64() no longer fails long after the decode.
			if want := in[:end]; got != want {
				t.Fatalf("%q (nocopy=%v): ReadNumber = %q, want %q", in, nocopy, got, want)
			}
			if _, perr := strconv.ParseFloat(got, 64); perr != nil {
				t.Fatalf("%q (nocopy=%v): accepted literal %q fails ParseFloat: %v", in, nocopy, got, perr)
			}
		}
	}
}
