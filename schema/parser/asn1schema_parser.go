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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 103,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 6,
	2, 36, 10, 2, 13, 2, 14, 2, 37, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 53, 10, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 65, 10, 8, 3, 9, 3,
	9, 3, 9, 3, 10, 5, 10, 71, 10, 10, 3, 10, 3, 10, 5, 10, 75, 10, 10, 3,
	10, 5, 10, 78, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	5, 12, 87, 10, 12, 3, 13, 3, 13, 5, 13, 91, 10, 13, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 2, 2, 18, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 3, 3, 2, 14, 17,
	2, 98, 2, 35, 3, 2, 2, 2, 4, 41, 3, 2, 2, 2, 6, 52, 3, 2, 2, 2, 8, 54,
	3, 2, 2, 2, 10, 56, 3, 2, 2, 2, 12, 60, 3, 2, 2, 2, 14, 64, 3, 2, 2, 2,
	16, 66, 3, 2, 2, 2, 18, 70, 3, 2, 2, 2, 20, 79, 3, 2, 2, 2, 22, 83, 3,
	2, 2, 2, 24, 88, 3, 2, 2, 2, 26, 94, 3, 2, 2, 2, 28, 96, 3, 2, 2, 2, 30,
	98, 3, 2, 2, 2, 32, 100, 3, 2, 2, 2, 34, 36, 5, 4, 3, 2, 35, 34, 3, 2,
	2, 2, 36, 37, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 39,
	3, 2, 2, 2, 39, 40, 7, 2, 2, 3, 40, 3, 3, 2, 2, 2, 41, 42, 5, 30, 16, 2,
	42, 43, 7, 3, 2, 2, 43, 44, 5, 6, 4, 2, 44, 45, 5, 8, 5, 2, 45, 5, 3, 2,
	2, 2, 46, 53, 7, 9, 2, 2, 47, 53, 7, 20, 2, 2, 48, 53, 7, 10, 2, 2, 49,
	53, 7, 19, 2, 2, 50, 53, 7, 18, 2, 2, 51, 53, 5, 26, 14, 2, 52, 46, 3,
	2, 2, 2, 52, 47, 3, 2, 2, 2, 52, 48, 3, 2, 2, 2, 52, 49, 3, 2, 2, 2, 52,
	50, 3, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 7, 3, 2, 2, 2, 54, 55, 5, 20, 11,
	2, 55, 9, 3, 2, 2, 2, 56, 57, 7, 6, 2, 2, 57, 58, 7, 21, 2, 2, 58, 59,
	7, 7, 2, 2, 59, 11, 3, 2, 2, 2, 60, 61, 7, 11, 2, 2, 61, 13, 3, 2, 2, 2,
	62, 65, 7, 12, 2, 2, 63, 65, 5, 16, 9, 2, 64, 62, 3, 2, 2, 2, 64, 63, 3,
	2, 2, 2, 65, 15, 3, 2, 2, 2, 66, 67, 7, 13, 2, 2, 67, 68, 5, 32, 17, 2,
	68, 17, 3, 2, 2, 2, 69, 71, 5, 12, 7, 2, 70, 69, 3, 2, 2, 2, 70, 71, 3,
	2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 75, 5, 6, 4, 2, 73, 75, 5, 30, 16, 2,
	74, 72, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 77, 3, 2, 2, 2, 76, 78, 5,
	14, 8, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 19, 3, 2, 2, 2, 79,
	80, 7, 4, 2, 2, 80, 81, 5, 22, 12, 2, 81, 82, 7, 5, 2, 2, 82, 21, 3, 2,
	2, 2, 83, 86, 5, 24, 13, 2, 84, 85, 7, 8, 2, 2, 85, 87, 5, 22, 12, 2, 86,
	84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 23, 3, 2, 2, 2, 88, 90, 5, 28,
	15, 2, 89, 91, 5, 10, 6, 2, 90, 89, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91,
	92, 3, 2, 2, 2, 92, 93, 5, 18, 10, 2, 93, 25, 3, 2, 2, 2, 94, 95, 9, 2,
	2, 2, 95, 27, 3, 2, 2, 2, 96, 97, 7, 22, 2, 2, 97, 29, 3, 2, 2, 2, 98,
	99, 7, 22, 2, 2, 99, 31, 3, 2, 2, 2, 100, 101, 7, 22, 2, 2, 101, 33, 3,
	2, 2, 2, 10, 37, 52, 64, 70, 74, 77, 86, 90,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'::='", "'{'", "'}'", "'['", "']'", "','", "'SEQUENCE'", "'INTEGER'",
	"'IMPLICIT'", "'OPTIONAL'", "'DEFAULT'", "'BMPString'", "'IA5String'",
	"'PrintableString'", "'UTF8String'",
}
var symbolicNames = []string{
	"", "ASSIGN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "COMMA", "SEQUENCE_LITERAL",
	"INTEGER_LITERAL", "IMPLICIT_LITERAL", "OPTIONAL_LITERAL", "DEFAULT_LITERAL",
	"BMPSTRING_LITERAL", "IA5STRING_LITERAL", "PRINTABLE_STRING_LITERAL", "UTF8_STRING_LITERAL",
	"BIT_STRING_LITERAL", "OCTET_STRING_LITERAL", "OBJECT_IDENTIFIER_LITERAL",
	"INTEGER", "WORD", "WHITESPACE",
}

var ruleNames = []string{
	"module", "assignment", "primitive", "primitive_definition", "tag", "context_flag",
	"encoding_flags", "default_value", "type_definition", "sequence_definition",
	"field_list", "field_definition", "string_literal", "field_name", "type_name",
	"value_identifier",
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
	ASN1SchemaParserASSIGN                    = 1
	ASN1SchemaParserLBRACE                    = 2
	ASN1SchemaParserRBRACE                    = 3
	ASN1SchemaParserLBRACKET                  = 4
	ASN1SchemaParserRBRACKET                  = 5
	ASN1SchemaParserCOMMA                     = 6
	ASN1SchemaParserSEQUENCE_LITERAL          = 7
	ASN1SchemaParserINTEGER_LITERAL           = 8
	ASN1SchemaParserIMPLICIT_LITERAL          = 9
	ASN1SchemaParserOPTIONAL_LITERAL          = 10
	ASN1SchemaParserDEFAULT_LITERAL           = 11
	ASN1SchemaParserBMPSTRING_LITERAL         = 12
	ASN1SchemaParserIA5STRING_LITERAL         = 13
	ASN1SchemaParserPRINTABLE_STRING_LITERAL  = 14
	ASN1SchemaParserUTF8_STRING_LITERAL       = 15
	ASN1SchemaParserBIT_STRING_LITERAL        = 16
	ASN1SchemaParserOCTET_STRING_LITERAL      = 17
	ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL = 18
	ASN1SchemaParserINTEGER                   = 19
	ASN1SchemaParserWORD                      = 20
	ASN1SchemaParserWHITESPACE                = 21
)

// ASN1SchemaParser rules.
const (
	ASN1SchemaParserRULE_module               = 0
	ASN1SchemaParserRULE_assignment           = 1
	ASN1SchemaParserRULE_primitive            = 2
	ASN1SchemaParserRULE_primitive_definition = 3
	ASN1SchemaParserRULE_tag                  = 4
	ASN1SchemaParserRULE_context_flag         = 5
	ASN1SchemaParserRULE_encoding_flags       = 6
	ASN1SchemaParserRULE_default_value        = 7
	ASN1SchemaParserRULE_type_definition      = 8
	ASN1SchemaParserRULE_sequence_definition  = 9
	ASN1SchemaParserRULE_field_list           = 10
	ASN1SchemaParserRULE_field_definition     = 11
	ASN1SchemaParserRULE_string_literal       = 12
	ASN1SchemaParserRULE_field_name           = 13
	ASN1SchemaParserRULE_type_name            = 14
	ASN1SchemaParserRULE_value_identifier     = 15
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ASN1SchemaParserWORD {
		{
			p.SetState(32)
			p.Assignment()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
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

func (s *AssignmentContext) Primitive_definition() IPrimitive_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitive_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitive_definitionContext)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Type_name()
	}
	{
		p.SetState(40)
		p.Match(ASN1SchemaParserASSIGN)
	}
	{
		p.SetState(41)
		p.Primitive()
	}
	{
		p.SetState(42)
		p.Primitive_definition()
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

func (s *PrimitiveContext) SEQUENCE_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserSEQUENCE_LITERAL, 0)
}

func (s *PrimitiveContext) OBJECT_IDENTIFIER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL, 0)
}

func (s *PrimitiveContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserINTEGER_LITERAL, 0)
}

func (s *PrimitiveContext) OCTET_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOCTET_STRING_LITERAL, 0)
}

func (s *PrimitiveContext) BIT_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserBIT_STRING_LITERAL, 0)
}

func (s *PrimitiveContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
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

	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserSEQUENCE_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Match(ASN1SchemaParserSEQUENCE_LITERAL)
		}

	case ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL)
		}

	case ASN1SchemaParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.Match(ASN1SchemaParserINTEGER_LITERAL)
		}

	case ASN1SchemaParserOCTET_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(47)
			p.Match(ASN1SchemaParserOCTET_STRING_LITERAL)
		}

	case ASN1SchemaParserBIT_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(48)
			p.Match(ASN1SchemaParserBIT_STRING_LITERAL)
		}

	case ASN1SchemaParserBMPSTRING_LITERAL, ASN1SchemaParserIA5STRING_LITERAL, ASN1SchemaParserPRINTABLE_STRING_LITERAL, ASN1SchemaParserUTF8_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(49)
			p.String_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimitive_definitionContext is an interface to support dynamic dispatch.
type IPrimitive_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitive_definitionContext differentiates from other interfaces.
	IsPrimitive_definitionContext()
}

type Primitive_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitive_definitionContext() *Primitive_definitionContext {
	var p = new(Primitive_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_primitive_definition
	return p
}

func (*Primitive_definitionContext) IsPrimitive_definitionContext() {}

func NewPrimitive_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primitive_definitionContext {
	var p = new(Primitive_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_primitive_definition

	return p
}

func (s *Primitive_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Primitive_definitionContext) Sequence_definition() ISequence_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_definitionContext)
}

func (s *Primitive_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitive_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primitive_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterPrimitive_definition(s)
	}
}

func (s *Primitive_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitPrimitive_definition(s)
	}
}

func (p *ASN1SchemaParser) Primitive_definition() (localctx IPrimitive_definitionContext) {
	localctx = NewPrimitive_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ASN1SchemaParserRULE_primitive_definition)

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
		p.SetState(52)
		p.Sequence_definition()
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
		p.SetState(54)
		p.Match(ASN1SchemaParserLBRACKET)
	}
	{
		p.SetState(55)
		p.Match(ASN1SchemaParserINTEGER)
	}
	{
		p.SetState(56)
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
		p.SetState(58)
		p.Match(ASN1SchemaParserIMPLICIT_LITERAL)
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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserOPTIONAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Match(ASN1SchemaParserOPTIONAL_LITERAL)
		}

	case ASN1SchemaParserDEFAULT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
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
		p.SetState(64)
		p.Match(ASN1SchemaParserDEFAULT_LITERAL)
	}
	{
		p.SetState(65)
		p.Value_identifier()
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

func (s *Type_definitionContext) Type_name() IType_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
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
	p.EnterRule(localctx, 16, ASN1SchemaParserRULE_type_definition)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserIMPLICIT_LITERAL {
		{
			p.SetState(67)
			p.Context_flag()
		}

	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserSEQUENCE_LITERAL, ASN1SchemaParserINTEGER_LITERAL, ASN1SchemaParserBMPSTRING_LITERAL, ASN1SchemaParserIA5STRING_LITERAL, ASN1SchemaParserPRINTABLE_STRING_LITERAL, ASN1SchemaParserUTF8_STRING_LITERAL, ASN1SchemaParserBIT_STRING_LITERAL, ASN1SchemaParserOCTET_STRING_LITERAL, ASN1SchemaParserOBJECT_IDENTIFIER_LITERAL:
		{
			p.SetState(70)
			p.Primitive()
		}

	case ASN1SchemaParserWORD:
		{
			p.SetState(71)
			p.Type_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserOPTIONAL_LITERAL || _la == ASN1SchemaParserDEFAULT_LITERAL {
		{
			p.SetState(74)
			p.Encoding_flags()
		}

	}

	return localctx
}

// ISequence_definitionContext is an interface to support dynamic dispatch.
type ISequence_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequence_definitionContext differentiates from other interfaces.
	IsSequence_definitionContext()
}

type Sequence_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequence_definitionContext() *Sequence_definitionContext {
	var p = new(Sequence_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ASN1SchemaParserRULE_sequence_definition
	return p
}

func (*Sequence_definitionContext) IsSequence_definitionContext() {}

func NewSequence_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sequence_definitionContext {
	var p = new(Sequence_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ASN1SchemaParserRULE_sequence_definition

	return p
}

func (s *Sequence_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Sequence_definitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserLBRACE, 0)
}

func (s *Sequence_definitionContext) Field_list() IField_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_listContext)
}

func (s *Sequence_definitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserRBRACE, 0)
}

func (s *Sequence_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sequence_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sequence_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.EnterSequence_definition(s)
	}
}

func (s *Sequence_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ASN1SchemaListener); ok {
		listenerT.ExitSequence_definition(s)
	}
}

func (p *ASN1SchemaParser) Sequence_definition() (localctx ISequence_definitionContext) {
	localctx = NewSequence_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ASN1SchemaParserRULE_sequence_definition)

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
		p.SetState(77)
		p.Match(ASN1SchemaParserLBRACE)
	}
	{
		p.SetState(78)
		p.Field_list()
	}
	{
		p.SetState(79)
		p.Match(ASN1SchemaParserRBRACE)
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
	p.EnterRule(localctx, 20, ASN1SchemaParserRULE_field_list)
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
		p.SetState(81)
		p.Field_definition()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserCOMMA {
		{
			p.SetState(82)
			p.Match(ASN1SchemaParserCOMMA)
		}
		{
			p.SetState(83)
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
	p.EnterRule(localctx, 22, ASN1SchemaParserRULE_field_definition)
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
		p.SetState(86)
		p.Field_name()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserLBRACKET {
		{
			p.SetState(87)
			p.Tag()
		}

	}
	{
		p.SetState(90)
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

func (s *String_literalContext) UTF8_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserUTF8_STRING_LITERAL, 0)
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
	p.EnterRule(localctx, 24, ASN1SchemaParserRULE_string_literal)
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
		p.SetState(92)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ASN1SchemaParserBMPSTRING_LITERAL)|(1<<ASN1SchemaParserIA5STRING_LITERAL)|(1<<ASN1SchemaParserPRINTABLE_STRING_LITERAL)|(1<<ASN1SchemaParserUTF8_STRING_LITERAL))) != 0) {
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
	p.EnterRule(localctx, 26, ASN1SchemaParserRULE_field_name)

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
		p.SetState(94)
		p.Match(ASN1SchemaParserWORD)
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

func (s *Type_nameContext) WORD() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserWORD, 0)
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
	p.EnterRule(localctx, 28, ASN1SchemaParserRULE_type_name)

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
		p.SetState(96)
		p.Match(ASN1SchemaParserWORD)
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
	p.EnterRule(localctx, 30, ASN1SchemaParserRULE_value_identifier)

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
		p.SetState(98)
		p.Match(ASN1SchemaParserWORD)
	}

	return localctx
}
