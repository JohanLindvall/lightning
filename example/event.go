// Package example is the README's opening example: a schema, the decoder the
// generator wrote for it (event_unmarshal.go, committed so the package builds
// anywhere), and runnable examples that check it against encoding/json. On
// pkg.go.dev they run in the browser, with nothing to install.
package example

import "time"

// Your own module would run the published generator, as the README shows:
//
//	//go:generate go run github.com/JohanLindvall/lightning@latest $GOFILE
//
// Inside this repository the generator is the local main package, so the
// directive below names it without a version. TestExampleDecoderIsCurrent fails
// whenever the committed decoder no longer matches what the generator writes.

//go:generate go run github.com/JohanLindvall/lightning $GOFILE

// Event is the struct from the README.
type Event struct {
	ID   int64     `json:"id"`
	Kind string    `json:"kind"`
	At   time.Time `json:"at"`
	Tags []string  `json:"tags"`
}
