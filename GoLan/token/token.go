package token

import (
	"strings"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT    = "IDENT"
	INT      = "INT"
	FLOAT    = "FLOAT"
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	EQUAL    = "=="
	NEQUAL   = "!="

	LT        = "<"
	GT        = ">"
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	FALSE    = "FALSE"
	TRUE     = "TRUE"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"else":   ELSE,
	"return": RETURN,
	"if":     IF,
}

func LookUpIdent(ident string) TokenType {
	if v, ok := keywords[ident]; ok {
		return v
	}
	return IDENT
}

func LoopUpNumberType(number string) TokenType {
	if strings.Contains(number, ".") {
		n := strings.Count(number, ".")
		if n > 1 {
			return ILLEGAL
		}
		return FLOAT
	}
	return INT
}
