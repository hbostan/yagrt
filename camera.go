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
	topLeftCorner Vector
	pixelWidth    float64
	pixelHeight   float64
}

// LookAt recomputes the coordinate system of the camera from gaze and up
func (c *Camera) LookAt(p, gaze, up Vector) {
	gaze = gaze.Normalize()
	up = up.Normalize()
	c.Position = p
	c.W = gaze.Negate().Normalize() // Careful Here
	c.U = gaze.Cross(up).Normalize()
	c.V = c.U.Cross(gaze).Normalize()
	planeMid := c.W.Mul(c.Distance)
	leftSide := c.U.Mul(c.NearPlane.Left)
	topSide := c.V.Mul(c.NearPlane.Top)
	c.topLeftCorner = c.Position.Sub(planeMid).Add(leftSide).Add(topSide)
	c.pixelWidth = (c.NearPlane.Right - c.NearPlane.Left) / float64(c.Resolution.Width)
	c.pixelHeight = (c.NearPlane.Top - c.NearPlane.Bottom) / float64(c.Resolution.Height)
}

// CastRay creates a Ray from the position of the camera to the (x,y) on image plane
func (c *Camera) CastRay(x, y int) Ray {
	xOffset := c.U.Mul(c.pixelWidth * (0.5 + float64(x)))
	yOffset := c.V.Mul(c.pixelHeight * (0.5 + float64(y)))
	dst := c.topLeftCorner.Add(xOffset).Sub(yOffset)
	dir := dst.Sub(c.Position).Normalize()
	return Ray{c.Position, dir}
}
