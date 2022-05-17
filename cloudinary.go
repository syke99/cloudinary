package cloudinary

import (
	"fmt"
	"github.com/syke99/cloudinary/assets/image"
	"github.com/syke99/cloudinary/assets/video"
	"github.com/syke99/cloudinary/internal/transformer"
	"github.com/syke99/cloudinary/internal/validator"
	"net/http"

	"github.com/syke99/cloudinary/config"
	"github.com/syke99/cloudinary/resources"
)

type Cloudinary struct {
	client      *http.Client
	config      config.Config
	validator   validator.Validator
	transformer transformer.Transformer
	cloudinary
}

type cloudinary interface {
	Image() image.Image
	Video() video.Video
}

// NewUnsignedCloudinary creates a new instance for unsigned interactions with the Cloudinary API
// Unsigned interactions currently not supported
//func NewUnsignedCloudinary(cloud string) *Cloudinary {
//	conf := config.Config{
//		Cloud: cloud,
//	}
//	return &Cloudinary{
//		config: conf,
//	}
//}

// NewSignedCloudinary creates a new instance for signed interactions with the Cloudinary API
func NewSignedCloudinary(client *http.Client, cloud string, key string, secret string) *Cloudinary {
	conf := config.Config{
		Cloud:     cloud,
		ApiKey:    key,
		ApiSecret: secret,
	}
	validator := validator.Validator{}
	transformer := transformer.Transformer{}
	return &Cloudinary{
		client:      client,
		config:      conf,
		validator:   validator,
		transformer: transformer,
	}
}

func (c Cloudinary) Image(name string) image.Image {
	reqUrl := fmt.Sprintf("%s/%s/image/", resources.BaseUrl, c.config.Cloud)
	uploadUrl := fmt.Sprintf("%s/v1_1/%s/image/", resources.BaseUrl, c.config.Cloud)

	return image.Image{Transformer: c.transformer}.NewImage(c.client, c.config, name, reqUrl, uploadUrl)
}

func (c Cloudinary) Video(name string) video.Video {
	reqUrl := fmt.Sprintf("%s/%s/video/", resources.BaseUrl, c.config.Cloud)
	uploadUrl := fmt.Sprintf("%s/v1_1/%s/image/", resources.BaseUrl, c.config.Cloud)

	return video.Video{Transformer: c.transformer}.NewVideo(c.client, c.config, name, reqUrl, uploadUrl)
}
