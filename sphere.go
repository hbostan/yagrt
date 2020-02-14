package yagrt

import "math"

// Sphere represents a sphere in 3D, it has an origin and a radius
type Sphere struct {
	Origin Vector
	Radius float64
}

// Intersect calculates the intersection point of a ray with the sphere
// It returns the distance of the intersection from the ray's origin if
// ray intersects the sphere, INF if it doesn't.
func (s *Sphere) Intersect(r Ray) float64 {
	to := r.Origin.Sub(s.Origin)
	a := r.Dir.Dot(r.Dir)
	b := 2 * to.Dot(r.Dir)
	c := to.Dot(to) - s.Radius*s.Radius
	d := b*b - 4*a*c
	if d > 0 {
		t := (-b - math.Sqrt(d)) / (2 * a)
		if t > 0 {
			return t
		}
	}
	return INF
}
