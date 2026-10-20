# Streaming JSON

[Home](../README.md) · [Documentation](README.md)

Walk an `io.Reader` through a bounded buffer with `github.com/JohanLindvall/lightning/pkg/json`.

## Streaming

The [in-memory walkers](json.md#key-lookups) take a complete `[]byte`, so a large answer read over the
network is buffered whole — `io.ReadAll` and then `ArrayEach` — before the
first byte is scanned. A `Reader` walks the document **as it arrives**, through
a bounded buffer:

```go
r := json.NewReader(resp.Body)
err := r.ArrayEach(func(series []byte) error {
    // series is one element of the matrix, complete, in the reader's buffer
    return json.ArrayEach(series, handlePoint, "values")
}, "data", "result")
```

- `NewReader(r io.Reader, opts ...ReaderOption) *Reader` — options are
  `WithBufferSize(n)` (the initial buffer, 64 KiB) and `WithMaxElement(n)` (the
  point at which one value is `ErrElementTooLarge` rather than a bigger
  allocation, 64 MiB).
- `(*Reader).ArrayEach(fn, keys...)` and `(*Reader).ObjectEach(fn, keys...)` —
  the two walkers, with the in-memory contract: the same path descent, the same
  sentinels, `ErrStop` from the callback to end a walk early, and a `null`
  where the container would be walked as an empty one.
- `(*Reader).Get(keys...)` — read until the value at the path is complete and
  return it.
- `(*Reader).Elements(keys...)` — `ArrayEach` as a range-over-func iterator,
  yielding `(value, nil)` per element and at most one `(nil, err)`.
- `(*Reader).Reset(io.Reader)` — start another document on the same reader, and
  the same buffer. A client making the same query repeatedly pays for the
  buffer once.
- `(*Reader).Consumed()` and `(*Reader).Buffered()` — how many bytes were read
  from the source, and the ones read past the end of the document; together
  they let the caller hand the rest of the stream on.

**The one contract difference — a value is valid only until the callback
returns.** The in-memory walkers hand back a window onto the caller's own
`data`, valid as long as that slice lives; the stream reader's window points
into its buffer, and the next refill may move or overwrite it. A callback that
keeps anything copies it: `UnescapeStringCopy` for a string body,
`strings.Clone` for a key, `append([]byte(nil), v...)` for raw bytes. This is
the one rule that a test, not the compiler, has to enforce for you.

**What it bounds.** The buffer holds at most one complete value plus what has
been read ahead of it. It grows only when a single value does not fit, and
never past `WithMaxElement`. Values the walk does not want — the members before
the one a key path names — are streamed past without ever being buffered whole,
so a 100 MB sibling costs the read and nothing else.

Reuse a reader with `Reset` to amortize buffer allocation. Streaming is useful
when buffering the whole document is expensive; for small documents already in
memory, use the byte-slice walkers. Measure both with your workload.

**What does not stream.** `Set`, `SetMany`, `SetPaths`, `StripDefaults`,
`Valid` and `DecodeAny` need the whole value by nature and stay `[]byte` APIs,
as does a generated `UnmarshalJSON`.

**Walking on.** A Reader takes one document, but not one path: a later call
continues from where the previous one stopped, forward only. The members
after the last value read are searched for the new path's next key; a path
that shares open objects with the previous one goes on inside them, one
that leaves them skips what remains and continues in the parent; a walk a
callback ended with `ErrStop` is finished first. A key the walk has already
passed — or the inside of a value read whole — answers `ErrKeyNotFound`,
and the search for it runs the object to its close, so a miss spends what
it passed over; after a keyless call has walked the root value nothing is
left to find. A
Prometheus envelope reads as it is written, in one pass:

```go
status, _ := r.Get("status")                       // "success"
kind, _ := r.Get("data", "resultType")             // "matrix"
_ = r.ArrayEach(series, "data", "result")          // the matrix, one element at a time
stats, _ := r.Get("data", "stats")                 // what follows the result inside data
warnings, _ := r.Get("warnings")                   // what follows data
```
