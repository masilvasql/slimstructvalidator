package i18n

import "fmt"

var currentLang = "pt"

var translations = map[string]string{
	"pt|required": "%s é obrigatório",
	"pt|min":      "%s deve ter pelo menos %s caracteres",
	"pt|email":    "%s deve ser um endereço de e-mail válido",
	"pt|max":      "%s deve ter no máximo %s caracteres",

	"en|required": "%s is required",
	"en|min":      "%s must have at least %s characters",
	"en|email":    "%s must be a valid email address",
	"en|max":      "%s must have at most %s characters",
}

func SetLanguage(lang string) {
	currentLang = lang
}

func T(key string, args ...any) string {
	langKey := fmt.Sprintf("%s|%s", currentLang, key)
	if tmpl, ok := translations[langKey]; ok {
		return fmt.Sprintf(tmpl, args...)
	}
	return key
}
