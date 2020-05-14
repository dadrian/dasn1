// Code generated from ASN1Schema.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ASN1Schema

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseASN1SchemaListener is a complete listener for a parse tree produced by ASN1SchemaParser.
type BaseASN1SchemaListener struct{}

var _ ASN1SchemaListener = &BaseASN1SchemaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseASN1SchemaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseASN1SchemaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseASN1SchemaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseASN1SchemaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseASN1SchemaListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseASN1SchemaListener) ExitTag(ctx *TagContext) {}

// EnterPrimitive is called when production primitive is entered.
func (s *BaseASN1SchemaListener) EnterPrimitive(ctx *PrimitiveContext) {}

// ExitPrimitive is called when production primitive is exited.
func (s *BaseASN1SchemaListener) ExitPrimitive(ctx *PrimitiveContext) {}

// EnterContext_flag is called when production context_flag is entered.
func (s *BaseASN1SchemaListener) EnterContext_flag(ctx *Context_flagContext) {}

// ExitContext_flag is called when production context_flag is exited.
func (s *BaseASN1SchemaListener) ExitContext_flag(ctx *Context_flagContext) {}

// EnterEncoding_flags is called when production encoding_flags is entered.
func (s *BaseASN1SchemaListener) EnterEncoding_flags(ctx *Encoding_flagsContext) {}

// ExitEncoding_flags is called when production encoding_flags is exited.
func (s *BaseASN1SchemaListener) ExitEncoding_flags(ctx *Encoding_flagsContext) {}

// EnterDefault_value is called when production default_value is entered.
func (s *BaseASN1SchemaListener) EnterDefault_value(ctx *Default_valueContext) {}

// ExitDefault_value is called when production default_value is exited.
func (s *BaseASN1SchemaListener) ExitDefault_value(ctx *Default_valueContext) {}

// EnterValue_identifier is called when production value_identifier is entered.
func (s *BaseASN1SchemaListener) EnterValue_identifier(ctx *Value_identifierContext) {}

// ExitValue_identifier is called when production value_identifier is exited.
func (s *BaseASN1SchemaListener) ExitValue_identifier(ctx *Value_identifierContext) {}

// EnterType_definition is called when production type_definition is entered.
func (s *BaseASN1SchemaListener) EnterType_definition(ctx *Type_definitionContext) {}

// ExitType_definition is called when production type_definition is exited.
func (s *BaseASN1SchemaListener) ExitType_definition(ctx *Type_definitionContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseASN1SchemaListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaseASN1SchemaListener) ExitType_name(ctx *Type_nameContext) {}

// EnterSequence_definition is called when production sequence_definition is entered.
func (s *BaseASN1SchemaListener) EnterSequence_definition(ctx *Sequence_definitionContext) {}

// ExitSequence_definition is called when production sequence_definition is exited.
func (s *BaseASN1SchemaListener) ExitSequence_definition(ctx *Sequence_definitionContext) {}

// EnterField_list is called when production field_list is entered.
func (s *BaseASN1SchemaListener) EnterField_list(ctx *Field_listContext) {}

// ExitField_list is called when production field_list is exited.
func (s *BaseASN1SchemaListener) ExitField_list(ctx *Field_listContext) {}

// EnterField_definition is called when production field_definition is entered.
func (s *BaseASN1SchemaListener) EnterField_definition(ctx *Field_definitionContext) {}

// ExitField_definition is called when production field_definition is exited.
func (s *BaseASN1SchemaListener) ExitField_definition(ctx *Field_definitionContext) {}

// EnterField_name is called when production field_name is entered.
func (s *BaseASN1SchemaListener) EnterField_name(ctx *Field_nameContext) {}

// ExitField_name is called when production field_name is exited.
func (s *BaseASN1SchemaListener) ExitField_name(ctx *Field_nameContext) {}
