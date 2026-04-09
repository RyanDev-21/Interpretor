package ast

import "github.com/RyanDev-21/token"

type Node interface {
	TokenLitreal() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLitreal()
	}
	return ""
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode()
func (l *LetStatement) TokenLiteral() string { return l.Value.TokenLitreal() }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
