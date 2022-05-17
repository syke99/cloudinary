package upload

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Uploader struct {
	params UploaderParameters
	uploader
}

type UploaderParameters struct {
	// RequiredParameters
	File          interface{}
	TimeStampUnix string
	Signature     string
	// OptionalParameters
	PublicId                string
	Folder                  string
	UseFilename             bool
	UniqueFilename          bool
	FilenameOverride        string
	ResourceType            string
	Type                    string
	AccessControl           []string
	AccessMode              string
	DiscardOriginalFilename bool
	Overwrite               bool
	// -Resource Data
	Tags                  []string
	Context               []string
	Metadata              []string
	Colors                bool
	Faces                 bool
	QualityAnalysis       bool
	AccessibilityAnalysis bool
	CinemagraphAnalysis   bool
	ImageMetadata         bool
	Phash                 bool
	ResponsiveBreakpoints []string
	AutoTagging           float64
	Categorization        string
	Detection             string
	OCR                   string
	Exif                  bool
	// -Manipulations
	Eager                string
	EagerAsync           bool
	EagerNotificationUrl string
	Transformation       string
	Format               string
	CustomCoordinates    string
	FaceCoordinates      string
	BackgroundRemoval    string
	RawConvert           string
	// -Additional Options
	AllowedFormats    string
	Async             bool
	Backup            bool
	Callback          string
	Eval              string
	Headers           string
	Invalidate        bool
	Moderation        string
	NotificationUrl   string
	Proxy             string
	ReturnDeleteToken bool
}

type UploaderResponse struct {
	AssetId          string          `json:"asset_id"`
	PublicId         string          `json:"public_id"`
	Version          int             `json:"version"`
	VersionId        string          `json:"version_id"`
	Signature        string          `json:"signature"`
	Width            string          `json:"width"`
	Height           string          `json:"height"`
	Format           string          `json:"format"`
	ResourceType     string          `json:"resource_type"`
	CreatedAt        string          `json:"created_at"`
	Tags             []responseTags  `json:"tags"`
	Pages            int             `json:"pages"`
	Bytes            int             `json:"bytes"`
	Type             string          `json:"type"`
	ETag             string          `json:"etag"`
	Placeholder      bool            `json:"placeholder"`
	Url              string          `json:"url"`
	SecureUrl        string          `json:"secure_url"`
	AccessMode       string          `json:"access_mode"`
	OriginalFilename string          `json:"original_filename"`
	Eager            []responseEager `json:"eager"`
}

type responseTags struct {
	Tag string `json:"tag"`
}

type responseEager struct {
	Transformation string `json:"transformation"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Url            string `json:"url"`
	SecureUrl      string `json:"secure_url"`
}

type uploader interface {
	SortUploadParameters(UploaderParameters) []string
	GenerateSignature([]string, string) string
	UploadMedia(UploaderParameters) UploaderResponse
}

func (u Uploader) SortUploadParameters(params UploaderParameters) []string {

	var sortedParams []string

	prms := reflect.ValueOf(&params).Elem()

	for i := 0; i < prms.NumField(); i++ {
		fieldName := strings.ToLower(prms.Type().Field(i).Name)
		if fieldName == "file" ||
			fieldName == "timestampunix" {
			continue
		}

		if fieldName == "accesscontrol" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("access_control=%s", fieldValues[0])

			param := strings.Join(fieldValues, ", access_control=")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "tags" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("tags=%s", fieldValues[0])

			param := strings.Join(fieldValues, ", tags=")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "context" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("context=%s", fieldValues[0])

			param := strings.Join(fieldValues, ", context=")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "metadata" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("metadata=%s", fieldValues[0])

			param := strings.Join(fieldValues, ", metadata=")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "responsivebreakpoints" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("responsive_breakpoints=%s", fieldValues[0])

			param := strings.Join(fieldValues, ", responsive_breakpoints=")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "timestampunix" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("timestamp=%s", fieldValues[0])

			param := strings.Join(fieldValues, ", timestamp=")

			sortedParams = append(sortedParams, param)
		}

		fieldValue := prms.Field(i).Interface().(string)

		if fieldValue != "" {
			fieldValue = fmt.Sprintf("%s=%s", fieldName, fieldValue)
			sortedParams = append(sortedParams, fieldValue)
		}
	}

	sort.Strings(sortedParams)

	return sortedParams
}

func (u Uploader) GenerateSignature(sortedParams []string, apiKey string) string {
	var encryptedSignature []byte

	stringifiedParams := strings.Join(sortedParams, "&")

	paramsWithKey := fmt.Sprintf("%s%s", stringifiedParams, apiKey)

	hasher := sha1.New()

	hasher.Write([]byte(paramsWithKey))

	hasher.Sum(encryptedSignature)

	encodedSignature := hex.EncodeToString(encryptedSignature)

	return encodedSignature
}

func (u Uploader) UploadMedia(params UploaderParameters) UploaderResponse {
	var resp UploaderResponse

	return resp
}
