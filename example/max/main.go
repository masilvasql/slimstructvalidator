package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type User struct {
		Name  string `validate:"required,max=4" label:"Nome"`
		Email string `validate:"required,email" label:"E-mail"`
	}
	sv := slimstructvalidator.New()
	errs := sv.Validate(User{
		Name:  "Lucas",
		Email: "teste@teste.com",
	})

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Validação bem-sucedida!")
	}
}
