# Lightning documentation

[Home](../README.md) · [Quick start](../README.md#quick-start) · [API reference](https://pkg.go.dev/github.com/JohanLindvall/lightning/pkg/json) · [Contributing](../CONTRIBUTING.md)

Start with the quick start to generate a decoder, or use `pkg/json` directly when
you need only a few fields from a document.

## User guides

| Guide | What it covers |
|---|---|
| [Generator](generator.md) | Installation, supported types, root selection, reuse, embedding, and recursive schemas |
| [Tags and directives](options.md) | `nocopy`, `destructive`, alternate names, `lax`, `unwrap`, `number`, custom readers, and root directives |
| [Reading and parsing JSON](json.md) | Lookups, iteration, dynamic values, validity, errors, strings, and numbers |
| [Streaming](streaming.md) | Bounded buffering, callback lifetimes, and forward-only paths |
| [Editing JSON](editing.md) | `Set`, `SetMany`, `SetPaths`, `StripDefaults`, and checked edits |
| [Compatibility](compatibility.md) | Differences from `encoding/json`, input validation, and ownership contracts |
| [Performance and CPUs](performance.md) | Reproducing benchmarks, CPU features, and portable fallbacks |

The [runnable examples](https://pkg.go.dev/github.com/JohanLindvall/lightning/example)
work in your browser. Their [source](../example) includes a schema and a committed
generated decoder.

## API at a glance

| Goal | API |
|---|---|
| Decode into known Go types | [Generated `UnmarshalJSON`](generator.md) |
| Read a nested value | [`Lookup` or `Get`](json.md#key-lookups) |
| Read several values in one pass | [`GetMany` or `GetPaths`](json.md#key-lookups) |
| Visit members or elements | [`ObjectEach`, `ArrayEach`, `ArrayEachIndex`](json.md#key-lookups) |
| Walk an `io.Reader` | [`NewReader`](streaming.md) |
| Decode without a schema | [`DecodeAny` or `DecodeAnyNumber`](json.md#decoding-into-any) |
| Check Lightning's JSON grammar | [`Valid`](json.md#checking-validity) |
| Edit untrusted input | [`…Checked` edit functions](editing.md#checked-edits) |

`pkg/json` is the public API. `pkg/unstable` supports generated decoders and is
not a stable API for application code.

## Developer guides

Read [CONTRIBUTING.md](../CONTRIBUTING.md) for setup and review expectations.
[AGENTS.md](../AGENTS.md) is the task-based entry point for coding agents.

| Guide | Read when… |
|---|---|
| [Repository layout](development/layout.md) | Finding source files, tests, or benchmark cases |
| [Workflow and conventions](development/workflow.md) | Building, formatting, linting, or preparing an assembly change |
| [Testing](development/testing.md) | Adding regression coverage or changing a contract |
| [Benchmarking](development/benchmarking.md) | Measuring a change or designing a benchmark |
| [Generator internals](development/codegen.md) | Changing templates, type discovery, imports, or generated reads |
| [Runtime internals](development/runtime.md) | Changing numbers, strings, allocation, whitespace, or validation |
| [SIMD and assembly](development/simd.md) | Changing scanners, kernels, or CPU dispatch |
| [Toolkit internals](development/toolkit.md) | Changing lookup, editing, streaming, or escaping |

## Help

Use [GitHub issues](https://github.com/JohanLindvall/lightning/issues/new/choose)
for questions, bug reports, and feature requests. A small Go example and JSON
input make it easier to help.

## Credits

Several of the hot-path techniques are borrowed from prior art:

- **[simdjson](https://github.com/simdjson/simdjson)** (Geoff Langdale and Daniel
  Lemire) and its Go port **[minio/simdjson-go](https://github.com/minio/simdjson-go)** —
  the SWAR "parse four digits at once" bit trick used in the float and integer
  scanners, the nibble-table shuffle classification that the arm64 NEON
  `indexStructural` uses to find structural bytes, and the branchless
  escaped-quote detection + quote-mask prefix-XOR that builds the *inside-string*
  bitmask for the whole-container skip.
- **[sonic-rs](https://github.com/cloudwego/sonic-rs)** (ByteDance/CloudWeGo) and
  the **[JSONSki](https://github.com/AutomataLab/JSONSki)** paper — the SIMD
  container-skip used by `SkipValue`/`Get`: balance a container's brackets over
  the inside-string bitmask, so string contents are absorbed into the bulk scan
  rather than re-scanned per string. (The arm64 `maskBlock` movemask uses the
  well-known byte-LSB gather, `(x & 0x01…01) * 0x0102…80 >> 56`.)
- The float parser in `scanFloat` layers two published algorithms: William
  Clinger's fast path for exactly-representable values, then the
  **[Eisel–Lemire](https://github.com/fastfloat/fast_float)** algorithm (Michael
  Eisel and Daniel Lemire) for the rest — bit-for-bit identical to
  `strconv.ParseFloat` on the values it accepts.
- The benchmark corpus draws on
  [minio/simdjson-go](https://github.com/minio/simdjson-go/tree/master/testdata)
  and [go-json-experiment/jsonbench](https://github.com/go-json-experiment/jsonbench)
  test data, and the comparison suite measures against
  [encoding/json](https://pkg.go.dev/encoding/json),
  [mailru/easyjson](https://github.com/mailru/easyjson),
  [bytedance/sonic](https://github.com/bytedance/sonic),
  [goccy/go-json](https://github.com/goccy/go-json), and
  [go-json-experiment/json](https://github.com/go-json-experiment/json) (json/v2).
