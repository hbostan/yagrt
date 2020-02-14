package yagrt

// Ray represents a ray sent from camera into the scene
// it has an origin and a direction
type Ray struct {
	Origin Vector
	Dir    Vector
}
