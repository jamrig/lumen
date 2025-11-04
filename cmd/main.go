package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/jamrig/lumen/internal/lumen"
	"github.com/jamrig/lumen/internal/lumen/maths"
	"github.com/jamrig/lumen/internal/lumen/shapes"
)

func main() {
	camera := lumen.NewCamera()

	scene := lumen.NewScene()
	scene.Add(shapes.NewSphere(maths.NewVec3(0, 0, -1), 0.5))
	scene.Add(shapes.NewSphere(maths.NewVec3(0, -100.5, -1), 100))

	startTime := time.Now()
	img := camera.Render(scene)
	fmt.Printf("Rendering took %s\n", time.Since(startTime))

	file, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("failed to encode image: %v", err)
	}
}
