package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
)

func init() {
	i18n.SetLanguage(i18n.PT_BR)
}

func main() {

	type User struct {
		Name   string `validate:"required,alpha" label:"Name"`
		Admin  bool   `validate:"eq=false" label:"Admin"`
		Gender string `validate:"oneOf=Male Female" label:"Gender"`
		Age    int    `validate:"numeric,min=18,max=65" label:"Age"`
		Email  string `validate:"email" label:"Email"`
	}

	sv := slimstructvalidator.New()

	user := User{
		Name:   "John123",
		Admin:  true,
		Gender: "Maleee",
		Age:    17,
		Email:  "invalid-email",
	}

	errs := sv.Validate(user)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Validation passed!")
	}
}
