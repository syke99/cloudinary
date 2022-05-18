package image

import (
	"fmt"
	"github.com/syke99/cloudinary/internal/api/request"
	"github.com/syke99/cloudinary/internal/api/upload"
	"github.com/syke99/cloudinary/internal/config"
	"github.com/syke99/cloudinary/internal/internal_resources"
	"github.com/syke99/cloudinary/internal/transformer"
	"github.com/syke99/cloudinary/internal/transformer/transformations"
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
	uploader        upload.Uploader
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
	i.validator = config.Validator
	i.uploader = config.Uploader

	return i
}

func (i Image) AddExtension(ext string) Image {
	i.transformations = i.transformer.AddExtension(i.transformations, ext)
	return i
}

func (i Image) AddAngle(angle transformations.Angle) Image {
	i.transformations = i.transformer.AddAngle(i.transformations, angle)
	return i
}

func (i Image) AddAspectRatio(ar transformations.AspectRatio) Image {
	i.transformations = i.transformer.AddAspectRatio(i.transformations, ar)
	return i
}

func (i Image) AddBackground(background transformations.Background) Image {
	i.transformations = i.transformer.AddBackground(i.transformations, background)
	return i
}

func (i Image) AddBorder(border transformations.Border) Image {
	i.transformations = i.transformer.AddBorder(i.transformations, border)
	return i
}

func (i Image) AddCropOrResize(resize transformations.CropResize) Image {
	i.transformations = i.transformer.AddCropOrResize(i.transformations, resize)
	return i
}

func (i Image) AddColor(color transformations.Color) Image {
	i.transformations = i.transformer.AddColor(i.transformations, color)
	return i
}

func (i Image) AddColorSpace(space transformations.ColorSpace) Image {
	i.transformations = i.transformer.AddColorSpace(i.transformations, space)
	return i
}

func (i Image) AddDefaultImage(defaultImage transformations.DefaultImage) Image {
	i.transformations = i.transformer.AddDefaultImage(i.transformations, defaultImage)
	return i
}

func (i Image) AddDelay(delay transformations.Delay) Image {
	i.transformations = i.transformer.AddDelay(i.transformations, delay)
	return i
}

func (i Image) AddDensity(density transformations.Density) Image {
	i.transformations = i.transformer.AddDensity(i.transformations, density)
	return i
}

func (i Image) AddDPR(dpr transformations.DPR) Image {
	i.transformations = i.transformer.AddDPR(i.transformations, dpr)
	return i
}

func (i Image) AddEffect(effect transformations.Effect) Image {
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

	sortedParams := i.uploader.SortUploadParameters(params)

	signature := i.uploader.GenerateSignature(sortedParams, i.config.ApiKey)

	return i.uploader.UploadMedia(i.client, params, i.config.ApiKey, signature, i.UploadUrl), nil
}
