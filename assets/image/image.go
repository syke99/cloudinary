package image

import (
	"fmt"
	"github.com/syke99/cloudinary/api/request"
	"github.com/syke99/cloudinary/api/upload"
	"github.com/syke99/cloudinary/config"
	"github.com/syke99/cloudinary/internal/internal_resources"
	transformer2 "github.com/syke99/cloudinary/internal/transformer"
	"github.com/syke99/cloudinary/internal/validator"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Image struct {
	client          *http.Client
	validator       validator.Validator
	config          config.Config
	Transformer     transformer2.Transformer
	transformations []string
	Name            string
	Ext             string
	ReqUrl          string
	UploadUrl       string
	image
}

type image interface {
	NewImage(*http.Client, string, string, string, transformer2.Transformer, config.Config) Image
	AddExtension(string) Image
	AddAngle(transformer2.Angle) Image
	AddAspectRatio(transformer2.AspectRatio) Image
	AddBackground(transformer2.Background) Image
	AddBorder(transformer2.Border) Image
	AddCropOrResize(transformer2.CropResize) Image
	AddColor(transformer2.Color) Image
	AddColorSpace(transformer2.ColorSpace) Image
	AddDefaultImage(transformer2.DefaultImage) Image
	AddDelay(transformer2.Delay) Image
	AddDensity(transformer2.Density) Image
	AddDPR(transformer2.DPR) Image
	AddEffect(transformer2.Effect) Image
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
	UploadImage(upload.UploaderParameters, string) (interface{}, error)
}

func (i Image) NewImage(client *http.Client, name string, reqUrl string, uploadUrl string, transformer transformer2.Transformer, config config.Config) Image {
	i.client = client
	i.config = config
	i.Transformer = transformer
	i.transformations = []string{}
	i.Name = name
	i.Ext = ""
	i.ReqUrl = reqUrl
	i.UploadUrl = uploadUrl
	i.validator = validator.Validator{}

	return i
}

func (i Image) AddExtension(ext string) Image {
	i.transformations = i.Transformer.AddExtension(i.transformations, ext)
	return i
}

func (i Image) AddAngle(angle transformer2.Angle) Image {
	i.transformations = i.Transformer.AddAngle(i.transformations, angle)
	return i
}

func (i Image) AddAspectRatio(ar transformer2.AspectRatio) Image {
	i.transformations = i.Transformer.AddAspectRatio(i.transformations, ar)
	return i
}

func (i Image) AddBackground(background transformer2.Background) Image {
	i.transformations = i.Transformer.AddBackground(i.transformations, background)
	return i
}

func (i Image) AddBorder(border transformer2.Border) Image {
	i.transformations = i.Transformer.AddBorder(i.transformations, border)
	return i
}

func (i Image) AddCropOrResize(resize transformer2.CropResize) Image {
	i.transformations = i.Transformer.AddCropOrResize(i.transformations, resize)
	return i
}

func (i Image) AddColor(color transformer2.Color) Image {
	i.transformations = i.Transformer.AddColor(i.transformations, color)
	return i
}

func (i Image) AddColorSpace(space transformer2.ColorSpace) Image {
	i.transformations = i.Transformer.AddColorSpace(i.transformations, space)
	return i
}

func (i Image) AddDefaultImage(defaultImage transformer2.DefaultImage) Image {
	i.transformations = i.Transformer.AddDefaultImage(i.transformations, defaultImage)
	return i
}

func (i Image) AddDelay(delay transformer2.Delay) Image {
	i.transformations = i.Transformer.AddDelay(i.transformations, delay)
	return i
}

func (i Image) AddDensity(density transformer2.Density) Image {
	i.transformations = i.Transformer.AddDensity(i.transformations, density)
	return i
}

func (i Image) AddDPR(dpr transformer2.DPR) Image {
	i.transformations = i.Transformer.AddDPR(i.transformations, dpr)
	return i
}

func (i Image) AddEffect(effect transformer2.Effect) Image {
	i.transformations = i.Transformer.AddEffect(i.transformations, effect)
	return i
}

func (i Image) AddFormat() Image {
	i.transformations = i.Transformer.AddFormat(i.transformations)
	return i
}

func (i Image) AddFlag() Image {
	i.transformations = i.Transformer.AddFlag(i.transformations)
	return i
}

func (i Image) AddCustomFunction() Image {
	i.transformations = i.Transformer.AddCustomFunction(i.transformations)
	return i
}

func (i Image) AddGravity() Image {
	i.transformations = i.Transformer.AddGravity(i.transformations)
	return i
}

func (i Image) AddHeight() Image {
	i.transformations = i.Transformer.AddHeight(i.transformations)
	return i
}

func (i Image) AddIf() Image {
	i.transformations = i.Transformer.AddIf(i.transformations)
	return i
}

func (i Image) AddLayer() Image {
	i.transformations = i.Transformer.AddLayer(i.transformations)
	return i
}

func (i Image) AddOpacity() Image {
	i.transformations = i.Transformer.AddOpacity(i.transformations)
	return i
}

func (i Image) AddPrefix() Image {
	i.transformations = i.Transformer.AddPrefix(i.transformations)
	return i
}

func (i Image) AddPageOrFileLayer() Image {
	i.transformations = i.Transformer.AddPageOrFileLayer(i.transformations)
	return i
}

func (i Image) AddQuality() Image {
	i.transformations = i.Transformer.AddQuality(i.transformations)
	return i
}

func (i Image) AddRoundCorners() Image {
	i.transformations = i.Transformer.AddRoundCorners(i.transformations)
	return i
}

func (i Image) AddNamedTransformation() Image {
	i.transformations = i.Transformer.AddNamedTransformation(i.transformations)
	return i
}

func (i Image) AddUnderlay() Image {
	i.transformations = i.Transformer.AddUnderlay(i.transformations)
	return i
}

func (i Image) AddWidth() Image {
	i.transformations = i.Transformer.AddWidth(i.transformations)
	return i
}

func (i Image) AddXY() Image {
	i.transformations = i.Transformer.AddXY(i.transformations)
	return i
}

func (i Image) AddZoom() Image {
	i.transformations = i.Transformer.AddZoom(i.transformations)
	return i
}

func (i Image) AddVariable() Image {
	i.transformations = i.Transformer.AddVariable(i.transformations)
	return i
}

func (i Image) RequestImage(delivery string) ([]byte, error) {
	err := i.validator.ValidateDeliveryType(delivery)
	if err != nil {
		return []byte{}, err
	}

	r := request.Request{}
	reqUrl := fmt.Sprintf("%s/%s/%s/%s", i.ReqUrl, delivery, i.Name, i.Ext)
	return r.RequestMedia(i.client, reqUrl), nil
}

func (i Image) UploadImage(params upload.UploaderParameters, apiKey string) (interface{}, error) {
	if reflect.ValueOf(params).Len() == 0 {
		return nil, internal_resources.NoUploadParamsSupplied
	}

	for _, transformation := range i.transformations {
		params.Transformation += transformation
	}
	params.TimeStampUnix = strconv.FormatInt(time.Now().Unix(), 10)
	u := upload.Uploader{}

	sortedParams := u.SortUploadParameters(params)

	signature := u.GenerateSignature(sortedParams, apiKey)

	return u.UploadMedia(i.client, params, apiKey, signature, i.UploadUrl), nil
}
