package transformer

type AspectRatio struct {
	Width  float32
	Height float32
}

func NewAspectRatio(w float32, h float32) AspectRatio {
	return AspectRatio{
		Width:  w,
		Height: h,
	}
}
