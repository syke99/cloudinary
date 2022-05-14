package transformations

import (
	"errors"
	"github.com/syke99/cloudinary/resources"
)

type Angle struct {
	Degrees int
	Mode    string
}

func NewAngle(degrees int, mode string) (Angle, error) {
	err := validateAngleMode(mode)
	if err != nil {
		return Angle{}, err
	}

	return Angle{
		Degrees: degrees,
		Mode:    mode,
	}, nil
}

func validateAngleMode(mode string) error {
	if mode != resources.AutoRight ||
		mode != resources.AutoLeft ||
		mode != resources.VFlip ||
		mode != resources.Hflip ||
		mode != resources.Ignore {
		return errors.New("angle mode provided not valid")
	}
	return nil
}
