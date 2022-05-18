package transformations

import (
	"errors"
	"fmt"
	"github.com/syke99/cloudinary/internal/validator"
	"github.com/syke99/cloudinary/pkg/resources"
)

type ColorSpace struct {
	Mode string
}

func NewColorSpace(validator validator.Validator, mode string, publicId string) (ColorSpace, error) {
	err := validator.ValidateColorSpaceMode(mode)
	if err != nil {
		return ColorSpace{}, err
	}

	if mode == resources.CsIcc && publicId == "" {
		return ColorSpace{}, errors.New(fmt.Sprintf("%s used but no public id passed in, which must be supplied to be valid", resources.CsIcc))
	}

	if mode == resources.CsIcc && publicId != "" {
		m := mode + publicId

		return ColorSpace{
			Mode: m,
		}, nil
	}

	return ColorSpace{
		Mode: mode,
	}, nil
}
