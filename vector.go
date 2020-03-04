package yagrt

import (
	"math"
)

// Vector is a 3D vector containing 3 float64s for each dimension
type Vector struct {
	X, Y, Z float64
}

// Length returns the length of a vector
func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Dot calculates the dot product of two vectors
func (v Vector) Dot(o Vector) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Cross calculates the cross product of two vectors (v.Cross(o) = v X o)
func (v Vector) Cross(o Vector) Vector {
	a := v.Y*o.Z - v.Z*o.Y
	b := v.Z*o.X - v.X*o.Z
	c := v.X*o.Y - v.Y*o.X
	return Vector{a, b, c}
}

// Normalize normalizes the vector by dividing it by its length
func (v Vector) Normalize() Vector {
	return Vector{v.X, v.Y, v.Z}.Div(v.Length())
}

// Add adds two vectors together
func (v Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Sub subtracts two vectors
func (v Vector) Sub(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Mul multiplies the vector with the given multiplier.
func (v Vector) Mul(m float64) Vector {
	return Vector{v.X * m, v.Y * m, v.Z * m}
}

// DivScalar divies the vector by the given divisor.
func (v Vector) Div(d float64) Vector {
	return Vector{v.X / d, v.Y / d, v.Z / d}
}

func (a Vector) MulVector(b Vector) Vector {
	return Vector{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (a Vector) DivVector(b Vector) Vector {
	return Vector{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

func (a Vector) Min(b Vector) Vector {
	return Vector{math.Min(a.X, b.X), math.Min(a.Y, b.Y), math.Min(a.Z, b.Z)}
}

func (a Vector) Max(b Vector) Vector {
	return Vector{math.Max(a.X, b.X), math.Max(a.Y, b.Y), math.Max(a.Z, b.Z)}
}

func (n Vector) Reflect(i Vector) Vector {
	return i.Sub(n.Mul(2 * n.Dot(i)))
}

// Negate flips the sign of a vector, same as Vector.Mul(-1)
func (v Vector) Negate() Vector {
	return Vector{-v.X, -v.Y, -v.Z}
}
