package json

import "github.com/JohanLindvall/lightning/pkg/unstable"

// Set returns the JSON document in with the value at the object-key path keys
// replaced by the raw JSON value rawVal, written into out. When the path does
// not exist it is created: a missing member is inserted into its parent object,
// and any missing intermediate object — or a non-object value found where the
// path still needs to descend — is created as nested objects. With no keys the
// whole document is replaced by rawVal.
//
// rawVal is inserted verbatim and must be a single well-formed JSON value. Newly
// created keys are written as JSON strings without escaping, so they should not
// contain characters that need escaping. Inter-token whitespace in in is
// tolerated and preserved outside the edited span.
//
// out is filled from out[:0] and returned; pass a reusable buffer to avoid
// allocation (a nil out allocates). out must not alias in.
func Set(in, out, rawVal []byte, keys []string) []byte {
	start, end, insert, member, comma, nested := setSpan(in, rawVal, keys)
	out = append(out[:0], in[:start]...)
	switch {
	case member == nil:
		out = append(out, insert...)
	case nested:
		// A non-object intermediate is replaced by the remaining path built as
		// nested objects — appendMember's member form wrapped in braces —
		// streamed straight into out instead of materializing a temporary.
		out = append(out, '{')
		out = appendMember(out, member, rawVal)
		out = append(out, '}')
	default:
		// Append the new member straight into out (which keeps its capacity across
		// calls) instead of materializing a throwaway insert buffer per call.
		if comma {
			out = append(out, ',')
		}
		out = appendMember(out, member, rawVal)
	}
	return append(out, in[end:]...)
}

// setSpan locates the edit Set must make: the half-open byte span [start, end)
// of in to replace and the bytes to put there. It walks keys from the root,
// descending through existing objects; on the leaf it targets the existing value
// (replace) or the insertion point in its parent object (create), and on a
// missing or non-object intermediate it builds the remaining path as nested
// objects.
// The insert to make is either a byte span (insert, with member nil) or built
// from the remaining key path member: an object member to append (comma =
// whether a leading separator is needed) or, with nested set, a nested-object
// value replacing a non-object intermediate. Set streams both member forms
// directly into its output buffer rather than setSpan allocating a throwaway
// slice.
func setSpan(in, rawVal []byte, keys []string) (start, end int, insert []byte, member []string, comma, nested bool) {
	i := unstable.SkipWS(in, 0)
	for level := 0; level < len(keys); level++ {
		j := unstable.SkipWS(in, i)
		if j >= len(in) || in[j] != '{' {
			// The path still needs to descend but there is no object here: replace
			// this value with the remaining keys built as nested objects.
			return i, skipValueOrEnd(in, i), nil, keys[level:], false, true
		}
		afterBrace := j + 1
		p := unstable.SkipWS(in, afterBrace)
		empty := p >= len(in) || in[p] == '}'
		found, valStart, valEnd := false, 0, 0
		lastValEnd := afterBrace // end of the last member's value, for appending
		for p < len(in) && in[p] != '}' {
			k, np, err := unstable.ReadKey(in, p)
			if err != nil {
				break
			}
			p = unstable.SkipWS(in, np)
			if p >= len(in) || in[p] != ':' {
				break
			}
			p = unstable.SkipWS(in, p+1)
			vs := p
			ve := skipValueOrEnd(in, p)
			lastValEnd = ve
			if k == keys[level] {
				found, valStart, valEnd = true, vs, ve
				break
			}
			p = unstable.SkipWS(in, ve)
			if p < len(in) && in[p] == ',' {
				p = unstable.SkipWS(in, p+1)
				continue
			}
			break
		}
		if found {
			if level == len(keys)-1 {
				return valStart, valEnd, rawVal, nil, false, false
			}
			i = valStart // descend into the existing value
			continue
		}
		// Key absent at this level: create the member. Into an empty object it
		// goes right after '{' (no separator); otherwise it is appended after the
		// last member (with a leading comma). Set writes the member directly.
		if empty {
			return afterBrace, afterBrace, nil, keys[level:], false, false
		}
		return lastValEnd, lastValEnd, nil, keys[level:], true, false
	}
	// No keys: replace the whole document value.
	return i, skipValueOrEnd(in, i), rawVal, nil, false, false
}

// appendMember appends the object member `"keys[0]":{"keys[1]":...rawVal...}`
// used when inserting a fresh key into an existing object: the deeper keys are
// written as nested objects around rawVal, streamed directly into dst — no
// intermediate buffer, so a multi-key create allocates nothing beyond dst's own
// growth.
func appendMember(dst []byte, keys []string, rawVal []byte) []byte {
	dst = append(dst, '"')
	dst = append(dst, keys[0]...)
	dst = append(dst, '"', ':')
	for _, k := range keys[1:] {
		dst = append(dst, '{', '"')
		dst = append(dst, k...)
		dst = append(dst, '"', ':')
	}
	dst = append(dst, rawVal...)
	for range keys[1:] {
		dst = append(dst, '}')
	}
	return dst
}

// skipValueOrEnd is SkipValue but tolerant: on a malformed value it returns the
// end of input rather than an error, keeping Set best-effort like the rest of
// the scanning helpers.
func skipValueOrEnd(data []byte, i int) int {
	if e, err := unstable.SkipValue(data, i); err == nil {
		return e
	}
	return len(data)
}

// SetMany sets several of the root object's own keys in a single pass: keys[i]'s
// value becomes the raw JSON rawVal[i], whether the key already exists (its value
// is replaced in place) or not (the member is appended). It is the multi-key
// counterpart of Set restricted to the top level — SetMany is to Set what GetMany
// is to Get — editing the document in one walk where N separate Set calls would
// rescan and rewrite it N times.
//
// Like Set it is best-effort and writes into out (filled from out[:0] and
// returned); out must not alias in, which is never modified. rawVal[i] is
// inserted verbatim and must be a single well-formed JSON value; created keys are
// written as plain JSON strings without escaping. A non-object root is replaced
// by a fresh object holding all the given members. Inter-token whitespace outside
// the replaced values and the insertion point is preserved. If rawVal is shorter
// than keys the extra keys are ignored.
func SetMany(in, out []byte, rawVal [][]byte, keys []string) []byte {
	out = out[:0]
	n := len(keys)
	if n > len(rawVal) {
		n = len(rawVal)
	}
	j := unstable.SkipWS(in, 0)
	if j >= len(in) || in[j] != '{' {
		// Non-object root: replace the whole document with a flat object.
		out = append(out, '{')
		for m := 0; m < n; m++ {
			if m > 0 {
				out = append(out, ',')
			}
			out = appendMember(out, keys[m:m+1], rawVal[m])
		}
		return append(out, '}')
	}
	// The found flags live on the stack for the common small key set; only an
	// oversized set falls back to a heap slice. The backing never escapes.
	var foundBuf [64]bool
	var found []bool
	if n <= len(foundBuf) {
		found = foundBuf[:n]
	} else {
		found = make([]bool, n)
	}
	prev := 0 // bytes of in already copied into out
	afterBrace := j + 1
	p := unstable.SkipWS(in, afterBrace)
	empty := p >= len(in) || in[p] == '}'
	lastValEnd := afterBrace // insertion point for new members: after the last value
	for p < len(in) && in[p] != '}' {
		k, np, err := unstable.ReadKey(in, p)
		if err != nil {
			break
		}
		q := unstable.SkipWS(in, np)
		if q >= len(in) || in[q] != ':' {
			break
		}
		q = unstable.SkipWS(in, q+1)
		vs := q
		ve := skipValueOrEnd(in, q)
		lastValEnd = ve
		// Replace this member's value if its key was requested and not yet set
		// (the first occurrence wins, as in GetMany). Key sets are small, so the
		// linear scan is cheaper than a map. As in getMany, the string compare
		// leads: it fails on a length mismatch for almost every member, skipping
		// the found[m] load on the overwhelmingly common non-matching key.
		for m := 0; m < n; m++ {
			if keys[m] == k && !found[m] {
				out = append(out, in[prev:vs]...) // copy through up to the old value
				out = append(out, rawVal[m]...)   // ... and substitute it
				prev = ve
				found[m] = true
				break
			}
		}
		p = unstable.SkipWS(in, ve)
		if p < len(in) && in[p] == ',' {
			p = unstable.SkipWS(in, p+1)
			continue
		}
		break
	}
	// Copy the untouched remainder up to the insertion point, append any keys
	// that were not found, then copy the tail (the closing brace onward).
	out = append(out, in[prev:lastValEnd]...)
	needComma := !empty
	for m := 0; m < n; m++ {
		if found[m] {
			continue
		}
		if needComma {
			out = append(out, ',')
		}
		out = appendMember(out, keys[m:m+1], rawVal[m])
		needComma = true
	}
	return append(out, in[lastValEnd:]...)
}

// SetPaths sets several object-key PATHS in a single pass — the multi-path form of
// Set, as SetMany is its multi-key form. paths[i]'s value becomes the raw JSON
// rawVal[i], replacing the value already at that path or creating it (building any
// missing intermediate objects); paths that share a prefix are edited and created
// together, so the document is walked and rewritten once where N separate Set calls
// would rescan it N times.
//
// Like Set/SetMany it is best-effort and writes into out (filled from out[:0] and
// returned); out must not alias in, which is never modified. Each rawVal[i] is
// inserted verbatim and must be one well-formed JSON value; created keys are written
// as plain JSON strings without escaping. A nil/empty path replaces the whole
// document. When one requested path is a prefix of another, the shorter (which
// replaces the whole subtree) wins and the longer is ignored. Inter-token whitespace
// outside the edited spans is preserved. If rawVal is shorter than paths the surplus
// paths are ignored.
func SetPaths(in, out []byte, rawVal [][]byte, paths [][]string) []byte {
	out = out[:0]
	n := len(paths)
	if n > len(rawVal) {
		n = len(rawVal)
	}
	i := unstable.SkipWS(in, 0)
	// A nil/empty path replaces the whole document (the first such path wins).
	for m := 0; m < n; m++ {
		if len(paths[m]) == 0 {
			out = append(out, in[:i]...)
			out = append(out, rawVal[m]...)
			return append(out, in[skipValueOrEnd(in, i):]...)
		}
	}
	if n == 0 {
		return append(out, in...)
	}
	// The walk's transient index sets (idx here; matched/recurse/create in
	// setObject; sub in appendMembers) are all small non-escaping locals the
	// compiler keeps on the stack, so a SetPaths call allocates nothing but out.
	idx := make([]int, n)
	for m := range idx {
		idx[m] = m
	}
	if i >= len(in) || in[i] != '{' {
		// Non-object root: replace it with a fresh merged object of all paths.
		out = append(out, in[:i]...)
		out = appendMergedObject(out, paths, rawVal, idx, 0)
		return append(out, in[skipValueOrEnd(in, i):]...)
	}
	out = append(out, in[:i]...)
	out, end := setObject(in, out, i, 0, idx, paths, rawVal)
	return append(out, in[end:]...)
}

// setObject edits the JSON object at in[i] (in[i] == '{') for the active paths
// (indices into paths) at the given depth, appending the rewritten object to out
// and returning out and the index in `in` just past the object's '}'. A member on a
// path is replaced (path ends there), recursed into (path continues into an object),
// or replaced by a fresh nested object (path continues but the value is not an
// object); paths whose key is absent are created after the last member. out must
// already hold in[:i].
func setObject(in, out []byte, i, depth int, active []int, paths [][]string, rawVal [][]byte) ([]byte, int) {
	prev := i // next byte of in not yet copied into out
	p := unstable.SkipWS(in, i+1)
	empty := p >= len(in) || in[p] == '}'
	lastValEnd := i + 1 // insertion point for created members: after the last value
	matched := make([]bool, len(active))
	for p < len(in) && in[p] != '}' {
		k, np, err := unstable.ReadKey(in, p)
		if err != nil {
			break
		}
		q := unstable.SkipWS(in, np)
		if q >= len(in) || in[q] != ':' {
			break
		}
		q = unstable.SkipWS(in, q+1)
		vs := q
		ve := skipValueOrEnd(in, q)
		lastValEnd = ve

		ending := -1 // an active path ending at this key (replace its value)
		var recurse []int
		for m, a := range active {
			if paths[a][depth] != k {
				continue
			}
			matched[m] = true
			if depth+1 == len(paths[a]) {
				if ending < 0 {
					ending = a
				}
			} else {
				recurse = append(recurse, a)
			}
		}
		if ending >= 0 {
			out = append(out, in[prev:vs]...) // copy through up to the old value...
			out = append(out, rawVal[ending]...)
			prev = ve
		} else if len(recurse) > 0 {
			out = append(out, in[prev:vs]...)
			if vs < len(in) && in[vs] == '{' {
				out, _ = setObject(in, out, vs, depth+1, recurse, paths, rawVal)
			} else {
				// Path continues but this value is not an object: replace it.
				out = appendMergedObject(out, paths, rawVal, recurse, depth+1)
			}
			prev = ve
		}

		p = unstable.SkipWS(in, ve)
		if p < len(in) && in[p] == ',' {
			p = unstable.SkipWS(in, p+1)
			continue
		}
		break
	}
	objEnd := p
	if objEnd < len(in) && in[objEnd] == '}' {
		objEnd++
	} else {
		objEnd = len(in) // malformed; best-effort
	}
	// Copy through to the insertion point, create any absent paths (grouped by key
	// so prefix-sharing paths become one member), then copy the closing brace.
	out = append(out, in[prev:lastValEnd]...)
	var create []int
	for m, a := range active {
		if !matched[m] {
			create = append(create, a)
		}
	}
	out = appendMembers(out, paths, rawVal, create, depth, !empty)
	return append(out, in[lastValEnd:objEnd]...), objEnd
}

// appendMembers appends the comma-separated object members for the paths idx at the
// given depth (no surrounding braces), grouping them by their key at that level so a
// shared prefix yields one member. Per key it uses a path that ends there (its
// rawVal) or a freshly built nested object for the deeper paths. If lead is true a
// comma precedes the first member (to follow members already written).
//
// The per-key sub-set of deeper paths is built in a small per-frame stack array:
// grown with `var sub []int` appends it was the one set in the walk the compiler
// heap-allocated (append growth from nil defeats stack placement even though the
// slice never escapes), costing an allocation per created level. A key sharing
// more than eight deeper paths overflows the array and append falls back to the
// heap — rare, and only ever a spill, never a correctness issue.
func appendMembers(out []byte, paths [][]string, rawVal [][]byte, idx []int, depth int, lead bool) []byte {
	var subbuf [8]int
	for pos, a := range idx {
		key := paths[a][depth]
		dup := false
		for _, b := range idx[:pos] {
			if paths[b][depth] == key {
				dup = true
				break
			}
		}
		if dup {
			continue
		}
		if lead {
			out = append(out, ',')
		}
		lead = true
		out = append(out, '"')
		out = append(out, key...)
		out = append(out, '"', ':')
		ending := -1
		sub := subbuf[:0]
		for _, b := range idx {
			if paths[b][depth] != key {
				continue
			}
			if depth+1 == len(paths[b]) {
				if ending < 0 {
					ending = b
				}
			} else {
				sub = append(sub, b)
			}
		}
		if ending >= 0 {
			out = append(out, rawVal[ending]...)
		} else {
			out = appendMergedObject(out, paths, rawVal, sub, depth+1)
		}
	}
	return out
}

// appendMergedObject appends '{' + appendMembers + '}': the object built from the
// paths idx at depth, used to create a member value or a non-object replacement.
func appendMergedObject(out []byte, paths [][]string, rawVal [][]byte, idx []int, depth int) []byte {
	out = append(out, '{')
	out = appendMembers(out, paths, rawVal, idx, depth, false)
	return append(out, '}')
}
