package entities

import (
	"fmt"

	"github.com/go-playground/validator"
)

type Student struct {
	ID   int64  `json:"id"`
	Name string `json:"name" validate:"required,max=50,min=3"`
	CPF  string `json:"cpf" validate:"required,max=15,min=11"`
	RG   string `json:"rg" validate:"required,max=10,min=5"`
}

func ValidateStudent(student *Student) error {
	v := validator.New()
	err := v.Struct(student)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		errorMessage := make(map[string]interface{})
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage[err.Field()] = err
		}
		return fmt.Errorf("%v", errorMessage)
	}
	return nil
}
