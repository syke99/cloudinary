package transformer

import (
	"github.com/syke99/cloudinary/validator"
)

type Background struct {
	Color   string
	Auto    BackgroundAuto
	Blurred BackgroundBlurred
}

type BackgroundAuto struct {
	Mode      string
	Number    int
	Direction string
	Colors    []string
}

type BackgroundBlurred struct {
	Intensity  int
	Brightness uint
}

func NewBackgroundAuto(mode string, number int, direction string, colors []string) (BackgroundAuto, error) {
	var err error
	err = validator.ValidateBackgroundAutoMode(mode)
	if err != nil {
		return BackgroundAuto{}, err
	}

	err = validator.ValidateBackgroundAutoDirection(direction)
	if err != nil {
		return BackgroundAuto{}, err
	}

	return BackgroundAuto{
		Mode:      mode,
		Number:    number,
		Direction: direction,
		Colors:    colors,
	}, nil
}

func NewBackgroundBlurred(intensity int, brightness uint) BackgroundBlurred {
	return BackgroundBlurred{
		Intensity:  intensity,
		Brightness: brightness,
	}
}

func NewBackground(color string, bgauto BackgroundAuto, bgblurred BackgroundBlurred) Background {
	return Background{
		Color:   color,
		Auto:    bgauto,
		Blurred: bgblurred,
	}
}
