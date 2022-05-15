package transformer

type AudioFrequency struct {
	FrequencyValue string
	Iaf            bool
}

func NewAudioFrequency(freq string, iaf bool) AudioFrequency {
	return AudioFrequency{
		FrequencyValue: freq,
		Iaf:            iaf,
	}
}
