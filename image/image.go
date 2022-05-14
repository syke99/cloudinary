package image

import (
	"fmt"
	"github.com/syke99/cloudinary/request"
	"github.com/syke99/cloudinary/transformations"
	"github.com/syke99/cloudinary/validation"
)

type Image struct {
	Transformations transformations.Transformations
	Name            string
	Ext             string
	Url             string
	image
}

type image interface {
	AddExtension(string) Image
	AddAngle(transformations.Angle) Image
	AddAspectRatio(transformations.AspectRatio) Image
	AddBackground(transformations.Background) Image
	AddBorder(transformations.Border) Image
	AddCropOrResize(transformations.CropResize) Image
	AddColor(transformations.Color) Image
	AddColorSpace(transformations.ColorSpace) Image
	AddDefaultImage(transformations.DefaultImage) Image
	AddDelay(transformations.Delay) Image
	AddDensity(transformations.Density) Image
	AddDPR(transformations.DPR) Image
	AddEffect(transformations.Effect) Image
	AddFormat() Image
	AddFlag() Image
	AddCustomFunction() Image
	AddGravity() Image
	AddHeight() Image
	AddIf() Image
	AddLayer() Image
	AddOpacity() Image
	AddPrefix() Image
	AddPageOrFileLayer() Image
	AddQuality() Image
	AddRoundCorners() Image
	AddNamedTransformation() Image
	AddUnderlay() Image
	AddWidth() Image
	AddXY() Image
	AddZoom() Image
	AddVariable() Image
	RequestImage(string) ([]byte, error)
}

func (i Image) AddExtension(ext string) Image {
	return i.Transformations.AddExtension(i, ext).(Image)
}

func (i Image) AddAngle(angle transformations.Angle) Image {
	return i.Transformations.AddAngle(i, angle).(Image)
}

func (i Image) AddAspectRatio(ar transformations.AspectRatio) Image {
	return i.Transformations.AddAspectRatio(i, ar).(Image)
}

func (i Image) AddBackground(background transformations.Background) Image {
	return i.Transformations.AddBackground(i, background).(Image)
}

func (i Image) AddBorder(border transformations.Border) Image {
	return i.Transformations.AddBorder(i, border).(Image)
}

func (i Image) AddCropOrResize(resize transformations.CropResize) Image {
	return i.Transformations.AddCropOrResize(i, resize).(Image)
}

func (i Image) AddColor(color transformations.Color) Image {
	return i.Transformations.AddColor(i, color).(Image)
}

func (i Image) AddColorSpace(space transformations.ColorSpace) Image {
	return i.Transformations.AddColorSpace(i, space).(Image)
}

func (i Image) AddDefaultImage(defaultImage transformations.DefaultImage) Image {
	return i.Transformations.AddDefaultImage(i, defaultImage).(Image)
}

func (i Image) AddDelay(delay transformations.Delay) Image {
	return i.Transformations.AddDelay(i, delay).(Image)
}

func (i Image) AddDensity(density transformations.Density) Image {
	return i.Transformations.AddDensity(i, density).(Image)
}

func (i Image) AddDPR(dpr transformations.DPR) Image {
	return i.Transformations.AddDPR(i, dpr).(Image)
}

func (i Image) AddEffect(effect transformations.Effect) Image {
	return i.Transformations.AddEffect(i, effect).(Image)
}

func (i Image) AddFormat() Image {
	return i.Transformations.AddFormat(i).(Image)
}

func (i Image) AddFlag() Image {
	return i.Transformations.AddFlag(i).(Image)
}

func (i Image) AddCustomFunction() Image {
	return i.Transformations.AddCustomFunction(i).(Image)
}

func (i Image) AddGravity() Image {
	return i.Transformations.AddGravity(i).(Image)
}

func (i Image) AddHeight() Image {
	return i.Transformations.AddHeight(i).(Image)
}

func (i Image) AddIf() Image {
	return i.Transformations.AddIf(i).(Image)
}

func (i Image) AddLayer() Image {
	return i.Transformations.AddLayer(i).(Image)
}

func (i Image) AddOpacity() Image {
	return i.Transformations.AddOpacity(i).(Image)
}

func (i Image) AddPrefix() Image {
	return i.Transformations.AddPrefix(i).(Image)
}

func (i Image) AddPageOrFileLayer() Image {
	return i.Transformations.AddPageOrFileLayer(i).(Image)
}

func (i Image) AddQuality() Image {
	return i.Transformations.AddQuality(i).(Image)
}

func (i Image) AddRoundCorners() Image {
	return i.Transformations.AddRoundCorners(i).(Image)
}

func (i Image) AddNamedTransformation() Image {
	return i.Transformations.AddNamedTransformation(i).(Image)
}

func (i Image) AddUnderlay() Image {
	return i.Transformations.AddUnderlay(i).(Image)
}

func (i Image) AddWidth() Image {
	return i.Transformations.AddWidth(i).(Image)
}

func (i Image) AddXY() Image {
	return i.Transformations.AddXY(i).(Image)
}

func (i Image) AddZoom() Image {
	return i.Transformations.AddZoom(i).(Image)
}

func (i Image) AddVariable() Image {
	return i.Transformations.AddVariable(i).(Image)
}

func (i Image) RequestImage(delivery string) ([]byte, error) {
	err := validation.ValidateDeliveryType(delivery)
	if err != nil {
		return []byte{}, err
	}

	r := request.Request{}
	reqUrl := fmt.Sprintf("%s/%s/%s/%s", i.Url, delivery, i.Name, i.Ext)
	return r.RequestMedia(reqUrl), nil
}
