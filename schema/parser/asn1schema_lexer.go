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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 107,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 6, 2, 33, 10, 2, 13, 2,
	14, 2, 34, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 41, 10, 3, 12, 3, 14, 3, 44, 11,
	3, 3, 4, 6, 4, 47, 10, 4, 13, 4, 14, 4, 48, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 2, 2,
	16, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 27, 15, 29, 16, 3, 2, 5, 5, 2, 11, 12, 15, 15, 34, 34,
	4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 2, 109, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 3, 32, 3, 2, 2, 2, 5, 38, 3, 2, 2, 2,
	7, 46, 3, 2, 2, 2, 9, 50, 3, 2, 2, 2, 11, 54, 3, 2, 2, 2, 13, 56, 3, 2,
	2, 2, 15, 58, 3, 2, 2, 2, 17, 60, 3, 2, 2, 2, 19, 62, 3, 2, 2, 2, 21, 64,
	3, 2, 2, 2, 23, 73, 3, 2, 2, 2, 25, 81, 3, 2, 2, 2, 27, 90, 3, 2, 2, 2,
	29, 99, 3, 2, 2, 2, 31, 33, 9, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3,
	2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36,
	37, 8, 2, 2, 2, 37, 4, 3, 2, 2, 2, 38, 42, 9, 3, 2, 2, 39, 41, 9, 4, 2,
	2, 40, 39, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43,
	3, 2, 2, 2, 43, 6, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 47, 4, 50, 59, 2,
	46, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3,
	2, 2, 2, 49, 8, 3, 2, 2, 2, 50, 51, 7, 60, 2, 2, 51, 52, 7, 60, 2, 2, 52,
	53, 7, 63, 2, 2, 53, 10, 3, 2, 2, 2, 54, 55, 7, 125, 2, 2, 55, 12, 3, 2,
	2, 2, 56, 57, 7, 127, 2, 2, 57, 14, 3, 2, 2, 2, 58, 59, 7, 93, 2, 2, 59,
	16, 3, 2, 2, 2, 60, 61, 7, 95, 2, 2, 61, 18, 3, 2, 2, 2, 62, 63, 7, 46,
	2, 2, 63, 20, 3, 2, 2, 2, 64, 65, 7, 85, 2, 2, 65, 66, 7, 71, 2, 2, 66,
	67, 7, 83, 2, 2, 67, 68, 7, 87, 2, 2, 68, 69, 7, 71, 2, 2, 69, 70, 7, 80,
	2, 2, 70, 71, 7, 69, 2, 2, 71, 72, 7, 71, 2, 2, 72, 22, 3, 2, 2, 2, 73,
	74, 7, 75, 2, 2, 74, 75, 7, 80, 2, 2, 75, 76, 7, 86, 2, 2, 76, 77, 7, 71,
	2, 2, 77, 78, 7, 73, 2, 2, 78, 79, 7, 71, 2, 2, 79, 80, 7, 84, 2, 2, 80,
	24, 3, 2, 2, 2, 81, 82, 7, 75, 2, 2, 82, 83, 7, 79, 2, 2, 83, 84, 7, 82,
	2, 2, 84, 85, 7, 78, 2, 2, 85, 86, 7, 75, 2, 2, 86, 87, 7, 69, 2, 2, 87,
	88, 7, 75, 2, 2, 88, 89, 7, 86, 2, 2, 89, 26, 3, 2, 2, 2, 90, 91, 7, 81,
	2, 2, 91, 92, 7, 82, 2, 2, 92, 93, 7, 86, 2, 2, 93, 94, 7, 75, 2, 2, 94,
	95, 7, 81, 2, 2, 95, 96, 7, 80, 2, 2, 96, 97, 7, 67, 2, 2, 97, 98, 7, 78,
	2, 2, 98, 28, 3, 2, 2, 2, 99, 100, 7, 70, 2, 2, 100, 101, 7, 71, 2, 2,
	101, 102, 7, 72, 2, 2, 102, 103, 7, 67, 2, 2, 103, 104, 7, 87, 2, 2, 104,
	105, 7, 78, 2, 2, 105, 106, 7, 86, 2, 2, 106, 30, 3, 2, 2, 2, 6, 2, 34,
	42, 48, 3, 8, 2, 2,
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
	"", "", "", "", "'::='", "'{'", "'}'", "'['", "']'", "','", "'SEQUENCE'",
	"'INTEGER'", "'IMPLICIT'", "'OPTIONAL'", "'DEFAULT'",
}

var lexerSymbolicNames = []string{
	"", "WHITESPACE", "WORD", "INTEGER", "ASSIGN", "LBRACE", "RBRACE", "LBRACKET",
	"RBRACKET", "COMMA", "SEQUENCE_LITERAL", "INTEGER_LITERAL", "IMPLICIT_LITERAL",
	"OPTIONAL_LITERAL", "DEFAULT_LITERAL",
}

var lexerRuleNames = []string{
	"WHITESPACE", "WORD", "INTEGER", "ASSIGN", "LBRACE", "RBRACE", "LBRACKET",
	"RBRACKET", "COMMA", "SEQUENCE_LITERAL", "INTEGER_LITERAL", "IMPLICIT_LITERAL",
	"OPTIONAL_LITERAL", "DEFAULT_LITERAL",
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
	ASN1SchemaLexerWHITESPACE       = 1
	ASN1SchemaLexerWORD             = 2
	ASN1SchemaLexerINTEGER          = 3
	ASN1SchemaLexerASSIGN           = 4
	ASN1SchemaLexerLBRACE           = 5
	ASN1SchemaLexerRBRACE           = 6
	ASN1SchemaLexerLBRACKET         = 7
	ASN1SchemaLexerRBRACKET         = 8
	ASN1SchemaLexerCOMMA            = 9
	ASN1SchemaLexerSEQUENCE_LITERAL = 10
	ASN1SchemaLexerINTEGER_LITERAL  = 11
	ASN1SchemaLexerIMPLICIT_LITERAL = 12
	ASN1SchemaLexerOPTIONAL_LITERAL = 13
	ASN1SchemaLexerDEFAULT_LITERAL  = 14
)
