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
	i.Transformations = i.Transformer.AddExtension(i.Transformations, ext)
	return i
}

func (i Image) AddAngle(angle transformations.Angle) Image {
	i.Transformations = i.Transformer.AddAngle(i.Transformations, angle)
	return i
}

func (i Image) AddAspectRatio(ar transformations.AspectRatio) Image {
	i.Transformations = i.Transformer.AddAspectRatio(i.Transformations, ar)
	return i
}

func (i Image) AddBackground(background transformations.Background) Image {
	i.Transformations = i.Transformer.AddBackground(i.Transformations, background)
	return i
}

func (i Image) AddBorder(border transformations.Border) Image {
	i.Transformations = i.Transformer.AddBorder(i.Transformations, border)
	return i
}

func (i Image) AddCropOrResize(resize transformations.CropResize) Image {
	i.Transformations = i.Transformer.AddCropOrResize(i.Transformations, resize)
	return i
}

func (i Image) AddColor(color transformations.Color) Image {
	i.Transformations = i.Transformer.AddColor(i.Transformations, color)
	return i
}

func (i Image) AddColorSpace(space transformations.ColorSpace) Image {
	i.Transformations = i.Transformer.AddColorSpace(i.Transformations, space)
	return i
}

func (i Image) AddDefaultImage(defaultImage transformations.DefaultImage) Image {
	i.Transformations = i.Transformer.AddDefaultImage(i.Transformations, defaultImage)
	return i
}

func (i Image) AddDelay(delay transformations.Delay) Image {
	i.Transformations = i.Transformer.AddDelay(i.Transformations, delay)
	return i
}

func (i Image) AddDensity(density transformations.Density) Image {
	i.Transformations = i.Transformer.AddDensity(i.Transformations, density)
	return i
}

func (i Image) AddDPR(dpr transformations.DPR) Image {
	i.Transformations = i.Transformer.AddDPR(i.Transformations, dpr)
	return i
}

func (i Image) AddEffect(effect transformations.Effect) Image {
	i.Transformations = i.Transformer.AddEffect(i.Transformations, effect)
	return i
}

func (i Image) AddFormat() Image {
	i.Transformations = i.Transformer.AddFormat(i.Transformations)
	return i
}

func (i Image) AddFlag() Image {
	i.Transformations = i.Transformer.AddFlag(i.Transformations)
	return i
}

func (i Image) AddCustomFunction() Image {
	i.Transformations = i.Transformer.AddCustomFunction(i.Transformations)
	return i
}

func (i Image) AddGravity() Image {
	i.Transformations = i.Transformer.AddGravity(i.Transformations)
	return i
}

func (i Image) AddHeight() Image {
	i.Transformations = i.Transformer.AddHeight(i.Transformations)
	return i
}

func (i Image) AddIf() Image {
	i.Transformations = i.Transformer.AddIf(i.Transformations)
	return i
}

func (i Image) AddLayer() Image {
	i.Transformations = i.Transformer.AddLayer(i.Transformations)
	return i
}

func (i Image) AddOpacity() Image {
	i.Transformations = i.Transformer.AddOpacity(i.Transformations)
	return i
}

func (i Image) AddPrefix() Image {
	i.Transformations = i.Transformer.AddPrefix(i.Transformations)
	return i
}

func (i Image) AddPageOrFileLayer() Image {
	i.Transformations = i.Transformer.AddPageOrFileLayer(i.Transformations)
	return i
}

func (i Image) AddQuality() Image {
	i.Transformations = i.Transformer.AddQuality(i.Transformations)
	return i
}

func (i Image) AddRoundCorners() Image {
	i.Transformations = i.Transformer.AddRoundCorners(i.Transformations)
	return i
}

func (i Image) AddNamedTransformation() Image {
	i.Transformations = i.Transformer.AddNamedTransformation(i.Transformations)
	return i
}

func (i Image) AddUnderlay() Image {
	i.Transformations = i.Transformer.AddUnderlay(i.Transformations)
	return i
}

func (i Image) AddWidth() Image {
	i.Transformations = i.Transformer.AddWidth(i.Transformations)
	return i
}

func (i Image) AddXY() Image {
	i.Transformations = i.Transformer.AddXY(i.Transformations)
	return i
}

func (i Image) AddZoom() Image {
	i.Transformations = i.Transformer.AddZoom(i.Transformations)
	return i
}

func (i Image) AddVariable() Image {
	i.Transformations = i.Transformer.AddVariable(i.Transformations)
	return i
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
