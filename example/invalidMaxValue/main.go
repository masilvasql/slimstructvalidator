package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type Product struct {
		Name string `validate:"required,max=zzz" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	errs := sv.Validate(Product{
		Name: "Notebook",
	})

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Validação bem-sucedida!")
	}
}
