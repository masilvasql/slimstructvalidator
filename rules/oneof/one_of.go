package oneof

import (
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
	"strings"
)

func init() {
	core.RegisterRuleWithParam("oneOf", OneOf)
}

func OneOf(value interface{}, param string, label string, kind reflect.Kind) *core.FieldError {
	v := reflect.ValueOf(value)
	field := label

	if field == "" {
		field = "Este campo"
	}

	if v.Kind() == reflect.String || v.Kind() == reflect.Int || v.Kind() == reflect.Float64 {
		valid := false
		for _, p := range strings.Split(param, " ") {
			if value == p {
				valid = true
				break
			}
		}

		param = strings.Join(strings.Split(param, " "), ", ")

		if !valid {
			return &core.FieldError{
				Field:   field,
				Tag:     "oneOf",
				Value:   value,
				Message: i18n.T("oneOf", label, param),
			}
		}
	}

	return nil
}
