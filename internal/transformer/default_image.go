package transformer

type DefaultImage struct {
	ImageAsset string
}

func NewDefaultImage(imageAsset string) DefaultImage {
	return DefaultImage{
		ImageAsset: imageAsset,
	}
}
