package validator

import (
	"github.com/syke99/cloudinary/internal/internal_resources"
	"github.com/syke99/cloudinary/pkg/resources"
)

type Validator struct {
	validator
}

type validator interface {
	ValidateDeliveryType(string) error
	ValidateAngleMode(mode string) error
	ValidateAudioCodec(codec string) error
	ValidateBackgroundAutoMode(mode string) error
	ValidateBackgroundAutoDirection(direction string) error
	ValidateColorSpaceMode(mode string) error
	ValidateFunctionType(functionType string) error
}

func (v Validator) ValidateDeliveryType(delivery string) error {
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
		return internal_resources.InvalidDeliveryType
	}

	return nil
}

func (v Validator) ValidateAngleMode(mode string) error {
	if mode != resources.AutoRight ||
		mode != resources.AutoLeft ||
		mode != resources.VFlip ||
		mode != resources.Hflip ||
		mode != resources.Ignore {
		return internal_resources.InvalidAngleMode
	}
	return nil
}

func (v Validator) ValidateAudioCodec(codec string) error {
	if codec != resources.None ||
		codec != resources.Aac ||
		codec != resources.Vorbis ||
		codec != resources.Mp3 ||
		codec != resources.Opus {
		return internal_resources.InvalidAudioCodec
	}
	return nil
}

func (v Validator) ValidateBackgroundAutoMode(mode string) error {
	if mode != resources.Border ||
		mode != resources.Predominant ||
		mode != resources.BorderContrast ||
		mode != resources.PredominantContrast ||
		mode != resources.BorderGradient ||
		mode != resources.PredominantGradientContrast ||
		mode != resources.BorderGradient ||
		mode != resources.BorderGradientContrast {
		return internal_resources.InvalidBackgroundAutoMode
	}
	return nil
}

func (v Validator) ValidateBackgroundAutoDirection(direction string) error {
	if direction != resources.Horizontal ||
		direction != resources.Vertical ||
		direction != resources.DiagonalDesc ||
		direction != resources.DiagonalAsc {
		return internal_resources.InvalidBackgroundAutoDirection
	}
	return nil
}

func (v Validator) ValidateColorSpaceMode(mode string) error {
	if mode != resources.Srgb ||
		mode != resources.TinySrgb ||
		mode != resources.Cmyk ||
		mode != resources.NoCmyk ||
		mode != resources.KeepCmyk ||
		mode != resources.SrgbTrueColor ||
		mode != resources.CsIcc ||
		mode != resources.Copy {
		return internal_resources.InvalidColorSpaceMode
	}

	return nil
}

func (v Validator) ValidateFunctionType(functionType string) error {
	if functionType != resources.Remote ||
		functionType != resources.Wasm {
		return internal_resources.InvalidCustomFunctionType
	}
	return nil
}
