package numeric_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPassWhenValid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		Age int `validate:"numeric" label:"Idade"`
	}

	sv := slimstructvalidator.New()
	u := user{
		Age: 25,
	}
	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))
}

func TestShouldFailWhenInvalid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)

	type user struct {
		BarCode string `validate:"numeric" label:"BarCode"`
	}

	sv := slimstructvalidator.New()
	u := user{
		BarCode: "1234@abc",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "BarCode", errs[0].Field)
	assert.Equal(t, "numeric", errs[0].Tag)
	assert.Equal(t, "1234@abc", errs[0].Value)
	assert.Equal(t, "O campo BarCode deve conter apenas números", errs[0].Message)
}

func TestShouldPassWhenValid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		BarCode string `validate:"numeric" label:"BarCode"`
	}

	sv := slimstructvalidator.New()
	u := user{
		BarCode: "1234567890",
	}

	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))

}

func TestShouldFailWhenInvalid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)

	type user struct {
		BarCode string `validate:"numeric" label:"BarCode"`
	}

	sv := slimstructvalidator.New()
	u := user{
		BarCode: "1234@abc",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "BarCode", errs[0].Field)
	assert.Equal(t, "numeric", errs[0].Tag)
	assert.Equal(t, "1234@abc", errs[0].Value)
	assert.Equal(t, "The field BarCode must contain only numbers", errs[0].Message)
}

func TestShouldPassWhenValid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		BarCode string `validate:"numeric" label:"BarCode"`
	}

	sv := slimstructvalidator.New()
	u := user{
		BarCode: "1234567890",
	}

	errs := sv.Validate(u)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))

}

func TestShouldFailWhenInvalid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)

	type user struct {
		BarCode string `validate:"numeric" label:"BarCode"`
	}

	sv := slimstructvalidator.New()
	u := user{
		BarCode: "1234@abc",
	}
	errs := sv.Validate(u)
	assert.Len(t, errs, 1)
	assert.Equal(t, "BarCode", errs[0].Field)
	assert.Equal(t, "numeric", errs[0].Tag)
	assert.Equal(t, "1234@abc", errs[0].Value)
	assert.Equal(t, "El campo BarCode debe contener solo números", errs[0].Message)
}
