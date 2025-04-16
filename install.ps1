# Script de instalação do JotLang
Write-Host "Instalando JotLang..."

# Criar diretórios necessários
New-Item -ItemType Directory -Force -Path "stdlib"
New-Item -ItemType Directory -Force -Path "stdlib/types"
New-Item -ItemType Directory -Force -Path "stdlib/io"
New-Item -ItemType Directory -Force -Path "stdlib/http"
New-Item -ItemType Directory -Force -Path "stdlib/websocket"
New-Item -ItemType Directory -Force -Path "stdlib/database"

# Criar arquivos da biblioteca padrão
# Types
@"
class Types {
    prop string String
    prop int Int
    prop float Float
    prop bool Bool

    fn ToString(value: any): string {
        return value.ToString()
    }

    fn ToInt(value: string): int {
        return int.Parse(value)
    }

    fn ToFloat(value: string): float {
        return float.Parse(value)
    }

    fn ToBool(value: string): bool {
        return bool.Parse(value)
    }
}
"@ | Out-File -FilePath "stdlib/types/types.jt" -Encoding utf8

# IO
@"
class IO {
    prop string FilePath

    fn ReadFile(path: string): string {
        return File.ReadAllText(path)
    }

    fn WriteFile(path: string, content: string): void {
        File.WriteAllText(path, content)
    }

    fn AppendFile(path: string, content: string): void {
        File.AppendAllText(path, content)
    }

    fn FileExists(path: string): bool {
        return File.Exists(path)
    }

    fn DeleteFile(path: string): void {
        File.Delete(path)
    }
}
"@ | Out-File -FilePath "stdlib/io/io.jt" -Encoding utf8

# HTTP
@"
class Request {
    prop string Method
    prop string Path
    prop map[string]string Headers
    prop string Body
}

class Response {
    prop int StatusCode
    prop map[string]string Headers
    prop string Body

    fn Send(status: int, body: string): void {
        this.StatusCode = status
        this.Body = body
    }
}

class Server {
    prop string Host
    prop int Port
    prop map[string]Function Routes

    fn Start(): void {
        print("Servidor iniciado em " + this.Host + ":" + this.Port)
    }

    fn Stop(): void {
        print("Servidor parado")
    }

    fn AddRoute(path: string, handler: Function): void {
        this.Routes[path] = handler
    }
}
"@ | Out-File -FilePath "stdlib/http/http.jt" -Encoding utf8

# WebSocket
@"
class Connection {
    prop string Id
    prop bool IsOpen
    prop Function OnMessage

    fn Send(message: string): void {
        if this.IsOpen {
            print("Enviando mensagem: " + message)
        }
    }

    fn Close(): void {
        this.IsOpen = false
    }
}

class Server {
    prop string Host
    prop int Port
    prop map[string]Connection Connections

    fn Start(): void {
        print("Servidor WebSocket iniciado em " + this.Host + ":" + this.Port)
    }

    fn Stop(): void {
        print("Servidor WebSocket parado")
    }
}
"@ | Out-File -FilePath "stdlib/websocket/websocket.jt" -Encoding utf8

# Database
@"
class Database {
    prop string ConnectionString
    prop bool IsConnected

    fn Connect(): void {
        this.IsConnected = true
        print("Conectado ao banco de dados")
    }

    fn Disconnect(): void {
        this.IsConnected = false
        print("Desconectado do banco de dados")
    }
}

class Transaction {
    prop bool IsActive

    fn Begin(): void {
        this.IsActive = true
    }

    fn Commit(): void {
        this.IsActive = false
    }

    fn Rollback(): void {
        this.IsActive = false
    }
}

class Entity {
    prop int Id
    prop string CreatedAt
    prop string UpdatedAt
}

class Repository<T> {
    prop Database Db

    fn Create(entity: T): void {
        print("Criando entidade")
    }

    fn Read(id: int): T {
        print("Lendo entidade")
        return null
    }

    fn Update(entity: T): void {
        print("Atualizando entidade")
    }

    fn Delete(id: int): void {
        print("Deletando entidade")
    }
}
"@ | Out-File -FilePath "stdlib/database/database.jt" -Encoding utf8

# Criar arquivo go.mod se não existir
if (-not (Test-Path "go.mod")) {
    @"
module jotlango

go 1.21

require (
    github.com/gorilla/websocket v1.5.1
    github.com/joho/godotenv v1.5.1
)
"@ | Out-File -FilePath "go.mod" -Encoding utf8
}

# Instalar dependências
go mod tidy

# Compilar o programa
go build -o jot.exe main.go

# Criar diretório para o JotLang no PATH do usuário
$jotPath = "$env:USERPROFILE\JotLang"
if (-not (Test-Path $jotPath)) {
    New-Item -ItemType Directory -Path $jotPath
}

# Mover o executável para o diretório
Copy-Item -Force jot.exe $jotPath

# Criar arquivo batch para executar o JotLang
@"
@echo off
if "%1"=="run" (
    %~dp0jot.exe run %2
) else (
    echo Comando desconhecido: %1
    echo Uso: jot run <arquivo.jt>
)
"@ | Out-File -FilePath "$jotPath\jot.bat" -Encoding utf8

# Adicionar ao PATH se ainda não estiver
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (-not $userPath.Contains($jotPath)) {
    [Environment]::SetEnvironmentVariable("Path", "$userPath;$jotPath", "User")
}

Write-Host "JotLang instalado com sucesso!"
Write-Host "Agora você pode usar 'jot run arquivo.jt' em qualquer lugar."
Write-Host "Reinicie o PowerShell para que as alterações no PATH tenham efeito." 