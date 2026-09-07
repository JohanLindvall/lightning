package json

import (
	"errors"
	"fmt"
	"io"
	"iter"

	"github.com/JohanLindvall/lightning/pkg/unstable"
)

// Reader walks ONE JSON document as it arrives from an io.Reader, through a
// buffer that holds a bounded part of it rather than the whole thing. It is
// for the answer that is large and whose useful working set is small — a
// Prometheus matrix of a thousand series, a table export, an event log — where
// io.ReadAll before the first ArrayEach costs the document's whole size in
// memory and delays every callback until the download ends.
//
// # The one contract difference
//
// The in-memory walkers hand fn a window onto the caller's own data, valid for
// as long as that slice lives. THE STREAM READER'S WINDOW IS VALID ONLY UNTIL
// fn RETURNS. It points into the reader's buffer, and the next refill may move
// or overwrite it. A callback that keeps anything must copy it —
// [UnescapeStringCopy] for a string body, strings.Clone for a key,
// append([]byte(nil), v...) for raw bytes. Everything else is the in-memory
// contract: the same path descent by keys, the same callback rule, the same
// sentinels, [ErrStop] to end a walk early, and a JSON null where a container
// would be walked as an empty one.
//
// # What it costs
//
// The buffer starts at 64 KiB, grows only when one value does not fit, and
// never past [WithMaxElement] (64 MiB by default), after which the walk fails
// with [ErrElementTooLarge] rather than allocating without bound. Values the
// walk does not want — the members before the one a key path names — are
// streamed past without being buffered whole, so a 100 MB sibling costs
// nothing but the read.
//
// A Reader is not safe for concurrent use and reads one document.
type Reader struct {
	r   io.Reader
	buf []byte
	// hold is the first byte the walk still needs: a value being assembled, or
	// the key of the member whose value is being assembled (fn is handed both
	// at once, so the key has to survive the value's refills). pos is the scan
	// cursor. Both are adjusted when fill compacts, so no caller may hold a
	// raw index across a call that can refill.
	hold, pos, end int
	maxBuf         int
	consumed       int64
	rerr           error // the reader's error, once it has given one
	sc             unstable.ValueScanner
}

// ErrElementTooLarge reports a value larger than the reader's buffer limit.
// The walk cannot continue past it: the value has to be buffered whole to be
// handed to a callback, and the limit is what says the reader would rather
// fail than allocate.
var ErrElementTooLarge = errors.New("json: element exceeds the buffer limit")

// ReaderOption configures a [Reader]; see [WithBufferSize] and
// [WithMaxElement].
type ReaderOption func(*readerConfig)

type readerConfig struct{ bufSize, maxElem int }

const (
	defaultReaderBuffer = 64 << 10
	defaultMaxElement   = 64 << 20
	// valueRetries is how many times a value that does not fit the buffered
	// bytes is refilled and re-offered to SkipValue before the resumable
	// scanner takes it; see value.
	valueRetries = 4
)

// WithBufferSize sets the initial buffer, 64 KiB by default. It is a starting
// point, not a limit: the buffer grows to fit one value, up to
// [WithMaxElement]. Sizing it near the largest expected value avoids the
// doubling copies; sizing it small does not change what the reader accepts.
func WithBufferSize(n int) ReaderOption {
	return func(c *readerConfig) {
		if n > 0 {
			c.bufSize = n
		}
	}
}

// WithMaxElement sets how large one value may be before the walk fails with
// [ErrElementTooLarge], 64 MiB by default. It bounds the reader's memory: a
// hostile or simply unexpected document cannot make it allocate more.
func WithMaxElement(n int) ReaderOption {
	return func(c *readerConfig) {
		if n > 0 {
			c.maxElem = n
		}
	}
}

// NewReader starts a document on r.
func NewReader(r io.Reader, opts ...ReaderOption) *Reader {
	c := readerConfig{bufSize: defaultReaderBuffer, maxElem: defaultMaxElement}
	for _, o := range opts {
		o(&c)
	}
	if c.bufSize > c.maxElem {
		c.bufSize = c.maxElem
	}
	return &Reader{r: r, buf: make([]byte, c.bufSize), maxBuf: c.maxElem}
}

// Reset starts a new document on rd, keeping the buffer this reader has
// already grown. A caller decoding a sequence of responses reuses one Reader
// rather than allocating a buffer per document — the whole point of a bounded
// buffer is that it is paid for once.
func (r *Reader) Reset(rd io.Reader) {
	r.r = rd
	r.hold, r.pos, r.end = 0, 0, 0
	r.consumed, r.rerr = 0, nil
}

// Consumed reports how many bytes have been READ from the underlying reader,
// which is ahead of the walk: the reader buffers in blocks and reads past the
// end of the document. A caller that means to hand the rest of the stream on
// takes [Reader.Buffered] first — those bytes were read but not walked.
func (r *Reader) Consumed() int64 { return r.consumed }

// Buffered returns the bytes read from the underlying reader that the walk has
// not consumed, valid until the next call on r. With the reader itself it is
// the remainder of the stream: io.MultiReader(bytes.NewReader(clone), src).
func (r *Reader) Buffered() []byte { return r.buf[r.pos:r.end] }

// ArrayEach calls fn for every element of the array reached by the path keys
// (the document's root array when there are none), in order.
//
// It is [ArrayEach] over a stream, with that function's contract in every
// respect but one: the value handed to fn is valid ONLY UNTIL fn RETURNS (see
// [Reader]). A null where the array would be is an array with no elements, fn
// returning [ErrStop] ends the walk with a nil result, and the sentinels are
// the same — with [ErrElementTooLarge] added for an element past the limit and
// the underlying reader's error wrapped when it fails.
func (r *Reader) ArrayEach(fn func(value []byte) error, keys ...string) error {
	if err := r.enter(keys); err != nil {
		return err
	}
	c, err := r.space()
	if err != nil {
		return err
	}
	if c != '[' {
		return r.emptyOrExpect(unstable.ErrExpectArray)
	}
	r.pos++
	if c, err = r.space(); err != nil {
		return err
	}
	if c == ']' {
		r.pos++
		return nil
	}
	for {
		r.hold = r.pos
		// value's fast path, written out: the element is already buffered in
		// full except where a refill landed in the middle of it, and the call
		// and frame it saves are a tenth of this walk. See value for what the
		// number test is about.
		end, err := unstable.SkipValue(r.buf[:r.end], r.pos)
		if err != nil || (end == r.end && r.rerr == nil && startsNumber(r.buf[r.pos])) {
			if end, err = r.value(); err != nil {
				return err
			}
		}
		if err := fn(r.buf[r.pos:end]); err != nil {
			if errors.Is(err, ErrStop) {
				return nil
			}
			return err
		}
		r.hold, r.pos = end, end
		done, err2 := r.afterElement(']')
		if done || err2 != nil {
			return err2
		}
	}
}

// ObjectEach calls fn for every member of the object reached by the path keys
// (the document's root object when there are none), in order.
//
// It is [ObjectEach] over a stream; see [Reader.ArrayEach] for the one
// difference. BOTH the key and the value handed to fn are valid only until fn
// returns — the key is a window onto the buffer too when it holds no escape.
func (r *Reader) ObjectEach(fn func(key string, value []byte) error, keys ...string) error {
	if err := r.enter(keys); err != nil {
		return err
	}
	c, err := r.space()
	if err != nil {
		return err
	}
	if c != '{' {
		return r.emptyOrExpect(unstable.ErrExpectObject)
	}
	r.pos++
	for {
		c, err := r.space()
		if err != nil {
			return err
		}
		if c == '}' {
			r.pos++
			return nil
		}
		// fn is handed the key and the value together, so the key has to
		// outlive the scan of the value — which may refill several times. The
		// hold stays at the key's first byte so those bytes are never dropped,
		// and the key STRING is decoded only once the value is complete: a
		// compaction moves the bytes within the buffer, and a string made
		// before it would still point at the offset they left.
		if c != '"' {
			return unstable.ErrInvalidJSON
		}
		r.hold = r.pos
		kend, err := r.value()
		if err != nil {
			return err
		}
		r.pos = kend
		if err := r.colon(); err != nil {
			return err
		}
		end, err := unstable.SkipValue(r.buf[:r.end], r.pos)
		if err != nil || (end == r.end && r.rerr == nil && startsNumber(r.buf[r.pos])) {
			if end, err = r.value(); err != nil {
				return err
			}
		}
		// pos and hold are read AFTER value, for the same reason.
		key, _, err := unstable.ReadKey(r.buf[:r.end], r.hold)
		if err != nil {
			return err
		}
		if err := fn(key, r.buf[r.pos:end]); err != nil {
			if errors.Is(err, ErrStop) {
				return nil
			}
			return err
		}
		r.hold, r.pos = end, end
		done, err := r.afterElement('}')
		if done || err != nil {
			return err
		}
	}
}

// Get reads until the value at the path keys is complete and returns it. The
// result is valid until the next call on r; keep it with
// append([]byte(nil), v...).
//
// Everything the path passes over on the way is streamed past rather than
// buffered, so a value late in a large object costs the reads and not the
// memory. A missing key is ErrKeyNotFound.
func (r *Reader) Get(keys ...string) ([]byte, error) {
	if err := r.enter(keys); err != nil {
		return nil, err
	}
	if _, err := r.space(); err != nil {
		return nil, err
	}
	r.hold = r.pos
	end, err := r.value()
	if err != nil {
		return nil, err
	}
	v := r.buf[r.pos:end]
	r.hold, r.pos = end, end
	return v, nil
}

// Elements is [Reader.ArrayEach] as a range-over-func iterator: it yields each
// element, and then at most one (nil, err) if the walk fails. The value's
// lifetime is the loop body — the same rule as the callback's, and for the
// same reason. Breaking out of the loop ends the walk cleanly.
func (r *Reader) Elements(keys ...string) iter.Seq2[[]byte, error] {
	return func(yield func([]byte, error) bool) {
		err := r.ArrayEach(func(v []byte) error {
			if !yield(v, nil) {
				return ErrStop
			}
			return nil
		}, keys...)
		if err != nil {
			yield(nil, err)
		}
	}
}

// enter descends the path, leaving pos at the first byte of the value the last
// key names (or at the document's first value when there are none). Members
// passed over are streamed past, never buffered whole.
func (r *Reader) enter(keys []string) error {
	for _, key := range keys {
		c, err := r.space()
		if err != nil {
			return err
		}
		if c != '{' {
			return unstable.ErrExpectObject
		}
		r.pos++
		for {
			if c, err = r.space(); err != nil {
				return err
			}
			if c == '}' {
				return unstable.ErrKeyNotFound
			}
			r.hold = r.pos
			k, err := r.key()
			if err != nil {
				return err
			}
			// The key has been compared, so it need not outlive the value.
			match := k == key
			r.hold = r.pos
			if err := r.colon(); err != nil {
				return err
			}
			if match {
				break
			}
			if err := r.skip(); err != nil {
				return err
			}
			r.hold = r.pos
			done, err := r.afterElement('}')
			if err != nil {
				return err
			}
			if done {
				return unstable.ErrKeyNotFound
			}
		}
	}
	return nil
}

// afterElement reads the separator that follows a member or an element: the
// container's closer ends the walk, a comma continues it, anything else is
// malformed. It positions pos at the next member or element.
func (r *Reader) afterElement(closer byte) (bool, error) {
	// The separator is the very next byte on compact input and one space away
	// on pretty input, so both tests are written out here rather than reached
	// through space's frame: this runs once per element.
	c := byte(0)
	if r.pos < r.end && r.buf[r.pos] > ' ' {
		c = r.buf[r.pos]
	} else {
		var err error
		if c, err = r.space(); err != nil {
			return false, err
		}
	}
	switch c {
	case closer:
		r.pos++
		return true, nil
	case ',':
		// The next member or element starts after the separator's whitespace,
		// which is where the in-memory walkers leave their cursor too.
		r.pos++
		r.hold = r.pos
		if r.pos < r.end && r.buf[r.pos] > ' ' {
			return false, nil
		}
		_, err := r.space()
		return false, err
	}
	return false, unstable.ErrInvalidJSON
}

// colon reads the separator between a key and its value and leaves pos on the
// value's first byte.
func (r *Reader) colon() error {
	c, err := r.peek()
	if err != nil {
		return err
	}
	if c != ':' {
		return unstable.ErrExpectColon
	}
	r.pos++
	_, err = r.peek()
	return err
}

// key reads a member's key, which must be buffered whole (keys are small, and
// the hold the caller set keeps it alive across the value's refills). The
// returned string aliases the buffer unless the key holds an escape.
func (r *Reader) key() (string, error) {
	if r.pos >= r.end || r.buf[r.pos] != '"' {
		return "", unstable.ErrInvalidJSON
	}
	end, err := r.value()
	if err != nil {
		return "", err
	}
	k, _, err := unstable.ReadKey(r.buf[:end], r.pos)
	if err != nil {
		return "", err
	}
	r.pos = end
	return k, nil
}

// emptyOrExpect answers a value that is not the container the walker wanted: a
// null is that container with nothing in it (unmarshalling null into a map or
// a slice leaves it nil, and "values":null is a series with no points), and
// anything else — a misspelt literal included — is the caller's sentinel.
func (r *Reader) emptyOrExpect(want error) error {
	if r.buf[r.pos] != 'n' {
		return want
	}
	// The literal must end at a token boundary, so that a null reached by key
	// and followed by its object's '}' is one and nullx is not. Whitespace
	// here is this package's own, every byte <= 0x20; see isNullToken.
	for r.end-r.pos < 5 && r.rerr == nil {
		if err := r.fill(); err != nil && !errors.Is(err, io.EOF) {
			return r.readErr(err)
		}
	}
	if r.end-r.pos < 4 || string(r.buf[r.pos:r.pos+4]) != "null" {
		return want
	}
	if r.end-r.pos == 4 {
		r.pos += 4
		r.hold = r.pos
		return nil
	}
	switch c := r.buf[r.pos+4]; {
	case c <= ' ', c == ',', c == '}', c == ']':
		r.pos += 4
		r.hold = r.pos
		return nil
	}
	return want
}

// peek leaves pos on the next byte that is not whitespace, reading more when
// the buffer runs out, and returns it. It does NOT move the hold: this is what
// runs while a key is pinned, waiting for its value.
func (r *Reader) peek() (byte, error) {
	for {
		r.pos = unstable.SkipWS(r.buf[:r.end], r.pos)
		if r.pos < r.end {
			return r.buf[r.pos], nil
		}
		if err := r.fill(); err != nil {
			return 0, r.readErr(err)
		}
	}
}

// space is peek where nothing is pinned: the whitespace it walks over is
// dropped rather than held, so a run of it longer than the buffer costs
// nothing. Every point that does not need an earlier byte uses this one.
func (r *Reader) space() (byte, error) {
	for {
		r.pos = unstable.SkipWS(r.buf[:r.end], r.pos)
		r.hold = r.pos
		if r.pos < r.end {
			return r.buf[r.pos], nil
		}
		if err := r.fill(); err != nil {
			return 0, r.readErr(err)
		}
	}
}

// readErr turns a fill failure into the error a walk reports: a document that
// ends is truncated, a limit is itself, and anything else the underlying
// reader says is returned wrapped so errors.Is reaches it.
func (r *Reader) readErr(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, io.EOF):
		return unstable.ErrTruncated
	case errors.Is(err, ErrElementTooLarge):
		return err
	}
	return fmt.Errorf("json: read: %w", err)
}

// value leaves the value at pos complete in the buffer and returns the index
// just past it. pos is unchanged (fill may move both).
//
// The whole value is scanned by SkipValue when it is already buffered, which
// is what keeps the streaming walk's answer identical to the in-memory one on
// every document that fits; only a value that arrives across refills goes
// through the resumable scanner, and then it is scanned once, not once per
// refill.
func (r *Reader) value() (int, error) {
	// A value that straddles the end of the buffered bytes is refilled and
	// handed to SkipValue again a few times before the resumable scanner takes
	// over: SkipValue is the SIMD skip, several times faster than the
	// byte-state machine, so re-running it over a value of a few kilobytes
	// costs less than scanning those kilobytes the slow way once. What the
	// attempt limit buys is the guarantee the scanner exists for — total
	// re-scanning stays a small multiple of the value's size instead of the
	// O(n²) that "refill and retry from the first byte" is when a value
	// arrives in many small chunks (a one-byte reader is the pathological
	// case, and TestStreamHugeElementIsLinear is where it would hang).
	for attempt := 0; ; attempt++ {
		end, err := unstable.SkipValue(r.buf[:r.end], r.pos)
		switch {
		case err == nil && (end < r.end || r.rerr != nil || !startsNumber(r.buf[r.pos])):
			// A number token that ends exactly where the buffered bytes do may
			// still continue in the next chunk; every other kind is decided by
			// a byte SkipValue has already seen.
			return end, nil
		case err != nil && r.rerr != nil:
			// Nothing more can arrive to change the verdict, so this is the
			// in-memory walkers' error for the same bytes.
			return 0, r.wrap(err)
		}
		if attempt == valueRetries {
			break
		}
		if err := r.fill(); err != nil && !errors.Is(err, io.EOF) {
			return 0, r.readErr(err)
		}
	}
	r.sc.Reset()
	scanned := r.pos
	for {
		final := r.rerr != nil
		n, done, err := r.sc.Feed(r.buf[scanned:r.end], final)
		scanned += n
		if err != nil {
			return 0, r.wrap(err)
		}
		if done {
			return scanned, nil
		}
		if final {
			return 0, r.wrap(unstable.ErrTruncated)
		}
		before := r.pos
		if err := r.fill(); err != nil && !errors.Is(err, io.EOF) {
			return 0, r.readErr(err)
		}
		scanned -= before - r.pos
	}
}

// skip streams past the value at pos without buffering it whole: the bytes it
// has consumed are dropped at each refill, so a member the path does not want
// costs the reads and nothing else.
func (r *Reader) skip() error {
	r.sc.Reset()
	for {
		final := r.rerr != nil
		n, done, err := r.sc.Feed(r.buf[r.pos:r.end], final)
		r.pos += n
		r.hold = r.pos
		if err != nil {
			return r.wrap(err)
		}
		if done {
			return nil
		}
		if final {
			return r.wrap(unstable.ErrTruncated)
		}
		if err := r.fill(); err != nil && !errors.Is(err, io.EOF) {
			return r.readErr(err)
		}
	}
}

// fill reads more of the document, compacting the buffer when it is full and
// growing it — up to the limit — when compacting would not help because the
// value being assembled starts at the front. It reports io.EOF when the
// document's source is exhausted, which is not by itself an error: the caller
// decides whether it needed those bytes.
func (r *Reader) fill() error {
	if r.rerr != nil {
		return r.rerr
	}
	if r.end == len(r.buf) {
		if r.hold > 0 {
			copy(r.buf, r.buf[r.hold:r.end])
			shift := r.hold
			r.end -= shift
			r.pos -= shift
			r.hold = 0
		} else {
			if len(r.buf) >= r.maxBuf {
				return ErrElementTooLarge
			}
			n := len(r.buf) * 2
			if n > r.maxBuf {
				n = r.maxBuf
			}
			grown := make([]byte, n)
			copy(grown, r.buf[:r.end])
			r.buf = grown
		}
	}
	// An io.Reader may return (0, nil); bufio calls that no progress after a
	// hundred of them and so does this.
	for i := 0; i < 100; i++ {
		n, err := r.r.Read(r.buf[r.end:])
		r.end += n
		r.consumed += int64(n)
		if err != nil {
			r.rerr = err
			if n > 0 {
				return nil
			}
			return err
		}
		if n > 0 {
			return nil
		}
	}
	r.rerr = io.ErrNoProgress
	return r.rerr
}

// wrap turns the reader's own failure into the error the walk reports: a
// document that ends early is ErrTruncated only when the stream ended
// cleanly, and any other failure of the underlying reader is returned as
// itself so errors.Is reaches it.
func (r *Reader) wrap(err error) error {
	if errors.Is(err, unstable.ErrTruncated) && r.rerr != nil && !errors.Is(r.rerr, io.EOF) {
		return fmt.Errorf("json: read: %w", r.rerr)
	}
	return err
}

// startsNumber reports whether c opens the number token SkipValue's dispatch
// leaves to its default arm — every byte but the five a value can otherwise
// begin with.
func startsNumber(c byte) bool {
	return c != '"' && c != '{' && c != '[' && c != 't' && c != 'f' && c != 'n'
}
