package yagrt

// Camera is the point where we look into the scene
// it holds a position, and the u, v, w vectors.
type Camera struct {
	Position Vector
	U        Vector
	V        Vector
	W        Vector
}

// LookAt recomputes the coordinate system of the camera from gaze and up
func (c *Camera) LookAt(p, gaze, up Vector) {
	c.Position = p
	c.W = gaze.Sub(p).Normalize()
	c.U = up.Cross(c.W).Normalize()
	c.V = c.W.Cross(c.U).Normalize()
}

// CastRay creates a Ray from the position of the camera to the (x,y) on
// image plane
func (c *Camera) CastRay(x, y, w, h int) Ray {
	aspect := float64(w) / float64(h)
	px := (float64(x)/(float64(w)-1))*2 - 1
	py := (float64(y)/(float64(h)-1))*2 - 1
	d := Vector{}
	d = d.Add(c.U.MulScalar(px * aspect))
	d = d.Add(c.V.MulScalar(py))
	d = d.Add(c.W.MulScalar(2.5))
	return Ray{c.Position, d}
}
