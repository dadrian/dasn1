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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 73, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 37,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 5, 8, 45, 10, 8, 3, 8, 3, 8,
	5, 8, 49, 10, 8, 3, 8, 5, 8, 52, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 63, 10, 11, 3, 12, 3, 12, 5, 12, 67,
	10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 2, 3, 3, 2, 12, 13, 2, 66, 2, 26, 3, 2, 2, 2, 4,
	30, 3, 2, 2, 2, 6, 32, 3, 2, 2, 2, 8, 36, 3, 2, 2, 2, 10, 38, 3, 2, 2,
	2, 12, 41, 3, 2, 2, 2, 14, 44, 3, 2, 2, 2, 16, 53, 3, 2, 2, 2, 18, 55,
	3, 2, 2, 2, 20, 59, 3, 2, 2, 2, 22, 64, 3, 2, 2, 2, 24, 70, 3, 2, 2, 2,
	26, 27, 7, 9, 2, 2, 27, 28, 7, 5, 2, 2, 28, 29, 7, 10, 2, 2, 29, 3, 3,
	2, 2, 2, 30, 31, 9, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 33, 7, 14, 2, 2, 33,
	7, 3, 2, 2, 2, 34, 37, 5, 10, 6, 2, 35, 37, 7, 15, 2, 2, 36, 34, 3, 2,
	2, 2, 36, 35, 3, 2, 2, 2, 37, 9, 3, 2, 2, 2, 38, 39, 7, 16, 2, 2, 39, 40,
	5, 12, 7, 2, 40, 11, 3, 2, 2, 2, 41, 42, 7, 4, 2, 2, 42, 13, 3, 2, 2, 2,
	43, 45, 5, 6, 4, 2, 44, 43, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 48, 3,
	2, 2, 2, 46, 49, 5, 4, 3, 2, 47, 49, 5, 16, 9, 2, 48, 46, 3, 2, 2, 2, 48,
	47, 3, 2, 2, 2, 49, 51, 3, 2, 2, 2, 50, 52, 5, 8, 5, 2, 51, 50, 3, 2, 2,
	2, 51, 52, 3, 2, 2, 2, 52, 15, 3, 2, 2, 2, 53, 54, 7, 4, 2, 2, 54, 17,
	3, 2, 2, 2, 55, 56, 7, 7, 2, 2, 56, 57, 5, 20, 11, 2, 57, 58, 7, 8, 2,
	2, 58, 19, 3, 2, 2, 2, 59, 62, 5, 22, 12, 2, 60, 61, 7, 11, 2, 2, 61, 63,
	5, 22, 12, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 21, 3, 2, 2,
	2, 64, 66, 5, 24, 13, 2, 65, 67, 5, 2, 2, 2, 66, 65, 3, 2, 2, 2, 66, 67,
	3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 5, 14, 8, 2, 69, 23, 3, 2, 2, 2,
	70, 71, 7, 4, 2, 2, 71, 25, 3, 2, 2, 2, 8, 36, 44, 48, 51, 62, 66,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "'::='", "'{'", "'}'", "'['", "']'", "','", "'SEQUENCE'",
	"'INTEGER'", "'IMPLICIT'", "'OPTIONAL'", "'DEFAULT'",
}
var symbolicNames = []string{
	"", "WHITESPACE", "WORD", "INTEGER", "ASSIGN", "LBRACE", "RBRACE", "LBRACKET",
	"RBRACKET", "COMMA", "SEQUENCE_LITERAL", "INTEGER_LITERAL", "IMPLICIT_LITERAL",
	"OPTIONAL_LITERAL", "DEFAULT_LITERAL",
}

var ruleNames = []string{
	"tag", "primitive", "context_flag", "encoding_flags", "default_value",
	"value_identifier", "type_definition", "type_name", "sequence_definition",
	"field_list", "field_definition", "field_name",
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
	ASN1SchemaParserEOF              = antlr.TokenEOF
	ASN1SchemaParserWHITESPACE       = 1
	ASN1SchemaParserWORD             = 2
	ASN1SchemaParserINTEGER          = 3
	ASN1SchemaParserASSIGN           = 4
	ASN1SchemaParserLBRACE           = 5
	ASN1SchemaParserRBRACE           = 6
	ASN1SchemaParserLBRACKET         = 7
	ASN1SchemaParserRBRACKET         = 8
	ASN1SchemaParserCOMMA            = 9
	ASN1SchemaParserSEQUENCE_LITERAL = 10
	ASN1SchemaParserINTEGER_LITERAL  = 11
	ASN1SchemaParserIMPLICIT_LITERAL = 12
	ASN1SchemaParserOPTIONAL_LITERAL = 13
	ASN1SchemaParserDEFAULT_LITERAL  = 14
)

// ASN1SchemaParser rules.
const (
	ASN1SchemaParserRULE_tag                 = 0
	ASN1SchemaParserRULE_primitive           = 1
	ASN1SchemaParserRULE_context_flag        = 2
	ASN1SchemaParserRULE_encoding_flags      = 3
	ASN1SchemaParserRULE_default_value       = 4
	ASN1SchemaParserRULE_value_identifier    = 5
	ASN1SchemaParserRULE_type_definition     = 6
	ASN1SchemaParserRULE_type_name           = 7
	ASN1SchemaParserRULE_sequence_definition = 8
	ASN1SchemaParserRULE_field_list          = 9
	ASN1SchemaParserRULE_field_definition    = 10
	ASN1SchemaParserRULE_field_name          = 11
)

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
	p.EnterRule(localctx, 0, ASN1SchemaParserRULE_tag)

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
		p.SetState(24)
		p.Match(ASN1SchemaParserLBRACKET)
	}
	{
		p.SetState(25)
		p.Match(ASN1SchemaParserINTEGER)
	}
	{
		p.SetState(26)
		p.Match(ASN1SchemaParserRBRACKET)
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

func (s *PrimitiveContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserINTEGER_LITERAL, 0)
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
	p.EnterRule(localctx, 2, ASN1SchemaParserRULE_primitive)
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
		p.SetState(28)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ASN1SchemaParserSEQUENCE_LITERAL || _la == ASN1SchemaParserINTEGER_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.EnterRule(localctx, 4, ASN1SchemaParserRULE_context_flag)

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
		p.SetState(30)
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

func (s *Encoding_flagsContext) Default_value() IDefault_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefault_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefault_valueContext)
}

func (s *Encoding_flagsContext) OPTIONAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserOPTIONAL_LITERAL, 0)
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
	p.EnterRule(localctx, 6, ASN1SchemaParserRULE_encoding_flags)

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

	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserDEFAULT_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Default_value()
		}

	case ASN1SchemaParserOPTIONAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Match(ASN1SchemaParserOPTIONAL_LITERAL)
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
	p.EnterRule(localctx, 8, ASN1SchemaParserRULE_default_value)

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
		p.SetState(36)
		p.Match(ASN1SchemaParserDEFAULT_LITERAL)
	}
	{
		p.SetState(37)
		p.Value_identifier()
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
	p.EnterRule(localctx, 10, ASN1SchemaParserRULE_value_identifier)

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
		p.Match(ASN1SchemaParserWORD)
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
	p.EnterRule(localctx, 12, ASN1SchemaParserRULE_type_definition)
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserIMPLICIT_LITERAL {
		{
			p.SetState(41)
			p.Context_flag()
		}

	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ASN1SchemaParserSEQUENCE_LITERAL, ASN1SchemaParserINTEGER_LITERAL:
		{
			p.SetState(44)
			p.Primitive()
		}

	case ASN1SchemaParserWORD:
		{
			p.SetState(45)
			p.Type_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserOPTIONAL_LITERAL || _la == ASN1SchemaParserDEFAULT_LITERAL {
		{
			p.SetState(48)
			p.Encoding_flags()
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
	p.EnterRule(localctx, 14, ASN1SchemaParserRULE_type_name)

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
		p.SetState(51)
		p.Match(ASN1SchemaParserWORD)
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
	p.EnterRule(localctx, 16, ASN1SchemaParserRULE_sequence_definition)

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
		p.SetState(53)
		p.Match(ASN1SchemaParserLBRACE)
	}
	{
		p.SetState(54)
		p.Field_list()
	}
	{
		p.SetState(55)
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

func (s *Field_listContext) AllField_definition() []IField_definitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IField_definitionContext)(nil)).Elem())
	var tst = make([]IField_definitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IField_definitionContext)
		}
	}

	return tst
}

func (s *Field_listContext) Field_definition(i int) IField_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_definitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IField_definitionContext)
}

func (s *Field_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ASN1SchemaParserCOMMA, 0)
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
	p.EnterRule(localctx, 18, ASN1SchemaParserRULE_field_list)
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
		p.SetState(57)
		p.Field_definition()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserCOMMA {
		{
			p.SetState(58)
			p.Match(ASN1SchemaParserCOMMA)
		}
		{
			p.SetState(59)
			p.Field_definition()
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
	p.EnterRule(localctx, 20, ASN1SchemaParserRULE_field_definition)
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
		p.SetState(62)
		p.Field_name()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ASN1SchemaParserLBRACKET {
		{
			p.SetState(63)
			p.Tag()
		}

	}
	{
		p.SetState(66)
		p.Type_definition()
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
	p.EnterRule(localctx, 22, ASN1SchemaParserRULE_field_name)

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
		p.SetState(68)
		p.Match(ASN1SchemaParserWORD)
	}

	return localctx
}
