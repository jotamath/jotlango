# Tutorial de Primeiros Passos

## 🚀 Instalação

1. Instale a JotLang:
```bash
go install github.com/jotamath/jotlango/cmd/jot
```

2. Verifique a instalação:
```bash
jot --version
```

## 📝 Primeiro Programa

1. Crie um novo projeto:
```bash
jot new meu-projeto
cd meu-projeto
```

2. Edite o arquivo `main.jt`:
```jt
class App {
    fn Main() : void {
        print("Olá, JotLang!")
    }
}

call Main
```

3. Execute o programa:
```bash
jot run main.jt
```

## 📚 Conceitos Básicos

### Variáveis e Tipos

```jt
// Números
idade = 25
preco = 19.99

// Texto
nome = "João"

// Lista
frutas = ["maçã", "banana", "laranja"]

// Mapa
pessoa = {
    "nome" = "Maria",
    "idade" = 30
}
```

### Funções

```jt
fn Soma(a int, b int) : int {
    return a + b
}

fn Saudacao(nome string) : void {
    print("Olá, {nome}!")
}
```

### Classes

```jt
class Pessoa {
    prop string Nome
    prop int Idade

    fn Apresentar() : void {
        print("Meu nome é {Nome} e tenho {Idade} anos.")
    }
}

// Criar instância
p = new Pessoa {
    Nome = "João",
    Idade = 25
}

// Chamar método
call p.Apresentar()
```

## 🎯 Projeto Prático: Todo List

Vamos criar uma lista de tarefas simples:

```jt
class Tarefa {
    prop string Titulo
    prop bool Concluida

    fn MarcarComoConcluida() : void {
        Concluida = true
        print("Tarefa '{Titulo}' concluída!")
    }
}

class TodoList {
    prop list<Tarefa> Tarefas

    fn AdicionarTarefa(titulo string) : void {
        tarefa = new Tarefa {
            Titulo = titulo,
            Concluida = false
        }
        Tarefas.add(tarefa)
        print("Tarefa '{titulo}' adicionada!")
    }

    fn ListarTarefas() : void {
        print("=== Lista de Tarefas ===")
        for tarefa in Tarefas {
            status = if tarefa.Concluida then "✓" else " "
            print("[{status}] {tarefa.Titulo}")
        }
    }
}

// Criar lista
lista = new TodoList { Tarefas = [] }

// Adicionar tarefas
call lista.AdicionarTarefa("Aprender JotLang")
call lista.AdicionarTarefa("Criar um projeto")

// Listar tarefas
call lista.ListarTarefas()

// Marcar como concluída
tarefa = lista.Tarefas[0]
call tarefa.MarcarComoConcluida()

// Listar novamente
call lista.ListarTarefas()
```

## 📡 API Simples

Vamos criar uma API básica:

```jt
@route("/api/tarefas")
class TarefaApi {
    prop list<Tarefa> Tarefas

    fn Listar() : list<Tarefa> {
        return Tarefas
    }

    fn Adicionar(titulo string) : Tarefa {
        tarefa = new Tarefa {
            Titulo = titulo,
            Concluida = false
        }
        Tarefas.add(tarefa)
        return tarefa
    }
}

// Criar API
api = new TarefaApi { Tarefas = [] }

// Adicionar tarefa
call api.Adicionar("Tarefa via API")

// Listar tarefas
tarefas = call api.Listar()
for tarefa in tarefas {
    print("Tarefa: {tarefa.Titulo}")
}
```

## 🚀 Próximos Passos

1. Explore os exemplos em `examples/`
2. Leia a documentação completa em `docs/`
3. Crie seu próprio projeto
4. Contribua com a comunidade 