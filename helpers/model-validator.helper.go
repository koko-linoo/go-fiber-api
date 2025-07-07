package helpers

import (
	"fmt"

	"github.com/go-playground/validator"
)

var check = validator.New()

func Validate(u any) *[]string {

	messages := make([]string, 0)

	err := check.Struct(u)
	if err != nil {
		switch e := err.(type) {
		case validator.ValidationErrors:
			for _, fieldErr := range e {
				messages = append(messages, fmt.Sprintf("%s is %s", fieldErr.Field(), fieldErr.Tag()))
			}
		default:
			messages = append(messages, e.Error())
		}

		return &messages
	}
	return nil
}
