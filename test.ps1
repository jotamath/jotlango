# Script de teste
Write-Host "Executando testes..." -ForegroundColor Green

# Verificar se o interpretador existe
if (-not (Test-Path "jot.exe")) {
    Write-Host "Interpretador não encontrado. Execute install.ps1 primeiro." -ForegroundColor Red
    exit 1
}

# Executar o teste
.\jot.exe run test.jt

Write-Host "Teste concluído!" -ForegroundColor Green 