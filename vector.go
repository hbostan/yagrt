package yagrt

import "math"

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Dot(o Vector) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector) Cross(o Vector) Vector {
	x := v.Y*o.Z - v.Z*o.Y
	y := v.Z*o.X - v.X*o.Z
	z := v.X*o.Y - v.Y*o.X
	return Vector{x, y, z}
}

func (v Vector) Normalize() Vector {
	d := v.Length()
	return Vector{v.X / d, v.Y / d, v.Z / d}
}

func (v Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

func (v Vector) Sub(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

func (v Vector) Mul(m float64) Vector {
	return Vector{v.X * m, v.Y * m, v.Z * m}
}

func (v Vector) Div(d float64) Vector {
	return Vector{v.X / d, v.Y / d, v.Z / d}
}
