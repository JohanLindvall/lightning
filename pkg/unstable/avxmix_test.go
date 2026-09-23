package unstable

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// avxMixViolations scans amd64 assembly source for the SSE/AVX transition
// penalty: a legacy-SSE instruction — any non-VEX vector instruction on an X
// register — executed while the upper halves of the vector registers are
// dirty, i.e. after an instruction that writes a Y or Z register and before
// the next VZEROUPPER. On Intel cores that costs a microcode assist
// (assists.sse_avx_mix) per transition; AMD does not penalise it, so a body
// measured only on Zen 4 passes every benchmark there with the bug in place.
// countKernel shipped that way — MOVQ R11, X1 after a ymm load, once per
// presized array — and cost update_center +31% and synthea_fhir +24% on
// Meteor Lake. A RET or a tail JMP while dirty is the same bug one call
// later, and is reported too.
//
// It follows program order, not control flow: a function's first vector code
// may be SSE (the string scanners' first block, before they switch to AVX2)
// and SSE after a VZEROUPPER is clean, which program order captures for every
// body in the tree; a backward jump from YMM code into earlier SSE code would
// escape it, and none exists. Macros are expanded in place, so an instruction
// inside CLASSIFY or CONVERT counts where the macro is used.
func avxMixViolations(src string) []string {
	lines := strings.Split(src, "\n")
	macros := map[string][]string{}
	var body []string
	for k := 0; k < len(lines); k++ {
		l := lines[k]
		m := regexp.MustCompile(`^#define\s+(\w+)`).FindStringSubmatch(l)
		if m == nil {
			body = append(body, l)
			continue
		}
		var def []string
		for {
			code := strings.TrimSuffix(strings.TrimSpace(stripAsmComment(lines[k])), `\`)
			def = append(def, code)
			if !strings.HasSuffix(strings.TrimSpace(stripAsmComment(lines[k])), `\`) || k+1 >= len(lines) {
				break
			}
			k++
		}
		def[0] = strings.TrimSpace(strings.TrimPrefix(def[0], m[0]))
		if i := strings.Index(def[0], ")"); strings.HasPrefix(def[0], "(") && i >= 0 {
			def[0] = def[0][i+1:]
		}
		macros[m[1]] = def
	}
	var expand func(string, int) []string
	expand = func(l string, depth int) []string {
		code := strings.TrimSpace(stripAsmComment(l))
		name := regexp.MustCompile(`^(\w+)\b`).FindString(code)
		if def, ok := macros[name]; ok && depth < 8 {
			var out []string
			for _, d := range def {
				for _, s := range strings.Split(d, ";") {
					out = append(out, expand(s, depth+1)...)
				}
			}
			return out
		}
		return []string{code}
	}
	// Only registers 0-15 carry the state: legacy SSE cannot address 16-31,
	// so a body confined to them (the VBMI structural scan) needs no
	// VZEROUPPER, the trick glibc's -evex string functions use.
	yz := regexp.MustCompile(`\b[YZ]([0-9]|1[0-5])\b`)
	xreg := regexp.MustCompile(`\bX\d+\b`)
	var bad []string
	fn, dirty := "", false
	for _, l := range body {
		if m := regexp.MustCompile(`^TEXT\s+·(\w+)`).FindStringSubmatch(l); m != nil {
			fn, dirty = m[1], false
			continue
		}
		for _, code := range expand(l, 0) {
			op := regexp.MustCompile(`^(\w+)\s*(.*)$`).FindStringSubmatch(code)
			if op == nil || strings.HasSuffix(op[1], ":") {
				continue
			}
			mn, args := op[1], op[2]
			switch {
			case mn == "VZEROUPPER" || mn == "VZEROALL":
				dirty = false
			case strings.HasPrefix(mn, "V") && yz.MatchString(args):
				dirty = true
			case dirty && (mn == "RET" || mn == "JMP" && strings.Contains(args, "(SB)")):
				// Leaving dirty hands the penalty to the caller, whose first
				// SSE instruction is often the ABI wrapper's XORPS X15, X15
				// (or to the tail-called body, which assumes a clean state).
				bad = append(bad, fn+": "+code+" with the upper state dirty")
			case dirty && !strings.HasPrefix(mn, "V") && !strings.HasPrefix(mn, "K") && xreg.MatchString(args):
				bad = append(bad, fn+": "+code)
			}
		}
	}
	return bad
}

// stripAsmComment drops a // or /* */ comment from an assembly line.
func stripAsmComment(l string) string {
	if i := strings.Index(l, "//"); i >= 0 {
		l = l[:i]
	}
	if i := strings.Index(l, "/*"); i >= 0 {
		l = l[:i]
	}
	return l
}

// TestNoSSEAfterAVX holds every amd64 assembly file to avxMixViolations. It
// reads the source, so it runs on every host: the only CPUs that feel the bug
// are Intel's, and they are not what every CI runner or development box is.
func TestNoSSEAfterAVX(t *testing.T) {
	files, err := filepath.Glob("*_amd64.s")
	if err != nil || len(files) == 0 {
		t.Fatalf("no amd64 assembly found: %v", err)
	}
	for _, f := range files {
		src, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range avxMixViolations(string(src)) {
			t.Errorf("%s: legacy SSE with dirty upper vector state: %s", f, v)
		}
	}
}

// TestAVXMixCheckFindsTheBug is the check's premise: the countKernel prologue
// that shipped with the bug, and the shapes the rule must accept.
func TestAVXMixCheckFindsTheBug(t *testing.T) {
	for _, c := range []struct {
		src  string
		want int
	}{
		{"TEXT ·f(SB), NOSPLIT, $0\n\tVMOVDQU cbcClose<>(SB), Y0\n\tMOVQ    R11, X1\n\tVPBROADCASTB X1, Y1\n\tVZEROUPPER\n\tRET\n", 1},
		{"TEXT ·f(SB), NOSPLIT, $0\n\tMOVQ    R11, X1\n\tVMOVDQU cbcClose<>(SB), Y0\n\tVZEROUPPER\n\tRET\n", 0},
		{"TEXT ·f(SB), NOSPLIT, $0\n\tVMOVDQU (SI), Y0\n\tVZEROUPPER\n\tMOVOU (SI), X2\n\tRET\n", 0},
		{"#define M \\\n\tVPCMPEQB Y1, Y2, Y3 \\\n\tVPMOVMSKB Y3, AX\n\nTEXT ·f(SB), NOSPLIT, $0\n\tM\n\tPCMPEQB X1, X2\n\tVZEROUPPER\n\tRET\n", 1},
		{"TEXT ·f(SB), NOSPLIT, $0\n\tVMOVDQU (SI), Y0\n\tVMOVQ R11, X1\n\tVZEROUPPER\n\tRET\n", 0},
		{"TEXT ·f(SB), NOSPLIT, $0\n\tVMOVDQU (SI), Y0\n\tRET\n", 1},
		{"TEXT ·f(SB), NOSPLIT, $0\n\tVMOVDQU (SI), Y0\n\tJMP ·g(SB)\n", 1},
		{"TEXT ·f(SB), NOSPLIT, $0\n\tVMOVDQU (SI), Y0\n\tJMP loop\n", 0},
		{"TEXT ·f(SB), NOSPLIT, $0\n\tVMOVDQU64 (SI), Z16\n\tVPCMPEQB Z17, Z16, K1\n\tRET\n", 0},
	} {
		if got := len(avxMixViolations(c.src)); got != c.want {
			t.Errorf("%d violations, want %d, in:\n%s", got, c.want, c.src)
		}
	}
}
