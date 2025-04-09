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

		if !field.CanInterface() {
			continue
		}

		// 🌀 Validação recursiva se for struct (e não tempo ou slice/map/etc)
		if field.Kind() == reflect.Struct && field.Type().PkgPath() != "time" {
			nestedErrs := ValidateStruct(field.Interface())
			errs = append(errs, nestedErrs...)
			continue
		}

		tag := fieldType.Tag.Get("validate")
		label := fieldType.Tag.Get("label")

		if tag == "" {
			continue
		}

		tags := strings.Split(tag, ",")
		for _, t := range tags {
			ruleName := t
			param := ""

			err := validateRulesWithParam(ruleName, param, t, field, label)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			err = validateSimpleRules(ruleName, field, label)
			if err != nil {
				errs = append(errs, err)
				continue
			}
		}
	}

	return errs
}

func validateRulesWithParam(ruleName string, param string, t string, field reflect.Value, label string) *FieldError {
	if parts := strings.SplitN(t, "=", 2); len(parts) == 2 {
		ruleName = parts[0]
		param = parts[1]
	}

	if fn, ok := registeredRulesWithParam[ruleName]; ok {
		if err := fn(field.Interface(), param, label, field.Kind()); err != nil {
			return err
		}
	}
	return nil
}

func validateSimpleRules(ruleName string, field reflect.Value, label string) *FieldError {
	if fn, ok := registeredRules[ruleName]; ok {
		if err := fn(field.Interface(), label, field.Kind()); err != nil {
			return err
		}
	}
	return nil
}
