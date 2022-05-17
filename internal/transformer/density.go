package transformer

type Density struct {
	DotsPerInch int
}

func NewDensity(dpi int) Density {
	return Density{
		DotsPerInch: dpi,
	}
}
