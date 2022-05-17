package cloudinary

import (
	"fmt"
	"github.com/syke99/cloudinary/assets/image"
	"github.com/syke99/cloudinary/assets/video"
	"github.com/syke99/cloudinary/internal/config"
	"github.com/syke99/cloudinary/internal/transformer"
	"github.com/syke99/cloudinary/internal/transformer/transformations"
	"github.com/syke99/cloudinary/internal/validator"
	"net/http"

	"github.com/syke99/cloudinary/resources"
)

type Cloudinary struct {
	client      *http.Client
	config      config.CloudinaryConfig
	validator   validator.Validator
	transformer transformer.Transformer
	cloudinary
}

type cloudinary interface {
	Image() image.Image
	Video() video.Video
}

// NewUnsignedCloudinary creates a new instance for unsigned interactions with the Cloudinary API
// Unsigned interactions currently not supported
//func NewUnsignedCloudinary(cloud string) *Cloudinary {
//	conf := config.Config{
//		Cloud: cloud,
//	}
//	return &Cloudinary{
//		config: conf,
//	}
//}

// NewSignedCloudinary creates a new instance for signed interactions with the Cloudinary API
func NewSignedCloudinary(client *http.Client, cloud string, key string, secret string) *Cloudinary {
	conf := config.CloudinaryConfig{
		Cloud:     cloud,
		ApiKey:    key,
		ApiSecret: secret,
	}
	validator := validator.Validator{}
	transformer := transformer.Transformer{}
	return &Cloudinary{
		client:      client,
		config:      conf,
		validator:   validator,
		transformer: transformer,
	}
}

func (c *Cloudinary) Image(name string) image.Image {
	reqUrl := fmt.Sprintf("%s/%s/image/", resources.BaseUrl, c.config.Cloud)
	uploadUrl := fmt.Sprintf("%s/v1_1/%s/image/", resources.BaseUrl, c.config.Cloud)

	config := config.MediaConfig{
		Client:      c.client,
		Config:      c.config,
		Transformer: c.transformer,
		Name:        name,
		ReqUrl:      reqUrl,
		UploadUrl:   uploadUrl,
	}

	return image.Image{}.ConfigureImage(config)
}

func (c *Cloudinary) Video(name string) video.Video {
	reqUrl := fmt.Sprintf("%s/%s/video/", resources.BaseUrl, c.config.Cloud)
	uploadUrl := fmt.Sprintf("%s/v1_1/%s/image/", resources.BaseUrl, c.config.Cloud)

	config := config.MediaConfig{
		Client:      c.client,
		Config:      c.config,
		Transformer: c.transformer,
		Name:        name,
		ReqUrl:      reqUrl,
		UploadUrl:   uploadUrl,
	}

	return video.Video{}.ConfigureVideo(config)
}

func (c Cloudinary) NewAngle(degrees int, mode string) (transformations.Angle, error) {
	transformation, err := transformations.NewAngle(c.validator, degrees, mode)

	if err != nil {
		return transformations.Angle{}, err
	}

	return transformation, err
}

func (c Cloudinary) NewAspectRatio(width float32, height float32) transformations.AspectRatio {
	transformation := transformations.NewAspectRatio(width, height)

	return transformation
}

func (c Cloudinary) NewAudioCodec(codec string) (transformations.AudioCodec, error) {
	transformation, err := transformations.NewAudioCodec(c.validator, codec)

	if err != nil {
		return transformations.AudioCodec{}, err
	}

	return transformation, nil
}

func (c Cloudinary) NewAudioFrequency(frequency string, iaf bool) transformations.AudioFrequency {
	return transformations.NewAudioFrequency(frequency, iaf)
}

func (c Cloudinary) NewBackgroundAuto(mode string, number int, direction string, colors []string) (transformations.BackgroundAuto, error) {
	transformation, err := transformations.NewBackgroundAuto(c.validator, mode, number, direction, colors)
	if err != nil {
		return transformations.BackgroundAuto{}, err
	}

	return transformation, nil
}

func (c Cloudinary) NewBackgroundBlurred(intensity int, brightness uint) transformations.BackgroundBlurred {
	return transformations.NewBackgroundBlurred(intensity, brightness)
}

func (c Cloudinary) NewBackground(color string, auto transformations.BackgroundAuto, blurred transformations.BackgroundBlurred) transformations.Background {
	return transformations.NewBackground(color, auto, blurred)
}

func (c Cloudinary) NewBitRate(bitrate int, constant bool) transformations.BitRate {
	return transformations.NewBitRate(bitrate, constant)
}

func (c Cloudinary) NewBorder(width string, style string, color string) transformations.Border {
	return transformations.NewBorder(width, style, color)
}

func (c Cloudinary) NewColor(color string) transformations.Color {
	return transformations.NewColor(color)
}

func (c Cloudinary) NewColorSpace(mode string, publicId string) (transformations.ColorSpace, error) {
	transformation, err := transformations.NewColorSpace(c.validator, mode, publicId)
	if err != nil {
		return transformations.ColorSpace{}, err
	}

	return transformation, nil
}

func (c Cloudinary) NewCustomFunction(functionType string, source string) (transformations.CustomFunction, error) {
	transformation, err := transformations.NewCustomFunction(c.validator, functionType, source)
	if err != nil {
		return transformations.CustomFunction{}, err
	}

	return transformation, nil
}

func (c Cloudinary) NewDefaultImage(imgAsset string) transformations.DefaultImage {
	return transformations.NewDefaultImage(imgAsset)
}

func (c Cloudinary) NewDelay(time int) transformations.Delay {
	return transformations.NewDelay(time)
}

func (c Cloudinary) NewDensity(dpi int) transformations.Density {
	return transformations.NewDensity(dpi)
}

func (c Cloudinary) NewDPR(pixelRatio float32, auto bool) transformations.DPR {
	return transformations.NewDPR(pixelRatio, auto)
}

func (c Cloudinary) NewDuration(time float32) transformations.Duration {
	return transformations.NewDuration(time)
}
