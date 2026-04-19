package parser

import (
	"fmt"
	"strconv"

	"github.com/RyanDev-21/ast"
	"github.com/RyanDev-21/lexer"
	"github.com/RyanDev-21/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[token.TokenType]int{
	token.EQUAL:    EQUALS,
	token.NEQUAL:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

type Parser struct {
	l             *lexer.Lexer
	curToken      token.Token
	peekToken     token.Token
	prefixParseFn map[token.TokenType]prefixParseFn
	infixParseFn  map[token.TokenType]infixParseFn
	errors        []string
}
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.prefixParseFn = make(map[token.TokenType]prefixParseFn)
	p.infixParseFn = make(map[token.TokenType]infixParseFn)
	p.registerPrefixParseFn(token.IDENT, p.parseIdentifier)
	p.registerPrefixParseFn(token.INT, p.parseIntLiteral)
	p.registerPrefixParseFn(token.TRUE, p.parseBoolLiteral)
	p.registerPrefixParseFn(token.FALSE, p.parseBoolLiteral)
	p.registerPrefixParseFn(token.BANG, p.parsePrefixExpression)
	p.registerPrefixParseFn(token.PLUS, p.parsePrefixExpression)
	p.registerPrefixParseFn(token.MINUS, p.parsePrefixExpression)
	p.registerPrefixParseFn(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefixParseFn(token.IF, p.parseConditionExpression)
	p.registerPrefixParseFn(token.FUNCTION, p.parseFunctionLiteral)

	p.registerInfixParseFn(token.LPAREN, p.parseCallExpression)
	p.registerInfixParseFn(token.SLASH, p.parseInfixExpression)
	p.registerInfixParseFn(token.PLUS, p.parseInfixExpression)
	p.registerInfixParseFn(token.MINUS, p.parseInfixExpression)
	p.registerInfixParseFn(token.EQUAL, p.parseInfixExpression)
	p.registerInfixParseFn(token.NEQUAL, p.parseInfixExpression)
	p.registerInfixParseFn(token.ASTERISK, p.parseInfixExpression)
	p.registerInfixParseFn(token.LT, p.parseInfixExpression)
	p.registerInfixParseFn(token.GT, p.parseInfixExpression)
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) noPrefixFnErrors(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noInfixFnErrors(t token.TokenType) {
	msg := fmt.Sprintf("no infix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerPrefixParseFn(tok token.TokenType, fn prefixParseFn) {
	p.prefixParseFn[tok] = fn
}

func (p *Parser) registerInfixParseFn(tok token.TokenType, fn infixParseFn) {
	p.infixParseFn[tok] = fn
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatements()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatements() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.nextPeek(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// fn(x,y){return x+y}
func (p *Parser) parseFunctionLiteral() ast.Expression {
	fl := &ast.FunctionLiteral{Token: p.curToken}
	if !p.nextPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	fl.Params = p.parseParams()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	fl.Body = p.parseBlockStatement()
	return fl
}

// (x,y,z) || ()
func (p *Parser) parseParams() []*ast.Identifier {
	var params []*ast.Identifier
	if p.nextPeek(token.RPAREN) {
		p.nextToken()
		return params
	}
	p.nextToken()
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	params = append(params, ident)

	for p.nextPeek(token.COMMA) {
		p.nextToken()                                                         // here we reach the comma  token
		p.nextToken()                                                         // here we reach the next token
		ident = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} // create the ident
		params = append(params, ident)
	}
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return params
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}

// add(1+2,2+3) || add()
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	callExp := &ast.CallExpression{Token: p.curToken, Function: function}
	callExp.Arguments = p.parseArguments()
	return callExp
}

func (p *Parser) parseArguments() []ast.Expression {
	args := []ast.Expression{}
	if p.nextPeek(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))
	for p.nextPeek(token.COMMA) {
		p.nextToken() // same stuff with if params
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return args
}

func (p *Parser) parseConditionExpression() ast.Expression {
	exp := &ast.IfExpression{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	exp.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	exp.Consequence = p.parseBlockStatement()
	if p.nextPeek(token.ELSE) {
		p.nextToken()
		if !p.expectPeek(token.LBRACE) {
			return nil
		}
		exp.Alternative = p.parseBlockStatement()
	}
	return exp
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatements()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()

	}
	return block
}

func (p *Parser) parseIntLiteral() ast.Expression {
	expression := &ast.IntegerLiteral{
		Token: p.curToken,
	}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("coult not parse the string into int%s", err)
		p.errors = append(p.errors, msg)
		return nil
	}

	expression.Value = value
	return expression
}

func (p *Parser) parseBoolLiteral() ast.Expression {
	expression := &ast.Boolean{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE),
	}
	return expression
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedece := p.curPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedece)
	return exp
}

// add(10,20)
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFn[p.curToken.Type]
	if prefix == nil {
		p.noPrefixFnErrors(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	for !p.nextPeek(token.SEMICOLON) && precedence < p.peekPrecedences() {
		infix := p.infixParseFn[p.peekToken.Type]
		if infix == nil {
			p.noInfixFnErrors(p.peekToken.Type)
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.nextPeek(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected token to be %s ,got %s instead ", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) nextPeek(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekPrecedences() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}
