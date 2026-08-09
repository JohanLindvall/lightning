package json

import "github.com/JohanLindvall/lightning/pkg/unstable"

// GetMany looks up several top-level members of the JSON object in data at once,
// in a single pass over the object, and returns their raw value bytes in the
// same order as keys. Where Get walks one key path and N separate Get calls
// would rescan the object N times, GetMany scans it once and distributes each
// member to the matching requested key, so extracting a handful of fields from a
// wide record costs one walk rather than one per field.
//
// The results are written into out[:0] — pass a slice to reuse (a nil out
// allocates) — and out is returned with length len(keys): out[n] is the value
// for keys[n], or nil if that key is absent from the object. Each value aliases
// data and follows Get's conventions (quotes kept for strings, the full span for
// objects and arrays, the literal token for scalars); a present member whose
// value is JSON null yields the four bytes "null", distinct from a missing key's
// nil. A duplicate key in the document takes its first occurrence; a key in keys
// not present in the object stays nil without being an error.
//
// A missing key is not an error — that is what the nil slot reports. Any other
// failure (a non-object root, malformed or truncated JSON) returns a non-nil
// error together with out as filled so far. Like Get, GetMany returns as soon as
// every requested key has been found and does not validate the rest of the
// object.
func GetMany(data []byte, keys []string, out [][]byte) ([][]byte, error) {
	return getMany(data, keys, out, false)
}

// GetManyCompact is GetMany for compact JSON — input with no whitespace between
// tokens, as compact serializers emit — skipping the inter-token whitespace
// scans GetMany makes (leading whitespace at the document start is still
// tolerated). On such input it behaves identically to GetMany but faster; given
// input that does contain inter-token whitespace it may report an error.
func GetManyCompact(data []byte, keys []string, out [][]byte) ([][]byte, error) {
	return getMany(data, keys, out, true)
}

func getMany(data []byte, keys []string, out [][]byte, compact bool) ([][]byte, error) {
	// One sized allocation (or a cleared reuse) instead of growing a nil out
	// by repeated append.
	if cap(out) < len(keys) {
		out = make([][]byte, len(keys))
	} else {
		out = out[:len(keys)]
		clear(out)
	}
	i := unstable.SkipWS(data, 0)
	if i >= len(data) {
		return out, unstable.ErrTruncated
	}
	if data[i] != '{' {
		return out, unstable.ErrExpectObject
	}
	i++
	found := 0
	for {
		i = unstable.SkipWSCompact(data, i, compact)
		if i >= len(data) {
			return out, unstable.ErrTruncated
		}
		if data[i] == '}' {
			return out, nil
		}
		// The key read with its no-escape fast path inline, the trick the generated
		// decoders use (see g.readKey in main.go): unstable.ReadKey costs 212 and so
		// never inlines, but the common case — a quoted key with no backslash — is
		// just an IndexCloseOrEscape scan and an UnsafeStr alias, both of which do
		// inline here. Only an escaped key or an error calls ReadKey. A helper
		// wrapping this does not work: it exceeds the inline budget once the SIMD
		// scan inlines into it, so the block is expanded at each site.
		var key string
		var ni int
		if i >= len(data) || data[i] != '"' {
			return out, unstable.ErrInvalidJSON
		}
		ks := i + 1
		if k := unstable.IndexCloseOrEscape(data[ks:]); ks+k < len(data) && data[ks+k] == '"' {
			key, ni = unstable.UnsafeStr(data[ks:ks+k]), ks+k+1
		} else {
			var err error
			if key, ni, err = unstable.ReadKey(data, i); err != nil {
				return out, err
			}
		}
		i = unstable.SkipWSCompact(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return out, unstable.ErrExpectColon
		}
		i = unstable.SkipWSCompact(data, i+1, compact)
		start := i
		end, err := unstable.SkipValue(data, i)
		if err != nil {
			return out, err
		}
		// Distribute this member to every requested key it matches that is still
		// unset, so the first occurrence in the document wins and duplicate
		// entries in keys are all filled. Key sets are small, so the linear scan
		// is cheaper than building a map. The string compare is tested before the
		// out[n] == nil guard: k == key fails on a length mismatch for almost
		// every member (the document keys rarely match a requested length), so
		// leading with it skips the out[n] slice-header load and its bounds check
		// on the overwhelmingly common non-matching member.
		for n, k := range keys {
			if k == key && out[n] == nil {
				out[n] = data[start:end]
				found++
			}
		}
		if found == len(keys) {
			return out, nil // every requested key found; skip the rest
		}
		i = unstable.SkipWSCompact(data, end, compact)
		if i >= len(data) {
			return out, unstable.ErrTruncated
		}
		switch data[i] {
		case '}':
			return out, nil
		case ',':
			i++
		default:
			return out, unstable.ErrInvalidJSON
		}
	}
}

// GetPaths looks up several object-key PATHS in data at once, in a single pass —
// the multi-path form of Get (as GetMany is the multi-key form). Each paths[n] is a
// key path like Get's variadic keys; out[n] receives the raw value bytes at that
// path (aliasing data, following Get's conventions) or nil if the path is absent. A
// nil/empty path selects the document root.
//
// The document is walked once: a member is descended into only when it lies on at
// least one requested path, and paths that share a prefix share that descent — so N
// nested lookups, especially under a common parent, cost one traversal instead of
// N. Results follow GetMany's rules (first occurrence wins; duplicate paths are all
// filled), and first-occurrence-wins governs the descent as well as the capture: a
// duplicate key in the document is not descended a second time, so GetPaths reports
// what Get reports for the same path. out is filled from out[:0] and returned.
//
// What each path resolves to does not depend on what else was requested, malformed
// input included. Descending for a deeper path is stricter than skipping a value —
// Get(doc, "a") on {"a":{"b" 1}} returns the value while Get(doc, "a", "x") reports
// ErrExpectColon — so a subtree only a co-requested path enters may fail to be
// walked; that failure is reported as the call's error but never retracts another
// path's value, which is the one Get would have given it alone. The error is the
// first such failure in document order, and out is returned filled as far as the
// walk got.
func GetPaths(data []byte, paths [][]string, out [][]byte) ([][]byte, error) {
	return getPaths(data, paths, out, false)
}

// GetPathsCompact is GetPaths for compact JSON — input with no whitespace between
// tokens — skipping the inter-token whitespace scans (leading whitespace at the
// document start is still tolerated). Faster on compact input; may error on input
// that does contain inter-token whitespace.
func GetPathsCompact(data []byte, paths [][]string, out [][]byte) ([][]byte, error) {
	return getPaths(data, paths, out, true)
}

func getPaths(data []byte, paths [][]string, out [][]byte, compact bool) ([][]byte, error) {
	// One sized allocation (or a cleared reuse) instead of growing a nil out
	// by repeated append.
	if cap(out) < len(paths) {
		out = make([][]byte, len(paths))
	} else {
		out = out[:len(paths)]
		clear(out)
	}
	i := unstable.SkipWS(data, 0)
	if i >= len(data) {
		return out, unstable.ErrTruncated
	}
	// One shared scratch holds the active-index set for every recursion level. The
	// walk is depth-first, so a child's set is appended right after its parent's and
	// the deepest root-to-leaf chain holds at most one set (<= len(paths)) per level.
	// Pre-sizing to len(paths)*(maxDepth+1) means append never reallocates, so the
	// parent sub-slices that point into the backing array stay valid — one allocation
	// for the whole traversal instead of a slice per descended object.
	maxDepth := 0
	for _, p := range paths {
		if len(p) > maxDepth {
			maxDepth = len(p)
		}
	}
	// The shared active-index scratch is len(paths)*(maxDepth+1) ints. For the common
	// small lookup that fits in a stack array, avoiding the heap allocation entirely;
	// larger sets fall back to a single make. The backing does not escape — out only
	// ever aliases data, never scratch.
	var stackbuf [32]int
	need := len(paths) * (maxDepth + 1)
	var scratch []int
	if need <= len(stackbuf) {
		scratch = stackbuf[:0]
	} else {
		scratch = make([]int, 0, need)
	}
	// An empty path selects the whole root value; everything else needs an object to
	// descend and joins the depth-0 active set. The root span is scanned once and
	// handed to *every* empty path — duplicate paths are all filled, as the doc
	// says and as the object walk below does; the earlier flag guarded the
	// assignment as well as the scan, so a second empty path came back nil.
	var root []byte
	rootCaptured := false
	for n, p := range paths {
		if len(p) == 0 {
			if !rootCaptured {
				end, err := unstable.SkipValue(data, i)
				if err != nil {
					return out, err
				}
				rootCaptured = true
				root = data[i:end]
			}
			out[n] = root
		} else {
			scratch = append(scratch, n)
		}
	}
	active := scratch
	if len(active) == 0 {
		return out, nil
	}
	if data[i] != '{' {
		return out, unstable.ErrExpectObject
	}
	free := scratch[len(active):cap(scratch)]
	_, err, _ := walkPaths(data, i, 0, active, free, paths, out, compact)
	return out, err
}

// walkPaths walks the object at data[i] (data[i] == '{'), matching active (indices
// into paths) at the given depth, and returns the index just past the object's '}'.
// A member whose key extends an active path either captures the value (path ends
// there) or is recursed into (path continues and the value is an object); members
// off every active path are skipped. free is the unused tail of the shared scratch,
// where a child's active set is built (never reallocated — see getPaths).
//
// The three results are (end, err, fatal). fatal distinguishes the two ways a walk
// can report an error, which the caller must treat differently: fatal means this
// frame could not be walked and end is only the offset it gave up at, while a
// non-fatal error means the frame completed — end is past its '}' — and err is
// merely the first failure met *inside* it, recovered from so the rest of the
// document could still be served (see the recovery at the descent below).
func walkPaths(data []byte, i, depth int, active, free []int, paths [][]string, out [][]byte, compact bool) (int, error, bool) {
	var firstErr error // first recovered failure; reported once the frame completes
	i++                // step over '{'
	for {
		i = unstable.SkipWSCompact(data, i, compact)
		if i >= len(data) {
			return i, unstable.ErrTruncated, true
		}
		if data[i] == '}' {
			return i + 1, firstErr, false
		}
		// Key read with the no-escape fast path inline; see getMany.
		var key string
		var ni int
		if i >= len(data) || data[i] != '"' {
			return i, unstable.ErrInvalidJSON, true
		}
		ks := i + 1
		if k := unstable.IndexCloseOrEscape(data[ks:]); ks+k < len(data) && data[ks+k] == '"' {
			key, ni = unstable.UnsafeStr(data[ks:ks+k]), ks+k+1
		} else {
			var err error
			if key, ni, err = unstable.ReadKey(data, i); err != nil {
				return ni, err, true
			}
		}
		i = unstable.SkipWSCompact(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return i, unstable.ErrExpectColon, true
		}
		i = unstable.SkipWSCompact(data, i+1, compact)
		start := i

		// Classify the active paths against this key: those that continue past it
		// (recurse into the value) and whether any end exactly here (capture). The
		// recurse set is built in free[:0] — the scratch tail — so it does not
		// allocate; siblings reuse the same region (truncated back on return).
		recurse := free[:0]
		ending := false
		for _, p := range active {
			if paths[p][depth] != key {
				continue
			}
			if depth+1 == len(paths[p]) {
				ending = true
			} else {
				recurse = append(recurse, p)
			}
		}

		var end int
		if len(recurse) > 0 && start < len(data) && data[start] == '{' {
			werr, fatal := error(nil), false
			end, werr, fatal = walkPaths(data, start, depth+1, recurse, free[len(recurse):], paths, out, compact)
			if werr != nil {
				// A descent is stricter than a skip — it reads keys and colons
				// where SkipValue only balances brackets — so a member only a
				// deeper path enters can fail where a solo lookup of a shorter
				// path succeeds. recoverDescent re-skips the member with that same
				// lenient SkipValue, so what this key's other paths get, and what
				// the members after it get, is what they would have got had the
				// deeper path not been requested. The failure is not swallowed: it
				// is held in firstErr and returned once the frame completes.
				var ok bool
				if end, ok = recoverDescent(data, start, end, fatal); !ok {
					return end, werr, true // not even skippable: give up here
				}
				if firstErr == nil {
					firstErr = werr
				}
			}
		} else {
			// start may be past the end here (a truncated "key:" with no
			// value); SkipValue bounds-checks and returns ErrTruncated.
			var err error
			if end, err = unstable.SkipValue(data, start); err != nil {
				return end, err, true
			}
		}
		if ending || len(recurse) > 0 {
			// This key consumed some paths: capture those that end here, and drop
			// every path it matched from the active set so a later duplicate of the
			// key is not descended a second time. Get stops at a key's first
			// occurrence and GetPaths is documented as its multi-path form, but the
			// descent had no such bookkeeping — on {"o":{"a":1},"o":{"b":2}} Get(o,b)
			// reported ErrKeyNotFound while GetPaths(o,b) walked into the second "o"
			// and returned 2. (The capture side was already first-occurrence-wins via
			// the out[p] == nil guard, which also keeps duplicate *requested* paths
			// all filled.) One pass does both, so a matched member costs no more
			// scans of active than the capture loop alone did, and off-path members —
			// the overwhelmingly common case — reach neither.
			//
			// Compacting active in place is safe: it is the parent frame's recurse
			// slice, which the parent does not read after this call returns (it
			// rebuilds it from free[:0] for the next member), and shrinking it never
			// moves free, so the shared scratch's no-reallocation invariant holds.
			w := 0
			for _, p := range active {
				if paths[p][depth] != key {
					active[w] = p
					w++
					continue
				}
				if out[p] == nil && depth+1 == len(paths[p]) {
					out[p] = data[start:end]
				}
			}
			active = active[:w]
		}

		i = unstable.SkipWSCompact(data, end, compact)
		if i >= len(data) {
			return i, unstable.ErrTruncated, true
		}
		switch data[i] {
		case '}':
			return i + 1, firstErr, false
		case ',':
			i++
		default:
			return i, unstable.ErrInvalidJSON, true
		}
	}
}

// recoverDescent resolves a member whose descent reported an error. A non-fatal
// error leaves the recursion's end usable as it is; a fatal one means the frame
// gave up mid-object, so the member's extent is re-taken with the lenient
// SkipValue a solo lookup of a shorter path would have used. ok is false only
// when even that fails, i.e. the member cannot be stepped over at all and the
// walk must stop.
//
// It is a separate function to keep walkPaths' loop free of a branch that only
// malformed input reaches.
func recoverDescent(data []byte, start, end int, fatal bool) (int, bool) {
	if !fatal {
		return end, true
	}
	e, err := unstable.SkipValue(data, start)
	if err != nil {
		return end, false
	}
	return e, true
}

// Get walks the object-key path keys into the JSON document data and returns
// the raw bytes of the value found at that path, without reporting a value type.
// The returned slice aliases data — for a string it includes the surrounding
// quotes (escapes left intact), for an object or array it spans the matching
// '{'..'}' or '['..']', and for a scalar it is the literal token.
//
// Each key descends one level: at every step the current value must be a JSON
// object that contains the key. A missing key returns ErrKeyNotFound; attempting
// to descend through a value that is not an object returns ErrExpectObject (with
// the index left at that value).
// With no keys Get returns the whole value at the document root. The second
// return value is the offset in data at which the returned value begins, and
// leading whitespace is tolerated at every level.
func Get(data []byte, keys ...string) ([]byte, int, error) {
	return get(data, false, keys...)
}

// GetCompact is Get for compact JSON — input with no whitespace between tokens,
// as compact serializers emit — skipping the inter-token whitespace scans Get
// makes while descending the key path (leading whitespace at the document start
// is still tolerated). On such input it behaves identically to Get but faster;
// given input that does contain inter-token whitespace it may report an error.
func GetCompact(data []byte, keys ...string) ([]byte, int, error) {
	return get(data, true, keys...)
}

// Lookup is Get without the offset return: it walks the object-key path keys into
// data and returns just the raw bytes of the value found there (aliasing data,
// following Get's conventions). Use it when the value's position in data is not
// needed — the common read-only case — and Get when it is (for example to splice
// the value back into the document). ErrKeyNotFound reports a missing key.
func Lookup(data []byte, keys ...string) ([]byte, error) {
	v, _, err := get(data, false, keys...)
	return v, err
}

// LookupCompact is Lookup for compact JSON — input with no whitespace between
// tokens — skipping the inter-token whitespace scans (leading whitespace at the
// document start is still tolerated). Faster on compact input; may report an
// error on input that does contain inter-token whitespace.
func LookupCompact(data []byte, keys ...string) ([]byte, error) {
	v, _, err := get(data, true, keys...)
	return v, err
}

func get(data []byte, compact bool, keys ...string) ([]byte, int, error) {
	i := unstable.SkipWS(data, 0)
	for _, key := range keys {
		var err error
		i, err = objectField(data, i, key, compact)
		if err != nil {
			return nil, i, err
		}
	}
	end, err := unstable.SkipValue(data, i)
	if err != nil {
		return nil, i, err
	}
	return data[i:end], i, nil
}

// ObjectEach calls fn once for every member of the JSON object reached by the
// object-key path keys in data, without reporting a value type. fn receives the
// member's decoded key and the raw bytes of its value. The value always aliases
// data, and so does the key when it contains no escapes (an escaped key is
// decoded into a fresh string); the caller must keep data unchanged while they
// are in use. The value follows the same conventions as Get — quotes kept for
// strings, the full span for objects and arrays, the literal token for scalars.
//
// With no keys ObjectEach iterates the document's root object; otherwise each
// key descends one level and the value at the end of the path must itself be an
// object (ErrExpectObject if not, ErrKeyNotFound if a key is missing). If fn
// returns a non-nil error, iteration stops and that error is returned.
// Non-target members along the path are skipped without allocating.
func ObjectEach(data []byte, fn func(key string, value []byte) error, keys ...string) error {
	return objectEach(data, fn, false, keys...)
}

// ObjectEachCompact is ObjectEach for compact JSON — input with no whitespace
// between tokens, as compact serializers emit — skipping the inter-token
// whitespace scans ObjectEach makes (leading whitespace at the document start is
// still tolerated). On such input it behaves identically to ObjectEach but
// faster; given input that does contain inter-token whitespace it may report an
// error.
func ObjectEachCompact(data []byte, fn func(key string, value []byte) error, keys ...string) error {
	return objectEach(data, fn, true, keys...)
}

func objectEach(data []byte, fn func(key string, value []byte) error, compact bool, keys ...string) error {
	i := unstable.SkipWS(data, 0)
	for _, key := range keys {
		var err error
		i, err = objectField(data, i, key, compact)
		if err != nil {
			return err
		}
	}
	i = unstable.SkipWSCompact(data, i, compact)
	if i >= len(data) {
		return unstable.ErrTruncated
	}
	if data[i] != '{' {
		return unstable.ErrExpectObject
	}
	i++
	for {
		i = unstable.SkipWSCompact(data, i, compact)
		if i >= len(data) {
			return unstable.ErrTruncated
		}
		if data[i] == '}' {
			return nil
		}
		// Key read with the no-escape fast path inline; see getMany.
		var key string
		var ni int
		if i >= len(data) || data[i] != '"' {
			return unstable.ErrInvalidJSON
		}
		ks := i + 1
		if k := unstable.IndexCloseOrEscape(data[ks:]); ks+k < len(data) && data[ks+k] == '"' {
			key, ni = unstable.UnsafeStr(data[ks:ks+k]), ks+k+1
		} else {
			var err error
			if key, ni, err = unstable.ReadKey(data, i); err != nil {
				return err
			}
		}
		i = unstable.SkipWSCompact(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return unstable.ErrExpectColon
		}
		i = unstable.SkipWSCompact(data, i+1, compact)
		start := i
		end, err := unstable.SkipValue(data, i)
		if err != nil {
			return err
		}
		if err := fn(key, data[start:end]); err != nil {
			return err
		}
		i = unstable.SkipWSCompact(data, end, compact)
		if i >= len(data) {
			return unstable.ErrTruncated
		}
		switch data[i] {
		case '}':
			return nil
		case ',':
			i++
		default:
			return unstable.ErrInvalidJSON
		}
	}
}

// ArrayEach calls fn once for every element of the JSON array reached by the
// object-key path keys in data — the array counterpart of ObjectEach. fn
// receives the element's raw bytes, aliasing data (so the caller must keep
// data unchanged while they are in use), following the same conventions as
// Get — quotes kept for strings, the full span for objects and arrays, the
// literal token for scalars.
//
// With no keys ArrayEach iterates the document's root array; otherwise each
// key descends one level and the value at the end of the path must itself be
// an array (ErrExpectArray if not, ErrKeyNotFound if a key is missing). If fn
// returns a non-nil error, iteration stops and that error is returned.
// Non-target members along the path are skipped without allocating.
func ArrayEach(data []byte, fn func(value []byte) error, keys ...string) error {
	return arrayEach(data, fn, false, keys...)
}

// ArrayEachCompact is ArrayEach for compact JSON — input with no whitespace
// between tokens, as compact serializers emit — skipping the inter-token
// whitespace scans ArrayEach makes (leading whitespace at the document start
// is still tolerated). On such input it behaves identically to ArrayEach but
// faster; given input that does contain inter-token whitespace it may report
// an error.
func ArrayEachCompact(data []byte, fn func(value []byte) error, keys ...string) error {
	return arrayEach(data, fn, true, keys...)
}

func arrayEach(data []byte, fn func(value []byte) error, compact bool, keys ...string) error {
	i := unstable.SkipWS(data, 0)
	for _, key := range keys {
		var err error
		i, err = objectField(data, i, key, compact)
		if err != nil {
			return err
		}
	}
	i = unstable.SkipWSCompact(data, i, compact)
	if i >= len(data) {
		return unstable.ErrTruncated
	}
	if data[i] != '[' {
		return unstable.ErrExpectArray
	}
	i++
	i = unstable.SkipWSCompact(data, i, compact)
	if i >= len(data) {
		return unstable.ErrTruncated
	}
	if data[i] == ']' {
		return nil
	}
	for {
		start := i
		end, err := unstable.SkipValue(data, i)
		if err != nil {
			return err
		}
		if err := fn(data[start:end]); err != nil {
			return err
		}
		i = unstable.SkipWSCompact(data, end, compact)
		if i >= len(data) {
			return unstable.ErrTruncated
		}
		switch data[i] {
		case ']':
			return nil
		case ',':
			i = unstable.SkipWSCompact(data, i+1, compact)
			if i >= len(data) {
				return unstable.ErrTruncated
			}
		default:
			return unstable.ErrInvalidJSON
		}
	}
}

// objectField scans the JSON object at data[i] (after any leading whitespace)
// for the member named key and returns the index of its value, with the value's
// own leading whitespace already skipped. It returns ErrExpectObject if data[i]
// is not an object and ErrKeyNotFound if the object has no such key. It reuses
// the scanner primitives (ReadKey, SkipValue, SkipWS) so non-target members are
// skipped without allocating.
func objectField(data []byte, i int, key string, compact bool) (int, error) {
	i = unstable.SkipWSCompact(data, i, compact)
	if i >= len(data) {
		return i, unstable.ErrTruncated
	}
	if data[i] != '{' {
		return i, unstable.ErrExpectObject
	}
	i++
	for {
		i = unstable.SkipWSCompact(data, i, compact)
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == '}' {
			return i, unstable.ErrKeyNotFound
		}
		// Key read with the no-escape fast path inline; see getMany.
		var k string
		var ni int
		if i >= len(data) || data[i] != '"' {
			return i, unstable.ErrInvalidJSON
		}
		ks := i + 1
		if n := unstable.IndexCloseOrEscape(data[ks:]); ks+n < len(data) && data[ks+n] == '"' {
			k, ni = unstable.UnsafeStr(data[ks:ks+n]), ks+n+1
		} else {
			var err error
			if k, ni, err = unstable.ReadKey(data, i); err != nil {
				return ni, err
			}
		}
		i = unstable.SkipWSCompact(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return i, unstable.ErrExpectColon
		}
		i = unstable.SkipWSCompact(data, i+1, compact)
		if k == key {
			return i, nil
		}
		end, err := unstable.SkipValue(data, i)
		if err != nil {
			return end, err
		}
		i = unstable.SkipWSCompact(data, end, compact)
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		switch data[i] {
		case '}':
			return i, unstable.ErrKeyNotFound
		case ',':
			i++
		default:
			return i, unstable.ErrInvalidJSON
		}
	}
}
