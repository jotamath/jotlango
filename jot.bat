@echo off
if "%1"=="run" (
    go run main.go run %2
) else (
    echo Comando desconhecido: %1
    echo Uso: jot run <arquivo.jt>
) 