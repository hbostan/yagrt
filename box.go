package yagrt

import "math"

type Axis uint8

const (
	AxisNone Axis = iota
	AxisX
	AxisY
	AxisZ
)

type Box struct {
	Min Vector
	Max Vector
}

func (a Box) Extend(b Box) Box {
	return Box{a.Min.Min(b.Min), a.Max.Max(b.Max)}
}

func (a Box) Center() Vector {
	return a.Min.Add(a.Max).Mul(0.5)
}

func (b *Box) Intersect(r Ray, hit *Hit) bool {
	//hit.Shape = nil
	invDir := Vector{1 / r.Dir.X, 1 / r.Dir.Y, 1 / r.Dir.Z}
	n := b.Min.Sub(r.Origin).MulVector(invDir)
	f := b.Max.Sub(r.Origin).MulVector(invDir)
	n, f = n.Min(f), n.Max(f)
	t0 := math.Max(math.Max(n.X, n.Y), n.Z)
	t1 := math.Min(math.Min(f.X, f.Y), f.Z)
	hit.T = math.Min(t0, t1)
	hit.Shape = b
	return t1 < t0
}

func (b *Box) BoundingBox() Box {
	return Box{b.Min, b.Max}
}

func (b *Box) Material() *Material {
	return nil
}

func (b *Box) SplitAxis() Axis {
	width := b.Max.X - b.Min.X
	height := b.Max.Y - b.Min.Y
	depth := b.Max.Z - b.Min.Z
	axis := AxisX
	if width > height && width > depth {
		axis = AxisX
	}
	if height > width && height > depth {
		axis = AxisY
	}
	if depth > width && depth > height {
		axis = AxisZ
	}
	return axis
}
