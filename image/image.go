package image

import (
	"fmt"

	"github.com/syke99/cloudinary/request"
	"github.com/syke99/cloudinary/transformations"
	"github.com/syke99/cloudinary/validation"
)

type Image struct {
	Transformer     transformations.Transformer
	Transformations []string
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
	return i.Transformer.AddExtension(i, ext).(Image)
}

func (i Image) AddAngle(angle transformations.Angle) Image {
	return i.Transformer.AddAngle(i, angle).(Image)
}

func (i Image) AddAspectRatio(ar transformations.AspectRatio) Image {
	return i.Transformer.AddAspectRatio(i, ar).(Image)
}

func (i Image) AddBackground(background transformations.Background) Image {
	return i.Transformer.AddBackground(i, background).(Image)
}

func (i Image) AddBorder(border transformations.Border) Image {
	return i.Transformer.AddBorder(i, border).(Image)
}

func (i Image) AddCropOrResize(resize transformations.CropResize) Image {
	return i.Transformer.AddCropOrResize(i, resize).(Image)
}

func (i Image) AddColor(color transformations.Color) Image {
	return i.Transformer.AddColor(i, color).(Image)
}

func (i Image) AddColorSpace(space transformations.ColorSpace) Image {
	return i.Transformer.AddColorSpace(i, space).(Image)
}

func (i Image) AddDefaultImage(defaultImage transformations.DefaultImage) Image {
	return i.Transformer.AddDefaultImage(i, defaultImage).(Image)
}

func (i Image) AddDelay(delay transformations.Delay) Image {
	return i.Transformer.AddDelay(i, delay).(Image)
}

func (i Image) AddDensity(density transformations.Density) Image {
	return i.Transformer.AddDensity(i, density).(Image)
}

func (i Image) AddDPR(dpr transformations.DPR) Image {
	return i.Transformer.AddDPR(i, dpr).(Image)
}

func (i Image) AddEffect(effect transformations.Effect) Image {
	return i.Transformer.AddEffect(i, effect).(Image)
}

func (i Image) AddFormat() Image {
	return i.Transformer.AddFormat(i).(Image)
}

func (i Image) AddFlag() Image {
	return i.Transformer.AddFlag(i).(Image)
}

func (i Image) AddCustomFunction() Image {
	return i.Transformer.AddCustomFunction(i).(Image)
}

func (i Image) AddGravity() Image {
	return i.Transformer.AddGravity(i).(Image)
}

func (i Image) AddHeight() Image {
	return i.Transformer.AddHeight(i).(Image)
}

func (i Image) AddIf() Image {
	return i.Transformer.AddIf(i).(Image)
}

func (i Image) AddLayer() Image {
	return i.Transformer.AddLayer(i).(Image)
}

func (i Image) AddOpacity() Image {
	return i.Transformer.AddOpacity(i).(Image)
}

func (i Image) AddPrefix() Image {
	return i.Transformer.AddPrefix(i).(Image)
}

func (i Image) AddPageOrFileLayer() Image {
	return i.Transformer.AddPageOrFileLayer(i).(Image)
}

func (i Image) AddQuality() Image {
	return i.Transformer.AddQuality(i).(Image)
}

func (i Image) AddRoundCorners() Image {
	return i.Transformer.AddRoundCorners(i).(Image)
}

func (i Image) AddNamedTransformation() Image {
	return i.Transformer.AddNamedTransformation(i).(Image)
}

func (i Image) AddUnderlay() Image {
	return i.Transformer.AddUnderlay(i).(Image)
}

func (i Image) AddWidth() Image {
	return i.Transformer.AddWidth(i).(Image)
}

func (i Image) AddXY() Image {
	return i.Transformer.AddXY(i).(Image)
}

func (i Image) AddZoom() Image {
	return i.Transformer.AddZoom(i).(Image)
}

func (i Image) AddVariable() Image {
	return i.Transformer.AddVariable(i).(Image)
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
