package video

import (
	"fmt"
	"github.com/syke99/cloudinary/resources"
	"github.com/syke99/cloudinary/upload"
	"reflect"

	"github.com/syke99/cloudinary/config"
	"github.com/syke99/cloudinary/request"
	"github.com/syke99/cloudinary/transformer"
	"github.com/syke99/cloudinary/validator"
)

type Video struct {
	config          config.Config
	Transformer     transformer.Transformer
	transformations []string
	Name            string
	Ext             string
	Url             string
	video
}

type video interface {
	NewVideo(string, string, transformer.Transformer, config.Config) Video
	AddExtension(string) Video
	AddAngle(transformer.Angle) Video
	AddAudioCodec(transformer.AudioCodec) Video
	AddAudioFrequency(transformer.AudioFrequency) Video
	AddBackground(transformer.Background) Video
	AddBorder(transformer.Border) Video
	AddBitrate(transformer.BitRate) Video
	AddCropOrResize(transformer.CropResize) Video
	AddColor(transformer.Color) Video
	AddColorSpace(transformer.ColorSpace) Video
	AddDelay(transformer.Delay) Video
	AddDPR(transformer.DPR) Video
	AddDuration(transformer.Duration) Video
	AddEffect(transformer.Effect) Video
	AddEndOffset() Video
	AddFormat() Video
	AddFlag() Video
	AddFPS() Video
	AddGravity() Video
	AddHeight() Video
	AddIf() Video
	AddKeyframeInterval() Video
	AddLayer() Video
	AddQuality() Video
	AddRoundCorners() Video
	AddStartOffset() Video
	AddStreamingProfile() Video
	AddNamedTransformation() Video
	AddUnderlay() Video
	AddVideoCodec() Video
	AddVideoSampling() Video
	AddWidth() Video
	AddXY() Video
	AddVariable() Video
	RequestVideo(string) ([]byte, error)
	UploadVideo(upload.UploaderParameters, string) interface{}
}

func (v Video) NewVideo(name string, url string, transformer transformer.Transformer, config config.Config) Video {
	v.config = config
	v.Transformer = transformer
	v.transformations = []string{}
	v.Name = name
	v.Ext = ""
	v.Url = url
	return v
}

func (v Video) AddExtension(ext string) Video {
	v.transformations = v.Transformer.AddExtension(v.transformations, ext)
	return v
}

func (v Video) AddAngle(angle transformer.Angle) Video {
	v.transformations = v.Transformer.AddAngle(v.transformations, angle)
	return v
}

func (v Video) AddAudioCodec(ac transformer.AudioCodec) Video {
	v.transformations = v.Transformer.AddAudioCodec(v.transformations, ac)
	return v
}

func (v Video) AddAudioFrequency(af transformer.AudioFrequency) Video {
	v.transformations = v.Transformer.AddAudioFrequency(v.transformations, af)
	return v
}

func (v Video) AddBackground(background transformer.Background) Video {
	v.transformations = v.Transformer.AddBackground(v.transformations, background)
	return v
}

func (v Video) AddBorder(border transformer.Border) Video {
	v.transformations = v.Transformer.AddBorder(v.transformations, border)
	return v
}

func (v Video) AddBitrate(br transformer.BitRate) Video {
	v.transformations = v.Transformer.AddBitrate(v.transformations, br)
	return v
}

func (v Video) AddCropOrResize(cr transformer.CropResize) Video {
	v.transformations = v.Transformer.AddCropOrResize(v.transformations, cr)
	return v
}

func (v Video) AddColor(color transformer.Color) Video {
	v.transformations = v.Transformer.AddColor(v.transformations, color)
	return v
}

func (v Video) AddColorSpace(colorSpace transformer.ColorSpace) Video {
	v.transformations = v.Transformer.AddColorSpace(v.transformations, colorSpace)
	return v
}

func (v Video) AddDelay(delay transformer.Delay) Video {
	v.transformations = v.Transformer.AddDelay(v.transformations, delay)
	return v
}

func (v Video) AddDPR(dpr transformer.DPR) Video {
	v.transformations = v.Transformer.AddDPR(v.transformations, dpr)
	return v
}

func (v Video) AddDuration(duration transformer.Duration) Video {
	v.transformations = v.Transformer.AddDuration(v.transformations, duration)
	return v
}

func (v Video) AddEffect(effect transformer.Effect) Video {
	v.transformations = v.Transformer.AddEffect(v.transformations, effect)
	return v
}

func (v Video) AddEndOffset() Video {
	v.transformations = v.Transformer.AddEndOffset(v.transformations)
	return v
}

func (v Video) AddFormat() Video {
	v.transformations = v.Transformer.AddFormat(v.transformations)
	return v
}

func (v Video) AddFlag() Video {
	v.transformations = v.Transformer.AddFlag(v.transformations)
	return v
}

func (v Video) AddFPS() Video {
	v.transformations = v.Transformer.AddFPS(v.transformations)
	return v
}

func (v Video) AddGravity() Video {
	v.transformations = v.Transformer.AddGravity(v.transformations)
	return v
}

func (v Video) AddHeight() Video {
	v.transformations = v.Transformer.AddHeight(v.transformations)
	return v
}

func (v Video) AddIf() Video {
	v.transformations = v.Transformer.AddIf(v.transformations)
	return v
}

func (v Video) AddKeyframeInterval() Video {
	v.transformations = v.Transformer.AddKeyframeInterval(v.transformations)
	return v
}

func (v Video) AddLayer() Video {
	v.transformations = v.Transformer.AddLayer(v.transformations)
	return v
}

func (v Video) AddQuality() Video {
	v.transformations = v.Transformer.AddQuality(v.transformations)
	return v
}

func (v Video) AddRoundCorners() Video {
	v.transformations = v.Transformer.AddRoundCorners(v.transformations)
	return v
}

func (v Video) AddStartOffset() Video {
	v.transformations = v.Transformer.AddStartOffset(v.transformations)
	return v
}

func (v Video) AddStreamingProfile() Video {
	v.transformations = v.Transformer.AddStreamingProfile(v.transformations)
	return v
}

func (v Video) AddNamedTransformation() Video {
	v.transformations = v.Transformer.AddNamedTransformation(v.transformations)
	return v
}

func (v Video) AddUnderlay() Video {
	v.transformations = v.Transformer.AddUnderlay(v.transformations)
	return v
}

func (v Video) AddVideoCodec() Video {
	v.transformations = v.Transformer.AddVideoCodec(v.transformations)
	return v
}

func (v Video) AddVideoSampling() Video {
	v.transformations = v.Transformer.AddVideoSampling(v.transformations)
	return v
}

func (v Video) AddWidth() Video {
	v.transformations = v.Transformer.AddWidth(v.transformations)
	return v
}

func (v Video) AddXY() Video {
	v.transformations = v.Transformer.AddXY(v.transformations)
	return v
}

func (v Video) AddVariable() Video {
	v.transformations = v.Transformer.AddVariable(v.transformations)
	return v
}

func (v Video) RequestVideo(delivery string) ([]byte, error) {
	err := validator.ValidateDeliveryType(delivery)
	if err != nil {
		return []byte{}, err
	}

	r := request.Request{}
	reqUrl := fmt.Sprintf("%s/%s/%s/%s", v.Url, delivery, v.Name, v.Ext)
	return r.RequestMedia(reqUrl), nil
}

func (v Video) UploadVideo(params upload.UploaderParameters, apiKey string) (interface{}, error) {
	if reflect.ValueOf(params).Len() == 0 {
		return nil, resources.NoUploadParamsSupplied
	}

	for _, transformation := range v.transformations {
		params.Transformation += transformation
	}
	u := upload.Uploader{}

	sortedParams := u.SortUploadParameters(params)

	signature := u.GenerateSignature(sortedParams, apiKey)

	// just a placeholder, will be removed once u.UploadMedia
	// is updated with correct call to Cloudinary API
	println(signature)

	return u.UploadMedia(params), nil
}
