package json

import (
	"fmt"
	"strings"
	"testing"
)

// TestArrayEachDispatchMatchesSkipValue pins the one-compare number dispatch
// arrayEach and arrayEachIndex use in place of a SkipValue call. The claim is
// that uint(c-'-') <= 12 — '-', '.', '/' and the ten digits — selects exactly
// the leading bytes SkipValue sends to its own number arm, so the walk is
// unchanged for every input. The check is over EVERY leading byte, because a
// predicate that is a byte too wide would silently measure a string, a
// container or a literal with the number scanner.
//
// ObjectEach is the reference: it is the same walk over the same values and it
// still calls SkipValue, so any disagreement is the dispatch.
func TestArrayEachDispatchMatchesSkipValue(t *testing.T) {
	// Tails that make the leading byte the start of a real token, a malformed
	// one, and one that runs to the end of the document.
	// Every value kind, whole and well-formed: if the dispatch ever routed a
	// string, a container or a literal to the number scanner, these are what
	// would measure wrong. (A byte-level predicate that is too NARROW is not a
	// defect — the byte simply takes SkipValue, which routes it to the same
	// scanner — so what is worth pinning is the too-wide direction.)
	for _, tok := range []string{
		`"abc"`, `""`, `"a,b]c}"`, `"with \"escape\" and \\"`, `"[{"`,
		`{}`, `{"a":1}`, `{"a":[1,{"b":"]"}]}`, `[]`, `[1,2]`, `[[1],[2]]`,
		`true`, `false`, `null`, `0`, `-1.5e-7`, `1788087600123`,
	} {
		arr := "[" + tok + "," + tok + "]"
		obj := `{"a":` + tok + `,"b":` + tok + "}"
		var gotArr, gotObj []string
		errArr := ArrayEach([]byte(arr), func(v []byte) error { gotArr = append(gotArr, string(v)); return nil })
		errObj := ObjectEach([]byte(obj), func(_ string, v []byte) error { gotObj = append(gotObj, string(v)); return nil })
		if errArr != nil || errObj != nil {
			t.Fatalf("%s: ArrayEach err=%v, ObjectEach err=%v", tok, errArr, errObj)
		}
		if fmt.Sprint(gotArr) != fmt.Sprint(gotObj) || len(gotArr) != 2 || gotArr[0] != tok {
			t.Fatalf("%s: ArrayEach %q, ObjectEach %q", tok, gotArr, gotObj)
		}
	}

	tails := []string{"", "1", "x", `"`, "23e4", "0.5", "true", "\x00"}
	for b := 0x21; b < 256; b++ {
		// A byte that opens or closes a container, or a comma, would make the
		// two documents differ in structure rather than in the value under
		// test; <= 0x20 is whitespace to SkipWS, where an empty element and an
		// empty member value legitimately differ.
		if strings.IndexByte(`[]{}",`, byte(b)) >= 0 {
			continue
		}
		for _, tail := range tails {
			tok := string(rune(b)) + tail
			arr := "[" + tok + "]"
			obj := `{"k":` + tok + "}"

			var gotArr, gotObj [][]byte
			errArr := ArrayEach([]byte(arr), func(v []byte) error {
				gotArr = append(gotArr, append([]byte(nil), v...))
				return nil
			})
			errObj := ObjectEach([]byte(obj), func(_ string, v []byte) error {
				gotObj = append(gotObj, append([]byte(nil), v...))
				return nil
			})
			// The two walkers report different sentinels for a missing
			// container, but for the value itself they must agree on the span
			// and on whether it was accepted at all.
			if (errArr == nil) != (errObj == nil) {
				t.Fatalf("byte %#02x tail %q: ArrayEach err=%v, ObjectEach err=%v", b, tail, errArr, errObj)
			}
			if len(gotArr) != len(gotObj) {
				t.Fatalf("byte %#02x tail %q: ArrayEach gave %d values, ObjectEach %d", b, tail, len(gotArr), len(gotObj))
			}
			for i := range gotArr {
				if string(gotArr[i]) != string(gotObj[i]) {
					t.Fatalf("byte %#02x tail %q: value %d is %q from ArrayEach and %q from ObjectEach",
						b, tail, i, gotArr[i], gotObj[i])
				}
			}
			// The indexed walker is separate code and must agree too.
			var gotIdx [][]byte
			errIdx := ArrayEachIndex([]byte(arr), func(n int, v []byte) error {
				if n != len(gotIdx) {
					t.Fatalf("byte %#02x: index %d at position %d", b, n, len(gotIdx))
				}
				gotIdx = append(gotIdx, append([]byte(nil), v...))
				return nil
			})
			if (errIdx == nil) != (errArr == nil) || len(gotIdx) != len(gotArr) {
				t.Fatalf("byte %#02x tail %q: ArrayEachIndex err=%v n=%d, ArrayEach err=%v n=%d",
					b, tail, errIdx, len(gotIdx), errArr, len(gotArr))
			}
			for i := range gotIdx {
				if string(gotIdx[i]) != string(gotArr[i]) {
					t.Fatalf("byte %#02x tail %q: value %d is %q indexed and %q plain", b, tail, i, gotIdx[i], gotArr[i])
				}
			}
		}
	}
}

// TestArrayEachNumberSpansMatchObjectEach walks arrays of every number shape
// the scanner measures, including the ones its grammar deliberately over-
// accepts (1.2.3, +-e.), and holds the spans to ObjectEach's.
func TestArrayEachNumberSpansMatchObjectEach(t *testing.T) {
	toks := []string{"0", "-0", "7", "-7", "200", "1788087600", "1788087600123",
		"9223372036854775807", "18446744073709551615", "1.5", "-65.613617000000029",
		"1e3", "1E3", "1.5e-7", "1e+308", "0.000698875266", "1.2.3", "+-e.", "-", ".", "e", "+",
		"00", "0123", "1-2", "1e", "1.", ".5"}
	for _, tok := range toks {
		for _, n := range []int{1, 2, 5} {
			arr, obj := "[", "{"
			for i := 0; i < n; i++ {
				if i > 0 {
					arr, obj = arr+",", obj+","
				}
				arr += tok
				obj += fmt.Sprintf("%q:%s", fmt.Sprintf("k%d", i), tok)
			}
			arr, obj = arr+"]", obj+"}"
			var gotArr, gotObj []string
			errArr := ArrayEach([]byte(arr), func(v []byte) error { gotArr = append(gotArr, string(v)); return nil })
			errObj := ObjectEach([]byte(obj), func(_ string, v []byte) error { gotObj = append(gotObj, string(v)); return nil })
			if (errArr == nil) != (errObj == nil) || fmt.Sprint(gotArr) != fmt.Sprint(gotObj) {
				t.Fatalf("%q x%d: array %v %v, object %v %v", tok, n, gotArr, errArr, gotObj, errObj)
			}
		}
	}
}
