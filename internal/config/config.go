package config

import (
	"github.com/syke99/cloudinary/internal/api/upload"
	"github.com/syke99/cloudinary/internal/transformer"
	"github.com/syke99/cloudinary/internal/validator"
	"net/http"
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
