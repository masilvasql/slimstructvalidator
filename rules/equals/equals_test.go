package equals_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualsShouldPassWhenEqual_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Name string `validate:"eq=John" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "John",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestEqualsShouldFailWhenNotEqual_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Name string `validate:"eq=John" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "Lucas",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nome", errs[0].Field)
	assert.Equal(t, "eq", errs[0].Tag)
	assert.Equal(t, "Lucas", errs[0].Value)
	assert.Equal(t, "O campo Nome deve ser igual a John", errs[0].Message)
}

func TestEqualsShouldPassWhenEqual_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Name string `validate:"eq=John" label:"Name"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "John",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestEqualsShouldFailWhenNotEqual_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Name string `validate:"eq=John" label:"Name"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "Lucas",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Name", errs[0].Field)
	assert.Equal(t, "eq", errs[0].Tag)
	assert.Equal(t, "Lucas", errs[0].Value)
	assert.Equal(t, "The field Name must be equal to John", errs[0].Message)
}

func TestEqualsShouldPassWhenEqual_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Name string `validate:"eq=John" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "John",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestEqualsShouldFailWhenNotEqual_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Name string `validate:"eq=John" label:"Nombre"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Name: "Lucas",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Nombre", errs[0].Field)
	assert.Equal(t, "eq", errs[0].Tag)
	assert.Equal(t, "Lucas", errs[0].Value)
	assert.Equal(t, "El campo Nombre debe ser igual a John", errs[0].Message)
}
