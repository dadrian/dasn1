// Code generated from ASN1Schema.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ASN1Schema

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ASN1SchemaListener is a complete listener for a parse tree produced by ASN1SchemaParser.
type ASN1SchemaListener interface {
	antlr.ParseTreeListener

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrimitive is called when entering the primitive production.
	EnterPrimitive(c *PrimitiveContext)

	// EnterPrimitive_name is called when entering the primitive_name production.
	EnterPrimitive_name(c *Primitive_nameContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterContext_flag is called when entering the context_flag production.
	EnterContext_flag(c *Context_flagContext)

	// EnterEncoding_flags is called when entering the encoding_flags production.
	EnterEncoding_flags(c *Encoding_flagsContext)

	// EnterDefault_value is called when entering the default_value production.
	EnterDefault_value(c *Default_valueContext)

	// EnterSize_constraint is called when entering the size_constraint production.
	EnterSize_constraint(c *Size_constraintContext)

	// EnterType_constraints is called when entering the type_constraints production.
	EnterType_constraints(c *Type_constraintsContext)

	// EnterType_definition is called when entering the type_definition production.
	EnterType_definition(c *Type_definitionContext)

	// EnterInteger_integer is called when entering the integer_integer production.
	EnterInteger_integer(c *Integer_integerContext)

	// EnterInteger_enum_value_list is called when entering the integer_enum_value_list production.
	EnterInteger_enum_value_list(c *Integer_enum_value_listContext)

	// EnterInteger_enum_value is called when entering the integer_enum_value production.
	EnterInteger_enum_value(c *Integer_enum_valueContext)

	// EnterOid is called when entering the oid production.
	EnterOid(c *OidContext)

	// EnterOctet_string is called when entering the octet_string production.
	EnterOctet_string(c *Octet_stringContext)

	// EnterBit_string is called when entering the bit_string production.
	EnterBit_string(c *Bit_stringContext)

	// EnterString_string is called when entering the string_string production.
	EnterString_string(c *String_stringContext)

	// EnterSequence is called when entering the sequence production.
	EnterSequence(c *SequenceContext)

	// EnterSequence_list is called when entering the sequence_list production.
	EnterSequence_list(c *Sequence_listContext)

	// EnterSequence_of is called when entering the sequence_of production.
	EnterSequence_of(c *Sequence_ofContext)

	// EnterChoice is called when entering the choice production.
	EnterChoice(c *ChoiceContext)

	// EnterField_list is called when entering the field_list production.
	EnterField_list(c *Field_listContext)

	// EnterField_definition is called when entering the field_definition production.
	EnterField_definition(c *Field_definitionContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterValue_identifier is called when entering the value_identifier production.
	EnterValue_identifier(c *Value_identifierContext)

	// EnterField_name is called when entering the field_name production.
	EnterField_name(c *Field_nameContext)

	// EnterCustom_type_name is called when entering the custom_type_name production.
	EnterCustom_type_name(c *Custom_type_nameContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrimitive is called when exiting the primitive production.
	ExitPrimitive(c *PrimitiveContext)

	// ExitPrimitive_name is called when exiting the primitive_name production.
	ExitPrimitive_name(c *Primitive_nameContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitContext_flag is called when exiting the context_flag production.
	ExitContext_flag(c *Context_flagContext)

	// ExitEncoding_flags is called when exiting the encoding_flags production.
	ExitEncoding_flags(c *Encoding_flagsContext)

	// ExitDefault_value is called when exiting the default_value production.
	ExitDefault_value(c *Default_valueContext)

	// ExitSize_constraint is called when exiting the size_constraint production.
	ExitSize_constraint(c *Size_constraintContext)

	// ExitType_constraints is called when exiting the type_constraints production.
	ExitType_constraints(c *Type_constraintsContext)

	// ExitType_definition is called when exiting the type_definition production.
	ExitType_definition(c *Type_definitionContext)

	// ExitInteger_integer is called when exiting the integer_integer production.
	ExitInteger_integer(c *Integer_integerContext)

	// ExitInteger_enum_value_list is called when exiting the integer_enum_value_list production.
	ExitInteger_enum_value_list(c *Integer_enum_value_listContext)

	// ExitInteger_enum_value is called when exiting the integer_enum_value production.
	ExitInteger_enum_value(c *Integer_enum_valueContext)

	// ExitOid is called when exiting the oid production.
	ExitOid(c *OidContext)

	// ExitOctet_string is called when exiting the octet_string production.
	ExitOctet_string(c *Octet_stringContext)

	// ExitBit_string is called when exiting the bit_string production.
	ExitBit_string(c *Bit_stringContext)

	// ExitString_string is called when exiting the string_string production.
	ExitString_string(c *String_stringContext)

	// ExitSequence is called when exiting the sequence production.
	ExitSequence(c *SequenceContext)

	// ExitSequence_list is called when exiting the sequence_list production.
	ExitSequence_list(c *Sequence_listContext)

	// ExitSequence_of is called when exiting the sequence_of production.
	ExitSequence_of(c *Sequence_ofContext)

	// ExitChoice is called when exiting the choice production.
	ExitChoice(c *ChoiceContext)

	// ExitField_list is called when exiting the field_list production.
	ExitField_list(c *Field_listContext)

	// ExitField_definition is called when exiting the field_definition production.
	ExitField_definition(c *Field_definitionContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitValue_identifier is called when exiting the value_identifier production.
	ExitValue_identifier(c *Value_identifierContext)

	// ExitField_name is called when exiting the field_name production.
	ExitField_name(c *Field_nameContext)

	// ExitCustom_type_name is called when exiting the custom_type_name production.
	ExitCustom_type_name(c *Custom_type_nameContext)
}
