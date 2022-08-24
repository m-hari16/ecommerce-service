package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func PanicIfError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func TranslateError(err error) (errs []string) {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		errs = append(errs, ToSnake(e.Field())+" is "+e.Tag())
	}
	return errs
}
