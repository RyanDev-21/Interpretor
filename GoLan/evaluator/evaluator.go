package evaluator

import (
	"github.com/RyanDev-21/ast"
	"github.com/RyanDev-21/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.NULLobj{}
)

const PENDAS = "+-*/"

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolObj(node.Value)
	case *ast.PrefixExpression: //!!false
		// fist iteration will be right = ! ,operator = !
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return evalBlockStatements(node.Statements)
	case *ast.IfExpression:
		return evalIfStatement(node)
	case *ast.ReturnStatement:
		value := Eval(node.ReturnValue)
		return &object.ReturnValue{Value: value}
	}
	return nil
}

func evalProgram(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}
	return result
}

func evalBlockStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
		if result != nil && result.Type() == object.ObjReturnValue {
			return result
		}
	}
	return result
}

func evalIfStatement(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)
	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return NULL
	}
}

func isTruthy(object object.Object) bool {
	switch object {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func nativeBoolObj(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangExpressoin(right)
	case "-":
		return evalMinuxExpressoin(right)
	default:
		return NULL
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.ObjNum && right.Type() == object.ObjNum:
		return evalNumberInfixExpression(
			operator,
			left.(object.Number),
			right.(object.Number),
		)
	case operator == "==":
		return nativeBoolObj(left == right)
	case operator != "!=":
		return nativeBoolObj(left != right)
	default:
		return NULL

	}
}

func evalNumberInfixExpression(operator string, left, right object.Number) object.Object {
	switch {
	case left.NumberType() == object.ObjInt && right.NumberType() == object.ObjInt:
		leftValue := left.(*object.Integer).Value
		rightValue := right.(*object.Integer).Value
		return evalIntegerValue(operator, leftValue, rightValue)
	case left.NumberType() == object.ObjFloat && right.NumberType() == object.ObjFloat:
		leftValue := left.(*object.Float).Value
		rightValue := right.(*object.Float).Value
		return evalFloatValue(operator, leftValue, rightValue)
	default:
		return NULL
	}
}

func evalIntegerValue(operator string, leftValue, rightValue int64) object.Object {
	switch operator {
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}
	case "==":
		return nativeBoolObj(leftValue == rightValue)
	case "!=":
		return nativeBoolObj(leftValue != rightValue)
	case ">":
		return nativeBoolObj(leftValue > rightValue)
	case "<":
		return nativeBoolObj(leftValue < rightValue)
	default:
		return NULL
	}
}

func evalFloatValue(operator string, leftValue, rightValue float32) object.Object {
	switch operator {
	case "-":
		return &object.Float{Value: leftValue - rightValue}
	case "+":
		return &object.Float{Value: leftValue + rightValue}
	case "*":
		return &object.Float{Value: leftValue * rightValue}
	case "/":
		return &object.Float{Value: leftValue / rightValue}
	case "==":
		return nativeBoolObj(leftValue == rightValue)
	case "!=":
		return nativeBoolObj(leftValue != rightValue)
	case ">":
		return nativeBoolObj(leftValue > rightValue)
	case "<":
		return nativeBoolObj(leftValue < rightValue)
	default:
		return NULL
	}
}

func evalBangExpressoin(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return FALSE
	default:
		return FALSE

	}
}

func evalMinuxExpressoin(right object.Object) object.Object {
	if right.Type() == object.ObjNULL {
		return NULL
	}
	switch objType := right.(object.Number); objType.NumberType() {
	case object.ObjInt:
		value := right.(*object.Integer).Value
		return &object.Integer{Value: -value}
	case object.ObjFloat:
		value := right.(*object.Float).Value
		return &object.Float{Value: -value}
	default:
		return NULL
	}
}
