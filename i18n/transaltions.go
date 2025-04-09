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

	"en-US|required":     "%s is required",
	"en-US|email":        "%s must be a valid email address",
	"en-US|min":          "%s must have a value or size greater than or equal to %s",
	"en-US|max":          "%s must have a value or size of at most %s",
	"en-US|invalidValue": "value of tag %s is invalid: %s",

	"es-ES|required":     "%s es obligatorio",
	"es-ES|email":        "%s debe ser una dirección de correo electrónico válida",
	"es-ES|min":          "%s debe tener un valor o tamaño mayor o igual a %s",
	"es-ES|max":          "%s debe tener un valor o tamaño máximo de %s",
	"es-ES|invalidValue": "valor de la etiqueta %s inválido: %s",
}
