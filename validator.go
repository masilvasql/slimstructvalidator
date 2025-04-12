package slimstructvalidator

import (
	"github.com/masilvasql/slimstructvalidator/core"
	_ "github.com/masilvasql/slimstructvalidator/rules/alpha"
	_ "github.com/masilvasql/slimstructvalidator/rules/alphanum"
	_ "github.com/masilvasql/slimstructvalidator/rules/email"
	_ "github.com/masilvasql/slimstructvalidator/rules/equals"
	_ "github.com/masilvasql/slimstructvalidator/rules/max"
	_ "github.com/masilvasql/slimstructvalidator/rules/min"
	_ "github.com/masilvasql/slimstructvalidator/rules/numeric"
	_ "github.com/masilvasql/slimstructvalidator/rules/oneof"
	_ "github.com/masilvasql/slimstructvalidator/rules/required"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(input interface{}) []*core.FieldError {

	return core.ValidateStruct(input)

}
