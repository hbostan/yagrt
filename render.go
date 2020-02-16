package yagrt

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func Render(sceneFile string) {
	scene := ParseScene(sceneFile)
	for i, camera := range scene.Cameras {
		fmt.Printf("Rendering %v. Camera (%v)", i, camera.ImageName)
		f, err := os.OpenFile(camera.ImageName, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		image := image.NewNRGBA(image.Rect(0, 0, camera.Resolution.Width, camera.Resolution.Height))
		for y := 0; y < camera.Resolution.Height; y++ {
			for x := 0; x < camera.Resolution.Width; x++ {
				c := scene.Sample(camera.CastRay(x, y))
				r := uint8(math.Min(255, math.Max(0, c.R)))
				g := uint8(math.Min(255, math.Max(0, c.G)))
				b := uint8(math.Min(255, math.Max(0, c.B)))
				image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
			}
		}
		if err = png.Encode(f, image); err != nil {
			return
		}
	}
}
