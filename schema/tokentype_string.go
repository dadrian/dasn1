// Code generated by "stringer -type=TokenType"; DO NOT EDIT.

package schema

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SPACE-0]
	_ = x[COMMENT-1]
	_ = x[ASSIGN-2]
	_ = x[LBRACE-3]
	_ = x[RBRACE-4]
	_ = x[LBRACKET-5]
	_ = x[RBRACKET-6]
	_ = x[LPAREN-7]
	_ = x[RPAREN-8]
	_ = x[COMMA-9]
	_ = x[SEQUENCE_LITERAL-10]
	_ = x[INTEGER_LITERAL-11]
	_ = x[IMPLICIT_LITERAL-12]
	_ = x[EXPLICIT_LITERAL-13]
	_ = x[OPTIONAL_LITERAL-14]
	_ = x[DEFAULT_LITERAL-15]
	_ = x[DEFINED_BY_LITERAL-16]
	_ = x[ANY_LITERAL-17]
	_ = x[OF_LITERAL-18]
	_ = x[CHOICE_LITERAL-19]
	_ = x[SIZE_LITERAL-20]
	_ = x[MAX-21]
	_ = x[BMPSTRING_LITERAL-22]
	_ = x[IA5STRING_LITERAL-23]
	_ = x[PRINTABLE_STRING_LITERAL-24]
	_ = x[TELETEX_STRING_LITERAL-25]
	_ = x[UNIVERSAL_STRING_LITERAL-26]
	_ = x[UTF8_STRING_LITERAL-27]
	_ = x[BIT_STRING_LITERAL-28]
	_ = x[OCTET_STRING_LITERAL-29]
	_ = x[OBJECT_IDENTIFIER_LITERAL-30]
	_ = x[DOTDOT-31]
	_ = x[INTEGER-32]
	_ = x[WORD-33]
}

const _TokenType_name = "SPACECOMMENTASSIGNLBRACERBRACELBRACKETRBRACKETLPARENRPARENCOMMASEQUENCE_LITERALINTEGER_LITERALIMPLICIT_LITERALEXPLICIT_LITERALOPTIONAL_LITERALDEFAULT_LITERALDEFINED_BY_LITERALANY_LITERALOF_LITERALCHOICE_LITERALSIZE_LITERALMAXBMPSTRING_LITERALIA5STRING_LITERALPRINTABLE_STRING_LITERALTELETEX_STRING_LITERALUNIVERSAL_STRING_LITERALUTF8_STRING_LITERALBIT_STRING_LITERALOCTET_STRING_LITERALOBJECT_IDENTIFIER_LITERALDOTDOTINTEGERWORD"

var _TokenType_index = [...]uint16{0, 5, 12, 18, 24, 30, 38, 46, 52, 58, 63, 79, 94, 110, 126, 142, 157, 175, 186, 196, 210, 222, 225, 242, 259, 283, 305, 329, 348, 366, 386, 411, 417, 424, 428}

func (i TokenType) String() string {
	if i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}