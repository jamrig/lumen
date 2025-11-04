package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/jamrig/lumen/internal/lumen"
)

func main() {
	camera := lumen.NewCamera()

	world := lumen.NewHittableList()
	world.Add(lumen.NewSphere(lumen.NewVec3(0, 0, -1), 0.5))
	world.Add(lumen.NewSphere(lumen.NewVec3(0, -100.5, -1), 100))

	startTime := time.Now()
	img := camera.Render(world)
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
