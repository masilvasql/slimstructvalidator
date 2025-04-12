Slim Struct Validator

[![Go Reference](https://pkg.go.dev/badge/github.com/masilvasql/slimstructvalidator.svg)](https://pkg.go.dev/github.com/masilvasql/slimstructvalidator)

Uma biblioteca Go leve e simples para validação de campos de structs usando tags.

## Funcionalidades

*   **Validação Baseada em Tags:** Defina regras de validação diretamente nas tags (`validate:"..."`) das suas structs.
*   **Regras Comuns:** Inclui validadores para 
    * `required`: Campo obrigatório
    * `email`: Valida e-mails
    * `min=X`: Tamanho ou valor mínimo
    * `max=X`: Tamanho ou valor máximo
    * `alpha` : Apenas letras
    * `alphaNum` : Letras e números
    * `numeric` : Apenas números
    * `equals` : Igualdade entre campos
    * `oneOf` : Um dos valores permitidos
    * `nestedstructure` : Validações em structs aninhadas


*   **Mensagens de Erro Amigáveis:** Use a tag `label:"..."` para nomes de campo mais claros nas mensagens de erro.
*   **Internacionalização (i18n):** Suporte integrado para tradução de mensagens de erro.
*   **Simples de Usar**

: ** API direta e fácil de integrar.

## Instalação

```bash
go get github.com/masilvasql/slimstructvalidator
```

#### possui exemplos em: example/main.go