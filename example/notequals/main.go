package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type User struct {
		Name  string
		Admin bool `validate:"ne=false" label:"Admin"`
	}

	sv := slimstructvalidator.New()
	user := User{
		Name:  "John",
		Admin: false,
	}

	errs := sv.Validate(user)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err.Message)
		}
	} else {
		fmt.Println("Validation passed!")
	}
}
