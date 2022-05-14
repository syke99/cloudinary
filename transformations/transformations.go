package transformations

import (
	"fmt"
	"github.com/syke99/cloudinary/image"
	"github.com/syke99/cloudinary/video"
	"reflect"
)

type Transformations struct {
	transformations
}

type transformations interface {
	AddExtension(interface{}, string) interface{}
	AddAngle(interface{}, Angle) interface{}
	AddAspectRatio(interface{}, AspectRatio) interface{}
	AddAudioCodec(interface{}, AudioCodec) interface{}
	AddAudioFrequency(interface{}, AudioFrequency) interface{}
	AddBackground(interface{}, Background) interface{}
	AddBorder(interface{}, Border) interface{}
	AddBitrate(interface{}, BitRate) interface{}
	AddCropOrResize(interface{}, CropResize) interface{}
	AddColor(interface{}, Color) interface{}
	AddColorSpace(interface{}, ColorSpace) interface{}
	AddDefaultImage(interface{}, DefaultImage) interface{}
	AddDelay(interface{}, Delay) interface{}
	AddDensity(interface{}, Density) interface{}
	AddDPR(interface{}, DPR) interface{}
	AddDuration(interface{}, Duration) interface{}
	AddEffect(interface{}, Effect) interface{}
	AddEndOffset(interface{}) interface{}
	AddFormat(interface{}) interface{}
	AddFlag(interface{}) interface{}
	AddFPS(interface{}) interface{}
	AddCustomFunction(interface{}) interface{}
	AddGravity(interface{}) interface{}
	AddHeight(interface{}) interface{}
	AddIf(interface{}) interface{}
	AddKeyframeInterval(interface{}) interface{}
	AddLayer(interface{}) interface{}
	AddOpacity(interface{}) interface{}
	AddPrefix(interface{}) interface{}
	AddPageOrFileLayer(interface{}) interface{}
	AddQuality(interface{}) interface{}
	AddRoundCorners(interface{}) interface{}
	AddStartOffset(interface{}) interface{}
	AddStreamingProfile(interface{}) interface{}
	AddNamedTransformation(interface{}) interface{}
	AddUnderlay(interface{}) interface{}
	AddVideoCodec(interface{}) interface{}
	AddVideoSampling(interface{}) interface{}
	AddWidth(interface{}) interface{}
	AddXY(interface{}) interface{}
	AddZoom(interface{}) interface{}
	AddVariable(interface{}) interface{}
}

func (t Transformations) AddExtension(media interface{}, ext string) interface{} {
	switch media.(type) {
	case image.Image:
		reflect.ValueOf(media).FieldByName("Ext").Set(reflect.ValueOf(fmt.Sprintf(".%s", ext)))
	case video.Video:
		reflect.ValueOf(media).FieldByName("Ext").Set(reflect.ValueOf(fmt.Sprintf(".%s", ext)))
	}
	return media
}

func (t Transformations) AddAngle(media interface{}, angle Angle) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddAspectRatio(media interface{}, ar AspectRatio) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddAudioCodec(media interface{}, ac AudioCodec) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddAudioFrequency(media interface{}, af AudioFrequency) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddBackground(media interface{}, background Background) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddBorder(media interface{}, border Border) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddBitrate(media interface{}, bitrate BitRate) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddCropOrResize(media interface{}, resize CropResize) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddColor(media interface{}, color Color) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddColorSpace(media interface{}, space ColorSpace) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddDefaultImage(media interface{}, defaultImage DefaultImage) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddDelay(media interface{}, delay Delay) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddDensity(media interface{}, density Density) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddDPR(media interface{}, dpr DPR) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddDuration(media interface{}, duration Duration) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddEffect(media interface{}, effect Effect) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddEndOffset(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddFormat(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddFlag(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddFPS(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddCustomFunction(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddGravity(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddHeight(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddIf(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddKeyframeInterval(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddLayer(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddOpacity(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddPrefix(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddPageOrFileLayer(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddQuality(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddRoundCorners(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddStartOffset(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddStreamingProfile(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddNamedTransformation(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddUnderlay(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddVideoCodec(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddVideoSampling(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddWidth(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddXY(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddZoom(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}

func (t Transformations) AddVariable(media interface{}) interface{} {
	var url string
	switch media.(type) {
	case image.Image:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	case video.Video:
		url = reflect.ValueOf(media).FieldByName("Url").String()
	}
	//this is just a placeholder to curb errors during development, will be removed
	println(url)
	return media
}
