package json

import "github.com/JohanLindvall/lightning/pkg/unstable"

// Valid reports whether data is a single well-formed JSON document: exactly one
// value, optionally surrounded by whitespace, with no trailing content.
//
// Valid is the precondition check for this library's own decoders: it accepts
// exactly what DecodeAny accepts, so "Valid(data)" answers "will DecodeAny read
// this?". That equivalence is the one this package tests and holds itself to —
// FuzzValidMatchesDecodeAny runs the two against each other and requires the same
// verdict on every input — and it holds because Valid reuses the decoder's own
// readers rather than reimplementing them.
//
// For a *generated* UnmarshalJSON, Valid is a guide rather than an equivalence,
// and nothing tests it as one: a generated decoder shares these scanners (so the
// divergences from encoding/json listed below apply to it as well) but then adds
// the schema, which cuts both ways. It is stricter where the schema is — a value
// Valid accepts is rejected if its type does not fit the field it lands in — and
// looser in two distinct ways. It is looser where it does not look: an unknown
// field's value is skipped with SkipValue's lenient bracket balancing, not parsed,
// so malformed bytes inside a value the schema ignores can decode even though Valid
// rejects the document. And it is looser where it *does* look, on the numeric
// readers: ReadInt64OrNull and ReadUint64OrNull consume a run of digit-ish bytes
// and stop, so a known int field decodes "1.2.3", "1e" and "1-2" as 1 with no
// error, each of which Valid rejects. The second is the surprising one — the
// leniency is not confined to the parts of the document the schema skips.
// Read Valid as "this is well-formed JSON as lightning defines it", not as
// "every generated decoder will accept this". (A ",lax" field is the one place a
// generated decoder runs this same grammar walk: unstable.SkipValueStrict is what
// lets it tell a tolerated type mismatch from a syntax error it must report.)
//
// That is a deliberately different question from encoding/json.Valid, whose answer
// is "does this match the JSON grammar?", and the two differ in a few places
// because the decoder's scanners are tuned for input already known to be JSON:
//
//   - Numbers follow the scanner's arithmetic, not the JSON grammar. A leading
//     '+', a leading zero, and an empty fraction (+1, 01, 5., .5) are accepted;
//     a magnitude that overflows a float64 (1e309) is rejected, because the
//     decoder cannot represent it. encoding/json.Valid does the opposite on both
//     counts.
//   - Whitespace between tokens is every byte <= 0x20, not just the four bytes
//     JSON names, because that test is one compare instead of four on the
//     decoder's hottest loop.
//   - A raw control byte inside a string is accepted: the string scanners stop at
//     '"' and '\\' only.
//
// Everything else is checked strictly, as the decoder checks it: a trailing comma,
// a non-string object key, a missing colon, an unknown escape, a \u escape without
// four hex digits, an unterminated string, a mismatched bracket, a bad keyword and
// trailing bytes after the value are all rejected. Semantics are not checked —
// duplicate keys are accepted, and neither UTF-8 well-formedness nor surrogate
// pairing is verified (an unpaired \uD800 decodes to U+FFFD rather than failing).
//
// Nesting deeper than unstable.MaxDepth (10000, matching encoding/json's bound) is
// rejected rather than descended into.
//
// Valid neither allocates nor builds the decoded value: strings are scanned in
// place, numbers are read with the decoder's own reader and discarded, and the
// nesting state is a stack-resident bitset. It is therefore far cheaper than
// calling DecodeAny to see whether it errors, and unlike SkipValue's lenient
// bracket balancing (which would accept a trailing comma) it is a real parse.
//
// The walk itself is unstable.SkipValueStrict, which the generated decoders share
// for their ",lax" fields; Valid adds only the "nothing but whitespace follows"
// rule that makes it a whole-document check rather than a single-value one.
func Valid(data []byte) bool {
	i := unstable.SkipWS(data, 0)
	i, err := unstable.SkipValueStrict(data, i)
	if err != nil {
		return false
	}
	return unstable.SkipWS(data, i) == len(data)
}
