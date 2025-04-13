package url

import (
	"github.com/masilvasql/slimstructvalidator/core"
	"github.com/masilvasql/slimstructvalidator/i18n"
	"reflect"
	"regexp"
)

var urlRegex = regexp.MustCompile(`^(https?:\/\/)?([a-zA-Z0-9\-]+\.)+[a-zA-Z]{2,}(:\d+)?(\/[^\s]*)?$`)

func init() {
	core.RegisterRule("url", Email)
}

func Email(value interface{}, label string, kind reflect.Kind) *core.FieldError {
	str, ok := value.(string)
	if !ok || str == "" {
		return nil
	}

	if !urlRegex.MatchString(str) {
		field := label
		if field == "" {
			field = "Este campo"
		}

		return &core.FieldError{
			Field:   field,
			Tag:     "url",
			Value:   value,
			Message: i18n.T("url", field),
		}
	}
	return nil
}
