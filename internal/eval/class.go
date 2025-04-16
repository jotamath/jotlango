package eval

import (
	"bytes"
	"fmt"
	"jotlango/internal/object"
)

type Class struct {
	Name       string
	Properties map[string]object.Object
	Methods    map[string]*object.Function
}

func (c *Class) Type() object.ObjectType { return "CLASS" }
func (c *Class) Inspect() string {
	var out bytes.Buffer
	out.WriteString("class " + c.Name + " {")
	for name := range c.Properties {
		out.WriteString(fmt.Sprintf("\n  prop %s", name))
	}
	for name := range c.Methods {
		out.WriteString(fmt.Sprintf("\n  fn %s", name))
	}
	out.WriteString("\n}")
	return out.String()
}

func NewClass(name string) *Class {
	return &Class{
		Name:       name,
		Properties: make(map[string]object.Object),
		Methods:    make(map[string]*object.Function),
	}
}
