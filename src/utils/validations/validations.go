package validations

import (
	"context"
	"fmt"
	genericModels "playlist-app/src/models"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var bffValidator *validator.Validate

var customErrMapping = map[string]string{
	"required": "this field is required",
}

func GetBFFValidator(ctx context.Context) *validator.Validate {
	if bffValidator == nil {
		GetNewBFFValidator()
	}
	return bffValidator
}

func GetNewBFFValidator() {
	v := validator.New()
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		if tag := field.Tag.Get("json"); tag != "" && tag != "-" {
			return tag
		}
		return field.Name
	})
	bffValidator = v
}

func FormatValidationErrors(validationErrors validator.ValidationErrors) ([]genericModels.ErrorMessage, string) {
	var (
		formatted          []genericModels.ErrorMessage
		joinedMessageParts []string
	)

	for _, e := range validationErrors {
		msg := customErrMapping[e.Tag()]
		if msg == "" {
			switch e.Tag() {
			case " ":
				msg = fmt.Sprintf("%s field is required", e.Field())
			default:
				msg = e.Error()
			}
		}
		formatted = append(formatted, genericModels.ErrorMessage{
			Key:          e.Field(),
			ErrorMessage: msg,
		})
		joinedMessageParts = append(joinedMessageParts, e.Field()+","+msg)
	}
	return formatted, strings.Join(joinedMessageParts, ",")
}
