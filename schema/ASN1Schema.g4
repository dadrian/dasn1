grammar ASN1Schema;

// Tokens
WHITESPACE: [ \r\n\t]+ -> skip;
WORD: ('a'..'z'|'A'..'Z') ('0'..'9'|'a'..'z'|'A'..'Z')* ;
INTEGER: ('0'..'9')+;
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

// Rules
tag: LBRACKET INTEGER RBRACKET;
primitive: SEQUENCE_LITERAL | INTEGER_LITERAL;
context_flag: IMPLICIT_LITERAL;
encoding_flags: default_value | OPTIONAL_LITERAL;
default_value: DEFAULT_LITERAL value_identifier;
value_identifier: WORD;

type_definition: context_flag? (primitive | type_name) encoding_flags?;
type_name: WORD;

sequence_definition: LBRACE field_list RBRACE;
field_list: field_definition (COMMA field_definition)?;
field_definition: field_name tag? type_definition;
field_name: WORD;


