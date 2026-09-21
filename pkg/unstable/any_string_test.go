package unstable

import (
	"strings"
	"testing"
)

func TestDecodeValueStringMatchesReader(t *testing.T) {
	readers := []struct {
		name string
		read func([]byte, int) (any, int, error)
	}{
		{"normal", DecodeValue},
		{"compact", DecodeValueCompact},
		{"number", DecodeValueNumber},
		{"number_compact", DecodeValueNumberCompact},
	}
	for _, reader := range readers {
		t.Run(reader.name, func(t *testing.T) {
			for _, n := range []int{0, 1, 15, 16, 31, 32, 63, 64, 128, 1024} {
				for _, body := range []string{
					``, `\n`, `\u4e16\u754c`, `\ud83d\ude00`, `\ud800`, `\"\\\/`,
					`\u12`, `\uZZZZ`, `\q`, string([]byte{0, 0x80, 0xff}),
				} {
					doc := []byte(`prefix"` + strings.Repeat("x", n) + body + `"tail`)
					const start = len("prefix")
					for limit := start + 1; limit <= len(doc); limit++ {
						in := doc[:limit]
						want, wantEnd, wantErr := ReadStringOrNull(in, start)
						got, end, err := reader.read(in, start)
						if got != want || end != wantEnd || err != wantErr {
							t.Fatalf("%q: (%#v, %d, %v), want (%q, %d, %v)", in, got, end, err, want, wantEnd, wantErr)
						}
					}
				}
			}
			for _, body := range []string{strings.Repeat("clean", 64), strings.Repeat("clean", 64) + `\nend`} {
				in := []byte(`"` + body + `"`)
				got, _, err := reader.read(in, 0)
				if err != nil {
					t.Fatal(err)
				}
				want := strings.Clone(got.(string))
				clear(in)
				if got != want {
					t.Fatal("decoded string aliases input")
				}
			}
		})
	}
}
