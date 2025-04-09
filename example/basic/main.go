package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

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
