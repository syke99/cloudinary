package transformations

import (
	"errors"
	"github.com/syke99/cloudinary/resources"
)

type AudioCodec struct {
	CodecValue string
}

func NewAudioCodec(codec string) (AudioCodec, error) {
	err := validateAudioCodec(codec)
	if err != nil {
		return AudioCodec{}, err
	}

	return AudioCodec{
		CodecValue: codec,
	}, nil
}

func validateAudioCodec(codec string) error {
	if codec != resources.None ||
		codec != resources.Aac ||
		codec != resources.Vorbis ||
		codec != resources.Mp3 ||
		codec != resources.Opus {
		return errors.New("audio codec provided not valid")
	}
	return nil
}
