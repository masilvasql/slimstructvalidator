package alphanum_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlphaNumShouldPassWhenValid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Name string `validate:"alphaNum" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "John123",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestAlphaNumShouldFailWhenInvalid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Name string `validate:"alphaNum" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "John@123",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nome", errs[0].Field)
	assert.Equal(t, "alphaNum", errs[0].Tag)
	assert.Equal(t, "John@123", errs[0].Value)
	assert.Equal(t, "O campo Nome deve conter apenas letras e números", errs[0].Message)
}

func TestAlphaNumShouldPassWhenValid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Name string `validate:"alphaNum" label:"Name"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "John123",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestAlphaNumShouldFailWhenInvalid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Name string `validate:"alphaNum" label:"Name"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "John@123",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Name", errs[0].Field)
	assert.Equal(t, "alphaNum", errs[0].Tag)
	assert.Equal(t, "John@123", errs[0].Value)
	assert.Equal(t, "The field Name must contain only letters and numbers", errs[0].Message)
}

func TestAlphaNumShouldPassWhenValid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Name string `validate:"alphaNum" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "John123",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestAlphaNumShouldFailWhenInvalid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Name string `validate:"alphaNum" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "John@123",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nombre", errs[0].Field)
	assert.Equal(t, "alphaNum", errs[0].Tag)
	assert.Equal(t, "John@123", errs[0].Value)
	assert.Equal(t, "El campo Nombre debe contener solo letras y números", errs[0].Message)
}
