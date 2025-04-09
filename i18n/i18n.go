package i18n

import "fmt"

func SetLanguage(language Language) {
	switch language {
	case PT_BR:
		currentLang = "pt-BR"
	case EN_US:
		currentLang = "en-US"
	case ES_ES:
		currentLang = "es-ES"
	default:
		currentLang = "pt-BR"
	}
}

func T(key string, args ...any) string {
	langKey := fmt.Sprintf("%s|%s", currentLang, key)
	if tmpl, ok := translations[langKey]; ok {
		return fmt.Sprintf(tmpl, args...)
	}
	return key
}

func TError(errorMessage string, tag string, args ...any) string {
	langKey := fmt.Sprintf("%s|%s", currentLang, errorMessage)
	if tmpl, ok := translations[langKey]; ok {
		allArgs := append([]any{tag}, args...)
		return fmt.Sprintf(tmpl, allArgs...)
	}
	return errorMessage
}
