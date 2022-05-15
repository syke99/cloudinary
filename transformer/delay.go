package transformer

type Delay struct {
	Time int
}

func NewDelay(time int) Delay {
	return Delay{
		Time: time,
	}
}
