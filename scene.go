package yagrt

import (
	"image"
	"image/color"
	"math"
)

type Scene struct {
	BackgroundColor  Color
	ShadowRayEpsilon float64
	IntersectEpsilon float64
	Cameras          []Camera
	AmbientLight     Color
	PointLights      []PointLight
	Materials        []Material
	VertexData       []Vector
	Shapes           []Shape
}

func (s Scene) Render(camIdx int, image *image.NRGBA) {
	camera := s.Cameras[camIdx]

	for y := 0; y < camera.Resolution.Height; y++ {
		for x := 0; x < camera.Resolution.Width; x++ {
			var r, g, b uint8
			r, g, b = 0, 0, 0
			hit := s.Intersect(camera.CastRay(x, y))
			if hit != nil {
				r = uint8((float64(y) / float64(camera.Resolution.Height)) * 255)
				g = uint8((float64(x) / float64(camera.Resolution.Width)) * 255)
				b = uint8(((float64(y) / float64(camera.Resolution.Height)) + (float64(x) / float64(camera.Resolution.Width))) / 2 * 255)
			}
			image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
		}
	}
}

func (s *Scene) Intersect(r Ray) *Hit {
	var closestHit *Hit = nil
	for _, shape := range s.Shapes {
		hit := shape.Intersect(r)
		if hit != nil && hit.T > s.IntersectEpsilon && (closestHit == nil || hit.T < closestHit.T) {
			closestHit = hit
		}
	}
	return closestHit
}

func (s *Scene) Shadow(r Ray) bool {
	for _, shape := range s.Shapes {
		hit := shape.Intersect(r)
		if hit != nil && hit.T < INF && hit.T > 0 {
			return true
		}
	}
	return false
}

func (s *Scene) Sample(r Ray) Color {
	col := Color{}
	if hit := s.Intersect(r); hit != nil {
		//fmt.Println(hit)
		intersect := hit.Ray.Origin
		normal := hit.Ray.Dir
		mat := hit.Shape.Material()
		w_0 := r.Origin.Sub(intersect).Normalize()
		// Ambient
		col = col.Add(mat.AmbientReflectance.MulColor(s.AmbientLight))
		// Light it up

		for _, light := range s.PointLights {
			lightDistance := light.Position.Sub(intersect)
			w_i := lightDistance.Normalize()
			shadowRay := Ray{intersect.Add((w_i.Mul(s.ShadowRayEpsilon))), w_i}
			if s.Shadow(shadowRay) {
				continue
			}
			lightDistanceSq := lightDistance.Length() * lightDistance.Length()
			diffCosTheta := normal.Dot(w_i)
			// Diffuse
			col = col.Add((mat.DiffuseReflectance.MulColor(light.Intensity)).Mul(diffCosTheta / lightDistanceSq))
			// Specular
			specCosTheta := math.Max(0, normal.Dot((w_0.Add(w_i)).Normalize()))
			col = col.Add((mat.SpecularReflectance.MulColor(light.Intensity)).Mul((math.Pow(specCosTheta, mat.PhongExponent)) / lightDistanceSq))
		}
	}
	return col
}
