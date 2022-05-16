package image

import (
	"fmt"

	"github.com/syke99/cloudinary/config"
	"github.com/syke99/cloudinary/request"
	"github.com/syke99/cloudinary/transformer"
	"github.com/syke99/cloudinary/validator"
)

type Image struct {
	Transformer transformer.Transformer
	transformer []string
	Name        string
	Ext         string
	Url         string
	image
}

type image interface {
	AddExtension(string) Image
	AddAngle(transformer.Angle) Image
	AddAspectRatio(transformer.AspectRatio) Image
	AddBackground(transformer.Background) Image
	AddBorder(transformer.Border) Image
	AddCropOrResize(transformer.CropResize) Image
	AddColor(transformer.Color) Image
	AddColorSpace(transformer.ColorSpace) Image
	AddDefaultImage(transformer.DefaultImage) Image
	AddDelay(transformer.Delay) Image
	AddDensity(transformer.Density) Image
	AddDPR(transformer.DPR) Image
	AddEffect(transformer.Effect) Image
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
	RequestImage(string, config.Config) ([]byte, error)
}

func (i Image) AddExtension(ext string) Image {
	i.transformer = i.Transformer.AddExtension(i.transformer, ext)
	return i
}

func (i Image) AddAngle(angle transformer.Angle) Image {
	i.transformer = i.Transformer.AddAngle(i.transformer, angle)
	return i
}

func (i Image) AddAspectRatio(ar transformer.AspectRatio) Image {
	i.transformer = i.Transformer.AddAspectRatio(i.transformer, ar)
	return i
}

func (i Image) AddBackground(background transformer.Background) Image {
	i.transformer = i.Transformer.AddBackground(i.transformer, background)
	return i
}

func (i Image) AddBorder(border transformer.Border) Image {
	i.transformer = i.Transformer.AddBorder(i.transformer, border)
	return i
}

func (i Image) AddCropOrResize(resize transformer.CropResize) Image {
	i.transformer = i.Transformer.AddCropOrResize(i.transformer, resize)
	return i
}

func (i Image) AddColor(color transformer.Color) Image {
	i.transformer = i.Transformer.AddColor(i.transformer, color)
	return i
}

func (i Image) AddColorSpace(space transformer.ColorSpace) Image {
	i.transformer = i.Transformer.AddColorSpace(i.transformer, space)
	return i
}

func (i Image) AddDefaultImage(defaultImage transformer.DefaultImage) Image {
	i.transformer = i.Transformer.AddDefaultImage(i.transformer, defaultImage)
	return i
}

func (i Image) AddDelay(delay transformer.Delay) Image {
	i.transformer = i.Transformer.AddDelay(i.transformer, delay)
	return i
}

func (i Image) AddDensity(density transformer.Density) Image {
	i.transformer = i.Transformer.AddDensity(i.transformer, density)
	return i
}

func (i Image) AddDPR(dpr transformer.DPR) Image {
	i.transformer = i.Transformer.AddDPR(i.transformer, dpr)
	return i
}

func (i Image) AddEffect(effect transformer.Effect) Image {
	i.transformer = i.Transformer.AddEffect(i.transformer, effect)
	return i
}

func (i Image) AddFormat() Image {
	i.transformer = i.Transformer.AddFormat(i.transformer)
	return i
}

func (i Image) AddFlag() Image {
	i.transformer = i.Transformer.AddFlag(i.transformer)
	return i
}

func (i Image) AddCustomFunction() Image {
	i.transformer = i.Transformer.AddCustomFunction(i.transformer)
	return i
}

func (i Image) AddGravity() Image {
	i.transformer = i.Transformer.AddGravity(i.transformer)
	return i
}

func (i Image) AddHeight() Image {
	i.transformer = i.Transformer.AddHeight(i.transformer)
	return i
}

func (i Image) AddIf() Image {
	i.transformer = i.Transformer.AddIf(i.transformer)
	return i
}

func (i Image) AddLayer() Image {
	i.transformer = i.Transformer.AddLayer(i.transformer)
	return i
}

func (i Image) AddOpacity() Image {
	i.transformer = i.Transformer.AddOpacity(i.transformer)
	return i
}

func (i Image) AddPrefix() Image {
	i.transformer = i.Transformer.AddPrefix(i.transformer)
	return i
}

func (i Image) AddPageOrFileLayer() Image {
	i.transformer = i.Transformer.AddPageOrFileLayer(i.transformer)
	return i
}

func (i Image) AddQuality() Image {
	i.transformer = i.Transformer.AddQuality(i.transformer)
	return i
}

func (i Image) AddRoundCorners() Image {
	i.transformer = i.Transformer.AddRoundCorners(i.transformer)
	return i
}

func (i Image) AddNamedTransformation() Image {
	i.transformer = i.Transformer.AddNamedTransformation(i.transformer)
	return i
}

func (i Image) AddUnderlay() Image {
	i.transformer = i.Transformer.AddUnderlay(i.transformer)
	return i
}

func (i Image) AddWidth() Image {
	i.transformer = i.Transformer.AddWidth(i.transformer)
	return i
}

func (i Image) AddXY() Image {
	i.transformer = i.Transformer.AddXY(i.transformer)
	return i
}

func (i Image) AddZoom() Image {
	i.transformer = i.Transformer.AddZoom(i.transformer)
	return i
}

func (i Image) AddVariable() Image {
	i.transformer = i.Transformer.AddVariable(i.transformer)
	return i
}

func (i Image) RequestImage(delivery string, config config.Config) ([]byte, error) {
	err := validator.ValidateDeliveryType(delivery)
	if err != nil {
		return []byte{}, err
	}

	r := request.Request{}
	reqUrl := fmt.Sprintf("%s/%s/%s/%s", i.Url, delivery, i.Name, i.Ext)
	return r.RequestMedia(reqUrl), nil
}
