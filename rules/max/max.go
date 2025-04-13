package max

import (
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
	"strconv"
)

func init() {
	core.RegisterRuleWithParam("max", Max)
}

func Max(value interface{}, param string, label string, kind reflect.Kind) *core.FieldError {
	maxVal, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return &core.FieldError{
			Field:   label,
			Tag:     "max",
			Value:   value,
			Message: i18n.TError("invalidValue", "max", param),
		}
	}

	v := reflect.ValueOf(value)
	field := label

	if field == "" {
		field = "Este campo"
	}

	var valid bool
	switch kind {
	case reflect.String, reflect.Slice, reflect.Array:
		valid = float64(v.Len()) <= maxVal
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valid = float64(v.Int()) <= maxVal
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		valid = float64(v.Uint()) <= maxVal
	case reflect.Float32, reflect.Float64:
		valid = v.Float() <= maxVal
	default:
		valid = true
	}

	strMin := strconv.FormatFloat(maxVal, 'f', -1, 64)

	if !valid {
		return &core.FieldError{
			Field:   field,
			Tag:     "max",
			Value:   value,
			Message: i18n.T("max", label, strMin),
		}
	}

	return nil
}
