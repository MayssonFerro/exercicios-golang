# Exercícios em Go

Este repositório contém exercícios práticos de programação em Go.

## Tutorial - Como compilar e executar

### Opção 1: Compilar e executar diretamente
```bash
go run file.go
```

### Opção 2: Criar executável e depois executar
```bash
# Compilar e criar executável
go build file.go

# Executar o arquivo gerado
./file.exe
```

## Notas importantes

- Cada arquivo `.go` deve ser compilado individualmente para evitar conflitos de múltiplas funções `main()`
- Use `go run arquivo.go` para execução direta
- Use `go build arquivo.go` para criar executável ```bash