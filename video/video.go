package video

import (
	"fmt"
	"github.com/syke99/cloudinary/request"
	"github.com/syke99/cloudinary/transformations"
	"github.com/syke99/cloudinary/validation"
)

type Video struct {
	Transformer     transformations.Transformer
	Transformations []string
	Name            string
	Ext             string
	Url             string
	video
}

type video interface {
	AddExtension(string) Video
	AddAngle(transformations.Angle) Video
	AddAudioCodec(transformations.AudioCodec) Video
	AddAudioFrequency(transformations.AudioFrequency) Video
	AddBackground(transformations.Background) Video
	AddBorder(transformations.Border) Video
	AddBitrate(transformations.BitRate) Video
	AddCropOrResize(transformations.CropResize) Video
	AddColor(transformations.Color) Video
	AddColorSpace(transformations.ColorSpace) Video
	AddDelay(transformations.Delay) Video
	AddDPR(transformations.DPR) Video
	AddDuration(transformations.Duration) Video
	AddEffect(transformations.Effect) Video
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
}

func (v Video) AddExtension(ext string) Video {
	return v.Transformer.AddExtension(v, ext).(Video)
}

func (v Video) AddAngle(angle transformations.Angle) Video {
	return v.Transformer.AddAngle(v, angle).(Video)
}

func (v Video) AddAudioCodec(ac transformations.AudioCodec) Video {
	return v.Transformer.AddAudioCodec(v, ac).(Video)
}

func (v Video) AddAudioFrequency(af transformations.AudioFrequency) Video {
	return v.Transformer.AddAudioFrequency(v, af).(Video)
}

func (v Video) AddBackground(background transformations.Background) Video {
	return v.Transformer.AddBackground(v, background).(Video)
}

func (v Video) AddBorder(border transformations.Border) Video {
	return v.Transformer.AddBorder(v, border).(Video)
}

func (v Video) AddBitrate(br transformations.BitRate) Video {
	return v.Transformer.AddBitrate(v, br).(Video)
}

func (v Video) AddCropOrResize(cr transformations.CropResize) Video {
	return v.Transformer.AddCropOrResize(v, cr).(Video)
}

func (v Video) AddColor(color transformations.Color) Video {
	return v.Transformer.AddColor(v, color).(Video)
}

func (v Video) AddColorSpace(colorSpace transformations.ColorSpace) Video {
	return v.Transformer.AddColorSpace(v, colorSpace).(Video)
}

func (v Video) AddDelay(delay transformations.Delay) Video {
	return v.Transformer.AddDelay(v, delay).(Video)
}

func (v Video) AddDPR(dpr transformations.DPR) Video {
	return v.Transformer.AddDPR(v, dpr).(Video)
}

func (v Video) AddDuration(duration transformations.Duration) Video {
	return v.Transformer.AddDuration(v, duration).(Video)
}

func (v Video) AddEffect(effect transformations.Effect) Video {
	return v.Transformer.AddEffect(v, effect).(Video)
}

func (v Video) AddEndOffset() Video {
	return v.Transformer.AddEndOffset(v).(Video)
}

func (v Video) AddFormat() Video {
	return v.Transformer.AddFormat(v).(Video)
}

func (v Video) AddFlag() Video {
	return v.Transformer.AddFlag(v).(Video)
}

func (v Video) AddFPS() Video {
	return v.Transformer.AddFPS(v).(Video)
}

func (v Video) AddGravity() Video {
	return v.Transformer.AddGravity(v).(Video)
}

func (v Video) AddHeight() Video {
	return v.Transformer.AddHeight(v).(Video)
}

func (v Video) AddIf() Video {
	return v.Transformer.AddIf(v).(Video)
}

func (v Video) AddKeyframeInterval() Video {
	return v.Transformer.AddKeyframeInterval(v).(Video)
}

func (v Video) AddLayer() Video {
	return v.Transformer.AddLayer(v).(Video)
}

func (v Video) AddQuality() Video {
	return v.Transformer.AddQuality(v).(Video)
}

func (v Video) AddRoundCorners() Video {
	return v.Transformer.AddRoundCorners(v).(Video)
}

func (v Video) AddStartOffset() Video {
	return v.Transformer.AddStartOffset(v).(Video)
}

func (v Video) AddStreamingProfile() Video {
	return v.Transformer.AddStreamingProfile(v).(Video)
}

func (v Video) AddNamedTransformation() Video {
	return v.Transformer.AddNamedTransformation(v).(Video)
}

func (v Video) AddUnderlay() Video {
	return v.Transformer.AddUnderlay(v).(Video)
}

func (v Video) AddVideoCodec() Video {
	return v.Transformer.AddVideoCodec(v).(Video)
}

func (v Video) AddVideoSampling() Video {
	return v.Transformer.AddVideoSampling(v).(Video)
}

func (v Video) AddWidth() Video {
	return v.Transformer.AddWidth(v).(Video)
}

func (v Video) AddXY() Video {
	return v.Transformer.AddXY(v).(Video)
}

func (v Video) AddVariable() Video {
	return v.Transformer.AddVariable(v).(Video)
}

func (v Video) RequestVideo(delivery string) ([]byte, error) {
	err := validation.ValidateDeliveryType(delivery)
	if err != nil {
		return []byte{}, err
	}

	r := request.Request{}
	reqUrl := fmt.Sprintf("%s/%s/%s/%s", v.Url, delivery, v.Name, v.Ext)
	return r.RequestMedia(reqUrl), nil
}
