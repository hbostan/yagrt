package yagrt

// Color holds three float64s representing RGB values of a pixel
type Color struct {
	R, G, B float64
}

// Add sums two colors channel-wise
func (c Color) Add(o Color) Color {
	return Color{c.R + o.R, c.G + o.G, c.B + o.B}
}

// Sub subtracts two colors channel-wise
func (c Color) Sub(o Color) Color {
	return Color{c.R - o.R, c.G - o.G, c.B - o.B}
}

// Mul multiplies a color with a scalar value
func (c Color) Mul(m float64) Color {
	return Color{c.R * m, c.G * m, c.B * m}
}

// MulColor multiples a color with another color channel-wise
func (c Color) MulColor(o Color) Color {
	return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}

// Div divides a color by a scalar value
func (c Color) Div(d float64) Color {
	return Color{c.R / d, c.G / d, c.B / d}
}
