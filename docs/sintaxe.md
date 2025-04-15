# Tabela de Sintaxe da JotLang

## 1. Declarações Básicas

| Construção | Sintaxe | Exemplo |
|------------|---------|---------|
| Classe | `class Nome { ... }` | `class Pessoa { ... }` |
| Propriedade | `prop tipo nome` | `prop string Nome` |
| Função | `fn nome(param1 tipo, param2 tipo) : retorno { ... }` | `fn Soma(a int, b int) : int { ... }` |
| Construtor | `fn New() : void { ... }` | `fn New() : void { Nome = "João" }` |
| Variável | `var nome = valor` | `var idade = 25` |

## 2. Tipos de Dados

| Tipo | Descrição | Exemplo |
|------|-----------|---------|
| `int` | Número inteiro | `25` |
| `float` | Número decimal | `3.14` |
| `string` | Texto | `"Olá"` |
| `bool` | Booleano | `true`, `false` |
| `void` | Sem retorno | `fn Print() : void` |

## 3. Operadores

| Operador | Descrição | Exemplo |
|----------|-----------|---------|
| `=` | Atribuição | `var x = 10` |
| `+` | Adição/Concatenação | `a + b` |
| `-` | Subtração | `a - b` |
| `*` | Multiplicação | `a * b` |
| `/` | Divisão | `a / b` |
| `==` | Igualdade | `a == b` |
| `!=` | Diferença | `a != b` |
| `>` | Maior que | `a > b` |
| `<` | Menor que | `a < b` |
| `>=` | Maior ou igual | `a >= b` |
| `<=` | Menor ou igual | `a <= b` |

## 4. Estruturas de Controle

| Construção | Sintaxe | Exemplo |
|------------|---------|---------|
| If | `if condição { ... }` | `if idade > 18 { ... }` |
| If-Else | `if condição { ... } else { ... }` | `if idade > 18 { ... } else { ... }` |
| For | `for condição { ... }` | `for i < 10 { ... }` |
| While | `while condição { ... }` | `while i < 10 { ... }` |
| Return | `return valor` | `return a + b` |

## 5. Funções Especiais

| Função | Descrição | Exemplo |
|--------|-----------|---------|
| `print()` | Exibe texto | `print("Olá")` |
| `len()` | Tamanho de string/lista | `len(nome)` |
| `type()` | Tipo do valor | `type(idade)` |

## 6. Comentários

| Tipo | Sintaxe | Exemplo |
|------|---------|---------|
| Linha | `// comentário` | `// Esta é uma linha` |
| Bloco | `/* comentário */` | `/* Este é um bloco */` |

## 7. Instanciação de Classes

| Construção | Sintaxe | Exemplo |
|------------|---------|---------|
| Criar objeto | `var obj = new Classe()` | `var p = new Pessoa()` |
| Acessar propriedade | `obj.propriedade` | `p.Nome` |
| Chamar método | `obj.metodo()` | `p.Apresentar()` |

## 8. Arrays e Listas

| Construção | Sintaxe | Exemplo |
|------------|---------|---------|
| Declaração | `var lista = [tipo]` | `var numeros = [int]` |
| Inicialização | `var lista = [valor1, valor2]` | `var numeros = [1, 2, 3]` |
| Acesso | `lista[indice]` | `numeros[0]` |

## 9. Dicionários

| Construção | Sintaxe | Exemplo |
|------------|---------|---------|
| Declaração | `var dict = {tipo:tipo}` | `var pessoas = {string:Pessoa}` |
| Inicialização | `var dict = {chave:valor}` | `var idades = {"João":25}` |
| Acesso | `dict[chave]` | `idades["João"]` |

## 10. Tratamento de Erros

| Construção | Sintaxe | Exemplo |
|------------|---------|---------|
| Try-Catch | `try { ... } catch (erro) { ... }` | `try { ... } catch (e) { print(e) }` |
| Throw | `throw mensagem` | `throw "Erro ocorreu"` |

# Sintaxe da JotLang

## 📝 Estrutura Básica

### Classes

```jt
class NomeDaClasse {
    // Propriedades e métodos
}
```

### Propriedades

```jt
prop tipo Nome
```

Tipos disponíveis:
- `string` - Texto
- `int` - Números inteiros
- `bool` - Verdadeiro/Falso
- `float` - Números decimais
- `list<tipo>` - Listas
- `map<chave,valor>` - Mapas

### Funções

```jt
fn Nome(param1 tipo, param2 tipo) : tipoRetorno {
    // Corpo da função
}
```

### Instâncias

```jt
new NomeDaClasse {
    Propriedade1 = valor1,
    Propriedade2 = valor2
}
```

## 🔄 Controle de Fluxo

### Condicionais

```jt
if condicao then
    // código
else
    // código
```

### Loops

```jt
for item in lista {
    // código
}
```

## 📊 Operações Matemáticas

### Básicas

```jt
soma = 10 + 20
subtracao = 30 - 10
multiplicacao = 5 * 4
divisao = 20 / 4
```

### Funções Matemáticas

```jt
raiz = sqrt(16)
potencia = pow(2, 3)
```

## 📦 Estruturas de Dados

### Listas

```jt
lista = [1, 2, 3, 4, 5]
lista.add(6)
lista.remove(3)
```

### Mapas

```jt
mapa = {
    "chave1" = "valor1",
    "chave2" = "valor2"
}
```

## 🖨️ Saída

```jt
print("Texto {variavel}")
```

## 🔄 Concorrência

```jt
async fn Processar() : void {
    // Código assíncrono
}

call Processar
```

## 📡 APIs

```jt
@route("/api/endpoint")
class ApiController {
    fn Get() : void {
        // Código do endpoint
    }
}
```

## 🗃️ Banco de Dados

```jt
@entity Usuario {
    prop string Nome
    prop string Email
}

@repository
class UsuarioRepository {
    fn BuscarPorEmail(email string) : Usuario {
        // Query
    }
}
```

## 🔒 Autenticação

```jt
@auth
class AuthService {
    fn Login(email string, senha string) : bool {
        // Lógica de autenticação
    }
}
```

## 📝 Anotações

```jt
@route("/path")
@auth
@cache
fn Metodo() : void {
    // Código
}
```

## 🔄 Módulos

```jt
import "modulo"
import "modulo/submodulo"
```

## Tipos Básicos

```jt
// Tipos primitivos
string nome = "JotLang"
int idade = 25
float altura = 1.75
bool ativo = true

// Tipos compostos
list<string> nomes = ["João", "Maria", "Pedro"]
map<string, int> idades = {
    "João": 25,
    "Maria": 30,
    "Pedro": 28
}

// Conversão de tipos
string numero = "42"
int valor = Types.ToInt(numero)
```

## Classes e Objetos

```jt
class Pessoa {
    prop string Nome
    prop int Idade
    prop list<string> Hobbies

    fn Cumprimentar() : string {
        return "Olá, meu nome é {Nome}!"
    }
}

// Instanciação
pessoa = Pessoa.New()
pessoa.Nome = "João"
pessoa.Idade = 25
pessoa.Hobbies = ["Programar", "Ler", "Correr"]

// Uso
mensagem = pessoa.Cumprimentar()
```

## I/O Básico

```jt
// Saída
IO.Print("Olá, mundo!")
IO.Println("Linha com quebra")
IO.Printf("Nome: %s, Idade: %d", nome, idade)

// Entrada
nome = IO.ReadLine()
idade = IO.ReadInt()
altura = IO.ReadFloat()

// Arquivos
conteudo = IO.ReadFile("arquivo.txt")
IO.WriteFile("novo.txt", "Conteúdo")
IO.AppendFile("log.txt", "Nova entrada")
```

## HTTP

```jt
// Servidor
server = Server.New()
server.Port = 3000

// Rotas
server.Get("/", fn(req Request) : Response {
    return Response.Html("<h1>Olá, mundo!</h1>")
})

server.Post("/api/users", fn(req Request) : Response {
    return Response.Json({
        "message": "Usuário criado",
        "data": req.Body
    })
})

// Middleware
server.Use(fn(req Request, next fn() : Response) : Response {
    print("Request recebida: {req.Method} {req.Path}")
    return next()
})

// Iniciar servidor
server.Listen(3000)
```

## WebSocket

```jt
// Servidor
ws = WebSocket.Server.New()
ws.Port = 8080
ws.Path = "/chat"

// Eventos
ws.OnConnect = fn(conn Connection) : void {
    print("Cliente conectado: {conn.Id}")
    ws.Broadcast("Novo usuário conectado")
}

ws.OnMessage = fn(conn Connection, message string) : void {
    print("Mensagem recebida: {message}")
    ws.Broadcast("Usuário {conn.Id}: {message}")
}

ws.OnDisconnect = fn(conn Connection) : void {
    print("Cliente desconectado: {conn.Id}")
    ws.Broadcast("Usuário {conn.Id} saiu")
}

// Iniciar servidor
ws.Listen()
```

## Banco de Dados

```jt
// Conexão
db = Database.New()
db.ConnectionString = "postgres://user:pass@localhost:5432/db"
db.Provider = "postgres"
db.Connect()

// Query
resultados = db.Query("SELECT * FROM usuarios WHERE idade > ?", 18)

// Transação
tx = db.BeginTransaction()
try {
    db.Execute("INSERT INTO usuarios (nome, idade) VALUES (?, ?)", "João", 25)
    tx.Commit()
} catch {
    tx.Rollback()
}

// Entidade
class Usuario : Entity {
    prop string Nome
    prop int Idade
}

// Repository
repo = Repository<Usuario>.New(db, "usuarios")
usuarios = repo.FindAll()
joao = repo.FindById("1")
```

## Operações Matemáticas

```jt
// Básicas
soma = 10 + 20
subtracao = 30 - 15
multiplicacao = 6 * 8
divisao = 100 / 5

// Funções
raiz = Math.Sqrt(16)
potencia = Math.Pow(2, 8)
seno = Math.Sin(Math.PI / 4)
cosseno = Math.Cos(Math.PI / 4)
```

## Controle de Fluxo

```jt
// Condicionais
if idade >= 18 then
    print("Maior de idade")
else
    print("Menor de idade")

// Loops
for i in 1..10 {
    print(i)
}

for nome in nomes {
    print(nome)
}

while condicao {
    // código
}
``` 