package yagrt

import "math"

// Scene describes a scene with cameras, lights, shapes and their materials
type Scene struct {
	BackgroundColor Color
	Cameras         []Camera
	AmbientLight    Color
	PointLights     []PointLight
	Materials       []Material
	VertexData      []Vector
	Shapes          []Shape
}

// Intersect tries to intersect a given ray with each object in the scene and
// returns the closest intersection as a Hit struct
func (s *Scene) Intersect(r Ray) *Hit {
	var closestHit *Hit = nil
	for _, shape := range s.Shapes {
		hit := shape.Intersect(r)
		if hit != nil && hit.T > HitEpsilon && (closestHit == nil || hit.T < closestHit.T) {
			closestHit = hit
		}
	}
	return closestHit
}

// Sample is used to sample a scene to get a color for that sample
func (s *Scene) Sample(r Ray) Color {
	col := s.BackgroundColor
	if hit := s.Intersect(r); hit != nil {
		intersect := r.Origin.Add(r.Dir.Mul(hit.T))
		normal := hit.Normal
		mat := hit.Shape.Material()

		// Ambient
		col = col.Add(mat.AmbientReflectance.MulColor(s.AmbientLight))
		// Light it up!
		for _, light := range s.PointLights {
			intersect = intersect.Add(normal.Mul(ShadowEpsilon))
			lightVector := light.Position.Sub(intersect)
			lightDirection := lightVector.Normalize()
			lightDistance := lightVector.Length()
			shadowRay := Ray{intersect, lightDirection}
			if shadowHit := s.Intersect(shadowRay); shadowHit != nil {
				shadowDist := shadowRay.Dir.Mul(shadowHit.T).Length()
				if shadowDist > ShadowEpsilon && shadowDist < lightDistance-ShadowEpsilon {
					continue
				}
			}
			// Diffuse
			diffCosTheta := math.Max(0, lightDirection.Dot(normal))
			attenuatedLight := light.Intensity.Div(lightDistance * lightDistance)
			col = col.Add(mat.DiffuseReflectance.MulColor(attenuatedLight).Mul(diffCosTheta))
			// Specular
			halfVector := lightDirection.Sub(r.Dir.Normalize()).Normalize()
			specCosTheta := math.Max(0, halfVector.Dot(normal))
			col = col.Add(mat.SpecularReflectance.MulColor(attenuatedLight).Mul((math.Pow(specCosTheta, mat.PhongExponent)) / lightDistance * lightDistance))
		}
	}
	return col
}
