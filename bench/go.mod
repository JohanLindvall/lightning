// The bench module is separate from the main lightning module so that the
// benchmark-only dependency on mailru/easyjson does not leak into the main
// module's dependency graph. It depends on lightning via a local replace.
module github.com/JohanLindvall/lightning/bench

go 1.25.0

require (
	github.com/JohanLindvall/lightning v0.0.0
	github.com/buger/jsonparser v1.2.0
	github.com/mailru/easyjson v0.9.2
)

require (
	github.com/josharian/intern v1.0.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
)

replace github.com/JohanLindvall/lightning => ../
