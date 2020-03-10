package yagrt

import "math"

// Camera is the point where we look into the scene
// it holds a position, and the u, v, w vectors.
type Camera struct {
	Position Vector
	U, V, W  Vector
	Scale    float64
}

// LookAt recomputes the coordinate system of the camera from gaze and up
func (c *Camera) LookAt(p, gaze, up Vector, fovY float64) {
	c.Position = p
	c.W = gaze.Sub(p).Normalize()
	c.U = up.Cross(c.W).Normalize()
	c.V = c.W.Cross(c.U).Normalize()
	c.Scale = 1 / math.Tan(fovY*math.Pi/360)
}

// CastRay creates a Ray from the position of the camera to the (x,y) on image plane
func (c *Camera) CastRay(x, y, w, h int) Ray {
	aspect := float64(w) / float64(h)
	px := (float64(x)/(float64(w)-1))*2 - 1
	py := -((float64(y)/(float64(h)-1))*2 - 1)
	d := Vector{}
	d = d.Add(c.U.Mul(px * aspect))
	d = d.Add(c.V.Mul(py))
	d = d.Add(c.W.Mul(c.Scale))
	return Ray{c.Position, d}
}
