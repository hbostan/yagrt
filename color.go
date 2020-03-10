package yagrt

type Color struct {
	R, G, B float64
}

func (c Color) Add(o Color) Color {
	return Color{c.R + o.R, c.G + o.G, c.B + o.B}
}

func (c Color) Sub(o Color) Color {
	return Color{c.R - o.R, c.G - o.G, c.B - o.B}
}

func (c Color) Mul(m float64) Color {
	return Color{c.R * m, c.G * m, c.B * m}
}

func (c Color) Div(d float64) Color {
	return Color{c.R / d, c.G / d, c.B / d}
}

func (c Color) MulColor(o Color) Color {
	return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}
