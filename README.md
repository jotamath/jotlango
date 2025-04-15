# JotLang 🚀

Uma linguagem de programação moderna e intuitiva, inspirada no melhor do Go, TypeScript e Python.

## ✨ Características

- **Sintaxe limpa e moderna**
- **Tipagem forte com inferência**
- **Orientação a objetos com classes**
- **Funcionalidades de primeira classe**
- **Concorrência nativa**
- **Biblioteca padrão robusta**
- **Interoperabilidade com Go**

## 🚀 Primeiros Passos

### Instalação

```bash
# Clone o repositório
git clone https://github.com/jotamath/jotlango.git
cd jotlango

# Instale as dependências
go mod download

# Instale a CLI
go install ./cmd/jot
```

### Hello World

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

## 📚 Documentação

- [Sintaxe](docs/sintaxe.md) - Guia completo da sintaxe
- [Tutorial](docs/tutorial.md) - Tutorial passo a passo
- [API](docs/api.md) - Documentação da API
- [Exemplos](examples/) - Exemplos práticos

## 🏗️ Exemplos

### API REST

```jt
class Api {
    fn Main() : void {
        server = Server.New()
        server.Port = 3000

        server.Get("/", fn(req Request) : Response {
            return Response.Json({
                "message": "Olá, Mundo!"
            })
        })

        server.Listen(3000)
    }
}
```

### WebSocket

```jt
class Chat {
    fn Main() : void {
        ws = WebSocket.Server.New()
        ws.Port = 8080

        ws.OnMessage = fn(conn Connection, msg string) : void {
            ws.Broadcast("Usuário {conn.Id}: {msg}")
        }

        ws.Listen()
    }
}
```

### Banco de Dados

```jt
class App {
    fn Main() : void {
        db = Database.New()
        db.ConnectionString = "postgres://user:pass@localhost:5432/db"
        db.Connect()

        usuarios = db.Query("SELECT * FROM usuarios")
        for usuario in usuarios {
            print("Nome: {usuario.nome}")
        }
    }
}
```

## 🛠️ Ferramentas

- **CLI** - Compilador e executor
- **LSP** - Suporte a IDEs
- **Debugger** - Depurador integrado
- **Formatter** - Formatador de código
- **Linter** - Analisador estático

## 🤝 Contribuição

1. Faça um fork do projeto
2. Crie uma branch (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova funcionalidade'`)
4. Push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👥 Comunidade

- [Discord](https://discord.gg/jotlang)
- [Twitter](https://twitter.com/jotlang)
- [GitHub Issues](https://github.com/jotamath/jotlango/issues)

---

Feito com ❤️ pela comunidade JotLang
