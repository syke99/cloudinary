package cloudinary

import (
	"github.com/syke99/cloudinary/internal/transformer"
)

func (c Cloudinary) NewAngle(degrees int, mode string) (transformer.Angle, error) {
	transformation, err := transformer.NewAngle(c.validator, degrees, mode)

	if err != nil {
		return transformer.Angle{}, err
	}

	return transformation, err
}

func (c Cloudinary) NewAspectRatio(width float32, height float32) transformer.AspectRatio {
	transformation := transformer.NewAspectRatio(width, height)

	return transformation
}

func (c Cloudinary) NewAudioCodec(codec string) (transformer.AudioCodec, error) {
	transformation, err := transformer.NewAudioCodec(c.validator, codec)

	if err != nil {
		return transformer.AudioCodec{}, err
	}

	return transformation, nil
}

func (c Cloudinary) NewAudioFrequency(frequency string, iaf bool) transformer.AudioFrequency {
	return transformer.NewAudioFrequency(frequency, iaf)
}

func (c Cloudinary) NewBackgroundAuto(mode string, number int, direction string, colors []string) (transformer.BackgroundAuto, error) {
	transformation, err := transformer.NewBackgroundAuto(c.validator, mode, number, direction, colors)
	if err != nil {
		return transformer.BackgroundAuto{}, err
	}

	return transformation, nil
}

func (c Cloudinary) NewBackgroundBlurred(intensity int, brightness uint) transformer.BackgroundBlurred {
	return transformer.NewBackgroundBlurred(intensity, brightness)
}

func (c Cloudinary) NewBackground(color string, auto transformer.BackgroundAuto, blurred transformer.BackgroundBlurred) transformer.Background {
	return transformer.NewBackground(color, auto, blurred)
}

func (c Cloudinary) NewBitRate(bitrate int, constant bool) transformer.BitRate {
	return transformer.NewBitRate(bitrate, constant)
}

func (c Cloudinary) NewBorder(width string, style string, color string) transformer.Border {
	return transformer.NewBorder(width, style, color)
}

func (c Cloudinary) NewColor(color string) transformer.Color {
	return transformer.NewColor(color)
}

func (c Cloudinary) NewColorSpace(mode string, publicId string) (transformer.ColorSpace, error) {
	transformation, err := transformer.NewColorSpace(c.validator, mode, publicId)
	if err != nil {
		return transformer.ColorSpace{}, err
	}

	return transformation, nil
}

func (c Cloudinary) NewCustomFunction(functionType string, source string) (transformer.CustomFunction, error) {
	transformation, err := transformer.NewCustomFunction(c.validator, functionType, source)
	if err != nil {
		return transformer.CustomFunction{}, err
	}

	return transformation, nil
}

func (c Cloudinary) NewDefaultImage(imgAsset string) transformer.DefaultImage {
	return transformer.NewDefaultImage(imgAsset)
}

func (c Cloudinary) NewDelay(time int) transformer.Delay {
	return transformer.NewDelay(time)
}

func (c Cloudinary) NewDensity(dpi int) transformer.Density {
	return transformer.NewDensity(dpi)
}

func (c Cloudinary) NewDPR(pixelRatio float32, auto bool) transformer.DPR {
	return transformer.NewDPR(pixelRatio, auto)
}

func (c Cloudinary) NewDuration(time float32) transformer.Duration {
	return transformer.NewDuration(time)
}
