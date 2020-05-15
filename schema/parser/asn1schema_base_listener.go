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

// EnterModule is called when production module is entered.
func (s *BaseASN1SchemaListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseASN1SchemaListener) ExitModule(ctx *ModuleContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseASN1SchemaListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseASN1SchemaListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterPrimitive is called when production primitive is entered.
func (s *BaseASN1SchemaListener) EnterPrimitive(ctx *PrimitiveContext) {}

// ExitPrimitive is called when production primitive is exited.
func (s *BaseASN1SchemaListener) ExitPrimitive(ctx *PrimitiveContext) {}

// EnterPrimitive_name is called when production primitive_name is entered.
func (s *BaseASN1SchemaListener) EnterPrimitive_name(ctx *Primitive_nameContext) {}

// ExitPrimitive_name is called when production primitive_name is exited.
func (s *BaseASN1SchemaListener) ExitPrimitive_name(ctx *Primitive_nameContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseASN1SchemaListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseASN1SchemaListener) ExitTag(ctx *TagContext) {}

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

// EnterType_definition is called when production type_definition is entered.
func (s *BaseASN1SchemaListener) EnterType_definition(ctx *Type_definitionContext) {}

// ExitType_definition is called when production type_definition is exited.
func (s *BaseASN1SchemaListener) ExitType_definition(ctx *Type_definitionContext) {}

// EnterInteger_integer is called when production integer_integer is entered.
func (s *BaseASN1SchemaListener) EnterInteger_integer(ctx *Integer_integerContext) {}

// ExitInteger_integer is called when production integer_integer is exited.
func (s *BaseASN1SchemaListener) ExitInteger_integer(ctx *Integer_integerContext) {}

// EnterOid is called when production oid is entered.
func (s *BaseASN1SchemaListener) EnterOid(ctx *OidContext) {}

// ExitOid is called when production oid is exited.
func (s *BaseASN1SchemaListener) ExitOid(ctx *OidContext) {}

// EnterOctet_string is called when production octet_string is entered.
func (s *BaseASN1SchemaListener) EnterOctet_string(ctx *Octet_stringContext) {}

// ExitOctet_string is called when production octet_string is exited.
func (s *BaseASN1SchemaListener) ExitOctet_string(ctx *Octet_stringContext) {}

// EnterBit_string is called when production bit_string is entered.
func (s *BaseASN1SchemaListener) EnterBit_string(ctx *Bit_stringContext) {}

// ExitBit_string is called when production bit_string is exited.
func (s *BaseASN1SchemaListener) ExitBit_string(ctx *Bit_stringContext) {}

// EnterString_string is called when production string_string is entered.
func (s *BaseASN1SchemaListener) EnterString_string(ctx *String_stringContext) {}

// ExitString_string is called when production string_string is exited.
func (s *BaseASN1SchemaListener) ExitString_string(ctx *String_stringContext) {}

// EnterSequence is called when production sequence is entered.
func (s *BaseASN1SchemaListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production sequence is exited.
func (s *BaseASN1SchemaListener) ExitSequence(ctx *SequenceContext) {}

// EnterSequence_list is called when production sequence_list is entered.
func (s *BaseASN1SchemaListener) EnterSequence_list(ctx *Sequence_listContext) {}

// ExitSequence_list is called when production sequence_list is exited.
func (s *BaseASN1SchemaListener) ExitSequence_list(ctx *Sequence_listContext) {}

// EnterSequence_of is called when production sequence_of is entered.
func (s *BaseASN1SchemaListener) EnterSequence_of(ctx *Sequence_ofContext) {}

// ExitSequence_of is called when production sequence_of is exited.
func (s *BaseASN1SchemaListener) ExitSequence_of(ctx *Sequence_ofContext) {}

// EnterField_list is called when production field_list is entered.
func (s *BaseASN1SchemaListener) EnterField_list(ctx *Field_listContext) {}

// ExitField_list is called when production field_list is exited.
func (s *BaseASN1SchemaListener) ExitField_list(ctx *Field_listContext) {}

// EnterField_definition is called when production field_definition is entered.
func (s *BaseASN1SchemaListener) EnterField_definition(ctx *Field_definitionContext) {}

// ExitField_definition is called when production field_definition is exited.
func (s *BaseASN1SchemaListener) ExitField_definition(ctx *Field_definitionContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseASN1SchemaListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseASN1SchemaListener) ExitString_literal(ctx *String_literalContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseASN1SchemaListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaseASN1SchemaListener) ExitType_name(ctx *Type_nameContext) {}

// EnterField_name is called when production field_name is entered.
func (s *BaseASN1SchemaListener) EnterField_name(ctx *Field_nameContext) {}

// ExitField_name is called when production field_name is exited.
func (s *BaseASN1SchemaListener) ExitField_name(ctx *Field_nameContext) {}

// EnterCustom_type_name is called when production custom_type_name is entered.
func (s *BaseASN1SchemaListener) EnterCustom_type_name(ctx *Custom_type_nameContext) {}

// ExitCustom_type_name is called when production custom_type_name is exited.
func (s *BaseASN1SchemaListener) ExitCustom_type_name(ctx *Custom_type_nameContext) {}

// EnterValue_identifier is called when production value_identifier is entered.
func (s *BaseASN1SchemaListener) EnterValue_identifier(ctx *Value_identifierContext) {}

// ExitValue_identifier is called when production value_identifier is exited.
func (s *BaseASN1SchemaListener) ExitValue_identifier(ctx *Value_identifierContext) {}
