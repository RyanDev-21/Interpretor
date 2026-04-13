package parser

import (
	"fmt"
	"testing"

	"github.com/RyanDev-21/ast"
	"github.com/RyanDev-21/lexer"
)

//	func TestLetStatements(t *testing.T) {
//		// t.Run("testing for letStatement", func(t *testing.T) {
//		// 	input := `
//		// let x  5;
//		// let  = 10;
//		// let foobar ;
//		// `
//		// 	l := lexer.New(input)
//		// 	p := New(l)
//		// 	program := p.ParseProgram()
//		// 	checkParserErrors(t, p)
//		// 	if program == nil {
//		// 		t.Fatalf("parseprogram() returned nil")
//		// 	}
//		// 	if len(program.Statements) != 3 {
//		// 		t.Fatalf("program.statements does not contain 3 statements. got=%d",
//		// 			len(program.Statements))
//		// 	}
//		// 	tests := []struct {
//		// 		expectedidentifier string
//		// 	}{
//		// 		{"x"},
//		// 		{"y"},
//		// 		{"foobar"},
//		// 	}
//		// 	for i, tt := range tests {
//		// 		stmt := program.Statements[i]
//		// 		if !testLetStatement(t, stmt, tt.expectedidentifier) {
//		// 			return
//		// 		}
//		// 	}
//		// })
//		t.Run("testing with returnStatement", func(t *testing.T) {
//			input := `
//		return 5;
//		return 10+5;
//		return add(10,4);
//		`
//			l := lexer.New(input)
//			p := New(l)
//			program := p.ParseProgram()
//			checkParserErrors(t, p)
//			if program == nil {
//				t.Fatalf("parseprogram() returned nil")
//			}
//			if len(program.Statements) != 3 {
//				t.Fatalf("program.statements does not contain 3 statements. got=%d",
//					len(program.Statements))
//			}
//
//			for _, stmt := range program.Statements {
//				returnStmt, ok := stmt.(*ast.ReturnStatement)
//				if !ok {
//					t.Fatalf("stmt not *ast.ReturnStatement got %T", stmt)
//					continue
//				}
//				if returnStmt.TokenLiteral() != "return" {
//					t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
//						returnStmt.TokenLiteral())
//				}
//			}
//		})
//	}
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func TestExpressionStatment(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program statements not match %d ", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statements doesn't statisfy the ExpressionStatement got %v", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("statements doesn't statisfy the ast.Identifier got %v", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Fatalf("ident values doesn't match got %v expected:%v", ident.Value, "foobar")
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func TestIntegerExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program statements not match %d ", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statements doesn't statisfy the ExpressionStatement got %v", program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error:%q", msg)
	}
	t.FailNow()
}
