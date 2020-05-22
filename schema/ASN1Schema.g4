grammar ASN1Schema;

// Tokens

fragment SPACE: '\t' | ' ' | '\r' | '\n'| '\u000C';
COMMENT: '--' ~[\n]* -> skip;
ASSIGN: '::=';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
LPAREN: '(';
RPAREN: ')';
COMMA: ',';
SEQUENCE_LITERAL: 'SEQUENCE';
INTEGER_LITERAL: 'INTEGER';
IMPLICIT_LITERAL: 'IMPLICIT';
EXPLICIT_LITERAL: 'EXPLICIT';
OPTIONAL_LITERAL: 'OPTIONAL';
DEFAULT_LITERAL: 'DEFAULT';
DEFINED_BY_LITERAL: 'DEFINED' SPACE+ 'BY';
ANY_LITERAL: 'ANY';
OF_LITERAL: 'OF';
CHOICE_LITERAL: 'CHOICE';
SIZE_LITERAL: 'SIZE';
MAX: 'MAX';
BMPSTRING_LITERAL: 'BMPString';
IA5STRING_LITERAL: 'IA5String';
PRINTABLE_STRING_LITERAL: 'PrintableString';
TELETEX_STRING_LITERAL: 'TeletexString';
UNIVERSAL_STRING_LITERAL: 'UniversalString';
UTF8_STRING_LITERAL: 'UTF8String';
BIT_STRING_LITERAL: 'BIT' SPACE+ 'STRING';
OCTET_STRING_LITERAL: 'OCTET' SPACE+ 'STRING';
OBJECT_IDENTIFIER_LITERAL: 'OBJECT' SPACE+ 'IDENTIFIER';
DOTDOT: '..';
INTEGER: ('0'..'9')+;
WORD: ('a'..'z'|'A'..'Z') ('0'..'9'|'a'..'z'|'A'..'Z'|'-')* ;
WHITESPACE: SPACE+ -> skip;

// Rules
module: assignment+ EOF;
assignment: (type_name ASSIGN primitive) | oid_assignment;

primitive: sequence | choice | integer_integer | oid | octet_string | bit_string | string_string | any;
primitive_name: SEQUENCE_LITERAL | INTEGER_LITERAL | OBJECT_IDENTIFIER_LITERAL | OCTET_STRING_LITERAL | BIT_STRING_LITERAL | string_literal | ANY_LITERAL;

tag: LBRACKET INTEGER RBRACKET;
context_flag: IMPLICIT_LITERAL | EXPLICIT_LITERAL;
encoding_flags: OPTIONAL_LITERAL | default_value;
default_value: DEFAULT_LITERAL value_identifier;

size_constraint: SIZE_LITERAL LPAREN INTEGER DOTDOT (INTEGER | field_name | MAX) RPAREN;
type_constraints: size_constraint;

type_definition: context_flag? (primitive | (type_name (LPAREN type_constraints RPAREN)?)) encoding_flags?;

integer_integer: INTEGER_LITERAL (LBRACE integer_enum_value_list RBRACE)?;
integer_value: LPAREN INTEGER+ RPAREN;

integer_enum_value_list: integer_enum_value (COMMA integer_enum_value_list)?;
integer_enum_value: field_name integer_value;

integer_oid_value_list: (integer_oid_value integer_oid_value_list?) | integer_oid_value INTEGER+;
integer_oid_value: integer_oid_reference | integer_enum_value;

oid: OBJECT_IDENTIFIER_LITERAL;
oid_assignment: field_name OBJECT_IDENTIFIER_LITERAL ASSIGN LBRACE integer_oid_value_list RBRACE;
octet_string: OCTET_STRING_LITERAL type_constraints?;
bit_string: BIT_STRING_LITERAL type_constraints?;
string_string: string_literal type_constraints?;

sequence: sequence_list | sequence_of;
sequence_list: SEQUENCE_LITERAL LBRACE field_list RBRACE;
sequence_of: SEQUENCE_LITERAL size_constraint? OF_LITERAL type_name;

choice: CHOICE_LITERAL LBRACE field_list RBRACE;

any: ANY_LITERAL DEFINED_BY_LITERAL field_name;

field_list: field_definition (COMMA field_list)?;
field_definition: field_name tag? type_definition;

string_literal: BMPSTRING_LITERAL | IA5STRING_LITERAL | PRINTABLE_STRING_LITERAL | TELETEX_STRING_LITERAL | UTF8_STRING_LITERAL | UNIVERSAL_STRING_LITERAL;

type_name: primitive_name | custom_type_name;

integer_oid_reference: field_name;
value_identifier: INTEGER | WORD;
field_name: WORD;
custom_type_name: WORD;

