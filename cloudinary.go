package cloudinary

import (
	"fmt"

	"github.com/syke99/cloudinary/config"
	"github.com/syke99/cloudinary/image"
	"github.com/syke99/cloudinary/resources"
	"github.com/syke99/cloudinary/transformer"
	"github.com/syke99/cloudinary/video"
)

type Cloudinary struct {
	config config.Config
	cloudinary
}

type cloudinary interface {
	Image() image.Image
	Video() video.Video
}

// NewUnsignedCloudinary creates a new instance for unsigned interactions with the Cloudinary API
func NewUnsignedCloudinary(cloud string) *Cloudinary {
	conf := config.Config{
		Cloud: cloud,
	}
	return &Cloudinary{
		config: conf,
	}
}

// NewSignedCloudinary creates a new instance for signed interactions with the Cloudinary API
func NewSignedCloudinary(cloud string, key string, secret string) *Cloudinary {
	conf := config.Config{
		Cloud:     cloud,
		ApiKey:    key,
		ApiSecret: secret,
	}
	return &Cloudinary{
		config: conf,
	}
}

func (c Cloudinary) Image(name string) image.Image {
	url := fmt.Sprintf("%s/%s/image/", resources.BaseUrl, c.config.Cloud)

	return image.Image{}.NewImage(name, url, transformer.Transformer{}, c.config)
}

func (c Cloudinary) Video(name string) video.Video {
	url := fmt.Sprintf("%s/%s/video/", resources.BaseUrl, c.config.Cloud)

	return video.Video{}.NewVideo(name, url, transformer.Transformer{}, c.config)
}
