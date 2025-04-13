package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
)

func main() {
	type User struct {
		Name    string
		WebSite string `validate:"url" label:"WebSite"`
	}

	sv := slimstructvalidator.New()
	user := User{
		Name:    "John",
		WebSite: "invalid-url",
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
