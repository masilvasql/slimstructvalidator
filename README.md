Slim Struct Validator

[![Go Reference](https://pkg.go.dev/badge/github.com/masilvasql/slimstructvalidator.svg)](https://pkg.go.dev/github.com/masilvasql/slimstructvalidator)

Uma biblioteca Go leve e simples para validação de campos de structs usando tags.

## Funcionalidades

*   **Validação Baseada em Tags:** Defina regras de validação diretamente nas tags (`validate:"..."`) das suas structs.
*   **Regras Comuns:** Inclui validadores para `required`, `email`, `min=X`, `max=X`.
*   **Mensagens de Erro Amigáveis:** Use a tag `label:"..."` para nomes de campo mais claros nas mensagens de erro.
*   **Internacionalização (i18n):** Suporte integrado para tradução de mensagens de erro.
*   **Simples de Usar**

: ** API direta e fácil de integrar.

## Instalação

```bash
go get github.com/masilvasql/slimstructvalidator
```

#### possui exemplos em: example/main.go