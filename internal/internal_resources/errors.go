package internal_resources

import "errors"

// Various Errors to be returned by this Go Package

var InvalidDeliveryType = errors.New("delivery type supplied is not a valid asset type for Cloudinary API")
var NoUploadParamsSupplied = errors.New("no upload parameters supplied")
var InvalidAngleMode = errors.New("angle mode provided not valid")
var InvalidAudioCodec = errors.New("audio codec provided not valid")
var InvalidBackgroundAutoMode = errors.New("backgroundauto mode provided not valid")
var InvalidBackgroundAutoDirection = errors.New("backgroundauto direction provided not valid")
var InvalidColorSpaceMode = errors.New("colorspace mode provided invalid")
var InvalidCustomFunctionType = errors.New("custom function type provided not valid")
