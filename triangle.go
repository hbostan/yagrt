package yagrt

type Mesh struct {
	Triangles []Triangle
	Mat       Material
}

func (m *Mesh) Material() *Material {
	return &m.Mat
}

func (m *Mesh) Intersect(r Ray) *Hit {
	for _, triangle := range m.Triangles {
		if hit := triangle.Intersect(r); hit != nil {
			hit.Shape = m
			return hit
		}
	}
	return nil
}

type Triangle struct {
	V0, V1, V2 Vector
	Mat        Material
}

func (t *Triangle) Intersect(r Ray) *Hit {
	v1 := t.V0.Sub(t.V1)
	v2 := t.V0.Sub(t.V2)
	v3 := r.Dir
	det := Determinant(&v1, &v2, &v3)
	if det == 0 {
		return nil
	}
	b := (t.V0.Sub(r.Origin)).DivScalar(det)
	beta := Determinant(&b, &v2, &v3)
	if beta < 0 || beta > 1 {
		return nil
	}
	gamma := Determinant(&v1, &b, &v3)
	if gamma < 0 || beta+gamma > 1 {
		return nil
	}
	d := Determinant(&v1, &v2, &b)
	if d < 0 {
		return nil
	}
	p := r.Origin.Add(r.Dir.Mul(d))
	return &Hit{T: d, Shape: t, Ray: Ray{p, t.Normal()}}

}

func (t *Triangle) Normal() Vector {
	return t.V1.Sub(t.V0).Cross((t.V2.Sub(t.V0))).Normalize()
}

func (t *Triangle) Material() *Material {
	return &t.Mat
}
