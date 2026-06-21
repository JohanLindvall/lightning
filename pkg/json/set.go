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
	start, end, insert := setSpan(in, rawVal, keys)
	out = append(out[:0], in[:start]...)
	out = append(out, insert...)
	return append(out, in[end:]...)
}

// setSpan locates the edit Set must make: the half-open byte span [start, end)
// of in to replace and the bytes to put there. It walks keys from the root,
// descending through existing objects; on the leaf it targets the existing value
// (replace) or the insertion point in its parent object (create), and on a
// missing or non-object intermediate it builds the remaining path as nested
// objects.
func setSpan(in, rawVal []byte, keys []string) (start, end int, insert []byte) {
	i := unstable.SkipWS(in, 0)
	for level := 0; level < len(keys); level++ {
		j := unstable.SkipWS(in, i)
		if j >= len(in) || in[j] != '{' {
			// The path still needs to descend but there is no object here: replace
			// this value with the remaining keys built as nested objects.
			return i, skipValueOrEnd(in, i), nestValue(keys[level:], rawVal)
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
				return valStart, valEnd, rawVal
			}
			i = valStart // descend into the existing value
			continue
		}
		// Key absent at this level: create the member. Into an empty object it
		// goes right after '{'; otherwise it is appended after the last member.
		if empty {
			return afterBrace, afterBrace, appendMember(nil, keys[level:], rawVal)
		}
		return lastValEnd, lastValEnd, appendMember([]byte{','}, keys[level:], rawVal)
	}
	// No keys: replace the whole document value.
	return i, skipValueOrEnd(in, i), rawVal
}

// nestValue builds the JSON value for a key path: rawVal itself when keys is
// empty, otherwise rawVal wrapped in one object per key, outermost first
// (["a","b"] -> {"a":{"b":rawVal}}).
func nestValue(keys []string, rawVal []byte) []byte {
	if len(keys) == 0 {
		return rawVal
	}
	out := make([]byte, 0, len(rawVal)+len(keys)*8)
	for _, k := range keys {
		out = append(out, '{', '"')
		out = append(out, k...)
		out = append(out, '"', ':')
	}
	out = append(out, rawVal...)
	for range keys {
		out = append(out, '}')
	}
	return out
}

// appendMember appends the object member `"keys[0]":<nestValue(keys[1:])>` used
// when inserting a fresh key into an existing object.
func appendMember(dst []byte, keys []string, rawVal []byte) []byte {
	dst = append(dst, '"')
	dst = append(dst, keys[0]...)
	dst = append(dst, '"', ':')
	return append(dst, nestValue(keys[1:], rawVal)...)
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
	found := make([]bool, n)
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
		// linear scan is cheaper than a map.
		for m := 0; m < n; m++ {
			if !found[m] && keys[m] == k {
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
func appendMembers(out []byte, paths [][]string, rawVal [][]byte, idx []int, depth int, lead bool) []byte {
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
		var sub []int
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
