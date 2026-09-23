package conformance

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

// geoDocStd and pointRingStd are the methodless twins
// TestFloat64PointsMatchStdlib compares with (see
// TestStdlibTwinsAreReflectionOnly).
type geoDocStd GeoDoc

type pointRingStd PointRing

// geoNum draws a number from both sides of the points walk's line: shapes it
// converts (15-digit coordinates, shortest forms, integers, signed zeros) and
// shapes it hands back to the scalar reader (exponents, 20+ digits, and the
// Eisel-Lemire cases eiselLemire64 refines, which cluster on values exact in
// binary like 1234567890123456.5).
func geoNum(rng *rand.Rand) string {
	switch r := rng.Intn(24); {
	case r < 8:
		return strconv.FormatFloat((rng.Float64()-0.5)*360, 'f', 15, 64)
	case r < 12:
		return strconv.FormatFloat((rng.Float64()-0.5)*360, 'f', -1, 64)
	case r < 15:
		return strconv.Itoa(rng.Intn(2000) - 1000)
	case r < 17:
		return strconv.FormatFloat(rng.Float64(), 'f', 6, 64)
	default:
		return []string{"1e5", "-2.5E-3", "6.02e23", "12345678901234567890123", "1234567890123456.5",
			"-9007199254740993.25", "-0", "0", "-0.0"}[rng.Intn(9)]
	}
}

// geoRing builds a ring of count points of n numbers each; odd mixes in
// nulls and points of the wrong length, which decode like any fixed array.
func geoRing(rng *rand.Rand, n, count int, odd bool) string {
	var b strings.Builder
	b.WriteByte('[')
	for k := 0; k < count; k++ {
		if k > 0 {
			b.WriteString([]string{",", ", ", " , ", ",\n    "}[rng.Intn(4)])
		}
		if odd && rng.Intn(30) == 0 {
			b.WriteString("null")
			continue
		}
		m := n
		if odd && rng.Intn(15) == 0 {
			m = []int{0, 1, n + 1}[rng.Intn(3)]
		}
		b.WriteByte('[')
		for j := 0; j < m; j++ {
			if j > 0 {
				b.WriteString([]string{",", ", ", " , "}[rng.Intn(3)])
			}
			b.WriteString(geoNum(rng))
		}
		b.WriteByte(']')
	}
	b.WriteByte(']')
	return b.String()
}

// sameFloats compares two point slices bit for bit, treating nil and empty
// as equal: lightning leaves a fresh target nil for [] where encoding/json
// makes it empty, a difference the rest of the package already documents.
func sameFloats[T [2]float64 | [3]float64](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		for j := range len(a[k]) {
			if math.Float64bits(a[k][j]) != math.Float64bits(b[k][j]) {
				return false
			}
		}
	}
	return true
}

func sameGeo(a, b GeoDoc) bool {
	if !sameFloats(a.Ring2, b.Ring2) || !sameFloats(a.Ring3, b.Ring3) || !sameFloats(a.Lax, b.Lax) ||
		a.Tail != b.Tail || len(a.Polygon) != len(b.Polygon) {
		return false
	}
	for k := range a.Polygon {
		if !sameFloats(a.Polygon[k], b.Polygon[k]) {
			return false
		}
	}
	return true
}

// TestFloat64PointsMatchStdlib holds the generated decode of [][N]float64 —
// unstable.DecodeFloat64Points, whose SIMD points walk takes runs of points in
// one call and hands every other shape back to the per-point reader — to
// encoding/json: in a field, nested in a polygon, under lax and at a named
// root; over rings of every size and separator, numbers either side of the
// walk's line, nulls and wrong-length points; fresh and reused; and, for
// malformed documents (truncated, a trailing comma, a missing comma), in
// agreeing that they are errors.
func TestFloat64PointsMatchStdlib(t *testing.T) {
	rng := rand.New(rand.NewSource(3))
	size := func() int {
		if rng.Intn(5) == 0 {
			return 60 + rng.Intn(500)
		}
		return rng.Intn(60)
	}
	var prev GeoDoc
	var prevStd geoDocStd
	for it := 0; it < 1500; it++ {
		odd := rng.Intn(3) == 0
		doc := fmt.Sprintf(`{"ring2":%s,"ring3": %s, "polygon":[%s, %s],"lax":%s,"tail":7}`,
			geoRing(rng, 2, size(), odd), geoRing(rng, 3, size(), odd),
			geoRing(rng, 2, size(), odd), geoRing(rng, 2, size(), false), geoRing(rng, 2, size(), false))
		malformed := false
		switch rng.Intn(8) {
		case 0:
			doc, malformed = doc[:rng.Intn(len(doc))], true
		case 1:
			if k := strings.Index(doc, "]]"); k >= 0 {
				doc, malformed = doc[:k]+"],]"+doc[k+2:], true // a trailing comma in a ring
			}
		case 2:
			if k := strings.LastIndex(doc, "],["); k >= 0 {
				doc, malformed = doc[:k]+"] ["+doc[k+3:], true // a missing comma between points
			}
		}
		var got GeoDoc
		gerr := got.UnmarshalJSON([]byte(doc))
		var want geoDocStd
		werr := json.Unmarshal([]byte(doc), &want)
		if (gerr == nil) != (werr == nil) || (gerr == nil) == malformed {
			t.Fatalf("iteration %d: lightning err=%v, encoding/json err=%v, malformed=%v on %.300q", it, gerr, werr, malformed, doc)
		}
		if gerr != nil {
			continue
		}
		if !sameGeo(got, GeoDoc(want)) {
			t.Fatalf("iteration %d: decoded differently from encoding/json on %.300q", it, doc)
		}
		// Reused: decoding into last round's value must replace its rings,
		// as encoding/json does. Nulls are left out: a null point in a reused
		// slot is where the two already part (encoding/json leaves the slot's
		// old values, this zeroes it).
		if !odd {
			rerr := prev.UnmarshalJSON([]byte(doc))
			serr := json.Unmarshal([]byte(doc), &prevStd)
			if rerr != nil || serr != nil || !sameGeo(prev, GeoDoc(prevStd)) {
				t.Fatalf("iteration %d: reused decode differs (err %v / %v) on %.300q", it, rerr, serr, doc)
			}
		}
	}

	// A named root, and a lax field given something that is not a ring.
	for it := 0; it < 300; it++ {
		ring := geoRing(rng, 2, size(), rng.Intn(3) == 0)
		var got PointRing
		var want pointRingStd
		gerr := got.UnmarshalJSON([]byte(ring))
		werr := json.Unmarshal([]byte(ring), &want)
		if gerr != nil || werr != nil || !sameFloats(got, want) {
			t.Fatalf("root: err %v / %v on %.300q", gerr, werr, ring)
		}
	}
	var l GeoDoc
	if err := l.UnmarshalJSON([]byte(`{"lax":{"x":[1,2]},"tail":3}`)); err != nil || l.Lax != nil || l.Tail != 3 {
		t.Fatalf("lax mismatch: err=%v %+v", err, l)
	}
	if err := l.UnmarshalJSON([]byte(`{"lax":[[1,2],[3,"x"]],"tail":4}`)); err != nil || l.Lax != nil || l.Tail != 4 {
		t.Fatalf("lax element mismatch: err=%v %+v", err, l)
	}
	if err := l.UnmarshalJSON([]byte(`{"lax":[[1,2],[3,4],]}`)); err == nil {
		t.Fatalf("lax accepted a trailing comma: %+v", l)
	}
}
