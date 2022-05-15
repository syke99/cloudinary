package validator

import (
	"errors"

	"github.com/syke99/cloudinary/resources"
)

func ValidateDeliveryType(delivery string) error {
	if delivery != resources.Upload ||
		delivery != resources.Private ||
		delivery != resources.Authenticated ||
		delivery != resources.List ||
		delivery != resources.Fetch ||
		delivery != resources.Facebook ||
		delivery != resources.Twitter ||
		delivery != resources.TwitterName ||
		delivery != resources.Gravatar ||
		delivery != resources.YouTube ||
		delivery != resources.Hulu ||
		delivery != resources.Vimeo ||
		delivery != resources.Animoto ||
		delivery != resources.WorldStarHipHop ||
		delivery != resources.DailyMotion ||
		delivery != resources.Multi ||
		delivery != resources.Delivery {
		return resources.InvalidDeliveryType
	}

	return nil
}

func ValidateAngleMode(mode string) error {
	if mode != resources.AutoRight ||
		mode != resources.AutoLeft ||
		mode != resources.VFlip ||
		mode != resources.Hflip ||
		mode != resources.Ignore {
		return errors.New("angle mode provided not valid")
	}
	return nil
}

func ValidateAudioCodec(codec string) error {
	if codec != resources.None ||
		codec != resources.Aac ||
		codec != resources.Vorbis ||
		codec != resources.Mp3 ||
		codec != resources.Opus {
		return errors.New("audio codec provided not valid")
	}
	return nil
}

func ValidateBackgroundAutoMode(mode string) error {
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

func ValidateBackgroundAutoDirection(direction string) error {
	if direction != resources.Horizontal ||
		direction != resources.Vertical ||
		direction != resources.DiagonalDesc ||
		direction != resources.DiagonalAsc {
		return errors.New("backgroundauto direction provided not valid")
	}
	return nil
}

func ValidateColorSpaceMode(mode string) error {
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

func ValidateFunctionType(functionType string) error {
	if functionType != resources.Remote ||
		functionType != resources.Wasm {
		return errors.New("custom function type provided not valid")
	}
	return nil
}