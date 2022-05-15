package cloudinary

import (
	"fmt"

	"github.com/syke99/cloudinary/image"
	"github.com/syke99/cloudinary/resources"
	"github.com/syke99/cloudinary/transformer"
	"github.com/syke99/cloudinary/video"
)

type Cloudinary struct {
	cloud string
	cloudinary
}

type cloudinary interface {
	Image() image.Image
	Video() video.Video
}

func NewCloudinary(cloud string) Cloudinary {
	return Cloudinary{
		cloud: cloud,
	}
}

func (c Cloudinary) Image(name string) image.Image {
	url := fmt.Sprintf("%s/%s/image/", resources.BaseUrl, c.cloud)

	return image.Image{
		Transformer: transformer.Transformer{},
		Name:        name,
		Url:         url,
	}
}

func (c Cloudinary) Video(name string) video.Video {
	url := fmt.Sprintf("%s/%s/video/", resources.BaseUrl, c.cloud)

	return video.Video{
		Transformer: transformer.Transformer{},
		Name:        name,
		Url:         url,
	}
}
