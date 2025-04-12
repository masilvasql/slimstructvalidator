package numeric

import (
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
	"unicode"
)

func init() {
	core.RegisterRule("numeric", Numeric)
}

func Numeric(value interface{}, label string, _ reflect.Kind) *core.FieldError {
	v := reflect.ValueOf(value)
	//permitir positivos e negativos
	if v.Kind() == reflect.String {
		str := v.String()
		for _, r := range str {
			if !unicode.IsNumber(r) && r != '-' && r != '.' {
				return &core.FieldError{
					Field:   label,
					Tag:     "numeric",
					Value:   value,
					Message: i18n.T("numeric", label),
				}
			}
		}
	}

	return nil
}
