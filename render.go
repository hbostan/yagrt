package yagrt

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sync"
	"time"
)

// SubRender is a Goroutine which renders the given part of an image
func SubRender(xStart int, yStart int, xEnd int, yEnd int, scene *Scene, camera *Camera, image *image.NRGBA, wg *sync.WaitGroup) {
	defer wg.Done()
	for y := yStart; y < yEnd; y++ {
		for x := xStart; x < xEnd; x++ {
			c := scene.Sample(camera.CastRay(x, y))
			r := uint8(math.Min(255, math.Max(0, c.R)))
			g := uint8(math.Min(255, math.Max(0, c.G)))
			b := uint8(math.Min(255, math.Max(0, c.B)))
			image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
		}
	}
}

// Render parses the scene file and then spawns goroutines for each line to quickly
// a scene, in the end it encodes the image into a png with the same name as the
// scene file.
func Render(sceneFile string) {
	var wg sync.WaitGroup
	fmt.Println("Parsing Scene")
	start := time.Now()
	scene := ParseScene(sceneFile)
	scene.BVH.DebugPrint(0)
	fmt.Printf("Parsing Done: %v\n", time.Since(start))

	for i, camera := range scene.Cameras {
		fmt.Printf("Rendering Camera %v (%v)\n", i+1, camera.ImageName)
		f, err := os.OpenFile("outputs/"+camera.ImageName, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		image := image.NewNRGBA(image.Rect(0, 0, camera.Resolution.Width, camera.Resolution.Height))
		for y := 0; y < camera.Resolution.Height; y++ {
			wg.Add(1)
			//go SubRender(0, y, camera.Resolution.Width, y+1, scene, &camera, image, &wg)
			SubRender(0, y, camera.Resolution.Width, camera.Resolution.Height, scene, &camera, image, &wg)
		}
		wg.Wait()
		if err = png.Encode(f, image); err != nil {
			return
		}
	}
}
