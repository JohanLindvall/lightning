// Command lightning generates a fast, allocation-light json.Unmarshaler
// implementation for the struct types declared in a Go source file.
//
// Usage:
//
//	go run . bench/1/data.go
//
// For each input file FOO.go it writes FOO_unmarshal.go next to it, containing
// an UnmarshalJSON method for every top-level struct type. The generated code
// calls the shared scanner helpers in pkg/support.
//
// Supported field types: string, bool, all sized int/uint kinds, float32,
// float64, json.RawMessage, nested (named or anonymous) structs, slices, maps
// with string keys, pointers, and interface{}/any (decoded into the usual
// Go representation of an arbitrary JSON value).
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
		fset:         fset,
		pkg:          file.Name.Name,
		structTypes:  map[string]*ast.StructType{},
		order:        nil,
		used:         map[string]bool{},
		memo:         map[string]string{},
		compactTypes: map[string]bool{},
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

	// Collect every top-level struct type, in source order. A type carrying the
	// //lightning:compact directive (on the type or its declaration) gets a
	// decoder that omits the inter-token SkipWS calls; see (*gen).skipWS.
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
			if st, ok := ts.Type.(*ast.StructType); ok {
				g.structTypes[ts.Name.Name] = st
				g.order = append(g.order, ts.Name.Name)
				if hasCompactDirective(gd.Doc, ts.Doc) {
					g.compactTypes[ts.Name.Name] = true
				}
			}
		}
	}
	if len(g.order) == 0 {
		return fmt.Errorf("no top-level struct types found")
	}

	var methods []string
	for _, name := range g.order {
		g.compact = g.compactTypes[name]
		g.prefix = "lightning" + g.pathFrag + name
		methods = append(methods, g.genUnmarshal(name))
	}
	if len(g.errs) > 0 {
		return errors.Join(g.errs...)
	}

	src := g.assemble(inPath, methods)
	formatted, ferr := format.Source([]byte(src))
	out := formatted
	if ferr != nil {
		// Emit the unformatted source so the failure can be inspected.
		out = []byte(src)
		fmt.Fprintf(os.Stderr, "%s: gofmt failed (writing unformatted): %v\n", inPath, ferr)
	}

	outPath := strings.TrimSuffix(inPath, ".go") + "_unmarshal.go"
	if err := os.WriteFile(outPath, out, 0o644); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "wrote %s\n", outPath)
	return nil
}

// gen holds the state accumulated while walking the type graph.
type gen struct {
	fset        *token.FileSet
	pkg         string
	structTypes map[string]*ast.StructType
	order       []string

	used map[string]bool   // reserved decoder function names
	memo map[string]string // type-key -> decoder function name (dedupe)

	compactTypes map[string]bool // top-level types tagged //lightning:compact
	compact      bool            // whether the type currently being generated is compact
	pathFrag     string          // import path sanitized into an identifier fragment
	prefix       string          // current per-type prefix for generated function names

	decoders []string // generated decoder functions, in creation order
	errs     []error  // generation errors; reported together after the walk

	needBool, needUint, needFloat, needAny, needJSON, needNoCopyStr, needTime bool
}

func (g *gen) typeStr(e ast.Expr) string {
	var b strings.Builder
	printer.Fprint(&b, g.fset, e)
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

// hasCompactDirective reports whether any of the given doc-comment groups
// carries the //lightning:compact directive.
func hasCompactDirective(groups ...*ast.CommentGroup) bool {
	for _, cg := range groups {
		if cg == nil {
			continue
		}
		for _, c := range cg.List {
			if strings.TrimSpace(strings.TrimPrefix(c.Text, "//")) == "lightning:compact" {
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

// skipWSExpr returns the expression giving the first non-whitespace index at or
// after src. For a //lightning:compact type the input is asserted to hold no
// whitespace between tokens (the form compact JSON serializers emit), so it
// collapses to src itself; otherwise it is the support.SkipWS call the generator
// has always emitted. Only the per-key/element calls inside objects, arrays and
// maps are elided this way; UnmarshalJSON keeps its leading/trailing SkipWS so a
// document framed with surrounding whitespace (e.g. a trailing newline) still
// decodes.
func (g *gen) skipWSExpr(src string) string {
	if g.compact {
		return src
	}
	return "support.SkipWS(data, " + src + ")"
}

// skipWS returns the statement assigning dst the first non-whitespace index at
// or after src. In compact mode it becomes a plain dst = src, or nothing when
// dst and src are identical (the assignment would be a no-op).
func (g *gen) skipWS(dst, src string) string {
	if g.compact && dst == src {
		return ""
	}
	return dst + " = " + g.skipWSExpr(src)
}

// cmark and csuf distinguish a compact decoder from its non-compact counterpart
// in, respectively, memo keys and generated function names, so the same type
// reached from both a compact and a non-compact root yields two decoders.
func (g *gen) cmark() string {
	if g.compact {
		return "compact:"
	}
	return ""
}

func (g *gen) csuf() string {
	if g.compact {
		return "Compact"
	}
	return ""
}

// genUnmarshal emits the UnmarshalJSON method for a named struct type and makes
// sure its decoder (and everything it reaches) is generated.
func (g *gen) genUnmarshal(name string) string {
	fn := g.namedStruct(name)
	return fmt.Sprintf(`// UnmarshalJSON parses JSON into the %[1]s. Fields whose json tag carries the
// "nocopy" option alias the input data instead of copying it; if any are
// present, the caller must keep data unchanged while the result is in use.
func (v *%[1]s) UnmarshalJSON(data []byte) error {
	i := support.SkipWS(data, 0)
	if i >= len(data) {
		return support.ErrTruncated
	}
	if data[i] == 'n' {
		end, err := support.ExpectNull(data, i)
		if err != nil {
			return err
		}
		if support.SkipWS(data, end) != len(data) {
			return support.ErrInvalidJSON
		}
		return nil
	}
	end, err := %[2]s(v, data, i)
	if err != nil {
		return err
	}
	if support.SkipWS(data, end) != len(data) {
		return support.ErrInvalidJSON
	}
	return nil
}`, name, fn)
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

func (g *gen) genStructBody(fn, paramType string, st *ast.StructType) {
	var cases strings.Builder
	seen := map[string]bool{} // json keys already claimed within this struct
	for _, f := range st.Fields.List {
		if len(f.Names) == 0 {
			fmt.Fprintf(os.Stderr, "warning: skipping embedded field %s\n", g.typeStr(f.Type))
			continue
		}
		tagNames, skip, nocopy, lax := jsonTag(f.Tag)
		if skip {
			continue
		}
		for _, nm := range f.Names {
			field := nm.Name
			if !ast.IsExported(field) {
				continue
			}
			keys := tagNames
			if len(keys) == 0 {
				keys = []string{field}
			}
			quoted := make([]string, len(keys))
			for i, k := range keys {
				if seen[k] {
					g.errs = append(g.errs, fmt.Errorf("%s: json name %q is mapped more than once", strings.TrimPrefix(paramType, "*"), k))
				}
				seen[k] = true
				quoted[i] = strconv.Quote(k)
			}
			var code string
			if lax {
				code = g.laxField("v."+field, f.Type, field, nocopy)
			} else {
				code = g.field("v."+field, f.Type, field, nocopy, false)
			}
			fmt.Fprintf(&cases, "\tcase %s:\n%s\n", strings.Join(quoted, ", "), code)
		}
	}

	body := fmt.Sprintf(`func %[1]s(v %[2]s, data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, support.ErrTruncated
	}
	if data[i] == 'n' {
		return support.ExpectNull(data, i)
	}
	if data[i] != '{' {
		return i, support.ErrExpectObject
	}
	i++
	for {
		%[4]s
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		if data[i] == '}' {
			return i + 1, nil
		}
		key, ni, err := support.ReadKey(data, i)
		if err != nil {
			return ni, err
		}
		%[5]s
		if i >= len(data) || data[i] != ':' {
			return i, support.ErrExpectColon
		}
		%[6]s
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		switch key {
%[3]s
		default:
			end, err := support.SkipValue(data, i)
			if err != nil {
				return end, err
			}
			i = end
		}
		%[7]s
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		if data[i] == '}' {
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, support.ErrInvalidJSON
		}
		i++
	}
}`, fn, paramType, cases.String(), g.skipWS("i", "i"), g.skipWS("i", "ni"), g.skipWS("i", "i+1"), g.skipWS("i", "i"))
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
	end, err := support.ExpectNull(data, i)
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
		fmt.Fprintf(os.Stderr, "warning: unknown type %q for %s; value skipped\n", t.Name, dest)
		return g.skipEmit()

	case *ast.SelectorExpr:
		if t.Sel.Name == "RawMessage" || t.Sel.Name == "RawValue" {
			return g.rawMessage(dest, nocopy)
		}
		if isTime(t) {
			return g.timeRead(dest, lax)
		}
		fmt.Fprintf(os.Stderr, "warning: unsupported type %s for %s; value skipped\n", g.typeStr(t), dest)
		return g.skipEmit()

	case *ast.StructType:
		return g.callDecoder(dest, g.anonStruct(t, hint))

	case *ast.ArrayType:
		if t.Len != nil {
			fmt.Fprintf(os.Stderr, "warning: fixed-size array for %s; value skipped\n", dest)
			return g.skipEmit()
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
		fmt.Fprintf(os.Stderr, "warning: unsupported type %s for %s; value skipped\n", g.typeStr(expr), dest)
		return g.skipEmit()
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
	end, err = support.SkipValue(data, i)
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
		reader := "support.ReadStringOrNull"
		if nocopy {
			reader = "support.ReadStringNoCopyOrNull"
			g.needNoCopyStr = true
		}
		return fmt.Sprintf(`s, end, err := %s(data, i)
if err != nil {
	return end, err
}
%s = s
i = end`, reader, dest)

	case name == "bool":
		g.needBool = true
		return fmt.Sprintf(`b, end, err := support.ReadBoolOrNull(data, i)
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
		return fmt.Sprintf(`f, end, err := support.ReadFloat64OrNull(data, i)
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
		return fmt.Sprintf(`n, end, err := support.ReadInt64OrNull(data, i)
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
		return fmt.Sprintf(`n, end, err := support.ReadUint64OrNull(data, i)
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
end, err := support.SkipValue(data, i)
if err != nil {
	return end, err
}
if data[start] != 'n' {
	%s
}
i = end`, assign)
}

func (g *gen) timeRead(dest string, lax bool) string {
	reader := "support.ReadTimeOrNull"
	if lax {
		// lax accepts RFC 3339 (with 'T' or space) and numeric Unix timestamps.
		reader = "support.ReadTimeLaxOrNull"
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
	return fmt.Sprintf(`val, end, err := support.DecodeValue(data, i)
if err != nil {
	return end, err
}
%s = val
i = end`, dest)
}

func (g *gen) skipEmit() string {
	return `end, err := support.SkipValue(data, i)
if err != nil {
	return end, err
}
i = end`
}

func (g *gen) sliceDecoder(elt ast.Expr, hint string, nocopy, lax bool) string {
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
	inner := g.field("el", elt, singular(hint)+"Entry", nocopy, lax)
	presize := g.slicePresize(elt, eltStr)
	body := fmt.Sprintf(`func %[1]s(out *[]%[2]s, data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, support.ErrTruncated
	}
	if data[i] == 'n' {
		end, err := support.ExpectNull(data, i)
		if err != nil {
			return end, err
		}
		*out = nil
		return end, nil
	}
	if data[i] != '[' {
		return i, support.ErrExpectArray
	}
%[4]s	i++
	for {
		%[5]s
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		var el %[2]s
		%[3]s
		*out = append(*out, el)
		%[5]s
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, support.ErrInvalidJSON
		}
		i++
	}
}`, fn, eltStr, inner, presize, g.skipWS("i", "i"))
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
			counter = "support.CountArrayScalars"
		case t.Name == "string":
			counter = "support.CountArrayElements"
		}
	case *ast.StructType, *ast.ArrayType, *ast.MapType:
		counter = "support.CountArrayElements"
	case *ast.SelectorExpr:
		if isTime(t) || isRaw(t) {
			counter = "support.CountArrayElements"
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
		fmt.Fprintf(os.Stderr, "warning: map key type %s unsupported for %s; value skipped\n", g.typeStr(keyExpr), hint)
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
	body := fmt.Sprintf(`func %[1]s(out *map[string]%[2]s, data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, support.ErrTruncated
	}
	if data[i] == 'n' {
		end, err := support.ExpectNull(data, i)
		if err != nil {
			return end, err
		}
		*out = nil
		return end, nil
	}
	if data[i] != '{' {
		return i, support.ErrExpectObject
	}
	i++
	m := *out
	if m == nil {
		m = make(map[string]%[2]s)
	}
	for {
		%[4]s
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		if data[i] == '}' {
			*out = m
			return i + 1, nil
		}
		key, ni, err := support.ReadKey(data, i)
		if err != nil {
			return ni, err
		}
		%[5]s
		if i >= len(data) || data[i] != ':' {
			return i, support.ErrExpectColon
		}
		%[6]s
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		var val %[2]s
		%[3]s
		m[string([]byte(key))] = val
		%[4]s
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		if data[i] == '}' {
			*out = m
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, support.ErrInvalidJSON
		}
		i++
	}
}`, fn, valStr, inner, g.skipWS("i", "i"), g.skipWS("i", "ni"), g.skipWS("i", "i+1"))
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

// supportPkg is the import path of the shared scanner package whose exported
// helpers the generated decoders call.
const supportPkg = "github.com/JohanLindvall/lightning/pkg/support"

func (g *gen) assemble(inPath string, methods []string) string {
	imports := []string{strconv.Quote(supportPkg)}
	if g.needJSON {
		imports = append(imports, `json "encoding/json"`)
	}
	if g.needTime {
		imports = append(imports, `"time"`)
	}

	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated by the lightning generator from %s; DO NOT EDIT.\n\n", inPath)
	fmt.Fprintf(&b, "package %s\n\n", g.pkg)
	b.WriteString("import (\n")
	for _, im := range imports {
		fmt.Fprintf(&b, "\t%s\n", im)
	}
	b.WriteString(")\n\n")

	for _, m := range methods {
		b.WriteString(m)
		b.WriteString("\n\n")
	}
	for _, d := range g.decoders {
		b.WriteString(d)
		b.WriteString("\n\n")
	}
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
func jsonTag(tag *ast.BasicLit) (names []string, skip, nocopy, lax bool) {
	if tag == nil {
		return nil, false, false, false
	}
	s, err := strconv.Unquote(tag.Value)
	if err != nil {
		return nil, false, false, false
	}
	v := reflect.StructTag(s).Get("json")
	if v == "" {
		return nil, false, false, false
	}
	parts := strings.Split(v, ",")
	if parts[0] == "-" && len(parts) == 1 {
		return nil, true, false, false
	}
	for _, o := range parts[1:] {
		switch o {
		case "nocopy":
			nocopy = true
		case "lax":
			lax = true
		}
	}
	for _, n := range strings.Split(parts[0], "|") {
		if n != "" {
			names = append(names, n)
		}
	}
	return names, false, nocopy, lax
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

var intKinds = map[string]bool{
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true, "rune": true,
}

var uintKinds = map[string]bool{
	"uint": true, "uint8": true, "uint16": true, "uint32": true,
	"uint64": true, "uintptr": true, "byte": true,
}
