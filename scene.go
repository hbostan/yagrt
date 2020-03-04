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
	BVH             *BVHNode
}

func NewScene(bc Color, cams []Camera, al Color, pl []PointLight, mats []Material, vdata []Vector, shapes []Shape) *Scene {
	// BVH Construction Here
	return &Scene{bc, cams, al, pl, mats, vdata, shapes, NewBVH(shapes)}
}

// Intersect tries to intersect a given ray with each object in the scene and
// returns the closest intersection as a Hit struct
func (s *Scene) Intersect(r Ray, hit *Hit) bool {
	// var closestHit Hit
	rayHit := false
	rayHit = s.BVH.Intersect(r, hit)
	// for _, shape := range s.Shapes {
	// 	var innerHit Hit
	// 	if isHit := shape.Intersect(r, &innerHit); isHit && innerHit.T > HitEpsilon && (!rayHit || innerHit.T < closestHit.T) {
	// 		rayHit = true
	// 		closestHit = innerHit
	// 	}
	// }
	// hit.T = closestHit.T
	// hit.Shape = closestHit.Shape
	// hit.Normal = closestHit.Normal
	return rayHit
}

// Sample is used to sample a scene to get a color for that sample
func (s *Scene) Sample(r Ray) Color {
	col := s.BackgroundColor
	var hit Hit
	if isHit := s.Intersect(r, &hit); isHit {
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
			var shadowHit Hit
			if inShadow := s.Intersect(shadowRay, &shadowHit); inShadow {
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
