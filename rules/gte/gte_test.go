package gte_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPassWhenIsValidValue_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)
	type user struct {
		Age int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: 100,
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Empty(t, err)

}

func TestShouldPassWhenIsValidValueWithSlice_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)
	type user struct {
		Age []int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: []int{100, 101},
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Empty(t, err)
}

func TestShouldFailWhenIsInvalidValue_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)
	type user struct {
		Age int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: 99,
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Len(t, err, 1)
	assert.Equal(t, "Idade", err[0].Field)
	assert.Equal(t, "gte", err[0].Tag)
	assert.Equal(t, 99, err[0].Value)
	assert.Equal(t, "O campo Idade deve ser maior ou igual a 100", err[0].Message)
}

func TestShouldFailWhenIsInvalidValueWithSlice_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)
	type user struct {
		Age []int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: []int{99, 101},
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Len(t, err, 1)
	assert.Equal(t, "Idade", err[0].Field)
	assert.Equal(t, "gte", err[0].Tag)
	assert.Equal(t, "O campo Idade deve ser maior ou igual a 100", err[0].Message)
}

func TestShouldPassWhenIsValidValue_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)
	type user struct {
		Age int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: 100,
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Empty(t, err)
}

func TestShouldPassWhenIsValidValueWithSlice_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)
	type user struct {
		Age []int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: []int{100, 101},
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Empty(t, err)
}

func TestShouldFailWhenIsInvalidValue_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)
	type user struct {
		Age int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: 99,
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Len(t, err, 1)
	assert.Equal(t, "Idade", err[0].Field)
	assert.Equal(t, "gte", err[0].Tag)
	assert.Equal(t, 99, err[0].Value)
	assert.Equal(t, "The field Idade must be greater than or equal to 100", err[0].Message)
}

func TestShouldFailWhenIsInvalidValueWithSlice_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)
	type user struct {
		Age []int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: []int{99, 101},
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Len(t, err, 1)
	assert.Equal(t, "Idade", err[0].Field)
	assert.Equal(t, "gte", err[0].Tag)
	assert.Equal(t, "The field Idade must be greater than or equal to 100", err[0].Message)
}

func TestShouldPassWhenIsValidValue_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)
	type user struct {
		Age int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: 100,
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Empty(t, err)
}

func TestShouldPassWhenIsValidValueWithSlice_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)
	type user struct {
		Age []int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: []int{100, 101},
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Empty(t, err)
}

func TestShouldFailWhenIsInvalidValue_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)
	type user struct {
		Age int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: 99,
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Len(t, err, 1)
	assert.Equal(t, "Idade", err[0].Field)
	assert.Equal(t, "gte", err[0].Tag)
	assert.Equal(t, 99, err[0].Value)
	assert.Equal(t, "El campo Idade debe ser mayor o igual a 100", err[0].Message)
}

func TestShouldFailWhenIsInvalidValueWithSlice_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)
	type user struct {
		Age []int `validate:"gte=100" label:"Idade"`
	}

	u := user{
		Age: []int{99, 101},
	}
	sv := slimstructvalidator.New()

	err := sv.Validate(u)
	assert.Len(t, err, 1)
	assert.Equal(t, "Idade", err[0].Field)
	assert.Equal(t, "gte", err[0].Tag)
	assert.Equal(t, "El campo Idade debe ser mayor o igual a 100", err[0].Message)
}
