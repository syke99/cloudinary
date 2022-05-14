package transformations

import (
	"errors"
	"github.com/syke99/cloudinary/resources"
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
	err = validateBackgroundAutoMode(mode)
	if err != nil {
		return BackgroundAuto{}, err
	}

	err = validateBackgroundAutoDirection(direction)
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

func validateBackgroundAutoMode(mode string) error {
	if mode != resources.Border ||
		mode != resources.Predominant ||
		mode != resources.BorderContrast ||
		mode != resources.PredominantContrast ||
		mode != resources.BorderGradient ||
		mode != resources.PredominantGradientContrast ||
		mode != resources.BorderGradient ||
		mode != resources.BorderGradientContrast {
		return errors.New("backgroundauto mode provided not valid")
	}
	return nil
}

func validateBackgroundAutoDirection(direction string) error {
	if direction != resources.Horizontal ||
		direction != resources.Vertical ||
		direction != resources.DiagonalDesc ||
		direction != resources.DiagonalAsc {
		return errors.New("backgroundauto direction provided not valid")
	}
	return nil
}
