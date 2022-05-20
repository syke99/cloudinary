package upload

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/syke99/cloudinary/pkg/api/models/api_handlers/uploader"
)

type Uploader struct {
	params uploader.UploaderParameters
	upload
}

type upload interface {
	SortUploadParameters(uploader.UploaderParameters) []string
	GenerateSignature([]string, string) string
	UploadMedia(*http.Client, uploader.UploaderParameters, string, string, string) uploader.UploaderResponse
}

func (u Uploader) SortUploadParameters(params uploader.UploaderParameters) []string {

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

func (u Uploader) UploadMedia(client *http.Client, params uploader.UploaderParameters, apiKey string, signature string, uploadUrl string) uploader.UploaderResponse {
	var resp uploader.UploaderResponse

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
	resBytes, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(resBytes, &resp)

	return resp
}
