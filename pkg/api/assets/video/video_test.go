package video

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syke99/cloudinary/internal/config"
	"github.com/syke99/cloudinary/internal/internal_resources/test"
	"github.com/syke99/cloudinary/internal/transformer"
	"github.com/syke99/cloudinary/internal/validator"
	"github.com/syke99/cloudinary/pkg/api/upload"
)

func TestNewVideo(t *testing.T) {
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

	testVideo := NewVideo(mockedMediaConfig)

	assert.Equal(t, test.TestMediaName, testVideo.Name)
	assert.Equal(t, test.TestMediaReqUrl, testVideo.ReqUrl)
	assert.Equal(t, test.TestMediaUploadUrl, testVideo.UploadUrl)
}
