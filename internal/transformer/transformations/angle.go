package transformations

import (
	"github.com/syke99/cloudinary/internal/validator"
)

type Angle struct {
	Degrees int
	Mode    string
}

func NewAngle(validator validator.Validator, degrees int, mode string) (Angle, error) {
	err := validator.ValidateAngleMode(mode)
	if err != nil {
		return Angle{}, err
	}

	return Angle{
		Degrees: degrees,
		Mode:    mode,
	}, nil
}
