package json

import (
	"strconv"
	"strings"
	"testing"
)

// skipHeavyObject builds {"f0":{record},"f1":{record},...,"target":42}: the
// target key is last, so Get must skip every preceding member. The values are
// small nested objects/arrays/strings — the dense-container shape where the
// in-string-mask skip in unstable.SkipValue beats the per-string indexStructural
// path. This is the canonical Get-on-a-big-document workload.
func skipHeavyObject(n int) []byte {
	var b strings.Builder
	b.WriteByte('{')
	for i := 0; i < n; i++ {
		b.WriteString(`"f`)
		b.WriteString(strconv.Itoa(i))
		b.WriteString(`":{"id":1234567,"name":"a moderately long name value","tags":[1,2,3,4],"active":true},`)
	}
	b.WriteString(`"target":42}`)
	return []byte(b.String())
}

func BenchmarkGetSkipHeavy(b *testing.B) {
	data := skipHeavyObject(500)
	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v, _, err := Get(data, "target")
		if err != nil || string(v) != "42" {
			b.Fatalf("got %q err %v", v, err)
		}
	}
}
