package transformations

import (
	"github.com/syke99/cloudinary/internal/validator"
)

type CustomFunction struct {
	FunctionType string
	Source       string
}

func NewCustomFunction(validator validator.Validator, functionType string, source string) (CustomFunction, error) {
	err := validator.ValidateFunctionType(functionType)
	if err != nil {
		return CustomFunction{}, err
	}

	return CustomFunction{
		FunctionType: functionType,
		Source:       source,
	}, nil
}
