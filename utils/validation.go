package utils

import (
	"strings"

	"github.com/Scholak/go-json-validation/lang/en"
	"github.com/Scholak/go-json-validation/lang/tr"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Set default language
var lang string = "tr"

func ValidateStruct(data interface{}) map[string]string {
	errors := map[string]string{}
	var errorMessage string

	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			// Defining messages according to their error tag
			switch err.Tag() {

			case "required":
				if lang == "en" {
					errorMessage = en.ValidataionErrors["required"]
				} else if lang == "tr" {
					errorMessage = tr.ValidataionErrors["required"]
				}

			case "min":
				if lang == "en" {
					errorMessage = en.ValidataionErrors["min"]
				} else if lang == "tr" {
					errorMessage = tr.ValidataionErrors["min"]
				}
				errorMessage = strings.Replace(errorMessage, ":value", err.Param(), -1)

			case "max":
				if lang == "en" {
					errorMessage = en.ValidataionErrors["max"]
				} else if lang == "tr" {
					errorMessage = tr.ValidataionErrors["max"]
				}
				errorMessage = strings.Replace(errorMessage, ":value", err.Param(), -1)
			}
			errors[err.Field()] = strings.Replace(errorMessage, ":attribute", err.Field(), -1)
		}
	}

	// Checks the errors map is empty or not. If it is empty, it returnss nil.
	if len(errors) == 0 {
		return nil
	} else {
		return errors
	}
}
