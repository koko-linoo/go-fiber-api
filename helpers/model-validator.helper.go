package helpers

import (
	"fmt"

	"github.com/go-playground/validator"
)

var check = validator.New()

func Validate(u any) *[]string {

	if err := check.Struct(u); err != nil {
		errs := err.(validator.ValidationErrors)
		messages := make([]string, 0)
		for _, fieldErr := range errs {
			messages = append(messages, fmt.Sprintf("%s is %s", fieldErr.Field(), fieldErr.Tag()))
		}
		return &messages
	}

	return nil
}
