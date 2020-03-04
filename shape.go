package yagrt

// Shape is an object a ray can intersect with
type Shape interface {
	// Incersects a ray with a shape and returns the Hit data.
	Intersect(r Ray, hit *Hit) bool
	// Material returns the material of a shape.
	Material() *Material
	BoundingBox() Box
}
