# Lightning ⚡

[![CI](https://github.com/JohanLindvall/lightning/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/JohanLindvall/lightning/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/JohanLindvall/lightning.svg)](https://pkg.go.dev/github.com/JohanLindvall/lightning/pkg/json)
[![Go version](https://img.shields.io/github/go-mod/go-version/JohanLindvall/lightning)](go.mod)
[![Version](https://img.shields.io/github/v/tag/JohanLindvall/lightning?sort=semver&label=version)](https://github.com/JohanLindvall/lightning/tags)
[![License: MIT](https://img.shields.io/github/license/JohanLindvall/lightning)](LICENSE)

**Fast JSON decoding for Go, generated from your types.** Lightning writes
reflection-free `UnmarshalJSON` methods and provides a toolkit for reading,
streaming, editing, and validating JSON without a schema.

[Quick start](#quick-start) · [Documentation](docs/README.md) ·
[Run an example in your browser](https://pkg.go.dev/github.com/JohanLindvall/lightning/example#example-package) ·
[Benchmarks](#benchmarks) · [Contributing](CONTRIBUTING.md)

## Why Lightning?

- **Keep your Go types and call sites.** Generated methods implement
  `json.Unmarshaler`; `encoding/json.Unmarshal` calls them automatically.
  Call `v.UnmarshalJSON(data)` directly to skip that entry point.
- **Control allocations.** Reuse slices, maps, and pointers across decodes.
  Opt into string aliasing or in-place unescaping when you own the input buffer.
- **Use the CPU you have.** Runtime dispatch selects SIMD scanners on amd64 and
  arm64, with portable fallbacks on other architectures.
- **Read only what you need.** The public `pkg/json` toolkit looks up paths,
  walks arrays, streams large documents, and edits raw JSON without generation.

Lightning has deliberate differences from `encoding/json`, including exact key
matching, null handling, and numeric overflow. Read the
[compatibility guide](docs/compatibility.md) before migrating.

## Quick start

Requires **Go 1.25 or newer**. In a new directory:

```sh
go mod init example.com/events
go get github.com/JohanLindvall/lightning@latest
```

Save this as `event.go`:

```go
package main

import (
    "fmt"
    "log"
    "time"
)

//go:generate go run github.com/JohanLindvall/lightning $GOFILE

type Event struct {
    ID   int64     `json:"id"`
    Kind string    `json:"kind"`
    At   time.Time `json:"at"`
    Tags []string  `json:"tags"`
}

func main() {
    data := []byte(`{"id":42,"kind":"deploy","at":"2026-10-01T12:00:00Z","tags":["prod"]}`)
    var e Event
    if err := e.UnmarshalJSON(data); err != nil {
        log.Fatal(err)
    }
    fmt.Println(e.ID, e.Kind, e.At.Format(time.RFC3339), e.Tags)
}
```

Generate the decoder and run it:

```sh
go generate ./...
go mod tidy
go run .
# 42 deploy 2026-10-01T12:00:00Z [prod]
```

Generation writes `event_unmarshal.go` beside your schema. The directive uses the
version in your `go.mod`, keeping the generator and its runtime together. Run
`go generate ./...` again after changing the schema or upgrading Lightning.

See the [generator guide](docs/generator.md) for installation options, supported
types, reuse, and custom unmarshalers, or [tags and directives](docs/options.md)
for `nocopy`, `strict`, `emptyslice`, `lax`, and other options.

## Use the toolkit without generation

Import `github.com/JohanLindvall/lightning/pkg/json` to work directly with bytes.
Install it in your module:

```sh
go get github.com/JohanLindvall/lightning/pkg/json@latest
```

This standalone example looks up a nested string:

```go
package main

import (
    "fmt"
    "log"

    json "github.com/JohanLindvall/lightning/pkg/json"
)

func main() {
    data := []byte(`{"user":{"name":"Ada"}}`)
    raw, err := json.Lookup(data, "user", "name")
    if err != nil {
        log.Fatal(err)
    }
    name, err := json.String(raw)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name) // Ada
}
```

Lookup results alias the input; keep it unchanged while using them. Use the
[checked edit functions](docs/editing.md#checked-edits) when editing untrusted
input. `Valid` follows Lightning's documented grammar; it is not a substitute for
`encoding/json.Valid` when you need exactly the standard library's acceptance rules.

## Documentation

<!-- Stable fragment links lead to the detailed guides below. -->
<a id="features"></a>
<a id="api-at-a-glance"></a>
<a id="the-generator"></a>
<a id="the-toolkit--pkgjson"></a>
<a id="installation"></a>
<a id="how-it-works"></a>
<a id="supported-types"></a>
<a id="reusing-a-decode-target"></a>
<a id="differences-from-encodingjson"></a>
<a id="root-types"></a>
<a id="declarations-that-get-no-method"></a>
<a id="the-nocopy-tag-option"></a>
<a id="the-lightningdestructive-directive"></a>
<a id="alternate-field-names"></a>
<a id="the-lax-tag-option"></a>
<a id="the-unwrap-tag-option"></a>
<a id="the-number-tag-option"></a>
<a id="the-reader-tag-option"></a>
<a id="comment-directives"></a>
<a id="lightningemptyslice"></a>
<a id="lightningstrict"></a>
<a id="lightningcompact"></a>
<a id="lightningnocopy"></a>
<a id="lightningarena"></a>
<a id="generated-function-names"></a>
<a id="key-lookups"></a>
<a id="streaming"></a>
<a id="decoding-into-any"></a>
<a id="checking-validity"></a>
<a id="errors"></a>
<a id="string-escaping-and-unescaping"></a>
<a id="number-parsing"></a>
<a id="stripping-default-fields"></a>
<a id="setting-a-value"></a>
<a id="checked-edits"></a>
<a id="simd-scanning"></a>
<a id="cpu-requirements"></a>
<a id="layout"></a>
<a id="limits-and-untrusted-input"></a>
<a id="credits"></a>

| I want to… | Start here |
|---|---|
| Generate a decoder or reuse a target | [Generator guide](docs/generator.md) |
| Configure tags and directives | [Options reference](docs/options.md) |
| Look up values, validate, escape, or parse | [JSON toolkit](docs/json.md) · [API reference](https://pkg.go.dev/github.com/JohanLindvall/lightning/pkg/json) |
| Process a large response as it arrives | [Streaming](docs/streaming.md) |
| Set values or remove defaults | [Editing JSON](docs/editing.md) |
| Check migration and buffer ownership rules | [Compatibility and input contracts](docs/compatibility.md) |
| Understand CPU support or measure performance | [Performance guide](docs/performance.md) |
| Work on Lightning itself | [Contributing](CONTRIBUTING.md) · [Repository layout](docs/development/layout.md) |

## Benchmarks

The comparison suite covers real-world documents and synthetic workloads, with
generated decoders, `DecodeAny`, `encoding/json`, easyjson, sonic, goccy/go-json,
and json/v2. It reports latency, bytes allocated, and allocation counts.

| Architecture | Decoder comparisons | Toolkit and runtime benchmarks |
|---|---|---|
| amd64 | [Results](bench/results_amd64.md) | [Results](bench/pkg_results_amd64.md) |
| arm64 | [Results](bench/results_arm64.md) | [Results](bench/pkg_results_arm64.md) |

These are single-run CI snapshots, not a guarantee for your workload. Check the
CPU and toolchain headers, compare same-run ratios, and measure your own data.
The [performance guide](docs/performance.md) explains how to reproduce the suite.

## Help and contributing

Questions, bug reports, documentation fixes, tests, and performance work are
welcome. [Open an issue](https://github.com/JohanLindvall/lightning/issues/new/choose)
with a small example, or follow [CONTRIBUTING.md](CONTRIBUTING.md) to run the project
and prepare a pull request. You can try the
[runnable examples](https://pkg.go.dev/github.com/JohanLindvall/lightning/example)
without installing anything.

If Lightning helps your project, a star helps other Go developers discover it.

## License

[MIT](LICENSE) © Johan Lindvall. See the [credits](docs/README.md#credits) for
algorithms and benchmark sources.
