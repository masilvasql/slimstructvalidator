package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type User struct {
		Name string `validate:"alphaNum" label:"Nome"`
	}

	sv := slimstructvalidator.New()
	user := User{
		Name: "John123@!",
	}

	errs := sv.Validate(user)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s\n", err.Message)
		}
	} else {
		fmt.Println("Validation passed!")
	}
}
