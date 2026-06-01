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
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
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
	file, err := parser.ParseFile(fset, inPath, nil, parser.SkipObjectResolution)
	if err != nil {
		return err
	}

	g := &gen{
		fset:        fset,
		pkg:         file.Name.Name,
		structTypes: map[string]*ast.StructType{},
		order:       nil,
		used:        map[string]bool{},
		memo:        map[string]string{},
	}

	// Collect every top-level struct type, in source order.
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
			}
		}
	}
	if len(g.order) == 0 {
		return fmt.Errorf("no top-level struct types found")
	}

	var methods []string
	for _, name := range g.order {
		methods = append(methods, g.genUnmarshal(name))
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

	decoders []string // generated decoder functions, in creation order

	needBool, needUint, needFloat, needAny, needJSON, needNoCopyStr bool
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
	key := "named:" + name
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	fn := "decode" + name
	g.used[fn] = true
	g.memo[key] = fn // set before body so recursive types terminate
	g.genStructBody(fn, "*"+name, g.structTypes[name])
	return fn
}

// anonStruct returns the decoder for an anonymous struct type, named after the
// field it was first reached through.
func (g *gen) anonStruct(t *ast.StructType, hint string) string {
	key := "anon:" + g.typeStr(t)
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	fn := g.uniq("decode" + cap1(hint))
	g.memo[key] = fn
	g.genStructBody(fn, "*"+g.typeStr(t), t)
	return fn
}

func (g *gen) genStructBody(fn, paramType string, st *ast.StructType) {
	var cases strings.Builder
	for _, f := range st.Fields.List {
		if len(f.Names) == 0 {
			fmt.Fprintf(os.Stderr, "warning: skipping embedded field %s\n", g.typeStr(f.Type))
			continue
		}
		tagName, skip, nocopy := jsonTag(f.Tag)
		if skip {
			continue
		}
		for _, nm := range f.Names {
			field := nm.Name
			if !ast.IsExported(field) {
				continue
			}
			key := field
			if tagName != "" {
				key = tagName
			}
			code := g.field("v."+field, f.Type, field, nocopy)
			fmt.Fprintf(&cases, "\tcase %q:\n%s\n", key, code)
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
		i = support.SkipWS(data, i)
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
		i = support.SkipWS(data, ni)
		if i >= len(data) || data[i] != ':' {
			return i, support.ErrExpectColon
		}
		i = support.SkipWS(data, i+1)
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
		i = support.SkipWS(data, i)
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
}`, fn, paramType, cases.String())
	g.decoders = append(g.decoders, body)
}

// field returns the code that decodes a JSON value at data[i] into the Go
// lvalue dest, advancing i and returning (end, err) on failure. hint suggests a
// name for any decoder this value needs. nocopy requests that string/raw leaves
// alias the input bytes; it propagates through slices, maps and pointers but
// stops at struct boundaries, where each field's own tag governs.
func (g *gen) field(dest string, expr ast.Expr, hint string, nocopy bool) string {
	switch t := expr.(type) {
	case *ast.ParenExpr:
		return g.field(dest, t.X, hint, nocopy)

	case *ast.StarExpr:
		inner := g.field("(*"+dest+")", t.X, hint, nocopy)
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
			// A struct's own field tags govern its nocopy behavior.
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
		fmt.Fprintf(os.Stderr, "warning: unsupported type %s for %s; value skipped\n", g.typeStr(t), dest)
		return g.skipEmit()

	case *ast.StructType:
		return g.callDecoder(dest, g.anonStruct(t, hint))

	case *ast.ArrayType:
		if t.Len != nil {
			fmt.Fprintf(os.Stderr, "warning: fixed-size array for %s; value skipped\n", dest)
			return g.skipEmit()
		}
		return g.callDecoder(dest, g.sliceDecoder(t.Elt, hint, nocopy))

	case *ast.MapType:
		fn := g.mapDecoder(t.Key, t.Value, hint, nocopy)
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

func (g *gen) sliceDecoder(elt ast.Expr, hint string, nocopy bool) string {
	suffix := ""
	if nocopy {
		suffix = "NoCopy"
	}
	key := "slice:" + suffix + ":" + g.typeStr(elt)
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	var fn string
	if g.isStruct(elt) {
		fn = g.uniq("decode" + cap1(hint) + suffix)
	} else {
		fn = g.uniq("decode" + g.baseName(elt) + suffix + "Slice")
	}
	g.memo[key] = fn
	if isRaw(elt) {
		g.needJSON = true
	}
	eltStr := g.typeStr(elt)
	inner := g.field("el", elt, singular(hint)+"Entry", nocopy)
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
	i++
	for {
		i = support.SkipWS(data, i)
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		var el %[2]s
		%[3]s
		*out = append(*out, el)
		i = support.SkipWS(data, i)
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
}`, fn, eltStr, inner)
	g.decoders = append(g.decoders, body)
	return fn
}

func (g *gen) mapDecoder(keyExpr, valExpr ast.Expr, hint string, nocopy bool) string {
	if g.typeStr(keyExpr) != "string" {
		fmt.Fprintf(os.Stderr, "warning: map key type %s unsupported for %s; value skipped\n", g.typeStr(keyExpr), hint)
		return ""
	}
	suffix := ""
	if nocopy {
		suffix = "NoCopy"
	}
	key := "map:" + suffix + ":" + g.typeStr(valExpr)
	if fn, ok := g.memo[key]; ok {
		return fn
	}
	fn := g.uniq("decode" + cap1(hint) + suffix + "Map")
	g.memo[key] = fn
	if isRaw(valExpr) {
		g.needJSON = true
	}
	valStr := g.typeStr(valExpr)
	inner := g.field("val", valExpr, hint+"Value", nocopy)
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
		i = support.SkipWS(data, i)
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
		i = support.SkipWS(data, ni)
		if i >= len(data) || data[i] != ':' {
			return i, support.ErrExpectColon
		}
		i = support.SkipWS(data, i+1)
		if i >= len(data) {
			return i, support.ErrTruncated
		}
		var val %[2]s
		%[3]s
		m[string([]byte(key))] = val
		i = support.SkipWS(data, i)
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
}`, fn, valStr, inner)
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

// jsonTag returns the JSON key from a struct tag, whether the field is excluded
// (`json:"-"`), and whether the "nocopy" option is present. nocopy makes string
// and raw fields alias the input bytes instead of copying them.
func jsonTag(tag *ast.BasicLit) (name string, skip, nocopy bool) {
	if tag == nil {
		return "", false, false
	}
	s, err := strconv.Unquote(tag.Value)
	if err != nil {
		return "", false, false
	}
	v := reflect.StructTag(s).Get("json")
	if v == "" {
		return "", false, false
	}
	parts := strings.Split(v, ",")
	if parts[0] == "-" && len(parts) == 1 {
		return "", true, false
	}
	for _, o := range parts[1:] {
		if o == "nocopy" {
			nocopy = true
		}
	}
	return parts[0], false, nocopy
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

func isRaw(expr ast.Expr) bool {
	if p, ok := expr.(*ast.ParenExpr); ok {
		return isRaw(p.X)
	}
	sel, ok := expr.(*ast.SelectorExpr)
	return ok && (sel.Sel.Name == "RawMessage" || sel.Sel.Name == "RawValue")
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
