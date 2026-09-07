package json

import "github.com/JohanLindvall/lightning/pkg/unstable"

// Kind is what a JSON value is: the six value kinds of the grammar, and
// KindInvalid for bytes no value starts with.
type Kind uint8

const (
	// KindInvalid is an empty input or a leading byte no JSON value begins with.
	KindInvalid Kind = iota
	// KindNull is the literal null.
	KindNull
	// KindBool is the literal true or false.
	KindBool
	// KindNumber is a value that opens like a number: a digit or a sign.
	KindNumber
	// KindString is a value that opens with a quote.
	KindString
	// KindArray is a value that opens with '['.
	KindArray
	// KindObject is a value that opens with '{'.
	KindObject
)

var kindNames = [...]string{"invalid", "null", "bool", "number", "string", "array", "object"}

// String names the kind, lower-case, for messages and tests.
func (k Kind) String() string {
	if int(k) < len(kindNames) {
		return kindNames[k]
	}
	return "invalid"
}

// KindOf classifies a JSON value: the token or container a read function
// returned — [Get], [Lookup], [GetMany], [GetPaths], [ObjectEach],
// [ArrayEach] all hand back the raw bytes and none reports what they hold —
// or a whole document. The read functions decided the kind while scanning
// and threw it away; without this every caller branched on raw[0] itself.
//
// Leading and trailing whitespace is tolerated, as [Get] tolerates it at a
// document's root — and by the same rule, which counts every byte <= 0x20 (see
// the literal cases below). A literal is matched whole — null, true and false with
// nothing else around them — so a misspelling (nul, truex) is KindInvalid, while
// a string, number, array or object is classified by its opening byte alone,
// exactly as the scanner dispatches on it: KindOf says what a value IS, and
// [Valid] says whether it is well-formed. A number opens with a digit, '-'
// or '+' (the sign the rest of this library accepts).
func KindOf(raw []byte) Kind {
	i := unstable.SkipWS(raw, 0)
	if uint(i) >= uint(len(raw)) {
		return KindInvalid
	}
	c := raw[i]
	if k := kindOfByte[c]; k != kindLiteral {
		return Kind(k)
	}
	// A literal must be the whole of what remains. The comparison is against a
	// constant, which the compiler does as one word load and one compare, and
	// what follows is measured with the same SkipWS that skipped what came
	// before: this package's whitespace is every byte <= 0x20, the
	// one-compare shortcut its decoder and Valid take, and a KindOf that
	// tolerated a different set at the two ends of the same value would answer
	// for a document neither of them would accept.
	switch c {
	case 'n':
		if len(raw)-i >= 4 && string(raw[i:i+4]) == "null" && unstable.SkipWS(raw, i+4) == len(raw) {
			return KindNull
		}
	case 't':
		if len(raw)-i >= 4 && string(raw[i:i+4]) == "true" && unstable.SkipWS(raw, i+4) == len(raw) {
			return KindBool
		}
	default:
		if len(raw)-i >= 5 && string(raw[i:i+5]) == "false" && unstable.SkipWS(raw, i+5) == len(raw) {
			return KindBool
		}
	}
	return KindInvalid
}

// kindOfByte maps a value's opening byte to its kind — the dispatch the
// scanner makes, as one load — with kindLiteral standing for the three bytes
// a literal can start with, which need the rest of the token read before they
// can be answered.
const kindLiteral = 0xff

var kindOfByte = [256]uint8{
	'"': uint8(KindString),
	'{': uint8(KindObject),
	'[': uint8(KindArray),
	'-': uint8(KindNumber), '+': uint8(KindNumber),
	'0': uint8(KindNumber), '1': uint8(KindNumber), '2': uint8(KindNumber),
	'3': uint8(KindNumber), '4': uint8(KindNumber), '5': uint8(KindNumber),
	'6': uint8(KindNumber), '7': uint8(KindNumber), '8': uint8(KindNumber),
	'9': uint8(KindNumber),
	'n': kindLiteral, 't': kindLiteral, 'f': kindLiteral,
	// Every other byte is left zero, which is KindInvalid: no JSON value
	// begins with it. Written as a literal rather than built in an init
	// function so the table is data, not startup work.
}
