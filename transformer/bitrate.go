package transformer

type BitRate struct {
	BitRateValue int
	Constant     bool
}

func NewBitRate(bitrate int, constant bool) BitRate {
	return BitRate{
		BitRateValue: bitrate,
		Constant:     constant,
	}
}
