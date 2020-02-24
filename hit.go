package yagrt

// Hit contains information about an intersection point
type Hit struct {
	T      float64
	Shape  Shape
	Normal Vector
}
