package gte

import (
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
	"strconv"
)

func init() {
	core.RegisterRuleWithParam("gte", Gte)
}

func Gte(value interface{}, param string, label string, kind reflect.Kind) *core.FieldError {
	gteVal, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return &core.FieldError{
			Field:   label,
			Tag:     "gte",
			Value:   value,
			Message: i18n.TError("invalidValue", "gte", param),
		}
	}

	v := reflect.ValueOf(value)
	field := label

	if field == "" {
		field = "Este campo"
	}

	var valid bool
	switch kind {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			switch item.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valid = float64(item.Int()) >= gteVal
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				valid = float64(item.Uint()) >= gteVal
			case reflect.Float32, reflect.Float64:
				valid = item.Float() >= gteVal
			default:
				valid = true
			}
			if !valid {
				break
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valid = float64(v.Int()) >= gteVal
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		valid = float64(v.Uint()) >= gteVal
	case reflect.Float32, reflect.Float64:
		valid = v.Float() >= gteVal
	default:
		valid = true
	}

	if !valid {
		return &core.FieldError{
			Field:   field,
			Tag:     "gte",
			Value:   value,
			Message: i18n.T("gte", label, param),
		}
	}

	return nil
}
