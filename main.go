// Command lightning generates a fast, allocation-light json.Unmarshaler
// implementation for the struct types declared in a Go source file.
//
// Usage:
//
//	go run . bench/1/data.go
//
// For each input file FOO.go it writes FOO_unmarshal.go next to it, containing
// an UnmarshalJSON method for every top-level struct type. The generated code
// calls the shared scanner helpers in pkg/unstable.
//
// Supported field types: string, bool, all sized int/uint kinds, float32,
// float64, json.Number, time.Time, json.RawMessage, nested (named or anonymous)
// structs, slices, maps with string keys, pointers, and interface{}/any (decoded
// into the usual Go representation of an arbitrary JSON value).
package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go run . <file.go> [more.go ...]")
		os.Exit(2)
	}
	for _, in := range os.Args[1:] {
		if err := generate(in); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", in, err)
			os.Exit(1)
		}
	}
}

func generate(inPath string) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, inPath, nil, parser.SkipObjectResolution|parser.ParseComments)
	if err != nil {
		return err
	}

	g := &gen{
		fset:             fset,
		pkg:              file.Name.Name,
		structTypes:      map[string]*ast.StructType{},
		sliceTypes:       map[string]*ast.ArrayType{},
		mapTypes:         map[string]*ast.MapType{},
		order:            nil,
		used:             map[string]bool{},
		memo:             map[string]string{},
		compactTypes:     map[string]bool{},
		nocopyTypes:      map[string]bool{},
		destructiveTypes: map[string]bool{},
	}

	// Generated helper functions are named "lightning<ImportPath><Type>decode...",
	// a system-derived prefix that makes them unique across files and packages so
	// decoders for several types can share one package without their helpers
	// colliding — no manual annotation needed. The import path is resolved from
	// the enclosing go.mod; if that fails (e.g. GOPATH mode) the package name
	// stands in. The UnmarshalJSON methods keep their exact name (json.Unmarshaler
	// requires it) and never collide, being keyed by their receiver type.
	if ip, ok := importPathFor(inPath); ok {
		g.pathFrag = sanitizeIdent(ip)
	} else {
		g.pathFrag = sanitizeIdent(file.Name.Name)
	}

	// Collect every top-level struct type, in source order, recording the
	// //lightning:compact / :nocopy / :destructive directives each carries.
	for _, d := range file.Decls {
		gd, ok := d.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, s := range gd.Specs {
			ts, ok := s.(*ast.TypeSpec)
			if !ok {
				continue
			}
			switch t := ts.Type.(type) {
			case *ast.StructType:
				g.structTypes[ts.Name.Name] = t
				g.order = append(g.order, ts.Name.Name)
			case *ast.ArrayType:
				// A named slice root type (type Foo []T) for array-root documents.
				if t.Len == nil {
					g.sliceTypes[ts.Name.Name] = t
					g.order = append(g.order, ts.Name.Name)
				} else {
					continue
				}
			case *ast.MapType:
				// A named map root type (type Foo map[string]V) for object-root
				// documents that are data maps rather than fixed records.
				g.mapTypes[ts.Name.Name] = t
				g.order = append(g.order, ts.Name.Name)
			default:
				continue
			}
			if hasDirective("lightning:compact", gd.Doc, ts.Doc) {
				g.compactTypes[ts.Name.Name] = true
			}
			if hasDirective("lightning:nocopy", gd.Doc, ts.Doc) {
				g.nocopyTypes[ts.Name.Name] = true
			}
			if hasDirective("lightning:destructive", gd.Doc, ts.Doc) {
				g.destructiveTypes[ts.Name.Name] = true
			}
		}
	}
	if len(g.order) == 0 {
		return fmt.Errorf("no top-level struct or slice types found")
	}

	// A top-level type that is nested inside another (referenced by one of its
	// fields, directly or through slices/maps/pointers/anonymous structs) gets
	// only an internal decode function, not its own UnmarshalJSON method: the
	// referencing root emits the function when it reaches the field. Emitting a
	// method on such a nested type would attach json.Unmarshaler to it, so the
	// `type logStd Log` reflection baselines (stdlib, sonic, json/v2) used in the
	// benchmarks would route through the generated decoder instead. Named nested
	// types are needed for recursive or shared schemas (e.g. a tree node, a FHIR
	// code) that an anonymous struct cannot express. With one top-level type — the
	// common case — nothing is referenced, so every type still gets its method.
	referenced := map[string]bool{}
	for _, name := range g.order {
		if st, ok := g.structTypes[name]; ok {
			for _, f := range st.Fields.List {
				g.markReferenced(f.Type, name, referenced)
			}
		} else if at, ok := g.sliceTypes[name]; ok {
			g.markReferenced(at.Elt, name, referenced)
		} else if mt, ok := g.mapTypes[name]; ok {
			g.markReferenced(mt.Value, name, referenced)
		}
	}
	allReferenced := len(referenced) == len(g.order) // degenerate (fully cyclic); emit all

	var methods []string
	for _, name := range g.order {
		if referenced[name] && !allReferenced {
			continue // nested in another type; its decoder is emitted there
		}
		g.compact = g.compactTypes[name]
		g.destructive = g.destructiveTypes[name]
		// A destructive root aliases the very buffer it decodes into, so it is nocopy
		// too; a plain //lightning:nocopy root (slice/map) aliases its keys/elements.
		g.nocopy = g.destructive || g.nocopyTypes[name]
		g.prefix = "lightning" + g.pathFrag + name
		methods = append(methods, g.genUnmarshal(name))
	}
	if len(g.errs) > 0 {
		return errors.Join(g.errs...)
	}

	src := g.assemble(inPath, methods)
	out, ferr := format.Source([]byte(src))
	if ferr != nil {
		out = []byte(src) // unformatted, for inspecting the malformed output
	}

	outPath := strings.TrimSuffix(inPath, ".go") + "_unmarshal.go"
	if err := os.WriteFile(outPath, out, 0o644); err != nil {
		return err
	}
	if ferr != nil {
		// Unparseable output means a generator bug; fail (the file is written
		// above so it can be inspected).
		return fmt.Errorf("generated code did not parse, wrote unformatted %s: %w", outPath, ferr)
	}
	fmt.Fprintf(os.Stderr, "wrote %s\n", outPath)
	return nil
}

// gen holds the state accumulated while walking the type graph.
type gen struct {
	fset        *token.FileSet
	pkg         string
	structTypes map[string]*ast.StructType
	sliceTypes  map[string]*ast.ArrayType // named slice root types (type X []T)
	mapTypes    map[string]*ast.MapType   // named map root types (type X map[string]V)
	order       []string

	used map[string]bool   // reserved decoder function names
	memo map[string]string // type-key -> decoder function name (dedupe)

	// Comment-directive selectors, keyed by top-level type name:
	//   //lightning:compact      -> compactTypes      (elide inter-token SkipWS)
	//   //lightning:nocopy       -> nocopyTypes       (slice/map root aliases keys/elements)
	//   //lightning:destructive  -> destructiveTypes  (unescape strings into the input
	//                                                   buffer, destroying it; implies nocopy)
	compactTypes     map[string]bool
	nocopyTypes      map[string]bool
	destructiveTypes map[string]bool

	// Working flags for the root type currently being generated, derived from the
	// directive sets above.
	compact     bool
	destructive bool   // //lightning:destructive: unescape strings in place
	nocopy      bool   // //lightning:nocopy root, or a destructive root (which aliases what it decodes)
	pathFrag    string // import path sanitized into an identifier fragment
	prefix      string // current per-type prefix for generated function names

	decoders []string // generated decoder functions, in creation order
	errs     []error  // generation errors; reported together after the walk

	needBool, needUint, needFloat, needAny, needJSON, needNoCopyStr, needTime bool
}

func (g *gen) typeStr(e ast.Expr) string {
	var b strings.Builder
	_ = printer.Fprint(&b, g.fset, e) // writing to a strings.Builder never fails
	return b.String()
}

func (g *gen) uniq(base string) string {
	name := base
	for n := 2; g.used[name]; n++ {
		name = base + strconv.Itoa(n)
	}
	g.used[name] = true
	return name
}

// hasDirective reports whether any of the given doc-comment groups carries the
// //<name> directive on its own line — e.g. //lightning:compact. Both the type
// spec's doc and the enclosing GenDecl's doc are checked, so the directive may sit
// above the type inside a `type (...)` block or above a standalone `type X`.
//
// Whitespace anywhere in the directive is ignored: a space after the //, a
// trailing space, or spaces around the colon are all tolerated, so
// "//lightning:compact", "// lightning:compact " and "// lightning: compact" are
// equivalent. strings.Fields splits on any run of whitespace and joining the
// pieces with "" collapses it all away before the comparison.
func hasDirective(name string, groups ...*ast.CommentGroup) bool {
	for _, cg := range groups {
		if cg == nil {
			continue
		}
		for _, c := range cg.List {
			if strings.Join(strings.Fields(strings.TrimPrefix(c.Text, "//")), "") == name {
				return true
			}
		}
	}
	return false
}

// importPathFor resolves the Go import path of the package that inPath belongs
// to by finding the enclosing go.mod and joining its module path with the
// directory's path relative to the module root. It reports false when no go.mod
// is found (e.g. GOPATH mode), so the caller can fall back to the package name.
func importPathFor(inPath string) (string, bool) {
	abs, err := filepath.Abs(inPath)
	if err != nil {
		return "", false
	}
	dir := filepath.Dir(abs)
	for d := dir; ; {
		if data, err := os.ReadFile(filepath.Join(d, "go.mod")); err == nil {
			mod := modulePath(data)
			if mod == "" {
				return "", false
			}
			rel, err := filepath.Rel(d, dir)
			if err != nil {
				return "", false
			}
			if rel = filepath.ToSlash(rel); rel == "." {
				return mod, true
			}
			return mod + "/" + rel, true
		}
		parent := filepath.Dir(d)
		if parent == d { // reached the filesystem root without a go.mod
			return "", false
		}
		d = parent
	}
}

// modulePath returns the module path from the "module" line of a go.mod file's
// contents, or "" if absent.
func modulePath(goMod []byte) string {
	for _, line := range strings.Split(string(goMod), "\n") {
		if rest, ok := strings.CutPrefix(strings.TrimSpace(line), "module "); ok {
			return strings.TrimSpace(rest)
		}
	}
	return ""
}

// sanitizeIdent folds an arbitrary string (an import path or package name) into
// a camel-cased run of [A-Za-z0-9], so it can sit inside a generated identifier.
// Each maximal alphanumeric segment is capitalized: "github.com/foo/bar-baz"
// becomes "GithubComFooBarBaz".
func sanitizeIdent(s string) string {
	var b strings.Builder
	upNext := true
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9':
			if upNext {
				b.WriteRune(unicode.ToUpper(r))
				upNext = false
			} else {
				b.WriteRune(r)
			}
		default:
			upNext = true
		}
	}
	return b.String()
}

// decFn applies the current per-type prefix to a generated decoder function's
// name, yielding the system-unique "lightning<ImportPath><Type>decode..." form.
func (g *gen) decFn(name string) string { return g.prefix + name }

// skipWS returns the statement assigning dst the first non-whitespace index at
// or after src. In compact mode it becomes a plain dst = src, or nothing when
// dst and src are identical (the assignment would be a no-op).
func (g *gen) skipWS(dst, src string) string {
	if g.compact {
		if dst == src {
			return ""
		}
		return dst + " = " + src
	}
	// Skip inter-token whitespace inline. The 0- and 1-byte cases (a compact
	// stream, or the single space after a token in lightly-spaced JSON) resolve
	// with one or two compares and never call out; only a real run of two or more
	// whitespace bytes — the indentation of pretty-printed input — is handed to
	// the SWAR unstable.SkipWSRun, which eats it eight bytes at a time. Doing this
	// at the call site rather than in unstable.SkipWS sidesteps the inliner budget
	// that would otherwise force every skip through a function call.
	set := ""
	if dst != src {
		set = dst + " = " + src + "\n\t"
	}
	return fmt.Sprintf(`%[1]sif %[2]s < len(data) && data[%[2]s] <= ' ' {
		%[2]s++
		if %[2]s < len(data) && data[%[2]s] <= ' ' {
			%[2]s = unstable.SkipWSRun(data, %[2]s+1)
		}
	}`, set, dst)
}

// readKey emits the object-key read with its no-escape fast path written inline,
// the way skipWS inlines the whitespace fast path. The common case — a quoted
// key with no backslash escape — is an IndexCloseOrEscape scan plus an UnsafeStr
// alias at the call site, costing no ReadKey call; only an escaped key (or an
// error) falls back to unstable.ReadKey. It declares key and ni for the loop that
// follows. The enclosing decoder returns (int, error).
func (g *gen) readKey() string {
	return `var key string
	var ni int
	if i >= len(data) || data[i] != '"' {
		return i, unstable.ErrInvalidJSON
	}
	ks := i + 1
	if k := unstable.IndexCloseOrEscape(data[ks:]); ks+k < len(data) && data[ks+k] == '"' {
		key, ni = unstable.UnsafeStr(data[ks:ks+k]), ks+k+1
	} else {
		var err error
		if key, ni, err = unstable.ReadKey(data, i); err != nil {
			return ni, err
		}
	}`
}

// cmark and csuf distinguish a decoder's compact and destructive variants from their
// plain counterparts in, respectively, memo keys and generated function names, so
// the same type reached from roots with different directive combinations yields a
// distinct decoder per combination.
func (g *gen) cmark() string {
	var m string
	if g.compact {
		m += "compact:"
	}
	if g.destructive {
		m += "destructive:"
	}
	return m
}

func (g *gen) csuf() string {
	var s string
	if g.compact {
		s += "Compact"
	}
	if g.destructive {
		s += "Destructive"
	}
	return s
}

// genUnmarshal emits the UnmarshalJSON method for a named root type and makes sure
// its decoder (and everything it reaches under the current g.compact/g.destructive/
// g.nocopy flags) is generated. For a //lightning:destructive type the doc comment
// warns that the method mutates its input.
func (g *gen) genUnmarshal(name string) string {
	// The decode call differs for a slice or map root: its decoder takes *[]T or
	// *map[string]V, and the receiver *Name (whose underlying type matches) is
	// converted to it.
	var call string
	nocopy := g.nocopy // //lightning:nocopy or :destructive root: alias the slice/map root's keys/elements
	switch {
	case g.sliceTypes[name] != nil:
		at := g.sliceTypes[name]
		call = fmt.Sprintf("%s((*[]%s)(v), data, i)", g.sliceDecoder(at.Elt, name, nocopy, false), g.typeStr(at.Elt))
	case g.mapTypes[name] != nil:
		mt := g.mapTypes[name]
		call = fmt.Sprintf("%s((*map[string]%s)(v), data, i)", g.mapDecoder(mt.Key, mt.Value, name, nocopy, false), g.typeStr(mt.Value))
	default:
		call = g.namedStruct(name) + "(v, data, i)"
	}
	var doc string
	if g.destructive {
		doc = fmt.Sprintf(`// UnmarshalJSON parses JSON into the %[1]s, unescaping every "nocopy" string
// field directly into data instead of allocating. This MUTATES and effectively
// destroys data: the bytes of every escaped string are overwritten, so data is no
// longer valid JSON afterward and must not be read or aliased again.`, name)
	} else {
		doc = fmt.Sprintf(`// UnmarshalJSON parses JSON into the %[1]s. Fields whose json tag carries the
// "nocopy" option alias the input data instead of copying it; if any are
// present, the caller must keep data unchanged while the result is in use.`, name)
	}
	return fmt.Sprintf(`%[2]s
func (v *%[1]s) UnmarshalJSON(data []byte) error {
	i := unstable.SkipWS(data, 0)
	if i >= len(data) {
		return unstable.ErrTruncated
	}
	if data[i] == 'n' {
		end, err := unstable.ExpectNull(data, i)
		if err != nil {
			return err
		}
		if unstable.SkipWS(data, end) != len(data) {
			return unstable.ErrInvalidJSON
		}
		return nil
	}
	end, err := %[3]s
	if err != nil {
		return err
	}
	if unstable.SkipWS(data, end) != len(data) {
		return unstable.ErrInvalidJSON
	}
	return nil
}`, name, doc, call)
}

// namedStruct returns (generating on first use) the decoder for a named struct.
func (g *gen) namedStruct(name string) string {
	key := g.prefix + g.cmark() + "named:" + name
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	fn := g.decFn("decode" + name + g.csuf())
	g.used[fn] = true
	g.memo[key] = fn // set before body so recursive types terminate
	g.genStructBody(fn, "*"+name, g.structTypes[name])
	return fn
}

// anonStruct returns the decoder for an anonymous struct type, named after the
// field it was first reached through.
func (g *gen) anonStruct(t *ast.StructType, hint string) string {
	key := g.prefix + g.cmark() + "anon:" + g.typeStr(t)
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	fn := g.uniq(g.decFn("decode" + cap1(hint) + g.csuf()))
	g.memo[key] = fn
	g.genStructBody(fn, "*"+g.typeStr(t), t)
	return fn
}

// fieldInfo describes one decodable field reached while walking a struct and its
// embedded structs: the JSON name(s) it answers to, the Go lvalue to decode into
// (e.g. "v.A" or the promoted "v.Inner.B"), its type and tag options, the
// embedding depth (0 for a direct field), whether it carries an explicit JSON
// name, and the pointer-embed allocation guards that must run before the lvalue
// is usable.
type fieldInfo struct {
	keys   []string
	dest   string
	typ    ast.Expr
	nocopy bool
	lax    bool
	unwrap bool
	tagged bool
	depth  int
	allocs []string
}

// collectFields appends to out the decodable fields of st, flattening embedded
// structs per Go's field promotion: an embedded struct's exported fields are
// promoted into the parent (an embedded *struct is allocated on demand via an
// allocs guard), an embedded field with an explicit JSON tag name is a plain
// named field rather than promoted, and an embedded non-struct (or a type whose
// definition isn't visible, e.g. from another package) is keyed by its type
// name. prefix is the Go access prefix ("v.", "v.Inner."), depth the embedding
// depth, and seen guards against the (illegal-in-Go, but cheap to defend) cycle.
func (g *gen) collectFields(st *ast.StructType, prefix string, depth int, allocs []string, seen map[string]bool, out *[]fieldInfo) {
	for _, f := range st.Fields.List {
		tagNames, skip, nocopy, lax, unwrap := jsonTag(f.Tag)
		if skip {
			continue
		}
		if len(f.Names) == 0 { // embedded (anonymous) field
			name := embeddedName(f.Type)
			if name == "" {
				g.errs = append(g.errs, fmt.Errorf("unsupported embedded field %s", g.typeStr(f.Type)))
				continue
			}
			if len(tagNames) == 0 {
				if est, pointee, isPtr, ok := g.embeddedStruct(f.Type); ok {
					child := allocs
					if isPtr {
						child = append(append([]string(nil), allocs...),
							fmt.Sprintf("if %[1]s == nil {\n%[1]s = new(%[2]s)\n}", prefix+name, pointee))
					}
					if k := g.typeStr(f.Type); !seen[k] {
						seen[k] = true
						g.collectFields(est, prefix+name+".", depth+1, child, seen, out)
						delete(seen, k)
					}
					continue
				}
			}
			// Tagged embed, or a non-struct/opaque embed: a plain named field
			// keyed by the tag name(s) or, lacking those, the embedded type name.
			keys := tagNames
			if len(keys) == 0 {
				keys = []string{name}
			}
			*out = append(*out, fieldInfo{keys: keys, dest: prefix + name, typ: f.Type,
				nocopy: nocopy, lax: lax, unwrap: unwrap, tagged: len(tagNames) > 0, depth: depth, allocs: allocs})
			continue
		}
		for _, nm := range f.Names {
			if !ast.IsExported(nm.Name) {
				continue
			}
			keys := tagNames
			if len(keys) == 0 {
				keys = []string{nm.Name}
			}
			*out = append(*out, fieldInfo{keys: keys, dest: prefix + nm.Name, typ: f.Type,
				nocopy: nocopy, lax: lax, unwrap: unwrap, tagged: len(tagNames) > 0, depth: depth, allocs: allocs})
		}
	}
}

// embeddedName returns the Go selector for an embedded field — the unqualified
// name of its type (Inner, *Inner, pkg.Inner all select v.Inner) — or "" for a
// shape that cannot be embedded.
func embeddedName(expr ast.Expr) string {
	switch t := unparen(expr).(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return embeddedName(t.X)
	case *ast.SelectorExpr:
		return t.Sel.Name
	}
	return ""
}

// embeddedStruct resolves an embedded field's type to the struct whose fields
// should be promoted, reporting the pointee type name (for new()) and whether the
// embed is a pointer. It returns ok=false when the type is not a struct defined
// in this file (a scalar, or a type from another package whose fields aren't
// visible), in which case the embed is handled as a plain named field.
func (g *gen) embeddedStruct(expr ast.Expr) (st *ast.StructType, pointee string, isPtr, ok bool) {
	e := unparen(expr)
	if star, isStar := e.(*ast.StarExpr); isStar {
		isPtr = true
		e = unparen(star.X)
	}
	switch t := e.(type) {
	case *ast.Ident:
		if s, found := g.structTypes[t.Name]; found {
			return s, t.Name, isPtr, true
		}
	case *ast.StructType:
		return t, g.typeStr(t), isPtr, true
	}
	return nil, "", false, false
}

func (g *gen) genStructBody(fn, paramType string, st *ast.StructType) {
	var fields []fieldInfo
	g.collectFields(st, "v.", 0, nil, map[string]bool{}, &fields)

	// Resolve each JSON name to its dominant field by Go's promotion rules: the
	// shallowest field wins; a tie at the shallowest depth is broken only by a
	// single tagged field, otherwise the name is ambiguous. An ambiguous clash
	// among the struct's own (depth-0) fields is a user error; a deeper one is
	// dropped silently, as encoding/json does.
	type cand struct {
		idx, depth int
		tagged     bool
	}
	byKey := map[string][]cand{}
	for i, f := range fields {
		for _, k := range f.keys {
			byKey[k] = append(byKey[k], cand{i, f.depth, f.tagged})
		}
	}
	winner := map[string]int{}
	for k, cs := range byKey {
		min := cs[0].depth
		for _, c := range cs {
			if c.depth < min {
				min = c.depth
			}
		}
		var atMin, taggedAtMin []cand
		for _, c := range cs {
			if c.depth == min {
				atMin = append(atMin, c)
				if c.tagged {
					taggedAtMin = append(taggedAtMin, c)
				}
			}
		}
		switch {
		case len(atMin) == 1:
			winner[k] = atMin[0].idx
		case len(taggedAtMin) == 1:
			winner[k] = taggedAtMin[0].idx
		case min == 0:
			g.errs = append(g.errs, fmt.Errorf("%s: json name %q is mapped more than once", strings.TrimPrefix(paramType, "*"), k))
		}
	}

	// Emit one case per field, in collection order, listing the names it won.
	var cases strings.Builder
	for i, f := range fields {
		var won []string
		for _, k := range f.keys {
			if w, ok := winner[k]; ok && w == i {
				won = append(won, k)
			}
		}
		if len(won) == 0 {
			continue
		}
		quoted := make([]string, len(won))
		for j, k := range won {
			quoted[j] = strconv.Quote(k)
		}
		hint := f.dest[strings.LastIndexByte(f.dest, '.')+1:]
		var code string
		if f.lax {
			code = g.laxField(f.dest, f.typ, hint, f.nocopy)
		} else {
			code = g.field(f.dest, f.typ, hint, f.nocopy, false)
		}
		if f.unwrap {
			code = unwrapField(code)
		}
		if len(f.allocs) > 0 {
			code = strings.Join(f.allocs, "\n") + "\n" + code
		}
		fmt.Fprintf(&cases, "\tcase %s:\n%s\n", strings.Join(quoted, ", "), code)
	}

	// The loop top is reached only after '{' (first iteration) or after a
	// comma, and a '}' after a member returns from the post-value check — so a
	// '}' at the loop top on a non-first iteration is exactly a trailing comma
	// ({"a":1,}), rejected as encoding/json does. The `first` flag keeps the
	// loop shape byte-identical to the lenient form otherwise (one register
	// move per member): a *rotated* loop (closer checked before the loop and
	// after each member only) measured cloudflare +11% — the wide decoder is
	// that layout-sensitive — so don't restructure, flag.
	body := fmt.Sprintf(`func %[1]s(v %[2]s, data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, unstable.ErrTruncated
	}
	if data[i] == 'n' {
		return unstable.ExpectNull(data, i)
	}
	if data[i] != '{' {
		return i, unstable.ErrExpectObject
	}
	i++
	for first := true; ; first = false {
		%[4]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == '}' {
			if first {
				return i + 1, nil
			}
			return i, unstable.ErrInvalidJSON
		}
		%[8]s
		%[5]s
		if i >= len(data) || data[i] != ':' {
			return i, unstable.ErrExpectColon
		}
		%[6]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		switch key {
%[3]s
		default:
			end, err := unstable.SkipValue(data, i)
			if err != nil {
				return end, err
			}
			i = end
		}
		%[7]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == '}' {
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, unstable.ErrInvalidJSON
		}
		i++
	}
}`, fn, paramType, cases.String(), g.skipWS("i", "i"), g.skipWS("i", "ni"), g.skipWS("i", "i+1"), g.skipWS("i", "i"), g.readKey())
	g.decoders = append(g.decoders, body)
}

// field returns the code that decodes a JSON value at data[i] into the Go
// lvalue dest, advancing i and returning (end, err) on failure. hint suggests a
// name for any decoder this value needs. nocopy requests that string/raw leaves
// alias the input bytes; lax requests the lenient time parser for time.Time
// leaves. Both propagate through slices, maps and pointers but stop at struct
// boundaries, where each field's own tag governs.
func (g *gen) field(dest string, expr ast.Expr, hint string, nocopy, lax bool) string {
	switch t := expr.(type) {
	case *ast.ParenExpr:
		return g.field(dest, t.X, hint, nocopy, lax)

	case *ast.StarExpr:
		// new(T) spells the pointee type, so its package must be imported.
		if isTime(t.X) {
			g.needTime = true
		}
		if isRaw(t.X) {
			g.needJSON = true
		}
		inner := g.field("(*"+dest+")", t.X, hint, nocopy, lax)
		return fmt.Sprintf(`if data[i] == 'n' {
	end, err := unstable.ExpectNull(data, i)
	if err != nil {
		return end, err
	}
	%[1]s = nil
	i = end
} else {
	%[1]s = new(%[2]s)
	%[3]s
}`, dest, g.typeStr(t.X), inner)

	case *ast.Ident:
		if isScalar(t.Name) {
			return g.scalar(dest, t.Name, nocopy)
		}
		if _, ok := g.structTypes[t.Name]; ok {
			// A struct's own field tags govern its nocopy/lax behavior.
			return g.callDecoder(dest, g.namedStruct(t.Name))
		}
		if t.Name == "any" {
			return g.anyValue(dest)
		}
		return g.unsupportedf("unknown type %q for %s", t.Name, dest)

	case *ast.SelectorExpr:
		if t.Sel.Name == "RawMessage" || t.Sel.Name == "RawValue" {
			return g.rawMessage(dest, nocopy)
		}
		if isNumber(t) {
			return g.numberRead(dest, nocopy)
		}
		if isTime(t) {
			return g.timeRead(dest, lax)
		}
		return g.unsupportedf("unsupported type %s for %s", g.typeStr(t), dest)

	case *ast.StructType:
		return g.callDecoder(dest, g.anonStruct(t, hint))

	case *ast.ArrayType:
		if t.Len != nil {
			if fn := batchArrayFn(t.Elt); fn != "" {
				// Batched fixed-array read: pass the array as a slice; one call
				// per point instead of a generated decoder call plus a reader
				// call per element.
				return fmt.Sprintf(`end, err := %s((%s)[:], data, i)
	if err != nil {
		return end, err
	}
	i = end`, fn, dest)
			}
			return g.callDecoder(dest, g.arrayDecoder(t, hint, nocopy, lax))
		}
		return g.callDecoder(dest, g.sliceDecoder(t.Elt, hint, nocopy, lax))

	case *ast.MapType:
		fn := g.mapDecoder(t.Key, t.Value, hint, nocopy, lax)
		if fn == "" {
			return g.skipEmit()
		}
		return g.callDecoder(dest, fn)

	case *ast.InterfaceType:
		return g.anyValue(dest)

	default:
		return g.unsupportedf("unsupported type %s for %s", g.typeStr(expr), dest)
	}
}

// callDecoder invokes a (struct/slice/map) decoder function on &dest.
func (g *gen) callDecoder(dest, fn string) string {
	return fmt.Sprintf(`end, err := %s(&%s, data, i)
if err != nil {
	return end, err
}
i = end`, fn, dest)
}

// unsupportedf records a generation error for a type the generator cannot decode
// and returns skip code so the walk can continue and surface every such field at
// once. Because generate() returns the joined g.errs (and writes no output) when
// any are present, an unsupported field fails the run with a non-zero exit rather
// than silently producing a decoder that drops it.
func (g *gen) unsupportedf(format string, args ...any) string {
	g.errs = append(g.errs, fmt.Errorf(format, args...))
	return g.skipEmit()
}

// laxField emits the decode for a field carrying the "lax" tag option. The value
// is decoded into a scratch variable; on success it is committed to dest, and on
// any error the input value is skipped and dest left unset. Because a well-formed
// JSON value of the wrong type is skippable while genuinely malformed JSON is
// not, only type mismatches are swallowed: a syntax error still propagates.
func (g *gen) laxField(dest string, expr ast.Expr, hint string, nocopy bool) string {
	fn := g.valueDecoder(expr, hint, nocopy, true)
	if fn == "" {
		// Unsupported type: skipping already leaves the field unset.
		return g.skipEmit()
	}
	return fmt.Sprintf(`var lax %[1]s
end, err := %[2]s(&lax, data, i)
if err != nil {
	end, err = unstable.SkipValue(data, i)
	if err != nil {
		return end, err
	}
} else {
	%[3]s = lax
}
i = end`, g.typeStr(expr), fn, dest)
}

// valueDecoder returns the name of a function decoding the JSON value at data[i]
// into *T, where T is the field type. Struct, slice and map types reuse their
// existing decoders; everything else gets a thin wrapper around the inline field
// code so a lax field can decode into a scratch value. It returns "" for types
// whose value would otherwise be skipped (e.g. an unsupported map key type).
func (g *gen) valueDecoder(expr ast.Expr, hint string, nocopy, lax bool) string {
	switch t := unparen(expr).(type) {
	case *ast.Ident:
		if _, ok := g.structTypes[t.Name]; ok {
			return g.namedStruct(t.Name)
		}
	case *ast.StructType:
		return g.anonStruct(t, hint)
	case *ast.ArrayType:
		if t.Len == nil {
			return g.sliceDecoder(t.Elt, hint, nocopy, lax)
		}
		return g.arrayDecoder(t, hint, nocopy, lax)
	case *ast.MapType:
		return g.mapDecoder(t.Key, t.Value, hint, nocopy, lax)
	}
	suffix := ""
	if nocopy {
		suffix = "NoCopy"
	}
	if lax {
		suffix += "Lax"
	}
	key := "value:" + suffix + ":" + g.typeStr(expr)
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	fn := g.uniq("decode" + g.baseName(expr) + suffix + "Value")
	g.memo[key] = fn
	if isRaw(expr) {
		g.needJSON = true
	}
	if isTime(expr) {
		g.needTime = true
	}
	inner := g.field("(*v)", expr, hint, nocopy, lax)
	body := fmt.Sprintf(`func %[1]s(v *%[2]s, data []byte, i int) (int, error) {
	%[3]s
	return i, nil
}`, fn, g.typeStr(expr), inner)
	g.decoders = append(g.decoders, body)
	return fn
}

func (g *gen) scalar(dest, name string, nocopy bool) string {
	switch {
	case name == "string":
		reader := "unstable.ReadStringOrNull"
		if nocopy {
			reader = "unstable.ReadStringNoCopyOrNull"
			g.needNoCopyStr = true
			if g.destructive {
				// //lightning:destructive — unescape escaped strings into the input
				// buffer instead of allocating, aliasing the result (destroys the input).
				reader = "unstable.ReadStringDestructiveOrNull"
			}
		}
		return fmt.Sprintf(`s, end, err := %s(data, i)
if err != nil {
	return end, err
}
%s = s
i = end`, reader, dest)

	case name == "bool":
		g.needBool = true
		return fmt.Sprintf(`b, end, err := unstable.ReadBoolOrNull(data, i)
if err != nil {
	return end, err
}
%s = b
i = end`, dest)

	case name == "float32" || name == "float64":
		g.needFloat = true
		val := "f"
		if name == "float32" {
			val = "float32(f)"
		}
		return fmt.Sprintf(`f, end, err := unstable.ReadFloat64OrNull(data, i)
if err != nil {
	return end, err
}
%s = %s
i = end`, dest, val)

	case intKinds[name]:
		val := "n"
		if name != "int64" {
			val = name + "(n)"
		}
		return fmt.Sprintf(`n, end, err := unstable.ReadInt64OrNull(data, i)
if err != nil {
	return end, err
}
%s = %s
i = end`, dest, val)

	case uintKinds[name]:
		g.needUint = true
		val := "n"
		if name != "uint64" {
			val = name + "(n)"
		}
		return fmt.Sprintf(`n, end, err := unstable.ReadUint64OrNull(data, i)
if err != nil {
	return end, err
}
%s = %s
i = end`, dest, val)
	}
	panic("unreachable scalar: " + name)
}

func (g *gen) rawMessage(dest string, nocopy bool) string {
	// Default copies the bytes (like encoding/json's json.RawMessage); nocopy
	// aliases the input, which the caller must then keep unchanged.
	assign := fmt.Sprintf("%s = append(%s[:0], data[start:end]...)", dest, dest)
	if nocopy {
		assign = fmt.Sprintf("%s = data[start:end]", dest)
	}
	return fmt.Sprintf(`start := i
end, err := unstable.SkipValue(data, i)
if err != nil {
	return end, err
}
if data[start] != 'n' {
	%s
}
i = end`, assign)
}

// numberRead emits the read for a json.Number field: capture the number token's
// raw literal as a string and convert it to json.Number. nocopy aliases the input
// (json.Number's underlying type is string, so the conversion never copies).
func (g *gen) numberRead(dest string, nocopy bool) string {
	reader := "unstable.ReadNumberOrNull"
	if nocopy {
		reader = "unstable.ReadNumberNoCopyOrNull"
	}
	return fmt.Sprintf(`s, end, err := %s(data, i)
if err != nil {
	return end, err
}
%s = json.Number(s)
i = end`, reader, dest)
}

func (g *gen) timeRead(dest string, lax bool) string {
	reader := "unstable.ReadTimeOrNull"
	if lax {
		// lax accepts RFC 3339 (with 'T' or space) and numeric Unix timestamps.
		reader = "unstable.ReadTimeLaxOrNull"
	}
	return fmt.Sprintf(`t, end, err := %s(data, i)
if err != nil {
	return end, err
}
%s = t
i = end`, reader, dest)
}

func (g *gen) anyValue(dest string) string {
	g.needAny = true
	g.needBool = true
	g.needFloat = true
	decode := "unstable.DecodeValue"
	if g.compact {
		decode = "unstable.DecodeValueCompact"
	}
	return fmt.Sprintf(`val, end, err := %s(data, i)
if err != nil {
	return end, err
}
%s = val
i = end`, decode, dest)
}

func (g *gen) skipEmit() string {
	return `end, err := unstable.SkipValue(data, i)
if err != nil {
	return end, err
}
i = end`
}

// unwrapField wraps a field's normal decode (inner) so it runs against the JSON
// embedded in a string value (the "unwrap" option). unstable.Unwrap reads the
// string, unescapes it, and base64-decodes it when needed; the embedded bytes
// are then decoded by inner inside a closure that rebinds data and i to them, so
// inner's own (int, error) returns stay scoped to the embedded document while the
// outer cursor advances past the original string. A null or empty string leaves
// the field at its zero value.
func unwrapField(inner string) string {
	return fmt.Sprintf(`body, bend, berr := unstable.Unwrap(data, i)
if berr != nil {
	return bend, berr
}
if len(body) > 0 {
	if _, ierr := func(data []byte, i int) (int, error) {
		i = unstable.SkipWS(data, i)
%s
		return i, nil
	}(body, 0); ierr != nil {
		return bend, ierr
	}
}
i = bend`, inner)
}

// arrayDecoder returns (generating on first use) the decoder for a fixed-size
// array type t (t.Len != nil). It mirrors sliceDecoder but writes elements by
// index instead of appending: the array is zeroed, up to len(out) elements are
// decoded, and any beyond that are discarded. This matches encoding/json — fill
// the leading elements, leave a short JSON array's tail zero, drop a long JSON
// array's extras — and, like it, leaves the array untouched on a JSON null.
func (g *gen) arrayDecoder(t *ast.ArrayType, hint string, nocopy, lax bool) string {
	suffix := ""
	if nocopy {
		suffix = "NoCopy"
	}
	if lax {
		suffix += "Lax"
	}
	arrType := g.typeStr(t)
	key := g.prefix + g.cmark() + "array:" + suffix + ":" + arrType
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	fn := g.uniq(g.decFn("decode" + cap1(hint) + suffix + "Array" + g.csuf()))
	g.memo[key] = fn
	if isRaw(t.Elt) {
		g.needJSON = true
	}
	if isTime(t.Elt) {
		g.needTime = true
	}
	elem := g.field("(*out)[idx]", t.Elt, hint, nocopy, lax)
	// Trailing commas are rejected by the first-iteration flag, as in
	// genStructBody.
	body := fmt.Sprintf(`func %[1]s(out *%[2]s, data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, unstable.ErrTruncated
	}
	if data[i] == 'n' {
		return unstable.ExpectNull(data, i)
	}
	if data[i] != '[' {
		return i, unstable.ErrExpectArray
	}
	*out = %[2]s{}
	i++
	idx := 0
	for first := true; ; first = false {
		%[3]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == ']' {
			if first {
				return i + 1, nil
			}
			return i, unstable.ErrInvalidJSON
		}
		if idx < len(out) {
			%[4]s
		} else {
			end, err := unstable.SkipValue(data, i)
			if err != nil {
				return end, err
			}
			i = end
		}
		idx++
		%[3]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, unstable.ErrInvalidJSON
		}
		i++
	}
}`, fn, arrType, g.skipWS("i", "i"), elem)
	g.decoders = append(g.decoders, body)
	return fn
}

// batchSliceFn returns the pkg/unstable batched reader that decodes a whole
// []T in one call when T is a bare float64/int/uint kind, or "" when the
// element needs a generated loop. The batched readers keep the element loop
// next to the private scanFloat / inlined SWAR digit parse, so each number
// costs at most one call instead of the generated loop's per-element reader
// call; presize (CountArrayScalars) and null handling live inside them, and
// nocopy/lax are no-ops for these element kinds. The generic Int/Uint readers
// infer T from the *[]T argument, so no instantiation is spelled.
func batchSliceFn(elt ast.Expr) string {
	id, ok := unparen(elt).(*ast.Ident)
	if !ok {
		return ""
	}
	switch {
	case id.Name == "float64":
		return "unstable.DecodeFloat64Slice"
	case intKinds[id.Name]:
		return "unstable.DecodeIntSlice"
	case uintKinds[id.Name]:
		return "unstable.DecodeUintSlice"
	}
	return ""
}

// batchArrayFn is batchSliceFn for a fixed-size array element ([N]float64
// coordinate points and the like); the reader takes the array as a slice
// (dest[:]) and zeroes/fills/skips exactly as the generated fixed-array
// decoder did.
func batchArrayFn(elt ast.Expr) string {
	id, ok := unparen(elt).(*ast.Ident)
	if !ok {
		return ""
	}
	switch {
	case id.Name == "float64":
		return "unstable.DecodeFloat64Array"
	case intKinds[id.Name]:
		return "unstable.DecodeIntArray"
	case uintKinds[id.Name]:
		return "unstable.DecodeUintArray"
	}
	return ""
}

func (g *gen) sliceDecoder(elt ast.Expr, hint string, nocopy, lax bool) string {
	if fn := batchSliceFn(elt); fn != "" {
		return fn
	}
	suffix := ""
	if nocopy {
		suffix = "NoCopy"
	}
	if lax {
		suffix += "Lax"
	}
	key := g.prefix + g.cmark() + "slice:" + suffix + ":" + g.typeStr(elt)
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	var fn string
	if g.isStruct(elt) {
		fn = g.uniq(g.decFn("decode" + cap1(hint) + suffix + g.csuf()))
	} else {
		fn = g.uniq(g.decFn("decode" + g.baseName(elt) + suffix + "Slice" + g.csuf()))
	}
	g.memo[key] = fn
	if isRaw(elt) {
		g.needJSON = true
	}
	if isTime(elt) {
		g.needTime = true
	}
	eltStr := g.typeStr(elt)
	inner := g.field("(*out)[len(*out)-1]", elt, singular(hint)+"Entry", nocopy, lax)
	presize := g.slicePresize(elt, eltStr)
	grow := fmt.Sprintf(`var zero %[1]s
		*out = append(*out, zero)`, eltStr)
	if presize == "" {
		// No cheap counter sizes this array (see slicePresize), so instead of
		// letting append grow the backing 1→2→4→…, the *first* element brings a
		// static capacity hint sized to ~256 bytes of elements (a compile-time
		// constant via unsafe.Sizeof). Unlike the rejected counting presizes this
		// costs no extra scan: a too-large hint only wastes spare cap on a short
		// array, a too-small one regrows exactly as before. Empty arrays never
		// reach the first append, so `[]` still yields a nil slice, and decoding
		// into a reused non-nil slice keeps appending unchanged.
		grow = fmt.Sprintf(`var zero %[1]s
		if *out == nil {
			*out = make([]%[1]s, 1, max(4, 256/max(1, int(unsafe.Sizeof(zero)))))
		} else {
			*out = append(*out, zero)
		}`, eltStr)
	}
	// Trailing commas ([1,]) are rejected by the first-iteration flag, as in
	// genStructBody (see there — a rotated loop regressed cloudflare +11%).
	body := fmt.Sprintf(`func %[1]s(out *[]%[2]s, data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, unstable.ErrTruncated
	}
	if data[i] == 'n' {
		end, err := unstable.ExpectNull(data, i)
		if err != nil {
			return end, err
		}
		*out = nil
		return end, nil
	}
	if data[i] != '[' {
		return i, unstable.ErrExpectArray
	}
%[4]s	i++
	for first := true; ; first = false {
		%[5]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == ']' {
			if first {
				return i + 1, nil
			}
			return i, unstable.ErrInvalidJSON
		}
		// Grow the slice by one zero element and decode in place into the new
		// slot, so the element never lives in an escaping local (which would
		// cost a heap allocation per element for slices of structs/pointers).
		%[6]s
		%[3]s
		%[5]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, unstable.ErrInvalidJSON
		}
		i++
	}
}`, fn, eltStr, inner, presize, g.skipWS("i", "i"), grow)
	g.decoders = append(g.decoders, body)
	return fn
}

// slicePresize returns the code that allocates out to the array's element count
// before the decode loop, sparing the repeated reallocation that append-driven
// growth incurs. It presizes only when out is nil (so decoding into a reused
// slice keeps appending) and only for element types where counting is cheap
// relative to decoding: scalars (numbers/booleans) get the vectorized
// CountArrayScalars; strings, time.Time, raw messages and nested structs/slices,
// whose per-element decode is costly enough to amortize a structural pass, get
// the general CountArrayElements. Tiny fixed-cost elements where a counting scan
// would not pay for itself get no presize (an empty string keeps the loop as-is).
func (g *gen) slicePresize(elt ast.Expr, eltStr string) string {
	counter := ""
	switch t := unparen(elt).(type) {
	case *ast.Ident:
		switch {
		case t.Name == "bool" || t.Name == "float32" || t.Name == "float64" ||
			intKinds[t.Name] || uintKinds[t.Name]:
			counter = "unstable.CountArrayScalars"
		case t.Name == "string":
			counter = "unstable.CountArrayElements"
		default:
			// A *named* struct element gets the same presize decision as the
			// equivalent anonymous struct below; without this a schema that
			// names a shared record type (a FHIR coding, say) silently lost
			// the CountArrayObjects/CountArrayElements presize its inline twin
			// receives.
			if st, ok := g.structTypes[t.Name]; ok {
				switch {
				case isFlatScalarStringStruct(st):
					counter = "unstable.CountArrayObjects"
				case g.structSkipIsCheap(st):
					counter = "unstable.CountArrayElements"
				}
			}
		}
		// Presize a slice of structs only when each element is cheap to skip — a
		// flat record of scalars and strings, like a Cloudflare log line. If an
		// element nests a slice, array, map, or interface (the citm_catalog
		// performances → seatCategories → areas tree, or a citylots Feature whose
		// geometry holds a [][][]float64), the CountArrayElements scan that sizes
		// the slice must descend through every element's whole subtree — which the
		// decoder then walks again — costing far more than the reallocations
		// presize avoids. There, let the slice append instead.
	case *ast.StructType:
		switch {
		case isFlatScalarStringStruct(t):
			// A flat struct of only number/bool/string fields has no '[', ']' or
			// nested '{' of its own, so its array is sized by the cheap
			// CountArrayObjects (find ']', count '{') instead of a SkipValue per
			// element. Exact for number/bool fields (citm price entries); for string
			// fields it is a presize hint — a '['/']'/'{' inside a string value can
			// mis-size the slice but never misdecodes — which avoids the per-element
			// skipObject that dominates arrays of small {name,version}-style records
			// (update_center dependencies/developers, apache_builds jobs).
			counter = "unstable.CountArrayObjects"
		case g.structSkipIsCheap(t):
			counter = "unstable.CountArrayElements"
		}
	case *ast.MapType:
		counter = "unstable.CountArrayElements"
	case *ast.ArrayType:
		// elt is itself a slice/array. A fixed-size array element ([N]T, e.g. a
		// [3]float64 coordinate point) is skipped here: counting the outer slice
		// means structurally descending through every element's N values, which
		// the element decoders then re-parse — a full extra pass over the numbers,
		// plus zeroing the presized backing — that costs more than the doubling
		// growth it would save. Let that slice append instead.
		//
		// Otherwise presize only when the element is a leaf (a scalar slice like
		// []float64, or a slice of structs/strings) — there the count is the
		// number of inner slices and presizing avoids real reallocation. For a
		// slice *of slices of slices*, the outer dimension is typically small, yet
		// counting it would deep-scan every element that the inner decoders then
		// re-scan; skip presize there too so the outer slice just appends.
		if t.Len == nil {
			if _, nested := unparen(t.Elt).(*ast.ArrayType); !nested {
				counter = "unstable.CountArrayElements"
			}
		}
	case *ast.StarExpr:
		// A []*T presizes by the same rules as []T: the pointer wrapper changes
		// per-element allocation, not the JSON shape being counted. Without this
		// a []*Foo silently lost the presize its []Foo twin gets (the same class
		// of gap the named-struct-element case above closed).
		return g.slicePresize(t.X, eltStr)
	case *ast.SelectorExpr:
		switch {
		case isNumber(t):
			// A json.Number element is a bare number token — a scalar — so the
			// cheaper comma-counting scan sizes the slice.
			counter = "unstable.CountArrayScalars"
		case isTime(t):
			// An RFC 3339 timestamp or Unix-timestamp number never contains a
			// comma or bracket, so the cheap comma-counting scan sizes a
			// []time.Time too — far cheaper than CountArrayElements, which would
			// skipString past every element. (time-array −15%.)
			counter = "unstable.CountArrayScalars"
		case isRaw(t):
			counter = "unstable.CountArrayElements"
		}
	}
	if counter == "" {
		return ""
	}
	return fmt.Sprintf(`	if *out == nil {
		if n := %s(data, i); n > 0 {
			*out = make([]%s, 0, n)
		}
	}
`, counter, eltStr)
}

func (g *gen) mapDecoder(keyExpr, valExpr ast.Expr, hint string, nocopy, lax bool) string {
	if g.typeStr(keyExpr) != "string" {
		g.errs = append(g.errs, fmt.Errorf("unsupported map key type %s for %s", g.typeStr(keyExpr), hint))
		return ""
	}
	suffix := ""
	if nocopy {
		suffix = "NoCopy"
	}
	if lax {
		suffix += "Lax"
	}
	key := g.prefix + g.cmark() + "map:" + suffix + ":" + g.typeStr(valExpr)
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	fn := g.uniq(g.decFn("decode" + cap1(hint) + suffix + "Map" + g.csuf()))
	g.memo[key] = fn
	if isRaw(valExpr) {
		g.needJSON = true
	}
	if isTime(valExpr) {
		g.needTime = true
	}
	valStr := g.typeStr(valExpr)
	inner := g.field("val", valExpr, hint+"Value", nocopy, lax)
	// With nocopy the key aliases the input (ReadKey already returns an alias for
	// an unescaped key); otherwise it is copied so the map owns it.
	keyAssign := "m[string([]byte(key))] = val"
	if nocopy {
		keyAssign = "m[key] = val"
	}
	// Trailing commas are rejected by the first-iteration flag, as in
	// genStructBody.
	body := fmt.Sprintf(`func %[1]s(out *map[string]%[2]s, data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, unstable.ErrTruncated
	}
	if data[i] == 'n' {
		end, err := unstable.ExpectNull(data, i)
		if err != nil {
			return end, err
		}
		*out = nil
		return end, nil
	}
	if data[i] != '{' {
		return i, unstable.ErrExpectObject
	}
	i++
	m := *out
	if m == nil {
		m = make(map[string]%[2]s)
	}
	for first := true; ; first = false {
		%[4]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == '}' {
			if first {
				*out = m
				return i + 1, nil
			}
			return i, unstable.ErrInvalidJSON
		}
		%[7]s
		%[5]s
		if i >= len(data) || data[i] != ':' {
			return i, unstable.ErrExpectColon
		}
		%[6]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		var val %[2]s
		%[3]s
		%[8]s
		%[4]s
		if i >= len(data) {
			return i, unstable.ErrTruncated
		}
		if data[i] == '}' {
			*out = m
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, unstable.ErrInvalidJSON
		}
		i++
	}
}`, fn, valStr, inner, g.skipWS("i", "i"), g.skipWS("i", "ni"), g.skipWS("i", "i+1"), g.readKey(), keyAssign)
	g.decoders = append(g.decoders, body)
	return fn
}

func (g *gen) isStruct(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.ParenExpr:
		return g.isStruct(t.X)
	case *ast.StructType:
		return true
	case *ast.Ident:
		_, ok := g.structTypes[t.Name]
		return ok
	}
	return false
}

// structSkipIsCheap reports whether skipping a JSON value of expr (transitively,
// descending through structs, pointers, and named struct types) is cheap and
// bounded — which decides whether presizing a slice of this element type pays
// off (see slicePresize). A scalar, string, time, or raw leaf, and a struct
// built only from those, skips in a single bounded SIMD pass; presizing a slice
// of such records (a Cloudflare log line, say) is worth the count.
//
// A slice, array, map, or interface field is not cheap: skipping it is unbounded
// and recursive, so a counting scan would re-traverse the whole subtree the
// decoder walks anyway — far more work than the ~log2(n) backing-array
// reallocations presize would save. That covers a multi-dimensional [][][]float64
// of GeoJSON coordinates and the citm_catalog performances → seatCategories →
// areas tree alike.
func (g *gen) structSkipIsCheap(expr ast.Expr) bool {
	return g.structSkipIsCheapSeen(expr, map[string]bool{})
}

func (g *gen) structSkipIsCheapSeen(expr ast.Expr, seen map[string]bool) bool {
	switch t := unparen(expr).(type) {
	case *ast.ArrayType, *ast.MapType, *ast.InterfaceType:
		return false
	case *ast.StarExpr:
		return g.structSkipIsCheapSeen(t.X, seen)
	case *ast.StructType:
		for _, f := range t.Fields.List {
			if !g.structSkipIsCheapSeen(f.Type, seen) {
				return false
			}
		}
		return true
	case *ast.Ident:
		if seen[t.Name] {
			return true // recursive type already being walked; don't recount
		}
		seen[t.Name] = true
		if st, ok := g.structTypes[t.Name]; ok {
			return g.structSkipIsCheapSeen(st, seen)
		}
		return true // a scalar named type (int, string, ...)
	}
	return true // selector leaves (time.Time, json.RawMessage)
}

func (g *gen) baseName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.ParenExpr:
		return g.baseName(t.X)
	case *ast.Ident:
		return cap1(t.Name)
	case *ast.SelectorExpr:
		return cap1(t.Sel.Name)
	case *ast.StarExpr:
		return "Ptr" + g.baseName(t.X)
	case *ast.ArrayType:
		return "SliceOf" + g.baseName(t.Elt)
	case *ast.MapType:
		return "Map"
	case *ast.InterfaceType:
		return "Any"
	case *ast.StructType:
		return "Struct"
	}
	return "Value"
}

// unstablePkg is the import path of the shared scanner package whose exported
// helpers the generated decoders call.
const unstablePkg = "github.com/JohanLindvall/lightning/pkg/unstable"

func (g *gen) assemble(inPath string, methods []string) string {
	var body strings.Builder
	for _, m := range methods {
		body.WriteString(m)
		body.WriteString("\n\n")
	}
	for _, d := range g.decoders {
		body.WriteString(d)
		body.WriteString("\n\n")
	}
	code := body.String()

	// "time" and "encoding/json" are needed only when a spelled type actually
	// names them (time.Time, or json.RawMessage/RawValue/Number in an anonymous
	// struct, slice or map element). Decide from the generated text rather than from
	// heuristic flags: those both miss types nested inside an anonymous struct and
	// over-fire on a nocopy raw field whose decode emits no json. reference. The
	// tokens appear only as real type usages, never in the generated comments.
	imports := []string{strconv.Quote(unstablePkg)}
	if strings.Contains(code, "json.RawMessage") || strings.Contains(code, "json.RawValue") || strings.Contains(code, "json.Number") {
		imports = append(imports, `json "encoding/json"`)
	}
	if strings.Contains(code, "time.Time") {
		imports = append(imports, `"time"`)
	}
	if strings.Contains(code, "unsafe.Sizeof") {
		imports = append(imports, `"unsafe"`)
	}

	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated by the lightning generator from %s; DO NOT EDIT.\n\n", inPath)
	fmt.Fprintf(&b, "package %s\n\n", g.pkg)
	b.WriteString("import (\n")
	for _, im := range imports {
		fmt.Fprintf(&b, "\t%s\n", im)
	}
	b.WriteString(")\n\n")
	b.WriteString(code)
	return b.String()
}

// jsonTag returns the JSON key(s) from a struct tag, whether the field is
// excluded (`json:"-"`), and whether the "nocopy" and "lax" options are present.
// The name field may list several pipe-separated names
// (`json:"Status|EdgeResponseStatus"`), each of which maps the same field; any
// of them in the input fills the field. Comma-separated options follow the name,
// as in encoding/json; the recognized options are "nocopy", which makes string
// and raw fields alias the input bytes instead of copying them, and "lax", which
// makes a type mismatch on the field's value a no-op (the value is skipped and
// the field left unset) rather than an error.
func jsonTag(tag *ast.BasicLit) (names []string, skip, nocopy, lax, unwrap bool) {
	if tag == nil {
		return nil, false, false, false, false
	}
	s, err := strconv.Unquote(tag.Value)
	if err != nil {
		return nil, false, false, false, false
	}
	v := reflect.StructTag(s).Get("json")
	if v == "" {
		return nil, false, false, false, false
	}
	parts := strings.Split(v, ",")
	if parts[0] == "-" && len(parts) == 1 {
		return nil, true, false, false, false
	}
	for _, o := range parts[1:] {
		switch o {
		case "nocopy":
			nocopy = true
		case "lax":
			lax = true
		case "unwrap":
			unwrap = true
		}
	}
	for _, n := range strings.Split(parts[0], "|") {
		if n != "" {
			names = append(names, n)
		}
	}
	return names, false, nocopy, lax, unwrap
}

func cap1(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func singular(s string) string {
	switch {
	case strings.HasSuffix(s, "ies") && len(s) > 3:
		return s[:len(s)-3] + "y"
	case strings.HasSuffix(s, "s") && len(s) > 1:
		return s[:len(s)-1]
	}
	return s
}

// markReferenced records, in ref, every top-level struct type named anywhere
// within expr — through pointers, slices, maps, and anonymous structs — except
// self, so a recursive type referencing only itself is not treated as nested in
// another. It stops at named types (they are walked on their own pass), so it
// captures direct references rather than recursing through the whole graph.
func (g *gen) markReferenced(expr ast.Expr, self string, ref map[string]bool) {
	switch t := unparen(expr).(type) {
	case *ast.Ident:
		if _, ok := g.structTypes[t.Name]; ok && t.Name != self {
			ref[t.Name] = true
		}
	case *ast.StarExpr:
		g.markReferenced(t.X, self, ref)
	case *ast.ArrayType:
		g.markReferenced(t.Elt, self, ref)
	case *ast.MapType:
		g.markReferenced(t.Value, self, ref)
	case *ast.StructType:
		for _, f := range t.Fields.List {
			g.markReferenced(f.Type, self, ref)
		}
	}
}

// unparen strips any enclosing parentheses from a type expression.
func unparen(expr ast.Expr) ast.Expr {
	for {
		p, ok := expr.(*ast.ParenExpr)
		if !ok {
			return expr
		}
		expr = p.X
	}
}

func isRaw(expr ast.Expr) bool {
	if p, ok := expr.(*ast.ParenExpr); ok {
		return isRaw(p.X)
	}
	sel, ok := expr.(*ast.SelectorExpr)
	return ok && (sel.Sel.Name == "RawMessage" || sel.Sel.Name == "RawValue")
}

func isNumber(expr ast.Expr) bool {
	if p, ok := expr.(*ast.ParenExpr); ok {
		return isNumber(p.X)
	}
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Number" {
		return false
	}
	pkg, ok := sel.X.(*ast.Ident)
	return ok && pkg.Name == "json"
}

func isTime(expr ast.Expr) bool {
	if p, ok := expr.(*ast.ParenExpr); ok {
		return isTime(p.X)
	}
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Time" {
		return false
	}
	pkg, ok := sel.X.(*ast.Ident)
	return ok && pkg.Name == "time"
}

func isScalar(name string) bool {
	switch name {
	case "string", "bool", "float32", "float64":
		return true
	}
	return intKinds[name] || uintKinds[name]
}

// isBracketFreeStruct reports whether every field of t is a plain number/bool
// scalar, so the struct's JSON contains no string, '[' or ']' and no nested '{'.
// An array of such structs can be counted by CountArrayObjects (find ']', count
// '{'). Any string, pointer, array, map, embedded/named struct, or selector field
// (time.Time, json.Number, json.RawMessage) disqualifies it — its JSON could hold
// a bracket. Pipe-renamed or "-"/option-tagged fields are fine; only the field
// *type* matters. A struct with no fields is excluded (nothing to size against).
// isFlatScalarStringStruct reports whether every field is a plain number, bool or
// string identifier — no nested struct/array/map/pointer. Such a struct's JSON
// contains no '[', ']' or nested '{' of its own (only its single object braces and
// its scalar/string values), so an array of them can be counted by CountArrayObjects
// ('{' before the first ']'). Bracket-free (number/bool only) is exact; with string
// fields it is a presize hint, since a bracket inside a string value could throw the
// count off — harmless, as a miscount only mis-sizes the slice.
func isFlatScalarStringStruct(t *ast.StructType) bool {
	if t.Fields == nil || len(t.Fields.List) == 0 {
		return false
	}
	for _, f := range t.Fields.List {
		id, ok := unparen(f.Type).(*ast.Ident)
		if !ok {
			return false
		}
		switch {
		case id.Name == "string" || id.Name == "bool" ||
			id.Name == "float32" || id.Name == "float64",
			intKinds[id.Name], uintKinds[id.Name]:
		default:
			return false
		}
	}
	return true
}

var intKinds = map[string]bool{
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true, "rune": true,
}

var uintKinds = map[string]bool{
	"uint": true, "uint8": true, "uint16": true, "uint32": true,
	"uint64": true, "uintptr": true, "byte": true,
}
