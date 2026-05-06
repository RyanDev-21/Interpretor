package object

import (
	"fmt"
)

type ObjectType string

const (
	objInt  = "INTEGER"
	objBool = "BOOLEAN"
	objNULL = "NULL"
)

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
	NULL  = &NULLobj{}
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return objInt }

type Float struct {
	Value float32
}

func (i *Float) Inspect() string  { return fmt.Sprintf("%v", i.Value) }
func (i *Float) Type() ObjectType { return objInt }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%v", b.Value) }
func (b *Boolean) Type() ObjectType { return objBool }

type NULLobj struct{}

func (n *NULLobj) Inspect() string  { return "null" }
func (n *NULLobj) Type() ObjectType { return objNULL }
