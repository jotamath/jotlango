package core

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Estruturas para armazenar o estado do interpretador
type Variable struct {
	Type  string
	Value interface{}
}

type Function struct {
	Parameters []string
	Body       string
	ReturnType string
}

type Class struct {
	Properties map[string]Variable
	Methods    map[string]Function
}

type Interpreter struct {
	Variables map[string]Variable
	Functions map[string]Function
	Classes   map[string]Class
	Stack     []map[string]Variable
	Instances map[string]map[string]Variable
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		Variables: make(map[string]Variable),
		Functions: make(map[string]Function),
		Classes:   make(map[string]Class),
		Stack:     make([]map[string]Variable, 0),
		Instances: make(map[string]map[string]Variable),
	}
}

func Run(code string) (string, error) {
	interpreter := NewInterpreter()
	lines := strings.Split(code, "\n")
	var result strings.Builder

	// Primeira passagem: definir classes e funções
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		// Processar declaração de classe
		if strings.HasPrefix(line, "class ") {
			className := strings.TrimSpace(strings.TrimPrefix(line, "class "))
			interpreter.Classes[className] = Class{
				Properties: make(map[string]Variable),
				Methods:    make(map[string]Function),
			}
			continue
		}

		// Processar propriedades da classe
		if strings.HasPrefix(line, "prop ") {
			propDef := strings.TrimSpace(strings.TrimPrefix(line, "prop "))
			parts := strings.Split(propDef, " ")
			if len(parts) >= 2 {
				propType := parts[0]
				propName := parts[1]
				interpreter.Variables[propName] = Variable{
					Type:  propType,
					Value: getDefaultValue(propType),
				}
			}
			continue
		}

		// Processar declaração de função
		if strings.HasPrefix(line, "fn ") {
			funcDef := strings.TrimSpace(strings.TrimPrefix(line, "fn "))
			parts := strings.Split(funcDef, "(")
			funcName := strings.TrimSpace(parts[0])
			returnType := "void"
			if strings.Contains(funcDef, ":") {
				returnType = strings.TrimSpace(strings.Split(funcDef, ":")[1])
			}
			params := extractParameters(line)
			body := extractFunctionBody(lines[i:])
			interpreter.Functions[funcName] = Function{
				Parameters: params,
				Body:       body,
				ReturnType: returnType,
			}
			i += strings.Count(body, "\n")
			continue
		}
	}

	// Segunda passagem: executar o código
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		// Processar declaração de variável
		if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			varName := strings.TrimSpace(parts[0])
			varValue := strings.TrimSpace(parts[1])
			interpreter.Variables[varName] = interpreter.evaluateExpression(varValue)
			continue
		}

		// Processar print
		if strings.HasPrefix(line, "print(") {
			content := line[strings.Index(line, "(")+1 : strings.LastIndex(line, ")")]
			evaluated := interpreter.evaluateExpression(content)
			output := fmt.Sprintf("%v", evaluated.Value)
			result.WriteString(output + "\n")
			continue
		}

		// Processar chamada de método
		if strings.Contains(line, ".") && !strings.Contains(line, "=") {
			parts := strings.Split(line, ".")
			instanceName := strings.TrimSpace(parts[0])
			methodName := strings.TrimSpace(parts[1])
			if _, ok := interpreter.Instances[instanceName]; ok {
				if method, ok := interpreter.Functions[methodName]; ok {
					// Executar método
					executeFunction(method, interpreter, &result)
				}
			}
			continue
		}

		// Executar método Main se encontrado
		if strings.Contains(line, "fn Main()") {
			if method, ok := interpreter.Functions["Main"]; ok {
				executeFunction(method, interpreter, &result)
			}
		}
	}

	return result.String(), nil
}

func executeFunction(fn Function, interpreter *Interpreter, result *strings.Builder) {
	// Criar novo escopo para a função
	interpreter.Stack = append(interpreter.Stack, make(map[string]Variable))

	// Processar corpo da função
	lines := strings.Split(fn.Body, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Processar print dentro da função
		if strings.HasPrefix(line, "print(") {
			content := line[strings.Index(line, "(")+1 : strings.LastIndex(line, ")")]
			evaluated := interpreter.evaluateExpression(content)
			output := fmt.Sprintf("%v", evaluated.Value)
			result.WriteString(output + "\n")
			continue
		}

		// Processar declaração de variável
		if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			varName := strings.TrimSpace(parts[0])
			varValue := strings.TrimSpace(parts[1])
			interpreter.Stack[len(interpreter.Stack)-1][varName] = interpreter.evaluateExpression(varValue)
			continue
		}
	}

	// Remover escopo da função
	interpreter.Stack = interpreter.Stack[:len(interpreter.Stack)-1]
}

func getDefaultValue(typeName string) interface{} {
	switch typeName {
	case "string":
		return ""
	case "int":
		return 0
	case "float":
		return 0.0
	case "bool":
		return false
	default:
		return nil
	}
}

func extractParameters(line string) []string {
	re := regexp.MustCompile(`\((.*?)\)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) < 2 {
		return []string{}
	}
	params := strings.Split(matches[1], ",")
	for i := range params {
		params[i] = strings.TrimSpace(params[i])
	}
	return params
}

func extractFunctionBody(lines []string) string {
	var body strings.Builder
	depth := 0
	for _, line := range lines {
		if strings.Contains(line, "{") {
			depth++
		}
		if strings.Contains(line, "}") {
			depth--
		}
		if depth > 0 {
			body.WriteString(line + "\n")
		}
		if depth == 0 && strings.Contains(line, "}") {
			break
		}
	}
	return body.String()
}

func (i *Interpreter) evaluateExpression(expr string) Variable {
	// Remove espaços em branco
	expr = strings.TrimSpace(expr)

	// Verifica se é uma string com concatenação
	if strings.Contains(expr, "+") {
		parts := strings.Split(expr, "+")
		var result strings.Builder

		for _, part := range parts {
			part = strings.TrimSpace(part)

			// Se for uma string literal
			if strings.HasPrefix(part, "\"") && strings.HasSuffix(part, "\"") {
				result.WriteString(part[1 : len(part)-1])
			} else {
				// Se for uma variável
				if val, ok := i.Variables[part]; ok {
					result.WriteString(fmt.Sprintf("%v", val.Value))
				}
			}
		}

		return Variable{Type: "string", Value: result.String()}
	}

	// Verifica se é uma string simples
	if strings.HasPrefix(expr, "\"") && strings.HasSuffix(expr, "\"") {
		return Variable{Type: "string", Value: expr[1 : len(expr)-1]}
	}

	// Verifica se é um número
	if num, err := strconv.Atoi(expr); err == nil {
		return Variable{Type: "int", Value: num}
	}

	// Verifica se é um número decimal
	if num, err := strconv.ParseFloat(expr, 64); err == nil {
		return Variable{Type: "float", Value: num}
	}

	// Verifica se é um booleano
	if expr == "true" {
		return Variable{Type: "bool", Value: true}
	}
	if expr == "false" {
		return Variable{Type: "bool", Value: false}
	}

	// Verifica se é uma variável
	if val, ok := i.Variables[expr]; ok {
		return val
	}

	// Se não for nenhum dos casos acima, retorna uma string vazia
	return Variable{Type: "string", Value: ""}
}

func evaluateCondition(cond string, interpreter *Interpreter) bool {
	// Implementar avaliação de condições
	return true
}
