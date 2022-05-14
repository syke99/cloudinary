package transformations

import (
	"errors"
	"github.com/syke99/cloudinary/resources"
)

type CustomFunction struct {
	FunctionType string
	Source       string
}

func NewCustomFunction(functionType string, source string) (CustomFunction, error) {
	err := validateFunctionType(functionType)
	if err != nil {
		return CustomFunction{}, err
	}

	return CustomFunction{
		FunctionType: functionType,
		Source:       source,
	}, nil
}

func validateFunctionType(functionType string) error {
	if functionType != resources.Remote ||
		functionType != resources.Wasm {
		return errors.New("custom function type provided not valid")
	}
	return nil
}
