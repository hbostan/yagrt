package yagrt

// Material contains different values used in shading
type Material struct {
	AmbientReflectance  Color
	DiffuseReflectance  Color
	SpecularReflectance Color
	PhongExponent       float64
}
