package transformer

import "github.com/syke99/cloudinary/validator"

type CustomFunction struct {
	FunctionType string
	Source       string
}

func NewCustomFunction(functionType string, source string) (CustomFunction, error) {
	err := validator.ValidateFunctionType(functionType)
	if err != nil {
		return CustomFunction{}, err
	}

	return CustomFunction{
		FunctionType: functionType,
		Source:       source,
	}, nil
}