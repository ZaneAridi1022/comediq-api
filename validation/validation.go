package validation

import "gopkg.in/go-playground/validator.v8"

var Validate *validator.Validate

func Init() {
	config := &validator.Config{TagName: "validate"}
	Validate = validator.New(config)
}
