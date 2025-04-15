package eval

import (
	"fmt"
	"jotlango/internal/parser"
	"strings"
)

type Object interface {
	Type() string
	Inspect() string
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
		outer: nil,
	}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() string    { return "INTEGER" }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type Float struct {
	Value float64
}

func (f *Float) Type() string    { return "FLOAT" }
func (f *Float) Inspect() string { return fmt.Sprintf("%f", f.Value) }

type String struct {
	Value string
}

func (s *String) Type() string    { return "STRING" }
func (s *String) Inspect() string { return s.Value }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string    { return "BOOLEAN" }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() string    { return "NULL" }
func (n *Null) Inspect() string { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() string    { return "RETURN_VALUE" }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() string    { return "ERROR" }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*parser.Identifier
	Body       *parser.BlockStatement
	Env        *Environment
}

func (f *Function) Type() string { return "FUNCTION" }
func (f *Function) Inspect() string {
	return fmt.Sprintf("fn(%s) {\n%s\n}", f.Parameters, f.Body.String())
}

type Class struct {
	Name       string
	Properties map[string]Object
	Methods    map[string]*Function
}

func (c *Class) Type() string { return "CLASS" }
func (c *Class) Inspect() string {
	return fmt.Sprintf("class %s {\n%s\n}", c.Name, c.Properties)
}

type Evaluator struct {
	env *Environment
}

func NewEvaluator() *Evaluator {
	return &Evaluator{
		env: NewEnvironment(),
	}
}

func (e *Evaluator) Eval(node parser.Node) Object {
	switch node := node.(type) {
	case *parser.Program:
		return e.evalProgram(node)
	case *parser.ExpressionStatement:
		return e.Eval(node.Expression)
	case *parser.IntegerLiteral:
		return &Integer{Value: node.Value}
	case *parser.FloatLiteral:
		return &Float{Value: node.Value}
	case *parser.StringLiteral:
		return &String{Value: node.Value}
	case *parser.Boolean:
		return &Boolean{Value: node.Value}
	case *parser.PrefixExpression:
		right := e.Eval(node.Right)
		if isError(right) {
			return right
		}
		return e.evalPrefixExpression(node.Operator, right)
	case *parser.InfixExpression:
		left := e.Eval(node.Left)
		if isError(left) {
			return left
		}
		right := e.Eval(node.Right)
		if isError(right) {
			return right
		}
		return e.evalInfixExpression(node.Operator, left, right)
	case *parser.BlockStatement:
		return e.evalBlockStatement(node)
	case *parser.IfExpression:
		return e.evalIfExpression(node)
	case *parser.ReturnStatement:
		val := e.Eval(node.ReturnValue)
		if isError(val) {
			return val
		}
		return &ReturnValue{Value: val}
	case *parser.LetStatement:
		val := e.Eval(node.Value)
		if isError(val) {
			return val
		}
		e.env.Set(node.Name.Value, val)
	case *parser.Identifier:
		return e.evalIdentifier(node)
	case *parser.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &Function{Parameters: params, Body: body, Env: e.env}
	case *parser.CallExpression:
		function := e.Eval(node.Function)
		if isError(function) {
			return function
		}
		args := e.evalExpressions(node.Arguments)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return e.applyFunction(function, args)
	case *parser.ClassStatement:
		return e.evalClassStatement(node)
	case *parser.InstanceLiteral:
		return e.evalInstanceLiteral(node)
	case *parser.PropertyAccess:
		return e.evalPropertyAccess(node)
	}
	return nil
}

func (e *Evaluator) evalProgram(program *parser.Program) Object {
	var result Object
	for _, statement := range program.Statements {
		result = e.Eval(statement)
		switch result := result.(type) {
		case *ReturnValue:
			return result.Value
		case *Error:
			return result
		}
	}
	return result
}

func (e *Evaluator) evalBlockStatement(block *parser.BlockStatement) Object {
	var result Object
	for _, statement := range block.Statements {
		result = e.Eval(statement)
		if result != nil {
			rt := result.Type()
			if rt == "RETURN_VALUE" || rt == "ERROR" {
				return result
			}
		}
	}
	return result
}

func (e *Evaluator) evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return e.evalBangOperatorExpression(right)
	case "-":
		return e.evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func (e *Evaluator) evalInfixExpression(operator string, left, right Object) Object {
	switch {
	case left.Type() == "INTEGER" && right.Type() == "INTEGER":
		return e.evalIntegerInfixExpression(operator, left, right)
	case left.Type() == "FLOAT" && right.Type() == "FLOAT":
		return e.evalFloatInfixExpression(operator, left, right)
	case left.Type() == "STRING" && right.Type() == "STRING":
		return e.evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return &Boolean{Value: left == right}
	case operator == "!=":
		return &Boolean{Value: left != right}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func (e *Evaluator) evalIntegerInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Integer).Value
	rightVal := right.(*Integer).Value
	switch operator {
	case "+":
		return &Integer{Value: leftVal + rightVal}
	case "-":
		return &Integer{Value: leftVal - rightVal}
	case "*":
		return &Integer{Value: leftVal * rightVal}
	case "/":
		return &Integer{Value: leftVal / rightVal}
	case "<":
		return &Boolean{Value: leftVal < rightVal}
	case ">":
		return &Boolean{Value: leftVal > rightVal}
	case "==":
		return &Boolean{Value: leftVal == rightVal}
	case "!=":
		return &Boolean{Value: leftVal != rightVal}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func (e *Evaluator) evalFloatInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Float).Value
	rightVal := right.(*Float).Value
	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "/":
		return &Float{Value: leftVal / rightVal}
	case "<":
		return &Boolean{Value: leftVal < rightVal}
	case ">":
		return &Boolean{Value: leftVal > rightVal}
	case "==":
		return &Boolean{Value: leftVal == rightVal}
	case "!=":
		return &Boolean{Value: leftVal != rightVal}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func (e *Evaluator) evalStringInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*String).Value
	rightVal := right.(*String).Value
	switch operator {
	case "+":
		return &String{Value: leftVal + rightVal}
	case "==":
		return &Boolean{Value: leftVal == rightVal}
	case "!=":
		return &Boolean{Value: leftVal != rightVal}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func (e *Evaluator) evalBangOperatorExpression(right Object) Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	default:
		return FALSE
	}
}

func (e *Evaluator) evalMinusPrefixOperatorExpression(right Object) Object {
	switch right.Type() {
	case "INTEGER":
		value := right.(*Integer).Value
		return &Integer{Value: -value}
	case "FLOAT":
		value := right.(*Float).Value
		return &Float{Value: -value}
	default:
		return newError("unknown operator: -%s", right.Type())
	}
}

func (e *Evaluator) evalIfExpression(ie *parser.IfExpression) Object {
	condition := e.Eval(ie.Condition)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return e.Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return e.Eval(ie.Alternative)
	} else {
		return NULL
	}
}

func (e *Evaluator) evalIdentifier(node *parser.Identifier) Object {
	if val, ok := e.env.Get(node.Value); ok {
		return val
	}
	return newError("identifier not found: " + node.Value)
}

func (e *Evaluator) evalExpressions(exps []parser.Expression) []Object {
	var result []Object
	for _, exp := range exps {
		evaluated := e.Eval(exp)
		if isError(evaluated) {
			return []Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func (e *Evaluator) applyFunction(fn Object, args []Object) Object {
	function, ok := fn.(*Function)
	if !ok {
		return newError("not a function: %s", fn.Type())
	}
	extendedEnv := e.extendFunctionEnv(function, args)
	evaluator := &Evaluator{env: extendedEnv}
	result := evaluator.Eval(function.Body)
	return unwrapReturnValue(result)
}

func (e *Evaluator) extendFunctionEnv(fn *Function, args []Object) *Environment {
	env := NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}
	return env
}

func (e *Evaluator) evalClassStatement(node *parser.ClassStatement) Object {
	class := &Class{
		Name:       node.Name.Value,
		Properties: make(map[string]Object),
		Methods:    make(map[string]*Function),
	}
	for _, prop := range node.Properties {
		val := getDefaultValue(prop.Type.Value)
		class.Properties[prop.Name.Value] = val
	}
	for _, method := range node.Methods {
		fn := &Function{
			Parameters: method.Parameters,
			Body:       method.Body,
			Env:        e.env,
		}
		class.Methods[method.Name.Value] = fn
	}
	return class
}

func (e *Evaluator) evalInstanceLiteral(node *parser.InstanceLiteral) Object {
	class := e.Eval(node.Class)
	if isError(class) {
		return class
	}
	classObj, ok := class.(*Class)
	if !ok {
		return newError("not a class: %s", class.Type())
	}
	instance := &Class{
		Name:       classObj.Name,
		Properties: make(map[string]Object),
		Methods:    make(map[string]*Function),
	}
	for name, prop := range classObj.Properties {
		instance.Properties[name] = prop
	}
	for name, method := range classObj.Methods {
		instance.Methods[name] = method
	}
	return instance
}

func (e *Evaluator) evalPropertyAccess(node *parser.PropertyAccess) Object {
	obj := e.Eval(node.Object)
	if isError(obj) {
		return obj
	}
	instance, ok := obj.(*Class)
	if !ok {
		return newError("not an instance: %s", obj.Type())
	}
	if val, ok := instance.Properties[node.Property.Value]; ok {
		return val
	}
	if method, ok := instance.Methods[node.Property.Value]; ok {
		return method
	}
	return newError("property not found: %s", node.Property.Value)
}

func isTruthy(obj Object) bool {
	switch obj {
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

func unwrapReturnValue(obj Object) Object {
	if returnValue, ok := obj.(*ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == "ERROR"
	}
	return false
}

func getDefaultValue(typeName string) Object {
	switch strings.ToLower(typeName) {
	case "int":
		return &Integer{Value: 0}
	case "float":
		return &Float{Value: 0.0}
	case "string":
		return &String{Value: ""}
	case "bool":
		return &Boolean{Value: false}
	default:
		return &Null{}
	}
}

var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)
