// Code generated from ASN1Schema.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ASN1Schema

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 258,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 6, 2, 74, 10, 2, 13, 2, 14, 2,
	75, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 85, 10, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 95, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 104, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 5, 8, 114, 10, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 126, 10, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 12, 5, 12, 133, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 5, 12, 141, 10, 12, 5, 12, 143, 10, 12, 3, 12, 5, 12, 146, 10, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 153, 10, 13, 3, 14, 3, 14, 6,
	14, 157, 10, 14, 13, 14, 14, 14, 158, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	5, 15, 166, 10, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 173, 10,
	17, 3, 17, 3, 17, 6, 17, 177, 10, 17, 13, 17, 14, 17, 178, 5, 17, 181,
	10, 17, 3, 18, 3, 18, 5, 18, 185, 10, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21, 198, 10, 21, 3, 22,
	3, 22, 5, 22, 202, 10, 22, 3, 23, 3, 23, 5, 23, 206, 10, 23, 3, 24, 3,
	24, 5, 24, 210, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26,
	5, 26, 219, 10, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5, 29, 236, 10, 29,
	3, 30, 3, 30, 5, 30, 240, 10, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3,
	32, 5, 32, 248, 10, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 2, 2, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
	64, 66, 68, 70, 2, 5, 3, 2, 14, 15, 3, 2, 24, 29, 3, 2, 34, 35, 2, 259,
	2, 73, 3, 2, 2, 2, 4, 84, 3, 2, 2, 2, 6, 94, 3, 2, 2, 2, 8, 103, 3, 2,
	2, 2, 10, 105, 3, 2, 2, 2, 12, 109, 3, 2, 2, 2, 14, 113, 3, 2, 2, 2, 16,
	115, 3, 2, 2, 2, 18, 118, 3, 2, 2, 2, 20, 129, 3, 2, 2, 2, 22, 132, 3,
	2, 2, 2, 24, 147, 3, 2, 2, 2, 26, 154, 3, 2, 2, 2, 28, 162, 3, 2, 2, 2,
	30, 167, 3, 2, 2, 2, 32, 180, 3, 2, 2, 2, 34, 184, 3, 2, 2, 2, 36, 186,
	3, 2, 2, 2, 38, 188, 3, 2, 2, 2, 40, 195, 3, 2, 2, 2, 42, 199, 3, 2, 2,
	2, 44, 203, 3, 2, 2, 2, 46, 209, 3, 2, 2, 2, 48, 211, 3, 2, 2, 2, 50, 216,
	3, 2, 2, 2, 52, 223, 3, 2, 2, 2, 54, 228, 3, 2, 2, 2, 56, 232, 3, 2, 2,
	2, 58, 237, 3, 2, 2, 2, 60, 243, 3, 2, 2, 2, 62, 247, 3, 2, 2, 2, 64, 249,
	3, 2, 2, 2, 66, 251, 3, 2, 2, 2, 68, 253, 3, 2, 2, 2, 70, 255, 3, 2, 2,
	2, 72, 74, 5, 4, 3, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73,
	3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 7, 2, 2, 3,
	78, 3, 3, 2, 2, 2, 79, 80, 5, 62, 32, 2, 80, 81, 7, 4, 2, 2, 81, 82, 5,
	6, 4, 2, 82, 85, 3, 2, 2, 2, 83, 85, 5, 38, 20, 2, 84, 79, 3, 2, 2, 2,
	84, 83, 3, 2, 2, 2, 85, 5, 3, 2, 2, 2, 86, 95, 5, 46, 24, 2, 87, 95, 5,
	52, 27, 2, 88, 95, 5, 24, 13, 2, 89, 95, 5, 36, 19, 2, 90, 95, 5, 40, 21,
	2, 91, 95, 5, 42, 22, 2, 92, 95, 5, 44, 23, 2, 93, 95, 5, 54, 28, 2, 94,
	86, 3, 2, 2, 2, 94, 87, 3, 2, 2, 2, 94, 88, 3, 2, 2, 2, 94, 89, 3, 2, 2,
	2, 94, 90, 3, 2, 2, 2, 94, 91, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 93,
	3, 2, 2, 2, 95, 7, 3, 2, 2, 2, 96, 104, 7, 12, 2, 2, 97, 104, 7, 13, 2,
	2, 98, 104, 7, 32, 2, 2, 99, 104, 7, 31, 2, 2, 100, 104, 7, 30, 2, 2, 101,
	104, 5, 60, 31, 2, 102, 104, 7, 19, 2, 2, 103, 96, 3, 2, 2, 2, 103, 97,
	3, 2, 2, 2, 103, 98, 3, 2, 2, 2, 103, 99, 3, 2, 2, 2, 103, 100, 3, 2, 2,
	2, 103, 101, 3, 2, 2, 2, 103, 102, 3, 2, 2, 2, 104, 9, 3, 2, 2, 2, 105,
	106, 7, 7, 2, 2, 106, 107, 7, 34, 2, 2, 107, 108, 7, 8, 2, 2, 108, 11,
	3, 2, 2, 2, 109, 110, 9, 2, 2, 2, 110, 13, 3, 2, 2, 2, 111, 114, 7, 16,
	2, 2, 112, 114, 5, 16, 9, 2, 113, 111, 3, 2, 2, 2, 113, 112, 3, 2, 2, 2,
	114, 15, 3, 2, 2, 2, 115, 116, 7, 17, 2, 2, 116, 117, 5, 66, 34, 2, 117,
	17, 3, 2, 2, 2, 118, 119, 7, 22, 2, 2, 119, 120, 7, 9, 2, 2, 120, 121,
	7, 34, 2, 2, 121, 125, 7, 33, 2, 2, 122, 126, 7, 34, 2, 2, 123, 126, 5,
	68, 35, 2, 124, 126, 7, 23, 2, 2, 125, 122, 3, 2, 2, 2, 125, 123, 3, 2,
	2, 2, 125, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 7, 10, 2, 2,
	128, 19, 3, 2, 2, 2, 129, 130, 5, 18, 10, 2, 130, 21, 3, 2, 2, 2, 131,
	133, 5, 12, 7, 2, 132, 131, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 142,
	3, 2, 2, 2, 134, 143, 5, 6, 4, 2, 135, 140, 5, 62, 32, 2, 136, 137, 7,
	9, 2, 2, 137, 138, 5, 20, 11, 2, 138, 139, 7, 10, 2, 2, 139, 141, 3, 2,
	2, 2, 140, 136, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 143, 3, 2, 2, 2,
	142, 134, 3, 2, 2, 2, 142, 135, 3, 2, 2, 2, 143, 145, 3, 2, 2, 2, 144,
	146, 5, 14, 8, 2, 145, 144, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 23,
	3, 2, 2, 2, 147, 152, 7, 13, 2, 2, 148, 149, 7, 5, 2, 2, 149, 150, 5, 28,
	15, 2, 150, 151, 7, 6, 2, 2, 151, 153, 3, 2, 2, 2, 152, 148, 3, 2, 2, 2,
	152, 153, 3, 2, 2, 2, 153, 25, 3, 2, 2, 2, 154, 156, 7, 9, 2, 2, 155, 157,
	7, 34, 2, 2, 156, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 156, 3, 2,
	2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161, 7, 10, 2, 2,
	161, 27, 3, 2, 2, 2, 162, 165, 5, 30, 16, 2, 163, 164, 7, 11, 2, 2, 164,
	166, 5, 28, 15, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 29,
	3, 2, 2, 2, 167, 168, 5, 68, 35, 2, 168, 169, 5, 26, 14, 2, 169, 31, 3,
	2, 2, 2, 170, 172, 5, 34, 18, 2, 171, 173, 5, 32, 17, 2, 172, 171, 3, 2,
	2, 2, 172, 173, 3, 2, 2, 2, 173, 181, 3, 2, 2, 2, 174, 176, 5, 34, 18,
	2, 175, 177, 7, 34, 2, 2, 176, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178,
	176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 3, 2, 2, 2, 180, 170,
	3, 2, 2, 2, 180, 174, 3, 2, 2, 2, 181, 33, 3, 2, 2, 2, 182, 185, 5, 64,
	33, 2, 183, 185, 5, 30, 16, 2, 184, 182, 3, 2, 2, 2, 184, 183, 3, 2, 2,
	2, 185, 35, 3, 2, 2, 2, 186, 187, 7, 32, 2, 2, 187, 37, 3, 2, 2, 2, 188,
	189, 5, 68, 35, 2, 189, 190, 7, 32, 2, 2, 190, 191, 7, 4, 2, 2, 191, 192,
	7, 5, 2, 2, 192, 193, 5, 32, 17, 2, 193, 194, 7, 6, 2, 2, 194, 39, 3, 2,
	2, 2, 195, 197, 7, 31, 2, 2, 196, 198, 5, 20, 11, 2, 197, 196, 3, 2, 2,
	2, 197, 198, 3, 2, 2, 2, 198, 41, 3, 2, 2, 2, 199, 201, 7, 30, 2, 2, 200,
	202, 5, 20, 11, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 43,
	3, 2, 2, 2, 203, 205, 5, 60, 31, 2, 204, 206, 5, 20, 11, 2, 205, 204, 3,
	2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 45, 3, 2, 2, 2, 207, 210, 5, 48, 25,
	2, 208, 210, 5, 50, 26, 2, 209, 207, 3, 2, 2, 2, 209, 208, 3, 2, 2, 2,
	210, 47, 3, 2, 2, 2, 211, 212, 7, 12, 2, 2, 212, 213, 7, 5, 2, 2, 213,
	214, 5, 56, 29, 2, 214, 215, 7, 6, 2, 2, 215, 49, 3, 2, 2, 2, 216, 218,
	7, 12, 2, 2, 217, 219, 5, 18, 10, 2, 218, 217, 3, 2, 2, 2, 218, 219, 3,
	2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 7, 20, 2, 2, 221, 222, 5, 62,
	32, 2, 222, 51, 3, 2, 2, 2, 223, 224, 7, 21, 2, 2, 224, 225, 7, 5, 2, 2,
	225, 226, 5, 56, 29, 2, 226, 227, 7, 6, 2, 2, 227, 53, 3, 2, 2, 2, 228,
	229, 7, 19, 2, 2, 229, 230, 7, 18, 2, 2, 230, 231, 5, 68, 35, 2, 231, 55,
	3, 2, 2, 2, 232, 235, 5, 58, 30, 2, 233, 234, 7, 11, 2, 2, 234, 236, 5,
	56, 29, 2, 235, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 57, 3, 2, 2,
	2, 237, 239, 5, 68, 35, 2, 238, 240, 5, 10, 6, 2, 239, 238, 3, 2, 2, 2,
	239, 240, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 242, 5, 22, 12, 2, 242,
	59, 3, 2, 2, 2, 243, 244, 9, 3, 2, 2, 244, 61, 3, 2, 2, 2, 245, 248, 5,
	8, 5, 2, 246, 248, 5, 70, 36, 2, 247, 245, 3, 2, 2, 2, 247, 246, 3, 2,
	2, 2, 248, 63, 3, 2, 2, 2, 249, 250, 5, 68, 35, 2, 250, 65, 3, 2, 2, 2,
	251, 252, 9, 4, 2, 2, 252, 67, 3, 2, 2, 2, 253, 254, 7, 35, 2, 2, 254,
	69, 3, 2, 2, 2, 255, 256, 7, 35, 2, 2, 256, 71, 3, 2, 2, 2, 27, 75, 84,
	94, 103, 113, 125, 132, 140, 142, 145, 152, 158, 165, 172, 178, 180, 184,
	197, 201, 205, 209, 218, 235, 239, 247,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'::='", "'{'", "'}'", "'['", "']'", "'('", "')'", "','", "'SEQUENCE'",
	"'INTEGER'", "'IMPLICIT'", "'EXPLICIT'", "'OPTIONAL'", "'DEFAULT'", "",
	"'ANY'", "'OF'", "'CHOICE'", "'SIZE'", "'MAX'", "'BMPString'", "'IA5String'",
	"'PrintableString'", "'TeletexString'", "'UniversalString'", "'UTF8String'",
	"", "", "", "'..'",
}
var symbolicNames = []string{
	"", "COMMENT", "ASSIGN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "LPAREN",
	"RPAREN", "COMMA", "SEQUENCE_LITERAL", "INTEGER_LITERAL", "IMPLICIT_LITERAL",
	"EXPLICIT_LITERAL", "OPTIONAL_LITERAL", "DEFAULT_LITERAL", "DEFINED_BY_LITERAL",
	"ANY_LITERAL", "OF_LITERAL", "CHOICE_LITERAL", "SIZE_LITERAL", "MAX", "BMPSTRING_LITERAL",
	"IA5STRING_LITERAL", "PRINTABLE_STRING_LITERAL", "TELETEX_STRING_LITERAL",
	"UNIVERSAL_STRING_LITERAL", "UTF8_STRING_LITERAL", "BIT_STRING_LITERAL",
	"OCTET_STRING_LITERAL", "OBJECT_IDENTIFIER_LITERAL", "DOTDOT", "INTEGER",
	"WORD", "WHITESPACE",
}

var ruleNames = []string{
	"module", "assignment", "primitive", "primitive_name", "tag", "context_flag",
	"encoding_flags", "default_value", "size_constraint", "type_constraints",
	"type_definition", "integer_integer", "integer_value", "integer_enum_value_list",
	"integer_enum_value", "integer_oid_value_list", "integer_oid_value", "oid",
	"oid_assignment", "octet_string", "bit_string", "string_string", "sequence",
	"sequence_list", "sequence_of", "choice", "any", "field_list", "field_definition",
	"string_literal", "type_name", "integer_oid_reference", "value_identifier",
	"field_name", "custom_type_name",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ASN1SchemaParser struct {
	*antlr.BaseParser
}

func NewASN1SchemaParser(input antlr.TokenStream) *ASN1SchemaParser {
	this := new(ASN1SchemaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ASN1Schema.g4"

	return this
}

// ASN1SchemaParser tokens.
const (
	ASN1SchemaParserEOF                       = antlr.TokenEOF
	ASN1SchemaParserCOMMENT                   = 1
	ASN1SchemaParserASSIGN                    = 2
	ASN1SchemaParserLBRACE                    = 3
	ASN1SchemaParserRBRACE                    = 4
	ASN1SchemaParserLBRACKET                  = 5
	ASN1SchemaParserRBRACKET                  = 6
	ASN1SchemaParserLPAREN                    = 7
	ASN1SchemaParserRPAREN                    = 8
	ASN1SchemaParserCOMMA                     = 9
	ASN1SchemaParserSEQUENCE_LITERAL          = 10
	ASN1SchemaParserINTEGER_LITERAL           = 11
	ASN1SchemaParserIMPLICIT_LITERAL          = 12
	ASN1SchemaParserEXPLICIT_LITERAL          = 13
	ASN1SchemaParserOPTIONAL_LITERAL          = 14
	ASN1SchemaParserDEFAULT_LITERAL           = 15
	ASN1SchemaParserDEFINED_BY_LITERAL        = 16
	ASN1SchemaParserANY_LITERAL               = 17
	ASN1SchemaParserOF_LITERAL                = 18
	ASN1SchemaParserCHOICE_LITERAL            = 19
	ASN1SchemaParserSIZE_LITERAL              = 20
	ASN1SchemaParserMAX                       = 21
	ASN1SchemaParserBMPSTRING_LITERAL         = 22
	ASN1SchemaParserIA5STRING_LITERAL         = 23
	ASN1SchemaParserPRINTABLE_STRING_LITERAL  = 24
	ASN1SchemaParserTELETEX_STRING_LITERAL    = 25
	ASN1SchemaParserUNIVERSAL_STRING_LITERAL  = 26
	ASN1SchemaParserUTF8_STRING_LITERAL       = 27
	ASN1SchemaParserBIT_STRING_LITERAL        = 28
	ASN1SchemaParserOCTET_STRING_LITERAL      = 29
	ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL = 30
	ASN1SchemaParserDOTDOT                    = 31
	ASN1SchemaParserINTEGER                   = 32
	ASN1SchemaParserWORD                      = 33
	ASN1SchemaParserWHITESPACE                = 34
)

// ASN1SchemaParser rules.
const (
	ASN1SchemaParserRULE_module                  = 0
	ASN1SchemaParserRULE_assignment              = 1
	ASN1SchemaParserRULE_primitive               = 2
	ASN1SchemaParserRULE_primitive_name          = 3
	ASN1SchemaParserRULE_tag                     = 4
	ASN1SchemaParserRULE_context_flag            = 5
	ASN1SchemaParserRULE_encoding_flags          = 6
	ASN1SchemaParserRULE_default_value           = 7
	ASN1SchemaParserRULE_size_constraint         = 8
	ASN1SchemaParserRULE_type_constraints        = 9
	ASN1SchemaParserRULE_type_definition         = 10
	ASN1SchemaParserRULE_integer_integer         = 11
	ASN1SchemaParserRULE_integer_value           = 12
	ASN1SchemaParserRULE_integer_enum_value_list = 13
	ASN1SchemaParserRULE_integer_enum_value      = 14
	ASN1SchemaParserRULE_integer_oid_value_list  = 15
	ASN1SchemaParserRULE_integer_oid_value       = 16
	ASN1SchemaParserRULE_oid                     = 17
	ASN1SchemaParserRULE_oid_assignment          = 18
	ASN1SchemaParserRULE_octet_string            = 19
	ASN1SchemaParserRULE_bit_string              = 20
	ASN1SchemaParserRULE_string_string           = 21
	ASN1SchemaParserRULE_sequence                = 22
	ASN1SchemaParserRULE_sequence_list           = 23
	ASN1SchemaParserRULE_sequence_of             = 24
	ASN1SchemaParserRULE_choice                  = 25
	ASN1SchemaParserRULE_any                     = 26
	ASN1SchemaParserRULE_field_list              = 27
	ASN1SchemaParserRULE_field_definition        = 28
	ASN1SchemaParserRULE_string_literal          = 29
	ASN1SchemaParserRULE_type_name               = 30
	ASN1SchemaParserRULE_integer_oid_reference   = 31
	ASN1SchemaParserRULE_value_identifier        = 32
	ASN1SchemaParserRULE_field_name              = 33
	ASN1SchemaParserRULE_custom_type_name        = 34
)

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) EOF() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserEOF, 0)
}

func (s *ModuleContext) AllAssignment() []IAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignmentContext)(nil)).Elem())
	var tst = make([]IAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignmentContext)
		}
	}

	return tst
}

func (s *ModuleContext) Assignment(i int) IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *ASN1SchemaParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ASN1SchemaParserRULE_module)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(ASN1SchemaParserSEQUENCE_LITERAL-10))|(1<<(ASN1SchemaParserINTEGER_LITERAL-10))|(1<<(ASN1SchemaParserANY_LITERAL-10))|(1<<(ASN1SchemaParserBMPSTRING_LITERAL-10))|(1<<(ASN1SchemaParserIA5STRING_LITERAL-10))|(1<<(ASN1SchemaParserPRINTABLE_STRING_LITERAL-10))|(1<<(ASN1SchemaParserTELETEX_STRING_LITERAL-10))|(1<<(ASN1SchemaParserUNIVERSAL_STRING_LITERAL-10))|(1<<(ASN1SchemaParserUTF8_STRING_LITERAL-10))|(1<<(ASN1SchemaParserBIT_STRING_LITERAL-10))|(1<<(ASN1SchemaParserOCTET_STRING_LITERAL-10))|(1<<(ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL-10))|(1<<(ASN1SchemaParserWORD-10)))) != 0) {
		{
			p.SetState(70)
			p.Assignment()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(ASN1SchemaParserEOF)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Type_name() IType_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserASSIGN, 0)
}

func (s *AssignmentContext) Primitive() IPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *AssignmentContext) Oid_assignment() IOid_assignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOid_assignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOid_assignmentContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *ASN1SchemaParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ASN1SchemaParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Type_name()
		}
		{
			p.SetState(78)
			p.Match(ASN1SchemaParserASSIGN)
		}
		{
			p.SetState(79)
			p.Primitive()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Oid_assignment()
		}

	}

	return localctx
}

// IPrimitiveContext is an interface to support dynamic dispatch.
type IPrimitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveContext differentiates from other interfaces.
	IsPrimitiveContext()
}

type PrimitiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveContext() *PrimitiveContext {
	var p = new(PrimitiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_primitive
	return p
}

func (*PrimitiveContext) IsPrimitiveContext() {}

func NewPrimitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveContext {
	var p = new(PrimitiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_primitive

	return p
}

func (s *PrimitiveContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *PrimitiveContext) Choice() IChoiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChoiceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChoiceContext)
}

func (s *PrimitiveContext) Integer_integer() IInteger_integerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_integerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_integerContext)
}

func (s *PrimitiveContext) Oid() IOidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOidContext)
}

func (s *PrimitiveContext) Octet_string() IOctet_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOctet_stringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOctet_stringContext)
}

func (s *PrimitiveContext) Bit_string() IBit_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBit_stringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBit_stringContext)
}

func (s *PrimitiveContext) String_string() IString_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_stringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_stringContext)
}

func (s *PrimitiveContext) Any() IAnyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnyContext)
}

func (s *PrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterPrimitive(s)
	}
}

func (s *PrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitPrimitive(s)
	}
}

func (p *ASN1SchemaParser) Primitive() (localctx IPrimitiveContext) {
	localctx = NewPrimitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ASN1SchemaParserRULE_primitive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserSEQUENCE_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Sequence()
		}

	case ASN1SchemaParserCHOICE_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Choice()
		}

	case ASN1SchemaParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Integer_integer()
		}

	case ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.Oid()
		}

	case ASN1SchemaParserOCTET_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(88)
			p.Octet_string()
		}

	case ASN1SchemaParserBIT_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(89)
			p.Bit_string()
		}

	case ASN1SchemaParserBMPSTRING_LITERAL, ASN1SchemaParserIA5STRING_LITERAL, ASN1SchemaParserPRINTABLE_STRING_LITERAL, ASN1SchemaParserTELETEX_STRING_LITERAL, ASN1SchemaParserUNIVERSAL_STRING_LITERAL, ASN1SchemaParserUTF8_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(90)
			p.String_string()
		}

	case ASN1SchemaParserANY_LITERAL:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(91)
			p.Any()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimitive_nameContext is an interface to support dynamic dispatch.
type IPrimitive_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitive_nameContext differentiates from other interfaces.
	IsPrimitive_nameContext()
}

type Primitive_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitive_nameContext() *Primitive_nameContext {
	var p = new(Primitive_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_primitive_name
	return p
}

func (*Primitive_nameContext) IsPrimitive_nameContext() {}

func NewPrimitive_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primitive_nameContext {
	var p = new(Primitive_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_primitive_name

	return p
}

func (s *Primitive_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Primitive_nameContext) SEQUENCE_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserSEQUENCE_LITERAL, 0)
}

func (s *Primitive_nameContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserINTEGER_LITERAL, 0)
}

func (s *Primitive_nameContext) OBJECT_IDENTIFIER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL, 0)
}

func (s *Primitive_nameContext) OCTET_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOCTET_STRING_LITERAL, 0)
}

func (s *Primitive_nameContext) BIT_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserBIT_STRING_LITERAL, 0)
}

func (s *Primitive_nameContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Primitive_nameContext) ANY_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserANY_LITERAL, 0)
}

func (s *Primitive_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitive_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primitive_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterPrimitive_name(s)
	}
}

func (s *Primitive_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitPrimitive_name(s)
	}
}

func (p *ASN1SchemaParser) Primitive_name() (localctx IPrimitive_nameContext) {
	localctx = NewPrimitive_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ASN1SchemaParserRULE_primitive_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserSEQUENCE_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(ASN1SchemaParserSEQUENCE_LITERAL)
		}

	case ASN1SchemaParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Match(ASN1SchemaParserINTEGER_LITERAL)
		}

	case ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Match(ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL)
		}

	case ASN1SchemaParserOCTET_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(97)
			p.Match(ASN1SchemaParserOCTET_STRING_LITERAL)
		}

	case ASN1SchemaParserBIT_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.Match(ASN1SchemaParserBIT_STRING_LITERAL)
		}

	case ASN1SchemaParserBMPSTRING_LITERAL, ASN1SchemaParserIA5STRING_LITERAL, ASN1SchemaParserPRINTABLE_STRING_LITERAL, ASN1SchemaParserTELETEX_STRING_LITERAL, ASN1SchemaParserUNIVERSAL_STRING_LITERAL, ASN1SchemaParserUTF8_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(99)
			p.String_literal()
		}

	case ASN1SchemaParserANY_LITERAL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(100)
			p.Match(ASN1SchemaParserANY_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserLBRACKET, 0)
}

func (s *TagContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserINTEGER, 0)
}

func (s *TagContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserRBRACKET, 0)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *ASN1SchemaParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ASN1SchemaParserRULE_tag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(ASN1SchemaParserLBRACKET)
	}
	{
		p.SetState(104)
		p.Match(ASN1SchemaParserINTEGER)
	}
	{
		p.SetState(105)
		p.Match(ASN1SchemaParserRBRACKET)
	}

	return localctx
}

// IContext_flagContext is an interface to support dynamic dispatch.
type IContext_flagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContext_flagContext differentiates from other interfaces.
	IsContext_flagContext()
}

type Context_flagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContext_flagContext() *Context_flagContext {
	var p = new(Context_flagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_context_flag
	return p
}

func (*Context_flagContext) IsContext_flagContext() {}

func NewContext_flagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Context_flagContext {
	var p = new(Context_flagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_context_flag

	return p
}

func (s *Context_flagContext) GetParser() antlr.Parser { return s.parser }

func (s *Context_flagContext) IMPLICIT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserIMPLICIT_LITERAL, 0)
}

func (s *Context_flagContext) EXPLICIT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserEXPLICIT_LITERAL, 0)
}

func (s *Context_flagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Context_flagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Context_flagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterContext_flag(s)
	}
}

func (s *Context_flagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitContext_flag(s)
	}
}

func (p *ASN1SchemaParser) Context_flag() (localctx IContext_flagContext) {
	localctx = NewContext_flagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ASN1SchemaParserRULE_context_flag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ASN1SchemaParserIMPLICIT_LITERAL || _la == ASN1SchemaParserEXPLICIT_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEncoding_flagsContext is an interface to support dynamic dispatch.
type IEncoding_flagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncoding_flagsContext differentiates from other interfaces.
	IsEncoding_flagsContext()
}

type Encoding_flagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncoding_flagsContext() *Encoding_flagsContext {
	var p = new(Encoding_flagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_encoding_flags
	return p
}

func (*Encoding_flagsContext) IsEncoding_flagsContext() {}

func NewEncoding_flagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Encoding_flagsContext {
	var p = new(Encoding_flagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_encoding_flags

	return p
}

func (s *Encoding_flagsContext) GetParser() antlr.Parser { return s.parser }

func (s *Encoding_flagsContext) OPTIONAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOPTIONAL_LITERAL, 0)
}

func (s *Encoding_flagsContext) Default_value() IDefault_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefault_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefault_valueContext)
}

func (s *Encoding_flagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Encoding_flagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Encoding_flagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterEncoding_flags(s)
	}
}

func (s *Encoding_flagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitEncoding_flags(s)
	}
}

func (p *ASN1SchemaParser) Encoding_flags() (localctx IEncoding_flagsContext) {
	localctx = NewEncoding_flagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ASN1SchemaParserRULE_encoding_flags)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserOPTIONAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Match(ASN1SchemaParserOPTIONAL_LITERAL)
		}

	case ASN1SchemaParserDEFAULT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Default_value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefault_valueContext is an interface to support dynamic dispatch.
type IDefault_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefault_valueContext differentiates from other interfaces.
	IsDefault_valueContext()
}

type Default_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_valueContext() *Default_valueContext {
	var p = new(Default_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_default_value
	return p
}

func (*Default_valueContext) IsDefault_valueContext() {}

func NewDefault_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_valueContext {
	var p = new(Default_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_default_value

	return p
}

func (s *Default_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_valueContext) DEFAULT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserDEFAULT_LITERAL, 0)
}

func (s *Default_valueContext) Value_identifier() IValue_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_identifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValue_identifierContext)
}

func (s *Default_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Default_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterDefault_value(s)
	}
}

func (s *Default_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitDefault_value(s)
	}
}

func (p *ASN1SchemaParser) Default_value() (localctx IDefault_valueContext) {
	localctx = NewDefault_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ASN1SchemaParserRULE_default_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(ASN1SchemaParserDEFAULT_LITERAL)
	}
	{
		p.SetState(114)
		p.Value_identifier()
	}

	return localctx
}

// ISize_constraintContext is an interface to support dynamic dispatch.
type ISize_constraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSize_constraintContext differentiates from other interfaces.
	IsSize_constraintContext()
}

type Size_constraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySize_constraintContext() *Size_constraintContext {
	var p = new(Size_constraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_size_constraint
	return p
}

func (*Size_constraintContext) IsSize_constraintContext() {}

func NewSize_constraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Size_constraintContext {
	var p = new(Size_constraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_size_constraint

	return p
}

func (s *Size_constraintContext) GetParser() antlr.Parser { return s.parser }

func (s *Size_constraintContext) SIZE_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserSIZE_LITERAL, 0)
}

func (s *Size_constraintContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserLPAREN, 0)
}

func (s *Size_constraintContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(ASN1SchemaParserINTEGER)
}

func (s *Size_constraintContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserINTEGER, i)
}

func (s *Size_constraintContext) DOTDOT() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserDOTDOT, 0)
}

func (s *Size_constraintContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserRPAREN, 0)
}

func (s *Size_constraintContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Size_constraintContext) MAX() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserMAX, 0)
}

func (s *Size_constraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Size_constraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Size_constraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterSize_constraint(s)
	}
}

func (s *Size_constraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitSize_constraint(s)
	}
}

func (p *ASN1SchemaParser) Size_constraint() (localctx ISize_constraintContext) {
	localctx = NewSize_constraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ASN1SchemaParserRULE_size_constraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(ASN1SchemaParserSIZE_LITERAL)
	}
	{
		p.SetState(117)
		p.Match(ASN1SchemaParserLPAREN)
	}
	{
		p.SetState(118)
		p.Match(ASN1SchemaParserINTEGER)
	}
	{
		p.SetState(119)
		p.Match(ASN1SchemaParserDOTDOT)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserINTEGER:
		{
			p.SetState(120)
			p.Match(ASN1SchemaParserINTEGER)
		}

	case ASN1SchemaParserWORD:
		{
			p.SetState(121)
			p.Field_name()
		}

	case ASN1SchemaParserMAX:
		{
			p.SetState(122)
			p.Match(ASN1SchemaParserMAX)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(125)
		p.Match(ASN1SchemaParserRPAREN)
	}

	return localctx
}

// IType_constraintsContext is an interface to support dynamic dispatch.
type IType_constraintsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_constraintsContext differentiates from other interfaces.
	IsType_constraintsContext()
}

type Type_constraintsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_constraintsContext() *Type_constraintsContext {
	var p = new(Type_constraintsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_type_constraints
	return p
}

func (*Type_constraintsContext) IsType_constraintsContext() {}

func NewType_constraintsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_constraintsContext {
	var p = new(Type_constraintsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_type_constraints

	return p
}

func (s *Type_constraintsContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_constraintsContext) Size_constraint() ISize_constraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISize_constraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISize_constraintContext)
}

func (s *Type_constraintsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_constraintsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_constraintsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterType_constraints(s)
	}
}

func (s *Type_constraintsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitType_constraints(s)
	}
}

func (p *ASN1SchemaParser) Type_constraints() (localctx IType_constraintsContext) {
	localctx = NewType_constraintsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ASN1SchemaParserRULE_type_constraints)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Size_constraint()
	}

	return localctx
}

// IType_definitionContext is an interface to support dynamic dispatch.
type IType_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_definitionContext differentiates from other interfaces.
	IsType_definitionContext()
}

type Type_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_definitionContext() *Type_definitionContext {
	var p = new(Type_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_type_definition
	return p
}

func (*Type_definitionContext) IsType_definitionContext() {}

func NewType_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_definitionContext {
	var p = new(Type_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_type_definition

	return p
}

func (s *Type_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_definitionContext) Primitive() IPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *Type_definitionContext) Context_flag() IContext_flagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContext_flagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContext_flagContext)
}

func (s *Type_definitionContext) Encoding_flags() IEncoding_flagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncoding_flagsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEncoding_flagsContext)
}

func (s *Type_definitionContext) Type_name() IType_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
}

func (s *Type_definitionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserLPAREN, 0)
}

func (s *Type_definitionContext) Type_constraints() IType_constraintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_constraintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_constraintsContext)
}

func (s *Type_definitionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserRPAREN, 0)
}

func (s *Type_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterType_definition(s)
	}
}

func (s *Type_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitType_definition(s)
	}
}

func (p *ASN1SchemaParser) Type_definition() (localctx IType_definitionContext) {
	localctx = NewType_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ASN1SchemaParserRULE_type_definition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserIMPLICIT_LITERAL || _la == ASN1SchemaParserEXPLICIT_LITERAL {
		{
			p.SetState(129)
			p.Context_flag()
		}

	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(132)
			p.Primitive()
		}

	case 2:
		{
			p.SetState(133)
			p.Type_name()
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ASN1SchemaParserLPAREN {
			{
				p.SetState(134)
				p.Match(ASN1SchemaParserLPAREN)
			}
			{
				p.SetState(135)
				p.Type_constraints()
			}
			{
				p.SetState(136)
				p.Match(ASN1SchemaParserRPAREN)
			}

		}

	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserOPTIONAL_LITERAL || _la == ASN1SchemaParserDEFAULT_LITERAL {
		{
			p.SetState(142)
			p.Encoding_flags()
		}

	}

	return localctx
}

// IInteger_integerContext is an interface to support dynamic dispatch.
type IInteger_integerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_integerContext differentiates from other interfaces.
	IsInteger_integerContext()
}

type Integer_integerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_integerContext() *Integer_integerContext {
	var p = new(Integer_integerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_integer_integer
	return p
}

func (*Integer_integerContext) IsInteger_integerContext() {}

func NewInteger_integerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_integerContext {
	var p = new(Integer_integerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_integer_integer

	return p
}

func (s *Integer_integerContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_integerContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserINTEGER_LITERAL, 0)
}

func (s *Integer_integerContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserLBRACE, 0)
}

func (s *Integer_integerContext) Integer_enum_value_list() IInteger_enum_value_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_enum_value_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_enum_value_listContext)
}

func (s *Integer_integerContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserRBRACE, 0)
}

func (s *Integer_integerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_integerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_integerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterInteger_integer(s)
	}
}

func (s *Integer_integerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitInteger_integer(s)
	}
}

func (p *ASN1SchemaParser) Integer_integer() (localctx IInteger_integerContext) {
	localctx = NewInteger_integerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ASN1SchemaParserRULE_integer_integer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(ASN1SchemaParserINTEGER_LITERAL)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserLBRACE {
		{
			p.SetState(146)
			p.Match(ASN1SchemaParserLBRACE)
		}
		{
			p.SetState(147)
			p.Integer_enum_value_list()
		}
		{
			p.SetState(148)
			p.Match(ASN1SchemaParserRBRACE)
		}

	}

	return localctx
}

// IInteger_valueContext is an interface to support dynamic dispatch.
type IInteger_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_valueContext differentiates from other interfaces.
	IsInteger_valueContext()
}

type Integer_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_valueContext() *Integer_valueContext {
	var p = new(Integer_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_integer_value
	return p
}

func (*Integer_valueContext) IsInteger_valueContext() {}

func NewInteger_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_valueContext {
	var p = new(Integer_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_integer_value

	return p
}

func (s *Integer_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_valueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserLPAREN, 0)
}

func (s *Integer_valueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserRPAREN, 0)
}

func (s *Integer_valueContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(ASN1SchemaParserINTEGER)
}

func (s *Integer_valueContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserINTEGER, i)
}

func (s *Integer_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterInteger_value(s)
	}
}

func (s *Integer_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitInteger_value(s)
	}
}

func (p *ASN1SchemaParser) Integer_value() (localctx IInteger_valueContext) {
	localctx = NewInteger_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ASN1SchemaParserRULE_integer_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(ASN1SchemaParserLPAREN)
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ASN1SchemaParserINTEGER {
		{
			p.SetState(153)
			p.Match(ASN1SchemaParserINTEGER)
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(158)
		p.Match(ASN1SchemaParserRPAREN)
	}

	return localctx
}

// IInteger_enum_value_listContext is an interface to support dynamic dispatch.
type IInteger_enum_value_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_enum_value_listContext differentiates from other interfaces.
	IsInteger_enum_value_listContext()
}

type Integer_enum_value_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_enum_value_listContext() *Integer_enum_value_listContext {
	var p = new(Integer_enum_value_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_integer_enum_value_list
	return p
}

func (*Integer_enum_value_listContext) IsInteger_enum_value_listContext() {}

func NewInteger_enum_value_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_enum_value_listContext {
	var p = new(Integer_enum_value_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_integer_enum_value_list

	return p
}

func (s *Integer_enum_value_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_enum_value_listContext) Integer_enum_value() IInteger_enum_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_enum_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_enum_valueContext)
}

func (s *Integer_enum_value_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserCOMMA, 0)
}

func (s *Integer_enum_value_listContext) Integer_enum_value_list() IInteger_enum_value_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_enum_value_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_enum_value_listContext)
}

func (s *Integer_enum_value_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_enum_value_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_enum_value_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterInteger_enum_value_list(s)
	}
}

func (s *Integer_enum_value_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitInteger_enum_value_list(s)
	}
}

func (p *ASN1SchemaParser) Integer_enum_value_list() (localctx IInteger_enum_value_listContext) {
	localctx = NewInteger_enum_value_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ASN1SchemaParserRULE_integer_enum_value_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Integer_enum_value()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserCOMMA {
		{
			p.SetState(161)
			p.Match(ASN1SchemaParserCOMMA)
		}
		{
			p.SetState(162)
			p.Integer_enum_value_list()
		}

	}

	return localctx
}

// IInteger_enum_valueContext is an interface to support dynamic dispatch.
type IInteger_enum_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_enum_valueContext differentiates from other interfaces.
	IsInteger_enum_valueContext()
}

type Integer_enum_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_enum_valueContext() *Integer_enum_valueContext {
	var p = new(Integer_enum_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_integer_enum_value
	return p
}

func (*Integer_enum_valueContext) IsInteger_enum_valueContext() {}

func NewInteger_enum_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_enum_valueContext {
	var p = new(Integer_enum_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_integer_enum_value

	return p
}

func (s *Integer_enum_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_enum_valueContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Integer_enum_valueContext) Integer_value() IInteger_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_valueContext)
}

func (s *Integer_enum_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_enum_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_enum_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterInteger_enum_value(s)
	}
}

func (s *Integer_enum_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitInteger_enum_value(s)
	}
}

func (p *ASN1SchemaParser) Integer_enum_value() (localctx IInteger_enum_valueContext) {
	localctx = NewInteger_enum_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ASN1SchemaParserRULE_integer_enum_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Field_name()
	}
	{
		p.SetState(166)
		p.Integer_value()
	}

	return localctx
}

// IInteger_oid_value_listContext is an interface to support dynamic dispatch.
type IInteger_oid_value_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_oid_value_listContext differentiates from other interfaces.
	IsInteger_oid_value_listContext()
}

type Integer_oid_value_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_oid_value_listContext() *Integer_oid_value_listContext {
	var p = new(Integer_oid_value_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_integer_oid_value_list
	return p
}

func (*Integer_oid_value_listContext) IsInteger_oid_value_listContext() {}

func NewInteger_oid_value_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_oid_value_listContext {
	var p = new(Integer_oid_value_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_integer_oid_value_list

	return p
}

func (s *Integer_oid_value_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_oid_value_listContext) Integer_oid_value() IInteger_oid_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_oid_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_oid_valueContext)
}

func (s *Integer_oid_value_listContext) Integer_oid_value_list() IInteger_oid_value_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_oid_value_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_oid_value_listContext)
}

func (s *Integer_oid_value_listContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(ASN1SchemaParserINTEGER)
}

func (s *Integer_oid_value_listContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserINTEGER, i)
}

func (s *Integer_oid_value_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_oid_value_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_oid_value_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterInteger_oid_value_list(s)
	}
}

func (s *Integer_oid_value_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitInteger_oid_value_list(s)
	}
}

func (p *ASN1SchemaParser) Integer_oid_value_list() (localctx IInteger_oid_value_listContext) {
	localctx = NewInteger_oid_value_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ASN1SchemaParserRULE_integer_oid_value_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Integer_oid_value()
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ASN1SchemaParserWORD {
			{
				p.SetState(169)
				p.Integer_oid_value_list()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Integer_oid_value()
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ASN1SchemaParserINTEGER {
			{
				p.SetState(173)
				p.Match(ASN1SchemaParserINTEGER)
			}

			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IInteger_oid_valueContext is an interface to support dynamic dispatch.
type IInteger_oid_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_oid_valueContext differentiates from other interfaces.
	IsInteger_oid_valueContext()
}

type Integer_oid_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_oid_valueContext() *Integer_oid_valueContext {
	var p = new(Integer_oid_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_integer_oid_value
	return p
}

func (*Integer_oid_valueContext) IsInteger_oid_valueContext() {}

func NewInteger_oid_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_oid_valueContext {
	var p = new(Integer_oid_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_integer_oid_value

	return p
}

func (s *Integer_oid_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_oid_valueContext) Integer_oid_reference() IInteger_oid_referenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_oid_referenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_oid_referenceContext)
}

func (s *Integer_oid_valueContext) Integer_enum_value() IInteger_enum_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_enum_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_enum_valueContext)
}

func (s *Integer_oid_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_oid_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_oid_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterInteger_oid_value(s)
	}
}

func (s *Integer_oid_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitInteger_oid_value(s)
	}
}

func (p *ASN1SchemaParser) Integer_oid_value() (localctx IInteger_oid_valueContext) {
	localctx = NewInteger_oid_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ASN1SchemaParserRULE_integer_oid_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Integer_oid_reference()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Integer_enum_value()
		}

	}

	return localctx
}

// IOidContext is an interface to support dynamic dispatch.
type IOidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOidContext differentiates from other interfaces.
	IsOidContext()
}

type OidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOidContext() *OidContext {
	var p = new(OidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_oid
	return p
}

func (*OidContext) IsOidContext() {}

func NewOidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OidContext {
	var p = new(OidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_oid

	return p
}

func (s *OidContext) GetParser() antlr.Parser { return s.parser }

func (s *OidContext) OBJECT_IDENTIFIER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL, 0)
}

func (s *OidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterOid(s)
	}
}

func (s *OidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitOid(s)
	}
}

func (p *ASN1SchemaParser) Oid() (localctx IOidContext) {
	localctx = NewOidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ASN1SchemaParserRULE_oid)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL)
	}

	return localctx
}

// IOid_assignmentContext is an interface to support dynamic dispatch.
type IOid_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOid_assignmentContext differentiates from other interfaces.
	IsOid_assignmentContext()
}

type Oid_assignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOid_assignmentContext() *Oid_assignmentContext {
	var p = new(Oid_assignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_oid_assignment
	return p
}

func (*Oid_assignmentContext) IsOid_assignmentContext() {}

func NewOid_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Oid_assignmentContext {
	var p = new(Oid_assignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_oid_assignment

	return p
}

func (s *Oid_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Oid_assignmentContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Oid_assignmentContext) OBJECT_IDENTIFIER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL, 0)
}

func (s *Oid_assignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserASSIGN, 0)
}

func (s *Oid_assignmentContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserLBRACE, 0)
}

func (s *Oid_assignmentContext) Integer_oid_value_list() IInteger_oid_value_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_oid_value_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_oid_value_listContext)
}

func (s *Oid_assignmentContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserRBRACE, 0)
}

func (s *Oid_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Oid_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Oid_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterOid_assignment(s)
	}
}

func (s *Oid_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitOid_assignment(s)
	}
}

func (p *ASN1SchemaParser) Oid_assignment() (localctx IOid_assignmentContext) {
	localctx = NewOid_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ASN1SchemaParserRULE_oid_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Field_name()
	}
	{
		p.SetState(187)
		p.Match(ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL)
	}
	{
		p.SetState(188)
		p.Match(ASN1SchemaParserASSIGN)
	}
	{
		p.SetState(189)
		p.Match(ASN1SchemaParserLBRACE)
	}
	{
		p.SetState(190)
		p.Integer_oid_value_list()
	}
	{
		p.SetState(191)
		p.Match(ASN1SchemaParserRBRACE)
	}

	return localctx
}

// IOctet_stringContext is an interface to support dynamic dispatch.
type IOctet_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOctet_stringContext differentiates from other interfaces.
	IsOctet_stringContext()
}

type Octet_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOctet_stringContext() *Octet_stringContext {
	var p = new(Octet_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_octet_string
	return p
}

func (*Octet_stringContext) IsOctet_stringContext() {}

func NewOctet_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Octet_stringContext {
	var p = new(Octet_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_octet_string

	return p
}

func (s *Octet_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Octet_stringContext) OCTET_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOCTET_STRING_LITERAL, 0)
}

func (s *Octet_stringContext) Type_constraints() IType_constraintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_constraintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_constraintsContext)
}

func (s *Octet_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Octet_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Octet_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterOctet_string(s)
	}
}

func (s *Octet_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitOctet_string(s)
	}
}

func (p *ASN1SchemaParser) Octet_string() (localctx IOctet_stringContext) {
	localctx = NewOctet_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ASN1SchemaParserRULE_octet_string)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(ASN1SchemaParserOCTET_STRING_LITERAL)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserSIZE_LITERAL {
		{
			p.SetState(194)
			p.Type_constraints()
		}

	}

	return localctx
}

// IBit_stringContext is an interface to support dynamic dispatch.
type IBit_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBit_stringContext differentiates from other interfaces.
	IsBit_stringContext()
}

type Bit_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBit_stringContext() *Bit_stringContext {
	var p = new(Bit_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_bit_string
	return p
}

func (*Bit_stringContext) IsBit_stringContext() {}

func NewBit_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bit_stringContext {
	var p = new(Bit_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_bit_string

	return p
}

func (s *Bit_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Bit_stringContext) BIT_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserBIT_STRING_LITERAL, 0)
}

func (s *Bit_stringContext) Type_constraints() IType_constraintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_constraintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_constraintsContext)
}

func (s *Bit_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bit_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bit_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterBit_string(s)
	}
}

func (s *Bit_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitBit_string(s)
	}
}

func (p *ASN1SchemaParser) Bit_string() (localctx IBit_stringContext) {
	localctx = NewBit_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ASN1SchemaParserRULE_bit_string)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(ASN1SchemaParserBIT_STRING_LITERAL)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserSIZE_LITERAL {
		{
			p.SetState(198)
			p.Type_constraints()
		}

	}

	return localctx
}

// IString_stringContext is an interface to support dynamic dispatch.
type IString_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_stringContext differentiates from other interfaces.
	IsString_stringContext()
}

type String_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_stringContext() *String_stringContext {
	var p = new(String_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_string_string
	return p
}

func (*String_stringContext) IsString_stringContext() {}

func NewString_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_stringContext {
	var p = new(String_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_string_string

	return p
}

func (s *String_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *String_stringContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *String_stringContext) Type_constraints() IType_constraintsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_constraintsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_constraintsContext)
}

func (s *String_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterString_string(s)
	}
}

func (s *String_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitString_string(s)
	}
}

func (p *ASN1SchemaParser) String_string() (localctx IString_stringContext) {
	localctx = NewString_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ASN1SchemaParserRULE_string_string)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.String_literal()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserSIZE_LITERAL {
		{
			p.SetState(202)
			p.Type_constraints()
		}

	}

	return localctx
}

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_sequence
	return p
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) Sequence_list() ISequence_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_listContext)
}

func (s *SequenceContext) Sequence_of() ISequence_ofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_ofContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_ofContext)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitSequence(s)
	}
}

func (p *ASN1SchemaParser) Sequence() (localctx ISequenceContext) {
	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ASN1SchemaParserRULE_sequence)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Sequence_list()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Sequence_of()
		}

	}

	return localctx
}

// ISequence_listContext is an interface to support dynamic dispatch.
type ISequence_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequence_listContext differentiates from other interfaces.
	IsSequence_listContext()
}

type Sequence_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequence_listContext() *Sequence_listContext {
	var p = new(Sequence_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_sequence_list
	return p
}

func (*Sequence_listContext) IsSequence_listContext() {}

func NewSequence_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sequence_listContext {
	var p = new(Sequence_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_sequence_list

	return p
}

func (s *Sequence_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Sequence_listContext) SEQUENCE_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserSEQUENCE_LITERAL, 0)
}

func (s *Sequence_listContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserLBRACE, 0)
}

func (s *Sequence_listContext) Field_list() IField_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_listContext)
}

func (s *Sequence_listContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserRBRACE, 0)
}

func (s *Sequence_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sequence_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sequence_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterSequence_list(s)
	}
}

func (s *Sequence_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitSequence_list(s)
	}
}

func (p *ASN1SchemaParser) Sequence_list() (localctx ISequence_listContext) {
	localctx = NewSequence_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ASN1SchemaParserRULE_sequence_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(ASN1SchemaParserSEQUENCE_LITERAL)
	}
	{
		p.SetState(210)
		p.Match(ASN1SchemaParserLBRACE)
	}
	{
		p.SetState(211)
		p.Field_list()
	}
	{
		p.SetState(212)
		p.Match(ASN1SchemaParserRBRACE)
	}

	return localctx
}

// ISequence_ofContext is an interface to support dynamic dispatch.
type ISequence_ofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequence_ofContext differentiates from other interfaces.
	IsSequence_ofContext()
}

type Sequence_ofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequence_ofContext() *Sequence_ofContext {
	var p = new(Sequence_ofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_sequence_of
	return p
}

func (*Sequence_ofContext) IsSequence_ofContext() {}

func NewSequence_ofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sequence_ofContext {
	var p = new(Sequence_ofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_sequence_of

	return p
}

func (s *Sequence_ofContext) GetParser() antlr.Parser { return s.parser }

func (s *Sequence_ofContext) SEQUENCE_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserSEQUENCE_LITERAL, 0)
}

func (s *Sequence_ofContext) OF_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOF_LITERAL, 0)
}

func (s *Sequence_ofContext) Type_name() IType_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
}

func (s *Sequence_ofContext) Size_constraint() ISize_constraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISize_constraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISize_constraintContext)
}

func (s *Sequence_ofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sequence_ofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sequence_ofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterSequence_of(s)
	}
}

func (s *Sequence_ofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitSequence_of(s)
	}
}

func (p *ASN1SchemaParser) Sequence_of() (localctx ISequence_ofContext) {
	localctx = NewSequence_ofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ASN1SchemaParserRULE_sequence_of)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(ASN1SchemaParserSEQUENCE_LITERAL)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserSIZE_LITERAL {
		{
			p.SetState(215)
			p.Size_constraint()
		}

	}
	{
		p.SetState(218)
		p.Match(ASN1SchemaParserOF_LITERAL)
	}
	{
		p.SetState(219)
		p.Type_name()
	}

	return localctx
}

// IChoiceContext is an interface to support dynamic dispatch.
type IChoiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChoiceContext differentiates from other interfaces.
	IsChoiceContext()
}

type ChoiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChoiceContext() *ChoiceContext {
	var p = new(ChoiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_choice
	return p
}

func (*ChoiceContext) IsChoiceContext() {}

func NewChoiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceContext {
	var p = new(ChoiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_choice

	return p
}

func (s *ChoiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceContext) CHOICE_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserCHOICE_LITERAL, 0)
}

func (s *ChoiceContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserLBRACE, 0)
}

func (s *ChoiceContext) Field_list() IField_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_listContext)
}

func (s *ChoiceContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserRBRACE, 0)
}

func (s *ChoiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChoiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterChoice(s)
	}
}

func (s *ChoiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitChoice(s)
	}
}

func (p *ASN1SchemaParser) Choice() (localctx IChoiceContext) {
	localctx = NewChoiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ASN1SchemaParserRULE_choice)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(ASN1SchemaParserCHOICE_LITERAL)
	}
	{
		p.SetState(222)
		p.Match(ASN1SchemaParserLBRACE)
	}
	{
		p.SetState(223)
		p.Field_list()
	}
	{
		p.SetState(224)
		p.Match(ASN1SchemaParserRBRACE)
	}

	return localctx
}

// IAnyContext is an interface to support dynamic dispatch.
type IAnyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnyContext differentiates from other interfaces.
	IsAnyContext()
}

type AnyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyContext() *AnyContext {
	var p = new(AnyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_any
	return p
}

func (*AnyContext) IsAnyContext() {}

func NewAnyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyContext {
	var p = new(AnyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_any

	return p
}

func (s *AnyContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyContext) ANY_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserANY_LITERAL, 0)
}

func (s *AnyContext) DEFINED_BY_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserDEFINED_BY_LITERAL, 0)
}

func (s *AnyContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *AnyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterAny(s)
	}
}

func (s *AnyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitAny(s)
	}
}

func (p *ASN1SchemaParser) Any() (localctx IAnyContext) {
	localctx = NewAnyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ASN1SchemaParserRULE_any)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(ASN1SchemaParserANY_LITERAL)
	}
	{
		p.SetState(227)
		p.Match(ASN1SchemaParserDEFINED_BY_LITERAL)
	}
	{
		p.SetState(228)
		p.Field_name()
	}

	return localctx
}

// IField_listContext is an interface to support dynamic dispatch.
type IField_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_listContext differentiates from other interfaces.
	IsField_listContext()
}

type Field_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_listContext() *Field_listContext {
	var p = new(Field_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_field_list
	return p
}

func (*Field_listContext) IsField_listContext() {}

func NewField_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_listContext {
	var p = new(Field_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_field_list

	return p
}

func (s *Field_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_listContext) Field_definition() IField_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_definitionContext)
}

func (s *Field_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserCOMMA, 0)
}

func (s *Field_listContext) Field_list() IField_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_listContext)
}

func (s *Field_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterField_list(s)
	}
}

func (s *Field_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitField_list(s)
	}
}

func (p *ASN1SchemaParser) Field_list() (localctx IField_listContext) {
	localctx = NewField_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ASN1SchemaParserRULE_field_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Field_definition()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserCOMMA {
		{
			p.SetState(231)
			p.Match(ASN1SchemaParserCOMMA)
		}
		{
			p.SetState(232)
			p.Field_list()
		}

	}

	return localctx
}

// IField_definitionContext is an interface to support dynamic dispatch.
type IField_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_definitionContext differentiates from other interfaces.
	IsField_definitionContext()
}

type Field_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_definitionContext() *Field_definitionContext {
	var p = new(Field_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_field_definition
	return p
}

func (*Field_definitionContext) IsField_definitionContext() {}

func NewField_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_definitionContext {
	var p = new(Field_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_field_definition

	return p
}

func (s *Field_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_definitionContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Field_definitionContext) Type_definition() IType_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_definitionContext)
}

func (s *Field_definitionContext) Tag() ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *Field_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterField_definition(s)
	}
}

func (s *Field_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitField_definition(s)
	}
}

func (p *ASN1SchemaParser) Field_definition() (localctx IField_definitionContext) {
	localctx = NewField_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ASN1SchemaParserRULE_field_definition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Field_name()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserLBRACKET {
		{
			p.SetState(236)
			p.Tag()
		}

	}
	{
		p.SetState(239)
		p.Type_definition()
	}

	return localctx
}

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_literalContext differentiates from other interfaces.
	IsString_literalContext()
}

type String_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_literalContext() *String_literalContext {
	var p = new(String_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_string_literal
	return p
}

func (*String_literalContext) IsString_literalContext() {}

func NewString_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_literalContext {
	var p = new(String_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_string_literal

	return p
}

func (s *String_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *String_literalContext) BMPSTRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserBMPSTRING_LITERAL, 0)
}

func (s *String_literalContext) IA5STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserIA5STRING_LITERAL, 0)
}

func (s *String_literalContext) PRINTABLE_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserPRINTABLE_STRING_LITERAL, 0)
}

func (s *String_literalContext) TELETEX_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserTELETEX_STRING_LITERAL, 0)
}

func (s *String_literalContext) UTF8_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserUTF8_STRING_LITERAL, 0)
}

func (s *String_literalContext) UNIVERSAL_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserUNIVERSAL_STRING_LITERAL, 0)
}

func (s *String_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterString_literal(s)
	}
}

func (s *String_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitString_literal(s)
	}
}

func (p *ASN1SchemaParser) String_literal() (localctx IString_literalContext) {
	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ASN1SchemaParserRULE_string_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ASN1SchemaParserBMPSTRING_LITERAL)|(1<<ASN1SchemaParserIA5STRING_LITERAL)|(1<<ASN1SchemaParserPRINTABLE_STRING_LITERAL)|(1<<ASN1SchemaParserTELETEX_STRING_LITERAL)|(1<<ASN1SchemaParserUNIVERSAL_STRING_LITERAL)|(1<<ASN1SchemaParserUTF8_STRING_LITERAL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IType_nameContext is an interface to support dynamic dispatch.
type IType_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_nameContext differentiates from other interfaces.
	IsType_nameContext()
}

type Type_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_nameContext() *Type_nameContext {
	var p = new(Type_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_type_name
	return p
}

func (*Type_nameContext) IsType_nameContext() {}

func NewType_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_nameContext {
	var p = new(Type_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_type_name

	return p
}

func (s *Type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_nameContext) Primitive_name() IPrimitive_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitive_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitive_nameContext)
}

func (s *Type_nameContext) Custom_type_name() ICustom_type_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICustom_type_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICustom_type_nameContext)
}

func (s *Type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterType_name(s)
	}
}

func (s *Type_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitType_name(s)
	}
}

func (p *ASN1SchemaParser) Type_name() (localctx IType_nameContext) {
	localctx = NewType_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ASN1SchemaParserRULE_type_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserSEQUENCE_LITERAL, ASN1SchemaParserINTEGER_LITERAL, ASN1SchemaParserANY_LITERAL, ASN1SchemaParserBMPSTRING_LITERAL, ASN1SchemaParserIA5STRING_LITERAL, ASN1SchemaParserPRINTABLE_STRING_LITERAL, ASN1SchemaParserTELETEX_STRING_LITERAL, ASN1SchemaParserUNIVERSAL_STRING_LITERAL, ASN1SchemaParserUTF8_STRING_LITERAL, ASN1SchemaParserBIT_STRING_LITERAL, ASN1SchemaParserOCTET_STRING_LITERAL, ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.Primitive_name()
		}

	case ASN1SchemaParserWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.Custom_type_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInteger_oid_referenceContext is an interface to support dynamic dispatch.
type IInteger_oid_referenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_oid_referenceContext differentiates from other interfaces.
	IsInteger_oid_referenceContext()
}

type Integer_oid_referenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_oid_referenceContext() *Integer_oid_referenceContext {
	var p = new(Integer_oid_referenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_integer_oid_reference
	return p
}

func (*Integer_oid_referenceContext) IsInteger_oid_referenceContext() {}

func NewInteger_oid_referenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_oid_referenceContext {
	var p = new(Integer_oid_referenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_integer_oid_reference

	return p
}

func (s *Integer_oid_referenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_oid_referenceContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Integer_oid_referenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_oid_referenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_oid_referenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterInteger_oid_reference(s)
	}
}

func (s *Integer_oid_referenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitInteger_oid_reference(s)
	}
}

func (p *ASN1SchemaParser) Integer_oid_reference() (localctx IInteger_oid_referenceContext) {
	localctx = NewInteger_oid_referenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ASN1SchemaParserRULE_integer_oid_reference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Field_name()
	}

	return localctx
}

// IValue_identifierContext is an interface to support dynamic dispatch.
type IValue_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValue_identifierContext differentiates from other interfaces.
	IsValue_identifierContext()
}

type Value_identifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_identifierContext() *Value_identifierContext {
	var p = new(Value_identifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_value_identifier
	return p
}

func (*Value_identifierContext) IsValue_identifierContext() {}

func NewValue_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_identifierContext {
	var p = new(Value_identifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_value_identifier

	return p
}

func (s *Value_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_identifierContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserINTEGER, 0)
}

func (s *Value_identifierContext) WORD() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserWORD, 0)
}

func (s *Value_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterValue_identifier(s)
	}
}

func (s *Value_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitValue_identifier(s)
	}
}

func (p *ASN1SchemaParser) Value_identifier() (localctx IValue_identifierContext) {
	localctx = NewValue_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ASN1SchemaParserRULE_value_identifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ASN1SchemaParserINTEGER || _la == ASN1SchemaParserWORD) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IField_nameContext is an interface to support dynamic dispatch.
type IField_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_nameContext differentiates from other interfaces.
	IsField_nameContext()
}

type Field_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_nameContext() *Field_nameContext {
	var p = new(Field_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_field_name
	return p
}

func (*Field_nameContext) IsField_nameContext() {}

func NewField_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_nameContext {
	var p = new(Field_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_field_name

	return p
}

func (s *Field_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_nameContext) WORD() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserWORD, 0)
}

func (s *Field_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterField_name(s)
	}
}

func (s *Field_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitField_name(s)
	}
}

func (p *ASN1SchemaParser) Field_name() (localctx IField_nameContext) {
	localctx = NewField_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ASN1SchemaParserRULE_field_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(ASN1SchemaParserWORD)
	}

	return localctx
}

// ICustom_type_nameContext is an interface to support dynamic dispatch.
type ICustom_type_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCustom_type_nameContext differentiates from other interfaces.
	IsCustom_type_nameContext()
}

type Custom_type_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCustom_type_nameContext() *Custom_type_nameContext {
	var p = new(Custom_type_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_custom_type_name
	return p
}

func (*Custom_type_nameContext) IsCustom_type_nameContext() {}

func NewCustom_type_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Custom_type_nameContext {
	var p = new(Custom_type_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_custom_type_name

	return p
}

func (s *Custom_type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Custom_type_nameContext) WORD() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserWORD, 0)
}

func (s *Custom_type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Custom_type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Custom_type_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterCustom_type_name(s)
	}
}

func (s *Custom_type_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitCustom_type_name(s)
	}
}

func (p *ASN1SchemaParser) Custom_type_name() (localctx ICustom_type_nameContext) {
	localctx = NewCustom_type_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ASN1SchemaParserRULE_custom_type_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(ASN1SchemaParserWORD)
	}

	return localctx
}
