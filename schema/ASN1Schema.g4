grammar ASN1Schema;

// Tokens
fragment SPACE: '\t' | ' ' | '\r' | '\n'| '\u000C';
ASSIGN: '::=';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
COMMA: ',';
SEQUENCE_LITERAL: 'SEQUENCE';
INTEGER_LITERAL: 'INTEGER';
IMPLICIT_LITERAL: 'IMPLICIT';
OPTIONAL_LITERAL: 'OPTIONAL';
DEFAULT_LITERAL: 'DEFAULT';
BMPSTRING_LITERAL: 'BMPString';
IA5STRING_LITERAL: 'IA5String';
PRINTABLE_STRING_LITERAL: 'PrintableString';
UTF8_STRING_LITERAL: 'UTF8String';
BIT_STRING_LITERAL: 'BIT' SPACE+ 'STRING';
OCTET_STRING_LITERAL: 'OCTET' SPACE+ 'STRING';
OBJECT_IDENTIFIER_LITERAL: 'OBJECT' SPACE+ 'IDENTIFIER';
INTEGER: ('0'..'9')+;
WORD: ('a'..'z'|'A'..'Z') ('0'..'9'|'a'..'z'|'A'..'Z')* ;
WHITESPACE: SPACE+ -> skip;

// Rules
module: assignment+ EOF;
assignment: type_name ASSIGN primitive primitive_definition;

primitive: SEQUENCE_LITERAL | OBJECT_IDENTIFIER_LITERAL | INTEGER_LITERAL | OCTET_STRING_LITERAL | BIT_STRING_LITERAL | string_literal;
primitive_definition: sequence_definition;  // | interger_definition, ...

tag: LBRACKET INTEGER RBRACKET;
context_flag: IMPLICIT_LITERAL;
encoding_flags: OPTIONAL_LITERAL | default_value;
default_value: DEFAULT_LITERAL value_identifier;

type_definition: context_flag? (primitive | type_name) encoding_flags?;

sequence_definition: LBRACE field_list RBRACE;
field_list: field_definition (COMMA field_list)?;
field_definition: field_name tag? type_definition;

string_literal: BMPSTRING_LITERAL | IA5STRING_LITERAL | PRINTABLE_STRING_LITERAL | UTF8_STRING_LITERAL;

field_name: WORD;
type_name: WORD;
value_identifier: WORD;

