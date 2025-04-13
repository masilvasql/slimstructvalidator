package neuals_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPassWhenNotEquals_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type test struct {
		Name string `validate:"ne=Lucas" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "John",
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenEquals_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type test struct {
		Name string `validate:"ne=Lucas" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "Lucas",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nome", errs[0].Field)
	assert.Equal(t, "ne", errs[0].Tag)
	assert.Equal(t, "Lucas", errs[0].Value)
	assert.Equal(t, "O campo Nome não deve ter valor igual a Lucas", errs[0].Message)
}

func TestShouldPassWhenNotEquals_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type test struct {
		Name string `validate:"ne=Lucas" label:"Name"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "John",
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenEquals_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type test struct {
		Name string `validate:"ne=Lucas" label:"Name"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "Lucas",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Name", errs[0].Field)
	assert.Equal(t, "ne", errs[0].Tag)
	assert.Equal(t, "Lucas", errs[0].Value)
	assert.Equal(t, "The field Name must not be equal to Lucas", errs[0].Message)
}

func TestShouldPassWhenNotEquals_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type test struct {
		Name string `validate:"ne=Lucas" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "John",
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenEquals_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type test struct {
		Name string `validate:"ne=Lucas" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Name: "Lucas",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nombre", errs[0].Field)
	assert.Equal(t, "ne", errs[0].Tag)
	assert.Equal(t, "Lucas", errs[0].Value)
	assert.Equal(t, "El campo Nombre no debe ser igual a Lucas", errs[0].Message)
}
