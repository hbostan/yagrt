package yagrt

// Shape is an object a ray can intersect with
type Shape interface {
	// Incersects a ray with a shape and returns the Hit data.
	Intersect(r Ray) *Hit
	// Material returns the material of a shape.
	Material() *Material
}
