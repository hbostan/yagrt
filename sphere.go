package yagrt

import "math"

// Sphere represents a sphere in 3D, it has an origin and a radius
type Sphere struct {
	Origin Vector
	Radius float64
	Mat    Material
}

// Intersect calculates the intersection point of a ray with the sphere
// It returns the distance of the intersection from the ray's origin if
// ray intersects the sphere, INF if it doesn't.
func (s *Sphere) Intersect(r Ray) *Hit {
	to := r.Origin.Sub(s.Origin)
	a := r.Dir.Dot(r.Dir)
	b := 2 * r.Dir.Dot(to)
	c := to.Dot(to) - s.Radius*s.Radius
	d := b*b - 4*a*c
	t := INF
	if d < -EPS {
		return nil
	} else if d < EPS {
		t = -b / 2 * a
	} else {
		d = math.Sqrt(d)
		t1 := (-b + d) / (2 * a)
		t2 := (-b - d) / (2 * a)
		t = math.Min(t1, t2)
		if t2 < 0 {
			t = t1
		} else {
			t = t2
		}
	}
	p := r.Origin.Add(r.Dir.Mul(t))
	n := s.Normal(p)
	return &Hit{T: t, Shape: s, Ray: Ray{p, n}}
}

func (s *Sphere) Normal(p Vector) Vector {
	return p.Sub(s.Origin).Normalize()
}

func (s *Sphere) Material() *Material {
	return &s.Mat
}
