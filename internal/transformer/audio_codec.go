package transformer

import (
	"github.com/syke99/cloudinary/internal/validator"
)

type AudioCodec struct {
	CodecValue string
}

func NewAudioCodec(validator validator.Validator, codec string) (AudioCodec, error) {
	err := validator.ValidateAudioCodec(codec)
	if err != nil {
		return AudioCodec{}, err
	}

	return AudioCodec{
		CodecValue: codec,
	}, nil
}
