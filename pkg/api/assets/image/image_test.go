package image

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syke99/cloudinary/internal/config"
	"github.com/syke99/cloudinary/internal/internal_resources/test"
	"github.com/syke99/cloudinary/internal/transformer"
	"github.com/syke99/cloudinary/internal/validator"
	"github.com/syke99/cloudinary/pkg/api/actions/upload"
)

func TestNewImage(t *testing.T) {
	mockedMediaConfig := config.MediaConfig{
		Client:      &http.Client{},
		Config:      config.CloudinaryConfig{},
		Validator:   validator.Validator{},
		Transformer: transformer.Transformer{},
		Uploader:    upload.Uploader{},
		Name:        "testing",
		ReqUrl:      "https://testing.request.url.com",
		UploadUrl:   "https://testing.upload.url.com",
	}

	testImage := NewImage(mockedMediaConfig)

	assert.NotNil(t, testImage)
	assert.Equal(t, test.TestMediaName, testImage.Name)
	assert.Equal(t, test.TestMediaReqUrl, testImage.ReqUrl)
	assert.Equal(t, test.TestMediaUploadUrl, testImage.UploadUrl)
}
