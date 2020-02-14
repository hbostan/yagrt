package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/hbostann/yagrt"
)

func main() {
	sphere := yagrt.Sphere{yagrt.Vector{}, 1}
	camera := yagrt.Camera{}
	camera.LookAt(yagrt.Vector{0, 0, -5}, yagrt.Vector{}, yagrt.Vector{0, 1, 0})
	w, h := 1024, 1024

	f, err := os.OpenFile("out.png", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	image := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			image.SetNRGBA(x, y, color.NRGBA{0, 0, 0, 255})
			ray := camera.CastRay(x, y, w, h)
			t := sphere.Intersect(ray)
			if t < yagrt.INF {
				r := uint8(rand.Intn(255))
				g := uint8(rand.Intn(255))
				b := uint8(rand.Intn(255))
				image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
			}
		}
	}
	if err = png.Encode(f, image); err != nil {
		return
	}
}
