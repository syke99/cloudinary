package resources

import "errors"

// Various Errors to be returned by this Go Package

var InvalidDeliveryType = errors.New("delivery type supplied is not a valid asset type for Cloudinary API")
var NoUploadParamsSupplied = errors.New("no upload parameters supplied")
