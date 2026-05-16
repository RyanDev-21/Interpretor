package evaluator

import (
	"fmt"
	"strings"

	"github.com/RyanDev-21/ast"
	"github.com/RyanDev-21/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.NULLobj{}
)

const PENDAS = "+-*/"

var builtins = map[string]*object.BuiltIn{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
	"includes": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments.got=%d, want=2", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				if args[1].Type() == object.ObjStr {
					arg2 := args[1].(*object.String)
					return &object.Boolean{Value: strings.Contains(arg.Value, arg2.Value)}
				} else {
					return newError("wrong type for second arguments ,got=%v ,want=STRING", args[1].Type())
				}
			default:
				return newError("wrong type for first arguments, got=%v,want=STRING", arg.Type())
			}
		},
	},
}

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolObj(node.Value)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ast.PrefixExpression: //!!false
		// fist iteration will be right = ! ,operator = !
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return evalBlockStatements(node.Statements, env)
	case *ast.IfExpression:
		return evalIfStatement(node, env)
	case *ast.ReturnStatement:
		value := Eval(node.ReturnValue, env)
		if isError(value) {
			return value
		}
		return &object.ReturnValue{Value: value}
	case *ast.FunctionLiteral:
		params := node.Params
		body := node.Body
		return &object.Function{Params: params, Body: body, Env: env}
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpression(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	}
	return nil
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedenv := extendNewEnv(fn, args)
		evaluated := Eval(fn.Body, extendedenv)
		return unwrapReturnValue(evaluated)
	case *object.BuiltIn:
		return fn.Fn(args...)
	default:
		return newError("not a function : %s", fn.Type())
	}
}

func unwrapReturnValue(eval object.Object) object.Object {
	if val, ok := eval.(*object.ReturnValue); ok {
		return val
	}
	return eval
}

func extendNewEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Params {
		env.Set(param.Value, args[paramIdx])
	}
	return env
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		val, ok := builtins[node.Value]
		if !ok {
			return newError("identifier not found: %s", node.Value)
		}
		return val
	}
	return val
}

func evalExpression(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, exp := range exps {
		eval := Eval(exp, env)
		if isError(eval) {
			return []object.Object{eval}
		}
		result = append(result, eval)
	}
	return result
}

func evalProgram(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func evalBlockStatements(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()
			if rt == object.ObjReturnValue || rt == object.ObjError {
				return result
			}
		}
	}
	return result
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func evalIfStatement(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
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
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	case left.Type() == object.ObjNum && right.Type() == object.ObjNum:
		return evalNumberInfixExpression(
			operator,
			left.(object.Number),
			right.(object.Number),
		)
	case left.Type() == object.ObjStr && right.Type() == object.ObjStr && operator == "+":
		return evalStringInfixExpression(
			left,
			right,
		)
	case operator == "==":
		return nativeBoolObj(left == right)
	case operator == "!=":
		return nativeBoolObj(left != right)

	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())

	}
}

func evalStringInfixExpression(left, right object.Object) object.Object {
	leftStr := left.(*object.String).Value
	rightStr := right.(*object.String).Value
	return &object.String{Value: leftStr + rightStr}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ObjError
	}
	return false
}

func evalNumberInfixExpression(operator string, left, right object.Number) object.Object {
	switch {
	case left.NumberType() != right.NumberType():
		return newError("number type mismatch:%s %s", left.NumberType(), right.NumberType())
	case left.NumberType() == object.ObjInt && right.NumberType() == object.ObjInt:
		leftValue := left.(*object.Integer).Value
		rightValue := right.(*object.Integer).Value
		return evalIntegerValue(operator, leftValue, rightValue)
	case left.NumberType() == object.ObjFloat && right.NumberType() == object.ObjFloat:
		leftValue := left.(*object.Float).Value
		rightValue := right.(*object.Float).Value
		return evalFloatValue(operator, leftValue, rightValue)
	default:
		return newError("unknown type :%s %s", left.NumberType(), right.NumberType())
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
		return newError("unknown operator:%s", operator)
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
		return newError("unknown operator:%s", operator)
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
	if right.Type() == object.ObjNum {
		switch objType := right.(object.Number); objType.NumberType() {
		case object.ObjInt:
			value := right.(*object.Integer).Value
			return &object.Integer{Value: -value}
		case object.ObjFloat:
			value := right.(*object.Float).Value
			return &object.Float{Value: -value}
		}
	}
	return newError("unknown operator: -%s", right.Type())
}
