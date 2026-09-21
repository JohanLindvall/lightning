package unstable

import (
	"bytes"
	stdjson "encoding/json"
	"reflect"
	"strconv"
	"testing"
)

func TestDecodeValueArrayBoundaries(t *testing.T) {
	elements := []string{`{"key":"value\n"}`, `[1,2,3]`, `true`, `null`, `"text"`, `12345.5`, `[]`, `{}`}
	for _, n := range []int{0, 1, 2, 7, 8, 9, 15, 16, 17, 33, 257} {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			array := []byte{'['}
			for i := 0; i < n; i++ {
				if i != 0 {
					array = append(array, ',')
				}
				array = append(array, elements[i%len(elements)]...)
			}
			array = append(array, ']')
			// Sibling arrays exercise separate scratch buffers in recursive calls.
			doc := append([]byte{'['}, array...)
			doc = append(doc, ',')
			doc = append(doc, array...)
			doc = append(doc, ']')
			for _, mode := range []struct {
				name            string
				compact, number bool
				read            func([]byte, int) (any, int, error)
			}{
				{"plain", false, false, DecodeValue},
				{"compact", true, false, DecodeValueCompact},
				{"number", false, true, DecodeValueNumber},
				{"number_compact", true, true, DecodeValueNumberCompact},
			} {
				t.Run(mode.name, func(t *testing.T) {
					input := doc
					if !mode.compact {
						var pretty bytes.Buffer
						if err := stdjson.Indent(&pretty, doc, "", "  "); err != nil {
							t.Fatal(err)
						}
						input = pretty.Bytes()
					}
					oracle := stdjson.NewDecoder(bytes.NewReader(input))
					if mode.number {
						oracle.UseNumber()
					}
					var want any
					if err := oracle.Decode(&want); err != nil {
						t.Fatal(err)
					}
					const start = 3
					padded := append([]byte("   "), input...)
					padded = append(padded, ",false]"...)
					got, end, err := mode.read(padded, start)
					if err != nil || end != start+len(input) {
						t.Fatalf("end=%d err=%v, want end=%d", end, err, start+len(input))
					}
					// Output owns its strings and backings, including values buffered
					// while another array is decoded recursively.
					clear(padded)
					if !reflect.DeepEqual(got, want) {
						t.Fatalf("decoded value differs from encoding/json")
					}
					if n > 0 {
						siblings := got.([]any)
						siblings[0].([]any)[0] = "changed"
						if !reflect.DeepEqual(siblings[1], want.([]any)[1]) {
							t.Fatal("sibling arrays share a backing")
						}
					}
				})
			}
		})
	}
}

func TestDecodeValueArrayTruncation(t *testing.T) {
	for _, n := range []int{7, 8, 9, 16, 17} {
		doc := []byte("[" + string(bytes.Repeat([]byte(`{"x":[1,"a\nb"]},`), n-1)) + `null]`)
		for _, read := range []func([]byte, int) (any, int, error){
			DecodeValue, DecodeValueCompact, DecodeValueNumber, DecodeValueNumberCompact,
		} {
			for cut := 0; cut < len(doc); cut++ {
				_, end, err := read(doc[:cut], 0)
				if err == nil || end < 0 || end > cut {
					t.Fatalf("n=%d prefix=%d: end=%d err=%v", n, cut, end, err)
				}
			}
			bad := append(append([]byte(nil), doc[:len(doc)-1]...), ',', ']')
			if _, _, err := read(bad, 0); err != ErrInvalidJSON {
				t.Fatalf("n=%d trailing comma: %v", n, err)
			}
		}
	}
}
