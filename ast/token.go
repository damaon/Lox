package ast

import "fmt"

type TokenType int
const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE

	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

  
	// One or two character tokens.
	BANG
 	BANG_EQUAL

	EQUAL
 	EQUAL_EQUAL

	GREATER
 	GREATER_EQUAL

	LESS
 	LESS_EQUAL

  
	// Literals.
	IDENTIFIER
	STRING
	NUMBER

  
	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR

	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Token struct {
	Typ TokenType
	Lexeme string
	Literal string
	Line int
}

func (t *Token) ToString() string {	
	return fmt.Sprintf("%v %v %v %d", t.Typ, t.Lexeme, t.Literal, t.Line)
}