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
	out = out[:0]
	for range keys {
		out = append(out, nil)
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
		key, ni, err := unstable.ReadKey(data, i)
		if err != nil {
			return out, err
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
// filled). out is filled from out[:0] and returned.
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
	out = out[:0]
	for range paths {
		out = append(out, nil)
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
	// descend and joins the depth-0 active set.
	rootCaptured := false
	for n, p := range paths {
		if len(p) == 0 {
			if !rootCaptured {
				end, err := unstable.SkipValue(data, i)
				if err != nil {
					return out, err
				}
				rootCaptured = true
				out[n] = data[i:end]
			}
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
	_, err := walkPaths(data, i, 0, active, free, paths, out, compact)
	return out, err
}

// walkPaths walks the object at data[i] (data[i] == '{'), matching active (indices
// into paths) at the given depth, and returns the index just past the object's '}'.
// A member whose key extends an active path either captures the value (path ends
// there) or is recursed into (path continues and the value is an object); members
// off every active path are skipped. free is the unused tail of the shared scratch,
// where a child's active set is built (never reallocated — see getPaths).
func walkPaths(data []byte, i, depth int, active, free []int, paths [][]string, out [][]byte, compact bool) (int, error) {
	i++ // step over '{'
	for {
		i = unstable.SkipWSCompact(data, i, compact)
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == '}' {
			return i + 1, nil
		}
		key, ni, err := unstable.ReadKey(data, i)
		if err != nil {
			return ni, err
		}
		i = unstable.SkipWSCompact(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return i, unstable.ErrExpectColon
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
		if len(recurse) > 0 && data[start] == '{' {
			end, err = walkPaths(data, start, depth+1, recurse, free[len(recurse):], paths, out, compact)
		} else {
			end, err = unstable.SkipValue(data, start)
		}
		if err != nil {
			return end, err
		}
		if ending {
			for _, p := range active {
				if out[p] == nil && depth+1 == len(paths[p]) && paths[p][depth] == key {
					out[p] = data[start:end]
				}
			}
		}

		i = unstable.SkipWSCompact(data, end, compact)
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		switch data[i] {
		case '}':
			return i + 1, nil
		case ',':
			i++
		default:
			return i, unstable.ErrInvalidJSON
		}
	}
}

// Get walks the object-key path keys into the JSON document data and returns
// the raw bytes of the value found at that path, without reporting a value type.
// The returned slice aliases data — for a string it includes the surrounding
// quotes (escapes left intact), for an object or array it spans the matching
// '{'..'}' or '['..']', and for a scalar it is the literal token.
//
// Each key descends one level: at every step the current value must be a JSON
// object that contains the key, or Get returns ErrKeyNotFound (and, for a
// non-object where descent was attempted, the index is left at that value).
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
// member's decoded key and the raw bytes of its value; both alias data (so the caller
// must keep data unchanged while they are in use), and the value follows the
// same conventions as Get — quotes kept for strings, the full span for objects
// and arrays, the literal token for scalars.
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
		key, ni, err := unstable.ReadKey(data, i)
		if err != nil {
			return err
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
		k, ni, err := unstable.ReadKey(data, i)
		if err != nil {
			return ni, err
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
