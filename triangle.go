package yagrt

// Mesh defies a set of Triangles and a material for them
type Mesh struct {
	Triangles []Triangle
	Mat       Material
}

// Material returns the material of a mesh
func (m *Mesh) Material() *Material {
	return &m.Mat
}

// Intersect returns Hit data for a given ray
func (m *Mesh) Intersect(r Ray) *Hit {
	var closestHit *Hit
	for _, triangle := range m.Triangles {
		if hit := triangle.Intersect(r); hit != nil && (closestHit == nil || hit.T < closestHit.T) {
			closestHit = hit
			closestHit.Shape = m
		}
	}
	return closestHit
}

// Triangle consists of three points and a material
type Triangle struct {
	V0, V1, V2 Vector
	Mat        Material
}

// Intersect returns the Hit data for a given ray
func (t *Triangle) Intersect(r Ray) *Hit {
	edge1 := t.V0.Sub(t.V1)
	edge2 := t.V0.Sub(t.V2)
	replaceVec := t.V0.Sub(r.Origin)

	// If we hit the backface of a triangle return
	if r.Dir.Dot(t.Normal()) > HitEpsilon {
		return nil
	}

	area := Determinant(&edge1, &edge2, &r.Dir)
	beta := Determinant(&replaceVec, &edge2, &r.Dir) / area
	if beta < 0 {
		return nil
	}
	gamma := Determinant(&edge1, &replaceVec, &r.Dir) / area
	if gamma < 0 || beta+gamma > 1 {
		return nil
	}
	// For smooth shading
	//alpha := 1 - beta - gamma
	d := Determinant(&edge1, &edge2, &replaceVec) / area

	if d < HitEpsilon {
		return nil
	}

	return &Hit{T: d, Shape: t, Normal: t.Normal()}

}

// Normal calculates the normal vector for a triangle
func (t *Triangle) Normal() Vector {
	e1 := t.V1.Sub(t.V0)
	e2 := t.V2.Sub(t.V0)
	return e1.Cross(e2).Normalize()
}

// Material returns the material of a triangle
func (t *Triangle) Material() *Material {
	return &t.Mat
}
