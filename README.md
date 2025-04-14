Slim Struct Validator

[![Go Reference](https://pkg.go.dev/badge/github.com/masilvasql/slimstructvalidator.svg)](https://pkg.go.dev/github.com/masilvasql/slimstructvalidator)

Uma biblioteca Go leve e simples para validação de campos de structs usando tags.

## Funcionalidades

*   **Validação Baseada em Tags:** Defina regras de validação diretamente nas tags (`validate:"..."`) das suas structs.
*   **Regras Comuns:** Inclui validadores para 
    * `alpha` : Apenas letras
    * `alphaNum` : Letras e números
    * `email`: Valida e-mails
    * `eq` : Valores devem ser iguais ao valor passados na tag
    * `max=X`: Tamanho ou valor máximo
    * `min=X`: Tamanho ou valor mínimo
    * `ne` : Valores não devem ser iguais ao valor passados na tag
    * `numeric` : Apenas números
    * `oneOf` : Um dos valores permitidos
    * `required`: Campo obrigatório
    * `url` : Permite URLs válidas
    * `gte` : Maior ou igual ao valor passado na tag

Não precisa da TAG, mas também efetua validação de `nestedstructure` quando uma struct possui outra struct como campo.


*   **Mensagens de Erro Amigáveis:** Use a tag `label:"..."` para nomes de campo mais claros nas mensagens de erro.
*   **Internacionalização (i18n):** Suporte integrado para tradução de mensagens de erro.
*   **Simples de Usar**

: ** API direta e fácil de integrar.

## Instalação

```bash
go get github.com/masilvasql/slimstructvalidator
```

#### possui exemplos em: example/main.go