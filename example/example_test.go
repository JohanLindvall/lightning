package example_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/JohanLindvall/lightning/example"
)

// The document both examples decode. The "debug" member has no field in Event,
// so it is skipped, exactly as encoding/json skips it.
var data = []byte(`{
	"id": 42,
	"kind": "deploy",
	"at": "2026-10-06T12:00:00Z",
	"tags": ["prod", "eu-west"],
	"debug": {"trace": [1, 2, 3]}
}`)

// eventStd has Event's fields and none of its methods, so encoding/json decodes
// it by reflection instead of handing the bytes to the generated UnmarshalJSON.
// It is the baseline the generated decoder is checked against.
type eventStd example.Event

// Example decodes a document through the generated method and checks the result
// against encoding/json.
func Example() {
	var e example.Event
	if err := e.UnmarshalJSON(data); err != nil {
		panic(err)
	}
	fmt.Println(e.ID, e.Kind, e.At.Format(time.RFC3339), e.Tags)

	// The generated method makes Event a json.Unmarshaler, so json.Unmarshal
	// calls it with no change at the call site.
	var viaUnmarshal example.Event
	if err := json.Unmarshal(data, &viaUnmarshal); err != nil {
		panic(err)
	}

	var std eventStd
	if err := json.Unmarshal(data, &std); err != nil {
		panic(err)
	}
	got, _ := json.Marshal(e)
	again, _ := json.Marshal(viaUnmarshal)
	want, _ := json.Marshal(std)
	fmt.Println("same as encoding/json:", string(got) == string(want) && string(again) == string(want))

	// Output:
	// 42 deploy 2026-10-06T12:00:00Z [prod eu-west]
	// same as encoding/json: true
}

// Example_allocations counts the heap allocations of one decode. It has no
// Output section — the standard library's count changes between Go versions —
// so go test only compiles it; run it on pkg.go.dev to see the numbers.
func Example_allocations() {
	lightning := testing.AllocsPerRun(100, func() {
		var e example.Event
		_ = e.UnmarshalJSON(data)
	})
	stdlib := testing.AllocsPerRun(100, func() {
		var e eventStd
		_ = json.Unmarshal(data, &e)
	})
	fmt.Printf("allocations per decode: lightning %.0f, encoding/json %.0f\n", lightning, stdlib)
}
