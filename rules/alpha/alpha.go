package alpha

import (
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
	"unicode"
)

func init() {
	core.RegisterRule("alpha", Alpha)
}

func Alpha(value interface{}, label string, _ reflect.Kind) *core.FieldError {
	v := reflect.ValueOf(value)
	field := label

	if field == "" {
		field = "Este campo"
	}

	if v.Kind() == reflect.String {
		str := v.String()
		for _, r := range str {
			if !unicode.IsLetter(r) {
				return &core.FieldError{
					Field:   field,
					Tag:     "alpha",
					Value:   value,
					Message: i18n.T("alpha", label),
				}
			}
		}
	}

	return nil
}
