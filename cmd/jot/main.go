package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jotamath/jotlango/core"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: jot run <arquivo.jt>")
		os.Exit(1)
	}

	if os.Args[1] != "run" {
		fmt.Println("Comando inválido. Use: jot run <arquivo.jt>")
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Println("Arquivo não especificado")
		os.Exit(1)
	}

	file := os.Args[2]
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("Arquivo não encontrado: %s\n", file)
		os.Exit(1)
	}

	// Verificar extensão do arquivo
	if filepath.Ext(file) != ".jt" {
		fmt.Println("Arquivo deve ter extensão .jt")
		os.Exit(1)
	}

	// Ler o arquivo
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		os.Exit(1)
	}

	// Executar o código
	result, err := core.Run(string(content))
	if err != nil {
		fmt.Printf("Erro ao executar código: %v\n", err)
		os.Exit(1)
	}

	// Mostrar resultado
	if result != "" {
		fmt.Print(result)
	}
}
