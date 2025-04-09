package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
)

func init() {
	i18n.SetLanguage(i18n.EN_US)
}

func main() {
	type User struct {
		Name string `validate:"required" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	errs := sv.Validate(User{})

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		println("Validação bem-sucedida!")
	}
}
