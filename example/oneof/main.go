package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type User struct {
		Name   string
		Gender string `validate:"oneOf=Male Female" label:"Gender"`
	}

	sv := slimstructvalidator.New()
	user := User{
		Name:   "John",
		Gender: "test",
	}

	errs := sv.Validate(user)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err.Field, err.Message)
		}
	} else {
		fmt.Println("Validation passed!")
	}
}
