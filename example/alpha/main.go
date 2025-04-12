package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type User struct {
		Name string `validate:"required,alpha" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	user := User{
		Name: "John123",
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
