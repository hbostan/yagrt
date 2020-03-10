package yagrt

type Shape interface {
	Intersect(r Ray, hit *Hit) bool
	Color() Color
	Normal(Vector) Vector
	RandomPoint() Vector
}
