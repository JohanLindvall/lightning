package unstable

import "testing"

func TestDecodeValueEmptyArrayOwnership(t *testing.T) {
	data := []byte("[]")
	for _, read := range []func([]byte, int) (any, int, error){
		DecodeValue, DecodeValueCompact, DecodeValueNumber, DecodeValueNumberCompact,
	} {
		v, end, err := read(data, 0)
		if err != nil || end != len(data) {
			t.Fatalf("end=%d err=%v", end, err)
		}
		a, ok := v.([]any)
		if !ok || a == nil || len(a) != 0 || cap(a) != 0 {
			t.Fatalf("got %#v; want a non-nil empty []any with zero capacity", v)
		}
		a = append(a, "owned by caller")
		v, _, err = read(data, 0)
		if err != nil || len(v.([]any)) != 0 || a[0] != "owned by caller" {
			t.Fatalf("appending changed another decode: %#v, %v", v, err)
		}
		if allocs := testing.AllocsPerRun(100, func() {
			if _, _, err := read(data, 0); err != nil {
				panic(err)
			}
		}); allocs != 0 {
			t.Fatalf("empty array allocated %v times", allocs)
		}
	}
}
