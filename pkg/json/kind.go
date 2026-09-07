package json

import "github.com/JohanLindvall/lightning/pkg/unstable"

// Kind is what a JSON value is: the six value kinds of the grammar, and
// Invalid for bytes no value starts with.
type Kind uint8

const (
	// Invalid is an empty input or a leading byte no JSON value begins with.
	Invalid Kind = iota
	// Null is the literal null.
	Null
	// Bool is the literal true or false.
	Bool
	// Number is a value that opens like a number: a digit or a sign.
	Number
	// String is a value that opens with a quote.
	String
	// Array is a value that opens with '['.
	Array
	// Object is a value that opens with '{'.
	Object
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
// document's root. A literal is matched whole — null, true and false with
// nothing else around them — so a misspelling (nul, truex) is Invalid, while
// a string, number, array or object is classified by its opening byte alone,
// exactly as the scanner dispatches on it: KindOf says what a value IS, and
// [Valid] says whether it is well-formed. A number opens with a digit, '-'
// or '+' (the sign the rest of this library accepts).
func KindOf(raw []byte) Kind {
	i := unstable.SkipWS(raw, 0)
	if i >= len(raw) {
		return Invalid
	}
	switch raw[i] {
	case '"':
		return String
	case '{':
		return Object
	case '[':
		return Array
	case '-', '+', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return Number
	case 'n', 't', 'f':
		// A literal is the whole of what remains, whitespace aside.
		end := len(raw)
		for end > i && isWS(raw[end-1]) {
			end--
		}
		switch string(raw[i:end]) {
		case "null":
			return Null
		case "true", "false":
			return Bool
		}
	}
	return Invalid
}

// isWS is the JSON grammar's whitespace set.
func isWS(c byte) bool { return c == ' ' || c == '\t' || c == '\n' || c == '\r' }
