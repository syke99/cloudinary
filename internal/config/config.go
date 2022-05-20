package config

import (
	"net/http"

	"github.com/syke99/cloudinary/internal/transformer"
	"github.com/syke99/cloudinary/internal/validator"
	"github.com/syke99/cloudinary/pkg/api/actions/upload"
)

type CloudinaryConfig struct {
	Cloud     string
	ApiKey    string
	ApiSecret string
}

type MediaConfig struct {
	Client      *http.Client
	Config      CloudinaryConfig
	Validator   validator.Validator
	Transformer transformer.Transformer
	Uploader    upload.Uploader
	Name        string
	ReqUrl      string
	UploadUrl   string
}
