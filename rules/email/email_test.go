package email_test

import (
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldPassWhenEmailIsValid_PT_BR(t *testing.T) {

	i18n.SetLanguage(i18n.PT_BR)

	type test struct {
		Email string `validate:"email" label:"Email"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Email: "john123@gmail.com",
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))

}

func TestShouldFailWhenEmailIsInvalid_PT_BR(t *testing.T) {
	i18n.SetLanguage(i18n.PT_BR)
	type test struct {
		Email string `validate:"email" label:"Email"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Email: "john123@gmail",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Email", errs[0].Field)
	assert.Equal(t, "email", errs[0].Tag)
	assert.Equal(t, "john123@gmail", errs[0].Value)
	assert.Equal(t, "O campo Email deve ser um endereço de e-mail válido", errs[0].Message)
}

func TestShouldPassWhenEmailIsValid_EN_US(t *testing.T) {

	i18n.SetLanguage(i18n.EN_US)

	type test struct {
		Email string `validate:"email" label:"Email"`
	}
	sv := slimstructvalidator.New()
	tst := test{
		Email: "john123@gmail.com",
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
}

func TestShouldFailWhenEmailIsInvalid_EN_US(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)
	type test struct {
		Email string `validate:"email" label:"Email"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Email: "john123@gmail",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Email", errs[0].Field)
	assert.Equal(t, "email", errs[0].Tag)
	assert.Equal(t, "john123@gmail", errs[0].Value)
	assert.Equal(t, "The field Email must be a valid email address", errs[0].Message)
}

func TestShouldFailWhenEmailIsEmpty(t *testing.T) {
	i18n.SetLanguage(i18n.EN_US)
	type test struct {
		Email string `validate:"email" label:"Email"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Email: "",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Email", errs[0].Field)
	assert.Equal(t, "email", errs[0].Tag)
	assert.Equal(t, "", errs[0].Value)
	assert.Equal(t, "The field Email must be a valid email address", errs[0].Message)
}

func TestShouldPassWhenEmailIsValid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)
	type test struct {
		Email string `validate:"email" label:"Email"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Email: "john123@gmail.com",
	}
	errs := sv.Validate(tst)
	assert.Empty(t, errs)
	assert.Len(t, errs, 0)
	assert.Equal(t, 0, len(errs))

}

func TestShouldFailWhenEmailIsInvalid_ES_ES(t *testing.T) {
	i18n.SetLanguage(i18n.ES_ES)
	type test struct {
		Email string `validate:"email" label:"Email"`
	}

	sv := slimstructvalidator.New()
	tst := test{
		Email: "john123@gmail",
	}
	errs := sv.Validate(tst)
	assert.Len(t, errs, 1)
	assert.Equal(t, "Email", errs[0].Field)
	assert.Equal(t, "email", errs[0].Tag)
	assert.Equal(t, "john123@gmail", errs[0].Value)
	assert.Equal(t, "El campo Email debe ser una dirección de correo electrónico válida", errs[0].Message)
}
