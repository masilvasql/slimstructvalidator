package required

import (
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
)

func init() {
	core.RegisterRule("required", Required)
}

func Required(value interface{}, label string, kind reflect.Kind) *core.FieldError {
	v := reflect.ValueOf(value)
	zero := reflect.Zero(v.Type()).Interface()

	if reflect.DeepEqual(value, zero) {
		field := label
		if field == "" {
			field = "Este campo"
		}

		return &core.FieldError{
			Field:   field,
			Tag:     "required",
			Value:   value,
			Message: i18n.T("required", label),
		}
	}
	return nil
}
