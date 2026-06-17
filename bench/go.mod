// The bench module is separate from the main lightning module so that the
// benchmark-only dependencies (mailru/easyjson, bytedance/sonic) do not leak
// into the main module's dependency graph. It depends on lightning via a local
// replace.
module github.com/JohanLindvall/lightning/bench

go 1.26

require (
	github.com/JohanLindvall/lightning v0.0.0
	github.com/buger/jsonparser v1.2.0
	github.com/bytedance/sonic v1.15.2
	github.com/go-json-experiment/json v0.0.0-20260601182631-00ed12fed2a6
	github.com/goccy/go-json v0.10.6
	github.com/mailru/easyjson v0.9.2
	github.com/tidwall/gjson v1.19.0
)

require (
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic/loader v0.5.1 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/sys v0.45.0 // indirect
)

replace github.com/JohanLindvall/lightning => ../
