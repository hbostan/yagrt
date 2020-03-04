package yagrt

// Mesh defies a set of Triangles and a material for them
type Mesh struct {
	Triangles []*Triangle
	Mat       Material
	Box       Box
}

func NewMesh(triangles []*Triangle, mat Material) *Mesh {
	box := triangles[0].BoundingBox()
	for _, t := range triangles {
		box = box.Extend(t.BoundingBox())
	}
	return &Mesh{triangles, mat, box}
}

func (m *Mesh) BoundingBox() Box {
	return m.Box
}

// Material returns the material of a mesh
func (m *Mesh) Material() *Material {
	return &m.Mat
}

// Intersect returns Hit data for a given ray
func (m *Mesh) Intersect(r Ray, hit *Hit) bool {
	var closestHit Hit
	rayHit := false
	for _, triangle := range m.Triangles {
		var innerHit Hit
		if isHit := triangle.Intersect(r, &innerHit); isHit && (!rayHit || innerHit.T < closestHit.T) {
			rayHit = true
			closestHit = innerHit
		}
	}
	hit.T = closestHit.T
	hit.Shape = m
	hit.Normal = closestHit.Normal
	return true
}

// Triangle consists of three points and a material
type Triangle struct {
	V0, V1, V2 Vector
	Mat        Material
	Normal     Vector
	Box        Box
}

func NewTriangle(v0, v1, v2 Vector, mat Material) *Triangle {
	normal := v1.Sub(v0).Cross(v2.Sub(v0)).Normalize()
	min := v0.Min(v1).Min(v2)
	max := v0.Max(v1).Max(v2)
	box := Box{min, max}
	return &Triangle{v0, v1, v2, mat, normal, box}
}

// Intersect returns the Hit data for a given ray
func (t *Triangle) Intersect(r Ray, hit *Hit) bool {
	edge1 := t.V0.Sub(t.V1)
	edge2 := t.V0.Sub(t.V2)
	replaceVec := t.V0.Sub(r.Origin)

	// If we hit the backface of a triangle return
	if r.Dir.Dot(t.Normal) > HitEpsilon {
		return false
	}

	area := Determinant(&edge1, &edge2, &r.Dir)
	beta := Determinant(&replaceVec, &edge2, &r.Dir) / area
	if beta < 0 {
		return false
	}
	gamma := Determinant(&edge1, &replaceVec, &r.Dir) / area
	if gamma < 0 || beta+gamma > 1 {
		return false
	}
	// For smooth shading
	//alpha := 1 - beta - gamma
	d := Determinant(&edge1, &edge2, &replaceVec) / area

	if d < HitEpsilon {
		return false
	}
	hit.T = d
	hit.Shape = t
	hit.Normal = t.Normal
	return true

}

func (t *Triangle) BoundingBox() Box {
	return t.Box
}

// Normal calculates the normal vector for a triangle
func (t *Triangle) GetNormal() Vector {
	e1 := t.V1.Sub(t.V0)
	e2 := t.V2.Sub(t.V0)
	return e1.Cross(e2).Normalize()
}

// Material returns the material of a triangle
func (t *Triangle) Material() *Material {
	return &t.Mat
}
