package upload

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
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
	UploadMedia(*http.Client, UploaderParameters, string, string, string) UploaderResponse
}

func (u Uploader) SortUploadParameters(params UploaderParameters) []string {

	var sortedParams []string

	prms := reflect.ValueOf(&params).Elem()

	for i := 0; i < prms.NumField(); i++ {
		// prepend an underscore to each capital
		// letter in each fieldName
		reg := regexp.MustCompile("([A-Z])")
		replacement := "_$1"
		newName := reg.ReplaceAllString(prms.Type().Field(i).Name, replacement)

		// remove the leading underscore so fieldName will
		// begin with a letter
		fieldName := newName[1:]

		// convert all characters in fieldName to
		// lowercase
		fieldName = strings.ToLower(fieldName)
		if fieldName == "file" {
			continue
		}

		if fieldName == "access_control" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("access_control=%s", fieldValues[0])

			param := strings.Join(fieldValues, ",")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "tags" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("tags=%s", fieldValues[0])

			param := strings.Join(fieldValues, ",")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "context" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("content=%s", fieldValues[0])

			param := strings.Join(fieldValues, ",")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "metadata" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("metadata=%s", fieldValues[0])

			param := strings.Join(fieldValues, ",")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "responsive_breakpoints" {
			fieldValues := prms.Field(i).Interface().([]string)

			fieldValues[0] = fmt.Sprintf("responsive_breakpoints=%s", fieldValues[0])

			param := strings.Join(fieldValues, ",")

			sortedParams = append(sortedParams, param)
		}

		if fieldName == "time_stamp_unix" {
			fieldValue := prms.Field(i).Interface().(string)

			fieldValue = fmt.Sprintf("timestamp=%s", fieldValue)

			sortedParams = append(sortedParams, fieldValue)
		}

		fieldValue := prms.Field(i).Interface().(string)

		if fieldValue != "" {
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

func (u Uploader) UploadMedia(client *http.Client, params UploaderParameters, apiKey string, signature string, uploadUrl string) UploaderResponse {
	var resp UploaderResponse

	formData := url.Values{}
	formData.Add("api_key", apiKey)
	formData.Add("signature", signature)

	prms := reflect.ValueOf(&params).Elem()

	for i := 0; i < prms.NumField(); i++ {
		// prepend an underscore to each capital
		// letter in each fieldName
		reg := regexp.MustCompile("([A-Z])")
		replacement := "_$1"
		newName := reg.ReplaceAllString(prms.Type().Field(i).Name, replacement)

		// remove the leading underscore so fieldName will
		// begin with a letter
		fieldName := newName[1:]

		// convert all characters in fieldName to
		// lowercase
		fieldName = strings.ToLower(fieldName)

		if fieldName == "file" {
			if prms.Type().Field(i).Type.String() == "[]byte" {
				encodedFile := base64.StdEncoding.EncodeToString(prms.Field(i).Interface().([]byte))
				formData.Add("file", encodedFile)
				// so that we don't add a second "file" key/value
				// pair, continue through to the next iteration
				continue
			}

			fieldValue := prms.Field(i).Interface().(string)
			formData.Add("file", fieldValue)
		}

		if fieldName == "access_control" ||
			fieldName == "tags" ||
			fieldName == "context" ||
			fieldName == "metadata" ||
			fieldName == "responsive_breakpoints" {
			value := strings.Join(prms.Field(i).Interface().([]string), ",")

			formData.Add(fieldName, value)
		}

		if fieldName == "time_stamp_unix" {
			fieldValue := prms.Field(i).Interface().(string)
			name := "timestamp="

			fieldValue = fmt.Sprintf("timestamp=%s", fieldValue)

			formData.Add(name, fieldValue)
		}
	}

	req, _ := http.NewRequest(http.MethodPost, uploadUrl, strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// will need to add error handling here
	res, _ := client.Do(req)

	// unmarshall res into resp

	json.NewDecoder(res.Body).Decode(&resp)

	return resp
}
