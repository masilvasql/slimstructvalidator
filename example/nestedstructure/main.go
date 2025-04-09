package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type Country struct {
		Name string `validate:"required" label:"Nome País"`
	}
	type Address struct {
		Street  string  `validate:"required" label:"Rua"`
		Country Country `validate:"required" label:"País"`
	}

	type User struct {
		Name    string  `validate:"required,min=3" label:"Nome"`
		Email   string  `validate:"required,email" label:"E-mail"`
		Address Address `validate:"required" label:"Endereço"`
	}

	sv := slimstructvalidator.New()
	errs := sv.Validate(User{
		Name:    "Lucas",
		Email:   "teste@teste.com",
		Address: Address{},
	})
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Validação bem-sucedida!")
	}
}
