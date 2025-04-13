package alphanum

import (
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
	"unicode"
)

func init() {
	core.RegisterRule("alphaNum", AlphaNum)
}

func AlphaNum(value interface{}, label string, kind reflect.Kind) *core.FieldError {
	v := reflect.ValueOf(value)
	field := label

	if field == "" {
		field = "Este Campo"
	}

	if v.Kind() == reflect.String {
		str := v.String()
		for _, r := range str {
			if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
				return &core.FieldError{
					Field:   field,
					Tag:     "alphaNum",
					Value:   value,
					Message: i18n.T("alphaNum", label),
				}
			}
		}
	}
	return nil
}
