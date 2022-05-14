package transformations

import (
	"errors"
	"fmt"
	"github.com/syke99/cloudinary/resources"
)

type ColorSpace struct {
	Mode string
}

func NewColorSpace(mode string, publicId string) (ColorSpace, error) {
	err := validateColorSpaceMode(mode)
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

func validateColorSpaceMode(mode string) error {
	if mode != resources.Srgb ||
		mode != resources.TinySrgb ||
		mode != resources.Cmyk ||
		mode != resources.NoCmyk ||
		mode != resources.KeepCmyk ||
		mode != resources.SrgbTrueColor ||
		mode != resources.CsIcc ||
		mode != resources.Copy {
		return errors.New("colorspace mode provided invalid")
	}

	return nil
}
