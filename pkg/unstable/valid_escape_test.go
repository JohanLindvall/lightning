package unstable

import (
	"math/rand"
	"strings"
	"testing"
)

// Keep the original escape grammar and offsets as a scalar oracle. It uses
// neither the decoder's lookup tables nor the optimized literal-run scanner.
func strictStringEscapedReference(data []byte, i int) (int, error) {
	for {
		i++ // backslash
		if i >= len(data) {
			return i, ErrTruncated
		}
		switch data[i] {
		case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
			i++
		case 'u':
			if len(data)-i-1 < 4 {
				return i, ErrTruncated
			}
			for _, c := range data[i+1 : i+5] {
				validHex := c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
				if !validHex {
					return i, ErrBadUnicode
				}
			}
			i += 5
		default:
			return i, ErrBadEscape
		}
		for i < len(data) && data[i] != '"' && data[i] != '\\' {
			i++
		}
		if i == len(data) {
			return i, ErrTruncated
		}
		if data[i] == '"' {
			return i + 1, nil
		}
	}
}

func checkStrictEscaped(t *testing.T, data []byte, start int) {
	t.Helper()
	wantEnd, wantErr := strictStringEscapedReference(data, start)
	end, err := strictStringEscaped(data, start)
	if end != wantEnd || err != wantErr {
		t.Fatalf("%q at %d: (%d, %v), want (%d, %v)", data, start, end, err, wantEnd, wantErr)
	}
}

func TestStrictStringEscapedByteClasses(t *testing.T) {
	for _, start := range []int{0, 1, 15, 16, 31, 32, 63, 64, 129} {
		prefix := strings.Repeat("x", start)
		for c := 0; c < 256; c++ {
			checkStrictEscaped(t, []byte(prefix+`\`+string([]byte{byte(c)})+`"tail`), start)
			for digit := 0; digit < 4; digit++ {
				data := []byte(prefix + `\u0aF9\n"tail`)
				data[start+2+digit] = byte(c)
				checkStrictEscaped(t, data, start)
			}
		}
	}
}

func TestStrictStringEscapedBoundaries(t *testing.T) {
	for _, n := range []int{0, 1, 15, 16, 31, 32, 63, 64, 127, 128} {
		for _, body := range []string{
			`\n\t\r\b\f\/\\\"`,
			`\u0000\u00aF\u4E16\u754c`,
			`\ud83d\ude00\ud800\udfff`,
			`\n` + strings.Repeat("x", n) + `\u1234\n`,
			`\n` + strings.Repeat("x", n) + `\u12Z4`,
			`\n` + string([]byte{0, 0x1f, 0x80, 0xff}) + `\t`,
		} {
			data := []byte(strings.Repeat("p", n) + body + `"tail`)
			for end := n + 1; end <= len(data); end++ {
				checkStrictEscaped(t, data[:end], n)
			}
		}
	}
	// Include malformed bodies with arbitrary bytes, not just valid JSON.
	rng := rand.New(rand.NewSource(42))
	for range 20000 {
		data := make([]byte, 1+rng.Intn(256))
		_, _ = rng.Read(data)
		data[0] = '\\'
		checkStrictEscaped(t, data, 0)
	}
}

func FuzzStrictStringEscaped(f *testing.F) {
	for _, body := range []string{`n"`, `u1234"`, `ud83d\ude00"`, `u12`, `q`, `n\t\u0000"`, `n` + strings.Repeat("x", 64) + `\uFFFF"`} {
		f.Add([]byte(body), byte(0))
	}
	f.Fuzz(func(t *testing.T, body []byte, offset byte) {
		start := int(offset)
		data := make([]byte, start+1, start+1+len(body))
		data[start] = '\\'
		data = append(data, body...)
		checkStrictEscaped(t, data, start)
	})
}
