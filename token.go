package main

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

type Token struct {
	typ TokenType
	lexeme string
	literal string // object?
	line int
}

func (t *Token) toString(){	
	fmt.Println("%v %v %v", t.typ, t.lexeme, t.literal)
}