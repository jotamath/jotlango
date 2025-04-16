package eval

import (
	"fmt"
	"jotlango/internal/object"
	"strings"
)

var builtins = map[string]*object.Builtin{
	"print": {
		Fn: func(args ...object.Object) object.Object {
			var output []string
			for _, arg := range args {
				output = append(output, arg.Inspect())
			}
			fmt.Println(strings.Join(output, " "))
			return object.NULL
		},
	},
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Number{Value: float64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
}
