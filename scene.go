package yagrt

import "math"

type Scene struct {
	Shapes []Shape
	Lights []Shape
}

func (s *Scene) AddShape(shape Shape) {
	s.Shapes = append(s.Shapes, shape)
}

func (s *Scene) AddLight(light Shape) {
	s.Lights = append(s.Lights, light)
}

func (s *Scene) Intersect(r Ray, hit *Hit) bool {
	ok := false
	for _, shape := range s.Shapes {
		innerHit := Hit{}
		innerHit.T = INF
		if isHit := shape.Intersect(r, &innerHit); isHit {
			if innerHit.T < hit.T {
				*hit = innerHit
				ok = true
			}
		}
	}
	return ok
}

func (s *Scene) Shadow(r Ray) bool {
	for _, shape := range s.Shapes {
		hit := Hit{}
		if isHit := shape.Intersect(r, &hit); isHit {
			return true
		}
	}
	return false
}

func (s *Scene) Light(r Ray) Color {
	color := Color{}
	for _, light := range s.Lights {
		shadowRay := Ray{r.Origin, light.RandomPoint().Sub(r.Origin).Normalize()}
		if s.Shadow(shadowRay) {
			continue
		}
		cos := math.Max(0, shadowRay.Direction.Dot(r.Direction))
		color = color.Add(light.Color().Mul(cos))
	}
	return color
}

func (s *Scene) Sample(r Ray) Color {
	hit := Hit{}
	hit.T = INF
	if isHit := s.Intersect(r, &hit); isHit {
		return hit.Shape.Color().MulColor(s.Light(hit.Ray))
	}
	return Color{}
}
