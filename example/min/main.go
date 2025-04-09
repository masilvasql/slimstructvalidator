package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type Product struct {
		Name  string  `validate:"required,min=3" label:"Nome"`
		Value float64 `validate:"required,min=1.00,max=1000" label:"Valor"`
	}

	sv := slimstructvalidator.New()
	errs := sv.Validate(Product{
		Name:  "PC",
		Value: 0.01,
	})

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Validação bem-sucedida!")
	}
}
