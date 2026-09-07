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
	// skipProbe bounds the bytes skip offers SkipValue before handing the
	// value to the scanner. It is what keeps the probe's cost a property of
	// the VALUE rather than of the buffer: without it a reader given a 4 MB
	// buffer would scan 4 MB to find out that a 100 MB sibling does not fit.
	// Above it a value is large enough that the probe is noise either way
	// (the scan is ~0.23 instructions a byte against the scanner's ~3.5, so
	// a wasted probe is ~7% of what the scanner then spends on the same
	// bytes, and less the larger the value gets).
	skipProbe = 4 << 10
)

// errMoreInput is not an error any caller sees: it is how a walker's inline
// value scan says "the buffered bytes do not settle this", which sends the
// value to valueMore exactly as a real error does — valueMore reads more and
// reports whatever the scan then finds, and the two cases are not worth
// telling apart before it does.
var errMoreInput = errors.New("json: value continues past the buffered bytes")

// atByte is the shape every point of a walk takes on compact input: the next
// byte is buffered and is not whitespace, so neither a refill nor a whitespace
// run can be involved. It is a helper rather than two tests written out at
// each site because it has to inline, and the functions it fronts — space,
// peek, colon, afterElement — each carry a fill loop and cost two to three
// times the inliner's budget. A member of an object passes through three of
// them.
func (r *Reader) atByte() (byte, bool) {
	if uint(r.pos) < uint(r.end) {
		if c := r.buf[r.pos]; c > ' ' {
			return c, true
		}
	}
	return 0, false
}

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
		// SkipValue's number and string arms, written out, which is what the
		// in-memory walkers do (see arrayEach in get.go) and for the same
		// reason: an element that is already buffered costs a call and a frame
		// otherwise, and on an array of numbers that is a third of the walk.
		// It has to be written out rather than reached through a helper —
		// anything holding these three calls is far past the inline budget, so
		// a helper would put back the frame it exists to remove. The dispatch
		// agrees with SkipValue's byte for byte (every byte from '-' to '9'
		// reaches its number arm), and it is uint(c)-'-' rather than
		// uint(c-'-') so that no truncation is needed before the compare; see
		// get.go for both.
		//
		// Only a number can be left undecided by the bytes the scan has seen:
		// a token ending exactly where the buffered bytes do may continue in
		// the next chunk. Every other kind is settled by a byte already read —
		// a closing quote, the balanced bracket, the literal's last letter —
		// so writing the arms out takes that test off them as well.
		buf := r.buf[:r.end]
		var end int
		var err error
		switch c := buf[r.pos]; {
		case uint(c)-'-' <= 12:
			if end, err = unstable.SkipNumber(buf, r.pos); err == nil && end == r.end && r.rerr == nil {
				err = errMoreInput
			}
		case c == '"':
			end, err = unstable.SkipString(buf, r.pos)
		default:
			end, err = unstable.SkipValue(buf, r.pos)
		}
		if err != nil {
			if end, err = r.valueMore(); err != nil {
				return err
			}
		}
		if err := fn(r.buf[r.pos:end]); err != nil {
			if err == ErrStop || errors.Is(err, ErrStop) {
				return nil
			}
			return err
		}
		r.hold, r.pos = end, end
		// afterElement's fast path, written out: on compact input the
		// separator is the very next byte and the one after it opens the next
		// element, which is once per element through a frame otherwise.
		// Nothing is consumed before the fallback, so the two cannot disagree
		// — and the loop's own top sets the hold, so the comma arm need not.
		if c, ok := r.atByte(); ok {
			if c == ']' {
				r.pos++
				return nil
			}
			if c == ',' && uint(r.pos+1) < uint(r.end) && r.buf[r.pos+1] > ' ' {
				r.pos++
				continue
			}
		}
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
		// space's fast path, written out; see atByte.
		c, ok := r.atByte()
		if !ok {
			var err error
			if c, err = r.space(); err != nil {
				return err
			}
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
		// The no-escape key read, inline: get.go's readKey trick, which
		// settles the key's end AND whether it holds an escape in one
		// vectorized scan. Both of the calls it replaces are real — SkipString
		// is past the inline budget, and String would test the same bytes a
		// second time to reach the same verdict — and on a record they are two
		// of the three frames per member the in-memory walker does not have.
		// A key that straddles the buffered bytes, or holds an escape, takes
		// the ordinary path.
		kbuf := r.buf[:r.end]
		clean := false
		var kend int
		var err error
		if k := unstable.IndexCloseOrEscapeAt(kbuf, r.pos+1); uint(k) < uint(r.end) && kbuf[k] == '"' {
			kend, clean = k+1, true
		} else if kend, err = unstable.SkipString(kbuf, r.pos); err != nil {
			if kend, err = r.valueMore(); err != nil {
				return err
			}
		}
		klen := kend - r.pos
		r.pos = kend
		// colon's fast path: the colon and the value's first byte are the next
		// two bytes on compact input. Nothing is consumed before the fallback,
		// so the two cannot disagree.
		if uint(r.pos+1) < uint(r.end) && r.buf[r.pos] == ':' && r.buf[r.pos+1] > ' ' {
			r.pos++
		} else if err := r.colon(); err != nil {
			return err
		}
		// SkipValue's number and string arms, written out; see ArrayEach. The
		// buffer is taken again here rather than reused: the key's own scan
		// may have refilled, and a refill can compact.
		buf := r.buf[:r.end]
		var end int
		switch c := buf[r.pos]; {
		case uint(c)-'-' <= 12:
			if end, err = unstable.SkipNumber(buf, r.pos); err == nil && end == r.end && r.rerr == nil {
				err = errMoreInput
			}
		case c == '"':
			end, err = unstable.SkipString(buf, r.pos)
		default:
			end, err = unstable.SkipValue(buf, r.pos)
		}
		if err != nil {
			if end, err = r.valueMore(); err != nil {
				return err
			}
		}
		// The key is decoded AFTER the value, for the reason above, and from
		// the hold rather than from the index the key scan returned — but its
		// LENGTH survives a compaction, because compaction shifts hold and
		// every index above it by the same amount. So hold+klen is the key's
		// end whether or not the value's refills moved the bytes, and nothing
		// has to be scanned again to find it. The scan above has already ruled
		// out an escape in the common case, which leaves the body to alias.
		var key string
		if clean {
			key = unstable.UnsafeStr(r.buf[r.hold+1 : r.hold+klen-1])
		} else if key, err = String(r.buf[r.hold : r.hold+klen]); err != nil {
			return err
		}
		if err := fn(key, r.buf[r.pos:end]); err != nil {
			if err == ErrStop || errors.Is(err, ErrStop) {
				return nil
			}
			return err
		}
		r.hold, r.pos = end, end
		// afterElement's fast path, written out; see ArrayEach.
		if c, ok := r.atByte(); ok {
			if c == '}' {
				r.pos++
				return nil
			}
			if c == ',' && uint(r.pos+1) < uint(r.end) && r.buf[r.pos+1] > ' ' {
				r.pos++
				continue
			}
		}
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
			// The fast paths of space, colon and afterElement, written out
			// exactly as in the walkers: a descent reads a key and steps over
			// a value per member, and the three frames around that are most of
			// what it costs on a record whose members are buffered.
			var ok bool
			if c, ok = r.atByte(); !ok {
				if c, err = r.space(); err != nil {
					return err
				}
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
			if uint(r.pos+1) < uint(r.end) && r.buf[r.pos] == ':' && r.buf[r.pos+1] > ' ' {
				r.pos++
			} else if err := r.colon(); err != nil {
				return err
			}
			if match {
				break
			}
			if err := r.skip(); err != nil {
				return err
			}
			r.hold = r.pos
			if c, ok := r.atByte(); ok {
				if c == '}' {
					r.pos++
					return unstable.ErrKeyNotFound
				}
				if c == ',' && uint(r.pos+1) < uint(r.end) && r.buf[r.pos+1] > ' ' {
					r.pos++
					r.hold = r.pos
					continue
				}
			}
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
	if uint(r.pos) >= uint(r.end) || r.buf[r.pos] != '"' {
		return "", unstable.ErrInvalidJSON
	}
	// The no-escape key read, inline; see ObjectEach. A key on the way to a
	// path is compared and dropped, so nothing here has to outlive the walk.
	buf := r.buf[:r.end]
	if k := unstable.IndexCloseOrEscapeAt(buf, r.pos+1); uint(k) < uint(r.end) && buf[k] == '"' {
		s := unstable.UnsafeStr(buf[r.pos+1 : k])
		r.pos = k + 1
		return s, nil
	}
	// A key that straddles the buffered bytes, or holds an escape: SkipString
	// settles the token without value's refill loop, and String decodes it.
	end, err := unstable.SkipString(buf, r.pos)
	if err != nil {
		if end, err = r.valueMore(); err != nil {
			return "", err
		}
	}
	k, err := String(r.buf[r.pos:end])
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
// A value that is already buffered is settled by SkipValue, which is what
// keeps the streaming walk's answer identical to the in-memory one on every
// document that fits; a value that is not goes to valueMore.
func (r *Reader) value() (int, error) {
	end, err := unstable.SkipValue(r.buf[:r.end], r.pos)
	switch {
	case err == nil && (end < r.end || r.rerr != nil || !startsNumber(r.buf[r.pos])):
		// A number token that ends exactly where the buffered bytes do may
		// still continue in the next chunk; every other kind is decided by a
		// byte SkipValue has already seen.
		return end, nil
	case err != nil && r.rerr != nil:
		// Nothing more can arrive to change the verdict, so this is the
		// in-memory walkers' error for the same bytes.
		return 0, r.wrap(err)
	}
	return r.valueMore()
}

// valueMore assembles a value whose buffered bytes have already been scanned
// and did not settle it — which is what every walker's own inline scan has
// just found out, and what value falls through to. It reads more and then
// hands the value to the resumable scanner, which reaches the same end index
// SkipValue would (ResetFast puts it on the same block scan) while scanning
// each byte ONCE, however many refills the value arrives in.
//
// The refill comes first because repeating a scan that has just failed on
// exactly these bytes can only fail again: it is the caller's scan, not an
// independent one. What used to stand here instead was a bounded number of
// refill-and-retry rounds before the scanner took over, on the reasoning that
// SkipValue was several times faster per byte than the byte-state machine and
// so worth re-running over a value of a few kilobytes. That reasoning ended
// when the scanner got the block scan: a retry re-reads the whole value prefix
// to learn what the previous one already knew, and dropping the rounds took a
// document of 148 KB elements read through a 64 KiB buffer from 24.4M
// instructions to 1.5M (-94%), while a matrix of 5.6 KB elements — where a
// retry usually did succeed — is unchanged, because one scan of the value
// replaces one scan of its prefix plus one of the whole.
func (r *Reader) valueMore() (int, error) {
	if r.rerr == nil {
		if err := r.fill(); err != nil && !errors.Is(err, io.EOF) {
			return 0, r.readErr(err)
		}
	}
	r.sc.ResetFast()
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
	// The value a path steps over is usually buffered already — a record, not
	// a megabyte — and then SkipValue settles it in one SIMD pass, where the
	// resumable scanner walks it a structural byte at a time. The fast path is
	// tried ONCE, on the bytes that have arrived, and never refills to make a
	// value fit: that is what keeps a sibling larger than the buffer streaming
	// past rather than growing it to ErrElementTooLarge.
	//
	// Only SkipValue's SUCCESS is taken. A failure means either malformed
	// bytes or a value that continues in the next chunk, and the two are not
	// distinguishable here, so the scanner re-reads the value from its first
	// byte and reports what it finds — the errors this walk has always given.
	// The one behaviour that moves is a value SkipValue accepts and the
	// scanner does not (skipfast.go's divergence classes, all malformed):
	// those are now stepped over, which is what the in-memory walkers do with
	// a member they are not asked for.
	//
	// end < lim is the whole test for "settled": the scan stopped on a byte it
	// had in front of it — a closing quote, the balancing bracket, the byte
	// after a number — so nothing that arrives later can change the answer.
	// A value that ends exactly at the last byte offered is left to the
	// scanner, which is correct rather than fast and happens once in a buffer.
	lim := r.end
	if lim-r.pos > skipProbe {
		lim = r.pos + skipProbe
	}
	if end, err := unstable.SkipValue(r.buf[:lim], r.pos); err == nil && end < lim {
		r.pos, r.hold = end, end
		return nil
	}
	// A value too large for the buffer is where the streaming skip earns its
	// keep, and ResetFast is what makes it cost the read rather than several
	// times the read: the same block scan the probe above just tried, resumed
	// across refills. Its answer on a malformed value is that scan's; see
	// ResetFast.
	r.sc.ResetFast()
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
