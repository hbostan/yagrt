package yagrt

// Camera is the point where we look into the scene
// it holds a position, and the u, v, w vectors.
type Camera struct {
	Position  Vector
	U         Vector
	V         Vector
	W         Vector
	NearPlane struct {
		Left   float64
		Right  float64
		Bottom float64
		Top    float64
	}
	Distance   float64
	Resolution struct {
		Width  int
		Height int
	}
	ImageName     string
	TopLeftCorner Vector
}

// LookAt recomputes the coordinate system of the camera from gaze and up
func (c *Camera) LookAt(p, gaze, up Vector) {
	c.Position = p
	c.W = gaze.Sub(c.Position).Normalize().Negate()
	c.U = up.Cross(c.W).Normalize()
	c.V = c.W.Cross(c.U).Normalize()
	c.TopLeftCorner = c.Position.Sub(c.W.Mul(c.Distance)).Add(c.U.Mul(c.NearPlane.Left)).Add(c.V.Mul(c.NearPlane.Top))
}

// CastRay creates a Ray from the position of the camera to the (x,y) on
// image plane
// TODO: Check type conversions
func (c *Camera) CastRay(x, y int) Ray {
	su := (c.NearPlane.Right - c.NearPlane.Left) * (float64(x) + 0.5) / float64(c.Resolution.Width)
	sv := (c.NearPlane.Top - c.NearPlane.Bottom) * (float64(y) + 0.5) / float64(c.Resolution.Height)
	trg := c.TopLeftCorner.Add(c.U.Mul(su)).Sub(c.V.Mul(sv))
	dir := trg.Sub(c.Position)
	return Ray{c.Position, dir}
}
