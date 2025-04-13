package i18n

var currentLang = "pt-BR"

type Language string

const (
	PT_BR Language = "pt-BR"
	EN_US Language = "en-US"
	ES_ES Language = "es-ES"
)

var translations = map[string]string{
	"pt-BR|required":     "%s é obrigatório",
	"pt-BR|email":        "%s deve ser um endereço de e-mail válido",
	"pt-BR|min":          "%s deve ter valor ou tamanho maior ou igual a %s",
	"pt-BR|max":          "%s deve ter o valor ou tamanho máximo de %s",
	"pt-BR|invalidValue": "valor da tag %s inválido: %s",
	"pt-BR|alpha":        "%s deve conter apenas letras",
	"pt-BR|alphaNum":     "%s deve conter apenas letras e números",
	"pt-BR|numeric":      "%s deve conter apenas números",
	"pt-BR|oneOf":        "%s deve ser um dos valores: %s",
	"pt-BR|eq":           "%s deve ser igual a %s",
	"pt-BR|ne":           "%s não deve ter valor igual a %s",
	"pt-BR|url":          "%s deve ser uma URL válida",

	"en-US|required":     "%s is required",
	"en-US|email":        "%s must be a valid email address",
	"en-US|min":          "%s must have a value or size greater than or equal to %s",
	"en-US|max":          "%s must have a value or size of at most %s",
	"en-US|invalidValue": "value of tag %s is invalid: %s",
	"en-US|alpha":        "%s must contain only letters",
	"en-US|alphaNum":     "%s must contain only letters and numbers",
	"en-US|numeric":      "%s must contain only numbers",
	"en-US|oneOf":        "%s must be one of the values: %s",
	"en-US|eq":           "%s must be equal to %s",
	"en-US|ne":           "%s must not be equal to %s",
	"en-US|url":          "%s must be a valid URL",

	"es-ES|required":     "%s es obligatorio",
	"es-ES|email":        "%s debe ser una dirección de correo electrónico válida",
	"es-ES|min":          "%s debe tener un valor o tamaño mayor o igual a %s",
	"es-ES|max":          "%s debe tener un valor o tamaño máximo de %s",
	"es-ES|invalidValue": "valor de la etiqueta %s inválido: %s",
	"es-ES|alpha":        "%s debe contener solo letras",
	"es-ES|alphaNum":     "%s debe contener solo letras y números",
	"es-ES|numeric":      "%s debe contener solo números",
	"es-ES|oneOf":        "%s debe ser uno de los valores: %s",
	"es-ES|eq":           "%s debe ser igual a %s",
	"es-ES|ne":           "%s no debe ser igual a %s",
	"es-ES|url":          "%s debe ser una URL válida",
}
