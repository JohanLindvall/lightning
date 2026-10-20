---
name: Bug report
about: Report incorrect decoding, generation, or toolkit behavior.
---

## What happened?

Describe the actual behavior and what you expected.

## Reproduction

Include the smallest Go types, JSON input, and generator command or toolkit call
that show the problem. Include the error or generated-code compiler output.

```go
// Minimal example
```

```json
{}
```

## Environment

- Lightning version:
- `go version`:
- Operating system and architecture:
- CPU model (for SIMD or performance issues):

## Additional context

For a comparison with `encoding/json`, use a type without a generated
`UnmarshalJSON` method so the baseline does not call Lightning.
For a performance issue, include the benchmark command and baseline results.
