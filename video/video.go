package video

import (
	"fmt"

	"github.com/syke99/cloudinary/request"
	"github.com/syke99/cloudinary/transformer"
	"github.com/syke99/cloudinary/validator"
)

type Video struct {
	Transformer     transformer.Transformer
	Transformations []string
	Name            string
	Ext             string
	Url             string
	video
}

type video interface {
	AddExtension(string) *Video
	AddAngle(transformer.Angle) *Video
	AddAudioCodec(transformer.AudioCodec) *Video
	AddAudioFrequency(transformer.AudioFrequency) *Video
	AddBackground(transformer.Background) *Video
	AddBorder(transformer.Border) *Video
	AddBitrate(transformer.BitRate) *Video
	AddCropOrResize(transformer.CropResize) *Video
	AddColor(transformer.Color) *Video
	AddColorSpace(transformer.ColorSpace) *Video
	AddDelay(transformer.Delay) *Video
	AddDPR(transformer.DPR) *Video
	AddDuration(transformer.Duration) *Video
	AddEffect(transformer.Effect) *Video
	AddEndOffset() *Video
	AddFormat() *Video
	AddFlag() *Video
	AddFPS() *Video
	AddGravity() *Video
	AddHeight() *Video
	AddIf() *Video
	AddKeyframeInterval() *Video
	AddLayer() *Video
	AddQuality() *Video
	AddRoundCorners() *Video
	AddStartOffset() *Video
	AddStreamingProfile() *Video
	AddNamedTransformation() *Video
	AddUnderlay() *Video
	AddVideoCodec() *Video
	AddVideoSampling() *Video
	AddWidth() *Video
	AddXY() *Video
	AddVariable() *Video
	RequestVideo(string) ([]byte, error)
}

func (v *Video) AddExtension(ext string) *Video {
	v.Transformations = v.Transformer.AddExtension(v.Transformations, ext)
	return v
}

func (v *Video) AddAngle(angle transformer.Angle) *Video {
	v.Transformations = v.Transformer.AddAngle(v.Transformations, angle)
	return v
}

func (v *Video) AddAudioCodec(ac transformer.AudioCodec) *Video {
	v.Transformations = v.Transformer.AddAudioCodec(v.Transformations, ac)
	return v
}

func (v *Video) AddAudioFrequency(af transformer.AudioFrequency) *Video {
	v.Transformations = v.Transformer.AddAudioFrequency(v.Transformations, af)
	return v
}

func (v *Video) AddBackground(background transformer.Background) *Video {
	v.Transformations = v.Transformer.AddBackground(v.Transformations, background)
	return v
}

func (v *Video) AddBorder(border transformer.Border) *Video {
	v.Transformations = v.Transformer.AddBorder(v.Transformations, border)
	return v
}

func (v *Video) AddBitrate(br transformer.BitRate) *Video {
	v.Transformations = v.Transformer.AddBitrate(v.Transformations, br)
	return v
}

func (v *Video) AddCropOrResize(cr transformer.CropResize) *Video {
	v.Transformations = v.Transformer.AddCropOrResize(v.Transformations, cr)
	return v
}

func (v *Video) AddColor(color transformer.Color) *Video {
	v.Transformations = v.Transformer.AddColor(v.Transformations, color)
	return v
}

func (v *Video) AddColorSpace(colorSpace transformer.ColorSpace) *Video {
	v.Transformations = v.Transformer.AddColorSpace(v.Transformations, colorSpace)
	return v
}

func (v *Video) AddDelay(delay transformer.Delay) *Video {
	v.Transformations = v.Transformer.AddDelay(v.Transformations, delay)
	return v
}

func (v *Video) AddDPR(dpr transformer.DPR) *Video {
	v.Transformations = v.Transformer.AddDPR(v.Transformations, dpr)
	return v
}

func (v *Video) AddDuration(duration transformer.Duration) *Video {
	v.Transformations = v.Transformer.AddDuration(v.Transformations, duration)
	return v
}

func (v *Video) AddEffect(effect transformer.Effect) *Video {
	v.Transformations = v.Transformer.AddEffect(v.Transformations, effect)
	return v
}

func (v *Video) AddEndOffset() *Video {
	v.Transformations = v.Transformer.AddEndOffset(v.Transformations)
	return v
}

func (v *Video) AddFormat() *Video {
	v.Transformations = v.Transformer.AddFormat(v.Transformations)
	return v
}

func (v *Video) AddFlag() *Video {
	v.Transformations = v.Transformer.AddFlag(v.Transformations)
	return v
}

func (v *Video) AddFPS() *Video {
	v.Transformations = v.Transformer.AddFPS(v.Transformations)
	return v
}

func (v *Video) AddGravity() *Video {
	v.Transformations = v.Transformer.AddGravity(v.Transformations)
	return v
}

func (v *Video) AddHeight() *Video {
	v.Transformations = v.Transformer.AddHeight(v.Transformations)
	return v
}

func (v *Video) AddIf() *Video {
	v.Transformations = v.Transformer.AddIf(v.Transformations)
	return v
}

func (v *Video) AddKeyframeInterval() *Video {
	v.Transformations = v.Transformer.AddKeyframeInterval(v.Transformations)
	return v
}

func (v *Video) AddLayer() *Video {
	v.Transformations = v.Transformer.AddLayer(v.Transformations)
	return v
}

func (v *Video) AddQuality() *Video {
	v.Transformations = v.Transformer.AddQuality(v.Transformations)
	return v
}

func (v *Video) AddRoundCorners() *Video {
	v.Transformations = v.Transformer.AddRoundCorners(v.Transformations)
	return v
}

func (v *Video) AddStartOffset() *Video {
	v.Transformations = v.Transformer.AddStartOffset(v.Transformations)
	return v
}

func (v *Video) AddStreamingProfile() *Video {
	v.Transformations = v.Transformer.AddStreamingProfile(v.Transformations)
	return v
}

func (v *Video) AddNamedTransformation() *Video {
	v.Transformations = v.Transformer.AddNamedTransformation(v.Transformations)
	return v
}

func (v *Video) AddUnderlay() *Video {
	v.Transformations = v.Transformer.AddUnderlay(v.Transformations)
	return v
}

func (v *Video) AddVideoCodec() *Video {
	v.Transformations = v.Transformer.AddVideoCodec(v.Transformations)
	return v
}

func (v *Video) AddVideoSampling() *Video {
	v.Transformations = v.Transformer.AddVideoSampling(v.Transformations)
	return v
}

func (v *Video) AddWidth() *Video {
	v.Transformations = v.Transformer.AddWidth(v.Transformations)
	return v
}

func (v *Video) AddXY() *Video {
	v.Transformations = v.Transformer.AddXY(v.Transformations)
	return v
}

func (v *Video) AddVariable() *Video {
	v.Transformations = v.Transformer.AddVariable(v.Transformations)
	return v
}

func (v *Video) RequestVideo(delivery string) ([]byte, error) {
	err := validator.ValidateDeliveryType(delivery)
	if err != nil {
		return []byte{}, err
	}

	r := request.Request{}
	reqUrl := fmt.Sprintf("%s/%s/%s/%s", v.Url, delivery, v.Name, v.Ext)
	return r.RequestMedia(reqUrl), nil
}
