package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/jamrig/lumen/internal/lumen"
)

func main() {

	aspectRatio := 16.0 / 9.0
	imageWidth := 1920
	imageHeight := int(math.Floor(float64(imageWidth) / aspectRatio))

	camera := lumen.NewCamera(lumen.NewVec3(0, 0, 0), 1, 2.0*(float64(imageWidth)/float64(imageHeight)), 2.0, imageWidth, imageHeight)
	fmt.Println(camera)

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	world := lumen.NewHittableList()
	world.Add(lumen.NewSphere(lumen.NewVec3(0, 0, -1), 0.5))
	world.Add(lumen.NewSphere(lumen.NewVec3(0, -100.5, -1), 100))

	for j := range imageHeight {
		for i := range imageWidth {
			img.SetRGBA(i, j, camera.ColorAtPixel(i, j, world))
		}
	}

	file, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("failed to encode image: %v", err)
	}
}
