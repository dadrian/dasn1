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
BIT_LITERAL: 'BIT';
STRING_LITERAL: 'STRING';
BIT_STRING_LITERAL: 'BIT' SPACE 'STRING';
INTEGER: ('0'..'9')+;
WORD: ('a'..'z'|'A'..'Z') ('0'..'9'|'a'..'z'|'A'..'Z')* ;
WHITESPACE: SPACE+ -> skip;

// Rules
module: assignment;
assignment: type_name ASSIGN primitive primitive_definition;

primitive: SEQUENCE_LITERAL | INTEGER_LITERAL | BIT_STRING_LITERAL;
primitive_definition: sequence_definition;  // | interger_definition, ...

tag: LBRACKET INTEGER RBRACKET;
context_flag: IMPLICIT_LITERAL;
encoding_flags: default_value | OPTIONAL_LITERAL;
default_value: DEFAULT_LITERAL value_identifier;
value_identifier: WORD;

type_definition: context_flag? (primitive | type_name) encoding_flags?;
type_name: WORD;

sequence_definition: LBRACE field_list RBRACE;
field_list: field_definition (COMMA field_list)?;
field_definition: field_name tag? type_definition;
field_name: WORD;


