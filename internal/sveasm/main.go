// Command sveasm keeps the hand-encoded SVE instructions in the arm64 assembly
// honest by deriving them, rather than trusting that they were pasted correctly.
//
// The Go assembler has no SVE mnemonics (checked on Go 1.26), so pkg/unstable's
// SVE2 scanners spell those instructions as WORD constants. A raw 32-bit
// constant is unreviewable and a single wrong nibble is a silently different
// instruction, so the mnemonic — not the hex — is the source of truth: every
// such line carries the instruction it means in its comment,
//
//	WORD $0x45218402 // match p2.b, p1/z, z0.b, z1.b
//	WORD $0x25211c41 // whilelo p1.b, x2, x1 // active lanes = min(VL, len-x2)
//
// and this command assembles those mnemonics with the GNU assembler and either
// rewrites the constants from them (-w, the generate direction) or reports any
// that disagree (the default, the CI gate). Anything after a second "//" is
// prose and is preserved untouched.
//
// Usage:
//
//	sveasm [-w] file.s...
//
// The assembler is picked in this order: $SVEAS, then aarch64-linux-gnu-as (the
// cross-assembler, so the check also runs on an amd64 host), then plain "as" if
// it accepts aarch64. Ubuntu ships the cross-assembler in binutils-aarch64-linux-gnu
// and the native one in binutils.
//
// Note this validates the ENCODING against the mnemonic, not the mnemonic
// against the algorithm. That a scanner computes the right answer is the job of
// the differential tests (TestIndexVariantsFlip, TestIndexArm64Lengths), which
// run the assembled code against a scalar oracle.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// wordLine matches an emitted SVE instruction: the WORD constant and the comment
// after it. The comment is split into mnemonic and prose by splitComment rather
// than by this regexp, because a mnemonic contains single slashes of its own
// (every predicate operand is written "p1/z") and only a DOUBLE slash starts the
// prose.
var wordLine = regexp.MustCompile(`^(\s*WORD\s+\$0x)([0-9a-fA-F]{8})(\s*//\s*)(.*)$`)

// splitComment divides a WORD line's comment into the mnemonic and the trailing
// prose, which begins at the first "//" — as opposed to the single slashes that
// appear inside the mnemonic itself.
func splitComment(comment string) (mnemonic, prose string) {
	if i := strings.Index(comment, "//"); i >= 0 {
		return strings.TrimSpace(comment[:i]), comment[i:]
	}
	return strings.TrimSpace(comment), ""
}

// objdumpLine matches one disassembled instruction: "   4:\t2538cb82 \tmov\tz2.b, #92".
var objdumpLine = regexp.MustCompile(`^\s*[0-9a-f]+:\s+([0-9a-f]{8})\s+(.*)$`)

// arch is the .arch directive the mnemonics are assembled under. SVE2 covers
// MATCH; everything else the scanners use is base SVE.
const arch = "armv8.6-a+sve2"

func main() {
	write := flag.Bool("w", false, "rewrite the WORD constants from their mnemonics instead of checking them")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: sveasm [-w] file.s...")
		os.Exit(2)
	}

	as, objdump, err := findTools()
	if err != nil {
		fmt.Fprintf(os.Stderr, "sveasm: %v\n", err)
		os.Exit(1)
	}

	bad := 0
	for _, path := range flag.Args() {
		n, err := process(path, as, objdump, *write)
		if err != nil {
			fmt.Fprintf(os.Stderr, "sveasm: %s: %v\n", path, err)
			os.Exit(1)
		}
		bad += n
	}
	if bad > 0 {
		fmt.Fprintf(os.Stderr, "\nsveasm: %d encoding(s) disagree with their mnemonic; run `make sveasm` to regenerate\n", bad)
		os.Exit(1)
	}
}

// process checks (or rewrites) one assembly file, returning how many WORD
// constants disagreed with their mnemonic.
func process(path, as, objdump string, write bool) (int, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(src), "\n")

	// Collect the mnemonics in file order.
	type site struct {
		line     int    // index into lines
		hex      string // as written in the file
		mnemonic string
	}
	var sites []site
	for i, line := range lines {
		m := wordLine.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		mn, _ := splitComment(m[4])
		if mn == "" {
			return 0, fmt.Errorf("line %d: WORD constant has no mnemonic in its comment", i+1)
		}
		sites = append(sites, site{line: i, hex: strings.ToLower(m[2]), mnemonic: mn})
	}
	if len(sites) == 0 {
		fmt.Printf("%s: no WORD constants\n", path)
		return 0, nil
	}

	mnemonics := make([]string, len(sites))
	for i, s := range sites {
		mnemonics[i] = s.mnemonic
	}
	encs, err := assemble(as, objdump, mnemonics)
	if err != nil {
		return 0, err
	}
	// One mnemonic must produce exactly one instruction, or the positional
	// mapping back to the WORD lines is meaningless.
	if len(encs) != len(sites) {
		return 0, fmt.Errorf("assembled %d instructions from %d mnemonics: a mnemonic expanded to more than one instruction", len(encs), len(sites))
	}

	// For a mismatch the useful half of the message is what the constant in the
	// file actually encodes — decode the stale values, in one batch.
	var stale []string
	for i, s := range sites {
		if s.hex != encs[i] {
			stale = append(stale, s.hex)
		}
	}
	staleText := map[string]string{}
	if len(stale) > 0 && !write {
		if texts, err := disassemble(as, objdump, stale); err == nil {
			for i, h := range stale {
				staleText[h] = texts[i]
			}
		}
	}

	bad := 0
	for i, s := range sites {
		if s.hex == encs[i] {
			continue
		}
		bad++
		if write {
			m := wordLine.FindStringSubmatch(lines[s.line])
			mn, prose := splitComment(m[4])
			rest := mn
			if prose != "" {
				rest += " " + prose
			}
			lines[s.line] = m[1] + encs[i] + m[3] + rest
			fmt.Printf("%s:%d: %s -> %s  (%s)\n", path, s.line+1, s.hex, encs[i], s.mnemonic)
		} else {
			actual := staleText[s.hex]
			if actual == "" {
				actual = "(not a valid instruction)"
			}
			fmt.Fprintf(os.Stderr, "%s:%d: 0x%s encodes %q, but the comment says %q (which is 0x%s)\n",
				path, s.line+1, s.hex, actual, s.mnemonic, encs[i])
		}
	}
	if write {
		if bad > 0 {
			if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0o644); err != nil {
				return 0, err
			}
			fmt.Printf("%s: rewrote %d of %d encodings\n", path, bad, len(sites))
		} else {
			fmt.Printf("%s: %d encodings already up to date\n", path, len(sites))
		}
		return 0, nil // -w resolves the disagreements rather than reporting them
	}
	if bad == 0 {
		fmt.Printf("%s: %d encodings match their mnemonics\n", path, len(sites))
	}
	return bad, nil
}

// assemble runs the mnemonics through the assembler and returns, per mnemonic,
// its 8-hex-digit encoding.
func assemble(as, objdump string, mnemonics []string) (encs []string, err error) {
	dir, err := os.MkdirTemp("", "sveasm")
	if err != nil {
		return nil, err
	}
	defer func() { _ = os.RemoveAll(dir) }()

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\t.arch %s\n\t.text\n", arch)
	for _, m := range mnemonics {
		fmt.Fprintf(&buf, "\t%s\n", m)
	}
	srcPath := filepath.Join(dir, "sve.s")
	objPath := filepath.Join(dir, "sve.o")
	if err := os.WriteFile(srcPath, buf.Bytes(), 0o644); err != nil {
		return nil, err
	}

	var stderr bytes.Buffer
	cmd := exec.Command(as, "-o", objPath, srcPath)
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s failed: %v\n%s", as, err, cleanAsErrors(stderr.String(), srcPath, mnemonics))
	}

	out, err := exec.Command(objdump, "-d", objPath).Output()
	if err != nil {
		return nil, fmt.Errorf("%s failed: %v", objdump, err)
	}
	for _, line := range strings.Split(string(out), "\n") {
		m := objdumpLine.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		encs = append(encs, m[1])
	}
	return encs, nil
}

// disassemble decodes raw 32-bit words back into mnemonics, so a mismatch can be
// reported as what the constant in the file actually means. Uses .inst, which
// emits the word verbatim without the assembler trying to interpret it.
func disassemble(as, objdump string, hexes []string) ([]string, error) {
	dir, err := os.MkdirTemp("", "svedis")
	if err != nil {
		return nil, err
	}
	defer func() { _ = os.RemoveAll(dir) }()

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\t.arch %s\n\t.text\n", arch)
	for _, h := range hexes {
		fmt.Fprintf(&buf, "\t.inst 0x%s\n", h)
	}
	src := filepath.Join(dir, "d.s")
	obj := filepath.Join(dir, "d.o")
	if err := os.WriteFile(src, buf.Bytes(), 0o644); err != nil {
		return nil, err
	}
	if err := exec.Command(as, "-o", obj, src).Run(); err != nil {
		return nil, err
	}
	out, err := exec.Command(objdump, "-d", obj).Output()
	if err != nil {
		return nil, err
	}
	var texts []string
	for _, line := range strings.Split(string(out), "\n") {
		if m := objdumpLine.FindStringSubmatch(line); m != nil {
			texts = append(texts, strings.Join(strings.Fields(strings.ReplaceAll(m[2], "\t", " ")), " "))
		}
	}
	if len(texts) != len(hexes) {
		return nil, fmt.Errorf("disassembled %d of %d words", len(texts), len(hexes))
	}
	return texts, nil
}

// cleanAsErrors rewrites the assembler's "sve.s:7: Error: ..." line references
// into the mnemonic that failed, since the temporary file is gone by the time
// anyone reads the message.
func cleanAsErrors(msg, srcPath string, mnemonics []string) string {
	re := regexp.MustCompile(regexp.QuoteMeta(srcPath) + `:(\d+):`)
	return re.ReplaceAllStringFunc(msg, func(s string) string {
		var n int
		if _, err := fmt.Sscanf(re.FindStringSubmatch(s)[1], "%d", &n); err != nil {
			return s
		}
		// Line 1 is .arch, line 2 is .text, so mnemonic i is on line i+3.
		if i := n - 3; i >= 0 && i < len(mnemonics) {
			return fmt.Sprintf("%q:", mnemonics[i])
		}
		return s
	})
}

// findTools locates an assembler and disassembler that understand aarch64. The
// cross-tools are preferred so the check runs on an amd64 host too; a native
// arm64 host has them under the plain names.
func findTools() (as, objdump string, err error) {
	if v := os.Getenv("SVEAS"); v != "" {
		return v, envOr("SVEOBJDUMP", "objdump"), nil
	}
	for _, pair := range [][2]string{
		{"aarch64-linux-gnu-as", "aarch64-linux-gnu-objdump"},
		{"as", "objdump"},
	} {
		asPath, err1 := exec.LookPath(pair[0])
		odPath, err2 := exec.LookPath(pair[1])
		if err1 != nil || err2 != nil {
			continue
		}
		if !acceptsAArch64(asPath) {
			continue
		}
		return asPath, odPath, nil
	}
	return "", "", fmt.Errorf("no aarch64 assembler found: install binutils-aarch64-linux-gnu (any host) " +
		"or binutils (on an arm64 host), or set $SVEAS to one")
}

// acceptsAArch64 reports whether this assembler can assemble an SVE2
// instruction — the cheapest possible probe of "is this the right target".
func acceptsAArch64(as string) bool {
	dir, err := os.MkdirTemp("", "sveprobe")
	if err != nil {
		return false
	}
	defer func() { _ = os.RemoveAll(dir) }()
	src := filepath.Join(dir, "p.s")
	if err := os.WriteFile(src, []byte("\t.arch "+arch+"\n\t.text\n\tmatch p2.b, p1/z, z0.b, z1.b\n"), 0o644); err != nil {
		return false
	}
	cmd := exec.Command(as, "-o", filepath.Join(dir, "p.o"), src)
	cmd.Stderr = nil
	return cmd.Run() == nil
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
