package core

import (
	"reflect"
	"strings"
)

var registeredRules = map[string]RuleFunc{}
var registeredRulesWithParam = map[string]RuleWithParamFunc{}

type RuleFunc func(value interface{}, label string, kind reflect.Kind) *FieldError
type RuleWithParamFunc func(value interface{}, param string, label string, kind reflect.Kind) *FieldError

func RegisterRule(tag string, fn RuleFunc) {
	registeredRules[tag] = fn
}

func RegisterRuleWithParam(tag string, fn RuleWithParamFunc) {
	registeredRulesWithParam[tag] = fn
}

func ValidateStruct(s interface{}) []*FieldError {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	var errs []*FieldError

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("validate")
		label := fieldType.Tag.Get("label")

		if tag == "" {
			continue
		}

		tags := strings.Split(tag, ",")
		for _, t := range tags {
			ruleName := t
			param := ""

			// Captura regras no formato "min=3"
			if parts := strings.SplitN(t, "=", 2); len(parts) == 2 {
				ruleName = parts[0]
				param = parts[1]
			}

			// Regras com param
			if fn, ok := registeredRulesWithParam[ruleName]; ok {
				if err := fn(field.Interface(), param, label, field.Kind()); err != nil {
					errs = append(errs, err)
				}
				continue
			}

			// Regras simples
			if fn, ok := registeredRules[ruleName]; ok {
				if err := fn(field.Interface(), label, field.Kind()); err != nil {
					errs = append(errs, err)
				}
			}
		}
	}

	return errs
}
