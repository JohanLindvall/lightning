// Command lightning generates a fast, allocation-light json.Unmarshaler
// implementation for the struct types declared in a Go source file.
//
// Usage:
//
//	lightning <file.go> [more.go ...]
//
// It is normally run from a //go:generate directive:
//
//	//go:generate go run github.com/JohanLindvall/lightning $GOFILE
//
// or after go install github.com/JohanLindvall/lightning@latest:
//
//	//go:generate lightning $GOFILE
//
// For each input file FOO.go it writes FOO_unmarshal.go next to it, containing
// an UnmarshalJSON method for every top-level struct, slice, or map type (a
// type nested inside another gets an internal decode function instead, emitted
// by the type that reaches it; where a group of types reference each other in a
// cycle no other type enters, every member of that cycle gets the method).
// Generic types and type aliases get neither — no method can be
// declared on them — and are skipped with a warning. The generated code calls
// the shared scanner helpers in pkg/unstable.
//
// Supported field types: string, bool, all sized int/uint kinds, float32,
// float64, json.Number, time.Time, json.RawMessage, nested (named or anonymous)
// structs, slices, fixed-size arrays ([N]T), maps with string keys, pointers,
// and the empty interface any/interface{} (decoded into the usual Go
// representation of an arbitrary JSON value; an interface with methods is not
// supported). A []byte field follows encoding/json: both the stdlib's base64
// string form and a JSON array of numbers decode into it.
//
// Numbers are converted to the field's declared type without a range check,
// which is where the decoder deliberately parts company with encoding/json: an
// out-of-range integer wraps to the field's width (256 into a uint8 is 0) and a
// float32 beyond its range becomes ±Inf, where encoding/json would return an
// UnmarshalTypeError. See the scalar method for the full account.
//
// Field mapping follows the `json:"..."` struct tag: a tag renames the key,
// `json:"-"` omits the field, and a `Name|Alias` tag maps several JSON keys onto
// one field. Untagged exported fields use the Go field name as the key. Unlike
// encoding/json, key matching is EXACT and case-sensitive — there is no
// case-insensitive fallback — which is both faster and less surprising. Unknown
// members in the input are skipped.
//
// Per-type behavior is selected by `//lightning:` doc-comment directives on the
// type: `//lightning:compact` (assume whitespace-free input), `//lightning:nocopy`
// (alias string/[]byte leaves into the input instead of copying — the caller must
// then keep the input alive and unchanged), `//lightning:destructive` (unescape
// strings in place, mutating and destroying the input buffer; implies nocopy),
// and `//lightning:arena` (batch small numeric-slice backings into per-decode
// arena chunks — a retained slice pins its chunk, so it is for decode-and-discard
// use). See README.md for the full directive and struct-tag reference.
//
// Non-goals. lightning is decode-only: it generates UnmarshalJSON, not
// MarshalJSON. It operates on a complete []byte, not an io.Reader — the zero-copy
// design lets generated decoders and the pkg/json toolkit alias the input rather
// than stream it, so the whole document must be in memory. For dynamic reads,
// edits, and reshaping of JSON whose schema is not known at generate time, use the
// pkg/json toolkit.
package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <file.go> [more.go ...]\n", filepath.Base(os.Args[0]))
		os.Exit(2)
	}
	for _, in := range os.Args[1:] {
		if err := generate(in); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", in, err)
			os.Exit(1)
		}
	}
}

// generate reads one schema file and writes its FOO_unmarshal.go, reporting
// progress and non-fatal diagnostics on stderr.
func generate(inPath string) error { return generateTo(inPath, os.Stderr) }

// generateTo is generate with the diagnostic stream injected. Warnings are the
// generator's only channel for "this compiles, but it is not what your source
// asks for" (a misplaced directive, an unrecognized struct-tag option), so the
// tests have to be able to read them; passing the writer down beats a package
// global, which the parallel table test in generator_test.go would race on.
func generateTo(inPath string, warn io.Writer) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, inPath, nil, parser.SkipObjectResolution|parser.ParseComments)
	if err != nil {
		return err
	}

	g := &gen{
		warn:             warn,
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
		arenaTypes:       map[string]bool{},
		typeDirectives:   map[string][]string{},
		depthFns:         map[string]bool{},
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

	g.collectQualifiers(file)

	// Collect every top-level struct type, in source order, recording the
	// //lightning:compact / :nocopy / :destructive / :arena directives each
	// carries. Directive use is validated as it is collected: a typo'd
	// directive silently generating a decoder without the requested behavior
	// (copying where the user believes nocopy, accepting whitespace where the
	// user believes compact) is worse than failing, so an unknown
	// //lightning:* name is an error; a known directive somewhere it cannot
	// take effect only warns, since the generated code is still correct.
	attached := map[*ast.CommentGroup]bool{}
	for _, d := range file.Decls {
		gd, ok := d.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		attached[gd.Doc] = true
		// A directive on the type(...) block applies to every spec inside, so
		// unknown names in the block doc are diagnosed once, here.
		for _, dir := range lightningDirectives(gd.Doc) {
			if !knownDirectives[dir] {
				g.errs = append(g.errs, fmt.Errorf("unknown directive //lightning:%s", dir))
			}
		}
		for _, s := range gd.Specs {
			ts, ok := s.(*ast.TypeSpec)
			if !ok {
				continue
			}
			attached[ts.Doc] = true
			for _, dir := range lightningDirectives(ts.Doc) {
				if !knownDirectives[dir] {
					g.errs = append(g.errs, fmt.Errorf("unknown directive //lightning:%s on type %s", dir, ts.Name.Name))
				}
			}
			// Two declaration forms carry a struct/slice/map type that no method
			// can be attached to, and both reach the type switch below as an
			// ordinary TypeSpec — the switch reads ts.Type and never looks at
			// how the type was declared. Without these tests the generator
			// collected them like any other root and exited 0 having written a
			// package that does not compile ("cannot use generic type
			// Root[T any] without instantiation"; "invalid receiver type
			// *Root"). Skipped with a warning rather than failed, following the
			// fixed-size-array case just below: a schema file may legitimately
			// hold a generic helper or a compatibility alias that was never
			// meant as a root, and failing the whole run over an incidental
			// declaration is worse than naming it and moving on. If that leaves
			// nothing at all, the "no top-level struct, slice or map types
			// found" error below still fires.
			if ts.TypeParams != nil {
				g.warnf("generic type %s gets no UnmarshalJSON: a decoder cannot be generated for a type with type parameters", ts.Name.Name)
				continue
			}
			if ts.Assign.IsValid() {
				// An alias to a struct literal stays resolvable as a FIELD
				// type: `type Legacy = struct{...}` is not a defined type, so
				// no method can be attached, but a generated decoder taking a
				// *Legacy is perfectly legal and that is how such an alias
				// already decoded. Registering it without adding it to g.order
				// keeps that working while denying it the method it cannot
				// have.
				if st, isStruct := ts.Type.(*ast.StructType); isStruct {
					g.structTypes[ts.Name.Name] = st
				}
				g.warnf("type alias %s gets no UnmarshalJSON: a method can only be declared on a defined type (drop the '=' to define one)", ts.Name.Name)
				continue
			}
			dirs := lightningDirectives(gd.Doc, ts.Doc)
			isStruct := false
			switch t := ts.Type.(type) {
			case *ast.StructType:
				g.structTypes[ts.Name.Name] = t
				g.order = append(g.order, ts.Name.Name)
				isStruct = true
			case *ast.ArrayType:
				// A named slice root type (type Foo []T) for array-root documents.
				if t.Len == nil {
					g.sliceTypes[ts.Name.Name] = t
					g.order = append(g.order, ts.Name.Name)
				} else {
					g.warnDirectives(dirs, ts.Name.Name, "a fixed-size array type gets no UnmarshalJSON")
					continue
				}
			case *ast.MapType:
				// A named map root type (type Foo map[string]V) for object-root
				// documents that are data maps rather than fixed records.
				g.mapTypes[ts.Name.Name] = t
				g.order = append(g.order, ts.Name.Name)
			default:
				g.warnDirectives(dirs, ts.Name.Name, "not a struct, slice or map type")
				continue
			}
			if len(dirs) > 0 {
				g.typeDirectives[ts.Name.Name] = dirs
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
			if hasDirective("lightning:arena", gd.Doc, ts.Doc) {
				g.arenaTypes[ts.Name.Name] = true
			}
			// nocopy is consumed only by slice/map roots (a struct's aliasing is
			// governed per field by the ,nocopy tag), so on a struct root the
			// bare directive silently did nothing.
			if isStruct && g.nocopyTypes[ts.Name.Name] && !g.destructiveTypes[ts.Name.Name] {
				g.warnf("//lightning:nocopy on struct type %s has no effect; use the ,nocopy tag on its string/raw fields", ts.Name.Name)
			}
		}
	}
	// Everything above reads directives the parser *attached* to a type
	// declaration. A //lightning: comment anywhere else reaches none of it: it
	// selects nothing and it escapes the typo check, both in silence. The two
	// ways that happens are easy to write by accident — a blank line between the
	// directive and its type is enough for go/ast to stop treating the comment as
	// that type's doc, and a directive stranded above the package clause belongs
	// to the file — and the cost is invisible (an //lightning:arena that never
	// fires is ~29k allocations per decode back on marine_ik-shaped input). Report
	// it the way a misplaced-but-attached directive is reported: a warning, since
	// the generated code is still correct, just not what the source asks for.
	for _, cg := range file.Comments {
		if attached[cg] {
			continue
		}
		for _, c := range cg.List {
			dir, ok := directiveIn(c)
			if !ok {
				continue
			}
			pos := g.fset.Position(c.Pos())
			if knownDirectives[dir] {
				g.warnf("//lightning:%s at %s has no effect: it is not attached to a type declaration (a blank line between a directive and its type detaches it)", dir, pos)
			} else {
				g.warnf("//lightning:%s at %s is not a known directive, and is not attached to a type declaration", dir, pos)
			}
		}
	}
	if len(g.order) == 0 {
		return errors.New("no top-level struct, slice or map types found")
	}
	g.checkReservedNames()

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
	refs := make(map[string]map[string]bool, len(g.order))
	referenced := map[string]bool{}
	for _, name := range g.order {
		r := map[string]bool{}
		if st, ok := g.structTypes[name]; ok {
			for _, f := range st.Fields.List {
				g.markReferenced(f.Type, name, r)
			}
		} else if at, ok := g.sliceTypes[name]; ok {
			g.markReferenced(at.Elt, name, r)
		} else if mt, ok := g.mapTypes[name]; ok {
			g.markReferenced(mt.Value, name, r)
		}
		refs[name] = r
		for u := range r {
			referenced[u] = true
		}
	}
	emitted := g.entryTypes(refs, referenced)

	// A referenced type's decoder is emitted by the root that reaches it, under
	// that root's directives — its own directives are never read, so tell the
	// user rather than silently generating the plain variant.
	for _, name := range g.order {
		if !emitted[name] {
			g.warnDirectives(g.typeDirectives[name], name,
				"the type is nested inside another type; its decoder follows the enclosing root's directives")
		}
	}

	// Decide, before any code is emitted, which decoders must thread a recursion
	// depth: only those belonging to a type graph that loops back on itself.
	g.computeDepthThreading()

	var methods []string
	for _, name := range g.order {
		if !emitted[name] {
			continue // nested in another type; its decoder is emitted there
		}
		g.compact = g.compactTypes[name]
		g.destructive = g.destructiveTypes[name]
		// A destructive root aliases the very buffer it decodes into, so it is nocopy
		// too; a plain //lightning:nocopy root (slice/map) aliases its keys/elements.
		g.nocopy = g.destructive || g.nocopyTypes[name]
		g.arena = g.arenaTypes[name]
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
	// Best effort: a generated file is not worth failing over an unwritable
	// diagnostic stream. Discarded explicitly because warn is an io.Writer, so
	// errcheck's built-in os.Stderr exclusion does not apply here.
	_, _ = fmt.Fprintf(warn, "wrote %s\n", outPath)
	return nil
}

// gen holds the state accumulated while walking the type graph.
type gen struct {
	warn        io.Writer // diagnostic stream (os.Stderr in a real run)
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
	//   //lightning:arena        -> arenaTypes        (batch small scalar-slice backings
	//                                                   into per-decode arena chunks)
	compactTypes     map[string]bool
	nocopyTypes      map[string]bool
	destructiveTypes map[string]bool
	arenaTypes       map[string]bool
	typeDirectives   map[string][]string // every directive a type carries, for misplacement warnings

	// Working flags for the root type currently being generated, derived from the
	// directive sets above.
	compact     bool
	destructive bool   // //lightning:destructive: unescape strings in place
	nocopy      bool   // //lightning:nocopy root, or a destructive root (which aliases what it decodes)
	arena       bool   // //lightning:arena: thread a per-decode unstable.Arena through the decoders
	pathFrag    string // import path sanitized into an identifier fragment
	prefix      string // current per-type prefix for generated function names

	// Recursion-depth threading, for schemas whose type graph has a cycle (a tree
	// node: type Node struct { Kids []*Node }). Such a schema's decoders call each
	// other in a loop, so a deeply nested document recurses once per level — and a
	// Go stack overflow is fatal, beyond recover's reach. threadDepth marks the
	// named types that can reach a cycle; every decoder generated for one of them
	// takes an extra `depth int`, and the struct decoders among them refuse to
	// descend past unstable.MaxDepth. A schema with no cycle (the overwhelmingly
	// common case, and every benchmark) is unaffected: not one signature changes,
	// so the hot paths are byte-identical. See computeDepthThreading.
	threadDepth map[string]bool // named type -> its decoders carry depth
	depthFns    map[string]bool // generated function name -> takes a depth param
	depthArg    string          // what a call site inside the current body passes

	// The identifiers the *input file* uses to qualify encoding/json and time,
	// and whether the generated file ends up naming them. A qualifier is ""
	// when the input does not import that package at all, which is what makes
	// `mypkg.RawMessage` a foreign type rather than encoding/json's — see
	// isRaw/isNumber/isTime. The need* flags are set where a qualified name is
	// actually emitted (typeStr, numberRead) and decide the generated file's
	// import block; see assemble for why they are not derived from the text.
	jsonQual   string
	timeQual   string
	needJSON   bool
	needTime   bool
	needUnsafe bool

	decoders []string        // generated decoder functions, in creation order
	errs     []error         // generation errors; reported together after the walk
	warned   map[string]bool // diagnostics already reported (see warnf)

}

// collectQualifiers records how the input file spells encoding/json and time.
// The generated file is a *sibling* with its own import block, so it must both
// recognize the schema's qualifier (an aliased `import ej "encoding/json"` makes
// ej.RawMessage the real thing and json.RawMessage a foreign type) and re-import
// under that same alias, since every type it emits is printed from the schema's
// own AST. A blank or dot import contributes no qualifier: `_` names nothing and
// a dot-imported RawMessage arrives as a bare identifier, which the generator
// does not resolve.
func (g *gen) collectQualifiers(file *ast.File) {
	for _, imp := range file.Imports {
		path, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			continue
		}
		name := path[strings.LastIndexByte(path, '/')+1:]
		if imp.Name != nil {
			name = imp.Name.Name
		}
		if name == "_" || name == "." {
			continue
		}
		switch path {
		case "encoding/json":
			g.jsonQual = name
		case "time":
			g.timeQual = name
		}
	}
}

// typeStr renders a type expression as the Go source text the generated file
// will contain — and, on the way through, records which of the two importable
// packages that text names. Every emitted type name funnels through here (a
// decoder's parameter type, a `var zero T`, a `new(T)`, a make), so it is the
// one place that can decide the imports from what is genuinely emitted rather
// than from a scan of the finished output. Uses that do not emit — memo keys,
// the map-key-type comparison, an error message — are paired with a real
// emission or with a failed run that writes no file, so noting them is harmless.
func (g *gen) typeStr(e ast.Expr) string {
	g.noteQualifiers(e)
	var b strings.Builder
	_ = printer.Fprint(&b, g.fset, e) // writing to a strings.Builder never fails
	return b.String()
}

// noteQualifiers marks the packages a type expression names as needed by the
// generated file. It walks the whole expression rather than testing its head,
// because the type text may nest a qualified name arbitrarily deep — an
// anonymous struct's field (struct{ W time.Time }, emitted verbatim as a
// decoder's parameter type) is exactly the case a per-field flag missed.
func (g *gen) noteQualifiers(e ast.Expr) {
	ast.Inspect(e, func(n ast.Node) bool {
		sel, ok := n.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		if id, ok := sel.X.(*ast.Ident); ok {
			switch id.Name {
			case g.jsonQual:
				g.needJSON = true
			case g.timeQual:
				g.needTime = true
			}
		}
		return true
	})
}

// reservedIdents are the identifiers the generated code puts in scope around the
// schema's own type names. A schema type sharing one of these names is captured
// by the generated declaration and stops being a type where the decoders name it
// — `var zero zero` makes the following `make([]zero, ...)` illegal, a decoder
// body that says `new(data)` or `var zero data` resolves data to the []byte
// parameter, and a package-level `unstable` collides outright with the scanner
// import. Every one of those is the generator's characteristic failure: it exits
// 0 having written a package that does not compile, so the collision is reported
// as an error at generate time instead (see checkReservedNames).
//
// The set is derived by reading the format strings in this file — the only place
// generated text is written — in three groups:
//
//  1. Parameters and locals every emitted decoder declares: the signatures of
//     genUnmarshal / genStructBody / valueDecoder / arrayDecoder / sliceDecoder /
//     mapDecoder, and the locals in readKey, skipUnknown, scalar, rawMessage,
//     numberRead, timeRead, anyValue, skipEmit, laxField, unwrapField,
//     callDecoder and the pointer arm of field. Names this generator prefixes
//     with "lightning" (lightningArrStart, the dispatch labels) are excluded —
//     the prefix exists precisely so they cannot collide.
//  2. The packages the generated file imports: an import name and a
//     package-level declaration may not share an identifier. The schema's own
//     encoding/json and time qualifiers are checked separately in
//     checkReservedNames, since they are whatever the input file imports them as.
//  3. The predeclared identifiers the generated code names of its own accord. A
//     package-level declaration shadows the universe block, so `type max
//     struct{...}` breaks the capacity hint's max(4, ...) and `type nil
//     struct{...}` breaks every `return end, nil`. Predeclared names that reach
//     the output only as an echo of the schema's own type text (bool, float64,
//     uint16, any — typeStr prints what the field declares) are deliberately
//     not here: nothing in the generated code puts them in scope.
//
// Groups 1 and 2 were cross-checked against reality rather than only read off
// the templates — parsing the decoders generated for the conformance suite and
// all 30 benchmark schemas yields exactly this list of declared parameters,
// locals and imports (plus the blank identifier and lightningArrStart).
//
// Keep it in step with the templates: adding a local to an emitted body means
// adding its name here.
var reservedIdents = map[string]bool{
	// 1. decoder parameters and locals
	"a": true, "b": true, "bend": true, "berr": true, "body": true, "data": true,
	"depth": true, "end": true, "err": true, "f": true, "first": true, "i": true,
	"idx": true, "ierr": true, "k": true, "key": true, "ks": true, "lax": true,
	"m": true, "n": true, "ni": true, "out": true, "s": true, "start": true,
	"t": true, "v": true, "val": true, "zero": true,
	// 2. imports of the generated file
	"unstable": true, "unsafe": true,
	// 3. predeclared identifiers the generated code uses
	"append": true, "byte": true, "cap": true, "error": true, "false": true,
	"int": true, "len": true, "make": true, "max": true, "new": true,
	"nil": true, "string": true, "true": true,
}

// checkReservedNames reports every collected type whose name the generated code
// would capture (see reservedIdents). It is an error, not a warning: unlike a
// misplaced directive the result does not compile, and the fix — rename the type
// — is the user's to make. Renaming the generated locals instead was considered
// and rejected: it would rewrite every committed generated file to buy nothing.
func (g *gen) checkReservedNames() {
	for _, name := range g.order {
		// jsonQual/timeQual are the identifiers the *input* file imports
		// encoding/json and time under, and the generated file re-imports them
		// under the same name; "" means the schema does not import that package,
		// and no type name is empty.
		if reservedIdents[name] || name == g.jsonQual || name == g.timeQual {
			g.errs = append(g.errs, fmt.Errorf("type name %q collides with an identifier used by the generated code; rename it", name))
		}
	}
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
// knownDirectives is the set of //lightning:* directive names the generator
// understands; anything else in the lightning: namespace is a typo and an error.
var knownDirectives = map[string]bool{
	"compact":     true,
	"nocopy":      true,
	"destructive": true,
	"arena":       true,
}

// directiveIn returns the name of the //lightning:* directive a single comment
// line carries — the part after the colon — normalized the way hasDirective
// matches: whitespace anywhere is collapsed first, so "//lightning:compact",
// "// lightning:compact " and "// lightning: compact" are the one directive. A
// line that merely mentions a directive mid-sentence does not match, since the
// collapsed text has to *start* with "lightning:".
func directiveIn(c *ast.Comment) (string, bool) {
	text := strings.Join(strings.Fields(strings.TrimPrefix(c.Text, "//")), "")
	return strings.CutPrefix(text, "lightning:")
}

// lightningDirectives returns the name of every //lightning:* directive in the
// given comment groups.
func lightningDirectives(groups ...*ast.CommentGroup) []string {
	var out []string
	for _, cg := range groups {
		if cg == nil {
			continue
		}
		for _, c := range cg.List {
			if dir, ok := directiveIn(c); ok {
				out = append(out, dir)
			}
		}
	}
	return out
}

// warnf reports a non-fatal diagnostic on the generator's diagnostic stream:
// the generated code is correct, but not what the source appears to ask for.
//
// Identical messages are reported once. A field's diagnostics are produced while
// its struct's decoder is generated, and one struct yields several decoders — a
// shared named type reached from two roots, or the compact/destructive/arena
// variants of one root — so without this a single mistyped tag option would be
// echoed once per variant.
func (g *gen) warnf(format string, args ...any) {
	msg := fmt.Sprintf("warning: "+format+"\n", args...)
	if g.warned[msg] {
		return
	}
	if g.warned == nil {
		g.warned = map[string]bool{}
	}
	g.warned[msg] = true
	_, _ = fmt.Fprint(g.warn, msg) // best effort; see the write at the end of generateTo
}

// warnDirectives warns that every directive in dirs is ineffective on the named
// type, with the reason it cannot take effect there.
func (g *gen) warnDirectives(dirs []string, typeName, reason string) {
	for _, dir := range dirs {
		if knownDirectives[dir] {
			g.warnf("//lightning:%s on type %s has no effect: %s", dir, typeName, reason)
		}
	}
}

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

// cmark and csuf distinguish a decoder's compact, destructive and arena variants
// from their plain counterparts in, respectively, memo keys and generated function
// names, so the same type reached from roots with different directive combinations
// yields a distinct decoder per combination.
func (g *gen) cmark() string {
	var m string
	if g.compact {
		m += "compact:"
	}
	if g.destructive {
		m += "destructive:"
	}
	if g.arena {
		m += "arena:"
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
	if g.arena {
		s += "Arena"
	}
	return s
}

// arenaParam and arenaArg thread the per-decode arena through an
// //lightning:arena root's decoders, the way depthParam/depthArgFor thread the
// recursion depth — except unconditionally within the variant: every decoder
// generated for an arena root takes the parameter, so call sites need no
// per-function bookkeeping, and a root without the directive (arenaParam == "")
// generates byte-identical code. The batched scalar-slice readers are the only
// consumers (see batchSliceFn's Arena rerouting in sliceDecoder); decoders the
// arena merely passes through carry it unused.
func (g *gen) arenaParam() string {
	if g.arena {
		return ", a *unstable.Arena"
	}
	return ""
}

func (g *gen) arenaArg() string {
	if g.arena {
		return ", a"
	}
	return ""
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
	// The root call seeds the recursion depth at 0 (g.depthArg's zero value outside
	// any body), and contributes nothing when the schema has no cycle. An arena root
	// likewise seeds the decode's arena: the method declares it (arenaDecl) and every
	// decoder threads the pointer down.
	g.depthArg = "0"
	var arenaDecl, rootArenaArg string
	if g.arena {
		arenaDecl = "var a unstable.Arena\n\t"
		rootArenaArg = ", &a"
	}
	switch {
	case g.sliceTypes[name] != nil:
		at := g.sliceTypes[name]
		fn := g.sliceDecoder(at.Elt, name, nocopy, false, true)
		call = fmt.Sprintf("%s((*[]%s)(v), data, i%s%s)", fn, g.typeStr(at.Elt), g.depthArgFor(fn), rootArenaArg)
	case g.mapTypes[name] != nil:
		mt := g.mapTypes[name]
		fn := g.mapDecoder(mt.Key, mt.Value, name, nocopy, false)
		call = fmt.Sprintf("%s((*map[string]%s)(v), data, i%s%s)", fn, g.typeStr(mt.Value), g.depthArgFor(fn), rootArenaArg)
	default:
		fn := g.namedStruct(name)
		call = fmt.Sprintf("%s(v, data, i%s%s)", fn, g.depthArgFor(fn), rootArenaArg)
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
	if g.arena {
		doc += `
// Small numeric-slice backings are carved from shared per-call arena chunks
// (//lightning:arena): retaining such a slice keeps its whole chunk reachable,
// so decode, process, and discard the result together rather than holding on
// to a few small slices from it.`
	}
	return fmt.Sprintf(`%[2]s
func (v *%[1]s) UnmarshalJSON(data []byte) error {
	%[4]si := unstable.SkipWS(data, 0)
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
}`, name, doc, call, arenaDecl)
}

// namedStruct returns (generating on first use) the decoder for a named struct.
func (g *gen) namedStruct(name string) string {
	key := g.prefix + g.cmark() + "named:" + name
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	// Through g.uniq like every other emitter (anonStruct, valueDecoder,
	// arrayDecoder, sliceDecoder, mapDecoder), not a bare g.used write: reserving
	// the name without first READING the reservation set let a named struct take
	// a name an anonymous struct had already been given — `type Item` plus a
	// field `Item struct{...}` both wanted decodeItem, and the loser's rename
	// never happened, so the generator exited 0 having emitted the function
	// twice. It was order-dependent (uniq renames whichever emitter runs second),
	// which made a plain field reorder flip a working schema to one that does not
	// compile.
	fn := g.uniq(g.decFn("decode" + name + g.csuf()))
	g.memo[key] = fn // set before body so recursive types terminate
	// Likewise set before the body: a recursive schema calls back into fn while
	// fn's own body is still being generated, and that call must spell the same
	// signature.
	g.markDepthFn(fn, g.threadDepth[name])
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
	g.markDepthFn(fn, g.exprThreadsDepth(t))
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
		tag := jsonTag(f.Tag)
		tagNames, nocopy, lax, unwrap := tag.names, tag.nocopy, tag.lax, tag.unwrap
		if len(tag.unknown) > 0 {
			g.warnTagOptions(tag.unknown, fieldLabel(f))
		}
		if tag.skip {
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

// fieldArm is one field's dispatch entry: the JSON names it answers to (already
// resolved through Go's promotion rules) and the code that decodes it.
type fieldArm struct {
	keys []string
	code string
}

// maxInlineCmp is cmd/compile's maxRewriteLen: the longest string constant it will
// compare with inline word loads (2 * RegSize on amd64/arm64). A comparison against
// a longer constant becomes a CALL to runtime.memequal, which also forces the live
// registers to be spilled and reloaded around it and gives the caller a stack frame.
const maxInlineCmp = 16

// skipUnknown is the code that skips an object member whose key matched no field.
const skipUnknown = `end, err := unstable.SkipValue(data, i)
		if err != nil {
			return end, err
		}
		i = end`

// keyDispatch emits the statement that matches key against each field's names and
// runs that field's decode, skipping the value when nothing matches.
//
// When every name fits maxInlineCmp it is a plain `switch key`, which is already
// optimal: cmd/compile buckets the cases by length itself and compares each with
// inline word loads. That keeps the emitted code byte-identical for the great
// majority of schemas.
//
// A name longer than maxInlineCmp is the case worth restructuring, because `switch
// key` compiles it to a runtime.memequal call — cloudflare's 45-field decoder made 10
// such calls, and memeqbody plus memequal were ~3.5% of its profile. Then the dispatch
// becomes `switch len(key)`, and inside each length bucket the names are compared in
// <=maxInlineCmp-byte chunks, so no comparison ever calls memequal. Buckets whose
// names all fit keep a nested `switch key` so short keys are dispatched exactly as
// well as before.
//
// Non-matching keys reach the skip through a goto rather than a copy of it per
// bucket: a wide struct has many buckets, and duplicating the skip in each would add
// real code to the hot loop for no benefit. The matched path costs nothing — it falls
// out of the switch — which is why this is a goto and not a `handled` flag.
func (g *gen) keyDispatch(arms []fieldArm) string {
	maxLen := 0
	for _, a := range arms {
		for _, k := range a.keys {
			if len(k) > maxLen {
				maxLen = len(k)
			}
		}
	}
	if maxLen <= maxInlineCmp {
		var cases strings.Builder
		for _, a := range arms {
			quoted := make([]string, len(a.keys))
			for j, k := range a.keys {
				quoted[j] = strconv.Quote(k)
			}
			fmt.Fprintf(&cases, "\tcase %s:\n%s\n", strings.Join(quoted, ", "), a.code)
		}
		return fmt.Sprintf("switch key {\n%s\n\t\tdefault:\n\t\t\t%s\n\t\t}", cases.String(), skipUnknown)
	}

	// Group each field's names by length. A field whose names differ in length (a
	// pipe-separated alternate spelling) appears in more than one bucket, which
	// duplicates its decode code — rare, and the alternative is a second dispatch.
	byLen := map[int][]fieldArm{}
	var lens []int
	for _, a := range arms {
		perLen := map[int][]string{}
		for _, k := range a.keys {
			perLen[len(k)] = append(perLen[len(k)], k)
		}
		for l, ks := range perLen {
			if _, seen := byLen[l]; !seen {
				lens = append(lens, l)
			}
			byLen[l] = append(byLen[l], fieldArm{keys: ks, code: a.code})
		}
	}
	slices.Sort(lens)

	var b strings.Builder
	b.WriteString("switch len(key) {\n")
	for _, l := range lens {
		fmt.Fprintf(&b, "\t\tcase %d:\n", l)
		if l <= maxInlineCmp {
			b.WriteString("\t\t\tswitch key {\n")
			for _, a := range byLen[l] {
				quoted := make([]string, len(a.keys))
				for j, k := range a.keys {
					quoted[j] = strconv.Quote(k)
				}
				fmt.Fprintf(&b, "\t\t\tcase %s:\n%s\n", strings.Join(quoted, ", "), a.code)
			}
			b.WriteString("\t\t\tdefault:\n\t\t\t\tgoto lightningSkipKey\n\t\t\t}\n")
			continue
		}
		for n, a := range byLen[l] {
			kw := "if"
			if n > 0 {
				kw = "} else if"
			}
			conds := make([]string, len(a.keys))
			for j, k := range a.keys {
				conds[j] = chunkedKeyEq(k)
			}
			fmt.Fprintf(&b, "\t\t\t%s %s {\n%s\n", kw, strings.Join(conds, " || "), a.code)
		}
		b.WriteString("\t\t\t} else {\n\t\t\t\tgoto lightningSkipKey\n\t\t\t}\n")
	}
	b.WriteString("\t\tdefault:\n\t\t\tgoto lightningSkipKey\n\t\t}\n")
	// The matched path jumps clear of the skip. The skip sits in its own block so
	// that jump does not cross a variable declaration, which Go forbids.
	b.WriteString("\t\tgoto lightningKeyDone\n")
	b.WriteString("\tlightningSkipKey:\n\t\t{\n\t\t\t" + skipUnknown + "\n\t\t}\n")
	b.WriteString("\tlightningKeyDone:")
	return b.String()
}

// chunkedKeyEq compares key against the constant k without ever exceeding
// maxInlineCmp bytes in one comparison, so cmd/compile keeps every piece as inline
// word loads instead of a runtime.memequal call. The caller has already established
// len(key) == len(k) by switching on it, which is also what lets the slice bounds be
// proved and their checks elided.
func chunkedKeyEq(k string) string {
	if len(k) <= maxInlineCmp {
		return "key == " + strconv.Quote(k)
	}
	var parts []string
	for off := 0; off < len(k); off += maxInlineCmp {
		end := off + maxInlineCmp
		if end >= len(k) {
			parts = append(parts, fmt.Sprintf("key[%d:] == %s", off, strconv.Quote(k[off:])))
			break
		}
		parts = append(parts, fmt.Sprintf("key[%d:%d] == %s", off, end, strconv.Quote(k[off:end])))
	}
	return strings.Join(parts, " && ")
}

func (g *gen) genStructBody(fn, paramType string, st *ast.StructType) {
	// A struct frame is the one that repeats as a recursive document nests, so
	// inside this body a callee's depth argument is depth+1 (see enterBody).
	prevDepth := g.enterBody(fn, true)
	defer func() { g.depthArg = prevDepth }()

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

	// Collect the winning key names and decode code for each field, in collection
	// order, then let dispatch decide how to match a key against them.
	var arms []fieldArm
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
		arms = append(arms, fieldArm{keys: won, code: code})
	}
	dispatch := g.keyDispatch(arms)

	// The loop top is reached only after '{' (first iteration) or after a
	// comma, and a '}' after a member returns from the post-value check — so a
	// '}' at the loop top on a non-first iteration is exactly a trailing comma
	// ({"a":1,}), rejected as encoding/json does. The `first` flag keeps the
	// loop shape byte-identical to the lenient form otherwise (one register
	// move per member): a *rotated* loop (closer checked before the loop and
	// after each member only) measured cloudflare +11% — the wide decoder is
	// that layout-sensitive — so don't restructure, flag.
	body := fmt.Sprintf(`func %[1]s(v %[2]s, data []byte, i int%[9]s) (int, error) {
	%[10]sif i >= len(data) {
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
		%[3]s
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
}`, fn, paramType, dispatch, g.skipWS("i", "i"), g.skipWS("i", "ni"), g.skipWS("i", "i+1"), g.skipWS("i", "i"), g.readKey(), g.depthParam(fn)+g.arenaParam(), g.depthGuard(fn))
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
		inner := g.field("(*"+dest+")", t.X, hint, nocopy, lax)
		// The allocation is guarded so a reused non-nil target decodes into the
		// existing pointee, matching encoding/json ("Unmarshal unmarshals the
		// JSON into the value pointed at by the pointer") and sparing pointer-
		// dense schemas one allocation per field per decode on reuse. Like the
		// stdlib, a reused pointee keeps field values the document omits.
		return fmt.Sprintf(`if data[i] == 'n' {
	end, err := unstable.ExpectNull(data, i)
	if err != nil {
		return end, err
	}
	%[1]s = nil
	i = end
} else {
	if %[1]s == nil {
		%[1]s = new(%[2]s)
	}
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
		if g.isRaw(t) {
			return g.rawMessage(dest, nocopy)
		}
		if g.isNumber(t) {
			return g.numberRead(dest, nocopy)
		}
		if g.isTime(t) {
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
		return g.callDecoder(dest, g.sliceDecoder(t.Elt, hint, nocopy, lax, false))

	case *ast.MapType:
		fn := g.mapDecoder(t.Key, t.Value, hint, nocopy, lax)
		if fn == "" {
			return g.skipEmit()
		}
		return g.callDecoder(dest, fn)

	case *ast.InterfaceType:
		// Only the empty interface holds an arbitrary JSON value: DecodeValue
		// returns an `any`, which assigns to `any` and to nothing else. Every
		// *ast.InterfaceType used to route here, so a field declared
		// `interface{ Foo() }` generated `v.E = val` with val an any — the
		// generator exiting 0 on a package that does not compile.
		if !isAnyInterface(t) {
			return g.unsupportedf("unsupported type %s for %s: only the empty interface (any) can hold an arbitrary JSON value", g.typeStr(t), dest)
		}
		return g.anyValue(dest)

	default:
		return g.unsupportedf("unsupported type %s for %s", g.typeStr(expr), dest)
	}
}

// callDecoder invokes a (struct/slice/map) decoder function on &dest, passing the
// recursion depth on when fn is one of the decoders that thread it, and the arena
// when generating an //lightning:arena variant (whose decoders all take it).
func (g *gen) callDecoder(dest, fn string) string {
	return fmt.Sprintf(`end, err := %s(&%s, data, i%s%s)
if err != nil {
	return end, err
}
i = end`, fn, dest, g.depthArgFor(fn), g.arenaArg())
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
end, err := %[2]s(&lax, data, i%[4]s%[5]s)
if err != nil {
	end, err = unstable.SkipValue(data, i)
	if err != nil {
		return end, err
	}
} else {
	%[3]s = lax
}
i = end`, g.typeStr(expr), fn, dest, g.depthArgFor(fn), g.arenaArg())
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
			return g.sliceDecoder(t.Elt, hint, nocopy, lax, false)
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
	// The memo key and function name carry the per-root prefix and the
	// compact/destructive/arena marker exactly like every other emitter
	// (named/anon/array/slice/map): without them, roots with different
	// directives sharing a lax field type shared one decoder — a plain root
	// could pick up a destructive sibling's in-place unescape (mutating the
	// caller's buffer) or a compact sibling's whitespace rejection, and two
	// files with the same lax field type generated colliding top-level names.
	key := g.prefix + g.cmark() + "value:" + suffix + ":" + g.typeStr(expr)
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	fn := g.uniq(g.decFn("decode" + g.baseName(expr) + suffix + "Value" + g.csuf()))
	g.memo[key] = fn
	g.markDepthFn(fn, g.exprThreadsDepth(expr))
	prevDepth := g.enterBody(fn, false)
	inner := g.field("(*v)", expr, hint, nocopy, lax)
	g.depthArg = prevDepth
	body := fmt.Sprintf(`func %[1]s(v *%[2]s, data []byte, i int%[4]s) (int, error) {
	%[3]s
	return i, nil
}`, fn, g.typeStr(expr), inner, g.depthParam(fn)+g.arenaParam())
	g.decoders = append(g.decoders, body)
	return fn
}

// scalar emits the read for a leaf field of a builtin kind.
//
// Range behavior differs from encoding/json, deliberately and in the decoder's
// favor — record it here, where the conversions are emitted, because nothing in
// the generated code says so:
//
//   - Integers are read as int64/uint64 and then CONVERTED to the field's width,
//     so an out-of-range value wraps rather than being rejected: 256 into a uint8
//     is 0, 300 into an int8 is 44, and a value past int64 wraps in the reader
//     itself (9223372036854775808 into an int64 is -9223372036854775808).
//     encoding/json returns an UnmarshalTypeError for each of those and leaves
//     the field alone. A negative literal in an unsigned field is still an error
//     (the reader rejects the '-'), and a fractional or exponent literal in an
//     integer field fails as a syntax error at the container level, since the
//     reader stops at the '.'/'e' the enclosing loop then does not accept.
//   - float32 is read as a float64 and converted, so a value beyond float32's
//     range becomes ±Inf (1e39) instead of an UnmarshalTypeError. float64 itself
//     rejects an out-of-range literal (1e400) in the reader.
//
// The reason is that a range check costs a compare and a branch per number on
// paths that are otherwise a few instructions per digit, to detect input a
// schema-matched producer cannot emit. It is a documented trade, not an
// oversight: a decoder that must reject out-of-range numbers has to check the
// decoded field itself.
func (g *gen) scalar(dest, name string, nocopy bool) string {
	switch {
	case name == "string":
		reader := "unstable.ReadStringOrNull"
		if nocopy {
			reader = "unstable.ReadStringNoCopyOrNull"
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
		return fmt.Sprintf(`b, end, err := unstable.ReadBoolOrNull(data, i)
if err != nil {
	return end, err
}
%s = b
i = end`, dest)

	case name == "float32" || name == "float64":
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
	// Defensive: every caller gates on isScalar, so this fires only if isScalar
	// and the switch above drift apart. Routed through the error accumulator so
	// the user gets the uniform "unsupported type" diagnostic, not a stack trace.
	return g.unsupportedf("internal: scalar kind %q not handled for %s", name, dest)
}

// rawMessage emits the read for a json.RawMessage field: the value's span is
// found with SkipValue and handed over verbatim.
//
// A JSON null is captured like any other value — the field ends up holding the
// four bytes "null" — which is what encoding/json does: json.RawMessage is an
// Unmarshaler, and "to unmarshal JSON into a value implementing Unmarshaler,
// Unmarshal calls that value's UnmarshalJSON method, INCLUDING when the input is
// a JSON null", so RawMessage.UnmarshalJSON stores the literal. Leaving the
// field untouched on null (what this emitted before) broke that two ways: a
// fresh target could not tell a null value from an absent key — the very
// distinction a raw field exists to preserve — and a target being reused across
// documents silently kept the PREVIOUS document's value where the new one says
// null, which is the reuse pattern this library encourages elsewhere.
//
// Default copies the bytes (like encoding/json's json.RawMessage); nocopy
// aliases the input, which the caller must then keep unchanged. Both spellings
// capture null identically — the aliased "null" is four bytes of the input like
// any other span.
func (g *gen) rawMessage(dest string, nocopy bool) string {
	assign := fmt.Sprintf("%s = append(%s[:0], data[start:end]...)", dest, dest)
	if nocopy {
		assign = fmt.Sprintf("%s = data[start:end]", dest)
	}
	return fmt.Sprintf(`start := i
end, err := unstable.SkipValue(data, i)
if err != nil {
	return end, err
}
%s
i = end`, assign)
}

// numberRead emits the read for a json.Number field: capture the number token's
// raw literal as a string and convert it to json.Number. nocopy aliases the input
// (json.Number's underlying type is string, so the conversion never copies).
//
// This is the one place a qualified name is emitted without going through
// typeStr — the conversion is spelled by hand — so it is also where the
// encoding/json import is claimed for it. It uses the schema's own qualifier, so
// an aliased import stays consistent between the conversion and the import block.
func (g *gen) numberRead(dest string, nocopy bool) string {
	reader := "unstable.ReadNumberOrNull"
	if nocopy {
		reader = "unstable.ReadNumberNoCopyOrNull"
	}
	g.needJSON = true
	return fmt.Sprintf(`s, end, err := %s(data, i)
if err != nil {
	return end, err
}
%s = %s.Number(s)
i = end`, reader, dest, g.jsonQual)
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
		if i >= len(data) {
			// All-whitespace body: like an empty string, leave the zero value.
			// inner assumes a value byte at data[i] (a pointer field's null
			// probe reads it unguarded), so it must not run here.
			return i, nil
		}
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
	g.markDepthFn(fn, g.exprThreadsDepth(t.Elt))
	// A bare float64/int/uint element routes to the batched pkg/unstable reader
	// even when the array is reached through a decoder function (a lax field's
	// scratch decode) rather than field's inline emission — one call per array
	// instead of a non-inlinable reader call per element, matching what field
	// emits for the same [N]T. The wrapper only adapts the signature.
	if bfn := batchArrayFn(t.Elt); bfn != "" {
		body := fmt.Sprintf(`func %[1]s(out *%[2]s, data []byte, i int%[4]s) (int, error) {
	return %[3]s((*out)[:], data, i)
}`, fn, arrType, bfn, g.depthParam(fn)+g.arenaParam())
		g.decoders = append(g.decoders, body)
		return fn
	}
	prevDepth := g.enterBody(fn, false)
	elem := g.field("(*out)[idx]", t.Elt, hint, nocopy, lax)
	g.depthArg = prevDepth
	// Trailing commas are rejected by the first-iteration flag, as in
	// genStructBody.
	body := fmt.Sprintf(`func %[1]s(out *%[2]s, data []byte, i int%[5]s) (int, error) {
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
}`, fn, arrType, g.skipWS("i", "i"), elem, g.depthParam(fn)+g.arenaParam())
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
	case id.Name == "byte" || id.Name == "uint8":
		// []byte follows encoding/json: a base64 string or a numeric array.
		// Fixed-size [N]byte stays numeric-only (batchArrayFn), also like the
		// stdlib.
		return "unstable.DecodeByteSlice"
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

// sliceDecoder returns (generating on first use) the decoder for a slice of elt.
//
// The emitted body resets the slice's length before the loop, because decoding an
// array into a slice replaces its contents rather than appending to them (the rule
// encoding/json documents). The reset is *guarded* on a non-empty length rather
// than done unconditionally: a fresh decode — the common case, where the slice is
// already empty — then pays a load and a not-taken branch instead of a three-word
// slice-header store. Storing unconditionally measured cloudflare +1.26%
// (p=0.001, n=8) for its three slice fields; with the guard every slice-heavy case
// (cloudflare, marine_ik, citm_catalog, mesh, canada, large-json) is flat.
//
// Note the reset text lives inside a fmt.Sprintf template, so any '%' written into
// it must be doubled — an unescaped one is silently emitted into every generated
// file as "%!(MISSING)", which is why the measurement note above is here and not
// in the generated comment.
func (g *gen) sliceDecoder(elt ast.Expr, hint string, nocopy, lax, root bool) string {
	if fn := batchSliceFn(elt); fn != "" {
		if g.arena {
			// The ...Arena twin carves the presized backing from the decode's
			// arena; callDecoder appends the `a` argument for every call under
			// an arena root, which matches this reader's extra parameter.
			return fn + "Arena"
		}
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
	if root {
		// A named slice root's decoder grows by progress extrapolation (see the
		// presize == "" branch below), which is only sound where the array is
		// known to span the rest of the document — the root position. Keep it
		// memo-distinct from any field decoder for the same element type, which
		// must keep the flat-2x growth.
		key = "root:" + key
	}
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
	g.markDepthFn(fn, g.exprThreadsDepth(elt))
	eltStr := g.typeStr(elt)
	prevDepth := g.enterBody(fn, false)
	inner := g.field("(*out)[len(*out)-1]", elt, singular(hint)+"Entry", nocopy, lax)
	g.depthArg = prevDepth
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
		growCall := "unstable.GrowSlice(*out)"
		if root {
			// A slice ROOT's growth extrapolates the final element count from
			// decode progress instead of blind doubling: at the grow point the
			// loop knows the array's '[' index (captured in lightningArrStart
			// below — the only extra code, no extra scanning), the current
			// position and the document end, so GrowSliceEst sizes the new
			// backing by len * (end-start) / (i-start), padded and clamped (see
			// its doc). The premise — the array spans the rest of the document,
			// making progress a faithful density sample — is structurally true
			// exactly at the root, so a root array of large records reaches its
			// final capacity in one or two grows instead of log2(n) doublings
			// of dead backing + memmove (github_events-shaped documents).
			// Nested slices keep the measured flat-2x GrowSlice: for them the
			// document tail says nothing about the array's own length, the
			// estimate always saturates its clamp, and 4x-and-up growth was
			// measured and rejected for exactly those shapes (see GrowSlice).
			presize = "	lightningArrStart := i\n"
			growCall = "unstable.GrowSliceEst(*out, lightningArrStart, i, len(data))"
		}
		// The capacity hint below is the generated code's only use of unsafe, so
		// this is where that import is claimed (the same emit-site rule as
		// typeStr's noteQualifiers — a JSON key literal reading "unsafe.Sizeof"
		// used to claim it through the old text scan).
		g.needUnsafe = true
		grow = fmt.Sprintf(`var zero %[1]s
		if *out == nil {
			*out = make([]%[1]s, 1, max(4, 256/max(1, int(unsafe.Sizeof(zero)))))
		} else {
			if len(*out) == cap(*out) {
				*out = %[2]s
			}
			*out = append(*out, zero)
		}`, eltStr, growCall)
	}
	// Trailing commas ([1,]) are rejected by the first-iteration flag, as in
	// genStructBody (see there — a rotated loop regressed cloudflare +11%).
	body := fmt.Sprintf(`func %[1]s(out *[]%[2]s, data []byte, i int%[7]s) (int, error) {
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
	// Decoding an array into a slice replaces its contents rather than appending
	// to them, matching encoding/json ("Unmarshal resets the slice length to zero
	// and then appends each element"). The backing array is kept, so reuse stays
	// allocation-free; a nil slice stays nil through [:0], leaving the presize and
	// first-append capacity hint below to fire only for the fresh case.
	if len(*out) != 0 {
		*out = (*out)[:0]
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
}`, fn, eltStr, inner, presize, g.skipWS("i", "i"), grow, g.depthParam(fn)+g.arenaParam())
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
		case g.isNumber(t):
			// A json.Number element is a bare number token — a scalar — so the
			// cheaper comma-counting scan sizes the slice.
			counter = "unstable.CountArrayScalars"
		case g.isTime(t):
			// An RFC 3339 timestamp or Unix-timestamp number never contains a
			// comma or bracket, so the cheap comma-counting scan sizes a
			// []time.Time too — far cheaper than CountArrayElements, which would
			// skipString past every element. (time-array −15%.)
			counter = "unstable.CountArrayScalars"
		case g.isRaw(t):
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
	g.markDepthFn(fn, g.exprThreadsDepth(valExpr))
	valStr := g.typeStr(valExpr)
	prevDepth := g.enterBody(fn, false)
	inner := g.field("val", valExpr, hint+"Value", nocopy, lax)
	g.depthArg = prevDepth
	// With nocopy the key aliases the input (ReadKey already returns an alias for
	// an unescaped key); otherwise it is copied so the map owns it.
	keyAssign := "m[string([]byte(key))] = val"
	if nocopy {
		keyAssign = "m[key] = val"
	}
	// Trailing commas are rejected by the first-iteration flag, as in
	// genStructBody.
	body := fmt.Sprintf(`func %[1]s(out *map[string]%[2]s, data []byte, i int%[9]s) (int, error) {
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
}`, fn, valStr, inner, g.skipWS("i", "i"), g.skipWS("i", "ni"), g.skipWS("i", "i+1"), g.readKey(), keyAssign, g.depthParam(fn)+g.arenaParam())
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

	// "encoding/json", "time" and "unsafe" are imported only when the generated
	// code actually names them: a json.Number conversion or a type text that
	// spells json.RawMessage / time.Time (in an anonymous struct, a slice or map
	// element, a pointee), and the unsafe.Sizeof in a slice's capacity hint.
	//
	// The flags are set where those names are EMITTED (typeStr's noteQualifiers,
	// numberRead, sliceDecoder). This used to be decided by substring-scanning
	// the finished text, on the premise that the tokens "appear only as real type
	// usages" — but JSON key literals are emitted into the dispatch switch as
	// string constants, so a field tagged `json:"time.Time"` produced `case
	// "time.Time":` and an import of "time" that nothing used: a schema that
	// generated cleanly and did not compile. Setting the flag at the emit site
	// rather than per field also answers what the text scan was introduced to
	// fix — it sees a type nested inside an anonymous struct (typeStr renders the
	// whole struct text) and it does not fire for a nocopy raw field, whose
	// decode emits no qualified name at all.
	//
	// The qualifier is the schema's own (see collectQualifiers), so a file that
	// imports encoding/json under an alias generates code spelling that alias and
	// importing it under the same name.
	imports := []string{strconv.Quote(unstablePkg)}
	if g.needJSON {
		imports = append(imports, fmt.Sprintf("%s %q", g.jsonQual, "encoding/json"))
	}
	if g.needTime {
		if g.timeQual == "time" {
			imports = append(imports, `"time"`)
		} else {
			imports = append(imports, fmt.Sprintf("%s %q", g.timeQual, "time"))
		}
	}
	if g.needUnsafe {
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

// tagInfo is one field's parsed `json:"..."` tag.
type tagInfo struct {
	names   []string // the JSON key(s) the field answers to
	skip    bool     // json:"-": never read from the input
	nocopy  bool
	lax     bool
	unwrap  bool
	unknown []string // options the generator does not act on, for diagnostics
}

// jsonTag parses a field's `json:"..."` struct tag. The name may list several
// pipe-separated names (`json:"Status|EdgeResponseStatus"`), each of which maps
// the same field; any of them in the input fills it. Comma-separated options
// follow the name, as in encoding/json. The generator acts on three:
//
//   - "nocopy": string and raw fields alias the input bytes instead of copying
//     them, so the caller must keep the input alive and unchanged;
//   - "lax": a type mismatch on this field's value is a no-op (the value is
//     skipped and the field left unset) rather than an error;
//   - "unwrap": the value is a string whose contents are themselves the JSON to
//     decode into the field.
//
// "omitempty" is accepted in silence: it is encode-only, so it means nothing to
// a decoder, and it is on real structs everywhere. Any other option is returned
// in unknown for the caller to warn about — a typo'd ",nocpy" turning off the
// aliasing the author asked for is a silent performance regression, and
// encoding/json's ",string" is a silent *semantic* one.
func jsonTag(tag *ast.BasicLit) tagInfo {
	if tag == nil {
		return tagInfo{}
	}
	s, err := strconv.Unquote(tag.Value)
	if err != nil {
		return tagInfo{}
	}
	v := reflect.StructTag(s).Get("json")
	if v == "" {
		return tagInfo{}
	}
	parts := strings.Split(v, ",")
	if parts[0] == "-" && len(parts) == 1 {
		return tagInfo{skip: true}
	}
	var t tagInfo
	for _, o := range parts[1:] {
		switch o {
		case "nocopy":
			t.nocopy = true
		case "lax":
			t.lax = true
		case "unwrap":
			t.unwrap = true
		case "", "omitempty":
			// A trailing comma, and the stdlib's encode-only option.
		default:
			t.unknown = append(t.unknown, o)
		}
	}
	for _, n := range strings.Split(parts[0], "|") {
		if n != "" {
			t.names = append(t.names, n)
		}
	}
	return t
}

// warnTagOptions reports the json-tag options the generator does not act on.
//
// This warns where an unknown //lightning: directive is a hard error, and the
// asymmetry is deliberate: a //lightning: directive can only have come from this
// generator, so an unrecognized one is unambiguously a mistake, while a json tag
// is encoding/json's vocabulary too. A schema migrating from the stdlib carries
// options that are simply meaningless here, and failing the build on them would
// break working code for no safety gain. A warning still catches what silence
// cost — the typo that quietly disables aliasing, and ",string", whose absence
// changes what the schema decodes.
func (g *gen) warnTagOptions(opts []string, field string) {
	for _, o := range opts {
		if o == "string" {
			// encoding/json's ",string" decodes the value from a JSON string
			// wrapping the number/bool. Ignoring it silently would decode a
			// migrated schema differently from the stdlib it came from.
			g.warnf("json tag option %q on %s is not implemented: the value is decoded with the field's declared Go type (for a whole JSON document embedded in a string, see the unwrap option)", o, field)
			continue
		}
		g.warnf("unrecognized json tag option %q on %s; it is ignored (this generator understands nocopy, lax and unwrap)", o, field)
	}
}

// fieldLabel names a struct field for a diagnostic: its declared name(s), or,
// for an embedded field, the type it embeds.
func fieldLabel(f *ast.Field) string {
	if len(f.Names) == 0 {
		if n := embeddedName(f.Type); n != "" {
			return "embedded field " + n
		}
		return "embedded field"
	}
	names := make([]string, len(f.Names))
	for i, n := range f.Names {
		names[i] = n.Name
	}
	return "field " + strings.Join(names, ", ")
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

// namedRefs returns the top-level named types referenced directly by name's own
// definition — through fields, pointers, slices, arrays, maps and anonymous
// structs. Unlike markReferenced it *keeps* a self-reference (a self-recursive type
// is exactly what the cycle search is looking for) and it stops at named types
// rather than recursing through them, so the result is the type graph's edge set
// out of name.
func (g *gen) namedRefs(name string) map[string]bool {
	out := map[string]bool{}
	var walk func(ast.Expr)
	walk = func(expr ast.Expr) {
		switch t := unparen(expr).(type) {
		case *ast.Ident:
			if g.structTypes[t.Name] != nil || g.sliceTypes[t.Name] != nil || g.mapTypes[t.Name] != nil {
				out[t.Name] = true
			}
		case *ast.StarExpr:
			walk(t.X)
		case *ast.ArrayType:
			walk(t.Elt)
		case *ast.MapType:
			walk(t.Value)
		case *ast.StructType:
			for _, f := range t.Fields.List {
				walk(f.Type)
			}
		}
	}
	switch {
	case g.structTypes[name] != nil:
		for _, f := range g.structTypes[name].Fields.List {
			walk(f.Type)
		}
	case g.sliceTypes[name] != nil:
		walk(g.sliceTypes[name].Elt)
	case g.mapTypes[name] != nil:
		walk(g.mapTypes[name].Value)
	}
	return out
}

// computeDepthThreading fills g.threadDepth with the named types whose decoders
// must carry a recursion depth: those lying on a reference cycle, plus those that
// can reach one (they sit above it in the call chain and have to pass the counter
// down).
//
// A cycle is found by asking, for each type, whether it can reach itself. The
// graphs here are a single file's type declarations — a handful of nodes — so the
// naive per-node DFS closure is far cheaper than its own setup, and it needs no
// SCC bookkeeping.
//
// Only reaching a cycle earns the parameter, so a schema without one generates
// exactly the code it did before.
func (g *gen) computeDepthThreading() {
	g.threadDepth = map[string]bool{}

	refs := make(map[string]map[string]bool, len(g.order))
	for _, n := range g.order {
		refs[n] = g.namedRefs(n)
	}
	reach := make(map[string]map[string]bool, len(g.order))
	for _, n := range g.order {
		seen := map[string]bool{}
		var dfs func(string)
		dfs = func(cur string) {
			for u := range refs[cur] {
				if !seen[u] {
					seen[u] = true
					dfs(u)
				}
			}
		}
		dfs(n)
		reach[n] = seen
	}
	// A type on a cycle is one that reaches itself.
	cyclic := map[string]bool{}
	for _, n := range g.order {
		if reach[n][n] {
			cyclic[n] = true
		}
	}
	for _, n := range g.order {
		if cyclic[n] {
			g.threadDepth[n] = true
			continue
		}
		for u := range reach[n] {
			if cyclic[u] {
				g.threadDepth[n] = true
				break
			}
		}
	}
}

// exprThreadsDepth reports whether a decoder generated for the type expression
// expr must carry a depth parameter — that is, whether expr mentions a named type
// that can reach a cycle. It is the composite-helper counterpart of a
// g.threadDepth lookup on a named type.
func (g *gen) exprThreadsDepth(expr ast.Expr) bool {
	switch t := unparen(expr).(type) {
	case *ast.Ident:
		return g.threadDepth[t.Name]
	case *ast.StarExpr:
		return g.exprThreadsDepth(t.X)
	case *ast.ArrayType:
		return g.exprThreadsDepth(t.Elt)
	case *ast.MapType:
		return g.exprThreadsDepth(t.Value)
	case *ast.StructType:
		for _, f := range t.Fields.List {
			if g.exprThreadsDepth(f.Type) {
				return true
			}
		}
	}
	return false
}

// markDepthFn records that fn takes a depth parameter, and must be called before
// fn's body is generated so that a recursive call reached while generating it sees
// the right signature — the same reason the memo entry is set early.
func (g *gen) markDepthFn(fn string, threads bool) {
	if threads {
		g.depthFns[fn] = true
	}
}

// depthParam is the extra parameter fn's signature carries, and depthArg what a
// call site inside the body currently being generated passes for it. Both are ""
// for a decoder that does not thread depth, which is why a cycle-free schema
// produces byte-identical output.
func (g *gen) depthParam(fn string) string {
	if g.depthFns[fn] {
		return ", depth int"
	}
	return ""
}

func (g *gen) depthArgFor(fn string) string {
	if !g.depthFns[fn] {
		return ""
	}
	arg := g.depthArg
	if arg == "" {
		arg = "0"
	}
	return ", " + arg
}

// enterBody sets what call sites in fn's body pass for a callee's depth parameter,
// returning the previous value for the caller to restore (bodies are generated
// recursively, so this is a stack discipline).
//
// A struct decoder passes depth+1: its frame is the one that repeats as the
// document nests, so it is where a level is counted. A composite helper (slice,
// array, map, lax value wrapper) threads depth unchanged — it sits between two
// struct frames rather than adding a level of its own.
func (g *gen) enterBody(fn string, isStruct bool) string {
	prev := g.depthArg
	switch {
	case !g.depthFns[fn]:
		g.depthArg = "0"
	case isStruct:
		g.depthArg = "depth+1"
	default:
		g.depthArg = "depth"
	}
	return prev
}

// depthGuard is the bound itself, emitted at the top of a depth-threading struct
// decoder. Every cycle in a decodable schema runs through a named struct — a named
// slice or map type is only decodable at the root, not as a field type — so
// guarding the struct decoders bounds every cycle.
func (g *gen) depthGuard(fn string) string {
	if !g.depthFns[fn] {
		return ""
	}
	return `if depth >= unstable.MaxDepth {
		return i, unstable.ErrMaxDepth
	}
	`
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

// entryTypes decides which collected types get an UnmarshalJSON method: a type
// is skipped — its decoder emitted by the type that reaches it — only when it is
// reachable from a type that actually gets one. refs is the reference graph
// (name -> the top-level types its own definition names, self excluded, as
// markReferenced builds it) and referenced its union.
//
// The obvious rule, "skip everything some other type references", loses a cycle:
// each member of a mutually recursive pair references the other, so both are
// referenced, both are skipped, and the pair gets no method AND no decoder — the
// user's code then fails to compile on the missing UnmarshalJSON, with the
// generator having exited 0 in silence. That was papered over by a special case
// for the file where *every* type is referenced (emit all of them), which only
// fired when the cycle was the whole file: declaring one unrelated type
// alongside the pair brought the hole back.
//
// The fixpoint generalizes that special case instead of guarding it. Start from
// the types nothing references, mark everything reachable from them, and note
// that a type which is neither emitted nor reachable can only be sitting on a
// cycle no emitted type enters; promote one such cycle and recompute, until
// every type is emitted or covered.
//
// Which cycle, and how much of it, is the whole design — both answers exist to
// make the result independent of DECLARATION ORDER, since "reorder two type
// declarations and a different type becomes decodable" is not a property anyone
// can reason about.
//
//   - Promote the candidate's whole strongly-connected component, not the
//     candidate alone. Promoting one member would leave the rest merely
//     reachable from it, so `MutA ↔ MutB` would give a method to whichever came
//     first in the file — and would downgrade the fully-cyclic file, which the
//     old special case gave a method per member, so a cycle-only schema would
//     silently lose an UnmarshalJSON. Every member of a cycle is mutually
//     reachable with every other, so the SCC is exactly "the cycle".
//
//   - Take the candidate from a SOURCE component of what is still uncovered:
//     one that no other uncovered type outside it reaches. Scanning for the
//     first uncovered type in source order is not enough, and the gap only
//     opens when the file has NO entry type at all (every type referenced by
//     something, so the first round starts with nothing covered). Then a record
//     that merely hangs off a cycle — reachable FROM it, never reaching back —
//     is uncovered like everything else, and if it happens to be declared first
//     it is picked and handed a method, where declaring it after the cycle
//     leaves it covered and correctly nested. Requiring a source makes the
//     first round behave exactly as a file that does have an entry type: the
//     hanger-on is never a candidate, so it stays nested and its
//     `type recordStd Record` reflection baseline keeps reaching encoding/json.
//
// Among source components the first in source order is taken, and that choice
// is immaterial: distinct source components are mutually unreachable, so none
// of them can cover another and every one is promoted before the fixpoint ends.
//
// A source always exists while anything is uncovered. Only uncovered-internal
// edges can matter — if an emitted-or-covered type referenced an uncovered one,
// that one would be covered by definition — so the condensation of the induced
// subgraph is a finite non-empty DAG, and a finite non-empty DAG has a source.
func (g *gen) entryTypes(refs map[string]map[string]bool, referenced map[string]bool) map[string]bool {
	emitted := make(map[string]bool, len(g.order))
	for _, n := range g.order {
		if !referenced[n] {
			emitted[n] = true
		}
	}
	for {
		reach := make(map[string]map[string]bool, len(g.order))
		for _, n := range g.order {
			reach[n] = reachableFrom(refs, n)
		}
		covered := map[string]bool{}
		for _, n := range g.order {
			if emitted[n] {
				for u := range reach[n] {
					covered[u] = true
				}
			}
		}
		// The first uncovered type sitting in a source component. reach[u][n]
		// with !reach[n][u] is "u is strictly upstream of n": u reaches n and n
		// cannot reach back, so n's component is not a source. Mutual
		// reachability is the same component and does not disqualify it.
		promote, first := "", ""
		for _, n := range g.order {
			if emitted[n] || covered[n] {
				continue
			}
			if first == "" {
				first = n
			}
			source := true
			for _, u := range g.order {
				if u == n || emitted[u] || covered[u] {
					continue
				}
				if reach[u][n] && !reach[n][u] {
					source = false
					break
				}
			}
			if source {
				promote = n
				break
			}
		}
		if promote == "" {
			// Unreachable per the DAG argument above. Falling back to the first
			// uncovered type keeps the walk terminating and every type reachable
			// from something emitted, which is the property callers depend on —
			// a silent "no method and no decoder" is the bug this function
			// exists to prevent, so never let a missing source stop the loop.
			promote = first
		}
		if promote == "" {
			return emitted
		}
		// promote and everything mutually reachable with it: the cycle it sits
		// on, and any other cycle that shares a member with it.
		emitted[promote] = true
		for _, n := range g.order {
			if reach[promote][n] && reach[n][promote] {
				emitted[n] = true
			}
		}
	}
}

// reachableFrom returns every type reachable from start by following refs. start
// itself is in the result exactly when it lies on a cycle, which is what makes
// the mutual-reachability test in entryTypes a strongly-connected-component test.
func reachableFrom(refs map[string]map[string]bool, start string) map[string]bool {
	seen := map[string]bool{}
	var walk func(string)
	walk = func(cur string) {
		for u := range refs[cur] {
			if !seen[u] {
				seen[u] = true
				walk(u)
			}
		}
	}
	walk(start)
	return seen
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

// qualified reports whether expr is the selector pkg.name, where pkg is the
// identifier the *input file* imports that package under (see
// collectQualifiers) — never a hardcoded "json"/"time", so an aliased import
// resolves and, more importantly, a same-named type from somewhere else does
// not. An empty qual means the package is not imported at all, so nothing
// matches it: a bare `mypkg.RawMessage` is then reported as the unsupported
// type it is rather than decoded as encoding/json's and emitted with no import.
func qualified(expr ast.Expr, qual, name string) bool {
	if qual == "" {
		return false
	}
	sel, ok := unparen(expr).(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != name {
		return false
	}
	pkg, ok := sel.X.(*ast.Ident)
	return ok && pkg.Name == qual
}

// isRaw is the single spelling of the "this is encoding/json's raw-JSON type"
// predicate, used by both the field emitter and slicePresize. It was duplicated
// in field() once, and the copies are exactly the kind that drift apart.
// json.RawValue is the json/v2 spelling of the same thing.
func (g *gen) isRaw(expr ast.Expr) bool {
	return qualified(expr, g.jsonQual, "RawMessage") || qualified(expr, g.jsonQual, "RawValue")
}

func (g *gen) isNumber(expr ast.Expr) bool { return qualified(expr, g.jsonQual, "Number") }

func (g *gen) isTime(expr ast.Expr) bool { return qualified(expr, g.timeQual, "Time") }

// isAnyInterface reports whether an interface type is `any` under another
// spelling — the only interface a decoded value can be assigned to. That is the
// empty interface itself, and one whose elements are all embedded empty
// interfaces (`interface{ any }`, `interface{ interface{} }`), which are the
// same type. A method (a field with a name), an embedded named interface, or a
// type-set element makes it something else, and the field is unsupported.
func isAnyInterface(t *ast.InterfaceType) bool {
	// Methods is the whole element list, embedded elements and type-set terms
	// included, not just the methods its name suggests.
	if t.Methods == nil {
		return true
	}
	for _, f := range t.Methods.List {
		if len(f.Names) != 0 { // a method
			return false
		}
		switch e := unparen(f.Type).(type) {
		case *ast.Ident:
			if e.Name != "any" {
				return false
			}
		case *ast.InterfaceType:
			if !isAnyInterface(e) {
				return false
			}
		default: // a union, ~T, or a qualified interface name
			return false
		}
	}
	return true
}

func isScalar(name string) bool {
	switch name {
	case "string", "bool", "float32", "float64":
		return true
	}
	return intKinds[name] || uintKinds[name]
}

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
