// Code generated from ASN1Schema.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 140,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 6,
	17, 123, 10, 17, 13, 17, 14, 17, 124, 3, 18, 3, 18, 7, 18, 129, 10, 18,
	12, 18, 14, 18, 132, 11, 18, 3, 19, 6, 19, 135, 10, 19, 13, 19, 14, 19,
	136, 3, 19, 3, 19, 2, 2, 20, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
	8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
	35, 18, 37, 19, 3, 2, 5, 5, 2, 11, 12, 14, 15, 34, 34, 4, 2, 67, 92, 99,
	124, 5, 2, 50, 59, 67, 92, 99, 124, 2, 141, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 3, 39, 3, 2, 2, 2, 5, 41, 3, 2, 2, 2, 7, 45, 3, 2, 2, 2, 9, 47, 3, 2,
	2, 2, 11, 49, 3, 2, 2, 2, 13, 51, 3, 2, 2, 2, 15, 53, 3, 2, 2, 2, 17, 55,
	3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 72, 3, 2, 2, 2, 23, 81, 3, 2, 2, 2,
	25, 90, 3, 2, 2, 2, 27, 98, 3, 2, 2, 2, 29, 102, 3, 2, 2, 2, 31, 109, 3,
	2, 2, 2, 33, 122, 3, 2, 2, 2, 35, 126, 3, 2, 2, 2, 37, 134, 3, 2, 2, 2,
	39, 40, 9, 2, 2, 2, 40, 4, 3, 2, 2, 2, 41, 42, 7, 60, 2, 2, 42, 43, 7,
	60, 2, 2, 43, 44, 7, 63, 2, 2, 44, 6, 3, 2, 2, 2, 45, 46, 7, 125, 2, 2,
	46, 8, 3, 2, 2, 2, 47, 48, 7, 127, 2, 2, 48, 10, 3, 2, 2, 2, 49, 50, 7,
	93, 2, 2, 50, 12, 3, 2, 2, 2, 51, 52, 7, 95, 2, 2, 52, 14, 3, 2, 2, 2,
	53, 54, 7, 46, 2, 2, 54, 16, 3, 2, 2, 2, 55, 56, 7, 85, 2, 2, 56, 57, 7,
	71, 2, 2, 57, 58, 7, 83, 2, 2, 58, 59, 7, 87, 2, 2, 59, 60, 7, 71, 2, 2,
	60, 61, 7, 80, 2, 2, 61, 62, 7, 69, 2, 2, 62, 63, 7, 71, 2, 2, 63, 18,
	3, 2, 2, 2, 64, 65, 7, 75, 2, 2, 65, 66, 7, 80, 2, 2, 66, 67, 7, 86, 2,
	2, 67, 68, 7, 71, 2, 2, 68, 69, 7, 73, 2, 2, 69, 70, 7, 71, 2, 2, 70, 71,
	7, 84, 2, 2, 71, 20, 3, 2, 2, 2, 72, 73, 7, 75, 2, 2, 73, 74, 7, 79, 2,
	2, 74, 75, 7, 82, 2, 2, 75, 76, 7, 78, 2, 2, 76, 77, 7, 75, 2, 2, 77, 78,
	7, 69, 2, 2, 78, 79, 7, 75, 2, 2, 79, 80, 7, 86, 2, 2, 80, 22, 3, 2, 2,
	2, 81, 82, 7, 81, 2, 2, 82, 83, 7, 82, 2, 2, 83, 84, 7, 86, 2, 2, 84, 85,
	7, 75, 2, 2, 85, 86, 7, 81, 2, 2, 86, 87, 7, 80, 2, 2, 87, 88, 7, 67, 2,
	2, 88, 89, 7, 78, 2, 2, 89, 24, 3, 2, 2, 2, 90, 91, 7, 70, 2, 2, 91, 92,
	7, 71, 2, 2, 92, 93, 7, 72, 2, 2, 93, 94, 7, 67, 2, 2, 94, 95, 7, 87, 2,
	2, 95, 96, 7, 78, 2, 2, 96, 97, 7, 86, 2, 2, 97, 26, 3, 2, 2, 2, 98, 99,
	7, 68, 2, 2, 99, 100, 7, 75, 2, 2, 100, 101, 7, 86, 2, 2, 101, 28, 3, 2,
	2, 2, 102, 103, 7, 85, 2, 2, 103, 104, 7, 86, 2, 2, 104, 105, 7, 84, 2,
	2, 105, 106, 7, 75, 2, 2, 106, 107, 7, 80, 2, 2, 107, 108, 7, 73, 2, 2,
	108, 30, 3, 2, 2, 2, 109, 110, 7, 68, 2, 2, 110, 111, 7, 75, 2, 2, 111,
	112, 7, 86, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 5, 3, 2, 2, 114, 115,
	7, 85, 2, 2, 115, 116, 7, 86, 2, 2, 116, 117, 7, 84, 2, 2, 117, 118, 7,
	75, 2, 2, 118, 119, 7, 80, 2, 2, 119, 120, 7, 73, 2, 2, 120, 32, 3, 2,
	2, 2, 121, 123, 4, 50, 59, 2, 122, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2,
	2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 34, 3, 2, 2, 2, 126,
	130, 9, 3, 2, 2, 127, 129, 9, 4, 2, 2, 128, 127, 3, 2, 2, 2, 129, 132,
	3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 36, 3, 2,
	2, 2, 132, 130, 3, 2, 2, 2, 133, 135, 5, 3, 2, 2, 134, 133, 3, 2, 2, 2,
	135, 136, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 139, 8, 19, 2, 2, 139, 38, 3, 2, 2, 2, 6, 2, 124,
	130, 136, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'::='", "'{'", "'}'", "'['", "']'", "','", "'SEQUENCE'", "'INTEGER'",
	"'IMPLICIT'", "'OPTIONAL'", "'DEFAULT'", "'BIT'", "'STRING'",
}

var lexerSymbolicNames = []string{
	"", "ASSIGN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "COMMA", "SEQUENCE_LITERAL",
	"INTEGER_LITERAL", "IMPLICIT_LITERAL", "OPTIONAL_LITERAL", "DEFAULT_LITERAL",
	"BIT_LITERAL", "STRING_LITERAL", "BIT_STRING_LITERAL", "INTEGER", "WORD",
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"SPACE", "ASSIGN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "COMMA",
	"SEQUENCE_LITERAL", "INTEGER_LITERAL", "IMPLICIT_LITERAL", "OPTIONAL_LITERAL",
	"DEFAULT_LITERAL", "BIT_LITERAL", "STRING_LITERAL", "BIT_STRING_LITERAL",
	"INTEGER", "WORD", "WHITESPACE",
}

type ASN1SchemaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewASN1SchemaLexer(input antlr.CharStream) *ASN1SchemaLexer {

	l := new(ASN1SchemaLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ASN1Schema.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ASN1SchemaLexer tokens.
const (
	ASN1SchemaLexerASSIGN             = 1
	ASN1SchemaLexerLBRACE             = 2
	ASN1SchemaLexerRBRACE             = 3
	ASN1SchemaLexerLBRACKET           = 4
	ASN1SchemaLexerRBRACKET           = 5
	ASN1SchemaLexerCOMMA              = 6
	ASN1SchemaLexerSEQUENCE_LITERAL   = 7
	ASN1SchemaLexerINTEGER_LITERAL    = 8
	ASN1SchemaLexerIMPLICIT_LITERAL   = 9
	ASN1SchemaLexerOPTIONAL_LITERAL   = 10
	ASN1SchemaLexerDEFAULT_LITERAL    = 11
	ASN1SchemaLexerBIT_LITERAL        = 12
	ASN1SchemaLexerSTRING_LITERAL     = 13
	ASN1SchemaLexerBIT_STRING_LITERAL = 14
	ASN1SchemaLexerINTEGER            = 15
	ASN1SchemaLexerWORD               = 16
	ASN1SchemaLexerWHITESPACE         = 17
)
