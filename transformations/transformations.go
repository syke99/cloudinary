package transformations

import (
	"fmt"
	"github.com/syke99/cloudinary/image"
	"github.com/syke99/cloudinary/video"
	"reflect"
)

type Transformer struct {
	transformer
}

type transformer[Image image.Image, Video video.Video] interface {
	AddExtension(Image, Video, string) any
	AddAngle(any, Angle) any
	AddAspectRatio(any, AspectRatio) any
	AddAudioCodec(any, AudioCodec) any
	AddAudioFrequency(any, AudioFrequency) any
	AddBackground(any, Background) any
	AddBorder(any, Border) any
	AddBitrate(any, BitRate) any
	AddCropOrResize(any, CropResize) any
	AddColor(any, Color) any
	AddColorSpace(any, ColorSpace) any
	AddDefaultImage(any, DefaultImage) any
	AddDelay(any, Delay) any
	AddDensity(any, Density) any
	AddDPR(any, DPR) any
	AddDuration(any, Duration) any
	AddEffect(any, Effect) any
	AddEndOffset(any) any
	AddFormat(any) any
	AddFlag(any) any
	AddFPS(any) any
	AddCustomFunction(any) any
	AddGravity(any) any
	AddHeight(any) any
	AddIf(any) any
	AddKeyframeInterval(any) any
	AddLayer(any) any
	AddOpacity(any) any
	AddPrefix(any) any
	AddPageOrFileLayer(any) any
	AddQuality(any) any
	AddRoundCorners(any) any
	AddStartOffset(any) any
	AddStreamingProfile(any) any
	AddNamedTransformation(any) any
	AddUnderlay(any) any
	AddVideoCodec(any) any
	AddVideoSampling(any) any
	AddWidth(any) any
	AddXY(any) any
	AddZoom(any) any
	AddVariable(any) any
}

func (t Transformer) AddExtension(media any, ext string) any {
	var med any
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		reflect.ValueOf(media).FieldByName("Ext").Set(reflect.ValueOf(fmt.Sprintf(".%s", ext)))
	case video.Video:
		vid := media.(video.Video)
		med = vid
		reflect.ValueOf(media).FieldByName("Ext").Set(reflect.ValueOf(fmt.Sprintf(".%s", ext)))
	}
	return med
}

func (t Transformer) AddAngle(media any, angle Angle) any {
	var med any
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		reflect.ValueOf(media).FieldByName("Ext").Set(reflect.ValueOf(fmt.Sprintf(".%s", ext)))
	case video.Video:
		vid := media.(video.Video)
		med = vid
		reflect.ValueOf(media).FieldByName("Ext").Set(reflect.ValueOf(fmt.Sprintf(".%s", ext)))
	}
	return med
}

func (t Transformer) AddAspectRatio(media any, ar AspectRatio) any {
	var med any
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		reflect.ValueOf(media).FieldByName("Ext").Set(reflect.ValueOf(fmt.Sprintf(".%s", ar)))
	case video.Video:
		vid := media.(video.Video)
		med = vid
		reflect.ValueOf(media).FieldByName("Ext").Set(reflect.ValueOf(fmt.Sprintf(".%s", ar)))
	}
	return med
}

func (t Transformer) AddAudioCodec(media any, ac AudioCodec) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddAudioFrequency(media any, af AudioFrequency) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddBackground(media any, background Background) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddBorder(media any, border Border) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddBitrate(media any, bitrate BitRate) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddCropOrResize(media any, resize CropResize) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddColor(media any, color Color) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddColorSpace(media any, space ColorSpace) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddDefaultImage(media any, defaultImage DefaultImage) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddDelay(media any, delay Delay) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddDensity(media any, density Density) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddDPR(media any, dpr DPR) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddDuration(media any, duration Duration) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddEffect(media any, effect Effect) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddEndOffset(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddFormat(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddFlag(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddFPS(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddCustomFunction(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddGravity(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddHeight(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddIf(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddKeyframeInterval(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddLayer(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddOpacity(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddPrefix(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddPageOrFileLayer(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddQuality(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddRoundCorners(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddStartOffset(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddStreamingProfile(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddNamedTransformation(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddUnderlay(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddVideoCodec(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddVideoSampling(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddWidth(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddXY(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddZoom(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}

func (t Transformer) AddVariable(media any) any {
	var med any
	var url string
	switch media.(type) {
	case image.Image:
		img := media.(image.Image)
		med = img
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		vid := media.(video.Video)
		med = vid
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return med
}
