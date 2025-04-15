package main

import (
	"fmt"
	"jotlango/internal/eval"
	"jotlango/internal/lexer"
	"jotlango/internal/parser"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: jot run <arquivo>")
		os.Exit(1)
	}

	command := os.Args[1]
	file := os.Args[2]

	if command != "run" {
		fmt.Println("Comando desconhecido:", command)
		os.Exit(1)
	}

	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		os.Exit(1)
	}

	l := lexer.NewLexer(string(content))
	p := parser.NewParser(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Println("Erros de parsing:")
		for _, err := range p.Errors() {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	evaluator := eval.NewEvaluator()
	result := evaluator.Eval(program)

	if result != nil {
		fmt.Println(result.Inspect())
	}
}
