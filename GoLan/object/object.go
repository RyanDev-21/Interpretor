package object

import (
	"fmt"
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
	}
}

type Environment struct {
	store map[string]Object
}

func (e *Environment) Get(name string) (Object, bool) {
	v, ok := e.store[name]
	return v, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
