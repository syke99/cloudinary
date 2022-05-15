package transformer

type DPR struct {
	PixelRatio float32
	Auto       bool
}

func NewDPR(pixelRatio float32, auto bool) DPR {
	return DPR{
		PixelRatio: pixelRatio,
		Auto:       auto,
	}
}
