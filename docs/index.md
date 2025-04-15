---
layout: default
title: JotLang - Documentação
nav_order: 1
---

# Bem-vindo à Documentação do JotLang 🌱

<p align="center">
  <img src="assets/jot.png" alt="Jot - O mascote da JotLang" width="200"/>
</p>

JotLang é uma linguagem de programação moderna e expressiva, projetada para ser a melhor amiga do desenvolvedor backend. Com foco em performance e simplicidade, a JotLang oferece uma experiência de desenvolvimento agradável e poderosa.

## 🚀 Começando

### Instalação Rápida

```bash
# Clone o repositório
git clone https://github.com/jotamath/jotlango.git
cd jotlango

# Instale as dependências
go mod download

# Instale a CLI
go install ./cmd/jot
```

### Primeiro Programa

```jt
// main.jt
class App {
    fn Main() : void {
        print("Olá, Mundo!")
    }
}
```

Execute com:
```bash
jot run main.jt
```

## 📚 Conteúdo da Documentação

### [Guia de Sintaxe](sintaxe.md)
- Tipos de dados
- Classes e objetos
- Funções e métodos
- Controle de fluxo
- Operações matemáticas

### [Tutorial](tutorial.md)
- Primeiros passos
- Projetos práticos
- Integração com APIs
- Desenvolvimento web

### [Framework Web](framework.md)
- Servidor HTTP
- Sistema de rotas
- Middleware
- Integração com React/Next.js

### [Exemplos](exemplos/)
- [Hello World](exemplos/HelloWorld.jt)
- [Todo List](exemplos/TodoApi.jt)
- [Autenticação](exemplos/AuthApi.jt)
- [Middleware](exemplos/middleware/)

## 🎯 Recursos Principais

### Sistema de Tipos Forte
```jt
class Usuario {
    prop string nome
    prop int idade
    prop list<string> interesses
}
```

### Classes e Objetos
```jt
class Calculadora {
    fn Soma(int a, int b) -> int {
        return a + b
    }
}

calc = Calculadora.New()
resultado = calc.Soma(5, 3)
```

### Operações Matemáticas
```jt
fn CalcularArea(float raio) -> float {
    return 3.14159 * pow(raio, 2)
}
```

### Controle de Fluxo
```jt
fn VerificarIdade(int idade) -> string {
    if idade >= 18 then
        return "Maior de idade"
    else
        return "Menor de idade"
}
```

## 🤝 Contribuindo

Contribuições são bem-vindas! Consulte nosso [guia de contribuição](contribuicao.md) para saber como ajudar no desenvolvimento da JotLang.

## 📄 Licença

Este projeto está licenciado sob a [MIT License](license.md). 