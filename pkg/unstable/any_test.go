package unstable

import (
	"errors"
	"reflect"
	"testing"
)

// --- DecodeValue (and decodeAnyObject/decodeAnyArray) ----------------------

func TestDecodeValue(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`"hi"`), 0)
		if err != nil || v != "hi" || end != 4 {
			t.Fatalf("got %v, %d, %v", v, end, err)
		}
	})
	t.Run("number", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`3.5`), 0)
		if err != nil || v != 3.5 {
			t.Fatalf("got %v, %v", v, err)
		}
	})
	t.Run("true", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`true`), 0)
		if err != nil || v != true {
			t.Fatalf("got %v, %v", v, err)
		}
	})
	t.Run("false", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`false`), 0)
		if err != nil || v != false {
			t.Fatalf("got %v, %v", v, err)
		}
	})
	t.Run("null", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`null`), 0)
		if err != nil || v != nil || end != 4 {
			t.Fatalf("got %v, %d, %v", v, end, err)
		}
	})
	t.Run("empty truncated", func(t *testing.T) {
		_, _, err := DecodeValue(nil, 0)
		if !errors.Is(err, ErrTruncated) {
			t.Fatalf("err = %v, want ErrTruncated", err)
		}
	})

	t.Run("object", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`{"a":1,"b":"x","c":true}`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		m, ok := v.(map[string]any)
		if !ok {
			t.Fatalf("got %T, want map", v)
		}
		if m["a"] != float64(1) || m["b"] != "x" || m["c"] != true {
			t.Errorf("unexpected map: %v", m)
		}
		if end != 24 {
			t.Errorf("end = %d, want 24", end)
		}
	})

	t.Run("empty object", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`{}`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		m := v.(map[string]any)
		if len(m) != 0 || end != 2 {
			t.Errorf("got %v, end %d", m, end)
		}
	})

	t.Run("object with whitespace", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`{ "a" : 1 , "b" : 2 }`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		m := v.(map[string]any)
		if m["a"] != float64(1) || m["b"] != float64(2) {
			t.Errorf("unexpected map: %v", m)
		}
	})

	t.Run("nested array", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`[1,"two",[3,4],{"k":5}]`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		a := v.([]any)
		if len(a) != 4 {
			t.Fatalf("len = %d, want 4", len(a))
		}
		if a[0] != float64(1) || a[1] != "two" {
			t.Errorf("unexpected array head: %v", a)
		}
		inner := a[2].([]any)
		if inner[0] != float64(3) || inner[1] != float64(4) {
			t.Errorf("unexpected inner array: %v", inner)
		}
		obj := a[3].(map[string]any)
		if obj["k"] != float64(5) {
			t.Errorf("unexpected inner object: %v", obj)
		}
	})

	t.Run("empty array", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`[]`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if len(v.([]any)) != 0 || end != 2 {
			t.Errorf("got %v, end %d", v, end)
		}
	})

	// Error paths for the object/array decoders.
	objErrs := []struct {
		name    string
		in      string
		wantErr error
	}{
		{"obj truncated after brace", `{`, ErrTruncated},
		{"obj bad key", `{1:2}`, ErrInvalidJSON},
		{"obj missing colon", `{"a" 1}`, ErrExpectColon},
		{"obj truncated after value", `{"a":1`, ErrTruncated},
		{"obj missing comma", `{"a":1 "b":2}`, ErrInvalidJSON},
		{"obj bad value", `{"a":}`, ErrBadNumber},
		{"obj truncated value", `{"a":`, ErrTruncated},
		{"arr truncated after bracket", `[`, ErrTruncated},
		{"arr truncated after value", `[1`, ErrTruncated},
		{"arr missing comma", `[1 2]`, ErrInvalidJSON},
		{"arr bad value", `[}]`, ErrBadNumber},
	}
	for _, tt := range objErrs {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := DecodeValue([]byte(tt.in), 0)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecodeValueCompact(t *testing.T) {
	compact := []byte(`{"a":1,"b":[true,null,"x"],"c":{"d":2.5}}`)
	want, end1, err1 := DecodeValue(compact, 0)
	got, end2, err2 := DecodeValueCompact(compact, 0)
	if err1 != nil || err2 != nil {
		t.Fatalf("DecodeValue err=%v, DecodeValueCompact err=%v", err1, err2)
	}
	if end1 != end2 {
		t.Fatalf("end mismatch: %d vs %d", end1, end2)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("compact decode != non-compact\n got %#v\nwant %#v", got, want)
	}
	// Like DecodeValue, the compact variant assumes data[i] is the value start
	// (the caller has skipped any leading whitespace), so both decode a bare
	// scalar identically.
	if v, _, err := DecodeValueCompact([]byte(`42`), 0); err != nil || v != 42.0 {
		t.Fatalf("scalar: got %v, %v", v, err)
	}
}

// TestDaysFromCivilCached pins the table-based fast path to the general
// algorithm across (and beyond) the cached year range, including leap years.
