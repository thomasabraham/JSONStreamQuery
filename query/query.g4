// query.g4
grammar query;

// NOTE: Parser rule names must start with a lowercase letter 
// and lexer rules must start with a capital letter.

// Tokens; lexer rules
SELECT: 'SELECT';
STAR: '*';
FROM: 'FROM';
STREAM: 'STREAM';
WS: [ \n\r\t]+ -> skip;

// Rules; parser rules

start : statement EOF;

statement: SELECT STAR FROM STREAM;
