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
OF_LITERAL: 'OF';
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
assignment: type_name ASSIGN primitive;

primitive: sequence | integer_integer | oid | octet_string | bit_string | string_string;
primitive_name: SEQUENCE_LITERAL | INTEGER_LITERAL | OBJECT_IDENTIFIER_LITERAL | OCTET_STRING_LITERAL | BIT_STRING_LITERAL | string_literal;

tag: LBRACKET INTEGER RBRACKET;
context_flag: IMPLICIT_LITERAL;
encoding_flags: OPTIONAL_LITERAL | default_value;
default_value: DEFAULT_LITERAL value_identifier;

type_definition: context_flag? (primitive | type_name) encoding_flags?;

integer_integer: INTEGER_LITERAL;
oid: OBJECT_IDENTIFIER_LITERAL;
octet_string: OCTET_STRING_LITERAL;
bit_string: BIT_STRING_LITERAL;
string_string: string_literal;

sequence: sequence_list | sequence_of;
sequence_list: SEQUENCE_LITERAL LBRACE field_list RBRACE;
sequence_of: SEQUENCE_LITERAL OF_LITERAL type_name;
field_list: field_definition (COMMA field_list)?;
field_definition: field_name tag? type_definition;

string_literal: BMPSTRING_LITERAL | IA5STRING_LITERAL | PRINTABLE_STRING_LITERAL | UTF8_STRING_LITERAL;

type_name: primitive_name | custom_type_name;

field_name: WORD;
custom_type_name: WORD;
value_identifier: WORD;

