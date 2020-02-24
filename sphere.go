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
	if d < HitEpsilon {
		return nil
	}
	d = math.Sqrt(d)
	t1 := (-b + d) / (2 * a)
	t2 := (-b - d) / (2 * a)
	var small float64
	var big float64
	if small = t1; t2 < t1 {
		small = t2
	}
	if big = t1; t2 > t1 {
		big = t2
	}

	if small < HitEpsilon {
		if big < HitEpsilon {
			return nil
		}
		small = big
	}

	p := r.Origin.Add(r.Dir.Mul(small))
	n := s.Normal(p)
	return &Hit{T: small, Shape: s, Normal: n}
}

// Normal returns the sufrace normal of a sphere
func (s *Sphere) Normal(p Vector) Vector {
	return p.Sub(s.Origin).Normalize()
}

// Material returns the material of a sphere
func (s *Sphere) Material() *Material {
	return &s.Mat
}
