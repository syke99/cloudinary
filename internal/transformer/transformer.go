package transformer

import "github.com/syke99/cloudinary/internal/transformer/transformations"

type Transformer struct {
	transformer
}

type transformer interface {
	AddExtension([]string, string) []string
	AddAngle([]string, transformations.Angle) []string
	AddAspectRatio([]string, transformations.AspectRatio) []string
	AddAudioCodec([]string, transformations.AudioCodec) []string
	AddAudioFrequency([]string, transformations.AudioFrequency) []string
	AddBackground([]string, transformations.Background) []string
	AddBorder([]string, transformations.Border) []string
	AddBitrate([]string, transformations.BitRate) []string
	AddCropOrResize([]string, transformations.CropResize) []string
	AddColor([]string, transformations.Color) []string
	AddColorSpace([]string, transformations.ColorSpace) []string
	AddDefaultImage([]string, transformations.DefaultImage) []string
	AddDelay([]string, transformations.Delay) []string
	AddDensity([]string, transformations.Density) []string
	AddDPR([]string, transformations.DPR) []string
	AddDuration([]string, transformations.Duration) []string
	AddEffect([]string, transformations.Effect) []string
	AddEndOffset([]string) []string
	AddFormat([]string) []string
	AddFlag([]string) []string
	AddFPS([]string) []string
	AddCustomFunction([]string) []string
	AddGravity([]string) []string
	AddHeight([]string) []string
	AddIf([]string) []string
	AddKeyframeInterval([]string) []string
	AddLayer([]string) []string
	AddOpacity([]string) []string
	AddPrefix([]string) []string
	AddPageOrFileLayer([]string) []string
	AddQuality([]string) []string
	AddRoundCorners([]string) []string
	AddStartOffset([]string) []string
	AddStreamingProfile([]string) []string
	AddNamedTransformation([]string) []string
	AddUnderlay([]string) []string
	AddVideoCodec([]string) []string
	AddVideoSampling([]string) []string
	AddWidth([]string) []string
	AddXY([]string) []string
	AddZoom([]string) []string
	AddVariable([]string) []string
}

func (t Transformer) AddExtension(transformations []string, ext string) []string {
	var med []string
	return med
}

func (t Transformer) AddAngle(transformations []string, angle transformations.Angle) []string {
	var med []string
	return med
}

func (t Transformer) AddAspectRatio(transformations []string, ar transformations.AspectRatio) []string {
	var med []string
	return med
}

func (t Transformer) AddAudioCodec(transformations []string, ac transformations.AudioCodec) []string {
	var med []string
	return med
}

func (t Transformer) AddAudioFrequency(transformations []string, af transformations.AudioFrequency) []string {
	var med []string
	return med
}

func (t Transformer) AddBackground(transformations []string, background transformations.Background) []string {
	var med []string
	return med
}

func (t Transformer) AddBorder(transformations []string, border transformations.Border) []string {
	var med []string
	return med
}

func (t Transformer) AddBitrate(transformations []string, bitrate transformations.BitRate) []string {
	var med []string
	return med
}

func (t Transformer) AddCropOrResize(transformations []string, resize transformations.CropResize) []string {
	var med []string
	return med
}

func (t Transformer) AddColor(transformations []string, color transformations.Color) []string {
	var med []string
	return med
}

func (t Transformer) AddColorSpace(transformations []string, space transformations.ColorSpace) []string {
	var med []string
	return med
}

func (t Transformer) AddDefaultImage(transformations []string, defaultImage transformations.DefaultImage) []string {
	var med []string
	return med
}

func (t Transformer) AddDelay(transformations []string, delay transformations.Delay) []string {
	var med []string
	return med
}

func (t Transformer) AddDensity(transformations []string, density transformations.Density) []string {
	var med []string
	return med
}

func (t Transformer) AddDPR(transformations []string, dpr transformations.DPR) []string {
	var med []string
	return med
}

func (t Transformer) AddDuration(transformations []string, duration transformations.Duration) []string {
	var med []string
	return med
}

func (t Transformer) AddEffect(transformations []string, effect transformations.Effect) []string {
	var med []string
	return med
}

func (t Transformer) AddEndOffset(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddFormat(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddFlag(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddFPS(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddCustomFunction(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddGravity(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddHeight(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddIf(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddKeyframeInterval(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddLayer(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddOpacity(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddPrefix(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddPageOrFileLayer(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddQuality(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddRoundCorners(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddStartOffset(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddStreamingProfile(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddNamedTransformation(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddUnderlay(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddVideoCodec(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddVideoSampling(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddWidth(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddXY(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddZoom(transformations []string) []string {
	var med []string
	return med
}

func (t Transformer) AddVariable(transformations []string) []string {
	var med []string
	return med
}
