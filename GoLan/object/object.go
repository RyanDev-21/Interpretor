package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/RyanDev-21/ast"
)

type ObjectType string

const (
	ObjInt         = "INTEGER"
	ObjFloat       = "FLOAT"
	ObjBool        = "BOOLEAN"
	ObjNULL        = "NULL"
	ObjNum         = "NUMBER"
	ObjReturnValue = "RETURN"
	ObjError       = "ERROR"
	ObjFunction    = "FUNCTOIN"
	ObjStr         = "STRING"
	ObjBuiltIn     = "BUILTIN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
type Number interface {
	NumberType() ObjectType
}

type Error struct {
	Message string
}

func (e *Error) Inspect() string  { return fmt.Sprint("Error: " + e.Message) }
func (e *Error) Type() ObjectType { return ObjError }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string  { return fmt.Sprintf("%v", rv.Value) }
func (rv *ReturnValue) Type() ObjectType { return ObjReturnValue }

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string        { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) NumberType() ObjectType { return ObjInt }
func (i *Integer) Type() ObjectType       { return ObjNum }

type Float struct {
	Value float32
}

func (i *Float) Inspect() string        { return fmt.Sprintf("%v", i.Value) }
func (i *Float) NumberType() ObjectType { return ObjFloat }
func (i *Float) Type() ObjectType       { return ObjNum }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%v", b.Value) }
func (b *Boolean) Type() ObjectType { return ObjBool }

type NULLobj struct{}

func (n *NULLobj) Inspect() string  { return "null" }
func (n *NULLobj) Type() ObjectType { return ObjNULL }

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{
		store: s,
		outer: nil,
	}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	v, ok := e.store[name]
	if !ok && e.outer != nil {
		v, ok = e.outer.Get(name)
	}
	return v, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

type Function struct {
	Params []*ast.Identifier
	Body   *ast.BlockStatement
	Env    *Environment
}

func (f *Function) Type() ObjectType { return ObjFunction }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Params {
		params = append(params, p.Value)
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString("){\n")
	out.WriteString(f.Body.String())
	out.WriteString("}\n")
	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return ObjStr }
func (s *String) Inspect() string {
	return s.Value
}

type (
	BuiltInFunction func(args ...Object) Object
	BuiltIn         struct {
		Fn BuiltInFunction
	}
)

func (b *BuiltIn) Type() ObjectType { return ObjBuiltIn }
func (b *BuiltIn) Inspect() string {
	return "builtIn function"
}
