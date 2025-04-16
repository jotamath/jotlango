package eval

import (
	"bytes"
	"fmt"
	"jotlango/internal/object"
)

type Instance struct {
	Class      *Class
	Properties map[string]object.Object
}

func (i *Instance) Type() object.ObjectType { return "INSTANCE" }
func (i *Instance) Inspect() string {
	var out bytes.Buffer
	out.WriteString(i.Class.Name + " {")
	for name, value := range i.Properties {
		out.WriteString(fmt.Sprintf("\n  %s: %s", name, value.Inspect()))
	}
	out.WriteString("\n}")
	return out.String()
}

func NewInstance(class *Class) *Instance {
	instance := &Instance{
		Class:      class,
		Properties: make(map[string]object.Object),
	}

	// Copiar propriedades da classe para a instância
	for name, value := range class.Properties {
		instance.Properties[name] = value
	}

	return instance
}
