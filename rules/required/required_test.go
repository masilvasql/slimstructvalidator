package required_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPassWhenRequired_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Name string `validate:"required" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "Lucas",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenRequired_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Name string `validate:"required" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nome", errs[0].Field)
	assert.Equal(t, "required", errs[0].Tag)
	assert.Equal(t, "", errs[0].Value)
	assert.Equal(t, "O campo Nome é obrigatório", errs[0].Message)
}

func TestShouldPassWhenRequired_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Name string `validate:"required" label:"Name"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "Lucas",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenRequired_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Name string `validate:"required" label:"Name"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Name", errs[0].Field)
	assert.Equal(t, "required", errs[0].Tag)
	assert.Equal(t, "", errs[0].Value)
	assert.Equal(t, "The field Name is required", errs[0].Message)
}

func TestShouldPassWhenRequired_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Name string `validate:"required" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "Lucas",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenRequired_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Name string `validate:"required" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nombre", errs[0].Field)
	assert.Equal(t, "required", errs[0].Tag)
	assert.Equal(t, "", errs[0].Value)
	assert.Equal(t, "El campo Nombre es obligatorio", errs[0].Message)
}
