package main

import (
	"fmt"
	"io/ioutil"
	"jotlang/internal/eval"
	"jotlang/internal/lexer"
	"jotlang/internal/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: jotlang <arquivo.jt>")
		os.Exit(1)
	}

	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo %s: %s\n", filename, err)
		os.Exit(1)
	}

	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
		os.Exit(1)
	}

	env := eval.NewEnvironment()
	result := eval.Eval(program, env)

	if result != nil {
		fmt.Println(result.Inspect())
	}
}

func printParserErrors(errors []string) {
	fmt.Println("Erros de parser:")
	for _, msg := range errors {
		fmt.Printf("\t%s\n", msg)
	}
}
