package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type Product struct {
		Name    string `validate:"required" label:"Nome"`
		BarCode string `validate:"required,numeric" label:"Código de Barras"`
	}

	sv := slimstructvalidator.New()
	product := Product{
		Name:    "Produto1",
		BarCode: "-1234567@abc",
	}

	errs := sv.Validate(product)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Validation passed!")
	}
}
