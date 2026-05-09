package object

import (
	"fmt"
)

type ObjectType string

const (
	ObjInt   = "INTEGER"
	ObjFloat = "FLOAT"
	ObjBool  = "BOOLEAN"
	ObjNULL  = "NULL"
	ObjNum   = "NUMBER"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
type Number interface {
	NumberType() ObjectType
}

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
