package main

import (
	"fmt"
	"github.com/masilvasql/slimstructvalidator"
	"github.com/masilvasql/slimstructvalidator/i18n"
)

func main() {
	basicExample()
	fmt.Println("--------------------------------")
	mailValidate()
	fmt.Println("--------------------------------")
	minValidate()
	fmt.Println("--------------------------------")
	maxValidate()
}

func mailValidate() {
	fmt.Print("Mail Validate Start\n")

	type User struct {
		Name  string `validate:"required" label:"Nome"`
		Email string `validate:"required,email" label:"E-mail"`
	}

	sv := slimstructvalidator.New()
	errs := sv.Validate(User{
		Name:  "Lucas",
		Email: "teste.teste.com",
	})

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Validação bem-sucedida!")
	}

	fmt.Println("Mail Validate End\n")
}

func basicExample() {
	i18n.SetLanguage("pt") // ou "pt"
	fmt.Print("Basic Example Start\n")
	type User struct {
		Name  string `validate:"required" label:"Nome"`
		Email string `validate:"required,email" label:"E-mail"`
	}

	sv := slimstructvalidator.New()
	errs := sv.Validate(User{})

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		println("Validação bem-sucedida!")
	}

	fmt.Println("Basic Example End\n")

}

func minValidate() {
	fmt.Print("Min Validate Start\n")

	type User struct {
		Name  string `validate:"required,min=3" label:"Nome"`
		Email string `validate:"required,email" label:"E-mail"`
	}

	sv := slimstructvalidator.New()
	errs := sv.Validate(User{
		Name:  "Lu",
		Email: "teste@teste.com",
	})

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%s: %s\n", err.Field, err.Message)
		}
	} else {
		println("Validação bem-sucedida!")
	}

	fmt.Println("Min Validate End\n")
}

func maxValidate() {
	fmt.Print("Max Validate Start\n")
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
	fmt.Println("Max Validate End\n")
}
