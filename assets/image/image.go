package image

import (
	"fmt"
	"github.com/syke99/cloudinary/api/request"
	"github.com/syke99/cloudinary/api/upload"
	"github.com/syke99/cloudinary/internal/config"
	"github.com/syke99/cloudinary/internal/internal_resources"
	"github.com/syke99/cloudinary/internal/transformer"
	"github.com/syke99/cloudinary/internal/validator"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Image struct {
	client          *http.Client
	validator       validator.Validator
	config          config.CloudinaryConfig
	transformer     transformer.Transformer
	transformations []string
	Name            string
	Ext             string
	ReqUrl          string
	UploadUrl       string
	image
}

type image interface {
	ConfigureImage(config.MediaConfig) Image
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
	RequestImage(string) ([]byte, error)
	UploadImage(upload.UploaderParameters) (interface{}, error)
}

func (i Image) ConfigureImage(config config.MediaConfig) Image {
	i.client = config.Client
	i.config = config.Config
	i.transformer = config.Transformer
	i.transformations = []string{}
	i.Name = config.Name
	i.Ext = ""
	i.ReqUrl = config.ReqUrl
	i.UploadUrl = config.UploadUrl
	i.validator = validator.Validator{}

	return i
}

func (i Image) AddExtension(ext string) Image {
	i.transformations = i.transformer.AddExtension(i.transformations, ext)
	return i
}

func (i Image) AddAngle(angle transformer.Angle) Image {
	i.transformations = i.transformer.AddAngle(i.transformations, angle)
	return i
}

func (i Image) AddAspectRatio(ar transformer.AspectRatio) Image {
	i.transformations = i.transformer.AddAspectRatio(i.transformations, ar)
	return i
}

func (i Image) AddBackground(background transformer.Background) Image {
	i.transformations = i.transformer.AddBackground(i.transformations, background)
	return i
}

func (i Image) AddBorder(border transformer.Border) Image {
	i.transformations = i.transformer.AddBorder(i.transformations, border)
	return i
}

func (i Image) AddCropOrResize(resize transformer.CropResize) Image {
	i.transformations = i.transformer.AddCropOrResize(i.transformations, resize)
	return i
}

func (i Image) AddColor(color transformer.Color) Image {
	i.transformations = i.transformer.AddColor(i.transformations, color)
	return i
}

func (i Image) AddColorSpace(space transformer.ColorSpace) Image {
	i.transformations = i.transformer.AddColorSpace(i.transformations, space)
	return i
}

func (i Image) AddDefaultImage(defaultImage transformer.DefaultImage) Image {
	i.transformations = i.transformer.AddDefaultImage(i.transformations, defaultImage)
	return i
}

func (i Image) AddDelay(delay transformer.Delay) Image {
	i.transformations = i.transformer.AddDelay(i.transformations, delay)
	return i
}

func (i Image) AddDensity(density transformer.Density) Image {
	i.transformations = i.transformer.AddDensity(i.transformations, density)
	return i
}

func (i Image) AddDPR(dpr transformer.DPR) Image {
	i.transformations = i.transformer.AddDPR(i.transformations, dpr)
	return i
}

func (i Image) AddEffect(effect transformer.Effect) Image {
	i.transformations = i.transformer.AddEffect(i.transformations, effect)
	return i
}

func (i Image) AddFormat() Image {
	i.transformations = i.transformer.AddFormat(i.transformations)
	return i
}

func (i Image) AddFlag() Image {
	i.transformations = i.transformer.AddFlag(i.transformations)
	return i
}

func (i Image) AddCustomFunction() Image {
	i.transformations = i.transformer.AddCustomFunction(i.transformations)
	return i
}

func (i Image) AddGravity() Image {
	i.transformations = i.transformer.AddGravity(i.transformations)
	return i
}

func (i Image) AddHeight() Image {
	i.transformations = i.transformer.AddHeight(i.transformations)
	return i
}

func (i Image) AddIf() Image {
	i.transformations = i.transformer.AddIf(i.transformations)
	return i
}

func (i Image) AddLayer() Image {
	i.transformations = i.transformer.AddLayer(i.transformations)
	return i
}

func (i Image) AddOpacity() Image {
	i.transformations = i.transformer.AddOpacity(i.transformations)
	return i
}

func (i Image) AddPrefix() Image {
	i.transformations = i.transformer.AddPrefix(i.transformations)
	return i
}

func (i Image) AddPageOrFileLayer() Image {
	i.transformations = i.transformer.AddPageOrFileLayer(i.transformations)
	return i
}

func (i Image) AddQuality() Image {
	i.transformations = i.transformer.AddQuality(i.transformations)
	return i
}

func (i Image) AddRoundCorners() Image {
	i.transformations = i.transformer.AddRoundCorners(i.transformations)
	return i
}

func (i Image) AddNamedTransformation() Image {
	i.transformations = i.transformer.AddNamedTransformation(i.transformations)
	return i
}

func (i Image) AddUnderlay() Image {
	i.transformations = i.transformer.AddUnderlay(i.transformations)
	return i
}

func (i Image) AddWidth() Image {
	i.transformations = i.transformer.AddWidth(i.transformations)
	return i
}

func (i Image) AddXY() Image {
	i.transformations = i.transformer.AddXY(i.transformations)
	return i
}

func (i Image) AddZoom() Image {
	i.transformations = i.transformer.AddZoom(i.transformations)
	return i
}

func (i Image) AddVariable() Image {
	i.transformations = i.transformer.AddVariable(i.transformations)
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

func (i Image) UploadImage(params upload.UploaderParameters) (interface{}, error) {
	if reflect.ValueOf(params).Len() == 0 {
		return nil, internal_resources.NoUploadParamsSupplied
	}

	for _, transformation := range i.transformations {
		params.Transformation += transformation
	}
	params.TimeStampUnix = strconv.FormatInt(time.Now().Unix(), 10)
	u := upload.Uploader{}

	sortedParams := u.SortUploadParameters(params)

	signature := u.GenerateSignature(sortedParams, i.config.ApiKey)

	return u.UploadMedia(i.client, params, i.config.ApiKey, signature, i.UploadUrl), nil
}
