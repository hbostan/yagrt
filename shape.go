package yagrt

type Shape interface {
	Intersect(r Ray) *Hit
	Material() *Material
}
