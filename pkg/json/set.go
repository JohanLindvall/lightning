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
// created keys are written as JSON strings without escaping — skipping that pass
// is part of what keeps Set allocation-free — so a key containing '"', '\\' or a
// control byte does not round-trip, and a key crafted to close its own string
// (x":1,"role) splices attacker-chosen members into the result, which is
// well-formed JSON and so passes any later validity check. Keys must therefore
// come from the program, not from the input: SetChecked rejects an unsafe key
// with ErrUnsafeKey. Inter-token whitespace in in is tolerated and preserved
// outside the edited span.
//
// A key that occurs more than once in the same object is edited at its first
// occurrence, as GetMany reads the first occurrence.
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
			// The key read with its no-escape fast path inline, as in getMany
			// (see get.go): unstable.ReadKey never inlines, but the common case —
			// a quoted key with no backslash — is just an IndexCloseOrEscape scan
			// and an UnsafeStr alias, both of which do inline here. Only an
			// escaped key or an error calls ReadKey; a non-quote byte breaks out
			// exactly as ReadKey's ErrInvalidJSON did.
			var k string
			var np int
			if in[p] != '"' {
				break
			}
			ks := p + 1
			if kk := unstable.IndexCloseOrEscape(in[ks:]); ks+kk < len(in) && in[ks+kk] == '"' {
				k, np = unstable.UnsafeStr(in[ks:ks+kk]), ks+kk+1
			} else {
				var err error
				if k, np, err = unstable.ReadKey(in, p); err != nil {
					break
				}
			}
			p = unstable.SkipWS(in, np)
			if p >= len(in) || in[p] != ':' {
				break
			}
			p = unstable.SkipWS(in, p+1)
			vs := p
			if k == keys[level] {
				found, valStart = true, vs
				if level == len(keys)-1 {
					valEnd = skipValueOrEnd(in, p)
				}
				// Not the leaf: descending into this value — the next level walks
				// it member by member anyway, so don't pre-skip the whole subtree
				// here (that would scan every on-path container twice).
				break
			}
			ve := skipValueOrEnd(in, p)
			lastValEnd = ve
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
// written as plain JSON strings without escaping, so they must come from the
// program rather than from untrusted input (see Set — SetManyChecked rejects an
// unsafe key with ErrUnsafeKey). A non-object root is replaced by a fresh object
// holding all the given members, in place of that root value: whitespace before it
// and any bytes after it survive, as in Set and SetPaths. Inter-token whitespace
// outside the replaced values and the insertion point is preserved. If rawVal is
// shorter than keys the extra keys are ignored.
//
// A key occurring more than once in the document is edited at its first occurrence
// only, as in Set and SetPaths. A key listed twice in keys is likewise set once,
// from its first entry — the later entries are ignored rather than appending a
// second member — so a degenerate request cannot make SetMany write a
// duplicate-key document.
func SetMany(in, out []byte, rawVal [][]byte, keys []string) []byte {
	out = out[:0]
	n := len(keys)
	if n > len(rawVal) {
		n = len(rawVal)
	}
	j := unstable.SkipWS(in, 0)
	if j >= len(in) || in[j] != '{' {
		return setManyNonObject(in, out, rawVal, keys, n, j)
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
	prev := 0   // bytes of in already copied into out
	nfound := 0 // requested keys found so far, for the all-found early exit
	// A key requested more than once is served by its first entry alone: the
	// later entries are marked found up front, so they neither edit a document
	// duplicate of that key nor append a second member with it. Without this a
	// degenerate request emitted a duplicate-key *document* — SetMany(`{"a":1}`,
	// "a"=7, "a"=8) gave {"a":7,"a":8} where SetPaths gives {"a":7} (its
	// matched[] consumes both entries at the first occurrence and appendMembers
	// dedups by key). Entry 0 is never a duplicate, so nfound < n still holds
	// for n > 0 and the all-found exit below keeps its meaning.
	for m := 1; m < n; m++ {
		if dupKey(keys, m) {
			found[m] = true
			nfound++
		}
	}
	afterBrace := j + 1
	p := unstable.SkipWS(in, afterBrace)
	empty := p >= len(in) || in[p] == '}'
	lastValEnd := afterBrace // insertion point for new members: after the last value
	for p < len(in) && in[p] != '}' {
		// Key read with the no-escape fast path inline; see setSpan.
		var k string
		var np int
		if in[p] != '"' {
			break
		}
		ks := p + 1
		if kk := unstable.IndexCloseOrEscape(in[ks:]); ks+kk < len(in) && in[ks+kk] == '"' {
			k, np = unstable.UnsafeStr(in[ks:ks+kk]), ks+kk+1
		} else {
			var err error
			if k, np, err = unstable.ReadKey(in, p); err != nil {
				break
			}
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
				nfound++
				break
			}
		}
		if nfound == n {
			// Every requested key has been found and replaced, so no member
			// remains to append at the close, and no later member can be edited
			// (a document duplicate of a found key is left as-is either way —
			// first occurrence wins). The rest of the input passes through
			// verbatim; copy it and skip the scan, as in getMany's all-found
			// early exit.
			return append(out, in[prev:]...)
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

// setManyNonObject is SetMany's non-object-root case: the root VALUE at in[j] is
// replaced by a flat object of the requested members, keeping the whitespace
// before it and whatever follows it — Set and SetPaths both copy in[:i] and
// in[end:] around their replacement, and SetMany dropped both, so Set(" 5") gave
// " {\"a\":7}" and SetMany gave "{\"a\":7}". It is a separate function because it
// is the rare branch: written inline it grew SetMany's body (and the walk's
// i-cache footprint) by ~100 instructions for a case that returns immediately.
func setManyNonObject(in, out []byte, rawVal [][]byte, keys []string, n, j int) []byte {
	out = append(out, in[:j]...)
	out = append(out, '{')
	lead := false
	for m := 0; m < n; m++ {
		if dupKey(keys, m) {
			continue // a repeat of an earlier key: the first entry wins
		}
		if lead {
			out = append(out, ',')
		}
		lead = true
		out = appendMember(out, keys[m:m+1], rawVal[m])
	}
	out = append(out, '}')
	return append(out, in[skipValueOrEnd(in, j):]...)
}

// dupKey reports whether keys[m] repeats an earlier entry — a second request for
// a key already served, which SetMany ignores rather than acting on twice. Key
// sets are small, so the pairwise scan beats building a set (appendMembers dedups
// the same way, for the same reason).
func dupKey(keys []string, m int) bool {
	for _, k := range keys[:m] {
		if k == keys[m] {
			return true
		}
	}
	return false
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
// as plain JSON strings without escaping, at every level of a path, so they must come
// from the program rather than from untrusted input (see Set — SetPathsChecked
// rejects an unsafe path element with ErrUnsafeKey). A nil/empty path replaces the
// whole document. When one requested path is a prefix of another, the shorter (which
// replaces the whole subtree) wins and the longer is ignored. A key occurring more
// than once in the same object is edited at its first occurrence only, as in Set and
// SetMany. Inter-token whitespace outside the edited spans is preserved. If rawVal is
// shorter than paths the surplus paths are ignored.
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
	nmatched := 0 // active paths matched so far, for the root all-matched early exit
	for p < len(in) && in[p] != '}' {
		// Key read with the no-escape fast path inline; see setSpan.
		var k string
		var np int
		if in[p] != '"' {
			break
		}
		ks := p + 1
		if kk := unstable.IndexCloseOrEscape(in[ks:]); ks+kk < len(in) && in[ks+kk] == '"' {
			k, np = unstable.UnsafeStr(in[ks:ks+kk]), ks+kk+1
		} else {
			var err error
			if k, np, err = unstable.ReadKey(in, p); err != nil {
				break
			}
		}
		q := unstable.SkipWS(in, np)
		if q >= len(in) || in[q] != ':' {
			break
		}
		q = unstable.SkipWS(in, q+1)
		vs := q

		ending := -1 // an active path ending at this key (replace its value)
		var recurse []int
		// An already-matched path is skipped rather than re-applied, so a key that
		// occurs twice in one object is edited at its first occurrence only — the
		// rule Set and SetMany already follow. Without the matched[m] test every
		// duplicate was rewritten, and since the root frame's all-matched early exit
		// below returns before ever reaching a duplicate, whether a path's duplicates
		// got edited depended on whether some *unrelated* path had also been found.
		// The key compare still leads: it fails on almost every member, so the
		// matched[m] load is skipped on the overwhelmingly common non-matching key.
		for m, a := range active {
			if paths[a][depth] != k || matched[m] {
				continue
			}
			matched[m] = true
			nmatched++
			if depth+1 == len(paths[a]) {
				if ending < 0 {
					ending = a
				}
			} else {
				recurse = append(recurse, a)
			}
		}
		var ve int
		if ending >= 0 {
			ve = skipValueOrEnd(in, q)
			out = append(out, in[prev:vs]...) // copy through up to the old value...
			out = append(out, rawVal[ending]...)
			prev = ve
		} else if len(recurse) > 0 && vs < len(in) && in[vs] == '{' {
			// Recursing into this object: the recursion walks it member by member
			// and reports where it ends, so don't pre-skip the whole subtree here
			// (that would scan every on-path container twice).
			out = append(out, in[prev:vs]...)
			out, ve = setObject(in, out, vs, depth+1, recurse, paths, rawVal)
			prev = ve
		} else if len(recurse) > 0 {
			// Path continues but this value is not an object: replace it.
			ve = skipValueOrEnd(in, q)
			out = append(out, in[prev:vs]...)
			out = appendMergedObject(out, paths, rawVal, recurse, depth+1)
			prev = ve
		} else {
			ve = skipValueOrEnd(in, q)
		}
		lastValEnd = ve

		if depth == 0 && nmatched == len(active) {
			// Root-frame early exit, the all-found exit of getMany/SetMany:
			// every path routed through this frame has been consumed — a path
			// ending at a matched key had its value replaced (or was ignored
			// under the shorter-prefix-wins rule), and a path continuing past
			// it was finished by the recursion or merged replacement, which
			// create their own missing members at the deeper close. create
			// below only ever holds unmatched paths, so nothing remains to
			// append at this frame's close, and the rest of the input passes
			// through verbatim (a document duplicate of a matched key is left
			// as-is — first occurrence wins, as in Set/SetMany). Only the root
			// frame can splice-and-return: it needs no end offset beyond "copy
			// the tail" (end = len(in) makes SetPaths' in[end:] append a
			// no-op), while a nested frame would still have to locate its '}'
			// for the parent — a scan as costly as the walk this exit skips.
			return append(out, in[prev:]...), len(in)
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
