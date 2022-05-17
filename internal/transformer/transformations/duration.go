package transformations

type Duration struct {
	Time float32
}

func NewDuration(time float32) Duration {
	return Duration{
		Time: time,
	}
}
