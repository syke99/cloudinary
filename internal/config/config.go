package config

import (
	"github.com/syke99/cloudinary/internal/transformer"
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
	Transformer transformer.Transformer
	Name        string
	ReqUrl      string
	UploadUrl   string
}
