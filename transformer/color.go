package transformer

type Color struct {
	Color string
}

func NewColor(color string) Color {
	return Color{
		Color: color,
	}
}
