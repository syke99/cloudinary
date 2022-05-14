package validation

import "github.com/syke99/cloudinary/resources"

func ValidateDeliveryType(delivery string) error {
	if delivery != resources.Upload ||
		delivery != resources.Private ||
		delivery != resources.Authenticated ||
		delivery != resources.List ||
		delivery != resources.Fetch ||
		delivery != resources.Facebook ||
		delivery != resources.Twitter ||
		delivery != resources.TwitterName ||
		delivery != resources.Gravatar ||
		delivery != resources.YouTube ||
		delivery != resources.Hulu ||
		delivery != resources.Vimeo ||
		delivery != resources.Animoto ||
		delivery != resources.WorldStarHipHop ||
		delivery != resources.DailyMotion ||
		delivery != resources.Multi ||
		delivery != resources.Delivery {
		return resources.InvalidDeliveryType
	}

	return nil
}
