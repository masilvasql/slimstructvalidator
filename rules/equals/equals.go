package equals

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
)

func init() {
	core.RegisterRuleWithParam("eq", Equals)
}
func Equals(value interface{}, param string, label string, kind reflect.Kind) *core.FieldError {
	v := reflect.ValueOf(value)
	field := label

	if field == "" {
		field = "Este campo"
	}

	var valueString string
	switch kind {
	case reflect.String:
		valueString = v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valueString = fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		valueString = fmt.Sprintf("%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		valueString = fmt.Sprintf("%f", v.Float())
	case reflect.Bool:
		valueString = fmt.Sprintf("%t", v.Bool())
	default:
		valueString = fmt.Sprintf("%v", value)
	}

	if valueString != param {
		return &core.FieldError{
			Field:   field,
			Tag:     "eq",
			Value:   value,
			Message: i18n.T("eq", label, param), // Ex: "%s deve ser igual a %s"
		}
	}

	return nil
}
