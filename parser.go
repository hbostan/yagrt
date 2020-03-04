package yagrt

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/beevik/etree"
)

// ParseScene parses the given xml scene and fill in a Scene struct
// TODO: Fix this crappy code
func ParseScene(filename string) *Scene {
	var backgroundColor, ambientLight Color
	var cameras []Camera
	var pointLights []PointLight
	var materials []Material
	var vertexData []Vector
	var shapes []Shape

	doc := etree.NewDocument()
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Cannot find XML:", filename)
	}
	if err := doc.ReadFromFile(filename); err != nil {
		panic(err)
	}
	root := doc.SelectElement("Scene")

	// Background Color
	elem := root.SelectElement("BackgroundColor")
	if elem != nil {
		col, err := colorFromString(strings.TrimSpace(elem.Text()))
		if err != nil {
			fmt.Println("Cannot set BackgroundColor")
		}
		backgroundColor = col
	}

	// Shadow Epsilon
	elem = root.SelectElement("ShadowRayEpsilon")
	if elem != nil {
		eps, err := ParseFloat(strings.TrimSpace(elem.Text()))
		if err != nil {
			fmt.Println("Cannot set BackgroundColor")
		}
		ShadowEpsilon = eps
	}

	// IntersectionTestEpsilon
	elem = root.SelectElement("IntersectionTestEpsilon")
	if elem != nil {
		eps, err := ParseFloat(strings.TrimSpace(elem.Text()))
		if err != nil {
			fmt.Println("Cannot set BackgroundColor")
		}
		HitEpsilon = eps
	}
	// Cameras
	elem = root.SelectElement("Cameras")
	cams := elem.SelectElements("Camera")
	for _, cam := range cams {
		camera := parseCamera(cam)
		cameras = append(cameras, camera)
	}

	elem = root.SelectElement("Lights")
	// Ambient Light
	amblight := elem.SelectElement("AmbientLight")
	amb, err := colorFromString(amblight.Text())
	if err != nil {
		fmt.Println("Cannot parse Ambient Light")
	}
	ambientLight = amb
	// Lights
	lights := elem.SelectElements("PointLight")
	for _, l := range lights {
		light := parseLight(l)
		pointLights = append(pointLights, light)
	}

	// Materials
	elem = root.SelectElement("Materials")
	mats := elem.SelectElements("Material")
	for _, mat := range mats {
		material := parseMaterial(mat)
		materials = append(materials, material)
	}

	// VertexData
	elem = root.SelectElement("VertexData")
	vertices := strings.Split(strings.TrimSpace(elem.Text()), "\n")
	for i, vert := range vertices {
		vert = strings.TrimSpace(vert)
		if len(vert) == 0 {
			continue
		}
		vec, err := vectorFromString(vert)
		if err != nil {
			fmt.Printf("Cannot parse %v Vertex %v", i, err)
			continue
		}
		vertexData = append(vertexData, vec)
	}

	// Objects
	elem = root.SelectElement("Objects")
	spheres := elem.SelectElements("Sphere")
	for _, sph := range spheres {
		sphere := parseSphere(sph, vertexData, materials)
		shapes = append(shapes, sphere)
	}

	triangles := elem.SelectElements("Triangle")
	for _, trg := range triangles {
		triangle := parseTriangle(trg, vertexData, materials)
		shapes = append(shapes, triangle)
	}

	meshes := elem.SelectElements("Mesh")
	for _, msh := range meshes {
		mesh := parseMesh(msh, vertexData, materials)
		shapes = append(shapes, mesh)
	}
	return NewScene(
		backgroundColor,
		cameras,
		ambientLight,
		pointLights,
		materials,
		vertexData,
		shapes,
	)
}

func parseMesh(msh *etree.Element, vertexData []Vector, materials []Material) *Mesh {
	elem := msh.SelectElement("Material")
	mat, err := strconv.ParseInt(strings.TrimSpace(elem.Text()), 10, 64)
	if err != nil {
		fmt.Println("Cannot parse mesh material")
	}
	meshmat := materials[int(mat)-1]
	var triangles []*Triangle
	faces := strings.Split(strings.TrimSpace(msh.SelectElement("Faces").Text()), "\n")
	for i, face := range faces {
		indices := strings.Fields(strings.TrimSpace(face))
		v1, err1 := strconv.ParseInt(indices[0], 10, 64)
		v2, err2 := strconv.ParseInt(indices[1], 10, 64)
		v3, err3 := strconv.ParseInt(indices[2], 10, 64)
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Printf("Cannot parse indices of face with id %v\n", i)
		}
		triangles = append(triangles, NewTriangle(vertexData[int(v1)-1], vertexData[int(v2)-1], vertexData[int(v3)-1], materials[int(mat)-1]))
	}
	return NewMesh(triangles, meshmat)
}

func parseTriangle(trg *etree.Element, vertexData []Vector, materials []Material) *Triangle {

	elem := trg.SelectElement("Material")
	mat, err := strconv.ParseInt(strings.TrimSpace(elem.Text()), 10, 64)
	if err != nil {
		fmt.Println("Cannot parse triangle material")
	}
	elem = trg.SelectElement("Indices")
	indices := strings.Fields(strings.TrimSpace(elem.Text()))
	v1, err1 := strconv.ParseInt(indices[0], 10, 64)
	v2, err2 := strconv.ParseInt(indices[1], 10, 64)
	v3, err3 := strconv.ParseInt(indices[2], 10, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Printf("Cannot parse indices of triangle with id %v\n", trg.SelectAttr("id"))
	}

	return NewTriangle(vertexData[int(v1)-1], vertexData[int(v2)-1], vertexData[int(v3)-1], materials[int(mat)-1])
}

func parseSphere(sph *etree.Element, vertexData []Vector, materials []Material) *Sphere {
	elem := sph.SelectElement("Material")
	mat, err := strconv.ParseInt(strings.TrimSpace(elem.Text()), 10, 64)
	if err != nil {
		fmt.Println("Cannot parse sphere material")
	}
	elem = sph.SelectElement("Radius")
	rad, err := ParseFloat(strings.TrimSpace(elem.Text()))
	if err != nil {
		fmt.Println("Cannot parse sphere Radius")
	}
	elem = sph.SelectElement("Center")
	cen, err := strconv.ParseInt(strings.TrimSpace(elem.Text()), 10, 64)
	if err != nil {
		fmt.Println("Cannot parse sphere Center")
	}
	return NewSphere(vertexData[cen-1], rad, materials[(int(mat)-1)])
}

func parseMaterial(mat *etree.Element) Material {
	var material Material
	var ar, dr, sr Color
	var ph float64
	var err error
	elem := mat.SelectElement("AmbientReflectance")
	if elem != nil {
		ar, err = colorFromString(strings.TrimSpace(elem.Text()))
		if err != nil {
			fmt.Println("Cannot parse AmbientReflectance of Material")
		}
	} else {
		fmt.Println("No AmbientReflectance using (0, 0, 0)")
	}
	elem = mat.SelectElement("DiffuseReflectance")
	if elem != nil {
		dr, err = colorFromString(strings.TrimSpace(elem.Text()))
		if err != nil {
			fmt.Println("Cannot parse DiffuseReflectance of Material")
		}
	} else {
		fmt.Println("No DiffuseReflectance using (0, 0, 0)")
	}
	elem = mat.SelectElement("SpecularReflectance")
	if elem != nil {
		sr, err = colorFromString(strings.TrimSpace(elem.Text()))
		if err != nil {
			fmt.Println("Cannot parse SpecularReflectance of Material")
		}
	} else {
		fmt.Println("No SpecularReflectance using (0, 0, 0)")
	}
	elem = mat.SelectElement("PhongExponent")
	if elem != nil {
		ph, err = ParseFloat(strings.TrimSpace(elem.Text()))
		if err != nil {
			fmt.Println("Cannot parse PhongExponent of Material")
		}
	}

	material.AmbientReflectance = ar
	material.DiffuseReflectance = dr
	material.SpecularReflectance = sr
	material.PhongExponent = ph
	return material
}

func parseLight(light *etree.Element) PointLight {
	var pointLight PointLight
	elem := light.SelectElement("Position")
	pos, err := vectorFromString(strings.TrimSpace(elem.Text()))
	if err != nil {
		fmt.Println("Cannot parse Light position")
	}
	elem = light.SelectElement("Intensity")
	ints, err := colorFromString(strings.TrimSpace(elem.Text()))
	if err != nil {
		fmt.Println("Cannot parse Light Intensity")
	}
	pointLight.Position = pos
	pointLight.Intensity = ints
	return pointLight
}

func parseCamera(cam *etree.Element) Camera {
	var camera Camera
	var pos, gaze, up Vector
	//fmt.Printf("%v (id=%v)\n", cam.Tag, cam.SelectAttr("id").Value)
	// Position
	camElem := cam.SelectElement("Position")
	if camElem != nil {
		vec, err := vectorFromString(strings.TrimSpace(camElem.Text()))
		if err != nil {
			fmt.Println("Cannot get camera position")
		}
		pos = vec
	}
	// Gaze
	camElem = cam.SelectElement("Gaze")
	if camElem != nil {
		vec, err := vectorFromString(strings.TrimSpace(camElem.Text()))
		if err != nil {
			fmt.Println("Cannot get camera position")
		}
		gaze = vec
	}
	// Up
	camElem = cam.SelectElement("Up")
	if camElem != nil {
		vec, err := vectorFromString(strings.TrimSpace(camElem.Text()))
		if err != nil {
			fmt.Println("Cannot get camera position")
		}
		up = vec
	}

	// NearPlane
	camElem = cam.SelectElement("NearPlane")
	if camElem != nil {
		nums := strings.Fields(strings.TrimSpace(camElem.Text()))
		if len(nums) != 4 {
			fmt.Println("NearPlane should consists of l, r, b, t elements")
		}
		var l, r, b, t float64
		var err error
		if l, err = ParseFloat(nums[0]); err != nil {
			fmt.Println("Nearplane Left value parse error")
		}
		if r, err = ParseFloat(nums[1]); err != nil {
			fmt.Println("Nearplane Right value parse error")
		}
		if b, err = ParseFloat(nums[2]); err != nil {
			fmt.Println("Nearplane Bottom value parse error")
		}
		if t, err = ParseFloat(nums[3]); err != nil {
			fmt.Println("Nearplane Top value parse error")
		}
		camera.NearPlane.Left = l
		camera.NearPlane.Right = r
		camera.NearPlane.Bottom = b
		camera.NearPlane.Top = t
	}

	// NearDistance
	camElem = cam.SelectElement("NearDistance")
	if camElem != nil {
		d, err := ParseFloat(strings.TrimSpace(camElem.Text()))
		if err != nil {
			fmt.Println("Cannot parse NearDistance")
		}
		camera.Distance = d
	}

	// ImageResolution
	camElem = cam.SelectElement("ImageResolution")
	if camElem != nil {
		nums := strings.Fields(strings.TrimSpace(camElem.Text()))
		if len(nums) != 2 {
			fmt.Println("Resolution must be 2D")
		}
		width, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			fmt.Println("Cannot parse Width of Resolution")
		}
		height, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			fmt.Println("Cannot parse Height of Resolution")
		}
		camera.Resolution.Width = int(width)
		camera.Resolution.Height = int(height)
	}

	// ImageName
	camElem = cam.SelectElement("ImageName")
	if camElem != nil {
		camera.ImageName = strings.TrimSpace(camElem.Text())
	}
	// INITIALIZE CAMERA DIRECTION
	camera.LookAt(pos, gaze, up)
	return camera
}

func vectorFromString(s string) (Vector, error) {
	nums := strings.Fields(s)
	if len(nums) != 3 {
		return Vector{}, fmt.Errorf("Cannot convert to vector without 3 elements")
	}
	var x, y, z float64
	var err error
	if x, err = ParseFloat(nums[0]); err != nil {
		return Vector{}, fmt.Errorf("Cannot parse X value of Vector")
	}
	if y, err = ParseFloat(nums[1]); err != nil {
		return Vector{}, fmt.Errorf("Cannot parse Y value of Vector")
	}
	if z, err = ParseFloat(nums[2]); err != nil {
		return Vector{}, fmt.Errorf("Cannot parse Z value of Vector")
	}
	return Vector{x, y, z}, nil
}

func colorFromString(s string) (Color, error) {
	nums := strings.Fields(s)
	if len(nums) != 3 {
		return Color{}, fmt.Errorf("Cannot convert to vector without 3 elements")
	}
	var r, g, b float64
	var err error
	if r, err = ParseFloat(nums[0]); err != nil {
		return Color{}, fmt.Errorf("Cannot parse X value of Vector")
	}
	if g, err = ParseFloat(nums[1]); err != nil {
		return Color{}, fmt.Errorf("Cannot parse Y value of Vector")
	}
	if b, err = ParseFloat(nums[2]); err != nil {
		return Color{}, fmt.Errorf("Cannot parse Z value of Vector")
	}
	return Color{r, g, b}, nil
}
