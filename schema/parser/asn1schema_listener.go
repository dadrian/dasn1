// Code generated from ASN1Schema.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ASN1Schema

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ASN1SchemaListener is a complete listener for a parse tree produced by ASN1SchemaParser.
type ASN1SchemaListener interface {
	antlr.ParseTreeListener

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterPrimitive is called when entering the primitive production.
	EnterPrimitive(c *PrimitiveContext)

	// EnterContext_flag is called when entering the context_flag production.
	EnterContext_flag(c *Context_flagContext)

	// EnterEncoding_flags is called when entering the encoding_flags production.
	EnterEncoding_flags(c *Encoding_flagsContext)

	// EnterDefault_value is called when entering the default_value production.
	EnterDefault_value(c *Default_valueContext)

	// EnterValue_identifier is called when entering the value_identifier production.
	EnterValue_identifier(c *Value_identifierContext)

	// EnterType_definition is called when entering the type_definition production.
	EnterType_definition(c *Type_definitionContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterSequence_definition is called when entering the sequence_definition production.
	EnterSequence_definition(c *Sequence_definitionContext)

	// EnterField_list is called when entering the field_list production.
	EnterField_list(c *Field_listContext)

	// EnterField_definition is called when entering the field_definition production.
	EnterField_definition(c *Field_definitionContext)

	// EnterField_name is called when entering the field_name production.
	EnterField_name(c *Field_nameContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitPrimitive is called when exiting the primitive production.
	ExitPrimitive(c *PrimitiveContext)

	// ExitContext_flag is called when exiting the context_flag production.
	ExitContext_flag(c *Context_flagContext)

	// ExitEncoding_flags is called when exiting the encoding_flags production.
	ExitEncoding_flags(c *Encoding_flagsContext)

	// ExitDefault_value is called when exiting the default_value production.
	ExitDefault_value(c *Default_valueContext)

	// ExitValue_identifier is called when exiting the value_identifier production.
	ExitValue_identifier(c *Value_identifierContext)

	// ExitType_definition is called when exiting the type_definition production.
	ExitType_definition(c *Type_definitionContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitSequence_definition is called when exiting the sequence_definition production.
	ExitSequence_definition(c *Sequence_definitionContext)

	// ExitField_list is called when exiting the field_list production.
	ExitField_list(c *Field_listContext)

	// ExitField_definition is called when exiting the field_definition production.
	ExitField_definition(c *Field_definitionContext)

	// ExitField_name is called when exiting the field_name production.
	ExitField_name(c *Field_nameContext)
}
