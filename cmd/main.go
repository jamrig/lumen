package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/jamrig/lumen/internal/lumen"
	"github.com/jamrig/lumen/internal/lumen/material"
	"github.com/jamrig/lumen/internal/lumen/maths"
	"github.com/jamrig/lumen/internal/lumen/shapes"
)

func main() {
	camera := lumen.NewCamera()

	// materials
	materialGround := material.NewLambertianMaterial(maths.NewColor(0.8, 0.8, 0.0))
	materialCenter := material.NewLambertianMaterial(maths.NewColor(0.1, 0.2, 0.5))
	materialLeft := material.NewMetalMaterial(maths.NewColor(0.8, 0.8, 0.8))
	materialRight := material.NewMetalMaterial(maths.NewColor(0.8, 0.6, 0.2))

	// scene
	scene := lumen.NewScene()
	scene.Add(shapes.NewSphere(maths.NewVec3(0.0, -100.5, -1.0), 100, materialGround))
	scene.Add(shapes.NewSphere(maths.NewVec3(0.0, 0.0, -1.2), 0.5, materialCenter))
	scene.Add(shapes.NewSphere(maths.NewVec3(-1.0, 0.0, -1.0), 0.5, materialLeft))
	scene.Add(shapes.NewSphere(maths.NewVec3(1.0, 0.0, -1.0), 0.5, materialRight))

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
