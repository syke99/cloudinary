package video

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

type Video struct {
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
	video
}

type video interface {
	ConfigureVideo(config.MediaConfig) *Video
	AddExtension(string) *Video
	AddAngle(transformations.Angle) *Video
	AddAudioCodec(transformations.AudioCodec) *Video
	AddAudioFrequency(transformations.AudioFrequency) *Video
	AddBackground(transformations.Background) *Video
	AddBorder(transformations.Border) *Video
	AddBitrate(transformations.BitRate) *Video
	AddCropOrResize(transformations.CropResize) *Video
	AddColor(transformations.Color) *Video
	AddColorSpace(transformations.ColorSpace) *Video
	AddDelay(transformations.Delay) *Video
	AddDPR(transformations.DPR) *Video
	AddDuration(transformations.Duration) *Video
	AddEffect(transformations.Effect) *Video
	AddEndOffset() *Video
	AddFormat() *Video
	AddFlag() *Video
	AddFPS() *Video
	AddGravity() *Video
	AddHeight() *Video
	AddIf() *Video
	AddKeyframeInterval() *Video
	AddLayer() *Video
	AddQuality() *Video
	AddRoundCorners() *Video
	AddStartOffset() *Video
	AddStreamingProfile() *Video
	AddNamedTransformation() *Video
	AddUnderlay() *Video
	AddVideoCodec() *Video
	AddVideoSampling() *Video
	AddWidth() *Video
	AddXY() *Video
	AddVariable() *Video
	RequestVideo(string) ([]byte, error)
	UploadVideo(upload.UploaderParameters) (upload.UploaderResponse, error)
}

func NewVideo(config config.MediaConfig) *Video {
	v := Video{
		client:          config.Client,
		config:          config.Config,
		transformer:     config.Transformer,
		transformations: []string{},
		Name:            config.Name,
		Ext:             "",
		ReqUrl:          config.ReqUrl,
		UploadUrl:       config.UploadUrl,
		validator:       config.Validator,
		uploader:        config.Uploader,
	}

	return &v
}

func (v *Video) AddExtension(ext string) *Video {
	v.transformations = v.transformer.AddExtension(v.transformations, ext)
	return v
}

func (v *Video) AddAngle(angle transformations.Angle) *Video {
	v.transformations = v.transformer.AddAngle(v.transformations, angle)
	return v
}

func (v *Video) AddAudioCodec(ac transformations.AudioCodec) *Video {
	v.transformations = v.transformer.AddAudioCodec(v.transformations, ac)
	return v
}

func (v *Video) AddAudioFrequency(af transformations.AudioFrequency) *Video {
	v.transformations = v.transformer.AddAudioFrequency(v.transformations, af)
	return v
}

func (v *Video) AddBackground(background transformations.Background) *Video {
	v.transformations = v.transformer.AddBackground(v.transformations, background)
	return v
}

func (v *Video) AddBorder(border transformations.Border) *Video {
	v.transformations = v.transformer.AddBorder(v.transformations, border)
	return v
}

func (v *Video) AddBitrate(br transformations.BitRate) *Video {
	v.transformations = v.transformer.AddBitrate(v.transformations, br)
	return v
}

func (v *Video) AddCropOrResize(cr transformations.CropResize) *Video {
	v.transformations = v.transformer.AddCropOrResize(v.transformations, cr)
	return v
}

func (v *Video) AddColor(color transformations.Color) *Video {
	v.transformations = v.transformer.AddColor(v.transformations, color)
	return v
}

func (v *Video) AddColorSpace(colorSpace transformations.ColorSpace) *Video {
	v.transformations = v.transformer.AddColorSpace(v.transformations, colorSpace)
	return v
}

func (v *Video) AddDelay(delay transformations.Delay) *Video {
	v.transformations = v.transformer.AddDelay(v.transformations, delay)
	return v
}

func (v *Video) AddDPR(dpr transformations.DPR) *Video {
	v.transformations = v.transformer.AddDPR(v.transformations, dpr)
	return v
}

func (v *Video) AddDuration(duration transformations.Duration) *Video {
	v.transformations = v.transformer.AddDuration(v.transformations, duration)
	return v
}

func (v *Video) AddEffect(effect transformations.Effect) *Video {
	v.transformations = v.transformer.AddEffect(v.transformations, effect)
	return v
}

func (v *Video) AddEndOffset() *Video {
	v.transformations = v.transformer.AddEndOffset(v.transformations)
	return v
}

func (v *Video) AddFormat() *Video {
	v.transformations = v.transformer.AddFormat(v.transformations)
	return v
}

func (v *Video) AddFlag() *Video {
	v.transformations = v.transformer.AddFlag(v.transformations)
	return v
}

func (v *Video) AddFPS() *Video {
	v.transformations = v.transformer.AddFPS(v.transformations)
	return v
}

func (v *Video) AddGravity() *Video {
	v.transformations = v.transformer.AddGravity(v.transformations)
	return v
}

func (v *Video) AddHeight() *Video {
	v.transformations = v.transformer.AddHeight(v.transformations)
	return v
}

func (v *Video) AddIf() *Video {
	v.transformations = v.transformer.AddIf(v.transformations)
	return v
}

func (v *Video) AddKeyframeInterval() *Video {
	v.transformations = v.transformer.AddKeyframeInterval(v.transformations)
	return v
}

func (v *Video) AddLayer() *Video {
	v.transformations = v.transformer.AddLayer(v.transformations)
	return v
}

func (v *Video) AddQuality() *Video {
	v.transformations = v.transformer.AddQuality(v.transformations)
	return v
}

func (v *Video) AddRoundCorners() *Video {
	v.transformations = v.transformer.AddRoundCorners(v.transformations)
	return v
}

func (v *Video) AddStartOffset() *Video {
	v.transformations = v.transformer.AddStartOffset(v.transformations)
	return v
}

func (v *Video) AddStreamingProfile() *Video {
	v.transformations = v.transformer.AddStreamingProfile(v.transformations)
	return v
}

func (v *Video) AddNamedTransformation() *Video {
	v.transformations = v.transformer.AddNamedTransformation(v.transformations)
	return v
}

func (v *Video) AddUnderlay() *Video {
	v.transformations = v.transformer.AddUnderlay(v.transformations)
	return v
}

func (v *Video) AddVideoCodec() *Video {
	v.transformations = v.transformer.AddVideoCodec(v.transformations)
	return v
}

func (v *Video) AddVideoSampling() *Video {
	v.transformations = v.transformer.AddVideoSampling(v.transformations)
	return v
}

func (v *Video) AddWidth() *Video {
	v.transformations = v.transformer.AddWidth(v.transformations)
	return v
}

func (v *Video) AddXY() *Video {
	v.transformations = v.transformer.AddXY(v.transformations)
	return v
}

func (v *Video) AddVariable() *Video {
	v.transformations = v.transformer.AddVariable(v.transformations)
	return v
}

func (v *Video) RequestVideo(delivery string) ([]byte, error) {
	err := v.validator.ValidateDeliveryType(delivery)
	if err != nil {
		return []byte{}, err
	}

	r := request.Request{}
	reqUrl := fmt.Sprintf("%s/%s/%s/%s", v.ReqUrl, delivery, v.Name, v.Ext)
	return r.RequestMedia(v.client, reqUrl), nil
}

func (v *Video) UploadVideo(params upload.UploaderParameters) (upload.UploaderResponse, error) {
	if reflect.ValueOf(params).Len() == 0 {
		return upload.UploaderResponse{}, internal_resources.NoUploadParamsSupplied
	}

	for _, transformation := range v.transformations {
		params.Transformation += transformation
	}
	params.TimeStampUnix = strconv.FormatInt(time.Now().Unix(), 10)

	sortedParams := v.uploader.SortUploadParameters(params)

	signature := v.uploader.GenerateSignature(sortedParams, v.config.ApiKey)

	return v.uploader.UploadMedia(v.client, params, v.config.ApiKey, signature, v.UploadUrl), nil
}
