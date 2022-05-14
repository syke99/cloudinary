package video

import (
	"fmt"
	"github.com/syke99/cloudinary/request"
	"github.com/syke99/cloudinary/transformations"
	"github.com/syke99/cloudinary/validation"
)

type Video struct {
	Transformations transformations.Transformations
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
	return v.Transformations.AddExtension(v, ext).(Video)
}

func (v Video) AddAngle(angle transformations.Angle) Video {
	return v.Transformations.AddAngle(v, angle).(Video)
}

func (v Video) AddAudioCodec(ac transformations.AudioCodec) Video {
	return v.Transformations.AddAudioCodec(v, ac).(Video)
}

func (v Video) AddAudioFrequency(af transformations.AudioFrequency) Video {
	return v.Transformations.AddAudioFrequency(v, af).(Video)
}

func (v Video) AddBackground(background transformations.Background) Video {
	return v.Transformations.AddBackground(v, background).(Video)
}

func (v Video) AddBorder(border transformations.Border) Video {
	return v.Transformations.AddBorder(v, border).(Video)
}

func (v Video) AddBitrate(br transformations.BitRate) Video {
	return v.Transformations.AddBitrate(v, br).(Video)
}

func (v Video) AddCropOrResize(cr transformations.CropResize) Video {
	return v.Transformations.AddCropOrResize(v, cr).(Video)
}

func (v Video) AddColor(color transformations.Color) Video {
	return v.Transformations.AddColor(v, color).(Video)
}

func (v Video) AddColorSpace(colorSpace transformations.ColorSpace) Video {
	return v.Transformations.AddColorSpace(v, colorSpace).(Video)
}

func (v Video) AddDelay(delay transformations.Delay) Video {
	return v.Transformations.AddDelay(v, delay).(Video)
}

func (v Video) AddDPR(dpr transformations.DPR) Video {
	return v.Transformations.AddDPR(v, dpr).(Video)
}

func (v Video) AddDuration(duration transformations.Duration) Video {
	return v.Transformations.AddDuration(v, duration).(Video)
}

func (v Video) AddEffect(effect transformations.Effect) Video {
	return v.Transformations.AddEffect(v, effect).(Video)
}

func (v Video) AddEndOffset() Video {
	return v.Transformations.AddEndOffset(v).(Video)
}

func (v Video) AddFormat() Video {
	return v.Transformations.AddFormat(v).(Video)
}

func (v Video) AddFlag() Video {
	return v.Transformations.AddFlag(v).(Video)
}

func (v Video) AddFPS() Video {
	return v.Transformations.AddFPS(v).(Video)
}

func (v Video) AddGravity() Video {
	return v.Transformations.AddGravity(v).(Video)
}

func (v Video) AddHeight() Video {
	return v.Transformations.AddHeight(v).(Video)
}

func (v Video) AddIf() Video {
	return v.Transformations.AddIf(v).(Video)
}

func (v Video) AddKeyframeInterval() Video {
	return v.Transformations.AddKeyframeInterval(v).(Video)
}

func (v Video) AddLayer() Video {
	return v.Transformations.AddLayer(v).(Video)
}

func (v Video) AddQuality() Video {
	return v.Transformations.AddQuality(v).(Video)
}

func (v Video) AddRoundCorners() Video {
	return v.Transformations.AddRoundCorners(v).(Video)
}

func (v Video) AddStartOffset() Video {
	return v.Transformations.AddStartOffset(v).(Video)
}

func (v Video) AddStreamingProfile() Video {
	return v.Transformations.AddStreamingProfile(v).(Video)
}

func (v Video) AddNamedTransformation() Video {
	return v.Transformations.AddNamedTransformation(v).(Video)
}

func (v Video) AddUnderlay() Video {
	return v.Transformations.AddUnderlay(v).(Video)
}

func (v Video) AddVideoCodec() Video {
	return v.Transformations.AddVideoCodec(v).(Video)
}

func (v Video) AddVideoSampling() Video {
	return v.Transformations.AddVideoSampling(v).(Video)
}

func (v Video) AddWidth() Video {
	return v.Transformations.AddWidth(v).(Video)
}

func (v Video) AddXY() Video {
	return v.Transformations.AddXY(v).(Video)
}

func (v Video) AddVariable() Video {
	return v.Transformations.AddVariable(v).(Video)
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
