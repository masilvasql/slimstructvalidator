package email

import (
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func init() {
	core.RegisterRule("email", Email)
}

func Email(value interface{}, label string, kind reflect.Kind) *core.FieldError {
	str, ok := value.(string)
	if !ok {
		return nil
	}

	if !emailRegex.MatchString(str) || len(str) == 0 {
		field := label
		if field == "" {
			field = "Este campo"
		}

		return &core.FieldError{
			Field:   field,
			Tag:     "email",
			Value:   value,
			Message: i18n.T("email", field),
		}
	}
	return nil
}
