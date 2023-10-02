package helper

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	errorhelper "bnt/bnt-indent-service/internal/core/helper/error-helper"
)

func Validate(data interface{}) error {
	LogEvent("INFO", "Validating "+reflect.TypeOf(data).String()+" Data...")
	err := validator.New().Struct(data)
	if err != nil {
		var fieldErrors []validator.FieldError
		LogEvent("ERROR", "Error validating struct: "+err.Error())
		for _, errs := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, errs)
		}
		return errorhelper.ErrorArrayToError(fieldErrors)
	}
	LogEvent("INFO", reflect.TypeOf(data).String()+" Data Validated Successfully...")
	return nil
}
