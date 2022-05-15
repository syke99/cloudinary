package transformer

type Border struct {
	Width string
	Style string
	Color string
}

func NewBorder(width string, style string, color string) Border {
	return Border{
		Width: width,
		Style: style,
		Color: color,
	}
}
