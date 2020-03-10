package yagrt

import (
	"math"
	"math/rand"
)

type Sphere struct {
	Center Vector
	Radius float64
	Col    Color
}

func (s *Sphere) Intersect(r Ray, hit *Hit) bool {
	to := r.Origin.Sub(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2 * to.Dot(r.Direction)
	c := to.Dot(to) - s.Radius*s.Radius
	d := b*b - 4*a*c
	if d > 0 {
		t := (-b - math.Sqrt(d)) / (2 * a)
		if t > 0 {
			p := r.Origin.Add(r.Direction.Mul(t))
			hit.T = t
			hit.Ray = Ray{p, s.Normal(p)}
			hit.Shape = s
			return true
		}
	}
	return false
}

func (s *Sphere) Color() Color {
	return s.Col
}

func (s *Sphere) Normal(position Vector) Vector {
	return position.Sub(s.Center).Normalize()
}

func (s *Sphere) RandomPoint() Vector {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		z := rand.Float64()*2 - 1
		v := Vector{x, y, z}
		if v.Length() <= 1 {
			return v.Mul(s.Radius).Add(s.Center)
		}
	}
}
