package url_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPassWhenValid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Website string `validate:"url" label:"Website"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Website: "https://example.com",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenInvalid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Website string `validate:"url" label:"Website"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Website: "invalid-url",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Website", errs[0].Field)
	assert.Equal(t, "url", errs[0].Tag)
	assert.Equal(t, "invalid-url", errs[0].Value)
	assert.Equal(t, "O campo Website deve ser uma URL válida", errs[0].Message)
}

func TestShouldPassWhenValid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Website string `validate:"url" label:"Website"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Website: "https://example.com",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenInvalid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		Website string `validate:"url" label:"Website"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Website: "invalid-url",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Website", errs[0].Field)
	assert.Equal(t, "url", errs[0].Tag)
	assert.Equal(t, "invalid-url", errs[0].Value)
	assert.Equal(t, "The field Website must be a valid URL", errs[0].Message)
}

func TestShouldPassWhenValid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Website string `validate:"url" label:"Website"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Website: "https://example.com",
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenInvalid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		Website string `validate:"url" label:"Website"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Website: "invalid-url",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Website", errs[0].Field)
	assert.Equal(t, "url", errs[0].Tag)
	assert.Equal(t, "invalid-url", errs[0].Value)
	assert.Equal(t, "El campo Website debe ser una URL válida", errs[0].Message)
}
