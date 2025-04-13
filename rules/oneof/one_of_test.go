package oneof_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPassWhenOneOfIsValid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type test struct {
		Name string `validate:"oneOf=Lucas John" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "Lucas",
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenOneOfIsInvalid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type test struct {
		Name string `validate:"oneOf=Lucas John" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "João",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nome", errs[0].Field)
	assert.Equal(t, "oneOf", errs[0].Tag)
	assert.Equal(t, "João", errs[0].Value)
	assert.Equal(t, "O campo Nome deve ser um dos valores: Lucas, John", errs[0].Message)
}

func TestShouldPassWhenOneOfIsValid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type test struct {
		Name string `validate:"oneOf=Lucas John" label:"Name"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "Lucas",
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenOneOfIsInvalid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type test struct {
		Name string `validate:"oneOf=Lucas John" label:"Name"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "João",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Name", errs[0].Field)
	assert.Equal(t, "oneOf", errs[0].Tag)
	assert.Equal(t, "João", errs[0].Value)
	assert.Equal(t, "The field Name must be one of the values: Lucas, John", errs[0].Message)
}

func TestShouldPassWhenOneOfIsValid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type test struct {
		Name string `validate:"oneOf=Lucas John" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "Lucas",
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenOneOfIsInvalid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type test struct {
		Name string `validate:"oneOf=Lucas John" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "João",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nombre", errs[0].Field)
	assert.Equal(t, "oneOf", errs[0].Tag)
	assert.Equal(t, "João", errs[0].Value)
	assert.Equal(t, "El campo Nombre debe ser uno de los valores: Lucas, John", errs[0].Message)
}
