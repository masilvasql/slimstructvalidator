package min_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPassWhenMinValueIsValid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type test struct {
		Age int `validate:"min=18" label:"Idade"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Age: 18,
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenMinValueIsInvalid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type test struct {
		Age int `validate:"min=18" label:"Idade"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Age: 17,
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Idade", errs[0].Field)
	assert.Equal(t, "min", errs[0].Tag)
	assert.Equal(t, 17, errs[0].Value)
	assert.Equal(t, "O campo Idade deve ter valor ou tamanho maior ou igual a 18", errs[0].Message)
}

func TestShouldPassWhenMinValueIsValid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type test struct {
		Age int `validate:"min=18" label:"Age"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Age: 18,
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenMinValueIsInvalid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type test struct {
		Age int `validate:"min=18" label:"Age"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Age: 17,
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Age", errs[0].Field)
	assert.Equal(t, "min", errs[0].Tag)
	assert.Equal(t, 17, errs[0].Value)
	assert.Equal(t, "The field Age must have a value or size greater than or equal to 18", errs[0].Message)
}

func TestShouldPassWhenMinValueIsValid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type test struct {
		Age int `validate:"min=18" label:"Edad"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Age: 18,
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenMinValueIsInvalid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type test struct {
		Age int `validate:"min=18" label:"Edad"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Age: 17,
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Edad", errs[0].Field)
	assert.Equal(t, "min", errs[0].Tag)
	assert.Equal(t, 17, errs[0].Value)
	assert.Equal(t, "El campo Edad debe tener un valor o tamaño mayor o igual a 18", errs[0].Message)
}
