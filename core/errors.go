package core

type FieldError struct {
	Field   string
	Tag     string
	Value   interface{}
	Message string
}
