package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/hbostann/yagrt"
)

func main() {
	scene := yagrt.Scene{}
	scene.AddShape(&yagrt.Sphere{yagrt.Vector{-1, 1, 2}, 1, yagrt.Color{1, 0, 0}})
	scene.AddShape(&yagrt.Sphere{yagrt.Vector{0, 0, 0}, 1, yagrt.Color{0, 1, 0}})
	scene.AddShape(&yagrt.Sphere{yagrt.Vector{1, -1, 3}, 1, yagrt.Color{0, 0, 1}})
	scene.AddLight(&yagrt.Sphere{yagrt.Vector{-2.5, -2.5, -1}, 0.1, yagrt.Color{1, 1, 1}})

	camera := yagrt.Camera{}
	camera.LookAt(yagrt.Vector{0, 0, -5}, yagrt.Vector{0, 0, 0}, yagrt.Vector{0, 1, 0}, 45)
	image := yagrt.Render(&scene, &camera, 800, 600, 16)
	file, err := os.Create("outputs/out.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, image)
	if err != nil {
		fmt.Printf("Error encoding %v\n", file)
		panic(err)
	}
}
