package i18n

var currentLang = "pt-BR"

type Language string

const (
	PT_BR Language = "pt-BR"
	EN_US Language = "en-US"
	ES_ES Language = "es-ES"
)

var translations = map[string]string{
	"pt-BR|required":     "O campo %s é obrigatório",
	"pt-BR|email":        "O campo %s deve ser um endereço de e-mail válido",
	"pt-BR|min":          "O campo %s deve ter valor ou tamanho maior ou igual a %s",
	"pt-BR|max":          "O campo %s deve ter o valor ou tamanho máximo de %s",
	"pt-BR|invalidValue": "valor da tag %s inválido: %s",
	"pt-BR|alpha":        "O campo %s deve conter apenas letras",
	"pt-BR|alphaNum":     "O campo %s deve conter apenas letras e números",
	"pt-BR|numeric":      "O campo %s deve conter apenas números",
	"pt-BR|oneOf":        "O campo %s deve ser um dos valores: %s",
	"pt-BR|eq":           "O campo %s deve ser igual a %s",
	"pt-BR|ne":           "O campo %s não deve ter valor igual a %s",
	"pt-BR|url":          "O campo %s deve ser uma URL válida",

	"en-US|required":     "The field %s is required",
	"en-US|email":        "The field %s must be a valid email address",
	"en-US|min":          "The field %s must have a value or size greater than or equal to %s",
	"en-US|max":          "The field %s must have a value or size of at most %s",
	"en-US|invalidValue": "value of tag %s is invalid: %s",
	"en-US|alpha":        "The field %s must contain only letters",
	"en-US|alphaNum":     "The field %s must contain only letters and numbers",
	"en-US|numeric":      "The field %s must contain only numbers",
	"en-US|oneOf":        "The field %s must be one of the values: %s",
	"en-US|eq":           "The field %s must be equal to %s",
	"en-US|ne":           "The field %s must not be equal to %s",
	"en-US|url":          "The field %s must be a valid URL",

	"es-ES|required":     "El campo %s es obligatorio",
	"es-ES|email":        "El campo %s debe ser una dirección de correo electrónico válida",
	"es-ES|min":          "El campo %s debe tener un valor o tamaño mayor o igual a %s",
	"es-ES|max":          "El campo %s debe tener un valor o tamaño máximo de %s",
	"es-ES|invalidValue": "valor de la etiqueta %s inválido: %s",
	"es-ES|alpha":        "El campo %s debe contener solo letras",
	"es-ES|alphaNum":     "El campo %s debe contener solo letras y números",
	"es-ES|numeric":      "El campo %s debe contener solo números",
	"es-ES|oneOf":        "El campo %s debe ser uno de los valores: %s",
	"es-ES|eq":           "El campo %s debe ser igual a %s",
	"es-ES|ne":           "El campo %s no debe ser igual a %s",
	"es-ES|url":          "El campo %s debe ser una URL válida",
}
