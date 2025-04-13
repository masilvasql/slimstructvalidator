package max_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPassWhenValidMaxValue_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Age int `validate:"max=100" label:"Idade"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Age: 99,
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenInvalidMaxValue_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Age int `validate:"max=100" label:"Idade"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Age: 101,
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Idade", errs[0].Field)
	assert.Equal(t, "max", errs[0].Tag)
	assert.Equal(t, 101, errs[0].Value)
	assert.Equal(t, "O campo Idade deve ter o valor ou tamanho máximo de 100", errs[0].Message)
}

func TestShouldPassWhenValidMaxValue_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Age int `validate:"max=100" label:"Age"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Age: 99,
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenInvalidMaxValue_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Age int `validate:"max=100" label:"Age"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Age: 101,
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Age", errs[0].Field)
	assert.Equal(t, "max", errs[0].Tag)
	assert.Equal(t, 101, errs[0].Value)
	assert.Equal(t, "The field Age must have a value or size of at most 100", errs[0].Message)
}

func TestShouldPassWhenValidMaxValue_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Age int `validate:"max=100" label:"Edad"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Age: 99,
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenInvalidMaxValue_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Age int `validate:"max=100" label:"Edad"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Age: 101,
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Edad", errs[0].Field)
	assert.Equal(t, "max", errs[0].Tag)
	assert.Equal(t, 101, errs[0].Value)
	assert.Equal(t, "El campo Edad debe tener un valor o tamaño máximo de 100", errs[0].Message)
}
