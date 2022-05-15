package transformer

import "github.com/syke99/cloudinary/validator"

type Angle struct {
	Degrees int
	Mode    string
}

func NewAngle(degrees int, mode string) (Angle, error) {
	err := validator.ValidateAngleMode(mode)
	if err != nil {
		return Angle{}, err
	}

	return Angle{
		Degrees: degrees,
		Mode:    mode,
	}, nil
}
